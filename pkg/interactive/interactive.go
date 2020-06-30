package interactive

import (
	"bytes"
	"errors"
	"fmt"
	"math/rand"
	"sort"

	"github.com/similarweb/kubectl-interactive/prompt"

	"k8s.io/cli-runtime/pkg/resource"
)

const (
	// defaultNamespace precent the namespace if not set
	defaultNamespace = "default"

	// interactiveResourceText show the current text when the user needs to select resource from the given list
	interactiveResourceText = "select %s from the list"

	// interactiveTextValidation show message validaton to approve the working resource
	interactiveTextValidation = "Do you want to perform these action? (Only 'yes' will be accepted)"
)

// Config describe the interactive configuration command
type Config struct {
	SelectCluster   bool
	AllNamespaces   bool
	Random          bool
	Namespace       string
	Like            string
	KubeConfigPaths []string
}

// Interactive describe the interactive instance
type Interactive struct {
	query  *QueryOptions
	config *Config
	ctx    Context
}

// NewInteractive creates new instance of interactive actions
func NewInteractive(config *Config) (*Interactive, error) {

	interactive := Interactive{
		config: config,
	}

	contexts, err := NewContexts(config.KubeConfigPaths)

	if err != nil {
		return nil, err
	}

	if config.SelectCluster {
		clusterBuf := &bytes.Buffer{}
		contexts.PrintClusters(clusterBuf)
		// Print all the available clusters to STDOUT
		fmt.Print(clusterBuf.String())

		// Select the cluster that the user wants to work with
		selectedCluster := prompt.InteractiveNumber("select context", len(contexts.GetContexts())+1)
		selectedContext := contexts.GetContexts()[selectedCluster-1]
		contexts.SetContext(selectedContext)
	}

	// Set the selected/default context.
	interactive.ctx = contexts.GetCurrentContext()
	interactive.query = NewQueryOptions(interactive.ctx.Name)

	return &interactive, nil
}

// SelectResource return selected resource information
func (i *Interactive) SelectResource(resourceType string) (*resource.Info, error) {

	namespace := defaultNamespace
	req := i.query.builder.
		Unstructured().
		ResourceTypeOrNameArgs(true, resourceType).
		Latest().
		Flatten()

	if !i.config.AllNamespaces {

		if i.ctx.Data.Namespace != "" {
			namespace = i.ctx.Data.Namespace
		}

		if i.config.Namespace != "" {
			namespace = i.config.Namespace
		}
		req.NamespaceParam(namespace)
	}

	resp := req.Do()
	infos, err := resp.Infos()

	if err != nil {
		return nil, err
	}

	buf := &bytes.Buffer{}

	sort.Slice(infos, func(i, j int) bool {
		return infos[i].Name < infos[j].Name
	})
	resourcesCount, err := PrintResources(infos, i.config.Like, buf)
	if err != nil {
		return nil, err
	}

	if resourcesCount == 0 {
		if i.config.AllNamespaces {
			return nil, errors.New("no resources found")
		}
		return nil, fmt.Errorf("no resources found in %s namespace", namespace)
	}
	fmt.Print(buf.String())

	var selectedResource int
	// Select random resource from resources responses
	if i.config.Random {
		selectedResource = i.randomInteger(1, resourcesCount)
	} else {
		selectedResource = prompt.InteractiveNumber(fmt.Sprintf(interactiveResourceText, resourceType), resourcesCount)
	}

	// Validate resource when single resource was found
	if resourcesCount == 1 {
		stdInText := prompt.InteractiveText(interactiveTextValidation)
		if stdInText != "yes" {
			return nil, errors.New("request canceld")
		}
	}

	resource := infos[selectedResource-1]
	return resource, nil

}

// randomInteger returns random number in range
func (i *Interactive) randomInteger(min int, max int) int {
	return rand.Intn(max-min) + min
}
