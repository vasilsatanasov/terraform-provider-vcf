# CHANGELOG

## [v0.6.0](https://github.com/vmware/terraform-provider-vcf/releases/tag/v0.6.0)

> Release Date: Nov 23 2023

FEATURES:
* Support for CA configuration
* Support for CSR generation
* Support for replacing certificate of a resource in a Domain via configured CA
* Support for replacing certificate of a resource in a Domain via external CA

## [v0.5.0](https://github.com/vmware/terraform-provider-vcf/releases/tag/v0.5.0)

> Release Date: Oct 9th 2023

FEATURES:
* Add support for management domain deployment (bringup) [\#38](https://github.com/vmware/terraform-provider-vcf/issues/38)

**Note:** Provider has two mutually exclusive modes of operation: CloudBuilder mode (for bringup) and SDDCManager mode

## [v0.4.0](https://github.com/vmware/terraform-provider-vcf/releases/tag/v0.4.0)

> Release Date: Sep 11th 2023

BREAKING CHANGES:

* Removed the attribute "host_id" from the "vcf_host" resource and replaced it with just "id" as per Terraform standard practice. This way users can refer to the Host UUID (in cluster host spec for example) in the standard way, e.g. "vcf_host.host1.id"
* Replaced attribute "nsx_cluster_ref" from the "vcf_domain" datasource with a richer "nsx_configuration", that additionally contains IPs, Names and DNS Names of NSX-T Manager Nodes
* Renamed attribute "nsx_configuration.nsx_manager_node.dns_name" in "vcf_domain" to "nsx_configuration.nsx_manager_node.fqdn" for clarity
* Renamed attribute "vcenter" to "vcenter_configuration" in "vcf_domain" resource and "vcf_domain" datasource
* Replaced attribute "dns_name" in "vcenter_configuration" in "vcf_domain" resource with "fqdn"
* Replaced attribute "vcenter_fqdn" and "vcenter_id from the "vcf_domain" datasource with "vcenter_configuration" subresource, that contains "id" and "fqdn" attributes. 

FEATURES:
* Extend support for host resource: import [\#36](https://github.com/vmware/terraform-provider-vcf/issues/36)
* Add support for workload domain resource: import [\#35](https://github.com/vmware/terraform-provider-vcf/issues/35)
* Add support for configuration of NSX host TEP pool (static / DHCP) in r/vcf_domain [\#54](https://github.com/vmware/terraform-provider-vcf/issues/54)

**Note:** Management domain cannot be imported, but can be used as datasource

BUG FIXES:
* Include "domain_id" attribute to both imported cluster and cluster datasource [\#49](https://github.com/vmware/terraform-provider-vcf/issues/49)

## [v0.3.0](https://github.com/vmware/terraform-provider-vcf/releases/tag/v0.3.0)

> Release Date: Aug 22nd 2023

FEATURES:
* Add support for workload domain cluster resource: read, add, update, delete [\#32](https://github.com/vmware/terraform-provider-vcf/issues/32)
* Add support for workload domain cluster data source [\#32](https://github.com/vmware/terraform-provider-vcf/issues/34)
* Extend support for workload domain cluster resource: import [\#33](https://github.com/vmware/terraform-provider-vcf/issues/33)
* Extend support for workload domain cluster: expand and contract [\#37](https://github.com/vmware/terraform-provider-vcf/issues/37)

BUG FIXES:
* Fix IsEmpty not checking for boolean [\#45](https://github.com/vmware/terraform-provider-vcf/pull/45)



## [v0.2.0](https://github.com/vmware/terraform-provider-vcf/releases/tag/v0.2.0)

> Release Date: Jul 25th 2023

Add support for creating/deleting workload domains and being used as datasource.

## [v0.1.0](https://github.com/vmware/terraform-provider-vcf/releases/tag/v0.1.0)

> Release Date: Jun 5th 2023

Initial release, adding support for commissioning/decommissioning hosts, creating/destroying
network pools, creating/destroying SSO user, turning on/off the telemetry (CEIP).
