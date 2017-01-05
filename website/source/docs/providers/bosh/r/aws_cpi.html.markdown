---
layout: "bosh"
page_title: "Bosh: bosh_aws_cpi"
sidebar_current: "docs-bosh-resource-cpi-aws"
description: |-
  Provides a Bosh AWS Cloud Provider Interface resource.
---

# bosh\_aws\_cpi

Provides a Bosh AWS [Cloud Provider Interface](http://bosh.io/docs) resource which is an abstraction of the underlying IaaS. It can be used to manage Bosh deployments on an AWS IaaS environment.

## Example Usage

The following declares an AWS CPI.

```
# Set the variable values in *.tfvars file
# or using -var="..." CLI option

variable "aws_cpi_release_url" {}
variable "aws_cpi_release_sha1" {}

variable "aws_access_key_id" {}
variable "aws_secret_access_key" {}

# Configure the Bosh OpenStack CPI

resource "bosh_aws_cpi" "us-west-1" {
    
    cpi_release_url = "${var.aws_cpi_release_url}"
    cpi_release_sha1 = "${var.aws_cpi_release_sha1}"
        
    access_key_id = "${var.aws_access_key_id}"
    secret_access_key = "${var.aws_secret_access_key}"
    default_key_name = "default"
    default_security_groups = [ "default" ]
    region = "us-west-1"
}
```

## Argument Reference

The following arguments are supported:

* `cpi_release_url` - (required) The Bosh release URL for the CPI. A local filesystem path can be specified via "file://" url protocol
* `cpi_release_sha1` - (required) The SHA1 signature of the release archive
* `access_key_id` - (required) AWS Access Key ID
* `secret_access_key` - (required) AWS Secret Key
* `default_key_name` - (required) Name of key pair to use for managed instances
* `default_security_groups` - (required) Array of security groups
* `region` - (required) AWS region

## Attributes Reference

The following attributes are exported:

* `id` - A GUID identifying the CPI
