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

package custompagesv1_test

import (
	"bytes"
	"context"
	"fmt"
	"github.com/IBM/go-sdk-core/v4/core"
	"github.com/IBM/networking-go-sdk/custompagesv1"
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

var _ = Describe(`CustomPagesV1`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		It(`Instantiate service client`, func() {
			customPagesService, serviceErr := custompagesv1.NewCustomPagesV1(&custompagesv1.CustomPagesV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
				Crn: core.StringPtr(crn),
				ZoneIdentifier: core.StringPtr(zoneIdentifier),
			})
			Expect(customPagesService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			customPagesService, serviceErr := custompagesv1.NewCustomPagesV1(&custompagesv1.CustomPagesV1Options{
				URL: "{BAD_URL_STRING",
				Crn: core.StringPtr(crn),
				ZoneIdentifier: core.StringPtr(zoneIdentifier),
			})
			Expect(customPagesService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			customPagesService, serviceErr := custompagesv1.NewCustomPagesV1(&custompagesv1.CustomPagesV1Options{
				URL: "https://custompagesv1/api",
				Crn: core.StringPtr(crn),
				ZoneIdentifier: core.StringPtr(zoneIdentifier),
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(customPagesService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Validation Error`, func() {
			customPagesService, serviceErr := custompagesv1.NewCustomPagesV1(&custompagesv1.CustomPagesV1Options{})
			Expect(customPagesService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"CUSTOM_PAGES_URL": "https://custompagesv1/api",
				"CUSTOM_PAGES_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				customPagesService, serviceErr := custompagesv1.NewCustomPagesV1UsingExternalConfig(&custompagesv1.CustomPagesV1Options{
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(customPagesService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := customPagesService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != customPagesService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(customPagesService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(customPagesService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				customPagesService, serviceErr := custompagesv1.NewCustomPagesV1UsingExternalConfig(&custompagesv1.CustomPagesV1Options{
					URL: "https://testService/api",
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(customPagesService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(customPagesService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := customPagesService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != customPagesService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(customPagesService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(customPagesService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				customPagesService, serviceErr := custompagesv1.NewCustomPagesV1UsingExternalConfig(&custompagesv1.CustomPagesV1Options{
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				err := customPagesService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(customPagesService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(customPagesService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := customPagesService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != customPagesService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(customPagesService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(customPagesService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"CUSTOM_PAGES_URL": "https://custompagesv1/api",
				"CUSTOM_PAGES_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			customPagesService, serviceErr := custompagesv1.NewCustomPagesV1UsingExternalConfig(&custompagesv1.CustomPagesV1Options{
				Crn: core.StringPtr(crn),
				ZoneIdentifier: core.StringPtr(zoneIdentifier),
			})

			It(`Instantiate service client with error`, func() {
				Expect(customPagesService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"CUSTOM_PAGES_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			customPagesService, serviceErr := custompagesv1.NewCustomPagesV1UsingExternalConfig(&custompagesv1.CustomPagesV1Options{
				URL: "{BAD_URL_STRING",
				Crn: core.StringPtr(crn),
				ZoneIdentifier: core.StringPtr(zoneIdentifier),
			})

			It(`Instantiate service client with error`, func() {
				Expect(customPagesService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = custompagesv1.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`ListInstanceCustomPages(listInstanceCustomPagesOptions *ListInstanceCustomPagesOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		listInstanceCustomPagesPath := "/v1/testString/custom_pages"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listInstanceCustomPagesPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListInstanceCustomPages with error: Operation response processing error`, func() {
				customPagesService, serviceErr := custompagesv1.NewCustomPagesV1(&custompagesv1.CustomPagesV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(customPagesService).ToNot(BeNil())

				// Construct an instance of the ListInstanceCustomPagesOptions model
				listInstanceCustomPagesOptionsModel := new(custompagesv1.ListInstanceCustomPagesOptions)
				listInstanceCustomPagesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := customPagesService.ListInstanceCustomPages(listInstanceCustomPagesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				customPagesService.EnableRetries(0, 0)
				result, response, operationErr = customPagesService.ListInstanceCustomPages(listInstanceCustomPagesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ListInstanceCustomPages(listInstanceCustomPagesOptions *ListInstanceCustomPagesOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		listInstanceCustomPagesPath := "/v1/testString/custom_pages"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listInstanceCustomPagesPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": [{"id": "basic_challenge", "description": "Basic Challenge", "required_tokens": ["::CAPTCHA_BOX::"], "preview_target": "block:basic-sec-captcha", "created_on": "2019-01-01T12:00:00", "modified_on": "2019-01-01T12:00:00", "url": "https://www.example.com/basic_challenge_error.html", "state": "customized"}], "result_info": {"page": 1, "per_page": 20, "total_pages": 1, "count": 10, "total_count": 10}}`)
				}))
			})
			It(`Invoke ListInstanceCustomPages successfully`, func() {
				customPagesService, serviceErr := custompagesv1.NewCustomPagesV1(&custompagesv1.CustomPagesV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(customPagesService).ToNot(BeNil())
				customPagesService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := customPagesService.ListInstanceCustomPages(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListInstanceCustomPagesOptions model
				listInstanceCustomPagesOptionsModel := new(custompagesv1.ListInstanceCustomPagesOptions)
				listInstanceCustomPagesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = customPagesService.ListInstanceCustomPages(listInstanceCustomPagesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = customPagesService.ListInstanceCustomPagesWithContext(ctx, listInstanceCustomPagesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				customPagesService.DisableRetries()
				result, response, operationErr = customPagesService.ListInstanceCustomPages(listInstanceCustomPagesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = customPagesService.ListInstanceCustomPagesWithContext(ctx, listInstanceCustomPagesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke ListInstanceCustomPages with error: Operation request error`, func() {
				customPagesService, serviceErr := custompagesv1.NewCustomPagesV1(&custompagesv1.CustomPagesV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(customPagesService).ToNot(BeNil())

				// Construct an instance of the ListInstanceCustomPagesOptions model
				listInstanceCustomPagesOptionsModel := new(custompagesv1.ListInstanceCustomPagesOptions)
				listInstanceCustomPagesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := customPagesService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := customPagesService.ListInstanceCustomPages(listInstanceCustomPagesOptionsModel)
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
	Describe(`GetInstanceCustomPage(getInstanceCustomPageOptions *GetInstanceCustomPageOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getInstanceCustomPagePath := "/v1/testString/custom_pages/basic_challenge"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getInstanceCustomPagePath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetInstanceCustomPage with error: Operation response processing error`, func() {
				customPagesService, serviceErr := custompagesv1.NewCustomPagesV1(&custompagesv1.CustomPagesV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(customPagesService).ToNot(BeNil())

				// Construct an instance of the GetInstanceCustomPageOptions model
				getInstanceCustomPageOptionsModel := new(custompagesv1.GetInstanceCustomPageOptions)
				getInstanceCustomPageOptionsModel.PageIdentifier = core.StringPtr("basic_challenge")
				getInstanceCustomPageOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := customPagesService.GetInstanceCustomPage(getInstanceCustomPageOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				customPagesService.EnableRetries(0, 0)
				result, response, operationErr = customPagesService.GetInstanceCustomPage(getInstanceCustomPageOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetInstanceCustomPage(getInstanceCustomPageOptions *GetInstanceCustomPageOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getInstanceCustomPagePath := "/v1/testString/custom_pages/basic_challenge"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getInstanceCustomPagePath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": "basic_challenge", "description": "Basic Challenge", "required_tokens": ["::CAPTCHA_BOX::"], "preview_target": "block:basic-sec-captcha", "created_on": "2019-01-01T12:00:00", "modified_on": "2019-01-01T12:00:00", "url": "https://www.example.com/basic_challenge_error.html", "state": "customized"}}`)
				}))
			})
			It(`Invoke GetInstanceCustomPage successfully`, func() {
				customPagesService, serviceErr := custompagesv1.NewCustomPagesV1(&custompagesv1.CustomPagesV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(customPagesService).ToNot(BeNil())
				customPagesService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := customPagesService.GetInstanceCustomPage(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetInstanceCustomPageOptions model
				getInstanceCustomPageOptionsModel := new(custompagesv1.GetInstanceCustomPageOptions)
				getInstanceCustomPageOptionsModel.PageIdentifier = core.StringPtr("basic_challenge")
				getInstanceCustomPageOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = customPagesService.GetInstanceCustomPage(getInstanceCustomPageOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = customPagesService.GetInstanceCustomPageWithContext(ctx, getInstanceCustomPageOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				customPagesService.DisableRetries()
				result, response, operationErr = customPagesService.GetInstanceCustomPage(getInstanceCustomPageOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = customPagesService.GetInstanceCustomPageWithContext(ctx, getInstanceCustomPageOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke GetInstanceCustomPage with error: Operation validation and request error`, func() {
				customPagesService, serviceErr := custompagesv1.NewCustomPagesV1(&custompagesv1.CustomPagesV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(customPagesService).ToNot(BeNil())

				// Construct an instance of the GetInstanceCustomPageOptions model
				getInstanceCustomPageOptionsModel := new(custompagesv1.GetInstanceCustomPageOptions)
				getInstanceCustomPageOptionsModel.PageIdentifier = core.StringPtr("basic_challenge")
				getInstanceCustomPageOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := customPagesService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := customPagesService.GetInstanceCustomPage(getInstanceCustomPageOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetInstanceCustomPageOptions model with no property values
				getInstanceCustomPageOptionsModelNew := new(custompagesv1.GetInstanceCustomPageOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = customPagesService.GetInstanceCustomPage(getInstanceCustomPageOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateInstanceCustomPage(updateInstanceCustomPageOptions *UpdateInstanceCustomPageOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		updateInstanceCustomPagePath := "/v1/testString/custom_pages/basic_challenge"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateInstanceCustomPagePath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateInstanceCustomPage with error: Operation response processing error`, func() {
				customPagesService, serviceErr := custompagesv1.NewCustomPagesV1(&custompagesv1.CustomPagesV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(customPagesService).ToNot(BeNil())

				// Construct an instance of the UpdateInstanceCustomPageOptions model
				updateInstanceCustomPageOptionsModel := new(custompagesv1.UpdateInstanceCustomPageOptions)
				updateInstanceCustomPageOptionsModel.PageIdentifier = core.StringPtr("basic_challenge")
				updateInstanceCustomPageOptionsModel.URL = core.StringPtr("https://www.example.com/basic_challenge_error.html")
				updateInstanceCustomPageOptionsModel.State = core.StringPtr("customized")
				updateInstanceCustomPageOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := customPagesService.UpdateInstanceCustomPage(updateInstanceCustomPageOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				customPagesService.EnableRetries(0, 0)
				result, response, operationErr = customPagesService.UpdateInstanceCustomPage(updateInstanceCustomPageOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UpdateInstanceCustomPage(updateInstanceCustomPageOptions *UpdateInstanceCustomPageOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		updateInstanceCustomPagePath := "/v1/testString/custom_pages/basic_challenge"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateInstanceCustomPagePath))
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
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": "basic_challenge", "description": "Basic Challenge", "required_tokens": ["::CAPTCHA_BOX::"], "preview_target": "block:basic-sec-captcha", "created_on": "2019-01-01T12:00:00", "modified_on": "2019-01-01T12:00:00", "url": "https://www.example.com/basic_challenge_error.html", "state": "customized"}}`)
				}))
			})
			It(`Invoke UpdateInstanceCustomPage successfully`, func() {
				customPagesService, serviceErr := custompagesv1.NewCustomPagesV1(&custompagesv1.CustomPagesV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(customPagesService).ToNot(BeNil())
				customPagesService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := customPagesService.UpdateInstanceCustomPage(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateInstanceCustomPageOptions model
				updateInstanceCustomPageOptionsModel := new(custompagesv1.UpdateInstanceCustomPageOptions)
				updateInstanceCustomPageOptionsModel.PageIdentifier = core.StringPtr("basic_challenge")
				updateInstanceCustomPageOptionsModel.URL = core.StringPtr("https://www.example.com/basic_challenge_error.html")
				updateInstanceCustomPageOptionsModel.State = core.StringPtr("customized")
				updateInstanceCustomPageOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = customPagesService.UpdateInstanceCustomPage(updateInstanceCustomPageOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = customPagesService.UpdateInstanceCustomPageWithContext(ctx, updateInstanceCustomPageOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				customPagesService.DisableRetries()
				result, response, operationErr = customPagesService.UpdateInstanceCustomPage(updateInstanceCustomPageOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = customPagesService.UpdateInstanceCustomPageWithContext(ctx, updateInstanceCustomPageOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke UpdateInstanceCustomPage with error: Operation validation and request error`, func() {
				customPagesService, serviceErr := custompagesv1.NewCustomPagesV1(&custompagesv1.CustomPagesV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(customPagesService).ToNot(BeNil())

				// Construct an instance of the UpdateInstanceCustomPageOptions model
				updateInstanceCustomPageOptionsModel := new(custompagesv1.UpdateInstanceCustomPageOptions)
				updateInstanceCustomPageOptionsModel.PageIdentifier = core.StringPtr("basic_challenge")
				updateInstanceCustomPageOptionsModel.URL = core.StringPtr("https://www.example.com/basic_challenge_error.html")
				updateInstanceCustomPageOptionsModel.State = core.StringPtr("customized")
				updateInstanceCustomPageOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := customPagesService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := customPagesService.UpdateInstanceCustomPage(updateInstanceCustomPageOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateInstanceCustomPageOptions model with no property values
				updateInstanceCustomPageOptionsModelNew := new(custompagesv1.UpdateInstanceCustomPageOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = customPagesService.UpdateInstanceCustomPage(updateInstanceCustomPageOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListZoneCustomPages(listZoneCustomPagesOptions *ListZoneCustomPagesOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		listZoneCustomPagesPath := "/v1/testString/zones/testString/custom_pages"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listZoneCustomPagesPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListZoneCustomPages with error: Operation response processing error`, func() {
				customPagesService, serviceErr := custompagesv1.NewCustomPagesV1(&custompagesv1.CustomPagesV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(customPagesService).ToNot(BeNil())

				// Construct an instance of the ListZoneCustomPagesOptions model
				listZoneCustomPagesOptionsModel := new(custompagesv1.ListZoneCustomPagesOptions)
				listZoneCustomPagesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := customPagesService.ListZoneCustomPages(listZoneCustomPagesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				customPagesService.EnableRetries(0, 0)
				result, response, operationErr = customPagesService.ListZoneCustomPages(listZoneCustomPagesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ListZoneCustomPages(listZoneCustomPagesOptions *ListZoneCustomPagesOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		listZoneCustomPagesPath := "/v1/testString/zones/testString/custom_pages"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listZoneCustomPagesPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": [{"id": "basic_challenge", "description": "Basic Challenge", "required_tokens": ["::CAPTCHA_BOX::"], "preview_target": "block:basic-sec-captcha", "created_on": "2019-01-01T12:00:00", "modified_on": "2019-01-01T12:00:00", "url": "https://www.example.com/basic_challenge_error.html", "state": "customized"}], "result_info": {"page": 1, "per_page": 20, "total_pages": 1, "count": 10, "total_count": 10}}`)
				}))
			})
			It(`Invoke ListZoneCustomPages successfully`, func() {
				customPagesService, serviceErr := custompagesv1.NewCustomPagesV1(&custompagesv1.CustomPagesV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(customPagesService).ToNot(BeNil())
				customPagesService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := customPagesService.ListZoneCustomPages(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListZoneCustomPagesOptions model
				listZoneCustomPagesOptionsModel := new(custompagesv1.ListZoneCustomPagesOptions)
				listZoneCustomPagesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = customPagesService.ListZoneCustomPages(listZoneCustomPagesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = customPagesService.ListZoneCustomPagesWithContext(ctx, listZoneCustomPagesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				customPagesService.DisableRetries()
				result, response, operationErr = customPagesService.ListZoneCustomPages(listZoneCustomPagesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = customPagesService.ListZoneCustomPagesWithContext(ctx, listZoneCustomPagesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke ListZoneCustomPages with error: Operation request error`, func() {
				customPagesService, serviceErr := custompagesv1.NewCustomPagesV1(&custompagesv1.CustomPagesV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(customPagesService).ToNot(BeNil())

				// Construct an instance of the ListZoneCustomPagesOptions model
				listZoneCustomPagesOptionsModel := new(custompagesv1.ListZoneCustomPagesOptions)
				listZoneCustomPagesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := customPagesService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := customPagesService.ListZoneCustomPages(listZoneCustomPagesOptionsModel)
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
	Describe(`GetZoneCustomPage(getZoneCustomPageOptions *GetZoneCustomPageOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getZoneCustomPagePath := "/v1/testString/zones/testString/custom_pages/basic_challenge"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getZoneCustomPagePath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetZoneCustomPage with error: Operation response processing error`, func() {
				customPagesService, serviceErr := custompagesv1.NewCustomPagesV1(&custompagesv1.CustomPagesV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(customPagesService).ToNot(BeNil())

				// Construct an instance of the GetZoneCustomPageOptions model
				getZoneCustomPageOptionsModel := new(custompagesv1.GetZoneCustomPageOptions)
				getZoneCustomPageOptionsModel.PageIdentifier = core.StringPtr("basic_challenge")
				getZoneCustomPageOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := customPagesService.GetZoneCustomPage(getZoneCustomPageOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				customPagesService.EnableRetries(0, 0)
				result, response, operationErr = customPagesService.GetZoneCustomPage(getZoneCustomPageOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetZoneCustomPage(getZoneCustomPageOptions *GetZoneCustomPageOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getZoneCustomPagePath := "/v1/testString/zones/testString/custom_pages/basic_challenge"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getZoneCustomPagePath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": "basic_challenge", "description": "Basic Challenge", "required_tokens": ["::CAPTCHA_BOX::"], "preview_target": "block:basic-sec-captcha", "created_on": "2019-01-01T12:00:00", "modified_on": "2019-01-01T12:00:00", "url": "https://www.example.com/basic_challenge_error.html", "state": "customized"}}`)
				}))
			})
			It(`Invoke GetZoneCustomPage successfully`, func() {
				customPagesService, serviceErr := custompagesv1.NewCustomPagesV1(&custompagesv1.CustomPagesV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(customPagesService).ToNot(BeNil())
				customPagesService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := customPagesService.GetZoneCustomPage(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetZoneCustomPageOptions model
				getZoneCustomPageOptionsModel := new(custompagesv1.GetZoneCustomPageOptions)
				getZoneCustomPageOptionsModel.PageIdentifier = core.StringPtr("basic_challenge")
				getZoneCustomPageOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = customPagesService.GetZoneCustomPage(getZoneCustomPageOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = customPagesService.GetZoneCustomPageWithContext(ctx, getZoneCustomPageOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				customPagesService.DisableRetries()
				result, response, operationErr = customPagesService.GetZoneCustomPage(getZoneCustomPageOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = customPagesService.GetZoneCustomPageWithContext(ctx, getZoneCustomPageOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke GetZoneCustomPage with error: Operation validation and request error`, func() {
				customPagesService, serviceErr := custompagesv1.NewCustomPagesV1(&custompagesv1.CustomPagesV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(customPagesService).ToNot(BeNil())

				// Construct an instance of the GetZoneCustomPageOptions model
				getZoneCustomPageOptionsModel := new(custompagesv1.GetZoneCustomPageOptions)
				getZoneCustomPageOptionsModel.PageIdentifier = core.StringPtr("basic_challenge")
				getZoneCustomPageOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := customPagesService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := customPagesService.GetZoneCustomPage(getZoneCustomPageOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetZoneCustomPageOptions model with no property values
				getZoneCustomPageOptionsModelNew := new(custompagesv1.GetZoneCustomPageOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = customPagesService.GetZoneCustomPage(getZoneCustomPageOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateZoneCustomPage(updateZoneCustomPageOptions *UpdateZoneCustomPageOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		updateZoneCustomPagePath := "/v1/testString/zones/testString/custom_pages/basic_challenge"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateZoneCustomPagePath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateZoneCustomPage with error: Operation response processing error`, func() {
				customPagesService, serviceErr := custompagesv1.NewCustomPagesV1(&custompagesv1.CustomPagesV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(customPagesService).ToNot(BeNil())

				// Construct an instance of the UpdateZoneCustomPageOptions model
				updateZoneCustomPageOptionsModel := new(custompagesv1.UpdateZoneCustomPageOptions)
				updateZoneCustomPageOptionsModel.PageIdentifier = core.StringPtr("basic_challenge")
				updateZoneCustomPageOptionsModel.URL = core.StringPtr("https://www.example.com/basic_challenge_error.html")
				updateZoneCustomPageOptionsModel.State = core.StringPtr("customized")
				updateZoneCustomPageOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := customPagesService.UpdateZoneCustomPage(updateZoneCustomPageOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				customPagesService.EnableRetries(0, 0)
				result, response, operationErr = customPagesService.UpdateZoneCustomPage(updateZoneCustomPageOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UpdateZoneCustomPage(updateZoneCustomPageOptions *UpdateZoneCustomPageOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		updateZoneCustomPagePath := "/v1/testString/zones/testString/custom_pages/basic_challenge"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateZoneCustomPagePath))
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
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": "basic_challenge", "description": "Basic Challenge", "required_tokens": ["::CAPTCHA_BOX::"], "preview_target": "block:basic-sec-captcha", "created_on": "2019-01-01T12:00:00", "modified_on": "2019-01-01T12:00:00", "url": "https://www.example.com/basic_challenge_error.html", "state": "customized"}}`)
				}))
			})
			It(`Invoke UpdateZoneCustomPage successfully`, func() {
				customPagesService, serviceErr := custompagesv1.NewCustomPagesV1(&custompagesv1.CustomPagesV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(customPagesService).ToNot(BeNil())
				customPagesService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := customPagesService.UpdateZoneCustomPage(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateZoneCustomPageOptions model
				updateZoneCustomPageOptionsModel := new(custompagesv1.UpdateZoneCustomPageOptions)
				updateZoneCustomPageOptionsModel.PageIdentifier = core.StringPtr("basic_challenge")
				updateZoneCustomPageOptionsModel.URL = core.StringPtr("https://www.example.com/basic_challenge_error.html")
				updateZoneCustomPageOptionsModel.State = core.StringPtr("customized")
				updateZoneCustomPageOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = customPagesService.UpdateZoneCustomPage(updateZoneCustomPageOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = customPagesService.UpdateZoneCustomPageWithContext(ctx, updateZoneCustomPageOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				customPagesService.DisableRetries()
				result, response, operationErr = customPagesService.UpdateZoneCustomPage(updateZoneCustomPageOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = customPagesService.UpdateZoneCustomPageWithContext(ctx, updateZoneCustomPageOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke UpdateZoneCustomPage with error: Operation validation and request error`, func() {
				customPagesService, serviceErr := custompagesv1.NewCustomPagesV1(&custompagesv1.CustomPagesV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(customPagesService).ToNot(BeNil())

				// Construct an instance of the UpdateZoneCustomPageOptions model
				updateZoneCustomPageOptionsModel := new(custompagesv1.UpdateZoneCustomPageOptions)
				updateZoneCustomPageOptionsModel.PageIdentifier = core.StringPtr("basic_challenge")
				updateZoneCustomPageOptionsModel.URL = core.StringPtr("https://www.example.com/basic_challenge_error.html")
				updateZoneCustomPageOptionsModel.State = core.StringPtr("customized")
				updateZoneCustomPageOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := customPagesService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := customPagesService.UpdateZoneCustomPage(updateZoneCustomPageOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateZoneCustomPageOptions model with no property values
				updateZoneCustomPageOptionsModelNew := new(custompagesv1.UpdateZoneCustomPageOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = customPagesService.UpdateZoneCustomPage(updateZoneCustomPageOptionsModelNew)
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
			zoneIdentifier := "testString"
			customPagesService, _ := custompagesv1.NewCustomPagesV1(&custompagesv1.CustomPagesV1Options{
				URL:           "http://custompagesv1modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
				Crn: core.StringPtr(crn),
				ZoneIdentifier: core.StringPtr(zoneIdentifier),
			})
			It(`Invoke NewGetInstanceCustomPageOptions successfully`, func() {
				// Construct an instance of the GetInstanceCustomPageOptions model
				pageIdentifier := "basic_challenge"
				getInstanceCustomPageOptionsModel := customPagesService.NewGetInstanceCustomPageOptions(pageIdentifier)
				getInstanceCustomPageOptionsModel.SetPageIdentifier("basic_challenge")
				getInstanceCustomPageOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getInstanceCustomPageOptionsModel).ToNot(BeNil())
				Expect(getInstanceCustomPageOptionsModel.PageIdentifier).To(Equal(core.StringPtr("basic_challenge")))
				Expect(getInstanceCustomPageOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetZoneCustomPageOptions successfully`, func() {
				// Construct an instance of the GetZoneCustomPageOptions model
				pageIdentifier := "basic_challenge"
				getZoneCustomPageOptionsModel := customPagesService.NewGetZoneCustomPageOptions(pageIdentifier)
				getZoneCustomPageOptionsModel.SetPageIdentifier("basic_challenge")
				getZoneCustomPageOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getZoneCustomPageOptionsModel).ToNot(BeNil())
				Expect(getZoneCustomPageOptionsModel.PageIdentifier).To(Equal(core.StringPtr("basic_challenge")))
				Expect(getZoneCustomPageOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListInstanceCustomPagesOptions successfully`, func() {
				// Construct an instance of the ListInstanceCustomPagesOptions model
				listInstanceCustomPagesOptionsModel := customPagesService.NewListInstanceCustomPagesOptions()
				listInstanceCustomPagesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listInstanceCustomPagesOptionsModel).ToNot(BeNil())
				Expect(listInstanceCustomPagesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListZoneCustomPagesOptions successfully`, func() {
				// Construct an instance of the ListZoneCustomPagesOptions model
				listZoneCustomPagesOptionsModel := customPagesService.NewListZoneCustomPagesOptions()
				listZoneCustomPagesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listZoneCustomPagesOptionsModel).ToNot(BeNil())
				Expect(listZoneCustomPagesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateInstanceCustomPageOptions successfully`, func() {
				// Construct an instance of the UpdateInstanceCustomPageOptions model
				pageIdentifier := "basic_challenge"
				updateInstanceCustomPageOptionsModel := customPagesService.NewUpdateInstanceCustomPageOptions(pageIdentifier)
				updateInstanceCustomPageOptionsModel.SetPageIdentifier("basic_challenge")
				updateInstanceCustomPageOptionsModel.SetURL("https://www.example.com/basic_challenge_error.html")
				updateInstanceCustomPageOptionsModel.SetState("customized")
				updateInstanceCustomPageOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateInstanceCustomPageOptionsModel).ToNot(BeNil())
				Expect(updateInstanceCustomPageOptionsModel.PageIdentifier).To(Equal(core.StringPtr("basic_challenge")))
				Expect(updateInstanceCustomPageOptionsModel.URL).To(Equal(core.StringPtr("https://www.example.com/basic_challenge_error.html")))
				Expect(updateInstanceCustomPageOptionsModel.State).To(Equal(core.StringPtr("customized")))
				Expect(updateInstanceCustomPageOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateZoneCustomPageOptions successfully`, func() {
				// Construct an instance of the UpdateZoneCustomPageOptions model
				pageIdentifier := "basic_challenge"
				updateZoneCustomPageOptionsModel := customPagesService.NewUpdateZoneCustomPageOptions(pageIdentifier)
				updateZoneCustomPageOptionsModel.SetPageIdentifier("basic_challenge")
				updateZoneCustomPageOptionsModel.SetURL("https://www.example.com/basic_challenge_error.html")
				updateZoneCustomPageOptionsModel.SetState("customized")
				updateZoneCustomPageOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateZoneCustomPageOptionsModel).ToNot(BeNil())
				Expect(updateZoneCustomPageOptionsModel.PageIdentifier).To(Equal(core.StringPtr("basic_challenge")))
				Expect(updateZoneCustomPageOptionsModel.URL).To(Equal(core.StringPtr("https://www.example.com/basic_challenge_error.html")))
				Expect(updateZoneCustomPageOptionsModel.State).To(Equal(core.StringPtr("customized")))
				Expect(updateZoneCustomPageOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
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
