#fix ZSH -  perform field splitting - http://zsh.sourceforge.net/Doc/Release/Options.html
if [ -n "$ZSH_VERSION" ]; then
    setopt SH_WORD_SPLIT
fi

# helper functions
if [ -z "$FUBECTL_MSYS_HACK" ]; then
    alias _kctl_tty="kubectl"
    alias _inline_fzf="fzf --multi --ansi -i -1 --height=50% --reverse -0 --header-lines=1 --inline-info --border"
    alias _inline_fzf_nh="fzf --multi --ansi -i -1 --height=50% --reverse -0 --inline-info --border"
else # hack for Msys2 shell, where fzf doesn't support mintty
    alias _kctl_tty="winpty kubectl"
    alias _inline_fzf="fzf --multi --ansi -i -1 --reverse -0 --header-lines=1 --inline-info"
    alias _inline_fzf_nh="fzf --multi --ansi -i -1 --reverse -0 --inline-info"
fi

function _isClusterSpaceObject() {
  # caller is responsible for assuring non-empty "$1"
    obj="$1"
    kubectl api-resources --namespaced=false \
      | awk '(apiidx){print substr($0, 0, apiidx),substr($0, kindidx) } (!apiidx){ apiidx=index($0, " APIVERSION");kindidx=index($0, " KIND")}' \
      | grep -iq "\<${obj}\>"
}

# [k] like g for git but 233% as effective!
alias k="kubectl"

# [kw] watch resources of any KIND in current namespace, usage: kw KIND[,KIND2,...]
alias kw="watch kubectl get"

# [ka] get all pods from current namespace
alias ka="kubectl get pods"

# [kall] get all pods from cluster
alias kall="kubectl get pods --all-namespaces"

# [kwa] watch all pods in current namespace
alias kwa="watch kubectl get pods"

# [kwall] watch all pods in cluster
alias kwall="watch kubectl get pods --all-namespaces"

# TODO use "open" instead of "xdg-open" on a mac - also redirect xdg-open std{out,err} to /dev/null
# [kp] open kubernetes dashboard with proxy
alias kp="xdg-open 'http://localhost:8001/api/v1/namespaces/kube-system/services/https:kubernetes-dashboard:/proxy/' & kubectl proxy"

# [kwatchn] watch a resource of KIND in current namespace, usage: kwatchn [KIND] - if KIND is empty then pod is used
function kwatchn() {
    local kind="${1:-pod}"
    kubectl get "${kind}" | _inline_fzf | awk '{print $1}' | xargs -r watch kubectl get "${kind}"
}

# [kwatch] watch a resource of KIND in cluster, usage: kwatch [KIND] - if KIND is empty then pod is used
function kwatch() {
    local kind="${1:-pod}"
    if _isClusterSpaceObject "$kind" ; then
        kubectl get "${kind}" | _inline_fzf | awk '{print $1}' | xargs -r watch kubectl get "${kind}"
    else
        kubectl get "${kind}" --all-namespaces | _inline_fzf | awk '{print $1, $2}' | xargs -r watch kubectl get "${kind}" -n
    fi
}

# [kcmd] create a pod from IMAGE (ubuntu by default) and execute CMD (bash by default), usage: kcmd [CMD] [IMAGE]
function kcmd() {
    local cmd="$1"
    local image="${2:-ubuntu}"
    local ns="$(kubectl get ns | _inline_fzf | awk '{print $1}')"
    if [ -n "$cmd" ]; then
        kubectl run shell-$RANDOM --namespace $ns --rm -i --tty --image ${image} -- /bin/sh -c "${cmd}"
    else
        kubectl run shell-$RANDOM --namespace $ns --rm -i --tty --image ${image} -- /bin/bash
    fi
}

# [kube_ctx_name] get the current context
function kube_ctx_name() {
    kubectl config current-context
}

