[![Build Status](https://travis-ci.com/IBM/networking-go-sdk.svg?branch=master)](https://travis-ci.com/IBM/networking-go-sdk)
[![Release](https://img.shields.io/github/v/release/IBM/networking-go-sdk)](https://github.com/IBM/networking-go-sdk/releases/latest)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/IBM/networking-go-sdk)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![semantic-release](https://img.shields.io/badge/%20%20%F0%9F%93%A6%F0%9F%9A%80-semantic--release-e10079.svg)](https://github.com/semantic-release/semantic-release)

# IBM Cloud Networking Go SDK Version 0.17.0

Go client library to interact with the various [IBM Cloud Networking Service APIs](https://cloud.ibm.com/apidocs?category=<networking>).

## Table of Contents

<!--
  The TOC below is generated using the `markdown-toc` node package.

      https://github.com/jonschlinkert/markdown-toc

  You should regenerate the TOC after making changes to this file.

      npx markdown-toc -i README.md
  -->

<!-- toc -->

- [Overview](#overview)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
  - [`go get` command](#go-get-command)
  - [Go modules](#go-modules)
  - [`dep` dependency manager](#dep-dependency-manager)
- [Using the SDK](#using-the-sdk)
- [Questions](#questions)
- [Issues](#issues)
- [Open source @ IBM](#open-source--ibm)
- [Contributing](#contributing)
- [License](#license)

<!-- tocstop -->

## Overview

The IBM Cloud Networking Go SDK allows developers to programmatically interact with the following IBM Cloud services:

| Service Name                                                                                                         | Package name                   |
| -------------------------------------------------------------------------------------------------------------------- | ------------------------------ |
| [Transit Gateway Service](https://cloud.ibm.com/docs/transit-gateway)                                                | transitgatewayapisv1           |
| [Direct Link Service](https://cloud.ibm.com/apidocs/direct_link?code=go)                                             | directlinkv1                   |
| [Direct Link Provider Service](https://cloud.ibm.com/apidocs/direct_link_provider_api?code=go)                       | directlinkproviderv2           |
| [CIS: Cache](https://cloud.ibm.com/apidocs/cis?code=go#purge-all)                                                    | cachingapiv1                   |
| [CIS: IP](https://cloud.ibm.com/apidocs/cis?code=go#list-of-all-ip-addresses-used-by-the-cis-proxy)                  | cisipapiv1                     |
| [CIS: Custom Pages](https://cloud.ibm.com/apidocs/cis?code=go#list-all-custom-pages-for-a-given-instance)            | custompagesv1                  |
| [CIS: DNS Records Bulk](https://cloud.ibm.com/apidocs/cis?code=go#export-zone-file)                                  | dnsrecordbulkv1                |
| [CIS: DNS Records](https://cloud.ibm.com/apidocs/cis?code=go#export-zone-file)                                       | dnsrecordsv1                   |
| [CIS: Firewall Access Rules](https://cloud.ibm.com/apidocs/cis?code=go#list-instance-level-firewall-access-rules)    | firewallaccessrulesv1          |
| [CIS: Security Level Settings](https://cloud.ibm.com/apidocs/cis?code=go#get-security-level-setting)                 | firewallapiv1                  |
| [CIS: GLB Events](https://cloud.ibm.com/apidocs/cis?code=go#list-all-load-balancer-events)                           | globalloadbalancereventsv1     |
| [CIS: GLB Monitor](https://cloud.ibm.com/apidocs/cis?code=go#list-all-load-balancer-monitors)                        | globalloadbalancermonitorv1    |
| [CIS: GLB Pools](https://cloud.ibm.com/apidocs/cis?code=go#list-all-pools)                                           | globalloadbalancerpoolsv0      |
| [CIS: GLB Service](https://cloud.ibm.com/apidocs/cis?code=go#list-all-load-balancers)                                | globalloadbalancerv1           |
| [CIS: Page Rules](https://cloud.ibm.com/apidocs/cis?code=go#get-page-rule)                                           | pageruleapiv1                  |
| [CIS: Range Application](https://cloud.ibm.com/apidocs/cis?code=go#list-range-applications)                          | rangeapplicationsv1            |
| [CIS: Routing](https://cloud.ibm.com/apidocs/cis?code=go#get-routing-feature-smart-routing-setting)                  | routingv1                      |
| [CIS: Security Events](https://cloud.ibm.com/apidocs/cis?code=go#logs-of-the-mitigations-performed-by-firewall-feat) | securityeventsapiv1            |
| [CIS: SSL/TLS](https://cloud.ibm.com/apidocs/cis?code=go#list-all-certificates)                                      | sslcertificateapiv1            |
| [CIS: User Agent Blocking Rules](https://cloud.ibm.com/apidocs/cis?code=go#list-all-user-agent-blocking-rules)       | useragentblockingrulesv1       |
| [CIS: WAF Settings](https://cloud.ibm.com/apidocs/cis?code=go#get-waf-setting)                                       | wafapiv1                       |
| [CIS: WAF Rule Groups](https://cloud.ibm.com/apidocs/cis?code=go#list-all-waf-rule-groups)                           | wafrulegroupsapiv1             |
| [CIS: WAF Rule Packages](https://cloud.ibm.com/apidocs/cis?code=go#list-all-waf-rule-packages)                       | wafrulepackagesapiv1           |
| [CIS: WAF Rules](https://cloud.ibm.com/apidocs/cis?code=go#list-all-waf-rules)                                       | wafrulesapiv1                  |
| [CIS: Zone Firewall Access Rules](https://cloud.ibm.com/apidocs/cis?code=go#list-all-firewall-access-rules)          | zonefirewallaccessrulesv1      |
| [CIS: Zone Lockdown](https://cloud.ibm.com/apidocs/cis?code=go#list-all-lockdown-rules)                              | zonelockdownv1                 |
| [CIS: Zone Rate Limits](https://cloud.ibm.com/apidocs/cis?code=go#list-all-rate-limits)                              | zoneratelimitsv1               |
| [CIS: Zone Settings](https://cloud.ibm.com/apidocs/cis?code=go#get-zone-dnssec)                                      | zonessettingsv1                |
| [CIS: Zones](https://cloud.ibm.com/apidocs/cis?code=go#list-all-zones)                                               | zonesv1                        |
| [PDNS: DNS Zones](https://cloud.ibm.com/apidocs/dns-svcs?code=go#list-dns-zones)                                 | ~dnszonesv1~ dnssvcsv1 |
| [PDNS: Resource Records](https://cloud.ibm.com/apidocs/dns-svcs?code=go#list-resource-records) | ~resourcerecordsv1~ dnssvcsv1 |
| [PDNS: Permitted Networks](https://cloud.ibm.com/apidocs/dns-svcs?code=go#list-permitted-networks) | ~permittednetworksfordnszonesv1~ dnssvcsv1|
| [PDNS: Global Load Balancers](https://cloud.ibm.com/apidocs/dns-svcs?code=go) | ~globalloadbalancersv1~ dnssvcsv1 |
| [PDNS: DNS Services](https://cloud.ibm.com/apidocs/dns-svcs?code=go) | dnssvcsv1 |

## Prerequisites

[ibm-cloud-onboarding]: https://cloud.ibm.com/registration

- An [IBM Cloud][ibm-cloud-onboarding] account.
- An IAM API key to allow the SDK to access your account. Create one [here](https://cloud.ibm.com/iam/apikeys).
- Go version 1.12 or above.

## Installation

The current version of this SDK: 0.17.0

There are a few different ways to download and install the Networking Go SDK project for use by your
Go application:

#### `go get` command

Use this command to download and install the SDK to allow your Go application to
use it:

```
go get -u github.com/IBM/networking-go-sdk
```

#### Go modules

If your application is using Go modules, you can add a suitable import to your
Go application, like this:

```go
import (
  "github.com/IBM/networking-go-sdk/transitgatewayapisv1"
  "github.com/IBM/networking-go-sdk/directlinkv1"
  "github.com/IBM/networking-go-sdk/directlinkproviderv2"
  "github.com/IBM/networking-go-sdk/cachingapiv1"
  "github.com/IBM/networking-go-sdk/cisipapiv1"
  "github.com/IBM/networking-go-sdk/custompagesv1"
  "github.com/IBM/networking-go-sdk/dnsrecordbulkv1"
  "github.com/IBM/networking-go-sdk/dnsrecordsv1"
  "github.com/IBM/networking-go-sdk/firewallaccessrulesv1"
  "github.com/IBM/networking-go-sdk/firewallapiv1"
  "github.com/IBM/networking-go-sdk/globalloadbalancereventsv1"
  "github.com/IBM/networking-go-sdk/globalloadbalancermonitorv1"
  "github.com/IBM/networking-go-sdk/globalloadbalancerpoolsv0"
  "github.com/IBM/networking-go-sdk/globalloadbalancerv1"
  "github.com/IBM/networking-go-sdk/pageruleapiv1"
  "github.com/IBM/networking-go-sdk/rangeapplicationsv1"
  "github.com/IBM/networking-go-sdk/routingv1"
  "github.com/IBM/networking-go-sdk/securityeventsapiv1"
  "github.com/IBM/networking-go-sdk/sslcertificateapiv1"
  "github.com/IBM/networking-go-sdk/useragentblockingrulesv1"
  "github.com/IBM/networking-go-sdk/wafapiv1"
  "github.com/IBM/networking-go-sdk/wafrulegroupsapiv1"
  "github.com/IBM/networking-go-sdk/wafrulepackagesapiv1"
  "github.com/IBM/networking-go-sdk/wafrulesapiv1"
  "github.com/IBM/networking-go-sdk/zonefirewallaccessrulesv1"
  "github.com/IBM/networking-go-sdk/zonelockdownv1"
  "github.com/IBM/networking-go-sdk/zoneratelimitsv1"
  "github.com/IBM/networking-go-sdk/zonessettingsv1"
  "github.com/IBM/networking-go-sdk/zonesv1"
  "github.com/IBM/networking-go-sdk/dnszonesv1"
  "github.com/IBM/networking-go-sdk/resourcerecordsv1"
  "github.com/IBM/networking-go-sdk/permittednetworksfordnszonesv1"`
  "github.com/IBM/networking-go-sdk/globalloadbalancersv1"
  "github.com/IBM/networking-go-sdk/dnssvcsv1"
)
```

then run `go mod tidy` to download and install the new dependency and update your Go application's
`go.mod` file.

#### `dep` dependency manager

If your application is using the `dep` dependency management tool, you can add a dependency
to your `Gopkg.toml` file. Here is an example:

```
[[constraint]]
  name = "github.com/IBM/networking-go-sdk"
  version = "0.17.0"

```

then run `dep ensure`.

## Using the SDK

For general SDK usage information, please see [this link](https://github.com/IBM/ibm-cloud-sdk-common/blob/master/README.md)

## Questions

If you are having difficulties using this SDK or have a question about the IBM Cloud services,
please ask a question at
[Stack Overflow](http://stackoverflow.com/questions/ask?tags=ibm-cloud).

## Issues

If you encounter an issue with the project, you are welcome to submit a
[bug report](https://github.com/IBM/networking-go-sdk/issues).
Before that, please search for similar issues. It's possible that someone has already reported the problem.

## Deprecation Notice

For deprecation notice, please see [this link](https://github.com/IBM/networking-go-sdk/blob/master/DEPRECATION-NOTICE.md)

## Open source @ IBM

Find more open source projects on the [IBM Github Page](http://ibm.github.io/)

## Contributing

See [CONTRIBUTING](CONTRIBUTING.md).

## License

This SDK project is released under the Apache 2.0 license.
The license's full text can be found in [LICENSE](LICENSE).
