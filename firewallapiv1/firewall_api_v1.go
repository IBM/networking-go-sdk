/**
 * (C) Copyright IBM Corp. 2020.
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
 * IBM OpenAPI SDK Code Generator Version: 3.20.0-debb9f29-20201203-202043
 */

// Package firewallapiv1 : Operations and models for the FirewallApiV1 service
package firewallapiv1

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	common "github.com/IBM/networking-go-sdk/common"
)

// FirewallApiV1 : Firewall API
//
// Version: 1.0.0
type FirewallApiV1 struct {
	Service *core.BaseService

	// cloud resource name.
	Crn *string

	// zone identifier.
	ZoneIdentifier *string
}

// DefaultServiceURL is the default URL to make service requests to.
const DefaultServiceURL = "https://api.cis.cloud.ibm.com"

// DefaultServiceName is the default key used to find external configuration information.
const DefaultServiceName = "firewall_api"

// FirewallApiV1Options : Service options
type FirewallApiV1Options struct {
	ServiceName   string
	URL           string
	Authenticator core.Authenticator

	// cloud resource name.
	Crn *string `validate:"required"`

	// zone identifier.
	ZoneIdentifier *string `validate:"required"`
}

// NewFirewallApiV1UsingExternalConfig : constructs an instance of FirewallApiV1 with passed in options and external configuration.
func NewFirewallApiV1UsingExternalConfig(options *FirewallApiV1Options) (firewallApi *FirewallApiV1, err error) {
	if options.ServiceName == "" {
		options.ServiceName = DefaultServiceName
	}

	if options.Authenticator == nil {
		options.Authenticator, err = core.GetAuthenticatorFromEnvironment(options.ServiceName)
		if err != nil {
			return
		}
	}

	firewallApi, err = NewFirewallApiV1(options)
	if err != nil {
		return
	}

	err = firewallApi.Service.ConfigureService(options.ServiceName)
	if err != nil {
		return
	}

	if options.URL != "" {
		err = firewallApi.Service.SetServiceURL(options.URL)
	}
	return
}

// NewFirewallApiV1 : constructs an instance of FirewallApiV1 with passed in options.
func NewFirewallApiV1(options *FirewallApiV1Options) (service *FirewallApiV1, err error) {
	serviceOptions := &core.ServiceOptions{
		URL:           DefaultServiceURL,
		Authenticator: options.Authenticator,
	}

	err = core.ValidateStruct(options, "options")
	if err != nil {
		return
	}

	baseService, err := core.NewBaseService(serviceOptions)
	if err != nil {
		return
	}

	if options.URL != "" {
		err = baseService.SetServiceURL(options.URL)
		if err != nil {
			return
		}
	}

	service = &FirewallApiV1{
		Service:        baseService,
		Crn:            options.Crn,
		ZoneIdentifier: options.ZoneIdentifier,
	}

	return
}

// GetServiceURLForRegion returns the service URL to be used for the specified region
func GetServiceURLForRegion(region string) (string, error) {
	return "", fmt.Errorf("service does not support regional URLs")
}

// Clone makes a copy of "firewallApi" suitable for processing requests.
func (firewallApi *FirewallApiV1) Clone() *FirewallApiV1 {
	if core.IsNil(firewallApi) {
		return nil
	}
	clone := *firewallApi
	clone.Service = firewallApi.Service.Clone()
	return &clone
}

// SetServiceURL sets the service URL
func (firewallApi *FirewallApiV1) SetServiceURL(url string) error {
	return firewallApi.Service.SetServiceURL(url)
}

// GetServiceURL returns the service URL
func (firewallApi *FirewallApiV1) GetServiceURL() string {
	return firewallApi.Service.GetServiceURL()
}

// SetDefaultHeaders sets HTTP headers to be sent in every request
func (firewallApi *FirewallApiV1) SetDefaultHeaders(headers http.Header) {
	firewallApi.Service.SetDefaultHeaders(headers)
}

// SetEnableGzipCompression sets the service's EnableGzipCompression field
func (firewallApi *FirewallApiV1) SetEnableGzipCompression(enableGzip bool) {
	firewallApi.Service.SetEnableGzipCompression(enableGzip)
}

