# Patching Cluster Nodes

AKS Engine allows users to control node OS patching through the `linuxProfile` API model configuration object.

## Enable Unattended Upgrades

By setting configuration property `linuxProfile.enableUnattendedUpgrades` to `false`, users are able to disable unatteded upgrades.
If set to `true`, the default OS behavior remains unaltered.

On Azure Stack Hub, the default value of `linuxProfile.enableUnattendedUpgrades` is `true`.

AKS Engine's mechanism to prevent node patching is to create file `/etc/apt/apt.conf.d/99periodic` with the following content:

```
APT::Periodic::Update-Package-Lists "0";
APT::Periodic::Download-Upgradeable-Packages "0";
APT::Periodic::AutocleanInterval "0";
APT::Periodic::Unattended-Upgrade "0";
```

More information [here](https://help.ubuntu.com/community/AutomaticSecurityUpdates) for Ubuntu Server's default behavior.

## Run Unattended Upgrades On Bootstrap

Setting the configuration property `linuxProfile.runUnattendedUpgradesOnBootstrap` to `true`
forces the execution of the unattended upgrade process right after each Linux node VM comes online for the first time.

When `linuxProfile.runUnattendedUpgradesOnBootstrap` is set to `true`,
the upgrade process will be triggered regardless of the value of `linuxProfile.enableUnattendedUpgrades`.

On Azure Stack Hub, the default value of `linuxProfile.runUnattendedUpgradesOnBootstrap` is `false`.

## When to Enable Node Patching

The decision of enabling unattended upgrades or not is a trade off between security and consistency.

Enabling unattended upgrades ensures that the latest security patches will be downloaded and installed.
On the contrary, disabling them allows the adoptiong of the `immutable infrastructure` paradigm.

## Orchestrate Node Reboots

Some OS patches are only effective after a node reboot.
Deploying the [kured](https://github.com/weaveworks/kured) daemonset ensure that the nodes are rebooted in a non-disruptive way.

The `kured` daemonset can be installed by following this [instructions](https://kured.dev/docs/installation/):

```bash
# Find the appropiate version for the target cluster https://kured.dev/docs/installation/#kubernetes--os-compatibility
version=$(curl -s https://api.github.com/repos/kubereboot/kured/releases | jq -r '.[0].tag_name')

# Download resources yaml
curl -sLO "https://github.com/kubereboot/kured/releases/download/${version}/kured-${version}-dockerhub.yaml"

# Add nodeSelector to the DaemonSet pod spec to constraint to Linux nodes
# spec:
#   nodeSelector:
#     kubernetes.io/os: linux

# Override tolerations to the DaemonSet pod spec to deploy to control plane nodes
# spec:
#   tolerations:
#   - key: node-role.kubernetes.io/control-plane
#     effect: NoSchedule
#     value: "true"
#   - key: node-role.kubernetes.io/master
#     effect: NoSchedule
#     value: "true"

# Check https://kured.dev/docs/configuration/ for configuration options
kubectl apply -f kured-${version}-dockerhub.yaml
```
