// Code generated by pg-bindings generator. DO NOT EDIT.

package postgres

import (
	"context"
	"strings"
	"time"

	"github.com/hashicorp/go-multierror"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
	"github.com/stackrox/rox/central/metrics"
	"github.com/stackrox/rox/central/role/resources"
	v1 "github.com/stackrox/rox/generated/api/v1"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/logging"
	ops "github.com/stackrox/rox/pkg/metrics"
	"github.com/stackrox/rox/pkg/postgres"
	"github.com/stackrox/rox/pkg/postgres/pgutils"
	pkgSchema "github.com/stackrox/rox/pkg/postgres/schema"
	"github.com/stackrox/rox/pkg/sac"
	"github.com/stackrox/rox/pkg/search"
	pgSearch "github.com/stackrox/rox/pkg/search/postgres"
	"github.com/stackrox/rox/pkg/sync"
	"gorm.io/gorm"
)

const (
	baseTable = "role_bindings"
	storeName = "K8SRoleBinding"

	batchAfter = 100

	// using copyFrom, we may not even want to batch.  It would probably be simpler
	// to deal with failures if we just sent it all.  Something to think about as we
	// proceed and move into more e2e and larger performance testing
	batchSize       = 10000
	deleteBatchSize = 5000
)

var (
	log            = logging.LoggerForModule()
	schema         = pkgSchema.RoleBindingsSchema
	targetResource = resources.K8sRoleBinding
)

// Store is the interface to interact with the storage for storage.K8SRoleBinding
type Store interface {
	Upsert(ctx context.Context, obj *storage.K8SRoleBinding) error
	UpsertMany(ctx context.Context, objs []*storage.K8SRoleBinding) error
	Delete(ctx context.Context, id string) error
	DeleteByQuery(ctx context.Context, q *v1.Query) error
	DeleteMany(ctx context.Context, identifiers []string) error

	Count(ctx context.Context) (int, error)
	Exists(ctx context.Context, id string) (bool, error)

	Get(ctx context.Context, id string) (*storage.K8SRoleBinding, bool, error)
	GetByQuery(ctx context.Context, query *v1.Query) ([]*storage.K8SRoleBinding, error)
	GetMany(ctx context.Context, identifiers []string) ([]*storage.K8SRoleBinding, []int, error)
	GetIDs(ctx context.Context) ([]string, error)

	Walk(ctx context.Context, fn func(obj *storage.K8SRoleBinding) error) error
}

type storeImpl struct {
	*pgSearch.GenericStore[storage.K8SRoleBinding, *storage.K8SRoleBinding]
	db    postgres.DB
	mutex sync.RWMutex
}

// New returns a new Store instance using the provided sql instance.
func New(db postgres.DB) Store {
	return &storeImpl{
		db: db,
		GenericStore: pgSearch.NewGenericStore[storage.K8SRoleBinding, *storage.K8SRoleBinding](
			db,
			schema,
			pkGetter,
			metricsSetAcquireDBConnDuration,
			metricsSetPostgresOperationDurationTime,
			targetResource,
		),
	}
}

// region Helper functions

func pkGetter(obj *storage.K8SRoleBinding) string {
	return obj.GetId()
}

func metricsSetPostgresOperationDurationTime(start time.Time, op ops.Op) {
	metrics.SetPostgresOperationDurationTime(start, op, storeName)
}

func metricsSetAcquireDBConnDuration(start time.Time, op ops.Op) {
	metrics.SetAcquireDBConnDuration(start, op, storeName)
}

