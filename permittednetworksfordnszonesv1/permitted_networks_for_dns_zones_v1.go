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

// Package permittednetworksfordnszonesv1 : Operations and models for the PermittedNetworksForDnsZonesV1 service
package permittednetworksfordnszonesv1

import (
	"encoding/json"
	"fmt"
	"github.com/IBM/go-sdk-core/v4/core"
	common "github.com/IBM/networking-go-sdk/common"
	"reflect"
)

// PermittedNetworksForDnsZonesV1 : Permitted Networks for DNS Zones
//
// Version: 1.0.0
type PermittedNetworksForDnsZonesV1 struct {
	Service *core.BaseService
}

// DefaultServiceURL is the default URL to make service requests to.
const DefaultServiceURL = "https://api.dns-svcs.cloud.ibm.com/v1"

// DefaultServiceName is the default key used to find external configuration information.
const DefaultServiceName = "permitted_networks_for_dns_zones"

// PermittedNetworksForDnsZonesV1Options : Service options
type PermittedNetworksForDnsZonesV1Options struct {
	ServiceName   string
	URL           string
	Authenticator core.Authenticator
}

// NewPermittedNetworksForDnsZonesV1UsingExternalConfig : constructs an instance of PermittedNetworksForDnsZonesV1 with passed in options and external configuration.
func NewPermittedNetworksForDnsZonesV1UsingExternalConfig(options *PermittedNetworksForDnsZonesV1Options) (permittedNetworksForDnsZones *PermittedNetworksForDnsZonesV1, err error) {
	if options.ServiceName == "" {
		options.ServiceName = DefaultServiceName
	}

	if options.Authenticator == nil {
		options.Authenticator, err = core.GetAuthenticatorFromEnvironment(options.ServiceName)
		if err != nil {
			return
		}
	}

	permittedNetworksForDnsZones, err = NewPermittedNetworksForDnsZonesV1(options)
	if err != nil {
		return
	}

	err = permittedNetworksForDnsZones.Service.ConfigureService(options.ServiceName)
	if err != nil {
		return
	}

	if options.URL != "" {
		err = permittedNetworksForDnsZones.Service.SetServiceURL(options.URL)
	}
	return
}

// NewPermittedNetworksForDnsZonesV1 : constructs an instance of PermittedNetworksForDnsZonesV1 with passed in options.
func NewPermittedNetworksForDnsZonesV1(options *PermittedNetworksForDnsZonesV1Options) (service *PermittedNetworksForDnsZonesV1, err error) {
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

	service = &PermittedNetworksForDnsZonesV1{
		Service: baseService,
	}

	return
}

// SetServiceURL sets the service URL
func (permittedNetworksForDnsZones *PermittedNetworksForDnsZonesV1) SetServiceURL(url string) error {
	return permittedNetworksForDnsZones.Service.SetServiceURL(url)
}

