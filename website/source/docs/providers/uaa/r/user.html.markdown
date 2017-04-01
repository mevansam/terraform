---
layout: "uaa"
page_title: "UAA: uaa_user"
sidebar_current: "docs-uaa-resource-user"
description: |-
  Provides a UAA user resource.
---

# uaa\_user

Provides a UAA resource for managing users.

## Example Usage

The following example creates a user and attaches additional UAA roles to grant administrator rights to that user.

```
resource "uaa_user" "admin-service-user" {
    name = "cf-admin"
    password = "Passw0rd"
    given_name = "John"
    family_name = "Doe"
    groups = [ "cloud_controller.admin", "scim.read", "scim.write" ]
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the user. This will also be the users login name
* `password` - (Optional) The user's password
* `origin` - (Optional) The user authentcation origin. By default this will be `UAA`. For users authenticated by LDAP this should be `ldap`
* `given_name` - (Optional) The given name of the user
* `family_name` - (Optional) The family name of the user
* `email` - (Optional) The email address of the user
* `groups` - (Optional) Any UAA `groups` / `scopes` to associated the user with

## Attributes Reference

The following attributes are exported:

* `id` - The GUID of the User
* `email` - If not provided this attribute will be assigned the same value as the `name`, assuming that the username is the user's email address
