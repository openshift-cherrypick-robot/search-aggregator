// Copyright (c) 2021 Red Hat, Inc.
// Copyright Contributors to the Open Cluster Management project
module github.com/stolostron/search-aggregator

go 1.18

require (
	github.com/golang/glog v1.0.0
	github.com/gomodule/redigo v1.8.9
	github.com/gorilla/mux v1.8.0
	github.com/kennygrant/sanitize v1.2.4
	github.com/olekukonko/tablewriter v0.0.5 // indirect
	github.com/redislabs/redisgraph-go v2.0.2+incompatible
	github.com/stolostron/klusterlet-addon-controller v0.0.0-20220714103120-43778f205c45
	github.com/stolostron/multicloud-operators-foundation v1.0.0-2021-10-26-20-16-14.0.20220110023249-172fb944faa9
	github.com/stretchr/testify v1.7.1
	k8s.io/apimachinery v0.24.3
	k8s.io/client-go v13.0.0+incompatible
	open-cluster-management.io/api v0.8.0

)

require (
	github.com/Masterminds/semver v1.5.0 // indirect
	github.com/PuerkitoBio/purell v1.1.1 // indirect
	github.com/PuerkitoBio/urlesc v0.0.0-20170810143723-de5bf2ad4578 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/emicklei/go-restful/v3 v3.8.0 // indirect
	github.com/evanphx/json-patch v5.6.0+incompatible // indirect
	github.com/go-logr/logr v1.2.3 // indirect
	github.com/go-openapi/jsonpointer v0.19.5 // indirect
	github.com/go-openapi/jsonreference v0.19.5 // indirect
	github.com/go-openapi/swag v0.19.14 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/gnostic v0.6.9 // indirect
	github.com/google/go-cmp v0.5.8 // indirect
	github.com/google/gofuzz v1.2.0 // indirect
	github.com/imdario/mergo v0.3.13 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/mailru/easyjson v0.7.6 // indirect
	github.com/mattn/go-runewidth v0.0.13 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/rivo/uniseg v0.2.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/stolostron/cluster-lifecycle-api v0.0.0-20220714081119-eae2fe1f05fd // indirect
	golang.org/x/net v0.0.0-20220708220712-1185a9018129 // indirect
	golang.org/x/oauth2 v0.0.0-20220718184931-c8730f7fcb92 // indirect
	golang.org/x/sys v0.0.0-20220715151400-c0bba94af5f8 // indirect
	golang.org/x/term v0.0.0-20220526004731-065cf7ba2467 // indirect
	golang.org/x/text v0.3.7 // indirect
	golang.org/x/time v0.0.0-20220609170525-579cf78fd858 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/protobuf v1.28.0 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	k8s.io/api v0.24.3 // indirect
	k8s.io/klog/v2 v2.70.1 // indirect
	k8s.io/kube-openapi v0.0.0-20220627174259-011e075b9cb8 // indirect
	k8s.io/utils v0.0.0-20220713171938-56c0de1e6f5e // indirect
	sigs.k8s.io/controller-runtime v0.12.3 // indirect
	sigs.k8s.io/json v0.0.0-20211208200746-9f7c6b3444d2 // indirect
	sigs.k8s.io/structured-merge-diff/v4 v4.2.1 // indirect
	sigs.k8s.io/yaml v1.3.0 // indirect
)

