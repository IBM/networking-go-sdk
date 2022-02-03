/**
 * (C) Copyright IBM Corp. 2021.
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
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
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
			directLinkProviderService, serviceErr := directlinkproviderv2.NewDirectLinkProviderV2(&directlinkproviderv2.DirectLinkProviderV2Options{
				Authenticator: &core.NoAuthAuthenticator{},
				Version:       core.StringPtr(version),
			})
			Expect(directLinkProviderService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			directLinkProviderService, serviceErr := directlinkproviderv2.NewDirectLinkProviderV2(&directlinkproviderv2.DirectLinkProviderV2Options{
				URL:     "{BAD_URL_STRING",
				Version: core.StringPtr(version),
			})
			Expect(directLinkProviderService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			directLinkProviderService, serviceErr := directlinkproviderv2.NewDirectLinkProviderV2(&directlinkproviderv2.DirectLinkProviderV2Options{
				URL:     "https://directlinkproviderv2/api",
				Version: core.StringPtr(version),
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(directLinkProviderService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Validation Error`, func() {
			directLinkProviderService, serviceErr := directlinkproviderv2.NewDirectLinkProviderV2(&directlinkproviderv2.DirectLinkProviderV2Options{})
			Expect(directLinkProviderService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
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
				directLinkProviderService, serviceErr := directlinkproviderv2.NewDirectLinkProviderV2UsingExternalConfig(&directlinkproviderv2.DirectLinkProviderV2Options{
					Version: core.StringPtr(version),
				})
				Expect(directLinkProviderService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := directLinkProviderService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != directLinkProviderService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(directLinkProviderService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(directLinkProviderService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				directLinkProviderService, serviceErr := directlinkproviderv2.NewDirectLinkProviderV2UsingExternalConfig(&directlinkproviderv2.DirectLinkProviderV2Options{
					URL:     "https://testService/api",
					Version: core.StringPtr(version),
				})
				Expect(directLinkProviderService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(directLinkProviderService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := directLinkProviderService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != directLinkProviderService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(directLinkProviderService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(directLinkProviderService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				directLinkProviderService, serviceErr := directlinkproviderv2.NewDirectLinkProviderV2UsingExternalConfig(&directlinkproviderv2.DirectLinkProviderV2Options{
					Version: core.StringPtr(version),
				})
				err := directLinkProviderService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(directLinkProviderService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(directLinkProviderService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := directLinkProviderService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != directLinkProviderService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(directLinkProviderService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(directLinkProviderService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"DIRECT_LINK_PROVIDER_URL":       "https://directlinkproviderv2/api",
				"DIRECT_LINK_PROVIDER_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			directLinkProviderService, serviceErr := directlinkproviderv2.NewDirectLinkProviderV2UsingExternalConfig(&directlinkproviderv2.DirectLinkProviderV2Options{
				Version: core.StringPtr(version),
			})

			It(`Instantiate service client with error`, func() {
				Expect(directLinkProviderService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"DIRECT_LINK_PROVIDER_AUTH_TYPE": "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			directLinkProviderService, serviceErr := directlinkproviderv2.NewDirectLinkProviderV2UsingExternalConfig(&directlinkproviderv2.DirectLinkProviderV2Options{
				URL:     "{BAD_URL_STRING",
				Version: core.StringPtr(version),
			})

			It(`Instantiate service client with error`, func() {
				Expect(directLinkProviderService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = directlinkproviderv2.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`ListProviderGateways(listProviderGatewaysOptions *ListProviderGatewaysOptions) - Operation response error`, func() {
		version := "testString"
		listProviderGatewaysPath := "/gateways"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listProviderGatewaysPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(1))}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListProviderGateways with error: Operation response processing error`, func() {
				directLinkProviderService, serviceErr := directlinkproviderv2.NewDirectLinkProviderV2(&directlinkproviderv2.DirectLinkProviderV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkProviderService).ToNot(BeNil())

				// Construct an instance of the ListProviderGatewaysOptions model
				listProviderGatewaysOptionsModel := new(directlinkproviderv2.ListProviderGatewaysOptions)
				listProviderGatewaysOptionsModel.Start = core.StringPtr("testString")
				listProviderGatewaysOptionsModel.Limit = core.Int64Ptr(int64(1))
				listProviderGatewaysOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := directLinkProviderService.ListProviderGateways(listProviderGatewaysOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				directLinkProviderService.EnableRetries(0, 0)
				result, response, operationErr = directLinkProviderService.ListProviderGateways(listProviderGatewaysOptionsModel)
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
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listProviderGatewaysPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(1))}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"first": {"href": "https://directlink.cloud.ibm.com/provider/v2/gateways?limit=100"}, "limit": 100, "next": {"href": "https://directlink.cloud.ibm.com/provider/v2/gateways?start=8c4a91a3e2cbd233b5a5b33436855fc2&limit=100", "start": "8c4a91a3e2cbd233b5a5b33436855fc2"}, "total_count": 132, "gateways": [{"bgp_asn": 64999, "bgp_cer_cidr": "10.254.30.78/30", "bgp_ibm_asn": 13884, "bgp_ibm_cidr": "10.254.30.77/30", "bgp_status": "active", "change_request": {"type": "create_gateway"}, "created_at": "2019-01-01T12:00:00.000Z", "crn": "crn:v1:bluemix:public:directlink:dal03:a/4111d05f36894e3cb9b46a43556d9000::connect:ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4", "customer_account_id": "4111d05f36894e3cb9b46a43556d9000", "id": "ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4", "name": "myGateway", "operational_status": "configuring", "port": {"id": "fffdcb1a-fee4-41c7-9e11-9cd99e65c777"}, "provider_api_managed": true, "speed_mbps": 1000, "type": "connect", "vlan": 10}]}`)
				}))
			})
			It(`Invoke ListProviderGateways successfully with retries`, func() {
				directLinkProviderService, serviceErr := directlinkproviderv2.NewDirectLinkProviderV2(&directlinkproviderv2.DirectLinkProviderV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkProviderService).ToNot(BeNil())
				directLinkProviderService.EnableRetries(0, 0)

				// Construct an instance of the ListProviderGatewaysOptions model
				listProviderGatewaysOptionsModel := new(directlinkproviderv2.ListProviderGatewaysOptions)
				listProviderGatewaysOptionsModel.Start = core.StringPtr("testString")
				listProviderGatewaysOptionsModel.Limit = core.Int64Ptr(int64(1))
				listProviderGatewaysOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := directLinkProviderService.ListProviderGatewaysWithContext(ctx, listProviderGatewaysOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				directLinkProviderService.DisableRetries()
				result, response, operationErr := directLinkProviderService.ListProviderGateways(listProviderGatewaysOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = directLinkProviderService.ListProviderGatewaysWithContext(ctx, listProviderGatewaysOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listProviderGatewaysPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(1))}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"first": {"href": "https://directlink.cloud.ibm.com/provider/v2/gateways?limit=100"}, "limit": 100, "next": {"href": "https://directlink.cloud.ibm.com/provider/v2/gateways?start=8c4a91a3e2cbd233b5a5b33436855fc2&limit=100", "start": "8c4a91a3e2cbd233b5a5b33436855fc2"}, "total_count": 132, "gateways": [{"bgp_asn": 64999, "bgp_cer_cidr": "10.254.30.78/30", "bgp_ibm_asn": 13884, "bgp_ibm_cidr": "10.254.30.77/30", "bgp_status": "active", "change_request": {"type": "create_gateway"}, "created_at": "2019-01-01T12:00:00.000Z", "crn": "crn:v1:bluemix:public:directlink:dal03:a/4111d05f36894e3cb9b46a43556d9000::connect:ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4", "customer_account_id": "4111d05f36894e3cb9b46a43556d9000", "id": "ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4", "name": "myGateway", "operational_status": "configuring", "port": {"id": "fffdcb1a-fee4-41c7-9e11-9cd99e65c777"}, "provider_api_managed": true, "speed_mbps": 1000, "type": "connect", "vlan": 10}]}`)
				}))
			})
			It(`Invoke ListProviderGateways successfully`, func() {
				directLinkProviderService, serviceErr := directlinkproviderv2.NewDirectLinkProviderV2(&directlinkproviderv2.DirectLinkProviderV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkProviderService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := directLinkProviderService.ListProviderGateways(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListProviderGatewaysOptions model
				listProviderGatewaysOptionsModel := new(directlinkproviderv2.ListProviderGatewaysOptions)
				listProviderGatewaysOptionsModel.Start = core.StringPtr("testString")
				listProviderGatewaysOptionsModel.Limit = core.Int64Ptr(int64(1))
				listProviderGatewaysOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = directLinkProviderService.ListProviderGateways(listProviderGatewaysOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListProviderGateways with error: Operation request error`, func() {
				directLinkProviderService, serviceErr := directlinkproviderv2.NewDirectLinkProviderV2(&directlinkproviderv2.DirectLinkProviderV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkProviderService).ToNot(BeNil())

				// Construct an instance of the ListProviderGatewaysOptions model
				listProviderGatewaysOptionsModel := new(directlinkproviderv2.ListProviderGatewaysOptions)
				listProviderGatewaysOptionsModel.Start = core.StringPtr("testString")
				listProviderGatewaysOptionsModel.Limit = core.Int64Ptr(int64(1))
				listProviderGatewaysOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := directLinkProviderService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := directLinkProviderService.ListProviderGateways(listProviderGatewaysOptionsModel)
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
			It(`Invoke ListProviderGateways successfully`, func() {
				directLinkProviderService, serviceErr := directlinkproviderv2.NewDirectLinkProviderV2(&directlinkproviderv2.DirectLinkProviderV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkProviderService).ToNot(BeNil())

				// Construct an instance of the ListProviderGatewaysOptions model
				listProviderGatewaysOptionsModel := new(directlinkproviderv2.ListProviderGatewaysOptions)
				listProviderGatewaysOptionsModel.Start = core.StringPtr("testString")
				listProviderGatewaysOptionsModel.Limit = core.Int64Ptr(int64(1))
				listProviderGatewaysOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := directLinkProviderService.ListProviderGateways(listProviderGatewaysOptionsModel)
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
				responseObject := new(directlinkproviderv2.ProviderGatewayCollection)
				nextObject := new(directlinkproviderv2.ProviderGatewayCollectionNext)
				nextObject.Start = core.StringPtr("abc-123")
				responseObject.Next = nextObject

				value, err := responseObject.GetNextStart()
				Expect(err).To(BeNil())
				Expect(value).To(Equal(core.StringPtr("abc-123")))
			})
			It(`Invoke GetNextStart without a "Next" property in the response`, func() {
				responseObject := new(directlinkproviderv2.ProviderGatewayCollection)

				value, err := responseObject.GetNextStart()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
		})
	})
	Describe(`CreateProviderGateway(createProviderGatewayOptions *CreateProviderGatewayOptions) - Operation response error`, func() {
		version := "testString"
		createProviderGatewayPath := "/gateways"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createProviderGatewayPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["check_only"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateProviderGateway with error: Operation response processing error`, func() {
				directLinkProviderService, serviceErr := directlinkproviderv2.NewDirectLinkProviderV2(&directlinkproviderv2.DirectLinkProviderV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkProviderService).ToNot(BeNil())

				// Construct an instance of the ProviderGatewayPortIdentity model
				providerGatewayPortIdentityModel := new(directlinkproviderv2.ProviderGatewayPortIdentity)
				providerGatewayPortIdentityModel.ID = core.StringPtr("fffdcb1a-fee4-41c7-9e11-9cd99e65c777")

				// Construct an instance of the CreateProviderGatewayOptions model
				createProviderGatewayOptionsModel := new(directlinkproviderv2.CreateProviderGatewayOptions)
				createProviderGatewayOptionsModel.BgpAsn = core.Int64Ptr(int64(64999))
				createProviderGatewayOptionsModel.CustomerAccountID = core.StringPtr("4111d05f36894e3cb9b46a43556d9000")
				createProviderGatewayOptionsModel.Name = core.StringPtr("myGateway")
				createProviderGatewayOptionsModel.Port = providerGatewayPortIdentityModel
				createProviderGatewayOptionsModel.SpeedMbps = core.Int64Ptr(int64(1000))
				createProviderGatewayOptionsModel.BgpCerCidr = core.StringPtr("10.254.30.78/30")
				createProviderGatewayOptionsModel.BgpIbmCidr = core.StringPtr("10.254.30.77/30")
				createProviderGatewayOptionsModel.Vlan = core.Int64Ptr(int64(10))
				createProviderGatewayOptionsModel.CheckOnly = core.StringPtr("testString")
				createProviderGatewayOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := directLinkProviderService.CreateProviderGateway(createProviderGatewayOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				directLinkProviderService.EnableRetries(0, 0)
				result, response, operationErr = directLinkProviderService.CreateProviderGateway(createProviderGatewayOptionsModel)
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
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createProviderGatewayPath))
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
					Expect(req.URL.Query()["check_only"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"bgp_asn": 64999, "bgp_cer_cidr": "10.254.30.78/30", "bgp_ibm_asn": 13884, "bgp_ibm_cidr": "10.254.30.77/30", "bgp_status": "active", "change_request": {"type": "create_gateway"}, "created_at": "2019-01-01T12:00:00.000Z", "crn": "crn:v1:bluemix:public:directlink:dal03:a/4111d05f36894e3cb9b46a43556d9000::connect:ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4", "customer_account_id": "4111d05f36894e3cb9b46a43556d9000", "id": "ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4", "name": "myGateway", "operational_status": "configuring", "port": {"id": "fffdcb1a-fee4-41c7-9e11-9cd99e65c777"}, "provider_api_managed": true, "speed_mbps": 1000, "type": "connect", "vlan": 10}`)
				}))
			})
			It(`Invoke CreateProviderGateway successfully with retries`, func() {
				directLinkProviderService, serviceErr := directlinkproviderv2.NewDirectLinkProviderV2(&directlinkproviderv2.DirectLinkProviderV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkProviderService).ToNot(BeNil())
				directLinkProviderService.EnableRetries(0, 0)

				// Construct an instance of the ProviderGatewayPortIdentity model
				providerGatewayPortIdentityModel := new(directlinkproviderv2.ProviderGatewayPortIdentity)
				providerGatewayPortIdentityModel.ID = core.StringPtr("fffdcb1a-fee4-41c7-9e11-9cd99e65c777")

				// Construct an instance of the CreateProviderGatewayOptions model
				createProviderGatewayOptionsModel := new(directlinkproviderv2.CreateProviderGatewayOptions)
				createProviderGatewayOptionsModel.BgpAsn = core.Int64Ptr(int64(64999))
				createProviderGatewayOptionsModel.CustomerAccountID = core.StringPtr("4111d05f36894e3cb9b46a43556d9000")
				createProviderGatewayOptionsModel.Name = core.StringPtr("myGateway")
				createProviderGatewayOptionsModel.Port = providerGatewayPortIdentityModel
				createProviderGatewayOptionsModel.SpeedMbps = core.Int64Ptr(int64(1000))
				createProviderGatewayOptionsModel.BgpCerCidr = core.StringPtr("10.254.30.78/30")
				createProviderGatewayOptionsModel.BgpIbmCidr = core.StringPtr("10.254.30.77/30")
				createProviderGatewayOptionsModel.Vlan = core.Int64Ptr(int64(10))
				createProviderGatewayOptionsModel.CheckOnly = core.StringPtr("testString")
				createProviderGatewayOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := directLinkProviderService.CreateProviderGatewayWithContext(ctx, createProviderGatewayOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				directLinkProviderService.DisableRetries()
				result, response, operationErr := directLinkProviderService.CreateProviderGateway(createProviderGatewayOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = directLinkProviderService.CreateProviderGatewayWithContext(ctx, createProviderGatewayOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createProviderGatewayPath))
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
					Expect(req.URL.Query()["check_only"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"bgp_asn": 64999, "bgp_cer_cidr": "10.254.30.78/30", "bgp_ibm_asn": 13884, "bgp_ibm_cidr": "10.254.30.77/30", "bgp_status": "active", "change_request": {"type": "create_gateway"}, "created_at": "2019-01-01T12:00:00.000Z", "crn": "crn:v1:bluemix:public:directlink:dal03:a/4111d05f36894e3cb9b46a43556d9000::connect:ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4", "customer_account_id": "4111d05f36894e3cb9b46a43556d9000", "id": "ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4", "name": "myGateway", "operational_status": "configuring", "port": {"id": "fffdcb1a-fee4-41c7-9e11-9cd99e65c777"}, "provider_api_managed": true, "speed_mbps": 1000, "type": "connect", "vlan": 10}`)
				}))
			})
			It(`Invoke CreateProviderGateway successfully`, func() {
				directLinkProviderService, serviceErr := directlinkproviderv2.NewDirectLinkProviderV2(&directlinkproviderv2.DirectLinkProviderV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkProviderService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := directLinkProviderService.CreateProviderGateway(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ProviderGatewayPortIdentity model
				providerGatewayPortIdentityModel := new(directlinkproviderv2.ProviderGatewayPortIdentity)
				providerGatewayPortIdentityModel.ID = core.StringPtr("fffdcb1a-fee4-41c7-9e11-9cd99e65c777")

				// Construct an instance of the CreateProviderGatewayOptions model
				createProviderGatewayOptionsModel := new(directlinkproviderv2.CreateProviderGatewayOptions)
				createProviderGatewayOptionsModel.BgpAsn = core.Int64Ptr(int64(64999))
				createProviderGatewayOptionsModel.CustomerAccountID = core.StringPtr("4111d05f36894e3cb9b46a43556d9000")
				createProviderGatewayOptionsModel.Name = core.StringPtr("myGateway")
				createProviderGatewayOptionsModel.Port = providerGatewayPortIdentityModel
				createProviderGatewayOptionsModel.SpeedMbps = core.Int64Ptr(int64(1000))
				createProviderGatewayOptionsModel.BgpCerCidr = core.StringPtr("10.254.30.78/30")
				createProviderGatewayOptionsModel.BgpIbmCidr = core.StringPtr("10.254.30.77/30")
				createProviderGatewayOptionsModel.Vlan = core.Int64Ptr(int64(10))
				createProviderGatewayOptionsModel.CheckOnly = core.StringPtr("testString")
				createProviderGatewayOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = directLinkProviderService.CreateProviderGateway(createProviderGatewayOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateProviderGateway with error: Operation validation and request error`, func() {
				directLinkProviderService, serviceErr := directlinkproviderv2.NewDirectLinkProviderV2(&directlinkproviderv2.DirectLinkProviderV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkProviderService).ToNot(BeNil())

				// Construct an instance of the ProviderGatewayPortIdentity model
				providerGatewayPortIdentityModel := new(directlinkproviderv2.ProviderGatewayPortIdentity)
				providerGatewayPortIdentityModel.ID = core.StringPtr("fffdcb1a-fee4-41c7-9e11-9cd99e65c777")

				// Construct an instance of the CreateProviderGatewayOptions model
				createProviderGatewayOptionsModel := new(directlinkproviderv2.CreateProviderGatewayOptions)
				createProviderGatewayOptionsModel.BgpAsn = core.Int64Ptr(int64(64999))
				createProviderGatewayOptionsModel.CustomerAccountID = core.StringPtr("4111d05f36894e3cb9b46a43556d9000")
				createProviderGatewayOptionsModel.Name = core.StringPtr("myGateway")
				createProviderGatewayOptionsModel.Port = providerGatewayPortIdentityModel
				createProviderGatewayOptionsModel.SpeedMbps = core.Int64Ptr(int64(1000))
				createProviderGatewayOptionsModel.BgpCerCidr = core.StringPtr("10.254.30.78/30")
				createProviderGatewayOptionsModel.BgpIbmCidr = core.StringPtr("10.254.30.77/30")
				createProviderGatewayOptionsModel.Vlan = core.Int64Ptr(int64(10))
				createProviderGatewayOptionsModel.CheckOnly = core.StringPtr("testString")
				createProviderGatewayOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := directLinkProviderService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := directLinkProviderService.CreateProviderGateway(createProviderGatewayOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateProviderGatewayOptions model with no property values
				createProviderGatewayOptionsModelNew := new(directlinkproviderv2.CreateProviderGatewayOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = directLinkProviderService.CreateProviderGateway(createProviderGatewayOptionsModelNew)
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
			It(`Invoke CreateProviderGateway successfully`, func() {
				directLinkProviderService, serviceErr := directlinkproviderv2.NewDirectLinkProviderV2(&directlinkproviderv2.DirectLinkProviderV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkProviderService).ToNot(BeNil())

				// Construct an instance of the ProviderGatewayPortIdentity model
				providerGatewayPortIdentityModel := new(directlinkproviderv2.ProviderGatewayPortIdentity)
				providerGatewayPortIdentityModel.ID = core.StringPtr("fffdcb1a-fee4-41c7-9e11-9cd99e65c777")

				// Construct an instance of the CreateProviderGatewayOptions model
				createProviderGatewayOptionsModel := new(directlinkproviderv2.CreateProviderGatewayOptions)
				createProviderGatewayOptionsModel.BgpAsn = core.Int64Ptr(int64(64999))
				createProviderGatewayOptionsModel.CustomerAccountID = core.StringPtr("4111d05f36894e3cb9b46a43556d9000")
				createProviderGatewayOptionsModel.Name = core.StringPtr("myGateway")
				createProviderGatewayOptionsModel.Port = providerGatewayPortIdentityModel
				createProviderGatewayOptionsModel.SpeedMbps = core.Int64Ptr(int64(1000))
				createProviderGatewayOptionsModel.BgpCerCidr = core.StringPtr("10.254.30.78/30")
				createProviderGatewayOptionsModel.BgpIbmCidr = core.StringPtr("10.254.30.77/30")
				createProviderGatewayOptionsModel.Vlan = core.Int64Ptr(int64(10))
				createProviderGatewayOptionsModel.CheckOnly = core.StringPtr("testString")
				createProviderGatewayOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := directLinkProviderService.CreateProviderGateway(createProviderGatewayOptionsModel)
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
	Describe(`DeleteProviderGateway(deleteProviderGatewayOptions *DeleteProviderGatewayOptions) - Operation response error`, func() {
		version := "testString"
		deleteProviderGatewayPath := "/gateways/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteProviderGatewayPath))
					Expect(req.Method).To(Equal("DELETE"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke DeleteProviderGateway with error: Operation response processing error`, func() {
				directLinkProviderService, serviceErr := directlinkproviderv2.NewDirectLinkProviderV2(&directlinkproviderv2.DirectLinkProviderV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkProviderService).ToNot(BeNil())

				// Construct an instance of the DeleteProviderGatewayOptions model
				deleteProviderGatewayOptionsModel := new(directlinkproviderv2.DeleteProviderGatewayOptions)
				deleteProviderGatewayOptionsModel.ID = core.StringPtr("testString")
				deleteProviderGatewayOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := directLinkProviderService.DeleteProviderGateway(deleteProviderGatewayOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				directLinkProviderService.EnableRetries(0, 0)
				result, response, operationErr = directLinkProviderService.DeleteProviderGateway(deleteProviderGatewayOptionsModel)
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
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteProviderGatewayPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"bgp_asn": 64999, "bgp_cer_cidr": "10.254.30.78/30", "bgp_ibm_asn": 13884, "bgp_ibm_cidr": "10.254.30.77/30", "bgp_status": "active", "change_request": {"type": "create_gateway"}, "created_at": "2019-01-01T12:00:00.000Z", "crn": "crn:v1:bluemix:public:directlink:dal03:a/4111d05f36894e3cb9b46a43556d9000::connect:ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4", "customer_account_id": "4111d05f36894e3cb9b46a43556d9000", "id": "ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4", "name": "myGateway", "operational_status": "configuring", "port": {"id": "fffdcb1a-fee4-41c7-9e11-9cd99e65c777"}, "provider_api_managed": true, "speed_mbps": 1000, "type": "connect", "vlan": 10}`)
				}))
			})
			It(`Invoke DeleteProviderGateway successfully with retries`, func() {
				directLinkProviderService, serviceErr := directlinkproviderv2.NewDirectLinkProviderV2(&directlinkproviderv2.DirectLinkProviderV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkProviderService).ToNot(BeNil())
				directLinkProviderService.EnableRetries(0, 0)

				// Construct an instance of the DeleteProviderGatewayOptions model
				deleteProviderGatewayOptionsModel := new(directlinkproviderv2.DeleteProviderGatewayOptions)
				deleteProviderGatewayOptionsModel.ID = core.StringPtr("testString")
				deleteProviderGatewayOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := directLinkProviderService.DeleteProviderGatewayWithContext(ctx, deleteProviderGatewayOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				directLinkProviderService.DisableRetries()
				result, response, operationErr := directLinkProviderService.DeleteProviderGateway(deleteProviderGatewayOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = directLinkProviderService.DeleteProviderGatewayWithContext(ctx, deleteProviderGatewayOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(deleteProviderGatewayPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"bgp_asn": 64999, "bgp_cer_cidr": "10.254.30.78/30", "bgp_ibm_asn": 13884, "bgp_ibm_cidr": "10.254.30.77/30", "bgp_status": "active", "change_request": {"type": "create_gateway"}, "created_at": "2019-01-01T12:00:00.000Z", "crn": "crn:v1:bluemix:public:directlink:dal03:a/4111d05f36894e3cb9b46a43556d9000::connect:ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4", "customer_account_id": "4111d05f36894e3cb9b46a43556d9000", "id": "ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4", "name": "myGateway", "operational_status": "configuring", "port": {"id": "fffdcb1a-fee4-41c7-9e11-9cd99e65c777"}, "provider_api_managed": true, "speed_mbps": 1000, "type": "connect", "vlan": 10}`)
				}))
			})
			It(`Invoke DeleteProviderGateway successfully`, func() {
				directLinkProviderService, serviceErr := directlinkproviderv2.NewDirectLinkProviderV2(&directlinkproviderv2.DirectLinkProviderV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkProviderService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := directLinkProviderService.DeleteProviderGateway(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DeleteProviderGatewayOptions model
				deleteProviderGatewayOptionsModel := new(directlinkproviderv2.DeleteProviderGatewayOptions)
				deleteProviderGatewayOptionsModel.ID = core.StringPtr("testString")
				deleteProviderGatewayOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = directLinkProviderService.DeleteProviderGateway(deleteProviderGatewayOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke DeleteProviderGateway with error: Operation validation and request error`, func() {
				directLinkProviderService, serviceErr := directlinkproviderv2.NewDirectLinkProviderV2(&directlinkproviderv2.DirectLinkProviderV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkProviderService).ToNot(BeNil())

				// Construct an instance of the DeleteProviderGatewayOptions model
				deleteProviderGatewayOptionsModel := new(directlinkproviderv2.DeleteProviderGatewayOptions)
				deleteProviderGatewayOptionsModel.ID = core.StringPtr("testString")
				deleteProviderGatewayOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := directLinkProviderService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := directLinkProviderService.DeleteProviderGateway(deleteProviderGatewayOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the DeleteProviderGatewayOptions model with no property values
				deleteProviderGatewayOptionsModelNew := new(directlinkproviderv2.DeleteProviderGatewayOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = directLinkProviderService.DeleteProviderGateway(deleteProviderGatewayOptionsModelNew)
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
			It(`Invoke DeleteProviderGateway successfully`, func() {
				directLinkProviderService, serviceErr := directlinkproviderv2.NewDirectLinkProviderV2(&directlinkproviderv2.DirectLinkProviderV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkProviderService).ToNot(BeNil())

				// Construct an instance of the DeleteProviderGatewayOptions model
				deleteProviderGatewayOptionsModel := new(directlinkproviderv2.DeleteProviderGatewayOptions)
				deleteProviderGatewayOptionsModel.ID = core.StringPtr("testString")
				deleteProviderGatewayOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := directLinkProviderService.DeleteProviderGateway(deleteProviderGatewayOptionsModel)
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
	Describe(`GetProviderGateway(getProviderGatewayOptions *GetProviderGatewayOptions) - Operation response error`, func() {
		version := "testString"
		getProviderGatewayPath := "/gateways/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getProviderGatewayPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetProviderGateway with error: Operation response processing error`, func() {
				directLinkProviderService, serviceErr := directlinkproviderv2.NewDirectLinkProviderV2(&directlinkproviderv2.DirectLinkProviderV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkProviderService).ToNot(BeNil())

				// Construct an instance of the GetProviderGatewayOptions model
				getProviderGatewayOptionsModel := new(directlinkproviderv2.GetProviderGatewayOptions)
				getProviderGatewayOptionsModel.ID = core.StringPtr("testString")
				getProviderGatewayOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := directLinkProviderService.GetProviderGateway(getProviderGatewayOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				directLinkProviderService.EnableRetries(0, 0)
				result, response, operationErr = directLinkProviderService.GetProviderGateway(getProviderGatewayOptionsModel)
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
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getProviderGatewayPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"bgp_asn": 64999, "bgp_cer_cidr": "10.254.30.78/30", "bgp_ibm_asn": 13884, "bgp_ibm_cidr": "10.254.30.77/30", "bgp_status": "active", "change_request": {"type": "create_gateway"}, "created_at": "2019-01-01T12:00:00.000Z", "crn": "crn:v1:bluemix:public:directlink:dal03:a/4111d05f36894e3cb9b46a43556d9000::connect:ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4", "customer_account_id": "4111d05f36894e3cb9b46a43556d9000", "id": "ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4", "name": "myGateway", "operational_status": "configuring", "port": {"id": "fffdcb1a-fee4-41c7-9e11-9cd99e65c777"}, "provider_api_managed": true, "speed_mbps": 1000, "type": "connect", "vlan": 10}`)
				}))
			})
			It(`Invoke GetProviderGateway successfully with retries`, func() {
				directLinkProviderService, serviceErr := directlinkproviderv2.NewDirectLinkProviderV2(&directlinkproviderv2.DirectLinkProviderV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkProviderService).ToNot(BeNil())
				directLinkProviderService.EnableRetries(0, 0)

				// Construct an instance of the GetProviderGatewayOptions model
				getProviderGatewayOptionsModel := new(directlinkproviderv2.GetProviderGatewayOptions)
				getProviderGatewayOptionsModel.ID = core.StringPtr("testString")
				getProviderGatewayOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := directLinkProviderService.GetProviderGatewayWithContext(ctx, getProviderGatewayOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				directLinkProviderService.DisableRetries()
				result, response, operationErr := directLinkProviderService.GetProviderGateway(getProviderGatewayOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = directLinkProviderService.GetProviderGatewayWithContext(ctx, getProviderGatewayOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getProviderGatewayPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"bgp_asn": 64999, "bgp_cer_cidr": "10.254.30.78/30", "bgp_ibm_asn": 13884, "bgp_ibm_cidr": "10.254.30.77/30", "bgp_status": "active", "change_request": {"type": "create_gateway"}, "created_at": "2019-01-01T12:00:00.000Z", "crn": "crn:v1:bluemix:public:directlink:dal03:a/4111d05f36894e3cb9b46a43556d9000::connect:ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4", "customer_account_id": "4111d05f36894e3cb9b46a43556d9000", "id": "ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4", "name": "myGateway", "operational_status": "configuring", "port": {"id": "fffdcb1a-fee4-41c7-9e11-9cd99e65c777"}, "provider_api_managed": true, "speed_mbps": 1000, "type": "connect", "vlan": 10}`)
				}))
			})
			It(`Invoke GetProviderGateway successfully`, func() {
				directLinkProviderService, serviceErr := directlinkproviderv2.NewDirectLinkProviderV2(&directlinkproviderv2.DirectLinkProviderV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkProviderService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := directLinkProviderService.GetProviderGateway(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetProviderGatewayOptions model
				getProviderGatewayOptionsModel := new(directlinkproviderv2.GetProviderGatewayOptions)
				getProviderGatewayOptionsModel.ID = core.StringPtr("testString")
				getProviderGatewayOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = directLinkProviderService.GetProviderGateway(getProviderGatewayOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetProviderGateway with error: Operation validation and request error`, func() {
				directLinkProviderService, serviceErr := directlinkproviderv2.NewDirectLinkProviderV2(&directlinkproviderv2.DirectLinkProviderV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkProviderService).ToNot(BeNil())

				// Construct an instance of the GetProviderGatewayOptions model
				getProviderGatewayOptionsModel := new(directlinkproviderv2.GetProviderGatewayOptions)
				getProviderGatewayOptionsModel.ID = core.StringPtr("testString")
				getProviderGatewayOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := directLinkProviderService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := directLinkProviderService.GetProviderGateway(getProviderGatewayOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetProviderGatewayOptions model with no property values
				getProviderGatewayOptionsModelNew := new(directlinkproviderv2.GetProviderGatewayOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = directLinkProviderService.GetProviderGateway(getProviderGatewayOptionsModelNew)
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
			It(`Invoke GetProviderGateway successfully`, func() {
				directLinkProviderService, serviceErr := directlinkproviderv2.NewDirectLinkProviderV2(&directlinkproviderv2.DirectLinkProviderV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkProviderService).ToNot(BeNil())

				// Construct an instance of the GetProviderGatewayOptions model
				getProviderGatewayOptionsModel := new(directlinkproviderv2.GetProviderGatewayOptions)
				getProviderGatewayOptionsModel.ID = core.StringPtr("testString")
				getProviderGatewayOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := directLinkProviderService.GetProviderGateway(getProviderGatewayOptionsModel)
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
	Describe(`UpdateProviderGateway(updateProviderGatewayOptions *UpdateProviderGatewayOptions) - Operation response error`, func() {
		version := "testString"
		updateProviderGatewayPath := "/gateways/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateProviderGatewayPath))
					Expect(req.Method).To(Equal("PATCH"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateProviderGateway with error: Operation response processing error`, func() {
				directLinkProviderService, serviceErr := directlinkproviderv2.NewDirectLinkProviderV2(&directlinkproviderv2.DirectLinkProviderV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkProviderService).ToNot(BeNil())

				// Construct an instance of the UpdateProviderGatewayOptions model
				updateProviderGatewayOptionsModel := new(directlinkproviderv2.UpdateProviderGatewayOptions)
				updateProviderGatewayOptionsModel.ID = core.StringPtr("testString")
				updateProviderGatewayOptionsModel.BgpAsn = core.Int64Ptr(int64(64999))
				updateProviderGatewayOptionsModel.BgpCerCidr = core.StringPtr("169.254.0.10/30")
				updateProviderGatewayOptionsModel.BgpIbmCidr = core.StringPtr("169.254.0.9/30")
				updateProviderGatewayOptionsModel.Name = core.StringPtr("myNewGateway")
				updateProviderGatewayOptionsModel.SpeedMbps = core.Int64Ptr(int64(1000))
				updateProviderGatewayOptionsModel.Vlan = core.Int64Ptr(int64(10))
				updateProviderGatewayOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := directLinkProviderService.UpdateProviderGateway(updateProviderGatewayOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				directLinkProviderService.EnableRetries(0, 0)
				result, response, operationErr = directLinkProviderService.UpdateProviderGateway(updateProviderGatewayOptionsModel)
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
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateProviderGatewayPath))
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
					fmt.Fprintf(res, "%s", `{"bgp_asn": 64999, "bgp_cer_cidr": "10.254.30.78/30", "bgp_ibm_asn": 13884, "bgp_ibm_cidr": "10.254.30.77/30", "bgp_status": "active", "change_request": {"type": "create_gateway"}, "created_at": "2019-01-01T12:00:00.000Z", "crn": "crn:v1:bluemix:public:directlink:dal03:a/4111d05f36894e3cb9b46a43556d9000::connect:ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4", "customer_account_id": "4111d05f36894e3cb9b46a43556d9000", "id": "ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4", "name": "myGateway", "operational_status": "configuring", "port": {"id": "fffdcb1a-fee4-41c7-9e11-9cd99e65c777"}, "provider_api_managed": true, "speed_mbps": 1000, "type": "connect", "vlan": 10}`)
				}))
			})
			It(`Invoke UpdateProviderGateway successfully with retries`, func() {
				directLinkProviderService, serviceErr := directlinkproviderv2.NewDirectLinkProviderV2(&directlinkproviderv2.DirectLinkProviderV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkProviderService).ToNot(BeNil())
				directLinkProviderService.EnableRetries(0, 0)

				// Construct an instance of the UpdateProviderGatewayOptions model
				updateProviderGatewayOptionsModel := new(directlinkproviderv2.UpdateProviderGatewayOptions)
				updateProviderGatewayOptionsModel.ID = core.StringPtr("testString")
				updateProviderGatewayOptionsModel.BgpAsn = core.Int64Ptr(int64(64999))
				updateProviderGatewayOptionsModel.BgpCerCidr = core.StringPtr("169.254.0.10/30")
				updateProviderGatewayOptionsModel.BgpIbmCidr = core.StringPtr("169.254.0.9/30")
				updateProviderGatewayOptionsModel.Name = core.StringPtr("myNewGateway")
				updateProviderGatewayOptionsModel.SpeedMbps = core.Int64Ptr(int64(1000))
				updateProviderGatewayOptionsModel.Vlan = core.Int64Ptr(int64(10))
				updateProviderGatewayOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := directLinkProviderService.UpdateProviderGatewayWithContext(ctx, updateProviderGatewayOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				directLinkProviderService.DisableRetries()
				result, response, operationErr := directLinkProviderService.UpdateProviderGateway(updateProviderGatewayOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = directLinkProviderService.UpdateProviderGatewayWithContext(ctx, updateProviderGatewayOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(updateProviderGatewayPath))
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
					fmt.Fprintf(res, "%s", `{"bgp_asn": 64999, "bgp_cer_cidr": "10.254.30.78/30", "bgp_ibm_asn": 13884, "bgp_ibm_cidr": "10.254.30.77/30", "bgp_status": "active", "change_request": {"type": "create_gateway"}, "created_at": "2019-01-01T12:00:00.000Z", "crn": "crn:v1:bluemix:public:directlink:dal03:a/4111d05f36894e3cb9b46a43556d9000::connect:ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4", "customer_account_id": "4111d05f36894e3cb9b46a43556d9000", "id": "ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4", "name": "myGateway", "operational_status": "configuring", "port": {"id": "fffdcb1a-fee4-41c7-9e11-9cd99e65c777"}, "provider_api_managed": true, "speed_mbps": 1000, "type": "connect", "vlan": 10}`)
				}))
			})
			It(`Invoke UpdateProviderGateway successfully`, func() {
				directLinkProviderService, serviceErr := directlinkproviderv2.NewDirectLinkProviderV2(&directlinkproviderv2.DirectLinkProviderV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkProviderService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := directLinkProviderService.UpdateProviderGateway(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateProviderGatewayOptions model
				updateProviderGatewayOptionsModel := new(directlinkproviderv2.UpdateProviderGatewayOptions)
				updateProviderGatewayOptionsModel.ID = core.StringPtr("testString")
				updateProviderGatewayOptionsModel.BgpAsn = core.Int64Ptr(int64(64999))
				updateProviderGatewayOptionsModel.BgpCerCidr = core.StringPtr("169.254.0.10/30")
				updateProviderGatewayOptionsModel.BgpIbmCidr = core.StringPtr("169.254.0.9/30")
				updateProviderGatewayOptionsModel.Name = core.StringPtr("myNewGateway")
				updateProviderGatewayOptionsModel.SpeedMbps = core.Int64Ptr(int64(1000))
				updateProviderGatewayOptionsModel.Vlan = core.Int64Ptr(int64(10))
				updateProviderGatewayOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = directLinkProviderService.UpdateProviderGateway(updateProviderGatewayOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateProviderGateway with error: Operation validation and request error`, func() {
				directLinkProviderService, serviceErr := directlinkproviderv2.NewDirectLinkProviderV2(&directlinkproviderv2.DirectLinkProviderV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkProviderService).ToNot(BeNil())

				// Construct an instance of the UpdateProviderGatewayOptions model
				updateProviderGatewayOptionsModel := new(directlinkproviderv2.UpdateProviderGatewayOptions)
				updateProviderGatewayOptionsModel.ID = core.StringPtr("testString")
				updateProviderGatewayOptionsModel.BgpAsn = core.Int64Ptr(int64(64999))
				updateProviderGatewayOptionsModel.BgpCerCidr = core.StringPtr("169.254.0.10/30")
				updateProviderGatewayOptionsModel.BgpIbmCidr = core.StringPtr("169.254.0.9/30")
				updateProviderGatewayOptionsModel.Name = core.StringPtr("myNewGateway")
				updateProviderGatewayOptionsModel.SpeedMbps = core.Int64Ptr(int64(1000))
				updateProviderGatewayOptionsModel.Vlan = core.Int64Ptr(int64(10))
				updateProviderGatewayOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := directLinkProviderService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := directLinkProviderService.UpdateProviderGateway(updateProviderGatewayOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateProviderGatewayOptions model with no property values
				updateProviderGatewayOptionsModelNew := new(directlinkproviderv2.UpdateProviderGatewayOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = directLinkProviderService.UpdateProviderGateway(updateProviderGatewayOptionsModelNew)
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
			It(`Invoke UpdateProviderGateway successfully`, func() {
				directLinkProviderService, serviceErr := directlinkproviderv2.NewDirectLinkProviderV2(&directlinkproviderv2.DirectLinkProviderV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkProviderService).ToNot(BeNil())

				// Construct an instance of the UpdateProviderGatewayOptions model
				updateProviderGatewayOptionsModel := new(directlinkproviderv2.UpdateProviderGatewayOptions)
				updateProviderGatewayOptionsModel.ID = core.StringPtr("testString")
				updateProviderGatewayOptionsModel.BgpAsn = core.Int64Ptr(int64(64999))
				updateProviderGatewayOptionsModel.BgpCerCidr = core.StringPtr("169.254.0.10/30")
				updateProviderGatewayOptionsModel.BgpIbmCidr = core.StringPtr("169.254.0.9/30")
				updateProviderGatewayOptionsModel.Name = core.StringPtr("myNewGateway")
				updateProviderGatewayOptionsModel.SpeedMbps = core.Int64Ptr(int64(1000))
				updateProviderGatewayOptionsModel.Vlan = core.Int64Ptr(int64(10))
				updateProviderGatewayOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := directLinkProviderService.UpdateProviderGateway(updateProviderGatewayOptionsModel)
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
	Describe(`ListProviderPorts(listProviderPortsOptions *ListProviderPortsOptions) - Operation response error`, func() {
		version := "testString"
		listProviderPortsPath := "/ports"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listProviderPortsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(1))}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListProviderPorts with error: Operation response processing error`, func() {
				directLinkProviderService, serviceErr := directlinkproviderv2.NewDirectLinkProviderV2(&directlinkproviderv2.DirectLinkProviderV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkProviderService).ToNot(BeNil())

				// Construct an instance of the ListProviderPortsOptions model
				listProviderPortsOptionsModel := new(directlinkproviderv2.ListProviderPortsOptions)
				listProviderPortsOptionsModel.Start = core.StringPtr("testString")
				listProviderPortsOptionsModel.Limit = core.Int64Ptr(int64(1))
				listProviderPortsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := directLinkProviderService.ListProviderPorts(listProviderPortsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				directLinkProviderService.EnableRetries(0, 0)
				result, response, operationErr = directLinkProviderService.ListProviderPorts(listProviderPortsOptionsModel)
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
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listProviderPortsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(1))}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"first": {"href": "https://directlink.cloud.ibm.com/provider/v2/ports?limit=100"}, "limit": 100, "next": {"href": "https://directlink.cloud.ibm.com/provider/v2/ports?start=9d5a91a3e2cbd233b5a5b33436855ed1&limit=100", "start": "9d5a91a3e2cbd233b5a5b33436855ed1"}, "total_count": 132, "ports": [{"id": "01122b9b-820f-4c44-8a31-77f1f0806765", "label": "XCR-FRK-CS-SEC-01", "location_display_name": "Dallas 03", "location_name": "dal03", "provider_name": "provider_1", "supported_link_speeds": [19]}]}`)
				}))
			})
			It(`Invoke ListProviderPorts successfully with retries`, func() {
				directLinkProviderService, serviceErr := directlinkproviderv2.NewDirectLinkProviderV2(&directlinkproviderv2.DirectLinkProviderV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkProviderService).ToNot(BeNil())
				directLinkProviderService.EnableRetries(0, 0)

				// Construct an instance of the ListProviderPortsOptions model
				listProviderPortsOptionsModel := new(directlinkproviderv2.ListProviderPortsOptions)
				listProviderPortsOptionsModel.Start = core.StringPtr("testString")
				listProviderPortsOptionsModel.Limit = core.Int64Ptr(int64(1))
				listProviderPortsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := directLinkProviderService.ListProviderPortsWithContext(ctx, listProviderPortsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				directLinkProviderService.DisableRetries()
				result, response, operationErr := directLinkProviderService.ListProviderPorts(listProviderPortsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = directLinkProviderService.ListProviderPortsWithContext(ctx, listProviderPortsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listProviderPortsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(1))}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"first": {"href": "https://directlink.cloud.ibm.com/provider/v2/ports?limit=100"}, "limit": 100, "next": {"href": "https://directlink.cloud.ibm.com/provider/v2/ports?start=9d5a91a3e2cbd233b5a5b33436855ed1&limit=100", "start": "9d5a91a3e2cbd233b5a5b33436855ed1"}, "total_count": 132, "ports": [{"id": "01122b9b-820f-4c44-8a31-77f1f0806765", "label": "XCR-FRK-CS-SEC-01", "location_display_name": "Dallas 03", "location_name": "dal03", "provider_name": "provider_1", "supported_link_speeds": [19]}]}`)
				}))
			})
			It(`Invoke ListProviderPorts successfully`, func() {
				directLinkProviderService, serviceErr := directlinkproviderv2.NewDirectLinkProviderV2(&directlinkproviderv2.DirectLinkProviderV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkProviderService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := directLinkProviderService.ListProviderPorts(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListProviderPortsOptions model
				listProviderPortsOptionsModel := new(directlinkproviderv2.ListProviderPortsOptions)
				listProviderPortsOptionsModel.Start = core.StringPtr("testString")
				listProviderPortsOptionsModel.Limit = core.Int64Ptr(int64(1))
				listProviderPortsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = directLinkProviderService.ListProviderPorts(listProviderPortsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListProviderPorts with error: Operation request error`, func() {
				directLinkProviderService, serviceErr := directlinkproviderv2.NewDirectLinkProviderV2(&directlinkproviderv2.DirectLinkProviderV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkProviderService).ToNot(BeNil())

				// Construct an instance of the ListProviderPortsOptions model
				listProviderPortsOptionsModel := new(directlinkproviderv2.ListProviderPortsOptions)
				listProviderPortsOptionsModel.Start = core.StringPtr("testString")
				listProviderPortsOptionsModel.Limit = core.Int64Ptr(int64(1))
				listProviderPortsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := directLinkProviderService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := directLinkProviderService.ListProviderPorts(listProviderPortsOptionsModel)
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
			It(`Invoke ListProviderPorts successfully`, func() {
				directLinkProviderService, serviceErr := directlinkproviderv2.NewDirectLinkProviderV2(&directlinkproviderv2.DirectLinkProviderV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkProviderService).ToNot(BeNil())

				// Construct an instance of the ListProviderPortsOptions model
				listProviderPortsOptionsModel := new(directlinkproviderv2.ListProviderPortsOptions)
				listProviderPortsOptionsModel.Start = core.StringPtr("testString")
				listProviderPortsOptionsModel.Limit = core.Int64Ptr(int64(1))
				listProviderPortsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := directLinkProviderService.ListProviderPorts(listProviderPortsOptionsModel)
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
				responseObject := new(directlinkproviderv2.ProviderPortCollection)
				nextObject := new(directlinkproviderv2.ProviderPortCollectionNext)
				nextObject.Start = core.StringPtr("abc-123")
				responseObject.Next = nextObject

				value, err := responseObject.GetNextStart()
				Expect(err).To(BeNil())
				Expect(value).To(Equal(core.StringPtr("abc-123")))
			})
			It(`Invoke GetNextStart without a "Next" property in the response`, func() {
				responseObject := new(directlinkproviderv2.ProviderPortCollection)

				value, err := responseObject.GetNextStart()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
		})
	})
	Describe(`GetProviderPort(getProviderPortOptions *GetProviderPortOptions) - Operation response error`, func() {
		version := "testString"
		getProviderPortPath := "/ports/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getProviderPortPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetProviderPort with error: Operation response processing error`, func() {
				directLinkProviderService, serviceErr := directlinkproviderv2.NewDirectLinkProviderV2(&directlinkproviderv2.DirectLinkProviderV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkProviderService).ToNot(BeNil())

				// Construct an instance of the GetProviderPortOptions model
				getProviderPortOptionsModel := new(directlinkproviderv2.GetProviderPortOptions)
				getProviderPortOptionsModel.ID = core.StringPtr("testString")
				getProviderPortOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := directLinkProviderService.GetProviderPort(getProviderPortOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				directLinkProviderService.EnableRetries(0, 0)
				result, response, operationErr = directLinkProviderService.GetProviderPort(getProviderPortOptionsModel)
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
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getProviderPortPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "01122b9b-820f-4c44-8a31-77f1f0806765", "label": "XCR-FRK-CS-SEC-01", "location_display_name": "Dallas 03", "location_name": "dal03", "provider_name": "provider_1", "supported_link_speeds": [19]}`)
				}))
			})
			It(`Invoke GetProviderPort successfully with retries`, func() {
				directLinkProviderService, serviceErr := directlinkproviderv2.NewDirectLinkProviderV2(&directlinkproviderv2.DirectLinkProviderV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkProviderService).ToNot(BeNil())
				directLinkProviderService.EnableRetries(0, 0)

				// Construct an instance of the GetProviderPortOptions model
				getProviderPortOptionsModel := new(directlinkproviderv2.GetProviderPortOptions)
				getProviderPortOptionsModel.ID = core.StringPtr("testString")
				getProviderPortOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := directLinkProviderService.GetProviderPortWithContext(ctx, getProviderPortOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				directLinkProviderService.DisableRetries()
				result, response, operationErr := directLinkProviderService.GetProviderPort(getProviderPortOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = directLinkProviderService.GetProviderPortWithContext(ctx, getProviderPortOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getProviderPortPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "01122b9b-820f-4c44-8a31-77f1f0806765", "label": "XCR-FRK-CS-SEC-01", "location_display_name": "Dallas 03", "location_name": "dal03", "provider_name": "provider_1", "supported_link_speeds": [19]}`)
				}))
			})
			It(`Invoke GetProviderPort successfully`, func() {
				directLinkProviderService, serviceErr := directlinkproviderv2.NewDirectLinkProviderV2(&directlinkproviderv2.DirectLinkProviderV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkProviderService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := directLinkProviderService.GetProviderPort(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetProviderPortOptions model
				getProviderPortOptionsModel := new(directlinkproviderv2.GetProviderPortOptions)
				getProviderPortOptionsModel.ID = core.StringPtr("testString")
				getProviderPortOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = directLinkProviderService.GetProviderPort(getProviderPortOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetProviderPort with error: Operation validation and request error`, func() {
				directLinkProviderService, serviceErr := directlinkproviderv2.NewDirectLinkProviderV2(&directlinkproviderv2.DirectLinkProviderV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkProviderService).ToNot(BeNil())

				// Construct an instance of the GetProviderPortOptions model
				getProviderPortOptionsModel := new(directlinkproviderv2.GetProviderPortOptions)
				getProviderPortOptionsModel.ID = core.StringPtr("testString")
				getProviderPortOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := directLinkProviderService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := directLinkProviderService.GetProviderPort(getProviderPortOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetProviderPortOptions model with no property values
				getProviderPortOptionsModelNew := new(directlinkproviderv2.GetProviderPortOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = directLinkProviderService.GetProviderPort(getProviderPortOptionsModelNew)
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
			It(`Invoke GetProviderPort successfully`, func() {
				directLinkProviderService, serviceErr := directlinkproviderv2.NewDirectLinkProviderV2(&directlinkproviderv2.DirectLinkProviderV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(directLinkProviderService).ToNot(BeNil())

				// Construct an instance of the GetProviderPortOptions model
				getProviderPortOptionsModel := new(directlinkproviderv2.GetProviderPortOptions)
				getProviderPortOptionsModel.ID = core.StringPtr("testString")
				getProviderPortOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := directLinkProviderService.GetProviderPort(getProviderPortOptionsModel)
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
			directLinkProviderService, _ := directlinkproviderv2.NewDirectLinkProviderV2(&directlinkproviderv2.DirectLinkProviderV2Options{
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
				createProviderGatewayOptionsCustomerAccountID := "4111d05f36894e3cb9b46a43556d9000"
				createProviderGatewayOptionsName := "myGateway"
				var createProviderGatewayOptionsPort *directlinkproviderv2.ProviderGatewayPortIdentity = nil
				createProviderGatewayOptionsSpeedMbps := int64(1000)
				createProviderGatewayOptionsModel := directLinkProviderService.NewCreateProviderGatewayOptions(createProviderGatewayOptionsBgpAsn, createProviderGatewayOptionsCustomerAccountID, createProviderGatewayOptionsName, createProviderGatewayOptionsPort, createProviderGatewayOptionsSpeedMbps)
				createProviderGatewayOptionsModel.SetBgpAsn(int64(64999))
				createProviderGatewayOptionsModel.SetCustomerAccountID("4111d05f36894e3cb9b46a43556d9000")
				createProviderGatewayOptionsModel.SetName("myGateway")
				createProviderGatewayOptionsModel.SetPort(providerGatewayPortIdentityModel)
				createProviderGatewayOptionsModel.SetSpeedMbps(int64(1000))
				createProviderGatewayOptionsModel.SetBgpCerCidr("10.254.30.78/30")
				createProviderGatewayOptionsModel.SetBgpIbmCidr("10.254.30.77/30")
				createProviderGatewayOptionsModel.SetVlan(int64(10))
				createProviderGatewayOptionsModel.SetCheckOnly("testString")
				createProviderGatewayOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createProviderGatewayOptionsModel).ToNot(BeNil())
				Expect(createProviderGatewayOptionsModel.BgpAsn).To(Equal(core.Int64Ptr(int64(64999))))
				Expect(createProviderGatewayOptionsModel.CustomerAccountID).To(Equal(core.StringPtr("4111d05f36894e3cb9b46a43556d9000")))
				Expect(createProviderGatewayOptionsModel.Name).To(Equal(core.StringPtr("myGateway")))
				Expect(createProviderGatewayOptionsModel.Port).To(Equal(providerGatewayPortIdentityModel))
				Expect(createProviderGatewayOptionsModel.SpeedMbps).To(Equal(core.Int64Ptr(int64(1000))))
				Expect(createProviderGatewayOptionsModel.BgpCerCidr).To(Equal(core.StringPtr("10.254.30.78/30")))
				Expect(createProviderGatewayOptionsModel.BgpIbmCidr).To(Equal(core.StringPtr("10.254.30.77/30")))
				Expect(createProviderGatewayOptionsModel.Vlan).To(Equal(core.Int64Ptr(int64(10))))
				Expect(createProviderGatewayOptionsModel.CheckOnly).To(Equal(core.StringPtr("testString")))
				Expect(createProviderGatewayOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteProviderGatewayOptions successfully`, func() {
				// Construct an instance of the DeleteProviderGatewayOptions model
				id := "testString"
				deleteProviderGatewayOptionsModel := directLinkProviderService.NewDeleteProviderGatewayOptions(id)
				deleteProviderGatewayOptionsModel.SetID("testString")
				deleteProviderGatewayOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteProviderGatewayOptionsModel).ToNot(BeNil())
				Expect(deleteProviderGatewayOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(deleteProviderGatewayOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetProviderGatewayOptions successfully`, func() {
				// Construct an instance of the GetProviderGatewayOptions model
				id := "testString"
				getProviderGatewayOptionsModel := directLinkProviderService.NewGetProviderGatewayOptions(id)
				getProviderGatewayOptionsModel.SetID("testString")
				getProviderGatewayOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getProviderGatewayOptionsModel).ToNot(BeNil())
				Expect(getProviderGatewayOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(getProviderGatewayOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetProviderPortOptions successfully`, func() {
				// Construct an instance of the GetProviderPortOptions model
				id := "testString"
				getProviderPortOptionsModel := directLinkProviderService.NewGetProviderPortOptions(id)
				getProviderPortOptionsModel.SetID("testString")
				getProviderPortOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getProviderPortOptionsModel).ToNot(BeNil())
				Expect(getProviderPortOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(getProviderPortOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListProviderGatewaysOptions successfully`, func() {
				// Construct an instance of the ListProviderGatewaysOptions model
				listProviderGatewaysOptionsModel := directLinkProviderService.NewListProviderGatewaysOptions()
				listProviderGatewaysOptionsModel.SetStart("testString")
				listProviderGatewaysOptionsModel.SetLimit(int64(1))
				listProviderGatewaysOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listProviderGatewaysOptionsModel).ToNot(BeNil())
				Expect(listProviderGatewaysOptionsModel.Start).To(Equal(core.StringPtr("testString")))
				Expect(listProviderGatewaysOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(1))))
				Expect(listProviderGatewaysOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListProviderPortsOptions successfully`, func() {
				// Construct an instance of the ListProviderPortsOptions model
				listProviderPortsOptionsModel := directLinkProviderService.NewListProviderPortsOptions()
				listProviderPortsOptionsModel.SetStart("testString")
				listProviderPortsOptionsModel.SetLimit(int64(1))
				listProviderPortsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listProviderPortsOptionsModel).ToNot(BeNil())
				Expect(listProviderPortsOptionsModel.Start).To(Equal(core.StringPtr("testString")))
				Expect(listProviderPortsOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(1))))
				Expect(listProviderPortsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewProviderGatewayPortIdentity successfully`, func() {
				id := "fffdcb1a-fee4-41c7-9e11-9cd99e65c777"
				_model, err := directLinkProviderService.NewProviderGatewayPortIdentity(id)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewUpdateProviderGatewayOptions successfully`, func() {
				// Construct an instance of the UpdateProviderGatewayOptions model
				id := "testString"
				updateProviderGatewayOptionsModel := directLinkProviderService.NewUpdateProviderGatewayOptions(id)
				updateProviderGatewayOptionsModel.SetID("testString")
				updateProviderGatewayOptionsModel.SetBgpAsn(int64(64999))
				updateProviderGatewayOptionsModel.SetBgpCerCidr("169.254.0.10/30")
				updateProviderGatewayOptionsModel.SetBgpIbmCidr("169.254.0.9/30")
				updateProviderGatewayOptionsModel.SetName("myNewGateway")
				updateProviderGatewayOptionsModel.SetSpeedMbps(int64(1000))
				updateProviderGatewayOptionsModel.SetVlan(int64(10))
				updateProviderGatewayOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateProviderGatewayOptionsModel).ToNot(BeNil())
				Expect(updateProviderGatewayOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(updateProviderGatewayOptionsModel.BgpAsn).To(Equal(core.Int64Ptr(int64(64999))))
				Expect(updateProviderGatewayOptionsModel.BgpCerCidr).To(Equal(core.StringPtr("169.254.0.10/30")))
				Expect(updateProviderGatewayOptionsModel.BgpIbmCidr).To(Equal(core.StringPtr("169.254.0.9/30")))
				Expect(updateProviderGatewayOptionsModel.Name).To(Equal(core.StringPtr("myNewGateway")))
				Expect(updateProviderGatewayOptionsModel.SpeedMbps).To(Equal(core.Int64Ptr(int64(1000))))
				Expect(updateProviderGatewayOptionsModel.Vlan).To(Equal(core.Int64Ptr(int64(10))))
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