// GetEnableGzipCompression returns the service's EnableGzipCompression field
func (firewallApi *FirewallApiV1) GetEnableGzipCompression() bool {
	return firewallApi.Service.GetEnableGzipCompression()
}

// EnableRetries enables automatic retries for requests invoked for this service instance.
// If either parameter is specified as 0, then a default value is used instead.
func (firewallApi *FirewallApiV1) EnableRetries(maxRetries int, maxRetryInterval time.Duration) {
	firewallApi.Service.EnableRetries(maxRetries, maxRetryInterval)
}

// DisableRetries disables automatic retries for requests invoked for this service instance.
func (firewallApi *FirewallApiV1) DisableRetries() {
	firewallApi.Service.DisableRetries()
}

// GetSecurityLevelSetting : Get security level setting
// For a given zone identifier, get security level setting.
func (firewallApi *FirewallApiV1) GetSecurityLevelSetting(getSecurityLevelSettingOptions *GetSecurityLevelSettingOptions) (result *SecurityLevelSettingResp, response *core.DetailedResponse, err error) {
	return firewallApi.GetSecurityLevelSettingWithContext(context.Background(), getSecurityLevelSettingOptions)
}

// GetSecurityLevelSettingWithContext is an alternate form of the GetSecurityLevelSetting method which supports a Context parameter
func (firewallApi *FirewallApiV1) GetSecurityLevelSettingWithContext(ctx context.Context, getSecurityLevelSettingOptions *GetSecurityLevelSettingOptions) (result *SecurityLevelSettingResp, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getSecurityLevelSettingOptions, "getSecurityLevelSettingOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"crn":             *firewallApi.Crn,
		"zone_identifier": *firewallApi.ZoneIdentifier,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = firewallApi.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(firewallApi.Service.Options.URL, `/v1/{crn}/zones/{zone_identifier}/settings/security_level`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getSecurityLevelSettingOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("firewall_api", "V1", "GetSecurityLevelSetting")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = firewallApi.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalSecurityLevelSettingResp)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// SetSecurityLevelSetting : Set security level setting
// For a given zone identifier, set security level setting.
func (firewallApi *FirewallApiV1) SetSecurityLevelSetting(setSecurityLevelSettingOptions *SetSecurityLevelSettingOptions) (result *SecurityLevelSettingResp, response *core.DetailedResponse, err error) {
	return firewallApi.SetSecurityLevelSettingWithContext(context.Background(), setSecurityLevelSettingOptions)
}

// SetSecurityLevelSettingWithContext is an alternate form of the SetSecurityLevelSetting method which supports a Context parameter
func (firewallApi *FirewallApiV1) SetSecurityLevelSettingWithContext(ctx context.Context, setSecurityLevelSettingOptions *SetSecurityLevelSettingOptions) (result *SecurityLevelSettingResp, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(setSecurityLevelSettingOptions, "setSecurityLevelSettingOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"crn":             *firewallApi.Crn,
		"zone_identifier": *firewallApi.ZoneIdentifier,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = firewallApi.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(firewallApi.Service.Options.URL, `/v1/{crn}/zones/{zone_identifier}/settings/security_level`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range setSecurityLevelSettingOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("firewall_api", "V1", "SetSecurityLevelSetting")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if setSecurityLevelSettingOptions.Value != nil {
		body["value"] = setSecurityLevelSettingOptions.Value
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = firewallApi.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalSecurityLevelSettingResp)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// GetSecurityLevelSettingOptions : The GetSecurityLevelSetting options.
type GetSecurityLevelSettingOptions struct {

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetSecurityLevelSettingOptions : Instantiate GetSecurityLevelSettingOptions
func (*FirewallApiV1) NewGetSecurityLevelSettingOptions() *GetSecurityLevelSettingOptions {
	return &GetSecurityLevelSettingOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *GetSecurityLevelSettingOptions) SetHeaders(param map[string]string) *GetSecurityLevelSettingOptions {
	options.Headers = param
	return options
}

// SecurityLevelSettingRespMessagesItem : SecurityLevelSettingRespMessagesItem struct
type SecurityLevelSettingRespMessagesItem struct {
	// messages.
	Status *string `json:"status,omitempty"`
}

// UnmarshalSecurityLevelSettingRespMessagesItem unmarshals an instance of SecurityLevelSettingRespMessagesItem from the specified map of raw messages.
func UnmarshalSecurityLevelSettingRespMessagesItem(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SecurityLevelSettingRespMessagesItem)
	err = core.UnmarshalPrimitive(m, "status", &obj.Status)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SecurityLevelSettingRespResult : result object.
type SecurityLevelSettingRespResult struct {
	// identifier.
	ID *string `json:"id" validate:"required"`

	// value.
	Value *string `json:"value" validate:"required"`

	// editable.
	Editable *bool `json:"editable" validate:"required"`

	// modified date.
	ModifiedOn *string `json:"modified_on" validate:"required"`
}

// Constants associated with the SecurityLevelSettingRespResult.ID property.
// identifier.
const (
	SecurityLevelSettingRespResult_ID_SecurityLevel = "security_level"
)

// UnmarshalSecurityLevelSettingRespResult unmarshals an instance of SecurityLevelSettingRespResult from the specified map of raw messages.
func UnmarshalSecurityLevelSettingRespResult(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SecurityLevelSettingRespResult)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "value", &obj.Value)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "editable", &obj.Editable)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "modified_on", &obj.ModifiedOn)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SetSecurityLevelSettingOptions : The SetSecurityLevelSetting options.
type SetSecurityLevelSettingOptions struct {
	// security level.
	Value *string `json:"value,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the SetSecurityLevelSettingOptions.Value property.
// security level.
const (
	SetSecurityLevelSettingOptions_Value_EssentiallyOff = "essentially_off"
	SetSecurityLevelSettingOptions_Value_High           = "high"
	SetSecurityLevelSettingOptions_Value_Low            = "low"
	SetSecurityLevelSettingOptions_Value_Medium         = "medium"
	SetSecurityLevelSettingOptions_Value_UnderAttack    = "under_attack"
)

// NewSetSecurityLevelSettingOptions : Instantiate SetSecurityLevelSettingOptions
func (*FirewallApiV1) NewSetSecurityLevelSettingOptions() *SetSecurityLevelSettingOptions {
	return &SetSecurityLevelSettingOptions{}
}

// SetValue : Allow user to set Value
func (options *SetSecurityLevelSettingOptions) SetValue(value string) *SetSecurityLevelSettingOptions {
	options.Value = core.StringPtr(value)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *SetSecurityLevelSettingOptions) SetHeaders(param map[string]string) *SetSecurityLevelSettingOptions {
	options.Headers = param
	return options
}

// ResultInfo : result information.
type ResultInfo struct {
	// output pages.
	Page *int64 `json:"page" validate:"required"`

	// output per page.
	PerPage *int64 `json:"per_page" validate:"required"`

	// firewall hit count.
	Count *int64 `json:"count" validate:"required"`

	// total count.
	TotalCount *int64 `json:"total_count" validate:"required"`
}

// UnmarshalResultInfo unmarshals an instance of ResultInfo from the specified map of raw messages.
func UnmarshalResultInfo(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ResultInfo)
	err = core.UnmarshalPrimitive(m, "page", &obj.Page)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "per_page", &obj.PerPage)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "count", &obj.Count)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "total_count", &obj.TotalCount)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SecurityLevelSettingResp : security level setting response.
type SecurityLevelSettingResp struct {
	// result object.
	Result *SecurityLevelSettingRespResult `json:"result" validate:"required"`

	// result information.
	ResultInfo *ResultInfo `json:"result_info" validate:"required"`

	// success response.
	Success *bool `json:"success" validate:"required"`

	// array of errors.
	Errors [][]string `json:"errors" validate:"required"`

	// array of messages.
	Messages []SecurityLevelSettingRespMessagesItem `json:"messages" validate:"required"`
}

// UnmarshalSecurityLevelSettingResp unmarshals an instance of SecurityLevelSettingResp from the specified map of raw messages.
func UnmarshalSecurityLevelSettingResp(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SecurityLevelSettingResp)
	err = core.UnmarshalModel(m, "result", &obj.Result, UnmarshalSecurityLevelSettingRespResult)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "result_info", &obj.ResultInfo, UnmarshalResultInfo)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "success", &obj.Success)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "errors", &obj.Errors)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "messages", &obj.Messages, UnmarshalSecurityLevelSettingRespMessagesItem)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}
