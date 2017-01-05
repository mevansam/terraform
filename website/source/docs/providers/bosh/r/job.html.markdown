---
layout: "bosh"
page_title: "Bosh: bosh_job"
sidebar_current: "docs-bosh-resource-deployment-job"
description: |-
  Provides a Bosh Job resource.
---

# bosh\_job

Provides a Bosh Job resource which will be applied to an [Instance Group](http://bosh.io/docs/manifest-v2.html#instance-groups). This resource refers to a particular release [job](http://bosh.io/docs/jobs.html) and the resource should declare the configuration parameters to apply when the job is installed to an instance group.

## Example Usage

The folowing declares jobs required to deploy a Bosh environment. 

```
resource "bosh_job" "nats" {   
    name = "nats"
    release = "${bosh_release.bosh.id}"
}
resource "bosh_job" "postgres" {   
    name = "postgres"
    release = "${bosh_release.bosh.id}"
}
resource "bosh_job" "registry" {   
    name = "registry"
    release = "${bosh_release.bosh.id}"
}
resource "bosh_job" "director" {   
    name = "director"
    release = "${bosh_release.bosh.id}"
}
resource "bosh_job" "blobstore" {   
    name = "blobstore"
    release = "${bosh_release.bosh.id}"
}
resource "bosh_job" "health_monitor" {   
    name = "health_monitor"
    release = "${bosh_release.bosh.id}"
}
```

This example references the following resources:

* [bosh_release](/docs/providers/bosh/r/release.html)

## Argument Reference

The following arguments are supported:

* `name` - (String, required) The job name defined in the release
* `release` - (String, required) The ID of a [bosh_release](/docs/providers/bosh/r/release.html) resource that provides a Bosh release specification.
* `consumes` - (Map, optional) Links consumes by the job.
* `provides` - (Map, optional) Links provided by the job.
* `properties` - (String, optional) Specifies a job property collection as a YAML string. You can use Terraform interpolation syntax to interpolate attributes and variables declared elsewhere in the Terraform configuration. If a resource ID is provided in the value field of a YAML key, then that resource's exported properties will be inserted.

## Attributes Reference

The following attributes are exported:

* `id` - A GUID identifying the job
