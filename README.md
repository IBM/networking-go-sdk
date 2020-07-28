[![Build Status](https://travis.ibm.com/CloudEngineering/go-sdk-template.svg?token=eW5FVD71iyte6tTby8gr&branch=master)](https://travis.ibm.com/CloudEngineering/go-sdk-template.svg?token=eW5FVD71iyte6tTby8gr&branch=master)
[![semantic-release](https://img.shields.io/badge/%20%20%F0%9F%93%A6%F0%9F%9A%80-semantic--release-e10079.svg)](https://github.com/semantic-release/semantic-release)

# IBM Cloud Networking Go SDK Version 0.1.1

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

| Service Name                                                                                   | Package name                |
| ---------------------------------------------------------------------------------------------- | --------------------------- |
| [Transit Gateway Service](https://cloud.ibm.com/docs/transit-gateway)                          | transitgatewayapisv1        |
| [Direct Link Service](https://cloud.ibm.com/apidocs/direct_link?code=go)                       | directlinkv1            |
| [CIS: Cache](https://cloud.ibm.com/apidocs/cis/cache)                                          | cachingapiv1                |
| [CIS: IP](https://cloud.ibm.com/apidocs/cis/ip)                                                | cisipapiv1                  |
| [CIS: Custom Pages](https://cloud.ibm.com/apidocs/cis)                                         | custompagesv1               |
| [CIS: DNS Records Bulk](https://cloud.ibm.com/apidocs/cis/dnsrecords)                          | dnsrecordbulkv1             |
| [CIS: DNS Records](https://cloud.ibm.com/apidocs/cis/dnsrecords)                               | dnsrecordsv1                |
| [CIS: Firewall Access Rules](https://cloud.ibm.com/apidocs/cis/firewall-access-rule)           | firewallaccessrulesv1       |
| [CIS: Security Level Settings](https://cloud.ibm.com/apidocs/cis/security-level-settings)      | firewallapiv1               |
| [CIS: GLB Events](https://cloud.ibm.com/apidocs/cis/glb-events)                                | globalloadbalancereventsv1  |
| [CIS: GLB Monitor](https://cloud.ibm.com/apidocs/cis/glb-monitor)                              | globalloadbalancermonitorv1 |
| [CIS: GLB Pools](https://cloud.ibm.com/apidocs/cis/glb-pool)                                   | globalloadbalancerpoolsv0   |
| [CIS: GLB Service](https://cloud.ibm.com/apidocs/cis/glb)                                      | globalloadbalancerv1        |
| [CIS: Page Rules](https://cloud.ibm.com/apidocs/cis/page-rules)                                | pageruleapiv1               |
| [CIS: Range Application](https://cloud.ibm.com/apidocs/cis/range)                              | rangeapplicationsv1         |
| [CIS: Routing](https://cloud.ibm.com/apidocs/cis/routing)                                      | routingv1                   |
| [CIS: Security Events](https://cloud.ibm.com/apidocs/cis)                                      | securityeventsapiv1         |
| [CIS: SSL/TLS](https://cloud.ibm.com/apidocs/cis/tls)                                          | sslcertificateapiv1         |
| [CIS: User Agent Blocking Rules](https://cloud.ibm.com/apidocs/cis/user-agent-rules)           | useragentblockingrulesv1    |
| [CIS: WAF Settings](https://cloud.ibm.com/apidocs/cis/waf)                                     | wafapiv1                    |
| [CIS: WAF Rule Groups](https://cloud.ibm.com/apidocs/cis/waf-groups)                           | wafrulegroupsapiv1          |
| [CIS: WAF Rule Packages](https://cloud.ibm.com/apidocs/cis/waf-packages)                       | wafrulepackagesapiv1        |
| [CIS: WAF Rules](https://cloud.ibm.com/apidocs/cis/waf-rules)                                  | wafrulesapiv1               |
| [CIS: Zone Firewall Access Rules](https://cloud.ibm.com/apidocs/cis/zone-firewall-access-rule) | zonefirewallaccessrulesv1   |
| [CIS: Zone Lockdown](https://cloud.ibm.com/apidocs/cis/zone-lockdown)                          | zonelockdownv1              |
| [CIS: Zone Rate Limits](https://cloud.ibm.com/apidocs/cis)                                     | zoneratelimitsv1            |
| [CIS: Zone Settings](https://cloud.ibm.com/apidocs/cis/zonesettings)                           | zonessettingsv1             |
| [CIS: Zones](https://cloud.ibm.com/apidocs/cis/zones)                                          | zonesv1                     |

## Prerequisites

[ibm-cloud-onboarding]: https://cloud.ibm.com/registration

- An [IBM Cloud][ibm-cloud-onboarding] account.
- An IAM API key to allow the SDK to access your account. Create one [here](https://cloud.ibm.com/iam/apikeys).
- Go version 1.12 or above.

## Installation

The current version of this SDK: 0.1.1

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
  version = "0.1.1"

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
[bug report](github.com/IBM/networking-go-sdk/issues).
Before that, please search for similar issues. It's possible that someone has already reported the problem.

## Open source @ IBM

Find more open source projects on the [IBM Github Page](http://ibm.github.io/)

## Contributing

See [CONTRIBUTING](CONTRIBUTING.md).

## License

This SDK project is released under the Apache 2.0 license.
The license's full text can be found in [LICENSE](LICENSE).
