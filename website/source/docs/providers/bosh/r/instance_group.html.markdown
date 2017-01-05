---
layout: "bosh"
page_title: "Bosh: bosh_instance_group"
sidebar_current: "docs-bosh-resource-deployment-instance-group"
description: |-
  Provides a Bosh Instance Group resource.
---

# bosh\_instance\_group

Provides a Bosh [Instance Group](http://bosh.io/docs/manifest-v2.html#instance-groups) resource. An instance group declares the mapping between release jobs and VM instances. 

## Example Usage

The folowing declares an instance group for deploying services required by a Bosh Director environment. 


```
# Set the variable values in *.tfvars file
# or using -var="..." CLI option

variable "bosh_public_ip" {}
variable "bosh_services_password" {}

# Configure the instance group for the director services

resource "bosh_instance_group" "bosh" {
    
    name = "bosh"
    instances = 1
    vm_type = "${bosh_vm_type.large.id}"
    disk_type = "${bosh_disk_type.small.id}"
    stemcell = "${bosh_stemcell.ubuntu.id}"

    network {
        id = "${bosh_manual_network.private.id}
        static_ips = [ "${cidrhost(var.subnet_cidr, 254)}" ]
        default = [ "dns", "gateway" ]
    }
    network {
        id = "${bosh_vip_network.public.id}
        static_ips = [ "${var.bosh_public_ip}" ]
    }

    jobs = [
        "${bosh_job.nats.id}",
        "${bosh_job.postgres.id}",
        "${bosh_job.registry.id}",
        "${bosh_job.director.id}",
        "${bosh_job.blobstore.id}",
        "${bosh_job.health_monitor.id}"
    ]

    properties = <<EOF
---
nats:
    address: 127.0.0.1
    user: nats
    password: ${var.bosh_services_password}

postgres:
    listen_address: 127.0.0.1
    host: 127.0.0.1
    user: postgres
    password: ${var.bosh_services_password}
    database: bosh
    adapter: postgres

registry:
    address: ${var.bosh_public_ip}
    host: ${var.bosh_public_ip}
    port: 25777
    username: admin
    password: ${var.bosh_services_password}
    db: ${bosh_job.postgres.id}
    http:
        user: admin
        password: ${var.bosh_services_password}
        port: 25777

director:
    address: 127.0.0.1
    name: ${var.bosh_name}
    db: ${bosh_job.postgres.id}
    cpi_job: openstack_cpi
    max_threads: 3
    user_management:
        provider: local
        local:
            users:
            - { name: ${var.bosh_user}, password: ${var.bosh_password} }
            - { name: hm, password: ${var.bosh_services_password} }

blobstore:
    address: ${var.bosh_public_ip}
    port: 25250
    provider: dav
    director:
        user: director
        password: ${var.bosh_services_password}
    agent:
        user: agent
        password: ${var.bosh_services_password}

hm:
    resurrector_enabled: true
    director_account:
        user: hm
        password: ${var.bosh_services_password}

agent:
    mbus: nats://nats:passw0rd@${var.bosh_public_ip}:4222

ntp:
- 0.pool.ntp.org
- 1.pool.ntp.org

EOF
}
```

This example references the following resources:

* [bosh_vm_type](/docs/providers/bosh/r/vm_type.html)
* [bosh_disk_type](/docs/providers/bosh/r/disk_type.html)
* [bosh_stemcell](/docs/providers/bosh/r/stemcell.html)
* [bosh_manual_network](/docs/providers/bosh/r/manual_network.html)
* [bosh_vip_network](/docs/providers/bosh/r/vip_network.html)
* [bosh_job](/docs/providers/bosh/r/job.html)

## Argument Reference

The following arguments are supported:

* `name` - (String, required) A unique name used to identify and reference instance group.

* `azs` - (Array, optional) List of AZs associated with this instance group (should only be used when using first class AZs). Example: [z1, z2].

* `instances` - (Integer, required) The number of instances in this group. Each instance is a VM.

* `jobs` - (Array, required) The IDs of the [job](/docs/providers/bosh/r/job.html) resources to apply to this instance group

* `vm_type` - (String, required) The ID of a [bosh_vm_type](/docs/providers/bosh/r/vm_type.html) resource.

* `vm_extensions` - (Array, optional) An array of IDs of [bosh_vm_extension](/docs/providers/bosh/r/vm_extension.html) resources.

* `stemcell` - (String, required) The ID of a [bosh_stemcell](/docs/providers/bosh/r/stemcell.html) resource.

* `persistent_disk_type` - (String, optional) A valid disk type name from the cloud config. Read more about persistent disks

* `network` - (Array, required) Specifies the networks this instance requires. Each network can have the following properties specified:

  - `id` - (String, required) A [bosh_manual_network](/docs/providers/bosh/r/manual_network.html), [bosh_dynamic_network](/docs/providers/bosh/r/dynamic_network.html) or [bosh_vip_network](/docs/providers/bosh/r/vip_network.html) resource id.
  - `static_ips` - (Array, optional) Array of IP addresses reserved for the instances on the network.
  - `default` - (Array, optional) Specifies which network components (DNS, Gateway) BOSH populates by default from this network. This property is required if more than one network is specified.

* `update` - (Hash, optional) Specific update settings for this instance group. Use this to override global job update settings on a per-instance-group basis.

* `migrated_from` - (Array, optional) Specific migration settings for this instance group. Use this to rename and/or migrate instance groups.

* `lifecycle` - (String, optional) Specifies the kind of workload the instance group represents. Valid values are service and errand; defaults to service. A service runs indefinitely and restarts if it fails. An errand starts with a manual trigger and does not restart if it fails.

* `properties` - (String, optional) Specifies instance group property collection as a YAML string. You can use Terraform interpolation syntax to interpolate attributes and variables declared elsewhere in the Terraform configuration. If a resource ID is provided in the value field of a YAML key, then that resource's exported properties will be inserted.

## Attributes Reference

The following attributes are exported:

* `id` - A GUID identifying the instance group
