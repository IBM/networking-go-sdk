/**
 * (C) Copyright IBM Corp. 2026.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

/*
 * IBM OpenAPI SDK Code Generator Version: 3.117.0-7f07c563-20260915-094553
 */

// Package aisecurityforappsv1 : Operations and models for the AiSecurityForAppsV1 service
package aisecurityforappsv1

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strings"
	"time"

	common "github.com/IBM/networking-go-sdk/common"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/go-openapi/strfmt"
)

// AiSecurityForAppsV1 : AI Security for Apps
//
// API Version: 1.0.0
type AiSecurityForAppsV1 struct {
	Service *core.BaseService

	// Full url-encoded CRN of the service instance.
	Crn *string

	// Zone identifier to identify the zone.
	ZoneIdentifier *string
}

// DefaultServiceURL is the default URL to make service requests to.
const DefaultServiceURL = "https://api.cis.cloud.ibm.com"

// DefaultServiceName is the default key used to find external configuration information.
const DefaultServiceName = "ai_security_for_apps"

// AiSecurityForAppsV1Options : Service options
type AiSecurityForAppsV1Options struct {
	ServiceName   string
	URL           string
	Authenticator core.Authenticator

	// Full url-encoded CRN of the service instance.
	Crn *string `validate:"required"`

	// Zone identifier to identify the zone.
	ZoneIdentifier *string `validate:"required"`
}

// NewAiSecurityForAppsV1UsingExternalConfig : constructs an instance of AiSecurityForAppsV1 with passed in options and external configuration.
func NewAiSecurityForAppsV1UsingExternalConfig(options *AiSecurityForAppsV1Options) (aiSecurityForApps *AiSecurityForAppsV1, err error) {
	if options.ServiceName == "" {
		options.ServiceName = DefaultServiceName
	}

	if options.Authenticator == nil {
		options.Authenticator, err = core.GetAuthenticatorFromEnvironment(options.ServiceName)
		if err != nil {
			err = core.SDKErrorf(err, "", "env-auth-error", common.GetComponentInfo())
			return
		}
	}

	aiSecurityForApps, err = NewAiSecurityForAppsV1(options)
	err = core.RepurposeSDKProblem(err, "new-client-error")
	if err != nil {
		return
	}

	err = aiSecurityForApps.Service.ConfigureService(options.ServiceName)
	if err != nil {
		err = core.SDKErrorf(err, "", "client-config-error", common.GetComponentInfo())
		return
	}

	if options.URL != "" {
		err = aiSecurityForApps.Service.SetServiceURL(options.URL)
		err = core.RepurposeSDKProblem(err, "url-set-error")
	}
	return
}

// NewAiSecurityForAppsV1 : constructs an instance of AiSecurityForAppsV1 with passed in options.
func NewAiSecurityForAppsV1(options *AiSecurityForAppsV1Options) (service *AiSecurityForAppsV1, err error) {
	serviceOptions := &core.ServiceOptions{
		URL:           DefaultServiceURL,
		Authenticator: options.Authenticator,
	}

	err = core.ValidateStruct(options, "options")
	if err != nil {
		err = core.SDKErrorf(err, "", "invalid-global-options", common.GetComponentInfo())
		return
	}

	baseService, err := core.NewBaseService(serviceOptions)
	if err != nil {
		err = core.SDKErrorf(err, "", "new-base-error", common.GetComponentInfo())
		return
	}

	if options.URL != "" {
		err = baseService.SetServiceURL(options.URL)
		if err != nil {
			err = core.SDKErrorf(err, "", "set-url-error", common.GetComponentInfo())
			return
		}
	}

	service = &AiSecurityForAppsV1{
		Service: baseService,
		Crn: options.Crn,
		ZoneIdentifier: options.ZoneIdentifier,
	}

	return
}

// GetServiceURLForRegion returns the service URL to be used for the specified region
func GetServiceURLForRegion(region string) (string, error) {
	return "", core.SDKErrorf(nil, "service does not support regional URLs", "no-regional-support", common.GetComponentInfo())
}

// Clone makes a copy of "aiSecurityForApps" suitable for processing requests.
func (aiSecurityForApps *AiSecurityForAppsV1) Clone() *AiSecurityForAppsV1 {
	if core.IsNil(aiSecurityForApps) {
		return nil
	}
	clone := *aiSecurityForApps
	clone.Service = aiSecurityForApps.Service.Clone()
	return &clone
}

// SetServiceURL sets the service URL
func (aiSecurityForApps *AiSecurityForAppsV1) SetServiceURL(url string) error {
	err := aiSecurityForApps.Service.SetServiceURL(url)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-set-error", common.GetComponentInfo())
	}
	return err
}

// GetServiceURL returns the service URL
func (aiSecurityForApps *AiSecurityForAppsV1) GetServiceURL() string {
	return aiSecurityForApps.Service.GetServiceURL()
}

// SetDefaultHeaders sets HTTP headers to be sent in every request
func (aiSecurityForApps *AiSecurityForAppsV1) SetDefaultHeaders(headers http.Header) {
	aiSecurityForApps.Service.SetDefaultHeaders(headers)
}

// SetEnableGzipCompression sets the service's EnableGzipCompression field
func (aiSecurityForApps *AiSecurityForAppsV1) SetEnableGzipCompression(enableGzip bool) {
	aiSecurityForApps.Service.SetEnableGzipCompression(enableGzip)
}

// GetEnableGzipCompression returns the service's EnableGzipCompression field
func (aiSecurityForApps *AiSecurityForAppsV1) GetEnableGzipCompression() bool {
	return aiSecurityForApps.Service.GetEnableGzipCompression()
}

// EnableRetries enables automatic retries for requests invoked for this service instance.
// If either parameter is specified as 0, then a default value is used instead.
func (aiSecurityForApps *AiSecurityForAppsV1) EnableRetries(maxRetries int, maxRetryInterval time.Duration) {
	aiSecurityForApps.Service.EnableRetries(maxRetries, maxRetryInterval)
}

// DisableRetries disables automatic retries for requests invoked for this service instance.
func (aiSecurityForApps *AiSecurityForAppsV1) DisableRetries() {
	aiSecurityForApps.Service.DisableRetries()
}