// ListPermittedNetworks : List permitted networks
// List the permitted networks for a given DNS zone.
func (permittedNetworksForDnsZones *PermittedNetworksForDnsZonesV1) ListPermittedNetworks(listPermittedNetworksOptions *ListPermittedNetworksOptions) (result *ListPermittedNetworks, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listPermittedNetworksOptions, "listPermittedNetworksOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listPermittedNetworksOptions, "listPermittedNetworksOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"instances", "dnszones", "permitted_networks"}
	pathParameters := []string{*listPermittedNetworksOptions.InstanceID, *listPermittedNetworksOptions.DnszoneID}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(permittedNetworksForDnsZones.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range listPermittedNetworksOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("permitted_networks_for_dns_zones", "V1", "ListPermittedNetworks")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if listPermittedNetworksOptions.XCorrelationID != nil {
		builder.AddHeader("X-Correlation-ID", fmt.Sprint(*listPermittedNetworksOptions.XCorrelationID))
	}

	if listPermittedNetworksOptions.Offset != nil {
		builder.AddQuery("offset", fmt.Sprint(*listPermittedNetworksOptions.Offset))
	}
	if listPermittedNetworksOptions.Limit != nil {
		builder.AddQuery("limit", fmt.Sprint(*listPermittedNetworksOptions.Limit))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = permittedNetworksForDnsZones.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalListPermittedNetworks)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// CreatePermittedNetwork : Create a permitted network
// Create a permitted network for a given DNS zone.
func (permittedNetworksForDnsZones *PermittedNetworksForDnsZonesV1) CreatePermittedNetwork(createPermittedNetworkOptions *CreatePermittedNetworkOptions) (result *PermittedNetwork, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createPermittedNetworkOptions, "createPermittedNetworkOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createPermittedNetworkOptions, "createPermittedNetworkOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"instances", "dnszones", "permitted_networks"}
	pathParameters := []string{*createPermittedNetworkOptions.InstanceID, *createPermittedNetworkOptions.DnszoneID}

	builder := core.NewRequestBuilder(core.POST)
	_, err = builder.ConstructHTTPURL(permittedNetworksForDnsZones.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range createPermittedNetworkOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("permitted_networks_for_dns_zones", "V1", "CreatePermittedNetwork")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if createPermittedNetworkOptions.XCorrelationID != nil {
		builder.AddHeader("X-Correlation-ID", fmt.Sprint(*createPermittedNetworkOptions.XCorrelationID))
	}

	body := make(map[string]interface{})
	if createPermittedNetworkOptions.Type != nil {
		body["type"] = createPermittedNetworkOptions.Type
	}
	if createPermittedNetworkOptions.PermittedNetwork != nil {
		body["permitted_network"] = createPermittedNetworkOptions.PermittedNetwork
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
	response, err = permittedNetworksForDnsZones.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalPermittedNetwork)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// DeletePermittedNetwork : Remove a permitted network
// Remove a permitted network.
func (permittedNetworksForDnsZones *PermittedNetworksForDnsZonesV1) DeletePermittedNetwork(deletePermittedNetworkOptions *DeletePermittedNetworkOptions) (result *PermittedNetwork, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deletePermittedNetworkOptions, "deletePermittedNetworkOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deletePermittedNetworkOptions, "deletePermittedNetworkOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"instances", "dnszones", "permitted_networks"}
	pathParameters := []string{*deletePermittedNetworkOptions.InstanceID, *deletePermittedNetworkOptions.DnszoneID, *deletePermittedNetworkOptions.PermittedNetworkID}

	builder := core.NewRequestBuilder(core.DELETE)
	_, err = builder.ConstructHTTPURL(permittedNetworksForDnsZones.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range deletePermittedNetworkOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("permitted_networks_for_dns_zones", "V1", "DeletePermittedNetwork")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if deletePermittedNetworkOptions.XCorrelationID != nil {
		builder.AddHeader("X-Correlation-ID", fmt.Sprint(*deletePermittedNetworkOptions.XCorrelationID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = permittedNetworksForDnsZones.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalPermittedNetwork)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// GetPermittedNetwork : Get a permitted network
// Get details of a permitted network.
func (permittedNetworksForDnsZones *PermittedNetworksForDnsZonesV1) GetPermittedNetwork(getPermittedNetworkOptions *GetPermittedNetworkOptions) (result *PermittedNetwork, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getPermittedNetworkOptions, "getPermittedNetworkOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getPermittedNetworkOptions, "getPermittedNetworkOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"instances", "dnszones", "permitted_networks"}
	pathParameters := []string{*getPermittedNetworkOptions.InstanceID, *getPermittedNetworkOptions.DnszoneID, *getPermittedNetworkOptions.PermittedNetworkID}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(permittedNetworksForDnsZones.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range getPermittedNetworkOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("permitted_networks_for_dns_zones", "V1", "GetPermittedNetwork")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if getPermittedNetworkOptions.XCorrelationID != nil {
		builder.AddHeader("X-Correlation-ID", fmt.Sprint(*getPermittedNetworkOptions.XCorrelationID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = permittedNetworksForDnsZones.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalPermittedNetwork)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// CreatePermittedNetworkOptions : The CreatePermittedNetwork options.
type CreatePermittedNetworkOptions struct {
	// The unique identifier of a service instance.
	InstanceID *string `json:"instance_id" validate:"required"`

	// The unique identifier of a DNS zone.
	DnszoneID *string `json:"dnszone_id" validate:"required"`

	// The type of a permitted network.
	Type *string `json:"type,omitempty"`

	// Permitted network data for VPC.
	PermittedNetwork *PermittedNetworkVpc `json:"permitted_network,omitempty"`

	// Uniquely identifying a request.
	XCorrelationID *string `json:"X-Correlation-ID,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the CreatePermittedNetworkOptions.Type property.
// The type of a permitted network.
const (
	CreatePermittedNetworkOptions_Type_Vpc = "vpc"
)

// NewCreatePermittedNetworkOptions : Instantiate CreatePermittedNetworkOptions
func (*PermittedNetworksForDnsZonesV1) NewCreatePermittedNetworkOptions(instanceID string, dnszoneID string) *CreatePermittedNetworkOptions {
	return &CreatePermittedNetworkOptions{
		InstanceID: core.StringPtr(instanceID),
		DnszoneID: core.StringPtr(dnszoneID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (options *CreatePermittedNetworkOptions) SetInstanceID(instanceID string) *CreatePermittedNetworkOptions {
	options.InstanceID = core.StringPtr(instanceID)
	return options
}

// SetDnszoneID : Allow user to set DnszoneID
func (options *CreatePermittedNetworkOptions) SetDnszoneID(dnszoneID string) *CreatePermittedNetworkOptions {
	options.DnszoneID = core.StringPtr(dnszoneID)
	return options
}

// SetType : Allow user to set Type
func (options *CreatePermittedNetworkOptions) SetType(typeVar string) *CreatePermittedNetworkOptions {
	options.Type = core.StringPtr(typeVar)
	return options
}

// SetPermittedNetwork : Allow user to set PermittedNetwork
func (options *CreatePermittedNetworkOptions) SetPermittedNetwork(permittedNetwork *PermittedNetworkVpc) *CreatePermittedNetworkOptions {
	options.PermittedNetwork = permittedNetwork
	return options
}

// SetXCorrelationID : Allow user to set XCorrelationID
func (options *CreatePermittedNetworkOptions) SetXCorrelationID(xCorrelationID string) *CreatePermittedNetworkOptions {
	options.XCorrelationID = core.StringPtr(xCorrelationID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *CreatePermittedNetworkOptions) SetHeaders(param map[string]string) *CreatePermittedNetworkOptions {
	options.Headers = param
	return options
}

// DeletePermittedNetworkOptions : The DeletePermittedNetwork options.
type DeletePermittedNetworkOptions struct {
	// The unique identifier of a service instance.
	InstanceID *string `json:"instance_id" validate:"required"`

	// The unique identifier of a DNS zone.
	DnszoneID *string `json:"dnszone_id" validate:"required"`

	// The unique identifier of a permitted network.
	PermittedNetworkID *string `json:"permitted_network_id" validate:"required"`

	// Uniquely identifying a request.
	XCorrelationID *string `json:"X-Correlation-ID,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeletePermittedNetworkOptions : Instantiate DeletePermittedNetworkOptions
func (*PermittedNetworksForDnsZonesV1) NewDeletePermittedNetworkOptions(instanceID string, dnszoneID string, permittedNetworkID string) *DeletePermittedNetworkOptions {
	return &DeletePermittedNetworkOptions{
		InstanceID: core.StringPtr(instanceID),
		DnszoneID: core.StringPtr(dnszoneID),
		PermittedNetworkID: core.StringPtr(permittedNetworkID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (options *DeletePermittedNetworkOptions) SetInstanceID(instanceID string) *DeletePermittedNetworkOptions {
	options.InstanceID = core.StringPtr(instanceID)
	return options
}

// SetDnszoneID : Allow user to set DnszoneID
func (options *DeletePermittedNetworkOptions) SetDnszoneID(dnszoneID string) *DeletePermittedNetworkOptions {
	options.DnszoneID = core.StringPtr(dnszoneID)
	return options
}

// SetPermittedNetworkID : Allow user to set PermittedNetworkID
func (options *DeletePermittedNetworkOptions) SetPermittedNetworkID(permittedNetworkID string) *DeletePermittedNetworkOptions {
	options.PermittedNetworkID = core.StringPtr(permittedNetworkID)
	return options
}

// SetXCorrelationID : Allow user to set XCorrelationID
func (options *DeletePermittedNetworkOptions) SetXCorrelationID(xCorrelationID string) *DeletePermittedNetworkOptions {
	options.XCorrelationID = core.StringPtr(xCorrelationID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *DeletePermittedNetworkOptions) SetHeaders(param map[string]string) *DeletePermittedNetworkOptions {
	options.Headers = param
	return options
}

// GetPermittedNetworkOptions : The GetPermittedNetwork options.
type GetPermittedNetworkOptions struct {
	// The unique identifier of a service instance.
	InstanceID *string `json:"instance_id" validate:"required"`

	// The unique identifier of a DNS zone.
	DnszoneID *string `json:"dnszone_id" validate:"required"`

	// The unique identifier of a permitted network.
	PermittedNetworkID *string `json:"permitted_network_id" validate:"required"`

	// Uniquely identifying a request.
	XCorrelationID *string `json:"X-Correlation-ID,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetPermittedNetworkOptions : Instantiate GetPermittedNetworkOptions
func (*PermittedNetworksForDnsZonesV1) NewGetPermittedNetworkOptions(instanceID string, dnszoneID string, permittedNetworkID string) *GetPermittedNetworkOptions {
	return &GetPermittedNetworkOptions{
		InstanceID: core.StringPtr(instanceID),
		DnszoneID: core.StringPtr(dnszoneID),
		PermittedNetworkID: core.StringPtr(permittedNetworkID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (options *GetPermittedNetworkOptions) SetInstanceID(instanceID string) *GetPermittedNetworkOptions {
	options.InstanceID = core.StringPtr(instanceID)
	return options
}

// SetDnszoneID : Allow user to set DnszoneID
func (options *GetPermittedNetworkOptions) SetDnszoneID(dnszoneID string) *GetPermittedNetworkOptions {
	options.DnszoneID = core.StringPtr(dnszoneID)
	return options
}

// SetPermittedNetworkID : Allow user to set PermittedNetworkID
func (options *GetPermittedNetworkOptions) SetPermittedNetworkID(permittedNetworkID string) *GetPermittedNetworkOptions {
	options.PermittedNetworkID = core.StringPtr(permittedNetworkID)
	return options
}

// SetXCorrelationID : Allow user to set XCorrelationID
func (options *GetPermittedNetworkOptions) SetXCorrelationID(xCorrelationID string) *GetPermittedNetworkOptions {
	options.XCorrelationID = core.StringPtr(xCorrelationID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetPermittedNetworkOptions) SetHeaders(param map[string]string) *GetPermittedNetworkOptions {
	options.Headers = param
	return options
}

// ListPermittedNetworksOptions : The ListPermittedNetworks options.
type ListPermittedNetworksOptions struct {
	// The unique identifier of a service instance.
	InstanceID *string `json:"instance_id" validate:"required"`

	// The unique identifier of a DNS zone.
	DnszoneID *string `json:"dnszone_id" validate:"required"`

	// Uniquely identifying a request.
	XCorrelationID *string `json:"X-Correlation-ID,omitempty"`

	// Specify how many permitted networks to skip over, the default value is 0.
	Offset *string `json:"offset,omitempty"`

	// Specify how many permitted networks are returned, the default value is 10.
	Limit *string `json:"limit,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListPermittedNetworksOptions : Instantiate ListPermittedNetworksOptions
func (*PermittedNetworksForDnsZonesV1) NewListPermittedNetworksOptions(instanceID string, dnszoneID string) *ListPermittedNetworksOptions {
	return &ListPermittedNetworksOptions{
		InstanceID: core.StringPtr(instanceID),
		DnszoneID: core.StringPtr(dnszoneID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (options *ListPermittedNetworksOptions) SetInstanceID(instanceID string) *ListPermittedNetworksOptions {
	options.InstanceID = core.StringPtr(instanceID)
	return options
}

// SetDnszoneID : Allow user to set DnszoneID
func (options *ListPermittedNetworksOptions) SetDnszoneID(dnszoneID string) *ListPermittedNetworksOptions {
	options.DnszoneID = core.StringPtr(dnszoneID)
	return options
}

// SetXCorrelationID : Allow user to set XCorrelationID
func (options *ListPermittedNetworksOptions) SetXCorrelationID(xCorrelationID string) *ListPermittedNetworksOptions {
	options.XCorrelationID = core.StringPtr(xCorrelationID)
	return options
}

// SetOffset : Allow user to set Offset
func (options *ListPermittedNetworksOptions) SetOffset(offset string) *ListPermittedNetworksOptions {
	options.Offset = core.StringPtr(offset)
	return options
}

// SetLimit : Allow user to set Limit
func (options *ListPermittedNetworksOptions) SetLimit(limit string) *ListPermittedNetworksOptions {
	options.Limit = core.StringPtr(limit)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ListPermittedNetworksOptions) SetHeaders(param map[string]string) *ListPermittedNetworksOptions {
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

// ListPermittedNetworks : List permitted networks response.
type ListPermittedNetworks struct {
	// An array of permitted networks.
	PermittedNetworks []PermittedNetwork `json:"permitted_networks" validate:"required"`

	// Specify how many permitted networks to skip over, the default value is 0.
	Offset *int64 `json:"offset" validate:"required"`

	// Specify how many permitted networks are returned, the default value is 10.
	Limit *int64 `json:"limit" validate:"required"`

	// Total number of permitted networks.
	TotalCount *int64 `json:"total_count" validate:"required"`

	// href.
	First *FirstHref `json:"first" validate:"required"`

	// href.
	Next *NextHref `json:"next,omitempty"`
}


// UnmarshalListPermittedNetworks unmarshals an instance of ListPermittedNetworks from the specified map of raw messages.
func UnmarshalListPermittedNetworks(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ListPermittedNetworks)
	err = core.UnmarshalModel(m, "permitted_networks", &obj.PermittedNetworks, UnmarshalPermittedNetwork)
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

// PermittedNetwork : Permitted network details.
type PermittedNetwork struct {
	// Unique identifier of a permitted network.
	ID *string `json:"id,omitempty"`

	// The time when a permitted network is created.
	CreatedOn *string `json:"created_on,omitempty"`

	// The recent time when a permitted network is modified.
	ModifiedOn *string `json:"modified_on,omitempty"`

	// Permitted network data for VPC.
	PermittedNetwork *PermittedNetworkVpc `json:"permitted_network,omitempty"`

	// The type of a permitted network.
	Type *string `json:"type,omitempty"`

	// The state of a permitted network.
	State *string `json:"state,omitempty"`
}

// Constants associated with the PermittedNetwork.Type property.
// The type of a permitted network.
const (
	PermittedNetwork_Type_Vpc = "vpc"
)

// Constants associated with the PermittedNetwork.State property.
// The state of a permitted network.
const (
	PermittedNetwork_State_Active = "ACTIVE"
	PermittedNetwork_State_RemovalInProgress = "REMOVAL_IN_PROGRESS"
)


// UnmarshalPermittedNetwork unmarshals an instance of PermittedNetwork from the specified map of raw messages.
func UnmarshalPermittedNetwork(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PermittedNetwork)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
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
	err = core.UnmarshalModel(m, "permitted_network", &obj.PermittedNetwork, UnmarshalPermittedNetworkVpc)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "state", &obj.State)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// PermittedNetworkVpc : Permitted network data for VPC.
type PermittedNetworkVpc struct {
	// CRN string uniquely identifies a VPC.
	VpcCrn *string `json:"vpc_crn" validate:"required"`
}


// NewPermittedNetworkVpc : Instantiate PermittedNetworkVpc (Generic Model Constructor)
func (*PermittedNetworksForDnsZonesV1) NewPermittedNetworkVpc(vpcCrn string) (model *PermittedNetworkVpc, err error) {
	model = &PermittedNetworkVpc{
		VpcCrn: core.StringPtr(vpcCrn),
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalPermittedNetworkVpc unmarshals an instance of PermittedNetworkVpc from the specified map of raw messages.
func UnmarshalPermittedNetworkVpc(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PermittedNetworkVpc)
	err = core.UnmarshalPrimitive(m, "vpc_crn", &obj.VpcCrn)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}
