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

// Package cachingapiv1 : Operations and models for the CachingApiV1 service
package cachingapiv1

import (
	"encoding/json"
	"github.com/IBM/go-sdk-core/v4/core"
	common "github.com/IBM/networking-go-sdk/common"
	"reflect"
)

// CachingApiV1 : This document describes CIS caching  API.
//
// Version: 1.0.0
type CachingApiV1 struct {
	Service *core.BaseService

	// cloud resource name.
	Crn *string

	// zone id.
	ZoneID *string
}

// DefaultServiceURL is the default URL to make service requests to.
const DefaultServiceURL = "https://api.cis.cloud.ibm.com"

// DefaultServiceName is the default key used to find external configuration information.
const DefaultServiceName = "caching_api"

// CachingApiV1Options : Service options
type CachingApiV1Options struct {
	ServiceName   string
	URL           string
	Authenticator core.Authenticator

	// cloud resource name.
	Crn *string `validate:"required"`

	// zone id.
	ZoneID *string `validate:"required"`
}

// NewCachingApiV1UsingExternalConfig : constructs an instance of CachingApiV1 with passed in options and external configuration.
func NewCachingApiV1UsingExternalConfig(options *CachingApiV1Options) (cachingApi *CachingApiV1, err error) {
	if options.ServiceName == "" {
		options.ServiceName = DefaultServiceName
	}

	if options.Authenticator == nil {
		options.Authenticator, err = core.GetAuthenticatorFromEnvironment(options.ServiceName)
		if err != nil {
			return
		}
	}

	cachingApi, err = NewCachingApiV1(options)
	if err != nil {
		return
	}

	err = cachingApi.Service.ConfigureService(options.ServiceName)
	if err != nil {
		return
	}

	if options.URL != "" {
		err = cachingApi.Service.SetServiceURL(options.URL)
	}
	return
}

// NewCachingApiV1 : constructs an instance of CachingApiV1 with passed in options.
func NewCachingApiV1(options *CachingApiV1Options) (service *CachingApiV1, err error) {
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

	service = &CachingApiV1{
		Service: baseService,
		Crn: options.Crn,
		ZoneID: options.ZoneID,
	}

	return
}

// SetServiceURL sets the service URL
func (cachingApi *CachingApiV1) SetServiceURL(url string) error {
	return cachingApi.Service.SetServiceURL(url)
}

