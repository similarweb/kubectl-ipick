![Lint](https://github.com/similarweb/kubectl-interactive/workflows/Lint/badge.svg)
![Fmt](https://github.com/similarweb/kubectl-interactive/workflows/Fmt/badge.svg)

# kubectl-interactive
A kubectl plugin that is easy-to-use, flexible, and interactive 

You will have the ability to roam around the Kubernetes world with a blazing fast selection menu

# Introduction
When you look for a Kubernetes resource to edit (or any kubectl aciton), let's say a configmap.

First, you will have to identify the name of the configmap you wish to edit.

You can list all of the configmaps in your cluster with:

```kubectl get configmap```

Next, find the config map and tell Kubernetes to edit it:

```kubectl edit configmap <configmap_name>```

kubectl-interactive makes our day to day work with kubectl much more faster by listing the resources you wish to edit within a single command.


* Choose the relevant configmap you wish to edit 
        

```
➜ kubectl interactive edit configmap --all-namespaces   

ID    NAME                                 NAMESPACE
1     cluster-info                         kube-public
2     coredns                              kube-system
3     nginx                                kube-system
select configmap from the list:
```

* Wildcard Filtering using `--like` makes it is easy to find matches without having to know the full string value needed to complete the retrieval.

```
➜  ~ kubectl interactive describe configmap --all-namespaces --like nginx 

ID    NAME                              NAMESPACE
1     nginx                             kube-system
2     nginx-load-balancer-conf          kube-system
select configmap from the list:
```


# Usage
```
➜ kubectl interactive --help  

Kubectl-interactive is an interactive kubectl plugin which wraps kubectl commands.

Usage:
  interactive command [resource name] [flags]

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
  ```


## Installation


### Binary

Supporting OS

```
# Linux
OS=LINUX

# Mac
OS=DARWIN
```

```
TAG=$(curl --silent "https://api.github.com/repos/similarweb/kubectl-interactive/releases/latest" |grep '"tag_name":' | sed -E 's/.*"v([^"]+)".*/\1/')

curl -LO https://github.com/similarweb/kubectl-interactive/releases/download/v${TAG}/kubectl-interactive_${TAG}_${OS:-Linux}_x86_64.tar.gz

mkdir -p /tmp/kubectl-interactive
tar -xzvf kubectl-interactive_${TAG}_${OS}_x86_64.tar.gz -C /tmp/kubectl-interactive
chmod +x /tmp/kubectl-interactive/kubectl-interactive

sudo mv /tmp/kubectl-interactive/kubectl-interactive /usr/local/bin

```

### Krew
Working progress