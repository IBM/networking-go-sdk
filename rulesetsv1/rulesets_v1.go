/**
 * (C) Copyright IBM Corp. 2024.
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
 * IBM OpenAPI SDK Code Generator Version: 3.85.0-75c38f8f-20240206-210220
 */

// Package rulesetsv1 : Operations and models for the RulesetsV1 service
package rulesetsv1

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

// RulesetsV1 : Rulesets Engine
//
// API Version: 1.0.1
type RulesetsV1 struct {
	Service *core.BaseService

	// Full url-encoded CRN of the service instance.
	Crn *string

	// zone identifier.
	ZoneIdentifier *string
}

// DefaultServiceURL is the default URL to make service requests to.
const DefaultServiceURL = "https://api.cis.cloud.ibm.com"

// DefaultServiceName is the default key used to find external configuration information.
const DefaultServiceName = "rulesets"

// RulesetsV1Options : Service options
type RulesetsV1Options struct {
	ServiceName   string
	URL           string
	Authenticator core.Authenticator

	// Full url-encoded CRN of the service instance.
	Crn *string `validate:"required"`

	// zone identifier.
	ZoneIdentifier *string `validate:"required"`
}

// NewRulesetsV1UsingExternalConfig : constructs an instance of RulesetsV1 with passed in options and external configuration.
func NewRulesetsV1UsingExternalConfig(options *RulesetsV1Options) (rulesets *RulesetsV1, err error) {
	if options.ServiceName == "" {
		options.ServiceName = DefaultServiceName
	}

	if options.Authenticator == nil {
		options.Authenticator, err = core.GetAuthenticatorFromEnvironment(options.ServiceName)
		if err != nil {
			return
		}
	}

	rulesets, err = NewRulesetsV1(options)
	if err != nil {
		return
	}

	err = rulesets.Service.ConfigureService(options.ServiceName)
	if err != nil {
		return
	}

	if options.URL != "" {
		err = rulesets.Service.SetServiceURL(options.URL)
	}
	return
}

