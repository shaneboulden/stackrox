package alertmanager

import (
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/search"
)

// An AlertFilterOption modifies the query builder to filter existing alerts in the DB.
type AlertFilterOption interface {
	apply(*search.QueryBuilder)
	specifiedPolicyID() string
	removedDeploymentID() string
}

type alertFilterOptionImpl struct {
	applyFunc         func(*search.QueryBuilder)
	policyID          string
	removedDeployment string
}

func (a *alertFilterOptionImpl) apply(qb *search.QueryBuilder) {
	a.applyFunc(qb)
}

func (a *alertFilterOptionImpl) specifiedPolicyID() string {
	return a.policyID
}

func (a *alertFilterOptionImpl) removedDeploymentID() string {
	return a.removedDeployment
}

// WithPolicyID returns an AlertFilterOption that filters by policy id.
func WithPolicyID(policyID string) AlertFilterOption {
	return &alertFilterOptionImpl{
		applyFunc: func(qb *search.QueryBuilder) {
			qb.AddExactMatches(search.PolicyID, policyID)
		},
		policyID: policyID,
	}
}

// WithDeploymentID returns an AlertFilterOption that filters by deployment id.
func WithDeploymentID(deploymentID string, isRemove bool) AlertFilterOption {
	opt := &alertFilterOptionImpl{
		applyFunc: func(qb *search.QueryBuilder) {
			qb.AddExactMatches(search.DeploymentID, deploymentID)
		},
	}
	if isRemove {
		opt.removedDeployment = deploymentID
	}
	return opt
}

// WithClusterID returns an AlertFilterOption that filters by cluster.
func WithClusterID(clusterID string) AlertFilterOption {
	return &alertFilterOptionImpl{
		applyFunc: func(qb *search.QueryBuilder) {
			qb.AddExactMatches(search.ClusterID, clusterID)
		},
	}
}

// WithoutResourceType returns an AlertFilterOption that filters _out_ the specified resource type.
func WithoutResourceType(resourceType storage.ListAlert_ResourceType) AlertFilterOption {
	return &alertFilterOptionImpl{
		applyFunc: func(qb *search.QueryBuilder) {
			qb.AddStrings(search.ResourceType, search.NegateQueryString(resourceType.String()))
		},
	}
}

// WithLifecycleStage returns an AlertFilterOptions that filters by lifecycle stage.
func WithLifecycleStage(lifecycleStage storage.LifecycleStage) AlertFilterOption {
	return &alertFilterOptionImpl{
		applyFunc: func(qb *search.QueryBuilder) {
			qb.AddExactMatches(search.LifecycleStage, lifecycleStage.String())
		},
	}
}