# [kube_ctx_namespace] get current namespace
function kube_ctx_namespace() {
    local default_ns="$(kubectl config view --minify|grep namespace: |sed 's/namespace: //g'|tr -d ' ')"
    default_ns="${default_ns:-default}"
    echo "$default_ns"
}

# [kgetn] get resource of KIND from current namespace, usage: kgetn [KIND] - if KIND is empty then pod is used
function kgetn() {
    local kind="${1:-pod}"
    kubectl get "$kind" | _inline_fzf | awk '{print $1}' | xargs -r kubectl get -o yaml "$kind"
}

# [kget] get resource of KIND from cluster, usage: kget [KIND] - if KIND is empty then pod is used
function kget() {
    local kind="${1:-pod}"
    if _isClusterSpaceObject "$kind" ; then
        kubectl get "$kind" | _inline_fzf | awk '{print $1}' | xargs -r kubectl get -o yaml "$kind"
    else
        kubectl get "$kind" --all-namespaces | _inline_fzf | awk '{print $1, $2}' | xargs -r kubectl get -o yaml "$kind" -n
    fi
}

# [kexp] as former `--export` field removes unwanted metadata - usage: COMMAND | kexp
function kexp() {
    if [ -t 0 ]; then
        echo "kexp has no piped input!"
        echo "usage: COMMAND | kexp"
    else
        # remove not neat fields
        kubectl neat
    fi
}

# [kget-exp] get a resource by its YAML as former `--export` flag
function kget-exp() {
    kget "$@" | kexp
}

# [kedn] edit resource of KIND from current namespace, usage: kedn [KIND] - if KIND is empty then pod is used
function kedn() {
    local kind="${1:-pod}"
    kubectl edit "${kind}" "$(kubectl get "${kind}" | _inline_fzf | awk '{print $1}')"
}

# [ked] edit resource of KIND from cluster, usage: ked [KIND] - if KIND is empty then pod is used
function ked() {
    local kind="${1:-pod}"
    local edit_args
    if _isClusterSpaceObject $kind ; then
        edit_args=( $(kubectl get "$kind" | _inline_fzf | awk '{print $1}') )
    else
        edit_args=( $(kubectl get "$kind" --all-namespaces | _inline_fzf | awk '{print "-n", $1, $2}') )
    fi
    kubectl edit "$kind" ${edit_args[*]}
}

# [kdesn] describe resource of KIND in current namespace, usage: kdesn [KIND] - if KIND is empty then pod is used
function kdesn() {
    local kind="${1:-pod}"
    kubectl get "$kind" | _inline_fzf | awk '{print $1}' | xargs -r kubectl describe "$kind"
}

# [kdes] describe resource of KIND in cluster, usage: kdes [KIND] - if KIND is empty then pod is used
function kdes() {
    local kind="${1:-pod}"
    if _isClusterSpaceObject "$kind" ; then
        kubectl get "$kind" | _inline_fzf | awk '{print $1}' | xargs -r kubectl describe "$kind"
    else
        kubectl get "$kind" --all-namespaces | _inline_fzf | awk '{print $1, $2}' | xargs -r kubectl describe "$kind" -n
    fi
}

# [kdeln] delete resource of KIND in current namespace, usage: kdeln [KIND] - if KIND is empty then pod is used
function kdeln() {
    local kind="${1:-pod}"
    kubectl get "$kind" | _inline_fzf | awk '{print $1}' | xargs -r -p kubectl delete "$kind"
}

# [kdel] delete resource of KIND in cluster, usage: kdel [KIND] - if KIND is empty then pod is used
function kdel() {
    local kind="${1:-pod}"
    if _isClusterSpaceObject "$kind" ; then
        kubectl get "$kind" | _inline_fzf | awk '{print $1}' | xargs -r -p kubectl delete "$kind"
    else
        kubectl get "$kind" --all-namespaces | _inline_fzf | awk '{print $1, $2}' | xargs -r -p kubectl delete "$kind" -n
    fi
}

