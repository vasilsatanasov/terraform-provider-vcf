/* Copyright 2023 VMware, Inc.
   SPDX-License-Identifier: MPL-2.0 */

package constants

const (

	// VcfTestUrl URL of a VCF instance, used for Acceptance tests.
	VcfTestUrl = "VCF_TEST_URL"
	// VcfTestUsername username of SSO user, used for Acceptance tests.
	VcfTestUsername = "VCF_TEST_USERNAME"
	// VcfTestPassword passowrd of SSO user, used for Acceptance tests.
	VcfTestPassword = "VCF_TEST_PASSWORD"

	// VcfTestHost1Fqdn the FQDN of the first ESXi host, that has not been commissioned
	// with the SDDC Manager.
	VcfTestHost1Fqdn = "VCF_TEST_HOST1_FQDN"

	// VcfTestHost1Pass the SSH pass of the first ESXi host, that has not been commissioned
	// with the SDDC Manager.
	VcfTestHost1Pass = "VCF_TEST_HOST1_PASS"

	// VcfTestHost2Fqdn the FQDN of the second ESXi host, that has not been commissioned
	// with the SDDC Manager.
	VcfTestHost2Fqdn = "VCF_TEST_HOST2_FQDN"

	// VcfTestHost2Pass the SSH pass of the second ESXi host, that has not been commissioned
	// with the SDDC Manager.
	VcfTestHost2Pass = "VCF_TEST_HOST2_PASS"

	// VcfTestHost3Fqdn the FQDN of the third ESXi host, that has not been commissioned
	// with the SDDC Manager.
	VcfTestHost3Fqdn = "VCF_TEST_HOST3_FQDN"

	// VcfTestHost3Pass the SSH pass of the third ESXi host, that has not been commissioned
	// with the SDDC Manager.
	VcfTestHost3Pass = "VCF_TEST_HOST3_PASS"

	// VcfTestHost4Fqdn the FQDN of the forth ESXi host, that has not been commissioned
	// with the SDDC Manager.
	VcfTestHost4Fqdn = "VCF_TEST_HOST4_FQDN"

	// VcfTestHost4Pass the SSH pass of the forth ESXi host, that has not been commissioned
	// with the SDDC Manager.
	VcfTestHost4Pass = "VCF_TEST_HOST4_PASS"

	// VcfTestNsxtLicenseKey license key for NSXT required for domain and cluster acceptance tests.
	VcfTestNsxtLicenseKey = "VCF_TEST_NSXT_LICENSE_KEY"

	// VcfTestEsxiLicenseKey license key for NSXT required for domain and cluster acceptance tests.
	VcfTestEsxiLicenseKey = "VCF_TEST_ESXI_LICENSE_KEY"

	// VcfTestVsanLicenseKey license key for VSAN required for domain and cluster acceptance tests.
	VcfTestVsanLicenseKey = "VCF_TEST_VSAN_LICENSE_KEY"

	// VcfTestGatewayIp the gateway IP required for domain and cluster acceptance tests.
	VcfTestGatewayIp = "VCF_TEST_GATEWAY_IP"

	// VcfTestNetworkPoolName used in vcf_network_pool Acceptance tests.
	VcfTestNetworkPoolName = "terraform-test-pool"
)
