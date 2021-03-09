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

package directlinkv1_test

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"time"

	"github.com/IBM/go-sdk-core/v4/core"
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
			testService, testServiceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
				Version:       core.StringPtr(version),
			})
			Expect(testService).ToNot(BeNil())
			Expect(testServiceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			testService, testServiceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
				URL:     "{BAD_URL_STRING",
				Version: core.StringPtr(version),
			})
			Expect(testService).To(BeNil())
			Expect(testServiceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			testService, testServiceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
				URL:     "https://directlinkv1/api",
				Version: core.StringPtr(version),
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(testService).To(BeNil())
			Expect(testServiceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Validation Error`, func() {
			testService, testServiceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{})
			Expect(testService).To(BeNil())
			Expect(testServiceErr).ToNot(BeNil())
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
				testService, testServiceErr := directlinkv1.NewDirectLinkV1UsingExternalConfig(&directlinkv1.DirectLinkV1Options{
					Version: core.StringPtr(version),
				})
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := directlinkv1.NewDirectLinkV1UsingExternalConfig(&directlinkv1.DirectLinkV1Options{
					URL:     "https://testService/api",
					Version: core.StringPtr(version),
				})
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				Expect(testService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := directlinkv1.NewDirectLinkV1UsingExternalConfig(&directlinkv1.DirectLinkV1Options{
					Version: core.StringPtr(version),
				})
				err := testService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				Expect(testService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"DIRECT_LINK_URL":       "https://directlinkv1/api",
				"DIRECT_LINK_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			testService, testServiceErr := directlinkv1.NewDirectLinkV1UsingExternalConfig(&directlinkv1.DirectLinkV1Options{
				Version: core.StringPtr(version),
			})

			It(`Instantiate service client with error`, func() {
				Expect(testService).To(BeNil())
				Expect(testServiceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"DIRECT_LINK_AUTH_TYPE": "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			testService, testServiceErr := directlinkv1.NewDirectLinkV1UsingExternalConfig(&directlinkv1.DirectLinkV1Options{
				URL:     "{BAD_URL_STRING",
				Version: core.StringPtr(version),
			})

			It(`Instantiate service client with error`, func() {
				Expect(testService).To(BeNil())
				Expect(testServiceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`ListGateways(listGatewaysOptions *ListGatewaysOptions) - Operation response error`, func() {
		version := "testString"
		listGatewaysPath := "/gateways"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(listGatewaysPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListGateways with error: Operation response processing error`, func() {
				testService, testServiceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ListGatewaysOptions model
				listGatewaysOptionsModel := new(directlinkv1.ListGatewaysOptions)
				listGatewaysOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.ListGateways(listGatewaysOptionsModel)
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
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(listGatewaysPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"gateways": [{"bgp_asn": 64999, "bgp_base_cidr": "BgpBaseCidr", "bgp_cer_cidr": "10.254.30.78/30", "bgp_ibm_asn": 13884, "bgp_ibm_cidr": "10.254.30.77/30", "bgp_status": "active", "carrier_name": "myCarrierName", "change_request": {"type": "create_gateway"}, "completion_notice_reject_reason": "The completion notice file was blank", "created_at": "2019-01-01T12:00:00", "crn": "crn:v1:bluemix:public:directlink:dal03:a/4111d05f36894e3cb9b46a43556d9000::dedicated:ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4", "cross_connect_router": "xcr01.dal03", "customer_name": "newCustomerName", "global": true, "id": "ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4", "link_status": "up", "location_display_name": "Dallas 03", "location_name": "dal03", "macsec_config": {"active": true, "active_cak": {"crn": "crn:v1:bluemix:public:hs-crypto:us-south:a/4111d05f36894e3cb9b46a43556d9000:abc111b8-37aa-4034-9def-f2607c87aaaa:key:bbb222bc-430a-4de9-9aad-84e5bb022222", "status": "Status"}, "cipher_suite": "gcm_aes_xpn_256", "confidentiality_offset": 0, "cryptographic_algorithm": "aes_256_cmac", "fallback_cak": {"crn": "crn:v1:bluemix:public:hs-crypto:us-south:a/4111d05f36894e3cb9b46a43556d9000:abc111b8-37aa-4034-9def-f2607c87aaaa:key:bbb222bc-430a-4de9-9aad-84e5bb022222", "status": "Status"}, "key_server_priority": 255, "primary_cak": {"crn": "crn:v1:bluemix:public:hs-crypto:us-south:a/4111d05f36894e3cb9b46a43556d9000:abc111b8-37aa-4034-9def-f2607c87aaaa:key:bbb222bc-430a-4de9-9aad-84e5bb022222", "status": "Status"}, "sak_expiry_time": 3600, "security_policy": "should_secure", "status": "secured", "window_size": 64}, "metered": false, "name": "myGateway", "operational_status": "awaiting_completion_notice", "port": {"id": "54321b1a-fee4-41c7-9e11-9cd99e000aaa"}, "provider_api_managed": false, "resource_group": {"id": "56969d6043e9465c883cb9f7363e78e8"}, "speed_mbps": 1000, "type": "dedicated", "vlan": 10}]}`)
				}))
			})
			It(`Invoke ListGateways successfully`, func() {
				testService, testServiceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.ListGateways(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListGatewaysOptions model
				listGatewaysOptionsModel := new(directlinkv1.ListGatewaysOptions)
				listGatewaysOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.ListGateways(listGatewaysOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke ListGateways with error: Operation request error`, func() {
				testService, testServiceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ListGatewaysOptions model
				listGatewaysOptionsModel := new(directlinkv1.ListGatewaysOptions)
				listGatewaysOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.ListGateways(listGatewaysOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
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
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(createGatewayPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateGateway with error: Operation response processing error`, func() {
				testService, testServiceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GatewayMacsecConfigTemplateFallbackCak model
				gatewayMacsecConfigTemplateFallbackCakModel := new(directlinkv1.GatewayMacsecConfigTemplateFallbackCak)
				gatewayMacsecConfigTemplateFallbackCakModel.Crn = core.StringPtr("crn:v1:bluemix:public:hs-crypto:us-south:a/4111d05f36894e3cb9b46a43556d9000:abc111b8-37aa-4034-9def-f2607c87aaaa:key:bbb222bc-430a-4de9-9aad-84e5bb022222")

				// Construct an instance of the GatewayMacsecConfigTemplatePrimaryCak model
				gatewayMacsecConfigTemplatePrimaryCakModel := new(directlinkv1.GatewayMacsecConfigTemplatePrimaryCak)
				gatewayMacsecConfigTemplatePrimaryCakModel.Crn = core.StringPtr("crn:v1:bluemix:public:hs-crypto:us-south:a/4111d05f36894e3cb9b46a43556d9000:abc111b8-37aa-4034-9def-f2607c87aaaa:key:bbb222bc-430a-4de9-9aad-84e5bb022222")

				// Construct an instance of the GatewayMacsecConfigTemplate model
				gatewayMacsecConfigTemplateModel := new(directlinkv1.GatewayMacsecConfigTemplate)
				gatewayMacsecConfigTemplateModel.Active = core.BoolPtr(true)
				gatewayMacsecConfigTemplateModel.FallbackCak = gatewayMacsecConfigTemplateFallbackCakModel
				gatewayMacsecConfigTemplateModel.PrimaryCak = gatewayMacsecConfigTemplatePrimaryCakModel
				gatewayMacsecConfigTemplateModel.WindowSize = core.Int64Ptr(int64(148809600))

				// Construct an instance of the ResourceGroupIdentity model
				resourceGroupIdentityModel := new(directlinkv1.ResourceGroupIdentity)
				resourceGroupIdentityModel.ID = core.StringPtr("56969d6043e9465c883cb9f7363e78e8")

				// Construct an instance of the GatewayTemplateGatewayTypeDedicatedTemplate model
				gatewayTemplateModel := new(directlinkv1.GatewayTemplateGatewayTypeDedicatedTemplate)
				gatewayTemplateModel.BgpAsn = core.Int64Ptr(int64(64999))
				gatewayTemplateModel.BgpBaseCidr = core.StringPtr("testString")
				gatewayTemplateModel.BgpCerCidr = core.StringPtr("169.254.0.10/30")
				gatewayTemplateModel.BgpIbmCidr = core.StringPtr("169.254.0.9/30")
				gatewayTemplateModel.Global = core.BoolPtr(true)
				gatewayTemplateModel.Metered = core.BoolPtr(false)
				gatewayTemplateModel.Name = core.StringPtr("myGateway")
				gatewayTemplateModel.ResourceGroup = resourceGroupIdentityModel
				gatewayTemplateModel.SpeedMbps = core.Int64Ptr(int64(1000))
				gatewayTemplateModel.Type = core.StringPtr("dedicated")
				gatewayTemplateModel.CarrierName = core.StringPtr("myCarrierName")
				gatewayTemplateModel.CrossConnectRouter = core.StringPtr("xcr01.dal03")
				gatewayTemplateModel.CustomerName = core.StringPtr("newCustomerName")
				gatewayTemplateModel.LocationName = core.StringPtr("dal03")
				gatewayTemplateModel.MacsecConfig = gatewayMacsecConfigTemplateModel

				// Construct an instance of the CreateGatewayOptions model
				createGatewayOptionsModel := new(directlinkv1.CreateGatewayOptions)
				createGatewayOptionsModel.GatewayTemplate = gatewayTemplateModel
				createGatewayOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.CreateGateway(createGatewayOptionsModel)
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
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(createGatewayPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, `{"bgp_asn": 64999, "bgp_base_cidr": "BgpBaseCidr", "bgp_cer_cidr": "10.254.30.78/30", "bgp_ibm_asn": 13884, "bgp_ibm_cidr": "10.254.30.77/30", "bgp_status": "active", "carrier_name": "myCarrierName", "change_request": {"type": "create_gateway"}, "completion_notice_reject_reason": "The completion notice file was blank", "created_at": "2019-01-01T12:00:00", "crn": "crn:v1:bluemix:public:directlink:dal03:a/4111d05f36894e3cb9b46a43556d9000::dedicated:ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4", "cross_connect_router": "xcr01.dal03", "customer_name": "newCustomerName", "global": true, "id": "ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4", "link_status": "up", "location_display_name": "Dallas 03", "location_name": "dal03", "macsec_config": {"active": true, "active_cak": {"crn": "crn:v1:bluemix:public:hs-crypto:us-south:a/4111d05f36894e3cb9b46a43556d9000:abc111b8-37aa-4034-9def-f2607c87aaaa:key:bbb222bc-430a-4de9-9aad-84e5bb022222", "status": "Status"}, "cipher_suite": "gcm_aes_xpn_256", "confidentiality_offset": 0, "cryptographic_algorithm": "aes_256_cmac", "fallback_cak": {"crn": "crn:v1:bluemix:public:hs-crypto:us-south:a/4111d05f36894e3cb9b46a43556d9000:abc111b8-37aa-4034-9def-f2607c87aaaa:key:bbb222bc-430a-4de9-9aad-84e5bb022222", "status": "Status"}, "key_server_priority": 255, "primary_cak": {"crn": "crn:v1:bluemix:public:hs-crypto:us-south:a/4111d05f36894e3cb9b46a43556d9000:abc111b8-37aa-4034-9def-f2607c87aaaa:key:bbb222bc-430a-4de9-9aad-84e5bb022222", "status": "Status"}, "sak_expiry_time": 3600, "security_policy": "should_secure", "status": "secured", "window_size": 64}, "metered": false, "name": "myGateway", "operational_status": "awaiting_completion_notice", "port": {"id": "54321b1a-fee4-41c7-9e11-9cd99e000aaa"}, "provider_api_managed": false, "resource_group": {"id": "56969d6043e9465c883cb9f7363e78e8"}, "speed_mbps": 1000, "type": "dedicated", "vlan": 10}`)
				}))
			})
			It(`Invoke CreateGateway successfully`, func() {
				testService, testServiceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.CreateGateway(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GatewayMacsecConfigTemplateFallbackCak model
				gatewayMacsecConfigTemplateFallbackCakModel := new(directlinkv1.GatewayMacsecConfigTemplateFallbackCak)
				gatewayMacsecConfigTemplateFallbackCakModel.Crn = core.StringPtr("crn:v1:bluemix:public:hs-crypto:us-south:a/4111d05f36894e3cb9b46a43556d9000:abc111b8-37aa-4034-9def-f2607c87aaaa:key:bbb222bc-430a-4de9-9aad-84e5bb022222")

				// Construct an instance of the GatewayMacsecConfigTemplatePrimaryCak model
				gatewayMacsecConfigTemplatePrimaryCakModel := new(directlinkv1.GatewayMacsecConfigTemplatePrimaryCak)
				gatewayMacsecConfigTemplatePrimaryCakModel.Crn = core.StringPtr("crn:v1:bluemix:public:hs-crypto:us-south:a/4111d05f36894e3cb9b46a43556d9000:abc111b8-37aa-4034-9def-f2607c87aaaa:key:bbb222bc-430a-4de9-9aad-84e5bb022222")

				// Construct an instance of the GatewayMacsecConfigTemplate model
				gatewayMacsecConfigTemplateModel := new(directlinkv1.GatewayMacsecConfigTemplate)
				gatewayMacsecConfigTemplateModel.Active = core.BoolPtr(true)
				gatewayMacsecConfigTemplateModel.FallbackCak = gatewayMacsecConfigTemplateFallbackCakModel
				gatewayMacsecConfigTemplateModel.PrimaryCak = gatewayMacsecConfigTemplatePrimaryCakModel
				gatewayMacsecConfigTemplateModel.WindowSize = core.Int64Ptr(int64(148809600))

				// Construct an instance of the ResourceGroupIdentity model
				resourceGroupIdentityModel := new(directlinkv1.ResourceGroupIdentity)
				resourceGroupIdentityModel.ID = core.StringPtr("56969d6043e9465c883cb9f7363e78e8")

				// Construct an instance of the GatewayTemplateGatewayTypeDedicatedTemplate model
				gatewayTemplateModel := new(directlinkv1.GatewayTemplateGatewayTypeDedicatedTemplate)
				gatewayTemplateModel.BgpAsn = core.Int64Ptr(int64(64999))
				gatewayTemplateModel.BgpBaseCidr = core.StringPtr("testString")
				gatewayTemplateModel.BgpCerCidr = core.StringPtr("169.254.0.10/30")
				gatewayTemplateModel.BgpIbmCidr = core.StringPtr("169.254.0.9/30")
				gatewayTemplateModel.Global = core.BoolPtr(true)
				gatewayTemplateModel.Metered = core.BoolPtr(false)
				gatewayTemplateModel.Name = core.StringPtr("myGateway")
				gatewayTemplateModel.ResourceGroup = resourceGroupIdentityModel
				gatewayTemplateModel.SpeedMbps = core.Int64Ptr(int64(1000))
				gatewayTemplateModel.Type = core.StringPtr("dedicated")
				gatewayTemplateModel.CarrierName = core.StringPtr("myCarrierName")
				gatewayTemplateModel.CrossConnectRouter = core.StringPtr("xcr01.dal03")
				gatewayTemplateModel.CustomerName = core.StringPtr("newCustomerName")
				gatewayTemplateModel.LocationName = core.StringPtr("dal03")
				gatewayTemplateModel.MacsecConfig = gatewayMacsecConfigTemplateModel

				// Construct an instance of the CreateGatewayOptions model
				createGatewayOptionsModel := new(directlinkv1.CreateGatewayOptions)
				createGatewayOptionsModel.GatewayTemplate = gatewayTemplateModel
				createGatewayOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.CreateGateway(createGatewayOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke CreateGateway with error: Operation validation and request error`, func() {
				testService, testServiceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GatewayMacsecConfigTemplateFallbackCak model
				gatewayMacsecConfigTemplateFallbackCakModel := new(directlinkv1.GatewayMacsecConfigTemplateFallbackCak)
				gatewayMacsecConfigTemplateFallbackCakModel.Crn = core.StringPtr("crn:v1:bluemix:public:hs-crypto:us-south:a/4111d05f36894e3cb9b46a43556d9000:abc111b8-37aa-4034-9def-f2607c87aaaa:key:bbb222bc-430a-4de9-9aad-84e5bb022222")

				// Construct an instance of the GatewayMacsecConfigTemplatePrimaryCak model
				gatewayMacsecConfigTemplatePrimaryCakModel := new(directlinkv1.GatewayMacsecConfigTemplatePrimaryCak)
				gatewayMacsecConfigTemplatePrimaryCakModel.Crn = core.StringPtr("crn:v1:bluemix:public:hs-crypto:us-south:a/4111d05f36894e3cb9b46a43556d9000:abc111b8-37aa-4034-9def-f2607c87aaaa:key:bbb222bc-430a-4de9-9aad-84e5bb022222")

				// Construct an instance of the GatewayMacsecConfigTemplate model
				gatewayMacsecConfigTemplateModel := new(directlinkv1.GatewayMacsecConfigTemplate)
				gatewayMacsecConfigTemplateModel.Active = core.BoolPtr(true)
				gatewayMacsecConfigTemplateModel.FallbackCak = gatewayMacsecConfigTemplateFallbackCakModel
				gatewayMacsecConfigTemplateModel.PrimaryCak = gatewayMacsecConfigTemplatePrimaryCakModel
				gatewayMacsecConfigTemplateModel.WindowSize = core.Int64Ptr(int64(148809600))

				// Construct an instance of the ResourceGroupIdentity model
				resourceGroupIdentityModel := new(directlinkv1.ResourceGroupIdentity)
				resourceGroupIdentityModel.ID = core.StringPtr("56969d6043e9465c883cb9f7363e78e8")

				// Construct an instance of the GatewayTemplateGatewayTypeDedicatedTemplate model
				gatewayTemplateModel := new(directlinkv1.GatewayTemplateGatewayTypeDedicatedTemplate)
				gatewayTemplateModel.BgpAsn = core.Int64Ptr(int64(64999))
				gatewayTemplateModel.BgpBaseCidr = core.StringPtr("testString")
				gatewayTemplateModel.BgpCerCidr = core.StringPtr("169.254.0.10/30")
				gatewayTemplateModel.BgpIbmCidr = core.StringPtr("169.254.0.9/30")
				gatewayTemplateModel.Global = core.BoolPtr(true)
				gatewayTemplateModel.Metered = core.BoolPtr(false)
				gatewayTemplateModel.Name = core.StringPtr("myGateway")
				gatewayTemplateModel.ResourceGroup = resourceGroupIdentityModel
				gatewayTemplateModel.SpeedMbps = core.Int64Ptr(int64(1000))
				gatewayTemplateModel.Type = core.StringPtr("dedicated")
				gatewayTemplateModel.CarrierName = core.StringPtr("myCarrierName")
				gatewayTemplateModel.CrossConnectRouter = core.StringPtr("xcr01.dal03")
				gatewayTemplateModel.CustomerName = core.StringPtr("newCustomerName")
				gatewayTemplateModel.LocationName = core.StringPtr("dal03")
				gatewayTemplateModel.MacsecConfig = gatewayMacsecConfigTemplateModel

				// Construct an instance of the CreateGatewayOptions model
				createGatewayOptionsModel := new(directlinkv1.CreateGatewayOptions)
				createGatewayOptionsModel.GatewayTemplate = gatewayTemplateModel
				createGatewayOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.CreateGateway(createGatewayOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateGatewayOptions model with no property values
				createGatewayOptionsModelNew := new(directlinkv1.CreateGatewayOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.CreateGateway(createGatewayOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`DeleteGateway(deleteGatewayOptions *DeleteGatewayOptions)`, func() {
		version := "testString"
		deleteGatewayPath := "/gateways/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(deleteGatewayPath))
					Expect(req.Method).To(Equal("DELETE"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteGateway successfully`, func() {
				testService, testServiceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := testService.DeleteGateway(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteGatewayOptions model
				deleteGatewayOptionsModel := new(directlinkv1.DeleteGatewayOptions)
				deleteGatewayOptionsModel.ID = core.StringPtr("testString")
				deleteGatewayOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = testService.DeleteGateway(deleteGatewayOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteGateway with error: Operation validation and request error`, func() {
				testService, testServiceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the DeleteGatewayOptions model
				deleteGatewayOptionsModel := new(directlinkv1.DeleteGatewayOptions)
				deleteGatewayOptionsModel.ID = core.StringPtr("testString")
				deleteGatewayOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := testService.DeleteGateway(deleteGatewayOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteGatewayOptions model with no property values
				deleteGatewayOptionsModelNew := new(directlinkv1.DeleteGatewayOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = testService.DeleteGateway(deleteGatewayOptionsModelNew)
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
		getGatewayPath := "/gateways/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getGatewayPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetGateway with error: Operation response processing error`, func() {
				testService, testServiceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetGatewayOptions model
				getGatewayOptionsModel := new(directlinkv1.GetGatewayOptions)
				getGatewayOptionsModel.ID = core.StringPtr("testString")
				getGatewayOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.GetGateway(getGatewayOptionsModel)
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
		getGatewayPath := "/gateways/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getGatewayPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"bgp_asn": 64999, "bgp_base_cidr": "BgpBaseCidr", "bgp_cer_cidr": "10.254.30.78/30", "bgp_ibm_asn": 13884, "bgp_ibm_cidr": "10.254.30.77/30", "bgp_status": "active", "carrier_name": "myCarrierName", "change_request": {"type": "create_gateway"}, "completion_notice_reject_reason": "The completion notice file was blank", "created_at": "2019-01-01T12:00:00", "crn": "crn:v1:bluemix:public:directlink:dal03:a/4111d05f36894e3cb9b46a43556d9000::dedicated:ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4", "cross_connect_router": "xcr01.dal03", "customer_name": "newCustomerName", "global": true, "id": "ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4", "link_status": "up", "location_display_name": "Dallas 03", "location_name": "dal03", "macsec_config": {"active": true, "active_cak": {"crn": "crn:v1:bluemix:public:hs-crypto:us-south:a/4111d05f36894e3cb9b46a43556d9000:abc111b8-37aa-4034-9def-f2607c87aaaa:key:bbb222bc-430a-4de9-9aad-84e5bb022222", "status": "Status"}, "cipher_suite": "gcm_aes_xpn_256", "confidentiality_offset": 0, "cryptographic_algorithm": "aes_256_cmac", "fallback_cak": {"crn": "crn:v1:bluemix:public:hs-crypto:us-south:a/4111d05f36894e3cb9b46a43556d9000:abc111b8-37aa-4034-9def-f2607c87aaaa:key:bbb222bc-430a-4de9-9aad-84e5bb022222", "status": "Status"}, "key_server_priority": 255, "primary_cak": {"crn": "crn:v1:bluemix:public:hs-crypto:us-south:a/4111d05f36894e3cb9b46a43556d9000:abc111b8-37aa-4034-9def-f2607c87aaaa:key:bbb222bc-430a-4de9-9aad-84e5bb022222", "status": "Status"}, "sak_expiry_time": 3600, "security_policy": "should_secure", "status": "secured", "window_size": 64}, "metered": false, "name": "myGateway", "operational_status": "awaiting_completion_notice", "port": {"id": "54321b1a-fee4-41c7-9e11-9cd99e000aaa"}, "provider_api_managed": false, "resource_group": {"id": "56969d6043e9465c883cb9f7363e78e8"}, "speed_mbps": 1000, "type": "dedicated", "vlan": 10}`)
				}))
			})
			It(`Invoke GetGateway successfully`, func() {
				testService, testServiceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetGateway(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetGatewayOptions model
				getGatewayOptionsModel := new(directlinkv1.GetGatewayOptions)
				getGatewayOptionsModel.ID = core.StringPtr("testString")
				getGatewayOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetGateway(getGatewayOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke GetGateway with error: Operation validation and request error`, func() {
				testService, testServiceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetGatewayOptions model
				getGatewayOptionsModel := new(directlinkv1.GetGatewayOptions)
				getGatewayOptionsModel.ID = core.StringPtr("testString")
				getGatewayOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.GetGateway(getGatewayOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetGatewayOptions model with no property values
				getGatewayOptionsModelNew := new(directlinkv1.GetGatewayOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.GetGateway(getGatewayOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateGateway(updateGatewayOptions *UpdateGatewayOptions) - Operation response error`, func() {
		version := "testString"
		updateGatewayPath := "/gateways/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(updateGatewayPath))
					Expect(req.Method).To(Equal("PATCH"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateGateway with error: Operation response processing error`, func() {
				testService, testServiceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GatewayMacsecConfigPatchTemplateFallbackCak model
				gatewayMacsecConfigPatchTemplateFallbackCakModel := new(directlinkv1.GatewayMacsecConfigPatchTemplateFallbackCak)
				gatewayMacsecConfigPatchTemplateFallbackCakModel.Crn = core.StringPtr("crn:v1:bluemix:public:hs-crypto:us-south:a/4111d05f36894e3cb9b46a43556d9000:abc111b8-37aa-4034-9def-f2607c87aaaa:key:bbb222bc-430a-4de9-9aad-84e5bb022222")

				// Construct an instance of the GatewayMacsecConfigPatchTemplatePrimaryCak model
				gatewayMacsecConfigPatchTemplatePrimaryCakModel := new(directlinkv1.GatewayMacsecConfigPatchTemplatePrimaryCak)
				gatewayMacsecConfigPatchTemplatePrimaryCakModel.Crn = core.StringPtr("crn:v1:bluemix:public:hs-crypto:us-south:a/4111d05f36894e3cb9b46a43556d9000:abc111b8-37aa-4034-9def-f2607c87aaaa:key:bbb222bc-430a-4de9-9aad-84e5bb022222")

				// Construct an instance of the GatewayMacsecConfigPatchTemplate model
				gatewayMacsecConfigPatchTemplateModel := new(directlinkv1.GatewayMacsecConfigPatchTemplate)
				gatewayMacsecConfigPatchTemplateModel.Active = core.BoolPtr(true)
				gatewayMacsecConfigPatchTemplateModel.FallbackCak = gatewayMacsecConfigPatchTemplateFallbackCakModel
				gatewayMacsecConfigPatchTemplateModel.PrimaryCak = gatewayMacsecConfigPatchTemplatePrimaryCakModel
				gatewayMacsecConfigPatchTemplateModel.WindowSize = core.Int64Ptr(int64(512))

				// Construct an instance of the UpdateGatewayOptions model
				updateGatewayOptionsModel := new(directlinkv1.UpdateGatewayOptions)
				updateGatewayOptionsModel.ID = core.StringPtr("testString")
				updateGatewayOptionsModel.Global = core.BoolPtr(true)
				updateGatewayOptionsModel.LoaRejectReason = core.StringPtr("The port mentioned was incorrect")
				updateGatewayOptionsModel.MacsecConfig = gatewayMacsecConfigPatchTemplateModel
				updateGatewayOptionsModel.Metered = core.BoolPtr(false)
				updateGatewayOptionsModel.Name = core.StringPtr("testGateway")
				updateGatewayOptionsModel.OperationalStatus = core.StringPtr("loa_accepted")
				updateGatewayOptionsModel.SpeedMbps = core.Int64Ptr(int64(1000))
				updateGatewayOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.UpdateGateway(updateGatewayOptionsModel)
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
		updateGatewayPath := "/gateways/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(updateGatewayPath))
					Expect(req.Method).To(Equal("PATCH"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"bgp_asn": 64999, "bgp_base_cidr": "BgpBaseCidr", "bgp_cer_cidr": "10.254.30.78/30", "bgp_ibm_asn": 13884, "bgp_ibm_cidr": "10.254.30.77/30", "bgp_status": "active", "carrier_name": "myCarrierName", "change_request": {"type": "create_gateway"}, "completion_notice_reject_reason": "The completion notice file was blank", "created_at": "2019-01-01T12:00:00", "crn": "crn:v1:bluemix:public:directlink:dal03:a/4111d05f36894e3cb9b46a43556d9000::dedicated:ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4", "cross_connect_router": "xcr01.dal03", "customer_name": "newCustomerName", "global": true, "id": "ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4", "link_status": "up", "location_display_name": "Dallas 03", "location_name": "dal03", "macsec_config": {"active": true, "active_cak": {"crn": "crn:v1:bluemix:public:hs-crypto:us-south:a/4111d05f36894e3cb9b46a43556d9000:abc111b8-37aa-4034-9def-f2607c87aaaa:key:bbb222bc-430a-4de9-9aad-84e5bb022222", "status": "Status"}, "cipher_suite": "gcm_aes_xpn_256", "confidentiality_offset": 0, "cryptographic_algorithm": "aes_256_cmac", "fallback_cak": {"crn": "crn:v1:bluemix:public:hs-crypto:us-south:a/4111d05f36894e3cb9b46a43556d9000:abc111b8-37aa-4034-9def-f2607c87aaaa:key:bbb222bc-430a-4de9-9aad-84e5bb022222", "status": "Status"}, "key_server_priority": 255, "primary_cak": {"crn": "crn:v1:bluemix:public:hs-crypto:us-south:a/4111d05f36894e3cb9b46a43556d9000:abc111b8-37aa-4034-9def-f2607c87aaaa:key:bbb222bc-430a-4de9-9aad-84e5bb022222", "status": "Status"}, "sak_expiry_time": 3600, "security_policy": "should_secure", "status": "secured", "window_size": 64}, "metered": false, "name": "myGateway", "operational_status": "awaiting_completion_notice", "port": {"id": "54321b1a-fee4-41c7-9e11-9cd99e000aaa"}, "provider_api_managed": false, "resource_group": {"id": "56969d6043e9465c883cb9f7363e78e8"}, "speed_mbps": 1000, "type": "dedicated", "vlan": 10}`)
				}))
			})
			It(`Invoke UpdateGateway successfully`, func() {
				testService, testServiceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.UpdateGateway(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GatewayMacsecConfigPatchTemplateFallbackCak model
				gatewayMacsecConfigPatchTemplateFallbackCakModel := new(directlinkv1.GatewayMacsecConfigPatchTemplateFallbackCak)
				gatewayMacsecConfigPatchTemplateFallbackCakModel.Crn = core.StringPtr("crn:v1:bluemix:public:hs-crypto:us-south:a/4111d05f36894e3cb9b46a43556d9000:abc111b8-37aa-4034-9def-f2607c87aaaa:key:bbb222bc-430a-4de9-9aad-84e5bb022222")

				// Construct an instance of the GatewayMacsecConfigPatchTemplatePrimaryCak model
				gatewayMacsecConfigPatchTemplatePrimaryCakModel := new(directlinkv1.GatewayMacsecConfigPatchTemplatePrimaryCak)
				gatewayMacsecConfigPatchTemplatePrimaryCakModel.Crn = core.StringPtr("crn:v1:bluemix:public:hs-crypto:us-south:a/4111d05f36894e3cb9b46a43556d9000:abc111b8-37aa-4034-9def-f2607c87aaaa:key:bbb222bc-430a-4de9-9aad-84e5bb022222")

				// Construct an instance of the GatewayMacsecConfigPatchTemplate model
				gatewayMacsecConfigPatchTemplateModel := new(directlinkv1.GatewayMacsecConfigPatchTemplate)
				gatewayMacsecConfigPatchTemplateModel.Active = core.BoolPtr(true)
				gatewayMacsecConfigPatchTemplateModel.FallbackCak = gatewayMacsecConfigPatchTemplateFallbackCakModel
				gatewayMacsecConfigPatchTemplateModel.PrimaryCak = gatewayMacsecConfigPatchTemplatePrimaryCakModel
				gatewayMacsecConfigPatchTemplateModel.WindowSize = core.Int64Ptr(int64(512))

				// Construct an instance of the UpdateGatewayOptions model
				updateGatewayOptionsModel := new(directlinkv1.UpdateGatewayOptions)
				updateGatewayOptionsModel.ID = core.StringPtr("testString")
				updateGatewayOptionsModel.Global = core.BoolPtr(true)
				updateGatewayOptionsModel.LoaRejectReason = core.StringPtr("The port mentioned was incorrect")
				updateGatewayOptionsModel.MacsecConfig = gatewayMacsecConfigPatchTemplateModel
				updateGatewayOptionsModel.Metered = core.BoolPtr(false)
				updateGatewayOptionsModel.Name = core.StringPtr("testGateway")
				updateGatewayOptionsModel.OperationalStatus = core.StringPtr("loa_accepted")
				updateGatewayOptionsModel.SpeedMbps = core.Int64Ptr(int64(1000))
				updateGatewayOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.UpdateGateway(updateGatewayOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke UpdateGateway with error: Operation validation and request error`, func() {
				testService, testServiceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GatewayMacsecConfigPatchTemplateFallbackCak model
				gatewayMacsecConfigPatchTemplateFallbackCakModel := new(directlinkv1.GatewayMacsecConfigPatchTemplateFallbackCak)
				gatewayMacsecConfigPatchTemplateFallbackCakModel.Crn = core.StringPtr("crn:v1:bluemix:public:hs-crypto:us-south:a/4111d05f36894e3cb9b46a43556d9000:abc111b8-37aa-4034-9def-f2607c87aaaa:key:bbb222bc-430a-4de9-9aad-84e5bb022222")

				// Construct an instance of the GatewayMacsecConfigPatchTemplatePrimaryCak model
				gatewayMacsecConfigPatchTemplatePrimaryCakModel := new(directlinkv1.GatewayMacsecConfigPatchTemplatePrimaryCak)
				gatewayMacsecConfigPatchTemplatePrimaryCakModel.Crn = core.StringPtr("crn:v1:bluemix:public:hs-crypto:us-south:a/4111d05f36894e3cb9b46a43556d9000:abc111b8-37aa-4034-9def-f2607c87aaaa:key:bbb222bc-430a-4de9-9aad-84e5bb022222")

				// Construct an instance of the GatewayMacsecConfigPatchTemplate model
				gatewayMacsecConfigPatchTemplateModel := new(directlinkv1.GatewayMacsecConfigPatchTemplate)
				gatewayMacsecConfigPatchTemplateModel.Active = core.BoolPtr(true)
				gatewayMacsecConfigPatchTemplateModel.FallbackCak = gatewayMacsecConfigPatchTemplateFallbackCakModel
				gatewayMacsecConfigPatchTemplateModel.PrimaryCak = gatewayMacsecConfigPatchTemplatePrimaryCakModel
				gatewayMacsecConfigPatchTemplateModel.WindowSize = core.Int64Ptr(int64(512))

				// Construct an instance of the UpdateGatewayOptions model
				updateGatewayOptionsModel := new(directlinkv1.UpdateGatewayOptions)
				updateGatewayOptionsModel.ID = core.StringPtr("testString")
				updateGatewayOptionsModel.Global = core.BoolPtr(true)
				updateGatewayOptionsModel.LoaRejectReason = core.StringPtr("The port mentioned was incorrect")
				updateGatewayOptionsModel.MacsecConfig = gatewayMacsecConfigPatchTemplateModel
				updateGatewayOptionsModel.Metered = core.BoolPtr(false)
				updateGatewayOptionsModel.Name = core.StringPtr("testGateway")
				updateGatewayOptionsModel.OperationalStatus = core.StringPtr("loa_accepted")
				updateGatewayOptionsModel.SpeedMbps = core.Int64Ptr(int64(1000))
				updateGatewayOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.UpdateGateway(updateGatewayOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateGatewayOptions model with no property values
				updateGatewayOptionsModelNew := new(directlinkv1.UpdateGatewayOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.UpdateGateway(updateGatewayOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateGatewayAction(createGatewayActionOptions *CreateGatewayActionOptions) - Operation response error`, func() {
		version := "testString"
		createGatewayActionPath := "/gateways/testString/actions"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(createGatewayActionPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateGatewayAction with error: Operation response processing error`, func() {
				testService, testServiceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GatewayActionTemplateUpdatesItemGatewayClientSpeedUpdate model
				gatewayActionTemplateUpdatesItemModel := new(directlinkv1.GatewayActionTemplateUpdatesItemGatewayClientSpeedUpdate)
				gatewayActionTemplateUpdatesItemModel.SpeedMbps = core.Int64Ptr(int64(500))

				// Construct an instance of the ResourceGroupIdentity model
				resourceGroupIdentityModel := new(directlinkv1.ResourceGroupIdentity)
				resourceGroupIdentityModel.ID = core.StringPtr("56969d6043e9465c883cb9f7363e78e8")

				// Construct an instance of the CreateGatewayActionOptions model
				createGatewayActionOptionsModel := new(directlinkv1.CreateGatewayActionOptions)
				createGatewayActionOptionsModel.ID = core.StringPtr("testString")
				createGatewayActionOptionsModel.Action = core.StringPtr("create_gateway_approve")
				createGatewayActionOptionsModel.Global = core.BoolPtr(true)
				createGatewayActionOptionsModel.Metered = core.BoolPtr(false)
				createGatewayActionOptionsModel.ResourceGroup = resourceGroupIdentityModel
				createGatewayActionOptionsModel.Updates = []directlinkv1.GatewayActionTemplateUpdatesItemIntf{gatewayActionTemplateUpdatesItemModel}
				createGatewayActionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.CreateGatewayAction(createGatewayActionOptionsModel)
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
		createGatewayActionPath := "/gateways/testString/actions"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(createGatewayActionPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"bgp_asn": 64999, "bgp_base_cidr": "BgpBaseCidr", "bgp_cer_cidr": "10.254.30.78/30", "bgp_ibm_asn": 13884, "bgp_ibm_cidr": "10.254.30.77/30", "bgp_status": "active", "carrier_name": "myCarrierName", "change_request": {"type": "create_gateway"}, "completion_notice_reject_reason": "The completion notice file was blank", "created_at": "2019-01-01T12:00:00", "crn": "crn:v1:bluemix:public:directlink:dal03:a/4111d05f36894e3cb9b46a43556d9000::dedicated:ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4", "cross_connect_router": "xcr01.dal03", "customer_name": "newCustomerName", "global": true, "id": "ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4", "link_status": "up", "location_display_name": "Dallas 03", "location_name": "dal03", "macsec_config": {"active": true, "active_cak": {"crn": "crn:v1:bluemix:public:hs-crypto:us-south:a/4111d05f36894e3cb9b46a43556d9000:abc111b8-37aa-4034-9def-f2607c87aaaa:key:bbb222bc-430a-4de9-9aad-84e5bb022222", "status": "Status"}, "cipher_suite": "gcm_aes_xpn_256", "confidentiality_offset": 0, "cryptographic_algorithm": "aes_256_cmac", "fallback_cak": {"crn": "crn:v1:bluemix:public:hs-crypto:us-south:a/4111d05f36894e3cb9b46a43556d9000:abc111b8-37aa-4034-9def-f2607c87aaaa:key:bbb222bc-430a-4de9-9aad-84e5bb022222", "status": "Status"}, "key_server_priority": 255, "primary_cak": {"crn": "crn:v1:bluemix:public:hs-crypto:us-south:a/4111d05f36894e3cb9b46a43556d9000:abc111b8-37aa-4034-9def-f2607c87aaaa:key:bbb222bc-430a-4de9-9aad-84e5bb022222", "status": "Status"}, "sak_expiry_time": 3600, "security_policy": "should_secure", "status": "secured", "window_size": 64}, "metered": false, "name": "myGateway", "operational_status": "awaiting_completion_notice", "port": {"id": "54321b1a-fee4-41c7-9e11-9cd99e000aaa"}, "provider_api_managed": false, "resource_group": {"id": "56969d6043e9465c883cb9f7363e78e8"}, "speed_mbps": 1000, "type": "dedicated", "vlan": 10}`)
				}))
			})
			It(`Invoke CreateGatewayAction successfully`, func() {
				testService, testServiceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.CreateGatewayAction(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GatewayActionTemplateUpdatesItemGatewayClientSpeedUpdate model
				gatewayActionTemplateUpdatesItemModel := new(directlinkv1.GatewayActionTemplateUpdatesItemGatewayClientSpeedUpdate)
				gatewayActionTemplateUpdatesItemModel.SpeedMbps = core.Int64Ptr(int64(500))

				// Construct an instance of the ResourceGroupIdentity model
				resourceGroupIdentityModel := new(directlinkv1.ResourceGroupIdentity)
				resourceGroupIdentityModel.ID = core.StringPtr("56969d6043e9465c883cb9f7363e78e8")

				// Construct an instance of the CreateGatewayActionOptions model
				createGatewayActionOptionsModel := new(directlinkv1.CreateGatewayActionOptions)
				createGatewayActionOptionsModel.ID = core.StringPtr("testString")
				createGatewayActionOptionsModel.Action = core.StringPtr("create_gateway_approve")
				createGatewayActionOptionsModel.Global = core.BoolPtr(true)
				createGatewayActionOptionsModel.Metered = core.BoolPtr(false)
				createGatewayActionOptionsModel.ResourceGroup = resourceGroupIdentityModel
				createGatewayActionOptionsModel.Updates = []directlinkv1.GatewayActionTemplateUpdatesItemIntf{gatewayActionTemplateUpdatesItemModel}
				createGatewayActionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.CreateGatewayAction(createGatewayActionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke CreateGatewayAction with error: Operation validation and request error`, func() {
				testService, testServiceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GatewayActionTemplateUpdatesItemGatewayClientSpeedUpdate model
				gatewayActionTemplateUpdatesItemModel := new(directlinkv1.GatewayActionTemplateUpdatesItemGatewayClientSpeedUpdate)
				gatewayActionTemplateUpdatesItemModel.SpeedMbps = core.Int64Ptr(int64(500))

				// Construct an instance of the ResourceGroupIdentity model
				resourceGroupIdentityModel := new(directlinkv1.ResourceGroupIdentity)
				resourceGroupIdentityModel.ID = core.StringPtr("56969d6043e9465c883cb9f7363e78e8")

				// Construct an instance of the CreateGatewayActionOptions model
				createGatewayActionOptionsModel := new(directlinkv1.CreateGatewayActionOptions)
				createGatewayActionOptionsModel.ID = core.StringPtr("testString")
				createGatewayActionOptionsModel.Action = core.StringPtr("create_gateway_approve")
				createGatewayActionOptionsModel.Global = core.BoolPtr(true)
				createGatewayActionOptionsModel.Metered = core.BoolPtr(false)
				createGatewayActionOptionsModel.ResourceGroup = resourceGroupIdentityModel
				createGatewayActionOptionsModel.Updates = []directlinkv1.GatewayActionTemplateUpdatesItemIntf{gatewayActionTemplateUpdatesItemModel}
				createGatewayActionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.CreateGatewayAction(createGatewayActionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateGatewayActionOptions model with no property values
				createGatewayActionOptionsModelNew := new(directlinkv1.CreateGatewayActionOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.CreateGatewayAction(createGatewayActionOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ListGatewayCompletionNotice(listGatewayCompletionNoticeOptions *ListGatewayCompletionNoticeOptions)`, func() {
		version := "testString"
		listGatewayCompletionNoticePath := "/gateways/testString/completion_notice"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(listGatewayCompletionNoticePath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/pdf")
					res.WriteHeader(200)
					fmt.Fprintf(res, `Contents of response byte-stream...`)
				}))
			})
			It(`Invoke ListGatewayCompletionNotice successfully`, func() {
				testService, testServiceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.ListGatewayCompletionNotice(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListGatewayCompletionNoticeOptions model
				listGatewayCompletionNoticeOptionsModel := new(directlinkv1.ListGatewayCompletionNoticeOptions)
				listGatewayCompletionNoticeOptionsModel.ID = core.StringPtr("testString")
				listGatewayCompletionNoticeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.ListGatewayCompletionNotice(listGatewayCompletionNoticeOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke ListGatewayCompletionNotice with error: Operation validation and request error`, func() {
				testService, testServiceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ListGatewayCompletionNoticeOptions model
				listGatewayCompletionNoticeOptionsModel := new(directlinkv1.ListGatewayCompletionNoticeOptions)
				listGatewayCompletionNoticeOptionsModel.ID = core.StringPtr("testString")
				listGatewayCompletionNoticeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.ListGatewayCompletionNotice(listGatewayCompletionNoticeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListGatewayCompletionNoticeOptions model with no property values
				listGatewayCompletionNoticeOptionsModelNew := new(directlinkv1.ListGatewayCompletionNoticeOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.ListGatewayCompletionNotice(listGatewayCompletionNoticeOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`CreateGatewayCompletionNotice(createGatewayCompletionNoticeOptions *CreateGatewayCompletionNoticeOptions)`, func() {
		version := "testString"
		createGatewayCompletionNoticePath := "/gateways/testString/completion_notice"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(createGatewayCompletionNoticePath))
					Expect(req.Method).To(Equal("PUT"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					res.WriteHeader(204)
				}))
			})
			It(`Invoke CreateGatewayCompletionNotice successfully`, func() {
				testService, testServiceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := testService.CreateGatewayCompletionNotice(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the CreateGatewayCompletionNoticeOptions model
				createGatewayCompletionNoticeOptionsModel := new(directlinkv1.CreateGatewayCompletionNoticeOptions)
				createGatewayCompletionNoticeOptionsModel.ID = core.StringPtr("testString")
				createGatewayCompletionNoticeOptionsModel.Upload = CreateMockReader("This is a mock file.")
				createGatewayCompletionNoticeOptionsModel.UploadContentType = core.StringPtr("testString")
				createGatewayCompletionNoticeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = testService.CreateGatewayCompletionNotice(createGatewayCompletionNoticeOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke CreateGatewayCompletionNotice with error: Param validation error`, func() {
				testService, testServiceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the CreateGatewayCompletionNoticeOptions model
				createGatewayCompletionNoticeOptionsModel := new(directlinkv1.CreateGatewayCompletionNoticeOptions)
				// Invoke operation with invalid options model (negative test)
				response, operationErr := testService.CreateGatewayCompletionNotice(createGatewayCompletionNoticeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			It(`Invoke CreateGatewayCompletionNotice with error: Operation validation and request error`, func() {
				testService, testServiceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the CreateGatewayCompletionNoticeOptions model
				createGatewayCompletionNoticeOptionsModel := new(directlinkv1.CreateGatewayCompletionNoticeOptions)
				createGatewayCompletionNoticeOptionsModel.ID = core.StringPtr("testString")
				createGatewayCompletionNoticeOptionsModel.Upload = CreateMockReader("This is a mock file.")
				createGatewayCompletionNoticeOptionsModel.UploadContentType = core.StringPtr("testString")
				createGatewayCompletionNoticeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := testService.CreateGatewayCompletionNotice(createGatewayCompletionNoticeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the CreateGatewayCompletionNoticeOptions model with no property values
				createGatewayCompletionNoticeOptionsModelNew := new(directlinkv1.CreateGatewayCompletionNoticeOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = testService.CreateGatewayCompletionNotice(createGatewayCompletionNoticeOptionsModelNew)
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
		listGatewayLetterOfAuthorizationPath := "/gateways/testString/letter_of_authorization"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(listGatewayLetterOfAuthorizationPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/pdf")
					res.WriteHeader(200)
					fmt.Fprintf(res, `Contents of response byte-stream...`)
				}))
			})
			It(`Invoke ListGatewayLetterOfAuthorization successfully`, func() {
				testService, testServiceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.ListGatewayLetterOfAuthorization(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListGatewayLetterOfAuthorizationOptions model
				listGatewayLetterOfAuthorizationOptionsModel := new(directlinkv1.ListGatewayLetterOfAuthorizationOptions)
				listGatewayLetterOfAuthorizationOptionsModel.ID = core.StringPtr("testString")
				listGatewayLetterOfAuthorizationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.ListGatewayLetterOfAuthorization(listGatewayLetterOfAuthorizationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke ListGatewayLetterOfAuthorization with error: Operation validation and request error`, func() {
				testService, testServiceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ListGatewayLetterOfAuthorizationOptions model
				listGatewayLetterOfAuthorizationOptionsModel := new(directlinkv1.ListGatewayLetterOfAuthorizationOptions)
				listGatewayLetterOfAuthorizationOptionsModel.ID = core.StringPtr("testString")
				listGatewayLetterOfAuthorizationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.ListGatewayLetterOfAuthorization(listGatewayLetterOfAuthorizationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListGatewayLetterOfAuthorizationOptions model with no property values
				listGatewayLetterOfAuthorizationOptionsModelNew := new(directlinkv1.ListGatewayLetterOfAuthorizationOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.ListGatewayLetterOfAuthorization(listGatewayLetterOfAuthorizationOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetGatewayStatistics(getGatewayStatisticsOptions *GetGatewayStatisticsOptions) - Operation response error`, func() {
		version := "testString"
		getGatewayStatisticsPath := "/gateways/testString/statistics"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getGatewayStatisticsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["type"]).To(Equal([]string{"macsec_mka"}))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetGatewayStatistics with error: Operation response processing error`, func() {
				testService, testServiceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetGatewayStatisticsOptions model
				getGatewayStatisticsOptionsModel := new(directlinkv1.GetGatewayStatisticsOptions)
				getGatewayStatisticsOptionsModel.ID = core.StringPtr("testString")
				getGatewayStatisticsOptionsModel.Type = core.StringPtr("macsec_mka")
				getGatewayStatisticsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.GetGatewayStatistics(getGatewayStatisticsOptionsModel)
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
		getGatewayStatisticsPath := "/gateways/testString/statistics"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getGatewayStatisticsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["type"]).To(Equal([]string{"macsec_mka"}))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"statistics": [{"created_at": "2019-01-01T12:00:00", "data": "MKA statistics text...", "type": "macsec_policy"}]}`)
				}))
			})
			It(`Invoke GetGatewayStatistics successfully`, func() {
				testService, testServiceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetGatewayStatistics(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetGatewayStatisticsOptions model
				getGatewayStatisticsOptionsModel := new(directlinkv1.GetGatewayStatisticsOptions)
				getGatewayStatisticsOptionsModel.ID = core.StringPtr("testString")
				getGatewayStatisticsOptionsModel.Type = core.StringPtr("macsec_mka")
				getGatewayStatisticsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetGatewayStatistics(getGatewayStatisticsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke GetGatewayStatistics with error: Operation validation and request error`, func() {
				testService, testServiceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetGatewayStatisticsOptions model
				getGatewayStatisticsOptionsModel := new(directlinkv1.GetGatewayStatisticsOptions)
				getGatewayStatisticsOptionsModel.ID = core.StringPtr("testString")
				getGatewayStatisticsOptionsModel.Type = core.StringPtr("macsec_mka")
				getGatewayStatisticsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.GetGatewayStatistics(getGatewayStatisticsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetGatewayStatisticsOptions model with no property values
				getGatewayStatisticsOptionsModelNew := new(directlinkv1.GetGatewayStatisticsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.GetGatewayStatistics(getGatewayStatisticsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Service constructor tests`, func() {
		version := "testString"
		It(`Instantiate service client`, func() {
			testService, testServiceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
				Version:       core.StringPtr(version),
			})
			Expect(testService).ToNot(BeNil())
			Expect(testServiceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			testService, testServiceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
				URL:     "{BAD_URL_STRING",
				Version: core.StringPtr(version),
			})
			Expect(testService).To(BeNil())
			Expect(testServiceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			testService, testServiceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
				URL:     "https://directlinkv1/api",
				Version: core.StringPtr(version),
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(testService).To(BeNil())
			Expect(testServiceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Validation Error`, func() {
			testService, testServiceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{})
			Expect(testService).To(BeNil())
			Expect(testServiceErr).ToNot(BeNil())
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
				testService, testServiceErr := directlinkv1.NewDirectLinkV1UsingExternalConfig(&directlinkv1.DirectLinkV1Options{
					Version: core.StringPtr(version),
				})
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := directlinkv1.NewDirectLinkV1UsingExternalConfig(&directlinkv1.DirectLinkV1Options{
					URL:     "https://testService/api",
					Version: core.StringPtr(version),
				})
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				Expect(testService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := directlinkv1.NewDirectLinkV1UsingExternalConfig(&directlinkv1.DirectLinkV1Options{
					Version: core.StringPtr(version),
				})
				err := testService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				Expect(testService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"DIRECT_LINK_URL":       "https://directlinkv1/api",
				"DIRECT_LINK_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			testService, testServiceErr := directlinkv1.NewDirectLinkV1UsingExternalConfig(&directlinkv1.DirectLinkV1Options{
				Version: core.StringPtr(version),
			})

			It(`Instantiate service client with error`, func() {
				Expect(testService).To(BeNil())
				Expect(testServiceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"DIRECT_LINK_AUTH_TYPE": "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			testService, testServiceErr := directlinkv1.NewDirectLinkV1UsingExternalConfig(&directlinkv1.DirectLinkV1Options{
				URL:     "{BAD_URL_STRING",
				Version: core.StringPtr(version),
			})

			It(`Instantiate service client with error`, func() {
				Expect(testService).To(BeNil())
				Expect(testServiceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`ListOfferingTypeLocations(listOfferingTypeLocationsOptions *ListOfferingTypeLocationsOptions) - Operation response error`, func() {
		version := "testString"
		listOfferingTypeLocationsPath := "/offering_types/dedicated/locations"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(listOfferingTypeLocationsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListOfferingTypeLocations with error: Operation response processing error`, func() {
				testService, testServiceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ListOfferingTypeLocationsOptions model
				listOfferingTypeLocationsOptionsModel := new(directlinkv1.ListOfferingTypeLocationsOptions)
				listOfferingTypeLocationsOptionsModel.OfferingType = core.StringPtr("dedicated")
				listOfferingTypeLocationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.ListOfferingTypeLocations(listOfferingTypeLocationsOptionsModel)
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
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(listOfferingTypeLocationsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"locations": [{"billing_location": "us", "building_colocation_owner": "MyProvider", "display_name": "Dallas 9", "location_type": "PoP", "macsec_enabled": false, "market": "Dallas", "market_geography": "N/S America", "mzr": true, "name": "dal03", "offering_type": "dedicated", "provision_enabled": true, "vpc_region": "us-south"}]}`)
				}))
			})
			It(`Invoke ListOfferingTypeLocations successfully`, func() {
				testService, testServiceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.ListOfferingTypeLocations(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListOfferingTypeLocationsOptions model
				listOfferingTypeLocationsOptionsModel := new(directlinkv1.ListOfferingTypeLocationsOptions)
				listOfferingTypeLocationsOptionsModel.OfferingType = core.StringPtr("dedicated")
				listOfferingTypeLocationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.ListOfferingTypeLocations(listOfferingTypeLocationsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke ListOfferingTypeLocations with error: Operation validation and request error`, func() {
				testService, testServiceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ListOfferingTypeLocationsOptions model
				listOfferingTypeLocationsOptionsModel := new(directlinkv1.ListOfferingTypeLocationsOptions)
				listOfferingTypeLocationsOptionsModel.OfferingType = core.StringPtr("dedicated")
				listOfferingTypeLocationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.ListOfferingTypeLocations(listOfferingTypeLocationsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListOfferingTypeLocationsOptions model with no property values
				listOfferingTypeLocationsOptionsModelNew := new(directlinkv1.ListOfferingTypeLocationsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.ListOfferingTypeLocations(listOfferingTypeLocationsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
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
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(listOfferingTypeLocationCrossConnectRoutersPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListOfferingTypeLocationCrossConnectRouters with error: Operation response processing error`, func() {
				testService, testServiceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ListOfferingTypeLocationCrossConnectRoutersOptions model
				listOfferingTypeLocationCrossConnectRoutersOptionsModel := new(directlinkv1.ListOfferingTypeLocationCrossConnectRoutersOptions)
				listOfferingTypeLocationCrossConnectRoutersOptionsModel.OfferingType = core.StringPtr("dedicated")
				listOfferingTypeLocationCrossConnectRoutersOptionsModel.LocationName = core.StringPtr("testString")
				listOfferingTypeLocationCrossConnectRoutersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.ListOfferingTypeLocationCrossConnectRouters(listOfferingTypeLocationCrossConnectRoutersOptionsModel)
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
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(listOfferingTypeLocationCrossConnectRoutersPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"cross_connect_routers": [{"router_name": "xcr01.dal03", "total_connections": 1}]}`)
				}))
			})
			It(`Invoke ListOfferingTypeLocationCrossConnectRouters successfully`, func() {
				testService, testServiceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.ListOfferingTypeLocationCrossConnectRouters(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListOfferingTypeLocationCrossConnectRoutersOptions model
				listOfferingTypeLocationCrossConnectRoutersOptionsModel := new(directlinkv1.ListOfferingTypeLocationCrossConnectRoutersOptions)
				listOfferingTypeLocationCrossConnectRoutersOptionsModel.OfferingType = core.StringPtr("dedicated")
				listOfferingTypeLocationCrossConnectRoutersOptionsModel.LocationName = core.StringPtr("testString")
				listOfferingTypeLocationCrossConnectRoutersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.ListOfferingTypeLocationCrossConnectRouters(listOfferingTypeLocationCrossConnectRoutersOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke ListOfferingTypeLocationCrossConnectRouters with error: Operation validation and request error`, func() {
				testService, testServiceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ListOfferingTypeLocationCrossConnectRoutersOptions model
				listOfferingTypeLocationCrossConnectRoutersOptionsModel := new(directlinkv1.ListOfferingTypeLocationCrossConnectRoutersOptions)
				listOfferingTypeLocationCrossConnectRoutersOptionsModel.OfferingType = core.StringPtr("dedicated")
				listOfferingTypeLocationCrossConnectRoutersOptionsModel.LocationName = core.StringPtr("testString")
				listOfferingTypeLocationCrossConnectRoutersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.ListOfferingTypeLocationCrossConnectRouters(listOfferingTypeLocationCrossConnectRoutersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListOfferingTypeLocationCrossConnectRoutersOptions model with no property values
				listOfferingTypeLocationCrossConnectRoutersOptionsModelNew := new(directlinkv1.ListOfferingTypeLocationCrossConnectRoutersOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.ListOfferingTypeLocationCrossConnectRouters(listOfferingTypeLocationCrossConnectRoutersOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
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
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(listOfferingTypeSpeedsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListOfferingTypeSpeeds with error: Operation response processing error`, func() {
				testService, testServiceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ListOfferingTypeSpeedsOptions model
				listOfferingTypeSpeedsOptionsModel := new(directlinkv1.ListOfferingTypeSpeedsOptions)
				listOfferingTypeSpeedsOptionsModel.OfferingType = core.StringPtr("dedicated")
				listOfferingTypeSpeedsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.ListOfferingTypeSpeeds(listOfferingTypeSpeedsOptionsModel)
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
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(listOfferingTypeSpeedsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"speeds": [{"link_speed": 2000, "macsec_enabled": false}]}`)
				}))
			})
			It(`Invoke ListOfferingTypeSpeeds successfully`, func() {
				testService, testServiceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.ListOfferingTypeSpeeds(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListOfferingTypeSpeedsOptions model
				listOfferingTypeSpeedsOptionsModel := new(directlinkv1.ListOfferingTypeSpeedsOptions)
				listOfferingTypeSpeedsOptionsModel.OfferingType = core.StringPtr("dedicated")
				listOfferingTypeSpeedsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.ListOfferingTypeSpeeds(listOfferingTypeSpeedsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke ListOfferingTypeSpeeds with error: Operation validation and request error`, func() {
				testService, testServiceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ListOfferingTypeSpeedsOptions model
				listOfferingTypeSpeedsOptionsModel := new(directlinkv1.ListOfferingTypeSpeedsOptions)
				listOfferingTypeSpeedsOptionsModel.OfferingType = core.StringPtr("dedicated")
				listOfferingTypeSpeedsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.ListOfferingTypeSpeeds(listOfferingTypeSpeedsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListOfferingTypeSpeedsOptions model with no property values
				listOfferingTypeSpeedsOptionsModelNew := new(directlinkv1.ListOfferingTypeSpeedsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.ListOfferingTypeSpeeds(listOfferingTypeSpeedsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Service constructor tests`, func() {
		version := "testString"
		It(`Instantiate service client`, func() {
			testService, testServiceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
				Version:       core.StringPtr(version),
			})
			Expect(testService).ToNot(BeNil())
			Expect(testServiceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			testService, testServiceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
				URL:     "{BAD_URL_STRING",
				Version: core.StringPtr(version),
			})
			Expect(testService).To(BeNil())
			Expect(testServiceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			testService, testServiceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
				URL:     "https://directlinkv1/api",
				Version: core.StringPtr(version),
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(testService).To(BeNil())
			Expect(testServiceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Validation Error`, func() {
			testService, testServiceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{})
			Expect(testService).To(BeNil())
			Expect(testServiceErr).ToNot(BeNil())
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
				testService, testServiceErr := directlinkv1.NewDirectLinkV1UsingExternalConfig(&directlinkv1.DirectLinkV1Options{
					Version: core.StringPtr(version),
				})
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := directlinkv1.NewDirectLinkV1UsingExternalConfig(&directlinkv1.DirectLinkV1Options{
					URL:     "https://testService/api",
					Version: core.StringPtr(version),
				})
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				Expect(testService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := directlinkv1.NewDirectLinkV1UsingExternalConfig(&directlinkv1.DirectLinkV1Options{
					Version: core.StringPtr(version),
				})
				err := testService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				Expect(testService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"DIRECT_LINK_URL":       "https://directlinkv1/api",
				"DIRECT_LINK_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			testService, testServiceErr := directlinkv1.NewDirectLinkV1UsingExternalConfig(&directlinkv1.DirectLinkV1Options{
				Version: core.StringPtr(version),
			})

			It(`Instantiate service client with error`, func() {
				Expect(testService).To(BeNil())
				Expect(testServiceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"DIRECT_LINK_AUTH_TYPE": "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			testService, testServiceErr := directlinkv1.NewDirectLinkV1UsingExternalConfig(&directlinkv1.DirectLinkV1Options{
				URL:     "{BAD_URL_STRING",
				Version: core.StringPtr(version),
			})

			It(`Instantiate service client with error`, func() {
				Expect(testService).To(BeNil())
				Expect(testServiceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`ListPorts(listPortsOptions *ListPortsOptions) - Operation response error`, func() {
		version := "testString"
		listPortsPath := "/ports"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(listPortsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(38))}))

					Expect(req.URL.Query()["location_name"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListPorts with error: Operation response processing error`, func() {
				testService, testServiceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ListPortsOptions model
				listPortsOptionsModel := new(directlinkv1.ListPortsOptions)
				listPortsOptionsModel.Start = core.StringPtr("testString")
				listPortsOptionsModel.Limit = core.Int64Ptr(int64(38))
				listPortsOptionsModel.LocationName = core.StringPtr("testString")
				listPortsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.ListPorts(listPortsOptionsModel)
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
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(listPortsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(38))}))

					Expect(req.URL.Query()["location_name"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"first": {"href": "https://directlink.cloud.ibm.com/v1/ports?limit=100"}, "limit": 100, "next": {"href": "https://directlink.cloud.ibm.com/v1/ports?start=9d5a91a3e2cbd233b5a5b33436855ed1&limit=100", "start": "9d5a91a3e2cbd233b5a5b33436855ed1"}, "total_count": 132, "ports": [{"direct_link_count": 1, "id": "01122b9b-820f-4c44-8a31-77f1f0806765", "label": "XCR-FRK-CS-SEC-01", "location_display_name": "Dallas 03", "location_name": "dal03", "provider_name": "provider_1", "supported_link_speeds": [19]}]}`)
				}))
			})
			It(`Invoke ListPorts successfully`, func() {
				testService, testServiceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.ListPorts(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListPortsOptions model
				listPortsOptionsModel := new(directlinkv1.ListPortsOptions)
				listPortsOptionsModel.Start = core.StringPtr("testString")
				listPortsOptionsModel.Limit = core.Int64Ptr(int64(38))
				listPortsOptionsModel.LocationName = core.StringPtr("testString")
				listPortsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.ListPorts(listPortsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke ListPorts with error: Operation request error`, func() {
				testService, testServiceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ListPortsOptions model
				listPortsOptionsModel := new(directlinkv1.ListPortsOptions)
				listPortsOptionsModel.Start = core.StringPtr("testString")
				listPortsOptionsModel.Limit = core.Int64Ptr(int64(38))
				listPortsOptionsModel.LocationName = core.StringPtr("testString")
				listPortsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.ListPorts(listPortsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetPort(getPortOptions *GetPortOptions) - Operation response error`, func() {
		version := "testString"
		getPortPath := "/ports/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getPortPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetPort with error: Operation response processing error`, func() {
				testService, testServiceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetPortOptions model
				getPortOptionsModel := new(directlinkv1.GetPortOptions)
				getPortOptionsModel.ID = core.StringPtr("testString")
				getPortOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.GetPort(getPortOptionsModel)
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
		getPortPath := "/ports/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getPortPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"direct_link_count": 1, "id": "01122b9b-820f-4c44-8a31-77f1f0806765", "label": "XCR-FRK-CS-SEC-01", "location_display_name": "Dallas 03", "location_name": "dal03", "provider_name": "provider_1", "supported_link_speeds": [19]}`)
				}))
			})
			It(`Invoke GetPort successfully`, func() {
				testService, testServiceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetPort(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetPortOptions model
				getPortOptionsModel := new(directlinkv1.GetPortOptions)
				getPortOptionsModel.ID = core.StringPtr("testString")
				getPortOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetPort(getPortOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke GetPort with error: Operation validation and request error`, func() {
				testService, testServiceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetPortOptions model
				getPortOptionsModel := new(directlinkv1.GetPortOptions)
				getPortOptionsModel.ID = core.StringPtr("testString")
				getPortOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.GetPort(getPortOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetPortOptions model with no property values
				getPortOptionsModelNew := new(directlinkv1.GetPortOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.GetPort(getPortOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Service constructor tests`, func() {
		version := "testString"
		It(`Instantiate service client`, func() {
			testService, testServiceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
				Version:       core.StringPtr(version),
			})
			Expect(testService).ToNot(BeNil())
			Expect(testServiceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			testService, testServiceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
				URL:     "{BAD_URL_STRING",
				Version: core.StringPtr(version),
			})
			Expect(testService).To(BeNil())
			Expect(testServiceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			testService, testServiceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
				URL:     "https://directlinkv1/api",
				Version: core.StringPtr(version),
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(testService).To(BeNil())
			Expect(testServiceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Validation Error`, func() {
			testService, testServiceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{})
			Expect(testService).To(BeNil())
			Expect(testServiceErr).ToNot(BeNil())
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
				testService, testServiceErr := directlinkv1.NewDirectLinkV1UsingExternalConfig(&directlinkv1.DirectLinkV1Options{
					Version: core.StringPtr(version),
				})
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := directlinkv1.NewDirectLinkV1UsingExternalConfig(&directlinkv1.DirectLinkV1Options{
					URL:     "https://testService/api",
					Version: core.StringPtr(version),
				})
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				Expect(testService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := directlinkv1.NewDirectLinkV1UsingExternalConfig(&directlinkv1.DirectLinkV1Options{
					Version: core.StringPtr(version),
				})
				err := testService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				Expect(testService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"DIRECT_LINK_URL":       "https://directlinkv1/api",
				"DIRECT_LINK_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			testService, testServiceErr := directlinkv1.NewDirectLinkV1UsingExternalConfig(&directlinkv1.DirectLinkV1Options{
				Version: core.StringPtr(version),
			})

			It(`Instantiate service client with error`, func() {
				Expect(testService).To(BeNil())
				Expect(testServiceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"DIRECT_LINK_AUTH_TYPE": "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			testService, testServiceErr := directlinkv1.NewDirectLinkV1UsingExternalConfig(&directlinkv1.DirectLinkV1Options{
				URL:     "{BAD_URL_STRING",
				Version: core.StringPtr(version),
			})

			It(`Instantiate service client with error`, func() {
				Expect(testService).To(BeNil())
				Expect(testServiceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`ListGatewayVirtualConnections(listGatewayVirtualConnectionsOptions *ListGatewayVirtualConnectionsOptions) - Operation response error`, func() {
		version := "testString"
		listGatewayVirtualConnectionsPath := "/gateways/testString/virtual_connections"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(listGatewayVirtualConnectionsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListGatewayVirtualConnections with error: Operation response processing error`, func() {
				testService, testServiceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ListGatewayVirtualConnectionsOptions model
				listGatewayVirtualConnectionsOptionsModel := new(directlinkv1.ListGatewayVirtualConnectionsOptions)
				listGatewayVirtualConnectionsOptionsModel.GatewayID = core.StringPtr("testString")
				listGatewayVirtualConnectionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.ListGatewayVirtualConnections(listGatewayVirtualConnectionsOptionsModel)
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
		listGatewayVirtualConnectionsPath := "/gateways/testString/virtual_connections"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(listGatewayVirtualConnectionsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"virtual_connections": [{"created_at": "2019-01-01T12:00:00", "id": "ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4", "name": "newVC", "network_account": "00aa14a2e0fb102c8995ebefff865555", "network_id": "crn:v1:bluemix:public:is:us-east:a/28e4d90ac7504be69447111122223333::vpc:aaa81ac8-5e96-42a0-a4b7-6c2e2d1bbbbb", "status": "attached", "type": "vpc"}]}`)
				}))
			})
			It(`Invoke ListGatewayVirtualConnections successfully`, func() {
				testService, testServiceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.ListGatewayVirtualConnections(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListGatewayVirtualConnectionsOptions model
				listGatewayVirtualConnectionsOptionsModel := new(directlinkv1.ListGatewayVirtualConnectionsOptions)
				listGatewayVirtualConnectionsOptionsModel.GatewayID = core.StringPtr("testString")
				listGatewayVirtualConnectionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.ListGatewayVirtualConnections(listGatewayVirtualConnectionsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke ListGatewayVirtualConnections with error: Operation validation and request error`, func() {
				testService, testServiceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ListGatewayVirtualConnectionsOptions model
				listGatewayVirtualConnectionsOptionsModel := new(directlinkv1.ListGatewayVirtualConnectionsOptions)
				listGatewayVirtualConnectionsOptionsModel.GatewayID = core.StringPtr("testString")
				listGatewayVirtualConnectionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.ListGatewayVirtualConnections(listGatewayVirtualConnectionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListGatewayVirtualConnectionsOptions model with no property values
				listGatewayVirtualConnectionsOptionsModelNew := new(directlinkv1.ListGatewayVirtualConnectionsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.ListGatewayVirtualConnections(listGatewayVirtualConnectionsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateGatewayVirtualConnection(createGatewayVirtualConnectionOptions *CreateGatewayVirtualConnectionOptions) - Operation response error`, func() {
		version := "testString"
		createGatewayVirtualConnectionPath := "/gateways/testString/virtual_connections"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(createGatewayVirtualConnectionPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateGatewayVirtualConnection with error: Operation response processing error`, func() {
				testService, testServiceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the CreateGatewayVirtualConnectionOptions model
				createGatewayVirtualConnectionOptionsModel := new(directlinkv1.CreateGatewayVirtualConnectionOptions)
				createGatewayVirtualConnectionOptionsModel.GatewayID = core.StringPtr("testString")
				createGatewayVirtualConnectionOptionsModel.Name = core.StringPtr("newVC")
				createGatewayVirtualConnectionOptionsModel.Type = core.StringPtr("vpc")
				createGatewayVirtualConnectionOptionsModel.NetworkID = core.StringPtr("crn:v1:bluemix:public:is:us-east:a/28e4d90ac7504be69447111122223333::vpc:aaa81ac8-5e96-42a0-a4b7-6c2e2d1bbbbb")
				createGatewayVirtualConnectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.CreateGatewayVirtualConnection(createGatewayVirtualConnectionOptionsModel)
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
		createGatewayVirtualConnectionPath := "/gateways/testString/virtual_connections"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(createGatewayVirtualConnectionPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, `{"created_at": "2019-01-01T12:00:00", "id": "ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4", "name": "newVC", "network_account": "00aa14a2e0fb102c8995ebefff865555", "network_id": "crn:v1:bluemix:public:is:us-east:a/28e4d90ac7504be69447111122223333::vpc:aaa81ac8-5e96-42a0-a4b7-6c2e2d1bbbbb", "status": "attached", "type": "vpc"}`)
				}))
			})
			It(`Invoke CreateGatewayVirtualConnection successfully`, func() {
				testService, testServiceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.CreateGatewayVirtualConnection(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateGatewayVirtualConnectionOptions model
				createGatewayVirtualConnectionOptionsModel := new(directlinkv1.CreateGatewayVirtualConnectionOptions)
				createGatewayVirtualConnectionOptionsModel.GatewayID = core.StringPtr("testString")
				createGatewayVirtualConnectionOptionsModel.Name = core.StringPtr("newVC")
				createGatewayVirtualConnectionOptionsModel.Type = core.StringPtr("vpc")
				createGatewayVirtualConnectionOptionsModel.NetworkID = core.StringPtr("crn:v1:bluemix:public:is:us-east:a/28e4d90ac7504be69447111122223333::vpc:aaa81ac8-5e96-42a0-a4b7-6c2e2d1bbbbb")
				createGatewayVirtualConnectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.CreateGatewayVirtualConnection(createGatewayVirtualConnectionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke CreateGatewayVirtualConnection with error: Operation validation and request error`, func() {
				testService, testServiceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the CreateGatewayVirtualConnectionOptions model
				createGatewayVirtualConnectionOptionsModel := new(directlinkv1.CreateGatewayVirtualConnectionOptions)
				createGatewayVirtualConnectionOptionsModel.GatewayID = core.StringPtr("testString")
				createGatewayVirtualConnectionOptionsModel.Name = core.StringPtr("newVC")
				createGatewayVirtualConnectionOptionsModel.Type = core.StringPtr("vpc")
				createGatewayVirtualConnectionOptionsModel.NetworkID = core.StringPtr("crn:v1:bluemix:public:is:us-east:a/28e4d90ac7504be69447111122223333::vpc:aaa81ac8-5e96-42a0-a4b7-6c2e2d1bbbbb")
				createGatewayVirtualConnectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.CreateGatewayVirtualConnection(createGatewayVirtualConnectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateGatewayVirtualConnectionOptions model with no property values
				createGatewayVirtualConnectionOptionsModelNew := new(directlinkv1.CreateGatewayVirtualConnectionOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.CreateGatewayVirtualConnection(createGatewayVirtualConnectionOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`DeleteGatewayVirtualConnection(deleteGatewayVirtualConnectionOptions *DeleteGatewayVirtualConnectionOptions)`, func() {
		version := "testString"
		deleteGatewayVirtualConnectionPath := "/gateways/testString/virtual_connections/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(deleteGatewayVirtualConnectionPath))
					Expect(req.Method).To(Equal("DELETE"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteGatewayVirtualConnection successfully`, func() {
				testService, testServiceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := testService.DeleteGatewayVirtualConnection(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteGatewayVirtualConnectionOptions model
				deleteGatewayVirtualConnectionOptionsModel := new(directlinkv1.DeleteGatewayVirtualConnectionOptions)
				deleteGatewayVirtualConnectionOptionsModel.GatewayID = core.StringPtr("testString")
				deleteGatewayVirtualConnectionOptionsModel.ID = core.StringPtr("testString")
				deleteGatewayVirtualConnectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = testService.DeleteGatewayVirtualConnection(deleteGatewayVirtualConnectionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteGatewayVirtualConnection with error: Operation validation and request error`, func() {
				testService, testServiceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the DeleteGatewayVirtualConnectionOptions model
				deleteGatewayVirtualConnectionOptionsModel := new(directlinkv1.DeleteGatewayVirtualConnectionOptions)
				deleteGatewayVirtualConnectionOptionsModel.GatewayID = core.StringPtr("testString")
				deleteGatewayVirtualConnectionOptionsModel.ID = core.StringPtr("testString")
				deleteGatewayVirtualConnectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := testService.DeleteGatewayVirtualConnection(deleteGatewayVirtualConnectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteGatewayVirtualConnectionOptions model with no property values
				deleteGatewayVirtualConnectionOptionsModelNew := new(directlinkv1.DeleteGatewayVirtualConnectionOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = testService.DeleteGatewayVirtualConnection(deleteGatewayVirtualConnectionOptionsModelNew)
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
		getGatewayVirtualConnectionPath := "/gateways/testString/virtual_connections/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getGatewayVirtualConnectionPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetGatewayVirtualConnection with error: Operation response processing error`, func() {
				testService, testServiceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetGatewayVirtualConnectionOptions model
				getGatewayVirtualConnectionOptionsModel := new(directlinkv1.GetGatewayVirtualConnectionOptions)
				getGatewayVirtualConnectionOptionsModel.GatewayID = core.StringPtr("testString")
				getGatewayVirtualConnectionOptionsModel.ID = core.StringPtr("testString")
				getGatewayVirtualConnectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.GetGatewayVirtualConnection(getGatewayVirtualConnectionOptionsModel)
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
		getGatewayVirtualConnectionPath := "/gateways/testString/virtual_connections/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getGatewayVirtualConnectionPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"created_at": "2019-01-01T12:00:00", "id": "ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4", "name": "newVC", "network_account": "00aa14a2e0fb102c8995ebefff865555", "network_id": "crn:v1:bluemix:public:is:us-east:a/28e4d90ac7504be69447111122223333::vpc:aaa81ac8-5e96-42a0-a4b7-6c2e2d1bbbbb", "status": "attached", "type": "vpc"}`)
				}))
			})
			It(`Invoke GetGatewayVirtualConnection successfully`, func() {
				testService, testServiceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetGatewayVirtualConnection(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetGatewayVirtualConnectionOptions model
				getGatewayVirtualConnectionOptionsModel := new(directlinkv1.GetGatewayVirtualConnectionOptions)
				getGatewayVirtualConnectionOptionsModel.GatewayID = core.StringPtr("testString")
				getGatewayVirtualConnectionOptionsModel.ID = core.StringPtr("testString")
				getGatewayVirtualConnectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetGatewayVirtualConnection(getGatewayVirtualConnectionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke GetGatewayVirtualConnection with error: Operation validation and request error`, func() {
				testService, testServiceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetGatewayVirtualConnectionOptions model
				getGatewayVirtualConnectionOptionsModel := new(directlinkv1.GetGatewayVirtualConnectionOptions)
				getGatewayVirtualConnectionOptionsModel.GatewayID = core.StringPtr("testString")
				getGatewayVirtualConnectionOptionsModel.ID = core.StringPtr("testString")
				getGatewayVirtualConnectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.GetGatewayVirtualConnection(getGatewayVirtualConnectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetGatewayVirtualConnectionOptions model with no property values
				getGatewayVirtualConnectionOptionsModelNew := new(directlinkv1.GetGatewayVirtualConnectionOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.GetGatewayVirtualConnection(getGatewayVirtualConnectionOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateGatewayVirtualConnection(updateGatewayVirtualConnectionOptions *UpdateGatewayVirtualConnectionOptions) - Operation response error`, func() {
		version := "testString"
		updateGatewayVirtualConnectionPath := "/gateways/testString/virtual_connections/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(updateGatewayVirtualConnectionPath))
					Expect(req.Method).To(Equal("PATCH"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateGatewayVirtualConnection with error: Operation response processing error`, func() {
				testService, testServiceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the UpdateGatewayVirtualConnectionOptions model
				updateGatewayVirtualConnectionOptionsModel := new(directlinkv1.UpdateGatewayVirtualConnectionOptions)
				updateGatewayVirtualConnectionOptionsModel.GatewayID = core.StringPtr("testString")
				updateGatewayVirtualConnectionOptionsModel.ID = core.StringPtr("testString")
				updateGatewayVirtualConnectionOptionsModel.Name = core.StringPtr("newConnectionName")
				updateGatewayVirtualConnectionOptionsModel.Status = core.StringPtr("attached")
				updateGatewayVirtualConnectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.UpdateGatewayVirtualConnection(updateGatewayVirtualConnectionOptionsModel)
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
		updateGatewayVirtualConnectionPath := "/gateways/testString/virtual_connections/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(updateGatewayVirtualConnectionPath))
					Expect(req.Method).To(Equal("PATCH"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"created_at": "2019-01-01T12:00:00", "id": "ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4", "name": "newVC", "network_account": "00aa14a2e0fb102c8995ebefff865555", "network_id": "crn:v1:bluemix:public:is:us-east:a/28e4d90ac7504be69447111122223333::vpc:aaa81ac8-5e96-42a0-a4b7-6c2e2d1bbbbb", "status": "attached", "type": "vpc"}`)
				}))
			})
			It(`Invoke UpdateGatewayVirtualConnection successfully`, func() {
				testService, testServiceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.UpdateGatewayVirtualConnection(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateGatewayVirtualConnectionOptions model
				updateGatewayVirtualConnectionOptionsModel := new(directlinkv1.UpdateGatewayVirtualConnectionOptions)
				updateGatewayVirtualConnectionOptionsModel.GatewayID = core.StringPtr("testString")
				updateGatewayVirtualConnectionOptionsModel.ID = core.StringPtr("testString")
				updateGatewayVirtualConnectionOptionsModel.Name = core.StringPtr("newConnectionName")
				updateGatewayVirtualConnectionOptionsModel.Status = core.StringPtr("attached")
				updateGatewayVirtualConnectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.UpdateGatewayVirtualConnection(updateGatewayVirtualConnectionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke UpdateGatewayVirtualConnection with error: Operation validation and request error`, func() {
				testService, testServiceErr := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the UpdateGatewayVirtualConnectionOptions model
				updateGatewayVirtualConnectionOptionsModel := new(directlinkv1.UpdateGatewayVirtualConnectionOptions)
				updateGatewayVirtualConnectionOptionsModel.GatewayID = core.StringPtr("testString")
				updateGatewayVirtualConnectionOptionsModel.ID = core.StringPtr("testString")
				updateGatewayVirtualConnectionOptionsModel.Name = core.StringPtr("newConnectionName")
				updateGatewayVirtualConnectionOptionsModel.Status = core.StringPtr("attached")
				updateGatewayVirtualConnectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.UpdateGatewayVirtualConnection(updateGatewayVirtualConnectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateGatewayVirtualConnectionOptions model with no property values
				updateGatewayVirtualConnectionOptionsModelNew := new(directlinkv1.UpdateGatewayVirtualConnectionOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.UpdateGatewayVirtualConnection(updateGatewayVirtualConnectionOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
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
			testService, _ := directlinkv1.NewDirectLinkV1(&directlinkv1.DirectLinkV1Options{
				URL:           "http://directlinkv1modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
				Version:       core.StringPtr(version),
			})
			It(`Invoke NewCreateGatewayActionOptions successfully`, func() {
				// Construct an instance of the GatewayActionTemplateUpdatesItemGatewayClientSpeedUpdate model
				gatewayActionTemplateUpdatesItemModel := new(directlinkv1.GatewayActionTemplateUpdatesItemGatewayClientSpeedUpdate)
				Expect(gatewayActionTemplateUpdatesItemModel).ToNot(BeNil())
				gatewayActionTemplateUpdatesItemModel.SpeedMbps = core.Int64Ptr(int64(500))
				Expect(gatewayActionTemplateUpdatesItemModel.SpeedMbps).To(Equal(core.Int64Ptr(int64(500))))

				// Construct an instance of the ResourceGroupIdentity model
				resourceGroupIdentityModel := new(directlinkv1.ResourceGroupIdentity)
				Expect(resourceGroupIdentityModel).ToNot(BeNil())
				resourceGroupIdentityModel.ID = core.StringPtr("56969d6043e9465c883cb9f7363e78e8")
				Expect(resourceGroupIdentityModel.ID).To(Equal(core.StringPtr("56969d6043e9465c883cb9f7363e78e8")))

				// Construct an instance of the CreateGatewayActionOptions model
				id := "testString"
				createGatewayActionOptionsAction := "create_gateway_approve"
				createGatewayActionOptionsModel := testService.NewCreateGatewayActionOptions(id, createGatewayActionOptionsAction)
				createGatewayActionOptionsModel.SetID("testString")
				createGatewayActionOptionsModel.SetAction("create_gateway_approve")
				createGatewayActionOptionsModel.SetGlobal(true)
				createGatewayActionOptionsModel.SetMetered(false)
				createGatewayActionOptionsModel.SetResourceGroup(resourceGroupIdentityModel)
				createGatewayActionOptionsModel.SetUpdates([]directlinkv1.GatewayActionTemplateUpdatesItemIntf{gatewayActionTemplateUpdatesItemModel})
				createGatewayActionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createGatewayActionOptionsModel).ToNot(BeNil())
				Expect(createGatewayActionOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(createGatewayActionOptionsModel.Action).To(Equal(core.StringPtr("create_gateway_approve")))
				Expect(createGatewayActionOptionsModel.Global).To(Equal(core.BoolPtr(true)))
				Expect(createGatewayActionOptionsModel.Metered).To(Equal(core.BoolPtr(false)))
				Expect(createGatewayActionOptionsModel.ResourceGroup).To(Equal(resourceGroupIdentityModel))
				Expect(createGatewayActionOptionsModel.Updates).To(Equal([]directlinkv1.GatewayActionTemplateUpdatesItemIntf{gatewayActionTemplateUpdatesItemModel}))
				Expect(createGatewayActionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateGatewayCompletionNoticeOptions successfully`, func() {
				// Construct an instance of the CreateGatewayCompletionNoticeOptions model
				id := "testString"
				createGatewayCompletionNoticeOptionsModel := testService.NewCreateGatewayCompletionNoticeOptions(id)
				createGatewayCompletionNoticeOptionsModel.SetID("testString")
				createGatewayCompletionNoticeOptionsModel.SetUpload(CreateMockReader("This is a mock file."))
				createGatewayCompletionNoticeOptionsModel.SetUploadContentType("testString")
				createGatewayCompletionNoticeOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createGatewayCompletionNoticeOptionsModel).ToNot(BeNil())
				Expect(createGatewayCompletionNoticeOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(createGatewayCompletionNoticeOptionsModel.Upload).To(Equal(CreateMockReader("This is a mock file.")))
				Expect(createGatewayCompletionNoticeOptionsModel.UploadContentType).To(Equal(core.StringPtr("testString")))
				Expect(createGatewayCompletionNoticeOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateGatewayOptions successfully`, func() {
				// Construct an instance of the GatewayMacsecConfigTemplateFallbackCak model
				gatewayMacsecConfigTemplateFallbackCakModel := new(directlinkv1.GatewayMacsecConfigTemplateFallbackCak)
				Expect(gatewayMacsecConfigTemplateFallbackCakModel).ToNot(BeNil())
				gatewayMacsecConfigTemplateFallbackCakModel.Crn = core.StringPtr("crn:v1:bluemix:public:hs-crypto:us-south:a/4111d05f36894e3cb9b46a43556d9000:abc111b8-37aa-4034-9def-f2607c87aaaa:key:bbb222bc-430a-4de9-9aad-84e5bb022222")
				Expect(gatewayMacsecConfigTemplateFallbackCakModel.Crn).To(Equal(core.StringPtr("crn:v1:bluemix:public:hs-crypto:us-south:a/4111d05f36894e3cb9b46a43556d9000:abc111b8-37aa-4034-9def-f2607c87aaaa:key:bbb222bc-430a-4de9-9aad-84e5bb022222")))

				// Construct an instance of the GatewayMacsecConfigTemplatePrimaryCak model
				gatewayMacsecConfigTemplatePrimaryCakModel := new(directlinkv1.GatewayMacsecConfigTemplatePrimaryCak)
				Expect(gatewayMacsecConfigTemplatePrimaryCakModel).ToNot(BeNil())
				gatewayMacsecConfigTemplatePrimaryCakModel.Crn = core.StringPtr("crn:v1:bluemix:public:hs-crypto:us-south:a/4111d05f36894e3cb9b46a43556d9000:abc111b8-37aa-4034-9def-f2607c87aaaa:key:bbb222bc-430a-4de9-9aad-84e5bb022222")
				Expect(gatewayMacsecConfigTemplatePrimaryCakModel.Crn).To(Equal(core.StringPtr("crn:v1:bluemix:public:hs-crypto:us-south:a/4111d05f36894e3cb9b46a43556d9000:abc111b8-37aa-4034-9def-f2607c87aaaa:key:bbb222bc-430a-4de9-9aad-84e5bb022222")))

				// Construct an instance of the GatewayMacsecConfigTemplate model
				gatewayMacsecConfigTemplateModel := new(directlinkv1.GatewayMacsecConfigTemplate)
				Expect(gatewayMacsecConfigTemplateModel).ToNot(BeNil())
				gatewayMacsecConfigTemplateModel.Active = core.BoolPtr(true)
				gatewayMacsecConfigTemplateModel.FallbackCak = gatewayMacsecConfigTemplateFallbackCakModel
				gatewayMacsecConfigTemplateModel.PrimaryCak = gatewayMacsecConfigTemplatePrimaryCakModel
				gatewayMacsecConfigTemplateModel.WindowSize = core.Int64Ptr(int64(148809600))
				Expect(gatewayMacsecConfigTemplateModel.Active).To(Equal(core.BoolPtr(true)))
				Expect(gatewayMacsecConfigTemplateModel.FallbackCak).To(Equal(gatewayMacsecConfigTemplateFallbackCakModel))
				Expect(gatewayMacsecConfigTemplateModel.PrimaryCak).To(Equal(gatewayMacsecConfigTemplatePrimaryCakModel))
				Expect(gatewayMacsecConfigTemplateModel.WindowSize).To(Equal(core.Int64Ptr(int64(148809600))))

				// Construct an instance of the ResourceGroupIdentity model
				resourceGroupIdentityModel := new(directlinkv1.ResourceGroupIdentity)
				Expect(resourceGroupIdentityModel).ToNot(BeNil())
				resourceGroupIdentityModel.ID = core.StringPtr("56969d6043e9465c883cb9f7363e78e8")
				Expect(resourceGroupIdentityModel.ID).To(Equal(core.StringPtr("56969d6043e9465c883cb9f7363e78e8")))

				// Construct an instance of the GatewayTemplateGatewayTypeDedicatedTemplate model
				gatewayTemplateModel := new(directlinkv1.GatewayTemplateGatewayTypeDedicatedTemplate)
				Expect(gatewayTemplateModel).ToNot(BeNil())
				gatewayTemplateModel.BgpAsn = core.Int64Ptr(int64(64999))
				gatewayTemplateModel.BgpBaseCidr = core.StringPtr("testString")
				gatewayTemplateModel.BgpCerCidr = core.StringPtr("169.254.0.10/30")
				gatewayTemplateModel.BgpIbmCidr = core.StringPtr("169.254.0.9/30")
				gatewayTemplateModel.Global = core.BoolPtr(true)
				gatewayTemplateModel.Metered = core.BoolPtr(false)
				gatewayTemplateModel.Name = core.StringPtr("myGateway")
				gatewayTemplateModel.ResourceGroup = resourceGroupIdentityModel
				gatewayTemplateModel.SpeedMbps = core.Int64Ptr(int64(1000))
				gatewayTemplateModel.Type = core.StringPtr("dedicated")
				gatewayTemplateModel.CarrierName = core.StringPtr("myCarrierName")
				gatewayTemplateModel.CrossConnectRouter = core.StringPtr("xcr01.dal03")
				gatewayTemplateModel.CustomerName = core.StringPtr("newCustomerName")
				gatewayTemplateModel.LocationName = core.StringPtr("dal03")
				gatewayTemplateModel.MacsecConfig = gatewayMacsecConfigTemplateModel
				Expect(gatewayTemplateModel.BgpAsn).To(Equal(core.Int64Ptr(int64(64999))))
				Expect(gatewayTemplateModel.BgpBaseCidr).To(Equal(core.StringPtr("testString")))
				Expect(gatewayTemplateModel.BgpCerCidr).To(Equal(core.StringPtr("169.254.0.10/30")))
				Expect(gatewayTemplateModel.BgpIbmCidr).To(Equal(core.StringPtr("169.254.0.9/30")))
				Expect(gatewayTemplateModel.Global).To(Equal(core.BoolPtr(true)))
				Expect(gatewayTemplateModel.Metered).To(Equal(core.BoolPtr(false)))
				Expect(gatewayTemplateModel.Name).To(Equal(core.StringPtr("myGateway")))
				Expect(gatewayTemplateModel.ResourceGroup).To(Equal(resourceGroupIdentityModel))
				Expect(gatewayTemplateModel.SpeedMbps).To(Equal(core.Int64Ptr(int64(1000))))
				Expect(gatewayTemplateModel.Type).To(Equal(core.StringPtr("dedicated")))
				Expect(gatewayTemplateModel.CarrierName).To(Equal(core.StringPtr("myCarrierName")))
				Expect(gatewayTemplateModel.CrossConnectRouter).To(Equal(core.StringPtr("xcr01.dal03")))
				Expect(gatewayTemplateModel.CustomerName).To(Equal(core.StringPtr("newCustomerName")))
				Expect(gatewayTemplateModel.LocationName).To(Equal(core.StringPtr("dal03")))
				Expect(gatewayTemplateModel.MacsecConfig).To(Equal(gatewayMacsecConfigTemplateModel))

				// Construct an instance of the CreateGatewayOptions model
				var gatewayTemplate directlinkv1.GatewayTemplateIntf = nil
				createGatewayOptionsModel := testService.NewCreateGatewayOptions(gatewayTemplate)
				createGatewayOptionsModel.SetGatewayTemplate(gatewayTemplateModel)
				createGatewayOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createGatewayOptionsModel).ToNot(BeNil())
				Expect(createGatewayOptionsModel.GatewayTemplate).To(Equal(gatewayTemplateModel))
				Expect(createGatewayOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateGatewayVirtualConnectionOptions successfully`, func() {
				// Construct an instance of the CreateGatewayVirtualConnectionOptions model
				gatewayID := "testString"
				createGatewayVirtualConnectionOptionsName := "newVC"
				createGatewayVirtualConnectionOptionsType := "vpc"
				createGatewayVirtualConnectionOptionsModel := testService.NewCreateGatewayVirtualConnectionOptions(gatewayID, createGatewayVirtualConnectionOptionsName, createGatewayVirtualConnectionOptionsType)
				createGatewayVirtualConnectionOptionsModel.SetGatewayID("testString")
				createGatewayVirtualConnectionOptionsModel.SetName("newVC")
				createGatewayVirtualConnectionOptionsModel.SetType("vpc")
				createGatewayVirtualConnectionOptionsModel.SetNetworkID("crn:v1:bluemix:public:is:us-east:a/28e4d90ac7504be69447111122223333::vpc:aaa81ac8-5e96-42a0-a4b7-6c2e2d1bbbbb")
				createGatewayVirtualConnectionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createGatewayVirtualConnectionOptionsModel).ToNot(BeNil())
				Expect(createGatewayVirtualConnectionOptionsModel.GatewayID).To(Equal(core.StringPtr("testString")))
				Expect(createGatewayVirtualConnectionOptionsModel.Name).To(Equal(core.StringPtr("newVC")))
				Expect(createGatewayVirtualConnectionOptionsModel.Type).To(Equal(core.StringPtr("vpc")))
				Expect(createGatewayVirtualConnectionOptionsModel.NetworkID).To(Equal(core.StringPtr("crn:v1:bluemix:public:is:us-east:a/28e4d90ac7504be69447111122223333::vpc:aaa81ac8-5e96-42a0-a4b7-6c2e2d1bbbbb")))
				Expect(createGatewayVirtualConnectionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteGatewayOptions successfully`, func() {
				// Construct an instance of the DeleteGatewayOptions model
				id := "testString"
				deleteGatewayOptionsModel := testService.NewDeleteGatewayOptions(id)
				deleteGatewayOptionsModel.SetID("testString")
				deleteGatewayOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteGatewayOptionsModel).ToNot(BeNil())
				Expect(deleteGatewayOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(deleteGatewayOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteGatewayVirtualConnectionOptions successfully`, func() {
				// Construct an instance of the DeleteGatewayVirtualConnectionOptions model
				gatewayID := "testString"
				id := "testString"
				deleteGatewayVirtualConnectionOptionsModel := testService.NewDeleteGatewayVirtualConnectionOptions(gatewayID, id)
				deleteGatewayVirtualConnectionOptionsModel.SetGatewayID("testString")
				deleteGatewayVirtualConnectionOptionsModel.SetID("testString")
				deleteGatewayVirtualConnectionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteGatewayVirtualConnectionOptionsModel).ToNot(BeNil())
				Expect(deleteGatewayVirtualConnectionOptionsModel.GatewayID).To(Equal(core.StringPtr("testString")))
				Expect(deleteGatewayVirtualConnectionOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(deleteGatewayVirtualConnectionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGatewayMacsecConfigPatchTemplateFallbackCak successfully`, func() {
				crn := "crn:v1:bluemix:public:hs-crypto:us-south:a/4111d05f36894e3cb9b46a43556d9000:abc111b8-37aa-4034-9def-f2607c87aaaa:key:bbb222bc-430a-4de9-9aad-84e5bb022222"
				model, err := testService.NewGatewayMacsecConfigPatchTemplateFallbackCak(crn)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewGatewayMacsecConfigPatchTemplatePrimaryCak successfully`, func() {
				crn := "crn:v1:bluemix:public:hs-crypto:us-south:a/4111d05f36894e3cb9b46a43556d9000:abc111b8-37aa-4034-9def-f2607c87aaaa:key:bbb222bc-430a-4de9-9aad-84e5bb022222"
				model, err := testService.NewGatewayMacsecConfigPatchTemplatePrimaryCak(crn)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewGatewayMacsecConfigTemplate successfully`, func() {
				active := true
				var primaryCak *directlinkv1.GatewayMacsecConfigTemplatePrimaryCak = nil
				_, err := testService.NewGatewayMacsecConfigTemplate(active, primaryCak)
				Expect(err).ToNot(BeNil())
			})
			It(`Invoke NewGatewayMacsecConfigTemplateFallbackCak successfully`, func() {
				crn := "crn:v1:bluemix:public:hs-crypto:us-south:a/4111d05f36894e3cb9b46a43556d9000:abc111b8-37aa-4034-9def-f2607c87aaaa:key:bbb222bc-430a-4de9-9aad-84e5bb022222"
				model, err := testService.NewGatewayMacsecConfigTemplateFallbackCak(crn)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewGatewayMacsecConfigTemplatePrimaryCak successfully`, func() {
				crn := "crn:v1:bluemix:public:hs-crypto:us-south:a/4111d05f36894e3cb9b46a43556d9000:abc111b8-37aa-4034-9def-f2607c87aaaa:key:bbb222bc-430a-4de9-9aad-84e5bb022222"
				model, err := testService.NewGatewayMacsecConfigTemplatePrimaryCak(crn)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewGatewayPortIdentity successfully`, func() {
				id := "fffdcb1a-fee4-41c7-9e11-9cd99e65c777"
				model, err := testService.NewGatewayPortIdentity(id)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewGetGatewayOptions successfully`, func() {
				// Construct an instance of the GetGatewayOptions model
				id := "testString"
				getGatewayOptionsModel := testService.NewGetGatewayOptions(id)
				getGatewayOptionsModel.SetID("testString")
				getGatewayOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getGatewayOptionsModel).ToNot(BeNil())
				Expect(getGatewayOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(getGatewayOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetGatewayStatisticsOptions successfully`, func() {
				// Construct an instance of the GetGatewayStatisticsOptions model
				id := "testString"
				typeVar := "macsec_mka"
				getGatewayStatisticsOptionsModel := testService.NewGetGatewayStatisticsOptions(id, typeVar)
				getGatewayStatisticsOptionsModel.SetID("testString")
				getGatewayStatisticsOptionsModel.SetType("macsec_mka")
				getGatewayStatisticsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getGatewayStatisticsOptionsModel).ToNot(BeNil())
				Expect(getGatewayStatisticsOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(getGatewayStatisticsOptionsModel.Type).To(Equal(core.StringPtr("macsec_mka")))
				Expect(getGatewayStatisticsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetGatewayVirtualConnectionOptions successfully`, func() {
				// Construct an instance of the GetGatewayVirtualConnectionOptions model
				gatewayID := "testString"
				id := "testString"
				getGatewayVirtualConnectionOptionsModel := testService.NewGetGatewayVirtualConnectionOptions(gatewayID, id)
				getGatewayVirtualConnectionOptionsModel.SetGatewayID("testString")
				getGatewayVirtualConnectionOptionsModel.SetID("testString")
				getGatewayVirtualConnectionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getGatewayVirtualConnectionOptionsModel).ToNot(BeNil())
				Expect(getGatewayVirtualConnectionOptionsModel.GatewayID).To(Equal(core.StringPtr("testString")))
				Expect(getGatewayVirtualConnectionOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(getGatewayVirtualConnectionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetPortOptions successfully`, func() {
				// Construct an instance of the GetPortOptions model
				id := "testString"
				getPortOptionsModel := testService.NewGetPortOptions(id)
				getPortOptionsModel.SetID("testString")
				getPortOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getPortOptionsModel).ToNot(BeNil())
				Expect(getPortOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(getPortOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListGatewayCompletionNoticeOptions successfully`, func() {
				// Construct an instance of the ListGatewayCompletionNoticeOptions model
				id := "testString"
				listGatewayCompletionNoticeOptionsModel := testService.NewListGatewayCompletionNoticeOptions(id)
				listGatewayCompletionNoticeOptionsModel.SetID("testString")
				listGatewayCompletionNoticeOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listGatewayCompletionNoticeOptionsModel).ToNot(BeNil())
				Expect(listGatewayCompletionNoticeOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(listGatewayCompletionNoticeOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListGatewayLetterOfAuthorizationOptions successfully`, func() {
				// Construct an instance of the ListGatewayLetterOfAuthorizationOptions model
				id := "testString"
				listGatewayLetterOfAuthorizationOptionsModel := testService.NewListGatewayLetterOfAuthorizationOptions(id)
				listGatewayLetterOfAuthorizationOptionsModel.SetID("testString")
				listGatewayLetterOfAuthorizationOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listGatewayLetterOfAuthorizationOptionsModel).ToNot(BeNil())
				Expect(listGatewayLetterOfAuthorizationOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(listGatewayLetterOfAuthorizationOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListGatewayVirtualConnectionsOptions successfully`, func() {
				// Construct an instance of the ListGatewayVirtualConnectionsOptions model
				gatewayID := "testString"
				listGatewayVirtualConnectionsOptionsModel := testService.NewListGatewayVirtualConnectionsOptions(gatewayID)
				listGatewayVirtualConnectionsOptionsModel.SetGatewayID("testString")
				listGatewayVirtualConnectionsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listGatewayVirtualConnectionsOptionsModel).ToNot(BeNil())
				Expect(listGatewayVirtualConnectionsOptionsModel.GatewayID).To(Equal(core.StringPtr("testString")))
				Expect(listGatewayVirtualConnectionsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListGatewaysOptions successfully`, func() {
				// Construct an instance of the ListGatewaysOptions model
				listGatewaysOptionsModel := testService.NewListGatewaysOptions()
				listGatewaysOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listGatewaysOptionsModel).ToNot(BeNil())
				Expect(listGatewaysOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListOfferingTypeLocationCrossConnectRoutersOptions successfully`, func() {
				// Construct an instance of the ListOfferingTypeLocationCrossConnectRoutersOptions model
				offeringType := "dedicated"
				locationName := "testString"
				listOfferingTypeLocationCrossConnectRoutersOptionsModel := testService.NewListOfferingTypeLocationCrossConnectRoutersOptions(offeringType, locationName)
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
				listOfferingTypeLocationsOptionsModel := testService.NewListOfferingTypeLocationsOptions(offeringType)
				listOfferingTypeLocationsOptionsModel.SetOfferingType("dedicated")
				listOfferingTypeLocationsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listOfferingTypeLocationsOptionsModel).ToNot(BeNil())
				Expect(listOfferingTypeLocationsOptionsModel.OfferingType).To(Equal(core.StringPtr("dedicated")))
				Expect(listOfferingTypeLocationsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListOfferingTypeSpeedsOptions successfully`, func() {
				// Construct an instance of the ListOfferingTypeSpeedsOptions model
				offeringType := "dedicated"
				listOfferingTypeSpeedsOptionsModel := testService.NewListOfferingTypeSpeedsOptions(offeringType)
				listOfferingTypeSpeedsOptionsModel.SetOfferingType("dedicated")
				listOfferingTypeSpeedsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listOfferingTypeSpeedsOptionsModel).ToNot(BeNil())
				Expect(listOfferingTypeSpeedsOptionsModel.OfferingType).To(Equal(core.StringPtr("dedicated")))
				Expect(listOfferingTypeSpeedsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListPortsOptions successfully`, func() {
				// Construct an instance of the ListPortsOptions model
				listPortsOptionsModel := testService.NewListPortsOptions()
				listPortsOptionsModel.SetStart("testString")
				listPortsOptionsModel.SetLimit(int64(38))
				listPortsOptionsModel.SetLocationName("testString")
				listPortsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listPortsOptionsModel).ToNot(BeNil())
				Expect(listPortsOptionsModel.Start).To(Equal(core.StringPtr("testString")))
				Expect(listPortsOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(38))))
				Expect(listPortsOptionsModel.LocationName).To(Equal(core.StringPtr("testString")))
				Expect(listPortsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewResourceGroupIdentity successfully`, func() {
				id := "56969d6043e9465c883cb9f7363e78e8"
				model, err := testService.NewResourceGroupIdentity(id)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewUpdateGatewayOptions successfully`, func() {
				// Construct an instance of the GatewayMacsecConfigPatchTemplateFallbackCak model
				gatewayMacsecConfigPatchTemplateFallbackCakModel := new(directlinkv1.GatewayMacsecConfigPatchTemplateFallbackCak)
				Expect(gatewayMacsecConfigPatchTemplateFallbackCakModel).ToNot(BeNil())
				gatewayMacsecConfigPatchTemplateFallbackCakModel.Crn = core.StringPtr("crn:v1:bluemix:public:hs-crypto:us-south:a/4111d05f36894e3cb9b46a43556d9000:abc111b8-37aa-4034-9def-f2607c87aaaa:key:bbb222bc-430a-4de9-9aad-84e5bb022222")
				Expect(gatewayMacsecConfigPatchTemplateFallbackCakModel.Crn).To(Equal(core.StringPtr("crn:v1:bluemix:public:hs-crypto:us-south:a/4111d05f36894e3cb9b46a43556d9000:abc111b8-37aa-4034-9def-f2607c87aaaa:key:bbb222bc-430a-4de9-9aad-84e5bb022222")))

				// Construct an instance of the GatewayMacsecConfigPatchTemplatePrimaryCak model
				gatewayMacsecConfigPatchTemplatePrimaryCakModel := new(directlinkv1.GatewayMacsecConfigPatchTemplatePrimaryCak)
				Expect(gatewayMacsecConfigPatchTemplatePrimaryCakModel).ToNot(BeNil())
				gatewayMacsecConfigPatchTemplatePrimaryCakModel.Crn = core.StringPtr("crn:v1:bluemix:public:hs-crypto:us-south:a/4111d05f36894e3cb9b46a43556d9000:abc111b8-37aa-4034-9def-f2607c87aaaa:key:bbb222bc-430a-4de9-9aad-84e5bb022222")
				Expect(gatewayMacsecConfigPatchTemplatePrimaryCakModel.Crn).To(Equal(core.StringPtr("crn:v1:bluemix:public:hs-crypto:us-south:a/4111d05f36894e3cb9b46a43556d9000:abc111b8-37aa-4034-9def-f2607c87aaaa:key:bbb222bc-430a-4de9-9aad-84e5bb022222")))

				// Construct an instance of the GatewayMacsecConfigPatchTemplate model
				gatewayMacsecConfigPatchTemplateModel := new(directlinkv1.GatewayMacsecConfigPatchTemplate)
				Expect(gatewayMacsecConfigPatchTemplateModel).ToNot(BeNil())
				gatewayMacsecConfigPatchTemplateModel.Active = core.BoolPtr(true)
				gatewayMacsecConfigPatchTemplateModel.FallbackCak = gatewayMacsecConfigPatchTemplateFallbackCakModel
				gatewayMacsecConfigPatchTemplateModel.PrimaryCak = gatewayMacsecConfigPatchTemplatePrimaryCakModel
				gatewayMacsecConfigPatchTemplateModel.WindowSize = core.Int64Ptr(int64(512))
				Expect(gatewayMacsecConfigPatchTemplateModel.Active).To(Equal(core.BoolPtr(true)))
				Expect(gatewayMacsecConfigPatchTemplateModel.FallbackCak).To(Equal(gatewayMacsecConfigPatchTemplateFallbackCakModel))
				Expect(gatewayMacsecConfigPatchTemplateModel.PrimaryCak).To(Equal(gatewayMacsecConfigPatchTemplatePrimaryCakModel))
				Expect(gatewayMacsecConfigPatchTemplateModel.WindowSize).To(Equal(core.Int64Ptr(int64(512))))

				// Construct an instance of the UpdateGatewayOptions model
				id := "testString"
				updateGatewayOptionsModel := testService.NewUpdateGatewayOptions(id)
				updateGatewayOptionsModel.SetID("testString")
				updateGatewayOptionsModel.SetGlobal(true)
				updateGatewayOptionsModel.SetLoaRejectReason("The port mentioned was incorrect")
				updateGatewayOptionsModel.SetMacsecConfig(gatewayMacsecConfigPatchTemplateModel)
				updateGatewayOptionsModel.SetMetered(false)
				updateGatewayOptionsModel.SetName("testGateway")
				updateGatewayOptionsModel.SetOperationalStatus("loa_accepted")
				updateGatewayOptionsModel.SetSpeedMbps(int64(1000))
				updateGatewayOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateGatewayOptionsModel).ToNot(BeNil())
				Expect(updateGatewayOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(updateGatewayOptionsModel.Global).To(Equal(core.BoolPtr(true)))
				Expect(updateGatewayOptionsModel.LoaRejectReason).To(Equal(core.StringPtr("The port mentioned was incorrect")))
				Expect(updateGatewayOptionsModel.MacsecConfig).To(Equal(gatewayMacsecConfigPatchTemplateModel))
				Expect(updateGatewayOptionsModel.Metered).To(Equal(core.BoolPtr(false)))
				Expect(updateGatewayOptionsModel.Name).To(Equal(core.StringPtr("testGateway")))
				Expect(updateGatewayOptionsModel.OperationalStatus).To(Equal(core.StringPtr("loa_accepted")))
				Expect(updateGatewayOptionsModel.SpeedMbps).To(Equal(core.Int64Ptr(int64(1000))))
				Expect(updateGatewayOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateGatewayVirtualConnectionOptions successfully`, func() {
				// Construct an instance of the UpdateGatewayVirtualConnectionOptions model
				gatewayID := "testString"
				id := "testString"
				updateGatewayVirtualConnectionOptionsModel := testService.NewUpdateGatewayVirtualConnectionOptions(gatewayID, id)
				updateGatewayVirtualConnectionOptionsModel.SetGatewayID("testString")
				updateGatewayVirtualConnectionOptionsModel.SetID("testString")
				updateGatewayVirtualConnectionOptionsModel.SetName("newConnectionName")
				updateGatewayVirtualConnectionOptionsModel.SetStatus("attached")
				updateGatewayVirtualConnectionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateGatewayVirtualConnectionOptionsModel).ToNot(BeNil())
				Expect(updateGatewayVirtualConnectionOptionsModel.GatewayID).To(Equal(core.StringPtr("testString")))
				Expect(updateGatewayVirtualConnectionOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(updateGatewayVirtualConnectionOptionsModel.Name).To(Equal(core.StringPtr("newConnectionName")))
				Expect(updateGatewayVirtualConnectionOptionsModel.Status).To(Equal(core.StringPtr("attached")))
				Expect(updateGatewayVirtualConnectionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGatewayTemplateGatewayTypeConnectTemplate successfully`, func() {
				bgpAsn := int64(64999)
				global := true
				metered := false
				name := "myGateway"
				speedMbps := int64(1000)
				typeVar := "dedicated"
				var port *directlinkv1.GatewayPortIdentity = nil
				_, err := testService.NewGatewayTemplateGatewayTypeConnectTemplate(bgpAsn, global, metered, name, speedMbps, typeVar, port)
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
				model, err := testService.NewGatewayTemplateGatewayTypeDedicatedTemplate(bgpAsn, global, metered, name, speedMbps, typeVar, carrierName, crossConnectRouter, customerName, locationName)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
		})
	})
	Describe(`Utility function tests`, func() {
		It(`Invoke CreateMockByteArray() successfully`, func() {
			mockByteArray := CreateMockByteArray("This is a test")
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
			mockDate := CreateMockDate()
			Expect(mockDate).ToNot(BeNil())
		})
		It(`Invoke CreateMockDateTime() successfully`, func() {
			mockDateTime := CreateMockDateTime()
			Expect(mockDateTime).ToNot(BeNil())
		})
	})
})

//
// Utility functions used by the generated test code
//

func CreateMockByteArray(mockData string) *[]byte {
	ba := make([]byte, 0)
	ba = append(ba, mockData...)
	return &ba
}

func CreateMockUUID(mockData string) *strfmt.UUID {
	uuid := strfmt.UUID(mockData)
	return &uuid
}

func CreateMockReader(mockData string) io.ReadCloser {
	return ioutil.NopCloser(bytes.NewReader([]byte(mockData)))
}

func CreateMockDate() *strfmt.Date {
	d := strfmt.Date(time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC))
	return &d
}

func CreateMockDateTime() *strfmt.DateTime {
	d := strfmt.DateTime(time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC))
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
