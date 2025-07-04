/**
 * (C) Copyright IBM Corp. 2025.
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
 * IBM OpenAPI SDK Code Generator Version: 3.98.0-8be2046a-20241205-162752
 */

// Package zonessettingsv1 : Operations and models for the ZonesSettingsV1 service
package zonessettingsv1

import (
	"context"
	"encoding/json"
	"net/http"
	"reflect"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	common "github.com/IBM/networking-go-sdk/common"
	"github.com/go-openapi/strfmt"
)

// ZonesSettingsV1 : CIS Zones Settings
//
// API Version: 1.0.1
type ZonesSettingsV1 struct {
	Service *core.BaseService

	// Full url-encoded cloud resource name (CRN) of resource instance.
	Crn *string

	// Zone identifier.
	ZoneIdentifier *string
}

// DefaultServiceURL is the default URL to make service requests to.
const DefaultServiceURL = "https://api.cis.cloud.ibm.com"

// DefaultServiceName is the default key used to find external configuration information.
const DefaultServiceName = "zones_settings"

// ZonesSettingsV1Options : Service options
type ZonesSettingsV1Options struct {
	ServiceName   string
	URL           string
	Authenticator core.Authenticator

	// Full url-encoded cloud resource name (CRN) of resource instance.
	Crn *string `validate:"required"`

	// Zone identifier.
	ZoneIdentifier *string `validate:"required"`
}

// NewZonesSettingsV1UsingExternalConfig : constructs an instance of ZonesSettingsV1 with passed in options and external configuration.
func NewZonesSettingsV1UsingExternalConfig(options *ZonesSettingsV1Options) (zonesSettings *ZonesSettingsV1, err error) {
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

	zonesSettings, err = NewZonesSettingsV1(options)
	err = core.RepurposeSDKProblem(err, "new-client-error")
	if err != nil {
		return
	}

	err = zonesSettings.Service.ConfigureService(options.ServiceName)
	if err != nil {
		err = core.SDKErrorf(err, "", "client-config-error", common.GetComponentInfo())
		return
	}

	if options.URL != "" {
		err = zonesSettings.Service.SetServiceURL(options.URL)
		err = core.RepurposeSDKProblem(err, "url-set-error")
	}
	return
}

// NewZonesSettingsV1 : constructs an instance of ZonesSettingsV1 with passed in options.
func NewZonesSettingsV1(options *ZonesSettingsV1Options) (service *ZonesSettingsV1, err error) {
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

	service = &ZonesSettingsV1{
		Service:        baseService,
		Crn:            options.Crn,
		ZoneIdentifier: options.ZoneIdentifier,
	}

	return
}

// GetServiceURLForRegion returns the service URL to be used for the specified region
func GetServiceURLForRegion(region string) (string, error) {
	return "", core.SDKErrorf(nil, "service does not support regional URLs", "no-regional-support", common.GetComponentInfo())
}

// Clone makes a copy of "zonesSettings" suitable for processing requests.
func (zonesSettings *ZonesSettingsV1) Clone() *ZonesSettingsV1 {
	if core.IsNil(zonesSettings) {
		return nil
	}
	clone := *zonesSettings
	clone.Service = zonesSettings.Service.Clone()
	return &clone
}

// SetServiceURL sets the service URL
func (zonesSettings *ZonesSettingsV1) SetServiceURL(url string) error {
	err := zonesSettings.Service.SetServiceURL(url)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-set-error", common.GetComponentInfo())
	}
	return err
}

// GetServiceURL returns the service URL
func (zonesSettings *ZonesSettingsV1) GetServiceURL() string {
	return zonesSettings.Service.GetServiceURL()
}

// SetDefaultHeaders sets HTTP headers to be sent in every request
func (zonesSettings *ZonesSettingsV1) SetDefaultHeaders(headers http.Header) {
	zonesSettings.Service.SetDefaultHeaders(headers)
}

// SetEnableGzipCompression sets the service's EnableGzipCompression field
func (zonesSettings *ZonesSettingsV1) SetEnableGzipCompression(enableGzip bool) {
	zonesSettings.Service.SetEnableGzipCompression(enableGzip)
}

// GetEnableGzipCompression returns the service's EnableGzipCompression field
func (zonesSettings *ZonesSettingsV1) GetEnableGzipCompression() bool {
	return zonesSettings.Service.GetEnableGzipCompression()
}

// EnableRetries enables automatic retries for requests invoked for this service instance.
// If either parameter is specified as 0, then a default value is used instead.
func (zonesSettings *ZonesSettingsV1) EnableRetries(maxRetries int, maxRetryInterval time.Duration) {
	zonesSettings.Service.EnableRetries(maxRetries, maxRetryInterval)
}

// DisableRetries disables automatic retries for requests invoked for this service instance.
func (zonesSettings *ZonesSettingsV1) DisableRetries() {
	zonesSettings.Service.DisableRetries()
}