function _klog_usage() {
  cat <<'EOF'
Usage: klog[n] [LINECOUNT] [options]

First argument is interpreted as LINECOUNT if it matches integer syntax.
Additional `options` are passed on (see `kubectl logs --help` for details).
EOF
}

# [klogn] fetch log from container in current namespace
function klogn() {
    [ "$1" = "--help" ] && _klog_usage && return
    local line_count=10
    if [[ $1 =~ ^[-]{0,1}[0-9]+$ ]]; then
        line_count="$1"
        shift
    fi

    local pod=$(kubectl get po | _inline_fzf | awk '{print $1}')
    [ -z "${pod}" ] && printf "klogn: no pods found. no logs can be shown.\n" && return
    local containers_out=$(kubectl get po "${pod}" -o=jsonpath='{.spec.containers[*].name} {.spec.initContainers[*].name}' | sed 's/ $//')
    local container_choosen=$(echo "$containers_out" |  tr ' ' "\n" | _inline_fzf_nh)
    _kctl_tty logs "${pod}" -c "${container_choosen}" --tail="${line_count}" "$@"
}

# [klog] fetch log from container in the cluster
function klog() {
    [ "$1" = "--help" ] && _klog_usage && return
    local line_count=10
    if [[ $1 =~ ^[-]{0,1}[0-9]+$ ]]; then
        line_count="$1"
        shift
    fi

    local arg_pair=$(kubectl get po --all-namespaces | _inline_fzf | awk '{print $1, $2}')
    [ -z "$arg_pair" ] && printf "klog: no pods found. no logs can be shown.\n" && return
    local containers_out=$(echo "$arg_pair" | xargs kubectl get po -o=jsonpath='{.spec.containers[*].name} {.spec.initContainers[*].name}' -n | sed 's/ $//')
    local container_choosen=$(echo "$containers_out" |  tr ' ' "\n" | _inline_fzf_nh)
    _kctl_tty logs -n ${arg_pair} -c "${container_choosen}" --tail="${line_count}" "$@"
}

# [kkonfig] select a file in current directory and set it as $KUBECONFIG
function kkonfig() {
    local kubeconfig
    kubeconfig=$(_inline_fzf_nh) || return
    export KUBECONFIG=$PWD/$kubeconfig
    exec $SHELL
}

# [kexn] execute command in container from current namespace, usage: kexn CMD [ARGS]
function kexn() {
    [ -z "$1" ] && printf "kexn: missing argument(s).\nUsage: kexn CMD [ARGS]\n" && return 255
    local arg_pair=$(kubectl get po | _inline_fzf | awk '{print $1}')
    [ -z "$arg_pair" ] && printf "kexn: no pods found. no execution.\n" && return
    local containers_out=$(echo "$arg_pair" | xargs kubectl get po -o=jsonpath='{.spec.containers[*].name}' -n)
    local container_choosen=$(echo "$containers_out" |  tr ' ' "\n" | _inline_fzf_nh)
    _kctl_tty exec -it ${arg_pair} -c "${container_choosen}" -- "$@"
}

# [kex] execute command in container from cluster, usage: kex CMD [ARGS]
function kex() {
    [ -z "$1" ] && printf "kex: missing argument(s).\nUsage: kex CMD [ARGS]\n" && return 255
    local arg_pair=$(kubectl get po --all-namespaces | _inline_fzf | awk '{print $1, $2}')
    [ -z "$arg_pair" ] && printf "kex: no pods found. no execution.\n" && return
    local containers_out=$(echo "$arg_pair" | xargs kubectl get po -o=jsonpath='{.spec.containers[*].name}' -n)
    local container_choosen=$(echo "$containers_out" |  tr ' ' "\n" | _inline_fzf_nh)
    _kctl_tty exec -it -n ${arg_pair} -c "${container_choosen}" -- "$@"
}

