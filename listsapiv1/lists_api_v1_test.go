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

package listsapiv1_test

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
	"github.com/IBM/networking-go-sdk/listsapiv1"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe(`ListsApiV1`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		crn := "testString"
		itemID := "testString"
		listID := "testString"
		operationID := "testString"
		It(`Instantiate service client`, func() {
			listsApiService, serviceErr := listsapiv1.NewListsApiV1(&listsapiv1.ListsApiV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
				Crn:           core.StringPtr(crn),
				ItemID:        core.StringPtr(itemID),
				ListID:        core.StringPtr(listID),
				OperationID:   core.StringPtr(operationID),
			})
			Expect(listsApiService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			listsApiService, serviceErr := listsapiv1.NewListsApiV1(&listsapiv1.ListsApiV1Options{
				URL:         "{BAD_URL_STRING",
				Crn:         core.StringPtr(crn),
				ItemID:      core.StringPtr(itemID),
				ListID:      core.StringPtr(listID),
				OperationID: core.StringPtr(operationID),
			})
			Expect(listsApiService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			listsApiService, serviceErr := listsapiv1.NewListsApiV1(&listsapiv1.ListsApiV1Options{
				URL:         "https://listsapiv1/api",
				Crn:         core.StringPtr(crn),
				ItemID:      core.StringPtr(itemID),
				ListID:      core.StringPtr(listID),
				OperationID: core.StringPtr(operationID),
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(listsApiService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Validation Error`, func() {
			listsApiService, serviceErr := listsapiv1.NewListsApiV1(&listsapiv1.ListsApiV1Options{})
			Expect(listsApiService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		crn := "testString"
		itemID := "testString"
		listID := "testString"
		operationID := "testString"
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"LISTS_API_URL":       "https://listsapiv1/api",
				"LISTS_API_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				listsApiService, serviceErr := listsapiv1.NewListsApiV1UsingExternalConfig(&listsapiv1.ListsApiV1Options{
					Crn:         core.StringPtr(crn),
					ItemID:      core.StringPtr(itemID),
					ListID:      core.StringPtr(listID),
					OperationID: core.StringPtr(operationID),
				})
				Expect(listsApiService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := listsApiService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != listsApiService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(listsApiService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(listsApiService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				listsApiService, serviceErr := listsapiv1.NewListsApiV1UsingExternalConfig(&listsapiv1.ListsApiV1Options{
					URL:         "https://testService/api",
					Crn:         core.StringPtr(crn),
					ItemID:      core.StringPtr(itemID),
					ListID:      core.StringPtr(listID),
					OperationID: core.StringPtr(operationID),
				})
				Expect(listsApiService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(listsApiService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := listsApiService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != listsApiService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(listsApiService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(listsApiService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				listsApiService, serviceErr := listsapiv1.NewListsApiV1UsingExternalConfig(&listsapiv1.ListsApiV1Options{
					Crn:         core.StringPtr(crn),
					ItemID:      core.StringPtr(itemID),
					ListID:      core.StringPtr(listID),
					OperationID: core.StringPtr(operationID),
				})
				err := listsApiService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(listsApiService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(listsApiService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := listsApiService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != listsApiService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(listsApiService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(listsApiService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"LISTS_API_URL":       "https://listsapiv1/api",
				"LISTS_API_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			listsApiService, serviceErr := listsapiv1.NewListsApiV1UsingExternalConfig(&listsapiv1.ListsApiV1Options{
				Crn:         core.StringPtr(crn),
				ItemID:      core.StringPtr(itemID),
				ListID:      core.StringPtr(listID),
				OperationID: core.StringPtr(operationID),
			})

			It(`Instantiate service client with error`, func() {
				Expect(listsApiService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"LISTS_API_AUTH_TYPE": "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			listsApiService, serviceErr := listsapiv1.NewListsApiV1UsingExternalConfig(&listsapiv1.ListsApiV1Options{
				URL:         "{BAD_URL_STRING",
				Crn:         core.StringPtr(crn),
				ItemID:      core.StringPtr(itemID),
				ListID:      core.StringPtr(listID),
				OperationID: core.StringPtr(operationID),
			})

			It(`Instantiate service client with error`, func() {
				Expect(listsApiService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = listsapiv1.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`GetManagedLists(getManagedListsOptions *GetManagedListsOptions) - Operation response error`, func() {
		crn := "testString"
		itemID := "testString"
		listID := "testString"
		operationID := "testString"
		getManagedListsPath := "/v1/testString/rules/managed_lists"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getManagedListsPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetManagedLists with error: Operation response processing error`, func() {
				listsApiService, serviceErr := listsapiv1.NewListsApiV1(&listsapiv1.ListsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ItemID:        core.StringPtr(itemID),
					ListID:        core.StringPtr(listID),
					OperationID:   core.StringPtr(operationID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(listsApiService).ToNot(BeNil())

				// Construct an instance of the GetManagedListsOptions model
				getManagedListsOptionsModel := new(listsapiv1.GetManagedListsOptions)
				getManagedListsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := listsApiService.GetManagedLists(getManagedListsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				listsApiService.EnableRetries(0, 0)
				result, response, operationErr = listsApiService.GetManagedLists(getManagedListsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetManagedLists(getManagedListsOptions *GetManagedListsOptions)`, func() {
		crn := "testString"
		itemID := "testString"
		listID := "testString"
		operationID := "testString"
		getManagedListsPath := "/v1/testString/rules/managed_lists"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getManagedListsPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": [{"name": "cf.malware", "description": "Description", "kind": "ip"}]}`)
				}))
			})
			It(`Invoke GetManagedLists successfully with retries`, func() {
				listsApiService, serviceErr := listsapiv1.NewListsApiV1(&listsapiv1.ListsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ItemID:        core.StringPtr(itemID),
					ListID:        core.StringPtr(listID),
					OperationID:   core.StringPtr(operationID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(listsApiService).ToNot(BeNil())
				listsApiService.EnableRetries(0, 0)

				// Construct an instance of the GetManagedListsOptions model
				getManagedListsOptionsModel := new(listsapiv1.GetManagedListsOptions)
				getManagedListsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := listsApiService.GetManagedListsWithContext(ctx, getManagedListsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				listsApiService.DisableRetries()
				result, response, operationErr := listsApiService.GetManagedLists(getManagedListsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = listsApiService.GetManagedListsWithContext(ctx, getManagedListsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getManagedListsPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": [{"name": "cf.malware", "description": "Description", "kind": "ip"}]}`)
				}))
			})
			It(`Invoke GetManagedLists successfully`, func() {
				listsApiService, serviceErr := listsapiv1.NewListsApiV1(&listsapiv1.ListsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ItemID:        core.StringPtr(itemID),
					ListID:        core.StringPtr(listID),
					OperationID:   core.StringPtr(operationID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(listsApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := listsApiService.GetManagedLists(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetManagedListsOptions model
				getManagedListsOptionsModel := new(listsapiv1.GetManagedListsOptions)
				getManagedListsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = listsApiService.GetManagedLists(getManagedListsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetManagedLists with error: Operation request error`, func() {
				listsApiService, serviceErr := listsapiv1.NewListsApiV1(&listsapiv1.ListsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ItemID:        core.StringPtr(itemID),
					ListID:        core.StringPtr(listID),
					OperationID:   core.StringPtr(operationID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(listsApiService).ToNot(BeNil())

				// Construct an instance of the GetManagedListsOptions model
				getManagedListsOptionsModel := new(listsapiv1.GetManagedListsOptions)
				getManagedListsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := listsApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := listsApiService.GetManagedLists(getManagedListsOptionsModel)
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
			It(`Invoke GetManagedLists successfully`, func() {
				listsApiService, serviceErr := listsapiv1.NewListsApiV1(&listsapiv1.ListsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ItemID:        core.StringPtr(itemID),
					ListID:        core.StringPtr(listID),
					OperationID:   core.StringPtr(operationID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(listsApiService).ToNot(BeNil())

				// Construct an instance of the GetManagedListsOptions model
				getManagedListsOptionsModel := new(listsapiv1.GetManagedListsOptions)
				getManagedListsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := listsApiService.GetManagedLists(getManagedListsOptionsModel)
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
	Describe(`GetCustomLists(getCustomListsOptions *GetCustomListsOptions) - Operation response error`, func() {
		crn := "testString"
		itemID := "testString"
		listID := "testString"
		operationID := "testString"
		getCustomListsPath := "/v1/testString/rules/lists"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getCustomListsPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetCustomLists with error: Operation response processing error`, func() {
				listsApiService, serviceErr := listsapiv1.NewListsApiV1(&listsapiv1.ListsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ItemID:        core.StringPtr(itemID),
					ListID:        core.StringPtr(listID),
					OperationID:   core.StringPtr(operationID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(listsApiService).ToNot(BeNil())

				// Construct an instance of the GetCustomListsOptions model
				getCustomListsOptionsModel := new(listsapiv1.GetCustomListsOptions)
				getCustomListsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := listsApiService.GetCustomLists(getCustomListsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				listsApiService.EnableRetries(0, 0)
				result, response, operationErr = listsApiService.GetCustomLists(getCustomListsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetCustomLists(getCustomListsOptions *GetCustomListsOptions)`, func() {
		crn := "testString"
		itemID := "testString"
		listID := "testString"
		operationID := "testString"
		getCustomListsPath := "/v1/testString/rules/lists"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getCustomListsPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": [{"name": "good_ips", "id": "ID", "description": "Description", "kind": "ip", "num_items": 10, "num_referencing_filters": 5}]}`)
				}))
			})
			It(`Invoke GetCustomLists successfully with retries`, func() {
				listsApiService, serviceErr := listsapiv1.NewListsApiV1(&listsapiv1.ListsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ItemID:        core.StringPtr(itemID),
					ListID:        core.StringPtr(listID),
					OperationID:   core.StringPtr(operationID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(listsApiService).ToNot(BeNil())
				listsApiService.EnableRetries(0, 0)

				// Construct an instance of the GetCustomListsOptions model
				getCustomListsOptionsModel := new(listsapiv1.GetCustomListsOptions)
				getCustomListsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := listsApiService.GetCustomListsWithContext(ctx, getCustomListsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				listsApiService.DisableRetries()
				result, response, operationErr := listsApiService.GetCustomLists(getCustomListsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = listsApiService.GetCustomListsWithContext(ctx, getCustomListsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getCustomListsPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": [{"name": "good_ips", "id": "ID", "description": "Description", "kind": "ip", "num_items": 10, "num_referencing_filters": 5}]}`)
				}))
			})
			It(`Invoke GetCustomLists successfully`, func() {
				listsApiService, serviceErr := listsapiv1.NewListsApiV1(&listsapiv1.ListsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ItemID:        core.StringPtr(itemID),
					ListID:        core.StringPtr(listID),
					OperationID:   core.StringPtr(operationID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(listsApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := listsApiService.GetCustomLists(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetCustomListsOptions model
				getCustomListsOptionsModel := new(listsapiv1.GetCustomListsOptions)
				getCustomListsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = listsApiService.GetCustomLists(getCustomListsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetCustomLists with error: Operation request error`, func() {
				listsApiService, serviceErr := listsapiv1.NewListsApiV1(&listsapiv1.ListsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ItemID:        core.StringPtr(itemID),
					ListID:        core.StringPtr(listID),
					OperationID:   core.StringPtr(operationID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(listsApiService).ToNot(BeNil())

				// Construct an instance of the GetCustomListsOptions model
				getCustomListsOptionsModel := new(listsapiv1.GetCustomListsOptions)
				getCustomListsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := listsApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := listsApiService.GetCustomLists(getCustomListsOptionsModel)
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
			It(`Invoke GetCustomLists successfully`, func() {
				listsApiService, serviceErr := listsapiv1.NewListsApiV1(&listsapiv1.ListsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ItemID:        core.StringPtr(itemID),
					ListID:        core.StringPtr(listID),
					OperationID:   core.StringPtr(operationID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(listsApiService).ToNot(BeNil())

				// Construct an instance of the GetCustomListsOptions model
				getCustomListsOptionsModel := new(listsapiv1.GetCustomListsOptions)
				getCustomListsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := listsApiService.GetCustomLists(getCustomListsOptionsModel)
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
	Describe(`CreateCustomLists(createCustomListsOptions *CreateCustomListsOptions) - Operation response error`, func() {
		crn := "testString"
		itemID := "testString"
		listID := "testString"
		operationID := "testString"
		createCustomListsPath := "/v1/testString/rules/lists"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createCustomListsPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateCustomLists with error: Operation response processing error`, func() {
				listsApiService, serviceErr := listsapiv1.NewListsApiV1(&listsapiv1.ListsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ItemID:        core.StringPtr(itemID),
					ListID:        core.StringPtr(listID),
					OperationID:   core.StringPtr(operationID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(listsApiService).ToNot(BeNil())

				// Construct an instance of the CreateCustomListsOptions model
				createCustomListsOptionsModel := new(listsapiv1.CreateCustomListsOptions)
				createCustomListsOptionsModel.Kind = core.StringPtr("ip")
				createCustomListsOptionsModel.Name = core.StringPtr("testString")
				createCustomListsOptionsModel.Description = core.StringPtr("testString")
				createCustomListsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := listsApiService.CreateCustomLists(createCustomListsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				listsApiService.EnableRetries(0, 0)
				result, response, operationErr = listsApiService.CreateCustomLists(createCustomListsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateCustomLists(createCustomListsOptions *CreateCustomListsOptions)`, func() {
		crn := "testString"
		itemID := "testString"
		listID := "testString"
		operationID := "testString"
		createCustomListsPath := "/v1/testString/rules/lists"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createCustomListsPath))
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
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"name": "good_ips", "id": "ID", "description": "Description", "kind": "ip", "num_items": 10, "num_referencing_filters": 5}}`)
				}))
			})
			It(`Invoke CreateCustomLists successfully with retries`, func() {
				listsApiService, serviceErr := listsapiv1.NewListsApiV1(&listsapiv1.ListsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ItemID:        core.StringPtr(itemID),
					ListID:        core.StringPtr(listID),
					OperationID:   core.StringPtr(operationID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(listsApiService).ToNot(BeNil())
				listsApiService.EnableRetries(0, 0)

				// Construct an instance of the CreateCustomListsOptions model
				createCustomListsOptionsModel := new(listsapiv1.CreateCustomListsOptions)
				createCustomListsOptionsModel.Kind = core.StringPtr("ip")
				createCustomListsOptionsModel.Name = core.StringPtr("testString")
				createCustomListsOptionsModel.Description = core.StringPtr("testString")
				createCustomListsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := listsApiService.CreateCustomListsWithContext(ctx, createCustomListsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				listsApiService.DisableRetries()
				result, response, operationErr := listsApiService.CreateCustomLists(createCustomListsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = listsApiService.CreateCustomListsWithContext(ctx, createCustomListsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createCustomListsPath))
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
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"name": "good_ips", "id": "ID", "description": "Description", "kind": "ip", "num_items": 10, "num_referencing_filters": 5}}`)
				}))
			})
			It(`Invoke CreateCustomLists successfully`, func() {
				listsApiService, serviceErr := listsapiv1.NewListsApiV1(&listsapiv1.ListsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ItemID:        core.StringPtr(itemID),
					ListID:        core.StringPtr(listID),
					OperationID:   core.StringPtr(operationID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(listsApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := listsApiService.CreateCustomLists(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateCustomListsOptions model
				createCustomListsOptionsModel := new(listsapiv1.CreateCustomListsOptions)
				createCustomListsOptionsModel.Kind = core.StringPtr("ip")
				createCustomListsOptionsModel.Name = core.StringPtr("testString")
				createCustomListsOptionsModel.Description = core.StringPtr("testString")
				createCustomListsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = listsApiService.CreateCustomLists(createCustomListsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateCustomLists with error: Operation request error`, func() {
				listsApiService, serviceErr := listsapiv1.NewListsApiV1(&listsapiv1.ListsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ItemID:        core.StringPtr(itemID),
					ListID:        core.StringPtr(listID),
					OperationID:   core.StringPtr(operationID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(listsApiService).ToNot(BeNil())

				// Construct an instance of the CreateCustomListsOptions model
				createCustomListsOptionsModel := new(listsapiv1.CreateCustomListsOptions)
				createCustomListsOptionsModel.Kind = core.StringPtr("ip")
				createCustomListsOptionsModel.Name = core.StringPtr("testString")
				createCustomListsOptionsModel.Description = core.StringPtr("testString")
				createCustomListsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := listsApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := listsApiService.CreateCustomLists(createCustomListsOptionsModel)
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
			It(`Invoke CreateCustomLists successfully`, func() {
				listsApiService, serviceErr := listsapiv1.NewListsApiV1(&listsapiv1.ListsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ItemID:        core.StringPtr(itemID),
					ListID:        core.StringPtr(listID),
					OperationID:   core.StringPtr(operationID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(listsApiService).ToNot(BeNil())

				// Construct an instance of the CreateCustomListsOptions model
				createCustomListsOptionsModel := new(listsapiv1.CreateCustomListsOptions)
				createCustomListsOptionsModel.Kind = core.StringPtr("ip")
				createCustomListsOptionsModel.Name = core.StringPtr("testString")
				createCustomListsOptionsModel.Description = core.StringPtr("testString")
				createCustomListsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := listsApiService.CreateCustomLists(createCustomListsOptionsModel)
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
	Describe(`GetCustomList(getCustomListOptions *GetCustomListOptions) - Operation response error`, func() {
		crn := "testString"
		itemID := "testString"
		listID := "testString"
		operationID := "testString"
		getCustomListPath := "/v1/testString/rules/lists/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getCustomListPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetCustomList with error: Operation response processing error`, func() {
				listsApiService, serviceErr := listsapiv1.NewListsApiV1(&listsapiv1.ListsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ItemID:        core.StringPtr(itemID),
					ListID:        core.StringPtr(listID),
					OperationID:   core.StringPtr(operationID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(listsApiService).ToNot(BeNil())

				// Construct an instance of the GetCustomListOptions model
				getCustomListOptionsModel := new(listsapiv1.GetCustomListOptions)
				getCustomListOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := listsApiService.GetCustomList(getCustomListOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				listsApiService.EnableRetries(0, 0)
				result, response, operationErr = listsApiService.GetCustomList(getCustomListOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetCustomList(getCustomListOptions *GetCustomListOptions)`, func() {
		crn := "testString"
		itemID := "testString"
		listID := "testString"
		operationID := "testString"
		getCustomListPath := "/v1/testString/rules/lists/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getCustomListPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"name": "good_ips", "id": "ID", "description": "Description", "kind": "ip", "num_items": 10, "num_referencing_filters": 5}}`)
				}))
			})
			It(`Invoke GetCustomList successfully with retries`, func() {
				listsApiService, serviceErr := listsapiv1.NewListsApiV1(&listsapiv1.ListsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ItemID:        core.StringPtr(itemID),
					ListID:        core.StringPtr(listID),
					OperationID:   core.StringPtr(operationID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(listsApiService).ToNot(BeNil())
				listsApiService.EnableRetries(0, 0)

				// Construct an instance of the GetCustomListOptions model
				getCustomListOptionsModel := new(listsapiv1.GetCustomListOptions)
				getCustomListOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := listsApiService.GetCustomListWithContext(ctx, getCustomListOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				listsApiService.DisableRetries()
				result, response, operationErr := listsApiService.GetCustomList(getCustomListOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = listsApiService.GetCustomListWithContext(ctx, getCustomListOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getCustomListPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"name": "good_ips", "id": "ID", "description": "Description", "kind": "ip", "num_items": 10, "num_referencing_filters": 5}}`)
				}))
			})
			It(`Invoke GetCustomList successfully`, func() {
				listsApiService, serviceErr := listsapiv1.NewListsApiV1(&listsapiv1.ListsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ItemID:        core.StringPtr(itemID),
					ListID:        core.StringPtr(listID),
					OperationID:   core.StringPtr(operationID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(listsApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := listsApiService.GetCustomList(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetCustomListOptions model
				getCustomListOptionsModel := new(listsapiv1.GetCustomListOptions)
				getCustomListOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = listsApiService.GetCustomList(getCustomListOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetCustomList with error: Operation request error`, func() {
				listsApiService, serviceErr := listsapiv1.NewListsApiV1(&listsapiv1.ListsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ItemID:        core.StringPtr(itemID),
					ListID:        core.StringPtr(listID),
					OperationID:   core.StringPtr(operationID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(listsApiService).ToNot(BeNil())

				// Construct an instance of the GetCustomListOptions model
				getCustomListOptionsModel := new(listsapiv1.GetCustomListOptions)
				getCustomListOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := listsApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := listsApiService.GetCustomList(getCustomListOptionsModel)
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
			It(`Invoke GetCustomList successfully`, func() {
				listsApiService, serviceErr := listsapiv1.NewListsApiV1(&listsapiv1.ListsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ItemID:        core.StringPtr(itemID),
					ListID:        core.StringPtr(listID),
					OperationID:   core.StringPtr(operationID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(listsApiService).ToNot(BeNil())

				// Construct an instance of the GetCustomListOptions model
				getCustomListOptionsModel := new(listsapiv1.GetCustomListOptions)
				getCustomListOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := listsApiService.GetCustomList(getCustomListOptionsModel)
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
	Describe(`UpdateCustomList(updateCustomListOptions *UpdateCustomListOptions) - Operation response error`, func() {
		crn := "testString"
		itemID := "testString"
		listID := "testString"
		operationID := "testString"
		updateCustomListPath := "/v1/testString/rules/lists/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateCustomListPath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateCustomList with error: Operation response processing error`, func() {
				listsApiService, serviceErr := listsapiv1.NewListsApiV1(&listsapiv1.ListsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ItemID:        core.StringPtr(itemID),
					ListID:        core.StringPtr(listID),
					OperationID:   core.StringPtr(operationID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(listsApiService).ToNot(BeNil())

				// Construct an instance of the UpdateCustomListOptions model
				updateCustomListOptionsModel := new(listsapiv1.UpdateCustomListOptions)
				updateCustomListOptionsModel.Description = core.StringPtr("testString")
				updateCustomListOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := listsApiService.UpdateCustomList(updateCustomListOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				listsApiService.EnableRetries(0, 0)
				result, response, operationErr = listsApiService.UpdateCustomList(updateCustomListOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateCustomList(updateCustomListOptions *UpdateCustomListOptions)`, func() {
		crn := "testString"
		itemID := "testString"
		listID := "testString"
		operationID := "testString"
		updateCustomListPath := "/v1/testString/rules/lists/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateCustomListPath))
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
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"name": "good_ips", "id": "ID", "description": "Description", "kind": "ip", "num_items": 10, "num_referencing_filters": 5}}`)
				}))
			})
			It(`Invoke UpdateCustomList successfully with retries`, func() {
				listsApiService, serviceErr := listsapiv1.NewListsApiV1(&listsapiv1.ListsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ItemID:        core.StringPtr(itemID),
					ListID:        core.StringPtr(listID),
					OperationID:   core.StringPtr(operationID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(listsApiService).ToNot(BeNil())
				listsApiService.EnableRetries(0, 0)

				// Construct an instance of the UpdateCustomListOptions model
				updateCustomListOptionsModel := new(listsapiv1.UpdateCustomListOptions)
				updateCustomListOptionsModel.Description = core.StringPtr("testString")
				updateCustomListOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := listsApiService.UpdateCustomListWithContext(ctx, updateCustomListOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				listsApiService.DisableRetries()
				result, response, operationErr := listsApiService.UpdateCustomList(updateCustomListOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = listsApiService.UpdateCustomListWithContext(ctx, updateCustomListOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(updateCustomListPath))
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
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"name": "good_ips", "id": "ID", "description": "Description", "kind": "ip", "num_items": 10, "num_referencing_filters": 5}}`)
				}))
			})
			It(`Invoke UpdateCustomList successfully`, func() {
				listsApiService, serviceErr := listsapiv1.NewListsApiV1(&listsapiv1.ListsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ItemID:        core.StringPtr(itemID),
					ListID:        core.StringPtr(listID),
					OperationID:   core.StringPtr(operationID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(listsApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := listsApiService.UpdateCustomList(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateCustomListOptions model
				updateCustomListOptionsModel := new(listsapiv1.UpdateCustomListOptions)
				updateCustomListOptionsModel.Description = core.StringPtr("testString")
				updateCustomListOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = listsApiService.UpdateCustomList(updateCustomListOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateCustomList with error: Operation request error`, func() {
				listsApiService, serviceErr := listsapiv1.NewListsApiV1(&listsapiv1.ListsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ItemID:        core.StringPtr(itemID),
					ListID:        core.StringPtr(listID),
					OperationID:   core.StringPtr(operationID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(listsApiService).ToNot(BeNil())

				// Construct an instance of the UpdateCustomListOptions model
				updateCustomListOptionsModel := new(listsapiv1.UpdateCustomListOptions)
				updateCustomListOptionsModel.Description = core.StringPtr("testString")
				updateCustomListOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := listsApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := listsApiService.UpdateCustomList(updateCustomListOptionsModel)
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
			It(`Invoke UpdateCustomList successfully`, func() {
				listsApiService, serviceErr := listsapiv1.NewListsApiV1(&listsapiv1.ListsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ItemID:        core.StringPtr(itemID),
					ListID:        core.StringPtr(listID),
					OperationID:   core.StringPtr(operationID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(listsApiService).ToNot(BeNil())

				// Construct an instance of the UpdateCustomListOptions model
				updateCustomListOptionsModel := new(listsapiv1.UpdateCustomListOptions)
				updateCustomListOptionsModel.Description = core.StringPtr("testString")
				updateCustomListOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := listsApiService.UpdateCustomList(updateCustomListOptionsModel)
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
	Describe(`DeleteCustomList(deleteCustomListOptions *DeleteCustomListOptions) - Operation response error`, func() {
		crn := "testString"
		itemID := "testString"
		listID := "testString"
		operationID := "testString"
		deleteCustomListPath := "/v1/testString/rules/lists/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteCustomListPath))
					Expect(req.Method).To(Equal("DELETE"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke DeleteCustomList with error: Operation response processing error`, func() {
				listsApiService, serviceErr := listsapiv1.NewListsApiV1(&listsapiv1.ListsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ItemID:        core.StringPtr(itemID),
					ListID:        core.StringPtr(listID),
					OperationID:   core.StringPtr(operationID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(listsApiService).ToNot(BeNil())

				// Construct an instance of the DeleteCustomListOptions model
				deleteCustomListOptionsModel := new(listsapiv1.DeleteCustomListOptions)
				deleteCustomListOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := listsApiService.DeleteCustomList(deleteCustomListOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				listsApiService.EnableRetries(0, 0)
				result, response, operationErr = listsApiService.DeleteCustomList(deleteCustomListOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteCustomList(deleteCustomListOptions *DeleteCustomListOptions)`, func() {
		crn := "testString"
		itemID := "testString"
		listID := "testString"
		operationID := "testString"
		deleteCustomListPath := "/v1/testString/rules/lists/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteCustomListPath))
					Expect(req.Method).To(Equal("DELETE"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": "34b12448945f11eaa1b71c4d701ab86e"}}`)
				}))
			})
			It(`Invoke DeleteCustomList successfully with retries`, func() {
				listsApiService, serviceErr := listsapiv1.NewListsApiV1(&listsapiv1.ListsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ItemID:        core.StringPtr(itemID),
					ListID:        core.StringPtr(listID),
					OperationID:   core.StringPtr(operationID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(listsApiService).ToNot(BeNil())
				listsApiService.EnableRetries(0, 0)

				// Construct an instance of the DeleteCustomListOptions model
				deleteCustomListOptionsModel := new(listsapiv1.DeleteCustomListOptions)
				deleteCustomListOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := listsApiService.DeleteCustomListWithContext(ctx, deleteCustomListOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				listsApiService.DisableRetries()
				result, response, operationErr := listsApiService.DeleteCustomList(deleteCustomListOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = listsApiService.DeleteCustomListWithContext(ctx, deleteCustomListOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(deleteCustomListPath))
					Expect(req.Method).To(Equal("DELETE"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": "34b12448945f11eaa1b71c4d701ab86e"}}`)
				}))
			})
			It(`Invoke DeleteCustomList successfully`, func() {
				listsApiService, serviceErr := listsapiv1.NewListsApiV1(&listsapiv1.ListsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ItemID:        core.StringPtr(itemID),
					ListID:        core.StringPtr(listID),
					OperationID:   core.StringPtr(operationID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(listsApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := listsApiService.DeleteCustomList(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DeleteCustomListOptions model
				deleteCustomListOptionsModel := new(listsapiv1.DeleteCustomListOptions)
				deleteCustomListOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = listsApiService.DeleteCustomList(deleteCustomListOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke DeleteCustomList with error: Operation request error`, func() {
				listsApiService, serviceErr := listsapiv1.NewListsApiV1(&listsapiv1.ListsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ItemID:        core.StringPtr(itemID),
					ListID:        core.StringPtr(listID),
					OperationID:   core.StringPtr(operationID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(listsApiService).ToNot(BeNil())

				// Construct an instance of the DeleteCustomListOptions model
				deleteCustomListOptionsModel := new(listsapiv1.DeleteCustomListOptions)
				deleteCustomListOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := listsApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := listsApiService.DeleteCustomList(deleteCustomListOptionsModel)
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
			It(`Invoke DeleteCustomList successfully`, func() {
				listsApiService, serviceErr := listsapiv1.NewListsApiV1(&listsapiv1.ListsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ItemID:        core.StringPtr(itemID),
					ListID:        core.StringPtr(listID),
					OperationID:   core.StringPtr(operationID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(listsApiService).ToNot(BeNil())

				// Construct an instance of the DeleteCustomListOptions model
				deleteCustomListOptionsModel := new(listsapiv1.DeleteCustomListOptions)
				deleteCustomListOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := listsApiService.DeleteCustomList(deleteCustomListOptionsModel)
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
	Describe(`GetListItems(getListItemsOptions *GetListItemsOptions) - Operation response error`, func() {
		crn := "testString"
		itemID := "testString"
		listID := "testString"
		operationID := "testString"
		getListItemsPath := "/v1/testString/rules/lists/testString/items"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getListItemsPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetListItems with error: Operation response processing error`, func() {
				listsApiService, serviceErr := listsapiv1.NewListsApiV1(&listsapiv1.ListsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ItemID:        core.StringPtr(itemID),
					ListID:        core.StringPtr(listID),
					OperationID:   core.StringPtr(operationID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(listsApiService).ToNot(BeNil())

				// Construct an instance of the GetListItemsOptions model
				getListItemsOptionsModel := new(listsapiv1.GetListItemsOptions)
				getListItemsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := listsApiService.GetListItems(getListItemsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				listsApiService.EnableRetries(0, 0)
				result, response, operationErr = listsApiService.GetListItems(getListItemsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetListItems(getListItemsOptions *GetListItemsOptions)`, func() {
		crn := "testString"
		itemID := "testString"
		listID := "testString"
		operationID := "testString"
		getListItemsPath := "/v1/testString/rules/lists/testString/items"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getListItemsPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": [{"id": "70c2009751b24ffc9ed1ab462ba957b4", "asn": 19604, "comment": "My list of developer IPs.", "hostname": "cloud.ibm.com", "ip": "172.64.0.0/13", "created_on": "2025-03-21T16:19:21Z", "modified_on": "2025-03-21T16:19:37Z"}]}`)
				}))
			})
			It(`Invoke GetListItems successfully with retries`, func() {
				listsApiService, serviceErr := listsapiv1.NewListsApiV1(&listsapiv1.ListsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ItemID:        core.StringPtr(itemID),
					ListID:        core.StringPtr(listID),
					OperationID:   core.StringPtr(operationID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(listsApiService).ToNot(BeNil())
				listsApiService.EnableRetries(0, 0)

				// Construct an instance of the GetListItemsOptions model
				getListItemsOptionsModel := new(listsapiv1.GetListItemsOptions)
				getListItemsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := listsApiService.GetListItemsWithContext(ctx, getListItemsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				listsApiService.DisableRetries()
				result, response, operationErr := listsApiService.GetListItems(getListItemsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = listsApiService.GetListItemsWithContext(ctx, getListItemsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getListItemsPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": [{"id": "70c2009751b24ffc9ed1ab462ba957b4", "asn": 19604, "comment": "My list of developer IPs.", "hostname": "cloud.ibm.com", "ip": "172.64.0.0/13", "created_on": "2025-03-21T16:19:21Z", "modified_on": "2025-03-21T16:19:37Z"}]}`)
				}))
			})
			It(`Invoke GetListItems successfully`, func() {
				listsApiService, serviceErr := listsapiv1.NewListsApiV1(&listsapiv1.ListsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ItemID:        core.StringPtr(itemID),
					ListID:        core.StringPtr(listID),
					OperationID:   core.StringPtr(operationID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(listsApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := listsApiService.GetListItems(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetListItemsOptions model
				getListItemsOptionsModel := new(listsapiv1.GetListItemsOptions)
				getListItemsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = listsApiService.GetListItems(getListItemsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetListItems with error: Operation request error`, func() {
				listsApiService, serviceErr := listsapiv1.NewListsApiV1(&listsapiv1.ListsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ItemID:        core.StringPtr(itemID),
					ListID:        core.StringPtr(listID),
					OperationID:   core.StringPtr(operationID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(listsApiService).ToNot(BeNil())

				// Construct an instance of the GetListItemsOptions model
				getListItemsOptionsModel := new(listsapiv1.GetListItemsOptions)
				getListItemsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := listsApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := listsApiService.GetListItems(getListItemsOptionsModel)
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
			It(`Invoke GetListItems successfully`, func() {
				listsApiService, serviceErr := listsapiv1.NewListsApiV1(&listsapiv1.ListsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ItemID:        core.StringPtr(itemID),
					ListID:        core.StringPtr(listID),
					OperationID:   core.StringPtr(operationID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(listsApiService).ToNot(BeNil())

				// Construct an instance of the GetListItemsOptions model
				getListItemsOptionsModel := new(listsapiv1.GetListItemsOptions)
				getListItemsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := listsApiService.GetListItems(getListItemsOptionsModel)
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
	Describe(`CreateListItems(createListItemsOptions *CreateListItemsOptions) - Operation response error`, func() {
		crn := "testString"
		itemID := "testString"
		listID := "testString"
		operationID := "testString"
		createListItemsPath := "/v1/testString/rules/lists/testString/items"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createListItemsPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateListItems with error: Operation response processing error`, func() {
				listsApiService, serviceErr := listsapiv1.NewListsApiV1(&listsapiv1.ListsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ItemID:        core.StringPtr(itemID),
					ListID:        core.StringPtr(listID),
					OperationID:   core.StringPtr(operationID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(listsApiService).ToNot(BeNil())

				// Construct an instance of the CreateListItemsReqItem model
				createListItemsReqItemModel := new(listsapiv1.CreateListItemsReqItem)
				createListItemsReqItemModel.Asn = core.Float64Ptr(float64(19604))
				createListItemsReqItemModel.Comment = core.StringPtr("My list of developer IPs.")
				createListItemsReqItemModel.Hostname = core.StringPtr("cloud.ibm.com")
				createListItemsReqItemModel.Ip = core.StringPtr("172.64.0.0/13")

				// Construct an instance of the CreateListItemsOptions model
				createListItemsOptionsModel := new(listsapiv1.CreateListItemsOptions)
				createListItemsOptionsModel.CreateListItemsReqItem = []listsapiv1.CreateListItemsReqItem{*createListItemsReqItemModel}
				createListItemsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := listsApiService.CreateListItems(createListItemsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				listsApiService.EnableRetries(0, 0)
				result, response, operationErr = listsApiService.CreateListItems(createListItemsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateListItems(createListItemsOptions *CreateListItemsOptions)`, func() {
		crn := "testString"
		itemID := "testString"
		listID := "testString"
		operationID := "testString"
		createListItemsPath := "/v1/testString/rules/lists/testString/items"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createListItemsPath))
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
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"operation_id": "53d73a83d33d4e3b8791764a9b9f2412"}}`)
				}))
			})
			It(`Invoke CreateListItems successfully with retries`, func() {
				listsApiService, serviceErr := listsapiv1.NewListsApiV1(&listsapiv1.ListsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ItemID:        core.StringPtr(itemID),
					ListID:        core.StringPtr(listID),
					OperationID:   core.StringPtr(operationID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(listsApiService).ToNot(BeNil())
				listsApiService.EnableRetries(0, 0)

				// Construct an instance of the CreateListItemsReqItem model
				createListItemsReqItemModel := new(listsapiv1.CreateListItemsReqItem)
				createListItemsReqItemModel.Asn = core.Float64Ptr(float64(19604))
				createListItemsReqItemModel.Comment = core.StringPtr("My list of developer IPs.")
				createListItemsReqItemModel.Hostname = core.StringPtr("cloud.ibm.com")
				createListItemsReqItemModel.Ip = core.StringPtr("172.64.0.0/13")

				// Construct an instance of the CreateListItemsOptions model
				createListItemsOptionsModel := new(listsapiv1.CreateListItemsOptions)
				createListItemsOptionsModel.CreateListItemsReqItem = []listsapiv1.CreateListItemsReqItem{*createListItemsReqItemModel}
				createListItemsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := listsApiService.CreateListItemsWithContext(ctx, createListItemsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				listsApiService.DisableRetries()
				result, response, operationErr := listsApiService.CreateListItems(createListItemsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = listsApiService.CreateListItemsWithContext(ctx, createListItemsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createListItemsPath))
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
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"operation_id": "53d73a83d33d4e3b8791764a9b9f2412"}}`)
				}))
			})
			It(`Invoke CreateListItems successfully`, func() {
				listsApiService, serviceErr := listsapiv1.NewListsApiV1(&listsapiv1.ListsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ItemID:        core.StringPtr(itemID),
					ListID:        core.StringPtr(listID),
					OperationID:   core.StringPtr(operationID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(listsApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := listsApiService.CreateListItems(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateListItemsReqItem model
				createListItemsReqItemModel := new(listsapiv1.CreateListItemsReqItem)
				createListItemsReqItemModel.Asn = core.Float64Ptr(float64(19604))
				createListItemsReqItemModel.Comment = core.StringPtr("My list of developer IPs.")
				createListItemsReqItemModel.Hostname = core.StringPtr("cloud.ibm.com")
				createListItemsReqItemModel.Ip = core.StringPtr("172.64.0.0/13")

				// Construct an instance of the CreateListItemsOptions model
				createListItemsOptionsModel := new(listsapiv1.CreateListItemsOptions)
				createListItemsOptionsModel.CreateListItemsReqItem = []listsapiv1.CreateListItemsReqItem{*createListItemsReqItemModel}
				createListItemsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = listsApiService.CreateListItems(createListItemsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateListItems with error: Operation request error`, func() {
				listsApiService, serviceErr := listsapiv1.NewListsApiV1(&listsapiv1.ListsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ItemID:        core.StringPtr(itemID),
					ListID:        core.StringPtr(listID),
					OperationID:   core.StringPtr(operationID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(listsApiService).ToNot(BeNil())

				// Construct an instance of the CreateListItemsReqItem model
				createListItemsReqItemModel := new(listsapiv1.CreateListItemsReqItem)
				createListItemsReqItemModel.Asn = core.Float64Ptr(float64(19604))
				createListItemsReqItemModel.Comment = core.StringPtr("My list of developer IPs.")
				createListItemsReqItemModel.Hostname = core.StringPtr("cloud.ibm.com")
				createListItemsReqItemModel.Ip = core.StringPtr("172.64.0.0/13")

				// Construct an instance of the CreateListItemsOptions model
				createListItemsOptionsModel := new(listsapiv1.CreateListItemsOptions)
				createListItemsOptionsModel.CreateListItemsReqItem = []listsapiv1.CreateListItemsReqItem{*createListItemsReqItemModel}
				createListItemsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := listsApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := listsApiService.CreateListItems(createListItemsOptionsModel)
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
			It(`Invoke CreateListItems successfully`, func() {
				listsApiService, serviceErr := listsapiv1.NewListsApiV1(&listsapiv1.ListsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ItemID:        core.StringPtr(itemID),
					ListID:        core.StringPtr(listID),
					OperationID:   core.StringPtr(operationID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(listsApiService).ToNot(BeNil())

				// Construct an instance of the CreateListItemsReqItem model
				createListItemsReqItemModel := new(listsapiv1.CreateListItemsReqItem)
				createListItemsReqItemModel.Asn = core.Float64Ptr(float64(19604))
				createListItemsReqItemModel.Comment = core.StringPtr("My list of developer IPs.")
				createListItemsReqItemModel.Hostname = core.StringPtr("cloud.ibm.com")
				createListItemsReqItemModel.Ip = core.StringPtr("172.64.0.0/13")

				// Construct an instance of the CreateListItemsOptions model
				createListItemsOptionsModel := new(listsapiv1.CreateListItemsOptions)
				createListItemsOptionsModel.CreateListItemsReqItem = []listsapiv1.CreateListItemsReqItem{*createListItemsReqItemModel}
				createListItemsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := listsApiService.CreateListItems(createListItemsOptionsModel)
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
	Describe(`DeleteListItems(deleteListItemsOptions *DeleteListItemsOptions) - Operation response error`, func() {
		crn := "testString"
		itemID := "testString"
		listID := "testString"
		operationID := "testString"
		deleteListItemsPath := "/v1/testString/rules/lists/testString/items"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteListItemsPath))
					Expect(req.Method).To(Equal("DELETE"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke DeleteListItems with error: Operation response processing error`, func() {
				listsApiService, serviceErr := listsapiv1.NewListsApiV1(&listsapiv1.ListsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ItemID:        core.StringPtr(itemID),
					ListID:        core.StringPtr(listID),
					OperationID:   core.StringPtr(operationID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(listsApiService).ToNot(BeNil())

				// Construct an instance of the DeleteListItemsReqItemsItem model
				deleteListItemsReqItemsItemModel := new(listsapiv1.DeleteListItemsReqItemsItem)
				deleteListItemsReqItemsItemModel.ID = core.StringPtr("70c2009751b24ffc9ed1ab462ba957b4")

				// Construct an instance of the DeleteListItemsOptions model
				deleteListItemsOptionsModel := new(listsapiv1.DeleteListItemsOptions)
				deleteListItemsOptionsModel.Items = []listsapiv1.DeleteListItemsReqItemsItem{*deleteListItemsReqItemsItemModel}
				deleteListItemsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := listsApiService.DeleteListItems(deleteListItemsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				listsApiService.EnableRetries(0, 0)
				result, response, operationErr = listsApiService.DeleteListItems(deleteListItemsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteListItems(deleteListItemsOptions *DeleteListItemsOptions)`, func() {
		crn := "testString"
		itemID := "testString"
		listID := "testString"
		operationID := "testString"
		deleteListItemsPath := "/v1/testString/rules/lists/testString/items"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteListItemsPath))
					Expect(req.Method).To(Equal("DELETE"))

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
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"operation_id": "53d73a83d33d4e3b8791764a9b9f2412"}}`)
				}))
			})
			It(`Invoke DeleteListItems successfully with retries`, func() {
				listsApiService, serviceErr := listsapiv1.NewListsApiV1(&listsapiv1.ListsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ItemID:        core.StringPtr(itemID),
					ListID:        core.StringPtr(listID),
					OperationID:   core.StringPtr(operationID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(listsApiService).ToNot(BeNil())
				listsApiService.EnableRetries(0, 0)

				// Construct an instance of the DeleteListItemsReqItemsItem model
				deleteListItemsReqItemsItemModel := new(listsapiv1.DeleteListItemsReqItemsItem)
				deleteListItemsReqItemsItemModel.ID = core.StringPtr("70c2009751b24ffc9ed1ab462ba957b4")

				// Construct an instance of the DeleteListItemsOptions model
				deleteListItemsOptionsModel := new(listsapiv1.DeleteListItemsOptions)
				deleteListItemsOptionsModel.Items = []listsapiv1.DeleteListItemsReqItemsItem{*deleteListItemsReqItemsItemModel}
				deleteListItemsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := listsApiService.DeleteListItemsWithContext(ctx, deleteListItemsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				listsApiService.DisableRetries()
				result, response, operationErr := listsApiService.DeleteListItems(deleteListItemsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = listsApiService.DeleteListItemsWithContext(ctx, deleteListItemsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(deleteListItemsPath))
					Expect(req.Method).To(Equal("DELETE"))

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
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"operation_id": "53d73a83d33d4e3b8791764a9b9f2412"}}`)
				}))
			})
			It(`Invoke DeleteListItems successfully`, func() {
				listsApiService, serviceErr := listsapiv1.NewListsApiV1(&listsapiv1.ListsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ItemID:        core.StringPtr(itemID),
					ListID:        core.StringPtr(listID),
					OperationID:   core.StringPtr(operationID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(listsApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := listsApiService.DeleteListItems(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DeleteListItemsReqItemsItem model
				deleteListItemsReqItemsItemModel := new(listsapiv1.DeleteListItemsReqItemsItem)
				deleteListItemsReqItemsItemModel.ID = core.StringPtr("70c2009751b24ffc9ed1ab462ba957b4")

				// Construct an instance of the DeleteListItemsOptions model
				deleteListItemsOptionsModel := new(listsapiv1.DeleteListItemsOptions)
				deleteListItemsOptionsModel.Items = []listsapiv1.DeleteListItemsReqItemsItem{*deleteListItemsReqItemsItemModel}
				deleteListItemsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = listsApiService.DeleteListItems(deleteListItemsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke DeleteListItems with error: Operation request error`, func() {
				listsApiService, serviceErr := listsapiv1.NewListsApiV1(&listsapiv1.ListsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ItemID:        core.StringPtr(itemID),
					ListID:        core.StringPtr(listID),
					OperationID:   core.StringPtr(operationID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(listsApiService).ToNot(BeNil())

				// Construct an instance of the DeleteListItemsReqItemsItem model
				deleteListItemsReqItemsItemModel := new(listsapiv1.DeleteListItemsReqItemsItem)
				deleteListItemsReqItemsItemModel.ID = core.StringPtr("70c2009751b24ffc9ed1ab462ba957b4")

				// Construct an instance of the DeleteListItemsOptions model
				deleteListItemsOptionsModel := new(listsapiv1.DeleteListItemsOptions)
				deleteListItemsOptionsModel.Items = []listsapiv1.DeleteListItemsReqItemsItem{*deleteListItemsReqItemsItemModel}
				deleteListItemsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := listsApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := listsApiService.DeleteListItems(deleteListItemsOptionsModel)
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
			It(`Invoke DeleteListItems successfully`, func() {
				listsApiService, serviceErr := listsapiv1.NewListsApiV1(&listsapiv1.ListsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ItemID:        core.StringPtr(itemID),
					ListID:        core.StringPtr(listID),
					OperationID:   core.StringPtr(operationID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(listsApiService).ToNot(BeNil())

				// Construct an instance of the DeleteListItemsReqItemsItem model
				deleteListItemsReqItemsItemModel := new(listsapiv1.DeleteListItemsReqItemsItem)
				deleteListItemsReqItemsItemModel.ID = core.StringPtr("70c2009751b24ffc9ed1ab462ba957b4")

				// Construct an instance of the DeleteListItemsOptions model
				deleteListItemsOptionsModel := new(listsapiv1.DeleteListItemsOptions)
				deleteListItemsOptionsModel.Items = []listsapiv1.DeleteListItemsReqItemsItem{*deleteListItemsReqItemsItemModel}
				deleteListItemsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := listsApiService.DeleteListItems(deleteListItemsOptionsModel)
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
	Describe(`UpdateListItems(updateListItemsOptions *UpdateListItemsOptions) - Operation response error`, func() {
		crn := "testString"
		itemID := "testString"
		listID := "testString"
		operationID := "testString"
		updateListItemsPath := "/v1/testString/rules/lists/testString/items"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateListItemsPath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateListItems with error: Operation response processing error`, func() {
				listsApiService, serviceErr := listsapiv1.NewListsApiV1(&listsapiv1.ListsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ItemID:        core.StringPtr(itemID),
					ListID:        core.StringPtr(listID),
					OperationID:   core.StringPtr(operationID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(listsApiService).ToNot(BeNil())

				// Construct an instance of the CreateListItemsReqItem model
				createListItemsReqItemModel := new(listsapiv1.CreateListItemsReqItem)
				createListItemsReqItemModel.Asn = core.Float64Ptr(float64(19604))
				createListItemsReqItemModel.Comment = core.StringPtr("My list of developer IPs.")
				createListItemsReqItemModel.Hostname = core.StringPtr("cloud.ibm.com")
				createListItemsReqItemModel.Ip = core.StringPtr("172.64.0.0/13")

				// Construct an instance of the UpdateListItemsOptions model
				updateListItemsOptionsModel := new(listsapiv1.UpdateListItemsOptions)
				updateListItemsOptionsModel.CreateListItemsReqItem = []listsapiv1.CreateListItemsReqItem{*createListItemsReqItemModel}
				updateListItemsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := listsApiService.UpdateListItems(updateListItemsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				listsApiService.EnableRetries(0, 0)
				result, response, operationErr = listsApiService.UpdateListItems(updateListItemsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateListItems(updateListItemsOptions *UpdateListItemsOptions)`, func() {
		crn := "testString"
		itemID := "testString"
		listID := "testString"
		operationID := "testString"
		updateListItemsPath := "/v1/testString/rules/lists/testString/items"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateListItemsPath))
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
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"operation_id": "53d73a83d33d4e3b8791764a9b9f2412"}}`)
				}))
			})
			It(`Invoke UpdateListItems successfully with retries`, func() {
				listsApiService, serviceErr := listsapiv1.NewListsApiV1(&listsapiv1.ListsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ItemID:        core.StringPtr(itemID),
					ListID:        core.StringPtr(listID),
					OperationID:   core.StringPtr(operationID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(listsApiService).ToNot(BeNil())
				listsApiService.EnableRetries(0, 0)

				// Construct an instance of the CreateListItemsReqItem model
				createListItemsReqItemModel := new(listsapiv1.CreateListItemsReqItem)
				createListItemsReqItemModel.Asn = core.Float64Ptr(float64(19604))
				createListItemsReqItemModel.Comment = core.StringPtr("My list of developer IPs.")
				createListItemsReqItemModel.Hostname = core.StringPtr("cloud.ibm.com")
				createListItemsReqItemModel.Ip = core.StringPtr("172.64.0.0/13")

				// Construct an instance of the UpdateListItemsOptions model
				updateListItemsOptionsModel := new(listsapiv1.UpdateListItemsOptions)
				updateListItemsOptionsModel.CreateListItemsReqItem = []listsapiv1.CreateListItemsReqItem{*createListItemsReqItemModel}
				updateListItemsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := listsApiService.UpdateListItemsWithContext(ctx, updateListItemsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				listsApiService.DisableRetries()
				result, response, operationErr := listsApiService.UpdateListItems(updateListItemsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = listsApiService.UpdateListItemsWithContext(ctx, updateListItemsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(updateListItemsPath))
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
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"operation_id": "53d73a83d33d4e3b8791764a9b9f2412"}}`)
				}))
			})
			It(`Invoke UpdateListItems successfully`, func() {
				listsApiService, serviceErr := listsapiv1.NewListsApiV1(&listsapiv1.ListsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ItemID:        core.StringPtr(itemID),
					ListID:        core.StringPtr(listID),
					OperationID:   core.StringPtr(operationID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(listsApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := listsApiService.UpdateListItems(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateListItemsReqItem model
				createListItemsReqItemModel := new(listsapiv1.CreateListItemsReqItem)
				createListItemsReqItemModel.Asn = core.Float64Ptr(float64(19604))
				createListItemsReqItemModel.Comment = core.StringPtr("My list of developer IPs.")
				createListItemsReqItemModel.Hostname = core.StringPtr("cloud.ibm.com")
				createListItemsReqItemModel.Ip = core.StringPtr("172.64.0.0/13")

				// Construct an instance of the UpdateListItemsOptions model
				updateListItemsOptionsModel := new(listsapiv1.UpdateListItemsOptions)
				updateListItemsOptionsModel.CreateListItemsReqItem = []listsapiv1.CreateListItemsReqItem{*createListItemsReqItemModel}
				updateListItemsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = listsApiService.UpdateListItems(updateListItemsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateListItems with error: Operation request error`, func() {
				listsApiService, serviceErr := listsapiv1.NewListsApiV1(&listsapiv1.ListsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ItemID:        core.StringPtr(itemID),
					ListID:        core.StringPtr(listID),
					OperationID:   core.StringPtr(operationID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(listsApiService).ToNot(BeNil())

				// Construct an instance of the CreateListItemsReqItem model
				createListItemsReqItemModel := new(listsapiv1.CreateListItemsReqItem)
				createListItemsReqItemModel.Asn = core.Float64Ptr(float64(19604))
				createListItemsReqItemModel.Comment = core.StringPtr("My list of developer IPs.")
				createListItemsReqItemModel.Hostname = core.StringPtr("cloud.ibm.com")
				createListItemsReqItemModel.Ip = core.StringPtr("172.64.0.0/13")

				// Construct an instance of the UpdateListItemsOptions model
				updateListItemsOptionsModel := new(listsapiv1.UpdateListItemsOptions)
				updateListItemsOptionsModel.CreateListItemsReqItem = []listsapiv1.CreateListItemsReqItem{*createListItemsReqItemModel}
				updateListItemsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := listsApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := listsApiService.UpdateListItems(updateListItemsOptionsModel)
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
			It(`Invoke UpdateListItems successfully`, func() {
				listsApiService, serviceErr := listsapiv1.NewListsApiV1(&listsapiv1.ListsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ItemID:        core.StringPtr(itemID),
					ListID:        core.StringPtr(listID),
					OperationID:   core.StringPtr(operationID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(listsApiService).ToNot(BeNil())

				// Construct an instance of the CreateListItemsReqItem model
				createListItemsReqItemModel := new(listsapiv1.CreateListItemsReqItem)
				createListItemsReqItemModel.Asn = core.Float64Ptr(float64(19604))
				createListItemsReqItemModel.Comment = core.StringPtr("My list of developer IPs.")
				createListItemsReqItemModel.Hostname = core.StringPtr("cloud.ibm.com")
				createListItemsReqItemModel.Ip = core.StringPtr("172.64.0.0/13")

				// Construct an instance of the UpdateListItemsOptions model
				updateListItemsOptionsModel := new(listsapiv1.UpdateListItemsOptions)
				updateListItemsOptionsModel.CreateListItemsReqItem = []listsapiv1.CreateListItemsReqItem{*createListItemsReqItemModel}
				updateListItemsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := listsApiService.UpdateListItems(updateListItemsOptionsModel)
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
	Describe(`GetListItem(getListItemOptions *GetListItemOptions) - Operation response error`, func() {
		crn := "testString"
		itemID := "testString"
		listID := "testString"
		operationID := "testString"
		getListItemPath := "/v1/testString/rules/lists/testString/items/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getListItemPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetListItem with error: Operation response processing error`, func() {
				listsApiService, serviceErr := listsapiv1.NewListsApiV1(&listsapiv1.ListsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ItemID:        core.StringPtr(itemID),
					ListID:        core.StringPtr(listID),
					OperationID:   core.StringPtr(operationID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(listsApiService).ToNot(BeNil())

				// Construct an instance of the GetListItemOptions model
				getListItemOptionsModel := new(listsapiv1.GetListItemOptions)
				getListItemOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := listsApiService.GetListItem(getListItemOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				listsApiService.EnableRetries(0, 0)
				result, response, operationErr = listsApiService.GetListItem(getListItemOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetListItem(getListItemOptions *GetListItemOptions)`, func() {
		crn := "testString"
		itemID := "testString"
		listID := "testString"
		operationID := "testString"
		getListItemPath := "/v1/testString/rules/lists/testString/items/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getListItemPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": "70c2009751b24ffc9ed1ab462ba957b4", "asn": 19604, "comment": "My list of developer IPs.", "hostname": "cloud.ibm.com", "ip": "172.64.0.0/13", "created_on": "2025-03-21T16:19:21Z", "modified_on": "2025-03-21T16:19:37Z"}}`)
				}))
			})
			It(`Invoke GetListItem successfully with retries`, func() {
				listsApiService, serviceErr := listsapiv1.NewListsApiV1(&listsapiv1.ListsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ItemID:        core.StringPtr(itemID),
					ListID:        core.StringPtr(listID),
					OperationID:   core.StringPtr(operationID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(listsApiService).ToNot(BeNil())
				listsApiService.EnableRetries(0, 0)

				// Construct an instance of the GetListItemOptions model
				getListItemOptionsModel := new(listsapiv1.GetListItemOptions)
				getListItemOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := listsApiService.GetListItemWithContext(ctx, getListItemOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				listsApiService.DisableRetries()
				result, response, operationErr := listsApiService.GetListItem(getListItemOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = listsApiService.GetListItemWithContext(ctx, getListItemOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getListItemPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": "70c2009751b24ffc9ed1ab462ba957b4", "asn": 19604, "comment": "My list of developer IPs.", "hostname": "cloud.ibm.com", "ip": "172.64.0.0/13", "created_on": "2025-03-21T16:19:21Z", "modified_on": "2025-03-21T16:19:37Z"}}`)
				}))
			})
			It(`Invoke GetListItem successfully`, func() {
				listsApiService, serviceErr := listsapiv1.NewListsApiV1(&listsapiv1.ListsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ItemID:        core.StringPtr(itemID),
					ListID:        core.StringPtr(listID),
					OperationID:   core.StringPtr(operationID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(listsApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := listsApiService.GetListItem(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetListItemOptions model
				getListItemOptionsModel := new(listsapiv1.GetListItemOptions)
				getListItemOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = listsApiService.GetListItem(getListItemOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetListItem with error: Operation request error`, func() {
				listsApiService, serviceErr := listsapiv1.NewListsApiV1(&listsapiv1.ListsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ItemID:        core.StringPtr(itemID),
					ListID:        core.StringPtr(listID),
					OperationID:   core.StringPtr(operationID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(listsApiService).ToNot(BeNil())

				// Construct an instance of the GetListItemOptions model
				getListItemOptionsModel := new(listsapiv1.GetListItemOptions)
				getListItemOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := listsApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := listsApiService.GetListItem(getListItemOptionsModel)
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
			It(`Invoke GetListItem successfully`, func() {
				listsApiService, serviceErr := listsapiv1.NewListsApiV1(&listsapiv1.ListsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ItemID:        core.StringPtr(itemID),
					ListID:        core.StringPtr(listID),
					OperationID:   core.StringPtr(operationID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(listsApiService).ToNot(BeNil())

				// Construct an instance of the GetListItemOptions model
				getListItemOptionsModel := new(listsapiv1.GetListItemOptions)
				getListItemOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := listsApiService.GetListItem(getListItemOptionsModel)
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
	Describe(`GetOperationStatus(getOperationStatusOptions *GetOperationStatusOptions) - Operation response error`, func() {
		crn := "testString"
		itemID := "testString"
		listID := "testString"
		operationID := "testString"
		getOperationStatusPath := "/v1/testString/rules/lists/bulk_operations/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getOperationStatusPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetOperationStatus with error: Operation response processing error`, func() {
				listsApiService, serviceErr := listsapiv1.NewListsApiV1(&listsapiv1.ListsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ItemID:        core.StringPtr(itemID),
					ListID:        core.StringPtr(listID),
					OperationID:   core.StringPtr(operationID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(listsApiService).ToNot(BeNil())

				// Construct an instance of the GetOperationStatusOptions model
				getOperationStatusOptionsModel := new(listsapiv1.GetOperationStatusOptions)
				getOperationStatusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := listsApiService.GetOperationStatus(getOperationStatusOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				listsApiService.EnableRetries(0, 0)
				result, response, operationErr = listsApiService.GetOperationStatus(getOperationStatusOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetOperationStatus(getOperationStatusOptions *GetOperationStatusOptions)`, func() {
		crn := "testString"
		itemID := "testString"
		listID := "testString"
		operationID := "testString"
		getOperationStatusPath := "/v1/testString/rules/lists/bulk_operations/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getOperationStatusPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": "0147be950d5c42b8b47c07792c5015e3", "status": "completed", "completed": "2025-03-21T16:07:41.782564Z", "error": "Error"}}`)
				}))
			})
			It(`Invoke GetOperationStatus successfully with retries`, func() {
				listsApiService, serviceErr := listsapiv1.NewListsApiV1(&listsapiv1.ListsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ItemID:        core.StringPtr(itemID),
					ListID:        core.StringPtr(listID),
					OperationID:   core.StringPtr(operationID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(listsApiService).ToNot(BeNil())
				listsApiService.EnableRetries(0, 0)

				// Construct an instance of the GetOperationStatusOptions model
				getOperationStatusOptionsModel := new(listsapiv1.GetOperationStatusOptions)
				getOperationStatusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := listsApiService.GetOperationStatusWithContext(ctx, getOperationStatusOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				listsApiService.DisableRetries()
				result, response, operationErr := listsApiService.GetOperationStatus(getOperationStatusOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = listsApiService.GetOperationStatusWithContext(ctx, getOperationStatusOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getOperationStatusPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": "0147be950d5c42b8b47c07792c5015e3", "status": "completed", "completed": "2025-03-21T16:07:41.782564Z", "error": "Error"}}`)
				}))
			})
			It(`Invoke GetOperationStatus successfully`, func() {
				listsApiService, serviceErr := listsapiv1.NewListsApiV1(&listsapiv1.ListsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ItemID:        core.StringPtr(itemID),
					ListID:        core.StringPtr(listID),
					OperationID:   core.StringPtr(operationID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(listsApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := listsApiService.GetOperationStatus(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetOperationStatusOptions model
				getOperationStatusOptionsModel := new(listsapiv1.GetOperationStatusOptions)
				getOperationStatusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = listsApiService.GetOperationStatus(getOperationStatusOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetOperationStatus with error: Operation request error`, func() {
				listsApiService, serviceErr := listsapiv1.NewListsApiV1(&listsapiv1.ListsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ItemID:        core.StringPtr(itemID),
					ListID:        core.StringPtr(listID),
					OperationID:   core.StringPtr(operationID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(listsApiService).ToNot(BeNil())

				// Construct an instance of the GetOperationStatusOptions model
				getOperationStatusOptionsModel := new(listsapiv1.GetOperationStatusOptions)
				getOperationStatusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := listsApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := listsApiService.GetOperationStatus(getOperationStatusOptionsModel)
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
			It(`Invoke GetOperationStatus successfully`, func() {
				listsApiService, serviceErr := listsapiv1.NewListsApiV1(&listsapiv1.ListsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ItemID:        core.StringPtr(itemID),
					ListID:        core.StringPtr(listID),
					OperationID:   core.StringPtr(operationID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(listsApiService).ToNot(BeNil())

				// Construct an instance of the GetOperationStatusOptions model
				getOperationStatusOptionsModel := new(listsapiv1.GetOperationStatusOptions)
				getOperationStatusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := listsApiService.GetOperationStatus(getOperationStatusOptionsModel)
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
			itemID := "testString"
			listID := "testString"
			operationID := "testString"
			listsApiService, _ := listsapiv1.NewListsApiV1(&listsapiv1.ListsApiV1Options{
				URL:           "http://listsapiv1modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
				Crn:           core.StringPtr(crn),
				ItemID:        core.StringPtr(itemID),
				ListID:        core.StringPtr(listID),
				OperationID:   core.StringPtr(operationID),
			})
			It(`Invoke NewCreateCustomListsOptions successfully`, func() {
				// Construct an instance of the CreateCustomListsOptions model
				createCustomListsOptionsModel := listsApiService.NewCreateCustomListsOptions()
				createCustomListsOptionsModel.SetKind("ip")
				createCustomListsOptionsModel.SetName("testString")
				createCustomListsOptionsModel.SetDescription("testString")
				createCustomListsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createCustomListsOptionsModel).ToNot(BeNil())
				Expect(createCustomListsOptionsModel.Kind).To(Equal(core.StringPtr("ip")))
				Expect(createCustomListsOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(createCustomListsOptionsModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(createCustomListsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateListItemsOptions successfully`, func() {
				// Construct an instance of the CreateListItemsReqItem model
				createListItemsReqItemModel := new(listsapiv1.CreateListItemsReqItem)
				Expect(createListItemsReqItemModel).ToNot(BeNil())
				createListItemsReqItemModel.Asn = core.Float64Ptr(float64(19604))
				createListItemsReqItemModel.Comment = core.StringPtr("My list of developer IPs.")
				createListItemsReqItemModel.Hostname = core.StringPtr("cloud.ibm.com")
				createListItemsReqItemModel.Ip = core.StringPtr("172.64.0.0/13")
				Expect(createListItemsReqItemModel.Asn).To(Equal(core.Float64Ptr(float64(19604))))
				Expect(createListItemsReqItemModel.Comment).To(Equal(core.StringPtr("My list of developer IPs.")))
				Expect(createListItemsReqItemModel.Hostname).To(Equal(core.StringPtr("cloud.ibm.com")))
				Expect(createListItemsReqItemModel.Ip).To(Equal(core.StringPtr("172.64.0.0/13")))

				// Construct an instance of the CreateListItemsOptions model
				createListItemsOptionsModel := listsApiService.NewCreateListItemsOptions()
				createListItemsOptionsModel.SetCreateListItemsReqItem([]listsapiv1.CreateListItemsReqItem{*createListItemsReqItemModel})
				createListItemsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createListItemsOptionsModel).ToNot(BeNil())
				Expect(createListItemsOptionsModel.CreateListItemsReqItem).To(Equal([]listsapiv1.CreateListItemsReqItem{*createListItemsReqItemModel}))
				Expect(createListItemsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteCustomListOptions successfully`, func() {
				// Construct an instance of the DeleteCustomListOptions model
				deleteCustomListOptionsModel := listsApiService.NewDeleteCustomListOptions()
				deleteCustomListOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteCustomListOptionsModel).ToNot(BeNil())
				Expect(deleteCustomListOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteListItemsOptions successfully`, func() {
				// Construct an instance of the DeleteListItemsReqItemsItem model
				deleteListItemsReqItemsItemModel := new(listsapiv1.DeleteListItemsReqItemsItem)
				Expect(deleteListItemsReqItemsItemModel).ToNot(BeNil())
				deleteListItemsReqItemsItemModel.ID = core.StringPtr("70c2009751b24ffc9ed1ab462ba957b4")
				Expect(deleteListItemsReqItemsItemModel.ID).To(Equal(core.StringPtr("70c2009751b24ffc9ed1ab462ba957b4")))

				// Construct an instance of the DeleteListItemsOptions model
				deleteListItemsOptionsModel := listsApiService.NewDeleteListItemsOptions()
				deleteListItemsOptionsModel.SetItems([]listsapiv1.DeleteListItemsReqItemsItem{*deleteListItemsReqItemsItemModel})
				deleteListItemsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteListItemsOptionsModel).ToNot(BeNil())
				Expect(deleteListItemsOptionsModel.Items).To(Equal([]listsapiv1.DeleteListItemsReqItemsItem{*deleteListItemsReqItemsItemModel}))
				Expect(deleteListItemsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetCustomListOptions successfully`, func() {
				// Construct an instance of the GetCustomListOptions model
				getCustomListOptionsModel := listsApiService.NewGetCustomListOptions()
				getCustomListOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getCustomListOptionsModel).ToNot(BeNil())
				Expect(getCustomListOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetCustomListsOptions successfully`, func() {
				// Construct an instance of the GetCustomListsOptions model
				getCustomListsOptionsModel := listsApiService.NewGetCustomListsOptions()
				getCustomListsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getCustomListsOptionsModel).ToNot(BeNil())
				Expect(getCustomListsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetListItemOptions successfully`, func() {
				// Construct an instance of the GetListItemOptions model
				getListItemOptionsModel := listsApiService.NewGetListItemOptions()
				getListItemOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getListItemOptionsModel).ToNot(BeNil())
				Expect(getListItemOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetListItemsOptions successfully`, func() {
				// Construct an instance of the GetListItemsOptions model
				getListItemsOptionsModel := listsApiService.NewGetListItemsOptions()
				getListItemsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getListItemsOptionsModel).ToNot(BeNil())
				Expect(getListItemsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetManagedListsOptions successfully`, func() {
				// Construct an instance of the GetManagedListsOptions model
				getManagedListsOptionsModel := listsApiService.NewGetManagedListsOptions()
				getManagedListsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getManagedListsOptionsModel).ToNot(BeNil())
				Expect(getManagedListsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetOperationStatusOptions successfully`, func() {
				// Construct an instance of the GetOperationStatusOptions model
				getOperationStatusOptionsModel := listsApiService.NewGetOperationStatusOptions()
				getOperationStatusOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getOperationStatusOptionsModel).ToNot(BeNil())
				Expect(getOperationStatusOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateCustomListOptions successfully`, func() {
				// Construct an instance of the UpdateCustomListOptions model
				updateCustomListOptionsModel := listsApiService.NewUpdateCustomListOptions()
				updateCustomListOptionsModel.SetDescription("testString")
				updateCustomListOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateCustomListOptionsModel).ToNot(BeNil())
				Expect(updateCustomListOptionsModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(updateCustomListOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateListItemsOptions successfully`, func() {
				// Construct an instance of the CreateListItemsReqItem model
				createListItemsReqItemModel := new(listsapiv1.CreateListItemsReqItem)
				Expect(createListItemsReqItemModel).ToNot(BeNil())
				createListItemsReqItemModel.Asn = core.Float64Ptr(float64(19604))
				createListItemsReqItemModel.Comment = core.StringPtr("My list of developer IPs.")
				createListItemsReqItemModel.Hostname = core.StringPtr("cloud.ibm.com")
				createListItemsReqItemModel.Ip = core.StringPtr("172.64.0.0/13")
				Expect(createListItemsReqItemModel.Asn).To(Equal(core.Float64Ptr(float64(19604))))
				Expect(createListItemsReqItemModel.Comment).To(Equal(core.StringPtr("My list of developer IPs.")))
				Expect(createListItemsReqItemModel.Hostname).To(Equal(core.StringPtr("cloud.ibm.com")))
				Expect(createListItemsReqItemModel.Ip).To(Equal(core.StringPtr("172.64.0.0/13")))

				// Construct an instance of the UpdateListItemsOptions model
				updateListItemsOptionsModel := listsApiService.NewUpdateListItemsOptions()
				updateListItemsOptionsModel.SetCreateListItemsReqItem([]listsapiv1.CreateListItemsReqItem{*createListItemsReqItemModel})
				updateListItemsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateListItemsOptionsModel).ToNot(BeNil())
				Expect(updateListItemsOptionsModel.CreateListItemsReqItem).To(Equal([]listsapiv1.CreateListItemsReqItem{*createListItemsReqItemModel}))
				Expect(updateListItemsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
		})
	})
	Describe(`Model unmarshaling tests`, func() {
		It(`Invoke UnmarshalCreateListItemsReqItem successfully`, func() {
			// Construct an instance of the model.
			model := new(listsapiv1.CreateListItemsReqItem)
			model.Asn = core.Float64Ptr(float64(19604))
			model.Comment = core.StringPtr("My list of developer IPs.")
			model.Hostname = core.StringPtr("cloud.ibm.com")
			model.Ip = core.StringPtr("172.64.0.0/13")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *listsapiv1.CreateListItemsReqItem
			err = listsapiv1.UnmarshalCreateListItemsReqItem(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalDeleteListItemsReqItemsItem successfully`, func() {
			// Construct an instance of the model.
			model := new(listsapiv1.DeleteListItemsReqItemsItem)
			model.ID = core.StringPtr("70c2009751b24ffc9ed1ab462ba957b4")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *listsapiv1.DeleteListItemsReqItemsItem
			err = listsapiv1.UnmarshalDeleteListItemsReqItemsItem(raw, &result)
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
