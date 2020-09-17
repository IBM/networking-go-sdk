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
 * IBM OpenAPI SDK Code Generator Version: 3.12.0-64fe8d3f-20200820-144050
 */
 

// Package globalloadbalancersv1 : Operations and models for the GlobalLoadBalancersV1 service
package globalloadbalancersv1

import (
	"encoding/json"
	"fmt"
	"github.com/IBM/go-sdk-core/v4/core"
	common "github.com/IBM/networking-go-sdk/common"
	"reflect"
)

// GlobalLoadBalancersV1 : Global Load Balancers
//
// Version: 1.0.0
type GlobalLoadBalancersV1 struct {
	Service *core.BaseService
}

// DefaultServiceURL is the default URL to make service requests to.
const DefaultServiceURL = "https://api.dns-svcs.cloud.ibm.com/v1"

// DefaultServiceName is the default key used to find external configuration information.
const DefaultServiceName = "global_load_balancers"

// GlobalLoadBalancersV1Options : Service options
type GlobalLoadBalancersV1Options struct {
	ServiceName   string
	URL           string
	Authenticator core.Authenticator
}

// NewGlobalLoadBalancersV1UsingExternalConfig : constructs an instance of GlobalLoadBalancersV1 with passed in options and external configuration.
func NewGlobalLoadBalancersV1UsingExternalConfig(options *GlobalLoadBalancersV1Options) (globalLoadBalancers *GlobalLoadBalancersV1, err error) {
	if options.ServiceName == "" {
		options.ServiceName = DefaultServiceName
	}

	if options.Authenticator == nil {
		options.Authenticator, err = core.GetAuthenticatorFromEnvironment(options.ServiceName)
		if err != nil {
			return
		}
	}

	globalLoadBalancers, err = NewGlobalLoadBalancersV1(options)
	if err != nil {
		return
	}

	err = globalLoadBalancers.Service.ConfigureService(options.ServiceName)
	if err != nil {
		return
	}

	if options.URL != "" {
		err = globalLoadBalancers.Service.SetServiceURL(options.URL)
	}
	return
}

