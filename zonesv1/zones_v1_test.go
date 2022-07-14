/**
 * (C) Copyright IBM Corp. 2022.
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

package zonesv1_test

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
	"github.com/IBM/networking-go-sdk/zonesv1"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe(`ZonesV1`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		crn := "testString"
		It(`Instantiate service client`, func() {
			zonesService, serviceErr := zonesv1.NewZonesV1(&zonesv1.ZonesV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
				Crn:           core.StringPtr(crn),
			})
			Expect(zonesService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			zonesService, serviceErr := zonesv1.NewZonesV1(&zonesv1.ZonesV1Options{
				URL: "{BAD_URL_STRING",
				Crn: core.StringPtr(crn),
			})
			Expect(zonesService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			zonesService, serviceErr := zonesv1.NewZonesV1(&zonesv1.ZonesV1Options{
				URL: "https://zonesv1/api",
				Crn: core.StringPtr(crn),
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(zonesService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Validation Error`, func() {
			zonesService, serviceErr := zonesv1.NewZonesV1(&zonesv1.ZonesV1Options{})
			Expect(zonesService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		crn := "testString"
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"ZONES_URL":       "https://zonesv1/api",
				"ZONES_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				zonesService, serviceErr := zonesv1.NewZonesV1UsingExternalConfig(&zonesv1.ZonesV1Options{
					Crn: core.StringPtr(crn),
				})
				Expect(zonesService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := zonesService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != zonesService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(zonesService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(zonesService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				zonesService, serviceErr := zonesv1.NewZonesV1UsingExternalConfig(&zonesv1.ZonesV1Options{
					URL: "https://testService/api",
					Crn: core.StringPtr(crn),
				})
				Expect(zonesService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(zonesService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := zonesService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != zonesService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(zonesService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(zonesService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				zonesService, serviceErr := zonesv1.NewZonesV1UsingExternalConfig(&zonesv1.ZonesV1Options{
					Crn: core.StringPtr(crn),
				})
				err := zonesService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(zonesService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(zonesService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := zonesService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != zonesService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(zonesService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(zonesService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"ZONES_URL":       "https://zonesv1/api",
				"ZONES_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			zonesService, serviceErr := zonesv1.NewZonesV1UsingExternalConfig(&zonesv1.ZonesV1Options{
				Crn: core.StringPtr(crn),
			})

			It(`Instantiate service client with error`, func() {
				Expect(zonesService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"ZONES_AUTH_TYPE": "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			zonesService, serviceErr := zonesv1.NewZonesV1UsingExternalConfig(&zonesv1.ZonesV1Options{
				URL: "{BAD_URL_STRING",
				Crn: core.StringPtr(crn),
			})

			It(`Instantiate service client with error`, func() {
				Expect(zonesService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = zonesv1.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`ListZones(listZonesOptions *ListZonesOptions) - Operation response error`, func() {
		crn := "testString"
		listZonesPath := "/v1/testString/zones"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listZonesPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["page"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					Expect(req.URL.Query()["per_page"]).To(Equal([]string{fmt.Sprint(int64(5))}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListZones with error: Operation response processing error`, func() {
				zonesService, serviceErr := zonesv1.NewZonesV1(&zonesv1.ZonesV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesService).ToNot(BeNil())

				// Construct an instance of the ListZonesOptions model
				listZonesOptionsModel := new(zonesv1.ListZonesOptions)
				listZonesOptionsModel.Page = core.Int64Ptr(int64(38))
				listZonesOptionsModel.PerPage = core.Int64Ptr(int64(5))
				listZonesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := zonesService.ListZones(listZonesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				zonesService.EnableRetries(0, 0)
				result, response, operationErr = zonesService.ListZones(listZonesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListZones(listZonesOptions *ListZonesOptions)`, func() {
		crn := "testString"
		listZonesPath := "/v1/testString/zones"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listZonesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["page"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					Expect(req.URL.Query()["per_page"]).To(Equal([]string{fmt.Sprint(int64(5))}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": [{"id": "f1aba936b94213e5b8dca0c0dbf1f9cc", "created_on": "2014-01-01T05:20:00.12345Z", "modified_on": "2014-01-01T05:20:00.12345Z", "name": "test-example.com", "original_registrar": "GoDaddy", "original_dnshost": "NameCheap", "status": "active", "paused": false, "original_name_servers": ["ns1.originaldnshost.com"], "name_servers": ["ns001.name.cloud.ibm.com"], "type": "full", "verification_key": "476754457-428595283", "cname_suffix": "cdn.cloudflare.net"}], "result_info": {"page": 1, "per_page": 20, "count": 1, "total_count": 2000}}`)
				}))
			})
			It(`Invoke ListZones successfully with retries`, func() {
				zonesService, serviceErr := zonesv1.NewZonesV1(&zonesv1.ZonesV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesService).ToNot(BeNil())
				zonesService.EnableRetries(0, 0)

				// Construct an instance of the ListZonesOptions model
				listZonesOptionsModel := new(zonesv1.ListZonesOptions)
				listZonesOptionsModel.Page = core.Int64Ptr(int64(38))
				listZonesOptionsModel.PerPage = core.Int64Ptr(int64(5))
				listZonesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := zonesService.ListZonesWithContext(ctx, listZonesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				zonesService.DisableRetries()
				result, response, operationErr := zonesService.ListZones(listZonesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = zonesService.ListZonesWithContext(ctx, listZonesOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listZonesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["page"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					Expect(req.URL.Query()["per_page"]).To(Equal([]string{fmt.Sprint(int64(5))}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": [{"id": "f1aba936b94213e5b8dca0c0dbf1f9cc", "created_on": "2014-01-01T05:20:00.12345Z", "modified_on": "2014-01-01T05:20:00.12345Z", "name": "test-example.com", "original_registrar": "GoDaddy", "original_dnshost": "NameCheap", "status": "active", "paused": false, "original_name_servers": ["ns1.originaldnshost.com"], "name_servers": ["ns001.name.cloud.ibm.com"], "type": "full", "verification_key": "476754457-428595283", "cname_suffix": "cdn.cloudflare.net"}], "result_info": {"page": 1, "per_page": 20, "count": 1, "total_count": 2000}}`)
				}))
			})
			It(`Invoke ListZones successfully`, func() {
				zonesService, serviceErr := zonesv1.NewZonesV1(&zonesv1.ZonesV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := zonesService.ListZones(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListZonesOptions model
				listZonesOptionsModel := new(zonesv1.ListZonesOptions)
				listZonesOptionsModel.Page = core.Int64Ptr(int64(38))
				listZonesOptionsModel.PerPage = core.Int64Ptr(int64(5))
				listZonesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = zonesService.ListZones(listZonesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListZones with error: Operation request error`, func() {
				zonesService, serviceErr := zonesv1.NewZonesV1(&zonesv1.ZonesV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesService).ToNot(BeNil())

				// Construct an instance of the ListZonesOptions model
				listZonesOptionsModel := new(zonesv1.ListZonesOptions)
				listZonesOptionsModel.Page = core.Int64Ptr(int64(38))
				listZonesOptionsModel.PerPage = core.Int64Ptr(int64(5))
				listZonesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := zonesService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := zonesService.ListZones(listZonesOptionsModel)
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
			It(`Invoke ListZones successfully`, func() {
				zonesService, serviceErr := zonesv1.NewZonesV1(&zonesv1.ZonesV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesService).ToNot(BeNil())

				// Construct an instance of the ListZonesOptions model
				listZonesOptionsModel := new(zonesv1.ListZonesOptions)
				listZonesOptionsModel.Page = core.Int64Ptr(int64(38))
				listZonesOptionsModel.PerPage = core.Int64Ptr(int64(5))
				listZonesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := zonesService.ListZones(listZonesOptionsModel)
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
	Describe(`CreateZone(createZoneOptions *CreateZoneOptions) - Operation response error`, func() {
		crn := "testString"
		createZonePath := "/v1/testString/zones"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createZonePath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateZone with error: Operation response processing error`, func() {
				zonesService, serviceErr := zonesv1.NewZonesV1(&zonesv1.ZonesV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesService).ToNot(BeNil())

				// Construct an instance of the CreateZoneOptions model
				createZoneOptionsModel := new(zonesv1.CreateZoneOptions)
				createZoneOptionsModel.Name = core.StringPtr("test-example.com")
				createZoneOptionsModel.Type = core.StringPtr("full")
				createZoneOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := zonesService.CreateZone(createZoneOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				zonesService.EnableRetries(0, 0)
				result, response, operationErr = zonesService.CreateZone(createZoneOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateZone(createZoneOptions *CreateZoneOptions)`, func() {
		crn := "testString"
		createZonePath := "/v1/testString/zones"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createZonePath))
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

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": "f1aba936b94213e5b8dca0c0dbf1f9cc", "created_on": "2014-01-01T05:20:00.12345Z", "modified_on": "2014-01-01T05:20:00.12345Z", "name": "test-example.com", "original_registrar": "GoDaddy", "original_dnshost": "NameCheap", "status": "active", "paused": false, "original_name_servers": ["ns1.originaldnshost.com"], "name_servers": ["ns001.name.cloud.ibm.com"], "type": "full", "verification_key": "476754457-428595283", "cname_suffix": "cdn.cloudflare.net"}}`)
				}))
			})
			It(`Invoke CreateZone successfully with retries`, func() {
				zonesService, serviceErr := zonesv1.NewZonesV1(&zonesv1.ZonesV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesService).ToNot(BeNil())
				zonesService.EnableRetries(0, 0)

				// Construct an instance of the CreateZoneOptions model
				createZoneOptionsModel := new(zonesv1.CreateZoneOptions)
				createZoneOptionsModel.Name = core.StringPtr("test-example.com")
				createZoneOptionsModel.Type = core.StringPtr("full")
				createZoneOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := zonesService.CreateZoneWithContext(ctx, createZoneOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				zonesService.DisableRetries()
				result, response, operationErr := zonesService.CreateZone(createZoneOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = zonesService.CreateZoneWithContext(ctx, createZoneOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createZonePath))
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

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": "f1aba936b94213e5b8dca0c0dbf1f9cc", "created_on": "2014-01-01T05:20:00.12345Z", "modified_on": "2014-01-01T05:20:00.12345Z", "name": "test-example.com", "original_registrar": "GoDaddy", "original_dnshost": "NameCheap", "status": "active", "paused": false, "original_name_servers": ["ns1.originaldnshost.com"], "name_servers": ["ns001.name.cloud.ibm.com"], "type": "full", "verification_key": "476754457-428595283", "cname_suffix": "cdn.cloudflare.net"}}`)
				}))
			})
			It(`Invoke CreateZone successfully`, func() {
				zonesService, serviceErr := zonesv1.NewZonesV1(&zonesv1.ZonesV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := zonesService.CreateZone(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateZoneOptions model
				createZoneOptionsModel := new(zonesv1.CreateZoneOptions)
				createZoneOptionsModel.Name = core.StringPtr("test-example.com")
				createZoneOptionsModel.Type = core.StringPtr("full")
				createZoneOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = zonesService.CreateZone(createZoneOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateZone with error: Operation request error`, func() {
				zonesService, serviceErr := zonesv1.NewZonesV1(&zonesv1.ZonesV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesService).ToNot(BeNil())

				// Construct an instance of the CreateZoneOptions model
				createZoneOptionsModel := new(zonesv1.CreateZoneOptions)
				createZoneOptionsModel.Name = core.StringPtr("test-example.com")
				createZoneOptionsModel.Type = core.StringPtr("full")
				createZoneOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := zonesService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := zonesService.CreateZone(createZoneOptionsModel)
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
			It(`Invoke CreateZone successfully`, func() {
				zonesService, serviceErr := zonesv1.NewZonesV1(&zonesv1.ZonesV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesService).ToNot(BeNil())

				// Construct an instance of the CreateZoneOptions model
				createZoneOptionsModel := new(zonesv1.CreateZoneOptions)
				createZoneOptionsModel.Name = core.StringPtr("test-example.com")
				createZoneOptionsModel.Type = core.StringPtr("full")
				createZoneOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := zonesService.CreateZone(createZoneOptionsModel)
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
	Describe(`DeleteZone(deleteZoneOptions *DeleteZoneOptions) - Operation response error`, func() {
		crn := "testString"
		deleteZonePath := "/v1/testString/zones/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteZonePath))
					Expect(req.Method).To(Equal("DELETE"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke DeleteZone with error: Operation response processing error`, func() {
				zonesService, serviceErr := zonesv1.NewZonesV1(&zonesv1.ZonesV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesService).ToNot(BeNil())

				// Construct an instance of the DeleteZoneOptions model
				deleteZoneOptionsModel := new(zonesv1.DeleteZoneOptions)
				deleteZoneOptionsModel.ZoneIdentifier = core.StringPtr("testString")
				deleteZoneOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := zonesService.DeleteZone(deleteZoneOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				zonesService.EnableRetries(0, 0)
				result, response, operationErr = zonesService.DeleteZone(deleteZoneOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteZone(deleteZoneOptions *DeleteZoneOptions)`, func() {
		crn := "testString"
		deleteZonePath := "/v1/testString/zones/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteZonePath))
					Expect(req.Method).To(Equal("DELETE"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": "f1aba936b94213e5b8dca0c0dbf1f9cc"}}`)
				}))
			})
			It(`Invoke DeleteZone successfully with retries`, func() {
				zonesService, serviceErr := zonesv1.NewZonesV1(&zonesv1.ZonesV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesService).ToNot(BeNil())
				zonesService.EnableRetries(0, 0)

				// Construct an instance of the DeleteZoneOptions model
				deleteZoneOptionsModel := new(zonesv1.DeleteZoneOptions)
				deleteZoneOptionsModel.ZoneIdentifier = core.StringPtr("testString")
				deleteZoneOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := zonesService.DeleteZoneWithContext(ctx, deleteZoneOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				zonesService.DisableRetries()
				result, response, operationErr := zonesService.DeleteZone(deleteZoneOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = zonesService.DeleteZoneWithContext(ctx, deleteZoneOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(deleteZonePath))
					Expect(req.Method).To(Equal("DELETE"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": "f1aba936b94213e5b8dca0c0dbf1f9cc"}}`)
				}))
			})
			It(`Invoke DeleteZone successfully`, func() {
				zonesService, serviceErr := zonesv1.NewZonesV1(&zonesv1.ZonesV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := zonesService.DeleteZone(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DeleteZoneOptions model
				deleteZoneOptionsModel := new(zonesv1.DeleteZoneOptions)
				deleteZoneOptionsModel.ZoneIdentifier = core.StringPtr("testString")
				deleteZoneOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = zonesService.DeleteZone(deleteZoneOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke DeleteZone with error: Operation validation and request error`, func() {
				zonesService, serviceErr := zonesv1.NewZonesV1(&zonesv1.ZonesV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesService).ToNot(BeNil())

				// Construct an instance of the DeleteZoneOptions model
				deleteZoneOptionsModel := new(zonesv1.DeleteZoneOptions)
				deleteZoneOptionsModel.ZoneIdentifier = core.StringPtr("testString")
				deleteZoneOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := zonesService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := zonesService.DeleteZone(deleteZoneOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the DeleteZoneOptions model with no property values
				deleteZoneOptionsModelNew := new(zonesv1.DeleteZoneOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = zonesService.DeleteZone(deleteZoneOptionsModelNew)
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
			It(`Invoke DeleteZone successfully`, func() {
				zonesService, serviceErr := zonesv1.NewZonesV1(&zonesv1.ZonesV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesService).ToNot(BeNil())

				// Construct an instance of the DeleteZoneOptions model
				deleteZoneOptionsModel := new(zonesv1.DeleteZoneOptions)
				deleteZoneOptionsModel.ZoneIdentifier = core.StringPtr("testString")
				deleteZoneOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := zonesService.DeleteZone(deleteZoneOptionsModel)
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
	Describe(`GetZone(getZoneOptions *GetZoneOptions) - Operation response error`, func() {
		crn := "testString"
		getZonePath := "/v1/testString/zones/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getZonePath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetZone with error: Operation response processing error`, func() {
				zonesService, serviceErr := zonesv1.NewZonesV1(&zonesv1.ZonesV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesService).ToNot(BeNil())

				// Construct an instance of the GetZoneOptions model
				getZoneOptionsModel := new(zonesv1.GetZoneOptions)
				getZoneOptionsModel.ZoneIdentifier = core.StringPtr("testString")
				getZoneOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := zonesService.GetZone(getZoneOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				zonesService.EnableRetries(0, 0)
				result, response, operationErr = zonesService.GetZone(getZoneOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetZone(getZoneOptions *GetZoneOptions)`, func() {
		crn := "testString"
		getZonePath := "/v1/testString/zones/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getZonePath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": "f1aba936b94213e5b8dca0c0dbf1f9cc", "created_on": "2014-01-01T05:20:00.12345Z", "modified_on": "2014-01-01T05:20:00.12345Z", "name": "test-example.com", "original_registrar": "GoDaddy", "original_dnshost": "NameCheap", "status": "active", "paused": false, "original_name_servers": ["ns1.originaldnshost.com"], "name_servers": ["ns001.name.cloud.ibm.com"], "type": "full", "verification_key": "476754457-428595283", "cname_suffix": "cdn.cloudflare.net"}}`)
				}))
			})
			It(`Invoke GetZone successfully with retries`, func() {
				zonesService, serviceErr := zonesv1.NewZonesV1(&zonesv1.ZonesV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesService).ToNot(BeNil())
				zonesService.EnableRetries(0, 0)

				// Construct an instance of the GetZoneOptions model
				getZoneOptionsModel := new(zonesv1.GetZoneOptions)
				getZoneOptionsModel.ZoneIdentifier = core.StringPtr("testString")
				getZoneOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := zonesService.GetZoneWithContext(ctx, getZoneOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				zonesService.DisableRetries()
				result, response, operationErr := zonesService.GetZone(getZoneOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = zonesService.GetZoneWithContext(ctx, getZoneOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getZonePath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": "f1aba936b94213e5b8dca0c0dbf1f9cc", "created_on": "2014-01-01T05:20:00.12345Z", "modified_on": "2014-01-01T05:20:00.12345Z", "name": "test-example.com", "original_registrar": "GoDaddy", "original_dnshost": "NameCheap", "status": "active", "paused": false, "original_name_servers": ["ns1.originaldnshost.com"], "name_servers": ["ns001.name.cloud.ibm.com"], "type": "full", "verification_key": "476754457-428595283", "cname_suffix": "cdn.cloudflare.net"}}`)
				}))
			})
			It(`Invoke GetZone successfully`, func() {
				zonesService, serviceErr := zonesv1.NewZonesV1(&zonesv1.ZonesV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := zonesService.GetZone(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetZoneOptions model
				getZoneOptionsModel := new(zonesv1.GetZoneOptions)
				getZoneOptionsModel.ZoneIdentifier = core.StringPtr("testString")
				getZoneOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = zonesService.GetZone(getZoneOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetZone with error: Operation validation and request error`, func() {
				zonesService, serviceErr := zonesv1.NewZonesV1(&zonesv1.ZonesV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesService).ToNot(BeNil())

				// Construct an instance of the GetZoneOptions model
				getZoneOptionsModel := new(zonesv1.GetZoneOptions)
				getZoneOptionsModel.ZoneIdentifier = core.StringPtr("testString")
				getZoneOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := zonesService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := zonesService.GetZone(getZoneOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetZoneOptions model with no property values
				getZoneOptionsModelNew := new(zonesv1.GetZoneOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = zonesService.GetZone(getZoneOptionsModelNew)
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
			It(`Invoke GetZone successfully`, func() {
				zonesService, serviceErr := zonesv1.NewZonesV1(&zonesv1.ZonesV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesService).ToNot(BeNil())

				// Construct an instance of the GetZoneOptions model
				getZoneOptionsModel := new(zonesv1.GetZoneOptions)
				getZoneOptionsModel.ZoneIdentifier = core.StringPtr("testString")
				getZoneOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := zonesService.GetZone(getZoneOptionsModel)
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
	Describe(`UpdateZone(updateZoneOptions *UpdateZoneOptions) - Operation response error`, func() {
		crn := "testString"
		updateZonePath := "/v1/testString/zones/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateZonePath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateZone with error: Operation response processing error`, func() {
				zonesService, serviceErr := zonesv1.NewZonesV1(&zonesv1.ZonesV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesService).ToNot(BeNil())

				// Construct an instance of the UpdateZoneOptions model
				updateZoneOptionsModel := new(zonesv1.UpdateZoneOptions)
				updateZoneOptionsModel.ZoneIdentifier = core.StringPtr("testString")
				updateZoneOptionsModel.Paused = core.BoolPtr(false)
				updateZoneOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := zonesService.UpdateZone(updateZoneOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				zonesService.EnableRetries(0, 0)
				result, response, operationErr = zonesService.UpdateZone(updateZoneOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateZone(updateZoneOptions *UpdateZoneOptions)`, func() {
		crn := "testString"
		updateZonePath := "/v1/testString/zones/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateZonePath))
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

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": "f1aba936b94213e5b8dca0c0dbf1f9cc", "created_on": "2014-01-01T05:20:00.12345Z", "modified_on": "2014-01-01T05:20:00.12345Z", "name": "test-example.com", "original_registrar": "GoDaddy", "original_dnshost": "NameCheap", "status": "active", "paused": false, "original_name_servers": ["ns1.originaldnshost.com"], "name_servers": ["ns001.name.cloud.ibm.com"], "type": "full", "verification_key": "476754457-428595283", "cname_suffix": "cdn.cloudflare.net"}}`)
				}))
			})
			It(`Invoke UpdateZone successfully with retries`, func() {
				zonesService, serviceErr := zonesv1.NewZonesV1(&zonesv1.ZonesV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesService).ToNot(BeNil())
				zonesService.EnableRetries(0, 0)

				// Construct an instance of the UpdateZoneOptions model
				updateZoneOptionsModel := new(zonesv1.UpdateZoneOptions)
				updateZoneOptionsModel.ZoneIdentifier = core.StringPtr("testString")
				updateZoneOptionsModel.Paused = core.BoolPtr(false)
				updateZoneOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := zonesService.UpdateZoneWithContext(ctx, updateZoneOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				zonesService.DisableRetries()
				result, response, operationErr := zonesService.UpdateZone(updateZoneOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = zonesService.UpdateZoneWithContext(ctx, updateZoneOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(updateZonePath))
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

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": "f1aba936b94213e5b8dca0c0dbf1f9cc", "created_on": "2014-01-01T05:20:00.12345Z", "modified_on": "2014-01-01T05:20:00.12345Z", "name": "test-example.com", "original_registrar": "GoDaddy", "original_dnshost": "NameCheap", "status": "active", "paused": false, "original_name_servers": ["ns1.originaldnshost.com"], "name_servers": ["ns001.name.cloud.ibm.com"], "type": "full", "verification_key": "476754457-428595283", "cname_suffix": "cdn.cloudflare.net"}}`)
				}))
			})
			It(`Invoke UpdateZone successfully`, func() {
				zonesService, serviceErr := zonesv1.NewZonesV1(&zonesv1.ZonesV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := zonesService.UpdateZone(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateZoneOptions model
				updateZoneOptionsModel := new(zonesv1.UpdateZoneOptions)
				updateZoneOptionsModel.ZoneIdentifier = core.StringPtr("testString")
				updateZoneOptionsModel.Paused = core.BoolPtr(false)
				updateZoneOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = zonesService.UpdateZone(updateZoneOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateZone with error: Operation validation and request error`, func() {
				zonesService, serviceErr := zonesv1.NewZonesV1(&zonesv1.ZonesV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesService).ToNot(BeNil())

				// Construct an instance of the UpdateZoneOptions model
				updateZoneOptionsModel := new(zonesv1.UpdateZoneOptions)
				updateZoneOptionsModel.ZoneIdentifier = core.StringPtr("testString")
				updateZoneOptionsModel.Paused = core.BoolPtr(false)
				updateZoneOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := zonesService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := zonesService.UpdateZone(updateZoneOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateZoneOptions model with no property values
				updateZoneOptionsModelNew := new(zonesv1.UpdateZoneOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = zonesService.UpdateZone(updateZoneOptionsModelNew)
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
			It(`Invoke UpdateZone successfully`, func() {
				zonesService, serviceErr := zonesv1.NewZonesV1(&zonesv1.ZonesV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesService).ToNot(BeNil())

				// Construct an instance of the UpdateZoneOptions model
				updateZoneOptionsModel := new(zonesv1.UpdateZoneOptions)
				updateZoneOptionsModel.ZoneIdentifier = core.StringPtr("testString")
				updateZoneOptionsModel.Paused = core.BoolPtr(false)
				updateZoneOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := zonesService.UpdateZone(updateZoneOptionsModel)
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
	Describe(`ZoneActivationCheck(zoneActivationCheckOptions *ZoneActivationCheckOptions) - Operation response error`, func() {
		crn := "testString"
		zoneActivationCheckPath := "/v1/testString/zones/testString/activation_check"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(zoneActivationCheckPath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ZoneActivationCheck with error: Operation response processing error`, func() {
				zonesService, serviceErr := zonesv1.NewZonesV1(&zonesv1.ZonesV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesService).ToNot(BeNil())

				// Construct an instance of the ZoneActivationCheckOptions model
				zoneActivationCheckOptionsModel := new(zonesv1.ZoneActivationCheckOptions)
				zoneActivationCheckOptionsModel.ZoneIdentifier = core.StringPtr("testString")
				zoneActivationCheckOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := zonesService.ZoneActivationCheck(zoneActivationCheckOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				zonesService.EnableRetries(0, 0)
				result, response, operationErr = zonesService.ZoneActivationCheck(zoneActivationCheckOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ZoneActivationCheck(zoneActivationCheckOptions *ZoneActivationCheckOptions)`, func() {
		crn := "testString"
		zoneActivationCheckPath := "/v1/testString/zones/testString/activation_check"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(zoneActivationCheckPath))
					Expect(req.Method).To(Equal("PUT"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": "f1aba936b94213e5b8dca0c0dbf1f9cc"}}`)
				}))
			})
			It(`Invoke ZoneActivationCheck successfully with retries`, func() {
				zonesService, serviceErr := zonesv1.NewZonesV1(&zonesv1.ZonesV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesService).ToNot(BeNil())
				zonesService.EnableRetries(0, 0)

				// Construct an instance of the ZoneActivationCheckOptions model
				zoneActivationCheckOptionsModel := new(zonesv1.ZoneActivationCheckOptions)
				zoneActivationCheckOptionsModel.ZoneIdentifier = core.StringPtr("testString")
				zoneActivationCheckOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := zonesService.ZoneActivationCheckWithContext(ctx, zoneActivationCheckOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				zonesService.DisableRetries()
				result, response, operationErr := zonesService.ZoneActivationCheck(zoneActivationCheckOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = zonesService.ZoneActivationCheckWithContext(ctx, zoneActivationCheckOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(zoneActivationCheckPath))
					Expect(req.Method).To(Equal("PUT"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": "f1aba936b94213e5b8dca0c0dbf1f9cc"}}`)
				}))
			})
			It(`Invoke ZoneActivationCheck successfully`, func() {
				zonesService, serviceErr := zonesv1.NewZonesV1(&zonesv1.ZonesV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := zonesService.ZoneActivationCheck(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ZoneActivationCheckOptions model
				zoneActivationCheckOptionsModel := new(zonesv1.ZoneActivationCheckOptions)
				zoneActivationCheckOptionsModel.ZoneIdentifier = core.StringPtr("testString")
				zoneActivationCheckOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = zonesService.ZoneActivationCheck(zoneActivationCheckOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ZoneActivationCheck with error: Operation validation and request error`, func() {
				zonesService, serviceErr := zonesv1.NewZonesV1(&zonesv1.ZonesV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesService).ToNot(BeNil())

				// Construct an instance of the ZoneActivationCheckOptions model
				zoneActivationCheckOptionsModel := new(zonesv1.ZoneActivationCheckOptions)
				zoneActivationCheckOptionsModel.ZoneIdentifier = core.StringPtr("testString")
				zoneActivationCheckOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := zonesService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := zonesService.ZoneActivationCheck(zoneActivationCheckOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ZoneActivationCheckOptions model with no property values
				zoneActivationCheckOptionsModelNew := new(zonesv1.ZoneActivationCheckOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = zonesService.ZoneActivationCheck(zoneActivationCheckOptionsModelNew)
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
			It(`Invoke ZoneActivationCheck successfully`, func() {
				zonesService, serviceErr := zonesv1.NewZonesV1(&zonesv1.ZonesV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesService).ToNot(BeNil())

				// Construct an instance of the ZoneActivationCheckOptions model
				zoneActivationCheckOptionsModel := new(zonesv1.ZoneActivationCheckOptions)
				zoneActivationCheckOptionsModel.ZoneIdentifier = core.StringPtr("testString")
				zoneActivationCheckOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := zonesService.ZoneActivationCheck(zoneActivationCheckOptionsModel)
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
			crn := "testString"
			zonesService, _ := zonesv1.NewZonesV1(&zonesv1.ZonesV1Options{
				URL:           "http://zonesv1modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
				Crn:           core.StringPtr(crn),
			})
			It(`Invoke NewCreateZoneOptions successfully`, func() {
				// Construct an instance of the CreateZoneOptions model
				createZoneOptionsModel := zonesService.NewCreateZoneOptions()
				createZoneOptionsModel.SetName("test-example.com")
				createZoneOptionsModel.SetType("full")
				createZoneOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createZoneOptionsModel).ToNot(BeNil())
				Expect(createZoneOptionsModel.Name).To(Equal(core.StringPtr("test-example.com")))
				Expect(createZoneOptionsModel.Type).To(Equal(core.StringPtr("full")))
				Expect(createZoneOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteZoneOptions successfully`, func() {
				// Construct an instance of the DeleteZoneOptions model
				zoneIdentifier := "testString"
				deleteZoneOptionsModel := zonesService.NewDeleteZoneOptions(zoneIdentifier)
				deleteZoneOptionsModel.SetZoneIdentifier("testString")
				deleteZoneOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteZoneOptionsModel).ToNot(BeNil())
				Expect(deleteZoneOptionsModel.ZoneIdentifier).To(Equal(core.StringPtr("testString")))
				Expect(deleteZoneOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetZoneOptions successfully`, func() {
				// Construct an instance of the GetZoneOptions model
				zoneIdentifier := "testString"
				getZoneOptionsModel := zonesService.NewGetZoneOptions(zoneIdentifier)
				getZoneOptionsModel.SetZoneIdentifier("testString")
				getZoneOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getZoneOptionsModel).ToNot(BeNil())
				Expect(getZoneOptionsModel.ZoneIdentifier).To(Equal(core.StringPtr("testString")))
				Expect(getZoneOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListZonesOptions successfully`, func() {
				// Construct an instance of the ListZonesOptions model
				listZonesOptionsModel := zonesService.NewListZonesOptions()
				listZonesOptionsModel.SetPage(int64(38))
				listZonesOptionsModel.SetPerPage(int64(5))
				listZonesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listZonesOptionsModel).ToNot(BeNil())
				Expect(listZonesOptionsModel.Page).To(Equal(core.Int64Ptr(int64(38))))
				Expect(listZonesOptionsModel.PerPage).To(Equal(core.Int64Ptr(int64(5))))
				Expect(listZonesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateZoneOptions successfully`, func() {
				// Construct an instance of the UpdateZoneOptions model
				zoneIdentifier := "testString"
				updateZoneOptionsModel := zonesService.NewUpdateZoneOptions(zoneIdentifier)
				updateZoneOptionsModel.SetZoneIdentifier("testString")
				updateZoneOptionsModel.SetPaused(false)
				updateZoneOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateZoneOptionsModel).ToNot(BeNil())
				Expect(updateZoneOptionsModel.ZoneIdentifier).To(Equal(core.StringPtr("testString")))
				Expect(updateZoneOptionsModel.Paused).To(Equal(core.BoolPtr(false)))
				Expect(updateZoneOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewZoneActivationCheckOptions successfully`, func() {
				// Construct an instance of the ZoneActivationCheckOptions model
				zoneIdentifier := "testString"
				zoneActivationCheckOptionsModel := zonesService.NewZoneActivationCheckOptions(zoneIdentifier)
				zoneActivationCheckOptionsModel.SetZoneIdentifier("testString")
				zoneActivationCheckOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(zoneActivationCheckOptionsModel).ToNot(BeNil())
				Expect(zoneActivationCheckOptionsModel.ZoneIdentifier).To(Equal(core.StringPtr("testString")))
				Expect(zoneActivationCheckOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
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
