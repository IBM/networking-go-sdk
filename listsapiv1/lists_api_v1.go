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
 * IBM OpenAPI SDK Code Generator Version: 3.104.0-b4a47c49-20250418-184351
 */

// Package listsapiv1 : Operations and models for the ListsApiV1 service
package listsapiv1

import (
	"context"
	"encoding/json"
	"net/http"
	"reflect"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	common "github.com/IBM/networking-go-sdk/common"
)

// ListsApiV1 : CIS Lists
//
// API Version: 1.0.0
type ListsApiV1 struct {
	Service *core.BaseService

	// Full URL-encoded CRN of the service instance.
	Crn *string

	// List item identifier.
	ItemID *string

	// List identifier.
	ListID *string

	// List operation identifier.
	OperationID *string
}

// DefaultServiceURL is the default URL to make service requests to.
const DefaultServiceURL = "https://api.cis.cloud.ibm.com"

// DefaultServiceName is the default key used to find external configuration information.
const DefaultServiceName = "lists_api"

// ListsApiV1Options : Service options
type ListsApiV1Options struct {
	ServiceName   string
	URL           string
	Authenticator core.Authenticator

	// Full URL-encoded CRN of the service instance.
	Crn *string `validate:"required"`

	// List item identifier.
	ItemID *string `validate:"required"`

	// List identifier.
	ListID *string `validate:"required"`

	// List operation identifier.
	OperationID *string `validate:"required"`
}

// NewListsApiV1UsingExternalConfig : constructs an instance of ListsApiV1 with passed in options and external configuration.
func NewListsApiV1UsingExternalConfig(options *ListsApiV1Options) (listsApi *ListsApiV1, err error) {
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

	listsApi, err = NewListsApiV1(options)
	err = core.RepurposeSDKProblem(err, "new-client-error")
	if err != nil {
		return
	}

	err = listsApi.Service.ConfigureService(options.ServiceName)
	if err != nil {
		err = core.SDKErrorf(err, "", "client-config-error", common.GetComponentInfo())
		return
	}

	if options.URL != "" {
		err = listsApi.Service.SetServiceURL(options.URL)
		err = core.RepurposeSDKProblem(err, "url-set-error")
	}
	return
}

// NewListsApiV1 : constructs an instance of ListsApiV1 with passed in options.
func NewListsApiV1(options *ListsApiV1Options) (service *ListsApiV1, err error) {
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

	service = &ListsApiV1{
		Service:     baseService,
		Crn:         options.Crn,
		ItemID:      options.ItemID,
		ListID:      options.ListID,
		OperationID: options.OperationID,
	}

	return
}

// GetServiceURLForRegion returns the service URL to be used for the specified region
func GetServiceURLForRegion(region string) (string, error) {
	return "", core.SDKErrorf(nil, "service does not support regional URLs", "no-regional-support", common.GetComponentInfo())
}

// Clone makes a copy of "listsApi" suitable for processing requests.
func (listsApi *ListsApiV1) Clone() *ListsApiV1 {
	if core.IsNil(listsApi) {
		return nil
	}
	clone := *listsApi
	clone.Service = listsApi.Service.Clone()
	return &clone
}

// SetServiceURL sets the service URL
func (listsApi *ListsApiV1) SetServiceURL(url string) error {
	err := listsApi.Service.SetServiceURL(url)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-set-error", common.GetComponentInfo())
	}
	return err
}

// GetServiceURL returns the service URL
func (listsApi *ListsApiV1) GetServiceURL() string {
	return listsApi.Service.GetServiceURL()
}

// SetDefaultHeaders sets HTTP headers to be sent in every request
func (listsApi *ListsApiV1) SetDefaultHeaders(headers http.Header) {
	listsApi.Service.SetDefaultHeaders(headers)
}

// SetEnableGzipCompression sets the service's EnableGzipCompression field
func (listsApi *ListsApiV1) SetEnableGzipCompression(enableGzip bool) {
	listsApi.Service.SetEnableGzipCompression(enableGzip)
}

// GetEnableGzipCompression returns the service's EnableGzipCompression field
func (listsApi *ListsApiV1) GetEnableGzipCompression() bool {
	return listsApi.Service.GetEnableGzipCompression()
}

// EnableRetries enables automatic retries for requests invoked for this service instance.
// If either parameter is specified as 0, then a default value is used instead.
func (listsApi *ListsApiV1) EnableRetries(maxRetries int, maxRetryInterval time.Duration) {
	listsApi.Service.EnableRetries(maxRetries, maxRetryInterval)
}

// DisableRetries disables automatic retries for requests invoked for this service instance.
func (listsApi *ListsApiV1) DisableRetries() {
	listsApi.Service.DisableRetries()
}

