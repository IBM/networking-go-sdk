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

package logpushjobsapiv1_test

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
	"github.com/IBM/networking-go-sdk/logpushjobsapiv1"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe(`LogpushJobsApiV1`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		crn := "testString"
		zoneID := "testString"
		dataset := "http_requests"
		It(`Instantiate service client`, func() {
			logpushJobsApiService, serviceErr := logpushjobsapiv1.NewLogpushJobsApiV1(&logpushjobsapiv1.LogpushJobsApiV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
				Crn:           core.StringPtr(crn),
				ZoneID:        core.StringPtr(zoneID),
				Dataset:       core.StringPtr(dataset),
			})
			Expect(logpushJobsApiService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			logpushJobsApiService, serviceErr := logpushjobsapiv1.NewLogpushJobsApiV1(&logpushjobsapiv1.LogpushJobsApiV1Options{
				URL:     "{BAD_URL_STRING",
				Crn:     core.StringPtr(crn),
				ZoneID:  core.StringPtr(zoneID),
				Dataset: core.StringPtr(dataset),
			})
			Expect(logpushJobsApiService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			logpushJobsApiService, serviceErr := logpushjobsapiv1.NewLogpushJobsApiV1(&logpushjobsapiv1.LogpushJobsApiV1Options{
				URL:     "https://logpushjobsapiv1/api",
				Crn:     core.StringPtr(crn),
				ZoneID:  core.StringPtr(zoneID),
				Dataset: core.StringPtr(dataset),
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
		zoneID := "testString"
		dataset := "http_requests"
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
					ZoneID:  core.StringPtr(zoneID),
					Dataset: core.StringPtr(dataset),
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
					ZoneID:  core.StringPtr(zoneID),
					Dataset: core.StringPtr(dataset),
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
					ZoneID:  core.StringPtr(zoneID),
					Dataset: core.StringPtr(dataset),
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
				ZoneID:  core.StringPtr(zoneID),
				Dataset: core.StringPtr(dataset),
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
				ZoneID:  core.StringPtr(zoneID),
				Dataset: core.StringPtr(dataset),
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
	Describe(`GetLogpushJobs(getLogpushJobsOptions *GetLogpushJobsOptions) - Operation response error`, func() {
		crn := "testString"
		zoneID := "testString"
		dataset := "http_requests"
		getLogpushJobsPath := "/v1/testString/zones/testString/logpush/jobs"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getLogpushJobsPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetLogpushJobs with error: Operation response processing error`, func() {
				logpushJobsApiService, serviceErr := logpushjobsapiv1.NewLogpushJobsApiV1(&logpushjobsapiv1.LogpushJobsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ZoneID:        core.StringPtr(zoneID),
					Dataset:       core.StringPtr(dataset),
				})
				Expect(serviceErr).To(BeNil())
				Expect(logpushJobsApiService).ToNot(BeNil())

				// Construct an instance of the GetLogpushJobsOptions model
				getLogpushJobsOptionsModel := new(logpushjobsapiv1.GetLogpushJobsOptions)
				getLogpushJobsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := logpushJobsApiService.GetLogpushJobs(getLogpushJobsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				logpushJobsApiService.EnableRetries(0, 0)
				result, response, operationErr = logpushJobsApiService.GetLogpushJobs(getLogpushJobsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetLogpushJobs(getLogpushJobsOptions *GetLogpushJobsOptions)`, func() {
		crn := "testString"
		zoneID := "testString"
		dataset := "http_requests"
		getLogpushJobsPath := "/v1/testString/zones/testString/logpush/jobs"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getLogpushJobsPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": [{"id": 5850, "name": "My log push job", "enabled": false, "dataset": "firewall_events", "frequency": "high", "logpull_options": "timestamps=rfc3339&timestamps=rfc3339", "destination_conf": "cos://cos-bucket001?region=us-south&instance-id=231f5467-3072-4cb9-9e39-a906fa3032ea", "last_complete": "2022-01-15T16:33:31.834209Z", "last_error": "2022-01-15T16:33:31.834209Z", "error_message": "ErrorMessage"}]}`)
				}))
			})
			It(`Invoke GetLogpushJobs successfully with retries`, func() {
				logpushJobsApiService, serviceErr := logpushjobsapiv1.NewLogpushJobsApiV1(&logpushjobsapiv1.LogpushJobsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ZoneID:        core.StringPtr(zoneID),
					Dataset:       core.StringPtr(dataset),
				})
				Expect(serviceErr).To(BeNil())
				Expect(logpushJobsApiService).ToNot(BeNil())
				logpushJobsApiService.EnableRetries(0, 0)

				// Construct an instance of the GetLogpushJobsOptions model
				getLogpushJobsOptionsModel := new(logpushjobsapiv1.GetLogpushJobsOptions)
				getLogpushJobsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := logpushJobsApiService.GetLogpushJobsWithContext(ctx, getLogpushJobsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				logpushJobsApiService.DisableRetries()
				result, response, operationErr := logpushJobsApiService.GetLogpushJobs(getLogpushJobsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = logpushJobsApiService.GetLogpushJobsWithContext(ctx, getLogpushJobsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getLogpushJobsPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": [{"id": 5850, "name": "My log push job", "enabled": false, "dataset": "firewall_events", "frequency": "high", "logpull_options": "timestamps=rfc3339&timestamps=rfc3339", "destination_conf": "cos://cos-bucket001?region=us-south&instance-id=231f5467-3072-4cb9-9e39-a906fa3032ea", "last_complete": "2022-01-15T16:33:31.834209Z", "last_error": "2022-01-15T16:33:31.834209Z", "error_message": "ErrorMessage"}]}`)
				}))
			})
			It(`Invoke GetLogpushJobs successfully`, func() {
				logpushJobsApiService, serviceErr := logpushjobsapiv1.NewLogpushJobsApiV1(&logpushjobsapiv1.LogpushJobsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ZoneID:        core.StringPtr(zoneID),
					Dataset:       core.StringPtr(dataset),
				})
				Expect(serviceErr).To(BeNil())
				Expect(logpushJobsApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := logpushJobsApiService.GetLogpushJobs(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetLogpushJobsOptions model
				getLogpushJobsOptionsModel := new(logpushjobsapiv1.GetLogpushJobsOptions)
				getLogpushJobsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = logpushJobsApiService.GetLogpushJobs(getLogpushJobsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetLogpushJobs with error: Operation request error`, func() {
				logpushJobsApiService, serviceErr := logpushjobsapiv1.NewLogpushJobsApiV1(&logpushjobsapiv1.LogpushJobsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ZoneID:        core.StringPtr(zoneID),
					Dataset:       core.StringPtr(dataset),
				})
				Expect(serviceErr).To(BeNil())
				Expect(logpushJobsApiService).ToNot(BeNil())

				// Construct an instance of the GetLogpushJobsOptions model
				getLogpushJobsOptionsModel := new(logpushjobsapiv1.GetLogpushJobsOptions)
				getLogpushJobsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := logpushJobsApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := logpushJobsApiService.GetLogpushJobs(getLogpushJobsOptionsModel)
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
			It(`Invoke GetLogpushJobs successfully`, func() {
				logpushJobsApiService, serviceErr := logpushjobsapiv1.NewLogpushJobsApiV1(&logpushjobsapiv1.LogpushJobsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ZoneID:        core.StringPtr(zoneID),
					Dataset:       core.StringPtr(dataset),
				})
				Expect(serviceErr).To(BeNil())
				Expect(logpushJobsApiService).ToNot(BeNil())

				// Construct an instance of the GetLogpushJobsOptions model
				getLogpushJobsOptionsModel := new(logpushjobsapiv1.GetLogpushJobsOptions)
				getLogpushJobsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := logpushJobsApiService.GetLogpushJobs(getLogpushJobsOptionsModel)
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
	Describe(`CreateLogpushJob(createLogpushJobOptions *CreateLogpushJobOptions) - Operation response error`, func() {
		crn := "testString"
		zoneID := "testString"
		dataset := "http_requests"
		createLogpushJobPath := "/v1/testString/zones/testString/logpush/jobs"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createLogpushJobPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateLogpushJob with error: Operation response processing error`, func() {
				logpushJobsApiService, serviceErr := logpushjobsapiv1.NewLogpushJobsApiV1(&logpushjobsapiv1.LogpushJobsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ZoneID:        core.StringPtr(zoneID),
					Dataset:       core.StringPtr(dataset),
				})
				Expect(serviceErr).To(BeNil())
				Expect(logpushJobsApiService).ToNot(BeNil())

				// Construct an instance of the CreateLogpushJobOptions model
				createLogpushJobOptionsModel := new(logpushjobsapiv1.CreateLogpushJobOptions)
				createLogpushJobOptionsModel.Name = core.StringPtr("My log push job")
				createLogpushJobOptionsModel.Enabled = core.BoolPtr(false)
				createLogpushJobOptionsModel.LogpullOptions = core.StringPtr("timestamps=rfc3339&timestamps=rfc3339")
				createLogpushJobOptionsModel.DestinationConf = core.StringPtr("cos://cos-bucket001?region=us-south&instance-id=231f5467-3072-4cb9-9e39-a906fa3032ea")
				createLogpushJobOptionsModel.OwnershipChallenge = core.StringPtr("00000000000000000000000000000000")
				createLogpushJobOptionsModel.Dataset = core.StringPtr("firewall_events")
				createLogpushJobOptionsModel.Frequency = core.StringPtr("high")
				createLogpushJobOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := logpushJobsApiService.CreateLogpushJob(createLogpushJobOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				logpushJobsApiService.EnableRetries(0, 0)
				result, response, operationErr = logpushJobsApiService.CreateLogpushJob(createLogpushJobOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateLogpushJob(createLogpushJobOptions *CreateLogpushJobOptions)`, func() {
		crn := "testString"
		zoneID := "testString"
		dataset := "http_requests"
		createLogpushJobPath := "/v1/testString/zones/testString/logpush/jobs"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createLogpushJobPath))
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
			It(`Invoke CreateLogpushJob successfully with retries`, func() {
				logpushJobsApiService, serviceErr := logpushjobsapiv1.NewLogpushJobsApiV1(&logpushjobsapiv1.LogpushJobsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ZoneID:        core.StringPtr(zoneID),
					Dataset:       core.StringPtr(dataset),
				})
				Expect(serviceErr).To(BeNil())
				Expect(logpushJobsApiService).ToNot(BeNil())
				logpushJobsApiService.EnableRetries(0, 0)

				// Construct an instance of the CreateLogpushJobOptions model
				createLogpushJobOptionsModel := new(logpushjobsapiv1.CreateLogpushJobOptions)
				createLogpushJobOptionsModel.Name = core.StringPtr("My log push job")
				createLogpushJobOptionsModel.Enabled = core.BoolPtr(false)
				createLogpushJobOptionsModel.LogpullOptions = core.StringPtr("timestamps=rfc3339&timestamps=rfc3339")
				createLogpushJobOptionsModel.DestinationConf = core.StringPtr("cos://cos-bucket001?region=us-south&instance-id=231f5467-3072-4cb9-9e39-a906fa3032ea")
				createLogpushJobOptionsModel.OwnershipChallenge = core.StringPtr("00000000000000000000000000000000")
				createLogpushJobOptionsModel.Dataset = core.StringPtr("firewall_events")
				createLogpushJobOptionsModel.Frequency = core.StringPtr("high")
				createLogpushJobOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := logpushJobsApiService.CreateLogpushJobWithContext(ctx, createLogpushJobOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				logpushJobsApiService.DisableRetries()
				result, response, operationErr := logpushJobsApiService.CreateLogpushJob(createLogpushJobOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = logpushJobsApiService.CreateLogpushJobWithContext(ctx, createLogpushJobOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createLogpushJobPath))
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
			It(`Invoke CreateLogpushJob successfully`, func() {
				logpushJobsApiService, serviceErr := logpushjobsapiv1.NewLogpushJobsApiV1(&logpushjobsapiv1.LogpushJobsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ZoneID:        core.StringPtr(zoneID),
					Dataset:       core.StringPtr(dataset),
				})
				Expect(serviceErr).To(BeNil())
				Expect(logpushJobsApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := logpushJobsApiService.CreateLogpushJob(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateLogpushJobOptions model
				createLogpushJobOptionsModel := new(logpushjobsapiv1.CreateLogpushJobOptions)
				createLogpushJobOptionsModel.Name = core.StringPtr("My log push job")
				createLogpushJobOptionsModel.Enabled = core.BoolPtr(false)
				createLogpushJobOptionsModel.LogpullOptions = core.StringPtr("timestamps=rfc3339&timestamps=rfc3339")
				createLogpushJobOptionsModel.DestinationConf = core.StringPtr("cos://cos-bucket001?region=us-south&instance-id=231f5467-3072-4cb9-9e39-a906fa3032ea")
				createLogpushJobOptionsModel.OwnershipChallenge = core.StringPtr("00000000000000000000000000000000")
				createLogpushJobOptionsModel.Dataset = core.StringPtr("firewall_events")
				createLogpushJobOptionsModel.Frequency = core.StringPtr("high")
				createLogpushJobOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = logpushJobsApiService.CreateLogpushJob(createLogpushJobOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateLogpushJob with error: Operation request error`, func() {
				logpushJobsApiService, serviceErr := logpushjobsapiv1.NewLogpushJobsApiV1(&logpushjobsapiv1.LogpushJobsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ZoneID:        core.StringPtr(zoneID),
					Dataset:       core.StringPtr(dataset),
				})
				Expect(serviceErr).To(BeNil())
				Expect(logpushJobsApiService).ToNot(BeNil())

				// Construct an instance of the CreateLogpushJobOptions model
				createLogpushJobOptionsModel := new(logpushjobsapiv1.CreateLogpushJobOptions)
				createLogpushJobOptionsModel.Name = core.StringPtr("My log push job")
				createLogpushJobOptionsModel.Enabled = core.BoolPtr(false)
				createLogpushJobOptionsModel.LogpullOptions = core.StringPtr("timestamps=rfc3339&timestamps=rfc3339")
				createLogpushJobOptionsModel.DestinationConf = core.StringPtr("cos://cos-bucket001?region=us-south&instance-id=231f5467-3072-4cb9-9e39-a906fa3032ea")
				createLogpushJobOptionsModel.OwnershipChallenge = core.StringPtr("00000000000000000000000000000000")
				createLogpushJobOptionsModel.Dataset = core.StringPtr("firewall_events")
				createLogpushJobOptionsModel.Frequency = core.StringPtr("high")
				createLogpushJobOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := logpushJobsApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := logpushJobsApiService.CreateLogpushJob(createLogpushJobOptionsModel)
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
			It(`Invoke CreateLogpushJob successfully`, func() {
				logpushJobsApiService, serviceErr := logpushjobsapiv1.NewLogpushJobsApiV1(&logpushjobsapiv1.LogpushJobsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ZoneID:        core.StringPtr(zoneID),
					Dataset:       core.StringPtr(dataset),
				})
				Expect(serviceErr).To(BeNil())
				Expect(logpushJobsApiService).ToNot(BeNil())

				// Construct an instance of the CreateLogpushJobOptions model
				createLogpushJobOptionsModel := new(logpushjobsapiv1.CreateLogpushJobOptions)
				createLogpushJobOptionsModel.Name = core.StringPtr("My log push job")
				createLogpushJobOptionsModel.Enabled = core.BoolPtr(false)
				createLogpushJobOptionsModel.LogpullOptions = core.StringPtr("timestamps=rfc3339&timestamps=rfc3339")
				createLogpushJobOptionsModel.DestinationConf = core.StringPtr("cos://cos-bucket001?region=us-south&instance-id=231f5467-3072-4cb9-9e39-a906fa3032ea")
				createLogpushJobOptionsModel.OwnershipChallenge = core.StringPtr("00000000000000000000000000000000")
				createLogpushJobOptionsModel.Dataset = core.StringPtr("firewall_events")
				createLogpushJobOptionsModel.Frequency = core.StringPtr("high")
				createLogpushJobOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := logpushJobsApiService.CreateLogpushJob(createLogpushJobOptionsModel)
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
	Describe(`GetLogpushJob(getLogpushJobOptions *GetLogpushJobOptions) - Operation response error`, func() {
		crn := "testString"
		zoneID := "testString"
		dataset := "http_requests"
		getLogpushJobPath := "/v1/testString/zones/testString/logpush/jobs/38"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getLogpushJobPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetLogpushJob with error: Operation response processing error`, func() {
				logpushJobsApiService, serviceErr := logpushjobsapiv1.NewLogpushJobsApiV1(&logpushjobsapiv1.LogpushJobsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ZoneID:        core.StringPtr(zoneID),
					Dataset:       core.StringPtr(dataset),
				})
				Expect(serviceErr).To(BeNil())
				Expect(logpushJobsApiService).ToNot(BeNil())

				// Construct an instance of the GetLogpushJobOptions model
				getLogpushJobOptionsModel := new(logpushjobsapiv1.GetLogpushJobOptions)
				getLogpushJobOptionsModel.JobID = core.Int64Ptr(int64(38))
				getLogpushJobOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := logpushJobsApiService.GetLogpushJob(getLogpushJobOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				logpushJobsApiService.EnableRetries(0, 0)
				result, response, operationErr = logpushJobsApiService.GetLogpushJob(getLogpushJobOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetLogpushJob(getLogpushJobOptions *GetLogpushJobOptions)`, func() {
		crn := "testString"
		zoneID := "testString"
		dataset := "http_requests"
		getLogpushJobPath := "/v1/testString/zones/testString/logpush/jobs/38"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getLogpushJobPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": 5850, "name": "My log push job", "enabled": false, "dataset": "firewall_events", "frequency": "high", "logpull_options": "timestamps=rfc3339&timestamps=rfc3339", "destination_conf": "cos://cos-bucket001?region=us-south&instance-id=231f5467-3072-4cb9-9e39-a906fa3032ea", "last_complete": "2022-01-15T16:33:31.834209Z", "last_error": "2022-01-15T16:33:31.834209Z", "error_message": "ErrorMessage"}}`)
				}))
			})
			It(`Invoke GetLogpushJob successfully with retries`, func() {
				logpushJobsApiService, serviceErr := logpushjobsapiv1.NewLogpushJobsApiV1(&logpushjobsapiv1.LogpushJobsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ZoneID:        core.StringPtr(zoneID),
					Dataset:       core.StringPtr(dataset),
				})
				Expect(serviceErr).To(BeNil())
				Expect(logpushJobsApiService).ToNot(BeNil())
				logpushJobsApiService.EnableRetries(0, 0)

				// Construct an instance of the GetLogpushJobOptions model
				getLogpushJobOptionsModel := new(logpushjobsapiv1.GetLogpushJobOptions)
				getLogpushJobOptionsModel.JobID = core.Int64Ptr(int64(38))
				getLogpushJobOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := logpushJobsApiService.GetLogpushJobWithContext(ctx, getLogpushJobOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				logpushJobsApiService.DisableRetries()
				result, response, operationErr := logpushJobsApiService.GetLogpushJob(getLogpushJobOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = logpushJobsApiService.GetLogpushJobWithContext(ctx, getLogpushJobOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getLogpushJobPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": 5850, "name": "My log push job", "enabled": false, "dataset": "firewall_events", "frequency": "high", "logpull_options": "timestamps=rfc3339&timestamps=rfc3339", "destination_conf": "cos://cos-bucket001?region=us-south&instance-id=231f5467-3072-4cb9-9e39-a906fa3032ea", "last_complete": "2022-01-15T16:33:31.834209Z", "last_error": "2022-01-15T16:33:31.834209Z", "error_message": "ErrorMessage"}}`)
				}))
			})
			It(`Invoke GetLogpushJob successfully`, func() {
				logpushJobsApiService, serviceErr := logpushjobsapiv1.NewLogpushJobsApiV1(&logpushjobsapiv1.LogpushJobsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ZoneID:        core.StringPtr(zoneID),
					Dataset:       core.StringPtr(dataset),
				})
				Expect(serviceErr).To(BeNil())
				Expect(logpushJobsApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := logpushJobsApiService.GetLogpushJob(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetLogpushJobOptions model
				getLogpushJobOptionsModel := new(logpushjobsapiv1.GetLogpushJobOptions)
				getLogpushJobOptionsModel.JobID = core.Int64Ptr(int64(38))
				getLogpushJobOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = logpushJobsApiService.GetLogpushJob(getLogpushJobOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetLogpushJob with error: Operation validation and request error`, func() {
				logpushJobsApiService, serviceErr := logpushjobsapiv1.NewLogpushJobsApiV1(&logpushjobsapiv1.LogpushJobsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ZoneID:        core.StringPtr(zoneID),
					Dataset:       core.StringPtr(dataset),
				})
				Expect(serviceErr).To(BeNil())
				Expect(logpushJobsApiService).ToNot(BeNil())

				// Construct an instance of the GetLogpushJobOptions model
				getLogpushJobOptionsModel := new(logpushjobsapiv1.GetLogpushJobOptions)
				getLogpushJobOptionsModel.JobID = core.Int64Ptr(int64(38))
				getLogpushJobOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := logpushJobsApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := logpushJobsApiService.GetLogpushJob(getLogpushJobOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetLogpushJobOptions model with no property values
				getLogpushJobOptionsModelNew := new(logpushjobsapiv1.GetLogpushJobOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = logpushJobsApiService.GetLogpushJob(getLogpushJobOptionsModelNew)
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
			It(`Invoke GetLogpushJob successfully`, func() {
				logpushJobsApiService, serviceErr := logpushjobsapiv1.NewLogpushJobsApiV1(&logpushjobsapiv1.LogpushJobsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ZoneID:        core.StringPtr(zoneID),
					Dataset:       core.StringPtr(dataset),
				})
				Expect(serviceErr).To(BeNil())
				Expect(logpushJobsApiService).ToNot(BeNil())

				// Construct an instance of the GetLogpushJobOptions model
				getLogpushJobOptionsModel := new(logpushjobsapiv1.GetLogpushJobOptions)
				getLogpushJobOptionsModel.JobID = core.Int64Ptr(int64(38))
				getLogpushJobOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := logpushJobsApiService.GetLogpushJob(getLogpushJobOptionsModel)
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
	Describe(`UpdateLogpushJob(updateLogpushJobOptions *UpdateLogpushJobOptions) - Operation response error`, func() {
		crn := "testString"
		zoneID := "testString"
		dataset := "http_requests"
		updateLogpushJobPath := "/v1/testString/zones/testString/logpush/jobs/38"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateLogpushJobPath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateLogpushJob with error: Operation response processing error`, func() {
				logpushJobsApiService, serviceErr := logpushjobsapiv1.NewLogpushJobsApiV1(&logpushjobsapiv1.LogpushJobsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ZoneID:        core.StringPtr(zoneID),
					Dataset:       core.StringPtr(dataset),
				})
				Expect(serviceErr).To(BeNil())
				Expect(logpushJobsApiService).ToNot(BeNil())

				// Construct an instance of the UpdateLogpushJobOptions model
				updateLogpushJobOptionsModel := new(logpushjobsapiv1.UpdateLogpushJobOptions)
				updateLogpushJobOptionsModel.JobID = core.Int64Ptr(int64(38))
				updateLogpushJobOptionsModel.Enabled = core.BoolPtr(false)
				updateLogpushJobOptionsModel.LogpullOptions = core.StringPtr("timestamps=rfc3339&timestamps=rfc3339")
				updateLogpushJobOptionsModel.DestinationConf = core.StringPtr("cos://cos-bucket001?region=us-south&instance-id=231f5467-3072-4cb9-9e39-a906fa3032ea")
				updateLogpushJobOptionsModel.OwnershipChallenge = core.StringPtr("00000000000000000000000000000000")
				updateLogpushJobOptionsModel.Frequency = core.StringPtr("high")
				updateLogpushJobOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := logpushJobsApiService.UpdateLogpushJob(updateLogpushJobOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				logpushJobsApiService.EnableRetries(0, 0)
				result, response, operationErr = logpushJobsApiService.UpdateLogpushJob(updateLogpushJobOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateLogpushJob(updateLogpushJobOptions *UpdateLogpushJobOptions)`, func() {
		crn := "testString"
		zoneID := "testString"
		dataset := "http_requests"
		updateLogpushJobPath := "/v1/testString/zones/testString/logpush/jobs/38"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateLogpushJobPath))
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
			It(`Invoke UpdateLogpushJob successfully with retries`, func() {
				logpushJobsApiService, serviceErr := logpushjobsapiv1.NewLogpushJobsApiV1(&logpushjobsapiv1.LogpushJobsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ZoneID:        core.StringPtr(zoneID),
					Dataset:       core.StringPtr(dataset),
				})
				Expect(serviceErr).To(BeNil())
				Expect(logpushJobsApiService).ToNot(BeNil())
				logpushJobsApiService.EnableRetries(0, 0)

				// Construct an instance of the UpdateLogpushJobOptions model
				updateLogpushJobOptionsModel := new(logpushjobsapiv1.UpdateLogpushJobOptions)
				updateLogpushJobOptionsModel.JobID = core.Int64Ptr(int64(38))
				updateLogpushJobOptionsModel.Enabled = core.BoolPtr(false)
				updateLogpushJobOptionsModel.LogpullOptions = core.StringPtr("timestamps=rfc3339&timestamps=rfc3339")
				updateLogpushJobOptionsModel.DestinationConf = core.StringPtr("cos://cos-bucket001?region=us-south&instance-id=231f5467-3072-4cb9-9e39-a906fa3032ea")
				updateLogpushJobOptionsModel.OwnershipChallenge = core.StringPtr("00000000000000000000000000000000")
				updateLogpushJobOptionsModel.Frequency = core.StringPtr("high")
				updateLogpushJobOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := logpushJobsApiService.UpdateLogpushJobWithContext(ctx, updateLogpushJobOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				logpushJobsApiService.DisableRetries()
				result, response, operationErr := logpushJobsApiService.UpdateLogpushJob(updateLogpushJobOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = logpushJobsApiService.UpdateLogpushJobWithContext(ctx, updateLogpushJobOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(updateLogpushJobPath))
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
			It(`Invoke UpdateLogpushJob successfully`, func() {
				logpushJobsApiService, serviceErr := logpushjobsapiv1.NewLogpushJobsApiV1(&logpushjobsapiv1.LogpushJobsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ZoneID:        core.StringPtr(zoneID),
					Dataset:       core.StringPtr(dataset),
				})
				Expect(serviceErr).To(BeNil())
				Expect(logpushJobsApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := logpushJobsApiService.UpdateLogpushJob(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateLogpushJobOptions model
				updateLogpushJobOptionsModel := new(logpushjobsapiv1.UpdateLogpushJobOptions)
				updateLogpushJobOptionsModel.JobID = core.Int64Ptr(int64(38))
				updateLogpushJobOptionsModel.Enabled = core.BoolPtr(false)
				updateLogpushJobOptionsModel.LogpullOptions = core.StringPtr("timestamps=rfc3339&timestamps=rfc3339")
				updateLogpushJobOptionsModel.DestinationConf = core.StringPtr("cos://cos-bucket001?region=us-south&instance-id=231f5467-3072-4cb9-9e39-a906fa3032ea")
				updateLogpushJobOptionsModel.OwnershipChallenge = core.StringPtr("00000000000000000000000000000000")
				updateLogpushJobOptionsModel.Frequency = core.StringPtr("high")
				updateLogpushJobOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = logpushJobsApiService.UpdateLogpushJob(updateLogpushJobOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateLogpushJob with error: Operation validation and request error`, func() {
				logpushJobsApiService, serviceErr := logpushjobsapiv1.NewLogpushJobsApiV1(&logpushjobsapiv1.LogpushJobsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ZoneID:        core.StringPtr(zoneID),
					Dataset:       core.StringPtr(dataset),
				})
				Expect(serviceErr).To(BeNil())
				Expect(logpushJobsApiService).ToNot(BeNil())

				// Construct an instance of the UpdateLogpushJobOptions model
				updateLogpushJobOptionsModel := new(logpushjobsapiv1.UpdateLogpushJobOptions)
				updateLogpushJobOptionsModel.JobID = core.Int64Ptr(int64(38))
				updateLogpushJobOptionsModel.Enabled = core.BoolPtr(false)
				updateLogpushJobOptionsModel.LogpullOptions = core.StringPtr("timestamps=rfc3339&timestamps=rfc3339")
				updateLogpushJobOptionsModel.DestinationConf = core.StringPtr("cos://cos-bucket001?region=us-south&instance-id=231f5467-3072-4cb9-9e39-a906fa3032ea")
				updateLogpushJobOptionsModel.OwnershipChallenge = core.StringPtr("00000000000000000000000000000000")
				updateLogpushJobOptionsModel.Frequency = core.StringPtr("high")
				updateLogpushJobOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := logpushJobsApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := logpushJobsApiService.UpdateLogpushJob(updateLogpushJobOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateLogpushJobOptions model with no property values
				updateLogpushJobOptionsModelNew := new(logpushjobsapiv1.UpdateLogpushJobOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = logpushJobsApiService.UpdateLogpushJob(updateLogpushJobOptionsModelNew)
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
			It(`Invoke UpdateLogpushJob successfully`, func() {
				logpushJobsApiService, serviceErr := logpushjobsapiv1.NewLogpushJobsApiV1(&logpushjobsapiv1.LogpushJobsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ZoneID:        core.StringPtr(zoneID),
					Dataset:       core.StringPtr(dataset),
				})
				Expect(serviceErr).To(BeNil())
				Expect(logpushJobsApiService).ToNot(BeNil())

				// Construct an instance of the UpdateLogpushJobOptions model
				updateLogpushJobOptionsModel := new(logpushjobsapiv1.UpdateLogpushJobOptions)
				updateLogpushJobOptionsModel.JobID = core.Int64Ptr(int64(38))
				updateLogpushJobOptionsModel.Enabled = core.BoolPtr(false)
				updateLogpushJobOptionsModel.LogpullOptions = core.StringPtr("timestamps=rfc3339&timestamps=rfc3339")
				updateLogpushJobOptionsModel.DestinationConf = core.StringPtr("cos://cos-bucket001?region=us-south&instance-id=231f5467-3072-4cb9-9e39-a906fa3032ea")
				updateLogpushJobOptionsModel.OwnershipChallenge = core.StringPtr("00000000000000000000000000000000")
				updateLogpushJobOptionsModel.Frequency = core.StringPtr("high")
				updateLogpushJobOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := logpushJobsApiService.UpdateLogpushJob(updateLogpushJobOptionsModel)
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
	Describe(`DeleteLogpushJob(deleteLogpushJobOptions *DeleteLogpushJobOptions) - Operation response error`, func() {
		crn := "testString"
		zoneID := "testString"
		dataset := "http_requests"
		deleteLogpushJobPath := "/v1/testString/zones/testString/logpush/jobs/38"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteLogpushJobPath))
					Expect(req.Method).To(Equal("DELETE"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke DeleteLogpushJob with error: Operation response processing error`, func() {
				logpushJobsApiService, serviceErr := logpushjobsapiv1.NewLogpushJobsApiV1(&logpushjobsapiv1.LogpushJobsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ZoneID:        core.StringPtr(zoneID),
					Dataset:       core.StringPtr(dataset),
				})
				Expect(serviceErr).To(BeNil())
				Expect(logpushJobsApiService).ToNot(BeNil())

				// Construct an instance of the DeleteLogpushJobOptions model
				deleteLogpushJobOptionsModel := new(logpushjobsapiv1.DeleteLogpushJobOptions)
				deleteLogpushJobOptionsModel.JobID = core.Int64Ptr(int64(38))
				deleteLogpushJobOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := logpushJobsApiService.DeleteLogpushJob(deleteLogpushJobOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				logpushJobsApiService.EnableRetries(0, 0)
				result, response, operationErr = logpushJobsApiService.DeleteLogpushJob(deleteLogpushJobOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteLogpushJob(deleteLogpushJobOptions *DeleteLogpushJobOptions)`, func() {
		crn := "testString"
		zoneID := "testString"
		dataset := "http_requests"
		deleteLogpushJobPath := "/v1/testString/zones/testString/logpush/jobs/38"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteLogpushJobPath))
					Expect(req.Method).To(Equal("DELETE"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"anyKey": "anyValue"}}`)
				}))
			})
			It(`Invoke DeleteLogpushJob successfully with retries`, func() {
				logpushJobsApiService, serviceErr := logpushjobsapiv1.NewLogpushJobsApiV1(&logpushjobsapiv1.LogpushJobsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ZoneID:        core.StringPtr(zoneID),
					Dataset:       core.StringPtr(dataset),
				})
				Expect(serviceErr).To(BeNil())
				Expect(logpushJobsApiService).ToNot(BeNil())
				logpushJobsApiService.EnableRetries(0, 0)

				// Construct an instance of the DeleteLogpushJobOptions model
				deleteLogpushJobOptionsModel := new(logpushjobsapiv1.DeleteLogpushJobOptions)
				deleteLogpushJobOptionsModel.JobID = core.Int64Ptr(int64(38))
				deleteLogpushJobOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := logpushJobsApiService.DeleteLogpushJobWithContext(ctx, deleteLogpushJobOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				logpushJobsApiService.DisableRetries()
				result, response, operationErr := logpushJobsApiService.DeleteLogpushJob(deleteLogpushJobOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = logpushJobsApiService.DeleteLogpushJobWithContext(ctx, deleteLogpushJobOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(deleteLogpushJobPath))
					Expect(req.Method).To(Equal("DELETE"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"anyKey": "anyValue"}}`)
				}))
			})
			It(`Invoke DeleteLogpushJob successfully`, func() {
				logpushJobsApiService, serviceErr := logpushjobsapiv1.NewLogpushJobsApiV1(&logpushjobsapiv1.LogpushJobsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ZoneID:        core.StringPtr(zoneID),
					Dataset:       core.StringPtr(dataset),
				})
				Expect(serviceErr).To(BeNil())
				Expect(logpushJobsApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := logpushJobsApiService.DeleteLogpushJob(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DeleteLogpushJobOptions model
				deleteLogpushJobOptionsModel := new(logpushjobsapiv1.DeleteLogpushJobOptions)
				deleteLogpushJobOptionsModel.JobID = core.Int64Ptr(int64(38))
				deleteLogpushJobOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = logpushJobsApiService.DeleteLogpushJob(deleteLogpushJobOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke DeleteLogpushJob with error: Operation validation and request error`, func() {
				logpushJobsApiService, serviceErr := logpushjobsapiv1.NewLogpushJobsApiV1(&logpushjobsapiv1.LogpushJobsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ZoneID:        core.StringPtr(zoneID),
					Dataset:       core.StringPtr(dataset),
				})
				Expect(serviceErr).To(BeNil())
				Expect(logpushJobsApiService).ToNot(BeNil())

				// Construct an instance of the DeleteLogpushJobOptions model
				deleteLogpushJobOptionsModel := new(logpushjobsapiv1.DeleteLogpushJobOptions)
				deleteLogpushJobOptionsModel.JobID = core.Int64Ptr(int64(38))
				deleteLogpushJobOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := logpushJobsApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := logpushJobsApiService.DeleteLogpushJob(deleteLogpushJobOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the DeleteLogpushJobOptions model with no property values
				deleteLogpushJobOptionsModelNew := new(logpushjobsapiv1.DeleteLogpushJobOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = logpushJobsApiService.DeleteLogpushJob(deleteLogpushJobOptionsModelNew)
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
			It(`Invoke DeleteLogpushJob successfully`, func() {
				logpushJobsApiService, serviceErr := logpushjobsapiv1.NewLogpushJobsApiV1(&logpushjobsapiv1.LogpushJobsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ZoneID:        core.StringPtr(zoneID),
					Dataset:       core.StringPtr(dataset),
				})
				Expect(serviceErr).To(BeNil())
				Expect(logpushJobsApiService).ToNot(BeNil())

				// Construct an instance of the DeleteLogpushJobOptions model
				deleteLogpushJobOptionsModel := new(logpushjobsapiv1.DeleteLogpushJobOptions)
				deleteLogpushJobOptionsModel.JobID = core.Int64Ptr(int64(38))
				deleteLogpushJobOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := logpushJobsApiService.DeleteLogpushJob(deleteLogpushJobOptionsModel)
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
	Describe(`ListFieldsForDataset(listFieldsForDatasetOptions *ListFieldsForDatasetOptions) - Operation response error`, func() {
		crn := "testString"
		zoneID := "testString"
		dataset := "http_requests"
		listFieldsForDatasetPath := "/v1/testString/zones/testString/logpush/datasets/http_requests/fields"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listFieldsForDatasetPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListFieldsForDataset with error: Operation response processing error`, func() {
				logpushJobsApiService, serviceErr := logpushjobsapiv1.NewLogpushJobsApiV1(&logpushjobsapiv1.LogpushJobsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ZoneID:        core.StringPtr(zoneID),
					Dataset:       core.StringPtr(dataset),
				})
				Expect(serviceErr).To(BeNil())
				Expect(logpushJobsApiService).ToNot(BeNil())

				// Construct an instance of the ListFieldsForDatasetOptions model
				listFieldsForDatasetOptionsModel := new(logpushjobsapiv1.ListFieldsForDatasetOptions)
				listFieldsForDatasetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := logpushJobsApiService.ListFieldsForDataset(listFieldsForDatasetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				logpushJobsApiService.EnableRetries(0, 0)
				result, response, operationErr = logpushJobsApiService.ListFieldsForDataset(listFieldsForDatasetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListFieldsForDataset(listFieldsForDatasetOptions *ListFieldsForDatasetOptions)`, func() {
		crn := "testString"
		zoneID := "testString"
		dataset := "http_requests"
		listFieldsForDatasetPath := "/v1/testString/zones/testString/logpush/datasets/http_requests/fields"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listFieldsForDatasetPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"anyKey": "anyValue"}}`)
				}))
			})
			It(`Invoke ListFieldsForDataset successfully with retries`, func() {
				logpushJobsApiService, serviceErr := logpushjobsapiv1.NewLogpushJobsApiV1(&logpushjobsapiv1.LogpushJobsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ZoneID:        core.StringPtr(zoneID),
					Dataset:       core.StringPtr(dataset),
				})
				Expect(serviceErr).To(BeNil())
				Expect(logpushJobsApiService).ToNot(BeNil())
				logpushJobsApiService.EnableRetries(0, 0)

				// Construct an instance of the ListFieldsForDatasetOptions model
				listFieldsForDatasetOptionsModel := new(logpushjobsapiv1.ListFieldsForDatasetOptions)
				listFieldsForDatasetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := logpushJobsApiService.ListFieldsForDatasetWithContext(ctx, listFieldsForDatasetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				logpushJobsApiService.DisableRetries()
				result, response, operationErr := logpushJobsApiService.ListFieldsForDataset(listFieldsForDatasetOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = logpushJobsApiService.ListFieldsForDatasetWithContext(ctx, listFieldsForDatasetOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listFieldsForDatasetPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"anyKey": "anyValue"}}`)
				}))
			})
			It(`Invoke ListFieldsForDataset successfully`, func() {
				logpushJobsApiService, serviceErr := logpushjobsapiv1.NewLogpushJobsApiV1(&logpushjobsapiv1.LogpushJobsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ZoneID:        core.StringPtr(zoneID),
					Dataset:       core.StringPtr(dataset),
				})
				Expect(serviceErr).To(BeNil())
				Expect(logpushJobsApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := logpushJobsApiService.ListFieldsForDataset(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListFieldsForDatasetOptions model
				listFieldsForDatasetOptionsModel := new(logpushjobsapiv1.ListFieldsForDatasetOptions)
				listFieldsForDatasetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = logpushJobsApiService.ListFieldsForDataset(listFieldsForDatasetOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListFieldsForDataset with error: Operation request error`, func() {
				logpushJobsApiService, serviceErr := logpushjobsapiv1.NewLogpushJobsApiV1(&logpushjobsapiv1.LogpushJobsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ZoneID:        core.StringPtr(zoneID),
					Dataset:       core.StringPtr(dataset),
				})
				Expect(serviceErr).To(BeNil())
				Expect(logpushJobsApiService).ToNot(BeNil())

				// Construct an instance of the ListFieldsForDatasetOptions model
				listFieldsForDatasetOptionsModel := new(logpushjobsapiv1.ListFieldsForDatasetOptions)
				listFieldsForDatasetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := logpushJobsApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := logpushJobsApiService.ListFieldsForDataset(listFieldsForDatasetOptionsModel)
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
			It(`Invoke ListFieldsForDataset successfully`, func() {
				logpushJobsApiService, serviceErr := logpushjobsapiv1.NewLogpushJobsApiV1(&logpushjobsapiv1.LogpushJobsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ZoneID:        core.StringPtr(zoneID),
					Dataset:       core.StringPtr(dataset),
				})
				Expect(serviceErr).To(BeNil())
				Expect(logpushJobsApiService).ToNot(BeNil())

				// Construct an instance of the ListFieldsForDatasetOptions model
				listFieldsForDatasetOptionsModel := new(logpushjobsapiv1.ListFieldsForDatasetOptions)
				listFieldsForDatasetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := logpushJobsApiService.ListFieldsForDataset(listFieldsForDatasetOptionsModel)
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
	Describe(`ListLogpushJobsForDataset(listLogpushJobsForDatasetOptions *ListLogpushJobsForDatasetOptions) - Operation response error`, func() {
		crn := "testString"
		zoneID := "testString"
		dataset := "http_requests"
		listLogpushJobsForDatasetPath := "/v1/testString/zones/testString/logpush/datasets/http_requests/jobs"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listLogpushJobsForDatasetPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListLogpushJobsForDataset with error: Operation response processing error`, func() {
				logpushJobsApiService, serviceErr := logpushjobsapiv1.NewLogpushJobsApiV1(&logpushjobsapiv1.LogpushJobsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ZoneID:        core.StringPtr(zoneID),
					Dataset:       core.StringPtr(dataset),
				})
				Expect(serviceErr).To(BeNil())
				Expect(logpushJobsApiService).ToNot(BeNil())

				// Construct an instance of the ListLogpushJobsForDatasetOptions model
				listLogpushJobsForDatasetOptionsModel := new(logpushjobsapiv1.ListLogpushJobsForDatasetOptions)
				listLogpushJobsForDatasetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := logpushJobsApiService.ListLogpushJobsForDataset(listLogpushJobsForDatasetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				logpushJobsApiService.EnableRetries(0, 0)
				result, response, operationErr = logpushJobsApiService.ListLogpushJobsForDataset(listLogpushJobsForDatasetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListLogpushJobsForDataset(listLogpushJobsForDatasetOptions *ListLogpushJobsForDatasetOptions)`, func() {
		crn := "testString"
		zoneID := "testString"
		dataset := "http_requests"
		listLogpushJobsForDatasetPath := "/v1/testString/zones/testString/logpush/datasets/http_requests/jobs"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listLogpushJobsForDatasetPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": 5850, "name": "My log push job", "enabled": false, "dataset": "firewall_events", "frequency": "high", "logpull_options": "timestamps=rfc3339&timestamps=rfc3339", "destination_conf": "cos://cos-bucket001?region=us-south&instance-id=231f5467-3072-4cb9-9e39-a906fa3032ea", "last_complete": "2022-01-15T16:33:31.834209Z", "last_error": "2022-01-15T16:33:31.834209Z", "error_message": "ErrorMessage"}}`)
				}))
			})
			It(`Invoke ListLogpushJobsForDataset successfully with retries`, func() {
				logpushJobsApiService, serviceErr := logpushjobsapiv1.NewLogpushJobsApiV1(&logpushjobsapiv1.LogpushJobsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ZoneID:        core.StringPtr(zoneID),
					Dataset:       core.StringPtr(dataset),
				})
				Expect(serviceErr).To(BeNil())
				Expect(logpushJobsApiService).ToNot(BeNil())
				logpushJobsApiService.EnableRetries(0, 0)

				// Construct an instance of the ListLogpushJobsForDatasetOptions model
				listLogpushJobsForDatasetOptionsModel := new(logpushjobsapiv1.ListLogpushJobsForDatasetOptions)
				listLogpushJobsForDatasetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := logpushJobsApiService.ListLogpushJobsForDatasetWithContext(ctx, listLogpushJobsForDatasetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				logpushJobsApiService.DisableRetries()
				result, response, operationErr := logpushJobsApiService.ListLogpushJobsForDataset(listLogpushJobsForDatasetOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = logpushJobsApiService.ListLogpushJobsForDatasetWithContext(ctx, listLogpushJobsForDatasetOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listLogpushJobsForDatasetPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": 5850, "name": "My log push job", "enabled": false, "dataset": "firewall_events", "frequency": "high", "logpull_options": "timestamps=rfc3339&timestamps=rfc3339", "destination_conf": "cos://cos-bucket001?region=us-south&instance-id=231f5467-3072-4cb9-9e39-a906fa3032ea", "last_complete": "2022-01-15T16:33:31.834209Z", "last_error": "2022-01-15T16:33:31.834209Z", "error_message": "ErrorMessage"}}`)
				}))
			})
			It(`Invoke ListLogpushJobsForDataset successfully`, func() {
				logpushJobsApiService, serviceErr := logpushjobsapiv1.NewLogpushJobsApiV1(&logpushjobsapiv1.LogpushJobsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ZoneID:        core.StringPtr(zoneID),
					Dataset:       core.StringPtr(dataset),
				})
				Expect(serviceErr).To(BeNil())
				Expect(logpushJobsApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := logpushJobsApiService.ListLogpushJobsForDataset(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListLogpushJobsForDatasetOptions model
				listLogpushJobsForDatasetOptionsModel := new(logpushjobsapiv1.ListLogpushJobsForDatasetOptions)
				listLogpushJobsForDatasetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = logpushJobsApiService.ListLogpushJobsForDataset(listLogpushJobsForDatasetOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListLogpushJobsForDataset with error: Operation request error`, func() {
				logpushJobsApiService, serviceErr := logpushjobsapiv1.NewLogpushJobsApiV1(&logpushjobsapiv1.LogpushJobsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ZoneID:        core.StringPtr(zoneID),
					Dataset:       core.StringPtr(dataset),
				})
				Expect(serviceErr).To(BeNil())
				Expect(logpushJobsApiService).ToNot(BeNil())

				// Construct an instance of the ListLogpushJobsForDatasetOptions model
				listLogpushJobsForDatasetOptionsModel := new(logpushjobsapiv1.ListLogpushJobsForDatasetOptions)
				listLogpushJobsForDatasetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := logpushJobsApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := logpushJobsApiService.ListLogpushJobsForDataset(listLogpushJobsForDatasetOptionsModel)
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
			It(`Invoke ListLogpushJobsForDataset successfully`, func() {
				logpushJobsApiService, serviceErr := logpushjobsapiv1.NewLogpushJobsApiV1(&logpushjobsapiv1.LogpushJobsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ZoneID:        core.StringPtr(zoneID),
					Dataset:       core.StringPtr(dataset),
				})
				Expect(serviceErr).To(BeNil())
				Expect(logpushJobsApiService).ToNot(BeNil())

				// Construct an instance of the ListLogpushJobsForDatasetOptions model
				listLogpushJobsForDatasetOptionsModel := new(logpushjobsapiv1.ListLogpushJobsForDatasetOptions)
				listLogpushJobsForDatasetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := logpushJobsApiService.ListLogpushJobsForDataset(listLogpushJobsForDatasetOptionsModel)
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
	Describe(`GetLogpushOwnership(getLogpushOwnershipOptions *GetLogpushOwnershipOptions) - Operation response error`, func() {
		crn := "testString"
		zoneID := "testString"
		dataset := "http_requests"
		getLogpushOwnershipPath := "/v1/testString/zones/testString/logpush/ownership"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getLogpushOwnershipPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetLogpushOwnership with error: Operation response processing error`, func() {
				logpushJobsApiService, serviceErr := logpushjobsapiv1.NewLogpushJobsApiV1(&logpushjobsapiv1.LogpushJobsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ZoneID:        core.StringPtr(zoneID),
					Dataset:       core.StringPtr(dataset),
				})
				Expect(serviceErr).To(BeNil())
				Expect(logpushJobsApiService).ToNot(BeNil())

				// Construct an instance of the GetLogpushOwnershipOptions model
				getLogpushOwnershipOptionsModel := new(logpushjobsapiv1.GetLogpushOwnershipOptions)
				getLogpushOwnershipOptionsModel.DestinationConf = core.StringPtr("cos://cos-bucket001?region=us-south&instance-id=231f5467-3072-4cb9-9e39-a906fa3032ea")
				getLogpushOwnershipOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := logpushJobsApiService.GetLogpushOwnership(getLogpushOwnershipOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				logpushJobsApiService.EnableRetries(0, 0)
				result, response, operationErr = logpushJobsApiService.GetLogpushOwnership(getLogpushOwnershipOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetLogpushOwnership(getLogpushOwnershipOptions *GetLogpushOwnershipOptions)`, func() {
		crn := "testString"
		zoneID := "testString"
		dataset := "http_requests"
		getLogpushOwnershipPath := "/v1/testString/zones/testString/logpush/ownership"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getLogpushOwnershipPath))
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
			It(`Invoke GetLogpushOwnership successfully with retries`, func() {
				logpushJobsApiService, serviceErr := logpushjobsapiv1.NewLogpushJobsApiV1(&logpushjobsapiv1.LogpushJobsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ZoneID:        core.StringPtr(zoneID),
					Dataset:       core.StringPtr(dataset),
				})
				Expect(serviceErr).To(BeNil())
				Expect(logpushJobsApiService).ToNot(BeNil())
				logpushJobsApiService.EnableRetries(0, 0)

				// Construct an instance of the GetLogpushOwnershipOptions model
				getLogpushOwnershipOptionsModel := new(logpushjobsapiv1.GetLogpushOwnershipOptions)
				getLogpushOwnershipOptionsModel.DestinationConf = core.StringPtr("cos://cos-bucket001?region=us-south&instance-id=231f5467-3072-4cb9-9e39-a906fa3032ea")
				getLogpushOwnershipOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := logpushJobsApiService.GetLogpushOwnershipWithContext(ctx, getLogpushOwnershipOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				logpushJobsApiService.DisableRetries()
				result, response, operationErr := logpushJobsApiService.GetLogpushOwnership(getLogpushOwnershipOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = logpushJobsApiService.GetLogpushOwnershipWithContext(ctx, getLogpushOwnershipOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getLogpushOwnershipPath))
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
			It(`Invoke GetLogpushOwnership successfully`, func() {
				logpushJobsApiService, serviceErr := logpushjobsapiv1.NewLogpushJobsApiV1(&logpushjobsapiv1.LogpushJobsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ZoneID:        core.StringPtr(zoneID),
					Dataset:       core.StringPtr(dataset),
				})
				Expect(serviceErr).To(BeNil())
				Expect(logpushJobsApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := logpushJobsApiService.GetLogpushOwnership(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetLogpushOwnershipOptions model
				getLogpushOwnershipOptionsModel := new(logpushjobsapiv1.GetLogpushOwnershipOptions)
				getLogpushOwnershipOptionsModel.DestinationConf = core.StringPtr("cos://cos-bucket001?region=us-south&instance-id=231f5467-3072-4cb9-9e39-a906fa3032ea")
				getLogpushOwnershipOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = logpushJobsApiService.GetLogpushOwnership(getLogpushOwnershipOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetLogpushOwnership with error: Operation request error`, func() {
				logpushJobsApiService, serviceErr := logpushjobsapiv1.NewLogpushJobsApiV1(&logpushjobsapiv1.LogpushJobsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ZoneID:        core.StringPtr(zoneID),
					Dataset:       core.StringPtr(dataset),
				})
				Expect(serviceErr).To(BeNil())
				Expect(logpushJobsApiService).ToNot(BeNil())

				// Construct an instance of the GetLogpushOwnershipOptions model
				getLogpushOwnershipOptionsModel := new(logpushjobsapiv1.GetLogpushOwnershipOptions)
				getLogpushOwnershipOptionsModel.DestinationConf = core.StringPtr("cos://cos-bucket001?region=us-south&instance-id=231f5467-3072-4cb9-9e39-a906fa3032ea")
				getLogpushOwnershipOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := logpushJobsApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := logpushJobsApiService.GetLogpushOwnership(getLogpushOwnershipOptionsModel)
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
			It(`Invoke GetLogpushOwnership successfully`, func() {
				logpushJobsApiService, serviceErr := logpushjobsapiv1.NewLogpushJobsApiV1(&logpushjobsapiv1.LogpushJobsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ZoneID:        core.StringPtr(zoneID),
					Dataset:       core.StringPtr(dataset),
				})
				Expect(serviceErr).To(BeNil())
				Expect(logpushJobsApiService).ToNot(BeNil())

				// Construct an instance of the GetLogpushOwnershipOptions model
				getLogpushOwnershipOptionsModel := new(logpushjobsapiv1.GetLogpushOwnershipOptions)
				getLogpushOwnershipOptionsModel.DestinationConf = core.StringPtr("cos://cos-bucket001?region=us-south&instance-id=231f5467-3072-4cb9-9e39-a906fa3032ea")
				getLogpushOwnershipOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := logpushJobsApiService.GetLogpushOwnership(getLogpushOwnershipOptionsModel)
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
	Describe(`ValidateLogpushOwnershipChallenge(validateLogpushOwnershipChallengeOptions *ValidateLogpushOwnershipChallengeOptions) - Operation response error`, func() {
		crn := "testString"
		zoneID := "testString"
		dataset := "http_requests"
		validateLogpushOwnershipChallengePath := "/v1/testString/zones/testString/logpush/ownership/validate"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(validateLogpushOwnershipChallengePath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ValidateLogpushOwnershipChallenge with error: Operation response processing error`, func() {
				logpushJobsApiService, serviceErr := logpushjobsapiv1.NewLogpushJobsApiV1(&logpushjobsapiv1.LogpushJobsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ZoneID:        core.StringPtr(zoneID),
					Dataset:       core.StringPtr(dataset),
				})
				Expect(serviceErr).To(BeNil())
				Expect(logpushJobsApiService).ToNot(BeNil())

				// Construct an instance of the ValidateLogpushOwnershipChallengeOptions model
				validateLogpushOwnershipChallengeOptionsModel := new(logpushjobsapiv1.ValidateLogpushOwnershipChallengeOptions)
				validateLogpushOwnershipChallengeOptionsModel.DestinationConf = core.StringPtr("cos://cos-bucket001?region=us-south&instance-id=231f5467-3072-4cb9-9e39-a906fa3032ea")
				validateLogpushOwnershipChallengeOptionsModel.OwnershipChallenge = core.StringPtr("00000000000000000000")
				validateLogpushOwnershipChallengeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := logpushJobsApiService.ValidateLogpushOwnershipChallenge(validateLogpushOwnershipChallengeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				logpushJobsApiService.EnableRetries(0, 0)
				result, response, operationErr = logpushJobsApiService.ValidateLogpushOwnershipChallenge(validateLogpushOwnershipChallengeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ValidateLogpushOwnershipChallenge(validateLogpushOwnershipChallengeOptions *ValidateLogpushOwnershipChallengeOptions)`, func() {
		crn := "testString"
		zoneID := "testString"
		dataset := "http_requests"
		validateLogpushOwnershipChallengePath := "/v1/testString/zones/testString/logpush/ownership/validate"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(validateLogpushOwnershipChallengePath))
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
			It(`Invoke ValidateLogpushOwnershipChallenge successfully with retries`, func() {
				logpushJobsApiService, serviceErr := logpushjobsapiv1.NewLogpushJobsApiV1(&logpushjobsapiv1.LogpushJobsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ZoneID:        core.StringPtr(zoneID),
					Dataset:       core.StringPtr(dataset),
				})
				Expect(serviceErr).To(BeNil())
				Expect(logpushJobsApiService).ToNot(BeNil())
				logpushJobsApiService.EnableRetries(0, 0)

				// Construct an instance of the ValidateLogpushOwnershipChallengeOptions model
				validateLogpushOwnershipChallengeOptionsModel := new(logpushjobsapiv1.ValidateLogpushOwnershipChallengeOptions)
				validateLogpushOwnershipChallengeOptionsModel.DestinationConf = core.StringPtr("cos://cos-bucket001?region=us-south&instance-id=231f5467-3072-4cb9-9e39-a906fa3032ea")
				validateLogpushOwnershipChallengeOptionsModel.OwnershipChallenge = core.StringPtr("00000000000000000000")
				validateLogpushOwnershipChallengeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := logpushJobsApiService.ValidateLogpushOwnershipChallengeWithContext(ctx, validateLogpushOwnershipChallengeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				logpushJobsApiService.DisableRetries()
				result, response, operationErr := logpushJobsApiService.ValidateLogpushOwnershipChallenge(validateLogpushOwnershipChallengeOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = logpushJobsApiService.ValidateLogpushOwnershipChallengeWithContext(ctx, validateLogpushOwnershipChallengeOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(validateLogpushOwnershipChallengePath))
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
			It(`Invoke ValidateLogpushOwnershipChallenge successfully`, func() {
				logpushJobsApiService, serviceErr := logpushjobsapiv1.NewLogpushJobsApiV1(&logpushjobsapiv1.LogpushJobsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ZoneID:        core.StringPtr(zoneID),
					Dataset:       core.StringPtr(dataset),
				})
				Expect(serviceErr).To(BeNil())
				Expect(logpushJobsApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := logpushJobsApiService.ValidateLogpushOwnershipChallenge(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ValidateLogpushOwnershipChallengeOptions model
				validateLogpushOwnershipChallengeOptionsModel := new(logpushjobsapiv1.ValidateLogpushOwnershipChallengeOptions)
				validateLogpushOwnershipChallengeOptionsModel.DestinationConf = core.StringPtr("cos://cos-bucket001?region=us-south&instance-id=231f5467-3072-4cb9-9e39-a906fa3032ea")
				validateLogpushOwnershipChallengeOptionsModel.OwnershipChallenge = core.StringPtr("00000000000000000000")
				validateLogpushOwnershipChallengeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = logpushJobsApiService.ValidateLogpushOwnershipChallenge(validateLogpushOwnershipChallengeOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ValidateLogpushOwnershipChallenge with error: Operation request error`, func() {
				logpushJobsApiService, serviceErr := logpushjobsapiv1.NewLogpushJobsApiV1(&logpushjobsapiv1.LogpushJobsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ZoneID:        core.StringPtr(zoneID),
					Dataset:       core.StringPtr(dataset),
				})
				Expect(serviceErr).To(BeNil())
				Expect(logpushJobsApiService).ToNot(BeNil())

				// Construct an instance of the ValidateLogpushOwnershipChallengeOptions model
				validateLogpushOwnershipChallengeOptionsModel := new(logpushjobsapiv1.ValidateLogpushOwnershipChallengeOptions)
				validateLogpushOwnershipChallengeOptionsModel.DestinationConf = core.StringPtr("cos://cos-bucket001?region=us-south&instance-id=231f5467-3072-4cb9-9e39-a906fa3032ea")
				validateLogpushOwnershipChallengeOptionsModel.OwnershipChallenge = core.StringPtr("00000000000000000000")
				validateLogpushOwnershipChallengeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := logpushJobsApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := logpushJobsApiService.ValidateLogpushOwnershipChallenge(validateLogpushOwnershipChallengeOptionsModel)
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
			It(`Invoke ValidateLogpushOwnershipChallenge successfully`, func() {
				logpushJobsApiService, serviceErr := logpushjobsapiv1.NewLogpushJobsApiV1(&logpushjobsapiv1.LogpushJobsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ZoneID:        core.StringPtr(zoneID),
					Dataset:       core.StringPtr(dataset),
				})
				Expect(serviceErr).To(BeNil())
				Expect(logpushJobsApiService).ToNot(BeNil())

				// Construct an instance of the ValidateLogpushOwnershipChallengeOptions model
				validateLogpushOwnershipChallengeOptionsModel := new(logpushjobsapiv1.ValidateLogpushOwnershipChallengeOptions)
				validateLogpushOwnershipChallengeOptionsModel.DestinationConf = core.StringPtr("cos://cos-bucket001?region=us-south&instance-id=231f5467-3072-4cb9-9e39-a906fa3032ea")
				validateLogpushOwnershipChallengeOptionsModel.OwnershipChallenge = core.StringPtr("00000000000000000000")
				validateLogpushOwnershipChallengeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := logpushJobsApiService.ValidateLogpushOwnershipChallenge(validateLogpushOwnershipChallengeOptionsModel)
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
	Describe(`GetLogpushJobsV2(getLogpushJobsV2Options *GetLogpushJobsV2Options) - Operation response error`, func() {
		crn := "testString"
		zoneID := "testString"
		dataset := "http_requests"
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
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetLogpushJobsV2 with error: Operation response processing error`, func() {
				logpushJobsApiService, serviceErr := logpushjobsapiv1.NewLogpushJobsApiV1(&logpushjobsapiv1.LogpushJobsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ZoneID:        core.StringPtr(zoneID),
					Dataset:       core.StringPtr(dataset),
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
		zoneID := "testString"
		dataset := "http_requests"
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
					ZoneID:        core.StringPtr(zoneID),
					Dataset:       core.StringPtr(dataset),
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
					ZoneID:        core.StringPtr(zoneID),
					Dataset:       core.StringPtr(dataset),
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
					ZoneID:        core.StringPtr(zoneID),
					Dataset:       core.StringPtr(dataset),
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
					ZoneID:        core.StringPtr(zoneID),
					Dataset:       core.StringPtr(dataset),
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
		zoneID := "testString"
		dataset := "http_requests"
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
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateLogpushJobV2 with error: Operation response processing error`, func() {
				logpushJobsApiService, serviceErr := logpushjobsapiv1.NewLogpushJobsApiV1(&logpushjobsapiv1.LogpushJobsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ZoneID:        core.StringPtr(zoneID),
					Dataset:       core.StringPtr(dataset),
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
		zoneID := "testString"
		dataset := "http_requests"
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
					ZoneID:        core.StringPtr(zoneID),
					Dataset:       core.StringPtr(dataset),
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
					ZoneID:        core.StringPtr(zoneID),
					Dataset:       core.StringPtr(dataset),
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
					ZoneID:        core.StringPtr(zoneID),
					Dataset:       core.StringPtr(dataset),
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
					ZoneID:        core.StringPtr(zoneID),
					Dataset:       core.StringPtr(dataset),
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
		zoneID := "testString"
		dataset := "http_requests"
		getLogpushJobV2Path := "/v2/testString/zones/testString/logpush/jobs/38"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getLogpushJobV2Path))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetLogpushJobV2 with error: Operation response processing error`, func() {
				logpushJobsApiService, serviceErr := logpushjobsapiv1.NewLogpushJobsApiV1(&logpushjobsapiv1.LogpushJobsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ZoneID:        core.StringPtr(zoneID),
					Dataset:       core.StringPtr(dataset),
				})
				Expect(serviceErr).To(BeNil())
				Expect(logpushJobsApiService).ToNot(BeNil())

				// Construct an instance of the GetLogpushJobV2Options model
				getLogpushJobV2OptionsModel := new(logpushjobsapiv1.GetLogpushJobV2Options)
				getLogpushJobV2OptionsModel.JobID = core.Int64Ptr(int64(38))
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
		zoneID := "testString"
		dataset := "http_requests"
		getLogpushJobV2Path := "/v2/testString/zones/testString/logpush/jobs/38"
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
					ZoneID:        core.StringPtr(zoneID),
					Dataset:       core.StringPtr(dataset),
				})
				Expect(serviceErr).To(BeNil())
				Expect(logpushJobsApiService).ToNot(BeNil())
				logpushJobsApiService.EnableRetries(0, 0)

				// Construct an instance of the GetLogpushJobV2Options model
				getLogpushJobV2OptionsModel := new(logpushjobsapiv1.GetLogpushJobV2Options)
				getLogpushJobV2OptionsModel.JobID = core.Int64Ptr(int64(38))
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
					ZoneID:        core.StringPtr(zoneID),
					Dataset:       core.StringPtr(dataset),
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
				getLogpushJobV2OptionsModel.JobID = core.Int64Ptr(int64(38))
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
					ZoneID:        core.StringPtr(zoneID),
					Dataset:       core.StringPtr(dataset),
				})
				Expect(serviceErr).To(BeNil())
				Expect(logpushJobsApiService).ToNot(BeNil())

				// Construct an instance of the GetLogpushJobV2Options model
				getLogpushJobV2OptionsModel := new(logpushjobsapiv1.GetLogpushJobV2Options)
				getLogpushJobV2OptionsModel.JobID = core.Int64Ptr(int64(38))
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
					ZoneID:        core.StringPtr(zoneID),
					Dataset:       core.StringPtr(dataset),
				})
				Expect(serviceErr).To(BeNil())
				Expect(logpushJobsApiService).ToNot(BeNil())

				// Construct an instance of the GetLogpushJobV2Options model
				getLogpushJobV2OptionsModel := new(logpushjobsapiv1.GetLogpushJobV2Options)
				getLogpushJobV2OptionsModel.JobID = core.Int64Ptr(int64(38))
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
		zoneID := "testString"
		dataset := "http_requests"
		updateLogpushJobV2Path := "/v2/testString/zones/testString/logpush/jobs/38"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateLogpushJobV2Path))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateLogpushJobV2 with error: Operation response processing error`, func() {
				logpushJobsApiService, serviceErr := logpushjobsapiv1.NewLogpushJobsApiV1(&logpushjobsapiv1.LogpushJobsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ZoneID:        core.StringPtr(zoneID),
					Dataset:       core.StringPtr(dataset),
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
				updateLogpushJobV2OptionsModel.JobID = core.Int64Ptr(int64(38))
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
		zoneID := "testString"
		dataset := "http_requests"
		updateLogpushJobV2Path := "/v2/testString/zones/testString/logpush/jobs/38"
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
					ZoneID:        core.StringPtr(zoneID),
					Dataset:       core.StringPtr(dataset),
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
				updateLogpushJobV2OptionsModel.JobID = core.Int64Ptr(int64(38))
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
					ZoneID:        core.StringPtr(zoneID),
					Dataset:       core.StringPtr(dataset),
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
				updateLogpushJobV2OptionsModel.JobID = core.Int64Ptr(int64(38))
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
					ZoneID:        core.StringPtr(zoneID),
					Dataset:       core.StringPtr(dataset),
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
				updateLogpushJobV2OptionsModel.JobID = core.Int64Ptr(int64(38))
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
					ZoneID:        core.StringPtr(zoneID),
					Dataset:       core.StringPtr(dataset),
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
				updateLogpushJobV2OptionsModel.JobID = core.Int64Ptr(int64(38))
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
		zoneID := "testString"
		dataset := "http_requests"
		deleteLogpushJobV2Path := "/v2/testString/zones/testString/logpush/jobs/38"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteLogpushJobV2Path))
					Expect(req.Method).To(Equal("DELETE"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke DeleteLogpushJobV2 with error: Operation response processing error`, func() {
				logpushJobsApiService, serviceErr := logpushjobsapiv1.NewLogpushJobsApiV1(&logpushjobsapiv1.LogpushJobsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ZoneID:        core.StringPtr(zoneID),
					Dataset:       core.StringPtr(dataset),
				})
				Expect(serviceErr).To(BeNil())
				Expect(logpushJobsApiService).ToNot(BeNil())

				// Construct an instance of the DeleteLogpushJobV2Options model
				deleteLogpushJobV2OptionsModel := new(logpushjobsapiv1.DeleteLogpushJobV2Options)
				deleteLogpushJobV2OptionsModel.JobID = core.Int64Ptr(int64(38))
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
		zoneID := "testString"
		dataset := "http_requests"
		deleteLogpushJobV2Path := "/v2/testString/zones/testString/logpush/jobs/38"
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
					ZoneID:        core.StringPtr(zoneID),
					Dataset:       core.StringPtr(dataset),
				})
				Expect(serviceErr).To(BeNil())
				Expect(logpushJobsApiService).ToNot(BeNil())
				logpushJobsApiService.EnableRetries(0, 0)

				// Construct an instance of the DeleteLogpushJobV2Options model
				deleteLogpushJobV2OptionsModel := new(logpushjobsapiv1.DeleteLogpushJobV2Options)
				deleteLogpushJobV2OptionsModel.JobID = core.Int64Ptr(int64(38))
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
					ZoneID:        core.StringPtr(zoneID),
					Dataset:       core.StringPtr(dataset),
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
				deleteLogpushJobV2OptionsModel.JobID = core.Int64Ptr(int64(38))
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
					ZoneID:        core.StringPtr(zoneID),
					Dataset:       core.StringPtr(dataset),
				})
				Expect(serviceErr).To(BeNil())
				Expect(logpushJobsApiService).ToNot(BeNil())

				// Construct an instance of the DeleteLogpushJobV2Options model
				deleteLogpushJobV2OptionsModel := new(logpushjobsapiv1.DeleteLogpushJobV2Options)
				deleteLogpushJobV2OptionsModel.JobID = core.Int64Ptr(int64(38))
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
					ZoneID:        core.StringPtr(zoneID),
					Dataset:       core.StringPtr(dataset),
				})
				Expect(serviceErr).To(BeNil())
				Expect(logpushJobsApiService).ToNot(BeNil())

				// Construct an instance of the DeleteLogpushJobV2Options model
				deleteLogpushJobV2OptionsModel := new(logpushjobsapiv1.DeleteLogpushJobV2Options)
				deleteLogpushJobV2OptionsModel.JobID = core.Int64Ptr(int64(38))
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
		zoneID := "testString"
		dataset := "http_requests"
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
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetLogpushOwnershipV2 with error: Operation response processing error`, func() {
				logpushJobsApiService, serviceErr := logpushjobsapiv1.NewLogpushJobsApiV1(&logpushjobsapiv1.LogpushJobsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ZoneID:        core.StringPtr(zoneID),
					Dataset:       core.StringPtr(dataset),
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
		zoneID := "testString"
		dataset := "http_requests"
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
					ZoneID:        core.StringPtr(zoneID),
					Dataset:       core.StringPtr(dataset),
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
					ZoneID:        core.StringPtr(zoneID),
					Dataset:       core.StringPtr(dataset),
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
					ZoneID:        core.StringPtr(zoneID),
					Dataset:       core.StringPtr(dataset),
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
					ZoneID:        core.StringPtr(zoneID),
					Dataset:       core.StringPtr(dataset),
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
		zoneID := "testString"
		dataset := "http_requests"
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
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ValidateLogpushOwnershipChallengeV2 with error: Operation response processing error`, func() {
				logpushJobsApiService, serviceErr := logpushjobsapiv1.NewLogpushJobsApiV1(&logpushjobsapiv1.LogpushJobsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ZoneID:        core.StringPtr(zoneID),
					Dataset:       core.StringPtr(dataset),
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
		zoneID := "testString"
		dataset := "http_requests"
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
					ZoneID:        core.StringPtr(zoneID),
					Dataset:       core.StringPtr(dataset),
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
					ZoneID:        core.StringPtr(zoneID),
					Dataset:       core.StringPtr(dataset),
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
					ZoneID:        core.StringPtr(zoneID),
					Dataset:       core.StringPtr(dataset),
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
					ZoneID:        core.StringPtr(zoneID),
					Dataset:       core.StringPtr(dataset),
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
		zoneID := "testString"
		dataset := "http_requests"
		listFieldsForDatasetV2Path := "/v2/testString/zones/testString/logpush/datasets/http_requests/fields"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listFieldsForDatasetV2Path))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListFieldsForDatasetV2 with error: Operation response processing error`, func() {
				logpushJobsApiService, serviceErr := logpushjobsapiv1.NewLogpushJobsApiV1(&logpushjobsapiv1.LogpushJobsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ZoneID:        core.StringPtr(zoneID),
					Dataset:       core.StringPtr(dataset),
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
		zoneID := "testString"
		dataset := "http_requests"
		listFieldsForDatasetV2Path := "/v2/testString/zones/testString/logpush/datasets/http_requests/fields"
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
					ZoneID:        core.StringPtr(zoneID),
					Dataset:       core.StringPtr(dataset),
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
					ZoneID:        core.StringPtr(zoneID),
					Dataset:       core.StringPtr(dataset),
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
					ZoneID:        core.StringPtr(zoneID),
					Dataset:       core.StringPtr(dataset),
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
					ZoneID:        core.StringPtr(zoneID),
					Dataset:       core.StringPtr(dataset),
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
		zoneID := "testString"
		dataset := "http_requests"
		listLogpushJobsForDatasetV2Path := "/v2/testString/zones/testString/logpush/datasets/http_requests/jobs"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listLogpushJobsForDatasetV2Path))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListLogpushJobsForDatasetV2 with error: Operation response processing error`, func() {
				logpushJobsApiService, serviceErr := logpushjobsapiv1.NewLogpushJobsApiV1(&logpushjobsapiv1.LogpushJobsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ZoneID:        core.StringPtr(zoneID),
					Dataset:       core.StringPtr(dataset),
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
		zoneID := "testString"
		dataset := "http_requests"
		listLogpushJobsForDatasetV2Path := "/v2/testString/zones/testString/logpush/datasets/http_requests/jobs"
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
					ZoneID:        core.StringPtr(zoneID),
					Dataset:       core.StringPtr(dataset),
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
					ZoneID:        core.StringPtr(zoneID),
					Dataset:       core.StringPtr(dataset),
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
					ZoneID:        core.StringPtr(zoneID),
					Dataset:       core.StringPtr(dataset),
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
					ZoneID:        core.StringPtr(zoneID),
					Dataset:       core.StringPtr(dataset),
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
	Describe(`Model constructor tests`, func() {
		Context(`Using a service client instance`, func() {
			crn := "testString"
			zoneID := "testString"
			dataset := "http_requests"
			logpushJobsApiService, _ := logpushjobsapiv1.NewLogpushJobsApiV1(&logpushjobsapiv1.LogpushJobsApiV1Options{
				URL:           "http://logpushjobsapiv1modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
				Crn:           core.StringPtr(crn),
				ZoneID:        core.StringPtr(zoneID),
				Dataset:       core.StringPtr(dataset),
			})
			It(`Invoke NewCreateLogpushJobOptions successfully`, func() {
				// Construct an instance of the CreateLogpushJobOptions model
				createLogpushJobOptionsModel := logpushJobsApiService.NewCreateLogpushJobOptions()
				createLogpushJobOptionsModel.SetName("My log push job")
				createLogpushJobOptionsModel.SetEnabled(false)
				createLogpushJobOptionsModel.SetLogpullOptions("timestamps=rfc3339&timestamps=rfc3339")
				createLogpushJobOptionsModel.SetDestinationConf("cos://cos-bucket001?region=us-south&instance-id=231f5467-3072-4cb9-9e39-a906fa3032ea")
				createLogpushJobOptionsModel.SetOwnershipChallenge("00000000000000000000000000000000")
				createLogpushJobOptionsModel.SetDataset("firewall_events")
				createLogpushJobOptionsModel.SetFrequency("high")
				createLogpushJobOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createLogpushJobOptionsModel).ToNot(BeNil())
				Expect(createLogpushJobOptionsModel.Name).To(Equal(core.StringPtr("My log push job")))
				Expect(createLogpushJobOptionsModel.Enabled).To(Equal(core.BoolPtr(false)))
				Expect(createLogpushJobOptionsModel.LogpullOptions).To(Equal(core.StringPtr("timestamps=rfc3339&timestamps=rfc3339")))
				Expect(createLogpushJobOptionsModel.DestinationConf).To(Equal(core.StringPtr("cos://cos-bucket001?region=us-south&instance-id=231f5467-3072-4cb9-9e39-a906fa3032ea")))
				Expect(createLogpushJobOptionsModel.OwnershipChallenge).To(Equal(core.StringPtr("00000000000000000000000000000000")))
				Expect(createLogpushJobOptionsModel.Dataset).To(Equal(core.StringPtr("firewall_events")))
				Expect(createLogpushJobOptionsModel.Frequency).To(Equal(core.StringPtr("high")))
				Expect(createLogpushJobOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
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
			It(`Invoke NewDeleteLogpushJobOptions successfully`, func() {
				// Construct an instance of the DeleteLogpushJobOptions model
				jobID := int64(38)
				deleteLogpushJobOptionsModel := logpushJobsApiService.NewDeleteLogpushJobOptions(jobID)
				deleteLogpushJobOptionsModel.SetJobID(int64(38))
				deleteLogpushJobOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteLogpushJobOptionsModel).ToNot(BeNil())
				Expect(deleteLogpushJobOptionsModel.JobID).To(Equal(core.Int64Ptr(int64(38))))
				Expect(deleteLogpushJobOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteLogpushJobV2Options successfully`, func() {
				// Construct an instance of the DeleteLogpushJobV2Options model
				jobID := int64(38)
				deleteLogpushJobV2OptionsModel := logpushJobsApiService.NewDeleteLogpushJobV2Options(jobID)
				deleteLogpushJobV2OptionsModel.SetJobID(int64(38))
				deleteLogpushJobV2OptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteLogpushJobV2OptionsModel).ToNot(BeNil())
				Expect(deleteLogpushJobV2OptionsModel.JobID).To(Equal(core.Int64Ptr(int64(38))))
				Expect(deleteLogpushJobV2OptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetLogpushJobOptions successfully`, func() {
				// Construct an instance of the GetLogpushJobOptions model
				jobID := int64(38)
				getLogpushJobOptionsModel := logpushJobsApiService.NewGetLogpushJobOptions(jobID)
				getLogpushJobOptionsModel.SetJobID(int64(38))
				getLogpushJobOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getLogpushJobOptionsModel).ToNot(BeNil())
				Expect(getLogpushJobOptionsModel.JobID).To(Equal(core.Int64Ptr(int64(38))))
				Expect(getLogpushJobOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetLogpushJobV2Options successfully`, func() {
				// Construct an instance of the GetLogpushJobV2Options model
				jobID := int64(38)
				getLogpushJobV2OptionsModel := logpushJobsApiService.NewGetLogpushJobV2Options(jobID)
				getLogpushJobV2OptionsModel.SetJobID(int64(38))
				getLogpushJobV2OptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getLogpushJobV2OptionsModel).ToNot(BeNil())
				Expect(getLogpushJobV2OptionsModel.JobID).To(Equal(core.Int64Ptr(int64(38))))
				Expect(getLogpushJobV2OptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetLogpushJobsOptions successfully`, func() {
				// Construct an instance of the GetLogpushJobsOptions model
				getLogpushJobsOptionsModel := logpushJobsApiService.NewGetLogpushJobsOptions()
				getLogpushJobsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getLogpushJobsOptionsModel).ToNot(BeNil())
				Expect(getLogpushJobsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetLogpushJobsV2Options successfully`, func() {
				// Construct an instance of the GetLogpushJobsV2Options model
				getLogpushJobsV2OptionsModel := logpushJobsApiService.NewGetLogpushJobsV2Options()
				getLogpushJobsV2OptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getLogpushJobsV2OptionsModel).ToNot(BeNil())
				Expect(getLogpushJobsV2OptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetLogpushOwnershipOptions successfully`, func() {
				// Construct an instance of the GetLogpushOwnershipOptions model
				getLogpushOwnershipOptionsModel := logpushJobsApiService.NewGetLogpushOwnershipOptions()
				getLogpushOwnershipOptionsModel.SetDestinationConf("cos://cos-bucket001?region=us-south&instance-id=231f5467-3072-4cb9-9e39-a906fa3032ea")
				getLogpushOwnershipOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getLogpushOwnershipOptionsModel).ToNot(BeNil())
				Expect(getLogpushOwnershipOptionsModel.DestinationConf).To(Equal(core.StringPtr("cos://cos-bucket001?region=us-south&instance-id=231f5467-3072-4cb9-9e39-a906fa3032ea")))
				Expect(getLogpushOwnershipOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
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
			It(`Invoke NewListFieldsForDatasetOptions successfully`, func() {
				// Construct an instance of the ListFieldsForDatasetOptions model
				listFieldsForDatasetOptionsModel := logpushJobsApiService.NewListFieldsForDatasetOptions()
				listFieldsForDatasetOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listFieldsForDatasetOptionsModel).ToNot(BeNil())
				Expect(listFieldsForDatasetOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListFieldsForDatasetV2Options successfully`, func() {
				// Construct an instance of the ListFieldsForDatasetV2Options model
				listFieldsForDatasetV2OptionsModel := logpushJobsApiService.NewListFieldsForDatasetV2Options()
				listFieldsForDatasetV2OptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listFieldsForDatasetV2OptionsModel).ToNot(BeNil())
				Expect(listFieldsForDatasetV2OptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListLogpushJobsForDatasetOptions successfully`, func() {
				// Construct an instance of the ListLogpushJobsForDatasetOptions model
				listLogpushJobsForDatasetOptionsModel := logpushJobsApiService.NewListLogpushJobsForDatasetOptions()
				listLogpushJobsForDatasetOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listLogpushJobsForDatasetOptionsModel).ToNot(BeNil())
				Expect(listLogpushJobsForDatasetOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListLogpushJobsForDatasetV2Options successfully`, func() {
				// Construct an instance of the ListLogpushJobsForDatasetV2Options model
				listLogpushJobsForDatasetV2OptionsModel := logpushJobsApiService.NewListLogpushJobsForDatasetV2Options()
				listLogpushJobsForDatasetV2OptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listLogpushJobsForDatasetV2OptionsModel).ToNot(BeNil())
				Expect(listLogpushJobsForDatasetV2OptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateLogpushJobOptions successfully`, func() {
				// Construct an instance of the UpdateLogpushJobOptions model
				jobID := int64(38)
				updateLogpushJobOptionsModel := logpushJobsApiService.NewUpdateLogpushJobOptions(jobID)
				updateLogpushJobOptionsModel.SetJobID(int64(38))
				updateLogpushJobOptionsModel.SetEnabled(false)
				updateLogpushJobOptionsModel.SetLogpullOptions("timestamps=rfc3339&timestamps=rfc3339")
				updateLogpushJobOptionsModel.SetDestinationConf("cos://cos-bucket001?region=us-south&instance-id=231f5467-3072-4cb9-9e39-a906fa3032ea")
				updateLogpushJobOptionsModel.SetOwnershipChallenge("00000000000000000000000000000000")
				updateLogpushJobOptionsModel.SetFrequency("high")
				updateLogpushJobOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateLogpushJobOptionsModel).ToNot(BeNil())
				Expect(updateLogpushJobOptionsModel.JobID).To(Equal(core.Int64Ptr(int64(38))))
				Expect(updateLogpushJobOptionsModel.Enabled).To(Equal(core.BoolPtr(false)))
				Expect(updateLogpushJobOptionsModel.LogpullOptions).To(Equal(core.StringPtr("timestamps=rfc3339&timestamps=rfc3339")))
				Expect(updateLogpushJobOptionsModel.DestinationConf).To(Equal(core.StringPtr("cos://cos-bucket001?region=us-south&instance-id=231f5467-3072-4cb9-9e39-a906fa3032ea")))
				Expect(updateLogpushJobOptionsModel.OwnershipChallenge).To(Equal(core.StringPtr("00000000000000000000000000000000")))
				Expect(updateLogpushJobOptionsModel.Frequency).To(Equal(core.StringPtr("high")))
				Expect(updateLogpushJobOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
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
				jobID := int64(38)
				updateLogpushJobV2OptionsModel := logpushJobsApiService.NewUpdateLogpushJobV2Options(jobID)
				updateLogpushJobV2OptionsModel.SetJobID(int64(38))
				updateLogpushJobV2OptionsModel.SetUpdateLogpushJobV2Request(updateLogpushJobV2RequestModel)
				updateLogpushJobV2OptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateLogpushJobV2OptionsModel).ToNot(BeNil())
				Expect(updateLogpushJobV2OptionsModel.JobID).To(Equal(core.Int64Ptr(int64(38))))
				Expect(updateLogpushJobV2OptionsModel.UpdateLogpushJobV2Request).To(Equal(updateLogpushJobV2RequestModel))
				Expect(updateLogpushJobV2OptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewValidateLogpushOwnershipChallengeOptions successfully`, func() {
				// Construct an instance of the ValidateLogpushOwnershipChallengeOptions model
				validateLogpushOwnershipChallengeOptionsModel := logpushJobsApiService.NewValidateLogpushOwnershipChallengeOptions()
				validateLogpushOwnershipChallengeOptionsModel.SetDestinationConf("cos://cos-bucket001?region=us-south&instance-id=231f5467-3072-4cb9-9e39-a906fa3032ea")
				validateLogpushOwnershipChallengeOptionsModel.SetOwnershipChallenge("00000000000000000000")
				validateLogpushOwnershipChallengeOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(validateLogpushOwnershipChallengeOptionsModel).ToNot(BeNil())
				Expect(validateLogpushOwnershipChallengeOptionsModel.DestinationConf).To(Equal(core.StringPtr("cos://cos-bucket001?region=us-south&instance-id=231f5467-3072-4cb9-9e39-a906fa3032ea")))
				Expect(validateLogpushOwnershipChallengeOptionsModel.OwnershipChallenge).To(Equal(core.StringPtr("00000000000000000000")))
				Expect(validateLogpushOwnershipChallengeOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
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
			It(`Invoke NewCreateLogpushJobV2RequestLogpushJobLogdnaReq successfully`, func() {
				logdna := map[string]interface{}{"anyKey": "anyValue"}
				_model, err := logpushJobsApiService.NewCreateLogpushJobV2RequestLogpushJobLogdnaReq(logdna)
				Expect(_model).ToNot(BeNil())
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
