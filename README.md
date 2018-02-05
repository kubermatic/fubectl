# Fubectl
Because it's fancy-kubectl!
---
### What do I need?
* fzf
* kubectl
* jq

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
ksearch | `grep-regex`| search for string in resources
kcs ||Context list
kcs ||Context set
kdebug|| Start debugging Pod in Cluster
