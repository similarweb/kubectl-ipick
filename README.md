![Lint](https://github.com/similarweb/kubectl-interactive/workflows/Lint/badge.svg)
![Fmt](https://github.com/similarweb/kubectl-interactive/workflows/Fmt/badge.svg)
# kubectl-interactive
A kubectl plugin that is easy-to-use, flexible, and interactive 
You will have the ability to roam around the Kubernetes world with a blazing fast selection menu
# Introduction
Let's say we'd like to edit a resource in kubernetes using native kubectl and compare it to using kubectl-interactive.
First, we'll start by trying to filter out by the relevant resource type and identify the correct resource.
So we'll run:

```kubectl get configmap```

In some cases, we just noticed it's not the right namespace, so we'll run again:

```kubectl get configmap -n <namespace>```

Then, we'll need to go over the list, identify our resource and copy the resource ID and trigger edit command.

```kubectl edit configmap <configmap_id> -n <namespace>```

This is basically what we need to do for every single resource in our cluster as resource IDs keeps changing.
Instead, we'll leverage kubectl-interactive to minimize the number of operations.
In order to modify a single resource with a single operation, we can run:

```kubectl interactive edit configmap --all-namespaces```

This command will query all namespaces and present a list for us to pick from, once we've chosen the right resource, the plugin will run the command we've passed as argument (edit). And we're done!
```
➜ kubectl interactive edit configmap --all-namespaces
ID    NAME                                 NAMESPACE
1     cluster-info                         kube-public
2     coredns                              kube-system
3     nginx                                kube-system
select configmap from the list:
```
You can even use the "like" filter on top of the list to show only resources that contains the supplied string.
```
➜  ~ kubectl interactive describe configmap --all-namespaces --like nginx 
ID    NAME                              NAMESPACE
1     nginx                             kube-system
2     nginx-load-balancer-conf          kube-system
select configmap from the list:
```
Recommendation: make an alias to kubectl interactive to increase the speed of your work even more!

```alias kia='kubectl interactive'```

Or use it globally

```alias gkia='kubectl interactive --all-namespaces'```

kubectl-interactive makes our day to day work with kubectl much more faster by listing the resources you wish to edit within a single command.
![kubectl-interactive](/docs/images/demo.gif)
# Installation
### Manual
Supported OS
```
# Linux
export OS=LINUX
# Mac
export OS=DARWIN
# Windows
export OS=WINDOWS
```
Execute:
```
# Get the latest kubectl interactive version
TAG=$(curl --silent "https://api.github.com/repos/similarweb/kubectl-interactive/releases/latest" |grep '"tag_name":' | sed -E 's/.*"v([^"]+)".*/\1/')
```
```
# Download the relevant OS version of the plugin
curl -L https://github.com/similarweb/kubectl-interactive/releases/download/v${TAG}/kubectl-interactive_${TAG}_${OS:-Linux}_x86_64.tar.gz | tar xz && chmod +x kubectl-interactive && mv kubectl-interactive /usr/local/bin
```
```
# Make your life easier and...
# Add the following Alias to your .bashrc|.zshrch|.bash_profile
# Run kubectl interactive plugin on all namespaces
alias kia='f(){ kubectl interactive "$@" --all-namespaces;  unset -f f; }; f'
```
# Usage
```
➜ kubectl interactive --help 
Kubectl-interactive is an interactive kubectl plugin which wraps kubectl commands.
Examples:
  # Print an interactive list of namespaces and describe the chosen one
  kubectl interactive describe namespaces
  # Print an interactive list of pods filtered by --like <filter> and describe the chosen one
  kubectl interactive describe pods --like nginx
  # Print an interactive list of configmap filtered by -n <namespace> and edit the chosen one
  kubectl interactive edit configmap -n kube-system
  # Print an interactive list of pods filtered by --like <filter> and -f <exec extra flags>  and exec the chosen one
  kubectl interactive exec --like nginx -f "it bash"
  # Print an interactive list of pods filtered by --like <filter> and -f <exec extra flags>  and show the chosen pod logs
  kubectl interactive logs --like nginx -f "-f"
  # Print an interactive list of deployments and delete the chosen one
  kubectl interactive delete deployment
Usage:
  interactive command [resource name] [flags]
  interactive [command]
Available Commands:
  help        Help about any command
  version     Print the kubectl-interactive version
Flags:
  -A, --all-namespaces           If present, list the requested object(s) across all namespaces
  -f, --flags string             Append kubectl flags
  -h, --help                     help for interactive
      --kubeconfig-path string   By default the configuration will take from ~/.kube/config unless the flag is present
  -l, --like string              If present, the requested resources response will be filter by given value
  -v, --log-level string         log level (trace|debug|info|warn|error|fatal|panic) (default "error")
  -n, --namespace string         If present, the namespace scope for this CLI request
  -r, --random                   If present, one of the resources will select automatically
  -s, --select-cluster           Select cluster from .kube config file
Use "interactive [command] --help" for more information about a command.
  ```
# Contributing
All pull requests and issues are more than welcome! 
Please see [Contribution guidelines](./CONTRIBUTING.md).