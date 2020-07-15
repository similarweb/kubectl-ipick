![Lint](https://github.com/similarweb/kubectl-ipick/workflows/Lint/badge.svg)
![Fmt](https://github.com/similarweb/kubectl-ipick/workflows/Fmt/badge.svg)
# kubectl-ipick
A kubectl plugin that is easy-to-use, flexible, and interactive 
You will have the ability to roam around the Kubernetes world with a blazing fast selection menu
# Introduction
Let's say we'd like to edit a resource in kubernetes using native kubectl and compare it to using kubectl-ipick.
First, we'll start by trying to filter out by the relevant resource type and identify the correct resource.
So we'll run:

```kubectl get configmap```

In some cases, we just noticed it's not the right namespace, so we'll run again:

```kubectl get configmap -n <namespace>```

Then, we'll need to go over the list, identify our resource and copy the resource ID and trigger edit command.

```kubectl edit configmap <configmap_id> -n <namespace>```

This is basically what we need to do for every single resource in our cluster as resource IDs keeps changing.
Instead, we'll leverage kubectl-ipick to minimize the number of operations.
In order to modify a single resource with a single operation, we can run:

```kubectl ipick edit configmap --all-namespaces```

This command will query all namespaces and present a list for us to pick from, once we've chosen the right resource, the plugin will run the command we've passed as argument (edit). And we're done!
```
➜ kubectl ipick edit configmap --all-namespaces
ID    NAME                                 NAMESPACE
1     cluster-info                         kube-public
2     coredns                              kube-system
3     nginx                                kube-system
select configmap from the list:
```
You can even use the "like" filter on top of the list to show only resources that contains the supplied string.
```
➜  ~ kubectl ipick describe configmap --all-namespaces --like nginx 
ID    NAME                              NAMESPACE
1     nginx                             kube-system
2     nginx-load-balancer-conf          kube-system
select configmap from the list:
```
Recommendation: make an alias to kubectl ipick to increase the speed of your work even more!

```alias kp='kubectl ipick'```

kubectl-ipick makes our day to day work with kubectl much more faster by listing the resources you wish to edit within a single command.
![kubectl-ipick](/docs/images/usage.gif)
# Installation

### Via krew
Installation via krew (https://github.com/GoogleContainerTools/krew)

```
kubectl krew install ipick
```

### Manual
Supported OS
```
# Linux
export OS=Linux
# Mac
export OS=Darwin
# Windows
export OS=Windows
```
Execute:
```
# Get the latest kubectl ipick version
TAG=$(curl --silent "https://api.github.com/repos/similarweb/kubectl-ipick/releases/latest" |grep '"tag_name":' | sed -E 's/.*"v([^"]+)".*/\1/')
```
```
# Download the relevant OS version of the plugin
curl -L https://github.com/similarweb/kubectl-ipick/releases/download/v${TAG}/kubectl-ipick_v${TAG}_${OS:-Linux}_x86_64.tar.gz | tar xz && chmod +x kubectl-ipick && mv kubectl-ipick /usr/local/bin
```
```
# Make your life easier and...
# Add the following Alias to your .bashrc|.zshrch|.bash_profile
# Run kubectl ipick plugin on all namespaces
alias kp='f(){ kubectl ipick "$@" ;  unset -f f; }; f'
```
# Usage
```
➜ kubectl ipick --help 

Kubectl-ipick is an interactive kubectl plugin which wraps kubectl commands.

Examples:

  # Print an interactive list of namespaces and describe the chosen one
  kubectl ipick describe namespaces

  # Print an interactive list of pods filtered by --like <filter> and describe the chosen one
  kubectl ipick describe pods --like nginx

  # Print an interactive list of configmap filtered by -n <namespace> and edit the chosen one
  kubectl ipick edit configmap -n kube-system

  # Print an interactive list of pods filtered by --like <filter> and -- <exec extra flags>  and exec the chosen one
  kubectl ipick exec --like nginx -- -it bash

  # Print an interactive list of pods filtered by --like <filter> and -- <exec extra flags>  and show the chosen pod logs
  kubectl ipick logs --like nginx -- -f

  # Print an interactive list of deployments and delete the chosen one
  kubectl ipick delete deployment

Usage:
  ipick command [resource name] [flags]
  ipick [command]

Available Commands:
  help        Help about any command
  version     Print the kubectl-ipick version

Flags:
  -A, --all-namespaces           If present, list the requested object(s) across all namespaces
  -h, --help                     help for ipick
      --kubeconfig-path string   By default the configuration will take from ~/.kube/config unless the flag is present
  -l, --like string              If present, the requested resources response will be filter by given value
  -v, --log-level string         log level (trace|debug|info|warn|error|fatal|panic) (default "error")
  -n, --namespace string         If present, the namespace scope for this CLI request
  -r, --random                   If present, one of the resources will select automatically
  -s, --select-cluster           Select cluster from .kube config file

Use "ipick [command] --help" for more information about a command.

```

# Contributing
All pull requests and issues are more than welcome! 
Please see [Contribution guidelines](./CONTRIBUTING.md).