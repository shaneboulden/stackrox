// Code generated by pg-bindings generator. DO NOT EDIT.

package postgres

import (
	"context"
	"time"

	"github.com/gogo/protobuf/proto"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/stackrox/rox/central/metrics"
	pkgSchema "github.com/stackrox/rox/central/postgres/schema"
	v1 "github.com/stackrox/rox/generated/api/v1"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/logging"
	ops "github.com/stackrox/rox/pkg/metrics"
	"github.com/stackrox/rox/pkg/postgres/pgutils"
	"github.com/stackrox/rox/pkg/search/postgres"
)

const (
	baseTable = "image_component_cve_relations"

	existsStmt = "SELECT EXISTS(SELECT 1 FROM image_component_cve_relations WHERE Id = $1 AND ImageComponentId = $2 AND ImageCveId = $3)"
	getStmt    = "SELECT serialized FROM image_component_cve_relations WHERE Id = $1 AND ImageComponentId = $2 AND ImageCveId = $3"
	deleteStmt = "DELETE FROM image_component_cve_relations WHERE Id = $1 AND ImageComponentId = $2 AND ImageCveId = $3"

	walkStmt    = "SELECT serialized FROM image_component_cve_relations"
	getManyStmt = "SELECT serialized FROM image_component_cve_relations WHERE Id = ANY($1::text[])"

	batchAfter = 100

	// using copyFrom, we may not even want to batch.  It would probably be simpler
	// to deal with failures if we just sent it all.  Something to think about as we
	// proceed and move into more e2e and larger performance testing
	batchSize = 10000
)

var (
	log    = logging.LoggerForModule()
	schema = pkgSchema.ImageComponentCveRelationsSchema
)

type Store interface {
	Count(ctx context.Context) (int, error)
	Exists(ctx context.Context, id string, imageComponentId string, imageCveId string) (bool, error)
	Get(ctx context.Context, id string, imageComponentId string, imageCveId string) (*storage.ComponentCVEEdge, bool, error)
	GetIDs(ctx context.Context) ([]string, error)
	GetMany(ctx context.Context, ids []string) ([]*storage.ComponentCVEEdge, []int, error)

	Walk(ctx context.Context, fn func(obj *storage.ComponentCVEEdge) error) error

	AckKeysIndexed(ctx context.Context, keys ...string) error
	GetKeysToIndex(ctx context.Context) ([]string, error)
}

type storeImpl struct {
	db *pgxpool.Pool
}

// New returns a new Store instance using the provided sql instance.
func New(ctx context.Context, db *pgxpool.Pool) Store {
	pgutils.CreateTable(ctx, db, pkgSchema.CreateTableImageComponentsStmt)
	pgutils.CreateTable(ctx, db, pkgSchema.CreateTableImageCvesStmt)
	pgutils.CreateTable(ctx, db, pkgSchema.CreateTableImageComponentCveRelationsStmt)

	return &storeImpl{
		db: db,
	}
}

// Count returns the number of objects in the store
func (s *storeImpl) Count(ctx context.Context) (int, error) {
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.Count, "ComponentCVEEdge")

	var sacQueryFilter *v1.Query

	return postgres.RunCountRequestForSchema(schema, sacQueryFilter, s.db)
}

// Exists returns if the id exists in the store
func (s *storeImpl) Exists(ctx context.Context, id string, imageComponentId string, imageCveId string) (bool, error) {
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.Exists, "ComponentCVEEdge")

	row := s.db.QueryRow(ctx, existsStmt, id, imageComponentId, imageCveId)
	var exists bool
	if err := row.Scan(&exists); err != nil {
		return false, pgutils.ErrNilIfNoRows(err)
	}
	return exists, nil
}

// Get returns the object, if it exists from the store
func (s *storeImpl) Get(ctx context.Context, id string, imageComponentId string, imageCveId string) (*storage.ComponentCVEEdge, bool, error) {
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.Get, "ComponentCVEEdge")

	conn, release, err := s.acquireConn(ctx, ops.Get, "ComponentCVEEdge")
	if err != nil {
		return nil, false, err
	}
	defer release()

	row := conn.QueryRow(ctx, getStmt, id, imageComponentId, imageCveId)
	var data []byte
	if err := row.Scan(&data); err != nil {
		return nil, false, pgutils.ErrNilIfNoRows(err)
	}

	var msg storage.ComponentCVEEdge
	if err := proto.Unmarshal(data, &msg); err != nil {
		return nil, false, err
	}
	return &msg, true, nil
}

