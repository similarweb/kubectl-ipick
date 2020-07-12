package cmd

import (
	"fmt"
	"os"
	"os/user"
	"strings"

	"github.com/similarweb/kubectl-interactive/command"
	"github.com/similarweb/kubectl-interactive/pkg/interactive"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

const (
	// commandName describe the plugin command name
	commandName = "interactive"

	// defaultKubeConfigPath set the default kubeconfig path
	defaultKubeConfigPath = ".kube/config"
)

var (
	// If present, user need to select one of the cluster from the kubeconfig configuration
	selectCluster bool

	// If present, search the resource in all namespaces
	allNamespaces bool

	// If presend, select randomly resource
	random bool

	// If present search the resource from the given namespace
	namespace string

	// If present, prompt only resources that include the given value
	like string

	// Append kubectl flags
	flags string

	// Set application log level
	lv string

	kubeConfigPath string

	// ignoreNamespaceSet will not set -n <namespace> while the action is one of the list
	ignoreNamespaceSet = []string{"node", "nodes", "ns", "namespace", "namespaces"}
)

var rootCmd = &cobra.Command{
	Use:   fmt.Sprintf("%s command [resource name]", commandName),
	Short: fmt.Sprintf("kubectl-%s is interactive plugin for kubectl", commandName),
	Args:  cobra.MinimumNArgs(1),
	Long: strings.ReplaceAll(`
Kubectl-{COMMAND_NAME} is an interactive kubectl plugin which wraps kubectl commands.

Examples:

  # Print an interactive list of namespaces and describe the chosen one
  kubectl {COMMAND_NAME} describe namespaces

  # Print an interactive list of pods filtered by --like <filter> and describe the chosen one
  kubectl {COMMAND_NAME} describe pods --like nginx

  # Print an interactive list of configmap filtered by -n <namespace> and edit the chosen one
  kubectl {COMMAND_NAME} edit configmap -n kube-system

  # Print an interactive list of pods filtered by --like <filter> and -f <exec extra flags>  and exec the chosen one
  kubectl {COMMAND_NAME} exec --like nginx -f "it bash"

  # Print an interactive list of pods filtered by --like <filter> and -f <exec extra flags>  and show the chosen pod logs
  kubectl {COMMAND_NAME} logs --like nginx -f "-f"

  # Print an interactive list of deployments and delete the chosen one
  kubectl {COMMAND_NAME} delete deployment

`, "{COMMAND_NAME}", commandName),
	Run: func(cmd *cobra.Command, args []string) {

		log.WithFields(log.Fields{"args": args, "len": len(args)}).Debug("initialize interactive plugin")

		// resourceType describes the available types of Kubernetes resources (pod|configmap and etc)
		var resourceType string

		// The kubectl actions for example: describe, edit, delete, exec and all kubectl actions
		action := args[0]

		// creates kubectl comand
		commandArgs := []string{action}

		// When interactive plugin gets only one argument (Example: exec|log) we are mapping the command type to resource type
		if len(args) == 1 {
			switch args[0] {
			case "logs", "exec":
				resourceType = "pod"
			case "drain", "cordon", "uncordon":
				resourceType = "node"
			default:
				log.WithField("action", args[0]).Fatal("command not supported")
			}
		} else {
			// Set the given resource name (Example: edit configmap)
			resourceType = args[1]
			commandArgs = append(commandArgs, resourceType)
		}

		log.WithField("resource_type", resourceType).Info("given resource type")

		var workingKubeConfig string
		if kubeConfigPath == "" {
			// Get the user home directory (~/) to find the full kubeconfig path
			usr, err := user.Current()
			if err != nil {
				log.WithError(err).Fatal("could not get user home directory path")
			}

			workingKubeConfig = fmt.Sprintf("%s/%s", usr.HomeDir, defaultKubeConfigPath)
		} else {
			workingKubeConfig = kubeConfigPath
		}

		if !fileExists(workingKubeConfig) {
			log.WithField("kubeconfig_path", workingKubeConfig).Fatal("kubeconfig file not found in path")
		}

		kubeConfigPaths := []string{workingKubeConfig}

		// Set interactive configuration
		config := &interactive.Config{
			SelectCluster:   selectCluster,   // If present, the user needs to select one of the clusters from the kubeconfig
			AllNamespaces:   allNamespaces,   // If present, search the resource in all namespaces
			Namespace:       namespace,       // If present search the resource from a given namespace
			Like:            like,            // If present, filter the resources which contain the given value
			Random:          random,          // If present, select random resource
			KubeConfigPaths: kubeConfigPaths, // Kubeconfig file paths
		}

		r, err := interactive.NewInteractive(config)
		if err != nil {
			log.Fatal(err)
		}

		resource, err := r.SelectResource(resourceType)
		if err != nil {
			log.Fatal(err)
		}

		// add the resource name to kubectl command
		commandArgs = append(commandArgs, resource.Name)

		_, found := find(ignoreNamespaceSet, resourceType)
		if !found {
			// Adding namespace flag
			commandArgs = append(commandArgs, "-n", resource.Namespace)
		}

		// Append extra from to kubectl command.
		// For example kubectl interactive exec -f "-it sh"
		commandArgs = append(commandArgs, strings.Split(flags, " ")...)

		err = command.Run("kubectl", commandArgs)
		if err != nil {
			log.Fatal(err)
		}

	},
}

// Execute adds all child commands to the root command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// init cli flags
func init() {
	cobra.OnInitialize(initLogger)

	rootCmd.PersistentFlags().BoolVarP(&selectCluster, "select-cluster", "s", false, "Select cluster from .kube config file")
	rootCmd.PersistentFlags().BoolVarP(&allNamespaces, "all-namespaces", "A", false, "If present, list the requested object(s) across all namespaces")
	rootCmd.PersistentFlags().StringVarP(&namespace, "namespace", "n", "", "If present, the namespace scope for this CLI request")
	rootCmd.PersistentFlags().StringVarP(&kubeConfigPath, "kubeconfig-path", "", "", fmt.Sprintf("By default the configuration will take from ~/%s unless the flag is present", defaultKubeConfigPath))
	rootCmd.PersistentFlags().StringVarP(&like, "like", "l", "", "If present, the requested resources response will be filter by given value")
	rootCmd.PersistentFlags().StringVarP(&flags, "flags", "f", "", "Append kubectl flags")
	rootCmd.PersistentFlags().BoolVarP(&random, "random", "r", false, "If present, one of the resources will select automatically")
	rootCmd.PersistentFlags().StringVarP(&lv, "log-level", "v", "error", "log level (trace|debug|info|warn|error|fatal|panic)")

}

// initLogger sets application log level
func initLogger() {
	switch lv {
	case "trace":
		log.SetLevel(log.TraceLevel)
	case "debug":
		log.SetLevel(log.DebugLevel)
	case "info":
		log.SetLevel(log.InfoLevel)
	case "warn":
		log.SetLevel(log.WarnLevel)
	case "error":
		log.SetLevel(log.ErrorLevel)
	case "fatal":
		log.SetLevel(log.FatalLevel)
	case "panic":
		log.SetLevel(log.PanicLevel)
	default:
		log.SetLevel(log.ErrorLevel)
	}
}

// Find string in slice
func find(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

// fileExists check if file exists
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
