package printers

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"

	log "github.com/sirupsen/logrus"
	"k8s.io/cli-runtime/pkg/printers"
)

// PrintTableConfig defines the headers value
type PrintTableConfig struct {
	Key    string
	Header string
}

// Table will render given data to the given io.writer
func Table(tableConfig []PrintTableConfig, data []byte, out io.Writer) error {

	w := printers.GetNewTabWriter(out)

	var dataList []map[string]interface{}
	err := json.Unmarshal(data, &dataList)
	if err != nil {
		log.Error("could not unmarshal table data")
		return err
	}

	columnNames := []string{"ID"}
	for _, cell := range tableConfig {
		columnNames = append(columnNames, strings.ToUpper(cell.Header))
	}

	fmt.Fprintf(w, "%s\n", strings.Join(columnNames, "\t"))

	for i, details := range dataList {
		rowData := []string{fmt.Sprintf("%v", i+1)}
		for _, row := range tableConfig {
			rowData = append(rowData, fmt.Sprintf("%v", details[row.Key]))
		}
		fmt.Fprintf(w, "%s\n", strings.Join(rowData, "\t"))
	}

	w.Flush()
	return nil

}
