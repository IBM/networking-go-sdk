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

package directlinkv1_test

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/networking-go-sdk/directlinkv1"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe(`DirectLinkV1`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		version := "testString"
		It(`Instantiate service client`, func() {
			directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
				Version:       core.StringPtr(version),
				ServiceName:   "dl_services",
			})
			Expect(directLinkService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
				URL:     "{BAD_URL_STRING",
				Version: core.StringPtr(version),
			})
			Expect(directLinkService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
				URL:     "https://directlinkv1/api",
				Version: core.StringPtr(version),
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(directLinkService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Validation Error`, func() {
			directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{})
			Expect(directLinkService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		version := "testString"
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"DIRECT_LINK_URL":       "https://directlinkv1/api",
				"DIRECT_LINK_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1UsingExternalConfig(&directlinkv1.DirectLinkV1Options{
					Version: core.StringPtr(version),
				})
				Expect(directLinkService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := directLinkService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != directLinkService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(directLinkService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(directLinkService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1UsingExternalConfig(&directlinkv1.DirectLinkV1Options{
					URL:     "https://testService/api",
					Version: core.StringPtr(version),
				})
				Expect(directLinkService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := directLinkService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != directLinkService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(directLinkService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(directLinkService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1UsingExternalConfig(&directlinkv1.DirectLinkV1Options{
					Version: core.StringPtr(version),
				})
				err := directLinkService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := directLinkService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != directLinkService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(directLinkService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(directLinkService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"DIRECT_LINK_URL":       "https://directlinkv1/api",
				"DIRECT_LINK_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			directLinkService, serviceErr := directlinkv1.NewDirectLinkV1UsingExternalConfig(&directlinkv1.DirectLinkV1Options{
				Version: core.StringPtr(version),
			})

			It(`Instantiate service client with error`, func() {
				Expect(directLinkService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"DIRECT_LINK_AUTH_TYPE": "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			directLinkService, serviceErr := directlinkv1.NewDirectLinkV1UsingExternalConfig(&directlinkv1.DirectLinkV1Options{
				URL:     "{BAD_URL_STRING",
				Version: core.StringPtr(version),
			})

			It(`Instantiate service client with error`, func() {
				Expect(directLinkService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = directlinkv1.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`ListGateways(listGatewaysOptions *ListGatewaysOptions) - Operation response error`, func() {
		version := "testString"
		listGatewaysPath := "/gateways"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listGatewaysPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListGateways with error: Operation response processing error`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Construct an instance of the ListGatewaysOptions model
				listGatewaysOptionsModel := new(directlinkv1.ListGatewaysOptions)
				listGatewaysOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := directLinkService.ListGateways(listGatewaysOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				directLinkService.EnableRetries(0, 0)
				result, response, operationErr = directLinkService.ListGateways(listGatewaysOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListGateways(listGatewaysOptions *ListGatewaysOptions)`, func() {
		version := "testString"
		listGatewaysPath := "/gateways"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listGatewaysPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"gateways": [{"as_prepends": [{"created_at": "2019-01-01T12:00:00.000Z", "id": "ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4", "length": 4, "policy": "import", "prefix": "172.17.0.0/16", "specific_prefixes": ["192.168.3.0/24"], "updated_at": "2019-01-01T12:00:00.000Z"}], "authentication_key": {"crn": "crn:v1:bluemix:public:kms:us-south:a/4111d05f36894e3cb9b46a43556d9000:abc111b8-37aa-4034-9def-f2607c87aaaa:key:bbb222bc-430a-4de9-9aad-84e5bb022222"}, "bfd_config": {"bfd_status": "up", "bfd_status_updated_at": "2020-08-20T06:58:41.909Z", "interval": 2000, "multiplier": 10}, "bgp_asn": 64999, "bgp_base_cidr": "BgpBaseCidr", "bgp_cer_cidr": "10.254.30.78/30", "bgp_ibm_asn": 13884, "bgp_ibm_cidr": "10.254.30.77/30", "bgp_status": "active", "bgp_status_updated_at": "2020-08-20T06:58:41.909Z", "carrier_name": "myCarrierName", "change_request": {"type": "create_gateway"}, "completion_notice_reject_reason": "The completion notice file was blank", "connection_mode": "transit", "created_at": "2019-01-01T12:00:00.000Z", "crn": "crn:v1:bluemix:public:directlink:dal03:a/4111d05f36894e3cb9b46a43556d9000::dedicated:ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4", "cross_account": false, "cross_connect_router": "xcr01.dal03", "customer_name": "newCustomerName", "default_export_route_filter": "permit", "default_import_route_filter": "permit", "global": true, "id": "ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4", "link_status": "up", "link_status_updated_at": "2020-08-20T06:58:41.909Z", "location_display_name": "Dallas 03", "location_name": "dal03", "macsec": {"active": true, "security_policy": "must_secure", "status": "secured", "status_reasons": [{"code": "macsec_cak_failed", "message": "The authentication_key failed configuration.", "more_info": "https://cloud.ibm.com/docs/dl/TODO_ADD_DOCS_LINK"}]}, "macsec_capability": "non_macsec", "metered": false, "name": "myGateway", "operational_status": "awaiting_completion_notice", "operational_status_reasons": [{"code": "authentication_key_failed", "message": "The authentication_key failed configuration.", "more_info": "https://cloud.ibm.com/docs/dl/TODO_ADD_DOCS_LINK"}], "patch_panel_completion_notice": "patch panel configuration details", "port": {"id": "54321b1a-fee4-41c7-9e11-9cd99e000aaa"}, "provider_api_managed": false, "resource_group": {"id": "56969d6043e9465c883cb9f7363e78e8"}, "speed_mbps": 1000, "type": "dedicated", "vlan": 10}]}`)
				}))
			})
			It(`Invoke ListGateways successfully with retries`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())
				directLinkService.EnableRetries(0, 0)

				// Construct an instance of the ListGatewaysOptions model
				listGatewaysOptionsModel := new(directlinkv1.ListGatewaysOptions)
				listGatewaysOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := directLinkService.ListGatewaysWithContext(ctx, listGatewaysOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				directLinkService.DisableRetries()
				result, response, operationErr := directLinkService.ListGateways(listGatewaysOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = directLinkService.ListGatewaysWithContext(ctx, listGatewaysOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listGatewaysPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"gateways": [{"as_prepends": [{"created_at": "2019-01-01T12:00:00.000Z", "id": "ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4", "length": 4, "policy": "import", "prefix": "172.17.0.0/16", "specific_prefixes": ["192.168.3.0/24"], "updated_at": "2019-01-01T12:00:00.000Z"}], "authentication_key": {"crn": "crn:v1:bluemix:public:kms:us-south:a/4111d05f36894e3cb9b46a43556d9000:abc111b8-37aa-4034-9def-f2607c87aaaa:key:bbb222bc-430a-4de9-9aad-84e5bb022222"}, "bfd_config": {"bfd_status": "up", "bfd_status_updated_at": "2020-08-20T06:58:41.909Z", "interval": 2000, "multiplier": 10}, "bgp_asn": 64999, "bgp_base_cidr": "BgpBaseCidr", "bgp_cer_cidr": "10.254.30.78/30", "bgp_ibm_asn": 13884, "bgp_ibm_cidr": "10.254.30.77/30", "bgp_status": "active", "bgp_status_updated_at": "2020-08-20T06:58:41.909Z", "carrier_name": "myCarrierName", "change_request": {"type": "create_gateway"}, "completion_notice_reject_reason": "The completion notice file was blank", "connection_mode": "transit", "created_at": "2019-01-01T12:00:00.000Z", "crn": "crn:v1:bluemix:public:directlink:dal03:a/4111d05f36894e3cb9b46a43556d9000::dedicated:ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4", "cross_account": false, "cross_connect_router": "xcr01.dal03", "customer_name": "newCustomerName", "default_export_route_filter": "permit", "default_import_route_filter": "permit", "global": true, "id": "ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4", "link_status": "up", "link_status_updated_at": "2020-08-20T06:58:41.909Z", "location_display_name": "Dallas 03", "location_name": "dal03", "macsec": {"active": true, "security_policy": "must_secure", "status": "secured", "status_reasons": [{"code": "macsec_cak_failed", "message": "The authentication_key failed configuration.", "more_info": "https://cloud.ibm.com/docs/dl/TODO_ADD_DOCS_LINK"}]}, "macsec_capability": "non_macsec", "metered": false, "name": "myGateway", "operational_status": "awaiting_completion_notice", "operational_status_reasons": [{"code": "authentication_key_failed", "message": "The authentication_key failed configuration.", "more_info": "https://cloud.ibm.com/docs/dl/TODO_ADD_DOCS_LINK"}], "patch_panel_completion_notice": "patch panel configuration details", "port": {"id": "54321b1a-fee4-41c7-9e11-9cd99e000aaa"}, "provider_api_managed": false, "resource_group": {"id": "56969d6043e9465c883cb9f7363e78e8"}, "speed_mbps": 1000, "type": "dedicated", "vlan": 10}]}`)
				}))
			})
			It(`Invoke ListGateways successfully`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := directLinkService.ListGateways(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListGatewaysOptions model
				listGatewaysOptionsModel := new(directlinkv1.ListGatewaysOptions)
				listGatewaysOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = directLinkService.ListGateways(listGatewaysOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListGateways with error: Operation request error`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Construct an instance of the ListGatewaysOptions model
				listGatewaysOptionsModel := new(directlinkv1.ListGatewaysOptions)
				listGatewaysOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := directLinkService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := directLinkService.ListGateways(listGatewaysOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ListGateways successfully`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Construct an instance of the ListGatewaysOptions model
				listGatewaysOptionsModel := new(directlinkv1.ListGatewaysOptions)
				listGatewaysOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := directLinkService.ListGateways(listGatewaysOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateGateway(createGatewayOptions *CreateGatewayOptions) - Operation response error`, func() {
		version := "testString"
		createGatewayPath := "/gateways"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createGatewayPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateGateway with error: Operation response processing error`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Construct an instance of the AsPrependTemplate model
				asPrependTemplateModel := new(directlinkv1.AsPrependTemplate)
				asPrependTemplateModel.Length = core.Int64Ptr(int64(4))
				asPrependTemplateModel.Policy = core.StringPtr("import")
				asPrependTemplateModel.Prefix = core.StringPtr("172.17.0.0/16")
				asPrependTemplateModel.SpecificPrefixes = []string{"192.168.3.0/24"}

				// Construct an instance of the AuthenticationKeyIdentityKeyProtectAuthenticationKeyIdentity model
				authenticationKeyIdentityModel := new(directlinkv1.AuthenticationKeyIdentityKeyProtectAuthenticationKeyIdentity)
				authenticationKeyIdentityModel.Crn = core.StringPtr("crn:v1:bluemix:public:kms:us-south:a/4111d05f36894e3cb9b46a43556d9000:abc111b8-37aa-4034-9def-f2607c87aaaa:key:bbb222bc-430a-4de9-9aad-84e5bb022222")

				// Construct an instance of the GatewayBfdConfigTemplate model
				gatewayBfdConfigTemplateModel := new(directlinkv1.GatewayBfdConfigTemplate)
				gatewayBfdConfigTemplateModel.Interval = core.Int64Ptr(int64(2000))
				gatewayBfdConfigTemplateModel.Multiplier = core.Int64Ptr(int64(10))

				// Construct an instance of the GatewayTemplateRouteFilter model
				gatewayTemplateRouteFilterModel := new(directlinkv1.GatewayTemplateRouteFilter)
				gatewayTemplateRouteFilterModel.Action = core.StringPtr("permit")
				gatewayTemplateRouteFilterModel.Ge = core.Int64Ptr(int64(25))
				gatewayTemplateRouteFilterModel.Le = core.Int64Ptr(int64(30))
				gatewayTemplateRouteFilterModel.Prefix = core.StringPtr("192.168.100.0/24")

				// Construct an instance of the ResourceGroupIdentity model
				resourceGroupIdentityModel := new(directlinkv1.ResourceGroupIdentity)
				resourceGroupIdentityModel.ID = core.StringPtr("56969d6043e9465c883cb9f7363e78e8")

				// Construct an instance of the HpcsKeyIdentity model
				hpcsKeyIdentityModel := new(directlinkv1.HpcsKeyIdentity)
				hpcsKeyIdentityModel.Crn = core.StringPtr("crn:v1:bluemix:public:hs-crypto:us-south:a/4111d05f36894e3cb9b46a43556d9000:abc111b8-37aa-4034-9def-f2607c87aaaa:key:bbb222bc-430a-4de9-9aad-84e5bb022222")

				// Construct an instance of the GatewayMacsecCakPrototype model
				gatewayMacsecCakPrototypeModel := new(directlinkv1.GatewayMacsecCakPrototype)
				gatewayMacsecCakPrototypeModel.Key = hpcsKeyIdentityModel
				gatewayMacsecCakPrototypeModel.Name = core.StringPtr("1000")
				gatewayMacsecCakPrototypeModel.Session = core.StringPtr("primary")

				// Construct an instance of the SakRekeyPrototypeSakRekeyTimerModePrototype model
				sakRekeyPrototypeModel := new(directlinkv1.SakRekeyPrototypeSakRekeyTimerModePrototype)
				sakRekeyPrototypeModel.Interval = core.Int64Ptr(int64(3600))
				sakRekeyPrototypeModel.Mode = core.StringPtr("timer")

				// Construct an instance of the GatewayMacsecPrototype model
				gatewayMacsecPrototypeModel := new(directlinkv1.GatewayMacsecPrototype)
				gatewayMacsecPrototypeModel.Active = core.BoolPtr(true)
				gatewayMacsecPrototypeModel.Caks = []directlinkv1.GatewayMacsecCakPrototype{*gatewayMacsecCakPrototypeModel}
				gatewayMacsecPrototypeModel.SakRekey = sakRekeyPrototypeModel
				gatewayMacsecPrototypeModel.SecurityPolicy = core.StringPtr("must_secure")
				gatewayMacsecPrototypeModel.WindowSize = core.Int64Ptr(int64(64))

				// Construct an instance of the GatewayTemplateGatewayTypeDedicatedTemplate model
				gatewayTemplateModel := new(directlinkv1.GatewayTemplateGatewayTypeDedicatedTemplate)
				gatewayTemplateModel.AsPrepends = []directlinkv1.AsPrependTemplate{*asPrependTemplateModel}
				gatewayTemplateModel.AuthenticationKey = authenticationKeyIdentityModel
				gatewayTemplateModel.BfdConfig = gatewayBfdConfigTemplateModel
				gatewayTemplateModel.BgpAsn = core.Int64Ptr(int64(64999))
				gatewayTemplateModel.BgpBaseCidr = core.StringPtr("testString")
				gatewayTemplateModel.BgpCerCidr = core.StringPtr("169.254.0.10/30")
				gatewayTemplateModel.BgpIbmCidr = core.StringPtr("169.254.0.9/30")
				gatewayTemplateModel.ConnectionMode = core.StringPtr("transit")
				gatewayTemplateModel.DefaultExportRouteFilter = core.StringPtr("permit")
				gatewayTemplateModel.DefaultImportRouteFilter = core.StringPtr("permit")
				gatewayTemplateModel.ExportRouteFilters = []directlinkv1.GatewayTemplateRouteFilter{*gatewayTemplateRouteFilterModel}
				gatewayTemplateModel.Global = core.BoolPtr(true)
				gatewayTemplateModel.ImportRouteFilters = []directlinkv1.GatewayTemplateRouteFilter{*gatewayTemplateRouteFilterModel}
				gatewayTemplateModel.Metered = core.BoolPtr(false)
				gatewayTemplateModel.Name = core.StringPtr("myGateway")
				gatewayTemplateModel.PatchPanelCompletionNotice = core.StringPtr("patch panel configuration details")
				gatewayTemplateModel.ResourceGroup = resourceGroupIdentityModel
				gatewayTemplateModel.SpeedMbps = core.Int64Ptr(int64(1000))
				gatewayTemplateModel.Type = core.StringPtr("dedicated")
				gatewayTemplateModel.CarrierName = core.StringPtr("myCarrierName")
				gatewayTemplateModel.CrossConnectRouter = core.StringPtr("xcr01.dal03")
				gatewayTemplateModel.CustomerName = core.StringPtr("newCustomerName")
				gatewayTemplateModel.LocationName = core.StringPtr("dal03")
				gatewayTemplateModel.Macsec = gatewayMacsecPrototypeModel
				gatewayTemplateModel.MacsecCapability = core.StringPtr("non_macsec")
				gatewayTemplateModel.Vlan = core.Int64Ptr(int64(10))

				// Construct an instance of the CreateGatewayOptions model
				createGatewayOptionsModel := new(directlinkv1.CreateGatewayOptions)
				createGatewayOptionsModel.GatewayTemplate = gatewayTemplateModel
				createGatewayOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := directLinkService.CreateGateway(createGatewayOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				directLinkService.EnableRetries(0, 0)
				result, response, operationErr = directLinkService.CreateGateway(createGatewayOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateGateway(createGatewayOptions *CreateGatewayOptions)`, func() {
		version := "testString"
		createGatewayPath := "/gateways"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createGatewayPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"as_prepends": [{"created_at": "2019-01-01T12:00:00.000Z", "id": "ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4", "length": 4, "policy": "import", "prefix": "172.17.0.0/16", "specific_prefixes": ["192.168.3.0/24"], "updated_at": "2019-01-01T12:00:00.000Z"}], "authentication_key": {"crn": "crn:v1:bluemix:public:kms:us-south:a/4111d05f36894e3cb9b46a43556d9000:abc111b8-37aa-4034-9def-f2607c87aaaa:key:bbb222bc-430a-4de9-9aad-84e5bb022222"}, "bfd_config": {"bfd_status": "up", "bfd_status_updated_at": "2020-08-20T06:58:41.909Z", "interval": 2000, "multiplier": 10}, "bgp_asn": 64999, "bgp_base_cidr": "BgpBaseCidr", "bgp_cer_cidr": "10.254.30.78/30", "bgp_ibm_asn": 13884, "bgp_ibm_cidr": "10.254.30.77/30", "bgp_status": "active", "bgp_status_updated_at": "2020-08-20T06:58:41.909Z", "carrier_name": "myCarrierName", "change_request": {"type": "create_gateway"}, "completion_notice_reject_reason": "The completion notice file was blank", "connection_mode": "transit", "created_at": "2019-01-01T12:00:00.000Z", "crn": "crn:v1:bluemix:public:directlink:dal03:a/4111d05f36894e3cb9b46a43556d9000::dedicated:ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4", "cross_account": false, "cross_connect_router": "xcr01.dal03", "customer_name": "newCustomerName", "default_export_route_filter": "permit", "default_import_route_filter": "permit", "global": true, "id": "ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4", "link_status": "up", "link_status_updated_at": "2020-08-20T06:58:41.909Z", "location_display_name": "Dallas 03", "location_name": "dal03", "macsec": {"active": true, "security_policy": "must_secure", "status": "secured", "status_reasons": [{"code": "macsec_cak_failed", "message": "The authentication_key failed configuration.", "more_info": "https://cloud.ibm.com/docs/dl/TODO_ADD_DOCS_LINK"}]}, "macsec_capability": "non_macsec", "metered": false, "name": "myGateway", "operational_status": "awaiting_completion_notice", "operational_status_reasons": [{"code": "authentication_key_failed", "message": "The authentication_key failed configuration.", "more_info": "https://cloud.ibm.com/docs/dl/TODO_ADD_DOCS_LINK"}], "patch_panel_completion_notice": "patch panel configuration details", "port": {"id": "54321b1a-fee4-41c7-9e11-9cd99e000aaa"}, "provider_api_managed": false, "resource_group": {"id": "56969d6043e9465c883cb9f7363e78e8"}, "speed_mbps": 1000, "type": "dedicated", "vlan": 10}`)
				}))
			})
			It(`Invoke CreateGateway successfully with retries`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())
				directLinkService.EnableRetries(0, 0)

				// Construct an instance of the AsPrependTemplate model
				asPrependTemplateModel := new(directlinkv1.AsPrependTemplate)
				asPrependTemplateModel.Length = core.Int64Ptr(int64(4))
				asPrependTemplateModel.Policy = core.StringPtr("import")
				asPrependTemplateModel.Prefix = core.StringPtr("172.17.0.0/16")
				asPrependTemplateModel.SpecificPrefixes = []string{"192.168.3.0/24"}

				// Construct an instance of the AuthenticationKeyIdentityKeyProtectAuthenticationKeyIdentity model
				authenticationKeyIdentityModel := new(directlinkv1.AuthenticationKeyIdentityKeyProtectAuthenticationKeyIdentity)
				authenticationKeyIdentityModel.Crn = core.StringPtr("crn:v1:bluemix:public:kms:us-south:a/4111d05f36894e3cb9b46a43556d9000:abc111b8-37aa-4034-9def-f2607c87aaaa:key:bbb222bc-430a-4de9-9aad-84e5bb022222")

				// Construct an instance of the GatewayBfdConfigTemplate model
				gatewayBfdConfigTemplateModel := new(directlinkv1.GatewayBfdConfigTemplate)
				gatewayBfdConfigTemplateModel.Interval = core.Int64Ptr(int64(2000))
				gatewayBfdConfigTemplateModel.Multiplier = core.Int64Ptr(int64(10))

				// Construct an instance of the GatewayTemplateRouteFilter model
				gatewayTemplateRouteFilterModel := new(directlinkv1.GatewayTemplateRouteFilter)
				gatewayTemplateRouteFilterModel.Action = core.StringPtr("permit")
				gatewayTemplateRouteFilterModel.Ge = core.Int64Ptr(int64(25))
				gatewayTemplateRouteFilterModel.Le = core.Int64Ptr(int64(30))
				gatewayTemplateRouteFilterModel.Prefix = core.StringPtr("192.168.100.0/24")

				// Construct an instance of the ResourceGroupIdentity model
				resourceGroupIdentityModel := new(directlinkv1.ResourceGroupIdentity)
				resourceGroupIdentityModel.ID = core.StringPtr("56969d6043e9465c883cb9f7363e78e8")

				// Construct an instance of the HpcsKeyIdentity model
				hpcsKeyIdentityModel := new(directlinkv1.HpcsKeyIdentity)
				hpcsKeyIdentityModel.Crn = core.StringPtr("crn:v1:bluemix:public:hs-crypto:us-south:a/4111d05f36894e3cb9b46a43556d9000:abc111b8-37aa-4034-9def-f2607c87aaaa:key:bbb222bc-430a-4de9-9aad-84e5bb022222")

				// Construct an instance of the GatewayMacsecCakPrototype model
				gatewayMacsecCakPrototypeModel := new(directlinkv1.GatewayMacsecCakPrototype)
				gatewayMacsecCakPrototypeModel.Key = hpcsKeyIdentityModel
				gatewayMacsecCakPrototypeModel.Name = core.StringPtr("1000")
				gatewayMacsecCakPrototypeModel.Session = core.StringPtr("primary")

				// Construct an instance of the SakRekeyPrototypeSakRekeyTimerModePrototype model
				sakRekeyPrototypeModel := new(directlinkv1.SakRekeyPrototypeSakRekeyTimerModePrototype)
				sakRekeyPrototypeModel.Interval = core.Int64Ptr(int64(3600))
				sakRekeyPrototypeModel.Mode = core.StringPtr("timer")

				// Construct an instance of the GatewayMacsecPrototype model
				gatewayMacsecPrototypeModel := new(directlinkv1.GatewayMacsecPrototype)
				gatewayMacsecPrototypeModel.Active = core.BoolPtr(true)
				gatewayMacsecPrototypeModel.Caks = []directlinkv1.GatewayMacsecCakPrototype{*gatewayMacsecCakPrototypeModel}
				gatewayMacsecPrototypeModel.SakRekey = sakRekeyPrototypeModel
				gatewayMacsecPrototypeModel.SecurityPolicy = core.StringPtr("must_secure")
				gatewayMacsecPrototypeModel.WindowSize = core.Int64Ptr(int64(64))

				// Construct an instance of the GatewayTemplateGatewayTypeDedicatedTemplate model
				gatewayTemplateModel := new(directlinkv1.GatewayTemplateGatewayTypeDedicatedTemplate)
				gatewayTemplateModel.AsPrepends = []directlinkv1.AsPrependTemplate{*asPrependTemplateModel}
				gatewayTemplateModel.AuthenticationKey = authenticationKeyIdentityModel
				gatewayTemplateModel.BfdConfig = gatewayBfdConfigTemplateModel
				gatewayTemplateModel.BgpAsn = core.Int64Ptr(int64(64999))
				gatewayTemplateModel.BgpBaseCidr = core.StringPtr("testString")
				gatewayTemplateModel.BgpCerCidr = core.StringPtr("169.254.0.10/30")
				gatewayTemplateModel.BgpIbmCidr = core.StringPtr("169.254.0.9/30")
				gatewayTemplateModel.ConnectionMode = core.StringPtr("transit")
				gatewayTemplateModel.DefaultExportRouteFilter = core.StringPtr("permit")
				gatewayTemplateModel.DefaultImportRouteFilter = core.StringPtr("permit")
				gatewayTemplateModel.ExportRouteFilters = []directlinkv1.GatewayTemplateRouteFilter{*gatewayTemplateRouteFilterModel}
				gatewayTemplateModel.Global = core.BoolPtr(true)
				gatewayTemplateModel.ImportRouteFilters = []directlinkv1.GatewayTemplateRouteFilter{*gatewayTemplateRouteFilterModel}
				gatewayTemplateModel.Metered = core.BoolPtr(false)
				gatewayTemplateModel.Name = core.StringPtr("myGateway")
				gatewayTemplateModel.PatchPanelCompletionNotice = core.StringPtr("patch panel configuration details")
				gatewayTemplateModel.ResourceGroup = resourceGroupIdentityModel
				gatewayTemplateModel.SpeedMbps = core.Int64Ptr(int64(1000))
				gatewayTemplateModel.Type = core.StringPtr("dedicated")
				gatewayTemplateModel.CarrierName = core.StringPtr("myCarrierName")
				gatewayTemplateModel.CrossConnectRouter = core.StringPtr("xcr01.dal03")
				gatewayTemplateModel.CustomerName = core.StringPtr("newCustomerName")
				gatewayTemplateModel.LocationName = core.StringPtr("dal03")
				gatewayTemplateModel.Macsec = gatewayMacsecPrototypeModel
				gatewayTemplateModel.MacsecCapability = core.StringPtr("non_macsec")
				gatewayTemplateModel.Vlan = core.Int64Ptr(int64(10))

				// Construct an instance of the CreateGatewayOptions model
				createGatewayOptionsModel := new(directlinkv1.CreateGatewayOptions)
				createGatewayOptionsModel.GatewayTemplate = gatewayTemplateModel
				createGatewayOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := directLinkService.CreateGatewayWithContext(ctx, createGatewayOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				directLinkService.DisableRetries()
				result, response, operationErr := directLinkService.CreateGateway(createGatewayOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = directLinkService.CreateGatewayWithContext(ctx, createGatewayOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createGatewayPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"as_prepends": [{"created_at": "2019-01-01T12:00:00.000Z", "id": "ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4", "length": 4, "policy": "import", "prefix": "172.17.0.0/16", "specific_prefixes": ["192.168.3.0/24"], "updated_at": "2019-01-01T12:00:00.000Z"}], "authentication_key": {"crn": "crn:v1:bluemix:public:kms:us-south:a/4111d05f36894e3cb9b46a43556d9000:abc111b8-37aa-4034-9def-f2607c87aaaa:key:bbb222bc-430a-4de9-9aad-84e5bb022222"}, "bfd_config": {"bfd_status": "up", "bfd_status_updated_at": "2020-08-20T06:58:41.909Z", "interval": 2000, "multiplier": 10}, "bgp_asn": 64999, "bgp_base_cidr": "BgpBaseCidr", "bgp_cer_cidr": "10.254.30.78/30", "bgp_ibm_asn": 13884, "bgp_ibm_cidr": "10.254.30.77/30", "bgp_status": "active", "bgp_status_updated_at": "2020-08-20T06:58:41.909Z", "carrier_name": "myCarrierName", "change_request": {"type": "create_gateway"}, "completion_notice_reject_reason": "The completion notice file was blank", "connection_mode": "transit", "created_at": "2019-01-01T12:00:00.000Z", "crn": "crn:v1:bluemix:public:directlink:dal03:a/4111d05f36894e3cb9b46a43556d9000::dedicated:ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4", "cross_account": false, "cross_connect_router": "xcr01.dal03", "customer_name": "newCustomerName", "default_export_route_filter": "permit", "default_import_route_filter": "permit", "global": true, "id": "ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4", "link_status": "up", "link_status_updated_at": "2020-08-20T06:58:41.909Z", "location_display_name": "Dallas 03", "location_name": "dal03", "macsec": {"active": true, "security_policy": "must_secure", "status": "secured", "status_reasons": [{"code": "macsec_cak_failed", "message": "The authentication_key failed configuration.", "more_info": "https://cloud.ibm.com/docs/dl/TODO_ADD_DOCS_LINK"}]}, "macsec_capability": "non_macsec", "metered": false, "name": "myGateway", "operational_status": "awaiting_completion_notice", "operational_status_reasons": [{"code": "authentication_key_failed", "message": "The authentication_key failed configuration.", "more_info": "https://cloud.ibm.com/docs/dl/TODO_ADD_DOCS_LINK"}], "patch_panel_completion_notice": "patch panel configuration details", "port": {"id": "54321b1a-fee4-41c7-9e11-9cd99e000aaa"}, "provider_api_managed": false, "resource_group": {"id": "56969d6043e9465c883cb9f7363e78e8"}, "speed_mbps": 1000, "type": "dedicated", "vlan": 10}`)
				}))
			})
			It(`Invoke CreateGateway successfully`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := directLinkService.CreateGateway(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the AsPrependTemplate model
				asPrependTemplateModel := new(directlinkv1.AsPrependTemplate)
				asPrependTemplateModel.Length = core.Int64Ptr(int64(4))
				asPrependTemplateModel.Policy = core.StringPtr("import")
				asPrependTemplateModel.Prefix = core.StringPtr("172.17.0.0/16")
				asPrependTemplateModel.SpecificPrefixes = []string{"192.168.3.0/24"}

				// Construct an instance of the AuthenticationKeyIdentityKeyProtectAuthenticationKeyIdentity model
				authenticationKeyIdentityModel := new(directlinkv1.AuthenticationKeyIdentityKeyProtectAuthenticationKeyIdentity)
				authenticationKeyIdentityModel.Crn = core.StringPtr("crn:v1:bluemix:public:kms:us-south:a/4111d05f36894e3cb9b46a43556d9000:abc111b8-37aa-4034-9def-f2607c87aaaa:key:bbb222bc-430a-4de9-9aad-84e5bb022222")

				// Construct an instance of the GatewayBfdConfigTemplate model
				gatewayBfdConfigTemplateModel := new(directlinkv1.GatewayBfdConfigTemplate)
				gatewayBfdConfigTemplateModel.Interval = core.Int64Ptr(int64(2000))
				gatewayBfdConfigTemplateModel.Multiplier = core.Int64Ptr(int64(10))

				// Construct an instance of the GatewayTemplateRouteFilter model
				gatewayTemplateRouteFilterModel := new(directlinkv1.GatewayTemplateRouteFilter)
				gatewayTemplateRouteFilterModel.Action = core.StringPtr("permit")
				gatewayTemplateRouteFilterModel.Ge = core.Int64Ptr(int64(25))
				gatewayTemplateRouteFilterModel.Le = core.Int64Ptr(int64(30))
				gatewayTemplateRouteFilterModel.Prefix = core.StringPtr("192.168.100.0/24")

				// Construct an instance of the ResourceGroupIdentity model
				resourceGroupIdentityModel := new(directlinkv1.ResourceGroupIdentity)
				resourceGroupIdentityModel.ID = core.StringPtr("56969d6043e9465c883cb9f7363e78e8")

				// Construct an instance of the HpcsKeyIdentity model
				hpcsKeyIdentityModel := new(directlinkv1.HpcsKeyIdentity)
				hpcsKeyIdentityModel.Crn = core.StringPtr("crn:v1:bluemix:public:hs-crypto:us-south:a/4111d05f36894e3cb9b46a43556d9000:abc111b8-37aa-4034-9def-f2607c87aaaa:key:bbb222bc-430a-4de9-9aad-84e5bb022222")

				// Construct an instance of the GatewayMacsecCakPrototype model
				gatewayMacsecCakPrototypeModel := new(directlinkv1.GatewayMacsecCakPrototype)
				gatewayMacsecCakPrototypeModel.Key = hpcsKeyIdentityModel
				gatewayMacsecCakPrototypeModel.Name = core.StringPtr("1000")
				gatewayMacsecCakPrototypeModel.Session = core.StringPtr("primary")

				// Construct an instance of the SakRekeyPrototypeSakRekeyTimerModePrototype model
				sakRekeyPrototypeModel := new(directlinkv1.SakRekeyPrototypeSakRekeyTimerModePrototype)
				sakRekeyPrototypeModel.Interval = core.Int64Ptr(int64(3600))
				sakRekeyPrototypeModel.Mode = core.StringPtr("timer")

				// Construct an instance of the GatewayMacsecPrototype model
				gatewayMacsecPrototypeModel := new(directlinkv1.GatewayMacsecPrototype)
				gatewayMacsecPrototypeModel.Active = core.BoolPtr(true)
				gatewayMacsecPrototypeModel.Caks = []directlinkv1.GatewayMacsecCakPrototype{*gatewayMacsecCakPrototypeModel}
				gatewayMacsecPrototypeModel.SakRekey = sakRekeyPrototypeModel
				gatewayMacsecPrototypeModel.SecurityPolicy = core.StringPtr("must_secure")
				gatewayMacsecPrototypeModel.WindowSize = core.Int64Ptr(int64(64))

				// Construct an instance of the GatewayTemplateGatewayTypeDedicatedTemplate model
				gatewayTemplateModel := new(directlinkv1.GatewayTemplateGatewayTypeDedicatedTemplate)
				gatewayTemplateModel.AsPrepends = []directlinkv1.AsPrependTemplate{*asPrependTemplateModel}
				gatewayTemplateModel.AuthenticationKey = authenticationKeyIdentityModel
				gatewayTemplateModel.BfdConfig = gatewayBfdConfigTemplateModel
				gatewayTemplateModel.BgpAsn = core.Int64Ptr(int64(64999))
				gatewayTemplateModel.BgpBaseCidr = core.StringPtr("testString")
				gatewayTemplateModel.BgpCerCidr = core.StringPtr("169.254.0.10/30")
				gatewayTemplateModel.BgpIbmCidr = core.StringPtr("169.254.0.9/30")
				gatewayTemplateModel.ConnectionMode = core.StringPtr("transit")
				gatewayTemplateModel.DefaultExportRouteFilter = core.StringPtr("permit")
				gatewayTemplateModel.DefaultImportRouteFilter = core.StringPtr("permit")
				gatewayTemplateModel.ExportRouteFilters = []directlinkv1.GatewayTemplateRouteFilter{*gatewayTemplateRouteFilterModel}
				gatewayTemplateModel.Global = core.BoolPtr(true)
				gatewayTemplateModel.ImportRouteFilters = []directlinkv1.GatewayTemplateRouteFilter{*gatewayTemplateRouteFilterModel}
				gatewayTemplateModel.Metered = core.BoolPtr(false)
				gatewayTemplateModel.Name = core.StringPtr("myGateway")
				gatewayTemplateModel.PatchPanelCompletionNotice = core.StringPtr("patch panel configuration details")
				gatewayTemplateModel.ResourceGroup = resourceGroupIdentityModel
				gatewayTemplateModel.SpeedMbps = core.Int64Ptr(int64(1000))
				gatewayTemplateModel.Type = core.StringPtr("dedicated")
				gatewayTemplateModel.CarrierName = core.StringPtr("myCarrierName")
				gatewayTemplateModel.CrossConnectRouter = core.StringPtr("xcr01.dal03")
				gatewayTemplateModel.CustomerName = core.StringPtr("newCustomerName")
				gatewayTemplateModel.LocationName = core.StringPtr("dal03")
				gatewayTemplateModel.Macsec = gatewayMacsecPrototypeModel
				gatewayTemplateModel.MacsecCapability = core.StringPtr("non_macsec")
				gatewayTemplateModel.Vlan = core.Int64Ptr(int64(10))

				// Construct an instance of the CreateGatewayOptions model
				createGatewayOptionsModel := new(directlinkv1.CreateGatewayOptions)
				createGatewayOptionsModel.GatewayTemplate = gatewayTemplateModel
				createGatewayOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = directLinkService.CreateGateway(createGatewayOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateGateway with error: Operation validation and request error`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Construct an instance of the AsPrependTemplate model
				asPrependTemplateModel := new(directlinkv1.AsPrependTemplate)
				asPrependTemplateModel.Length = core.Int64Ptr(int64(4))
				asPrependTemplateModel.Policy = core.StringPtr("import")
				asPrependTemplateModel.Prefix = core.StringPtr("172.17.0.0/16")
				asPrependTemplateModel.SpecificPrefixes = []string{"192.168.3.0/24"}

				// Construct an instance of the AuthenticationKeyIdentityKeyProtectAuthenticationKeyIdentity model
				authenticationKeyIdentityModel := new(directlinkv1.AuthenticationKeyIdentityKeyProtectAuthenticationKeyIdentity)
				authenticationKeyIdentityModel.Crn = core.StringPtr("crn:v1:bluemix:public:kms:us-south:a/4111d05f36894e3cb9b46a43556d9000:abc111b8-37aa-4034-9def-f2607c87aaaa:key:bbb222bc-430a-4de9-9aad-84e5bb022222")

				// Construct an instance of the GatewayBfdConfigTemplate model
				gatewayBfdConfigTemplateModel := new(directlinkv1.GatewayBfdConfigTemplate)
				gatewayBfdConfigTemplateModel.Interval = core.Int64Ptr(int64(2000))
				gatewayBfdConfigTemplateModel.Multiplier = core.Int64Ptr(int64(10))

				// Construct an instance of the GatewayTemplateRouteFilter model
				gatewayTemplateRouteFilterModel := new(directlinkv1.GatewayTemplateRouteFilter)
				gatewayTemplateRouteFilterModel.Action = core.StringPtr("permit")
				gatewayTemplateRouteFilterModel.Ge = core.Int64Ptr(int64(25))
				gatewayTemplateRouteFilterModel.Le = core.Int64Ptr(int64(30))
				gatewayTemplateRouteFilterModel.Prefix = core.StringPtr("192.168.100.0/24")

				// Construct an instance of the ResourceGroupIdentity model
				resourceGroupIdentityModel := new(directlinkv1.ResourceGroupIdentity)
				resourceGroupIdentityModel.ID = core.StringPtr("56969d6043e9465c883cb9f7363e78e8")

				// Construct an instance of the HpcsKeyIdentity model
				hpcsKeyIdentityModel := new(directlinkv1.HpcsKeyIdentity)
				hpcsKeyIdentityModel.Crn = core.StringPtr("crn:v1:bluemix:public:hs-crypto:us-south:a/4111d05f36894e3cb9b46a43556d9000:abc111b8-37aa-4034-9def-f2607c87aaaa:key:bbb222bc-430a-4de9-9aad-84e5bb022222")

				// Construct an instance of the GatewayMacsecCakPrototype model
				gatewayMacsecCakPrototypeModel := new(directlinkv1.GatewayMacsecCakPrototype)
				gatewayMacsecCakPrototypeModel.Key = hpcsKeyIdentityModel
				gatewayMacsecCakPrototypeModel.Name = core.StringPtr("1000")
				gatewayMacsecCakPrototypeModel.Session = core.StringPtr("primary")

				// Construct an instance of the SakRekeyPrototypeSakRekeyTimerModePrototype model
				sakRekeyPrototypeModel := new(directlinkv1.SakRekeyPrototypeSakRekeyTimerModePrototype)
				sakRekeyPrototypeModel.Interval = core.Int64Ptr(int64(3600))
				sakRekeyPrototypeModel.Mode = core.StringPtr("timer")

				// Construct an instance of the GatewayMacsecPrototype model
				gatewayMacsecPrototypeModel := new(directlinkv1.GatewayMacsecPrototype)
				gatewayMacsecPrototypeModel.Active = core.BoolPtr(true)
				gatewayMacsecPrototypeModel.Caks = []directlinkv1.GatewayMacsecCakPrototype{*gatewayMacsecCakPrototypeModel}
				gatewayMacsecPrototypeModel.SakRekey = sakRekeyPrototypeModel
				gatewayMacsecPrototypeModel.SecurityPolicy = core.StringPtr("must_secure")
				gatewayMacsecPrototypeModel.WindowSize = core.Int64Ptr(int64(64))

				// Construct an instance of the GatewayTemplateGatewayTypeDedicatedTemplate model
				gatewayTemplateModel := new(directlinkv1.GatewayTemplateGatewayTypeDedicatedTemplate)
				gatewayTemplateModel.AsPrepends = []directlinkv1.AsPrependTemplate{*asPrependTemplateModel}
				gatewayTemplateModel.AuthenticationKey = authenticationKeyIdentityModel
				gatewayTemplateModel.BfdConfig = gatewayBfdConfigTemplateModel
				gatewayTemplateModel.BgpAsn = core.Int64Ptr(int64(64999))
				gatewayTemplateModel.BgpBaseCidr = core.StringPtr("testString")
				gatewayTemplateModel.BgpCerCidr = core.StringPtr("169.254.0.10/30")
				gatewayTemplateModel.BgpIbmCidr = core.StringPtr("169.254.0.9/30")
				gatewayTemplateModel.ConnectionMode = core.StringPtr("transit")
				gatewayTemplateModel.DefaultExportRouteFilter = core.StringPtr("permit")
				gatewayTemplateModel.DefaultImportRouteFilter = core.StringPtr("permit")
				gatewayTemplateModel.ExportRouteFilters = []directlinkv1.GatewayTemplateRouteFilter{*gatewayTemplateRouteFilterModel}
				gatewayTemplateModel.Global = core.BoolPtr(true)
				gatewayTemplateModel.ImportRouteFilters = []directlinkv1.GatewayTemplateRouteFilter{*gatewayTemplateRouteFilterModel}
				gatewayTemplateModel.Metered = core.BoolPtr(false)
				gatewayTemplateModel.Name = core.StringPtr("myGateway")
				gatewayTemplateModel.PatchPanelCompletionNotice = core.StringPtr("patch panel configuration details")
				gatewayTemplateModel.ResourceGroup = resourceGroupIdentityModel
				gatewayTemplateModel.SpeedMbps = core.Int64Ptr(int64(1000))
				gatewayTemplateModel.Type = core.StringPtr("dedicated")
				gatewayTemplateModel.CarrierName = core.StringPtr("myCarrierName")
				gatewayTemplateModel.CrossConnectRouter = core.StringPtr("xcr01.dal03")
				gatewayTemplateModel.CustomerName = core.StringPtr("newCustomerName")
				gatewayTemplateModel.LocationName = core.StringPtr("dal03")
				gatewayTemplateModel.Macsec = gatewayMacsecPrototypeModel
				gatewayTemplateModel.MacsecCapability = core.StringPtr("non_macsec")
				gatewayTemplateModel.Vlan = core.Int64Ptr(int64(10))

				// Construct an instance of the CreateGatewayOptions model
				createGatewayOptionsModel := new(directlinkv1.CreateGatewayOptions)
				createGatewayOptionsModel.GatewayTemplate = gatewayTemplateModel
				createGatewayOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := directLinkService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := directLinkService.CreateGateway(createGatewayOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateGatewayOptions model with no property values
				createGatewayOptionsModelNew := new(directlinkv1.CreateGatewayOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = directLinkService.CreateGateway(createGatewayOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(201)
				}))
			})
			It(`Invoke CreateGateway successfully`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Construct an instance of the AsPrependTemplate model
				asPrependTemplateModel := new(directlinkv1.AsPrependTemplate)
				asPrependTemplateModel.Length = core.Int64Ptr(int64(4))
				asPrependTemplateModel.Policy = core.StringPtr("import")
				asPrependTemplateModel.Prefix = core.StringPtr("172.17.0.0/16")
				asPrependTemplateModel.SpecificPrefixes = []string{"192.168.3.0/24"}

				// Construct an instance of the AuthenticationKeyIdentityKeyProtectAuthenticationKeyIdentity model
				authenticationKeyIdentityModel := new(directlinkv1.AuthenticationKeyIdentityKeyProtectAuthenticationKeyIdentity)
				authenticationKeyIdentityModel.Crn = core.StringPtr("crn:v1:bluemix:public:kms:us-south:a/4111d05f36894e3cb9b46a43556d9000:abc111b8-37aa-4034-9def-f2607c87aaaa:key:bbb222bc-430a-4de9-9aad-84e5bb022222")

				// Construct an instance of the GatewayBfdConfigTemplate model
				gatewayBfdConfigTemplateModel := new(directlinkv1.GatewayBfdConfigTemplate)
				gatewayBfdConfigTemplateModel.Interval = core.Int64Ptr(int64(2000))
				gatewayBfdConfigTemplateModel.Multiplier = core.Int64Ptr(int64(10))

				// Construct an instance of the GatewayTemplateRouteFilter model
				gatewayTemplateRouteFilterModel := new(directlinkv1.GatewayTemplateRouteFilter)
				gatewayTemplateRouteFilterModel.Action = core.StringPtr("permit")
				gatewayTemplateRouteFilterModel.Ge = core.Int64Ptr(int64(25))
				gatewayTemplateRouteFilterModel.Le = core.Int64Ptr(int64(30))
				gatewayTemplateRouteFilterModel.Prefix = core.StringPtr("192.168.100.0/24")

				// Construct an instance of the ResourceGroupIdentity model
				resourceGroupIdentityModel := new(directlinkv1.ResourceGroupIdentity)
				resourceGroupIdentityModel.ID = core.StringPtr("56969d6043e9465c883cb9f7363e78e8")

				// Construct an instance of the HpcsKeyIdentity model
				hpcsKeyIdentityModel := new(directlinkv1.HpcsKeyIdentity)
				hpcsKeyIdentityModel.Crn = core.StringPtr("crn:v1:bluemix:public:hs-crypto:us-south:a/4111d05f36894e3cb9b46a43556d9000:abc111b8-37aa-4034-9def-f2607c87aaaa:key:bbb222bc-430a-4de9-9aad-84e5bb022222")

				// Construct an instance of the GatewayMacsecCakPrototype model
				gatewayMacsecCakPrototypeModel := new(directlinkv1.GatewayMacsecCakPrototype)
				gatewayMacsecCakPrototypeModel.Key = hpcsKeyIdentityModel
				gatewayMacsecCakPrototypeModel.Name = core.StringPtr("1000")
				gatewayMacsecCakPrototypeModel.Session = core.StringPtr("primary")

				// Construct an instance of the SakRekeyPrototypeSakRekeyTimerModePrototype model
				sakRekeyPrototypeModel := new(directlinkv1.SakRekeyPrototypeSakRekeyTimerModePrototype)
				sakRekeyPrototypeModel.Interval = core.Int64Ptr(int64(3600))
				sakRekeyPrototypeModel.Mode = core.StringPtr("timer")

				// Construct an instance of the GatewayMacsecPrototype model
				gatewayMacsecPrototypeModel := new(directlinkv1.GatewayMacsecPrototype)
				gatewayMacsecPrototypeModel.Active = core.BoolPtr(true)
				gatewayMacsecPrototypeModel.Caks = []directlinkv1.GatewayMacsecCakPrototype{*gatewayMacsecCakPrototypeModel}
				gatewayMacsecPrototypeModel.SakRekey = sakRekeyPrototypeModel
				gatewayMacsecPrototypeModel.SecurityPolicy = core.StringPtr("must_secure")
				gatewayMacsecPrototypeModel.WindowSize = core.Int64Ptr(int64(64))

				// Construct an instance of the GatewayTemplateGatewayTypeDedicatedTemplate model
				gatewayTemplateModel := new(directlinkv1.GatewayTemplateGatewayTypeDedicatedTemplate)
				gatewayTemplateModel.AsPrepends = []directlinkv1.AsPrependTemplate{*asPrependTemplateModel}
				gatewayTemplateModel.AuthenticationKey = authenticationKeyIdentityModel
				gatewayTemplateModel.BfdConfig = gatewayBfdConfigTemplateModel
				gatewayTemplateModel.BgpAsn = core.Int64Ptr(int64(64999))
				gatewayTemplateModel.BgpBaseCidr = core.StringPtr("testString")
				gatewayTemplateModel.BgpCerCidr = core.StringPtr("169.254.0.10/30")
				gatewayTemplateModel.BgpIbmCidr = core.StringPtr("169.254.0.9/30")
				gatewayTemplateModel.ConnectionMode = core.StringPtr("transit")
				gatewayTemplateModel.DefaultExportRouteFilter = core.StringPtr("permit")
				gatewayTemplateModel.DefaultImportRouteFilter = core.StringPtr("permit")
				gatewayTemplateModel.ExportRouteFilters = []directlinkv1.GatewayTemplateRouteFilter{*gatewayTemplateRouteFilterModel}
				gatewayTemplateModel.Global = core.BoolPtr(true)
				gatewayTemplateModel.ImportRouteFilters = []directlinkv1.GatewayTemplateRouteFilter{*gatewayTemplateRouteFilterModel}
				gatewayTemplateModel.Metered = core.BoolPtr(false)
				gatewayTemplateModel.Name = core.StringPtr("myGateway")
				gatewayTemplateModel.PatchPanelCompletionNotice = core.StringPtr("patch panel configuration details")
				gatewayTemplateModel.ResourceGroup = resourceGroupIdentityModel
				gatewayTemplateModel.SpeedMbps = core.Int64Ptr(int64(1000))
				gatewayTemplateModel.Type = core.StringPtr("dedicated")
				gatewayTemplateModel.CarrierName = core.StringPtr("myCarrierName")
				gatewayTemplateModel.CrossConnectRouter = core.StringPtr("xcr01.dal03")
				gatewayTemplateModel.CustomerName = core.StringPtr("newCustomerName")
				gatewayTemplateModel.LocationName = core.StringPtr("dal03")
				gatewayTemplateModel.Macsec = gatewayMacsecPrototypeModel
				gatewayTemplateModel.MacsecCapability = core.StringPtr("non_macsec")
				gatewayTemplateModel.Vlan = core.Int64Ptr(int64(10))

				// Construct an instance of the CreateGatewayOptions model
				createGatewayOptionsModel := new(directlinkv1.CreateGatewayOptions)
				createGatewayOptionsModel.GatewayTemplate = gatewayTemplateModel
				createGatewayOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := directLinkService.CreateGateway(createGatewayOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteGateway(deleteGatewayOptions *DeleteGatewayOptions)`, func() {
		version := "testString"
		deleteGatewayPath := "/gateways/0a06fb9b-820f-4c44-8a31-77f1f0806d28"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteGatewayPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteGateway successfully`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := directLinkService.DeleteGateway(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteGatewayOptions model
				deleteGatewayOptionsModel := new(directlinkv1.DeleteGatewayOptions)
				deleteGatewayOptionsModel.ID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				deleteGatewayOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = directLinkService.DeleteGateway(deleteGatewayOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteGateway with error: Operation validation and request error`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Construct an instance of the DeleteGatewayOptions model
				deleteGatewayOptionsModel := new(directlinkv1.DeleteGatewayOptions)
				deleteGatewayOptionsModel.ID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				deleteGatewayOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := directLinkService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := directLinkService.DeleteGateway(deleteGatewayOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteGatewayOptions model with no property values
				deleteGatewayOptionsModelNew := new(directlinkv1.DeleteGatewayOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = directLinkService.DeleteGateway(deleteGatewayOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetGateway(getGatewayOptions *GetGatewayOptions) - Operation response error`, func() {
		version := "testString"
		getGatewayPath := "/gateways/0a06fb9b-820f-4c44-8a31-77f1f0806d28"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getGatewayPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetGateway with error: Operation response processing error`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Construct an instance of the GetGatewayOptions model
				getGatewayOptionsModel := new(directlinkv1.GetGatewayOptions)
				getGatewayOptionsModel.ID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				getGatewayOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := directLinkService.GetGateway(getGatewayOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				directLinkService.EnableRetries(0, 0)
				result, response, operationErr = directLinkService.GetGateway(getGatewayOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetGateway(getGatewayOptions *GetGatewayOptions)`, func() {
		version := "testString"
		getGatewayPath := "/gateways/0a06fb9b-820f-4c44-8a31-77f1f0806d28"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getGatewayPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"as_prepends": [{"created_at": "2019-01-01T12:00:00.000Z", "id": "ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4", "length": 4, "policy": "import", "prefix": "172.17.0.0/16", "specific_prefixes": ["192.168.3.0/24"], "updated_at": "2019-01-01T12:00:00.000Z"}], "authentication_key": {"crn": "crn:v1:bluemix:public:kms:us-south:a/4111d05f36894e3cb9b46a43556d9000:abc111b8-37aa-4034-9def-f2607c87aaaa:key:bbb222bc-430a-4de9-9aad-84e5bb022222"}, "bfd_config": {"bfd_status": "up", "bfd_status_updated_at": "2020-08-20T06:58:41.909Z", "interval": 2000, "multiplier": 10}, "bgp_asn": 64999, "bgp_base_cidr": "BgpBaseCidr", "bgp_cer_cidr": "10.254.30.78/30", "bgp_ibm_asn": 13884, "bgp_ibm_cidr": "10.254.30.77/30", "bgp_status": "active", "bgp_status_updated_at": "2020-08-20T06:58:41.909Z", "carrier_name": "myCarrierName", "change_request": {"type": "create_gateway"}, "completion_notice_reject_reason": "The completion notice file was blank", "connection_mode": "transit", "created_at": "2019-01-01T12:00:00.000Z", "crn": "crn:v1:bluemix:public:directlink:dal03:a/4111d05f36894e3cb9b46a43556d9000::dedicated:ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4", "cross_account": false, "cross_connect_router": "xcr01.dal03", "customer_name": "newCustomerName", "default_export_route_filter": "permit", "default_import_route_filter": "permit", "global": true, "id": "ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4", "link_status": "up", "link_status_updated_at": "2020-08-20T06:58:41.909Z", "location_display_name": "Dallas 03", "location_name": "dal03", "macsec": {"active": true, "security_policy": "must_secure", "status": "secured", "status_reasons": [{"code": "macsec_cak_failed", "message": "The authentication_key failed configuration.", "more_info": "https://cloud.ibm.com/docs/dl/TODO_ADD_DOCS_LINK"}]}, "macsec_capability": "non_macsec", "metered": false, "name": "myGateway", "operational_status": "awaiting_completion_notice", "operational_status_reasons": [{"code": "authentication_key_failed", "message": "The authentication_key failed configuration.", "more_info": "https://cloud.ibm.com/docs/dl/TODO_ADD_DOCS_LINK"}], "patch_panel_completion_notice": "patch panel configuration details", "port": {"id": "54321b1a-fee4-41c7-9e11-9cd99e000aaa"}, "provider_api_managed": false, "resource_group": {"id": "56969d6043e9465c883cb9f7363e78e8"}, "speed_mbps": 1000, "type": "dedicated", "vlan": 10}`)
				}))
			})
			It(`Invoke GetGateway successfully with retries`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())
				directLinkService.EnableRetries(0, 0)

				// Construct an instance of the GetGatewayOptions model
				getGatewayOptionsModel := new(directlinkv1.GetGatewayOptions)
				getGatewayOptionsModel.ID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				getGatewayOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := directLinkService.GetGatewayWithContext(ctx, getGatewayOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				directLinkService.DisableRetries()
				result, response, operationErr := directLinkService.GetGateway(getGatewayOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = directLinkService.GetGatewayWithContext(ctx, getGatewayOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getGatewayPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"as_prepends": [{"created_at": "2019-01-01T12:00:00.000Z", "id": "ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4", "length": 4, "policy": "import", "prefix": "172.17.0.0/16", "specific_prefixes": ["192.168.3.0/24"], "updated_at": "2019-01-01T12:00:00.000Z"}], "authentication_key": {"crn": "crn:v1:bluemix:public:kms:us-south:a/4111d05f36894e3cb9b46a43556d9000:abc111b8-37aa-4034-9def-f2607c87aaaa:key:bbb222bc-430a-4de9-9aad-84e5bb022222"}, "bfd_config": {"bfd_status": "up", "bfd_status_updated_at": "2020-08-20T06:58:41.909Z", "interval": 2000, "multiplier": 10}, "bgp_asn": 64999, "bgp_base_cidr": "BgpBaseCidr", "bgp_cer_cidr": "10.254.30.78/30", "bgp_ibm_asn": 13884, "bgp_ibm_cidr": "10.254.30.77/30", "bgp_status": "active", "bgp_status_updated_at": "2020-08-20T06:58:41.909Z", "carrier_name": "myCarrierName", "change_request": {"type": "create_gateway"}, "completion_notice_reject_reason": "The completion notice file was blank", "connection_mode": "transit", "created_at": "2019-01-01T12:00:00.000Z", "crn": "crn:v1:bluemix:public:directlink:dal03:a/4111d05f36894e3cb9b46a43556d9000::dedicated:ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4", "cross_account": false, "cross_connect_router": "xcr01.dal03", "customer_name": "newCustomerName", "default_export_route_filter": "permit", "default_import_route_filter": "permit", "global": true, "id": "ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4", "link_status": "up", "link_status_updated_at": "2020-08-20T06:58:41.909Z", "location_display_name": "Dallas 03", "location_name": "dal03", "macsec": {"active": true, "security_policy": "must_secure", "status": "secured", "status_reasons": [{"code": "macsec_cak_failed", "message": "The authentication_key failed configuration.", "more_info": "https://cloud.ibm.com/docs/dl/TODO_ADD_DOCS_LINK"}]}, "macsec_capability": "non_macsec", "metered": false, "name": "myGateway", "operational_status": "awaiting_completion_notice", "operational_status_reasons": [{"code": "authentication_key_failed", "message": "The authentication_key failed configuration.", "more_info": "https://cloud.ibm.com/docs/dl/TODO_ADD_DOCS_LINK"}], "patch_panel_completion_notice": "patch panel configuration details", "port": {"id": "54321b1a-fee4-41c7-9e11-9cd99e000aaa"}, "provider_api_managed": false, "resource_group": {"id": "56969d6043e9465c883cb9f7363e78e8"}, "speed_mbps": 1000, "type": "dedicated", "vlan": 10}`)
				}))
			})
			It(`Invoke GetGateway successfully`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := directLinkService.GetGateway(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetGatewayOptions model
				getGatewayOptionsModel := new(directlinkv1.GetGatewayOptions)
				getGatewayOptionsModel.ID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				getGatewayOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = directLinkService.GetGateway(getGatewayOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetGateway with error: Operation validation and request error`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Construct an instance of the GetGatewayOptions model
				getGatewayOptionsModel := new(directlinkv1.GetGatewayOptions)
				getGatewayOptionsModel.ID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				getGatewayOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := directLinkService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := directLinkService.GetGateway(getGatewayOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetGatewayOptions model with no property values
				getGatewayOptionsModelNew := new(directlinkv1.GetGatewayOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = directLinkService.GetGateway(getGatewayOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetGateway successfully`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Construct an instance of the GetGatewayOptions model
				getGatewayOptionsModel := new(directlinkv1.GetGatewayOptions)
				getGatewayOptionsModel.ID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				getGatewayOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := directLinkService.GetGateway(getGatewayOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateGateway(updateGatewayOptions *UpdateGatewayOptions) - Operation response error`, func() {
		version := "testString"
		updateGatewayPath := "/gateways/0a06fb9b-820f-4c44-8a31-77f1f0806d28"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateGatewayPath))
					Expect(req.Method).To(Equal("PATCH"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateGateway with error: Operation response processing error`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Construct an instance of the AuthenticationKeyIdentityKeyProtectAuthenticationKeyIdentity model
				authenticationKeyIdentityModel := new(directlinkv1.AuthenticationKeyIdentityKeyProtectAuthenticationKeyIdentity)
				authenticationKeyIdentityModel.Crn = core.StringPtr("crn:v1:bluemix:public:kms:us-south:a/4111d05f36894e3cb9b46a43556d9000:abc111b8-37aa-4034-9def-f2607c87aaaa:key:bbb222bc-430a-4de9-9aad-84e5bb022222")

				// Construct an instance of the GatewayBfdPatchTemplate model
				gatewayBfdPatchTemplateModel := new(directlinkv1.GatewayBfdPatchTemplate)
				gatewayBfdPatchTemplateModel.Interval = core.Int64Ptr(int64(2000))
				gatewayBfdPatchTemplateModel.Multiplier = core.Int64Ptr(int64(10))

				// Construct an instance of the GatewayPatchTemplate model
				gatewayPatchTemplateModel := new(directlinkv1.GatewayPatchTemplate)
				gatewayPatchTemplateModel.AuthenticationKey = authenticationKeyIdentityModel
				gatewayPatchTemplateModel.BfdConfig = gatewayBfdPatchTemplateModel
				gatewayPatchTemplateModel.BgpAsn = core.Int64Ptr(int64(64999))
				gatewayPatchTemplateModel.BgpCerCidr = core.StringPtr("169.254.0.10/30")
				gatewayPatchTemplateModel.BgpIbmCidr = core.StringPtr("169.254.0.9/30")
				gatewayPatchTemplateModel.ConnectionMode = core.StringPtr("transit")
				gatewayPatchTemplateModel.DefaultExportRouteFilter = core.StringPtr("permit")
				gatewayPatchTemplateModel.DefaultImportRouteFilter = core.StringPtr("permit")
				gatewayPatchTemplateModel.Global = core.BoolPtr(true)
				gatewayPatchTemplateModel.LoaRejectReason = core.StringPtr("The port mentioned was incorrect")
				gatewayPatchTemplateModel.Metered = core.BoolPtr(false)
				gatewayPatchTemplateModel.Name = core.StringPtr("testGateway")
				gatewayPatchTemplateModel.OperationalStatus = core.StringPtr("loa_accepted")
				gatewayPatchTemplateModel.PatchPanelCompletionNotice = core.StringPtr("patch panel configuration details")
				gatewayPatchTemplateModel.SpeedMbps = core.Int64Ptr(int64(1000))
				gatewayPatchTemplateModel.Vlan = core.Int64Ptr(int64(10))
				gatewayPatchTemplateModelAsPatch, asPatchErr := gatewayPatchTemplateModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateGatewayOptions model
				updateGatewayOptionsModel := new(directlinkv1.UpdateGatewayOptions)
				updateGatewayOptionsModel.ID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				updateGatewayOptionsModel.GatewayPatchTemplatePatch = gatewayPatchTemplateModelAsPatch
				updateGatewayOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := directLinkService.UpdateGateway(updateGatewayOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				directLinkService.EnableRetries(0, 0)
				result, response, operationErr = directLinkService.UpdateGateway(updateGatewayOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateGateway(updateGatewayOptions *UpdateGatewayOptions)`, func() {
		version := "testString"
		updateGatewayPath := "/gateways/0a06fb9b-820f-4c44-8a31-77f1f0806d28"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateGatewayPath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"as_prepends": [{"created_at": "2019-01-01T12:00:00.000Z", "id": "ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4", "length": 4, "policy": "import", "prefix": "172.17.0.0/16", "specific_prefixes": ["192.168.3.0/24"], "updated_at": "2019-01-01T12:00:00.000Z"}], "authentication_key": {"crn": "crn:v1:bluemix:public:kms:us-south:a/4111d05f36894e3cb9b46a43556d9000:abc111b8-37aa-4034-9def-f2607c87aaaa:key:bbb222bc-430a-4de9-9aad-84e5bb022222"}, "bfd_config": {"bfd_status": "up", "bfd_status_updated_at": "2020-08-20T06:58:41.909Z", "interval": 2000, "multiplier": 10}, "bgp_asn": 64999, "bgp_base_cidr": "BgpBaseCidr", "bgp_cer_cidr": "10.254.30.78/30", "bgp_ibm_asn": 13884, "bgp_ibm_cidr": "10.254.30.77/30", "bgp_status": "active", "bgp_status_updated_at": "2020-08-20T06:58:41.909Z", "carrier_name": "myCarrierName", "change_request": {"type": "create_gateway"}, "completion_notice_reject_reason": "The completion notice file was blank", "connection_mode": "transit", "created_at": "2019-01-01T12:00:00.000Z", "crn": "crn:v1:bluemix:public:directlink:dal03:a/4111d05f36894e3cb9b46a43556d9000::dedicated:ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4", "cross_account": false, "cross_connect_router": "xcr01.dal03", "customer_name": "newCustomerName", "default_export_route_filter": "permit", "default_import_route_filter": "permit", "global": true, "id": "ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4", "link_status": "up", "link_status_updated_at": "2020-08-20T06:58:41.909Z", "location_display_name": "Dallas 03", "location_name": "dal03", "macsec": {"active": true, "security_policy": "must_secure", "status": "secured", "status_reasons": [{"code": "macsec_cak_failed", "message": "The authentication_key failed configuration.", "more_info": "https://cloud.ibm.com/docs/dl/TODO_ADD_DOCS_LINK"}]}, "macsec_capability": "non_macsec", "metered": false, "name": "myGateway", "operational_status": "awaiting_completion_notice", "operational_status_reasons": [{"code": "authentication_key_failed", "message": "The authentication_key failed configuration.", "more_info": "https://cloud.ibm.com/docs/dl/TODO_ADD_DOCS_LINK"}], "patch_panel_completion_notice": "patch panel configuration details", "port": {"id": "54321b1a-fee4-41c7-9e11-9cd99e000aaa"}, "provider_api_managed": false, "resource_group": {"id": "56969d6043e9465c883cb9f7363e78e8"}, "speed_mbps": 1000, "type": "dedicated", "vlan": 10}`)
				}))
			})
			It(`Invoke UpdateGateway successfully with retries`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())
				directLinkService.EnableRetries(0, 0)

				// Construct an instance of the AuthenticationKeyIdentityKeyProtectAuthenticationKeyIdentity model
				authenticationKeyIdentityModel := new(directlinkv1.AuthenticationKeyIdentityKeyProtectAuthenticationKeyIdentity)
				authenticationKeyIdentityModel.Crn = core.StringPtr("crn:v1:bluemix:public:kms:us-south:a/4111d05f36894e3cb9b46a43556d9000:abc111b8-37aa-4034-9def-f2607c87aaaa:key:bbb222bc-430a-4de9-9aad-84e5bb022222")

				// Construct an instance of the GatewayBfdPatchTemplate model
				gatewayBfdPatchTemplateModel := new(directlinkv1.GatewayBfdPatchTemplate)
				gatewayBfdPatchTemplateModel.Interval = core.Int64Ptr(int64(2000))
				gatewayBfdPatchTemplateModel.Multiplier = core.Int64Ptr(int64(10))

				// Construct an instance of the GatewayPatchTemplate model
				gatewayPatchTemplateModel := new(directlinkv1.GatewayPatchTemplate)
				gatewayPatchTemplateModel.AuthenticationKey = authenticationKeyIdentityModel
				gatewayPatchTemplateModel.BfdConfig = gatewayBfdPatchTemplateModel
				gatewayPatchTemplateModel.BgpAsn = core.Int64Ptr(int64(64999))
				gatewayPatchTemplateModel.BgpCerCidr = core.StringPtr("169.254.0.10/30")
				gatewayPatchTemplateModel.BgpIbmCidr = core.StringPtr("169.254.0.9/30")
				gatewayPatchTemplateModel.ConnectionMode = core.StringPtr("transit")
				gatewayPatchTemplateModel.DefaultExportRouteFilter = core.StringPtr("permit")
				gatewayPatchTemplateModel.DefaultImportRouteFilter = core.StringPtr("permit")
				gatewayPatchTemplateModel.Global = core.BoolPtr(true)
				gatewayPatchTemplateModel.LoaRejectReason = core.StringPtr("The port mentioned was incorrect")
				gatewayPatchTemplateModel.Metered = core.BoolPtr(false)
				gatewayPatchTemplateModel.Name = core.StringPtr("testGateway")
				gatewayPatchTemplateModel.OperationalStatus = core.StringPtr("loa_accepted")
				gatewayPatchTemplateModel.PatchPanelCompletionNotice = core.StringPtr("patch panel configuration details")
				gatewayPatchTemplateModel.SpeedMbps = core.Int64Ptr(int64(1000))
				gatewayPatchTemplateModel.Vlan = core.Int64Ptr(int64(10))
				gatewayPatchTemplateModelAsPatch, asPatchErr := gatewayPatchTemplateModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateGatewayOptions model
				updateGatewayOptionsModel := new(directlinkv1.UpdateGatewayOptions)
				updateGatewayOptionsModel.ID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				updateGatewayOptionsModel.GatewayPatchTemplatePatch = gatewayPatchTemplateModelAsPatch
				updateGatewayOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := directLinkService.UpdateGatewayWithContext(ctx, updateGatewayOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				directLinkService.DisableRetries()
				result, response, operationErr := directLinkService.UpdateGateway(updateGatewayOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = directLinkService.UpdateGatewayWithContext(ctx, updateGatewayOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateGatewayPath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"as_prepends": [{"created_at": "2019-01-01T12:00:00.000Z", "id": "ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4", "length": 4, "policy": "import", "prefix": "172.17.0.0/16", "specific_prefixes": ["192.168.3.0/24"], "updated_at": "2019-01-01T12:00:00.000Z"}], "authentication_key": {"crn": "crn:v1:bluemix:public:kms:us-south:a/4111d05f36894e3cb9b46a43556d9000:abc111b8-37aa-4034-9def-f2607c87aaaa:key:bbb222bc-430a-4de9-9aad-84e5bb022222"}, "bfd_config": {"bfd_status": "up", "bfd_status_updated_at": "2020-08-20T06:58:41.909Z", "interval": 2000, "multiplier": 10}, "bgp_asn": 64999, "bgp_base_cidr": "BgpBaseCidr", "bgp_cer_cidr": "10.254.30.78/30", "bgp_ibm_asn": 13884, "bgp_ibm_cidr": "10.254.30.77/30", "bgp_status": "active", "bgp_status_updated_at": "2020-08-20T06:58:41.909Z", "carrier_name": "myCarrierName", "change_request": {"type": "create_gateway"}, "completion_notice_reject_reason": "The completion notice file was blank", "connection_mode": "transit", "created_at": "2019-01-01T12:00:00.000Z", "crn": "crn:v1:bluemix:public:directlink:dal03:a/4111d05f36894e3cb9b46a43556d9000::dedicated:ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4", "cross_account": false, "cross_connect_router": "xcr01.dal03", "customer_name": "newCustomerName", "default_export_route_filter": "permit", "default_import_route_filter": "permit", "global": true, "id": "ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4", "link_status": "up", "link_status_updated_at": "2020-08-20T06:58:41.909Z", "location_display_name": "Dallas 03", "location_name": "dal03", "macsec": {"active": true, "security_policy": "must_secure", "status": "secured", "status_reasons": [{"code": "macsec_cak_failed", "message": "The authentication_key failed configuration.", "more_info": "https://cloud.ibm.com/docs/dl/TODO_ADD_DOCS_LINK"}]}, "macsec_capability": "non_macsec", "metered": false, "name": "myGateway", "operational_status": "awaiting_completion_notice", "operational_status_reasons": [{"code": "authentication_key_failed", "message": "The authentication_key failed configuration.", "more_info": "https://cloud.ibm.com/docs/dl/TODO_ADD_DOCS_LINK"}], "patch_panel_completion_notice": "patch panel configuration details", "port": {"id": "54321b1a-fee4-41c7-9e11-9cd99e000aaa"}, "provider_api_managed": false, "resource_group": {"id": "56969d6043e9465c883cb9f7363e78e8"}, "speed_mbps": 1000, "type": "dedicated", "vlan": 10}`)
				}))
			})
			It(`Invoke UpdateGateway successfully`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := directLinkService.UpdateGateway(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the AuthenticationKeyIdentityKeyProtectAuthenticationKeyIdentity model
				authenticationKeyIdentityModel := new(directlinkv1.AuthenticationKeyIdentityKeyProtectAuthenticationKeyIdentity)
				authenticationKeyIdentityModel.Crn = core.StringPtr("crn:v1:bluemix:public:kms:us-south:a/4111d05f36894e3cb9b46a43556d9000:abc111b8-37aa-4034-9def-f2607c87aaaa:key:bbb222bc-430a-4de9-9aad-84e5bb022222")

				// Construct an instance of the GatewayBfdPatchTemplate model
				gatewayBfdPatchTemplateModel := new(directlinkv1.GatewayBfdPatchTemplate)
				gatewayBfdPatchTemplateModel.Interval = core.Int64Ptr(int64(2000))
				gatewayBfdPatchTemplateModel.Multiplier = core.Int64Ptr(int64(10))

				// Construct an instance of the GatewayPatchTemplate model
				gatewayPatchTemplateModel := new(directlinkv1.GatewayPatchTemplate)
				gatewayPatchTemplateModel.AuthenticationKey = authenticationKeyIdentityModel
				gatewayPatchTemplateModel.BfdConfig = gatewayBfdPatchTemplateModel
				gatewayPatchTemplateModel.BgpAsn = core.Int64Ptr(int64(64999))
				gatewayPatchTemplateModel.BgpCerCidr = core.StringPtr("169.254.0.10/30")
				gatewayPatchTemplateModel.BgpIbmCidr = core.StringPtr("169.254.0.9/30")
				gatewayPatchTemplateModel.ConnectionMode = core.StringPtr("transit")
				gatewayPatchTemplateModel.DefaultExportRouteFilter = core.StringPtr("permit")
				gatewayPatchTemplateModel.DefaultImportRouteFilter = core.StringPtr("permit")
				gatewayPatchTemplateModel.Global = core.BoolPtr(true)
				gatewayPatchTemplateModel.LoaRejectReason = core.StringPtr("The port mentioned was incorrect")
				gatewayPatchTemplateModel.Metered = core.BoolPtr(false)
				gatewayPatchTemplateModel.Name = core.StringPtr("testGateway")
				gatewayPatchTemplateModel.OperationalStatus = core.StringPtr("loa_accepted")
				gatewayPatchTemplateModel.PatchPanelCompletionNotice = core.StringPtr("patch panel configuration details")
				gatewayPatchTemplateModel.SpeedMbps = core.Int64Ptr(int64(1000))
				gatewayPatchTemplateModel.Vlan = core.Int64Ptr(int64(10))
				gatewayPatchTemplateModelAsPatch, asPatchErr := gatewayPatchTemplateModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateGatewayOptions model
				updateGatewayOptionsModel := new(directlinkv1.UpdateGatewayOptions)
				updateGatewayOptionsModel.ID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				updateGatewayOptionsModel.GatewayPatchTemplatePatch = gatewayPatchTemplateModelAsPatch
				updateGatewayOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = directLinkService.UpdateGateway(updateGatewayOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateGateway with error: Operation validation and request error`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Construct an instance of the AuthenticationKeyIdentityKeyProtectAuthenticationKeyIdentity model
				authenticationKeyIdentityModel := new(directlinkv1.AuthenticationKeyIdentityKeyProtectAuthenticationKeyIdentity)
				authenticationKeyIdentityModel.Crn = core.StringPtr("crn:v1:bluemix:public:kms:us-south:a/4111d05f36894e3cb9b46a43556d9000:abc111b8-37aa-4034-9def-f2607c87aaaa:key:bbb222bc-430a-4de9-9aad-84e5bb022222")

				// Construct an instance of the GatewayBfdPatchTemplate model
				gatewayBfdPatchTemplateModel := new(directlinkv1.GatewayBfdPatchTemplate)
				gatewayBfdPatchTemplateModel.Interval = core.Int64Ptr(int64(2000))
				gatewayBfdPatchTemplateModel.Multiplier = core.Int64Ptr(int64(10))

				// Construct an instance of the GatewayPatchTemplate model
				gatewayPatchTemplateModel := new(directlinkv1.GatewayPatchTemplate)
				gatewayPatchTemplateModel.AuthenticationKey = authenticationKeyIdentityModel
				gatewayPatchTemplateModel.BfdConfig = gatewayBfdPatchTemplateModel
				gatewayPatchTemplateModel.BgpAsn = core.Int64Ptr(int64(64999))
				gatewayPatchTemplateModel.BgpCerCidr = core.StringPtr("169.254.0.10/30")
				gatewayPatchTemplateModel.BgpIbmCidr = core.StringPtr("169.254.0.9/30")
				gatewayPatchTemplateModel.ConnectionMode = core.StringPtr("transit")
				gatewayPatchTemplateModel.DefaultExportRouteFilter = core.StringPtr("permit")
				gatewayPatchTemplateModel.DefaultImportRouteFilter = core.StringPtr("permit")
				gatewayPatchTemplateModel.Global = core.BoolPtr(true)
				gatewayPatchTemplateModel.LoaRejectReason = core.StringPtr("The port mentioned was incorrect")
				gatewayPatchTemplateModel.Metered = core.BoolPtr(false)
				gatewayPatchTemplateModel.Name = core.StringPtr("testGateway")
				gatewayPatchTemplateModel.OperationalStatus = core.StringPtr("loa_accepted")
				gatewayPatchTemplateModel.PatchPanelCompletionNotice = core.StringPtr("patch panel configuration details")
				gatewayPatchTemplateModel.SpeedMbps = core.Int64Ptr(int64(1000))
				gatewayPatchTemplateModel.Vlan = core.Int64Ptr(int64(10))
				gatewayPatchTemplateModelAsPatch, asPatchErr := gatewayPatchTemplateModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateGatewayOptions model
				updateGatewayOptionsModel := new(directlinkv1.UpdateGatewayOptions)
				updateGatewayOptionsModel.ID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				updateGatewayOptionsModel.GatewayPatchTemplatePatch = gatewayPatchTemplateModelAsPatch
				updateGatewayOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := directLinkService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := directLinkService.UpdateGateway(updateGatewayOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateGatewayOptions model with no property values
				updateGatewayOptionsModelNew := new(directlinkv1.UpdateGatewayOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = directLinkService.UpdateGateway(updateGatewayOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke UpdateGateway successfully`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Construct an instance of the AuthenticationKeyIdentityKeyProtectAuthenticationKeyIdentity model
				authenticationKeyIdentityModel := new(directlinkv1.AuthenticationKeyIdentityKeyProtectAuthenticationKeyIdentity)
				authenticationKeyIdentityModel.Crn = core.StringPtr("crn:v1:bluemix:public:kms:us-south:a/4111d05f36894e3cb9b46a43556d9000:abc111b8-37aa-4034-9def-f2607c87aaaa:key:bbb222bc-430a-4de9-9aad-84e5bb022222")

				// Construct an instance of the GatewayBfdPatchTemplate model
				gatewayBfdPatchTemplateModel := new(directlinkv1.GatewayBfdPatchTemplate)
				gatewayBfdPatchTemplateModel.Interval = core.Int64Ptr(int64(2000))
				gatewayBfdPatchTemplateModel.Multiplier = core.Int64Ptr(int64(10))

				// Construct an instance of the GatewayPatchTemplate model
				gatewayPatchTemplateModel := new(directlinkv1.GatewayPatchTemplate)
				gatewayPatchTemplateModel.AuthenticationKey = authenticationKeyIdentityModel
				gatewayPatchTemplateModel.BfdConfig = gatewayBfdPatchTemplateModel
				gatewayPatchTemplateModel.BgpAsn = core.Int64Ptr(int64(64999))
				gatewayPatchTemplateModel.BgpCerCidr = core.StringPtr("169.254.0.10/30")
				gatewayPatchTemplateModel.BgpIbmCidr = core.StringPtr("169.254.0.9/30")
				gatewayPatchTemplateModel.ConnectionMode = core.StringPtr("transit")
				gatewayPatchTemplateModel.DefaultExportRouteFilter = core.StringPtr("permit")
				gatewayPatchTemplateModel.DefaultImportRouteFilter = core.StringPtr("permit")
				gatewayPatchTemplateModel.Global = core.BoolPtr(true)
				gatewayPatchTemplateModel.LoaRejectReason = core.StringPtr("The port mentioned was incorrect")
				gatewayPatchTemplateModel.Metered = core.BoolPtr(false)
				gatewayPatchTemplateModel.Name = core.StringPtr("testGateway")
				gatewayPatchTemplateModel.OperationalStatus = core.StringPtr("loa_accepted")
				gatewayPatchTemplateModel.PatchPanelCompletionNotice = core.StringPtr("patch panel configuration details")
				gatewayPatchTemplateModel.SpeedMbps = core.Int64Ptr(int64(1000))
				gatewayPatchTemplateModel.Vlan = core.Int64Ptr(int64(10))
				gatewayPatchTemplateModelAsPatch, asPatchErr := gatewayPatchTemplateModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateGatewayOptions model
				updateGatewayOptionsModel := new(directlinkv1.UpdateGatewayOptions)
				updateGatewayOptionsModel.ID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				updateGatewayOptionsModel.GatewayPatchTemplatePatch = gatewayPatchTemplateModelAsPatch
				updateGatewayOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := directLinkService.UpdateGateway(updateGatewayOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateGatewayAction(createGatewayActionOptions *CreateGatewayActionOptions) - Operation response error`, func() {
		version := "testString"
		createGatewayActionPath := "/gateways/0a06fb9b-820f-4c44-8a31-77f1f0806d28/actions"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createGatewayActionPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateGatewayAction with error: Operation response processing error`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Construct an instance of the AsPrependTemplate model
				asPrependTemplateModel := new(directlinkv1.AsPrependTemplate)
				asPrependTemplateModel.Length = core.Int64Ptr(int64(4))
				asPrependTemplateModel.Policy = core.StringPtr("import")
				asPrependTemplateModel.Prefix = core.StringPtr("172.17.0.0/16")
				asPrependTemplateModel.SpecificPrefixes = []string{"192.168.3.0/24"}

				// Construct an instance of the AuthenticationKeyIdentityKeyProtectAuthenticationKeyIdentity model
				authenticationKeyIdentityModel := new(directlinkv1.AuthenticationKeyIdentityKeyProtectAuthenticationKeyIdentity)
				authenticationKeyIdentityModel.Crn = core.StringPtr("crn:v1:bluemix:public:kms:us-south:a/4111d05f36894e3cb9b46a43556d9000:abc111b8-37aa-4034-9def-f2607c87aaaa:key:bbb222bc-430a-4de9-9aad-84e5bb022222")

				// Construct an instance of the GatewayBfdConfigActionTemplate model
				gatewayBfdConfigActionTemplateModel := new(directlinkv1.GatewayBfdConfigActionTemplate)
				gatewayBfdConfigActionTemplateModel.Interval = core.Int64Ptr(int64(2000))
				gatewayBfdConfigActionTemplateModel.Multiplier = core.Int64Ptr(int64(10))

				// Construct an instance of the GatewayTemplateRouteFilter model
				gatewayTemplateRouteFilterModel := new(directlinkv1.GatewayTemplateRouteFilter)
				gatewayTemplateRouteFilterModel.Action = core.StringPtr("permit")
				gatewayTemplateRouteFilterModel.Ge = core.Int64Ptr(int64(25))
				gatewayTemplateRouteFilterModel.Le = core.Int64Ptr(int64(30))
				gatewayTemplateRouteFilterModel.Prefix = core.StringPtr("192.168.100.0/24")

				// Construct an instance of the ResourceGroupIdentity model
				resourceGroupIdentityModel := new(directlinkv1.ResourceGroupIdentity)
				resourceGroupIdentityModel.ID = core.StringPtr("56969d6043e9465c883cb9f7363e78e8")

				// Construct an instance of the GatewayActionTemplateUpdatesItemGatewayClientSpeedUpdate model
				gatewayActionTemplateUpdatesItemModel := new(directlinkv1.GatewayActionTemplateUpdatesItemGatewayClientSpeedUpdate)
				gatewayActionTemplateUpdatesItemModel.SpeedMbps = core.Int64Ptr(int64(500))

				// Construct an instance of the CreateGatewayActionOptions model
				createGatewayActionOptionsModel := new(directlinkv1.CreateGatewayActionOptions)
				createGatewayActionOptionsModel.ID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				createGatewayActionOptionsModel.Action = core.StringPtr("create_gateway_approve")
				createGatewayActionOptionsModel.AsPrepends = []directlinkv1.AsPrependTemplate{*asPrependTemplateModel}
				createGatewayActionOptionsModel.AuthenticationKey = authenticationKeyIdentityModel
				createGatewayActionOptionsModel.BfdConfig = gatewayBfdConfigActionTemplateModel
				createGatewayActionOptionsModel.ConnectionMode = core.StringPtr("transit")
				createGatewayActionOptionsModel.DefaultExportRouteFilter = core.StringPtr("permit")
				createGatewayActionOptionsModel.DefaultImportRouteFilter = core.StringPtr("permit")
				createGatewayActionOptionsModel.ExportRouteFilters = []directlinkv1.GatewayTemplateRouteFilter{*gatewayTemplateRouteFilterModel}
				createGatewayActionOptionsModel.Global = core.BoolPtr(true)
				createGatewayActionOptionsModel.ImportRouteFilters = []directlinkv1.GatewayTemplateRouteFilter{*gatewayTemplateRouteFilterModel}
				createGatewayActionOptionsModel.Metered = core.BoolPtr(false)
				createGatewayActionOptionsModel.ResourceGroup = resourceGroupIdentityModel
				createGatewayActionOptionsModel.Updates = []directlinkv1.GatewayActionTemplateUpdatesItemIntf{gatewayActionTemplateUpdatesItemModel}
				createGatewayActionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := directLinkService.CreateGatewayAction(createGatewayActionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				directLinkService.EnableRetries(0, 0)
				result, response, operationErr = directLinkService.CreateGatewayAction(createGatewayActionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateGatewayAction(createGatewayActionOptions *CreateGatewayActionOptions)`, func() {
		version := "testString"
		createGatewayActionPath := "/gateways/0a06fb9b-820f-4c44-8a31-77f1f0806d28/actions"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createGatewayActionPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"as_prepends": [{"created_at": "2019-01-01T12:00:00.000Z", "id": "ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4", "length": 4, "policy": "import", "prefix": "172.17.0.0/16", "specific_prefixes": ["192.168.3.0/24"], "updated_at": "2019-01-01T12:00:00.000Z"}], "authentication_key": {"crn": "crn:v1:bluemix:public:kms:us-south:a/4111d05f36894e3cb9b46a43556d9000:abc111b8-37aa-4034-9def-f2607c87aaaa:key:bbb222bc-430a-4de9-9aad-84e5bb022222"}, "bfd_config": {"bfd_status": "up", "bfd_status_updated_at": "2020-08-20T06:58:41.909Z", "interval": 2000, "multiplier": 10}, "bgp_asn": 64999, "bgp_base_cidr": "BgpBaseCidr", "bgp_cer_cidr": "10.254.30.78/30", "bgp_ibm_asn": 13884, "bgp_ibm_cidr": "10.254.30.77/30", "bgp_status": "active", "bgp_status_updated_at": "2020-08-20T06:58:41.909Z", "carrier_name": "myCarrierName", "change_request": {"type": "create_gateway"}, "completion_notice_reject_reason": "The completion notice file was blank", "connection_mode": "transit", "created_at": "2019-01-01T12:00:00.000Z", "crn": "crn:v1:bluemix:public:directlink:dal03:a/4111d05f36894e3cb9b46a43556d9000::dedicated:ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4", "cross_account": false, "cross_connect_router": "xcr01.dal03", "customer_name": "newCustomerName", "default_export_route_filter": "permit", "default_import_route_filter": "permit", "global": true, "id": "ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4", "link_status": "up", "link_status_updated_at": "2020-08-20T06:58:41.909Z", "location_display_name": "Dallas 03", "location_name": "dal03", "macsec": {"active": true, "security_policy": "must_secure", "status": "secured", "status_reasons": [{"code": "macsec_cak_failed", "message": "The authentication_key failed configuration.", "more_info": "https://cloud.ibm.com/docs/dl/TODO_ADD_DOCS_LINK"}]}, "macsec_capability": "non_macsec", "metered": false, "name": "myGateway", "operational_status": "awaiting_completion_notice", "operational_status_reasons": [{"code": "authentication_key_failed", "message": "The authentication_key failed configuration.", "more_info": "https://cloud.ibm.com/docs/dl/TODO_ADD_DOCS_LINK"}], "patch_panel_completion_notice": "patch panel configuration details", "port": {"id": "54321b1a-fee4-41c7-9e11-9cd99e000aaa"}, "provider_api_managed": false, "resource_group": {"id": "56969d6043e9465c883cb9f7363e78e8"}, "speed_mbps": 1000, "type": "dedicated", "vlan": 10}`)
				}))
			})
			It(`Invoke CreateGatewayAction successfully with retries`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())
				directLinkService.EnableRetries(0, 0)

				// Construct an instance of the AsPrependTemplate model
				asPrependTemplateModel := new(directlinkv1.AsPrependTemplate)
				asPrependTemplateModel.Length = core.Int64Ptr(int64(4))
				asPrependTemplateModel.Policy = core.StringPtr("import")
				asPrependTemplateModel.Prefix = core.StringPtr("172.17.0.0/16")
				asPrependTemplateModel.SpecificPrefixes = []string{"192.168.3.0/24"}

				// Construct an instance of the AuthenticationKeyIdentityKeyProtectAuthenticationKeyIdentity model
				authenticationKeyIdentityModel := new(directlinkv1.AuthenticationKeyIdentityKeyProtectAuthenticationKeyIdentity)
				authenticationKeyIdentityModel.Crn = core.StringPtr("crn:v1:bluemix:public:kms:us-south:a/4111d05f36894e3cb9b46a43556d9000:abc111b8-37aa-4034-9def-f2607c87aaaa:key:bbb222bc-430a-4de9-9aad-84e5bb022222")

				// Construct an instance of the GatewayBfdConfigActionTemplate model
				gatewayBfdConfigActionTemplateModel := new(directlinkv1.GatewayBfdConfigActionTemplate)
				gatewayBfdConfigActionTemplateModel.Interval = core.Int64Ptr(int64(2000))
				gatewayBfdConfigActionTemplateModel.Multiplier = core.Int64Ptr(int64(10))

				// Construct an instance of the GatewayTemplateRouteFilter model
				gatewayTemplateRouteFilterModel := new(directlinkv1.GatewayTemplateRouteFilter)
				gatewayTemplateRouteFilterModel.Action = core.StringPtr("permit")
				gatewayTemplateRouteFilterModel.Ge = core.Int64Ptr(int64(25))
				gatewayTemplateRouteFilterModel.Le = core.Int64Ptr(int64(30))
				gatewayTemplateRouteFilterModel.Prefix = core.StringPtr("192.168.100.0/24")

				// Construct an instance of the ResourceGroupIdentity model
				resourceGroupIdentityModel := new(directlinkv1.ResourceGroupIdentity)
				resourceGroupIdentityModel.ID = core.StringPtr("56969d6043e9465c883cb9f7363e78e8")

				// Construct an instance of the GatewayActionTemplateUpdatesItemGatewayClientSpeedUpdate model
				gatewayActionTemplateUpdatesItemModel := new(directlinkv1.GatewayActionTemplateUpdatesItemGatewayClientSpeedUpdate)
				gatewayActionTemplateUpdatesItemModel.SpeedMbps = core.Int64Ptr(int64(500))

				// Construct an instance of the CreateGatewayActionOptions model
				createGatewayActionOptionsModel := new(directlinkv1.CreateGatewayActionOptions)
				createGatewayActionOptionsModel.ID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				createGatewayActionOptionsModel.Action = core.StringPtr("create_gateway_approve")
				createGatewayActionOptionsModel.AsPrepends = []directlinkv1.AsPrependTemplate{*asPrependTemplateModel}
				createGatewayActionOptionsModel.AuthenticationKey = authenticationKeyIdentityModel
				createGatewayActionOptionsModel.BfdConfig = gatewayBfdConfigActionTemplateModel
				createGatewayActionOptionsModel.ConnectionMode = core.StringPtr("transit")
				createGatewayActionOptionsModel.DefaultExportRouteFilter = core.StringPtr("permit")
				createGatewayActionOptionsModel.DefaultImportRouteFilter = core.StringPtr("permit")
				createGatewayActionOptionsModel.ExportRouteFilters = []directlinkv1.GatewayTemplateRouteFilter{*gatewayTemplateRouteFilterModel}
				createGatewayActionOptionsModel.Global = core.BoolPtr(true)
				createGatewayActionOptionsModel.ImportRouteFilters = []directlinkv1.GatewayTemplateRouteFilter{*gatewayTemplateRouteFilterModel}
				createGatewayActionOptionsModel.Metered = core.BoolPtr(false)
				createGatewayActionOptionsModel.ResourceGroup = resourceGroupIdentityModel
				createGatewayActionOptionsModel.Updates = []directlinkv1.GatewayActionTemplateUpdatesItemIntf{gatewayActionTemplateUpdatesItemModel}
				createGatewayActionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := directLinkService.CreateGatewayActionWithContext(ctx, createGatewayActionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				directLinkService.DisableRetries()
				result, response, operationErr := directLinkService.CreateGatewayAction(createGatewayActionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = directLinkService.CreateGatewayActionWithContext(ctx, createGatewayActionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createGatewayActionPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"as_prepends": [{"created_at": "2019-01-01T12:00:00.000Z", "id": "ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4", "length": 4, "policy": "import", "prefix": "172.17.0.0/16", "specific_prefixes": ["192.168.3.0/24"], "updated_at": "2019-01-01T12:00:00.000Z"}], "authentication_key": {"crn": "crn:v1:bluemix:public:kms:us-south:a/4111d05f36894e3cb9b46a43556d9000:abc111b8-37aa-4034-9def-f2607c87aaaa:key:bbb222bc-430a-4de9-9aad-84e5bb022222"}, "bfd_config": {"bfd_status": "up", "bfd_status_updated_at": "2020-08-20T06:58:41.909Z", "interval": 2000, "multiplier": 10}, "bgp_asn": 64999, "bgp_base_cidr": "BgpBaseCidr", "bgp_cer_cidr": "10.254.30.78/30", "bgp_ibm_asn": 13884, "bgp_ibm_cidr": "10.254.30.77/30", "bgp_status": "active", "bgp_status_updated_at": "2020-08-20T06:58:41.909Z", "carrier_name": "myCarrierName", "change_request": {"type": "create_gateway"}, "completion_notice_reject_reason": "The completion notice file was blank", "connection_mode": "transit", "created_at": "2019-01-01T12:00:00.000Z", "crn": "crn:v1:bluemix:public:directlink:dal03:a/4111d05f36894e3cb9b46a43556d9000::dedicated:ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4", "cross_account": false, "cross_connect_router": "xcr01.dal03", "customer_name": "newCustomerName", "default_export_route_filter": "permit", "default_import_route_filter": "permit", "global": true, "id": "ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4", "link_status": "up", "link_status_updated_at": "2020-08-20T06:58:41.909Z", "location_display_name": "Dallas 03", "location_name": "dal03", "macsec": {"active": true, "security_policy": "must_secure", "status": "secured", "status_reasons": [{"code": "macsec_cak_failed", "message": "The authentication_key failed configuration.", "more_info": "https://cloud.ibm.com/docs/dl/TODO_ADD_DOCS_LINK"}]}, "macsec_capability": "non_macsec", "metered": false, "name": "myGateway", "operational_status": "awaiting_completion_notice", "operational_status_reasons": [{"code": "authentication_key_failed", "message": "The authentication_key failed configuration.", "more_info": "https://cloud.ibm.com/docs/dl/TODO_ADD_DOCS_LINK"}], "patch_panel_completion_notice": "patch panel configuration details", "port": {"id": "54321b1a-fee4-41c7-9e11-9cd99e000aaa"}, "provider_api_managed": false, "resource_group": {"id": "56969d6043e9465c883cb9f7363e78e8"}, "speed_mbps": 1000, "type": "dedicated", "vlan": 10}`)
				}))
			})
			It(`Invoke CreateGatewayAction successfully`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := directLinkService.CreateGatewayAction(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the AsPrependTemplate model
				asPrependTemplateModel := new(directlinkv1.AsPrependTemplate)
				asPrependTemplateModel.Length = core.Int64Ptr(int64(4))
				asPrependTemplateModel.Policy = core.StringPtr("import")
				asPrependTemplateModel.Prefix = core.StringPtr("172.17.0.0/16")
				asPrependTemplateModel.SpecificPrefixes = []string{"192.168.3.0/24"}

				// Construct an instance of the AuthenticationKeyIdentityKeyProtectAuthenticationKeyIdentity model
				authenticationKeyIdentityModel := new(directlinkv1.AuthenticationKeyIdentityKeyProtectAuthenticationKeyIdentity)
				authenticationKeyIdentityModel.Crn = core.StringPtr("crn:v1:bluemix:public:kms:us-south:a/4111d05f36894e3cb9b46a43556d9000:abc111b8-37aa-4034-9def-f2607c87aaaa:key:bbb222bc-430a-4de9-9aad-84e5bb022222")

				// Construct an instance of the GatewayBfdConfigActionTemplate model
				gatewayBfdConfigActionTemplateModel := new(directlinkv1.GatewayBfdConfigActionTemplate)
				gatewayBfdConfigActionTemplateModel.Interval = core.Int64Ptr(int64(2000))
				gatewayBfdConfigActionTemplateModel.Multiplier = core.Int64Ptr(int64(10))

				// Construct an instance of the GatewayTemplateRouteFilter model
				gatewayTemplateRouteFilterModel := new(directlinkv1.GatewayTemplateRouteFilter)
				gatewayTemplateRouteFilterModel.Action = core.StringPtr("permit")
				gatewayTemplateRouteFilterModel.Ge = core.Int64Ptr(int64(25))
				gatewayTemplateRouteFilterModel.Le = core.Int64Ptr(int64(30))
				gatewayTemplateRouteFilterModel.Prefix = core.StringPtr("192.168.100.0/24")

				// Construct an instance of the ResourceGroupIdentity model
				resourceGroupIdentityModel := new(directlinkv1.ResourceGroupIdentity)
				resourceGroupIdentityModel.ID = core.StringPtr("56969d6043e9465c883cb9f7363e78e8")

				// Construct an instance of the GatewayActionTemplateUpdatesItemGatewayClientSpeedUpdate model
				gatewayActionTemplateUpdatesItemModel := new(directlinkv1.GatewayActionTemplateUpdatesItemGatewayClientSpeedUpdate)
				gatewayActionTemplateUpdatesItemModel.SpeedMbps = core.Int64Ptr(int64(500))

				// Construct an instance of the CreateGatewayActionOptions model
				createGatewayActionOptionsModel := new(directlinkv1.CreateGatewayActionOptions)
				createGatewayActionOptionsModel.ID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				createGatewayActionOptionsModel.Action = core.StringPtr("create_gateway_approve")
				createGatewayActionOptionsModel.AsPrepends = []directlinkv1.AsPrependTemplate{*asPrependTemplateModel}
				createGatewayActionOptionsModel.AuthenticationKey = authenticationKeyIdentityModel
				createGatewayActionOptionsModel.BfdConfig = gatewayBfdConfigActionTemplateModel
				createGatewayActionOptionsModel.ConnectionMode = core.StringPtr("transit")
				createGatewayActionOptionsModel.DefaultExportRouteFilter = core.StringPtr("permit")
				createGatewayActionOptionsModel.DefaultImportRouteFilter = core.StringPtr("permit")
				createGatewayActionOptionsModel.ExportRouteFilters = []directlinkv1.GatewayTemplateRouteFilter{*gatewayTemplateRouteFilterModel}
				createGatewayActionOptionsModel.Global = core.BoolPtr(true)
				createGatewayActionOptionsModel.ImportRouteFilters = []directlinkv1.GatewayTemplateRouteFilter{*gatewayTemplateRouteFilterModel}
				createGatewayActionOptionsModel.Metered = core.BoolPtr(false)
				createGatewayActionOptionsModel.ResourceGroup = resourceGroupIdentityModel
				createGatewayActionOptionsModel.Updates = []directlinkv1.GatewayActionTemplateUpdatesItemIntf{gatewayActionTemplateUpdatesItemModel}
				createGatewayActionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = directLinkService.CreateGatewayAction(createGatewayActionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateGatewayAction with error: Operation validation and request error`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Construct an instance of the AsPrependTemplate model
				asPrependTemplateModel := new(directlinkv1.AsPrependTemplate)
				asPrependTemplateModel.Length = core.Int64Ptr(int64(4))
				asPrependTemplateModel.Policy = core.StringPtr("import")
				asPrependTemplateModel.Prefix = core.StringPtr("172.17.0.0/16")
				asPrependTemplateModel.SpecificPrefixes = []string{"192.168.3.0/24"}

				// Construct an instance of the AuthenticationKeyIdentityKeyProtectAuthenticationKeyIdentity model
				authenticationKeyIdentityModel := new(directlinkv1.AuthenticationKeyIdentityKeyProtectAuthenticationKeyIdentity)
				authenticationKeyIdentityModel.Crn = core.StringPtr("crn:v1:bluemix:public:kms:us-south:a/4111d05f36894e3cb9b46a43556d9000:abc111b8-37aa-4034-9def-f2607c87aaaa:key:bbb222bc-430a-4de9-9aad-84e5bb022222")

				// Construct an instance of the GatewayBfdConfigActionTemplate model
				gatewayBfdConfigActionTemplateModel := new(directlinkv1.GatewayBfdConfigActionTemplate)
				gatewayBfdConfigActionTemplateModel.Interval = core.Int64Ptr(int64(2000))
				gatewayBfdConfigActionTemplateModel.Multiplier = core.Int64Ptr(int64(10))

				// Construct an instance of the GatewayTemplateRouteFilter model
				gatewayTemplateRouteFilterModel := new(directlinkv1.GatewayTemplateRouteFilter)
				gatewayTemplateRouteFilterModel.Action = core.StringPtr("permit")
				gatewayTemplateRouteFilterModel.Ge = core.Int64Ptr(int64(25))
				gatewayTemplateRouteFilterModel.Le = core.Int64Ptr(int64(30))
				gatewayTemplateRouteFilterModel.Prefix = core.StringPtr("192.168.100.0/24")

				// Construct an instance of the ResourceGroupIdentity model
				resourceGroupIdentityModel := new(directlinkv1.ResourceGroupIdentity)
				resourceGroupIdentityModel.ID = core.StringPtr("56969d6043e9465c883cb9f7363e78e8")

				// Construct an instance of the GatewayActionTemplateUpdatesItemGatewayClientSpeedUpdate model
				gatewayActionTemplateUpdatesItemModel := new(directlinkv1.GatewayActionTemplateUpdatesItemGatewayClientSpeedUpdate)
				gatewayActionTemplateUpdatesItemModel.SpeedMbps = core.Int64Ptr(int64(500))

				// Construct an instance of the CreateGatewayActionOptions model
				createGatewayActionOptionsModel := new(directlinkv1.CreateGatewayActionOptions)
				createGatewayActionOptionsModel.ID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				createGatewayActionOptionsModel.Action = core.StringPtr("create_gateway_approve")
				createGatewayActionOptionsModel.AsPrepends = []directlinkv1.AsPrependTemplate{*asPrependTemplateModel}
				createGatewayActionOptionsModel.AuthenticationKey = authenticationKeyIdentityModel
				createGatewayActionOptionsModel.BfdConfig = gatewayBfdConfigActionTemplateModel
				createGatewayActionOptionsModel.ConnectionMode = core.StringPtr("transit")
				createGatewayActionOptionsModel.DefaultExportRouteFilter = core.StringPtr("permit")
				createGatewayActionOptionsModel.DefaultImportRouteFilter = core.StringPtr("permit")
				createGatewayActionOptionsModel.ExportRouteFilters = []directlinkv1.GatewayTemplateRouteFilter{*gatewayTemplateRouteFilterModel}
				createGatewayActionOptionsModel.Global = core.BoolPtr(true)
				createGatewayActionOptionsModel.ImportRouteFilters = []directlinkv1.GatewayTemplateRouteFilter{*gatewayTemplateRouteFilterModel}
				createGatewayActionOptionsModel.Metered = core.BoolPtr(false)
				createGatewayActionOptionsModel.ResourceGroup = resourceGroupIdentityModel
				createGatewayActionOptionsModel.Updates = []directlinkv1.GatewayActionTemplateUpdatesItemIntf{gatewayActionTemplateUpdatesItemModel}
				createGatewayActionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := directLinkService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := directLinkService.CreateGatewayAction(createGatewayActionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateGatewayActionOptions model with no property values
				createGatewayActionOptionsModelNew := new(directlinkv1.CreateGatewayActionOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = directLinkService.CreateGatewayAction(createGatewayActionOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke CreateGatewayAction successfully`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Construct an instance of the AsPrependTemplate model
				asPrependTemplateModel := new(directlinkv1.AsPrependTemplate)
				asPrependTemplateModel.Length = core.Int64Ptr(int64(4))
				asPrependTemplateModel.Policy = core.StringPtr("import")
				asPrependTemplateModel.Prefix = core.StringPtr("172.17.0.0/16")
				asPrependTemplateModel.SpecificPrefixes = []string{"192.168.3.0/24"}

				// Construct an instance of the AuthenticationKeyIdentityKeyProtectAuthenticationKeyIdentity model
				authenticationKeyIdentityModel := new(directlinkv1.AuthenticationKeyIdentityKeyProtectAuthenticationKeyIdentity)
				authenticationKeyIdentityModel.Crn = core.StringPtr("crn:v1:bluemix:public:kms:us-south:a/4111d05f36894e3cb9b46a43556d9000:abc111b8-37aa-4034-9def-f2607c87aaaa:key:bbb222bc-430a-4de9-9aad-84e5bb022222")

				// Construct an instance of the GatewayBfdConfigActionTemplate model
				gatewayBfdConfigActionTemplateModel := new(directlinkv1.GatewayBfdConfigActionTemplate)
				gatewayBfdConfigActionTemplateModel.Interval = core.Int64Ptr(int64(2000))
				gatewayBfdConfigActionTemplateModel.Multiplier = core.Int64Ptr(int64(10))

				// Construct an instance of the GatewayTemplateRouteFilter model
				gatewayTemplateRouteFilterModel := new(directlinkv1.GatewayTemplateRouteFilter)
				gatewayTemplateRouteFilterModel.Action = core.StringPtr("permit")
				gatewayTemplateRouteFilterModel.Ge = core.Int64Ptr(int64(25))
				gatewayTemplateRouteFilterModel.Le = core.Int64Ptr(int64(30))
				gatewayTemplateRouteFilterModel.Prefix = core.StringPtr("192.168.100.0/24")

				// Construct an instance of the ResourceGroupIdentity model
				resourceGroupIdentityModel := new(directlinkv1.ResourceGroupIdentity)
				resourceGroupIdentityModel.ID = core.StringPtr("56969d6043e9465c883cb9f7363e78e8")

				// Construct an instance of the GatewayActionTemplateUpdatesItemGatewayClientSpeedUpdate model
				gatewayActionTemplateUpdatesItemModel := new(directlinkv1.GatewayActionTemplateUpdatesItemGatewayClientSpeedUpdate)
				gatewayActionTemplateUpdatesItemModel.SpeedMbps = core.Int64Ptr(int64(500))

				// Construct an instance of the CreateGatewayActionOptions model
				createGatewayActionOptionsModel := new(directlinkv1.CreateGatewayActionOptions)
				createGatewayActionOptionsModel.ID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				createGatewayActionOptionsModel.Action = core.StringPtr("create_gateway_approve")
				createGatewayActionOptionsModel.AsPrepends = []directlinkv1.AsPrependTemplate{*asPrependTemplateModel}
				createGatewayActionOptionsModel.AuthenticationKey = authenticationKeyIdentityModel
				createGatewayActionOptionsModel.BfdConfig = gatewayBfdConfigActionTemplateModel
				createGatewayActionOptionsModel.ConnectionMode = core.StringPtr("transit")
				createGatewayActionOptionsModel.DefaultExportRouteFilter = core.StringPtr("permit")
				createGatewayActionOptionsModel.DefaultImportRouteFilter = core.StringPtr("permit")
				createGatewayActionOptionsModel.ExportRouteFilters = []directlinkv1.GatewayTemplateRouteFilter{*gatewayTemplateRouteFilterModel}
				createGatewayActionOptionsModel.Global = core.BoolPtr(true)
				createGatewayActionOptionsModel.ImportRouteFilters = []directlinkv1.GatewayTemplateRouteFilter{*gatewayTemplateRouteFilterModel}
				createGatewayActionOptionsModel.Metered = core.BoolPtr(false)
				createGatewayActionOptionsModel.ResourceGroup = resourceGroupIdentityModel
				createGatewayActionOptionsModel.Updates = []directlinkv1.GatewayActionTemplateUpdatesItemIntf{gatewayActionTemplateUpdatesItemModel}
				createGatewayActionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := directLinkService.CreateGatewayAction(createGatewayActionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListGatewayCompletionNotice(listGatewayCompletionNoticeOptions *ListGatewayCompletionNoticeOptions)`, func() {
		version := "testString"
		listGatewayCompletionNoticePath := "/gateways/0a06fb9b-820f-4c44-8a31-77f1f0806d28/completion_notice"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listGatewayCompletionNoticePath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/pdf")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `This is a mock binary response.`)
				}))
			})
			It(`Invoke ListGatewayCompletionNotice successfully with retries`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())
				directLinkService.EnableRetries(0, 0)

				// Construct an instance of the ListGatewayCompletionNoticeOptions model
				listGatewayCompletionNoticeOptionsModel := new(directlinkv1.ListGatewayCompletionNoticeOptions)
				listGatewayCompletionNoticeOptionsModel.ID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				listGatewayCompletionNoticeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := directLinkService.ListGatewayCompletionNoticeWithContext(ctx, listGatewayCompletionNoticeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				directLinkService.DisableRetries()
				result, response, operationErr := directLinkService.ListGatewayCompletionNotice(listGatewayCompletionNoticeOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = directLinkService.ListGatewayCompletionNoticeWithContext(ctx, listGatewayCompletionNoticeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listGatewayCompletionNoticePath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/pdf")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `This is a mock binary response.`)
				}))
			})
			It(`Invoke ListGatewayCompletionNotice successfully`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := directLinkService.ListGatewayCompletionNotice(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListGatewayCompletionNoticeOptions model
				listGatewayCompletionNoticeOptionsModel := new(directlinkv1.ListGatewayCompletionNoticeOptions)
				listGatewayCompletionNoticeOptionsModel.ID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				listGatewayCompletionNoticeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = directLinkService.ListGatewayCompletionNotice(listGatewayCompletionNoticeOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListGatewayCompletionNotice with error: Operation validation and request error`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Construct an instance of the ListGatewayCompletionNoticeOptions model
				listGatewayCompletionNoticeOptionsModel := new(directlinkv1.ListGatewayCompletionNoticeOptions)
				listGatewayCompletionNoticeOptionsModel.ID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				listGatewayCompletionNoticeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := directLinkService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := directLinkService.ListGatewayCompletionNotice(listGatewayCompletionNoticeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListGatewayCompletionNoticeOptions model with no property values
				listGatewayCompletionNoticeOptionsModelNew := new(directlinkv1.ListGatewayCompletionNoticeOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = directLinkService.ListGatewayCompletionNotice(listGatewayCompletionNoticeOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ListGatewayCompletionNotice successfully`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Construct an instance of the ListGatewayCompletionNoticeOptions model
				listGatewayCompletionNoticeOptionsModel := new(directlinkv1.ListGatewayCompletionNoticeOptions)
				listGatewayCompletionNoticeOptionsModel.ID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				listGatewayCompletionNoticeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := directLinkService.ListGatewayCompletionNotice(listGatewayCompletionNoticeOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify empty byte buffer.
				Expect(result).ToNot(BeNil())
				buffer, operationErr := io.ReadAll(result)
				Expect(operationErr).To(BeNil())
				Expect(buffer).ToNot(BeNil())
				Expect(len(buffer)).To(Equal(0))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateGatewayCompletionNotice(createGatewayCompletionNoticeOptions *CreateGatewayCompletionNoticeOptions)`, func() {
		version := "testString"
		createGatewayCompletionNoticePath := "/gateways/0a06fb9b-820f-4c44-8a31-77f1f0806d28/completion_notice"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createGatewayCompletionNoticePath))
					Expect(req.Method).To(Equal("PUT"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.WriteHeader(204)
				}))
			})
			It(`Invoke CreateGatewayCompletionNotice successfully`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := directLinkService.CreateGatewayCompletionNotice(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the CreateGatewayCompletionNoticeOptions model
				createGatewayCompletionNoticeOptionsModel := new(directlinkv1.CreateGatewayCompletionNoticeOptions)
				createGatewayCompletionNoticeOptionsModel.ID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				createGatewayCompletionNoticeOptionsModel.Upload = CreateMockReader("This is a mock file.")
				createGatewayCompletionNoticeOptionsModel.UploadContentType = core.StringPtr("testString")
				createGatewayCompletionNoticeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = directLinkService.CreateGatewayCompletionNotice(createGatewayCompletionNoticeOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke CreateGatewayCompletionNotice with error: Param validation error`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Construct an instance of the CreateGatewayCompletionNoticeOptions model
				createGatewayCompletionNoticeOptionsModel := new(directlinkv1.CreateGatewayCompletionNoticeOptions)
				// Invoke operation with invalid options model (negative test)
				response, operationErr := directLinkService.CreateGatewayCompletionNotice(createGatewayCompletionNoticeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			It(`Invoke CreateGatewayCompletionNotice with error: Operation validation and request error`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Construct an instance of the CreateGatewayCompletionNoticeOptions model
				createGatewayCompletionNoticeOptionsModel := new(directlinkv1.CreateGatewayCompletionNoticeOptions)
				createGatewayCompletionNoticeOptionsModel.ID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				createGatewayCompletionNoticeOptionsModel.Upload = CreateMockReader("This is a mock file.")
				createGatewayCompletionNoticeOptionsModel.UploadContentType = core.StringPtr("testString")
				createGatewayCompletionNoticeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := directLinkService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := directLinkService.CreateGatewayCompletionNotice(createGatewayCompletionNoticeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the CreateGatewayCompletionNoticeOptions model with no property values
				createGatewayCompletionNoticeOptionsModelNew := new(directlinkv1.CreateGatewayCompletionNoticeOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = directLinkService.CreateGatewayCompletionNotice(createGatewayCompletionNoticeOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListGatewayLetterOfAuthorization(listGatewayLetterOfAuthorizationOptions *ListGatewayLetterOfAuthorizationOptions)`, func() {
		version := "testString"
		listGatewayLetterOfAuthorizationPath := "/gateways/0a06fb9b-820f-4c44-8a31-77f1f0806d28/letter_of_authorization"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listGatewayLetterOfAuthorizationPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/pdf")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `This is a mock binary response.`)
				}))
			})
			It(`Invoke ListGatewayLetterOfAuthorization successfully with retries`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())
				directLinkService.EnableRetries(0, 0)

				// Construct an instance of the ListGatewayLetterOfAuthorizationOptions model
				listGatewayLetterOfAuthorizationOptionsModel := new(directlinkv1.ListGatewayLetterOfAuthorizationOptions)
				listGatewayLetterOfAuthorizationOptionsModel.ID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				listGatewayLetterOfAuthorizationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := directLinkService.ListGatewayLetterOfAuthorizationWithContext(ctx, listGatewayLetterOfAuthorizationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				directLinkService.DisableRetries()
				result, response, operationErr := directLinkService.ListGatewayLetterOfAuthorization(listGatewayLetterOfAuthorizationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = directLinkService.ListGatewayLetterOfAuthorizationWithContext(ctx, listGatewayLetterOfAuthorizationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listGatewayLetterOfAuthorizationPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/pdf")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `This is a mock binary response.`)
				}))
			})
			It(`Invoke ListGatewayLetterOfAuthorization successfully`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := directLinkService.ListGatewayLetterOfAuthorization(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListGatewayLetterOfAuthorizationOptions model
				listGatewayLetterOfAuthorizationOptionsModel := new(directlinkv1.ListGatewayLetterOfAuthorizationOptions)
				listGatewayLetterOfAuthorizationOptionsModel.ID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				listGatewayLetterOfAuthorizationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = directLinkService.ListGatewayLetterOfAuthorization(listGatewayLetterOfAuthorizationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListGatewayLetterOfAuthorization with error: Operation validation and request error`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Construct an instance of the ListGatewayLetterOfAuthorizationOptions model
				listGatewayLetterOfAuthorizationOptionsModel := new(directlinkv1.ListGatewayLetterOfAuthorizationOptions)
				listGatewayLetterOfAuthorizationOptionsModel.ID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				listGatewayLetterOfAuthorizationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := directLinkService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := directLinkService.ListGatewayLetterOfAuthorization(listGatewayLetterOfAuthorizationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListGatewayLetterOfAuthorizationOptions model with no property values
				listGatewayLetterOfAuthorizationOptionsModelNew := new(directlinkv1.ListGatewayLetterOfAuthorizationOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = directLinkService.ListGatewayLetterOfAuthorization(listGatewayLetterOfAuthorizationOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ListGatewayLetterOfAuthorization successfully`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Construct an instance of the ListGatewayLetterOfAuthorizationOptions model
				listGatewayLetterOfAuthorizationOptionsModel := new(directlinkv1.ListGatewayLetterOfAuthorizationOptions)
				listGatewayLetterOfAuthorizationOptionsModel.ID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				listGatewayLetterOfAuthorizationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := directLinkService.ListGatewayLetterOfAuthorization(listGatewayLetterOfAuthorizationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify empty byte buffer.
				Expect(result).ToNot(BeNil())
				buffer, operationErr := io.ReadAll(result)
				Expect(operationErr).To(BeNil())
				Expect(buffer).ToNot(BeNil())
				Expect(len(buffer)).To(Equal(0))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetGatewayStatistics(getGatewayStatisticsOptions *GetGatewayStatisticsOptions) - Operation response error`, func() {
		version := "testString"
		getGatewayStatisticsPath := "/gateways/0a06fb9b-820f-4c44-8a31-77f1f0806d28/statistics"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getGatewayStatisticsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["type"]).To(Equal([]string{"macsec_mka_session"}))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetGatewayStatistics with error: Operation response processing error`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Construct an instance of the GetGatewayStatisticsOptions model
				getGatewayStatisticsOptionsModel := new(directlinkv1.GetGatewayStatisticsOptions)
				getGatewayStatisticsOptionsModel.ID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				getGatewayStatisticsOptionsModel.Type = core.StringPtr("macsec_mka_session")
				getGatewayStatisticsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := directLinkService.GetGatewayStatistics(getGatewayStatisticsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				directLinkService.EnableRetries(0, 0)
				result, response, operationErr = directLinkService.GetGatewayStatistics(getGatewayStatisticsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetGatewayStatistics(getGatewayStatisticsOptions *GetGatewayStatisticsOptions)`, func() {
		version := "testString"
		getGatewayStatisticsPath := "/gateways/0a06fb9b-820f-4c44-8a31-77f1f0806d28/statistics"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getGatewayStatisticsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["type"]).To(Equal([]string{"macsec_mka_session"}))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"statistics": [{"created_at": "2020-08-20T06:58:41.909Z", "data": "MKA statistics text...", "type": "macsec_policy"}]}`)
				}))
			})
			It(`Invoke GetGatewayStatistics successfully with retries`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())
				directLinkService.EnableRetries(0, 0)

				// Construct an instance of the GetGatewayStatisticsOptions model
				getGatewayStatisticsOptionsModel := new(directlinkv1.GetGatewayStatisticsOptions)
				getGatewayStatisticsOptionsModel.ID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				getGatewayStatisticsOptionsModel.Type = core.StringPtr("macsec_mka_session")
				getGatewayStatisticsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := directLinkService.GetGatewayStatisticsWithContext(ctx, getGatewayStatisticsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				directLinkService.DisableRetries()
				result, response, operationErr := directLinkService.GetGatewayStatistics(getGatewayStatisticsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = directLinkService.GetGatewayStatisticsWithContext(ctx, getGatewayStatisticsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getGatewayStatisticsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["type"]).To(Equal([]string{"macsec_mka_session"}))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"statistics": [{"created_at": "2020-08-20T06:58:41.909Z", "data": "MKA statistics text...", "type": "macsec_policy"}]}`)
				}))
			})
			It(`Invoke GetGatewayStatistics successfully`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := directLinkService.GetGatewayStatistics(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetGatewayStatisticsOptions model
				getGatewayStatisticsOptionsModel := new(directlinkv1.GetGatewayStatisticsOptions)
				getGatewayStatisticsOptionsModel.ID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				getGatewayStatisticsOptionsModel.Type = core.StringPtr("macsec_mka_session")
				getGatewayStatisticsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = directLinkService.GetGatewayStatistics(getGatewayStatisticsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetGatewayStatistics with error: Operation validation and request error`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Construct an instance of the GetGatewayStatisticsOptions model
				getGatewayStatisticsOptionsModel := new(directlinkv1.GetGatewayStatisticsOptions)
				getGatewayStatisticsOptionsModel.ID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				getGatewayStatisticsOptionsModel.Type = core.StringPtr("macsec_mka_session")
				getGatewayStatisticsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := directLinkService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := directLinkService.GetGatewayStatistics(getGatewayStatisticsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetGatewayStatisticsOptions model with no property values
				getGatewayStatisticsOptionsModelNew := new(directlinkv1.GetGatewayStatisticsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = directLinkService.GetGatewayStatistics(getGatewayStatisticsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetGatewayStatistics successfully`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Construct an instance of the GetGatewayStatisticsOptions model
				getGatewayStatisticsOptionsModel := new(directlinkv1.GetGatewayStatisticsOptions)
				getGatewayStatisticsOptionsModel.ID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				getGatewayStatisticsOptionsModel.Type = core.StringPtr("macsec_mka_session")
				getGatewayStatisticsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := directLinkService.GetGatewayStatistics(getGatewayStatisticsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetGatewayStatus(getGatewayStatusOptions *GetGatewayStatusOptions) - Operation response error`, func() {
		version := "testString"
		getGatewayStatusPath := "/gateways/0a06fb9b-820f-4c44-8a31-77f1f0806d28/status"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getGatewayStatusPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["type"]).To(Equal([]string{"bgp"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetGatewayStatus with error: Operation response processing error`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Construct an instance of the GetGatewayStatusOptions model
				getGatewayStatusOptionsModel := new(directlinkv1.GetGatewayStatusOptions)
				getGatewayStatusOptionsModel.ID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				getGatewayStatusOptionsModel.Type = core.StringPtr("bgp")
				getGatewayStatusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := directLinkService.GetGatewayStatus(getGatewayStatusOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				directLinkService.EnableRetries(0, 0)
				result, response, operationErr = directLinkService.GetGatewayStatus(getGatewayStatusOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetGatewayStatus(getGatewayStatusOptions *GetGatewayStatusOptions)`, func() {
		version := "testString"
		getGatewayStatusPath := "/gateways/0a06fb9b-820f-4c44-8a31-77f1f0806d28/status"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getGatewayStatusPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["type"]).To(Equal([]string{"bgp"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"status": [{"type": "bgp", "updated_at": "2020-08-20T06:58:41.909Z", "value": "active"}]}`)
				}))
			})
			It(`Invoke GetGatewayStatus successfully with retries`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())
				directLinkService.EnableRetries(0, 0)

				// Construct an instance of the GetGatewayStatusOptions model
				getGatewayStatusOptionsModel := new(directlinkv1.GetGatewayStatusOptions)
				getGatewayStatusOptionsModel.ID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				getGatewayStatusOptionsModel.Type = core.StringPtr("bgp")
				getGatewayStatusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := directLinkService.GetGatewayStatusWithContext(ctx, getGatewayStatusOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				directLinkService.DisableRetries()
				result, response, operationErr := directLinkService.GetGatewayStatus(getGatewayStatusOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = directLinkService.GetGatewayStatusWithContext(ctx, getGatewayStatusOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getGatewayStatusPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["type"]).To(Equal([]string{"bgp"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"status": [{"type": "bgp", "updated_at": "2020-08-20T06:58:41.909Z", "value": "active"}]}`)
				}))
			})
			It(`Invoke GetGatewayStatus successfully`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := directLinkService.GetGatewayStatus(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetGatewayStatusOptions model
				getGatewayStatusOptionsModel := new(directlinkv1.GetGatewayStatusOptions)
				getGatewayStatusOptionsModel.ID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				getGatewayStatusOptionsModel.Type = core.StringPtr("bgp")
				getGatewayStatusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = directLinkService.GetGatewayStatus(getGatewayStatusOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetGatewayStatus with error: Operation validation and request error`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Construct an instance of the GetGatewayStatusOptions model
				getGatewayStatusOptionsModel := new(directlinkv1.GetGatewayStatusOptions)
				getGatewayStatusOptionsModel.ID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				getGatewayStatusOptionsModel.Type = core.StringPtr("bgp")
				getGatewayStatusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := directLinkService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := directLinkService.GetGatewayStatus(getGatewayStatusOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetGatewayStatusOptions model with no property values
				getGatewayStatusOptionsModelNew := new(directlinkv1.GetGatewayStatusOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = directLinkService.GetGatewayStatus(getGatewayStatusOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetGatewayStatus successfully`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Construct an instance of the GetGatewayStatusOptions model
				getGatewayStatusOptionsModel := new(directlinkv1.GetGatewayStatusOptions)
				getGatewayStatusOptionsModel.ID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				getGatewayStatusOptionsModel.Type = core.StringPtr("bgp")
				getGatewayStatusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := directLinkService.GetGatewayStatus(getGatewayStatusOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListGatewayAsPrepends(listGatewayAsPrependsOptions *ListGatewayAsPrependsOptions) - Operation response error`, func() {
		version := "testString"
		listGatewayAsPrependsPath := "/gateways/0a06fb9b-820f-4c44-8a31-77f1f0806d28/as_prepends"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listGatewayAsPrependsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListGatewayAsPrepends with error: Operation response processing error`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Construct an instance of the ListGatewayAsPrependsOptions model
				listGatewayAsPrependsOptionsModel := new(directlinkv1.ListGatewayAsPrependsOptions)
				listGatewayAsPrependsOptionsModel.GatewayID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				listGatewayAsPrependsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := directLinkService.ListGatewayAsPrepends(listGatewayAsPrependsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				directLinkService.EnableRetries(0, 0)
				result, response, operationErr = directLinkService.ListGatewayAsPrepends(listGatewayAsPrependsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListGatewayAsPrepends(listGatewayAsPrependsOptions *ListGatewayAsPrependsOptions)`, func() {
		version := "testString"
		listGatewayAsPrependsPath := "/gateways/0a06fb9b-820f-4c44-8a31-77f1f0806d28/as_prepends"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listGatewayAsPrependsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"as_prepends": [{"created_at": "2019-01-01T12:00:00.000Z", "id": "ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4", "length": 4, "policy": "import", "specific_prefixes": ["192.168.3.0/24"], "updated_at": "2019-01-01T12:00:00.000Z"}]}`)
				}))
			})
			It(`Invoke ListGatewayAsPrepends successfully with retries`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())
				directLinkService.EnableRetries(0, 0)

				// Construct an instance of the ListGatewayAsPrependsOptions model
				listGatewayAsPrependsOptionsModel := new(directlinkv1.ListGatewayAsPrependsOptions)
				listGatewayAsPrependsOptionsModel.GatewayID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				listGatewayAsPrependsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := directLinkService.ListGatewayAsPrependsWithContext(ctx, listGatewayAsPrependsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				directLinkService.DisableRetries()
				result, response, operationErr := directLinkService.ListGatewayAsPrepends(listGatewayAsPrependsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = directLinkService.ListGatewayAsPrependsWithContext(ctx, listGatewayAsPrependsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listGatewayAsPrependsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"as_prepends": [{"created_at": "2019-01-01T12:00:00.000Z", "id": "ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4", "length": 4, "policy": "import", "specific_prefixes": ["192.168.3.0/24"], "updated_at": "2019-01-01T12:00:00.000Z"}]}`)
				}))
			})
			It(`Invoke ListGatewayAsPrepends successfully`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := directLinkService.ListGatewayAsPrepends(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListGatewayAsPrependsOptions model
				listGatewayAsPrependsOptionsModel := new(directlinkv1.ListGatewayAsPrependsOptions)
				listGatewayAsPrependsOptionsModel.GatewayID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				listGatewayAsPrependsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = directLinkService.ListGatewayAsPrepends(listGatewayAsPrependsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListGatewayAsPrepends with error: Operation validation and request error`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Construct an instance of the ListGatewayAsPrependsOptions model
				listGatewayAsPrependsOptionsModel := new(directlinkv1.ListGatewayAsPrependsOptions)
				listGatewayAsPrependsOptionsModel.GatewayID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				listGatewayAsPrependsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := directLinkService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := directLinkService.ListGatewayAsPrepends(listGatewayAsPrependsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListGatewayAsPrependsOptions model with no property values
				listGatewayAsPrependsOptionsModelNew := new(directlinkv1.ListGatewayAsPrependsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = directLinkService.ListGatewayAsPrepends(listGatewayAsPrependsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ListGatewayAsPrepends successfully`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Construct an instance of the ListGatewayAsPrependsOptions model
				listGatewayAsPrependsOptionsModel := new(directlinkv1.ListGatewayAsPrependsOptions)
				listGatewayAsPrependsOptionsModel.GatewayID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				listGatewayAsPrependsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := directLinkService.ListGatewayAsPrepends(listGatewayAsPrependsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ReplaceGatewayAsPrepends(replaceGatewayAsPrependsOptions *ReplaceGatewayAsPrependsOptions) - Operation response error`, func() {
		version := "testString"
		replaceGatewayAsPrependsPath := "/gateways/0a06fb9b-820f-4c44-8a31-77f1f0806d28/as_prepends"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(replaceGatewayAsPrependsPath))
					Expect(req.Method).To(Equal("PUT"))
					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "W/\"96d225c4-56bd-43d9-98fc-d7148e5c5028\"")))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ReplaceGatewayAsPrepends with error: Operation response processing error`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Construct an instance of the AsPrependPrefixArrayTemplate model
				asPrependPrefixArrayTemplateModel := new(directlinkv1.AsPrependPrefixArrayTemplate)
				asPrependPrefixArrayTemplateModel.Length = core.Int64Ptr(int64(4))
				asPrependPrefixArrayTemplateModel.Policy = core.StringPtr("import")
				asPrependPrefixArrayTemplateModel.SpecificPrefixes = []string{"192.168.3.0/24"}

				// Construct an instance of the ReplaceGatewayAsPrependsOptions model
				replaceGatewayAsPrependsOptionsModel := new(directlinkv1.ReplaceGatewayAsPrependsOptions)
				replaceGatewayAsPrependsOptionsModel.GatewayID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				replaceGatewayAsPrependsOptionsModel.IfMatch = core.StringPtr("W/\"96d225c4-56bd-43d9-98fc-d7148e5c5028\"")
				replaceGatewayAsPrependsOptionsModel.AsPrepends = []directlinkv1.AsPrependPrefixArrayTemplate{*asPrependPrefixArrayTemplateModel}
				replaceGatewayAsPrependsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := directLinkService.ReplaceGatewayAsPrepends(replaceGatewayAsPrependsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				directLinkService.EnableRetries(0, 0)
				result, response, operationErr = directLinkService.ReplaceGatewayAsPrepends(replaceGatewayAsPrependsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ReplaceGatewayAsPrepends(replaceGatewayAsPrependsOptions *ReplaceGatewayAsPrependsOptions)`, func() {
		version := "testString"
		replaceGatewayAsPrependsPath := "/gateways/0a06fb9b-820f-4c44-8a31-77f1f0806d28/as_prepends"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(replaceGatewayAsPrependsPath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "W/\"96d225c4-56bd-43d9-98fc-d7148e5c5028\"")))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"as_prepends": [{"created_at": "2019-01-01T12:00:00.000Z", "id": "ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4", "length": 4, "policy": "import", "specific_prefixes": ["192.168.3.0/24"], "updated_at": "2019-01-01T12:00:00.000Z"}]}`)
				}))
			})
			It(`Invoke ReplaceGatewayAsPrepends successfully with retries`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())
				directLinkService.EnableRetries(0, 0)

				// Construct an instance of the AsPrependPrefixArrayTemplate model
				asPrependPrefixArrayTemplateModel := new(directlinkv1.AsPrependPrefixArrayTemplate)
				asPrependPrefixArrayTemplateModel.Length = core.Int64Ptr(int64(4))
				asPrependPrefixArrayTemplateModel.Policy = core.StringPtr("import")
				asPrependPrefixArrayTemplateModel.SpecificPrefixes = []string{"192.168.3.0/24"}

				// Construct an instance of the ReplaceGatewayAsPrependsOptions model
				replaceGatewayAsPrependsOptionsModel := new(directlinkv1.ReplaceGatewayAsPrependsOptions)
				replaceGatewayAsPrependsOptionsModel.GatewayID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				replaceGatewayAsPrependsOptionsModel.IfMatch = core.StringPtr("W/\"96d225c4-56bd-43d9-98fc-d7148e5c5028\"")
				replaceGatewayAsPrependsOptionsModel.AsPrepends = []directlinkv1.AsPrependPrefixArrayTemplate{*asPrependPrefixArrayTemplateModel}
				replaceGatewayAsPrependsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := directLinkService.ReplaceGatewayAsPrependsWithContext(ctx, replaceGatewayAsPrependsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				directLinkService.DisableRetries()
				result, response, operationErr := directLinkService.ReplaceGatewayAsPrepends(replaceGatewayAsPrependsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = directLinkService.ReplaceGatewayAsPrependsWithContext(ctx, replaceGatewayAsPrependsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(replaceGatewayAsPrependsPath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "W/\"96d225c4-56bd-43d9-98fc-d7148e5c5028\"")))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"as_prepends": [{"created_at": "2019-01-01T12:00:00.000Z", "id": "ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4", "length": 4, "policy": "import", "specific_prefixes": ["192.168.3.0/24"], "updated_at": "2019-01-01T12:00:00.000Z"}]}`)
				}))
			})
			It(`Invoke ReplaceGatewayAsPrepends successfully`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := directLinkService.ReplaceGatewayAsPrepends(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the AsPrependPrefixArrayTemplate model
				asPrependPrefixArrayTemplateModel := new(directlinkv1.AsPrependPrefixArrayTemplate)
				asPrependPrefixArrayTemplateModel.Length = core.Int64Ptr(int64(4))
				asPrependPrefixArrayTemplateModel.Policy = core.StringPtr("import")
				asPrependPrefixArrayTemplateModel.SpecificPrefixes = []string{"192.168.3.0/24"}

				// Construct an instance of the ReplaceGatewayAsPrependsOptions model
				replaceGatewayAsPrependsOptionsModel := new(directlinkv1.ReplaceGatewayAsPrependsOptions)
				replaceGatewayAsPrependsOptionsModel.GatewayID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				replaceGatewayAsPrependsOptionsModel.IfMatch = core.StringPtr("W/\"96d225c4-56bd-43d9-98fc-d7148e5c5028\"")
				replaceGatewayAsPrependsOptionsModel.AsPrepends = []directlinkv1.AsPrependPrefixArrayTemplate{*asPrependPrefixArrayTemplateModel}
				replaceGatewayAsPrependsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = directLinkService.ReplaceGatewayAsPrepends(replaceGatewayAsPrependsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ReplaceGatewayAsPrepends with error: Operation validation and request error`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Construct an instance of the AsPrependPrefixArrayTemplate model
				asPrependPrefixArrayTemplateModel := new(directlinkv1.AsPrependPrefixArrayTemplate)
				asPrependPrefixArrayTemplateModel.Length = core.Int64Ptr(int64(4))
				asPrependPrefixArrayTemplateModel.Policy = core.StringPtr("import")
				asPrependPrefixArrayTemplateModel.SpecificPrefixes = []string{"192.168.3.0/24"}

				// Construct an instance of the ReplaceGatewayAsPrependsOptions model
				replaceGatewayAsPrependsOptionsModel := new(directlinkv1.ReplaceGatewayAsPrependsOptions)
				replaceGatewayAsPrependsOptionsModel.GatewayID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				replaceGatewayAsPrependsOptionsModel.IfMatch = core.StringPtr("W/\"96d225c4-56bd-43d9-98fc-d7148e5c5028\"")
				replaceGatewayAsPrependsOptionsModel.AsPrepends = []directlinkv1.AsPrependPrefixArrayTemplate{*asPrependPrefixArrayTemplateModel}
				replaceGatewayAsPrependsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := directLinkService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := directLinkService.ReplaceGatewayAsPrepends(replaceGatewayAsPrependsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ReplaceGatewayAsPrependsOptions model with no property values
				replaceGatewayAsPrependsOptionsModelNew := new(directlinkv1.ReplaceGatewayAsPrependsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = directLinkService.ReplaceGatewayAsPrepends(replaceGatewayAsPrependsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(201)
				}))
			})
			It(`Invoke ReplaceGatewayAsPrepends successfully`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Construct an instance of the AsPrependPrefixArrayTemplate model
				asPrependPrefixArrayTemplateModel := new(directlinkv1.AsPrependPrefixArrayTemplate)
				asPrependPrefixArrayTemplateModel.Length = core.Int64Ptr(int64(4))
				asPrependPrefixArrayTemplateModel.Policy = core.StringPtr("import")
				asPrependPrefixArrayTemplateModel.SpecificPrefixes = []string{"192.168.3.0/24"}

				// Construct an instance of the ReplaceGatewayAsPrependsOptions model
				replaceGatewayAsPrependsOptionsModel := new(directlinkv1.ReplaceGatewayAsPrependsOptions)
				replaceGatewayAsPrependsOptionsModel.GatewayID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				replaceGatewayAsPrependsOptionsModel.IfMatch = core.StringPtr("W/\"96d225c4-56bd-43d9-98fc-d7148e5c5028\"")
				replaceGatewayAsPrependsOptionsModel.AsPrepends = []directlinkv1.AsPrependPrefixArrayTemplate{*asPrependPrefixArrayTemplateModel}
				replaceGatewayAsPrependsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := directLinkService.ReplaceGatewayAsPrepends(replaceGatewayAsPrependsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListGatewayExportRouteFilters(listGatewayExportRouteFiltersOptions *ListGatewayExportRouteFiltersOptions) - Operation response error`, func() {
		version := "testString"
		listGatewayExportRouteFiltersPath := "/gateways/0a06fb9b-820f-4c44-8a31-77f1f0806d28/export_route_filters"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listGatewayExportRouteFiltersPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListGatewayExportRouteFilters with error: Operation response processing error`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Construct an instance of the ListGatewayExportRouteFiltersOptions model
				listGatewayExportRouteFiltersOptionsModel := new(directlinkv1.ListGatewayExportRouteFiltersOptions)
				listGatewayExportRouteFiltersOptionsModel.GatewayID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				listGatewayExportRouteFiltersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := directLinkService.ListGatewayExportRouteFilters(listGatewayExportRouteFiltersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				directLinkService.EnableRetries(0, 0)
				result, response, operationErr = directLinkService.ListGatewayExportRouteFilters(listGatewayExportRouteFiltersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListGatewayExportRouteFilters(listGatewayExportRouteFiltersOptions *ListGatewayExportRouteFiltersOptions)`, func() {
		version := "testString"
		listGatewayExportRouteFiltersPath := "/gateways/0a06fb9b-820f-4c44-8a31-77f1f0806d28/export_route_filters"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listGatewayExportRouteFiltersPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"export_route_filters": [{"action": "permit", "before": "1a15dcab-7e40-45e1-b7c5-bc690eaa9782", "created_at": "2020-11-02T20:40:29.622Z", "ge": 25, "id": "1a15dcab-7e40-45e1-b7c5-bc690eaa9782", "le": 30, "prefix": "192.168.100.0/24", "updated_at": "2020-11-02T20:40:29.622Z"}]}`)
				}))
			})
			It(`Invoke ListGatewayExportRouteFilters successfully with retries`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())
				directLinkService.EnableRetries(0, 0)

				// Construct an instance of the ListGatewayExportRouteFiltersOptions model
				listGatewayExportRouteFiltersOptionsModel := new(directlinkv1.ListGatewayExportRouteFiltersOptions)
				listGatewayExportRouteFiltersOptionsModel.GatewayID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				listGatewayExportRouteFiltersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := directLinkService.ListGatewayExportRouteFiltersWithContext(ctx, listGatewayExportRouteFiltersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				directLinkService.DisableRetries()
				result, response, operationErr := directLinkService.ListGatewayExportRouteFilters(listGatewayExportRouteFiltersOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = directLinkService.ListGatewayExportRouteFiltersWithContext(ctx, listGatewayExportRouteFiltersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listGatewayExportRouteFiltersPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"export_route_filters": [{"action": "permit", "before": "1a15dcab-7e40-45e1-b7c5-bc690eaa9782", "created_at": "2020-11-02T20:40:29.622Z", "ge": 25, "id": "1a15dcab-7e40-45e1-b7c5-bc690eaa9782", "le": 30, "prefix": "192.168.100.0/24", "updated_at": "2020-11-02T20:40:29.622Z"}]}`)
				}))
			})
			It(`Invoke ListGatewayExportRouteFilters successfully`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := directLinkService.ListGatewayExportRouteFilters(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListGatewayExportRouteFiltersOptions model
				listGatewayExportRouteFiltersOptionsModel := new(directlinkv1.ListGatewayExportRouteFiltersOptions)
				listGatewayExportRouteFiltersOptionsModel.GatewayID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				listGatewayExportRouteFiltersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = directLinkService.ListGatewayExportRouteFilters(listGatewayExportRouteFiltersOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListGatewayExportRouteFilters with error: Operation validation and request error`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Construct an instance of the ListGatewayExportRouteFiltersOptions model
				listGatewayExportRouteFiltersOptionsModel := new(directlinkv1.ListGatewayExportRouteFiltersOptions)
				listGatewayExportRouteFiltersOptionsModel.GatewayID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				listGatewayExportRouteFiltersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := directLinkService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := directLinkService.ListGatewayExportRouteFilters(listGatewayExportRouteFiltersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListGatewayExportRouteFiltersOptions model with no property values
				listGatewayExportRouteFiltersOptionsModelNew := new(directlinkv1.ListGatewayExportRouteFiltersOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = directLinkService.ListGatewayExportRouteFilters(listGatewayExportRouteFiltersOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ListGatewayExportRouteFilters successfully`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Construct an instance of the ListGatewayExportRouteFiltersOptions model
				listGatewayExportRouteFiltersOptionsModel := new(directlinkv1.ListGatewayExportRouteFiltersOptions)
				listGatewayExportRouteFiltersOptionsModel.GatewayID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				listGatewayExportRouteFiltersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := directLinkService.ListGatewayExportRouteFilters(listGatewayExportRouteFiltersOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateGatewayExportRouteFilter(createGatewayExportRouteFilterOptions *CreateGatewayExportRouteFilterOptions) - Operation response error`, func() {
		version := "testString"
		createGatewayExportRouteFilterPath := "/gateways/0a06fb9b-820f-4c44-8a31-77f1f0806d28/export_route_filters"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createGatewayExportRouteFilterPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateGatewayExportRouteFilter with error: Operation response processing error`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Construct an instance of the CreateGatewayExportRouteFilterOptions model
				createGatewayExportRouteFilterOptionsModel := new(directlinkv1.CreateGatewayExportRouteFilterOptions)
				createGatewayExportRouteFilterOptionsModel.GatewayID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				createGatewayExportRouteFilterOptionsModel.Action = core.StringPtr("permit")
				createGatewayExportRouteFilterOptionsModel.Prefix = core.StringPtr("192.168.100.0/24")
				createGatewayExportRouteFilterOptionsModel.Before = core.StringPtr("1a15dcab-7e40-45e1-b7c5-bc690eaa9782")
				createGatewayExportRouteFilterOptionsModel.Ge = core.Int64Ptr(int64(25))
				createGatewayExportRouteFilterOptionsModel.Le = core.Int64Ptr(int64(30))
				createGatewayExportRouteFilterOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := directLinkService.CreateGatewayExportRouteFilter(createGatewayExportRouteFilterOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				directLinkService.EnableRetries(0, 0)
				result, response, operationErr = directLinkService.CreateGatewayExportRouteFilter(createGatewayExportRouteFilterOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateGatewayExportRouteFilter(createGatewayExportRouteFilterOptions *CreateGatewayExportRouteFilterOptions)`, func() {
		version := "testString"
		createGatewayExportRouteFilterPath := "/gateways/0a06fb9b-820f-4c44-8a31-77f1f0806d28/export_route_filters"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createGatewayExportRouteFilterPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"action": "permit", "before": "1a15dcab-7e40-45e1-b7c5-bc690eaa9782", "created_at": "2020-11-02T20:40:29.622Z", "ge": 25, "id": "1a15dcab-7e40-45e1-b7c5-bc690eaa9782", "le": 30, "prefix": "192.168.100.0/24", "updated_at": "2020-11-02T20:40:29.622Z"}`)
				}))
			})
			It(`Invoke CreateGatewayExportRouteFilter successfully with retries`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())
				directLinkService.EnableRetries(0, 0)

				// Construct an instance of the CreateGatewayExportRouteFilterOptions model
				createGatewayExportRouteFilterOptionsModel := new(directlinkv1.CreateGatewayExportRouteFilterOptions)
				createGatewayExportRouteFilterOptionsModel.GatewayID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				createGatewayExportRouteFilterOptionsModel.Action = core.StringPtr("permit")
				createGatewayExportRouteFilterOptionsModel.Prefix = core.StringPtr("192.168.100.0/24")
				createGatewayExportRouteFilterOptionsModel.Before = core.StringPtr("1a15dcab-7e40-45e1-b7c5-bc690eaa9782")
				createGatewayExportRouteFilterOptionsModel.Ge = core.Int64Ptr(int64(25))
				createGatewayExportRouteFilterOptionsModel.Le = core.Int64Ptr(int64(30))
				createGatewayExportRouteFilterOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := directLinkService.CreateGatewayExportRouteFilterWithContext(ctx, createGatewayExportRouteFilterOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				directLinkService.DisableRetries()
				result, response, operationErr := directLinkService.CreateGatewayExportRouteFilter(createGatewayExportRouteFilterOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = directLinkService.CreateGatewayExportRouteFilterWithContext(ctx, createGatewayExportRouteFilterOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createGatewayExportRouteFilterPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"action": "permit", "before": "1a15dcab-7e40-45e1-b7c5-bc690eaa9782", "created_at": "2020-11-02T20:40:29.622Z", "ge": 25, "id": "1a15dcab-7e40-45e1-b7c5-bc690eaa9782", "le": 30, "prefix": "192.168.100.0/24", "updated_at": "2020-11-02T20:40:29.622Z"}`)
				}))
			})
			It(`Invoke CreateGatewayExportRouteFilter successfully`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := directLinkService.CreateGatewayExportRouteFilter(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateGatewayExportRouteFilterOptions model
				createGatewayExportRouteFilterOptionsModel := new(directlinkv1.CreateGatewayExportRouteFilterOptions)
				createGatewayExportRouteFilterOptionsModel.GatewayID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				createGatewayExportRouteFilterOptionsModel.Action = core.StringPtr("permit")
				createGatewayExportRouteFilterOptionsModel.Prefix = core.StringPtr("192.168.100.0/24")
				createGatewayExportRouteFilterOptionsModel.Before = core.StringPtr("1a15dcab-7e40-45e1-b7c5-bc690eaa9782")
				createGatewayExportRouteFilterOptionsModel.Ge = core.Int64Ptr(int64(25))
				createGatewayExportRouteFilterOptionsModel.Le = core.Int64Ptr(int64(30))
				createGatewayExportRouteFilterOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = directLinkService.CreateGatewayExportRouteFilter(createGatewayExportRouteFilterOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateGatewayExportRouteFilter with error: Operation validation and request error`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Construct an instance of the CreateGatewayExportRouteFilterOptions model
				createGatewayExportRouteFilterOptionsModel := new(directlinkv1.CreateGatewayExportRouteFilterOptions)
				createGatewayExportRouteFilterOptionsModel.GatewayID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				createGatewayExportRouteFilterOptionsModel.Action = core.StringPtr("permit")
				createGatewayExportRouteFilterOptionsModel.Prefix = core.StringPtr("192.168.100.0/24")
				createGatewayExportRouteFilterOptionsModel.Before = core.StringPtr("1a15dcab-7e40-45e1-b7c5-bc690eaa9782")
				createGatewayExportRouteFilterOptionsModel.Ge = core.Int64Ptr(int64(25))
				createGatewayExportRouteFilterOptionsModel.Le = core.Int64Ptr(int64(30))
				createGatewayExportRouteFilterOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := directLinkService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := directLinkService.CreateGatewayExportRouteFilter(createGatewayExportRouteFilterOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateGatewayExportRouteFilterOptions model with no property values
				createGatewayExportRouteFilterOptionsModelNew := new(directlinkv1.CreateGatewayExportRouteFilterOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = directLinkService.CreateGatewayExportRouteFilter(createGatewayExportRouteFilterOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(201)
				}))
			})
			It(`Invoke CreateGatewayExportRouteFilter successfully`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Construct an instance of the CreateGatewayExportRouteFilterOptions model
				createGatewayExportRouteFilterOptionsModel := new(directlinkv1.CreateGatewayExportRouteFilterOptions)
				createGatewayExportRouteFilterOptionsModel.GatewayID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				createGatewayExportRouteFilterOptionsModel.Action = core.StringPtr("permit")
				createGatewayExportRouteFilterOptionsModel.Prefix = core.StringPtr("192.168.100.0/24")
				createGatewayExportRouteFilterOptionsModel.Before = core.StringPtr("1a15dcab-7e40-45e1-b7c5-bc690eaa9782")
				createGatewayExportRouteFilterOptionsModel.Ge = core.Int64Ptr(int64(25))
				createGatewayExportRouteFilterOptionsModel.Le = core.Int64Ptr(int64(30))
				createGatewayExportRouteFilterOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := directLinkService.CreateGatewayExportRouteFilter(createGatewayExportRouteFilterOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ReplaceGatewayExportRouteFilters(replaceGatewayExportRouteFiltersOptions *ReplaceGatewayExportRouteFiltersOptions) - Operation response error`, func() {
		version := "testString"
		replaceGatewayExportRouteFiltersPath := "/gateways/0a06fb9b-820f-4c44-8a31-77f1f0806d28/export_route_filters"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(replaceGatewayExportRouteFiltersPath))
					Expect(req.Method).To(Equal("PUT"))
					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "W/\"96d225c4-56bd-43d9-98fc-d7148e5c5028\"")))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ReplaceGatewayExportRouteFilters with error: Operation response processing error`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Construct an instance of the GatewayTemplateRouteFilter model
				gatewayTemplateRouteFilterModel := new(directlinkv1.GatewayTemplateRouteFilter)
				gatewayTemplateRouteFilterModel.Action = core.StringPtr("permit")
				gatewayTemplateRouteFilterModel.Ge = core.Int64Ptr(int64(25))
				gatewayTemplateRouteFilterModel.Le = core.Int64Ptr(int64(30))
				gatewayTemplateRouteFilterModel.Prefix = core.StringPtr("192.168.100.0/24")

				// Construct an instance of the ReplaceGatewayExportRouteFiltersOptions model
				replaceGatewayExportRouteFiltersOptionsModel := new(directlinkv1.ReplaceGatewayExportRouteFiltersOptions)
				replaceGatewayExportRouteFiltersOptionsModel.GatewayID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				replaceGatewayExportRouteFiltersOptionsModel.IfMatch = core.StringPtr("W/\"96d225c4-56bd-43d9-98fc-d7148e5c5028\"")
				replaceGatewayExportRouteFiltersOptionsModel.ExportRouteFilters = []directlinkv1.GatewayTemplateRouteFilter{*gatewayTemplateRouteFilterModel}
				replaceGatewayExportRouteFiltersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := directLinkService.ReplaceGatewayExportRouteFilters(replaceGatewayExportRouteFiltersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				directLinkService.EnableRetries(0, 0)
				result, response, operationErr = directLinkService.ReplaceGatewayExportRouteFilters(replaceGatewayExportRouteFiltersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ReplaceGatewayExportRouteFilters(replaceGatewayExportRouteFiltersOptions *ReplaceGatewayExportRouteFiltersOptions)`, func() {
		version := "testString"
		replaceGatewayExportRouteFiltersPath := "/gateways/0a06fb9b-820f-4c44-8a31-77f1f0806d28/export_route_filters"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(replaceGatewayExportRouteFiltersPath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "W/\"96d225c4-56bd-43d9-98fc-d7148e5c5028\"")))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"export_route_filters": [{"action": "permit", "before": "1a15dcab-7e40-45e1-b7c5-bc690eaa9782", "created_at": "2020-11-02T20:40:29.622Z", "ge": 25, "id": "1a15dcab-7e40-45e1-b7c5-bc690eaa9782", "le": 30, "prefix": "192.168.100.0/24", "updated_at": "2020-11-02T20:40:29.622Z"}]}`)
				}))
			})
			It(`Invoke ReplaceGatewayExportRouteFilters successfully with retries`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())
				directLinkService.EnableRetries(0, 0)

				// Construct an instance of the GatewayTemplateRouteFilter model
				gatewayTemplateRouteFilterModel := new(directlinkv1.GatewayTemplateRouteFilter)
				gatewayTemplateRouteFilterModel.Action = core.StringPtr("permit")
				gatewayTemplateRouteFilterModel.Ge = core.Int64Ptr(int64(25))
				gatewayTemplateRouteFilterModel.Le = core.Int64Ptr(int64(30))
				gatewayTemplateRouteFilterModel.Prefix = core.StringPtr("192.168.100.0/24")

				// Construct an instance of the ReplaceGatewayExportRouteFiltersOptions model
				replaceGatewayExportRouteFiltersOptionsModel := new(directlinkv1.ReplaceGatewayExportRouteFiltersOptions)
				replaceGatewayExportRouteFiltersOptionsModel.GatewayID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				replaceGatewayExportRouteFiltersOptionsModel.IfMatch = core.StringPtr("W/\"96d225c4-56bd-43d9-98fc-d7148e5c5028\"")
				replaceGatewayExportRouteFiltersOptionsModel.ExportRouteFilters = []directlinkv1.GatewayTemplateRouteFilter{*gatewayTemplateRouteFilterModel}
				replaceGatewayExportRouteFiltersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := directLinkService.ReplaceGatewayExportRouteFiltersWithContext(ctx, replaceGatewayExportRouteFiltersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				directLinkService.DisableRetries()
				result, response, operationErr := directLinkService.ReplaceGatewayExportRouteFilters(replaceGatewayExportRouteFiltersOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = directLinkService.ReplaceGatewayExportRouteFiltersWithContext(ctx, replaceGatewayExportRouteFiltersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(replaceGatewayExportRouteFiltersPath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "W/\"96d225c4-56bd-43d9-98fc-d7148e5c5028\"")))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"export_route_filters": [{"action": "permit", "before": "1a15dcab-7e40-45e1-b7c5-bc690eaa9782", "created_at": "2020-11-02T20:40:29.622Z", "ge": 25, "id": "1a15dcab-7e40-45e1-b7c5-bc690eaa9782", "le": 30, "prefix": "192.168.100.0/24", "updated_at": "2020-11-02T20:40:29.622Z"}]}`)
				}))
			})
			It(`Invoke ReplaceGatewayExportRouteFilters successfully`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := directLinkService.ReplaceGatewayExportRouteFilters(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GatewayTemplateRouteFilter model
				gatewayTemplateRouteFilterModel := new(directlinkv1.GatewayTemplateRouteFilter)
				gatewayTemplateRouteFilterModel.Action = core.StringPtr("permit")
				gatewayTemplateRouteFilterModel.Ge = core.Int64Ptr(int64(25))
				gatewayTemplateRouteFilterModel.Le = core.Int64Ptr(int64(30))
				gatewayTemplateRouteFilterModel.Prefix = core.StringPtr("192.168.100.0/24")

				// Construct an instance of the ReplaceGatewayExportRouteFiltersOptions model
				replaceGatewayExportRouteFiltersOptionsModel := new(directlinkv1.ReplaceGatewayExportRouteFiltersOptions)
				replaceGatewayExportRouteFiltersOptionsModel.GatewayID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				replaceGatewayExportRouteFiltersOptionsModel.IfMatch = core.StringPtr("W/\"96d225c4-56bd-43d9-98fc-d7148e5c5028\"")
				replaceGatewayExportRouteFiltersOptionsModel.ExportRouteFilters = []directlinkv1.GatewayTemplateRouteFilter{*gatewayTemplateRouteFilterModel}
				replaceGatewayExportRouteFiltersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = directLinkService.ReplaceGatewayExportRouteFilters(replaceGatewayExportRouteFiltersOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ReplaceGatewayExportRouteFilters with error: Operation validation and request error`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Construct an instance of the GatewayTemplateRouteFilter model
				gatewayTemplateRouteFilterModel := new(directlinkv1.GatewayTemplateRouteFilter)
				gatewayTemplateRouteFilterModel.Action = core.StringPtr("permit")
				gatewayTemplateRouteFilterModel.Ge = core.Int64Ptr(int64(25))
				gatewayTemplateRouteFilterModel.Le = core.Int64Ptr(int64(30))
				gatewayTemplateRouteFilterModel.Prefix = core.StringPtr("192.168.100.0/24")

				// Construct an instance of the ReplaceGatewayExportRouteFiltersOptions model
				replaceGatewayExportRouteFiltersOptionsModel := new(directlinkv1.ReplaceGatewayExportRouteFiltersOptions)
				replaceGatewayExportRouteFiltersOptionsModel.GatewayID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				replaceGatewayExportRouteFiltersOptionsModel.IfMatch = core.StringPtr("W/\"96d225c4-56bd-43d9-98fc-d7148e5c5028\"")
				replaceGatewayExportRouteFiltersOptionsModel.ExportRouteFilters = []directlinkv1.GatewayTemplateRouteFilter{*gatewayTemplateRouteFilterModel}
				replaceGatewayExportRouteFiltersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := directLinkService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := directLinkService.ReplaceGatewayExportRouteFilters(replaceGatewayExportRouteFiltersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ReplaceGatewayExportRouteFiltersOptions model with no property values
				replaceGatewayExportRouteFiltersOptionsModelNew := new(directlinkv1.ReplaceGatewayExportRouteFiltersOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = directLinkService.ReplaceGatewayExportRouteFilters(replaceGatewayExportRouteFiltersOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(201)
				}))
			})
			It(`Invoke ReplaceGatewayExportRouteFilters successfully`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Construct an instance of the GatewayTemplateRouteFilter model
				gatewayTemplateRouteFilterModel := new(directlinkv1.GatewayTemplateRouteFilter)
				gatewayTemplateRouteFilterModel.Action = core.StringPtr("permit")
				gatewayTemplateRouteFilterModel.Ge = core.Int64Ptr(int64(25))
				gatewayTemplateRouteFilterModel.Le = core.Int64Ptr(int64(30))
				gatewayTemplateRouteFilterModel.Prefix = core.StringPtr("192.168.100.0/24")

				// Construct an instance of the ReplaceGatewayExportRouteFiltersOptions model
				replaceGatewayExportRouteFiltersOptionsModel := new(directlinkv1.ReplaceGatewayExportRouteFiltersOptions)
				replaceGatewayExportRouteFiltersOptionsModel.GatewayID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				replaceGatewayExportRouteFiltersOptionsModel.IfMatch = core.StringPtr("W/\"96d225c4-56bd-43d9-98fc-d7148e5c5028\"")
				replaceGatewayExportRouteFiltersOptionsModel.ExportRouteFilters = []directlinkv1.GatewayTemplateRouteFilter{*gatewayTemplateRouteFilterModel}
				replaceGatewayExportRouteFiltersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := directLinkService.ReplaceGatewayExportRouteFilters(replaceGatewayExportRouteFiltersOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteGatewayExportRouteFilter(deleteGatewayExportRouteFilterOptions *DeleteGatewayExportRouteFilterOptions)`, func() {
		version := "testString"
		deleteGatewayExportRouteFilterPath := "/gateways/0a06fb9b-820f-4c44-8a31-77f1f0806d28/export_route_filters/0a06fb9b-820f-4c44-8a31-77f1f0806d28"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteGatewayExportRouteFilterPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteGatewayExportRouteFilter successfully`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := directLinkService.DeleteGatewayExportRouteFilter(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteGatewayExportRouteFilterOptions model
				deleteGatewayExportRouteFilterOptionsModel := new(directlinkv1.DeleteGatewayExportRouteFilterOptions)
				deleteGatewayExportRouteFilterOptionsModel.GatewayID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				deleteGatewayExportRouteFilterOptionsModel.ID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				deleteGatewayExportRouteFilterOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = directLinkService.DeleteGatewayExportRouteFilter(deleteGatewayExportRouteFilterOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteGatewayExportRouteFilter with error: Operation validation and request error`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Construct an instance of the DeleteGatewayExportRouteFilterOptions model
				deleteGatewayExportRouteFilterOptionsModel := new(directlinkv1.DeleteGatewayExportRouteFilterOptions)
				deleteGatewayExportRouteFilterOptionsModel.GatewayID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				deleteGatewayExportRouteFilterOptionsModel.ID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				deleteGatewayExportRouteFilterOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := directLinkService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := directLinkService.DeleteGatewayExportRouteFilter(deleteGatewayExportRouteFilterOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteGatewayExportRouteFilterOptions model with no property values
				deleteGatewayExportRouteFilterOptionsModelNew := new(directlinkv1.DeleteGatewayExportRouteFilterOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = directLinkService.DeleteGatewayExportRouteFilter(deleteGatewayExportRouteFilterOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetGatewayExportRouteFilter(getGatewayExportRouteFilterOptions *GetGatewayExportRouteFilterOptions) - Operation response error`, func() {
		version := "testString"
		getGatewayExportRouteFilterPath := "/gateways/0a06fb9b-820f-4c44-8a31-77f1f0806d28/export_route_filters/0a06fb9b-820f-4c44-8a31-77f1f0806d28"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getGatewayExportRouteFilterPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetGatewayExportRouteFilter with error: Operation response processing error`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Construct an instance of the GetGatewayExportRouteFilterOptions model
				getGatewayExportRouteFilterOptionsModel := new(directlinkv1.GetGatewayExportRouteFilterOptions)
				getGatewayExportRouteFilterOptionsModel.GatewayID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				getGatewayExportRouteFilterOptionsModel.ID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				getGatewayExportRouteFilterOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := directLinkService.GetGatewayExportRouteFilter(getGatewayExportRouteFilterOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				directLinkService.EnableRetries(0, 0)
				result, response, operationErr = directLinkService.GetGatewayExportRouteFilter(getGatewayExportRouteFilterOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetGatewayExportRouteFilter(getGatewayExportRouteFilterOptions *GetGatewayExportRouteFilterOptions)`, func() {
		version := "testString"
		getGatewayExportRouteFilterPath := "/gateways/0a06fb9b-820f-4c44-8a31-77f1f0806d28/export_route_filters/0a06fb9b-820f-4c44-8a31-77f1f0806d28"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getGatewayExportRouteFilterPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"action": "permit", "before": "1a15dcab-7e40-45e1-b7c5-bc690eaa9782", "created_at": "2020-11-02T20:40:29.622Z", "ge": 25, "id": "1a15dcab-7e40-45e1-b7c5-bc690eaa9782", "le": 30, "prefix": "192.168.100.0/24", "updated_at": "2020-11-02T20:40:29.622Z"}`)
				}))
			})
			It(`Invoke GetGatewayExportRouteFilter successfully with retries`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())
				directLinkService.EnableRetries(0, 0)

				// Construct an instance of the GetGatewayExportRouteFilterOptions model
				getGatewayExportRouteFilterOptionsModel := new(directlinkv1.GetGatewayExportRouteFilterOptions)
				getGatewayExportRouteFilterOptionsModel.GatewayID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				getGatewayExportRouteFilterOptionsModel.ID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				getGatewayExportRouteFilterOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := directLinkService.GetGatewayExportRouteFilterWithContext(ctx, getGatewayExportRouteFilterOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				directLinkService.DisableRetries()
				result, response, operationErr := directLinkService.GetGatewayExportRouteFilter(getGatewayExportRouteFilterOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = directLinkService.GetGatewayExportRouteFilterWithContext(ctx, getGatewayExportRouteFilterOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getGatewayExportRouteFilterPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"action": "permit", "before": "1a15dcab-7e40-45e1-b7c5-bc690eaa9782", "created_at": "2020-11-02T20:40:29.622Z", "ge": 25, "id": "1a15dcab-7e40-45e1-b7c5-bc690eaa9782", "le": 30, "prefix": "192.168.100.0/24", "updated_at": "2020-11-02T20:40:29.622Z"}`)
				}))
			})
			It(`Invoke GetGatewayExportRouteFilter successfully`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := directLinkService.GetGatewayExportRouteFilter(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetGatewayExportRouteFilterOptions model
				getGatewayExportRouteFilterOptionsModel := new(directlinkv1.GetGatewayExportRouteFilterOptions)
				getGatewayExportRouteFilterOptionsModel.GatewayID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				getGatewayExportRouteFilterOptionsModel.ID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				getGatewayExportRouteFilterOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = directLinkService.GetGatewayExportRouteFilter(getGatewayExportRouteFilterOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetGatewayExportRouteFilter with error: Operation validation and request error`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Construct an instance of the GetGatewayExportRouteFilterOptions model
				getGatewayExportRouteFilterOptionsModel := new(directlinkv1.GetGatewayExportRouteFilterOptions)
				getGatewayExportRouteFilterOptionsModel.GatewayID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				getGatewayExportRouteFilterOptionsModel.ID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				getGatewayExportRouteFilterOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := directLinkService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := directLinkService.GetGatewayExportRouteFilter(getGatewayExportRouteFilterOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetGatewayExportRouteFilterOptions model with no property values
				getGatewayExportRouteFilterOptionsModelNew := new(directlinkv1.GetGatewayExportRouteFilterOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = directLinkService.GetGatewayExportRouteFilter(getGatewayExportRouteFilterOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetGatewayExportRouteFilter successfully`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Construct an instance of the GetGatewayExportRouteFilterOptions model
				getGatewayExportRouteFilterOptionsModel := new(directlinkv1.GetGatewayExportRouteFilterOptions)
				getGatewayExportRouteFilterOptionsModel.GatewayID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				getGatewayExportRouteFilterOptionsModel.ID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				getGatewayExportRouteFilterOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := directLinkService.GetGatewayExportRouteFilter(getGatewayExportRouteFilterOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateGatewayExportRouteFilter(updateGatewayExportRouteFilterOptions *UpdateGatewayExportRouteFilterOptions) - Operation response error`, func() {
		version := "testString"
		updateGatewayExportRouteFilterPath := "/gateways/0a06fb9b-820f-4c44-8a31-77f1f0806d28/export_route_filters/0a06fb9b-820f-4c44-8a31-77f1f0806d28"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateGatewayExportRouteFilterPath))
					Expect(req.Method).To(Equal("PATCH"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateGatewayExportRouteFilter with error: Operation response processing error`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Construct an instance of the UpdateRouteFilterTemplate model
				updateRouteFilterTemplateModel := new(directlinkv1.UpdateRouteFilterTemplate)
				updateRouteFilterTemplateModel.Action = core.StringPtr("permit")
				updateRouteFilterTemplateModel.Before = core.StringPtr("1a15dcab-7e40-45e1-b7c5-bc690eaa9782")
				updateRouteFilterTemplateModel.Ge = core.Int64Ptr(int64(25))
				updateRouteFilterTemplateModel.Le = core.Int64Ptr(int64(30))
				updateRouteFilterTemplateModel.Prefix = core.StringPtr("192.168.100.0/24")
				updateRouteFilterTemplateModelAsPatch, asPatchErr := updateRouteFilterTemplateModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateGatewayExportRouteFilterOptions model
				updateGatewayExportRouteFilterOptionsModel := new(directlinkv1.UpdateGatewayExportRouteFilterOptions)
				updateGatewayExportRouteFilterOptionsModel.GatewayID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				updateGatewayExportRouteFilterOptionsModel.ID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				updateGatewayExportRouteFilterOptionsModel.UpdateRouteFilterTemplatePatch = updateRouteFilterTemplateModelAsPatch
				updateGatewayExportRouteFilterOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := directLinkService.UpdateGatewayExportRouteFilter(updateGatewayExportRouteFilterOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				directLinkService.EnableRetries(0, 0)
				result, response, operationErr = directLinkService.UpdateGatewayExportRouteFilter(updateGatewayExportRouteFilterOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateGatewayExportRouteFilter(updateGatewayExportRouteFilterOptions *UpdateGatewayExportRouteFilterOptions)`, func() {
		version := "testString"
		updateGatewayExportRouteFilterPath := "/gateways/0a06fb9b-820f-4c44-8a31-77f1f0806d28/export_route_filters/0a06fb9b-820f-4c44-8a31-77f1f0806d28"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateGatewayExportRouteFilterPath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"action": "permit", "before": "1a15dcab-7e40-45e1-b7c5-bc690eaa9782", "created_at": "2020-11-02T20:40:29.622Z", "ge": 25, "id": "1a15dcab-7e40-45e1-b7c5-bc690eaa9782", "le": 30, "prefix": "192.168.100.0/24", "updated_at": "2020-11-02T20:40:29.622Z"}`)
				}))
			})
			It(`Invoke UpdateGatewayExportRouteFilter successfully with retries`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())
				directLinkService.EnableRetries(0, 0)

				// Construct an instance of the UpdateRouteFilterTemplate model
				updateRouteFilterTemplateModel := new(directlinkv1.UpdateRouteFilterTemplate)
				updateRouteFilterTemplateModel.Action = core.StringPtr("permit")
				updateRouteFilterTemplateModel.Before = core.StringPtr("1a15dcab-7e40-45e1-b7c5-bc690eaa9782")
				updateRouteFilterTemplateModel.Ge = core.Int64Ptr(int64(25))
				updateRouteFilterTemplateModel.Le = core.Int64Ptr(int64(30))
				updateRouteFilterTemplateModel.Prefix = core.StringPtr("192.168.100.0/24")
				updateRouteFilterTemplateModelAsPatch, asPatchErr := updateRouteFilterTemplateModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateGatewayExportRouteFilterOptions model
				updateGatewayExportRouteFilterOptionsModel := new(directlinkv1.UpdateGatewayExportRouteFilterOptions)
				updateGatewayExportRouteFilterOptionsModel.GatewayID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				updateGatewayExportRouteFilterOptionsModel.ID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				updateGatewayExportRouteFilterOptionsModel.UpdateRouteFilterTemplatePatch = updateRouteFilterTemplateModelAsPatch
				updateGatewayExportRouteFilterOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := directLinkService.UpdateGatewayExportRouteFilterWithContext(ctx, updateGatewayExportRouteFilterOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				directLinkService.DisableRetries()
				result, response, operationErr := directLinkService.UpdateGatewayExportRouteFilter(updateGatewayExportRouteFilterOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = directLinkService.UpdateGatewayExportRouteFilterWithContext(ctx, updateGatewayExportRouteFilterOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateGatewayExportRouteFilterPath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"action": "permit", "before": "1a15dcab-7e40-45e1-b7c5-bc690eaa9782", "created_at": "2020-11-02T20:40:29.622Z", "ge": 25, "id": "1a15dcab-7e40-45e1-b7c5-bc690eaa9782", "le": 30, "prefix": "192.168.100.0/24", "updated_at": "2020-11-02T20:40:29.622Z"}`)
				}))
			})
			It(`Invoke UpdateGatewayExportRouteFilter successfully`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := directLinkService.UpdateGatewayExportRouteFilter(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateRouteFilterTemplate model
				updateRouteFilterTemplateModel := new(directlinkv1.UpdateRouteFilterTemplate)
				updateRouteFilterTemplateModel.Action = core.StringPtr("permit")
				updateRouteFilterTemplateModel.Before = core.StringPtr("1a15dcab-7e40-45e1-b7c5-bc690eaa9782")
				updateRouteFilterTemplateModel.Ge = core.Int64Ptr(int64(25))
				updateRouteFilterTemplateModel.Le = core.Int64Ptr(int64(30))
				updateRouteFilterTemplateModel.Prefix = core.StringPtr("192.168.100.0/24")
				updateRouteFilterTemplateModelAsPatch, asPatchErr := updateRouteFilterTemplateModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateGatewayExportRouteFilterOptions model
				updateGatewayExportRouteFilterOptionsModel := new(directlinkv1.UpdateGatewayExportRouteFilterOptions)
				updateGatewayExportRouteFilterOptionsModel.GatewayID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				updateGatewayExportRouteFilterOptionsModel.ID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				updateGatewayExportRouteFilterOptionsModel.UpdateRouteFilterTemplatePatch = updateRouteFilterTemplateModelAsPatch
				updateGatewayExportRouteFilterOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = directLinkService.UpdateGatewayExportRouteFilter(updateGatewayExportRouteFilterOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateGatewayExportRouteFilter with error: Operation validation and request error`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Construct an instance of the UpdateRouteFilterTemplate model
				updateRouteFilterTemplateModel := new(directlinkv1.UpdateRouteFilterTemplate)
				updateRouteFilterTemplateModel.Action = core.StringPtr("permit")
				updateRouteFilterTemplateModel.Before = core.StringPtr("1a15dcab-7e40-45e1-b7c5-bc690eaa9782")
				updateRouteFilterTemplateModel.Ge = core.Int64Ptr(int64(25))
				updateRouteFilterTemplateModel.Le = core.Int64Ptr(int64(30))
				updateRouteFilterTemplateModel.Prefix = core.StringPtr("192.168.100.0/24")
				updateRouteFilterTemplateModelAsPatch, asPatchErr := updateRouteFilterTemplateModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateGatewayExportRouteFilterOptions model
				updateGatewayExportRouteFilterOptionsModel := new(directlinkv1.UpdateGatewayExportRouteFilterOptions)
				updateGatewayExportRouteFilterOptionsModel.GatewayID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				updateGatewayExportRouteFilterOptionsModel.ID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				updateGatewayExportRouteFilterOptionsModel.UpdateRouteFilterTemplatePatch = updateRouteFilterTemplateModelAsPatch
				updateGatewayExportRouteFilterOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := directLinkService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := directLinkService.UpdateGatewayExportRouteFilter(updateGatewayExportRouteFilterOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateGatewayExportRouteFilterOptions model with no property values
				updateGatewayExportRouteFilterOptionsModelNew := new(directlinkv1.UpdateGatewayExportRouteFilterOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = directLinkService.UpdateGatewayExportRouteFilter(updateGatewayExportRouteFilterOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke UpdateGatewayExportRouteFilter successfully`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Construct an instance of the UpdateRouteFilterTemplate model
				updateRouteFilterTemplateModel := new(directlinkv1.UpdateRouteFilterTemplate)
				updateRouteFilterTemplateModel.Action = core.StringPtr("permit")
				updateRouteFilterTemplateModel.Before = core.StringPtr("1a15dcab-7e40-45e1-b7c5-bc690eaa9782")
				updateRouteFilterTemplateModel.Ge = core.Int64Ptr(int64(25))
				updateRouteFilterTemplateModel.Le = core.Int64Ptr(int64(30))
				updateRouteFilterTemplateModel.Prefix = core.StringPtr("192.168.100.0/24")
				updateRouteFilterTemplateModelAsPatch, asPatchErr := updateRouteFilterTemplateModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateGatewayExportRouteFilterOptions model
				updateGatewayExportRouteFilterOptionsModel := new(directlinkv1.UpdateGatewayExportRouteFilterOptions)
				updateGatewayExportRouteFilterOptionsModel.GatewayID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				updateGatewayExportRouteFilterOptionsModel.ID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				updateGatewayExportRouteFilterOptionsModel.UpdateRouteFilterTemplatePatch = updateRouteFilterTemplateModelAsPatch
				updateGatewayExportRouteFilterOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := directLinkService.UpdateGatewayExportRouteFilter(updateGatewayExportRouteFilterOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListGatewayImportRouteFilters(listGatewayImportRouteFiltersOptions *ListGatewayImportRouteFiltersOptions) - Operation response error`, func() {
		version := "testString"
		listGatewayImportRouteFiltersPath := "/gateways/0a06fb9b-820f-4c44-8a31-77f1f0806d28/import_route_filters"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listGatewayImportRouteFiltersPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListGatewayImportRouteFilters with error: Operation response processing error`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Construct an instance of the ListGatewayImportRouteFiltersOptions model
				listGatewayImportRouteFiltersOptionsModel := new(directlinkv1.ListGatewayImportRouteFiltersOptions)
				listGatewayImportRouteFiltersOptionsModel.GatewayID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				listGatewayImportRouteFiltersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := directLinkService.ListGatewayImportRouteFilters(listGatewayImportRouteFiltersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				directLinkService.EnableRetries(0, 0)
				result, response, operationErr = directLinkService.ListGatewayImportRouteFilters(listGatewayImportRouteFiltersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListGatewayImportRouteFilters(listGatewayImportRouteFiltersOptions *ListGatewayImportRouteFiltersOptions)`, func() {
		version := "testString"
		listGatewayImportRouteFiltersPath := "/gateways/0a06fb9b-820f-4c44-8a31-77f1f0806d28/import_route_filters"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listGatewayImportRouteFiltersPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"import_route_filters": [{"action": "permit", "before": "1a15dcab-7e40-45e1-b7c5-bc690eaa9782", "created_at": "2020-11-02T20:40:29.622Z", "ge": 25, "id": "1a15dcab-7e40-45e1-b7c5-bc690eaa9782", "le": 30, "prefix": "192.168.100.0/24", "updated_at": "2020-11-02T20:40:29.622Z"}]}`)
				}))
			})
			It(`Invoke ListGatewayImportRouteFilters successfully with retries`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())
				directLinkService.EnableRetries(0, 0)

				// Construct an instance of the ListGatewayImportRouteFiltersOptions model
				listGatewayImportRouteFiltersOptionsModel := new(directlinkv1.ListGatewayImportRouteFiltersOptions)
				listGatewayImportRouteFiltersOptionsModel.GatewayID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				listGatewayImportRouteFiltersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := directLinkService.ListGatewayImportRouteFiltersWithContext(ctx, listGatewayImportRouteFiltersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				directLinkService.DisableRetries()
				result, response, operationErr := directLinkService.ListGatewayImportRouteFilters(listGatewayImportRouteFiltersOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = directLinkService.ListGatewayImportRouteFiltersWithContext(ctx, listGatewayImportRouteFiltersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listGatewayImportRouteFiltersPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"import_route_filters": [{"action": "permit", "before": "1a15dcab-7e40-45e1-b7c5-bc690eaa9782", "created_at": "2020-11-02T20:40:29.622Z", "ge": 25, "id": "1a15dcab-7e40-45e1-b7c5-bc690eaa9782", "le": 30, "prefix": "192.168.100.0/24", "updated_at": "2020-11-02T20:40:29.622Z"}]}`)
				}))
			})
			It(`Invoke ListGatewayImportRouteFilters successfully`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := directLinkService.ListGatewayImportRouteFilters(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListGatewayImportRouteFiltersOptions model
				listGatewayImportRouteFiltersOptionsModel := new(directlinkv1.ListGatewayImportRouteFiltersOptions)
				listGatewayImportRouteFiltersOptionsModel.GatewayID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				listGatewayImportRouteFiltersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = directLinkService.ListGatewayImportRouteFilters(listGatewayImportRouteFiltersOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListGatewayImportRouteFilters with error: Operation validation and request error`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Construct an instance of the ListGatewayImportRouteFiltersOptions model
				listGatewayImportRouteFiltersOptionsModel := new(directlinkv1.ListGatewayImportRouteFiltersOptions)
				listGatewayImportRouteFiltersOptionsModel.GatewayID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				listGatewayImportRouteFiltersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := directLinkService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := directLinkService.ListGatewayImportRouteFilters(listGatewayImportRouteFiltersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListGatewayImportRouteFiltersOptions model with no property values
				listGatewayImportRouteFiltersOptionsModelNew := new(directlinkv1.ListGatewayImportRouteFiltersOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = directLinkService.ListGatewayImportRouteFilters(listGatewayImportRouteFiltersOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ListGatewayImportRouteFilters successfully`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Construct an instance of the ListGatewayImportRouteFiltersOptions model
				listGatewayImportRouteFiltersOptionsModel := new(directlinkv1.ListGatewayImportRouteFiltersOptions)
				listGatewayImportRouteFiltersOptionsModel.GatewayID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				listGatewayImportRouteFiltersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := directLinkService.ListGatewayImportRouteFilters(listGatewayImportRouteFiltersOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateGatewayImportRouteFilter(createGatewayImportRouteFilterOptions *CreateGatewayImportRouteFilterOptions) - Operation response error`, func() {
		version := "testString"
		createGatewayImportRouteFilterPath := "/gateways/0a06fb9b-820f-4c44-8a31-77f1f0806d28/import_route_filters"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createGatewayImportRouteFilterPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateGatewayImportRouteFilter with error: Operation response processing error`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Construct an instance of the CreateGatewayImportRouteFilterOptions model
				createGatewayImportRouteFilterOptionsModel := new(directlinkv1.CreateGatewayImportRouteFilterOptions)
				createGatewayImportRouteFilterOptionsModel.GatewayID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				createGatewayImportRouteFilterOptionsModel.Action = core.StringPtr("permit")
				createGatewayImportRouteFilterOptionsModel.Prefix = core.StringPtr("192.168.100.0/24")
				createGatewayImportRouteFilterOptionsModel.Before = core.StringPtr("1a15dcab-7e40-45e1-b7c5-bc690eaa9782")
				createGatewayImportRouteFilterOptionsModel.Ge = core.Int64Ptr(int64(25))
				createGatewayImportRouteFilterOptionsModel.Le = core.Int64Ptr(int64(30))
				createGatewayImportRouteFilterOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := directLinkService.CreateGatewayImportRouteFilter(createGatewayImportRouteFilterOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				directLinkService.EnableRetries(0, 0)
				result, response, operationErr = directLinkService.CreateGatewayImportRouteFilter(createGatewayImportRouteFilterOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateGatewayImportRouteFilter(createGatewayImportRouteFilterOptions *CreateGatewayImportRouteFilterOptions)`, func() {
		version := "testString"
		createGatewayImportRouteFilterPath := "/gateways/0a06fb9b-820f-4c44-8a31-77f1f0806d28/import_route_filters"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createGatewayImportRouteFilterPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"action": "permit", "before": "1a15dcab-7e40-45e1-b7c5-bc690eaa9782", "created_at": "2020-11-02T20:40:29.622Z", "ge": 25, "id": "1a15dcab-7e40-45e1-b7c5-bc690eaa9782", "le": 30, "prefix": "192.168.100.0/24", "updated_at": "2020-11-02T20:40:29.622Z"}`)
				}))
			})
			It(`Invoke CreateGatewayImportRouteFilter successfully with retries`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())
				directLinkService.EnableRetries(0, 0)

				// Construct an instance of the CreateGatewayImportRouteFilterOptions model
				createGatewayImportRouteFilterOptionsModel := new(directlinkv1.CreateGatewayImportRouteFilterOptions)
				createGatewayImportRouteFilterOptionsModel.GatewayID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				createGatewayImportRouteFilterOptionsModel.Action = core.StringPtr("permit")
				createGatewayImportRouteFilterOptionsModel.Prefix = core.StringPtr("192.168.100.0/24")
				createGatewayImportRouteFilterOptionsModel.Before = core.StringPtr("1a15dcab-7e40-45e1-b7c5-bc690eaa9782")
				createGatewayImportRouteFilterOptionsModel.Ge = core.Int64Ptr(int64(25))
				createGatewayImportRouteFilterOptionsModel.Le = core.Int64Ptr(int64(30))
				createGatewayImportRouteFilterOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := directLinkService.CreateGatewayImportRouteFilterWithContext(ctx, createGatewayImportRouteFilterOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				directLinkService.DisableRetries()
				result, response, operationErr := directLinkService.CreateGatewayImportRouteFilter(createGatewayImportRouteFilterOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = directLinkService.CreateGatewayImportRouteFilterWithContext(ctx, createGatewayImportRouteFilterOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createGatewayImportRouteFilterPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"action": "permit", "before": "1a15dcab-7e40-45e1-b7c5-bc690eaa9782", "created_at": "2020-11-02T20:40:29.622Z", "ge": 25, "id": "1a15dcab-7e40-45e1-b7c5-bc690eaa9782", "le": 30, "prefix": "192.168.100.0/24", "updated_at": "2020-11-02T20:40:29.622Z"}`)
				}))
			})
			It(`Invoke CreateGatewayImportRouteFilter successfully`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := directLinkService.CreateGatewayImportRouteFilter(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateGatewayImportRouteFilterOptions model
				createGatewayImportRouteFilterOptionsModel := new(directlinkv1.CreateGatewayImportRouteFilterOptions)
				createGatewayImportRouteFilterOptionsModel.GatewayID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				createGatewayImportRouteFilterOptionsModel.Action = core.StringPtr("permit")
				createGatewayImportRouteFilterOptionsModel.Prefix = core.StringPtr("192.168.100.0/24")
				createGatewayImportRouteFilterOptionsModel.Before = core.StringPtr("1a15dcab-7e40-45e1-b7c5-bc690eaa9782")
				createGatewayImportRouteFilterOptionsModel.Ge = core.Int64Ptr(int64(25))
				createGatewayImportRouteFilterOptionsModel.Le = core.Int64Ptr(int64(30))
				createGatewayImportRouteFilterOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = directLinkService.CreateGatewayImportRouteFilter(createGatewayImportRouteFilterOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateGatewayImportRouteFilter with error: Operation validation and request error`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Construct an instance of the CreateGatewayImportRouteFilterOptions model
				createGatewayImportRouteFilterOptionsModel := new(directlinkv1.CreateGatewayImportRouteFilterOptions)
				createGatewayImportRouteFilterOptionsModel.GatewayID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				createGatewayImportRouteFilterOptionsModel.Action = core.StringPtr("permit")
				createGatewayImportRouteFilterOptionsModel.Prefix = core.StringPtr("192.168.100.0/24")
				createGatewayImportRouteFilterOptionsModel.Before = core.StringPtr("1a15dcab-7e40-45e1-b7c5-bc690eaa9782")
				createGatewayImportRouteFilterOptionsModel.Ge = core.Int64Ptr(int64(25))
				createGatewayImportRouteFilterOptionsModel.Le = core.Int64Ptr(int64(30))
				createGatewayImportRouteFilterOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := directLinkService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := directLinkService.CreateGatewayImportRouteFilter(createGatewayImportRouteFilterOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateGatewayImportRouteFilterOptions model with no property values
				createGatewayImportRouteFilterOptionsModelNew := new(directlinkv1.CreateGatewayImportRouteFilterOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = directLinkService.CreateGatewayImportRouteFilter(createGatewayImportRouteFilterOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(201)
				}))
			})
			It(`Invoke CreateGatewayImportRouteFilter successfully`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Construct an instance of the CreateGatewayImportRouteFilterOptions model
				createGatewayImportRouteFilterOptionsModel := new(directlinkv1.CreateGatewayImportRouteFilterOptions)
				createGatewayImportRouteFilterOptionsModel.GatewayID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				createGatewayImportRouteFilterOptionsModel.Action = core.StringPtr("permit")
				createGatewayImportRouteFilterOptionsModel.Prefix = core.StringPtr("192.168.100.0/24")
				createGatewayImportRouteFilterOptionsModel.Before = core.StringPtr("1a15dcab-7e40-45e1-b7c5-bc690eaa9782")
				createGatewayImportRouteFilterOptionsModel.Ge = core.Int64Ptr(int64(25))
				createGatewayImportRouteFilterOptionsModel.Le = core.Int64Ptr(int64(30))
				createGatewayImportRouteFilterOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := directLinkService.CreateGatewayImportRouteFilter(createGatewayImportRouteFilterOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ReplaceGatewayImportRouteFilters(replaceGatewayImportRouteFiltersOptions *ReplaceGatewayImportRouteFiltersOptions) - Operation response error`, func() {
		version := "testString"
		replaceGatewayImportRouteFiltersPath := "/gateways/0a06fb9b-820f-4c44-8a31-77f1f0806d28/import_route_filters"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(replaceGatewayImportRouteFiltersPath))
					Expect(req.Method).To(Equal("PUT"))
					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "W/\"96d225c4-56bd-43d9-98fc-d7148e5c5028\"")))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ReplaceGatewayImportRouteFilters with error: Operation response processing error`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Construct an instance of the GatewayTemplateRouteFilter model
				gatewayTemplateRouteFilterModel := new(directlinkv1.GatewayTemplateRouteFilter)
				gatewayTemplateRouteFilterModel.Action = core.StringPtr("permit")
				gatewayTemplateRouteFilterModel.Ge = core.Int64Ptr(int64(25))
				gatewayTemplateRouteFilterModel.Le = core.Int64Ptr(int64(30))
				gatewayTemplateRouteFilterModel.Prefix = core.StringPtr("192.168.100.0/24")

				// Construct an instance of the ReplaceGatewayImportRouteFiltersOptions model
				replaceGatewayImportRouteFiltersOptionsModel := new(directlinkv1.ReplaceGatewayImportRouteFiltersOptions)
				replaceGatewayImportRouteFiltersOptionsModel.GatewayID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				replaceGatewayImportRouteFiltersOptionsModel.IfMatch = core.StringPtr("W/\"96d225c4-56bd-43d9-98fc-d7148e5c5028\"")
				replaceGatewayImportRouteFiltersOptionsModel.ImportRouteFilters = []directlinkv1.GatewayTemplateRouteFilter{*gatewayTemplateRouteFilterModel}
				replaceGatewayImportRouteFiltersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := directLinkService.ReplaceGatewayImportRouteFilters(replaceGatewayImportRouteFiltersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				directLinkService.EnableRetries(0, 0)
				result, response, operationErr = directLinkService.ReplaceGatewayImportRouteFilters(replaceGatewayImportRouteFiltersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ReplaceGatewayImportRouteFilters(replaceGatewayImportRouteFiltersOptions *ReplaceGatewayImportRouteFiltersOptions)`, func() {
		version := "testString"
		replaceGatewayImportRouteFiltersPath := "/gateways/0a06fb9b-820f-4c44-8a31-77f1f0806d28/import_route_filters"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(replaceGatewayImportRouteFiltersPath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "W/\"96d225c4-56bd-43d9-98fc-d7148e5c5028\"")))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"import_route_filters": [{"action": "permit", "before": "1a15dcab-7e40-45e1-b7c5-bc690eaa9782", "created_at": "2020-11-02T20:40:29.622Z", "ge": 25, "id": "1a15dcab-7e40-45e1-b7c5-bc690eaa9782", "le": 30, "prefix": "192.168.100.0/24", "updated_at": "2020-11-02T20:40:29.622Z"}]}`)
				}))
			})
			It(`Invoke ReplaceGatewayImportRouteFilters successfully with retries`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())
				directLinkService.EnableRetries(0, 0)

				// Construct an instance of the GatewayTemplateRouteFilter model
				gatewayTemplateRouteFilterModel := new(directlinkv1.GatewayTemplateRouteFilter)
				gatewayTemplateRouteFilterModel.Action = core.StringPtr("permit")
				gatewayTemplateRouteFilterModel.Ge = core.Int64Ptr(int64(25))
				gatewayTemplateRouteFilterModel.Le = core.Int64Ptr(int64(30))
				gatewayTemplateRouteFilterModel.Prefix = core.StringPtr("192.168.100.0/24")

				// Construct an instance of the ReplaceGatewayImportRouteFiltersOptions model
				replaceGatewayImportRouteFiltersOptionsModel := new(directlinkv1.ReplaceGatewayImportRouteFiltersOptions)
				replaceGatewayImportRouteFiltersOptionsModel.GatewayID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				replaceGatewayImportRouteFiltersOptionsModel.IfMatch = core.StringPtr("W/\"96d225c4-56bd-43d9-98fc-d7148e5c5028\"")
				replaceGatewayImportRouteFiltersOptionsModel.ImportRouteFilters = []directlinkv1.GatewayTemplateRouteFilter{*gatewayTemplateRouteFilterModel}
				replaceGatewayImportRouteFiltersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := directLinkService.ReplaceGatewayImportRouteFiltersWithContext(ctx, replaceGatewayImportRouteFiltersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				directLinkService.DisableRetries()
				result, response, operationErr := directLinkService.ReplaceGatewayImportRouteFilters(replaceGatewayImportRouteFiltersOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = directLinkService.ReplaceGatewayImportRouteFiltersWithContext(ctx, replaceGatewayImportRouteFiltersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(replaceGatewayImportRouteFiltersPath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "W/\"96d225c4-56bd-43d9-98fc-d7148e5c5028\"")))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"import_route_filters": [{"action": "permit", "before": "1a15dcab-7e40-45e1-b7c5-bc690eaa9782", "created_at": "2020-11-02T20:40:29.622Z", "ge": 25, "id": "1a15dcab-7e40-45e1-b7c5-bc690eaa9782", "le": 30, "prefix": "192.168.100.0/24", "updated_at": "2020-11-02T20:40:29.622Z"}]}`)
				}))
			})
			It(`Invoke ReplaceGatewayImportRouteFilters successfully`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := directLinkService.ReplaceGatewayImportRouteFilters(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GatewayTemplateRouteFilter model
				gatewayTemplateRouteFilterModel := new(directlinkv1.GatewayTemplateRouteFilter)
				gatewayTemplateRouteFilterModel.Action = core.StringPtr("permit")
				gatewayTemplateRouteFilterModel.Ge = core.Int64Ptr(int64(25))
				gatewayTemplateRouteFilterModel.Le = core.Int64Ptr(int64(30))
				gatewayTemplateRouteFilterModel.Prefix = core.StringPtr("192.168.100.0/24")

				// Construct an instance of the ReplaceGatewayImportRouteFiltersOptions model
				replaceGatewayImportRouteFiltersOptionsModel := new(directlinkv1.ReplaceGatewayImportRouteFiltersOptions)
				replaceGatewayImportRouteFiltersOptionsModel.GatewayID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				replaceGatewayImportRouteFiltersOptionsModel.IfMatch = core.StringPtr("W/\"96d225c4-56bd-43d9-98fc-d7148e5c5028\"")
				replaceGatewayImportRouteFiltersOptionsModel.ImportRouteFilters = []directlinkv1.GatewayTemplateRouteFilter{*gatewayTemplateRouteFilterModel}
				replaceGatewayImportRouteFiltersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = directLinkService.ReplaceGatewayImportRouteFilters(replaceGatewayImportRouteFiltersOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ReplaceGatewayImportRouteFilters with error: Operation validation and request error`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Construct an instance of the GatewayTemplateRouteFilter model
				gatewayTemplateRouteFilterModel := new(directlinkv1.GatewayTemplateRouteFilter)
				gatewayTemplateRouteFilterModel.Action = core.StringPtr("permit")
				gatewayTemplateRouteFilterModel.Ge = core.Int64Ptr(int64(25))
				gatewayTemplateRouteFilterModel.Le = core.Int64Ptr(int64(30))
				gatewayTemplateRouteFilterModel.Prefix = core.StringPtr("192.168.100.0/24")

				// Construct an instance of the ReplaceGatewayImportRouteFiltersOptions model
				replaceGatewayImportRouteFiltersOptionsModel := new(directlinkv1.ReplaceGatewayImportRouteFiltersOptions)
				replaceGatewayImportRouteFiltersOptionsModel.GatewayID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				replaceGatewayImportRouteFiltersOptionsModel.IfMatch = core.StringPtr("W/\"96d225c4-56bd-43d9-98fc-d7148e5c5028\"")
				replaceGatewayImportRouteFiltersOptionsModel.ImportRouteFilters = []directlinkv1.GatewayTemplateRouteFilter{*gatewayTemplateRouteFilterModel}
				replaceGatewayImportRouteFiltersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := directLinkService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := directLinkService.ReplaceGatewayImportRouteFilters(replaceGatewayImportRouteFiltersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ReplaceGatewayImportRouteFiltersOptions model with no property values
				replaceGatewayImportRouteFiltersOptionsModelNew := new(directlinkv1.ReplaceGatewayImportRouteFiltersOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = directLinkService.ReplaceGatewayImportRouteFilters(replaceGatewayImportRouteFiltersOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(201)
				}))
			})
			It(`Invoke ReplaceGatewayImportRouteFilters successfully`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Construct an instance of the GatewayTemplateRouteFilter model
				gatewayTemplateRouteFilterModel := new(directlinkv1.GatewayTemplateRouteFilter)
				gatewayTemplateRouteFilterModel.Action = core.StringPtr("permit")
				gatewayTemplateRouteFilterModel.Ge = core.Int64Ptr(int64(25))
				gatewayTemplateRouteFilterModel.Le = core.Int64Ptr(int64(30))
				gatewayTemplateRouteFilterModel.Prefix = core.StringPtr("192.168.100.0/24")

				// Construct an instance of the ReplaceGatewayImportRouteFiltersOptions model
				replaceGatewayImportRouteFiltersOptionsModel := new(directlinkv1.ReplaceGatewayImportRouteFiltersOptions)
				replaceGatewayImportRouteFiltersOptionsModel.GatewayID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				replaceGatewayImportRouteFiltersOptionsModel.IfMatch = core.StringPtr("W/\"96d225c4-56bd-43d9-98fc-d7148e5c5028\"")
				replaceGatewayImportRouteFiltersOptionsModel.ImportRouteFilters = []directlinkv1.GatewayTemplateRouteFilter{*gatewayTemplateRouteFilterModel}
				replaceGatewayImportRouteFiltersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := directLinkService.ReplaceGatewayImportRouteFilters(replaceGatewayImportRouteFiltersOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteGatewayImportRouteFilter(deleteGatewayImportRouteFilterOptions *DeleteGatewayImportRouteFilterOptions)`, func() {
		version := "testString"
		deleteGatewayImportRouteFilterPath := "/gateways/0a06fb9b-820f-4c44-8a31-77f1f0806d28/import_route_filters/0a06fb9b-820f-4c44-8a31-77f1f0806d28"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteGatewayImportRouteFilterPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteGatewayImportRouteFilter successfully`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := directLinkService.DeleteGatewayImportRouteFilter(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteGatewayImportRouteFilterOptions model
				deleteGatewayImportRouteFilterOptionsModel := new(directlinkv1.DeleteGatewayImportRouteFilterOptions)
				deleteGatewayImportRouteFilterOptionsModel.GatewayID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				deleteGatewayImportRouteFilterOptionsModel.ID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				deleteGatewayImportRouteFilterOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = directLinkService.DeleteGatewayImportRouteFilter(deleteGatewayImportRouteFilterOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteGatewayImportRouteFilter with error: Operation validation and request error`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Construct an instance of the DeleteGatewayImportRouteFilterOptions model
				deleteGatewayImportRouteFilterOptionsModel := new(directlinkv1.DeleteGatewayImportRouteFilterOptions)
				deleteGatewayImportRouteFilterOptionsModel.GatewayID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				deleteGatewayImportRouteFilterOptionsModel.ID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				deleteGatewayImportRouteFilterOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := directLinkService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := directLinkService.DeleteGatewayImportRouteFilter(deleteGatewayImportRouteFilterOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteGatewayImportRouteFilterOptions model with no property values
				deleteGatewayImportRouteFilterOptionsModelNew := new(directlinkv1.DeleteGatewayImportRouteFilterOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = directLinkService.DeleteGatewayImportRouteFilter(deleteGatewayImportRouteFilterOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetGatewayImportRouteFilter(getGatewayImportRouteFilterOptions *GetGatewayImportRouteFilterOptions) - Operation response error`, func() {
		version := "testString"
		getGatewayImportRouteFilterPath := "/gateways/0a06fb9b-820f-4c44-8a31-77f1f0806d28/import_route_filters/0a06fb9b-820f-4c44-8a31-77f1f0806d28"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getGatewayImportRouteFilterPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetGatewayImportRouteFilter with error: Operation response processing error`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Construct an instance of the GetGatewayImportRouteFilterOptions model
				getGatewayImportRouteFilterOptionsModel := new(directlinkv1.GetGatewayImportRouteFilterOptions)
				getGatewayImportRouteFilterOptionsModel.GatewayID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				getGatewayImportRouteFilterOptionsModel.ID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				getGatewayImportRouteFilterOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := directLinkService.GetGatewayImportRouteFilter(getGatewayImportRouteFilterOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				directLinkService.EnableRetries(0, 0)
				result, response, operationErr = directLinkService.GetGatewayImportRouteFilter(getGatewayImportRouteFilterOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetGatewayImportRouteFilter(getGatewayImportRouteFilterOptions *GetGatewayImportRouteFilterOptions)`, func() {
		version := "testString"
		getGatewayImportRouteFilterPath := "/gateways/0a06fb9b-820f-4c44-8a31-77f1f0806d28/import_route_filters/0a06fb9b-820f-4c44-8a31-77f1f0806d28"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getGatewayImportRouteFilterPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"action": "permit", "before": "1a15dcab-7e40-45e1-b7c5-bc690eaa9782", "created_at": "2020-11-02T20:40:29.622Z", "ge": 25, "id": "1a15dcab-7e40-45e1-b7c5-bc690eaa9782", "le": 30, "prefix": "192.168.100.0/24", "updated_at": "2020-11-02T20:40:29.622Z"}`)
				}))
			})
			It(`Invoke GetGatewayImportRouteFilter successfully with retries`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())
				directLinkService.EnableRetries(0, 0)

				// Construct an instance of the GetGatewayImportRouteFilterOptions model
				getGatewayImportRouteFilterOptionsModel := new(directlinkv1.GetGatewayImportRouteFilterOptions)
				getGatewayImportRouteFilterOptionsModel.GatewayID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				getGatewayImportRouteFilterOptionsModel.ID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				getGatewayImportRouteFilterOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := directLinkService.GetGatewayImportRouteFilterWithContext(ctx, getGatewayImportRouteFilterOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				directLinkService.DisableRetries()
				result, response, operationErr := directLinkService.GetGatewayImportRouteFilter(getGatewayImportRouteFilterOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = directLinkService.GetGatewayImportRouteFilterWithContext(ctx, getGatewayImportRouteFilterOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getGatewayImportRouteFilterPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"action": "permit", "before": "1a15dcab-7e40-45e1-b7c5-bc690eaa9782", "created_at": "2020-11-02T20:40:29.622Z", "ge": 25, "id": "1a15dcab-7e40-45e1-b7c5-bc690eaa9782", "le": 30, "prefix": "192.168.100.0/24", "updated_at": "2020-11-02T20:40:29.622Z"}`)
				}))
			})
			It(`Invoke GetGatewayImportRouteFilter successfully`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := directLinkService.GetGatewayImportRouteFilter(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetGatewayImportRouteFilterOptions model
				getGatewayImportRouteFilterOptionsModel := new(directlinkv1.GetGatewayImportRouteFilterOptions)
				getGatewayImportRouteFilterOptionsModel.GatewayID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				getGatewayImportRouteFilterOptionsModel.ID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				getGatewayImportRouteFilterOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = directLinkService.GetGatewayImportRouteFilter(getGatewayImportRouteFilterOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetGatewayImportRouteFilter with error: Operation validation and request error`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Construct an instance of the GetGatewayImportRouteFilterOptions model
				getGatewayImportRouteFilterOptionsModel := new(directlinkv1.GetGatewayImportRouteFilterOptions)
				getGatewayImportRouteFilterOptionsModel.GatewayID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				getGatewayImportRouteFilterOptionsModel.ID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				getGatewayImportRouteFilterOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := directLinkService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := directLinkService.GetGatewayImportRouteFilter(getGatewayImportRouteFilterOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetGatewayImportRouteFilterOptions model with no property values
				getGatewayImportRouteFilterOptionsModelNew := new(directlinkv1.GetGatewayImportRouteFilterOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = directLinkService.GetGatewayImportRouteFilter(getGatewayImportRouteFilterOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetGatewayImportRouteFilter successfully`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Construct an instance of the GetGatewayImportRouteFilterOptions model
				getGatewayImportRouteFilterOptionsModel := new(directlinkv1.GetGatewayImportRouteFilterOptions)
				getGatewayImportRouteFilterOptionsModel.GatewayID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				getGatewayImportRouteFilterOptionsModel.ID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				getGatewayImportRouteFilterOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := directLinkService.GetGatewayImportRouteFilter(getGatewayImportRouteFilterOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateGatewayImportRouteFilter(updateGatewayImportRouteFilterOptions *UpdateGatewayImportRouteFilterOptions) - Operation response error`, func() {
		version := "testString"
		updateGatewayImportRouteFilterPath := "/gateways/0a06fb9b-820f-4c44-8a31-77f1f0806d28/import_route_filters/0a06fb9b-820f-4c44-8a31-77f1f0806d28"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateGatewayImportRouteFilterPath))
					Expect(req.Method).To(Equal("PATCH"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateGatewayImportRouteFilter with error: Operation response processing error`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Construct an instance of the UpdateRouteFilterTemplate model
				updateRouteFilterTemplateModel := new(directlinkv1.UpdateRouteFilterTemplate)
				updateRouteFilterTemplateModel.Action = core.StringPtr("permit")
				updateRouteFilterTemplateModel.Before = core.StringPtr("1a15dcab-7e40-45e1-b7c5-bc690eaa9782")
				updateRouteFilterTemplateModel.Ge = core.Int64Ptr(int64(25))
				updateRouteFilterTemplateModel.Le = core.Int64Ptr(int64(30))
				updateRouteFilterTemplateModel.Prefix = core.StringPtr("192.168.100.0/24")
				updateRouteFilterTemplateModelAsPatch, asPatchErr := updateRouteFilterTemplateModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateGatewayImportRouteFilterOptions model
				updateGatewayImportRouteFilterOptionsModel := new(directlinkv1.UpdateGatewayImportRouteFilterOptions)
				updateGatewayImportRouteFilterOptionsModel.GatewayID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				updateGatewayImportRouteFilterOptionsModel.ID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				updateGatewayImportRouteFilterOptionsModel.UpdateRouteFilterTemplatePatch = updateRouteFilterTemplateModelAsPatch
				updateGatewayImportRouteFilterOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := directLinkService.UpdateGatewayImportRouteFilter(updateGatewayImportRouteFilterOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				directLinkService.EnableRetries(0, 0)
				result, response, operationErr = directLinkService.UpdateGatewayImportRouteFilter(updateGatewayImportRouteFilterOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateGatewayImportRouteFilter(updateGatewayImportRouteFilterOptions *UpdateGatewayImportRouteFilterOptions)`, func() {
		version := "testString"
		updateGatewayImportRouteFilterPath := "/gateways/0a06fb9b-820f-4c44-8a31-77f1f0806d28/import_route_filters/0a06fb9b-820f-4c44-8a31-77f1f0806d28"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateGatewayImportRouteFilterPath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"action": "permit", "before": "1a15dcab-7e40-45e1-b7c5-bc690eaa9782", "created_at": "2020-11-02T20:40:29.622Z", "ge": 25, "id": "1a15dcab-7e40-45e1-b7c5-bc690eaa9782", "le": 30, "prefix": "192.168.100.0/24", "updated_at": "2020-11-02T20:40:29.622Z"}`)
				}))
			})
			It(`Invoke UpdateGatewayImportRouteFilter successfully with retries`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())
				directLinkService.EnableRetries(0, 0)

				// Construct an instance of the UpdateRouteFilterTemplate model
				updateRouteFilterTemplateModel := new(directlinkv1.UpdateRouteFilterTemplate)
				updateRouteFilterTemplateModel.Action = core.StringPtr("permit")
				updateRouteFilterTemplateModel.Before = core.StringPtr("1a15dcab-7e40-45e1-b7c5-bc690eaa9782")
				updateRouteFilterTemplateModel.Ge = core.Int64Ptr(int64(25))
				updateRouteFilterTemplateModel.Le = core.Int64Ptr(int64(30))
				updateRouteFilterTemplateModel.Prefix = core.StringPtr("192.168.100.0/24")
				updateRouteFilterTemplateModelAsPatch, asPatchErr := updateRouteFilterTemplateModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateGatewayImportRouteFilterOptions model
				updateGatewayImportRouteFilterOptionsModel := new(directlinkv1.UpdateGatewayImportRouteFilterOptions)
				updateGatewayImportRouteFilterOptionsModel.GatewayID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				updateGatewayImportRouteFilterOptionsModel.ID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				updateGatewayImportRouteFilterOptionsModel.UpdateRouteFilterTemplatePatch = updateRouteFilterTemplateModelAsPatch
				updateGatewayImportRouteFilterOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := directLinkService.UpdateGatewayImportRouteFilterWithContext(ctx, updateGatewayImportRouteFilterOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				directLinkService.DisableRetries()
				result, response, operationErr := directLinkService.UpdateGatewayImportRouteFilter(updateGatewayImportRouteFilterOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = directLinkService.UpdateGatewayImportRouteFilterWithContext(ctx, updateGatewayImportRouteFilterOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateGatewayImportRouteFilterPath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"action": "permit", "before": "1a15dcab-7e40-45e1-b7c5-bc690eaa9782", "created_at": "2020-11-02T20:40:29.622Z", "ge": 25, "id": "1a15dcab-7e40-45e1-b7c5-bc690eaa9782", "le": 30, "prefix": "192.168.100.0/24", "updated_at": "2020-11-02T20:40:29.622Z"}`)
				}))
			})
			It(`Invoke UpdateGatewayImportRouteFilter successfully`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := directLinkService.UpdateGatewayImportRouteFilter(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateRouteFilterTemplate model
				updateRouteFilterTemplateModel := new(directlinkv1.UpdateRouteFilterTemplate)
				updateRouteFilterTemplateModel.Action = core.StringPtr("permit")
				updateRouteFilterTemplateModel.Before = core.StringPtr("1a15dcab-7e40-45e1-b7c5-bc690eaa9782")
				updateRouteFilterTemplateModel.Ge = core.Int64Ptr(int64(25))
				updateRouteFilterTemplateModel.Le = core.Int64Ptr(int64(30))
				updateRouteFilterTemplateModel.Prefix = core.StringPtr("192.168.100.0/24")
				updateRouteFilterTemplateModelAsPatch, asPatchErr := updateRouteFilterTemplateModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateGatewayImportRouteFilterOptions model
				updateGatewayImportRouteFilterOptionsModel := new(directlinkv1.UpdateGatewayImportRouteFilterOptions)
				updateGatewayImportRouteFilterOptionsModel.GatewayID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				updateGatewayImportRouteFilterOptionsModel.ID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				updateGatewayImportRouteFilterOptionsModel.UpdateRouteFilterTemplatePatch = updateRouteFilterTemplateModelAsPatch
				updateGatewayImportRouteFilterOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = directLinkService.UpdateGatewayImportRouteFilter(updateGatewayImportRouteFilterOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateGatewayImportRouteFilter with error: Operation validation and request error`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Construct an instance of the UpdateRouteFilterTemplate model
				updateRouteFilterTemplateModel := new(directlinkv1.UpdateRouteFilterTemplate)
				updateRouteFilterTemplateModel.Action = core.StringPtr("permit")
				updateRouteFilterTemplateModel.Before = core.StringPtr("1a15dcab-7e40-45e1-b7c5-bc690eaa9782")
				updateRouteFilterTemplateModel.Ge = core.Int64Ptr(int64(25))
				updateRouteFilterTemplateModel.Le = core.Int64Ptr(int64(30))
				updateRouteFilterTemplateModel.Prefix = core.StringPtr("192.168.100.0/24")
				updateRouteFilterTemplateModelAsPatch, asPatchErr := updateRouteFilterTemplateModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateGatewayImportRouteFilterOptions model
				updateGatewayImportRouteFilterOptionsModel := new(directlinkv1.UpdateGatewayImportRouteFilterOptions)
				updateGatewayImportRouteFilterOptionsModel.GatewayID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				updateGatewayImportRouteFilterOptionsModel.ID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				updateGatewayImportRouteFilterOptionsModel.UpdateRouteFilterTemplatePatch = updateRouteFilterTemplateModelAsPatch
				updateGatewayImportRouteFilterOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := directLinkService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := directLinkService.UpdateGatewayImportRouteFilter(updateGatewayImportRouteFilterOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateGatewayImportRouteFilterOptions model with no property values
				updateGatewayImportRouteFilterOptionsModelNew := new(directlinkv1.UpdateGatewayImportRouteFilterOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = directLinkService.UpdateGatewayImportRouteFilter(updateGatewayImportRouteFilterOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke UpdateGatewayImportRouteFilter successfully`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Construct an instance of the UpdateRouteFilterTemplate model
				updateRouteFilterTemplateModel := new(directlinkv1.UpdateRouteFilterTemplate)
				updateRouteFilterTemplateModel.Action = core.StringPtr("permit")
				updateRouteFilterTemplateModel.Before = core.StringPtr("1a15dcab-7e40-45e1-b7c5-bc690eaa9782")
				updateRouteFilterTemplateModel.Ge = core.Int64Ptr(int64(25))
				updateRouteFilterTemplateModel.Le = core.Int64Ptr(int64(30))
				updateRouteFilterTemplateModel.Prefix = core.StringPtr("192.168.100.0/24")
				updateRouteFilterTemplateModelAsPatch, asPatchErr := updateRouteFilterTemplateModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateGatewayImportRouteFilterOptions model
				updateGatewayImportRouteFilterOptionsModel := new(directlinkv1.UpdateGatewayImportRouteFilterOptions)
				updateGatewayImportRouteFilterOptionsModel.GatewayID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				updateGatewayImportRouteFilterOptionsModel.ID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				updateGatewayImportRouteFilterOptionsModel.UpdateRouteFilterTemplatePatch = updateRouteFilterTemplateModelAsPatch
				updateGatewayImportRouteFilterOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := directLinkService.UpdateGatewayImportRouteFilter(updateGatewayImportRouteFilterOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UnsetGatewayMacsec(unsetGatewayMacsecOptions *UnsetGatewayMacsecOptions)`, func() {
		version := "testString"
		unsetGatewayMacsecPath := "/gateways/0a06fb9b-820f-4c44-8a31-77f1f0806d28/macsec"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(unsetGatewayMacsecPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.WriteHeader(204)
				}))
			})
			It(`Invoke UnsetGatewayMacsec successfully`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := directLinkService.UnsetGatewayMacsec(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the UnsetGatewayMacsecOptions model
				unsetGatewayMacsecOptionsModel := new(directlinkv1.UnsetGatewayMacsecOptions)
				unsetGatewayMacsecOptionsModel.ID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				unsetGatewayMacsecOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = directLinkService.UnsetGatewayMacsec(unsetGatewayMacsecOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke UnsetGatewayMacsec with error: Operation validation and request error`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Construct an instance of the UnsetGatewayMacsecOptions model
				unsetGatewayMacsecOptionsModel := new(directlinkv1.UnsetGatewayMacsecOptions)
				unsetGatewayMacsecOptionsModel.ID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				unsetGatewayMacsecOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := directLinkService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := directLinkService.UnsetGatewayMacsec(unsetGatewayMacsecOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the UnsetGatewayMacsecOptions model with no property values
				unsetGatewayMacsecOptionsModelNew := new(directlinkv1.UnsetGatewayMacsecOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = directLinkService.UnsetGatewayMacsec(unsetGatewayMacsecOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetGatewayMacsec(getGatewayMacsecOptions *GetGatewayMacsecOptions) - Operation response error`, func() {
		version := "testString"
		getGatewayMacsecPath := "/gateways/0a06fb9b-820f-4c44-8a31-77f1f0806d28/macsec"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getGatewayMacsecPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetGatewayMacsec with error: Operation response processing error`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Construct an instance of the GetGatewayMacsecOptions model
				getGatewayMacsecOptionsModel := new(directlinkv1.GetGatewayMacsecOptions)
				getGatewayMacsecOptionsModel.ID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				getGatewayMacsecOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := directLinkService.GetGatewayMacsec(getGatewayMacsecOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				directLinkService.EnableRetries(0, 0)
				result, response, operationErr = directLinkService.GetGatewayMacsec(getGatewayMacsecOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetGatewayMacsec(getGatewayMacsecOptions *GetGatewayMacsecOptions)`, func() {
		version := "testString"
		getGatewayMacsecPath := "/gateways/0a06fb9b-820f-4c44-8a31-77f1f0806d28/macsec"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getGatewayMacsecPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"active": true, "cipher_suite": "gcm_aes_xpn_256", "confidentiality_offset": 0, "created_at": "2020-11-02T20:40:29.622Z", "key_server_priority": 255, "sak_rekey": {"interval": 3600, "mode": "timer"}, "security_policy": "must_secure", "status": "secured", "status_reasons": [{"code": "macsec_cak_failed", "message": "The authentication_key failed configuration.", "more_info": "https://cloud.ibm.com/docs/dl/TODO_ADD_DOCS_LINK"}], "updated_at": "2020-11-02T20:40:29.622Z", "window_size": 512}`)
				}))
			})
			It(`Invoke GetGatewayMacsec successfully with retries`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())
				directLinkService.EnableRetries(0, 0)

				// Construct an instance of the GetGatewayMacsecOptions model
				getGatewayMacsecOptionsModel := new(directlinkv1.GetGatewayMacsecOptions)
				getGatewayMacsecOptionsModel.ID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				getGatewayMacsecOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := directLinkService.GetGatewayMacsecWithContext(ctx, getGatewayMacsecOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				directLinkService.DisableRetries()
				result, response, operationErr := directLinkService.GetGatewayMacsec(getGatewayMacsecOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = directLinkService.GetGatewayMacsecWithContext(ctx, getGatewayMacsecOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getGatewayMacsecPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"active": true, "cipher_suite": "gcm_aes_xpn_256", "confidentiality_offset": 0, "created_at": "2020-11-02T20:40:29.622Z", "key_server_priority": 255, "sak_rekey": {"interval": 3600, "mode": "timer"}, "security_policy": "must_secure", "status": "secured", "status_reasons": [{"code": "macsec_cak_failed", "message": "The authentication_key failed configuration.", "more_info": "https://cloud.ibm.com/docs/dl/TODO_ADD_DOCS_LINK"}], "updated_at": "2020-11-02T20:40:29.622Z", "window_size": 512}`)
				}))
			})
			It(`Invoke GetGatewayMacsec successfully`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := directLinkService.GetGatewayMacsec(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetGatewayMacsecOptions model
				getGatewayMacsecOptionsModel := new(directlinkv1.GetGatewayMacsecOptions)
				getGatewayMacsecOptionsModel.ID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				getGatewayMacsecOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = directLinkService.GetGatewayMacsec(getGatewayMacsecOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetGatewayMacsec with error: Operation validation and request error`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Construct an instance of the GetGatewayMacsecOptions model
				getGatewayMacsecOptionsModel := new(directlinkv1.GetGatewayMacsecOptions)
				getGatewayMacsecOptionsModel.ID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				getGatewayMacsecOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := directLinkService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := directLinkService.GetGatewayMacsec(getGatewayMacsecOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetGatewayMacsecOptions model with no property values
				getGatewayMacsecOptionsModelNew := new(directlinkv1.GetGatewayMacsecOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = directLinkService.GetGatewayMacsec(getGatewayMacsecOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetGatewayMacsec successfully`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Construct an instance of the GetGatewayMacsecOptions model
				getGatewayMacsecOptionsModel := new(directlinkv1.GetGatewayMacsecOptions)
				getGatewayMacsecOptionsModel.ID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				getGatewayMacsecOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := directLinkService.GetGatewayMacsec(getGatewayMacsecOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateGatewayMacsec(updateGatewayMacsecOptions *UpdateGatewayMacsecOptions) - Operation response error`, func() {
		version := "testString"
		updateGatewayMacsecPath := "/gateways/0a06fb9b-820f-4c44-8a31-77f1f0806d28/macsec"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateGatewayMacsecPath))
					Expect(req.Method).To(Equal("PATCH"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateGatewayMacsec with error: Operation response processing error`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Construct an instance of the SakRekeyPatchSakRekeyTimerModePatch model
				sakRekeyPatchModel := new(directlinkv1.SakRekeyPatchSakRekeyTimerModePatch)
				sakRekeyPatchModel.Interval = core.Int64Ptr(int64(3600))
				sakRekeyPatchModel.Mode = core.StringPtr("timer")

				// Construct an instance of the GatewayMacsecPatch model
				gatewayMacsecPatchModel := new(directlinkv1.GatewayMacsecPatch)
				gatewayMacsecPatchModel.Active = core.BoolPtr(true)
				gatewayMacsecPatchModel.SakRekey = sakRekeyPatchModel
				gatewayMacsecPatchModel.SecurityPolicy = core.StringPtr("must_secure")
				gatewayMacsecPatchModel.WindowSize = core.Int64Ptr(int64(64))
				gatewayMacsecPatchModelAsPatch, asPatchErr := gatewayMacsecPatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateGatewayMacsecOptions model
				updateGatewayMacsecOptionsModel := new(directlinkv1.UpdateGatewayMacsecOptions)
				updateGatewayMacsecOptionsModel.ID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				updateGatewayMacsecOptionsModel.GatewayMacsecPatch = gatewayMacsecPatchModelAsPatch
				updateGatewayMacsecOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := directLinkService.UpdateGatewayMacsec(updateGatewayMacsecOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				directLinkService.EnableRetries(0, 0)
				result, response, operationErr = directLinkService.UpdateGatewayMacsec(updateGatewayMacsecOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateGatewayMacsec(updateGatewayMacsecOptions *UpdateGatewayMacsecOptions)`, func() {
		version := "testString"
		updateGatewayMacsecPath := "/gateways/0a06fb9b-820f-4c44-8a31-77f1f0806d28/macsec"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateGatewayMacsecPath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"active": true, "cipher_suite": "gcm_aes_xpn_256", "confidentiality_offset": 0, "created_at": "2020-11-02T20:40:29.622Z", "key_server_priority": 255, "sak_rekey": {"interval": 3600, "mode": "timer"}, "security_policy": "must_secure", "status": "secured", "status_reasons": [{"code": "macsec_cak_failed", "message": "The authentication_key failed configuration.", "more_info": "https://cloud.ibm.com/docs/dl/TODO_ADD_DOCS_LINK"}], "updated_at": "2020-11-02T20:40:29.622Z", "window_size": 512}`)
				}))
			})
			It(`Invoke UpdateGatewayMacsec successfully with retries`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())
				directLinkService.EnableRetries(0, 0)

				// Construct an instance of the SakRekeyPatchSakRekeyTimerModePatch model
				sakRekeyPatchModel := new(directlinkv1.SakRekeyPatchSakRekeyTimerModePatch)
				sakRekeyPatchModel.Interval = core.Int64Ptr(int64(3600))
				sakRekeyPatchModel.Mode = core.StringPtr("timer")

				// Construct an instance of the GatewayMacsecPatch model
				gatewayMacsecPatchModel := new(directlinkv1.GatewayMacsecPatch)
				gatewayMacsecPatchModel.Active = core.BoolPtr(true)
				gatewayMacsecPatchModel.SakRekey = sakRekeyPatchModel
				gatewayMacsecPatchModel.SecurityPolicy = core.StringPtr("must_secure")
				gatewayMacsecPatchModel.WindowSize = core.Int64Ptr(int64(64))
				gatewayMacsecPatchModelAsPatch, asPatchErr := gatewayMacsecPatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateGatewayMacsecOptions model
				updateGatewayMacsecOptionsModel := new(directlinkv1.UpdateGatewayMacsecOptions)
				updateGatewayMacsecOptionsModel.ID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				updateGatewayMacsecOptionsModel.GatewayMacsecPatch = gatewayMacsecPatchModelAsPatch
				updateGatewayMacsecOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := directLinkService.UpdateGatewayMacsecWithContext(ctx, updateGatewayMacsecOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				directLinkService.DisableRetries()
				result, response, operationErr := directLinkService.UpdateGatewayMacsec(updateGatewayMacsecOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = directLinkService.UpdateGatewayMacsecWithContext(ctx, updateGatewayMacsecOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateGatewayMacsecPath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"active": true, "cipher_suite": "gcm_aes_xpn_256", "confidentiality_offset": 0, "created_at": "2020-11-02T20:40:29.622Z", "key_server_priority": 255, "sak_rekey": {"interval": 3600, "mode": "timer"}, "security_policy": "must_secure", "status": "secured", "status_reasons": [{"code": "macsec_cak_failed", "message": "The authentication_key failed configuration.", "more_info": "https://cloud.ibm.com/docs/dl/TODO_ADD_DOCS_LINK"}], "updated_at": "2020-11-02T20:40:29.622Z", "window_size": 512}`)
				}))
			})
			It(`Invoke UpdateGatewayMacsec successfully`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := directLinkService.UpdateGatewayMacsec(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the SakRekeyPatchSakRekeyTimerModePatch model
				sakRekeyPatchModel := new(directlinkv1.SakRekeyPatchSakRekeyTimerModePatch)
				sakRekeyPatchModel.Interval = core.Int64Ptr(int64(3600))
				sakRekeyPatchModel.Mode = core.StringPtr("timer")

				// Construct an instance of the GatewayMacsecPatch model
				gatewayMacsecPatchModel := new(directlinkv1.GatewayMacsecPatch)
				gatewayMacsecPatchModel.Active = core.BoolPtr(true)
				gatewayMacsecPatchModel.SakRekey = sakRekeyPatchModel
				gatewayMacsecPatchModel.SecurityPolicy = core.StringPtr("must_secure")
				gatewayMacsecPatchModel.WindowSize = core.Int64Ptr(int64(64))
				gatewayMacsecPatchModelAsPatch, asPatchErr := gatewayMacsecPatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateGatewayMacsecOptions model
				updateGatewayMacsecOptionsModel := new(directlinkv1.UpdateGatewayMacsecOptions)
				updateGatewayMacsecOptionsModel.ID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				updateGatewayMacsecOptionsModel.GatewayMacsecPatch = gatewayMacsecPatchModelAsPatch
				updateGatewayMacsecOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = directLinkService.UpdateGatewayMacsec(updateGatewayMacsecOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateGatewayMacsec with error: Operation validation and request error`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Construct an instance of the SakRekeyPatchSakRekeyTimerModePatch model
				sakRekeyPatchModel := new(directlinkv1.SakRekeyPatchSakRekeyTimerModePatch)
				sakRekeyPatchModel.Interval = core.Int64Ptr(int64(3600))
				sakRekeyPatchModel.Mode = core.StringPtr("timer")

				// Construct an instance of the GatewayMacsecPatch model
				gatewayMacsecPatchModel := new(directlinkv1.GatewayMacsecPatch)
				gatewayMacsecPatchModel.Active = core.BoolPtr(true)
				gatewayMacsecPatchModel.SakRekey = sakRekeyPatchModel
				gatewayMacsecPatchModel.SecurityPolicy = core.StringPtr("must_secure")
				gatewayMacsecPatchModel.WindowSize = core.Int64Ptr(int64(64))
				gatewayMacsecPatchModelAsPatch, asPatchErr := gatewayMacsecPatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateGatewayMacsecOptions model
				updateGatewayMacsecOptionsModel := new(directlinkv1.UpdateGatewayMacsecOptions)
				updateGatewayMacsecOptionsModel.ID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				updateGatewayMacsecOptionsModel.GatewayMacsecPatch = gatewayMacsecPatchModelAsPatch
				updateGatewayMacsecOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := directLinkService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := directLinkService.UpdateGatewayMacsec(updateGatewayMacsecOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateGatewayMacsecOptions model with no property values
				updateGatewayMacsecOptionsModelNew := new(directlinkv1.UpdateGatewayMacsecOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = directLinkService.UpdateGatewayMacsec(updateGatewayMacsecOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke UpdateGatewayMacsec successfully`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Construct an instance of the SakRekeyPatchSakRekeyTimerModePatch model
				sakRekeyPatchModel := new(directlinkv1.SakRekeyPatchSakRekeyTimerModePatch)
				sakRekeyPatchModel.Interval = core.Int64Ptr(int64(3600))
				sakRekeyPatchModel.Mode = core.StringPtr("timer")

				// Construct an instance of the GatewayMacsecPatch model
				gatewayMacsecPatchModel := new(directlinkv1.GatewayMacsecPatch)
				gatewayMacsecPatchModel.Active = core.BoolPtr(true)
				gatewayMacsecPatchModel.SakRekey = sakRekeyPatchModel
				gatewayMacsecPatchModel.SecurityPolicy = core.StringPtr("must_secure")
				gatewayMacsecPatchModel.WindowSize = core.Int64Ptr(int64(64))
				gatewayMacsecPatchModelAsPatch, asPatchErr := gatewayMacsecPatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateGatewayMacsecOptions model
				updateGatewayMacsecOptionsModel := new(directlinkv1.UpdateGatewayMacsecOptions)
				updateGatewayMacsecOptionsModel.ID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				updateGatewayMacsecOptionsModel.GatewayMacsecPatch = gatewayMacsecPatchModelAsPatch
				updateGatewayMacsecOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := directLinkService.UpdateGatewayMacsec(updateGatewayMacsecOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`SetGatewayMacsec(setGatewayMacsecOptions *SetGatewayMacsecOptions) - Operation response error`, func() {
		version := "testString"
		setGatewayMacsecPath := "/gateways/0a06fb9b-820f-4c44-8a31-77f1f0806d28/macsec"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(setGatewayMacsecPath))
					Expect(req.Method).To(Equal("PUT"))
					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "W/\"96d225c4-56bd-43d9-98fc-d7148e5c5028\"")))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke SetGatewayMacsec with error: Operation response processing error`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Construct an instance of the HpcsKeyIdentity model
				hpcsKeyIdentityModel := new(directlinkv1.HpcsKeyIdentity)
				hpcsKeyIdentityModel.Crn = core.StringPtr("crn:v1:bluemix:public:hs-crypto:us-south:a/4111d05f36894e3cb9b46a43556d9000:abc111b8-37aa-4034-9def-f2607c87aaaa:key:bbb222bc-430a-4de9-9aad-84e5bb022222")

				// Construct an instance of the GatewayMacsecCakPrototype model
				gatewayMacsecCakPrototypeModel := new(directlinkv1.GatewayMacsecCakPrototype)
				gatewayMacsecCakPrototypeModel.Key = hpcsKeyIdentityModel
				gatewayMacsecCakPrototypeModel.Name = core.StringPtr("1000")
				gatewayMacsecCakPrototypeModel.Session = core.StringPtr("primary")

				// Construct an instance of the SakRekeyPrototypeSakRekeyTimerModePrototype model
				sakRekeyPrototypeModel := new(directlinkv1.SakRekeyPrototypeSakRekeyTimerModePrototype)
				sakRekeyPrototypeModel.Interval = core.Int64Ptr(int64(3600))
				sakRekeyPrototypeModel.Mode = core.StringPtr("timer")

				// Construct an instance of the SetGatewayMacsecOptions model
				setGatewayMacsecOptionsModel := new(directlinkv1.SetGatewayMacsecOptions)
				setGatewayMacsecOptionsModel.ID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				setGatewayMacsecOptionsModel.Active = core.BoolPtr(true)
				setGatewayMacsecOptionsModel.Caks = []directlinkv1.GatewayMacsecCakPrototype{*gatewayMacsecCakPrototypeModel}
				setGatewayMacsecOptionsModel.SakRekey = sakRekeyPrototypeModel
				setGatewayMacsecOptionsModel.SecurityPolicy = core.StringPtr("must_secure")
				setGatewayMacsecOptionsModel.WindowSize = core.Int64Ptr(int64(64))
				setGatewayMacsecOptionsModel.IfMatch = core.StringPtr("W/\"96d225c4-56bd-43d9-98fc-d7148e5c5028\"")
				setGatewayMacsecOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := directLinkService.SetGatewayMacsec(setGatewayMacsecOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				directLinkService.EnableRetries(0, 0)
				result, response, operationErr = directLinkService.SetGatewayMacsec(setGatewayMacsecOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`SetGatewayMacsec(setGatewayMacsecOptions *SetGatewayMacsecOptions)`, func() {
		version := "testString"
		setGatewayMacsecPath := "/gateways/0a06fb9b-820f-4c44-8a31-77f1f0806d28/macsec"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(setGatewayMacsecPath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "W/\"96d225c4-56bd-43d9-98fc-d7148e5c5028\"")))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"active": true, "cipher_suite": "gcm_aes_xpn_256", "confidentiality_offset": 0, "created_at": "2020-11-02T20:40:29.622Z", "key_server_priority": 255, "sak_rekey": {"interval": 3600, "mode": "timer"}, "security_policy": "must_secure", "status": "secured", "status_reasons": [{"code": "macsec_cak_failed", "message": "The authentication_key failed configuration.", "more_info": "https://cloud.ibm.com/docs/dl/TODO_ADD_DOCS_LINK"}], "updated_at": "2020-11-02T20:40:29.622Z", "window_size": 512}`)
				}))
			})
			It(`Invoke SetGatewayMacsec successfully with retries`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())
				directLinkService.EnableRetries(0, 0)

				// Construct an instance of the HpcsKeyIdentity model
				hpcsKeyIdentityModel := new(directlinkv1.HpcsKeyIdentity)
				hpcsKeyIdentityModel.Crn = core.StringPtr("crn:v1:bluemix:public:hs-crypto:us-south:a/4111d05f36894e3cb9b46a43556d9000:abc111b8-37aa-4034-9def-f2607c87aaaa:key:bbb222bc-430a-4de9-9aad-84e5bb022222")

				// Construct an instance of the GatewayMacsecCakPrototype model
				gatewayMacsecCakPrototypeModel := new(directlinkv1.GatewayMacsecCakPrototype)
				gatewayMacsecCakPrototypeModel.Key = hpcsKeyIdentityModel
				gatewayMacsecCakPrototypeModel.Name = core.StringPtr("1000")
				gatewayMacsecCakPrototypeModel.Session = core.StringPtr("primary")

				// Construct an instance of the SakRekeyPrototypeSakRekeyTimerModePrototype model
				sakRekeyPrototypeModel := new(directlinkv1.SakRekeyPrototypeSakRekeyTimerModePrototype)
				sakRekeyPrototypeModel.Interval = core.Int64Ptr(int64(3600))
				sakRekeyPrototypeModel.Mode = core.StringPtr("timer")

				// Construct an instance of the SetGatewayMacsecOptions model
				setGatewayMacsecOptionsModel := new(directlinkv1.SetGatewayMacsecOptions)
				setGatewayMacsecOptionsModel.ID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				setGatewayMacsecOptionsModel.Active = core.BoolPtr(true)
				setGatewayMacsecOptionsModel.Caks = []directlinkv1.GatewayMacsecCakPrototype{*gatewayMacsecCakPrototypeModel}
				setGatewayMacsecOptionsModel.SakRekey = sakRekeyPrototypeModel
				setGatewayMacsecOptionsModel.SecurityPolicy = core.StringPtr("must_secure")
				setGatewayMacsecOptionsModel.WindowSize = core.Int64Ptr(int64(64))
				setGatewayMacsecOptionsModel.IfMatch = core.StringPtr("W/\"96d225c4-56bd-43d9-98fc-d7148e5c5028\"")
				setGatewayMacsecOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := directLinkService.SetGatewayMacsecWithContext(ctx, setGatewayMacsecOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				directLinkService.DisableRetries()
				result, response, operationErr := directLinkService.SetGatewayMacsec(setGatewayMacsecOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = directLinkService.SetGatewayMacsecWithContext(ctx, setGatewayMacsecOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(setGatewayMacsecPath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "W/\"96d225c4-56bd-43d9-98fc-d7148e5c5028\"")))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"active": true, "cipher_suite": "gcm_aes_xpn_256", "confidentiality_offset": 0, "created_at": "2020-11-02T20:40:29.622Z", "key_server_priority": 255, "sak_rekey": {"interval": 3600, "mode": "timer"}, "security_policy": "must_secure", "status": "secured", "status_reasons": [{"code": "macsec_cak_failed", "message": "The authentication_key failed configuration.", "more_info": "https://cloud.ibm.com/docs/dl/TODO_ADD_DOCS_LINK"}], "updated_at": "2020-11-02T20:40:29.622Z", "window_size": 512}`)
				}))
			})
			It(`Invoke SetGatewayMacsec successfully`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := directLinkService.SetGatewayMacsec(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the HpcsKeyIdentity model
				hpcsKeyIdentityModel := new(directlinkv1.HpcsKeyIdentity)
				hpcsKeyIdentityModel.Crn = core.StringPtr("crn:v1:bluemix:public:hs-crypto:us-south:a/4111d05f36894e3cb9b46a43556d9000:abc111b8-37aa-4034-9def-f2607c87aaaa:key:bbb222bc-430a-4de9-9aad-84e5bb022222")

				// Construct an instance of the GatewayMacsecCakPrototype model
				gatewayMacsecCakPrototypeModel := new(directlinkv1.GatewayMacsecCakPrototype)
				gatewayMacsecCakPrototypeModel.Key = hpcsKeyIdentityModel
				gatewayMacsecCakPrototypeModel.Name = core.StringPtr("1000")
				gatewayMacsecCakPrototypeModel.Session = core.StringPtr("primary")

				// Construct an instance of the SakRekeyPrototypeSakRekeyTimerModePrototype model
				sakRekeyPrototypeModel := new(directlinkv1.SakRekeyPrototypeSakRekeyTimerModePrototype)
				sakRekeyPrototypeModel.Interval = core.Int64Ptr(int64(3600))
				sakRekeyPrototypeModel.Mode = core.StringPtr("timer")

				// Construct an instance of the SetGatewayMacsecOptions model
				setGatewayMacsecOptionsModel := new(directlinkv1.SetGatewayMacsecOptions)
				setGatewayMacsecOptionsModel.ID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				setGatewayMacsecOptionsModel.Active = core.BoolPtr(true)
				setGatewayMacsecOptionsModel.Caks = []directlinkv1.GatewayMacsecCakPrototype{*gatewayMacsecCakPrototypeModel}
				setGatewayMacsecOptionsModel.SakRekey = sakRekeyPrototypeModel
				setGatewayMacsecOptionsModel.SecurityPolicy = core.StringPtr("must_secure")
				setGatewayMacsecOptionsModel.WindowSize = core.Int64Ptr(int64(64))
				setGatewayMacsecOptionsModel.IfMatch = core.StringPtr("W/\"96d225c4-56bd-43d9-98fc-d7148e5c5028\"")
				setGatewayMacsecOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = directLinkService.SetGatewayMacsec(setGatewayMacsecOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke SetGatewayMacsec with error: Operation validation and request error`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Construct an instance of the HpcsKeyIdentity model
				hpcsKeyIdentityModel := new(directlinkv1.HpcsKeyIdentity)
				hpcsKeyIdentityModel.Crn = core.StringPtr("crn:v1:bluemix:public:hs-crypto:us-south:a/4111d05f36894e3cb9b46a43556d9000:abc111b8-37aa-4034-9def-f2607c87aaaa:key:bbb222bc-430a-4de9-9aad-84e5bb022222")

				// Construct an instance of the GatewayMacsecCakPrototype model
				gatewayMacsecCakPrototypeModel := new(directlinkv1.GatewayMacsecCakPrototype)
				gatewayMacsecCakPrototypeModel.Key = hpcsKeyIdentityModel
				gatewayMacsecCakPrototypeModel.Name = core.StringPtr("1000")
				gatewayMacsecCakPrototypeModel.Session = core.StringPtr("primary")

				// Construct an instance of the SakRekeyPrototypeSakRekeyTimerModePrototype model
				sakRekeyPrototypeModel := new(directlinkv1.SakRekeyPrototypeSakRekeyTimerModePrototype)
				sakRekeyPrototypeModel.Interval = core.Int64Ptr(int64(3600))
				sakRekeyPrototypeModel.Mode = core.StringPtr("timer")

				// Construct an instance of the SetGatewayMacsecOptions model
				setGatewayMacsecOptionsModel := new(directlinkv1.SetGatewayMacsecOptions)
				setGatewayMacsecOptionsModel.ID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				setGatewayMacsecOptionsModel.Active = core.BoolPtr(true)
				setGatewayMacsecOptionsModel.Caks = []directlinkv1.GatewayMacsecCakPrototype{*gatewayMacsecCakPrototypeModel}
				setGatewayMacsecOptionsModel.SakRekey = sakRekeyPrototypeModel
				setGatewayMacsecOptionsModel.SecurityPolicy = core.StringPtr("must_secure")
				setGatewayMacsecOptionsModel.WindowSize = core.Int64Ptr(int64(64))
				setGatewayMacsecOptionsModel.IfMatch = core.StringPtr("W/\"96d225c4-56bd-43d9-98fc-d7148e5c5028\"")
				setGatewayMacsecOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := directLinkService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := directLinkService.SetGatewayMacsec(setGatewayMacsecOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the SetGatewayMacsecOptions model with no property values
				setGatewayMacsecOptionsModelNew := new(directlinkv1.SetGatewayMacsecOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = directLinkService.SetGatewayMacsec(setGatewayMacsecOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke SetGatewayMacsec successfully`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Construct an instance of the HpcsKeyIdentity model
				hpcsKeyIdentityModel := new(directlinkv1.HpcsKeyIdentity)
				hpcsKeyIdentityModel.Crn = core.StringPtr("crn:v1:bluemix:public:hs-crypto:us-south:a/4111d05f36894e3cb9b46a43556d9000:abc111b8-37aa-4034-9def-f2607c87aaaa:key:bbb222bc-430a-4de9-9aad-84e5bb022222")

				// Construct an instance of the GatewayMacsecCakPrototype model
				gatewayMacsecCakPrototypeModel := new(directlinkv1.GatewayMacsecCakPrototype)
				gatewayMacsecCakPrototypeModel.Key = hpcsKeyIdentityModel
				gatewayMacsecCakPrototypeModel.Name = core.StringPtr("1000")
				gatewayMacsecCakPrototypeModel.Session = core.StringPtr("primary")

				// Construct an instance of the SakRekeyPrototypeSakRekeyTimerModePrototype model
				sakRekeyPrototypeModel := new(directlinkv1.SakRekeyPrototypeSakRekeyTimerModePrototype)
				sakRekeyPrototypeModel.Interval = core.Int64Ptr(int64(3600))
				sakRekeyPrototypeModel.Mode = core.StringPtr("timer")

				// Construct an instance of the SetGatewayMacsecOptions model
				setGatewayMacsecOptionsModel := new(directlinkv1.SetGatewayMacsecOptions)
				setGatewayMacsecOptionsModel.ID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				setGatewayMacsecOptionsModel.Active = core.BoolPtr(true)
				setGatewayMacsecOptionsModel.Caks = []directlinkv1.GatewayMacsecCakPrototype{*gatewayMacsecCakPrototypeModel}
				setGatewayMacsecOptionsModel.SakRekey = sakRekeyPrototypeModel
				setGatewayMacsecOptionsModel.SecurityPolicy = core.StringPtr("must_secure")
				setGatewayMacsecOptionsModel.WindowSize = core.Int64Ptr(int64(64))
				setGatewayMacsecOptionsModel.IfMatch = core.StringPtr("W/\"96d225c4-56bd-43d9-98fc-d7148e5c5028\"")
				setGatewayMacsecOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := directLinkService.SetGatewayMacsec(setGatewayMacsecOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListGatewayMacsecCaks(listGatewayMacsecCaksOptions *ListGatewayMacsecCaksOptions) - Operation response error`, func() {
		version := "testString"
		listGatewayMacsecCaksPath := "/gateways/0a06fb9b-820f-4c44-8a31-77f1f0806d28/macsec/caks"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listGatewayMacsecCaksPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListGatewayMacsecCaks with error: Operation response processing error`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Construct an instance of the ListGatewayMacsecCaksOptions model
				listGatewayMacsecCaksOptionsModel := new(directlinkv1.ListGatewayMacsecCaksOptions)
				listGatewayMacsecCaksOptionsModel.ID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				listGatewayMacsecCaksOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := directLinkService.ListGatewayMacsecCaks(listGatewayMacsecCaksOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				directLinkService.EnableRetries(0, 0)
				result, response, operationErr = directLinkService.ListGatewayMacsecCaks(listGatewayMacsecCaksOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListGatewayMacsecCaks(listGatewayMacsecCaksOptions *ListGatewayMacsecCaksOptions)`, func() {
		version := "testString"
		listGatewayMacsecCaksPath := "/gateways/0a06fb9b-820f-4c44-8a31-77f1f0806d28/macsec/caks"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listGatewayMacsecCaksPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"caks": [{"active_delta": {"key": {"crn": "crn:v1:bluemix:public:hs-crypto:us-south:a/4111d05f36894e3cb9b46a43556d9000:abc111b8-37aa-4034-9def-f2607c87aaaa:key:bbb222bc-430a-4de9-9aad-84e5bb022222"}, "name": "1000"}, "created_at": "2020-11-02T20:40:29.622Z", "id": "ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4", "key": {"crn": "crn:v1:bluemix:public:hs-crypto:us-south:a/4111d05f36894e3cb9b46a43556d9000:abc111b8-37aa-4034-9def-f2607c87aaaa:key:bbb222bc-430a-4de9-9aad-84e5bb022222"}, "name": "1000", "session": "primary", "status": "active", "updated_at": "2020-11-02T20:40:29.622Z"}]}`)
				}))
			})
			It(`Invoke ListGatewayMacsecCaks successfully with retries`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())
				directLinkService.EnableRetries(0, 0)

				// Construct an instance of the ListGatewayMacsecCaksOptions model
				listGatewayMacsecCaksOptionsModel := new(directlinkv1.ListGatewayMacsecCaksOptions)
				listGatewayMacsecCaksOptionsModel.ID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				listGatewayMacsecCaksOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := directLinkService.ListGatewayMacsecCaksWithContext(ctx, listGatewayMacsecCaksOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				directLinkService.DisableRetries()
				result, response, operationErr := directLinkService.ListGatewayMacsecCaks(listGatewayMacsecCaksOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = directLinkService.ListGatewayMacsecCaksWithContext(ctx, listGatewayMacsecCaksOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listGatewayMacsecCaksPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"caks": [{"active_delta": {"key": {"crn": "crn:v1:bluemix:public:hs-crypto:us-south:a/4111d05f36894e3cb9b46a43556d9000:abc111b8-37aa-4034-9def-f2607c87aaaa:key:bbb222bc-430a-4de9-9aad-84e5bb022222"}, "name": "1000"}, "created_at": "2020-11-02T20:40:29.622Z", "id": "ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4", "key": {"crn": "crn:v1:bluemix:public:hs-crypto:us-south:a/4111d05f36894e3cb9b46a43556d9000:abc111b8-37aa-4034-9def-f2607c87aaaa:key:bbb222bc-430a-4de9-9aad-84e5bb022222"}, "name": "1000", "session": "primary", "status": "active", "updated_at": "2020-11-02T20:40:29.622Z"}]}`)
				}))
			})
			It(`Invoke ListGatewayMacsecCaks successfully`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := directLinkService.ListGatewayMacsecCaks(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListGatewayMacsecCaksOptions model
				listGatewayMacsecCaksOptionsModel := new(directlinkv1.ListGatewayMacsecCaksOptions)
				listGatewayMacsecCaksOptionsModel.ID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				listGatewayMacsecCaksOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = directLinkService.ListGatewayMacsecCaks(listGatewayMacsecCaksOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListGatewayMacsecCaks with error: Operation validation and request error`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Construct an instance of the ListGatewayMacsecCaksOptions model
				listGatewayMacsecCaksOptionsModel := new(directlinkv1.ListGatewayMacsecCaksOptions)
				listGatewayMacsecCaksOptionsModel.ID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				listGatewayMacsecCaksOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := directLinkService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := directLinkService.ListGatewayMacsecCaks(listGatewayMacsecCaksOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListGatewayMacsecCaksOptions model with no property values
				listGatewayMacsecCaksOptionsModelNew := new(directlinkv1.ListGatewayMacsecCaksOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = directLinkService.ListGatewayMacsecCaks(listGatewayMacsecCaksOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ListGatewayMacsecCaks successfully`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Construct an instance of the ListGatewayMacsecCaksOptions model
				listGatewayMacsecCaksOptionsModel := new(directlinkv1.ListGatewayMacsecCaksOptions)
				listGatewayMacsecCaksOptionsModel.ID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				listGatewayMacsecCaksOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := directLinkService.ListGatewayMacsecCaks(listGatewayMacsecCaksOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateGatewayMacsecCak(createGatewayMacsecCakOptions *CreateGatewayMacsecCakOptions) - Operation response error`, func() {
		version := "testString"
		createGatewayMacsecCakPath := "/gateways/0a06fb9b-820f-4c44-8a31-77f1f0806d28/macsec/caks"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createGatewayMacsecCakPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateGatewayMacsecCak with error: Operation response processing error`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Construct an instance of the HpcsKeyIdentity model
				hpcsKeyIdentityModel := new(directlinkv1.HpcsKeyIdentity)
				hpcsKeyIdentityModel.Crn = core.StringPtr("crn:v1:bluemix:public:hs-crypto:us-south:a/4111d05f36894e3cb9b46a43556d9000:abc111b8-37aa-4034-9def-f2607c87aaaa:key:bbb222bc-430a-4de9-9aad-84e5bb022222")

				// Construct an instance of the CreateGatewayMacsecCakOptions model
				createGatewayMacsecCakOptionsModel := new(directlinkv1.CreateGatewayMacsecCakOptions)
				createGatewayMacsecCakOptionsModel.ID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				createGatewayMacsecCakOptionsModel.Key = hpcsKeyIdentityModel
				createGatewayMacsecCakOptionsModel.Name = core.StringPtr("1000")
				createGatewayMacsecCakOptionsModel.Session = core.StringPtr("primary")
				createGatewayMacsecCakOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := directLinkService.CreateGatewayMacsecCak(createGatewayMacsecCakOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				directLinkService.EnableRetries(0, 0)
				result, response, operationErr = directLinkService.CreateGatewayMacsecCak(createGatewayMacsecCakOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateGatewayMacsecCak(createGatewayMacsecCakOptions *CreateGatewayMacsecCakOptions)`, func() {
		version := "testString"
		createGatewayMacsecCakPath := "/gateways/0a06fb9b-820f-4c44-8a31-77f1f0806d28/macsec/caks"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createGatewayMacsecCakPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"active_delta": {"key": {"crn": "crn:v1:bluemix:public:hs-crypto:us-south:a/4111d05f36894e3cb9b46a43556d9000:abc111b8-37aa-4034-9def-f2607c87aaaa:key:bbb222bc-430a-4de9-9aad-84e5bb022222"}, "name": "1000"}, "created_at": "2020-11-02T20:40:29.622Z", "id": "ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4", "key": {"crn": "crn:v1:bluemix:public:hs-crypto:us-south:a/4111d05f36894e3cb9b46a43556d9000:abc111b8-37aa-4034-9def-f2607c87aaaa:key:bbb222bc-430a-4de9-9aad-84e5bb022222"}, "name": "1000", "session": "primary", "status": "active", "updated_at": "2020-11-02T20:40:29.622Z"}`)
				}))
			})
			It(`Invoke CreateGatewayMacsecCak successfully with retries`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())
				directLinkService.EnableRetries(0, 0)

				// Construct an instance of the HpcsKeyIdentity model
				hpcsKeyIdentityModel := new(directlinkv1.HpcsKeyIdentity)
				hpcsKeyIdentityModel.Crn = core.StringPtr("crn:v1:bluemix:public:hs-crypto:us-south:a/4111d05f36894e3cb9b46a43556d9000:abc111b8-37aa-4034-9def-f2607c87aaaa:key:bbb222bc-430a-4de9-9aad-84e5bb022222")

				// Construct an instance of the CreateGatewayMacsecCakOptions model
				createGatewayMacsecCakOptionsModel := new(directlinkv1.CreateGatewayMacsecCakOptions)
				createGatewayMacsecCakOptionsModel.ID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				createGatewayMacsecCakOptionsModel.Key = hpcsKeyIdentityModel
				createGatewayMacsecCakOptionsModel.Name = core.StringPtr("1000")
				createGatewayMacsecCakOptionsModel.Session = core.StringPtr("primary")
				createGatewayMacsecCakOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := directLinkService.CreateGatewayMacsecCakWithContext(ctx, createGatewayMacsecCakOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				directLinkService.DisableRetries()
				result, response, operationErr := directLinkService.CreateGatewayMacsecCak(createGatewayMacsecCakOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = directLinkService.CreateGatewayMacsecCakWithContext(ctx, createGatewayMacsecCakOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createGatewayMacsecCakPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"active_delta": {"key": {"crn": "crn:v1:bluemix:public:hs-crypto:us-south:a/4111d05f36894e3cb9b46a43556d9000:abc111b8-37aa-4034-9def-f2607c87aaaa:key:bbb222bc-430a-4de9-9aad-84e5bb022222"}, "name": "1000"}, "created_at": "2020-11-02T20:40:29.622Z", "id": "ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4", "key": {"crn": "crn:v1:bluemix:public:hs-crypto:us-south:a/4111d05f36894e3cb9b46a43556d9000:abc111b8-37aa-4034-9def-f2607c87aaaa:key:bbb222bc-430a-4de9-9aad-84e5bb022222"}, "name": "1000", "session": "primary", "status": "active", "updated_at": "2020-11-02T20:40:29.622Z"}`)
				}))
			})
			It(`Invoke CreateGatewayMacsecCak successfully`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := directLinkService.CreateGatewayMacsecCak(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the HpcsKeyIdentity model
				hpcsKeyIdentityModel := new(directlinkv1.HpcsKeyIdentity)
				hpcsKeyIdentityModel.Crn = core.StringPtr("crn:v1:bluemix:public:hs-crypto:us-south:a/4111d05f36894e3cb9b46a43556d9000:abc111b8-37aa-4034-9def-f2607c87aaaa:key:bbb222bc-430a-4de9-9aad-84e5bb022222")

				// Construct an instance of the CreateGatewayMacsecCakOptions model
				createGatewayMacsecCakOptionsModel := new(directlinkv1.CreateGatewayMacsecCakOptions)
				createGatewayMacsecCakOptionsModel.ID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				createGatewayMacsecCakOptionsModel.Key = hpcsKeyIdentityModel
				createGatewayMacsecCakOptionsModel.Name = core.StringPtr("1000")
				createGatewayMacsecCakOptionsModel.Session = core.StringPtr("primary")
				createGatewayMacsecCakOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = directLinkService.CreateGatewayMacsecCak(createGatewayMacsecCakOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateGatewayMacsecCak with error: Operation validation and request error`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Construct an instance of the HpcsKeyIdentity model
				hpcsKeyIdentityModel := new(directlinkv1.HpcsKeyIdentity)
				hpcsKeyIdentityModel.Crn = core.StringPtr("crn:v1:bluemix:public:hs-crypto:us-south:a/4111d05f36894e3cb9b46a43556d9000:abc111b8-37aa-4034-9def-f2607c87aaaa:key:bbb222bc-430a-4de9-9aad-84e5bb022222")

				// Construct an instance of the CreateGatewayMacsecCakOptions model
				createGatewayMacsecCakOptionsModel := new(directlinkv1.CreateGatewayMacsecCakOptions)
				createGatewayMacsecCakOptionsModel.ID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				createGatewayMacsecCakOptionsModel.Key = hpcsKeyIdentityModel
				createGatewayMacsecCakOptionsModel.Name = core.StringPtr("1000")
				createGatewayMacsecCakOptionsModel.Session = core.StringPtr("primary")
				createGatewayMacsecCakOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := directLinkService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := directLinkService.CreateGatewayMacsecCak(createGatewayMacsecCakOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateGatewayMacsecCakOptions model with no property values
				createGatewayMacsecCakOptionsModelNew := new(directlinkv1.CreateGatewayMacsecCakOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = directLinkService.CreateGatewayMacsecCak(createGatewayMacsecCakOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(201)
				}))
			})
			It(`Invoke CreateGatewayMacsecCak successfully`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Construct an instance of the HpcsKeyIdentity model
				hpcsKeyIdentityModel := new(directlinkv1.HpcsKeyIdentity)
				hpcsKeyIdentityModel.Crn = core.StringPtr("crn:v1:bluemix:public:hs-crypto:us-south:a/4111d05f36894e3cb9b46a43556d9000:abc111b8-37aa-4034-9def-f2607c87aaaa:key:bbb222bc-430a-4de9-9aad-84e5bb022222")

				// Construct an instance of the CreateGatewayMacsecCakOptions model
				createGatewayMacsecCakOptionsModel := new(directlinkv1.CreateGatewayMacsecCakOptions)
				createGatewayMacsecCakOptionsModel.ID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				createGatewayMacsecCakOptionsModel.Key = hpcsKeyIdentityModel
				createGatewayMacsecCakOptionsModel.Name = core.StringPtr("1000")
				createGatewayMacsecCakOptionsModel.Session = core.StringPtr("primary")
				createGatewayMacsecCakOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := directLinkService.CreateGatewayMacsecCak(createGatewayMacsecCakOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteGatewayMacsecCak(deleteGatewayMacsecCakOptions *DeleteGatewayMacsecCakOptions)`, func() {
		version := "testString"
		deleteGatewayMacsecCakPath := "/gateways/0a06fb9b-820f-4c44-8a31-77f1f0806d28/macsec/caks/ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteGatewayMacsecCakPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteGatewayMacsecCak successfully`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := directLinkService.DeleteGatewayMacsecCak(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteGatewayMacsecCakOptions model
				deleteGatewayMacsecCakOptionsModel := new(directlinkv1.DeleteGatewayMacsecCakOptions)
				deleteGatewayMacsecCakOptionsModel.ID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				deleteGatewayMacsecCakOptionsModel.CakID = core.StringPtr("ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4")
				deleteGatewayMacsecCakOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = directLinkService.DeleteGatewayMacsecCak(deleteGatewayMacsecCakOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteGatewayMacsecCak with error: Operation validation and request error`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Construct an instance of the DeleteGatewayMacsecCakOptions model
				deleteGatewayMacsecCakOptionsModel := new(directlinkv1.DeleteGatewayMacsecCakOptions)
				deleteGatewayMacsecCakOptionsModel.ID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				deleteGatewayMacsecCakOptionsModel.CakID = core.StringPtr("ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4")
				deleteGatewayMacsecCakOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := directLinkService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := directLinkService.DeleteGatewayMacsecCak(deleteGatewayMacsecCakOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteGatewayMacsecCakOptions model with no property values
				deleteGatewayMacsecCakOptionsModelNew := new(directlinkv1.DeleteGatewayMacsecCakOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = directLinkService.DeleteGatewayMacsecCak(deleteGatewayMacsecCakOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetGatewayMacsecCak(getGatewayMacsecCakOptions *GetGatewayMacsecCakOptions) - Operation response error`, func() {
		version := "testString"
		getGatewayMacsecCakPath := "/gateways/0a06fb9b-820f-4c44-8a31-77f1f0806d28/macsec/caks/ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getGatewayMacsecCakPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetGatewayMacsecCak with error: Operation response processing error`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Construct an instance of the GetGatewayMacsecCakOptions model
				getGatewayMacsecCakOptionsModel := new(directlinkv1.GetGatewayMacsecCakOptions)
				getGatewayMacsecCakOptionsModel.ID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				getGatewayMacsecCakOptionsModel.CakID = core.StringPtr("ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4")
				getGatewayMacsecCakOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := directLinkService.GetGatewayMacsecCak(getGatewayMacsecCakOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				directLinkService.EnableRetries(0, 0)
				result, response, operationErr = directLinkService.GetGatewayMacsecCak(getGatewayMacsecCakOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetGatewayMacsecCak(getGatewayMacsecCakOptions *GetGatewayMacsecCakOptions)`, func() {
		version := "testString"
		getGatewayMacsecCakPath := "/gateways/0a06fb9b-820f-4c44-8a31-77f1f0806d28/macsec/caks/ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getGatewayMacsecCakPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"active_delta": {"key": {"crn": "crn:v1:bluemix:public:hs-crypto:us-south:a/4111d05f36894e3cb9b46a43556d9000:abc111b8-37aa-4034-9def-f2607c87aaaa:key:bbb222bc-430a-4de9-9aad-84e5bb022222"}, "name": "1000"}, "created_at": "2020-11-02T20:40:29.622Z", "id": "ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4", "key": {"crn": "crn:v1:bluemix:public:hs-crypto:us-south:a/4111d05f36894e3cb9b46a43556d9000:abc111b8-37aa-4034-9def-f2607c87aaaa:key:bbb222bc-430a-4de9-9aad-84e5bb022222"}, "name": "1000", "session": "primary", "status": "active", "updated_at": "2020-11-02T20:40:29.622Z"}`)
				}))
			})
			It(`Invoke GetGatewayMacsecCak successfully with retries`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())
				directLinkService.EnableRetries(0, 0)

				// Construct an instance of the GetGatewayMacsecCakOptions model
				getGatewayMacsecCakOptionsModel := new(directlinkv1.GetGatewayMacsecCakOptions)
				getGatewayMacsecCakOptionsModel.ID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				getGatewayMacsecCakOptionsModel.CakID = core.StringPtr("ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4")
				getGatewayMacsecCakOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := directLinkService.GetGatewayMacsecCakWithContext(ctx, getGatewayMacsecCakOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				directLinkService.DisableRetries()
				result, response, operationErr := directLinkService.GetGatewayMacsecCak(getGatewayMacsecCakOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = directLinkService.GetGatewayMacsecCakWithContext(ctx, getGatewayMacsecCakOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getGatewayMacsecCakPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"active_delta": {"key": {"crn": "crn:v1:bluemix:public:hs-crypto:us-south:a/4111d05f36894e3cb9b46a43556d9000:abc111b8-37aa-4034-9def-f2607c87aaaa:key:bbb222bc-430a-4de9-9aad-84e5bb022222"}, "name": "1000"}, "created_at": "2020-11-02T20:40:29.622Z", "id": "ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4", "key": {"crn": "crn:v1:bluemix:public:hs-crypto:us-south:a/4111d05f36894e3cb9b46a43556d9000:abc111b8-37aa-4034-9def-f2607c87aaaa:key:bbb222bc-430a-4de9-9aad-84e5bb022222"}, "name": "1000", "session": "primary", "status": "active", "updated_at": "2020-11-02T20:40:29.622Z"}`)
				}))
			})
			It(`Invoke GetGatewayMacsecCak successfully`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := directLinkService.GetGatewayMacsecCak(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetGatewayMacsecCakOptions model
				getGatewayMacsecCakOptionsModel := new(directlinkv1.GetGatewayMacsecCakOptions)
				getGatewayMacsecCakOptionsModel.ID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				getGatewayMacsecCakOptionsModel.CakID = core.StringPtr("ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4")
				getGatewayMacsecCakOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = directLinkService.GetGatewayMacsecCak(getGatewayMacsecCakOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetGatewayMacsecCak with error: Operation validation and request error`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Construct an instance of the GetGatewayMacsecCakOptions model
				getGatewayMacsecCakOptionsModel := new(directlinkv1.GetGatewayMacsecCakOptions)
				getGatewayMacsecCakOptionsModel.ID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				getGatewayMacsecCakOptionsModel.CakID = core.StringPtr("ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4")
				getGatewayMacsecCakOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := directLinkService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := directLinkService.GetGatewayMacsecCak(getGatewayMacsecCakOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetGatewayMacsecCakOptions model with no property values
				getGatewayMacsecCakOptionsModelNew := new(directlinkv1.GetGatewayMacsecCakOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = directLinkService.GetGatewayMacsecCak(getGatewayMacsecCakOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetGatewayMacsecCak successfully`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Construct an instance of the GetGatewayMacsecCakOptions model
				getGatewayMacsecCakOptionsModel := new(directlinkv1.GetGatewayMacsecCakOptions)
				getGatewayMacsecCakOptionsModel.ID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				getGatewayMacsecCakOptionsModel.CakID = core.StringPtr("ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4")
				getGatewayMacsecCakOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := directLinkService.GetGatewayMacsecCak(getGatewayMacsecCakOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateGatewayMacsecCak(updateGatewayMacsecCakOptions *UpdateGatewayMacsecCakOptions) - Operation response error`, func() {
		version := "testString"
		updateGatewayMacsecCakPath := "/gateways/0a06fb9b-820f-4c44-8a31-77f1f0806d28/macsec/caks/ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateGatewayMacsecCakPath))
					Expect(req.Method).To(Equal("PATCH"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateGatewayMacsecCak with error: Operation response processing error`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Construct an instance of the HpcsKeyIdentity model
				hpcsKeyIdentityModel := new(directlinkv1.HpcsKeyIdentity)
				hpcsKeyIdentityModel.Crn = core.StringPtr("crn:v1:bluemix:public:hs-crypto:us-south:a/4111d05f36894e3cb9b46a43556d9000:abc111b8-37aa-4034-9def-f2607c87aaaa:key:bbb222bc-430a-4de9-9aad-84e5bb022222")

				// Construct an instance of the GatewayMacsecCakPatch model
				gatewayMacsecCakPatchModel := new(directlinkv1.GatewayMacsecCakPatch)
				gatewayMacsecCakPatchModel.Key = hpcsKeyIdentityModel
				gatewayMacsecCakPatchModel.Name = core.StringPtr("1000")
				gatewayMacsecCakPatchModelAsPatch, asPatchErr := gatewayMacsecCakPatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateGatewayMacsecCakOptions model
				updateGatewayMacsecCakOptionsModel := new(directlinkv1.UpdateGatewayMacsecCakOptions)
				updateGatewayMacsecCakOptionsModel.ID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				updateGatewayMacsecCakOptionsModel.CakID = core.StringPtr("ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4")
				updateGatewayMacsecCakOptionsModel.GatewayMacsecCakPatch = gatewayMacsecCakPatchModelAsPatch
				updateGatewayMacsecCakOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := directLinkService.UpdateGatewayMacsecCak(updateGatewayMacsecCakOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				directLinkService.EnableRetries(0, 0)
				result, response, operationErr = directLinkService.UpdateGatewayMacsecCak(updateGatewayMacsecCakOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateGatewayMacsecCak(updateGatewayMacsecCakOptions *UpdateGatewayMacsecCakOptions)`, func() {
		version := "testString"
		updateGatewayMacsecCakPath := "/gateways/0a06fb9b-820f-4c44-8a31-77f1f0806d28/macsec/caks/ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateGatewayMacsecCakPath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"active_delta": {"key": {"crn": "crn:v1:bluemix:public:hs-crypto:us-south:a/4111d05f36894e3cb9b46a43556d9000:abc111b8-37aa-4034-9def-f2607c87aaaa:key:bbb222bc-430a-4de9-9aad-84e5bb022222"}, "name": "1000"}, "created_at": "2020-11-02T20:40:29.622Z", "id": "ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4", "key": {"crn": "crn:v1:bluemix:public:hs-crypto:us-south:a/4111d05f36894e3cb9b46a43556d9000:abc111b8-37aa-4034-9def-f2607c87aaaa:key:bbb222bc-430a-4de9-9aad-84e5bb022222"}, "name": "1000", "session": "primary", "status": "active", "updated_at": "2020-11-02T20:40:29.622Z"}`)
				}))
			})
			It(`Invoke UpdateGatewayMacsecCak successfully with retries`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())
				directLinkService.EnableRetries(0, 0)

				// Construct an instance of the HpcsKeyIdentity model
				hpcsKeyIdentityModel := new(directlinkv1.HpcsKeyIdentity)
				hpcsKeyIdentityModel.Crn = core.StringPtr("crn:v1:bluemix:public:hs-crypto:us-south:a/4111d05f36894e3cb9b46a43556d9000:abc111b8-37aa-4034-9def-f2607c87aaaa:key:bbb222bc-430a-4de9-9aad-84e5bb022222")

				// Construct an instance of the GatewayMacsecCakPatch model
				gatewayMacsecCakPatchModel := new(directlinkv1.GatewayMacsecCakPatch)
				gatewayMacsecCakPatchModel.Key = hpcsKeyIdentityModel
				gatewayMacsecCakPatchModel.Name = core.StringPtr("1000")
				gatewayMacsecCakPatchModelAsPatch, asPatchErr := gatewayMacsecCakPatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateGatewayMacsecCakOptions model
				updateGatewayMacsecCakOptionsModel := new(directlinkv1.UpdateGatewayMacsecCakOptions)
				updateGatewayMacsecCakOptionsModel.ID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				updateGatewayMacsecCakOptionsModel.CakID = core.StringPtr("ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4")
				updateGatewayMacsecCakOptionsModel.GatewayMacsecCakPatch = gatewayMacsecCakPatchModelAsPatch
				updateGatewayMacsecCakOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := directLinkService.UpdateGatewayMacsecCakWithContext(ctx, updateGatewayMacsecCakOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				directLinkService.DisableRetries()
				result, response, operationErr := directLinkService.UpdateGatewayMacsecCak(updateGatewayMacsecCakOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = directLinkService.UpdateGatewayMacsecCakWithContext(ctx, updateGatewayMacsecCakOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateGatewayMacsecCakPath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"active_delta": {"key": {"crn": "crn:v1:bluemix:public:hs-crypto:us-south:a/4111d05f36894e3cb9b46a43556d9000:abc111b8-37aa-4034-9def-f2607c87aaaa:key:bbb222bc-430a-4de9-9aad-84e5bb022222"}, "name": "1000"}, "created_at": "2020-11-02T20:40:29.622Z", "id": "ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4", "key": {"crn": "crn:v1:bluemix:public:hs-crypto:us-south:a/4111d05f36894e3cb9b46a43556d9000:abc111b8-37aa-4034-9def-f2607c87aaaa:key:bbb222bc-430a-4de9-9aad-84e5bb022222"}, "name": "1000", "session": "primary", "status": "active", "updated_at": "2020-11-02T20:40:29.622Z"}`)
				}))
			})
			It(`Invoke UpdateGatewayMacsecCak successfully`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := directLinkService.UpdateGatewayMacsecCak(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the HpcsKeyIdentity model
				hpcsKeyIdentityModel := new(directlinkv1.HpcsKeyIdentity)
				hpcsKeyIdentityModel.Crn = core.StringPtr("crn:v1:bluemix:public:hs-crypto:us-south:a/4111d05f36894e3cb9b46a43556d9000:abc111b8-37aa-4034-9def-f2607c87aaaa:key:bbb222bc-430a-4de9-9aad-84e5bb022222")

				// Construct an instance of the GatewayMacsecCakPatch model
				gatewayMacsecCakPatchModel := new(directlinkv1.GatewayMacsecCakPatch)
				gatewayMacsecCakPatchModel.Key = hpcsKeyIdentityModel
				gatewayMacsecCakPatchModel.Name = core.StringPtr("1000")
				gatewayMacsecCakPatchModelAsPatch, asPatchErr := gatewayMacsecCakPatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateGatewayMacsecCakOptions model
				updateGatewayMacsecCakOptionsModel := new(directlinkv1.UpdateGatewayMacsecCakOptions)
				updateGatewayMacsecCakOptionsModel.ID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				updateGatewayMacsecCakOptionsModel.CakID = core.StringPtr("ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4")
				updateGatewayMacsecCakOptionsModel.GatewayMacsecCakPatch = gatewayMacsecCakPatchModelAsPatch
				updateGatewayMacsecCakOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = directLinkService.UpdateGatewayMacsecCak(updateGatewayMacsecCakOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateGatewayMacsecCak with error: Operation validation and request error`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Construct an instance of the HpcsKeyIdentity model
				hpcsKeyIdentityModel := new(directlinkv1.HpcsKeyIdentity)
				hpcsKeyIdentityModel.Crn = core.StringPtr("crn:v1:bluemix:public:hs-crypto:us-south:a/4111d05f36894e3cb9b46a43556d9000:abc111b8-37aa-4034-9def-f2607c87aaaa:key:bbb222bc-430a-4de9-9aad-84e5bb022222")

				// Construct an instance of the GatewayMacsecCakPatch model
				gatewayMacsecCakPatchModel := new(directlinkv1.GatewayMacsecCakPatch)
				gatewayMacsecCakPatchModel.Key = hpcsKeyIdentityModel
				gatewayMacsecCakPatchModel.Name = core.StringPtr("1000")
				gatewayMacsecCakPatchModelAsPatch, asPatchErr := gatewayMacsecCakPatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateGatewayMacsecCakOptions model
				updateGatewayMacsecCakOptionsModel := new(directlinkv1.UpdateGatewayMacsecCakOptions)
				updateGatewayMacsecCakOptionsModel.ID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				updateGatewayMacsecCakOptionsModel.CakID = core.StringPtr("ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4")
				updateGatewayMacsecCakOptionsModel.GatewayMacsecCakPatch = gatewayMacsecCakPatchModelAsPatch
				updateGatewayMacsecCakOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := directLinkService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := directLinkService.UpdateGatewayMacsecCak(updateGatewayMacsecCakOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateGatewayMacsecCakOptions model with no property values
				updateGatewayMacsecCakOptionsModelNew := new(directlinkv1.UpdateGatewayMacsecCakOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = directLinkService.UpdateGatewayMacsecCak(updateGatewayMacsecCakOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke UpdateGatewayMacsecCak successfully`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Construct an instance of the HpcsKeyIdentity model
				hpcsKeyIdentityModel := new(directlinkv1.HpcsKeyIdentity)
				hpcsKeyIdentityModel.Crn = core.StringPtr("crn:v1:bluemix:public:hs-crypto:us-south:a/4111d05f36894e3cb9b46a43556d9000:abc111b8-37aa-4034-9def-f2607c87aaaa:key:bbb222bc-430a-4de9-9aad-84e5bb022222")

				// Construct an instance of the GatewayMacsecCakPatch model
				gatewayMacsecCakPatchModel := new(directlinkv1.GatewayMacsecCakPatch)
				gatewayMacsecCakPatchModel.Key = hpcsKeyIdentityModel
				gatewayMacsecCakPatchModel.Name = core.StringPtr("1000")
				gatewayMacsecCakPatchModelAsPatch, asPatchErr := gatewayMacsecCakPatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateGatewayMacsecCakOptions model
				updateGatewayMacsecCakOptionsModel := new(directlinkv1.UpdateGatewayMacsecCakOptions)
				updateGatewayMacsecCakOptionsModel.ID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				updateGatewayMacsecCakOptionsModel.CakID = core.StringPtr("ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4")
				updateGatewayMacsecCakOptionsModel.GatewayMacsecCakPatch = gatewayMacsecCakPatchModelAsPatch
				updateGatewayMacsecCakOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := directLinkService.UpdateGatewayMacsecCak(updateGatewayMacsecCakOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListGatewayRouteReports(listGatewayRouteReportsOptions *ListGatewayRouteReportsOptions) - Operation response error`, func() {
		version := "testString"
		listGatewayRouteReportsPath := "/gateways/0a06fb9b-820f-4c44-8a31-77f1f0806d28/route_reports"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listGatewayRouteReportsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListGatewayRouteReports with error: Operation response processing error`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Construct an instance of the ListGatewayRouteReportsOptions model
				listGatewayRouteReportsOptionsModel := new(directlinkv1.ListGatewayRouteReportsOptions)
				listGatewayRouteReportsOptionsModel.GatewayID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				listGatewayRouteReportsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := directLinkService.ListGatewayRouteReports(listGatewayRouteReportsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				directLinkService.EnableRetries(0, 0)
				result, response, operationErr = directLinkService.ListGatewayRouteReports(listGatewayRouteReportsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListGatewayRouteReports(listGatewayRouteReportsOptions *ListGatewayRouteReportsOptions)`, func() {
		version := "testString"
		listGatewayRouteReportsPath := "/gateways/0a06fb9b-820f-4c44-8a31-77f1f0806d28/route_reports"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listGatewayRouteReportsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"route_reports": [{"advertised_routes": [{"as_path": "64999 64999 64998 I", "prefix": "172.17.0.0/16"}], "created_at": "2019-01-01T12:00:00.000Z", "gateway_routes": [{"prefix": "172.17.0.0/16"}], "id": "1a15dcab-7e26-45e1-b7c5-bc690eaa9724", "on_prem_routes": [{"as_path": "64999 64999 64998 I", "next_hop": "172.17.0.0", "prefix": "172.17.0.0/16"}], "overlapping_routes": [{"routes": [{"prefix": "172.17.0.0/16", "type": "virtual_connection", "virtual_connection_id": "d2d985d8-1d8e-4e8b-96cd-cee2290ecaff"}]}], "status": "complete", "updated_at": "2019-01-01T12:00:00.000Z", "virtual_connection_routes": [{"routes": [{"active": true, "local_preference": "200", "prefix": "172.17.0.0/16"}], "virtual_connection_id": "3c265a62-91da-4261-a950-950b6af0eb58", "virtual_connection_name": "vpc1", "virtual_connection_type": "vpc"}]}]}`)
				}))
			})
			It(`Invoke ListGatewayRouteReports successfully with retries`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())
				directLinkService.EnableRetries(0, 0)

				// Construct an instance of the ListGatewayRouteReportsOptions model
				listGatewayRouteReportsOptionsModel := new(directlinkv1.ListGatewayRouteReportsOptions)
				listGatewayRouteReportsOptionsModel.GatewayID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				listGatewayRouteReportsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := directLinkService.ListGatewayRouteReportsWithContext(ctx, listGatewayRouteReportsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				directLinkService.DisableRetries()
				result, response, operationErr := directLinkService.ListGatewayRouteReports(listGatewayRouteReportsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = directLinkService.ListGatewayRouteReportsWithContext(ctx, listGatewayRouteReportsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listGatewayRouteReportsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"route_reports": [{"advertised_routes": [{"as_path": "64999 64999 64998 I", "prefix": "172.17.0.0/16"}], "created_at": "2019-01-01T12:00:00.000Z", "gateway_routes": [{"prefix": "172.17.0.0/16"}], "id": "1a15dcab-7e26-45e1-b7c5-bc690eaa9724", "on_prem_routes": [{"as_path": "64999 64999 64998 I", "next_hop": "172.17.0.0", "prefix": "172.17.0.0/16"}], "overlapping_routes": [{"routes": [{"prefix": "172.17.0.0/16", "type": "virtual_connection", "virtual_connection_id": "d2d985d8-1d8e-4e8b-96cd-cee2290ecaff"}]}], "status": "complete", "updated_at": "2019-01-01T12:00:00.000Z", "virtual_connection_routes": [{"routes": [{"active": true, "local_preference": "200", "prefix": "172.17.0.0/16"}], "virtual_connection_id": "3c265a62-91da-4261-a950-950b6af0eb58", "virtual_connection_name": "vpc1", "virtual_connection_type": "vpc"}]}]}`)
				}))
			})
			It(`Invoke ListGatewayRouteReports successfully`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := directLinkService.ListGatewayRouteReports(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListGatewayRouteReportsOptions model
				listGatewayRouteReportsOptionsModel := new(directlinkv1.ListGatewayRouteReportsOptions)
				listGatewayRouteReportsOptionsModel.GatewayID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				listGatewayRouteReportsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = directLinkService.ListGatewayRouteReports(listGatewayRouteReportsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListGatewayRouteReports with error: Operation validation and request error`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Construct an instance of the ListGatewayRouteReportsOptions model
				listGatewayRouteReportsOptionsModel := new(directlinkv1.ListGatewayRouteReportsOptions)
				listGatewayRouteReportsOptionsModel.GatewayID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				listGatewayRouteReportsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := directLinkService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := directLinkService.ListGatewayRouteReports(listGatewayRouteReportsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListGatewayRouteReportsOptions model with no property values
				listGatewayRouteReportsOptionsModelNew := new(directlinkv1.ListGatewayRouteReportsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = directLinkService.ListGatewayRouteReports(listGatewayRouteReportsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ListGatewayRouteReports successfully`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Construct an instance of the ListGatewayRouteReportsOptions model
				listGatewayRouteReportsOptionsModel := new(directlinkv1.ListGatewayRouteReportsOptions)
				listGatewayRouteReportsOptionsModel.GatewayID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				listGatewayRouteReportsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := directLinkService.ListGatewayRouteReports(listGatewayRouteReportsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateGatewayRouteReport(createGatewayRouteReportOptions *CreateGatewayRouteReportOptions) - Operation response error`, func() {
		version := "testString"
		createGatewayRouteReportPath := "/gateways/0a06fb9b-820f-4c44-8a31-77f1f0806d28/route_reports"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createGatewayRouteReportPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateGatewayRouteReport with error: Operation response processing error`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Construct an instance of the CreateGatewayRouteReportOptions model
				createGatewayRouteReportOptionsModel := new(directlinkv1.CreateGatewayRouteReportOptions)
				createGatewayRouteReportOptionsModel.GatewayID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				createGatewayRouteReportOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := directLinkService.CreateGatewayRouteReport(createGatewayRouteReportOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				directLinkService.EnableRetries(0, 0)
				result, response, operationErr = directLinkService.CreateGatewayRouteReport(createGatewayRouteReportOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateGatewayRouteReport(createGatewayRouteReportOptions *CreateGatewayRouteReportOptions)`, func() {
		version := "testString"
		createGatewayRouteReportPath := "/gateways/0a06fb9b-820f-4c44-8a31-77f1f0806d28/route_reports"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createGatewayRouteReportPath))
					Expect(req.Method).To(Equal("POST"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"advertised_routes": [{"as_path": "64999 64999 64998 I", "prefix": "172.17.0.0/16"}], "created_at": "2019-01-01T12:00:00.000Z", "gateway_routes": [{"prefix": "172.17.0.0/16"}], "id": "1a15dcab-7e26-45e1-b7c5-bc690eaa9724", "on_prem_routes": [{"as_path": "64999 64999 64998 I", "next_hop": "172.17.0.0", "prefix": "172.17.0.0/16"}], "overlapping_routes": [{"routes": [{"prefix": "172.17.0.0/16", "type": "virtual_connection", "virtual_connection_id": "d2d985d8-1d8e-4e8b-96cd-cee2290ecaff"}]}], "status": "complete", "updated_at": "2019-01-01T12:00:00.000Z", "virtual_connection_routes": [{"routes": [{"active": true, "local_preference": "200", "prefix": "172.17.0.0/16"}], "virtual_connection_id": "3c265a62-91da-4261-a950-950b6af0eb58", "virtual_connection_name": "vpc1", "virtual_connection_type": "vpc"}]}`)
				}))
			})
			It(`Invoke CreateGatewayRouteReport successfully with retries`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())
				directLinkService.EnableRetries(0, 0)

				// Construct an instance of the CreateGatewayRouteReportOptions model
				createGatewayRouteReportOptionsModel := new(directlinkv1.CreateGatewayRouteReportOptions)
				createGatewayRouteReportOptionsModel.GatewayID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				createGatewayRouteReportOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := directLinkService.CreateGatewayRouteReportWithContext(ctx, createGatewayRouteReportOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				directLinkService.DisableRetries()
				result, response, operationErr := directLinkService.CreateGatewayRouteReport(createGatewayRouteReportOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = directLinkService.CreateGatewayRouteReportWithContext(ctx, createGatewayRouteReportOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createGatewayRouteReportPath))
					Expect(req.Method).To(Equal("POST"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"advertised_routes": [{"as_path": "64999 64999 64998 I", "prefix": "172.17.0.0/16"}], "created_at": "2019-01-01T12:00:00.000Z", "gateway_routes": [{"prefix": "172.17.0.0/16"}], "id": "1a15dcab-7e26-45e1-b7c5-bc690eaa9724", "on_prem_routes": [{"as_path": "64999 64999 64998 I", "next_hop": "172.17.0.0", "prefix": "172.17.0.0/16"}], "overlapping_routes": [{"routes": [{"prefix": "172.17.0.0/16", "type": "virtual_connection", "virtual_connection_id": "d2d985d8-1d8e-4e8b-96cd-cee2290ecaff"}]}], "status": "complete", "updated_at": "2019-01-01T12:00:00.000Z", "virtual_connection_routes": [{"routes": [{"active": true, "local_preference": "200", "prefix": "172.17.0.0/16"}], "virtual_connection_id": "3c265a62-91da-4261-a950-950b6af0eb58", "virtual_connection_name": "vpc1", "virtual_connection_type": "vpc"}]}`)
				}))
			})
			It(`Invoke CreateGatewayRouteReport successfully`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := directLinkService.CreateGatewayRouteReport(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateGatewayRouteReportOptions model
				createGatewayRouteReportOptionsModel := new(directlinkv1.CreateGatewayRouteReportOptions)
				createGatewayRouteReportOptionsModel.GatewayID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				createGatewayRouteReportOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = directLinkService.CreateGatewayRouteReport(createGatewayRouteReportOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateGatewayRouteReport with error: Operation validation and request error`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Construct an instance of the CreateGatewayRouteReportOptions model
				createGatewayRouteReportOptionsModel := new(directlinkv1.CreateGatewayRouteReportOptions)
				createGatewayRouteReportOptionsModel.GatewayID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				createGatewayRouteReportOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := directLinkService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := directLinkService.CreateGatewayRouteReport(createGatewayRouteReportOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateGatewayRouteReportOptions model with no property values
				createGatewayRouteReportOptionsModelNew := new(directlinkv1.CreateGatewayRouteReportOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = directLinkService.CreateGatewayRouteReport(createGatewayRouteReportOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(202)
				}))
			})
			It(`Invoke CreateGatewayRouteReport successfully`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Construct an instance of the CreateGatewayRouteReportOptions model
				createGatewayRouteReportOptionsModel := new(directlinkv1.CreateGatewayRouteReportOptions)
				createGatewayRouteReportOptionsModel.GatewayID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				createGatewayRouteReportOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := directLinkService.CreateGatewayRouteReport(createGatewayRouteReportOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteGatewayRouteReport(deleteGatewayRouteReportOptions *DeleteGatewayRouteReportOptions)`, func() {
		version := "testString"
		deleteGatewayRouteReportPath := "/gateways/0a06fb9b-820f-4c44-8a31-77f1f0806d28/route_reports/0a06fb9b-820f-4c44-8a31-77f1f0806d28"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteGatewayRouteReportPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteGatewayRouteReport successfully`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := directLinkService.DeleteGatewayRouteReport(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteGatewayRouteReportOptions model
				deleteGatewayRouteReportOptionsModel := new(directlinkv1.DeleteGatewayRouteReportOptions)
				deleteGatewayRouteReportOptionsModel.GatewayID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				deleteGatewayRouteReportOptionsModel.ID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				deleteGatewayRouteReportOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = directLinkService.DeleteGatewayRouteReport(deleteGatewayRouteReportOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteGatewayRouteReport with error: Operation validation and request error`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Construct an instance of the DeleteGatewayRouteReportOptions model
				deleteGatewayRouteReportOptionsModel := new(directlinkv1.DeleteGatewayRouteReportOptions)
				deleteGatewayRouteReportOptionsModel.GatewayID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				deleteGatewayRouteReportOptionsModel.ID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				deleteGatewayRouteReportOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := directLinkService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := directLinkService.DeleteGatewayRouteReport(deleteGatewayRouteReportOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteGatewayRouteReportOptions model with no property values
				deleteGatewayRouteReportOptionsModelNew := new(directlinkv1.DeleteGatewayRouteReportOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = directLinkService.DeleteGatewayRouteReport(deleteGatewayRouteReportOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetGatewayRouteReport(getGatewayRouteReportOptions *GetGatewayRouteReportOptions) - Operation response error`, func() {
		version := "testString"
		getGatewayRouteReportPath := "/gateways/0a06fb9b-820f-4c44-8a31-77f1f0806d28/route_reports/0a06fb9b-820f-4c44-8a31-77f1f0806d28"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getGatewayRouteReportPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetGatewayRouteReport with error: Operation response processing error`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Construct an instance of the GetGatewayRouteReportOptions model
				getGatewayRouteReportOptionsModel := new(directlinkv1.GetGatewayRouteReportOptions)
				getGatewayRouteReportOptionsModel.GatewayID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				getGatewayRouteReportOptionsModel.ID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				getGatewayRouteReportOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := directLinkService.GetGatewayRouteReport(getGatewayRouteReportOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				directLinkService.EnableRetries(0, 0)
				result, response, operationErr = directLinkService.GetGatewayRouteReport(getGatewayRouteReportOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetGatewayRouteReport(getGatewayRouteReportOptions *GetGatewayRouteReportOptions)`, func() {
		version := "testString"
		getGatewayRouteReportPath := "/gateways/0a06fb9b-820f-4c44-8a31-77f1f0806d28/route_reports/0a06fb9b-820f-4c44-8a31-77f1f0806d28"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getGatewayRouteReportPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"advertised_routes": [{"as_path": "64999 64999 64998 I", "prefix": "172.17.0.0/16"}], "created_at": "2019-01-01T12:00:00.000Z", "gateway_routes": [{"prefix": "172.17.0.0/16"}], "id": "1a15dcab-7e26-45e1-b7c5-bc690eaa9724", "on_prem_routes": [{"as_path": "64999 64999 64998 I", "next_hop": "172.17.0.0", "prefix": "172.17.0.0/16"}], "overlapping_routes": [{"routes": [{"prefix": "172.17.0.0/16", "type": "virtual_connection", "virtual_connection_id": "d2d985d8-1d8e-4e8b-96cd-cee2290ecaff"}]}], "status": "complete", "updated_at": "2019-01-01T12:00:00.000Z", "virtual_connection_routes": [{"routes": [{"active": true, "local_preference": "200", "prefix": "172.17.0.0/16"}], "virtual_connection_id": "3c265a62-91da-4261-a950-950b6af0eb58", "virtual_connection_name": "vpc1", "virtual_connection_type": "vpc"}]}`)
				}))
			})
			It(`Invoke GetGatewayRouteReport successfully with retries`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())
				directLinkService.EnableRetries(0, 0)

				// Construct an instance of the GetGatewayRouteReportOptions model
				getGatewayRouteReportOptionsModel := new(directlinkv1.GetGatewayRouteReportOptions)
				getGatewayRouteReportOptionsModel.GatewayID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				getGatewayRouteReportOptionsModel.ID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				getGatewayRouteReportOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := directLinkService.GetGatewayRouteReportWithContext(ctx, getGatewayRouteReportOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				directLinkService.DisableRetries()
				result, response, operationErr := directLinkService.GetGatewayRouteReport(getGatewayRouteReportOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = directLinkService.GetGatewayRouteReportWithContext(ctx, getGatewayRouteReportOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getGatewayRouteReportPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"advertised_routes": [{"as_path": "64999 64999 64998 I", "prefix": "172.17.0.0/16"}], "created_at": "2019-01-01T12:00:00.000Z", "gateway_routes": [{"prefix": "172.17.0.0/16"}], "id": "1a15dcab-7e26-45e1-b7c5-bc690eaa9724", "on_prem_routes": [{"as_path": "64999 64999 64998 I", "next_hop": "172.17.0.0", "prefix": "172.17.0.0/16"}], "overlapping_routes": [{"routes": [{"prefix": "172.17.0.0/16", "type": "virtual_connection", "virtual_connection_id": "d2d985d8-1d8e-4e8b-96cd-cee2290ecaff"}]}], "status": "complete", "updated_at": "2019-01-01T12:00:00.000Z", "virtual_connection_routes": [{"routes": [{"active": true, "local_preference": "200", "prefix": "172.17.0.0/16"}], "virtual_connection_id": "3c265a62-91da-4261-a950-950b6af0eb58", "virtual_connection_name": "vpc1", "virtual_connection_type": "vpc"}]}`)
				}))
			})
			It(`Invoke GetGatewayRouteReport successfully`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := directLinkService.GetGatewayRouteReport(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetGatewayRouteReportOptions model
				getGatewayRouteReportOptionsModel := new(directlinkv1.GetGatewayRouteReportOptions)
				getGatewayRouteReportOptionsModel.GatewayID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				getGatewayRouteReportOptionsModel.ID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				getGatewayRouteReportOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = directLinkService.GetGatewayRouteReport(getGatewayRouteReportOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetGatewayRouteReport with error: Operation validation and request error`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Construct an instance of the GetGatewayRouteReportOptions model
				getGatewayRouteReportOptionsModel := new(directlinkv1.GetGatewayRouteReportOptions)
				getGatewayRouteReportOptionsModel.GatewayID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				getGatewayRouteReportOptionsModel.ID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				getGatewayRouteReportOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := directLinkService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := directLinkService.GetGatewayRouteReport(getGatewayRouteReportOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetGatewayRouteReportOptions model with no property values
				getGatewayRouteReportOptionsModelNew := new(directlinkv1.GetGatewayRouteReportOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = directLinkService.GetGatewayRouteReport(getGatewayRouteReportOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetGatewayRouteReport successfully`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Construct an instance of the GetGatewayRouteReportOptions model
				getGatewayRouteReportOptionsModel := new(directlinkv1.GetGatewayRouteReportOptions)
				getGatewayRouteReportOptionsModel.GatewayID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				getGatewayRouteReportOptionsModel.ID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				getGatewayRouteReportOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := directLinkService.GetGatewayRouteReport(getGatewayRouteReportOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListGatewayVirtualConnections(listGatewayVirtualConnectionsOptions *ListGatewayVirtualConnectionsOptions) - Operation response error`, func() {
		version := "testString"
		listGatewayVirtualConnectionsPath := "/gateways/0a06fb9b-820f-4c44-8a31-77f1f0806d28/virtual_connections"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listGatewayVirtualConnectionsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListGatewayVirtualConnections with error: Operation response processing error`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Construct an instance of the ListGatewayVirtualConnectionsOptions model
				listGatewayVirtualConnectionsOptionsModel := new(directlinkv1.ListGatewayVirtualConnectionsOptions)
				listGatewayVirtualConnectionsOptionsModel.GatewayID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				listGatewayVirtualConnectionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := directLinkService.ListGatewayVirtualConnections(listGatewayVirtualConnectionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				directLinkService.EnableRetries(0, 0)
				result, response, operationErr = directLinkService.ListGatewayVirtualConnections(listGatewayVirtualConnectionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListGatewayVirtualConnections(listGatewayVirtualConnectionsOptions *ListGatewayVirtualConnectionsOptions)`, func() {
		version := "testString"
		listGatewayVirtualConnectionsPath := "/gateways/0a06fb9b-820f-4c44-8a31-77f1f0806d28/virtual_connections"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listGatewayVirtualConnectionsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"virtual_connections": [{"created_at": "2019-01-01T12:00:00.000Z", "id": "ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4", "name": "newVC", "network_account": "00aa14a2e0fb102c8995ebefff865555", "network_id": "crn:v1:bluemix:public:is:us-east:a/28e4d90ac7504be69447111122223333::vpc:aaa81ac8-5e96-42a0-a4b7-6c2e2d1bbbbb", "status": "attached", "type": "vpc"}]}`)
				}))
			})
			It(`Invoke ListGatewayVirtualConnections successfully with retries`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())
				directLinkService.EnableRetries(0, 0)

				// Construct an instance of the ListGatewayVirtualConnectionsOptions model
				listGatewayVirtualConnectionsOptionsModel := new(directlinkv1.ListGatewayVirtualConnectionsOptions)
				listGatewayVirtualConnectionsOptionsModel.GatewayID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				listGatewayVirtualConnectionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := directLinkService.ListGatewayVirtualConnectionsWithContext(ctx, listGatewayVirtualConnectionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				directLinkService.DisableRetries()
				result, response, operationErr := directLinkService.ListGatewayVirtualConnections(listGatewayVirtualConnectionsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = directLinkService.ListGatewayVirtualConnectionsWithContext(ctx, listGatewayVirtualConnectionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listGatewayVirtualConnectionsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"virtual_connections": [{"created_at": "2019-01-01T12:00:00.000Z", "id": "ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4", "name": "newVC", "network_account": "00aa14a2e0fb102c8995ebefff865555", "network_id": "crn:v1:bluemix:public:is:us-east:a/28e4d90ac7504be69447111122223333::vpc:aaa81ac8-5e96-42a0-a4b7-6c2e2d1bbbbb", "status": "attached", "type": "vpc"}]}`)
				}))
			})
			It(`Invoke ListGatewayVirtualConnections successfully`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := directLinkService.ListGatewayVirtualConnections(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListGatewayVirtualConnectionsOptions model
				listGatewayVirtualConnectionsOptionsModel := new(directlinkv1.ListGatewayVirtualConnectionsOptions)
				listGatewayVirtualConnectionsOptionsModel.GatewayID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				listGatewayVirtualConnectionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = directLinkService.ListGatewayVirtualConnections(listGatewayVirtualConnectionsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListGatewayVirtualConnections with error: Operation validation and request error`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Construct an instance of the ListGatewayVirtualConnectionsOptions model
				listGatewayVirtualConnectionsOptionsModel := new(directlinkv1.ListGatewayVirtualConnectionsOptions)
				listGatewayVirtualConnectionsOptionsModel.GatewayID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				listGatewayVirtualConnectionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := directLinkService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := directLinkService.ListGatewayVirtualConnections(listGatewayVirtualConnectionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListGatewayVirtualConnectionsOptions model with no property values
				listGatewayVirtualConnectionsOptionsModelNew := new(directlinkv1.ListGatewayVirtualConnectionsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = directLinkService.ListGatewayVirtualConnections(listGatewayVirtualConnectionsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ListGatewayVirtualConnections successfully`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Construct an instance of the ListGatewayVirtualConnectionsOptions model
				listGatewayVirtualConnectionsOptionsModel := new(directlinkv1.ListGatewayVirtualConnectionsOptions)
				listGatewayVirtualConnectionsOptionsModel.GatewayID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				listGatewayVirtualConnectionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := directLinkService.ListGatewayVirtualConnections(listGatewayVirtualConnectionsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateGatewayVirtualConnection(createGatewayVirtualConnectionOptions *CreateGatewayVirtualConnectionOptions) - Operation response error`, func() {
		version := "testString"
		createGatewayVirtualConnectionPath := "/gateways/0a06fb9b-820f-4c44-8a31-77f1f0806d28/virtual_connections"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createGatewayVirtualConnectionPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateGatewayVirtualConnection with error: Operation response processing error`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Construct an instance of the CreateGatewayVirtualConnectionOptions model
				createGatewayVirtualConnectionOptionsModel := new(directlinkv1.CreateGatewayVirtualConnectionOptions)
				createGatewayVirtualConnectionOptionsModel.GatewayID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				createGatewayVirtualConnectionOptionsModel.Name = core.StringPtr("newVC")
				createGatewayVirtualConnectionOptionsModel.Type = core.StringPtr("vpc")
				createGatewayVirtualConnectionOptionsModel.NetworkID = core.StringPtr("crn:v1:bluemix:public:is:us-east:a/28e4d90ac7504be69447111122223333::vpc:aaa81ac8-5e96-42a0-a4b7-6c2e2d1bbbbb")
				createGatewayVirtualConnectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := directLinkService.CreateGatewayVirtualConnection(createGatewayVirtualConnectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				directLinkService.EnableRetries(0, 0)
				result, response, operationErr = directLinkService.CreateGatewayVirtualConnection(createGatewayVirtualConnectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateGatewayVirtualConnection(createGatewayVirtualConnectionOptions *CreateGatewayVirtualConnectionOptions)`, func() {
		version := "testString"
		createGatewayVirtualConnectionPath := "/gateways/0a06fb9b-820f-4c44-8a31-77f1f0806d28/virtual_connections"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createGatewayVirtualConnectionPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"created_at": "2019-01-01T12:00:00.000Z", "id": "ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4", "name": "newVC", "network_account": "00aa14a2e0fb102c8995ebefff865555", "network_id": "crn:v1:bluemix:public:is:us-east:a/28e4d90ac7504be69447111122223333::vpc:aaa81ac8-5e96-42a0-a4b7-6c2e2d1bbbbb", "status": "attached", "type": "vpc"}`)
				}))
			})
			It(`Invoke CreateGatewayVirtualConnection successfully with retries`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())
				directLinkService.EnableRetries(0, 0)

				// Construct an instance of the CreateGatewayVirtualConnectionOptions model
				createGatewayVirtualConnectionOptionsModel := new(directlinkv1.CreateGatewayVirtualConnectionOptions)
				createGatewayVirtualConnectionOptionsModel.GatewayID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				createGatewayVirtualConnectionOptionsModel.Name = core.StringPtr("newVC")
				createGatewayVirtualConnectionOptionsModel.Type = core.StringPtr("vpc")
				createGatewayVirtualConnectionOptionsModel.NetworkID = core.StringPtr("crn:v1:bluemix:public:is:us-east:a/28e4d90ac7504be69447111122223333::vpc:aaa81ac8-5e96-42a0-a4b7-6c2e2d1bbbbb")
				createGatewayVirtualConnectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := directLinkService.CreateGatewayVirtualConnectionWithContext(ctx, createGatewayVirtualConnectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				directLinkService.DisableRetries()
				result, response, operationErr := directLinkService.CreateGatewayVirtualConnection(createGatewayVirtualConnectionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = directLinkService.CreateGatewayVirtualConnectionWithContext(ctx, createGatewayVirtualConnectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createGatewayVirtualConnectionPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"created_at": "2019-01-01T12:00:00.000Z", "id": "ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4", "name": "newVC", "network_account": "00aa14a2e0fb102c8995ebefff865555", "network_id": "crn:v1:bluemix:public:is:us-east:a/28e4d90ac7504be69447111122223333::vpc:aaa81ac8-5e96-42a0-a4b7-6c2e2d1bbbbb", "status": "attached", "type": "vpc"}`)
				}))
			})
			It(`Invoke CreateGatewayVirtualConnection successfully`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := directLinkService.CreateGatewayVirtualConnection(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateGatewayVirtualConnectionOptions model
				createGatewayVirtualConnectionOptionsModel := new(directlinkv1.CreateGatewayVirtualConnectionOptions)
				createGatewayVirtualConnectionOptionsModel.GatewayID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				createGatewayVirtualConnectionOptionsModel.Name = core.StringPtr("newVC")
				createGatewayVirtualConnectionOptionsModel.Type = core.StringPtr("vpc")
				createGatewayVirtualConnectionOptionsModel.NetworkID = core.StringPtr("crn:v1:bluemix:public:is:us-east:a/28e4d90ac7504be69447111122223333::vpc:aaa81ac8-5e96-42a0-a4b7-6c2e2d1bbbbb")
				createGatewayVirtualConnectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = directLinkService.CreateGatewayVirtualConnection(createGatewayVirtualConnectionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateGatewayVirtualConnection with error: Operation validation and request error`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Construct an instance of the CreateGatewayVirtualConnectionOptions model
				createGatewayVirtualConnectionOptionsModel := new(directlinkv1.CreateGatewayVirtualConnectionOptions)
				createGatewayVirtualConnectionOptionsModel.GatewayID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				createGatewayVirtualConnectionOptionsModel.Name = core.StringPtr("newVC")
				createGatewayVirtualConnectionOptionsModel.Type = core.StringPtr("vpc")
				createGatewayVirtualConnectionOptionsModel.NetworkID = core.StringPtr("crn:v1:bluemix:public:is:us-east:a/28e4d90ac7504be69447111122223333::vpc:aaa81ac8-5e96-42a0-a4b7-6c2e2d1bbbbb")
				createGatewayVirtualConnectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := directLinkService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := directLinkService.CreateGatewayVirtualConnection(createGatewayVirtualConnectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateGatewayVirtualConnectionOptions model with no property values
				createGatewayVirtualConnectionOptionsModelNew := new(directlinkv1.CreateGatewayVirtualConnectionOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = directLinkService.CreateGatewayVirtualConnection(createGatewayVirtualConnectionOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(201)
				}))
			})
			It(`Invoke CreateGatewayVirtualConnection successfully`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Construct an instance of the CreateGatewayVirtualConnectionOptions model
				createGatewayVirtualConnectionOptionsModel := new(directlinkv1.CreateGatewayVirtualConnectionOptions)
				createGatewayVirtualConnectionOptionsModel.GatewayID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				createGatewayVirtualConnectionOptionsModel.Name = core.StringPtr("newVC")
				createGatewayVirtualConnectionOptionsModel.Type = core.StringPtr("vpc")
				createGatewayVirtualConnectionOptionsModel.NetworkID = core.StringPtr("crn:v1:bluemix:public:is:us-east:a/28e4d90ac7504be69447111122223333::vpc:aaa81ac8-5e96-42a0-a4b7-6c2e2d1bbbbb")
				createGatewayVirtualConnectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := directLinkService.CreateGatewayVirtualConnection(createGatewayVirtualConnectionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteGatewayVirtualConnection(deleteGatewayVirtualConnectionOptions *DeleteGatewayVirtualConnectionOptions)`, func() {
		version := "testString"
		deleteGatewayVirtualConnectionPath := "/gateways/0a06fb9b-820f-4c44-8a31-77f1f0806d28/virtual_connections/0a06fb9b-820f-4c44-8a31-77f1f0806d28"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteGatewayVirtualConnectionPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteGatewayVirtualConnection successfully`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := directLinkService.DeleteGatewayVirtualConnection(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteGatewayVirtualConnectionOptions model
				deleteGatewayVirtualConnectionOptionsModel := new(directlinkv1.DeleteGatewayVirtualConnectionOptions)
				deleteGatewayVirtualConnectionOptionsModel.GatewayID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				deleteGatewayVirtualConnectionOptionsModel.ID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				deleteGatewayVirtualConnectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = directLinkService.DeleteGatewayVirtualConnection(deleteGatewayVirtualConnectionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteGatewayVirtualConnection with error: Operation validation and request error`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Construct an instance of the DeleteGatewayVirtualConnectionOptions model
				deleteGatewayVirtualConnectionOptionsModel := new(directlinkv1.DeleteGatewayVirtualConnectionOptions)
				deleteGatewayVirtualConnectionOptionsModel.GatewayID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				deleteGatewayVirtualConnectionOptionsModel.ID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				deleteGatewayVirtualConnectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := directLinkService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := directLinkService.DeleteGatewayVirtualConnection(deleteGatewayVirtualConnectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteGatewayVirtualConnectionOptions model with no property values
				deleteGatewayVirtualConnectionOptionsModelNew := new(directlinkv1.DeleteGatewayVirtualConnectionOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = directLinkService.DeleteGatewayVirtualConnection(deleteGatewayVirtualConnectionOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetGatewayVirtualConnection(getGatewayVirtualConnectionOptions *GetGatewayVirtualConnectionOptions) - Operation response error`, func() {
		version := "testString"
		getGatewayVirtualConnectionPath := "/gateways/0a06fb9b-820f-4c44-8a31-77f1f0806d28/virtual_connections/0a06fb9b-820f-4c44-8a31-77f1f0806d28"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getGatewayVirtualConnectionPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetGatewayVirtualConnection with error: Operation response processing error`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Construct an instance of the GetGatewayVirtualConnectionOptions model
				getGatewayVirtualConnectionOptionsModel := new(directlinkv1.GetGatewayVirtualConnectionOptions)
				getGatewayVirtualConnectionOptionsModel.GatewayID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				getGatewayVirtualConnectionOptionsModel.ID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				getGatewayVirtualConnectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := directLinkService.GetGatewayVirtualConnection(getGatewayVirtualConnectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				directLinkService.EnableRetries(0, 0)
				result, response, operationErr = directLinkService.GetGatewayVirtualConnection(getGatewayVirtualConnectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetGatewayVirtualConnection(getGatewayVirtualConnectionOptions *GetGatewayVirtualConnectionOptions)`, func() {
		version := "testString"
		getGatewayVirtualConnectionPath := "/gateways/0a06fb9b-820f-4c44-8a31-77f1f0806d28/virtual_connections/0a06fb9b-820f-4c44-8a31-77f1f0806d28"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getGatewayVirtualConnectionPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"created_at": "2019-01-01T12:00:00.000Z", "id": "ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4", "name": "newVC", "network_account": "00aa14a2e0fb102c8995ebefff865555", "network_id": "crn:v1:bluemix:public:is:us-east:a/28e4d90ac7504be69447111122223333::vpc:aaa81ac8-5e96-42a0-a4b7-6c2e2d1bbbbb", "status": "attached", "type": "vpc"}`)
				}))
			})
			It(`Invoke GetGatewayVirtualConnection successfully with retries`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())
				directLinkService.EnableRetries(0, 0)

				// Construct an instance of the GetGatewayVirtualConnectionOptions model
				getGatewayVirtualConnectionOptionsModel := new(directlinkv1.GetGatewayVirtualConnectionOptions)
				getGatewayVirtualConnectionOptionsModel.GatewayID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				getGatewayVirtualConnectionOptionsModel.ID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				getGatewayVirtualConnectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := directLinkService.GetGatewayVirtualConnectionWithContext(ctx, getGatewayVirtualConnectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				directLinkService.DisableRetries()
				result, response, operationErr := directLinkService.GetGatewayVirtualConnection(getGatewayVirtualConnectionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = directLinkService.GetGatewayVirtualConnectionWithContext(ctx, getGatewayVirtualConnectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getGatewayVirtualConnectionPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"created_at": "2019-01-01T12:00:00.000Z", "id": "ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4", "name": "newVC", "network_account": "00aa14a2e0fb102c8995ebefff865555", "network_id": "crn:v1:bluemix:public:is:us-east:a/28e4d90ac7504be69447111122223333::vpc:aaa81ac8-5e96-42a0-a4b7-6c2e2d1bbbbb", "status": "attached", "type": "vpc"}`)
				}))
			})
			It(`Invoke GetGatewayVirtualConnection successfully`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := directLinkService.GetGatewayVirtualConnection(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetGatewayVirtualConnectionOptions model
				getGatewayVirtualConnectionOptionsModel := new(directlinkv1.GetGatewayVirtualConnectionOptions)
				getGatewayVirtualConnectionOptionsModel.GatewayID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				getGatewayVirtualConnectionOptionsModel.ID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				getGatewayVirtualConnectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = directLinkService.GetGatewayVirtualConnection(getGatewayVirtualConnectionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetGatewayVirtualConnection with error: Operation validation and request error`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Construct an instance of the GetGatewayVirtualConnectionOptions model
				getGatewayVirtualConnectionOptionsModel := new(directlinkv1.GetGatewayVirtualConnectionOptions)
				getGatewayVirtualConnectionOptionsModel.GatewayID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				getGatewayVirtualConnectionOptionsModel.ID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				getGatewayVirtualConnectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := directLinkService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := directLinkService.GetGatewayVirtualConnection(getGatewayVirtualConnectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetGatewayVirtualConnectionOptions model with no property values
				getGatewayVirtualConnectionOptionsModelNew := new(directlinkv1.GetGatewayVirtualConnectionOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = directLinkService.GetGatewayVirtualConnection(getGatewayVirtualConnectionOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetGatewayVirtualConnection successfully`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Construct an instance of the GetGatewayVirtualConnectionOptions model
				getGatewayVirtualConnectionOptionsModel := new(directlinkv1.GetGatewayVirtualConnectionOptions)
				getGatewayVirtualConnectionOptionsModel.GatewayID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				getGatewayVirtualConnectionOptionsModel.ID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				getGatewayVirtualConnectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := directLinkService.GetGatewayVirtualConnection(getGatewayVirtualConnectionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateGatewayVirtualConnection(updateGatewayVirtualConnectionOptions *UpdateGatewayVirtualConnectionOptions) - Operation response error`, func() {
		version := "testString"
		updateGatewayVirtualConnectionPath := "/gateways/0a06fb9b-820f-4c44-8a31-77f1f0806d28/virtual_connections/0a06fb9b-820f-4c44-8a31-77f1f0806d28"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateGatewayVirtualConnectionPath))
					Expect(req.Method).To(Equal("PATCH"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateGatewayVirtualConnection with error: Operation response processing error`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Construct an instance of the GatewayVirtualConnectionPatchTemplate model
				gatewayVirtualConnectionPatchTemplateModel := new(directlinkv1.GatewayVirtualConnectionPatchTemplate)
				gatewayVirtualConnectionPatchTemplateModel.Name = core.StringPtr("newConnectionName")
				gatewayVirtualConnectionPatchTemplateModel.Status = core.StringPtr("attached")
				gatewayVirtualConnectionPatchTemplateModelAsPatch, asPatchErr := gatewayVirtualConnectionPatchTemplateModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateGatewayVirtualConnectionOptions model
				updateGatewayVirtualConnectionOptionsModel := new(directlinkv1.UpdateGatewayVirtualConnectionOptions)
				updateGatewayVirtualConnectionOptionsModel.GatewayID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				updateGatewayVirtualConnectionOptionsModel.ID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				updateGatewayVirtualConnectionOptionsModel.GatewayVirtualConnectionPatchTemplatePatch = gatewayVirtualConnectionPatchTemplateModelAsPatch
				updateGatewayVirtualConnectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := directLinkService.UpdateGatewayVirtualConnection(updateGatewayVirtualConnectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				directLinkService.EnableRetries(0, 0)
				result, response, operationErr = directLinkService.UpdateGatewayVirtualConnection(updateGatewayVirtualConnectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateGatewayVirtualConnection(updateGatewayVirtualConnectionOptions *UpdateGatewayVirtualConnectionOptions)`, func() {
		version := "testString"
		updateGatewayVirtualConnectionPath := "/gateways/0a06fb9b-820f-4c44-8a31-77f1f0806d28/virtual_connections/0a06fb9b-820f-4c44-8a31-77f1f0806d28"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateGatewayVirtualConnectionPath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"created_at": "2019-01-01T12:00:00.000Z", "id": "ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4", "name": "newVC", "network_account": "00aa14a2e0fb102c8995ebefff865555", "network_id": "crn:v1:bluemix:public:is:us-east:a/28e4d90ac7504be69447111122223333::vpc:aaa81ac8-5e96-42a0-a4b7-6c2e2d1bbbbb", "status": "attached", "type": "vpc"}`)
				}))
			})
			It(`Invoke UpdateGatewayVirtualConnection successfully with retries`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())
				directLinkService.EnableRetries(0, 0)

				// Construct an instance of the GatewayVirtualConnectionPatchTemplate model
				gatewayVirtualConnectionPatchTemplateModel := new(directlinkv1.GatewayVirtualConnectionPatchTemplate)
				gatewayVirtualConnectionPatchTemplateModel.Name = core.StringPtr("newConnectionName")
				gatewayVirtualConnectionPatchTemplateModel.Status = core.StringPtr("attached")
				gatewayVirtualConnectionPatchTemplateModelAsPatch, asPatchErr := gatewayVirtualConnectionPatchTemplateModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateGatewayVirtualConnectionOptions model
				updateGatewayVirtualConnectionOptionsModel := new(directlinkv1.UpdateGatewayVirtualConnectionOptions)
				updateGatewayVirtualConnectionOptionsModel.GatewayID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				updateGatewayVirtualConnectionOptionsModel.ID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				updateGatewayVirtualConnectionOptionsModel.GatewayVirtualConnectionPatchTemplatePatch = gatewayVirtualConnectionPatchTemplateModelAsPatch
				updateGatewayVirtualConnectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := directLinkService.UpdateGatewayVirtualConnectionWithContext(ctx, updateGatewayVirtualConnectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				directLinkService.DisableRetries()
				result, response, operationErr := directLinkService.UpdateGatewayVirtualConnection(updateGatewayVirtualConnectionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = directLinkService.UpdateGatewayVirtualConnectionWithContext(ctx, updateGatewayVirtualConnectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateGatewayVirtualConnectionPath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"created_at": "2019-01-01T12:00:00.000Z", "id": "ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4", "name": "newVC", "network_account": "00aa14a2e0fb102c8995ebefff865555", "network_id": "crn:v1:bluemix:public:is:us-east:a/28e4d90ac7504be69447111122223333::vpc:aaa81ac8-5e96-42a0-a4b7-6c2e2d1bbbbb", "status": "attached", "type": "vpc"}`)
				}))
			})
			It(`Invoke UpdateGatewayVirtualConnection successfully`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := directLinkService.UpdateGatewayVirtualConnection(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GatewayVirtualConnectionPatchTemplate model
				gatewayVirtualConnectionPatchTemplateModel := new(directlinkv1.GatewayVirtualConnectionPatchTemplate)
				gatewayVirtualConnectionPatchTemplateModel.Name = core.StringPtr("newConnectionName")
				gatewayVirtualConnectionPatchTemplateModel.Status = core.StringPtr("attached")
				gatewayVirtualConnectionPatchTemplateModelAsPatch, asPatchErr := gatewayVirtualConnectionPatchTemplateModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateGatewayVirtualConnectionOptions model
				updateGatewayVirtualConnectionOptionsModel := new(directlinkv1.UpdateGatewayVirtualConnectionOptions)
				updateGatewayVirtualConnectionOptionsModel.GatewayID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				updateGatewayVirtualConnectionOptionsModel.ID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				updateGatewayVirtualConnectionOptionsModel.GatewayVirtualConnectionPatchTemplatePatch = gatewayVirtualConnectionPatchTemplateModelAsPatch
				updateGatewayVirtualConnectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = directLinkService.UpdateGatewayVirtualConnection(updateGatewayVirtualConnectionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateGatewayVirtualConnection with error: Operation validation and request error`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Construct an instance of the GatewayVirtualConnectionPatchTemplate model
				gatewayVirtualConnectionPatchTemplateModel := new(directlinkv1.GatewayVirtualConnectionPatchTemplate)
				gatewayVirtualConnectionPatchTemplateModel.Name = core.StringPtr("newConnectionName")
				gatewayVirtualConnectionPatchTemplateModel.Status = core.StringPtr("attached")
				gatewayVirtualConnectionPatchTemplateModelAsPatch, asPatchErr := gatewayVirtualConnectionPatchTemplateModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateGatewayVirtualConnectionOptions model
				updateGatewayVirtualConnectionOptionsModel := new(directlinkv1.UpdateGatewayVirtualConnectionOptions)
				updateGatewayVirtualConnectionOptionsModel.GatewayID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				updateGatewayVirtualConnectionOptionsModel.ID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				updateGatewayVirtualConnectionOptionsModel.GatewayVirtualConnectionPatchTemplatePatch = gatewayVirtualConnectionPatchTemplateModelAsPatch
				updateGatewayVirtualConnectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := directLinkService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := directLinkService.UpdateGatewayVirtualConnection(updateGatewayVirtualConnectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateGatewayVirtualConnectionOptions model with no property values
				updateGatewayVirtualConnectionOptionsModelNew := new(directlinkv1.UpdateGatewayVirtualConnectionOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = directLinkService.UpdateGatewayVirtualConnection(updateGatewayVirtualConnectionOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke UpdateGatewayVirtualConnection successfully`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Construct an instance of the GatewayVirtualConnectionPatchTemplate model
				gatewayVirtualConnectionPatchTemplateModel := new(directlinkv1.GatewayVirtualConnectionPatchTemplate)
				gatewayVirtualConnectionPatchTemplateModel.Name = core.StringPtr("newConnectionName")
				gatewayVirtualConnectionPatchTemplateModel.Status = core.StringPtr("attached")
				gatewayVirtualConnectionPatchTemplateModelAsPatch, asPatchErr := gatewayVirtualConnectionPatchTemplateModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateGatewayVirtualConnectionOptions model
				updateGatewayVirtualConnectionOptionsModel := new(directlinkv1.UpdateGatewayVirtualConnectionOptions)
				updateGatewayVirtualConnectionOptionsModel.GatewayID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				updateGatewayVirtualConnectionOptionsModel.ID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				updateGatewayVirtualConnectionOptionsModel.GatewayVirtualConnectionPatchTemplatePatch = gatewayVirtualConnectionPatchTemplateModelAsPatch
				updateGatewayVirtualConnectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := directLinkService.UpdateGatewayVirtualConnection(updateGatewayVirtualConnectionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListOfferingTypeLocations(listOfferingTypeLocationsOptions *ListOfferingTypeLocationsOptions) - Operation response error`, func() {
		version := "testString"
		listOfferingTypeLocationsPath := "/offering_types/dedicated/locations"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listOfferingTypeLocationsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListOfferingTypeLocations with error: Operation response processing error`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Construct an instance of the ListOfferingTypeLocationsOptions model
				listOfferingTypeLocationsOptionsModel := new(directlinkv1.ListOfferingTypeLocationsOptions)
				listOfferingTypeLocationsOptionsModel.OfferingType = core.StringPtr("dedicated")
				listOfferingTypeLocationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := directLinkService.ListOfferingTypeLocations(listOfferingTypeLocationsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				directLinkService.EnableRetries(0, 0)
				result, response, operationErr = directLinkService.ListOfferingTypeLocations(listOfferingTypeLocationsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListOfferingTypeLocations(listOfferingTypeLocationsOptions *ListOfferingTypeLocationsOptions)`, func() {
		version := "testString"
		listOfferingTypeLocationsPath := "/offering_types/dedicated/locations"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listOfferingTypeLocationsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"locations": [{"billing_location": "us", "building_colocation_owner": "MyProvider", "display_name": "Dallas 9", "location_type": "PoP", "macsec_enabled": false, "market": "Dallas", "market_geography": "N/S America", "mzr": true, "name": "dal03", "offering_type": "dedicated", "provision_enabled": true, "vpc_region": "us-south"}]}`)
				}))
			})
			It(`Invoke ListOfferingTypeLocations successfully with retries`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())
				directLinkService.EnableRetries(0, 0)

				// Construct an instance of the ListOfferingTypeLocationsOptions model
				listOfferingTypeLocationsOptionsModel := new(directlinkv1.ListOfferingTypeLocationsOptions)
				listOfferingTypeLocationsOptionsModel.OfferingType = core.StringPtr("dedicated")
				listOfferingTypeLocationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := directLinkService.ListOfferingTypeLocationsWithContext(ctx, listOfferingTypeLocationsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				directLinkService.DisableRetries()
				result, response, operationErr := directLinkService.ListOfferingTypeLocations(listOfferingTypeLocationsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = directLinkService.ListOfferingTypeLocationsWithContext(ctx, listOfferingTypeLocationsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listOfferingTypeLocationsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"locations": [{"billing_location": "us", "building_colocation_owner": "MyProvider", "display_name": "Dallas 9", "location_type": "PoP", "macsec_enabled": false, "market": "Dallas", "market_geography": "N/S America", "mzr": true, "name": "dal03", "offering_type": "dedicated", "provision_enabled": true, "vpc_region": "us-south"}]}`)
				}))
			})
			It(`Invoke ListOfferingTypeLocations successfully`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := directLinkService.ListOfferingTypeLocations(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListOfferingTypeLocationsOptions model
				listOfferingTypeLocationsOptionsModel := new(directlinkv1.ListOfferingTypeLocationsOptions)
				listOfferingTypeLocationsOptionsModel.OfferingType = core.StringPtr("dedicated")
				listOfferingTypeLocationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = directLinkService.ListOfferingTypeLocations(listOfferingTypeLocationsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListOfferingTypeLocations with error: Operation validation and request error`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Construct an instance of the ListOfferingTypeLocationsOptions model
				listOfferingTypeLocationsOptionsModel := new(directlinkv1.ListOfferingTypeLocationsOptions)
				listOfferingTypeLocationsOptionsModel.OfferingType = core.StringPtr("dedicated")
				listOfferingTypeLocationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := directLinkService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := directLinkService.ListOfferingTypeLocations(listOfferingTypeLocationsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListOfferingTypeLocationsOptions model with no property values
				listOfferingTypeLocationsOptionsModelNew := new(directlinkv1.ListOfferingTypeLocationsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = directLinkService.ListOfferingTypeLocations(listOfferingTypeLocationsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ListOfferingTypeLocations successfully`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Construct an instance of the ListOfferingTypeLocationsOptions model
				listOfferingTypeLocationsOptionsModel := new(directlinkv1.ListOfferingTypeLocationsOptions)
				listOfferingTypeLocationsOptionsModel.OfferingType = core.StringPtr("dedicated")
				listOfferingTypeLocationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := directLinkService.ListOfferingTypeLocations(listOfferingTypeLocationsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListOfferingTypeLocationCrossConnectRouters(listOfferingTypeLocationCrossConnectRoutersOptions *ListOfferingTypeLocationCrossConnectRoutersOptions) - Operation response error`, func() {
		version := "testString"
		listOfferingTypeLocationCrossConnectRoutersPath := "/offering_types/dedicated/locations/testString/cross_connect_routers"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listOfferingTypeLocationCrossConnectRoutersPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListOfferingTypeLocationCrossConnectRouters with error: Operation response processing error`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Construct an instance of the ListOfferingTypeLocationCrossConnectRoutersOptions model
				listOfferingTypeLocationCrossConnectRoutersOptionsModel := new(directlinkv1.ListOfferingTypeLocationCrossConnectRoutersOptions)
				listOfferingTypeLocationCrossConnectRoutersOptionsModel.OfferingType = core.StringPtr("dedicated")
				listOfferingTypeLocationCrossConnectRoutersOptionsModel.LocationName = core.StringPtr("testString")
				listOfferingTypeLocationCrossConnectRoutersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := directLinkService.ListOfferingTypeLocationCrossConnectRouters(listOfferingTypeLocationCrossConnectRoutersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				directLinkService.EnableRetries(0, 0)
				result, response, operationErr = directLinkService.ListOfferingTypeLocationCrossConnectRouters(listOfferingTypeLocationCrossConnectRoutersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListOfferingTypeLocationCrossConnectRouters(listOfferingTypeLocationCrossConnectRoutersOptions *ListOfferingTypeLocationCrossConnectRoutersOptions)`, func() {
		version := "testString"
		listOfferingTypeLocationCrossConnectRoutersPath := "/offering_types/dedicated/locations/testString/cross_connect_routers"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listOfferingTypeLocationCrossConnectRoutersPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"cross_connect_routers": [{"capabilities": ["Capabilities"], "router_name": "xcr01.dal03", "total_connections": 1}]}`)
				}))
			})
			It(`Invoke ListOfferingTypeLocationCrossConnectRouters successfully with retries`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())
				directLinkService.EnableRetries(0, 0)

				// Construct an instance of the ListOfferingTypeLocationCrossConnectRoutersOptions model
				listOfferingTypeLocationCrossConnectRoutersOptionsModel := new(directlinkv1.ListOfferingTypeLocationCrossConnectRoutersOptions)
				listOfferingTypeLocationCrossConnectRoutersOptionsModel.OfferingType = core.StringPtr("dedicated")
				listOfferingTypeLocationCrossConnectRoutersOptionsModel.LocationName = core.StringPtr("testString")
				listOfferingTypeLocationCrossConnectRoutersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := directLinkService.ListOfferingTypeLocationCrossConnectRoutersWithContext(ctx, listOfferingTypeLocationCrossConnectRoutersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				directLinkService.DisableRetries()
				result, response, operationErr := directLinkService.ListOfferingTypeLocationCrossConnectRouters(listOfferingTypeLocationCrossConnectRoutersOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = directLinkService.ListOfferingTypeLocationCrossConnectRoutersWithContext(ctx, listOfferingTypeLocationCrossConnectRoutersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listOfferingTypeLocationCrossConnectRoutersPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"cross_connect_routers": [{"capabilities": ["Capabilities"], "router_name": "xcr01.dal03", "total_connections": 1}]}`)
				}))
			})
			It(`Invoke ListOfferingTypeLocationCrossConnectRouters successfully`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := directLinkService.ListOfferingTypeLocationCrossConnectRouters(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListOfferingTypeLocationCrossConnectRoutersOptions model
				listOfferingTypeLocationCrossConnectRoutersOptionsModel := new(directlinkv1.ListOfferingTypeLocationCrossConnectRoutersOptions)
				listOfferingTypeLocationCrossConnectRoutersOptionsModel.OfferingType = core.StringPtr("dedicated")
				listOfferingTypeLocationCrossConnectRoutersOptionsModel.LocationName = core.StringPtr("testString")
				listOfferingTypeLocationCrossConnectRoutersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = directLinkService.ListOfferingTypeLocationCrossConnectRouters(listOfferingTypeLocationCrossConnectRoutersOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListOfferingTypeLocationCrossConnectRouters with error: Operation validation and request error`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Construct an instance of the ListOfferingTypeLocationCrossConnectRoutersOptions model
				listOfferingTypeLocationCrossConnectRoutersOptionsModel := new(directlinkv1.ListOfferingTypeLocationCrossConnectRoutersOptions)
				listOfferingTypeLocationCrossConnectRoutersOptionsModel.OfferingType = core.StringPtr("dedicated")
				listOfferingTypeLocationCrossConnectRoutersOptionsModel.LocationName = core.StringPtr("testString")
				listOfferingTypeLocationCrossConnectRoutersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := directLinkService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := directLinkService.ListOfferingTypeLocationCrossConnectRouters(listOfferingTypeLocationCrossConnectRoutersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListOfferingTypeLocationCrossConnectRoutersOptions model with no property values
				listOfferingTypeLocationCrossConnectRoutersOptionsModelNew := new(directlinkv1.ListOfferingTypeLocationCrossConnectRoutersOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = directLinkService.ListOfferingTypeLocationCrossConnectRouters(listOfferingTypeLocationCrossConnectRoutersOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ListOfferingTypeLocationCrossConnectRouters successfully`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Construct an instance of the ListOfferingTypeLocationCrossConnectRoutersOptions model
				listOfferingTypeLocationCrossConnectRoutersOptionsModel := new(directlinkv1.ListOfferingTypeLocationCrossConnectRoutersOptions)
				listOfferingTypeLocationCrossConnectRoutersOptionsModel.OfferingType = core.StringPtr("dedicated")
				listOfferingTypeLocationCrossConnectRoutersOptionsModel.LocationName = core.StringPtr("testString")
				listOfferingTypeLocationCrossConnectRoutersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := directLinkService.ListOfferingTypeLocationCrossConnectRouters(listOfferingTypeLocationCrossConnectRoutersOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListOfferingTypeSpeeds(listOfferingTypeSpeedsOptions *ListOfferingTypeSpeedsOptions) - Operation response error`, func() {
		version := "testString"
		listOfferingTypeSpeedsPath := "/offering_types/dedicated/speeds"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listOfferingTypeSpeedsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListOfferingTypeSpeeds with error: Operation response processing error`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Construct an instance of the ListOfferingTypeSpeedsOptions model
				listOfferingTypeSpeedsOptionsModel := new(directlinkv1.ListOfferingTypeSpeedsOptions)
				listOfferingTypeSpeedsOptionsModel.OfferingType = core.StringPtr("dedicated")
				listOfferingTypeSpeedsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := directLinkService.ListOfferingTypeSpeeds(listOfferingTypeSpeedsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				directLinkService.EnableRetries(0, 0)
				result, response, operationErr = directLinkService.ListOfferingTypeSpeeds(listOfferingTypeSpeedsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListOfferingTypeSpeeds(listOfferingTypeSpeedsOptions *ListOfferingTypeSpeedsOptions)`, func() {
		version := "testString"
		listOfferingTypeSpeedsPath := "/offering_types/dedicated/speeds"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listOfferingTypeSpeedsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"speeds": [{"capabilities": ["Capabilities"], "link_speed": 2000, "macsec_enabled": false}]}`)
				}))
			})
			It(`Invoke ListOfferingTypeSpeeds successfully with retries`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())
				directLinkService.EnableRetries(0, 0)

				// Construct an instance of the ListOfferingTypeSpeedsOptions model
				listOfferingTypeSpeedsOptionsModel := new(directlinkv1.ListOfferingTypeSpeedsOptions)
				listOfferingTypeSpeedsOptionsModel.OfferingType = core.StringPtr("dedicated")
				listOfferingTypeSpeedsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := directLinkService.ListOfferingTypeSpeedsWithContext(ctx, listOfferingTypeSpeedsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				directLinkService.DisableRetries()
				result, response, operationErr := directLinkService.ListOfferingTypeSpeeds(listOfferingTypeSpeedsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = directLinkService.ListOfferingTypeSpeedsWithContext(ctx, listOfferingTypeSpeedsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listOfferingTypeSpeedsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"speeds": [{"capabilities": ["Capabilities"], "link_speed": 2000, "macsec_enabled": false}]}`)
				}))
			})
			It(`Invoke ListOfferingTypeSpeeds successfully`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := directLinkService.ListOfferingTypeSpeeds(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListOfferingTypeSpeedsOptions model
				listOfferingTypeSpeedsOptionsModel := new(directlinkv1.ListOfferingTypeSpeedsOptions)
				listOfferingTypeSpeedsOptionsModel.OfferingType = core.StringPtr("dedicated")
				listOfferingTypeSpeedsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = directLinkService.ListOfferingTypeSpeeds(listOfferingTypeSpeedsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListOfferingTypeSpeeds with error: Operation validation and request error`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Construct an instance of the ListOfferingTypeSpeedsOptions model
				listOfferingTypeSpeedsOptionsModel := new(directlinkv1.ListOfferingTypeSpeedsOptions)
				listOfferingTypeSpeedsOptionsModel.OfferingType = core.StringPtr("dedicated")
				listOfferingTypeSpeedsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := directLinkService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := directLinkService.ListOfferingTypeSpeeds(listOfferingTypeSpeedsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListOfferingTypeSpeedsOptions model with no property values
				listOfferingTypeSpeedsOptionsModelNew := new(directlinkv1.ListOfferingTypeSpeedsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = directLinkService.ListOfferingTypeSpeeds(listOfferingTypeSpeedsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ListOfferingTypeSpeeds successfully`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Construct an instance of the ListOfferingTypeSpeedsOptions model
				listOfferingTypeSpeedsOptionsModel := new(directlinkv1.ListOfferingTypeSpeedsOptions)
				listOfferingTypeSpeedsOptionsModel.OfferingType = core.StringPtr("dedicated")
				listOfferingTypeSpeedsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := directLinkService.ListOfferingTypeSpeeds(listOfferingTypeSpeedsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListPorts(listPortsOptions *ListPortsOptions) - Operation response error`, func() {
		version := "testString"
		listPortsPath := "/ports"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listPortsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["location_name"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListPorts with error: Operation response processing error`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Construct an instance of the ListPortsOptions model
				listPortsOptionsModel := new(directlinkv1.ListPortsOptions)
				listPortsOptionsModel.Start = core.StringPtr("testString")
				listPortsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listPortsOptionsModel.LocationName = core.StringPtr("testString")
				listPortsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := directLinkService.ListPorts(listPortsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				directLinkService.EnableRetries(0, 0)
				result, response, operationErr = directLinkService.ListPorts(listPortsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListPorts(listPortsOptions *ListPortsOptions)`, func() {
		version := "testString"
		listPortsPath := "/ports"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listPortsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["location_name"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"first": {"href": "https://directlink.cloud.ibm.com/v1/ports?limit=100"}, "limit": 100, "next": {"href": "https://directlink.cloud.ibm.com/v1/ports?start=9d5a91a3e2cbd233b5a5b33436855ed1&limit=100", "start": "9d5a91a3e2cbd233b5a5b33436855ed1"}, "total_count": 132, "ports": [{"direct_link_count": 1, "id": "01122b9b-820f-4c44-8a31-77f1f0806765", "label": "XCR-FRK-CS-SEC-01", "location_display_name": "Dallas 03", "location_name": "dal03", "provider_name": "provider_1", "supported_link_speeds": [19]}]}`)
				}))
			})
			It(`Invoke ListPorts successfully with retries`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())
				directLinkService.EnableRetries(0, 0)

				// Construct an instance of the ListPortsOptions model
				listPortsOptionsModel := new(directlinkv1.ListPortsOptions)
				listPortsOptionsModel.Start = core.StringPtr("testString")
				listPortsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listPortsOptionsModel.LocationName = core.StringPtr("testString")
				listPortsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := directLinkService.ListPortsWithContext(ctx, listPortsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				directLinkService.DisableRetries()
				result, response, operationErr := directLinkService.ListPorts(listPortsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = directLinkService.ListPortsWithContext(ctx, listPortsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listPortsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["location_name"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"first": {"href": "https://directlink.cloud.ibm.com/v1/ports?limit=100"}, "limit": 100, "next": {"href": "https://directlink.cloud.ibm.com/v1/ports?start=9d5a91a3e2cbd233b5a5b33436855ed1&limit=100", "start": "9d5a91a3e2cbd233b5a5b33436855ed1"}, "total_count": 132, "ports": [{"direct_link_count": 1, "id": "01122b9b-820f-4c44-8a31-77f1f0806765", "label": "XCR-FRK-CS-SEC-01", "location_display_name": "Dallas 03", "location_name": "dal03", "provider_name": "provider_1", "supported_link_speeds": [19]}]}`)
				}))
			})
			It(`Invoke ListPorts successfully`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := directLinkService.ListPorts(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListPortsOptions model
				listPortsOptionsModel := new(directlinkv1.ListPortsOptions)
				listPortsOptionsModel.Start = core.StringPtr("testString")
				listPortsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listPortsOptionsModel.LocationName = core.StringPtr("testString")
				listPortsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = directLinkService.ListPorts(listPortsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListPorts with error: Operation request error`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Construct an instance of the ListPortsOptions model
				listPortsOptionsModel := new(directlinkv1.ListPortsOptions)
				listPortsOptionsModel.Start = core.StringPtr("testString")
				listPortsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listPortsOptionsModel.LocationName = core.StringPtr("testString")
				listPortsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := directLinkService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := directLinkService.ListPorts(listPortsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ListPorts successfully`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Construct an instance of the ListPortsOptions model
				listPortsOptionsModel := new(directlinkv1.ListPortsOptions)
				listPortsOptionsModel.Start = core.StringPtr("testString")
				listPortsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listPortsOptionsModel.LocationName = core.StringPtr("testString")
				listPortsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := directLinkService.ListPorts(listPortsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Test pagination helper method on response`, func() {
			It(`Invoke GetNextStart successfully`, func() {
				responseObject := new(directlinkv1.PortCollection)
				nextObject := new(directlinkv1.PortsPaginatedCollectionNext)
				nextObject.Start = core.StringPtr("abc-123")
				responseObject.Next = nextObject

				value, err := responseObject.GetNextStart()
				Expect(err).To(BeNil())
				Expect(value).To(Equal(core.StringPtr("abc-123")))
			})
			It(`Invoke GetNextStart without a "Next" property in the response`, func() {
				responseObject := new(directlinkv1.PortCollection)

				value, err := responseObject.GetNextStart()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
		})
		Context(`Using mock server endpoint - paginated response`, func() {
			BeforeEach(func() {
				var requestNumber int = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listPortsPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					requestNumber++
					if requestNumber == 1 {
						fmt.Fprintf(res, "%s", `{"next":{"start":"1"},"total_count":2,"limit":1,"ports":[{"direct_link_count":1,"id":"01122b9b-820f-4c44-8a31-77f1f0806765","label":"XCR-FRK-CS-SEC-01","location_display_name":"Dallas 03","location_name":"dal03","provider_name":"provider_1","supported_link_speeds":[19]}]}`)
					} else if requestNumber == 2 {
						fmt.Fprintf(res, "%s", `{"total_count":2,"limit":1,"ports":[{"direct_link_count":1,"id":"01122b9b-820f-4c44-8a31-77f1f0806765","label":"XCR-FRK-CS-SEC-01","location_display_name":"Dallas 03","location_name":"dal03","provider_name":"provider_1","supported_link_speeds":[19]}]}`)
					} else {
						res.WriteHeader(400)
					}
				}))
			})
			It(`Use PortsPager.GetNext successfully`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				listPortsOptionsModel := &directlinkv1.ListPortsOptions{
					Limit:        core.Int64Ptr(int64(10)),
					LocationName: core.StringPtr("testString"),
				}

				pager, err := directLinkService.NewPortsPager(listPortsOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				var allResults []directlinkv1.Port
				for pager.HasNext() {
					nextPage, err := pager.GetNext()
					Expect(err).To(BeNil())
					Expect(nextPage).ToNot(BeNil())
					allResults = append(allResults, nextPage...)
				}
				Expect(len(allResults)).To(Equal(2))
			})
			It(`Use PortsPager.GetAll successfully`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				listPortsOptionsModel := &directlinkv1.ListPortsOptions{
					Limit:        core.Int64Ptr(int64(10)),
					LocationName: core.StringPtr("testString"),
				}

				pager, err := directLinkService.NewPortsPager(listPortsOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				allResults, err := pager.GetAll()
				Expect(err).To(BeNil())
				Expect(allResults).ToNot(BeNil())
				Expect(len(allResults)).To(Equal(2))
			})
		})
	})
	Describe(`GetPort(getPortOptions *GetPortOptions) - Operation response error`, func() {
		version := "testString"
		getPortPath := "/ports/0a06fb9b-820f-4c44-8a31-77f1f0806d28"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getPortPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetPort with error: Operation response processing error`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Construct an instance of the GetPortOptions model
				getPortOptionsModel := new(directlinkv1.GetPortOptions)
				getPortOptionsModel.ID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				getPortOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := directLinkService.GetPort(getPortOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				directLinkService.EnableRetries(0, 0)
				result, response, operationErr = directLinkService.GetPort(getPortOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetPort(getPortOptions *GetPortOptions)`, func() {
		version := "testString"
		getPortPath := "/ports/0a06fb9b-820f-4c44-8a31-77f1f0806d28"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getPortPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"direct_link_count": 1, "id": "01122b9b-820f-4c44-8a31-77f1f0806765", "label": "XCR-FRK-CS-SEC-01", "location_display_name": "Dallas 03", "location_name": "dal03", "provider_name": "provider_1", "supported_link_speeds": [19]}`)
				}))
			})
			It(`Invoke GetPort successfully with retries`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())
				directLinkService.EnableRetries(0, 0)

				// Construct an instance of the GetPortOptions model
				getPortOptionsModel := new(directlinkv1.GetPortOptions)
				getPortOptionsModel.ID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				getPortOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := directLinkService.GetPortWithContext(ctx, getPortOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				directLinkService.DisableRetries()
				result, response, operationErr := directLinkService.GetPort(getPortOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = directLinkService.GetPortWithContext(ctx, getPortOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getPortPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"direct_link_count": 1, "id": "01122b9b-820f-4c44-8a31-77f1f0806765", "label": "XCR-FRK-CS-SEC-01", "location_display_name": "Dallas 03", "location_name": "dal03", "provider_name": "provider_1", "supported_link_speeds": [19]}`)
				}))
			})
			It(`Invoke GetPort successfully`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := directLinkService.GetPort(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetPortOptions model
				getPortOptionsModel := new(directlinkv1.GetPortOptions)
				getPortOptionsModel.ID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				getPortOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = directLinkService.GetPort(getPortOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetPort with error: Operation validation and request error`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Construct an instance of the GetPortOptions model
				getPortOptionsModel := new(directlinkv1.GetPortOptions)
				getPortOptionsModel.ID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				getPortOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := directLinkService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := directLinkService.GetPort(getPortOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetPortOptions model with no property values
				getPortOptionsModelNew := new(directlinkv1.GetPortOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = directLinkService.GetPort(getPortOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetPort successfully`, func() {
				directLinkService, serviceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkService).ToNot(BeNil())

				// Construct an instance of the GetPortOptions model
				getPortOptionsModel := new(directlinkv1.GetPortOptions)
				getPortOptionsModel.ID = core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				getPortOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := directLinkService.GetPort(getPortOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Model constructor tests`, func() {
		Context(`Using a service client instance`, func() {
			version := "testString"
			directLinkService, _ := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
				URL:           "http://directlinkv1modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
				Version:       core.StringPtr(version),
			})
			It(`Invoke NewAsPrependPrefixArrayTemplate successfully`, func() {
				length := int64(4)
				policy := "import"
				_model, err := directLinkService.NewAsPrependPrefixArrayTemplate(length, policy)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewAsPrependTemplate successfully`, func() {
				length := int64(4)
				policy := "import"
				_model, err := directLinkService.NewAsPrependTemplate(length, policy)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewCreateGatewayActionOptions successfully`, func() {
				// Construct an instance of the AsPrependTemplate model
				asPrependTemplateModel := new(directlinkv1.AsPrependTemplate)
				Expect(asPrependTemplateModel).ToNot(BeNil())
				asPrependTemplateModel.Length = core.Int64Ptr(int64(4))
				asPrependTemplateModel.Policy = core.StringPtr("import")
				asPrependTemplateModel.Prefix = core.StringPtr("172.17.0.0/16")
				asPrependTemplateModel.SpecificPrefixes = []string{"192.168.3.0/24"}
				Expect(asPrependTemplateModel.Length).To(Equal(core.Int64Ptr(int64(4))))
				Expect(asPrependTemplateModel.Policy).To(Equal(core.StringPtr("import")))
				Expect(asPrependTemplateModel.Prefix).To(Equal(core.StringPtr("172.17.0.0/16")))
				Expect(asPrependTemplateModel.SpecificPrefixes).To(Equal([]string{"192.168.3.0/24"}))

				// Construct an instance of the AuthenticationKeyIdentityKeyProtectAuthenticationKeyIdentity model
				authenticationKeyIdentityModel := new(directlinkv1.AuthenticationKeyIdentityKeyProtectAuthenticationKeyIdentity)
				Expect(authenticationKeyIdentityModel).ToNot(BeNil())
				authenticationKeyIdentityModel.Crn = core.StringPtr("crn:v1:bluemix:public:kms:us-south:a/4111d05f36894e3cb9b46a43556d9000:abc111b8-37aa-4034-9def-f2607c87aaaa:key:bbb222bc-430a-4de9-9aad-84e5bb022222")
				Expect(authenticationKeyIdentityModel.Crn).To(Equal(core.StringPtr("crn:v1:bluemix:public:kms:us-south:a/4111d05f36894e3cb9b46a43556d9000:abc111b8-37aa-4034-9def-f2607c87aaaa:key:bbb222bc-430a-4de9-9aad-84e5bb022222")))

				// Construct an instance of the GatewayBfdConfigActionTemplate model
				gatewayBfdConfigActionTemplateModel := new(directlinkv1.GatewayBfdConfigActionTemplate)
				Expect(gatewayBfdConfigActionTemplateModel).ToNot(BeNil())
				gatewayBfdConfigActionTemplateModel.Interval = core.Int64Ptr(int64(2000))
				gatewayBfdConfigActionTemplateModel.Multiplier = core.Int64Ptr(int64(10))
				Expect(gatewayBfdConfigActionTemplateModel.Interval).To(Equal(core.Int64Ptr(int64(2000))))
				Expect(gatewayBfdConfigActionTemplateModel.Multiplier).To(Equal(core.Int64Ptr(int64(10))))

				// Construct an instance of the GatewayTemplateRouteFilter model
				gatewayTemplateRouteFilterModel := new(directlinkv1.GatewayTemplateRouteFilter)
				Expect(gatewayTemplateRouteFilterModel).ToNot(BeNil())
				gatewayTemplateRouteFilterModel.Action = core.StringPtr("permit")
				gatewayTemplateRouteFilterModel.Ge = core.Int64Ptr(int64(25))
				gatewayTemplateRouteFilterModel.Le = core.Int64Ptr(int64(30))
				gatewayTemplateRouteFilterModel.Prefix = core.StringPtr("192.168.100.0/24")
				Expect(gatewayTemplateRouteFilterModel.Action).To(Equal(core.StringPtr("permit")))
				Expect(gatewayTemplateRouteFilterModel.Ge).To(Equal(core.Int64Ptr(int64(25))))
				Expect(gatewayTemplateRouteFilterModel.Le).To(Equal(core.Int64Ptr(int64(30))))
				Expect(gatewayTemplateRouteFilterModel.Prefix).To(Equal(core.StringPtr("192.168.100.0/24")))

				// Construct an instance of the ResourceGroupIdentity model
				resourceGroupIdentityModel := new(directlinkv1.ResourceGroupIdentity)
				Expect(resourceGroupIdentityModel).ToNot(BeNil())
				resourceGroupIdentityModel.ID = core.StringPtr("56969d6043e9465c883cb9f7363e78e8")
				Expect(resourceGroupIdentityModel.ID).To(Equal(core.StringPtr("56969d6043e9465c883cb9f7363e78e8")))

				// Construct an instance of the GatewayActionTemplateUpdatesItemGatewayClientSpeedUpdate model
				gatewayActionTemplateUpdatesItemModel := new(directlinkv1.GatewayActionTemplateUpdatesItemGatewayClientSpeedUpdate)
				Expect(gatewayActionTemplateUpdatesItemModel).ToNot(BeNil())
				gatewayActionTemplateUpdatesItemModel.SpeedMbps = core.Int64Ptr(int64(500))
				Expect(gatewayActionTemplateUpdatesItemModel.SpeedMbps).To(Equal(core.Int64Ptr(int64(500))))

				// Construct an instance of the CreateGatewayActionOptions model
				id := "0a06fb9b-820f-4c44-8a31-77f1f0806d28"
				createGatewayActionOptionsModel := directLinkService.NewCreateGatewayActionOptions(id)
				createGatewayActionOptionsModel.SetID("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				createGatewayActionOptionsModel.SetAction("create_gateway_approve")
				createGatewayActionOptionsModel.SetAsPrepends([]directlinkv1.AsPrependTemplate{*asPrependTemplateModel})
				createGatewayActionOptionsModel.SetAuthenticationKey(authenticationKeyIdentityModel)
				createGatewayActionOptionsModel.SetBfdConfig(gatewayBfdConfigActionTemplateModel)
				createGatewayActionOptionsModel.SetConnectionMode("transit")
				createGatewayActionOptionsModel.SetDefaultExportRouteFilter("permit")
				createGatewayActionOptionsModel.SetDefaultImportRouteFilter("permit")
				createGatewayActionOptionsModel.SetExportRouteFilters([]directlinkv1.GatewayTemplateRouteFilter{*gatewayTemplateRouteFilterModel})
				createGatewayActionOptionsModel.SetGlobal(true)
				createGatewayActionOptionsModel.SetImportRouteFilters([]directlinkv1.GatewayTemplateRouteFilter{*gatewayTemplateRouteFilterModel})
				createGatewayActionOptionsModel.SetMetered(false)
				createGatewayActionOptionsModel.SetResourceGroup(resourceGroupIdentityModel)
				createGatewayActionOptionsModel.SetUpdates([]directlinkv1.GatewayActionTemplateUpdatesItemIntf{gatewayActionTemplateUpdatesItemModel})
				createGatewayActionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createGatewayActionOptionsModel).ToNot(BeNil())
				Expect(createGatewayActionOptionsModel.ID).To(Equal(core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")))
				Expect(createGatewayActionOptionsModel.Action).To(Equal(core.StringPtr("create_gateway_approve")))
				Expect(createGatewayActionOptionsModel.AsPrepends).To(Equal([]directlinkv1.AsPrependTemplate{*asPrependTemplateModel}))
				Expect(createGatewayActionOptionsModel.AuthenticationKey).To(Equal(authenticationKeyIdentityModel))
				Expect(createGatewayActionOptionsModel.BfdConfig).To(Equal(gatewayBfdConfigActionTemplateModel))
				Expect(createGatewayActionOptionsModel.ConnectionMode).To(Equal(core.StringPtr("transit")))
				Expect(createGatewayActionOptionsModel.DefaultExportRouteFilter).To(Equal(core.StringPtr("permit")))
				Expect(createGatewayActionOptionsModel.DefaultImportRouteFilter).To(Equal(core.StringPtr("permit")))
				Expect(createGatewayActionOptionsModel.ExportRouteFilters).To(Equal([]directlinkv1.GatewayTemplateRouteFilter{*gatewayTemplateRouteFilterModel}))
				Expect(createGatewayActionOptionsModel.Global).To(Equal(core.BoolPtr(true)))
				Expect(createGatewayActionOptionsModel.ImportRouteFilters).To(Equal([]directlinkv1.GatewayTemplateRouteFilter{*gatewayTemplateRouteFilterModel}))
				Expect(createGatewayActionOptionsModel.Metered).To(Equal(core.BoolPtr(false)))
				Expect(createGatewayActionOptionsModel.ResourceGroup).To(Equal(resourceGroupIdentityModel))
				Expect(createGatewayActionOptionsModel.Updates).To(Equal([]directlinkv1.GatewayActionTemplateUpdatesItemIntf{gatewayActionTemplateUpdatesItemModel}))
				Expect(createGatewayActionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateGatewayCompletionNoticeOptions successfully`, func() {
				// Construct an instance of the CreateGatewayCompletionNoticeOptions model
				id := "0a06fb9b-820f-4c44-8a31-77f1f0806d28"
				createGatewayCompletionNoticeOptionsModel := directLinkService.NewCreateGatewayCompletionNoticeOptions(id)
				createGatewayCompletionNoticeOptionsModel.SetID("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				createGatewayCompletionNoticeOptionsModel.SetUpload(CreateMockReader("This is a mock file."))
				createGatewayCompletionNoticeOptionsModel.SetUploadContentType("testString")
				createGatewayCompletionNoticeOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createGatewayCompletionNoticeOptionsModel).ToNot(BeNil())
				Expect(createGatewayCompletionNoticeOptionsModel.ID).To(Equal(core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")))
				Expect(createGatewayCompletionNoticeOptionsModel.Upload).To(Equal(CreateMockReader("This is a mock file.")))
				Expect(createGatewayCompletionNoticeOptionsModel.UploadContentType).To(Equal(core.StringPtr("testString")))
				Expect(createGatewayCompletionNoticeOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateGatewayExportRouteFilterOptions successfully`, func() {
				// Construct an instance of the CreateGatewayExportRouteFilterOptions model
				gatewayID := "0a06fb9b-820f-4c44-8a31-77f1f0806d28"
				createGatewayExportRouteFilterOptionsAction := "permit"
				createGatewayExportRouteFilterOptionsPrefix := "192.168.100.0/24"
				createGatewayExportRouteFilterOptionsModel := directLinkService.NewCreateGatewayExportRouteFilterOptions(gatewayID, createGatewayExportRouteFilterOptionsAction, createGatewayExportRouteFilterOptionsPrefix)
				createGatewayExportRouteFilterOptionsModel.SetGatewayID("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				createGatewayExportRouteFilterOptionsModel.SetAction("permit")
				createGatewayExportRouteFilterOptionsModel.SetPrefix("192.168.100.0/24")
				createGatewayExportRouteFilterOptionsModel.SetBefore("1a15dcab-7e40-45e1-b7c5-bc690eaa9782")
				createGatewayExportRouteFilterOptionsModel.SetGe(int64(25))
				createGatewayExportRouteFilterOptionsModel.SetLe(int64(30))
				createGatewayExportRouteFilterOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createGatewayExportRouteFilterOptionsModel).ToNot(BeNil())
				Expect(createGatewayExportRouteFilterOptionsModel.GatewayID).To(Equal(core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")))
				Expect(createGatewayExportRouteFilterOptionsModel.Action).To(Equal(core.StringPtr("permit")))
				Expect(createGatewayExportRouteFilterOptionsModel.Prefix).To(Equal(core.StringPtr("192.168.100.0/24")))
				Expect(createGatewayExportRouteFilterOptionsModel.Before).To(Equal(core.StringPtr("1a15dcab-7e40-45e1-b7c5-bc690eaa9782")))
				Expect(createGatewayExportRouteFilterOptionsModel.Ge).To(Equal(core.Int64Ptr(int64(25))))
				Expect(createGatewayExportRouteFilterOptionsModel.Le).To(Equal(core.Int64Ptr(int64(30))))
				Expect(createGatewayExportRouteFilterOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateGatewayImportRouteFilterOptions successfully`, func() {
				// Construct an instance of the CreateGatewayImportRouteFilterOptions model
				gatewayID := "0a06fb9b-820f-4c44-8a31-77f1f0806d28"
				createGatewayImportRouteFilterOptionsAction := "permit"
				createGatewayImportRouteFilterOptionsPrefix := "192.168.100.0/24"
				createGatewayImportRouteFilterOptionsModel := directLinkService.NewCreateGatewayImportRouteFilterOptions(gatewayID, createGatewayImportRouteFilterOptionsAction, createGatewayImportRouteFilterOptionsPrefix)
				createGatewayImportRouteFilterOptionsModel.SetGatewayID("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				createGatewayImportRouteFilterOptionsModel.SetAction("permit")
				createGatewayImportRouteFilterOptionsModel.SetPrefix("192.168.100.0/24")
				createGatewayImportRouteFilterOptionsModel.SetBefore("1a15dcab-7e40-45e1-b7c5-bc690eaa9782")
				createGatewayImportRouteFilterOptionsModel.SetGe(int64(25))
				createGatewayImportRouteFilterOptionsModel.SetLe(int64(30))
				createGatewayImportRouteFilterOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createGatewayImportRouteFilterOptionsModel).ToNot(BeNil())
				Expect(createGatewayImportRouteFilterOptionsModel.GatewayID).To(Equal(core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")))
				Expect(createGatewayImportRouteFilterOptionsModel.Action).To(Equal(core.StringPtr("permit")))
				Expect(createGatewayImportRouteFilterOptionsModel.Prefix).To(Equal(core.StringPtr("192.168.100.0/24")))
				Expect(createGatewayImportRouteFilterOptionsModel.Before).To(Equal(core.StringPtr("1a15dcab-7e40-45e1-b7c5-bc690eaa9782")))
				Expect(createGatewayImportRouteFilterOptionsModel.Ge).To(Equal(core.Int64Ptr(int64(25))))
				Expect(createGatewayImportRouteFilterOptionsModel.Le).To(Equal(core.Int64Ptr(int64(30))))
				Expect(createGatewayImportRouteFilterOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateGatewayMacsecCakOptions successfully`, func() {
				// Construct an instance of the HpcsKeyIdentity model
				hpcsKeyIdentityModel := new(directlinkv1.HpcsKeyIdentity)
				Expect(hpcsKeyIdentityModel).ToNot(BeNil())
				hpcsKeyIdentityModel.Crn = core.StringPtr("crn:v1:bluemix:public:hs-crypto:us-south:a/4111d05f36894e3cb9b46a43556d9000:abc111b8-37aa-4034-9def-f2607c87aaaa:key:bbb222bc-430a-4de9-9aad-84e5bb022222")
				Expect(hpcsKeyIdentityModel.Crn).To(Equal(core.StringPtr("crn:v1:bluemix:public:hs-crypto:us-south:a/4111d05f36894e3cb9b46a43556d9000:abc111b8-37aa-4034-9def-f2607c87aaaa:key:bbb222bc-430a-4de9-9aad-84e5bb022222")))

				// Construct an instance of the CreateGatewayMacsecCakOptions model
				id := "0a06fb9b-820f-4c44-8a31-77f1f0806d28"
				var createGatewayMacsecCakOptionsKey *directlinkv1.HpcsKeyIdentity = nil
				createGatewayMacsecCakOptionsName := "1000"
				createGatewayMacsecCakOptionsSession := "primary"
				createGatewayMacsecCakOptionsModel := directLinkService.NewCreateGatewayMacsecCakOptions(id, createGatewayMacsecCakOptionsKey, createGatewayMacsecCakOptionsName, createGatewayMacsecCakOptionsSession)
				createGatewayMacsecCakOptionsModel.SetID("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				createGatewayMacsecCakOptionsModel.SetKey(hpcsKeyIdentityModel)
				createGatewayMacsecCakOptionsModel.SetName("1000")
				createGatewayMacsecCakOptionsModel.SetSession("primary")
				createGatewayMacsecCakOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createGatewayMacsecCakOptionsModel).ToNot(BeNil())
				Expect(createGatewayMacsecCakOptionsModel.ID).To(Equal(core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")))
				Expect(createGatewayMacsecCakOptionsModel.Key).To(Equal(hpcsKeyIdentityModel))
				Expect(createGatewayMacsecCakOptionsModel.Name).To(Equal(core.StringPtr("1000")))
				Expect(createGatewayMacsecCakOptionsModel.Session).To(Equal(core.StringPtr("primary")))
				Expect(createGatewayMacsecCakOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateGatewayOptions successfully`, func() {
				// Construct an instance of the AsPrependTemplate model
				asPrependTemplateModel := new(directlinkv1.AsPrependTemplate)
				Expect(asPrependTemplateModel).ToNot(BeNil())
				asPrependTemplateModel.Length = core.Int64Ptr(int64(4))
				asPrependTemplateModel.Policy = core.StringPtr("import")
				asPrependTemplateModel.Prefix = core.StringPtr("172.17.0.0/16")
				asPrependTemplateModel.SpecificPrefixes = []string{"192.168.3.0/24"}
				Expect(asPrependTemplateModel.Length).To(Equal(core.Int64Ptr(int64(4))))
				Expect(asPrependTemplateModel.Policy).To(Equal(core.StringPtr("import")))
				Expect(asPrependTemplateModel.Prefix).To(Equal(core.StringPtr("172.17.0.0/16")))
				Expect(asPrependTemplateModel.SpecificPrefixes).To(Equal([]string{"192.168.3.0/24"}))

				// Construct an instance of the AuthenticationKeyIdentityKeyProtectAuthenticationKeyIdentity model
				authenticationKeyIdentityModel := new(directlinkv1.AuthenticationKeyIdentityKeyProtectAuthenticationKeyIdentity)
				Expect(authenticationKeyIdentityModel).ToNot(BeNil())
				authenticationKeyIdentityModel.Crn = core.StringPtr("crn:v1:bluemix:public:kms:us-south:a/4111d05f36894e3cb9b46a43556d9000:abc111b8-37aa-4034-9def-f2607c87aaaa:key:bbb222bc-430a-4de9-9aad-84e5bb022222")
				Expect(authenticationKeyIdentityModel.Crn).To(Equal(core.StringPtr("crn:v1:bluemix:public:kms:us-south:a/4111d05f36894e3cb9b46a43556d9000:abc111b8-37aa-4034-9def-f2607c87aaaa:key:bbb222bc-430a-4de9-9aad-84e5bb022222")))

				// Construct an instance of the GatewayBfdConfigTemplate model
				gatewayBfdConfigTemplateModel := new(directlinkv1.GatewayBfdConfigTemplate)
				Expect(gatewayBfdConfigTemplateModel).ToNot(BeNil())
				gatewayBfdConfigTemplateModel.Interval = core.Int64Ptr(int64(2000))
				gatewayBfdConfigTemplateModel.Multiplier = core.Int64Ptr(int64(10))
				Expect(gatewayBfdConfigTemplateModel.Interval).To(Equal(core.Int64Ptr(int64(2000))))
				Expect(gatewayBfdConfigTemplateModel.Multiplier).To(Equal(core.Int64Ptr(int64(10))))

				// Construct an instance of the GatewayTemplateRouteFilter model
				gatewayTemplateRouteFilterModel := new(directlinkv1.GatewayTemplateRouteFilter)
				Expect(gatewayTemplateRouteFilterModel).ToNot(BeNil())
				gatewayTemplateRouteFilterModel.Action = core.StringPtr("permit")
				gatewayTemplateRouteFilterModel.Ge = core.Int64Ptr(int64(25))
				gatewayTemplateRouteFilterModel.Le = core.Int64Ptr(int64(30))
				gatewayTemplateRouteFilterModel.Prefix = core.StringPtr("192.168.100.0/24")
				Expect(gatewayTemplateRouteFilterModel.Action).To(Equal(core.StringPtr("permit")))
				Expect(gatewayTemplateRouteFilterModel.Ge).To(Equal(core.Int64Ptr(int64(25))))
				Expect(gatewayTemplateRouteFilterModel.Le).To(Equal(core.Int64Ptr(int64(30))))
				Expect(gatewayTemplateRouteFilterModel.Prefix).To(Equal(core.StringPtr("192.168.100.0/24")))

				// Construct an instance of the ResourceGroupIdentity model
				resourceGroupIdentityModel := new(directlinkv1.ResourceGroupIdentity)
				Expect(resourceGroupIdentityModel).ToNot(BeNil())
				resourceGroupIdentityModel.ID = core.StringPtr("56969d6043e9465c883cb9f7363e78e8")
				Expect(resourceGroupIdentityModel.ID).To(Equal(core.StringPtr("56969d6043e9465c883cb9f7363e78e8")))

				// Construct an instance of the HpcsKeyIdentity model
				hpcsKeyIdentityModel := new(directlinkv1.HpcsKeyIdentity)
				Expect(hpcsKeyIdentityModel).ToNot(BeNil())
				hpcsKeyIdentityModel.Crn = core.StringPtr("crn:v1:bluemix:public:hs-crypto:us-south:a/4111d05f36894e3cb9b46a43556d9000:abc111b8-37aa-4034-9def-f2607c87aaaa:key:bbb222bc-430a-4de9-9aad-84e5bb022222")
				Expect(hpcsKeyIdentityModel.Crn).To(Equal(core.StringPtr("crn:v1:bluemix:public:hs-crypto:us-south:a/4111d05f36894e3cb9b46a43556d9000:abc111b8-37aa-4034-9def-f2607c87aaaa:key:bbb222bc-430a-4de9-9aad-84e5bb022222")))

				// Construct an instance of the GatewayMacsecCakPrototype model
				gatewayMacsecCakPrototypeModel := new(directlinkv1.GatewayMacsecCakPrototype)
				Expect(gatewayMacsecCakPrototypeModel).ToNot(BeNil())
				gatewayMacsecCakPrototypeModel.Key = hpcsKeyIdentityModel
				gatewayMacsecCakPrototypeModel.Name = core.StringPtr("1000")
				gatewayMacsecCakPrototypeModel.Session = core.StringPtr("primary")
				Expect(gatewayMacsecCakPrototypeModel.Key).To(Equal(hpcsKeyIdentityModel))
				Expect(gatewayMacsecCakPrototypeModel.Name).To(Equal(core.StringPtr("1000")))
				Expect(gatewayMacsecCakPrototypeModel.Session).To(Equal(core.StringPtr("primary")))

				// Construct an instance of the SakRekeyPrototypeSakRekeyTimerModePrototype model
				sakRekeyPrototypeModel := new(directlinkv1.SakRekeyPrototypeSakRekeyTimerModePrototype)
				Expect(sakRekeyPrototypeModel).ToNot(BeNil())
				sakRekeyPrototypeModel.Interval = core.Int64Ptr(int64(3600))
				sakRekeyPrototypeModel.Mode = core.StringPtr("timer")
				Expect(sakRekeyPrototypeModel.Interval).To(Equal(core.Int64Ptr(int64(3600))))
				Expect(sakRekeyPrototypeModel.Mode).To(Equal(core.StringPtr("timer")))

				// Construct an instance of the GatewayMacsecPrototype model
				gatewayMacsecPrototypeModel := new(directlinkv1.GatewayMacsecPrototype)
				Expect(gatewayMacsecPrototypeModel).ToNot(BeNil())
				gatewayMacsecPrototypeModel.Active = core.BoolPtr(true)
				gatewayMacsecPrototypeModel.Caks = []directlinkv1.GatewayMacsecCakPrototype{*gatewayMacsecCakPrototypeModel}
				gatewayMacsecPrototypeModel.SakRekey = sakRekeyPrototypeModel
				gatewayMacsecPrototypeModel.SecurityPolicy = core.StringPtr("must_secure")
				gatewayMacsecPrototypeModel.WindowSize = core.Int64Ptr(int64(64))
				Expect(gatewayMacsecPrototypeModel.Active).To(Equal(core.BoolPtr(true)))
				Expect(gatewayMacsecPrototypeModel.Caks).To(Equal([]directlinkv1.GatewayMacsecCakPrototype{*gatewayMacsecCakPrototypeModel}))
				Expect(gatewayMacsecPrototypeModel.SakRekey).To(Equal(sakRekeyPrototypeModel))
				Expect(gatewayMacsecPrototypeModel.SecurityPolicy).To(Equal(core.StringPtr("must_secure")))
				Expect(gatewayMacsecPrototypeModel.WindowSize).To(Equal(core.Int64Ptr(int64(64))))

				// Construct an instance of the GatewayTemplateGatewayTypeDedicatedTemplate model
				gatewayTemplateModel := new(directlinkv1.GatewayTemplateGatewayTypeDedicatedTemplate)
				Expect(gatewayTemplateModel).ToNot(BeNil())
				gatewayTemplateModel.AsPrepends = []directlinkv1.AsPrependTemplate{*asPrependTemplateModel}
				gatewayTemplateModel.AuthenticationKey = authenticationKeyIdentityModel
				gatewayTemplateModel.BfdConfig = gatewayBfdConfigTemplateModel
				gatewayTemplateModel.BgpAsn = core.Int64Ptr(int64(64999))
				gatewayTemplateModel.BgpBaseCidr = core.StringPtr("testString")
				gatewayTemplateModel.BgpCerCidr = core.StringPtr("169.254.0.10/30")
				gatewayTemplateModel.BgpIbmCidr = core.StringPtr("169.254.0.9/30")
				gatewayTemplateModel.ConnectionMode = core.StringPtr("transit")
				gatewayTemplateModel.DefaultExportRouteFilter = core.StringPtr("permit")
				gatewayTemplateModel.DefaultImportRouteFilter = core.StringPtr("permit")
				gatewayTemplateModel.ExportRouteFilters = []directlinkv1.GatewayTemplateRouteFilter{*gatewayTemplateRouteFilterModel}
				gatewayTemplateModel.Global = core.BoolPtr(true)
				gatewayTemplateModel.ImportRouteFilters = []directlinkv1.GatewayTemplateRouteFilter{*gatewayTemplateRouteFilterModel}
				gatewayTemplateModel.Metered = core.BoolPtr(false)
				gatewayTemplateModel.Name = core.StringPtr("myGateway")
				gatewayTemplateModel.PatchPanelCompletionNotice = core.StringPtr("patch panel configuration details")
				gatewayTemplateModel.ResourceGroup = resourceGroupIdentityModel
				gatewayTemplateModel.SpeedMbps = core.Int64Ptr(int64(1000))
				gatewayTemplateModel.Type = core.StringPtr("dedicated")
				gatewayTemplateModel.CarrierName = core.StringPtr("myCarrierName")
				gatewayTemplateModel.CrossConnectRouter = core.StringPtr("xcr01.dal03")
				gatewayTemplateModel.CustomerName = core.StringPtr("newCustomerName")
				gatewayTemplateModel.LocationName = core.StringPtr("dal03")
				gatewayTemplateModel.Macsec = gatewayMacsecPrototypeModel
				gatewayTemplateModel.MacsecCapability = core.StringPtr("non_macsec")
				gatewayTemplateModel.Vlan = core.Int64Ptr(int64(10))
				Expect(gatewayTemplateModel.AsPrepends).To(Equal([]directlinkv1.AsPrependTemplate{*asPrependTemplateModel}))
				Expect(gatewayTemplateModel.AuthenticationKey).To(Equal(authenticationKeyIdentityModel))
				Expect(gatewayTemplateModel.BfdConfig).To(Equal(gatewayBfdConfigTemplateModel))
				Expect(gatewayTemplateModel.BgpAsn).To(Equal(core.Int64Ptr(int64(64999))))
				Expect(gatewayTemplateModel.BgpBaseCidr).To(Equal(core.StringPtr("testString")))
				Expect(gatewayTemplateModel.BgpCerCidr).To(Equal(core.StringPtr("169.254.0.10/30")))
				Expect(gatewayTemplateModel.BgpIbmCidr).To(Equal(core.StringPtr("169.254.0.9/30")))
				Expect(gatewayTemplateModel.ConnectionMode).To(Equal(core.StringPtr("transit")))
				Expect(gatewayTemplateModel.DefaultExportRouteFilter).To(Equal(core.StringPtr("permit")))
				Expect(gatewayTemplateModel.DefaultImportRouteFilter).To(Equal(core.StringPtr("permit")))
				Expect(gatewayTemplateModel.ExportRouteFilters).To(Equal([]directlinkv1.GatewayTemplateRouteFilter{*gatewayTemplateRouteFilterModel}))
				Expect(gatewayTemplateModel.Global).To(Equal(core.BoolPtr(true)))
				Expect(gatewayTemplateModel.ImportRouteFilters).To(Equal([]directlinkv1.GatewayTemplateRouteFilter{*gatewayTemplateRouteFilterModel}))
				Expect(gatewayTemplateModel.Metered).To(Equal(core.BoolPtr(false)))
				Expect(gatewayTemplateModel.Name).To(Equal(core.StringPtr("myGateway")))
				Expect(gatewayTemplateModel.PatchPanelCompletionNotice).To(Equal(core.StringPtr("patch panel configuration details")))
				Expect(gatewayTemplateModel.ResourceGroup).To(Equal(resourceGroupIdentityModel))
				Expect(gatewayTemplateModel.SpeedMbps).To(Equal(core.Int64Ptr(int64(1000))))
				Expect(gatewayTemplateModel.Type).To(Equal(core.StringPtr("dedicated")))
				Expect(gatewayTemplateModel.CarrierName).To(Equal(core.StringPtr("myCarrierName")))
				Expect(gatewayTemplateModel.CrossConnectRouter).To(Equal(core.StringPtr("xcr01.dal03")))
				Expect(gatewayTemplateModel.CustomerName).To(Equal(core.StringPtr("newCustomerName")))
				Expect(gatewayTemplateModel.LocationName).To(Equal(core.StringPtr("dal03")))
				Expect(gatewayTemplateModel.Macsec).To(Equal(gatewayMacsecPrototypeModel))
				Expect(gatewayTemplateModel.MacsecCapability).To(Equal(core.StringPtr("non_macsec")))
				Expect(gatewayTemplateModel.Vlan).To(Equal(core.Int64Ptr(int64(10))))

				// Construct an instance of the CreateGatewayOptions model
				var gatewayTemplate directlinkv1.GatewayTemplateIntf = nil
				createGatewayOptionsModel := directLinkService.NewCreateGatewayOptions(gatewayTemplate)
				createGatewayOptionsModel.SetGatewayTemplate(gatewayTemplateModel)
				createGatewayOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createGatewayOptionsModel).ToNot(BeNil())
				Expect(createGatewayOptionsModel.GatewayTemplate).To(Equal(gatewayTemplateModel))
				Expect(createGatewayOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateGatewayRouteReportOptions successfully`, func() {
				// Construct an instance of the CreateGatewayRouteReportOptions model
				gatewayID := "0a06fb9b-820f-4c44-8a31-77f1f0806d28"
				createGatewayRouteReportOptionsModel := directLinkService.NewCreateGatewayRouteReportOptions(gatewayID)
				createGatewayRouteReportOptionsModel.SetGatewayID("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				createGatewayRouteReportOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createGatewayRouteReportOptionsModel).ToNot(BeNil())
				Expect(createGatewayRouteReportOptionsModel.GatewayID).To(Equal(core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")))
				Expect(createGatewayRouteReportOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateGatewayVirtualConnectionOptions successfully`, func() {
				// Construct an instance of the CreateGatewayVirtualConnectionOptions model
				gatewayID := "0a06fb9b-820f-4c44-8a31-77f1f0806d28"
				createGatewayVirtualConnectionOptionsName := "newVC"
				createGatewayVirtualConnectionOptionsType := "vpc"
				createGatewayVirtualConnectionOptionsModel := directLinkService.NewCreateGatewayVirtualConnectionOptions(gatewayID, createGatewayVirtualConnectionOptionsName, createGatewayVirtualConnectionOptionsType)
				createGatewayVirtualConnectionOptionsModel.SetGatewayID("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				createGatewayVirtualConnectionOptionsModel.SetName("newVC")
				createGatewayVirtualConnectionOptionsModel.SetType("vpc")
				createGatewayVirtualConnectionOptionsModel.SetNetworkID("crn:v1:bluemix:public:is:us-east:a/28e4d90ac7504be69447111122223333::vpc:aaa81ac8-5e96-42a0-a4b7-6c2e2d1bbbbb")
				createGatewayVirtualConnectionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createGatewayVirtualConnectionOptionsModel).ToNot(BeNil())
				Expect(createGatewayVirtualConnectionOptionsModel.GatewayID).To(Equal(core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")))
				Expect(createGatewayVirtualConnectionOptionsModel.Name).To(Equal(core.StringPtr("newVC")))
				Expect(createGatewayVirtualConnectionOptionsModel.Type).To(Equal(core.StringPtr("vpc")))
				Expect(createGatewayVirtualConnectionOptionsModel.NetworkID).To(Equal(core.StringPtr("crn:v1:bluemix:public:is:us-east:a/28e4d90ac7504be69447111122223333::vpc:aaa81ac8-5e96-42a0-a4b7-6c2e2d1bbbbb")))
				Expect(createGatewayVirtualConnectionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteGatewayExportRouteFilterOptions successfully`, func() {
				// Construct an instance of the DeleteGatewayExportRouteFilterOptions model
				gatewayID := "0a06fb9b-820f-4c44-8a31-77f1f0806d28"
				id := "0a06fb9b-820f-4c44-8a31-77f1f0806d28"
				deleteGatewayExportRouteFilterOptionsModel := directLinkService.NewDeleteGatewayExportRouteFilterOptions(gatewayID, id)
				deleteGatewayExportRouteFilterOptionsModel.SetGatewayID("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				deleteGatewayExportRouteFilterOptionsModel.SetID("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				deleteGatewayExportRouteFilterOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteGatewayExportRouteFilterOptionsModel).ToNot(BeNil())
				Expect(deleteGatewayExportRouteFilterOptionsModel.GatewayID).To(Equal(core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")))
				Expect(deleteGatewayExportRouteFilterOptionsModel.ID).To(Equal(core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")))
				Expect(deleteGatewayExportRouteFilterOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteGatewayImportRouteFilterOptions successfully`, func() {
				// Construct an instance of the DeleteGatewayImportRouteFilterOptions model
				gatewayID := "0a06fb9b-820f-4c44-8a31-77f1f0806d28"
				id := "0a06fb9b-820f-4c44-8a31-77f1f0806d28"
				deleteGatewayImportRouteFilterOptionsModel := directLinkService.NewDeleteGatewayImportRouteFilterOptions(gatewayID, id)
				deleteGatewayImportRouteFilterOptionsModel.SetGatewayID("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				deleteGatewayImportRouteFilterOptionsModel.SetID("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				deleteGatewayImportRouteFilterOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteGatewayImportRouteFilterOptionsModel).ToNot(BeNil())
				Expect(deleteGatewayImportRouteFilterOptionsModel.GatewayID).To(Equal(core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")))
				Expect(deleteGatewayImportRouteFilterOptionsModel.ID).To(Equal(core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")))
				Expect(deleteGatewayImportRouteFilterOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteGatewayMacsecCakOptions successfully`, func() {
				// Construct an instance of the DeleteGatewayMacsecCakOptions model
				id := "0a06fb9b-820f-4c44-8a31-77f1f0806d28"
				cakID := "ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4"
				deleteGatewayMacsecCakOptionsModel := directLinkService.NewDeleteGatewayMacsecCakOptions(id, cakID)
				deleteGatewayMacsecCakOptionsModel.SetID("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				deleteGatewayMacsecCakOptionsModel.SetCakID("ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4")
				deleteGatewayMacsecCakOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteGatewayMacsecCakOptionsModel).ToNot(BeNil())
				Expect(deleteGatewayMacsecCakOptionsModel.ID).To(Equal(core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")))
				Expect(deleteGatewayMacsecCakOptionsModel.CakID).To(Equal(core.StringPtr("ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4")))
				Expect(deleteGatewayMacsecCakOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteGatewayOptions successfully`, func() {
				// Construct an instance of the DeleteGatewayOptions model
				id := "0a06fb9b-820f-4c44-8a31-77f1f0806d28"
				deleteGatewayOptionsModel := directLinkService.NewDeleteGatewayOptions(id)
				deleteGatewayOptionsModel.SetID("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				deleteGatewayOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteGatewayOptionsModel).ToNot(BeNil())
				Expect(deleteGatewayOptionsModel.ID).To(Equal(core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")))
				Expect(deleteGatewayOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteGatewayRouteReportOptions successfully`, func() {
				// Construct an instance of the DeleteGatewayRouteReportOptions model
				gatewayID := "0a06fb9b-820f-4c44-8a31-77f1f0806d28"
				id := "0a06fb9b-820f-4c44-8a31-77f1f0806d28"
				deleteGatewayRouteReportOptionsModel := directLinkService.NewDeleteGatewayRouteReportOptions(gatewayID, id)
				deleteGatewayRouteReportOptionsModel.SetGatewayID("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				deleteGatewayRouteReportOptionsModel.SetID("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				deleteGatewayRouteReportOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteGatewayRouteReportOptionsModel).ToNot(BeNil())
				Expect(deleteGatewayRouteReportOptionsModel.GatewayID).To(Equal(core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")))
				Expect(deleteGatewayRouteReportOptionsModel.ID).To(Equal(core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")))
				Expect(deleteGatewayRouteReportOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteGatewayVirtualConnectionOptions successfully`, func() {
				// Construct an instance of the DeleteGatewayVirtualConnectionOptions model
				gatewayID := "0a06fb9b-820f-4c44-8a31-77f1f0806d28"
				id := "0a06fb9b-820f-4c44-8a31-77f1f0806d28"
				deleteGatewayVirtualConnectionOptionsModel := directLinkService.NewDeleteGatewayVirtualConnectionOptions(gatewayID, id)
				deleteGatewayVirtualConnectionOptionsModel.SetGatewayID("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				deleteGatewayVirtualConnectionOptionsModel.SetID("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				deleteGatewayVirtualConnectionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteGatewayVirtualConnectionOptionsModel).ToNot(BeNil())
				Expect(deleteGatewayVirtualConnectionOptionsModel.GatewayID).To(Equal(core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")))
				Expect(deleteGatewayVirtualConnectionOptionsModel.ID).To(Equal(core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")))
				Expect(deleteGatewayVirtualConnectionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGatewayBfdConfigActionTemplate successfully`, func() {
				interval := int64(2000)
				_model, err := directLinkService.NewGatewayBfdConfigActionTemplate(interval)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewGatewayBfdConfigTemplate successfully`, func() {
				interval := int64(2000)
				_model, err := directLinkService.NewGatewayBfdConfigTemplate(interval)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewGatewayMacsecCakPrototype successfully`, func() {
				var key *directlinkv1.HpcsKeyIdentity = nil
				name := "1000"
				session := "primary"
				_, err := directLinkService.NewGatewayMacsecCakPrototype(key, name, session)
				Expect(err).ToNot(BeNil())
			})
			It(`Invoke NewGatewayMacsecPrototype successfully`, func() {
				active := true
				caks := []directlinkv1.GatewayMacsecCakPrototype{}
				var sakRekey directlinkv1.SakRekeyPrototypeIntf = nil
				securityPolicy := "must_secure"
				_, err := directLinkService.NewGatewayMacsecPrototype(active, caks, sakRekey, securityPolicy)
				Expect(err).ToNot(BeNil())
			})
			It(`Invoke NewGatewayPortIdentity successfully`, func() {
				id := "fffdcb1a-fee4-41c7-9e11-9cd99e65c777"
				_model, err := directLinkService.NewGatewayPortIdentity(id)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewGatewayTemplateRouteFilter successfully`, func() {
				action := "permit"
				prefix := "192.168.100.0/24"
				_model, err := directLinkService.NewGatewayTemplateRouteFilter(action, prefix)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewGetGatewayExportRouteFilterOptions successfully`, func() {
				// Construct an instance of the GetGatewayExportRouteFilterOptions model
				gatewayID := "0a06fb9b-820f-4c44-8a31-77f1f0806d28"
				id := "0a06fb9b-820f-4c44-8a31-77f1f0806d28"
				getGatewayExportRouteFilterOptionsModel := directLinkService.NewGetGatewayExportRouteFilterOptions(gatewayID, id)
				getGatewayExportRouteFilterOptionsModel.SetGatewayID("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				getGatewayExportRouteFilterOptionsModel.SetID("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				getGatewayExportRouteFilterOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getGatewayExportRouteFilterOptionsModel).ToNot(BeNil())
				Expect(getGatewayExportRouteFilterOptionsModel.GatewayID).To(Equal(core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")))
				Expect(getGatewayExportRouteFilterOptionsModel.ID).To(Equal(core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")))
				Expect(getGatewayExportRouteFilterOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetGatewayImportRouteFilterOptions successfully`, func() {
				// Construct an instance of the GetGatewayImportRouteFilterOptions model
				gatewayID := "0a06fb9b-820f-4c44-8a31-77f1f0806d28"
				id := "0a06fb9b-820f-4c44-8a31-77f1f0806d28"
				getGatewayImportRouteFilterOptionsModel := directLinkService.NewGetGatewayImportRouteFilterOptions(gatewayID, id)
				getGatewayImportRouteFilterOptionsModel.SetGatewayID("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				getGatewayImportRouteFilterOptionsModel.SetID("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				getGatewayImportRouteFilterOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getGatewayImportRouteFilterOptionsModel).ToNot(BeNil())
				Expect(getGatewayImportRouteFilterOptionsModel.GatewayID).To(Equal(core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")))
				Expect(getGatewayImportRouteFilterOptionsModel.ID).To(Equal(core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")))
				Expect(getGatewayImportRouteFilterOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetGatewayMacsecCakOptions successfully`, func() {
				// Construct an instance of the GetGatewayMacsecCakOptions model
				id := "0a06fb9b-820f-4c44-8a31-77f1f0806d28"
				cakID := "ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4"
				getGatewayMacsecCakOptionsModel := directLinkService.NewGetGatewayMacsecCakOptions(id, cakID)
				getGatewayMacsecCakOptionsModel.SetID("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				getGatewayMacsecCakOptionsModel.SetCakID("ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4")
				getGatewayMacsecCakOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getGatewayMacsecCakOptionsModel).ToNot(BeNil())
				Expect(getGatewayMacsecCakOptionsModel.ID).To(Equal(core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")))
				Expect(getGatewayMacsecCakOptionsModel.CakID).To(Equal(core.StringPtr("ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4")))
				Expect(getGatewayMacsecCakOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetGatewayMacsecOptions successfully`, func() {
				// Construct an instance of the GetGatewayMacsecOptions model
				id := "0a06fb9b-820f-4c44-8a31-77f1f0806d28"
				getGatewayMacsecOptionsModel := directLinkService.NewGetGatewayMacsecOptions(id)
				getGatewayMacsecOptionsModel.SetID("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				getGatewayMacsecOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getGatewayMacsecOptionsModel).ToNot(BeNil())
				Expect(getGatewayMacsecOptionsModel.ID).To(Equal(core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")))
				Expect(getGatewayMacsecOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetGatewayOptions successfully`, func() {
				// Construct an instance of the GetGatewayOptions model
				id := "0a06fb9b-820f-4c44-8a31-77f1f0806d28"
				getGatewayOptionsModel := directLinkService.NewGetGatewayOptions(id)
				getGatewayOptionsModel.SetID("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				getGatewayOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getGatewayOptionsModel).ToNot(BeNil())
				Expect(getGatewayOptionsModel.ID).To(Equal(core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")))
				Expect(getGatewayOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetGatewayRouteReportOptions successfully`, func() {
				// Construct an instance of the GetGatewayRouteReportOptions model
				gatewayID := "0a06fb9b-820f-4c44-8a31-77f1f0806d28"
				id := "0a06fb9b-820f-4c44-8a31-77f1f0806d28"
				getGatewayRouteReportOptionsModel := directLinkService.NewGetGatewayRouteReportOptions(gatewayID, id)
				getGatewayRouteReportOptionsModel.SetGatewayID("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				getGatewayRouteReportOptionsModel.SetID("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				getGatewayRouteReportOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getGatewayRouteReportOptionsModel).ToNot(BeNil())
				Expect(getGatewayRouteReportOptionsModel.GatewayID).To(Equal(core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")))
				Expect(getGatewayRouteReportOptionsModel.ID).To(Equal(core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")))
				Expect(getGatewayRouteReportOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetGatewayStatisticsOptions successfully`, func() {
				// Construct an instance of the GetGatewayStatisticsOptions model
				id := "0a06fb9b-820f-4c44-8a31-77f1f0806d28"
				typeVar := "macsec_mka_session"
				getGatewayStatisticsOptionsModel := directLinkService.NewGetGatewayStatisticsOptions(id, typeVar)
				getGatewayStatisticsOptionsModel.SetID("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				getGatewayStatisticsOptionsModel.SetType("macsec_mka_session")
				getGatewayStatisticsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getGatewayStatisticsOptionsModel).ToNot(BeNil())
				Expect(getGatewayStatisticsOptionsModel.ID).To(Equal(core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")))
				Expect(getGatewayStatisticsOptionsModel.Type).To(Equal(core.StringPtr("macsec_mka_session")))
				Expect(getGatewayStatisticsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetGatewayStatusOptions successfully`, func() {
				// Construct an instance of the GetGatewayStatusOptions model
				id := "0a06fb9b-820f-4c44-8a31-77f1f0806d28"
				getGatewayStatusOptionsModel := directLinkService.NewGetGatewayStatusOptions(id)
				getGatewayStatusOptionsModel.SetID("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				getGatewayStatusOptionsModel.SetType("bgp")
				getGatewayStatusOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getGatewayStatusOptionsModel).ToNot(BeNil())
				Expect(getGatewayStatusOptionsModel.ID).To(Equal(core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")))
				Expect(getGatewayStatusOptionsModel.Type).To(Equal(core.StringPtr("bgp")))
				Expect(getGatewayStatusOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetGatewayVirtualConnectionOptions successfully`, func() {
				// Construct an instance of the GetGatewayVirtualConnectionOptions model
				gatewayID := "0a06fb9b-820f-4c44-8a31-77f1f0806d28"
				id := "0a06fb9b-820f-4c44-8a31-77f1f0806d28"
				getGatewayVirtualConnectionOptionsModel := directLinkService.NewGetGatewayVirtualConnectionOptions(gatewayID, id)
				getGatewayVirtualConnectionOptionsModel.SetGatewayID("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				getGatewayVirtualConnectionOptionsModel.SetID("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				getGatewayVirtualConnectionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getGatewayVirtualConnectionOptionsModel).ToNot(BeNil())
				Expect(getGatewayVirtualConnectionOptionsModel.GatewayID).To(Equal(core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")))
				Expect(getGatewayVirtualConnectionOptionsModel.ID).To(Equal(core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")))
				Expect(getGatewayVirtualConnectionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetPortOptions successfully`, func() {
				// Construct an instance of the GetPortOptions model
				id := "0a06fb9b-820f-4c44-8a31-77f1f0806d28"
				getPortOptionsModel := directLinkService.NewGetPortOptions(id)
				getPortOptionsModel.SetID("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				getPortOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getPortOptionsModel).ToNot(BeNil())
				Expect(getPortOptionsModel.ID).To(Equal(core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")))
				Expect(getPortOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewHpcsKeyIdentity successfully`, func() {
				crn := "crn:v1:bluemix:public:hs-crypto:us-south:a/4111d05f36894e3cb9b46a43556d9000:abc111b8-37aa-4034-9def-f2607c87aaaa:key:bbb222bc-430a-4de9-9aad-84e5bb022222"
				_model, err := directLinkService.NewHpcsKeyIdentity(crn)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewListGatewayAsPrependsOptions successfully`, func() {
				// Construct an instance of the ListGatewayAsPrependsOptions model
				gatewayID := "0a06fb9b-820f-4c44-8a31-77f1f0806d28"
				listGatewayAsPrependsOptionsModel := directLinkService.NewListGatewayAsPrependsOptions(gatewayID)
				listGatewayAsPrependsOptionsModel.SetGatewayID("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				listGatewayAsPrependsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listGatewayAsPrependsOptionsModel).ToNot(BeNil())
				Expect(listGatewayAsPrependsOptionsModel.GatewayID).To(Equal(core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")))
				Expect(listGatewayAsPrependsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListGatewayCompletionNoticeOptions successfully`, func() {
				// Construct an instance of the ListGatewayCompletionNoticeOptions model
				id := "0a06fb9b-820f-4c44-8a31-77f1f0806d28"
				listGatewayCompletionNoticeOptionsModel := directLinkService.NewListGatewayCompletionNoticeOptions(id)
				listGatewayCompletionNoticeOptionsModel.SetID("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				listGatewayCompletionNoticeOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listGatewayCompletionNoticeOptionsModel).ToNot(BeNil())
				Expect(listGatewayCompletionNoticeOptionsModel.ID).To(Equal(core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")))
				Expect(listGatewayCompletionNoticeOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListGatewayExportRouteFiltersOptions successfully`, func() {
				// Construct an instance of the ListGatewayExportRouteFiltersOptions model
				gatewayID := "0a06fb9b-820f-4c44-8a31-77f1f0806d28"
				listGatewayExportRouteFiltersOptionsModel := directLinkService.NewListGatewayExportRouteFiltersOptions(gatewayID)
				listGatewayExportRouteFiltersOptionsModel.SetGatewayID("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				listGatewayExportRouteFiltersOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listGatewayExportRouteFiltersOptionsModel).ToNot(BeNil())
				Expect(listGatewayExportRouteFiltersOptionsModel.GatewayID).To(Equal(core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")))
				Expect(listGatewayExportRouteFiltersOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListGatewayImportRouteFiltersOptions successfully`, func() {
				// Construct an instance of the ListGatewayImportRouteFiltersOptions model
				gatewayID := "0a06fb9b-820f-4c44-8a31-77f1f0806d28"
				listGatewayImportRouteFiltersOptionsModel := directLinkService.NewListGatewayImportRouteFiltersOptions(gatewayID)
				listGatewayImportRouteFiltersOptionsModel.SetGatewayID("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				listGatewayImportRouteFiltersOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listGatewayImportRouteFiltersOptionsModel).ToNot(BeNil())
				Expect(listGatewayImportRouteFiltersOptionsModel.GatewayID).To(Equal(core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")))
				Expect(listGatewayImportRouteFiltersOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListGatewayLetterOfAuthorizationOptions successfully`, func() {
				// Construct an instance of the ListGatewayLetterOfAuthorizationOptions model
				id := "0a06fb9b-820f-4c44-8a31-77f1f0806d28"
				listGatewayLetterOfAuthorizationOptionsModel := directLinkService.NewListGatewayLetterOfAuthorizationOptions(id)
				listGatewayLetterOfAuthorizationOptionsModel.SetID("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				listGatewayLetterOfAuthorizationOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listGatewayLetterOfAuthorizationOptionsModel).ToNot(BeNil())
				Expect(listGatewayLetterOfAuthorizationOptionsModel.ID).To(Equal(core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")))
				Expect(listGatewayLetterOfAuthorizationOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListGatewayMacsecCaksOptions successfully`, func() {
				// Construct an instance of the ListGatewayMacsecCaksOptions model
				id := "0a06fb9b-820f-4c44-8a31-77f1f0806d28"
				listGatewayMacsecCaksOptionsModel := directLinkService.NewListGatewayMacsecCaksOptions(id)
				listGatewayMacsecCaksOptionsModel.SetID("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				listGatewayMacsecCaksOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listGatewayMacsecCaksOptionsModel).ToNot(BeNil())
				Expect(listGatewayMacsecCaksOptionsModel.ID).To(Equal(core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")))
				Expect(listGatewayMacsecCaksOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListGatewayRouteReportsOptions successfully`, func() {
				// Construct an instance of the ListGatewayRouteReportsOptions model
				gatewayID := "0a06fb9b-820f-4c44-8a31-77f1f0806d28"
				listGatewayRouteReportsOptionsModel := directLinkService.NewListGatewayRouteReportsOptions(gatewayID)
				listGatewayRouteReportsOptionsModel.SetGatewayID("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				listGatewayRouteReportsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listGatewayRouteReportsOptionsModel).ToNot(BeNil())
				Expect(listGatewayRouteReportsOptionsModel.GatewayID).To(Equal(core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")))
				Expect(listGatewayRouteReportsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListGatewayVirtualConnectionsOptions successfully`, func() {
				// Construct an instance of the ListGatewayVirtualConnectionsOptions model
				gatewayID := "0a06fb9b-820f-4c44-8a31-77f1f0806d28"
				listGatewayVirtualConnectionsOptionsModel := directLinkService.NewListGatewayVirtualConnectionsOptions(gatewayID)
				listGatewayVirtualConnectionsOptionsModel.SetGatewayID("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				listGatewayVirtualConnectionsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listGatewayVirtualConnectionsOptionsModel).ToNot(BeNil())
				Expect(listGatewayVirtualConnectionsOptionsModel.GatewayID).To(Equal(core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")))
				Expect(listGatewayVirtualConnectionsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListGatewaysOptions successfully`, func() {
				// Construct an instance of the ListGatewaysOptions model
				listGatewaysOptionsModel := directLinkService.NewListGatewaysOptions()
				listGatewaysOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listGatewaysOptionsModel).ToNot(BeNil())
				Expect(listGatewaysOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListOfferingTypeLocationCrossConnectRoutersOptions successfully`, func() {
				// Construct an instance of the ListOfferingTypeLocationCrossConnectRoutersOptions model
				offeringType := "dedicated"
				locationName := "testString"
				listOfferingTypeLocationCrossConnectRoutersOptionsModel := directLinkService.NewListOfferingTypeLocationCrossConnectRoutersOptions(offeringType, locationName)
				listOfferingTypeLocationCrossConnectRoutersOptionsModel.SetOfferingType("dedicated")
				listOfferingTypeLocationCrossConnectRoutersOptionsModel.SetLocationName("testString")
				listOfferingTypeLocationCrossConnectRoutersOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listOfferingTypeLocationCrossConnectRoutersOptionsModel).ToNot(BeNil())
				Expect(listOfferingTypeLocationCrossConnectRoutersOptionsModel.OfferingType).To(Equal(core.StringPtr("dedicated")))
				Expect(listOfferingTypeLocationCrossConnectRoutersOptionsModel.LocationName).To(Equal(core.StringPtr("testString")))
				Expect(listOfferingTypeLocationCrossConnectRoutersOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListOfferingTypeLocationsOptions successfully`, func() {
				// Construct an instance of the ListOfferingTypeLocationsOptions model
				offeringType := "dedicated"
				listOfferingTypeLocationsOptionsModel := directLinkService.NewListOfferingTypeLocationsOptions(offeringType)
				listOfferingTypeLocationsOptionsModel.SetOfferingType("dedicated")
				listOfferingTypeLocationsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listOfferingTypeLocationsOptionsModel).ToNot(BeNil())
				Expect(listOfferingTypeLocationsOptionsModel.OfferingType).To(Equal(core.StringPtr("dedicated")))
				Expect(listOfferingTypeLocationsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListOfferingTypeSpeedsOptions successfully`, func() {
				// Construct an instance of the ListOfferingTypeSpeedsOptions model
				offeringType := "dedicated"
				listOfferingTypeSpeedsOptionsModel := directLinkService.NewListOfferingTypeSpeedsOptions(offeringType)
				listOfferingTypeSpeedsOptionsModel.SetOfferingType("dedicated")
				listOfferingTypeSpeedsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listOfferingTypeSpeedsOptionsModel).ToNot(BeNil())
				Expect(listOfferingTypeSpeedsOptionsModel.OfferingType).To(Equal(core.StringPtr("dedicated")))
				Expect(listOfferingTypeSpeedsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListPortsOptions successfully`, func() {
				// Construct an instance of the ListPortsOptions model
				listPortsOptionsModel := directLinkService.NewListPortsOptions()
				listPortsOptionsModel.SetStart("testString")
				listPortsOptionsModel.SetLimit(int64(10))
				listPortsOptionsModel.SetLocationName("testString")
				listPortsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listPortsOptionsModel).ToNot(BeNil())
				Expect(listPortsOptionsModel.Start).To(Equal(core.StringPtr("testString")))
				Expect(listPortsOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(10))))
				Expect(listPortsOptionsModel.LocationName).To(Equal(core.StringPtr("testString")))
				Expect(listPortsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewReplaceGatewayAsPrependsOptions successfully`, func() {
				// Construct an instance of the AsPrependPrefixArrayTemplate model
				asPrependPrefixArrayTemplateModel := new(directlinkv1.AsPrependPrefixArrayTemplate)
				Expect(asPrependPrefixArrayTemplateModel).ToNot(BeNil())
				asPrependPrefixArrayTemplateModel.Length = core.Int64Ptr(int64(4))
				asPrependPrefixArrayTemplateModel.Policy = core.StringPtr("import")
				asPrependPrefixArrayTemplateModel.SpecificPrefixes = []string{"192.168.3.0/24"}
				Expect(asPrependPrefixArrayTemplateModel.Length).To(Equal(core.Int64Ptr(int64(4))))
				Expect(asPrependPrefixArrayTemplateModel.Policy).To(Equal(core.StringPtr("import")))
				Expect(asPrependPrefixArrayTemplateModel.SpecificPrefixes).To(Equal([]string{"192.168.3.0/24"}))

				// Construct an instance of the ReplaceGatewayAsPrependsOptions model
				gatewayID := "0a06fb9b-820f-4c44-8a31-77f1f0806d28"
				ifMatch := "W/\"96d225c4-56bd-43d9-98fc-d7148e5c5028\""
				replaceGatewayAsPrependsOptionsModel := directLinkService.NewReplaceGatewayAsPrependsOptions(gatewayID, ifMatch)
				replaceGatewayAsPrependsOptionsModel.SetGatewayID("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				replaceGatewayAsPrependsOptionsModel.SetIfMatch("W/\"96d225c4-56bd-43d9-98fc-d7148e5c5028\"")
				replaceGatewayAsPrependsOptionsModel.SetAsPrepends([]directlinkv1.AsPrependPrefixArrayTemplate{*asPrependPrefixArrayTemplateModel})
				replaceGatewayAsPrependsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(replaceGatewayAsPrependsOptionsModel).ToNot(BeNil())
				Expect(replaceGatewayAsPrependsOptionsModel.GatewayID).To(Equal(core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")))
				Expect(replaceGatewayAsPrependsOptionsModel.IfMatch).To(Equal(core.StringPtr("W/\"96d225c4-56bd-43d9-98fc-d7148e5c5028\"")))
				Expect(replaceGatewayAsPrependsOptionsModel.AsPrepends).To(Equal([]directlinkv1.AsPrependPrefixArrayTemplate{*asPrependPrefixArrayTemplateModel}))
				Expect(replaceGatewayAsPrependsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewReplaceGatewayExportRouteFiltersOptions successfully`, func() {
				// Construct an instance of the GatewayTemplateRouteFilter model
				gatewayTemplateRouteFilterModel := new(directlinkv1.GatewayTemplateRouteFilter)
				Expect(gatewayTemplateRouteFilterModel).ToNot(BeNil())
				gatewayTemplateRouteFilterModel.Action = core.StringPtr("permit")
				gatewayTemplateRouteFilterModel.Ge = core.Int64Ptr(int64(25))
				gatewayTemplateRouteFilterModel.Le = core.Int64Ptr(int64(30))
				gatewayTemplateRouteFilterModel.Prefix = core.StringPtr("192.168.100.0/24")
				Expect(gatewayTemplateRouteFilterModel.Action).To(Equal(core.StringPtr("permit")))
				Expect(gatewayTemplateRouteFilterModel.Ge).To(Equal(core.Int64Ptr(int64(25))))
				Expect(gatewayTemplateRouteFilterModel.Le).To(Equal(core.Int64Ptr(int64(30))))
				Expect(gatewayTemplateRouteFilterModel.Prefix).To(Equal(core.StringPtr("192.168.100.0/24")))

				// Construct an instance of the ReplaceGatewayExportRouteFiltersOptions model
				gatewayID := "0a06fb9b-820f-4c44-8a31-77f1f0806d28"
				ifMatch := "W/\"96d225c4-56bd-43d9-98fc-d7148e5c5028\""
				replaceGatewayExportRouteFiltersOptionsModel := directLinkService.NewReplaceGatewayExportRouteFiltersOptions(gatewayID, ifMatch)
				replaceGatewayExportRouteFiltersOptionsModel.SetGatewayID("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				replaceGatewayExportRouteFiltersOptionsModel.SetIfMatch("W/\"96d225c4-56bd-43d9-98fc-d7148e5c5028\"")
				replaceGatewayExportRouteFiltersOptionsModel.SetExportRouteFilters([]directlinkv1.GatewayTemplateRouteFilter{*gatewayTemplateRouteFilterModel})
				replaceGatewayExportRouteFiltersOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(replaceGatewayExportRouteFiltersOptionsModel).ToNot(BeNil())
				Expect(replaceGatewayExportRouteFiltersOptionsModel.GatewayID).To(Equal(core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")))
				Expect(replaceGatewayExportRouteFiltersOptionsModel.IfMatch).To(Equal(core.StringPtr("W/\"96d225c4-56bd-43d9-98fc-d7148e5c5028\"")))
				Expect(replaceGatewayExportRouteFiltersOptionsModel.ExportRouteFilters).To(Equal([]directlinkv1.GatewayTemplateRouteFilter{*gatewayTemplateRouteFilterModel}))
				Expect(replaceGatewayExportRouteFiltersOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewReplaceGatewayImportRouteFiltersOptions successfully`, func() {
				// Construct an instance of the GatewayTemplateRouteFilter model
				gatewayTemplateRouteFilterModel := new(directlinkv1.GatewayTemplateRouteFilter)
				Expect(gatewayTemplateRouteFilterModel).ToNot(BeNil())
				gatewayTemplateRouteFilterModel.Action = core.StringPtr("permit")
				gatewayTemplateRouteFilterModel.Ge = core.Int64Ptr(int64(25))
				gatewayTemplateRouteFilterModel.Le = core.Int64Ptr(int64(30))
				gatewayTemplateRouteFilterModel.Prefix = core.StringPtr("192.168.100.0/24")
				Expect(gatewayTemplateRouteFilterModel.Action).To(Equal(core.StringPtr("permit")))
				Expect(gatewayTemplateRouteFilterModel.Ge).To(Equal(core.Int64Ptr(int64(25))))
				Expect(gatewayTemplateRouteFilterModel.Le).To(Equal(core.Int64Ptr(int64(30))))
				Expect(gatewayTemplateRouteFilterModel.Prefix).To(Equal(core.StringPtr("192.168.100.0/24")))

				// Construct an instance of the ReplaceGatewayImportRouteFiltersOptions model
				gatewayID := "0a06fb9b-820f-4c44-8a31-77f1f0806d28"
				ifMatch := "W/\"96d225c4-56bd-43d9-98fc-d7148e5c5028\""
				replaceGatewayImportRouteFiltersOptionsModel := directLinkService.NewReplaceGatewayImportRouteFiltersOptions(gatewayID, ifMatch)
				replaceGatewayImportRouteFiltersOptionsModel.SetGatewayID("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				replaceGatewayImportRouteFiltersOptionsModel.SetIfMatch("W/\"96d225c4-56bd-43d9-98fc-d7148e5c5028\"")
				replaceGatewayImportRouteFiltersOptionsModel.SetImportRouteFilters([]directlinkv1.GatewayTemplateRouteFilter{*gatewayTemplateRouteFilterModel})
				replaceGatewayImportRouteFiltersOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(replaceGatewayImportRouteFiltersOptionsModel).ToNot(BeNil())
				Expect(replaceGatewayImportRouteFiltersOptionsModel.GatewayID).To(Equal(core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")))
				Expect(replaceGatewayImportRouteFiltersOptionsModel.IfMatch).To(Equal(core.StringPtr("W/\"96d225c4-56bd-43d9-98fc-d7148e5c5028\"")))
				Expect(replaceGatewayImportRouteFiltersOptionsModel.ImportRouteFilters).To(Equal([]directlinkv1.GatewayTemplateRouteFilter{*gatewayTemplateRouteFilterModel}))
				Expect(replaceGatewayImportRouteFiltersOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewResourceGroupIdentity successfully`, func() {
				id := "56969d6043e9465c883cb9f7363e78e8"
				_model, err := directLinkService.NewResourceGroupIdentity(id)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewSetGatewayMacsecOptions successfully`, func() {
				// Construct an instance of the HpcsKeyIdentity model
				hpcsKeyIdentityModel := new(directlinkv1.HpcsKeyIdentity)
				Expect(hpcsKeyIdentityModel).ToNot(BeNil())
				hpcsKeyIdentityModel.Crn = core.StringPtr("crn:v1:bluemix:public:hs-crypto:us-south:a/4111d05f36894e3cb9b46a43556d9000:abc111b8-37aa-4034-9def-f2607c87aaaa:key:bbb222bc-430a-4de9-9aad-84e5bb022222")
				Expect(hpcsKeyIdentityModel.Crn).To(Equal(core.StringPtr("crn:v1:bluemix:public:hs-crypto:us-south:a/4111d05f36894e3cb9b46a43556d9000:abc111b8-37aa-4034-9def-f2607c87aaaa:key:bbb222bc-430a-4de9-9aad-84e5bb022222")))

				// Construct an instance of the GatewayMacsecCakPrototype model
				gatewayMacsecCakPrototypeModel := new(directlinkv1.GatewayMacsecCakPrototype)
				Expect(gatewayMacsecCakPrototypeModel).ToNot(BeNil())
				gatewayMacsecCakPrototypeModel.Key = hpcsKeyIdentityModel
				gatewayMacsecCakPrototypeModel.Name = core.StringPtr("1000")
				gatewayMacsecCakPrototypeModel.Session = core.StringPtr("primary")
				Expect(gatewayMacsecCakPrototypeModel.Key).To(Equal(hpcsKeyIdentityModel))
				Expect(gatewayMacsecCakPrototypeModel.Name).To(Equal(core.StringPtr("1000")))
				Expect(gatewayMacsecCakPrototypeModel.Session).To(Equal(core.StringPtr("primary")))

				// Construct an instance of the SakRekeyPrototypeSakRekeyTimerModePrototype model
				sakRekeyPrototypeModel := new(directlinkv1.SakRekeyPrototypeSakRekeyTimerModePrototype)
				Expect(sakRekeyPrototypeModel).ToNot(BeNil())
				sakRekeyPrototypeModel.Interval = core.Int64Ptr(int64(3600))
				sakRekeyPrototypeModel.Mode = core.StringPtr("timer")
				Expect(sakRekeyPrototypeModel.Interval).To(Equal(core.Int64Ptr(int64(3600))))
				Expect(sakRekeyPrototypeModel.Mode).To(Equal(core.StringPtr("timer")))

				// Construct an instance of the SetGatewayMacsecOptions model
				id := "0a06fb9b-820f-4c44-8a31-77f1f0806d28"
				setGatewayMacsecOptionsActive := true
				setGatewayMacsecOptionsCaks := []directlinkv1.GatewayMacsecCakPrototype{}
				var setGatewayMacsecOptionsSakRekey directlinkv1.SakRekeyPrototypeIntf = nil
				setGatewayMacsecOptionsSecurityPolicy := "must_secure"
				setGatewayMacsecOptionsModel := directLinkService.NewSetGatewayMacsecOptions(id, setGatewayMacsecOptionsActive, setGatewayMacsecOptionsCaks, setGatewayMacsecOptionsSakRekey, setGatewayMacsecOptionsSecurityPolicy)
				setGatewayMacsecOptionsModel.SetID("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				setGatewayMacsecOptionsModel.SetActive(true)
				setGatewayMacsecOptionsModel.SetCaks([]directlinkv1.GatewayMacsecCakPrototype{*gatewayMacsecCakPrototypeModel})
				setGatewayMacsecOptionsModel.SetSakRekey(sakRekeyPrototypeModel)
				setGatewayMacsecOptionsModel.SetSecurityPolicy("must_secure")
				setGatewayMacsecOptionsModel.SetWindowSize(int64(64))
				setGatewayMacsecOptionsModel.SetIfMatch("W/\"96d225c4-56bd-43d9-98fc-d7148e5c5028\"")
				setGatewayMacsecOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(setGatewayMacsecOptionsModel).ToNot(BeNil())
				Expect(setGatewayMacsecOptionsModel.ID).To(Equal(core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")))
				Expect(setGatewayMacsecOptionsModel.Active).To(Equal(core.BoolPtr(true)))
				Expect(setGatewayMacsecOptionsModel.Caks).To(Equal([]directlinkv1.GatewayMacsecCakPrototype{*gatewayMacsecCakPrototypeModel}))
				Expect(setGatewayMacsecOptionsModel.SakRekey).To(Equal(sakRekeyPrototypeModel))
				Expect(setGatewayMacsecOptionsModel.SecurityPolicy).To(Equal(core.StringPtr("must_secure")))
				Expect(setGatewayMacsecOptionsModel.WindowSize).To(Equal(core.Int64Ptr(int64(64))))
				Expect(setGatewayMacsecOptionsModel.IfMatch).To(Equal(core.StringPtr("W/\"96d225c4-56bd-43d9-98fc-d7148e5c5028\"")))
				Expect(setGatewayMacsecOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUnsetGatewayMacsecOptions successfully`, func() {
				// Construct an instance of the UnsetGatewayMacsecOptions model
				id := "0a06fb9b-820f-4c44-8a31-77f1f0806d28"
				unsetGatewayMacsecOptionsModel := directLinkService.NewUnsetGatewayMacsecOptions(id)
				unsetGatewayMacsecOptionsModel.SetID("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				unsetGatewayMacsecOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(unsetGatewayMacsecOptionsModel).ToNot(BeNil())
				Expect(unsetGatewayMacsecOptionsModel.ID).To(Equal(core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")))
				Expect(unsetGatewayMacsecOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateGatewayExportRouteFilterOptions successfully`, func() {
				// Construct an instance of the UpdateGatewayExportRouteFilterOptions model
				gatewayID := "0a06fb9b-820f-4c44-8a31-77f1f0806d28"
				id := "0a06fb9b-820f-4c44-8a31-77f1f0806d28"
				updateRouteFilterTemplatePatch := map[string]interface{}{"anyKey": "anyValue"}
				updateGatewayExportRouteFilterOptionsModel := directLinkService.NewUpdateGatewayExportRouteFilterOptions(gatewayID, id, updateRouteFilterTemplatePatch)
				updateGatewayExportRouteFilterOptionsModel.SetGatewayID("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				updateGatewayExportRouteFilterOptionsModel.SetID("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				updateGatewayExportRouteFilterOptionsModel.SetUpdateRouteFilterTemplatePatch(map[string]interface{}{"anyKey": "anyValue"})
				updateGatewayExportRouteFilterOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateGatewayExportRouteFilterOptionsModel).ToNot(BeNil())
				Expect(updateGatewayExportRouteFilterOptionsModel.GatewayID).To(Equal(core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")))
				Expect(updateGatewayExportRouteFilterOptionsModel.ID).To(Equal(core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")))
				Expect(updateGatewayExportRouteFilterOptionsModel.UpdateRouteFilterTemplatePatch).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))
				Expect(updateGatewayExportRouteFilterOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateGatewayImportRouteFilterOptions successfully`, func() {
				// Construct an instance of the UpdateGatewayImportRouteFilterOptions model
				gatewayID := "0a06fb9b-820f-4c44-8a31-77f1f0806d28"
				id := "0a06fb9b-820f-4c44-8a31-77f1f0806d28"
				updateRouteFilterTemplatePatch := map[string]interface{}{"anyKey": "anyValue"}
				updateGatewayImportRouteFilterOptionsModel := directLinkService.NewUpdateGatewayImportRouteFilterOptions(gatewayID, id, updateRouteFilterTemplatePatch)
				updateGatewayImportRouteFilterOptionsModel.SetGatewayID("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				updateGatewayImportRouteFilterOptionsModel.SetID("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				updateGatewayImportRouteFilterOptionsModel.SetUpdateRouteFilterTemplatePatch(map[string]interface{}{"anyKey": "anyValue"})
				updateGatewayImportRouteFilterOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateGatewayImportRouteFilterOptionsModel).ToNot(BeNil())
				Expect(updateGatewayImportRouteFilterOptionsModel.GatewayID).To(Equal(core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")))
				Expect(updateGatewayImportRouteFilterOptionsModel.ID).To(Equal(core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")))
				Expect(updateGatewayImportRouteFilterOptionsModel.UpdateRouteFilterTemplatePatch).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))
				Expect(updateGatewayImportRouteFilterOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateGatewayMacsecCakOptions successfully`, func() {
				// Construct an instance of the UpdateGatewayMacsecCakOptions model
				id := "0a06fb9b-820f-4c44-8a31-77f1f0806d28"
				cakID := "ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4"
				gatewayMacsecCakPatch := map[string]interface{}{"anyKey": "anyValue"}
				updateGatewayMacsecCakOptionsModel := directLinkService.NewUpdateGatewayMacsecCakOptions(id, cakID, gatewayMacsecCakPatch)
				updateGatewayMacsecCakOptionsModel.SetID("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				updateGatewayMacsecCakOptionsModel.SetCakID("ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4")
				updateGatewayMacsecCakOptionsModel.SetGatewayMacsecCakPatch(map[string]interface{}{"anyKey": "anyValue"})
				updateGatewayMacsecCakOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateGatewayMacsecCakOptionsModel).ToNot(BeNil())
				Expect(updateGatewayMacsecCakOptionsModel.ID).To(Equal(core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")))
				Expect(updateGatewayMacsecCakOptionsModel.CakID).To(Equal(core.StringPtr("ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4")))
				Expect(updateGatewayMacsecCakOptionsModel.GatewayMacsecCakPatch).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))
				Expect(updateGatewayMacsecCakOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateGatewayMacsecOptions successfully`, func() {
				// Construct an instance of the UpdateGatewayMacsecOptions model
				id := "0a06fb9b-820f-4c44-8a31-77f1f0806d28"
				gatewayMacsecPatch := map[string]interface{}{"anyKey": "anyValue"}
				updateGatewayMacsecOptionsModel := directLinkService.NewUpdateGatewayMacsecOptions(id, gatewayMacsecPatch)
				updateGatewayMacsecOptionsModel.SetID("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				updateGatewayMacsecOptionsModel.SetGatewayMacsecPatch(map[string]interface{}{"anyKey": "anyValue"})
				updateGatewayMacsecOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateGatewayMacsecOptionsModel).ToNot(BeNil())
				Expect(updateGatewayMacsecOptionsModel.ID).To(Equal(core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")))
				Expect(updateGatewayMacsecOptionsModel.GatewayMacsecPatch).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))
				Expect(updateGatewayMacsecOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateGatewayOptions successfully`, func() {
				// Construct an instance of the UpdateGatewayOptions model
				id := "0a06fb9b-820f-4c44-8a31-77f1f0806d28"
				gatewayPatchTemplatePatch := map[string]interface{}{"anyKey": "anyValue"}
				updateGatewayOptionsModel := directLinkService.NewUpdateGatewayOptions(id, gatewayPatchTemplatePatch)
				updateGatewayOptionsModel.SetID("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				updateGatewayOptionsModel.SetGatewayPatchTemplatePatch(map[string]interface{}{"anyKey": "anyValue"})
				updateGatewayOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateGatewayOptionsModel).ToNot(BeNil())
				Expect(updateGatewayOptionsModel.ID).To(Equal(core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")))
				Expect(updateGatewayOptionsModel.GatewayPatchTemplatePatch).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))
				Expect(updateGatewayOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateGatewayVirtualConnectionOptions successfully`, func() {
				// Construct an instance of the UpdateGatewayVirtualConnectionOptions model
				gatewayID := "0a06fb9b-820f-4c44-8a31-77f1f0806d28"
				id := "0a06fb9b-820f-4c44-8a31-77f1f0806d28"
				gatewayVirtualConnectionPatchTemplatePatch := map[string]interface{}{"anyKey": "anyValue"}
				updateGatewayVirtualConnectionOptionsModel := directLinkService.NewUpdateGatewayVirtualConnectionOptions(gatewayID, id, gatewayVirtualConnectionPatchTemplatePatch)
				updateGatewayVirtualConnectionOptionsModel.SetGatewayID("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				updateGatewayVirtualConnectionOptionsModel.SetID("0a06fb9b-820f-4c44-8a31-77f1f0806d28")
				updateGatewayVirtualConnectionOptionsModel.SetGatewayVirtualConnectionPatchTemplatePatch(map[string]interface{}{"anyKey": "anyValue"})
				updateGatewayVirtualConnectionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateGatewayVirtualConnectionOptionsModel).ToNot(BeNil())
				Expect(updateGatewayVirtualConnectionOptionsModel.GatewayID).To(Equal(core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")))
				Expect(updateGatewayVirtualConnectionOptionsModel.ID).To(Equal(core.StringPtr("0a06fb9b-820f-4c44-8a31-77f1f0806d28")))
				Expect(updateGatewayVirtualConnectionOptionsModel.GatewayVirtualConnectionPatchTemplatePatch).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))
				Expect(updateGatewayVirtualConnectionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewAuthenticationKeyIdentityHpcsAuthenticationKeyIdentity successfully`, func() {
				crn := "crn:v1:bluemix:public:hs-crypto:us-south:a/4111d05f36894e3cb9b46a43556d9000:abc111b8-37aa-4034-9def-f2607c87aaaa:key:bbb222bc-430a-4de9-9aad-84e5bb022222"
				_model, err := directLinkService.NewAuthenticationKeyIdentityHpcsAuthenticationKeyIdentity(crn)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewAuthenticationKeyIdentityKeyProtectAuthenticationKeyIdentity successfully`, func() {
				crn := "crn:v1:bluemix:public:kms:us-south:a/4111d05f36894e3cb9b46a43556d9000:abc111b8-37aa-4034-9def-f2607c87aaaa:key:bbb222bc-430a-4de9-9aad-84e5bb022222"
				_model, err := directLinkService.NewAuthenticationKeyIdentityKeyProtectAuthenticationKeyIdentity(crn)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewGatewayTemplateGatewayTypeConnectTemplate successfully`, func() {
				bgpAsn := int64(64999)
				global := true
				metered := false
				name := "myGateway"
				speedMbps := int64(1000)
				typeVar := "dedicated"
				var port *directlinkv1.GatewayPortIdentity = nil
				_, err := directLinkService.NewGatewayTemplateGatewayTypeConnectTemplate(bgpAsn, global, metered, name, speedMbps, typeVar, port)
				Expect(err).ToNot(BeNil())
			})
			It(`Invoke NewGatewayTemplateGatewayTypeDedicatedTemplate successfully`, func() {
				bgpAsn := int64(64999)
				global := true
				metered := false
				name := "myGateway"
				speedMbps := int64(1000)
				typeVar := "dedicated"
				carrierName := "myCarrierName"
				crossConnectRouter := "xcr01.dal03"
				customerName := "newCustomerName"
				locationName := "dal03"
				_model, err := directLinkService.NewGatewayTemplateGatewayTypeDedicatedTemplate(bgpAsn, global, metered, name, speedMbps, typeVar, carrierName, crossConnectRouter, customerName, locationName)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewSakRekeyPatchSakRekeyPacketNumberRolloverModePatch successfully`, func() {
				mode := "packet_number_rollover"
				_model, err := directLinkService.NewSakRekeyPatchSakRekeyPacketNumberRolloverModePatch(mode)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewSakRekeyPatchSakRekeyTimerModePatch successfully`, func() {
				interval := int64(3600)
				mode := "timer"
				_model, err := directLinkService.NewSakRekeyPatchSakRekeyTimerModePatch(interval, mode)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewSakRekeyPrototypeSakRekeyPacketNumberRolloverModePrototype successfully`, func() {
				mode := "packet_number_rollover"
				_model, err := directLinkService.NewSakRekeyPrototypeSakRekeyPacketNumberRolloverModePrototype(mode)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewSakRekeyPrototypeSakRekeyTimerModePrototype successfully`, func() {
				interval := int64(3600)
				mode := "timer"
				_model, err := directLinkService.NewSakRekeyPrototypeSakRekeyTimerModePrototype(interval, mode)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
		})
	})
	Describe(`Model unmarshaling tests`, func() {
		It(`Invoke UnmarshalAsPrependPrefixArrayTemplate successfully`, func() {
			// Construct an instance of the model.
			model := new(directlinkv1.AsPrependPrefixArrayTemplate)
			model.Length = core.Int64Ptr(int64(4))
			model.Policy = core.StringPtr("import")
			model.SpecificPrefixes = []string{"192.168.3.0/24"}

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *directlinkv1.AsPrependPrefixArrayTemplate
			err = directlinkv1.UnmarshalAsPrependPrefixArrayTemplate(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalAsPrependTemplate successfully`, func() {
			// Construct an instance of the model.
			model := new(directlinkv1.AsPrependTemplate)
			model.Length = core.Int64Ptr(int64(4))
			model.Policy = core.StringPtr("import")
			model.Prefix = core.StringPtr("172.17.0.0/16")
			model.SpecificPrefixes = []string{"192.168.3.0/24"}

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *directlinkv1.AsPrependTemplate
			err = directlinkv1.UnmarshalAsPrependTemplate(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalAuthenticationKeyIdentity successfully`, func() {
			// Construct an instance of the model.
			model := new(directlinkv1.AuthenticationKeyIdentity)
			model.Crn = core.StringPtr("crn:v1:bluemix:public:kms:us-south:a/4111d05f36894e3cb9b46a43556d9000:abc111b8-37aa-4034-9def-f2607c87aaaa:key:bbb222bc-430a-4de9-9aad-84e5bb022222")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *directlinkv1.AuthenticationKeyIdentity
			err = directlinkv1.UnmarshalAuthenticationKeyIdentity(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalGatewayActionTemplateUpdatesItem successfully`, func() {
			// Construct an instance of the model.
			model := new(directlinkv1.GatewayActionTemplateUpdatesItem)
			model.SpeedMbps = core.Int64Ptr(int64(500))
			model.BgpCerCidr = core.StringPtr("169.254.0.10/30")
			model.BgpIbmCidr = core.StringPtr("169.254.0.9/30")
			model.BgpAsn = core.Int64Ptr(int64(64999))
			model.Vlan = core.Int64Ptr(int64(10))

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *directlinkv1.GatewayActionTemplateUpdatesItem
			err = directlinkv1.UnmarshalGatewayActionTemplateUpdatesItem(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalGatewayBfdConfigActionTemplate successfully`, func() {
			// Construct an instance of the model.
			model := new(directlinkv1.GatewayBfdConfigActionTemplate)
			model.Interval = core.Int64Ptr(int64(2000))
			model.Multiplier = core.Int64Ptr(int64(10))

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *directlinkv1.GatewayBfdConfigActionTemplate
			err = directlinkv1.UnmarshalGatewayBfdConfigActionTemplate(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalGatewayBfdConfigTemplate successfully`, func() {
			// Construct an instance of the model.
			model := new(directlinkv1.GatewayBfdConfigTemplate)
			model.Interval = core.Int64Ptr(int64(2000))
			model.Multiplier = core.Int64Ptr(int64(10))

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *directlinkv1.GatewayBfdConfigTemplate
			err = directlinkv1.UnmarshalGatewayBfdConfigTemplate(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalGatewayBfdPatchTemplate successfully`, func() {
			// Construct an instance of the model.
			model := new(directlinkv1.GatewayBfdPatchTemplate)
			model.Interval = core.Int64Ptr(int64(2000))
			model.Multiplier = core.Int64Ptr(int64(10))

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *directlinkv1.GatewayBfdPatchTemplate
			err = directlinkv1.UnmarshalGatewayBfdPatchTemplate(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalGatewayMacsecCakPatch successfully`, func() {
			// Construct an instance of the model.
			model := new(directlinkv1.GatewayMacsecCakPatch)
			model.Key = nil
			model.Name = core.StringPtr("1000")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *directlinkv1.GatewayMacsecCakPatch
			err = directlinkv1.UnmarshalGatewayMacsecCakPatch(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalGatewayMacsecCakPrototype successfully`, func() {
			// Construct an instance of the model.
			model := new(directlinkv1.GatewayMacsecCakPrototype)
			model.Key = nil
			model.Name = core.StringPtr("1000")
			model.Session = core.StringPtr("primary")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *directlinkv1.GatewayMacsecCakPrototype
			err = directlinkv1.UnmarshalGatewayMacsecCakPrototype(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalGatewayMacsecPatch successfully`, func() {
			// Construct an instance of the model.
			model := new(directlinkv1.GatewayMacsecPatch)
			model.Active = core.BoolPtr(true)
			model.SakRekey = nil
			model.SecurityPolicy = core.StringPtr("must_secure")
			model.WindowSize = core.Int64Ptr(int64(64))

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *directlinkv1.GatewayMacsecPatch
			err = directlinkv1.UnmarshalGatewayMacsecPatch(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalGatewayMacsecPrototype successfully`, func() {
			// Construct an instance of the model.
			model := new(directlinkv1.GatewayMacsecPrototype)
			model.Active = core.BoolPtr(true)
			model.Caks = nil
			model.SakRekey = nil
			model.SecurityPolicy = core.StringPtr("must_secure")
			model.WindowSize = core.Int64Ptr(int64(64))

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *directlinkv1.GatewayMacsecPrototype
			err = directlinkv1.UnmarshalGatewayMacsecPrototype(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalGatewayPatchTemplate successfully`, func() {
			// Construct an instance of the model.
			model := new(directlinkv1.GatewayPatchTemplate)
			model.AuthenticationKey = nil
			model.BfdConfig = nil
			model.BgpAsn = core.Int64Ptr(int64(64999))
			model.BgpCerCidr = core.StringPtr("169.254.0.10/30")
			model.BgpIbmCidr = core.StringPtr("169.254.0.9/30")
			model.ConnectionMode = core.StringPtr("transit")
			model.DefaultExportRouteFilter = core.StringPtr("permit")
			model.DefaultImportRouteFilter = core.StringPtr("permit")
			model.Global = core.BoolPtr(true)
			model.LoaRejectReason = core.StringPtr("The port mentioned was incorrect")
			model.Metered = core.BoolPtr(false)
			model.Name = core.StringPtr("testGateway")
			model.OperationalStatus = core.StringPtr("loa_accepted")
			model.PatchPanelCompletionNotice = core.StringPtr("patch panel configuration details")
			model.SpeedMbps = core.Int64Ptr(int64(1000))
			model.Vlan = core.Int64Ptr(int64(10))

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *directlinkv1.GatewayPatchTemplate
			err = directlinkv1.UnmarshalGatewayPatchTemplate(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalGatewayPortIdentity successfully`, func() {
			// Construct an instance of the model.
			model := new(directlinkv1.GatewayPortIdentity)
			model.ID = core.StringPtr("fffdcb1a-fee4-41c7-9e11-9cd99e65c777")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *directlinkv1.GatewayPortIdentity
			err = directlinkv1.UnmarshalGatewayPortIdentity(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalGatewayTemplate successfully`, func() {
			// Construct an instance of the model.
			model := new(directlinkv1.GatewayTemplate)
			model.AsPrepends = nil
			model.AuthenticationKey = nil
			model.BfdConfig = nil
			model.BgpAsn = core.Int64Ptr(int64(64999))
			model.BgpBaseCidr = core.StringPtr("testString")
			model.BgpCerCidr = core.StringPtr("169.254.0.10/30")
			model.BgpIbmCidr = core.StringPtr("169.254.0.9/30")
			model.ConnectionMode = core.StringPtr("transit")
			model.DefaultExportRouteFilter = core.StringPtr("permit")
			model.DefaultImportRouteFilter = core.StringPtr("permit")
			model.ExportRouteFilters = nil
			model.Global = core.BoolPtr(true)
			model.ImportRouteFilters = nil
			model.Metered = core.BoolPtr(false)
			model.Name = core.StringPtr("myGateway")
			model.PatchPanelCompletionNotice = core.StringPtr("patch panel configuration details")
			model.ResourceGroup = nil
			model.SpeedMbps = core.Int64Ptr(int64(1000))
			model.Type = core.StringPtr("dedicated")
			model.CarrierName = core.StringPtr("myCarrierName")
			model.CrossConnectRouter = core.StringPtr("xcr01.dal03")
			model.CustomerName = core.StringPtr("newCustomerName")
			model.LocationName = core.StringPtr("dal03")
			model.Macsec = nil
			model.MacsecCapability = core.StringPtr("non_macsec")
			model.Vlan = core.Int64Ptr(int64(10))
			model.Port = nil

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *directlinkv1.GatewayTemplate
			err = directlinkv1.UnmarshalGatewayTemplate(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalGatewayTemplateRouteFilter successfully`, func() {
			// Construct an instance of the model.
			model := new(directlinkv1.GatewayTemplateRouteFilter)
			model.Action = core.StringPtr("permit")
			model.Ge = core.Int64Ptr(int64(25))
			model.Le = core.Int64Ptr(int64(30))
			model.Prefix = core.StringPtr("192.168.100.0/24")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *directlinkv1.GatewayTemplateRouteFilter
			err = directlinkv1.UnmarshalGatewayTemplateRouteFilter(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalGatewayVirtualConnectionPatchTemplate successfully`, func() {
			// Construct an instance of the model.
			model := new(directlinkv1.GatewayVirtualConnectionPatchTemplate)
			model.Name = core.StringPtr("newConnectionName")
			model.Status = core.StringPtr("attached")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *directlinkv1.GatewayVirtualConnectionPatchTemplate
			err = directlinkv1.UnmarshalGatewayVirtualConnectionPatchTemplate(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalHpcsKeyIdentity successfully`, func() {
			// Construct an instance of the model.
			model := new(directlinkv1.HpcsKeyIdentity)
			model.Crn = core.StringPtr("crn:v1:bluemix:public:hs-crypto:us-south:a/4111d05f36894e3cb9b46a43556d9000:abc111b8-37aa-4034-9def-f2607c87aaaa:key:bbb222bc-430a-4de9-9aad-84e5bb022222")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *directlinkv1.HpcsKeyIdentity
			err = directlinkv1.UnmarshalHpcsKeyIdentity(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalResourceGroupIdentity successfully`, func() {
			// Construct an instance of the model.
			model := new(directlinkv1.ResourceGroupIdentity)
			model.ID = core.StringPtr("56969d6043e9465c883cb9f7363e78e8")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *directlinkv1.ResourceGroupIdentity
			err = directlinkv1.UnmarshalResourceGroupIdentity(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalSakRekeyPatch successfully`, func() {
			// Construct an instance of the model.
			model := new(directlinkv1.SakRekeyPatch)
			model.Interval = core.Int64Ptr(int64(3600))
			model.Mode = core.StringPtr("timer")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *directlinkv1.SakRekeyPatch
			err = directlinkv1.UnmarshalSakRekeyPatch(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalSakRekeyPrototype successfully`, func() {
			// Construct an instance of the model.
			model := new(directlinkv1.SakRekeyPrototype)
			model.Interval = core.Int64Ptr(int64(3600))
			model.Mode = core.StringPtr("timer")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *directlinkv1.SakRekeyPrototype
			err = directlinkv1.UnmarshalSakRekeyPrototype(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalUpdateRouteFilterTemplate successfully`, func() {
			// Construct an instance of the model.
			model := new(directlinkv1.UpdateRouteFilterTemplate)
			model.Action = core.StringPtr("permit")
			model.Before = core.StringPtr("1a15dcab-7e40-45e1-b7c5-bc690eaa9782")
			model.Ge = core.Int64Ptr(int64(25))
			model.Le = core.Int64Ptr(int64(30))
			model.Prefix = core.StringPtr("192.168.100.0/24")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *directlinkv1.UpdateRouteFilterTemplate
			err = directlinkv1.UnmarshalUpdateRouteFilterTemplate(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalAuthenticationKeyIdentityHpcsAuthenticationKeyIdentity successfully`, func() {
			// Construct an instance of the model.
			model := new(directlinkv1.AuthenticationKeyIdentityHpcsAuthenticationKeyIdentity)
			model.Crn = core.StringPtr("crn:v1:bluemix:public:hs-crypto:us-south:a/4111d05f36894e3cb9b46a43556d9000:abc111b8-37aa-4034-9def-f2607c87aaaa:key:bbb222bc-430a-4de9-9aad-84e5bb022222")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *directlinkv1.AuthenticationKeyIdentityHpcsAuthenticationKeyIdentity
			err = directlinkv1.UnmarshalAuthenticationKeyIdentityHpcsAuthenticationKeyIdentity(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalAuthenticationKeyIdentityKeyProtectAuthenticationKeyIdentity successfully`, func() {
			// Construct an instance of the model.
			model := new(directlinkv1.AuthenticationKeyIdentityKeyProtectAuthenticationKeyIdentity)
			model.Crn = core.StringPtr("crn:v1:bluemix:public:kms:us-south:a/4111d05f36894e3cb9b46a43556d9000:abc111b8-37aa-4034-9def-f2607c87aaaa:key:bbb222bc-430a-4de9-9aad-84e5bb022222")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *directlinkv1.AuthenticationKeyIdentityKeyProtectAuthenticationKeyIdentity
			err = directlinkv1.UnmarshalAuthenticationKeyIdentityKeyProtectAuthenticationKeyIdentity(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalGatewayActionTemplateUpdatesItemGatewayClientBGPASNUpdate successfully`, func() {
			// Construct an instance of the model.
			model := new(directlinkv1.GatewayActionTemplateUpdatesItemGatewayClientBGPASNUpdate)
			model.BgpAsn = core.Int64Ptr(int64(64999))

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *directlinkv1.GatewayActionTemplateUpdatesItemGatewayClientBGPASNUpdate
			err = directlinkv1.UnmarshalGatewayActionTemplateUpdatesItemGatewayClientBGPASNUpdate(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalGatewayActionTemplateUpdatesItemGatewayClientBGPIPUpdate successfully`, func() {
			// Construct an instance of the model.
			model := new(directlinkv1.GatewayActionTemplateUpdatesItemGatewayClientBGPIPUpdate)
			model.BgpCerCidr = core.StringPtr("169.254.0.10/30")
			model.BgpIbmCidr = core.StringPtr("169.254.0.9/30")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *directlinkv1.GatewayActionTemplateUpdatesItemGatewayClientBGPIPUpdate
			err = directlinkv1.UnmarshalGatewayActionTemplateUpdatesItemGatewayClientBGPIPUpdate(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalGatewayActionTemplateUpdatesItemGatewayClientSpeedUpdate successfully`, func() {
			// Construct an instance of the model.
			model := new(directlinkv1.GatewayActionTemplateUpdatesItemGatewayClientSpeedUpdate)
			model.SpeedMbps = core.Int64Ptr(int64(500))

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *directlinkv1.GatewayActionTemplateUpdatesItemGatewayClientSpeedUpdate
			err = directlinkv1.UnmarshalGatewayActionTemplateUpdatesItemGatewayClientSpeedUpdate(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalGatewayActionTemplateUpdatesItemGatewayClientVLANUpdate successfully`, func() {
			// Construct an instance of the model.
			model := new(directlinkv1.GatewayActionTemplateUpdatesItemGatewayClientVLANUpdate)
			model.Vlan = core.Int64Ptr(int64(10))

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *directlinkv1.GatewayActionTemplateUpdatesItemGatewayClientVLANUpdate
			err = directlinkv1.UnmarshalGatewayActionTemplateUpdatesItemGatewayClientVLANUpdate(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalGatewayTemplateGatewayTypeConnectTemplate successfully`, func() {
			// Construct an instance of the model.
			model := new(directlinkv1.GatewayTemplateGatewayTypeConnectTemplate)
			model.AsPrepends = nil
			model.AuthenticationKey = nil
			model.BfdConfig = nil
			model.BgpAsn = core.Int64Ptr(int64(64999))
			model.BgpBaseCidr = core.StringPtr("testString")
			model.BgpCerCidr = core.StringPtr("169.254.0.10/30")
			model.BgpIbmCidr = core.StringPtr("169.254.0.9/30")
			model.ConnectionMode = core.StringPtr("transit")
			model.DefaultExportRouteFilter = core.StringPtr("permit")
			model.DefaultImportRouteFilter = core.StringPtr("permit")
			model.ExportRouteFilters = nil
			model.Global = core.BoolPtr(true)
			model.ImportRouteFilters = nil
			model.Metered = core.BoolPtr(false)
			model.Name = core.StringPtr("myGateway")
			model.PatchPanelCompletionNotice = core.StringPtr("patch panel configuration details")
			model.ResourceGroup = nil
			model.SpeedMbps = core.Int64Ptr(int64(1000))
			model.Type = core.StringPtr("dedicated")
			model.Port = nil

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *directlinkv1.GatewayTemplateGatewayTypeConnectTemplate
			err = directlinkv1.UnmarshalGatewayTemplateGatewayTypeConnectTemplate(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalGatewayTemplateGatewayTypeDedicatedTemplate successfully`, func() {
			// Construct an instance of the model.
			model := new(directlinkv1.GatewayTemplateGatewayTypeDedicatedTemplate)
			model.AsPrepends = nil
			model.AuthenticationKey = nil
			model.BfdConfig = nil
			model.BgpAsn = core.Int64Ptr(int64(64999))
			model.BgpBaseCidr = core.StringPtr("testString")
			model.BgpCerCidr = core.StringPtr("169.254.0.10/30")
			model.BgpIbmCidr = core.StringPtr("169.254.0.9/30")
			model.ConnectionMode = core.StringPtr("transit")
			model.DefaultExportRouteFilter = core.StringPtr("permit")
			model.DefaultImportRouteFilter = core.StringPtr("permit")
			model.ExportRouteFilters = nil
			model.Global = core.BoolPtr(true)
			model.ImportRouteFilters = nil
			model.Metered = core.BoolPtr(false)
			model.Name = core.StringPtr("myGateway")
			model.PatchPanelCompletionNotice = core.StringPtr("patch panel configuration details")
			model.ResourceGroup = nil
			model.SpeedMbps = core.Int64Ptr(int64(1000))
			model.Type = core.StringPtr("dedicated")
			model.CarrierName = core.StringPtr("myCarrierName")
			model.CrossConnectRouter = core.StringPtr("xcr01.dal03")
			model.CustomerName = core.StringPtr("newCustomerName")
			model.LocationName = core.StringPtr("dal03")
			model.Macsec = nil
			model.MacsecCapability = core.StringPtr("non_macsec")
			model.Vlan = core.Int64Ptr(int64(10))

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *directlinkv1.GatewayTemplateGatewayTypeDedicatedTemplate
			err = directlinkv1.UnmarshalGatewayTemplateGatewayTypeDedicatedTemplate(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalSakRekeyPatchSakRekeyPacketNumberRolloverModePatch successfully`, func() {
			// Construct an instance of the model.
			model := new(directlinkv1.SakRekeyPatchSakRekeyPacketNumberRolloverModePatch)
			model.Mode = core.StringPtr("packet_number_rollover")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *directlinkv1.SakRekeyPatchSakRekeyPacketNumberRolloverModePatch
			err = directlinkv1.UnmarshalSakRekeyPatchSakRekeyPacketNumberRolloverModePatch(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalSakRekeyPatchSakRekeyTimerModePatch successfully`, func() {
			// Construct an instance of the model.
			model := new(directlinkv1.SakRekeyPatchSakRekeyTimerModePatch)
			model.Interval = core.Int64Ptr(int64(3600))
			model.Mode = core.StringPtr("timer")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *directlinkv1.SakRekeyPatchSakRekeyTimerModePatch
			err = directlinkv1.UnmarshalSakRekeyPatchSakRekeyTimerModePatch(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalSakRekeyPrototypeSakRekeyPacketNumberRolloverModePrototype successfully`, func() {
			// Construct an instance of the model.
			model := new(directlinkv1.SakRekeyPrototypeSakRekeyPacketNumberRolloverModePrototype)
			model.Mode = core.StringPtr("packet_number_rollover")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *directlinkv1.SakRekeyPrototypeSakRekeyPacketNumberRolloverModePrototype
			err = directlinkv1.UnmarshalSakRekeyPrototypeSakRekeyPacketNumberRolloverModePrototype(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalSakRekeyPrototypeSakRekeyTimerModePrototype successfully`, func() {
			// Construct an instance of the model.
			model := new(directlinkv1.SakRekeyPrototypeSakRekeyTimerModePrototype)
			model.Interval = core.Int64Ptr(int64(3600))
			model.Mode = core.StringPtr("timer")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *directlinkv1.SakRekeyPrototypeSakRekeyTimerModePrototype
			err = directlinkv1.UnmarshalSakRekeyPrototypeSakRekeyTimerModePrototype(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
	})
	Describe(`Utility function tests`, func() {
		It(`Invoke CreateMockByteArray() successfully`, func() {
			mockByteArray := CreateMockByteArray("VGhpcyBpcyBhIHRlc3Qgb2YgdGhlIGVtZXJnZW5jeSBicm9hZGNhc3Qgc3lzdGVt")
			Expect(mockByteArray).ToNot(BeNil())
		})
		It(`Invoke CreateMockUUID() successfully`, func() {
			mockUUID := CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
			Expect(mockUUID).ToNot(BeNil())
		})
		It(`Invoke CreateMockReader() successfully`, func() {
			mockReader := CreateMockReader("This is a test.")
			Expect(mockReader).ToNot(BeNil())
		})
		It(`Invoke CreateMockDate() successfully`, func() {
			mockDate := CreateMockDate("2019-01-01")
			Expect(mockDate).ToNot(BeNil())
		})
		It(`Invoke CreateMockDateTime() successfully`, func() {
			mockDateTime := CreateMockDateTime("2019-01-01T12:00:00.000Z")
			Expect(mockDateTime).ToNot(BeNil())
		})
	})
})

//
// Utility functions used by the generated test code
//

func CreateMockByteArray(encodedString string) *[]byte {
	ba, err := base64.StdEncoding.DecodeString(encodedString)
	if err != nil {
		panic(err)
	}
	return &ba
}

func CreateMockUUID(mockData string) *strfmt.UUID {
	uuid := strfmt.UUID(mockData)
	return &uuid
}

func CreateMockReader(mockData string) io.ReadCloser {
	return io.NopCloser(bytes.NewReader([]byte(mockData)))
}

func CreateMockDate(mockData string) *strfmt.Date {
	d, err := core.ParseDate(mockData)
	if err != nil {
		return nil
	}
	return &d
}

func CreateMockDateTime(mockData string) *strfmt.DateTime {
	d, err := core.ParseDateTime(mockData)
	if err != nil {
		return nil
	}
	return &d
}

func SetTestEnvironment(testEnvironment map[string]string) {
	for key, value := range testEnvironment {
		os.Setenv(key, value)
	}
}

func ClearTestEnvironment(testEnvironment map[string]string) {
	for key := range testEnvironment {
		os.Unsetenv(key)
	}
}
