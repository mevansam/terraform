---
layout: "bosh"
page_title: "Bosh: bosh_stemcell"
sidebar_current: "docs-bosh-resource-deployment-stemcell"
description: |-
  Provides a Bosh Stemcell resource.
---

# bosh\_stemcell

Provides a Bosh [Stemcell](http://bosh.io/docs/stemcell.html) resource.

## Example Usage

The following declares a Bosh stemcell for an Ubuntu VM image.

```
# Set the variable values in *.tfvars file
# or using -var="..." CLI option

variable "stemcell_version" {}
variable "stemcell_sha1" {}

# Configure the Bosh Release

resource "bosh_stemcell" "ubuntu" {
    
    alias = "trusty"
    os = "ubuntu"
    version = "${var.stemcell_version}"

    url = "https://bosh.io/d/stemcells/bosh-openstack-kvm-ubuntu-trusty-go_agent?v=${self.version}"
    sha1 = "${var.stemcell_sha1}"
}
```

## Argument Reference

The following arguments are supported:

* `alias` - (Required) The name of the Bosh stemcell
* `os` - (Optional) The Operating System of the stemcell
* `version` - (Optional) The version of the stemcell. If this value is omitted then the 'latest' stemcell in the blobstore will be used.
* `url` - (Optional) The URL from which the stemcell can be download. A local filesystem path can be specified via "file://" url protocol.
* `sha1` - (Optional) The SHA1 signature of the stemcell archive

## Attributes Reference

The following attributes are exported:

* `id` - A GUID identifying the stemcell