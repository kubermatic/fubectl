# Fubectl
Because it's fancy-kubectl!

## What do I need?
* [fzf](https://github.com/junegunn/fzf)
* [kubectl](https://github.com/kubernetes/kubernetes)
* [jq](https://stedolan.github.io/jq/)

## What can I do?
Command|Params|Description
---|---|---
k | `*` | Like g for git but 133% more effective!
kall||Get all pods
kwall||Watch all pods
kp ||Open kubernetes dashboard
kdes | `kind` | Describe resource
kdel | `kind` | Delete resource
klog | \[`lines` `extra-flag`\] |Fetch log from a container
kex |`*`| execute command in container
kfor | `port` | port-forward a local port to a pod
ksearch | `grep-regex`| search for string in resources
kcs ||Context list
kcs ||Context set
kdebug|| Start debugging Pod in Cluster

## Extra!
Do you wan't to have the current kubecontext in your prompt?:
```bash
export PS1="\[$(kube_ctx_name)\] $PS1"
```
