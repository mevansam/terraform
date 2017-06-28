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
* `hostname` - (Optional) The hostname which will be used to create a default route to the app on the default application domain. If this argument is not provided a route can be created and mapped to this application using the [`cf_route`](/docs/providers/cf/r/route.html). 
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

* `url` - (Optional, String) The URL for the application binary. A local path may be referenced via "`file://...`".

* `git` - (Optional, String) The git location to pull the application source directly from source control.

  - `url` - (Required, String) The git URL for the application repository.
  - `branch` - (Optional, String) The branch or tag of the repository.
  - `key` - (Optional, String) The git private key to access a private repo via SSH.
  - `user` - (Optional, String) Git user for accessing a private repo.
  - `password` - (Optional, String) Git password for accessing a private repo.

* `github_release` - (Optional, String) The Buildpack archive published as a github release.

  - `owner` - (Required, String) The github owner or organization name
  - `repo` - (Required, String) The repository containing the release
  - `token` - (Optional, String) Github API token to use to access Github
  - `version` - (Optional, String) The version or tag of the release.
  - `filename` - (Required, String) The name of the published file. The values `zipball` or `tarball` will download the published    
  
### Service bindings

Modifying this argument will cause the application to be restaged.

* `service_binding` - (Optional, Array) Service instances to bind to.

  - `service` - (Required, String) The service instance GUID.
  - `params` - (Optional, Map) A list of key/value parameters used by the service broker to create the binding.

### Blue-Green Deployment Strategy

* `blue_green` - (Optional) Defines a blue-green app deployment strategy. When doing a blue-green deployment the actual name of the application will be timestamped to differentiate between the current live application and the more recent staged application.

  - `stage_route` - (Required, String) The GUID of the route where the staged application will be available.
  - `live_route` - (Required, String) The GUID of the route where the live application will be available.
  - `validation_script` - (Optional, String) The validation script to execute against the stage application before mapping the live route to the staged application.

### Environment Variables

Modifying this argument will cause the application to be restaged.

* `environment` - (Optional, Map) Key/value pairs of all the environment variables to run in your app. Does not include any system or service variables.

### Health Checks

* `health-check-http-endpoint` -(Optional, String) The endpoint for the http health check type. The default is '/'.
* `health-check-type` - (Optional, String) The health check type which can be on of "`port`", "`process`", "`http`" or "`none`".

## Attributes Reference

The following attributes are exported along with any defaults for the inputs attributes.

* `id` - The GUID of the application

