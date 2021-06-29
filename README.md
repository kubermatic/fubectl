# fubectl
Because it's fancy-kubectl!

## Prerequisites?
* [fzf](https://github.com/junegunn/fzf)
* [kubectl](https://github.com/kubernetes/kubernetes)
* [kubectl tree](https://github.com/ahmetb/kubectl-tree)
* [kubectl neat](https://github.com/itaysk/kubectl-neat) - only for kexp to work
* [jq](https://stedolan.github.io/jq/)

## Installation

You can directly download the [`fubectl.source`](https://rawgit.com/kubermatic/fubectl/master/fubectl.source)
and save it in some directory.

Download:
```bash
curl -LO https://rawgit.com/kubermatic/fubectl/master/fubectl.source
```

then add to your .bashrc/.zshrc file:
```bash
[ -f <path-to>/fubectl.source ] && source <path-to>/fubectl.source
```

Alternatively you can install fubectl using the ZSH plugin manager of your
choice.

## What can it do?

### k - alias for kubectl

Like g for git but 133% more effective!

Examples:
 - `k get nodes`
 - `k get pods`
 - `k version --short`

Usage:
![kGif](./demo_src/k.gif)

---

### kw - alias for 'watch kubectl'


Examples:
 - `kw nodes`
 - `kw pods`
 - `kw nodes,pods,services`

---

### kall - All pods in all namespaces

Get all pods

Usage:
![kGif](./demo_src/kall.gif)

---

### kwall - Watch all pods in all namespaces

Watch all pods in all namespaces every 2 seconds.

Usage:
![kGif](./demo_src/kwall.gif)

---

### kdes - Describe a resource

Examples:
- `kdes pod`
- `kdes service`
- `kdes nodes`

Usage:
![kGif](./demo_src/kdes.gif)

---

### kdel - Delete a resource

Examples:
- `kdel pod`
- `kdel secret`
- `kdel pvc`

Usage:
![kGif](./demo_src/kdel.gif)

---

### klog - Print the logs for a container in a pod

Examples:
- `klog` - Print the last 10 lines
- `klog 100` - Print the last 100 lines
- `klog 250 -f` - Print the last 250 lines and follow the output, like `tail -f`
- `klog 50 -p` - Print the last 50 lines of the previous container

Usage:
![kGif](./demo_src/klog.gif)

---

### kex - Execute a command in a container

Examples:
- `kex bash` - Start a bash in a container
- `kex date` - Print the date in a container

Usage:
![kGif](./demo_src/kex.gif)

---

### kfor - Forward one or more local ports to a pod

Examples:
- `kfor 8000` - Forwards port 8000 to a pod
- `kfor 8000:80` Fowards local port 8000 to a pod's port 80

Usage:
![kGif](./demo_src/kfor.gif)

---

### ksearch - Search for string in resources

Examples:
- `// TODO`

Usage:
![kGif](./demo_src/ksearch.gif)

---

### kcl - Displays one or many contexts from the kubeconfig file
Context list

Usage:
![kGif](./demo_src/kcl.gif)
---
### kcs - Sets the current context

Usage:
![kGif](./demo_src/kcs.gif)

---

### kcns - Switch the default namespace

`kcns` - Set the current default namespace from list
`kcns kube-system` - Set kube-system as default namespace immediately

Usage:
![kGif](./demo_src/kcns.gif)
---

### kdebug - Start a debugging Pod in a Cluster

Usage:
![kGif](./demo_src/kdebug.gif)

---

### kp - Open the Kubernetes dashboard

Opens `localhost:8001/ui` in your browser and runs `kubectl proxy`

---

## Extra!
Do you want to have the current kubecontext in your prompt?:
```bash
export PS1="\[$(kube_ctx_name)\] $PS1"
```

for the current namespace (this is currently slow, because it calls kubectl every time):
```bash
export PS1="\[$(kube_ctx_namespace)\] $PS1"
```


## Troubleshooting

If you encounter issues [file an issue][1] or talk to us on the [#fubectl channel][12] on the [Kubermatic Slack][15].

## Contributing

Thanks for taking the time to join our community and start contributing!

Feedback and discussion are available on [Kubermatic Slack][15].

### Before you start

* Please familiarize yourself with the [Code of Conduct][4] before contributing.
* See [CONTRIBUTING.md][2] for instructions on the developer certificate of origin that we require.

### Pull requests

* We welcome pull requests. Feel free to dig through the [issues][1] and jump in.

[1]: https://github.com/kubermatic/fubectl/issues
[2]: https://github.com/kubermatic/fubectl/blob/master/CONTRIBUTING.md
[4]: https://github.com/kubermatic/fubectl/blob/master/CODE_OF_CONDUCT.md

[12]: https://kubermatic.slack.com/messages/fubectl
[15]: http://slack.kubermatic.io/
