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

package globalloadbalancermonitorv1_test

import (
	"bytes"
	"context"
	"fmt"
	"github.com/IBM/go-sdk-core/v4/core"
	"github.com/IBM/networking-go-sdk/globalloadbalancermonitorv1"
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

var _ = Describe(`GlobalLoadBalancerMonitorV1`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		crn := "testString"
		It(`Instantiate service client`, func() {
			globalLoadBalancerMonitorService, serviceErr := globalloadbalancermonitorv1.NewGlobalLoadBalancerMonitorV1(&globalloadbalancermonitorv1.GlobalLoadBalancerMonitorV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
				Crn: core.StringPtr(crn),
			})
			Expect(globalLoadBalancerMonitorService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			globalLoadBalancerMonitorService, serviceErr := globalloadbalancermonitorv1.NewGlobalLoadBalancerMonitorV1(&globalloadbalancermonitorv1.GlobalLoadBalancerMonitorV1Options{
				URL: "{BAD_URL_STRING",
				Crn: core.StringPtr(crn),
			})
			Expect(globalLoadBalancerMonitorService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			globalLoadBalancerMonitorService, serviceErr := globalloadbalancermonitorv1.NewGlobalLoadBalancerMonitorV1(&globalloadbalancermonitorv1.GlobalLoadBalancerMonitorV1Options{
				URL: "https://globalloadbalancermonitorv1/api",
				Crn: core.StringPtr(crn),
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(globalLoadBalancerMonitorService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Validation Error`, func() {
			globalLoadBalancerMonitorService, serviceErr := globalloadbalancermonitorv1.NewGlobalLoadBalancerMonitorV1(&globalloadbalancermonitorv1.GlobalLoadBalancerMonitorV1Options{})
			Expect(globalLoadBalancerMonitorService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		crn := "testString"
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"GLOBAL_LOAD_BALANCER_MONITOR_URL": "https://globalloadbalancermonitorv1/api",
				"GLOBAL_LOAD_BALANCER_MONITOR_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				globalLoadBalancerMonitorService, serviceErr := globalloadbalancermonitorv1.NewGlobalLoadBalancerMonitorV1UsingExternalConfig(&globalloadbalancermonitorv1.GlobalLoadBalancerMonitorV1Options{
					Crn: core.StringPtr(crn),
				})
				Expect(globalLoadBalancerMonitorService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := globalLoadBalancerMonitorService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != globalLoadBalancerMonitorService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(globalLoadBalancerMonitorService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(globalLoadBalancerMonitorService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				globalLoadBalancerMonitorService, serviceErr := globalloadbalancermonitorv1.NewGlobalLoadBalancerMonitorV1UsingExternalConfig(&globalloadbalancermonitorv1.GlobalLoadBalancerMonitorV1Options{
					URL: "https://testService/api",
					Crn: core.StringPtr(crn),
				})
				Expect(globalLoadBalancerMonitorService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(globalLoadBalancerMonitorService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := globalLoadBalancerMonitorService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != globalLoadBalancerMonitorService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(globalLoadBalancerMonitorService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(globalLoadBalancerMonitorService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				globalLoadBalancerMonitorService, serviceErr := globalloadbalancermonitorv1.NewGlobalLoadBalancerMonitorV1UsingExternalConfig(&globalloadbalancermonitorv1.GlobalLoadBalancerMonitorV1Options{
					Crn: core.StringPtr(crn),
				})
				err := globalLoadBalancerMonitorService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(globalLoadBalancerMonitorService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(globalLoadBalancerMonitorService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := globalLoadBalancerMonitorService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != globalLoadBalancerMonitorService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(globalLoadBalancerMonitorService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(globalLoadBalancerMonitorService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"GLOBAL_LOAD_BALANCER_MONITOR_URL": "https://globalloadbalancermonitorv1/api",
				"GLOBAL_LOAD_BALANCER_MONITOR_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			globalLoadBalancerMonitorService, serviceErr := globalloadbalancermonitorv1.NewGlobalLoadBalancerMonitorV1UsingExternalConfig(&globalloadbalancermonitorv1.GlobalLoadBalancerMonitorV1Options{
				Crn: core.StringPtr(crn),
			})

			It(`Instantiate service client with error`, func() {
				Expect(globalLoadBalancerMonitorService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"GLOBAL_LOAD_BALANCER_MONITOR_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			globalLoadBalancerMonitorService, serviceErr := globalloadbalancermonitorv1.NewGlobalLoadBalancerMonitorV1UsingExternalConfig(&globalloadbalancermonitorv1.GlobalLoadBalancerMonitorV1Options{
				URL: "{BAD_URL_STRING",
				Crn: core.StringPtr(crn),
			})

			It(`Instantiate service client with error`, func() {
				Expect(globalLoadBalancerMonitorService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = globalloadbalancermonitorv1.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`ListAllLoadBalancerMonitors(listAllLoadBalancerMonitorsOptions *ListAllLoadBalancerMonitorsOptions) - Operation response error`, func() {
		crn := "testString"
		listAllLoadBalancerMonitorsPath := "/v1/testString/load_balancers/monitors"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listAllLoadBalancerMonitorsPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListAllLoadBalancerMonitors with error: Operation response processing error`, func() {
				globalLoadBalancerMonitorService, serviceErr := globalloadbalancermonitorv1.NewGlobalLoadBalancerMonitorV1(&globalloadbalancermonitorv1.GlobalLoadBalancerMonitorV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalLoadBalancerMonitorService).ToNot(BeNil())

				// Construct an instance of the ListAllLoadBalancerMonitorsOptions model
				listAllLoadBalancerMonitorsOptionsModel := new(globalloadbalancermonitorv1.ListAllLoadBalancerMonitorsOptions)
				listAllLoadBalancerMonitorsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := globalLoadBalancerMonitorService.ListAllLoadBalancerMonitors(listAllLoadBalancerMonitorsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				globalLoadBalancerMonitorService.EnableRetries(0, 0)
				result, response, operationErr = globalLoadBalancerMonitorService.ListAllLoadBalancerMonitors(listAllLoadBalancerMonitorsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ListAllLoadBalancerMonitors(listAllLoadBalancerMonitorsOptions *ListAllLoadBalancerMonitorsOptions)`, func() {
		crn := "testString"
		listAllLoadBalancerMonitorsPath := "/v1/testString/load_balancers/monitors"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listAllLoadBalancerMonitorsPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": [{"id": "f1aba936b94213e5b8dca0c0dbf1f9cc", "created_on": "2014-01-01T05:20:00.12345Z", "modified_on": "2014-01-01T05:20:00.12345Z", "type": "http", "description": "Login page monitor", "method": "GET", "port": 8080, "path": "/", "timeout": 5, "retries": 2, "interval": 60, "expected_body": "alive", "expected_codes": "2xx", "follow_redirects": true, "allow_insecure": true, "header": {"mapKey": ["Inner"]}}], "result_info": {"page": 1, "per_page": 20, "count": 1, "total_count": 2000}}`)
				}))
			})
			It(`Invoke ListAllLoadBalancerMonitors successfully`, func() {
				globalLoadBalancerMonitorService, serviceErr := globalloadbalancermonitorv1.NewGlobalLoadBalancerMonitorV1(&globalloadbalancermonitorv1.GlobalLoadBalancerMonitorV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalLoadBalancerMonitorService).ToNot(BeNil())
				globalLoadBalancerMonitorService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := globalLoadBalancerMonitorService.ListAllLoadBalancerMonitors(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListAllLoadBalancerMonitorsOptions model
				listAllLoadBalancerMonitorsOptionsModel := new(globalloadbalancermonitorv1.ListAllLoadBalancerMonitorsOptions)
				listAllLoadBalancerMonitorsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = globalLoadBalancerMonitorService.ListAllLoadBalancerMonitors(listAllLoadBalancerMonitorsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = globalLoadBalancerMonitorService.ListAllLoadBalancerMonitorsWithContext(ctx, listAllLoadBalancerMonitorsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				globalLoadBalancerMonitorService.DisableRetries()
				result, response, operationErr = globalLoadBalancerMonitorService.ListAllLoadBalancerMonitors(listAllLoadBalancerMonitorsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = globalLoadBalancerMonitorService.ListAllLoadBalancerMonitorsWithContext(ctx, listAllLoadBalancerMonitorsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke ListAllLoadBalancerMonitors with error: Operation request error`, func() {
				globalLoadBalancerMonitorService, serviceErr := globalloadbalancermonitorv1.NewGlobalLoadBalancerMonitorV1(&globalloadbalancermonitorv1.GlobalLoadBalancerMonitorV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalLoadBalancerMonitorService).ToNot(BeNil())

				// Construct an instance of the ListAllLoadBalancerMonitorsOptions model
				listAllLoadBalancerMonitorsOptionsModel := new(globalloadbalancermonitorv1.ListAllLoadBalancerMonitorsOptions)
				listAllLoadBalancerMonitorsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := globalLoadBalancerMonitorService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := globalLoadBalancerMonitorService.ListAllLoadBalancerMonitors(listAllLoadBalancerMonitorsOptionsModel)
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
	Describe(`CreateLoadBalancerMonitor(createLoadBalancerMonitorOptions *CreateLoadBalancerMonitorOptions) - Operation response error`, func() {
		crn := "testString"
		createLoadBalancerMonitorPath := "/v1/testString/load_balancers/monitors"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createLoadBalancerMonitorPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateLoadBalancerMonitor with error: Operation response processing error`, func() {
				globalLoadBalancerMonitorService, serviceErr := globalloadbalancermonitorv1.NewGlobalLoadBalancerMonitorV1(&globalloadbalancermonitorv1.GlobalLoadBalancerMonitorV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalLoadBalancerMonitorService).ToNot(BeNil())

				// Construct an instance of the CreateLoadBalancerMonitorOptions model
				createLoadBalancerMonitorOptionsModel := new(globalloadbalancermonitorv1.CreateLoadBalancerMonitorOptions)
				createLoadBalancerMonitorOptionsModel.Type = core.StringPtr("http")
				createLoadBalancerMonitorOptionsModel.Description = core.StringPtr("Login page monitor")
				createLoadBalancerMonitorOptionsModel.Method = core.StringPtr("GET")
				createLoadBalancerMonitorOptionsModel.Port = core.Int64Ptr(int64(8080))
				createLoadBalancerMonitorOptionsModel.Path = core.StringPtr("/")
				createLoadBalancerMonitorOptionsModel.Timeout = core.Int64Ptr(int64(5))
				createLoadBalancerMonitorOptionsModel.Retries = core.Int64Ptr(int64(2))
				createLoadBalancerMonitorOptionsModel.Interval = core.Int64Ptr(int64(60))
				createLoadBalancerMonitorOptionsModel.ExpectedCodes = core.StringPtr("2xx")
				createLoadBalancerMonitorOptionsModel.FollowRedirects = core.BoolPtr(true)
				createLoadBalancerMonitorOptionsModel.ExpectedBody = core.StringPtr("alive")
				createLoadBalancerMonitorOptionsModel.AllowInsecure = core.BoolPtr(true)
				createLoadBalancerMonitorOptionsModel.Header = make(map[string][]string)
				createLoadBalancerMonitorOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := globalLoadBalancerMonitorService.CreateLoadBalancerMonitor(createLoadBalancerMonitorOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				globalLoadBalancerMonitorService.EnableRetries(0, 0)
				result, response, operationErr = globalLoadBalancerMonitorService.CreateLoadBalancerMonitor(createLoadBalancerMonitorOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`CreateLoadBalancerMonitor(createLoadBalancerMonitorOptions *CreateLoadBalancerMonitorOptions)`, func() {
		crn := "testString"
		createLoadBalancerMonitorPath := "/v1/testString/load_balancers/monitors"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createLoadBalancerMonitorPath))
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
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": "f1aba936b94213e5b8dca0c0dbf1f9cc", "created_on": "2014-01-01T05:20:00.12345Z", "modified_on": "2014-01-01T05:20:00.12345Z", "type": "http", "description": "Login page monitor", "method": "GET", "port": 8080, "path": "/", "timeout": 5, "retries": 2, "interval": 60, "expected_body": "alive", "expected_codes": "2xx", "follow_redirects": true, "allow_insecure": true, "header": {"mapKey": ["Inner"]}}}`)
				}))
			})
			It(`Invoke CreateLoadBalancerMonitor successfully`, func() {
				globalLoadBalancerMonitorService, serviceErr := globalloadbalancermonitorv1.NewGlobalLoadBalancerMonitorV1(&globalloadbalancermonitorv1.GlobalLoadBalancerMonitorV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalLoadBalancerMonitorService).ToNot(BeNil())
				globalLoadBalancerMonitorService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := globalLoadBalancerMonitorService.CreateLoadBalancerMonitor(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateLoadBalancerMonitorOptions model
				createLoadBalancerMonitorOptionsModel := new(globalloadbalancermonitorv1.CreateLoadBalancerMonitorOptions)
				createLoadBalancerMonitorOptionsModel.Type = core.StringPtr("http")
				createLoadBalancerMonitorOptionsModel.Description = core.StringPtr("Login page monitor")
				createLoadBalancerMonitorOptionsModel.Method = core.StringPtr("GET")
				createLoadBalancerMonitorOptionsModel.Port = core.Int64Ptr(int64(8080))
				createLoadBalancerMonitorOptionsModel.Path = core.StringPtr("/")
				createLoadBalancerMonitorOptionsModel.Timeout = core.Int64Ptr(int64(5))
				createLoadBalancerMonitorOptionsModel.Retries = core.Int64Ptr(int64(2))
				createLoadBalancerMonitorOptionsModel.Interval = core.Int64Ptr(int64(60))
				createLoadBalancerMonitorOptionsModel.ExpectedCodes = core.StringPtr("2xx")
				createLoadBalancerMonitorOptionsModel.FollowRedirects = core.BoolPtr(true)
				createLoadBalancerMonitorOptionsModel.ExpectedBody = core.StringPtr("alive")
				createLoadBalancerMonitorOptionsModel.AllowInsecure = core.BoolPtr(true)
				createLoadBalancerMonitorOptionsModel.Header = make(map[string][]string)
				createLoadBalancerMonitorOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = globalLoadBalancerMonitorService.CreateLoadBalancerMonitor(createLoadBalancerMonitorOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = globalLoadBalancerMonitorService.CreateLoadBalancerMonitorWithContext(ctx, createLoadBalancerMonitorOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				globalLoadBalancerMonitorService.DisableRetries()
				result, response, operationErr = globalLoadBalancerMonitorService.CreateLoadBalancerMonitor(createLoadBalancerMonitorOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = globalLoadBalancerMonitorService.CreateLoadBalancerMonitorWithContext(ctx, createLoadBalancerMonitorOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke CreateLoadBalancerMonitor with error: Operation request error`, func() {
				globalLoadBalancerMonitorService, serviceErr := globalloadbalancermonitorv1.NewGlobalLoadBalancerMonitorV1(&globalloadbalancermonitorv1.GlobalLoadBalancerMonitorV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalLoadBalancerMonitorService).ToNot(BeNil())

				// Construct an instance of the CreateLoadBalancerMonitorOptions model
				createLoadBalancerMonitorOptionsModel := new(globalloadbalancermonitorv1.CreateLoadBalancerMonitorOptions)
				createLoadBalancerMonitorOptionsModel.Type = core.StringPtr("http")
				createLoadBalancerMonitorOptionsModel.Description = core.StringPtr("Login page monitor")
				createLoadBalancerMonitorOptionsModel.Method = core.StringPtr("GET")
				createLoadBalancerMonitorOptionsModel.Port = core.Int64Ptr(int64(8080))
				createLoadBalancerMonitorOptionsModel.Path = core.StringPtr("/")
				createLoadBalancerMonitorOptionsModel.Timeout = core.Int64Ptr(int64(5))
				createLoadBalancerMonitorOptionsModel.Retries = core.Int64Ptr(int64(2))
				createLoadBalancerMonitorOptionsModel.Interval = core.Int64Ptr(int64(60))
				createLoadBalancerMonitorOptionsModel.ExpectedCodes = core.StringPtr("2xx")
				createLoadBalancerMonitorOptionsModel.FollowRedirects = core.BoolPtr(true)
				createLoadBalancerMonitorOptionsModel.ExpectedBody = core.StringPtr("alive")
				createLoadBalancerMonitorOptionsModel.AllowInsecure = core.BoolPtr(true)
				createLoadBalancerMonitorOptionsModel.Header = make(map[string][]string)
				createLoadBalancerMonitorOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := globalLoadBalancerMonitorService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := globalLoadBalancerMonitorService.CreateLoadBalancerMonitor(createLoadBalancerMonitorOptionsModel)
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
	Describe(`EditLoadBalancerMonitor(editLoadBalancerMonitorOptions *EditLoadBalancerMonitorOptions) - Operation response error`, func() {
		crn := "testString"
		editLoadBalancerMonitorPath := "/v1/testString/load_balancers/monitors/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(editLoadBalancerMonitorPath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke EditLoadBalancerMonitor with error: Operation response processing error`, func() {
				globalLoadBalancerMonitorService, serviceErr := globalloadbalancermonitorv1.NewGlobalLoadBalancerMonitorV1(&globalloadbalancermonitorv1.GlobalLoadBalancerMonitorV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalLoadBalancerMonitorService).ToNot(BeNil())

				// Construct an instance of the EditLoadBalancerMonitorOptions model
				editLoadBalancerMonitorOptionsModel := new(globalloadbalancermonitorv1.EditLoadBalancerMonitorOptions)
				editLoadBalancerMonitorOptionsModel.MonitorIdentifier = core.StringPtr("testString")
				editLoadBalancerMonitorOptionsModel.Type = core.StringPtr("http")
				editLoadBalancerMonitorOptionsModel.Description = core.StringPtr("Login page monitor")
				editLoadBalancerMonitorOptionsModel.Method = core.StringPtr("GET")
				editLoadBalancerMonitorOptionsModel.Port = core.Int64Ptr(int64(8080))
				editLoadBalancerMonitorOptionsModel.Path = core.StringPtr("/")
				editLoadBalancerMonitorOptionsModel.Timeout = core.Int64Ptr(int64(5))
				editLoadBalancerMonitorOptionsModel.Retries = core.Int64Ptr(int64(2))
				editLoadBalancerMonitorOptionsModel.Interval = core.Int64Ptr(int64(60))
				editLoadBalancerMonitorOptionsModel.ExpectedCodes = core.StringPtr("2xx")
				editLoadBalancerMonitorOptionsModel.FollowRedirects = core.BoolPtr(true)
				editLoadBalancerMonitorOptionsModel.ExpectedBody = core.StringPtr("alive")
				editLoadBalancerMonitorOptionsModel.AllowInsecure = core.BoolPtr(true)
				editLoadBalancerMonitorOptionsModel.Header = make(map[string][]string)
				editLoadBalancerMonitorOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := globalLoadBalancerMonitorService.EditLoadBalancerMonitor(editLoadBalancerMonitorOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				globalLoadBalancerMonitorService.EnableRetries(0, 0)
				result, response, operationErr = globalLoadBalancerMonitorService.EditLoadBalancerMonitor(editLoadBalancerMonitorOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`EditLoadBalancerMonitor(editLoadBalancerMonitorOptions *EditLoadBalancerMonitorOptions)`, func() {
		crn := "testString"
		editLoadBalancerMonitorPath := "/v1/testString/load_balancers/monitors/testString"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(editLoadBalancerMonitorPath))
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

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": "f1aba936b94213e5b8dca0c0dbf1f9cc", "created_on": "2014-01-01T05:20:00.12345Z", "modified_on": "2014-01-01T05:20:00.12345Z", "type": "http", "description": "Login page monitor", "method": "GET", "port": 8080, "path": "/", "timeout": 5, "retries": 2, "interval": 60, "expected_body": "alive", "expected_codes": "2xx", "follow_redirects": true, "allow_insecure": true, "header": {"mapKey": ["Inner"]}}}`)
				}))
			})
			It(`Invoke EditLoadBalancerMonitor successfully`, func() {
				globalLoadBalancerMonitorService, serviceErr := globalloadbalancermonitorv1.NewGlobalLoadBalancerMonitorV1(&globalloadbalancermonitorv1.GlobalLoadBalancerMonitorV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalLoadBalancerMonitorService).ToNot(BeNil())
				globalLoadBalancerMonitorService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := globalLoadBalancerMonitorService.EditLoadBalancerMonitor(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the EditLoadBalancerMonitorOptions model
				editLoadBalancerMonitorOptionsModel := new(globalloadbalancermonitorv1.EditLoadBalancerMonitorOptions)
				editLoadBalancerMonitorOptionsModel.MonitorIdentifier = core.StringPtr("testString")
				editLoadBalancerMonitorOptionsModel.Type = core.StringPtr("http")
				editLoadBalancerMonitorOptionsModel.Description = core.StringPtr("Login page monitor")
				editLoadBalancerMonitorOptionsModel.Method = core.StringPtr("GET")
				editLoadBalancerMonitorOptionsModel.Port = core.Int64Ptr(int64(8080))
				editLoadBalancerMonitorOptionsModel.Path = core.StringPtr("/")
				editLoadBalancerMonitorOptionsModel.Timeout = core.Int64Ptr(int64(5))
				editLoadBalancerMonitorOptionsModel.Retries = core.Int64Ptr(int64(2))
				editLoadBalancerMonitorOptionsModel.Interval = core.Int64Ptr(int64(60))
				editLoadBalancerMonitorOptionsModel.ExpectedCodes = core.StringPtr("2xx")
				editLoadBalancerMonitorOptionsModel.FollowRedirects = core.BoolPtr(true)
				editLoadBalancerMonitorOptionsModel.ExpectedBody = core.StringPtr("alive")
				editLoadBalancerMonitorOptionsModel.AllowInsecure = core.BoolPtr(true)
				editLoadBalancerMonitorOptionsModel.Header = make(map[string][]string)
				editLoadBalancerMonitorOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = globalLoadBalancerMonitorService.EditLoadBalancerMonitor(editLoadBalancerMonitorOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = globalLoadBalancerMonitorService.EditLoadBalancerMonitorWithContext(ctx, editLoadBalancerMonitorOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				globalLoadBalancerMonitorService.DisableRetries()
				result, response, operationErr = globalLoadBalancerMonitorService.EditLoadBalancerMonitor(editLoadBalancerMonitorOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = globalLoadBalancerMonitorService.EditLoadBalancerMonitorWithContext(ctx, editLoadBalancerMonitorOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke EditLoadBalancerMonitor with error: Operation validation and request error`, func() {
				globalLoadBalancerMonitorService, serviceErr := globalloadbalancermonitorv1.NewGlobalLoadBalancerMonitorV1(&globalloadbalancermonitorv1.GlobalLoadBalancerMonitorV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalLoadBalancerMonitorService).ToNot(BeNil())

				// Construct an instance of the EditLoadBalancerMonitorOptions model
				editLoadBalancerMonitorOptionsModel := new(globalloadbalancermonitorv1.EditLoadBalancerMonitorOptions)
				editLoadBalancerMonitorOptionsModel.MonitorIdentifier = core.StringPtr("testString")
				editLoadBalancerMonitorOptionsModel.Type = core.StringPtr("http")
				editLoadBalancerMonitorOptionsModel.Description = core.StringPtr("Login page monitor")
				editLoadBalancerMonitorOptionsModel.Method = core.StringPtr("GET")
				editLoadBalancerMonitorOptionsModel.Port = core.Int64Ptr(int64(8080))
				editLoadBalancerMonitorOptionsModel.Path = core.StringPtr("/")
				editLoadBalancerMonitorOptionsModel.Timeout = core.Int64Ptr(int64(5))
				editLoadBalancerMonitorOptionsModel.Retries = core.Int64Ptr(int64(2))
				editLoadBalancerMonitorOptionsModel.Interval = core.Int64Ptr(int64(60))
				editLoadBalancerMonitorOptionsModel.ExpectedCodes = core.StringPtr("2xx")
				editLoadBalancerMonitorOptionsModel.FollowRedirects = core.BoolPtr(true)
				editLoadBalancerMonitorOptionsModel.ExpectedBody = core.StringPtr("alive")
				editLoadBalancerMonitorOptionsModel.AllowInsecure = core.BoolPtr(true)
				editLoadBalancerMonitorOptionsModel.Header = make(map[string][]string)
				editLoadBalancerMonitorOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := globalLoadBalancerMonitorService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := globalLoadBalancerMonitorService.EditLoadBalancerMonitor(editLoadBalancerMonitorOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the EditLoadBalancerMonitorOptions model with no property values
				editLoadBalancerMonitorOptionsModelNew := new(globalloadbalancermonitorv1.EditLoadBalancerMonitorOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = globalLoadBalancerMonitorService.EditLoadBalancerMonitor(editLoadBalancerMonitorOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteLoadBalancerMonitor(deleteLoadBalancerMonitorOptions *DeleteLoadBalancerMonitorOptions) - Operation response error`, func() {
		crn := "testString"
		deleteLoadBalancerMonitorPath := "/v1/testString/load_balancers/monitors/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteLoadBalancerMonitorPath))
					Expect(req.Method).To(Equal("DELETE"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke DeleteLoadBalancerMonitor with error: Operation response processing error`, func() {
				globalLoadBalancerMonitorService, serviceErr := globalloadbalancermonitorv1.NewGlobalLoadBalancerMonitorV1(&globalloadbalancermonitorv1.GlobalLoadBalancerMonitorV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalLoadBalancerMonitorService).ToNot(BeNil())

				// Construct an instance of the DeleteLoadBalancerMonitorOptions model
				deleteLoadBalancerMonitorOptionsModel := new(globalloadbalancermonitorv1.DeleteLoadBalancerMonitorOptions)
				deleteLoadBalancerMonitorOptionsModel.MonitorIdentifier = core.StringPtr("testString")
				deleteLoadBalancerMonitorOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := globalLoadBalancerMonitorService.DeleteLoadBalancerMonitor(deleteLoadBalancerMonitorOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				globalLoadBalancerMonitorService.EnableRetries(0, 0)
				result, response, operationErr = globalLoadBalancerMonitorService.DeleteLoadBalancerMonitor(deleteLoadBalancerMonitorOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`DeleteLoadBalancerMonitor(deleteLoadBalancerMonitorOptions *DeleteLoadBalancerMonitorOptions)`, func() {
		crn := "testString"
		deleteLoadBalancerMonitorPath := "/v1/testString/load_balancers/monitors/testString"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteLoadBalancerMonitorPath))
					Expect(req.Method).To(Equal("DELETE"))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": "f1aba936b94213e5b8dca0c0dbf1f9cc"}}`)
				}))
			})
			It(`Invoke DeleteLoadBalancerMonitor successfully`, func() {
				globalLoadBalancerMonitorService, serviceErr := globalloadbalancermonitorv1.NewGlobalLoadBalancerMonitorV1(&globalloadbalancermonitorv1.GlobalLoadBalancerMonitorV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalLoadBalancerMonitorService).ToNot(BeNil())
				globalLoadBalancerMonitorService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := globalLoadBalancerMonitorService.DeleteLoadBalancerMonitor(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DeleteLoadBalancerMonitorOptions model
				deleteLoadBalancerMonitorOptionsModel := new(globalloadbalancermonitorv1.DeleteLoadBalancerMonitorOptions)
				deleteLoadBalancerMonitorOptionsModel.MonitorIdentifier = core.StringPtr("testString")
				deleteLoadBalancerMonitorOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = globalLoadBalancerMonitorService.DeleteLoadBalancerMonitor(deleteLoadBalancerMonitorOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = globalLoadBalancerMonitorService.DeleteLoadBalancerMonitorWithContext(ctx, deleteLoadBalancerMonitorOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				globalLoadBalancerMonitorService.DisableRetries()
				result, response, operationErr = globalLoadBalancerMonitorService.DeleteLoadBalancerMonitor(deleteLoadBalancerMonitorOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = globalLoadBalancerMonitorService.DeleteLoadBalancerMonitorWithContext(ctx, deleteLoadBalancerMonitorOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke DeleteLoadBalancerMonitor with error: Operation validation and request error`, func() {
				globalLoadBalancerMonitorService, serviceErr := globalloadbalancermonitorv1.NewGlobalLoadBalancerMonitorV1(&globalloadbalancermonitorv1.GlobalLoadBalancerMonitorV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalLoadBalancerMonitorService).ToNot(BeNil())

				// Construct an instance of the DeleteLoadBalancerMonitorOptions model
				deleteLoadBalancerMonitorOptionsModel := new(globalloadbalancermonitorv1.DeleteLoadBalancerMonitorOptions)
				deleteLoadBalancerMonitorOptionsModel.MonitorIdentifier = core.StringPtr("testString")
				deleteLoadBalancerMonitorOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := globalLoadBalancerMonitorService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := globalLoadBalancerMonitorService.DeleteLoadBalancerMonitor(deleteLoadBalancerMonitorOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the DeleteLoadBalancerMonitorOptions model with no property values
				deleteLoadBalancerMonitorOptionsModelNew := new(globalloadbalancermonitorv1.DeleteLoadBalancerMonitorOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = globalLoadBalancerMonitorService.DeleteLoadBalancerMonitor(deleteLoadBalancerMonitorOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetLoadBalancerMonitor(getLoadBalancerMonitorOptions *GetLoadBalancerMonitorOptions) - Operation response error`, func() {
		crn := "testString"
		getLoadBalancerMonitorPath := "/v1/testString/load_balancers/monitors/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getLoadBalancerMonitorPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetLoadBalancerMonitor with error: Operation response processing error`, func() {
				globalLoadBalancerMonitorService, serviceErr := globalloadbalancermonitorv1.NewGlobalLoadBalancerMonitorV1(&globalloadbalancermonitorv1.GlobalLoadBalancerMonitorV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalLoadBalancerMonitorService).ToNot(BeNil())

				// Construct an instance of the GetLoadBalancerMonitorOptions model
				getLoadBalancerMonitorOptionsModel := new(globalloadbalancermonitorv1.GetLoadBalancerMonitorOptions)
				getLoadBalancerMonitorOptionsModel.MonitorIdentifier = core.StringPtr("testString")
				getLoadBalancerMonitorOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := globalLoadBalancerMonitorService.GetLoadBalancerMonitor(getLoadBalancerMonitorOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				globalLoadBalancerMonitorService.EnableRetries(0, 0)
				result, response, operationErr = globalLoadBalancerMonitorService.GetLoadBalancerMonitor(getLoadBalancerMonitorOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetLoadBalancerMonitor(getLoadBalancerMonitorOptions *GetLoadBalancerMonitorOptions)`, func() {
		crn := "testString"
		getLoadBalancerMonitorPath := "/v1/testString/load_balancers/monitors/testString"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getLoadBalancerMonitorPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": "f1aba936b94213e5b8dca0c0dbf1f9cc", "created_on": "2014-01-01T05:20:00.12345Z", "modified_on": "2014-01-01T05:20:00.12345Z", "type": "http", "description": "Login page monitor", "method": "GET", "port": 8080, "path": "/", "timeout": 5, "retries": 2, "interval": 60, "expected_body": "alive", "expected_codes": "2xx", "follow_redirects": true, "allow_insecure": true, "header": {"mapKey": ["Inner"]}}}`)
				}))
			})
			It(`Invoke GetLoadBalancerMonitor successfully`, func() {
				globalLoadBalancerMonitorService, serviceErr := globalloadbalancermonitorv1.NewGlobalLoadBalancerMonitorV1(&globalloadbalancermonitorv1.GlobalLoadBalancerMonitorV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalLoadBalancerMonitorService).ToNot(BeNil())
				globalLoadBalancerMonitorService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := globalLoadBalancerMonitorService.GetLoadBalancerMonitor(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetLoadBalancerMonitorOptions model
				getLoadBalancerMonitorOptionsModel := new(globalloadbalancermonitorv1.GetLoadBalancerMonitorOptions)
				getLoadBalancerMonitorOptionsModel.MonitorIdentifier = core.StringPtr("testString")
				getLoadBalancerMonitorOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = globalLoadBalancerMonitorService.GetLoadBalancerMonitor(getLoadBalancerMonitorOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = globalLoadBalancerMonitorService.GetLoadBalancerMonitorWithContext(ctx, getLoadBalancerMonitorOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				globalLoadBalancerMonitorService.DisableRetries()
				result, response, operationErr = globalLoadBalancerMonitorService.GetLoadBalancerMonitor(getLoadBalancerMonitorOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = globalLoadBalancerMonitorService.GetLoadBalancerMonitorWithContext(ctx, getLoadBalancerMonitorOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke GetLoadBalancerMonitor with error: Operation validation and request error`, func() {
				globalLoadBalancerMonitorService, serviceErr := globalloadbalancermonitorv1.NewGlobalLoadBalancerMonitorV1(&globalloadbalancermonitorv1.GlobalLoadBalancerMonitorV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalLoadBalancerMonitorService).ToNot(BeNil())

				// Construct an instance of the GetLoadBalancerMonitorOptions model
				getLoadBalancerMonitorOptionsModel := new(globalloadbalancermonitorv1.GetLoadBalancerMonitorOptions)
				getLoadBalancerMonitorOptionsModel.MonitorIdentifier = core.StringPtr("testString")
				getLoadBalancerMonitorOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := globalLoadBalancerMonitorService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := globalLoadBalancerMonitorService.GetLoadBalancerMonitor(getLoadBalancerMonitorOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetLoadBalancerMonitorOptions model with no property values
				getLoadBalancerMonitorOptionsModelNew := new(globalloadbalancermonitorv1.GetLoadBalancerMonitorOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = globalLoadBalancerMonitorService.GetLoadBalancerMonitor(getLoadBalancerMonitorOptionsModelNew)
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
			crn := "testString"
			globalLoadBalancerMonitorService, _ := globalloadbalancermonitorv1.NewGlobalLoadBalancerMonitorV1(&globalloadbalancermonitorv1.GlobalLoadBalancerMonitorV1Options{
				URL:           "http://globalloadbalancermonitorv1modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
				Crn: core.StringPtr(crn),
			})
			It(`Invoke NewCreateLoadBalancerMonitorOptions successfully`, func() {
				// Construct an instance of the CreateLoadBalancerMonitorOptions model
				createLoadBalancerMonitorOptionsModel := globalLoadBalancerMonitorService.NewCreateLoadBalancerMonitorOptions()
				createLoadBalancerMonitorOptionsModel.SetType("http")
				createLoadBalancerMonitorOptionsModel.SetDescription("Login page monitor")
				createLoadBalancerMonitorOptionsModel.SetMethod("GET")
				createLoadBalancerMonitorOptionsModel.SetPort(int64(8080))
				createLoadBalancerMonitorOptionsModel.SetPath("/")
				createLoadBalancerMonitorOptionsModel.SetTimeout(int64(5))
				createLoadBalancerMonitorOptionsModel.SetRetries(int64(2))
				createLoadBalancerMonitorOptionsModel.SetInterval(int64(60))
				createLoadBalancerMonitorOptionsModel.SetExpectedCodes("2xx")
				createLoadBalancerMonitorOptionsModel.SetFollowRedirects(true)
				createLoadBalancerMonitorOptionsModel.SetExpectedBody("alive")
				createLoadBalancerMonitorOptionsModel.SetAllowInsecure(true)
				createLoadBalancerMonitorOptionsModel.SetHeader(make(map[string][]string))
				createLoadBalancerMonitorOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createLoadBalancerMonitorOptionsModel).ToNot(BeNil())
				Expect(createLoadBalancerMonitorOptionsModel.Type).To(Equal(core.StringPtr("http")))
				Expect(createLoadBalancerMonitorOptionsModel.Description).To(Equal(core.StringPtr("Login page monitor")))
				Expect(createLoadBalancerMonitorOptionsModel.Method).To(Equal(core.StringPtr("GET")))
				Expect(createLoadBalancerMonitorOptionsModel.Port).To(Equal(core.Int64Ptr(int64(8080))))
				Expect(createLoadBalancerMonitorOptionsModel.Path).To(Equal(core.StringPtr("/")))
				Expect(createLoadBalancerMonitorOptionsModel.Timeout).To(Equal(core.Int64Ptr(int64(5))))
				Expect(createLoadBalancerMonitorOptionsModel.Retries).To(Equal(core.Int64Ptr(int64(2))))
				Expect(createLoadBalancerMonitorOptionsModel.Interval).To(Equal(core.Int64Ptr(int64(60))))
				Expect(createLoadBalancerMonitorOptionsModel.ExpectedCodes).To(Equal(core.StringPtr("2xx")))
				Expect(createLoadBalancerMonitorOptionsModel.FollowRedirects).To(Equal(core.BoolPtr(true)))
				Expect(createLoadBalancerMonitorOptionsModel.ExpectedBody).To(Equal(core.StringPtr("alive")))
				Expect(createLoadBalancerMonitorOptionsModel.AllowInsecure).To(Equal(core.BoolPtr(true)))
				Expect(createLoadBalancerMonitorOptionsModel.Header).To(Equal(make(map[string][]string)))
				Expect(createLoadBalancerMonitorOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteLoadBalancerMonitorOptions successfully`, func() {
				// Construct an instance of the DeleteLoadBalancerMonitorOptions model
				monitorIdentifier := "testString"
				deleteLoadBalancerMonitorOptionsModel := globalLoadBalancerMonitorService.NewDeleteLoadBalancerMonitorOptions(monitorIdentifier)
				deleteLoadBalancerMonitorOptionsModel.SetMonitorIdentifier("testString")
				deleteLoadBalancerMonitorOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteLoadBalancerMonitorOptionsModel).ToNot(BeNil())
				Expect(deleteLoadBalancerMonitorOptionsModel.MonitorIdentifier).To(Equal(core.StringPtr("testString")))
				Expect(deleteLoadBalancerMonitorOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewEditLoadBalancerMonitorOptions successfully`, func() {
				// Construct an instance of the EditLoadBalancerMonitorOptions model
				monitorIdentifier := "testString"
				editLoadBalancerMonitorOptionsModel := globalLoadBalancerMonitorService.NewEditLoadBalancerMonitorOptions(monitorIdentifier)
				editLoadBalancerMonitorOptionsModel.SetMonitorIdentifier("testString")
				editLoadBalancerMonitorOptionsModel.SetType("http")
				editLoadBalancerMonitorOptionsModel.SetDescription("Login page monitor")
				editLoadBalancerMonitorOptionsModel.SetMethod("GET")
				editLoadBalancerMonitorOptionsModel.SetPort(int64(8080))
				editLoadBalancerMonitorOptionsModel.SetPath("/")
				editLoadBalancerMonitorOptionsModel.SetTimeout(int64(5))
				editLoadBalancerMonitorOptionsModel.SetRetries(int64(2))
				editLoadBalancerMonitorOptionsModel.SetInterval(int64(60))
				editLoadBalancerMonitorOptionsModel.SetExpectedCodes("2xx")
				editLoadBalancerMonitorOptionsModel.SetFollowRedirects(true)
				editLoadBalancerMonitorOptionsModel.SetExpectedBody("alive")
				editLoadBalancerMonitorOptionsModel.SetAllowInsecure(true)
				editLoadBalancerMonitorOptionsModel.SetHeader(make(map[string][]string))
				editLoadBalancerMonitorOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(editLoadBalancerMonitorOptionsModel).ToNot(BeNil())
				Expect(editLoadBalancerMonitorOptionsModel.MonitorIdentifier).To(Equal(core.StringPtr("testString")))
				Expect(editLoadBalancerMonitorOptionsModel.Type).To(Equal(core.StringPtr("http")))
				Expect(editLoadBalancerMonitorOptionsModel.Description).To(Equal(core.StringPtr("Login page monitor")))
				Expect(editLoadBalancerMonitorOptionsModel.Method).To(Equal(core.StringPtr("GET")))
				Expect(editLoadBalancerMonitorOptionsModel.Port).To(Equal(core.Int64Ptr(int64(8080))))
				Expect(editLoadBalancerMonitorOptionsModel.Path).To(Equal(core.StringPtr("/")))
				Expect(editLoadBalancerMonitorOptionsModel.Timeout).To(Equal(core.Int64Ptr(int64(5))))
				Expect(editLoadBalancerMonitorOptionsModel.Retries).To(Equal(core.Int64Ptr(int64(2))))
				Expect(editLoadBalancerMonitorOptionsModel.Interval).To(Equal(core.Int64Ptr(int64(60))))
				Expect(editLoadBalancerMonitorOptionsModel.ExpectedCodes).To(Equal(core.StringPtr("2xx")))
				Expect(editLoadBalancerMonitorOptionsModel.FollowRedirects).To(Equal(core.BoolPtr(true)))
				Expect(editLoadBalancerMonitorOptionsModel.ExpectedBody).To(Equal(core.StringPtr("alive")))
				Expect(editLoadBalancerMonitorOptionsModel.AllowInsecure).To(Equal(core.BoolPtr(true)))
				Expect(editLoadBalancerMonitorOptionsModel.Header).To(Equal(make(map[string][]string)))
				Expect(editLoadBalancerMonitorOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetLoadBalancerMonitorOptions successfully`, func() {
				// Construct an instance of the GetLoadBalancerMonitorOptions model
				monitorIdentifier := "testString"
				getLoadBalancerMonitorOptionsModel := globalLoadBalancerMonitorService.NewGetLoadBalancerMonitorOptions(monitorIdentifier)
				getLoadBalancerMonitorOptionsModel.SetMonitorIdentifier("testString")
				getLoadBalancerMonitorOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getLoadBalancerMonitorOptionsModel).ToNot(BeNil())
				Expect(getLoadBalancerMonitorOptionsModel.MonitorIdentifier).To(Equal(core.StringPtr("testString")))
				Expect(getLoadBalancerMonitorOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListAllLoadBalancerMonitorsOptions successfully`, func() {
				// Construct an instance of the ListAllLoadBalancerMonitorsOptions model
				listAllLoadBalancerMonitorsOptionsModel := globalLoadBalancerMonitorService.NewListAllLoadBalancerMonitorsOptions()
				listAllLoadBalancerMonitorsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listAllLoadBalancerMonitorsOptionsModel).ToNot(BeNil())
				Expect(listAllLoadBalancerMonitorsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
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