// GetAiSecuritySettings : Get AI Security for Apps settings
// Get AI Security for Apps enabled/disabled setting for a given zone.
func (aiSecurityForApps *AiSecurityForAppsV1) GetAiSecuritySettings(getAiSecuritySettingsOptions *GetAiSecuritySettingsOptions) (result *AiSecuritySettingsResp, response *core.DetailedResponse, err error) {
	result, response, err = aiSecurityForApps.GetAiSecuritySettingsWithContext(context.Background(), getAiSecuritySettingsOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// GetAiSecuritySettingsWithContext is an alternate form of the GetAiSecuritySettings method which supports a Context parameter
func (aiSecurityForApps *AiSecurityForAppsV1) GetAiSecuritySettingsWithContext(ctx context.Context, getAiSecuritySettingsOptions *GetAiSecuritySettingsOptions) (result *AiSecuritySettingsResp, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getAiSecuritySettingsOptions, "getAiSecuritySettingsOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"crn": *aiSecurityForApps.Crn,
		"zone_identifier": *aiSecurityForApps.ZoneIdentifier,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = aiSecurityForApps.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(aiSecurityForApps.Service.Options.URL, `/v1/{crn}/zones/{zone_identifier}/ai_security/settings`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("ai_security_for_apps", "V1", "GetAiSecuritySettings")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range getAiSecuritySettingsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = aiSecurityForApps.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "get_ai_security_settings", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalAiSecuritySettingsResp)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// ReplaceZoneAiSecuritySettings : Update AI Security for Apps settings
// Enable or disable AI Security for Apps for a given zone.
func (aiSecurityForApps *AiSecurityForAppsV1) ReplaceZoneAiSecuritySettings(replaceZoneAiSecuritySettingsOptions *ReplaceZoneAiSecuritySettingsOptions) (result *AiSecuritySettingsResp, response *core.DetailedResponse, err error) {
	result, response, err = aiSecurityForApps.ReplaceZoneAiSecuritySettingsWithContext(context.Background(), replaceZoneAiSecuritySettingsOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// ReplaceZoneAiSecuritySettingsWithContext is an alternate form of the ReplaceZoneAiSecuritySettings method which supports a Context parameter
func (aiSecurityForApps *AiSecurityForAppsV1) ReplaceZoneAiSecuritySettingsWithContext(ctx context.Context, replaceZoneAiSecuritySettingsOptions *ReplaceZoneAiSecuritySettingsOptions) (result *AiSecuritySettingsResp, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(replaceZoneAiSecuritySettingsOptions, "replaceZoneAiSecuritySettingsOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"crn": *aiSecurityForApps.Crn,
		"zone_identifier": *aiSecurityForApps.ZoneIdentifier,
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = aiSecurityForApps.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(aiSecurityForApps.Service.Options.URL, `/v1/{crn}/zones/{zone_identifier}/ai_security/settings`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("ai_security_for_apps", "V1", "ReplaceZoneAiSecuritySettings")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range replaceZoneAiSecuritySettingsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if replaceZoneAiSecuritySettingsOptions.Enabled != nil {
		body["enabled"] = replaceZoneAiSecuritySettingsOptions.Enabled
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		err = core.SDKErrorf(err, "", "set-json-body-error", common.GetComponentInfo())
		return
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = aiSecurityForApps.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "replace_zone_ai_security_settings", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalAiSecuritySettingsResp)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// GetApiGatewayDiscovery : Get API Gateway discovery
// Retrieve discovered operations for a zone rendered as OpenAPI schemas. Use this to identify AI-powered endpoints,
// save them to Endpoint Management, and label them to enable AI Security for Apps scanning.
func (aiSecurityForApps *AiSecurityForAppsV1) GetApiGatewayDiscovery(getApiGatewayDiscoveryOptions *GetApiGatewayDiscoveryOptions) (result *ApiGatewayDiscoveryResp, response *core.DetailedResponse, err error) {
	result, response, err = aiSecurityForApps.GetApiGatewayDiscoveryWithContext(context.Background(), getApiGatewayDiscoveryOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// GetApiGatewayDiscoveryWithContext is an alternate form of the GetApiGatewayDiscovery method which supports a Context parameter
func (aiSecurityForApps *AiSecurityForAppsV1) GetApiGatewayDiscoveryWithContext(ctx context.Context, getApiGatewayDiscoveryOptions *GetApiGatewayDiscoveryOptions) (result *ApiGatewayDiscoveryResp, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getApiGatewayDiscoveryOptions, "getApiGatewayDiscoveryOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"crn": *aiSecurityForApps.Crn,
		"zone_identifier": *aiSecurityForApps.ZoneIdentifier,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = aiSecurityForApps.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(aiSecurityForApps.Service.Options.URL, `/v1/{crn}/zones/{zone_identifier}/api_gateway/discovery`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("ai_security_for_apps", "V1", "GetApiGatewayDiscovery")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range getApiGatewayDiscoveryOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = aiSecurityForApps.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "get_api_gateway_discovery", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalApiGatewayDiscoveryResp)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// ListApiGatewayDiscoveryOperations : List API Gateway discovery operations
// Retrieve the most up-to-date list of discovered operations for a zone.
func (aiSecurityForApps *AiSecurityForAppsV1) ListApiGatewayDiscoveryOperations(listApiGatewayDiscoveryOperationsOptions *ListApiGatewayDiscoveryOperationsOptions) (result *DiscoveryOperationsListResp, response *core.DetailedResponse, err error) {
	result, response, err = aiSecurityForApps.ListApiGatewayDiscoveryOperationsWithContext(context.Background(), listApiGatewayDiscoveryOperationsOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// ListApiGatewayDiscoveryOperationsWithContext is an alternate form of the ListApiGatewayDiscoveryOperations method which supports a Context parameter
func (aiSecurityForApps *AiSecurityForAppsV1) ListApiGatewayDiscoveryOperationsWithContext(ctx context.Context, listApiGatewayDiscoveryOperationsOptions *ListApiGatewayDiscoveryOperationsOptions) (result *DiscoveryOperationsListResp, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(listApiGatewayDiscoveryOperationsOptions, "listApiGatewayDiscoveryOperationsOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"crn": *aiSecurityForApps.Crn,
		"zone_identifier": *aiSecurityForApps.ZoneIdentifier,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = aiSecurityForApps.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(aiSecurityForApps.Service.Options.URL, `/v1/{crn}/zones/{zone_identifier}/api_gateway/discovery/operations`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("ai_security_for_apps", "V1", "ListApiGatewayDiscoveryOperations")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range listApiGatewayDiscoveryOperationsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if listApiGatewayDiscoveryOperationsOptions.Diff != nil {
		builder.AddQuery("diff", fmt.Sprint(*listApiGatewayDiscoveryOperationsOptions.Diff))
	}
	if listApiGatewayDiscoveryOperationsOptions.Direction != nil {
		builder.AddQuery("direction", fmt.Sprint(*listApiGatewayDiscoveryOperationsOptions.Direction))
	}
	if listApiGatewayDiscoveryOperationsOptions.Endpoint != nil {
		builder.AddQuery("endpoint", fmt.Sprint(*listApiGatewayDiscoveryOperationsOptions.Endpoint))
	}
	if listApiGatewayDiscoveryOperationsOptions.Host != nil {
		builder.AddQuery("host", strings.Join(listApiGatewayDiscoveryOperationsOptions.Host, ","))
	}
	if listApiGatewayDiscoveryOperationsOptions.Method != nil {
		builder.AddQuery("method", strings.Join(listApiGatewayDiscoveryOperationsOptions.Method, ","))
	}
	if listApiGatewayDiscoveryOperationsOptions.Order != nil {
		builder.AddQuery("order", fmt.Sprint(*listApiGatewayDiscoveryOperationsOptions.Order))
	}
	if listApiGatewayDiscoveryOperationsOptions.Origin != nil {
		builder.AddQuery("origin", fmt.Sprint(*listApiGatewayDiscoveryOperationsOptions.Origin))
	}
	if listApiGatewayDiscoveryOperationsOptions.State != nil {
		builder.AddQuery("state", fmt.Sprint(*listApiGatewayDiscoveryOperationsOptions.State))
	}
	if listApiGatewayDiscoveryOperationsOptions.Page != nil {
		builder.AddQuery("page", fmt.Sprint(*listApiGatewayDiscoveryOperationsOptions.Page))
	}
	if listApiGatewayDiscoveryOperationsOptions.PerPage != nil {
		builder.AddQuery("per_page", fmt.Sprint(*listApiGatewayDiscoveryOperationsOptions.PerPage))
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = aiSecurityForApps.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "list_api_gateway_discovery_operations", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDiscoveryOperationsListResp)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// UpdateZoneApiGatewayDiscoveryOperation : Bulk update discovered operation states
// Bulk update the state of one or more discovered operations. Use to mark operations as saved (promoting to Endpoint
// Management) or ignored.
func (aiSecurityForApps *AiSecurityForAppsV1) UpdateZoneApiGatewayDiscoveryOperation(updateZoneApiGatewayDiscoveryOperationOptions *UpdateZoneApiGatewayDiscoveryOperationOptions) (result *DiscoveryOperationsPatchResp, response *core.DetailedResponse, err error) {
	result, response, err = aiSecurityForApps.UpdateZoneApiGatewayDiscoveryOperationWithContext(context.Background(), updateZoneApiGatewayDiscoveryOperationOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// UpdateZoneApiGatewayDiscoveryOperationWithContext is an alternate form of the UpdateZoneApiGatewayDiscoveryOperation method which supports a Context parameter
func (aiSecurityForApps *AiSecurityForAppsV1) UpdateZoneApiGatewayDiscoveryOperationWithContext(ctx context.Context, updateZoneApiGatewayDiscoveryOperationOptions *UpdateZoneApiGatewayDiscoveryOperationOptions) (result *DiscoveryOperationsPatchResp, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(updateZoneApiGatewayDiscoveryOperationOptions, "updateZoneApiGatewayDiscoveryOperationOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"crn": *aiSecurityForApps.Crn,
		"zone_identifier": *aiSecurityForApps.ZoneIdentifier,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = aiSecurityForApps.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(aiSecurityForApps.Service.Options.URL, `/v1/{crn}/zones/{zone_identifier}/api_gateway/discovery/operations`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("ai_security_for_apps", "V1", "UpdateZoneApiGatewayDiscoveryOperation")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range updateZoneApiGatewayDiscoveryOperationOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	if updateZoneApiGatewayDiscoveryOperationOptions.RequestBody != nil {
		_, err = builder.SetBodyContentJSON(updateZoneApiGatewayDiscoveryOperationOptions.RequestBody)
		if err != nil {
			err = core.SDKErrorf(err, "", "set-json-body-error", common.GetComponentInfo())
			return
		}
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = aiSecurityForApps.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "update_zone_api_gateway_discovery_operation", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDiscoveryOperationsPatchResp)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// CreateZoneApiGatewayOperation : Create API Gateway operations in bulk
// Create API Gateway operations in bulk for a zone, saving them to Endpoint Management.
func (aiSecurityForApps *AiSecurityForAppsV1) CreateZoneApiGatewayOperation(createZoneApiGatewayOperationOptions *CreateZoneApiGatewayOperationOptions) (result *ApiGatewayOperationsResp, response *core.DetailedResponse, err error) {
	result, response, err = aiSecurityForApps.CreateZoneApiGatewayOperationWithContext(context.Background(), createZoneApiGatewayOperationOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// CreateZoneApiGatewayOperationWithContext is an alternate form of the CreateZoneApiGatewayOperation method which supports a Context parameter
func (aiSecurityForApps *AiSecurityForAppsV1) CreateZoneApiGatewayOperationWithContext(ctx context.Context, createZoneApiGatewayOperationOptions *CreateZoneApiGatewayOperationOptions) (result *ApiGatewayOperationsResp, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(createZoneApiGatewayOperationOptions, "createZoneApiGatewayOperationOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"crn": *aiSecurityForApps.Crn,
		"zone_identifier": *aiSecurityForApps.ZoneIdentifier,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = aiSecurityForApps.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(aiSecurityForApps.Service.Options.URL, `/v1/{crn}/zones/{zone_identifier}/api_gateway/operations`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("ai_security_for_apps", "V1", "CreateZoneApiGatewayOperation")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range createZoneApiGatewayOperationOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	if createZoneApiGatewayOperationOptions.ApiGatewayOperation != nil {
		_, err = builder.SetBodyContentJSON(createZoneApiGatewayOperationOptions.ApiGatewayOperation)
		if err != nil {
			err = core.SDKErrorf(err, "", "set-json-body-error", common.GetComponentInfo())
			return
		}
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = aiSecurityForApps.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "create_zone_api_gateway_operation", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalApiGatewayOperationsResp)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// CreateApiGatewayOperationItem : Create a single API Gateway operation
// Create a single API Gateway operation for a zone, saving it to Endpoint Management.
func (aiSecurityForApps *AiSecurityForAppsV1) CreateApiGatewayOperationItem(createApiGatewayOperationItemOptions *CreateApiGatewayOperationItemOptions) (result *ApiGatewayOperationItemResp, response *core.DetailedResponse, err error) {
	result, response, err = aiSecurityForApps.CreateApiGatewayOperationItemWithContext(context.Background(), createApiGatewayOperationItemOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// CreateApiGatewayOperationItemWithContext is an alternate form of the CreateApiGatewayOperationItem method which supports a Context parameter
func (aiSecurityForApps *AiSecurityForAppsV1) CreateApiGatewayOperationItemWithContext(ctx context.Context, createApiGatewayOperationItemOptions *CreateApiGatewayOperationItemOptions) (result *ApiGatewayOperationItemResp, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(createApiGatewayOperationItemOptions, "createApiGatewayOperationItemOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"crn": *aiSecurityForApps.Crn,
		"zone_identifier": *aiSecurityForApps.ZoneIdentifier,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = aiSecurityForApps.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(aiSecurityForApps.Service.Options.URL, `/v1/{crn}/zones/{zone_identifier}/api_gateway/operations/item`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("ai_security_for_apps", "V1", "CreateApiGatewayOperationItem")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range createApiGatewayOperationItemOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if createApiGatewayOperationItemOptions.Method != nil {
		body["method"] = createApiGatewayOperationItemOptions.Method
	}
	if createApiGatewayOperationItemOptions.Host != nil {
		body["host"] = createApiGatewayOperationItemOptions.Host
	}
	if createApiGatewayOperationItemOptions.Endpoint != nil {
		body["endpoint"] = createApiGatewayOperationItemOptions.Endpoint
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		err = core.SDKErrorf(err, "", "set-json-body-error", common.GetComponentInfo())
		return
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = aiSecurityForApps.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "create_api_gateway_operation_item", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalApiGatewayOperationItemResp)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// UpdateApiGatewayOperationLabels : Add or remove labels from API Gateway operations
// Add or remove labels from one or more API Gateway operations. Apply the built-in LLM label to endpoints that receive
// LLM traffic to enable IBM AI Security for Apps to scan those endpoints for prompt injection, PII, and unsafe topics.
func (aiSecurityForApps *AiSecurityForAppsV1) UpdateApiGatewayOperationLabels(updateApiGatewayOperationLabelsOptions *UpdateApiGatewayOperationLabelsOptions) (result *ApiGatewayOperationsLabelsResp, response *core.DetailedResponse, err error) {
	result, response, err = aiSecurityForApps.UpdateApiGatewayOperationLabelsWithContext(context.Background(), updateApiGatewayOperationLabelsOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// UpdateApiGatewayOperationLabelsWithContext is an alternate form of the UpdateApiGatewayOperationLabels method which supports a Context parameter
func (aiSecurityForApps *AiSecurityForAppsV1) UpdateApiGatewayOperationLabelsWithContext(ctx context.Context, updateApiGatewayOperationLabelsOptions *UpdateApiGatewayOperationLabelsOptions) (result *ApiGatewayOperationsLabelsResp, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(updateApiGatewayOperationLabelsOptions, "updateApiGatewayOperationLabelsOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"crn": *aiSecurityForApps.Crn,
		"zone_identifier": *aiSecurityForApps.ZoneIdentifier,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = aiSecurityForApps.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(aiSecurityForApps.Service.Options.URL, `/v1/{crn}/zones/{zone_identifier}/api_gateway/operations/labels`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("ai_security_for_apps", "V1", "UpdateApiGatewayOperationLabels")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range updateApiGatewayOperationLabelsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if updateApiGatewayOperationLabelsOptions.User != nil {
		body["user"] = updateApiGatewayOperationLabelsOptions.User
	}
	if updateApiGatewayOperationLabelsOptions.Managed != nil {
		body["managed"] = updateApiGatewayOperationLabelsOptions.Managed
	}
	if updateApiGatewayOperationLabelsOptions.Selector != nil {
		body["selector"] = updateApiGatewayOperationLabelsOptions.Selector
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		err = core.SDKErrorf(err, "", "set-json-body-error", common.GetComponentInfo())
		return
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = aiSecurityForApps.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "update_api_gateway_operation_labels", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalApiGatewayOperationsLabelsResp)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// GetZoneApiGatewayOperation : Retrieve information about an operation
// Retrieve information about a specific operation on a zone.
func (aiSecurityForApps *AiSecurityForAppsV1) GetZoneApiGatewayOperation(getZoneApiGatewayOperationOptions *GetZoneApiGatewayOperationOptions) (result *ApiGatewayOperationItemResp, response *core.DetailedResponse, err error) {
	result, response, err = aiSecurityForApps.GetZoneApiGatewayOperationWithContext(context.Background(), getZoneApiGatewayOperationOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// GetZoneApiGatewayOperationWithContext is an alternate form of the GetZoneApiGatewayOperation method which supports a Context parameter
func (aiSecurityForApps *AiSecurityForAppsV1) GetZoneApiGatewayOperationWithContext(ctx context.Context, getZoneApiGatewayOperationOptions *GetZoneApiGatewayOperationOptions) (result *ApiGatewayOperationItemResp, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getZoneApiGatewayOperationOptions, "getZoneApiGatewayOperationOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(getZoneApiGatewayOperationOptions, "getZoneApiGatewayOperationOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"crn": *aiSecurityForApps.Crn,
		"zone_identifier": *aiSecurityForApps.ZoneIdentifier,
		"operation_id": *getZoneApiGatewayOperationOptions.OperationID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = aiSecurityForApps.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(aiSecurityForApps.Service.Options.URL, `/v1/{crn}/zones/{zone_identifier}/api_gateway/operations/{operation_id}`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("ai_security_for_apps", "V1", "GetZoneApiGatewayOperation")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range getZoneApiGatewayOperationOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = aiSecurityForApps.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "get_zone_api_gateway_operation", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalApiGatewayOperationItemResp)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// DeleteZoneApiGatewayOperation : Delete an operation
// Delete an operation from a zone.
func (aiSecurityForApps *AiSecurityForAppsV1) DeleteZoneApiGatewayOperation(deleteZoneApiGatewayOperationOptions *DeleteZoneApiGatewayOperationOptions) (response *core.DetailedResponse, err error) {
	response, err = aiSecurityForApps.DeleteZoneApiGatewayOperationWithContext(context.Background(), deleteZoneApiGatewayOperationOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// DeleteZoneApiGatewayOperationWithContext is an alternate form of the DeleteZoneApiGatewayOperation method which supports a Context parameter
func (aiSecurityForApps *AiSecurityForAppsV1) DeleteZoneApiGatewayOperationWithContext(ctx context.Context, deleteZoneApiGatewayOperationOptions *DeleteZoneApiGatewayOperationOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteZoneApiGatewayOperationOptions, "deleteZoneApiGatewayOperationOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(deleteZoneApiGatewayOperationOptions, "deleteZoneApiGatewayOperationOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"crn": *aiSecurityForApps.Crn,
		"zone_identifier": *aiSecurityForApps.ZoneIdentifier,
		"operation_id": *deleteZoneApiGatewayOperationOptions.OperationID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = aiSecurityForApps.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(aiSecurityForApps.Service.Options.URL, `/v1/{crn}/zones/{zone_identifier}/api_gateway/operations/{operation_id}`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("ai_security_for_apps", "V1", "DeleteZoneApiGatewayOperation")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range deleteZoneApiGatewayOperationOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	response, err = aiSecurityForApps.Service.Request(request, nil)
	if err != nil {
		core.EnrichHTTPProblem(err, "delete_zone_api_gateway_operation", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}

	return
}

// GetApiGatewaySchemas : Get API Gateway schemas
// Retrieve API Gateway schemas for a specified zone rendered as OpenAPI schemas.
func (aiSecurityForApps *AiSecurityForAppsV1) GetApiGatewaySchemas(getApiGatewaySchemasOptions *GetApiGatewaySchemasOptions) (result *ApiGatewaySchemasResp, response *core.DetailedResponse, err error) {
	result, response, err = aiSecurityForApps.GetApiGatewaySchemasWithContext(context.Background(), getApiGatewaySchemasOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// GetApiGatewaySchemasWithContext is an alternate form of the GetApiGatewaySchemas method which supports a Context parameter
func (aiSecurityForApps *AiSecurityForAppsV1) GetApiGatewaySchemasWithContext(ctx context.Context, getApiGatewaySchemasOptions *GetApiGatewaySchemasOptions) (result *ApiGatewaySchemasResp, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getApiGatewaySchemasOptions, "getApiGatewaySchemasOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"crn": *aiSecurityForApps.Crn,
		"zone_identifier": *aiSecurityForApps.ZoneIdentifier,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = aiSecurityForApps.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(aiSecurityForApps.Service.Options.URL, `/v1/{crn}/zones/{zone_identifier}/api_gateway/schemas`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("ai_security_for_apps", "V1", "GetApiGatewaySchemas")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range getApiGatewaySchemasOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = aiSecurityForApps.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "get_api_gateway_schemas", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalApiGatewaySchemasResp)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}
func getServiceComponentInfo() *core.ProblemComponent {
	return core.NewProblemComponent(DefaultServiceName, "1.0.0")
}

// AiSecuritySettingsRespResult : Container for response information.
type AiSecuritySettingsRespResult struct {
	// Whether AI Security for Apps is enabled on the zone.
	Enabled *bool `json:"enabled,omitempty"`
}

// UnmarshalAiSecuritySettingsRespResult unmarshals an instance of AiSecuritySettingsRespResult from the specified map of raw messages.
func UnmarshalAiSecuritySettingsRespResult(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AiSecuritySettingsRespResult)
	err = core.UnmarshalPrimitive(m, "enabled", &obj.Enabled)
	if err != nil {
		err = core.SDKErrorf(err, "", "enabled-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ApiGatewayOperationItemRespResult : ApiGatewayOperationItemRespResult struct
type ApiGatewayOperationItemRespResult struct {
	// UUID of the created operation.
	OperationID *string `json:"operation_id,omitempty"`

	Method *string `json:"method,omitempty"`

	Host *string `json:"host,omitempty"`

	Endpoint *string `json:"endpoint,omitempty"`
}

// UnmarshalApiGatewayOperationItemRespResult unmarshals an instance of ApiGatewayOperationItemRespResult from the specified map of raw messages.
func UnmarshalApiGatewayOperationItemRespResult(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ApiGatewayOperationItemRespResult)
	err = core.UnmarshalPrimitive(m, "operation_id", &obj.OperationID)
	if err != nil {
		err = core.SDKErrorf(err, "", "operation_id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "method", &obj.Method)
	if err != nil {
		err = core.SDKErrorf(err, "", "method-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "host", &obj.Host)
	if err != nil {
		err = core.SDKErrorf(err, "", "host-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "endpoint", &obj.Endpoint)
	if err != nil {
		err = core.SDKErrorf(err, "", "endpoint-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ApiGatewayOperationsLabelsInputManaged : Managed labels to apply (e.g. cf-llm).
type ApiGatewayOperationsLabelsInputManaged struct {
	// Array of managed label strings.
	Labels []string `json:"labels,omitempty"`
}

// UnmarshalApiGatewayOperationsLabelsInputManaged unmarshals an instance of ApiGatewayOperationsLabelsInputManaged from the specified map of raw messages.
func UnmarshalApiGatewayOperationsLabelsInputManaged(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ApiGatewayOperationsLabelsInputManaged)
	err = core.UnmarshalPrimitive(m, "labels", &obj.Labels)
	if err != nil {
		err = core.SDKErrorf(err, "", "labels-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ApiGatewayOperationsLabelsInputSelector : Selector specifying which operations to label.
type ApiGatewayOperationsLabelsInputSelector struct {
	// Operations to include in the label operation.
	Include *ApiGatewayOperationsLabelsInputSelectorInclude `json:"include" validate:"required"`
}

// NewApiGatewayOperationsLabelsInputSelector : Instantiate ApiGatewayOperationsLabelsInputSelector (Generic Model Constructor)
func (*AiSecurityForAppsV1) NewApiGatewayOperationsLabelsInputSelector(include *ApiGatewayOperationsLabelsInputSelectorInclude) (_model *ApiGatewayOperationsLabelsInputSelector, err error) {
	_model = &ApiGatewayOperationsLabelsInputSelector{
		Include: include,
	}
	err = core.ValidateStruct(_model, "required parameters")
	if err != nil {
		err = core.SDKErrorf(err, "", "model-missing-required", common.GetComponentInfo())
	}
	return
}

// UnmarshalApiGatewayOperationsLabelsInputSelector unmarshals an instance of ApiGatewayOperationsLabelsInputSelector from the specified map of raw messages.
func UnmarshalApiGatewayOperationsLabelsInputSelector(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ApiGatewayOperationsLabelsInputSelector)
	err = core.UnmarshalModel(m, "include", &obj.Include, UnmarshalApiGatewayOperationsLabelsInputSelectorInclude)
	if err != nil {
		err = core.SDKErrorf(err, "", "include-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ApiGatewayOperationsLabelsInputSelectorInclude : Operations to include in the label operation.
type ApiGatewayOperationsLabelsInputSelectorInclude struct {
	// Array of operation UUIDs to label.
	OperationIds []string `json:"operation_ids,omitempty"`
}

// UnmarshalApiGatewayOperationsLabelsInputSelectorInclude unmarshals an instance of ApiGatewayOperationsLabelsInputSelectorInclude from the specified map of raw messages.
func UnmarshalApiGatewayOperationsLabelsInputSelectorInclude(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ApiGatewayOperationsLabelsInputSelectorInclude)
	err = core.UnmarshalPrimitive(m, "operation_ids", &obj.OperationIds)
	if err != nil {
		err = core.SDKErrorf(err, "", "operation_ids-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ApiGatewayOperationsLabelsInputUser : User-defined labels to apply.
type ApiGatewayOperationsLabelsInputUser struct {
	// Array of user-defined label strings.
	Labels []string `json:"labels,omitempty"`
}

// UnmarshalApiGatewayOperationsLabelsInputUser unmarshals an instance of ApiGatewayOperationsLabelsInputUser from the specified map of raw messages.
func UnmarshalApiGatewayOperationsLabelsInputUser(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ApiGatewayOperationsLabelsInputUser)
	err = core.UnmarshalPrimitive(m, "labels", &obj.Labels)
	if err != nil {
		err = core.SDKErrorf(err, "", "labels-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ApiGatewayOperationsLabelsRespResultItem : ApiGatewayOperationsLabelsRespResultItem struct
type ApiGatewayOperationsLabelsRespResultItem struct {
	OperationID *string `json:"operation_id,omitempty"`

	Labels []string `json:"labels,omitempty"`
}

// UnmarshalApiGatewayOperationsLabelsRespResultItem unmarshals an instance of ApiGatewayOperationsLabelsRespResultItem from the specified map of raw messages.
func UnmarshalApiGatewayOperationsLabelsRespResultItem(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ApiGatewayOperationsLabelsRespResultItem)
	err = core.UnmarshalPrimitive(m, "operation_id", &obj.OperationID)
	if err != nil {
		err = core.SDKErrorf(err, "", "operation_id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "labels", &obj.Labels)
	if err != nil {
		err = core.SDKErrorf(err, "", "labels-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ApiGatewayOperationsRespResultItem : ApiGatewayOperationsRespResultItem struct
type ApiGatewayOperationsRespResultItem struct {
	// UUID of the created operation.
	OperationID *string `json:"operation_id,omitempty"`

	Method *string `json:"method,omitempty"`

	Host *string `json:"host,omitempty"`

	Endpoint *string `json:"endpoint,omitempty"`
}

// UnmarshalApiGatewayOperationsRespResultItem unmarshals an instance of ApiGatewayOperationsRespResultItem from the specified map of raw messages.
func UnmarshalApiGatewayOperationsRespResultItem(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ApiGatewayOperationsRespResultItem)
	err = core.UnmarshalPrimitive(m, "operation_id", &obj.OperationID)
	if err != nil {
		err = core.SDKErrorf(err, "", "operation_id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "method", &obj.Method)
	if err != nil {
		err = core.SDKErrorf(err, "", "method-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "host", &obj.Host)
	if err != nil {
		err = core.SDKErrorf(err, "", "host-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "endpoint", &obj.Endpoint)
	if err != nil {
		err = core.SDKErrorf(err, "", "endpoint-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CreateApiGatewayOperationItemOptions : The CreateApiGatewayOperationItem options.
type CreateApiGatewayOperationItemOptions struct {
	// The HTTP method for the operation.
	Method *string `json:"method,omitempty"`

	// RFC3986-compliant host.
	Host *string `json:"host,omitempty"`

	// The endpoint path. Must start with /.
	Endpoint *string `json:"endpoint,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// Constants associated with the CreateApiGatewayOperationItemOptions.Method property.
// The HTTP method for the operation.
const (
	CreateApiGatewayOperationItemOptions_Method_Delete = "DELETE"
	CreateApiGatewayOperationItemOptions_Method_Get = "GET"
	CreateApiGatewayOperationItemOptions_Method_Head = "HEAD"
	CreateApiGatewayOperationItemOptions_Method_Options = "OPTIONS"
	CreateApiGatewayOperationItemOptions_Method_Patch = "PATCH"
	CreateApiGatewayOperationItemOptions_Method_Post = "POST"
	CreateApiGatewayOperationItemOptions_Method_Put = "PUT"
)

// NewCreateApiGatewayOperationItemOptions : Instantiate CreateApiGatewayOperationItemOptions
func (*AiSecurityForAppsV1) NewCreateApiGatewayOperationItemOptions() *CreateApiGatewayOperationItemOptions {
	return &CreateApiGatewayOperationItemOptions{}
}

// SetMethod : Allow user to set Method
func (_options *CreateApiGatewayOperationItemOptions) SetMethod(method string) *CreateApiGatewayOperationItemOptions {
	_options.Method = core.StringPtr(method)
	return _options
}

// SetHost : Allow user to set Host
func (_options *CreateApiGatewayOperationItemOptions) SetHost(host string) *CreateApiGatewayOperationItemOptions {
	_options.Host = core.StringPtr(host)
	return _options
}

// SetEndpoint : Allow user to set Endpoint
func (_options *CreateApiGatewayOperationItemOptions) SetEndpoint(endpoint string) *CreateApiGatewayOperationItemOptions {
	_options.Endpoint = core.StringPtr(endpoint)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateApiGatewayOperationItemOptions) SetHeaders(param map[string]string) *CreateApiGatewayOperationItemOptions {
	options.Headers = param
	return options
}

// CreateZoneApiGatewayOperationOptions : The CreateZoneApiGatewayOperation options.
type CreateZoneApiGatewayOperationOptions struct {
	// List of operations to create.
	ApiGatewayOperation []ApiGatewayOperation `json:"api_gateway_operation,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewCreateZoneApiGatewayOperationOptions : Instantiate CreateZoneApiGatewayOperationOptions
func (*AiSecurityForAppsV1) NewCreateZoneApiGatewayOperationOptions() *CreateZoneApiGatewayOperationOptions {
	return &CreateZoneApiGatewayOperationOptions{}
}

// SetApiGatewayOperation : Allow user to set ApiGatewayOperation
func (_options *CreateZoneApiGatewayOperationOptions) SetApiGatewayOperation(apiGatewayOperation []ApiGatewayOperation) *CreateZoneApiGatewayOperationOptions {
	_options.ApiGatewayOperation = apiGatewayOperation
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateZoneApiGatewayOperationOptions) SetHeaders(param map[string]string) *CreateZoneApiGatewayOperationOptions {
	options.Headers = param
	return options
}

// DeleteZoneApiGatewayOperationOptions : The DeleteZoneApiGatewayOperation options.
type DeleteZoneApiGatewayOperationOptions struct {
	// UUID of the API Gateway operation.
	OperationID *string `json:"operation_id" validate:"required,ne="`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewDeleteZoneApiGatewayOperationOptions : Instantiate DeleteZoneApiGatewayOperationOptions
func (*AiSecurityForAppsV1) NewDeleteZoneApiGatewayOperationOptions(operationID string) *DeleteZoneApiGatewayOperationOptions {
	return &DeleteZoneApiGatewayOperationOptions{
		OperationID: core.StringPtr(operationID),
	}
}

// SetOperationID : Allow user to set OperationID
func (_options *DeleteZoneApiGatewayOperationOptions) SetOperationID(operationID string) *DeleteZoneApiGatewayOperationOptions {
	_options.OperationID = core.StringPtr(operationID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteZoneApiGatewayOperationOptions) SetHeaders(param map[string]string) *DeleteZoneApiGatewayOperationOptions {
	options.Headers = param
	return options
}

// DiscoveryOperationFeatures : DiscoveryOperationFeatures struct
type DiscoveryOperationFeatures struct {
	TrafficStats *DiscoveryOperationFeaturesTrafficStats `json:"traffic_stats,omitempty"`
}

// UnmarshalDiscoveryOperationFeatures unmarshals an instance of DiscoveryOperationFeatures from the specified map of raw messages.
func UnmarshalDiscoveryOperationFeatures(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DiscoveryOperationFeatures)
	err = core.UnmarshalModel(m, "traffic_stats", &obj.TrafficStats, UnmarshalDiscoveryOperationFeaturesTrafficStats)
	if err != nil {
		err = core.SDKErrorf(err, "", "traffic_stats-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DiscoveryOperationFeaturesTrafficStats : DiscoveryOperationFeaturesTrafficStats struct
type DiscoveryOperationFeaturesTrafficStats struct {
	LastUpdated *strfmt.DateTime `json:"last_updated,omitempty"`

	// The period in seconds over which statistics were computed.
	PeriodSeconds *int64 `json:"period_seconds,omitempty"`

	// The average number of requests seen during this period.
	Requests *float64 `json:"requests,omitempty"`
}

// UnmarshalDiscoveryOperationFeaturesTrafficStats unmarshals an instance of DiscoveryOperationFeaturesTrafficStats from the specified map of raw messages.
func UnmarshalDiscoveryOperationFeaturesTrafficStats(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DiscoveryOperationFeaturesTrafficStats)
	err = core.UnmarshalPrimitive(m, "last_updated", &obj.LastUpdated)
	if err != nil {
		err = core.SDKErrorf(err, "", "last_updated-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "period_seconds", &obj.PeriodSeconds)
	if err != nil {
		err = core.SDKErrorf(err, "", "period_seconds-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "requests", &obj.Requests)
	if err != nil {
		err = core.SDKErrorf(err, "", "requests-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetAiSecuritySettingsOptions : The GetAiSecuritySettings options.
type GetAiSecuritySettingsOptions struct {

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewGetAiSecuritySettingsOptions : Instantiate GetAiSecuritySettingsOptions
func (*AiSecurityForAppsV1) NewGetAiSecuritySettingsOptions() *GetAiSecuritySettingsOptions {
	return &GetAiSecuritySettingsOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *GetAiSecuritySettingsOptions) SetHeaders(param map[string]string) *GetAiSecuritySettingsOptions {
	options.Headers = param
	return options
}

// GetApiGatewayDiscoveryOptions : The GetApiGatewayDiscovery options.
type GetApiGatewayDiscoveryOptions struct {

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewGetApiGatewayDiscoveryOptions : Instantiate GetApiGatewayDiscoveryOptions
func (*AiSecurityForAppsV1) NewGetApiGatewayDiscoveryOptions() *GetApiGatewayDiscoveryOptions {
	return &GetApiGatewayDiscoveryOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *GetApiGatewayDiscoveryOptions) SetHeaders(param map[string]string) *GetApiGatewayDiscoveryOptions {
	options.Headers = param
	return options
}

// GetApiGatewaySchemasOptions : The GetApiGatewaySchemas options.
type GetApiGatewaySchemasOptions struct {

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewGetApiGatewaySchemasOptions : Instantiate GetApiGatewaySchemasOptions
func (*AiSecurityForAppsV1) NewGetApiGatewaySchemasOptions() *GetApiGatewaySchemasOptions {
	return &GetApiGatewaySchemasOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *GetApiGatewaySchemasOptions) SetHeaders(param map[string]string) *GetApiGatewaySchemasOptions {
	options.Headers = param
	return options
}

// GetZoneApiGatewayOperationOptions : The GetZoneApiGatewayOperation options.
type GetZoneApiGatewayOperationOptions struct {
	// UUID of the API Gateway operation.
	OperationID *string `json:"operation_id" validate:"required,ne="`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewGetZoneApiGatewayOperationOptions : Instantiate GetZoneApiGatewayOperationOptions
func (*AiSecurityForAppsV1) NewGetZoneApiGatewayOperationOptions(operationID string) *GetZoneApiGatewayOperationOptions {
	return &GetZoneApiGatewayOperationOptions{
		OperationID: core.StringPtr(operationID),
	}
}

// SetOperationID : Allow user to set OperationID
func (_options *GetZoneApiGatewayOperationOptions) SetOperationID(operationID string) *GetZoneApiGatewayOperationOptions {
	_options.OperationID = core.StringPtr(operationID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetZoneApiGatewayOperationOptions) SetHeaders(param map[string]string) *GetZoneApiGatewayOperationOptions {
	options.Headers = param
	return options
}

// ListApiGatewayDiscoveryOperationsOptions : The ListApiGatewayDiscoveryOperations options.
type ListApiGatewayDiscoveryOperationsOptions struct {
	// When true, only return operations not yet saved into API Shield Endpoint Management.
	Diff *bool `json:"diff,omitempty"`

	// Direction to order results.
	Direction *string `json:"direction,omitempty"`

	// Filter results to only include endpoints containing this pattern.
	Endpoint *string `json:"endpoint,omitempty"`

	// Filter results to only include the specified hosts.
	Host []string `json:"host,omitempty"`

	// Filter results to only include the specified HTTP methods.
	Method []string `json:"method,omitempty"`

	// Field to order results by.
	Order *string `json:"order,omitempty"`

	// Filter by discovery engine source.
	Origin *string `json:"origin,omitempty"`

	// Filter results by discovery state (review/saved/ignored).
	State *string `json:"state,omitempty"`

	// Page number of paginated results.
	Page *int64 `json:"page,omitempty"`

	// Maximum number of results per page.
	PerPage *int64 `json:"per_page,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// Constants associated with the ListApiGatewayDiscoveryOperationsOptions.Direction property.
// Direction to order results.
const (
	ListApiGatewayDiscoveryOperationsOptions_Direction_Asc = "asc"
	ListApiGatewayDiscoveryOperationsOptions_Direction_Desc = "desc"
)

// Constants associated with the ListApiGatewayDiscoveryOperationsOptions.Order property.
// Field to order results by.
const (
	ListApiGatewayDiscoveryOperationsOptions_Order_Endpoint = "endpoint"
	ListApiGatewayDiscoveryOperationsOptions_Order_Host = "host"
	ListApiGatewayDiscoveryOperationsOptions_Order_Method = "method"
	ListApiGatewayDiscoveryOperationsOptions_Order_TrafficStatsLastUpdated = "traffic_stats.last_updated"
	ListApiGatewayDiscoveryOperationsOptions_Order_TrafficStatsRequests = "traffic_stats.requests"
)

// Constants associated with the ListApiGatewayDiscoveryOperationsOptions.Origin property.
// Filter by discovery engine source.
const (
	ListApiGatewayDiscoveryOperationsOptions_Origin_Labeldiscovery = "LabelDiscovery"
	ListApiGatewayDiscoveryOperationsOptions_Origin_Ml = "ML"
	ListApiGatewayDiscoveryOperationsOptions_Origin_Sessionidentifier = "SessionIdentifier"
)

// Constants associated with the ListApiGatewayDiscoveryOperationsOptions.State property.
// Filter results by discovery state (review/saved/ignored).
const (
	ListApiGatewayDiscoveryOperationsOptions_State_Ignored = "ignored"
	ListApiGatewayDiscoveryOperationsOptions_State_Review = "review"
	ListApiGatewayDiscoveryOperationsOptions_State_Saved = "saved"
)

// NewListApiGatewayDiscoveryOperationsOptions : Instantiate ListApiGatewayDiscoveryOperationsOptions
func (*AiSecurityForAppsV1) NewListApiGatewayDiscoveryOperationsOptions() *ListApiGatewayDiscoveryOperationsOptions {
	return &ListApiGatewayDiscoveryOperationsOptions{}
}

// SetDiff : Allow user to set Diff
func (_options *ListApiGatewayDiscoveryOperationsOptions) SetDiff(diff bool) *ListApiGatewayDiscoveryOperationsOptions {
	_options.Diff = core.BoolPtr(diff)
	return _options
}

// SetDirection : Allow user to set Direction
func (_options *ListApiGatewayDiscoveryOperationsOptions) SetDirection(direction string) *ListApiGatewayDiscoveryOperationsOptions {
	_options.Direction = core.StringPtr(direction)
	return _options
}

// SetEndpoint : Allow user to set Endpoint
func (_options *ListApiGatewayDiscoveryOperationsOptions) SetEndpoint(endpoint string) *ListApiGatewayDiscoveryOperationsOptions {
	_options.Endpoint = core.StringPtr(endpoint)
	return _options
}

// SetHost : Allow user to set Host
func (_options *ListApiGatewayDiscoveryOperationsOptions) SetHost(host []string) *ListApiGatewayDiscoveryOperationsOptions {
	_options.Host = host
	return _options
}

// SetMethod : Allow user to set Method
func (_options *ListApiGatewayDiscoveryOperationsOptions) SetMethod(method []string) *ListApiGatewayDiscoveryOperationsOptions {
	_options.Method = method
	return _options
}

// SetOrder : Allow user to set Order
func (_options *ListApiGatewayDiscoveryOperationsOptions) SetOrder(order string) *ListApiGatewayDiscoveryOperationsOptions {
	_options.Order = core.StringPtr(order)
	return _options
}

// SetOrigin : Allow user to set Origin
func (_options *ListApiGatewayDiscoveryOperationsOptions) SetOrigin(origin string) *ListApiGatewayDiscoveryOperationsOptions {
	_options.Origin = core.StringPtr(origin)
	return _options
}

// SetState : Allow user to set State
func (_options *ListApiGatewayDiscoveryOperationsOptions) SetState(state string) *ListApiGatewayDiscoveryOperationsOptions {
	_options.State = core.StringPtr(state)
	return _options
}

// SetPage : Allow user to set Page
func (_options *ListApiGatewayDiscoveryOperationsOptions) SetPage(page int64) *ListApiGatewayDiscoveryOperationsOptions {
	_options.Page = core.Int64Ptr(page)
	return _options
}

// SetPerPage : Allow user to set PerPage
func (_options *ListApiGatewayDiscoveryOperationsOptions) SetPerPage(perPage int64) *ListApiGatewayDiscoveryOperationsOptions {
	_options.PerPage = core.Int64Ptr(perPage)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListApiGatewayDiscoveryOperationsOptions) SetHeaders(param map[string]string) *ListApiGatewayDiscoveryOperationsOptions {
	options.Headers = param
	return options
}

// ReplaceZoneAiSecuritySettingsOptions : The ReplaceZoneAiSecuritySettings options.
type ReplaceZoneAiSecuritySettingsOptions struct {
	// Set to true to enable AI Security for Apps, false to disable.
	Enabled *bool `json:"enabled,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewReplaceZoneAiSecuritySettingsOptions : Instantiate ReplaceZoneAiSecuritySettingsOptions
func (*AiSecurityForAppsV1) NewReplaceZoneAiSecuritySettingsOptions() *ReplaceZoneAiSecuritySettingsOptions {
	return &ReplaceZoneAiSecuritySettingsOptions{}
}

// SetEnabled : Allow user to set Enabled
func (_options *ReplaceZoneAiSecuritySettingsOptions) SetEnabled(enabled bool) *ReplaceZoneAiSecuritySettingsOptions {
	_options.Enabled = core.BoolPtr(enabled)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ReplaceZoneAiSecuritySettingsOptions) SetHeaders(param map[string]string) *ReplaceZoneAiSecuritySettingsOptions {
	options.Headers = param
	return options
}

// UpdateApiGatewayOperationLabelsOptions : The UpdateApiGatewayOperationLabels options.
type UpdateApiGatewayOperationLabelsOptions struct {
	// User-defined labels to apply.
	User *ApiGatewayOperationsLabelsInputUser `json:"user,omitempty"`

	// Managed labels to apply (e.g. cf-llm).
	Managed *ApiGatewayOperationsLabelsInputManaged `json:"managed,omitempty"`

	// Selector specifying which operations to label.
	Selector *ApiGatewayOperationsLabelsInputSelector `json:"selector,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewUpdateApiGatewayOperationLabelsOptions : Instantiate UpdateApiGatewayOperationLabelsOptions
func (*AiSecurityForAppsV1) NewUpdateApiGatewayOperationLabelsOptions() *UpdateApiGatewayOperationLabelsOptions {
	return &UpdateApiGatewayOperationLabelsOptions{}
}

// SetUser : Allow user to set User
func (_options *UpdateApiGatewayOperationLabelsOptions) SetUser(user *ApiGatewayOperationsLabelsInputUser) *UpdateApiGatewayOperationLabelsOptions {
	_options.User = user
	return _options
}

// SetManaged : Allow user to set Managed
func (_options *UpdateApiGatewayOperationLabelsOptions) SetManaged(managed *ApiGatewayOperationsLabelsInputManaged) *UpdateApiGatewayOperationLabelsOptions {
	_options.Managed = managed
	return _options
}

// SetSelector : Allow user to set Selector
func (_options *UpdateApiGatewayOperationLabelsOptions) SetSelector(selector *ApiGatewayOperationsLabelsInputSelector) *UpdateApiGatewayOperationLabelsOptions {
	_options.Selector = selector
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateApiGatewayOperationLabelsOptions) SetHeaders(param map[string]string) *UpdateApiGatewayOperationLabelsOptions {
	options.Headers = param
	return options
}

// UpdateZoneApiGatewayDiscoveryOperationOptions : The UpdateZoneApiGatewayDiscoveryOperation options.
type UpdateZoneApiGatewayDiscoveryOperationOptions struct {
	// List of operation state updates.
	RequestBody map[string]interface{} `json:"request_body,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewUpdateZoneApiGatewayDiscoveryOperationOptions : Instantiate UpdateZoneApiGatewayDiscoveryOperationOptions
func (*AiSecurityForAppsV1) NewUpdateZoneApiGatewayDiscoveryOperationOptions() *UpdateZoneApiGatewayDiscoveryOperationOptions {
	return &UpdateZoneApiGatewayDiscoveryOperationOptions{}
}

// SetRequestBody : Allow user to set RequestBody
func (_options *UpdateZoneApiGatewayDiscoveryOperationOptions) SetRequestBody(requestBody map[string]interface{}) *UpdateZoneApiGatewayDiscoveryOperationOptions {
	_options.RequestBody = requestBody
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateZoneApiGatewayDiscoveryOperationOptions) SetHeaders(param map[string]string) *UpdateZoneApiGatewayDiscoveryOperationOptions {
	options.Headers = param
	return options
}

// AiSecuritySettingsResp : AI Security for Apps settings response.
type AiSecuritySettingsResp struct {
	// Was operation successful.
	Success *bool `json:"success" validate:"required"`

	// Array of errors encountered.
	Errors [][]string `json:"errors" validate:"required"`

	// Array of messages returned.
	Messages [][]string `json:"messages" validate:"required"`

	// Container for response information.
	Result *AiSecuritySettingsRespResult `json:"result" validate:"required"`
}

// UnmarshalAiSecuritySettingsResp unmarshals an instance of AiSecuritySettingsResp from the specified map of raw messages.
func UnmarshalAiSecuritySettingsResp(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AiSecuritySettingsResp)
	err = core.UnmarshalPrimitive(m, "success", &obj.Success)
	if err != nil {
		err = core.SDKErrorf(err, "", "success-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "errors", &obj.Errors)
	if err != nil {
		err = core.SDKErrorf(err, "", "errors-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "messages", &obj.Messages)
	if err != nil {
		err = core.SDKErrorf(err, "", "messages-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "result", &obj.Result, UnmarshalAiSecuritySettingsRespResult)
	if err != nil {
		err = core.SDKErrorf(err, "", "result-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ApiGatewayDiscoveryResp : API Gateway discovery response (OpenAPI schema format).
type ApiGatewayDiscoveryResp struct {
	// Was operation successful.
	Success *bool `json:"success" validate:"required"`

	// Array of errors encountered.
	Errors [][]string `json:"errors" validate:"required"`

	// Array of messages returned.
	Messages [][]string `json:"messages" validate:"required"`

	// Discovered operations rendered as an OpenAPI schema document.
	Result map[string]interface{} `json:"result" validate:"required"`
}

// UnmarshalApiGatewayDiscoveryResp unmarshals an instance of ApiGatewayDiscoveryResp from the specified map of raw messages.
func UnmarshalApiGatewayDiscoveryResp(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ApiGatewayDiscoveryResp)
	err = core.UnmarshalPrimitive(m, "success", &obj.Success)
	if err != nil {
		err = core.SDKErrorf(err, "", "success-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "errors", &obj.Errors)
	if err != nil {
		err = core.SDKErrorf(err, "", "errors-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "messages", &obj.Messages)
	if err != nil {
		err = core.SDKErrorf(err, "", "messages-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "result", &obj.Result)
	if err != nil {
		err = core.SDKErrorf(err, "", "result-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ApiGatewayOperation : An API Gateway operation definition.
type ApiGatewayOperation struct {
	// The HTTP method for the operation.
	Method *string `json:"method" validate:"required"`

	// RFC3986-compliant host.
	Host *string `json:"host" validate:"required"`

	// The endpoint path. Must start with /.
	Endpoint *string `json:"endpoint" validate:"required"`
}

// Constants associated with the ApiGatewayOperation.Method property.
// The HTTP method for the operation.
const (
	ApiGatewayOperation_Method_Delete = "DELETE"
	ApiGatewayOperation_Method_Get = "GET"
	ApiGatewayOperation_Method_Head = "HEAD"
	ApiGatewayOperation_Method_Options = "OPTIONS"
	ApiGatewayOperation_Method_Patch = "PATCH"
	ApiGatewayOperation_Method_Post = "POST"
	ApiGatewayOperation_Method_Put = "PUT"
)

// NewApiGatewayOperation : Instantiate ApiGatewayOperation (Generic Model Constructor)
func (*AiSecurityForAppsV1) NewApiGatewayOperation(method string, host string, endpoint string) (_model *ApiGatewayOperation, err error) {
	_model = &ApiGatewayOperation{
		Method: core.StringPtr(method),
		Host: core.StringPtr(host),
		Endpoint: core.StringPtr(endpoint),
	}
	err = core.ValidateStruct(_model, "required parameters")
	if err != nil {
		err = core.SDKErrorf(err, "", "model-missing-required", common.GetComponentInfo())
	}
	return
}

// UnmarshalApiGatewayOperation unmarshals an instance of ApiGatewayOperation from the specified map of raw messages.
func UnmarshalApiGatewayOperation(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ApiGatewayOperation)
	err = core.UnmarshalPrimitive(m, "method", &obj.Method)
	if err != nil {
		err = core.SDKErrorf(err, "", "method-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "host", &obj.Host)
	if err != nil {
		err = core.SDKErrorf(err, "", "host-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "endpoint", &obj.Endpoint)
	if err != nil {
		err = core.SDKErrorf(err, "", "endpoint-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ApiGatewayOperationItemResp : Single API Gateway operation create response.
type ApiGatewayOperationItemResp struct {
	// Was operation successful.
	Success *bool `json:"success" validate:"required"`

	// Array of errors encountered.
	Errors [][]string `json:"errors" validate:"required"`

	// Array of messages returned.
	Messages [][]string `json:"messages" validate:"required"`

	Result *ApiGatewayOperationItemRespResult `json:"result" validate:"required"`
}

// UnmarshalApiGatewayOperationItemResp unmarshals an instance of ApiGatewayOperationItemResp from the specified map of raw messages.
func UnmarshalApiGatewayOperationItemResp(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ApiGatewayOperationItemResp)
	err = core.UnmarshalPrimitive(m, "success", &obj.Success)
	if err != nil {
		err = core.SDKErrorf(err, "", "success-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "errors", &obj.Errors)
	if err != nil {
		err = core.SDKErrorf(err, "", "errors-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "messages", &obj.Messages)
	if err != nil {
		err = core.SDKErrorf(err, "", "messages-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "result", &obj.Result, UnmarshalApiGatewayOperationItemRespResult)
	if err != nil {
		err = core.SDKErrorf(err, "", "result-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ApiGatewayOperationsLabelsResp : API Gateway operations labels update response.
type ApiGatewayOperationsLabelsResp struct {
	// Was operation successful.
	Success *bool `json:"success" validate:"required"`

	// Array of errors encountered.
	Errors [][]string `json:"errors" validate:"required"`

	// Array of messages returned.
	Messages [][]string `json:"messages" validate:"required"`

	// List of operations with their updated label sets.
	Result []ApiGatewayOperationsLabelsRespResultItem `json:"result" validate:"required"`
}

// UnmarshalApiGatewayOperationsLabelsResp unmarshals an instance of ApiGatewayOperationsLabelsResp from the specified map of raw messages.
func UnmarshalApiGatewayOperationsLabelsResp(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ApiGatewayOperationsLabelsResp)
	err = core.UnmarshalPrimitive(m, "success", &obj.Success)
	if err != nil {
		err = core.SDKErrorf(err, "", "success-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "errors", &obj.Errors)
	if err != nil {
		err = core.SDKErrorf(err, "", "errors-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "messages", &obj.Messages)
	if err != nil {
		err = core.SDKErrorf(err, "", "messages-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "result", &obj.Result, UnmarshalApiGatewayOperationsLabelsRespResultItem)
	if err != nil {
		err = core.SDKErrorf(err, "", "result-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ApiGatewayOperationsResp : API Gateway bulk operations create response.
type ApiGatewayOperationsResp struct {
	// Was operation successful.
	Success *bool `json:"success" validate:"required"`

	// Array of errors encountered.
	Errors [][]string `json:"errors" validate:"required"`

	// Array of messages returned.
	Messages [][]string `json:"messages" validate:"required"`

	// List of created operations.
	Result []ApiGatewayOperationsRespResultItem `json:"result" validate:"required"`
}

// UnmarshalApiGatewayOperationsResp unmarshals an instance of ApiGatewayOperationsResp from the specified map of raw messages.
func UnmarshalApiGatewayOperationsResp(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ApiGatewayOperationsResp)
	err = core.UnmarshalPrimitive(m, "success", &obj.Success)
	if err != nil {
		err = core.SDKErrorf(err, "", "success-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "errors", &obj.Errors)
	if err != nil {
		err = core.SDKErrorf(err, "", "errors-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "messages", &obj.Messages)
	if err != nil {
		err = core.SDKErrorf(err, "", "messages-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "result", &obj.Result, UnmarshalApiGatewayOperationsRespResultItem)
	if err != nil {
		err = core.SDKErrorf(err, "", "result-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ApiGatewaySchemasResp : API Gateway schemas response (OpenAPI schema format).
type ApiGatewaySchemasResp struct {
	// Was operation successful.
	Success *bool `json:"success" validate:"required"`

	// Array of errors encountered.
	Errors [][]string `json:"errors" validate:"required"`

	// Array of messages returned.
	Messages [][]string `json:"messages" validate:"required"`

	// API Gateway schemas rendered as an OpenAPI schema document.
	Result map[string]interface{} `json:"result" validate:"required"`
}

// UnmarshalApiGatewaySchemasResp unmarshals an instance of ApiGatewaySchemasResp from the specified map of raw messages.
func UnmarshalApiGatewaySchemasResp(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ApiGatewaySchemasResp)
	err = core.UnmarshalPrimitive(m, "success", &obj.Success)
	if err != nil {
		err = core.SDKErrorf(err, "", "success-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "errors", &obj.Errors)
	if err != nil {
		err = core.SDKErrorf(err, "", "errors-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "messages", &obj.Messages)
	if err != nil {
		err = core.SDKErrorf(err, "", "messages-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "result", &obj.Result)
	if err != nil {
		err = core.SDKErrorf(err, "", "result-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DiscoveryOperation : A discovered API operation.
type DiscoveryOperation struct {
	// UUID of the discovered operation.
	ID *string `json:"id,omitempty"`

	// The endpoint path. May contain path parameter templates in curly braces (e.g. /api/user/{var1}/details).
	Endpoint *string `json:"endpoint,omitempty"`

	// RFC3986-compliant host.
	Host *string `json:"host,omitempty"`

	// The HTTP method used to access the endpoint.
	Method *string `json:"method,omitempty"`

	LastUpdated *strfmt.DateTime `json:"last_updated,omitempty"`

	// API discovery engine(s) that discovered this operation.
	Origin []string `json:"origin,omitempty"`

	// State of the operation in API Discovery. review - not yet saved to Endpoint Management; saved - saved to Endpoint
	// Management; ignored - marked as ignored.
	State *string `json:"state,omitempty"`

	Features *DiscoveryOperationFeatures `json:"features,omitempty"`
}

// Constants associated with the DiscoveryOperation.Method property.
// The HTTP method used to access the endpoint.
const (
	DiscoveryOperation_Method_Connect = "CONNECT"
	DiscoveryOperation_Method_Delete = "DELETE"
	DiscoveryOperation_Method_Get = "GET"
	DiscoveryOperation_Method_Head = "HEAD"
	DiscoveryOperation_Method_Options = "OPTIONS"
	DiscoveryOperation_Method_Patch = "PATCH"
	DiscoveryOperation_Method_Post = "POST"
	DiscoveryOperation_Method_Put = "PUT"
	DiscoveryOperation_Method_Trace = "TRACE"
)

// Constants associated with the DiscoveryOperation.Origin property.
const (
	DiscoveryOperation_Origin_Labeldiscovery = "LabelDiscovery"
	DiscoveryOperation_Origin_Ml = "ML"
	DiscoveryOperation_Origin_Sessionidentifier = "SessionIdentifier"
)

// Constants associated with the DiscoveryOperation.State property.
// State of the operation in API Discovery. review - not yet saved to Endpoint Management; saved - saved to Endpoint
// Management; ignored - marked as ignored.
const (
	DiscoveryOperation_State_Ignored = "ignored"
	DiscoveryOperation_State_Review = "review"
	DiscoveryOperation_State_Saved = "saved"
)

// UnmarshalDiscoveryOperation unmarshals an instance of DiscoveryOperation from the specified map of raw messages.
func UnmarshalDiscoveryOperation(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DiscoveryOperation)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		err = core.SDKErrorf(err, "", "id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "endpoint", &obj.Endpoint)
	if err != nil {
		err = core.SDKErrorf(err, "", "endpoint-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "host", &obj.Host)
	if err != nil {
		err = core.SDKErrorf(err, "", "host-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "method", &obj.Method)
	if err != nil {
		err = core.SDKErrorf(err, "", "method-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "last_updated", &obj.LastUpdated)
	if err != nil {
		err = core.SDKErrorf(err, "", "last_updated-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "origin", &obj.Origin)
	if err != nil {
		err = core.SDKErrorf(err, "", "origin-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "state", &obj.State)
	if err != nil {
		err = core.SDKErrorf(err, "", "state-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "features", &obj.Features, UnmarshalDiscoveryOperationFeatures)
	if err != nil {
		err = core.SDKErrorf(err, "", "features-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DiscoveryOperationsListResp : API Gateway discovery operations list response.
type DiscoveryOperationsListResp struct {
	// Was operation successful.
	Success *bool `json:"success" validate:"required"`

	// Array of errors encountered.
	Errors [][]string `json:"errors" validate:"required"`

	// Array of messages returned.
	Messages [][]string `json:"messages" validate:"required"`

	Result []DiscoveryOperation `json:"result" validate:"required"`

	ResultInfo *ResultInfo `json:"result_info,omitempty"`
}

// UnmarshalDiscoveryOperationsListResp unmarshals an instance of DiscoveryOperationsListResp from the specified map of raw messages.
func UnmarshalDiscoveryOperationsListResp(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DiscoveryOperationsListResp)
	err = core.UnmarshalPrimitive(m, "success", &obj.Success)
	if err != nil {
		err = core.SDKErrorf(err, "", "success-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "errors", &obj.Errors)
	if err != nil {
		err = core.SDKErrorf(err, "", "errors-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "messages", &obj.Messages)
	if err != nil {
		err = core.SDKErrorf(err, "", "messages-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "result", &obj.Result, UnmarshalDiscoveryOperation)
	if err != nil {
		err = core.SDKErrorf(err, "", "result-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "result_info", &obj.ResultInfo, UnmarshalResultInfo)
	if err != nil {
		err = core.SDKErrorf(err, "", "result_info-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DiscoveryOperationsPatchResp : API Gateway discovery operations bulk patch response.
type DiscoveryOperationsPatchResp struct {
	// Was operation successful.
	Success *bool `json:"success" validate:"required"`

	// Array of errors encountered.
	Errors [][]string `json:"errors" validate:"required"`

	// Array of messages returned.
	Messages [][]string `json:"messages" validate:"required"`

	Result []DiscoveryOperation `json:"result" validate:"required"`
}

// UnmarshalDiscoveryOperationsPatchResp unmarshals an instance of DiscoveryOperationsPatchResp from the specified map of raw messages.
func UnmarshalDiscoveryOperationsPatchResp(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DiscoveryOperationsPatchResp)
	err = core.UnmarshalPrimitive(m, "success", &obj.Success)
	if err != nil {
		err = core.SDKErrorf(err, "", "success-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "errors", &obj.Errors)
	if err != nil {
		err = core.SDKErrorf(err, "", "errors-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "messages", &obj.Messages)
	if err != nil {
		err = core.SDKErrorf(err, "", "messages-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "result", &obj.Result, UnmarshalDiscoveryOperation)
	if err != nil {
		err = core.SDKErrorf(err, "", "result-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ResultInfo : ResultInfo struct
type ResultInfo struct {
	// Total number of results for the requested service.
	Count *int64 `json:"count,omitempty"`

	// Current page within paginated list of results.
	Page *int64 `json:"page,omitempty"`

	// Number of results per page.
	PerPage *int64 `json:"per_page,omitempty"`

	// Total number of results.
	TotalCount *int64 `json:"total_count,omitempty"`
}

// UnmarshalResultInfo unmarshals an instance of ResultInfo from the specified map of raw messages.
func UnmarshalResultInfo(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ResultInfo)
	err = core.UnmarshalPrimitive(m, "count", &obj.Count)
	if err != nil {
		err = core.SDKErrorf(err, "", "count-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "page", &obj.Page)
	if err != nil {
		err = core.SDKErrorf(err, "", "page-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "per_page", &obj.PerPage)
	if err != nil {
		err = core.SDKErrorf(err, "", "per_page-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "total_count", &obj.TotalCount)
	if err != nil {
		err = core.SDKErrorf(err, "", "total_count-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}
