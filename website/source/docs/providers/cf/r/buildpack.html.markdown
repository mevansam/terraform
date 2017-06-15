---
layout: "cf"
page_title: "Cloud Foundry: cf_buildpack"
sidebar_current: "docs-cf-resource-buildpack"
description: |-
  Provides a Cloud Foundry Buildpack resource.
---

# cf\_buildpack

Provides a Cloud Foundry resource for managing Cloud Foundry [buildpacks](https://docs.cloudfoundry.org/adminguide/buildpacks.html).

## Example Usage

The following example creates a Cloud Foundry Buildpack .

```
resource "cf_buildpack" "tomee" {
    name = "tomcat-enterprise-edition"
    path = "https://github.com/cloudfoundry-community/tomee-buildpack"
    position = "12"
    enable = true
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the Buildpack.
* `position` - (Optional, Number) Specifies where to place the buildpack in the detection priority list. For more information, see the [Buildpack Detection](https://docs.cloudfoundry.org/buildpacks/detection.html) topic.
* `enabled` - (Optional, Boolean) Specifies whether to allow apps to be pushed with the buildpack, and defaults to true.
* `locked` - (Optional, Boolean) Specifies whether buildpack is locked to prevent further updates, and defaults to false.

### Buildpack location

One of the following arguments must be declared to locate buildpack source or archive to be uploaded.

* `url` - (Optional) Specifies the location of the buildpack to upload from. It can be a URL to a zip file, a github repository or a local directory via "`file://...`".
* `git` - (Optional) The git location to pull the builpack source directly from source control.
  - `url` - (Required) The git URL for the application repository.
  - `branch` - (Optional) The branch or tag of the repository.
  - `key` - (Optional) The git private key to access a private repo via SSH.
  - `user` - (Optional) Git user for accessing a private repo.
  - `password` - (Optional) Git password for accessing a private repo.

## Attributes Reference

The following attributes are exported:

* `id` - The GUID of the organization