// GetManagedLists : List Managed Lists
// List available managed lists for your instance.
func (listsApi *ListsApiV1) GetManagedLists(getManagedListsOptions *GetManagedListsOptions) (result *ManagedListsResp, response *core.DetailedResponse, err error) {
	result, response, err = listsApi.GetManagedListsWithContext(context.Background(), getManagedListsOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// GetManagedListsWithContext is an alternate form of the GetManagedLists method which supports a Context parameter
func (listsApi *ListsApiV1) GetManagedListsWithContext(ctx context.Context, getManagedListsOptions *GetManagedListsOptions) (result *ManagedListsResp, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getManagedListsOptions, "getManagedListsOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"crn": *listsApi.Crn,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = listsApi.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(listsApi.Service.Options.URL, `/v1/{crn}/rules/managed_lists`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range getManagedListsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("lists_api", "V1", "GetManagedLists")
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
	response, err = listsApi.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "get_managed_lists", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalManagedListsResp)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// GetCustomLists : List Custom Lists
// List the custom lists for your instance.
func (listsApi *ListsApiV1) GetCustomLists(getCustomListsOptions *GetCustomListsOptions) (result *CustomListsResp, response *core.DetailedResponse, err error) {
	result, response, err = listsApi.GetCustomListsWithContext(context.Background(), getCustomListsOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// GetCustomListsWithContext is an alternate form of the GetCustomLists method which supports a Context parameter
func (listsApi *ListsApiV1) GetCustomListsWithContext(ctx context.Context, getCustomListsOptions *GetCustomListsOptions) (result *CustomListsResp, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getCustomListsOptions, "getCustomListsOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"crn": *listsApi.Crn,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = listsApi.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(listsApi.Service.Options.URL, `/v1/{crn}/rules/lists`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range getCustomListsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("lists_api", "V1", "GetCustomLists")
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
	response, err = listsApi.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "get_custom_lists", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCustomListsResp)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// CreateCustomLists : Create Custom List
// Create a custom list for your instance.
func (listsApi *ListsApiV1) CreateCustomLists(createCustomListsOptions *CreateCustomListsOptions) (result *CustomListResp, response *core.DetailedResponse, err error) {
	result, response, err = listsApi.CreateCustomListsWithContext(context.Background(), createCustomListsOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// CreateCustomListsWithContext is an alternate form of the CreateCustomLists method which supports a Context parameter
func (listsApi *ListsApiV1) CreateCustomListsWithContext(ctx context.Context, createCustomListsOptions *CreateCustomListsOptions) (result *CustomListResp, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(createCustomListsOptions, "createCustomListsOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"crn": *listsApi.Crn,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = listsApi.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(listsApi.Service.Options.URL, `/v1/{crn}/rules/lists`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range createCustomListsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("lists_api", "V1", "CreateCustomLists")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if createCustomListsOptions.Kind != nil {
		body["kind"] = createCustomListsOptions.Kind
	}
	if createCustomListsOptions.Name != nil {
		body["name"] = createCustomListsOptions.Name
	}
	if createCustomListsOptions.Description != nil {
		body["description"] = createCustomListsOptions.Description
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
	response, err = listsApi.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "create_custom_lists", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCustomListResp)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// GetCustomList : Get Custom List
// Get a custom list for your instance.
func (listsApi *ListsApiV1) GetCustomList(getCustomListOptions *GetCustomListOptions) (result *CustomListResp, response *core.DetailedResponse, err error) {
	result, response, err = listsApi.GetCustomListWithContext(context.Background(), getCustomListOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// GetCustomListWithContext is an alternate form of the GetCustomList method which supports a Context parameter
func (listsApi *ListsApiV1) GetCustomListWithContext(ctx context.Context, getCustomListOptions *GetCustomListOptions) (result *CustomListResp, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getCustomListOptions, "getCustomListOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"crn":     *listsApi.Crn,
		"list_id": *listsApi.ListID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = listsApi.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(listsApi.Service.Options.URL, `/v1/{crn}/rules/lists/{list_id}`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range getCustomListOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("lists_api", "V1", "GetCustomList")
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
	response, err = listsApi.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "get_custom_list", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCustomListResp)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// UpdateCustomList : Update Custom List
// Update the description of a custom list.
func (listsApi *ListsApiV1) UpdateCustomList(updateCustomListOptions *UpdateCustomListOptions) (result *CustomListResp, response *core.DetailedResponse, err error) {
	result, response, err = listsApi.UpdateCustomListWithContext(context.Background(), updateCustomListOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// UpdateCustomListWithContext is an alternate form of the UpdateCustomList method which supports a Context parameter
func (listsApi *ListsApiV1) UpdateCustomListWithContext(ctx context.Context, updateCustomListOptions *UpdateCustomListOptions) (result *CustomListResp, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(updateCustomListOptions, "updateCustomListOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"crn":     *listsApi.Crn,
		"list_id": *listsApi.ListID,
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = listsApi.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(listsApi.Service.Options.URL, `/v1/{crn}/rules/lists/{list_id}`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range updateCustomListOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("lists_api", "V1", "UpdateCustomList")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if updateCustomListOptions.Description != nil {
		body["description"] = updateCustomListOptions.Description
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
	response, err = listsApi.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "update_custom_list", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCustomListResp)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// DeleteCustomList : Delete Custom List
// Delete a custom list for your instance.
func (listsApi *ListsApiV1) DeleteCustomList(deleteCustomListOptions *DeleteCustomListOptions) (result *DeleteResourceResp, response *core.DetailedResponse, err error) {
	result, response, err = listsApi.DeleteCustomListWithContext(context.Background(), deleteCustomListOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// DeleteCustomListWithContext is an alternate form of the DeleteCustomList method which supports a Context parameter
func (listsApi *ListsApiV1) DeleteCustomListWithContext(ctx context.Context, deleteCustomListOptions *DeleteCustomListOptions) (result *DeleteResourceResp, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(deleteCustomListOptions, "deleteCustomListOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"crn":     *listsApi.Crn,
		"list_id": *listsApi.ListID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = listsApi.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(listsApi.Service.Options.URL, `/v1/{crn}/rules/lists/{list_id}`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range deleteCustomListOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("lists_api", "V1", "DeleteCustomList")
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
	response, err = listsApi.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "delete_custom_list", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDeleteResourceResp)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// GetListItems : Get List Items
// Get the list items for a custom list.
func (listsApi *ListsApiV1) GetListItems(getListItemsOptions *GetListItemsOptions) (result *ListItemsResp, response *core.DetailedResponse, err error) {
	result, response, err = listsApi.GetListItemsWithContext(context.Background(), getListItemsOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// GetListItemsWithContext is an alternate form of the GetListItems method which supports a Context parameter
func (listsApi *ListsApiV1) GetListItemsWithContext(ctx context.Context, getListItemsOptions *GetListItemsOptions) (result *ListItemsResp, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getListItemsOptions, "getListItemsOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"crn":     *listsApi.Crn,
		"list_id": *listsApi.ListID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = listsApi.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(listsApi.Service.Options.URL, `/v1/{crn}/rules/lists/{list_id}/items`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range getListItemsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("lists_api", "V1", "GetListItems")
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
	response, err = listsApi.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "get_list_items", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalListItemsResp)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// CreateListItems : Create List Items
// Create list items for your custom list. This operation is asynchronous. To get current the operation status, use the
// get operation status API.
func (listsApi *ListsApiV1) CreateListItems(createListItemsOptions *CreateListItemsOptions) (result *ListOperationResp, response *core.DetailedResponse, err error) {
	result, response, err = listsApi.CreateListItemsWithContext(context.Background(), createListItemsOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// CreateListItemsWithContext is an alternate form of the CreateListItems method which supports a Context parameter
func (listsApi *ListsApiV1) CreateListItemsWithContext(ctx context.Context, createListItemsOptions *CreateListItemsOptions) (result *ListOperationResp, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(createListItemsOptions, "createListItemsOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"crn":     *listsApi.Crn,
		"list_id": *listsApi.ListID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = listsApi.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(listsApi.Service.Options.URL, `/v1/{crn}/rules/lists/{list_id}/items`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range createListItemsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("lists_api", "V1", "CreateListItems")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	if createListItemsOptions.CreateListItemsReqItem != nil {
		_, err = builder.SetBodyContentJSON(createListItemsOptions.CreateListItemsReqItem)
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
	response, err = listsApi.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "create_list_items", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalListOperationResp)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// DeleteListItems : Delete List Items
// Remove one or more list items from your custom list. This operation is asynchronous. To get current the operation
// status, use the get operation status API.
func (listsApi *ListsApiV1) DeleteListItems(deleteListItemsOptions *DeleteListItemsOptions) (result *ListOperationResp, response *core.DetailedResponse, err error) {
	result, response, err = listsApi.DeleteListItemsWithContext(context.Background(), deleteListItemsOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// DeleteListItemsWithContext is an alternate form of the DeleteListItems method which supports a Context parameter
func (listsApi *ListsApiV1) DeleteListItemsWithContext(ctx context.Context, deleteListItemsOptions *DeleteListItemsOptions) (result *ListOperationResp, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(deleteListItemsOptions, "deleteListItemsOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"crn":     *listsApi.Crn,
		"list_id": *listsApi.ListID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = listsApi.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(listsApi.Service.Options.URL, `/v1/{crn}/rules/lists/{list_id}/items`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range deleteListItemsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("lists_api", "V1", "DeleteListItems")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if deleteListItemsOptions.Items != nil {
		body["items"] = deleteListItemsOptions.Items
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
	response, err = listsApi.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "delete_list_items", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalListOperationResp)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// UpdateListItems : Update All List Items
// Update all list items for your custom list. This removes existing items from the list. This operation is
// asynchronous. To get current the operation status, use the get operation status API.
func (listsApi *ListsApiV1) UpdateListItems(updateListItemsOptions *UpdateListItemsOptions) (result *ListOperationResp, response *core.DetailedResponse, err error) {
	result, response, err = listsApi.UpdateListItemsWithContext(context.Background(), updateListItemsOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// UpdateListItemsWithContext is an alternate form of the UpdateListItems method which supports a Context parameter
func (listsApi *ListsApiV1) UpdateListItemsWithContext(ctx context.Context, updateListItemsOptions *UpdateListItemsOptions) (result *ListOperationResp, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(updateListItemsOptions, "updateListItemsOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"crn":     *listsApi.Crn,
		"list_id": *listsApi.ListID,
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = listsApi.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(listsApi.Service.Options.URL, `/v1/{crn}/rules/lists/{list_id}/items`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range updateListItemsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("lists_api", "V1", "UpdateListItems")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	if updateListItemsOptions.CreateListItemsReqItem != nil {
		_, err = builder.SetBodyContentJSON(updateListItemsOptions.CreateListItemsReqItem)
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
	response, err = listsApi.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "update_list_items", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalListOperationResp)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// GetListItem : Get List Item
// Get a specific list item.
func (listsApi *ListsApiV1) GetListItem(getListItemOptions *GetListItemOptions) (result *ListItemResp, response *core.DetailedResponse, err error) {
	result, response, err = listsApi.GetListItemWithContext(context.Background(), getListItemOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// GetListItemWithContext is an alternate form of the GetListItem method which supports a Context parameter
func (listsApi *ListsApiV1) GetListItemWithContext(ctx context.Context, getListItemOptions *GetListItemOptions) (result *ListItemResp, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getListItemOptions, "getListItemOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"crn":     *listsApi.Crn,
		"list_id": *listsApi.ListID,
		"item_id": *listsApi.ItemID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = listsApi.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(listsApi.Service.Options.URL, `/v1/{crn}/rules/lists/{list_id}/items/{item_id}`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range getListItemOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("lists_api", "V1", "GetListItem")
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
	response, err = listsApi.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "get_list_item", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalListItemResp)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// GetOperationStatus : Get List Operation Status
// Get the operation status for a custom list operation.
func (listsApi *ListsApiV1) GetOperationStatus(getOperationStatusOptions *GetOperationStatusOptions) (result *OperationStatusResp, response *core.DetailedResponse, err error) {
	result, response, err = listsApi.GetOperationStatusWithContext(context.Background(), getOperationStatusOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// GetOperationStatusWithContext is an alternate form of the GetOperationStatus method which supports a Context parameter
func (listsApi *ListsApiV1) GetOperationStatusWithContext(ctx context.Context, getOperationStatusOptions *GetOperationStatusOptions) (result *OperationStatusResp, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getOperationStatusOptions, "getOperationStatusOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"crn":          *listsApi.Crn,
		"operation_id": *listsApi.OperationID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = listsApi.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(listsApi.Service.Options.URL, `/v1/{crn}/rules/lists/bulk_operations/{operation_id}`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range getOperationStatusOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("lists_api", "V1", "GetOperationStatus")
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
	response, err = listsApi.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "get_operation_status", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalOperationStatusResp)
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

// CreateCustomListsOptions : The CreateCustomLists options.
type CreateCustomListsOptions struct {
	// The type of list. Each type supports specific list items (IP addresses, ASNs, hostnames or redirects).
	Kind *string `json:"kind,omitempty"`

	// An informative name for the list. Use this name in rule expressions.
	Name *string `json:"name,omitempty"`

	// An informative summary of the list.
	Description *string `json:"description,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// Constants associated with the CreateCustomListsOptions.Kind property.
// The type of list. Each type supports specific list items (IP addresses, ASNs, hostnames or redirects).
const (
	CreateCustomListsOptions_Kind_Asn      = "asn"
	CreateCustomListsOptions_Kind_Hostname = "hostname"
	CreateCustomListsOptions_Kind_Ip       = "ip"
	CreateCustomListsOptions_Kind_Redirect = "redirect"
)

// NewCreateCustomListsOptions : Instantiate CreateCustomListsOptions
func (*ListsApiV1) NewCreateCustomListsOptions() *CreateCustomListsOptions {
	return &CreateCustomListsOptions{}
}

// SetKind : Allow user to set Kind
func (_options *CreateCustomListsOptions) SetKind(kind string) *CreateCustomListsOptions {
	_options.Kind = core.StringPtr(kind)
	return _options
}

// SetName : Allow user to set Name
func (_options *CreateCustomListsOptions) SetName(name string) *CreateCustomListsOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetDescription : Allow user to set Description
func (_options *CreateCustomListsOptions) SetDescription(description string) *CreateCustomListsOptions {
	_options.Description = core.StringPtr(description)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateCustomListsOptions) SetHeaders(param map[string]string) *CreateCustomListsOptions {
	options.Headers = param
	return options
}

// CreateListItemsOptions : The CreateListItems options.
type CreateListItemsOptions struct {
	CreateListItemsReqItem []CreateListItemsReqItem `json:"CreateListItemsReqItem,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewCreateListItemsOptions : Instantiate CreateListItemsOptions
func (*ListsApiV1) NewCreateListItemsOptions() *CreateListItemsOptions {
	return &CreateListItemsOptions{}
}

// SetCreateListItemsReqItem : Allow user to set CreateListItemsReqItem
func (_options *CreateListItemsOptions) SetCreateListItemsReqItem(createListItemsReqItem []CreateListItemsReqItem) *CreateListItemsOptions {
	_options.CreateListItemsReqItem = createListItemsReqItem
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateListItemsOptions) SetHeaders(param map[string]string) *CreateListItemsOptions {
	options.Headers = param
	return options
}

// CreateListItemsReqItem : CreateListItemsReqItem struct
type CreateListItemsReqItem struct {
	// An autonomous system number.
	Asn *float64 `json:"asn,omitempty"`

	// An informative summary of the list item.
	Comment *string `json:"comment,omitempty"`

	// Valid characters for hostnames are ASCII(7) letters from a to z, the digits from 0 to 9, wildcards (*), and the
	// hyphen (-).
	Hostname *string `json:"hostname,omitempty"`

	// An IPv4 address, an IPv4 CIDR, or an IPv6 CIDR. IPv6 CIDRs are limited to a maximum of /64.
	Ip *string `json:"ip,omitempty"`
}

// UnmarshalCreateListItemsReqItem unmarshals an instance of CreateListItemsReqItem from the specified map of raw messages.
func UnmarshalCreateListItemsReqItem(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CreateListItemsReqItem)
	err = core.UnmarshalPrimitive(m, "asn", &obj.Asn)
	if err != nil {
		err = core.SDKErrorf(err, "", "asn-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "comment", &obj.Comment)
	if err != nil {
		err = core.SDKErrorf(err, "", "comment-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "hostname", &obj.Hostname)
	if err != nil {
		err = core.SDKErrorf(err, "", "hostname-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "ip", &obj.Ip)
	if err != nil {
		err = core.SDKErrorf(err, "", "ip-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DeleteCustomListOptions : The DeleteCustomList options.
type DeleteCustomListOptions struct {

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewDeleteCustomListOptions : Instantiate DeleteCustomListOptions
func (*ListsApiV1) NewDeleteCustomListOptions() *DeleteCustomListOptions {
	return &DeleteCustomListOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *DeleteCustomListOptions) SetHeaders(param map[string]string) *DeleteCustomListOptions {
	options.Headers = param
	return options
}

// DeleteListItemsOptions : The DeleteListItems options.
type DeleteListItemsOptions struct {
	Items []DeleteListItemsReqItemsItem `json:"items,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewDeleteListItemsOptions : Instantiate DeleteListItemsOptions
func (*ListsApiV1) NewDeleteListItemsOptions() *DeleteListItemsOptions {
	return &DeleteListItemsOptions{}
}

// SetItems : Allow user to set Items
func (_options *DeleteListItemsOptions) SetItems(items []DeleteListItemsReqItemsItem) *DeleteListItemsOptions {
	_options.Items = items
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteListItemsOptions) SetHeaders(param map[string]string) *DeleteListItemsOptions {
	options.Headers = param
	return options
}

// DeleteListItemsReqItemsItem : DeleteListItemsReqItemsItem struct
type DeleteListItemsReqItemsItem struct {
	ID *string `json:"id,omitempty"`
}

// UnmarshalDeleteListItemsReqItemsItem unmarshals an instance of DeleteListItemsReqItemsItem from the specified map of raw messages.
func UnmarshalDeleteListItemsReqItemsItem(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DeleteListItemsReqItemsItem)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		err = core.SDKErrorf(err, "", "id-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DeleteResourceRespResult : DeleteResourceRespResult struct
type DeleteResourceRespResult struct {
	ID *string `json:"id,omitempty"`
}

// UnmarshalDeleteResourceRespResult unmarshals an instance of DeleteResourceRespResult from the specified map of raw messages.
func UnmarshalDeleteResourceRespResult(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DeleteResourceRespResult)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		err = core.SDKErrorf(err, "", "id-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetCustomListOptions : The GetCustomList options.
type GetCustomListOptions struct {

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewGetCustomListOptions : Instantiate GetCustomListOptions
func (*ListsApiV1) NewGetCustomListOptions() *GetCustomListOptions {
	return &GetCustomListOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *GetCustomListOptions) SetHeaders(param map[string]string) *GetCustomListOptions {
	options.Headers = param
	return options
}

// GetCustomListsOptions : The GetCustomLists options.
type GetCustomListsOptions struct {

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewGetCustomListsOptions : Instantiate GetCustomListsOptions
func (*ListsApiV1) NewGetCustomListsOptions() *GetCustomListsOptions {
	return &GetCustomListsOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *GetCustomListsOptions) SetHeaders(param map[string]string) *GetCustomListsOptions {
	options.Headers = param
	return options
}

// GetListItemOptions : The GetListItem options.
type GetListItemOptions struct {

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewGetListItemOptions : Instantiate GetListItemOptions
func (*ListsApiV1) NewGetListItemOptions() *GetListItemOptions {
	return &GetListItemOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *GetListItemOptions) SetHeaders(param map[string]string) *GetListItemOptions {
	options.Headers = param
	return options
}

// GetListItemsOptions : The GetListItems options.
type GetListItemsOptions struct {

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewGetListItemsOptions : Instantiate GetListItemsOptions
func (*ListsApiV1) NewGetListItemsOptions() *GetListItemsOptions {
	return &GetListItemsOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *GetListItemsOptions) SetHeaders(param map[string]string) *GetListItemsOptions {
	options.Headers = param
	return options
}

// GetManagedListsOptions : The GetManagedLists options.
type GetManagedListsOptions struct {

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewGetManagedListsOptions : Instantiate GetManagedListsOptions
func (*ListsApiV1) NewGetManagedListsOptions() *GetManagedListsOptions {
	return &GetManagedListsOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *GetManagedListsOptions) SetHeaders(param map[string]string) *GetManagedListsOptions {
	options.Headers = param
	return options
}

// GetOperationStatusOptions : The GetOperationStatus options.
type GetOperationStatusOptions struct {

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewGetOperationStatusOptions : Instantiate GetOperationStatusOptions
func (*ListsApiV1) NewGetOperationStatusOptions() *GetOperationStatusOptions {
	return &GetOperationStatusOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *GetOperationStatusOptions) SetHeaders(param map[string]string) *GetOperationStatusOptions {
	options.Headers = param
	return options
}

// ListOperationRespResult : ListOperationRespResult struct
type ListOperationRespResult struct {
	OperationID *string `json:"operation_id,omitempty"`
}

// UnmarshalListOperationRespResult unmarshals an instance of ListOperationRespResult from the specified map of raw messages.
func UnmarshalListOperationRespResult(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ListOperationRespResult)
	err = core.UnmarshalPrimitive(m, "operation_id", &obj.OperationID)
	if err != nil {
		err = core.SDKErrorf(err, "", "operation_id-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ManagedListsResultItem : ManagedListsResultItem struct
type ManagedListsResultItem struct {
	// The name of the list to be referenced by rule expressions.
	Name *string `json:"name,omitempty"`

	// Describes the contents of the managed list.
	Description *string `json:"description,omitempty"`

	// The type of resource this list contains.
	Kind *string `json:"kind,omitempty"`
}

// UnmarshalManagedListsResultItem unmarshals an instance of ManagedListsResultItem from the specified map of raw messages.
func UnmarshalManagedListsResultItem(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ManagedListsResultItem)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		err = core.SDKErrorf(err, "", "name-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		err = core.SDKErrorf(err, "", "description-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "kind", &obj.Kind)
	if err != nil {
		err = core.SDKErrorf(err, "", "kind-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// OperationStatusRespResult : OperationStatusRespResult struct
type OperationStatusRespResult struct {
	ID *string `json:"id,omitempty"`

	Status *string `json:"status,omitempty"`

	Completed *string `json:"completed,omitempty"`

	// A message describing the error when the status is failed.
	Error *string `json:"error,omitempty"`
}

// Constants associated with the OperationStatusRespResult.Status property.
const (
	OperationStatusRespResult_Status_Completed = "completed"
	OperationStatusRespResult_Status_Failed    = "failed"
	OperationStatusRespResult_Status_Pending   = "pending"
	OperationStatusRespResult_Status_Running   = "running"
)

// UnmarshalOperationStatusRespResult unmarshals an instance of OperationStatusRespResult from the specified map of raw messages.
func UnmarshalOperationStatusRespResult(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(OperationStatusRespResult)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		err = core.SDKErrorf(err, "", "id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "status", &obj.Status)
	if err != nil {
		err = core.SDKErrorf(err, "", "status-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "completed", &obj.Completed)
	if err != nil {
		err = core.SDKErrorf(err, "", "completed-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "error", &obj.Error)
	if err != nil {
		err = core.SDKErrorf(err, "", "error-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// UpdateCustomListOptions : The UpdateCustomList options.
type UpdateCustomListOptions struct {
	// An informative summary of the list.
	Description *string `json:"description,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewUpdateCustomListOptions : Instantiate UpdateCustomListOptions
func (*ListsApiV1) NewUpdateCustomListOptions() *UpdateCustomListOptions {
	return &UpdateCustomListOptions{}
}

// SetDescription : Allow user to set Description
func (_options *UpdateCustomListOptions) SetDescription(description string) *UpdateCustomListOptions {
	_options.Description = core.StringPtr(description)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateCustomListOptions) SetHeaders(param map[string]string) *UpdateCustomListOptions {
	options.Headers = param
	return options
}

// UpdateListItemsOptions : The UpdateListItems options.
type UpdateListItemsOptions struct {
	CreateListItemsReqItem []CreateListItemsReqItem `json:"CreateListItemsReqItem,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewUpdateListItemsOptions : Instantiate UpdateListItemsOptions
func (*ListsApiV1) NewUpdateListItemsOptions() *UpdateListItemsOptions {
	return &UpdateListItemsOptions{}
}

// SetCreateListItemsReqItem : Allow user to set CreateListItemsReqItem
func (_options *UpdateListItemsOptions) SetCreateListItemsReqItem(createListItemsReqItem []CreateListItemsReqItem) *UpdateListItemsOptions {
	_options.CreateListItemsReqItem = createListItemsReqItem
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateListItemsOptions) SetHeaders(param map[string]string) *UpdateListItemsOptions {
	options.Headers = param
	return options
}

// CustomListResp : Create Custom List Response.
type CustomListResp struct {
	// Was operation successful.
	Success *bool `json:"success" validate:"required"`

	// Errors.
	Errors [][]string `json:"errors" validate:"required"`

	// Messages.
	Messages [][]string `json:"messages" validate:"required"`

	Result *CustomListResult `json:"result" validate:"required"`
}

// UnmarshalCustomListResp unmarshals an instance of CustomListResp from the specified map of raw messages.
func UnmarshalCustomListResp(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CustomListResp)
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
	err = core.UnmarshalModel(m, "result", &obj.Result, UnmarshalCustomListResult)
	if err != nil {
		err = core.SDKErrorf(err, "", "result-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CustomListResult : CustomListResult struct
type CustomListResult struct {
	// The name of the list to be referenced by rule expressions.
	Name *string `json:"name,omitempty"`

	// The unique ID of the list.
	ID *string `json:"id,omitempty"`

	// Describes the contents of the list.
	Description *string `json:"description,omitempty"`

	// The type of resource this list contains.
	Kind *string `json:"kind,omitempty"`

	// How many items the list contains.
	NumItems *float64 `json:"num_items,omitempty"`

	// How many times the list is used by rule expressions.
	NumReferencingFilters *float64 `json:"num_referencing_filters,omitempty"`
}

// UnmarshalCustomListResult unmarshals an instance of CustomListResult from the specified map of raw messages.
func UnmarshalCustomListResult(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CustomListResult)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		err = core.SDKErrorf(err, "", "name-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		err = core.SDKErrorf(err, "", "id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		err = core.SDKErrorf(err, "", "description-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "kind", &obj.Kind)
	if err != nil {
		err = core.SDKErrorf(err, "", "kind-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "num_items", &obj.NumItems)
	if err != nil {
		err = core.SDKErrorf(err, "", "num_items-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "num_referencing_filters", &obj.NumReferencingFilters)
	if err != nil {
		err = core.SDKErrorf(err, "", "num_referencing_filters-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CustomListsResp : List Custom Lists Response.
type CustomListsResp struct {
	// Was operation successful.
	Success *bool `json:"success" validate:"required"`

	// Errors.
	Errors [][]string `json:"errors" validate:"required"`

	// Messages.
	Messages [][]string `json:"messages" validate:"required"`

	Result []CustomListResult `json:"result" validate:"required"`
}

// UnmarshalCustomListsResp unmarshals an instance of CustomListsResp from the specified map of raw messages.
func UnmarshalCustomListsResp(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CustomListsResp)
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
	err = core.UnmarshalModel(m, "result", &obj.Result, UnmarshalCustomListResult)
	if err != nil {
		err = core.SDKErrorf(err, "", "result-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DeleteResourceResp : DeleteResourceResp struct
type DeleteResourceResp struct {
	// Was operation successful.
	Success *bool `json:"success" validate:"required"`

	// Errors.
	Errors [][]string `json:"errors" validate:"required"`

	// Messages.
	Messages [][]string `json:"messages" validate:"required"`

	Result *DeleteResourceRespResult `json:"result" validate:"required"`
}

// UnmarshalDeleteResourceResp unmarshals an instance of DeleteResourceResp from the specified map of raw messages.
func UnmarshalDeleteResourceResp(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DeleteResourceResp)
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
	err = core.UnmarshalModel(m, "result", &obj.Result, UnmarshalDeleteResourceRespResult)
	if err != nil {
		err = core.SDKErrorf(err, "", "result-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ListItem : ListItem struct
type ListItem struct {
	ID *string `json:"id,omitempty"`

	// An autonomous system number.
	Asn *float64 `json:"asn,omitempty"`

	// An informative summary of the list item.
	Comment *string `json:"comment,omitempty"`

	// Valid characters for hostnames are ASCII(7) letters from a to z, the digits from 0 to 9, wildcards (*), and the
	// hyphen (-).
	Hostname *string `json:"hostname,omitempty"`

	// An IPv4 address, an IPv4 CIDR, or an IPv6 CIDR. IPv6 CIDRs are limited to a maximum of /64.
	Ip *string `json:"ip,omitempty"`

	CreatedOn *string `json:"created_on,omitempty"`

	ModifiedOn *string `json:"modified_on,omitempty"`
}

// UnmarshalListItem unmarshals an instance of ListItem from the specified map of raw messages.
func UnmarshalListItem(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ListItem)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		err = core.SDKErrorf(err, "", "id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "asn", &obj.Asn)
	if err != nil {
		err = core.SDKErrorf(err, "", "asn-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "comment", &obj.Comment)
	if err != nil {
		err = core.SDKErrorf(err, "", "comment-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "hostname", &obj.Hostname)
	if err != nil {
		err = core.SDKErrorf(err, "", "hostname-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "ip", &obj.Ip)
	if err != nil {
		err = core.SDKErrorf(err, "", "ip-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "created_on", &obj.CreatedOn)
	if err != nil {
		err = core.SDKErrorf(err, "", "created_on-error", common.GetComponentInfo())
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

// ListItemResp : ListItemResp struct
type ListItemResp struct {
	// Was operation successful.
	Success *bool `json:"success" validate:"required"`

	// Errors.
	Errors [][]string `json:"errors" validate:"required"`

	// Messages.
	Messages [][]string `json:"messages" validate:"required"`

	Result *ListItem `json:"result" validate:"required"`
}

// UnmarshalListItemResp unmarshals an instance of ListItemResp from the specified map of raw messages.
func UnmarshalListItemResp(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ListItemResp)
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
	err = core.UnmarshalModel(m, "result", &obj.Result, UnmarshalListItem)
	if err != nil {
		err = core.SDKErrorf(err, "", "result-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ListItemsResp : ListItemsResp struct
type ListItemsResp struct {
	// Was operation successful.
	Success *bool `json:"success" validate:"required"`

	// Errors.
	Errors [][]string `json:"errors" validate:"required"`

	// Messages.
	Messages [][]string `json:"messages" validate:"required"`

	Result []ListItem `json:"result" validate:"required"`
}

// UnmarshalListItemsResp unmarshals an instance of ListItemsResp from the specified map of raw messages.
func UnmarshalListItemsResp(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ListItemsResp)
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
	err = core.UnmarshalModel(m, "result", &obj.Result, UnmarshalListItem)
	if err != nil {
		err = core.SDKErrorf(err, "", "result-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ListOperationResp : ListOperationResp struct
type ListOperationResp struct {
	// Was operation successful.
	Success *bool `json:"success" validate:"required"`

	// Errors.
	Errors [][]string `json:"errors" validate:"required"`

	// Messages.
	Messages [][]string `json:"messages" validate:"required"`

	Result *ListOperationRespResult `json:"result" validate:"required"`
}

// UnmarshalListOperationResp unmarshals an instance of ListOperationResp from the specified map of raw messages.
func UnmarshalListOperationResp(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ListOperationResp)
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
	err = core.UnmarshalModel(m, "result", &obj.Result, UnmarshalListOperationRespResult)
	if err != nil {
		err = core.SDKErrorf(err, "", "result-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ManagedListsResp : List Managed Lists Response.
type ManagedListsResp struct {
	// Was operation successful.
	Success *bool `json:"success" validate:"required"`

	// Errors.
	Errors [][]string `json:"errors" validate:"required"`

	// Messages.
	Messages [][]string `json:"messages" validate:"required"`

	Result []ManagedListsResultItem `json:"result" validate:"required"`
}

// UnmarshalManagedListsResp unmarshals an instance of ManagedListsResp from the specified map of raw messages.
func UnmarshalManagedListsResp(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ManagedListsResp)
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
	err = core.UnmarshalModel(m, "result", &obj.Result, UnmarshalManagedListsResultItem)
	if err != nil {
		err = core.SDKErrorf(err, "", "result-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// OperationStatusResp : OperationStatusResp struct
type OperationStatusResp struct {
	// Was operation successful.
	Success *bool `json:"success" validate:"required"`

	// Errors.
	Errors [][]string `json:"errors" validate:"required"`

	// Messages.
	Messages [][]string `json:"messages" validate:"required"`

	Result *OperationStatusRespResult `json:"result" validate:"required"`
}

// UnmarshalOperationStatusResp unmarshals an instance of OperationStatusResp from the specified map of raw messages.
func UnmarshalOperationStatusResp(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(OperationStatusResp)
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
	err = core.UnmarshalModel(m, "result", &obj.Result, UnmarshalOperationStatusRespResult)
	if err != nil {
		err = core.SDKErrorf(err, "", "result-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}
