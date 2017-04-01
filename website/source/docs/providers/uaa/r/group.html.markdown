---
layout: "uaa"
page_title: "UAA: uaa_group"
sidebar_current: "docs-uaa-resource-group"
description: |-
  Provides a UAA group resource.
---

# uaa\_group

Provides a UAA resource for managing groups.

## Example Usage

The following example creates a group.

```
resource "uaa_group" "mygroup" {
    name = "my.custom.group"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the UAA `group` / `scope`.

## Attributes Reference

The following attributes are exported:

* `id` - The GUID of the User
* `email` - If not provided this attribute will be assigned the same value as the `name`, assuming that the username is the user's email address
