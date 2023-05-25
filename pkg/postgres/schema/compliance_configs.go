// Code generated by pg-bindings generator. DO NOT EDIT.

package schema

import (
	"reflect"

	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/postgres"
	"github.com/stackrox/rox/pkg/postgres/walker"
)

var (
	// CreateTableComplianceConfigsStmt holds the create statement for table `compliance_configs`.
	CreateTableComplianceConfigsStmt = &postgres.CreateStmts{
		GormModel: (*ComplianceConfigs)(nil),
		Children:  []*postgres.CreateStmts{},
	}

	// ComplianceConfigsSchema is the go schema for table `compliance_configs`.
	ComplianceConfigsSchema = func() *walker.Schema {
		schema := GetSchemaForTable("compliance_configs")
		if schema != nil {
			return schema
		}
		schema = walker.Walk(reflect.TypeOf((*storage.ComplianceConfig)(nil)), "compliance_configs")
		RegisterTable(schema, CreateTableComplianceConfigsStmt)
		return schema
	}()
)

const (
	ComplianceConfigsTableName = "compliance_configs"
)

// ComplianceConfigs holds the Gorm model for Postgres table `compliance_configs`.
type ComplianceConfigs struct {
	StandardId string `gorm:"column:standardid;type:varchar;primaryKey"`
	Serialized []byte `gorm:"column:serialized;type:bytea"`
}