replace (
	github.com/buger/jsonparser => github.com/buger/jsonparser v1.1.1
	github.com/coreos/go-systemd => github.com/coreos/go-systemd/v22 v22.1.0
	github.com/docker/docker => github.com/docker/docker v1.13.1
	github.com/go-logr/logr => github.com/go-logr/logr v1.2.3
	github.com/hashicorp/go-getter => github.com/hashicorp/go-getter v1.5.11
	github.com/hashicorp/terraform => github.com/openshift/terraform v0.12.20-openshift-4
	github.com/hashicorp/terraform-plugin-sdk => github.com/openshift/hashicorp-terraform-plugin-sdk v1.14.0-openshift
	github.com/hashicorp/terraform@v0.13.4 => github.com/openshift/terraform v0.12.20-openshift-4
	github.com/kubevirt/terraform-provider-kubevirt => github.com/kubevirt/terraform-provider-kubevirt v0.0.0-20210628085519-5c4934a8bda8
	github.com/metal3-io/baremetal-operator => github.com/metal3-io/baremetal-operator v0.0.0-20211220105604-05d12b6768a9
	github.com/metal3-io/baremetal-operator/apis => github.com/metal3-io/baremetal-operator/apis v0.0.0-20220119160837-9b26aa7816ca
	github.com/metal3-io/baremetal-operator/pkg/hardwareutils => github.com/metal3-io/baremetal-operator/pkg/hardwareutils v0.0.0-20220119160837-9b26aa7816ca //indirect
	github.com/metal3-io/cluster-api-provider-baremetal => github.com/openshift/cluster-api-provider-baremetal v0.0.0-20220104154407-0716ee4cb33b
	github.com/openshift/client-go => github.com/openshift/client-go v3.9.0+incompatible
	github.com/openshift/hive/apis => github.com/openshift/hive/apis v0.0.0-20220121012553-a0671aa97ef3
	github.com/terraform-providers/terraform-provider-azurerm => github.com/terraform-providers/terraform-provider-azurerm v0.0.0-20200604143437-d38893bc4f78
	github.com/terraform-providers/terraform-provider-ignition/v2 => github.com/community-terraform-providers/terraform-provider-ignition/v2 v2.1.0
	golang.org/x/crypto => golang.org/x/crypto v0.0.0-20220622213112-05595931fe9d
	k8s.io/api => k8s.io/api v0.24.3
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.24.3
	k8s.io/apimachinery => k8s.io/apimachinery v0.24.3
	k8s.io/apiserver => k8s.io/apiserver v0.24.3
	k8s.io/cli-runtime => k8s.io/cli-runtime v0.24.3
	k8s.io/client-go => k8s.io/client-go v0.24.3
	k8s.io/cloud-provider => k8s.io/cloud-provider v0.24.3
	k8s.io/cloud-provides => k8s.io/cloud-provides v0.24.3
	k8s.io/cluster-bootstrap => k8s.io/cluster-bootstrap v0.24.3
	k8s.io/code-generator => k8s.io/code-generator v0.24.3
	k8s.io/component-base => k8s.io/component-base v0.24.3
	k8s.io/cri-api => k8s.io/cri-api v0.24.3
	k8s.io/csi-translation-lib => k8s.io/csi-translation-lib v0.24.3
	k8s.io/kube-aggregator => k8s.io/kube-aggregator v0.24.3
	k8s.io/kube-controller-manager => k8s.io/kube-controller-manager v0.24.3
	k8s.io/kube-proxy => k8s.io/kube-proxy v0.24.3
	k8s.io/kube-scheduler => k8s.io/kube-scheduler v0.24.3
	k8s.io/kubectl => k8s.io/kubectl v0.24.3
	k8s.io/kubelet => k8s.io/kubelet v0.24.3
	k8s.io/legacy-cloud-providers => k8s.io/legacy-cloud-providers v0.24.3
	k8s.io/metrics => k8s.io/metrics v0.24.3
	k8s.io/sample-apiserver => k8s.io/sample-apiserver v0.24.3
	kubevirt.io/client-go => kubevirt.io/client-go v0.55.0
	sigs.k8s.io/cluster-api-provider-aws => sigs.k8s.io/cluster-api-provider-aws v0.4.0
	sigs.k8s.io/cluster-api-provider-azure => sigs.k8s.io/cluster-api-provider-azure v0.4.0
	sigs.k8s.io/cluster-api-provider-openstack => sigs.k8s.io/cluster-api-provider-openstack v0.3.0
	sigs.k8s.io/yaml => sigs.k8s.io/yaml v1.3.0 // indirect
)
