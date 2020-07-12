package ipick

import (
	"errors"
	"fmt"
	"log"
	"math"
	"math/rand"
	"sort"
	"strings"
	"time"

	"github.com/similarweb/kubectl-ipick/prompt"

	"k8s.io/cli-runtime/pkg/resource"
)

const (
	// defaultNamespace precent the namespace if not set
	defaultNamespace = "default"

	// interactiveResourceText show the current text when the user needs to select resource from the given list
	interactiveResourceText = "Displaying %s "
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
type Ipick struct {
	query  *QueryOptions
	config *Config
	ctx    Context
}

// NewIpick creates new instance of interactive actions
func NewIpick(config *Config) (*Ipick, error) {

	ipick := Ipick{
		config: config,
	}

	contexts, err := NewContexts(config.KubeConfigPaths)

	if err != nil {
		log.Fatal(err)
	}
	if err != nil {
		return nil, err
	}

	if config.SelectCluster {

		selectedCluster, err := prompt.PickSelectionFromData("select context", contexts.GetContextNames())
		if err != nil {
			return nil, err
		}

		// Select the cluster that the user wants to work with
		selectedContext := contexts.GetContexts()[selectedCluster]
		contexts.SetContext(selectedContext)
		err = contexts.SwitchLocalContext()
		if err != nil {
			return nil, err
		}
	}

	// Set the selected/default context.
	ipick.ctx = contexts.GetCurrentContext()
	ipick.query = NewQueryOptions(ipick.ctx.Name)

	return &ipick, nil
}

// SelectResource return selected resource information
func (i *Ipick) SelectResource(resourceType string) (*resource.Info, error) {

	namespace := defaultNamespace

	req := i.query.builder.
		Unstructured().
		ResourceTypeOrNameArgs(true, resourceType).
		Latest().
		Flatten()

	// Set namespace to query builder when --all-namespace not set from the root command.
	if !i.config.AllNamespaces {
		// First take the current user context from .kubeconfig
		if i.ctx.Data.Namespace != "" {
			namespace = i.ctx.Data.Namespace
		}

		// If user set a namespace from root command
		if i.config.Namespace != "" {
			namespace = i.config.Namespace
		}

		// Set namespace to query builder
		req.NamespaceParam(namespace)
	}

	resp := req.Do()
	infos, err := resp.Infos()
	if err != nil {
		return nil, err
	}

	// Order resources info by name field to keep the same order
	sort.Slice(infos, func(i, j int) bool {
		return infos[i].Name < infos[j].Name && infos[i].Namespace < infos[j].Namespace
	})

	var filteredResourcesInfo []*resource.Info
	if i.config.Like != "" {
		filteredResourcesInfo = FilterResources(infos, i.config.Like)
	} else {
		filteredResourcesInfo = FilterResources(infos, i.config.Like)
	}

	// If query builder not found resources
	if len(filteredResourcesInfo) == 0 {
		if i.config.AllNamespaces {
			return nil, errors.New("no resources found")
		}
		return nil, fmt.Errorf("no resources found in %s namespace", namespace)
	}

	var selectedResource int
	// Select random resource from resources responses
	if i.config.Random {
		selectedResource = i.randomInteger(1, len(filteredResourcesInfo))
	} else {

		resourcesNames := []string{}
		namespaceChars := 0
		for _, resourceInfo := range filteredResourcesInfo {
			if len(resourceInfo.Name) > namespaceChars {
				namespaceChars = len(resourceInfo.Name)
			}
		}

		for _, resourceInfo := range filteredResourcesInfo {
			resourcesNames = append(resourcesNames, fmt.Sprintf("%s %s", i.addRightPadding(resourceInfo.Name, namespaceChars+10, " "), resourceInfo.Namespace))
		}
		selectedResource, err = prompt.PickSelectionFromData(fmt.Sprintf(interactiveResourceText, resourceType), resourcesNames)
		if err != nil {
			return nil, err
		}
	}

	resource := filteredResourcesInfo[selectedResource]
	return resource, nil

}

// randomInteger returns random number in range
func (i *Ipick) randomInteger(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min

}

// addLeftPadding add right padding string
func (i *Ipick) addRightPadding(input string, padLength int, padString string) string {
	var output string

	inputLength := len(input)
	padStringLength := len(padString)

	if inputLength >= padLength {
		return input
	}

	repeat := math.Ceil(float64(1) + (float64(padLength-padStringLength))/float64(padStringLength))

	output = input + strings.Repeat(padString, int(repeat))
	output = output[:padLength]

	return output
}
