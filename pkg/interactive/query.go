package interactive

import (
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/cli-runtime/pkg/resource"
)

// QueryOptions describe the restAPI query builder
type QueryOptions struct {
	builder     *resource.Builder
	configFlags *genericclioptions.ConfigFlags
}

// NewQueryOptions create a new restAPI resource query
func NewQueryOptions(context string) *QueryOptions {
	queryOptions := &QueryOptions{
		configFlags: genericclioptions.NewConfigFlags(false),
	}

	queryOptions.configFlags.Context = &context
	queryOptions.builder = resource.NewBuilder(queryOptions.configFlags)
	return queryOptions
}
