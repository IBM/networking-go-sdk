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

package transitgatewayapisv1_test

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/networking-go-sdk/transitgatewayapisv1"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
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

				clone := transitGatewayApisService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != transitGatewayApisService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(transitGatewayApisService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(transitGatewayApisService.Service.Options.Authenticator))
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

				clone := transitGatewayApisService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != transitGatewayApisService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(transitGatewayApisService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(transitGatewayApisService.Service.Options.Authenticator))
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

				clone := transitGatewayApisService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != transitGatewayApisService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(transitGatewayApisService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(transitGatewayApisService.Service.Options.Authenticator))
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
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = transitgatewayapisv1.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`ListTransitGateways(listTransitGatewaysOptions *ListTransitGatewaysOptions) - Operation response error`, func() {
		version := "testString"
		listTransitGatewaysPath := "/transit_gateways"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listTransitGatewaysPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
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
				listTransitGatewaysOptionsModel.Limit = core.Int64Ptr(int64(10))
				listTransitGatewaysOptionsModel.Start = core.StringPtr("testString")
				listTransitGatewaysOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := transitGatewayApisService.ListTransitGateways(listTransitGatewaysOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				transitGatewayApisService.EnableRetries(0, 0)
				result, response, operationErr = transitGatewayApisService.ListTransitGateways(listTransitGatewaysOptionsModel)
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
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listTransitGatewaysPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"first": {"href": "https://transit.cloud.ibm.com/v1/transit_gateways?limit=50"}, "limit": 50, "next": {"href": "https://transit.cloud.ibm.com/v1/transit_gateways?start=MjAyMC0wNS0wOFQxNDoxNzowMy45NzQ5NzNa&limit=50", "start": "MjAyMC0wNS0wOFQxNDoxNzowMy45NzQ5NzNa"}, "transit_gateways": [{"connection_count": 5, "connection_needs_attention": true, "created_at": "2019-01-01T12:00:00.000Z", "crn": "crn:v1:bluemix:public:transit:dal03:a/57a7d05f36894e3cb9b46a43556d903e::gateway:ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4", "global": true, "gre_enhanced_route_propagation": true, "id": "0a06fb9b-820f-4c44-8a31-77f1f0806d28", "location": "us-south", "name": "my-transit-gateway-in-TransitGateway", "resource_group": {"href": "https://resource-manager.bluemix.net/v1/resource_groups/56969d6043e9465c883cb9f7363e78e8", "id": "56969d6043e9465c883cb9f7363e78e8"}, "status": "available", "updated_at": "2019-01-01T12:00:00.000Z"}]}`)
				}))
			})
			It(`Invoke ListTransitGateways successfully with retries`, func() {
				transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(transitGatewayApisService).ToNot(BeNil())
				transitGatewayApisService.EnableRetries(0, 0)

				// Construct an instance of the ListTransitGatewaysOptions model
				listTransitGatewaysOptionsModel := new(transitgatewayapisv1.ListTransitGatewaysOptions)
				listTransitGatewaysOptionsModel.Limit = core.Int64Ptr(int64(10))
				listTransitGatewaysOptionsModel.Start = core.StringPtr("testString")
				listTransitGatewaysOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := transitGatewayApisService.ListTransitGatewaysWithContext(ctx, listTransitGatewaysOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				transitGatewayApisService.DisableRetries()
				result, response, operationErr := transitGatewayApisService.ListTransitGateways(listTransitGatewaysOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = transitGatewayApisService.ListTransitGatewaysWithContext(ctx, listTransitGatewaysOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listTransitGatewaysPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"first": {"href": "https://transit.cloud.ibm.com/v1/transit_gateways?limit=50"}, "limit": 50, "next": {"href": "https://transit.cloud.ibm.com/v1/transit_gateways?start=MjAyMC0wNS0wOFQxNDoxNzowMy45NzQ5NzNa&limit=50", "start": "MjAyMC0wNS0wOFQxNDoxNzowMy45NzQ5NzNa"}, "transit_gateways": [{"connection_count": 5, "connection_needs_attention": true, "created_at": "2019-01-01T12:00:00.000Z", "crn": "crn:v1:bluemix:public:transit:dal03:a/57a7d05f36894e3cb9b46a43556d903e::gateway:ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4", "global": true, "gre_enhanced_route_propagation": true, "id": "0a06fb9b-820f-4c44-8a31-77f1f0806d28", "location": "us-south", "name": "my-transit-gateway-in-TransitGateway", "resource_group": {"href": "https://resource-manager.bluemix.net/v1/resource_groups/56969d6043e9465c883cb9f7363e78e8", "id": "56969d6043e9465c883cb9f7363e78e8"}, "status": "available", "updated_at": "2019-01-01T12:00:00.000Z"}]}`)
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
				listTransitGatewaysOptionsModel.Limit = core.Int64Ptr(int64(10))
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
				listTransitGatewaysOptionsModel.Limit = core.Int64Ptr(int64(10))
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
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
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

				// Construct an instance of the ListTransitGatewaysOptions model
				listTransitGatewaysOptionsModel := new(transitgatewayapisv1.ListTransitGatewaysOptions)
				listTransitGatewaysOptionsModel.Limit = core.Int64Ptr(int64(10))
				listTransitGatewaysOptionsModel.Start = core.StringPtr("testString")
				listTransitGatewaysOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := transitGatewayApisService.ListTransitGateways(listTransitGatewaysOptionsModel)
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
				responseObject := new(transitgatewayapisv1.TransitGatewayCollection)
				nextObject := new(transitgatewayapisv1.PaginationNextTG)
				nextObject.Start = core.StringPtr("abc-123")
				responseObject.Next = nextObject
	
				value, err := responseObject.GetNextStart()
				Expect(err).To(BeNil())
				Expect(value).To(Equal(core.StringPtr("abc-123")))
			})
			It(`Invoke GetNextStart without a "Next" property in the response`, func() {
				responseObject := new(transitgatewayapisv1.TransitGatewayCollection)
	
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
					Expect(req.URL.EscapedPath()).To(Equal(listTransitGatewaysPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					requestNumber++
					if requestNumber == 1 {
						fmt.Fprintf(res, "%s", `{"next":{"start":"1"},"transit_gateways":[{"connection_count":5,"connection_needs_attention":true,"created_at":"2019-01-01T12:00:00.000Z","crn":"crn:v1:bluemix:public:transit:dal03:a/57a7d05f36894e3cb9b46a43556d903e::gateway:ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4","global":true,"gre_enhanced_route_propagation":true,"id":"0a06fb9b-820f-4c44-8a31-77f1f0806d28","location":"us-south","name":"my-transit-gateway-in-TransitGateway","resource_group":{"href":"https://resource-manager.bluemix.net/v1/resource_groups/56969d6043e9465c883cb9f7363e78e8","id":"56969d6043e9465c883cb9f7363e78e8"},"status":"available","updated_at":"2019-01-01T12:00:00.000Z"}],"total_count":2,"limit":1}`)
					} else if requestNumber == 2 {
						fmt.Fprintf(res, "%s", `{"transit_gateways":[{"connection_count":5,"connection_needs_attention":true,"created_at":"2019-01-01T12:00:00.000Z","crn":"crn:v1:bluemix:public:transit:dal03:a/57a7d05f36894e3cb9b46a43556d903e::gateway:ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4","global":true,"gre_enhanced_route_propagation":true,"id":"0a06fb9b-820f-4c44-8a31-77f1f0806d28","location":"us-south","name":"my-transit-gateway-in-TransitGateway","resource_group":{"href":"https://resource-manager.bluemix.net/v1/resource_groups/56969d6043e9465c883cb9f7363e78e8","id":"56969d6043e9465c883cb9f7363e78e8"},"status":"available","updated_at":"2019-01-01T12:00:00.000Z"}],"total_count":2,"limit":1}`)
					} else {
						res.WriteHeader(400)
					}
				}))
			})
			It(`Use TransitGatewaysPager.GetNext successfully`, func() {
				transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(transitGatewayApisService).ToNot(BeNil())

				listTransitGatewaysOptionsModel := &transitgatewayapisv1.ListTransitGatewaysOptions{
					Limit: core.Int64Ptr(int64(10)),
				}

				pager, err := transitGatewayApisService.NewTransitGatewaysPager(listTransitGatewaysOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				var allResults []transitgatewayapisv1.TransitGateway
				for pager.HasNext() {
					nextPage, err := pager.GetNext()
					Expect(err).To(BeNil())
					Expect(nextPage).ToNot(BeNil())
					allResults = append(allResults, nextPage...)
				}
				Expect(len(allResults)).To(Equal(2))
			})
			It(`Use TransitGatewaysPager.GetAll successfully`, func() {
				transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(transitGatewayApisService).ToNot(BeNil())

				listTransitGatewaysOptionsModel := &transitgatewayapisv1.ListTransitGatewaysOptions{
					Limit: core.Int64Ptr(int64(10)),
				}

				pager, err := transitGatewayApisService.NewTransitGatewaysPager(listTransitGatewaysOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				allResults, err := pager.GetAll()
				Expect(err).To(BeNil())
				Expect(allResults).ToNot(BeNil())
				Expect(len(allResults)).To(Equal(2))
			})
		})
	})
	Describe(`CreateTransitGateway(createTransitGatewayOptions *CreateTransitGatewayOptions) - Operation response error`, func() {
		version := "testString"
		createTransitGatewayPath := "/transit_gateways"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createTransitGatewayPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
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
				createTransitGatewayOptionsModel.Name = core.StringPtr("my-transit-gateway-in-TransitGateway")
				createTransitGatewayOptionsModel.Global = core.BoolPtr(true)
				createTransitGatewayOptionsModel.GreEnhancedRoutePropagation = core.BoolPtr(true)
				createTransitGatewayOptionsModel.ResourceGroup = resourceGroupIdentityModel
				createTransitGatewayOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := transitGatewayApisService.CreateTransitGateway(createTransitGatewayOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				transitGatewayApisService.EnableRetries(0, 0)
				result, response, operationErr = transitGatewayApisService.CreateTransitGateway(createTransitGatewayOptionsModel)
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
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createTransitGatewayPath))
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
					fmt.Fprintf(res, "%s", `{"connection_count": 5, "connection_needs_attention": true, "created_at": "2019-01-01T12:00:00.000Z", "crn": "crn:v1:bluemix:public:transit:dal03:a/57a7d05f36894e3cb9b46a43556d903e::gateway:ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4", "global": true, "gre_enhanced_route_propagation": true, "id": "0a06fb9b-820f-4c44-8a31-77f1f0806d28", "location": "us-south", "name": "my-transit-gateway-in-TransitGateway", "resource_group": {"href": "https://resource-manager.bluemix.net/v1/resource_groups/56969d6043e9465c883cb9f7363e78e8", "id": "56969d6043e9465c883cb9f7363e78e8"}, "status": "available", "updated_at": "2019-01-01T12:00:00.000Z"}`)
				}))
			})
			It(`Invoke CreateTransitGateway successfully with retries`, func() {
				transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(transitGatewayApisService).ToNot(BeNil())
				transitGatewayApisService.EnableRetries(0, 0)

				// Construct an instance of the ResourceGroupIdentity model
				resourceGroupIdentityModel := new(transitgatewayapisv1.ResourceGroupIdentity)
				resourceGroupIdentityModel.ID = core.StringPtr("56969d6043e9465c883cb9f7363e78e8")

				// Construct an instance of the CreateTransitGatewayOptions model
				createTransitGatewayOptionsModel := new(transitgatewayapisv1.CreateTransitGatewayOptions)
				createTransitGatewayOptionsModel.Location = core.StringPtr("us-south")
				createTransitGatewayOptionsModel.Name = core.StringPtr("my-transit-gateway-in-TransitGateway")
				createTransitGatewayOptionsModel.Global = core.BoolPtr(true)
				createTransitGatewayOptionsModel.GreEnhancedRoutePropagation = core.BoolPtr(true)
				createTransitGatewayOptionsModel.ResourceGroup = resourceGroupIdentityModel
				createTransitGatewayOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := transitGatewayApisService.CreateTransitGatewayWithContext(ctx, createTransitGatewayOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				transitGatewayApisService.DisableRetries()
				result, response, operationErr := transitGatewayApisService.CreateTransitGateway(createTransitGatewayOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = transitGatewayApisService.CreateTransitGatewayWithContext(ctx, createTransitGatewayOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createTransitGatewayPath))
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
					fmt.Fprintf(res, "%s", `{"connection_count": 5, "connection_needs_attention": true, "created_at": "2019-01-01T12:00:00.000Z", "crn": "crn:v1:bluemix:public:transit:dal03:a/57a7d05f36894e3cb9b46a43556d903e::gateway:ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4", "global": true, "gre_enhanced_route_propagation": true, "id": "0a06fb9b-820f-4c44-8a31-77f1f0806d28", "location": "us-south", "name": "my-transit-gateway-in-TransitGateway", "resource_group": {"href": "https://resource-manager.bluemix.net/v1/resource_groups/56969d6043e9465c883cb9f7363e78e8", "id": "56969d6043e9465c883cb9f7363e78e8"}, "status": "available", "updated_at": "2019-01-01T12:00:00.000Z"}`)
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
				createTransitGatewayOptionsModel.Name = core.StringPtr("my-transit-gateway-in-TransitGateway")
				createTransitGatewayOptionsModel.Global = core.BoolPtr(true)
				createTransitGatewayOptionsModel.GreEnhancedRoutePropagation = core.BoolPtr(true)
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
				createTransitGatewayOptionsModel.Name = core.StringPtr("my-transit-gateway-in-TransitGateway")
				createTransitGatewayOptionsModel.Global = core.BoolPtr(true)
				createTransitGatewayOptionsModel.GreEnhancedRoutePropagation = core.BoolPtr(true)
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
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(201)
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

				// Construct an instance of the ResourceGroupIdentity model
				resourceGroupIdentityModel := new(transitgatewayapisv1.ResourceGroupIdentity)
				resourceGroupIdentityModel.ID = core.StringPtr("56969d6043e9465c883cb9f7363e78e8")

				// Construct an instance of the CreateTransitGatewayOptions model
				createTransitGatewayOptionsModel := new(transitgatewayapisv1.CreateTransitGatewayOptions)
				createTransitGatewayOptionsModel.Location = core.StringPtr("us-south")
				createTransitGatewayOptionsModel.Name = core.StringPtr("my-transit-gateway-in-TransitGateway")
				createTransitGatewayOptionsModel.Global = core.BoolPtr(true)
				createTransitGatewayOptionsModel.GreEnhancedRoutePropagation = core.BoolPtr(true)
				createTransitGatewayOptionsModel.ResourceGroup = resourceGroupIdentityModel
				createTransitGatewayOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := transitGatewayApisService.CreateTransitGateway(createTransitGatewayOptionsModel)
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
	Describe(`DeleteTransitGateway(deleteTransitGatewayOptions *DeleteTransitGatewayOptions)`, func() {
		version := "testString"
		deleteTransitGatewayPath := "/transit_gateways/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteTransitGatewayPath))
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
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getTransitGatewayPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
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

				// Enable retries and test again
				transitGatewayApisService.EnableRetries(0, 0)
				result, response, operationErr = transitGatewayApisService.GetTransitGateway(getTransitGatewayOptionsModel)
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
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getTransitGatewayPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"connection_count": 5, "connection_needs_attention": true, "created_at": "2019-01-01T12:00:00.000Z", "crn": "crn:v1:bluemix:public:transit:dal03:a/57a7d05f36894e3cb9b46a43556d903e::gateway:ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4", "global": true, "gre_enhanced_route_propagation": true, "id": "0a06fb9b-820f-4c44-8a31-77f1f0806d28", "location": "us-south", "name": "my-transit-gateway-in-TransitGateway", "resource_group": {"href": "https://resource-manager.bluemix.net/v1/resource_groups/56969d6043e9465c883cb9f7363e78e8", "id": "56969d6043e9465c883cb9f7363e78e8"}, "status": "available", "updated_at": "2019-01-01T12:00:00.000Z"}`)
				}))
			})
			It(`Invoke GetTransitGateway successfully with retries`, func() {
				transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(transitGatewayApisService).ToNot(BeNil())
				transitGatewayApisService.EnableRetries(0, 0)

				// Construct an instance of the GetTransitGatewayOptions model
				getTransitGatewayOptionsModel := new(transitgatewayapisv1.GetTransitGatewayOptions)
				getTransitGatewayOptionsModel.ID = core.StringPtr("testString")
				getTransitGatewayOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := transitGatewayApisService.GetTransitGatewayWithContext(ctx, getTransitGatewayOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				transitGatewayApisService.DisableRetries()
				result, response, operationErr := transitGatewayApisService.GetTransitGateway(getTransitGatewayOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = transitGatewayApisService.GetTransitGatewayWithContext(ctx, getTransitGatewayOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getTransitGatewayPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"connection_count": 5, "connection_needs_attention": true, "created_at": "2019-01-01T12:00:00.000Z", "crn": "crn:v1:bluemix:public:transit:dal03:a/57a7d05f36894e3cb9b46a43556d903e::gateway:ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4", "global": true, "gre_enhanced_route_propagation": true, "id": "0a06fb9b-820f-4c44-8a31-77f1f0806d28", "location": "us-south", "name": "my-transit-gateway-in-TransitGateway", "resource_group": {"href": "https://resource-manager.bluemix.net/v1/resource_groups/56969d6043e9465c883cb9f7363e78e8", "id": "56969d6043e9465c883cb9f7363e78e8"}, "status": "available", "updated_at": "2019-01-01T12:00:00.000Z"}`)
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
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
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

				// Construct an instance of the GetTransitGatewayOptions model
				getTransitGatewayOptionsModel := new(transitgatewayapisv1.GetTransitGatewayOptions)
				getTransitGatewayOptionsModel.ID = core.StringPtr("testString")
				getTransitGatewayOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := transitGatewayApisService.GetTransitGateway(getTransitGatewayOptionsModel)
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
	Describe(`UpdateTransitGateway(updateTransitGatewayOptions *UpdateTransitGatewayOptions) - Operation response error`, func() {
		version := "testString"
		updateTransitGatewayPath := "/transit_gateways/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateTransitGatewayPath))
					Expect(req.Method).To(Equal("PATCH"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
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
				updateTransitGatewayOptionsModel.GreEnhancedRoutePropagation = core.BoolPtr(true)
				updateTransitGatewayOptionsModel.Name = core.StringPtr("my-resource")
				updateTransitGatewayOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := transitGatewayApisService.UpdateTransitGateway(updateTransitGatewayOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				transitGatewayApisService.EnableRetries(0, 0)
				result, response, operationErr = transitGatewayApisService.UpdateTransitGateway(updateTransitGatewayOptionsModel)
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
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateTransitGatewayPath))
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
					fmt.Fprintf(res, "%s", `{"connection_count": 5, "connection_needs_attention": true, "created_at": "2019-01-01T12:00:00.000Z", "crn": "crn:v1:bluemix:public:transit:dal03:a/57a7d05f36894e3cb9b46a43556d903e::gateway:ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4", "global": true, "gre_enhanced_route_propagation": true, "id": "0a06fb9b-820f-4c44-8a31-77f1f0806d28", "location": "us-south", "name": "my-transit-gateway-in-TransitGateway", "resource_group": {"href": "https://resource-manager.bluemix.net/v1/resource_groups/56969d6043e9465c883cb9f7363e78e8", "id": "56969d6043e9465c883cb9f7363e78e8"}, "status": "available", "updated_at": "2019-01-01T12:00:00.000Z"}`)
				}))
			})
			It(`Invoke UpdateTransitGateway successfully with retries`, func() {
				transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(transitGatewayApisService).ToNot(BeNil())
				transitGatewayApisService.EnableRetries(0, 0)

				// Construct an instance of the UpdateTransitGatewayOptions model
				updateTransitGatewayOptionsModel := new(transitgatewayapisv1.UpdateTransitGatewayOptions)
				updateTransitGatewayOptionsModel.ID = core.StringPtr("testString")
				updateTransitGatewayOptionsModel.Global = core.BoolPtr(true)
				updateTransitGatewayOptionsModel.GreEnhancedRoutePropagation = core.BoolPtr(true)
				updateTransitGatewayOptionsModel.Name = core.StringPtr("my-resource")
				updateTransitGatewayOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := transitGatewayApisService.UpdateTransitGatewayWithContext(ctx, updateTransitGatewayOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				transitGatewayApisService.DisableRetries()
				result, response, operationErr := transitGatewayApisService.UpdateTransitGateway(updateTransitGatewayOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = transitGatewayApisService.UpdateTransitGatewayWithContext(ctx, updateTransitGatewayOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(updateTransitGatewayPath))
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
					fmt.Fprintf(res, "%s", `{"connection_count": 5, "connection_needs_attention": true, "created_at": "2019-01-01T12:00:00.000Z", "crn": "crn:v1:bluemix:public:transit:dal03:a/57a7d05f36894e3cb9b46a43556d903e::gateway:ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4", "global": true, "gre_enhanced_route_propagation": true, "id": "0a06fb9b-820f-4c44-8a31-77f1f0806d28", "location": "us-south", "name": "my-transit-gateway-in-TransitGateway", "resource_group": {"href": "https://resource-manager.bluemix.net/v1/resource_groups/56969d6043e9465c883cb9f7363e78e8", "id": "56969d6043e9465c883cb9f7363e78e8"}, "status": "available", "updated_at": "2019-01-01T12:00:00.000Z"}`)
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
				updateTransitGatewayOptionsModel.GreEnhancedRoutePropagation = core.BoolPtr(true)
				updateTransitGatewayOptionsModel.Name = core.StringPtr("my-resource")
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
				updateTransitGatewayOptionsModel.GreEnhancedRoutePropagation = core.BoolPtr(true)
				updateTransitGatewayOptionsModel.Name = core.StringPtr("my-resource")
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
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
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

				// Construct an instance of the UpdateTransitGatewayOptions model
				updateTransitGatewayOptionsModel := new(transitgatewayapisv1.UpdateTransitGatewayOptions)
				updateTransitGatewayOptionsModel.ID = core.StringPtr("testString")
				updateTransitGatewayOptionsModel.Global = core.BoolPtr(true)
				updateTransitGatewayOptionsModel.GreEnhancedRoutePropagation = core.BoolPtr(true)
				updateTransitGatewayOptionsModel.Name = core.StringPtr("my-resource")
				updateTransitGatewayOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := transitGatewayApisService.UpdateTransitGateway(updateTransitGatewayOptionsModel)
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
	Describe(`ListConnections(listConnectionsOptions *ListConnectionsOptions) - Operation response error`, func() {
		version := "testString"
		listConnectionsPath := "/connections"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listConnectionsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["network_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["network_type"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
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
				listConnectionsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listConnectionsOptionsModel.Start = core.StringPtr("testString")
				listConnectionsOptionsModel.NetworkID = core.StringPtr("testString")
				listConnectionsOptionsModel.NetworkType = core.StringPtr("testString")
				listConnectionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := transitGatewayApisService.ListConnections(listConnectionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				transitGatewayApisService.EnableRetries(0, 0)
				result, response, operationErr = transitGatewayApisService.ListConnections(listConnectionsOptionsModel)
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
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listConnectionsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["network_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["network_type"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"connections": [{"base_network_type": "classic", "name": "Transit_Service_BWTN_SJ_DL", "network_id": "crn:v1:bluemix:public:is:us-south:a/123456::vpc:4727d842-f94f-4a2d-824a-9bc9b02c523b", "network_type": "vpc", "id": "1a15dca5-7e33-45e1-b7c5-bc690e569531", "base_connection_id": "975f58c1-afe7-469a-9727-7f3d720f2d32", "created_at": "2019-01-01T12:00:00.000Z", "local_bgp_asn": 64490, "local_gateway_ip": "192.168.100.1", "local_tunnel_ip": "192.168.129.2", "mtu": 9000, "network_account_id": "NetworkAccountID", "prefix_filters": [{"action": "permit", "before": "1a15dcab-7e40-45e1-b7c5-bc690eaa9782", "created_at": "2019-01-01T12:00:00.000Z", "ge": 0, "id": "1a15dcab-7e30-45e1-b7c5-bc690eaa9865", "le": 32, "prefix": "192.168.100.0/24", "updated_at": "2019-01-01T12:00:00.000Z"}], "prefix_filters_default": "permit", "remote_bgp_asn": 65010, "remote_gateway_ip": "10.242.63.12", "remote_tunnel_ip": "192.168.129.1", "request_status": "pending", "status": "attached", "transit_gateway": {"crn": "crn:v1:bluemix:public:transit:us-south:a/123456::gateway:456f58c1-afe7-123a-0a0a-7f3d720f1a44", "id": "456f58c1-afe7-123a-0a0a-7f3d720f1a44", "name": "my-transit-gw100"}, "tunnels": [{"base_network_type": "classic", "created_at": "2019-01-01T12:00:00.000Z", "id": "1a15dca5-7e33-45e1-b7c5-bc690e569531", "local_bgp_asn": 11, "local_gateway_ip": "10.242.63.12", "local_tunnel_ip": "192.168.100.20", "mtu": 9000, "name": "gre1", "network_account_id": "NetworkAccountID", "network_id": "crn:v1:bluemix:public:is:us-south:a/123456::vpc:4727d842-f94f-4a2d-824a-9bc9b02c523b", "remote_bgp_asn": 65010, "remote_gateway_ip": "10.242.33.22", "remote_tunnel_ip": "192.168.129.1", "status": "attached", "updated_at": "2019-01-01T12:00:00.000Z", "zone": {"name": "us-south-1"}}], "updated_at": "2019-01-01T12:00:00.000Z", "zone": {"name": "us-south-1"}}], "first": {"href": "https://transit.cloud.ibm.com/v1/connections?limit=50"}, "limit": 50, "next": {"href": "https://transit.cloud.ibm.com/v1/connections?start=MjAyMC0wNS0wOFQxNDoxNzowMy45NzQ5NzNa&limit=50", "start": "MjAyMC0wNS0wOFQxNDoxNzowMy45NzQ5NzNa"}}`)
				}))
			})
			It(`Invoke ListConnections successfully with retries`, func() {
				transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(transitGatewayApisService).ToNot(BeNil())
				transitGatewayApisService.EnableRetries(0, 0)

				// Construct an instance of the ListConnectionsOptions model
				listConnectionsOptionsModel := new(transitgatewayapisv1.ListConnectionsOptions)
				listConnectionsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listConnectionsOptionsModel.Start = core.StringPtr("testString")
				listConnectionsOptionsModel.NetworkID = core.StringPtr("testString")
				listConnectionsOptionsModel.NetworkType = core.StringPtr("testString")
				listConnectionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := transitGatewayApisService.ListConnectionsWithContext(ctx, listConnectionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				transitGatewayApisService.DisableRetries()
				result, response, operationErr := transitGatewayApisService.ListConnections(listConnectionsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = transitGatewayApisService.ListConnectionsWithContext(ctx, listConnectionsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listConnectionsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["network_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["network_type"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"connections": [{"base_network_type": "classic", "name": "Transit_Service_BWTN_SJ_DL", "network_id": "crn:v1:bluemix:public:is:us-south:a/123456::vpc:4727d842-f94f-4a2d-824a-9bc9b02c523b", "network_type": "vpc", "id": "1a15dca5-7e33-45e1-b7c5-bc690e569531", "base_connection_id": "975f58c1-afe7-469a-9727-7f3d720f2d32", "created_at": "2019-01-01T12:00:00.000Z", "local_bgp_asn": 64490, "local_gateway_ip": "192.168.100.1", "local_tunnel_ip": "192.168.129.2", "mtu": 9000, "network_account_id": "NetworkAccountID", "prefix_filters": [{"action": "permit", "before": "1a15dcab-7e40-45e1-b7c5-bc690eaa9782", "created_at": "2019-01-01T12:00:00.000Z", "ge": 0, "id": "1a15dcab-7e30-45e1-b7c5-bc690eaa9865", "le": 32, "prefix": "192.168.100.0/24", "updated_at": "2019-01-01T12:00:00.000Z"}], "prefix_filters_default": "permit", "remote_bgp_asn": 65010, "remote_gateway_ip": "10.242.63.12", "remote_tunnel_ip": "192.168.129.1", "request_status": "pending", "status": "attached", "transit_gateway": {"crn": "crn:v1:bluemix:public:transit:us-south:a/123456::gateway:456f58c1-afe7-123a-0a0a-7f3d720f1a44", "id": "456f58c1-afe7-123a-0a0a-7f3d720f1a44", "name": "my-transit-gw100"}, "tunnels": [{"base_network_type": "classic", "created_at": "2019-01-01T12:00:00.000Z", "id": "1a15dca5-7e33-45e1-b7c5-bc690e569531", "local_bgp_asn": 11, "local_gateway_ip": "10.242.63.12", "local_tunnel_ip": "192.168.100.20", "mtu": 9000, "name": "gre1", "network_account_id": "NetworkAccountID", "network_id": "crn:v1:bluemix:public:is:us-south:a/123456::vpc:4727d842-f94f-4a2d-824a-9bc9b02c523b", "remote_bgp_asn": 65010, "remote_gateway_ip": "10.242.33.22", "remote_tunnel_ip": "192.168.129.1", "status": "attached", "updated_at": "2019-01-01T12:00:00.000Z", "zone": {"name": "us-south-1"}}], "updated_at": "2019-01-01T12:00:00.000Z", "zone": {"name": "us-south-1"}}], "first": {"href": "https://transit.cloud.ibm.com/v1/connections?limit=50"}, "limit": 50, "next": {"href": "https://transit.cloud.ibm.com/v1/connections?start=MjAyMC0wNS0wOFQxNDoxNzowMy45NzQ5NzNa&limit=50", "start": "MjAyMC0wNS0wOFQxNDoxNzowMy45NzQ5NzNa"}}`)
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
				listConnectionsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listConnectionsOptionsModel.Start = core.StringPtr("testString")
				listConnectionsOptionsModel.NetworkID = core.StringPtr("testString")
				listConnectionsOptionsModel.NetworkType = core.StringPtr("testString")
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
				listConnectionsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listConnectionsOptionsModel.Start = core.StringPtr("testString")
				listConnectionsOptionsModel.NetworkID = core.StringPtr("testString")
				listConnectionsOptionsModel.NetworkType = core.StringPtr("testString")
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
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
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

				// Construct an instance of the ListConnectionsOptions model
				listConnectionsOptionsModel := new(transitgatewayapisv1.ListConnectionsOptions)
				listConnectionsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listConnectionsOptionsModel.Start = core.StringPtr("testString")
				listConnectionsOptionsModel.NetworkID = core.StringPtr("testString")
				listConnectionsOptionsModel.NetworkType = core.StringPtr("testString")
				listConnectionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := transitGatewayApisService.ListConnections(listConnectionsOptionsModel)
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
				responseObject := new(transitgatewayapisv1.TransitConnectionCollection)
				nextObject := new(transitgatewayapisv1.PaginationNextConnection)
				nextObject.Start = core.StringPtr("abc-123")
				responseObject.Next = nextObject
	
				value, err := responseObject.GetNextStart()
				Expect(err).To(BeNil())
				Expect(value).To(Equal(core.StringPtr("abc-123")))
			})
			It(`Invoke GetNextStart without a "Next" property in the response`, func() {
				responseObject := new(transitgatewayapisv1.TransitConnectionCollection)
	
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
					Expect(req.URL.EscapedPath()).To(Equal(listConnectionsPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					requestNumber++
					if requestNumber == 1 {
						fmt.Fprintf(res, "%s", `{"next":{"start":"1"},"total_count":2,"limit":1,"connections":[{"base_network_type":"classic","name":"Transit_Service_BWTN_SJ_DL","network_id":"crn:v1:bluemix:public:is:us-south:a/123456::vpc:4727d842-f94f-4a2d-824a-9bc9b02c523b","network_type":"vpc","id":"1a15dca5-7e33-45e1-b7c5-bc690e569531","base_connection_id":"975f58c1-afe7-469a-9727-7f3d720f2d32","created_at":"2019-01-01T12:00:00.000Z","local_bgp_asn":64490,"local_gateway_ip":"192.168.100.1","local_tunnel_ip":"192.168.129.2","mtu":9000,"network_account_id":"NetworkAccountID","prefix_filters":[{"action":"permit","before":"1a15dcab-7e40-45e1-b7c5-bc690eaa9782","created_at":"2019-01-01T12:00:00.000Z","ge":0,"id":"1a15dcab-7e30-45e1-b7c5-bc690eaa9865","le":32,"prefix":"192.168.100.0/24","updated_at":"2019-01-01T12:00:00.000Z"}],"prefix_filters_default":"permit","remote_bgp_asn":65010,"remote_gateway_ip":"10.242.63.12","remote_tunnel_ip":"192.168.129.1","request_status":"pending","status":"attached","transit_gateway":{"crn":"crn:v1:bluemix:public:transit:us-south:a/123456::gateway:456f58c1-afe7-123a-0a0a-7f3d720f1a44","id":"456f58c1-afe7-123a-0a0a-7f3d720f1a44","name":"my-transit-gw100"},"tunnels":[{"base_network_type":"classic","created_at":"2019-01-01T12:00:00.000Z","id":"1a15dca5-7e33-45e1-b7c5-bc690e569531","local_bgp_asn":11,"local_gateway_ip":"10.242.63.12","local_tunnel_ip":"192.168.100.20","mtu":9000,"name":"gre1","network_account_id":"NetworkAccountID","network_id":"crn:v1:bluemix:public:is:us-south:a/123456::vpc:4727d842-f94f-4a2d-824a-9bc9b02c523b","remote_bgp_asn":65010,"remote_gateway_ip":"10.242.33.22","remote_tunnel_ip":"192.168.129.1","status":"attached","updated_at":"2019-01-01T12:00:00.000Z","zone":{"name":"us-south-1"}}],"updated_at":"2019-01-01T12:00:00.000Z","zone":{"name":"us-south-1"}}]}`)
					} else if requestNumber == 2 {
						fmt.Fprintf(res, "%s", `{"total_count":2,"limit":1,"connections":[{"base_network_type":"classic","name":"Transit_Service_BWTN_SJ_DL","network_id":"crn:v1:bluemix:public:is:us-south:a/123456::vpc:4727d842-f94f-4a2d-824a-9bc9b02c523b","network_type":"vpc","id":"1a15dca5-7e33-45e1-b7c5-bc690e569531","base_connection_id":"975f58c1-afe7-469a-9727-7f3d720f2d32","created_at":"2019-01-01T12:00:00.000Z","local_bgp_asn":64490,"local_gateway_ip":"192.168.100.1","local_tunnel_ip":"192.168.129.2","mtu":9000,"network_account_id":"NetworkAccountID","prefix_filters":[{"action":"permit","before":"1a15dcab-7e40-45e1-b7c5-bc690eaa9782","created_at":"2019-01-01T12:00:00.000Z","ge":0,"id":"1a15dcab-7e30-45e1-b7c5-bc690eaa9865","le":32,"prefix":"192.168.100.0/24","updated_at":"2019-01-01T12:00:00.000Z"}],"prefix_filters_default":"permit","remote_bgp_asn":65010,"remote_gateway_ip":"10.242.63.12","remote_tunnel_ip":"192.168.129.1","request_status":"pending","status":"attached","transit_gateway":{"crn":"crn:v1:bluemix:public:transit:us-south:a/123456::gateway:456f58c1-afe7-123a-0a0a-7f3d720f1a44","id":"456f58c1-afe7-123a-0a0a-7f3d720f1a44","name":"my-transit-gw100"},"tunnels":[{"base_network_type":"classic","created_at":"2019-01-01T12:00:00.000Z","id":"1a15dca5-7e33-45e1-b7c5-bc690e569531","local_bgp_asn":11,"local_gateway_ip":"10.242.63.12","local_tunnel_ip":"192.168.100.20","mtu":9000,"name":"gre1","network_account_id":"NetworkAccountID","network_id":"crn:v1:bluemix:public:is:us-south:a/123456::vpc:4727d842-f94f-4a2d-824a-9bc9b02c523b","remote_bgp_asn":65010,"remote_gateway_ip":"10.242.33.22","remote_tunnel_ip":"192.168.129.1","status":"attached","updated_at":"2019-01-01T12:00:00.000Z","zone":{"name":"us-south-1"}}],"updated_at":"2019-01-01T12:00:00.000Z","zone":{"name":"us-south-1"}}]}`)
					} else {
						res.WriteHeader(400)
					}
				}))
			})
			It(`Use ConnectionsPager.GetNext successfully`, func() {
				transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(transitGatewayApisService).ToNot(BeNil())

				listConnectionsOptionsModel := &transitgatewayapisv1.ListConnectionsOptions{
					Limit: core.Int64Ptr(int64(10)),
					NetworkID: core.StringPtr("testString"),
					NetworkType: core.StringPtr("testString"),
				}

				pager, err := transitGatewayApisService.NewConnectionsPager(listConnectionsOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				var allResults []transitgatewayapisv1.TransitConnection
				for pager.HasNext() {
					nextPage, err := pager.GetNext()
					Expect(err).To(BeNil())
					Expect(nextPage).ToNot(BeNil())
					allResults = append(allResults, nextPage...)
				}
				Expect(len(allResults)).To(Equal(2))
			})
			It(`Use ConnectionsPager.GetAll successfully`, func() {
				transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(transitGatewayApisService).ToNot(BeNil())

				listConnectionsOptionsModel := &transitgatewayapisv1.ListConnectionsOptions{
					Limit: core.Int64Ptr(int64(10)),
					NetworkID: core.StringPtr("testString"),
					NetworkType: core.StringPtr("testString"),
				}

				pager, err := transitGatewayApisService.NewConnectionsPager(listConnectionsOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				allResults, err := pager.GetAll()
				Expect(err).To(BeNil())
				Expect(allResults).ToNot(BeNil())
				Expect(len(allResults)).To(Equal(2))
			})
		})
	})
	Describe(`ListTransitGatewayConnections(listTransitGatewayConnectionsOptions *ListTransitGatewayConnectionsOptions) - Operation response error`, func() {
		version := "testString"
		listTransitGatewayConnectionsPath := "/transit_gateways/testString/connections"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listTransitGatewayConnectionsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["name"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
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
				listTransitGatewayConnectionsOptionsModel.Start = core.StringPtr("testString")
				listTransitGatewayConnectionsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listTransitGatewayConnectionsOptionsModel.Name = core.StringPtr("testString")
				listTransitGatewayConnectionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := transitGatewayApisService.ListTransitGatewayConnections(listTransitGatewayConnectionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				transitGatewayApisService.EnableRetries(0, 0)
				result, response, operationErr = transitGatewayApisService.ListTransitGatewayConnections(listTransitGatewayConnectionsOptionsModel)
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
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listTransitGatewayConnectionsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["name"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"connections": [{"base_connection_id": "975f58c1-afe7-469a-9727-7f3d720f2d32", "base_network_type": "classic", "cidr": "192.168.0.0/24", "created_at": "2019-01-01T12:00:00.000Z", "id": "1a15dca5-7e33-45e1-b7c5-bc690e569531", "local_bgp_asn": 64490, "local_gateway_ip": "192.168.100.1", "local_tunnel_ip": "192.168.129.2", "mtu": 9000, "name": "Transit_Service_BWTN_SJ_DL", "network_account_id": "NetworkAccountID", "network_id": "crn:v1:bluemix:public:is:us-south:a/123456::vpc:4727d842-f94f-4a2d-824a-9bc9b02c523b", "network_type": "vpc", "prefix_filters": [{"action": "permit", "before": "1a15dcab-7e40-45e1-b7c5-bc690eaa9782", "created_at": "2019-01-01T12:00:00.000Z", "ge": 0, "id": "1a15dcab-7e30-45e1-b7c5-bc690eaa9865", "le": 32, "prefix": "192.168.100.0/24", "updated_at": "2019-01-01T12:00:00.000Z"}], "prefix_filters_default": "permit", "remote_bgp_asn": 65010, "remote_gateway_ip": "10.242.63.12", "remote_tunnel_ip": "192.168.129.1", "request_status": "pending", "status": "attached", "tunnels": [{"base_network_type": "classic", "created_at": "2019-01-01T12:00:00.000Z", "id": "1a15dca5-7e33-45e1-b7c5-bc690e569531", "local_bgp_asn": 11, "local_gateway_ip": "10.242.63.12", "local_tunnel_ip": "192.168.100.20", "mtu": 9000, "name": "gre1", "network_account_id": "NetworkAccountID", "network_id": "crn:v1:bluemix:public:is:us-south:a/123456::vpc:4727d842-f94f-4a2d-824a-9bc9b02c523b", "remote_bgp_asn": 65010, "remote_gateway_ip": "10.242.33.22", "remote_tunnel_ip": "192.168.129.1", "status": "attached", "updated_at": "2019-01-01T12:00:00.000Z", "zone": {"name": "us-south-1"}}], "updated_at": "2019-01-01T12:00:00.000Z", "zone": {"name": "us-south-1"}}], "first": {"href": "https://transit.cloud.ibm.com/v1/transit_gateways/{transit_gateway_id}/connections?limit=50"}, "limit": 50, "next": {"href": "https://transit.cloud.ibm.com/v1/transit_gateways/{transit_gateway_id}/connections?start=MjAyMC0wNS0wOFQxNDoxNzowMy45NzQ5NzNa&limit=50", "start": "MjAyMC0wNS0wOFQxNDoxNzowMy45NzQ5NzNa"}, "total_count": 500}`)
				}))
			})
			It(`Invoke ListTransitGatewayConnections successfully with retries`, func() {
				transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(transitGatewayApisService).ToNot(BeNil())
				transitGatewayApisService.EnableRetries(0, 0)

				// Construct an instance of the ListTransitGatewayConnectionsOptions model
				listTransitGatewayConnectionsOptionsModel := new(transitgatewayapisv1.ListTransitGatewayConnectionsOptions)
				listTransitGatewayConnectionsOptionsModel.TransitGatewayID = core.StringPtr("testString")
				listTransitGatewayConnectionsOptionsModel.Start = core.StringPtr("testString")
				listTransitGatewayConnectionsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listTransitGatewayConnectionsOptionsModel.Name = core.StringPtr("testString")
				listTransitGatewayConnectionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := transitGatewayApisService.ListTransitGatewayConnectionsWithContext(ctx, listTransitGatewayConnectionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				transitGatewayApisService.DisableRetries()
				result, response, operationErr := transitGatewayApisService.ListTransitGatewayConnections(listTransitGatewayConnectionsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = transitGatewayApisService.ListTransitGatewayConnectionsWithContext(ctx, listTransitGatewayConnectionsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listTransitGatewayConnectionsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["name"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"connections": [{"base_connection_id": "975f58c1-afe7-469a-9727-7f3d720f2d32", "base_network_type": "classic", "cidr": "192.168.0.0/24", "created_at": "2019-01-01T12:00:00.000Z", "id": "1a15dca5-7e33-45e1-b7c5-bc690e569531", "local_bgp_asn": 64490, "local_gateway_ip": "192.168.100.1", "local_tunnel_ip": "192.168.129.2", "mtu": 9000, "name": "Transit_Service_BWTN_SJ_DL", "network_account_id": "NetworkAccountID", "network_id": "crn:v1:bluemix:public:is:us-south:a/123456::vpc:4727d842-f94f-4a2d-824a-9bc9b02c523b", "network_type": "vpc", "prefix_filters": [{"action": "permit", "before": "1a15dcab-7e40-45e1-b7c5-bc690eaa9782", "created_at": "2019-01-01T12:00:00.000Z", "ge": 0, "id": "1a15dcab-7e30-45e1-b7c5-bc690eaa9865", "le": 32, "prefix": "192.168.100.0/24", "updated_at": "2019-01-01T12:00:00.000Z"}], "prefix_filters_default": "permit", "remote_bgp_asn": 65010, "remote_gateway_ip": "10.242.63.12", "remote_tunnel_ip": "192.168.129.1", "request_status": "pending", "status": "attached", "tunnels": [{"base_network_type": "classic", "created_at": "2019-01-01T12:00:00.000Z", "id": "1a15dca5-7e33-45e1-b7c5-bc690e569531", "local_bgp_asn": 11, "local_gateway_ip": "10.242.63.12", "local_tunnel_ip": "192.168.100.20", "mtu": 9000, "name": "gre1", "network_account_id": "NetworkAccountID", "network_id": "crn:v1:bluemix:public:is:us-south:a/123456::vpc:4727d842-f94f-4a2d-824a-9bc9b02c523b", "remote_bgp_asn": 65010, "remote_gateway_ip": "10.242.33.22", "remote_tunnel_ip": "192.168.129.1", "status": "attached", "updated_at": "2019-01-01T12:00:00.000Z", "zone": {"name": "us-south-1"}}], "updated_at": "2019-01-01T12:00:00.000Z", "zone": {"name": "us-south-1"}}], "first": {"href": "https://transit.cloud.ibm.com/v1/transit_gateways/{transit_gateway_id}/connections?limit=50"}, "limit": 50, "next": {"href": "https://transit.cloud.ibm.com/v1/transit_gateways/{transit_gateway_id}/connections?start=MjAyMC0wNS0wOFQxNDoxNzowMy45NzQ5NzNa&limit=50", "start": "MjAyMC0wNS0wOFQxNDoxNzowMy45NzQ5NzNa"}, "total_count": 500}`)
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
				listTransitGatewayConnectionsOptionsModel.Start = core.StringPtr("testString")
				listTransitGatewayConnectionsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listTransitGatewayConnectionsOptionsModel.Name = core.StringPtr("testString")
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
				listTransitGatewayConnectionsOptionsModel.Start = core.StringPtr("testString")
				listTransitGatewayConnectionsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listTransitGatewayConnectionsOptionsModel.Name = core.StringPtr("testString")
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
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
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

				// Construct an instance of the ListTransitGatewayConnectionsOptions model
				listTransitGatewayConnectionsOptionsModel := new(transitgatewayapisv1.ListTransitGatewayConnectionsOptions)
				listTransitGatewayConnectionsOptionsModel.TransitGatewayID = core.StringPtr("testString")
				listTransitGatewayConnectionsOptionsModel.Start = core.StringPtr("testString")
				listTransitGatewayConnectionsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listTransitGatewayConnectionsOptionsModel.Name = core.StringPtr("testString")
				listTransitGatewayConnectionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := transitGatewayApisService.ListTransitGatewayConnections(listTransitGatewayConnectionsOptionsModel)
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
				responseObject := new(transitgatewayapisv1.TransitGatewayConnectionCollection)
				nextObject := new(transitgatewayapisv1.PaginationNextTGWConnection)
				nextObject.Start = core.StringPtr("abc-123")
				responseObject.Next = nextObject
	
				value, err := responseObject.GetNextStart()
				Expect(err).To(BeNil())
				Expect(value).To(Equal(core.StringPtr("abc-123")))
			})
			It(`Invoke GetNextStart without a "Next" property in the response`, func() {
				responseObject := new(transitgatewayapisv1.TransitGatewayConnectionCollection)
	
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
					Expect(req.URL.EscapedPath()).To(Equal(listTransitGatewayConnectionsPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					requestNumber++
					if requestNumber == 1 {
						fmt.Fprintf(res, "%s", `{"next":{"start":"1"},"total_count":2,"limit":1,"connections":[{"base_connection_id":"975f58c1-afe7-469a-9727-7f3d720f2d32","base_network_type":"classic","cidr":"192.168.0.0/24","created_at":"2019-01-01T12:00:00.000Z","id":"1a15dca5-7e33-45e1-b7c5-bc690e569531","local_bgp_asn":64490,"local_gateway_ip":"192.168.100.1","local_tunnel_ip":"192.168.129.2","mtu":9000,"name":"Transit_Service_BWTN_SJ_DL","network_account_id":"NetworkAccountID","network_id":"crn:v1:bluemix:public:is:us-south:a/123456::vpc:4727d842-f94f-4a2d-824a-9bc9b02c523b","network_type":"vpc","prefix_filters":[{"action":"permit","before":"1a15dcab-7e40-45e1-b7c5-bc690eaa9782","created_at":"2019-01-01T12:00:00.000Z","ge":0,"id":"1a15dcab-7e30-45e1-b7c5-bc690eaa9865","le":32,"prefix":"192.168.100.0/24","updated_at":"2019-01-01T12:00:00.000Z"}],"prefix_filters_default":"permit","remote_bgp_asn":65010,"remote_gateway_ip":"10.242.63.12","remote_tunnel_ip":"192.168.129.1","request_status":"pending","status":"attached","tunnels":[{"base_network_type":"classic","created_at":"2019-01-01T12:00:00.000Z","id":"1a15dca5-7e33-45e1-b7c5-bc690e569531","local_bgp_asn":11,"local_gateway_ip":"10.242.63.12","local_tunnel_ip":"192.168.100.20","mtu":9000,"name":"gre1","network_account_id":"NetworkAccountID","network_id":"crn:v1:bluemix:public:is:us-south:a/123456::vpc:4727d842-f94f-4a2d-824a-9bc9b02c523b","remote_bgp_asn":65010,"remote_gateway_ip":"10.242.33.22","remote_tunnel_ip":"192.168.129.1","status":"attached","updated_at":"2019-01-01T12:00:00.000Z","zone":{"name":"us-south-1"}}],"updated_at":"2019-01-01T12:00:00.000Z","zone":{"name":"us-south-1"}}]}`)
					} else if requestNumber == 2 {
						fmt.Fprintf(res, "%s", `{"total_count":2,"limit":1,"connections":[{"base_connection_id":"975f58c1-afe7-469a-9727-7f3d720f2d32","base_network_type":"classic","cidr":"192.168.0.0/24","created_at":"2019-01-01T12:00:00.000Z","id":"1a15dca5-7e33-45e1-b7c5-bc690e569531","local_bgp_asn":64490,"local_gateway_ip":"192.168.100.1","local_tunnel_ip":"192.168.129.2","mtu":9000,"name":"Transit_Service_BWTN_SJ_DL","network_account_id":"NetworkAccountID","network_id":"crn:v1:bluemix:public:is:us-south:a/123456::vpc:4727d842-f94f-4a2d-824a-9bc9b02c523b","network_type":"vpc","prefix_filters":[{"action":"permit","before":"1a15dcab-7e40-45e1-b7c5-bc690eaa9782","created_at":"2019-01-01T12:00:00.000Z","ge":0,"id":"1a15dcab-7e30-45e1-b7c5-bc690eaa9865","le":32,"prefix":"192.168.100.0/24","updated_at":"2019-01-01T12:00:00.000Z"}],"prefix_filters_default":"permit","remote_bgp_asn":65010,"remote_gateway_ip":"10.242.63.12","remote_tunnel_ip":"192.168.129.1","request_status":"pending","status":"attached","tunnels":[{"base_network_type":"classic","created_at":"2019-01-01T12:00:00.000Z","id":"1a15dca5-7e33-45e1-b7c5-bc690e569531","local_bgp_asn":11,"local_gateway_ip":"10.242.63.12","local_tunnel_ip":"192.168.100.20","mtu":9000,"name":"gre1","network_account_id":"NetworkAccountID","network_id":"crn:v1:bluemix:public:is:us-south:a/123456::vpc:4727d842-f94f-4a2d-824a-9bc9b02c523b","remote_bgp_asn":65010,"remote_gateway_ip":"10.242.33.22","remote_tunnel_ip":"192.168.129.1","status":"attached","updated_at":"2019-01-01T12:00:00.000Z","zone":{"name":"us-south-1"}}],"updated_at":"2019-01-01T12:00:00.000Z","zone":{"name":"us-south-1"}}]}`)
					} else {
						res.WriteHeader(400)
					}
				}))
			})
			It(`Use TransitGatewayConnectionsPager.GetNext successfully`, func() {
				transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(transitGatewayApisService).ToNot(BeNil())

				listTransitGatewayConnectionsOptionsModel := &transitgatewayapisv1.ListTransitGatewayConnectionsOptions{
					TransitGatewayID: core.StringPtr("testString"),
					Limit: core.Int64Ptr(int64(10)),
					Name: core.StringPtr("testString"),
				}

				pager, err := transitGatewayApisService.NewTransitGatewayConnectionsPager(listTransitGatewayConnectionsOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				var allResults []transitgatewayapisv1.TransitGatewayConnectionCust
				for pager.HasNext() {
					nextPage, err := pager.GetNext()
					Expect(err).To(BeNil())
					Expect(nextPage).ToNot(BeNil())
					allResults = append(allResults, nextPage...)
				}
				Expect(len(allResults)).To(Equal(2))
			})
			It(`Use TransitGatewayConnectionsPager.GetAll successfully`, func() {
				transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(transitGatewayApisService).ToNot(BeNil())

				listTransitGatewayConnectionsOptionsModel := &transitgatewayapisv1.ListTransitGatewayConnectionsOptions{
					TransitGatewayID: core.StringPtr("testString"),
					Limit: core.Int64Ptr(int64(10)),
					Name: core.StringPtr("testString"),
				}

				pager, err := transitGatewayApisService.NewTransitGatewayConnectionsPager(listTransitGatewayConnectionsOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				allResults, err := pager.GetAll()
				Expect(err).To(BeNil())
				Expect(allResults).ToNot(BeNil())
				Expect(len(allResults)).To(Equal(2))
			})
		})
	})
	Describe(`CreateTransitGatewayConnection(createTransitGatewayConnectionOptions *CreateTransitGatewayConnectionOptions) - Operation response error`, func() {
		version := "testString"
		createTransitGatewayConnectionPath := "/transit_gateways/testString/connections"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createTransitGatewayConnectionPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
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

				// Construct an instance of the TransitGatewayConnectionPrefixFilter model
				transitGatewayConnectionPrefixFilterModel := new(transitgatewayapisv1.TransitGatewayConnectionPrefixFilter)
				transitGatewayConnectionPrefixFilterModel.Action = core.StringPtr("permit")
				transitGatewayConnectionPrefixFilterModel.Ge = core.Int64Ptr(int64(0))
				transitGatewayConnectionPrefixFilterModel.Le = core.Int64Ptr(int64(32))
				transitGatewayConnectionPrefixFilterModel.Prefix = core.StringPtr("192.168.100.0/24")

				// Construct an instance of the ZoneIdentityByName model
				zoneIdentityModel := new(transitgatewayapisv1.ZoneIdentityByName)
				zoneIdentityModel.Name = core.StringPtr("us-south-1")

				// Construct an instance of the TransitGatewayTunnelTemplate model
				transitGatewayTunnelTemplateModel := new(transitgatewayapisv1.TransitGatewayTunnelTemplate)
				transitGatewayTunnelTemplateModel.LocalGatewayIp = core.StringPtr("10.242.63.12")
				transitGatewayTunnelTemplateModel.LocalTunnelIp = core.StringPtr("192.168.100.20")
				transitGatewayTunnelTemplateModel.Name = core.StringPtr("gre1")
				transitGatewayTunnelTemplateModel.RemoteBgpAsn = core.Int64Ptr(int64(65010))
				transitGatewayTunnelTemplateModel.RemoteGatewayIp = core.StringPtr("10.242.33.22")
				transitGatewayTunnelTemplateModel.RemoteTunnelIp = core.StringPtr("192.168.129.1")
				transitGatewayTunnelTemplateModel.Zone = zoneIdentityModel

				// Construct an instance of the CreateTransitGatewayConnectionOptions model
				createTransitGatewayConnectionOptionsModel := new(transitgatewayapisv1.CreateTransitGatewayConnectionOptions)
				createTransitGatewayConnectionOptionsModel.TransitGatewayID = core.StringPtr("testString")
				createTransitGatewayConnectionOptionsModel.NetworkType = core.StringPtr("vpc")
				createTransitGatewayConnectionOptionsModel.BaseConnectionID = core.StringPtr("975f58c1-afe7-469a-9727-7f3d720f2d32")
				createTransitGatewayConnectionOptionsModel.BaseNetworkType = core.StringPtr("classic")
				createTransitGatewayConnectionOptionsModel.Cidr = core.StringPtr("192.168.0.0/24")
				createTransitGatewayConnectionOptionsModel.LocalGatewayIp = core.StringPtr("192.168.100.1")
				createTransitGatewayConnectionOptionsModel.LocalTunnelIp = core.StringPtr("192.168.129.2")
				createTransitGatewayConnectionOptionsModel.Name = core.StringPtr("Transit_Service_BWTN_SJ_DL")
				createTransitGatewayConnectionOptionsModel.NetworkAccountID = core.StringPtr("testString")
				createTransitGatewayConnectionOptionsModel.NetworkID = core.StringPtr("crn:v1:bluemix:public:is:us-south:a/123456::vpc:4727d842-f94f-4a2d-824a-9bc9b02c523b")
				createTransitGatewayConnectionOptionsModel.PrefixFilters = []transitgatewayapisv1.TransitGatewayConnectionPrefixFilter{*transitGatewayConnectionPrefixFilterModel}
				createTransitGatewayConnectionOptionsModel.PrefixFiltersDefault = core.StringPtr("permit")
				createTransitGatewayConnectionOptionsModel.RemoteBgpAsn = core.Int64Ptr(int64(65010))
				createTransitGatewayConnectionOptionsModel.RemoteGatewayIp = core.StringPtr("10.242.63.12")
				createTransitGatewayConnectionOptionsModel.RemoteTunnelIp = core.StringPtr("192.168.129.1")
				createTransitGatewayConnectionOptionsModel.Tunnels = []transitgatewayapisv1.TransitGatewayTunnelTemplate{*transitGatewayTunnelTemplateModel}
				createTransitGatewayConnectionOptionsModel.Zone = zoneIdentityModel
				createTransitGatewayConnectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := transitGatewayApisService.CreateTransitGatewayConnection(createTransitGatewayConnectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				transitGatewayApisService.EnableRetries(0, 0)
				result, response, operationErr = transitGatewayApisService.CreateTransitGatewayConnection(createTransitGatewayConnectionOptionsModel)
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
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createTransitGatewayConnectionPath))
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
					fmt.Fprintf(res, "%s", `{"base_connection_id": "975f58c1-afe7-469a-9727-7f3d720f2d32", "base_network_type": "classic", "cidr": "192.168.0.0/24", "created_at": "2019-01-01T12:00:00.000Z", "id": "1a15dca5-7e33-45e1-b7c5-bc690e569531", "local_bgp_asn": 64490, "local_gateway_ip": "192.168.100.1", "local_tunnel_ip": "192.168.129.2", "mtu": 9000, "name": "Transit_Service_BWTN_SJ_DL", "network_account_id": "NetworkAccountID", "network_id": "crn:v1:bluemix:public:is:us-south:a/123456::vpc:4727d842-f94f-4a2d-824a-9bc9b02c523b", "network_type": "vpc", "prefix_filters": [{"action": "permit", "before": "1a15dcab-7e40-45e1-b7c5-bc690eaa9782", "created_at": "2019-01-01T12:00:00.000Z", "ge": 0, "id": "1a15dcab-7e30-45e1-b7c5-bc690eaa9865", "le": 32, "prefix": "192.168.100.0/24", "updated_at": "2019-01-01T12:00:00.000Z"}], "prefix_filters_default": "permit", "remote_bgp_asn": 65010, "remote_gateway_ip": "10.242.63.12", "remote_tunnel_ip": "192.168.129.1", "request_status": "pending", "status": "attached", "tunnels": [{"base_network_type": "classic", "created_at": "2019-01-01T12:00:00.000Z", "id": "1a15dca5-7e33-45e1-b7c5-bc690e569531", "local_bgp_asn": 11, "local_gateway_ip": "10.242.63.12", "local_tunnel_ip": "192.168.100.20", "mtu": 9000, "name": "gre1", "network_account_id": "NetworkAccountID", "network_id": "crn:v1:bluemix:public:is:us-south:a/123456::vpc:4727d842-f94f-4a2d-824a-9bc9b02c523b", "remote_bgp_asn": 65010, "remote_gateway_ip": "10.242.33.22", "remote_tunnel_ip": "192.168.129.1", "status": "attached", "updated_at": "2019-01-01T12:00:00.000Z", "zone": {"name": "us-south-1"}}], "updated_at": "2019-01-01T12:00:00.000Z", "zone": {"name": "us-south-1"}}`)
				}))
			})
			It(`Invoke CreateTransitGatewayConnection successfully with retries`, func() {
				transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(transitGatewayApisService).ToNot(BeNil())
				transitGatewayApisService.EnableRetries(0, 0)

				// Construct an instance of the TransitGatewayConnectionPrefixFilter model
				transitGatewayConnectionPrefixFilterModel := new(transitgatewayapisv1.TransitGatewayConnectionPrefixFilter)
				transitGatewayConnectionPrefixFilterModel.Action = core.StringPtr("permit")
				transitGatewayConnectionPrefixFilterModel.Ge = core.Int64Ptr(int64(0))
				transitGatewayConnectionPrefixFilterModel.Le = core.Int64Ptr(int64(32))
				transitGatewayConnectionPrefixFilterModel.Prefix = core.StringPtr("192.168.100.0/24")

				// Construct an instance of the ZoneIdentityByName model
				zoneIdentityModel := new(transitgatewayapisv1.ZoneIdentityByName)
				zoneIdentityModel.Name = core.StringPtr("us-south-1")

				// Construct an instance of the TransitGatewayTunnelTemplate model
				transitGatewayTunnelTemplateModel := new(transitgatewayapisv1.TransitGatewayTunnelTemplate)
				transitGatewayTunnelTemplateModel.LocalGatewayIp = core.StringPtr("10.242.63.12")
				transitGatewayTunnelTemplateModel.LocalTunnelIp = core.StringPtr("192.168.100.20")
				transitGatewayTunnelTemplateModel.Name = core.StringPtr("gre1")
				transitGatewayTunnelTemplateModel.RemoteBgpAsn = core.Int64Ptr(int64(65010))
				transitGatewayTunnelTemplateModel.RemoteGatewayIp = core.StringPtr("10.242.33.22")
				transitGatewayTunnelTemplateModel.RemoteTunnelIp = core.StringPtr("192.168.129.1")
				transitGatewayTunnelTemplateModel.Zone = zoneIdentityModel

				// Construct an instance of the CreateTransitGatewayConnectionOptions model
				createTransitGatewayConnectionOptionsModel := new(transitgatewayapisv1.CreateTransitGatewayConnectionOptions)
				createTransitGatewayConnectionOptionsModel.TransitGatewayID = core.StringPtr("testString")
				createTransitGatewayConnectionOptionsModel.NetworkType = core.StringPtr("vpc")
				createTransitGatewayConnectionOptionsModel.BaseConnectionID = core.StringPtr("975f58c1-afe7-469a-9727-7f3d720f2d32")
				createTransitGatewayConnectionOptionsModel.BaseNetworkType = core.StringPtr("classic")
				createTransitGatewayConnectionOptionsModel.Cidr = core.StringPtr("192.168.0.0/24")
				createTransitGatewayConnectionOptionsModel.LocalGatewayIp = core.StringPtr("192.168.100.1")
				createTransitGatewayConnectionOptionsModel.LocalTunnelIp = core.StringPtr("192.168.129.2")
				createTransitGatewayConnectionOptionsModel.Name = core.StringPtr("Transit_Service_BWTN_SJ_DL")
				createTransitGatewayConnectionOptionsModel.NetworkAccountID = core.StringPtr("testString")
				createTransitGatewayConnectionOptionsModel.NetworkID = core.StringPtr("crn:v1:bluemix:public:is:us-south:a/123456::vpc:4727d842-f94f-4a2d-824a-9bc9b02c523b")
				createTransitGatewayConnectionOptionsModel.PrefixFilters = []transitgatewayapisv1.TransitGatewayConnectionPrefixFilter{*transitGatewayConnectionPrefixFilterModel}
				createTransitGatewayConnectionOptionsModel.PrefixFiltersDefault = core.StringPtr("permit")
				createTransitGatewayConnectionOptionsModel.RemoteBgpAsn = core.Int64Ptr(int64(65010))
				createTransitGatewayConnectionOptionsModel.RemoteGatewayIp = core.StringPtr("10.242.63.12")
				createTransitGatewayConnectionOptionsModel.RemoteTunnelIp = core.StringPtr("192.168.129.1")
				createTransitGatewayConnectionOptionsModel.Tunnels = []transitgatewayapisv1.TransitGatewayTunnelTemplate{*transitGatewayTunnelTemplateModel}
				createTransitGatewayConnectionOptionsModel.Zone = zoneIdentityModel
				createTransitGatewayConnectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := transitGatewayApisService.CreateTransitGatewayConnectionWithContext(ctx, createTransitGatewayConnectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				transitGatewayApisService.DisableRetries()
				result, response, operationErr := transitGatewayApisService.CreateTransitGatewayConnection(createTransitGatewayConnectionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = transitGatewayApisService.CreateTransitGatewayConnectionWithContext(ctx, createTransitGatewayConnectionOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createTransitGatewayConnectionPath))
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
					fmt.Fprintf(res, "%s", `{"base_connection_id": "975f58c1-afe7-469a-9727-7f3d720f2d32", "base_network_type": "classic", "cidr": "192.168.0.0/24", "created_at": "2019-01-01T12:00:00.000Z", "id": "1a15dca5-7e33-45e1-b7c5-bc690e569531", "local_bgp_asn": 64490, "local_gateway_ip": "192.168.100.1", "local_tunnel_ip": "192.168.129.2", "mtu": 9000, "name": "Transit_Service_BWTN_SJ_DL", "network_account_id": "NetworkAccountID", "network_id": "crn:v1:bluemix:public:is:us-south:a/123456::vpc:4727d842-f94f-4a2d-824a-9bc9b02c523b", "network_type": "vpc", "prefix_filters": [{"action": "permit", "before": "1a15dcab-7e40-45e1-b7c5-bc690eaa9782", "created_at": "2019-01-01T12:00:00.000Z", "ge": 0, "id": "1a15dcab-7e30-45e1-b7c5-bc690eaa9865", "le": 32, "prefix": "192.168.100.0/24", "updated_at": "2019-01-01T12:00:00.000Z"}], "prefix_filters_default": "permit", "remote_bgp_asn": 65010, "remote_gateway_ip": "10.242.63.12", "remote_tunnel_ip": "192.168.129.1", "request_status": "pending", "status": "attached", "tunnels": [{"base_network_type": "classic", "created_at": "2019-01-01T12:00:00.000Z", "id": "1a15dca5-7e33-45e1-b7c5-bc690e569531", "local_bgp_asn": 11, "local_gateway_ip": "10.242.63.12", "local_tunnel_ip": "192.168.100.20", "mtu": 9000, "name": "gre1", "network_account_id": "NetworkAccountID", "network_id": "crn:v1:bluemix:public:is:us-south:a/123456::vpc:4727d842-f94f-4a2d-824a-9bc9b02c523b", "remote_bgp_asn": 65010, "remote_gateway_ip": "10.242.33.22", "remote_tunnel_ip": "192.168.129.1", "status": "attached", "updated_at": "2019-01-01T12:00:00.000Z", "zone": {"name": "us-south-1"}}], "updated_at": "2019-01-01T12:00:00.000Z", "zone": {"name": "us-south-1"}}`)
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

				// Construct an instance of the TransitGatewayConnectionPrefixFilter model
				transitGatewayConnectionPrefixFilterModel := new(transitgatewayapisv1.TransitGatewayConnectionPrefixFilter)
				transitGatewayConnectionPrefixFilterModel.Action = core.StringPtr("permit")
				transitGatewayConnectionPrefixFilterModel.Ge = core.Int64Ptr(int64(0))
				transitGatewayConnectionPrefixFilterModel.Le = core.Int64Ptr(int64(32))
				transitGatewayConnectionPrefixFilterModel.Prefix = core.StringPtr("192.168.100.0/24")

				// Construct an instance of the ZoneIdentityByName model
				zoneIdentityModel := new(transitgatewayapisv1.ZoneIdentityByName)
				zoneIdentityModel.Name = core.StringPtr("us-south-1")

				// Construct an instance of the TransitGatewayTunnelTemplate model
				transitGatewayTunnelTemplateModel := new(transitgatewayapisv1.TransitGatewayTunnelTemplate)
				transitGatewayTunnelTemplateModel.LocalGatewayIp = core.StringPtr("10.242.63.12")
				transitGatewayTunnelTemplateModel.LocalTunnelIp = core.StringPtr("192.168.100.20")
				transitGatewayTunnelTemplateModel.Name = core.StringPtr("gre1")
				transitGatewayTunnelTemplateModel.RemoteBgpAsn = core.Int64Ptr(int64(65010))
				transitGatewayTunnelTemplateModel.RemoteGatewayIp = core.StringPtr("10.242.33.22")
				transitGatewayTunnelTemplateModel.RemoteTunnelIp = core.StringPtr("192.168.129.1")
				transitGatewayTunnelTemplateModel.Zone = zoneIdentityModel

				// Construct an instance of the CreateTransitGatewayConnectionOptions model
				createTransitGatewayConnectionOptionsModel := new(transitgatewayapisv1.CreateTransitGatewayConnectionOptions)
				createTransitGatewayConnectionOptionsModel.TransitGatewayID = core.StringPtr("testString")
				createTransitGatewayConnectionOptionsModel.NetworkType = core.StringPtr("vpc")
				createTransitGatewayConnectionOptionsModel.BaseConnectionID = core.StringPtr("975f58c1-afe7-469a-9727-7f3d720f2d32")
				createTransitGatewayConnectionOptionsModel.BaseNetworkType = core.StringPtr("classic")
				createTransitGatewayConnectionOptionsModel.Cidr = core.StringPtr("192.168.0.0/24")
				createTransitGatewayConnectionOptionsModel.LocalGatewayIp = core.StringPtr("192.168.100.1")
				createTransitGatewayConnectionOptionsModel.LocalTunnelIp = core.StringPtr("192.168.129.2")
				createTransitGatewayConnectionOptionsModel.Name = core.StringPtr("Transit_Service_BWTN_SJ_DL")
				createTransitGatewayConnectionOptionsModel.NetworkAccountID = core.StringPtr("testString")
				createTransitGatewayConnectionOptionsModel.NetworkID = core.StringPtr("crn:v1:bluemix:public:is:us-south:a/123456::vpc:4727d842-f94f-4a2d-824a-9bc9b02c523b")
				createTransitGatewayConnectionOptionsModel.PrefixFilters = []transitgatewayapisv1.TransitGatewayConnectionPrefixFilter{*transitGatewayConnectionPrefixFilterModel}
				createTransitGatewayConnectionOptionsModel.PrefixFiltersDefault = core.StringPtr("permit")
				createTransitGatewayConnectionOptionsModel.RemoteBgpAsn = core.Int64Ptr(int64(65010))
				createTransitGatewayConnectionOptionsModel.RemoteGatewayIp = core.StringPtr("10.242.63.12")
				createTransitGatewayConnectionOptionsModel.RemoteTunnelIp = core.StringPtr("192.168.129.1")
				createTransitGatewayConnectionOptionsModel.Tunnels = []transitgatewayapisv1.TransitGatewayTunnelTemplate{*transitGatewayTunnelTemplateModel}
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

				// Construct an instance of the TransitGatewayConnectionPrefixFilter model
				transitGatewayConnectionPrefixFilterModel := new(transitgatewayapisv1.TransitGatewayConnectionPrefixFilter)
				transitGatewayConnectionPrefixFilterModel.Action = core.StringPtr("permit")
				transitGatewayConnectionPrefixFilterModel.Ge = core.Int64Ptr(int64(0))
				transitGatewayConnectionPrefixFilterModel.Le = core.Int64Ptr(int64(32))
				transitGatewayConnectionPrefixFilterModel.Prefix = core.StringPtr("192.168.100.0/24")

				// Construct an instance of the ZoneIdentityByName model
				zoneIdentityModel := new(transitgatewayapisv1.ZoneIdentityByName)
				zoneIdentityModel.Name = core.StringPtr("us-south-1")

				// Construct an instance of the TransitGatewayTunnelTemplate model
				transitGatewayTunnelTemplateModel := new(transitgatewayapisv1.TransitGatewayTunnelTemplate)
				transitGatewayTunnelTemplateModel.LocalGatewayIp = core.StringPtr("10.242.63.12")
				transitGatewayTunnelTemplateModel.LocalTunnelIp = core.StringPtr("192.168.100.20")
				transitGatewayTunnelTemplateModel.Name = core.StringPtr("gre1")
				transitGatewayTunnelTemplateModel.RemoteBgpAsn = core.Int64Ptr(int64(65010))
				transitGatewayTunnelTemplateModel.RemoteGatewayIp = core.StringPtr("10.242.33.22")
				transitGatewayTunnelTemplateModel.RemoteTunnelIp = core.StringPtr("192.168.129.1")
				transitGatewayTunnelTemplateModel.Zone = zoneIdentityModel

				// Construct an instance of the CreateTransitGatewayConnectionOptions model
				createTransitGatewayConnectionOptionsModel := new(transitgatewayapisv1.CreateTransitGatewayConnectionOptions)
				createTransitGatewayConnectionOptionsModel.TransitGatewayID = core.StringPtr("testString")
				createTransitGatewayConnectionOptionsModel.NetworkType = core.StringPtr("vpc")
				createTransitGatewayConnectionOptionsModel.BaseConnectionID = core.StringPtr("975f58c1-afe7-469a-9727-7f3d720f2d32")
				createTransitGatewayConnectionOptionsModel.BaseNetworkType = core.StringPtr("classic")
				createTransitGatewayConnectionOptionsModel.Cidr = core.StringPtr("192.168.0.0/24")
				createTransitGatewayConnectionOptionsModel.LocalGatewayIp = core.StringPtr("192.168.100.1")
				createTransitGatewayConnectionOptionsModel.LocalTunnelIp = core.StringPtr("192.168.129.2")
				createTransitGatewayConnectionOptionsModel.Name = core.StringPtr("Transit_Service_BWTN_SJ_DL")
				createTransitGatewayConnectionOptionsModel.NetworkAccountID = core.StringPtr("testString")
				createTransitGatewayConnectionOptionsModel.NetworkID = core.StringPtr("crn:v1:bluemix:public:is:us-south:a/123456::vpc:4727d842-f94f-4a2d-824a-9bc9b02c523b")
				createTransitGatewayConnectionOptionsModel.PrefixFilters = []transitgatewayapisv1.TransitGatewayConnectionPrefixFilter{*transitGatewayConnectionPrefixFilterModel}
				createTransitGatewayConnectionOptionsModel.PrefixFiltersDefault = core.StringPtr("permit")
				createTransitGatewayConnectionOptionsModel.RemoteBgpAsn = core.Int64Ptr(int64(65010))
				createTransitGatewayConnectionOptionsModel.RemoteGatewayIp = core.StringPtr("10.242.63.12")
				createTransitGatewayConnectionOptionsModel.RemoteTunnelIp = core.StringPtr("192.168.129.1")
				createTransitGatewayConnectionOptionsModel.Tunnels = []transitgatewayapisv1.TransitGatewayTunnelTemplate{*transitGatewayTunnelTemplateModel}
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
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(201)
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

				// Construct an instance of the TransitGatewayConnectionPrefixFilter model
				transitGatewayConnectionPrefixFilterModel := new(transitgatewayapisv1.TransitGatewayConnectionPrefixFilter)
				transitGatewayConnectionPrefixFilterModel.Action = core.StringPtr("permit")
				transitGatewayConnectionPrefixFilterModel.Ge = core.Int64Ptr(int64(0))
				transitGatewayConnectionPrefixFilterModel.Le = core.Int64Ptr(int64(32))
				transitGatewayConnectionPrefixFilterModel.Prefix = core.StringPtr("192.168.100.0/24")

				// Construct an instance of the ZoneIdentityByName model
				zoneIdentityModel := new(transitgatewayapisv1.ZoneIdentityByName)
				zoneIdentityModel.Name = core.StringPtr("us-south-1")

				// Construct an instance of the TransitGatewayTunnelTemplate model
				transitGatewayTunnelTemplateModel := new(transitgatewayapisv1.TransitGatewayTunnelTemplate)
				transitGatewayTunnelTemplateModel.LocalGatewayIp = core.StringPtr("10.242.63.12")
				transitGatewayTunnelTemplateModel.LocalTunnelIp = core.StringPtr("192.168.100.20")
				transitGatewayTunnelTemplateModel.Name = core.StringPtr("gre1")
				transitGatewayTunnelTemplateModel.RemoteBgpAsn = core.Int64Ptr(int64(65010))
				transitGatewayTunnelTemplateModel.RemoteGatewayIp = core.StringPtr("10.242.33.22")
				transitGatewayTunnelTemplateModel.RemoteTunnelIp = core.StringPtr("192.168.129.1")
				transitGatewayTunnelTemplateModel.Zone = zoneIdentityModel

				// Construct an instance of the CreateTransitGatewayConnectionOptions model
				createTransitGatewayConnectionOptionsModel := new(transitgatewayapisv1.CreateTransitGatewayConnectionOptions)
				createTransitGatewayConnectionOptionsModel.TransitGatewayID = core.StringPtr("testString")
				createTransitGatewayConnectionOptionsModel.NetworkType = core.StringPtr("vpc")
				createTransitGatewayConnectionOptionsModel.BaseConnectionID = core.StringPtr("975f58c1-afe7-469a-9727-7f3d720f2d32")
				createTransitGatewayConnectionOptionsModel.BaseNetworkType = core.StringPtr("classic")
				createTransitGatewayConnectionOptionsModel.Cidr = core.StringPtr("192.168.0.0/24")
				createTransitGatewayConnectionOptionsModel.LocalGatewayIp = core.StringPtr("192.168.100.1")
				createTransitGatewayConnectionOptionsModel.LocalTunnelIp = core.StringPtr("192.168.129.2")
				createTransitGatewayConnectionOptionsModel.Name = core.StringPtr("Transit_Service_BWTN_SJ_DL")
				createTransitGatewayConnectionOptionsModel.NetworkAccountID = core.StringPtr("testString")
				createTransitGatewayConnectionOptionsModel.NetworkID = core.StringPtr("crn:v1:bluemix:public:is:us-south:a/123456::vpc:4727d842-f94f-4a2d-824a-9bc9b02c523b")
				createTransitGatewayConnectionOptionsModel.PrefixFilters = []transitgatewayapisv1.TransitGatewayConnectionPrefixFilter{*transitGatewayConnectionPrefixFilterModel}
				createTransitGatewayConnectionOptionsModel.PrefixFiltersDefault = core.StringPtr("permit")
				createTransitGatewayConnectionOptionsModel.RemoteBgpAsn = core.Int64Ptr(int64(65010))
				createTransitGatewayConnectionOptionsModel.RemoteGatewayIp = core.StringPtr("10.242.63.12")
				createTransitGatewayConnectionOptionsModel.RemoteTunnelIp = core.StringPtr("192.168.129.1")
				createTransitGatewayConnectionOptionsModel.Tunnels = []transitgatewayapisv1.TransitGatewayTunnelTemplate{*transitGatewayTunnelTemplateModel}
				createTransitGatewayConnectionOptionsModel.Zone = zoneIdentityModel
				createTransitGatewayConnectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := transitGatewayApisService.CreateTransitGatewayConnection(createTransitGatewayConnectionOptionsModel)
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
	Describe(`DeleteTransitGatewayConnection(deleteTransitGatewayConnectionOptions *DeleteTransitGatewayConnectionOptions)`, func() {
		version := "testString"
		deleteTransitGatewayConnectionPath := "/transit_gateways/testString/connections/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteTransitGatewayConnectionPath))
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
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getTransitGatewayConnectionPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
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

				// Enable retries and test again
				transitGatewayApisService.EnableRetries(0, 0)
				result, response, operationErr = transitGatewayApisService.GetTransitGatewayConnection(getTransitGatewayConnectionOptionsModel)
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
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getTransitGatewayConnectionPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"base_connection_id": "975f58c1-afe7-469a-9727-7f3d720f2d32", "base_network_type": "classic", "cidr": "192.168.0.0/24", "created_at": "2019-01-01T12:00:00.000Z", "id": "1a15dca5-7e33-45e1-b7c5-bc690e569531", "local_bgp_asn": 64490, "local_gateway_ip": "192.168.100.1", "local_tunnel_ip": "192.168.129.2", "mtu": 9000, "name": "Transit_Service_BWTN_SJ_DL", "network_account_id": "NetworkAccountID", "network_id": "crn:v1:bluemix:public:is:us-south:a/123456::vpc:4727d842-f94f-4a2d-824a-9bc9b02c523b", "network_type": "vpc", "prefix_filters": [{"action": "permit", "before": "1a15dcab-7e40-45e1-b7c5-bc690eaa9782", "created_at": "2019-01-01T12:00:00.000Z", "ge": 0, "id": "1a15dcab-7e30-45e1-b7c5-bc690eaa9865", "le": 32, "prefix": "192.168.100.0/24", "updated_at": "2019-01-01T12:00:00.000Z"}], "prefix_filters_default": "permit", "remote_bgp_asn": 65010, "remote_gateway_ip": "10.242.63.12", "remote_tunnel_ip": "192.168.129.1", "request_status": "pending", "status": "attached", "tunnels": [{"base_network_type": "classic", "created_at": "2019-01-01T12:00:00.000Z", "id": "1a15dca5-7e33-45e1-b7c5-bc690e569531", "local_bgp_asn": 11, "local_gateway_ip": "10.242.63.12", "local_tunnel_ip": "192.168.100.20", "mtu": 9000, "name": "gre1", "network_account_id": "NetworkAccountID", "network_id": "crn:v1:bluemix:public:is:us-south:a/123456::vpc:4727d842-f94f-4a2d-824a-9bc9b02c523b", "remote_bgp_asn": 65010, "remote_gateway_ip": "10.242.33.22", "remote_tunnel_ip": "192.168.129.1", "status": "attached", "updated_at": "2019-01-01T12:00:00.000Z", "zone": {"name": "us-south-1"}}], "updated_at": "2019-01-01T12:00:00.000Z", "zone": {"name": "us-south-1"}}`)
				}))
			})
			It(`Invoke GetTransitGatewayConnection successfully with retries`, func() {
				transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(transitGatewayApisService).ToNot(BeNil())
				transitGatewayApisService.EnableRetries(0, 0)

				// Construct an instance of the GetTransitGatewayConnectionOptions model
				getTransitGatewayConnectionOptionsModel := new(transitgatewayapisv1.GetTransitGatewayConnectionOptions)
				getTransitGatewayConnectionOptionsModel.TransitGatewayID = core.StringPtr("testString")
				getTransitGatewayConnectionOptionsModel.ID = core.StringPtr("testString")
				getTransitGatewayConnectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := transitGatewayApisService.GetTransitGatewayConnectionWithContext(ctx, getTransitGatewayConnectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				transitGatewayApisService.DisableRetries()
				result, response, operationErr := transitGatewayApisService.GetTransitGatewayConnection(getTransitGatewayConnectionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = transitGatewayApisService.GetTransitGatewayConnectionWithContext(ctx, getTransitGatewayConnectionOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getTransitGatewayConnectionPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"base_connection_id": "975f58c1-afe7-469a-9727-7f3d720f2d32", "base_network_type": "classic", "cidr": "192.168.0.0/24", "created_at": "2019-01-01T12:00:00.000Z", "id": "1a15dca5-7e33-45e1-b7c5-bc690e569531", "local_bgp_asn": 64490, "local_gateway_ip": "192.168.100.1", "local_tunnel_ip": "192.168.129.2", "mtu": 9000, "name": "Transit_Service_BWTN_SJ_DL", "network_account_id": "NetworkAccountID", "network_id": "crn:v1:bluemix:public:is:us-south:a/123456::vpc:4727d842-f94f-4a2d-824a-9bc9b02c523b", "network_type": "vpc", "prefix_filters": [{"action": "permit", "before": "1a15dcab-7e40-45e1-b7c5-bc690eaa9782", "created_at": "2019-01-01T12:00:00.000Z", "ge": 0, "id": "1a15dcab-7e30-45e1-b7c5-bc690eaa9865", "le": 32, "prefix": "192.168.100.0/24", "updated_at": "2019-01-01T12:00:00.000Z"}], "prefix_filters_default": "permit", "remote_bgp_asn": 65010, "remote_gateway_ip": "10.242.63.12", "remote_tunnel_ip": "192.168.129.1", "request_status": "pending", "status": "attached", "tunnels": [{"base_network_type": "classic", "created_at": "2019-01-01T12:00:00.000Z", "id": "1a15dca5-7e33-45e1-b7c5-bc690e569531", "local_bgp_asn": 11, "local_gateway_ip": "10.242.63.12", "local_tunnel_ip": "192.168.100.20", "mtu": 9000, "name": "gre1", "network_account_id": "NetworkAccountID", "network_id": "crn:v1:bluemix:public:is:us-south:a/123456::vpc:4727d842-f94f-4a2d-824a-9bc9b02c523b", "remote_bgp_asn": 65010, "remote_gateway_ip": "10.242.33.22", "remote_tunnel_ip": "192.168.129.1", "status": "attached", "updated_at": "2019-01-01T12:00:00.000Z", "zone": {"name": "us-south-1"}}], "updated_at": "2019-01-01T12:00:00.000Z", "zone": {"name": "us-south-1"}}`)
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
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
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

				// Construct an instance of the GetTransitGatewayConnectionOptions model
				getTransitGatewayConnectionOptionsModel := new(transitgatewayapisv1.GetTransitGatewayConnectionOptions)
				getTransitGatewayConnectionOptionsModel.TransitGatewayID = core.StringPtr("testString")
				getTransitGatewayConnectionOptionsModel.ID = core.StringPtr("testString")
				getTransitGatewayConnectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := transitGatewayApisService.GetTransitGatewayConnection(getTransitGatewayConnectionOptionsModel)
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
	Describe(`UpdateTransitGatewayConnection(updateTransitGatewayConnectionOptions *UpdateTransitGatewayConnectionOptions) - Operation response error`, func() {
		version := "testString"
		updateTransitGatewayConnectionPath := "/transit_gateways/testString/connections/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateTransitGatewayConnectionPath))
					Expect(req.Method).To(Equal("PATCH"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
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
				updateTransitGatewayConnectionOptionsModel.PrefixFiltersDefault = core.StringPtr("permit")
				updateTransitGatewayConnectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := transitGatewayApisService.UpdateTransitGatewayConnection(updateTransitGatewayConnectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				transitGatewayApisService.EnableRetries(0, 0)
				result, response, operationErr = transitGatewayApisService.UpdateTransitGatewayConnection(updateTransitGatewayConnectionOptionsModel)
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
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateTransitGatewayConnectionPath))
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
					fmt.Fprintf(res, "%s", `{"base_connection_id": "975f58c1-afe7-469a-9727-7f3d720f2d32", "base_network_type": "classic", "cidr": "192.168.0.0/24", "created_at": "2019-01-01T12:00:00.000Z", "id": "1a15dca5-7e33-45e1-b7c5-bc690e569531", "local_bgp_asn": 64490, "local_gateway_ip": "192.168.100.1", "local_tunnel_ip": "192.168.129.2", "mtu": 9000, "name": "Transit_Service_BWTN_SJ_DL", "network_account_id": "NetworkAccountID", "network_id": "crn:v1:bluemix:public:is:us-south:a/123456::vpc:4727d842-f94f-4a2d-824a-9bc9b02c523b", "network_type": "vpc", "prefix_filters": [{"action": "permit", "before": "1a15dcab-7e40-45e1-b7c5-bc690eaa9782", "created_at": "2019-01-01T12:00:00.000Z", "ge": 0, "id": "1a15dcab-7e30-45e1-b7c5-bc690eaa9865", "le": 32, "prefix": "192.168.100.0/24", "updated_at": "2019-01-01T12:00:00.000Z"}], "prefix_filters_default": "permit", "remote_bgp_asn": 65010, "remote_gateway_ip": "10.242.63.12", "remote_tunnel_ip": "192.168.129.1", "request_status": "pending", "status": "attached", "tunnels": [{"base_network_type": "classic", "created_at": "2019-01-01T12:00:00.000Z", "id": "1a15dca5-7e33-45e1-b7c5-bc690e569531", "local_bgp_asn": 11, "local_gateway_ip": "10.242.63.12", "local_tunnel_ip": "192.168.100.20", "mtu": 9000, "name": "gre1", "network_account_id": "NetworkAccountID", "network_id": "crn:v1:bluemix:public:is:us-south:a/123456::vpc:4727d842-f94f-4a2d-824a-9bc9b02c523b", "remote_bgp_asn": 65010, "remote_gateway_ip": "10.242.33.22", "remote_tunnel_ip": "192.168.129.1", "status": "attached", "updated_at": "2019-01-01T12:00:00.000Z", "zone": {"name": "us-south-1"}}], "updated_at": "2019-01-01T12:00:00.000Z", "zone": {"name": "us-south-1"}}`)
				}))
			})
			It(`Invoke UpdateTransitGatewayConnection successfully with retries`, func() {
				transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(transitGatewayApisService).ToNot(BeNil())
				transitGatewayApisService.EnableRetries(0, 0)

				// Construct an instance of the UpdateTransitGatewayConnectionOptions model
				updateTransitGatewayConnectionOptionsModel := new(transitgatewayapisv1.UpdateTransitGatewayConnectionOptions)
				updateTransitGatewayConnectionOptionsModel.TransitGatewayID = core.StringPtr("testString")
				updateTransitGatewayConnectionOptionsModel.ID = core.StringPtr("testString")
				updateTransitGatewayConnectionOptionsModel.Name = core.StringPtr("Transit_Service_BWTN_SJ_DL")
				updateTransitGatewayConnectionOptionsModel.PrefixFiltersDefault = core.StringPtr("permit")
				updateTransitGatewayConnectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := transitGatewayApisService.UpdateTransitGatewayConnectionWithContext(ctx, updateTransitGatewayConnectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				transitGatewayApisService.DisableRetries()
				result, response, operationErr := transitGatewayApisService.UpdateTransitGatewayConnection(updateTransitGatewayConnectionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = transitGatewayApisService.UpdateTransitGatewayConnectionWithContext(ctx, updateTransitGatewayConnectionOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(updateTransitGatewayConnectionPath))
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
					fmt.Fprintf(res, "%s", `{"base_connection_id": "975f58c1-afe7-469a-9727-7f3d720f2d32", "base_network_type": "classic", "cidr": "192.168.0.0/24", "created_at": "2019-01-01T12:00:00.000Z", "id": "1a15dca5-7e33-45e1-b7c5-bc690e569531", "local_bgp_asn": 64490, "local_gateway_ip": "192.168.100.1", "local_tunnel_ip": "192.168.129.2", "mtu": 9000, "name": "Transit_Service_BWTN_SJ_DL", "network_account_id": "NetworkAccountID", "network_id": "crn:v1:bluemix:public:is:us-south:a/123456::vpc:4727d842-f94f-4a2d-824a-9bc9b02c523b", "network_type": "vpc", "prefix_filters": [{"action": "permit", "before": "1a15dcab-7e40-45e1-b7c5-bc690eaa9782", "created_at": "2019-01-01T12:00:00.000Z", "ge": 0, "id": "1a15dcab-7e30-45e1-b7c5-bc690eaa9865", "le": 32, "prefix": "192.168.100.0/24", "updated_at": "2019-01-01T12:00:00.000Z"}], "prefix_filters_default": "permit", "remote_bgp_asn": 65010, "remote_gateway_ip": "10.242.63.12", "remote_tunnel_ip": "192.168.129.1", "request_status": "pending", "status": "attached", "tunnels": [{"base_network_type": "classic", "created_at": "2019-01-01T12:00:00.000Z", "id": "1a15dca5-7e33-45e1-b7c5-bc690e569531", "local_bgp_asn": 11, "local_gateway_ip": "10.242.63.12", "local_tunnel_ip": "192.168.100.20", "mtu": 9000, "name": "gre1", "network_account_id": "NetworkAccountID", "network_id": "crn:v1:bluemix:public:is:us-south:a/123456::vpc:4727d842-f94f-4a2d-824a-9bc9b02c523b", "remote_bgp_asn": 65010, "remote_gateway_ip": "10.242.33.22", "remote_tunnel_ip": "192.168.129.1", "status": "attached", "updated_at": "2019-01-01T12:00:00.000Z", "zone": {"name": "us-south-1"}}], "updated_at": "2019-01-01T12:00:00.000Z", "zone": {"name": "us-south-1"}}`)
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
				updateTransitGatewayConnectionOptionsModel.PrefixFiltersDefault = core.StringPtr("permit")
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
				updateTransitGatewayConnectionOptionsModel.PrefixFiltersDefault = core.StringPtr("permit")
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
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
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

				// Construct an instance of the UpdateTransitGatewayConnectionOptions model
				updateTransitGatewayConnectionOptionsModel := new(transitgatewayapisv1.UpdateTransitGatewayConnectionOptions)
				updateTransitGatewayConnectionOptionsModel.TransitGatewayID = core.StringPtr("testString")
				updateTransitGatewayConnectionOptionsModel.ID = core.StringPtr("testString")
				updateTransitGatewayConnectionOptionsModel.Name = core.StringPtr("Transit_Service_BWTN_SJ_DL")
				updateTransitGatewayConnectionOptionsModel.PrefixFiltersDefault = core.StringPtr("permit")
				updateTransitGatewayConnectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := transitGatewayApisService.UpdateTransitGatewayConnection(updateTransitGatewayConnectionOptionsModel)
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
	Describe(`CreateTransitGatewayConnectionActions(createTransitGatewayConnectionActionsOptions *CreateTransitGatewayConnectionActionsOptions)`, func() {
		version := "testString"
		createTransitGatewayConnectionActionsPath := "/transit_gateways/testString/connections/testString/actions"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createTransitGatewayConnectionActionsPath))
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
	Describe(`ListTransitGatewayGreTunnel(listTransitGatewayGreTunnelOptions *ListTransitGatewayGreTunnelOptions) - Operation response error`, func() {
		version := "testString"
		listTransitGatewayGreTunnelPath := "/transit_gateways/testString/connections/testString/tunnels"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listTransitGatewayGreTunnelPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListTransitGatewayGreTunnel with error: Operation response processing error`, func() {
				transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(transitGatewayApisService).ToNot(BeNil())

				// Construct an instance of the ListTransitGatewayGreTunnelOptions model
				listTransitGatewayGreTunnelOptionsModel := new(transitgatewayapisv1.ListTransitGatewayGreTunnelOptions)
				listTransitGatewayGreTunnelOptionsModel.TransitGatewayID = core.StringPtr("testString")
				listTransitGatewayGreTunnelOptionsModel.ID = core.StringPtr("testString")
				listTransitGatewayGreTunnelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := transitGatewayApisService.ListTransitGatewayGreTunnel(listTransitGatewayGreTunnelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				transitGatewayApisService.EnableRetries(0, 0)
				result, response, operationErr = transitGatewayApisService.ListTransitGatewayGreTunnel(listTransitGatewayGreTunnelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListTransitGatewayGreTunnel(listTransitGatewayGreTunnelOptions *ListTransitGatewayGreTunnelOptions)`, func() {
		version := "testString"
		listTransitGatewayGreTunnelPath := "/transit_gateways/testString/connections/testString/tunnels"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listTransitGatewayGreTunnelPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"tunnels": [{"base_network_type": "classic", "created_at": "2019-01-01T12:00:00.000Z", "id": "1a15dca5-7e33-45e1-b7c5-bc690e569531", "local_bgp_asn": 11, "local_gateway_ip": "10.242.63.12", "local_tunnel_ip": "192.168.100.20", "mtu": 9000, "name": "gre1", "network_account_id": "NetworkAccountID", "network_id": "crn:v1:bluemix:public:is:us-south:a/123456::vpc:4727d842-f94f-4a2d-824a-9bc9b02c523b", "remote_bgp_asn": 65010, "remote_gateway_ip": "10.242.33.22", "remote_tunnel_ip": "192.168.129.1", "status": "attached", "updated_at": "2019-01-01T12:00:00.000Z", "zone": {"name": "us-south-1"}}]}`)
				}))
			})
			It(`Invoke ListTransitGatewayGreTunnel successfully with retries`, func() {
				transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(transitGatewayApisService).ToNot(BeNil())
				transitGatewayApisService.EnableRetries(0, 0)

				// Construct an instance of the ListTransitGatewayGreTunnelOptions model
				listTransitGatewayGreTunnelOptionsModel := new(transitgatewayapisv1.ListTransitGatewayGreTunnelOptions)
				listTransitGatewayGreTunnelOptionsModel.TransitGatewayID = core.StringPtr("testString")
				listTransitGatewayGreTunnelOptionsModel.ID = core.StringPtr("testString")
				listTransitGatewayGreTunnelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := transitGatewayApisService.ListTransitGatewayGreTunnelWithContext(ctx, listTransitGatewayGreTunnelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				transitGatewayApisService.DisableRetries()
				result, response, operationErr := transitGatewayApisService.ListTransitGatewayGreTunnel(listTransitGatewayGreTunnelOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = transitGatewayApisService.ListTransitGatewayGreTunnelWithContext(ctx, listTransitGatewayGreTunnelOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listTransitGatewayGreTunnelPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"tunnels": [{"base_network_type": "classic", "created_at": "2019-01-01T12:00:00.000Z", "id": "1a15dca5-7e33-45e1-b7c5-bc690e569531", "local_bgp_asn": 11, "local_gateway_ip": "10.242.63.12", "local_tunnel_ip": "192.168.100.20", "mtu": 9000, "name": "gre1", "network_account_id": "NetworkAccountID", "network_id": "crn:v1:bluemix:public:is:us-south:a/123456::vpc:4727d842-f94f-4a2d-824a-9bc9b02c523b", "remote_bgp_asn": 65010, "remote_gateway_ip": "10.242.33.22", "remote_tunnel_ip": "192.168.129.1", "status": "attached", "updated_at": "2019-01-01T12:00:00.000Z", "zone": {"name": "us-south-1"}}]}`)
				}))
			})
			It(`Invoke ListTransitGatewayGreTunnel successfully`, func() {
				transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(transitGatewayApisService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := transitGatewayApisService.ListTransitGatewayGreTunnel(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListTransitGatewayGreTunnelOptions model
				listTransitGatewayGreTunnelOptionsModel := new(transitgatewayapisv1.ListTransitGatewayGreTunnelOptions)
				listTransitGatewayGreTunnelOptionsModel.TransitGatewayID = core.StringPtr("testString")
				listTransitGatewayGreTunnelOptionsModel.ID = core.StringPtr("testString")
				listTransitGatewayGreTunnelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = transitGatewayApisService.ListTransitGatewayGreTunnel(listTransitGatewayGreTunnelOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListTransitGatewayGreTunnel with error: Operation validation and request error`, func() {
				transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(transitGatewayApisService).ToNot(BeNil())

				// Construct an instance of the ListTransitGatewayGreTunnelOptions model
				listTransitGatewayGreTunnelOptionsModel := new(transitgatewayapisv1.ListTransitGatewayGreTunnelOptions)
				listTransitGatewayGreTunnelOptionsModel.TransitGatewayID = core.StringPtr("testString")
				listTransitGatewayGreTunnelOptionsModel.ID = core.StringPtr("testString")
				listTransitGatewayGreTunnelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := transitGatewayApisService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := transitGatewayApisService.ListTransitGatewayGreTunnel(listTransitGatewayGreTunnelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListTransitGatewayGreTunnelOptions model with no property values
				listTransitGatewayGreTunnelOptionsModelNew := new(transitgatewayapisv1.ListTransitGatewayGreTunnelOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = transitGatewayApisService.ListTransitGatewayGreTunnel(listTransitGatewayGreTunnelOptionsModelNew)
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
			It(`Invoke ListTransitGatewayGreTunnel successfully`, func() {
				transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(transitGatewayApisService).ToNot(BeNil())

				// Construct an instance of the ListTransitGatewayGreTunnelOptions model
				listTransitGatewayGreTunnelOptionsModel := new(transitgatewayapisv1.ListTransitGatewayGreTunnelOptions)
				listTransitGatewayGreTunnelOptionsModel.TransitGatewayID = core.StringPtr("testString")
				listTransitGatewayGreTunnelOptionsModel.ID = core.StringPtr("testString")
				listTransitGatewayGreTunnelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := transitGatewayApisService.ListTransitGatewayGreTunnel(listTransitGatewayGreTunnelOptionsModel)
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
	Describe(`CreateTransitGatewayGreTunnel(createTransitGatewayGreTunnelOptions *CreateTransitGatewayGreTunnelOptions) - Operation response error`, func() {
		version := "testString"
		createTransitGatewayGreTunnelPath := "/transit_gateways/testString/connections/testString/tunnels"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createTransitGatewayGreTunnelPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateTransitGatewayGreTunnel with error: Operation response processing error`, func() {
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

				// Construct an instance of the CreateTransitGatewayGreTunnelOptions model
				createTransitGatewayGreTunnelOptionsModel := new(transitgatewayapisv1.CreateTransitGatewayGreTunnelOptions)
				createTransitGatewayGreTunnelOptionsModel.TransitGatewayID = core.StringPtr("testString")
				createTransitGatewayGreTunnelOptionsModel.ID = core.StringPtr("testString")
				createTransitGatewayGreTunnelOptionsModel.LocalGatewayIp = core.StringPtr("10.242.63.12")
				createTransitGatewayGreTunnelOptionsModel.LocalTunnelIp = core.StringPtr("192.168.100.20")
				createTransitGatewayGreTunnelOptionsModel.Name = core.StringPtr("gre1")
				createTransitGatewayGreTunnelOptionsModel.RemoteGatewayIp = core.StringPtr("10.242.33.22")
				createTransitGatewayGreTunnelOptionsModel.RemoteTunnelIp = core.StringPtr("192.168.129.1")
				createTransitGatewayGreTunnelOptionsModel.Zone = zoneIdentityModel
				createTransitGatewayGreTunnelOptionsModel.RemoteBgpAsn = core.Int64Ptr(int64(65010))
				createTransitGatewayGreTunnelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := transitGatewayApisService.CreateTransitGatewayGreTunnel(createTransitGatewayGreTunnelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				transitGatewayApisService.EnableRetries(0, 0)
				result, response, operationErr = transitGatewayApisService.CreateTransitGatewayGreTunnel(createTransitGatewayGreTunnelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateTransitGatewayGreTunnel(createTransitGatewayGreTunnelOptions *CreateTransitGatewayGreTunnelOptions)`, func() {
		version := "testString"
		createTransitGatewayGreTunnelPath := "/transit_gateways/testString/connections/testString/tunnels"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createTransitGatewayGreTunnelPath))
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
					fmt.Fprintf(res, "%s", `{"base_network_type": "classic", "created_at": "2019-01-01T12:00:00.000Z", "id": "1a15dca5-7e33-45e1-b7c5-bc690e569531", "local_bgp_asn": 11, "local_gateway_ip": "10.242.63.12", "local_tunnel_ip": "192.168.100.20", "mtu": 9000, "name": "gre1", "network_account_id": "NetworkAccountID", "network_id": "crn:v1:bluemix:public:is:us-south:a/123456::vpc:4727d842-f94f-4a2d-824a-9bc9b02c523b", "remote_bgp_asn": 65010, "remote_gateway_ip": "10.242.33.22", "remote_tunnel_ip": "192.168.129.1", "status": "attached", "updated_at": "2019-01-01T12:00:00.000Z", "zone": {"name": "us-south-1"}}`)
				}))
			})
			It(`Invoke CreateTransitGatewayGreTunnel successfully with retries`, func() {
				transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(transitGatewayApisService).ToNot(BeNil())
				transitGatewayApisService.EnableRetries(0, 0)

				// Construct an instance of the ZoneIdentityByName model
				zoneIdentityModel := new(transitgatewayapisv1.ZoneIdentityByName)
				zoneIdentityModel.Name = core.StringPtr("us-south-1")

				// Construct an instance of the CreateTransitGatewayGreTunnelOptions model
				createTransitGatewayGreTunnelOptionsModel := new(transitgatewayapisv1.CreateTransitGatewayGreTunnelOptions)
				createTransitGatewayGreTunnelOptionsModel.TransitGatewayID = core.StringPtr("testString")
				createTransitGatewayGreTunnelOptionsModel.ID = core.StringPtr("testString")
				createTransitGatewayGreTunnelOptionsModel.LocalGatewayIp = core.StringPtr("10.242.63.12")
				createTransitGatewayGreTunnelOptionsModel.LocalTunnelIp = core.StringPtr("192.168.100.20")
				createTransitGatewayGreTunnelOptionsModel.Name = core.StringPtr("gre1")
				createTransitGatewayGreTunnelOptionsModel.RemoteGatewayIp = core.StringPtr("10.242.33.22")
				createTransitGatewayGreTunnelOptionsModel.RemoteTunnelIp = core.StringPtr("192.168.129.1")
				createTransitGatewayGreTunnelOptionsModel.Zone = zoneIdentityModel
				createTransitGatewayGreTunnelOptionsModel.RemoteBgpAsn = core.Int64Ptr(int64(65010))
				createTransitGatewayGreTunnelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := transitGatewayApisService.CreateTransitGatewayGreTunnelWithContext(ctx, createTransitGatewayGreTunnelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				transitGatewayApisService.DisableRetries()
				result, response, operationErr := transitGatewayApisService.CreateTransitGatewayGreTunnel(createTransitGatewayGreTunnelOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = transitGatewayApisService.CreateTransitGatewayGreTunnelWithContext(ctx, createTransitGatewayGreTunnelOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createTransitGatewayGreTunnelPath))
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
					fmt.Fprintf(res, "%s", `{"base_network_type": "classic", "created_at": "2019-01-01T12:00:00.000Z", "id": "1a15dca5-7e33-45e1-b7c5-bc690e569531", "local_bgp_asn": 11, "local_gateway_ip": "10.242.63.12", "local_tunnel_ip": "192.168.100.20", "mtu": 9000, "name": "gre1", "network_account_id": "NetworkAccountID", "network_id": "crn:v1:bluemix:public:is:us-south:a/123456::vpc:4727d842-f94f-4a2d-824a-9bc9b02c523b", "remote_bgp_asn": 65010, "remote_gateway_ip": "10.242.33.22", "remote_tunnel_ip": "192.168.129.1", "status": "attached", "updated_at": "2019-01-01T12:00:00.000Z", "zone": {"name": "us-south-1"}}`)
				}))
			})
			It(`Invoke CreateTransitGatewayGreTunnel successfully`, func() {
				transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(transitGatewayApisService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := transitGatewayApisService.CreateTransitGatewayGreTunnel(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ZoneIdentityByName model
				zoneIdentityModel := new(transitgatewayapisv1.ZoneIdentityByName)
				zoneIdentityModel.Name = core.StringPtr("us-south-1")

				// Construct an instance of the CreateTransitGatewayGreTunnelOptions model
				createTransitGatewayGreTunnelOptionsModel := new(transitgatewayapisv1.CreateTransitGatewayGreTunnelOptions)
				createTransitGatewayGreTunnelOptionsModel.TransitGatewayID = core.StringPtr("testString")
				createTransitGatewayGreTunnelOptionsModel.ID = core.StringPtr("testString")
				createTransitGatewayGreTunnelOptionsModel.LocalGatewayIp = core.StringPtr("10.242.63.12")
				createTransitGatewayGreTunnelOptionsModel.LocalTunnelIp = core.StringPtr("192.168.100.20")
				createTransitGatewayGreTunnelOptionsModel.Name = core.StringPtr("gre1")
				createTransitGatewayGreTunnelOptionsModel.RemoteGatewayIp = core.StringPtr("10.242.33.22")
				createTransitGatewayGreTunnelOptionsModel.RemoteTunnelIp = core.StringPtr("192.168.129.1")
				createTransitGatewayGreTunnelOptionsModel.Zone = zoneIdentityModel
				createTransitGatewayGreTunnelOptionsModel.RemoteBgpAsn = core.Int64Ptr(int64(65010))
				createTransitGatewayGreTunnelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = transitGatewayApisService.CreateTransitGatewayGreTunnel(createTransitGatewayGreTunnelOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateTransitGatewayGreTunnel with error: Operation validation and request error`, func() {
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

				// Construct an instance of the CreateTransitGatewayGreTunnelOptions model
				createTransitGatewayGreTunnelOptionsModel := new(transitgatewayapisv1.CreateTransitGatewayGreTunnelOptions)
				createTransitGatewayGreTunnelOptionsModel.TransitGatewayID = core.StringPtr("testString")
				createTransitGatewayGreTunnelOptionsModel.ID = core.StringPtr("testString")
				createTransitGatewayGreTunnelOptionsModel.LocalGatewayIp = core.StringPtr("10.242.63.12")
				createTransitGatewayGreTunnelOptionsModel.LocalTunnelIp = core.StringPtr("192.168.100.20")
				createTransitGatewayGreTunnelOptionsModel.Name = core.StringPtr("gre1")
				createTransitGatewayGreTunnelOptionsModel.RemoteGatewayIp = core.StringPtr("10.242.33.22")
				createTransitGatewayGreTunnelOptionsModel.RemoteTunnelIp = core.StringPtr("192.168.129.1")
				createTransitGatewayGreTunnelOptionsModel.Zone = zoneIdentityModel
				createTransitGatewayGreTunnelOptionsModel.RemoteBgpAsn = core.Int64Ptr(int64(65010))
				createTransitGatewayGreTunnelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := transitGatewayApisService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := transitGatewayApisService.CreateTransitGatewayGreTunnel(createTransitGatewayGreTunnelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateTransitGatewayGreTunnelOptions model with no property values
				createTransitGatewayGreTunnelOptionsModelNew := new(transitgatewayapisv1.CreateTransitGatewayGreTunnelOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = transitGatewayApisService.CreateTransitGatewayGreTunnel(createTransitGatewayGreTunnelOptionsModelNew)
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
			It(`Invoke CreateTransitGatewayGreTunnel successfully`, func() {
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

				// Construct an instance of the CreateTransitGatewayGreTunnelOptions model
				createTransitGatewayGreTunnelOptionsModel := new(transitgatewayapisv1.CreateTransitGatewayGreTunnelOptions)
				createTransitGatewayGreTunnelOptionsModel.TransitGatewayID = core.StringPtr("testString")
				createTransitGatewayGreTunnelOptionsModel.ID = core.StringPtr("testString")
				createTransitGatewayGreTunnelOptionsModel.LocalGatewayIp = core.StringPtr("10.242.63.12")
				createTransitGatewayGreTunnelOptionsModel.LocalTunnelIp = core.StringPtr("192.168.100.20")
				createTransitGatewayGreTunnelOptionsModel.Name = core.StringPtr("gre1")
				createTransitGatewayGreTunnelOptionsModel.RemoteGatewayIp = core.StringPtr("10.242.33.22")
				createTransitGatewayGreTunnelOptionsModel.RemoteTunnelIp = core.StringPtr("192.168.129.1")
				createTransitGatewayGreTunnelOptionsModel.Zone = zoneIdentityModel
				createTransitGatewayGreTunnelOptionsModel.RemoteBgpAsn = core.Int64Ptr(int64(65010))
				createTransitGatewayGreTunnelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := transitGatewayApisService.CreateTransitGatewayGreTunnel(createTransitGatewayGreTunnelOptionsModel)
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
	Describe(`DeleteTransitGatewayConnectionTunnels(deleteTransitGatewayConnectionTunnelsOptions *DeleteTransitGatewayConnectionTunnelsOptions)`, func() {
		version := "testString"
		deleteTransitGatewayConnectionTunnelsPath := "/transit_gateways/testString/connections/testString/tunnels/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteTransitGatewayConnectionTunnelsPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteTransitGatewayConnectionTunnels successfully`, func() {
				transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(transitGatewayApisService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := transitGatewayApisService.DeleteTransitGatewayConnectionTunnels(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteTransitGatewayConnectionTunnelsOptions model
				deleteTransitGatewayConnectionTunnelsOptionsModel := new(transitgatewayapisv1.DeleteTransitGatewayConnectionTunnelsOptions)
				deleteTransitGatewayConnectionTunnelsOptionsModel.TransitGatewayID = core.StringPtr("testString")
				deleteTransitGatewayConnectionTunnelsOptionsModel.ID = core.StringPtr("testString")
				deleteTransitGatewayConnectionTunnelsOptionsModel.GreTunnelID = core.StringPtr("testString")
				deleteTransitGatewayConnectionTunnelsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = transitGatewayApisService.DeleteTransitGatewayConnectionTunnels(deleteTransitGatewayConnectionTunnelsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteTransitGatewayConnectionTunnels with error: Operation validation and request error`, func() {
				transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(transitGatewayApisService).ToNot(BeNil())

				// Construct an instance of the DeleteTransitGatewayConnectionTunnelsOptions model
				deleteTransitGatewayConnectionTunnelsOptionsModel := new(transitgatewayapisv1.DeleteTransitGatewayConnectionTunnelsOptions)
				deleteTransitGatewayConnectionTunnelsOptionsModel.TransitGatewayID = core.StringPtr("testString")
				deleteTransitGatewayConnectionTunnelsOptionsModel.ID = core.StringPtr("testString")
				deleteTransitGatewayConnectionTunnelsOptionsModel.GreTunnelID = core.StringPtr("testString")
				deleteTransitGatewayConnectionTunnelsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := transitGatewayApisService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := transitGatewayApisService.DeleteTransitGatewayConnectionTunnels(deleteTransitGatewayConnectionTunnelsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteTransitGatewayConnectionTunnelsOptions model with no property values
				deleteTransitGatewayConnectionTunnelsOptionsModelNew := new(transitgatewayapisv1.DeleteTransitGatewayConnectionTunnelsOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = transitGatewayApisService.DeleteTransitGatewayConnectionTunnels(deleteTransitGatewayConnectionTunnelsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetTransitGatewayConnectionTunnels(getTransitGatewayConnectionTunnelsOptions *GetTransitGatewayConnectionTunnelsOptions) - Operation response error`, func() {
		version := "testString"
		getTransitGatewayConnectionTunnelsPath := "/transit_gateways/testString/connections/testString/tunnels/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getTransitGatewayConnectionTunnelsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetTransitGatewayConnectionTunnels with error: Operation response processing error`, func() {
				transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(transitGatewayApisService).ToNot(BeNil())

				// Construct an instance of the GetTransitGatewayConnectionTunnelsOptions model
				getTransitGatewayConnectionTunnelsOptionsModel := new(transitgatewayapisv1.GetTransitGatewayConnectionTunnelsOptions)
				getTransitGatewayConnectionTunnelsOptionsModel.TransitGatewayID = core.StringPtr("testString")
				getTransitGatewayConnectionTunnelsOptionsModel.ID = core.StringPtr("testString")
				getTransitGatewayConnectionTunnelsOptionsModel.GreTunnelID = core.StringPtr("testString")
				getTransitGatewayConnectionTunnelsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := transitGatewayApisService.GetTransitGatewayConnectionTunnels(getTransitGatewayConnectionTunnelsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				transitGatewayApisService.EnableRetries(0, 0)
				result, response, operationErr = transitGatewayApisService.GetTransitGatewayConnectionTunnels(getTransitGatewayConnectionTunnelsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetTransitGatewayConnectionTunnels(getTransitGatewayConnectionTunnelsOptions *GetTransitGatewayConnectionTunnelsOptions)`, func() {
		version := "testString"
		getTransitGatewayConnectionTunnelsPath := "/transit_gateways/testString/connections/testString/tunnels/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getTransitGatewayConnectionTunnelsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"base_network_type": "classic", "created_at": "2019-01-01T12:00:00.000Z", "id": "1a15dca5-7e33-45e1-b7c5-bc690e569531", "local_bgp_asn": 11, "local_gateway_ip": "10.242.63.12", "local_tunnel_ip": "192.168.100.20", "mtu": 9000, "name": "gre1", "network_account_id": "NetworkAccountID", "network_id": "crn:v1:bluemix:public:is:us-south:a/123456::vpc:4727d842-f94f-4a2d-824a-9bc9b02c523b", "remote_bgp_asn": 65010, "remote_gateway_ip": "10.242.33.22", "remote_tunnel_ip": "192.168.129.1", "status": "attached", "updated_at": "2019-01-01T12:00:00.000Z", "zone": {"name": "us-south-1"}}`)
				}))
			})
			It(`Invoke GetTransitGatewayConnectionTunnels successfully with retries`, func() {
				transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(transitGatewayApisService).ToNot(BeNil())
				transitGatewayApisService.EnableRetries(0, 0)

				// Construct an instance of the GetTransitGatewayConnectionTunnelsOptions model
				getTransitGatewayConnectionTunnelsOptionsModel := new(transitgatewayapisv1.GetTransitGatewayConnectionTunnelsOptions)
				getTransitGatewayConnectionTunnelsOptionsModel.TransitGatewayID = core.StringPtr("testString")
				getTransitGatewayConnectionTunnelsOptionsModel.ID = core.StringPtr("testString")
				getTransitGatewayConnectionTunnelsOptionsModel.GreTunnelID = core.StringPtr("testString")
				getTransitGatewayConnectionTunnelsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := transitGatewayApisService.GetTransitGatewayConnectionTunnelsWithContext(ctx, getTransitGatewayConnectionTunnelsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				transitGatewayApisService.DisableRetries()
				result, response, operationErr := transitGatewayApisService.GetTransitGatewayConnectionTunnels(getTransitGatewayConnectionTunnelsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = transitGatewayApisService.GetTransitGatewayConnectionTunnelsWithContext(ctx, getTransitGatewayConnectionTunnelsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getTransitGatewayConnectionTunnelsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"base_network_type": "classic", "created_at": "2019-01-01T12:00:00.000Z", "id": "1a15dca5-7e33-45e1-b7c5-bc690e569531", "local_bgp_asn": 11, "local_gateway_ip": "10.242.63.12", "local_tunnel_ip": "192.168.100.20", "mtu": 9000, "name": "gre1", "network_account_id": "NetworkAccountID", "network_id": "crn:v1:bluemix:public:is:us-south:a/123456::vpc:4727d842-f94f-4a2d-824a-9bc9b02c523b", "remote_bgp_asn": 65010, "remote_gateway_ip": "10.242.33.22", "remote_tunnel_ip": "192.168.129.1", "status": "attached", "updated_at": "2019-01-01T12:00:00.000Z", "zone": {"name": "us-south-1"}}`)
				}))
			})
			It(`Invoke GetTransitGatewayConnectionTunnels successfully`, func() {
				transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(transitGatewayApisService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := transitGatewayApisService.GetTransitGatewayConnectionTunnels(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetTransitGatewayConnectionTunnelsOptions model
				getTransitGatewayConnectionTunnelsOptionsModel := new(transitgatewayapisv1.GetTransitGatewayConnectionTunnelsOptions)
				getTransitGatewayConnectionTunnelsOptionsModel.TransitGatewayID = core.StringPtr("testString")
				getTransitGatewayConnectionTunnelsOptionsModel.ID = core.StringPtr("testString")
				getTransitGatewayConnectionTunnelsOptionsModel.GreTunnelID = core.StringPtr("testString")
				getTransitGatewayConnectionTunnelsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = transitGatewayApisService.GetTransitGatewayConnectionTunnels(getTransitGatewayConnectionTunnelsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetTransitGatewayConnectionTunnels with error: Operation validation and request error`, func() {
				transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(transitGatewayApisService).ToNot(BeNil())

				// Construct an instance of the GetTransitGatewayConnectionTunnelsOptions model
				getTransitGatewayConnectionTunnelsOptionsModel := new(transitgatewayapisv1.GetTransitGatewayConnectionTunnelsOptions)
				getTransitGatewayConnectionTunnelsOptionsModel.TransitGatewayID = core.StringPtr("testString")
				getTransitGatewayConnectionTunnelsOptionsModel.ID = core.StringPtr("testString")
				getTransitGatewayConnectionTunnelsOptionsModel.GreTunnelID = core.StringPtr("testString")
				getTransitGatewayConnectionTunnelsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := transitGatewayApisService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := transitGatewayApisService.GetTransitGatewayConnectionTunnels(getTransitGatewayConnectionTunnelsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetTransitGatewayConnectionTunnelsOptions model with no property values
				getTransitGatewayConnectionTunnelsOptionsModelNew := new(transitgatewayapisv1.GetTransitGatewayConnectionTunnelsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = transitGatewayApisService.GetTransitGatewayConnectionTunnels(getTransitGatewayConnectionTunnelsOptionsModelNew)
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
			It(`Invoke GetTransitGatewayConnectionTunnels successfully`, func() {
				transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(transitGatewayApisService).ToNot(BeNil())

				// Construct an instance of the GetTransitGatewayConnectionTunnelsOptions model
				getTransitGatewayConnectionTunnelsOptionsModel := new(transitgatewayapisv1.GetTransitGatewayConnectionTunnelsOptions)
				getTransitGatewayConnectionTunnelsOptionsModel.TransitGatewayID = core.StringPtr("testString")
				getTransitGatewayConnectionTunnelsOptionsModel.ID = core.StringPtr("testString")
				getTransitGatewayConnectionTunnelsOptionsModel.GreTunnelID = core.StringPtr("testString")
				getTransitGatewayConnectionTunnelsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := transitGatewayApisService.GetTransitGatewayConnectionTunnels(getTransitGatewayConnectionTunnelsOptionsModel)
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
	Describe(`UpdateTransitGatewayConnectionTunnels(updateTransitGatewayConnectionTunnelsOptions *UpdateTransitGatewayConnectionTunnelsOptions) - Operation response error`, func() {
		version := "testString"
		updateTransitGatewayConnectionTunnelsPath := "/transit_gateways/testString/connections/testString/tunnels/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateTransitGatewayConnectionTunnelsPath))
					Expect(req.Method).To(Equal("PATCH"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateTransitGatewayConnectionTunnels with error: Operation response processing error`, func() {
				transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(transitGatewayApisService).ToNot(BeNil())

				// Construct an instance of the TransitGatewayTunnelPatch model
				transitGatewayTunnelPatchModel := new(transitgatewayapisv1.TransitGatewayTunnelPatch)
				transitGatewayTunnelPatchModel.Name = core.StringPtr("gre2")
				transitGatewayTunnelPatchModelAsPatch, asPatchErr := transitGatewayTunnelPatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateTransitGatewayConnectionTunnelsOptions model
				updateTransitGatewayConnectionTunnelsOptionsModel := new(transitgatewayapisv1.UpdateTransitGatewayConnectionTunnelsOptions)
				updateTransitGatewayConnectionTunnelsOptionsModel.TransitGatewayID = core.StringPtr("testString")
				updateTransitGatewayConnectionTunnelsOptionsModel.ID = core.StringPtr("testString")
				updateTransitGatewayConnectionTunnelsOptionsModel.GreTunnelID = core.StringPtr("testString")
				updateTransitGatewayConnectionTunnelsOptionsModel.TransitGatewayTunnelPatch = transitGatewayTunnelPatchModelAsPatch
				updateTransitGatewayConnectionTunnelsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := transitGatewayApisService.UpdateTransitGatewayConnectionTunnels(updateTransitGatewayConnectionTunnelsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				transitGatewayApisService.EnableRetries(0, 0)
				result, response, operationErr = transitGatewayApisService.UpdateTransitGatewayConnectionTunnels(updateTransitGatewayConnectionTunnelsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateTransitGatewayConnectionTunnels(updateTransitGatewayConnectionTunnelsOptions *UpdateTransitGatewayConnectionTunnelsOptions)`, func() {
		version := "testString"
		updateTransitGatewayConnectionTunnelsPath := "/transit_gateways/testString/connections/testString/tunnels/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateTransitGatewayConnectionTunnelsPath))
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
					fmt.Fprintf(res, "%s", `{"base_network_type": "classic", "created_at": "2019-01-01T12:00:00.000Z", "id": "1a15dca5-7e33-45e1-b7c5-bc690e569531", "local_bgp_asn": 11, "local_gateway_ip": "10.242.63.12", "local_tunnel_ip": "192.168.100.20", "mtu": 9000, "name": "gre1", "network_account_id": "NetworkAccountID", "network_id": "crn:v1:bluemix:public:is:us-south:a/123456::vpc:4727d842-f94f-4a2d-824a-9bc9b02c523b", "remote_bgp_asn": 65010, "remote_gateway_ip": "10.242.33.22", "remote_tunnel_ip": "192.168.129.1", "status": "attached", "updated_at": "2019-01-01T12:00:00.000Z", "zone": {"name": "us-south-1"}}`)
				}))
			})
			It(`Invoke UpdateTransitGatewayConnectionTunnels successfully with retries`, func() {
				transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(transitGatewayApisService).ToNot(BeNil())
				transitGatewayApisService.EnableRetries(0, 0)

				// Construct an instance of the TransitGatewayTunnelPatch model
				transitGatewayTunnelPatchModel := new(transitgatewayapisv1.TransitGatewayTunnelPatch)
				transitGatewayTunnelPatchModel.Name = core.StringPtr("gre2")
				transitGatewayTunnelPatchModelAsPatch, asPatchErr := transitGatewayTunnelPatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateTransitGatewayConnectionTunnelsOptions model
				updateTransitGatewayConnectionTunnelsOptionsModel := new(transitgatewayapisv1.UpdateTransitGatewayConnectionTunnelsOptions)
				updateTransitGatewayConnectionTunnelsOptionsModel.TransitGatewayID = core.StringPtr("testString")
				updateTransitGatewayConnectionTunnelsOptionsModel.ID = core.StringPtr("testString")
				updateTransitGatewayConnectionTunnelsOptionsModel.GreTunnelID = core.StringPtr("testString")
				updateTransitGatewayConnectionTunnelsOptionsModel.TransitGatewayTunnelPatch = transitGatewayTunnelPatchModelAsPatch
				updateTransitGatewayConnectionTunnelsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := transitGatewayApisService.UpdateTransitGatewayConnectionTunnelsWithContext(ctx, updateTransitGatewayConnectionTunnelsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				transitGatewayApisService.DisableRetries()
				result, response, operationErr := transitGatewayApisService.UpdateTransitGatewayConnectionTunnels(updateTransitGatewayConnectionTunnelsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = transitGatewayApisService.UpdateTransitGatewayConnectionTunnelsWithContext(ctx, updateTransitGatewayConnectionTunnelsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(updateTransitGatewayConnectionTunnelsPath))
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
					fmt.Fprintf(res, "%s", `{"base_network_type": "classic", "created_at": "2019-01-01T12:00:00.000Z", "id": "1a15dca5-7e33-45e1-b7c5-bc690e569531", "local_bgp_asn": 11, "local_gateway_ip": "10.242.63.12", "local_tunnel_ip": "192.168.100.20", "mtu": 9000, "name": "gre1", "network_account_id": "NetworkAccountID", "network_id": "crn:v1:bluemix:public:is:us-south:a/123456::vpc:4727d842-f94f-4a2d-824a-9bc9b02c523b", "remote_bgp_asn": 65010, "remote_gateway_ip": "10.242.33.22", "remote_tunnel_ip": "192.168.129.1", "status": "attached", "updated_at": "2019-01-01T12:00:00.000Z", "zone": {"name": "us-south-1"}}`)
				}))
			})
			It(`Invoke UpdateTransitGatewayConnectionTunnels successfully`, func() {
				transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(transitGatewayApisService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := transitGatewayApisService.UpdateTransitGatewayConnectionTunnels(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the TransitGatewayTunnelPatch model
				transitGatewayTunnelPatchModel := new(transitgatewayapisv1.TransitGatewayTunnelPatch)
				transitGatewayTunnelPatchModel.Name = core.StringPtr("gre2")
				transitGatewayTunnelPatchModelAsPatch, asPatchErr := transitGatewayTunnelPatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateTransitGatewayConnectionTunnelsOptions model
				updateTransitGatewayConnectionTunnelsOptionsModel := new(transitgatewayapisv1.UpdateTransitGatewayConnectionTunnelsOptions)
				updateTransitGatewayConnectionTunnelsOptionsModel.TransitGatewayID = core.StringPtr("testString")
				updateTransitGatewayConnectionTunnelsOptionsModel.ID = core.StringPtr("testString")
				updateTransitGatewayConnectionTunnelsOptionsModel.GreTunnelID = core.StringPtr("testString")
				updateTransitGatewayConnectionTunnelsOptionsModel.TransitGatewayTunnelPatch = transitGatewayTunnelPatchModelAsPatch
				updateTransitGatewayConnectionTunnelsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = transitGatewayApisService.UpdateTransitGatewayConnectionTunnels(updateTransitGatewayConnectionTunnelsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateTransitGatewayConnectionTunnels with error: Operation validation and request error`, func() {
				transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(transitGatewayApisService).ToNot(BeNil())

				// Construct an instance of the TransitGatewayTunnelPatch model
				transitGatewayTunnelPatchModel := new(transitgatewayapisv1.TransitGatewayTunnelPatch)
				transitGatewayTunnelPatchModel.Name = core.StringPtr("gre2")
				transitGatewayTunnelPatchModelAsPatch, asPatchErr := transitGatewayTunnelPatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateTransitGatewayConnectionTunnelsOptions model
				updateTransitGatewayConnectionTunnelsOptionsModel := new(transitgatewayapisv1.UpdateTransitGatewayConnectionTunnelsOptions)
				updateTransitGatewayConnectionTunnelsOptionsModel.TransitGatewayID = core.StringPtr("testString")
				updateTransitGatewayConnectionTunnelsOptionsModel.ID = core.StringPtr("testString")
				updateTransitGatewayConnectionTunnelsOptionsModel.GreTunnelID = core.StringPtr("testString")
				updateTransitGatewayConnectionTunnelsOptionsModel.TransitGatewayTunnelPatch = transitGatewayTunnelPatchModelAsPatch
				updateTransitGatewayConnectionTunnelsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := transitGatewayApisService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := transitGatewayApisService.UpdateTransitGatewayConnectionTunnels(updateTransitGatewayConnectionTunnelsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateTransitGatewayConnectionTunnelsOptions model with no property values
				updateTransitGatewayConnectionTunnelsOptionsModelNew := new(transitgatewayapisv1.UpdateTransitGatewayConnectionTunnelsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = transitGatewayApisService.UpdateTransitGatewayConnectionTunnels(updateTransitGatewayConnectionTunnelsOptionsModelNew)
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
			It(`Invoke UpdateTransitGatewayConnectionTunnels successfully`, func() {
				transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(transitGatewayApisService).ToNot(BeNil())

				// Construct an instance of the TransitGatewayTunnelPatch model
				transitGatewayTunnelPatchModel := new(transitgatewayapisv1.TransitGatewayTunnelPatch)
				transitGatewayTunnelPatchModel.Name = core.StringPtr("gre2")
				transitGatewayTunnelPatchModelAsPatch, asPatchErr := transitGatewayTunnelPatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateTransitGatewayConnectionTunnelsOptions model
				updateTransitGatewayConnectionTunnelsOptionsModel := new(transitgatewayapisv1.UpdateTransitGatewayConnectionTunnelsOptions)
				updateTransitGatewayConnectionTunnelsOptionsModel.TransitGatewayID = core.StringPtr("testString")
				updateTransitGatewayConnectionTunnelsOptionsModel.ID = core.StringPtr("testString")
				updateTransitGatewayConnectionTunnelsOptionsModel.GreTunnelID = core.StringPtr("testString")
				updateTransitGatewayConnectionTunnelsOptionsModel.TransitGatewayTunnelPatch = transitGatewayTunnelPatchModelAsPatch
				updateTransitGatewayConnectionTunnelsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := transitGatewayApisService.UpdateTransitGatewayConnectionTunnels(updateTransitGatewayConnectionTunnelsOptionsModel)
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
	Describe(`ListGatewayLocations(listGatewayLocationsOptions *ListGatewayLocationsOptions) - Operation response error`, func() {
		version := "testString"
		listGatewayLocationsPath := "/locations"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listGatewayLocationsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
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

				// Enable retries and test again
				transitGatewayApisService.EnableRetries(0, 0)
				result, response, operationErr = transitGatewayApisService.ListGatewayLocations(listGatewayLocationsOptionsModel)
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
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listGatewayLocationsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"locations": [{"billing_location": "us", "name": "us-south", "type": "region"}]}`)
				}))
			})
			It(`Invoke ListGatewayLocations successfully with retries`, func() {
				transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(transitGatewayApisService).ToNot(BeNil())
				transitGatewayApisService.EnableRetries(0, 0)

				// Construct an instance of the ListGatewayLocationsOptions model
				listGatewayLocationsOptionsModel := new(transitgatewayapisv1.ListGatewayLocationsOptions)
				listGatewayLocationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := transitGatewayApisService.ListGatewayLocationsWithContext(ctx, listGatewayLocationsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				transitGatewayApisService.DisableRetries()
				result, response, operationErr := transitGatewayApisService.ListGatewayLocations(listGatewayLocationsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = transitGatewayApisService.ListGatewayLocationsWithContext(ctx, listGatewayLocationsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listGatewayLocationsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"locations": [{"billing_location": "us", "name": "us-south", "type": "region"}]}`)
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
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
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

				// Construct an instance of the ListGatewayLocationsOptions model
				listGatewayLocationsOptionsModel := new(transitgatewayapisv1.ListGatewayLocationsOptions)
				listGatewayLocationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := transitGatewayApisService.ListGatewayLocations(listGatewayLocationsOptionsModel)
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
	Describe(`GetGatewayLocation(getGatewayLocationOptions *GetGatewayLocationOptions) - Operation response error`, func() {
		version := "testString"
		getGatewayLocationPath := "/locations/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getGatewayLocationPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
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

				// Enable retries and test again
				transitGatewayApisService.EnableRetries(0, 0)
				result, response, operationErr = transitGatewayApisService.GetGatewayLocation(getGatewayLocationOptionsModel)
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
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getGatewayLocationPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"billing_location": "us", "name": "us-south", "type": "region", "local_connection_locations": [{"display_name": "Dallas", "name": "us-south", "supported_connection_types": ["SupportedConnectionTypes"], "type": "region"}], "zones": [{"name": "us-south-1"}]}`)
				}))
			})
			It(`Invoke GetGatewayLocation successfully with retries`, func() {
				transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(transitGatewayApisService).ToNot(BeNil())
				transitGatewayApisService.EnableRetries(0, 0)

				// Construct an instance of the GetGatewayLocationOptions model
				getGatewayLocationOptionsModel := new(transitgatewayapisv1.GetGatewayLocationOptions)
				getGatewayLocationOptionsModel.Name = core.StringPtr("testString")
				getGatewayLocationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := transitGatewayApisService.GetGatewayLocationWithContext(ctx, getGatewayLocationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				transitGatewayApisService.DisableRetries()
				result, response, operationErr := transitGatewayApisService.GetGatewayLocation(getGatewayLocationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = transitGatewayApisService.GetGatewayLocationWithContext(ctx, getGatewayLocationOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getGatewayLocationPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"billing_location": "us", "name": "us-south", "type": "region", "local_connection_locations": [{"display_name": "Dallas", "name": "us-south", "supported_connection_types": ["SupportedConnectionTypes"], "type": "region"}], "zones": [{"name": "us-south-1"}]}`)
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
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
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

				// Construct an instance of the GetGatewayLocationOptions model
				getGatewayLocationOptionsModel := new(transitgatewayapisv1.GetGatewayLocationOptions)
				getGatewayLocationOptionsModel.Name = core.StringPtr("testString")
				getGatewayLocationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := transitGatewayApisService.GetGatewayLocation(getGatewayLocationOptionsModel)
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
	Describe(`ListTransitGatewayConnectionPrefixFilters(listTransitGatewayConnectionPrefixFiltersOptions *ListTransitGatewayConnectionPrefixFiltersOptions) - Operation response error`, func() {
		version := "testString"
		listTransitGatewayConnectionPrefixFiltersPath := "/transit_gateways/testString/connections/testString/prefix_filters"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listTransitGatewayConnectionPrefixFiltersPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListTransitGatewayConnectionPrefixFilters with error: Operation response processing error`, func() {
				transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(transitGatewayApisService).ToNot(BeNil())

				// Construct an instance of the ListTransitGatewayConnectionPrefixFiltersOptions model
				listTransitGatewayConnectionPrefixFiltersOptionsModel := new(transitgatewayapisv1.ListTransitGatewayConnectionPrefixFiltersOptions)
				listTransitGatewayConnectionPrefixFiltersOptionsModel.TransitGatewayID = core.StringPtr("testString")
				listTransitGatewayConnectionPrefixFiltersOptionsModel.ID = core.StringPtr("testString")
				listTransitGatewayConnectionPrefixFiltersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := transitGatewayApisService.ListTransitGatewayConnectionPrefixFilters(listTransitGatewayConnectionPrefixFiltersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				transitGatewayApisService.EnableRetries(0, 0)
				result, response, operationErr = transitGatewayApisService.ListTransitGatewayConnectionPrefixFilters(listTransitGatewayConnectionPrefixFiltersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListTransitGatewayConnectionPrefixFilters(listTransitGatewayConnectionPrefixFiltersOptions *ListTransitGatewayConnectionPrefixFiltersOptions)`, func() {
		version := "testString"
		listTransitGatewayConnectionPrefixFiltersPath := "/transit_gateways/testString/connections/testString/prefix_filters"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listTransitGatewayConnectionPrefixFiltersPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"prefix_filters": [{"action": "permit", "before": "1a15dcab-7e40-45e1-b7c5-bc690eaa9782", "created_at": "2019-01-01T12:00:00.000Z", "ge": 0, "id": "1a15dcab-7e30-45e1-b7c5-bc690eaa9865", "le": 32, "prefix": "192.168.100.0/24", "updated_at": "2019-01-01T12:00:00.000Z"}]}`)
				}))
			})
			It(`Invoke ListTransitGatewayConnectionPrefixFilters successfully with retries`, func() {
				transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(transitGatewayApisService).ToNot(BeNil())
				transitGatewayApisService.EnableRetries(0, 0)

				// Construct an instance of the ListTransitGatewayConnectionPrefixFiltersOptions model
				listTransitGatewayConnectionPrefixFiltersOptionsModel := new(transitgatewayapisv1.ListTransitGatewayConnectionPrefixFiltersOptions)
				listTransitGatewayConnectionPrefixFiltersOptionsModel.TransitGatewayID = core.StringPtr("testString")
				listTransitGatewayConnectionPrefixFiltersOptionsModel.ID = core.StringPtr("testString")
				listTransitGatewayConnectionPrefixFiltersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := transitGatewayApisService.ListTransitGatewayConnectionPrefixFiltersWithContext(ctx, listTransitGatewayConnectionPrefixFiltersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				transitGatewayApisService.DisableRetries()
				result, response, operationErr := transitGatewayApisService.ListTransitGatewayConnectionPrefixFilters(listTransitGatewayConnectionPrefixFiltersOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = transitGatewayApisService.ListTransitGatewayConnectionPrefixFiltersWithContext(ctx, listTransitGatewayConnectionPrefixFiltersOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listTransitGatewayConnectionPrefixFiltersPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"prefix_filters": [{"action": "permit", "before": "1a15dcab-7e40-45e1-b7c5-bc690eaa9782", "created_at": "2019-01-01T12:00:00.000Z", "ge": 0, "id": "1a15dcab-7e30-45e1-b7c5-bc690eaa9865", "le": 32, "prefix": "192.168.100.0/24", "updated_at": "2019-01-01T12:00:00.000Z"}]}`)
				}))
			})
			It(`Invoke ListTransitGatewayConnectionPrefixFilters successfully`, func() {
				transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(transitGatewayApisService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := transitGatewayApisService.ListTransitGatewayConnectionPrefixFilters(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListTransitGatewayConnectionPrefixFiltersOptions model
				listTransitGatewayConnectionPrefixFiltersOptionsModel := new(transitgatewayapisv1.ListTransitGatewayConnectionPrefixFiltersOptions)
				listTransitGatewayConnectionPrefixFiltersOptionsModel.TransitGatewayID = core.StringPtr("testString")
				listTransitGatewayConnectionPrefixFiltersOptionsModel.ID = core.StringPtr("testString")
				listTransitGatewayConnectionPrefixFiltersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = transitGatewayApisService.ListTransitGatewayConnectionPrefixFilters(listTransitGatewayConnectionPrefixFiltersOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListTransitGatewayConnectionPrefixFilters with error: Operation validation and request error`, func() {
				transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(transitGatewayApisService).ToNot(BeNil())

				// Construct an instance of the ListTransitGatewayConnectionPrefixFiltersOptions model
				listTransitGatewayConnectionPrefixFiltersOptionsModel := new(transitgatewayapisv1.ListTransitGatewayConnectionPrefixFiltersOptions)
				listTransitGatewayConnectionPrefixFiltersOptionsModel.TransitGatewayID = core.StringPtr("testString")
				listTransitGatewayConnectionPrefixFiltersOptionsModel.ID = core.StringPtr("testString")
				listTransitGatewayConnectionPrefixFiltersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := transitGatewayApisService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := transitGatewayApisService.ListTransitGatewayConnectionPrefixFilters(listTransitGatewayConnectionPrefixFiltersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListTransitGatewayConnectionPrefixFiltersOptions model with no property values
				listTransitGatewayConnectionPrefixFiltersOptionsModelNew := new(transitgatewayapisv1.ListTransitGatewayConnectionPrefixFiltersOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = transitGatewayApisService.ListTransitGatewayConnectionPrefixFilters(listTransitGatewayConnectionPrefixFiltersOptionsModelNew)
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
			It(`Invoke ListTransitGatewayConnectionPrefixFilters successfully`, func() {
				transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(transitGatewayApisService).ToNot(BeNil())

				// Construct an instance of the ListTransitGatewayConnectionPrefixFiltersOptions model
				listTransitGatewayConnectionPrefixFiltersOptionsModel := new(transitgatewayapisv1.ListTransitGatewayConnectionPrefixFiltersOptions)
				listTransitGatewayConnectionPrefixFiltersOptionsModel.TransitGatewayID = core.StringPtr("testString")
				listTransitGatewayConnectionPrefixFiltersOptionsModel.ID = core.StringPtr("testString")
				listTransitGatewayConnectionPrefixFiltersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := transitGatewayApisService.ListTransitGatewayConnectionPrefixFilters(listTransitGatewayConnectionPrefixFiltersOptionsModel)
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
	Describe(`CreateTransitGatewayConnectionPrefixFilter(createTransitGatewayConnectionPrefixFilterOptions *CreateTransitGatewayConnectionPrefixFilterOptions) - Operation response error`, func() {
		version := "testString"
		createTransitGatewayConnectionPrefixFilterPath := "/transit_gateways/testString/connections/testString/prefix_filters"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createTransitGatewayConnectionPrefixFilterPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateTransitGatewayConnectionPrefixFilter with error: Operation response processing error`, func() {
				transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(transitGatewayApisService).ToNot(BeNil())

				// Construct an instance of the CreateTransitGatewayConnectionPrefixFilterOptions model
				createTransitGatewayConnectionPrefixFilterOptionsModel := new(transitgatewayapisv1.CreateTransitGatewayConnectionPrefixFilterOptions)
				createTransitGatewayConnectionPrefixFilterOptionsModel.TransitGatewayID = core.StringPtr("testString")
				createTransitGatewayConnectionPrefixFilterOptionsModel.ID = core.StringPtr("testString")
				createTransitGatewayConnectionPrefixFilterOptionsModel.Action = core.StringPtr("permit")
				createTransitGatewayConnectionPrefixFilterOptionsModel.Prefix = core.StringPtr("192.168.100.0/24")
				createTransitGatewayConnectionPrefixFilterOptionsModel.Before = core.StringPtr("1a15dcab-7e40-45e1-b7c5-bc690eaa9782")
				createTransitGatewayConnectionPrefixFilterOptionsModel.Ge = core.Int64Ptr(int64(0))
				createTransitGatewayConnectionPrefixFilterOptionsModel.Le = core.Int64Ptr(int64(32))
				createTransitGatewayConnectionPrefixFilterOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := transitGatewayApisService.CreateTransitGatewayConnectionPrefixFilter(createTransitGatewayConnectionPrefixFilterOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				transitGatewayApisService.EnableRetries(0, 0)
				result, response, operationErr = transitGatewayApisService.CreateTransitGatewayConnectionPrefixFilter(createTransitGatewayConnectionPrefixFilterOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateTransitGatewayConnectionPrefixFilter(createTransitGatewayConnectionPrefixFilterOptions *CreateTransitGatewayConnectionPrefixFilterOptions)`, func() {
		version := "testString"
		createTransitGatewayConnectionPrefixFilterPath := "/transit_gateways/testString/connections/testString/prefix_filters"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createTransitGatewayConnectionPrefixFilterPath))
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
					fmt.Fprintf(res, "%s", `{"action": "permit", "before": "1a15dcab-7e40-45e1-b7c5-bc690eaa9782", "created_at": "2019-01-01T12:00:00.000Z", "ge": 0, "id": "1a15dcab-7e30-45e1-b7c5-bc690eaa9865", "le": 32, "prefix": "192.168.100.0/24", "updated_at": "2019-01-01T12:00:00.000Z"}`)
				}))
			})
			It(`Invoke CreateTransitGatewayConnectionPrefixFilter successfully with retries`, func() {
				transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(transitGatewayApisService).ToNot(BeNil())
				transitGatewayApisService.EnableRetries(0, 0)

				// Construct an instance of the CreateTransitGatewayConnectionPrefixFilterOptions model
				createTransitGatewayConnectionPrefixFilterOptionsModel := new(transitgatewayapisv1.CreateTransitGatewayConnectionPrefixFilterOptions)
				createTransitGatewayConnectionPrefixFilterOptionsModel.TransitGatewayID = core.StringPtr("testString")
				createTransitGatewayConnectionPrefixFilterOptionsModel.ID = core.StringPtr("testString")
				createTransitGatewayConnectionPrefixFilterOptionsModel.Action = core.StringPtr("permit")
				createTransitGatewayConnectionPrefixFilterOptionsModel.Prefix = core.StringPtr("192.168.100.0/24")
				createTransitGatewayConnectionPrefixFilterOptionsModel.Before = core.StringPtr("1a15dcab-7e40-45e1-b7c5-bc690eaa9782")
				createTransitGatewayConnectionPrefixFilterOptionsModel.Ge = core.Int64Ptr(int64(0))
				createTransitGatewayConnectionPrefixFilterOptionsModel.Le = core.Int64Ptr(int64(32))
				createTransitGatewayConnectionPrefixFilterOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := transitGatewayApisService.CreateTransitGatewayConnectionPrefixFilterWithContext(ctx, createTransitGatewayConnectionPrefixFilterOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				transitGatewayApisService.DisableRetries()
				result, response, operationErr := transitGatewayApisService.CreateTransitGatewayConnectionPrefixFilter(createTransitGatewayConnectionPrefixFilterOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = transitGatewayApisService.CreateTransitGatewayConnectionPrefixFilterWithContext(ctx, createTransitGatewayConnectionPrefixFilterOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createTransitGatewayConnectionPrefixFilterPath))
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
					fmt.Fprintf(res, "%s", `{"action": "permit", "before": "1a15dcab-7e40-45e1-b7c5-bc690eaa9782", "created_at": "2019-01-01T12:00:00.000Z", "ge": 0, "id": "1a15dcab-7e30-45e1-b7c5-bc690eaa9865", "le": 32, "prefix": "192.168.100.0/24", "updated_at": "2019-01-01T12:00:00.000Z"}`)
				}))
			})
			It(`Invoke CreateTransitGatewayConnectionPrefixFilter successfully`, func() {
				transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(transitGatewayApisService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := transitGatewayApisService.CreateTransitGatewayConnectionPrefixFilter(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateTransitGatewayConnectionPrefixFilterOptions model
				createTransitGatewayConnectionPrefixFilterOptionsModel := new(transitgatewayapisv1.CreateTransitGatewayConnectionPrefixFilterOptions)
				createTransitGatewayConnectionPrefixFilterOptionsModel.TransitGatewayID = core.StringPtr("testString")
				createTransitGatewayConnectionPrefixFilterOptionsModel.ID = core.StringPtr("testString")
				createTransitGatewayConnectionPrefixFilterOptionsModel.Action = core.StringPtr("permit")
				createTransitGatewayConnectionPrefixFilterOptionsModel.Prefix = core.StringPtr("192.168.100.0/24")
				createTransitGatewayConnectionPrefixFilterOptionsModel.Before = core.StringPtr("1a15dcab-7e40-45e1-b7c5-bc690eaa9782")
				createTransitGatewayConnectionPrefixFilterOptionsModel.Ge = core.Int64Ptr(int64(0))
				createTransitGatewayConnectionPrefixFilterOptionsModel.Le = core.Int64Ptr(int64(32))
				createTransitGatewayConnectionPrefixFilterOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = transitGatewayApisService.CreateTransitGatewayConnectionPrefixFilter(createTransitGatewayConnectionPrefixFilterOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateTransitGatewayConnectionPrefixFilter with error: Operation validation and request error`, func() {
				transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(transitGatewayApisService).ToNot(BeNil())

				// Construct an instance of the CreateTransitGatewayConnectionPrefixFilterOptions model
				createTransitGatewayConnectionPrefixFilterOptionsModel := new(transitgatewayapisv1.CreateTransitGatewayConnectionPrefixFilterOptions)
				createTransitGatewayConnectionPrefixFilterOptionsModel.TransitGatewayID = core.StringPtr("testString")
				createTransitGatewayConnectionPrefixFilterOptionsModel.ID = core.StringPtr("testString")
				createTransitGatewayConnectionPrefixFilterOptionsModel.Action = core.StringPtr("permit")
				createTransitGatewayConnectionPrefixFilterOptionsModel.Prefix = core.StringPtr("192.168.100.0/24")
				createTransitGatewayConnectionPrefixFilterOptionsModel.Before = core.StringPtr("1a15dcab-7e40-45e1-b7c5-bc690eaa9782")
				createTransitGatewayConnectionPrefixFilterOptionsModel.Ge = core.Int64Ptr(int64(0))
				createTransitGatewayConnectionPrefixFilterOptionsModel.Le = core.Int64Ptr(int64(32))
				createTransitGatewayConnectionPrefixFilterOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := transitGatewayApisService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := transitGatewayApisService.CreateTransitGatewayConnectionPrefixFilter(createTransitGatewayConnectionPrefixFilterOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateTransitGatewayConnectionPrefixFilterOptions model with no property values
				createTransitGatewayConnectionPrefixFilterOptionsModelNew := new(transitgatewayapisv1.CreateTransitGatewayConnectionPrefixFilterOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = transitGatewayApisService.CreateTransitGatewayConnectionPrefixFilter(createTransitGatewayConnectionPrefixFilterOptionsModelNew)
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
			It(`Invoke CreateTransitGatewayConnectionPrefixFilter successfully`, func() {
				transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(transitGatewayApisService).ToNot(BeNil())

				// Construct an instance of the CreateTransitGatewayConnectionPrefixFilterOptions model
				createTransitGatewayConnectionPrefixFilterOptionsModel := new(transitgatewayapisv1.CreateTransitGatewayConnectionPrefixFilterOptions)
				createTransitGatewayConnectionPrefixFilterOptionsModel.TransitGatewayID = core.StringPtr("testString")
				createTransitGatewayConnectionPrefixFilterOptionsModel.ID = core.StringPtr("testString")
				createTransitGatewayConnectionPrefixFilterOptionsModel.Action = core.StringPtr("permit")
				createTransitGatewayConnectionPrefixFilterOptionsModel.Prefix = core.StringPtr("192.168.100.0/24")
				createTransitGatewayConnectionPrefixFilterOptionsModel.Before = core.StringPtr("1a15dcab-7e40-45e1-b7c5-bc690eaa9782")
				createTransitGatewayConnectionPrefixFilterOptionsModel.Ge = core.Int64Ptr(int64(0))
				createTransitGatewayConnectionPrefixFilterOptionsModel.Le = core.Int64Ptr(int64(32))
				createTransitGatewayConnectionPrefixFilterOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := transitGatewayApisService.CreateTransitGatewayConnectionPrefixFilter(createTransitGatewayConnectionPrefixFilterOptionsModel)
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
	Describe(`ReplaceTransitGatewayConnectionPrefixFilter(replaceTransitGatewayConnectionPrefixFilterOptions *ReplaceTransitGatewayConnectionPrefixFilterOptions) - Operation response error`, func() {
		version := "testString"
		replaceTransitGatewayConnectionPrefixFilterPath := "/transit_gateways/testString/connections/testString/prefix_filters"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(replaceTransitGatewayConnectionPrefixFilterPath))
					Expect(req.Method).To(Equal("PUT"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ReplaceTransitGatewayConnectionPrefixFilter with error: Operation response processing error`, func() {
				transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(transitGatewayApisService).ToNot(BeNil())

				// Construct an instance of the PrefixFilterPut model
				prefixFilterPutModel := new(transitgatewayapisv1.PrefixFilterPut)
				prefixFilterPutModel.Action = core.StringPtr("permit")
				prefixFilterPutModel.Ge = core.Int64Ptr(int64(0))
				prefixFilterPutModel.Le = core.Int64Ptr(int64(32))
				prefixFilterPutModel.Prefix = core.StringPtr("192.168.100.0/24")

				// Construct an instance of the ReplaceTransitGatewayConnectionPrefixFilterOptions model
				replaceTransitGatewayConnectionPrefixFilterOptionsModel := new(transitgatewayapisv1.ReplaceTransitGatewayConnectionPrefixFilterOptions)
				replaceTransitGatewayConnectionPrefixFilterOptionsModel.TransitGatewayID = core.StringPtr("testString")
				replaceTransitGatewayConnectionPrefixFilterOptionsModel.ID = core.StringPtr("testString")
				replaceTransitGatewayConnectionPrefixFilterOptionsModel.PrefixFilters = []transitgatewayapisv1.PrefixFilterPut{*prefixFilterPutModel}
				replaceTransitGatewayConnectionPrefixFilterOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := transitGatewayApisService.ReplaceTransitGatewayConnectionPrefixFilter(replaceTransitGatewayConnectionPrefixFilterOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				transitGatewayApisService.EnableRetries(0, 0)
				result, response, operationErr = transitGatewayApisService.ReplaceTransitGatewayConnectionPrefixFilter(replaceTransitGatewayConnectionPrefixFilterOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ReplaceTransitGatewayConnectionPrefixFilter(replaceTransitGatewayConnectionPrefixFilterOptions *ReplaceTransitGatewayConnectionPrefixFilterOptions)`, func() {
		version := "testString"
		replaceTransitGatewayConnectionPrefixFilterPath := "/transit_gateways/testString/connections/testString/prefix_filters"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(replaceTransitGatewayConnectionPrefixFilterPath))
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

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"prefix_filters": [{"action": "permit", "before": "1a15dcab-7e40-45e1-b7c5-bc690eaa9782", "created_at": "2019-01-01T12:00:00.000Z", "ge": 0, "id": "1a15dcab-7e30-45e1-b7c5-bc690eaa9865", "le": 32, "prefix": "192.168.100.0/24", "updated_at": "2019-01-01T12:00:00.000Z"}]}`)
				}))
			})
			It(`Invoke ReplaceTransitGatewayConnectionPrefixFilter successfully with retries`, func() {
				transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(transitGatewayApisService).ToNot(BeNil())
				transitGatewayApisService.EnableRetries(0, 0)

				// Construct an instance of the PrefixFilterPut model
				prefixFilterPutModel := new(transitgatewayapisv1.PrefixFilterPut)
				prefixFilterPutModel.Action = core.StringPtr("permit")
				prefixFilterPutModel.Ge = core.Int64Ptr(int64(0))
				prefixFilterPutModel.Le = core.Int64Ptr(int64(32))
				prefixFilterPutModel.Prefix = core.StringPtr("192.168.100.0/24")

				// Construct an instance of the ReplaceTransitGatewayConnectionPrefixFilterOptions model
				replaceTransitGatewayConnectionPrefixFilterOptionsModel := new(transitgatewayapisv1.ReplaceTransitGatewayConnectionPrefixFilterOptions)
				replaceTransitGatewayConnectionPrefixFilterOptionsModel.TransitGatewayID = core.StringPtr("testString")
				replaceTransitGatewayConnectionPrefixFilterOptionsModel.ID = core.StringPtr("testString")
				replaceTransitGatewayConnectionPrefixFilterOptionsModel.PrefixFilters = []transitgatewayapisv1.PrefixFilterPut{*prefixFilterPutModel}
				replaceTransitGatewayConnectionPrefixFilterOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := transitGatewayApisService.ReplaceTransitGatewayConnectionPrefixFilterWithContext(ctx, replaceTransitGatewayConnectionPrefixFilterOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				transitGatewayApisService.DisableRetries()
				result, response, operationErr := transitGatewayApisService.ReplaceTransitGatewayConnectionPrefixFilter(replaceTransitGatewayConnectionPrefixFilterOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = transitGatewayApisService.ReplaceTransitGatewayConnectionPrefixFilterWithContext(ctx, replaceTransitGatewayConnectionPrefixFilterOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(replaceTransitGatewayConnectionPrefixFilterPath))
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

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"prefix_filters": [{"action": "permit", "before": "1a15dcab-7e40-45e1-b7c5-bc690eaa9782", "created_at": "2019-01-01T12:00:00.000Z", "ge": 0, "id": "1a15dcab-7e30-45e1-b7c5-bc690eaa9865", "le": 32, "prefix": "192.168.100.0/24", "updated_at": "2019-01-01T12:00:00.000Z"}]}`)
				}))
			})
			It(`Invoke ReplaceTransitGatewayConnectionPrefixFilter successfully`, func() {
				transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(transitGatewayApisService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := transitGatewayApisService.ReplaceTransitGatewayConnectionPrefixFilter(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the PrefixFilterPut model
				prefixFilterPutModel := new(transitgatewayapisv1.PrefixFilterPut)
				prefixFilterPutModel.Action = core.StringPtr("permit")
				prefixFilterPutModel.Ge = core.Int64Ptr(int64(0))
				prefixFilterPutModel.Le = core.Int64Ptr(int64(32))
				prefixFilterPutModel.Prefix = core.StringPtr("192.168.100.0/24")

				// Construct an instance of the ReplaceTransitGatewayConnectionPrefixFilterOptions model
				replaceTransitGatewayConnectionPrefixFilterOptionsModel := new(transitgatewayapisv1.ReplaceTransitGatewayConnectionPrefixFilterOptions)
				replaceTransitGatewayConnectionPrefixFilterOptionsModel.TransitGatewayID = core.StringPtr("testString")
				replaceTransitGatewayConnectionPrefixFilterOptionsModel.ID = core.StringPtr("testString")
				replaceTransitGatewayConnectionPrefixFilterOptionsModel.PrefixFilters = []transitgatewayapisv1.PrefixFilterPut{*prefixFilterPutModel}
				replaceTransitGatewayConnectionPrefixFilterOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = transitGatewayApisService.ReplaceTransitGatewayConnectionPrefixFilter(replaceTransitGatewayConnectionPrefixFilterOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ReplaceTransitGatewayConnectionPrefixFilter with error: Operation validation and request error`, func() {
				transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(transitGatewayApisService).ToNot(BeNil())

				// Construct an instance of the PrefixFilterPut model
				prefixFilterPutModel := new(transitgatewayapisv1.PrefixFilterPut)
				prefixFilterPutModel.Action = core.StringPtr("permit")
				prefixFilterPutModel.Ge = core.Int64Ptr(int64(0))
				prefixFilterPutModel.Le = core.Int64Ptr(int64(32))
				prefixFilterPutModel.Prefix = core.StringPtr("192.168.100.0/24")

				// Construct an instance of the ReplaceTransitGatewayConnectionPrefixFilterOptions model
				replaceTransitGatewayConnectionPrefixFilterOptionsModel := new(transitgatewayapisv1.ReplaceTransitGatewayConnectionPrefixFilterOptions)
				replaceTransitGatewayConnectionPrefixFilterOptionsModel.TransitGatewayID = core.StringPtr("testString")
				replaceTransitGatewayConnectionPrefixFilterOptionsModel.ID = core.StringPtr("testString")
				replaceTransitGatewayConnectionPrefixFilterOptionsModel.PrefixFilters = []transitgatewayapisv1.PrefixFilterPut{*prefixFilterPutModel}
				replaceTransitGatewayConnectionPrefixFilterOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := transitGatewayApisService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := transitGatewayApisService.ReplaceTransitGatewayConnectionPrefixFilter(replaceTransitGatewayConnectionPrefixFilterOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ReplaceTransitGatewayConnectionPrefixFilterOptions model with no property values
				replaceTransitGatewayConnectionPrefixFilterOptionsModelNew := new(transitgatewayapisv1.ReplaceTransitGatewayConnectionPrefixFilterOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = transitGatewayApisService.ReplaceTransitGatewayConnectionPrefixFilter(replaceTransitGatewayConnectionPrefixFilterOptionsModelNew)
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
			It(`Invoke ReplaceTransitGatewayConnectionPrefixFilter successfully`, func() {
				transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(transitGatewayApisService).ToNot(BeNil())

				// Construct an instance of the PrefixFilterPut model
				prefixFilterPutModel := new(transitgatewayapisv1.PrefixFilterPut)
				prefixFilterPutModel.Action = core.StringPtr("permit")
				prefixFilterPutModel.Ge = core.Int64Ptr(int64(0))
				prefixFilterPutModel.Le = core.Int64Ptr(int64(32))
				prefixFilterPutModel.Prefix = core.StringPtr("192.168.100.0/24")

				// Construct an instance of the ReplaceTransitGatewayConnectionPrefixFilterOptions model
				replaceTransitGatewayConnectionPrefixFilterOptionsModel := new(transitgatewayapisv1.ReplaceTransitGatewayConnectionPrefixFilterOptions)
				replaceTransitGatewayConnectionPrefixFilterOptionsModel.TransitGatewayID = core.StringPtr("testString")
				replaceTransitGatewayConnectionPrefixFilterOptionsModel.ID = core.StringPtr("testString")
				replaceTransitGatewayConnectionPrefixFilterOptionsModel.PrefixFilters = []transitgatewayapisv1.PrefixFilterPut{*prefixFilterPutModel}
				replaceTransitGatewayConnectionPrefixFilterOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := transitGatewayApisService.ReplaceTransitGatewayConnectionPrefixFilter(replaceTransitGatewayConnectionPrefixFilterOptionsModel)
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
	Describe(`DeleteTransitGatewayConnectionPrefixFilter(deleteTransitGatewayConnectionPrefixFilterOptions *DeleteTransitGatewayConnectionPrefixFilterOptions)`, func() {
		version := "testString"
		deleteTransitGatewayConnectionPrefixFilterPath := "/transit_gateways/testString/connections/testString/prefix_filters/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteTransitGatewayConnectionPrefixFilterPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteTransitGatewayConnectionPrefixFilter successfully`, func() {
				transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(transitGatewayApisService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := transitGatewayApisService.DeleteTransitGatewayConnectionPrefixFilter(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteTransitGatewayConnectionPrefixFilterOptions model
				deleteTransitGatewayConnectionPrefixFilterOptionsModel := new(transitgatewayapisv1.DeleteTransitGatewayConnectionPrefixFilterOptions)
				deleteTransitGatewayConnectionPrefixFilterOptionsModel.TransitGatewayID = core.StringPtr("testString")
				deleteTransitGatewayConnectionPrefixFilterOptionsModel.ID = core.StringPtr("testString")
				deleteTransitGatewayConnectionPrefixFilterOptionsModel.FilterID = core.StringPtr("testString")
				deleteTransitGatewayConnectionPrefixFilterOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = transitGatewayApisService.DeleteTransitGatewayConnectionPrefixFilter(deleteTransitGatewayConnectionPrefixFilterOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteTransitGatewayConnectionPrefixFilter with error: Operation validation and request error`, func() {
				transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(transitGatewayApisService).ToNot(BeNil())

				// Construct an instance of the DeleteTransitGatewayConnectionPrefixFilterOptions model
				deleteTransitGatewayConnectionPrefixFilterOptionsModel := new(transitgatewayapisv1.DeleteTransitGatewayConnectionPrefixFilterOptions)
				deleteTransitGatewayConnectionPrefixFilterOptionsModel.TransitGatewayID = core.StringPtr("testString")
				deleteTransitGatewayConnectionPrefixFilterOptionsModel.ID = core.StringPtr("testString")
				deleteTransitGatewayConnectionPrefixFilterOptionsModel.FilterID = core.StringPtr("testString")
				deleteTransitGatewayConnectionPrefixFilterOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := transitGatewayApisService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := transitGatewayApisService.DeleteTransitGatewayConnectionPrefixFilter(deleteTransitGatewayConnectionPrefixFilterOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteTransitGatewayConnectionPrefixFilterOptions model with no property values
				deleteTransitGatewayConnectionPrefixFilterOptionsModelNew := new(transitgatewayapisv1.DeleteTransitGatewayConnectionPrefixFilterOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = transitGatewayApisService.DeleteTransitGatewayConnectionPrefixFilter(deleteTransitGatewayConnectionPrefixFilterOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetTransitGatewayConnectionPrefixFilter(getTransitGatewayConnectionPrefixFilterOptions *GetTransitGatewayConnectionPrefixFilterOptions) - Operation response error`, func() {
		version := "testString"
		getTransitGatewayConnectionPrefixFilterPath := "/transit_gateways/testString/connections/testString/prefix_filters/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getTransitGatewayConnectionPrefixFilterPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetTransitGatewayConnectionPrefixFilter with error: Operation response processing error`, func() {
				transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(transitGatewayApisService).ToNot(BeNil())

				// Construct an instance of the GetTransitGatewayConnectionPrefixFilterOptions model
				getTransitGatewayConnectionPrefixFilterOptionsModel := new(transitgatewayapisv1.GetTransitGatewayConnectionPrefixFilterOptions)
				getTransitGatewayConnectionPrefixFilterOptionsModel.TransitGatewayID = core.StringPtr("testString")
				getTransitGatewayConnectionPrefixFilterOptionsModel.ID = core.StringPtr("testString")
				getTransitGatewayConnectionPrefixFilterOptionsModel.FilterID = core.StringPtr("testString")
				getTransitGatewayConnectionPrefixFilterOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := transitGatewayApisService.GetTransitGatewayConnectionPrefixFilter(getTransitGatewayConnectionPrefixFilterOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				transitGatewayApisService.EnableRetries(0, 0)
				result, response, operationErr = transitGatewayApisService.GetTransitGatewayConnectionPrefixFilter(getTransitGatewayConnectionPrefixFilterOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetTransitGatewayConnectionPrefixFilter(getTransitGatewayConnectionPrefixFilterOptions *GetTransitGatewayConnectionPrefixFilterOptions)`, func() {
		version := "testString"
		getTransitGatewayConnectionPrefixFilterPath := "/transit_gateways/testString/connections/testString/prefix_filters/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getTransitGatewayConnectionPrefixFilterPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"action": "permit", "before": "1a15dcab-7e40-45e1-b7c5-bc690eaa9782", "created_at": "2019-01-01T12:00:00.000Z", "ge": 0, "id": "1a15dcab-7e30-45e1-b7c5-bc690eaa9865", "le": 32, "prefix": "192.168.100.0/24", "updated_at": "2019-01-01T12:00:00.000Z"}`)
				}))
			})
			It(`Invoke GetTransitGatewayConnectionPrefixFilter successfully with retries`, func() {
				transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(transitGatewayApisService).ToNot(BeNil())
				transitGatewayApisService.EnableRetries(0, 0)

				// Construct an instance of the GetTransitGatewayConnectionPrefixFilterOptions model
				getTransitGatewayConnectionPrefixFilterOptionsModel := new(transitgatewayapisv1.GetTransitGatewayConnectionPrefixFilterOptions)
				getTransitGatewayConnectionPrefixFilterOptionsModel.TransitGatewayID = core.StringPtr("testString")
				getTransitGatewayConnectionPrefixFilterOptionsModel.ID = core.StringPtr("testString")
				getTransitGatewayConnectionPrefixFilterOptionsModel.FilterID = core.StringPtr("testString")
				getTransitGatewayConnectionPrefixFilterOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := transitGatewayApisService.GetTransitGatewayConnectionPrefixFilterWithContext(ctx, getTransitGatewayConnectionPrefixFilterOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				transitGatewayApisService.DisableRetries()
				result, response, operationErr := transitGatewayApisService.GetTransitGatewayConnectionPrefixFilter(getTransitGatewayConnectionPrefixFilterOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = transitGatewayApisService.GetTransitGatewayConnectionPrefixFilterWithContext(ctx, getTransitGatewayConnectionPrefixFilterOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getTransitGatewayConnectionPrefixFilterPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"action": "permit", "before": "1a15dcab-7e40-45e1-b7c5-bc690eaa9782", "created_at": "2019-01-01T12:00:00.000Z", "ge": 0, "id": "1a15dcab-7e30-45e1-b7c5-bc690eaa9865", "le": 32, "prefix": "192.168.100.0/24", "updated_at": "2019-01-01T12:00:00.000Z"}`)
				}))
			})
			It(`Invoke GetTransitGatewayConnectionPrefixFilter successfully`, func() {
				transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(transitGatewayApisService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := transitGatewayApisService.GetTransitGatewayConnectionPrefixFilter(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetTransitGatewayConnectionPrefixFilterOptions model
				getTransitGatewayConnectionPrefixFilterOptionsModel := new(transitgatewayapisv1.GetTransitGatewayConnectionPrefixFilterOptions)
				getTransitGatewayConnectionPrefixFilterOptionsModel.TransitGatewayID = core.StringPtr("testString")
				getTransitGatewayConnectionPrefixFilterOptionsModel.ID = core.StringPtr("testString")
				getTransitGatewayConnectionPrefixFilterOptionsModel.FilterID = core.StringPtr("testString")
				getTransitGatewayConnectionPrefixFilterOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = transitGatewayApisService.GetTransitGatewayConnectionPrefixFilter(getTransitGatewayConnectionPrefixFilterOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetTransitGatewayConnectionPrefixFilter with error: Operation validation and request error`, func() {
				transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(transitGatewayApisService).ToNot(BeNil())

				// Construct an instance of the GetTransitGatewayConnectionPrefixFilterOptions model
				getTransitGatewayConnectionPrefixFilterOptionsModel := new(transitgatewayapisv1.GetTransitGatewayConnectionPrefixFilterOptions)
				getTransitGatewayConnectionPrefixFilterOptionsModel.TransitGatewayID = core.StringPtr("testString")
				getTransitGatewayConnectionPrefixFilterOptionsModel.ID = core.StringPtr("testString")
				getTransitGatewayConnectionPrefixFilterOptionsModel.FilterID = core.StringPtr("testString")
				getTransitGatewayConnectionPrefixFilterOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := transitGatewayApisService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := transitGatewayApisService.GetTransitGatewayConnectionPrefixFilter(getTransitGatewayConnectionPrefixFilterOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetTransitGatewayConnectionPrefixFilterOptions model with no property values
				getTransitGatewayConnectionPrefixFilterOptionsModelNew := new(transitgatewayapisv1.GetTransitGatewayConnectionPrefixFilterOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = transitGatewayApisService.GetTransitGatewayConnectionPrefixFilter(getTransitGatewayConnectionPrefixFilterOptionsModelNew)
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
			It(`Invoke GetTransitGatewayConnectionPrefixFilter successfully`, func() {
				transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(transitGatewayApisService).ToNot(BeNil())

				// Construct an instance of the GetTransitGatewayConnectionPrefixFilterOptions model
				getTransitGatewayConnectionPrefixFilterOptionsModel := new(transitgatewayapisv1.GetTransitGatewayConnectionPrefixFilterOptions)
				getTransitGatewayConnectionPrefixFilterOptionsModel.TransitGatewayID = core.StringPtr("testString")
				getTransitGatewayConnectionPrefixFilterOptionsModel.ID = core.StringPtr("testString")
				getTransitGatewayConnectionPrefixFilterOptionsModel.FilterID = core.StringPtr("testString")
				getTransitGatewayConnectionPrefixFilterOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := transitGatewayApisService.GetTransitGatewayConnectionPrefixFilter(getTransitGatewayConnectionPrefixFilterOptionsModel)
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
	Describe(`UpdateTransitGatewayConnectionPrefixFilter(updateTransitGatewayConnectionPrefixFilterOptions *UpdateTransitGatewayConnectionPrefixFilterOptions) - Operation response error`, func() {
		version := "testString"
		updateTransitGatewayConnectionPrefixFilterPath := "/transit_gateways/testString/connections/testString/prefix_filters/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateTransitGatewayConnectionPrefixFilterPath))
					Expect(req.Method).To(Equal("PATCH"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateTransitGatewayConnectionPrefixFilter with error: Operation response processing error`, func() {
				transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(transitGatewayApisService).ToNot(BeNil())

				// Construct an instance of the UpdateTransitGatewayConnectionPrefixFilterOptions model
				updateTransitGatewayConnectionPrefixFilterOptionsModel := new(transitgatewayapisv1.UpdateTransitGatewayConnectionPrefixFilterOptions)
				updateTransitGatewayConnectionPrefixFilterOptionsModel.TransitGatewayID = core.StringPtr("testString")
				updateTransitGatewayConnectionPrefixFilterOptionsModel.ID = core.StringPtr("testString")
				updateTransitGatewayConnectionPrefixFilterOptionsModel.FilterID = core.StringPtr("testString")
				updateTransitGatewayConnectionPrefixFilterOptionsModel.Action = core.StringPtr("permit")
				updateTransitGatewayConnectionPrefixFilterOptionsModel.Before = core.StringPtr("1a15dcab-7e40-45e1-b7c5-bc690eaa9782")
				updateTransitGatewayConnectionPrefixFilterOptionsModel.Ge = core.Int64Ptr(int64(0))
				updateTransitGatewayConnectionPrefixFilterOptionsModel.Le = core.Int64Ptr(int64(32))
				updateTransitGatewayConnectionPrefixFilterOptionsModel.Prefix = core.StringPtr("192.168.100.0/24")
				updateTransitGatewayConnectionPrefixFilterOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := transitGatewayApisService.UpdateTransitGatewayConnectionPrefixFilter(updateTransitGatewayConnectionPrefixFilterOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				transitGatewayApisService.EnableRetries(0, 0)
				result, response, operationErr = transitGatewayApisService.UpdateTransitGatewayConnectionPrefixFilter(updateTransitGatewayConnectionPrefixFilterOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateTransitGatewayConnectionPrefixFilter(updateTransitGatewayConnectionPrefixFilterOptions *UpdateTransitGatewayConnectionPrefixFilterOptions)`, func() {
		version := "testString"
		updateTransitGatewayConnectionPrefixFilterPath := "/transit_gateways/testString/connections/testString/prefix_filters/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateTransitGatewayConnectionPrefixFilterPath))
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
					fmt.Fprintf(res, "%s", `{"action": "permit", "before": "1a15dcab-7e40-45e1-b7c5-bc690eaa9782", "created_at": "2019-01-01T12:00:00.000Z", "ge": 0, "id": "1a15dcab-7e30-45e1-b7c5-bc690eaa9865", "le": 32, "prefix": "192.168.100.0/24", "updated_at": "2019-01-01T12:00:00.000Z"}`)
				}))
			})
			It(`Invoke UpdateTransitGatewayConnectionPrefixFilter successfully with retries`, func() {
				transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(transitGatewayApisService).ToNot(BeNil())
				transitGatewayApisService.EnableRetries(0, 0)

				// Construct an instance of the UpdateTransitGatewayConnectionPrefixFilterOptions model
				updateTransitGatewayConnectionPrefixFilterOptionsModel := new(transitgatewayapisv1.UpdateTransitGatewayConnectionPrefixFilterOptions)
				updateTransitGatewayConnectionPrefixFilterOptionsModel.TransitGatewayID = core.StringPtr("testString")
				updateTransitGatewayConnectionPrefixFilterOptionsModel.ID = core.StringPtr("testString")
				updateTransitGatewayConnectionPrefixFilterOptionsModel.FilterID = core.StringPtr("testString")
				updateTransitGatewayConnectionPrefixFilterOptionsModel.Action = core.StringPtr("permit")
				updateTransitGatewayConnectionPrefixFilterOptionsModel.Before = core.StringPtr("1a15dcab-7e40-45e1-b7c5-bc690eaa9782")
				updateTransitGatewayConnectionPrefixFilterOptionsModel.Ge = core.Int64Ptr(int64(0))
				updateTransitGatewayConnectionPrefixFilterOptionsModel.Le = core.Int64Ptr(int64(32))
				updateTransitGatewayConnectionPrefixFilterOptionsModel.Prefix = core.StringPtr("192.168.100.0/24")
				updateTransitGatewayConnectionPrefixFilterOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := transitGatewayApisService.UpdateTransitGatewayConnectionPrefixFilterWithContext(ctx, updateTransitGatewayConnectionPrefixFilterOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				transitGatewayApisService.DisableRetries()
				result, response, operationErr := transitGatewayApisService.UpdateTransitGatewayConnectionPrefixFilter(updateTransitGatewayConnectionPrefixFilterOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = transitGatewayApisService.UpdateTransitGatewayConnectionPrefixFilterWithContext(ctx, updateTransitGatewayConnectionPrefixFilterOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(updateTransitGatewayConnectionPrefixFilterPath))
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
					fmt.Fprintf(res, "%s", `{"action": "permit", "before": "1a15dcab-7e40-45e1-b7c5-bc690eaa9782", "created_at": "2019-01-01T12:00:00.000Z", "ge": 0, "id": "1a15dcab-7e30-45e1-b7c5-bc690eaa9865", "le": 32, "prefix": "192.168.100.0/24", "updated_at": "2019-01-01T12:00:00.000Z"}`)
				}))
			})
			It(`Invoke UpdateTransitGatewayConnectionPrefixFilter successfully`, func() {
				transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(transitGatewayApisService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := transitGatewayApisService.UpdateTransitGatewayConnectionPrefixFilter(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateTransitGatewayConnectionPrefixFilterOptions model
				updateTransitGatewayConnectionPrefixFilterOptionsModel := new(transitgatewayapisv1.UpdateTransitGatewayConnectionPrefixFilterOptions)
				updateTransitGatewayConnectionPrefixFilterOptionsModel.TransitGatewayID = core.StringPtr("testString")
				updateTransitGatewayConnectionPrefixFilterOptionsModel.ID = core.StringPtr("testString")
				updateTransitGatewayConnectionPrefixFilterOptionsModel.FilterID = core.StringPtr("testString")
				updateTransitGatewayConnectionPrefixFilterOptionsModel.Action = core.StringPtr("permit")
				updateTransitGatewayConnectionPrefixFilterOptionsModel.Before = core.StringPtr("1a15dcab-7e40-45e1-b7c5-bc690eaa9782")
				updateTransitGatewayConnectionPrefixFilterOptionsModel.Ge = core.Int64Ptr(int64(0))
				updateTransitGatewayConnectionPrefixFilterOptionsModel.Le = core.Int64Ptr(int64(32))
				updateTransitGatewayConnectionPrefixFilterOptionsModel.Prefix = core.StringPtr("192.168.100.0/24")
				updateTransitGatewayConnectionPrefixFilterOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = transitGatewayApisService.UpdateTransitGatewayConnectionPrefixFilter(updateTransitGatewayConnectionPrefixFilterOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateTransitGatewayConnectionPrefixFilter with error: Operation validation and request error`, func() {
				transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(transitGatewayApisService).ToNot(BeNil())

				// Construct an instance of the UpdateTransitGatewayConnectionPrefixFilterOptions model
				updateTransitGatewayConnectionPrefixFilterOptionsModel := new(transitgatewayapisv1.UpdateTransitGatewayConnectionPrefixFilterOptions)
				updateTransitGatewayConnectionPrefixFilterOptionsModel.TransitGatewayID = core.StringPtr("testString")
				updateTransitGatewayConnectionPrefixFilterOptionsModel.ID = core.StringPtr("testString")
				updateTransitGatewayConnectionPrefixFilterOptionsModel.FilterID = core.StringPtr("testString")
				updateTransitGatewayConnectionPrefixFilterOptionsModel.Action = core.StringPtr("permit")
				updateTransitGatewayConnectionPrefixFilterOptionsModel.Before = core.StringPtr("1a15dcab-7e40-45e1-b7c5-bc690eaa9782")
				updateTransitGatewayConnectionPrefixFilterOptionsModel.Ge = core.Int64Ptr(int64(0))
				updateTransitGatewayConnectionPrefixFilterOptionsModel.Le = core.Int64Ptr(int64(32))
				updateTransitGatewayConnectionPrefixFilterOptionsModel.Prefix = core.StringPtr("192.168.100.0/24")
				updateTransitGatewayConnectionPrefixFilterOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := transitGatewayApisService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := transitGatewayApisService.UpdateTransitGatewayConnectionPrefixFilter(updateTransitGatewayConnectionPrefixFilterOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateTransitGatewayConnectionPrefixFilterOptions model with no property values
				updateTransitGatewayConnectionPrefixFilterOptionsModelNew := new(transitgatewayapisv1.UpdateTransitGatewayConnectionPrefixFilterOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = transitGatewayApisService.UpdateTransitGatewayConnectionPrefixFilter(updateTransitGatewayConnectionPrefixFilterOptionsModelNew)
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
			It(`Invoke UpdateTransitGatewayConnectionPrefixFilter successfully`, func() {
				transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(transitGatewayApisService).ToNot(BeNil())

				// Construct an instance of the UpdateTransitGatewayConnectionPrefixFilterOptions model
				updateTransitGatewayConnectionPrefixFilterOptionsModel := new(transitgatewayapisv1.UpdateTransitGatewayConnectionPrefixFilterOptions)
				updateTransitGatewayConnectionPrefixFilterOptionsModel.TransitGatewayID = core.StringPtr("testString")
				updateTransitGatewayConnectionPrefixFilterOptionsModel.ID = core.StringPtr("testString")
				updateTransitGatewayConnectionPrefixFilterOptionsModel.FilterID = core.StringPtr("testString")
				updateTransitGatewayConnectionPrefixFilterOptionsModel.Action = core.StringPtr("permit")
				updateTransitGatewayConnectionPrefixFilterOptionsModel.Before = core.StringPtr("1a15dcab-7e40-45e1-b7c5-bc690eaa9782")
				updateTransitGatewayConnectionPrefixFilterOptionsModel.Ge = core.Int64Ptr(int64(0))
				updateTransitGatewayConnectionPrefixFilterOptionsModel.Le = core.Int64Ptr(int64(32))
				updateTransitGatewayConnectionPrefixFilterOptionsModel.Prefix = core.StringPtr("192.168.100.0/24")
				updateTransitGatewayConnectionPrefixFilterOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := transitGatewayApisService.UpdateTransitGatewayConnectionPrefixFilter(updateTransitGatewayConnectionPrefixFilterOptionsModel)
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
	Describe(`ListTransitGatewayRouteReports(listTransitGatewayRouteReportsOptions *ListTransitGatewayRouteReportsOptions) - Operation response error`, func() {
		version := "testString"
		listTransitGatewayRouteReportsPath := "/transit_gateways/testString/route_reports"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listTransitGatewayRouteReportsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListTransitGatewayRouteReports with error: Operation response processing error`, func() {
				transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(transitGatewayApisService).ToNot(BeNil())

				// Construct an instance of the ListTransitGatewayRouteReportsOptions model
				listTransitGatewayRouteReportsOptionsModel := new(transitgatewayapisv1.ListTransitGatewayRouteReportsOptions)
				listTransitGatewayRouteReportsOptionsModel.TransitGatewayID = core.StringPtr("testString")
				listTransitGatewayRouteReportsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := transitGatewayApisService.ListTransitGatewayRouteReports(listTransitGatewayRouteReportsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				transitGatewayApisService.EnableRetries(0, 0)
				result, response, operationErr = transitGatewayApisService.ListTransitGatewayRouteReports(listTransitGatewayRouteReportsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListTransitGatewayRouteReports(listTransitGatewayRouteReportsOptions *ListTransitGatewayRouteReportsOptions)`, func() {
		version := "testString"
		listTransitGatewayRouteReportsPath := "/transit_gateways/testString/route_reports"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listTransitGatewayRouteReportsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"route_reports": [{"connections": [{"bgps": [{"as_path": "(65201 4201065544) 4203065544", "is_used": true, "local_preference": "190", "prefix": "172.17.0.0/16"}], "id": "3c265a62-91da-4261-a950-950b6af0eb58", "name": "transit-connection-vpc1", "routes": [{"prefix": "192.168.0.0/16"}], "type": "vpc"}], "created_at": "2019-01-01T12:00:00.000Z", "id": "1a15dcab-7e26-45e1-b7c5-bc690eaa9724", "overlapping_routes": [{"routes": [{"connection_id": "d2d985d8-1d8e-4e8b-96cd-cee2290ecaff", "prefix": "Prefix"}]}], "status": "complete", "updated_at": "2019-01-01T12:00:00.000Z"}]}`)
				}))
			})
			It(`Invoke ListTransitGatewayRouteReports successfully with retries`, func() {
				transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(transitGatewayApisService).ToNot(BeNil())
				transitGatewayApisService.EnableRetries(0, 0)

				// Construct an instance of the ListTransitGatewayRouteReportsOptions model
				listTransitGatewayRouteReportsOptionsModel := new(transitgatewayapisv1.ListTransitGatewayRouteReportsOptions)
				listTransitGatewayRouteReportsOptionsModel.TransitGatewayID = core.StringPtr("testString")
				listTransitGatewayRouteReportsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := transitGatewayApisService.ListTransitGatewayRouteReportsWithContext(ctx, listTransitGatewayRouteReportsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				transitGatewayApisService.DisableRetries()
				result, response, operationErr := transitGatewayApisService.ListTransitGatewayRouteReports(listTransitGatewayRouteReportsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = transitGatewayApisService.ListTransitGatewayRouteReportsWithContext(ctx, listTransitGatewayRouteReportsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listTransitGatewayRouteReportsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"route_reports": [{"connections": [{"bgps": [{"as_path": "(65201 4201065544) 4203065544", "is_used": true, "local_preference": "190", "prefix": "172.17.0.0/16"}], "id": "3c265a62-91da-4261-a950-950b6af0eb58", "name": "transit-connection-vpc1", "routes": [{"prefix": "192.168.0.0/16"}], "type": "vpc"}], "created_at": "2019-01-01T12:00:00.000Z", "id": "1a15dcab-7e26-45e1-b7c5-bc690eaa9724", "overlapping_routes": [{"routes": [{"connection_id": "d2d985d8-1d8e-4e8b-96cd-cee2290ecaff", "prefix": "Prefix"}]}], "status": "complete", "updated_at": "2019-01-01T12:00:00.000Z"}]}`)
				}))
			})
			It(`Invoke ListTransitGatewayRouteReports successfully`, func() {
				transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(transitGatewayApisService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := transitGatewayApisService.ListTransitGatewayRouteReports(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListTransitGatewayRouteReportsOptions model
				listTransitGatewayRouteReportsOptionsModel := new(transitgatewayapisv1.ListTransitGatewayRouteReportsOptions)
				listTransitGatewayRouteReportsOptionsModel.TransitGatewayID = core.StringPtr("testString")
				listTransitGatewayRouteReportsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = transitGatewayApisService.ListTransitGatewayRouteReports(listTransitGatewayRouteReportsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListTransitGatewayRouteReports with error: Operation validation and request error`, func() {
				transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(transitGatewayApisService).ToNot(BeNil())

				// Construct an instance of the ListTransitGatewayRouteReportsOptions model
				listTransitGatewayRouteReportsOptionsModel := new(transitgatewayapisv1.ListTransitGatewayRouteReportsOptions)
				listTransitGatewayRouteReportsOptionsModel.TransitGatewayID = core.StringPtr("testString")
				listTransitGatewayRouteReportsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := transitGatewayApisService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := transitGatewayApisService.ListTransitGatewayRouteReports(listTransitGatewayRouteReportsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListTransitGatewayRouteReportsOptions model with no property values
				listTransitGatewayRouteReportsOptionsModelNew := new(transitgatewayapisv1.ListTransitGatewayRouteReportsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = transitGatewayApisService.ListTransitGatewayRouteReports(listTransitGatewayRouteReportsOptionsModelNew)
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
			It(`Invoke ListTransitGatewayRouteReports successfully`, func() {
				transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(transitGatewayApisService).ToNot(BeNil())

				// Construct an instance of the ListTransitGatewayRouteReportsOptions model
				listTransitGatewayRouteReportsOptionsModel := new(transitgatewayapisv1.ListTransitGatewayRouteReportsOptions)
				listTransitGatewayRouteReportsOptionsModel.TransitGatewayID = core.StringPtr("testString")
				listTransitGatewayRouteReportsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := transitGatewayApisService.ListTransitGatewayRouteReports(listTransitGatewayRouteReportsOptionsModel)
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
	Describe(`CreateTransitGatewayRouteReport(createTransitGatewayRouteReportOptions *CreateTransitGatewayRouteReportOptions) - Operation response error`, func() {
		version := "testString"
		createTransitGatewayRouteReportPath := "/transit_gateways/testString/route_reports"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createTransitGatewayRouteReportPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateTransitGatewayRouteReport with error: Operation response processing error`, func() {
				transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(transitGatewayApisService).ToNot(BeNil())

				// Construct an instance of the CreateTransitGatewayRouteReportOptions model
				createTransitGatewayRouteReportOptionsModel := new(transitgatewayapisv1.CreateTransitGatewayRouteReportOptions)
				createTransitGatewayRouteReportOptionsModel.TransitGatewayID = core.StringPtr("testString")
				createTransitGatewayRouteReportOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := transitGatewayApisService.CreateTransitGatewayRouteReport(createTransitGatewayRouteReportOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				transitGatewayApisService.EnableRetries(0, 0)
				result, response, operationErr = transitGatewayApisService.CreateTransitGatewayRouteReport(createTransitGatewayRouteReportOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateTransitGatewayRouteReport(createTransitGatewayRouteReportOptions *CreateTransitGatewayRouteReportOptions)`, func() {
		version := "testString"
		createTransitGatewayRouteReportPath := "/transit_gateways/testString/route_reports"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createTransitGatewayRouteReportPath))
					Expect(req.Method).To(Equal("POST"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"connections": [{"bgps": [{"as_path": "(65201 4201065544) 4203065544", "is_used": true, "local_preference": "190", "prefix": "172.17.0.0/16"}], "id": "3c265a62-91da-4261-a950-950b6af0eb58", "name": "transit-connection-vpc1", "routes": [{"prefix": "192.168.0.0/16"}], "type": "vpc"}], "created_at": "2019-01-01T12:00:00.000Z", "id": "1a15dcab-7e26-45e1-b7c5-bc690eaa9724", "overlapping_routes": [{"routes": [{"connection_id": "d2d985d8-1d8e-4e8b-96cd-cee2290ecaff", "prefix": "Prefix"}]}], "status": "complete", "updated_at": "2019-01-01T12:00:00.000Z"}`)
				}))
			})
			It(`Invoke CreateTransitGatewayRouteReport successfully with retries`, func() {
				transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(transitGatewayApisService).ToNot(BeNil())
				transitGatewayApisService.EnableRetries(0, 0)

				// Construct an instance of the CreateTransitGatewayRouteReportOptions model
				createTransitGatewayRouteReportOptionsModel := new(transitgatewayapisv1.CreateTransitGatewayRouteReportOptions)
				createTransitGatewayRouteReportOptionsModel.TransitGatewayID = core.StringPtr("testString")
				createTransitGatewayRouteReportOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := transitGatewayApisService.CreateTransitGatewayRouteReportWithContext(ctx, createTransitGatewayRouteReportOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				transitGatewayApisService.DisableRetries()
				result, response, operationErr := transitGatewayApisService.CreateTransitGatewayRouteReport(createTransitGatewayRouteReportOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = transitGatewayApisService.CreateTransitGatewayRouteReportWithContext(ctx, createTransitGatewayRouteReportOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createTransitGatewayRouteReportPath))
					Expect(req.Method).To(Equal("POST"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"connections": [{"bgps": [{"as_path": "(65201 4201065544) 4203065544", "is_used": true, "local_preference": "190", "prefix": "172.17.0.0/16"}], "id": "3c265a62-91da-4261-a950-950b6af0eb58", "name": "transit-connection-vpc1", "routes": [{"prefix": "192.168.0.0/16"}], "type": "vpc"}], "created_at": "2019-01-01T12:00:00.000Z", "id": "1a15dcab-7e26-45e1-b7c5-bc690eaa9724", "overlapping_routes": [{"routes": [{"connection_id": "d2d985d8-1d8e-4e8b-96cd-cee2290ecaff", "prefix": "Prefix"}]}], "status": "complete", "updated_at": "2019-01-01T12:00:00.000Z"}`)
				}))
			})
			It(`Invoke CreateTransitGatewayRouteReport successfully`, func() {
				transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(transitGatewayApisService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := transitGatewayApisService.CreateTransitGatewayRouteReport(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateTransitGatewayRouteReportOptions model
				createTransitGatewayRouteReportOptionsModel := new(transitgatewayapisv1.CreateTransitGatewayRouteReportOptions)
				createTransitGatewayRouteReportOptionsModel.TransitGatewayID = core.StringPtr("testString")
				createTransitGatewayRouteReportOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = transitGatewayApisService.CreateTransitGatewayRouteReport(createTransitGatewayRouteReportOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateTransitGatewayRouteReport with error: Operation validation and request error`, func() {
				transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(transitGatewayApisService).ToNot(BeNil())

				// Construct an instance of the CreateTransitGatewayRouteReportOptions model
				createTransitGatewayRouteReportOptionsModel := new(transitgatewayapisv1.CreateTransitGatewayRouteReportOptions)
				createTransitGatewayRouteReportOptionsModel.TransitGatewayID = core.StringPtr("testString")
				createTransitGatewayRouteReportOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := transitGatewayApisService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := transitGatewayApisService.CreateTransitGatewayRouteReport(createTransitGatewayRouteReportOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateTransitGatewayRouteReportOptions model with no property values
				createTransitGatewayRouteReportOptionsModelNew := new(transitgatewayapisv1.CreateTransitGatewayRouteReportOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = transitGatewayApisService.CreateTransitGatewayRouteReport(createTransitGatewayRouteReportOptionsModelNew)
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
			It(`Invoke CreateTransitGatewayRouteReport successfully`, func() {
				transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(transitGatewayApisService).ToNot(BeNil())

				// Construct an instance of the CreateTransitGatewayRouteReportOptions model
				createTransitGatewayRouteReportOptionsModel := new(transitgatewayapisv1.CreateTransitGatewayRouteReportOptions)
				createTransitGatewayRouteReportOptionsModel.TransitGatewayID = core.StringPtr("testString")
				createTransitGatewayRouteReportOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := transitGatewayApisService.CreateTransitGatewayRouteReport(createTransitGatewayRouteReportOptionsModel)
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
	Describe(`DeleteTransitGatewayRouteReport(deleteTransitGatewayRouteReportOptions *DeleteTransitGatewayRouteReportOptions)`, func() {
		version := "testString"
		deleteTransitGatewayRouteReportPath := "/transit_gateways/testString/route_reports/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteTransitGatewayRouteReportPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteTransitGatewayRouteReport successfully`, func() {
				transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(transitGatewayApisService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := transitGatewayApisService.DeleteTransitGatewayRouteReport(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteTransitGatewayRouteReportOptions model
				deleteTransitGatewayRouteReportOptionsModel := new(transitgatewayapisv1.DeleteTransitGatewayRouteReportOptions)
				deleteTransitGatewayRouteReportOptionsModel.TransitGatewayID = core.StringPtr("testString")
				deleteTransitGatewayRouteReportOptionsModel.ID = core.StringPtr("testString")
				deleteTransitGatewayRouteReportOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = transitGatewayApisService.DeleteTransitGatewayRouteReport(deleteTransitGatewayRouteReportOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteTransitGatewayRouteReport with error: Operation validation and request error`, func() {
				transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(transitGatewayApisService).ToNot(BeNil())

				// Construct an instance of the DeleteTransitGatewayRouteReportOptions model
				deleteTransitGatewayRouteReportOptionsModel := new(transitgatewayapisv1.DeleteTransitGatewayRouteReportOptions)
				deleteTransitGatewayRouteReportOptionsModel.TransitGatewayID = core.StringPtr("testString")
				deleteTransitGatewayRouteReportOptionsModel.ID = core.StringPtr("testString")
				deleteTransitGatewayRouteReportOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := transitGatewayApisService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := transitGatewayApisService.DeleteTransitGatewayRouteReport(deleteTransitGatewayRouteReportOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteTransitGatewayRouteReportOptions model with no property values
				deleteTransitGatewayRouteReportOptionsModelNew := new(transitgatewayapisv1.DeleteTransitGatewayRouteReportOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = transitGatewayApisService.DeleteTransitGatewayRouteReport(deleteTransitGatewayRouteReportOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetTransitGatewayRouteReport(getTransitGatewayRouteReportOptions *GetTransitGatewayRouteReportOptions) - Operation response error`, func() {
		version := "testString"
		getTransitGatewayRouteReportPath := "/transit_gateways/testString/route_reports/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getTransitGatewayRouteReportPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetTransitGatewayRouteReport with error: Operation response processing error`, func() {
				transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(transitGatewayApisService).ToNot(BeNil())

				// Construct an instance of the GetTransitGatewayRouteReportOptions model
				getTransitGatewayRouteReportOptionsModel := new(transitgatewayapisv1.GetTransitGatewayRouteReportOptions)
				getTransitGatewayRouteReportOptionsModel.TransitGatewayID = core.StringPtr("testString")
				getTransitGatewayRouteReportOptionsModel.ID = core.StringPtr("testString")
				getTransitGatewayRouteReportOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := transitGatewayApisService.GetTransitGatewayRouteReport(getTransitGatewayRouteReportOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				transitGatewayApisService.EnableRetries(0, 0)
				result, response, operationErr = transitGatewayApisService.GetTransitGatewayRouteReport(getTransitGatewayRouteReportOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetTransitGatewayRouteReport(getTransitGatewayRouteReportOptions *GetTransitGatewayRouteReportOptions)`, func() {
		version := "testString"
		getTransitGatewayRouteReportPath := "/transit_gateways/testString/route_reports/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getTransitGatewayRouteReportPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"connections": [{"bgps": [{"as_path": "(65201 4201065544) 4203065544", "is_used": true, "local_preference": "190", "prefix": "172.17.0.0/16"}], "id": "3c265a62-91da-4261-a950-950b6af0eb58", "name": "transit-connection-vpc1", "routes": [{"prefix": "192.168.0.0/16"}], "type": "vpc"}], "created_at": "2019-01-01T12:00:00.000Z", "id": "1a15dcab-7e26-45e1-b7c5-bc690eaa9724", "overlapping_routes": [{"routes": [{"connection_id": "d2d985d8-1d8e-4e8b-96cd-cee2290ecaff", "prefix": "Prefix"}]}], "status": "complete", "updated_at": "2019-01-01T12:00:00.000Z"}`)
				}))
			})
			It(`Invoke GetTransitGatewayRouteReport successfully with retries`, func() {
				transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(transitGatewayApisService).ToNot(BeNil())
				transitGatewayApisService.EnableRetries(0, 0)

				// Construct an instance of the GetTransitGatewayRouteReportOptions model
				getTransitGatewayRouteReportOptionsModel := new(transitgatewayapisv1.GetTransitGatewayRouteReportOptions)
				getTransitGatewayRouteReportOptionsModel.TransitGatewayID = core.StringPtr("testString")
				getTransitGatewayRouteReportOptionsModel.ID = core.StringPtr("testString")
				getTransitGatewayRouteReportOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := transitGatewayApisService.GetTransitGatewayRouteReportWithContext(ctx, getTransitGatewayRouteReportOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				transitGatewayApisService.DisableRetries()
				result, response, operationErr := transitGatewayApisService.GetTransitGatewayRouteReport(getTransitGatewayRouteReportOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = transitGatewayApisService.GetTransitGatewayRouteReportWithContext(ctx, getTransitGatewayRouteReportOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getTransitGatewayRouteReportPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"connections": [{"bgps": [{"as_path": "(65201 4201065544) 4203065544", "is_used": true, "local_preference": "190", "prefix": "172.17.0.0/16"}], "id": "3c265a62-91da-4261-a950-950b6af0eb58", "name": "transit-connection-vpc1", "routes": [{"prefix": "192.168.0.0/16"}], "type": "vpc"}], "created_at": "2019-01-01T12:00:00.000Z", "id": "1a15dcab-7e26-45e1-b7c5-bc690eaa9724", "overlapping_routes": [{"routes": [{"connection_id": "d2d985d8-1d8e-4e8b-96cd-cee2290ecaff", "prefix": "Prefix"}]}], "status": "complete", "updated_at": "2019-01-01T12:00:00.000Z"}`)
				}))
			})
			It(`Invoke GetTransitGatewayRouteReport successfully`, func() {
				transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(transitGatewayApisService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := transitGatewayApisService.GetTransitGatewayRouteReport(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetTransitGatewayRouteReportOptions model
				getTransitGatewayRouteReportOptionsModel := new(transitgatewayapisv1.GetTransitGatewayRouteReportOptions)
				getTransitGatewayRouteReportOptionsModel.TransitGatewayID = core.StringPtr("testString")
				getTransitGatewayRouteReportOptionsModel.ID = core.StringPtr("testString")
				getTransitGatewayRouteReportOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = transitGatewayApisService.GetTransitGatewayRouteReport(getTransitGatewayRouteReportOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetTransitGatewayRouteReport with error: Operation validation and request error`, func() {
				transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(transitGatewayApisService).ToNot(BeNil())

				// Construct an instance of the GetTransitGatewayRouteReportOptions model
				getTransitGatewayRouteReportOptionsModel := new(transitgatewayapisv1.GetTransitGatewayRouteReportOptions)
				getTransitGatewayRouteReportOptionsModel.TransitGatewayID = core.StringPtr("testString")
				getTransitGatewayRouteReportOptionsModel.ID = core.StringPtr("testString")
				getTransitGatewayRouteReportOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := transitGatewayApisService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := transitGatewayApisService.GetTransitGatewayRouteReport(getTransitGatewayRouteReportOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetTransitGatewayRouteReportOptions model with no property values
				getTransitGatewayRouteReportOptionsModelNew := new(transitgatewayapisv1.GetTransitGatewayRouteReportOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = transitGatewayApisService.GetTransitGatewayRouteReport(getTransitGatewayRouteReportOptionsModelNew)
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
			It(`Invoke GetTransitGatewayRouteReport successfully`, func() {
				transitGatewayApisService, serviceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(transitGatewayApisService).ToNot(BeNil())

				// Construct an instance of the GetTransitGatewayRouteReportOptions model
				getTransitGatewayRouteReportOptionsModel := new(transitgatewayapisv1.GetTransitGatewayRouteReportOptions)
				getTransitGatewayRouteReportOptionsModel.TransitGatewayID = core.StringPtr("testString")
				getTransitGatewayRouteReportOptionsModel.ID = core.StringPtr("testString")
				getTransitGatewayRouteReportOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := transitGatewayApisService.GetTransitGatewayRouteReport(getTransitGatewayRouteReportOptionsModel)
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
				// Construct an instance of the TransitGatewayConnectionPrefixFilter model
				transitGatewayConnectionPrefixFilterModel := new(transitgatewayapisv1.TransitGatewayConnectionPrefixFilter)
				Expect(transitGatewayConnectionPrefixFilterModel).ToNot(BeNil())
				transitGatewayConnectionPrefixFilterModel.Action = core.StringPtr("permit")
				transitGatewayConnectionPrefixFilterModel.Ge = core.Int64Ptr(int64(0))
				transitGatewayConnectionPrefixFilterModel.Le = core.Int64Ptr(int64(32))
				transitGatewayConnectionPrefixFilterModel.Prefix = core.StringPtr("192.168.100.0/24")
				Expect(transitGatewayConnectionPrefixFilterModel.Action).To(Equal(core.StringPtr("permit")))
				Expect(transitGatewayConnectionPrefixFilterModel.Ge).To(Equal(core.Int64Ptr(int64(0))))
				Expect(transitGatewayConnectionPrefixFilterModel.Le).To(Equal(core.Int64Ptr(int64(32))))
				Expect(transitGatewayConnectionPrefixFilterModel.Prefix).To(Equal(core.StringPtr("192.168.100.0/24")))

				// Construct an instance of the ZoneIdentityByName model
				zoneIdentityModel := new(transitgatewayapisv1.ZoneIdentityByName)
				Expect(zoneIdentityModel).ToNot(BeNil())
				zoneIdentityModel.Name = core.StringPtr("us-south-1")
				Expect(zoneIdentityModel.Name).To(Equal(core.StringPtr("us-south-1")))

				// Construct an instance of the TransitGatewayTunnelTemplate model
				transitGatewayTunnelTemplateModel := new(transitgatewayapisv1.TransitGatewayTunnelTemplate)
				Expect(transitGatewayTunnelTemplateModel).ToNot(BeNil())
				transitGatewayTunnelTemplateModel.LocalGatewayIp = core.StringPtr("10.242.63.12")
				transitGatewayTunnelTemplateModel.LocalTunnelIp = core.StringPtr("192.168.100.20")
				transitGatewayTunnelTemplateModel.Name = core.StringPtr("gre1")
				transitGatewayTunnelTemplateModel.RemoteBgpAsn = core.Int64Ptr(int64(65010))
				transitGatewayTunnelTemplateModel.RemoteGatewayIp = core.StringPtr("10.242.33.22")
				transitGatewayTunnelTemplateModel.RemoteTunnelIp = core.StringPtr("192.168.129.1")
				transitGatewayTunnelTemplateModel.Zone = zoneIdentityModel
				Expect(transitGatewayTunnelTemplateModel.LocalGatewayIp).To(Equal(core.StringPtr("10.242.63.12")))
				Expect(transitGatewayTunnelTemplateModel.LocalTunnelIp).To(Equal(core.StringPtr("192.168.100.20")))
				Expect(transitGatewayTunnelTemplateModel.Name).To(Equal(core.StringPtr("gre1")))
				Expect(transitGatewayTunnelTemplateModel.RemoteBgpAsn).To(Equal(core.Int64Ptr(int64(65010))))
				Expect(transitGatewayTunnelTemplateModel.RemoteGatewayIp).To(Equal(core.StringPtr("10.242.33.22")))
				Expect(transitGatewayTunnelTemplateModel.RemoteTunnelIp).To(Equal(core.StringPtr("192.168.129.1")))
				Expect(transitGatewayTunnelTemplateModel.Zone).To(Equal(zoneIdentityModel))

				// Construct an instance of the CreateTransitGatewayConnectionOptions model
				transitGatewayID := "testString"
				createTransitGatewayConnectionOptionsNetworkType := "vpc"
				createTransitGatewayConnectionOptionsModel := transitGatewayApisService.NewCreateTransitGatewayConnectionOptions(transitGatewayID, createTransitGatewayConnectionOptionsNetworkType)
				createTransitGatewayConnectionOptionsModel.SetTransitGatewayID("testString")
				createTransitGatewayConnectionOptionsModel.SetNetworkType("vpc")
				createTransitGatewayConnectionOptionsModel.SetBaseConnectionID("975f58c1-afe7-469a-9727-7f3d720f2d32")
				createTransitGatewayConnectionOptionsModel.SetBaseNetworkType("classic")
				createTransitGatewayConnectionOptionsModel.SetCidr("192.168.0.0/24")
				createTransitGatewayConnectionOptionsModel.SetLocalGatewayIp("192.168.100.1")
				createTransitGatewayConnectionOptionsModel.SetLocalTunnelIp("192.168.129.2")
				createTransitGatewayConnectionOptionsModel.SetName("Transit_Service_BWTN_SJ_DL")
				createTransitGatewayConnectionOptionsModel.SetNetworkAccountID("testString")
				createTransitGatewayConnectionOptionsModel.SetNetworkID("crn:v1:bluemix:public:is:us-south:a/123456::vpc:4727d842-f94f-4a2d-824a-9bc9b02c523b")
				createTransitGatewayConnectionOptionsModel.SetPrefixFilters([]transitgatewayapisv1.TransitGatewayConnectionPrefixFilter{*transitGatewayConnectionPrefixFilterModel})
				createTransitGatewayConnectionOptionsModel.SetPrefixFiltersDefault("permit")
				createTransitGatewayConnectionOptionsModel.SetRemoteBgpAsn(int64(65010))
				createTransitGatewayConnectionOptionsModel.SetRemoteGatewayIp("10.242.63.12")
				createTransitGatewayConnectionOptionsModel.SetRemoteTunnelIp("192.168.129.1")
				createTransitGatewayConnectionOptionsModel.SetTunnels([]transitgatewayapisv1.TransitGatewayTunnelTemplate{*transitGatewayTunnelTemplateModel})
				createTransitGatewayConnectionOptionsModel.SetZone(zoneIdentityModel)
				createTransitGatewayConnectionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createTransitGatewayConnectionOptionsModel).ToNot(BeNil())
				Expect(createTransitGatewayConnectionOptionsModel.TransitGatewayID).To(Equal(core.StringPtr("testString")))
				Expect(createTransitGatewayConnectionOptionsModel.NetworkType).To(Equal(core.StringPtr("vpc")))
				Expect(createTransitGatewayConnectionOptionsModel.BaseConnectionID).To(Equal(core.StringPtr("975f58c1-afe7-469a-9727-7f3d720f2d32")))
				Expect(createTransitGatewayConnectionOptionsModel.BaseNetworkType).To(Equal(core.StringPtr("classic")))
				Expect(createTransitGatewayConnectionOptionsModel.Cidr).To(Equal(core.StringPtr("192.168.0.0/24")))
				Expect(createTransitGatewayConnectionOptionsModel.LocalGatewayIp).To(Equal(core.StringPtr("192.168.100.1")))
				Expect(createTransitGatewayConnectionOptionsModel.LocalTunnelIp).To(Equal(core.StringPtr("192.168.129.2")))
				Expect(createTransitGatewayConnectionOptionsModel.Name).To(Equal(core.StringPtr("Transit_Service_BWTN_SJ_DL")))
				Expect(createTransitGatewayConnectionOptionsModel.NetworkAccountID).To(Equal(core.StringPtr("testString")))
				Expect(createTransitGatewayConnectionOptionsModel.NetworkID).To(Equal(core.StringPtr("crn:v1:bluemix:public:is:us-south:a/123456::vpc:4727d842-f94f-4a2d-824a-9bc9b02c523b")))
				Expect(createTransitGatewayConnectionOptionsModel.PrefixFilters).To(Equal([]transitgatewayapisv1.TransitGatewayConnectionPrefixFilter{*transitGatewayConnectionPrefixFilterModel}))
				Expect(createTransitGatewayConnectionOptionsModel.PrefixFiltersDefault).To(Equal(core.StringPtr("permit")))
				Expect(createTransitGatewayConnectionOptionsModel.RemoteBgpAsn).To(Equal(core.Int64Ptr(int64(65010))))
				Expect(createTransitGatewayConnectionOptionsModel.RemoteGatewayIp).To(Equal(core.StringPtr("10.242.63.12")))
				Expect(createTransitGatewayConnectionOptionsModel.RemoteTunnelIp).To(Equal(core.StringPtr("192.168.129.1")))
				Expect(createTransitGatewayConnectionOptionsModel.Tunnels).To(Equal([]transitgatewayapisv1.TransitGatewayTunnelTemplate{*transitGatewayTunnelTemplateModel}))
				Expect(createTransitGatewayConnectionOptionsModel.Zone).To(Equal(zoneIdentityModel))
				Expect(createTransitGatewayConnectionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateTransitGatewayConnectionPrefixFilterOptions successfully`, func() {
				// Construct an instance of the CreateTransitGatewayConnectionPrefixFilterOptions model
				transitGatewayID := "testString"
				id := "testString"
				createTransitGatewayConnectionPrefixFilterOptionsAction := "permit"
				createTransitGatewayConnectionPrefixFilterOptionsPrefix := "192.168.100.0/24"
				createTransitGatewayConnectionPrefixFilterOptionsModel := transitGatewayApisService.NewCreateTransitGatewayConnectionPrefixFilterOptions(transitGatewayID, id, createTransitGatewayConnectionPrefixFilterOptionsAction, createTransitGatewayConnectionPrefixFilterOptionsPrefix)
				createTransitGatewayConnectionPrefixFilterOptionsModel.SetTransitGatewayID("testString")
				createTransitGatewayConnectionPrefixFilterOptionsModel.SetID("testString")
				createTransitGatewayConnectionPrefixFilterOptionsModel.SetAction("permit")
				createTransitGatewayConnectionPrefixFilterOptionsModel.SetPrefix("192.168.100.0/24")
				createTransitGatewayConnectionPrefixFilterOptionsModel.SetBefore("1a15dcab-7e40-45e1-b7c5-bc690eaa9782")
				createTransitGatewayConnectionPrefixFilterOptionsModel.SetGe(int64(0))
				createTransitGatewayConnectionPrefixFilterOptionsModel.SetLe(int64(32))
				createTransitGatewayConnectionPrefixFilterOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createTransitGatewayConnectionPrefixFilterOptionsModel).ToNot(BeNil())
				Expect(createTransitGatewayConnectionPrefixFilterOptionsModel.TransitGatewayID).To(Equal(core.StringPtr("testString")))
				Expect(createTransitGatewayConnectionPrefixFilterOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(createTransitGatewayConnectionPrefixFilterOptionsModel.Action).To(Equal(core.StringPtr("permit")))
				Expect(createTransitGatewayConnectionPrefixFilterOptionsModel.Prefix).To(Equal(core.StringPtr("192.168.100.0/24")))
				Expect(createTransitGatewayConnectionPrefixFilterOptionsModel.Before).To(Equal(core.StringPtr("1a15dcab-7e40-45e1-b7c5-bc690eaa9782")))
				Expect(createTransitGatewayConnectionPrefixFilterOptionsModel.Ge).To(Equal(core.Int64Ptr(int64(0))))
				Expect(createTransitGatewayConnectionPrefixFilterOptionsModel.Le).To(Equal(core.Int64Ptr(int64(32))))
				Expect(createTransitGatewayConnectionPrefixFilterOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateTransitGatewayGreTunnelOptions successfully`, func() {
				// Construct an instance of the ZoneIdentityByName model
				zoneIdentityModel := new(transitgatewayapisv1.ZoneIdentityByName)
				Expect(zoneIdentityModel).ToNot(BeNil())
				zoneIdentityModel.Name = core.StringPtr("us-south-1")
				Expect(zoneIdentityModel.Name).To(Equal(core.StringPtr("us-south-1")))

				// Construct an instance of the CreateTransitGatewayGreTunnelOptions model
				transitGatewayID := "testString"
				id := "testString"
				createTransitGatewayGreTunnelOptionsLocalGatewayIp := "10.242.63.12"
				createTransitGatewayGreTunnelOptionsLocalTunnelIp := "192.168.100.20"
				createTransitGatewayGreTunnelOptionsName := "gre1"
				createTransitGatewayGreTunnelOptionsRemoteGatewayIp := "10.242.33.22"
				createTransitGatewayGreTunnelOptionsRemoteTunnelIp := "192.168.129.1"
				var createTransitGatewayGreTunnelOptionsZone transitgatewayapisv1.ZoneIdentityIntf = nil
				createTransitGatewayGreTunnelOptionsModel := transitGatewayApisService.NewCreateTransitGatewayGreTunnelOptions(transitGatewayID, id, createTransitGatewayGreTunnelOptionsLocalGatewayIp, createTransitGatewayGreTunnelOptionsLocalTunnelIp, createTransitGatewayGreTunnelOptionsName, createTransitGatewayGreTunnelOptionsRemoteGatewayIp, createTransitGatewayGreTunnelOptionsRemoteTunnelIp, createTransitGatewayGreTunnelOptionsZone)
				createTransitGatewayGreTunnelOptionsModel.SetTransitGatewayID("testString")
				createTransitGatewayGreTunnelOptionsModel.SetID("testString")
				createTransitGatewayGreTunnelOptionsModel.SetLocalGatewayIp("10.242.63.12")
				createTransitGatewayGreTunnelOptionsModel.SetLocalTunnelIp("192.168.100.20")
				createTransitGatewayGreTunnelOptionsModel.SetName("gre1")
				createTransitGatewayGreTunnelOptionsModel.SetRemoteGatewayIp("10.242.33.22")
				createTransitGatewayGreTunnelOptionsModel.SetRemoteTunnelIp("192.168.129.1")
				createTransitGatewayGreTunnelOptionsModel.SetZone(zoneIdentityModel)
				createTransitGatewayGreTunnelOptionsModel.SetRemoteBgpAsn(int64(65010))
				createTransitGatewayGreTunnelOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createTransitGatewayGreTunnelOptionsModel).ToNot(BeNil())
				Expect(createTransitGatewayGreTunnelOptionsModel.TransitGatewayID).To(Equal(core.StringPtr("testString")))
				Expect(createTransitGatewayGreTunnelOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(createTransitGatewayGreTunnelOptionsModel.LocalGatewayIp).To(Equal(core.StringPtr("10.242.63.12")))
				Expect(createTransitGatewayGreTunnelOptionsModel.LocalTunnelIp).To(Equal(core.StringPtr("192.168.100.20")))
				Expect(createTransitGatewayGreTunnelOptionsModel.Name).To(Equal(core.StringPtr("gre1")))
				Expect(createTransitGatewayGreTunnelOptionsModel.RemoteGatewayIp).To(Equal(core.StringPtr("10.242.33.22")))
				Expect(createTransitGatewayGreTunnelOptionsModel.RemoteTunnelIp).To(Equal(core.StringPtr("192.168.129.1")))
				Expect(createTransitGatewayGreTunnelOptionsModel.Zone).To(Equal(zoneIdentityModel))
				Expect(createTransitGatewayGreTunnelOptionsModel.RemoteBgpAsn).To(Equal(core.Int64Ptr(int64(65010))))
				Expect(createTransitGatewayGreTunnelOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateTransitGatewayOptions successfully`, func() {
				// Construct an instance of the ResourceGroupIdentity model
				resourceGroupIdentityModel := new(transitgatewayapisv1.ResourceGroupIdentity)
				Expect(resourceGroupIdentityModel).ToNot(BeNil())
				resourceGroupIdentityModel.ID = core.StringPtr("56969d6043e9465c883cb9f7363e78e8")
				Expect(resourceGroupIdentityModel.ID).To(Equal(core.StringPtr("56969d6043e9465c883cb9f7363e78e8")))

				// Construct an instance of the CreateTransitGatewayOptions model
				createTransitGatewayOptionsLocation := "us-south"
				createTransitGatewayOptionsName := "my-transit-gateway-in-TransitGateway"
				createTransitGatewayOptionsModel := transitGatewayApisService.NewCreateTransitGatewayOptions(createTransitGatewayOptionsLocation, createTransitGatewayOptionsName)
				createTransitGatewayOptionsModel.SetLocation("us-south")
				createTransitGatewayOptionsModel.SetName("my-transit-gateway-in-TransitGateway")
				createTransitGatewayOptionsModel.SetGlobal(true)
				createTransitGatewayOptionsModel.SetGreEnhancedRoutePropagation(true)
				createTransitGatewayOptionsModel.SetResourceGroup(resourceGroupIdentityModel)
				createTransitGatewayOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createTransitGatewayOptionsModel).ToNot(BeNil())
				Expect(createTransitGatewayOptionsModel.Location).To(Equal(core.StringPtr("us-south")))
				Expect(createTransitGatewayOptionsModel.Name).To(Equal(core.StringPtr("my-transit-gateway-in-TransitGateway")))
				Expect(createTransitGatewayOptionsModel.Global).To(Equal(core.BoolPtr(true)))
				Expect(createTransitGatewayOptionsModel.GreEnhancedRoutePropagation).To(Equal(core.BoolPtr(true)))
				Expect(createTransitGatewayOptionsModel.ResourceGroup).To(Equal(resourceGroupIdentityModel))
				Expect(createTransitGatewayOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateTransitGatewayRouteReportOptions successfully`, func() {
				// Construct an instance of the CreateTransitGatewayRouteReportOptions model
				transitGatewayID := "testString"
				createTransitGatewayRouteReportOptionsModel := transitGatewayApisService.NewCreateTransitGatewayRouteReportOptions(transitGatewayID)
				createTransitGatewayRouteReportOptionsModel.SetTransitGatewayID("testString")
				createTransitGatewayRouteReportOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createTransitGatewayRouteReportOptionsModel).ToNot(BeNil())
				Expect(createTransitGatewayRouteReportOptionsModel.TransitGatewayID).To(Equal(core.StringPtr("testString")))
				Expect(createTransitGatewayRouteReportOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
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
			It(`Invoke NewDeleteTransitGatewayConnectionPrefixFilterOptions successfully`, func() {
				// Construct an instance of the DeleteTransitGatewayConnectionPrefixFilterOptions model
				transitGatewayID := "testString"
				id := "testString"
				filterID := "testString"
				deleteTransitGatewayConnectionPrefixFilterOptionsModel := transitGatewayApisService.NewDeleteTransitGatewayConnectionPrefixFilterOptions(transitGatewayID, id, filterID)
				deleteTransitGatewayConnectionPrefixFilterOptionsModel.SetTransitGatewayID("testString")
				deleteTransitGatewayConnectionPrefixFilterOptionsModel.SetID("testString")
				deleteTransitGatewayConnectionPrefixFilterOptionsModel.SetFilterID("testString")
				deleteTransitGatewayConnectionPrefixFilterOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteTransitGatewayConnectionPrefixFilterOptionsModel).ToNot(BeNil())
				Expect(deleteTransitGatewayConnectionPrefixFilterOptionsModel.TransitGatewayID).To(Equal(core.StringPtr("testString")))
				Expect(deleteTransitGatewayConnectionPrefixFilterOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(deleteTransitGatewayConnectionPrefixFilterOptionsModel.FilterID).To(Equal(core.StringPtr("testString")))
				Expect(deleteTransitGatewayConnectionPrefixFilterOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteTransitGatewayConnectionTunnelsOptions successfully`, func() {
				// Construct an instance of the DeleteTransitGatewayConnectionTunnelsOptions model
				transitGatewayID := "testString"
				id := "testString"
				greTunnelID := "testString"
				deleteTransitGatewayConnectionTunnelsOptionsModel := transitGatewayApisService.NewDeleteTransitGatewayConnectionTunnelsOptions(transitGatewayID, id, greTunnelID)
				deleteTransitGatewayConnectionTunnelsOptionsModel.SetTransitGatewayID("testString")
				deleteTransitGatewayConnectionTunnelsOptionsModel.SetID("testString")
				deleteTransitGatewayConnectionTunnelsOptionsModel.SetGreTunnelID("testString")
				deleteTransitGatewayConnectionTunnelsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteTransitGatewayConnectionTunnelsOptionsModel).ToNot(BeNil())
				Expect(deleteTransitGatewayConnectionTunnelsOptionsModel.TransitGatewayID).To(Equal(core.StringPtr("testString")))
				Expect(deleteTransitGatewayConnectionTunnelsOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(deleteTransitGatewayConnectionTunnelsOptionsModel.GreTunnelID).To(Equal(core.StringPtr("testString")))
				Expect(deleteTransitGatewayConnectionTunnelsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
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
			It(`Invoke NewDeleteTransitGatewayRouteReportOptions successfully`, func() {
				// Construct an instance of the DeleteTransitGatewayRouteReportOptions model
				transitGatewayID := "testString"
				id := "testString"
				deleteTransitGatewayRouteReportOptionsModel := transitGatewayApisService.NewDeleteTransitGatewayRouteReportOptions(transitGatewayID, id)
				deleteTransitGatewayRouteReportOptionsModel.SetTransitGatewayID("testString")
				deleteTransitGatewayRouteReportOptionsModel.SetID("testString")
				deleteTransitGatewayRouteReportOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteTransitGatewayRouteReportOptionsModel).ToNot(BeNil())
				Expect(deleteTransitGatewayRouteReportOptionsModel.TransitGatewayID).To(Equal(core.StringPtr("testString")))
				Expect(deleteTransitGatewayRouteReportOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(deleteTransitGatewayRouteReportOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
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
			It(`Invoke NewGetTransitGatewayConnectionPrefixFilterOptions successfully`, func() {
				// Construct an instance of the GetTransitGatewayConnectionPrefixFilterOptions model
				transitGatewayID := "testString"
				id := "testString"
				filterID := "testString"
				getTransitGatewayConnectionPrefixFilterOptionsModel := transitGatewayApisService.NewGetTransitGatewayConnectionPrefixFilterOptions(transitGatewayID, id, filterID)
				getTransitGatewayConnectionPrefixFilterOptionsModel.SetTransitGatewayID("testString")
				getTransitGatewayConnectionPrefixFilterOptionsModel.SetID("testString")
				getTransitGatewayConnectionPrefixFilterOptionsModel.SetFilterID("testString")
				getTransitGatewayConnectionPrefixFilterOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getTransitGatewayConnectionPrefixFilterOptionsModel).ToNot(BeNil())
				Expect(getTransitGatewayConnectionPrefixFilterOptionsModel.TransitGatewayID).To(Equal(core.StringPtr("testString")))
				Expect(getTransitGatewayConnectionPrefixFilterOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(getTransitGatewayConnectionPrefixFilterOptionsModel.FilterID).To(Equal(core.StringPtr("testString")))
				Expect(getTransitGatewayConnectionPrefixFilterOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetTransitGatewayConnectionTunnelsOptions successfully`, func() {
				// Construct an instance of the GetTransitGatewayConnectionTunnelsOptions model
				transitGatewayID := "testString"
				id := "testString"
				greTunnelID := "testString"
				getTransitGatewayConnectionTunnelsOptionsModel := transitGatewayApisService.NewGetTransitGatewayConnectionTunnelsOptions(transitGatewayID, id, greTunnelID)
				getTransitGatewayConnectionTunnelsOptionsModel.SetTransitGatewayID("testString")
				getTransitGatewayConnectionTunnelsOptionsModel.SetID("testString")
				getTransitGatewayConnectionTunnelsOptionsModel.SetGreTunnelID("testString")
				getTransitGatewayConnectionTunnelsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getTransitGatewayConnectionTunnelsOptionsModel).ToNot(BeNil())
				Expect(getTransitGatewayConnectionTunnelsOptionsModel.TransitGatewayID).To(Equal(core.StringPtr("testString")))
				Expect(getTransitGatewayConnectionTunnelsOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(getTransitGatewayConnectionTunnelsOptionsModel.GreTunnelID).To(Equal(core.StringPtr("testString")))
				Expect(getTransitGatewayConnectionTunnelsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
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
			It(`Invoke NewGetTransitGatewayRouteReportOptions successfully`, func() {
				// Construct an instance of the GetTransitGatewayRouteReportOptions model
				transitGatewayID := "testString"
				id := "testString"
				getTransitGatewayRouteReportOptionsModel := transitGatewayApisService.NewGetTransitGatewayRouteReportOptions(transitGatewayID, id)
				getTransitGatewayRouteReportOptionsModel.SetTransitGatewayID("testString")
				getTransitGatewayRouteReportOptionsModel.SetID("testString")
				getTransitGatewayRouteReportOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getTransitGatewayRouteReportOptionsModel).ToNot(BeNil())
				Expect(getTransitGatewayRouteReportOptionsModel.TransitGatewayID).To(Equal(core.StringPtr("testString")))
				Expect(getTransitGatewayRouteReportOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(getTransitGatewayRouteReportOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListConnectionsOptions successfully`, func() {
				// Construct an instance of the ListConnectionsOptions model
				listConnectionsOptionsModel := transitGatewayApisService.NewListConnectionsOptions()
				listConnectionsOptionsModel.SetLimit(int64(10))
				listConnectionsOptionsModel.SetStart("testString")
				listConnectionsOptionsModel.SetNetworkID("testString")
				listConnectionsOptionsModel.SetNetworkType("testString")
				listConnectionsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listConnectionsOptionsModel).ToNot(BeNil())
				Expect(listConnectionsOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(10))))
				Expect(listConnectionsOptionsModel.Start).To(Equal(core.StringPtr("testString")))
				Expect(listConnectionsOptionsModel.NetworkID).To(Equal(core.StringPtr("testString")))
				Expect(listConnectionsOptionsModel.NetworkType).To(Equal(core.StringPtr("testString")))
				Expect(listConnectionsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListGatewayLocationsOptions successfully`, func() {
				// Construct an instance of the ListGatewayLocationsOptions model
				listGatewayLocationsOptionsModel := transitGatewayApisService.NewListGatewayLocationsOptions()
				listGatewayLocationsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listGatewayLocationsOptionsModel).ToNot(BeNil())
				Expect(listGatewayLocationsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListTransitGatewayConnectionPrefixFiltersOptions successfully`, func() {
				// Construct an instance of the ListTransitGatewayConnectionPrefixFiltersOptions model
				transitGatewayID := "testString"
				id := "testString"
				listTransitGatewayConnectionPrefixFiltersOptionsModel := transitGatewayApisService.NewListTransitGatewayConnectionPrefixFiltersOptions(transitGatewayID, id)
				listTransitGatewayConnectionPrefixFiltersOptionsModel.SetTransitGatewayID("testString")
				listTransitGatewayConnectionPrefixFiltersOptionsModel.SetID("testString")
				listTransitGatewayConnectionPrefixFiltersOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listTransitGatewayConnectionPrefixFiltersOptionsModel).ToNot(BeNil())
				Expect(listTransitGatewayConnectionPrefixFiltersOptionsModel.TransitGatewayID).To(Equal(core.StringPtr("testString")))
				Expect(listTransitGatewayConnectionPrefixFiltersOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(listTransitGatewayConnectionPrefixFiltersOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListTransitGatewayConnectionsOptions successfully`, func() {
				// Construct an instance of the ListTransitGatewayConnectionsOptions model
				transitGatewayID := "testString"
				listTransitGatewayConnectionsOptionsModel := transitGatewayApisService.NewListTransitGatewayConnectionsOptions(transitGatewayID)
				listTransitGatewayConnectionsOptionsModel.SetTransitGatewayID("testString")
				listTransitGatewayConnectionsOptionsModel.SetStart("testString")
				listTransitGatewayConnectionsOptionsModel.SetLimit(int64(10))
				listTransitGatewayConnectionsOptionsModel.SetName("testString")
				listTransitGatewayConnectionsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listTransitGatewayConnectionsOptionsModel).ToNot(BeNil())
				Expect(listTransitGatewayConnectionsOptionsModel.TransitGatewayID).To(Equal(core.StringPtr("testString")))
				Expect(listTransitGatewayConnectionsOptionsModel.Start).To(Equal(core.StringPtr("testString")))
				Expect(listTransitGatewayConnectionsOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(10))))
				Expect(listTransitGatewayConnectionsOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(listTransitGatewayConnectionsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListTransitGatewayGreTunnelOptions successfully`, func() {
				// Construct an instance of the ListTransitGatewayGreTunnelOptions model
				transitGatewayID := "testString"
				id := "testString"
				listTransitGatewayGreTunnelOptionsModel := transitGatewayApisService.NewListTransitGatewayGreTunnelOptions(transitGatewayID, id)
				listTransitGatewayGreTunnelOptionsModel.SetTransitGatewayID("testString")
				listTransitGatewayGreTunnelOptionsModel.SetID("testString")
				listTransitGatewayGreTunnelOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listTransitGatewayGreTunnelOptionsModel).ToNot(BeNil())
				Expect(listTransitGatewayGreTunnelOptionsModel.TransitGatewayID).To(Equal(core.StringPtr("testString")))
				Expect(listTransitGatewayGreTunnelOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(listTransitGatewayGreTunnelOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListTransitGatewayRouteReportsOptions successfully`, func() {
				// Construct an instance of the ListTransitGatewayRouteReportsOptions model
				transitGatewayID := "testString"
				listTransitGatewayRouteReportsOptionsModel := transitGatewayApisService.NewListTransitGatewayRouteReportsOptions(transitGatewayID)
				listTransitGatewayRouteReportsOptionsModel.SetTransitGatewayID("testString")
				listTransitGatewayRouteReportsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listTransitGatewayRouteReportsOptionsModel).ToNot(BeNil())
				Expect(listTransitGatewayRouteReportsOptionsModel.TransitGatewayID).To(Equal(core.StringPtr("testString")))
				Expect(listTransitGatewayRouteReportsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListTransitGatewaysOptions successfully`, func() {
				// Construct an instance of the ListTransitGatewaysOptions model
				listTransitGatewaysOptionsModel := transitGatewayApisService.NewListTransitGatewaysOptions()
				listTransitGatewaysOptionsModel.SetLimit(int64(10))
				listTransitGatewaysOptionsModel.SetStart("testString")
				listTransitGatewaysOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listTransitGatewaysOptionsModel).ToNot(BeNil())
				Expect(listTransitGatewaysOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(10))))
				Expect(listTransitGatewaysOptionsModel.Start).To(Equal(core.StringPtr("testString")))
				Expect(listTransitGatewaysOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewPrefixFilterPut successfully`, func() {
				action := "permit"
				prefix := "192.168.100.0/24"
				_model, err := transitGatewayApisService.NewPrefixFilterPut(action, prefix)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewReplaceTransitGatewayConnectionPrefixFilterOptions successfully`, func() {
				// Construct an instance of the PrefixFilterPut model
				prefixFilterPutModel := new(transitgatewayapisv1.PrefixFilterPut)
				Expect(prefixFilterPutModel).ToNot(BeNil())
				prefixFilterPutModel.Action = core.StringPtr("permit")
				prefixFilterPutModel.Ge = core.Int64Ptr(int64(0))
				prefixFilterPutModel.Le = core.Int64Ptr(int64(32))
				prefixFilterPutModel.Prefix = core.StringPtr("192.168.100.0/24")
				Expect(prefixFilterPutModel.Action).To(Equal(core.StringPtr("permit")))
				Expect(prefixFilterPutModel.Ge).To(Equal(core.Int64Ptr(int64(0))))
				Expect(prefixFilterPutModel.Le).To(Equal(core.Int64Ptr(int64(32))))
				Expect(prefixFilterPutModel.Prefix).To(Equal(core.StringPtr("192.168.100.0/24")))

				// Construct an instance of the ReplaceTransitGatewayConnectionPrefixFilterOptions model
				transitGatewayID := "testString"
				id := "testString"
				replaceTransitGatewayConnectionPrefixFilterOptionsPrefixFilters := []transitgatewayapisv1.PrefixFilterPut{}
				replaceTransitGatewayConnectionPrefixFilterOptionsModel := transitGatewayApisService.NewReplaceTransitGatewayConnectionPrefixFilterOptions(transitGatewayID, id, replaceTransitGatewayConnectionPrefixFilterOptionsPrefixFilters)
				replaceTransitGatewayConnectionPrefixFilterOptionsModel.SetTransitGatewayID("testString")
				replaceTransitGatewayConnectionPrefixFilterOptionsModel.SetID("testString")
				replaceTransitGatewayConnectionPrefixFilterOptionsModel.SetPrefixFilters([]transitgatewayapisv1.PrefixFilterPut{*prefixFilterPutModel})
				replaceTransitGatewayConnectionPrefixFilterOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(replaceTransitGatewayConnectionPrefixFilterOptionsModel).ToNot(BeNil())
				Expect(replaceTransitGatewayConnectionPrefixFilterOptionsModel.TransitGatewayID).To(Equal(core.StringPtr("testString")))
				Expect(replaceTransitGatewayConnectionPrefixFilterOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(replaceTransitGatewayConnectionPrefixFilterOptionsModel.PrefixFilters).To(Equal([]transitgatewayapisv1.PrefixFilterPut{*prefixFilterPutModel}))
				Expect(replaceTransitGatewayConnectionPrefixFilterOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewResourceGroupIdentity successfully`, func() {
				id := "56969d6043e9465c883cb9f7363e78e8"
				_model, err := transitGatewayApisService.NewResourceGroupIdentity(id)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewTransitGatewayConnectionPrefixFilter successfully`, func() {
				action := "permit"
				prefix := "192.168.100.0/24"
				_model, err := transitGatewayApisService.NewTransitGatewayConnectionPrefixFilter(action, prefix)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewTransitGatewayTunnelTemplate successfully`, func() {
				localGatewayIp := "10.242.63.12"
				localTunnelIp := "192.168.100.20"
				name := "gre1"
				remoteGatewayIp := "10.242.33.22"
				remoteTunnelIp := "192.168.129.1"
				var zone transitgatewayapisv1.ZoneIdentityIntf = nil
				_, err := transitGatewayApisService.NewTransitGatewayTunnelTemplate(localGatewayIp, localTunnelIp, name, remoteGatewayIp, remoteTunnelIp, zone)
				Expect(err).ToNot(BeNil())
			})
			It(`Invoke NewUpdateTransitGatewayConnectionOptions successfully`, func() {
				// Construct an instance of the UpdateTransitGatewayConnectionOptions model
				transitGatewayID := "testString"
				id := "testString"
				updateTransitGatewayConnectionOptionsModel := transitGatewayApisService.NewUpdateTransitGatewayConnectionOptions(transitGatewayID, id)
				updateTransitGatewayConnectionOptionsModel.SetTransitGatewayID("testString")
				updateTransitGatewayConnectionOptionsModel.SetID("testString")
				updateTransitGatewayConnectionOptionsModel.SetName("Transit_Service_BWTN_SJ_DL")
				updateTransitGatewayConnectionOptionsModel.SetPrefixFiltersDefault("permit")
				updateTransitGatewayConnectionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateTransitGatewayConnectionOptionsModel).ToNot(BeNil())
				Expect(updateTransitGatewayConnectionOptionsModel.TransitGatewayID).To(Equal(core.StringPtr("testString")))
				Expect(updateTransitGatewayConnectionOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(updateTransitGatewayConnectionOptionsModel.Name).To(Equal(core.StringPtr("Transit_Service_BWTN_SJ_DL")))
				Expect(updateTransitGatewayConnectionOptionsModel.PrefixFiltersDefault).To(Equal(core.StringPtr("permit")))
				Expect(updateTransitGatewayConnectionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateTransitGatewayConnectionPrefixFilterOptions successfully`, func() {
				// Construct an instance of the UpdateTransitGatewayConnectionPrefixFilterOptions model
				transitGatewayID := "testString"
				id := "testString"
				filterID := "testString"
				updateTransitGatewayConnectionPrefixFilterOptionsModel := transitGatewayApisService.NewUpdateTransitGatewayConnectionPrefixFilterOptions(transitGatewayID, id, filterID)
				updateTransitGatewayConnectionPrefixFilterOptionsModel.SetTransitGatewayID("testString")
				updateTransitGatewayConnectionPrefixFilterOptionsModel.SetID("testString")
				updateTransitGatewayConnectionPrefixFilterOptionsModel.SetFilterID("testString")
				updateTransitGatewayConnectionPrefixFilterOptionsModel.SetAction("permit")
				updateTransitGatewayConnectionPrefixFilterOptionsModel.SetBefore("1a15dcab-7e40-45e1-b7c5-bc690eaa9782")
				updateTransitGatewayConnectionPrefixFilterOptionsModel.SetGe(int64(0))
				updateTransitGatewayConnectionPrefixFilterOptionsModel.SetLe(int64(32))
				updateTransitGatewayConnectionPrefixFilterOptionsModel.SetPrefix("192.168.100.0/24")
				updateTransitGatewayConnectionPrefixFilterOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateTransitGatewayConnectionPrefixFilterOptionsModel).ToNot(BeNil())
				Expect(updateTransitGatewayConnectionPrefixFilterOptionsModel.TransitGatewayID).To(Equal(core.StringPtr("testString")))
				Expect(updateTransitGatewayConnectionPrefixFilterOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(updateTransitGatewayConnectionPrefixFilterOptionsModel.FilterID).To(Equal(core.StringPtr("testString")))
				Expect(updateTransitGatewayConnectionPrefixFilterOptionsModel.Action).To(Equal(core.StringPtr("permit")))
				Expect(updateTransitGatewayConnectionPrefixFilterOptionsModel.Before).To(Equal(core.StringPtr("1a15dcab-7e40-45e1-b7c5-bc690eaa9782")))
				Expect(updateTransitGatewayConnectionPrefixFilterOptionsModel.Ge).To(Equal(core.Int64Ptr(int64(0))))
				Expect(updateTransitGatewayConnectionPrefixFilterOptionsModel.Le).To(Equal(core.Int64Ptr(int64(32))))
				Expect(updateTransitGatewayConnectionPrefixFilterOptionsModel.Prefix).To(Equal(core.StringPtr("192.168.100.0/24")))
				Expect(updateTransitGatewayConnectionPrefixFilterOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateTransitGatewayConnectionTunnelsOptions successfully`, func() {
				// Construct an instance of the UpdateTransitGatewayConnectionTunnelsOptions model
				transitGatewayID := "testString"
				id := "testString"
				greTunnelID := "testString"
				transitGatewayTunnelPatch := map[string]interface{}{"anyKey": "anyValue"}
				updateTransitGatewayConnectionTunnelsOptionsModel := transitGatewayApisService.NewUpdateTransitGatewayConnectionTunnelsOptions(transitGatewayID, id, greTunnelID, transitGatewayTunnelPatch)
				updateTransitGatewayConnectionTunnelsOptionsModel.SetTransitGatewayID("testString")
				updateTransitGatewayConnectionTunnelsOptionsModel.SetID("testString")
				updateTransitGatewayConnectionTunnelsOptionsModel.SetGreTunnelID("testString")
				updateTransitGatewayConnectionTunnelsOptionsModel.SetTransitGatewayTunnelPatch(map[string]interface{}{"anyKey": "anyValue"})
				updateTransitGatewayConnectionTunnelsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateTransitGatewayConnectionTunnelsOptionsModel).ToNot(BeNil())
				Expect(updateTransitGatewayConnectionTunnelsOptionsModel.TransitGatewayID).To(Equal(core.StringPtr("testString")))
				Expect(updateTransitGatewayConnectionTunnelsOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(updateTransitGatewayConnectionTunnelsOptionsModel.GreTunnelID).To(Equal(core.StringPtr("testString")))
				Expect(updateTransitGatewayConnectionTunnelsOptionsModel.TransitGatewayTunnelPatch).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))
				Expect(updateTransitGatewayConnectionTunnelsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateTransitGatewayOptions successfully`, func() {
				// Construct an instance of the UpdateTransitGatewayOptions model
				id := "testString"
				updateTransitGatewayOptionsModel := transitGatewayApisService.NewUpdateTransitGatewayOptions(id)
				updateTransitGatewayOptionsModel.SetID("testString")
				updateTransitGatewayOptionsModel.SetGlobal(true)
				updateTransitGatewayOptionsModel.SetGreEnhancedRoutePropagation(true)
				updateTransitGatewayOptionsModel.SetName("my-resource")
				updateTransitGatewayOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateTransitGatewayOptionsModel).ToNot(BeNil())
				Expect(updateTransitGatewayOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(updateTransitGatewayOptionsModel.Global).To(Equal(core.BoolPtr(true)))
				Expect(updateTransitGatewayOptionsModel.GreEnhancedRoutePropagation).To(Equal(core.BoolPtr(true)))
				Expect(updateTransitGatewayOptionsModel.Name).To(Equal(core.StringPtr("my-resource")))
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
