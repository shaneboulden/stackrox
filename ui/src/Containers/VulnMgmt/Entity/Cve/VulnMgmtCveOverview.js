import React, { useContext } from 'react';
import { ExternalLink } from 'react-feather';
import { format } from 'date-fns';

import CollapsibleSection from 'Components/CollapsibleSection';
import Metadata from 'Components/Metadata';
import LabelChip from 'Components/LabelChip';
import Widget from 'Components/Widget';
import dateTimeFormat from 'constants/dateTimeFormat';
import entityTypes from 'constants/entityTypes';
import workflowStateContext from 'Containers/workflowStateContext';
import { getSeverityChipType } from 'utils/vulnerabilityUtils';
import { truncate } from 'utils/textUtils';
import RelatedEntitiesSideList from '../RelatedEntitiesSideList';

const emptyCve = {
    componentCount: 0,
    cve: '',
    cvss: 0,
    deploymentCount: 0,
    envImpact: 0,
    fixedByVersion: '',
    imageCount: 0,
    isFixable: false,
    lastModified: '',
    lastScanned: '',
    link: '',
    publishedOn: '',
    scoreVersion: '',
    summary: '',
    vectors: {}
};

const VulnMgmtCveOverview = ({ data, entityContext }) => {
    const workflowState = useContext(workflowStateContext);

    // guard against incomplete GraphQL-cached data
    const safeData = { ...emptyCve, ...data };

    const {
        cve,
        cvss,
        envImpact,
        vectors,
        isFixable,
        summary,
        link,
        lastScanned,
        publishedOn,
        lastModified,
        scoreVersion,
        componentCount,
        imageCount,
        deploymentCount
    } = safeData;

    const linkToNVD = (
        <a
            href={link}
            className="btn-sm btn-base no-underline p-1"
            target="_blank"
            rel="noopener noreferrer nofollow"
        >
            <span className="pr-1">View on NVD Website</span>
            <ExternalLink size={16} />
        </a>
    );

    const cvssScoreBreakdown = [
        {
            key: 'CVSS Score',
            value: cvss && cvss.toFixed(1)
        },
        {
            key: 'Vector',
            value: vectors && vectors.vector
        },
        {
            key: 'Impact Score',
            value: vectors && vectors.impactScore && vectors.impactScore.toFixed(1)
        },
        {
            key: 'Exploitability Score',
            value: vectors && vectors.exploitabilityScore && vectors.exploitabilityScore.toFixed(1)
        }
    ];

    const scanningDetails = [
        {
            key: 'Scanned',
            value: lastScanned ? format(lastScanned, dateTimeFormat) : 'N/A'
        },
        {
            key: 'Published',
            value: publishedOn ? format(publishedOn, dateTimeFormat) : 'N/A'
        },
        {
            key: 'Last modified',
            value: lastModified ? format(lastModified, dateTimeFormat) : 'N/A'
        },
        {
            key: 'Scoring version',
            value: scoreVersion && `CVSS ${scoreVersion}`
        }
    ];

    function getCountData(entityType) {
        switch (entityType) {
            case entityTypes.DEPLOYMENT:
                return deploymentCount;
            case entityTypes.IMAGE:
                return imageCount;
            case entityTypes.COMPONENT:
                return componentCount;
            default:
                return 0;
        }
    }

    const severityStyle = getSeverityChipType(cvss);
    const newEntityContext = { ...entityContext, [entityTypes.CVE]: cve };

    const truncatedSummary = truncate(summary);

    return (
        <div className="flex h-full">
            <div className="flex flex-col flex-grow min-w-0">
                <CollapsibleSection title="CVE summary">
                    <div className="mx-4 grid-dense grid-auto-fit grid grid-gap-6 xxxl:grid-gap-8 grid-columns-1 lg:grid-columns-2 xl:grid-columns-3 mb-4 pdf-page">
                        <Widget
                            header="Details"
                            headerComponents={linkToNVD}
                            className="bg-base-100 min-h-48 lg:s-2"
                        >
                            <div className="flex flex-col w-full bg-counts-widget">
                                <div className="bg-tertiary-200 text-2xl text-base-500 flex flex-col md:flex-row items-start md:items-center justify-between">
                                    <div className="w-full flex-grow p-4">
                                        <span className="text-tertiary-800">{cve}</span>
                                    </div>
                                    <div className="w-full flex border-t border-base-400 md:border-t-0 justify-end items-center">
                                        {// eslint-disable-next-line eqeqeq
                                        envImpact == Number(envImpact) && (
                                            <span className="w-full md:w-auto p-4 border-base-400 text-base-600 border-l whitespace-no-wrap">
                                                <span>
                                                    {' '}
                                                    {`Env. Impact: ${(envImpact * 100).toFixed(
                                                        0
                                                    )}%`}
                                                </span>
                                            </span>
                                        )}
                                        <span className="w-full md:w-auto p-4 border-base-400 border-l">
                                            <LabelChip
                                                text={`CVSS ${cvss && cvss.toFixed(1)}`}
                                                type={severityStyle}
                                            />
                                        </span>
                                        <span className="w-full md:w-auto p-4 border-base-400 border-l">
                                            {isFixable ? (
                                                <LabelChip text="Fixable" type="success" />
                                            ) : (
                                                <LabelChip text="Not fixable" type="base" />
                                            )}
                                        </span>
                                    </div>
                                </div>
                                <div className="p-4 pb-12 leading-loose">{truncatedSummary}</div>
                            </div>
                        </Widget>
                        <Metadata
                            className="bg-base-100 min-h-48"
                            keyValuePairs={cvssScoreBreakdown}
                            title="CVSS Score Breakdown"
                        />
                        <Metadata
                            className="bg-base-100 min-h-48 "
                            keyValuePairs={scanningDetails}
                            title="Scanning Details"
                        />
                    </div>
                </CollapsibleSection>
            </div>
            <RelatedEntitiesSideList
                entityType={entityTypes.CVE}
                workflowState={workflowState}
                getCountData={getCountData}
                entityContext={newEntityContext}
            />
        </div>
    );
};

export default VulnMgmtCveOverview;