# [kforn] port-forward a container port from current namesapce, usage: kforn LOCAL_PORT:CONTAINER_PORT
function kforn() {
    local port="$1"
    [ -z "$port" ] && printf "kforn: missing argument.\nUsage: kforn LOCAL_PORT:CONTAINER_PORT\n" && return 255
    local arg_pair="$(kubectl get po | _inline_fzf | awk '{print $1}')"
    [ -z "$arg_pair" ] && printf "kforn: no pods found. no forwarding.\n" && return
    _kctl_tty port-forward ${arg_pair} ${port}
}

# [kfor] port-forward a container port from cluster, usage: kfor LOCAL_PORT:CONTAINER_PORT
function kfor() {
    local port="$1"
    [ -z "$port" ] && printf "kfor: missing argument.\nUsage: kfor LOCAL_PORT:CONTAINER_PORT\n" && return 255
    local arg_pair="$(kubectl get po --all-namespaces | _inline_fzf | awk '{print $1, $2}')"
    [ -z "$arg_pair" ] && printf "kfor: no pods found. no forwarding.\n" && return
    _kctl_tty port-forward -n ${arg_pair} ${port}
}

# [ksearch] search for string in resources
function ksearch() {
    local search_query="$1"
    [ -z "$search_query" ] && printf "ksearch: missing argument.\nUsage: ksearch SEARCH_QUERY\n" && return 255
    for ns in $(kubectl get --export -o=json ns | jq -r '.items[] | .metadata.name'); do
        kubectl --namespace="${ns}" get --export -o=json \
            deployment,ingress,daemonset,secrets,configmap,service,serviceaccount,statefulsets,pod,endpoints,customresourcedefinition,events,networkpolicies,persistentvolumeclaims,persistentvolumes,replicasets,replicationcontrollers,statefulsets,storageclasses | \
        jq '.items[]' -c | \
        grep "$search_query" | \
        jq -r  '. | [.kind, .metadata.name] | @tsv' | \
        awk -v prefix="$ns" '{print "kubectl get -n " prefix " " $0}'
    done
}

# [kcl] context list
alias kcl='kubectl config get-contexts'

# [kcs] context set
function kcs() {
    local context="$(kubectl config get-contexts | _inline_fzf | cut -b4- | awk '{print $1}')"
    kubectl config set current-context "${context}"
}

# [kcns] context set default namespace
function kcns() {
    local ns="$1"
    if [ -z "$ns" ]; then
        ns="$(kubectl get ns | _inline_fzf | awk '{print $1}')"
    fi
    [ -z "$ns" ] && printf "kcns: no namespace selected/found.\nUsage: kcns [NAMESPACE]\n" && return
    kubectl config set-context "$(kubectl config current-context)" --namespace="${ns}"
}

# [kwns] watch pods in a namespace
function kwns() {
    local ns=$(kubectl get ns | _inline_fzf | awk '{print $1}')
    [ -z "$ns" ] && printf "kcns: no namespace selected/found.\nUsage: kwns\n" && return
    watch kubectl get pod -n "$ns"
}

# [ktreen] prints a tree of k8s objects from current namespace, usage: ktreen [KIND]
function ktreen() {
    local kind="$1"
    if [ -z "$kind" ]; then
        local kind="$(kubectl api-resources -o name | _inline_fzf | awk '{print $1}')"
    fi
    kubectl get "$kind" | _inline_fzf | awk '{print $1}' | xargs -r kubectl tree "$kind"
}

# [ktree] prints a tree of k8s objects from cluster, usage: ktree [KIND]
function ktree() {
    local kind="$1"
    if [ -z "$kind" ]; then
        local kind="$(kubectl api-resources -o name | _inline_fzf | awk '{print $1}')"
    fi
    if _isClusterSpaceObject "$kind" ; then
        kubectl get "$kind" | _inline_fzf | awk '{print $1}' | xargs -r kubectl tree "$kind"
    else
        kubectl get "$kind" --all-namespaces | _inline_fzf | awk '{print $1, $2}' | xargs -r kubectl tree "$kind" -n
    fi
}

