package interactive

import (
	"bytes"
	"testing"

	"k8s.io/cli-runtime/pkg/resource"
)

func TestPopulateResources(t *testing.T) {

	infos := []*resource.Info{
		{Namespace: "A", Name: "resource Name"},
		{Namespace: "A", Name: "resource Name-1"},
		{Namespace: "B", Name: "resource Name-2"},
	}

	t.Run("print", func(t *testing.T) {
		exceptTable := `ID    NAME              NAMESPACE
1     resource Name     A
2     resource Name-1   A
3     resource Name-2   B
`
		buf := &bytes.Buffer{}
		resourcesInfo, err := PopulateResources(infos, "", buf)

		if err != nil {
			t.Fatalf("unexpected populateClusters error, got %s expected %s", err.Error(), "nil")
		}

		if buf.String() != exceptTable {
			t.Fatalf("unexpected table content, got %s expected %s", buf.String(), exceptTable)
		}

		if len(resourcesInfo) != len(infos) {
			t.Fatalf("unexpected resource count, got %d expected %d", len(resourcesInfo), len(infos))
		}

	})
	t.Run("print filters", func(t *testing.T) {
		exceptTable := `ID    NAME              NAMESPACE
1     resource Name-1   A
2     resource Name-2   B
`
		buf := &bytes.Buffer{}
		resourcesInfo, err := PopulateResources(infos, "Name-", buf)

		if err != nil {
			t.Fatalf("unexpected PopulateClusters error, got %s expected %s", err.Error(), "nil")
		}

		if buf.String() != exceptTable {
			t.Fatalf("unexpected table content, got %s expected %s", buf.String(), exceptTable)
		}

		if len(resourcesInfo) != 2 {
			t.Fatalf("unexpected resource count, got %d expected %d", len(resourcesInfo), 2)
		}

	})
}
