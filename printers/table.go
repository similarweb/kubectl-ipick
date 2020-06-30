package printers

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"k8s.io/cli-runtime/pkg/printers"
)

// PrintTableConfig defines the headers value
type PrintTableConfig struct {
	Key    string
	Header string
}

// Table will render given data to the given io.writer
func Table(tableConfig []PrintTableConfig, data []byte, out io.Writer) {

	w := printers.GetNewTabWriter(out)

	var dataList []map[string]interface{}
	json.Unmarshal(data, &dataList)

	columnNames := []string{"ID"}
	for _, cell := range tableConfig {
		columnNames = append(columnNames, strings.ToUpper(cell.Header))
	}

	fmt.Fprintf(w, "%s\n", fmt.Sprintf(strings.Join(columnNames, "\t")))

	for i, details := range dataList {
		rowData := []string{fmt.Sprintf("%v", i+1)}
		for _, row := range tableConfig {
			rowData = append(rowData, fmt.Sprintf("%v", details[row.Key]))
		}
		fmt.Fprintf(w, "%s\n", strings.Join(rowData, "\t"))
	}

	w.Flush()

}
