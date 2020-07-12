package ipick

import (
	"testing"

	"k8s.io/cli-runtime/pkg/resource"
)

func TestFilterResources(t *testing.T) {

	infos := []*resource.Info{
		{Name: "test-1"},
		{Name: "test-2"},
		{Name: "test-3"},
		{Name: "foo-1"},
		{Name: "foo-2"},
	}

	filters := FilterResources(infos, "foo")

	if len(filters) != 2 {
		t.Errorf("incorrect filter results, got: %d, want: %d.", len(filters), 2)
	}

}
