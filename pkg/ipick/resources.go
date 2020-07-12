package ipick

import (
	"strings"

	"k8s.io/cli-runtime/pkg/resource"
)

// FilterResources returns resources given like string
func FilterResources(resourceInfo []*resource.Info, like string) []*resource.Info {

	resourcesResults := []*resource.Info{}
	for _, info := range resourceInfo {
		if like != "" && !strings.Contains(strings.ToLower(info.Name), strings.ToLower(like)) {
			continue
		}
		resourcesResults = append(resourcesResults, info)
	}

	return resourcesResults

}
