---
layout: "bosh"
page_title: "Bosh: bosh_env"
sidebar_current: "docs-bosh-resource-infrastructure-cloud-config"
description: |-
  Provides a Bosh Director Cloud Config resource.
---

# bosh\_cloud\_config

Provides a Bosh [Cloud Config](https://bosh.io/docs/cloud-config.html) resource.

## Example Usage

The following configures a Bosh director with the CPI specification 
declared by the Bosh provider.

```
resource "bosh_cloud_config" "mycloud" {
    
    name = "mycloud"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of this Cloud Config

## Attributes Reference

The following attributes are exported:

* `id` - A GUID identifying the cloud config