// PurgeAll : Purge all files under a specific domain
// All resources in CDN edge servers' cache should be removed. This may have dramatic affects on your origin server load
// after performing this action.
func (cachingApi *CachingApiV1) PurgeAll(purgeAllOptions *PurgeAllOptions) (result *PurgeAllResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(purgeAllOptions, "purgeAllOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1", "zones", "purge_cache/purge_all"}
	pathParameters := []string{*cachingApi.Crn, *cachingApi.ZoneID}

	builder := core.NewRequestBuilder(core.PUT)
	_, err = builder.ConstructHTTPURL(cachingApi.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range purgeAllOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("caching_api", "V1", "PurgeAll")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = cachingApi.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalPurgeAllResponse)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// PurgeByUrls : Purge individual files by URLs
// Granularly remove one or more files from CDN edge servers' cache either by specifying URLs.
func (cachingApi *CachingApiV1) PurgeByUrls(purgeByUrlsOptions *PurgeByUrlsOptions) (result *PurgeAllResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(purgeByUrlsOptions, "purgeByUrlsOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1", "zones", "purge_cache/purge_by_urls"}
	pathParameters := []string{*cachingApi.Crn, *cachingApi.ZoneID}

	builder := core.NewRequestBuilder(core.PUT)
	_, err = builder.ConstructHTTPURL(cachingApi.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range purgeByUrlsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("caching_api", "V1", "PurgeByUrls")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if purgeByUrlsOptions.Files != nil {
		body["files"] = purgeByUrlsOptions.Files
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
	response, err = cachingApi.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalPurgeAllResponse)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// PurgeByCacheTags : Purge files by Cache-Tags
// Granularly remove one or more files from CDN edge servers' cache either by specifying the associated Cache-Tags.
func (cachingApi *CachingApiV1) PurgeByCacheTags(purgeByCacheTagsOptions *PurgeByCacheTagsOptions) (result *PurgeAllResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(purgeByCacheTagsOptions, "purgeByCacheTagsOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1", "zones", "purge_cache/purge_by_cache_tags"}
	pathParameters := []string{*cachingApi.Crn, *cachingApi.ZoneID}

	builder := core.NewRequestBuilder(core.PUT)
	_, err = builder.ConstructHTTPURL(cachingApi.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range purgeByCacheTagsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("caching_api", "V1", "PurgeByCacheTags")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if purgeByCacheTagsOptions.Tags != nil {
		body["tags"] = purgeByCacheTagsOptions.Tags
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
	response, err = cachingApi.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalPurgeAllResponse)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// PurgeByHosts : Purge individual files by hostnames
// Granularly remove one or more files from CDN edge servers' cache either by specifying the hostnames.
func (cachingApi *CachingApiV1) PurgeByHosts(purgeByHostsOptions *PurgeByHostsOptions) (result *PurgeAllResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(purgeByHostsOptions, "purgeByHostsOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1", "zones", "purge_cache/purge_by_hosts"}
	pathParameters := []string{*cachingApi.Crn, *cachingApi.ZoneID}

	builder := core.NewRequestBuilder(core.PUT)
	_, err = builder.ConstructHTTPURL(cachingApi.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range purgeByHostsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("caching_api", "V1", "PurgeByHosts")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if purgeByHostsOptions.Hosts != nil {
		body["hosts"] = purgeByHostsOptions.Hosts
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
	response, err = cachingApi.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalPurgeAllResponse)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// GetBrowserCacheTTL : Get browser cache TTL setting
// Browser Cache TTL (in seconds) specifies how long CDN edge servers cached resources will remain on your visitors'
// computers.
func (cachingApi *CachingApiV1) GetBrowserCacheTTL(getBrowserCacheTtlOptions *GetBrowserCacheTtlOptions) (result *BrowserTTLResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getBrowserCacheTtlOptions, "getBrowserCacheTtlOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1", "zones", "settings/browser_cache_ttl"}
	pathParameters := []string{*cachingApi.Crn, *cachingApi.ZoneID}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(cachingApi.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range getBrowserCacheTtlOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("caching_api", "V1", "GetBrowserCacheTTL")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = cachingApi.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalBrowserTTLResponse)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// UpdateBrowserCacheTTL : Change browser cache TTL setting
// Browser Cache TTL (in seconds) specifies how long CDN edge servers cached resources will remain on your visitors'
// computers.
func (cachingApi *CachingApiV1) UpdateBrowserCacheTTL(updateBrowserCacheTtlOptions *UpdateBrowserCacheTtlOptions) (result *BrowserTTLResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(updateBrowserCacheTtlOptions, "updateBrowserCacheTtlOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1", "zones", "settings/browser_cache_ttl"}
	pathParameters := []string{*cachingApi.Crn, *cachingApi.ZoneID}

	builder := core.NewRequestBuilder(core.PATCH)
	_, err = builder.ConstructHTTPURL(cachingApi.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateBrowserCacheTtlOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("caching_api", "V1", "UpdateBrowserCacheTTL")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if updateBrowserCacheTtlOptions.Value != nil {
		body["value"] = updateBrowserCacheTtlOptions.Value
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
	response, err = cachingApi.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalBrowserTTLResponse)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// GetDevelopmentMode : Get development mode setting
// Get development mode setting.
func (cachingApi *CachingApiV1) GetDevelopmentMode(getDevelopmentModeOptions *GetDevelopmentModeOptions) (result *DeveopmentModeResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getDevelopmentModeOptions, "getDevelopmentModeOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1", "zones", "settings/development_mode"}
	pathParameters := []string{*cachingApi.Crn, *cachingApi.ZoneID}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(cachingApi.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range getDevelopmentModeOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("caching_api", "V1", "GetDevelopmentMode")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = cachingApi.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDeveopmentModeResponse)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// UpdateDevelopmentMode : Change development mode setting
// Change development mode setting.
func (cachingApi *CachingApiV1) UpdateDevelopmentMode(updateDevelopmentModeOptions *UpdateDevelopmentModeOptions) (result *DeveopmentModeResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(updateDevelopmentModeOptions, "updateDevelopmentModeOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1", "zones", "settings/development_mode"}
	pathParameters := []string{*cachingApi.Crn, *cachingApi.ZoneID}

	builder := core.NewRequestBuilder(core.PATCH)
	_, err = builder.ConstructHTTPURL(cachingApi.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateDevelopmentModeOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("caching_api", "V1", "UpdateDevelopmentMode")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if updateDevelopmentModeOptions.Value != nil {
		body["value"] = updateDevelopmentModeOptions.Value
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
	response, err = cachingApi.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDeveopmentModeResponse)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// GetQueryStringSort : Get Enable Query String Sort setting
// Get Enable Query String Sort setting.
func (cachingApi *CachingApiV1) GetQueryStringSort(getQueryStringSortOptions *GetQueryStringSortOptions) (result *EnableQueryStringSortResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getQueryStringSortOptions, "getQueryStringSortOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1", "zones", "settings/sort_query_string_for_cache"}
	pathParameters := []string{*cachingApi.Crn, *cachingApi.ZoneID}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(cachingApi.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range getQueryStringSortOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("caching_api", "V1", "GetQueryStringSort")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = cachingApi.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalEnableQueryStringSortResponse)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// UpdateQueryStringSort : Change Enable Query String Sort setting
// Change Enable Query String Sort setting.
func (cachingApi *CachingApiV1) UpdateQueryStringSort(updateQueryStringSortOptions *UpdateQueryStringSortOptions) (result *EnableQueryStringSortResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(updateQueryStringSortOptions, "updateQueryStringSortOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1", "zones", "settings/sort_query_string_for_cache"}
	pathParameters := []string{*cachingApi.Crn, *cachingApi.ZoneID}

	builder := core.NewRequestBuilder(core.PATCH)
	_, err = builder.ConstructHTTPURL(cachingApi.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateQueryStringSortOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("caching_api", "V1", "UpdateQueryStringSort")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if updateQueryStringSortOptions.Value != nil {
		body["value"] = updateQueryStringSortOptions.Value
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
	response, err = cachingApi.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalEnableQueryStringSortResponse)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// GetCacheLevel : Get cache level setting of a specific zone
// Get cache level setting of a specific zone.
func (cachingApi *CachingApiV1) GetCacheLevel(getCacheLevelOptions *GetCacheLevelOptions) (result *CacheLevelResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getCacheLevelOptions, "getCacheLevelOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1", "zones", "settings/cache_level"}
	pathParameters := []string{*cachingApi.Crn, *cachingApi.ZoneID}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(cachingApi.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range getCacheLevelOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("caching_api", "V1", "GetCacheLevel")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = cachingApi.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCacheLevelResponse)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// UpdateCacheLevel : Set cache level setting for a specific zone
// The `basic` setting will cache most static resources (i.e., css, images, and JavaScript). The `simplified` setting
// will ignore the query string when delivering a cached resource. The `aggressive` setting will cache all static
// resources, including ones with a query string.
func (cachingApi *CachingApiV1) UpdateCacheLevel(updateCacheLevelOptions *UpdateCacheLevelOptions) (result *CacheLevelResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(updateCacheLevelOptions, "updateCacheLevelOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1", "zones", "settings/cache_level"}
	pathParameters := []string{*cachingApi.Crn, *cachingApi.ZoneID}

	builder := core.NewRequestBuilder(core.PATCH)
	_, err = builder.ConstructHTTPURL(cachingApi.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateCacheLevelOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("caching_api", "V1", "UpdateCacheLevel")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if updateCacheLevelOptions.Value != nil {
		body["value"] = updateCacheLevelOptions.Value
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
	response, err = cachingApi.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCacheLevelResponse)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// BrowserTTLResponseResult : result object.
type BrowserTTLResponseResult struct {
	// ttl type.
	ID *string `json:"id,omitempty"`

	// ttl value.
	Value *int64 `json:"value,omitempty"`

	// editable.
	Editable *bool `json:"editable,omitempty"`

	// modified date.
	ModifiedOn *string `json:"modified_on,omitempty"`
}


// UnmarshalBrowserTTLResponseResult unmarshals an instance of BrowserTTLResponseResult from the specified map of raw messages.
func UnmarshalBrowserTTLResponseResult(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(BrowserTTLResponseResult)
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

// CacheLevelResponseResult : result.
type CacheLevelResponseResult struct {
	// cache level.
	ID *string `json:"id,omitempty"`

	// cache level.
	Value *string `json:"value,omitempty"`

	// editable value.
	Editable *bool `json:"editable,omitempty"`

	// modified date.
	ModifiedOn *string `json:"modified_on,omitempty"`
}


// UnmarshalCacheLevelResponseResult unmarshals an instance of CacheLevelResponseResult from the specified map of raw messages.
func UnmarshalCacheLevelResponseResult(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CacheLevelResponseResult)
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

// DeveopmentModeResponseResult : result object.
type DeveopmentModeResponseResult struct {
	// object id.
	ID *string `json:"id,omitempty"`

	// on/off value.
	Value *string `json:"value,omitempty"`

	// editable value.
	Editable *bool `json:"editable,omitempty"`

	// modified date.
	ModifiedOn *string `json:"modified_on,omitempty"`
}


// UnmarshalDeveopmentModeResponseResult unmarshals an instance of DeveopmentModeResponseResult from the specified map of raw messages.
func UnmarshalDeveopmentModeResponseResult(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DeveopmentModeResponseResult)
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

// EnableQueryStringSortResponseResult : result of sort query string.
type EnableQueryStringSortResponseResult struct {
	// cache id.
	ID *string `json:"id,omitempty"`

	// on/off value.
	Value *string `json:"value,omitempty"`

	// editable propery.
	Editable *bool `json:"editable,omitempty"`

	// modified date.
	ModifiedOn *string `json:"modified_on,omitempty"`
}


// UnmarshalEnableQueryStringSortResponseResult unmarshals an instance of EnableQueryStringSortResponseResult from the specified map of raw messages.
func UnmarshalEnableQueryStringSortResponseResult(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(EnableQueryStringSortResponseResult)
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

// GetBrowserCacheTtlOptions : The GetBrowserCacheTTL options.
type GetBrowserCacheTtlOptions struct {

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetBrowserCacheTtlOptions : Instantiate GetBrowserCacheTtlOptions
func (*CachingApiV1) NewGetBrowserCacheTtlOptions() *GetBrowserCacheTtlOptions {
	return &GetBrowserCacheTtlOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *GetBrowserCacheTtlOptions) SetHeaders(param map[string]string) *GetBrowserCacheTtlOptions {
	options.Headers = param
	return options
}

// GetCacheLevelOptions : The GetCacheLevel options.
type GetCacheLevelOptions struct {

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetCacheLevelOptions : Instantiate GetCacheLevelOptions
func (*CachingApiV1) NewGetCacheLevelOptions() *GetCacheLevelOptions {
	return &GetCacheLevelOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *GetCacheLevelOptions) SetHeaders(param map[string]string) *GetCacheLevelOptions {
	options.Headers = param
	return options
}

// GetDevelopmentModeOptions : The GetDevelopmentMode options.
type GetDevelopmentModeOptions struct {

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetDevelopmentModeOptions : Instantiate GetDevelopmentModeOptions
func (*CachingApiV1) NewGetDevelopmentModeOptions() *GetDevelopmentModeOptions {
	return &GetDevelopmentModeOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *GetDevelopmentModeOptions) SetHeaders(param map[string]string) *GetDevelopmentModeOptions {
	options.Headers = param
	return options
}

// GetQueryStringSortOptions : The GetQueryStringSort options.
type GetQueryStringSortOptions struct {

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetQueryStringSortOptions : Instantiate GetQueryStringSortOptions
func (*CachingApiV1) NewGetQueryStringSortOptions() *GetQueryStringSortOptions {
	return &GetQueryStringSortOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *GetQueryStringSortOptions) SetHeaders(param map[string]string) *GetQueryStringSortOptions {
	options.Headers = param
	return options
}

// PurgeAllOptions : The PurgeAll options.
type PurgeAllOptions struct {

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewPurgeAllOptions : Instantiate PurgeAllOptions
func (*CachingApiV1) NewPurgeAllOptions() *PurgeAllOptions {
	return &PurgeAllOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *PurgeAllOptions) SetHeaders(param map[string]string) *PurgeAllOptions {
	options.Headers = param
	return options
}

// PurgeAllResponseResult : purge object.
type PurgeAllResponseResult struct {
	// purge id.
	ID *string `json:"id,omitempty"`
}


// UnmarshalPurgeAllResponseResult unmarshals an instance of PurgeAllResponseResult from the specified map of raw messages.
func UnmarshalPurgeAllResponseResult(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PurgeAllResponseResult)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// PurgeByCacheTagsOptions : The PurgeByCacheTags options.
type PurgeByCacheTagsOptions struct {
	// array of tags.
	Tags []string `json:"tags,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewPurgeByCacheTagsOptions : Instantiate PurgeByCacheTagsOptions
func (*CachingApiV1) NewPurgeByCacheTagsOptions() *PurgeByCacheTagsOptions {
	return &PurgeByCacheTagsOptions{}
}

// SetTags : Allow user to set Tags
func (options *PurgeByCacheTagsOptions) SetTags(tags []string) *PurgeByCacheTagsOptions {
	options.Tags = tags
	return options
}

// SetHeaders : Allow user to set Headers
func (options *PurgeByCacheTagsOptions) SetHeaders(param map[string]string) *PurgeByCacheTagsOptions {
	options.Headers = param
	return options
}

// PurgeByHostsOptions : The PurgeByHosts options.
type PurgeByHostsOptions struct {
	// hosts name.
	Hosts []string `json:"hosts,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewPurgeByHostsOptions : Instantiate PurgeByHostsOptions
func (*CachingApiV1) NewPurgeByHostsOptions() *PurgeByHostsOptions {
	return &PurgeByHostsOptions{}
}

// SetHosts : Allow user to set Hosts
func (options *PurgeByHostsOptions) SetHosts(hosts []string) *PurgeByHostsOptions {
	options.Hosts = hosts
	return options
}

// SetHeaders : Allow user to set Headers
func (options *PurgeByHostsOptions) SetHeaders(param map[string]string) *PurgeByHostsOptions {
	options.Headers = param
	return options
}

// PurgeByUrlsOptions : The PurgeByUrls options.
type PurgeByUrlsOptions struct {
	// purge url array.
	Files []string `json:"files,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewPurgeByUrlsOptions : Instantiate PurgeByUrlsOptions
func (*CachingApiV1) NewPurgeByUrlsOptions() *PurgeByUrlsOptions {
	return &PurgeByUrlsOptions{}
}

// SetFiles : Allow user to set Files
func (options *PurgeByUrlsOptions) SetFiles(files []string) *PurgeByUrlsOptions {
	options.Files = files
	return options
}

// SetHeaders : Allow user to set Headers
func (options *PurgeByUrlsOptions) SetHeaders(param map[string]string) *PurgeByUrlsOptions {
	options.Headers = param
	return options
}

// UpdateBrowserCacheTtlOptions : The UpdateBrowserCacheTTL options.
type UpdateBrowserCacheTtlOptions struct {
	// ttl value.
	Value *int64 `json:"value,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdateBrowserCacheTtlOptions : Instantiate UpdateBrowserCacheTtlOptions
func (*CachingApiV1) NewUpdateBrowserCacheTtlOptions() *UpdateBrowserCacheTtlOptions {
	return &UpdateBrowserCacheTtlOptions{}
}

// SetValue : Allow user to set Value
func (options *UpdateBrowserCacheTtlOptions) SetValue(value int64) *UpdateBrowserCacheTtlOptions {
	options.Value = core.Int64Ptr(value)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateBrowserCacheTtlOptions) SetHeaders(param map[string]string) *UpdateBrowserCacheTtlOptions {
	options.Headers = param
	return options
}

// UpdateCacheLevelOptions : The UpdateCacheLevel options.
type UpdateCacheLevelOptions struct {
	// cache level.
	Value *string `json:"value,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the UpdateCacheLevelOptions.Value property.
// cache level.
const (
	UpdateCacheLevelOptions_Value_Aggressive = "aggressive"
	UpdateCacheLevelOptions_Value_Basic = "basic"
	UpdateCacheLevelOptions_Value_Simplified = "simplified"
)

// NewUpdateCacheLevelOptions : Instantiate UpdateCacheLevelOptions
func (*CachingApiV1) NewUpdateCacheLevelOptions() *UpdateCacheLevelOptions {
	return &UpdateCacheLevelOptions{}
}

// SetValue : Allow user to set Value
func (options *UpdateCacheLevelOptions) SetValue(value string) *UpdateCacheLevelOptions {
	options.Value = core.StringPtr(value)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateCacheLevelOptions) SetHeaders(param map[string]string) *UpdateCacheLevelOptions {
	options.Headers = param
	return options
}

// UpdateDevelopmentModeOptions : The UpdateDevelopmentMode options.
type UpdateDevelopmentModeOptions struct {
	// on/off value.
	Value *string `json:"value,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the UpdateDevelopmentModeOptions.Value property.
// on/off value.
const (
	UpdateDevelopmentModeOptions_Value_Off = "off"
	UpdateDevelopmentModeOptions_Value_On = "on"
)

// NewUpdateDevelopmentModeOptions : Instantiate UpdateDevelopmentModeOptions
func (*CachingApiV1) NewUpdateDevelopmentModeOptions() *UpdateDevelopmentModeOptions {
	return &UpdateDevelopmentModeOptions{}
}

// SetValue : Allow user to set Value
func (options *UpdateDevelopmentModeOptions) SetValue(value string) *UpdateDevelopmentModeOptions {
	options.Value = core.StringPtr(value)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateDevelopmentModeOptions) SetHeaders(param map[string]string) *UpdateDevelopmentModeOptions {
	options.Headers = param
	return options
}

// UpdateQueryStringSortOptions : The UpdateQueryStringSort options.
type UpdateQueryStringSortOptions struct {
	// on/off property value.
	Value *string `json:"value,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the UpdateQueryStringSortOptions.Value property.
// on/off property value.
const (
	UpdateQueryStringSortOptions_Value_Off = "off"
	UpdateQueryStringSortOptions_Value_On = "on"
)

// NewUpdateQueryStringSortOptions : Instantiate UpdateQueryStringSortOptions
func (*CachingApiV1) NewUpdateQueryStringSortOptions() *UpdateQueryStringSortOptions {
	return &UpdateQueryStringSortOptions{}
}

// SetValue : Allow user to set Value
func (options *UpdateQueryStringSortOptions) SetValue(value string) *UpdateQueryStringSortOptions {
	options.Value = core.StringPtr(value)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateQueryStringSortOptions) SetHeaders(param map[string]string) *UpdateQueryStringSortOptions {
	options.Headers = param
	return options
}

// BrowserTTLResponse : browser ttl response.
type BrowserTTLResponse struct {
	// success response.
	Success *bool `json:"success" validate:"required"`

	// errors.
	Errors [][]string `json:"errors" validate:"required"`

	// messages.
	Messages [][]string `json:"messages" validate:"required"`

	// result object.
	Result *BrowserTTLResponseResult `json:"result" validate:"required"`
}


// UnmarshalBrowserTTLResponse unmarshals an instance of BrowserTTLResponse from the specified map of raw messages.
func UnmarshalBrowserTTLResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(BrowserTTLResponse)
	err = core.UnmarshalPrimitive(m, "success", &obj.Success)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "errors", &obj.Errors)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "messages", &obj.Messages)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "result", &obj.Result, UnmarshalBrowserTTLResponseResult)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CacheLevelResponse : cache level response.
type CacheLevelResponse struct {
	// success response.
	Success *bool `json:"success" validate:"required"`

	// errors.
	Errors [][]string `json:"errors" validate:"required"`

	// messages.
	Messages [][]string `json:"messages" validate:"required"`

	// result.
	Result *CacheLevelResponseResult `json:"result" validate:"required"`
}


// UnmarshalCacheLevelResponse unmarshals an instance of CacheLevelResponse from the specified map of raw messages.
func UnmarshalCacheLevelResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CacheLevelResponse)
	err = core.UnmarshalPrimitive(m, "success", &obj.Success)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "errors", &obj.Errors)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "messages", &obj.Messages)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "result", &obj.Result, UnmarshalCacheLevelResponseResult)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DeveopmentModeResponse : development mode response.
type DeveopmentModeResponse struct {
	// success response.
	Success *bool `json:"success" validate:"required"`

	// errors.
	Errors [][]string `json:"errors" validate:"required"`

	// messages.
	Messages [][]string `json:"messages" validate:"required"`

	// result object.
	Result *DeveopmentModeResponseResult `json:"result" validate:"required"`
}


// UnmarshalDeveopmentModeResponse unmarshals an instance of DeveopmentModeResponse from the specified map of raw messages.
func UnmarshalDeveopmentModeResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DeveopmentModeResponse)
	err = core.UnmarshalPrimitive(m, "success", &obj.Success)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "errors", &obj.Errors)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "messages", &obj.Messages)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "result", &obj.Result, UnmarshalDeveopmentModeResponseResult)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// EnableQueryStringSortResponse : sort query string response.
type EnableQueryStringSortResponse struct {
	// success response true/false.
	Success *bool `json:"success" validate:"required"`

	// errors.
	Errors [][]string `json:"errors" validate:"required"`

	// messages.
	Messages [][]string `json:"messages" validate:"required"`

	// result of sort query string.
	Result *EnableQueryStringSortResponseResult `json:"result" validate:"required"`
}


// UnmarshalEnableQueryStringSortResponse unmarshals an instance of EnableQueryStringSortResponse from the specified map of raw messages.
func UnmarshalEnableQueryStringSortResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(EnableQueryStringSortResponse)
	err = core.UnmarshalPrimitive(m, "success", &obj.Success)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "errors", &obj.Errors)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "messages", &obj.Messages)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "result", &obj.Result, UnmarshalEnableQueryStringSortResponseResult)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// PurgeAllResponse : purge all response.
type PurgeAllResponse struct {
	// success response.
	Success *bool `json:"success" validate:"required"`

	// errors.
	Errors [][]string `json:"errors" validate:"required"`

	// messages.
	Messages [][]string `json:"messages" validate:"required"`

	// purge object.
	Result *PurgeAllResponseResult `json:"result" validate:"required"`
}


// UnmarshalPurgeAllResponse unmarshals an instance of PurgeAllResponse from the specified map of raw messages.
func UnmarshalPurgeAllResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PurgeAllResponse)
	err = core.UnmarshalPrimitive(m, "success", &obj.Success)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "errors", &obj.Errors)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "messages", &obj.Messages)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "result", &obj.Result, UnmarshalPurgeAllResponseResult)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}