// NewRulesetsV1 : constructs an instance of RulesetsV1 with passed in options.
func NewRulesetsV1(options *RulesetsV1Options) (service *RulesetsV1, err error) {
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

	service = &RulesetsV1{
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

// Clone makes a copy of "rulesets" suitable for processing requests.
func (rulesets *RulesetsV1) Clone() *RulesetsV1 {
	if core.IsNil(rulesets) {
		return nil
	}
	clone := *rulesets
	clone.Service = rulesets.Service.Clone()
	return &clone
}

// SetServiceURL sets the service URL
func (rulesets *RulesetsV1) SetServiceURL(url string) error {
	return rulesets.Service.SetServiceURL(url)
}

// GetServiceURL returns the service URL
func (rulesets *RulesetsV1) GetServiceURL() string {
	return rulesets.Service.GetServiceURL()
}

// SetDefaultHeaders sets HTTP headers to be sent in every request
func (rulesets *RulesetsV1) SetDefaultHeaders(headers http.Header) {
	rulesets.Service.SetDefaultHeaders(headers)
}

// SetEnableGzipCompression sets the service's EnableGzipCompression field
func (rulesets *RulesetsV1) SetEnableGzipCompression(enableGzip bool) {
	rulesets.Service.SetEnableGzipCompression(enableGzip)
}

// GetEnableGzipCompression returns the service's EnableGzipCompression field
func (rulesets *RulesetsV1) GetEnableGzipCompression() bool {
	return rulesets.Service.GetEnableGzipCompression()
}

// EnableRetries enables automatic retries for requests invoked for this service instance.
// If either parameter is specified as 0, then a default value is used instead.
func (rulesets *RulesetsV1) EnableRetries(maxRetries int, maxRetryInterval time.Duration) {
	rulesets.Service.EnableRetries(maxRetries, maxRetryInterval)
}

// DisableRetries disables automatic retries for requests invoked for this service instance.
func (rulesets *RulesetsV1) DisableRetries() {
	rulesets.Service.DisableRetries()
}

// GetAccountRulesets : List account rulesets
// List all rulesets at the account level.
func (rulesets *RulesetsV1) GetAccountRulesets(getAccountRulesetsOptions *GetAccountRulesetsOptions) (result *ListRulesetsResp, response *core.DetailedResponse, err error) {
	return rulesets.GetAccountRulesetsWithContext(context.Background(), getAccountRulesetsOptions)
}

// GetAccountRulesetsWithContext is an alternate form of the GetAccountRulesets method which supports a Context parameter
func (rulesets *RulesetsV1) GetAccountRulesetsWithContext(ctx context.Context, getAccountRulesetsOptions *GetAccountRulesetsOptions) (result *ListRulesetsResp, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getAccountRulesetsOptions, "getAccountRulesetsOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"crn": *rulesets.Crn,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = rulesets.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(rulesets.Service.Options.URL, `/v1/{crn}/rulesets`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getAccountRulesetsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("rulesets", "V1", "GetAccountRulesets")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = rulesets.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalListRulesetsResp)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetAccountRuleset : Get an account ruleset
// View a specific account ruleset.
func (rulesets *RulesetsV1) GetAccountRuleset(getAccountRulesetOptions *GetAccountRulesetOptions) (result *RulesetResp, response *core.DetailedResponse, err error) {
	return rulesets.GetAccountRulesetWithContext(context.Background(), getAccountRulesetOptions)
}

// GetAccountRulesetWithContext is an alternate form of the GetAccountRuleset method which supports a Context parameter
func (rulesets *RulesetsV1) GetAccountRulesetWithContext(ctx context.Context, getAccountRulesetOptions *GetAccountRulesetOptions) (result *RulesetResp, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getAccountRulesetOptions, "getAccountRulesetOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getAccountRulesetOptions, "getAccountRulesetOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"crn":        *rulesets.Crn,
		"ruleset_id": *getAccountRulesetOptions.RulesetID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = rulesets.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(rulesets.Service.Options.URL, `/v1/{crn}/rulesets/{ruleset_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getAccountRulesetOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("rulesets", "V1", "GetAccountRuleset")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = rulesets.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalRulesetResp)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// UpdateAccountRuleset : Update an account ruleset
// Update a specific account ruleset.
func (rulesets *RulesetsV1) UpdateAccountRuleset(updateAccountRulesetOptions *UpdateAccountRulesetOptions) (result *RulesetResp, response *core.DetailedResponse, err error) {
	return rulesets.UpdateAccountRulesetWithContext(context.Background(), updateAccountRulesetOptions)
}

// UpdateAccountRulesetWithContext is an alternate form of the UpdateAccountRuleset method which supports a Context parameter
func (rulesets *RulesetsV1) UpdateAccountRulesetWithContext(ctx context.Context, updateAccountRulesetOptions *UpdateAccountRulesetOptions) (result *RulesetResp, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateAccountRulesetOptions, "updateAccountRulesetOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateAccountRulesetOptions, "updateAccountRulesetOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"crn":        *rulesets.Crn,
		"ruleset_id": *updateAccountRulesetOptions.RulesetID,
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = rulesets.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(rulesets.Service.Options.URL, `/v1/{crn}/rulesets/{ruleset_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateAccountRulesetOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("rulesets", "V1", "UpdateAccountRuleset")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if updateAccountRulesetOptions.Description != nil {
		body["description"] = updateAccountRulesetOptions.Description
	}
	if updateAccountRulesetOptions.Kind != nil {
		body["kind"] = updateAccountRulesetOptions.Kind
	}
	if updateAccountRulesetOptions.Name != nil {
		body["name"] = updateAccountRulesetOptions.Name
	}
	if updateAccountRulesetOptions.Phase != nil {
		body["phase"] = updateAccountRulesetOptions.Phase
	}
	if updateAccountRulesetOptions.Rules != nil {
		body["rules"] = updateAccountRulesetOptions.Rules
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
	response, err = rulesets.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalRulesetResp)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeleteAccountRuleset : Delete an account ruleset
// Delete a specific account ruleset.
func (rulesets *RulesetsV1) DeleteAccountRuleset(deleteAccountRulesetOptions *DeleteAccountRulesetOptions) (response *core.DetailedResponse, err error) {
	return rulesets.DeleteAccountRulesetWithContext(context.Background(), deleteAccountRulesetOptions)
}

// DeleteAccountRulesetWithContext is an alternate form of the DeleteAccountRuleset method which supports a Context parameter
func (rulesets *RulesetsV1) DeleteAccountRulesetWithContext(ctx context.Context, deleteAccountRulesetOptions *DeleteAccountRulesetOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteAccountRulesetOptions, "deleteAccountRulesetOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteAccountRulesetOptions, "deleteAccountRulesetOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"crn":        *rulesets.Crn,
		"ruleset_id": *deleteAccountRulesetOptions.RulesetID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = rulesets.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(rulesets.Service.Options.URL, `/v1/{crn}/rulesets/{ruleset_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteAccountRulesetOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("rulesets", "V1", "DeleteAccountRuleset")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = rulesets.Service.Request(request, nil)

	return
}

// GetAccountRulesetVersions : List version of an account ruleset
// List all versions of a specific account ruleset.
func (rulesets *RulesetsV1) GetAccountRulesetVersions(getAccountRulesetVersionsOptions *GetAccountRulesetVersionsOptions) (result *ListRulesetsResp, response *core.DetailedResponse, err error) {
	return rulesets.GetAccountRulesetVersionsWithContext(context.Background(), getAccountRulesetVersionsOptions)
}

// GetAccountRulesetVersionsWithContext is an alternate form of the GetAccountRulesetVersions method which supports a Context parameter
func (rulesets *RulesetsV1) GetAccountRulesetVersionsWithContext(ctx context.Context, getAccountRulesetVersionsOptions *GetAccountRulesetVersionsOptions) (result *ListRulesetsResp, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getAccountRulesetVersionsOptions, "getAccountRulesetVersionsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getAccountRulesetVersionsOptions, "getAccountRulesetVersionsOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"crn":        *rulesets.Crn,
		"ruleset_id": *getAccountRulesetVersionsOptions.RulesetID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = rulesets.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(rulesets.Service.Options.URL, `/v1/{crn}/rulesets/{ruleset_id}/versions`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getAccountRulesetVersionsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("rulesets", "V1", "GetAccountRulesetVersions")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = rulesets.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalListRulesetsResp)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetAccountRulesetVersion : Get a specific version of an account ruleset
// View a specific version of a specific account ruleset.
func (rulesets *RulesetsV1) GetAccountRulesetVersion(getAccountRulesetVersionOptions *GetAccountRulesetVersionOptions) (result *RulesetResp, response *core.DetailedResponse, err error) {
	return rulesets.GetAccountRulesetVersionWithContext(context.Background(), getAccountRulesetVersionOptions)
}

// GetAccountRulesetVersionWithContext is an alternate form of the GetAccountRulesetVersion method which supports a Context parameter
func (rulesets *RulesetsV1) GetAccountRulesetVersionWithContext(ctx context.Context, getAccountRulesetVersionOptions *GetAccountRulesetVersionOptions) (result *RulesetResp, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getAccountRulesetVersionOptions, "getAccountRulesetVersionOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getAccountRulesetVersionOptions, "getAccountRulesetVersionOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"crn":             *rulesets.Crn,
		"ruleset_id":      *getAccountRulesetVersionOptions.RulesetID,
		"ruleset_version": *getAccountRulesetVersionOptions.RulesetVersion,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = rulesets.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(rulesets.Service.Options.URL, `/v1/{crn}/rulesets/{ruleset_id}/versions/{ruleset_version}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getAccountRulesetVersionOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("rulesets", "V1", "GetAccountRulesetVersion")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = rulesets.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalRulesetResp)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeleteAccountRulesetVersion : Delete a specific version of an account ruleset
// Delete a specific version of a specific account ruleset.
func (rulesets *RulesetsV1) DeleteAccountRulesetVersion(deleteAccountRulesetVersionOptions *DeleteAccountRulesetVersionOptions) (response *core.DetailedResponse, err error) {
	return rulesets.DeleteAccountRulesetVersionWithContext(context.Background(), deleteAccountRulesetVersionOptions)
}

// DeleteAccountRulesetVersionWithContext is an alternate form of the DeleteAccountRulesetVersion method which supports a Context parameter
func (rulesets *RulesetsV1) DeleteAccountRulesetVersionWithContext(ctx context.Context, deleteAccountRulesetVersionOptions *DeleteAccountRulesetVersionOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteAccountRulesetVersionOptions, "deleteAccountRulesetVersionOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteAccountRulesetVersionOptions, "deleteAccountRulesetVersionOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"crn":             *rulesets.Crn,
		"ruleset_id":      *deleteAccountRulesetVersionOptions.RulesetID,
		"ruleset_version": *deleteAccountRulesetVersionOptions.RulesetVersion,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = rulesets.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(rulesets.Service.Options.URL, `/v1/{crn}/rulesets/{ruleset_id}/versions/{ruleset_version}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteAccountRulesetVersionOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("rulesets", "V1", "DeleteAccountRulesetVersion")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = rulesets.Service.Request(request, nil)

	return
}

// GetAccountEntrypointRuleset : Get an account entrypoint ruleset
// Get the account ruleset for the given phase's entrypoint.
func (rulesets *RulesetsV1) GetAccountEntrypointRuleset(getAccountEntrypointRulesetOptions *GetAccountEntrypointRulesetOptions) (result *RulesetResp, response *core.DetailedResponse, err error) {
	return rulesets.GetAccountEntrypointRulesetWithContext(context.Background(), getAccountEntrypointRulesetOptions)
}

// GetAccountEntrypointRulesetWithContext is an alternate form of the GetAccountEntrypointRuleset method which supports a Context parameter
func (rulesets *RulesetsV1) GetAccountEntrypointRulesetWithContext(ctx context.Context, getAccountEntrypointRulesetOptions *GetAccountEntrypointRulesetOptions) (result *RulesetResp, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getAccountEntrypointRulesetOptions, "getAccountEntrypointRulesetOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getAccountEntrypointRulesetOptions, "getAccountEntrypointRulesetOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"crn":           *rulesets.Crn,
		"ruleset_phase": *getAccountEntrypointRulesetOptions.RulesetPhase,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = rulesets.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(rulesets.Service.Options.URL, `/v1/{crn}/rulesets/phases/{ruleset_phase}/entrypoint`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getAccountEntrypointRulesetOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("rulesets", "V1", "GetAccountEntrypointRuleset")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = rulesets.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalRulesetResp)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// UpdateAccountEntrypointRuleset : Update an account entrypoint ruleset
// Updates the account ruleset for the given phase's entry point.
func (rulesets *RulesetsV1) UpdateAccountEntrypointRuleset(updateAccountEntrypointRulesetOptions *UpdateAccountEntrypointRulesetOptions) (result *RulesetResp, response *core.DetailedResponse, err error) {
	return rulesets.UpdateAccountEntrypointRulesetWithContext(context.Background(), updateAccountEntrypointRulesetOptions)
}

// UpdateAccountEntrypointRulesetWithContext is an alternate form of the UpdateAccountEntrypointRuleset method which supports a Context parameter
func (rulesets *RulesetsV1) UpdateAccountEntrypointRulesetWithContext(ctx context.Context, updateAccountEntrypointRulesetOptions *UpdateAccountEntrypointRulesetOptions) (result *RulesetResp, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateAccountEntrypointRulesetOptions, "updateAccountEntrypointRulesetOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateAccountEntrypointRulesetOptions, "updateAccountEntrypointRulesetOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"crn":           *rulesets.Crn,
		"ruleset_phase": *updateAccountEntrypointRulesetOptions.RulesetPhase,
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = rulesets.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(rulesets.Service.Options.URL, `/v1/{crn}/rulesets/phases/{ruleset_phase}/entrypoint`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateAccountEntrypointRulesetOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("rulesets", "V1", "UpdateAccountEntrypointRuleset")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if updateAccountEntrypointRulesetOptions.Description != nil {
		body["description"] = updateAccountEntrypointRulesetOptions.Description
	}
	if updateAccountEntrypointRulesetOptions.Kind != nil {
		body["kind"] = updateAccountEntrypointRulesetOptions.Kind
	}
	if updateAccountEntrypointRulesetOptions.Name != nil {
		body["name"] = updateAccountEntrypointRulesetOptions.Name
	}
	if updateAccountEntrypointRulesetOptions.Phase != nil {
		body["phase"] = updateAccountEntrypointRulesetOptions.Phase
	}
	if updateAccountEntrypointRulesetOptions.Rules != nil {
		body["rules"] = updateAccountEntrypointRulesetOptions.Rules
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
	response, err = rulesets.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalRulesetResp)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetAccountEntryPointRulesetVersions : List an account entry point ruleset's versions
// Lists the account ruleset versions for the given phase's entry point.
func (rulesets *RulesetsV1) GetAccountEntryPointRulesetVersions(getAccountEntryPointRulesetVersionsOptions *GetAccountEntryPointRulesetVersionsOptions) (result *ListRulesetsResp, response *core.DetailedResponse, err error) {
	return rulesets.GetAccountEntryPointRulesetVersionsWithContext(context.Background(), getAccountEntryPointRulesetVersionsOptions)
}

// GetAccountEntryPointRulesetVersionsWithContext is an alternate form of the GetAccountEntryPointRulesetVersions method which supports a Context parameter
func (rulesets *RulesetsV1) GetAccountEntryPointRulesetVersionsWithContext(ctx context.Context, getAccountEntryPointRulesetVersionsOptions *GetAccountEntryPointRulesetVersionsOptions) (result *ListRulesetsResp, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getAccountEntryPointRulesetVersionsOptions, "getAccountEntryPointRulesetVersionsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getAccountEntryPointRulesetVersionsOptions, "getAccountEntryPointRulesetVersionsOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"crn":           *rulesets.Crn,
		"ruleset_phase": *getAccountEntryPointRulesetVersionsOptions.RulesetPhase,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = rulesets.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(rulesets.Service.Options.URL, `/v1/{crn}/rulesets/phases/{ruleset_phase}/entrypoint/versions`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getAccountEntryPointRulesetVersionsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("rulesets", "V1", "GetAccountEntryPointRulesetVersions")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = rulesets.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalListRulesetsResp)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetAccountEntryPointRulesetVersion : Get an account entry point ruleset version
// Fetches a specific version of an account entry point ruleset.
func (rulesets *RulesetsV1) GetAccountEntryPointRulesetVersion(getAccountEntryPointRulesetVersionOptions *GetAccountEntryPointRulesetVersionOptions) (result *RulesetResp, response *core.DetailedResponse, err error) {
	return rulesets.GetAccountEntryPointRulesetVersionWithContext(context.Background(), getAccountEntryPointRulesetVersionOptions)
}

// GetAccountEntryPointRulesetVersionWithContext is an alternate form of the GetAccountEntryPointRulesetVersion method which supports a Context parameter
func (rulesets *RulesetsV1) GetAccountEntryPointRulesetVersionWithContext(ctx context.Context, getAccountEntryPointRulesetVersionOptions *GetAccountEntryPointRulesetVersionOptions) (result *RulesetResp, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getAccountEntryPointRulesetVersionOptions, "getAccountEntryPointRulesetVersionOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getAccountEntryPointRulesetVersionOptions, "getAccountEntryPointRulesetVersionOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"crn":             *rulesets.Crn,
		"ruleset_phase":   *getAccountEntryPointRulesetVersionOptions.RulesetPhase,
		"ruleset_version": *getAccountEntryPointRulesetVersionOptions.RulesetVersion,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = rulesets.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(rulesets.Service.Options.URL, `/v1/{crn}/rulesets/phases/{ruleset_phase}/entrypoint/versions/{ruleset_version}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getAccountEntryPointRulesetVersionOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("rulesets", "V1", "GetAccountEntryPointRulesetVersion")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = rulesets.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalRulesetResp)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// CreateAccountRulesetRule : Create an account ruleset rule
// Create an account ruleset rule.
func (rulesets *RulesetsV1) CreateAccountRulesetRule(createAccountRulesetRuleOptions *CreateAccountRulesetRuleOptions) (result *RuleResp, response *core.DetailedResponse, err error) {
	return rulesets.CreateAccountRulesetRuleWithContext(context.Background(), createAccountRulesetRuleOptions)
}

// CreateAccountRulesetRuleWithContext is an alternate form of the CreateAccountRulesetRule method which supports a Context parameter
func (rulesets *RulesetsV1) CreateAccountRulesetRuleWithContext(ctx context.Context, createAccountRulesetRuleOptions *CreateAccountRulesetRuleOptions) (result *RuleResp, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createAccountRulesetRuleOptions, "createAccountRulesetRuleOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createAccountRulesetRuleOptions, "createAccountRulesetRuleOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"crn":        *rulesets.Crn,
		"ruleset_id": *createAccountRulesetRuleOptions.RulesetID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = rulesets.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(rulesets.Service.Options.URL, `/v1/{crn}/rulesets/{ruleset_id}/rules`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range createAccountRulesetRuleOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("rulesets", "V1", "CreateAccountRulesetRule")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if createAccountRulesetRuleOptions.Action != nil {
		body["action"] = createAccountRulesetRuleOptions.Action
	}
	if createAccountRulesetRuleOptions.ActionParameters != nil {
		body["action_parameters"] = createAccountRulesetRuleOptions.ActionParameters
	}
	if createAccountRulesetRuleOptions.Description != nil {
		body["description"] = createAccountRulesetRuleOptions.Description
	}
	if createAccountRulesetRuleOptions.Enabled != nil {
		body["enabled"] = createAccountRulesetRuleOptions.Enabled
	}
	if createAccountRulesetRuleOptions.Expression != nil {
		body["expression"] = createAccountRulesetRuleOptions.Expression
	}
	if createAccountRulesetRuleOptions.ID != nil {
		body["id"] = createAccountRulesetRuleOptions.ID
	}
	if createAccountRulesetRuleOptions.Logging != nil {
		body["logging"] = createAccountRulesetRuleOptions.Logging
	}
	if createAccountRulesetRuleOptions.Ref != nil {
		body["ref"] = createAccountRulesetRuleOptions.Ref
	}
	if createAccountRulesetRuleOptions.Position != nil {
		body["position"] = createAccountRulesetRuleOptions.Position
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
	response, err = rulesets.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalRuleResp)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// UpdateAccountRulesetRule : Update an account ruleset rule
// Update an account ruleset rule.
func (rulesets *RulesetsV1) UpdateAccountRulesetRule(updateAccountRulesetRuleOptions *UpdateAccountRulesetRuleOptions) (result *RuleResp, response *core.DetailedResponse, err error) {
	return rulesets.UpdateAccountRulesetRuleWithContext(context.Background(), updateAccountRulesetRuleOptions)
}

// UpdateAccountRulesetRuleWithContext is an alternate form of the UpdateAccountRulesetRule method which supports a Context parameter
func (rulesets *RulesetsV1) UpdateAccountRulesetRuleWithContext(ctx context.Context, updateAccountRulesetRuleOptions *UpdateAccountRulesetRuleOptions) (result *RuleResp, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateAccountRulesetRuleOptions, "updateAccountRulesetRuleOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateAccountRulesetRuleOptions, "updateAccountRulesetRuleOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"crn":        *rulesets.Crn,
		"ruleset_id": *updateAccountRulesetRuleOptions.RulesetID,
		"rule_id":    *updateAccountRulesetRuleOptions.RuleID,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = rulesets.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(rulesets.Service.Options.URL, `/v1/{crn}/rulesets/{ruleset_id}/rules/{rule_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateAccountRulesetRuleOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("rulesets", "V1", "UpdateAccountRulesetRule")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if updateAccountRulesetRuleOptions.Action != nil {
		body["action"] = updateAccountRulesetRuleOptions.Action
	}
	if updateAccountRulesetRuleOptions.ActionParameters != nil {
		body["action_parameters"] = updateAccountRulesetRuleOptions.ActionParameters
	}
	if updateAccountRulesetRuleOptions.Description != nil {
		body["description"] = updateAccountRulesetRuleOptions.Description
	}
	if updateAccountRulesetRuleOptions.Enabled != nil {
		body["enabled"] = updateAccountRulesetRuleOptions.Enabled
	}
	if updateAccountRulesetRuleOptions.Expression != nil {
		body["expression"] = updateAccountRulesetRuleOptions.Expression
	}
	if updateAccountRulesetRuleOptions.ID != nil {
		body["id"] = updateAccountRulesetRuleOptions.ID
	}
	if updateAccountRulesetRuleOptions.Logging != nil {
		body["logging"] = updateAccountRulesetRuleOptions.Logging
	}
	if updateAccountRulesetRuleOptions.Ref != nil {
		body["ref"] = updateAccountRulesetRuleOptions.Ref
	}
	if updateAccountRulesetRuleOptions.Position != nil {
		body["position"] = updateAccountRulesetRuleOptions.Position
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
	response, err = rulesets.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalRuleResp)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeleteAccountRulesetRule : Delete an account ruleset rule
// Delete an account ruleset rule.
func (rulesets *RulesetsV1) DeleteAccountRulesetRule(deleteAccountRulesetRuleOptions *DeleteAccountRulesetRuleOptions) (result *RuleResp, response *core.DetailedResponse, err error) {
	return rulesets.DeleteAccountRulesetRuleWithContext(context.Background(), deleteAccountRulesetRuleOptions)
}

// DeleteAccountRulesetRuleWithContext is an alternate form of the DeleteAccountRulesetRule method which supports a Context parameter
func (rulesets *RulesetsV1) DeleteAccountRulesetRuleWithContext(ctx context.Context, deleteAccountRulesetRuleOptions *DeleteAccountRulesetRuleOptions) (result *RuleResp, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteAccountRulesetRuleOptions, "deleteAccountRulesetRuleOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteAccountRulesetRuleOptions, "deleteAccountRulesetRuleOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"crn":        *rulesets.Crn,
		"ruleset_id": *deleteAccountRulesetRuleOptions.RulesetID,
		"rule_id":    *deleteAccountRulesetRuleOptions.RuleID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = rulesets.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(rulesets.Service.Options.URL, `/v1/{crn}/rulesets/{ruleset_id}/rules/{rule_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteAccountRulesetRuleOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("rulesets", "V1", "DeleteAccountRulesetRule")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = rulesets.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalRuleResp)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetAccountRulesetVersionByTag : List an account ruleset verion's rules by tag
// Lists rules by tag for a specific version of an account ruleset.
func (rulesets *RulesetsV1) GetAccountRulesetVersionByTag(getAccountRulesetVersionByTagOptions *GetAccountRulesetVersionByTagOptions) (result *RulesetResp, response *core.DetailedResponse, err error) {
	return rulesets.GetAccountRulesetVersionByTagWithContext(context.Background(), getAccountRulesetVersionByTagOptions)
}

// GetAccountRulesetVersionByTagWithContext is an alternate form of the GetAccountRulesetVersionByTag method which supports a Context parameter
func (rulesets *RulesetsV1) GetAccountRulesetVersionByTagWithContext(ctx context.Context, getAccountRulesetVersionByTagOptions *GetAccountRulesetVersionByTagOptions) (result *RulesetResp, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getAccountRulesetVersionByTagOptions, "getAccountRulesetVersionByTagOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getAccountRulesetVersionByTagOptions, "getAccountRulesetVersionByTagOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"crn":             *rulesets.Crn,
		"ruleset_id":      *getAccountRulesetVersionByTagOptions.RulesetID,
		"ruleset_version": *getAccountRulesetVersionByTagOptions.RulesetVersion,
		"rule_tag":        *getAccountRulesetVersionByTagOptions.RuleTag,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = rulesets.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(rulesets.Service.Options.URL, `/v1/{crn}/rulesets/{ruleset_id}/versions/{ruleset_version}/by_tag/{rule_tag}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getAccountRulesetVersionByTagOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("rulesets", "V1", "GetAccountRulesetVersionByTag")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = rulesets.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalRulesetResp)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetZoneRulesets : List zone rulesets
// List all rulesets at the zone level.
func (rulesets *RulesetsV1) GetZoneRulesets(getZoneRulesetsOptions *GetZoneRulesetsOptions) (result *ListRulesetsResp, response *core.DetailedResponse, err error) {
	return rulesets.GetZoneRulesetsWithContext(context.Background(), getZoneRulesetsOptions)
}

// GetZoneRulesetsWithContext is an alternate form of the GetZoneRulesets method which supports a Context parameter
func (rulesets *RulesetsV1) GetZoneRulesetsWithContext(ctx context.Context, getZoneRulesetsOptions *GetZoneRulesetsOptions) (result *ListRulesetsResp, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getZoneRulesetsOptions, "getZoneRulesetsOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"crn":             *rulesets.Crn,
		"zone_identifier": *rulesets.ZoneIdentifier,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = rulesets.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(rulesets.Service.Options.URL, `/v1/{crn}/zones/{zone_identifier}/rulesets`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getZoneRulesetsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("rulesets", "V1", "GetZoneRulesets")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = rulesets.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalListRulesetsResp)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetZoneRuleset : Get a zone ruleset
// View a specific zone ruleset.
func (rulesets *RulesetsV1) GetZoneRuleset(getZoneRulesetOptions *GetZoneRulesetOptions) (result *RulesetResp, response *core.DetailedResponse, err error) {
	return rulesets.GetZoneRulesetWithContext(context.Background(), getZoneRulesetOptions)
}

// GetZoneRulesetWithContext is an alternate form of the GetZoneRuleset method which supports a Context parameter
func (rulesets *RulesetsV1) GetZoneRulesetWithContext(ctx context.Context, getZoneRulesetOptions *GetZoneRulesetOptions) (result *RulesetResp, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getZoneRulesetOptions, "getZoneRulesetOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getZoneRulesetOptions, "getZoneRulesetOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"crn":             *rulesets.Crn,
		"zone_identifier": *rulesets.ZoneIdentifier,
		"ruleset_id":      *getZoneRulesetOptions.RulesetID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = rulesets.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(rulesets.Service.Options.URL, `/v1/{crn}/zones/{zone_identifier}/rulesets/{ruleset_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getZoneRulesetOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("rulesets", "V1", "GetZoneRuleset")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = rulesets.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalRulesetResp)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// UpdateZoneRuleset : Update a zone ruleset
// Update a specific zone ruleset.
func (rulesets *RulesetsV1) UpdateZoneRuleset(updateZoneRulesetOptions *UpdateZoneRulesetOptions) (result *RulesetResp, response *core.DetailedResponse, err error) {
	return rulesets.UpdateZoneRulesetWithContext(context.Background(), updateZoneRulesetOptions)
}

// UpdateZoneRulesetWithContext is an alternate form of the UpdateZoneRuleset method which supports a Context parameter
func (rulesets *RulesetsV1) UpdateZoneRulesetWithContext(ctx context.Context, updateZoneRulesetOptions *UpdateZoneRulesetOptions) (result *RulesetResp, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateZoneRulesetOptions, "updateZoneRulesetOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateZoneRulesetOptions, "updateZoneRulesetOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"crn":             *rulesets.Crn,
		"zone_identifier": *rulesets.ZoneIdentifier,
		"ruleset_id":      *updateZoneRulesetOptions.RulesetID,
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = rulesets.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(rulesets.Service.Options.URL, `/v1/{crn}/zones/{zone_identifier}/rulesets/{ruleset_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateZoneRulesetOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("rulesets", "V1", "UpdateZoneRuleset")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if updateZoneRulesetOptions.Description != nil {
		body["description"] = updateZoneRulesetOptions.Description
	}
	if updateZoneRulesetOptions.Kind != nil {
		body["kind"] = updateZoneRulesetOptions.Kind
	}
	if updateZoneRulesetOptions.Name != nil {
		body["name"] = updateZoneRulesetOptions.Name
	}
	if updateZoneRulesetOptions.Phase != nil {
		body["phase"] = updateZoneRulesetOptions.Phase
	}
	if updateZoneRulesetOptions.Rules != nil {
		body["rules"] = updateZoneRulesetOptions.Rules
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
	response, err = rulesets.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalRulesetResp)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeleteZoneRuleset : Delete a zone ruleset
// Delete a specific zone ruleset.
func (rulesets *RulesetsV1) DeleteZoneRuleset(deleteZoneRulesetOptions *DeleteZoneRulesetOptions) (response *core.DetailedResponse, err error) {
	return rulesets.DeleteZoneRulesetWithContext(context.Background(), deleteZoneRulesetOptions)
}

// DeleteZoneRulesetWithContext is an alternate form of the DeleteZoneRuleset method which supports a Context parameter
func (rulesets *RulesetsV1) DeleteZoneRulesetWithContext(ctx context.Context, deleteZoneRulesetOptions *DeleteZoneRulesetOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteZoneRulesetOptions, "deleteZoneRulesetOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteZoneRulesetOptions, "deleteZoneRulesetOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"crn":             *rulesets.Crn,
		"zone_identifier": *rulesets.ZoneIdentifier,
		"ruleset_id":      *deleteZoneRulesetOptions.RulesetID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = rulesets.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(rulesets.Service.Options.URL, `/v1/{crn}/zones/{zone_identifier}/rulesets/{ruleset_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteZoneRulesetOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("rulesets", "V1", "DeleteZoneRuleset")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = rulesets.Service.Request(request, nil)

	return
}

// GetZoneRulesetVersions : List version of a zone ruleset
// List all versions of a specific zone ruleset.
func (rulesets *RulesetsV1) GetZoneRulesetVersions(getZoneRulesetVersionsOptions *GetZoneRulesetVersionsOptions) (result *ListRulesetsResp, response *core.DetailedResponse, err error) {
	return rulesets.GetZoneRulesetVersionsWithContext(context.Background(), getZoneRulesetVersionsOptions)
}

// GetZoneRulesetVersionsWithContext is an alternate form of the GetZoneRulesetVersions method which supports a Context parameter
func (rulesets *RulesetsV1) GetZoneRulesetVersionsWithContext(ctx context.Context, getZoneRulesetVersionsOptions *GetZoneRulesetVersionsOptions) (result *ListRulesetsResp, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getZoneRulesetVersionsOptions, "getZoneRulesetVersionsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getZoneRulesetVersionsOptions, "getZoneRulesetVersionsOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"crn":             *rulesets.Crn,
		"zone_identifier": *rulesets.ZoneIdentifier,
		"ruleset_id":      *getZoneRulesetVersionsOptions.RulesetID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = rulesets.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(rulesets.Service.Options.URL, `/v1/{crn}/zones/{zone_identifier}/rulesets/{ruleset_id}/versions`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getZoneRulesetVersionsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("rulesets", "V1", "GetZoneRulesetVersions")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = rulesets.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalListRulesetsResp)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetZoneRulesetVersion : Get a specific version of a zone ruleset
// View a specific version of a specific zone ruleset.
func (rulesets *RulesetsV1) GetZoneRulesetVersion(getZoneRulesetVersionOptions *GetZoneRulesetVersionOptions) (result *RulesetResp, response *core.DetailedResponse, err error) {
	return rulesets.GetZoneRulesetVersionWithContext(context.Background(), getZoneRulesetVersionOptions)
}

// GetZoneRulesetVersionWithContext is an alternate form of the GetZoneRulesetVersion method which supports a Context parameter
func (rulesets *RulesetsV1) GetZoneRulesetVersionWithContext(ctx context.Context, getZoneRulesetVersionOptions *GetZoneRulesetVersionOptions) (result *RulesetResp, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getZoneRulesetVersionOptions, "getZoneRulesetVersionOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getZoneRulesetVersionOptions, "getZoneRulesetVersionOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"crn":             *rulesets.Crn,
		"zone_identifier": *rulesets.ZoneIdentifier,
		"ruleset_id":      *getZoneRulesetVersionOptions.RulesetID,
		"ruleset_version": *getZoneRulesetVersionOptions.RulesetVersion,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = rulesets.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(rulesets.Service.Options.URL, `/v1/{crn}/zones/{zone_identifier}/rulesets/{ruleset_id}/versions/{ruleset_version}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getZoneRulesetVersionOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("rulesets", "V1", "GetZoneRulesetVersion")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = rulesets.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalRulesetResp)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeleteZoneRulesetVersion : Delete a specific version of a zone ruleset
// Delete a specific version of a specific zone ruleset.
func (rulesets *RulesetsV1) DeleteZoneRulesetVersion(deleteZoneRulesetVersionOptions *DeleteZoneRulesetVersionOptions) (response *core.DetailedResponse, err error) {
	return rulesets.DeleteZoneRulesetVersionWithContext(context.Background(), deleteZoneRulesetVersionOptions)
}

// DeleteZoneRulesetVersionWithContext is an alternate form of the DeleteZoneRulesetVersion method which supports a Context parameter
func (rulesets *RulesetsV1) DeleteZoneRulesetVersionWithContext(ctx context.Context, deleteZoneRulesetVersionOptions *DeleteZoneRulesetVersionOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteZoneRulesetVersionOptions, "deleteZoneRulesetVersionOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteZoneRulesetVersionOptions, "deleteZoneRulesetVersionOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"crn":             *rulesets.Crn,
		"zone_identifier": *rulesets.ZoneIdentifier,
		"ruleset_id":      *deleteZoneRulesetVersionOptions.RulesetID,
		"ruleset_version": *deleteZoneRulesetVersionOptions.RulesetVersion,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = rulesets.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(rulesets.Service.Options.URL, `/v1/{crn}/zones/{zone_identifier}/rulesets/{ruleset_id}/versions/{ruleset_version}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteZoneRulesetVersionOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("rulesets", "V1", "DeleteZoneRulesetVersion")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = rulesets.Service.Request(request, nil)

	return
}

// GetZoneEntrypointRuleset : Get a zone entrypoint ruleset
// Get the zone ruleset for the given phase's entrypoint.
func (rulesets *RulesetsV1) GetZoneEntrypointRuleset(getZoneEntrypointRulesetOptions *GetZoneEntrypointRulesetOptions) (result *RulesetResp, response *core.DetailedResponse, err error) {
	return rulesets.GetZoneEntrypointRulesetWithContext(context.Background(), getZoneEntrypointRulesetOptions)
}

// GetZoneEntrypointRulesetWithContext is an alternate form of the GetZoneEntrypointRuleset method which supports a Context parameter
func (rulesets *RulesetsV1) GetZoneEntrypointRulesetWithContext(ctx context.Context, getZoneEntrypointRulesetOptions *GetZoneEntrypointRulesetOptions) (result *RulesetResp, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getZoneEntrypointRulesetOptions, "getZoneEntrypointRulesetOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getZoneEntrypointRulesetOptions, "getZoneEntrypointRulesetOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"crn":             *rulesets.Crn,
		"zone_identifier": *rulesets.ZoneIdentifier,
		"ruleset_phase":   *getZoneEntrypointRulesetOptions.RulesetPhase,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = rulesets.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(rulesets.Service.Options.URL, `/v1/{crn}/zones/{zone_identifier}/rulesets/phases/{ruleset_phase}/entrypoint`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getZoneEntrypointRulesetOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("rulesets", "V1", "GetZoneEntrypointRuleset")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = rulesets.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalRulesetResp)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// UpdateZoneEntrypointRuleset : Update a zone entrypoint ruleset
// Updates the account ruleset for the given phase's entry point.
func (rulesets *RulesetsV1) UpdateZoneEntrypointRuleset(updateZoneEntrypointRulesetOptions *UpdateZoneEntrypointRulesetOptions) (result *RulesetResp, response *core.DetailedResponse, err error) {
	return rulesets.UpdateZoneEntrypointRulesetWithContext(context.Background(), updateZoneEntrypointRulesetOptions)
}

// UpdateZoneEntrypointRulesetWithContext is an alternate form of the UpdateZoneEntrypointRuleset method which supports a Context parameter
func (rulesets *RulesetsV1) UpdateZoneEntrypointRulesetWithContext(ctx context.Context, updateZoneEntrypointRulesetOptions *UpdateZoneEntrypointRulesetOptions) (result *RulesetResp, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateZoneEntrypointRulesetOptions, "updateZoneEntrypointRulesetOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateZoneEntrypointRulesetOptions, "updateZoneEntrypointRulesetOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"crn":             *rulesets.Crn,
		"zone_identifier": *rulesets.ZoneIdentifier,
		"ruleset_phase":   *updateZoneEntrypointRulesetOptions.RulesetPhase,
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = rulesets.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(rulesets.Service.Options.URL, `/v1/{crn}/zones/{zone_identifier}/rulesets/phases/{ruleset_phase}/entrypoint`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateZoneEntrypointRulesetOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("rulesets", "V1", "UpdateZoneEntrypointRuleset")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if updateZoneEntrypointRulesetOptions.Description != nil {
		body["description"] = updateZoneEntrypointRulesetOptions.Description
	}
	if updateZoneEntrypointRulesetOptions.Kind != nil {
		body["kind"] = updateZoneEntrypointRulesetOptions.Kind
	}
	if updateZoneEntrypointRulesetOptions.Name != nil {
		body["name"] = updateZoneEntrypointRulesetOptions.Name
	}
	if updateZoneEntrypointRulesetOptions.Phase != nil {
		body["phase"] = updateZoneEntrypointRulesetOptions.Phase
	}
	if updateZoneEntrypointRulesetOptions.Rules != nil {
		body["rules"] = updateZoneEntrypointRulesetOptions.Rules
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
	response, err = rulesets.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalRulesetResp)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetZoneEntryPointRulesetVersions : List a zone entry point ruleset's versions
// Lists the zone ruleset versions for the given phase's entry point.
func (rulesets *RulesetsV1) GetZoneEntryPointRulesetVersions(getZoneEntryPointRulesetVersionsOptions *GetZoneEntryPointRulesetVersionsOptions) (result *ListRulesetsResp, response *core.DetailedResponse, err error) {
	return rulesets.GetZoneEntryPointRulesetVersionsWithContext(context.Background(), getZoneEntryPointRulesetVersionsOptions)
}

// GetZoneEntryPointRulesetVersionsWithContext is an alternate form of the GetZoneEntryPointRulesetVersions method which supports a Context parameter
func (rulesets *RulesetsV1) GetZoneEntryPointRulesetVersionsWithContext(ctx context.Context, getZoneEntryPointRulesetVersionsOptions *GetZoneEntryPointRulesetVersionsOptions) (result *ListRulesetsResp, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getZoneEntryPointRulesetVersionsOptions, "getZoneEntryPointRulesetVersionsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getZoneEntryPointRulesetVersionsOptions, "getZoneEntryPointRulesetVersionsOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"crn":             *rulesets.Crn,
		"zone_identifier": *rulesets.ZoneIdentifier,
		"ruleset_phase":   *getZoneEntryPointRulesetVersionsOptions.RulesetPhase,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = rulesets.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(rulesets.Service.Options.URL, `/v1/{crn}/zones/{zone_identifier}/rulesets/phases/{ruleset_phase}/entrypoint/versions`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getZoneEntryPointRulesetVersionsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("rulesets", "V1", "GetZoneEntryPointRulesetVersions")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = rulesets.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalListRulesetsResp)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetZoneEntryPointRulesetVersion : Get a zone entry point ruleset version
// Fetches a specific version of a zone entry point ruleset.
func (rulesets *RulesetsV1) GetZoneEntryPointRulesetVersion(getZoneEntryPointRulesetVersionOptions *GetZoneEntryPointRulesetVersionOptions) (result *RulesetResp, response *core.DetailedResponse, err error) {
	return rulesets.GetZoneEntryPointRulesetVersionWithContext(context.Background(), getZoneEntryPointRulesetVersionOptions)
}

// GetZoneEntryPointRulesetVersionWithContext is an alternate form of the GetZoneEntryPointRulesetVersion method which supports a Context parameter
func (rulesets *RulesetsV1) GetZoneEntryPointRulesetVersionWithContext(ctx context.Context, getZoneEntryPointRulesetVersionOptions *GetZoneEntryPointRulesetVersionOptions) (result *RulesetResp, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getZoneEntryPointRulesetVersionOptions, "getZoneEntryPointRulesetVersionOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getZoneEntryPointRulesetVersionOptions, "getZoneEntryPointRulesetVersionOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"crn":             *rulesets.Crn,
		"zone_identifier": *rulesets.ZoneIdentifier,
		"ruleset_phase":   *getZoneEntryPointRulesetVersionOptions.RulesetPhase,
		"ruleset_version": *getZoneEntryPointRulesetVersionOptions.RulesetVersion,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = rulesets.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(rulesets.Service.Options.URL, `/v1/{crn}/zones/{zone_identifier}/rulesets/phases/{ruleset_phase}/entrypoint/versions/{ruleset_version}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getZoneEntryPointRulesetVersionOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("rulesets", "V1", "GetZoneEntryPointRulesetVersion")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = rulesets.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalRulesetResp)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// CreateZoneRulesetRule : Create a zone ruleset rule
// Create a zone ruleset rule.
func (rulesets *RulesetsV1) CreateZoneRulesetRule(createZoneRulesetRuleOptions *CreateZoneRulesetRuleOptions) (result *RuleResp, response *core.DetailedResponse, err error) {
	return rulesets.CreateZoneRulesetRuleWithContext(context.Background(), createZoneRulesetRuleOptions)
}

// CreateZoneRulesetRuleWithContext is an alternate form of the CreateZoneRulesetRule method which supports a Context parameter
func (rulesets *RulesetsV1) CreateZoneRulesetRuleWithContext(ctx context.Context, createZoneRulesetRuleOptions *CreateZoneRulesetRuleOptions) (result *RuleResp, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createZoneRulesetRuleOptions, "createZoneRulesetRuleOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createZoneRulesetRuleOptions, "createZoneRulesetRuleOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"crn":             *rulesets.Crn,
		"zone_identifier": *rulesets.ZoneIdentifier,
		"ruleset_id":      *createZoneRulesetRuleOptions.RulesetID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = rulesets.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(rulesets.Service.Options.URL, `/v1/{crn}/zones/{zone_identifier}/rulesets/{ruleset_id}/rules`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range createZoneRulesetRuleOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("rulesets", "V1", "CreateZoneRulesetRule")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if createZoneRulesetRuleOptions.Action != nil {
		body["action"] = createZoneRulesetRuleOptions.Action
	}
	if createZoneRulesetRuleOptions.ActionParameters != nil {
		body["action_parameters"] = createZoneRulesetRuleOptions.ActionParameters
	}
	if createZoneRulesetRuleOptions.Description != nil {
		body["description"] = createZoneRulesetRuleOptions.Description
	}
	if createZoneRulesetRuleOptions.Enabled != nil {
		body["enabled"] = createZoneRulesetRuleOptions.Enabled
	}
	if createZoneRulesetRuleOptions.Expression != nil {
		body["expression"] = createZoneRulesetRuleOptions.Expression
	}
	if createZoneRulesetRuleOptions.ID != nil {
		body["id"] = createZoneRulesetRuleOptions.ID
	}
	if createZoneRulesetRuleOptions.Logging != nil {
		body["logging"] = createZoneRulesetRuleOptions.Logging
	}
	if createZoneRulesetRuleOptions.Ref != nil {
		body["ref"] = createZoneRulesetRuleOptions.Ref
	}
	if createZoneRulesetRuleOptions.Position != nil {
		body["position"] = createZoneRulesetRuleOptions.Position
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
	response, err = rulesets.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalRuleResp)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// UpdateZoneRulesetRule : Update a zone ruleset rule
// Update a zone ruleset rule.
func (rulesets *RulesetsV1) UpdateZoneRulesetRule(updateZoneRulesetRuleOptions *UpdateZoneRulesetRuleOptions) (result *RuleResp, response *core.DetailedResponse, err error) {
	return rulesets.UpdateZoneRulesetRuleWithContext(context.Background(), updateZoneRulesetRuleOptions)
}

// UpdateZoneRulesetRuleWithContext is an alternate form of the UpdateZoneRulesetRule method which supports a Context parameter
func (rulesets *RulesetsV1) UpdateZoneRulesetRuleWithContext(ctx context.Context, updateZoneRulesetRuleOptions *UpdateZoneRulesetRuleOptions) (result *RuleResp, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateZoneRulesetRuleOptions, "updateZoneRulesetRuleOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateZoneRulesetRuleOptions, "updateZoneRulesetRuleOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"crn":             *rulesets.Crn,
		"zone_identifier": *rulesets.ZoneIdentifier,
		"ruleset_id":      *updateZoneRulesetRuleOptions.RulesetID,
		"rule_id":         *updateZoneRulesetRuleOptions.RuleID,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = rulesets.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(rulesets.Service.Options.URL, `/v1/{crn}/zones/{zone_identifier}/rulesets/{ruleset_id}/rules/{rule_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateZoneRulesetRuleOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("rulesets", "V1", "UpdateZoneRulesetRule")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if updateZoneRulesetRuleOptions.Action != nil {
		body["action"] = updateZoneRulesetRuleOptions.Action
	}
	if updateZoneRulesetRuleOptions.ActionParameters != nil {
		body["action_parameters"] = updateZoneRulesetRuleOptions.ActionParameters
	}
	if updateZoneRulesetRuleOptions.Description != nil {
		body["description"] = updateZoneRulesetRuleOptions.Description
	}
	if updateZoneRulesetRuleOptions.Enabled != nil {
		body["enabled"] = updateZoneRulesetRuleOptions.Enabled
	}
	if updateZoneRulesetRuleOptions.Expression != nil {
		body["expression"] = updateZoneRulesetRuleOptions.Expression
	}
	if updateZoneRulesetRuleOptions.ID != nil {
		body["id"] = updateZoneRulesetRuleOptions.ID
	}
	if updateZoneRulesetRuleOptions.Logging != nil {
		body["logging"] = updateZoneRulesetRuleOptions.Logging
	}
	if updateZoneRulesetRuleOptions.Ref != nil {
		body["ref"] = updateZoneRulesetRuleOptions.Ref
	}
	if updateZoneRulesetRuleOptions.Position != nil {
		body["position"] = updateZoneRulesetRuleOptions.Position
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
	response, err = rulesets.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalRuleResp)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeleteZoneRulesetRule : Delete a zone ruleset rule
// Delete an account ruleset rule.
func (rulesets *RulesetsV1) DeleteZoneRulesetRule(deleteZoneRulesetRuleOptions *DeleteZoneRulesetRuleOptions) (result *RuleResp, response *core.DetailedResponse, err error) {
	return rulesets.DeleteZoneRulesetRuleWithContext(context.Background(), deleteZoneRulesetRuleOptions)
}

// DeleteZoneRulesetRuleWithContext is an alternate form of the DeleteZoneRulesetRule method which supports a Context parameter
func (rulesets *RulesetsV1) DeleteZoneRulesetRuleWithContext(ctx context.Context, deleteZoneRulesetRuleOptions *DeleteZoneRulesetRuleOptions) (result *RuleResp, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteZoneRulesetRuleOptions, "deleteZoneRulesetRuleOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteZoneRulesetRuleOptions, "deleteZoneRulesetRuleOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"crn":             *rulesets.Crn,
		"zone_identifier": *rulesets.ZoneIdentifier,
		"ruleset_id":      *deleteZoneRulesetRuleOptions.RulesetID,
		"rule_id":         *deleteZoneRulesetRuleOptions.RuleID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = rulesets.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(rulesets.Service.Options.URL, `/v1/{crn}/zones/{zone_identifier}/rulesets/{ruleset_id}/rules/{rule_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteZoneRulesetRuleOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("rulesets", "V1", "DeleteZoneRulesetRule")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = rulesets.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalRuleResp)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ActionParametersResponse : ActionParametersResponse struct
type ActionParametersResponse struct {
	// the content to return.
	Content *string `json:"content" validate:"required"`

	ContentType *string `json:"content_type" validate:"required"`

	// The status code to return.
	StatusCode *int64 `json:"status_code" validate:"required"`
}

// NewActionParametersResponse : Instantiate ActionParametersResponse (Generic Model Constructor)
func (*RulesetsV1) NewActionParametersResponse(content string, contentType string, statusCode int64) (_model *ActionParametersResponse, err error) {
	_model = &ActionParametersResponse{
		Content:     core.StringPtr(content),
		ContentType: core.StringPtr(contentType),
		StatusCode:  core.Int64Ptr(statusCode),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalActionParametersResponse unmarshals an instance of ActionParametersResponse from the specified map of raw messages.
func UnmarshalActionParametersResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ActionParametersResponse)
	err = core.UnmarshalPrimitive(m, "content", &obj.Content)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "content_type", &obj.ContentType)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "status_code", &obj.StatusCode)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CreateAccountRulesetRuleOptions : The CreateAccountRulesetRule options.
type CreateAccountRulesetRuleOptions struct {
	// ID of a specific ruleset.
	RulesetID *string `json:"ruleset_id" validate:"required,ne="`

	// What happens when theres a match for the rule expression.
	Action *string `json:"action,omitempty"`

	ActionParameters *ActionParameters `json:"action_parameters,omitempty"`

	Description *string `json:"description,omitempty"`

	Enabled *bool `json:"enabled,omitempty"`

	// The expression defining which traffic will match the rule.
	Expression *string `json:"expression,omitempty"`

	ID *string `json:"id,omitempty"`

	Logging *Logging `json:"logging,omitempty"`

	// The reference of the rule (the rule ID by default).
	Ref *string `json:"ref,omitempty"`

	Position *Position `json:"position,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCreateAccountRulesetRuleOptions : Instantiate CreateAccountRulesetRuleOptions
func (*RulesetsV1) NewCreateAccountRulesetRuleOptions(rulesetID string) *CreateAccountRulesetRuleOptions {
	return &CreateAccountRulesetRuleOptions{
		RulesetID: core.StringPtr(rulesetID),
	}
}

// SetRulesetID : Allow user to set RulesetID
func (_options *CreateAccountRulesetRuleOptions) SetRulesetID(rulesetID string) *CreateAccountRulesetRuleOptions {
	_options.RulesetID = core.StringPtr(rulesetID)
	return _options
}

// SetAction : Allow user to set Action
func (_options *CreateAccountRulesetRuleOptions) SetAction(action string) *CreateAccountRulesetRuleOptions {
	_options.Action = core.StringPtr(action)
	return _options
}

// SetActionParameters : Allow user to set ActionParameters
func (_options *CreateAccountRulesetRuleOptions) SetActionParameters(actionParameters *ActionParameters) *CreateAccountRulesetRuleOptions {
	_options.ActionParameters = actionParameters
	return _options
}

// SetDescription : Allow user to set Description
func (_options *CreateAccountRulesetRuleOptions) SetDescription(description string) *CreateAccountRulesetRuleOptions {
	_options.Description = core.StringPtr(description)
	return _options
}

// SetEnabled : Allow user to set Enabled
func (_options *CreateAccountRulesetRuleOptions) SetEnabled(enabled bool) *CreateAccountRulesetRuleOptions {
	_options.Enabled = core.BoolPtr(enabled)
	return _options
}

// SetExpression : Allow user to set Expression
func (_options *CreateAccountRulesetRuleOptions) SetExpression(expression string) *CreateAccountRulesetRuleOptions {
	_options.Expression = core.StringPtr(expression)
	return _options
}

// SetID : Allow user to set ID
func (_options *CreateAccountRulesetRuleOptions) SetID(id string) *CreateAccountRulesetRuleOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetLogging : Allow user to set Logging
func (_options *CreateAccountRulesetRuleOptions) SetLogging(logging *Logging) *CreateAccountRulesetRuleOptions {
	_options.Logging = logging
	return _options
}

// SetRef : Allow user to set Ref
func (_options *CreateAccountRulesetRuleOptions) SetRef(ref string) *CreateAccountRulesetRuleOptions {
	_options.Ref = core.StringPtr(ref)
	return _options
}

// SetPosition : Allow user to set Position
func (_options *CreateAccountRulesetRuleOptions) SetPosition(position *Position) *CreateAccountRulesetRuleOptions {
	_options.Position = position
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateAccountRulesetRuleOptions) SetHeaders(param map[string]string) *CreateAccountRulesetRuleOptions {
	options.Headers = param
	return options
}

// CreateZoneRulesetRuleOptions : The CreateZoneRulesetRule options.
type CreateZoneRulesetRuleOptions struct {
	// ID of a specific ruleset.
	RulesetID *string `json:"ruleset_id" validate:"required,ne="`

	// What happens when theres a match for the rule expression.
	Action *string `json:"action,omitempty"`

	ActionParameters *ActionParameters `json:"action_parameters,omitempty"`

	Description *string `json:"description,omitempty"`

	Enabled *bool `json:"enabled,omitempty"`

	// The expression defining which traffic will match the rule.
	Expression *string `json:"expression,omitempty"`

	ID *string `json:"id,omitempty"`

	Logging *Logging `json:"logging,omitempty"`

	// The reference of the rule (the rule ID by default).
	Ref *string `json:"ref,omitempty"`

	Position *Position `json:"position,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCreateZoneRulesetRuleOptions : Instantiate CreateZoneRulesetRuleOptions
func (*RulesetsV1) NewCreateZoneRulesetRuleOptions(rulesetID string) *CreateZoneRulesetRuleOptions {
	return &CreateZoneRulesetRuleOptions{
		RulesetID: core.StringPtr(rulesetID),
	}
}

// SetRulesetID : Allow user to set RulesetID
func (_options *CreateZoneRulesetRuleOptions) SetRulesetID(rulesetID string) *CreateZoneRulesetRuleOptions {
	_options.RulesetID = core.StringPtr(rulesetID)
	return _options
}

// SetAction : Allow user to set Action
func (_options *CreateZoneRulesetRuleOptions) SetAction(action string) *CreateZoneRulesetRuleOptions {
	_options.Action = core.StringPtr(action)
	return _options
}

// SetActionParameters : Allow user to set ActionParameters
func (_options *CreateZoneRulesetRuleOptions) SetActionParameters(actionParameters *ActionParameters) *CreateZoneRulesetRuleOptions {
	_options.ActionParameters = actionParameters
	return _options
}

// SetDescription : Allow user to set Description
func (_options *CreateZoneRulesetRuleOptions) SetDescription(description string) *CreateZoneRulesetRuleOptions {
	_options.Description = core.StringPtr(description)
	return _options
}

// SetEnabled : Allow user to set Enabled
func (_options *CreateZoneRulesetRuleOptions) SetEnabled(enabled bool) *CreateZoneRulesetRuleOptions {
	_options.Enabled = core.BoolPtr(enabled)
	return _options
}

// SetExpression : Allow user to set Expression
func (_options *CreateZoneRulesetRuleOptions) SetExpression(expression string) *CreateZoneRulesetRuleOptions {
	_options.Expression = core.StringPtr(expression)
	return _options
}

// SetID : Allow user to set ID
func (_options *CreateZoneRulesetRuleOptions) SetID(id string) *CreateZoneRulesetRuleOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetLogging : Allow user to set Logging
func (_options *CreateZoneRulesetRuleOptions) SetLogging(logging *Logging) *CreateZoneRulesetRuleOptions {
	_options.Logging = logging
	return _options
}

// SetRef : Allow user to set Ref
func (_options *CreateZoneRulesetRuleOptions) SetRef(ref string) *CreateZoneRulesetRuleOptions {
	_options.Ref = core.StringPtr(ref)
	return _options
}

// SetPosition : Allow user to set Position
func (_options *CreateZoneRulesetRuleOptions) SetPosition(position *Position) *CreateZoneRulesetRuleOptions {
	_options.Position = position
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateZoneRulesetRuleOptions) SetHeaders(param map[string]string) *CreateZoneRulesetRuleOptions {
	options.Headers = param
	return options
}

// DeleteAccountRulesetOptions : The DeleteAccountRuleset options.
type DeleteAccountRulesetOptions struct {
	// ID of a specific ruleset.
	RulesetID *string `json:"ruleset_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteAccountRulesetOptions : Instantiate DeleteAccountRulesetOptions
func (*RulesetsV1) NewDeleteAccountRulesetOptions(rulesetID string) *DeleteAccountRulesetOptions {
	return &DeleteAccountRulesetOptions{
		RulesetID: core.StringPtr(rulesetID),
	}
}

// SetRulesetID : Allow user to set RulesetID
func (_options *DeleteAccountRulesetOptions) SetRulesetID(rulesetID string) *DeleteAccountRulesetOptions {
	_options.RulesetID = core.StringPtr(rulesetID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteAccountRulesetOptions) SetHeaders(param map[string]string) *DeleteAccountRulesetOptions {
	options.Headers = param
	return options
}

// DeleteAccountRulesetRuleOptions : The DeleteAccountRulesetRule options.
type DeleteAccountRulesetRuleOptions struct {
	// ID of a specific ruleset.
	RulesetID *string `json:"ruleset_id" validate:"required,ne="`

	// ID of a specific rule.
	RuleID *string `json:"rule_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteAccountRulesetRuleOptions : Instantiate DeleteAccountRulesetRuleOptions
func (*RulesetsV1) NewDeleteAccountRulesetRuleOptions(rulesetID string, ruleID string) *DeleteAccountRulesetRuleOptions {
	return &DeleteAccountRulesetRuleOptions{
		RulesetID: core.StringPtr(rulesetID),
		RuleID:    core.StringPtr(ruleID),
	}
}

// SetRulesetID : Allow user to set RulesetID
func (_options *DeleteAccountRulesetRuleOptions) SetRulesetID(rulesetID string) *DeleteAccountRulesetRuleOptions {
	_options.RulesetID = core.StringPtr(rulesetID)
	return _options
}

// SetRuleID : Allow user to set RuleID
func (_options *DeleteAccountRulesetRuleOptions) SetRuleID(ruleID string) *DeleteAccountRulesetRuleOptions {
	_options.RuleID = core.StringPtr(ruleID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteAccountRulesetRuleOptions) SetHeaders(param map[string]string) *DeleteAccountRulesetRuleOptions {
	options.Headers = param
	return options
}

// DeleteAccountRulesetVersionOptions : The DeleteAccountRulesetVersion options.
type DeleteAccountRulesetVersionOptions struct {
	// ID of a specific ruleset.
	RulesetID *string `json:"ruleset_id" validate:"required,ne="`

	// The version of the ruleset.
	RulesetVersion *string `json:"ruleset_version" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteAccountRulesetVersionOptions : Instantiate DeleteAccountRulesetVersionOptions
func (*RulesetsV1) NewDeleteAccountRulesetVersionOptions(rulesetID string, rulesetVersion string) *DeleteAccountRulesetVersionOptions {
	return &DeleteAccountRulesetVersionOptions{
		RulesetID:      core.StringPtr(rulesetID),
		RulesetVersion: core.StringPtr(rulesetVersion),
	}
}

// SetRulesetID : Allow user to set RulesetID
func (_options *DeleteAccountRulesetVersionOptions) SetRulesetID(rulesetID string) *DeleteAccountRulesetVersionOptions {
	_options.RulesetID = core.StringPtr(rulesetID)
	return _options
}

// SetRulesetVersion : Allow user to set RulesetVersion
func (_options *DeleteAccountRulesetVersionOptions) SetRulesetVersion(rulesetVersion string) *DeleteAccountRulesetVersionOptions {
	_options.RulesetVersion = core.StringPtr(rulesetVersion)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteAccountRulesetVersionOptions) SetHeaders(param map[string]string) *DeleteAccountRulesetVersionOptions {
	options.Headers = param
	return options
}

// DeleteZoneRulesetOptions : The DeleteZoneRuleset options.
type DeleteZoneRulesetOptions struct {
	// ID of a specific ruleset.
	RulesetID *string `json:"ruleset_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteZoneRulesetOptions : Instantiate DeleteZoneRulesetOptions
func (*RulesetsV1) NewDeleteZoneRulesetOptions(rulesetID string) *DeleteZoneRulesetOptions {
	return &DeleteZoneRulesetOptions{
		RulesetID: core.StringPtr(rulesetID),
	}
}

// SetRulesetID : Allow user to set RulesetID
func (_options *DeleteZoneRulesetOptions) SetRulesetID(rulesetID string) *DeleteZoneRulesetOptions {
	_options.RulesetID = core.StringPtr(rulesetID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteZoneRulesetOptions) SetHeaders(param map[string]string) *DeleteZoneRulesetOptions {
	options.Headers = param
	return options
}

// DeleteZoneRulesetRuleOptions : The DeleteZoneRulesetRule options.
type DeleteZoneRulesetRuleOptions struct {
	// ID of a specific ruleset.
	RulesetID *string `json:"ruleset_id" validate:"required,ne="`

	// ID of a specific rule.
	RuleID *string `json:"rule_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteZoneRulesetRuleOptions : Instantiate DeleteZoneRulesetRuleOptions
func (*RulesetsV1) NewDeleteZoneRulesetRuleOptions(rulesetID string, ruleID string) *DeleteZoneRulesetRuleOptions {
	return &DeleteZoneRulesetRuleOptions{
		RulesetID: core.StringPtr(rulesetID),
		RuleID:    core.StringPtr(ruleID),
	}
}

// SetRulesetID : Allow user to set RulesetID
func (_options *DeleteZoneRulesetRuleOptions) SetRulesetID(rulesetID string) *DeleteZoneRulesetRuleOptions {
	_options.RulesetID = core.StringPtr(rulesetID)
	return _options
}

// SetRuleID : Allow user to set RuleID
func (_options *DeleteZoneRulesetRuleOptions) SetRuleID(ruleID string) *DeleteZoneRulesetRuleOptions {
	_options.RuleID = core.StringPtr(ruleID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteZoneRulesetRuleOptions) SetHeaders(param map[string]string) *DeleteZoneRulesetRuleOptions {
	options.Headers = param
	return options
}

// DeleteZoneRulesetVersionOptions : The DeleteZoneRulesetVersion options.
type DeleteZoneRulesetVersionOptions struct {
	// ID of a specific ruleset.
	RulesetID *string `json:"ruleset_id" validate:"required,ne="`

	// The version of the ruleset.
	RulesetVersion *string `json:"ruleset_version" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteZoneRulesetVersionOptions : Instantiate DeleteZoneRulesetVersionOptions
func (*RulesetsV1) NewDeleteZoneRulesetVersionOptions(rulesetID string, rulesetVersion string) *DeleteZoneRulesetVersionOptions {
	return &DeleteZoneRulesetVersionOptions{
		RulesetID:      core.StringPtr(rulesetID),
		RulesetVersion: core.StringPtr(rulesetVersion),
	}
}

// SetRulesetID : Allow user to set RulesetID
func (_options *DeleteZoneRulesetVersionOptions) SetRulesetID(rulesetID string) *DeleteZoneRulesetVersionOptions {
	_options.RulesetID = core.StringPtr(rulesetID)
	return _options
}

// SetRulesetVersion : Allow user to set RulesetVersion
func (_options *DeleteZoneRulesetVersionOptions) SetRulesetVersion(rulesetVersion string) *DeleteZoneRulesetVersionOptions {
	_options.RulesetVersion = core.StringPtr(rulesetVersion)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteZoneRulesetVersionOptions) SetHeaders(param map[string]string) *DeleteZoneRulesetVersionOptions {
	options.Headers = param
	return options
}

// GetAccountEntryPointRulesetVersionOptions : The GetAccountEntryPointRulesetVersion options.
type GetAccountEntryPointRulesetVersionOptions struct {
	// The phase of the ruleset.
	RulesetPhase *string `json:"ruleset_phase" validate:"required,ne="`

	// The version of the ruleset.
	RulesetVersion *string `json:"ruleset_version" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the GetAccountEntryPointRulesetVersionOptions.RulesetPhase property.
// The phase of the ruleset.
const (
	GetAccountEntryPointRulesetVersionOptions_RulesetPhase_DdosL4                         = "ddos_l4"
	GetAccountEntryPointRulesetVersionOptions_RulesetPhase_DdosL7                         = "ddos_l7"
	GetAccountEntryPointRulesetVersionOptions_RulesetPhase_HttpConfigSettings             = "http_config_settings"
	GetAccountEntryPointRulesetVersionOptions_RulesetPhase_HttpCustomErrors               = "http_custom_errors"
	GetAccountEntryPointRulesetVersionOptions_RulesetPhase_HttpLogCustomFields            = "http_log_custom_fields"
	GetAccountEntryPointRulesetVersionOptions_RulesetPhase_HttpRatelimit                  = "http_ratelimit"
	GetAccountEntryPointRulesetVersionOptions_RulesetPhase_HttpRequestCacheSettings       = "http_request_cache_settings"
	GetAccountEntryPointRulesetVersionOptions_RulesetPhase_HttpRequestDynamicRedirect     = "http_request_dynamic_redirect"
	GetAccountEntryPointRulesetVersionOptions_RulesetPhase_HttpRequestFirewallCustom      = "http_request_firewall_custom"
	GetAccountEntryPointRulesetVersionOptions_RulesetPhase_HttpRequestFirewallManaged     = "http_request_firewall_managed"
	GetAccountEntryPointRulesetVersionOptions_RulesetPhase_HttpRequestLateTransform       = "http_request_late_transform"
	GetAccountEntryPointRulesetVersionOptions_RulesetPhase_HttpRequestOrigin              = "http_request_origin"
	GetAccountEntryPointRulesetVersionOptions_RulesetPhase_HttpRequestRedirect            = "http_request_redirect"
	GetAccountEntryPointRulesetVersionOptions_RulesetPhase_HttpRequestSanitize            = "http_request_sanitize"
	GetAccountEntryPointRulesetVersionOptions_RulesetPhase_HttpRequestSbfm                = "http_request_sbfm"
	GetAccountEntryPointRulesetVersionOptions_RulesetPhase_HttpRequestSelectConfiguration = "http_request_select_configuration"
	GetAccountEntryPointRulesetVersionOptions_RulesetPhase_HttpRequestTransform           = "http_request_transform"
	GetAccountEntryPointRulesetVersionOptions_RulesetPhase_HttpResponseCompression        = "http_response_compression"
	GetAccountEntryPointRulesetVersionOptions_RulesetPhase_HttpResponseFirewallManaged    = "http_response_firewall_managed"
	GetAccountEntryPointRulesetVersionOptions_RulesetPhase_HttpResponseHeadersTransform   = "http_response_headers_transform"
)

// NewGetAccountEntryPointRulesetVersionOptions : Instantiate GetAccountEntryPointRulesetVersionOptions
func (*RulesetsV1) NewGetAccountEntryPointRulesetVersionOptions(rulesetPhase string, rulesetVersion string) *GetAccountEntryPointRulesetVersionOptions {
	return &GetAccountEntryPointRulesetVersionOptions{
		RulesetPhase:   core.StringPtr(rulesetPhase),
		RulesetVersion: core.StringPtr(rulesetVersion),
	}
}

// SetRulesetPhase : Allow user to set RulesetPhase
func (_options *GetAccountEntryPointRulesetVersionOptions) SetRulesetPhase(rulesetPhase string) *GetAccountEntryPointRulesetVersionOptions {
	_options.RulesetPhase = core.StringPtr(rulesetPhase)
	return _options
}

// SetRulesetVersion : Allow user to set RulesetVersion
func (_options *GetAccountEntryPointRulesetVersionOptions) SetRulesetVersion(rulesetVersion string) *GetAccountEntryPointRulesetVersionOptions {
	_options.RulesetVersion = core.StringPtr(rulesetVersion)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetAccountEntryPointRulesetVersionOptions) SetHeaders(param map[string]string) *GetAccountEntryPointRulesetVersionOptions {
	options.Headers = param
	return options
}

// GetAccountEntryPointRulesetVersionsOptions : The GetAccountEntryPointRulesetVersions options.
type GetAccountEntryPointRulesetVersionsOptions struct {
	// The phase of the ruleset.
	RulesetPhase *string `json:"ruleset_phase" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the GetAccountEntryPointRulesetVersionsOptions.RulesetPhase property.
// The phase of the ruleset.
const (
	GetAccountEntryPointRulesetVersionsOptions_RulesetPhase_DdosL4                         = "ddos_l4"
	GetAccountEntryPointRulesetVersionsOptions_RulesetPhase_DdosL7                         = "ddos_l7"
	GetAccountEntryPointRulesetVersionsOptions_RulesetPhase_HttpConfigSettings             = "http_config_settings"
	GetAccountEntryPointRulesetVersionsOptions_RulesetPhase_HttpCustomErrors               = "http_custom_errors"
	GetAccountEntryPointRulesetVersionsOptions_RulesetPhase_HttpLogCustomFields            = "http_log_custom_fields"
	GetAccountEntryPointRulesetVersionsOptions_RulesetPhase_HttpRatelimit                  = "http_ratelimit"
	GetAccountEntryPointRulesetVersionsOptions_RulesetPhase_HttpRequestCacheSettings       = "http_request_cache_settings"
	GetAccountEntryPointRulesetVersionsOptions_RulesetPhase_HttpRequestDynamicRedirect     = "http_request_dynamic_redirect"
	GetAccountEntryPointRulesetVersionsOptions_RulesetPhase_HttpRequestFirewallCustom      = "http_request_firewall_custom"
	GetAccountEntryPointRulesetVersionsOptions_RulesetPhase_HttpRequestFirewallManaged     = "http_request_firewall_managed"
	GetAccountEntryPointRulesetVersionsOptions_RulesetPhase_HttpRequestLateTransform       = "http_request_late_transform"
	GetAccountEntryPointRulesetVersionsOptions_RulesetPhase_HttpRequestOrigin              = "http_request_origin"
	GetAccountEntryPointRulesetVersionsOptions_RulesetPhase_HttpRequestRedirect            = "http_request_redirect"
	GetAccountEntryPointRulesetVersionsOptions_RulesetPhase_HttpRequestSanitize            = "http_request_sanitize"
	GetAccountEntryPointRulesetVersionsOptions_RulesetPhase_HttpRequestSbfm                = "http_request_sbfm"
	GetAccountEntryPointRulesetVersionsOptions_RulesetPhase_HttpRequestSelectConfiguration = "http_request_select_configuration"
	GetAccountEntryPointRulesetVersionsOptions_RulesetPhase_HttpRequestTransform           = "http_request_transform"
	GetAccountEntryPointRulesetVersionsOptions_RulesetPhase_HttpResponseCompression        = "http_response_compression"
	GetAccountEntryPointRulesetVersionsOptions_RulesetPhase_HttpResponseFirewallManaged    = "http_response_firewall_managed"
	GetAccountEntryPointRulesetVersionsOptions_RulesetPhase_HttpResponseHeadersTransform   = "http_response_headers_transform"
)

// NewGetAccountEntryPointRulesetVersionsOptions : Instantiate GetAccountEntryPointRulesetVersionsOptions
func (*RulesetsV1) NewGetAccountEntryPointRulesetVersionsOptions(rulesetPhase string) *GetAccountEntryPointRulesetVersionsOptions {
	return &GetAccountEntryPointRulesetVersionsOptions{
		RulesetPhase: core.StringPtr(rulesetPhase),
	}
}

// SetRulesetPhase : Allow user to set RulesetPhase
func (_options *GetAccountEntryPointRulesetVersionsOptions) SetRulesetPhase(rulesetPhase string) *GetAccountEntryPointRulesetVersionsOptions {
	_options.RulesetPhase = core.StringPtr(rulesetPhase)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetAccountEntryPointRulesetVersionsOptions) SetHeaders(param map[string]string) *GetAccountEntryPointRulesetVersionsOptions {
	options.Headers = param
	return options
}

// GetAccountEntrypointRulesetOptions : The GetAccountEntrypointRuleset options.
type GetAccountEntrypointRulesetOptions struct {
	// The phase of the ruleset.
	RulesetPhase *string `json:"ruleset_phase" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the GetAccountEntrypointRulesetOptions.RulesetPhase property.
// The phase of the ruleset.
const (
	GetAccountEntrypointRulesetOptions_RulesetPhase_DdosL4                         = "ddos_l4"
	GetAccountEntrypointRulesetOptions_RulesetPhase_DdosL7                         = "ddos_l7"
	GetAccountEntrypointRulesetOptions_RulesetPhase_HttpConfigSettings             = "http_config_settings"
	GetAccountEntrypointRulesetOptions_RulesetPhase_HttpCustomErrors               = "http_custom_errors"
	GetAccountEntrypointRulesetOptions_RulesetPhase_HttpLogCustomFields            = "http_log_custom_fields"
	GetAccountEntrypointRulesetOptions_RulesetPhase_HttpRatelimit                  = "http_ratelimit"
	GetAccountEntrypointRulesetOptions_RulesetPhase_HttpRequestCacheSettings       = "http_request_cache_settings"
	GetAccountEntrypointRulesetOptions_RulesetPhase_HttpRequestDynamicRedirect     = "http_request_dynamic_redirect"
	GetAccountEntrypointRulesetOptions_RulesetPhase_HttpRequestFirewallCustom      = "http_request_firewall_custom"
	GetAccountEntrypointRulesetOptions_RulesetPhase_HttpRequestFirewallManaged     = "http_request_firewall_managed"
	GetAccountEntrypointRulesetOptions_RulesetPhase_HttpRequestLateTransform       = "http_request_late_transform"
	GetAccountEntrypointRulesetOptions_RulesetPhase_HttpRequestOrigin              = "http_request_origin"
	GetAccountEntrypointRulesetOptions_RulesetPhase_HttpRequestRedirect            = "http_request_redirect"
	GetAccountEntrypointRulesetOptions_RulesetPhase_HttpRequestSanitize            = "http_request_sanitize"
	GetAccountEntrypointRulesetOptions_RulesetPhase_HttpRequestSbfm                = "http_request_sbfm"
	GetAccountEntrypointRulesetOptions_RulesetPhase_HttpRequestSelectConfiguration = "http_request_select_configuration"
	GetAccountEntrypointRulesetOptions_RulesetPhase_HttpRequestTransform           = "http_request_transform"
	GetAccountEntrypointRulesetOptions_RulesetPhase_HttpResponseCompression        = "http_response_compression"
	GetAccountEntrypointRulesetOptions_RulesetPhase_HttpResponseFirewallManaged    = "http_response_firewall_managed"
	GetAccountEntrypointRulesetOptions_RulesetPhase_HttpResponseHeadersTransform   = "http_response_headers_transform"
)

// NewGetAccountEntrypointRulesetOptions : Instantiate GetAccountEntrypointRulesetOptions
func (*RulesetsV1) NewGetAccountEntrypointRulesetOptions(rulesetPhase string) *GetAccountEntrypointRulesetOptions {
	return &GetAccountEntrypointRulesetOptions{
		RulesetPhase: core.StringPtr(rulesetPhase),
	}
}

// SetRulesetPhase : Allow user to set RulesetPhase
func (_options *GetAccountEntrypointRulesetOptions) SetRulesetPhase(rulesetPhase string) *GetAccountEntrypointRulesetOptions {
	_options.RulesetPhase = core.StringPtr(rulesetPhase)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetAccountEntrypointRulesetOptions) SetHeaders(param map[string]string) *GetAccountEntrypointRulesetOptions {
	options.Headers = param
	return options
}

// GetAccountRulesetOptions : The GetAccountRuleset options.
type GetAccountRulesetOptions struct {
	// ID of a specific ruleset.
	RulesetID *string `json:"ruleset_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetAccountRulesetOptions : Instantiate GetAccountRulesetOptions
func (*RulesetsV1) NewGetAccountRulesetOptions(rulesetID string) *GetAccountRulesetOptions {
	return &GetAccountRulesetOptions{
		RulesetID: core.StringPtr(rulesetID),
	}
}

// SetRulesetID : Allow user to set RulesetID
func (_options *GetAccountRulesetOptions) SetRulesetID(rulesetID string) *GetAccountRulesetOptions {
	_options.RulesetID = core.StringPtr(rulesetID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetAccountRulesetOptions) SetHeaders(param map[string]string) *GetAccountRulesetOptions {
	options.Headers = param
	return options
}

// GetAccountRulesetVersionByTagOptions : The GetAccountRulesetVersionByTag options.
type GetAccountRulesetVersionByTagOptions struct {
	// ID of a specific ruleset.
	RulesetID *string `json:"ruleset_id" validate:"required,ne="`

	// The version of the ruleset.
	RulesetVersion *string `json:"ruleset_version" validate:"required,ne="`

	// A category of the rule.
	RuleTag *string `json:"rule_tag" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetAccountRulesetVersionByTagOptions : Instantiate GetAccountRulesetVersionByTagOptions
func (*RulesetsV1) NewGetAccountRulesetVersionByTagOptions(rulesetID string, rulesetVersion string, ruleTag string) *GetAccountRulesetVersionByTagOptions {
	return &GetAccountRulesetVersionByTagOptions{
		RulesetID:      core.StringPtr(rulesetID),
		RulesetVersion: core.StringPtr(rulesetVersion),
		RuleTag:        core.StringPtr(ruleTag),
	}
}

// SetRulesetID : Allow user to set RulesetID
func (_options *GetAccountRulesetVersionByTagOptions) SetRulesetID(rulesetID string) *GetAccountRulesetVersionByTagOptions {
	_options.RulesetID = core.StringPtr(rulesetID)
	return _options
}

// SetRulesetVersion : Allow user to set RulesetVersion
func (_options *GetAccountRulesetVersionByTagOptions) SetRulesetVersion(rulesetVersion string) *GetAccountRulesetVersionByTagOptions {
	_options.RulesetVersion = core.StringPtr(rulesetVersion)
	return _options
}

// SetRuleTag : Allow user to set RuleTag
func (_options *GetAccountRulesetVersionByTagOptions) SetRuleTag(ruleTag string) *GetAccountRulesetVersionByTagOptions {
	_options.RuleTag = core.StringPtr(ruleTag)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetAccountRulesetVersionByTagOptions) SetHeaders(param map[string]string) *GetAccountRulesetVersionByTagOptions {
	options.Headers = param
	return options
}

// GetAccountRulesetVersionOptions : The GetAccountRulesetVersion options.
type GetAccountRulesetVersionOptions struct {
	// ID of a specific ruleset.
	RulesetID *string `json:"ruleset_id" validate:"required,ne="`

	// The version of the ruleset.
	RulesetVersion *string `json:"ruleset_version" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetAccountRulesetVersionOptions : Instantiate GetAccountRulesetVersionOptions
func (*RulesetsV1) NewGetAccountRulesetVersionOptions(rulesetID string, rulesetVersion string) *GetAccountRulesetVersionOptions {
	return &GetAccountRulesetVersionOptions{
		RulesetID:      core.StringPtr(rulesetID),
		RulesetVersion: core.StringPtr(rulesetVersion),
	}
}

// SetRulesetID : Allow user to set RulesetID
func (_options *GetAccountRulesetVersionOptions) SetRulesetID(rulesetID string) *GetAccountRulesetVersionOptions {
	_options.RulesetID = core.StringPtr(rulesetID)
	return _options
}

// SetRulesetVersion : Allow user to set RulesetVersion
func (_options *GetAccountRulesetVersionOptions) SetRulesetVersion(rulesetVersion string) *GetAccountRulesetVersionOptions {
	_options.RulesetVersion = core.StringPtr(rulesetVersion)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetAccountRulesetVersionOptions) SetHeaders(param map[string]string) *GetAccountRulesetVersionOptions {
	options.Headers = param
	return options
}

// GetAccountRulesetVersionsOptions : The GetAccountRulesetVersions options.
type GetAccountRulesetVersionsOptions struct {
	// ID of a specific ruleset.
	RulesetID *string `json:"ruleset_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetAccountRulesetVersionsOptions : Instantiate GetAccountRulesetVersionsOptions
func (*RulesetsV1) NewGetAccountRulesetVersionsOptions(rulesetID string) *GetAccountRulesetVersionsOptions {
	return &GetAccountRulesetVersionsOptions{
		RulesetID: core.StringPtr(rulesetID),
	}
}

// SetRulesetID : Allow user to set RulesetID
func (_options *GetAccountRulesetVersionsOptions) SetRulesetID(rulesetID string) *GetAccountRulesetVersionsOptions {
	_options.RulesetID = core.StringPtr(rulesetID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetAccountRulesetVersionsOptions) SetHeaders(param map[string]string) *GetAccountRulesetVersionsOptions {
	options.Headers = param
	return options
}

// GetAccountRulesetsOptions : The GetAccountRulesets options.
type GetAccountRulesetsOptions struct {

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetAccountRulesetsOptions : Instantiate GetAccountRulesetsOptions
func (*RulesetsV1) NewGetAccountRulesetsOptions() *GetAccountRulesetsOptions {
	return &GetAccountRulesetsOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *GetAccountRulesetsOptions) SetHeaders(param map[string]string) *GetAccountRulesetsOptions {
	options.Headers = param
	return options
}

// GetZoneEntryPointRulesetVersionOptions : The GetZoneEntryPointRulesetVersion options.
type GetZoneEntryPointRulesetVersionOptions struct {
	// The phase of the ruleset.
	RulesetPhase *string `json:"ruleset_phase" validate:"required,ne="`

	// The version of the ruleset.
	RulesetVersion *string `json:"ruleset_version" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the GetZoneEntryPointRulesetVersionOptions.RulesetPhase property.
// The phase of the ruleset.
const (
	GetZoneEntryPointRulesetVersionOptions_RulesetPhase_DdosL4                         = "ddos_l4"
	GetZoneEntryPointRulesetVersionOptions_RulesetPhase_DdosL7                         = "ddos_l7"
	GetZoneEntryPointRulesetVersionOptions_RulesetPhase_HttpConfigSettings             = "http_config_settings"
	GetZoneEntryPointRulesetVersionOptions_RulesetPhase_HttpCustomErrors               = "http_custom_errors"
	GetZoneEntryPointRulesetVersionOptions_RulesetPhase_HttpLogCustomFields            = "http_log_custom_fields"
	GetZoneEntryPointRulesetVersionOptions_RulesetPhase_HttpRatelimit                  = "http_ratelimit"
	GetZoneEntryPointRulesetVersionOptions_RulesetPhase_HttpRequestCacheSettings       = "http_request_cache_settings"
	GetZoneEntryPointRulesetVersionOptions_RulesetPhase_HttpRequestDynamicRedirect     = "http_request_dynamic_redirect"
	GetZoneEntryPointRulesetVersionOptions_RulesetPhase_HttpRequestFirewallCustom      = "http_request_firewall_custom"
	GetZoneEntryPointRulesetVersionOptions_RulesetPhase_HttpRequestFirewallManaged     = "http_request_firewall_managed"
	GetZoneEntryPointRulesetVersionOptions_RulesetPhase_HttpRequestLateTransform       = "http_request_late_transform"
	GetZoneEntryPointRulesetVersionOptions_RulesetPhase_HttpRequestOrigin              = "http_request_origin"
	GetZoneEntryPointRulesetVersionOptions_RulesetPhase_HttpRequestRedirect            = "http_request_redirect"
	GetZoneEntryPointRulesetVersionOptions_RulesetPhase_HttpRequestSanitize            = "http_request_sanitize"
	GetZoneEntryPointRulesetVersionOptions_RulesetPhase_HttpRequestSbfm                = "http_request_sbfm"
	GetZoneEntryPointRulesetVersionOptions_RulesetPhase_HttpRequestSelectConfiguration = "http_request_select_configuration"
	GetZoneEntryPointRulesetVersionOptions_RulesetPhase_HttpRequestTransform           = "http_request_transform"
	GetZoneEntryPointRulesetVersionOptions_RulesetPhase_HttpResponseCompression        = "http_response_compression"
	GetZoneEntryPointRulesetVersionOptions_RulesetPhase_HttpResponseFirewallManaged    = "http_response_firewall_managed"
	GetZoneEntryPointRulesetVersionOptions_RulesetPhase_HttpResponseHeadersTransform   = "http_response_headers_transform"
)

// NewGetZoneEntryPointRulesetVersionOptions : Instantiate GetZoneEntryPointRulesetVersionOptions
func (*RulesetsV1) NewGetZoneEntryPointRulesetVersionOptions(rulesetPhase string, rulesetVersion string) *GetZoneEntryPointRulesetVersionOptions {
	return &GetZoneEntryPointRulesetVersionOptions{
		RulesetPhase:   core.StringPtr(rulesetPhase),
		RulesetVersion: core.StringPtr(rulesetVersion),
	}
}

// SetRulesetPhase : Allow user to set RulesetPhase
func (_options *GetZoneEntryPointRulesetVersionOptions) SetRulesetPhase(rulesetPhase string) *GetZoneEntryPointRulesetVersionOptions {
	_options.RulesetPhase = core.StringPtr(rulesetPhase)
	return _options
}

// SetRulesetVersion : Allow user to set RulesetVersion
func (_options *GetZoneEntryPointRulesetVersionOptions) SetRulesetVersion(rulesetVersion string) *GetZoneEntryPointRulesetVersionOptions {
	_options.RulesetVersion = core.StringPtr(rulesetVersion)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetZoneEntryPointRulesetVersionOptions) SetHeaders(param map[string]string) *GetZoneEntryPointRulesetVersionOptions {
	options.Headers = param
	return options
}

// GetZoneEntryPointRulesetVersionsOptions : The GetZoneEntryPointRulesetVersions options.
type GetZoneEntryPointRulesetVersionsOptions struct {
	// The phase of the ruleset.
	RulesetPhase *string `json:"ruleset_phase" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the GetZoneEntryPointRulesetVersionsOptions.RulesetPhase property.
// The phase of the ruleset.
const (
	GetZoneEntryPointRulesetVersionsOptions_RulesetPhase_DdosL4                         = "ddos_l4"
	GetZoneEntryPointRulesetVersionsOptions_RulesetPhase_DdosL7                         = "ddos_l7"
	GetZoneEntryPointRulesetVersionsOptions_RulesetPhase_HttpConfigSettings             = "http_config_settings"
	GetZoneEntryPointRulesetVersionsOptions_RulesetPhase_HttpCustomErrors               = "http_custom_errors"
	GetZoneEntryPointRulesetVersionsOptions_RulesetPhase_HttpLogCustomFields            = "http_log_custom_fields"
	GetZoneEntryPointRulesetVersionsOptions_RulesetPhase_HttpRatelimit                  = "http_ratelimit"
	GetZoneEntryPointRulesetVersionsOptions_RulesetPhase_HttpRequestCacheSettings       = "http_request_cache_settings"
	GetZoneEntryPointRulesetVersionsOptions_RulesetPhase_HttpRequestDynamicRedirect     = "http_request_dynamic_redirect"
	GetZoneEntryPointRulesetVersionsOptions_RulesetPhase_HttpRequestFirewallCustom      = "http_request_firewall_custom"
	GetZoneEntryPointRulesetVersionsOptions_RulesetPhase_HttpRequestFirewallManaged     = "http_request_firewall_managed"
	GetZoneEntryPointRulesetVersionsOptions_RulesetPhase_HttpRequestLateTransform       = "http_request_late_transform"
	GetZoneEntryPointRulesetVersionsOptions_RulesetPhase_HttpRequestOrigin              = "http_request_origin"
	GetZoneEntryPointRulesetVersionsOptions_RulesetPhase_HttpRequestRedirect            = "http_request_redirect"
	GetZoneEntryPointRulesetVersionsOptions_RulesetPhase_HttpRequestSanitize            = "http_request_sanitize"
	GetZoneEntryPointRulesetVersionsOptions_RulesetPhase_HttpRequestSbfm                = "http_request_sbfm"
	GetZoneEntryPointRulesetVersionsOptions_RulesetPhase_HttpRequestSelectConfiguration = "http_request_select_configuration"
	GetZoneEntryPointRulesetVersionsOptions_RulesetPhase_HttpRequestTransform           = "http_request_transform"
	GetZoneEntryPointRulesetVersionsOptions_RulesetPhase_HttpResponseCompression        = "http_response_compression"
	GetZoneEntryPointRulesetVersionsOptions_RulesetPhase_HttpResponseFirewallManaged    = "http_response_firewall_managed"
	GetZoneEntryPointRulesetVersionsOptions_RulesetPhase_HttpResponseHeadersTransform   = "http_response_headers_transform"
)

// NewGetZoneEntryPointRulesetVersionsOptions : Instantiate GetZoneEntryPointRulesetVersionsOptions
func (*RulesetsV1) NewGetZoneEntryPointRulesetVersionsOptions(rulesetPhase string) *GetZoneEntryPointRulesetVersionsOptions {
	return &GetZoneEntryPointRulesetVersionsOptions{
		RulesetPhase: core.StringPtr(rulesetPhase),
	}
}

// SetRulesetPhase : Allow user to set RulesetPhase
func (_options *GetZoneEntryPointRulesetVersionsOptions) SetRulesetPhase(rulesetPhase string) *GetZoneEntryPointRulesetVersionsOptions {
	_options.RulesetPhase = core.StringPtr(rulesetPhase)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetZoneEntryPointRulesetVersionsOptions) SetHeaders(param map[string]string) *GetZoneEntryPointRulesetVersionsOptions {
	options.Headers = param
	return options
}

// GetZoneEntrypointRulesetOptions : The GetZoneEntrypointRuleset options.
type GetZoneEntrypointRulesetOptions struct {
	// The phase of the ruleset.
	RulesetPhase *string `json:"ruleset_phase" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the GetZoneEntrypointRulesetOptions.RulesetPhase property.
// The phase of the ruleset.
const (
	GetZoneEntrypointRulesetOptions_RulesetPhase_DdosL4                         = "ddos_l4"
	GetZoneEntrypointRulesetOptions_RulesetPhase_DdosL7                         = "ddos_l7"
	GetZoneEntrypointRulesetOptions_RulesetPhase_HttpConfigSettings             = "http_config_settings"
	GetZoneEntrypointRulesetOptions_RulesetPhase_HttpCustomErrors               = "http_custom_errors"
	GetZoneEntrypointRulesetOptions_RulesetPhase_HttpLogCustomFields            = "http_log_custom_fields"
	GetZoneEntrypointRulesetOptions_RulesetPhase_HttpRatelimit                  = "http_ratelimit"
	GetZoneEntrypointRulesetOptions_RulesetPhase_HttpRequestCacheSettings       = "http_request_cache_settings"
	GetZoneEntrypointRulesetOptions_RulesetPhase_HttpRequestDynamicRedirect     = "http_request_dynamic_redirect"
	GetZoneEntrypointRulesetOptions_RulesetPhase_HttpRequestFirewallCustom      = "http_request_firewall_custom"
	GetZoneEntrypointRulesetOptions_RulesetPhase_HttpRequestFirewallManaged     = "http_request_firewall_managed"
	GetZoneEntrypointRulesetOptions_RulesetPhase_HttpRequestLateTransform       = "http_request_late_transform"
	GetZoneEntrypointRulesetOptions_RulesetPhase_HttpRequestOrigin              = "http_request_origin"
	GetZoneEntrypointRulesetOptions_RulesetPhase_HttpRequestRedirect            = "http_request_redirect"
	GetZoneEntrypointRulesetOptions_RulesetPhase_HttpRequestSanitize            = "http_request_sanitize"
	GetZoneEntrypointRulesetOptions_RulesetPhase_HttpRequestSbfm                = "http_request_sbfm"
	GetZoneEntrypointRulesetOptions_RulesetPhase_HttpRequestSelectConfiguration = "http_request_select_configuration"
	GetZoneEntrypointRulesetOptions_RulesetPhase_HttpRequestTransform           = "http_request_transform"
	GetZoneEntrypointRulesetOptions_RulesetPhase_HttpResponseCompression        = "http_response_compression"
	GetZoneEntrypointRulesetOptions_RulesetPhase_HttpResponseFirewallManaged    = "http_response_firewall_managed"
	GetZoneEntrypointRulesetOptions_RulesetPhase_HttpResponseHeadersTransform   = "http_response_headers_transform"
)

// NewGetZoneEntrypointRulesetOptions : Instantiate GetZoneEntrypointRulesetOptions
func (*RulesetsV1) NewGetZoneEntrypointRulesetOptions(rulesetPhase string) *GetZoneEntrypointRulesetOptions {
	return &GetZoneEntrypointRulesetOptions{
		RulesetPhase: core.StringPtr(rulesetPhase),
	}
}

// SetRulesetPhase : Allow user to set RulesetPhase
func (_options *GetZoneEntrypointRulesetOptions) SetRulesetPhase(rulesetPhase string) *GetZoneEntrypointRulesetOptions {
	_options.RulesetPhase = core.StringPtr(rulesetPhase)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetZoneEntrypointRulesetOptions) SetHeaders(param map[string]string) *GetZoneEntrypointRulesetOptions {
	options.Headers = param
	return options
}

// GetZoneRulesetOptions : The GetZoneRuleset options.
type GetZoneRulesetOptions struct {
	// ID of a specific ruleset.
	RulesetID *string `json:"ruleset_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetZoneRulesetOptions : Instantiate GetZoneRulesetOptions
func (*RulesetsV1) NewGetZoneRulesetOptions(rulesetID string) *GetZoneRulesetOptions {
	return &GetZoneRulesetOptions{
		RulesetID: core.StringPtr(rulesetID),
	}
}

// SetRulesetID : Allow user to set RulesetID
func (_options *GetZoneRulesetOptions) SetRulesetID(rulesetID string) *GetZoneRulesetOptions {
	_options.RulesetID = core.StringPtr(rulesetID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetZoneRulesetOptions) SetHeaders(param map[string]string) *GetZoneRulesetOptions {
	options.Headers = param
	return options
}

// GetZoneRulesetVersionOptions : The GetZoneRulesetVersion options.
type GetZoneRulesetVersionOptions struct {
	// ID of a specific ruleset.
	RulesetID *string `json:"ruleset_id" validate:"required,ne="`

	// The version of the ruleset.
	RulesetVersion *string `json:"ruleset_version" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetZoneRulesetVersionOptions : Instantiate GetZoneRulesetVersionOptions
func (*RulesetsV1) NewGetZoneRulesetVersionOptions(rulesetID string, rulesetVersion string) *GetZoneRulesetVersionOptions {
	return &GetZoneRulesetVersionOptions{
		RulesetID:      core.StringPtr(rulesetID),
		RulesetVersion: core.StringPtr(rulesetVersion),
	}
}

// SetRulesetID : Allow user to set RulesetID
func (_options *GetZoneRulesetVersionOptions) SetRulesetID(rulesetID string) *GetZoneRulesetVersionOptions {
	_options.RulesetID = core.StringPtr(rulesetID)
	return _options
}

// SetRulesetVersion : Allow user to set RulesetVersion
func (_options *GetZoneRulesetVersionOptions) SetRulesetVersion(rulesetVersion string) *GetZoneRulesetVersionOptions {
	_options.RulesetVersion = core.StringPtr(rulesetVersion)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetZoneRulesetVersionOptions) SetHeaders(param map[string]string) *GetZoneRulesetVersionOptions {
	options.Headers = param
	return options
}

// GetZoneRulesetVersionsOptions : The GetZoneRulesetVersions options.
type GetZoneRulesetVersionsOptions struct {
	// ID of a specific ruleset.
	RulesetID *string `json:"ruleset_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetZoneRulesetVersionsOptions : Instantiate GetZoneRulesetVersionsOptions
func (*RulesetsV1) NewGetZoneRulesetVersionsOptions(rulesetID string) *GetZoneRulesetVersionsOptions {
	return &GetZoneRulesetVersionsOptions{
		RulesetID: core.StringPtr(rulesetID),
	}
}

// SetRulesetID : Allow user to set RulesetID
func (_options *GetZoneRulesetVersionsOptions) SetRulesetID(rulesetID string) *GetZoneRulesetVersionsOptions {
	_options.RulesetID = core.StringPtr(rulesetID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetZoneRulesetVersionsOptions) SetHeaders(param map[string]string) *GetZoneRulesetVersionsOptions {
	options.Headers = param
	return options
}

// GetZoneRulesetsOptions : The GetZoneRulesets options.
type GetZoneRulesetsOptions struct {

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetZoneRulesetsOptions : Instantiate GetZoneRulesetsOptions
func (*RulesetsV1) NewGetZoneRulesetsOptions() *GetZoneRulesetsOptions {
	return &GetZoneRulesetsOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *GetZoneRulesetsOptions) SetHeaders(param map[string]string) *GetZoneRulesetsOptions {
	options.Headers = param
	return options
}

// MessageSource : The source of this message.
type MessageSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer *string `json:"pointer" validate:"required"`
}

// UnmarshalMessageSource unmarshals an instance of MessageSource from the specified map of raw messages.
func UnmarshalMessageSource(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(MessageSource)
	err = core.UnmarshalPrimitive(m, "pointer", &obj.Pointer)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// UpdateAccountEntrypointRulesetOptions : The UpdateAccountEntrypointRuleset options.
type UpdateAccountEntrypointRulesetOptions struct {
	// The phase of the ruleset.
	RulesetPhase *string `json:"ruleset_phase" validate:"required,ne="`

	// description of the ruleset.
	Description *string `json:"description,omitempty"`

	Kind *string `json:"kind,omitempty"`

	// human readable name of the ruleset.
	Name *string `json:"name,omitempty"`

	// The phase of the ruleset.
	Phase *string `json:"phase,omitempty"`

	Rules []RuleCreate `json:"rules,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the UpdateAccountEntrypointRulesetOptions.RulesetPhase property.
// The phase of the ruleset.
const (
	UpdateAccountEntrypointRulesetOptions_RulesetPhase_DdosL4                         = "ddos_l4"
	UpdateAccountEntrypointRulesetOptions_RulesetPhase_DdosL7                         = "ddos_l7"
	UpdateAccountEntrypointRulesetOptions_RulesetPhase_HttpConfigSettings             = "http_config_settings"
	UpdateAccountEntrypointRulesetOptions_RulesetPhase_HttpCustomErrors               = "http_custom_errors"
	UpdateAccountEntrypointRulesetOptions_RulesetPhase_HttpLogCustomFields            = "http_log_custom_fields"
	UpdateAccountEntrypointRulesetOptions_RulesetPhase_HttpRatelimit                  = "http_ratelimit"
	UpdateAccountEntrypointRulesetOptions_RulesetPhase_HttpRequestCacheSettings       = "http_request_cache_settings"
	UpdateAccountEntrypointRulesetOptions_RulesetPhase_HttpRequestDynamicRedirect     = "http_request_dynamic_redirect"
	UpdateAccountEntrypointRulesetOptions_RulesetPhase_HttpRequestFirewallCustom      = "http_request_firewall_custom"
	UpdateAccountEntrypointRulesetOptions_RulesetPhase_HttpRequestFirewallManaged     = "http_request_firewall_managed"
	UpdateAccountEntrypointRulesetOptions_RulesetPhase_HttpRequestLateTransform       = "http_request_late_transform"
	UpdateAccountEntrypointRulesetOptions_RulesetPhase_HttpRequestOrigin              = "http_request_origin"
	UpdateAccountEntrypointRulesetOptions_RulesetPhase_HttpRequestRedirect            = "http_request_redirect"
	UpdateAccountEntrypointRulesetOptions_RulesetPhase_HttpRequestSanitize            = "http_request_sanitize"
	UpdateAccountEntrypointRulesetOptions_RulesetPhase_HttpRequestSbfm                = "http_request_sbfm"
	UpdateAccountEntrypointRulesetOptions_RulesetPhase_HttpRequestSelectConfiguration = "http_request_select_configuration"
	UpdateAccountEntrypointRulesetOptions_RulesetPhase_HttpRequestTransform           = "http_request_transform"
	UpdateAccountEntrypointRulesetOptions_RulesetPhase_HttpResponseCompression        = "http_response_compression"
	UpdateAccountEntrypointRulesetOptions_RulesetPhase_HttpResponseFirewallManaged    = "http_response_firewall_managed"
	UpdateAccountEntrypointRulesetOptions_RulesetPhase_HttpResponseHeadersTransform   = "http_response_headers_transform"
)

// Constants associated with the UpdateAccountEntrypointRulesetOptions.Kind property.
const (
	UpdateAccountEntrypointRulesetOptions_Kind_Custom  = "custom"
	UpdateAccountEntrypointRulesetOptions_Kind_Managed = "managed"
	UpdateAccountEntrypointRulesetOptions_Kind_Root    = "root"
	UpdateAccountEntrypointRulesetOptions_Kind_Zone    = "zone"
)

// Constants associated with the UpdateAccountEntrypointRulesetOptions.Phase property.
// The phase of the ruleset.
const (
	UpdateAccountEntrypointRulesetOptions_Phase_DdosL4                         = "ddos_l4"
	UpdateAccountEntrypointRulesetOptions_Phase_DdosL7                         = "ddos_l7"
	UpdateAccountEntrypointRulesetOptions_Phase_HttpConfigSettings             = "http_config_settings"
	UpdateAccountEntrypointRulesetOptions_Phase_HttpCustomErrors               = "http_custom_errors"
	UpdateAccountEntrypointRulesetOptions_Phase_HttpLogCustomFields            = "http_log_custom_fields"
	UpdateAccountEntrypointRulesetOptions_Phase_HttpRatelimit                  = "http_ratelimit"
	UpdateAccountEntrypointRulesetOptions_Phase_HttpRequestCacheSettings       = "http_request_cache_settings"
	UpdateAccountEntrypointRulesetOptions_Phase_HttpRequestDynamicRedirect     = "http_request_dynamic_redirect"
	UpdateAccountEntrypointRulesetOptions_Phase_HttpRequestFirewallCustom      = "http_request_firewall_custom"
	UpdateAccountEntrypointRulesetOptions_Phase_HttpRequestFirewallManaged     = "http_request_firewall_managed"
	UpdateAccountEntrypointRulesetOptions_Phase_HttpRequestLateTransform       = "http_request_late_transform"
	UpdateAccountEntrypointRulesetOptions_Phase_HttpRequestOrigin              = "http_request_origin"
	UpdateAccountEntrypointRulesetOptions_Phase_HttpRequestRedirect            = "http_request_redirect"
	UpdateAccountEntrypointRulesetOptions_Phase_HttpRequestSanitize            = "http_request_sanitize"
	UpdateAccountEntrypointRulesetOptions_Phase_HttpRequestSbfm                = "http_request_sbfm"
	UpdateAccountEntrypointRulesetOptions_Phase_HttpRequestSelectConfiguration = "http_request_select_configuration"
	UpdateAccountEntrypointRulesetOptions_Phase_HttpRequestTransform           = "http_request_transform"
	UpdateAccountEntrypointRulesetOptions_Phase_HttpResponseCompression        = "http_response_compression"
	UpdateAccountEntrypointRulesetOptions_Phase_HttpResponseFirewallManaged    = "http_response_firewall_managed"
	UpdateAccountEntrypointRulesetOptions_Phase_HttpResponseHeadersTransform   = "http_response_headers_transform"
)

// NewUpdateAccountEntrypointRulesetOptions : Instantiate UpdateAccountEntrypointRulesetOptions
func (*RulesetsV1) NewUpdateAccountEntrypointRulesetOptions(rulesetPhase string) *UpdateAccountEntrypointRulesetOptions {
	return &UpdateAccountEntrypointRulesetOptions{
		RulesetPhase: core.StringPtr(rulesetPhase),
	}
}

// SetRulesetPhase : Allow user to set RulesetPhase
func (_options *UpdateAccountEntrypointRulesetOptions) SetRulesetPhase(rulesetPhase string) *UpdateAccountEntrypointRulesetOptions {
	_options.RulesetPhase = core.StringPtr(rulesetPhase)
	return _options
}

// SetDescription : Allow user to set Description
func (_options *UpdateAccountEntrypointRulesetOptions) SetDescription(description string) *UpdateAccountEntrypointRulesetOptions {
	_options.Description = core.StringPtr(description)
	return _options
}

// SetKind : Allow user to set Kind
func (_options *UpdateAccountEntrypointRulesetOptions) SetKind(kind string) *UpdateAccountEntrypointRulesetOptions {
	_options.Kind = core.StringPtr(kind)
	return _options
}

// SetName : Allow user to set Name
func (_options *UpdateAccountEntrypointRulesetOptions) SetName(name string) *UpdateAccountEntrypointRulesetOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetPhase : Allow user to set Phase
func (_options *UpdateAccountEntrypointRulesetOptions) SetPhase(phase string) *UpdateAccountEntrypointRulesetOptions {
	_options.Phase = core.StringPtr(phase)
	return _options
}

// SetRules : Allow user to set Rules
func (_options *UpdateAccountEntrypointRulesetOptions) SetRules(rules []RuleCreate) *UpdateAccountEntrypointRulesetOptions {
	_options.Rules = rules
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateAccountEntrypointRulesetOptions) SetHeaders(param map[string]string) *UpdateAccountEntrypointRulesetOptions {
	options.Headers = param
	return options
}

// UpdateAccountRulesetOptions : The UpdateAccountRuleset options.
type UpdateAccountRulesetOptions struct {
	// ID of a specific ruleset.
	RulesetID *string `json:"ruleset_id" validate:"required,ne="`

	// description of the ruleset.
	Description *string `json:"description,omitempty"`

	Kind *string `json:"kind,omitempty"`

	// human readable name of the ruleset.
	Name *string `json:"name,omitempty"`

	// The phase of the ruleset.
	Phase *string `json:"phase,omitempty"`

	Rules []RuleCreate `json:"rules,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the UpdateAccountRulesetOptions.Kind property.
const (
	UpdateAccountRulesetOptions_Kind_Custom  = "custom"
	UpdateAccountRulesetOptions_Kind_Managed = "managed"
	UpdateAccountRulesetOptions_Kind_Root    = "root"
	UpdateAccountRulesetOptions_Kind_Zone    = "zone"
)

// Constants associated with the UpdateAccountRulesetOptions.Phase property.
// The phase of the ruleset.
const (
	UpdateAccountRulesetOptions_Phase_DdosL4                         = "ddos_l4"
	UpdateAccountRulesetOptions_Phase_DdosL7                         = "ddos_l7"
	UpdateAccountRulesetOptions_Phase_HttpConfigSettings             = "http_config_settings"
	UpdateAccountRulesetOptions_Phase_HttpCustomErrors               = "http_custom_errors"
	UpdateAccountRulesetOptions_Phase_HttpLogCustomFields            = "http_log_custom_fields"
	UpdateAccountRulesetOptions_Phase_HttpRatelimit                  = "http_ratelimit"
	UpdateAccountRulesetOptions_Phase_HttpRequestCacheSettings       = "http_request_cache_settings"
	UpdateAccountRulesetOptions_Phase_HttpRequestDynamicRedirect     = "http_request_dynamic_redirect"
	UpdateAccountRulesetOptions_Phase_HttpRequestFirewallCustom      = "http_request_firewall_custom"
	UpdateAccountRulesetOptions_Phase_HttpRequestFirewallManaged     = "http_request_firewall_managed"
	UpdateAccountRulesetOptions_Phase_HttpRequestLateTransform       = "http_request_late_transform"
	UpdateAccountRulesetOptions_Phase_HttpRequestOrigin              = "http_request_origin"
	UpdateAccountRulesetOptions_Phase_HttpRequestRedirect            = "http_request_redirect"
	UpdateAccountRulesetOptions_Phase_HttpRequestSanitize            = "http_request_sanitize"
	UpdateAccountRulesetOptions_Phase_HttpRequestSbfm                = "http_request_sbfm"
	UpdateAccountRulesetOptions_Phase_HttpRequestSelectConfiguration = "http_request_select_configuration"
	UpdateAccountRulesetOptions_Phase_HttpRequestTransform           = "http_request_transform"
	UpdateAccountRulesetOptions_Phase_HttpResponseCompression        = "http_response_compression"
	UpdateAccountRulesetOptions_Phase_HttpResponseFirewallManaged    = "http_response_firewall_managed"
	UpdateAccountRulesetOptions_Phase_HttpResponseHeadersTransform   = "http_response_headers_transform"
)

// NewUpdateAccountRulesetOptions : Instantiate UpdateAccountRulesetOptions
func (*RulesetsV1) NewUpdateAccountRulesetOptions(rulesetID string) *UpdateAccountRulesetOptions {
	return &UpdateAccountRulesetOptions{
		RulesetID: core.StringPtr(rulesetID),
	}
}

// SetRulesetID : Allow user to set RulesetID
func (_options *UpdateAccountRulesetOptions) SetRulesetID(rulesetID string) *UpdateAccountRulesetOptions {
	_options.RulesetID = core.StringPtr(rulesetID)
	return _options
}

// SetDescription : Allow user to set Description
func (_options *UpdateAccountRulesetOptions) SetDescription(description string) *UpdateAccountRulesetOptions {
	_options.Description = core.StringPtr(description)
	return _options
}

// SetKind : Allow user to set Kind
func (_options *UpdateAccountRulesetOptions) SetKind(kind string) *UpdateAccountRulesetOptions {
	_options.Kind = core.StringPtr(kind)
	return _options
}

// SetName : Allow user to set Name
func (_options *UpdateAccountRulesetOptions) SetName(name string) *UpdateAccountRulesetOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetPhase : Allow user to set Phase
func (_options *UpdateAccountRulesetOptions) SetPhase(phase string) *UpdateAccountRulesetOptions {
	_options.Phase = core.StringPtr(phase)
	return _options
}

// SetRules : Allow user to set Rules
func (_options *UpdateAccountRulesetOptions) SetRules(rules []RuleCreate) *UpdateAccountRulesetOptions {
	_options.Rules = rules
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateAccountRulesetOptions) SetHeaders(param map[string]string) *UpdateAccountRulesetOptions {
	options.Headers = param
	return options
}

// UpdateAccountRulesetRuleOptions : The UpdateAccountRulesetRule options.
type UpdateAccountRulesetRuleOptions struct {
	// ID of a specific ruleset.
	RulesetID *string `json:"ruleset_id" validate:"required,ne="`

	// ID of a specific rule.
	RuleID *string `json:"rule_id" validate:"required,ne="`

	// What happens when theres a match for the rule expression.
	Action *string `json:"action,omitempty"`

	ActionParameters *ActionParameters `json:"action_parameters,omitempty"`

	Description *string `json:"description,omitempty"`

	Enabled *bool `json:"enabled,omitempty"`

	// The expression defining which traffic will match the rule.
	Expression *string `json:"expression,omitempty"`

	ID *string `json:"id,omitempty"`

	Logging *Logging `json:"logging,omitempty"`

	// The reference of the rule (the rule ID by default).
	Ref *string `json:"ref,omitempty"`

	Position *Position `json:"position,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdateAccountRulesetRuleOptions : Instantiate UpdateAccountRulesetRuleOptions
func (*RulesetsV1) NewUpdateAccountRulesetRuleOptions(rulesetID string, ruleID string) *UpdateAccountRulesetRuleOptions {
	return &UpdateAccountRulesetRuleOptions{
		RulesetID: core.StringPtr(rulesetID),
		RuleID:    core.StringPtr(ruleID),
	}
}

// SetRulesetID : Allow user to set RulesetID
func (_options *UpdateAccountRulesetRuleOptions) SetRulesetID(rulesetID string) *UpdateAccountRulesetRuleOptions {
	_options.RulesetID = core.StringPtr(rulesetID)
	return _options
}

// SetRuleID : Allow user to set RuleID
func (_options *UpdateAccountRulesetRuleOptions) SetRuleID(ruleID string) *UpdateAccountRulesetRuleOptions {
	_options.RuleID = core.StringPtr(ruleID)
	return _options
}

// SetAction : Allow user to set Action
func (_options *UpdateAccountRulesetRuleOptions) SetAction(action string) *UpdateAccountRulesetRuleOptions {
	_options.Action = core.StringPtr(action)
	return _options
}

// SetActionParameters : Allow user to set ActionParameters
func (_options *UpdateAccountRulesetRuleOptions) SetActionParameters(actionParameters *ActionParameters) *UpdateAccountRulesetRuleOptions {
	_options.ActionParameters = actionParameters
	return _options
}

// SetDescription : Allow user to set Description
func (_options *UpdateAccountRulesetRuleOptions) SetDescription(description string) *UpdateAccountRulesetRuleOptions {
	_options.Description = core.StringPtr(description)
	return _options
}

// SetEnabled : Allow user to set Enabled
func (_options *UpdateAccountRulesetRuleOptions) SetEnabled(enabled bool) *UpdateAccountRulesetRuleOptions {
	_options.Enabled = core.BoolPtr(enabled)
	return _options
}

// SetExpression : Allow user to set Expression
func (_options *UpdateAccountRulesetRuleOptions) SetExpression(expression string) *UpdateAccountRulesetRuleOptions {
	_options.Expression = core.StringPtr(expression)
	return _options
}

// SetID : Allow user to set ID
func (_options *UpdateAccountRulesetRuleOptions) SetID(id string) *UpdateAccountRulesetRuleOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetLogging : Allow user to set Logging
func (_options *UpdateAccountRulesetRuleOptions) SetLogging(logging *Logging) *UpdateAccountRulesetRuleOptions {
	_options.Logging = logging
	return _options
}

// SetRef : Allow user to set Ref
func (_options *UpdateAccountRulesetRuleOptions) SetRef(ref string) *UpdateAccountRulesetRuleOptions {
	_options.Ref = core.StringPtr(ref)
	return _options
}

// SetPosition : Allow user to set Position
func (_options *UpdateAccountRulesetRuleOptions) SetPosition(position *Position) *UpdateAccountRulesetRuleOptions {
	_options.Position = position
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateAccountRulesetRuleOptions) SetHeaders(param map[string]string) *UpdateAccountRulesetRuleOptions {
	options.Headers = param
	return options
}

// UpdateZoneEntrypointRulesetOptions : The UpdateZoneEntrypointRuleset options.
type UpdateZoneEntrypointRulesetOptions struct {
	// The phase of the ruleset.
	RulesetPhase *string `json:"ruleset_phase" validate:"required,ne="`

	// description of the ruleset.
	Description *string `json:"description,omitempty"`

	Kind *string `json:"kind,omitempty"`

	// human readable name of the ruleset.
	Name *string `json:"name,omitempty"`

	// The phase of the ruleset.
	Phase *string `json:"phase,omitempty"`

	Rules []RuleCreate `json:"rules,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the UpdateZoneEntrypointRulesetOptions.RulesetPhase property.
// The phase of the ruleset.
const (
	UpdateZoneEntrypointRulesetOptions_RulesetPhase_DdosL4                         = "ddos_l4"
	UpdateZoneEntrypointRulesetOptions_RulesetPhase_DdosL7                         = "ddos_l7"
	UpdateZoneEntrypointRulesetOptions_RulesetPhase_HttpConfigSettings             = "http_config_settings"
	UpdateZoneEntrypointRulesetOptions_RulesetPhase_HttpCustomErrors               = "http_custom_errors"
	UpdateZoneEntrypointRulesetOptions_RulesetPhase_HttpLogCustomFields            = "http_log_custom_fields"
	UpdateZoneEntrypointRulesetOptions_RulesetPhase_HttpRatelimit                  = "http_ratelimit"
	UpdateZoneEntrypointRulesetOptions_RulesetPhase_HttpRequestCacheSettings       = "http_request_cache_settings"
	UpdateZoneEntrypointRulesetOptions_RulesetPhase_HttpRequestDynamicRedirect     = "http_request_dynamic_redirect"
	UpdateZoneEntrypointRulesetOptions_RulesetPhase_HttpRequestFirewallCustom      = "http_request_firewall_custom"
	UpdateZoneEntrypointRulesetOptions_RulesetPhase_HttpRequestFirewallManaged     = "http_request_firewall_managed"
	UpdateZoneEntrypointRulesetOptions_RulesetPhase_HttpRequestLateTransform       = "http_request_late_transform"
	UpdateZoneEntrypointRulesetOptions_RulesetPhase_HttpRequestOrigin              = "http_request_origin"
	UpdateZoneEntrypointRulesetOptions_RulesetPhase_HttpRequestRedirect            = "http_request_redirect"
	UpdateZoneEntrypointRulesetOptions_RulesetPhase_HttpRequestSanitize            = "http_request_sanitize"
	UpdateZoneEntrypointRulesetOptions_RulesetPhase_HttpRequestSbfm                = "http_request_sbfm"
	UpdateZoneEntrypointRulesetOptions_RulesetPhase_HttpRequestSelectConfiguration = "http_request_select_configuration"
	UpdateZoneEntrypointRulesetOptions_RulesetPhase_HttpRequestTransform           = "http_request_transform"
	UpdateZoneEntrypointRulesetOptions_RulesetPhase_HttpResponseCompression        = "http_response_compression"
	UpdateZoneEntrypointRulesetOptions_RulesetPhase_HttpResponseFirewallManaged    = "http_response_firewall_managed"
	UpdateZoneEntrypointRulesetOptions_RulesetPhase_HttpResponseHeadersTransform   = "http_response_headers_transform"
)

// Constants associated with the UpdateZoneEntrypointRulesetOptions.Kind property.
const (
	UpdateZoneEntrypointRulesetOptions_Kind_Custom  = "custom"
	UpdateZoneEntrypointRulesetOptions_Kind_Managed = "managed"
	UpdateZoneEntrypointRulesetOptions_Kind_Root    = "root"
	UpdateZoneEntrypointRulesetOptions_Kind_Zone    = "zone"
)

// Constants associated with the UpdateZoneEntrypointRulesetOptions.Phase property.
// The phase of the ruleset.
const (
	UpdateZoneEntrypointRulesetOptions_Phase_DdosL4                         = "ddos_l4"
	UpdateZoneEntrypointRulesetOptions_Phase_DdosL7                         = "ddos_l7"
	UpdateZoneEntrypointRulesetOptions_Phase_HttpConfigSettings             = "http_config_settings"
	UpdateZoneEntrypointRulesetOptions_Phase_HttpCustomErrors               = "http_custom_errors"
	UpdateZoneEntrypointRulesetOptions_Phase_HttpLogCustomFields            = "http_log_custom_fields"
	UpdateZoneEntrypointRulesetOptions_Phase_HttpRatelimit                  = "http_ratelimit"
	UpdateZoneEntrypointRulesetOptions_Phase_HttpRequestCacheSettings       = "http_request_cache_settings"
	UpdateZoneEntrypointRulesetOptions_Phase_HttpRequestDynamicRedirect     = "http_request_dynamic_redirect"
	UpdateZoneEntrypointRulesetOptions_Phase_HttpRequestFirewallCustom      = "http_request_firewall_custom"
	UpdateZoneEntrypointRulesetOptions_Phase_HttpRequestFirewallManaged     = "http_request_firewall_managed"
	UpdateZoneEntrypointRulesetOptions_Phase_HttpRequestLateTransform       = "http_request_late_transform"
	UpdateZoneEntrypointRulesetOptions_Phase_HttpRequestOrigin              = "http_request_origin"
	UpdateZoneEntrypointRulesetOptions_Phase_HttpRequestRedirect            = "http_request_redirect"
	UpdateZoneEntrypointRulesetOptions_Phase_HttpRequestSanitize            = "http_request_sanitize"
	UpdateZoneEntrypointRulesetOptions_Phase_HttpRequestSbfm                = "http_request_sbfm"
	UpdateZoneEntrypointRulesetOptions_Phase_HttpRequestSelectConfiguration = "http_request_select_configuration"
	UpdateZoneEntrypointRulesetOptions_Phase_HttpRequestTransform           = "http_request_transform"
	UpdateZoneEntrypointRulesetOptions_Phase_HttpResponseCompression        = "http_response_compression"
	UpdateZoneEntrypointRulesetOptions_Phase_HttpResponseFirewallManaged    = "http_response_firewall_managed"
	UpdateZoneEntrypointRulesetOptions_Phase_HttpResponseHeadersTransform   = "http_response_headers_transform"
)

// NewUpdateZoneEntrypointRulesetOptions : Instantiate UpdateZoneEntrypointRulesetOptions
func (*RulesetsV1) NewUpdateZoneEntrypointRulesetOptions(rulesetPhase string) *UpdateZoneEntrypointRulesetOptions {
	return &UpdateZoneEntrypointRulesetOptions{
		RulesetPhase: core.StringPtr(rulesetPhase),
	}
}

// SetRulesetPhase : Allow user to set RulesetPhase
func (_options *UpdateZoneEntrypointRulesetOptions) SetRulesetPhase(rulesetPhase string) *UpdateZoneEntrypointRulesetOptions {
	_options.RulesetPhase = core.StringPtr(rulesetPhase)
	return _options
}

// SetDescription : Allow user to set Description
func (_options *UpdateZoneEntrypointRulesetOptions) SetDescription(description string) *UpdateZoneEntrypointRulesetOptions {
	_options.Description = core.StringPtr(description)
	return _options
}

// SetKind : Allow user to set Kind
func (_options *UpdateZoneEntrypointRulesetOptions) SetKind(kind string) *UpdateZoneEntrypointRulesetOptions {
	_options.Kind = core.StringPtr(kind)
	return _options
}

// SetName : Allow user to set Name
func (_options *UpdateZoneEntrypointRulesetOptions) SetName(name string) *UpdateZoneEntrypointRulesetOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetPhase : Allow user to set Phase
func (_options *UpdateZoneEntrypointRulesetOptions) SetPhase(phase string) *UpdateZoneEntrypointRulesetOptions {
	_options.Phase = core.StringPtr(phase)
	return _options
}

// SetRules : Allow user to set Rules
func (_options *UpdateZoneEntrypointRulesetOptions) SetRules(rules []RuleCreate) *UpdateZoneEntrypointRulesetOptions {
	_options.Rules = rules
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateZoneEntrypointRulesetOptions) SetHeaders(param map[string]string) *UpdateZoneEntrypointRulesetOptions {
	options.Headers = param
	return options
}

// UpdateZoneRulesetOptions : The UpdateZoneRuleset options.
type UpdateZoneRulesetOptions struct {
	// ID of a specific ruleset.
	RulesetID *string `json:"ruleset_id" validate:"required,ne="`

	// description of the ruleset.
	Description *string `json:"description,omitempty"`

	Kind *string `json:"kind,omitempty"`

	// human readable name of the ruleset.
	Name *string `json:"name,omitempty"`

	// The phase of the ruleset.
	Phase *string `json:"phase,omitempty"`

	Rules []RuleCreate `json:"rules,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the UpdateZoneRulesetOptions.Kind property.
const (
	UpdateZoneRulesetOptions_Kind_Custom  = "custom"
	UpdateZoneRulesetOptions_Kind_Managed = "managed"
	UpdateZoneRulesetOptions_Kind_Root    = "root"
	UpdateZoneRulesetOptions_Kind_Zone    = "zone"
)

// Constants associated with the UpdateZoneRulesetOptions.Phase property.
// The phase of the ruleset.
const (
	UpdateZoneRulesetOptions_Phase_DdosL4                         = "ddos_l4"
	UpdateZoneRulesetOptions_Phase_DdosL7                         = "ddos_l7"
	UpdateZoneRulesetOptions_Phase_HttpConfigSettings             = "http_config_settings"
	UpdateZoneRulesetOptions_Phase_HttpCustomErrors               = "http_custom_errors"
	UpdateZoneRulesetOptions_Phase_HttpLogCustomFields            = "http_log_custom_fields"
	UpdateZoneRulesetOptions_Phase_HttpRatelimit                  = "http_ratelimit"
	UpdateZoneRulesetOptions_Phase_HttpRequestCacheSettings       = "http_request_cache_settings"
	UpdateZoneRulesetOptions_Phase_HttpRequestDynamicRedirect     = "http_request_dynamic_redirect"
	UpdateZoneRulesetOptions_Phase_HttpRequestFirewallCustom      = "http_request_firewall_custom"
	UpdateZoneRulesetOptions_Phase_HttpRequestFirewallManaged     = "http_request_firewall_managed"
	UpdateZoneRulesetOptions_Phase_HttpRequestLateTransform       = "http_request_late_transform"
	UpdateZoneRulesetOptions_Phase_HttpRequestOrigin              = "http_request_origin"
	UpdateZoneRulesetOptions_Phase_HttpRequestRedirect            = "http_request_redirect"
	UpdateZoneRulesetOptions_Phase_HttpRequestSanitize            = "http_request_sanitize"
	UpdateZoneRulesetOptions_Phase_HttpRequestSbfm                = "http_request_sbfm"
	UpdateZoneRulesetOptions_Phase_HttpRequestSelectConfiguration = "http_request_select_configuration"
	UpdateZoneRulesetOptions_Phase_HttpRequestTransform           = "http_request_transform"
	UpdateZoneRulesetOptions_Phase_HttpResponseCompression        = "http_response_compression"
	UpdateZoneRulesetOptions_Phase_HttpResponseFirewallManaged    = "http_response_firewall_managed"
	UpdateZoneRulesetOptions_Phase_HttpResponseHeadersTransform   = "http_response_headers_transform"
)

// NewUpdateZoneRulesetOptions : Instantiate UpdateZoneRulesetOptions
func (*RulesetsV1) NewUpdateZoneRulesetOptions(rulesetID string) *UpdateZoneRulesetOptions {
	return &UpdateZoneRulesetOptions{
		RulesetID: core.StringPtr(rulesetID),
	}
}

// SetRulesetID : Allow user to set RulesetID
func (_options *UpdateZoneRulesetOptions) SetRulesetID(rulesetID string) *UpdateZoneRulesetOptions {
	_options.RulesetID = core.StringPtr(rulesetID)
	return _options
}

// SetDescription : Allow user to set Description
func (_options *UpdateZoneRulesetOptions) SetDescription(description string) *UpdateZoneRulesetOptions {
	_options.Description = core.StringPtr(description)
	return _options
}

// SetKind : Allow user to set Kind
func (_options *UpdateZoneRulesetOptions) SetKind(kind string) *UpdateZoneRulesetOptions {
	_options.Kind = core.StringPtr(kind)
	return _options
}

// SetName : Allow user to set Name
func (_options *UpdateZoneRulesetOptions) SetName(name string) *UpdateZoneRulesetOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetPhase : Allow user to set Phase
func (_options *UpdateZoneRulesetOptions) SetPhase(phase string) *UpdateZoneRulesetOptions {
	_options.Phase = core.StringPtr(phase)
	return _options
}

// SetRules : Allow user to set Rules
func (_options *UpdateZoneRulesetOptions) SetRules(rules []RuleCreate) *UpdateZoneRulesetOptions {
	_options.Rules = rules
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateZoneRulesetOptions) SetHeaders(param map[string]string) *UpdateZoneRulesetOptions {
	options.Headers = param
	return options
}

// UpdateZoneRulesetRuleOptions : The UpdateZoneRulesetRule options.
type UpdateZoneRulesetRuleOptions struct {
	// ID of a specific ruleset.
	RulesetID *string `json:"ruleset_id" validate:"required,ne="`

	// ID of a specific rule.
	RuleID *string `json:"rule_id" validate:"required,ne="`

	// What happens when theres a match for the rule expression.
	Action *string `json:"action,omitempty"`

	ActionParameters *ActionParameters `json:"action_parameters,omitempty"`

	Description *string `json:"description,omitempty"`

	Enabled *bool `json:"enabled,omitempty"`

	// The expression defining which traffic will match the rule.
	Expression *string `json:"expression,omitempty"`

	ID *string `json:"id,omitempty"`

	Logging *Logging `json:"logging,omitempty"`

	// The reference of the rule (the rule ID by default).
	Ref *string `json:"ref,omitempty"`

	Position *Position `json:"position,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdateZoneRulesetRuleOptions : Instantiate UpdateZoneRulesetRuleOptions
func (*RulesetsV1) NewUpdateZoneRulesetRuleOptions(rulesetID string, ruleID string) *UpdateZoneRulesetRuleOptions {
	return &UpdateZoneRulesetRuleOptions{
		RulesetID: core.StringPtr(rulesetID),
		RuleID:    core.StringPtr(ruleID),
	}
}

// SetRulesetID : Allow user to set RulesetID
func (_options *UpdateZoneRulesetRuleOptions) SetRulesetID(rulesetID string) *UpdateZoneRulesetRuleOptions {
	_options.RulesetID = core.StringPtr(rulesetID)
	return _options
}

// SetRuleID : Allow user to set RuleID
func (_options *UpdateZoneRulesetRuleOptions) SetRuleID(ruleID string) *UpdateZoneRulesetRuleOptions {
	_options.RuleID = core.StringPtr(ruleID)
	return _options
}

// SetAction : Allow user to set Action
func (_options *UpdateZoneRulesetRuleOptions) SetAction(action string) *UpdateZoneRulesetRuleOptions {
	_options.Action = core.StringPtr(action)
	return _options
}

// SetActionParameters : Allow user to set ActionParameters
func (_options *UpdateZoneRulesetRuleOptions) SetActionParameters(actionParameters *ActionParameters) *UpdateZoneRulesetRuleOptions {
	_options.ActionParameters = actionParameters
	return _options
}

// SetDescription : Allow user to set Description
func (_options *UpdateZoneRulesetRuleOptions) SetDescription(description string) *UpdateZoneRulesetRuleOptions {
	_options.Description = core.StringPtr(description)
	return _options
}

// SetEnabled : Allow user to set Enabled
func (_options *UpdateZoneRulesetRuleOptions) SetEnabled(enabled bool) *UpdateZoneRulesetRuleOptions {
	_options.Enabled = core.BoolPtr(enabled)
	return _options
}

// SetExpression : Allow user to set Expression
func (_options *UpdateZoneRulesetRuleOptions) SetExpression(expression string) *UpdateZoneRulesetRuleOptions {
	_options.Expression = core.StringPtr(expression)
	return _options
}

// SetID : Allow user to set ID
func (_options *UpdateZoneRulesetRuleOptions) SetID(id string) *UpdateZoneRulesetRuleOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetLogging : Allow user to set Logging
func (_options *UpdateZoneRulesetRuleOptions) SetLogging(logging *Logging) *UpdateZoneRulesetRuleOptions {
	_options.Logging = logging
	return _options
}

// SetRef : Allow user to set Ref
func (_options *UpdateZoneRulesetRuleOptions) SetRef(ref string) *UpdateZoneRulesetRuleOptions {
	_options.Ref = core.StringPtr(ref)
	return _options
}

// SetPosition : Allow user to set Position
func (_options *UpdateZoneRulesetRuleOptions) SetPosition(position *Position) *UpdateZoneRulesetRuleOptions {
	_options.Position = position
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateZoneRulesetRuleOptions) SetHeaders(param map[string]string) *UpdateZoneRulesetRuleOptions {
	options.Headers = param
	return options
}

// ActionParameters : ActionParameters struct
type ActionParameters struct {
	// unique ID of the ruleset.
	ID *string `json:"id,omitempty"`

	Overrides *Overrides `json:"overrides,omitempty"`

	// The version of the ruleset. Use "latest" to get the latest version.
	Version *string `json:"version,omitempty"`

	// Ruleset ID of the ruleset to apply action to. Use "current" to apply to the current ruleset.
	Ruleset *string `json:"ruleset,omitempty"`

	// List of ruleset ids to apply action to. Use "current" to apply to the current ruleset.
	Rulesets []string `json:"rulesets,omitempty"`

	Response *ActionParametersResponse `json:"response,omitempty"`
}

// UnmarshalActionParameters unmarshals an instance of ActionParameters from the specified map of raw messages.
func UnmarshalActionParameters(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ActionParameters)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "overrides", &obj.Overrides, UnmarshalOverrides)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "version", &obj.Version)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "ruleset", &obj.Ruleset)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "rulesets", &obj.Rulesets)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "response", &obj.Response, UnmarshalActionParametersResponse)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CategoriesOverride : CategoriesOverride struct
type CategoriesOverride struct {
	// The category tag name to override.
	Category *string `json:"category,omitempty"`

	Enabled *bool `json:"enabled,omitempty"`

	// What happens when theres a match for the rule expression.
	Action *string `json:"action,omitempty"`
}

// UnmarshalCategoriesOverride unmarshals an instance of CategoriesOverride from the specified map of raw messages.
func UnmarshalCategoriesOverride(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CategoriesOverride)
	err = core.UnmarshalPrimitive(m, "category", &obj.Category)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "enabled", &obj.Enabled)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "action", &obj.Action)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ListRulesetsResp : List rulesets response.
type ListRulesetsResp struct {
	// Was operation successful.
	Success *bool `json:"success" validate:"required"`

	// Array of errors encountered.
	Errors []Message `json:"errors" validate:"required"`

	// Array of messages returned.
	Messages []Message `json:"messages" validate:"required"`

	// Container for response information.
	Result []ListedRuleset `json:"result" validate:"required"`
}

// UnmarshalListRulesetsResp unmarshals an instance of ListRulesetsResp from the specified map of raw messages.
func UnmarshalListRulesetsResp(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ListRulesetsResp)
	err = core.UnmarshalPrimitive(m, "success", &obj.Success)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "errors", &obj.Errors, UnmarshalMessage)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "messages", &obj.Messages, UnmarshalMessage)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "result", &obj.Result, UnmarshalListedRuleset)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ListedRuleset : ListedRuleset struct
type ListedRuleset struct {
	// description of the ruleset.
	Description *string `json:"description" validate:"required"`

	// unique ID of the ruleset.
	ID *string `json:"id" validate:"required"`

	Kind *string `json:"kind" validate:"required"`

	// The timestamp of when the resource was last modified.
	LastUpdated *string `json:"last_updated" validate:"required"`

	// human readable name of the ruleset.
	Name *string `json:"name" validate:"required"`

	// The phase of the ruleset.
	Phase *string `json:"phase" validate:"required"`

	// The version of the ruleset.
	Version *string `json:"version" validate:"required"`
}

// Constants associated with the ListedRuleset.Kind property.
const (
	ListedRuleset_Kind_Custom  = "custom"
	ListedRuleset_Kind_Managed = "managed"
	ListedRuleset_Kind_Root    = "root"
	ListedRuleset_Kind_Zone    = "zone"
)

// Constants associated with the ListedRuleset.Phase property.
// The phase of the ruleset.
const (
	ListedRuleset_Phase_DdosL4                         = "ddos_l4"
	ListedRuleset_Phase_DdosL7                         = "ddos_l7"
	ListedRuleset_Phase_HttpConfigSettings             = "http_config_settings"
	ListedRuleset_Phase_HttpCustomErrors               = "http_custom_errors"
	ListedRuleset_Phase_HttpLogCustomFields            = "http_log_custom_fields"
	ListedRuleset_Phase_HttpRatelimit                  = "http_ratelimit"
	ListedRuleset_Phase_HttpRequestCacheSettings       = "http_request_cache_settings"
	ListedRuleset_Phase_HttpRequestDynamicRedirect     = "http_request_dynamic_redirect"
	ListedRuleset_Phase_HttpRequestFirewallCustom      = "http_request_firewall_custom"
	ListedRuleset_Phase_HttpRequestFirewallManaged     = "http_request_firewall_managed"
	ListedRuleset_Phase_HttpRequestLateTransform       = "http_request_late_transform"
	ListedRuleset_Phase_HttpRequestOrigin              = "http_request_origin"
	ListedRuleset_Phase_HttpRequestRedirect            = "http_request_redirect"
	ListedRuleset_Phase_HttpRequestSanitize            = "http_request_sanitize"
	ListedRuleset_Phase_HttpRequestSbfm                = "http_request_sbfm"
	ListedRuleset_Phase_HttpRequestSelectConfiguration = "http_request_select_configuration"
	ListedRuleset_Phase_HttpRequestTransform           = "http_request_transform"
	ListedRuleset_Phase_HttpResponseCompression        = "http_response_compression"
	ListedRuleset_Phase_HttpResponseFirewallManaged    = "http_response_firewall_managed"
	ListedRuleset_Phase_HttpResponseHeadersTransform   = "http_response_headers_transform"
)

// UnmarshalListedRuleset unmarshals an instance of ListedRuleset from the specified map of raw messages.
func UnmarshalListedRuleset(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ListedRuleset)
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "kind", &obj.Kind)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "last_updated", &obj.LastUpdated)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "phase", &obj.Phase)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "version", &obj.Version)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Logging : Logging struct
type Logging struct {
	Enabled *bool `json:"enabled" validate:"required"`
}

// NewLogging : Instantiate Logging (Generic Model Constructor)
func (*RulesetsV1) NewLogging(enabled bool) (_model *Logging, err error) {
	_model = &Logging{
		Enabled: core.BoolPtr(enabled),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalLogging unmarshals an instance of Logging from the specified map of raw messages.
func UnmarshalLogging(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Logging)
	err = core.UnmarshalPrimitive(m, "enabled", &obj.Enabled)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Message : Message struct
type Message struct {
	// A unique code for this message.
	Code *int64 `json:"code,omitempty"`

	// A text description of this message.
	Message *string `json:"message" validate:"required"`

	// The source of this message.
	Source *MessageSource `json:"source,omitempty"`
}

// UnmarshalMessage unmarshals an instance of Message from the specified map of raw messages.
func UnmarshalMessage(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Message)
	err = core.UnmarshalPrimitive(m, "code", &obj.Code)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "message", &obj.Message)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "source", &obj.Source, UnmarshalMessageSource)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Overrides : Overrides struct
type Overrides struct {
	// What happens when theres a match for the rule expression.
	Action *string `json:"action,omitempty"`

	Enabled *bool `json:"enabled,omitempty"`

	// The sensitivity level of the rule.
	SensitivityLevel *string `json:"sensitivity_level,omitempty"`

	Rules []RulesOverride `json:"rules,omitempty"`

	Categories []CategoriesOverride `json:"categories,omitempty"`
}

// Constants associated with the Overrides.SensitivityLevel property.
// The sensitivity level of the rule.
const (
	Overrides_SensitivityLevel_High   = "high"
	Overrides_SensitivityLevel_Low    = "low"
	Overrides_SensitivityLevel_Medium = "medium"
)

// UnmarshalOverrides unmarshals an instance of Overrides from the specified map of raw messages.
func UnmarshalOverrides(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Overrides)
	err = core.UnmarshalPrimitive(m, "action", &obj.Action)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "enabled", &obj.Enabled)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "sensitivity_level", &obj.SensitivityLevel)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "rules", &obj.Rules, UnmarshalRulesOverride)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "categories", &obj.Categories, UnmarshalCategoriesOverride)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Position : Position struct
type Position struct {
	// The rule ID to place this rule before.
	Before *string `json:"before,omitempty"`

	// The rule ID to place this rule after.
	After *string `json:"after,omitempty"`

	// The index to place this rule at.
	Index *int64 `json:"index,omitempty"`
}

// UnmarshalPosition unmarshals an instance of Position from the specified map of raw messages.
func UnmarshalPosition(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Position)
	err = core.UnmarshalPrimitive(m, "before", &obj.Before)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "after", &obj.After)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "index", &obj.Index)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RuleCreate : RuleCreate struct
type RuleCreate struct {
	// What happens when theres a match for the rule expression.
	Action *string `json:"action" validate:"required"`

	ActionParameters *ActionParameters `json:"action_parameters,omitempty"`

	Description *string `json:"description,omitempty"`

	Enabled *bool `json:"enabled,omitempty"`

	// The expression defining which traffic will match the rule.
	Expression *string `json:"expression" validate:"required"`

	ID *string `json:"id,omitempty"`

	Logging *Logging `json:"logging,omitempty"`

	// The reference of the rule (the rule ID by default).
	Ref *string `json:"ref,omitempty"`

	Position *Position `json:"position,omitempty"`
}

// NewRuleCreate : Instantiate RuleCreate (Generic Model Constructor)
func (*RulesetsV1) NewRuleCreate(action string, expression string) (_model *RuleCreate, err error) {
	_model = &RuleCreate{
		Action:     core.StringPtr(action),
		Expression: core.StringPtr(expression),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalRuleCreate unmarshals an instance of RuleCreate from the specified map of raw messages.
func UnmarshalRuleCreate(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RuleCreate)
	err = core.UnmarshalPrimitive(m, "action", &obj.Action)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "action_parameters", &obj.ActionParameters, UnmarshalActionParameters)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "enabled", &obj.Enabled)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "expression", &obj.Expression)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "logging", &obj.Logging, UnmarshalLogging)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "ref", &obj.Ref)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "position", &obj.Position, UnmarshalPosition)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RuleDetails : RuleDetails struct
type RuleDetails struct {
	// unique ID of rule.
	ID *string `json:"id" validate:"required"`

	// The version of the rule.
	Version *string `json:"version" validate:"required"`

	// What happens when theres a match for the rule expression.
	Action *string `json:"action" validate:"required"`

	ActionParameters *ActionParameters `json:"action_parameters,omitempty"`

	// The expression defining which traffic will match the rule.
	Expression *string `json:"expression" validate:"required"`

	// The reference of the rule (the rule ID by default).
	Ref *string `json:"ref,omitempty"`

	Logging *Logging `json:"logging,omitempty"`

	// The timestamp of when the resource was last modified.
	LastUpdated *string `json:"last_updated" validate:"required"`
}

// UnmarshalRuleDetails unmarshals an instance of RuleDetails from the specified map of raw messages.
func UnmarshalRuleDetails(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RuleDetails)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "version", &obj.Version)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "action", &obj.Action)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "action_parameters", &obj.ActionParameters, UnmarshalActionParameters)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "expression", &obj.Expression)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "ref", &obj.Ref)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "logging", &obj.Logging, UnmarshalLogging)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "last_updated", &obj.LastUpdated)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RuleResp : List rules response.
type RuleResp struct {
	// Was operation successful.
	Success *bool `json:"success" validate:"required"`

	// Array of errors encountered.
	Errors []Message `json:"errors" validate:"required"`

	// Array of messages returned.
	Messages []Message `json:"messages" validate:"required"`

	Result *RuleDetails `json:"result" validate:"required"`
}

// UnmarshalRuleResp unmarshals an instance of RuleResp from the specified map of raw messages.
func UnmarshalRuleResp(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RuleResp)
	err = core.UnmarshalPrimitive(m, "success", &obj.Success)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "errors", &obj.Errors, UnmarshalMessage)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "messages", &obj.Messages, UnmarshalMessage)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "result", &obj.Result, UnmarshalRuleDetails)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RulesOverride : RulesOverride struct
type RulesOverride struct {
	ID *string `json:"id,omitempty"`

	Enabled *bool `json:"enabled,omitempty"`

	// What happens when theres a match for the rule expression.
	Action *string `json:"action,omitempty"`

	// The sensitivity level of the rule.
	SensitivityLevel *string `json:"sensitivity_level,omitempty"`
}

// Constants associated with the RulesOverride.SensitivityLevel property.
// The sensitivity level of the rule.
const (
	RulesOverride_SensitivityLevel_High   = "high"
	RulesOverride_SensitivityLevel_Low    = "low"
	RulesOverride_SensitivityLevel_Medium = "medium"
)

// UnmarshalRulesOverride unmarshals an instance of RulesOverride from the specified map of raw messages.
func UnmarshalRulesOverride(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RulesOverride)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "enabled", &obj.Enabled)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "action", &obj.Action)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "sensitivity_level", &obj.SensitivityLevel)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RulesetDetails : RulesetDetails struct
type RulesetDetails struct {
	// description of the ruleset.
	Description *string `json:"description" validate:"required"`

	// unique ID of the ruleset.
	ID *string `json:"id" validate:"required"`

	Kind *string `json:"kind" validate:"required"`

	// The timestamp of when the resource was last modified.
	LastUpdated *string `json:"last_updated" validate:"required"`

	// human readable name of the ruleset.
	Name *string `json:"name" validate:"required"`

	// The phase of the ruleset.
	Phase *string `json:"phase" validate:"required"`

	// The version of the ruleset.
	Version *string `json:"version" validate:"required"`

	Rules []RuleDetails `json:"rules" validate:"required"`
}

// Constants associated with the RulesetDetails.Kind property.
const (
	RulesetDetails_Kind_Custom  = "custom"
	RulesetDetails_Kind_Managed = "managed"
	RulesetDetails_Kind_Root    = "root"
	RulesetDetails_Kind_Zone    = "zone"
)

// Constants associated with the RulesetDetails.Phase property.
// The phase of the ruleset.
const (
	RulesetDetails_Phase_DdosL4                         = "ddos_l4"
	RulesetDetails_Phase_DdosL7                         = "ddos_l7"
	RulesetDetails_Phase_HttpConfigSettings             = "http_config_settings"
	RulesetDetails_Phase_HttpCustomErrors               = "http_custom_errors"
	RulesetDetails_Phase_HttpLogCustomFields            = "http_log_custom_fields"
	RulesetDetails_Phase_HttpRatelimit                  = "http_ratelimit"
	RulesetDetails_Phase_HttpRequestCacheSettings       = "http_request_cache_settings"
	RulesetDetails_Phase_HttpRequestDynamicRedirect     = "http_request_dynamic_redirect"
	RulesetDetails_Phase_HttpRequestFirewallCustom      = "http_request_firewall_custom"
	RulesetDetails_Phase_HttpRequestFirewallManaged     = "http_request_firewall_managed"
	RulesetDetails_Phase_HttpRequestLateTransform       = "http_request_late_transform"
	RulesetDetails_Phase_HttpRequestOrigin              = "http_request_origin"
	RulesetDetails_Phase_HttpRequestRedirect            = "http_request_redirect"
	RulesetDetails_Phase_HttpRequestSanitize            = "http_request_sanitize"
	RulesetDetails_Phase_HttpRequestSbfm                = "http_request_sbfm"
	RulesetDetails_Phase_HttpRequestSelectConfiguration = "http_request_select_configuration"
	RulesetDetails_Phase_HttpRequestTransform           = "http_request_transform"
	RulesetDetails_Phase_HttpResponseCompression        = "http_response_compression"
	RulesetDetails_Phase_HttpResponseFirewallManaged    = "http_response_firewall_managed"
	RulesetDetails_Phase_HttpResponseHeadersTransform   = "http_response_headers_transform"
)

// UnmarshalRulesetDetails unmarshals an instance of RulesetDetails from the specified map of raw messages.
func UnmarshalRulesetDetails(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RulesetDetails)
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "kind", &obj.Kind)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "last_updated", &obj.LastUpdated)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "phase", &obj.Phase)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "version", &obj.Version)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "rules", &obj.Rules, UnmarshalRuleDetails)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RulesetResp : Ruleset response.
type RulesetResp struct {
	// Was operation successful.
	Success *bool `json:"success" validate:"required"`

	// Array of errors encountered.
	Errors []Message `json:"errors" validate:"required"`

	// Array of messages returned.
	Messages []Message `json:"messages" validate:"required"`

	Result *RulesetDetails `json:"result" validate:"required"`
}

// UnmarshalRulesetResp unmarshals an instance of RulesetResp from the specified map of raw messages.
func UnmarshalRulesetResp(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RulesetResp)
	err = core.UnmarshalPrimitive(m, "success", &obj.Success)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "errors", &obj.Errors, UnmarshalMessage)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "messages", &obj.Messages, UnmarshalMessage)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "result", &obj.Result, UnmarshalRulesetDetails)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}
