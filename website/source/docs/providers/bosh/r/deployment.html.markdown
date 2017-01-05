---
layout: "bosh"
page_title: "Bosh: bosh_deployment"
sidebar_current: "docs-bosh-resource-deployment-deployment"
description: |-
  Provides a Bosh Director deployment resource.
---

# bosh\_deployment

Provides a Bosh Director deployment resource. This resource builds and applies the Bosh [deployment manifest](https://bosh.io/docs/manifest-v2.html).

## Example Usage

The following configures a Bosh director environment configured to orchestrate deployments on the IaaS configured.

```
resource "bosh_deployment" "bosh" {
    
    name = "bosh"
    target = "${bosh_openstack_cpi.us-west-datacenter.id}"

    instance_groups = [ 
        "${var.bosh_instance_group.bosh.id}"
    ]

    update {
      canaries = 1
      max_in_flight = 10
      canary_watch_time = "1000-30000"
      update_watch_time = "1000-30000"
    }
}
```

This example references the following resources:

* [bosh_openstack_cpi](/docs/providers/bosh/r/openstack_cpi.html)
* [bosh_release](/docs/providers/bosh/r/release.html)
* [bosh_stemcell](/docs/providers/bosh/r/stemcell.html)
* [bosh_manual_network](/docs/providers/bosh/r/manual_network.html)
* [bosh_vip_network](/docs/providers/bosh/r/vip_network.html)
* [bosh_job](/docs/providers/bosh/r/job.html)

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the deployment

* `target` - (Required) The ID of a CPI or Environment resource. The following resources can be referenced.

  - [bosh_aws_cpi](/docs/providers/bosh/r/aws_cpi.html)
  - [bosh_azure_cpi](/docs/providers/bosh/r/azure_cpi.html)
  - [bosh_openstack_cpi](/docs/providers/bosh/r/openstack_cpi.html)
  - [bosh_vcenter_cpi](/docs/providers/bosh/r/vcenter_cpi.html)
  - [bosh_env](/docs/providers/bosh/d/env.html)

* `instance_groups` - (Required) Array of IDs of [bosh_instance_group](/docs/providers/bosh/r/instance_group.html) resources.

* `update` - (Required) This specifies instance update properties. These properties control how BOSH updates instances during the deployment.

  - `canaries` - (Integer, required) The number of canary instances.
  - `max_in_flight` - (Integer, required) The maximum number of non-canary instances to update in parallel.
  - `canary_watch_time` - (String, required) Only applies to monit start operation.
  - `update_watch_time` - (String, required) Only applies to monit start operation.
  - `serial` - (Boolean, optional) If disabled (set to false), instance groups will be deployed in parallel, otherwise - sequentially. 

* `properties` - (String, optional) Specifies a global property collection as a YAML string. You can use Terraform interpolation syntax to interpolate attributes and variables declared elsewhere in the Terraform configuration. If a resource ID is provided in the value field of a YAML key, then that resource's exported properties will be inserted.

* `tags` - (Map, optional) Specifies key value pairs to be sent to the CPI for VM tagging. Combined with runtime config level tags during the deploy. Available in bosh-release v258+.

## Attributes Reference

The following attributes are exported:

* `id` - A GUID identifying the deployment
* `instance_ips` - A map of the deployed instance groups and the IPs of each group's instances.
