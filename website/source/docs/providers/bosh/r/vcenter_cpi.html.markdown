---
layout: "bosh"
page_title: "Bosh: bosh_vcenter_cpi"
sidebar_current: "docs-bosh-resource-cpi-vcenter"
description: |-
  Provides a Bosh vCenter Cloud Provider Interface resource.
---

# bosh\_vcenter\_cpi

Provides a Bosh vCenter [Cloud Provider Interface](http://bosh.io/docs) resource which is an abstraction of the underlying IaaS. It can be used to manage Bosh deployments on a vCenter environment.

## Example Usage

```
resource "bosh_vcenter_cpi" "vcenter" {    
    
    ...
}
```

## Argument Reference

The following arguments are supported:

* `address` - (required) vCenter IP
* `user` - (required) vCenter user for managing IaaS resources via the API
* `password` - (required) The user's password
* `datacenter` - (required) An array of logical vCenter data centers to deploy to. At least one data center must be declared
  - `name` - (required) The name of the vCenter data center
  - `clusters` - (required) An array of cluster to deploy resources to
  - `vm_folder` - (required) The logical folder within which Bosh managed VMs will be organized
  - `template_folder` - (required) The logical folder within which Bosh stemcells will be organized
  - `datastore_pattern` - (required) Regex to match datastores to use for Bosh VM deployments
  - `persistent_datastore_pattern` - (required) Regex to match datastores to use for deploying persistent disks
  - `disk_path` - (required) The logical folder within which Bosh managed disks will be organized

## Attributes Reference

The following attributes are exported:

* `id` - A GUID identifying the CPI