func insertIntoRoleBindings(ctx context.Context, batch *pgx.Batch, obj *storage.K8SRoleBinding) error {

	serialized, marshalErr := obj.Marshal()
	if marshalErr != nil {
		return marshalErr
	}

	values := []interface{}{
		// parent primary keys start
		pgutils.NilOrUUID(obj.GetId()),
		obj.GetName(),
		obj.GetNamespace(),
		pgutils.NilOrUUID(obj.GetClusterId()),
		obj.GetClusterName(),
		obj.GetClusterRole(),
		obj.GetLabels(),
		obj.GetAnnotations(),
		pgutils.NilOrUUID(obj.GetRoleId()),
		serialized,
	}

	finalStr := "INSERT INTO role_bindings (Id, Name, Namespace, ClusterId, ClusterName, ClusterRole, Labels, Annotations, RoleId, serialized) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) ON CONFLICT(Id) DO UPDATE SET Id = EXCLUDED.Id, Name = EXCLUDED.Name, Namespace = EXCLUDED.Namespace, ClusterId = EXCLUDED.ClusterId, ClusterName = EXCLUDED.ClusterName, ClusterRole = EXCLUDED.ClusterRole, Labels = EXCLUDED.Labels, Annotations = EXCLUDED.Annotations, RoleId = EXCLUDED.RoleId, serialized = EXCLUDED.serialized"
	batch.Queue(finalStr, values...)

	var query string

	for childIndex, child := range obj.GetSubjects() {
		if err := insertIntoRoleBindingsSubjects(ctx, batch, child, obj.GetId(), childIndex); err != nil {
			return err
		}
	}

	query = "delete from role_bindings_subjects where role_bindings_Id = $1 AND idx >= $2"
	batch.Queue(query, pgutils.NilOrUUID(obj.GetId()), len(obj.GetSubjects()))
	return nil
}

func insertIntoRoleBindingsSubjects(_ context.Context, batch *pgx.Batch, obj *storage.Subject, roleBindingID string, idx int) error {

	values := []interface{}{
		// parent primary keys start
		pgutils.NilOrUUID(roleBindingID),
		idx,
		obj.GetKind(),
		obj.GetName(),
	}

	finalStr := "INSERT INTO role_bindings_subjects (role_bindings_Id, idx, Kind, Name) VALUES($1, $2, $3, $4) ON CONFLICT(role_bindings_Id, idx) DO UPDATE SET role_bindings_Id = EXCLUDED.role_bindings_Id, idx = EXCLUDED.idx, Kind = EXCLUDED.Kind, Name = EXCLUDED.Name"
	batch.Queue(finalStr, values...)

	return nil
}

func (s *storeImpl) copyFromRoleBindings(ctx context.Context, tx *postgres.Tx, objs ...*storage.K8SRoleBinding) error {

	inputRows := [][]interface{}{}

	var err error

	// This is a copy so first we must delete the rows and re-add them
	// Which is essentially the desired behaviour of an upsert.
	var deletes []string

	copyCols := []string{

		"id",

		"name",

		"namespace",

		"clusterid",

		"clustername",

		"clusterrole",

		"labels",

		"annotations",

		"roleid",

		"serialized",
	}

	for idx, obj := range objs {
		// Todo: ROX-9499 Figure out how to more cleanly template around this issue.
		log.Debugf("This is here for now because there is an issue with pods_TerminatedInstances where the obj "+
			"in the loop is not used as it only consists of the parent ID and the index.  Putting this here as a stop gap "+
			"to simply use the object.  %s", obj)

		serialized, marshalErr := obj.Marshal()
		if marshalErr != nil {
			return marshalErr
		}

		inputRows = append(inputRows, []interface{}{

			pgutils.NilOrUUID(obj.GetId()),

			obj.GetName(),

			obj.GetNamespace(),

			pgutils.NilOrUUID(obj.GetClusterId()),

			obj.GetClusterName(),

			obj.GetClusterRole(),

			obj.GetLabels(),

			obj.GetAnnotations(),

			pgutils.NilOrUUID(obj.GetRoleId()),

			serialized,
		})

		// Add the ID to be deleted.
		deletes = append(deletes, obj.GetId())

		// if we hit our batch size we need to push the data
		if (idx+1)%batchSize == 0 || idx == len(objs)-1 {
			// copy does not upsert so have to delete first.  parent deletion cascades so only need to
			// delete for the top level parent

			if err := s.DeleteMany(ctx, deletes); err != nil {
				return err
			}
			// clear the inserts and vals for the next batch
			deletes = nil

			_, err = tx.CopyFrom(ctx, pgx.Identifier{"role_bindings"}, copyCols, pgx.CopyFromRows(inputRows))

			if err != nil {
				return err
			}

			// clear the input rows for the next batch
			inputRows = inputRows[:0]
		}
	}

	for idx, obj := range objs {
		_ = idx // idx may or may not be used depending on how nested we are, so avoid compile-time errors.

		if err = s.copyFromRoleBindingsSubjects(ctx, tx, obj.GetId(), obj.GetSubjects()...); err != nil {
			return err
		}
	}

	return err
}

