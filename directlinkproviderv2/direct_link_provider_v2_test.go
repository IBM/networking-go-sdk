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

package directlinkproviderv2_test

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
	"github.com/IBM/networking-go-sdk/directlinkproviderv2"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe(`DirectLinkProviderV2`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		version := "testString"
		It(`Instantiate service client`, func() {
			testService, testServiceErr := directlinkproviderv2.NewDirectLinkProviderV2(&directlinkproviderv2.DirectLinkProviderV2Options{
				Authenticator: &core.NoAuthAuthenticator{},
				Version:       core.StringPtr(version),
			})
			Expect(testService).ToNot(BeNil())
			Expect(testServiceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			testService, testServiceErr := directlinkproviderv2.NewDirectLinkProviderV2(&directlinkproviderv2.DirectLinkProviderV2Options{
				URL:     "{BAD_URL_STRING",
				Version: core.StringPtr(version),
			})
			Expect(testService).To(BeNil())
			Expect(testServiceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			testService, testServiceErr := directlinkproviderv2.NewDirectLinkProviderV2(&directlinkproviderv2.DirectLinkProviderV2Options{
				URL:     "https://directlinkproviderv2/api",
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
			testService, testServiceErr := directlinkproviderv2.NewDirectLinkProviderV2(&directlinkproviderv2.DirectLinkProviderV2Options{})
			Expect(testService).To(BeNil())
			Expect(testServiceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		version := "testString"
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"DIRECT_LINK_PROVIDER_URL":       "https://directlinkproviderv2/api",
				"DIRECT_LINK_PROVIDER_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := directlinkproviderv2.NewDirectLinkProviderV2UsingExternalConfig(&directlinkproviderv2.DirectLinkProviderV2Options{
					Version: core.StringPtr(version),
				})
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := directlinkproviderv2.NewDirectLinkProviderV2UsingExternalConfig(&directlinkproviderv2.DirectLinkProviderV2Options{
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
				testService, testServiceErr := directlinkproviderv2.NewDirectLinkProviderV2UsingExternalConfig(&directlinkproviderv2.DirectLinkProviderV2Options{
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
				"DIRECT_LINK_PROVIDER_URL":       "https://directlinkproviderv2/api",
				"DIRECT_LINK_PROVIDER_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			testService, testServiceErr := directlinkproviderv2.NewDirectLinkProviderV2UsingExternalConfig(&directlinkproviderv2.DirectLinkProviderV2Options{
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
				"DIRECT_LINK_PROVIDER_AUTH_TYPE": "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			testService, testServiceErr := directlinkproviderv2.NewDirectLinkProviderV2UsingExternalConfig(&directlinkproviderv2.DirectLinkProviderV2Options{
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
	Describe(`ListProviderGateways(listProviderGatewaysOptions *ListProviderGatewaysOptions) - Operation response error`, func() {
		version := "testString"
		listProviderGatewaysPath := "/gateways"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(listProviderGatewaysPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(38))}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListProviderGateways with error: Operation response processing error`, func() {
				testService, testServiceErr := directlinkproviderv2.NewDirectLinkProviderV2(&directlinkproviderv2.DirectLinkProviderV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ListProviderGatewaysOptions model
				listProviderGatewaysOptionsModel := new(directlinkproviderv2.ListProviderGatewaysOptions)
				listProviderGatewaysOptionsModel.Start = core.StringPtr("testString")
				listProviderGatewaysOptionsModel.Limit = core.Int64Ptr(int64(38))
				listProviderGatewaysOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.ListProviderGateways(listProviderGatewaysOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ListProviderGateways(listProviderGatewaysOptions *ListProviderGatewaysOptions)`, func() {
		version := "testString"
		listProviderGatewaysPath := "/gateways"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(listProviderGatewaysPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(38))}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"first": {"href": "https://directlink.cloud.ibm.com/provider/v2/gateways?limit=100"}, "limit": 100, "next": {"href": "https://directlink.cloud.ibm.com/provider/v2/gateways?start=8c4a91a3e2cbd233b5a5b33436855fc2&limit=100", "start": "8c4a91a3e2cbd233b5a5b33436855fc2"}, "total_count": 132, "gateways": [{"bgp_asn": 64999, "bgp_cer_cidr": "10.254.30.78/30", "bgp_ibm_asn": 13884, "bgp_ibm_cidr": "10.254.30.77/30", "bgp_status": "active", "change_request": {"type": "create_gateway"}, "created_at": "2019-01-01T12:00:00", "crn": "crn:v1:bluemix:public:directlink:dal03:a/57a7d05f36894e3cb9b46a43556d903e::connect:ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4", "customer_account_id": "57a7d05f36894e3cb9b46a43556d903e", "id": "ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4", "name": "myGateway", "operational_status": "create_pending", "port": {"id": "fffdcb1a-fee4-41c7-9e11-9cd99e65c777"}, "provider_api_managed": true, "speed_mbps": 1000, "type": "connect", "vlan": 10}]}`)
				}))
			})
			It(`Invoke ListProviderGateways successfully`, func() {
				testService, testServiceErr := directlinkproviderv2.NewDirectLinkProviderV2(&directlinkproviderv2.DirectLinkProviderV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.ListProviderGateways(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListProviderGatewaysOptions model
				listProviderGatewaysOptionsModel := new(directlinkproviderv2.ListProviderGatewaysOptions)
				listProviderGatewaysOptionsModel.Start = core.StringPtr("testString")
				listProviderGatewaysOptionsModel.Limit = core.Int64Ptr(int64(38))
				listProviderGatewaysOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.ListProviderGateways(listProviderGatewaysOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke ListProviderGateways with error: Operation request error`, func() {
				testService, testServiceErr := directlinkproviderv2.NewDirectLinkProviderV2(&directlinkproviderv2.DirectLinkProviderV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ListProviderGatewaysOptions model
				listProviderGatewaysOptionsModel := new(directlinkproviderv2.ListProviderGatewaysOptions)
				listProviderGatewaysOptionsModel.Start = core.StringPtr("testString")
				listProviderGatewaysOptionsModel.Limit = core.Int64Ptr(int64(38))
				listProviderGatewaysOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.ListProviderGateways(listProviderGatewaysOptionsModel)
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
	Describe(`CreateProviderGateway(createProviderGatewayOptions *CreateProviderGatewayOptions) - Operation response error`, func() {
		version := "testString"
		createProviderGatewayPath := "/gateways"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(createProviderGatewayPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["check_only"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateProviderGateway with error: Operation response processing error`, func() {
				testService, testServiceErr := directlinkproviderv2.NewDirectLinkProviderV2(&directlinkproviderv2.DirectLinkProviderV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ProviderGatewayPortIdentity model
				providerGatewayPortIdentityModel := new(directlinkproviderv2.ProviderGatewayPortIdentity)
				providerGatewayPortIdentityModel.ID = core.StringPtr("fffdcb1a-fee4-41c7-9e11-9cd99e65c777")

				// Construct an instance of the CreateProviderGatewayOptions model
				createProviderGatewayOptionsModel := new(directlinkproviderv2.CreateProviderGatewayOptions)
				createProviderGatewayOptionsModel.BgpAsn = core.Int64Ptr(int64(64999))
				createProviderGatewayOptionsModel.CustomerAccountID = core.StringPtr("57a7d05f36894e3cb9b46a43556d903e")
				createProviderGatewayOptionsModel.Name = core.StringPtr("myGateway")
				createProviderGatewayOptionsModel.Port = providerGatewayPortIdentityModel
				createProviderGatewayOptionsModel.SpeedMbps = core.Int64Ptr(int64(1000))
				createProviderGatewayOptionsModel.BgpCerCidr = core.StringPtr("10.254.30.78/30")
				createProviderGatewayOptionsModel.BgpIbmCidr = core.StringPtr("10.254.30.77/30")
				createProviderGatewayOptionsModel.CheckOnly = core.StringPtr("testString")
				createProviderGatewayOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.CreateProviderGateway(createProviderGatewayOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`CreateProviderGateway(createProviderGatewayOptions *CreateProviderGatewayOptions)`, func() {
		version := "testString"
		createProviderGatewayPath := "/gateways"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(createProviderGatewayPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["check_only"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, `{"bgp_asn": 64999, "bgp_cer_cidr": "10.254.30.78/30", "bgp_ibm_asn": 13884, "bgp_ibm_cidr": "10.254.30.77/30", "bgp_status": "active", "change_request": {"type": "create_gateway"}, "created_at": "2019-01-01T12:00:00", "crn": "crn:v1:bluemix:public:directlink:dal03:a/57a7d05f36894e3cb9b46a43556d903e::connect:ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4", "customer_account_id": "57a7d05f36894e3cb9b46a43556d903e", "id": "ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4", "name": "myGateway", "operational_status": "create_pending", "port": {"id": "fffdcb1a-fee4-41c7-9e11-9cd99e65c777"}, "provider_api_managed": true, "speed_mbps": 1000, "type": "connect", "vlan": 10}`)
				}))
			})
			It(`Invoke CreateProviderGateway successfully`, func() {
				testService, testServiceErr := directlinkproviderv2.NewDirectLinkProviderV2(&directlinkproviderv2.DirectLinkProviderV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.CreateProviderGateway(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ProviderGatewayPortIdentity model
				providerGatewayPortIdentityModel := new(directlinkproviderv2.ProviderGatewayPortIdentity)
				providerGatewayPortIdentityModel.ID = core.StringPtr("fffdcb1a-fee4-41c7-9e11-9cd99e65c777")

				// Construct an instance of the CreateProviderGatewayOptions model
				createProviderGatewayOptionsModel := new(directlinkproviderv2.CreateProviderGatewayOptions)
				createProviderGatewayOptionsModel.BgpAsn = core.Int64Ptr(int64(64999))
				createProviderGatewayOptionsModel.CustomerAccountID = core.StringPtr("57a7d05f36894e3cb9b46a43556d903e")
				createProviderGatewayOptionsModel.Name = core.StringPtr("myGateway")
				createProviderGatewayOptionsModel.Port = providerGatewayPortIdentityModel
				createProviderGatewayOptionsModel.SpeedMbps = core.Int64Ptr(int64(1000))
				createProviderGatewayOptionsModel.BgpCerCidr = core.StringPtr("10.254.30.78/30")
				createProviderGatewayOptionsModel.BgpIbmCidr = core.StringPtr("10.254.30.77/30")
				createProviderGatewayOptionsModel.CheckOnly = core.StringPtr("testString")
				createProviderGatewayOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.CreateProviderGateway(createProviderGatewayOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke CreateProviderGateway with error: Operation validation and request error`, func() {
				testService, testServiceErr := directlinkproviderv2.NewDirectLinkProviderV2(&directlinkproviderv2.DirectLinkProviderV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ProviderGatewayPortIdentity model
				providerGatewayPortIdentityModel := new(directlinkproviderv2.ProviderGatewayPortIdentity)
				providerGatewayPortIdentityModel.ID = core.StringPtr("fffdcb1a-fee4-41c7-9e11-9cd99e65c777")

				// Construct an instance of the CreateProviderGatewayOptions model
				createProviderGatewayOptionsModel := new(directlinkproviderv2.CreateProviderGatewayOptions)
				createProviderGatewayOptionsModel.BgpAsn = core.Int64Ptr(int64(64999))
				createProviderGatewayOptionsModel.CustomerAccountID = core.StringPtr("57a7d05f36894e3cb9b46a43556d903e")
				createProviderGatewayOptionsModel.Name = core.StringPtr("myGateway")
				createProviderGatewayOptionsModel.Port = providerGatewayPortIdentityModel
				createProviderGatewayOptionsModel.SpeedMbps = core.Int64Ptr(int64(1000))
				createProviderGatewayOptionsModel.BgpCerCidr = core.StringPtr("10.254.30.78/30")
				createProviderGatewayOptionsModel.BgpIbmCidr = core.StringPtr("10.254.30.77/30")
				createProviderGatewayOptionsModel.CheckOnly = core.StringPtr("testString")
				createProviderGatewayOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.CreateProviderGateway(createProviderGatewayOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateProviderGatewayOptions model with no property values
				createProviderGatewayOptionsModelNew := new(directlinkproviderv2.CreateProviderGatewayOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.CreateProviderGateway(createProviderGatewayOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteProviderGateway(deleteProviderGatewayOptions *DeleteProviderGatewayOptions) - Operation response error`, func() {
		version := "testString"
		deleteProviderGatewayPath := "/gateways/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(deleteProviderGatewayPath))
					Expect(req.Method).To(Equal("DELETE"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke DeleteProviderGateway with error: Operation response processing error`, func() {
				testService, testServiceErr := directlinkproviderv2.NewDirectLinkProviderV2(&directlinkproviderv2.DirectLinkProviderV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the DeleteProviderGatewayOptions model
				deleteProviderGatewayOptionsModel := new(directlinkproviderv2.DeleteProviderGatewayOptions)
				deleteProviderGatewayOptionsModel.ID = core.StringPtr("testString")
				deleteProviderGatewayOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.DeleteProviderGateway(deleteProviderGatewayOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`DeleteProviderGateway(deleteProviderGatewayOptions *DeleteProviderGatewayOptions)`, func() {
		version := "testString"
		deleteProviderGatewayPath := "/gateways/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(deleteProviderGatewayPath))
					Expect(req.Method).To(Equal("DELETE"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, `{"bgp_asn": 64999, "bgp_cer_cidr": "10.254.30.78/30", "bgp_ibm_asn": 13884, "bgp_ibm_cidr": "10.254.30.77/30", "bgp_status": "active", "change_request": {"type": "create_gateway"}, "created_at": "2019-01-01T12:00:00", "crn": "crn:v1:bluemix:public:directlink:dal03:a/57a7d05f36894e3cb9b46a43556d903e::connect:ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4", "customer_account_id": "57a7d05f36894e3cb9b46a43556d903e", "id": "ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4", "name": "myGateway", "operational_status": "create_pending", "port": {"id": "fffdcb1a-fee4-41c7-9e11-9cd99e65c777"}, "provider_api_managed": true, "speed_mbps": 1000, "type": "connect", "vlan": 10}`)
				}))
			})
			It(`Invoke DeleteProviderGateway successfully`, func() {
				testService, testServiceErr := directlinkproviderv2.NewDirectLinkProviderV2(&directlinkproviderv2.DirectLinkProviderV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.DeleteProviderGateway(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DeleteProviderGatewayOptions model
				deleteProviderGatewayOptionsModel := new(directlinkproviderv2.DeleteProviderGatewayOptions)
				deleteProviderGatewayOptionsModel.ID = core.StringPtr("testString")
				deleteProviderGatewayOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.DeleteProviderGateway(deleteProviderGatewayOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke DeleteProviderGateway with error: Operation validation and request error`, func() {
				testService, testServiceErr := directlinkproviderv2.NewDirectLinkProviderV2(&directlinkproviderv2.DirectLinkProviderV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the DeleteProviderGatewayOptions model
				deleteProviderGatewayOptionsModel := new(directlinkproviderv2.DeleteProviderGatewayOptions)
				deleteProviderGatewayOptionsModel.ID = core.StringPtr("testString")
				deleteProviderGatewayOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.DeleteProviderGateway(deleteProviderGatewayOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the DeleteProviderGatewayOptions model with no property values
				deleteProviderGatewayOptionsModelNew := new(directlinkproviderv2.DeleteProviderGatewayOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.DeleteProviderGateway(deleteProviderGatewayOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetProviderGateway(getProviderGatewayOptions *GetProviderGatewayOptions) - Operation response error`, func() {
		version := "testString"
		getProviderGatewayPath := "/gateways/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getProviderGatewayPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetProviderGateway with error: Operation response processing error`, func() {
				testService, testServiceErr := directlinkproviderv2.NewDirectLinkProviderV2(&directlinkproviderv2.DirectLinkProviderV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetProviderGatewayOptions model
				getProviderGatewayOptionsModel := new(directlinkproviderv2.GetProviderGatewayOptions)
				getProviderGatewayOptionsModel.ID = core.StringPtr("testString")
				getProviderGatewayOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.GetProviderGateway(getProviderGatewayOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetProviderGateway(getProviderGatewayOptions *GetProviderGatewayOptions)`, func() {
		version := "testString"
		getProviderGatewayPath := "/gateways/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getProviderGatewayPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"bgp_asn": 64999, "bgp_cer_cidr": "10.254.30.78/30", "bgp_ibm_asn": 13884, "bgp_ibm_cidr": "10.254.30.77/30", "bgp_status": "active", "change_request": {"type": "create_gateway"}, "created_at": "2019-01-01T12:00:00", "crn": "crn:v1:bluemix:public:directlink:dal03:a/57a7d05f36894e3cb9b46a43556d903e::connect:ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4", "customer_account_id": "57a7d05f36894e3cb9b46a43556d903e", "id": "ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4", "name": "myGateway", "operational_status": "create_pending", "port": {"id": "fffdcb1a-fee4-41c7-9e11-9cd99e65c777"}, "provider_api_managed": true, "speed_mbps": 1000, "type": "connect", "vlan": 10}`)
				}))
			})
			It(`Invoke GetProviderGateway successfully`, func() {
				testService, testServiceErr := directlinkproviderv2.NewDirectLinkProviderV2(&directlinkproviderv2.DirectLinkProviderV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetProviderGateway(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetProviderGatewayOptions model
				getProviderGatewayOptionsModel := new(directlinkproviderv2.GetProviderGatewayOptions)
				getProviderGatewayOptionsModel.ID = core.StringPtr("testString")
				getProviderGatewayOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetProviderGateway(getProviderGatewayOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke GetProviderGateway with error: Operation validation and request error`, func() {
				testService, testServiceErr := directlinkproviderv2.NewDirectLinkProviderV2(&directlinkproviderv2.DirectLinkProviderV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetProviderGatewayOptions model
				getProviderGatewayOptionsModel := new(directlinkproviderv2.GetProviderGatewayOptions)
				getProviderGatewayOptionsModel.ID = core.StringPtr("testString")
				getProviderGatewayOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.GetProviderGateway(getProviderGatewayOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetProviderGatewayOptions model with no property values
				getProviderGatewayOptionsModelNew := new(directlinkproviderv2.GetProviderGatewayOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.GetProviderGateway(getProviderGatewayOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateProviderGateway(updateProviderGatewayOptions *UpdateProviderGatewayOptions) - Operation response error`, func() {
		version := "testString"
		updateProviderGatewayPath := "/gateways/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(updateProviderGatewayPath))
					Expect(req.Method).To(Equal("PATCH"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateProviderGateway with error: Operation response processing error`, func() {
				testService, testServiceErr := directlinkproviderv2.NewDirectLinkProviderV2(&directlinkproviderv2.DirectLinkProviderV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the UpdateProviderGatewayOptions model
				updateProviderGatewayOptionsModel := new(directlinkproviderv2.UpdateProviderGatewayOptions)
				updateProviderGatewayOptionsModel.ID = core.StringPtr("testString")
				updateProviderGatewayOptionsModel.Name = core.StringPtr("myNewGateway")
				updateProviderGatewayOptionsModel.SpeedMbps = core.Int64Ptr(int64(1000))
				updateProviderGatewayOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.UpdateProviderGateway(updateProviderGatewayOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UpdateProviderGateway(updateProviderGatewayOptions *UpdateProviderGatewayOptions)`, func() {
		version := "testString"
		updateProviderGatewayPath := "/gateways/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(updateProviderGatewayPath))
					Expect(req.Method).To(Equal("PATCH"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"bgp_asn": 64999, "bgp_cer_cidr": "10.254.30.78/30", "bgp_ibm_asn": 13884, "bgp_ibm_cidr": "10.254.30.77/30", "bgp_status": "active", "change_request": {"type": "create_gateway"}, "created_at": "2019-01-01T12:00:00", "crn": "crn:v1:bluemix:public:directlink:dal03:a/57a7d05f36894e3cb9b46a43556d903e::connect:ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4", "customer_account_id": "57a7d05f36894e3cb9b46a43556d903e", "id": "ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4", "name": "myGateway", "operational_status": "create_pending", "port": {"id": "fffdcb1a-fee4-41c7-9e11-9cd99e65c777"}, "provider_api_managed": true, "speed_mbps": 1000, "type": "connect", "vlan": 10}`)
				}))
			})
			It(`Invoke UpdateProviderGateway successfully`, func() {
				testService, testServiceErr := directlinkproviderv2.NewDirectLinkProviderV2(&directlinkproviderv2.DirectLinkProviderV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.UpdateProviderGateway(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateProviderGatewayOptions model
				updateProviderGatewayOptionsModel := new(directlinkproviderv2.UpdateProviderGatewayOptions)
				updateProviderGatewayOptionsModel.ID = core.StringPtr("testString")
				updateProviderGatewayOptionsModel.Name = core.StringPtr("myNewGateway")
				updateProviderGatewayOptionsModel.SpeedMbps = core.Int64Ptr(int64(1000))
				updateProviderGatewayOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.UpdateProviderGateway(updateProviderGatewayOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke UpdateProviderGateway with error: Operation validation and request error`, func() {
				testService, testServiceErr := directlinkproviderv2.NewDirectLinkProviderV2(&directlinkproviderv2.DirectLinkProviderV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the UpdateProviderGatewayOptions model
				updateProviderGatewayOptionsModel := new(directlinkproviderv2.UpdateProviderGatewayOptions)
				updateProviderGatewayOptionsModel.ID = core.StringPtr("testString")
				updateProviderGatewayOptionsModel.Name = core.StringPtr("myNewGateway")
				updateProviderGatewayOptionsModel.SpeedMbps = core.Int64Ptr(int64(1000))
				updateProviderGatewayOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.UpdateProviderGateway(updateProviderGatewayOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateProviderGatewayOptions model with no property values
				updateProviderGatewayOptionsModelNew := new(directlinkproviderv2.UpdateProviderGatewayOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.UpdateProviderGateway(updateProviderGatewayOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListProviderPorts(listProviderPortsOptions *ListProviderPortsOptions) - Operation response error`, func() {
		version := "testString"
		listProviderPortsPath := "/ports"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(listProviderPortsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(38))}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListProviderPorts with error: Operation response processing error`, func() {
				testService, testServiceErr := directlinkproviderv2.NewDirectLinkProviderV2(&directlinkproviderv2.DirectLinkProviderV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ListProviderPortsOptions model
				listProviderPortsOptionsModel := new(directlinkproviderv2.ListProviderPortsOptions)
				listProviderPortsOptionsModel.Start = core.StringPtr("testString")
				listProviderPortsOptionsModel.Limit = core.Int64Ptr(int64(38))
				listProviderPortsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.ListProviderPorts(listProviderPortsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ListProviderPorts(listProviderPortsOptions *ListProviderPortsOptions)`, func() {
		version := "testString"
		listProviderPortsPath := "/ports"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(listProviderPortsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(38))}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"first": {"href": "https://directlink.cloud.ibm.com/provider/v2/ports?limit=100"}, "limit": 100, "next": {"href": "https://directlink.cloud.ibm.com/provider/v2/ports?start=9d5a91a3e2cbd233b5a5b33436855ed1&limit=100", "start": "9d5a91a3e2cbd233b5a5b33436855ed1"}, "total_count": 132, "ports": [{"id": "01122b9b-820f-4c44-8a31-77f1f0806765", "label": "XCR-FRK-CS-SEC-01", "location_display_name": "Dallas 03", "location_name": "dal03", "provider_name": "provider_1", "supported_link_speeds": [19]}]}`)
				}))
			})
			It(`Invoke ListProviderPorts successfully`, func() {
				testService, testServiceErr := directlinkproviderv2.NewDirectLinkProviderV2(&directlinkproviderv2.DirectLinkProviderV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.ListProviderPorts(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListProviderPortsOptions model
				listProviderPortsOptionsModel := new(directlinkproviderv2.ListProviderPortsOptions)
				listProviderPortsOptionsModel.Start = core.StringPtr("testString")
				listProviderPortsOptionsModel.Limit = core.Int64Ptr(int64(38))
				listProviderPortsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.ListProviderPorts(listProviderPortsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke ListProviderPorts with error: Operation request error`, func() {
				testService, testServiceErr := directlinkproviderv2.NewDirectLinkProviderV2(&directlinkproviderv2.DirectLinkProviderV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ListProviderPortsOptions model
				listProviderPortsOptionsModel := new(directlinkproviderv2.ListProviderPortsOptions)
				listProviderPortsOptionsModel.Start = core.StringPtr("testString")
				listProviderPortsOptionsModel.Limit = core.Int64Ptr(int64(38))
				listProviderPortsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.ListProviderPorts(listProviderPortsOptionsModel)
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
	Describe(`GetProviderPort(getProviderPortOptions *GetProviderPortOptions) - Operation response error`, func() {
		version := "testString"
		getProviderPortPath := "/ports/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getProviderPortPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetProviderPort with error: Operation response processing error`, func() {
				testService, testServiceErr := directlinkproviderv2.NewDirectLinkProviderV2(&directlinkproviderv2.DirectLinkProviderV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetProviderPortOptions model
				getProviderPortOptionsModel := new(directlinkproviderv2.GetProviderPortOptions)
				getProviderPortOptionsModel.ID = core.StringPtr("testString")
				getProviderPortOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.GetProviderPort(getProviderPortOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetProviderPort(getProviderPortOptions *GetProviderPortOptions)`, func() {
		version := "testString"
		getProviderPortPath := "/ports/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getProviderPortPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"id": "01122b9b-820f-4c44-8a31-77f1f0806765", "label": "XCR-FRK-CS-SEC-01", "location_display_name": "Dallas 03", "location_name": "dal03", "provider_name": "provider_1", "supported_link_speeds": [19]}`)
				}))
			})
			It(`Invoke GetProviderPort successfully`, func() {
				testService, testServiceErr := directlinkproviderv2.NewDirectLinkProviderV2(&directlinkproviderv2.DirectLinkProviderV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetProviderPort(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetProviderPortOptions model
				getProviderPortOptionsModel := new(directlinkproviderv2.GetProviderPortOptions)
				getProviderPortOptionsModel.ID = core.StringPtr("testString")
				getProviderPortOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetProviderPort(getProviderPortOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke GetProviderPort with error: Operation validation and request error`, func() {
				testService, testServiceErr := directlinkproviderv2.NewDirectLinkProviderV2(&directlinkproviderv2.DirectLinkProviderV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetProviderPortOptions model
				getProviderPortOptionsModel := new(directlinkproviderv2.GetProviderPortOptions)
				getProviderPortOptionsModel.ID = core.StringPtr("testString")
				getProviderPortOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.GetProviderPort(getProviderPortOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetProviderPortOptions model with no property values
				getProviderPortOptionsModelNew := new(directlinkproviderv2.GetProviderPortOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.GetProviderPort(getProviderPortOptionsModelNew)
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
			testService, _ := directlinkproviderv2.NewDirectLinkProviderV2(&directlinkproviderv2.DirectLinkProviderV2Options{
				URL:           "http://directlinkproviderv2modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
				Version:       core.StringPtr(version),
			})
			It(`Invoke NewCreateProviderGatewayOptions successfully`, func() {
				// Construct an instance of the ProviderGatewayPortIdentity model
				providerGatewayPortIdentityModel := new(directlinkproviderv2.ProviderGatewayPortIdentity)
				Expect(providerGatewayPortIdentityModel).ToNot(BeNil())
				providerGatewayPortIdentityModel.ID = core.StringPtr("fffdcb1a-fee4-41c7-9e11-9cd99e65c777")
				Expect(providerGatewayPortIdentityModel.ID).To(Equal(core.StringPtr("fffdcb1a-fee4-41c7-9e11-9cd99e65c777")))

				// Construct an instance of the CreateProviderGatewayOptions model
				createProviderGatewayOptionsBgpAsn := int64(64999)
				createProviderGatewayOptionsCustomerAccountID := "57a7d05f36894e3cb9b46a43556d903e"
				createProviderGatewayOptionsName := "myGateway"
				var createProviderGatewayOptionsPort *directlinkproviderv2.ProviderGatewayPortIdentity = nil
				createProviderGatewayOptionsSpeedMbps := int64(1000)
				createProviderGatewayOptionsModel := testService.NewCreateProviderGatewayOptions(createProviderGatewayOptionsBgpAsn, createProviderGatewayOptionsCustomerAccountID, createProviderGatewayOptionsName, createProviderGatewayOptionsPort, createProviderGatewayOptionsSpeedMbps)
				createProviderGatewayOptionsModel.SetBgpAsn(int64(64999))
				createProviderGatewayOptionsModel.SetCustomerAccountID("57a7d05f36894e3cb9b46a43556d903e")
				createProviderGatewayOptionsModel.SetName("myGateway")
				createProviderGatewayOptionsModel.SetPort(providerGatewayPortIdentityModel)
				createProviderGatewayOptionsModel.SetSpeedMbps(int64(1000))
				createProviderGatewayOptionsModel.SetBgpCerCidr("10.254.30.78/30")
				createProviderGatewayOptionsModel.SetBgpIbmCidr("10.254.30.77/30")
				createProviderGatewayOptionsModel.SetCheckOnly("testString")
				createProviderGatewayOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createProviderGatewayOptionsModel).ToNot(BeNil())
				Expect(createProviderGatewayOptionsModel.BgpAsn).To(Equal(core.Int64Ptr(int64(64999))))
				Expect(createProviderGatewayOptionsModel.CustomerAccountID).To(Equal(core.StringPtr("57a7d05f36894e3cb9b46a43556d903e")))
				Expect(createProviderGatewayOptionsModel.Name).To(Equal(core.StringPtr("myGateway")))
				Expect(createProviderGatewayOptionsModel.Port).To(Equal(providerGatewayPortIdentityModel))
				Expect(createProviderGatewayOptionsModel.SpeedMbps).To(Equal(core.Int64Ptr(int64(1000))))
				Expect(createProviderGatewayOptionsModel.BgpCerCidr).To(Equal(core.StringPtr("10.254.30.78/30")))
				Expect(createProviderGatewayOptionsModel.BgpIbmCidr).To(Equal(core.StringPtr("10.254.30.77/30")))
				Expect(createProviderGatewayOptionsModel.CheckOnly).To(Equal(core.StringPtr("testString")))
				Expect(createProviderGatewayOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteProviderGatewayOptions successfully`, func() {
				// Construct an instance of the DeleteProviderGatewayOptions model
				id := "testString"
				deleteProviderGatewayOptionsModel := testService.NewDeleteProviderGatewayOptions(id)
				deleteProviderGatewayOptionsModel.SetID("testString")
				deleteProviderGatewayOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteProviderGatewayOptionsModel).ToNot(BeNil())
				Expect(deleteProviderGatewayOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(deleteProviderGatewayOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetProviderGatewayOptions successfully`, func() {
				// Construct an instance of the GetProviderGatewayOptions model
				id := "testString"
				getProviderGatewayOptionsModel := testService.NewGetProviderGatewayOptions(id)
				getProviderGatewayOptionsModel.SetID("testString")
				getProviderGatewayOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getProviderGatewayOptionsModel).ToNot(BeNil())
				Expect(getProviderGatewayOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(getProviderGatewayOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetProviderPortOptions successfully`, func() {
				// Construct an instance of the GetProviderPortOptions model
				id := "testString"
				getProviderPortOptionsModel := testService.NewGetProviderPortOptions(id)
				getProviderPortOptionsModel.SetID("testString")
				getProviderPortOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getProviderPortOptionsModel).ToNot(BeNil())
				Expect(getProviderPortOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(getProviderPortOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListProviderGatewaysOptions successfully`, func() {
				// Construct an instance of the ListProviderGatewaysOptions model
				listProviderGatewaysOptionsModel := testService.NewListProviderGatewaysOptions()
				listProviderGatewaysOptionsModel.SetStart("testString")
				listProviderGatewaysOptionsModel.SetLimit(int64(38))
				listProviderGatewaysOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listProviderGatewaysOptionsModel).ToNot(BeNil())
				Expect(listProviderGatewaysOptionsModel.Start).To(Equal(core.StringPtr("testString")))
				Expect(listProviderGatewaysOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(38))))
				Expect(listProviderGatewaysOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListProviderPortsOptions successfully`, func() {
				// Construct an instance of the ListProviderPortsOptions model
				listProviderPortsOptionsModel := testService.NewListProviderPortsOptions()
				listProviderPortsOptionsModel.SetStart("testString")
				listProviderPortsOptionsModel.SetLimit(int64(38))
				listProviderPortsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listProviderPortsOptionsModel).ToNot(BeNil())
				Expect(listProviderPortsOptionsModel.Start).To(Equal(core.StringPtr("testString")))
				Expect(listProviderPortsOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(38))))
				Expect(listProviderPortsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewProviderGatewayPortIdentity successfully`, func() {
				id := "fffdcb1a-fee4-41c7-9e11-9cd99e65c777"
				model, err := testService.NewProviderGatewayPortIdentity(id)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewUpdateProviderGatewayOptions successfully`, func() {
				// Construct an instance of the UpdateProviderGatewayOptions model
				id := "testString"
				updateProviderGatewayOptionsModel := testService.NewUpdateProviderGatewayOptions(id)
				updateProviderGatewayOptionsModel.SetID("testString")
				updateProviderGatewayOptionsModel.SetName("myNewGateway")
				updateProviderGatewayOptionsModel.SetSpeedMbps(int64(1000))
				updateProviderGatewayOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateProviderGatewayOptionsModel).ToNot(BeNil())
				Expect(updateProviderGatewayOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(updateProviderGatewayOptionsModel.Name).To(Equal(core.StringPtr("myNewGateway")))
				Expect(updateProviderGatewayOptionsModel.SpeedMbps).To(Equal(core.Int64Ptr(int64(1000))))
				Expect(updateProviderGatewayOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
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