func (s *storeImpl) acquireConn(ctx context.Context, op ops.Op, typ string) (*pgxpool.Conn, func(), error) {
	defer metrics.SetAcquireDBConnDuration(time.Now(), op, typ)
	conn, err := s.db.Acquire(ctx)
	if err != nil {
		return nil, nil, err
	}
	return conn, conn.Release, nil
}

// GetIDs returns all the IDs for the store
func (s *storeImpl) GetIDs(ctx context.Context) ([]string, error) {
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.GetAll, "storage.ComponentCVEEdgeIDs")
	var sacQueryFilter *v1.Query

	result, err := postgres.RunSearchRequestForSchema(schema, sacQueryFilter, s.db)
	if err != nil {
		return nil, err
	}

	ids := make([]string, 0, len(result))
	for _, entry := range result {
		ids = append(ids, entry.ID)
	}

	return ids, nil
}

// GetMany returns the objects specified by the IDs or the index in the missing indices slice
func (s *storeImpl) GetMany(ctx context.Context, ids []string) ([]*storage.ComponentCVEEdge, []int, error) {
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.GetMany, "ComponentCVEEdge")

	conn, release, err := s.acquireConn(ctx, ops.GetMany, "ComponentCVEEdge")
	if err != nil {
		return nil, nil, err
	}
	defer release()

	rows, err := conn.Query(ctx, getManyStmt, ids)
	if err != nil {
		if err == pgx.ErrNoRows {
			missingIndices := make([]int, 0, len(ids))
			for i := range ids {
				missingIndices = append(missingIndices, i)
			}
			return nil, missingIndices, nil
		}
		return nil, nil, err
	}
	defer rows.Close()
	resultsByID := make(map[string]*storage.ComponentCVEEdge)
	for rows.Next() {
		var data []byte
		if err := rows.Scan(&data); err != nil {
			return nil, nil, err
		}
		msg := &storage.ComponentCVEEdge{}
		if err := proto.Unmarshal(data, msg); err != nil {
			return nil, nil, err
		}
		resultsByID[msg.GetId()] = msg
	}
	missingIndices := make([]int, 0, len(ids)-len(resultsByID))
	// It is important that the elems are populated in the same order as the input ids
	// slice, since some calling code relies on that to maintain order.
	elems := make([]*storage.ComponentCVEEdge, 0, len(resultsByID))
	for i, id := range ids {
		if result, ok := resultsByID[id]; !ok {
			missingIndices = append(missingIndices, i)
		} else {
			elems = append(elems, result)
		}
	}
	return elems, missingIndices, nil
}

// Walk iterates over all of the objects in the store and applies the closure
func (s *storeImpl) Walk(ctx context.Context, fn func(obj *storage.ComponentCVEEdge) error) error {
	rows, err := s.db.Query(ctx, walkStmt)
	if err != nil {
		return pgutils.ErrNilIfNoRows(err)
	}
	defer rows.Close()
	for rows.Next() {
		var data []byte
		if err := rows.Scan(&data); err != nil {
			return err
		}
		var msg storage.ComponentCVEEdge
		if err := proto.Unmarshal(data, &msg); err != nil {
			return err
		}
		if err := fn(&msg); err != nil {
			return err
		}
	}
	return nil
}

//// Used for testing

func dropTableImageComponentCveRelations(ctx context.Context, db *pgxpool.Pool) {
	_, _ = db.Exec(ctx, "DROP TABLE IF EXISTS image_component_cve_relations CASCADE")

}

func Destroy(ctx context.Context, db *pgxpool.Pool) {
	dropTableImageComponentCveRelations(ctx, db)
}

//// Stubs for satisfying legacy interfaces

// AckKeysIndexed acknowledges the passed keys were indexed
func (s *storeImpl) AckKeysIndexed(ctx context.Context, keys ...string) error {
	return nil
}

// GetKeysToIndex returns the keys that need to be indexed
func (s *storeImpl) GetKeysToIndex(ctx context.Context) ([]string, error) {
	return nil, nil
}
