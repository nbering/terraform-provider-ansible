---
title: Installation
layout: default
nav_order: 1
---

## Installation

Installation can be accomplished in two different ways:

1. Installing a pre-compiled release (recommended)
2. Compiling from source

### Installing a Pre-Compiled Release

Downloading and installing a pre-compiled `terraform-provider-ansible` release
is the recommended method of installation since it requires no additional tools
or libraries to be installed on your workstation.

1. Visit the [releases][4] page and download the latest release for your target
   architecture.

2. Unzip the downloaded file and copy the `terraform-provider-ansible` binary
   to a designated directory as described in Terraform's [plugin installation
   instructions][2].

### Compiling From Source

> Note: Terraform requires Go 1.11 or later to successfully compile.

> Note: Dependencies are no longer included in this repository. You may need the
[bazaar][10] version control utility to download some of Terraform's Go-lang
module dependencies.

If you'd like to take advantage of features not yet available in a pre-compiled
release, you can compile `terraform-provider-ansible` from source.

In order to compile, you will need to have Go installed on your workstation.
Official instructions on how to install Go can be found [here][5].

Alternatively, you can use [gimme][6] as a quick and easy way to install Go:

```shell
$ sudo wget -O /usr/local/bin/gimme https://raw.githubusercontent.com/travis-ci/gimme/master/gimme
$ sudo chmod +x /usr/local/bin/gimme
$ gimme 1.10
# copy the output to your `.bashrc` and source `.bashrc`.
```

Once you have a working Go installation, you can compile
`terraform-provider-ansible` by doing the following:

```shell
$ go get github.com/nbering/terraform-provider-ansible
$ cd $GOPATH/src/github.com/nbering/terraform-provider-ansible
$ make
```

You should now have a `terraform-provider-ansible` binary located at
`$GOPATH/bin/terraform-provider-ansible`. Copy this binary to a designated
directory as described in Terraform's [plugin installation instructions][2]

[2]: https://www.terraform.io/docs/plugins/basics.html#installing-a-plugin
[4]: https://github.com/nbering/terraform-provider-ansible/releases
[5]: https://golang.org/doc/install
[6]: https://github.com/travis-ci/gimme
