
<a name="v0.76.0"></a>
# [v0.76.0] - 2023-03-14
### Bug Fixes 🐞
- use cross-platform pause image as the containerd sandbox image on Windows ([#106](https://github.com/Azure/aks-engine-azurestack/issues/106))
- kubernetes-azurestack.json uses distro aks-ubuntu-20.04 ([#87](https://github.com/Azure/aks-engine-azurestack/issues/87))
- CoreDNS image not updated after cluster upgrade ([#75](https://github.com/Azure/aks-engine-azurestack/issues/75))
- change reference of cni config to scripts dir ([#71](https://github.com/Azure/aks-engine-azurestack/issues/71))
- add Azure CNI config script to Ubuntu VHD ([#70](https://github.com/Azure/aks-engine-azurestack/issues/70))
- unit test checking e2e configs ([#49](https://github.com/Azure/aks-engine-azurestack/issues/49))
- syntax error in Windows VHD script ([#41](https://github.com/Azure/aks-engine-azurestack/issues/41))
- ensure eth0 addr is set to NIC's primary addr ([#39](https://github.com/Azure/aks-engine-azurestack/issues/39))

### Build 🏭
- rename binary to aks-engine-azurestack ([#35](https://github.com/Azure/aks-engine-azurestack/issues/35))

### Code Style 🎶
- replace usage of deprecated "io/ioutil" golang package

### Continuous Integration 💜
- release wf tag the correct commit
- gen-release-changelog wf creates branch and commit
- update actions/checkout to v3 ([#111](https://github.com/Azure/aks-engine-azurestack/issues/111))
- chocolatey workflow ([#86](https://github.com/Azure/aks-engine-azurestack/issues/86))
- release workflows run no-egress scenarios ([#85](https://github.com/Azure/aks-engine-azurestack/issues/85))
- remove no-egress job from create branch action ([#79](https://github.com/Azure/aks-engine-azurestack/issues/79))
- PR gate runs E2E suite ([#69](https://github.com/Azure/aks-engine-azurestack/issues/69))
- PR checks consume SIG images ([#64](https://github.com/Azure/aks-engine-azurestack/issues/64))
- E2E PR check uses user assigned identity ([#54](https://github.com/Azure/aks-engine-azurestack/issues/54))
- fix variable name in e2e PR check ([#52](https://github.com/Azure/aks-engine-azurestack/issues/52))
- e2e PR check sets tenant ([#51](https://github.com/Azure/aks-engine-azurestack/issues/51))
- e2e PR check does not use AvailabilitySets ([#47](https://github.com/Azure/aks-engine-azurestack/issues/47))
- e2e PR check does not use custom VNET ([#46](https://github.com/Azure/aks-engine-azurestack/issues/46))
- e2e PR check does not use MSI ([#45](https://github.com/Azure/aks-engine-azurestack/issues/45))
- remove Ubuntu Gen2 pipeline ([#32](https://github.com/Azure/aks-engine-azurestack/issues/32))
- fixed ShellCheck errors ([#31](https://github.com/Azure/aks-engine-azurestack/issues/31))
- update golangci + fix lint errors ([#30](https://github.com/Azure/aks-engine-azurestack/issues/30))
- Linux VHD pipeline supports Ubuntu 20.04 ([#26](https://github.com/Azure/aks-engine-azurestack/issues/26))
- Linux VHD pipeline creates SIG image version ([#23](https://github.com/Azure/aks-engine-azurestack/issues/23))
- remove 1.24 no-egress requirement ([#4947](https://github.com/Azure/aks-engine-azurestack/issues/4947))
- remove 1.20 no-egress requirement ([#4944](https://github.com/Azure/aks-engine-azurestack/issues/4944))

### Documentation 📘
- Enable azurestack-csi-driver addon for mixed clusters ([#96](https://github.com/Azure/aks-engine-azurestack/issues/96))
- remove Azure as a target cloud ([#43](https://github.com/Azure/aks-engine-azurestack/issues/43))
- rename binary name in all markdown files ([#42](https://github.com/Azure/aks-engine-azurestack/issues/42))

### Features 🌈
- add support for Kubernetes v1.24.11 ([#109](https://github.com/Azure/aks-engine-azurestack/issues/109))
- migrate from Pod Security Policy to Pod Security admission ([#94](https://github.com/Azure/aks-engine-azurestack/issues/94))
- DISA Ubuntu 20.04 STIG compliance ([#83](https://github.com/Azure/aks-engine-azurestack/issues/83))
- support UbuntuServer 20.04 nodes ([#27](https://github.com/Azure/aks-engine-azurestack/issues/27))

### Maintenance 🔧
- support Kubernetes v1.25.7 ([#105](https://github.com/Azure/aks-engine-azurestack/issues/105))
- upgrade coredns to v1.9.4 ([#98](https://github.com/Azure/aks-engine-azurestack/issues/98))
- upgrade containerd to 1.5.16 ([#95](https://github.com/Azure/aks-engine-azurestack/issues/95))
- upgrade pause to v3.8 ([#93](https://github.com/Azure/aks-engine-azurestack/issues/93))
- update golang toolchain to v1.19 ([#90](https://github.com/Azure/aks-engine-azurestack/issues/90))
- update registries for nvidia and k8s.io components ([#88](https://github.com/Azure/aks-engine-azurestack/issues/88))
- remove package apache2-utils from VHD ([#82](https://github.com/Azure/aks-engine-azurestack/issues/82))
- update default windows image to jan 2023 ([#77](https://github.com/Azure/aks-engine-azurestack/issues/77))
- Update Windows VHD packer job to use Jan 2023 patches ([#76](https://github.com/Azure/aks-engine-azurestack/issues/76))
- change base image sku and version to azurestack ([#74](https://github.com/Azure/aks-engine-azurestack/issues/74))
- set fsType to ext4 in supported storage classes ([#73](https://github.com/Azure/aks-engine-azurestack/issues/73))
- enable v1.23.15 & v1.24.9, use ubuntu 20.04 as default, force containerd runtime ([#68](https://github.com/Azure/aks-engine-azurestack/issues/68))
- include relevant updates from v0.75.0 ([#56](https://github.com/Azure/aks-engine-azurestack/issues/56))
- include relevant updates from v0.74.0 ([#55](https://github.com/Azure/aks-engine-azurestack/issues/55))
- include relevant updates from v0.73.0 ([#53](https://github.com/Azure/aks-engine-azurestack/issues/53))
- remove kv-fluxvolume addon ([#48](https://github.com/Azure/aks-engine-azurestack/issues/48))
- prefer ADO for PR E2E check ([#38](https://github.com/Azure/aks-engine-azurestack/issues/38))
- added e2e to PR workflow ([#36](https://github.com/Azure/aks-engine-azurestack/issues/36))
- execute pull request workflow on master branch ([#34](https://github.com/Azure/aks-engine-azurestack/issues/34))
- removed docker registry image from VHD ([#28](https://github.com/Azure/aks-engine-azurestack/issues/28))
- update Go toolchain to 1.17 ([#29](https://github.com/Azure/aks-engine-azurestack/issues/29))
- upgrade runc to v1.1.4 ([#25](https://github.com/Azure/aks-engine-azurestack/issues/25))
- upgrade Azure CNI to v1.4.32 ([#24](https://github.com/Azure/aks-engine-azurestack/issues/24))
- Update go mod vendor directories ([#22](https://github.com/Azure/aks-engine-azurestack/issues/22))
- renamed go packages and modules ([#20](https://github.com/Azure/aks-engine-azurestack/issues/20))
- make jadarsie the sole approver ([#15](https://github.com/Azure/aks-engine-azurestack/issues/15))
- updating Windows VHDs for next aks-engine release ([#4936](https://github.com/Azure/aks-engine-azurestack/issues/4936))
- Update Azure constants ([#4935](https://github.com/Azure/aks-engine-azurestack/issues/4935))
- pin github actions go versions to go.mod ([#4934](https://github.com/Azure/aks-engine-azurestack/issues/4934))
- Update Azure constants ([#4931](https://github.com/Azure/aks-engine-azurestack/issues/4931))
- upgrade container runtimes ([#4932](https://github.com/Azure/aks-engine-azurestack/issues/4932))

### Security Fix 🛡️
- bump x/net and x/crypto ([#104](https://github.com/Azure/aks-engine-azurestack/issues/104))
- upgrade k8s client to v0.21.10 ([#21](https://github.com/Azure/aks-engine-azurestack/issues/21))

### Testing 💚
- e2e suite validates an existing PV works after a cluster upgrade ([#92](https://github.com/Azure/aks-engine-azurestack/issues/92))
- e2e sets ImageRef in all linux nodepools ([#65](https://github.com/Azure/aks-engine-azurestack/issues/65))

#### Please report any issues here: https://github.com/Azure/aks-engine-azurestack/issues/new
[Unreleased]: https://github.com/Azure/aks-engine-azurestack/compare/v0.76.0...HEAD
[v0.76.0]: https://github.com/Azure/aks-engine-azurestack/compare/v0.70.2...v0.76.0
