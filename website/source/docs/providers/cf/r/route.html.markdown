---
layout: "cf"
page_title: "Cloud Foundry: cf_route"
sidebar_current: "docs-cf-resource-route"
description: |-
  Provides a Cloud Foundry route resource.
---

# cf\_route

Provides a Cloud Foundry resource for managing Cloud Foundry application [routes](https://docs.cloudfoundry.org/devguide/deploy-apps/routes-domains.html).

## Example Usage

The following example creates an route for an application.

```
resource "cf_route" "default" {
    domain = "${data.cf_domain.apps.domain.id}"
    space = "${data.cf_space.dev.id}"
    hostname = "myapp"
}
```

## Argument Reference

The following arguments are supported:

- `domain` - (Required, String) The ID of the domain to map the host name to. If not provided the default application domain will be used.
- `space` - (Required, String) The ID of the space to create the route in.
- `hostname` - (Required, Optional) The application's host name. This is required for shared domains.
- `application` - (Optional, String) The ID of the application to map this route to.

The following argument applies only to HTTP routes.

- `port` - (Optional) The port to associate with the route for a TCP route. 

The following argument applies only to TCP routes.

- `path` - (Optional) A path for a HTTP route.

## Attributes Reference

The following attributes are exported along with any defaults for the inputs attributes.

* `id` - The GUID of the route
