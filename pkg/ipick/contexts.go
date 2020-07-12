package ipick

import (
	"fmt"
	"sort"

	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"

	log "github.com/sirupsen/logrus"

	"github.com/similarweb/kubectl-ipick/command"
)

// ContextsManager describe the context instance
type ContextsManager struct {
	clientapiConfig clientcmdapi.Config
	currentContext  Context
	contexts        []Context
}

// Context describe the .kube config context (see kubectl config get-contexts)
type Context struct {
	Data *clientcmdapi.Context
	Name string
}

// NewContexts manage the cluster contexts
func NewContexts(paths []string) (*ContextsManager, error) {

	clientConfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		&clientcmd.ClientConfigLoadingRules{Precedence: paths},
		&clientcmd.ConfigOverrides{})
	clientapiConfig, err := clientConfig.RawConfig()
	if err != nil {
		return nil, fmt.Errorf("could not init kubernetes client. error: %s", err.Error())
	}

	contexts := []Context{}

	var currentContext Context
	for contextName, context := range clientapiConfig.Contexts {
		log.WithField("name", contextName).Debug("context found")

		ctx := Context{
			Data: context,
			Name: contextName,
		}

		if contextName == clientapiConfig.CurrentContext {
			log.WithField("name", contextName).Debug("current context")
			currentContext = ctx
		}
		contexts = append(contexts, ctx)
	}

	// Order contexts struct by name
	sort.Slice(contexts, func(i, j int) bool {
		return contexts[i].Name < contexts[j].Name
	})

	return &ContextsManager{
		clientapiConfig: clientapiConfig,
		currentContext:  currentContext,
		contexts:        contexts,
	}, nil
}

// GetCurrentContext returns the current context
func (cn *ContextsManager) GetCurrentContext() Context {
	return cn.currentContext
}

// GetContexts returns all the available contexts
func (cn *ContextsManager) GetContexts() []Context {
	return cn.contexts
}

// SetContext will set new context to work with
func (cn *ContextsManager) SetContext(context Context) {
	cn.currentContext = context
}

// SwitchLocalContext will switch current local context
func (cn *ContextsManager) SwitchLocalContext() error {
	return command.Run("kubectl", []string{"config", "set-context", cn.currentContext.Name})
}

// GetContextNames returns list of context names
func (cn *ContextsManager) GetContextNames() []string {

	contextsName := []string{}
	for _, context := range cn.contexts {
		contextsName = append(contextsName, context.Name)
	}

	return contextsName
}
