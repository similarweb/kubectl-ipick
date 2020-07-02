package printers

import (
	"bytes"
	"encoding/json"
	"testing"
)

type tableContent struct {
	A string
	B string
	C string
}

func TestTable(t *testing.T) {

	tableHeaders := []PrintTableConfig{
		{Key: "A", Header: "A"},
		{Key: "B", Header: "B"},
	}

	tableContent := []tableContent{
		{A: "first", B: "first", C: "first"},
		{A: "second", B: "second", C: "second"},
	}

	exceptTable := `ID    A        B
1     first    first
2     second   second
`
	buf := &bytes.Buffer{}
	resourceBuffer, _ := json.Marshal(tableContent)

	err := Table(tableHeaders, resourceBuffer, buf)

	if err != nil {
		t.Fatalf("unexpected Table error, got %s expected %s", err.Error(), "nil")
	}
	if buf.String() != exceptTable {
		t.Fatalf("unexpected table content, got %s expected %s", buf.String(), exceptTable)
	}

}
