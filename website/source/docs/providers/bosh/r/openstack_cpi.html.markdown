---
layout: "bosh"
page_title: "Bosh: bosh_openstack_cpi"
sidebar_current: "docs-bosh-resource-cpi-openstack"
description: |-
  Provides a Bosh OpenStack Cloud Provider Interface resource.
---

# bosh\_openstack\_cpi

Provides a Bosh OpenStack [Cloud Provider Interface](http://bosh.io/docs) resource which is an abstraction of the underlying IaaS. It can be used to manage Bosh deployments on an OpenStack IaaS environment.

## Example Usage

The following declares an OpenStack CPI.

```
# Set the variable values in *.tfvars file
# or using -var="..." CLI option

variable "openstack_cpi_release_url" {}
variable "openstack_cpi_release_sha1" {}

variable "openstack_auth_url" {}
variable "openstack_project" {}
variable "openstack_domain" {}
variable "openstack_username" {}
variable "openstack_api_key" {}

# Configure the Bosh OpenStack CPI

resource "bosh_openstack_cpi" "us-west-datacenter" {
    
    cpi_release_url = "${var.openstack_cpi_release_url}"
    cpi_release_sha1 = "${var.openstack_cpi_release_sha1}"
        
    auth_url = "${var.openstack_auth_url}"
    project = "${var.openstack_project}"
    domain = "${var.openstack_domain}"
    username = "${var.openstack_username}"
    api_key = "${var.openstack_api_key}"
    region = "us-west-datacenter"
    default_key_name = "default"
    default_security_groups = [ "default" ]
}
```

## Argument Reference

The following arguments are supported:

* `cpi_release_url` - (required) The Bosh release URL for the CPI. A local filesystem path can be specified via "file://" url protocol
* `cpi_release_sha1` - (required) The SHA1 signature of the release archive
* `auth_url` - (required) The OpenStack Identity API endpoint
* `project` - (required) The OpenStack project name
* `domain` - (required) The OpenStack domain name
* `region` - (optional) The OpenStack region
* `username` - (required) The API user name
* `api_key` - (required) The API user's password
* `default_key_name` - (required) Name of the key-pair used when deploying instances
* `default_security_groups` - (required) Array of security groups to attach to deployed instance

## Attributes Reference

The following attributes are exported:

* `id` - A GUID identifying the CPI

