package common

import (
	"archive/zip"
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/graph-gophers/graphql-go"
	"github.com/pkg/errors"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/csv"
	"github.com/stackrox/rox/pkg/logging"
	"github.com/stackrox/rox/pkg/stringutils"
)

// This formatter is tightly coupled to the report query. The end goal is to use the CSVPrinter in roxctl, but
// as it stands it has some  limitations which are non trivial to fix, so in the interim we will format the report using
// the formatting logic in this file

var (
	csvHeader = []string{
		"Cluster",
		"Namespace",
		"Deployment",
		"Image",
		"Component",
		"CVE",
		"Fixable",
		"CVE Fixed In",
		"Severity",
		"CVSS",
		"Discovered At",
		"Reference",
	}
	log = logging.LoggerForModule()
)

// ImageVulnerability contains image CVE data for vuln report
type ImageVulnerability struct {
	Cve               string        `json:"cve,omitempty"`
	Severity          string        `json:"severity,omitempty"`
	FixedByVersion    string        `json:"fixedByVersion,omitempty"`
	IsFixable         bool          `json:"isFixable,omitempty"`
	DiscoveredAtImage *graphql.Time `json:"discoveredAtImage,omitempty"`
	Link              string        `json:"link,omitempty"`
	Cvss              float64       `json:"cvss,omitempty"`
}

// ImageComponent data for vuln report
type ImageComponent struct {
	Name                 string                `json:"name,omitempty"`
	ImageVulnerabilities []*ImageVulnerability `json:"imageVulnerabilities,omitempty"`
	Vulns                []*ImageVulnerability `json:"vulns,omitempty"`
}

// Image data for vuln report
type Image struct {
	Name            *storage.ImageName `json:"name,omitempty"`
	ImageComponents []*ImageComponent  `json:"imageComponents,omitempty"`
	Components      []*ImageComponent  `json:"components,omitempty"`
}

// Deployment data used for generating vuln reports
type Deployment struct {
	Cluster        *storage.Cluster `json:"cluster,omitempty"`
	ClusterName    string           `json:"clusterName,omitempty"`
	Namespace      string           `json:"namespace,omitempty"`
	DeploymentName string           `json:"name,omitempty"`
	Images         []*Image         `json:"images,omitempty"`
}

// DeployedImagesResult contains results of running a single cvefields query and scope query combination on deployments graphQL schema
type DeployedImagesResult struct {
	Deployments []*Deployment `json:"deployments,omitempty"`
}

// WatchedImagesResult contains results of running a single cvefields query and scope query combination on images graphQL schema
type WatchedImagesResult struct {
	Images []*Image `json:"images,omitempty"`
}

// ZippedCSVResult contains the compressed report data and the number of CVEs found in deployed and watched images
type ZippedCSVResult struct {
	ZippedCsv            *bytes.Buffer
	NumDeployedImageCVEs int
	NumWatchedImageCVEs  int
}

// Format takes in the results of vuln report query, converts to CSV and returns zipped CSV data and
// a flag if the report is empty or not. For v1 config pass an empty string.
func Format(deployedImagesResults []DeployedImagesResult, watchedImagesResults []WatchedImagesResult, configName string) (*ZippedCSVResult, error) {
	csvWriter := csv.NewGenericWriter(csvHeader, true)
	numDeployedImageCVEs := 0
	for _, r := range deployedImagesResults {
		for _, d := range r.Deployments {
			for _, i := range d.Images {
				for _, c := range i.getComponents() {
					for _, v := range c.getVulnerabilities() {
						numDeployedImageCVEs++
						discoveredTs := "Not Available"
						if v.DiscoveredAtImage != nil {
							discoveredTs = v.DiscoveredAtImage.Time.Format("January 02, 2006")
						}
						csvWriter.AddValue(csv.Value{
							d.GetClusterName(),
							d.Namespace,
							d.DeploymentName,
							i.Name.FullName,
							c.Name,
							v.Cve,
							strconv.FormatBool(v.IsFixable),
							v.FixedByVersion,
							strings.ToTitle(stringutils.GetUpTo(v.Severity, "_")),
							strconv.FormatFloat(v.Cvss, 'f', 2, 64),
							discoveredTs,
							v.Link,
						})
					}
				}
			}
		}
	}

	numWatchedImageCVEs := 0
	for _, r := range watchedImagesResults {
		for _, i := range r.Images {
			for _, c := range i.getComponents() {
				for _, v := range c.getVulnerabilities() {
					numWatchedImageCVEs++
					discoveredTs := "Not Available"
					if v.DiscoveredAtImage != nil {
						discoveredTs = v.DiscoveredAtImage.Time.Format("January 02, 2006")
					}
					csvWriter.AddValue(csv.Value{
						"",
						"",
						"",
						i.Name.FullName,
						c.Name,
						v.Cve,
						strconv.FormatBool(v.IsFixable),
						v.FixedByVersion,
						strings.ToTitle(stringutils.GetUpTo(v.Severity, "_")),
						strconv.FormatFloat(v.Cvss, 'f', 2, 64),
						discoveredTs,
						v.Link,
					})
				}
			}
		}
	}

	var buf bytes.Buffer
	err := csvWriter.WriteBytes(&buf)
	if err != nil {
		return nil, errors.Wrap(err, "error creating csv report")
	}

	var zipBuf bytes.Buffer
	zipWriter := zip.NewWriter(&zipBuf)
	var reportName string
	if len(configName) > 80 {
		configName = configName[0:80] + "..."
		reportName = fmt.Sprintf("RHACS_Vulnerability_Report_%s_%s.csv", configName, time.Now().Format("02_January_2006"))
	} else if configName == "" {
		reportName = fmt.Sprintf("RHACS_Vulnerability_Report_%s.csv", time.Now().Format("02_January_2006"))
	} else {
		reportName = fmt.Sprintf("RHACS_Vulnerability_Report_%s_%s.csv", configName, time.Now().Format("02_January_2006"))
	}

	zipFile, err := zipWriter.Create(reportName)
	if err != nil {
		return nil, errors.Wrap(err, "unable to create a zip file of the vuln report")

	}
	_, err = zipFile.Write(buf.Bytes())
	if err != nil {
		return nil, errors.Wrap(err, "unable to create a zip file of the vuln report")
	}
	err = zipWriter.Close()
	if err != nil {
		return nil, errors.Wrap(err, "unable to create a zip file of the vuln report")
	}

	return &ZippedCSVResult{
		ZippedCsv:            &zipBuf,
		NumDeployedImageCVEs: numDeployedImageCVEs,
		NumWatchedImageCVEs:  numWatchedImageCVEs,
	}, nil
}

// GetClusterName returns name of cluster containing the Deployment
func (dep *Deployment) GetClusterName() string {
	return dep.ClusterName
}

func (img *Image) getComponents() []*ImageComponent {
	return img.ImageComponents
}

func (component *ImageComponent) getVulnerabilities() []*ImageVulnerability {
	return component.ImageVulnerabilities
}
