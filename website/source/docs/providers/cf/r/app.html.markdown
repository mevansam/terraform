---
layout: "cf"
page_title: "Cloud Foundry: cf_app"
sidebar_current: "docs-cf-resource-app"
description: |-
  Provides a Cloud Foundry Application resource.
---

# cf\_app

Provides a Cloud Foundry application resource for managing Cloud Foundry [applications](https://docs.cloudfoundry.org/devguide/deploy-apps/deploy-app.html).

## Example Usage

The following example creates an application.

```
resource "cf_app" "spring-music" {
    name = "spring-music"
    url = "file:///Work/cloudfoundry/apps/spring-music/build/libs/spring-music.war"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the application in Cloud Foundry space.
* `space` - (Required) The GUID of the associated space.
* `instances` - (Optional, Number) The number of app instances that you want to start.
* `memory` - (Optional, String) The memory limit for all instances of an app. This attribute requires a unit of measurement: M, MB, G, or GB, in upper case or lower case.
* `disk_quota` - (Optional, String) The disk space for the app instance. This attribute requires a unit of measurement: M, MB, G, or GB, in upper case or lower case.
* `stack` - (Optional) The GUID of the stack the application will be deployed to. Use the [`cf_stack`](/docs/providers/cf/d/stack.html) data resource to lookup the stack GUID to overriding the default.
* `state` - (Optional, String) The desired state of the app. One of "`STOPPED`" or "`STARTED`".
* `command` - (Optional, String) A custom start command for the application. This overrides the start command provided by the buildpack.
* `buildpack` - (Optional, String) The custom buildpack to use. This will bypass the buildpack detect phase.
* `enable_ssh` - (Optional, Boolean) Whether to enable or disable SSH access to the container. Default is `true` unless disabled globally.
* `timeout` - (Optional, Number) Defines the number of seconds that Cloud Foundry waits for starting your application.

### Application Source / Binary

One of the following arguments must be declared to locate application source or archive to be pushed.

* `url` - (Optional) The URL for the application binary. A local path may be referenced via "`file://...`".
* `git` - (Optional) The git location to pull the application source directly from source control.
  - `url` - (Required) The git URL for the application repository.
  - `branch` - (Optional) The branch or tag of the repository.
  - `key` - (Optional) The git private key to access a private repo via SSH.
  - `user` - (Optional) Git user for accessing a private repo.
  - `password` - (Optional) Git password for accessing a private repo.

### Service bindings

* `service_binding` - (Optional, Array) Service instances to bind to.

  - `service` - (Required, String) The service instance GUID.
  - `params` - (Optional, Map) A list of key/value parameters used by the service broker to create the binding.

### Application Routes

* `route` - (Optional, Array) Application routes to associate with the application. For applications that serve requests via HTTP or TCP at least one route needs to be declared in order for the application to be externally accessible.

  - `hostname` - (Optional, String) The application's host name.
  - `domains` - (Optional, Array of strings) Array of one of more domain GUIDs the application routes will be created on.
  - `port` - (Optional) The port to associate with the route for a TCP route. 
  - `path` - (Optional) A path for a HTTP route.

### Environment Variables

* `environment` - (Optional, Map) Key/value pairs of all the environment variables to run in your app. Does not include any system or service variables.

### Health Checks

* `health-check-http-endpoint` -(Optional, String) The endpoint for the http health check type. The default is '/'.
* `health-check-type` - (Optional, String) The health check type which can be on of "`port`", "`process`", "`http`" or "`none`".


## Attributes Reference

The following attributes are exported along with any defaults for the inputs attributes.

* `id` - The GUID of the application

