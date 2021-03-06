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

package transitgatewayapisv1_test

import (
	"bytes"
	"fmt"
	"github.com/IBM/go-sdk-core/v4/core"
	"github.com/IBM/networking-go-sdk/transitgatewayapisv1"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"time"
)

var _ = Describe(`TransitGatewayApisV1`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		version := "testString"
		It(`Instantiate service client`, func() {
			transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
				Version: core.StringPtr(version),
			})
			Expect(transitGatewayApisService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
				URL: "{BAD_URL_STRING",
				Version: core.StringPtr(version),
			})
			Expect(transitGatewayApisService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
				URL: "https://transitgatewayapisv1/api",
				Version: core.StringPtr(version),
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(transitGatewayApisService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Validation Error`, func() {
			transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{})
			Expect(transitGatewayApisService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		version := "testString"
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"TRANSIT_GATEWAY_APIS_URL": "https://transitgatewayapisv1/api",
				"TRANSIT_GATEWAY_APIS_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1UsingExternalConfig(&transitgatewayapisv1.TransitGatewayApisV1Options{
					Version: core.StringPtr(version),
				})
				Expect(transitGatewayApisService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1UsingExternalConfig(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL: "https://testService/api",
					Version: core.StringPtr(version),
				})
				Expect(transitGatewayApisService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(transitGatewayApisService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1UsingExternalConfig(&transitgatewayapisv1.TransitGatewayApisV1Options{
					Version: core.StringPtr(version),
				})
				err := transitGatewayApisService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(transitGatewayApisService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(transitGatewayApisService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"TRANSIT_GATEWAY_APIS_URL": "https://transitgatewayapisv1/api",
				"TRANSIT_GATEWAY_APIS_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1UsingExternalConfig(&transitgatewayapisv1.TransitGatewayApisV1Options{
				Version: core.StringPtr(version),
			})

			It(`Instantiate service client with error`, func() {
				Expect(transitGatewayApisService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"TRANSIT_GATEWAY_APIS_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1UsingExternalConfig(&transitgatewayapisv1.TransitGatewayApisV1Options{
				URL: "{BAD_URL_STRING",
				Version: core.StringPtr(version),
			})

			It(`Instantiate service client with error`, func() {
				Expect(transitGatewayApisService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`ListConnections(listConnectionsOptions *ListConnectionsOptions) - Operation response error`, func() {
		version := "testString"
		listConnectionsPath := "/connections"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(listConnectionsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(1))}))

					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["network_id"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListConnections with error: Operation response processing error`, func() {
				transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(transitGatewayApisService).ToNot(BeNil())

				// Construct an instance of the ListConnectionsOptions model
				listConnectionsOptionsModel := new(transitgatewayapisv1.ListConnectionsOptions)
				listConnectionsOptionsModel.Limit = core.Int64Ptr(int64(1))
				listConnectionsOptionsModel.Start = core.StringPtr("testString")
				listConnectionsOptionsModel.NetworkID = core.StringPtr("testString")
				listConnectionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := transitGatewayApisService.ListConnections(listConnectionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ListConnections(listConnectionsOptions *ListConnectionsOptions)`, func() {
		version := "testString"
		listConnectionsPath := "/connections"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(listConnectionsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(1))}))

					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["network_id"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"connections": [{"base_connection_id": "975f58c1-afe7-469a-9727-7f3d720f2d32", "created_at": "2019-01-01T12:00:00", "id": "1a15dca5-7e33-45e1-b7c5-bc690e569531", "local_bgp_asn": 64490, "local_gateway_ip": "192.168.100.1", "local_tunnel_ip": "192.168.129.2", "mtu": 9000, "name": "Transit_Service_SJ_DL", "network_account_id": "28e4d90ac7504be694471ee66e70d0d5", "network_id": "crn:v1:bluemix:public:is:us-south:a/123456::vpc:4727d842-f94f-4a2d-824a-9bc9b02c523b", "network_type": "vpc", "remote_bgp_asn": 65010, "remote_gateway_ip": "10.242.63.12", "remote_tunnel_ip": "192.168.129.1", "request_status": "pending", "status": "attached", "transit_gateway": {"crn": "crn:v1:bluemix:public:transit:us-south:a/123456::gateway:456f58c1-afe7-123a-0a0a-7f3d720f1a44", "id": "456f58c1-afe7-123a-0a0a-7f3d720f1a44", "name": "my-transit-gw100"}, "updated_at": "2019-01-01T12:00:00", "zone": {"name": "us-south-1"}}], "first": {"href": "https://transit.cloud.ibm.com/v1/connections?limit=50"}, "limit": 50, "next": {"href": "https://transit.cloud.ibm.com/v1/connections?start=MjAyMC0wNS0wOVQxNjoyMDoyMC4yMjQ5NzNa&limit=50", "start": "MjAyMC0wNS0wOVQxNjoyMDoyMC4yMjQ5NzNa"}}`)
				}))
			})
			It(`Invoke ListConnections successfully`, func() {
				transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(transitGatewayApisService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := transitGatewayApisService.ListConnections(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListConnectionsOptions model
				listConnectionsOptionsModel := new(transitgatewayapisv1.ListConnectionsOptions)
				listConnectionsOptionsModel.Limit = core.Int64Ptr(int64(1))
				listConnectionsOptionsModel.Start = core.StringPtr("testString")
				listConnectionsOptionsModel.NetworkID = core.StringPtr("testString")
				listConnectionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = transitGatewayApisService.ListConnections(listConnectionsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke ListConnections with error: Operation request error`, func() {
				transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(transitGatewayApisService).ToNot(BeNil())

				// Construct an instance of the ListConnectionsOptions model
				listConnectionsOptionsModel := new(transitgatewayapisv1.ListConnectionsOptions)
				listConnectionsOptionsModel.Limit = core.Int64Ptr(int64(1))
				listConnectionsOptionsModel.Start = core.StringPtr("testString")
				listConnectionsOptionsModel.NetworkID = core.StringPtr("testString")
				listConnectionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := transitGatewayApisService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := transitGatewayApisService.ListConnections(listConnectionsOptionsModel)
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
	Describe(`Service constructor tests`, func() {
		version := "testString"
		It(`Instantiate service client`, func() {
			transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
				Version: core.StringPtr(version),
			})
			Expect(transitGatewayApisService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
				URL: "{BAD_URL_STRING",
				Version: core.StringPtr(version),
			})
			Expect(transitGatewayApisService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
				URL: "https://transitgatewayapisv1/api",
				Version: core.StringPtr(version),
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(transitGatewayApisService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Validation Error`, func() {
			transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{})
			Expect(transitGatewayApisService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		version := "testString"
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"TRANSIT_GATEWAY_APIS_URL": "https://transitgatewayapisv1/api",
				"TRANSIT_GATEWAY_APIS_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1UsingExternalConfig(&transitgatewayapisv1.TransitGatewayApisV1Options{
					Version: core.StringPtr(version),
				})
				Expect(transitGatewayApisService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1UsingExternalConfig(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL: "https://testService/api",
					Version: core.StringPtr(version),
				})
				Expect(transitGatewayApisService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(transitGatewayApisService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1UsingExternalConfig(&transitgatewayapisv1.TransitGatewayApisV1Options{
					Version: core.StringPtr(version),
				})
				err := transitGatewayApisService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(transitGatewayApisService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(transitGatewayApisService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"TRANSIT_GATEWAY_APIS_URL": "https://transitgatewayapisv1/api",
				"TRANSIT_GATEWAY_APIS_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1UsingExternalConfig(&transitgatewayapisv1.TransitGatewayApisV1Options{
				Version: core.StringPtr(version),
			})

			It(`Instantiate service client with error`, func() {
				Expect(transitGatewayApisService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"TRANSIT_GATEWAY_APIS_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1UsingExternalConfig(&transitgatewayapisv1.TransitGatewayApisV1Options{
				URL: "{BAD_URL_STRING",
				Version: core.StringPtr(version),
			})

			It(`Instantiate service client with error`, func() {
				Expect(transitGatewayApisService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`ListTransitGateways(listTransitGatewaysOptions *ListTransitGatewaysOptions) - Operation response error`, func() {
		version := "testString"
		listTransitGatewaysPath := "/transit_gateways"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(listTransitGatewaysPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(1))}))

					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListTransitGateways with error: Operation response processing error`, func() {
				transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(transitGatewayApisService).ToNot(BeNil())

				// Construct an instance of the ListTransitGatewaysOptions model
				listTransitGatewaysOptionsModel := new(transitgatewayapisv1.ListTransitGatewaysOptions)
				listTransitGatewaysOptionsModel.Limit = core.Int64Ptr(int64(1))
				listTransitGatewaysOptionsModel.Start = core.StringPtr("testString")
				listTransitGatewaysOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := transitGatewayApisService.ListTransitGateways(listTransitGatewaysOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ListTransitGateways(listTransitGatewaysOptions *ListTransitGatewaysOptions)`, func() {
		version := "testString"
		listTransitGatewaysPath := "/transit_gateways"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(listTransitGatewaysPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(1))}))

					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"first": {"href": "https://transit.cloud.ibm.com/v1/transit_gateways?limit=50"}, "limit": 50, "next": {"href": "https://transit.cloud.ibm.com/v1/transit_gateways?start=MjAyMC0wNS0wOFQxNDoxNzowMy45NzQ5NzNa&limit=50", "start": "MjAyMC0wNS0wOFQxNDoxNzowMy45NzQ5NzNa"}, "transit_gateways": [{"id": "ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4", "crn": "crn:v1:bluemix:public:transit:dal03:a/57a7d05f36894e3cb9b46a43556d903e::gateway:ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4", "name": "my-transit-gateway-in-TransitGateway", "location": "us-south", "created_at": "2019-01-01T12:00:00", "global": true, "resource_group": {"id": "56969d6043e9465c883cb9f7363e78e8", "href": "https://resource-manager.bluemix.net/v1/resource_groups/56969d6043e9465c883cb9f7363e78e8"}, "status": "available", "updated_at": "2019-01-01T12:00:00"}]}`)
				}))
			})
			It(`Invoke ListTransitGateways successfully`, func() {
				transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(transitGatewayApisService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := transitGatewayApisService.ListTransitGateways(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListTransitGatewaysOptions model
				listTransitGatewaysOptionsModel := new(transitgatewayapisv1.ListTransitGatewaysOptions)
				listTransitGatewaysOptionsModel.Limit = core.Int64Ptr(int64(1))
				listTransitGatewaysOptionsModel.Start = core.StringPtr("testString")
				listTransitGatewaysOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = transitGatewayApisService.ListTransitGateways(listTransitGatewaysOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke ListTransitGateways with error: Operation request error`, func() {
				transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(transitGatewayApisService).ToNot(BeNil())

				// Construct an instance of the ListTransitGatewaysOptions model
				listTransitGatewaysOptionsModel := new(transitgatewayapisv1.ListTransitGatewaysOptions)
				listTransitGatewaysOptionsModel.Limit = core.Int64Ptr(int64(1))
				listTransitGatewaysOptionsModel.Start = core.StringPtr("testString")
				listTransitGatewaysOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := transitGatewayApisService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := transitGatewayApisService.ListTransitGateways(listTransitGatewaysOptionsModel)
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
	Describe(`CreateTransitGateway(createTransitGatewayOptions *CreateTransitGatewayOptions) - Operation response error`, func() {
		version := "testString"
		createTransitGatewayPath := "/transit_gateways"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(createTransitGatewayPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateTransitGateway with error: Operation response processing error`, func() {
				transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(transitGatewayApisService).ToNot(BeNil())

				// Construct an instance of the ResourceGroupIdentity model
				resourceGroupIdentityModel := new(transitgatewayapisv1.ResourceGroupIdentity)
				resourceGroupIdentityModel.ID = core.StringPtr("56969d6043e9465c883cb9f7363e78e8")

				// Construct an instance of the CreateTransitGatewayOptions model
				createTransitGatewayOptionsModel := new(transitgatewayapisv1.CreateTransitGatewayOptions)
				createTransitGatewayOptionsModel.Location = core.StringPtr("us-south")
				createTransitGatewayOptionsModel.Name = core.StringPtr("Transit_Service_BWTN_SJ_DL")
				createTransitGatewayOptionsModel.Global = core.BoolPtr(true)
				createTransitGatewayOptionsModel.ResourceGroup = resourceGroupIdentityModel
				createTransitGatewayOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := transitGatewayApisService.CreateTransitGateway(createTransitGatewayOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`CreateTransitGateway(createTransitGatewayOptions *CreateTransitGatewayOptions)`, func() {
		version := "testString"
		createTransitGatewayPath := "/transit_gateways"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(createTransitGatewayPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, `{"id": "ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4", "crn": "crn:v1:bluemix:public:transit:dal03:a/57a7d05f36894e3cb9b46a43556d903e::gateway:ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4", "name": "my-transit-gateway-in-TransitGateway", "location": "us-south", "created_at": "2019-01-01T12:00:00", "global": true, "resource_group": {"id": "56969d6043e9465c883cb9f7363e78e8", "href": "https://resource-manager.bluemix.net/v1/resource_groups/56969d6043e9465c883cb9f7363e78e8"}, "status": "available", "updated_at": "2019-01-01T12:00:00"}`)
				}))
			})
			It(`Invoke CreateTransitGateway successfully`, func() {
				transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(transitGatewayApisService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := transitGatewayApisService.CreateTransitGateway(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ResourceGroupIdentity model
				resourceGroupIdentityModel := new(transitgatewayapisv1.ResourceGroupIdentity)
				resourceGroupIdentityModel.ID = core.StringPtr("56969d6043e9465c883cb9f7363e78e8")

				// Construct an instance of the CreateTransitGatewayOptions model
				createTransitGatewayOptionsModel := new(transitgatewayapisv1.CreateTransitGatewayOptions)
				createTransitGatewayOptionsModel.Location = core.StringPtr("us-south")
				createTransitGatewayOptionsModel.Name = core.StringPtr("Transit_Service_BWTN_SJ_DL")
				createTransitGatewayOptionsModel.Global = core.BoolPtr(true)
				createTransitGatewayOptionsModel.ResourceGroup = resourceGroupIdentityModel
				createTransitGatewayOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = transitGatewayApisService.CreateTransitGateway(createTransitGatewayOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke CreateTransitGateway with error: Operation validation and request error`, func() {
				transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(transitGatewayApisService).ToNot(BeNil())

				// Construct an instance of the ResourceGroupIdentity model
				resourceGroupIdentityModel := new(transitgatewayapisv1.ResourceGroupIdentity)
				resourceGroupIdentityModel.ID = core.StringPtr("56969d6043e9465c883cb9f7363e78e8")

				// Construct an instance of the CreateTransitGatewayOptions model
				createTransitGatewayOptionsModel := new(transitgatewayapisv1.CreateTransitGatewayOptions)
				createTransitGatewayOptionsModel.Location = core.StringPtr("us-south")
				createTransitGatewayOptionsModel.Name = core.StringPtr("Transit_Service_BWTN_SJ_DL")
				createTransitGatewayOptionsModel.Global = core.BoolPtr(true)
				createTransitGatewayOptionsModel.ResourceGroup = resourceGroupIdentityModel
				createTransitGatewayOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := transitGatewayApisService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := transitGatewayApisService.CreateTransitGateway(createTransitGatewayOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateTransitGatewayOptions model with no property values
				createTransitGatewayOptionsModelNew := new(transitgatewayapisv1.CreateTransitGatewayOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = transitGatewayApisService.CreateTransitGateway(createTransitGatewayOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`DeleteTransitGateway(deleteTransitGatewayOptions *DeleteTransitGatewayOptions)`, func() {
		version := "testString"
		deleteTransitGatewayPath := "/transit_gateways/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(deleteTransitGatewayPath))
					Expect(req.Method).To(Equal("DELETE"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteTransitGateway successfully`, func() {
				transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(transitGatewayApisService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := transitGatewayApisService.DeleteTransitGateway(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteTransitGatewayOptions model
				deleteTransitGatewayOptionsModel := new(transitgatewayapisv1.DeleteTransitGatewayOptions)
				deleteTransitGatewayOptionsModel.ID = core.StringPtr("testString")
				deleteTransitGatewayOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = transitGatewayApisService.DeleteTransitGateway(deleteTransitGatewayOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteTransitGateway with error: Operation validation and request error`, func() {
				transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(transitGatewayApisService).ToNot(BeNil())

				// Construct an instance of the DeleteTransitGatewayOptions model
				deleteTransitGatewayOptionsModel := new(transitgatewayapisv1.DeleteTransitGatewayOptions)
				deleteTransitGatewayOptionsModel.ID = core.StringPtr("testString")
				deleteTransitGatewayOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := transitGatewayApisService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := transitGatewayApisService.DeleteTransitGateway(deleteTransitGatewayOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteTransitGatewayOptions model with no property values
				deleteTransitGatewayOptionsModelNew := new(transitgatewayapisv1.DeleteTransitGatewayOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = transitGatewayApisService.DeleteTransitGateway(deleteTransitGatewayOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetTransitGateway(getTransitGatewayOptions *GetTransitGatewayOptions) - Operation response error`, func() {
		version := "testString"
		getTransitGatewayPath := "/transit_gateways/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getTransitGatewayPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetTransitGateway with error: Operation response processing error`, func() {
				transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(transitGatewayApisService).ToNot(BeNil())

				// Construct an instance of the GetTransitGatewayOptions model
				getTransitGatewayOptionsModel := new(transitgatewayapisv1.GetTransitGatewayOptions)
				getTransitGatewayOptionsModel.ID = core.StringPtr("testString")
				getTransitGatewayOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := transitGatewayApisService.GetTransitGateway(getTransitGatewayOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetTransitGateway(getTransitGatewayOptions *GetTransitGatewayOptions)`, func() {
		version := "testString"
		getTransitGatewayPath := "/transit_gateways/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getTransitGatewayPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"id": "ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4", "crn": "crn:v1:bluemix:public:transit:dal03:a/57a7d05f36894e3cb9b46a43556d903e::gateway:ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4", "name": "my-transit-gateway-in-TransitGateway", "location": "us-south", "created_at": "2019-01-01T12:00:00", "global": true, "resource_group": {"id": "56969d6043e9465c883cb9f7363e78e8", "href": "https://resource-manager.bluemix.net/v1/resource_groups/56969d6043e9465c883cb9f7363e78e8"}, "status": "available", "updated_at": "2019-01-01T12:00:00"}`)
				}))
			})
			It(`Invoke GetTransitGateway successfully`, func() {
				transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(transitGatewayApisService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := transitGatewayApisService.GetTransitGateway(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetTransitGatewayOptions model
				getTransitGatewayOptionsModel := new(transitgatewayapisv1.GetTransitGatewayOptions)
				getTransitGatewayOptionsModel.ID = core.StringPtr("testString")
				getTransitGatewayOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = transitGatewayApisService.GetTransitGateway(getTransitGatewayOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke GetTransitGateway with error: Operation validation and request error`, func() {
				transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(transitGatewayApisService).ToNot(BeNil())

				// Construct an instance of the GetTransitGatewayOptions model
				getTransitGatewayOptionsModel := new(transitgatewayapisv1.GetTransitGatewayOptions)
				getTransitGatewayOptionsModel.ID = core.StringPtr("testString")
				getTransitGatewayOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := transitGatewayApisService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := transitGatewayApisService.GetTransitGateway(getTransitGatewayOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetTransitGatewayOptions model with no property values
				getTransitGatewayOptionsModelNew := new(transitgatewayapisv1.GetTransitGatewayOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = transitGatewayApisService.GetTransitGateway(getTransitGatewayOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateTransitGateway(updateTransitGatewayOptions *UpdateTransitGatewayOptions) - Operation response error`, func() {
		version := "testString"
		updateTransitGatewayPath := "/transit_gateways/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(updateTransitGatewayPath))
					Expect(req.Method).To(Equal("PATCH"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateTransitGateway with error: Operation response processing error`, func() {
				transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(transitGatewayApisService).ToNot(BeNil())

				// Construct an instance of the UpdateTransitGatewayOptions model
				updateTransitGatewayOptionsModel := new(transitgatewayapisv1.UpdateTransitGatewayOptions)
				updateTransitGatewayOptionsModel.ID = core.StringPtr("testString")
				updateTransitGatewayOptionsModel.Global = core.BoolPtr(true)
				updateTransitGatewayOptionsModel.Name = core.StringPtr("my-transit-gateway")
				updateTransitGatewayOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := transitGatewayApisService.UpdateTransitGateway(updateTransitGatewayOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UpdateTransitGateway(updateTransitGatewayOptions *UpdateTransitGatewayOptions)`, func() {
		version := "testString"
		updateTransitGatewayPath := "/transit_gateways/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(updateTransitGatewayPath))
					Expect(req.Method).To(Equal("PATCH"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"id": "ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4", "crn": "crn:v1:bluemix:public:transit:dal03:a/57a7d05f36894e3cb9b46a43556d903e::gateway:ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4", "name": "my-transit-gateway-in-TransitGateway", "location": "us-south", "created_at": "2019-01-01T12:00:00", "global": true, "resource_group": {"id": "56969d6043e9465c883cb9f7363e78e8", "href": "https://resource-manager.bluemix.net/v1/resource_groups/56969d6043e9465c883cb9f7363e78e8"}, "status": "available", "updated_at": "2019-01-01T12:00:00"}`)
				}))
			})
			It(`Invoke UpdateTransitGateway successfully`, func() {
				transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(transitGatewayApisService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := transitGatewayApisService.UpdateTransitGateway(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateTransitGatewayOptions model
				updateTransitGatewayOptionsModel := new(transitgatewayapisv1.UpdateTransitGatewayOptions)
				updateTransitGatewayOptionsModel.ID = core.StringPtr("testString")
				updateTransitGatewayOptionsModel.Global = core.BoolPtr(true)
				updateTransitGatewayOptionsModel.Name = core.StringPtr("my-transit-gateway")
				updateTransitGatewayOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = transitGatewayApisService.UpdateTransitGateway(updateTransitGatewayOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke UpdateTransitGateway with error: Operation validation and request error`, func() {
				transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(transitGatewayApisService).ToNot(BeNil())

				// Construct an instance of the UpdateTransitGatewayOptions model
				updateTransitGatewayOptionsModel := new(transitgatewayapisv1.UpdateTransitGatewayOptions)
				updateTransitGatewayOptionsModel.ID = core.StringPtr("testString")
				updateTransitGatewayOptionsModel.Global = core.BoolPtr(true)
				updateTransitGatewayOptionsModel.Name = core.StringPtr("my-transit-gateway")
				updateTransitGatewayOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := transitGatewayApisService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := transitGatewayApisService.UpdateTransitGateway(updateTransitGatewayOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateTransitGatewayOptions model with no property values
				updateTransitGatewayOptionsModelNew := new(transitgatewayapisv1.UpdateTransitGatewayOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = transitGatewayApisService.UpdateTransitGateway(updateTransitGatewayOptionsModelNew)
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
			transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
				Version: core.StringPtr(version),
			})
			Expect(transitGatewayApisService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
				URL: "{BAD_URL_STRING",
				Version: core.StringPtr(version),
			})
			Expect(transitGatewayApisService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
				URL: "https://transitgatewayapisv1/api",
				Version: core.StringPtr(version),
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(transitGatewayApisService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Validation Error`, func() {
			transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{})
			Expect(transitGatewayApisService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		version := "testString"
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"TRANSIT_GATEWAY_APIS_URL": "https://transitgatewayapisv1/api",
				"TRANSIT_GATEWAY_APIS_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1UsingExternalConfig(&transitgatewayapisv1.TransitGatewayApisV1Options{
					Version: core.StringPtr(version),
				})
				Expect(transitGatewayApisService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1UsingExternalConfig(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL: "https://testService/api",
					Version: core.StringPtr(version),
				})
				Expect(transitGatewayApisService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(transitGatewayApisService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1UsingExternalConfig(&transitgatewayapisv1.TransitGatewayApisV1Options{
					Version: core.StringPtr(version),
				})
				err := transitGatewayApisService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(transitGatewayApisService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(transitGatewayApisService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"TRANSIT_GATEWAY_APIS_URL": "https://transitgatewayapisv1/api",
				"TRANSIT_GATEWAY_APIS_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1UsingExternalConfig(&transitgatewayapisv1.TransitGatewayApisV1Options{
				Version: core.StringPtr(version),
			})

			It(`Instantiate service client with error`, func() {
				Expect(transitGatewayApisService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"TRANSIT_GATEWAY_APIS_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1UsingExternalConfig(&transitgatewayapisv1.TransitGatewayApisV1Options{
				URL: "{BAD_URL_STRING",
				Version: core.StringPtr(version),
			})

			It(`Instantiate service client with error`, func() {
				Expect(transitGatewayApisService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`ListTransitGatewayConnections(listTransitGatewayConnectionsOptions *ListTransitGatewayConnectionsOptions) - Operation response error`, func() {
		version := "testString"
		listTransitGatewayConnectionsPath := "/transit_gateways/testString/connections"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(listTransitGatewayConnectionsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListTransitGatewayConnections with error: Operation response processing error`, func() {
				transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(transitGatewayApisService).ToNot(BeNil())

				// Construct an instance of the ListTransitGatewayConnectionsOptions model
				listTransitGatewayConnectionsOptionsModel := new(transitgatewayapisv1.ListTransitGatewayConnectionsOptions)
				listTransitGatewayConnectionsOptionsModel.TransitGatewayID = core.StringPtr("testString")
				listTransitGatewayConnectionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := transitGatewayApisService.ListTransitGatewayConnections(listTransitGatewayConnectionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ListTransitGatewayConnections(listTransitGatewayConnectionsOptions *ListTransitGatewayConnectionsOptions)`, func() {
		version := "testString"
		listTransitGatewayConnectionsPath := "/transit_gateways/testString/connections"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(listTransitGatewayConnectionsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"connections": [{"name": "Transit_Service_BWTN_SJ_DL", "network_id": "crn:v1:bluemix:public:is:us-south:a/123456::vpc:4727d842-f94f-4a2d-824a-9bc9b02c523b", "network_type": "vpc", "id": "1a15dca5-7e33-45e1-b7c5-bc690e569531", "base_connection_id": "975f58c1-afe7-469a-9727-7f3d720f2d32", "created_at": "2019-01-01T12:00:00", "local_bgp_asn": 64490, "local_gateway_ip": "192.168.100.1", "local_tunnel_ip": "192.168.129.2", "mtu": 9000, "network_account_id": "28e4d90ac7504be694471ee66e70d0d5", "remote_bgp_asn": 65010, "remote_gateway_ip": "10.242.63.12", "remote_tunnel_ip": "192.168.129.1", "request_status": "pending", "status": "attached", "updated_at": "2019-01-01T12:00:00", "zone": {"name": "us-south-1"}}]}`)
				}))
			})
			It(`Invoke ListTransitGatewayConnections successfully`, func() {
				transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(transitGatewayApisService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := transitGatewayApisService.ListTransitGatewayConnections(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListTransitGatewayConnectionsOptions model
				listTransitGatewayConnectionsOptionsModel := new(transitgatewayapisv1.ListTransitGatewayConnectionsOptions)
				listTransitGatewayConnectionsOptionsModel.TransitGatewayID = core.StringPtr("testString")
				listTransitGatewayConnectionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = transitGatewayApisService.ListTransitGatewayConnections(listTransitGatewayConnectionsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke ListTransitGatewayConnections with error: Operation validation and request error`, func() {
				transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(transitGatewayApisService).ToNot(BeNil())

				// Construct an instance of the ListTransitGatewayConnectionsOptions model
				listTransitGatewayConnectionsOptionsModel := new(transitgatewayapisv1.ListTransitGatewayConnectionsOptions)
				listTransitGatewayConnectionsOptionsModel.TransitGatewayID = core.StringPtr("testString")
				listTransitGatewayConnectionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := transitGatewayApisService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := transitGatewayApisService.ListTransitGatewayConnections(listTransitGatewayConnectionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListTransitGatewayConnectionsOptions model with no property values
				listTransitGatewayConnectionsOptionsModelNew := new(transitgatewayapisv1.ListTransitGatewayConnectionsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = transitGatewayApisService.ListTransitGatewayConnections(listTransitGatewayConnectionsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateTransitGatewayConnection(createTransitGatewayConnectionOptions *CreateTransitGatewayConnectionOptions) - Operation response error`, func() {
		version := "testString"
		createTransitGatewayConnectionPath := "/transit_gateways/testString/connections"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(createTransitGatewayConnectionPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateTransitGatewayConnection with error: Operation response processing error`, func() {
				transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(transitGatewayApisService).ToNot(BeNil())

				// Construct an instance of the ZoneIdentityByName model
				zoneIdentityModel := new(transitgatewayapisv1.ZoneIdentityByName)
				zoneIdentityModel.Name = core.StringPtr("us-south-1")

				// Construct an instance of the CreateTransitGatewayConnectionOptions model
				createTransitGatewayConnectionOptionsModel := new(transitgatewayapisv1.CreateTransitGatewayConnectionOptions)
				createTransitGatewayConnectionOptionsModel.TransitGatewayID = core.StringPtr("testString")
				createTransitGatewayConnectionOptionsModel.NetworkType = core.StringPtr("vpc")
				createTransitGatewayConnectionOptionsModel.BaseConnectionID = core.StringPtr("975f58c1-afe7-469a-9727-7f3d720f2d32")
				createTransitGatewayConnectionOptionsModel.LocalGatewayIp = core.StringPtr("192.168.100.1")
				createTransitGatewayConnectionOptionsModel.LocalTunnelIp = core.StringPtr("192.168.129.2")
				createTransitGatewayConnectionOptionsModel.Name = core.StringPtr("Transit_Service_BWTN_SJ_DL")
				createTransitGatewayConnectionOptionsModel.NetworkAccountID = core.StringPtr("28e4d90ac7504be694471ee66e70d0d5")
				createTransitGatewayConnectionOptionsModel.NetworkID = core.StringPtr("crn:v1:bluemix:public:is:us-south:a/123456::vpc:4727d842-f94f-4a2d-824a-9bc9b02c523b")
				createTransitGatewayConnectionOptionsModel.RemoteBgpAsn = core.StringPtr("65010")
				createTransitGatewayConnectionOptionsModel.RemoteGatewayIp = core.StringPtr("10.242.63.12")
				createTransitGatewayConnectionOptionsModel.RemoteTunnelIp = core.StringPtr("192.168.129.1")
				createTransitGatewayConnectionOptionsModel.Zone = zoneIdentityModel
				createTransitGatewayConnectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := transitGatewayApisService.CreateTransitGatewayConnection(createTransitGatewayConnectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`CreateTransitGatewayConnection(createTransitGatewayConnectionOptions *CreateTransitGatewayConnectionOptions)`, func() {
		version := "testString"
		createTransitGatewayConnectionPath := "/transit_gateways/testString/connections"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(createTransitGatewayConnectionPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, `{"name": "Transit_Service_BWTN_SJ_DL", "network_id": "crn:v1:bluemix:public:is:us-south:a/123456::vpc:4727d842-f94f-4a2d-824a-9bc9b02c523b", "network_type": "vpc", "id": "1a15dca5-7e33-45e1-b7c5-bc690e569531", "base_connection_id": "975f58c1-afe7-469a-9727-7f3d720f2d32", "created_at": "2019-01-01T12:00:00", "local_bgp_asn": 64490, "local_gateway_ip": "192.168.100.1", "local_tunnel_ip": "192.168.129.2", "mtu": 9000, "network_account_id": "28e4d90ac7504be694471ee66e70d0d5", "remote_bgp_asn": 65010, "remote_gateway_ip": "10.242.63.12", "remote_tunnel_ip": "192.168.129.1", "request_status": "pending", "status": "attached", "updated_at": "2019-01-01T12:00:00", "zone": {"name": "us-south-1"}}`)
				}))
			})
			It(`Invoke CreateTransitGatewayConnection successfully`, func() {
				transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(transitGatewayApisService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := transitGatewayApisService.CreateTransitGatewayConnection(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ZoneIdentityByName model
				zoneIdentityModel := new(transitgatewayapisv1.ZoneIdentityByName)
				zoneIdentityModel.Name = core.StringPtr("us-south-1")

				// Construct an instance of the CreateTransitGatewayConnectionOptions model
				createTransitGatewayConnectionOptionsModel := new(transitgatewayapisv1.CreateTransitGatewayConnectionOptions)
				createTransitGatewayConnectionOptionsModel.TransitGatewayID = core.StringPtr("testString")
				createTransitGatewayConnectionOptionsModel.NetworkType = core.StringPtr("vpc")
				createTransitGatewayConnectionOptionsModel.BaseConnectionID = core.StringPtr("975f58c1-afe7-469a-9727-7f3d720f2d32")
				createTransitGatewayConnectionOptionsModel.LocalGatewayIp = core.StringPtr("192.168.100.1")
				createTransitGatewayConnectionOptionsModel.LocalTunnelIp = core.StringPtr("192.168.129.2")
				createTransitGatewayConnectionOptionsModel.Name = core.StringPtr("Transit_Service_BWTN_SJ_DL")
				createTransitGatewayConnectionOptionsModel.NetworkAccountID = core.StringPtr("28e4d90ac7504be694471ee66e70d0d5")
				createTransitGatewayConnectionOptionsModel.NetworkID = core.StringPtr("crn:v1:bluemix:public:is:us-south:a/123456::vpc:4727d842-f94f-4a2d-824a-9bc9b02c523b")
				createTransitGatewayConnectionOptionsModel.RemoteBgpAsn = core.StringPtr("65010")
				createTransitGatewayConnectionOptionsModel.RemoteGatewayIp = core.StringPtr("10.242.63.12")
				createTransitGatewayConnectionOptionsModel.RemoteTunnelIp = core.StringPtr("192.168.129.1")
				createTransitGatewayConnectionOptionsModel.Zone = zoneIdentityModel
				createTransitGatewayConnectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = transitGatewayApisService.CreateTransitGatewayConnection(createTransitGatewayConnectionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke CreateTransitGatewayConnection with error: Operation validation and request error`, func() {
				transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(transitGatewayApisService).ToNot(BeNil())

				// Construct an instance of the ZoneIdentityByName model
				zoneIdentityModel := new(transitgatewayapisv1.ZoneIdentityByName)
				zoneIdentityModel.Name = core.StringPtr("us-south-1")

				// Construct an instance of the CreateTransitGatewayConnectionOptions model
				createTransitGatewayConnectionOptionsModel := new(transitgatewayapisv1.CreateTransitGatewayConnectionOptions)
				createTransitGatewayConnectionOptionsModel.TransitGatewayID = core.StringPtr("testString")
				createTransitGatewayConnectionOptionsModel.NetworkType = core.StringPtr("vpc")
				createTransitGatewayConnectionOptionsModel.BaseConnectionID = core.StringPtr("975f58c1-afe7-469a-9727-7f3d720f2d32")
				createTransitGatewayConnectionOptionsModel.LocalGatewayIp = core.StringPtr("192.168.100.1")
				createTransitGatewayConnectionOptionsModel.LocalTunnelIp = core.StringPtr("192.168.129.2")
				createTransitGatewayConnectionOptionsModel.Name = core.StringPtr("Transit_Service_BWTN_SJ_DL")
				createTransitGatewayConnectionOptionsModel.NetworkAccountID = core.StringPtr("28e4d90ac7504be694471ee66e70d0d5")
				createTransitGatewayConnectionOptionsModel.NetworkID = core.StringPtr("crn:v1:bluemix:public:is:us-south:a/123456::vpc:4727d842-f94f-4a2d-824a-9bc9b02c523b")
				createTransitGatewayConnectionOptionsModel.RemoteBgpAsn = core.StringPtr("65010")
				createTransitGatewayConnectionOptionsModel.RemoteGatewayIp = core.StringPtr("10.242.63.12")
				createTransitGatewayConnectionOptionsModel.RemoteTunnelIp = core.StringPtr("192.168.129.1")
				createTransitGatewayConnectionOptionsModel.Zone = zoneIdentityModel
				createTransitGatewayConnectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := transitGatewayApisService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := transitGatewayApisService.CreateTransitGatewayConnection(createTransitGatewayConnectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateTransitGatewayConnectionOptions model with no property values
				createTransitGatewayConnectionOptionsModelNew := new(transitgatewayapisv1.CreateTransitGatewayConnectionOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = transitGatewayApisService.CreateTransitGatewayConnection(createTransitGatewayConnectionOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`DeleteTransitGatewayConnection(deleteTransitGatewayConnectionOptions *DeleteTransitGatewayConnectionOptions)`, func() {
		version := "testString"
		deleteTransitGatewayConnectionPath := "/transit_gateways/testString/connections/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(deleteTransitGatewayConnectionPath))
					Expect(req.Method).To(Equal("DELETE"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteTransitGatewayConnection successfully`, func() {
				transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(transitGatewayApisService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := transitGatewayApisService.DeleteTransitGatewayConnection(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteTransitGatewayConnectionOptions model
				deleteTransitGatewayConnectionOptionsModel := new(transitgatewayapisv1.DeleteTransitGatewayConnectionOptions)
				deleteTransitGatewayConnectionOptionsModel.TransitGatewayID = core.StringPtr("testString")
				deleteTransitGatewayConnectionOptionsModel.ID = core.StringPtr("testString")
				deleteTransitGatewayConnectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = transitGatewayApisService.DeleteTransitGatewayConnection(deleteTransitGatewayConnectionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteTransitGatewayConnection with error: Operation validation and request error`, func() {
				transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(transitGatewayApisService).ToNot(BeNil())

				// Construct an instance of the DeleteTransitGatewayConnectionOptions model
				deleteTransitGatewayConnectionOptionsModel := new(transitgatewayapisv1.DeleteTransitGatewayConnectionOptions)
				deleteTransitGatewayConnectionOptionsModel.TransitGatewayID = core.StringPtr("testString")
				deleteTransitGatewayConnectionOptionsModel.ID = core.StringPtr("testString")
				deleteTransitGatewayConnectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := transitGatewayApisService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := transitGatewayApisService.DeleteTransitGatewayConnection(deleteTransitGatewayConnectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteTransitGatewayConnectionOptions model with no property values
				deleteTransitGatewayConnectionOptionsModelNew := new(transitgatewayapisv1.DeleteTransitGatewayConnectionOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = transitGatewayApisService.DeleteTransitGatewayConnection(deleteTransitGatewayConnectionOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetTransitGatewayConnection(getTransitGatewayConnectionOptions *GetTransitGatewayConnectionOptions) - Operation response error`, func() {
		version := "testString"
		getTransitGatewayConnectionPath := "/transit_gateways/testString/connections/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getTransitGatewayConnectionPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetTransitGatewayConnection with error: Operation response processing error`, func() {
				transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(transitGatewayApisService).ToNot(BeNil())

				// Construct an instance of the GetTransitGatewayConnectionOptions model
				getTransitGatewayConnectionOptionsModel := new(transitgatewayapisv1.GetTransitGatewayConnectionOptions)
				getTransitGatewayConnectionOptionsModel.TransitGatewayID = core.StringPtr("testString")
				getTransitGatewayConnectionOptionsModel.ID = core.StringPtr("testString")
				getTransitGatewayConnectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := transitGatewayApisService.GetTransitGatewayConnection(getTransitGatewayConnectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetTransitGatewayConnection(getTransitGatewayConnectionOptions *GetTransitGatewayConnectionOptions)`, func() {
		version := "testString"
		getTransitGatewayConnectionPath := "/transit_gateways/testString/connections/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getTransitGatewayConnectionPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"name": "Transit_Service_BWTN_SJ_DL", "network_id": "crn:v1:bluemix:public:is:us-south:a/123456::vpc:4727d842-f94f-4a2d-824a-9bc9b02c523b", "network_type": "vpc", "id": "1a15dca5-7e33-45e1-b7c5-bc690e569531", "base_connection_id": "975f58c1-afe7-469a-9727-7f3d720f2d32", "created_at": "2019-01-01T12:00:00", "local_bgp_asn": 64490, "local_gateway_ip": "192.168.100.1", "local_tunnel_ip": "192.168.129.2", "mtu": 9000, "network_account_id": "28e4d90ac7504be694471ee66e70d0d5", "remote_bgp_asn": 65010, "remote_gateway_ip": "10.242.63.12", "remote_tunnel_ip": "192.168.129.1", "request_status": "pending", "status": "attached", "updated_at": "2019-01-01T12:00:00", "zone": {"name": "us-south-1"}}`)
				}))
			})
			It(`Invoke GetTransitGatewayConnection successfully`, func() {
				transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(transitGatewayApisService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := transitGatewayApisService.GetTransitGatewayConnection(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetTransitGatewayConnectionOptions model
				getTransitGatewayConnectionOptionsModel := new(transitgatewayapisv1.GetTransitGatewayConnectionOptions)
				getTransitGatewayConnectionOptionsModel.TransitGatewayID = core.StringPtr("testString")
				getTransitGatewayConnectionOptionsModel.ID = core.StringPtr("testString")
				getTransitGatewayConnectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = transitGatewayApisService.GetTransitGatewayConnection(getTransitGatewayConnectionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke GetTransitGatewayConnection with error: Operation validation and request error`, func() {
				transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(transitGatewayApisService).ToNot(BeNil())

				// Construct an instance of the GetTransitGatewayConnectionOptions model
				getTransitGatewayConnectionOptionsModel := new(transitgatewayapisv1.GetTransitGatewayConnectionOptions)
				getTransitGatewayConnectionOptionsModel.TransitGatewayID = core.StringPtr("testString")
				getTransitGatewayConnectionOptionsModel.ID = core.StringPtr("testString")
				getTransitGatewayConnectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := transitGatewayApisService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := transitGatewayApisService.GetTransitGatewayConnection(getTransitGatewayConnectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetTransitGatewayConnectionOptions model with no property values
				getTransitGatewayConnectionOptionsModelNew := new(transitgatewayapisv1.GetTransitGatewayConnectionOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = transitGatewayApisService.GetTransitGatewayConnection(getTransitGatewayConnectionOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateTransitGatewayConnection(updateTransitGatewayConnectionOptions *UpdateTransitGatewayConnectionOptions) - Operation response error`, func() {
		version := "testString"
		updateTransitGatewayConnectionPath := "/transit_gateways/testString/connections/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(updateTransitGatewayConnectionPath))
					Expect(req.Method).To(Equal("PATCH"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateTransitGatewayConnection with error: Operation response processing error`, func() {
				transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(transitGatewayApisService).ToNot(BeNil())

				// Construct an instance of the UpdateTransitGatewayConnectionOptions model
				updateTransitGatewayConnectionOptionsModel := new(transitgatewayapisv1.UpdateTransitGatewayConnectionOptions)
				updateTransitGatewayConnectionOptionsModel.TransitGatewayID = core.StringPtr("testString")
				updateTransitGatewayConnectionOptionsModel.ID = core.StringPtr("testString")
				updateTransitGatewayConnectionOptionsModel.Name = core.StringPtr("Transit_Service_BWTN_SJ_DL")
				updateTransitGatewayConnectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := transitGatewayApisService.UpdateTransitGatewayConnection(updateTransitGatewayConnectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UpdateTransitGatewayConnection(updateTransitGatewayConnectionOptions *UpdateTransitGatewayConnectionOptions)`, func() {
		version := "testString"
		updateTransitGatewayConnectionPath := "/transit_gateways/testString/connections/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(updateTransitGatewayConnectionPath))
					Expect(req.Method).To(Equal("PATCH"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"name": "Transit_Service_BWTN_SJ_DL", "network_id": "crn:v1:bluemix:public:is:us-south:a/123456::vpc:4727d842-f94f-4a2d-824a-9bc9b02c523b", "network_type": "vpc", "id": "1a15dca5-7e33-45e1-b7c5-bc690e569531", "base_connection_id": "975f58c1-afe7-469a-9727-7f3d720f2d32", "created_at": "2019-01-01T12:00:00", "local_bgp_asn": 64490, "local_gateway_ip": "192.168.100.1", "local_tunnel_ip": "192.168.129.2", "mtu": 9000, "network_account_id": "28e4d90ac7504be694471ee66e70d0d5", "remote_bgp_asn": 65010, "remote_gateway_ip": "10.242.63.12", "remote_tunnel_ip": "192.168.129.1", "request_status": "pending", "status": "attached", "updated_at": "2019-01-01T12:00:00", "zone": {"name": "us-south-1"}}`)
				}))
			})
			It(`Invoke UpdateTransitGatewayConnection successfully`, func() {
				transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(transitGatewayApisService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := transitGatewayApisService.UpdateTransitGatewayConnection(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateTransitGatewayConnectionOptions model
				updateTransitGatewayConnectionOptionsModel := new(transitgatewayapisv1.UpdateTransitGatewayConnectionOptions)
				updateTransitGatewayConnectionOptionsModel.TransitGatewayID = core.StringPtr("testString")
				updateTransitGatewayConnectionOptionsModel.ID = core.StringPtr("testString")
				updateTransitGatewayConnectionOptionsModel.Name = core.StringPtr("Transit_Service_BWTN_SJ_DL")
				updateTransitGatewayConnectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = transitGatewayApisService.UpdateTransitGatewayConnection(updateTransitGatewayConnectionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke UpdateTransitGatewayConnection with error: Operation validation and request error`, func() {
				transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(transitGatewayApisService).ToNot(BeNil())

				// Construct an instance of the UpdateTransitGatewayConnectionOptions model
				updateTransitGatewayConnectionOptionsModel := new(transitgatewayapisv1.UpdateTransitGatewayConnectionOptions)
				updateTransitGatewayConnectionOptionsModel.TransitGatewayID = core.StringPtr("testString")
				updateTransitGatewayConnectionOptionsModel.ID = core.StringPtr("testString")
				updateTransitGatewayConnectionOptionsModel.Name = core.StringPtr("Transit_Service_BWTN_SJ_DL")
				updateTransitGatewayConnectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := transitGatewayApisService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := transitGatewayApisService.UpdateTransitGatewayConnection(updateTransitGatewayConnectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateTransitGatewayConnectionOptions model with no property values
				updateTransitGatewayConnectionOptionsModelNew := new(transitgatewayapisv1.UpdateTransitGatewayConnectionOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = transitGatewayApisService.UpdateTransitGatewayConnection(updateTransitGatewayConnectionOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`CreateTransitGatewayConnectionActions(createTransitGatewayConnectionActionsOptions *CreateTransitGatewayConnectionActionsOptions)`, func() {
		version := "testString"
		createTransitGatewayConnectionActionsPath := "/transit_gateways/testString/connections/testString/actions"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(createTransitGatewayConnectionActionsPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					res.WriteHeader(204)
				}))
			})
			It(`Invoke CreateTransitGatewayConnectionActions successfully`, func() {
				transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(transitGatewayApisService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := transitGatewayApisService.CreateTransitGatewayConnectionActions(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the CreateTransitGatewayConnectionActionsOptions model
				createTransitGatewayConnectionActionsOptionsModel := new(transitgatewayapisv1.CreateTransitGatewayConnectionActionsOptions)
				createTransitGatewayConnectionActionsOptionsModel.TransitGatewayID = core.StringPtr("testString")
				createTransitGatewayConnectionActionsOptionsModel.ID = core.StringPtr("testString")
				createTransitGatewayConnectionActionsOptionsModel.Action = core.StringPtr("approve")
				createTransitGatewayConnectionActionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = transitGatewayApisService.CreateTransitGatewayConnectionActions(createTransitGatewayConnectionActionsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke CreateTransitGatewayConnectionActions with error: Operation validation and request error`, func() {
				transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(transitGatewayApisService).ToNot(BeNil())

				// Construct an instance of the CreateTransitGatewayConnectionActionsOptions model
				createTransitGatewayConnectionActionsOptionsModel := new(transitgatewayapisv1.CreateTransitGatewayConnectionActionsOptions)
				createTransitGatewayConnectionActionsOptionsModel.TransitGatewayID = core.StringPtr("testString")
				createTransitGatewayConnectionActionsOptionsModel.ID = core.StringPtr("testString")
				createTransitGatewayConnectionActionsOptionsModel.Action = core.StringPtr("approve")
				createTransitGatewayConnectionActionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := transitGatewayApisService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := transitGatewayApisService.CreateTransitGatewayConnectionActions(createTransitGatewayConnectionActionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the CreateTransitGatewayConnectionActionsOptions model with no property values
				createTransitGatewayConnectionActionsOptionsModelNew := new(transitgatewayapisv1.CreateTransitGatewayConnectionActionsOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = transitGatewayApisService.CreateTransitGatewayConnectionActions(createTransitGatewayConnectionActionsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Service constructor tests`, func() {
		version := "testString"
		It(`Instantiate service client`, func() {
			transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
				Version: core.StringPtr(version),
			})
			Expect(transitGatewayApisService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
				URL: "{BAD_URL_STRING",
				Version: core.StringPtr(version),
			})
			Expect(transitGatewayApisService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
				URL: "https://transitgatewayapisv1/api",
				Version: core.StringPtr(version),
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(transitGatewayApisService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Validation Error`, func() {
			transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{})
			Expect(transitGatewayApisService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		version := "testString"
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"TRANSIT_GATEWAY_APIS_URL": "https://transitgatewayapisv1/api",
				"TRANSIT_GATEWAY_APIS_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1UsingExternalConfig(&transitgatewayapisv1.TransitGatewayApisV1Options{
					Version: core.StringPtr(version),
				})
				Expect(transitGatewayApisService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1UsingExternalConfig(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL: "https://testService/api",
					Version: core.StringPtr(version),
				})
				Expect(transitGatewayApisService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(transitGatewayApisService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1UsingExternalConfig(&transitgatewayapisv1.TransitGatewayApisV1Options{
					Version: core.StringPtr(version),
				})
				err := transitGatewayApisService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(transitGatewayApisService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(transitGatewayApisService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"TRANSIT_GATEWAY_APIS_URL": "https://transitgatewayapisv1/api",
				"TRANSIT_GATEWAY_APIS_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1UsingExternalConfig(&transitgatewayapisv1.TransitGatewayApisV1Options{
				Version: core.StringPtr(version),
			})

			It(`Instantiate service client with error`, func() {
				Expect(transitGatewayApisService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"TRANSIT_GATEWAY_APIS_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1UsingExternalConfig(&transitgatewayapisv1.TransitGatewayApisV1Options{
				URL: "{BAD_URL_STRING",
				Version: core.StringPtr(version),
			})

			It(`Instantiate service client with error`, func() {
				Expect(transitGatewayApisService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`ListGatewayLocations(listGatewayLocationsOptions *ListGatewayLocationsOptions) - Operation response error`, func() {
		version := "testString"
		listGatewayLocationsPath := "/locations"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(listGatewayLocationsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListGatewayLocations with error: Operation response processing error`, func() {
				transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(transitGatewayApisService).ToNot(BeNil())

				// Construct an instance of the ListGatewayLocationsOptions model
				listGatewayLocationsOptionsModel := new(transitgatewayapisv1.ListGatewayLocationsOptions)
				listGatewayLocationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := transitGatewayApisService.ListGatewayLocations(listGatewayLocationsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ListGatewayLocations(listGatewayLocationsOptions *ListGatewayLocationsOptions)`, func() {
		version := "testString"
		listGatewayLocationsPath := "/locations"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(listGatewayLocationsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"locations": [{"billing_location": "us", "name": "us-south", "type": "region"}]}`)
				}))
			})
			It(`Invoke ListGatewayLocations successfully`, func() {
				transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(transitGatewayApisService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := transitGatewayApisService.ListGatewayLocations(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListGatewayLocationsOptions model
				listGatewayLocationsOptionsModel := new(transitgatewayapisv1.ListGatewayLocationsOptions)
				listGatewayLocationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = transitGatewayApisService.ListGatewayLocations(listGatewayLocationsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke ListGatewayLocations with error: Operation request error`, func() {
				transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(transitGatewayApisService).ToNot(BeNil())

				// Construct an instance of the ListGatewayLocationsOptions model
				listGatewayLocationsOptionsModel := new(transitgatewayapisv1.ListGatewayLocationsOptions)
				listGatewayLocationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := transitGatewayApisService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := transitGatewayApisService.ListGatewayLocations(listGatewayLocationsOptionsModel)
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
	Describe(`GetGatewayLocation(getGatewayLocationOptions *GetGatewayLocationOptions) - Operation response error`, func() {
		version := "testString"
		getGatewayLocationPath := "/locations/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getGatewayLocationPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetGatewayLocation with error: Operation response processing error`, func() {
				transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(transitGatewayApisService).ToNot(BeNil())

				// Construct an instance of the GetGatewayLocationOptions model
				getGatewayLocationOptionsModel := new(transitgatewayapisv1.GetGatewayLocationOptions)
				getGatewayLocationOptionsModel.Name = core.StringPtr("testString")
				getGatewayLocationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := transitGatewayApisService.GetGatewayLocation(getGatewayLocationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetGatewayLocation(getGatewayLocationOptions *GetGatewayLocationOptions)`, func() {
		version := "testString"
		getGatewayLocationPath := "/locations/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getGatewayLocationPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"billing_location": "us", "name": "us-south", "type": "region", "local_connection_locations": [{"display_name": "Dallas", "name": "us-south", "type": "region"}]}`)
				}))
			})
			It(`Invoke GetGatewayLocation successfully`, func() {
				transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(transitGatewayApisService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := transitGatewayApisService.GetGatewayLocation(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetGatewayLocationOptions model
				getGatewayLocationOptionsModel := new(transitgatewayapisv1.GetGatewayLocationOptions)
				getGatewayLocationOptionsModel.Name = core.StringPtr("testString")
				getGatewayLocationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = transitGatewayApisService.GetGatewayLocation(getGatewayLocationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke GetGatewayLocation with error: Operation validation and request error`, func() {
				transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(transitGatewayApisService).ToNot(BeNil())

				// Construct an instance of the GetGatewayLocationOptions model
				getGatewayLocationOptionsModel := new(transitgatewayapisv1.GetGatewayLocationOptions)
				getGatewayLocationOptionsModel.Name = core.StringPtr("testString")
				getGatewayLocationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := transitGatewayApisService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := transitGatewayApisService.GetGatewayLocation(getGatewayLocationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetGatewayLocationOptions model with no property values
				getGatewayLocationOptionsModelNew := new(transitgatewayapisv1.GetGatewayLocationOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = transitGatewayApisService.GetGatewayLocation(getGatewayLocationOptionsModelNew)
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
			transitGatewayApisService, _ := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
				URL:           "http://transitgatewayapisv1modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
				Version: core.StringPtr(version),
			})
			It(`Invoke NewCreateTransitGatewayConnectionActionsOptions successfully`, func() {
				// Construct an instance of the CreateTransitGatewayConnectionActionsOptions model
				transitGatewayID := "testString"
				id := "testString"
				createTransitGatewayConnectionActionsOptionsAction := "approve"
				createTransitGatewayConnectionActionsOptionsModel := transitGatewayApisService.NewCreateTransitGatewayConnectionActionsOptions(transitGatewayID, id, createTransitGatewayConnectionActionsOptionsAction)
				createTransitGatewayConnectionActionsOptionsModel.SetTransitGatewayID("testString")
				createTransitGatewayConnectionActionsOptionsModel.SetID("testString")
				createTransitGatewayConnectionActionsOptionsModel.SetAction("approve")
				createTransitGatewayConnectionActionsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createTransitGatewayConnectionActionsOptionsModel).ToNot(BeNil())
				Expect(createTransitGatewayConnectionActionsOptionsModel.TransitGatewayID).To(Equal(core.StringPtr("testString")))
				Expect(createTransitGatewayConnectionActionsOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(createTransitGatewayConnectionActionsOptionsModel.Action).To(Equal(core.StringPtr("approve")))
				Expect(createTransitGatewayConnectionActionsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateTransitGatewayConnectionOptions successfully`, func() {
				// Construct an instance of the ZoneIdentityByName model
				zoneIdentityModel := new(transitgatewayapisv1.ZoneIdentityByName)
				Expect(zoneIdentityModel).ToNot(BeNil())
				zoneIdentityModel.Name = core.StringPtr("us-south-1")
				Expect(zoneIdentityModel.Name).To(Equal(core.StringPtr("us-south-1")))

				// Construct an instance of the CreateTransitGatewayConnectionOptions model
				transitGatewayID := "testString"
				createTransitGatewayConnectionOptionsNetworkType := "vpc"
				createTransitGatewayConnectionOptionsModel := transitGatewayApisService.NewCreateTransitGatewayConnectionOptions(transitGatewayID, createTransitGatewayConnectionOptionsNetworkType)
				createTransitGatewayConnectionOptionsModel.SetTransitGatewayID("testString")
				createTransitGatewayConnectionOptionsModel.SetNetworkType("vpc")
				createTransitGatewayConnectionOptionsModel.SetBaseConnectionID("975f58c1-afe7-469a-9727-7f3d720f2d32")
				createTransitGatewayConnectionOptionsModel.SetLocalGatewayIp("192.168.100.1")
				createTransitGatewayConnectionOptionsModel.SetLocalTunnelIp("192.168.129.2")
				createTransitGatewayConnectionOptionsModel.SetName("Transit_Service_BWTN_SJ_DL")
				createTransitGatewayConnectionOptionsModel.SetNetworkAccountID("28e4d90ac7504be694471ee66e70d0d5")
				createTransitGatewayConnectionOptionsModel.SetNetworkID("crn:v1:bluemix:public:is:us-south:a/123456::vpc:4727d842-f94f-4a2d-824a-9bc9b02c523b")
				createTransitGatewayConnectionOptionsModel.SetRemoteBgpAsn("65010")
				createTransitGatewayConnectionOptionsModel.SetRemoteGatewayIp("10.242.63.12")
				createTransitGatewayConnectionOptionsModel.SetRemoteTunnelIp("192.168.129.1")
				createTransitGatewayConnectionOptionsModel.SetZone(zoneIdentityModel)
				createTransitGatewayConnectionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createTransitGatewayConnectionOptionsModel).ToNot(BeNil())
				Expect(createTransitGatewayConnectionOptionsModel.TransitGatewayID).To(Equal(core.StringPtr("testString")))
				Expect(createTransitGatewayConnectionOptionsModel.NetworkType).To(Equal(core.StringPtr("vpc")))
				Expect(createTransitGatewayConnectionOptionsModel.BaseConnectionID).To(Equal(core.StringPtr("975f58c1-afe7-469a-9727-7f3d720f2d32")))
				Expect(createTransitGatewayConnectionOptionsModel.LocalGatewayIp).To(Equal(core.StringPtr("192.168.100.1")))
				Expect(createTransitGatewayConnectionOptionsModel.LocalTunnelIp).To(Equal(core.StringPtr("192.168.129.2")))
				Expect(createTransitGatewayConnectionOptionsModel.Name).To(Equal(core.StringPtr("Transit_Service_BWTN_SJ_DL")))
				Expect(createTransitGatewayConnectionOptionsModel.NetworkAccountID).To(Equal(core.StringPtr("28e4d90ac7504be694471ee66e70d0d5")))
				Expect(createTransitGatewayConnectionOptionsModel.NetworkID).To(Equal(core.StringPtr("crn:v1:bluemix:public:is:us-south:a/123456::vpc:4727d842-f94f-4a2d-824a-9bc9b02c523b")))
				Expect(createTransitGatewayConnectionOptionsModel.RemoteBgpAsn).To(Equal(core.StringPtr("65010")))
				Expect(createTransitGatewayConnectionOptionsModel.RemoteGatewayIp).To(Equal(core.StringPtr("10.242.63.12")))
				Expect(createTransitGatewayConnectionOptionsModel.RemoteTunnelIp).To(Equal(core.StringPtr("192.168.129.1")))
				Expect(createTransitGatewayConnectionOptionsModel.Zone).To(Equal(zoneIdentityModel))
				Expect(createTransitGatewayConnectionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateTransitGatewayOptions successfully`, func() {
				// Construct an instance of the ResourceGroupIdentity model
				resourceGroupIdentityModel := new(transitgatewayapisv1.ResourceGroupIdentity)
				Expect(resourceGroupIdentityModel).ToNot(BeNil())
				resourceGroupIdentityModel.ID = core.StringPtr("56969d6043e9465c883cb9f7363e78e8")
				Expect(resourceGroupIdentityModel.ID).To(Equal(core.StringPtr("56969d6043e9465c883cb9f7363e78e8")))

				// Construct an instance of the CreateTransitGatewayOptions model
				createTransitGatewayOptionsLocation := "us-south"
				createTransitGatewayOptionsName := "Transit_Service_BWTN_SJ_DL"
				createTransitGatewayOptionsModel := transitGatewayApisService.NewCreateTransitGatewayOptions(createTransitGatewayOptionsLocation, createTransitGatewayOptionsName)
				createTransitGatewayOptionsModel.SetLocation("us-south")
				createTransitGatewayOptionsModel.SetName("Transit_Service_BWTN_SJ_DL")
				createTransitGatewayOptionsModel.SetGlobal(true)
				createTransitGatewayOptionsModel.SetResourceGroup(resourceGroupIdentityModel)
				createTransitGatewayOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createTransitGatewayOptionsModel).ToNot(BeNil())
				Expect(createTransitGatewayOptionsModel.Location).To(Equal(core.StringPtr("us-south")))
				Expect(createTransitGatewayOptionsModel.Name).To(Equal(core.StringPtr("Transit_Service_BWTN_SJ_DL")))
				Expect(createTransitGatewayOptionsModel.Global).To(Equal(core.BoolPtr(true)))
				Expect(createTransitGatewayOptionsModel.ResourceGroup).To(Equal(resourceGroupIdentityModel))
				Expect(createTransitGatewayOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteTransitGatewayConnectionOptions successfully`, func() {
				// Construct an instance of the DeleteTransitGatewayConnectionOptions model
				transitGatewayID := "testString"
				id := "testString"
				deleteTransitGatewayConnectionOptionsModel := transitGatewayApisService.NewDeleteTransitGatewayConnectionOptions(transitGatewayID, id)
				deleteTransitGatewayConnectionOptionsModel.SetTransitGatewayID("testString")
				deleteTransitGatewayConnectionOptionsModel.SetID("testString")
				deleteTransitGatewayConnectionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteTransitGatewayConnectionOptionsModel).ToNot(BeNil())
				Expect(deleteTransitGatewayConnectionOptionsModel.TransitGatewayID).To(Equal(core.StringPtr("testString")))
				Expect(deleteTransitGatewayConnectionOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(deleteTransitGatewayConnectionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteTransitGatewayOptions successfully`, func() {
				// Construct an instance of the DeleteTransitGatewayOptions model
				id := "testString"
				deleteTransitGatewayOptionsModel := transitGatewayApisService.NewDeleteTransitGatewayOptions(id)
				deleteTransitGatewayOptionsModel.SetID("testString")
				deleteTransitGatewayOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteTransitGatewayOptionsModel).ToNot(BeNil())
				Expect(deleteTransitGatewayOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(deleteTransitGatewayOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetGatewayLocationOptions successfully`, func() {
				// Construct an instance of the GetGatewayLocationOptions model
				name := "testString"
				getGatewayLocationOptionsModel := transitGatewayApisService.NewGetGatewayLocationOptions(name)
				getGatewayLocationOptionsModel.SetName("testString")
				getGatewayLocationOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getGatewayLocationOptionsModel).ToNot(BeNil())
				Expect(getGatewayLocationOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(getGatewayLocationOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetTransitGatewayConnectionOptions successfully`, func() {
				// Construct an instance of the GetTransitGatewayConnectionOptions model
				transitGatewayID := "testString"
				id := "testString"
				getTransitGatewayConnectionOptionsModel := transitGatewayApisService.NewGetTransitGatewayConnectionOptions(transitGatewayID, id)
				getTransitGatewayConnectionOptionsModel.SetTransitGatewayID("testString")
				getTransitGatewayConnectionOptionsModel.SetID("testString")
				getTransitGatewayConnectionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getTransitGatewayConnectionOptionsModel).ToNot(BeNil())
				Expect(getTransitGatewayConnectionOptionsModel.TransitGatewayID).To(Equal(core.StringPtr("testString")))
				Expect(getTransitGatewayConnectionOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(getTransitGatewayConnectionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetTransitGatewayOptions successfully`, func() {
				// Construct an instance of the GetTransitGatewayOptions model
				id := "testString"
				getTransitGatewayOptionsModel := transitGatewayApisService.NewGetTransitGatewayOptions(id)
				getTransitGatewayOptionsModel.SetID("testString")
				getTransitGatewayOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getTransitGatewayOptionsModel).ToNot(BeNil())
				Expect(getTransitGatewayOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(getTransitGatewayOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListConnectionsOptions successfully`, func() {
				// Construct an instance of the ListConnectionsOptions model
				listConnectionsOptionsModel := transitGatewayApisService.NewListConnectionsOptions()
				listConnectionsOptionsModel.SetLimit(int64(1))
				listConnectionsOptionsModel.SetStart("testString")
				listConnectionsOptionsModel.SetNetworkID("testString")
				listConnectionsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listConnectionsOptionsModel).ToNot(BeNil())
				Expect(listConnectionsOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(1))))
				Expect(listConnectionsOptionsModel.Start).To(Equal(core.StringPtr("testString")))
				Expect(listConnectionsOptionsModel.NetworkID).To(Equal(core.StringPtr("testString")))
				Expect(listConnectionsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListGatewayLocationsOptions successfully`, func() {
				// Construct an instance of the ListGatewayLocationsOptions model
				listGatewayLocationsOptionsModel := transitGatewayApisService.NewListGatewayLocationsOptions()
				listGatewayLocationsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listGatewayLocationsOptionsModel).ToNot(BeNil())
				Expect(listGatewayLocationsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListTransitGatewayConnectionsOptions successfully`, func() {
				// Construct an instance of the ListTransitGatewayConnectionsOptions model
				transitGatewayID := "testString"
				listTransitGatewayConnectionsOptionsModel := transitGatewayApisService.NewListTransitGatewayConnectionsOptions(transitGatewayID)
				listTransitGatewayConnectionsOptionsModel.SetTransitGatewayID("testString")
				listTransitGatewayConnectionsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listTransitGatewayConnectionsOptionsModel).ToNot(BeNil())
				Expect(listTransitGatewayConnectionsOptionsModel.TransitGatewayID).To(Equal(core.StringPtr("testString")))
				Expect(listTransitGatewayConnectionsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListTransitGatewaysOptions successfully`, func() {
				// Construct an instance of the ListTransitGatewaysOptions model
				listTransitGatewaysOptionsModel := transitGatewayApisService.NewListTransitGatewaysOptions()
				listTransitGatewaysOptionsModel.SetLimit(int64(1))
				listTransitGatewaysOptionsModel.SetStart("testString")
				listTransitGatewaysOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listTransitGatewaysOptionsModel).ToNot(BeNil())
				Expect(listTransitGatewaysOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(1))))
				Expect(listTransitGatewaysOptionsModel.Start).To(Equal(core.StringPtr("testString")))
				Expect(listTransitGatewaysOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewResourceGroupIdentity successfully`, func() {
				id := "56969d6043e9465c883cb9f7363e78e8"
				model, err := transitGatewayApisService.NewResourceGroupIdentity(id)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewUpdateTransitGatewayConnectionOptions successfully`, func() {
				// Construct an instance of the UpdateTransitGatewayConnectionOptions model
				transitGatewayID := "testString"
				id := "testString"
				updateTransitGatewayConnectionOptionsModel := transitGatewayApisService.NewUpdateTransitGatewayConnectionOptions(transitGatewayID, id)
				updateTransitGatewayConnectionOptionsModel.SetTransitGatewayID("testString")
				updateTransitGatewayConnectionOptionsModel.SetID("testString")
				updateTransitGatewayConnectionOptionsModel.SetName("Transit_Service_BWTN_SJ_DL")
				updateTransitGatewayConnectionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateTransitGatewayConnectionOptionsModel).ToNot(BeNil())
				Expect(updateTransitGatewayConnectionOptionsModel.TransitGatewayID).To(Equal(core.StringPtr("testString")))
				Expect(updateTransitGatewayConnectionOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(updateTransitGatewayConnectionOptionsModel.Name).To(Equal(core.StringPtr("Transit_Service_BWTN_SJ_DL")))
				Expect(updateTransitGatewayConnectionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateTransitGatewayOptions successfully`, func() {
				// Construct an instance of the UpdateTransitGatewayOptions model
				id := "testString"
				updateTransitGatewayOptionsModel := transitGatewayApisService.NewUpdateTransitGatewayOptions(id)
				updateTransitGatewayOptionsModel.SetID("testString")
				updateTransitGatewayOptionsModel.SetGlobal(true)
				updateTransitGatewayOptionsModel.SetName("my-transit-gateway")
				updateTransitGatewayOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateTransitGatewayOptionsModel).ToNot(BeNil())
				Expect(updateTransitGatewayOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(updateTransitGatewayOptionsModel.Global).To(Equal(core.BoolPtr(true)))
				Expect(updateTransitGatewayOptionsModel.Name).To(Equal(core.StringPtr("my-transit-gateway")))
				Expect(updateTransitGatewayOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
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
