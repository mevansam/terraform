Terraform
=========

# Overview

[![wercker status](https://app.wercker.com/status/7bcf32b26184e3d19df3e12198666836/m/master "wercker status")](https://app.wercker.com/project/byKey/7bcf32b26184e3d19df3e12198666836)

This repository contains a custom build of the latest [Terraform](https://www.terraform.io) source including the following providers.

* [Cloud Foundry](https://github.com/mevansam/terraform/tree/ab_cf_provider_v2/builtin/providers/cf)
* [BIG IP](https://github.com/DealerDotCom/terraform-provider-bigip)

This build is provded as is with no warranty under the same [Mozilla Public License v2](https://www.mozilla.org/en-US/MPL/2.0/) as that of Terraform core. Please do not fork this repository as is used to publish the [website](https://terraform.cfapps.io/), which includes the documentation for the additional providers including download links for the pre-build binaries.

---

- Website: https://www.terraform.io
- [![Gitter chat](https://badges.gitter.im/hashicorp-terraform/Lobby.png)](https://gitter.im/hashicorp-terraform/Lobby)
- Mailing list: [Google Groups](http://groups.google.com/group/terraform-tool)

![Terraform](https://rawgithub.com/hashicorp/terraform/master/website/source/assets/images/logo-hashicorp.svg)

Terraform is a tool for building, changing, and versioning infrastructure safely and efficiently. Terraform can manage existing and popular service providers as well as custom in-house solutions.

The key features of Terraform are:

- **Infrastructure as Code**: Infrastructure is described using a high-level configuration syntax. This allows a blueprint of your datacenter to be versioned and treated as you would any other code. Additionally, infrastructure can be shared and re-used.

- **Execution Plans**: Terraform has a "planning" step where it generates an *execution plan*. The execution plan shows what Terraform will do when you call apply. This lets you avoid any surprises when Terraform manipulates infrastructure.

- **Resource Graph**: Terraform builds a graph of all your resources, and parallelizes the creation and modification of any non-dependent resources. Because of this, Terraform builds infrastructure as efficiently as possible, and operators get insight into dependencies in their infrastructure.

- **Change Automation**: Complex changesets can be applied to your infrastructure with minimal human interaction. With the previously mentioned execution plan and resource graph, you know exactly what Terraform will change and in what order, avoiding many possible human errors.

For more information, see the [introduction section](http://www.terraform.io/intro) of the Terraform website.
