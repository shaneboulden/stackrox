// Code generated by pg-bindings generator. DO NOT EDIT.

package schema

import (
	"fmt"
	"reflect"

	v1 "github.com/stackrox/rox/generated/api/v1"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/postgres"
	"github.com/stackrox/rox/pkg/postgres/walker"
	"github.com/stackrox/rox/pkg/search"
	"github.com/stackrox/rox/pkg/search/postgres/mapping"
)

var (
	// CreateTableListeningEndpointsStmt holds the create statement for table `listening_endpoints`.
	CreateTableListeningEndpointsStmt = &postgres.CreateStmts{
		GormModel: (*ListeningEndpoints)(nil),
		Children:  []*postgres.CreateStmts{},
	}

	// ListeningEndpointsSchema is the go schema for table `listening_endpoints`.
	ListeningEndpointsSchema = func() *walker.Schema {
		schema := GetSchemaForTable("listening_endpoints")
		if schema != nil {
			return schema
		}
		schema = walker.Walk(reflect.TypeOf((*storage.ProcessListeningOnPortStorage)(nil)), "listening_endpoints")
		referencedSchemas := map[string]*walker.Schema{
			"storage.ProcessIndicator": ProcessIndicatorsSchema,
			"storage.Deployment":       DeploymentsSchema,
		}

		schema.ResolveReferences(func(messageTypeName string) *walker.Schema {
			return referencedSchemas[fmt.Sprintf("storage.%s", messageTypeName)]
		})
		schema.SetOptionsMap(search.Walk(v1.SearchCategory_PROCESS_LISTENING_ON_PORT, "processlisteningonportstorage", (*storage.ProcessListeningOnPortStorage)(nil)))
		RegisterTable(schema, CreateTableListeningEndpointsStmt)
		mapping.RegisterCategoryToTable(v1.SearchCategory_PROCESS_LISTENING_ON_PORT, schema)
		return schema
	}()
)

const (
	ListeningEndpointsTableName = "listening_endpoints"
)

// ListeningEndpoints holds the Gorm model for Postgres table `listening_endpoints`.
type ListeningEndpoints struct {
	Id                 string             `gorm:"column:id;type:uuid;primaryKey"`
	Port               uint32             `gorm:"column:port;type:bigint"`
	Protocol           storage.L4Protocol `gorm:"column:protocol;type:integer"`
	ProcessIndicatorId string             `gorm:"column:processindicatorid;type:uuid;index:listeningendpoints_processindicatorid,type:btree"`
	Closed             bool               `gorm:"column:closed;type:bool;index:listeningendpoints_closed,type:btree"`
	DeploymentId       string             `gorm:"column:deploymentid;type:uuid;index:listeningendpoints_deploymentid,type:btree"`
	Serialized         []byte             `gorm:"column:serialized;type:bytea"`
}