// NewGlobalLoadBalancersV1 : constructs an instance of GlobalLoadBalancersV1 with passed in options.
func NewGlobalLoadBalancersV1(options *GlobalLoadBalancersV1Options) (service *GlobalLoadBalancersV1, err error) {
	serviceOptions := &core.ServiceOptions{
		URL:           DefaultServiceURL,
		Authenticator: options.Authenticator,
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

	service = &GlobalLoadBalancersV1{
		Service: baseService,
	}

	return
}

// SetServiceURL sets the service URL
func (globalLoadBalancers *GlobalLoadBalancersV1) SetServiceURL(url string) error {
	return globalLoadBalancers.Service.SetServiceURL(url)
}

// ListLoadBalancers : List load balancers
// List the Global Load Balancers for a given DNS zone.
func (globalLoadBalancers *GlobalLoadBalancersV1) ListLoadBalancers(listLoadBalancersOptions *ListLoadBalancersOptions) (result *ListLoadBalancers, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listLoadBalancersOptions, "listLoadBalancersOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listLoadBalancersOptions, "listLoadBalancersOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"instances", "dnszones", "load_balancers"}
	pathParameters := []string{*listLoadBalancersOptions.InstanceID, *listLoadBalancersOptions.DnszoneID}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(globalLoadBalancers.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range listLoadBalancersOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("global_load_balancers", "V1", "ListLoadBalancers")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if listLoadBalancersOptions.XCorrelationID != nil {
		builder.AddHeader("X-Correlation-ID", fmt.Sprint(*listLoadBalancersOptions.XCorrelationID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = globalLoadBalancers.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalListLoadBalancers)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// CreateLoadBalancer : Create a load balancer
// Create a load balancer for a given DNS zone.
func (globalLoadBalancers *GlobalLoadBalancersV1) CreateLoadBalancer(createLoadBalancerOptions *CreateLoadBalancerOptions) (result *LoadBalancer, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createLoadBalancerOptions, "createLoadBalancerOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createLoadBalancerOptions, "createLoadBalancerOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"instances", "dnszones", "load_balancers"}
	pathParameters := []string{*createLoadBalancerOptions.InstanceID, *createLoadBalancerOptions.DnszoneID}

	builder := core.NewRequestBuilder(core.POST)
	_, err = builder.ConstructHTTPURL(globalLoadBalancers.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range createLoadBalancerOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("global_load_balancers", "V1", "CreateLoadBalancer")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if createLoadBalancerOptions.XCorrelationID != nil {
		builder.AddHeader("X-Correlation-ID", fmt.Sprint(*createLoadBalancerOptions.XCorrelationID))
	}

	body := make(map[string]interface{})
	if createLoadBalancerOptions.Name != nil {
		body["name"] = createLoadBalancerOptions.Name
	}
	if createLoadBalancerOptions.Description != nil {
		body["description"] = createLoadBalancerOptions.Description
	}
	if createLoadBalancerOptions.Enabled != nil {
		body["enabled"] = createLoadBalancerOptions.Enabled
	}
	if createLoadBalancerOptions.TTL != nil {
		body["ttl"] = createLoadBalancerOptions.TTL
	}
	if createLoadBalancerOptions.FallbackPool != nil {
		body["fallback_pool"] = createLoadBalancerOptions.FallbackPool
	}
	if createLoadBalancerOptions.DefaultPools != nil {
		body["default_pools"] = createLoadBalancerOptions.DefaultPools
	}
	if createLoadBalancerOptions.AzPools != nil {
		body["az_pools"] = createLoadBalancerOptions.AzPools
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
	response, err = globalLoadBalancers.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalLoadBalancer)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// DeleteLoadBalancer : Delete a load balancer
// Delete a load balancer.
func (globalLoadBalancers *GlobalLoadBalancersV1) DeleteLoadBalancer(deleteLoadBalancerOptions *DeleteLoadBalancerOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteLoadBalancerOptions, "deleteLoadBalancerOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteLoadBalancerOptions, "deleteLoadBalancerOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"instances", "dnszones", "load_balancers"}
	pathParameters := []string{*deleteLoadBalancerOptions.InstanceID, *deleteLoadBalancerOptions.DnszoneID, *deleteLoadBalancerOptions.LbID}

	builder := core.NewRequestBuilder(core.DELETE)
	_, err = builder.ConstructHTTPURL(globalLoadBalancers.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteLoadBalancerOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("global_load_balancers", "V1", "DeleteLoadBalancer")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	if deleteLoadBalancerOptions.XCorrelationID != nil {
		builder.AddHeader("X-Correlation-ID", fmt.Sprint(*deleteLoadBalancerOptions.XCorrelationID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = globalLoadBalancers.Service.Request(request, nil)

	return
}

// GetLoadBalancer : Get a load balancer
// Get details of a load balancer.
func (globalLoadBalancers *GlobalLoadBalancersV1) GetLoadBalancer(getLoadBalancerOptions *GetLoadBalancerOptions) (result *LoadBalancer, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getLoadBalancerOptions, "getLoadBalancerOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getLoadBalancerOptions, "getLoadBalancerOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"instances", "dnszones", "load_balancers"}
	pathParameters := []string{*getLoadBalancerOptions.InstanceID, *getLoadBalancerOptions.DnszoneID, *getLoadBalancerOptions.LbID}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(globalLoadBalancers.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range getLoadBalancerOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("global_load_balancers", "V1", "GetLoadBalancer")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if getLoadBalancerOptions.XCorrelationID != nil {
		builder.AddHeader("X-Correlation-ID", fmt.Sprint(*getLoadBalancerOptions.XCorrelationID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = globalLoadBalancers.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalLoadBalancer)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// UpdateLoadBalancer : Update the properties of a load balancer
// Update the properties of a load balancer.
func (globalLoadBalancers *GlobalLoadBalancersV1) UpdateLoadBalancer(updateLoadBalancerOptions *UpdateLoadBalancerOptions) (result *LoadBalancer, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateLoadBalancerOptions, "updateLoadBalancerOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateLoadBalancerOptions, "updateLoadBalancerOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"instances", "dnszones", "load_balancers"}
	pathParameters := []string{*updateLoadBalancerOptions.InstanceID, *updateLoadBalancerOptions.DnszoneID, *updateLoadBalancerOptions.LbID}

	builder := core.NewRequestBuilder(core.PUT)
	_, err = builder.ConstructHTTPURL(globalLoadBalancers.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateLoadBalancerOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("global_load_balancers", "V1", "UpdateLoadBalancer")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if updateLoadBalancerOptions.XCorrelationID != nil {
		builder.AddHeader("X-Correlation-ID", fmt.Sprint(*updateLoadBalancerOptions.XCorrelationID))
	}

	body := make(map[string]interface{})
	if updateLoadBalancerOptions.Name != nil {
		body["name"] = updateLoadBalancerOptions.Name
	}
	if updateLoadBalancerOptions.Description != nil {
		body["description"] = updateLoadBalancerOptions.Description
	}
	if updateLoadBalancerOptions.Enabled != nil {
		body["enabled"] = updateLoadBalancerOptions.Enabled
	}
	if updateLoadBalancerOptions.TTL != nil {
		body["ttl"] = updateLoadBalancerOptions.TTL
	}
	if updateLoadBalancerOptions.FallbackPool != nil {
		body["fallback_pool"] = updateLoadBalancerOptions.FallbackPool
	}
	if updateLoadBalancerOptions.DefaultPools != nil {
		body["default_pools"] = updateLoadBalancerOptions.DefaultPools
	}
	if updateLoadBalancerOptions.AzPools != nil {
		body["az_pools"] = updateLoadBalancerOptions.AzPools
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
	response, err = globalLoadBalancers.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalLoadBalancer)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// ListPools : List load balancer pools
// List the load balancer pools.
func (globalLoadBalancers *GlobalLoadBalancersV1) ListPools(listPoolsOptions *ListPoolsOptions) (result *ListPools, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listPoolsOptions, "listPoolsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listPoolsOptions, "listPoolsOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"instances", "pools"}
	pathParameters := []string{*listPoolsOptions.InstanceID}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(globalLoadBalancers.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range listPoolsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("global_load_balancers", "V1", "ListPools")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if listPoolsOptions.XCorrelationID != nil {
		builder.AddHeader("X-Correlation-ID", fmt.Sprint(*listPoolsOptions.XCorrelationID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = globalLoadBalancers.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalListPools)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// CreatePool : Create a load balancer pool
// Create a load balancer pool.
func (globalLoadBalancers *GlobalLoadBalancersV1) CreatePool(createPoolOptions *CreatePoolOptions) (result *Pool, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createPoolOptions, "createPoolOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createPoolOptions, "createPoolOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"instances", "pools"}
	pathParameters := []string{*createPoolOptions.InstanceID}

	builder := core.NewRequestBuilder(core.POST)
	_, err = builder.ConstructHTTPURL(globalLoadBalancers.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range createPoolOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("global_load_balancers", "V1", "CreatePool")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if createPoolOptions.XCorrelationID != nil {
		builder.AddHeader("X-Correlation-ID", fmt.Sprint(*createPoolOptions.XCorrelationID))
	}

	body := make(map[string]interface{})
	if createPoolOptions.Name != nil {
		body["name"] = createPoolOptions.Name
	}
	if createPoolOptions.Description != nil {
		body["description"] = createPoolOptions.Description
	}
	if createPoolOptions.Enabled != nil {
		body["enabled"] = createPoolOptions.Enabled
	}
	if createPoolOptions.HealthyOriginsThreshold != nil {
		body["healthy_origins_threshold"] = createPoolOptions.HealthyOriginsThreshold
	}
	if createPoolOptions.Origins != nil {
		body["origins"] = createPoolOptions.Origins
	}
	if createPoolOptions.Monitor != nil {
		body["monitor"] = createPoolOptions.Monitor
	}
	if createPoolOptions.NotificationChannel != nil {
		body["notification_channel"] = createPoolOptions.NotificationChannel
	}
	if createPoolOptions.HealthcheckRegion != nil {
		body["healthcheck_region"] = createPoolOptions.HealthcheckRegion
	}
	if createPoolOptions.HealthcheckSubnets != nil {
		body["healthcheck_subnets"] = createPoolOptions.HealthcheckSubnets
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
	response, err = globalLoadBalancers.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalPool)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// DeletePool : Delete a load balancer pool
// Delete a load balancer pool.
func (globalLoadBalancers *GlobalLoadBalancersV1) DeletePool(deletePoolOptions *DeletePoolOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deletePoolOptions, "deletePoolOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deletePoolOptions, "deletePoolOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"instances", "pools"}
	pathParameters := []string{*deletePoolOptions.InstanceID, *deletePoolOptions.PoolID}

	builder := core.NewRequestBuilder(core.DELETE)
	_, err = builder.ConstructHTTPURL(globalLoadBalancers.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range deletePoolOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("global_load_balancers", "V1", "DeletePool")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	if deletePoolOptions.XCorrelationID != nil {
		builder.AddHeader("X-Correlation-ID", fmt.Sprint(*deletePoolOptions.XCorrelationID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = globalLoadBalancers.Service.Request(request, nil)

	return
}

// GetPool : Get a load balancer pool
// Get details of a load balancer pool.
func (globalLoadBalancers *GlobalLoadBalancersV1) GetPool(getPoolOptions *GetPoolOptions) (result *Pool, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getPoolOptions, "getPoolOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getPoolOptions, "getPoolOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"instances", "pools"}
	pathParameters := []string{*getPoolOptions.InstanceID, *getPoolOptions.PoolID}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(globalLoadBalancers.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range getPoolOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("global_load_balancers", "V1", "GetPool")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if getPoolOptions.XCorrelationID != nil {
		builder.AddHeader("X-Correlation-ID", fmt.Sprint(*getPoolOptions.XCorrelationID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = globalLoadBalancers.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalPool)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// UpdatePool : Update the properties of a load balancer pool
// Update the properties of a load balancer pool.
func (globalLoadBalancers *GlobalLoadBalancersV1) UpdatePool(updatePoolOptions *UpdatePoolOptions) (result *Pool, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updatePoolOptions, "updatePoolOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updatePoolOptions, "updatePoolOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"instances", "pools"}
	pathParameters := []string{*updatePoolOptions.InstanceID, *updatePoolOptions.PoolID}

	builder := core.NewRequestBuilder(core.PUT)
	_, err = builder.ConstructHTTPURL(globalLoadBalancers.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range updatePoolOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("global_load_balancers", "V1", "UpdatePool")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if updatePoolOptions.XCorrelationID != nil {
		builder.AddHeader("X-Correlation-ID", fmt.Sprint(*updatePoolOptions.XCorrelationID))
	}

	body := make(map[string]interface{})
	if updatePoolOptions.Name != nil {
		body["name"] = updatePoolOptions.Name
	}
	if updatePoolOptions.Description != nil {
		body["description"] = updatePoolOptions.Description
	}
	if updatePoolOptions.Enabled != nil {
		body["enabled"] = updatePoolOptions.Enabled
	}
	if updatePoolOptions.HealthyOriginsThreshold != nil {
		body["healthy_origins_threshold"] = updatePoolOptions.HealthyOriginsThreshold
	}
	if updatePoolOptions.Origins != nil {
		body["origins"] = updatePoolOptions.Origins
	}
	if updatePoolOptions.Monitor != nil {
		body["monitor"] = updatePoolOptions.Monitor
	}
	if updatePoolOptions.NotificationChannel != nil {
		body["notification_channel"] = updatePoolOptions.NotificationChannel
	}
	if updatePoolOptions.HealthcheckRegion != nil {
		body["healthcheck_region"] = updatePoolOptions.HealthcheckRegion
	}
	if updatePoolOptions.HealthcheckSubnets != nil {
		body["healthcheck_subnets"] = updatePoolOptions.HealthcheckSubnets
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
	response, err = globalLoadBalancers.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalPool)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// ListMonitors : List load balancer monitors
// List the load balancer monitors.
func (globalLoadBalancers *GlobalLoadBalancersV1) ListMonitors(listMonitorsOptions *ListMonitorsOptions) (result *ListMonitors, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listMonitorsOptions, "listMonitorsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listMonitorsOptions, "listMonitorsOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"instances", "monitors"}
	pathParameters := []string{*listMonitorsOptions.InstanceID}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(globalLoadBalancers.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range listMonitorsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("global_load_balancers", "V1", "ListMonitors")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if listMonitorsOptions.XCorrelationID != nil {
		builder.AddHeader("X-Correlation-ID", fmt.Sprint(*listMonitorsOptions.XCorrelationID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = globalLoadBalancers.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalListMonitors)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// CreateMonitor : Create a load balancer monitor
// Create a load balancer monitor.
func (globalLoadBalancers *GlobalLoadBalancersV1) CreateMonitor(createMonitorOptions *CreateMonitorOptions) (result *Monitor, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createMonitorOptions, "createMonitorOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createMonitorOptions, "createMonitorOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"instances", "monitors"}
	pathParameters := []string{*createMonitorOptions.InstanceID}

	builder := core.NewRequestBuilder(core.POST)
	_, err = builder.ConstructHTTPURL(globalLoadBalancers.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range createMonitorOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("global_load_balancers", "V1", "CreateMonitor")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if createMonitorOptions.XCorrelationID != nil {
		builder.AddHeader("X-Correlation-ID", fmt.Sprint(*createMonitorOptions.XCorrelationID))
	}

	body := make(map[string]interface{})
	if createMonitorOptions.Name != nil {
		body["name"] = createMonitorOptions.Name
	}
	if createMonitorOptions.Description != nil {
		body["description"] = createMonitorOptions.Description
	}
	if createMonitorOptions.Type != nil {
		body["type"] = createMonitorOptions.Type
	}
	if createMonitorOptions.Port != nil {
		body["port"] = createMonitorOptions.Port
	}
	if createMonitorOptions.Interval != nil {
		body["interval"] = createMonitorOptions.Interval
	}
	if createMonitorOptions.Retries != nil {
		body["retries"] = createMonitorOptions.Retries
	}
	if createMonitorOptions.Timeout != nil {
		body["timeout"] = createMonitorOptions.Timeout
	}
	if createMonitorOptions.Method != nil {
		body["method"] = createMonitorOptions.Method
	}
	if createMonitorOptions.Path != nil {
		body["path"] = createMonitorOptions.Path
	}
	if createMonitorOptions.HeadersVar != nil {
		body["headers"] = createMonitorOptions.HeadersVar
	}
	if createMonitorOptions.AllowInsecure != nil {
		body["allow_insecure"] = createMonitorOptions.AllowInsecure
	}
	if createMonitorOptions.ExpectedCodes != nil {
		body["expected_codes"] = createMonitorOptions.ExpectedCodes
	}
	if createMonitorOptions.ExpectedBody != nil {
		body["expected_body"] = createMonitorOptions.ExpectedBody
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
	response, err = globalLoadBalancers.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalMonitor)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// DeleteMonitor : Delete a load balancer monitor
// Delete a load balancer monitor.
func (globalLoadBalancers *GlobalLoadBalancersV1) DeleteMonitor(deleteMonitorOptions *DeleteMonitorOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteMonitorOptions, "deleteMonitorOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteMonitorOptions, "deleteMonitorOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"instances", "monitors"}
	pathParameters := []string{*deleteMonitorOptions.InstanceID, *deleteMonitorOptions.MonitorID}

	builder := core.NewRequestBuilder(core.DELETE)
	_, err = builder.ConstructHTTPURL(globalLoadBalancers.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteMonitorOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("global_load_balancers", "V1", "DeleteMonitor")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	if deleteMonitorOptions.XCorrelationID != nil {
		builder.AddHeader("X-Correlation-ID", fmt.Sprint(*deleteMonitorOptions.XCorrelationID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = globalLoadBalancers.Service.Request(request, nil)

	return
}

// GetMonitor : Get a load balancer monitor
// Get details of a load balancer monitor.
func (globalLoadBalancers *GlobalLoadBalancersV1) GetMonitor(getMonitorOptions *GetMonitorOptions) (result *Monitor, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getMonitorOptions, "getMonitorOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getMonitorOptions, "getMonitorOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"instances", "monitors"}
	pathParameters := []string{*getMonitorOptions.InstanceID, *getMonitorOptions.MonitorID}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(globalLoadBalancers.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range getMonitorOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("global_load_balancers", "V1", "GetMonitor")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if getMonitorOptions.XCorrelationID != nil {
		builder.AddHeader("X-Correlation-ID", fmt.Sprint(*getMonitorOptions.XCorrelationID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = globalLoadBalancers.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalMonitor)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// UpdateMonitor : Update the properties of a load balancer monitor
// Update the properties of a load balancer monitor.
func (globalLoadBalancers *GlobalLoadBalancersV1) UpdateMonitor(updateMonitorOptions *UpdateMonitorOptions) (result *Monitor, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateMonitorOptions, "updateMonitorOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateMonitorOptions, "updateMonitorOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"instances", "monitors"}
	pathParameters := []string{*updateMonitorOptions.InstanceID, *updateMonitorOptions.MonitorID}

	builder := core.NewRequestBuilder(core.PUT)
	_, err = builder.ConstructHTTPURL(globalLoadBalancers.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateMonitorOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("global_load_balancers", "V1", "UpdateMonitor")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if updateMonitorOptions.XCorrelationID != nil {
		builder.AddHeader("X-Correlation-ID", fmt.Sprint(*updateMonitorOptions.XCorrelationID))
	}

	body := make(map[string]interface{})
	if updateMonitorOptions.Name != nil {
		body["name"] = updateMonitorOptions.Name
	}
	if updateMonitorOptions.Description != nil {
		body["description"] = updateMonitorOptions.Description
	}
	if updateMonitorOptions.Type != nil {
		body["type"] = updateMonitorOptions.Type
	}
	if updateMonitorOptions.Port != nil {
		body["port"] = updateMonitorOptions.Port
	}
	if updateMonitorOptions.Interval != nil {
		body["interval"] = updateMonitorOptions.Interval
	}
	if updateMonitorOptions.Retries != nil {
		body["retries"] = updateMonitorOptions.Retries
	}
	if updateMonitorOptions.Timeout != nil {
		body["timeout"] = updateMonitorOptions.Timeout
	}
	if updateMonitorOptions.Method != nil {
		body["method"] = updateMonitorOptions.Method
	}
	if updateMonitorOptions.Path != nil {
		body["path"] = updateMonitorOptions.Path
	}
	if updateMonitorOptions.HeadersVar != nil {
		body["headers"] = updateMonitorOptions.HeadersVar
	}
	if updateMonitorOptions.AllowInsecure != nil {
		body["allow_insecure"] = updateMonitorOptions.AllowInsecure
	}
	if updateMonitorOptions.ExpectedCodes != nil {
		body["expected_codes"] = updateMonitorOptions.ExpectedCodes
	}
	if updateMonitorOptions.ExpectedBody != nil {
		body["expected_body"] = updateMonitorOptions.ExpectedBody
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
	response, err = globalLoadBalancers.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalMonitor)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// CreateLoadBalancerOptions : The CreateLoadBalancer options.
type CreateLoadBalancerOptions struct {
	// The unique identifier of a service instance.
	InstanceID *string `json:"instance_id" validate:"required"`

	// The unique identifier of a DNS zone.
	DnszoneID *string `json:"dnszone_id" validate:"required"`

	// Name of the load balancer.
	Name *string `json:"name,omitempty"`

	// Descriptive text of the load balancer.
	Description *string `json:"description,omitempty"`

	// Whether the load balancer is enabled.
	Enabled *bool `json:"enabled,omitempty"`

	// Time to live in second.
	TTL *int64 `json:"ttl,omitempty"`

	// The pool ID to use when all other pools are detected as unhealthy.
	FallbackPool *string `json:"fallback_pool,omitempty"`

	// A list of pool IDs ordered by their failover priority. Pools defined here are used by default, or when region_pools
	// are not configured for a given region.
	DefaultPools []string `json:"default_pools,omitempty"`

	// Map availability zones to pool ID's.
	AzPools []LoadBalancerAzPoolsItem `json:"az_pools,omitempty"`

	// Uniquely identifying a request.
	XCorrelationID *string `json:"X-Correlation-ID,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCreateLoadBalancerOptions : Instantiate CreateLoadBalancerOptions
func (*GlobalLoadBalancersV1) NewCreateLoadBalancerOptions(instanceID string, dnszoneID string) *CreateLoadBalancerOptions {
	return &CreateLoadBalancerOptions{
		InstanceID: core.StringPtr(instanceID),
		DnszoneID: core.StringPtr(dnszoneID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (options *CreateLoadBalancerOptions) SetInstanceID(instanceID string) *CreateLoadBalancerOptions {
	options.InstanceID = core.StringPtr(instanceID)
	return options
}

// SetDnszoneID : Allow user to set DnszoneID
func (options *CreateLoadBalancerOptions) SetDnszoneID(dnszoneID string) *CreateLoadBalancerOptions {
	options.DnszoneID = core.StringPtr(dnszoneID)
	return options
}

// SetName : Allow user to set Name
func (options *CreateLoadBalancerOptions) SetName(name string) *CreateLoadBalancerOptions {
	options.Name = core.StringPtr(name)
	return options
}

// SetDescription : Allow user to set Description
func (options *CreateLoadBalancerOptions) SetDescription(description string) *CreateLoadBalancerOptions {
	options.Description = core.StringPtr(description)
	return options
}

// SetEnabled : Allow user to set Enabled
func (options *CreateLoadBalancerOptions) SetEnabled(enabled bool) *CreateLoadBalancerOptions {
	options.Enabled = core.BoolPtr(enabled)
	return options
}

// SetTTL : Allow user to set TTL
func (options *CreateLoadBalancerOptions) SetTTL(ttl int64) *CreateLoadBalancerOptions {
	options.TTL = core.Int64Ptr(ttl)
	return options
}

// SetFallbackPool : Allow user to set FallbackPool
func (options *CreateLoadBalancerOptions) SetFallbackPool(fallbackPool string) *CreateLoadBalancerOptions {
	options.FallbackPool = core.StringPtr(fallbackPool)
	return options
}

// SetDefaultPools : Allow user to set DefaultPools
func (options *CreateLoadBalancerOptions) SetDefaultPools(defaultPools []string) *CreateLoadBalancerOptions {
	options.DefaultPools = defaultPools
	return options
}

// SetAzPools : Allow user to set AzPools
func (options *CreateLoadBalancerOptions) SetAzPools(azPools []LoadBalancerAzPoolsItem) *CreateLoadBalancerOptions {
	options.AzPools = azPools
	return options
}

// SetXCorrelationID : Allow user to set XCorrelationID
func (options *CreateLoadBalancerOptions) SetXCorrelationID(xCorrelationID string) *CreateLoadBalancerOptions {
	options.XCorrelationID = core.StringPtr(xCorrelationID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *CreateLoadBalancerOptions) SetHeaders(param map[string]string) *CreateLoadBalancerOptions {
	options.Headers = param
	return options
}

// CreateMonitorOptions : The CreateMonitor options.
type CreateMonitorOptions struct {
	// The unique identifier of a service instance.
	InstanceID *string `json:"instance_id" validate:"required"`

	// The name of the load balancer monitor.
	Name *string `json:"name,omitempty"`

	// Descriptive text of the load balancer monitor.
	Description *string `json:"description,omitempty"`

	// The protocol to use for the health check. Currently supported protocols are 'HTTP','HTTPS' and 'TCP'.
	Type *string `json:"type,omitempty"`

	// Port number to connect to for the health check. Required for TCP checks. HTTP and HTTPS checks should only define
	// the port when using a non-standard port (HTTP: default 80, HTTPS: default 443).
	Port *int64 `json:"port,omitempty"`

	// The interval between each health check. Shorter intervals may improve failover time, but will increase load on the
	// origins as we check from multiple locations.
	Interval *int64 `json:"interval,omitempty"`

	// The number of retries to attempt in case of a timeout before marking the origin as unhealthy. Retries are attempted
	// immediately.
	Retries *int64 `json:"retries,omitempty"`

	// The timeout (in seconds) before marking the health check as failed.
	Timeout *int64 `json:"timeout,omitempty"`

	// The method to use for the health check applicable to HTTP/HTTPS based checks, the default value is 'GET'.
	Method *string `json:"method,omitempty"`

	// The endpoint path to health check against. This parameter is only valid for HTTP and HTTPS monitors.
	Path *string `json:"path,omitempty"`

	// The HTTP request headers to send in the health check. It is recommended you set a Host header by default. The
	// User-Agent header cannot be overridden. This parameter is only valid for HTTP and HTTPS monitors.
	HeadersVar []HealthcheckHeader `json:"headers,omitempty"`

	// Do not validate the certificate when monitor use HTTPS. This parameter is currently only valid for HTTPS monitors.
	AllowInsecure *bool `json:"allow_insecure,omitempty"`

	// The expected HTTP response code or code range of the health check. This parameter is only valid for HTTP and HTTPS
	// monitors.
	ExpectedCodes *string `json:"expected_codes,omitempty"`

	// A case-insensitive sub-string to look for in the response body. If this string is not found, the origin will be
	// marked as unhealthy. This parameter is only valid for HTTP and HTTPS monitors.
	ExpectedBody *string `json:"expected_body,omitempty"`

	// Uniquely identifying a request.
	XCorrelationID *string `json:"X-Correlation-ID,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the CreateMonitorOptions.Type property.
// The protocol to use for the health check. Currently supported protocols are 'HTTP','HTTPS' and 'TCP'.
const (
	CreateMonitorOptions_Type_Http = "HTTP"
	CreateMonitorOptions_Type_Https = "HTTPS"
	CreateMonitorOptions_Type_Tcp = "TCP"
)

// Constants associated with the CreateMonitorOptions.Method property.
// The method to use for the health check applicable to HTTP/HTTPS based checks, the default value is 'GET'.
const (
	CreateMonitorOptions_Method_Get = "GET"
	CreateMonitorOptions_Method_Head = "HEAD"
)

// NewCreateMonitorOptions : Instantiate CreateMonitorOptions
func (*GlobalLoadBalancersV1) NewCreateMonitorOptions(instanceID string) *CreateMonitorOptions {
	return &CreateMonitorOptions{
		InstanceID: core.StringPtr(instanceID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (options *CreateMonitorOptions) SetInstanceID(instanceID string) *CreateMonitorOptions {
	options.InstanceID = core.StringPtr(instanceID)
	return options
}

// SetName : Allow user to set Name
func (options *CreateMonitorOptions) SetName(name string) *CreateMonitorOptions {
	options.Name = core.StringPtr(name)
	return options
}

// SetDescription : Allow user to set Description
func (options *CreateMonitorOptions) SetDescription(description string) *CreateMonitorOptions {
	options.Description = core.StringPtr(description)
	return options
}

// SetType : Allow user to set Type
func (options *CreateMonitorOptions) SetType(typeVar string) *CreateMonitorOptions {
	options.Type = core.StringPtr(typeVar)
	return options
}

// SetPort : Allow user to set Port
func (options *CreateMonitorOptions) SetPort(port int64) *CreateMonitorOptions {
	options.Port = core.Int64Ptr(port)
	return options
}

// SetInterval : Allow user to set Interval
func (options *CreateMonitorOptions) SetInterval(interval int64) *CreateMonitorOptions {
	options.Interval = core.Int64Ptr(interval)
	return options
}

// SetRetries : Allow user to set Retries
func (options *CreateMonitorOptions) SetRetries(retries int64) *CreateMonitorOptions {
	options.Retries = core.Int64Ptr(retries)
	return options
}

// SetTimeout : Allow user to set Timeout
func (options *CreateMonitorOptions) SetTimeout(timeout int64) *CreateMonitorOptions {
	options.Timeout = core.Int64Ptr(timeout)
	return options
}

// SetMethod : Allow user to set Method
func (options *CreateMonitorOptions) SetMethod(method string) *CreateMonitorOptions {
	options.Method = core.StringPtr(method)
	return options
}

// SetPath : Allow user to set Path
func (options *CreateMonitorOptions) SetPath(path string) *CreateMonitorOptions {
	options.Path = core.StringPtr(path)
	return options
}

// SetHeadersVar : Allow user to set HeadersVar
func (options *CreateMonitorOptions) SetHeadersVar(headersVar []HealthcheckHeader) *CreateMonitorOptions {
	options.HeadersVar = headersVar
	return options
}

// SetAllowInsecure : Allow user to set AllowInsecure
func (options *CreateMonitorOptions) SetAllowInsecure(allowInsecure bool) *CreateMonitorOptions {
	options.AllowInsecure = core.BoolPtr(allowInsecure)
	return options
}

// SetExpectedCodes : Allow user to set ExpectedCodes
func (options *CreateMonitorOptions) SetExpectedCodes(expectedCodes string) *CreateMonitorOptions {
	options.ExpectedCodes = core.StringPtr(expectedCodes)
	return options
}

// SetExpectedBody : Allow user to set ExpectedBody
func (options *CreateMonitorOptions) SetExpectedBody(expectedBody string) *CreateMonitorOptions {
	options.ExpectedBody = core.StringPtr(expectedBody)
	return options
}

// SetXCorrelationID : Allow user to set XCorrelationID
func (options *CreateMonitorOptions) SetXCorrelationID(xCorrelationID string) *CreateMonitorOptions {
	options.XCorrelationID = core.StringPtr(xCorrelationID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *CreateMonitorOptions) SetHeaders(param map[string]string) *CreateMonitorOptions {
	options.Headers = param
	return options
}

// CreatePoolOptions : The CreatePool options.
type CreatePoolOptions struct {
	// The unique identifier of a service instance.
	InstanceID *string `json:"instance_id" validate:"required"`

	// Name of the load balancer pool.
	Name *string `json:"name,omitempty"`

	// Descriptive text of the load balancer pool.
	Description *string `json:"description,omitempty"`

	// Whether the load balancer pool is enabled.
	Enabled *bool `json:"enabled,omitempty"`

	// The minimum number of origins that must be healthy for this pool to serve traffic. If the number of healthy origins
	// falls below this number, the pool will be marked unhealthy and we will failover to the next available pool.
	HealthyOriginsThreshold *int64 `json:"healthy_origins_threshold,omitempty"`

	// The list of origins within this pool. Traffic directed at this pool is balanced across all currently healthy
	// origins, provided the pool itself is healthy.
	Origins []OriginInput `json:"origins,omitempty"`

	// The ID of the load balancer monitor to be associated to this pool.
	Monitor *string `json:"monitor,omitempty"`

	// The notification channel.
	NotificationChannel *string `json:"notification_channel,omitempty"`

	// Health check region of VSIs.
	HealthcheckRegion *string `json:"healthcheck_region,omitempty"`

	// Health check subnet IDs of VSIs.
	HealthcheckSubnets []string `json:"healthcheck_subnets,omitempty"`

	// Uniquely identifying a request.
	XCorrelationID *string `json:"X-Correlation-ID,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the CreatePoolOptions.HealthcheckRegion property.
// Health check region of VSIs.
const (
	CreatePoolOptions_HealthcheckRegion_AuSyd = "au-syd"
	CreatePoolOptions_HealthcheckRegion_EuDu = "eu-du"
	CreatePoolOptions_HealthcheckRegion_EuGb = "eu-gb"
	CreatePoolOptions_HealthcheckRegion_JpTok = "jp-tok"
	CreatePoolOptions_HealthcheckRegion_UsEast = "us-east"
	CreatePoolOptions_HealthcheckRegion_UsSouth = "us-south"
)

// NewCreatePoolOptions : Instantiate CreatePoolOptions
func (*GlobalLoadBalancersV1) NewCreatePoolOptions(instanceID string) *CreatePoolOptions {
	return &CreatePoolOptions{
		InstanceID: core.StringPtr(instanceID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (options *CreatePoolOptions) SetInstanceID(instanceID string) *CreatePoolOptions {
	options.InstanceID = core.StringPtr(instanceID)
	return options
}

// SetName : Allow user to set Name
func (options *CreatePoolOptions) SetName(name string) *CreatePoolOptions {
	options.Name = core.StringPtr(name)
	return options
}

// SetDescription : Allow user to set Description
func (options *CreatePoolOptions) SetDescription(description string) *CreatePoolOptions {
	options.Description = core.StringPtr(description)
	return options
}

// SetEnabled : Allow user to set Enabled
func (options *CreatePoolOptions) SetEnabled(enabled bool) *CreatePoolOptions {
	options.Enabled = core.BoolPtr(enabled)
	return options
}

// SetHealthyOriginsThreshold : Allow user to set HealthyOriginsThreshold
func (options *CreatePoolOptions) SetHealthyOriginsThreshold(healthyOriginsThreshold int64) *CreatePoolOptions {
	options.HealthyOriginsThreshold = core.Int64Ptr(healthyOriginsThreshold)
	return options
}

// SetOrigins : Allow user to set Origins
func (options *CreatePoolOptions) SetOrigins(origins []OriginInput) *CreatePoolOptions {
	options.Origins = origins
	return options
}

// SetMonitor : Allow user to set Monitor
func (options *CreatePoolOptions) SetMonitor(monitor string) *CreatePoolOptions {
	options.Monitor = core.StringPtr(monitor)
	return options
}

// SetNotificationChannel : Allow user to set NotificationChannel
func (options *CreatePoolOptions) SetNotificationChannel(notificationChannel string) *CreatePoolOptions {
	options.NotificationChannel = core.StringPtr(notificationChannel)
	return options
}

// SetHealthcheckRegion : Allow user to set HealthcheckRegion
func (options *CreatePoolOptions) SetHealthcheckRegion(healthcheckRegion string) *CreatePoolOptions {
	options.HealthcheckRegion = core.StringPtr(healthcheckRegion)
	return options
}

// SetHealthcheckSubnets : Allow user to set HealthcheckSubnets
func (options *CreatePoolOptions) SetHealthcheckSubnets(healthcheckSubnets []string) *CreatePoolOptions {
	options.HealthcheckSubnets = healthcheckSubnets
	return options
}

// SetXCorrelationID : Allow user to set XCorrelationID
func (options *CreatePoolOptions) SetXCorrelationID(xCorrelationID string) *CreatePoolOptions {
	options.XCorrelationID = core.StringPtr(xCorrelationID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *CreatePoolOptions) SetHeaders(param map[string]string) *CreatePoolOptions {
	options.Headers = param
	return options
}

// DeleteLoadBalancerOptions : The DeleteLoadBalancer options.
type DeleteLoadBalancerOptions struct {
	// The unique identifier of a service instance.
	InstanceID *string `json:"instance_id" validate:"required"`

	// The unique identifier of a DNS zone.
	DnszoneID *string `json:"dnszone_id" validate:"required"`

	// The unique identifier of a load balancer.
	LbID *string `json:"lb_id" validate:"required"`

	// Uniquely identifying a request.
	XCorrelationID *string `json:"X-Correlation-ID,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteLoadBalancerOptions : Instantiate DeleteLoadBalancerOptions
func (*GlobalLoadBalancersV1) NewDeleteLoadBalancerOptions(instanceID string, dnszoneID string, lbID string) *DeleteLoadBalancerOptions {
	return &DeleteLoadBalancerOptions{
		InstanceID: core.StringPtr(instanceID),
		DnszoneID: core.StringPtr(dnszoneID),
		LbID: core.StringPtr(lbID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (options *DeleteLoadBalancerOptions) SetInstanceID(instanceID string) *DeleteLoadBalancerOptions {
	options.InstanceID = core.StringPtr(instanceID)
	return options
}

// SetDnszoneID : Allow user to set DnszoneID
func (options *DeleteLoadBalancerOptions) SetDnszoneID(dnszoneID string) *DeleteLoadBalancerOptions {
	options.DnszoneID = core.StringPtr(dnszoneID)
	return options
}

// SetLbID : Allow user to set LbID
func (options *DeleteLoadBalancerOptions) SetLbID(lbID string) *DeleteLoadBalancerOptions {
	options.LbID = core.StringPtr(lbID)
	return options
}

// SetXCorrelationID : Allow user to set XCorrelationID
func (options *DeleteLoadBalancerOptions) SetXCorrelationID(xCorrelationID string) *DeleteLoadBalancerOptions {
	options.XCorrelationID = core.StringPtr(xCorrelationID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteLoadBalancerOptions) SetHeaders(param map[string]string) *DeleteLoadBalancerOptions {
	options.Headers = param
	return options
}

// DeleteMonitorOptions : The DeleteMonitor options.
type DeleteMonitorOptions struct {
	// The unique identifier of a service instance.
	InstanceID *string `json:"instance_id" validate:"required"`

	// The unique identifier of a load balancer monitor.
	MonitorID *string `json:"monitor_id" validate:"required"`

	// Uniquely identifying a request.
	XCorrelationID *string `json:"X-Correlation-ID,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteMonitorOptions : Instantiate DeleteMonitorOptions
func (*GlobalLoadBalancersV1) NewDeleteMonitorOptions(instanceID string, monitorID string) *DeleteMonitorOptions {
	return &DeleteMonitorOptions{
		InstanceID: core.StringPtr(instanceID),
		MonitorID: core.StringPtr(monitorID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (options *DeleteMonitorOptions) SetInstanceID(instanceID string) *DeleteMonitorOptions {
	options.InstanceID = core.StringPtr(instanceID)
	return options
}

// SetMonitorID : Allow user to set MonitorID
func (options *DeleteMonitorOptions) SetMonitorID(monitorID string) *DeleteMonitorOptions {
	options.MonitorID = core.StringPtr(monitorID)
	return options
}

// SetXCorrelationID : Allow user to set XCorrelationID
func (options *DeleteMonitorOptions) SetXCorrelationID(xCorrelationID string) *DeleteMonitorOptions {
	options.XCorrelationID = core.StringPtr(xCorrelationID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteMonitorOptions) SetHeaders(param map[string]string) *DeleteMonitorOptions {
	options.Headers = param
	return options
}

// DeletePoolOptions : The DeletePool options.
type DeletePoolOptions struct {
	// The unique identifier of a service instance.
	InstanceID *string `json:"instance_id" validate:"required"`

	// The unique identifier of a load balancer pool.
	PoolID *string `json:"pool_id" validate:"required"`

	// Uniquely identifying a request.
	XCorrelationID *string `json:"X-Correlation-ID,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeletePoolOptions : Instantiate DeletePoolOptions
func (*GlobalLoadBalancersV1) NewDeletePoolOptions(instanceID string, poolID string) *DeletePoolOptions {
	return &DeletePoolOptions{
		InstanceID: core.StringPtr(instanceID),
		PoolID: core.StringPtr(poolID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (options *DeletePoolOptions) SetInstanceID(instanceID string) *DeletePoolOptions {
	options.InstanceID = core.StringPtr(instanceID)
	return options
}

// SetPoolID : Allow user to set PoolID
func (options *DeletePoolOptions) SetPoolID(poolID string) *DeletePoolOptions {
	options.PoolID = core.StringPtr(poolID)
	return options
}

// SetXCorrelationID : Allow user to set XCorrelationID
func (options *DeletePoolOptions) SetXCorrelationID(xCorrelationID string) *DeletePoolOptions {
	options.XCorrelationID = core.StringPtr(xCorrelationID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *DeletePoolOptions) SetHeaders(param map[string]string) *DeletePoolOptions {
	options.Headers = param
	return options
}

// GetLoadBalancerOptions : The GetLoadBalancer options.
type GetLoadBalancerOptions struct {
	// The unique identifier of a service instance.
	InstanceID *string `json:"instance_id" validate:"required"`

	// The unique identifier of a DNS zone.
	DnszoneID *string `json:"dnszone_id" validate:"required"`

	// The unique identifier of a load balancer.
	LbID *string `json:"lb_id" validate:"required"`

	// Uniquely identifying a request.
	XCorrelationID *string `json:"X-Correlation-ID,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetLoadBalancerOptions : Instantiate GetLoadBalancerOptions
func (*GlobalLoadBalancersV1) NewGetLoadBalancerOptions(instanceID string, dnszoneID string, lbID string) *GetLoadBalancerOptions {
	return &GetLoadBalancerOptions{
		InstanceID: core.StringPtr(instanceID),
		DnszoneID: core.StringPtr(dnszoneID),
		LbID: core.StringPtr(lbID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (options *GetLoadBalancerOptions) SetInstanceID(instanceID string) *GetLoadBalancerOptions {
	options.InstanceID = core.StringPtr(instanceID)
	return options
}

// SetDnszoneID : Allow user to set DnszoneID
func (options *GetLoadBalancerOptions) SetDnszoneID(dnszoneID string) *GetLoadBalancerOptions {
	options.DnszoneID = core.StringPtr(dnszoneID)
	return options
}

// SetLbID : Allow user to set LbID
func (options *GetLoadBalancerOptions) SetLbID(lbID string) *GetLoadBalancerOptions {
	options.LbID = core.StringPtr(lbID)
	return options
}

// SetXCorrelationID : Allow user to set XCorrelationID
func (options *GetLoadBalancerOptions) SetXCorrelationID(xCorrelationID string) *GetLoadBalancerOptions {
	options.XCorrelationID = core.StringPtr(xCorrelationID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetLoadBalancerOptions) SetHeaders(param map[string]string) *GetLoadBalancerOptions {
	options.Headers = param
	return options
}

// GetMonitorOptions : The GetMonitor options.
type GetMonitorOptions struct {
	// The unique identifier of a service instance.
	InstanceID *string `json:"instance_id" validate:"required"`

	// The unique identifier of a load balancer monitor.
	MonitorID *string `json:"monitor_id" validate:"required"`

	// Uniquely identifying a request.
	XCorrelationID *string `json:"X-Correlation-ID,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetMonitorOptions : Instantiate GetMonitorOptions
func (*GlobalLoadBalancersV1) NewGetMonitorOptions(instanceID string, monitorID string) *GetMonitorOptions {
	return &GetMonitorOptions{
		InstanceID: core.StringPtr(instanceID),
		MonitorID: core.StringPtr(monitorID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (options *GetMonitorOptions) SetInstanceID(instanceID string) *GetMonitorOptions {
	options.InstanceID = core.StringPtr(instanceID)
	return options
}

// SetMonitorID : Allow user to set MonitorID
func (options *GetMonitorOptions) SetMonitorID(monitorID string) *GetMonitorOptions {
	options.MonitorID = core.StringPtr(monitorID)
	return options
}

// SetXCorrelationID : Allow user to set XCorrelationID
func (options *GetMonitorOptions) SetXCorrelationID(xCorrelationID string) *GetMonitorOptions {
	options.XCorrelationID = core.StringPtr(xCorrelationID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetMonitorOptions) SetHeaders(param map[string]string) *GetMonitorOptions {
	options.Headers = param
	return options
}

// GetPoolOptions : The GetPool options.
type GetPoolOptions struct {
	// The unique identifier of a service instance.
	InstanceID *string `json:"instance_id" validate:"required"`

	// The unique identifier of a load balancer pool.
	PoolID *string `json:"pool_id" validate:"required"`

	// Uniquely identifying a request.
	XCorrelationID *string `json:"X-Correlation-ID,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetPoolOptions : Instantiate GetPoolOptions
func (*GlobalLoadBalancersV1) NewGetPoolOptions(instanceID string, poolID string) *GetPoolOptions {
	return &GetPoolOptions{
		InstanceID: core.StringPtr(instanceID),
		PoolID: core.StringPtr(poolID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (options *GetPoolOptions) SetInstanceID(instanceID string) *GetPoolOptions {
	options.InstanceID = core.StringPtr(instanceID)
	return options
}

// SetPoolID : Allow user to set PoolID
func (options *GetPoolOptions) SetPoolID(poolID string) *GetPoolOptions {
	options.PoolID = core.StringPtr(poolID)
	return options
}

// SetXCorrelationID : Allow user to set XCorrelationID
func (options *GetPoolOptions) SetXCorrelationID(xCorrelationID string) *GetPoolOptions {
	options.XCorrelationID = core.StringPtr(xCorrelationID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetPoolOptions) SetHeaders(param map[string]string) *GetPoolOptions {
	options.Headers = param
	return options
}

// ListLoadBalancersOptions : The ListLoadBalancers options.
type ListLoadBalancersOptions struct {
	// The unique identifier of a service instance.
	InstanceID *string `json:"instance_id" validate:"required"`

	// The unique identifier of a DNS zone.
	DnszoneID *string `json:"dnszone_id" validate:"required"`

	// Uniquely identifying a request.
	XCorrelationID *string `json:"X-Correlation-ID,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListLoadBalancersOptions : Instantiate ListLoadBalancersOptions
func (*GlobalLoadBalancersV1) NewListLoadBalancersOptions(instanceID string, dnszoneID string) *ListLoadBalancersOptions {
	return &ListLoadBalancersOptions{
		InstanceID: core.StringPtr(instanceID),
		DnszoneID: core.StringPtr(dnszoneID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (options *ListLoadBalancersOptions) SetInstanceID(instanceID string) *ListLoadBalancersOptions {
	options.InstanceID = core.StringPtr(instanceID)
	return options
}

// SetDnszoneID : Allow user to set DnszoneID
func (options *ListLoadBalancersOptions) SetDnszoneID(dnszoneID string) *ListLoadBalancersOptions {
	options.DnszoneID = core.StringPtr(dnszoneID)
	return options
}

// SetXCorrelationID : Allow user to set XCorrelationID
func (options *ListLoadBalancersOptions) SetXCorrelationID(xCorrelationID string) *ListLoadBalancersOptions {
	options.XCorrelationID = core.StringPtr(xCorrelationID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ListLoadBalancersOptions) SetHeaders(param map[string]string) *ListLoadBalancersOptions {
	options.Headers = param
	return options
}

// ListMonitorsOptions : The ListMonitors options.
type ListMonitorsOptions struct {
	// The unique identifier of a service instance.
	InstanceID *string `json:"instance_id" validate:"required"`

	// Uniquely identifying a request.
	XCorrelationID *string `json:"X-Correlation-ID,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListMonitorsOptions : Instantiate ListMonitorsOptions
func (*GlobalLoadBalancersV1) NewListMonitorsOptions(instanceID string) *ListMonitorsOptions {
	return &ListMonitorsOptions{
		InstanceID: core.StringPtr(instanceID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (options *ListMonitorsOptions) SetInstanceID(instanceID string) *ListMonitorsOptions {
	options.InstanceID = core.StringPtr(instanceID)
	return options
}

// SetXCorrelationID : Allow user to set XCorrelationID
func (options *ListMonitorsOptions) SetXCorrelationID(xCorrelationID string) *ListMonitorsOptions {
	options.XCorrelationID = core.StringPtr(xCorrelationID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ListMonitorsOptions) SetHeaders(param map[string]string) *ListMonitorsOptions {
	options.Headers = param
	return options
}

// ListPoolsOptions : The ListPools options.
type ListPoolsOptions struct {
	// The unique identifier of a service instance.
	InstanceID *string `json:"instance_id" validate:"required"`

	// Uniquely identifying a request.
	XCorrelationID *string `json:"X-Correlation-ID,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListPoolsOptions : Instantiate ListPoolsOptions
func (*GlobalLoadBalancersV1) NewListPoolsOptions(instanceID string) *ListPoolsOptions {
	return &ListPoolsOptions{
		InstanceID: core.StringPtr(instanceID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (options *ListPoolsOptions) SetInstanceID(instanceID string) *ListPoolsOptions {
	options.InstanceID = core.StringPtr(instanceID)
	return options
}

// SetXCorrelationID : Allow user to set XCorrelationID
func (options *ListPoolsOptions) SetXCorrelationID(xCorrelationID string) *ListPoolsOptions {
	options.XCorrelationID = core.StringPtr(xCorrelationID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ListPoolsOptions) SetHeaders(param map[string]string) *ListPoolsOptions {
	options.Headers = param
	return options
}

// LoadBalancerAzPoolsItem : LoadBalancerAzPoolsItem struct
type LoadBalancerAzPoolsItem struct {
	// Availability zone.
	AvailabilityZone *string `json:"availability_zone,omitempty"`

	// List of load balancer pools.
	Pools []string `json:"pools,omitempty"`
}


// UnmarshalLoadBalancerAzPoolsItem unmarshals an instance of LoadBalancerAzPoolsItem from the specified map of raw messages.
func UnmarshalLoadBalancerAzPoolsItem(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(LoadBalancerAzPoolsItem)
	err = core.UnmarshalPrimitive(m, "availability_zone", &obj.AvailabilityZone)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "pools", &obj.Pools)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// UpdateLoadBalancerOptions : The UpdateLoadBalancer options.
type UpdateLoadBalancerOptions struct {
	// The unique identifier of a service instance.
	InstanceID *string `json:"instance_id" validate:"required"`

	// The unique identifier of a DNS zone.
	DnszoneID *string `json:"dnszone_id" validate:"required"`

	// The unique identifier of a load balancer.
	LbID *string `json:"lb_id" validate:"required"`

	// Name of the load balancer.
	Name *string `json:"name,omitempty"`

	// Descriptive text of the load balancer.
	Description *string `json:"description,omitempty"`

	// Whether the load balancer is enabled.
	Enabled *bool `json:"enabled,omitempty"`

	// Time to live in second.
	TTL *int64 `json:"ttl,omitempty"`

	// The pool ID to use when all other pools are detected as unhealthy.
	FallbackPool *string `json:"fallback_pool,omitempty"`

	// A list of pool IDs ordered by their failover priority. Pools defined here are used by default, or when region_pools
	// are not configured for a given region.
	DefaultPools []string `json:"default_pools,omitempty"`

	// Map availability zones to pool ID's.
	AzPools []LoadBalancerAzPoolsItem `json:"az_pools,omitempty"`

	// Uniquely identifying a request.
	XCorrelationID *string `json:"X-Correlation-ID,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdateLoadBalancerOptions : Instantiate UpdateLoadBalancerOptions
func (*GlobalLoadBalancersV1) NewUpdateLoadBalancerOptions(instanceID string, dnszoneID string, lbID string) *UpdateLoadBalancerOptions {
	return &UpdateLoadBalancerOptions{
		InstanceID: core.StringPtr(instanceID),
		DnszoneID: core.StringPtr(dnszoneID),
		LbID: core.StringPtr(lbID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (options *UpdateLoadBalancerOptions) SetInstanceID(instanceID string) *UpdateLoadBalancerOptions {
	options.InstanceID = core.StringPtr(instanceID)
	return options
}

// SetDnszoneID : Allow user to set DnszoneID
func (options *UpdateLoadBalancerOptions) SetDnszoneID(dnszoneID string) *UpdateLoadBalancerOptions {
	options.DnszoneID = core.StringPtr(dnszoneID)
	return options
}

// SetLbID : Allow user to set LbID
func (options *UpdateLoadBalancerOptions) SetLbID(lbID string) *UpdateLoadBalancerOptions {
	options.LbID = core.StringPtr(lbID)
	return options
}

// SetName : Allow user to set Name
func (options *UpdateLoadBalancerOptions) SetName(name string) *UpdateLoadBalancerOptions {
	options.Name = core.StringPtr(name)
	return options
}

// SetDescription : Allow user to set Description
func (options *UpdateLoadBalancerOptions) SetDescription(description string) *UpdateLoadBalancerOptions {
	options.Description = core.StringPtr(description)
	return options
}

// SetEnabled : Allow user to set Enabled
func (options *UpdateLoadBalancerOptions) SetEnabled(enabled bool) *UpdateLoadBalancerOptions {
	options.Enabled = core.BoolPtr(enabled)
	return options
}

// SetTTL : Allow user to set TTL
func (options *UpdateLoadBalancerOptions) SetTTL(ttl int64) *UpdateLoadBalancerOptions {
	options.TTL = core.Int64Ptr(ttl)
	return options
}

// SetFallbackPool : Allow user to set FallbackPool
func (options *UpdateLoadBalancerOptions) SetFallbackPool(fallbackPool string) *UpdateLoadBalancerOptions {
	options.FallbackPool = core.StringPtr(fallbackPool)
	return options
}

// SetDefaultPools : Allow user to set DefaultPools
func (options *UpdateLoadBalancerOptions) SetDefaultPools(defaultPools []string) *UpdateLoadBalancerOptions {
	options.DefaultPools = defaultPools
	return options
}

// SetAzPools : Allow user to set AzPools
func (options *UpdateLoadBalancerOptions) SetAzPools(azPools []LoadBalancerAzPoolsItem) *UpdateLoadBalancerOptions {
	options.AzPools = azPools
	return options
}

// SetXCorrelationID : Allow user to set XCorrelationID
func (options *UpdateLoadBalancerOptions) SetXCorrelationID(xCorrelationID string) *UpdateLoadBalancerOptions {
	options.XCorrelationID = core.StringPtr(xCorrelationID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateLoadBalancerOptions) SetHeaders(param map[string]string) *UpdateLoadBalancerOptions {
	options.Headers = param
	return options
}

// UpdateMonitorOptions : The UpdateMonitor options.
type UpdateMonitorOptions struct {
	// The unique identifier of a service instance.
	InstanceID *string `json:"instance_id" validate:"required"`

	// The unique identifier of a load balancer monitor.
	MonitorID *string `json:"monitor_id" validate:"required"`

	// The name of the load balancer monitor.
	Name *string `json:"name,omitempty"`

	// Descriptive text of the load balancer monitor.
	Description *string `json:"description,omitempty"`

	// The protocol to use for the health check. Currently supported protocols are 'HTTP','HTTPS' and 'TCP'.
	Type *string `json:"type,omitempty"`

	// Port number to connect to for the health check. Required for TCP checks. HTTP and HTTPS checks should only define
	// the port when using a non-standard port (HTTP: default 80, HTTPS: default 443).
	Port *int64 `json:"port,omitempty"`

	// The interval between each health check. Shorter intervals may improve failover time, but will increase load on the
	// origins as we check from multiple locations.
	Interval *int64 `json:"interval,omitempty"`

	// The number of retries to attempt in case of a timeout before marking the origin as unhealthy. Retries are attempted
	// immediately.
	Retries *int64 `json:"retries,omitempty"`

	// The timeout (in seconds) before marking the health check as failed.
	Timeout *int64 `json:"timeout,omitempty"`

	// The method to use for the health check applicable to HTTP/HTTPS based checks, the default value is 'GET'.
	Method *string `json:"method,omitempty"`

	// The endpoint path to health check against. This parameter is only valid for HTTP and HTTPS monitors.
	Path *string `json:"path,omitempty"`

	// The HTTP request headers to send in the health check. It is recommended you set a Host header by default. The
	// User-Agent header cannot be overridden. This parameter is only valid for HTTP and HTTPS monitors.
	HeadersVar []HealthcheckHeader `json:"headers,omitempty"`

	// Do not validate the certificate when monitor use HTTPS. This parameter is currently only valid for HTTP and HTTPS
	// monitors.
	AllowInsecure *bool `json:"allow_insecure,omitempty"`

	// The expected HTTP response code or code range of the health check. This parameter is only valid for HTTP and HTTPS
	// monitors.
	ExpectedCodes *string `json:"expected_codes,omitempty"`

	// A case-insensitive sub-string to look for in the response body. If this string is not found, the origin will be
	// marked as unhealthy. This parameter is only valid for HTTP and HTTPS monitors.
	ExpectedBody *string `json:"expected_body,omitempty"`

	// Uniquely identifying a request.
	XCorrelationID *string `json:"X-Correlation-ID,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the UpdateMonitorOptions.Type property.
// The protocol to use for the health check. Currently supported protocols are 'HTTP','HTTPS' and 'TCP'.
const (
	UpdateMonitorOptions_Type_Http = "HTTP"
	UpdateMonitorOptions_Type_Https = "HTTPS"
	UpdateMonitorOptions_Type_Tcp = "TCP"
)

// Constants associated with the UpdateMonitorOptions.Method property.
// The method to use for the health check applicable to HTTP/HTTPS based checks, the default value is 'GET'.
const (
	UpdateMonitorOptions_Method_Get = "GET"
	UpdateMonitorOptions_Method_Head = "HEAD"
)

// NewUpdateMonitorOptions : Instantiate UpdateMonitorOptions
func (*GlobalLoadBalancersV1) NewUpdateMonitorOptions(instanceID string, monitorID string) *UpdateMonitorOptions {
	return &UpdateMonitorOptions{
		InstanceID: core.StringPtr(instanceID),
		MonitorID: core.StringPtr(monitorID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (options *UpdateMonitorOptions) SetInstanceID(instanceID string) *UpdateMonitorOptions {
	options.InstanceID = core.StringPtr(instanceID)
	return options
}

// SetMonitorID : Allow user to set MonitorID
func (options *UpdateMonitorOptions) SetMonitorID(monitorID string) *UpdateMonitorOptions {
	options.MonitorID = core.StringPtr(monitorID)
	return options
}

// SetName : Allow user to set Name
func (options *UpdateMonitorOptions) SetName(name string) *UpdateMonitorOptions {
	options.Name = core.StringPtr(name)
	return options
}

// SetDescription : Allow user to set Description
func (options *UpdateMonitorOptions) SetDescription(description string) *UpdateMonitorOptions {
	options.Description = core.StringPtr(description)
	return options
}

// SetType : Allow user to set Type
func (options *UpdateMonitorOptions) SetType(typeVar string) *UpdateMonitorOptions {
	options.Type = core.StringPtr(typeVar)
	return options
}

// SetPort : Allow user to set Port
func (options *UpdateMonitorOptions) SetPort(port int64) *UpdateMonitorOptions {
	options.Port = core.Int64Ptr(port)
	return options
}

// SetInterval : Allow user to set Interval
func (options *UpdateMonitorOptions) SetInterval(interval int64) *UpdateMonitorOptions {
	options.Interval = core.Int64Ptr(interval)
	return options
}

// SetRetries : Allow user to set Retries
func (options *UpdateMonitorOptions) SetRetries(retries int64) *UpdateMonitorOptions {
	options.Retries = core.Int64Ptr(retries)
	return options
}

// SetTimeout : Allow user to set Timeout
func (options *UpdateMonitorOptions) SetTimeout(timeout int64) *UpdateMonitorOptions {
	options.Timeout = core.Int64Ptr(timeout)
	return options
}

// SetMethod : Allow user to set Method
func (options *UpdateMonitorOptions) SetMethod(method string) *UpdateMonitorOptions {
	options.Method = core.StringPtr(method)
	return options
}

// SetPath : Allow user to set Path
func (options *UpdateMonitorOptions) SetPath(path string) *UpdateMonitorOptions {
	options.Path = core.StringPtr(path)
	return options
}

// SetHeadersVar : Allow user to set HeadersVar
func (options *UpdateMonitorOptions) SetHeadersVar(headersVar []HealthcheckHeader) *UpdateMonitorOptions {
	options.HeadersVar = headersVar
	return options
}

// SetAllowInsecure : Allow user to set AllowInsecure
func (options *UpdateMonitorOptions) SetAllowInsecure(allowInsecure bool) *UpdateMonitorOptions {
	options.AllowInsecure = core.BoolPtr(allowInsecure)
	return options
}

// SetExpectedCodes : Allow user to set ExpectedCodes
func (options *UpdateMonitorOptions) SetExpectedCodes(expectedCodes string) *UpdateMonitorOptions {
	options.ExpectedCodes = core.StringPtr(expectedCodes)
	return options
}

// SetExpectedBody : Allow user to set ExpectedBody
func (options *UpdateMonitorOptions) SetExpectedBody(expectedBody string) *UpdateMonitorOptions {
	options.ExpectedBody = core.StringPtr(expectedBody)
	return options
}

// SetXCorrelationID : Allow user to set XCorrelationID
func (options *UpdateMonitorOptions) SetXCorrelationID(xCorrelationID string) *UpdateMonitorOptions {
	options.XCorrelationID = core.StringPtr(xCorrelationID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateMonitorOptions) SetHeaders(param map[string]string) *UpdateMonitorOptions {
	options.Headers = param
	return options
}

// UpdatePoolOptions : The UpdatePool options.
type UpdatePoolOptions struct {
	// The unique identifier of a service instance.
	InstanceID *string `json:"instance_id" validate:"required"`

	// The unique identifier of a load balancer pool.
	PoolID *string `json:"pool_id" validate:"required"`

	// Name of the load balancer pool.
	Name *string `json:"name,omitempty"`

	// Descriptive text of the load balancer pool.
	Description *string `json:"description,omitempty"`

	// Whether the load balancer pool is enabled.
	Enabled *bool `json:"enabled,omitempty"`

	// The minimum number of origins that must be healthy for this pool to serve traffic. If the number of healthy origins
	// falls below this number, the pool will be marked unhealthy and we will failover to the next available pool.
	HealthyOriginsThreshold *int64 `json:"healthy_origins_threshold,omitempty"`

	// The list of origins within this pool. Traffic directed at this pool is balanced across all currently healthy
	// origins, provided the pool itself is healthy.
	Origins []OriginInput `json:"origins,omitempty"`

	// The ID of the load balancer monitor to be associated to this pool.
	Monitor *string `json:"monitor,omitempty"`

	// The notification channel.
	NotificationChannel *string `json:"notification_channel,omitempty"`

	// Health check region of VSIs.
	HealthcheckRegion *string `json:"healthcheck_region,omitempty"`

	// Health check subnet IDs of VSIs.
	HealthcheckSubnets []string `json:"healthcheck_subnets,omitempty"`

	// Uniquely identifying a request.
	XCorrelationID *string `json:"X-Correlation-ID,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the UpdatePoolOptions.HealthcheckRegion property.
// Health check region of VSIs.
const (
	UpdatePoolOptions_HealthcheckRegion_AuSyd = "au-syd"
	UpdatePoolOptions_HealthcheckRegion_EuDu = "eu-du"
	UpdatePoolOptions_HealthcheckRegion_EuGb = "eu-gb"
	UpdatePoolOptions_HealthcheckRegion_JpTok = "jp-tok"
	UpdatePoolOptions_HealthcheckRegion_UsEast = "us-east"
	UpdatePoolOptions_HealthcheckRegion_UsSouth = "us-south"
)

// NewUpdatePoolOptions : Instantiate UpdatePoolOptions
func (*GlobalLoadBalancersV1) NewUpdatePoolOptions(instanceID string, poolID string) *UpdatePoolOptions {
	return &UpdatePoolOptions{
		InstanceID: core.StringPtr(instanceID),
		PoolID: core.StringPtr(poolID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (options *UpdatePoolOptions) SetInstanceID(instanceID string) *UpdatePoolOptions {
	options.InstanceID = core.StringPtr(instanceID)
	return options
}

// SetPoolID : Allow user to set PoolID
func (options *UpdatePoolOptions) SetPoolID(poolID string) *UpdatePoolOptions {
	options.PoolID = core.StringPtr(poolID)
	return options
}

// SetName : Allow user to set Name
func (options *UpdatePoolOptions) SetName(name string) *UpdatePoolOptions {
	options.Name = core.StringPtr(name)
	return options
}

// SetDescription : Allow user to set Description
func (options *UpdatePoolOptions) SetDescription(description string) *UpdatePoolOptions {
	options.Description = core.StringPtr(description)
	return options
}

// SetEnabled : Allow user to set Enabled
func (options *UpdatePoolOptions) SetEnabled(enabled bool) *UpdatePoolOptions {
	options.Enabled = core.BoolPtr(enabled)
	return options
}

// SetHealthyOriginsThreshold : Allow user to set HealthyOriginsThreshold
func (options *UpdatePoolOptions) SetHealthyOriginsThreshold(healthyOriginsThreshold int64) *UpdatePoolOptions {
	options.HealthyOriginsThreshold = core.Int64Ptr(healthyOriginsThreshold)
	return options
}

// SetOrigins : Allow user to set Origins
func (options *UpdatePoolOptions) SetOrigins(origins []OriginInput) *UpdatePoolOptions {
	options.Origins = origins
	return options
}

// SetMonitor : Allow user to set Monitor
func (options *UpdatePoolOptions) SetMonitor(monitor string) *UpdatePoolOptions {
	options.Monitor = core.StringPtr(monitor)
	return options
}

// SetNotificationChannel : Allow user to set NotificationChannel
func (options *UpdatePoolOptions) SetNotificationChannel(notificationChannel string) *UpdatePoolOptions {
	options.NotificationChannel = core.StringPtr(notificationChannel)
	return options
}

// SetHealthcheckRegion : Allow user to set HealthcheckRegion
func (options *UpdatePoolOptions) SetHealthcheckRegion(healthcheckRegion string) *UpdatePoolOptions {
	options.HealthcheckRegion = core.StringPtr(healthcheckRegion)
	return options
}

// SetHealthcheckSubnets : Allow user to set HealthcheckSubnets
func (options *UpdatePoolOptions) SetHealthcheckSubnets(healthcheckSubnets []string) *UpdatePoolOptions {
	options.HealthcheckSubnets = healthcheckSubnets
	return options
}

// SetXCorrelationID : Allow user to set XCorrelationID
func (options *UpdatePoolOptions) SetXCorrelationID(xCorrelationID string) *UpdatePoolOptions {
	options.XCorrelationID = core.StringPtr(xCorrelationID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *UpdatePoolOptions) SetHeaders(param map[string]string) *UpdatePoolOptions {
	options.Headers = param
	return options
}

// FirstHref : href.
type FirstHref struct {
	// href.
	Href *string `json:"href,omitempty"`
}


// UnmarshalFirstHref unmarshals an instance of FirstHref from the specified map of raw messages.
func UnmarshalFirstHref(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(FirstHref)
	err = core.UnmarshalPrimitive(m, "href", &obj.Href)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// HealthcheckHeader : The HTTP header of health check request.
type HealthcheckHeader struct {
	// The name of HTTP request header.
	Name *string `json:"name" validate:"required"`

	// The value of HTTP request header.
	Value []string `json:"value" validate:"required"`
}


// NewHealthcheckHeader : Instantiate HealthcheckHeader (Generic Model Constructor)
func (*GlobalLoadBalancersV1) NewHealthcheckHeader(name string, value []string) (model *HealthcheckHeader, err error) {
	model = &HealthcheckHeader{
		Name: core.StringPtr(name),
		Value: value,
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalHealthcheckHeader unmarshals an instance of HealthcheckHeader from the specified map of raw messages.
func UnmarshalHealthcheckHeader(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(HealthcheckHeader)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "value", &obj.Value)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ListLoadBalancers : List Global Load Balancers response.
type ListLoadBalancers struct {
	// An array of Global Load Balancers.
	LoadBalancers []LoadBalancer `json:"load_balancers" validate:"required"`

	// Page number.
	Offset *int64 `json:"offset" validate:"required"`

	// Number of Global Load Balancers per page.
	Limit *int64 `json:"limit" validate:"required"`

	// Number of Global Load Balancers.
	Count *int64 `json:"count" validate:"required"`

	// Total number of Global Load Balancers.
	TotalCount *int64 `json:"total_count" validate:"required"`

	// href.
	First *FirstHref `json:"first" validate:"required"`

	// href.
	Next *NextHref `json:"next" validate:"required"`
}


// UnmarshalListLoadBalancers unmarshals an instance of ListLoadBalancers from the specified map of raw messages.
func UnmarshalListLoadBalancers(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ListLoadBalancers)
	err = core.UnmarshalModel(m, "load_balancers", &obj.LoadBalancers, UnmarshalLoadBalancer)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "offset", &obj.Offset)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "limit", &obj.Limit)
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
	err = core.UnmarshalModel(m, "first", &obj.First, UnmarshalFirstHref)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "next", &obj.Next, UnmarshalNextHref)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ListMonitors : List load balancer monitors response.
type ListMonitors struct {
	// An array of load balancer monitors.
	Monitors []Monitor `json:"monitors" validate:"required"`

	// Page number.
	Offset *int64 `json:"offset" validate:"required"`

	// Number of load balancer monitors per page.
	Limit *int64 `json:"limit" validate:"required"`

	// Number of load balancers.
	Count *int64 `json:"count" validate:"required"`

	// Total number of load balancers.
	TotalCount *int64 `json:"total_count" validate:"required"`

	// href.
	First *FirstHref `json:"first" validate:"required"`

	// href.
	Next *NextHref `json:"next" validate:"required"`
}


// UnmarshalListMonitors unmarshals an instance of ListMonitors from the specified map of raw messages.
func UnmarshalListMonitors(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ListMonitors)
	err = core.UnmarshalModel(m, "monitors", &obj.Monitors, UnmarshalMonitor)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "offset", &obj.Offset)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "limit", &obj.Limit)
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
	err = core.UnmarshalModel(m, "first", &obj.First, UnmarshalFirstHref)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "next", &obj.Next, UnmarshalNextHref)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ListPools : List load balancer pools response.
type ListPools struct {
	// An array of load balancer pools.
	Pools []Pool `json:"pools" validate:"required"`

	// Page number.
	Offset *int64 `json:"offset" validate:"required"`

	// Number of load balancer pools per page.
	Limit *int64 `json:"limit" validate:"required"`

	// Number of load balancers.
	Count *int64 `json:"count" validate:"required"`

	// Total number of load balancers.
	TotalCount *int64 `json:"total_count" validate:"required"`

	// href.
	First *FirstHref `json:"first" validate:"required"`

	// href.
	Next *NextHref `json:"next" validate:"required"`
}


// UnmarshalListPools unmarshals an instance of ListPools from the specified map of raw messages.
func UnmarshalListPools(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ListPools)
	err = core.UnmarshalModel(m, "pools", &obj.Pools, UnmarshalPool)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "offset", &obj.Offset)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "limit", &obj.Limit)
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
	err = core.UnmarshalModel(m, "first", &obj.First, UnmarshalFirstHref)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "next", &obj.Next, UnmarshalNextHref)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// LoadBalancer : Load balancer details.
type LoadBalancer struct {
	// Identifier of the load balancer.
	ID *string `json:"id,omitempty"`

	// Name of the load balancer.
	Name *string `json:"name,omitempty"`

	// Descriptive text of the load balancer.
	Description *string `json:"description,omitempty"`

	// Whether the load balancer is enabled.
	Enabled *bool `json:"enabled,omitempty"`

	// Time to live in second.
	TTL *int64 `json:"ttl,omitempty"`

	// Healthy state of the load balancer.
	Health *string `json:"health,omitempty"`

	// The pool ID to use when all other pools are detected as unhealthy.
	FallbackPool *string `json:"fallback_pool,omitempty"`

	// A list of pool IDs ordered by their failover priority. Pools defined here are used by default, or when region_pools
	// are not configured for a given region.
	DefaultPools []string `json:"default_pools,omitempty"`

	// Map availability zones to pool ID's.
	AzPools []LoadBalancerAzPoolsItem `json:"az_pools,omitempty"`

	// the time when a load balancer is created.
	CreatedOn *string `json:"created_on,omitempty"`

	// the recent time when a load balancer is modified.
	ModifiedOn *string `json:"modified_on,omitempty"`
}

// Constants associated with the LoadBalancer.Health property.
// Healthy state of the load balancer.
const (
	LoadBalancer_Health_Critical = "CRITICAL"
	LoadBalancer_Health_Degraded = "DEGRADED"
	LoadBalancer_Health_Healthy = "HEALTHY"
)


// UnmarshalLoadBalancer unmarshals an instance of LoadBalancer from the specified map of raw messages.
func UnmarshalLoadBalancer(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(LoadBalancer)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
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
	err = core.UnmarshalPrimitive(m, "ttl", &obj.TTL)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "health", &obj.Health)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "fallback_pool", &obj.FallbackPool)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "default_pools", &obj.DefaultPools)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "az_pools", &obj.AzPools, UnmarshalLoadBalancerAzPoolsItem)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_on", &obj.CreatedOn)
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

// Monitor : Load balancer monitor details.
type Monitor struct {
	// Identifier of the load balancer monitor.
	ID *string `json:"id,omitempty"`

	// The name of the load balancer monitor.
	Name *string `json:"name,omitempty"`

	// Descriptive text of the load balancer monitor.
	Description *string `json:"description,omitempty"`

	// The protocol to use for the health check. Currently supported protocols are 'HTTP','HTTPS' and 'TCP'.
	Type *string `json:"type,omitempty"`

	// Port number to connect to for the health check. Required for TCP checks. HTTP and HTTPS checks should only define
	// the port when using a non-standard port (HTTP: default 80, HTTPS: default 443).
	Port *int64 `json:"port,omitempty"`

	// The interval between each health check. Shorter intervals may improve failover time, but will increase load on the
	// origins as we check from multiple locations.
	Interval *int64 `json:"interval,omitempty"`

	// The number of retries to attempt in case of a timeout before marking the origin as unhealthy. Retries are attempted
	// immediately.
	Retries *int64 `json:"retries,omitempty"`

	// The timeout (in seconds) before marking the health check as failed.
	Timeout *int64 `json:"timeout,omitempty"`

	// The method to use for the health check applicable to HTTP/HTTPS based checks, the default value is 'GET'.
	Method *string `json:"method,omitempty"`

	// The endpoint path to health check against. This parameter is only valid for HTTP and HTTPS monitors.
	Path *string `json:"path,omitempty"`

	// The HTTP request headers to send in the health check. It is recommended you set a Host header by default. The
	// User-Agent header cannot be overridden. This parameter is only valid for HTTP and HTTPS monitors.
	HeadersVar []HealthcheckHeader `json:"headers,omitempty"`

	// Do not validate the certificate when monitor use HTTPS. This parameter is currently only valid for HTTPS monitors.
	AllowInsecure *bool `json:"allow_insecure,omitempty"`

	// The expected HTTP response code or code range of the health check. This parameter is only valid for HTTP and HTTPS
	// monitors.
	ExpectedCodes *string `json:"expected_codes,omitempty"`

	// A case-insensitive sub-string to look for in the response body. If this string is not found, the origin will be
	// marked as unhealthy. This parameter is only valid for HTTP and HTTPS monitors.
	ExpectedBody *string `json:"expected_body,omitempty"`

	// the time when a load balancer monitor is created.
	CreatedOn *string `json:"created_on,omitempty"`

	// the recent time when a load balancer monitor is modified.
	ModifiedOn *string `json:"modified_on,omitempty"`
}

// Constants associated with the Monitor.Method property.
// The method to use for the health check applicable to HTTP/HTTPS based checks, the default value is 'GET'.
const (
	Monitor_Method_Get = "GET"
	Monitor_Method_Head = "HEAD"
)


// UnmarshalMonitor unmarshals an instance of Monitor from the specified map of raw messages.
func UnmarshalMonitor(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Monitor)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "port", &obj.Port)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "interval", &obj.Interval)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "retries", &obj.Retries)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "timeout", &obj.Timeout)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "method", &obj.Method)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "path", &obj.Path)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "headers", &obj.HeadersVar, UnmarshalHealthcheckHeader)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "allow_insecure", &obj.AllowInsecure)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "expected_codes", &obj.ExpectedCodes)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "expected_body", &obj.ExpectedBody)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_on", &obj.CreatedOn)
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

// NextHref : href.
type NextHref struct {
	// href.
	Href *string `json:"href,omitempty"`
}


// UnmarshalNextHref unmarshals an instance of NextHref from the specified map of raw messages.
func UnmarshalNextHref(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(NextHref)
	err = core.UnmarshalPrimitive(m, "href", &obj.Href)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Origin : Origin server.
type Origin struct {
	// The name of the origin server.
	Name *string `json:"name,omitempty"`

	// Description of the origin server.
	Description *string `json:"description,omitempty"`

	// The address of the origin server. It can be a hostname or an IP address.
	Address *string `json:"address,omitempty"`

	// Whether the origin server is enabled.
	Enabled *bool `json:"enabled,omitempty"`

	// The health state of the origin server.
	Health *bool `json:"health,omitempty"`

	// The failure reason of the origin server if it is unhealthy.
	HealthFailureReason *string `json:"health_failure_reason,omitempty"`
}


// UnmarshalOrigin unmarshals an instance of Origin from the specified map of raw messages.
func UnmarshalOrigin(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Origin)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "address", &obj.Address)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "enabled", &obj.Enabled)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "health", &obj.Health)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "health_failure_reason", &obj.HealthFailureReason)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// OriginInput : The request data of origin server.
type OriginInput struct {
	// The name of the origin server.
	Name *string `json:"name,omitempty"`

	// Description of the origin server.
	Description *string `json:"description,omitempty"`

	// The address of the origin server. It can be a hostname or an IP address.
	Address *string `json:"address,omitempty"`

	// Whether the origin server is enabled.
	Enabled *bool `json:"enabled,omitempty"`
}


// UnmarshalOriginInput unmarshals an instance of OriginInput from the specified map of raw messages.
func UnmarshalOriginInput(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(OriginInput)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "address", &obj.Address)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "enabled", &obj.Enabled)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Pool : Load balancer pool details.
type Pool struct {
	// Identifier of the load balancer pool.
	ID *string `json:"id,omitempty"`

	// Name of the load balancer pool.
	Name *string `json:"name,omitempty"`

	// Descriptive text of the load balancer pool.
	Description *string `json:"description,omitempty"`

	// Whether the load balancer pool is enabled.
	Enabled *bool `json:"enabled,omitempty"`

	// The minimum number of origins that must be healthy for this pool to serve traffic. If the number of healthy origins
	// falls below this number, the pool will be marked unhealthy and we will failover to the next available pool.
	HealthyOriginsThreshold *int64 `json:"healthy_origins_threshold,omitempty"`

	// The list of origins within this pool. Traffic directed at this pool is balanced across all currently healthy
	// origins, provided the pool itself is healthy.
	Origins []Origin `json:"origins,omitempty"`

	// The ID of the load balancer monitor to be associated to this pool.
	Monitor *string `json:"monitor,omitempty"`

	// The notification channel.
	NotificationChannel *string `json:"notification_channel,omitempty"`

	// Healthy state of the load balancer pool.
	Health *string `json:"health,omitempty"`

	// Health check region of VSIs.
	HealthcheckRegion *string `json:"healthcheck_region,omitempty"`

	// Health check subnet IDs of VSIs.
	HealthcheckSubnets []string `json:"healthcheck_subnets,omitempty"`

	// the time when a load balancer pool is created.
	CreatedOn *string `json:"created_on,omitempty"`

	// the recent time when a load balancer pool is modified.
	ModifiedOn *string `json:"modified_on,omitempty"`
}

// Constants associated with the Pool.Health property.
// Healthy state of the load balancer pool.
const (
	Pool_Health_Critical = "CRITICAL"
	Pool_Health_Degraded = "DEGRADED"
	Pool_Health_Healthy = "HEALTHY"
)

// Constants associated with the Pool.HealthcheckRegion property.
// Health check region of VSIs.
const (
	Pool_HealthcheckRegion_AuSyd = "au-syd"
	Pool_HealthcheckRegion_EuDu = "eu-du"
	Pool_HealthcheckRegion_EuGb = "eu-gb"
	Pool_HealthcheckRegion_JpTok = "jp-tok"
	Pool_HealthcheckRegion_UsEast = "us-east"
	Pool_HealthcheckRegion_UsSouth = "us-south"
)


// UnmarshalPool unmarshals an instance of Pool from the specified map of raw messages.
func UnmarshalPool(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Pool)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
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
	err = core.UnmarshalPrimitive(m, "healthy_origins_threshold", &obj.HealthyOriginsThreshold)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "origins", &obj.Origins, UnmarshalOrigin)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "monitor", &obj.Monitor)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "notification_channel", &obj.NotificationChannel)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "health", &obj.Health)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "healthcheck_region", &obj.HealthcheckRegion)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "healthcheck_subnets", &obj.HealthcheckSubnets)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_on", &obj.CreatedOn)
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
