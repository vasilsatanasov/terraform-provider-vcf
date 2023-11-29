// Copyright 2023 Broadcom. All Rights Reserved.
// SPDX-License-Identifier: MPL-2.0

package certificates

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/vmware/terraform-provider-vcf/internal/api_client"
	"github.com/vmware/terraform-provider-vcf/internal/constants"
	validationutils "github.com/vmware/terraform-provider-vcf/internal/validation"
	"github.com/vmware/vcf-sdk-go/client"
	vcfclient "github.com/vmware/vcf-sdk-go/client"
	"github.com/vmware/vcf-sdk-go/client/certificates"
	"github.com/vmware/vcf-sdk-go/client/domains"
	"github.com/vmware/vcf-sdk-go/models"
	"strings"
	"time"
)

func GetFqdnOfResourceTypeInDomain(ctx context.Context, domainId, resourceType string, apiClient *client.VcfClient) (*string, error) {
	if apiClient == nil || len(resourceType) < 1 || len(domainId) < 1 {
		return nil, nil
	}

	endpointsParams := domains.NewGetDomainEndpointsParamsWithContext(ctx).
		WithTimeout(constants.DefaultVcfApiCallTimeout).WithID(domainId)

	endpointsResult, err := apiClient.Domains.GetDomainEndpoints(endpointsParams)
	if err != nil {
		return nil, err
	}

	for _, endpoint := range endpointsResult.Payload.Elements {
		if resourceType == *endpoint.Type {
			result := *endpoint.URL
			if strings.Contains(result, "https://") {
				result = strings.ReplaceAll(result, "https://", "")
			}
			return &result, nil
		}
	}
	return nil, nil
}

func ValidateResourceCertificates(ctx context.Context, client *vcfclient.VcfClient,
	domainId string, resourceCertificateSpecs []*models.ResourceCertificateSpec) diag.Diagnostics {
	validateResourceCertificatesParams := certificates.NewValidateResourceCertificatesParams().
		WithContext(ctx).WithTimeout(constants.DefaultVcfApiCallTimeout).
		WithID(domainId)
	validateResourceCertificatesParams.SetResourceCertificateSpecs(resourceCertificateSpecs)

	var validationResponse *models.CertificateValidationTask
	okResponse, acceptedResponse, err := client.Certificates.ValidateResourceCertificates(validateResourceCertificatesParams)
	if okResponse != nil {
		validationResponse = okResponse.Payload
	}
	if acceptedResponse != nil {
		validationResponse = acceptedResponse.Payload
	}
	if err != nil {
		return validationutils.ConvertVcfErrorToDiag(err)
	}
	if validationutils.HaveCertificateValidationsFailed(validationResponse) {
		return validationutils.ConvertCertificateValidationsResultToDiag(validationResponse)
	}
	validationId := validationResponse.ValidationID
	// Wait for certificate validation to fisnish
	if !validationutils.HasCertificateValidationFinished(validationResponse) {
		for {
			getResourceCertificatesValidationResultParams := certificates.NewGetResourceCertificatesValidationResultParams().
				WithContext(ctx).
				WithTimeout(constants.DefaultVcfApiCallTimeout).
				WithID(*validationId)
			getValidationResponse, err := client.Certificates.GetResourceCertificatesValidationResult(getResourceCertificatesValidationResultParams)
			if err != nil {
				return validationutils.ConvertVcfErrorToDiag(err)
			}
			validationResponse = getValidationResponse.Payload
			if validationutils.HasCertificateValidationFinished(validationResponse) {
				break
			}
			time.Sleep(10 * time.Second)
		}
	}
	if err != nil {
		return validationutils.ConvertVcfErrorToDiag(err)
	}
	if validationutils.HaveCertificateValidationsFailed(validationResponse) {
		return validationutils.ConvertCertificateValidationsResultToDiag(validationResponse)
	}

	return nil
}

func GetCertificateForResourceInDomain(ctx context.Context, client *vcfclient.VcfClient,
	domainId, resourceType string) (*models.Certificate, error) {
	resourceFqdn, err := GetFqdnOfResourceTypeInDomain(ctx, domainId, resourceType, client)
	if err != nil {
		return nil, err
	}
	if resourceFqdn == nil {
		return nil, fmt.Errorf("could not determine FQDN for resourceType %s in domain %s", resourceType, domainId)
	}

	viewCertificatesParams := certificates.NewViewCertificateParamsWithContext(ctx).
		WithTimeout(constants.DefaultVcfApiCallTimeout)
	viewCertificatesParams.SetDomainName(domainId)

	certificatesResponse, err := client.Certificates.ViewCertificate(viewCertificatesParams)
	if err != nil {
		return nil, err
	}

	allCertsForDomain := certificatesResponse.Payload.Elements
	for _, cert := range allCertsForDomain {
		if cert.IssuedTo != nil && *cert.IssuedTo == *resourceFqdn {
			return cert, nil
		}
	}
	return nil, nil
}

func GenerateCertificateForResource(ctx context.Context, client *api_client.SddcManagerClient,
	domainId, resourceType, resourceFqdn, caType *string) error {

	certificateGenerationSpec := &models.CertificatesGenerationSpec{
		CaType: caType,
		Resources: []*models.Resource{{
			Fqdn: *resourceFqdn,
			Type: resourceType,
		}},
	}
	generateCertificatesParam := certificates.NewGenerateCertificatesParamsWithContext(ctx).
		WithTimeout(constants.DefaultVcfApiCallTimeout).
		WithDomainName(*domainId)
	generateCertificatesParam.SetCertificateGenerationSpec(certificateGenerationSpec)

	var taskId string
	responseOk, responseAccepted, err := client.ApiClient.Certificates.GenerateCertificates(generateCertificatesParam)
	if err != nil {
		return err
	}
	if responseOk != nil {
		taskId = responseOk.Payload.ID
	}
	if responseAccepted != nil {
		taskId = responseAccepted.Payload.ID
	}
	err = client.WaitForTaskComplete(ctx, taskId, true)
	if err != nil {
		return err
	}
	return nil
}
