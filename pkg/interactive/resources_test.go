package interactive

import (
	"bytes"
	"testing"

	"k8s.io/cli-runtime/pkg/resource"
)

func TestPrintResources(t *testing.T) {

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
		PrintResources(infos, "", buf)

		if buf.String() != exceptTable {
			t.Fatalf("unexpected table content, got %s expected %s", buf.String(), exceptTable)
		}

	})
	t.Run("print filters", func(t *testing.T) {
		exceptTable := `ID    NAME              NAMESPACE
1     resource Name-1   A
2     resource Name-2   B
`
		buf := &bytes.Buffer{}
		PrintResources(infos, "Name-", buf)

		if buf.String() != exceptTable {
			t.Fatalf("unexpected table content, got %s expected %s", buf.String(), exceptTable)
		}

	})
}
