---
layout: "uaa"
page_title: "UAA: uaa_client"
sidebar_current: "docs-uaa-resource-client"
description: |-
  Provides a UAA client resource.
---

# uaa\_client

Provides a UAA resource for managing OAuth2 clients.

## Example Usage

The following example creates a client.

```
resource "uaa_client" "myclient" {
    name = "myclient"
    scope  = [<list>]

    authorized_grant_types = [<list>]
    authorities = [<list>]
    autoapprove = [<list>]

    redirect_uri = [<list>]
    signup_redirect_url = "<url>"

    access_token_validity = <seconds>
    refresh_token_validity = <seconds>
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the client.
* `scope` - 
* `authorized_grant_types` - 
* `authorities` - 
* `autoapprove` - 
* `redirect_uri` - 
* `signup_redirect_url` - 
* `access_token_validity` - 
* `access_token_validity` - 

## Attributes Reference

The following attributes are exported:

* `id` - The GUID of the User