// GetZoneDnssec : Get zone DNSSEC
// Get DNSSEC setting for a given zone.
func (zonesSettings *ZonesSettingsV1) GetZoneDnssec(getZoneDnssecOptions *GetZoneDnssecOptions) (result *ZonesDnssecResp, response *core.DetailedResponse, err error) {
	result, response, err = zonesSettings.GetZoneDnssecWithContext(context.Background(), getZoneDnssecOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// GetZoneDnssecWithContext is an alternate form of the GetZoneDnssec method which supports a Context parameter
func (zonesSettings *ZonesSettingsV1) GetZoneDnssecWithContext(ctx context.Context, getZoneDnssecOptions *GetZoneDnssecOptions) (result *ZonesDnssecResp, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getZoneDnssecOptions, "getZoneDnssecOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"crn":             *zonesSettings.Crn,
		"zone_identifier": *zonesSettings.ZoneIdentifier,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = zonesSettings.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(zonesSettings.Service.Options.URL, `/v1/{crn}/zones/{zone_identifier}/dnssec`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range getZoneDnssecOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("zones_settings", "V1", "GetZoneDnssec")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = zonesSettings.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "get_zone_dnssec", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalZonesDnssecResp)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// UpdateZoneDnssec : Update zone DNSSEC
// Update DNSSEC setting for given zone.
func (zonesSettings *ZonesSettingsV1) UpdateZoneDnssec(updateZoneDnssecOptions *UpdateZoneDnssecOptions) (result *ZonesDnssecResp, response *core.DetailedResponse, err error) {
	result, response, err = zonesSettings.UpdateZoneDnssecWithContext(context.Background(), updateZoneDnssecOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// UpdateZoneDnssecWithContext is an alternate form of the UpdateZoneDnssec method which supports a Context parameter
func (zonesSettings *ZonesSettingsV1) UpdateZoneDnssecWithContext(ctx context.Context, updateZoneDnssecOptions *UpdateZoneDnssecOptions) (result *ZonesDnssecResp, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(updateZoneDnssecOptions, "updateZoneDnssecOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"crn":             *zonesSettings.Crn,
		"zone_identifier": *zonesSettings.ZoneIdentifier,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = zonesSettings.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(zonesSettings.Service.Options.URL, `/v1/{crn}/zones/{zone_identifier}/dnssec`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range updateZoneDnssecOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("zones_settings", "V1", "UpdateZoneDnssec")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if updateZoneDnssecOptions.Status != nil {
		body["status"] = updateZoneDnssecOptions.Status
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
	response, err = zonesSettings.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "update_zone_dnssec", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalZonesDnssecResp)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// GetZoneCnameFlattening : Get zone CNAME flattening
// Get CNAME flattening setting for a given zone.
func (zonesSettings *ZonesSettingsV1) GetZoneCnameFlattening(getZoneCnameFlatteningOptions *GetZoneCnameFlatteningOptions) (result *ZonesCnameFlatteningResp, response *core.DetailedResponse, err error) {
	result, response, err = zonesSettings.GetZoneCnameFlatteningWithContext(context.Background(), getZoneCnameFlatteningOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// GetZoneCnameFlatteningWithContext is an alternate form of the GetZoneCnameFlattening method which supports a Context parameter
func (zonesSettings *ZonesSettingsV1) GetZoneCnameFlatteningWithContext(ctx context.Context, getZoneCnameFlatteningOptions *GetZoneCnameFlatteningOptions) (result *ZonesCnameFlatteningResp, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getZoneCnameFlatteningOptions, "getZoneCnameFlatteningOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"crn":             *zonesSettings.Crn,
		"zone_identifier": *zonesSettings.ZoneIdentifier,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = zonesSettings.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(zonesSettings.Service.Options.URL, `/v1/{crn}/zones/{zone_identifier}/settings/cname_flattening`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range getZoneCnameFlatteningOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("zones_settings", "V1", "GetZoneCnameFlattening")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = zonesSettings.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "get_zone_cname_flattening", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalZonesCnameFlatteningResp)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// UpdateZoneCnameFlattening : Update zone CNAME flattening
// Update CNAME flattening setting for given zone.
func (zonesSettings *ZonesSettingsV1) UpdateZoneCnameFlattening(updateZoneCnameFlatteningOptions *UpdateZoneCnameFlatteningOptions) (result *ZonesCnameFlatteningResp, response *core.DetailedResponse, err error) {
	result, response, err = zonesSettings.UpdateZoneCnameFlatteningWithContext(context.Background(), updateZoneCnameFlatteningOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// UpdateZoneCnameFlatteningWithContext is an alternate form of the UpdateZoneCnameFlattening method which supports a Context parameter
func (zonesSettings *ZonesSettingsV1) UpdateZoneCnameFlatteningWithContext(ctx context.Context, updateZoneCnameFlatteningOptions *UpdateZoneCnameFlatteningOptions) (result *ZonesCnameFlatteningResp, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(updateZoneCnameFlatteningOptions, "updateZoneCnameFlatteningOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"crn":             *zonesSettings.Crn,
		"zone_identifier": *zonesSettings.ZoneIdentifier,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = zonesSettings.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(zonesSettings.Service.Options.URL, `/v1/{crn}/zones/{zone_identifier}/settings/cname_flattening`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range updateZoneCnameFlatteningOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("zones_settings", "V1", "UpdateZoneCnameFlattening")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if updateZoneCnameFlatteningOptions.Value != nil {
		body["value"] = updateZoneCnameFlatteningOptions.Value
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
	response, err = zonesSettings.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "update_zone_cname_flattening", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalZonesCnameFlatteningResp)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// GetOpportunisticEncryption : Get opportunistic encryption setting
// Get opportunistic encryption setting for a zone.
func (zonesSettings *ZonesSettingsV1) GetOpportunisticEncryption(getOpportunisticEncryptionOptions *GetOpportunisticEncryptionOptions) (result *OpportunisticEncryptionResp, response *core.DetailedResponse, err error) {
	result, response, err = zonesSettings.GetOpportunisticEncryptionWithContext(context.Background(), getOpportunisticEncryptionOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// GetOpportunisticEncryptionWithContext is an alternate form of the GetOpportunisticEncryption method which supports a Context parameter
func (zonesSettings *ZonesSettingsV1) GetOpportunisticEncryptionWithContext(ctx context.Context, getOpportunisticEncryptionOptions *GetOpportunisticEncryptionOptions) (result *OpportunisticEncryptionResp, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getOpportunisticEncryptionOptions, "getOpportunisticEncryptionOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"crn":             *zonesSettings.Crn,
		"zone_identifier": *zonesSettings.ZoneIdentifier,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = zonesSettings.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(zonesSettings.Service.Options.URL, `/v1/{crn}/zones/{zone_identifier}/settings/opportunistic_encryption`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range getOpportunisticEncryptionOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("zones_settings", "V1", "GetOpportunisticEncryption")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = zonesSettings.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "get_opportunistic_encryption", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalOpportunisticEncryptionResp)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// UpdateOpportunisticEncryption : Update opportunistic encryption setting
// Update opportunistic encryption setting for a zone.
func (zonesSettings *ZonesSettingsV1) UpdateOpportunisticEncryption(updateOpportunisticEncryptionOptions *UpdateOpportunisticEncryptionOptions) (result *OpportunisticEncryptionResp, response *core.DetailedResponse, err error) {
	result, response, err = zonesSettings.UpdateOpportunisticEncryptionWithContext(context.Background(), updateOpportunisticEncryptionOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// UpdateOpportunisticEncryptionWithContext is an alternate form of the UpdateOpportunisticEncryption method which supports a Context parameter
func (zonesSettings *ZonesSettingsV1) UpdateOpportunisticEncryptionWithContext(ctx context.Context, updateOpportunisticEncryptionOptions *UpdateOpportunisticEncryptionOptions) (result *OpportunisticEncryptionResp, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(updateOpportunisticEncryptionOptions, "updateOpportunisticEncryptionOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"crn":             *zonesSettings.Crn,
		"zone_identifier": *zonesSettings.ZoneIdentifier,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = zonesSettings.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(zonesSettings.Service.Options.URL, `/v1/{crn}/zones/{zone_identifier}/settings/opportunistic_encryption`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range updateOpportunisticEncryptionOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("zones_settings", "V1", "UpdateOpportunisticEncryption")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if updateOpportunisticEncryptionOptions.Value != nil {
		body["value"] = updateOpportunisticEncryptionOptions.Value
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
	response, err = zonesSettings.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "update_opportunistic_encryption", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalOpportunisticEncryptionResp)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// GetOpportunisticOnion : Get opportunistic onion setting
// Get opportunistic onion setting for a zone.
func (zonesSettings *ZonesSettingsV1) GetOpportunisticOnion(getOpportunisticOnionOptions *GetOpportunisticOnionOptions) (result *OpportunisticOnionResp, response *core.DetailedResponse, err error) {
	result, response, err = zonesSettings.GetOpportunisticOnionWithContext(context.Background(), getOpportunisticOnionOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// GetOpportunisticOnionWithContext is an alternate form of the GetOpportunisticOnion method which supports a Context parameter
func (zonesSettings *ZonesSettingsV1) GetOpportunisticOnionWithContext(ctx context.Context, getOpportunisticOnionOptions *GetOpportunisticOnionOptions) (result *OpportunisticOnionResp, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getOpportunisticOnionOptions, "getOpportunisticOnionOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"crn":             *zonesSettings.Crn,
		"zone_identifier": *zonesSettings.ZoneIdentifier,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = zonesSettings.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(zonesSettings.Service.Options.URL, `/v1/{crn}/zones/{zone_identifier}/settings/opportunistic_onion`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range getOpportunisticOnionOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("zones_settings", "V1", "GetOpportunisticOnion")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = zonesSettings.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "get_opportunistic_onion", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalOpportunisticOnionResp)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// UpdateOpportunisticOnion : Update opportunistic onion setting
// Update opportunistic onion setting for a zone.
func (zonesSettings *ZonesSettingsV1) UpdateOpportunisticOnion(updateOpportunisticOnionOptions *UpdateOpportunisticOnionOptions) (result *OpportunisticOnionResp, response *core.DetailedResponse, err error) {
	result, response, err = zonesSettings.UpdateOpportunisticOnionWithContext(context.Background(), updateOpportunisticOnionOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// UpdateOpportunisticOnionWithContext is an alternate form of the UpdateOpportunisticOnion method which supports a Context parameter
func (zonesSettings *ZonesSettingsV1) UpdateOpportunisticOnionWithContext(ctx context.Context, updateOpportunisticOnionOptions *UpdateOpportunisticOnionOptions) (result *OpportunisticOnionResp, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(updateOpportunisticOnionOptions, "updateOpportunisticOnionOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"crn":             *zonesSettings.Crn,
		"zone_identifier": *zonesSettings.ZoneIdentifier,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = zonesSettings.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(zonesSettings.Service.Options.URL, `/v1/{crn}/zones/{zone_identifier}/settings/opportunistic_onion`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range updateOpportunisticOnionOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("zones_settings", "V1", "UpdateOpportunisticOnion")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if updateOpportunisticOnionOptions.Value != nil {
		body["value"] = updateOpportunisticOnionOptions.Value
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
	response, err = zonesSettings.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "update_opportunistic_onion", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalOpportunisticOnionResp)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// GetChallengeTTL : Get challenge TTL setting
// Get challenge TTL setting for a zone.
func (zonesSettings *ZonesSettingsV1) GetChallengeTTL(getChallengeTtlOptions *GetChallengeTtlOptions) (result *ChallengeTtlResp, response *core.DetailedResponse, err error) {
	result, response, err = zonesSettings.GetChallengeTTLWithContext(context.Background(), getChallengeTtlOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// GetChallengeTTLWithContext is an alternate form of the GetChallengeTTL method which supports a Context parameter
func (zonesSettings *ZonesSettingsV1) GetChallengeTTLWithContext(ctx context.Context, getChallengeTtlOptions *GetChallengeTtlOptions) (result *ChallengeTtlResp, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getChallengeTtlOptions, "getChallengeTtlOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"crn":             *zonesSettings.Crn,
		"zone_identifier": *zonesSettings.ZoneIdentifier,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = zonesSettings.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(zonesSettings.Service.Options.URL, `/v1/{crn}/zones/{zone_identifier}/settings/challenge_ttl`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range getChallengeTtlOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("zones_settings", "V1", "GetChallengeTTL")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = zonesSettings.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "get_challenge_ttl", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalChallengeTtlResp)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// UpdateChallengeTTL : Update challenge TTL setting
// Update challenge TTL setting for a zone.
func (zonesSettings *ZonesSettingsV1) UpdateChallengeTTL(updateChallengeTtlOptions *UpdateChallengeTtlOptions) (result *ChallengeTtlResp, response *core.DetailedResponse, err error) {
	result, response, err = zonesSettings.UpdateChallengeTTLWithContext(context.Background(), updateChallengeTtlOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// UpdateChallengeTTLWithContext is an alternate form of the UpdateChallengeTTL method which supports a Context parameter
func (zonesSettings *ZonesSettingsV1) UpdateChallengeTTLWithContext(ctx context.Context, updateChallengeTtlOptions *UpdateChallengeTtlOptions) (result *ChallengeTtlResp, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(updateChallengeTtlOptions, "updateChallengeTtlOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"crn":             *zonesSettings.Crn,
		"zone_identifier": *zonesSettings.ZoneIdentifier,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = zonesSettings.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(zonesSettings.Service.Options.URL, `/v1/{crn}/zones/{zone_identifier}/settings/challenge_ttl`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range updateChallengeTtlOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("zones_settings", "V1", "UpdateChallengeTTL")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if updateChallengeTtlOptions.Value != nil {
		body["value"] = updateChallengeTtlOptions.Value
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
	response, err = zonesSettings.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "update_challenge_ttl", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalChallengeTtlResp)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// GetAutomaticHttpsRewrites : Get automatic https rewrites setting
// Get automatic https rewrites setting for a zone.
func (zonesSettings *ZonesSettingsV1) GetAutomaticHttpsRewrites(getAutomaticHttpsRewritesOptions *GetAutomaticHttpsRewritesOptions) (result *AutomaticHttpsRewritesResp, response *core.DetailedResponse, err error) {
	result, response, err = zonesSettings.GetAutomaticHttpsRewritesWithContext(context.Background(), getAutomaticHttpsRewritesOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// GetAutomaticHttpsRewritesWithContext is an alternate form of the GetAutomaticHttpsRewrites method which supports a Context parameter
func (zonesSettings *ZonesSettingsV1) GetAutomaticHttpsRewritesWithContext(ctx context.Context, getAutomaticHttpsRewritesOptions *GetAutomaticHttpsRewritesOptions) (result *AutomaticHttpsRewritesResp, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getAutomaticHttpsRewritesOptions, "getAutomaticHttpsRewritesOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"crn":             *zonesSettings.Crn,
		"zone_identifier": *zonesSettings.ZoneIdentifier,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = zonesSettings.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(zonesSettings.Service.Options.URL, `/v1/{crn}/zones/{zone_identifier}/settings/automatic_https_rewrites`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range getAutomaticHttpsRewritesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("zones_settings", "V1", "GetAutomaticHttpsRewrites")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = zonesSettings.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "get_automatic_https_rewrites", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalAutomaticHttpsRewritesResp)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// UpdateAutomaticHttpsRewrites : Update automatic https rewrites setting
// Update automatic https rewrites setting for a zone.
func (zonesSettings *ZonesSettingsV1) UpdateAutomaticHttpsRewrites(updateAutomaticHttpsRewritesOptions *UpdateAutomaticHttpsRewritesOptions) (result *AutomaticHttpsRewritesResp, response *core.DetailedResponse, err error) {
	result, response, err = zonesSettings.UpdateAutomaticHttpsRewritesWithContext(context.Background(), updateAutomaticHttpsRewritesOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// UpdateAutomaticHttpsRewritesWithContext is an alternate form of the UpdateAutomaticHttpsRewrites method which supports a Context parameter
func (zonesSettings *ZonesSettingsV1) UpdateAutomaticHttpsRewritesWithContext(ctx context.Context, updateAutomaticHttpsRewritesOptions *UpdateAutomaticHttpsRewritesOptions) (result *AutomaticHttpsRewritesResp, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(updateAutomaticHttpsRewritesOptions, "updateAutomaticHttpsRewritesOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"crn":             *zonesSettings.Crn,
		"zone_identifier": *zonesSettings.ZoneIdentifier,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = zonesSettings.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(zonesSettings.Service.Options.URL, `/v1/{crn}/zones/{zone_identifier}/settings/automatic_https_rewrites`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range updateAutomaticHttpsRewritesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("zones_settings", "V1", "UpdateAutomaticHttpsRewrites")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if updateAutomaticHttpsRewritesOptions.Value != nil {
		body["value"] = updateAutomaticHttpsRewritesOptions.Value
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
	response, err = zonesSettings.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "update_automatic_https_rewrites", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalAutomaticHttpsRewritesResp)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// GetTrueClientIp : Get true client IP setting
// Get true client IP setting for a zone.
func (zonesSettings *ZonesSettingsV1) GetTrueClientIp(getTrueClientIpOptions *GetTrueClientIpOptions) (result *TrueClientIpResp, response *core.DetailedResponse, err error) {
	result, response, err = zonesSettings.GetTrueClientIpWithContext(context.Background(), getTrueClientIpOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// GetTrueClientIpWithContext is an alternate form of the GetTrueClientIp method which supports a Context parameter
func (zonesSettings *ZonesSettingsV1) GetTrueClientIpWithContext(ctx context.Context, getTrueClientIpOptions *GetTrueClientIpOptions) (result *TrueClientIpResp, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getTrueClientIpOptions, "getTrueClientIpOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"crn":             *zonesSettings.Crn,
		"zone_identifier": *zonesSettings.ZoneIdentifier,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = zonesSettings.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(zonesSettings.Service.Options.URL, `/v1/{crn}/zones/{zone_identifier}/settings/true_client_ip_header`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range getTrueClientIpOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("zones_settings", "V1", "GetTrueClientIp")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = zonesSettings.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "get_true_client_ip", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalTrueClientIpResp)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// UpdateTrueClientIp : Update true client IP setting
// Update true client IP setting for a zone.
func (zonesSettings *ZonesSettingsV1) UpdateTrueClientIp(updateTrueClientIpOptions *UpdateTrueClientIpOptions) (result *TrueClientIpResp, response *core.DetailedResponse, err error) {
	result, response, err = zonesSettings.UpdateTrueClientIpWithContext(context.Background(), updateTrueClientIpOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// UpdateTrueClientIpWithContext is an alternate form of the UpdateTrueClientIp method which supports a Context parameter
func (zonesSettings *ZonesSettingsV1) UpdateTrueClientIpWithContext(ctx context.Context, updateTrueClientIpOptions *UpdateTrueClientIpOptions) (result *TrueClientIpResp, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(updateTrueClientIpOptions, "updateTrueClientIpOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"crn":             *zonesSettings.Crn,
		"zone_identifier": *zonesSettings.ZoneIdentifier,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = zonesSettings.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(zonesSettings.Service.Options.URL, `/v1/{crn}/zones/{zone_identifier}/settings/true_client_ip_header`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range updateTrueClientIpOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("zones_settings", "V1", "UpdateTrueClientIp")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if updateTrueClientIpOptions.Value != nil {
		body["value"] = updateTrueClientIpOptions.Value
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
	response, err = zonesSettings.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "update_true_client_ip", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalTrueClientIpResp)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// GetAlwaysUseHttps : Get always use https setting
// Get always use https setting for a zone.
func (zonesSettings *ZonesSettingsV1) GetAlwaysUseHttps(getAlwaysUseHttpsOptions *GetAlwaysUseHttpsOptions) (result *AlwaysUseHttpsResp, response *core.DetailedResponse, err error) {
	result, response, err = zonesSettings.GetAlwaysUseHttpsWithContext(context.Background(), getAlwaysUseHttpsOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// GetAlwaysUseHttpsWithContext is an alternate form of the GetAlwaysUseHttps method which supports a Context parameter
func (zonesSettings *ZonesSettingsV1) GetAlwaysUseHttpsWithContext(ctx context.Context, getAlwaysUseHttpsOptions *GetAlwaysUseHttpsOptions) (result *AlwaysUseHttpsResp, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getAlwaysUseHttpsOptions, "getAlwaysUseHttpsOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"crn":             *zonesSettings.Crn,
		"zone_identifier": *zonesSettings.ZoneIdentifier,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = zonesSettings.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(zonesSettings.Service.Options.URL, `/v1/{crn}/zones/{zone_identifier}/settings/always_use_https`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range getAlwaysUseHttpsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("zones_settings", "V1", "GetAlwaysUseHttps")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = zonesSettings.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "get_always_use_https", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalAlwaysUseHttpsResp)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// UpdateAlwaysUseHttps : Update always use https setting
// Update always use https setting for a zone.
func (zonesSettings *ZonesSettingsV1) UpdateAlwaysUseHttps(updateAlwaysUseHttpsOptions *UpdateAlwaysUseHttpsOptions) (result *AlwaysUseHttpsResp, response *core.DetailedResponse, err error) {
	result, response, err = zonesSettings.UpdateAlwaysUseHttpsWithContext(context.Background(), updateAlwaysUseHttpsOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// UpdateAlwaysUseHttpsWithContext is an alternate form of the UpdateAlwaysUseHttps method which supports a Context parameter
func (zonesSettings *ZonesSettingsV1) UpdateAlwaysUseHttpsWithContext(ctx context.Context, updateAlwaysUseHttpsOptions *UpdateAlwaysUseHttpsOptions) (result *AlwaysUseHttpsResp, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(updateAlwaysUseHttpsOptions, "updateAlwaysUseHttpsOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"crn":             *zonesSettings.Crn,
		"zone_identifier": *zonesSettings.ZoneIdentifier,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = zonesSettings.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(zonesSettings.Service.Options.URL, `/v1/{crn}/zones/{zone_identifier}/settings/always_use_https`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range updateAlwaysUseHttpsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("zones_settings", "V1", "UpdateAlwaysUseHttps")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if updateAlwaysUseHttpsOptions.Value != nil {
		body["value"] = updateAlwaysUseHttpsOptions.Value
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
	response, err = zonesSettings.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "update_always_use_https", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalAlwaysUseHttpsResp)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// GetImageSizeOptimization : Get image size optimization setting
// Get image size optimization setting for a zone.
func (zonesSettings *ZonesSettingsV1) GetImageSizeOptimization(getImageSizeOptimizationOptions *GetImageSizeOptimizationOptions) (result *ImageSizeOptimizationResp, response *core.DetailedResponse, err error) {
	result, response, err = zonesSettings.GetImageSizeOptimizationWithContext(context.Background(), getImageSizeOptimizationOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// GetImageSizeOptimizationWithContext is an alternate form of the GetImageSizeOptimization method which supports a Context parameter
func (zonesSettings *ZonesSettingsV1) GetImageSizeOptimizationWithContext(ctx context.Context, getImageSizeOptimizationOptions *GetImageSizeOptimizationOptions) (result *ImageSizeOptimizationResp, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getImageSizeOptimizationOptions, "getImageSizeOptimizationOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"crn":             *zonesSettings.Crn,
		"zone_identifier": *zonesSettings.ZoneIdentifier,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = zonesSettings.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(zonesSettings.Service.Options.URL, `/v1/{crn}/zones/{zone_identifier}/settings/image_size_optimization`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range getImageSizeOptimizationOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("zones_settings", "V1", "GetImageSizeOptimization")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = zonesSettings.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "get_image_size_optimization", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalImageSizeOptimizationResp)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// UpdateImageSizeOptimization : Update image size optimization setting
// Update image size optimization setting for a zone.
func (zonesSettings *ZonesSettingsV1) UpdateImageSizeOptimization(updateImageSizeOptimizationOptions *UpdateImageSizeOptimizationOptions) (result *ImageSizeOptimizationResp, response *core.DetailedResponse, err error) {
	result, response, err = zonesSettings.UpdateImageSizeOptimizationWithContext(context.Background(), updateImageSizeOptimizationOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// UpdateImageSizeOptimizationWithContext is an alternate form of the UpdateImageSizeOptimization method which supports a Context parameter
func (zonesSettings *ZonesSettingsV1) UpdateImageSizeOptimizationWithContext(ctx context.Context, updateImageSizeOptimizationOptions *UpdateImageSizeOptimizationOptions) (result *ImageSizeOptimizationResp, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(updateImageSizeOptimizationOptions, "updateImageSizeOptimizationOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"crn":             *zonesSettings.Crn,
		"zone_identifier": *zonesSettings.ZoneIdentifier,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = zonesSettings.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(zonesSettings.Service.Options.URL, `/v1/{crn}/zones/{zone_identifier}/settings/image_size_optimization`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range updateImageSizeOptimizationOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("zones_settings", "V1", "UpdateImageSizeOptimization")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if updateImageSizeOptimizationOptions.Value != nil {
		body["value"] = updateImageSizeOptimizationOptions.Value
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
	response, err = zonesSettings.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "update_image_size_optimization", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalImageSizeOptimizationResp)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// GetScriptLoadOptimization : Get script load optimization setting
// Get script load optimization setting for a zone.
func (zonesSettings *ZonesSettingsV1) GetScriptLoadOptimization(getScriptLoadOptimizationOptions *GetScriptLoadOptimizationOptions) (result *ScriptLoadOptimizationResp, response *core.DetailedResponse, err error) {
	result, response, err = zonesSettings.GetScriptLoadOptimizationWithContext(context.Background(), getScriptLoadOptimizationOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// GetScriptLoadOptimizationWithContext is an alternate form of the GetScriptLoadOptimization method which supports a Context parameter
func (zonesSettings *ZonesSettingsV1) GetScriptLoadOptimizationWithContext(ctx context.Context, getScriptLoadOptimizationOptions *GetScriptLoadOptimizationOptions) (result *ScriptLoadOptimizationResp, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getScriptLoadOptimizationOptions, "getScriptLoadOptimizationOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"crn":             *zonesSettings.Crn,
		"zone_identifier": *zonesSettings.ZoneIdentifier,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = zonesSettings.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(zonesSettings.Service.Options.URL, `/v1/{crn}/zones/{zone_identifier}/settings/script_load_optimization`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range getScriptLoadOptimizationOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("zones_settings", "V1", "GetScriptLoadOptimization")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = zonesSettings.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "get_script_load_optimization", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalScriptLoadOptimizationResp)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// UpdateScriptLoadOptimization : Update script load optimization setting
// Update script load optimization setting for a zone.
func (zonesSettings *ZonesSettingsV1) UpdateScriptLoadOptimization(updateScriptLoadOptimizationOptions *UpdateScriptLoadOptimizationOptions) (result *ScriptLoadOptimizationResp, response *core.DetailedResponse, err error) {
	result, response, err = zonesSettings.UpdateScriptLoadOptimizationWithContext(context.Background(), updateScriptLoadOptimizationOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// UpdateScriptLoadOptimizationWithContext is an alternate form of the UpdateScriptLoadOptimization method which supports a Context parameter
func (zonesSettings *ZonesSettingsV1) UpdateScriptLoadOptimizationWithContext(ctx context.Context, updateScriptLoadOptimizationOptions *UpdateScriptLoadOptimizationOptions) (result *ScriptLoadOptimizationResp, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(updateScriptLoadOptimizationOptions, "updateScriptLoadOptimizationOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"crn":             *zonesSettings.Crn,
		"zone_identifier": *zonesSettings.ZoneIdentifier,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = zonesSettings.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(zonesSettings.Service.Options.URL, `/v1/{crn}/zones/{zone_identifier}/settings/script_load_optimization`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range updateScriptLoadOptimizationOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("zones_settings", "V1", "UpdateScriptLoadOptimization")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if updateScriptLoadOptimizationOptions.Value != nil {
		body["value"] = updateScriptLoadOptimizationOptions.Value
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
	response, err = zonesSettings.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "update_script_load_optimization", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalScriptLoadOptimizationResp)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// GetImageLoadOptimization : Get image load optimizationn setting
// Get image load optimizationn setting for a zone.
func (zonesSettings *ZonesSettingsV1) GetImageLoadOptimization(getImageLoadOptimizationOptions *GetImageLoadOptimizationOptions) (result *ImageLoadOptimizationResp, response *core.DetailedResponse, err error) {
	result, response, err = zonesSettings.GetImageLoadOptimizationWithContext(context.Background(), getImageLoadOptimizationOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// GetImageLoadOptimizationWithContext is an alternate form of the GetImageLoadOptimization method which supports a Context parameter
func (zonesSettings *ZonesSettingsV1) GetImageLoadOptimizationWithContext(ctx context.Context, getImageLoadOptimizationOptions *GetImageLoadOptimizationOptions) (result *ImageLoadOptimizationResp, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getImageLoadOptimizationOptions, "getImageLoadOptimizationOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"crn":             *zonesSettings.Crn,
		"zone_identifier": *zonesSettings.ZoneIdentifier,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = zonesSettings.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(zonesSettings.Service.Options.URL, `/v1/{crn}/zones/{zone_identifier}/settings/image_load_optimization`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range getImageLoadOptimizationOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("zones_settings", "V1", "GetImageLoadOptimization")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = zonesSettings.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "get_image_load_optimization", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalImageLoadOptimizationResp)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// UpdateImageLoadOptimization : Update image load optimizationn setting
// Update image load optimizationn setting for a zone.
func (zonesSettings *ZonesSettingsV1) UpdateImageLoadOptimization(updateImageLoadOptimizationOptions *UpdateImageLoadOptimizationOptions) (result *ImageLoadOptimizationResp, response *core.DetailedResponse, err error) {
	result, response, err = zonesSettings.UpdateImageLoadOptimizationWithContext(context.Background(), updateImageLoadOptimizationOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// UpdateImageLoadOptimizationWithContext is an alternate form of the UpdateImageLoadOptimization method which supports a Context parameter
func (zonesSettings *ZonesSettingsV1) UpdateImageLoadOptimizationWithContext(ctx context.Context, updateImageLoadOptimizationOptions *UpdateImageLoadOptimizationOptions) (result *ImageLoadOptimizationResp, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(updateImageLoadOptimizationOptions, "updateImageLoadOptimizationOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"crn":             *zonesSettings.Crn,
		"zone_identifier": *zonesSettings.ZoneIdentifier,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = zonesSettings.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(zonesSettings.Service.Options.URL, `/v1/{crn}/zones/{zone_identifier}/settings/image_load_optimization`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range updateImageLoadOptimizationOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("zones_settings", "V1", "UpdateImageLoadOptimization")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if updateImageLoadOptimizationOptions.Value != nil {
		body["value"] = updateImageLoadOptimizationOptions.Value
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
	response, err = zonesSettings.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "update_image_load_optimization", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalImageLoadOptimizationResp)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// GetMinify : Get minify setting
// Get minify setting for a zone.
func (zonesSettings *ZonesSettingsV1) GetMinify(getMinifyOptions *GetMinifyOptions) (result *MinifyResp, response *core.DetailedResponse, err error) {
	result, response, err = zonesSettings.GetMinifyWithContext(context.Background(), getMinifyOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// GetMinifyWithContext is an alternate form of the GetMinify method which supports a Context parameter
func (zonesSettings *ZonesSettingsV1) GetMinifyWithContext(ctx context.Context, getMinifyOptions *GetMinifyOptions) (result *MinifyResp, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getMinifyOptions, "getMinifyOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"crn":             *zonesSettings.Crn,
		"zone_identifier": *zonesSettings.ZoneIdentifier,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = zonesSettings.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(zonesSettings.Service.Options.URL, `/v1/{crn}/zones/{zone_identifier}/settings/minify`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range getMinifyOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("zones_settings", "V1", "GetMinify")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = zonesSettings.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "get_minify", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalMinifyResp)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// UpdateMinify : Update minify setting
// Update minify setting for a zone.
func (zonesSettings *ZonesSettingsV1) UpdateMinify(updateMinifyOptions *UpdateMinifyOptions) (result *MinifyResp, response *core.DetailedResponse, err error) {
	result, response, err = zonesSettings.UpdateMinifyWithContext(context.Background(), updateMinifyOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// UpdateMinifyWithContext is an alternate form of the UpdateMinify method which supports a Context parameter
func (zonesSettings *ZonesSettingsV1) UpdateMinifyWithContext(ctx context.Context, updateMinifyOptions *UpdateMinifyOptions) (result *MinifyResp, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(updateMinifyOptions, "updateMinifyOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"crn":             *zonesSettings.Crn,
		"zone_identifier": *zonesSettings.ZoneIdentifier,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = zonesSettings.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(zonesSettings.Service.Options.URL, `/v1/{crn}/zones/{zone_identifier}/settings/minify`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range updateMinifyOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("zones_settings", "V1", "UpdateMinify")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if updateMinifyOptions.Value != nil {
		body["value"] = updateMinifyOptions.Value
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
	response, err = zonesSettings.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "update_minify", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalMinifyResp)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// GetMinTlsVersion : Get minimum TLS version setting
// Get minimum TLS version setting for a zone.
func (zonesSettings *ZonesSettingsV1) GetMinTlsVersion(getMinTlsVersionOptions *GetMinTlsVersionOptions) (result *MinTlsVersionResp, response *core.DetailedResponse, err error) {
	result, response, err = zonesSettings.GetMinTlsVersionWithContext(context.Background(), getMinTlsVersionOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// GetMinTlsVersionWithContext is an alternate form of the GetMinTlsVersion method which supports a Context parameter
func (zonesSettings *ZonesSettingsV1) GetMinTlsVersionWithContext(ctx context.Context, getMinTlsVersionOptions *GetMinTlsVersionOptions) (result *MinTlsVersionResp, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getMinTlsVersionOptions, "getMinTlsVersionOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"crn":             *zonesSettings.Crn,
		"zone_identifier": *zonesSettings.ZoneIdentifier,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = zonesSettings.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(zonesSettings.Service.Options.URL, `/v1/{crn}/zones/{zone_identifier}/settings/min_tls_version`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range getMinTlsVersionOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("zones_settings", "V1", "GetMinTlsVersion")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = zonesSettings.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "get_min_tls_version", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalMinTlsVersionResp)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// UpdateMinTlsVersion : Update minimum TLS version setting
// Update minimum TLS version setting for a zone.
func (zonesSettings *ZonesSettingsV1) UpdateMinTlsVersion(updateMinTlsVersionOptions *UpdateMinTlsVersionOptions) (result *MinTlsVersionResp, response *core.DetailedResponse, err error) {
	result, response, err = zonesSettings.UpdateMinTlsVersionWithContext(context.Background(), updateMinTlsVersionOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// UpdateMinTlsVersionWithContext is an alternate form of the UpdateMinTlsVersion method which supports a Context parameter
func (zonesSettings *ZonesSettingsV1) UpdateMinTlsVersionWithContext(ctx context.Context, updateMinTlsVersionOptions *UpdateMinTlsVersionOptions) (result *MinTlsVersionResp, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(updateMinTlsVersionOptions, "updateMinTlsVersionOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"crn":             *zonesSettings.Crn,
		"zone_identifier": *zonesSettings.ZoneIdentifier,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = zonesSettings.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(zonesSettings.Service.Options.URL, `/v1/{crn}/zones/{zone_identifier}/settings/min_tls_version`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range updateMinTlsVersionOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("zones_settings", "V1", "UpdateMinTlsVersion")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if updateMinTlsVersionOptions.Value != nil {
		body["value"] = updateMinTlsVersionOptions.Value
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
	response, err = zonesSettings.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "update_min_tls_version", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalMinTlsVersionResp)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// GetIpGeolocation : Get IP geolocation setting
// Get IP geolocation setting for a zone.
func (zonesSettings *ZonesSettingsV1) GetIpGeolocation(getIpGeolocationOptions *GetIpGeolocationOptions) (result *IpGeolocationResp, response *core.DetailedResponse, err error) {
	result, response, err = zonesSettings.GetIpGeolocationWithContext(context.Background(), getIpGeolocationOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// GetIpGeolocationWithContext is an alternate form of the GetIpGeolocation method which supports a Context parameter
func (zonesSettings *ZonesSettingsV1) GetIpGeolocationWithContext(ctx context.Context, getIpGeolocationOptions *GetIpGeolocationOptions) (result *IpGeolocationResp, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getIpGeolocationOptions, "getIpGeolocationOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"crn":             *zonesSettings.Crn,
		"zone_identifier": *zonesSettings.ZoneIdentifier,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = zonesSettings.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(zonesSettings.Service.Options.URL, `/v1/{crn}/zones/{zone_identifier}/settings/ip_geolocation`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range getIpGeolocationOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("zones_settings", "V1", "GetIpGeolocation")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = zonesSettings.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "get_ip_geolocation", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalIpGeolocationResp)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// UpdateIpGeolocation : Update IP geolocation setting
// Update IP geolocation setting for a zone.
func (zonesSettings *ZonesSettingsV1) UpdateIpGeolocation(updateIpGeolocationOptions *UpdateIpGeolocationOptions) (result *IpGeolocationResp, response *core.DetailedResponse, err error) {
	result, response, err = zonesSettings.UpdateIpGeolocationWithContext(context.Background(), updateIpGeolocationOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// UpdateIpGeolocationWithContext is an alternate form of the UpdateIpGeolocation method which supports a Context parameter
func (zonesSettings *ZonesSettingsV1) UpdateIpGeolocationWithContext(ctx context.Context, updateIpGeolocationOptions *UpdateIpGeolocationOptions) (result *IpGeolocationResp, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(updateIpGeolocationOptions, "updateIpGeolocationOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"crn":             *zonesSettings.Crn,
		"zone_identifier": *zonesSettings.ZoneIdentifier,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = zonesSettings.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(zonesSettings.Service.Options.URL, `/v1/{crn}/zones/{zone_identifier}/settings/ip_geolocation`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range updateIpGeolocationOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("zones_settings", "V1", "UpdateIpGeolocation")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if updateIpGeolocationOptions.Value != nil {
		body["value"] = updateIpGeolocationOptions.Value
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
	response, err = zonesSettings.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "update_ip_geolocation", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalIpGeolocationResp)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// GetServerSideExclude : Get server side exclude setting
// Get server side exclude setting for a zone.
func (zonesSettings *ZonesSettingsV1) GetServerSideExclude(getServerSideExcludeOptions *GetServerSideExcludeOptions) (result *ServerSideExcludeResp, response *core.DetailedResponse, err error) {
	result, response, err = zonesSettings.GetServerSideExcludeWithContext(context.Background(), getServerSideExcludeOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// GetServerSideExcludeWithContext is an alternate form of the GetServerSideExclude method which supports a Context parameter
func (zonesSettings *ZonesSettingsV1) GetServerSideExcludeWithContext(ctx context.Context, getServerSideExcludeOptions *GetServerSideExcludeOptions) (result *ServerSideExcludeResp, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getServerSideExcludeOptions, "getServerSideExcludeOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"crn":             *zonesSettings.Crn,
		"zone_identifier": *zonesSettings.ZoneIdentifier,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = zonesSettings.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(zonesSettings.Service.Options.URL, `/v1/{crn}/zones/{zone_identifier}/settings/server_side_exclude`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range getServerSideExcludeOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("zones_settings", "V1", "GetServerSideExclude")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = zonesSettings.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "get_server_side_exclude", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalServerSideExcludeResp)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// UpdateServerSideExclude : Update server side exclude setting
// Update server side exclude setting for a zone.
func (zonesSettings *ZonesSettingsV1) UpdateServerSideExclude(updateServerSideExcludeOptions *UpdateServerSideExcludeOptions) (result *ServerSideExcludeResp, response *core.DetailedResponse, err error) {
	result, response, err = zonesSettings.UpdateServerSideExcludeWithContext(context.Background(), updateServerSideExcludeOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// UpdateServerSideExcludeWithContext is an alternate form of the UpdateServerSideExclude method which supports a Context parameter
func (zonesSettings *ZonesSettingsV1) UpdateServerSideExcludeWithContext(ctx context.Context, updateServerSideExcludeOptions *UpdateServerSideExcludeOptions) (result *ServerSideExcludeResp, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(updateServerSideExcludeOptions, "updateServerSideExcludeOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"crn":             *zonesSettings.Crn,
		"zone_identifier": *zonesSettings.ZoneIdentifier,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = zonesSettings.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(zonesSettings.Service.Options.URL, `/v1/{crn}/zones/{zone_identifier}/settings/server_side_exclude`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range updateServerSideExcludeOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("zones_settings", "V1", "UpdateServerSideExclude")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if updateServerSideExcludeOptions.Value != nil {
		body["value"] = updateServerSideExcludeOptions.Value
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
	response, err = zonesSettings.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "update_server_side_exclude", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalServerSideExcludeResp)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// GetSecurityHeader : Get HTTP strict transport security setting
// Get HTTP strict transport security setting for a zone.
func (zonesSettings *ZonesSettingsV1) GetSecurityHeader(getSecurityHeaderOptions *GetSecurityHeaderOptions) (result *SecurityHeaderResp, response *core.DetailedResponse, err error) {
	result, response, err = zonesSettings.GetSecurityHeaderWithContext(context.Background(), getSecurityHeaderOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// GetSecurityHeaderWithContext is an alternate form of the GetSecurityHeader method which supports a Context parameter
func (zonesSettings *ZonesSettingsV1) GetSecurityHeaderWithContext(ctx context.Context, getSecurityHeaderOptions *GetSecurityHeaderOptions) (result *SecurityHeaderResp, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getSecurityHeaderOptions, "getSecurityHeaderOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"crn":             *zonesSettings.Crn,
		"zone_identifier": *zonesSettings.ZoneIdentifier,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = zonesSettings.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(zonesSettings.Service.Options.URL, `/v1/{crn}/zones/{zone_identifier}/settings/security_header`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range getSecurityHeaderOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("zones_settings", "V1", "GetSecurityHeader")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = zonesSettings.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "get_security_header", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalSecurityHeaderResp)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// UpdateSecurityHeader : Update HTTP strict transport security setting
// Update HTTP strict transport security setting for a zone.
func (zonesSettings *ZonesSettingsV1) UpdateSecurityHeader(updateSecurityHeaderOptions *UpdateSecurityHeaderOptions) (result *SecurityHeaderResp, response *core.DetailedResponse, err error) {
	result, response, err = zonesSettings.UpdateSecurityHeaderWithContext(context.Background(), updateSecurityHeaderOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// UpdateSecurityHeaderWithContext is an alternate form of the UpdateSecurityHeader method which supports a Context parameter
func (zonesSettings *ZonesSettingsV1) UpdateSecurityHeaderWithContext(ctx context.Context, updateSecurityHeaderOptions *UpdateSecurityHeaderOptions) (result *SecurityHeaderResp, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(updateSecurityHeaderOptions, "updateSecurityHeaderOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"crn":             *zonesSettings.Crn,
		"zone_identifier": *zonesSettings.ZoneIdentifier,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = zonesSettings.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(zonesSettings.Service.Options.URL, `/v1/{crn}/zones/{zone_identifier}/settings/security_header`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range updateSecurityHeaderOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("zones_settings", "V1", "UpdateSecurityHeader")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if updateSecurityHeaderOptions.Value != nil {
		body["value"] = updateSecurityHeaderOptions.Value
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
	response, err = zonesSettings.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "update_security_header", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalSecurityHeaderResp)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// GetMobileRedirect : Get mobile redirect setting
// Get mobile redirect setting for a zone.
func (zonesSettings *ZonesSettingsV1) GetMobileRedirect(getMobileRedirectOptions *GetMobileRedirectOptions) (result *MobileRedirectResp, response *core.DetailedResponse, err error) {
	result, response, err = zonesSettings.GetMobileRedirectWithContext(context.Background(), getMobileRedirectOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// GetMobileRedirectWithContext is an alternate form of the GetMobileRedirect method which supports a Context parameter
func (zonesSettings *ZonesSettingsV1) GetMobileRedirectWithContext(ctx context.Context, getMobileRedirectOptions *GetMobileRedirectOptions) (result *MobileRedirectResp, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getMobileRedirectOptions, "getMobileRedirectOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"crn":             *zonesSettings.Crn,
		"zone_identifier": *zonesSettings.ZoneIdentifier,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = zonesSettings.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(zonesSettings.Service.Options.URL, `/v1/{crn}/zones/{zone_identifier}/settings/mobile_redirect`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range getMobileRedirectOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("zones_settings", "V1", "GetMobileRedirect")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = zonesSettings.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "get_mobile_redirect", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalMobileRedirectResp)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// UpdateMobileRedirect : Update mobile redirect setting
// Update mobile redirect setting for a zone.
func (zonesSettings *ZonesSettingsV1) UpdateMobileRedirect(updateMobileRedirectOptions *UpdateMobileRedirectOptions) (result *MobileRedirectResp, response *core.DetailedResponse, err error) {
	result, response, err = zonesSettings.UpdateMobileRedirectWithContext(context.Background(), updateMobileRedirectOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// UpdateMobileRedirectWithContext is an alternate form of the UpdateMobileRedirect method which supports a Context parameter
func (zonesSettings *ZonesSettingsV1) UpdateMobileRedirectWithContext(ctx context.Context, updateMobileRedirectOptions *UpdateMobileRedirectOptions) (result *MobileRedirectResp, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(updateMobileRedirectOptions, "updateMobileRedirectOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"crn":             *zonesSettings.Crn,
		"zone_identifier": *zonesSettings.ZoneIdentifier,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = zonesSettings.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(zonesSettings.Service.Options.URL, `/v1/{crn}/zones/{zone_identifier}/settings/mobile_redirect`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range updateMobileRedirectOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("zones_settings", "V1", "UpdateMobileRedirect")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if updateMobileRedirectOptions.Value != nil {
		body["value"] = updateMobileRedirectOptions.Value
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
	response, err = zonesSettings.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "update_mobile_redirect", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalMobileRedirectResp)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// GetPrefetchPreload : Get prefetch URLs from header setting
// Get prefetch URLs from header setting for a zone.
func (zonesSettings *ZonesSettingsV1) GetPrefetchPreload(getPrefetchPreloadOptions *GetPrefetchPreloadOptions) (result *PrefetchPreloadResp, response *core.DetailedResponse, err error) {
	result, response, err = zonesSettings.GetPrefetchPreloadWithContext(context.Background(), getPrefetchPreloadOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// GetPrefetchPreloadWithContext is an alternate form of the GetPrefetchPreload method which supports a Context parameter
func (zonesSettings *ZonesSettingsV1) GetPrefetchPreloadWithContext(ctx context.Context, getPrefetchPreloadOptions *GetPrefetchPreloadOptions) (result *PrefetchPreloadResp, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getPrefetchPreloadOptions, "getPrefetchPreloadOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"crn":             *zonesSettings.Crn,
		"zone_identifier": *zonesSettings.ZoneIdentifier,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = zonesSettings.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(zonesSettings.Service.Options.URL, `/v1/{crn}/zones/{zone_identifier}/settings/prefetch_preload`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range getPrefetchPreloadOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("zones_settings", "V1", "GetPrefetchPreload")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = zonesSettings.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "get_prefetch_preload", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalPrefetchPreloadResp)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// UpdatePrefetchPreload : Update prefetch URLs from header setting
// Update prefetch URLs from header setting for a zone.
func (zonesSettings *ZonesSettingsV1) UpdatePrefetchPreload(updatePrefetchPreloadOptions *UpdatePrefetchPreloadOptions) (result *PrefetchPreloadResp, response *core.DetailedResponse, err error) {
	result, response, err = zonesSettings.UpdatePrefetchPreloadWithContext(context.Background(), updatePrefetchPreloadOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// UpdatePrefetchPreloadWithContext is an alternate form of the UpdatePrefetchPreload method which supports a Context parameter
func (zonesSettings *ZonesSettingsV1) UpdatePrefetchPreloadWithContext(ctx context.Context, updatePrefetchPreloadOptions *UpdatePrefetchPreloadOptions) (result *PrefetchPreloadResp, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(updatePrefetchPreloadOptions, "updatePrefetchPreloadOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"crn":             *zonesSettings.Crn,
		"zone_identifier": *zonesSettings.ZoneIdentifier,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = zonesSettings.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(zonesSettings.Service.Options.URL, `/v1/{crn}/zones/{zone_identifier}/settings/prefetch_preload`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range updatePrefetchPreloadOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("zones_settings", "V1", "UpdatePrefetchPreload")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if updatePrefetchPreloadOptions.Value != nil {
		body["value"] = updatePrefetchPreloadOptions.Value
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
	response, err = zonesSettings.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "update_prefetch_preload", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalPrefetchPreloadResp)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// GetHttp2 : Get http/2 setting
// Get http/2 setting for a zone.
func (zonesSettings *ZonesSettingsV1) GetHttp2(getHttp2Options *GetHttp2Options) (result *Http2Resp, response *core.DetailedResponse, err error) {
	result, response, err = zonesSettings.GetHttp2WithContext(context.Background(), getHttp2Options)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// GetHttp2WithContext is an alternate form of the GetHttp2 method which supports a Context parameter
func (zonesSettings *ZonesSettingsV1) GetHttp2WithContext(ctx context.Context, getHttp2Options *GetHttp2Options) (result *Http2Resp, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getHttp2Options, "getHttp2Options")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"crn":             *zonesSettings.Crn,
		"zone_identifier": *zonesSettings.ZoneIdentifier,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = zonesSettings.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(zonesSettings.Service.Options.URL, `/v1/{crn}/zones/{zone_identifier}/settings/http2`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range getHttp2Options.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("zones_settings", "V1", "GetHttp2")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = zonesSettings.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "get_http2", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalHttp2Resp)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// UpdateHttp2 : Update http/2 setting
// Update http/2 setting for a zone.
func (zonesSettings *ZonesSettingsV1) UpdateHttp2(updateHttp2Options *UpdateHttp2Options) (result *Http2Resp, response *core.DetailedResponse, err error) {
	result, response, err = zonesSettings.UpdateHttp2WithContext(context.Background(), updateHttp2Options)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// UpdateHttp2WithContext is an alternate form of the UpdateHttp2 method which supports a Context parameter
func (zonesSettings *ZonesSettingsV1) UpdateHttp2WithContext(ctx context.Context, updateHttp2Options *UpdateHttp2Options) (result *Http2Resp, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(updateHttp2Options, "updateHttp2Options")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"crn":             *zonesSettings.Crn,
		"zone_identifier": *zonesSettings.ZoneIdentifier,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = zonesSettings.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(zonesSettings.Service.Options.URL, `/v1/{crn}/zones/{zone_identifier}/settings/http2`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range updateHttp2Options.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("zones_settings", "V1", "UpdateHttp2")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if updateHttp2Options.Value != nil {
		body["value"] = updateHttp2Options.Value
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
	response, err = zonesSettings.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "update_http2", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalHttp2Resp)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// GetHttp3 : Get http/3 setting
// Get http/3 setting for a zone.
func (zonesSettings *ZonesSettingsV1) GetHttp3(getHttp3Options *GetHttp3Options) (result *Http3Resp, response *core.DetailedResponse, err error) {
	result, response, err = zonesSettings.GetHttp3WithContext(context.Background(), getHttp3Options)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// GetHttp3WithContext is an alternate form of the GetHttp3 method which supports a Context parameter
func (zonesSettings *ZonesSettingsV1) GetHttp3WithContext(ctx context.Context, getHttp3Options *GetHttp3Options) (result *Http3Resp, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getHttp3Options, "getHttp3Options")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"crn":             *zonesSettings.Crn,
		"zone_identifier": *zonesSettings.ZoneIdentifier,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = zonesSettings.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(zonesSettings.Service.Options.URL, `/v1/{crn}/zones/{zone_identifier}/settings/http3`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range getHttp3Options.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("zones_settings", "V1", "GetHttp3")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = zonesSettings.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "get_http3", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalHttp3Resp)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// UpdateHttp3 : Update http/3 setting
// Update http/3 setting for a zone.
func (zonesSettings *ZonesSettingsV1) UpdateHttp3(updateHttp3Options *UpdateHttp3Options) (result *Http3Resp, response *core.DetailedResponse, err error) {
	result, response, err = zonesSettings.UpdateHttp3WithContext(context.Background(), updateHttp3Options)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// UpdateHttp3WithContext is an alternate form of the UpdateHttp3 method which supports a Context parameter
func (zonesSettings *ZonesSettingsV1) UpdateHttp3WithContext(ctx context.Context, updateHttp3Options *UpdateHttp3Options) (result *Http3Resp, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(updateHttp3Options, "updateHttp3Options")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"crn":             *zonesSettings.Crn,
		"zone_identifier": *zonesSettings.ZoneIdentifier,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = zonesSettings.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(zonesSettings.Service.Options.URL, `/v1/{crn}/zones/{zone_identifier}/settings/http3`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range updateHttp3Options.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("zones_settings", "V1", "UpdateHttp3")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if updateHttp3Options.Value != nil {
		body["value"] = updateHttp3Options.Value
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
	response, err = zonesSettings.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "update_http3", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalHttp3Resp)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// GetIpv6 : Get IPv6 compatibility setting
// Get IPv6 compatibility setting for a zone.
func (zonesSettings *ZonesSettingsV1) GetIpv6(getIpv6Options *GetIpv6Options) (result *Ipv6Resp, response *core.DetailedResponse, err error) {
	result, response, err = zonesSettings.GetIpv6WithContext(context.Background(), getIpv6Options)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// GetIpv6WithContext is an alternate form of the GetIpv6 method which supports a Context parameter
func (zonesSettings *ZonesSettingsV1) GetIpv6WithContext(ctx context.Context, getIpv6Options *GetIpv6Options) (result *Ipv6Resp, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getIpv6Options, "getIpv6Options")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"crn":             *zonesSettings.Crn,
		"zone_identifier": *zonesSettings.ZoneIdentifier,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = zonesSettings.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(zonesSettings.Service.Options.URL, `/v1/{crn}/zones/{zone_identifier}/settings/ipv6`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range getIpv6Options.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("zones_settings", "V1", "GetIpv6")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = zonesSettings.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "get_ipv6", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalIpv6Resp)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// UpdateIpv6 : Update IPv6 compatibility setting
// Update IPv6 compatibility setting for a zone.
func (zonesSettings *ZonesSettingsV1) UpdateIpv6(updateIpv6Options *UpdateIpv6Options) (result *Ipv6Resp, response *core.DetailedResponse, err error) {
	result, response, err = zonesSettings.UpdateIpv6WithContext(context.Background(), updateIpv6Options)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// UpdateIpv6WithContext is an alternate form of the UpdateIpv6 method which supports a Context parameter
func (zonesSettings *ZonesSettingsV1) UpdateIpv6WithContext(ctx context.Context, updateIpv6Options *UpdateIpv6Options) (result *Ipv6Resp, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(updateIpv6Options, "updateIpv6Options")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"crn":             *zonesSettings.Crn,
		"zone_identifier": *zonesSettings.ZoneIdentifier,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = zonesSettings.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(zonesSettings.Service.Options.URL, `/v1/{crn}/zones/{zone_identifier}/settings/ipv6`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range updateIpv6Options.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("zones_settings", "V1", "UpdateIpv6")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if updateIpv6Options.Value != nil {
		body["value"] = updateIpv6Options.Value
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
	response, err = zonesSettings.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "update_ipv6", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalIpv6Resp)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// GetWebSockets : Get web sockets setting
// Get web sockets setting for a zone.
func (zonesSettings *ZonesSettingsV1) GetWebSockets(getWebSocketsOptions *GetWebSocketsOptions) (result *WebsocketsResp, response *core.DetailedResponse, err error) {
	result, response, err = zonesSettings.GetWebSocketsWithContext(context.Background(), getWebSocketsOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// GetWebSocketsWithContext is an alternate form of the GetWebSockets method which supports a Context parameter
func (zonesSettings *ZonesSettingsV1) GetWebSocketsWithContext(ctx context.Context, getWebSocketsOptions *GetWebSocketsOptions) (result *WebsocketsResp, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getWebSocketsOptions, "getWebSocketsOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"crn":             *zonesSettings.Crn,
		"zone_identifier": *zonesSettings.ZoneIdentifier,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = zonesSettings.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(zonesSettings.Service.Options.URL, `/v1/{crn}/zones/{zone_identifier}/settings/websockets`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range getWebSocketsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("zones_settings", "V1", "GetWebSockets")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = zonesSettings.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "get_web_sockets", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalWebsocketsResp)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// UpdateWebSockets : Update web sockets setting
// Update web sockets setting for a zone.
func (zonesSettings *ZonesSettingsV1) UpdateWebSockets(updateWebSocketsOptions *UpdateWebSocketsOptions) (result *WebsocketsResp, response *core.DetailedResponse, err error) {
	result, response, err = zonesSettings.UpdateWebSocketsWithContext(context.Background(), updateWebSocketsOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// UpdateWebSocketsWithContext is an alternate form of the UpdateWebSockets method which supports a Context parameter
func (zonesSettings *ZonesSettingsV1) UpdateWebSocketsWithContext(ctx context.Context, updateWebSocketsOptions *UpdateWebSocketsOptions) (result *WebsocketsResp, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(updateWebSocketsOptions, "updateWebSocketsOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"crn":             *zonesSettings.Crn,
		"zone_identifier": *zonesSettings.ZoneIdentifier,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = zonesSettings.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(zonesSettings.Service.Options.URL, `/v1/{crn}/zones/{zone_identifier}/settings/websockets`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range updateWebSocketsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("zones_settings", "V1", "UpdateWebSockets")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if updateWebSocketsOptions.Value != nil {
		body["value"] = updateWebSocketsOptions.Value
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
	response, err = zonesSettings.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "update_web_sockets", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalWebsocketsResp)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// GetPseudoIpv4 : Get pseudo IPv4 setting
// Get pseudo IPv4 setting for a zone.
func (zonesSettings *ZonesSettingsV1) GetPseudoIpv4(getPseudoIpv4Options *GetPseudoIpv4Options) (result *PseudoIpv4Resp, response *core.DetailedResponse, err error) {
	result, response, err = zonesSettings.GetPseudoIpv4WithContext(context.Background(), getPseudoIpv4Options)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// GetPseudoIpv4WithContext is an alternate form of the GetPseudoIpv4 method which supports a Context parameter
func (zonesSettings *ZonesSettingsV1) GetPseudoIpv4WithContext(ctx context.Context, getPseudoIpv4Options *GetPseudoIpv4Options) (result *PseudoIpv4Resp, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getPseudoIpv4Options, "getPseudoIpv4Options")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"crn":             *zonesSettings.Crn,
		"zone_identifier": *zonesSettings.ZoneIdentifier,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = zonesSettings.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(zonesSettings.Service.Options.URL, `/v1/{crn}/zones/{zone_identifier}/settings/pseudo_ipv4`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range getPseudoIpv4Options.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("zones_settings", "V1", "GetPseudoIpv4")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = zonesSettings.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "get_pseudo_ipv4", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalPseudoIpv4Resp)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// UpdatePseudoIpv4 : Update pseudo IPv4 setting
// Update pseudo IPv4 setting for a zone.
func (zonesSettings *ZonesSettingsV1) UpdatePseudoIpv4(updatePseudoIpv4Options *UpdatePseudoIpv4Options) (result *PseudoIpv4Resp, response *core.DetailedResponse, err error) {
	result, response, err = zonesSettings.UpdatePseudoIpv4WithContext(context.Background(), updatePseudoIpv4Options)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// UpdatePseudoIpv4WithContext is an alternate form of the UpdatePseudoIpv4 method which supports a Context parameter
func (zonesSettings *ZonesSettingsV1) UpdatePseudoIpv4WithContext(ctx context.Context, updatePseudoIpv4Options *UpdatePseudoIpv4Options) (result *PseudoIpv4Resp, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(updatePseudoIpv4Options, "updatePseudoIpv4Options")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"crn":             *zonesSettings.Crn,
		"zone_identifier": *zonesSettings.ZoneIdentifier,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = zonesSettings.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(zonesSettings.Service.Options.URL, `/v1/{crn}/zones/{zone_identifier}/settings/pseudo_ipv4`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range updatePseudoIpv4Options.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("zones_settings", "V1", "UpdatePseudoIpv4")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if updatePseudoIpv4Options.Value != nil {
		body["value"] = updatePseudoIpv4Options.Value
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
	response, err = zonesSettings.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "update_pseudo_ipv4", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalPseudoIpv4Resp)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// GetResponseBuffering : Get response buffering setting
// Get response buffering setting for a zone.
func (zonesSettings *ZonesSettingsV1) GetResponseBuffering(getResponseBufferingOptions *GetResponseBufferingOptions) (result *ResponseBufferingResp, response *core.DetailedResponse, err error) {
	result, response, err = zonesSettings.GetResponseBufferingWithContext(context.Background(), getResponseBufferingOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// GetResponseBufferingWithContext is an alternate form of the GetResponseBuffering method which supports a Context parameter
func (zonesSettings *ZonesSettingsV1) GetResponseBufferingWithContext(ctx context.Context, getResponseBufferingOptions *GetResponseBufferingOptions) (result *ResponseBufferingResp, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getResponseBufferingOptions, "getResponseBufferingOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"crn":             *zonesSettings.Crn,
		"zone_identifier": *zonesSettings.ZoneIdentifier,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = zonesSettings.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(zonesSettings.Service.Options.URL, `/v1/{crn}/zones/{zone_identifier}/settings/response_buffering`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range getResponseBufferingOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("zones_settings", "V1", "GetResponseBuffering")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = zonesSettings.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "get_response_buffering", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalResponseBufferingResp)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// UpdateResponseBuffering : Update response buffering setting
// Update response buffering setting for a zone.
func (zonesSettings *ZonesSettingsV1) UpdateResponseBuffering(updateResponseBufferingOptions *UpdateResponseBufferingOptions) (result *ResponseBufferingResp, response *core.DetailedResponse, err error) {
	result, response, err = zonesSettings.UpdateResponseBufferingWithContext(context.Background(), updateResponseBufferingOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// UpdateResponseBufferingWithContext is an alternate form of the UpdateResponseBuffering method which supports a Context parameter
func (zonesSettings *ZonesSettingsV1) UpdateResponseBufferingWithContext(ctx context.Context, updateResponseBufferingOptions *UpdateResponseBufferingOptions) (result *ResponseBufferingResp, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(updateResponseBufferingOptions, "updateResponseBufferingOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"crn":             *zonesSettings.Crn,
		"zone_identifier": *zonesSettings.ZoneIdentifier,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = zonesSettings.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(zonesSettings.Service.Options.URL, `/v1/{crn}/zones/{zone_identifier}/settings/response_buffering`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range updateResponseBufferingOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("zones_settings", "V1", "UpdateResponseBuffering")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if updateResponseBufferingOptions.Value != nil {
		body["value"] = updateResponseBufferingOptions.Value
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
	response, err = zonesSettings.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "update_response_buffering", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalResponseBufferingResp)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// GetHotlinkProtection : Get hotlink protection setting
// Get hotlink protection setting for a zone.
func (zonesSettings *ZonesSettingsV1) GetHotlinkProtection(getHotlinkProtectionOptions *GetHotlinkProtectionOptions) (result *HotlinkProtectionResp, response *core.DetailedResponse, err error) {
	result, response, err = zonesSettings.GetHotlinkProtectionWithContext(context.Background(), getHotlinkProtectionOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// GetHotlinkProtectionWithContext is an alternate form of the GetHotlinkProtection method which supports a Context parameter
func (zonesSettings *ZonesSettingsV1) GetHotlinkProtectionWithContext(ctx context.Context, getHotlinkProtectionOptions *GetHotlinkProtectionOptions) (result *HotlinkProtectionResp, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getHotlinkProtectionOptions, "getHotlinkProtectionOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"crn":             *zonesSettings.Crn,
		"zone_identifier": *zonesSettings.ZoneIdentifier,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = zonesSettings.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(zonesSettings.Service.Options.URL, `/v1/{crn}/zones/{zone_identifier}/settings/hotlink_protection`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range getHotlinkProtectionOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("zones_settings", "V1", "GetHotlinkProtection")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = zonesSettings.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "get_hotlink_protection", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalHotlinkProtectionResp)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// UpdateHotlinkProtection : Update hotlink protection setting
// Update hotlink protection setting for a zone.
func (zonesSettings *ZonesSettingsV1) UpdateHotlinkProtection(updateHotlinkProtectionOptions *UpdateHotlinkProtectionOptions) (result *HotlinkProtectionResp, response *core.DetailedResponse, err error) {
	result, response, err = zonesSettings.UpdateHotlinkProtectionWithContext(context.Background(), updateHotlinkProtectionOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// UpdateHotlinkProtectionWithContext is an alternate form of the UpdateHotlinkProtection method which supports a Context parameter
func (zonesSettings *ZonesSettingsV1) UpdateHotlinkProtectionWithContext(ctx context.Context, updateHotlinkProtectionOptions *UpdateHotlinkProtectionOptions) (result *HotlinkProtectionResp, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(updateHotlinkProtectionOptions, "updateHotlinkProtectionOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"crn":             *zonesSettings.Crn,
		"zone_identifier": *zonesSettings.ZoneIdentifier,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = zonesSettings.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(zonesSettings.Service.Options.URL, `/v1/{crn}/zones/{zone_identifier}/settings/hotlink_protection`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range updateHotlinkProtectionOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("zones_settings", "V1", "UpdateHotlinkProtection")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if updateHotlinkProtectionOptions.Value != nil {
		body["value"] = updateHotlinkProtectionOptions.Value
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
	response, err = zonesSettings.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "update_hotlink_protection", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalHotlinkProtectionResp)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// GetMaxUpload : Get maximum upload size setting
// Get maximum upload size setting for a zone.
func (zonesSettings *ZonesSettingsV1) GetMaxUpload(getMaxUploadOptions *GetMaxUploadOptions) (result *MaxUploadResp, response *core.DetailedResponse, err error) {
	result, response, err = zonesSettings.GetMaxUploadWithContext(context.Background(), getMaxUploadOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// GetMaxUploadWithContext is an alternate form of the GetMaxUpload method which supports a Context parameter
func (zonesSettings *ZonesSettingsV1) GetMaxUploadWithContext(ctx context.Context, getMaxUploadOptions *GetMaxUploadOptions) (result *MaxUploadResp, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getMaxUploadOptions, "getMaxUploadOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"crn":             *zonesSettings.Crn,
		"zone_identifier": *zonesSettings.ZoneIdentifier,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = zonesSettings.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(zonesSettings.Service.Options.URL, `/v1/{crn}/zones/{zone_identifier}/settings/max_upload`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range getMaxUploadOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("zones_settings", "V1", "GetMaxUpload")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = zonesSettings.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "get_max_upload", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalMaxUploadResp)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// UpdateMaxUpload : Update maximum upload size setting
// Update maximum upload size setting for a zone.
func (zonesSettings *ZonesSettingsV1) UpdateMaxUpload(updateMaxUploadOptions *UpdateMaxUploadOptions) (result *MaxUploadResp, response *core.DetailedResponse, err error) {
	result, response, err = zonesSettings.UpdateMaxUploadWithContext(context.Background(), updateMaxUploadOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// UpdateMaxUploadWithContext is an alternate form of the UpdateMaxUpload method which supports a Context parameter
func (zonesSettings *ZonesSettingsV1) UpdateMaxUploadWithContext(ctx context.Context, updateMaxUploadOptions *UpdateMaxUploadOptions) (result *MaxUploadResp, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(updateMaxUploadOptions, "updateMaxUploadOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"crn":             *zonesSettings.Crn,
		"zone_identifier": *zonesSettings.ZoneIdentifier,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = zonesSettings.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(zonesSettings.Service.Options.URL, `/v1/{crn}/zones/{zone_identifier}/settings/max_upload`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range updateMaxUploadOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("zones_settings", "V1", "UpdateMaxUpload")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if updateMaxUploadOptions.Value != nil {
		body["value"] = updateMaxUploadOptions.Value
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
	response, err = zonesSettings.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "update_max_upload", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalMaxUploadResp)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// GetTlsClientAuth : Get TLS Client Auth setting
// Get TLS Client Auth setting for a zone.
func (zonesSettings *ZonesSettingsV1) GetTlsClientAuth(getTlsClientAuthOptions *GetTlsClientAuthOptions) (result *TlsClientAuthResp, response *core.DetailedResponse, err error) {
	result, response, err = zonesSettings.GetTlsClientAuthWithContext(context.Background(), getTlsClientAuthOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// GetTlsClientAuthWithContext is an alternate form of the GetTlsClientAuth method which supports a Context parameter
func (zonesSettings *ZonesSettingsV1) GetTlsClientAuthWithContext(ctx context.Context, getTlsClientAuthOptions *GetTlsClientAuthOptions) (result *TlsClientAuthResp, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getTlsClientAuthOptions, "getTlsClientAuthOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"crn":             *zonesSettings.Crn,
		"zone_identifier": *zonesSettings.ZoneIdentifier,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = zonesSettings.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(zonesSettings.Service.Options.URL, `/v1/{crn}/zones/{zone_identifier}/settings/tls_client_auth`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range getTlsClientAuthOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("zones_settings", "V1", "GetTlsClientAuth")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = zonesSettings.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "get_tls_client_auth", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalTlsClientAuthResp)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// UpdateTlsClientAuth : Update TLS Client Auth setting
// Update TLS Client Auth setting for a zone.
func (zonesSettings *ZonesSettingsV1) UpdateTlsClientAuth(updateTlsClientAuthOptions *UpdateTlsClientAuthOptions) (result *TlsClientAuthResp, response *core.DetailedResponse, err error) {
	result, response, err = zonesSettings.UpdateTlsClientAuthWithContext(context.Background(), updateTlsClientAuthOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// UpdateTlsClientAuthWithContext is an alternate form of the UpdateTlsClientAuth method which supports a Context parameter
func (zonesSettings *ZonesSettingsV1) UpdateTlsClientAuthWithContext(ctx context.Context, updateTlsClientAuthOptions *UpdateTlsClientAuthOptions) (result *TlsClientAuthResp, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(updateTlsClientAuthOptions, "updateTlsClientAuthOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"crn":             *zonesSettings.Crn,
		"zone_identifier": *zonesSettings.ZoneIdentifier,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = zonesSettings.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(zonesSettings.Service.Options.URL, `/v1/{crn}/zones/{zone_identifier}/settings/tls_client_auth`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range updateTlsClientAuthOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("zones_settings", "V1", "UpdateTlsClientAuth")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if updateTlsClientAuthOptions.Value != nil {
		body["value"] = updateTlsClientAuthOptions.Value
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
	response, err = zonesSettings.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "update_tls_client_auth", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalTlsClientAuthResp)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// GetBrotli : Get brotli setting
// Get brotli setting for a zone.
func (zonesSettings *ZonesSettingsV1) GetBrotli(getBrotliOptions *GetBrotliOptions) (result *BrotliResp, response *core.DetailedResponse, err error) {
	result, response, err = zonesSettings.GetBrotliWithContext(context.Background(), getBrotliOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// GetBrotliWithContext is an alternate form of the GetBrotli method which supports a Context parameter
func (zonesSettings *ZonesSettingsV1) GetBrotliWithContext(ctx context.Context, getBrotliOptions *GetBrotliOptions) (result *BrotliResp, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getBrotliOptions, "getBrotliOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"crn":             *zonesSettings.Crn,
		"zone_identifier": *zonesSettings.ZoneIdentifier,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = zonesSettings.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(zonesSettings.Service.Options.URL, `/v1/{crn}/zones/{zone_identifier}/settings/brotli`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range getBrotliOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("zones_settings", "V1", "GetBrotli")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = zonesSettings.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "get_brotli", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalBrotliResp)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// UpdateBrotli : Update brotli setting
// Update brotli setting for a zone.
func (zonesSettings *ZonesSettingsV1) UpdateBrotli(updateBrotliOptions *UpdateBrotliOptions) (result *BrotliResp, response *core.DetailedResponse, err error) {
	result, response, err = zonesSettings.UpdateBrotliWithContext(context.Background(), updateBrotliOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// UpdateBrotliWithContext is an alternate form of the UpdateBrotli method which supports a Context parameter
func (zonesSettings *ZonesSettingsV1) UpdateBrotliWithContext(ctx context.Context, updateBrotliOptions *UpdateBrotliOptions) (result *BrotliResp, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(updateBrotliOptions, "updateBrotliOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"crn":             *zonesSettings.Crn,
		"zone_identifier": *zonesSettings.ZoneIdentifier,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = zonesSettings.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(zonesSettings.Service.Options.URL, `/v1/{crn}/zones/{zone_identifier}/settings/brotli`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range updateBrotliOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("zones_settings", "V1", "UpdateBrotli")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if updateBrotliOptions.Value != nil {
		body["value"] = updateBrotliOptions.Value
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
	response, err = zonesSettings.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "update_brotli", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalBrotliResp)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// GetProxyReadTimeout : Get proxy read timeout setting
// Get proxy read timeout setting for a zone.
func (zonesSettings *ZonesSettingsV1) GetProxyReadTimeout(getProxyReadTimeoutOptions *GetProxyReadTimeoutOptions) (result *ProxyReadTimeoutResp, response *core.DetailedResponse, err error) {
	result, response, err = zonesSettings.GetProxyReadTimeoutWithContext(context.Background(), getProxyReadTimeoutOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// GetProxyReadTimeoutWithContext is an alternate form of the GetProxyReadTimeout method which supports a Context parameter
func (zonesSettings *ZonesSettingsV1) GetProxyReadTimeoutWithContext(ctx context.Context, getProxyReadTimeoutOptions *GetProxyReadTimeoutOptions) (result *ProxyReadTimeoutResp, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getProxyReadTimeoutOptions, "getProxyReadTimeoutOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"crn":             *zonesSettings.Crn,
		"zone_identifier": *zonesSettings.ZoneIdentifier,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = zonesSettings.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(zonesSettings.Service.Options.URL, `/v1/{crn}/zones/{zone_identifier}/settings/proxy_read_timeout`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range getProxyReadTimeoutOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("zones_settings", "V1", "GetProxyReadTimeout")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = zonesSettings.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "get_proxy_read_timeout", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalProxyReadTimeoutResp)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// UpdateProxyReadTimeout : Update proxy read timeout setting
// Update proxy read timeout setting for a zone.
func (zonesSettings *ZonesSettingsV1) UpdateProxyReadTimeout(updateProxyReadTimeoutOptions *UpdateProxyReadTimeoutOptions) (result *ProxyReadTimeoutResp, response *core.DetailedResponse, err error) {
	result, response, err = zonesSettings.UpdateProxyReadTimeoutWithContext(context.Background(), updateProxyReadTimeoutOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// UpdateProxyReadTimeoutWithContext is an alternate form of the UpdateProxyReadTimeout method which supports a Context parameter
func (zonesSettings *ZonesSettingsV1) UpdateProxyReadTimeoutWithContext(ctx context.Context, updateProxyReadTimeoutOptions *UpdateProxyReadTimeoutOptions) (result *ProxyReadTimeoutResp, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(updateProxyReadTimeoutOptions, "updateProxyReadTimeoutOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"crn":             *zonesSettings.Crn,
		"zone_identifier": *zonesSettings.ZoneIdentifier,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = zonesSettings.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(zonesSettings.Service.Options.URL, `/v1/{crn}/zones/{zone_identifier}/settings/proxy_read_timeout`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range updateProxyReadTimeoutOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("zones_settings", "V1", "UpdateProxyReadTimeout")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if updateProxyReadTimeoutOptions.Value != nil {
		body["value"] = updateProxyReadTimeoutOptions.Value
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
	response, err = zonesSettings.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "update_proxy_read_timeout", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalProxyReadTimeoutResp)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// GetBrowserCheck : Get browser check setting
// Get browser check setting for a zone.
func (zonesSettings *ZonesSettingsV1) GetBrowserCheck(getBrowserCheckOptions *GetBrowserCheckOptions) (result *BrowserCheckResp, response *core.DetailedResponse, err error) {
	result, response, err = zonesSettings.GetBrowserCheckWithContext(context.Background(), getBrowserCheckOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// GetBrowserCheckWithContext is an alternate form of the GetBrowserCheck method which supports a Context parameter
func (zonesSettings *ZonesSettingsV1) GetBrowserCheckWithContext(ctx context.Context, getBrowserCheckOptions *GetBrowserCheckOptions) (result *BrowserCheckResp, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getBrowserCheckOptions, "getBrowserCheckOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"crn":             *zonesSettings.Crn,
		"zone_identifier": *zonesSettings.ZoneIdentifier,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = zonesSettings.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(zonesSettings.Service.Options.URL, `/v1/{crn}/zones/{zone_identifier}/settings/browser_check`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range getBrowserCheckOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("zones_settings", "V1", "GetBrowserCheck")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = zonesSettings.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "get_browser_check", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalBrowserCheckResp)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// UpdateBrowserCheck : Update browser check setting
// Update browser check setting for a zone.
func (zonesSettings *ZonesSettingsV1) UpdateBrowserCheck(updateBrowserCheckOptions *UpdateBrowserCheckOptions) (result *BrowserCheckResp, response *core.DetailedResponse, err error) {
	result, response, err = zonesSettings.UpdateBrowserCheckWithContext(context.Background(), updateBrowserCheckOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// UpdateBrowserCheckWithContext is an alternate form of the UpdateBrowserCheck method which supports a Context parameter
func (zonesSettings *ZonesSettingsV1) UpdateBrowserCheckWithContext(ctx context.Context, updateBrowserCheckOptions *UpdateBrowserCheckOptions) (result *BrowserCheckResp, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(updateBrowserCheckOptions, "updateBrowserCheckOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"crn":             *zonesSettings.Crn,
		"zone_identifier": *zonesSettings.ZoneIdentifier,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = zonesSettings.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(zonesSettings.Service.Options.URL, `/v1/{crn}/zones/{zone_identifier}/settings/browser_check`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range updateBrowserCheckOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("zones_settings", "V1", "UpdateBrowserCheck")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if updateBrowserCheckOptions.Value != nil {
		body["value"] = updateBrowserCheckOptions.Value
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
	response, err = zonesSettings.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "update_browser_check", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalBrowserCheckResp)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// GetEnableErrorPagesOn : Get enable error pages on setting
// Get enable error pages on setting for a zone.
func (zonesSettings *ZonesSettingsV1) GetEnableErrorPagesOn(getEnableErrorPagesOnOptions *GetEnableErrorPagesOnOptions) (result *OriginErrorPagePassThruResp, response *core.DetailedResponse, err error) {
	result, response, err = zonesSettings.GetEnableErrorPagesOnWithContext(context.Background(), getEnableErrorPagesOnOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// GetEnableErrorPagesOnWithContext is an alternate form of the GetEnableErrorPagesOn method which supports a Context parameter
func (zonesSettings *ZonesSettingsV1) GetEnableErrorPagesOnWithContext(ctx context.Context, getEnableErrorPagesOnOptions *GetEnableErrorPagesOnOptions) (result *OriginErrorPagePassThruResp, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getEnableErrorPagesOnOptions, "getEnableErrorPagesOnOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"crn":             *zonesSettings.Crn,
		"zone_identifier": *zonesSettings.ZoneIdentifier,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = zonesSettings.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(zonesSettings.Service.Options.URL, `/v1/{crn}/zones/{zone_identifier}/settings/origin_error_page_pass_thru`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range getEnableErrorPagesOnOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("zones_settings", "V1", "GetEnableErrorPagesOn")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = zonesSettings.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "get_enable_error_pages_on", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalOriginErrorPagePassThruResp)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// UpdateEnableErrorPagesOn : Update enable error pages on setting
// Update enable error pages on setting for a zone.
func (zonesSettings *ZonesSettingsV1) UpdateEnableErrorPagesOn(updateEnableErrorPagesOnOptions *UpdateEnableErrorPagesOnOptions) (result *OriginErrorPagePassThruResp, response *core.DetailedResponse, err error) {
	result, response, err = zonesSettings.UpdateEnableErrorPagesOnWithContext(context.Background(), updateEnableErrorPagesOnOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// UpdateEnableErrorPagesOnWithContext is an alternate form of the UpdateEnableErrorPagesOn method which supports a Context parameter
func (zonesSettings *ZonesSettingsV1) UpdateEnableErrorPagesOnWithContext(ctx context.Context, updateEnableErrorPagesOnOptions *UpdateEnableErrorPagesOnOptions) (result *OriginErrorPagePassThruResp, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(updateEnableErrorPagesOnOptions, "updateEnableErrorPagesOnOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"crn":             *zonesSettings.Crn,
		"zone_identifier": *zonesSettings.ZoneIdentifier,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = zonesSettings.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(zonesSettings.Service.Options.URL, `/v1/{crn}/zones/{zone_identifier}/settings/origin_error_page_pass_thru`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range updateEnableErrorPagesOnOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("zones_settings", "V1", "UpdateEnableErrorPagesOn")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if updateEnableErrorPagesOnOptions.Value != nil {
		body["value"] = updateEnableErrorPagesOnOptions.Value
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
	response, err = zonesSettings.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "update_enable_error_pages_on", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalOriginErrorPagePassThruResp)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// GetWebApplicationFirewall : Get web application firewall setting
// Get web application firewall setting for a zone.
func (zonesSettings *ZonesSettingsV1) GetWebApplicationFirewall(getWebApplicationFirewallOptions *GetWebApplicationFirewallOptions) (result *WafResp, response *core.DetailedResponse, err error) {
	result, response, err = zonesSettings.GetWebApplicationFirewallWithContext(context.Background(), getWebApplicationFirewallOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// GetWebApplicationFirewallWithContext is an alternate form of the GetWebApplicationFirewall method which supports a Context parameter
func (zonesSettings *ZonesSettingsV1) GetWebApplicationFirewallWithContext(ctx context.Context, getWebApplicationFirewallOptions *GetWebApplicationFirewallOptions) (result *WafResp, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getWebApplicationFirewallOptions, "getWebApplicationFirewallOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"crn":             *zonesSettings.Crn,
		"zone_identifier": *zonesSettings.ZoneIdentifier,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = zonesSettings.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(zonesSettings.Service.Options.URL, `/v1/{crn}/zones/{zone_identifier}/settings/waf`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range getWebApplicationFirewallOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("zones_settings", "V1", "GetWebApplicationFirewall")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = zonesSettings.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "get_web_application_firewall", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalWafResp)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// UpdateWebApplicationFirewall : Update web application firewall setting
// A Web Application Firewall (WAF) blocks requests that contain malicious content.
func (zonesSettings *ZonesSettingsV1) UpdateWebApplicationFirewall(updateWebApplicationFirewallOptions *UpdateWebApplicationFirewallOptions) (result *WafResp, response *core.DetailedResponse, err error) {
	result, response, err = zonesSettings.UpdateWebApplicationFirewallWithContext(context.Background(), updateWebApplicationFirewallOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// UpdateWebApplicationFirewallWithContext is an alternate form of the UpdateWebApplicationFirewall method which supports a Context parameter
func (zonesSettings *ZonesSettingsV1) UpdateWebApplicationFirewallWithContext(ctx context.Context, updateWebApplicationFirewallOptions *UpdateWebApplicationFirewallOptions) (result *WafResp, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(updateWebApplicationFirewallOptions, "updateWebApplicationFirewallOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"crn":             *zonesSettings.Crn,
		"zone_identifier": *zonesSettings.ZoneIdentifier,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = zonesSettings.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(zonesSettings.Service.Options.URL, `/v1/{crn}/zones/{zone_identifier}/settings/waf`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range updateWebApplicationFirewallOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("zones_settings", "V1", "UpdateWebApplicationFirewall")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if updateWebApplicationFirewallOptions.Value != nil {
		body["value"] = updateWebApplicationFirewallOptions.Value
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
	response, err = zonesSettings.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "update_web_application_firewall", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalWafResp)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// GetCiphers : Get ciphers setting
// Get ciphers setting for a zone.
func (zonesSettings *ZonesSettingsV1) GetCiphers(getCiphersOptions *GetCiphersOptions) (result *CiphersResp, response *core.DetailedResponse, err error) {
	result, response, err = zonesSettings.GetCiphersWithContext(context.Background(), getCiphersOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// GetCiphersWithContext is an alternate form of the GetCiphers method which supports a Context parameter
func (zonesSettings *ZonesSettingsV1) GetCiphersWithContext(ctx context.Context, getCiphersOptions *GetCiphersOptions) (result *CiphersResp, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getCiphersOptions, "getCiphersOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"crn":             *zonesSettings.Crn,
		"zone_identifier": *zonesSettings.ZoneIdentifier,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = zonesSettings.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(zonesSettings.Service.Options.URL, `/v1/{crn}/zones/{zone_identifier}/settings/ciphers`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range getCiphersOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("zones_settings", "V1", "GetCiphers")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = zonesSettings.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "get_ciphers", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCiphersResp)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// UpdateCiphers : Update ciphers setting
// Update ciphers setting for a zone.
func (zonesSettings *ZonesSettingsV1) UpdateCiphers(updateCiphersOptions *UpdateCiphersOptions) (result *CiphersResp, response *core.DetailedResponse, err error) {
	result, response, err = zonesSettings.UpdateCiphersWithContext(context.Background(), updateCiphersOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// UpdateCiphersWithContext is an alternate form of the UpdateCiphers method which supports a Context parameter
func (zonesSettings *ZonesSettingsV1) UpdateCiphersWithContext(ctx context.Context, updateCiphersOptions *UpdateCiphersOptions) (result *CiphersResp, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(updateCiphersOptions, "updateCiphersOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"crn":             *zonesSettings.Crn,
		"zone_identifier": *zonesSettings.ZoneIdentifier,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = zonesSettings.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(zonesSettings.Service.Options.URL, `/v1/{crn}/zones/{zone_identifier}/settings/ciphers`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range updateCiphersOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("zones_settings", "V1", "UpdateCiphers")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if updateCiphersOptions.Value != nil {
		body["value"] = updateCiphersOptions.Value
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
	response, err = zonesSettings.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "update_ciphers", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCiphersResp)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// GetOriginMaxHttpVersion : Get origin max http version setting
// Get origin max http version setting for a zone.
func (zonesSettings *ZonesSettingsV1) GetOriginMaxHttpVersion(getOriginMaxHttpVersionOptions *GetOriginMaxHttpVersionOptions) (result *OriginMaxHttpVersionResp, response *core.DetailedResponse, err error) {
	result, response, err = zonesSettings.GetOriginMaxHttpVersionWithContext(context.Background(), getOriginMaxHttpVersionOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// GetOriginMaxHttpVersionWithContext is an alternate form of the GetOriginMaxHttpVersion method which supports a Context parameter
func (zonesSettings *ZonesSettingsV1) GetOriginMaxHttpVersionWithContext(ctx context.Context, getOriginMaxHttpVersionOptions *GetOriginMaxHttpVersionOptions) (result *OriginMaxHttpVersionResp, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getOriginMaxHttpVersionOptions, "getOriginMaxHttpVersionOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"crn":             *zonesSettings.Crn,
		"zone_identifier": *zonesSettings.ZoneIdentifier,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = zonesSettings.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(zonesSettings.Service.Options.URL, `/v1/{crn}/zones/{zone_identifier}/settings/origin_max_http_version`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range getOriginMaxHttpVersionOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("zones_settings", "V1", "GetOriginMaxHttpVersion")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = zonesSettings.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "get_origin_max_http_version", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalOriginMaxHttpVersionResp)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// UpdateOriginMaxHttpVersion : Update origin max http version setting
// Update origin max http version setting for a zone.
func (zonesSettings *ZonesSettingsV1) UpdateOriginMaxHttpVersion(updateOriginMaxHttpVersionOptions *UpdateOriginMaxHttpVersionOptions) (result *OriginMaxHttpVersionResp, response *core.DetailedResponse, err error) {
	result, response, err = zonesSettings.UpdateOriginMaxHttpVersionWithContext(context.Background(), updateOriginMaxHttpVersionOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// UpdateOriginMaxHttpVersionWithContext is an alternate form of the UpdateOriginMaxHttpVersion method which supports a Context parameter
func (zonesSettings *ZonesSettingsV1) UpdateOriginMaxHttpVersionWithContext(ctx context.Context, updateOriginMaxHttpVersionOptions *UpdateOriginMaxHttpVersionOptions) (result *OriginMaxHttpVersionResp, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(updateOriginMaxHttpVersionOptions, "updateOriginMaxHttpVersionOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"crn":             *zonesSettings.Crn,
		"zone_identifier": *zonesSettings.ZoneIdentifier,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = zonesSettings.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(zonesSettings.Service.Options.URL, `/v1/{crn}/zones/{zone_identifier}/settings/origin_max_http_version`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range updateOriginMaxHttpVersionOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("zones_settings", "V1", "UpdateOriginMaxHttpVersion")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if updateOriginMaxHttpVersionOptions.Value != nil {
		body["value"] = updateOriginMaxHttpVersionOptions.Value
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
	response, err = zonesSettings.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "update_origin_max_http_version", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalOriginMaxHttpVersionResp)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// GetOriginPostQuantumEncryption : Get origin post quantum encryption setting
// Get origin post quantum encryption setting for a zone.
func (zonesSettings *ZonesSettingsV1) GetOriginPostQuantumEncryption(getOriginPostQuantumEncryptionOptions *GetOriginPostQuantumEncryptionOptions) (result *OriginPostQuantumEncryptionResp, response *core.DetailedResponse, err error) {
	result, response, err = zonesSettings.GetOriginPostQuantumEncryptionWithContext(context.Background(), getOriginPostQuantumEncryptionOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// GetOriginPostQuantumEncryptionWithContext is an alternate form of the GetOriginPostQuantumEncryption method which supports a Context parameter
func (zonesSettings *ZonesSettingsV1) GetOriginPostQuantumEncryptionWithContext(ctx context.Context, getOriginPostQuantumEncryptionOptions *GetOriginPostQuantumEncryptionOptions) (result *OriginPostQuantumEncryptionResp, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getOriginPostQuantumEncryptionOptions, "getOriginPostQuantumEncryptionOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"crn":             *zonesSettings.Crn,
		"zone_identifier": *zonesSettings.ZoneIdentifier,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = zonesSettings.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(zonesSettings.Service.Options.URL, `/v1/{crn}/zones/{zone_identifier}/cache/origin_post_quantum_encryption`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range getOriginPostQuantumEncryptionOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("zones_settings", "V1", "GetOriginPostQuantumEncryption")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = zonesSettings.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "get_origin_post_quantum_encryption", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalOriginPostQuantumEncryptionResp)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// UpdateOriginPostQuantumEncryption : Update origin post quantum encryption setting
// Update origin post quantum encryption setting for a zone.
func (zonesSettings *ZonesSettingsV1) UpdateOriginPostQuantumEncryption(updateOriginPostQuantumEncryptionOptions *UpdateOriginPostQuantumEncryptionOptions) (result *OriginPostQuantumEncryptionResp, response *core.DetailedResponse, err error) {
	result, response, err = zonesSettings.UpdateOriginPostQuantumEncryptionWithContext(context.Background(), updateOriginPostQuantumEncryptionOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// UpdateOriginPostQuantumEncryptionWithContext is an alternate form of the UpdateOriginPostQuantumEncryption method which supports a Context parameter
func (zonesSettings *ZonesSettingsV1) UpdateOriginPostQuantumEncryptionWithContext(ctx context.Context, updateOriginPostQuantumEncryptionOptions *UpdateOriginPostQuantumEncryptionOptions) (result *OriginPostQuantumEncryptionResp, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(updateOriginPostQuantumEncryptionOptions, "updateOriginPostQuantumEncryptionOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"crn":             *zonesSettings.Crn,
		"zone_identifier": *zonesSettings.ZoneIdentifier,
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = zonesSettings.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(zonesSettings.Service.Options.URL, `/v1/{crn}/zones/{zone_identifier}/cache/origin_post_quantum_encryption`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range updateOriginPostQuantumEncryptionOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("zones_settings", "V1", "UpdateOriginPostQuantumEncryption")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if updateOriginPostQuantumEncryptionOptions.Value != nil {
		body["value"] = updateOriginPostQuantumEncryptionOptions.Value
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
	response, err = zonesSettings.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "update_origin_post_quantum_encryption", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalOriginPostQuantumEncryptionResp)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// GetLogRetention : Retrieves the current setting for log retention
// Get the current setting for log retention. This setting is available for Enterprise plans only. If this setting is
// turned on, then logs from the cloud edge are retained for the customers domain. Otherwise they will be discarded.
func (zonesSettings *ZonesSettingsV1) GetLogRetention(getLogRetentionOptions *GetLogRetentionOptions) (result *LogRetentionResp, response *core.DetailedResponse, err error) {
	result, response, err = zonesSettings.GetLogRetentionWithContext(context.Background(), getLogRetentionOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// GetLogRetentionWithContext is an alternate form of the GetLogRetention method which supports a Context parameter
func (zonesSettings *ZonesSettingsV1) GetLogRetentionWithContext(ctx context.Context, getLogRetentionOptions *GetLogRetentionOptions) (result *LogRetentionResp, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getLogRetentionOptions, "getLogRetentionOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(getLogRetentionOptions, "getLogRetentionOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"crn":             *getLogRetentionOptions.Crn,
		"zone_identifier": *getLogRetentionOptions.ZoneIdentifier,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = zonesSettings.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(zonesSettings.Service.Options.URL, `/v1/{crn}/zones/{zone_identifier}/logs/retention`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range getLogRetentionOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("zones_settings", "V1", "GetLogRetention")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = zonesSettings.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "get_log_retention", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalLogRetentionResp)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// UpdateLogRetention : Toggles the current setting for log retention
// Toggles the current setting for log retention.
func (zonesSettings *ZonesSettingsV1) UpdateLogRetention(updateLogRetentionOptions *UpdateLogRetentionOptions) (result *LogRetentionResp, response *core.DetailedResponse, err error) {
	result, response, err = zonesSettings.UpdateLogRetentionWithContext(context.Background(), updateLogRetentionOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// UpdateLogRetentionWithContext is an alternate form of the UpdateLogRetention method which supports a Context parameter
func (zonesSettings *ZonesSettingsV1) UpdateLogRetentionWithContext(ctx context.Context, updateLogRetentionOptions *UpdateLogRetentionOptions) (result *LogRetentionResp, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateLogRetentionOptions, "updateLogRetentionOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(updateLogRetentionOptions, "updateLogRetentionOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"crn":             *updateLogRetentionOptions.Crn,
		"zone_identifier": *updateLogRetentionOptions.ZoneIdentifier,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = zonesSettings.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(zonesSettings.Service.Options.URL, `/v1/{crn}/zones/{zone_identifier}/logs/retention`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range updateLogRetentionOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("zones_settings", "V1", "UpdateLogRetention")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if updateLogRetentionOptions.Flag != nil {
		body["flag"] = updateLogRetentionOptions.Flag
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
	response, err = zonesSettings.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "update_log_retention", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalLogRetentionResp)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}
func getServiceComponentInfo() *core.ProblemComponent {
	return core.NewProblemComponent(DefaultServiceName, "1.0.1")
}

// AlwaysUseHttpsRespResult : Container for response information.
type AlwaysUseHttpsRespResult struct {
	// ID.
	ID *string `json:"id" validate:"required"`

	// Value.
	Value *string `json:"value" validate:"required"`

	// Editable.
	Editable *bool `json:"editable" validate:"required"`

	// Modified date.
	ModifiedOn *strfmt.DateTime `json:"modified_on" validate:"required"`
}

// UnmarshalAlwaysUseHttpsRespResult unmarshals an instance of AlwaysUseHttpsRespResult from the specified map of raw messages.
func UnmarshalAlwaysUseHttpsRespResult(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AlwaysUseHttpsRespResult)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		err = core.SDKErrorf(err, "", "id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "value", &obj.Value)
	if err != nil {
		err = core.SDKErrorf(err, "", "value-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "editable", &obj.Editable)
	if err != nil {
		err = core.SDKErrorf(err, "", "editable-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "modified_on", &obj.ModifiedOn)
	if err != nil {
		err = core.SDKErrorf(err, "", "modified_on-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AutomaticHttpsRewritesRespResult : Container for response information.
type AutomaticHttpsRewritesRespResult struct {
	// ID.
	ID *string `json:"id" validate:"required"`

	// Value.
	Value *string `json:"value" validate:"required"`

	// Editable.
	Editable *bool `json:"editable" validate:"required"`

	// Modified date.
	ModifiedOn *strfmt.DateTime `json:"modified_on" validate:"required"`
}

// UnmarshalAutomaticHttpsRewritesRespResult unmarshals an instance of AutomaticHttpsRewritesRespResult from the specified map of raw messages.
func UnmarshalAutomaticHttpsRewritesRespResult(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AutomaticHttpsRewritesRespResult)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		err = core.SDKErrorf(err, "", "id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "value", &obj.Value)
	if err != nil {
		err = core.SDKErrorf(err, "", "value-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "editable", &obj.Editable)
	if err != nil {
		err = core.SDKErrorf(err, "", "editable-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "modified_on", &obj.ModifiedOn)
	if err != nil {
		err = core.SDKErrorf(err, "", "modified_on-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// BrotliRespResult : Container for response information.
type BrotliRespResult struct {
	// ID.
	ID *string `json:"id" validate:"required"`

	// Value.
	Value *string `json:"value" validate:"required"`

	// Editable.
	Editable *bool `json:"editable" validate:"required"`

	// Modified date.
	ModifiedOn *strfmt.DateTime `json:"modified_on" validate:"required"`
}

// UnmarshalBrotliRespResult unmarshals an instance of BrotliRespResult from the specified map of raw messages.
func UnmarshalBrotliRespResult(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(BrotliRespResult)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		err = core.SDKErrorf(err, "", "id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "value", &obj.Value)
	if err != nil {
		err = core.SDKErrorf(err, "", "value-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "editable", &obj.Editable)
	if err != nil {
		err = core.SDKErrorf(err, "", "editable-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "modified_on", &obj.ModifiedOn)
	if err != nil {
		err = core.SDKErrorf(err, "", "modified_on-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// BrowserCheckRespResult : Container for response information.
type BrowserCheckRespResult struct {
	// ID.
	ID *string `json:"id" validate:"required"`

	// Value.
	Value *string `json:"value" validate:"required"`

	// Editable.
	Editable *bool `json:"editable" validate:"required"`

	// Modified date.
	ModifiedOn *strfmt.DateTime `json:"modified_on" validate:"required"`
}

// UnmarshalBrowserCheckRespResult unmarshals an instance of BrowserCheckRespResult from the specified map of raw messages.
func UnmarshalBrowserCheckRespResult(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(BrowserCheckRespResult)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		err = core.SDKErrorf(err, "", "id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "value", &obj.Value)
	if err != nil {
		err = core.SDKErrorf(err, "", "value-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "editable", &obj.Editable)
	if err != nil {
		err = core.SDKErrorf(err, "", "editable-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "modified_on", &obj.ModifiedOn)
	if err != nil {
		err = core.SDKErrorf(err, "", "modified_on-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ChallengeTtlRespResult : Container for response information.
type ChallengeTtlRespResult struct {
	// ID.
	ID *string `json:"id" validate:"required"`

	// Value.
	Value *int64 `json:"value" validate:"required"`

	// Editable.
	Editable *bool `json:"editable" validate:"required"`

	// Modified date.
	ModifiedOn *strfmt.DateTime `json:"modified_on" validate:"required"`
}

// UnmarshalChallengeTtlRespResult unmarshals an instance of ChallengeTtlRespResult from the specified map of raw messages.
func UnmarshalChallengeTtlRespResult(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ChallengeTtlRespResult)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		err = core.SDKErrorf(err, "", "id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "value", &obj.Value)
	if err != nil {
		err = core.SDKErrorf(err, "", "value-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "editable", &obj.Editable)
	if err != nil {
		err = core.SDKErrorf(err, "", "editable-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "modified_on", &obj.ModifiedOn)
	if err != nil {
		err = core.SDKErrorf(err, "", "modified_on-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CiphersRespResult : Container for response information.
type CiphersRespResult struct {
	// ID.
	ID *string `json:"id" validate:"required"`

	// Value.
	Value []string `json:"value" validate:"required"`

	// Editable.
	Editable *bool `json:"editable" validate:"required"`

	// Modified date.
	ModifiedOn *strfmt.DateTime `json:"modified_on" validate:"required"`
}

// UnmarshalCiphersRespResult unmarshals an instance of CiphersRespResult from the specified map of raw messages.
func UnmarshalCiphersRespResult(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CiphersRespResult)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		err = core.SDKErrorf(err, "", "id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "value", &obj.Value)
	if err != nil {
		err = core.SDKErrorf(err, "", "value-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "editable", &obj.Editable)
	if err != nil {
		err = core.SDKErrorf(err, "", "editable-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "modified_on", &obj.ModifiedOn)
	if err != nil {
		err = core.SDKErrorf(err, "", "modified_on-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetAlwaysUseHttpsOptions : The GetAlwaysUseHttps options.
type GetAlwaysUseHttpsOptions struct {

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewGetAlwaysUseHttpsOptions : Instantiate GetAlwaysUseHttpsOptions
func (*ZonesSettingsV1) NewGetAlwaysUseHttpsOptions() *GetAlwaysUseHttpsOptions {
	return &GetAlwaysUseHttpsOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *GetAlwaysUseHttpsOptions) SetHeaders(param map[string]string) *GetAlwaysUseHttpsOptions {
	options.Headers = param
	return options
}

// GetAutomaticHttpsRewritesOptions : The GetAutomaticHttpsRewrites options.
type GetAutomaticHttpsRewritesOptions struct {

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewGetAutomaticHttpsRewritesOptions : Instantiate GetAutomaticHttpsRewritesOptions
func (*ZonesSettingsV1) NewGetAutomaticHttpsRewritesOptions() *GetAutomaticHttpsRewritesOptions {
	return &GetAutomaticHttpsRewritesOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *GetAutomaticHttpsRewritesOptions) SetHeaders(param map[string]string) *GetAutomaticHttpsRewritesOptions {
	options.Headers = param
	return options
}

// GetBrotliOptions : The GetBrotli options.
type GetBrotliOptions struct {

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewGetBrotliOptions : Instantiate GetBrotliOptions
func (*ZonesSettingsV1) NewGetBrotliOptions() *GetBrotliOptions {
	return &GetBrotliOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *GetBrotliOptions) SetHeaders(param map[string]string) *GetBrotliOptions {
	options.Headers = param
	return options
}

// GetBrowserCheckOptions : The GetBrowserCheck options.
type GetBrowserCheckOptions struct {

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewGetBrowserCheckOptions : Instantiate GetBrowserCheckOptions
func (*ZonesSettingsV1) NewGetBrowserCheckOptions() *GetBrowserCheckOptions {
	return &GetBrowserCheckOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *GetBrowserCheckOptions) SetHeaders(param map[string]string) *GetBrowserCheckOptions {
	options.Headers = param
	return options
}

// GetChallengeTtlOptions : The GetChallengeTTL options.
type GetChallengeTtlOptions struct {

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewGetChallengeTtlOptions : Instantiate GetChallengeTtlOptions
func (*ZonesSettingsV1) NewGetChallengeTtlOptions() *GetChallengeTtlOptions {
	return &GetChallengeTtlOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *GetChallengeTtlOptions) SetHeaders(param map[string]string) *GetChallengeTtlOptions {
	options.Headers = param
	return options
}

// GetCiphersOptions : The GetCiphers options.
type GetCiphersOptions struct {

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewGetCiphersOptions : Instantiate GetCiphersOptions
func (*ZonesSettingsV1) NewGetCiphersOptions() *GetCiphersOptions {
	return &GetCiphersOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *GetCiphersOptions) SetHeaders(param map[string]string) *GetCiphersOptions {
	options.Headers = param
	return options
}

// GetEnableErrorPagesOnOptions : The GetEnableErrorPagesOn options.
type GetEnableErrorPagesOnOptions struct {

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewGetEnableErrorPagesOnOptions : Instantiate GetEnableErrorPagesOnOptions
func (*ZonesSettingsV1) NewGetEnableErrorPagesOnOptions() *GetEnableErrorPagesOnOptions {
	return &GetEnableErrorPagesOnOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *GetEnableErrorPagesOnOptions) SetHeaders(param map[string]string) *GetEnableErrorPagesOnOptions {
	options.Headers = param
	return options
}

// GetHotlinkProtectionOptions : The GetHotlinkProtection options.
type GetHotlinkProtectionOptions struct {

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewGetHotlinkProtectionOptions : Instantiate GetHotlinkProtectionOptions
func (*ZonesSettingsV1) NewGetHotlinkProtectionOptions() *GetHotlinkProtectionOptions {
	return &GetHotlinkProtectionOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *GetHotlinkProtectionOptions) SetHeaders(param map[string]string) *GetHotlinkProtectionOptions {
	options.Headers = param
	return options
}

// GetHttp2Options : The GetHttp2 options.
type GetHttp2Options struct {

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewGetHttp2Options : Instantiate GetHttp2Options
func (*ZonesSettingsV1) NewGetHttp2Options() *GetHttp2Options {
	return &GetHttp2Options{}
}

// SetHeaders : Allow user to set Headers
func (options *GetHttp2Options) SetHeaders(param map[string]string) *GetHttp2Options {
	options.Headers = param
	return options
}

// GetHttp3Options : The GetHttp3 options.
type GetHttp3Options struct {

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewGetHttp3Options : Instantiate GetHttp3Options
func (*ZonesSettingsV1) NewGetHttp3Options() *GetHttp3Options {
	return &GetHttp3Options{}
}

// SetHeaders : Allow user to set Headers
func (options *GetHttp3Options) SetHeaders(param map[string]string) *GetHttp3Options {
	options.Headers = param
	return options
}

// GetImageLoadOptimizationOptions : The GetImageLoadOptimization options.
type GetImageLoadOptimizationOptions struct {

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewGetImageLoadOptimizationOptions : Instantiate GetImageLoadOptimizationOptions
func (*ZonesSettingsV1) NewGetImageLoadOptimizationOptions() *GetImageLoadOptimizationOptions {
	return &GetImageLoadOptimizationOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *GetImageLoadOptimizationOptions) SetHeaders(param map[string]string) *GetImageLoadOptimizationOptions {
	options.Headers = param
	return options
}

// GetImageSizeOptimizationOptions : The GetImageSizeOptimization options.
type GetImageSizeOptimizationOptions struct {

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewGetImageSizeOptimizationOptions : Instantiate GetImageSizeOptimizationOptions
func (*ZonesSettingsV1) NewGetImageSizeOptimizationOptions() *GetImageSizeOptimizationOptions {
	return &GetImageSizeOptimizationOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *GetImageSizeOptimizationOptions) SetHeaders(param map[string]string) *GetImageSizeOptimizationOptions {
	options.Headers = param
	return options
}

// GetIpGeolocationOptions : The GetIpGeolocation options.
type GetIpGeolocationOptions struct {

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewGetIpGeolocationOptions : Instantiate GetIpGeolocationOptions
func (*ZonesSettingsV1) NewGetIpGeolocationOptions() *GetIpGeolocationOptions {
	return &GetIpGeolocationOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *GetIpGeolocationOptions) SetHeaders(param map[string]string) *GetIpGeolocationOptions {
	options.Headers = param
	return options
}

// GetIpv6Options : The GetIpv6 options.
type GetIpv6Options struct {

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewGetIpv6Options : Instantiate GetIpv6Options
func (*ZonesSettingsV1) NewGetIpv6Options() *GetIpv6Options {
	return &GetIpv6Options{}
}

// SetHeaders : Allow user to set Headers
func (options *GetIpv6Options) SetHeaders(param map[string]string) *GetIpv6Options {
	options.Headers = param
	return options
}

// GetLogRetentionOptions : The GetLogRetention options.
type GetLogRetentionOptions struct {
	// Full url-encoded cloud resource name (CRN) of resource instance.
	Crn *string `json:"crn" validate:"required,ne="`

	// Zone identifier.
	ZoneIdentifier *string `json:"zone_identifier" validate:"required,ne="`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewGetLogRetentionOptions : Instantiate GetLogRetentionOptions
func (*ZonesSettingsV1) NewGetLogRetentionOptions(crn string, zoneIdentifier string) *GetLogRetentionOptions {
	return &GetLogRetentionOptions{
		Crn:            core.StringPtr(crn),
		ZoneIdentifier: core.StringPtr(zoneIdentifier),
	}
}

// SetCrn : Allow user to set Crn
func (_options *GetLogRetentionOptions) SetCrn(crn string) *GetLogRetentionOptions {
	_options.Crn = core.StringPtr(crn)
	return _options
}

// SetZoneIdentifier : Allow user to set ZoneIdentifier
func (_options *GetLogRetentionOptions) SetZoneIdentifier(zoneIdentifier string) *GetLogRetentionOptions {
	_options.ZoneIdentifier = core.StringPtr(zoneIdentifier)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetLogRetentionOptions) SetHeaders(param map[string]string) *GetLogRetentionOptions {
	options.Headers = param
	return options
}

// GetMaxUploadOptions : The GetMaxUpload options.
type GetMaxUploadOptions struct {

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewGetMaxUploadOptions : Instantiate GetMaxUploadOptions
func (*ZonesSettingsV1) NewGetMaxUploadOptions() *GetMaxUploadOptions {
	return &GetMaxUploadOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *GetMaxUploadOptions) SetHeaders(param map[string]string) *GetMaxUploadOptions {
	options.Headers = param
	return options
}

// GetMinTlsVersionOptions : The GetMinTlsVersion options.
type GetMinTlsVersionOptions struct {

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewGetMinTlsVersionOptions : Instantiate GetMinTlsVersionOptions
func (*ZonesSettingsV1) NewGetMinTlsVersionOptions() *GetMinTlsVersionOptions {
	return &GetMinTlsVersionOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *GetMinTlsVersionOptions) SetHeaders(param map[string]string) *GetMinTlsVersionOptions {
	options.Headers = param
	return options
}

// GetMinifyOptions : The GetMinify options.
type GetMinifyOptions struct {

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewGetMinifyOptions : Instantiate GetMinifyOptions
func (*ZonesSettingsV1) NewGetMinifyOptions() *GetMinifyOptions {
	return &GetMinifyOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *GetMinifyOptions) SetHeaders(param map[string]string) *GetMinifyOptions {
	options.Headers = param
	return options
}

// GetMobileRedirectOptions : The GetMobileRedirect options.
type GetMobileRedirectOptions struct {

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewGetMobileRedirectOptions : Instantiate GetMobileRedirectOptions
func (*ZonesSettingsV1) NewGetMobileRedirectOptions() *GetMobileRedirectOptions {
	return &GetMobileRedirectOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *GetMobileRedirectOptions) SetHeaders(param map[string]string) *GetMobileRedirectOptions {
	options.Headers = param
	return options
}

// GetOpportunisticEncryptionOptions : The GetOpportunisticEncryption options.
type GetOpportunisticEncryptionOptions struct {

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewGetOpportunisticEncryptionOptions : Instantiate GetOpportunisticEncryptionOptions
func (*ZonesSettingsV1) NewGetOpportunisticEncryptionOptions() *GetOpportunisticEncryptionOptions {
	return &GetOpportunisticEncryptionOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *GetOpportunisticEncryptionOptions) SetHeaders(param map[string]string) *GetOpportunisticEncryptionOptions {
	options.Headers = param
	return options
}

// GetOpportunisticOnionOptions : The GetOpportunisticOnion options.
type GetOpportunisticOnionOptions struct {

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewGetOpportunisticOnionOptions : Instantiate GetOpportunisticOnionOptions
func (*ZonesSettingsV1) NewGetOpportunisticOnionOptions() *GetOpportunisticOnionOptions {
	return &GetOpportunisticOnionOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *GetOpportunisticOnionOptions) SetHeaders(param map[string]string) *GetOpportunisticOnionOptions {
	options.Headers = param
	return options
}

// GetOriginMaxHttpVersionOptions : The GetOriginMaxHttpVersion options.
type GetOriginMaxHttpVersionOptions struct {

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewGetOriginMaxHttpVersionOptions : Instantiate GetOriginMaxHttpVersionOptions
func (*ZonesSettingsV1) NewGetOriginMaxHttpVersionOptions() *GetOriginMaxHttpVersionOptions {
	return &GetOriginMaxHttpVersionOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *GetOriginMaxHttpVersionOptions) SetHeaders(param map[string]string) *GetOriginMaxHttpVersionOptions {
	options.Headers = param
	return options
}

// GetOriginPostQuantumEncryptionOptions : The GetOriginPostQuantumEncryption options.
type GetOriginPostQuantumEncryptionOptions struct {

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewGetOriginPostQuantumEncryptionOptions : Instantiate GetOriginPostQuantumEncryptionOptions
func (*ZonesSettingsV1) NewGetOriginPostQuantumEncryptionOptions() *GetOriginPostQuantumEncryptionOptions {
	return &GetOriginPostQuantumEncryptionOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *GetOriginPostQuantumEncryptionOptions) SetHeaders(param map[string]string) *GetOriginPostQuantumEncryptionOptions {
	options.Headers = param
	return options
}

// GetPrefetchPreloadOptions : The GetPrefetchPreload options.
type GetPrefetchPreloadOptions struct {

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewGetPrefetchPreloadOptions : Instantiate GetPrefetchPreloadOptions
func (*ZonesSettingsV1) NewGetPrefetchPreloadOptions() *GetPrefetchPreloadOptions {
	return &GetPrefetchPreloadOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *GetPrefetchPreloadOptions) SetHeaders(param map[string]string) *GetPrefetchPreloadOptions {
	options.Headers = param
	return options
}

// GetProxyReadTimeoutOptions : The GetProxyReadTimeout options.
type GetProxyReadTimeoutOptions struct {

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewGetProxyReadTimeoutOptions : Instantiate GetProxyReadTimeoutOptions
func (*ZonesSettingsV1) NewGetProxyReadTimeoutOptions() *GetProxyReadTimeoutOptions {
	return &GetProxyReadTimeoutOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *GetProxyReadTimeoutOptions) SetHeaders(param map[string]string) *GetProxyReadTimeoutOptions {
	options.Headers = param
	return options
}

// GetPseudoIpv4Options : The GetPseudoIpv4 options.
type GetPseudoIpv4Options struct {

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewGetPseudoIpv4Options : Instantiate GetPseudoIpv4Options
func (*ZonesSettingsV1) NewGetPseudoIpv4Options() *GetPseudoIpv4Options {
	return &GetPseudoIpv4Options{}
}

// SetHeaders : Allow user to set Headers
func (options *GetPseudoIpv4Options) SetHeaders(param map[string]string) *GetPseudoIpv4Options {
	options.Headers = param
	return options
}

// GetResponseBufferingOptions : The GetResponseBuffering options.
type GetResponseBufferingOptions struct {

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewGetResponseBufferingOptions : Instantiate GetResponseBufferingOptions
func (*ZonesSettingsV1) NewGetResponseBufferingOptions() *GetResponseBufferingOptions {
	return &GetResponseBufferingOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *GetResponseBufferingOptions) SetHeaders(param map[string]string) *GetResponseBufferingOptions {
	options.Headers = param
	return options
}

// GetScriptLoadOptimizationOptions : The GetScriptLoadOptimization options.
type GetScriptLoadOptimizationOptions struct {

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewGetScriptLoadOptimizationOptions : Instantiate GetScriptLoadOptimizationOptions
func (*ZonesSettingsV1) NewGetScriptLoadOptimizationOptions() *GetScriptLoadOptimizationOptions {
	return &GetScriptLoadOptimizationOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *GetScriptLoadOptimizationOptions) SetHeaders(param map[string]string) *GetScriptLoadOptimizationOptions {
	options.Headers = param
	return options
}

// GetSecurityHeaderOptions : The GetSecurityHeader options.
type GetSecurityHeaderOptions struct {

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewGetSecurityHeaderOptions : Instantiate GetSecurityHeaderOptions
func (*ZonesSettingsV1) NewGetSecurityHeaderOptions() *GetSecurityHeaderOptions {
	return &GetSecurityHeaderOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *GetSecurityHeaderOptions) SetHeaders(param map[string]string) *GetSecurityHeaderOptions {
	options.Headers = param
	return options
}

// GetServerSideExcludeOptions : The GetServerSideExclude options.
type GetServerSideExcludeOptions struct {

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewGetServerSideExcludeOptions : Instantiate GetServerSideExcludeOptions
func (*ZonesSettingsV1) NewGetServerSideExcludeOptions() *GetServerSideExcludeOptions {
	return &GetServerSideExcludeOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *GetServerSideExcludeOptions) SetHeaders(param map[string]string) *GetServerSideExcludeOptions {
	options.Headers = param
	return options
}

// GetTlsClientAuthOptions : The GetTlsClientAuth options.
type GetTlsClientAuthOptions struct {

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewGetTlsClientAuthOptions : Instantiate GetTlsClientAuthOptions
func (*ZonesSettingsV1) NewGetTlsClientAuthOptions() *GetTlsClientAuthOptions {
	return &GetTlsClientAuthOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *GetTlsClientAuthOptions) SetHeaders(param map[string]string) *GetTlsClientAuthOptions {
	options.Headers = param
	return options
}

// GetTrueClientIpOptions : The GetTrueClientIp options.
type GetTrueClientIpOptions struct {

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewGetTrueClientIpOptions : Instantiate GetTrueClientIpOptions
func (*ZonesSettingsV1) NewGetTrueClientIpOptions() *GetTrueClientIpOptions {
	return &GetTrueClientIpOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *GetTrueClientIpOptions) SetHeaders(param map[string]string) *GetTrueClientIpOptions {
	options.Headers = param
	return options
}

// GetWebApplicationFirewallOptions : The GetWebApplicationFirewall options.
type GetWebApplicationFirewallOptions struct {

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewGetWebApplicationFirewallOptions : Instantiate GetWebApplicationFirewallOptions
func (*ZonesSettingsV1) NewGetWebApplicationFirewallOptions() *GetWebApplicationFirewallOptions {
	return &GetWebApplicationFirewallOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *GetWebApplicationFirewallOptions) SetHeaders(param map[string]string) *GetWebApplicationFirewallOptions {
	options.Headers = param
	return options
}

// GetWebSocketsOptions : The GetWebSockets options.
type GetWebSocketsOptions struct {

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewGetWebSocketsOptions : Instantiate GetWebSocketsOptions
func (*ZonesSettingsV1) NewGetWebSocketsOptions() *GetWebSocketsOptions {
	return &GetWebSocketsOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *GetWebSocketsOptions) SetHeaders(param map[string]string) *GetWebSocketsOptions {
	options.Headers = param
	return options
}

// GetZoneCnameFlatteningOptions : The GetZoneCnameFlattening options.
type GetZoneCnameFlatteningOptions struct {

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewGetZoneCnameFlatteningOptions : Instantiate GetZoneCnameFlatteningOptions
func (*ZonesSettingsV1) NewGetZoneCnameFlatteningOptions() *GetZoneCnameFlatteningOptions {
	return &GetZoneCnameFlatteningOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *GetZoneCnameFlatteningOptions) SetHeaders(param map[string]string) *GetZoneCnameFlatteningOptions {
	options.Headers = param
	return options
}

// GetZoneDnssecOptions : The GetZoneDnssec options.
type GetZoneDnssecOptions struct {

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewGetZoneDnssecOptions : Instantiate GetZoneDnssecOptions
func (*ZonesSettingsV1) NewGetZoneDnssecOptions() *GetZoneDnssecOptions {
	return &GetZoneDnssecOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *GetZoneDnssecOptions) SetHeaders(param map[string]string) *GetZoneDnssecOptions {
	options.Headers = param
	return options
}

// HotlinkProtectionRespResult : Container for response information.
type HotlinkProtectionRespResult struct {
	// ID.
	ID *string `json:"id" validate:"required"`

	// Value.
	Value *string `json:"value" validate:"required"`

	// Editable.
	Editable *bool `json:"editable" validate:"required"`

	// Modified date.
	ModifiedOn *strfmt.DateTime `json:"modified_on" validate:"required"`
}

// UnmarshalHotlinkProtectionRespResult unmarshals an instance of HotlinkProtectionRespResult from the specified map of raw messages.
func UnmarshalHotlinkProtectionRespResult(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(HotlinkProtectionRespResult)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		err = core.SDKErrorf(err, "", "id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "value", &obj.Value)
	if err != nil {
		err = core.SDKErrorf(err, "", "value-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "editable", &obj.Editable)
	if err != nil {
		err = core.SDKErrorf(err, "", "editable-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "modified_on", &obj.ModifiedOn)
	if err != nil {
		err = core.SDKErrorf(err, "", "modified_on-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Http2RespResult : Container for response information.
type Http2RespResult struct {
	// ID.
	ID *string `json:"id" validate:"required"`

	// Value.
	Value *string `json:"value" validate:"required"`

	// Editable.
	Editable *bool `json:"editable" validate:"required"`

	// Modified date.
	ModifiedOn *strfmt.DateTime `json:"modified_on" validate:"required"`
}

// UnmarshalHttp2RespResult unmarshals an instance of Http2RespResult from the specified map of raw messages.
func UnmarshalHttp2RespResult(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Http2RespResult)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		err = core.SDKErrorf(err, "", "id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "value", &obj.Value)
	if err != nil {
		err = core.SDKErrorf(err, "", "value-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "editable", &obj.Editable)
	if err != nil {
		err = core.SDKErrorf(err, "", "editable-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "modified_on", &obj.ModifiedOn)
	if err != nil {
		err = core.SDKErrorf(err, "", "modified_on-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Http3RespResult : Container for response information.
type Http3RespResult struct {
	// ID.
	ID *string `json:"id" validate:"required"`

	// Value.
	Value *string `json:"value" validate:"required"`

	// Editable.
	Editable *bool `json:"editable" validate:"required"`

	// Modified date.
	ModifiedOn *strfmt.DateTime `json:"modified_on" validate:"required"`
}

// UnmarshalHttp3RespResult unmarshals an instance of Http3RespResult from the specified map of raw messages.
func UnmarshalHttp3RespResult(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Http3RespResult)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		err = core.SDKErrorf(err, "", "id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "value", &obj.Value)
	if err != nil {
		err = core.SDKErrorf(err, "", "value-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "editable", &obj.Editable)
	if err != nil {
		err = core.SDKErrorf(err, "", "editable-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "modified_on", &obj.ModifiedOn)
	if err != nil {
		err = core.SDKErrorf(err, "", "modified_on-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ImageLoadOptimizationRespResult : Container for response information.
type ImageLoadOptimizationRespResult struct {
	// ID.
	ID *string `json:"id" validate:"required"`

	// Value.
	Value *string `json:"value" validate:"required"`

	// Editable.
	Editable *bool `json:"editable" validate:"required"`

	// Modified date.
	ModifiedOn *strfmt.DateTime `json:"modified_on" validate:"required"`
}

// UnmarshalImageLoadOptimizationRespResult unmarshals an instance of ImageLoadOptimizationRespResult from the specified map of raw messages.
func UnmarshalImageLoadOptimizationRespResult(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ImageLoadOptimizationRespResult)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		err = core.SDKErrorf(err, "", "id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "value", &obj.Value)
	if err != nil {
		err = core.SDKErrorf(err, "", "value-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "editable", &obj.Editable)
	if err != nil {
		err = core.SDKErrorf(err, "", "editable-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "modified_on", &obj.ModifiedOn)
	if err != nil {
		err = core.SDKErrorf(err, "", "modified_on-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ImageSizeOptimizationRespResult : Container for response information.
type ImageSizeOptimizationRespResult struct {
	// ID.
	ID *string `json:"id" validate:"required"`

	// Value.
	Value *string `json:"value" validate:"required"`

	// Editable.
	Editable *bool `json:"editable" validate:"required"`

	// Modified date.
	ModifiedOn *strfmt.DateTime `json:"modified_on" validate:"required"`
}

// UnmarshalImageSizeOptimizationRespResult unmarshals an instance of ImageSizeOptimizationRespResult from the specified map of raw messages.
func UnmarshalImageSizeOptimizationRespResult(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ImageSizeOptimizationRespResult)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		err = core.SDKErrorf(err, "", "id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "value", &obj.Value)
	if err != nil {
		err = core.SDKErrorf(err, "", "value-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "editable", &obj.Editable)
	if err != nil {
		err = core.SDKErrorf(err, "", "editable-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "modified_on", &obj.ModifiedOn)
	if err != nil {
		err = core.SDKErrorf(err, "", "modified_on-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// IpGeolocationRespResult : Container for response information.
type IpGeolocationRespResult struct {
	// ID.
	ID *string `json:"id" validate:"required"`

	// Value.
	Value *string `json:"value" validate:"required"`

	// Editable.
	Editable *bool `json:"editable" validate:"required"`

	// Modified date.
	ModifiedOn *strfmt.DateTime `json:"modified_on" validate:"required"`
}

// UnmarshalIpGeolocationRespResult unmarshals an instance of IpGeolocationRespResult from the specified map of raw messages.
func UnmarshalIpGeolocationRespResult(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(IpGeolocationRespResult)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		err = core.SDKErrorf(err, "", "id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "value", &obj.Value)
	if err != nil {
		err = core.SDKErrorf(err, "", "value-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "editable", &obj.Editable)
	if err != nil {
		err = core.SDKErrorf(err, "", "editable-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "modified_on", &obj.ModifiedOn)
	if err != nil {
		err = core.SDKErrorf(err, "", "modified_on-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Ipv6RespResult : Container for response information.
type Ipv6RespResult struct {
	// ID.
	ID *string `json:"id" validate:"required"`

	// Value.
	Value *string `json:"value" validate:"required"`

	// Editable.
	Editable *bool `json:"editable" validate:"required"`

	// Modified date.
	ModifiedOn *strfmt.DateTime `json:"modified_on" validate:"required"`
}

// UnmarshalIpv6RespResult unmarshals an instance of Ipv6RespResult from the specified map of raw messages.
func UnmarshalIpv6RespResult(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Ipv6RespResult)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		err = core.SDKErrorf(err, "", "id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "value", &obj.Value)
	if err != nil {
		err = core.SDKErrorf(err, "", "value-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "editable", &obj.Editable)
	if err != nil {
		err = core.SDKErrorf(err, "", "editable-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "modified_on", &obj.ModifiedOn)
	if err != nil {
		err = core.SDKErrorf(err, "", "modified_on-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// LogRetentionRespResult : LogRetentionRespResult struct
type LogRetentionRespResult struct {
	// Boolean flag indicating whether or not log retention is turned on or off.
	Flag *bool `json:"flag" validate:"required"`
}

// UnmarshalLogRetentionRespResult unmarshals an instance of LogRetentionRespResult from the specified map of raw messages.
func UnmarshalLogRetentionRespResult(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(LogRetentionRespResult)
	err = core.UnmarshalPrimitive(m, "flag", &obj.Flag)
	if err != nil {
		err = core.SDKErrorf(err, "", "flag-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// MaxUploadRespResult : Container for response information.
type MaxUploadRespResult struct {
	// ID.
	ID *string `json:"id" validate:"required"`

	// Value.
	Value *int64 `json:"value" validate:"required"`

	// Editable.
	Editable *bool `json:"editable" validate:"required"`

	// Modified date.
	ModifiedOn *strfmt.DateTime `json:"modified_on" validate:"required"`
}

// UnmarshalMaxUploadRespResult unmarshals an instance of MaxUploadRespResult from the specified map of raw messages.
func UnmarshalMaxUploadRespResult(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(MaxUploadRespResult)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		err = core.SDKErrorf(err, "", "id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "value", &obj.Value)
	if err != nil {
		err = core.SDKErrorf(err, "", "value-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "editable", &obj.Editable)
	if err != nil {
		err = core.SDKErrorf(err, "", "editable-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "modified_on", &obj.ModifiedOn)
	if err != nil {
		err = core.SDKErrorf(err, "", "modified_on-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// MinTlsVersionRespResult : Container for response information.
type MinTlsVersionRespResult struct {
	// ID.
	ID *string `json:"id" validate:"required"`

	// Value.
	Value *string `json:"value" validate:"required"`

	// Editable.
	Editable *bool `json:"editable" validate:"required"`

	// Modified date.
	ModifiedOn *strfmt.DateTime `json:"modified_on" validate:"required"`
}

// UnmarshalMinTlsVersionRespResult unmarshals an instance of MinTlsVersionRespResult from the specified map of raw messages.
func UnmarshalMinTlsVersionRespResult(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(MinTlsVersionRespResult)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		err = core.SDKErrorf(err, "", "id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "value", &obj.Value)
	if err != nil {
		err = core.SDKErrorf(err, "", "value-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "editable", &obj.Editable)
	if err != nil {
		err = core.SDKErrorf(err, "", "editable-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "modified_on", &obj.ModifiedOn)
	if err != nil {
		err = core.SDKErrorf(err, "", "modified_on-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// MinifyRespResult : Container for response information.
type MinifyRespResult struct {
	// ID.
	ID *string `json:"id" validate:"required"`

	// Value.
	Value *MinifyRespResultValue `json:"value" validate:"required"`

	// Editable.
	Editable *bool `json:"editable" validate:"required"`

	// Modified date.
	ModifiedOn *strfmt.DateTime `json:"modified_on" validate:"required"`
}

// UnmarshalMinifyRespResult unmarshals an instance of MinifyRespResult from the specified map of raw messages.
func UnmarshalMinifyRespResult(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(MinifyRespResult)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		err = core.SDKErrorf(err, "", "id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "value", &obj.Value, UnmarshalMinifyRespResultValue)
	if err != nil {
		err = core.SDKErrorf(err, "", "value-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "editable", &obj.Editable)
	if err != nil {
		err = core.SDKErrorf(err, "", "editable-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "modified_on", &obj.ModifiedOn)
	if err != nil {
		err = core.SDKErrorf(err, "", "modified_on-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// MinifyRespResultValue : Value.
type MinifyRespResultValue struct {
	// css.
	Css *string `json:"css" validate:"required"`

	// html.
	HTML *string `json:"html" validate:"required"`

	// js.
	Js *string `json:"js" validate:"required"`
}

// UnmarshalMinifyRespResultValue unmarshals an instance of MinifyRespResultValue from the specified map of raw messages.
func UnmarshalMinifyRespResultValue(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(MinifyRespResultValue)
	err = core.UnmarshalPrimitive(m, "css", &obj.Css)
	if err != nil {
		err = core.SDKErrorf(err, "", "css-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "html", &obj.HTML)
	if err != nil {
		err = core.SDKErrorf(err, "", "html-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "js", &obj.Js)
	if err != nil {
		err = core.SDKErrorf(err, "", "js-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// MinifySettingValue : Value.
type MinifySettingValue struct {
	// Automatically minify all CSS for your website.
	Css *string `json:"css" validate:"required"`

	// Automatically minify all HTML for your website.
	HTML *string `json:"html" validate:"required"`

	// Automatically minify all JavaScript for your website.
	Js *string `json:"js" validate:"required"`
}

// Constants associated with the MinifySettingValue.Css property.
// Automatically minify all CSS for your website.
const (
	MinifySettingValue_Css_Off = "off"
	MinifySettingValue_Css_On  = "on"
)

// Constants associated with the MinifySettingValue.HTML property.
// Automatically minify all HTML for your website.
const (
	MinifySettingValue_HTML_Off = "off"
	MinifySettingValue_HTML_On  = "on"
)

// Constants associated with the MinifySettingValue.Js property.
// Automatically minify all JavaScript for your website.
const (
	MinifySettingValue_Js_Off = "off"
	MinifySettingValue_Js_On  = "on"
)

// NewMinifySettingValue : Instantiate MinifySettingValue (Generic Model Constructor)
func (*ZonesSettingsV1) NewMinifySettingValue(css string, html string, js string) (_model *MinifySettingValue, err error) {
	_model = &MinifySettingValue{
		Css:  core.StringPtr(css),
		HTML: core.StringPtr(html),
		Js:   core.StringPtr(js),
	}
	err = core.ValidateStruct(_model, "required parameters")
	if err != nil {
		err = core.SDKErrorf(err, "", "model-missing-required", common.GetComponentInfo())
	}
	return
}

// UnmarshalMinifySettingValue unmarshals an instance of MinifySettingValue from the specified map of raw messages.
func UnmarshalMinifySettingValue(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(MinifySettingValue)
	err = core.UnmarshalPrimitive(m, "css", &obj.Css)
	if err != nil {
		err = core.SDKErrorf(err, "", "css-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "html", &obj.HTML)
	if err != nil {
		err = core.SDKErrorf(err, "", "html-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "js", &obj.Js)
	if err != nil {
		err = core.SDKErrorf(err, "", "js-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// MobileRedirecSettingValue : Value.
type MobileRedirecSettingValue struct {
	// Whether or not the mobile redirection is enabled.
	Status *string `json:"status" validate:"required"`

	// Which subdomain prefix you wish to redirect visitors on mobile devices to.
	MobileSubdomain *string `json:"mobile_subdomain" validate:"required"`

	// Whether to drop the current page path and redirect to the mobile subdomain URL root or to keep the path and redirect
	// to the same page on the mobile subdomain.
	StripURI *bool `json:"strip_uri" validate:"required"`
}

// Constants associated with the MobileRedirecSettingValue.Status property.
// Whether or not the mobile redirection is enabled.
const (
	MobileRedirecSettingValue_Status_Off = "off"
	MobileRedirecSettingValue_Status_On  = "on"
)

// NewMobileRedirecSettingValue : Instantiate MobileRedirecSettingValue (Generic Model Constructor)
func (*ZonesSettingsV1) NewMobileRedirecSettingValue(status string, mobileSubdomain string, stripURI bool) (_model *MobileRedirecSettingValue, err error) {
	_model = &MobileRedirecSettingValue{
		Status:          core.StringPtr(status),
		MobileSubdomain: core.StringPtr(mobileSubdomain),
		StripURI:        core.BoolPtr(stripURI),
	}
	err = core.ValidateStruct(_model, "required parameters")
	if err != nil {
		err = core.SDKErrorf(err, "", "model-missing-required", common.GetComponentInfo())
	}
	return
}

// UnmarshalMobileRedirecSettingValue unmarshals an instance of MobileRedirecSettingValue from the specified map of raw messages.
func UnmarshalMobileRedirecSettingValue(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(MobileRedirecSettingValue)
	err = core.UnmarshalPrimitive(m, "status", &obj.Status)
	if err != nil {
		err = core.SDKErrorf(err, "", "status-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "mobile_subdomain", &obj.MobileSubdomain)
	if err != nil {
		err = core.SDKErrorf(err, "", "mobile_subdomain-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "strip_uri", &obj.StripURI)
	if err != nil {
		err = core.SDKErrorf(err, "", "strip_uri-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// MobileRedirectRespResult : Container for response information.
type MobileRedirectRespResult struct {
	// ID.
	ID *string `json:"id" validate:"required"`

	// Value.
	Value *MobileRedirectRespResultValue `json:"value" validate:"required"`

	// Editable.
	Editable *bool `json:"editable" validate:"required"`

	// Modified date.
	ModifiedOn *strfmt.DateTime `json:"modified_on" validate:"required"`
}

// UnmarshalMobileRedirectRespResult unmarshals an instance of MobileRedirectRespResult from the specified map of raw messages.
func UnmarshalMobileRedirectRespResult(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(MobileRedirectRespResult)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		err = core.SDKErrorf(err, "", "id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "value", &obj.Value, UnmarshalMobileRedirectRespResultValue)
	if err != nil {
		err = core.SDKErrorf(err, "", "value-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "editable", &obj.Editable)
	if err != nil {
		err = core.SDKErrorf(err, "", "editable-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "modified_on", &obj.ModifiedOn)
	if err != nil {
		err = core.SDKErrorf(err, "", "modified_on-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// MobileRedirectRespResultValue : Value.
type MobileRedirectRespResultValue struct {
	// Whether or not the mobile redirection is enabled.
	Status *string `json:"status" validate:"required"`

	// Which subdomain prefix you wish to redirect visitors on mobile devices to.
	MobileSubdomain *string `json:"mobile_subdomain" validate:"required"`

	// Whether to drop the current page path and redirect to the mobile subdomain URL root or to keep the path and redirect
	// to the same page on the mobile subdomain.
	StripURI *bool `json:"strip_uri" validate:"required"`
}

// UnmarshalMobileRedirectRespResultValue unmarshals an instance of MobileRedirectRespResultValue from the specified map of raw messages.
func UnmarshalMobileRedirectRespResultValue(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(MobileRedirectRespResultValue)
	err = core.UnmarshalPrimitive(m, "status", &obj.Status)
	if err != nil {
		err = core.SDKErrorf(err, "", "status-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "mobile_subdomain", &obj.MobileSubdomain)
	if err != nil {
		err = core.SDKErrorf(err, "", "mobile_subdomain-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "strip_uri", &obj.StripURI)
	if err != nil {
		err = core.SDKErrorf(err, "", "strip_uri-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// OpportunisticEncryptionRespResult : Container for response information.
type OpportunisticEncryptionRespResult struct {
	// ID.
	ID *string `json:"id" validate:"required"`

	// Value.
	Value *string `json:"value" validate:"required"`

	// Editable.
	Editable *bool `json:"editable" validate:"required"`

	// Modified date.
	ModifiedOn *strfmt.DateTime `json:"modified_on" validate:"required"`
}

// UnmarshalOpportunisticEncryptionRespResult unmarshals an instance of OpportunisticEncryptionRespResult from the specified map of raw messages.
func UnmarshalOpportunisticEncryptionRespResult(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(OpportunisticEncryptionRespResult)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		err = core.SDKErrorf(err, "", "id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "value", &obj.Value)
	if err != nil {
		err = core.SDKErrorf(err, "", "value-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "editable", &obj.Editable)
	if err != nil {
		err = core.SDKErrorf(err, "", "editable-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "modified_on", &obj.ModifiedOn)
	if err != nil {
		err = core.SDKErrorf(err, "", "modified_on-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// OpportunisticOnionRespResult : Container for response information.
type OpportunisticOnionRespResult struct {
	// ID.
	ID *string `json:"id" validate:"required"`

	// Value.
	Value *string `json:"value" validate:"required"`

	// Editable.
	Editable *bool `json:"editable" validate:"required"`

	// Modified date.
	ModifiedOn *strfmt.DateTime `json:"modified_on" validate:"required"`
}

// UnmarshalOpportunisticOnionRespResult unmarshals an instance of OpportunisticOnionRespResult from the specified map of raw messages.
func UnmarshalOpportunisticOnionRespResult(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(OpportunisticOnionRespResult)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		err = core.SDKErrorf(err, "", "id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "value", &obj.Value)
	if err != nil {
		err = core.SDKErrorf(err, "", "value-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "editable", &obj.Editable)
	if err != nil {
		err = core.SDKErrorf(err, "", "editable-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "modified_on", &obj.ModifiedOn)
	if err != nil {
		err = core.SDKErrorf(err, "", "modified_on-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// OriginErrorPagePassThruRespResult : Container for response information.
type OriginErrorPagePassThruRespResult struct {
	// ID.
	ID *string `json:"id" validate:"required"`

	// Value.
	Value *string `json:"value" validate:"required"`

	// Editable.
	Editable *bool `json:"editable" validate:"required"`

	// Modified date.
	ModifiedOn *strfmt.DateTime `json:"modified_on" validate:"required"`
}

// UnmarshalOriginErrorPagePassThruRespResult unmarshals an instance of OriginErrorPagePassThruRespResult from the specified map of raw messages.
func UnmarshalOriginErrorPagePassThruRespResult(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(OriginErrorPagePassThruRespResult)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		err = core.SDKErrorf(err, "", "id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "value", &obj.Value)
	if err != nil {
		err = core.SDKErrorf(err, "", "value-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "editable", &obj.Editable)
	if err != nil {
		err = core.SDKErrorf(err, "", "editable-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "modified_on", &obj.ModifiedOn)
	if err != nil {
		err = core.SDKErrorf(err, "", "modified_on-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// OriginMaxHttpVersionRespResult : Container for response information.
type OriginMaxHttpVersionRespResult struct {
	// ID.
	ID *string `json:"id" validate:"required"`

	// Value.
	Value *string `json:"value" validate:"required"`

	// Editable.
	Editable *bool `json:"editable" validate:"required"`

	// Modified date.
	ModifiedOn *strfmt.DateTime `json:"modified_on" validate:"required"`
}

// UnmarshalOriginMaxHttpVersionRespResult unmarshals an instance of OriginMaxHttpVersionRespResult from the specified map of raw messages.
func UnmarshalOriginMaxHttpVersionRespResult(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(OriginMaxHttpVersionRespResult)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		err = core.SDKErrorf(err, "", "id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "value", &obj.Value)
	if err != nil {
		err = core.SDKErrorf(err, "", "value-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "editable", &obj.Editable)
	if err != nil {
		err = core.SDKErrorf(err, "", "editable-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "modified_on", &obj.ModifiedOn)
	if err != nil {
		err = core.SDKErrorf(err, "", "modified_on-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// OriginPostQuantumEncryptionRespResult : Container for response information.
type OriginPostQuantumEncryptionRespResult struct {
	// ID.
	ID *string `json:"id" validate:"required"`

	// Value.
	Value *string `json:"value" validate:"required"`

	// Editable.
	Editable *bool `json:"editable" validate:"required"`

	// Modified date.
	ModifiedOn *strfmt.DateTime `json:"modified_on" validate:"required"`
}

// Constants associated with the OriginPostQuantumEncryptionRespResult.Value property.
// Value.
const (
	OriginPostQuantumEncryptionRespResult_Value_Off       = "off"
	OriginPostQuantumEncryptionRespResult_Value_Preferred = "preferred"
	OriginPostQuantumEncryptionRespResult_Value_Supported = "supported"
)

// UnmarshalOriginPostQuantumEncryptionRespResult unmarshals an instance of OriginPostQuantumEncryptionRespResult from the specified map of raw messages.
func UnmarshalOriginPostQuantumEncryptionRespResult(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(OriginPostQuantumEncryptionRespResult)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		err = core.SDKErrorf(err, "", "id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "value", &obj.Value)
	if err != nil {
		err = core.SDKErrorf(err, "", "value-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "editable", &obj.Editable)
	if err != nil {
		err = core.SDKErrorf(err, "", "editable-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "modified_on", &obj.ModifiedOn)
	if err != nil {
		err = core.SDKErrorf(err, "", "modified_on-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// PrefetchPreloadRespResult : Container for response information.
type PrefetchPreloadRespResult struct {
	// ID.
	ID *string `json:"id" validate:"required"`

	// Value.
	Value *string `json:"value" validate:"required"`

	// Editable.
	Editable *bool `json:"editable" validate:"required"`

	// Modified date.
	ModifiedOn *strfmt.DateTime `json:"modified_on" validate:"required"`
}

// UnmarshalPrefetchPreloadRespResult unmarshals an instance of PrefetchPreloadRespResult from the specified map of raw messages.
func UnmarshalPrefetchPreloadRespResult(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PrefetchPreloadRespResult)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		err = core.SDKErrorf(err, "", "id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "value", &obj.Value)
	if err != nil {
		err = core.SDKErrorf(err, "", "value-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "editable", &obj.Editable)
	if err != nil {
		err = core.SDKErrorf(err, "", "editable-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "modified_on", &obj.ModifiedOn)
	if err != nil {
		err = core.SDKErrorf(err, "", "modified_on-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ProxyReadTimeoutRespResult : Container for response information.
type ProxyReadTimeoutRespResult struct {
	// ID.
	ID *string `json:"id" validate:"required"`

	// Value.
	Value *float64 `json:"value" validate:"required"`

	// Editable.
	Editable *bool `json:"editable" validate:"required"`

	// Modified date.
	ModifiedOn *strfmt.DateTime `json:"modified_on" validate:"required"`
}

// UnmarshalProxyReadTimeoutRespResult unmarshals an instance of ProxyReadTimeoutRespResult from the specified map of raw messages.
func UnmarshalProxyReadTimeoutRespResult(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ProxyReadTimeoutRespResult)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		err = core.SDKErrorf(err, "", "id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "value", &obj.Value)
	if err != nil {
		err = core.SDKErrorf(err, "", "value-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "editable", &obj.Editable)
	if err != nil {
		err = core.SDKErrorf(err, "", "editable-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "modified_on", &obj.ModifiedOn)
	if err != nil {
		err = core.SDKErrorf(err, "", "modified_on-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// PseudoIpv4RespResult : Container for response information.
type PseudoIpv4RespResult struct {
	// ID.
	ID *string `json:"id" validate:"required"`

	// Value.
	Value *string `json:"value" validate:"required"`

	// Editable.
	Editable *bool `json:"editable" validate:"required"`

	// Modified date.
	ModifiedOn *strfmt.DateTime `json:"modified_on" validate:"required"`
}

// UnmarshalPseudoIpv4RespResult unmarshals an instance of PseudoIpv4RespResult from the specified map of raw messages.
func UnmarshalPseudoIpv4RespResult(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PseudoIpv4RespResult)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		err = core.SDKErrorf(err, "", "id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "value", &obj.Value)
	if err != nil {
		err = core.SDKErrorf(err, "", "value-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "editable", &obj.Editable)
	if err != nil {
		err = core.SDKErrorf(err, "", "editable-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "modified_on", &obj.ModifiedOn)
	if err != nil {
		err = core.SDKErrorf(err, "", "modified_on-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ResponseBufferingRespResult : Container for response information.
type ResponseBufferingRespResult struct {
	// ID.
	ID *string `json:"id" validate:"required"`

	// Value.
	Value *string `json:"value" validate:"required"`

	// Editable.
	Editable *bool `json:"editable" validate:"required"`

	// Modified date.
	ModifiedOn *strfmt.DateTime `json:"modified_on" validate:"required"`
}

// UnmarshalResponseBufferingRespResult unmarshals an instance of ResponseBufferingRespResult from the specified map of raw messages.
func UnmarshalResponseBufferingRespResult(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ResponseBufferingRespResult)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		err = core.SDKErrorf(err, "", "id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "value", &obj.Value)
	if err != nil {
		err = core.SDKErrorf(err, "", "value-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "editable", &obj.Editable)
	if err != nil {
		err = core.SDKErrorf(err, "", "editable-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "modified_on", &obj.ModifiedOn)
	if err != nil {
		err = core.SDKErrorf(err, "", "modified_on-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ScriptLoadOptimizationRespResult : Container for response information.
type ScriptLoadOptimizationRespResult struct {
	// ID.
	ID *string `json:"id" validate:"required"`

	// Value.
	Value *string `json:"value" validate:"required"`

	// Editable.
	Editable *bool `json:"editable" validate:"required"`

	// Modified date.
	ModifiedOn *strfmt.DateTime `json:"modified_on" validate:"required"`
}

// UnmarshalScriptLoadOptimizationRespResult unmarshals an instance of ScriptLoadOptimizationRespResult from the specified map of raw messages.
func UnmarshalScriptLoadOptimizationRespResult(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ScriptLoadOptimizationRespResult)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		err = core.SDKErrorf(err, "", "id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "value", &obj.Value)
	if err != nil {
		err = core.SDKErrorf(err, "", "value-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "editable", &obj.Editable)
	if err != nil {
		err = core.SDKErrorf(err, "", "editable-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "modified_on", &obj.ModifiedOn)
	if err != nil {
		err = core.SDKErrorf(err, "", "modified_on-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SecurityHeaderRespResult : Container for response information.
type SecurityHeaderRespResult struct {
	// ID.
	ID *string `json:"id" validate:"required"`

	// Value.
	Value *SecurityHeaderRespResultValue `json:"value" validate:"required"`

	// Editable.
	Editable *bool `json:"editable" validate:"required"`

	// Modified date.
	ModifiedOn *strfmt.DateTime `json:"modified_on" validate:"required"`
}

// UnmarshalSecurityHeaderRespResult unmarshals an instance of SecurityHeaderRespResult from the specified map of raw messages.
func UnmarshalSecurityHeaderRespResult(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SecurityHeaderRespResult)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		err = core.SDKErrorf(err, "", "id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "value", &obj.Value, UnmarshalSecurityHeaderRespResultValue)
	if err != nil {
		err = core.SDKErrorf(err, "", "value-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "editable", &obj.Editable)
	if err != nil {
		err = core.SDKErrorf(err, "", "editable-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "modified_on", &obj.ModifiedOn)
	if err != nil {
		err = core.SDKErrorf(err, "", "modified_on-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SecurityHeaderRespResultValue : Value.
type SecurityHeaderRespResultValue struct {
	// Strict transport security.
	StrictTransportSecurity *SecurityHeaderRespResultValueStrictTransportSecurity `json:"strict_transport_security" validate:"required"`
}

// UnmarshalSecurityHeaderRespResultValue unmarshals an instance of SecurityHeaderRespResultValue from the specified map of raw messages.
func UnmarshalSecurityHeaderRespResultValue(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SecurityHeaderRespResultValue)
	err = core.UnmarshalModel(m, "strict_transport_security", &obj.StrictTransportSecurity, UnmarshalSecurityHeaderRespResultValueStrictTransportSecurity)
	if err != nil {
		err = core.SDKErrorf(err, "", "strict_transport_security-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SecurityHeaderRespResultValueStrictTransportSecurity : Strict transport security.
type SecurityHeaderRespResultValueStrictTransportSecurity struct {
	// Whether or not security header is enabled.
	Enabled *bool `json:"enabled" validate:"required"`

	// Max age in seconds.
	MaxAge *int64 `json:"max_age" validate:"required"`

	// Include all subdomains.
	IncludeSubdomains *bool `json:"include_subdomains" validate:"required"`

	// Whether or not to permit browsers to preload security_header config.
	Preload *bool `json:"preload" validate:"required"`

	// Whether or not to include 'X-Content-Type-Options:nosniff' header.
	Nosniff *bool `json:"nosniff" validate:"required"`
}

// UnmarshalSecurityHeaderRespResultValueStrictTransportSecurity unmarshals an instance of SecurityHeaderRespResultValueStrictTransportSecurity from the specified map of raw messages.
func UnmarshalSecurityHeaderRespResultValueStrictTransportSecurity(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SecurityHeaderRespResultValueStrictTransportSecurity)
	err = core.UnmarshalPrimitive(m, "enabled", &obj.Enabled)
	if err != nil {
		err = core.SDKErrorf(err, "", "enabled-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "max_age", &obj.MaxAge)
	if err != nil {
		err = core.SDKErrorf(err, "", "max_age-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "include_subdomains", &obj.IncludeSubdomains)
	if err != nil {
		err = core.SDKErrorf(err, "", "include_subdomains-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "preload", &obj.Preload)
	if err != nil {
		err = core.SDKErrorf(err, "", "preload-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "nosniff", &obj.Nosniff)
	if err != nil {
		err = core.SDKErrorf(err, "", "nosniff-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SecurityHeaderSettingValue : Value.
type SecurityHeaderSettingValue struct {
	// Strict transport security.
	StrictTransportSecurity *SecurityHeaderSettingValueStrictTransportSecurity `json:"strict_transport_security" validate:"required"`
}

// NewSecurityHeaderSettingValue : Instantiate SecurityHeaderSettingValue (Generic Model Constructor)
func (*ZonesSettingsV1) NewSecurityHeaderSettingValue(strictTransportSecurity *SecurityHeaderSettingValueStrictTransportSecurity) (_model *SecurityHeaderSettingValue, err error) {
	_model = &SecurityHeaderSettingValue{
		StrictTransportSecurity: strictTransportSecurity,
	}
	err = core.ValidateStruct(_model, "required parameters")
	if err != nil {
		err = core.SDKErrorf(err, "", "model-missing-required", common.GetComponentInfo())
	}
	return
}

// UnmarshalSecurityHeaderSettingValue unmarshals an instance of SecurityHeaderSettingValue from the specified map of raw messages.
func UnmarshalSecurityHeaderSettingValue(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SecurityHeaderSettingValue)
	err = core.UnmarshalModel(m, "strict_transport_security", &obj.StrictTransportSecurity, UnmarshalSecurityHeaderSettingValueStrictTransportSecurity)
	if err != nil {
		err = core.SDKErrorf(err, "", "strict_transport_security-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SecurityHeaderSettingValueStrictTransportSecurity : Strict transport security.
type SecurityHeaderSettingValueStrictTransportSecurity struct {
	// Whether or not security header is enabled.
	Enabled *bool `json:"enabled" validate:"required"`

	// Max age in seconds.
	MaxAge *int64 `json:"max_age" validate:"required"`

	// Include all subdomains.
	IncludeSubdomains *bool `json:"include_subdomains" validate:"required"`

	// Whether or not to permit browsers to preload security_header config.
	Preload *bool `json:"preload" validate:"required"`

	// Whether or not to include 'X-Content-Type-Options:nosniff' header.
	Nosniff *bool `json:"nosniff" validate:"required"`
}

// NewSecurityHeaderSettingValueStrictTransportSecurity : Instantiate SecurityHeaderSettingValueStrictTransportSecurity (Generic Model Constructor)
func (*ZonesSettingsV1) NewSecurityHeaderSettingValueStrictTransportSecurity(enabled bool, maxAge int64, includeSubdomains bool, preload bool, nosniff bool) (_model *SecurityHeaderSettingValueStrictTransportSecurity, err error) {
	_model = &SecurityHeaderSettingValueStrictTransportSecurity{
		Enabled:           core.BoolPtr(enabled),
		MaxAge:            core.Int64Ptr(maxAge),
		IncludeSubdomains: core.BoolPtr(includeSubdomains),
		Preload:           core.BoolPtr(preload),
		Nosniff:           core.BoolPtr(nosniff),
	}
	err = core.ValidateStruct(_model, "required parameters")
	if err != nil {
		err = core.SDKErrorf(err, "", "model-missing-required", common.GetComponentInfo())
	}
	return
}

// UnmarshalSecurityHeaderSettingValueStrictTransportSecurity unmarshals an instance of SecurityHeaderSettingValueStrictTransportSecurity from the specified map of raw messages.
func UnmarshalSecurityHeaderSettingValueStrictTransportSecurity(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SecurityHeaderSettingValueStrictTransportSecurity)
	err = core.UnmarshalPrimitive(m, "enabled", &obj.Enabled)
	if err != nil {
		err = core.SDKErrorf(err, "", "enabled-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "max_age", &obj.MaxAge)
	if err != nil {
		err = core.SDKErrorf(err, "", "max_age-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "include_subdomains", &obj.IncludeSubdomains)
	if err != nil {
		err = core.SDKErrorf(err, "", "include_subdomains-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "preload", &obj.Preload)
	if err != nil {
		err = core.SDKErrorf(err, "", "preload-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "nosniff", &obj.Nosniff)
	if err != nil {
		err = core.SDKErrorf(err, "", "nosniff-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ServerSideExcludeRespResult : Container for response information.
type ServerSideExcludeRespResult struct {
	// ID.
	ID *string `json:"id" validate:"required"`

	// Value.
	Value *string `json:"value" validate:"required"`

	// Editable.
	Editable *bool `json:"editable" validate:"required"`

	// Modified date.
	ModifiedOn *strfmt.DateTime `json:"modified_on" validate:"required"`
}

// UnmarshalServerSideExcludeRespResult unmarshals an instance of ServerSideExcludeRespResult from the specified map of raw messages.
func UnmarshalServerSideExcludeRespResult(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ServerSideExcludeRespResult)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		err = core.SDKErrorf(err, "", "id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "value", &obj.Value)
	if err != nil {
		err = core.SDKErrorf(err, "", "value-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "editable", &obj.Editable)
	if err != nil {
		err = core.SDKErrorf(err, "", "editable-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "modified_on", &obj.ModifiedOn)
	if err != nil {
		err = core.SDKErrorf(err, "", "modified_on-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// TlsClientAuthRespResult : Container for response information.
type TlsClientAuthRespResult struct {
	// ID.
	ID *string `json:"id" validate:"required"`

	// Value.
	Value *string `json:"value" validate:"required"`

	// Editable.
	Editable *bool `json:"editable" validate:"required"`

	// Modified date.
	ModifiedOn *strfmt.DateTime `json:"modified_on" validate:"required"`
}

// UnmarshalTlsClientAuthRespResult unmarshals an instance of TlsClientAuthRespResult from the specified map of raw messages.
func UnmarshalTlsClientAuthRespResult(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(TlsClientAuthRespResult)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		err = core.SDKErrorf(err, "", "id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "value", &obj.Value)
	if err != nil {
		err = core.SDKErrorf(err, "", "value-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "editable", &obj.Editable)
	if err != nil {
		err = core.SDKErrorf(err, "", "editable-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "modified_on", &obj.ModifiedOn)
	if err != nil {
		err = core.SDKErrorf(err, "", "modified_on-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// TrueClientIpRespResult : Container for response information.
type TrueClientIpRespResult struct {
	// ID.
	ID *string `json:"id" validate:"required"`

	// Value.
	Value *string `json:"value" validate:"required"`

	// Editable.
	Editable *bool `json:"editable" validate:"required"`

	// Modified date.
	ModifiedOn *strfmt.DateTime `json:"modified_on" validate:"required"`
}

// UnmarshalTrueClientIpRespResult unmarshals an instance of TrueClientIpRespResult from the specified map of raw messages.
func UnmarshalTrueClientIpRespResult(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(TrueClientIpRespResult)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		err = core.SDKErrorf(err, "", "id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "value", &obj.Value)
	if err != nil {
		err = core.SDKErrorf(err, "", "value-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "editable", &obj.Editable)
	if err != nil {
		err = core.SDKErrorf(err, "", "editable-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "modified_on", &obj.ModifiedOn)
	if err != nil {
		err = core.SDKErrorf(err, "", "modified_on-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// UpdateAlwaysUseHttpsOptions : The UpdateAlwaysUseHttps options.
type UpdateAlwaysUseHttpsOptions struct {
	// Value.
	Value *string `json:"value,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// Constants associated with the UpdateAlwaysUseHttpsOptions.Value property.
// Value.
const (
	UpdateAlwaysUseHttpsOptions_Value_Off = "off"
	UpdateAlwaysUseHttpsOptions_Value_On  = "on"
)

// NewUpdateAlwaysUseHttpsOptions : Instantiate UpdateAlwaysUseHttpsOptions
func (*ZonesSettingsV1) NewUpdateAlwaysUseHttpsOptions() *UpdateAlwaysUseHttpsOptions {
	return &UpdateAlwaysUseHttpsOptions{}
}

// SetValue : Allow user to set Value
func (_options *UpdateAlwaysUseHttpsOptions) SetValue(value string) *UpdateAlwaysUseHttpsOptions {
	_options.Value = core.StringPtr(value)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateAlwaysUseHttpsOptions) SetHeaders(param map[string]string) *UpdateAlwaysUseHttpsOptions {
	options.Headers = param
	return options
}

// UpdateAutomaticHttpsRewritesOptions : The UpdateAutomaticHttpsRewrites options.
type UpdateAutomaticHttpsRewritesOptions struct {
	// Value.
	Value *string `json:"value,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// Constants associated with the UpdateAutomaticHttpsRewritesOptions.Value property.
// Value.
const (
	UpdateAutomaticHttpsRewritesOptions_Value_Off = "off"
	UpdateAutomaticHttpsRewritesOptions_Value_On  = "on"
)

// NewUpdateAutomaticHttpsRewritesOptions : Instantiate UpdateAutomaticHttpsRewritesOptions
func (*ZonesSettingsV1) NewUpdateAutomaticHttpsRewritesOptions() *UpdateAutomaticHttpsRewritesOptions {
	return &UpdateAutomaticHttpsRewritesOptions{}
}

// SetValue : Allow user to set Value
func (_options *UpdateAutomaticHttpsRewritesOptions) SetValue(value string) *UpdateAutomaticHttpsRewritesOptions {
	_options.Value = core.StringPtr(value)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateAutomaticHttpsRewritesOptions) SetHeaders(param map[string]string) *UpdateAutomaticHttpsRewritesOptions {
	options.Headers = param
	return options
}

// UpdateBrotliOptions : The UpdateBrotli options.
type UpdateBrotliOptions struct {
	// Value.
	Value *string `json:"value,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// Constants associated with the UpdateBrotliOptions.Value property.
// Value.
const (
	UpdateBrotliOptions_Value_Off = "off"
	UpdateBrotliOptions_Value_On  = "on"
)

// NewUpdateBrotliOptions : Instantiate UpdateBrotliOptions
func (*ZonesSettingsV1) NewUpdateBrotliOptions() *UpdateBrotliOptions {
	return &UpdateBrotliOptions{}
}

// SetValue : Allow user to set Value
func (_options *UpdateBrotliOptions) SetValue(value string) *UpdateBrotliOptions {
	_options.Value = core.StringPtr(value)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateBrotliOptions) SetHeaders(param map[string]string) *UpdateBrotliOptions {
	options.Headers = param
	return options
}

// UpdateBrowserCheckOptions : The UpdateBrowserCheck options.
type UpdateBrowserCheckOptions struct {
	// Value.
	Value *string `json:"value,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// Constants associated with the UpdateBrowserCheckOptions.Value property.
// Value.
const (
	UpdateBrowserCheckOptions_Value_Off = "off"
	UpdateBrowserCheckOptions_Value_On  = "on"
)

// NewUpdateBrowserCheckOptions : Instantiate UpdateBrowserCheckOptions
func (*ZonesSettingsV1) NewUpdateBrowserCheckOptions() *UpdateBrowserCheckOptions {
	return &UpdateBrowserCheckOptions{}
}

// SetValue : Allow user to set Value
func (_options *UpdateBrowserCheckOptions) SetValue(value string) *UpdateBrowserCheckOptions {
	_options.Value = core.StringPtr(value)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateBrowserCheckOptions) SetHeaders(param map[string]string) *UpdateBrowserCheckOptions {
	options.Headers = param
	return options
}

// UpdateChallengeTtlOptions : The UpdateChallengeTTL options.
type UpdateChallengeTtlOptions struct {
	// Value.
	Value *int64 `json:"value,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewUpdateChallengeTtlOptions : Instantiate UpdateChallengeTtlOptions
func (*ZonesSettingsV1) NewUpdateChallengeTtlOptions() *UpdateChallengeTtlOptions {
	return &UpdateChallengeTtlOptions{}
}

// SetValue : Allow user to set Value
func (_options *UpdateChallengeTtlOptions) SetValue(value int64) *UpdateChallengeTtlOptions {
	_options.Value = core.Int64Ptr(value)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateChallengeTtlOptions) SetHeaders(param map[string]string) *UpdateChallengeTtlOptions {
	options.Headers = param
	return options
}

// UpdateCiphersOptions : The UpdateCiphers options.
type UpdateCiphersOptions struct {
	// Value.
	Value []string `json:"value,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// Constants associated with the UpdateCiphersOptions.Value property.
const (
	UpdateCiphersOptions_Value_Aes128GcmSha256            = "AES128-GCM-SHA256"
	UpdateCiphersOptions_Value_Aes128Sha                  = "AES128-SHA"
	UpdateCiphersOptions_Value_Aes128Sha256               = "AES128-SHA256"
	UpdateCiphersOptions_Value_Aes256GcmSha384            = "AES256-GCM-SHA384"
	UpdateCiphersOptions_Value_Aes256Sha                  = "AES256-SHA"
	UpdateCiphersOptions_Value_Aes256Sha256               = "AES256-SHA256"
	UpdateCiphersOptions_Value_DesCbc3Sha                 = "DES-CBC3-SHA"
	UpdateCiphersOptions_Value_EcdheEcdsaAes128GcmSha256  = "ECDHE-ECDSA-AES128-GCM-SHA256"
	UpdateCiphersOptions_Value_EcdheEcdsaAes128Sha        = "ECDHE-ECDSA-AES128-SHA"
	UpdateCiphersOptions_Value_EcdheEcdsaAes128Sha256     = "ECDHE-ECDSA-AES128-SHA256"
	UpdateCiphersOptions_Value_EcdheEcdsaAes256GcmSha384  = "ECDHE-ECDSA-AES256-GCM-SHA384"
	UpdateCiphersOptions_Value_EcdheEcdsaAes256Sha384     = "ECDHE-ECDSA-AES256-SHA384"
	UpdateCiphersOptions_Value_EcdheEcdsaChacha20Poly1305 = "ECDHE-ECDSA-CHACHA20-POLY1305"
	UpdateCiphersOptions_Value_EcdheRsaAes128GcmSha256    = "ECDHE-RSA-AES128-GCM-SHA256"
	UpdateCiphersOptions_Value_EcdheRsaAes128Sha          = "ECDHE-RSA-AES128-SHA"
	UpdateCiphersOptions_Value_EcdheRsaAes128Sha256       = "ECDHE-RSA-AES128-SHA256"
	UpdateCiphersOptions_Value_EcdheRsaAes256GcmSha384    = "ECDHE-RSA-AES256-GCM-SHA384"
	UpdateCiphersOptions_Value_EcdheRsaAes256Sha          = "ECDHE-RSA-AES256-SHA"
	UpdateCiphersOptions_Value_EcdheRsaAes256Sha384       = "ECDHE-RSA-AES256-SHA384"
	UpdateCiphersOptions_Value_EcdheRsaChacha20Poly1305   = "ECDHE-RSA-CHACHA20-POLY1305"
)

// NewUpdateCiphersOptions : Instantiate UpdateCiphersOptions
func (*ZonesSettingsV1) NewUpdateCiphersOptions() *UpdateCiphersOptions {
	return &UpdateCiphersOptions{}
}

// SetValue : Allow user to set Value
func (_options *UpdateCiphersOptions) SetValue(value []string) *UpdateCiphersOptions {
	_options.Value = value
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateCiphersOptions) SetHeaders(param map[string]string) *UpdateCiphersOptions {
	options.Headers = param
	return options
}

// UpdateEnableErrorPagesOnOptions : The UpdateEnableErrorPagesOn options.
type UpdateEnableErrorPagesOnOptions struct {
	// Value.
	Value *string `json:"value,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// Constants associated with the UpdateEnableErrorPagesOnOptions.Value property.
// Value.
const (
	UpdateEnableErrorPagesOnOptions_Value_Off = "off"
	UpdateEnableErrorPagesOnOptions_Value_On  = "on"
)

// NewUpdateEnableErrorPagesOnOptions : Instantiate UpdateEnableErrorPagesOnOptions
func (*ZonesSettingsV1) NewUpdateEnableErrorPagesOnOptions() *UpdateEnableErrorPagesOnOptions {
	return &UpdateEnableErrorPagesOnOptions{}
}

// SetValue : Allow user to set Value
func (_options *UpdateEnableErrorPagesOnOptions) SetValue(value string) *UpdateEnableErrorPagesOnOptions {
	_options.Value = core.StringPtr(value)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateEnableErrorPagesOnOptions) SetHeaders(param map[string]string) *UpdateEnableErrorPagesOnOptions {
	options.Headers = param
	return options
}

// UpdateHotlinkProtectionOptions : The UpdateHotlinkProtection options.
type UpdateHotlinkProtectionOptions struct {
	// Value.
	Value *string `json:"value,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// Constants associated with the UpdateHotlinkProtectionOptions.Value property.
// Value.
const (
	UpdateHotlinkProtectionOptions_Value_Off = "off"
	UpdateHotlinkProtectionOptions_Value_On  = "on"
)

// NewUpdateHotlinkProtectionOptions : Instantiate UpdateHotlinkProtectionOptions
func (*ZonesSettingsV1) NewUpdateHotlinkProtectionOptions() *UpdateHotlinkProtectionOptions {
	return &UpdateHotlinkProtectionOptions{}
}

// SetValue : Allow user to set Value
func (_options *UpdateHotlinkProtectionOptions) SetValue(value string) *UpdateHotlinkProtectionOptions {
	_options.Value = core.StringPtr(value)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateHotlinkProtectionOptions) SetHeaders(param map[string]string) *UpdateHotlinkProtectionOptions {
	options.Headers = param
	return options
}

// UpdateHttp2Options : The UpdateHttp2 options.
type UpdateHttp2Options struct {
	// Value.
	Value *string `json:"value,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// Constants associated with the UpdateHttp2Options.Value property.
// Value.
const (
	UpdateHttp2Options_Value_Off = "off"
	UpdateHttp2Options_Value_On  = "on"
)

// NewUpdateHttp2Options : Instantiate UpdateHttp2Options
func (*ZonesSettingsV1) NewUpdateHttp2Options() *UpdateHttp2Options {
	return &UpdateHttp2Options{}
}

// SetValue : Allow user to set Value
func (_options *UpdateHttp2Options) SetValue(value string) *UpdateHttp2Options {
	_options.Value = core.StringPtr(value)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateHttp2Options) SetHeaders(param map[string]string) *UpdateHttp2Options {
	options.Headers = param
	return options
}

// UpdateHttp3Options : The UpdateHttp3 options.
type UpdateHttp3Options struct {
	// Value.
	Value *string `json:"value,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// Constants associated with the UpdateHttp3Options.Value property.
// Value.
const (
	UpdateHttp3Options_Value_Off = "off"
	UpdateHttp3Options_Value_On  = "on"
)

// NewUpdateHttp3Options : Instantiate UpdateHttp3Options
func (*ZonesSettingsV1) NewUpdateHttp3Options() *UpdateHttp3Options {
	return &UpdateHttp3Options{}
}

// SetValue : Allow user to set Value
func (_options *UpdateHttp3Options) SetValue(value string) *UpdateHttp3Options {
	_options.Value = core.StringPtr(value)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateHttp3Options) SetHeaders(param map[string]string) *UpdateHttp3Options {
	options.Headers = param
	return options
}

// UpdateImageLoadOptimizationOptions : The UpdateImageLoadOptimization options.
type UpdateImageLoadOptimizationOptions struct {
	// Value.
	Value *string `json:"value,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// Constants associated with the UpdateImageLoadOptimizationOptions.Value property.
// Value.
const (
	UpdateImageLoadOptimizationOptions_Value_Off = "off"
	UpdateImageLoadOptimizationOptions_Value_On  = "on"
)

// NewUpdateImageLoadOptimizationOptions : Instantiate UpdateImageLoadOptimizationOptions
func (*ZonesSettingsV1) NewUpdateImageLoadOptimizationOptions() *UpdateImageLoadOptimizationOptions {
	return &UpdateImageLoadOptimizationOptions{}
}

// SetValue : Allow user to set Value
func (_options *UpdateImageLoadOptimizationOptions) SetValue(value string) *UpdateImageLoadOptimizationOptions {
	_options.Value = core.StringPtr(value)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateImageLoadOptimizationOptions) SetHeaders(param map[string]string) *UpdateImageLoadOptimizationOptions {
	options.Headers = param
	return options
}

// UpdateImageSizeOptimizationOptions : The UpdateImageSizeOptimization options.
type UpdateImageSizeOptimizationOptions struct {
	// Valid values are "lossy", "off", "lossless". "lossy" - The file size of JPEG images is reduced using lossy
	// compression, which may reduce visual quality. "off" - Disable Image Size Optimization. "lossless" - Reduce the size
	// of image files without impacting visual quality.
	Value *string `json:"value,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// Constants associated with the UpdateImageSizeOptimizationOptions.Value property.
// Valid values are "lossy", "off", "lossless". "lossy" - The file size of JPEG images is reduced using lossy
// compression, which may reduce visual quality. "off" - Disable Image Size Optimization. "lossless" - Reduce the size
// of image files without impacting visual quality.
const (
	UpdateImageSizeOptimizationOptions_Value_Lossless = "lossless"
	UpdateImageSizeOptimizationOptions_Value_Lossy    = "lossy"
	UpdateImageSizeOptimizationOptions_Value_Off      = "off"
)

// NewUpdateImageSizeOptimizationOptions : Instantiate UpdateImageSizeOptimizationOptions
func (*ZonesSettingsV1) NewUpdateImageSizeOptimizationOptions() *UpdateImageSizeOptimizationOptions {
	return &UpdateImageSizeOptimizationOptions{}
}

// SetValue : Allow user to set Value
func (_options *UpdateImageSizeOptimizationOptions) SetValue(value string) *UpdateImageSizeOptimizationOptions {
	_options.Value = core.StringPtr(value)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateImageSizeOptimizationOptions) SetHeaders(param map[string]string) *UpdateImageSizeOptimizationOptions {
	options.Headers = param
	return options
}

// UpdateIpGeolocationOptions : The UpdateIpGeolocation options.
type UpdateIpGeolocationOptions struct {
	// Value.
	Value *string `json:"value,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// Constants associated with the UpdateIpGeolocationOptions.Value property.
// Value.
const (
	UpdateIpGeolocationOptions_Value_Off = "off"
	UpdateIpGeolocationOptions_Value_On  = "on"
)

// NewUpdateIpGeolocationOptions : Instantiate UpdateIpGeolocationOptions
func (*ZonesSettingsV1) NewUpdateIpGeolocationOptions() *UpdateIpGeolocationOptions {
	return &UpdateIpGeolocationOptions{}
}

// SetValue : Allow user to set Value
func (_options *UpdateIpGeolocationOptions) SetValue(value string) *UpdateIpGeolocationOptions {
	_options.Value = core.StringPtr(value)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateIpGeolocationOptions) SetHeaders(param map[string]string) *UpdateIpGeolocationOptions {
	options.Headers = param
	return options
}

// UpdateIpv6Options : The UpdateIpv6 options.
type UpdateIpv6Options struct {
	// Value.
	Value *string `json:"value,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// Constants associated with the UpdateIpv6Options.Value property.
// Value.
const (
	UpdateIpv6Options_Value_Off = "off"
	UpdateIpv6Options_Value_On  = "on"
)

// NewUpdateIpv6Options : Instantiate UpdateIpv6Options
func (*ZonesSettingsV1) NewUpdateIpv6Options() *UpdateIpv6Options {
	return &UpdateIpv6Options{}
}

// SetValue : Allow user to set Value
func (_options *UpdateIpv6Options) SetValue(value string) *UpdateIpv6Options {
	_options.Value = core.StringPtr(value)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateIpv6Options) SetHeaders(param map[string]string) *UpdateIpv6Options {
	options.Headers = param
	return options
}

// UpdateLogRetentionOptions : The UpdateLogRetention options.
type UpdateLogRetentionOptions struct {
	// Full url-encoded cloud resource name (CRN) of resource instance.
	Crn *string `json:"crn" validate:"required,ne="`

	// Zone identifier.
	ZoneIdentifier *string `json:"zone_identifier" validate:"required,ne="`

	// True/false value to turn log retention on/off respectively.
	Flag *bool `json:"flag,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewUpdateLogRetentionOptions : Instantiate UpdateLogRetentionOptions
func (*ZonesSettingsV1) NewUpdateLogRetentionOptions(crn string, zoneIdentifier string) *UpdateLogRetentionOptions {
	return &UpdateLogRetentionOptions{
		Crn:            core.StringPtr(crn),
		ZoneIdentifier: core.StringPtr(zoneIdentifier),
	}
}

// SetCrn : Allow user to set Crn
func (_options *UpdateLogRetentionOptions) SetCrn(crn string) *UpdateLogRetentionOptions {
	_options.Crn = core.StringPtr(crn)
	return _options
}

// SetZoneIdentifier : Allow user to set ZoneIdentifier
func (_options *UpdateLogRetentionOptions) SetZoneIdentifier(zoneIdentifier string) *UpdateLogRetentionOptions {
	_options.ZoneIdentifier = core.StringPtr(zoneIdentifier)
	return _options
}

// SetFlag : Allow user to set Flag
func (_options *UpdateLogRetentionOptions) SetFlag(flag bool) *UpdateLogRetentionOptions {
	_options.Flag = core.BoolPtr(flag)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateLogRetentionOptions) SetHeaders(param map[string]string) *UpdateLogRetentionOptions {
	options.Headers = param
	return options
}

// UpdateMaxUploadOptions : The UpdateMaxUpload options.
type UpdateMaxUploadOptions struct {
	// Valid values(in MB) for "max_upload" are 100, 125, 150, 175, 200, 225, 250, 275, 300, 325, 350, 375, 400, 425, 450,
	// 475, 500. Values 225, 250, 275, 300, 325, 350, 375, 400, 425, 450, 475, 500 are only for Enterprise Plan.
	Value *int64 `json:"value,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewUpdateMaxUploadOptions : Instantiate UpdateMaxUploadOptions
func (*ZonesSettingsV1) NewUpdateMaxUploadOptions() *UpdateMaxUploadOptions {
	return &UpdateMaxUploadOptions{}
}

// SetValue : Allow user to set Value
func (_options *UpdateMaxUploadOptions) SetValue(value int64) *UpdateMaxUploadOptions {
	_options.Value = core.Int64Ptr(value)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateMaxUploadOptions) SetHeaders(param map[string]string) *UpdateMaxUploadOptions {
	options.Headers = param
	return options
}

// UpdateMinTlsVersionOptions : The UpdateMinTlsVersion options.
type UpdateMinTlsVersionOptions struct {
	// Value.
	Value *string `json:"value,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewUpdateMinTlsVersionOptions : Instantiate UpdateMinTlsVersionOptions
func (*ZonesSettingsV1) NewUpdateMinTlsVersionOptions() *UpdateMinTlsVersionOptions {
	return &UpdateMinTlsVersionOptions{}
}

// SetValue : Allow user to set Value
func (_options *UpdateMinTlsVersionOptions) SetValue(value string) *UpdateMinTlsVersionOptions {
	_options.Value = core.StringPtr(value)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateMinTlsVersionOptions) SetHeaders(param map[string]string) *UpdateMinTlsVersionOptions {
	options.Headers = param
	return options
}

// UpdateMinifyOptions : The UpdateMinify options.
type UpdateMinifyOptions struct {
	// Value.
	Value *MinifySettingValue `json:"value,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewUpdateMinifyOptions : Instantiate UpdateMinifyOptions
func (*ZonesSettingsV1) NewUpdateMinifyOptions() *UpdateMinifyOptions {
	return &UpdateMinifyOptions{}
}

// SetValue : Allow user to set Value
func (_options *UpdateMinifyOptions) SetValue(value *MinifySettingValue) *UpdateMinifyOptions {
	_options.Value = value
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateMinifyOptions) SetHeaders(param map[string]string) *UpdateMinifyOptions {
	options.Headers = param
	return options
}

// UpdateMobileRedirectOptions : The UpdateMobileRedirect options.
type UpdateMobileRedirectOptions struct {
	// Value.
	Value *MobileRedirecSettingValue `json:"value,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewUpdateMobileRedirectOptions : Instantiate UpdateMobileRedirectOptions
func (*ZonesSettingsV1) NewUpdateMobileRedirectOptions() *UpdateMobileRedirectOptions {
	return &UpdateMobileRedirectOptions{}
}

// SetValue : Allow user to set Value
func (_options *UpdateMobileRedirectOptions) SetValue(value *MobileRedirecSettingValue) *UpdateMobileRedirectOptions {
	_options.Value = value
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateMobileRedirectOptions) SetHeaders(param map[string]string) *UpdateMobileRedirectOptions {
	options.Headers = param
	return options
}

// UpdateOpportunisticEncryptionOptions : The UpdateOpportunisticEncryption options.
type UpdateOpportunisticEncryptionOptions struct {
	// Value.
	Value *string `json:"value,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// Constants associated with the UpdateOpportunisticEncryptionOptions.Value property.
// Value.
const (
	UpdateOpportunisticEncryptionOptions_Value_Off = "off"
	UpdateOpportunisticEncryptionOptions_Value_On  = "on"
)

// NewUpdateOpportunisticEncryptionOptions : Instantiate UpdateOpportunisticEncryptionOptions
func (*ZonesSettingsV1) NewUpdateOpportunisticEncryptionOptions() *UpdateOpportunisticEncryptionOptions {
	return &UpdateOpportunisticEncryptionOptions{}
}

// SetValue : Allow user to set Value
func (_options *UpdateOpportunisticEncryptionOptions) SetValue(value string) *UpdateOpportunisticEncryptionOptions {
	_options.Value = core.StringPtr(value)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateOpportunisticEncryptionOptions) SetHeaders(param map[string]string) *UpdateOpportunisticEncryptionOptions {
	options.Headers = param
	return options
}

// UpdateOpportunisticOnionOptions : The UpdateOpportunisticOnion options.
type UpdateOpportunisticOnionOptions struct {
	// Value.
	Value *string `json:"value,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// Constants associated with the UpdateOpportunisticOnionOptions.Value property.
// Value.
const (
	UpdateOpportunisticOnionOptions_Value_Off = "off"
	UpdateOpportunisticOnionOptions_Value_On  = "on"
)

// NewUpdateOpportunisticOnionOptions : Instantiate UpdateOpportunisticOnionOptions
func (*ZonesSettingsV1) NewUpdateOpportunisticOnionOptions() *UpdateOpportunisticOnionOptions {
	return &UpdateOpportunisticOnionOptions{}
}

// SetValue : Allow user to set Value
func (_options *UpdateOpportunisticOnionOptions) SetValue(value string) *UpdateOpportunisticOnionOptions {
	_options.Value = core.StringPtr(value)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateOpportunisticOnionOptions) SetHeaders(param map[string]string) *UpdateOpportunisticOnionOptions {
	options.Headers = param
	return options
}

// UpdateOriginMaxHttpVersionOptions : The UpdateOriginMaxHttpVersion options.
type UpdateOriginMaxHttpVersionOptions struct {
	// Value.
	Value *string `json:"value,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewUpdateOriginMaxHttpVersionOptions : Instantiate UpdateOriginMaxHttpVersionOptions
func (*ZonesSettingsV1) NewUpdateOriginMaxHttpVersionOptions() *UpdateOriginMaxHttpVersionOptions {
	return &UpdateOriginMaxHttpVersionOptions{}
}

// SetValue : Allow user to set Value
func (_options *UpdateOriginMaxHttpVersionOptions) SetValue(value string) *UpdateOriginMaxHttpVersionOptions {
	_options.Value = core.StringPtr(value)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateOriginMaxHttpVersionOptions) SetHeaders(param map[string]string) *UpdateOriginMaxHttpVersionOptions {
	options.Headers = param
	return options
}

// UpdateOriginPostQuantumEncryptionOptions : The UpdateOriginPostQuantumEncryption options.
type UpdateOriginPostQuantumEncryptionOptions struct {
	// Instructs CIS to use Post-Quantum (PQ) key agreement algorithms when connecting to your origin.
	// - `preferred`: Instructs CIS to opportunistically send a Post-Quantum keyshare in the first message to the origin
	// for fastest connections when the origin supports and prefers PQ.
	// - `supported`: The PQ algorithms are advertised but used only when requested by the origin.
	// - `off`: The PQ algorithms are not advertised.
	Value *string `json:"value,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// Constants associated with the UpdateOriginPostQuantumEncryptionOptions.Value property.
// Instructs CIS to use Post-Quantum (PQ) key agreement algorithms when connecting to your origin.
// - `preferred`: Instructs CIS to opportunistically send a Post-Quantum keyshare in the first message to the origin for
// fastest connections when the origin supports and prefers PQ.
// - `supported`: The PQ algorithms are advertised but used only when requested by the origin.
// - `off`: The PQ algorithms are not advertised.
const (
	UpdateOriginPostQuantumEncryptionOptions_Value_Off       = "off"
	UpdateOriginPostQuantumEncryptionOptions_Value_Preferred = "preferred"
	UpdateOriginPostQuantumEncryptionOptions_Value_Supported = "supported"
)

// NewUpdateOriginPostQuantumEncryptionOptions : Instantiate UpdateOriginPostQuantumEncryptionOptions
func (*ZonesSettingsV1) NewUpdateOriginPostQuantumEncryptionOptions() *UpdateOriginPostQuantumEncryptionOptions {
	return &UpdateOriginPostQuantumEncryptionOptions{}
}

// SetValue : Allow user to set Value
func (_options *UpdateOriginPostQuantumEncryptionOptions) SetValue(value string) *UpdateOriginPostQuantumEncryptionOptions {
	_options.Value = core.StringPtr(value)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateOriginPostQuantumEncryptionOptions) SetHeaders(param map[string]string) *UpdateOriginPostQuantumEncryptionOptions {
	options.Headers = param
	return options
}

// UpdatePrefetchPreloadOptions : The UpdatePrefetchPreload options.
type UpdatePrefetchPreloadOptions struct {
	// Value.
	Value *string `json:"value,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// Constants associated with the UpdatePrefetchPreloadOptions.Value property.
// Value.
const (
	UpdatePrefetchPreloadOptions_Value_Off = "off"
	UpdatePrefetchPreloadOptions_Value_On  = "on"
)

// NewUpdatePrefetchPreloadOptions : Instantiate UpdatePrefetchPreloadOptions
func (*ZonesSettingsV1) NewUpdatePrefetchPreloadOptions() *UpdatePrefetchPreloadOptions {
	return &UpdatePrefetchPreloadOptions{}
}

// SetValue : Allow user to set Value
func (_options *UpdatePrefetchPreloadOptions) SetValue(value string) *UpdatePrefetchPreloadOptions {
	_options.Value = core.StringPtr(value)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdatePrefetchPreloadOptions) SetHeaders(param map[string]string) *UpdatePrefetchPreloadOptions {
	options.Headers = param
	return options
}

// UpdateProxyReadTimeoutOptions : The UpdateProxyReadTimeout options.
type UpdateProxyReadTimeoutOptions struct {
	// Value.
	Value *float64 `json:"value,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewUpdateProxyReadTimeoutOptions : Instantiate UpdateProxyReadTimeoutOptions
func (*ZonesSettingsV1) NewUpdateProxyReadTimeoutOptions() *UpdateProxyReadTimeoutOptions {
	return &UpdateProxyReadTimeoutOptions{}
}

// SetValue : Allow user to set Value
func (_options *UpdateProxyReadTimeoutOptions) SetValue(value float64) *UpdateProxyReadTimeoutOptions {
	_options.Value = core.Float64Ptr(value)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateProxyReadTimeoutOptions) SetHeaders(param map[string]string) *UpdateProxyReadTimeoutOptions {
	options.Headers = param
	return options
}

// UpdatePseudoIpv4Options : The UpdatePseudoIpv4 options.
type UpdatePseudoIpv4Options struct {
	// Value.
	Value *string `json:"value,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// Constants associated with the UpdatePseudoIpv4Options.Value property.
// Value.
const (
	UpdatePseudoIpv4Options_Value_AddHeader       = "add_header"
	UpdatePseudoIpv4Options_Value_Off             = "off"
	UpdatePseudoIpv4Options_Value_OverwriteHeader = "overwrite_header"
)

// NewUpdatePseudoIpv4Options : Instantiate UpdatePseudoIpv4Options
func (*ZonesSettingsV1) NewUpdatePseudoIpv4Options() *UpdatePseudoIpv4Options {
	return &UpdatePseudoIpv4Options{}
}

// SetValue : Allow user to set Value
func (_options *UpdatePseudoIpv4Options) SetValue(value string) *UpdatePseudoIpv4Options {
	_options.Value = core.StringPtr(value)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdatePseudoIpv4Options) SetHeaders(param map[string]string) *UpdatePseudoIpv4Options {
	options.Headers = param
	return options
}

// UpdateResponseBufferingOptions : The UpdateResponseBuffering options.
type UpdateResponseBufferingOptions struct {
	// Value.
	Value *string `json:"value,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// Constants associated with the UpdateResponseBufferingOptions.Value property.
// Value.
const (
	UpdateResponseBufferingOptions_Value_Off = "off"
	UpdateResponseBufferingOptions_Value_On  = "on"
)

// NewUpdateResponseBufferingOptions : Instantiate UpdateResponseBufferingOptions
func (*ZonesSettingsV1) NewUpdateResponseBufferingOptions() *UpdateResponseBufferingOptions {
	return &UpdateResponseBufferingOptions{}
}

// SetValue : Allow user to set Value
func (_options *UpdateResponseBufferingOptions) SetValue(value string) *UpdateResponseBufferingOptions {
	_options.Value = core.StringPtr(value)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateResponseBufferingOptions) SetHeaders(param map[string]string) *UpdateResponseBufferingOptions {
	options.Headers = param
	return options
}

// UpdateScriptLoadOptimizationOptions : The UpdateScriptLoadOptimization options.
type UpdateScriptLoadOptimizationOptions struct {
	// Value.
	Value *string `json:"value,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// Constants associated with the UpdateScriptLoadOptimizationOptions.Value property.
// Value.
const (
	UpdateScriptLoadOptimizationOptions_Value_Off = "off"
	UpdateScriptLoadOptimizationOptions_Value_On  = "on"
)

// NewUpdateScriptLoadOptimizationOptions : Instantiate UpdateScriptLoadOptimizationOptions
func (*ZonesSettingsV1) NewUpdateScriptLoadOptimizationOptions() *UpdateScriptLoadOptimizationOptions {
	return &UpdateScriptLoadOptimizationOptions{}
}

// SetValue : Allow user to set Value
func (_options *UpdateScriptLoadOptimizationOptions) SetValue(value string) *UpdateScriptLoadOptimizationOptions {
	_options.Value = core.StringPtr(value)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateScriptLoadOptimizationOptions) SetHeaders(param map[string]string) *UpdateScriptLoadOptimizationOptions {
	options.Headers = param
	return options
}

// UpdateSecurityHeaderOptions : The UpdateSecurityHeader options.
type UpdateSecurityHeaderOptions struct {
	// Value.
	Value *SecurityHeaderSettingValue `json:"value,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewUpdateSecurityHeaderOptions : Instantiate UpdateSecurityHeaderOptions
func (*ZonesSettingsV1) NewUpdateSecurityHeaderOptions() *UpdateSecurityHeaderOptions {
	return &UpdateSecurityHeaderOptions{}
}

// SetValue : Allow user to set Value
func (_options *UpdateSecurityHeaderOptions) SetValue(value *SecurityHeaderSettingValue) *UpdateSecurityHeaderOptions {
	_options.Value = value
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateSecurityHeaderOptions) SetHeaders(param map[string]string) *UpdateSecurityHeaderOptions {
	options.Headers = param
	return options
}

// UpdateServerSideExcludeOptions : The UpdateServerSideExclude options.
type UpdateServerSideExcludeOptions struct {
	// Value.
	Value *string `json:"value,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// Constants associated with the UpdateServerSideExcludeOptions.Value property.
// Value.
const (
	UpdateServerSideExcludeOptions_Value_Off = "off"
	UpdateServerSideExcludeOptions_Value_On  = "on"
)

// NewUpdateServerSideExcludeOptions : Instantiate UpdateServerSideExcludeOptions
func (*ZonesSettingsV1) NewUpdateServerSideExcludeOptions() *UpdateServerSideExcludeOptions {
	return &UpdateServerSideExcludeOptions{}
}

// SetValue : Allow user to set Value
func (_options *UpdateServerSideExcludeOptions) SetValue(value string) *UpdateServerSideExcludeOptions {
	_options.Value = core.StringPtr(value)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateServerSideExcludeOptions) SetHeaders(param map[string]string) *UpdateServerSideExcludeOptions {
	options.Headers = param
	return options
}

// UpdateTlsClientAuthOptions : The UpdateTlsClientAuth options.
type UpdateTlsClientAuthOptions struct {
	// Value.
	Value *string `json:"value,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// Constants associated with the UpdateTlsClientAuthOptions.Value property.
// Value.
const (
	UpdateTlsClientAuthOptions_Value_Off = "off"
	UpdateTlsClientAuthOptions_Value_On  = "on"
)

// NewUpdateTlsClientAuthOptions : Instantiate UpdateTlsClientAuthOptions
func (*ZonesSettingsV1) NewUpdateTlsClientAuthOptions() *UpdateTlsClientAuthOptions {
	return &UpdateTlsClientAuthOptions{}
}

// SetValue : Allow user to set Value
func (_options *UpdateTlsClientAuthOptions) SetValue(value string) *UpdateTlsClientAuthOptions {
	_options.Value = core.StringPtr(value)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateTlsClientAuthOptions) SetHeaders(param map[string]string) *UpdateTlsClientAuthOptions {
	options.Headers = param
	return options
}

// UpdateTrueClientIpOptions : The UpdateTrueClientIp options.
type UpdateTrueClientIpOptions struct {
	// Value.
	Value *string `json:"value,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// Constants associated with the UpdateTrueClientIpOptions.Value property.
// Value.
const (
	UpdateTrueClientIpOptions_Value_Off = "off"
	UpdateTrueClientIpOptions_Value_On  = "on"
)

// NewUpdateTrueClientIpOptions : Instantiate UpdateTrueClientIpOptions
func (*ZonesSettingsV1) NewUpdateTrueClientIpOptions() *UpdateTrueClientIpOptions {
	return &UpdateTrueClientIpOptions{}
}

// SetValue : Allow user to set Value
func (_options *UpdateTrueClientIpOptions) SetValue(value string) *UpdateTrueClientIpOptions {
	_options.Value = core.StringPtr(value)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateTrueClientIpOptions) SetHeaders(param map[string]string) *UpdateTrueClientIpOptions {
	options.Headers = param
	return options
}

// UpdateWebApplicationFirewallOptions : The UpdateWebApplicationFirewall options.
type UpdateWebApplicationFirewallOptions struct {
	// Value.
	Value *string `json:"value,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// Constants associated with the UpdateWebApplicationFirewallOptions.Value property.
// Value.
const (
	UpdateWebApplicationFirewallOptions_Value_Off = "off"
	UpdateWebApplicationFirewallOptions_Value_On  = "on"
)

// NewUpdateWebApplicationFirewallOptions : Instantiate UpdateWebApplicationFirewallOptions
func (*ZonesSettingsV1) NewUpdateWebApplicationFirewallOptions() *UpdateWebApplicationFirewallOptions {
	return &UpdateWebApplicationFirewallOptions{}
}

// SetValue : Allow user to set Value
func (_options *UpdateWebApplicationFirewallOptions) SetValue(value string) *UpdateWebApplicationFirewallOptions {
	_options.Value = core.StringPtr(value)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateWebApplicationFirewallOptions) SetHeaders(param map[string]string) *UpdateWebApplicationFirewallOptions {
	options.Headers = param
	return options
}

// UpdateWebSocketsOptions : The UpdateWebSockets options.
type UpdateWebSocketsOptions struct {
	// Value.
	Value *string `json:"value,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// Constants associated with the UpdateWebSocketsOptions.Value property.
// Value.
const (
	UpdateWebSocketsOptions_Value_Off = "off"
	UpdateWebSocketsOptions_Value_On  = "on"
)

// NewUpdateWebSocketsOptions : Instantiate UpdateWebSocketsOptions
func (*ZonesSettingsV1) NewUpdateWebSocketsOptions() *UpdateWebSocketsOptions {
	return &UpdateWebSocketsOptions{}
}

// SetValue : Allow user to set Value
func (_options *UpdateWebSocketsOptions) SetValue(value string) *UpdateWebSocketsOptions {
	_options.Value = core.StringPtr(value)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateWebSocketsOptions) SetHeaders(param map[string]string) *UpdateWebSocketsOptions {
	options.Headers = param
	return options
}

// UpdateZoneCnameFlatteningOptions : The UpdateZoneCnameFlattening options.
type UpdateZoneCnameFlatteningOptions struct {
	// Valid values are "flatten_at_root", "flatten_all". "flatten_at_root" - Flatten CNAME at root domain. This is the
	// default value. "flatten_all" - Flatten all CNAME records under your domain.
	Value *string `json:"value,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// Constants associated with the UpdateZoneCnameFlatteningOptions.Value property.
// Valid values are "flatten_at_root", "flatten_all". "flatten_at_root" - Flatten CNAME at root domain. This is the
// default value. "flatten_all" - Flatten all CNAME records under your domain.
const (
	UpdateZoneCnameFlatteningOptions_Value_FlattenAll    = "flatten_all"
	UpdateZoneCnameFlatteningOptions_Value_FlattenAtRoot = "flatten_at_root"
)

// NewUpdateZoneCnameFlatteningOptions : Instantiate UpdateZoneCnameFlatteningOptions
func (*ZonesSettingsV1) NewUpdateZoneCnameFlatteningOptions() *UpdateZoneCnameFlatteningOptions {
	return &UpdateZoneCnameFlatteningOptions{}
}

// SetValue : Allow user to set Value
func (_options *UpdateZoneCnameFlatteningOptions) SetValue(value string) *UpdateZoneCnameFlatteningOptions {
	_options.Value = core.StringPtr(value)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateZoneCnameFlatteningOptions) SetHeaders(param map[string]string) *UpdateZoneCnameFlatteningOptions {
	options.Headers = param
	return options
}

// UpdateZoneDnssecOptions : The UpdateZoneDnssec options.
type UpdateZoneDnssecOptions struct {
	// Status.
	Status *string `json:"status,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// Constants associated with the UpdateZoneDnssecOptions.Status property.
// Status.
const (
	UpdateZoneDnssecOptions_Status_Active   = "active"
	UpdateZoneDnssecOptions_Status_Disabled = "disabled"
)

// NewUpdateZoneDnssecOptions : Instantiate UpdateZoneDnssecOptions
func (*ZonesSettingsV1) NewUpdateZoneDnssecOptions() *UpdateZoneDnssecOptions {
	return &UpdateZoneDnssecOptions{}
}

// SetStatus : Allow user to set Status
func (_options *UpdateZoneDnssecOptions) SetStatus(status string) *UpdateZoneDnssecOptions {
	_options.Status = core.StringPtr(status)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateZoneDnssecOptions) SetHeaders(param map[string]string) *UpdateZoneDnssecOptions {
	options.Headers = param
	return options
}

// WafRespResult : Container for response information.
type WafRespResult struct {
	// ID.
	ID *string `json:"id" validate:"required"`

	// Value.
	Value *string `json:"value" validate:"required"`

	// Editable.
	Editable *bool `json:"editable" validate:"required"`

	// Modified date.
	ModifiedOn *strfmt.DateTime `json:"modified_on" validate:"required"`
}

// UnmarshalWafRespResult unmarshals an instance of WafRespResult from the specified map of raw messages.
func UnmarshalWafRespResult(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(WafRespResult)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		err = core.SDKErrorf(err, "", "id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "value", &obj.Value)
	if err != nil {
		err = core.SDKErrorf(err, "", "value-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "editable", &obj.Editable)
	if err != nil {
		err = core.SDKErrorf(err, "", "editable-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "modified_on", &obj.ModifiedOn)
	if err != nil {
		err = core.SDKErrorf(err, "", "modified_on-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// WebsocketsRespResult : Container for response information.
type WebsocketsRespResult struct {
	// ID.
	ID *string `json:"id" validate:"required"`

	// Value.
	Value *string `json:"value" validate:"required"`

	// Editable.
	Editable *bool `json:"editable" validate:"required"`

	// Modified date.
	ModifiedOn *strfmt.DateTime `json:"modified_on" validate:"required"`
}

// UnmarshalWebsocketsRespResult unmarshals an instance of WebsocketsRespResult from the specified map of raw messages.
func UnmarshalWebsocketsRespResult(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(WebsocketsRespResult)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		err = core.SDKErrorf(err, "", "id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "value", &obj.Value)
	if err != nil {
		err = core.SDKErrorf(err, "", "value-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "editable", &obj.Editable)
	if err != nil {
		err = core.SDKErrorf(err, "", "editable-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "modified_on", &obj.ModifiedOn)
	if err != nil {
		err = core.SDKErrorf(err, "", "modified_on-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ZonesDnssecRespResult : Container for response information.
type ZonesDnssecRespResult struct {
	// Status.
	Status *string `json:"status,omitempty"`

	// Flags.
	Flags *int64 `json:"flags,omitempty"`

	// Algorithm.
	Algorithm *string `json:"algorithm,omitempty"`

	// Key type.
	KeyType *string `json:"key_type,omitempty"`

	// Digest type.
	DigestType *string `json:"digest_type,omitempty"`

	// Digest algorithm.
	DigestAlgorithm *string `json:"digest_algorithm,omitempty"`

	// Digest.
	Digest *string `json:"digest,omitempty"`

	// DS.
	Ds *string `json:"ds,omitempty"`

	// Key tag.
	KeyTag *int64 `json:"key_tag,omitempty"`

	// Public key.
	PublicKey *string `json:"public_key,omitempty"`
}

// Constants associated with the ZonesDnssecRespResult.Status property.
// Status.
const (
	ZonesDnssecRespResult_Status_Active          = "active"
	ZonesDnssecRespResult_Status_Disabled        = "disabled"
	ZonesDnssecRespResult_Status_Error           = "error"
	ZonesDnssecRespResult_Status_Pending         = "pending"
	ZonesDnssecRespResult_Status_PendingDisabled = "pending-disabled"
)

// UnmarshalZonesDnssecRespResult unmarshals an instance of ZonesDnssecRespResult from the specified map of raw messages.
func UnmarshalZonesDnssecRespResult(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ZonesDnssecRespResult)
	err = core.UnmarshalPrimitive(m, "status", &obj.Status)
	if err != nil {
		err = core.SDKErrorf(err, "", "status-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "flags", &obj.Flags)
	if err != nil {
		err = core.SDKErrorf(err, "", "flags-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "algorithm", &obj.Algorithm)
	if err != nil {
		err = core.SDKErrorf(err, "", "algorithm-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "key_type", &obj.KeyType)
	if err != nil {
		err = core.SDKErrorf(err, "", "key_type-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "digest_type", &obj.DigestType)
	if err != nil {
		err = core.SDKErrorf(err, "", "digest_type-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "digest_algorithm", &obj.DigestAlgorithm)
	if err != nil {
		err = core.SDKErrorf(err, "", "digest_algorithm-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "digest", &obj.Digest)
	if err != nil {
		err = core.SDKErrorf(err, "", "digest-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "ds", &obj.Ds)
	if err != nil {
		err = core.SDKErrorf(err, "", "ds-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "key_tag", &obj.KeyTag)
	if err != nil {
		err = core.SDKErrorf(err, "", "key_tag-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "public_key", &obj.PublicKey)
	if err != nil {
		err = core.SDKErrorf(err, "", "public_key-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AlwaysUseHttpsResp : Always use http response.
type AlwaysUseHttpsResp struct {
	// Container for response information.
	Result *AlwaysUseHttpsRespResult `json:"result" validate:"required"`

	// Was the get successful.
	Success *bool `json:"success" validate:"required"`

	// Array of errors encountered.
	Errors [][]string `json:"errors" validate:"required"`

	// Array of messages returned.
	Messages [][]string `json:"messages" validate:"required"`
}

// UnmarshalAlwaysUseHttpsResp unmarshals an instance of AlwaysUseHttpsResp from the specified map of raw messages.
func UnmarshalAlwaysUseHttpsResp(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AlwaysUseHttpsResp)
	err = core.UnmarshalModel(m, "result", &obj.Result, UnmarshalAlwaysUseHttpsRespResult)
	if err != nil {
		err = core.SDKErrorf(err, "", "result-error", common.GetComponentInfo())
		return
	}
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
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AutomaticHttpsRewritesResp : automatic https rewrite response.
type AutomaticHttpsRewritesResp struct {
	// Container for response information.
	Result *AutomaticHttpsRewritesRespResult `json:"result" validate:"required"`

	// Was the get successful.
	Success *bool `json:"success" validate:"required"`

	// Array of errors encountered.
	Errors [][]string `json:"errors" validate:"required"`

	// Array of messages returned.
	Messages [][]string `json:"messages" validate:"required"`
}

// UnmarshalAutomaticHttpsRewritesResp unmarshals an instance of AutomaticHttpsRewritesResp from the specified map of raw messages.
func UnmarshalAutomaticHttpsRewritesResp(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AutomaticHttpsRewritesResp)
	err = core.UnmarshalModel(m, "result", &obj.Result, UnmarshalAutomaticHttpsRewritesRespResult)
	if err != nil {
		err = core.SDKErrorf(err, "", "result-error", common.GetComponentInfo())
		return
	}
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
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// BrotliResp : Brotli response.
type BrotliResp struct {
	// Container for response information.
	Result *BrotliRespResult `json:"result" validate:"required"`

	// Was the get successful.
	Success *bool `json:"success" validate:"required"`

	// Array of errors encountered.
	Errors [][]string `json:"errors" validate:"required"`

	// Array of messages returned.
	Messages [][]string `json:"messages" validate:"required"`
}

// UnmarshalBrotliResp unmarshals an instance of BrotliResp from the specified map of raw messages.
func UnmarshalBrotliResp(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(BrotliResp)
	err = core.UnmarshalModel(m, "result", &obj.Result, UnmarshalBrotliRespResult)
	if err != nil {
		err = core.SDKErrorf(err, "", "result-error", common.GetComponentInfo())
		return
	}
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
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// BrowserCheckResp : Browser Check response.
type BrowserCheckResp struct {
	// Container for response information.
	Result *BrowserCheckRespResult `json:"result" validate:"required"`

	// Was the get successful.
	Success *bool `json:"success" validate:"required"`

	// Array of errors encountered.
	Errors [][]string `json:"errors" validate:"required"`

	// Array of messages returned.
	Messages [][]string `json:"messages" validate:"required"`
}

// UnmarshalBrowserCheckResp unmarshals an instance of BrowserCheckResp from the specified map of raw messages.
func UnmarshalBrowserCheckResp(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(BrowserCheckResp)
	err = core.UnmarshalModel(m, "result", &obj.Result, UnmarshalBrowserCheckRespResult)
	if err != nil {
		err = core.SDKErrorf(err, "", "result-error", common.GetComponentInfo())
		return
	}
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
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ChallengeTtlResp : challenge TTL response.
type ChallengeTtlResp struct {
	// Container for response information.
	Result *ChallengeTtlRespResult `json:"result" validate:"required"`

	// Was the get successful.
	Success *bool `json:"success" validate:"required"`

	// Array of errors encountered.
	Errors [][]string `json:"errors" validate:"required"`

	// Array of messages returned.
	Messages [][]string `json:"messages" validate:"required"`
}

// UnmarshalChallengeTtlResp unmarshals an instance of ChallengeTtlResp from the specified map of raw messages.
func UnmarshalChallengeTtlResp(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ChallengeTtlResp)
	err = core.UnmarshalModel(m, "result", &obj.Result, UnmarshalChallengeTtlRespResult)
	if err != nil {
		err = core.SDKErrorf(err, "", "result-error", common.GetComponentInfo())
		return
	}
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
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CiphersResp : Ciphers response.
type CiphersResp struct {
	// Container for response information.
	Result *CiphersRespResult `json:"result" validate:"required"`

	// Was the get successful.
	Success *bool `json:"success" validate:"required"`

	// Array of errors encountered.
	Errors [][]string `json:"errors" validate:"required"`

	// Array of messages returned.
	Messages [][]string `json:"messages" validate:"required"`
}

// UnmarshalCiphersResp unmarshals an instance of CiphersResp from the specified map of raw messages.
func UnmarshalCiphersResp(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CiphersResp)
	err = core.UnmarshalModel(m, "result", &obj.Result, UnmarshalCiphersRespResult)
	if err != nil {
		err = core.SDKErrorf(err, "", "result-error", common.GetComponentInfo())
		return
	}
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
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CnameFlatteningResponse : CNAME Flattening response.
type CnameFlatteningResponse struct {
	// id.
	ID *string `json:"id,omitempty"`

	// value.
	Value *string `json:"value,omitempty"`

	// Date when it is modified.
	ModifiedOn *strfmt.DateTime `json:"modified_on,omitempty"`

	// editable.
	Editable *bool `json:"editable,omitempty"`
}

// Constants associated with the CnameFlatteningResponse.Value property.
// value.
const (
	CnameFlatteningResponse_Value_FlattenAll    = "flatten_all"
	CnameFlatteningResponse_Value_FlattenAtRoot = "flatten_at_root"
)

// UnmarshalCnameFlatteningResponse unmarshals an instance of CnameFlatteningResponse from the specified map of raw messages.
func UnmarshalCnameFlatteningResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CnameFlatteningResponse)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		err = core.SDKErrorf(err, "", "id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "value", &obj.Value)
	if err != nil {
		err = core.SDKErrorf(err, "", "value-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "modified_on", &obj.ModifiedOn)
	if err != nil {
		err = core.SDKErrorf(err, "", "modified_on-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "editable", &obj.Editable)
	if err != nil {
		err = core.SDKErrorf(err, "", "editable-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// HotlinkProtectionResp : Hotlink Protection response.
type HotlinkProtectionResp struct {
	// Container for response information.
	Result *HotlinkProtectionRespResult `json:"result" validate:"required"`

	// Was the get successful.
	Success *bool `json:"success" validate:"required"`

	// Array of errors encountered.
	Errors [][]string `json:"errors" validate:"required"`

	// Array of messages returned.
	Messages [][]string `json:"messages" validate:"required"`
}

// UnmarshalHotlinkProtectionResp unmarshals an instance of HotlinkProtectionResp from the specified map of raw messages.
func UnmarshalHotlinkProtectionResp(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(HotlinkProtectionResp)
	err = core.UnmarshalModel(m, "result", &obj.Result, UnmarshalHotlinkProtectionRespResult)
	if err != nil {
		err = core.SDKErrorf(err, "", "result-error", common.GetComponentInfo())
		return
	}
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
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Http2Resp : HTTP2 Response.
type Http2Resp struct {
	// Container for response information.
	Result *Http2RespResult `json:"result" validate:"required"`

	// Was the get successful.
	Success *bool `json:"success" validate:"required"`

	// Array of errors encountered.
	Errors [][]string `json:"errors" validate:"required"`

	// Array of messages returned.
	Messages [][]string `json:"messages" validate:"required"`
}

// UnmarshalHttp2Resp unmarshals an instance of Http2Resp from the specified map of raw messages.
func UnmarshalHttp2Resp(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Http2Resp)
	err = core.UnmarshalModel(m, "result", &obj.Result, UnmarshalHttp2RespResult)
	if err != nil {
		err = core.SDKErrorf(err, "", "result-error", common.GetComponentInfo())
		return
	}
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
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Http3Resp : HTTP3 Response.
type Http3Resp struct {
	// Container for response information.
	Result *Http3RespResult `json:"result" validate:"required"`

	// Was the get successful.
	Success *bool `json:"success" validate:"required"`

	// Array of errors encountered.
	Errors [][]string `json:"errors" validate:"required"`

	// Array of messages returned.
	Messages [][]string `json:"messages" validate:"required"`
}

// UnmarshalHttp3Resp unmarshals an instance of Http3Resp from the specified map of raw messages.
func UnmarshalHttp3Resp(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Http3Resp)
	err = core.UnmarshalModel(m, "result", &obj.Result, UnmarshalHttp3RespResult)
	if err != nil {
		err = core.SDKErrorf(err, "", "result-error", common.GetComponentInfo())
		return
	}
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
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ImageLoadOptimizationResp : Image Load Optimization response.
type ImageLoadOptimizationResp struct {
	// Container for response information.
	Result *ImageLoadOptimizationRespResult `json:"result" validate:"required"`

	// Was the get successful.
	Success *bool `json:"success" validate:"required"`

	// Array of errors encountered.
	Errors [][]string `json:"errors" validate:"required"`

	// Array of messages returned.
	Messages [][]string `json:"messages" validate:"required"`
}

// UnmarshalImageLoadOptimizationResp unmarshals an instance of ImageLoadOptimizationResp from the specified map of raw messages.
func UnmarshalImageLoadOptimizationResp(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ImageLoadOptimizationResp)
	err = core.UnmarshalModel(m, "result", &obj.Result, UnmarshalImageLoadOptimizationRespResult)
	if err != nil {
		err = core.SDKErrorf(err, "", "result-error", common.GetComponentInfo())
		return
	}
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
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ImageSizeOptimizationResp : Image size optimization response.
type ImageSizeOptimizationResp struct {
	// Container for response information.
	Result *ImageSizeOptimizationRespResult `json:"result" validate:"required"`

	// Was the get successful.
	Success *bool `json:"success" validate:"required"`

	// Array of errors encountered.
	Errors [][]string `json:"errors" validate:"required"`

	// Array of messages returned.
	Messages [][]string `json:"messages" validate:"required"`
}

// UnmarshalImageSizeOptimizationResp unmarshals an instance of ImageSizeOptimizationResp from the specified map of raw messages.
func UnmarshalImageSizeOptimizationResp(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ImageSizeOptimizationResp)
	err = core.UnmarshalModel(m, "result", &obj.Result, UnmarshalImageSizeOptimizationRespResult)
	if err != nil {
		err = core.SDKErrorf(err, "", "result-error", common.GetComponentInfo())
		return
	}
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
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// IpGeolocationResp : IP Geolocation response.
type IpGeolocationResp struct {
	// Container for response information.
	Result *IpGeolocationRespResult `json:"result" validate:"required"`

	// Was the get successful.
	Success *bool `json:"success" validate:"required"`

	// Array of errors encountered.
	Errors [][]string `json:"errors" validate:"required"`

	// Array of messages returned.
	Messages [][]string `json:"messages" validate:"required"`
}

// UnmarshalIpGeolocationResp unmarshals an instance of IpGeolocationResp from the specified map of raw messages.
func UnmarshalIpGeolocationResp(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(IpGeolocationResp)
	err = core.UnmarshalModel(m, "result", &obj.Result, UnmarshalIpGeolocationRespResult)
	if err != nil {
		err = core.SDKErrorf(err, "", "result-error", common.GetComponentInfo())
		return
	}
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
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Ipv6Resp : IPv6 Response.
type Ipv6Resp struct {
	// Container for response information.
	Result *Ipv6RespResult `json:"result" validate:"required"`

	// Was the get successful.
	Success *bool `json:"success" validate:"required"`

	// Array of errors encountered.
	Errors [][]string `json:"errors" validate:"required"`

	// Array of messages returned.
	Messages [][]string `json:"messages" validate:"required"`
}

// UnmarshalIpv6Resp unmarshals an instance of Ipv6Resp from the specified map of raw messages.
func UnmarshalIpv6Resp(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Ipv6Resp)
	err = core.UnmarshalModel(m, "result", &obj.Result, UnmarshalIpv6RespResult)
	if err != nil {
		err = core.SDKErrorf(err, "", "result-error", common.GetComponentInfo())
		return
	}
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
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// LogRetentionResp : Schema for the response to a GET call for the log retention setting.
type LogRetentionResp struct {
	// Boolean flag indicating whether hte API call was successful in retrieving the requested data.
	Success *bool `json:"success" validate:"required"`

	Result *LogRetentionRespResult `json:"result" validate:"required"`

	// Array of errors messages.
	Errors []string `json:"errors" validate:"required"`

	// Array of additional messages.
	Messages []string `json:"messages" validate:"required"`
}

// UnmarshalLogRetentionResp unmarshals an instance of LogRetentionResp from the specified map of raw messages.
func UnmarshalLogRetentionResp(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(LogRetentionResp)
	err = core.UnmarshalPrimitive(m, "success", &obj.Success)
	if err != nil {
		err = core.SDKErrorf(err, "", "success-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "result", &obj.Result, UnmarshalLogRetentionRespResult)
	if err != nil {
		err = core.SDKErrorf(err, "", "result-error", common.GetComponentInfo())
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
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// MaxUploadResp : Maximum upload response.
type MaxUploadResp struct {
	// Container for response information.
	Result *MaxUploadRespResult `json:"result" validate:"required"`

	// Was the get successful.
	Success *bool `json:"success" validate:"required"`

	// Array of errors encountered.
	Errors [][]string `json:"errors" validate:"required"`

	// Array of messages returned.
	Messages [][]string `json:"messages" validate:"required"`
}

// UnmarshalMaxUploadResp unmarshals an instance of MaxUploadResp from the specified map of raw messages.
func UnmarshalMaxUploadResp(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(MaxUploadResp)
	err = core.UnmarshalModel(m, "result", &obj.Result, UnmarshalMaxUploadRespResult)
	if err != nil {
		err = core.SDKErrorf(err, "", "result-error", common.GetComponentInfo())
		return
	}
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
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// MinTlsVersionResp : Minimum TLS Version response.
type MinTlsVersionResp struct {
	// Container for response information.
	Result *MinTlsVersionRespResult `json:"result" validate:"required"`

	// Was the get successful.
	Success *bool `json:"success" validate:"required"`

	// Array of errors encountered.
	Errors [][]string `json:"errors" validate:"required"`

	// Array of messages returned.
	Messages [][]string `json:"messages" validate:"required"`
}

// UnmarshalMinTlsVersionResp unmarshals an instance of MinTlsVersionResp from the specified map of raw messages.
func UnmarshalMinTlsVersionResp(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(MinTlsVersionResp)
	err = core.UnmarshalModel(m, "result", &obj.Result, UnmarshalMinTlsVersionRespResult)
	if err != nil {
		err = core.SDKErrorf(err, "", "result-error", common.GetComponentInfo())
		return
	}
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
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// MinifyResp : Minify response.
type MinifyResp struct {
	// Container for response information.
	Result *MinifyRespResult `json:"result" validate:"required"`

	// Was the get successful.
	Success *bool `json:"success" validate:"required"`

	// Array of errors encountered.
	Errors [][]string `json:"errors" validate:"required"`

	// Array of messages returned.
	Messages [][]string `json:"messages" validate:"required"`
}

// UnmarshalMinifyResp unmarshals an instance of MinifyResp from the specified map of raw messages.
func UnmarshalMinifyResp(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(MinifyResp)
	err = core.UnmarshalModel(m, "result", &obj.Result, UnmarshalMinifyRespResult)
	if err != nil {
		err = core.SDKErrorf(err, "", "result-error", common.GetComponentInfo())
		return
	}
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
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// MobileRedirectResp : Mobile Redirect Response.
type MobileRedirectResp struct {
	// Container for response information.
	Result *MobileRedirectRespResult `json:"result" validate:"required"`

	// Was the get successful.
	Success *bool `json:"success" validate:"required"`

	// Array of errors encountered.
	Errors [][]string `json:"errors" validate:"required"`

	// Array of messages returned.
	Messages [][]string `json:"messages" validate:"required"`
}

// UnmarshalMobileRedirectResp unmarshals an instance of MobileRedirectResp from the specified map of raw messages.
func UnmarshalMobileRedirectResp(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(MobileRedirectResp)
	err = core.UnmarshalModel(m, "result", &obj.Result, UnmarshalMobileRedirectRespResult)
	if err != nil {
		err = core.SDKErrorf(err, "", "result-error", common.GetComponentInfo())
		return
	}
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
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// OpportunisticEncryptionResp : Opportunistic encryption response.
type OpportunisticEncryptionResp struct {
	// Container for response information.
	Result *OpportunisticEncryptionRespResult `json:"result" validate:"required"`

	// Was the get successful.
	Success *bool `json:"success" validate:"required"`

	// Array of errors encountered.
	Errors [][]string `json:"errors" validate:"required"`

	// Array of messages returned.
	Messages [][]string `json:"messages" validate:"required"`
}

// UnmarshalOpportunisticEncryptionResp unmarshals an instance of OpportunisticEncryptionResp from the specified map of raw messages.
func UnmarshalOpportunisticEncryptionResp(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(OpportunisticEncryptionResp)
	err = core.UnmarshalModel(m, "result", &obj.Result, UnmarshalOpportunisticEncryptionRespResult)
	if err != nil {
		err = core.SDKErrorf(err, "", "result-error", common.GetComponentInfo())
		return
	}
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
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// OpportunisticOnionResp : Opportunistic onion response.
type OpportunisticOnionResp struct {
	// Container for response information.
	Result *OpportunisticOnionRespResult `json:"result" validate:"required"`

	// Was the get successful.
	Success *bool `json:"success" validate:"required"`

	// Array of errors encountered.
	Errors [][]string `json:"errors" validate:"required"`

	// Array of messages returned.
	Messages [][]string `json:"messages" validate:"required"`
}

// UnmarshalOpportunisticOnionResp unmarshals an instance of OpportunisticOnionResp from the specified map of raw messages.
func UnmarshalOpportunisticOnionResp(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(OpportunisticOnionResp)
	err = core.UnmarshalModel(m, "result", &obj.Result, UnmarshalOpportunisticOnionRespResult)
	if err != nil {
		err = core.SDKErrorf(err, "", "result-error", common.GetComponentInfo())
		return
	}
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
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// OriginErrorPagePassThruResp : origin error page pass through response.
type OriginErrorPagePassThruResp struct {
	// Container for response information.
	Result *OriginErrorPagePassThruRespResult `json:"result" validate:"required"`

	// Was the get successful.
	Success *bool `json:"success" validate:"required"`

	// Array of errors encountered.
	Errors [][]string `json:"errors" validate:"required"`

	// Array of messages returned.
	Messages [][]string `json:"messages" validate:"required"`
}

// UnmarshalOriginErrorPagePassThruResp unmarshals an instance of OriginErrorPagePassThruResp from the specified map of raw messages.
func UnmarshalOriginErrorPagePassThruResp(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(OriginErrorPagePassThruResp)
	err = core.UnmarshalModel(m, "result", &obj.Result, UnmarshalOriginErrorPagePassThruRespResult)
	if err != nil {
		err = core.SDKErrorf(err, "", "result-error", common.GetComponentInfo())
		return
	}
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
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// OriginMaxHttpVersionResp : Origin max http version response.
type OriginMaxHttpVersionResp struct {
	// Container for response information.
	Result *OriginMaxHttpVersionRespResult `json:"result" validate:"required"`

	// Was the get successful.
	Success *bool `json:"success" validate:"required"`

	// Array of errors encountered.
	Errors [][]string `json:"errors" validate:"required"`

	// Array of messages returned.
	Messages [][]string `json:"messages" validate:"required"`
}

// UnmarshalOriginMaxHttpVersionResp unmarshals an instance of OriginMaxHttpVersionResp from the specified map of raw messages.
func UnmarshalOriginMaxHttpVersionResp(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(OriginMaxHttpVersionResp)
	err = core.UnmarshalModel(m, "result", &obj.Result, UnmarshalOriginMaxHttpVersionRespResult)
	if err != nil {
		err = core.SDKErrorf(err, "", "result-error", common.GetComponentInfo())
		return
	}
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
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// OriginPostQuantumEncryptionResp : Origin post quantum encryption response.
type OriginPostQuantumEncryptionResp struct {
	// Container for response information.
	Result *OriginPostQuantumEncryptionRespResult `json:"result" validate:"required"`

	// Was the get successful.
	Success *bool `json:"success" validate:"required"`

	// Array of errors encountered.
	Errors [][]string `json:"errors" validate:"required"`

	// Array of messages returned.
	Messages [][]string `json:"messages" validate:"required"`
}

// UnmarshalOriginPostQuantumEncryptionResp unmarshals an instance of OriginPostQuantumEncryptionResp from the specified map of raw messages.
func UnmarshalOriginPostQuantumEncryptionResp(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(OriginPostQuantumEncryptionResp)
	err = core.UnmarshalModel(m, "result", &obj.Result, UnmarshalOriginPostQuantumEncryptionRespResult)
	if err != nil {
		err = core.SDKErrorf(err, "", "result-error", common.GetComponentInfo())
		return
	}
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
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// PrefetchPreloadResp : Prefetch & Preload Response.
type PrefetchPreloadResp struct {
	// Container for response information.
	Result *PrefetchPreloadRespResult `json:"result" validate:"required"`

	// Was the get successful.
	Success *bool `json:"success" validate:"required"`

	// Array of errors encountered.
	Errors [][]string `json:"errors" validate:"required"`

	// Array of messages returned.
	Messages [][]string `json:"messages" validate:"required"`
}

// UnmarshalPrefetchPreloadResp unmarshals an instance of PrefetchPreloadResp from the specified map of raw messages.
func UnmarshalPrefetchPreloadResp(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PrefetchPreloadResp)
	err = core.UnmarshalModel(m, "result", &obj.Result, UnmarshalPrefetchPreloadRespResult)
	if err != nil {
		err = core.SDKErrorf(err, "", "result-error", common.GetComponentInfo())
		return
	}
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
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ProxyReadTimeoutResp : Proxy read timeout response.
type ProxyReadTimeoutResp struct {
	// Container for response information.
	Result *ProxyReadTimeoutRespResult `json:"result" validate:"required"`

	// Was the get successful.
	Success *bool `json:"success" validate:"required"`

	// Array of errors encountered.
	Errors [][]string `json:"errors" validate:"required"`

	// Array of messages returned.
	Messages [][]string `json:"messages" validate:"required"`
}

// UnmarshalProxyReadTimeoutResp unmarshals an instance of ProxyReadTimeoutResp from the specified map of raw messages.
func UnmarshalProxyReadTimeoutResp(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ProxyReadTimeoutResp)
	err = core.UnmarshalModel(m, "result", &obj.Result, UnmarshalProxyReadTimeoutRespResult)
	if err != nil {
		err = core.SDKErrorf(err, "", "result-error", common.GetComponentInfo())
		return
	}
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
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// PseudoIpv4Resp : Pseudo ipv4 response.
type PseudoIpv4Resp struct {
	// Container for response information.
	Result *PseudoIpv4RespResult `json:"result" validate:"required"`

	// Was the get successful.
	Success *bool `json:"success" validate:"required"`

	// Array of errors encountered.
	Errors [][]string `json:"errors" validate:"required"`

	// Array of messages returned.
	Messages [][]string `json:"messages" validate:"required"`
}

// UnmarshalPseudoIpv4Resp unmarshals an instance of PseudoIpv4Resp from the specified map of raw messages.
func UnmarshalPseudoIpv4Resp(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PseudoIpv4Resp)
	err = core.UnmarshalModel(m, "result", &obj.Result, UnmarshalPseudoIpv4RespResult)
	if err != nil {
		err = core.SDKErrorf(err, "", "result-error", common.GetComponentInfo())
		return
	}
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
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ResponseBufferingResp : Buffering response.
type ResponseBufferingResp struct {
	// Container for response information.
	Result *ResponseBufferingRespResult `json:"result" validate:"required"`

	// Was the get successful.
	Success *bool `json:"success" validate:"required"`

	// Array of errors encountered.
	Errors [][]string `json:"errors" validate:"required"`

	// Array of messages returned.
	Messages [][]string `json:"messages" validate:"required"`
}

// UnmarshalResponseBufferingResp unmarshals an instance of ResponseBufferingResp from the specified map of raw messages.
func UnmarshalResponseBufferingResp(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ResponseBufferingResp)
	err = core.UnmarshalModel(m, "result", &obj.Result, UnmarshalResponseBufferingRespResult)
	if err != nil {
		err = core.SDKErrorf(err, "", "result-error", common.GetComponentInfo())
		return
	}
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
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ScriptLoadOptimizationResp : Script load optimization response.
type ScriptLoadOptimizationResp struct {
	// Container for response information.
	Result *ScriptLoadOptimizationRespResult `json:"result" validate:"required"`

	// Was the get successful.
	Success *bool `json:"success" validate:"required"`

	// Array of errors encountered.
	Errors [][]string `json:"errors" validate:"required"`

	// Array of messages returned.
	Messages [][]string `json:"messages" validate:"required"`
}

// UnmarshalScriptLoadOptimizationResp unmarshals an instance of ScriptLoadOptimizationResp from the specified map of raw messages.
func UnmarshalScriptLoadOptimizationResp(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ScriptLoadOptimizationResp)
	err = core.UnmarshalModel(m, "result", &obj.Result, UnmarshalScriptLoadOptimizationRespResult)
	if err != nil {
		err = core.SDKErrorf(err, "", "result-error", common.GetComponentInfo())
		return
	}
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
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SecurityHeaderResp : Response of Security Header.
type SecurityHeaderResp struct {
	// Container for response information.
	Result *SecurityHeaderRespResult `json:"result" validate:"required"`

	// Was the get successful.
	Success *bool `json:"success" validate:"required"`

	// Array of errors encountered.
	Errors [][]string `json:"errors" validate:"required"`

	// Array of messages returned.
	Messages [][]string `json:"messages" validate:"required"`
}

// UnmarshalSecurityHeaderResp unmarshals an instance of SecurityHeaderResp from the specified map of raw messages.
func UnmarshalSecurityHeaderResp(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SecurityHeaderResp)
	err = core.UnmarshalModel(m, "result", &obj.Result, UnmarshalSecurityHeaderRespResult)
	if err != nil {
		err = core.SDKErrorf(err, "", "result-error", common.GetComponentInfo())
		return
	}
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
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ServerSideExcludeResp : Response of server side exclude.
type ServerSideExcludeResp struct {
	// Container for response information.
	Result *ServerSideExcludeRespResult `json:"result" validate:"required"`

	// Was the get successful.
	Success *bool `json:"success" validate:"required"`

	// Array of errors encountered.
	Errors [][]string `json:"errors" validate:"required"`

	// Array of messages returned.
	Messages [][]string `json:"messages" validate:"required"`
}

// UnmarshalServerSideExcludeResp unmarshals an instance of ServerSideExcludeResp from the specified map of raw messages.
func UnmarshalServerSideExcludeResp(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ServerSideExcludeResp)
	err = core.UnmarshalModel(m, "result", &obj.Result, UnmarshalServerSideExcludeRespResult)
	if err != nil {
		err = core.SDKErrorf(err, "", "result-error", common.GetComponentInfo())
		return
	}
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
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// TlsClientAuthResp : TLS Client authentication response.
type TlsClientAuthResp struct {
	// Container for response information.
	Result *TlsClientAuthRespResult `json:"result" validate:"required"`

	// Was the get successful.
	Success *bool `json:"success" validate:"required"`

	// Array of errors encountered.
	Errors [][]string `json:"errors" validate:"required"`

	// Array of messages returned.
	Messages [][]string `json:"messages" validate:"required"`
}

// UnmarshalTlsClientAuthResp unmarshals an instance of TlsClientAuthResp from the specified map of raw messages.
func UnmarshalTlsClientAuthResp(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(TlsClientAuthResp)
	err = core.UnmarshalModel(m, "result", &obj.Result, UnmarshalTlsClientAuthRespResult)
	if err != nil {
		err = core.SDKErrorf(err, "", "result-error", common.GetComponentInfo())
		return
	}
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
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// TrueClientIpResp : true client IP response.
type TrueClientIpResp struct {
	// Container for response information.
	Result *TrueClientIpRespResult `json:"result" validate:"required"`

	// Was the get successful.
	Success *bool `json:"success" validate:"required"`

	// Array of errors encountered.
	Errors [][]string `json:"errors" validate:"required"`

	// Array of messages returned.
	Messages [][]string `json:"messages" validate:"required"`
}

// UnmarshalTrueClientIpResp unmarshals an instance of TrueClientIpResp from the specified map of raw messages.
func UnmarshalTrueClientIpResp(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(TrueClientIpResp)
	err = core.UnmarshalModel(m, "result", &obj.Result, UnmarshalTrueClientIpRespResult)
	if err != nil {
		err = core.SDKErrorf(err, "", "result-error", common.GetComponentInfo())
		return
	}
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
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// WafResp : WAF Response.
type WafResp struct {
	// Container for response information.
	Result *WafRespResult `json:"result" validate:"required"`

	// Was the get successful.
	Success *bool `json:"success" validate:"required"`

	// Array of errors encountered.
	Errors [][]string `json:"errors" validate:"required"`

	// Array of messages returned.
	Messages [][]string `json:"messages" validate:"required"`
}

// UnmarshalWafResp unmarshals an instance of WafResp from the specified map of raw messages.
func UnmarshalWafResp(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(WafResp)
	err = core.UnmarshalModel(m, "result", &obj.Result, UnmarshalWafRespResult)
	if err != nil {
		err = core.SDKErrorf(err, "", "result-error", common.GetComponentInfo())
		return
	}
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
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// WebsocketsResp : Websocket Response.
type WebsocketsResp struct {
	// Container for response information.
	Result *WebsocketsRespResult `json:"result" validate:"required"`

	// Was the get successful.
	Success *bool `json:"success" validate:"required"`

	// Array of errors encountered.
	Errors [][]string `json:"errors" validate:"required"`

	// Array of messages returned.
	Messages [][]string `json:"messages" validate:"required"`
}

// UnmarshalWebsocketsResp unmarshals an instance of WebsocketsResp from the specified map of raw messages.
func UnmarshalWebsocketsResp(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(WebsocketsResp)
	err = core.UnmarshalModel(m, "result", &obj.Result, UnmarshalWebsocketsRespResult)
	if err != nil {
		err = core.SDKErrorf(err, "", "result-error", common.GetComponentInfo())
		return
	}
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
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ZonesCnameFlatteningResp : Zones CNAME flattening response.
type ZonesCnameFlatteningResp struct {
	// Was operation successful.
	Success *bool `json:"success" validate:"required"`

	// Array of errors encountered.
	Errors [][]string `json:"errors" validate:"required"`

	// Array of messages returned.
	Messages [][]string `json:"messages" validate:"required"`

	// CNAME Flattening response.
	Result *CnameFlatteningResponse `json:"result" validate:"required"`
}

// UnmarshalZonesCnameFlatteningResp unmarshals an instance of ZonesCnameFlatteningResp from the specified map of raw messages.
func UnmarshalZonesCnameFlatteningResp(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ZonesCnameFlatteningResp)
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
	err = core.UnmarshalModel(m, "result", &obj.Result, UnmarshalCnameFlatteningResponse)
	if err != nil {
		err = core.SDKErrorf(err, "", "result-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ZonesDnssecResp : Zones DNS Sec Response.
type ZonesDnssecResp struct {
	// Was operation successful.
	Success *bool `json:"success" validate:"required"`

	// Array of errors encountered.
	Errors [][]string `json:"errors" validate:"required"`

	// Array of messages returned.
	Messages [][]string `json:"messages" validate:"required"`

	// Container for response information.
	Result *ZonesDnssecRespResult `json:"result" validate:"required"`
}

// UnmarshalZonesDnssecResp unmarshals an instance of ZonesDnssecResp from the specified map of raw messages.
func UnmarshalZonesDnssecResp(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ZonesDnssecResp)
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
	err = core.UnmarshalModel(m, "result", &obj.Result, UnmarshalZonesDnssecRespResult)
	if err != nil {
		err = core.SDKErrorf(err, "", "result-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}
