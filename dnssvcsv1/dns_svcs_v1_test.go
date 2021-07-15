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

package dnssvcsv1_test

import (
	"bytes"
	"context"
	"fmt"
	"github.com/IBM/go-sdk-core/v4/core"
	"github.com/IBM/networking-go-sdk/dnssvcsv1"
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

var _ = Describe(`DnsSvcsV1`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(dnsSvcsService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(dnsSvcsService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
				URL: "https://dnssvcsv1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(dnsSvcsService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"DNS_SVCS_URL": "https://dnssvcsv1/api",
				"DNS_SVCS_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1UsingExternalConfig(&dnssvcsv1.DnsSvcsV1Options{
				})
				Expect(dnsSvcsService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := dnsSvcsService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != dnsSvcsService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(dnsSvcsService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(dnsSvcsService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1UsingExternalConfig(&dnssvcsv1.DnsSvcsV1Options{
					URL: "https://testService/api",
				})
				Expect(dnsSvcsService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(dnsSvcsService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := dnsSvcsService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != dnsSvcsService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(dnsSvcsService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(dnsSvcsService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1UsingExternalConfig(&dnssvcsv1.DnsSvcsV1Options{
				})
				err := dnsSvcsService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(dnsSvcsService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(dnsSvcsService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := dnsSvcsService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != dnsSvcsService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(dnsSvcsService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(dnsSvcsService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"DNS_SVCS_URL": "https://dnssvcsv1/api",
				"DNS_SVCS_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1UsingExternalConfig(&dnssvcsv1.DnsSvcsV1Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(dnsSvcsService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"DNS_SVCS_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1UsingExternalConfig(&dnssvcsv1.DnsSvcsV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(dnsSvcsService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = dnssvcsv1.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`ListDnszones(listDnszonesOptions *ListDnszonesOptions) - Operation response error`, func() {
		listDnszonesPath := "/instances/testString/dnszones"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listDnszonesPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(38))}))

					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(200))}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListDnszones with error: Operation response processing error`, func() {
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsSvcsService).ToNot(BeNil())

				// Construct an instance of the ListDnszonesOptions model
				listDnszonesOptionsModel := new(dnssvcsv1.ListDnszonesOptions)
				listDnszonesOptionsModel.InstanceID = core.StringPtr("testString")
				listDnszonesOptionsModel.XCorrelationID = core.StringPtr("testString")
				listDnszonesOptionsModel.Offset = core.Int64Ptr(int64(38))
				listDnszonesOptionsModel.Limit = core.Int64Ptr(int64(200))
				listDnszonesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := dnsSvcsService.ListDnszones(listDnszonesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				dnsSvcsService.EnableRetries(0, 0)
				result, response, operationErr = dnsSvcsService.ListDnszones(listDnszonesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ListDnszones(listDnszonesOptions *ListDnszonesOptions)`, func() {
		listDnszonesPath := "/instances/testString/dnszones"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listDnszonesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(38))}))

					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(200))}))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"dnszones": [{"id": "example.com:2d0f862b-67cc-41f3-b6a2-59860d0aa90e", "created_on": "2019-01-01T05:20:00.12345Z", "modified_on": "2019-01-01T05:20:00.12345Z", "instance_id": "1407a753-a93f-4bb0-9784-bcfc269ee1b3", "name": "example.com", "description": "The DNS zone is used for VPCs in us-east region", "state": "pending_network_add", "label": "us-east"}], "offset": 0, "limit": 10, "total_count": 10, "first": {"href": "https://api.dns-svcs.cloud.ibm.com/v1/instances/434f6c3e-6014-4124-a61d-2e910bca19b1/dnszones/example.com:d04d3a7a-7f6d-47d4-b811-08c5478fa1a4/resource_records?limit=20"}, "next": {"href": "https://api.dns-svcs.cloud.ibm.com/v1/instances/434f6c3e-6014-4124-a61d-2e910bca19b1/dnszones/example.com:d04d3a7a-7f6d-47d4-b811-08c5478fa1a4/resource_records?offset=20&limit=20"}}`)
				}))
			})
			It(`Invoke ListDnszones successfully`, func() {
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsSvcsService).ToNot(BeNil())
				dnsSvcsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := dnsSvcsService.ListDnszones(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListDnszonesOptions model
				listDnszonesOptionsModel := new(dnssvcsv1.ListDnszonesOptions)
				listDnszonesOptionsModel.InstanceID = core.StringPtr("testString")
				listDnszonesOptionsModel.XCorrelationID = core.StringPtr("testString")
				listDnszonesOptionsModel.Offset = core.Int64Ptr(int64(38))
				listDnszonesOptionsModel.Limit = core.Int64Ptr(int64(200))
				listDnszonesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = dnsSvcsService.ListDnszones(listDnszonesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = dnsSvcsService.ListDnszonesWithContext(ctx, listDnszonesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				dnsSvcsService.DisableRetries()
				result, response, operationErr = dnsSvcsService.ListDnszones(listDnszonesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = dnsSvcsService.ListDnszonesWithContext(ctx, listDnszonesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke ListDnszones with error: Operation validation and request error`, func() {
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsSvcsService).ToNot(BeNil())

				// Construct an instance of the ListDnszonesOptions model
				listDnszonesOptionsModel := new(dnssvcsv1.ListDnszonesOptions)
				listDnszonesOptionsModel.InstanceID = core.StringPtr("testString")
				listDnszonesOptionsModel.XCorrelationID = core.StringPtr("testString")
				listDnszonesOptionsModel.Offset = core.Int64Ptr(int64(38))
				listDnszonesOptionsModel.Limit = core.Int64Ptr(int64(200))
				listDnszonesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := dnsSvcsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := dnsSvcsService.ListDnszones(listDnszonesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListDnszonesOptions model with no property values
				listDnszonesOptionsModelNew := new(dnssvcsv1.ListDnszonesOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = dnsSvcsService.ListDnszones(listDnszonesOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateDnszone(createDnszoneOptions *CreateDnszoneOptions) - Operation response error`, func() {
		createDnszonePath := "/instances/testString/dnszones"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createDnszonePath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateDnszone with error: Operation response processing error`, func() {
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsSvcsService).ToNot(BeNil())

				// Construct an instance of the CreateDnszoneOptions model
				createDnszoneOptionsModel := new(dnssvcsv1.CreateDnszoneOptions)
				createDnszoneOptionsModel.InstanceID = core.StringPtr("testString")
				createDnszoneOptionsModel.Name = core.StringPtr("example.com")
				createDnszoneOptionsModel.Description = core.StringPtr("The DNS zone is used for VPCs in us-east region")
				createDnszoneOptionsModel.Label = core.StringPtr("us-east")
				createDnszoneOptionsModel.XCorrelationID = core.StringPtr("testString")
				createDnszoneOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := dnsSvcsService.CreateDnszone(createDnszoneOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				dnsSvcsService.EnableRetries(0, 0)
				result, response, operationErr = dnsSvcsService.CreateDnszone(createDnszoneOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`CreateDnszone(createDnszoneOptions *CreateDnszoneOptions)`, func() {
		createDnszonePath := "/instances/testString/dnszones"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createDnszonePath))
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

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "example.com:2d0f862b-67cc-41f3-b6a2-59860d0aa90e", "created_on": "2019-01-01T05:20:00.12345Z", "modified_on": "2019-01-01T05:20:00.12345Z", "instance_id": "1407a753-a93f-4bb0-9784-bcfc269ee1b3", "name": "example.com", "description": "The DNS zone is used for VPCs in us-east region", "state": "pending_network_add", "label": "us-east"}`)
				}))
			})
			It(`Invoke CreateDnszone successfully`, func() {
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsSvcsService).ToNot(BeNil())
				dnsSvcsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := dnsSvcsService.CreateDnszone(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateDnszoneOptions model
				createDnszoneOptionsModel := new(dnssvcsv1.CreateDnszoneOptions)
				createDnszoneOptionsModel.InstanceID = core.StringPtr("testString")
				createDnszoneOptionsModel.Name = core.StringPtr("example.com")
				createDnszoneOptionsModel.Description = core.StringPtr("The DNS zone is used for VPCs in us-east region")
				createDnszoneOptionsModel.Label = core.StringPtr("us-east")
				createDnszoneOptionsModel.XCorrelationID = core.StringPtr("testString")
				createDnszoneOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = dnsSvcsService.CreateDnszone(createDnszoneOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = dnsSvcsService.CreateDnszoneWithContext(ctx, createDnszoneOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				dnsSvcsService.DisableRetries()
				result, response, operationErr = dnsSvcsService.CreateDnszone(createDnszoneOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = dnsSvcsService.CreateDnszoneWithContext(ctx, createDnszoneOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke CreateDnszone with error: Operation validation and request error`, func() {
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsSvcsService).ToNot(BeNil())

				// Construct an instance of the CreateDnszoneOptions model
				createDnszoneOptionsModel := new(dnssvcsv1.CreateDnszoneOptions)
				createDnszoneOptionsModel.InstanceID = core.StringPtr("testString")
				createDnszoneOptionsModel.Name = core.StringPtr("example.com")
				createDnszoneOptionsModel.Description = core.StringPtr("The DNS zone is used for VPCs in us-east region")
				createDnszoneOptionsModel.Label = core.StringPtr("us-east")
				createDnszoneOptionsModel.XCorrelationID = core.StringPtr("testString")
				createDnszoneOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := dnsSvcsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := dnsSvcsService.CreateDnszone(createDnszoneOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateDnszoneOptions model with no property values
				createDnszoneOptionsModelNew := new(dnssvcsv1.CreateDnszoneOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = dnsSvcsService.CreateDnszone(createDnszoneOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`DeleteDnszone(deleteDnszoneOptions *DeleteDnszoneOptions)`, func() {
		deleteDnszonePath := "/instances/testString/dnszones/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteDnszonePath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteDnszone successfully`, func() {
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsSvcsService).ToNot(BeNil())
				dnsSvcsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				response, operationErr := dnsSvcsService.DeleteDnszone(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteDnszoneOptions model
				deleteDnszoneOptionsModel := new(dnssvcsv1.DeleteDnszoneOptions)
				deleteDnszoneOptionsModel.InstanceID = core.StringPtr("testString")
				deleteDnszoneOptionsModel.DnszoneID = core.StringPtr("testString")
				deleteDnszoneOptionsModel.XCorrelationID = core.StringPtr("testString")
				deleteDnszoneOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = dnsSvcsService.DeleteDnszone(deleteDnszoneOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Disable retries and test again
				dnsSvcsService.DisableRetries()
				response, operationErr = dnsSvcsService.DeleteDnszone(deleteDnszoneOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteDnszone with error: Operation validation and request error`, func() {
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsSvcsService).ToNot(BeNil())

				// Construct an instance of the DeleteDnszoneOptions model
				deleteDnszoneOptionsModel := new(dnssvcsv1.DeleteDnszoneOptions)
				deleteDnszoneOptionsModel.InstanceID = core.StringPtr("testString")
				deleteDnszoneOptionsModel.DnszoneID = core.StringPtr("testString")
				deleteDnszoneOptionsModel.XCorrelationID = core.StringPtr("testString")
				deleteDnszoneOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := dnsSvcsService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := dnsSvcsService.DeleteDnszone(deleteDnszoneOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteDnszoneOptions model with no property values
				deleteDnszoneOptionsModelNew := new(dnssvcsv1.DeleteDnszoneOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = dnsSvcsService.DeleteDnszone(deleteDnszoneOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetDnszone(getDnszoneOptions *GetDnszoneOptions) - Operation response error`, func() {
		getDnszonePath := "/instances/testString/dnszones/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getDnszonePath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetDnszone with error: Operation response processing error`, func() {
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsSvcsService).ToNot(BeNil())

				// Construct an instance of the GetDnszoneOptions model
				getDnszoneOptionsModel := new(dnssvcsv1.GetDnszoneOptions)
				getDnszoneOptionsModel.InstanceID = core.StringPtr("testString")
				getDnszoneOptionsModel.DnszoneID = core.StringPtr("testString")
				getDnszoneOptionsModel.XCorrelationID = core.StringPtr("testString")
				getDnszoneOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := dnsSvcsService.GetDnszone(getDnszoneOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				dnsSvcsService.EnableRetries(0, 0)
				result, response, operationErr = dnsSvcsService.GetDnszone(getDnszoneOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetDnszone(getDnszoneOptions *GetDnszoneOptions)`, func() {
		getDnszonePath := "/instances/testString/dnszones/testString"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getDnszonePath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "example.com:2d0f862b-67cc-41f3-b6a2-59860d0aa90e", "created_on": "2019-01-01T05:20:00.12345Z", "modified_on": "2019-01-01T05:20:00.12345Z", "instance_id": "1407a753-a93f-4bb0-9784-bcfc269ee1b3", "name": "example.com", "description": "The DNS zone is used for VPCs in us-east region", "state": "pending_network_add", "label": "us-east"}`)
				}))
			})
			It(`Invoke GetDnszone successfully`, func() {
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsSvcsService).ToNot(BeNil())
				dnsSvcsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := dnsSvcsService.GetDnszone(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetDnszoneOptions model
				getDnszoneOptionsModel := new(dnssvcsv1.GetDnszoneOptions)
				getDnszoneOptionsModel.InstanceID = core.StringPtr("testString")
				getDnszoneOptionsModel.DnszoneID = core.StringPtr("testString")
				getDnszoneOptionsModel.XCorrelationID = core.StringPtr("testString")
				getDnszoneOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = dnsSvcsService.GetDnszone(getDnszoneOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = dnsSvcsService.GetDnszoneWithContext(ctx, getDnszoneOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				dnsSvcsService.DisableRetries()
				result, response, operationErr = dnsSvcsService.GetDnszone(getDnszoneOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = dnsSvcsService.GetDnszoneWithContext(ctx, getDnszoneOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke GetDnszone with error: Operation validation and request error`, func() {
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsSvcsService).ToNot(BeNil())

				// Construct an instance of the GetDnszoneOptions model
				getDnszoneOptionsModel := new(dnssvcsv1.GetDnszoneOptions)
				getDnszoneOptionsModel.InstanceID = core.StringPtr("testString")
				getDnszoneOptionsModel.DnszoneID = core.StringPtr("testString")
				getDnszoneOptionsModel.XCorrelationID = core.StringPtr("testString")
				getDnszoneOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := dnsSvcsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := dnsSvcsService.GetDnszone(getDnszoneOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetDnszoneOptions model with no property values
				getDnszoneOptionsModelNew := new(dnssvcsv1.GetDnszoneOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = dnsSvcsService.GetDnszone(getDnszoneOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateDnszone(updateDnszoneOptions *UpdateDnszoneOptions) - Operation response error`, func() {
		updateDnszonePath := "/instances/testString/dnszones/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateDnszonePath))
					Expect(req.Method).To(Equal("PATCH"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateDnszone with error: Operation response processing error`, func() {
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsSvcsService).ToNot(BeNil())

				// Construct an instance of the UpdateDnszoneOptions model
				updateDnszoneOptionsModel := new(dnssvcsv1.UpdateDnszoneOptions)
				updateDnszoneOptionsModel.InstanceID = core.StringPtr("testString")
				updateDnszoneOptionsModel.DnszoneID = core.StringPtr("testString")
				updateDnszoneOptionsModel.Description = core.StringPtr("The DNS zone is used for VPCs in us-east region")
				updateDnszoneOptionsModel.Label = core.StringPtr("us-east")
				updateDnszoneOptionsModel.XCorrelationID = core.StringPtr("testString")
				updateDnszoneOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := dnsSvcsService.UpdateDnszone(updateDnszoneOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				dnsSvcsService.EnableRetries(0, 0)
				result, response, operationErr = dnsSvcsService.UpdateDnszone(updateDnszoneOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UpdateDnszone(updateDnszoneOptions *UpdateDnszoneOptions)`, func() {
		updateDnszonePath := "/instances/testString/dnszones/testString"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateDnszonePath))
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

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "example.com:2d0f862b-67cc-41f3-b6a2-59860d0aa90e", "created_on": "2019-01-01T05:20:00.12345Z", "modified_on": "2019-01-01T05:20:00.12345Z", "instance_id": "1407a753-a93f-4bb0-9784-bcfc269ee1b3", "name": "example.com", "description": "The DNS zone is used for VPCs in us-east region", "state": "pending_network_add", "label": "us-east"}`)
				}))
			})
			It(`Invoke UpdateDnszone successfully`, func() {
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsSvcsService).ToNot(BeNil())
				dnsSvcsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := dnsSvcsService.UpdateDnszone(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateDnszoneOptions model
				updateDnszoneOptionsModel := new(dnssvcsv1.UpdateDnszoneOptions)
				updateDnszoneOptionsModel.InstanceID = core.StringPtr("testString")
				updateDnszoneOptionsModel.DnszoneID = core.StringPtr("testString")
				updateDnszoneOptionsModel.Description = core.StringPtr("The DNS zone is used for VPCs in us-east region")
				updateDnszoneOptionsModel.Label = core.StringPtr("us-east")
				updateDnszoneOptionsModel.XCorrelationID = core.StringPtr("testString")
				updateDnszoneOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = dnsSvcsService.UpdateDnszone(updateDnszoneOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = dnsSvcsService.UpdateDnszoneWithContext(ctx, updateDnszoneOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				dnsSvcsService.DisableRetries()
				result, response, operationErr = dnsSvcsService.UpdateDnszone(updateDnszoneOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = dnsSvcsService.UpdateDnszoneWithContext(ctx, updateDnszoneOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke UpdateDnszone with error: Operation validation and request error`, func() {
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsSvcsService).ToNot(BeNil())

				// Construct an instance of the UpdateDnszoneOptions model
				updateDnszoneOptionsModel := new(dnssvcsv1.UpdateDnszoneOptions)
				updateDnszoneOptionsModel.InstanceID = core.StringPtr("testString")
				updateDnszoneOptionsModel.DnszoneID = core.StringPtr("testString")
				updateDnszoneOptionsModel.Description = core.StringPtr("The DNS zone is used for VPCs in us-east region")
				updateDnszoneOptionsModel.Label = core.StringPtr("us-east")
				updateDnszoneOptionsModel.XCorrelationID = core.StringPtr("testString")
				updateDnszoneOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := dnsSvcsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := dnsSvcsService.UpdateDnszone(updateDnszoneOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateDnszoneOptions model with no property values
				updateDnszoneOptionsModelNew := new(dnssvcsv1.UpdateDnszoneOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = dnsSvcsService.UpdateDnszone(updateDnszoneOptionsModelNew)
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
		It(`Instantiate service client`, func() {
			dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(dnsSvcsService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(dnsSvcsService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
				URL: "https://dnssvcsv1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(dnsSvcsService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"DNS_SVCS_URL": "https://dnssvcsv1/api",
				"DNS_SVCS_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1UsingExternalConfig(&dnssvcsv1.DnsSvcsV1Options{
				})
				Expect(dnsSvcsService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := dnsSvcsService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != dnsSvcsService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(dnsSvcsService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(dnsSvcsService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1UsingExternalConfig(&dnssvcsv1.DnsSvcsV1Options{
					URL: "https://testService/api",
				})
				Expect(dnsSvcsService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(dnsSvcsService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := dnsSvcsService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != dnsSvcsService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(dnsSvcsService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(dnsSvcsService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1UsingExternalConfig(&dnssvcsv1.DnsSvcsV1Options{
				})
				err := dnsSvcsService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(dnsSvcsService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(dnsSvcsService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := dnsSvcsService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != dnsSvcsService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(dnsSvcsService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(dnsSvcsService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"DNS_SVCS_URL": "https://dnssvcsv1/api",
				"DNS_SVCS_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1UsingExternalConfig(&dnssvcsv1.DnsSvcsV1Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(dnsSvcsService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"DNS_SVCS_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1UsingExternalConfig(&dnssvcsv1.DnsSvcsV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(dnsSvcsService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = dnssvcsv1.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`ListResourceRecords(listResourceRecordsOptions *ListResourceRecordsOptions) - Operation response error`, func() {
		listResourceRecordsPath := "/instances/testString/dnszones/testString/resource_records"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listResourceRecordsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(38))}))

					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(200))}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListResourceRecords with error: Operation response processing error`, func() {
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsSvcsService).ToNot(BeNil())

				// Construct an instance of the ListResourceRecordsOptions model
				listResourceRecordsOptionsModel := new(dnssvcsv1.ListResourceRecordsOptions)
				listResourceRecordsOptionsModel.InstanceID = core.StringPtr("testString")
				listResourceRecordsOptionsModel.DnszoneID = core.StringPtr("testString")
				listResourceRecordsOptionsModel.XCorrelationID = core.StringPtr("testString")
				listResourceRecordsOptionsModel.Offset = core.Int64Ptr(int64(38))
				listResourceRecordsOptionsModel.Limit = core.Int64Ptr(int64(200))
				listResourceRecordsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := dnsSvcsService.ListResourceRecords(listResourceRecordsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				dnsSvcsService.EnableRetries(0, 0)
				result, response, operationErr = dnsSvcsService.ListResourceRecords(listResourceRecordsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ListResourceRecords(listResourceRecordsOptions *ListResourceRecordsOptions)`, func() {
		listResourceRecordsPath := "/instances/testString/dnszones/testString/resource_records"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listResourceRecordsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(38))}))

					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(200))}))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"resource_records": [{"id": "SRV:5365b73c-ce6f-4d6f-ad9f-d9c131b26370", "created_on": "2019-01-01T05:20:00.12345Z", "modified_on": "2019-01-01T05:20:00.12345Z", "name": "_sip._udp.test.example.com", "type": "SRV", "ttl": 120, "rdata": {"anyKey": "anyValue"}, "service": "_sip", "protocol": "udp"}], "offset": 0, "limit": 20, "total_count": 200, "first": {"href": "https://api.dns-svcs.cloud.ibm.com/v1/instances/434f6c3e-6014-4124-a61d-2e910bca19b1/dnszones/example.com:d04d3a7a-7f6d-47d4-b811-08c5478fa1a4/resource_records?limit=20"}, "next": {"href": "https://api.dns-svcs.cloud.ibm.com/v1/instances/434f6c3e-6014-4124-a61d-2e910bca19b1/dnszones/example.com:d04d3a7a-7f6d-47d4-b811-08c5478fa1a4/resource_records?offset=20&limit=20"}}`)
				}))
			})
			It(`Invoke ListResourceRecords successfully`, func() {
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsSvcsService).ToNot(BeNil())
				dnsSvcsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := dnsSvcsService.ListResourceRecords(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListResourceRecordsOptions model
				listResourceRecordsOptionsModel := new(dnssvcsv1.ListResourceRecordsOptions)
				listResourceRecordsOptionsModel.InstanceID = core.StringPtr("testString")
				listResourceRecordsOptionsModel.DnszoneID = core.StringPtr("testString")
				listResourceRecordsOptionsModel.XCorrelationID = core.StringPtr("testString")
				listResourceRecordsOptionsModel.Offset = core.Int64Ptr(int64(38))
				listResourceRecordsOptionsModel.Limit = core.Int64Ptr(int64(200))
				listResourceRecordsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = dnsSvcsService.ListResourceRecords(listResourceRecordsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = dnsSvcsService.ListResourceRecordsWithContext(ctx, listResourceRecordsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				dnsSvcsService.DisableRetries()
				result, response, operationErr = dnsSvcsService.ListResourceRecords(listResourceRecordsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = dnsSvcsService.ListResourceRecordsWithContext(ctx, listResourceRecordsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke ListResourceRecords with error: Operation validation and request error`, func() {
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsSvcsService).ToNot(BeNil())

				// Construct an instance of the ListResourceRecordsOptions model
				listResourceRecordsOptionsModel := new(dnssvcsv1.ListResourceRecordsOptions)
				listResourceRecordsOptionsModel.InstanceID = core.StringPtr("testString")
				listResourceRecordsOptionsModel.DnszoneID = core.StringPtr("testString")
				listResourceRecordsOptionsModel.XCorrelationID = core.StringPtr("testString")
				listResourceRecordsOptionsModel.Offset = core.Int64Ptr(int64(38))
				listResourceRecordsOptionsModel.Limit = core.Int64Ptr(int64(200))
				listResourceRecordsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := dnsSvcsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := dnsSvcsService.ListResourceRecords(listResourceRecordsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListResourceRecordsOptions model with no property values
				listResourceRecordsOptionsModelNew := new(dnssvcsv1.ListResourceRecordsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = dnsSvcsService.ListResourceRecords(listResourceRecordsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateResourceRecord(createResourceRecordOptions *CreateResourceRecordOptions) - Operation response error`, func() {
		createResourceRecordPath := "/instances/testString/dnszones/testString/resource_records"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createResourceRecordPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateResourceRecord with error: Operation response processing error`, func() {
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsSvcsService).ToNot(BeNil())

				// Construct an instance of the ResourceRecordInputRdataRdataARecord model
				resourceRecordInputRdataModel := new(dnssvcsv1.ResourceRecordInputRdataRdataARecord)
				resourceRecordInputRdataModel.Ip = core.StringPtr("10.110.201.214")

				// Construct an instance of the CreateResourceRecordOptions model
				createResourceRecordOptionsModel := new(dnssvcsv1.CreateResourceRecordOptions)
				createResourceRecordOptionsModel.InstanceID = core.StringPtr("testString")
				createResourceRecordOptionsModel.DnszoneID = core.StringPtr("testString")
				createResourceRecordOptionsModel.Name = core.StringPtr("test.example.com")
				createResourceRecordOptionsModel.Type = core.StringPtr("SRV")
				createResourceRecordOptionsModel.Rdata = resourceRecordInputRdataModel
				createResourceRecordOptionsModel.TTL = core.Int64Ptr(int64(120))
				createResourceRecordOptionsModel.Service = core.StringPtr("_sip")
				createResourceRecordOptionsModel.Protocol = core.StringPtr("udp")
				createResourceRecordOptionsModel.XCorrelationID = core.StringPtr("testString")
				createResourceRecordOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := dnsSvcsService.CreateResourceRecord(createResourceRecordOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				dnsSvcsService.EnableRetries(0, 0)
				result, response, operationErr = dnsSvcsService.CreateResourceRecord(createResourceRecordOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`CreateResourceRecord(createResourceRecordOptions *CreateResourceRecordOptions)`, func() {
		createResourceRecordPath := "/instances/testString/dnszones/testString/resource_records"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createResourceRecordPath))
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

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "SRV:5365b73c-ce6f-4d6f-ad9f-d9c131b26370", "created_on": "2019-01-01T05:20:00.12345Z", "modified_on": "2019-01-01T05:20:00.12345Z", "name": "_sip._udp.test.example.com", "type": "SRV", "ttl": 120, "rdata": {"anyKey": "anyValue"}, "service": "_sip", "protocol": "udp"}`)
				}))
			})
			It(`Invoke CreateResourceRecord successfully`, func() {
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsSvcsService).ToNot(BeNil())
				dnsSvcsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := dnsSvcsService.CreateResourceRecord(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ResourceRecordInputRdataRdataARecord model
				resourceRecordInputRdataModel := new(dnssvcsv1.ResourceRecordInputRdataRdataARecord)
				resourceRecordInputRdataModel.Ip = core.StringPtr("10.110.201.214")

				// Construct an instance of the CreateResourceRecordOptions model
				createResourceRecordOptionsModel := new(dnssvcsv1.CreateResourceRecordOptions)
				createResourceRecordOptionsModel.InstanceID = core.StringPtr("testString")
				createResourceRecordOptionsModel.DnszoneID = core.StringPtr("testString")
				createResourceRecordOptionsModel.Name = core.StringPtr("test.example.com")
				createResourceRecordOptionsModel.Type = core.StringPtr("SRV")
				createResourceRecordOptionsModel.Rdata = resourceRecordInputRdataModel
				createResourceRecordOptionsModel.TTL = core.Int64Ptr(int64(120))
				createResourceRecordOptionsModel.Service = core.StringPtr("_sip")
				createResourceRecordOptionsModel.Protocol = core.StringPtr("udp")
				createResourceRecordOptionsModel.XCorrelationID = core.StringPtr("testString")
				createResourceRecordOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = dnsSvcsService.CreateResourceRecord(createResourceRecordOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = dnsSvcsService.CreateResourceRecordWithContext(ctx, createResourceRecordOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				dnsSvcsService.DisableRetries()
				result, response, operationErr = dnsSvcsService.CreateResourceRecord(createResourceRecordOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = dnsSvcsService.CreateResourceRecordWithContext(ctx, createResourceRecordOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke CreateResourceRecord with error: Operation validation and request error`, func() {
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsSvcsService).ToNot(BeNil())

				// Construct an instance of the ResourceRecordInputRdataRdataARecord model
				resourceRecordInputRdataModel := new(dnssvcsv1.ResourceRecordInputRdataRdataARecord)
				resourceRecordInputRdataModel.Ip = core.StringPtr("10.110.201.214")

				// Construct an instance of the CreateResourceRecordOptions model
				createResourceRecordOptionsModel := new(dnssvcsv1.CreateResourceRecordOptions)
				createResourceRecordOptionsModel.InstanceID = core.StringPtr("testString")
				createResourceRecordOptionsModel.DnszoneID = core.StringPtr("testString")
				createResourceRecordOptionsModel.Name = core.StringPtr("test.example.com")
				createResourceRecordOptionsModel.Type = core.StringPtr("SRV")
				createResourceRecordOptionsModel.Rdata = resourceRecordInputRdataModel
				createResourceRecordOptionsModel.TTL = core.Int64Ptr(int64(120))
				createResourceRecordOptionsModel.Service = core.StringPtr("_sip")
				createResourceRecordOptionsModel.Protocol = core.StringPtr("udp")
				createResourceRecordOptionsModel.XCorrelationID = core.StringPtr("testString")
				createResourceRecordOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := dnsSvcsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := dnsSvcsService.CreateResourceRecord(createResourceRecordOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateResourceRecordOptions model with no property values
				createResourceRecordOptionsModelNew := new(dnssvcsv1.CreateResourceRecordOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = dnsSvcsService.CreateResourceRecord(createResourceRecordOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`DeleteResourceRecord(deleteResourceRecordOptions *DeleteResourceRecordOptions)`, func() {
		deleteResourceRecordPath := "/instances/testString/dnszones/testString/resource_records/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteResourceRecordPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteResourceRecord successfully`, func() {
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsSvcsService).ToNot(BeNil())
				dnsSvcsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				response, operationErr := dnsSvcsService.DeleteResourceRecord(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteResourceRecordOptions model
				deleteResourceRecordOptionsModel := new(dnssvcsv1.DeleteResourceRecordOptions)
				deleteResourceRecordOptionsModel.InstanceID = core.StringPtr("testString")
				deleteResourceRecordOptionsModel.DnszoneID = core.StringPtr("testString")
				deleteResourceRecordOptionsModel.RecordID = core.StringPtr("testString")
				deleteResourceRecordOptionsModel.XCorrelationID = core.StringPtr("testString")
				deleteResourceRecordOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = dnsSvcsService.DeleteResourceRecord(deleteResourceRecordOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Disable retries and test again
				dnsSvcsService.DisableRetries()
				response, operationErr = dnsSvcsService.DeleteResourceRecord(deleteResourceRecordOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteResourceRecord with error: Operation validation and request error`, func() {
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsSvcsService).ToNot(BeNil())

				// Construct an instance of the DeleteResourceRecordOptions model
				deleteResourceRecordOptionsModel := new(dnssvcsv1.DeleteResourceRecordOptions)
				deleteResourceRecordOptionsModel.InstanceID = core.StringPtr("testString")
				deleteResourceRecordOptionsModel.DnszoneID = core.StringPtr("testString")
				deleteResourceRecordOptionsModel.RecordID = core.StringPtr("testString")
				deleteResourceRecordOptionsModel.XCorrelationID = core.StringPtr("testString")
				deleteResourceRecordOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := dnsSvcsService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := dnsSvcsService.DeleteResourceRecord(deleteResourceRecordOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteResourceRecordOptions model with no property values
				deleteResourceRecordOptionsModelNew := new(dnssvcsv1.DeleteResourceRecordOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = dnsSvcsService.DeleteResourceRecord(deleteResourceRecordOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetResourceRecord(getResourceRecordOptions *GetResourceRecordOptions) - Operation response error`, func() {
		getResourceRecordPath := "/instances/testString/dnszones/testString/resource_records/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getResourceRecordPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetResourceRecord with error: Operation response processing error`, func() {
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsSvcsService).ToNot(BeNil())

				// Construct an instance of the GetResourceRecordOptions model
				getResourceRecordOptionsModel := new(dnssvcsv1.GetResourceRecordOptions)
				getResourceRecordOptionsModel.InstanceID = core.StringPtr("testString")
				getResourceRecordOptionsModel.DnszoneID = core.StringPtr("testString")
				getResourceRecordOptionsModel.RecordID = core.StringPtr("testString")
				getResourceRecordOptionsModel.XCorrelationID = core.StringPtr("testString")
				getResourceRecordOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := dnsSvcsService.GetResourceRecord(getResourceRecordOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				dnsSvcsService.EnableRetries(0, 0)
				result, response, operationErr = dnsSvcsService.GetResourceRecord(getResourceRecordOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetResourceRecord(getResourceRecordOptions *GetResourceRecordOptions)`, func() {
		getResourceRecordPath := "/instances/testString/dnszones/testString/resource_records/testString"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getResourceRecordPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "SRV:5365b73c-ce6f-4d6f-ad9f-d9c131b26370", "created_on": "2019-01-01T05:20:00.12345Z", "modified_on": "2019-01-01T05:20:00.12345Z", "name": "_sip._udp.test.example.com", "type": "SRV", "ttl": 120, "rdata": {"anyKey": "anyValue"}, "service": "_sip", "protocol": "udp"}`)
				}))
			})
			It(`Invoke GetResourceRecord successfully`, func() {
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsSvcsService).ToNot(BeNil())
				dnsSvcsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := dnsSvcsService.GetResourceRecord(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetResourceRecordOptions model
				getResourceRecordOptionsModel := new(dnssvcsv1.GetResourceRecordOptions)
				getResourceRecordOptionsModel.InstanceID = core.StringPtr("testString")
				getResourceRecordOptionsModel.DnszoneID = core.StringPtr("testString")
				getResourceRecordOptionsModel.RecordID = core.StringPtr("testString")
				getResourceRecordOptionsModel.XCorrelationID = core.StringPtr("testString")
				getResourceRecordOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = dnsSvcsService.GetResourceRecord(getResourceRecordOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = dnsSvcsService.GetResourceRecordWithContext(ctx, getResourceRecordOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				dnsSvcsService.DisableRetries()
				result, response, operationErr = dnsSvcsService.GetResourceRecord(getResourceRecordOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = dnsSvcsService.GetResourceRecordWithContext(ctx, getResourceRecordOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke GetResourceRecord with error: Operation validation and request error`, func() {
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsSvcsService).ToNot(BeNil())

				// Construct an instance of the GetResourceRecordOptions model
				getResourceRecordOptionsModel := new(dnssvcsv1.GetResourceRecordOptions)
				getResourceRecordOptionsModel.InstanceID = core.StringPtr("testString")
				getResourceRecordOptionsModel.DnszoneID = core.StringPtr("testString")
				getResourceRecordOptionsModel.RecordID = core.StringPtr("testString")
				getResourceRecordOptionsModel.XCorrelationID = core.StringPtr("testString")
				getResourceRecordOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := dnsSvcsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := dnsSvcsService.GetResourceRecord(getResourceRecordOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetResourceRecordOptions model with no property values
				getResourceRecordOptionsModelNew := new(dnssvcsv1.GetResourceRecordOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = dnsSvcsService.GetResourceRecord(getResourceRecordOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateResourceRecord(updateResourceRecordOptions *UpdateResourceRecordOptions) - Operation response error`, func() {
		updateResourceRecordPath := "/instances/testString/dnszones/testString/resource_records/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateResourceRecordPath))
					Expect(req.Method).To(Equal("PUT"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateResourceRecord with error: Operation response processing error`, func() {
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsSvcsService).ToNot(BeNil())

				// Construct an instance of the ResourceRecordUpdateInputRdataRdataARecord model
				resourceRecordUpdateInputRdataModel := new(dnssvcsv1.ResourceRecordUpdateInputRdataRdataARecord)
				resourceRecordUpdateInputRdataModel.Ip = core.StringPtr("10.110.201.214")

				// Construct an instance of the UpdateResourceRecordOptions model
				updateResourceRecordOptionsModel := new(dnssvcsv1.UpdateResourceRecordOptions)
				updateResourceRecordOptionsModel.InstanceID = core.StringPtr("testString")
				updateResourceRecordOptionsModel.DnszoneID = core.StringPtr("testString")
				updateResourceRecordOptionsModel.RecordID = core.StringPtr("testString")
				updateResourceRecordOptionsModel.Name = core.StringPtr("test.example.com")
				updateResourceRecordOptionsModel.Rdata = resourceRecordUpdateInputRdataModel
				updateResourceRecordOptionsModel.TTL = core.Int64Ptr(int64(120))
				updateResourceRecordOptionsModel.Service = core.StringPtr("_sip")
				updateResourceRecordOptionsModel.Protocol = core.StringPtr("udp")
				updateResourceRecordOptionsModel.XCorrelationID = core.StringPtr("testString")
				updateResourceRecordOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := dnsSvcsService.UpdateResourceRecord(updateResourceRecordOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				dnsSvcsService.EnableRetries(0, 0)
				result, response, operationErr = dnsSvcsService.UpdateResourceRecord(updateResourceRecordOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UpdateResourceRecord(updateResourceRecordOptions *UpdateResourceRecordOptions)`, func() {
		updateResourceRecordPath := "/instances/testString/dnszones/testString/resource_records/testString"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateResourceRecordPath))
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

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "SRV:5365b73c-ce6f-4d6f-ad9f-d9c131b26370", "created_on": "2019-01-01T05:20:00.12345Z", "modified_on": "2019-01-01T05:20:00.12345Z", "name": "_sip._udp.test.example.com", "type": "SRV", "ttl": 120, "rdata": {"anyKey": "anyValue"}, "service": "_sip", "protocol": "udp"}`)
				}))
			})
			It(`Invoke UpdateResourceRecord successfully`, func() {
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsSvcsService).ToNot(BeNil())
				dnsSvcsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := dnsSvcsService.UpdateResourceRecord(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ResourceRecordUpdateInputRdataRdataARecord model
				resourceRecordUpdateInputRdataModel := new(dnssvcsv1.ResourceRecordUpdateInputRdataRdataARecord)
				resourceRecordUpdateInputRdataModel.Ip = core.StringPtr("10.110.201.214")

				// Construct an instance of the UpdateResourceRecordOptions model
				updateResourceRecordOptionsModel := new(dnssvcsv1.UpdateResourceRecordOptions)
				updateResourceRecordOptionsModel.InstanceID = core.StringPtr("testString")
				updateResourceRecordOptionsModel.DnszoneID = core.StringPtr("testString")
				updateResourceRecordOptionsModel.RecordID = core.StringPtr("testString")
				updateResourceRecordOptionsModel.Name = core.StringPtr("test.example.com")
				updateResourceRecordOptionsModel.Rdata = resourceRecordUpdateInputRdataModel
				updateResourceRecordOptionsModel.TTL = core.Int64Ptr(int64(120))
				updateResourceRecordOptionsModel.Service = core.StringPtr("_sip")
				updateResourceRecordOptionsModel.Protocol = core.StringPtr("udp")
				updateResourceRecordOptionsModel.XCorrelationID = core.StringPtr("testString")
				updateResourceRecordOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = dnsSvcsService.UpdateResourceRecord(updateResourceRecordOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = dnsSvcsService.UpdateResourceRecordWithContext(ctx, updateResourceRecordOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				dnsSvcsService.DisableRetries()
				result, response, operationErr = dnsSvcsService.UpdateResourceRecord(updateResourceRecordOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = dnsSvcsService.UpdateResourceRecordWithContext(ctx, updateResourceRecordOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke UpdateResourceRecord with error: Operation validation and request error`, func() {
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsSvcsService).ToNot(BeNil())

				// Construct an instance of the ResourceRecordUpdateInputRdataRdataARecord model
				resourceRecordUpdateInputRdataModel := new(dnssvcsv1.ResourceRecordUpdateInputRdataRdataARecord)
				resourceRecordUpdateInputRdataModel.Ip = core.StringPtr("10.110.201.214")

				// Construct an instance of the UpdateResourceRecordOptions model
				updateResourceRecordOptionsModel := new(dnssvcsv1.UpdateResourceRecordOptions)
				updateResourceRecordOptionsModel.InstanceID = core.StringPtr("testString")
				updateResourceRecordOptionsModel.DnszoneID = core.StringPtr("testString")
				updateResourceRecordOptionsModel.RecordID = core.StringPtr("testString")
				updateResourceRecordOptionsModel.Name = core.StringPtr("test.example.com")
				updateResourceRecordOptionsModel.Rdata = resourceRecordUpdateInputRdataModel
				updateResourceRecordOptionsModel.TTL = core.Int64Ptr(int64(120))
				updateResourceRecordOptionsModel.Service = core.StringPtr("_sip")
				updateResourceRecordOptionsModel.Protocol = core.StringPtr("udp")
				updateResourceRecordOptionsModel.XCorrelationID = core.StringPtr("testString")
				updateResourceRecordOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := dnsSvcsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := dnsSvcsService.UpdateResourceRecord(updateResourceRecordOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateResourceRecordOptions model with no property values
				updateResourceRecordOptionsModelNew := new(dnssvcsv1.UpdateResourceRecordOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = dnsSvcsService.UpdateResourceRecord(updateResourceRecordOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ExportResourceRecords(exportResourceRecordsOptions *ExportResourceRecordsOptions)`, func() {
		exportResourceRecordsPath := "/instances/testString/dnszones/testString/export_resource_records"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(exportResourceRecordsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "text/plain; charset=utf-8")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `This is a mock binary response.`)
				}))
			})
			It(`Invoke ExportResourceRecords successfully`, func() {
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsSvcsService).ToNot(BeNil())
				dnsSvcsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := dnsSvcsService.ExportResourceRecords(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ExportResourceRecordsOptions model
				exportResourceRecordsOptionsModel := new(dnssvcsv1.ExportResourceRecordsOptions)
				exportResourceRecordsOptionsModel.InstanceID = core.StringPtr("testString")
				exportResourceRecordsOptionsModel.DnszoneID = core.StringPtr("testString")
				exportResourceRecordsOptionsModel.XCorrelationID = core.StringPtr("testString")
				exportResourceRecordsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = dnsSvcsService.ExportResourceRecords(exportResourceRecordsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = dnsSvcsService.ExportResourceRecordsWithContext(ctx, exportResourceRecordsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				dnsSvcsService.DisableRetries()
				result, response, operationErr = dnsSvcsService.ExportResourceRecords(exportResourceRecordsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = dnsSvcsService.ExportResourceRecordsWithContext(ctx, exportResourceRecordsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke ExportResourceRecords with error: Operation validation and request error`, func() {
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsSvcsService).ToNot(BeNil())

				// Construct an instance of the ExportResourceRecordsOptions model
				exportResourceRecordsOptionsModel := new(dnssvcsv1.ExportResourceRecordsOptions)
				exportResourceRecordsOptionsModel.InstanceID = core.StringPtr("testString")
				exportResourceRecordsOptionsModel.DnszoneID = core.StringPtr("testString")
				exportResourceRecordsOptionsModel.XCorrelationID = core.StringPtr("testString")
				exportResourceRecordsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := dnsSvcsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := dnsSvcsService.ExportResourceRecords(exportResourceRecordsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ExportResourceRecordsOptions model with no property values
				exportResourceRecordsOptionsModelNew := new(dnssvcsv1.ExportResourceRecordsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = dnsSvcsService.ExportResourceRecords(exportResourceRecordsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ImportResourceRecords(importResourceRecordsOptions *ImportResourceRecordsOptions) - Operation response error`, func() {
		importResourceRecordsPath := "/instances/testString/dnszones/testString/import_resource_records"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(importResourceRecordsPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ImportResourceRecords with error: Operation response processing error`, func() {
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsSvcsService).ToNot(BeNil())

				// Construct an instance of the ImportResourceRecordsOptions model
				importResourceRecordsOptionsModel := new(dnssvcsv1.ImportResourceRecordsOptions)
				importResourceRecordsOptionsModel.InstanceID = core.StringPtr("testString")
				importResourceRecordsOptionsModel.DnszoneID = core.StringPtr("testString")
				importResourceRecordsOptionsModel.File = CreateMockReader("This is a mock file.")
				importResourceRecordsOptionsModel.FileContentType = core.StringPtr("testString")
				importResourceRecordsOptionsModel.XCorrelationID = core.StringPtr("testString")
				importResourceRecordsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := dnsSvcsService.ImportResourceRecords(importResourceRecordsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				dnsSvcsService.EnableRetries(0, 0)
				result, response, operationErr = dnsSvcsService.ImportResourceRecords(importResourceRecordsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ImportResourceRecords(importResourceRecordsOptions *ImportResourceRecordsOptions)`, func() {
		importResourceRecordsPath := "/instances/testString/dnszones/testString/import_resource_records"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(importResourceRecordsPath))
					Expect(req.Method).To(Equal("POST"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"total_records_parsed": 10, "records_added": 2, "records_failed": 0, "records_added_by_type": {"A": 10, "AAAA": 10, "CNAME": 10, "SRV": 10, "TXT": 10, "MX": 10, "PTR": 10}, "records_failed_by_type": {"A": 10, "AAAA": 10, "CNAME": 10, "SRV": 10, "TXT": 10, "MX": 10, "PTR": 10}, "messages": [{"code": "conflict", "message": "A type record conflict with other records"}], "errors": [{"resource_record": "test.example.com A 1.1.1.1", "error": {"code": "internal_server_error", "message": "An internal error occurred. Try again later."}}]}`)
				}))
			})
			It(`Invoke ImportResourceRecords successfully`, func() {
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsSvcsService).ToNot(BeNil())
				dnsSvcsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := dnsSvcsService.ImportResourceRecords(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ImportResourceRecordsOptions model
				importResourceRecordsOptionsModel := new(dnssvcsv1.ImportResourceRecordsOptions)
				importResourceRecordsOptionsModel.InstanceID = core.StringPtr("testString")
				importResourceRecordsOptionsModel.DnszoneID = core.StringPtr("testString")
				importResourceRecordsOptionsModel.File = CreateMockReader("This is a mock file.")
				importResourceRecordsOptionsModel.FileContentType = core.StringPtr("testString")
				importResourceRecordsOptionsModel.XCorrelationID = core.StringPtr("testString")
				importResourceRecordsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = dnsSvcsService.ImportResourceRecords(importResourceRecordsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = dnsSvcsService.ImportResourceRecordsWithContext(ctx, importResourceRecordsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				dnsSvcsService.DisableRetries()
				result, response, operationErr = dnsSvcsService.ImportResourceRecords(importResourceRecordsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = dnsSvcsService.ImportResourceRecordsWithContext(ctx, importResourceRecordsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke ImportResourceRecords with error: Param validation error`, func() {
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
					URL:  testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsSvcsService).ToNot(BeNil())

				// Construct an instance of the ImportResourceRecordsOptions model
				importResourceRecordsOptionsModel := new(dnssvcsv1.ImportResourceRecordsOptions)
				// Invoke operation with invalid options model (negative test)
				result, response, operationErr := dnsSvcsService.ImportResourceRecords(importResourceRecordsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			It(`Invoke ImportResourceRecords with error: Operation validation and request error`, func() {
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsSvcsService).ToNot(BeNil())

				// Construct an instance of the ImportResourceRecordsOptions model
				importResourceRecordsOptionsModel := new(dnssvcsv1.ImportResourceRecordsOptions)
				importResourceRecordsOptionsModel.InstanceID = core.StringPtr("testString")
				importResourceRecordsOptionsModel.DnszoneID = core.StringPtr("testString")
				importResourceRecordsOptionsModel.File = CreateMockReader("This is a mock file.")
				importResourceRecordsOptionsModel.FileContentType = core.StringPtr("testString")
				importResourceRecordsOptionsModel.XCorrelationID = core.StringPtr("testString")
				importResourceRecordsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := dnsSvcsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := dnsSvcsService.ImportResourceRecords(importResourceRecordsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ImportResourceRecordsOptions model with no property values
				importResourceRecordsOptionsModelNew := new(dnssvcsv1.ImportResourceRecordsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = dnsSvcsService.ImportResourceRecords(importResourceRecordsOptionsModelNew)
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
		It(`Instantiate service client`, func() {
			dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(dnsSvcsService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(dnsSvcsService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
				URL: "https://dnssvcsv1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(dnsSvcsService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"DNS_SVCS_URL": "https://dnssvcsv1/api",
				"DNS_SVCS_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1UsingExternalConfig(&dnssvcsv1.DnsSvcsV1Options{
				})
				Expect(dnsSvcsService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := dnsSvcsService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != dnsSvcsService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(dnsSvcsService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(dnsSvcsService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1UsingExternalConfig(&dnssvcsv1.DnsSvcsV1Options{
					URL: "https://testService/api",
				})
				Expect(dnsSvcsService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(dnsSvcsService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := dnsSvcsService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != dnsSvcsService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(dnsSvcsService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(dnsSvcsService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1UsingExternalConfig(&dnssvcsv1.DnsSvcsV1Options{
				})
				err := dnsSvcsService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(dnsSvcsService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(dnsSvcsService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := dnsSvcsService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != dnsSvcsService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(dnsSvcsService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(dnsSvcsService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"DNS_SVCS_URL": "https://dnssvcsv1/api",
				"DNS_SVCS_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1UsingExternalConfig(&dnssvcsv1.DnsSvcsV1Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(dnsSvcsService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"DNS_SVCS_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1UsingExternalConfig(&dnssvcsv1.DnsSvcsV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(dnsSvcsService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = dnssvcsv1.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`ListPermittedNetworks(listPermittedNetworksOptions *ListPermittedNetworksOptions) - Operation response error`, func() {
		listPermittedNetworksPath := "/instances/testString/dnszones/testString/permitted_networks"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listPermittedNetworksPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(38))}))

					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(200))}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListPermittedNetworks with error: Operation response processing error`, func() {
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsSvcsService).ToNot(BeNil())

				// Construct an instance of the ListPermittedNetworksOptions model
				listPermittedNetworksOptionsModel := new(dnssvcsv1.ListPermittedNetworksOptions)
				listPermittedNetworksOptionsModel.InstanceID = core.StringPtr("testString")
				listPermittedNetworksOptionsModel.DnszoneID = core.StringPtr("testString")
				listPermittedNetworksOptionsModel.XCorrelationID = core.StringPtr("testString")
				listPermittedNetworksOptionsModel.Offset = core.Int64Ptr(int64(38))
				listPermittedNetworksOptionsModel.Limit = core.Int64Ptr(int64(200))
				listPermittedNetworksOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := dnsSvcsService.ListPermittedNetworks(listPermittedNetworksOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				dnsSvcsService.EnableRetries(0, 0)
				result, response, operationErr = dnsSvcsService.ListPermittedNetworks(listPermittedNetworksOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ListPermittedNetworks(listPermittedNetworksOptions *ListPermittedNetworksOptions)`, func() {
		listPermittedNetworksPath := "/instances/testString/dnszones/testString/permitted_networks"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listPermittedNetworksPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(38))}))

					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(200))}))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"permitted_networks": [{"id": "fecd0173-3919-456b-b202-3029dfa1b0f7", "created_on": "2019-01-01T05:20:00.12345Z", "modified_on": "2019-01-01T05:20:00.12345Z", "permitted_network": {"vpc_crn": "crn:v1:bluemix:public:is:eu-de:a/bcf1865e99742d38d2d5fc3fb80a5496::vpc:6e6cc326-04d1-4c99-a289-efb3ae4193d6"}, "type": "vpc", "state": "ACTIVE"}], "offset": 0, "limit": 10, "total_count": 200, "first": {"href": "https://api.dns-svcs.cloud.ibm.com/v1/instances/434f6c3e-6014-4124-a61d-2e910bca19b1/dnszones/example.com:d04d3a7a-7f6d-47d4-b811-08c5478fa1a4/resource_records?limit=20"}, "next": {"href": "https://api.dns-svcs.cloud.ibm.com/v1/instances/434f6c3e-6014-4124-a61d-2e910bca19b1/dnszones/example.com:d04d3a7a-7f6d-47d4-b811-08c5478fa1a4/resource_records?offset=20&limit=20"}}`)
				}))
			})
			It(`Invoke ListPermittedNetworks successfully`, func() {
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsSvcsService).ToNot(BeNil())
				dnsSvcsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := dnsSvcsService.ListPermittedNetworks(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListPermittedNetworksOptions model
				listPermittedNetworksOptionsModel := new(dnssvcsv1.ListPermittedNetworksOptions)
				listPermittedNetworksOptionsModel.InstanceID = core.StringPtr("testString")
				listPermittedNetworksOptionsModel.DnszoneID = core.StringPtr("testString")
				listPermittedNetworksOptionsModel.XCorrelationID = core.StringPtr("testString")
				listPermittedNetworksOptionsModel.Offset = core.Int64Ptr(int64(38))
				listPermittedNetworksOptionsModel.Limit = core.Int64Ptr(int64(200))
				listPermittedNetworksOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = dnsSvcsService.ListPermittedNetworks(listPermittedNetworksOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = dnsSvcsService.ListPermittedNetworksWithContext(ctx, listPermittedNetworksOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				dnsSvcsService.DisableRetries()
				result, response, operationErr = dnsSvcsService.ListPermittedNetworks(listPermittedNetworksOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = dnsSvcsService.ListPermittedNetworksWithContext(ctx, listPermittedNetworksOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke ListPermittedNetworks with error: Operation validation and request error`, func() {
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsSvcsService).ToNot(BeNil())

				// Construct an instance of the ListPermittedNetworksOptions model
				listPermittedNetworksOptionsModel := new(dnssvcsv1.ListPermittedNetworksOptions)
				listPermittedNetworksOptionsModel.InstanceID = core.StringPtr("testString")
				listPermittedNetworksOptionsModel.DnszoneID = core.StringPtr("testString")
				listPermittedNetworksOptionsModel.XCorrelationID = core.StringPtr("testString")
				listPermittedNetworksOptionsModel.Offset = core.Int64Ptr(int64(38))
				listPermittedNetworksOptionsModel.Limit = core.Int64Ptr(int64(200))
				listPermittedNetworksOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := dnsSvcsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := dnsSvcsService.ListPermittedNetworks(listPermittedNetworksOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListPermittedNetworksOptions model with no property values
				listPermittedNetworksOptionsModelNew := new(dnssvcsv1.ListPermittedNetworksOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = dnsSvcsService.ListPermittedNetworks(listPermittedNetworksOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreatePermittedNetwork(createPermittedNetworkOptions *CreatePermittedNetworkOptions) - Operation response error`, func() {
		createPermittedNetworkPath := "/instances/testString/dnszones/testString/permitted_networks"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createPermittedNetworkPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreatePermittedNetwork with error: Operation response processing error`, func() {
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsSvcsService).ToNot(BeNil())

				// Construct an instance of the PermittedNetworkVpc model
				permittedNetworkVpcModel := new(dnssvcsv1.PermittedNetworkVpc)
				permittedNetworkVpcModel.VpcCrn = core.StringPtr("crn:v1:bluemix:public:is:eu-de:a/bcf1865e99742d38d2d5fc3fb80a5496::vpc:6e6cc326-04d1-4c99-a289-efb3ae4193d6")

				// Construct an instance of the CreatePermittedNetworkOptions model
				createPermittedNetworkOptionsModel := new(dnssvcsv1.CreatePermittedNetworkOptions)
				createPermittedNetworkOptionsModel.InstanceID = core.StringPtr("testString")
				createPermittedNetworkOptionsModel.DnszoneID = core.StringPtr("testString")
				createPermittedNetworkOptionsModel.Type = core.StringPtr("vpc")
				createPermittedNetworkOptionsModel.PermittedNetwork = permittedNetworkVpcModel
				createPermittedNetworkOptionsModel.XCorrelationID = core.StringPtr("testString")
				createPermittedNetworkOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := dnsSvcsService.CreatePermittedNetwork(createPermittedNetworkOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				dnsSvcsService.EnableRetries(0, 0)
				result, response, operationErr = dnsSvcsService.CreatePermittedNetwork(createPermittedNetworkOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`CreatePermittedNetwork(createPermittedNetworkOptions *CreatePermittedNetworkOptions)`, func() {
		createPermittedNetworkPath := "/instances/testString/dnszones/testString/permitted_networks"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createPermittedNetworkPath))
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

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "fecd0173-3919-456b-b202-3029dfa1b0f7", "created_on": "2019-01-01T05:20:00.12345Z", "modified_on": "2019-01-01T05:20:00.12345Z", "permitted_network": {"vpc_crn": "crn:v1:bluemix:public:is:eu-de:a/bcf1865e99742d38d2d5fc3fb80a5496::vpc:6e6cc326-04d1-4c99-a289-efb3ae4193d6"}, "type": "vpc", "state": "ACTIVE"}`)
				}))
			})
			It(`Invoke CreatePermittedNetwork successfully`, func() {
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsSvcsService).ToNot(BeNil())
				dnsSvcsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := dnsSvcsService.CreatePermittedNetwork(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the PermittedNetworkVpc model
				permittedNetworkVpcModel := new(dnssvcsv1.PermittedNetworkVpc)
				permittedNetworkVpcModel.VpcCrn = core.StringPtr("crn:v1:bluemix:public:is:eu-de:a/bcf1865e99742d38d2d5fc3fb80a5496::vpc:6e6cc326-04d1-4c99-a289-efb3ae4193d6")

				// Construct an instance of the CreatePermittedNetworkOptions model
				createPermittedNetworkOptionsModel := new(dnssvcsv1.CreatePermittedNetworkOptions)
				createPermittedNetworkOptionsModel.InstanceID = core.StringPtr("testString")
				createPermittedNetworkOptionsModel.DnszoneID = core.StringPtr("testString")
				createPermittedNetworkOptionsModel.Type = core.StringPtr("vpc")
				createPermittedNetworkOptionsModel.PermittedNetwork = permittedNetworkVpcModel
				createPermittedNetworkOptionsModel.XCorrelationID = core.StringPtr("testString")
				createPermittedNetworkOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = dnsSvcsService.CreatePermittedNetwork(createPermittedNetworkOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = dnsSvcsService.CreatePermittedNetworkWithContext(ctx, createPermittedNetworkOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				dnsSvcsService.DisableRetries()
				result, response, operationErr = dnsSvcsService.CreatePermittedNetwork(createPermittedNetworkOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = dnsSvcsService.CreatePermittedNetworkWithContext(ctx, createPermittedNetworkOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke CreatePermittedNetwork with error: Operation validation and request error`, func() {
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsSvcsService).ToNot(BeNil())

				// Construct an instance of the PermittedNetworkVpc model
				permittedNetworkVpcModel := new(dnssvcsv1.PermittedNetworkVpc)
				permittedNetworkVpcModel.VpcCrn = core.StringPtr("crn:v1:bluemix:public:is:eu-de:a/bcf1865e99742d38d2d5fc3fb80a5496::vpc:6e6cc326-04d1-4c99-a289-efb3ae4193d6")

				// Construct an instance of the CreatePermittedNetworkOptions model
				createPermittedNetworkOptionsModel := new(dnssvcsv1.CreatePermittedNetworkOptions)
				createPermittedNetworkOptionsModel.InstanceID = core.StringPtr("testString")
				createPermittedNetworkOptionsModel.DnszoneID = core.StringPtr("testString")
				createPermittedNetworkOptionsModel.Type = core.StringPtr("vpc")
				createPermittedNetworkOptionsModel.PermittedNetwork = permittedNetworkVpcModel
				createPermittedNetworkOptionsModel.XCorrelationID = core.StringPtr("testString")
				createPermittedNetworkOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := dnsSvcsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := dnsSvcsService.CreatePermittedNetwork(createPermittedNetworkOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreatePermittedNetworkOptions model with no property values
				createPermittedNetworkOptionsModelNew := new(dnssvcsv1.CreatePermittedNetworkOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = dnsSvcsService.CreatePermittedNetwork(createPermittedNetworkOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeletePermittedNetwork(deletePermittedNetworkOptions *DeletePermittedNetworkOptions) - Operation response error`, func() {
		deletePermittedNetworkPath := "/instances/testString/dnszones/testString/permitted_networks/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deletePermittedNetworkPath))
					Expect(req.Method).To(Equal("DELETE"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke DeletePermittedNetwork with error: Operation response processing error`, func() {
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsSvcsService).ToNot(BeNil())

				// Construct an instance of the DeletePermittedNetworkOptions model
				deletePermittedNetworkOptionsModel := new(dnssvcsv1.DeletePermittedNetworkOptions)
				deletePermittedNetworkOptionsModel.InstanceID = core.StringPtr("testString")
				deletePermittedNetworkOptionsModel.DnszoneID = core.StringPtr("testString")
				deletePermittedNetworkOptionsModel.PermittedNetworkID = core.StringPtr("testString")
				deletePermittedNetworkOptionsModel.XCorrelationID = core.StringPtr("testString")
				deletePermittedNetworkOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := dnsSvcsService.DeletePermittedNetwork(deletePermittedNetworkOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				dnsSvcsService.EnableRetries(0, 0)
				result, response, operationErr = dnsSvcsService.DeletePermittedNetwork(deletePermittedNetworkOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`DeletePermittedNetwork(deletePermittedNetworkOptions *DeletePermittedNetworkOptions)`, func() {
		deletePermittedNetworkPath := "/instances/testString/dnszones/testString/permitted_networks/testString"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deletePermittedNetworkPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"id": "fecd0173-3919-456b-b202-3029dfa1b0f7", "created_on": "2019-01-01T05:20:00.12345Z", "modified_on": "2019-01-01T05:20:00.12345Z", "permitted_network": {"vpc_crn": "crn:v1:bluemix:public:is:eu-de:a/bcf1865e99742d38d2d5fc3fb80a5496::vpc:6e6cc326-04d1-4c99-a289-efb3ae4193d6"}, "type": "vpc", "state": "ACTIVE"}`)
				}))
			})
			It(`Invoke DeletePermittedNetwork successfully`, func() {
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsSvcsService).ToNot(BeNil())
				dnsSvcsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := dnsSvcsService.DeletePermittedNetwork(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DeletePermittedNetworkOptions model
				deletePermittedNetworkOptionsModel := new(dnssvcsv1.DeletePermittedNetworkOptions)
				deletePermittedNetworkOptionsModel.InstanceID = core.StringPtr("testString")
				deletePermittedNetworkOptionsModel.DnszoneID = core.StringPtr("testString")
				deletePermittedNetworkOptionsModel.PermittedNetworkID = core.StringPtr("testString")
				deletePermittedNetworkOptionsModel.XCorrelationID = core.StringPtr("testString")
				deletePermittedNetworkOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = dnsSvcsService.DeletePermittedNetwork(deletePermittedNetworkOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = dnsSvcsService.DeletePermittedNetworkWithContext(ctx, deletePermittedNetworkOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				dnsSvcsService.DisableRetries()
				result, response, operationErr = dnsSvcsService.DeletePermittedNetwork(deletePermittedNetworkOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = dnsSvcsService.DeletePermittedNetworkWithContext(ctx, deletePermittedNetworkOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke DeletePermittedNetwork with error: Operation validation and request error`, func() {
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsSvcsService).ToNot(BeNil())

				// Construct an instance of the DeletePermittedNetworkOptions model
				deletePermittedNetworkOptionsModel := new(dnssvcsv1.DeletePermittedNetworkOptions)
				deletePermittedNetworkOptionsModel.InstanceID = core.StringPtr("testString")
				deletePermittedNetworkOptionsModel.DnszoneID = core.StringPtr("testString")
				deletePermittedNetworkOptionsModel.PermittedNetworkID = core.StringPtr("testString")
				deletePermittedNetworkOptionsModel.XCorrelationID = core.StringPtr("testString")
				deletePermittedNetworkOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := dnsSvcsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := dnsSvcsService.DeletePermittedNetwork(deletePermittedNetworkOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the DeletePermittedNetworkOptions model with no property values
				deletePermittedNetworkOptionsModelNew := new(dnssvcsv1.DeletePermittedNetworkOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = dnsSvcsService.DeletePermittedNetwork(deletePermittedNetworkOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetPermittedNetwork(getPermittedNetworkOptions *GetPermittedNetworkOptions) - Operation response error`, func() {
		getPermittedNetworkPath := "/instances/testString/dnszones/testString/permitted_networks/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getPermittedNetworkPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetPermittedNetwork with error: Operation response processing error`, func() {
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsSvcsService).ToNot(BeNil())

				// Construct an instance of the GetPermittedNetworkOptions model
				getPermittedNetworkOptionsModel := new(dnssvcsv1.GetPermittedNetworkOptions)
				getPermittedNetworkOptionsModel.InstanceID = core.StringPtr("testString")
				getPermittedNetworkOptionsModel.DnszoneID = core.StringPtr("testString")
				getPermittedNetworkOptionsModel.PermittedNetworkID = core.StringPtr("testString")
				getPermittedNetworkOptionsModel.XCorrelationID = core.StringPtr("testString")
				getPermittedNetworkOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := dnsSvcsService.GetPermittedNetwork(getPermittedNetworkOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				dnsSvcsService.EnableRetries(0, 0)
				result, response, operationErr = dnsSvcsService.GetPermittedNetwork(getPermittedNetworkOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetPermittedNetwork(getPermittedNetworkOptions *GetPermittedNetworkOptions)`, func() {
		getPermittedNetworkPath := "/instances/testString/dnszones/testString/permitted_networks/testString"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getPermittedNetworkPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "fecd0173-3919-456b-b202-3029dfa1b0f7", "created_on": "2019-01-01T05:20:00.12345Z", "modified_on": "2019-01-01T05:20:00.12345Z", "permitted_network": {"vpc_crn": "crn:v1:bluemix:public:is:eu-de:a/bcf1865e99742d38d2d5fc3fb80a5496::vpc:6e6cc326-04d1-4c99-a289-efb3ae4193d6"}, "type": "vpc", "state": "ACTIVE"}`)
				}))
			})
			It(`Invoke GetPermittedNetwork successfully`, func() {
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsSvcsService).ToNot(BeNil())
				dnsSvcsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := dnsSvcsService.GetPermittedNetwork(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetPermittedNetworkOptions model
				getPermittedNetworkOptionsModel := new(dnssvcsv1.GetPermittedNetworkOptions)
				getPermittedNetworkOptionsModel.InstanceID = core.StringPtr("testString")
				getPermittedNetworkOptionsModel.DnszoneID = core.StringPtr("testString")
				getPermittedNetworkOptionsModel.PermittedNetworkID = core.StringPtr("testString")
				getPermittedNetworkOptionsModel.XCorrelationID = core.StringPtr("testString")
				getPermittedNetworkOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = dnsSvcsService.GetPermittedNetwork(getPermittedNetworkOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = dnsSvcsService.GetPermittedNetworkWithContext(ctx, getPermittedNetworkOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				dnsSvcsService.DisableRetries()
				result, response, operationErr = dnsSvcsService.GetPermittedNetwork(getPermittedNetworkOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = dnsSvcsService.GetPermittedNetworkWithContext(ctx, getPermittedNetworkOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke GetPermittedNetwork with error: Operation validation and request error`, func() {
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsSvcsService).ToNot(BeNil())

				// Construct an instance of the GetPermittedNetworkOptions model
				getPermittedNetworkOptionsModel := new(dnssvcsv1.GetPermittedNetworkOptions)
				getPermittedNetworkOptionsModel.InstanceID = core.StringPtr("testString")
				getPermittedNetworkOptionsModel.DnszoneID = core.StringPtr("testString")
				getPermittedNetworkOptionsModel.PermittedNetworkID = core.StringPtr("testString")
				getPermittedNetworkOptionsModel.XCorrelationID = core.StringPtr("testString")
				getPermittedNetworkOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := dnsSvcsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := dnsSvcsService.GetPermittedNetwork(getPermittedNetworkOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetPermittedNetworkOptions model with no property values
				getPermittedNetworkOptionsModelNew := new(dnssvcsv1.GetPermittedNetworkOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = dnsSvcsService.GetPermittedNetwork(getPermittedNetworkOptionsModelNew)
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
		It(`Instantiate service client`, func() {
			dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(dnsSvcsService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(dnsSvcsService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
				URL: "https://dnssvcsv1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(dnsSvcsService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"DNS_SVCS_URL": "https://dnssvcsv1/api",
				"DNS_SVCS_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1UsingExternalConfig(&dnssvcsv1.DnsSvcsV1Options{
				})
				Expect(dnsSvcsService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := dnsSvcsService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != dnsSvcsService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(dnsSvcsService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(dnsSvcsService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1UsingExternalConfig(&dnssvcsv1.DnsSvcsV1Options{
					URL: "https://testService/api",
				})
				Expect(dnsSvcsService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(dnsSvcsService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := dnsSvcsService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != dnsSvcsService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(dnsSvcsService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(dnsSvcsService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1UsingExternalConfig(&dnssvcsv1.DnsSvcsV1Options{
				})
				err := dnsSvcsService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(dnsSvcsService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(dnsSvcsService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := dnsSvcsService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != dnsSvcsService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(dnsSvcsService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(dnsSvcsService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"DNS_SVCS_URL": "https://dnssvcsv1/api",
				"DNS_SVCS_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1UsingExternalConfig(&dnssvcsv1.DnsSvcsV1Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(dnsSvcsService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"DNS_SVCS_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1UsingExternalConfig(&dnssvcsv1.DnsSvcsV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(dnsSvcsService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = dnssvcsv1.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`ListLoadBalancers(listLoadBalancersOptions *ListLoadBalancersOptions) - Operation response error`, func() {
		listLoadBalancersPath := "/instances/testString/dnszones/testString/load_balancers"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listLoadBalancersPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListLoadBalancers with error: Operation response processing error`, func() {
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsSvcsService).ToNot(BeNil())

				// Construct an instance of the ListLoadBalancersOptions model
				listLoadBalancersOptionsModel := new(dnssvcsv1.ListLoadBalancersOptions)
				listLoadBalancersOptionsModel.InstanceID = core.StringPtr("testString")
				listLoadBalancersOptionsModel.DnszoneID = core.StringPtr("testString")
				listLoadBalancersOptionsModel.XCorrelationID = core.StringPtr("testString")
				listLoadBalancersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := dnsSvcsService.ListLoadBalancers(listLoadBalancersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				dnsSvcsService.EnableRetries(0, 0)
				result, response, operationErr = dnsSvcsService.ListLoadBalancers(listLoadBalancersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ListLoadBalancers(listLoadBalancersOptions *ListLoadBalancersOptions)`, func() {
		listLoadBalancersPath := "/instances/testString/dnszones/testString/load_balancers"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listLoadBalancersPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"load_balancers": [{"id": "5365b73c-ce6f-4d6f-ad9f-d9c131b26370", "name": "glb.example.com", "description": "Load balancer for glb.example.com.", "enabled": true, "ttl": 120, "health": "DEGRADED", "fallback_pool": "24ccf79a-4ae0-4769-b4c8-17f8f230072e", "default_pools": ["DefaultPools"], "az_pools": [{"availability_zone": "us-south-1", "pools": ["0fc0bb7c-2fab-476e-8b9b-40fa14bf8e3d"]}], "created_on": "2019-01-01T05:20:00.12345Z", "modified_on": "2019-01-01T05:20:00.12345Z"}], "offset": 1, "limit": 20, "count": 1, "total_count": 200, "first": {"href": "https://api.dns-svcs.cloud.ibm.com/v1/instances/434f6c3e-6014-4124-a61d-2e910bca19b1/dnszones/example.com:d04d3a7a-7f6d-47d4-b811-08c5478fa1a4/resource_records?limit=20"}, "next": {"href": "https://api.dns-svcs.cloud.ibm.com/v1/instances/434f6c3e-6014-4124-a61d-2e910bca19b1/dnszones/example.com:d04d3a7a-7f6d-47d4-b811-08c5478fa1a4/resource_records?offset=20&limit=20"}}`)
				}))
			})
			It(`Invoke ListLoadBalancers successfully`, func() {
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsSvcsService).ToNot(BeNil())
				dnsSvcsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := dnsSvcsService.ListLoadBalancers(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListLoadBalancersOptions model
				listLoadBalancersOptionsModel := new(dnssvcsv1.ListLoadBalancersOptions)
				listLoadBalancersOptionsModel.InstanceID = core.StringPtr("testString")
				listLoadBalancersOptionsModel.DnszoneID = core.StringPtr("testString")
				listLoadBalancersOptionsModel.XCorrelationID = core.StringPtr("testString")
				listLoadBalancersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = dnsSvcsService.ListLoadBalancers(listLoadBalancersOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = dnsSvcsService.ListLoadBalancersWithContext(ctx, listLoadBalancersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				dnsSvcsService.DisableRetries()
				result, response, operationErr = dnsSvcsService.ListLoadBalancers(listLoadBalancersOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = dnsSvcsService.ListLoadBalancersWithContext(ctx, listLoadBalancersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke ListLoadBalancers with error: Operation validation and request error`, func() {
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsSvcsService).ToNot(BeNil())

				// Construct an instance of the ListLoadBalancersOptions model
				listLoadBalancersOptionsModel := new(dnssvcsv1.ListLoadBalancersOptions)
				listLoadBalancersOptionsModel.InstanceID = core.StringPtr("testString")
				listLoadBalancersOptionsModel.DnszoneID = core.StringPtr("testString")
				listLoadBalancersOptionsModel.XCorrelationID = core.StringPtr("testString")
				listLoadBalancersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := dnsSvcsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := dnsSvcsService.ListLoadBalancers(listLoadBalancersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListLoadBalancersOptions model with no property values
				listLoadBalancersOptionsModelNew := new(dnssvcsv1.ListLoadBalancersOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = dnsSvcsService.ListLoadBalancers(listLoadBalancersOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateLoadBalancer(createLoadBalancerOptions *CreateLoadBalancerOptions) - Operation response error`, func() {
		createLoadBalancerPath := "/instances/testString/dnszones/testString/load_balancers"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createLoadBalancerPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateLoadBalancer with error: Operation response processing error`, func() {
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsSvcsService).ToNot(BeNil())

				// Construct an instance of the LoadBalancerAzPoolsItem model
				loadBalancerAzPoolsItemModel := new(dnssvcsv1.LoadBalancerAzPoolsItem)
				loadBalancerAzPoolsItemModel.AvailabilityZone = core.StringPtr("us-south-1")
				loadBalancerAzPoolsItemModel.Pools = []string{"0fc0bb7c-2fab-476e-8b9b-40fa14bf8e3d"}

				// Construct an instance of the CreateLoadBalancerOptions model
				createLoadBalancerOptionsModel := new(dnssvcsv1.CreateLoadBalancerOptions)
				createLoadBalancerOptionsModel.InstanceID = core.StringPtr("testString")
				createLoadBalancerOptionsModel.DnszoneID = core.StringPtr("testString")
				createLoadBalancerOptionsModel.Name = core.StringPtr("glb.example.com")
				createLoadBalancerOptionsModel.Description = core.StringPtr("Load balancer for glb.example.com.")
				createLoadBalancerOptionsModel.Enabled = core.BoolPtr(true)
				createLoadBalancerOptionsModel.TTL = core.Int64Ptr(int64(120))
				createLoadBalancerOptionsModel.FallbackPool = core.StringPtr("24ccf79a-4ae0-4769-b4c8-17f8f230072e")
				createLoadBalancerOptionsModel.DefaultPools = []string{"testString"}
				createLoadBalancerOptionsModel.AzPools = []dnssvcsv1.LoadBalancerAzPoolsItem{*loadBalancerAzPoolsItemModel}
				createLoadBalancerOptionsModel.XCorrelationID = core.StringPtr("testString")
				createLoadBalancerOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := dnsSvcsService.CreateLoadBalancer(createLoadBalancerOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				dnsSvcsService.EnableRetries(0, 0)
				result, response, operationErr = dnsSvcsService.CreateLoadBalancer(createLoadBalancerOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`CreateLoadBalancer(createLoadBalancerOptions *CreateLoadBalancerOptions)`, func() {
		createLoadBalancerPath := "/instances/testString/dnszones/testString/load_balancers"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createLoadBalancerPath))
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

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "5365b73c-ce6f-4d6f-ad9f-d9c131b26370", "name": "glb.example.com", "description": "Load balancer for glb.example.com.", "enabled": true, "ttl": 120, "health": "DEGRADED", "fallback_pool": "24ccf79a-4ae0-4769-b4c8-17f8f230072e", "default_pools": ["DefaultPools"], "az_pools": [{"availability_zone": "us-south-1", "pools": ["0fc0bb7c-2fab-476e-8b9b-40fa14bf8e3d"]}], "created_on": "2019-01-01T05:20:00.12345Z", "modified_on": "2019-01-01T05:20:00.12345Z"}`)
				}))
			})
			It(`Invoke CreateLoadBalancer successfully`, func() {
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsSvcsService).ToNot(BeNil())
				dnsSvcsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := dnsSvcsService.CreateLoadBalancer(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the LoadBalancerAzPoolsItem model
				loadBalancerAzPoolsItemModel := new(dnssvcsv1.LoadBalancerAzPoolsItem)
				loadBalancerAzPoolsItemModel.AvailabilityZone = core.StringPtr("us-south-1")
				loadBalancerAzPoolsItemModel.Pools = []string{"0fc0bb7c-2fab-476e-8b9b-40fa14bf8e3d"}

				// Construct an instance of the CreateLoadBalancerOptions model
				createLoadBalancerOptionsModel := new(dnssvcsv1.CreateLoadBalancerOptions)
				createLoadBalancerOptionsModel.InstanceID = core.StringPtr("testString")
				createLoadBalancerOptionsModel.DnszoneID = core.StringPtr("testString")
				createLoadBalancerOptionsModel.Name = core.StringPtr("glb.example.com")
				createLoadBalancerOptionsModel.Description = core.StringPtr("Load balancer for glb.example.com.")
				createLoadBalancerOptionsModel.Enabled = core.BoolPtr(true)
				createLoadBalancerOptionsModel.TTL = core.Int64Ptr(int64(120))
				createLoadBalancerOptionsModel.FallbackPool = core.StringPtr("24ccf79a-4ae0-4769-b4c8-17f8f230072e")
				createLoadBalancerOptionsModel.DefaultPools = []string{"testString"}
				createLoadBalancerOptionsModel.AzPools = []dnssvcsv1.LoadBalancerAzPoolsItem{*loadBalancerAzPoolsItemModel}
				createLoadBalancerOptionsModel.XCorrelationID = core.StringPtr("testString")
				createLoadBalancerOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = dnsSvcsService.CreateLoadBalancer(createLoadBalancerOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = dnsSvcsService.CreateLoadBalancerWithContext(ctx, createLoadBalancerOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				dnsSvcsService.DisableRetries()
				result, response, operationErr = dnsSvcsService.CreateLoadBalancer(createLoadBalancerOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = dnsSvcsService.CreateLoadBalancerWithContext(ctx, createLoadBalancerOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke CreateLoadBalancer with error: Operation validation and request error`, func() {
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsSvcsService).ToNot(BeNil())

				// Construct an instance of the LoadBalancerAzPoolsItem model
				loadBalancerAzPoolsItemModel := new(dnssvcsv1.LoadBalancerAzPoolsItem)
				loadBalancerAzPoolsItemModel.AvailabilityZone = core.StringPtr("us-south-1")
				loadBalancerAzPoolsItemModel.Pools = []string{"0fc0bb7c-2fab-476e-8b9b-40fa14bf8e3d"}

				// Construct an instance of the CreateLoadBalancerOptions model
				createLoadBalancerOptionsModel := new(dnssvcsv1.CreateLoadBalancerOptions)
				createLoadBalancerOptionsModel.InstanceID = core.StringPtr("testString")
				createLoadBalancerOptionsModel.DnszoneID = core.StringPtr("testString")
				createLoadBalancerOptionsModel.Name = core.StringPtr("glb.example.com")
				createLoadBalancerOptionsModel.Description = core.StringPtr("Load balancer for glb.example.com.")
				createLoadBalancerOptionsModel.Enabled = core.BoolPtr(true)
				createLoadBalancerOptionsModel.TTL = core.Int64Ptr(int64(120))
				createLoadBalancerOptionsModel.FallbackPool = core.StringPtr("24ccf79a-4ae0-4769-b4c8-17f8f230072e")
				createLoadBalancerOptionsModel.DefaultPools = []string{"testString"}
				createLoadBalancerOptionsModel.AzPools = []dnssvcsv1.LoadBalancerAzPoolsItem{*loadBalancerAzPoolsItemModel}
				createLoadBalancerOptionsModel.XCorrelationID = core.StringPtr("testString")
				createLoadBalancerOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := dnsSvcsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := dnsSvcsService.CreateLoadBalancer(createLoadBalancerOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateLoadBalancerOptions model with no property values
				createLoadBalancerOptionsModelNew := new(dnssvcsv1.CreateLoadBalancerOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = dnsSvcsService.CreateLoadBalancer(createLoadBalancerOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`DeleteLoadBalancer(deleteLoadBalancerOptions *DeleteLoadBalancerOptions)`, func() {
		deleteLoadBalancerPath := "/instances/testString/dnszones/testString/load_balancers/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteLoadBalancerPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteLoadBalancer successfully`, func() {
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsSvcsService).ToNot(BeNil())
				dnsSvcsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				response, operationErr := dnsSvcsService.DeleteLoadBalancer(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteLoadBalancerOptions model
				deleteLoadBalancerOptionsModel := new(dnssvcsv1.DeleteLoadBalancerOptions)
				deleteLoadBalancerOptionsModel.InstanceID = core.StringPtr("testString")
				deleteLoadBalancerOptionsModel.DnszoneID = core.StringPtr("testString")
				deleteLoadBalancerOptionsModel.LbID = core.StringPtr("testString")
				deleteLoadBalancerOptionsModel.XCorrelationID = core.StringPtr("testString")
				deleteLoadBalancerOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = dnsSvcsService.DeleteLoadBalancer(deleteLoadBalancerOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Disable retries and test again
				dnsSvcsService.DisableRetries()
				response, operationErr = dnsSvcsService.DeleteLoadBalancer(deleteLoadBalancerOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteLoadBalancer with error: Operation validation and request error`, func() {
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsSvcsService).ToNot(BeNil())

				// Construct an instance of the DeleteLoadBalancerOptions model
				deleteLoadBalancerOptionsModel := new(dnssvcsv1.DeleteLoadBalancerOptions)
				deleteLoadBalancerOptionsModel.InstanceID = core.StringPtr("testString")
				deleteLoadBalancerOptionsModel.DnszoneID = core.StringPtr("testString")
				deleteLoadBalancerOptionsModel.LbID = core.StringPtr("testString")
				deleteLoadBalancerOptionsModel.XCorrelationID = core.StringPtr("testString")
				deleteLoadBalancerOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := dnsSvcsService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := dnsSvcsService.DeleteLoadBalancer(deleteLoadBalancerOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteLoadBalancerOptions model with no property values
				deleteLoadBalancerOptionsModelNew := new(dnssvcsv1.DeleteLoadBalancerOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = dnsSvcsService.DeleteLoadBalancer(deleteLoadBalancerOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetLoadBalancer(getLoadBalancerOptions *GetLoadBalancerOptions) - Operation response error`, func() {
		getLoadBalancerPath := "/instances/testString/dnszones/testString/load_balancers/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getLoadBalancerPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetLoadBalancer with error: Operation response processing error`, func() {
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsSvcsService).ToNot(BeNil())

				// Construct an instance of the GetLoadBalancerOptions model
				getLoadBalancerOptionsModel := new(dnssvcsv1.GetLoadBalancerOptions)
				getLoadBalancerOptionsModel.InstanceID = core.StringPtr("testString")
				getLoadBalancerOptionsModel.DnszoneID = core.StringPtr("testString")
				getLoadBalancerOptionsModel.LbID = core.StringPtr("testString")
				getLoadBalancerOptionsModel.XCorrelationID = core.StringPtr("testString")
				getLoadBalancerOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := dnsSvcsService.GetLoadBalancer(getLoadBalancerOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				dnsSvcsService.EnableRetries(0, 0)
				result, response, operationErr = dnsSvcsService.GetLoadBalancer(getLoadBalancerOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetLoadBalancer(getLoadBalancerOptions *GetLoadBalancerOptions)`, func() {
		getLoadBalancerPath := "/instances/testString/dnszones/testString/load_balancers/testString"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getLoadBalancerPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "5365b73c-ce6f-4d6f-ad9f-d9c131b26370", "name": "glb.example.com", "description": "Load balancer for glb.example.com.", "enabled": true, "ttl": 120, "health": "DEGRADED", "fallback_pool": "24ccf79a-4ae0-4769-b4c8-17f8f230072e", "default_pools": ["DefaultPools"], "az_pools": [{"availability_zone": "us-south-1", "pools": ["0fc0bb7c-2fab-476e-8b9b-40fa14bf8e3d"]}], "created_on": "2019-01-01T05:20:00.12345Z", "modified_on": "2019-01-01T05:20:00.12345Z"}`)
				}))
			})
			It(`Invoke GetLoadBalancer successfully`, func() {
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsSvcsService).ToNot(BeNil())
				dnsSvcsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := dnsSvcsService.GetLoadBalancer(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetLoadBalancerOptions model
				getLoadBalancerOptionsModel := new(dnssvcsv1.GetLoadBalancerOptions)
				getLoadBalancerOptionsModel.InstanceID = core.StringPtr("testString")
				getLoadBalancerOptionsModel.DnszoneID = core.StringPtr("testString")
				getLoadBalancerOptionsModel.LbID = core.StringPtr("testString")
				getLoadBalancerOptionsModel.XCorrelationID = core.StringPtr("testString")
				getLoadBalancerOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = dnsSvcsService.GetLoadBalancer(getLoadBalancerOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = dnsSvcsService.GetLoadBalancerWithContext(ctx, getLoadBalancerOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				dnsSvcsService.DisableRetries()
				result, response, operationErr = dnsSvcsService.GetLoadBalancer(getLoadBalancerOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = dnsSvcsService.GetLoadBalancerWithContext(ctx, getLoadBalancerOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke GetLoadBalancer with error: Operation validation and request error`, func() {
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsSvcsService).ToNot(BeNil())

				// Construct an instance of the GetLoadBalancerOptions model
				getLoadBalancerOptionsModel := new(dnssvcsv1.GetLoadBalancerOptions)
				getLoadBalancerOptionsModel.InstanceID = core.StringPtr("testString")
				getLoadBalancerOptionsModel.DnszoneID = core.StringPtr("testString")
				getLoadBalancerOptionsModel.LbID = core.StringPtr("testString")
				getLoadBalancerOptionsModel.XCorrelationID = core.StringPtr("testString")
				getLoadBalancerOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := dnsSvcsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := dnsSvcsService.GetLoadBalancer(getLoadBalancerOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetLoadBalancerOptions model with no property values
				getLoadBalancerOptionsModelNew := new(dnssvcsv1.GetLoadBalancerOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = dnsSvcsService.GetLoadBalancer(getLoadBalancerOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateLoadBalancer(updateLoadBalancerOptions *UpdateLoadBalancerOptions) - Operation response error`, func() {
		updateLoadBalancerPath := "/instances/testString/dnszones/testString/load_balancers/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateLoadBalancerPath))
					Expect(req.Method).To(Equal("PUT"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateLoadBalancer with error: Operation response processing error`, func() {
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsSvcsService).ToNot(BeNil())

				// Construct an instance of the LoadBalancerAzPoolsItem model
				loadBalancerAzPoolsItemModel := new(dnssvcsv1.LoadBalancerAzPoolsItem)
				loadBalancerAzPoolsItemModel.AvailabilityZone = core.StringPtr("us-south-1")
				loadBalancerAzPoolsItemModel.Pools = []string{"0fc0bb7c-2fab-476e-8b9b-40fa14bf8e3d"}

				// Construct an instance of the UpdateLoadBalancerOptions model
				updateLoadBalancerOptionsModel := new(dnssvcsv1.UpdateLoadBalancerOptions)
				updateLoadBalancerOptionsModel.InstanceID = core.StringPtr("testString")
				updateLoadBalancerOptionsModel.DnszoneID = core.StringPtr("testString")
				updateLoadBalancerOptionsModel.LbID = core.StringPtr("testString")
				updateLoadBalancerOptionsModel.Name = core.StringPtr("glb.example.com")
				updateLoadBalancerOptionsModel.Description = core.StringPtr("Load balancer for glb.example.com.")
				updateLoadBalancerOptionsModel.Enabled = core.BoolPtr(true)
				updateLoadBalancerOptionsModel.TTL = core.Int64Ptr(int64(120))
				updateLoadBalancerOptionsModel.FallbackPool = core.StringPtr("24ccf79a-4ae0-4769-b4c8-17f8f230072e")
				updateLoadBalancerOptionsModel.DefaultPools = []string{"testString"}
				updateLoadBalancerOptionsModel.AzPools = []dnssvcsv1.LoadBalancerAzPoolsItem{*loadBalancerAzPoolsItemModel}
				updateLoadBalancerOptionsModel.XCorrelationID = core.StringPtr("testString")
				updateLoadBalancerOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := dnsSvcsService.UpdateLoadBalancer(updateLoadBalancerOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				dnsSvcsService.EnableRetries(0, 0)
				result, response, operationErr = dnsSvcsService.UpdateLoadBalancer(updateLoadBalancerOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UpdateLoadBalancer(updateLoadBalancerOptions *UpdateLoadBalancerOptions)`, func() {
		updateLoadBalancerPath := "/instances/testString/dnszones/testString/load_balancers/testString"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateLoadBalancerPath))
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

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "5365b73c-ce6f-4d6f-ad9f-d9c131b26370", "name": "glb.example.com", "description": "Load balancer for glb.example.com.", "enabled": true, "ttl": 120, "health": "DEGRADED", "fallback_pool": "24ccf79a-4ae0-4769-b4c8-17f8f230072e", "default_pools": ["DefaultPools"], "az_pools": [{"availability_zone": "us-south-1", "pools": ["0fc0bb7c-2fab-476e-8b9b-40fa14bf8e3d"]}], "created_on": "2019-01-01T05:20:00.12345Z", "modified_on": "2019-01-01T05:20:00.12345Z"}`)
				}))
			})
			It(`Invoke UpdateLoadBalancer successfully`, func() {
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsSvcsService).ToNot(BeNil())
				dnsSvcsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := dnsSvcsService.UpdateLoadBalancer(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the LoadBalancerAzPoolsItem model
				loadBalancerAzPoolsItemModel := new(dnssvcsv1.LoadBalancerAzPoolsItem)
				loadBalancerAzPoolsItemModel.AvailabilityZone = core.StringPtr("us-south-1")
				loadBalancerAzPoolsItemModel.Pools = []string{"0fc0bb7c-2fab-476e-8b9b-40fa14bf8e3d"}

				// Construct an instance of the UpdateLoadBalancerOptions model
				updateLoadBalancerOptionsModel := new(dnssvcsv1.UpdateLoadBalancerOptions)
				updateLoadBalancerOptionsModel.InstanceID = core.StringPtr("testString")
				updateLoadBalancerOptionsModel.DnszoneID = core.StringPtr("testString")
				updateLoadBalancerOptionsModel.LbID = core.StringPtr("testString")
				updateLoadBalancerOptionsModel.Name = core.StringPtr("glb.example.com")
				updateLoadBalancerOptionsModel.Description = core.StringPtr("Load balancer for glb.example.com.")
				updateLoadBalancerOptionsModel.Enabled = core.BoolPtr(true)
				updateLoadBalancerOptionsModel.TTL = core.Int64Ptr(int64(120))
				updateLoadBalancerOptionsModel.FallbackPool = core.StringPtr("24ccf79a-4ae0-4769-b4c8-17f8f230072e")
				updateLoadBalancerOptionsModel.DefaultPools = []string{"testString"}
				updateLoadBalancerOptionsModel.AzPools = []dnssvcsv1.LoadBalancerAzPoolsItem{*loadBalancerAzPoolsItemModel}
				updateLoadBalancerOptionsModel.XCorrelationID = core.StringPtr("testString")
				updateLoadBalancerOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = dnsSvcsService.UpdateLoadBalancer(updateLoadBalancerOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = dnsSvcsService.UpdateLoadBalancerWithContext(ctx, updateLoadBalancerOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				dnsSvcsService.DisableRetries()
				result, response, operationErr = dnsSvcsService.UpdateLoadBalancer(updateLoadBalancerOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = dnsSvcsService.UpdateLoadBalancerWithContext(ctx, updateLoadBalancerOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke UpdateLoadBalancer with error: Operation validation and request error`, func() {
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsSvcsService).ToNot(BeNil())

				// Construct an instance of the LoadBalancerAzPoolsItem model
				loadBalancerAzPoolsItemModel := new(dnssvcsv1.LoadBalancerAzPoolsItem)
				loadBalancerAzPoolsItemModel.AvailabilityZone = core.StringPtr("us-south-1")
				loadBalancerAzPoolsItemModel.Pools = []string{"0fc0bb7c-2fab-476e-8b9b-40fa14bf8e3d"}

				// Construct an instance of the UpdateLoadBalancerOptions model
				updateLoadBalancerOptionsModel := new(dnssvcsv1.UpdateLoadBalancerOptions)
				updateLoadBalancerOptionsModel.InstanceID = core.StringPtr("testString")
				updateLoadBalancerOptionsModel.DnszoneID = core.StringPtr("testString")
				updateLoadBalancerOptionsModel.LbID = core.StringPtr("testString")
				updateLoadBalancerOptionsModel.Name = core.StringPtr("glb.example.com")
				updateLoadBalancerOptionsModel.Description = core.StringPtr("Load balancer for glb.example.com.")
				updateLoadBalancerOptionsModel.Enabled = core.BoolPtr(true)
				updateLoadBalancerOptionsModel.TTL = core.Int64Ptr(int64(120))
				updateLoadBalancerOptionsModel.FallbackPool = core.StringPtr("24ccf79a-4ae0-4769-b4c8-17f8f230072e")
				updateLoadBalancerOptionsModel.DefaultPools = []string{"testString"}
				updateLoadBalancerOptionsModel.AzPools = []dnssvcsv1.LoadBalancerAzPoolsItem{*loadBalancerAzPoolsItemModel}
				updateLoadBalancerOptionsModel.XCorrelationID = core.StringPtr("testString")
				updateLoadBalancerOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := dnsSvcsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := dnsSvcsService.UpdateLoadBalancer(updateLoadBalancerOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateLoadBalancerOptions model with no property values
				updateLoadBalancerOptionsModelNew := new(dnssvcsv1.UpdateLoadBalancerOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = dnsSvcsService.UpdateLoadBalancer(updateLoadBalancerOptionsModelNew)
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
		It(`Instantiate service client`, func() {
			dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(dnsSvcsService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(dnsSvcsService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
				URL: "https://dnssvcsv1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(dnsSvcsService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"DNS_SVCS_URL": "https://dnssvcsv1/api",
				"DNS_SVCS_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1UsingExternalConfig(&dnssvcsv1.DnsSvcsV1Options{
				})
				Expect(dnsSvcsService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := dnsSvcsService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != dnsSvcsService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(dnsSvcsService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(dnsSvcsService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1UsingExternalConfig(&dnssvcsv1.DnsSvcsV1Options{
					URL: "https://testService/api",
				})
				Expect(dnsSvcsService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(dnsSvcsService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := dnsSvcsService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != dnsSvcsService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(dnsSvcsService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(dnsSvcsService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1UsingExternalConfig(&dnssvcsv1.DnsSvcsV1Options{
				})
				err := dnsSvcsService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(dnsSvcsService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(dnsSvcsService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := dnsSvcsService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != dnsSvcsService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(dnsSvcsService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(dnsSvcsService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"DNS_SVCS_URL": "https://dnssvcsv1/api",
				"DNS_SVCS_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1UsingExternalConfig(&dnssvcsv1.DnsSvcsV1Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(dnsSvcsService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"DNS_SVCS_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1UsingExternalConfig(&dnssvcsv1.DnsSvcsV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(dnsSvcsService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = dnssvcsv1.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`ListPools(listPoolsOptions *ListPoolsOptions) - Operation response error`, func() {
		listPoolsPath := "/instances/testString/pools"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listPoolsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListPools with error: Operation response processing error`, func() {
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsSvcsService).ToNot(BeNil())

				// Construct an instance of the ListPoolsOptions model
				listPoolsOptionsModel := new(dnssvcsv1.ListPoolsOptions)
				listPoolsOptionsModel.InstanceID = core.StringPtr("testString")
				listPoolsOptionsModel.XCorrelationID = core.StringPtr("testString")
				listPoolsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := dnsSvcsService.ListPools(listPoolsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				dnsSvcsService.EnableRetries(0, 0)
				result, response, operationErr = dnsSvcsService.ListPools(listPoolsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ListPools(listPoolsOptions *ListPoolsOptions)`, func() {
		listPoolsPath := "/instances/testString/pools"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listPoolsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"pools": [{"id": "5365b73c-ce6f-4d6f-ad9f-d9c131b26370", "name": "dal10-az-pool", "description": "Load balancer pool for dal10 availability zone.", "enabled": true, "healthy_origins_threshold": 1, "origins": [{"name": "app-server-1", "description": "description of the origin server", "address": "10.10.16.8", "enabled": true, "health": true, "health_failure_reason": "HealthFailureReason"}], "monitor": "7dd6841c-264e-11ea-88df-062967242a6a", "notification_channel": "https://mywebsite.com/dns/webhook", "health": "HEALTHY", "healthcheck_region": "us-south", "healthcheck_subnets": ["crn:v1:staging:public:is:us-south-1:a/01652b251c3ae2787110a995d8db0135::subnet:0716-b49ef064-0f89-4fb1-8212-135b12568f04"], "healthcheck_vsis": [{"subnet": "crn:v1:staging:public:is:us-south-1:a/01652b251c3ae2787110a995d8db0135::subnet:0716-b49ef064-0f89-4fb1-8212-135b12568f04", "ipv4_address": "10.10.16.8", "ipv4_cidr_block": "10.10.16.0/24", "vpc": "crn:v1:staging:public:is:us-south:a/01652b251c3ae2787110a995d8db0135::vpc:r134-8c426a0a-ec74-4c97-9c02-f6194c224d8a"}], "created_on": "2019-01-01T05:20:00.12345Z", "modified_on": "2019-01-01T05:20:00.12345Z"}], "offset": 1, "limit": 20, "count": 1, "total_count": 200, "first": {"href": "https://api.dns-svcs.cloud.ibm.com/v1/instances/434f6c3e-6014-4124-a61d-2e910bca19b1/dnszones/example.com:d04d3a7a-7f6d-47d4-b811-08c5478fa1a4/resource_records?limit=20"}, "next": {"href": "https://api.dns-svcs.cloud.ibm.com/v1/instances/434f6c3e-6014-4124-a61d-2e910bca19b1/dnszones/example.com:d04d3a7a-7f6d-47d4-b811-08c5478fa1a4/resource_records?offset=20&limit=20"}}`)
				}))
			})
			It(`Invoke ListPools successfully`, func() {
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsSvcsService).ToNot(BeNil())
				dnsSvcsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := dnsSvcsService.ListPools(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListPoolsOptions model
				listPoolsOptionsModel := new(dnssvcsv1.ListPoolsOptions)
				listPoolsOptionsModel.InstanceID = core.StringPtr("testString")
				listPoolsOptionsModel.XCorrelationID = core.StringPtr("testString")
				listPoolsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = dnsSvcsService.ListPools(listPoolsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = dnsSvcsService.ListPoolsWithContext(ctx, listPoolsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				dnsSvcsService.DisableRetries()
				result, response, operationErr = dnsSvcsService.ListPools(listPoolsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = dnsSvcsService.ListPoolsWithContext(ctx, listPoolsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke ListPools with error: Operation validation and request error`, func() {
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsSvcsService).ToNot(BeNil())

				// Construct an instance of the ListPoolsOptions model
				listPoolsOptionsModel := new(dnssvcsv1.ListPoolsOptions)
				listPoolsOptionsModel.InstanceID = core.StringPtr("testString")
				listPoolsOptionsModel.XCorrelationID = core.StringPtr("testString")
				listPoolsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := dnsSvcsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := dnsSvcsService.ListPools(listPoolsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListPoolsOptions model with no property values
				listPoolsOptionsModelNew := new(dnssvcsv1.ListPoolsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = dnsSvcsService.ListPools(listPoolsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreatePool(createPoolOptions *CreatePoolOptions) - Operation response error`, func() {
		createPoolPath := "/instances/testString/pools"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createPoolPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreatePool with error: Operation response processing error`, func() {
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsSvcsService).ToNot(BeNil())

				// Construct an instance of the OriginInput model
				originInputModel := new(dnssvcsv1.OriginInput)
				originInputModel.Name = core.StringPtr("app-server-1")
				originInputModel.Description = core.StringPtr("description of the origin server")
				originInputModel.Address = core.StringPtr("10.10.16.8")
				originInputModel.Enabled = core.BoolPtr(true)

				// Construct an instance of the CreatePoolOptions model
				createPoolOptionsModel := new(dnssvcsv1.CreatePoolOptions)
				createPoolOptionsModel.InstanceID = core.StringPtr("testString")
				createPoolOptionsModel.Name = core.StringPtr("dal10-az-pool")
				createPoolOptionsModel.Description = core.StringPtr("Load balancer pool for dal10 availability zone.")
				createPoolOptionsModel.Enabled = core.BoolPtr(true)
				createPoolOptionsModel.HealthyOriginsThreshold = core.Int64Ptr(int64(1))
				createPoolOptionsModel.Origins = []dnssvcsv1.OriginInput{*originInputModel}
				createPoolOptionsModel.Monitor = core.StringPtr("7dd6841c-264e-11ea-88df-062967242a6a")
				createPoolOptionsModel.NotificationChannel = core.StringPtr("https://mywebsite.com/dns/webhook")
				createPoolOptionsModel.HealthcheckRegion = core.StringPtr("us-south")
				createPoolOptionsModel.HealthcheckSubnets = []string{"crn:v1:staging:public:is:us-south-1:a/01652b251c3ae2787110a995d8db0135::subnet:0716-b49ef064-0f89-4fb1-8212-135b12568f04"}
				createPoolOptionsModel.XCorrelationID = core.StringPtr("testString")
				createPoolOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := dnsSvcsService.CreatePool(createPoolOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				dnsSvcsService.EnableRetries(0, 0)
				result, response, operationErr = dnsSvcsService.CreatePool(createPoolOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`CreatePool(createPoolOptions *CreatePoolOptions)`, func() {
		createPoolPath := "/instances/testString/pools"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createPoolPath))
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

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "5365b73c-ce6f-4d6f-ad9f-d9c131b26370", "name": "dal10-az-pool", "description": "Load balancer pool for dal10 availability zone.", "enabled": true, "healthy_origins_threshold": 1, "origins": [{"name": "app-server-1", "description": "description of the origin server", "address": "10.10.16.8", "enabled": true, "health": true, "health_failure_reason": "HealthFailureReason"}], "monitor": "7dd6841c-264e-11ea-88df-062967242a6a", "notification_channel": "https://mywebsite.com/dns/webhook", "health": "HEALTHY", "healthcheck_region": "us-south", "healthcheck_subnets": ["crn:v1:staging:public:is:us-south-1:a/01652b251c3ae2787110a995d8db0135::subnet:0716-b49ef064-0f89-4fb1-8212-135b12568f04"], "healthcheck_vsis": [{"subnet": "crn:v1:staging:public:is:us-south-1:a/01652b251c3ae2787110a995d8db0135::subnet:0716-b49ef064-0f89-4fb1-8212-135b12568f04", "ipv4_address": "10.10.16.8", "ipv4_cidr_block": "10.10.16.0/24", "vpc": "crn:v1:staging:public:is:us-south:a/01652b251c3ae2787110a995d8db0135::vpc:r134-8c426a0a-ec74-4c97-9c02-f6194c224d8a"}], "created_on": "2019-01-01T05:20:00.12345Z", "modified_on": "2019-01-01T05:20:00.12345Z"}`)
				}))
			})
			It(`Invoke CreatePool successfully`, func() {
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsSvcsService).ToNot(BeNil())
				dnsSvcsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := dnsSvcsService.CreatePool(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the OriginInput model
				originInputModel := new(dnssvcsv1.OriginInput)
				originInputModel.Name = core.StringPtr("app-server-1")
				originInputModel.Description = core.StringPtr("description of the origin server")
				originInputModel.Address = core.StringPtr("10.10.16.8")
				originInputModel.Enabled = core.BoolPtr(true)

				// Construct an instance of the CreatePoolOptions model
				createPoolOptionsModel := new(dnssvcsv1.CreatePoolOptions)
				createPoolOptionsModel.InstanceID = core.StringPtr("testString")
				createPoolOptionsModel.Name = core.StringPtr("dal10-az-pool")
				createPoolOptionsModel.Description = core.StringPtr("Load balancer pool for dal10 availability zone.")
				createPoolOptionsModel.Enabled = core.BoolPtr(true)
				createPoolOptionsModel.HealthyOriginsThreshold = core.Int64Ptr(int64(1))
				createPoolOptionsModel.Origins = []dnssvcsv1.OriginInput{*originInputModel}
				createPoolOptionsModel.Monitor = core.StringPtr("7dd6841c-264e-11ea-88df-062967242a6a")
				createPoolOptionsModel.NotificationChannel = core.StringPtr("https://mywebsite.com/dns/webhook")
				createPoolOptionsModel.HealthcheckRegion = core.StringPtr("us-south")
				createPoolOptionsModel.HealthcheckSubnets = []string{"crn:v1:staging:public:is:us-south-1:a/01652b251c3ae2787110a995d8db0135::subnet:0716-b49ef064-0f89-4fb1-8212-135b12568f04"}
				createPoolOptionsModel.XCorrelationID = core.StringPtr("testString")
				createPoolOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = dnsSvcsService.CreatePool(createPoolOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = dnsSvcsService.CreatePoolWithContext(ctx, createPoolOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				dnsSvcsService.DisableRetries()
				result, response, operationErr = dnsSvcsService.CreatePool(createPoolOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = dnsSvcsService.CreatePoolWithContext(ctx, createPoolOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke CreatePool with error: Operation validation and request error`, func() {
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsSvcsService).ToNot(BeNil())

				// Construct an instance of the OriginInput model
				originInputModel := new(dnssvcsv1.OriginInput)
				originInputModel.Name = core.StringPtr("app-server-1")
				originInputModel.Description = core.StringPtr("description of the origin server")
				originInputModel.Address = core.StringPtr("10.10.16.8")
				originInputModel.Enabled = core.BoolPtr(true)

				// Construct an instance of the CreatePoolOptions model
				createPoolOptionsModel := new(dnssvcsv1.CreatePoolOptions)
				createPoolOptionsModel.InstanceID = core.StringPtr("testString")
				createPoolOptionsModel.Name = core.StringPtr("dal10-az-pool")
				createPoolOptionsModel.Description = core.StringPtr("Load balancer pool for dal10 availability zone.")
				createPoolOptionsModel.Enabled = core.BoolPtr(true)
				createPoolOptionsModel.HealthyOriginsThreshold = core.Int64Ptr(int64(1))
				createPoolOptionsModel.Origins = []dnssvcsv1.OriginInput{*originInputModel}
				createPoolOptionsModel.Monitor = core.StringPtr("7dd6841c-264e-11ea-88df-062967242a6a")
				createPoolOptionsModel.NotificationChannel = core.StringPtr("https://mywebsite.com/dns/webhook")
				createPoolOptionsModel.HealthcheckRegion = core.StringPtr("us-south")
				createPoolOptionsModel.HealthcheckSubnets = []string{"crn:v1:staging:public:is:us-south-1:a/01652b251c3ae2787110a995d8db0135::subnet:0716-b49ef064-0f89-4fb1-8212-135b12568f04"}
				createPoolOptionsModel.XCorrelationID = core.StringPtr("testString")
				createPoolOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := dnsSvcsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := dnsSvcsService.CreatePool(createPoolOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreatePoolOptions model with no property values
				createPoolOptionsModelNew := new(dnssvcsv1.CreatePoolOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = dnsSvcsService.CreatePool(createPoolOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`DeletePool(deletePoolOptions *DeletePoolOptions)`, func() {
		deletePoolPath := "/instances/testString/pools/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deletePoolPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeletePool successfully`, func() {
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsSvcsService).ToNot(BeNil())
				dnsSvcsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				response, operationErr := dnsSvcsService.DeletePool(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeletePoolOptions model
				deletePoolOptionsModel := new(dnssvcsv1.DeletePoolOptions)
				deletePoolOptionsModel.InstanceID = core.StringPtr("testString")
				deletePoolOptionsModel.PoolID = core.StringPtr("testString")
				deletePoolOptionsModel.XCorrelationID = core.StringPtr("testString")
				deletePoolOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = dnsSvcsService.DeletePool(deletePoolOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Disable retries and test again
				dnsSvcsService.DisableRetries()
				response, operationErr = dnsSvcsService.DeletePool(deletePoolOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeletePool with error: Operation validation and request error`, func() {
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsSvcsService).ToNot(BeNil())

				// Construct an instance of the DeletePoolOptions model
				deletePoolOptionsModel := new(dnssvcsv1.DeletePoolOptions)
				deletePoolOptionsModel.InstanceID = core.StringPtr("testString")
				deletePoolOptionsModel.PoolID = core.StringPtr("testString")
				deletePoolOptionsModel.XCorrelationID = core.StringPtr("testString")
				deletePoolOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := dnsSvcsService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := dnsSvcsService.DeletePool(deletePoolOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeletePoolOptions model with no property values
				deletePoolOptionsModelNew := new(dnssvcsv1.DeletePoolOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = dnsSvcsService.DeletePool(deletePoolOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetPool(getPoolOptions *GetPoolOptions) - Operation response error`, func() {
		getPoolPath := "/instances/testString/pools/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getPoolPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetPool with error: Operation response processing error`, func() {
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsSvcsService).ToNot(BeNil())

				// Construct an instance of the GetPoolOptions model
				getPoolOptionsModel := new(dnssvcsv1.GetPoolOptions)
				getPoolOptionsModel.InstanceID = core.StringPtr("testString")
				getPoolOptionsModel.PoolID = core.StringPtr("testString")
				getPoolOptionsModel.XCorrelationID = core.StringPtr("testString")
				getPoolOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := dnsSvcsService.GetPool(getPoolOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				dnsSvcsService.EnableRetries(0, 0)
				result, response, operationErr = dnsSvcsService.GetPool(getPoolOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetPool(getPoolOptions *GetPoolOptions)`, func() {
		getPoolPath := "/instances/testString/pools/testString"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getPoolPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "5365b73c-ce6f-4d6f-ad9f-d9c131b26370", "name": "dal10-az-pool", "description": "Load balancer pool for dal10 availability zone.", "enabled": true, "healthy_origins_threshold": 1, "origins": [{"name": "app-server-1", "description": "description of the origin server", "address": "10.10.16.8", "enabled": true, "health": true, "health_failure_reason": "HealthFailureReason"}], "monitor": "7dd6841c-264e-11ea-88df-062967242a6a", "notification_channel": "https://mywebsite.com/dns/webhook", "health": "HEALTHY", "healthcheck_region": "us-south", "healthcheck_subnets": ["crn:v1:staging:public:is:us-south-1:a/01652b251c3ae2787110a995d8db0135::subnet:0716-b49ef064-0f89-4fb1-8212-135b12568f04"], "healthcheck_vsis": [{"subnet": "crn:v1:staging:public:is:us-south-1:a/01652b251c3ae2787110a995d8db0135::subnet:0716-b49ef064-0f89-4fb1-8212-135b12568f04", "ipv4_address": "10.10.16.8", "ipv4_cidr_block": "10.10.16.0/24", "vpc": "crn:v1:staging:public:is:us-south:a/01652b251c3ae2787110a995d8db0135::vpc:r134-8c426a0a-ec74-4c97-9c02-f6194c224d8a"}], "created_on": "2019-01-01T05:20:00.12345Z", "modified_on": "2019-01-01T05:20:00.12345Z"}`)
				}))
			})
			It(`Invoke GetPool successfully`, func() {
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsSvcsService).ToNot(BeNil())
				dnsSvcsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := dnsSvcsService.GetPool(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetPoolOptions model
				getPoolOptionsModel := new(dnssvcsv1.GetPoolOptions)
				getPoolOptionsModel.InstanceID = core.StringPtr("testString")
				getPoolOptionsModel.PoolID = core.StringPtr("testString")
				getPoolOptionsModel.XCorrelationID = core.StringPtr("testString")
				getPoolOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = dnsSvcsService.GetPool(getPoolOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = dnsSvcsService.GetPoolWithContext(ctx, getPoolOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				dnsSvcsService.DisableRetries()
				result, response, operationErr = dnsSvcsService.GetPool(getPoolOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = dnsSvcsService.GetPoolWithContext(ctx, getPoolOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke GetPool with error: Operation validation and request error`, func() {
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsSvcsService).ToNot(BeNil())

				// Construct an instance of the GetPoolOptions model
				getPoolOptionsModel := new(dnssvcsv1.GetPoolOptions)
				getPoolOptionsModel.InstanceID = core.StringPtr("testString")
				getPoolOptionsModel.PoolID = core.StringPtr("testString")
				getPoolOptionsModel.XCorrelationID = core.StringPtr("testString")
				getPoolOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := dnsSvcsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := dnsSvcsService.GetPool(getPoolOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetPoolOptions model with no property values
				getPoolOptionsModelNew := new(dnssvcsv1.GetPoolOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = dnsSvcsService.GetPool(getPoolOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdatePool(updatePoolOptions *UpdatePoolOptions) - Operation response error`, func() {
		updatePoolPath := "/instances/testString/pools/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updatePoolPath))
					Expect(req.Method).To(Equal("PUT"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdatePool with error: Operation response processing error`, func() {
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsSvcsService).ToNot(BeNil())

				// Construct an instance of the OriginInput model
				originInputModel := new(dnssvcsv1.OriginInput)
				originInputModel.Name = core.StringPtr("app-server-1")
				originInputModel.Description = core.StringPtr("description of the origin server")
				originInputModel.Address = core.StringPtr("10.10.16.8")
				originInputModel.Enabled = core.BoolPtr(true)

				// Construct an instance of the UpdatePoolOptions model
				updatePoolOptionsModel := new(dnssvcsv1.UpdatePoolOptions)
				updatePoolOptionsModel.InstanceID = core.StringPtr("testString")
				updatePoolOptionsModel.PoolID = core.StringPtr("testString")
				updatePoolOptionsModel.Name = core.StringPtr("dal10-az-pool")
				updatePoolOptionsModel.Description = core.StringPtr("Load balancer pool for dal10 availability zone.")
				updatePoolOptionsModel.Enabled = core.BoolPtr(true)
				updatePoolOptionsModel.HealthyOriginsThreshold = core.Int64Ptr(int64(1))
				updatePoolOptionsModel.Origins = []dnssvcsv1.OriginInput{*originInputModel}
				updatePoolOptionsModel.Monitor = core.StringPtr("7dd6841c-264e-11ea-88df-062967242a6a")
				updatePoolOptionsModel.NotificationChannel = core.StringPtr("https://mywebsite.com/dns/webhook")
				updatePoolOptionsModel.HealthcheckRegion = core.StringPtr("us-south")
				updatePoolOptionsModel.HealthcheckSubnets = []string{"crn:v1:staging:public:is:us-south-1:a/01652b251c3ae2787110a995d8db0135::subnet:0716-b49ef064-0f89-4fb1-8212-135b12568f04"}
				updatePoolOptionsModel.XCorrelationID = core.StringPtr("testString")
				updatePoolOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := dnsSvcsService.UpdatePool(updatePoolOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				dnsSvcsService.EnableRetries(0, 0)
				result, response, operationErr = dnsSvcsService.UpdatePool(updatePoolOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UpdatePool(updatePoolOptions *UpdatePoolOptions)`, func() {
		updatePoolPath := "/instances/testString/pools/testString"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updatePoolPath))
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

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "5365b73c-ce6f-4d6f-ad9f-d9c131b26370", "name": "dal10-az-pool", "description": "Load balancer pool for dal10 availability zone.", "enabled": true, "healthy_origins_threshold": 1, "origins": [{"name": "app-server-1", "description": "description of the origin server", "address": "10.10.16.8", "enabled": true, "health": true, "health_failure_reason": "HealthFailureReason"}], "monitor": "7dd6841c-264e-11ea-88df-062967242a6a", "notification_channel": "https://mywebsite.com/dns/webhook", "health": "HEALTHY", "healthcheck_region": "us-south", "healthcheck_subnets": ["crn:v1:staging:public:is:us-south-1:a/01652b251c3ae2787110a995d8db0135::subnet:0716-b49ef064-0f89-4fb1-8212-135b12568f04"], "healthcheck_vsis": [{"subnet": "crn:v1:staging:public:is:us-south-1:a/01652b251c3ae2787110a995d8db0135::subnet:0716-b49ef064-0f89-4fb1-8212-135b12568f04", "ipv4_address": "10.10.16.8", "ipv4_cidr_block": "10.10.16.0/24", "vpc": "crn:v1:staging:public:is:us-south:a/01652b251c3ae2787110a995d8db0135::vpc:r134-8c426a0a-ec74-4c97-9c02-f6194c224d8a"}], "created_on": "2019-01-01T05:20:00.12345Z", "modified_on": "2019-01-01T05:20:00.12345Z"}`)
				}))
			})
			It(`Invoke UpdatePool successfully`, func() {
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsSvcsService).ToNot(BeNil())
				dnsSvcsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := dnsSvcsService.UpdatePool(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the OriginInput model
				originInputModel := new(dnssvcsv1.OriginInput)
				originInputModel.Name = core.StringPtr("app-server-1")
				originInputModel.Description = core.StringPtr("description of the origin server")
				originInputModel.Address = core.StringPtr("10.10.16.8")
				originInputModel.Enabled = core.BoolPtr(true)

				// Construct an instance of the UpdatePoolOptions model
				updatePoolOptionsModel := new(dnssvcsv1.UpdatePoolOptions)
				updatePoolOptionsModel.InstanceID = core.StringPtr("testString")
				updatePoolOptionsModel.PoolID = core.StringPtr("testString")
				updatePoolOptionsModel.Name = core.StringPtr("dal10-az-pool")
				updatePoolOptionsModel.Description = core.StringPtr("Load balancer pool for dal10 availability zone.")
				updatePoolOptionsModel.Enabled = core.BoolPtr(true)
				updatePoolOptionsModel.HealthyOriginsThreshold = core.Int64Ptr(int64(1))
				updatePoolOptionsModel.Origins = []dnssvcsv1.OriginInput{*originInputModel}
				updatePoolOptionsModel.Monitor = core.StringPtr("7dd6841c-264e-11ea-88df-062967242a6a")
				updatePoolOptionsModel.NotificationChannel = core.StringPtr("https://mywebsite.com/dns/webhook")
				updatePoolOptionsModel.HealthcheckRegion = core.StringPtr("us-south")
				updatePoolOptionsModel.HealthcheckSubnets = []string{"crn:v1:staging:public:is:us-south-1:a/01652b251c3ae2787110a995d8db0135::subnet:0716-b49ef064-0f89-4fb1-8212-135b12568f04"}
				updatePoolOptionsModel.XCorrelationID = core.StringPtr("testString")
				updatePoolOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = dnsSvcsService.UpdatePool(updatePoolOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = dnsSvcsService.UpdatePoolWithContext(ctx, updatePoolOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				dnsSvcsService.DisableRetries()
				result, response, operationErr = dnsSvcsService.UpdatePool(updatePoolOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = dnsSvcsService.UpdatePoolWithContext(ctx, updatePoolOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke UpdatePool with error: Operation validation and request error`, func() {
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsSvcsService).ToNot(BeNil())

				// Construct an instance of the OriginInput model
				originInputModel := new(dnssvcsv1.OriginInput)
				originInputModel.Name = core.StringPtr("app-server-1")
				originInputModel.Description = core.StringPtr("description of the origin server")
				originInputModel.Address = core.StringPtr("10.10.16.8")
				originInputModel.Enabled = core.BoolPtr(true)

				// Construct an instance of the UpdatePoolOptions model
				updatePoolOptionsModel := new(dnssvcsv1.UpdatePoolOptions)
				updatePoolOptionsModel.InstanceID = core.StringPtr("testString")
				updatePoolOptionsModel.PoolID = core.StringPtr("testString")
				updatePoolOptionsModel.Name = core.StringPtr("dal10-az-pool")
				updatePoolOptionsModel.Description = core.StringPtr("Load balancer pool for dal10 availability zone.")
				updatePoolOptionsModel.Enabled = core.BoolPtr(true)
				updatePoolOptionsModel.HealthyOriginsThreshold = core.Int64Ptr(int64(1))
				updatePoolOptionsModel.Origins = []dnssvcsv1.OriginInput{*originInputModel}
				updatePoolOptionsModel.Monitor = core.StringPtr("7dd6841c-264e-11ea-88df-062967242a6a")
				updatePoolOptionsModel.NotificationChannel = core.StringPtr("https://mywebsite.com/dns/webhook")
				updatePoolOptionsModel.HealthcheckRegion = core.StringPtr("us-south")
				updatePoolOptionsModel.HealthcheckSubnets = []string{"crn:v1:staging:public:is:us-south-1:a/01652b251c3ae2787110a995d8db0135::subnet:0716-b49ef064-0f89-4fb1-8212-135b12568f04"}
				updatePoolOptionsModel.XCorrelationID = core.StringPtr("testString")
				updatePoolOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := dnsSvcsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := dnsSvcsService.UpdatePool(updatePoolOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdatePoolOptions model with no property values
				updatePoolOptionsModelNew := new(dnssvcsv1.UpdatePoolOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = dnsSvcsService.UpdatePool(updatePoolOptionsModelNew)
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
		It(`Instantiate service client`, func() {
			dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(dnsSvcsService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(dnsSvcsService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
				URL: "https://dnssvcsv1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(dnsSvcsService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"DNS_SVCS_URL": "https://dnssvcsv1/api",
				"DNS_SVCS_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1UsingExternalConfig(&dnssvcsv1.DnsSvcsV1Options{
				})
				Expect(dnsSvcsService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := dnsSvcsService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != dnsSvcsService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(dnsSvcsService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(dnsSvcsService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1UsingExternalConfig(&dnssvcsv1.DnsSvcsV1Options{
					URL: "https://testService/api",
				})
				Expect(dnsSvcsService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(dnsSvcsService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := dnsSvcsService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != dnsSvcsService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(dnsSvcsService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(dnsSvcsService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1UsingExternalConfig(&dnssvcsv1.DnsSvcsV1Options{
				})
				err := dnsSvcsService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(dnsSvcsService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(dnsSvcsService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := dnsSvcsService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != dnsSvcsService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(dnsSvcsService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(dnsSvcsService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"DNS_SVCS_URL": "https://dnssvcsv1/api",
				"DNS_SVCS_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1UsingExternalConfig(&dnssvcsv1.DnsSvcsV1Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(dnsSvcsService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"DNS_SVCS_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1UsingExternalConfig(&dnssvcsv1.DnsSvcsV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(dnsSvcsService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = dnssvcsv1.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`ListMonitors(listMonitorsOptions *ListMonitorsOptions) - Operation response error`, func() {
		listMonitorsPath := "/instances/testString/monitors"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listMonitorsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListMonitors with error: Operation response processing error`, func() {
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsSvcsService).ToNot(BeNil())

				// Construct an instance of the ListMonitorsOptions model
				listMonitorsOptionsModel := new(dnssvcsv1.ListMonitorsOptions)
				listMonitorsOptionsModel.InstanceID = core.StringPtr("testString")
				listMonitorsOptionsModel.XCorrelationID = core.StringPtr("testString")
				listMonitorsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := dnsSvcsService.ListMonitors(listMonitorsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				dnsSvcsService.EnableRetries(0, 0)
				result, response, operationErr = dnsSvcsService.ListMonitors(listMonitorsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ListMonitors(listMonitorsOptions *ListMonitorsOptions)`, func() {
		listMonitorsPath := "/instances/testString/monitors"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listMonitorsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"monitors": [{"id": "5365b73c-ce6f-4d6f-ad9f-d9c131b26370", "name": "healthcheck-monitor", "description": "Load balancer monitor for glb.example.com.", "type": "HTTPS", "port": 8080, "interval": 60, "retries": 2, "timeout": 5, "method": "GET", "path": "/health", "headers": [{"name": "Host", "value": ["origin.example.com"]}], "allow_insecure": false, "expected_codes": "200", "expected_body": "alive", "created_on": "2019-01-01T05:20:00.12345Z", "modified_on": "2019-01-01T05:20:00.12345Z"}], "offset": 1, "limit": 20, "count": 1, "total_count": 200, "first": {"href": "https://api.dns-svcs.cloud.ibm.com/v1/instances/434f6c3e-6014-4124-a61d-2e910bca19b1/dnszones/example.com:d04d3a7a-7f6d-47d4-b811-08c5478fa1a4/resource_records?limit=20"}, "next": {"href": "https://api.dns-svcs.cloud.ibm.com/v1/instances/434f6c3e-6014-4124-a61d-2e910bca19b1/dnszones/example.com:d04d3a7a-7f6d-47d4-b811-08c5478fa1a4/resource_records?offset=20&limit=20"}}`)
				}))
			})
			It(`Invoke ListMonitors successfully`, func() {
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsSvcsService).ToNot(BeNil())
				dnsSvcsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := dnsSvcsService.ListMonitors(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListMonitorsOptions model
				listMonitorsOptionsModel := new(dnssvcsv1.ListMonitorsOptions)
				listMonitorsOptionsModel.InstanceID = core.StringPtr("testString")
				listMonitorsOptionsModel.XCorrelationID = core.StringPtr("testString")
				listMonitorsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = dnsSvcsService.ListMonitors(listMonitorsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = dnsSvcsService.ListMonitorsWithContext(ctx, listMonitorsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				dnsSvcsService.DisableRetries()
				result, response, operationErr = dnsSvcsService.ListMonitors(listMonitorsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = dnsSvcsService.ListMonitorsWithContext(ctx, listMonitorsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke ListMonitors with error: Operation validation and request error`, func() {
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsSvcsService).ToNot(BeNil())

				// Construct an instance of the ListMonitorsOptions model
				listMonitorsOptionsModel := new(dnssvcsv1.ListMonitorsOptions)
				listMonitorsOptionsModel.InstanceID = core.StringPtr("testString")
				listMonitorsOptionsModel.XCorrelationID = core.StringPtr("testString")
				listMonitorsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := dnsSvcsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := dnsSvcsService.ListMonitors(listMonitorsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListMonitorsOptions model with no property values
				listMonitorsOptionsModelNew := new(dnssvcsv1.ListMonitorsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = dnsSvcsService.ListMonitors(listMonitorsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateMonitor(createMonitorOptions *CreateMonitorOptions) - Operation response error`, func() {
		createMonitorPath := "/instances/testString/monitors"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createMonitorPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateMonitor with error: Operation response processing error`, func() {
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsSvcsService).ToNot(BeNil())

				// Construct an instance of the HealthcheckHeader model
				healthcheckHeaderModel := new(dnssvcsv1.HealthcheckHeader)
				healthcheckHeaderModel.Name = core.StringPtr("Host")
				healthcheckHeaderModel.Value = []string{"origin.example.com"}

				// Construct an instance of the CreateMonitorOptions model
				createMonitorOptionsModel := new(dnssvcsv1.CreateMonitorOptions)
				createMonitorOptionsModel.InstanceID = core.StringPtr("testString")
				createMonitorOptionsModel.Name = core.StringPtr("healthcheck-monitor")
				createMonitorOptionsModel.Description = core.StringPtr("Load balancer monitor for glb.example.com.")
				createMonitorOptionsModel.Type = core.StringPtr("HTTPS")
				createMonitorOptionsModel.Port = core.Int64Ptr(int64(8080))
				createMonitorOptionsModel.Interval = core.Int64Ptr(int64(60))
				createMonitorOptionsModel.Retries = core.Int64Ptr(int64(2))
				createMonitorOptionsModel.Timeout = core.Int64Ptr(int64(5))
				createMonitorOptionsModel.Method = core.StringPtr("GET")
				createMonitorOptionsModel.Path = core.StringPtr("/health")
				createMonitorOptionsModel.HeadersVar = []dnssvcsv1.HealthcheckHeader{*healthcheckHeaderModel}
				createMonitorOptionsModel.AllowInsecure = core.BoolPtr(false)
				createMonitorOptionsModel.ExpectedCodes = core.StringPtr("200")
				createMonitorOptionsModel.ExpectedBody = core.StringPtr("alive")
				createMonitorOptionsModel.XCorrelationID = core.StringPtr("testString")
				createMonitorOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := dnsSvcsService.CreateMonitor(createMonitorOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				dnsSvcsService.EnableRetries(0, 0)
				result, response, operationErr = dnsSvcsService.CreateMonitor(createMonitorOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`CreateMonitor(createMonitorOptions *CreateMonitorOptions)`, func() {
		createMonitorPath := "/instances/testString/monitors"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createMonitorPath))
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

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "5365b73c-ce6f-4d6f-ad9f-d9c131b26370", "name": "healthcheck-monitor", "description": "Load balancer monitor for glb.example.com.", "type": "HTTPS", "port": 8080, "interval": 60, "retries": 2, "timeout": 5, "method": "GET", "path": "/health", "headers": [{"name": "Host", "value": ["origin.example.com"]}], "allow_insecure": false, "expected_codes": "200", "expected_body": "alive", "created_on": "2019-01-01T05:20:00.12345Z", "modified_on": "2019-01-01T05:20:00.12345Z"}`)
				}))
			})
			It(`Invoke CreateMonitor successfully`, func() {
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsSvcsService).ToNot(BeNil())
				dnsSvcsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := dnsSvcsService.CreateMonitor(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the HealthcheckHeader model
				healthcheckHeaderModel := new(dnssvcsv1.HealthcheckHeader)
				healthcheckHeaderModel.Name = core.StringPtr("Host")
				healthcheckHeaderModel.Value = []string{"origin.example.com"}

				// Construct an instance of the CreateMonitorOptions model
				createMonitorOptionsModel := new(dnssvcsv1.CreateMonitorOptions)
				createMonitorOptionsModel.InstanceID = core.StringPtr("testString")
				createMonitorOptionsModel.Name = core.StringPtr("healthcheck-monitor")
				createMonitorOptionsModel.Description = core.StringPtr("Load balancer monitor for glb.example.com.")
				createMonitorOptionsModel.Type = core.StringPtr("HTTPS")
				createMonitorOptionsModel.Port = core.Int64Ptr(int64(8080))
				createMonitorOptionsModel.Interval = core.Int64Ptr(int64(60))
				createMonitorOptionsModel.Retries = core.Int64Ptr(int64(2))
				createMonitorOptionsModel.Timeout = core.Int64Ptr(int64(5))
				createMonitorOptionsModel.Method = core.StringPtr("GET")
				createMonitorOptionsModel.Path = core.StringPtr("/health")
				createMonitorOptionsModel.HeadersVar = []dnssvcsv1.HealthcheckHeader{*healthcheckHeaderModel}
				createMonitorOptionsModel.AllowInsecure = core.BoolPtr(false)
				createMonitorOptionsModel.ExpectedCodes = core.StringPtr("200")
				createMonitorOptionsModel.ExpectedBody = core.StringPtr("alive")
				createMonitorOptionsModel.XCorrelationID = core.StringPtr("testString")
				createMonitorOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = dnsSvcsService.CreateMonitor(createMonitorOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = dnsSvcsService.CreateMonitorWithContext(ctx, createMonitorOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				dnsSvcsService.DisableRetries()
				result, response, operationErr = dnsSvcsService.CreateMonitor(createMonitorOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = dnsSvcsService.CreateMonitorWithContext(ctx, createMonitorOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke CreateMonitor with error: Operation validation and request error`, func() {
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsSvcsService).ToNot(BeNil())

				// Construct an instance of the HealthcheckHeader model
				healthcheckHeaderModel := new(dnssvcsv1.HealthcheckHeader)
				healthcheckHeaderModel.Name = core.StringPtr("Host")
				healthcheckHeaderModel.Value = []string{"origin.example.com"}

				// Construct an instance of the CreateMonitorOptions model
				createMonitorOptionsModel := new(dnssvcsv1.CreateMonitorOptions)
				createMonitorOptionsModel.InstanceID = core.StringPtr("testString")
				createMonitorOptionsModel.Name = core.StringPtr("healthcheck-monitor")
				createMonitorOptionsModel.Description = core.StringPtr("Load balancer monitor for glb.example.com.")
				createMonitorOptionsModel.Type = core.StringPtr("HTTPS")
				createMonitorOptionsModel.Port = core.Int64Ptr(int64(8080))
				createMonitorOptionsModel.Interval = core.Int64Ptr(int64(60))
				createMonitorOptionsModel.Retries = core.Int64Ptr(int64(2))
				createMonitorOptionsModel.Timeout = core.Int64Ptr(int64(5))
				createMonitorOptionsModel.Method = core.StringPtr("GET")
				createMonitorOptionsModel.Path = core.StringPtr("/health")
				createMonitorOptionsModel.HeadersVar = []dnssvcsv1.HealthcheckHeader{*healthcheckHeaderModel}
				createMonitorOptionsModel.AllowInsecure = core.BoolPtr(false)
				createMonitorOptionsModel.ExpectedCodes = core.StringPtr("200")
				createMonitorOptionsModel.ExpectedBody = core.StringPtr("alive")
				createMonitorOptionsModel.XCorrelationID = core.StringPtr("testString")
				createMonitorOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := dnsSvcsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := dnsSvcsService.CreateMonitor(createMonitorOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateMonitorOptions model with no property values
				createMonitorOptionsModelNew := new(dnssvcsv1.CreateMonitorOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = dnsSvcsService.CreateMonitor(createMonitorOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`DeleteMonitor(deleteMonitorOptions *DeleteMonitorOptions)`, func() {
		deleteMonitorPath := "/instances/testString/monitors/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteMonitorPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteMonitor successfully`, func() {
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsSvcsService).ToNot(BeNil())
				dnsSvcsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				response, operationErr := dnsSvcsService.DeleteMonitor(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteMonitorOptions model
				deleteMonitorOptionsModel := new(dnssvcsv1.DeleteMonitorOptions)
				deleteMonitorOptionsModel.InstanceID = core.StringPtr("testString")
				deleteMonitorOptionsModel.MonitorID = core.StringPtr("testString")
				deleteMonitorOptionsModel.XCorrelationID = core.StringPtr("testString")
				deleteMonitorOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = dnsSvcsService.DeleteMonitor(deleteMonitorOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Disable retries and test again
				dnsSvcsService.DisableRetries()
				response, operationErr = dnsSvcsService.DeleteMonitor(deleteMonitorOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteMonitor with error: Operation validation and request error`, func() {
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsSvcsService).ToNot(BeNil())

				// Construct an instance of the DeleteMonitorOptions model
				deleteMonitorOptionsModel := new(dnssvcsv1.DeleteMonitorOptions)
				deleteMonitorOptionsModel.InstanceID = core.StringPtr("testString")
				deleteMonitorOptionsModel.MonitorID = core.StringPtr("testString")
				deleteMonitorOptionsModel.XCorrelationID = core.StringPtr("testString")
				deleteMonitorOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := dnsSvcsService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := dnsSvcsService.DeleteMonitor(deleteMonitorOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteMonitorOptions model with no property values
				deleteMonitorOptionsModelNew := new(dnssvcsv1.DeleteMonitorOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = dnsSvcsService.DeleteMonitor(deleteMonitorOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetMonitor(getMonitorOptions *GetMonitorOptions) - Operation response error`, func() {
		getMonitorPath := "/instances/testString/monitors/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getMonitorPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetMonitor with error: Operation response processing error`, func() {
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsSvcsService).ToNot(BeNil())

				// Construct an instance of the GetMonitorOptions model
				getMonitorOptionsModel := new(dnssvcsv1.GetMonitorOptions)
				getMonitorOptionsModel.InstanceID = core.StringPtr("testString")
				getMonitorOptionsModel.MonitorID = core.StringPtr("testString")
				getMonitorOptionsModel.XCorrelationID = core.StringPtr("testString")
				getMonitorOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := dnsSvcsService.GetMonitor(getMonitorOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				dnsSvcsService.EnableRetries(0, 0)
				result, response, operationErr = dnsSvcsService.GetMonitor(getMonitorOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetMonitor(getMonitorOptions *GetMonitorOptions)`, func() {
		getMonitorPath := "/instances/testString/monitors/testString"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getMonitorPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "5365b73c-ce6f-4d6f-ad9f-d9c131b26370", "name": "healthcheck-monitor", "description": "Load balancer monitor for glb.example.com.", "type": "HTTPS", "port": 8080, "interval": 60, "retries": 2, "timeout": 5, "method": "GET", "path": "/health", "headers": [{"name": "Host", "value": ["origin.example.com"]}], "allow_insecure": false, "expected_codes": "200", "expected_body": "alive", "created_on": "2019-01-01T05:20:00.12345Z", "modified_on": "2019-01-01T05:20:00.12345Z"}`)
				}))
			})
			It(`Invoke GetMonitor successfully`, func() {
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsSvcsService).ToNot(BeNil())
				dnsSvcsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := dnsSvcsService.GetMonitor(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetMonitorOptions model
				getMonitorOptionsModel := new(dnssvcsv1.GetMonitorOptions)
				getMonitorOptionsModel.InstanceID = core.StringPtr("testString")
				getMonitorOptionsModel.MonitorID = core.StringPtr("testString")
				getMonitorOptionsModel.XCorrelationID = core.StringPtr("testString")
				getMonitorOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = dnsSvcsService.GetMonitor(getMonitorOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = dnsSvcsService.GetMonitorWithContext(ctx, getMonitorOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				dnsSvcsService.DisableRetries()
				result, response, operationErr = dnsSvcsService.GetMonitor(getMonitorOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = dnsSvcsService.GetMonitorWithContext(ctx, getMonitorOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke GetMonitor with error: Operation validation and request error`, func() {
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsSvcsService).ToNot(BeNil())

				// Construct an instance of the GetMonitorOptions model
				getMonitorOptionsModel := new(dnssvcsv1.GetMonitorOptions)
				getMonitorOptionsModel.InstanceID = core.StringPtr("testString")
				getMonitorOptionsModel.MonitorID = core.StringPtr("testString")
				getMonitorOptionsModel.XCorrelationID = core.StringPtr("testString")
				getMonitorOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := dnsSvcsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := dnsSvcsService.GetMonitor(getMonitorOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetMonitorOptions model with no property values
				getMonitorOptionsModelNew := new(dnssvcsv1.GetMonitorOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = dnsSvcsService.GetMonitor(getMonitorOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateMonitor(updateMonitorOptions *UpdateMonitorOptions) - Operation response error`, func() {
		updateMonitorPath := "/instances/testString/monitors/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateMonitorPath))
					Expect(req.Method).To(Equal("PUT"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateMonitor with error: Operation response processing error`, func() {
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsSvcsService).ToNot(BeNil())

				// Construct an instance of the HealthcheckHeader model
				healthcheckHeaderModel := new(dnssvcsv1.HealthcheckHeader)
				healthcheckHeaderModel.Name = core.StringPtr("Host")
				healthcheckHeaderModel.Value = []string{"origin.example.com"}

				// Construct an instance of the UpdateMonitorOptions model
				updateMonitorOptionsModel := new(dnssvcsv1.UpdateMonitorOptions)
				updateMonitorOptionsModel.InstanceID = core.StringPtr("testString")
				updateMonitorOptionsModel.MonitorID = core.StringPtr("testString")
				updateMonitorOptionsModel.Name = core.StringPtr("healthcheck-monitor")
				updateMonitorOptionsModel.Description = core.StringPtr("Load balancer monitor for glb.example.com.")
				updateMonitorOptionsModel.Type = core.StringPtr("HTTPS")
				updateMonitorOptionsModel.Port = core.Int64Ptr(int64(8080))
				updateMonitorOptionsModel.Interval = core.Int64Ptr(int64(60))
				updateMonitorOptionsModel.Retries = core.Int64Ptr(int64(2))
				updateMonitorOptionsModel.Timeout = core.Int64Ptr(int64(5))
				updateMonitorOptionsModel.Method = core.StringPtr("GET")
				updateMonitorOptionsModel.Path = core.StringPtr("/health")
				updateMonitorOptionsModel.HeadersVar = []dnssvcsv1.HealthcheckHeader{*healthcheckHeaderModel}
				updateMonitorOptionsModel.AllowInsecure = core.BoolPtr(false)
				updateMonitorOptionsModel.ExpectedCodes = core.StringPtr("200")
				updateMonitorOptionsModel.ExpectedBody = core.StringPtr("alive")
				updateMonitorOptionsModel.XCorrelationID = core.StringPtr("testString")
				updateMonitorOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := dnsSvcsService.UpdateMonitor(updateMonitorOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				dnsSvcsService.EnableRetries(0, 0)
				result, response, operationErr = dnsSvcsService.UpdateMonitor(updateMonitorOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UpdateMonitor(updateMonitorOptions *UpdateMonitorOptions)`, func() {
		updateMonitorPath := "/instances/testString/monitors/testString"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateMonitorPath))
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

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "5365b73c-ce6f-4d6f-ad9f-d9c131b26370", "name": "healthcheck-monitor", "description": "Load balancer monitor for glb.example.com.", "type": "HTTPS", "port": 8080, "interval": 60, "retries": 2, "timeout": 5, "method": "GET", "path": "/health", "headers": [{"name": "Host", "value": ["origin.example.com"]}], "allow_insecure": false, "expected_codes": "200", "expected_body": "alive", "created_on": "2019-01-01T05:20:00.12345Z", "modified_on": "2019-01-01T05:20:00.12345Z"}`)
				}))
			})
			It(`Invoke UpdateMonitor successfully`, func() {
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsSvcsService).ToNot(BeNil())
				dnsSvcsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := dnsSvcsService.UpdateMonitor(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the HealthcheckHeader model
				healthcheckHeaderModel := new(dnssvcsv1.HealthcheckHeader)
				healthcheckHeaderModel.Name = core.StringPtr("Host")
				healthcheckHeaderModel.Value = []string{"origin.example.com"}

				// Construct an instance of the UpdateMonitorOptions model
				updateMonitorOptionsModel := new(dnssvcsv1.UpdateMonitorOptions)
				updateMonitorOptionsModel.InstanceID = core.StringPtr("testString")
				updateMonitorOptionsModel.MonitorID = core.StringPtr("testString")
				updateMonitorOptionsModel.Name = core.StringPtr("healthcheck-monitor")
				updateMonitorOptionsModel.Description = core.StringPtr("Load balancer monitor for glb.example.com.")
				updateMonitorOptionsModel.Type = core.StringPtr("HTTPS")
				updateMonitorOptionsModel.Port = core.Int64Ptr(int64(8080))
				updateMonitorOptionsModel.Interval = core.Int64Ptr(int64(60))
				updateMonitorOptionsModel.Retries = core.Int64Ptr(int64(2))
				updateMonitorOptionsModel.Timeout = core.Int64Ptr(int64(5))
				updateMonitorOptionsModel.Method = core.StringPtr("GET")
				updateMonitorOptionsModel.Path = core.StringPtr("/health")
				updateMonitorOptionsModel.HeadersVar = []dnssvcsv1.HealthcheckHeader{*healthcheckHeaderModel}
				updateMonitorOptionsModel.AllowInsecure = core.BoolPtr(false)
				updateMonitorOptionsModel.ExpectedCodes = core.StringPtr("200")
				updateMonitorOptionsModel.ExpectedBody = core.StringPtr("alive")
				updateMonitorOptionsModel.XCorrelationID = core.StringPtr("testString")
				updateMonitorOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = dnsSvcsService.UpdateMonitor(updateMonitorOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = dnsSvcsService.UpdateMonitorWithContext(ctx, updateMonitorOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				dnsSvcsService.DisableRetries()
				result, response, operationErr = dnsSvcsService.UpdateMonitor(updateMonitorOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = dnsSvcsService.UpdateMonitorWithContext(ctx, updateMonitorOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke UpdateMonitor with error: Operation validation and request error`, func() {
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsSvcsService).ToNot(BeNil())

				// Construct an instance of the HealthcheckHeader model
				healthcheckHeaderModel := new(dnssvcsv1.HealthcheckHeader)
				healthcheckHeaderModel.Name = core.StringPtr("Host")
				healthcheckHeaderModel.Value = []string{"origin.example.com"}

				// Construct an instance of the UpdateMonitorOptions model
				updateMonitorOptionsModel := new(dnssvcsv1.UpdateMonitorOptions)
				updateMonitorOptionsModel.InstanceID = core.StringPtr("testString")
				updateMonitorOptionsModel.MonitorID = core.StringPtr("testString")
				updateMonitorOptionsModel.Name = core.StringPtr("healthcheck-monitor")
				updateMonitorOptionsModel.Description = core.StringPtr("Load balancer monitor for glb.example.com.")
				updateMonitorOptionsModel.Type = core.StringPtr("HTTPS")
				updateMonitorOptionsModel.Port = core.Int64Ptr(int64(8080))
				updateMonitorOptionsModel.Interval = core.Int64Ptr(int64(60))
				updateMonitorOptionsModel.Retries = core.Int64Ptr(int64(2))
				updateMonitorOptionsModel.Timeout = core.Int64Ptr(int64(5))
				updateMonitorOptionsModel.Method = core.StringPtr("GET")
				updateMonitorOptionsModel.Path = core.StringPtr("/health")
				updateMonitorOptionsModel.HeadersVar = []dnssvcsv1.HealthcheckHeader{*healthcheckHeaderModel}
				updateMonitorOptionsModel.AllowInsecure = core.BoolPtr(false)
				updateMonitorOptionsModel.ExpectedCodes = core.StringPtr("200")
				updateMonitorOptionsModel.ExpectedBody = core.StringPtr("alive")
				updateMonitorOptionsModel.XCorrelationID = core.StringPtr("testString")
				updateMonitorOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := dnsSvcsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := dnsSvcsService.UpdateMonitor(updateMonitorOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateMonitorOptions model with no property values
				updateMonitorOptionsModelNew := new(dnssvcsv1.UpdateMonitorOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = dnsSvcsService.UpdateMonitor(updateMonitorOptionsModelNew)
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
		It(`Instantiate service client`, func() {
			dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(dnsSvcsService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(dnsSvcsService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
				URL: "https://dnssvcsv1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(dnsSvcsService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"DNS_SVCS_URL": "https://dnssvcsv1/api",
				"DNS_SVCS_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1UsingExternalConfig(&dnssvcsv1.DnsSvcsV1Options{
				})
				Expect(dnsSvcsService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := dnsSvcsService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != dnsSvcsService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(dnsSvcsService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(dnsSvcsService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1UsingExternalConfig(&dnssvcsv1.DnsSvcsV1Options{
					URL: "https://testService/api",
				})
				Expect(dnsSvcsService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(dnsSvcsService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := dnsSvcsService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != dnsSvcsService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(dnsSvcsService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(dnsSvcsService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1UsingExternalConfig(&dnssvcsv1.DnsSvcsV1Options{
				})
				err := dnsSvcsService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(dnsSvcsService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(dnsSvcsService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := dnsSvcsService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != dnsSvcsService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(dnsSvcsService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(dnsSvcsService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"DNS_SVCS_URL": "https://dnssvcsv1/api",
				"DNS_SVCS_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1UsingExternalConfig(&dnssvcsv1.DnsSvcsV1Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(dnsSvcsService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"DNS_SVCS_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1UsingExternalConfig(&dnssvcsv1.DnsSvcsV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(dnsSvcsService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = dnssvcsv1.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`ListCustomResolvers(listCustomResolversOptions *ListCustomResolversOptions) - Operation response error`, func() {
		listCustomResolversPath := "/instances/testString/custom_resolvers"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listCustomResolversPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListCustomResolvers with error: Operation response processing error`, func() {
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsSvcsService).ToNot(BeNil())

				// Construct an instance of the ListCustomResolversOptions model
				listCustomResolversOptionsModel := new(dnssvcsv1.ListCustomResolversOptions)
				listCustomResolversOptionsModel.InstanceID = core.StringPtr("testString")
				listCustomResolversOptionsModel.XCorrelationID = core.StringPtr("testString")
				listCustomResolversOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := dnsSvcsService.ListCustomResolvers(listCustomResolversOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				dnsSvcsService.EnableRetries(0, 0)
				result, response, operationErr = dnsSvcsService.ListCustomResolvers(listCustomResolversOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ListCustomResolvers(listCustomResolversOptions *ListCustomResolversOptions)`, func() {
		listCustomResolversPath := "/instances/testString/custom_resolvers"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listCustomResolversPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"custom_resolvers": [{"id": "5365b73c-ce6f-4d6f-ad9f-d9c131b26370", "name": "my-resolver", "description": "custom resolver", "enabled": false, "health": "HEALTHY", "locations": [{"id": "9a234ede-c2b6-4c39-bc27-d39ec139ecdb", "subnet_crn": "crn:v1:bluemix:public:is:us-south-1:a/01652b251c3ae2787110a995d8db0135::subnet:0716-b49ef064-0f89-4fb1-8212-135b12568f04", "enabled": true, "healthy": true, "dns_server_ip": "10.10.16.8"}], "created_on": "2019-01-01T12:00:00", "modified_on": "2019-01-01T12:00:00"}]}`)
				}))
			})
			It(`Invoke ListCustomResolvers successfully`, func() {
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsSvcsService).ToNot(BeNil())
				dnsSvcsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := dnsSvcsService.ListCustomResolvers(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListCustomResolversOptions model
				listCustomResolversOptionsModel := new(dnssvcsv1.ListCustomResolversOptions)
				listCustomResolversOptionsModel.InstanceID = core.StringPtr("testString")
				listCustomResolversOptionsModel.XCorrelationID = core.StringPtr("testString")
				listCustomResolversOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = dnsSvcsService.ListCustomResolvers(listCustomResolversOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = dnsSvcsService.ListCustomResolversWithContext(ctx, listCustomResolversOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				dnsSvcsService.DisableRetries()
				result, response, operationErr = dnsSvcsService.ListCustomResolvers(listCustomResolversOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = dnsSvcsService.ListCustomResolversWithContext(ctx, listCustomResolversOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke ListCustomResolvers with error: Operation validation and request error`, func() {
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsSvcsService).ToNot(BeNil())

				// Construct an instance of the ListCustomResolversOptions model
				listCustomResolversOptionsModel := new(dnssvcsv1.ListCustomResolversOptions)
				listCustomResolversOptionsModel.InstanceID = core.StringPtr("testString")
				listCustomResolversOptionsModel.XCorrelationID = core.StringPtr("testString")
				listCustomResolversOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := dnsSvcsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := dnsSvcsService.ListCustomResolvers(listCustomResolversOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListCustomResolversOptions model with no property values
				listCustomResolversOptionsModelNew := new(dnssvcsv1.ListCustomResolversOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = dnsSvcsService.ListCustomResolvers(listCustomResolversOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateCustomResolver(createCustomResolverOptions *CreateCustomResolverOptions) - Operation response error`, func() {
		createCustomResolverPath := "/instances/testString/custom_resolvers"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createCustomResolverPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateCustomResolver with error: Operation response processing error`, func() {
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsSvcsService).ToNot(BeNil())

				// Construct an instance of the LocationInput model
				locationInputModel := new(dnssvcsv1.LocationInput)
				locationInputModel.SubnetCrn = core.StringPtr("crn:v1:bluemix:public:is:us-south-1:a/01652b251c3ae2787110a995d8db0135::subnet:0716-b49ef064-0f89-4fb1-8212-135b12568f04")
				locationInputModel.Enabled = core.BoolPtr(false)

				// Construct an instance of the CreateCustomResolverOptions model
				createCustomResolverOptionsModel := new(dnssvcsv1.CreateCustomResolverOptions)
				createCustomResolverOptionsModel.InstanceID = core.StringPtr("testString")
				createCustomResolverOptionsModel.Name = core.StringPtr("my-resolver")
				createCustomResolverOptionsModel.Description = core.StringPtr("custom resolver")
				createCustomResolverOptionsModel.Locations = []dnssvcsv1.LocationInput{*locationInputModel}
				createCustomResolverOptionsModel.XCorrelationID = core.StringPtr("testString")
				createCustomResolverOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := dnsSvcsService.CreateCustomResolver(createCustomResolverOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				dnsSvcsService.EnableRetries(0, 0)
				result, response, operationErr = dnsSvcsService.CreateCustomResolver(createCustomResolverOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`CreateCustomResolver(createCustomResolverOptions *CreateCustomResolverOptions)`, func() {
		createCustomResolverPath := "/instances/testString/custom_resolvers"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createCustomResolverPath))
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

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "5365b73c-ce6f-4d6f-ad9f-d9c131b26370", "name": "my-resolver", "description": "custom resolver", "enabled": false, "health": "HEALTHY", "locations": [{"id": "9a234ede-c2b6-4c39-bc27-d39ec139ecdb", "subnet_crn": "crn:v1:bluemix:public:is:us-south-1:a/01652b251c3ae2787110a995d8db0135::subnet:0716-b49ef064-0f89-4fb1-8212-135b12568f04", "enabled": true, "healthy": true, "dns_server_ip": "10.10.16.8"}], "created_on": "2019-01-01T12:00:00", "modified_on": "2019-01-01T12:00:00"}`)
				}))
			})
			It(`Invoke CreateCustomResolver successfully`, func() {
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsSvcsService).ToNot(BeNil())
				dnsSvcsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := dnsSvcsService.CreateCustomResolver(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the LocationInput model
				locationInputModel := new(dnssvcsv1.LocationInput)
				locationInputModel.SubnetCrn = core.StringPtr("crn:v1:bluemix:public:is:us-south-1:a/01652b251c3ae2787110a995d8db0135::subnet:0716-b49ef064-0f89-4fb1-8212-135b12568f04")
				locationInputModel.Enabled = core.BoolPtr(false)

				// Construct an instance of the CreateCustomResolverOptions model
				createCustomResolverOptionsModel := new(dnssvcsv1.CreateCustomResolverOptions)
				createCustomResolverOptionsModel.InstanceID = core.StringPtr("testString")
				createCustomResolverOptionsModel.Name = core.StringPtr("my-resolver")
				createCustomResolverOptionsModel.Description = core.StringPtr("custom resolver")
				createCustomResolverOptionsModel.Locations = []dnssvcsv1.LocationInput{*locationInputModel}
				createCustomResolverOptionsModel.XCorrelationID = core.StringPtr("testString")
				createCustomResolverOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = dnsSvcsService.CreateCustomResolver(createCustomResolverOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = dnsSvcsService.CreateCustomResolverWithContext(ctx, createCustomResolverOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				dnsSvcsService.DisableRetries()
				result, response, operationErr = dnsSvcsService.CreateCustomResolver(createCustomResolverOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = dnsSvcsService.CreateCustomResolverWithContext(ctx, createCustomResolverOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke CreateCustomResolver with error: Operation validation and request error`, func() {
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsSvcsService).ToNot(BeNil())

				// Construct an instance of the LocationInput model
				locationInputModel := new(dnssvcsv1.LocationInput)
				locationInputModel.SubnetCrn = core.StringPtr("crn:v1:bluemix:public:is:us-south-1:a/01652b251c3ae2787110a995d8db0135::subnet:0716-b49ef064-0f89-4fb1-8212-135b12568f04")
				locationInputModel.Enabled = core.BoolPtr(false)

				// Construct an instance of the CreateCustomResolverOptions model
				createCustomResolverOptionsModel := new(dnssvcsv1.CreateCustomResolverOptions)
				createCustomResolverOptionsModel.InstanceID = core.StringPtr("testString")
				createCustomResolverOptionsModel.Name = core.StringPtr("my-resolver")
				createCustomResolverOptionsModel.Description = core.StringPtr("custom resolver")
				createCustomResolverOptionsModel.Locations = []dnssvcsv1.LocationInput{*locationInputModel}
				createCustomResolverOptionsModel.XCorrelationID = core.StringPtr("testString")
				createCustomResolverOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := dnsSvcsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := dnsSvcsService.CreateCustomResolver(createCustomResolverOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateCustomResolverOptions model with no property values
				createCustomResolverOptionsModelNew := new(dnssvcsv1.CreateCustomResolverOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = dnsSvcsService.CreateCustomResolver(createCustomResolverOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`DeleteCustomResolver(deleteCustomResolverOptions *DeleteCustomResolverOptions)`, func() {
		deleteCustomResolverPath := "/instances/testString/custom_resolvers/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteCustomResolverPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteCustomResolver successfully`, func() {
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsSvcsService).ToNot(BeNil())
				dnsSvcsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				response, operationErr := dnsSvcsService.DeleteCustomResolver(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteCustomResolverOptions model
				deleteCustomResolverOptionsModel := new(dnssvcsv1.DeleteCustomResolverOptions)
				deleteCustomResolverOptionsModel.InstanceID = core.StringPtr("testString")
				deleteCustomResolverOptionsModel.ResolverID = core.StringPtr("testString")
				deleteCustomResolverOptionsModel.XCorrelationID = core.StringPtr("testString")
				deleteCustomResolverOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = dnsSvcsService.DeleteCustomResolver(deleteCustomResolverOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Disable retries and test again
				dnsSvcsService.DisableRetries()
				response, operationErr = dnsSvcsService.DeleteCustomResolver(deleteCustomResolverOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteCustomResolver with error: Operation validation and request error`, func() {
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsSvcsService).ToNot(BeNil())

				// Construct an instance of the DeleteCustomResolverOptions model
				deleteCustomResolverOptionsModel := new(dnssvcsv1.DeleteCustomResolverOptions)
				deleteCustomResolverOptionsModel.InstanceID = core.StringPtr("testString")
				deleteCustomResolverOptionsModel.ResolverID = core.StringPtr("testString")
				deleteCustomResolverOptionsModel.XCorrelationID = core.StringPtr("testString")
				deleteCustomResolverOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := dnsSvcsService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := dnsSvcsService.DeleteCustomResolver(deleteCustomResolverOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteCustomResolverOptions model with no property values
				deleteCustomResolverOptionsModelNew := new(dnssvcsv1.DeleteCustomResolverOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = dnsSvcsService.DeleteCustomResolver(deleteCustomResolverOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetCustomResolver(getCustomResolverOptions *GetCustomResolverOptions) - Operation response error`, func() {
		getCustomResolverPath := "/instances/testString/custom_resolvers/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getCustomResolverPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetCustomResolver with error: Operation response processing error`, func() {
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsSvcsService).ToNot(BeNil())

				// Construct an instance of the GetCustomResolverOptions model
				getCustomResolverOptionsModel := new(dnssvcsv1.GetCustomResolverOptions)
				getCustomResolverOptionsModel.InstanceID = core.StringPtr("testString")
				getCustomResolverOptionsModel.ResolverID = core.StringPtr("testString")
				getCustomResolverOptionsModel.XCorrelationID = core.StringPtr("testString")
				getCustomResolverOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := dnsSvcsService.GetCustomResolver(getCustomResolverOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				dnsSvcsService.EnableRetries(0, 0)
				result, response, operationErr = dnsSvcsService.GetCustomResolver(getCustomResolverOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetCustomResolver(getCustomResolverOptions *GetCustomResolverOptions)`, func() {
		getCustomResolverPath := "/instances/testString/custom_resolvers/testString"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getCustomResolverPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "5365b73c-ce6f-4d6f-ad9f-d9c131b26370", "name": "my-resolver", "description": "custom resolver", "enabled": false, "health": "HEALTHY", "locations": [{"id": "9a234ede-c2b6-4c39-bc27-d39ec139ecdb", "subnet_crn": "crn:v1:bluemix:public:is:us-south-1:a/01652b251c3ae2787110a995d8db0135::subnet:0716-b49ef064-0f89-4fb1-8212-135b12568f04", "enabled": true, "healthy": true, "dns_server_ip": "10.10.16.8"}], "created_on": "2019-01-01T12:00:00", "modified_on": "2019-01-01T12:00:00"}`)
				}))
			})
			It(`Invoke GetCustomResolver successfully`, func() {
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsSvcsService).ToNot(BeNil())
				dnsSvcsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := dnsSvcsService.GetCustomResolver(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetCustomResolverOptions model
				getCustomResolverOptionsModel := new(dnssvcsv1.GetCustomResolverOptions)
				getCustomResolverOptionsModel.InstanceID = core.StringPtr("testString")
				getCustomResolverOptionsModel.ResolverID = core.StringPtr("testString")
				getCustomResolverOptionsModel.XCorrelationID = core.StringPtr("testString")
				getCustomResolverOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = dnsSvcsService.GetCustomResolver(getCustomResolverOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = dnsSvcsService.GetCustomResolverWithContext(ctx, getCustomResolverOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				dnsSvcsService.DisableRetries()
				result, response, operationErr = dnsSvcsService.GetCustomResolver(getCustomResolverOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = dnsSvcsService.GetCustomResolverWithContext(ctx, getCustomResolverOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke GetCustomResolver with error: Operation validation and request error`, func() {
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsSvcsService).ToNot(BeNil())

				// Construct an instance of the GetCustomResolverOptions model
				getCustomResolverOptionsModel := new(dnssvcsv1.GetCustomResolverOptions)
				getCustomResolverOptionsModel.InstanceID = core.StringPtr("testString")
				getCustomResolverOptionsModel.ResolverID = core.StringPtr("testString")
				getCustomResolverOptionsModel.XCorrelationID = core.StringPtr("testString")
				getCustomResolverOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := dnsSvcsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := dnsSvcsService.GetCustomResolver(getCustomResolverOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetCustomResolverOptions model with no property values
				getCustomResolverOptionsModelNew := new(dnssvcsv1.GetCustomResolverOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = dnsSvcsService.GetCustomResolver(getCustomResolverOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateCustomResolver(updateCustomResolverOptions *UpdateCustomResolverOptions) - Operation response error`, func() {
		updateCustomResolverPath := "/instances/testString/custom_resolvers/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateCustomResolverPath))
					Expect(req.Method).To(Equal("PATCH"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateCustomResolver with error: Operation response processing error`, func() {
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsSvcsService).ToNot(BeNil())

				// Construct an instance of the UpdateCustomResolverOptions model
				updateCustomResolverOptionsModel := new(dnssvcsv1.UpdateCustomResolverOptions)
				updateCustomResolverOptionsModel.InstanceID = core.StringPtr("testString")
				updateCustomResolverOptionsModel.ResolverID = core.StringPtr("testString")
				updateCustomResolverOptionsModel.Name = core.StringPtr("my-resolver")
				updateCustomResolverOptionsModel.Description = core.StringPtr("custom resolver")
				updateCustomResolverOptionsModel.Enabled = core.BoolPtr(false)
				updateCustomResolverOptionsModel.XCorrelationID = core.StringPtr("testString")
				updateCustomResolverOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := dnsSvcsService.UpdateCustomResolver(updateCustomResolverOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				dnsSvcsService.EnableRetries(0, 0)
				result, response, operationErr = dnsSvcsService.UpdateCustomResolver(updateCustomResolverOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UpdateCustomResolver(updateCustomResolverOptions *UpdateCustomResolverOptions)`, func() {
		updateCustomResolverPath := "/instances/testString/custom_resolvers/testString"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateCustomResolverPath))
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

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "5365b73c-ce6f-4d6f-ad9f-d9c131b26370", "name": "my-resolver", "description": "custom resolver", "enabled": false, "health": "HEALTHY", "locations": [{"id": "9a234ede-c2b6-4c39-bc27-d39ec139ecdb", "subnet_crn": "crn:v1:bluemix:public:is:us-south-1:a/01652b251c3ae2787110a995d8db0135::subnet:0716-b49ef064-0f89-4fb1-8212-135b12568f04", "enabled": true, "healthy": true, "dns_server_ip": "10.10.16.8"}], "created_on": "2019-01-01T12:00:00", "modified_on": "2019-01-01T12:00:00"}`)
				}))
			})
			It(`Invoke UpdateCustomResolver successfully`, func() {
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsSvcsService).ToNot(BeNil())
				dnsSvcsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := dnsSvcsService.UpdateCustomResolver(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateCustomResolverOptions model
				updateCustomResolverOptionsModel := new(dnssvcsv1.UpdateCustomResolverOptions)
				updateCustomResolverOptionsModel.InstanceID = core.StringPtr("testString")
				updateCustomResolverOptionsModel.ResolverID = core.StringPtr("testString")
				updateCustomResolverOptionsModel.Name = core.StringPtr("my-resolver")
				updateCustomResolverOptionsModel.Description = core.StringPtr("custom resolver")
				updateCustomResolverOptionsModel.Enabled = core.BoolPtr(false)
				updateCustomResolverOptionsModel.XCorrelationID = core.StringPtr("testString")
				updateCustomResolverOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = dnsSvcsService.UpdateCustomResolver(updateCustomResolverOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = dnsSvcsService.UpdateCustomResolverWithContext(ctx, updateCustomResolverOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				dnsSvcsService.DisableRetries()
				result, response, operationErr = dnsSvcsService.UpdateCustomResolver(updateCustomResolverOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = dnsSvcsService.UpdateCustomResolverWithContext(ctx, updateCustomResolverOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke UpdateCustomResolver with error: Operation validation and request error`, func() {
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsSvcsService).ToNot(BeNil())

				// Construct an instance of the UpdateCustomResolverOptions model
				updateCustomResolverOptionsModel := new(dnssvcsv1.UpdateCustomResolverOptions)
				updateCustomResolverOptionsModel.InstanceID = core.StringPtr("testString")
				updateCustomResolverOptionsModel.ResolverID = core.StringPtr("testString")
				updateCustomResolverOptionsModel.Name = core.StringPtr("my-resolver")
				updateCustomResolverOptionsModel.Description = core.StringPtr("custom resolver")
				updateCustomResolverOptionsModel.Enabled = core.BoolPtr(false)
				updateCustomResolverOptionsModel.XCorrelationID = core.StringPtr("testString")
				updateCustomResolverOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := dnsSvcsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := dnsSvcsService.UpdateCustomResolver(updateCustomResolverOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateCustomResolverOptions model with no property values
				updateCustomResolverOptionsModelNew := new(dnssvcsv1.UpdateCustomResolverOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = dnsSvcsService.UpdateCustomResolver(updateCustomResolverOptionsModelNew)
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
		It(`Instantiate service client`, func() {
			dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(dnsSvcsService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(dnsSvcsService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
				URL: "https://dnssvcsv1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(dnsSvcsService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"DNS_SVCS_URL": "https://dnssvcsv1/api",
				"DNS_SVCS_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1UsingExternalConfig(&dnssvcsv1.DnsSvcsV1Options{
				})
				Expect(dnsSvcsService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := dnsSvcsService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != dnsSvcsService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(dnsSvcsService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(dnsSvcsService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1UsingExternalConfig(&dnssvcsv1.DnsSvcsV1Options{
					URL: "https://testService/api",
				})
				Expect(dnsSvcsService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(dnsSvcsService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := dnsSvcsService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != dnsSvcsService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(dnsSvcsService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(dnsSvcsService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1UsingExternalConfig(&dnssvcsv1.DnsSvcsV1Options{
				})
				err := dnsSvcsService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(dnsSvcsService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(dnsSvcsService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := dnsSvcsService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != dnsSvcsService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(dnsSvcsService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(dnsSvcsService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"DNS_SVCS_URL": "https://dnssvcsv1/api",
				"DNS_SVCS_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1UsingExternalConfig(&dnssvcsv1.DnsSvcsV1Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(dnsSvcsService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"DNS_SVCS_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1UsingExternalConfig(&dnssvcsv1.DnsSvcsV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(dnsSvcsService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = dnssvcsv1.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`AddCustomResolverLocation(addCustomResolverLocationOptions *AddCustomResolverLocationOptions) - Operation response error`, func() {
		addCustomResolverLocationPath := "/instances/testString/custom_resolvers/testString/locations"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(addCustomResolverLocationPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke AddCustomResolverLocation with error: Operation response processing error`, func() {
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsSvcsService).ToNot(BeNil())

				// Construct an instance of the AddCustomResolverLocationOptions model
				addCustomResolverLocationOptionsModel := new(dnssvcsv1.AddCustomResolverLocationOptions)
				addCustomResolverLocationOptionsModel.InstanceID = core.StringPtr("testString")
				addCustomResolverLocationOptionsModel.ResolverID = core.StringPtr("testString")
				addCustomResolverLocationOptionsModel.SubnetCrn = core.StringPtr("crn:v1:bluemix:public:is:us-south-1:a/01652b251c3ae2787110a995d8db0135::subnet:0716-b49ef064-0f89-4fb1-8212-135b12568f04")
				addCustomResolverLocationOptionsModel.Enabled = core.BoolPtr(false)
				addCustomResolverLocationOptionsModel.XCorrelationID = core.StringPtr("testString")
				addCustomResolverLocationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := dnsSvcsService.AddCustomResolverLocation(addCustomResolverLocationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				dnsSvcsService.EnableRetries(0, 0)
				result, response, operationErr = dnsSvcsService.AddCustomResolverLocation(addCustomResolverLocationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`AddCustomResolverLocation(addCustomResolverLocationOptions *AddCustomResolverLocationOptions)`, func() {
		addCustomResolverLocationPath := "/instances/testString/custom_resolvers/testString/locations"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(addCustomResolverLocationPath))
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

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "9a234ede-c2b6-4c39-bc27-d39ec139ecdb", "subnet_crn": "crn:v1:bluemix:public:is:us-south-1:a/01652b251c3ae2787110a995d8db0135::subnet:0716-b49ef064-0f89-4fb1-8212-135b12568f04", "enabled": true, "healthy": true, "dns_server_ip": "10.10.16.8"}`)
				}))
			})
			It(`Invoke AddCustomResolverLocation successfully`, func() {
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsSvcsService).ToNot(BeNil())
				dnsSvcsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := dnsSvcsService.AddCustomResolverLocation(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the AddCustomResolverLocationOptions model
				addCustomResolverLocationOptionsModel := new(dnssvcsv1.AddCustomResolverLocationOptions)
				addCustomResolverLocationOptionsModel.InstanceID = core.StringPtr("testString")
				addCustomResolverLocationOptionsModel.ResolverID = core.StringPtr("testString")
				addCustomResolverLocationOptionsModel.SubnetCrn = core.StringPtr("crn:v1:bluemix:public:is:us-south-1:a/01652b251c3ae2787110a995d8db0135::subnet:0716-b49ef064-0f89-4fb1-8212-135b12568f04")
				addCustomResolverLocationOptionsModel.Enabled = core.BoolPtr(false)
				addCustomResolverLocationOptionsModel.XCorrelationID = core.StringPtr("testString")
				addCustomResolverLocationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = dnsSvcsService.AddCustomResolverLocation(addCustomResolverLocationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = dnsSvcsService.AddCustomResolverLocationWithContext(ctx, addCustomResolverLocationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				dnsSvcsService.DisableRetries()
				result, response, operationErr = dnsSvcsService.AddCustomResolverLocation(addCustomResolverLocationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = dnsSvcsService.AddCustomResolverLocationWithContext(ctx, addCustomResolverLocationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke AddCustomResolverLocation with error: Operation validation and request error`, func() {
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsSvcsService).ToNot(BeNil())

				// Construct an instance of the AddCustomResolverLocationOptions model
				addCustomResolverLocationOptionsModel := new(dnssvcsv1.AddCustomResolverLocationOptions)
				addCustomResolverLocationOptionsModel.InstanceID = core.StringPtr("testString")
				addCustomResolverLocationOptionsModel.ResolverID = core.StringPtr("testString")
				addCustomResolverLocationOptionsModel.SubnetCrn = core.StringPtr("crn:v1:bluemix:public:is:us-south-1:a/01652b251c3ae2787110a995d8db0135::subnet:0716-b49ef064-0f89-4fb1-8212-135b12568f04")
				addCustomResolverLocationOptionsModel.Enabled = core.BoolPtr(false)
				addCustomResolverLocationOptionsModel.XCorrelationID = core.StringPtr("testString")
				addCustomResolverLocationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := dnsSvcsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := dnsSvcsService.AddCustomResolverLocation(addCustomResolverLocationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the AddCustomResolverLocationOptions model with no property values
				addCustomResolverLocationOptionsModelNew := new(dnssvcsv1.AddCustomResolverLocationOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = dnsSvcsService.AddCustomResolverLocation(addCustomResolverLocationOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateCustomResolverLocation(updateCustomResolverLocationOptions *UpdateCustomResolverLocationOptions) - Operation response error`, func() {
		updateCustomResolverLocationPath := "/instances/testString/custom_resolvers/testString/locations/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateCustomResolverLocationPath))
					Expect(req.Method).To(Equal("PATCH"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateCustomResolverLocation with error: Operation response processing error`, func() {
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsSvcsService).ToNot(BeNil())

				// Construct an instance of the UpdateCustomResolverLocationOptions model
				updateCustomResolverLocationOptionsModel := new(dnssvcsv1.UpdateCustomResolverLocationOptions)
				updateCustomResolverLocationOptionsModel.InstanceID = core.StringPtr("testString")
				updateCustomResolverLocationOptionsModel.ResolverID = core.StringPtr("testString")
				updateCustomResolverLocationOptionsModel.LocationID = core.StringPtr("testString")
				updateCustomResolverLocationOptionsModel.Enabled = core.BoolPtr(false)
				updateCustomResolverLocationOptionsModel.SubnetCrn = core.StringPtr("crn:v1:bluemix:public:is:us-south-1:a/01652b251c3ae2787110a995d8db0135::subnet:0716-b49ef064-0f89-4fb1-8212-135b12568f04")
				updateCustomResolverLocationOptionsModel.XCorrelationID = core.StringPtr("testString")
				updateCustomResolverLocationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := dnsSvcsService.UpdateCustomResolverLocation(updateCustomResolverLocationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				dnsSvcsService.EnableRetries(0, 0)
				result, response, operationErr = dnsSvcsService.UpdateCustomResolverLocation(updateCustomResolverLocationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UpdateCustomResolverLocation(updateCustomResolverLocationOptions *UpdateCustomResolverLocationOptions)`, func() {
		updateCustomResolverLocationPath := "/instances/testString/custom_resolvers/testString/locations/testString"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateCustomResolverLocationPath))
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

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "9a234ede-c2b6-4c39-bc27-d39ec139ecdb", "subnet_crn": "crn:v1:bluemix:public:is:us-south-1:a/01652b251c3ae2787110a995d8db0135::subnet:0716-b49ef064-0f89-4fb1-8212-135b12568f04", "enabled": true, "healthy": true, "dns_server_ip": "10.10.16.8"}`)
				}))
			})
			It(`Invoke UpdateCustomResolverLocation successfully`, func() {
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsSvcsService).ToNot(BeNil())
				dnsSvcsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := dnsSvcsService.UpdateCustomResolverLocation(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateCustomResolverLocationOptions model
				updateCustomResolverLocationOptionsModel := new(dnssvcsv1.UpdateCustomResolverLocationOptions)
				updateCustomResolverLocationOptionsModel.InstanceID = core.StringPtr("testString")
				updateCustomResolverLocationOptionsModel.ResolverID = core.StringPtr("testString")
				updateCustomResolverLocationOptionsModel.LocationID = core.StringPtr("testString")
				updateCustomResolverLocationOptionsModel.Enabled = core.BoolPtr(false)
				updateCustomResolverLocationOptionsModel.SubnetCrn = core.StringPtr("crn:v1:bluemix:public:is:us-south-1:a/01652b251c3ae2787110a995d8db0135::subnet:0716-b49ef064-0f89-4fb1-8212-135b12568f04")
				updateCustomResolverLocationOptionsModel.XCorrelationID = core.StringPtr("testString")
				updateCustomResolverLocationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = dnsSvcsService.UpdateCustomResolverLocation(updateCustomResolverLocationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = dnsSvcsService.UpdateCustomResolverLocationWithContext(ctx, updateCustomResolverLocationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				dnsSvcsService.DisableRetries()
				result, response, operationErr = dnsSvcsService.UpdateCustomResolverLocation(updateCustomResolverLocationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = dnsSvcsService.UpdateCustomResolverLocationWithContext(ctx, updateCustomResolverLocationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke UpdateCustomResolverLocation with error: Operation validation and request error`, func() {
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsSvcsService).ToNot(BeNil())

				// Construct an instance of the UpdateCustomResolverLocationOptions model
				updateCustomResolverLocationOptionsModel := new(dnssvcsv1.UpdateCustomResolverLocationOptions)
				updateCustomResolverLocationOptionsModel.InstanceID = core.StringPtr("testString")
				updateCustomResolverLocationOptionsModel.ResolverID = core.StringPtr("testString")
				updateCustomResolverLocationOptionsModel.LocationID = core.StringPtr("testString")
				updateCustomResolverLocationOptionsModel.Enabled = core.BoolPtr(false)
				updateCustomResolverLocationOptionsModel.SubnetCrn = core.StringPtr("crn:v1:bluemix:public:is:us-south-1:a/01652b251c3ae2787110a995d8db0135::subnet:0716-b49ef064-0f89-4fb1-8212-135b12568f04")
				updateCustomResolverLocationOptionsModel.XCorrelationID = core.StringPtr("testString")
				updateCustomResolverLocationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := dnsSvcsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := dnsSvcsService.UpdateCustomResolverLocation(updateCustomResolverLocationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateCustomResolverLocationOptions model with no property values
				updateCustomResolverLocationOptionsModelNew := new(dnssvcsv1.UpdateCustomResolverLocationOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = dnsSvcsService.UpdateCustomResolverLocation(updateCustomResolverLocationOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`DeleteCustomResolverLocation(deleteCustomResolverLocationOptions *DeleteCustomResolverLocationOptions)`, func() {
		deleteCustomResolverLocationPath := "/instances/testString/custom_resolvers/testString/locations/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteCustomResolverLocationPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteCustomResolverLocation successfully`, func() {
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsSvcsService).ToNot(BeNil())
				dnsSvcsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				response, operationErr := dnsSvcsService.DeleteCustomResolverLocation(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteCustomResolverLocationOptions model
				deleteCustomResolverLocationOptionsModel := new(dnssvcsv1.DeleteCustomResolverLocationOptions)
				deleteCustomResolverLocationOptionsModel.InstanceID = core.StringPtr("testString")
				deleteCustomResolverLocationOptionsModel.ResolverID = core.StringPtr("testString")
				deleteCustomResolverLocationOptionsModel.LocationID = core.StringPtr("testString")
				deleteCustomResolverLocationOptionsModel.XCorrelationID = core.StringPtr("testString")
				deleteCustomResolverLocationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = dnsSvcsService.DeleteCustomResolverLocation(deleteCustomResolverLocationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Disable retries and test again
				dnsSvcsService.DisableRetries()
				response, operationErr = dnsSvcsService.DeleteCustomResolverLocation(deleteCustomResolverLocationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteCustomResolverLocation with error: Operation validation and request error`, func() {
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsSvcsService).ToNot(BeNil())

				// Construct an instance of the DeleteCustomResolverLocationOptions model
				deleteCustomResolverLocationOptionsModel := new(dnssvcsv1.DeleteCustomResolverLocationOptions)
				deleteCustomResolverLocationOptionsModel.InstanceID = core.StringPtr("testString")
				deleteCustomResolverLocationOptionsModel.ResolverID = core.StringPtr("testString")
				deleteCustomResolverLocationOptionsModel.LocationID = core.StringPtr("testString")
				deleteCustomResolverLocationOptionsModel.XCorrelationID = core.StringPtr("testString")
				deleteCustomResolverLocationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := dnsSvcsService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := dnsSvcsService.DeleteCustomResolverLocation(deleteCustomResolverLocationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteCustomResolverLocationOptions model with no property values
				deleteCustomResolverLocationOptionsModelNew := new(dnssvcsv1.DeleteCustomResolverLocationOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = dnsSvcsService.DeleteCustomResolverLocation(deleteCustomResolverLocationOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(dnsSvcsService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(dnsSvcsService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
				URL: "https://dnssvcsv1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(dnsSvcsService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"DNS_SVCS_URL": "https://dnssvcsv1/api",
				"DNS_SVCS_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1UsingExternalConfig(&dnssvcsv1.DnsSvcsV1Options{
				})
				Expect(dnsSvcsService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := dnsSvcsService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != dnsSvcsService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(dnsSvcsService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(dnsSvcsService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1UsingExternalConfig(&dnssvcsv1.DnsSvcsV1Options{
					URL: "https://testService/api",
				})
				Expect(dnsSvcsService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(dnsSvcsService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := dnsSvcsService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != dnsSvcsService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(dnsSvcsService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(dnsSvcsService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1UsingExternalConfig(&dnssvcsv1.DnsSvcsV1Options{
				})
				err := dnsSvcsService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(dnsSvcsService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(dnsSvcsService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := dnsSvcsService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != dnsSvcsService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(dnsSvcsService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(dnsSvcsService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"DNS_SVCS_URL": "https://dnssvcsv1/api",
				"DNS_SVCS_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1UsingExternalConfig(&dnssvcsv1.DnsSvcsV1Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(dnsSvcsService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"DNS_SVCS_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1UsingExternalConfig(&dnssvcsv1.DnsSvcsV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(dnsSvcsService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = dnssvcsv1.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`ListForwardingRules(listForwardingRulesOptions *ListForwardingRulesOptions) - Operation response error`, func() {
		listForwardingRulesPath := "/instances/testString/custom_resolvers/testString/forwarding_rules"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listForwardingRulesPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListForwardingRules with error: Operation response processing error`, func() {
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsSvcsService).ToNot(BeNil())

				// Construct an instance of the ListForwardingRulesOptions model
				listForwardingRulesOptionsModel := new(dnssvcsv1.ListForwardingRulesOptions)
				listForwardingRulesOptionsModel.InstanceID = core.StringPtr("testString")
				listForwardingRulesOptionsModel.ResolverID = core.StringPtr("testString")
				listForwardingRulesOptionsModel.XCorrelationID = core.StringPtr("testString")
				listForwardingRulesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := dnsSvcsService.ListForwardingRules(listForwardingRulesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				dnsSvcsService.EnableRetries(0, 0)
				result, response, operationErr = dnsSvcsService.ListForwardingRules(listForwardingRulesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ListForwardingRules(listForwardingRulesOptions *ListForwardingRulesOptions)`, func() {
		listForwardingRulesPath := "/instances/testString/custom_resolvers/testString/forwarding_rules"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listForwardingRulesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"forwarding_rules": [{"id": "5365b73c-ce6f-4d6f-ad9f-d9c131b26370", "description": "forwarding rule", "type": "zone", "match": "example.com", "forward_to": ["161.26.0.7"], "created_on": "2019-01-01T12:00:00", "modified_on": "2019-01-01T12:00:00"}]}`)
				}))
			})
			It(`Invoke ListForwardingRules successfully`, func() {
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsSvcsService).ToNot(BeNil())
				dnsSvcsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := dnsSvcsService.ListForwardingRules(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListForwardingRulesOptions model
				listForwardingRulesOptionsModel := new(dnssvcsv1.ListForwardingRulesOptions)
				listForwardingRulesOptionsModel.InstanceID = core.StringPtr("testString")
				listForwardingRulesOptionsModel.ResolverID = core.StringPtr("testString")
				listForwardingRulesOptionsModel.XCorrelationID = core.StringPtr("testString")
				listForwardingRulesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = dnsSvcsService.ListForwardingRules(listForwardingRulesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = dnsSvcsService.ListForwardingRulesWithContext(ctx, listForwardingRulesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				dnsSvcsService.DisableRetries()
				result, response, operationErr = dnsSvcsService.ListForwardingRules(listForwardingRulesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = dnsSvcsService.ListForwardingRulesWithContext(ctx, listForwardingRulesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke ListForwardingRules with error: Operation validation and request error`, func() {
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsSvcsService).ToNot(BeNil())

				// Construct an instance of the ListForwardingRulesOptions model
				listForwardingRulesOptionsModel := new(dnssvcsv1.ListForwardingRulesOptions)
				listForwardingRulesOptionsModel.InstanceID = core.StringPtr("testString")
				listForwardingRulesOptionsModel.ResolverID = core.StringPtr("testString")
				listForwardingRulesOptionsModel.XCorrelationID = core.StringPtr("testString")
				listForwardingRulesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := dnsSvcsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := dnsSvcsService.ListForwardingRules(listForwardingRulesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListForwardingRulesOptions model with no property values
				listForwardingRulesOptionsModelNew := new(dnssvcsv1.ListForwardingRulesOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = dnsSvcsService.ListForwardingRules(listForwardingRulesOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateForwardingRule(createForwardingRuleOptions *CreateForwardingRuleOptions) - Operation response error`, func() {
		createForwardingRulePath := "/instances/testString/custom_resolvers/testString/forwarding_rules"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createForwardingRulePath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateForwardingRule with error: Operation response processing error`, func() {
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsSvcsService).ToNot(BeNil())

				// Construct an instance of the CreateForwardingRuleOptions model
				createForwardingRuleOptionsModel := new(dnssvcsv1.CreateForwardingRuleOptions)
				createForwardingRuleOptionsModel.InstanceID = core.StringPtr("testString")
				createForwardingRuleOptionsModel.ResolverID = core.StringPtr("testString")
				createForwardingRuleOptionsModel.Description = core.StringPtr("forwarding rule")
				createForwardingRuleOptionsModel.Type = core.StringPtr("zone")
				createForwardingRuleOptionsModel.Match = core.StringPtr("example.com")
				createForwardingRuleOptionsModel.ForwardTo = []string{"161.26.0.7"}
				createForwardingRuleOptionsModel.XCorrelationID = core.StringPtr("testString")
				createForwardingRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := dnsSvcsService.CreateForwardingRule(createForwardingRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				dnsSvcsService.EnableRetries(0, 0)
				result, response, operationErr = dnsSvcsService.CreateForwardingRule(createForwardingRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`CreateForwardingRule(createForwardingRuleOptions *CreateForwardingRuleOptions)`, func() {
		createForwardingRulePath := "/instances/testString/custom_resolvers/testString/forwarding_rules"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createForwardingRulePath))
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

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "5365b73c-ce6f-4d6f-ad9f-d9c131b26370", "description": "forwarding rule", "type": "zone", "match": "example.com", "forward_to": ["161.26.0.7"], "created_on": "2019-01-01T12:00:00", "modified_on": "2019-01-01T12:00:00"}`)
				}))
			})
			It(`Invoke CreateForwardingRule successfully`, func() {
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsSvcsService).ToNot(BeNil())
				dnsSvcsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := dnsSvcsService.CreateForwardingRule(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateForwardingRuleOptions model
				createForwardingRuleOptionsModel := new(dnssvcsv1.CreateForwardingRuleOptions)
				createForwardingRuleOptionsModel.InstanceID = core.StringPtr("testString")
				createForwardingRuleOptionsModel.ResolverID = core.StringPtr("testString")
				createForwardingRuleOptionsModel.Description = core.StringPtr("forwarding rule")
				createForwardingRuleOptionsModel.Type = core.StringPtr("zone")
				createForwardingRuleOptionsModel.Match = core.StringPtr("example.com")
				createForwardingRuleOptionsModel.ForwardTo = []string{"161.26.0.7"}
				createForwardingRuleOptionsModel.XCorrelationID = core.StringPtr("testString")
				createForwardingRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = dnsSvcsService.CreateForwardingRule(createForwardingRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = dnsSvcsService.CreateForwardingRuleWithContext(ctx, createForwardingRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				dnsSvcsService.DisableRetries()
				result, response, operationErr = dnsSvcsService.CreateForwardingRule(createForwardingRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = dnsSvcsService.CreateForwardingRuleWithContext(ctx, createForwardingRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke CreateForwardingRule with error: Operation validation and request error`, func() {
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsSvcsService).ToNot(BeNil())

				// Construct an instance of the CreateForwardingRuleOptions model
				createForwardingRuleOptionsModel := new(dnssvcsv1.CreateForwardingRuleOptions)
				createForwardingRuleOptionsModel.InstanceID = core.StringPtr("testString")
				createForwardingRuleOptionsModel.ResolverID = core.StringPtr("testString")
				createForwardingRuleOptionsModel.Description = core.StringPtr("forwarding rule")
				createForwardingRuleOptionsModel.Type = core.StringPtr("zone")
				createForwardingRuleOptionsModel.Match = core.StringPtr("example.com")
				createForwardingRuleOptionsModel.ForwardTo = []string{"161.26.0.7"}
				createForwardingRuleOptionsModel.XCorrelationID = core.StringPtr("testString")
				createForwardingRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := dnsSvcsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := dnsSvcsService.CreateForwardingRule(createForwardingRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateForwardingRuleOptions model with no property values
				createForwardingRuleOptionsModelNew := new(dnssvcsv1.CreateForwardingRuleOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = dnsSvcsService.CreateForwardingRule(createForwardingRuleOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`DeleteForwardingRule(deleteForwardingRuleOptions *DeleteForwardingRuleOptions)`, func() {
		deleteForwardingRulePath := "/instances/testString/custom_resolvers/testString/forwarding_rules/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteForwardingRulePath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteForwardingRule successfully`, func() {
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsSvcsService).ToNot(BeNil())
				dnsSvcsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				response, operationErr := dnsSvcsService.DeleteForwardingRule(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteForwardingRuleOptions model
				deleteForwardingRuleOptionsModel := new(dnssvcsv1.DeleteForwardingRuleOptions)
				deleteForwardingRuleOptionsModel.InstanceID = core.StringPtr("testString")
				deleteForwardingRuleOptionsModel.ResolverID = core.StringPtr("testString")
				deleteForwardingRuleOptionsModel.RuleID = core.StringPtr("testString")
				deleteForwardingRuleOptionsModel.XCorrelationID = core.StringPtr("testString")
				deleteForwardingRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = dnsSvcsService.DeleteForwardingRule(deleteForwardingRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Disable retries and test again
				dnsSvcsService.DisableRetries()
				response, operationErr = dnsSvcsService.DeleteForwardingRule(deleteForwardingRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteForwardingRule with error: Operation validation and request error`, func() {
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsSvcsService).ToNot(BeNil())

				// Construct an instance of the DeleteForwardingRuleOptions model
				deleteForwardingRuleOptionsModel := new(dnssvcsv1.DeleteForwardingRuleOptions)
				deleteForwardingRuleOptionsModel.InstanceID = core.StringPtr("testString")
				deleteForwardingRuleOptionsModel.ResolverID = core.StringPtr("testString")
				deleteForwardingRuleOptionsModel.RuleID = core.StringPtr("testString")
				deleteForwardingRuleOptionsModel.XCorrelationID = core.StringPtr("testString")
				deleteForwardingRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := dnsSvcsService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := dnsSvcsService.DeleteForwardingRule(deleteForwardingRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteForwardingRuleOptions model with no property values
				deleteForwardingRuleOptionsModelNew := new(dnssvcsv1.DeleteForwardingRuleOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = dnsSvcsService.DeleteForwardingRule(deleteForwardingRuleOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetForwardingRule(getForwardingRuleOptions *GetForwardingRuleOptions) - Operation response error`, func() {
		getForwardingRulePath := "/instances/testString/custom_resolvers/testString/forwarding_rules/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getForwardingRulePath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetForwardingRule with error: Operation response processing error`, func() {
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsSvcsService).ToNot(BeNil())

				// Construct an instance of the GetForwardingRuleOptions model
				getForwardingRuleOptionsModel := new(dnssvcsv1.GetForwardingRuleOptions)
				getForwardingRuleOptionsModel.InstanceID = core.StringPtr("testString")
				getForwardingRuleOptionsModel.ResolverID = core.StringPtr("testString")
				getForwardingRuleOptionsModel.RuleID = core.StringPtr("testString")
				getForwardingRuleOptionsModel.XCorrelationID = core.StringPtr("testString")
				getForwardingRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := dnsSvcsService.GetForwardingRule(getForwardingRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				dnsSvcsService.EnableRetries(0, 0)
				result, response, operationErr = dnsSvcsService.GetForwardingRule(getForwardingRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetForwardingRule(getForwardingRuleOptions *GetForwardingRuleOptions)`, func() {
		getForwardingRulePath := "/instances/testString/custom_resolvers/testString/forwarding_rules/testString"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getForwardingRulePath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "5365b73c-ce6f-4d6f-ad9f-d9c131b26370", "description": "forwarding rule", "type": "zone", "match": "example.com", "forward_to": ["161.26.0.7"], "created_on": "2019-01-01T12:00:00", "modified_on": "2019-01-01T12:00:00"}`)
				}))
			})
			It(`Invoke GetForwardingRule successfully`, func() {
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsSvcsService).ToNot(BeNil())
				dnsSvcsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := dnsSvcsService.GetForwardingRule(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetForwardingRuleOptions model
				getForwardingRuleOptionsModel := new(dnssvcsv1.GetForwardingRuleOptions)
				getForwardingRuleOptionsModel.InstanceID = core.StringPtr("testString")
				getForwardingRuleOptionsModel.ResolverID = core.StringPtr("testString")
				getForwardingRuleOptionsModel.RuleID = core.StringPtr("testString")
				getForwardingRuleOptionsModel.XCorrelationID = core.StringPtr("testString")
				getForwardingRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = dnsSvcsService.GetForwardingRule(getForwardingRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = dnsSvcsService.GetForwardingRuleWithContext(ctx, getForwardingRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				dnsSvcsService.DisableRetries()
				result, response, operationErr = dnsSvcsService.GetForwardingRule(getForwardingRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = dnsSvcsService.GetForwardingRuleWithContext(ctx, getForwardingRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke GetForwardingRule with error: Operation validation and request error`, func() {
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsSvcsService).ToNot(BeNil())

				// Construct an instance of the GetForwardingRuleOptions model
				getForwardingRuleOptionsModel := new(dnssvcsv1.GetForwardingRuleOptions)
				getForwardingRuleOptionsModel.InstanceID = core.StringPtr("testString")
				getForwardingRuleOptionsModel.ResolverID = core.StringPtr("testString")
				getForwardingRuleOptionsModel.RuleID = core.StringPtr("testString")
				getForwardingRuleOptionsModel.XCorrelationID = core.StringPtr("testString")
				getForwardingRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := dnsSvcsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := dnsSvcsService.GetForwardingRule(getForwardingRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetForwardingRuleOptions model with no property values
				getForwardingRuleOptionsModelNew := new(dnssvcsv1.GetForwardingRuleOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = dnsSvcsService.GetForwardingRule(getForwardingRuleOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateForwardingRule(updateForwardingRuleOptions *UpdateForwardingRuleOptions) - Operation response error`, func() {
		updateForwardingRulePath := "/instances/testString/custom_resolvers/testString/forwarding_rules/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateForwardingRulePath))
					Expect(req.Method).To(Equal("PATCH"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateForwardingRule with error: Operation response processing error`, func() {
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsSvcsService).ToNot(BeNil())

				// Construct an instance of the UpdateForwardingRuleOptions model
				updateForwardingRuleOptionsModel := new(dnssvcsv1.UpdateForwardingRuleOptions)
				updateForwardingRuleOptionsModel.InstanceID = core.StringPtr("testString")
				updateForwardingRuleOptionsModel.ResolverID = core.StringPtr("testString")
				updateForwardingRuleOptionsModel.RuleID = core.StringPtr("testString")
				updateForwardingRuleOptionsModel.Description = core.StringPtr("forwarding rule")
				updateForwardingRuleOptionsModel.Match = core.StringPtr("example.com")
				updateForwardingRuleOptionsModel.ForwardTo = []string{"161.26.0.7"}
				updateForwardingRuleOptionsModel.XCorrelationID = core.StringPtr("testString")
				updateForwardingRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := dnsSvcsService.UpdateForwardingRule(updateForwardingRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				dnsSvcsService.EnableRetries(0, 0)
				result, response, operationErr = dnsSvcsService.UpdateForwardingRule(updateForwardingRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UpdateForwardingRule(updateForwardingRuleOptions *UpdateForwardingRuleOptions)`, func() {
		updateForwardingRulePath := "/instances/testString/custom_resolvers/testString/forwarding_rules/testString"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateForwardingRulePath))
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

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "5365b73c-ce6f-4d6f-ad9f-d9c131b26370", "description": "forwarding rule", "type": "zone", "match": "example.com", "forward_to": ["161.26.0.7"], "created_on": "2019-01-01T12:00:00", "modified_on": "2019-01-01T12:00:00"}`)
				}))
			})
			It(`Invoke UpdateForwardingRule successfully`, func() {
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsSvcsService).ToNot(BeNil())
				dnsSvcsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := dnsSvcsService.UpdateForwardingRule(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateForwardingRuleOptions model
				updateForwardingRuleOptionsModel := new(dnssvcsv1.UpdateForwardingRuleOptions)
				updateForwardingRuleOptionsModel.InstanceID = core.StringPtr("testString")
				updateForwardingRuleOptionsModel.ResolverID = core.StringPtr("testString")
				updateForwardingRuleOptionsModel.RuleID = core.StringPtr("testString")
				updateForwardingRuleOptionsModel.Description = core.StringPtr("forwarding rule")
				updateForwardingRuleOptionsModel.Match = core.StringPtr("example.com")
				updateForwardingRuleOptionsModel.ForwardTo = []string{"161.26.0.7"}
				updateForwardingRuleOptionsModel.XCorrelationID = core.StringPtr("testString")
				updateForwardingRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = dnsSvcsService.UpdateForwardingRule(updateForwardingRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = dnsSvcsService.UpdateForwardingRuleWithContext(ctx, updateForwardingRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				dnsSvcsService.DisableRetries()
				result, response, operationErr = dnsSvcsService.UpdateForwardingRule(updateForwardingRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = dnsSvcsService.UpdateForwardingRuleWithContext(ctx, updateForwardingRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke UpdateForwardingRule with error: Operation validation and request error`, func() {
				dnsSvcsService, serviceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsSvcsService).ToNot(BeNil())

				// Construct an instance of the UpdateForwardingRuleOptions model
				updateForwardingRuleOptionsModel := new(dnssvcsv1.UpdateForwardingRuleOptions)
				updateForwardingRuleOptionsModel.InstanceID = core.StringPtr("testString")
				updateForwardingRuleOptionsModel.ResolverID = core.StringPtr("testString")
				updateForwardingRuleOptionsModel.RuleID = core.StringPtr("testString")
				updateForwardingRuleOptionsModel.Description = core.StringPtr("forwarding rule")
				updateForwardingRuleOptionsModel.Match = core.StringPtr("example.com")
				updateForwardingRuleOptionsModel.ForwardTo = []string{"161.26.0.7"}
				updateForwardingRuleOptionsModel.XCorrelationID = core.StringPtr("testString")
				updateForwardingRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := dnsSvcsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := dnsSvcsService.UpdateForwardingRule(updateForwardingRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateForwardingRuleOptions model with no property values
				updateForwardingRuleOptionsModelNew := new(dnssvcsv1.UpdateForwardingRuleOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = dnsSvcsService.UpdateForwardingRule(updateForwardingRuleOptionsModelNew)
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
			dnsSvcsService, _ := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
				URL:           "http://dnssvcsv1modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
			})
			It(`Invoke NewAddCustomResolverLocationOptions successfully`, func() {
				// Construct an instance of the AddCustomResolverLocationOptions model
				instanceID := "testString"
				resolverID := "testString"
				addCustomResolverLocationOptionsModel := dnsSvcsService.NewAddCustomResolverLocationOptions(instanceID, resolverID)
				addCustomResolverLocationOptionsModel.SetInstanceID("testString")
				addCustomResolverLocationOptionsModel.SetResolverID("testString")
				addCustomResolverLocationOptionsModel.SetSubnetCrn("crn:v1:bluemix:public:is:us-south-1:a/01652b251c3ae2787110a995d8db0135::subnet:0716-b49ef064-0f89-4fb1-8212-135b12568f04")
				addCustomResolverLocationOptionsModel.SetEnabled(false)
				addCustomResolverLocationOptionsModel.SetXCorrelationID("testString")
				addCustomResolverLocationOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(addCustomResolverLocationOptionsModel).ToNot(BeNil())
				Expect(addCustomResolverLocationOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(addCustomResolverLocationOptionsModel.ResolverID).To(Equal(core.StringPtr("testString")))
				Expect(addCustomResolverLocationOptionsModel.SubnetCrn).To(Equal(core.StringPtr("crn:v1:bluemix:public:is:us-south-1:a/01652b251c3ae2787110a995d8db0135::subnet:0716-b49ef064-0f89-4fb1-8212-135b12568f04")))
				Expect(addCustomResolverLocationOptionsModel.Enabled).To(Equal(core.BoolPtr(false)))
				Expect(addCustomResolverLocationOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(addCustomResolverLocationOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateCustomResolverOptions successfully`, func() {
				// Construct an instance of the LocationInput model
				locationInputModel := new(dnssvcsv1.LocationInput)
				Expect(locationInputModel).ToNot(BeNil())
				locationInputModel.SubnetCrn = core.StringPtr("crn:v1:bluemix:public:is:us-south-1:a/01652b251c3ae2787110a995d8db0135::subnet:0716-b49ef064-0f89-4fb1-8212-135b12568f04")
				locationInputModel.Enabled = core.BoolPtr(false)
				Expect(locationInputModel.SubnetCrn).To(Equal(core.StringPtr("crn:v1:bluemix:public:is:us-south-1:a/01652b251c3ae2787110a995d8db0135::subnet:0716-b49ef064-0f89-4fb1-8212-135b12568f04")))
				Expect(locationInputModel.Enabled).To(Equal(core.BoolPtr(false)))

				// Construct an instance of the CreateCustomResolverOptions model
				instanceID := "testString"
				createCustomResolverOptionsModel := dnsSvcsService.NewCreateCustomResolverOptions(instanceID)
				createCustomResolverOptionsModel.SetInstanceID("testString")
				createCustomResolverOptionsModel.SetName("my-resolver")
				createCustomResolverOptionsModel.SetDescription("custom resolver")
				createCustomResolverOptionsModel.SetLocations([]dnssvcsv1.LocationInput{*locationInputModel})
				createCustomResolverOptionsModel.SetXCorrelationID("testString")
				createCustomResolverOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createCustomResolverOptionsModel).ToNot(BeNil())
				Expect(createCustomResolverOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(createCustomResolverOptionsModel.Name).To(Equal(core.StringPtr("my-resolver")))
				Expect(createCustomResolverOptionsModel.Description).To(Equal(core.StringPtr("custom resolver")))
				Expect(createCustomResolverOptionsModel.Locations).To(Equal([]dnssvcsv1.LocationInput{*locationInputModel}))
				Expect(createCustomResolverOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(createCustomResolverOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateDnszoneOptions successfully`, func() {
				// Construct an instance of the CreateDnszoneOptions model
				instanceID := "testString"
				createDnszoneOptionsModel := dnsSvcsService.NewCreateDnszoneOptions(instanceID)
				createDnszoneOptionsModel.SetInstanceID("testString")
				createDnszoneOptionsModel.SetName("example.com")
				createDnszoneOptionsModel.SetDescription("The DNS zone is used for VPCs in us-east region")
				createDnszoneOptionsModel.SetLabel("us-east")
				createDnszoneOptionsModel.SetXCorrelationID("testString")
				createDnszoneOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createDnszoneOptionsModel).ToNot(BeNil())
				Expect(createDnszoneOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(createDnszoneOptionsModel.Name).To(Equal(core.StringPtr("example.com")))
				Expect(createDnszoneOptionsModel.Description).To(Equal(core.StringPtr("The DNS zone is used for VPCs in us-east region")))
				Expect(createDnszoneOptionsModel.Label).To(Equal(core.StringPtr("us-east")))
				Expect(createDnszoneOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(createDnszoneOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateForwardingRuleOptions successfully`, func() {
				// Construct an instance of the CreateForwardingRuleOptions model
				instanceID := "testString"
				resolverID := "testString"
				createForwardingRuleOptionsModel := dnsSvcsService.NewCreateForwardingRuleOptions(instanceID, resolverID)
				createForwardingRuleOptionsModel.SetInstanceID("testString")
				createForwardingRuleOptionsModel.SetResolverID("testString")
				createForwardingRuleOptionsModel.SetDescription("forwarding rule")
				createForwardingRuleOptionsModel.SetType("zone")
				createForwardingRuleOptionsModel.SetMatch("example.com")
				createForwardingRuleOptionsModel.SetForwardTo([]string{"161.26.0.7"})
				createForwardingRuleOptionsModel.SetXCorrelationID("testString")
				createForwardingRuleOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createForwardingRuleOptionsModel).ToNot(BeNil())
				Expect(createForwardingRuleOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(createForwardingRuleOptionsModel.ResolverID).To(Equal(core.StringPtr("testString")))
				Expect(createForwardingRuleOptionsModel.Description).To(Equal(core.StringPtr("forwarding rule")))
				Expect(createForwardingRuleOptionsModel.Type).To(Equal(core.StringPtr("zone")))
				Expect(createForwardingRuleOptionsModel.Match).To(Equal(core.StringPtr("example.com")))
				Expect(createForwardingRuleOptionsModel.ForwardTo).To(Equal([]string{"161.26.0.7"}))
				Expect(createForwardingRuleOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(createForwardingRuleOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateLoadBalancerOptions successfully`, func() {
				// Construct an instance of the LoadBalancerAzPoolsItem model
				loadBalancerAzPoolsItemModel := new(dnssvcsv1.LoadBalancerAzPoolsItem)
				Expect(loadBalancerAzPoolsItemModel).ToNot(BeNil())
				loadBalancerAzPoolsItemModel.AvailabilityZone = core.StringPtr("us-south-1")
				loadBalancerAzPoolsItemModel.Pools = []string{"0fc0bb7c-2fab-476e-8b9b-40fa14bf8e3d"}
				Expect(loadBalancerAzPoolsItemModel.AvailabilityZone).To(Equal(core.StringPtr("us-south-1")))
				Expect(loadBalancerAzPoolsItemModel.Pools).To(Equal([]string{"0fc0bb7c-2fab-476e-8b9b-40fa14bf8e3d"}))

				// Construct an instance of the CreateLoadBalancerOptions model
				instanceID := "testString"
				dnszoneID := "testString"
				createLoadBalancerOptionsModel := dnsSvcsService.NewCreateLoadBalancerOptions(instanceID, dnszoneID)
				createLoadBalancerOptionsModel.SetInstanceID("testString")
				createLoadBalancerOptionsModel.SetDnszoneID("testString")
				createLoadBalancerOptionsModel.SetName("glb.example.com")
				createLoadBalancerOptionsModel.SetDescription("Load balancer for glb.example.com.")
				createLoadBalancerOptionsModel.SetEnabled(true)
				createLoadBalancerOptionsModel.SetTTL(int64(120))
				createLoadBalancerOptionsModel.SetFallbackPool("24ccf79a-4ae0-4769-b4c8-17f8f230072e")
				createLoadBalancerOptionsModel.SetDefaultPools([]string{"testString"})
				createLoadBalancerOptionsModel.SetAzPools([]dnssvcsv1.LoadBalancerAzPoolsItem{*loadBalancerAzPoolsItemModel})
				createLoadBalancerOptionsModel.SetXCorrelationID("testString")
				createLoadBalancerOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createLoadBalancerOptionsModel).ToNot(BeNil())
				Expect(createLoadBalancerOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(createLoadBalancerOptionsModel.DnszoneID).To(Equal(core.StringPtr("testString")))
				Expect(createLoadBalancerOptionsModel.Name).To(Equal(core.StringPtr("glb.example.com")))
				Expect(createLoadBalancerOptionsModel.Description).To(Equal(core.StringPtr("Load balancer for glb.example.com.")))
				Expect(createLoadBalancerOptionsModel.Enabled).To(Equal(core.BoolPtr(true)))
				Expect(createLoadBalancerOptionsModel.TTL).To(Equal(core.Int64Ptr(int64(120))))
				Expect(createLoadBalancerOptionsModel.FallbackPool).To(Equal(core.StringPtr("24ccf79a-4ae0-4769-b4c8-17f8f230072e")))
				Expect(createLoadBalancerOptionsModel.DefaultPools).To(Equal([]string{"testString"}))
				Expect(createLoadBalancerOptionsModel.AzPools).To(Equal([]dnssvcsv1.LoadBalancerAzPoolsItem{*loadBalancerAzPoolsItemModel}))
				Expect(createLoadBalancerOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(createLoadBalancerOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateMonitorOptions successfully`, func() {
				// Construct an instance of the HealthcheckHeader model
				healthcheckHeaderModel := new(dnssvcsv1.HealthcheckHeader)
				Expect(healthcheckHeaderModel).ToNot(BeNil())
				healthcheckHeaderModel.Name = core.StringPtr("Host")
				healthcheckHeaderModel.Value = []string{"origin.example.com"}
				Expect(healthcheckHeaderModel.Name).To(Equal(core.StringPtr("Host")))
				Expect(healthcheckHeaderModel.Value).To(Equal([]string{"origin.example.com"}))

				// Construct an instance of the CreateMonitorOptions model
				instanceID := "testString"
				createMonitorOptionsModel := dnsSvcsService.NewCreateMonitorOptions(instanceID)
				createMonitorOptionsModel.SetInstanceID("testString")
				createMonitorOptionsModel.SetName("healthcheck-monitor")
				createMonitorOptionsModel.SetDescription("Load balancer monitor for glb.example.com.")
				createMonitorOptionsModel.SetType("HTTPS")
				createMonitorOptionsModel.SetPort(int64(8080))
				createMonitorOptionsModel.SetInterval(int64(60))
				createMonitorOptionsModel.SetRetries(int64(2))
				createMonitorOptionsModel.SetTimeout(int64(5))
				createMonitorOptionsModel.SetMethod("GET")
				createMonitorOptionsModel.SetPath("/health")
				createMonitorOptionsModel.SetHeadersVar([]dnssvcsv1.HealthcheckHeader{*healthcheckHeaderModel})
				createMonitorOptionsModel.SetAllowInsecure(false)
				createMonitorOptionsModel.SetExpectedCodes("200")
				createMonitorOptionsModel.SetExpectedBody("alive")
				createMonitorOptionsModel.SetXCorrelationID("testString")
				createMonitorOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createMonitorOptionsModel).ToNot(BeNil())
				Expect(createMonitorOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(createMonitorOptionsModel.Name).To(Equal(core.StringPtr("healthcheck-monitor")))
				Expect(createMonitorOptionsModel.Description).To(Equal(core.StringPtr("Load balancer monitor for glb.example.com.")))
				Expect(createMonitorOptionsModel.Type).To(Equal(core.StringPtr("HTTPS")))
				Expect(createMonitorOptionsModel.Port).To(Equal(core.Int64Ptr(int64(8080))))
				Expect(createMonitorOptionsModel.Interval).To(Equal(core.Int64Ptr(int64(60))))
				Expect(createMonitorOptionsModel.Retries).To(Equal(core.Int64Ptr(int64(2))))
				Expect(createMonitorOptionsModel.Timeout).To(Equal(core.Int64Ptr(int64(5))))
				Expect(createMonitorOptionsModel.Method).To(Equal(core.StringPtr("GET")))
				Expect(createMonitorOptionsModel.Path).To(Equal(core.StringPtr("/health")))
				Expect(createMonitorOptionsModel.HeadersVar).To(Equal([]dnssvcsv1.HealthcheckHeader{*healthcheckHeaderModel}))
				Expect(createMonitorOptionsModel.AllowInsecure).To(Equal(core.BoolPtr(false)))
				Expect(createMonitorOptionsModel.ExpectedCodes).To(Equal(core.StringPtr("200")))
				Expect(createMonitorOptionsModel.ExpectedBody).To(Equal(core.StringPtr("alive")))
				Expect(createMonitorOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(createMonitorOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreatePermittedNetworkOptions successfully`, func() {
				// Construct an instance of the PermittedNetworkVpc model
				permittedNetworkVpcModel := new(dnssvcsv1.PermittedNetworkVpc)
				Expect(permittedNetworkVpcModel).ToNot(BeNil())
				permittedNetworkVpcModel.VpcCrn = core.StringPtr("crn:v1:bluemix:public:is:eu-de:a/bcf1865e99742d38d2d5fc3fb80a5496::vpc:6e6cc326-04d1-4c99-a289-efb3ae4193d6")
				Expect(permittedNetworkVpcModel.VpcCrn).To(Equal(core.StringPtr("crn:v1:bluemix:public:is:eu-de:a/bcf1865e99742d38d2d5fc3fb80a5496::vpc:6e6cc326-04d1-4c99-a289-efb3ae4193d6")))

				// Construct an instance of the CreatePermittedNetworkOptions model
				instanceID := "testString"
				dnszoneID := "testString"
				createPermittedNetworkOptionsModel := dnsSvcsService.NewCreatePermittedNetworkOptions(instanceID, dnszoneID)
				createPermittedNetworkOptionsModel.SetInstanceID("testString")
				createPermittedNetworkOptionsModel.SetDnszoneID("testString")
				createPermittedNetworkOptionsModel.SetType("vpc")
				createPermittedNetworkOptionsModel.SetPermittedNetwork(permittedNetworkVpcModel)
				createPermittedNetworkOptionsModel.SetXCorrelationID("testString")
				createPermittedNetworkOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createPermittedNetworkOptionsModel).ToNot(BeNil())
				Expect(createPermittedNetworkOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(createPermittedNetworkOptionsModel.DnszoneID).To(Equal(core.StringPtr("testString")))
				Expect(createPermittedNetworkOptionsModel.Type).To(Equal(core.StringPtr("vpc")))
				Expect(createPermittedNetworkOptionsModel.PermittedNetwork).To(Equal(permittedNetworkVpcModel))
				Expect(createPermittedNetworkOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(createPermittedNetworkOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreatePoolOptions successfully`, func() {
				// Construct an instance of the OriginInput model
				originInputModel := new(dnssvcsv1.OriginInput)
				Expect(originInputModel).ToNot(BeNil())
				originInputModel.Name = core.StringPtr("app-server-1")
				originInputModel.Description = core.StringPtr("description of the origin server")
				originInputModel.Address = core.StringPtr("10.10.16.8")
				originInputModel.Enabled = core.BoolPtr(true)
				Expect(originInputModel.Name).To(Equal(core.StringPtr("app-server-1")))
				Expect(originInputModel.Description).To(Equal(core.StringPtr("description of the origin server")))
				Expect(originInputModel.Address).To(Equal(core.StringPtr("10.10.16.8")))
				Expect(originInputModel.Enabled).To(Equal(core.BoolPtr(true)))

				// Construct an instance of the CreatePoolOptions model
				instanceID := "testString"
				createPoolOptionsModel := dnsSvcsService.NewCreatePoolOptions(instanceID)
				createPoolOptionsModel.SetInstanceID("testString")
				createPoolOptionsModel.SetName("dal10-az-pool")
				createPoolOptionsModel.SetDescription("Load balancer pool for dal10 availability zone.")
				createPoolOptionsModel.SetEnabled(true)
				createPoolOptionsModel.SetHealthyOriginsThreshold(int64(1))
				createPoolOptionsModel.SetOrigins([]dnssvcsv1.OriginInput{*originInputModel})
				createPoolOptionsModel.SetMonitor("7dd6841c-264e-11ea-88df-062967242a6a")
				createPoolOptionsModel.SetNotificationChannel("https://mywebsite.com/dns/webhook")
				createPoolOptionsModel.SetHealthcheckRegion("us-south")
				createPoolOptionsModel.SetHealthcheckSubnets([]string{"crn:v1:staging:public:is:us-south-1:a/01652b251c3ae2787110a995d8db0135::subnet:0716-b49ef064-0f89-4fb1-8212-135b12568f04"})
				createPoolOptionsModel.SetXCorrelationID("testString")
				createPoolOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createPoolOptionsModel).ToNot(BeNil())
				Expect(createPoolOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(createPoolOptionsModel.Name).To(Equal(core.StringPtr("dal10-az-pool")))
				Expect(createPoolOptionsModel.Description).To(Equal(core.StringPtr("Load balancer pool for dal10 availability zone.")))
				Expect(createPoolOptionsModel.Enabled).To(Equal(core.BoolPtr(true)))
				Expect(createPoolOptionsModel.HealthyOriginsThreshold).To(Equal(core.Int64Ptr(int64(1))))
				Expect(createPoolOptionsModel.Origins).To(Equal([]dnssvcsv1.OriginInput{*originInputModel}))
				Expect(createPoolOptionsModel.Monitor).To(Equal(core.StringPtr("7dd6841c-264e-11ea-88df-062967242a6a")))
				Expect(createPoolOptionsModel.NotificationChannel).To(Equal(core.StringPtr("https://mywebsite.com/dns/webhook")))
				Expect(createPoolOptionsModel.HealthcheckRegion).To(Equal(core.StringPtr("us-south")))
				Expect(createPoolOptionsModel.HealthcheckSubnets).To(Equal([]string{"crn:v1:staging:public:is:us-south-1:a/01652b251c3ae2787110a995d8db0135::subnet:0716-b49ef064-0f89-4fb1-8212-135b12568f04"}))
				Expect(createPoolOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(createPoolOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateResourceRecordOptions successfully`, func() {
				// Construct an instance of the ResourceRecordInputRdataRdataARecord model
				resourceRecordInputRdataModel := new(dnssvcsv1.ResourceRecordInputRdataRdataARecord)
				Expect(resourceRecordInputRdataModel).ToNot(BeNil())
				resourceRecordInputRdataModel.Ip = core.StringPtr("10.110.201.214")
				Expect(resourceRecordInputRdataModel.Ip).To(Equal(core.StringPtr("10.110.201.214")))

				// Construct an instance of the CreateResourceRecordOptions model
				instanceID := "testString"
				dnszoneID := "testString"
				createResourceRecordOptionsModel := dnsSvcsService.NewCreateResourceRecordOptions(instanceID, dnszoneID)
				createResourceRecordOptionsModel.SetInstanceID("testString")
				createResourceRecordOptionsModel.SetDnszoneID("testString")
				createResourceRecordOptionsModel.SetName("test.example.com")
				createResourceRecordOptionsModel.SetType("SRV")
				createResourceRecordOptionsModel.SetRdata(resourceRecordInputRdataModel)
				createResourceRecordOptionsModel.SetTTL(int64(120))
				createResourceRecordOptionsModel.SetService("_sip")
				createResourceRecordOptionsModel.SetProtocol("udp")
				createResourceRecordOptionsModel.SetXCorrelationID("testString")
				createResourceRecordOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createResourceRecordOptionsModel).ToNot(BeNil())
				Expect(createResourceRecordOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(createResourceRecordOptionsModel.DnszoneID).To(Equal(core.StringPtr("testString")))
				Expect(createResourceRecordOptionsModel.Name).To(Equal(core.StringPtr("test.example.com")))
				Expect(createResourceRecordOptionsModel.Type).To(Equal(core.StringPtr("SRV")))
				Expect(createResourceRecordOptionsModel.Rdata).To(Equal(resourceRecordInputRdataModel))
				Expect(createResourceRecordOptionsModel.TTL).To(Equal(core.Int64Ptr(int64(120))))
				Expect(createResourceRecordOptionsModel.Service).To(Equal(core.StringPtr("_sip")))
				Expect(createResourceRecordOptionsModel.Protocol).To(Equal(core.StringPtr("udp")))
				Expect(createResourceRecordOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(createResourceRecordOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteCustomResolverLocationOptions successfully`, func() {
				// Construct an instance of the DeleteCustomResolverLocationOptions model
				instanceID := "testString"
				resolverID := "testString"
				locationID := "testString"
				deleteCustomResolverLocationOptionsModel := dnsSvcsService.NewDeleteCustomResolverLocationOptions(instanceID, resolverID, locationID)
				deleteCustomResolverLocationOptionsModel.SetInstanceID("testString")
				deleteCustomResolverLocationOptionsModel.SetResolverID("testString")
				deleteCustomResolverLocationOptionsModel.SetLocationID("testString")
				deleteCustomResolverLocationOptionsModel.SetXCorrelationID("testString")
				deleteCustomResolverLocationOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteCustomResolverLocationOptionsModel).ToNot(BeNil())
				Expect(deleteCustomResolverLocationOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(deleteCustomResolverLocationOptionsModel.ResolverID).To(Equal(core.StringPtr("testString")))
				Expect(deleteCustomResolverLocationOptionsModel.LocationID).To(Equal(core.StringPtr("testString")))
				Expect(deleteCustomResolverLocationOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(deleteCustomResolverLocationOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteCustomResolverOptions successfully`, func() {
				// Construct an instance of the DeleteCustomResolverOptions model
				instanceID := "testString"
				resolverID := "testString"
				deleteCustomResolverOptionsModel := dnsSvcsService.NewDeleteCustomResolverOptions(instanceID, resolverID)
				deleteCustomResolverOptionsModel.SetInstanceID("testString")
				deleteCustomResolverOptionsModel.SetResolverID("testString")
				deleteCustomResolverOptionsModel.SetXCorrelationID("testString")
				deleteCustomResolverOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteCustomResolverOptionsModel).ToNot(BeNil())
				Expect(deleteCustomResolverOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(deleteCustomResolverOptionsModel.ResolverID).To(Equal(core.StringPtr("testString")))
				Expect(deleteCustomResolverOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(deleteCustomResolverOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteDnszoneOptions successfully`, func() {
				// Construct an instance of the DeleteDnszoneOptions model
				instanceID := "testString"
				dnszoneID := "testString"
				deleteDnszoneOptionsModel := dnsSvcsService.NewDeleteDnszoneOptions(instanceID, dnszoneID)
				deleteDnszoneOptionsModel.SetInstanceID("testString")
				deleteDnszoneOptionsModel.SetDnszoneID("testString")
				deleteDnszoneOptionsModel.SetXCorrelationID("testString")
				deleteDnszoneOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteDnszoneOptionsModel).ToNot(BeNil())
				Expect(deleteDnszoneOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(deleteDnszoneOptionsModel.DnszoneID).To(Equal(core.StringPtr("testString")))
				Expect(deleteDnszoneOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(deleteDnszoneOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteForwardingRuleOptions successfully`, func() {
				// Construct an instance of the DeleteForwardingRuleOptions model
				instanceID := "testString"
				resolverID := "testString"
				ruleID := "testString"
				deleteForwardingRuleOptionsModel := dnsSvcsService.NewDeleteForwardingRuleOptions(instanceID, resolverID, ruleID)
				deleteForwardingRuleOptionsModel.SetInstanceID("testString")
				deleteForwardingRuleOptionsModel.SetResolverID("testString")
				deleteForwardingRuleOptionsModel.SetRuleID("testString")
				deleteForwardingRuleOptionsModel.SetXCorrelationID("testString")
				deleteForwardingRuleOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteForwardingRuleOptionsModel).ToNot(BeNil())
				Expect(deleteForwardingRuleOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(deleteForwardingRuleOptionsModel.ResolverID).To(Equal(core.StringPtr("testString")))
				Expect(deleteForwardingRuleOptionsModel.RuleID).To(Equal(core.StringPtr("testString")))
				Expect(deleteForwardingRuleOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(deleteForwardingRuleOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteLoadBalancerOptions successfully`, func() {
				// Construct an instance of the DeleteLoadBalancerOptions model
				instanceID := "testString"
				dnszoneID := "testString"
				lbID := "testString"
				deleteLoadBalancerOptionsModel := dnsSvcsService.NewDeleteLoadBalancerOptions(instanceID, dnszoneID, lbID)
				deleteLoadBalancerOptionsModel.SetInstanceID("testString")
				deleteLoadBalancerOptionsModel.SetDnszoneID("testString")
				deleteLoadBalancerOptionsModel.SetLbID("testString")
				deleteLoadBalancerOptionsModel.SetXCorrelationID("testString")
				deleteLoadBalancerOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteLoadBalancerOptionsModel).ToNot(BeNil())
				Expect(deleteLoadBalancerOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(deleteLoadBalancerOptionsModel.DnszoneID).To(Equal(core.StringPtr("testString")))
				Expect(deleteLoadBalancerOptionsModel.LbID).To(Equal(core.StringPtr("testString")))
				Expect(deleteLoadBalancerOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(deleteLoadBalancerOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteMonitorOptions successfully`, func() {
				// Construct an instance of the DeleteMonitorOptions model
				instanceID := "testString"
				monitorID := "testString"
				deleteMonitorOptionsModel := dnsSvcsService.NewDeleteMonitorOptions(instanceID, monitorID)
				deleteMonitorOptionsModel.SetInstanceID("testString")
				deleteMonitorOptionsModel.SetMonitorID("testString")
				deleteMonitorOptionsModel.SetXCorrelationID("testString")
				deleteMonitorOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteMonitorOptionsModel).ToNot(BeNil())
				Expect(deleteMonitorOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(deleteMonitorOptionsModel.MonitorID).To(Equal(core.StringPtr("testString")))
				Expect(deleteMonitorOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(deleteMonitorOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeletePermittedNetworkOptions successfully`, func() {
				// Construct an instance of the DeletePermittedNetworkOptions model
				instanceID := "testString"
				dnszoneID := "testString"
				permittedNetworkID := "testString"
				deletePermittedNetworkOptionsModel := dnsSvcsService.NewDeletePermittedNetworkOptions(instanceID, dnszoneID, permittedNetworkID)
				deletePermittedNetworkOptionsModel.SetInstanceID("testString")
				deletePermittedNetworkOptionsModel.SetDnszoneID("testString")
				deletePermittedNetworkOptionsModel.SetPermittedNetworkID("testString")
				deletePermittedNetworkOptionsModel.SetXCorrelationID("testString")
				deletePermittedNetworkOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deletePermittedNetworkOptionsModel).ToNot(BeNil())
				Expect(deletePermittedNetworkOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(deletePermittedNetworkOptionsModel.DnszoneID).To(Equal(core.StringPtr("testString")))
				Expect(deletePermittedNetworkOptionsModel.PermittedNetworkID).To(Equal(core.StringPtr("testString")))
				Expect(deletePermittedNetworkOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(deletePermittedNetworkOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeletePoolOptions successfully`, func() {
				// Construct an instance of the DeletePoolOptions model
				instanceID := "testString"
				poolID := "testString"
				deletePoolOptionsModel := dnsSvcsService.NewDeletePoolOptions(instanceID, poolID)
				deletePoolOptionsModel.SetInstanceID("testString")
				deletePoolOptionsModel.SetPoolID("testString")
				deletePoolOptionsModel.SetXCorrelationID("testString")
				deletePoolOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deletePoolOptionsModel).ToNot(BeNil())
				Expect(deletePoolOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(deletePoolOptionsModel.PoolID).To(Equal(core.StringPtr("testString")))
				Expect(deletePoolOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(deletePoolOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteResourceRecordOptions successfully`, func() {
				// Construct an instance of the DeleteResourceRecordOptions model
				instanceID := "testString"
				dnszoneID := "testString"
				recordID := "testString"
				deleteResourceRecordOptionsModel := dnsSvcsService.NewDeleteResourceRecordOptions(instanceID, dnszoneID, recordID)
				deleteResourceRecordOptionsModel.SetInstanceID("testString")
				deleteResourceRecordOptionsModel.SetDnszoneID("testString")
				deleteResourceRecordOptionsModel.SetRecordID("testString")
				deleteResourceRecordOptionsModel.SetXCorrelationID("testString")
				deleteResourceRecordOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteResourceRecordOptionsModel).ToNot(BeNil())
				Expect(deleteResourceRecordOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(deleteResourceRecordOptionsModel.DnszoneID).To(Equal(core.StringPtr("testString")))
				Expect(deleteResourceRecordOptionsModel.RecordID).To(Equal(core.StringPtr("testString")))
				Expect(deleteResourceRecordOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(deleteResourceRecordOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewExportResourceRecordsOptions successfully`, func() {
				// Construct an instance of the ExportResourceRecordsOptions model
				instanceID := "testString"
				dnszoneID := "testString"
				exportResourceRecordsOptionsModel := dnsSvcsService.NewExportResourceRecordsOptions(instanceID, dnszoneID)
				exportResourceRecordsOptionsModel.SetInstanceID("testString")
				exportResourceRecordsOptionsModel.SetDnszoneID("testString")
				exportResourceRecordsOptionsModel.SetXCorrelationID("testString")
				exportResourceRecordsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(exportResourceRecordsOptionsModel).ToNot(BeNil())
				Expect(exportResourceRecordsOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(exportResourceRecordsOptionsModel.DnszoneID).To(Equal(core.StringPtr("testString")))
				Expect(exportResourceRecordsOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(exportResourceRecordsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetCustomResolverOptions successfully`, func() {
				// Construct an instance of the GetCustomResolverOptions model
				instanceID := "testString"
				resolverID := "testString"
				getCustomResolverOptionsModel := dnsSvcsService.NewGetCustomResolverOptions(instanceID, resolverID)
				getCustomResolverOptionsModel.SetInstanceID("testString")
				getCustomResolverOptionsModel.SetResolverID("testString")
				getCustomResolverOptionsModel.SetXCorrelationID("testString")
				getCustomResolverOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getCustomResolverOptionsModel).ToNot(BeNil())
				Expect(getCustomResolverOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(getCustomResolverOptionsModel.ResolverID).To(Equal(core.StringPtr("testString")))
				Expect(getCustomResolverOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(getCustomResolverOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetDnszoneOptions successfully`, func() {
				// Construct an instance of the GetDnszoneOptions model
				instanceID := "testString"
				dnszoneID := "testString"
				getDnszoneOptionsModel := dnsSvcsService.NewGetDnszoneOptions(instanceID, dnszoneID)
				getDnszoneOptionsModel.SetInstanceID("testString")
				getDnszoneOptionsModel.SetDnszoneID("testString")
				getDnszoneOptionsModel.SetXCorrelationID("testString")
				getDnszoneOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getDnszoneOptionsModel).ToNot(BeNil())
				Expect(getDnszoneOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(getDnszoneOptionsModel.DnszoneID).To(Equal(core.StringPtr("testString")))
				Expect(getDnszoneOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(getDnszoneOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetForwardingRuleOptions successfully`, func() {
				// Construct an instance of the GetForwardingRuleOptions model
				instanceID := "testString"
				resolverID := "testString"
				ruleID := "testString"
				getForwardingRuleOptionsModel := dnsSvcsService.NewGetForwardingRuleOptions(instanceID, resolverID, ruleID)
				getForwardingRuleOptionsModel.SetInstanceID("testString")
				getForwardingRuleOptionsModel.SetResolverID("testString")
				getForwardingRuleOptionsModel.SetRuleID("testString")
				getForwardingRuleOptionsModel.SetXCorrelationID("testString")
				getForwardingRuleOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getForwardingRuleOptionsModel).ToNot(BeNil())
				Expect(getForwardingRuleOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(getForwardingRuleOptionsModel.ResolverID).To(Equal(core.StringPtr("testString")))
				Expect(getForwardingRuleOptionsModel.RuleID).To(Equal(core.StringPtr("testString")))
				Expect(getForwardingRuleOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(getForwardingRuleOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetLoadBalancerOptions successfully`, func() {
				// Construct an instance of the GetLoadBalancerOptions model
				instanceID := "testString"
				dnszoneID := "testString"
				lbID := "testString"
				getLoadBalancerOptionsModel := dnsSvcsService.NewGetLoadBalancerOptions(instanceID, dnszoneID, lbID)
				getLoadBalancerOptionsModel.SetInstanceID("testString")
				getLoadBalancerOptionsModel.SetDnszoneID("testString")
				getLoadBalancerOptionsModel.SetLbID("testString")
				getLoadBalancerOptionsModel.SetXCorrelationID("testString")
				getLoadBalancerOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getLoadBalancerOptionsModel).ToNot(BeNil())
				Expect(getLoadBalancerOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(getLoadBalancerOptionsModel.DnszoneID).To(Equal(core.StringPtr("testString")))
				Expect(getLoadBalancerOptionsModel.LbID).To(Equal(core.StringPtr("testString")))
				Expect(getLoadBalancerOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(getLoadBalancerOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetMonitorOptions successfully`, func() {
				// Construct an instance of the GetMonitorOptions model
				instanceID := "testString"
				monitorID := "testString"
				getMonitorOptionsModel := dnsSvcsService.NewGetMonitorOptions(instanceID, monitorID)
				getMonitorOptionsModel.SetInstanceID("testString")
				getMonitorOptionsModel.SetMonitorID("testString")
				getMonitorOptionsModel.SetXCorrelationID("testString")
				getMonitorOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getMonitorOptionsModel).ToNot(BeNil())
				Expect(getMonitorOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(getMonitorOptionsModel.MonitorID).To(Equal(core.StringPtr("testString")))
				Expect(getMonitorOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(getMonitorOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetPermittedNetworkOptions successfully`, func() {
				// Construct an instance of the GetPermittedNetworkOptions model
				instanceID := "testString"
				dnszoneID := "testString"
				permittedNetworkID := "testString"
				getPermittedNetworkOptionsModel := dnsSvcsService.NewGetPermittedNetworkOptions(instanceID, dnszoneID, permittedNetworkID)
				getPermittedNetworkOptionsModel.SetInstanceID("testString")
				getPermittedNetworkOptionsModel.SetDnszoneID("testString")
				getPermittedNetworkOptionsModel.SetPermittedNetworkID("testString")
				getPermittedNetworkOptionsModel.SetXCorrelationID("testString")
				getPermittedNetworkOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getPermittedNetworkOptionsModel).ToNot(BeNil())
				Expect(getPermittedNetworkOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(getPermittedNetworkOptionsModel.DnszoneID).To(Equal(core.StringPtr("testString")))
				Expect(getPermittedNetworkOptionsModel.PermittedNetworkID).To(Equal(core.StringPtr("testString")))
				Expect(getPermittedNetworkOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(getPermittedNetworkOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetPoolOptions successfully`, func() {
				// Construct an instance of the GetPoolOptions model
				instanceID := "testString"
				poolID := "testString"
				getPoolOptionsModel := dnsSvcsService.NewGetPoolOptions(instanceID, poolID)
				getPoolOptionsModel.SetInstanceID("testString")
				getPoolOptionsModel.SetPoolID("testString")
				getPoolOptionsModel.SetXCorrelationID("testString")
				getPoolOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getPoolOptionsModel).ToNot(BeNil())
				Expect(getPoolOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(getPoolOptionsModel.PoolID).To(Equal(core.StringPtr("testString")))
				Expect(getPoolOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(getPoolOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetResourceRecordOptions successfully`, func() {
				// Construct an instance of the GetResourceRecordOptions model
				instanceID := "testString"
				dnszoneID := "testString"
				recordID := "testString"
				getResourceRecordOptionsModel := dnsSvcsService.NewGetResourceRecordOptions(instanceID, dnszoneID, recordID)
				getResourceRecordOptionsModel.SetInstanceID("testString")
				getResourceRecordOptionsModel.SetDnszoneID("testString")
				getResourceRecordOptionsModel.SetRecordID("testString")
				getResourceRecordOptionsModel.SetXCorrelationID("testString")
				getResourceRecordOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getResourceRecordOptionsModel).ToNot(BeNil())
				Expect(getResourceRecordOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(getResourceRecordOptionsModel.DnszoneID).To(Equal(core.StringPtr("testString")))
				Expect(getResourceRecordOptionsModel.RecordID).To(Equal(core.StringPtr("testString")))
				Expect(getResourceRecordOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(getResourceRecordOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewImportResourceRecordsOptions successfully`, func() {
				// Construct an instance of the ImportResourceRecordsOptions model
				instanceID := "testString"
				dnszoneID := "testString"
				importResourceRecordsOptionsModel := dnsSvcsService.NewImportResourceRecordsOptions(instanceID, dnszoneID)
				importResourceRecordsOptionsModel.SetInstanceID("testString")
				importResourceRecordsOptionsModel.SetDnszoneID("testString")
				importResourceRecordsOptionsModel.SetFile(CreateMockReader("This is a mock file."))
				importResourceRecordsOptionsModel.SetFileContentType("testString")
				importResourceRecordsOptionsModel.SetXCorrelationID("testString")
				importResourceRecordsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(importResourceRecordsOptionsModel).ToNot(BeNil())
				Expect(importResourceRecordsOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(importResourceRecordsOptionsModel.DnszoneID).To(Equal(core.StringPtr("testString")))
				Expect(importResourceRecordsOptionsModel.File).To(Equal(CreateMockReader("This is a mock file.")))
				Expect(importResourceRecordsOptionsModel.FileContentType).To(Equal(core.StringPtr("testString")))
				Expect(importResourceRecordsOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(importResourceRecordsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListCustomResolversOptions successfully`, func() {
				// Construct an instance of the ListCustomResolversOptions model
				instanceID := "testString"
				listCustomResolversOptionsModel := dnsSvcsService.NewListCustomResolversOptions(instanceID)
				listCustomResolversOptionsModel.SetInstanceID("testString")
				listCustomResolversOptionsModel.SetXCorrelationID("testString")
				listCustomResolversOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listCustomResolversOptionsModel).ToNot(BeNil())
				Expect(listCustomResolversOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(listCustomResolversOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(listCustomResolversOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListDnszonesOptions successfully`, func() {
				// Construct an instance of the ListDnszonesOptions model
				instanceID := "testString"
				listDnszonesOptionsModel := dnsSvcsService.NewListDnszonesOptions(instanceID)
				listDnszonesOptionsModel.SetInstanceID("testString")
				listDnszonesOptionsModel.SetXCorrelationID("testString")
				listDnszonesOptionsModel.SetOffset(int64(38))
				listDnszonesOptionsModel.SetLimit(int64(200))
				listDnszonesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listDnszonesOptionsModel).ToNot(BeNil())
				Expect(listDnszonesOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(listDnszonesOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(listDnszonesOptionsModel.Offset).To(Equal(core.Int64Ptr(int64(38))))
				Expect(listDnszonesOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(200))))
				Expect(listDnszonesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListForwardingRulesOptions successfully`, func() {
				// Construct an instance of the ListForwardingRulesOptions model
				instanceID := "testString"
				resolverID := "testString"
				listForwardingRulesOptionsModel := dnsSvcsService.NewListForwardingRulesOptions(instanceID, resolverID)
				listForwardingRulesOptionsModel.SetInstanceID("testString")
				listForwardingRulesOptionsModel.SetResolverID("testString")
				listForwardingRulesOptionsModel.SetXCorrelationID("testString")
				listForwardingRulesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listForwardingRulesOptionsModel).ToNot(BeNil())
				Expect(listForwardingRulesOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(listForwardingRulesOptionsModel.ResolverID).To(Equal(core.StringPtr("testString")))
				Expect(listForwardingRulesOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(listForwardingRulesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListLoadBalancersOptions successfully`, func() {
				// Construct an instance of the ListLoadBalancersOptions model
				instanceID := "testString"
				dnszoneID := "testString"
				listLoadBalancersOptionsModel := dnsSvcsService.NewListLoadBalancersOptions(instanceID, dnszoneID)
				listLoadBalancersOptionsModel.SetInstanceID("testString")
				listLoadBalancersOptionsModel.SetDnszoneID("testString")
				listLoadBalancersOptionsModel.SetXCorrelationID("testString")
				listLoadBalancersOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listLoadBalancersOptionsModel).ToNot(BeNil())
				Expect(listLoadBalancersOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(listLoadBalancersOptionsModel.DnszoneID).To(Equal(core.StringPtr("testString")))
				Expect(listLoadBalancersOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(listLoadBalancersOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListMonitorsOptions successfully`, func() {
				// Construct an instance of the ListMonitorsOptions model
				instanceID := "testString"
				listMonitorsOptionsModel := dnsSvcsService.NewListMonitorsOptions(instanceID)
				listMonitorsOptionsModel.SetInstanceID("testString")
				listMonitorsOptionsModel.SetXCorrelationID("testString")
				listMonitorsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listMonitorsOptionsModel).ToNot(BeNil())
				Expect(listMonitorsOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(listMonitorsOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(listMonitorsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListPermittedNetworksOptions successfully`, func() {
				// Construct an instance of the ListPermittedNetworksOptions model
				instanceID := "testString"
				dnszoneID := "testString"
				listPermittedNetworksOptionsModel := dnsSvcsService.NewListPermittedNetworksOptions(instanceID, dnszoneID)
				listPermittedNetworksOptionsModel.SetInstanceID("testString")
				listPermittedNetworksOptionsModel.SetDnszoneID("testString")
				listPermittedNetworksOptionsModel.SetXCorrelationID("testString")
				listPermittedNetworksOptionsModel.SetOffset(int64(38))
				listPermittedNetworksOptionsModel.SetLimit(int64(200))
				listPermittedNetworksOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listPermittedNetworksOptionsModel).ToNot(BeNil())
				Expect(listPermittedNetworksOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(listPermittedNetworksOptionsModel.DnszoneID).To(Equal(core.StringPtr("testString")))
				Expect(listPermittedNetworksOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(listPermittedNetworksOptionsModel.Offset).To(Equal(core.Int64Ptr(int64(38))))
				Expect(listPermittedNetworksOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(200))))
				Expect(listPermittedNetworksOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListPoolsOptions successfully`, func() {
				// Construct an instance of the ListPoolsOptions model
				instanceID := "testString"
				listPoolsOptionsModel := dnsSvcsService.NewListPoolsOptions(instanceID)
				listPoolsOptionsModel.SetInstanceID("testString")
				listPoolsOptionsModel.SetXCorrelationID("testString")
				listPoolsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listPoolsOptionsModel).ToNot(BeNil())
				Expect(listPoolsOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(listPoolsOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(listPoolsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListResourceRecordsOptions successfully`, func() {
				// Construct an instance of the ListResourceRecordsOptions model
				instanceID := "testString"
				dnszoneID := "testString"
				listResourceRecordsOptionsModel := dnsSvcsService.NewListResourceRecordsOptions(instanceID, dnszoneID)
				listResourceRecordsOptionsModel.SetInstanceID("testString")
				listResourceRecordsOptionsModel.SetDnszoneID("testString")
				listResourceRecordsOptionsModel.SetXCorrelationID("testString")
				listResourceRecordsOptionsModel.SetOffset(int64(38))
				listResourceRecordsOptionsModel.SetLimit(int64(200))
				listResourceRecordsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listResourceRecordsOptionsModel).ToNot(BeNil())
				Expect(listResourceRecordsOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(listResourceRecordsOptionsModel.DnszoneID).To(Equal(core.StringPtr("testString")))
				Expect(listResourceRecordsOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(listResourceRecordsOptionsModel.Offset).To(Equal(core.Int64Ptr(int64(38))))
				Expect(listResourceRecordsOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(200))))
				Expect(listResourceRecordsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateCustomResolverLocationOptions successfully`, func() {
				// Construct an instance of the UpdateCustomResolverLocationOptions model
				instanceID := "testString"
				resolverID := "testString"
				locationID := "testString"
				updateCustomResolverLocationOptionsModel := dnsSvcsService.NewUpdateCustomResolverLocationOptions(instanceID, resolverID, locationID)
				updateCustomResolverLocationOptionsModel.SetInstanceID("testString")
				updateCustomResolverLocationOptionsModel.SetResolverID("testString")
				updateCustomResolverLocationOptionsModel.SetLocationID("testString")
				updateCustomResolverLocationOptionsModel.SetEnabled(false)
				updateCustomResolverLocationOptionsModel.SetSubnetCrn("crn:v1:bluemix:public:is:us-south-1:a/01652b251c3ae2787110a995d8db0135::subnet:0716-b49ef064-0f89-4fb1-8212-135b12568f04")
				updateCustomResolverLocationOptionsModel.SetXCorrelationID("testString")
				updateCustomResolverLocationOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateCustomResolverLocationOptionsModel).ToNot(BeNil())
				Expect(updateCustomResolverLocationOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(updateCustomResolverLocationOptionsModel.ResolverID).To(Equal(core.StringPtr("testString")))
				Expect(updateCustomResolverLocationOptionsModel.LocationID).To(Equal(core.StringPtr("testString")))
				Expect(updateCustomResolverLocationOptionsModel.Enabled).To(Equal(core.BoolPtr(false)))
				Expect(updateCustomResolverLocationOptionsModel.SubnetCrn).To(Equal(core.StringPtr("crn:v1:bluemix:public:is:us-south-1:a/01652b251c3ae2787110a995d8db0135::subnet:0716-b49ef064-0f89-4fb1-8212-135b12568f04")))
				Expect(updateCustomResolverLocationOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(updateCustomResolverLocationOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateCustomResolverOptions successfully`, func() {
				// Construct an instance of the UpdateCustomResolverOptions model
				instanceID := "testString"
				resolverID := "testString"
				updateCustomResolverOptionsModel := dnsSvcsService.NewUpdateCustomResolverOptions(instanceID, resolverID)
				updateCustomResolverOptionsModel.SetInstanceID("testString")
				updateCustomResolverOptionsModel.SetResolverID("testString")
				updateCustomResolverOptionsModel.SetName("my-resolver")
				updateCustomResolverOptionsModel.SetDescription("custom resolver")
				updateCustomResolverOptionsModel.SetEnabled(false)
				updateCustomResolverOptionsModel.SetXCorrelationID("testString")
				updateCustomResolverOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateCustomResolverOptionsModel).ToNot(BeNil())
				Expect(updateCustomResolverOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(updateCustomResolverOptionsModel.ResolverID).To(Equal(core.StringPtr("testString")))
				Expect(updateCustomResolverOptionsModel.Name).To(Equal(core.StringPtr("my-resolver")))
				Expect(updateCustomResolverOptionsModel.Description).To(Equal(core.StringPtr("custom resolver")))
				Expect(updateCustomResolverOptionsModel.Enabled).To(Equal(core.BoolPtr(false)))
				Expect(updateCustomResolverOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(updateCustomResolverOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateDnszoneOptions successfully`, func() {
				// Construct an instance of the UpdateDnszoneOptions model
				instanceID := "testString"
				dnszoneID := "testString"
				updateDnszoneOptionsModel := dnsSvcsService.NewUpdateDnszoneOptions(instanceID, dnszoneID)
				updateDnszoneOptionsModel.SetInstanceID("testString")
				updateDnszoneOptionsModel.SetDnszoneID("testString")
				updateDnszoneOptionsModel.SetDescription("The DNS zone is used for VPCs in us-east region")
				updateDnszoneOptionsModel.SetLabel("us-east")
				updateDnszoneOptionsModel.SetXCorrelationID("testString")
				updateDnszoneOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateDnszoneOptionsModel).ToNot(BeNil())
				Expect(updateDnszoneOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(updateDnszoneOptionsModel.DnszoneID).To(Equal(core.StringPtr("testString")))
				Expect(updateDnszoneOptionsModel.Description).To(Equal(core.StringPtr("The DNS zone is used for VPCs in us-east region")))
				Expect(updateDnszoneOptionsModel.Label).To(Equal(core.StringPtr("us-east")))
				Expect(updateDnszoneOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(updateDnszoneOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateForwardingRuleOptions successfully`, func() {
				// Construct an instance of the UpdateForwardingRuleOptions model
				instanceID := "testString"
				resolverID := "testString"
				ruleID := "testString"
				updateForwardingRuleOptionsModel := dnsSvcsService.NewUpdateForwardingRuleOptions(instanceID, resolverID, ruleID)
				updateForwardingRuleOptionsModel.SetInstanceID("testString")
				updateForwardingRuleOptionsModel.SetResolverID("testString")
				updateForwardingRuleOptionsModel.SetRuleID("testString")
				updateForwardingRuleOptionsModel.SetDescription("forwarding rule")
				updateForwardingRuleOptionsModel.SetMatch("example.com")
				updateForwardingRuleOptionsModel.SetForwardTo([]string{"161.26.0.7"})
				updateForwardingRuleOptionsModel.SetXCorrelationID("testString")
				updateForwardingRuleOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateForwardingRuleOptionsModel).ToNot(BeNil())
				Expect(updateForwardingRuleOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(updateForwardingRuleOptionsModel.ResolverID).To(Equal(core.StringPtr("testString")))
				Expect(updateForwardingRuleOptionsModel.RuleID).To(Equal(core.StringPtr("testString")))
				Expect(updateForwardingRuleOptionsModel.Description).To(Equal(core.StringPtr("forwarding rule")))
				Expect(updateForwardingRuleOptionsModel.Match).To(Equal(core.StringPtr("example.com")))
				Expect(updateForwardingRuleOptionsModel.ForwardTo).To(Equal([]string{"161.26.0.7"}))
				Expect(updateForwardingRuleOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(updateForwardingRuleOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateLoadBalancerOptions successfully`, func() {
				// Construct an instance of the LoadBalancerAzPoolsItem model
				loadBalancerAzPoolsItemModel := new(dnssvcsv1.LoadBalancerAzPoolsItem)
				Expect(loadBalancerAzPoolsItemModel).ToNot(BeNil())
				loadBalancerAzPoolsItemModel.AvailabilityZone = core.StringPtr("us-south-1")
				loadBalancerAzPoolsItemModel.Pools = []string{"0fc0bb7c-2fab-476e-8b9b-40fa14bf8e3d"}
				Expect(loadBalancerAzPoolsItemModel.AvailabilityZone).To(Equal(core.StringPtr("us-south-1")))
				Expect(loadBalancerAzPoolsItemModel.Pools).To(Equal([]string{"0fc0bb7c-2fab-476e-8b9b-40fa14bf8e3d"}))

				// Construct an instance of the UpdateLoadBalancerOptions model
				instanceID := "testString"
				dnszoneID := "testString"
				lbID := "testString"
				updateLoadBalancerOptionsModel := dnsSvcsService.NewUpdateLoadBalancerOptions(instanceID, dnszoneID, lbID)
				updateLoadBalancerOptionsModel.SetInstanceID("testString")
				updateLoadBalancerOptionsModel.SetDnszoneID("testString")
				updateLoadBalancerOptionsModel.SetLbID("testString")
				updateLoadBalancerOptionsModel.SetName("glb.example.com")
				updateLoadBalancerOptionsModel.SetDescription("Load balancer for glb.example.com.")
				updateLoadBalancerOptionsModel.SetEnabled(true)
				updateLoadBalancerOptionsModel.SetTTL(int64(120))
				updateLoadBalancerOptionsModel.SetFallbackPool("24ccf79a-4ae0-4769-b4c8-17f8f230072e")
				updateLoadBalancerOptionsModel.SetDefaultPools([]string{"testString"})
				updateLoadBalancerOptionsModel.SetAzPools([]dnssvcsv1.LoadBalancerAzPoolsItem{*loadBalancerAzPoolsItemModel})
				updateLoadBalancerOptionsModel.SetXCorrelationID("testString")
				updateLoadBalancerOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateLoadBalancerOptionsModel).ToNot(BeNil())
				Expect(updateLoadBalancerOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(updateLoadBalancerOptionsModel.DnszoneID).To(Equal(core.StringPtr("testString")))
				Expect(updateLoadBalancerOptionsModel.LbID).To(Equal(core.StringPtr("testString")))
				Expect(updateLoadBalancerOptionsModel.Name).To(Equal(core.StringPtr("glb.example.com")))
				Expect(updateLoadBalancerOptionsModel.Description).To(Equal(core.StringPtr("Load balancer for glb.example.com.")))
				Expect(updateLoadBalancerOptionsModel.Enabled).To(Equal(core.BoolPtr(true)))
				Expect(updateLoadBalancerOptionsModel.TTL).To(Equal(core.Int64Ptr(int64(120))))
				Expect(updateLoadBalancerOptionsModel.FallbackPool).To(Equal(core.StringPtr("24ccf79a-4ae0-4769-b4c8-17f8f230072e")))
				Expect(updateLoadBalancerOptionsModel.DefaultPools).To(Equal([]string{"testString"}))
				Expect(updateLoadBalancerOptionsModel.AzPools).To(Equal([]dnssvcsv1.LoadBalancerAzPoolsItem{*loadBalancerAzPoolsItemModel}))
				Expect(updateLoadBalancerOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(updateLoadBalancerOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateMonitorOptions successfully`, func() {
				// Construct an instance of the HealthcheckHeader model
				healthcheckHeaderModel := new(dnssvcsv1.HealthcheckHeader)
				Expect(healthcheckHeaderModel).ToNot(BeNil())
				healthcheckHeaderModel.Name = core.StringPtr("Host")
				healthcheckHeaderModel.Value = []string{"origin.example.com"}
				Expect(healthcheckHeaderModel.Name).To(Equal(core.StringPtr("Host")))
				Expect(healthcheckHeaderModel.Value).To(Equal([]string{"origin.example.com"}))

				// Construct an instance of the UpdateMonitorOptions model
				instanceID := "testString"
				monitorID := "testString"
				updateMonitorOptionsModel := dnsSvcsService.NewUpdateMonitorOptions(instanceID, monitorID)
				updateMonitorOptionsModel.SetInstanceID("testString")
				updateMonitorOptionsModel.SetMonitorID("testString")
				updateMonitorOptionsModel.SetName("healthcheck-monitor")
				updateMonitorOptionsModel.SetDescription("Load balancer monitor for glb.example.com.")
				updateMonitorOptionsModel.SetType("HTTPS")
				updateMonitorOptionsModel.SetPort(int64(8080))
				updateMonitorOptionsModel.SetInterval(int64(60))
				updateMonitorOptionsModel.SetRetries(int64(2))
				updateMonitorOptionsModel.SetTimeout(int64(5))
				updateMonitorOptionsModel.SetMethod("GET")
				updateMonitorOptionsModel.SetPath("/health")
				updateMonitorOptionsModel.SetHeadersVar([]dnssvcsv1.HealthcheckHeader{*healthcheckHeaderModel})
				updateMonitorOptionsModel.SetAllowInsecure(false)
				updateMonitorOptionsModel.SetExpectedCodes("200")
				updateMonitorOptionsModel.SetExpectedBody("alive")
				updateMonitorOptionsModel.SetXCorrelationID("testString")
				updateMonitorOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateMonitorOptionsModel).ToNot(BeNil())
				Expect(updateMonitorOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(updateMonitorOptionsModel.MonitorID).To(Equal(core.StringPtr("testString")))
				Expect(updateMonitorOptionsModel.Name).To(Equal(core.StringPtr("healthcheck-monitor")))
				Expect(updateMonitorOptionsModel.Description).To(Equal(core.StringPtr("Load balancer monitor for glb.example.com.")))
				Expect(updateMonitorOptionsModel.Type).To(Equal(core.StringPtr("HTTPS")))
				Expect(updateMonitorOptionsModel.Port).To(Equal(core.Int64Ptr(int64(8080))))
				Expect(updateMonitorOptionsModel.Interval).To(Equal(core.Int64Ptr(int64(60))))
				Expect(updateMonitorOptionsModel.Retries).To(Equal(core.Int64Ptr(int64(2))))
				Expect(updateMonitorOptionsModel.Timeout).To(Equal(core.Int64Ptr(int64(5))))
				Expect(updateMonitorOptionsModel.Method).To(Equal(core.StringPtr("GET")))
				Expect(updateMonitorOptionsModel.Path).To(Equal(core.StringPtr("/health")))
				Expect(updateMonitorOptionsModel.HeadersVar).To(Equal([]dnssvcsv1.HealthcheckHeader{*healthcheckHeaderModel}))
				Expect(updateMonitorOptionsModel.AllowInsecure).To(Equal(core.BoolPtr(false)))
				Expect(updateMonitorOptionsModel.ExpectedCodes).To(Equal(core.StringPtr("200")))
				Expect(updateMonitorOptionsModel.ExpectedBody).To(Equal(core.StringPtr("alive")))
				Expect(updateMonitorOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(updateMonitorOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdatePoolOptions successfully`, func() {
				// Construct an instance of the OriginInput model
				originInputModel := new(dnssvcsv1.OriginInput)
				Expect(originInputModel).ToNot(BeNil())
				originInputModel.Name = core.StringPtr("app-server-1")
				originInputModel.Description = core.StringPtr("description of the origin server")
				originInputModel.Address = core.StringPtr("10.10.16.8")
				originInputModel.Enabled = core.BoolPtr(true)
				Expect(originInputModel.Name).To(Equal(core.StringPtr("app-server-1")))
				Expect(originInputModel.Description).To(Equal(core.StringPtr("description of the origin server")))
				Expect(originInputModel.Address).To(Equal(core.StringPtr("10.10.16.8")))
				Expect(originInputModel.Enabled).To(Equal(core.BoolPtr(true)))

				// Construct an instance of the UpdatePoolOptions model
				instanceID := "testString"
				poolID := "testString"
				updatePoolOptionsModel := dnsSvcsService.NewUpdatePoolOptions(instanceID, poolID)
				updatePoolOptionsModel.SetInstanceID("testString")
				updatePoolOptionsModel.SetPoolID("testString")
				updatePoolOptionsModel.SetName("dal10-az-pool")
				updatePoolOptionsModel.SetDescription("Load balancer pool for dal10 availability zone.")
				updatePoolOptionsModel.SetEnabled(true)
				updatePoolOptionsModel.SetHealthyOriginsThreshold(int64(1))
				updatePoolOptionsModel.SetOrigins([]dnssvcsv1.OriginInput{*originInputModel})
				updatePoolOptionsModel.SetMonitor("7dd6841c-264e-11ea-88df-062967242a6a")
				updatePoolOptionsModel.SetNotificationChannel("https://mywebsite.com/dns/webhook")
				updatePoolOptionsModel.SetHealthcheckRegion("us-south")
				updatePoolOptionsModel.SetHealthcheckSubnets([]string{"crn:v1:staging:public:is:us-south-1:a/01652b251c3ae2787110a995d8db0135::subnet:0716-b49ef064-0f89-4fb1-8212-135b12568f04"})
				updatePoolOptionsModel.SetXCorrelationID("testString")
				updatePoolOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updatePoolOptionsModel).ToNot(BeNil())
				Expect(updatePoolOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(updatePoolOptionsModel.PoolID).To(Equal(core.StringPtr("testString")))
				Expect(updatePoolOptionsModel.Name).To(Equal(core.StringPtr("dal10-az-pool")))
				Expect(updatePoolOptionsModel.Description).To(Equal(core.StringPtr("Load balancer pool for dal10 availability zone.")))
				Expect(updatePoolOptionsModel.Enabled).To(Equal(core.BoolPtr(true)))
				Expect(updatePoolOptionsModel.HealthyOriginsThreshold).To(Equal(core.Int64Ptr(int64(1))))
				Expect(updatePoolOptionsModel.Origins).To(Equal([]dnssvcsv1.OriginInput{*originInputModel}))
				Expect(updatePoolOptionsModel.Monitor).To(Equal(core.StringPtr("7dd6841c-264e-11ea-88df-062967242a6a")))
				Expect(updatePoolOptionsModel.NotificationChannel).To(Equal(core.StringPtr("https://mywebsite.com/dns/webhook")))
				Expect(updatePoolOptionsModel.HealthcheckRegion).To(Equal(core.StringPtr("us-south")))
				Expect(updatePoolOptionsModel.HealthcheckSubnets).To(Equal([]string{"crn:v1:staging:public:is:us-south-1:a/01652b251c3ae2787110a995d8db0135::subnet:0716-b49ef064-0f89-4fb1-8212-135b12568f04"}))
				Expect(updatePoolOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(updatePoolOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateResourceRecordOptions successfully`, func() {
				// Construct an instance of the ResourceRecordUpdateInputRdataRdataARecord model
				resourceRecordUpdateInputRdataModel := new(dnssvcsv1.ResourceRecordUpdateInputRdataRdataARecord)
				Expect(resourceRecordUpdateInputRdataModel).ToNot(BeNil())
				resourceRecordUpdateInputRdataModel.Ip = core.StringPtr("10.110.201.214")
				Expect(resourceRecordUpdateInputRdataModel.Ip).To(Equal(core.StringPtr("10.110.201.214")))

				// Construct an instance of the UpdateResourceRecordOptions model
				instanceID := "testString"
				dnszoneID := "testString"
				recordID := "testString"
				updateResourceRecordOptionsModel := dnsSvcsService.NewUpdateResourceRecordOptions(instanceID, dnszoneID, recordID)
				updateResourceRecordOptionsModel.SetInstanceID("testString")
				updateResourceRecordOptionsModel.SetDnszoneID("testString")
				updateResourceRecordOptionsModel.SetRecordID("testString")
				updateResourceRecordOptionsModel.SetName("test.example.com")
				updateResourceRecordOptionsModel.SetRdata(resourceRecordUpdateInputRdataModel)
				updateResourceRecordOptionsModel.SetTTL(int64(120))
				updateResourceRecordOptionsModel.SetService("_sip")
				updateResourceRecordOptionsModel.SetProtocol("udp")
				updateResourceRecordOptionsModel.SetXCorrelationID("testString")
				updateResourceRecordOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateResourceRecordOptionsModel).ToNot(BeNil())
				Expect(updateResourceRecordOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(updateResourceRecordOptionsModel.DnszoneID).To(Equal(core.StringPtr("testString")))
				Expect(updateResourceRecordOptionsModel.RecordID).To(Equal(core.StringPtr("testString")))
				Expect(updateResourceRecordOptionsModel.Name).To(Equal(core.StringPtr("test.example.com")))
				Expect(updateResourceRecordOptionsModel.Rdata).To(Equal(resourceRecordUpdateInputRdataModel))
				Expect(updateResourceRecordOptionsModel.TTL).To(Equal(core.Int64Ptr(int64(120))))
				Expect(updateResourceRecordOptionsModel.Service).To(Equal(core.StringPtr("_sip")))
				Expect(updateResourceRecordOptionsModel.Protocol).To(Equal(core.StringPtr("udp")))
				Expect(updateResourceRecordOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(updateResourceRecordOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewHealthcheckHeader successfully`, func() {
				name := "Host"
				value := []string{"origin.example.com"}
				model, err := dnsSvcsService.NewHealthcheckHeader(name, value)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewLocationInput successfully`, func() {
				subnetCrn := "crn:v1:bluemix:public:is:us-south-1:a/01652b251c3ae2787110a995d8db0135::subnet:0716-b49ef064-0f89-4fb1-8212-135b12568f04"
				model, err := dnsSvcsService.NewLocationInput(subnetCrn)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewPermittedNetworkVpc successfully`, func() {
				vpcCrn := "crn:v1:bluemix:public:is:eu-de:a/bcf1865e99742d38d2d5fc3fb80a5496::vpc:6e6cc326-04d1-4c99-a289-efb3ae4193d6"
				model, err := dnsSvcsService.NewPermittedNetworkVpc(vpcCrn)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewResourceRecordInputRdataRdataARecord successfully`, func() {
				ip := "10.110.201.214"
				model, err := dnsSvcsService.NewResourceRecordInputRdataRdataARecord(ip)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewResourceRecordInputRdataRdataAaaaRecord successfully`, func() {
				ip := "2019::2019"
				model, err := dnsSvcsService.NewResourceRecordInputRdataRdataAaaaRecord(ip)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewResourceRecordInputRdataRdataCnameRecord successfully`, func() {
				cname := "www.example.com"
				model, err := dnsSvcsService.NewResourceRecordInputRdataRdataCnameRecord(cname)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewResourceRecordInputRdataRdataMxRecord successfully`, func() {
				exchange := "mail.example.com"
				preference := int64(10)
				model, err := dnsSvcsService.NewResourceRecordInputRdataRdataMxRecord(exchange, preference)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewResourceRecordInputRdataRdataPtrRecord successfully`, func() {
				ptrdname := "www.example.com"
				model, err := dnsSvcsService.NewResourceRecordInputRdataRdataPtrRecord(ptrdname)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewResourceRecordInputRdataRdataSrvRecord successfully`, func() {
				port := int64(80)
				priority := int64(10)
				target := "www.example.com"
				weight := int64(10)
				model, err := dnsSvcsService.NewResourceRecordInputRdataRdataSrvRecord(port, priority, target, weight)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewResourceRecordInputRdataRdataTxtRecord successfully`, func() {
				text := "This is a text record"
				model, err := dnsSvcsService.NewResourceRecordInputRdataRdataTxtRecord(text)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewResourceRecordUpdateInputRdataRdataARecord successfully`, func() {
				ip := "10.110.201.214"
				model, err := dnsSvcsService.NewResourceRecordUpdateInputRdataRdataARecord(ip)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewResourceRecordUpdateInputRdataRdataAaaaRecord successfully`, func() {
				ip := "2019::2019"
				model, err := dnsSvcsService.NewResourceRecordUpdateInputRdataRdataAaaaRecord(ip)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewResourceRecordUpdateInputRdataRdataCnameRecord successfully`, func() {
				cname := "www.example.com"
				model, err := dnsSvcsService.NewResourceRecordUpdateInputRdataRdataCnameRecord(cname)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewResourceRecordUpdateInputRdataRdataMxRecord successfully`, func() {
				exchange := "mail.example.com"
				preference := int64(10)
				model, err := dnsSvcsService.NewResourceRecordUpdateInputRdataRdataMxRecord(exchange, preference)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewResourceRecordUpdateInputRdataRdataPtrRecord successfully`, func() {
				ptrdname := "www.example.com"
				model, err := dnsSvcsService.NewResourceRecordUpdateInputRdataRdataPtrRecord(ptrdname)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewResourceRecordUpdateInputRdataRdataSrvRecord successfully`, func() {
				port := int64(80)
				priority := int64(10)
				target := "www.example.com"
				weight := int64(10)
				model, err := dnsSvcsService.NewResourceRecordUpdateInputRdataRdataSrvRecord(port, priority, target, weight)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewResourceRecordUpdateInputRdataRdataTxtRecord successfully`, func() {
				text := "This is a text record"
				model, err := dnsSvcsService.NewResourceRecordUpdateInputRdataRdataTxtRecord(text)
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
