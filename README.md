# Fubectl
Because it's fancy-kubectl!

## What do I need?
* [fzf](https://github.com/junegunn/fzf)
* [kubectl](https://github.com/kubernetes/kubernetes)
* [jq](https://stedolan.github.io/jq/)

## What can I do?

### k
Like g for git but 133% more effective!

Params:
* `*`

Usage:
![kGif](./demo_src/k.gif)
---
### kall
Get all pods

Usage:
![kGif](./demo_src/kall.gif)
---
### kwall
Watch all pods

Usage:
![kGif](./demo_src/kwall.gif)
---
### kdes
Describe a resource

Params:
* `kind`

Usage:
![kGif](./demo_src/kdes.gif)
---
### kdel
Delete a resource

Params:
* `kind`

Usage:
![kGif](./demo_src/kdel.gif)
---
### klog
Fetch log from a container

Params:
* \[`lines` `extra-flag`\](optional)

Usage:
![kGif](./demo_src/klog.gif)
---
### kex
Execute a command in a container

Params:
* `*`

Usage:
![kGif](./demo_src/kex.gif)
---
### kfor
Port-forward a local port to a pod

Params:
* `port` | `host-port:container-port`

Usage:
![kGif](./demo_src/kfor.gif)
---
### ksearch
Search for string in resources

Params:
* `grep-regex`

Usage:
![kGif](./demo_src/ksearch.gif)
---
### kcl
Context list

Usage:
![kGif](./demo_src/kcl.gif)
---
### kcs
Context set

Usage:
![kGif](./demo_src/kcs.gif)
---
### kdebug
Start debugging Pod in Cluster

Usage:
![kGif](./demo_src/kdebug.gif)
---
### kp
Opens a Kubernetes dashboard

---

## Installation

You can directly download the [`fubeconfig.source`](https://rawgit.com/realfake/fubectl/master/fubectl.source)
and save it in some directory.

Download:
```sh
curl -LO https://rawgit.com/realfake/fubectl/master/fubectl.source
```

then add to your .bashrc/.zshrc file:
```sh
[ -f <path-to>/fubeconfig.source ] && source <path-to>/fubeconfig.source
```

## Extra!
Do you wan't to have the current kubecontext in your prompt?:
```bash
export PS1="\[$(kube_ctx_name)\] $PS1"
```
