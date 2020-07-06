package interactive

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/similarweb/kubectl-interactive/printers"

	"k8s.io/cli-runtime/pkg/resource"
)

// tableData describe the present table
type tableData struct {
	Name      string
	Namespace string
}

// tableHeader set a table headers
var tableHeader = []printers.PrintTableConfig{
	{Header: "Name", Key: "Name"},
	{Header: "Namespace", Key: "Namespace"},
}

// PrintResources will add all the resource data the given buffer
func PrintResources(resourceInfo []*resource.Info, like string, buf *bytes.Buffer) ([]*resource.Info, error) {

	resourcesResults := []*resource.Info{}
	data := []tableData{}
	for _, info := range resourceInfo {
		if like != "" && !strings.Contains(strings.ToLower(info.Name), strings.ToLower(like)) {
			continue
		}
		resourcesResults = append(resourcesResults, info)
		data = append(data, tableData{
			Name:      info.Name,
			Namespace: info.Namespace,
		})
	}

	resourceBuffer, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("could not marshal contexts. error: %s", err.Error())
	}
	err = printers.Table(tableHeader, resourceBuffer, buf)
	if err != nil {
		return nil, err
	}
	return resourcesResults, nil

}
