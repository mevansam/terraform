---
layout: "bosh"
page_title: "Bosh: bosh_env"
sidebar_current: "docs-bosh-resource-infrastructure-vm-extension"
description: |-
  Provides a Bosh Director environment resource.
---

# bosh\_env

Provides a Bosh Director [VM Extensions](https://bosh.io/docs/cloud-config.html#vm-extensions) resource. 

## Example Usage

```
resource "vm_extension" "pub-lbs" {
    
    name = "pub-lbs"

    properties = <<EOF
---
elbs: [main]
EOF
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of VM extension.

* `properties` - (String, optional) Describes any IaaS-specific properties needed to configure VMs, as a YAML string. You can use Terraform interpolation syntax to interpolate attributes and variables declared elsewhere in the Terraform configuration. 

## Attributes Reference

The following attributes are exported:

* `id` - The GUID of the Space