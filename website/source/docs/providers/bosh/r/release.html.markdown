---
layout: "bosh"
page_title: "Bosh: bosh_release"
sidebar_current: "docs-bosh-resource-deployment-release"
description: |-
  Provides a Bosh Release resource.
---

# bosh\_release

Provides a Bosh [Release](http://bosh.io/docs/release.html) resource.

## Example Usage

The following declares a Bosh release containing the Bosh Director service bits.

```
# Set the variable values in *.tfvars file
# or using -var="..." CLI option

variable "bosh_release_version" {}
variable "bosh_release_sha1" {}

# Configure the Bosh Release

resource "bosh_release" "bosh" {
    
    name = "bosh"
    version = "${var.bosh_release_version}"

    url = "https://bosh.io/d/github.com/cloudfoundry/bosh?v=${self.version}"
    sha1 = "${var.bosh_release_sha1}"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the Bosh release
* `version` - (Optional) The version of the release. If this value is omitted then the 'latest' release in the blobstore will be used.
* `url` - (Optional) The URL from which the release can be download. A local filesystem path can be specified via "file://" url protocol.
* `sha1` - (Optional) The SHA1 signature of the release archive

## Attributes Reference

The following attributes are exported:

* `id` - A GUID identifying the release