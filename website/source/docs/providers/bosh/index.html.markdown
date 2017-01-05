---
layout: "bosh"
page_title: "Provider: Bosh"
sidebar_current: "docs-bosh-index"
description: |-
  The Bosh provider is used to provision a Bosh Director environment.
---

# Bosh Provider

The Bosh provider declares the Bosh Director environment which is responsible 
for orchestrating and managing Bosh deployments. 

Use the navigation to the left to read about the available resources.

## Example Usage

```
provider "bosh" {
    skip_ssl_validation = true
}
```

## Argument Reference

The following arguments are supported:

* `skip_ssl_validation` - (optional) Turns of SSL certificate verification. 
This can also be specified with the `BOSH_SKIP_SSL_VALIDATION` shell environment 
variable.

* `ca_cert` - (optional) The CA certificate for verifying the Bosh SSL endpoint. 
This can also be specified with the `CA_CERT` shell environment variable.

* `ca_cert_file` - (optional) Path to the file containing the Bosh director's CA 
certificate. This can also be specified with the `CA_CERT_FILE` shell environment 
variable.