# [konsole] create root shell on a node
function konsole() {
    local node_hostname="$(kubectl get node --label-columns=kubernetes.io/hostname | _inline_fzf | awk '{print $6}')"
    local ns="$(kubectl get ns | _inline_fzf | awk '{print $1}')"
    local name=shell-$RANDOM
    local overrides='
{
    "spec": {
        "hostPID": true,
        "hostNetwork": true,
        "tolerations": [
            {
                "operator": "Exists",
                "effect": "NoSchedule"
            },
            {
                "operator": "Exists",
                "effect": "NoExecute"
            }
        ],
        "containers": [
            {
                "name": "'$name'",
                "image": "alpine",
                "command": [
                    "/bin/sh"
                ],
                "args": [
                    "-c",
                    "nsenter -t 1 -m -u -i -n -p -- bash"
                ],
                "resources": null,
                "stdin": true,
                "stdinOnce": true,
                "terminationMessagePath": "/dev/termination-log",
                "terminationMessagePolicy": "File",
                "tty": true,
                "securityContext": {
                    "privileged": true
                }
            }
        ],
        "nodeSelector": {
            "kubernetes.io/hostname": "'$node_hostname'"
        }
    }
}
'
    kubectl run $name --namespace "$ns" --rm -it --image alpine --overrides="${overrides}"
}

# [ksecn] decode a value from a secret in current namespace, usage: ksecn
function ksecn() {
    local secret=$(kubectl get secret | _inline_fzf | awk '{print $1}')
    local key=$(kubectl get secret "${secret}" -o go-template='{{- range $k,$v := .data -}}{{- printf "%s\n" $k -}}{{- end -}}' | _inline_fzf_nh)
    kubectl get secret "${secret}" -o go-template='{{ index .data "'$key'" | base64decode }}'
}

# [ksec] decode a value from a secret in cluster, usage: ksec
function ksec() {
    local ns=$(kubectl get ns | _inline_fzf | awk '{print $1}')
    local secret=$(kubectl get secret -n "$ns" | _inline_fzf | awk '{print $1}')
    local key=$(kubectl get secret -n "$ns" "$secret" -o go-template='{{- range $k,$v := .data -}}{{- printf "%s\n" $k -}}{{- end -}}' | _inline_fzf_nh)
    kubectl get secret -n "$ns" "${secret}" -o go-template='{{ index .data "'$key'" | base64decode }}'
}

# [kinstall] Install the required kubectl plugins
function kinstall() {
    kubectl krew install tree
    kubectl krew install neat
}
# [kupdate] Updates kubectl plugins
function kupdate() {
    kubectl krew upgrade
}

#### Kubermatic KKP specific
# [kkp-cluster] Kubermatic KKP - extracts kubeconfig of user cluster and connects it in a new bash
function kkp-cluster() {
    TMP_KUBECONFIG=$(mktemp)
    local cluster="$(kubectl get cluster | _inline_fzf | awk '{print $1}')"
    kubectl get secret admin-kubeconfig -n cluster-$cluster -o go-template='{{ index .data "kubeconfig" | base64decode }}' > $TMP_KUBECONFIG
    KUBECONFIG=$TMP_KUBECONFIG $SHELL
}

# [khelp] show this help message
function khelp() {
    echo "Usage of fubectl"
    echo
    echo "Reduces repetitive interactions with kubectl"
    echo "Find more information at https://github.com/kubermatic/fubectl"
    echo
    echo "Usage:"
    if [ -n "$ZSH_VERSION" ]
    then
        grep -E '^# \[.+\]' "${(%):-%x}"
    else
        grep -E '^# \[.+\]' "${BASH_SOURCE[0]}"
    fi
}
