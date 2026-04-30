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

package logpushjobsapiv1_test

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
	"github.com/IBM/networking-go-sdk/logpushjobsapiv1"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe(`LogpushJobsApiV1`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		crn := "testString"
		dataset := "testString"
		zoneID := "testString"
		It(`Instantiate service client`, func() {
			logpushJobsApiService, serviceErr := logpushjobsapiv1.NewLogpushJobsApiV1(&logpushjobsapiv1.LogpushJobsApiV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
				Crn:           core.StringPtr(crn),
				Dataset:       core.StringPtr(dataset),
				ZoneID:        core.StringPtr(zoneID),
			})
			Expect(logpushJobsApiService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			logpushJobsApiService, serviceErr := logpushjobsapiv1.NewLogpushJobsApiV1(&logpushjobsapiv1.LogpushJobsApiV1Options{
				URL:     "{BAD_URL_STRING",
				Crn:     core.StringPtr(crn),
				Dataset: core.StringPtr(dataset),
				ZoneID:  core.StringPtr(zoneID),
			})
			Expect(logpushJobsApiService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			logpushJobsApiService, serviceErr := logpushjobsapiv1.NewLogpushJobsApiV1(&logpushjobsapiv1.LogpushJobsApiV1Options{
				URL:     "https://logpushjobsapiv1/api",
				Crn:     core.StringPtr(crn),
				Dataset: core.StringPtr(dataset),
				ZoneID:  core.StringPtr(zoneID),
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(logpushJobsApiService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Validation Error`, func() {
			logpushJobsApiService, serviceErr := logpushjobsapiv1.NewLogpushJobsApiV1(&logpushjobsapiv1.LogpushJobsApiV1Options{})
			Expect(logpushJobsApiService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		crn := "testString"
		dataset := "testString"
		zoneID := "testString"
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"LOGPUSH_JOBS_API_URL":       "https://logpushjobsapiv1/api",
				"LOGPUSH_JOBS_API_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				logpushJobsApiService, serviceErr := logpushjobsapiv1.NewLogpushJobsApiV1UsingExternalConfig(&logpushjobsapiv1.LogpushJobsApiV1Options{
					Crn:     core.StringPtr(crn),
					Dataset: core.StringPtr(dataset),
					ZoneID:  core.StringPtr(zoneID),
				})
				Expect(logpushJobsApiService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := logpushJobsApiService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != logpushJobsApiService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(logpushJobsApiService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(logpushJobsApiService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				logpushJobsApiService, serviceErr := logpushjobsapiv1.NewLogpushJobsApiV1UsingExternalConfig(&logpushjobsapiv1.LogpushJobsApiV1Options{
					URL:     "https://testService/api",
					Crn:     core.StringPtr(crn),
					Dataset: core.StringPtr(dataset),
					ZoneID:  core.StringPtr(zoneID),
				})
				Expect(logpushJobsApiService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(logpushJobsApiService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := logpushJobsApiService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != logpushJobsApiService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(logpushJobsApiService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(logpushJobsApiService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				logpushJobsApiService, serviceErr := logpushjobsapiv1.NewLogpushJobsApiV1UsingExternalConfig(&logpushjobsapiv1.LogpushJobsApiV1Options{
					Crn:     core.StringPtr(crn),
					Dataset: core.StringPtr(dataset),
					ZoneID:  core.StringPtr(zoneID),
				})
				err := logpushJobsApiService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(logpushJobsApiService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(logpushJobsApiService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := logpushJobsApiService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != logpushJobsApiService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(logpushJobsApiService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(logpushJobsApiService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"LOGPUSH_JOBS_API_URL":       "https://logpushjobsapiv1/api",
				"LOGPUSH_JOBS_API_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			logpushJobsApiService, serviceErr := logpushjobsapiv1.NewLogpushJobsApiV1UsingExternalConfig(&logpushjobsapiv1.LogpushJobsApiV1Options{
				Crn:     core.StringPtr(crn),
				Dataset: core.StringPtr(dataset),
				ZoneID:  core.StringPtr(zoneID),
			})

			It(`Instantiate service client with error`, func() {
				Expect(logpushJobsApiService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"LOGPUSH_JOBS_API_AUTH_TYPE": "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			logpushJobsApiService, serviceErr := logpushjobsapiv1.NewLogpushJobsApiV1UsingExternalConfig(&logpushjobsapiv1.LogpushJobsApiV1Options{
				URL:     "{BAD_URL_STRING",
				Crn:     core.StringPtr(crn),
				Dataset: core.StringPtr(dataset),
				ZoneID:  core.StringPtr(zoneID),
			})

			It(`Instantiate service client with error`, func() {
				Expect(logpushJobsApiService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = logpushjobsapiv1.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`GetLogpushJobsV2(getLogpushJobsV2Options *GetLogpushJobsV2Options) - Operation response error`, func() {
		crn := "testString"
		dataset := "testString"
		zoneID := "testString"
		getLogpushJobsV2Path := "/v2/testString/zones/testString/logpush/jobs"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getLogpushJobsV2Path))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetLogpushJobsV2 with error: Operation response processing error`, func() {
				logpushJobsApiService, serviceErr := logpushjobsapiv1.NewLogpushJobsApiV1(&logpushjobsapiv1.LogpushJobsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					Dataset:       core.StringPtr(dataset),
					ZoneID:        core.StringPtr(zoneID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(logpushJobsApiService).ToNot(BeNil())

				// Construct an instance of the GetLogpushJobsV2Options model
				getLogpushJobsV2OptionsModel := new(logpushjobsapiv1.GetLogpushJobsV2Options)
				getLogpushJobsV2OptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := logpushJobsApiService.GetLogpushJobsV2(getLogpushJobsV2OptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				logpushJobsApiService.EnableRetries(0, 0)
				result, response, operationErr = logpushJobsApiService.GetLogpushJobsV2(getLogpushJobsV2OptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetLogpushJobsV2(getLogpushJobsV2Options *GetLogpushJobsV2Options)`, func() {
		crn := "testString"
		dataset := "testString"
		zoneID := "testString"
		getLogpushJobsV2Path := "/v2/testString/zones/testString/logpush/jobs"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getLogpushJobsV2Path))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": [{"id": 5850, "name": "My log push job", "enabled": false, "dataset": "firewall_events", "frequency": "high", "logpull_options": "timestamps=rfc3339&timestamps=rfc3339", "destination_conf": "cos://cos-bucket001?region=us-south&instance-id=231f5467-3072-4cb9-9e39-a906fa3032ea", "last_complete": "2022-01-15T16:33:31.834209Z", "last_error": "2022-01-15T16:33:31.834209Z", "error_message": "ErrorMessage"}]}`)
				}))
			})
			It(`Invoke GetLogpushJobsV2 successfully with retries`, func() {
				logpushJobsApiService, serviceErr := logpushjobsapiv1.NewLogpushJobsApiV1(&logpushjobsapiv1.LogpushJobsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					Dataset:       core.StringPtr(dataset),
					ZoneID:        core.StringPtr(zoneID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(logpushJobsApiService).ToNot(BeNil())
				logpushJobsApiService.EnableRetries(0, 0)

				// Construct an instance of the GetLogpushJobsV2Options model
				getLogpushJobsV2OptionsModel := new(logpushjobsapiv1.GetLogpushJobsV2Options)
				getLogpushJobsV2OptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := logpushJobsApiService.GetLogpushJobsV2WithContext(ctx, getLogpushJobsV2OptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				logpushJobsApiService.DisableRetries()
				result, response, operationErr := logpushJobsApiService.GetLogpushJobsV2(getLogpushJobsV2OptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = logpushJobsApiService.GetLogpushJobsV2WithContext(ctx, getLogpushJobsV2OptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getLogpushJobsV2Path))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": [{"id": 5850, "name": "My log push job", "enabled": false, "dataset": "firewall_events", "frequency": "high", "logpull_options": "timestamps=rfc3339&timestamps=rfc3339", "destination_conf": "cos://cos-bucket001?region=us-south&instance-id=231f5467-3072-4cb9-9e39-a906fa3032ea", "last_complete": "2022-01-15T16:33:31.834209Z", "last_error": "2022-01-15T16:33:31.834209Z", "error_message": "ErrorMessage"}]}`)
				}))
			})
			It(`Invoke GetLogpushJobsV2 successfully`, func() {
				logpushJobsApiService, serviceErr := logpushjobsapiv1.NewLogpushJobsApiV1(&logpushjobsapiv1.LogpushJobsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					Dataset:       core.StringPtr(dataset),
					ZoneID:        core.StringPtr(zoneID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(logpushJobsApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := logpushJobsApiService.GetLogpushJobsV2(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetLogpushJobsV2Options model
				getLogpushJobsV2OptionsModel := new(logpushjobsapiv1.GetLogpushJobsV2Options)
				getLogpushJobsV2OptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = logpushJobsApiService.GetLogpushJobsV2(getLogpushJobsV2OptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetLogpushJobsV2 with error: Operation request error`, func() {
				logpushJobsApiService, serviceErr := logpushjobsapiv1.NewLogpushJobsApiV1(&logpushjobsapiv1.LogpushJobsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					Dataset:       core.StringPtr(dataset),
					ZoneID:        core.StringPtr(zoneID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(logpushJobsApiService).ToNot(BeNil())

				// Construct an instance of the GetLogpushJobsV2Options model
				getLogpushJobsV2OptionsModel := new(logpushjobsapiv1.GetLogpushJobsV2Options)
				getLogpushJobsV2OptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := logpushJobsApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := logpushJobsApiService.GetLogpushJobsV2(getLogpushJobsV2OptionsModel)
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
			It(`Invoke GetLogpushJobsV2 successfully`, func() {
				logpushJobsApiService, serviceErr := logpushjobsapiv1.NewLogpushJobsApiV1(&logpushjobsapiv1.LogpushJobsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					Dataset:       core.StringPtr(dataset),
					ZoneID:        core.StringPtr(zoneID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(logpushJobsApiService).ToNot(BeNil())

				// Construct an instance of the GetLogpushJobsV2Options model
				getLogpushJobsV2OptionsModel := new(logpushjobsapiv1.GetLogpushJobsV2Options)
				getLogpushJobsV2OptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := logpushJobsApiService.GetLogpushJobsV2(getLogpushJobsV2OptionsModel)
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
	Describe(`CreateLogpushJobV2(createLogpushJobV2Options *CreateLogpushJobV2Options) - Operation response error`, func() {
		crn := "testString"
		dataset := "testString"
		zoneID := "testString"
		createLogpushJobV2Path := "/v2/testString/zones/testString/logpush/jobs"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createLogpushJobV2Path))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateLogpushJobV2 with error: Operation response processing error`, func() {
				logpushJobsApiService, serviceErr := logpushjobsapiv1.NewLogpushJobsApiV1(&logpushjobsapiv1.LogpushJobsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					Dataset:       core.StringPtr(dataset),
					ZoneID:        core.StringPtr(zoneID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(logpushJobsApiService).ToNot(BeNil())

				// Construct an instance of the CreateLogpushJobV2RequestLogpushJobCosReq model
				createLogpushJobV2RequestModel := new(logpushjobsapiv1.CreateLogpushJobV2RequestLogpushJobCosReq)
				createLogpushJobV2RequestModel.Name = core.StringPtr("My log push job")
				createLogpushJobV2RequestModel.Enabled = core.BoolPtr(false)
				createLogpushJobV2RequestModel.LogpullOptions = core.StringPtr("timestamps=rfc3339&timestamps=rfc3339")
				createLogpushJobV2RequestModel.Cos = map[string]interface{}{"anyKey": "anyValue"}
				createLogpushJobV2RequestModel.OwnershipChallenge = core.StringPtr("00000000000000000000000000000000")
				createLogpushJobV2RequestModel.Dataset = core.StringPtr("http_requests")
				createLogpushJobV2RequestModel.Frequency = core.StringPtr("high")

				// Construct an instance of the CreateLogpushJobV2Options model
				createLogpushJobV2OptionsModel := new(logpushjobsapiv1.CreateLogpushJobV2Options)
				createLogpushJobV2OptionsModel.CreateLogpushJobV2Request = createLogpushJobV2RequestModel
				createLogpushJobV2OptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := logpushJobsApiService.CreateLogpushJobV2(createLogpushJobV2OptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				logpushJobsApiService.EnableRetries(0, 0)
				result, response, operationErr = logpushJobsApiService.CreateLogpushJobV2(createLogpushJobV2OptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateLogpushJobV2(createLogpushJobV2Options *CreateLogpushJobV2Options)`, func() {
		crn := "testString"
		dataset := "testString"
		zoneID := "testString"
		createLogpushJobV2Path := "/v2/testString/zones/testString/logpush/jobs"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createLogpushJobV2Path))
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
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": 5850, "name": "My log push job", "enabled": false, "dataset": "firewall_events", "frequency": "high", "logpull_options": "timestamps=rfc3339&timestamps=rfc3339", "destination_conf": "cos://cos-bucket001?region=us-south&instance-id=231f5467-3072-4cb9-9e39-a906fa3032ea", "last_complete": "2022-01-15T16:33:31.834209Z", "last_error": "2022-01-15T16:33:31.834209Z", "error_message": "ErrorMessage"}}`)
				}))
			})
			It(`Invoke CreateLogpushJobV2 successfully with retries`, func() {
				logpushJobsApiService, serviceErr := logpushjobsapiv1.NewLogpushJobsApiV1(&logpushjobsapiv1.LogpushJobsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					Dataset:       core.StringPtr(dataset),
					ZoneID:        core.StringPtr(zoneID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(logpushJobsApiService).ToNot(BeNil())
				logpushJobsApiService.EnableRetries(0, 0)

				// Construct an instance of the CreateLogpushJobV2RequestLogpushJobCosReq model
				createLogpushJobV2RequestModel := new(logpushjobsapiv1.CreateLogpushJobV2RequestLogpushJobCosReq)
				createLogpushJobV2RequestModel.Name = core.StringPtr("My log push job")
				createLogpushJobV2RequestModel.Enabled = core.BoolPtr(false)
				createLogpushJobV2RequestModel.LogpullOptions = core.StringPtr("timestamps=rfc3339&timestamps=rfc3339")
				createLogpushJobV2RequestModel.Cos = map[string]interface{}{"anyKey": "anyValue"}
				createLogpushJobV2RequestModel.OwnershipChallenge = core.StringPtr("00000000000000000000000000000000")
				createLogpushJobV2RequestModel.Dataset = core.StringPtr("http_requests")
				createLogpushJobV2RequestModel.Frequency = core.StringPtr("high")

				// Construct an instance of the CreateLogpushJobV2Options model
				createLogpushJobV2OptionsModel := new(logpushjobsapiv1.CreateLogpushJobV2Options)
				createLogpushJobV2OptionsModel.CreateLogpushJobV2Request = createLogpushJobV2RequestModel
				createLogpushJobV2OptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := logpushJobsApiService.CreateLogpushJobV2WithContext(ctx, createLogpushJobV2OptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				logpushJobsApiService.DisableRetries()
				result, response, operationErr := logpushJobsApiService.CreateLogpushJobV2(createLogpushJobV2OptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = logpushJobsApiService.CreateLogpushJobV2WithContext(ctx, createLogpushJobV2OptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createLogpushJobV2Path))
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
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": 5850, "name": "My log push job", "enabled": false, "dataset": "firewall_events", "frequency": "high", "logpull_options": "timestamps=rfc3339&timestamps=rfc3339", "destination_conf": "cos://cos-bucket001?region=us-south&instance-id=231f5467-3072-4cb9-9e39-a906fa3032ea", "last_complete": "2022-01-15T16:33:31.834209Z", "last_error": "2022-01-15T16:33:31.834209Z", "error_message": "ErrorMessage"}}`)
				}))
			})
			It(`Invoke CreateLogpushJobV2 successfully`, func() {
				logpushJobsApiService, serviceErr := logpushjobsapiv1.NewLogpushJobsApiV1(&logpushjobsapiv1.LogpushJobsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					Dataset:       core.StringPtr(dataset),
					ZoneID:        core.StringPtr(zoneID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(logpushJobsApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := logpushJobsApiService.CreateLogpushJobV2(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateLogpushJobV2RequestLogpushJobCosReq model
				createLogpushJobV2RequestModel := new(logpushjobsapiv1.CreateLogpushJobV2RequestLogpushJobCosReq)
				createLogpushJobV2RequestModel.Name = core.StringPtr("My log push job")
				createLogpushJobV2RequestModel.Enabled = core.BoolPtr(false)
				createLogpushJobV2RequestModel.LogpullOptions = core.StringPtr("timestamps=rfc3339&timestamps=rfc3339")
				createLogpushJobV2RequestModel.Cos = map[string]interface{}{"anyKey": "anyValue"}
				createLogpushJobV2RequestModel.OwnershipChallenge = core.StringPtr("00000000000000000000000000000000")
				createLogpushJobV2RequestModel.Dataset = core.StringPtr("http_requests")
				createLogpushJobV2RequestModel.Frequency = core.StringPtr("high")

				// Construct an instance of the CreateLogpushJobV2Options model
				createLogpushJobV2OptionsModel := new(logpushjobsapiv1.CreateLogpushJobV2Options)
				createLogpushJobV2OptionsModel.CreateLogpushJobV2Request = createLogpushJobV2RequestModel
				createLogpushJobV2OptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = logpushJobsApiService.CreateLogpushJobV2(createLogpushJobV2OptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateLogpushJobV2 with error: Operation request error`, func() {
				logpushJobsApiService, serviceErr := logpushjobsapiv1.NewLogpushJobsApiV1(&logpushjobsapiv1.LogpushJobsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					Dataset:       core.StringPtr(dataset),
					ZoneID:        core.StringPtr(zoneID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(logpushJobsApiService).ToNot(BeNil())

				// Construct an instance of the CreateLogpushJobV2RequestLogpushJobCosReq model
				createLogpushJobV2RequestModel := new(logpushjobsapiv1.CreateLogpushJobV2RequestLogpushJobCosReq)
				createLogpushJobV2RequestModel.Name = core.StringPtr("My log push job")
				createLogpushJobV2RequestModel.Enabled = core.BoolPtr(false)
				createLogpushJobV2RequestModel.LogpullOptions = core.StringPtr("timestamps=rfc3339&timestamps=rfc3339")
				createLogpushJobV2RequestModel.Cos = map[string]interface{}{"anyKey": "anyValue"}
				createLogpushJobV2RequestModel.OwnershipChallenge = core.StringPtr("00000000000000000000000000000000")
				createLogpushJobV2RequestModel.Dataset = core.StringPtr("http_requests")
				createLogpushJobV2RequestModel.Frequency = core.StringPtr("high")

				// Construct an instance of the CreateLogpushJobV2Options model
				createLogpushJobV2OptionsModel := new(logpushjobsapiv1.CreateLogpushJobV2Options)
				createLogpushJobV2OptionsModel.CreateLogpushJobV2Request = createLogpushJobV2RequestModel
				createLogpushJobV2OptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := logpushJobsApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := logpushJobsApiService.CreateLogpushJobV2(createLogpushJobV2OptionsModel)
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
			It(`Invoke CreateLogpushJobV2 successfully`, func() {
				logpushJobsApiService, serviceErr := logpushjobsapiv1.NewLogpushJobsApiV1(&logpushjobsapiv1.LogpushJobsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					Dataset:       core.StringPtr(dataset),
					ZoneID:        core.StringPtr(zoneID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(logpushJobsApiService).ToNot(BeNil())

				// Construct an instance of the CreateLogpushJobV2RequestLogpushJobCosReq model
				createLogpushJobV2RequestModel := new(logpushjobsapiv1.CreateLogpushJobV2RequestLogpushJobCosReq)
				createLogpushJobV2RequestModel.Name = core.StringPtr("My log push job")
				createLogpushJobV2RequestModel.Enabled = core.BoolPtr(false)
				createLogpushJobV2RequestModel.LogpullOptions = core.StringPtr("timestamps=rfc3339&timestamps=rfc3339")
				createLogpushJobV2RequestModel.Cos = map[string]interface{}{"anyKey": "anyValue"}
				createLogpushJobV2RequestModel.OwnershipChallenge = core.StringPtr("00000000000000000000000000000000")
				createLogpushJobV2RequestModel.Dataset = core.StringPtr("http_requests")
				createLogpushJobV2RequestModel.Frequency = core.StringPtr("high")

				// Construct an instance of the CreateLogpushJobV2Options model
				createLogpushJobV2OptionsModel := new(logpushjobsapiv1.CreateLogpushJobV2Options)
				createLogpushJobV2OptionsModel.CreateLogpushJobV2Request = createLogpushJobV2RequestModel
				createLogpushJobV2OptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := logpushJobsApiService.CreateLogpushJobV2(createLogpushJobV2OptionsModel)
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
	Describe(`GetLogpushJobV2(getLogpushJobV2Options *GetLogpushJobV2Options) - Operation response error`, func() {
		crn := "testString"
		dataset := "testString"
		zoneID := "testString"
		getLogpushJobV2Path := "/v2/testString/zones/testString/logpush/jobs/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getLogpushJobV2Path))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetLogpushJobV2 with error: Operation response processing error`, func() {
				logpushJobsApiService, serviceErr := logpushjobsapiv1.NewLogpushJobsApiV1(&logpushjobsapiv1.LogpushJobsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					Dataset:       core.StringPtr(dataset),
					ZoneID:        core.StringPtr(zoneID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(logpushJobsApiService).ToNot(BeNil())

				// Construct an instance of the GetLogpushJobV2Options model
				getLogpushJobV2OptionsModel := new(logpushjobsapiv1.GetLogpushJobV2Options)
				getLogpushJobV2OptionsModel.JobID = core.StringPtr("testString")
				getLogpushJobV2OptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := logpushJobsApiService.GetLogpushJobV2(getLogpushJobV2OptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				logpushJobsApiService.EnableRetries(0, 0)
				result, response, operationErr = logpushJobsApiService.GetLogpushJobV2(getLogpushJobV2OptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetLogpushJobV2(getLogpushJobV2Options *GetLogpushJobV2Options)`, func() {
		crn := "testString"
		dataset := "testString"
		zoneID := "testString"
		getLogpushJobV2Path := "/v2/testString/zones/testString/logpush/jobs/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getLogpushJobV2Path))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": 5850, "name": "My log push job", "enabled": false, "dataset": "firewall_events", "frequency": "high", "logpull_options": "timestamps=rfc3339&timestamps=rfc3339", "destination_conf": "cos://cos-bucket001?region=us-south&instance-id=231f5467-3072-4cb9-9e39-a906fa3032ea", "last_complete": "2022-01-15T16:33:31.834209Z", "last_error": "2022-01-15T16:33:31.834209Z", "error_message": "ErrorMessage"}}`)
				}))
			})
			It(`Invoke GetLogpushJobV2 successfully with retries`, func() {
				logpushJobsApiService, serviceErr := logpushjobsapiv1.NewLogpushJobsApiV1(&logpushjobsapiv1.LogpushJobsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					Dataset:       core.StringPtr(dataset),
					ZoneID:        core.StringPtr(zoneID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(logpushJobsApiService).ToNot(BeNil())
				logpushJobsApiService.EnableRetries(0, 0)

				// Construct an instance of the GetLogpushJobV2Options model
				getLogpushJobV2OptionsModel := new(logpushjobsapiv1.GetLogpushJobV2Options)
				getLogpushJobV2OptionsModel.JobID = core.StringPtr("testString")
				getLogpushJobV2OptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := logpushJobsApiService.GetLogpushJobV2WithContext(ctx, getLogpushJobV2OptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				logpushJobsApiService.DisableRetries()
				result, response, operationErr := logpushJobsApiService.GetLogpushJobV2(getLogpushJobV2OptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = logpushJobsApiService.GetLogpushJobV2WithContext(ctx, getLogpushJobV2OptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getLogpushJobV2Path))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": 5850, "name": "My log push job", "enabled": false, "dataset": "firewall_events", "frequency": "high", "logpull_options": "timestamps=rfc3339&timestamps=rfc3339", "destination_conf": "cos://cos-bucket001?region=us-south&instance-id=231f5467-3072-4cb9-9e39-a906fa3032ea", "last_complete": "2022-01-15T16:33:31.834209Z", "last_error": "2022-01-15T16:33:31.834209Z", "error_message": "ErrorMessage"}}`)
				}))
			})
			It(`Invoke GetLogpushJobV2 successfully`, func() {
				logpushJobsApiService, serviceErr := logpushjobsapiv1.NewLogpushJobsApiV1(&logpushjobsapiv1.LogpushJobsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					Dataset:       core.StringPtr(dataset),
					ZoneID:        core.StringPtr(zoneID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(logpushJobsApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := logpushJobsApiService.GetLogpushJobV2(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetLogpushJobV2Options model
				getLogpushJobV2OptionsModel := new(logpushjobsapiv1.GetLogpushJobV2Options)
				getLogpushJobV2OptionsModel.JobID = core.StringPtr("testString")
				getLogpushJobV2OptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = logpushJobsApiService.GetLogpushJobV2(getLogpushJobV2OptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetLogpushJobV2 with error: Operation validation and request error`, func() {
				logpushJobsApiService, serviceErr := logpushjobsapiv1.NewLogpushJobsApiV1(&logpushjobsapiv1.LogpushJobsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					Dataset:       core.StringPtr(dataset),
					ZoneID:        core.StringPtr(zoneID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(logpushJobsApiService).ToNot(BeNil())

				// Construct an instance of the GetLogpushJobV2Options model
				getLogpushJobV2OptionsModel := new(logpushjobsapiv1.GetLogpushJobV2Options)
				getLogpushJobV2OptionsModel.JobID = core.StringPtr("testString")
				getLogpushJobV2OptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := logpushJobsApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := logpushJobsApiService.GetLogpushJobV2(getLogpushJobV2OptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetLogpushJobV2Options model with no property values
				getLogpushJobV2OptionsModelNew := new(logpushjobsapiv1.GetLogpushJobV2Options)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = logpushJobsApiService.GetLogpushJobV2(getLogpushJobV2OptionsModelNew)
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
			It(`Invoke GetLogpushJobV2 successfully`, func() {
				logpushJobsApiService, serviceErr := logpushjobsapiv1.NewLogpushJobsApiV1(&logpushjobsapiv1.LogpushJobsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					Dataset:       core.StringPtr(dataset),
					ZoneID:        core.StringPtr(zoneID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(logpushJobsApiService).ToNot(BeNil())

				// Construct an instance of the GetLogpushJobV2Options model
				getLogpushJobV2OptionsModel := new(logpushjobsapiv1.GetLogpushJobV2Options)
				getLogpushJobV2OptionsModel.JobID = core.StringPtr("testString")
				getLogpushJobV2OptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := logpushJobsApiService.GetLogpushJobV2(getLogpushJobV2OptionsModel)
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
	Describe(`UpdateLogpushJobV2(updateLogpushJobV2Options *UpdateLogpushJobV2Options) - Operation response error`, func() {
		crn := "testString"
		dataset := "testString"
		zoneID := "testString"
		updateLogpushJobV2Path := "/v2/testString/zones/testString/logpush/jobs/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateLogpushJobV2Path))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateLogpushJobV2 with error: Operation response processing error`, func() {
				logpushJobsApiService, serviceErr := logpushjobsapiv1.NewLogpushJobsApiV1(&logpushjobsapiv1.LogpushJobsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					Dataset:       core.StringPtr(dataset),
					ZoneID:        core.StringPtr(zoneID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(logpushJobsApiService).ToNot(BeNil())

				// Construct an instance of the UpdateLogpushJobV2RequestLogpushJobsUpdateCosReq model
				updateLogpushJobV2RequestModel := new(logpushjobsapiv1.UpdateLogpushJobV2RequestLogpushJobsUpdateCosReq)
				updateLogpushJobV2RequestModel.Enabled = core.BoolPtr(false)
				updateLogpushJobV2RequestModel.LogpullOptions = core.StringPtr("timestamps=rfc3339&timestamps=rfc3339")
				updateLogpushJobV2RequestModel.Cos = map[string]interface{}{"anyKey": "anyValue"}
				updateLogpushJobV2RequestModel.OwnershipChallenge = core.StringPtr("00000000000000000000000000000000")
				updateLogpushJobV2RequestModel.Frequency = core.StringPtr("high")

				// Construct an instance of the UpdateLogpushJobV2Options model
				updateLogpushJobV2OptionsModel := new(logpushjobsapiv1.UpdateLogpushJobV2Options)
				updateLogpushJobV2OptionsModel.JobID = core.StringPtr("testString")
				updateLogpushJobV2OptionsModel.UpdateLogpushJobV2Request = updateLogpushJobV2RequestModel
				updateLogpushJobV2OptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := logpushJobsApiService.UpdateLogpushJobV2(updateLogpushJobV2OptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				logpushJobsApiService.EnableRetries(0, 0)
				result, response, operationErr = logpushJobsApiService.UpdateLogpushJobV2(updateLogpushJobV2OptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateLogpushJobV2(updateLogpushJobV2Options *UpdateLogpushJobV2Options)`, func() {
		crn := "testString"
		dataset := "testString"
		zoneID := "testString"
		updateLogpushJobV2Path := "/v2/testString/zones/testString/logpush/jobs/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateLogpushJobV2Path))
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
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": 5850, "name": "My log push job", "enabled": false, "dataset": "firewall_events", "frequency": "high", "logpull_options": "timestamps=rfc3339&timestamps=rfc3339", "destination_conf": "cos://cos-bucket001?region=us-south&instance-id=231f5467-3072-4cb9-9e39-a906fa3032ea", "last_complete": "2022-01-15T16:33:31.834209Z", "last_error": "2022-01-15T16:33:31.834209Z", "error_message": "ErrorMessage"}}`)
				}))
			})
			It(`Invoke UpdateLogpushJobV2 successfully with retries`, func() {
				logpushJobsApiService, serviceErr := logpushjobsapiv1.NewLogpushJobsApiV1(&logpushjobsapiv1.LogpushJobsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					Dataset:       core.StringPtr(dataset),
					ZoneID:        core.StringPtr(zoneID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(logpushJobsApiService).ToNot(BeNil())
				logpushJobsApiService.EnableRetries(0, 0)

				// Construct an instance of the UpdateLogpushJobV2RequestLogpushJobsUpdateCosReq model
				updateLogpushJobV2RequestModel := new(logpushjobsapiv1.UpdateLogpushJobV2RequestLogpushJobsUpdateCosReq)
				updateLogpushJobV2RequestModel.Enabled = core.BoolPtr(false)
				updateLogpushJobV2RequestModel.LogpullOptions = core.StringPtr("timestamps=rfc3339&timestamps=rfc3339")
				updateLogpushJobV2RequestModel.Cos = map[string]interface{}{"anyKey": "anyValue"}
				updateLogpushJobV2RequestModel.OwnershipChallenge = core.StringPtr("00000000000000000000000000000000")
				updateLogpushJobV2RequestModel.Frequency = core.StringPtr("high")

				// Construct an instance of the UpdateLogpushJobV2Options model
				updateLogpushJobV2OptionsModel := new(logpushjobsapiv1.UpdateLogpushJobV2Options)
				updateLogpushJobV2OptionsModel.JobID = core.StringPtr("testString")
				updateLogpushJobV2OptionsModel.UpdateLogpushJobV2Request = updateLogpushJobV2RequestModel
				updateLogpushJobV2OptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := logpushJobsApiService.UpdateLogpushJobV2WithContext(ctx, updateLogpushJobV2OptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				logpushJobsApiService.DisableRetries()
				result, response, operationErr := logpushJobsApiService.UpdateLogpushJobV2(updateLogpushJobV2OptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = logpushJobsApiService.UpdateLogpushJobV2WithContext(ctx, updateLogpushJobV2OptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(updateLogpushJobV2Path))
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

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": 5850, "name": "My log push job", "enabled": false, "dataset": "firewall_events", "frequency": "high", "logpull_options": "timestamps=rfc3339&timestamps=rfc3339", "destination_conf": "cos://cos-bucket001?region=us-south&instance-id=231f5467-3072-4cb9-9e39-a906fa3032ea", "last_complete": "2022-01-15T16:33:31.834209Z", "last_error": "2022-01-15T16:33:31.834209Z", "error_message": "ErrorMessage"}}`)
				}))
			})
			It(`Invoke UpdateLogpushJobV2 successfully`, func() {
				logpushJobsApiService, serviceErr := logpushjobsapiv1.NewLogpushJobsApiV1(&logpushjobsapiv1.LogpushJobsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					Dataset:       core.StringPtr(dataset),
					ZoneID:        core.StringPtr(zoneID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(logpushJobsApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := logpushJobsApiService.UpdateLogpushJobV2(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateLogpushJobV2RequestLogpushJobsUpdateCosReq model
				updateLogpushJobV2RequestModel := new(logpushjobsapiv1.UpdateLogpushJobV2RequestLogpushJobsUpdateCosReq)
				updateLogpushJobV2RequestModel.Enabled = core.BoolPtr(false)
				updateLogpushJobV2RequestModel.LogpullOptions = core.StringPtr("timestamps=rfc3339&timestamps=rfc3339")
				updateLogpushJobV2RequestModel.Cos = map[string]interface{}{"anyKey": "anyValue"}
				updateLogpushJobV2RequestModel.OwnershipChallenge = core.StringPtr("00000000000000000000000000000000")
				updateLogpushJobV2RequestModel.Frequency = core.StringPtr("high")

				// Construct an instance of the UpdateLogpushJobV2Options model
				updateLogpushJobV2OptionsModel := new(logpushjobsapiv1.UpdateLogpushJobV2Options)
				updateLogpushJobV2OptionsModel.JobID = core.StringPtr("testString")
				updateLogpushJobV2OptionsModel.UpdateLogpushJobV2Request = updateLogpushJobV2RequestModel
				updateLogpushJobV2OptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = logpushJobsApiService.UpdateLogpushJobV2(updateLogpushJobV2OptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateLogpushJobV2 with error: Operation validation and request error`, func() {
				logpushJobsApiService, serviceErr := logpushjobsapiv1.NewLogpushJobsApiV1(&logpushjobsapiv1.LogpushJobsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					Dataset:       core.StringPtr(dataset),
					ZoneID:        core.StringPtr(zoneID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(logpushJobsApiService).ToNot(BeNil())

				// Construct an instance of the UpdateLogpushJobV2RequestLogpushJobsUpdateCosReq model
				updateLogpushJobV2RequestModel := new(logpushjobsapiv1.UpdateLogpushJobV2RequestLogpushJobsUpdateCosReq)
				updateLogpushJobV2RequestModel.Enabled = core.BoolPtr(false)
				updateLogpushJobV2RequestModel.LogpullOptions = core.StringPtr("timestamps=rfc3339&timestamps=rfc3339")
				updateLogpushJobV2RequestModel.Cos = map[string]interface{}{"anyKey": "anyValue"}
				updateLogpushJobV2RequestModel.OwnershipChallenge = core.StringPtr("00000000000000000000000000000000")
				updateLogpushJobV2RequestModel.Frequency = core.StringPtr("high")

				// Construct an instance of the UpdateLogpushJobV2Options model
				updateLogpushJobV2OptionsModel := new(logpushjobsapiv1.UpdateLogpushJobV2Options)
				updateLogpushJobV2OptionsModel.JobID = core.StringPtr("testString")
				updateLogpushJobV2OptionsModel.UpdateLogpushJobV2Request = updateLogpushJobV2RequestModel
				updateLogpushJobV2OptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := logpushJobsApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := logpushJobsApiService.UpdateLogpushJobV2(updateLogpushJobV2OptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateLogpushJobV2Options model with no property values
				updateLogpushJobV2OptionsModelNew := new(logpushjobsapiv1.UpdateLogpushJobV2Options)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = logpushJobsApiService.UpdateLogpushJobV2(updateLogpushJobV2OptionsModelNew)
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
			It(`Invoke UpdateLogpushJobV2 successfully`, func() {
				logpushJobsApiService, serviceErr := logpushjobsapiv1.NewLogpushJobsApiV1(&logpushjobsapiv1.LogpushJobsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					Dataset:       core.StringPtr(dataset),
					ZoneID:        core.StringPtr(zoneID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(logpushJobsApiService).ToNot(BeNil())

				// Construct an instance of the UpdateLogpushJobV2RequestLogpushJobsUpdateCosReq model
				updateLogpushJobV2RequestModel := new(logpushjobsapiv1.UpdateLogpushJobV2RequestLogpushJobsUpdateCosReq)
				updateLogpushJobV2RequestModel.Enabled = core.BoolPtr(false)
				updateLogpushJobV2RequestModel.LogpullOptions = core.StringPtr("timestamps=rfc3339&timestamps=rfc3339")
				updateLogpushJobV2RequestModel.Cos = map[string]interface{}{"anyKey": "anyValue"}
				updateLogpushJobV2RequestModel.OwnershipChallenge = core.StringPtr("00000000000000000000000000000000")
				updateLogpushJobV2RequestModel.Frequency = core.StringPtr("high")

				// Construct an instance of the UpdateLogpushJobV2Options model
				updateLogpushJobV2OptionsModel := new(logpushjobsapiv1.UpdateLogpushJobV2Options)
				updateLogpushJobV2OptionsModel.JobID = core.StringPtr("testString")
				updateLogpushJobV2OptionsModel.UpdateLogpushJobV2Request = updateLogpushJobV2RequestModel
				updateLogpushJobV2OptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := logpushJobsApiService.UpdateLogpushJobV2(updateLogpushJobV2OptionsModel)
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
	Describe(`DeleteLogpushJobV2(deleteLogpushJobV2Options *DeleteLogpushJobV2Options) - Operation response error`, func() {
		crn := "testString"
		dataset := "testString"
		zoneID := "testString"
		deleteLogpushJobV2Path := "/v2/testString/zones/testString/logpush/jobs/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteLogpushJobV2Path))
					Expect(req.Method).To(Equal("DELETE"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke DeleteLogpushJobV2 with error: Operation response processing error`, func() {
				logpushJobsApiService, serviceErr := logpushjobsapiv1.NewLogpushJobsApiV1(&logpushjobsapiv1.LogpushJobsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					Dataset:       core.StringPtr(dataset),
					ZoneID:        core.StringPtr(zoneID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(logpushJobsApiService).ToNot(BeNil())

				// Construct an instance of the DeleteLogpushJobV2Options model
				deleteLogpushJobV2OptionsModel := new(logpushjobsapiv1.DeleteLogpushJobV2Options)
				deleteLogpushJobV2OptionsModel.JobID = core.StringPtr("testString")
				deleteLogpushJobV2OptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := logpushJobsApiService.DeleteLogpushJobV2(deleteLogpushJobV2OptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				logpushJobsApiService.EnableRetries(0, 0)
				result, response, operationErr = logpushJobsApiService.DeleteLogpushJobV2(deleteLogpushJobV2OptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteLogpushJobV2(deleteLogpushJobV2Options *DeleteLogpushJobV2Options)`, func() {
		crn := "testString"
		dataset := "testString"
		zoneID := "testString"
		deleteLogpushJobV2Path := "/v2/testString/zones/testString/logpush/jobs/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteLogpushJobV2Path))
					Expect(req.Method).To(Equal("DELETE"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"anyKey": "anyValue"}}`)
				}))
			})
			It(`Invoke DeleteLogpushJobV2 successfully with retries`, func() {
				logpushJobsApiService, serviceErr := logpushjobsapiv1.NewLogpushJobsApiV1(&logpushjobsapiv1.LogpushJobsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					Dataset:       core.StringPtr(dataset),
					ZoneID:        core.StringPtr(zoneID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(logpushJobsApiService).ToNot(BeNil())
				logpushJobsApiService.EnableRetries(0, 0)

				// Construct an instance of the DeleteLogpushJobV2Options model
				deleteLogpushJobV2OptionsModel := new(logpushjobsapiv1.DeleteLogpushJobV2Options)
				deleteLogpushJobV2OptionsModel.JobID = core.StringPtr("testString")
				deleteLogpushJobV2OptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := logpushJobsApiService.DeleteLogpushJobV2WithContext(ctx, deleteLogpushJobV2OptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				logpushJobsApiService.DisableRetries()
				result, response, operationErr := logpushJobsApiService.DeleteLogpushJobV2(deleteLogpushJobV2OptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = logpushJobsApiService.DeleteLogpushJobV2WithContext(ctx, deleteLogpushJobV2OptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(deleteLogpushJobV2Path))
					Expect(req.Method).To(Equal("DELETE"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"anyKey": "anyValue"}}`)
				}))
			})
			It(`Invoke DeleteLogpushJobV2 successfully`, func() {
				logpushJobsApiService, serviceErr := logpushjobsapiv1.NewLogpushJobsApiV1(&logpushjobsapiv1.LogpushJobsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					Dataset:       core.StringPtr(dataset),
					ZoneID:        core.StringPtr(zoneID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(logpushJobsApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := logpushJobsApiService.DeleteLogpushJobV2(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DeleteLogpushJobV2Options model
				deleteLogpushJobV2OptionsModel := new(logpushjobsapiv1.DeleteLogpushJobV2Options)
				deleteLogpushJobV2OptionsModel.JobID = core.StringPtr("testString")
				deleteLogpushJobV2OptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = logpushJobsApiService.DeleteLogpushJobV2(deleteLogpushJobV2OptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke DeleteLogpushJobV2 with error: Operation validation and request error`, func() {
				logpushJobsApiService, serviceErr := logpushjobsapiv1.NewLogpushJobsApiV1(&logpushjobsapiv1.LogpushJobsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					Dataset:       core.StringPtr(dataset),
					ZoneID:        core.StringPtr(zoneID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(logpushJobsApiService).ToNot(BeNil())

				// Construct an instance of the DeleteLogpushJobV2Options model
				deleteLogpushJobV2OptionsModel := new(logpushjobsapiv1.DeleteLogpushJobV2Options)
				deleteLogpushJobV2OptionsModel.JobID = core.StringPtr("testString")
				deleteLogpushJobV2OptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := logpushJobsApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := logpushJobsApiService.DeleteLogpushJobV2(deleteLogpushJobV2OptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the DeleteLogpushJobV2Options model with no property values
				deleteLogpushJobV2OptionsModelNew := new(logpushjobsapiv1.DeleteLogpushJobV2Options)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = logpushJobsApiService.DeleteLogpushJobV2(deleteLogpushJobV2OptionsModelNew)
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
			It(`Invoke DeleteLogpushJobV2 successfully`, func() {
				logpushJobsApiService, serviceErr := logpushjobsapiv1.NewLogpushJobsApiV1(&logpushjobsapiv1.LogpushJobsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					Dataset:       core.StringPtr(dataset),
					ZoneID:        core.StringPtr(zoneID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(logpushJobsApiService).ToNot(BeNil())

				// Construct an instance of the DeleteLogpushJobV2Options model
				deleteLogpushJobV2OptionsModel := new(logpushjobsapiv1.DeleteLogpushJobV2Options)
				deleteLogpushJobV2OptionsModel.JobID = core.StringPtr("testString")
				deleteLogpushJobV2OptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := logpushJobsApiService.DeleteLogpushJobV2(deleteLogpushJobV2OptionsModel)
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
	Describe(`GetLogpushOwnershipV2(getLogpushOwnershipV2Options *GetLogpushOwnershipV2Options) - Operation response error`, func() {
		crn := "testString"
		dataset := "testString"
		zoneID := "testString"
		getLogpushOwnershipV2Path := "/v2/testString/zones/testString/logpush/ownership"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getLogpushOwnershipV2Path))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetLogpushOwnershipV2 with error: Operation response processing error`, func() {
				logpushJobsApiService, serviceErr := logpushjobsapiv1.NewLogpushJobsApiV1(&logpushjobsapiv1.LogpushJobsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					Dataset:       core.StringPtr(dataset),
					ZoneID:        core.StringPtr(zoneID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(logpushJobsApiService).ToNot(BeNil())

				// Construct an instance of the GetLogpushOwnershipV2Options model
				getLogpushOwnershipV2OptionsModel := new(logpushjobsapiv1.GetLogpushOwnershipV2Options)
				getLogpushOwnershipV2OptionsModel.Cos = map[string]interface{}{"anyKey": "anyValue"}
				getLogpushOwnershipV2OptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := logpushJobsApiService.GetLogpushOwnershipV2(getLogpushOwnershipV2OptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				logpushJobsApiService.EnableRetries(0, 0)
				result, response, operationErr = logpushJobsApiService.GetLogpushOwnershipV2(getLogpushOwnershipV2OptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetLogpushOwnershipV2(getLogpushOwnershipV2Options *GetLogpushOwnershipV2Options)`, func() {
		crn := "testString"
		dataset := "testString"
		zoneID := "testString"
		getLogpushOwnershipV2Path := "/v2/testString/zones/testString/logpush/ownership"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getLogpushOwnershipV2Path))
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
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"filename": "logs/challenge-filename.txt", "valid": true, "messages": "Messages"}}`)
				}))
			})
			It(`Invoke GetLogpushOwnershipV2 successfully with retries`, func() {
				logpushJobsApiService, serviceErr := logpushjobsapiv1.NewLogpushJobsApiV1(&logpushjobsapiv1.LogpushJobsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					Dataset:       core.StringPtr(dataset),
					ZoneID:        core.StringPtr(zoneID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(logpushJobsApiService).ToNot(BeNil())
				logpushJobsApiService.EnableRetries(0, 0)

				// Construct an instance of the GetLogpushOwnershipV2Options model
				getLogpushOwnershipV2OptionsModel := new(logpushjobsapiv1.GetLogpushOwnershipV2Options)
				getLogpushOwnershipV2OptionsModel.Cos = map[string]interface{}{"anyKey": "anyValue"}
				getLogpushOwnershipV2OptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := logpushJobsApiService.GetLogpushOwnershipV2WithContext(ctx, getLogpushOwnershipV2OptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				logpushJobsApiService.DisableRetries()
				result, response, operationErr := logpushJobsApiService.GetLogpushOwnershipV2(getLogpushOwnershipV2OptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = logpushJobsApiService.GetLogpushOwnershipV2WithContext(ctx, getLogpushOwnershipV2OptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getLogpushOwnershipV2Path))
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
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"filename": "logs/challenge-filename.txt", "valid": true, "messages": "Messages"}}`)
				}))
			})
			It(`Invoke GetLogpushOwnershipV2 successfully`, func() {
				logpushJobsApiService, serviceErr := logpushjobsapiv1.NewLogpushJobsApiV1(&logpushjobsapiv1.LogpushJobsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					Dataset:       core.StringPtr(dataset),
					ZoneID:        core.StringPtr(zoneID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(logpushJobsApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := logpushJobsApiService.GetLogpushOwnershipV2(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetLogpushOwnershipV2Options model
				getLogpushOwnershipV2OptionsModel := new(logpushjobsapiv1.GetLogpushOwnershipV2Options)
				getLogpushOwnershipV2OptionsModel.Cos = map[string]interface{}{"anyKey": "anyValue"}
				getLogpushOwnershipV2OptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = logpushJobsApiService.GetLogpushOwnershipV2(getLogpushOwnershipV2OptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetLogpushOwnershipV2 with error: Operation request error`, func() {
				logpushJobsApiService, serviceErr := logpushjobsapiv1.NewLogpushJobsApiV1(&logpushjobsapiv1.LogpushJobsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					Dataset:       core.StringPtr(dataset),
					ZoneID:        core.StringPtr(zoneID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(logpushJobsApiService).ToNot(BeNil())

				// Construct an instance of the GetLogpushOwnershipV2Options model
				getLogpushOwnershipV2OptionsModel := new(logpushjobsapiv1.GetLogpushOwnershipV2Options)
				getLogpushOwnershipV2OptionsModel.Cos = map[string]interface{}{"anyKey": "anyValue"}
				getLogpushOwnershipV2OptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := logpushJobsApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := logpushJobsApiService.GetLogpushOwnershipV2(getLogpushOwnershipV2OptionsModel)
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
			It(`Invoke GetLogpushOwnershipV2 successfully`, func() {
				logpushJobsApiService, serviceErr := logpushjobsapiv1.NewLogpushJobsApiV1(&logpushjobsapiv1.LogpushJobsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					Dataset:       core.StringPtr(dataset),
					ZoneID:        core.StringPtr(zoneID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(logpushJobsApiService).ToNot(BeNil())

				// Construct an instance of the GetLogpushOwnershipV2Options model
				getLogpushOwnershipV2OptionsModel := new(logpushjobsapiv1.GetLogpushOwnershipV2Options)
				getLogpushOwnershipV2OptionsModel.Cos = map[string]interface{}{"anyKey": "anyValue"}
				getLogpushOwnershipV2OptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := logpushJobsApiService.GetLogpushOwnershipV2(getLogpushOwnershipV2OptionsModel)
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
	Describe(`ValidateLogpushOwnershipChallengeV2(validateLogpushOwnershipChallengeV2Options *ValidateLogpushOwnershipChallengeV2Options) - Operation response error`, func() {
		crn := "testString"
		dataset := "testString"
		zoneID := "testString"
		validateLogpushOwnershipChallengeV2Path := "/v2/testString/zones/testString/logpush/ownership/validate"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(validateLogpushOwnershipChallengeV2Path))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ValidateLogpushOwnershipChallengeV2 with error: Operation response processing error`, func() {
				logpushJobsApiService, serviceErr := logpushjobsapiv1.NewLogpushJobsApiV1(&logpushjobsapiv1.LogpushJobsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					Dataset:       core.StringPtr(dataset),
					ZoneID:        core.StringPtr(zoneID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(logpushJobsApiService).ToNot(BeNil())

				// Construct an instance of the ValidateLogpushOwnershipChallengeV2Options model
				validateLogpushOwnershipChallengeV2OptionsModel := new(logpushjobsapiv1.ValidateLogpushOwnershipChallengeV2Options)
				validateLogpushOwnershipChallengeV2OptionsModel.Cos = map[string]interface{}{"anyKey": "anyValue"}
				validateLogpushOwnershipChallengeV2OptionsModel.OwnershipChallenge = core.StringPtr("00000000000000000000")
				validateLogpushOwnershipChallengeV2OptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := logpushJobsApiService.ValidateLogpushOwnershipChallengeV2(validateLogpushOwnershipChallengeV2OptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				logpushJobsApiService.EnableRetries(0, 0)
				result, response, operationErr = logpushJobsApiService.ValidateLogpushOwnershipChallengeV2(validateLogpushOwnershipChallengeV2OptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ValidateLogpushOwnershipChallengeV2(validateLogpushOwnershipChallengeV2Options *ValidateLogpushOwnershipChallengeV2Options)`, func() {
		crn := "testString"
		dataset := "testString"
		zoneID := "testString"
		validateLogpushOwnershipChallengeV2Path := "/v2/testString/zones/testString/logpush/ownership/validate"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(validateLogpushOwnershipChallengeV2Path))
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
					fmt.Fprintf(res, "%s", `{"valid": true}`)
				}))
			})
			It(`Invoke ValidateLogpushOwnershipChallengeV2 successfully with retries`, func() {
				logpushJobsApiService, serviceErr := logpushjobsapiv1.NewLogpushJobsApiV1(&logpushjobsapiv1.LogpushJobsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					Dataset:       core.StringPtr(dataset),
					ZoneID:        core.StringPtr(zoneID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(logpushJobsApiService).ToNot(BeNil())
				logpushJobsApiService.EnableRetries(0, 0)

				// Construct an instance of the ValidateLogpushOwnershipChallengeV2Options model
				validateLogpushOwnershipChallengeV2OptionsModel := new(logpushjobsapiv1.ValidateLogpushOwnershipChallengeV2Options)
				validateLogpushOwnershipChallengeV2OptionsModel.Cos = map[string]interface{}{"anyKey": "anyValue"}
				validateLogpushOwnershipChallengeV2OptionsModel.OwnershipChallenge = core.StringPtr("00000000000000000000")
				validateLogpushOwnershipChallengeV2OptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := logpushJobsApiService.ValidateLogpushOwnershipChallengeV2WithContext(ctx, validateLogpushOwnershipChallengeV2OptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				logpushJobsApiService.DisableRetries()
				result, response, operationErr := logpushJobsApiService.ValidateLogpushOwnershipChallengeV2(validateLogpushOwnershipChallengeV2OptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = logpushJobsApiService.ValidateLogpushOwnershipChallengeV2WithContext(ctx, validateLogpushOwnershipChallengeV2OptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(validateLogpushOwnershipChallengeV2Path))
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
					fmt.Fprintf(res, "%s", `{"valid": true}`)
				}))
			})
			It(`Invoke ValidateLogpushOwnershipChallengeV2 successfully`, func() {
				logpushJobsApiService, serviceErr := logpushjobsapiv1.NewLogpushJobsApiV1(&logpushjobsapiv1.LogpushJobsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					Dataset:       core.StringPtr(dataset),
					ZoneID:        core.StringPtr(zoneID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(logpushJobsApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := logpushJobsApiService.ValidateLogpushOwnershipChallengeV2(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ValidateLogpushOwnershipChallengeV2Options model
				validateLogpushOwnershipChallengeV2OptionsModel := new(logpushjobsapiv1.ValidateLogpushOwnershipChallengeV2Options)
				validateLogpushOwnershipChallengeV2OptionsModel.Cos = map[string]interface{}{"anyKey": "anyValue"}
				validateLogpushOwnershipChallengeV2OptionsModel.OwnershipChallenge = core.StringPtr("00000000000000000000")
				validateLogpushOwnershipChallengeV2OptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = logpushJobsApiService.ValidateLogpushOwnershipChallengeV2(validateLogpushOwnershipChallengeV2OptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ValidateLogpushOwnershipChallengeV2 with error: Operation request error`, func() {
				logpushJobsApiService, serviceErr := logpushjobsapiv1.NewLogpushJobsApiV1(&logpushjobsapiv1.LogpushJobsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					Dataset:       core.StringPtr(dataset),
					ZoneID:        core.StringPtr(zoneID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(logpushJobsApiService).ToNot(BeNil())

				// Construct an instance of the ValidateLogpushOwnershipChallengeV2Options model
				validateLogpushOwnershipChallengeV2OptionsModel := new(logpushjobsapiv1.ValidateLogpushOwnershipChallengeV2Options)
				validateLogpushOwnershipChallengeV2OptionsModel.Cos = map[string]interface{}{"anyKey": "anyValue"}
				validateLogpushOwnershipChallengeV2OptionsModel.OwnershipChallenge = core.StringPtr("00000000000000000000")
				validateLogpushOwnershipChallengeV2OptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := logpushJobsApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := logpushJobsApiService.ValidateLogpushOwnershipChallengeV2(validateLogpushOwnershipChallengeV2OptionsModel)
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
			It(`Invoke ValidateLogpushOwnershipChallengeV2 successfully`, func() {
				logpushJobsApiService, serviceErr := logpushjobsapiv1.NewLogpushJobsApiV1(&logpushjobsapiv1.LogpushJobsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					Dataset:       core.StringPtr(dataset),
					ZoneID:        core.StringPtr(zoneID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(logpushJobsApiService).ToNot(BeNil())

				// Construct an instance of the ValidateLogpushOwnershipChallengeV2Options model
				validateLogpushOwnershipChallengeV2OptionsModel := new(logpushjobsapiv1.ValidateLogpushOwnershipChallengeV2Options)
				validateLogpushOwnershipChallengeV2OptionsModel.Cos = map[string]interface{}{"anyKey": "anyValue"}
				validateLogpushOwnershipChallengeV2OptionsModel.OwnershipChallenge = core.StringPtr("00000000000000000000")
				validateLogpushOwnershipChallengeV2OptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := logpushJobsApiService.ValidateLogpushOwnershipChallengeV2(validateLogpushOwnershipChallengeV2OptionsModel)
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
	Describe(`ListFieldsForDatasetV2(listFieldsForDatasetV2Options *ListFieldsForDatasetV2Options) - Operation response error`, func() {
		crn := "testString"
		dataset := "testString"
		zoneID := "testString"
		listFieldsForDatasetV2Path := "/v2/testString/zones/testString/logpush/datasets/testString/fields"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listFieldsForDatasetV2Path))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListFieldsForDatasetV2 with error: Operation response processing error`, func() {
				logpushJobsApiService, serviceErr := logpushjobsapiv1.NewLogpushJobsApiV1(&logpushjobsapiv1.LogpushJobsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					Dataset:       core.StringPtr(dataset),
					ZoneID:        core.StringPtr(zoneID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(logpushJobsApiService).ToNot(BeNil())

				// Construct an instance of the ListFieldsForDatasetV2Options model
				listFieldsForDatasetV2OptionsModel := new(logpushjobsapiv1.ListFieldsForDatasetV2Options)
				listFieldsForDatasetV2OptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := logpushJobsApiService.ListFieldsForDatasetV2(listFieldsForDatasetV2OptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				logpushJobsApiService.EnableRetries(0, 0)
				result, response, operationErr = logpushJobsApiService.ListFieldsForDatasetV2(listFieldsForDatasetV2OptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListFieldsForDatasetV2(listFieldsForDatasetV2Options *ListFieldsForDatasetV2Options)`, func() {
		crn := "testString"
		dataset := "testString"
		zoneID := "testString"
		listFieldsForDatasetV2Path := "/v2/testString/zones/testString/logpush/datasets/testString/fields"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listFieldsForDatasetV2Path))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"anyKey": "anyValue"}}`)
				}))
			})
			It(`Invoke ListFieldsForDatasetV2 successfully with retries`, func() {
				logpushJobsApiService, serviceErr := logpushjobsapiv1.NewLogpushJobsApiV1(&logpushjobsapiv1.LogpushJobsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					Dataset:       core.StringPtr(dataset),
					ZoneID:        core.StringPtr(zoneID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(logpushJobsApiService).ToNot(BeNil())
				logpushJobsApiService.EnableRetries(0, 0)

				// Construct an instance of the ListFieldsForDatasetV2Options model
				listFieldsForDatasetV2OptionsModel := new(logpushjobsapiv1.ListFieldsForDatasetV2Options)
				listFieldsForDatasetV2OptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := logpushJobsApiService.ListFieldsForDatasetV2WithContext(ctx, listFieldsForDatasetV2OptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				logpushJobsApiService.DisableRetries()
				result, response, operationErr := logpushJobsApiService.ListFieldsForDatasetV2(listFieldsForDatasetV2OptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = logpushJobsApiService.ListFieldsForDatasetV2WithContext(ctx, listFieldsForDatasetV2OptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listFieldsForDatasetV2Path))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"anyKey": "anyValue"}}`)
				}))
			})
			It(`Invoke ListFieldsForDatasetV2 successfully`, func() {
				logpushJobsApiService, serviceErr := logpushjobsapiv1.NewLogpushJobsApiV1(&logpushjobsapiv1.LogpushJobsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					Dataset:       core.StringPtr(dataset),
					ZoneID:        core.StringPtr(zoneID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(logpushJobsApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := logpushJobsApiService.ListFieldsForDatasetV2(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListFieldsForDatasetV2Options model
				listFieldsForDatasetV2OptionsModel := new(logpushjobsapiv1.ListFieldsForDatasetV2Options)
				listFieldsForDatasetV2OptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = logpushJobsApiService.ListFieldsForDatasetV2(listFieldsForDatasetV2OptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListFieldsForDatasetV2 with error: Operation request error`, func() {
				logpushJobsApiService, serviceErr := logpushjobsapiv1.NewLogpushJobsApiV1(&logpushjobsapiv1.LogpushJobsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					Dataset:       core.StringPtr(dataset),
					ZoneID:        core.StringPtr(zoneID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(logpushJobsApiService).ToNot(BeNil())

				// Construct an instance of the ListFieldsForDatasetV2Options model
				listFieldsForDatasetV2OptionsModel := new(logpushjobsapiv1.ListFieldsForDatasetV2Options)
				listFieldsForDatasetV2OptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := logpushJobsApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := logpushJobsApiService.ListFieldsForDatasetV2(listFieldsForDatasetV2OptionsModel)
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
			It(`Invoke ListFieldsForDatasetV2 successfully`, func() {
				logpushJobsApiService, serviceErr := logpushjobsapiv1.NewLogpushJobsApiV1(&logpushjobsapiv1.LogpushJobsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					Dataset:       core.StringPtr(dataset),
					ZoneID:        core.StringPtr(zoneID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(logpushJobsApiService).ToNot(BeNil())

				// Construct an instance of the ListFieldsForDatasetV2Options model
				listFieldsForDatasetV2OptionsModel := new(logpushjobsapiv1.ListFieldsForDatasetV2Options)
				listFieldsForDatasetV2OptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := logpushJobsApiService.ListFieldsForDatasetV2(listFieldsForDatasetV2OptionsModel)
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
	Describe(`ListLogpushJobsForDatasetV2(listLogpushJobsForDatasetV2Options *ListLogpushJobsForDatasetV2Options) - Operation response error`, func() {
		crn := "testString"
		dataset := "testString"
		zoneID := "testString"
		listLogpushJobsForDatasetV2Path := "/v2/testString/zones/testString/logpush/datasets/testString/jobs"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listLogpushJobsForDatasetV2Path))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListLogpushJobsForDatasetV2 with error: Operation response processing error`, func() {
				logpushJobsApiService, serviceErr := logpushjobsapiv1.NewLogpushJobsApiV1(&logpushjobsapiv1.LogpushJobsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					Dataset:       core.StringPtr(dataset),
					ZoneID:        core.StringPtr(zoneID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(logpushJobsApiService).ToNot(BeNil())

				// Construct an instance of the ListLogpushJobsForDatasetV2Options model
				listLogpushJobsForDatasetV2OptionsModel := new(logpushjobsapiv1.ListLogpushJobsForDatasetV2Options)
				listLogpushJobsForDatasetV2OptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := logpushJobsApiService.ListLogpushJobsForDatasetV2(listLogpushJobsForDatasetV2OptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				logpushJobsApiService.EnableRetries(0, 0)
				result, response, operationErr = logpushJobsApiService.ListLogpushJobsForDatasetV2(listLogpushJobsForDatasetV2OptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListLogpushJobsForDatasetV2(listLogpushJobsForDatasetV2Options *ListLogpushJobsForDatasetV2Options)`, func() {
		crn := "testString"
		dataset := "testString"
		zoneID := "testString"
		listLogpushJobsForDatasetV2Path := "/v2/testString/zones/testString/logpush/datasets/testString/jobs"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listLogpushJobsForDatasetV2Path))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": 5850, "name": "My log push job", "enabled": false, "dataset": "firewall_events", "frequency": "high", "logpull_options": "timestamps=rfc3339&timestamps=rfc3339", "destination_conf": "cos://cos-bucket001?region=us-south&instance-id=231f5467-3072-4cb9-9e39-a906fa3032ea", "last_complete": "2022-01-15T16:33:31.834209Z", "last_error": "2022-01-15T16:33:31.834209Z", "error_message": "ErrorMessage"}}`)
				}))
			})
			It(`Invoke ListLogpushJobsForDatasetV2 successfully with retries`, func() {
				logpushJobsApiService, serviceErr := logpushjobsapiv1.NewLogpushJobsApiV1(&logpushjobsapiv1.LogpushJobsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					Dataset:       core.StringPtr(dataset),
					ZoneID:        core.StringPtr(zoneID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(logpushJobsApiService).ToNot(BeNil())
				logpushJobsApiService.EnableRetries(0, 0)

				// Construct an instance of the ListLogpushJobsForDatasetV2Options model
				listLogpushJobsForDatasetV2OptionsModel := new(logpushjobsapiv1.ListLogpushJobsForDatasetV2Options)
				listLogpushJobsForDatasetV2OptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := logpushJobsApiService.ListLogpushJobsForDatasetV2WithContext(ctx, listLogpushJobsForDatasetV2OptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				logpushJobsApiService.DisableRetries()
				result, response, operationErr := logpushJobsApiService.ListLogpushJobsForDatasetV2(listLogpushJobsForDatasetV2OptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = logpushJobsApiService.ListLogpushJobsForDatasetV2WithContext(ctx, listLogpushJobsForDatasetV2OptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listLogpushJobsForDatasetV2Path))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": 5850, "name": "My log push job", "enabled": false, "dataset": "firewall_events", "frequency": "high", "logpull_options": "timestamps=rfc3339&timestamps=rfc3339", "destination_conf": "cos://cos-bucket001?region=us-south&instance-id=231f5467-3072-4cb9-9e39-a906fa3032ea", "last_complete": "2022-01-15T16:33:31.834209Z", "last_error": "2022-01-15T16:33:31.834209Z", "error_message": "ErrorMessage"}}`)
				}))
			})
			It(`Invoke ListLogpushJobsForDatasetV2 successfully`, func() {
				logpushJobsApiService, serviceErr := logpushjobsapiv1.NewLogpushJobsApiV1(&logpushjobsapiv1.LogpushJobsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					Dataset:       core.StringPtr(dataset),
					ZoneID:        core.StringPtr(zoneID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(logpushJobsApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := logpushJobsApiService.ListLogpushJobsForDatasetV2(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListLogpushJobsForDatasetV2Options model
				listLogpushJobsForDatasetV2OptionsModel := new(logpushjobsapiv1.ListLogpushJobsForDatasetV2Options)
				listLogpushJobsForDatasetV2OptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = logpushJobsApiService.ListLogpushJobsForDatasetV2(listLogpushJobsForDatasetV2OptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListLogpushJobsForDatasetV2 with error: Operation request error`, func() {
				logpushJobsApiService, serviceErr := logpushjobsapiv1.NewLogpushJobsApiV1(&logpushjobsapiv1.LogpushJobsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					Dataset:       core.StringPtr(dataset),
					ZoneID:        core.StringPtr(zoneID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(logpushJobsApiService).ToNot(BeNil())

				// Construct an instance of the ListLogpushJobsForDatasetV2Options model
				listLogpushJobsForDatasetV2OptionsModel := new(logpushjobsapiv1.ListLogpushJobsForDatasetV2Options)
				listLogpushJobsForDatasetV2OptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := logpushJobsApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := logpushJobsApiService.ListLogpushJobsForDatasetV2(listLogpushJobsForDatasetV2OptionsModel)
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
			It(`Invoke ListLogpushJobsForDatasetV2 successfully`, func() {
				logpushJobsApiService, serviceErr := logpushjobsapiv1.NewLogpushJobsApiV1(&logpushjobsapiv1.LogpushJobsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					Dataset:       core.StringPtr(dataset),
					ZoneID:        core.StringPtr(zoneID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(logpushJobsApiService).ToNot(BeNil())

				// Construct an instance of the ListLogpushJobsForDatasetV2Options model
				listLogpushJobsForDatasetV2OptionsModel := new(logpushjobsapiv1.ListLogpushJobsForDatasetV2Options)
				listLogpushJobsForDatasetV2OptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := logpushJobsApiService.ListLogpushJobsForDatasetV2(listLogpushJobsForDatasetV2OptionsModel)
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
	Describe(`GetLogsRetention(getLogsRetentionOptions *GetLogsRetentionOptions) - Operation response error`, func() {
		crn := "testString"
		dataset := "testString"
		zoneID := "testString"
		getLogsRetentionPath := "/v1/testString/zones/testString/logs/retention"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getLogsRetentionPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetLogsRetention with error: Operation response processing error`, func() {
				logpushJobsApiService, serviceErr := logpushjobsapiv1.NewLogpushJobsApiV1(&logpushjobsapiv1.LogpushJobsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					Dataset:       core.StringPtr(dataset),
					ZoneID:        core.StringPtr(zoneID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(logpushJobsApiService).ToNot(BeNil())

				// Construct an instance of the GetLogsRetentionOptions model
				getLogsRetentionOptionsModel := new(logpushjobsapiv1.GetLogsRetentionOptions)
				getLogsRetentionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := logpushJobsApiService.GetLogsRetention(getLogsRetentionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				logpushJobsApiService.EnableRetries(0, 0)
				result, response, operationErr = logpushJobsApiService.GetLogsRetention(getLogsRetentionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetLogsRetention(getLogsRetentionOptions *GetLogsRetentionOptions)`, func() {
		crn := "testString"
		dataset := "testString"
		zoneID := "testString"
		getLogsRetentionPath := "/v1/testString/zones/testString/logs/retention"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getLogsRetentionPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"result": {"flag": true}, "success": true, "errors": [["Errors"]], "messages": [["Messages"]]}`)
				}))
			})
			It(`Invoke GetLogsRetention successfully with retries`, func() {
				logpushJobsApiService, serviceErr := logpushjobsapiv1.NewLogpushJobsApiV1(&logpushjobsapiv1.LogpushJobsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					Dataset:       core.StringPtr(dataset),
					ZoneID:        core.StringPtr(zoneID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(logpushJobsApiService).ToNot(BeNil())
				logpushJobsApiService.EnableRetries(0, 0)

				// Construct an instance of the GetLogsRetentionOptions model
				getLogsRetentionOptionsModel := new(logpushjobsapiv1.GetLogsRetentionOptions)
				getLogsRetentionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := logpushJobsApiService.GetLogsRetentionWithContext(ctx, getLogsRetentionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				logpushJobsApiService.DisableRetries()
				result, response, operationErr := logpushJobsApiService.GetLogsRetention(getLogsRetentionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = logpushJobsApiService.GetLogsRetentionWithContext(ctx, getLogsRetentionOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getLogsRetentionPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"result": {"flag": true}, "success": true, "errors": [["Errors"]], "messages": [["Messages"]]}`)
				}))
			})
			It(`Invoke GetLogsRetention successfully`, func() {
				logpushJobsApiService, serviceErr := logpushjobsapiv1.NewLogpushJobsApiV1(&logpushjobsapiv1.LogpushJobsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					Dataset:       core.StringPtr(dataset),
					ZoneID:        core.StringPtr(zoneID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(logpushJobsApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := logpushJobsApiService.GetLogsRetention(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetLogsRetentionOptions model
				getLogsRetentionOptionsModel := new(logpushjobsapiv1.GetLogsRetentionOptions)
				getLogsRetentionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = logpushJobsApiService.GetLogsRetention(getLogsRetentionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetLogsRetention with error: Operation request error`, func() {
				logpushJobsApiService, serviceErr := logpushjobsapiv1.NewLogpushJobsApiV1(&logpushjobsapiv1.LogpushJobsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					Dataset:       core.StringPtr(dataset),
					ZoneID:        core.StringPtr(zoneID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(logpushJobsApiService).ToNot(BeNil())

				// Construct an instance of the GetLogsRetentionOptions model
				getLogsRetentionOptionsModel := new(logpushjobsapiv1.GetLogsRetentionOptions)
				getLogsRetentionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := logpushJobsApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := logpushJobsApiService.GetLogsRetention(getLogsRetentionOptionsModel)
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
			It(`Invoke GetLogsRetention successfully`, func() {
				logpushJobsApiService, serviceErr := logpushjobsapiv1.NewLogpushJobsApiV1(&logpushjobsapiv1.LogpushJobsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					Dataset:       core.StringPtr(dataset),
					ZoneID:        core.StringPtr(zoneID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(logpushJobsApiService).ToNot(BeNil())

				// Construct an instance of the GetLogsRetentionOptions model
				getLogsRetentionOptionsModel := new(logpushjobsapiv1.GetLogsRetentionOptions)
				getLogsRetentionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := logpushJobsApiService.GetLogsRetention(getLogsRetentionOptionsModel)
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
	Describe(`CreateLogRetention(createLogRetentionOptions *CreateLogRetentionOptions) - Operation response error`, func() {
		crn := "testString"
		dataset := "testString"
		zoneID := "testString"
		createLogRetentionPath := "/v1/testString/zones/testString/logs/retention"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createLogRetentionPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateLogRetention with error: Operation response processing error`, func() {
				logpushJobsApiService, serviceErr := logpushjobsapiv1.NewLogpushJobsApiV1(&logpushjobsapiv1.LogpushJobsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					Dataset:       core.StringPtr(dataset),
					ZoneID:        core.StringPtr(zoneID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(logpushJobsApiService).ToNot(BeNil())

				// Construct an instance of the CreateLogRetentionOptions model
				createLogRetentionOptionsModel := new(logpushjobsapiv1.CreateLogRetentionOptions)
				createLogRetentionOptionsModel.Flag = core.BoolPtr(false)
				createLogRetentionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := logpushJobsApiService.CreateLogRetention(createLogRetentionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				logpushJobsApiService.EnableRetries(0, 0)
				result, response, operationErr = logpushJobsApiService.CreateLogRetention(createLogRetentionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateLogRetention(createLogRetentionOptions *CreateLogRetentionOptions)`, func() {
		crn := "testString"
		dataset := "testString"
		zoneID := "testString"
		createLogRetentionPath := "/v1/testString/zones/testString/logs/retention"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createLogRetentionPath))
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
					fmt.Fprintf(res, "%s", `{"result": {"flag": true}, "success": true, "errors": [["Errors"]], "messages": [["Messages"]]}`)
				}))
			})
			It(`Invoke CreateLogRetention successfully with retries`, func() {
				logpushJobsApiService, serviceErr := logpushjobsapiv1.NewLogpushJobsApiV1(&logpushjobsapiv1.LogpushJobsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					Dataset:       core.StringPtr(dataset),
					ZoneID:        core.StringPtr(zoneID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(logpushJobsApiService).ToNot(BeNil())
				logpushJobsApiService.EnableRetries(0, 0)

				// Construct an instance of the CreateLogRetentionOptions model
				createLogRetentionOptionsModel := new(logpushjobsapiv1.CreateLogRetentionOptions)
				createLogRetentionOptionsModel.Flag = core.BoolPtr(false)
				createLogRetentionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := logpushJobsApiService.CreateLogRetentionWithContext(ctx, createLogRetentionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				logpushJobsApiService.DisableRetries()
				result, response, operationErr := logpushJobsApiService.CreateLogRetention(createLogRetentionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = logpushJobsApiService.CreateLogRetentionWithContext(ctx, createLogRetentionOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createLogRetentionPath))
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
					fmt.Fprintf(res, "%s", `{"result": {"flag": true}, "success": true, "errors": [["Errors"]], "messages": [["Messages"]]}`)
				}))
			})
			It(`Invoke CreateLogRetention successfully`, func() {
				logpushJobsApiService, serviceErr := logpushjobsapiv1.NewLogpushJobsApiV1(&logpushjobsapiv1.LogpushJobsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					Dataset:       core.StringPtr(dataset),
					ZoneID:        core.StringPtr(zoneID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(logpushJobsApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := logpushJobsApiService.CreateLogRetention(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateLogRetentionOptions model
				createLogRetentionOptionsModel := new(logpushjobsapiv1.CreateLogRetentionOptions)
				createLogRetentionOptionsModel.Flag = core.BoolPtr(false)
				createLogRetentionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = logpushJobsApiService.CreateLogRetention(createLogRetentionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateLogRetention with error: Operation request error`, func() {
				logpushJobsApiService, serviceErr := logpushjobsapiv1.NewLogpushJobsApiV1(&logpushjobsapiv1.LogpushJobsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					Dataset:       core.StringPtr(dataset),
					ZoneID:        core.StringPtr(zoneID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(logpushJobsApiService).ToNot(BeNil())

				// Construct an instance of the CreateLogRetentionOptions model
				createLogRetentionOptionsModel := new(logpushjobsapiv1.CreateLogRetentionOptions)
				createLogRetentionOptionsModel.Flag = core.BoolPtr(false)
				createLogRetentionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := logpushJobsApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := logpushJobsApiService.CreateLogRetention(createLogRetentionOptionsModel)
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
			It(`Invoke CreateLogRetention successfully`, func() {
				logpushJobsApiService, serviceErr := logpushjobsapiv1.NewLogpushJobsApiV1(&logpushjobsapiv1.LogpushJobsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					Dataset:       core.StringPtr(dataset),
					ZoneID:        core.StringPtr(zoneID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(logpushJobsApiService).ToNot(BeNil())

				// Construct an instance of the CreateLogRetentionOptions model
				createLogRetentionOptionsModel := new(logpushjobsapiv1.CreateLogRetentionOptions)
				createLogRetentionOptionsModel.Flag = core.BoolPtr(false)
				createLogRetentionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := logpushJobsApiService.CreateLogRetention(createLogRetentionOptionsModel)
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
			dataset := "testString"
			zoneID := "testString"
			logpushJobsApiService, _ := logpushjobsapiv1.NewLogpushJobsApiV1(&logpushjobsapiv1.LogpushJobsApiV1Options{
				URL:           "http://logpushjobsapiv1modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
				Crn:           core.StringPtr(crn),
				Dataset:       core.StringPtr(dataset),
				ZoneID:        core.StringPtr(zoneID),
			})
			It(`Invoke NewCreateLogRetentionOptions successfully`, func() {
				// Construct an instance of the CreateLogRetentionOptions model
				createLogRetentionOptionsModel := logpushJobsApiService.NewCreateLogRetentionOptions()
				createLogRetentionOptionsModel.SetFlag(false)
				createLogRetentionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createLogRetentionOptionsModel).ToNot(BeNil())
				Expect(createLogRetentionOptionsModel.Flag).To(Equal(core.BoolPtr(false)))
				Expect(createLogRetentionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateLogpushJobV2Options successfully`, func() {
				// Construct an instance of the CreateLogpushJobV2RequestLogpushJobCosReq model
				createLogpushJobV2RequestModel := new(logpushjobsapiv1.CreateLogpushJobV2RequestLogpushJobCosReq)
				Expect(createLogpushJobV2RequestModel).ToNot(BeNil())
				createLogpushJobV2RequestModel.Name = core.StringPtr("My log push job")
				createLogpushJobV2RequestModel.Enabled = core.BoolPtr(false)
				createLogpushJobV2RequestModel.LogpullOptions = core.StringPtr("timestamps=rfc3339&timestamps=rfc3339")
				createLogpushJobV2RequestModel.Cos = map[string]interface{}{"anyKey": "anyValue"}
				createLogpushJobV2RequestModel.OwnershipChallenge = core.StringPtr("00000000000000000000000000000000")
				createLogpushJobV2RequestModel.Dataset = core.StringPtr("http_requests")
				createLogpushJobV2RequestModel.Frequency = core.StringPtr("high")
				Expect(createLogpushJobV2RequestModel.Name).To(Equal(core.StringPtr("My log push job")))
				Expect(createLogpushJobV2RequestModel.Enabled).To(Equal(core.BoolPtr(false)))
				Expect(createLogpushJobV2RequestModel.LogpullOptions).To(Equal(core.StringPtr("timestamps=rfc3339&timestamps=rfc3339")))
				Expect(createLogpushJobV2RequestModel.Cos).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))
				Expect(createLogpushJobV2RequestModel.OwnershipChallenge).To(Equal(core.StringPtr("00000000000000000000000000000000")))
				Expect(createLogpushJobV2RequestModel.Dataset).To(Equal(core.StringPtr("http_requests")))
				Expect(createLogpushJobV2RequestModel.Frequency).To(Equal(core.StringPtr("high")))

				// Construct an instance of the CreateLogpushJobV2Options model
				createLogpushJobV2OptionsModel := logpushJobsApiService.NewCreateLogpushJobV2Options()
				createLogpushJobV2OptionsModel.SetCreateLogpushJobV2Request(createLogpushJobV2RequestModel)
				createLogpushJobV2OptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createLogpushJobV2OptionsModel).ToNot(BeNil())
				Expect(createLogpushJobV2OptionsModel.CreateLogpushJobV2Request).To(Equal(createLogpushJobV2RequestModel))
				Expect(createLogpushJobV2OptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteLogpushJobV2Options successfully`, func() {
				// Construct an instance of the DeleteLogpushJobV2Options model
				jobID := "testString"
				deleteLogpushJobV2OptionsModel := logpushJobsApiService.NewDeleteLogpushJobV2Options(jobID)
				deleteLogpushJobV2OptionsModel.SetJobID("testString")
				deleteLogpushJobV2OptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteLogpushJobV2OptionsModel).ToNot(BeNil())
				Expect(deleteLogpushJobV2OptionsModel.JobID).To(Equal(core.StringPtr("testString")))
				Expect(deleteLogpushJobV2OptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetLogpushJobV2Options successfully`, func() {
				// Construct an instance of the GetLogpushJobV2Options model
				jobID := "testString"
				getLogpushJobV2OptionsModel := logpushJobsApiService.NewGetLogpushJobV2Options(jobID)
				getLogpushJobV2OptionsModel.SetJobID("testString")
				getLogpushJobV2OptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getLogpushJobV2OptionsModel).ToNot(BeNil())
				Expect(getLogpushJobV2OptionsModel.JobID).To(Equal(core.StringPtr("testString")))
				Expect(getLogpushJobV2OptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetLogpushJobsV2Options successfully`, func() {
				// Construct an instance of the GetLogpushJobsV2Options model
				getLogpushJobsV2OptionsModel := logpushJobsApiService.NewGetLogpushJobsV2Options()
				getLogpushJobsV2OptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getLogpushJobsV2OptionsModel).ToNot(BeNil())
				Expect(getLogpushJobsV2OptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetLogpushOwnershipV2Options successfully`, func() {
				// Construct an instance of the GetLogpushOwnershipV2Options model
				getLogpushOwnershipV2OptionsModel := logpushJobsApiService.NewGetLogpushOwnershipV2Options()
				getLogpushOwnershipV2OptionsModel.SetCos(map[string]interface{}{"anyKey": "anyValue"})
				getLogpushOwnershipV2OptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getLogpushOwnershipV2OptionsModel).ToNot(BeNil())
				Expect(getLogpushOwnershipV2OptionsModel.Cos).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))
				Expect(getLogpushOwnershipV2OptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetLogsRetentionOptions successfully`, func() {
				// Construct an instance of the GetLogsRetentionOptions model
				getLogsRetentionOptionsModel := logpushJobsApiService.NewGetLogsRetentionOptions()
				getLogsRetentionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getLogsRetentionOptionsModel).ToNot(BeNil())
				Expect(getLogsRetentionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListFieldsForDatasetV2Options successfully`, func() {
				// Construct an instance of the ListFieldsForDatasetV2Options model
				listFieldsForDatasetV2OptionsModel := logpushJobsApiService.NewListFieldsForDatasetV2Options()
				listFieldsForDatasetV2OptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listFieldsForDatasetV2OptionsModel).ToNot(BeNil())
				Expect(listFieldsForDatasetV2OptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListLogpushJobsForDatasetV2Options successfully`, func() {
				// Construct an instance of the ListLogpushJobsForDatasetV2Options model
				listLogpushJobsForDatasetV2OptionsModel := logpushJobsApiService.NewListLogpushJobsForDatasetV2Options()
				listLogpushJobsForDatasetV2OptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listLogpushJobsForDatasetV2OptionsModel).ToNot(BeNil())
				Expect(listLogpushJobsForDatasetV2OptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewLogpushJobIbmclReqIbmcl successfully`, func() {
				instanceID := "90d208cc-e1dd-4fb2-a938-358e5996f056"
				region := "eu-es"
				apiKey := "XXXXXXXXXXXXXX"
				_model, err := logpushJobsApiService.NewLogpushJobIbmclReqIbmcl(instanceID, region, apiKey)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewUpdateLogpushJobV2Options successfully`, func() {
				// Construct an instance of the UpdateLogpushJobV2RequestLogpushJobsUpdateCosReq model
				updateLogpushJobV2RequestModel := new(logpushjobsapiv1.UpdateLogpushJobV2RequestLogpushJobsUpdateCosReq)
				Expect(updateLogpushJobV2RequestModel).ToNot(BeNil())
				updateLogpushJobV2RequestModel.Enabled = core.BoolPtr(false)
				updateLogpushJobV2RequestModel.LogpullOptions = core.StringPtr("timestamps=rfc3339&timestamps=rfc3339")
				updateLogpushJobV2RequestModel.Cos = map[string]interface{}{"anyKey": "anyValue"}
				updateLogpushJobV2RequestModel.OwnershipChallenge = core.StringPtr("00000000000000000000000000000000")
				updateLogpushJobV2RequestModel.Frequency = core.StringPtr("high")
				Expect(updateLogpushJobV2RequestModel.Enabled).To(Equal(core.BoolPtr(false)))
				Expect(updateLogpushJobV2RequestModel.LogpullOptions).To(Equal(core.StringPtr("timestamps=rfc3339&timestamps=rfc3339")))
				Expect(updateLogpushJobV2RequestModel.Cos).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))
				Expect(updateLogpushJobV2RequestModel.OwnershipChallenge).To(Equal(core.StringPtr("00000000000000000000000000000000")))
				Expect(updateLogpushJobV2RequestModel.Frequency).To(Equal(core.StringPtr("high")))

				// Construct an instance of the UpdateLogpushJobV2Options model
				jobID := "testString"
				updateLogpushJobV2OptionsModel := logpushJobsApiService.NewUpdateLogpushJobV2Options(jobID)
				updateLogpushJobV2OptionsModel.SetJobID("testString")
				updateLogpushJobV2OptionsModel.SetUpdateLogpushJobV2Request(updateLogpushJobV2RequestModel)
				updateLogpushJobV2OptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateLogpushJobV2OptionsModel).ToNot(BeNil())
				Expect(updateLogpushJobV2OptionsModel.JobID).To(Equal(core.StringPtr("testString")))
				Expect(updateLogpushJobV2OptionsModel.UpdateLogpushJobV2Request).To(Equal(updateLogpushJobV2RequestModel))
				Expect(updateLogpushJobV2OptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewValidateLogpushOwnershipChallengeV2Options successfully`, func() {
				// Construct an instance of the ValidateLogpushOwnershipChallengeV2Options model
				validateLogpushOwnershipChallengeV2OptionsModel := logpushJobsApiService.NewValidateLogpushOwnershipChallengeV2Options()
				validateLogpushOwnershipChallengeV2OptionsModel.SetCos(map[string]interface{}{"anyKey": "anyValue"})
				validateLogpushOwnershipChallengeV2OptionsModel.SetOwnershipChallenge("00000000000000000000")
				validateLogpushOwnershipChallengeV2OptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(validateLogpushOwnershipChallengeV2OptionsModel).ToNot(BeNil())
				Expect(validateLogpushOwnershipChallengeV2OptionsModel.Cos).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))
				Expect(validateLogpushOwnershipChallengeV2OptionsModel.OwnershipChallenge).To(Equal(core.StringPtr("00000000000000000000")))
				Expect(validateLogpushOwnershipChallengeV2OptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateLogpushJobV2RequestLogpushJobCosReq successfully`, func() {
				cos := map[string]interface{}{"anyKey": "anyValue"}
				ownershipChallenge := "00000000000000000000000000000000"
				_model, err := logpushJobsApiService.NewCreateLogpushJobV2RequestLogpushJobCosReq(cos, ownershipChallenge)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewCreateLogpushJobV2RequestLogpushJobGenericReq successfully`, func() {
				destinationConf := "s3://mybucket/logs?region=us-west-2"
				_model, err := logpushJobsApiService.NewCreateLogpushJobV2RequestLogpushJobGenericReq(destinationConf)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewCreateLogpushJobV2RequestLogpushJobIbmclReq successfully`, func() {
				var ibmcl *logpushjobsapiv1.LogpushJobIbmclReqIbmcl = nil
				_, err := logpushJobsApiService.NewCreateLogpushJobV2RequestLogpushJobIbmclReq(ibmcl)
				Expect(err).ToNot(BeNil())
			})
			It(`Invoke NewCreateLogpushJobV2RequestLogpushJobLogdnaReq successfully`, func() {
				logdna := map[string]interface{}{"anyKey": "anyValue"}
				_model, err := logpushJobsApiService.NewCreateLogpushJobV2RequestLogpushJobLogdnaReq(logdna)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
		})
	})
	Describe(`Model unmarshaling tests`, func() {
		It(`Invoke UnmarshalCreateLogpushJobV2Request successfully`, func() {
			// Construct an instance of the model.
			model := new(logpushjobsapiv1.CreateLogpushJobV2Request)
			model.Name = core.StringPtr("My log push job")
			model.Enabled = core.BoolPtr(false)
			model.LogpullOptions = core.StringPtr("timestamps=rfc3339&timestamps=rfc3339")
			model.Cos = map[string]interface{}{"anyKey": "anyValue"}
			model.OwnershipChallenge = core.StringPtr("00000000000000000000000000000000")
			model.Dataset = core.StringPtr("http_requests")
			model.Frequency = core.StringPtr("high")
			model.Logdna = map[string]interface{}{"anyKey": "anyValue"}
			model.Ibmcl = nil
			model.DestinationConf = core.StringPtr("s3://mybucket/logs?region=us-west-2")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *logpushjobsapiv1.CreateLogpushJobV2Request
			err = logpushjobsapiv1.UnmarshalCreateLogpushJobV2Request(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalLogpushJobIbmclReqIbmcl successfully`, func() {
			// Construct an instance of the model.
			model := new(logpushjobsapiv1.LogpushJobIbmclReqIbmcl)
			model.InstanceID = core.StringPtr("90d208cc-e1dd-4fb2-a938-358e5996f056")
			model.Region = core.StringPtr("eu-es")
			model.ApiKey = core.StringPtr("XXXXXXXXXXXXXX")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *logpushjobsapiv1.LogpushJobIbmclReqIbmcl
			err = logpushjobsapiv1.UnmarshalLogpushJobIbmclReqIbmcl(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalLogpushJobsUpdateIbmclReqIbmcl successfully`, func() {
			// Construct an instance of the model.
			model := new(logpushjobsapiv1.LogpushJobsUpdateIbmclReqIbmcl)
			model.InstanceID = core.StringPtr("90d208cc-e1dd-4fb2-a938-358e5996f056")
			model.Region = core.StringPtr("eu-es")
			model.ApiKey = core.StringPtr("XXXXXXXXXXXXXX")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *logpushjobsapiv1.LogpushJobsUpdateIbmclReqIbmcl
			err = logpushjobsapiv1.UnmarshalLogpushJobsUpdateIbmclReqIbmcl(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalUpdateLogpushJobV2Request successfully`, func() {
			// Construct an instance of the model.
			model := new(logpushjobsapiv1.UpdateLogpushJobV2Request)
			model.Enabled = core.BoolPtr(false)
			model.LogpullOptions = core.StringPtr("timestamps=rfc3339&timestamps=rfc3339")
			model.Cos = map[string]interface{}{"anyKey": "anyValue"}
			model.OwnershipChallenge = core.StringPtr("00000000000000000000000000000000")
			model.Frequency = core.StringPtr("high")
			model.Logdna = map[string]interface{}{"anyKey": "anyValue"}
			model.Ibmcl = nil
			model.Name = core.StringPtr("My log push job")
			model.DestinationConf = core.StringPtr("s3://mybucket/logs?region=us-west-2")
			model.Dataset = core.StringPtr("http_requests")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *logpushjobsapiv1.UpdateLogpushJobV2Request
			err = logpushjobsapiv1.UnmarshalUpdateLogpushJobV2Request(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalCreateLogpushJobV2RequestLogpushJobCosReq successfully`, func() {
			// Construct an instance of the model.
			model := new(logpushjobsapiv1.CreateLogpushJobV2RequestLogpushJobCosReq)
			model.Name = core.StringPtr("My log push job")
			model.Enabled = core.BoolPtr(false)
			model.LogpullOptions = core.StringPtr("timestamps=rfc3339&timestamps=rfc3339")
			model.Cos = map[string]interface{}{"anyKey": "anyValue"}
			model.OwnershipChallenge = core.StringPtr("00000000000000000000000000000000")
			model.Dataset = core.StringPtr("http_requests")
			model.Frequency = core.StringPtr("high")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *logpushjobsapiv1.CreateLogpushJobV2RequestLogpushJobCosReq
			err = logpushjobsapiv1.UnmarshalCreateLogpushJobV2RequestLogpushJobCosReq(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalCreateLogpushJobV2RequestLogpushJobGenericReq successfully`, func() {
			// Construct an instance of the model.
			model := new(logpushjobsapiv1.CreateLogpushJobV2RequestLogpushJobGenericReq)
			model.Name = core.StringPtr("My log push job")
			model.Enabled = core.BoolPtr(false)
			model.LogpullOptions = core.StringPtr("timestamps=rfc3339&timestamps=rfc3339")
			model.DestinationConf = core.StringPtr("s3://mybucket/logs?region=us-west-2")
			model.Dataset = core.StringPtr("http_requests")
			model.Frequency = core.StringPtr("high")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *logpushjobsapiv1.CreateLogpushJobV2RequestLogpushJobGenericReq
			err = logpushjobsapiv1.UnmarshalCreateLogpushJobV2RequestLogpushJobGenericReq(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalCreateLogpushJobV2RequestLogpushJobIbmclReq successfully`, func() {
			// Construct an instance of the model.
			model := new(logpushjobsapiv1.CreateLogpushJobV2RequestLogpushJobIbmclReq)
			model.Name = core.StringPtr("My log push job")
			model.Enabled = core.BoolPtr(false)
			model.LogpullOptions = core.StringPtr("timestamps=rfc3339&timestamps=rfc3339")
			model.Ibmcl = nil
			model.Dataset = core.StringPtr("http_requests")
			model.Frequency = core.StringPtr("high")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *logpushjobsapiv1.CreateLogpushJobV2RequestLogpushJobIbmclReq
			err = logpushjobsapiv1.UnmarshalCreateLogpushJobV2RequestLogpushJobIbmclReq(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalCreateLogpushJobV2RequestLogpushJobLogdnaReq successfully`, func() {
			// Construct an instance of the model.
			model := new(logpushjobsapiv1.CreateLogpushJobV2RequestLogpushJobLogdnaReq)
			model.Name = core.StringPtr("My log push job")
			model.Enabled = core.BoolPtr(false)
			model.LogpullOptions = core.StringPtr("timestamps=rfc3339&timestamps=rfc3339")
			model.Logdna = map[string]interface{}{"anyKey": "anyValue"}
			model.Dataset = core.StringPtr("http_requests")
			model.Frequency = core.StringPtr("high")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *logpushjobsapiv1.CreateLogpushJobV2RequestLogpushJobLogdnaReq
			err = logpushjobsapiv1.UnmarshalCreateLogpushJobV2RequestLogpushJobLogdnaReq(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalUpdateLogpushJobV2RequestLogpushJobsUpdateCosReq successfully`, func() {
			// Construct an instance of the model.
			model := new(logpushjobsapiv1.UpdateLogpushJobV2RequestLogpushJobsUpdateCosReq)
			model.Enabled = core.BoolPtr(false)
			model.LogpullOptions = core.StringPtr("timestamps=rfc3339&timestamps=rfc3339")
			model.Cos = map[string]interface{}{"anyKey": "anyValue"}
			model.OwnershipChallenge = core.StringPtr("00000000000000000000000000000000")
			model.Frequency = core.StringPtr("high")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *logpushjobsapiv1.UpdateLogpushJobV2RequestLogpushJobsUpdateCosReq
			err = logpushjobsapiv1.UnmarshalUpdateLogpushJobV2RequestLogpushJobsUpdateCosReq(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalUpdateLogpushJobV2RequestLogpushJobsUpdateGenericReq successfully`, func() {
			// Construct an instance of the model.
			model := new(logpushjobsapiv1.UpdateLogpushJobV2RequestLogpushJobsUpdateGenericReq)
			model.Name = core.StringPtr("My log push job")
			model.Enabled = core.BoolPtr(false)
			model.LogpullOptions = core.StringPtr("timestamps=rfc3339&timestamps=rfc3339")
			model.DestinationConf = core.StringPtr("s3://mybucket/logs?region=us-west-2")
			model.Dataset = core.StringPtr("http_requests")
			model.Frequency = core.StringPtr("high")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *logpushjobsapiv1.UpdateLogpushJobV2RequestLogpushJobsUpdateGenericReq
			err = logpushjobsapiv1.UnmarshalUpdateLogpushJobV2RequestLogpushJobsUpdateGenericReq(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalUpdateLogpushJobV2RequestLogpushJobsUpdateIbmclReq successfully`, func() {
			// Construct an instance of the model.
			model := new(logpushjobsapiv1.UpdateLogpushJobV2RequestLogpushJobsUpdateIbmclReq)
			model.Enabled = core.BoolPtr(false)
			model.LogpullOptions = core.StringPtr("timestamps=rfc3339&timestamps=rfc3339")
			model.Ibmcl = nil
			model.Frequency = core.StringPtr("high")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *logpushjobsapiv1.UpdateLogpushJobV2RequestLogpushJobsUpdateIbmclReq
			err = logpushjobsapiv1.UnmarshalUpdateLogpushJobV2RequestLogpushJobsUpdateIbmclReq(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalUpdateLogpushJobV2RequestLogpushJobsUpdateLogdnaReq successfully`, func() {
			// Construct an instance of the model.
			model := new(logpushjobsapiv1.UpdateLogpushJobV2RequestLogpushJobsUpdateLogdnaReq)
			model.Enabled = core.BoolPtr(false)
			model.LogpullOptions = core.StringPtr("timestamps=rfc3339&timestamps=rfc3339")
			model.Logdna = map[string]interface{}{"anyKey": "anyValue"}
			model.Frequency = core.StringPtr("high")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *logpushjobsapiv1.UpdateLogpushJobV2RequestLogpushJobsUpdateLogdnaReq
			err = logpushjobsapiv1.UnmarshalUpdateLogpushJobV2RequestLogpushJobsUpdateLogdnaReq(raw, &result)
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