func (s *storeImpl) copyFromRoleBindingsSubjects(ctx context.Context, tx *postgres.Tx, roleBindingID string, objs ...*storage.Subject) error {

	inputRows := [][]interface{}{}

	var err error

	copyCols := []string{

		"role_bindings_id",

		"idx",

		"kind",

		"name",
	}

	for idx, obj := range objs {
		// Todo: ROX-9499 Figure out how to more cleanly template around this issue.
		log.Debugf("This is here for now because there is an issue with pods_TerminatedInstances where the obj "+
			"in the loop is not used as it only consists of the parent ID and the index.  Putting this here as a stop gap "+
			"to simply use the object.  %s", obj)

		inputRows = append(inputRows, []interface{}{

			pgutils.NilOrUUID(roleBindingID),

			idx,

			obj.GetKind(),

			obj.GetName(),
		})

		// if we hit our batch size we need to push the data
		if (idx+1)%batchSize == 0 || idx == len(objs)-1 {
			// copy does not upsert so have to delete first.  parent deletion cascades so only need to
			// delete for the top level parent

			_, err = tx.CopyFrom(ctx, pgx.Identifier{"role_bindings_subjects"}, copyCols, pgx.CopyFromRows(inputRows))

			if err != nil {
				return err
			}

			// clear the input rows for the next batch
			inputRows = inputRows[:0]
		}
	}

	return err
}

func (s *storeImpl) copyFrom(ctx context.Context, objs ...*storage.K8SRoleBinding) error {
	conn, err := s.AcquireConn(ctx, ops.Get)
	if err != nil {
		return err
	}
	defer conn.Release()

	tx, err := conn.Begin(ctx)
	if err != nil {
		return err
	}

	if err := s.copyFromRoleBindings(ctx, tx, objs...); err != nil {
		if err := tx.Rollback(ctx); err != nil {
			return err
		}
		return err
	}
	if err := tx.Commit(ctx); err != nil {
		return err
	}
	return nil
}

func (s *storeImpl) upsert(ctx context.Context, objs ...*storage.K8SRoleBinding) error {
	conn, err := s.AcquireConn(ctx, ops.Get)
	if err != nil {
		return err
	}
	defer conn.Release()

	for _, obj := range objs {
		batch := &pgx.Batch{}
		if err := insertIntoRoleBindings(ctx, batch, obj); err != nil {
			return err
		}
		batchResults := conn.SendBatch(ctx, batch)
		var result *multierror.Error
		for i := 0; i < batch.Len(); i++ {
			_, err := batchResults.Exec()
			result = multierror.Append(result, err)
		}
		if err := batchResults.Close(); err != nil {
			return err
		}
		if err := result.ErrorOrNil(); err != nil {
			return err
		}
	}
	return nil
}

// endregion Helper functions

//// Interface functions

// Upsert saves the current state of an object in storage.
func (s *storeImpl) Upsert(ctx context.Context, obj *storage.K8SRoleBinding) error {
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.Upsert, "K8SRoleBinding")

	scopeChecker := sac.GlobalAccessScopeChecker(ctx).AccessMode(storage.Access_READ_WRITE_ACCESS).Resource(targetResource).
		ClusterID(obj.GetClusterId()).Namespace(obj.GetNamespace())
	if !scopeChecker.IsAllowed() {
		return sac.ErrResourceAccessDenied
	}

	return pgutils.Retry(func() error {
		return s.upsert(ctx, obj)
	})
}

// UpsertMany saves the state of multiple objects in the storage.
func (s *storeImpl) UpsertMany(ctx context.Context, objs []*storage.K8SRoleBinding) error {
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.UpdateMany, "K8SRoleBinding")

	scopeChecker := sac.GlobalAccessScopeChecker(ctx).AccessMode(storage.Access_READ_WRITE_ACCESS).Resource(targetResource)
	if !scopeChecker.IsAllowed() {
		var deniedIDs []string
		for _, obj := range objs {
			subScopeChecker := scopeChecker.ClusterID(obj.GetClusterId()).Namespace(obj.GetNamespace())
			if !subScopeChecker.IsAllowed() {
				deniedIDs = append(deniedIDs, obj.GetId())
			}
		}
		if len(deniedIDs) != 0 {
			return errors.Wrapf(sac.ErrResourceAccessDenied, "modifying k8SRoleBindings with IDs [%s] was denied", strings.Join(deniedIDs, ", "))
		}
	}

	return pgutils.Retry(func() error {
		// Lock since copyFrom requires a delete first before being executed.  If multiple processes are updating
		// same subset of rows, both deletes could occur before the copyFrom resulting in unique constraint
		// violations
		if len(objs) < batchAfter {
			s.mutex.RLock()
			defer s.mutex.RUnlock()

			return s.upsert(ctx, objs...)
		}
		s.mutex.Lock()
		defer s.mutex.Unlock()

		return s.copyFrom(ctx, objs...)
	})
}

// DeleteMany removes the objects associated to the specified IDs from the store.
func (s *storeImpl) DeleteMany(ctx context.Context, identifiers []string) error {
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.RemoveMany, "K8SRoleBinding")

	var sacQueryFilter *v1.Query

	sacQueryFilter, err := pgSearch.GetReadWriteSACQuery(ctx, targetResource)
	if err != nil {
		return err
	}

	// Batch the deletes
	localBatchSize := deleteBatchSize
	numRecordsToDelete := len(identifiers)
	for {
		if len(identifiers) == 0 {
			break
		}

		if len(identifiers) < localBatchSize {
			localBatchSize = len(identifiers)
		}

		identifierBatch := identifiers[:localBatchSize]
		q := search.ConjunctionQuery(
			sacQueryFilter,
			search.NewQueryBuilder().AddDocIDs(identifierBatch...).ProtoQuery(),
		)

		if err := pgSearch.RunDeleteRequestForSchema(ctx, schema, q, s.db); err != nil {
			return errors.Wrapf(err, "unable to delete the records.  Successfully deleted %d out of %d", numRecordsToDelete-len(identifiers), numRecordsToDelete)
		}

		// Move the slice forward to start the next batch
		identifiers = identifiers[localBatchSize:]
	}

	return nil
}

//// Interface functions - END

//// Used for testing

// CreateTableAndNewStore returns a new Store instance for testing.
func CreateTableAndNewStore(ctx context.Context, db postgres.DB, gormDB *gorm.DB) Store {
	pkgSchema.ApplySchemaForTable(ctx, gormDB, baseTable)
	return New(db)
}

// Destroy drops the tables associated with the target object type.
func Destroy(ctx context.Context, db postgres.DB) {
	dropTableRoleBindings(ctx, db)
}

func dropTableRoleBindings(ctx context.Context, db postgres.DB) {
	_, _ = db.Exec(ctx, "DROP TABLE IF EXISTS role_bindings CASCADE")
	dropTableRoleBindingsSubjects(ctx, db)

}

func dropTableRoleBindingsSubjects(ctx context.Context, db postgres.DB) {
	_, _ = db.Exec(ctx, "DROP TABLE IF EXISTS role_bindings_subjects CASCADE")

}

//// Used for testing - END
