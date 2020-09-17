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

package edgefunctionsapiv1_test

import (
	"bytes"
	"fmt"
	"github.com/IBM/go-sdk-core/v4/core"
	"github.com/IBM/networking-go-sdk/edgefunctionsapiv1"
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

var _ = Describe(`EdgeFunctionsApiV1`, func() {
	var testServer *httptest.Server
    Describe(`Service constructor tests`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		It(`Instantiate service client`, func() {
			testService, testServiceErr := edgefunctionsapiv1.NewEdgeFunctionsApiV1(&edgefunctionsapiv1.EdgeFunctionsApiV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
				Crn: core.StringPtr(crn),
				ZoneIdentifier: core.StringPtr(zoneIdentifier),
			})
			Expect(testService).ToNot(BeNil())
			Expect(testServiceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			testService, testServiceErr := edgefunctionsapiv1.NewEdgeFunctionsApiV1(&edgefunctionsapiv1.EdgeFunctionsApiV1Options{
				URL: "{BAD_URL_STRING",
				Crn: core.StringPtr(crn),
				ZoneIdentifier: core.StringPtr(zoneIdentifier),
			})
			Expect(testService).To(BeNil())
			Expect(testServiceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			testService, testServiceErr := edgefunctionsapiv1.NewEdgeFunctionsApiV1(&edgefunctionsapiv1.EdgeFunctionsApiV1Options{
				URL: "https://edgefunctionsapiv1/api",
				Crn: core.StringPtr(crn),
				ZoneIdentifier: core.StringPtr(zoneIdentifier),
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(testService).To(BeNil())
			Expect(testServiceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Validation Error`, func() {
			testService, testServiceErr := edgefunctionsapiv1.NewEdgeFunctionsApiV1(&edgefunctionsapiv1.EdgeFunctionsApiV1Options{})
			Expect(testService).To(BeNil())
			Expect(testServiceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"EDGE_FUNCTIONS_API_URL": "https://edgefunctionsapiv1/api",
				"EDGE_FUNCTIONS_API_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := edgefunctionsapiv1.NewEdgeFunctionsApiV1UsingExternalConfig(&edgefunctionsapiv1.EdgeFunctionsApiV1Options{
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := edgefunctionsapiv1.NewEdgeFunctionsApiV1UsingExternalConfig(&edgefunctionsapiv1.EdgeFunctionsApiV1Options{
					URL: "https://testService/api",
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				Expect(testService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := edgefunctionsapiv1.NewEdgeFunctionsApiV1UsingExternalConfig(&edgefunctionsapiv1.EdgeFunctionsApiV1Options{
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				err := testService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				Expect(testService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"EDGE_FUNCTIONS_API_URL": "https://edgefunctionsapiv1/api",
				"EDGE_FUNCTIONS_API_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			testService, testServiceErr := edgefunctionsapiv1.NewEdgeFunctionsApiV1UsingExternalConfig(&edgefunctionsapiv1.EdgeFunctionsApiV1Options{
				Crn: core.StringPtr(crn),
				ZoneIdentifier: core.StringPtr(zoneIdentifier),
			})

			It(`Instantiate service client with error`, func() {
				Expect(testService).To(BeNil())
				Expect(testServiceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"EDGE_FUNCTIONS_API_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			testService, testServiceErr := edgefunctionsapiv1.NewEdgeFunctionsApiV1UsingExternalConfig(&edgefunctionsapiv1.EdgeFunctionsApiV1Options{
				URL: "{BAD_URL_STRING",
				Crn: core.StringPtr(crn),
				ZoneIdentifier: core.StringPtr(zoneIdentifier),
			})

			It(`Instantiate service client with error`, func() {
				Expect(testService).To(BeNil())
				Expect(testServiceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`ListEdgeFunctionsActions(listEdgeFunctionsActionsOptions *ListEdgeFunctionsActionsOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		listEdgeFunctionsActionsPath := "/v1/testString/workers/scripts"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(listEdgeFunctionsActionsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListEdgeFunctionsActions with error: Operation response processing error`, func() {
				testService, testServiceErr := edgefunctionsapiv1.NewEdgeFunctionsApiV1(&edgefunctionsapiv1.EdgeFunctionsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ListEdgeFunctionsActionsOptions model
				listEdgeFunctionsActionsOptionsModel := new(edgefunctionsapiv1.ListEdgeFunctionsActionsOptions)
				listEdgeFunctionsActionsOptionsModel.XCorrelationID = core.StringPtr("testString")
				listEdgeFunctionsActionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.ListEdgeFunctionsActions(listEdgeFunctionsActionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ListEdgeFunctionsActions(listEdgeFunctionsActionsOptions *ListEdgeFunctionsActionsOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		listEdgeFunctionsActionsPath := "/v1/testString/workers/scripts"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(listEdgeFunctionsActionsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"result": [{"script": "addEventListener('fetch', event => { event.respondWith(fetch(event.request)) })", "etag": "ea95132c15732412d22c1476fa83f27a", "handlers": ["fetch"], "modified_on": "2019-01-01T12:00:00", "created_on": "2019-01-01T12:00:00", "routes": [{"id": "9a7806061c88ada191ed06f989cc3dac", "pattern": "example.net/*", "script": "this-is_my_script-01", "request_limit_fail_open": false}]}], "success": true, "errors": ["Errors"], "messages": ["Messages"]}`)
				}))
			})
			It(`Invoke ListEdgeFunctionsActions successfully`, func() {
				testService, testServiceErr := edgefunctionsapiv1.NewEdgeFunctionsApiV1(&edgefunctionsapiv1.EdgeFunctionsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.ListEdgeFunctionsActions(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListEdgeFunctionsActionsOptions model
				listEdgeFunctionsActionsOptionsModel := new(edgefunctionsapiv1.ListEdgeFunctionsActionsOptions)
				listEdgeFunctionsActionsOptionsModel.XCorrelationID = core.StringPtr("testString")
 				listEdgeFunctionsActionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.ListEdgeFunctionsActions(listEdgeFunctionsActionsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke ListEdgeFunctionsActions with error: Operation request error`, func() {
				testService, testServiceErr := edgefunctionsapiv1.NewEdgeFunctionsApiV1(&edgefunctionsapiv1.EdgeFunctionsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ListEdgeFunctionsActionsOptions model
				listEdgeFunctionsActionsOptionsModel := new(edgefunctionsapiv1.ListEdgeFunctionsActionsOptions)
				listEdgeFunctionsActionsOptionsModel.XCorrelationID = core.StringPtr("testString")
				listEdgeFunctionsActionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.ListEdgeFunctionsActions(listEdgeFunctionsActionsOptionsModel)
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
	Describe(`UpdateEdgeFunctionsAction(updateEdgeFunctionsActionOptions *UpdateEdgeFunctionsActionOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		updateEdgeFunctionsActionPath := "/v1/testString/workers/scripts/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(updateEdgeFunctionsActionPath))
					Expect(req.Method).To(Equal("PUT"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateEdgeFunctionsAction with error: Operation response processing error`, func() {
				testService, testServiceErr := edgefunctionsapiv1.NewEdgeFunctionsApiV1(&edgefunctionsapiv1.EdgeFunctionsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the UpdateEdgeFunctionsActionOptions model
				updateEdgeFunctionsActionOptionsModel := new(edgefunctionsapiv1.UpdateEdgeFunctionsActionOptions)
				updateEdgeFunctionsActionOptionsModel.ScriptName = core.StringPtr("testString")
				updateEdgeFunctionsActionOptionsModel.EdgeFunctionsAction = CreateMockReader("This is a mock file.")
				updateEdgeFunctionsActionOptionsModel.XCorrelationID = core.StringPtr("testString")
				updateEdgeFunctionsActionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.UpdateEdgeFunctionsAction(updateEdgeFunctionsActionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UpdateEdgeFunctionsAction(updateEdgeFunctionsActionOptions *UpdateEdgeFunctionsActionOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		updateEdgeFunctionsActionPath := "/v1/testString/workers/scripts/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(updateEdgeFunctionsActionPath))
					Expect(req.Method).To(Equal("PUT"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"result": {"script": "addEventListener('fetch', event => { event.respondWith(fetch(event.request)) })", "etag": "ea95132c15732412d22c1476fa83f27a", "handlers": ["fetch"], "modified_on": "2019-01-01T12:00:00", "created_on": "2019-01-01T12:00:00", "routes": [{"id": "9a7806061c88ada191ed06f989cc3dac", "pattern": "example.net/*", "script": "this-is_my_script-01", "request_limit_fail_open": false}]}, "success": true, "errors": ["Errors"], "messages": ["Messages"]}`)
				}))
			})
			It(`Invoke UpdateEdgeFunctionsAction successfully`, func() {
				testService, testServiceErr := edgefunctionsapiv1.NewEdgeFunctionsApiV1(&edgefunctionsapiv1.EdgeFunctionsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.UpdateEdgeFunctionsAction(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateEdgeFunctionsActionOptions model
				updateEdgeFunctionsActionOptionsModel := new(edgefunctionsapiv1.UpdateEdgeFunctionsActionOptions)
				updateEdgeFunctionsActionOptionsModel.ScriptName = core.StringPtr("testString")
				updateEdgeFunctionsActionOptionsModel.EdgeFunctionsAction = CreateMockReader("This is a mock file.")
				updateEdgeFunctionsActionOptionsModel.XCorrelationID = core.StringPtr("testString")
 				updateEdgeFunctionsActionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.UpdateEdgeFunctionsAction(updateEdgeFunctionsActionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke UpdateEdgeFunctionsAction with error: Operation validation and request error`, func() {
				testService, testServiceErr := edgefunctionsapiv1.NewEdgeFunctionsApiV1(&edgefunctionsapiv1.EdgeFunctionsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the UpdateEdgeFunctionsActionOptions model
				updateEdgeFunctionsActionOptionsModel := new(edgefunctionsapiv1.UpdateEdgeFunctionsActionOptions)
				updateEdgeFunctionsActionOptionsModel.ScriptName = core.StringPtr("testString")
				updateEdgeFunctionsActionOptionsModel.EdgeFunctionsAction = CreateMockReader("This is a mock file.")
				updateEdgeFunctionsActionOptionsModel.XCorrelationID = core.StringPtr("testString")
				updateEdgeFunctionsActionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.UpdateEdgeFunctionsAction(updateEdgeFunctionsActionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateEdgeFunctionsActionOptions model with no property values
				updateEdgeFunctionsActionOptionsModelNew := new(edgefunctionsapiv1.UpdateEdgeFunctionsActionOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.UpdateEdgeFunctionsAction(updateEdgeFunctionsActionOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetEdgeFunctionsAction(getEdgeFunctionsActionOptions *GetEdgeFunctionsActionOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getEdgeFunctionsActionPath := "/v1/testString/workers/scripts/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getEdgeFunctionsActionPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/javascript")
					res.WriteHeader(200)
					fmt.Fprintf(res, `"unknown property type: OperationResponse"`)
				}))
			})
			It(`Invoke GetEdgeFunctionsAction successfully`, func() {
				testService, testServiceErr := edgefunctionsapiv1.NewEdgeFunctionsApiV1(&edgefunctionsapiv1.EdgeFunctionsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetEdgeFunctionsAction(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetEdgeFunctionsActionOptions model
				getEdgeFunctionsActionOptionsModel := new(edgefunctionsapiv1.GetEdgeFunctionsActionOptions)
				getEdgeFunctionsActionOptionsModel.ScriptName = core.StringPtr("testString")
				getEdgeFunctionsActionOptionsModel.XCorrelationID = core.StringPtr("testString")
 				getEdgeFunctionsActionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetEdgeFunctionsAction(getEdgeFunctionsActionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke GetEdgeFunctionsAction with error: Operation validation and request error`, func() {
				testService, testServiceErr := edgefunctionsapiv1.NewEdgeFunctionsApiV1(&edgefunctionsapiv1.EdgeFunctionsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetEdgeFunctionsActionOptions model
				getEdgeFunctionsActionOptionsModel := new(edgefunctionsapiv1.GetEdgeFunctionsActionOptions)
				getEdgeFunctionsActionOptionsModel.ScriptName = core.StringPtr("testString")
				getEdgeFunctionsActionOptionsModel.XCorrelationID = core.StringPtr("testString")
				getEdgeFunctionsActionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.GetEdgeFunctionsAction(getEdgeFunctionsActionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetEdgeFunctionsActionOptions model with no property values
				getEdgeFunctionsActionOptionsModelNew := new(edgefunctionsapiv1.GetEdgeFunctionsActionOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.GetEdgeFunctionsAction(getEdgeFunctionsActionOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteEdgeFunctionsAction(deleteEdgeFunctionsActionOptions *DeleteEdgeFunctionsActionOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		deleteEdgeFunctionsActionPath := "/v1/testString/workers/scripts/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(deleteEdgeFunctionsActionPath))
					Expect(req.Method).To(Equal("DELETE"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke DeleteEdgeFunctionsAction with error: Operation response processing error`, func() {
				testService, testServiceErr := edgefunctionsapiv1.NewEdgeFunctionsApiV1(&edgefunctionsapiv1.EdgeFunctionsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the DeleteEdgeFunctionsActionOptions model
				deleteEdgeFunctionsActionOptionsModel := new(edgefunctionsapiv1.DeleteEdgeFunctionsActionOptions)
				deleteEdgeFunctionsActionOptionsModel.ScriptName = core.StringPtr("testString")
				deleteEdgeFunctionsActionOptionsModel.XCorrelationID = core.StringPtr("testString")
				deleteEdgeFunctionsActionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.DeleteEdgeFunctionsAction(deleteEdgeFunctionsActionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`DeleteEdgeFunctionsAction(deleteEdgeFunctionsActionOptions *DeleteEdgeFunctionsActionOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		deleteEdgeFunctionsActionPath := "/v1/testString/workers/scripts/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(deleteEdgeFunctionsActionPath))
					Expect(req.Method).To(Equal("DELETE"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"result": {"id": "9a7806061c88ada191ed06f989cc3dac"}, "success": true, "errors": ["Errors"], "messages": ["Messages"]}`)
				}))
			})
			It(`Invoke DeleteEdgeFunctionsAction successfully`, func() {
				testService, testServiceErr := edgefunctionsapiv1.NewEdgeFunctionsApiV1(&edgefunctionsapiv1.EdgeFunctionsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.DeleteEdgeFunctionsAction(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DeleteEdgeFunctionsActionOptions model
				deleteEdgeFunctionsActionOptionsModel := new(edgefunctionsapiv1.DeleteEdgeFunctionsActionOptions)
				deleteEdgeFunctionsActionOptionsModel.ScriptName = core.StringPtr("testString")
				deleteEdgeFunctionsActionOptionsModel.XCorrelationID = core.StringPtr("testString")
 				deleteEdgeFunctionsActionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.DeleteEdgeFunctionsAction(deleteEdgeFunctionsActionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke DeleteEdgeFunctionsAction with error: Operation validation and request error`, func() {
				testService, testServiceErr := edgefunctionsapiv1.NewEdgeFunctionsApiV1(&edgefunctionsapiv1.EdgeFunctionsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the DeleteEdgeFunctionsActionOptions model
				deleteEdgeFunctionsActionOptionsModel := new(edgefunctionsapiv1.DeleteEdgeFunctionsActionOptions)
				deleteEdgeFunctionsActionOptionsModel.ScriptName = core.StringPtr("testString")
				deleteEdgeFunctionsActionOptionsModel.XCorrelationID = core.StringPtr("testString")
				deleteEdgeFunctionsActionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.DeleteEdgeFunctionsAction(deleteEdgeFunctionsActionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the DeleteEdgeFunctionsActionOptions model with no property values
				deleteEdgeFunctionsActionOptionsModelNew := new(edgefunctionsapiv1.DeleteEdgeFunctionsActionOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.DeleteEdgeFunctionsAction(deleteEdgeFunctionsActionOptionsModelNew)
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
		crn := "testString"
		zoneIdentifier := "testString"
		It(`Instantiate service client`, func() {
			testService, testServiceErr := edgefunctionsapiv1.NewEdgeFunctionsApiV1(&edgefunctionsapiv1.EdgeFunctionsApiV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
				Crn: core.StringPtr(crn),
				ZoneIdentifier: core.StringPtr(zoneIdentifier),
			})
			Expect(testService).ToNot(BeNil())
			Expect(testServiceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			testService, testServiceErr := edgefunctionsapiv1.NewEdgeFunctionsApiV1(&edgefunctionsapiv1.EdgeFunctionsApiV1Options{
				URL: "{BAD_URL_STRING",
				Crn: core.StringPtr(crn),
				ZoneIdentifier: core.StringPtr(zoneIdentifier),
			})
			Expect(testService).To(BeNil())
			Expect(testServiceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			testService, testServiceErr := edgefunctionsapiv1.NewEdgeFunctionsApiV1(&edgefunctionsapiv1.EdgeFunctionsApiV1Options{
				URL: "https://edgefunctionsapiv1/api",
				Crn: core.StringPtr(crn),
				ZoneIdentifier: core.StringPtr(zoneIdentifier),
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(testService).To(BeNil())
			Expect(testServiceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Validation Error`, func() {
			testService, testServiceErr := edgefunctionsapiv1.NewEdgeFunctionsApiV1(&edgefunctionsapiv1.EdgeFunctionsApiV1Options{})
			Expect(testService).To(BeNil())
			Expect(testServiceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"EDGE_FUNCTIONS_API_URL": "https://edgefunctionsapiv1/api",
				"EDGE_FUNCTIONS_API_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := edgefunctionsapiv1.NewEdgeFunctionsApiV1UsingExternalConfig(&edgefunctionsapiv1.EdgeFunctionsApiV1Options{
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := edgefunctionsapiv1.NewEdgeFunctionsApiV1UsingExternalConfig(&edgefunctionsapiv1.EdgeFunctionsApiV1Options{
					URL: "https://testService/api",
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				Expect(testService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := edgefunctionsapiv1.NewEdgeFunctionsApiV1UsingExternalConfig(&edgefunctionsapiv1.EdgeFunctionsApiV1Options{
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				err := testService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				Expect(testService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"EDGE_FUNCTIONS_API_URL": "https://edgefunctionsapiv1/api",
				"EDGE_FUNCTIONS_API_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			testService, testServiceErr := edgefunctionsapiv1.NewEdgeFunctionsApiV1UsingExternalConfig(&edgefunctionsapiv1.EdgeFunctionsApiV1Options{
				Crn: core.StringPtr(crn),
				ZoneIdentifier: core.StringPtr(zoneIdentifier),
			})

			It(`Instantiate service client with error`, func() {
				Expect(testService).To(BeNil())
				Expect(testServiceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"EDGE_FUNCTIONS_API_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			testService, testServiceErr := edgefunctionsapiv1.NewEdgeFunctionsApiV1UsingExternalConfig(&edgefunctionsapiv1.EdgeFunctionsApiV1Options{
				URL: "{BAD_URL_STRING",
				Crn: core.StringPtr(crn),
				ZoneIdentifier: core.StringPtr(zoneIdentifier),
			})

			It(`Instantiate service client with error`, func() {
				Expect(testService).To(BeNil())
				Expect(testServiceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`CreateEdgeFunctionsTrigger(createEdgeFunctionsTriggerOptions *CreateEdgeFunctionsTriggerOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		createEdgeFunctionsTriggerPath := "/v1/testString/zones/testString/workers/routes"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(createEdgeFunctionsTriggerPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateEdgeFunctionsTrigger with error: Operation response processing error`, func() {
				testService, testServiceErr := edgefunctionsapiv1.NewEdgeFunctionsApiV1(&edgefunctionsapiv1.EdgeFunctionsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the CreateEdgeFunctionsTriggerOptions model
				createEdgeFunctionsTriggerOptionsModel := new(edgefunctionsapiv1.CreateEdgeFunctionsTriggerOptions)
				createEdgeFunctionsTriggerOptionsModel.Pattern = core.StringPtr("example.net/*")
				createEdgeFunctionsTriggerOptionsModel.Script = core.StringPtr("this-is_my_script-01")
				createEdgeFunctionsTriggerOptionsModel.XCorrelationID = core.StringPtr("testString")
				createEdgeFunctionsTriggerOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.CreateEdgeFunctionsTrigger(createEdgeFunctionsTriggerOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`CreateEdgeFunctionsTrigger(createEdgeFunctionsTriggerOptions *CreateEdgeFunctionsTriggerOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		createEdgeFunctionsTriggerPath := "/v1/testString/zones/testString/workers/routes"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(createEdgeFunctionsTriggerPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"result": {"id": "9a7806061c88ada191ed06f989cc3dac"}, "success": true, "errors": ["Errors"], "messages": ["Messages"]}`)
				}))
			})
			It(`Invoke CreateEdgeFunctionsTrigger successfully`, func() {
				testService, testServiceErr := edgefunctionsapiv1.NewEdgeFunctionsApiV1(&edgefunctionsapiv1.EdgeFunctionsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.CreateEdgeFunctionsTrigger(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateEdgeFunctionsTriggerOptions model
				createEdgeFunctionsTriggerOptionsModel := new(edgefunctionsapiv1.CreateEdgeFunctionsTriggerOptions)
				createEdgeFunctionsTriggerOptionsModel.Pattern = core.StringPtr("example.net/*")
				createEdgeFunctionsTriggerOptionsModel.Script = core.StringPtr("this-is_my_script-01")
				createEdgeFunctionsTriggerOptionsModel.XCorrelationID = core.StringPtr("testString")
 				createEdgeFunctionsTriggerOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.CreateEdgeFunctionsTrigger(createEdgeFunctionsTriggerOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke CreateEdgeFunctionsTrigger with error: Operation request error`, func() {
				testService, testServiceErr := edgefunctionsapiv1.NewEdgeFunctionsApiV1(&edgefunctionsapiv1.EdgeFunctionsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the CreateEdgeFunctionsTriggerOptions model
				createEdgeFunctionsTriggerOptionsModel := new(edgefunctionsapiv1.CreateEdgeFunctionsTriggerOptions)
				createEdgeFunctionsTriggerOptionsModel.Pattern = core.StringPtr("example.net/*")
				createEdgeFunctionsTriggerOptionsModel.Script = core.StringPtr("this-is_my_script-01")
				createEdgeFunctionsTriggerOptionsModel.XCorrelationID = core.StringPtr("testString")
				createEdgeFunctionsTriggerOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.CreateEdgeFunctionsTrigger(createEdgeFunctionsTriggerOptionsModel)
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
	Describe(`ListEdgeFunctionsTriggers(listEdgeFunctionsTriggersOptions *ListEdgeFunctionsTriggersOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		listEdgeFunctionsTriggersPath := "/v1/testString/zones/testString/workers/routes"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(listEdgeFunctionsTriggersPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListEdgeFunctionsTriggers with error: Operation response processing error`, func() {
				testService, testServiceErr := edgefunctionsapiv1.NewEdgeFunctionsApiV1(&edgefunctionsapiv1.EdgeFunctionsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ListEdgeFunctionsTriggersOptions model
				listEdgeFunctionsTriggersOptionsModel := new(edgefunctionsapiv1.ListEdgeFunctionsTriggersOptions)
				listEdgeFunctionsTriggersOptionsModel.XCorrelationID = core.StringPtr("testString")
				listEdgeFunctionsTriggersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.ListEdgeFunctionsTriggers(listEdgeFunctionsTriggersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ListEdgeFunctionsTriggers(listEdgeFunctionsTriggersOptions *ListEdgeFunctionsTriggersOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		listEdgeFunctionsTriggersPath := "/v1/testString/zones/testString/workers/routes"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(listEdgeFunctionsTriggersPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"result": [{"id": "9a7806061c88ada191ed06f989cc3dac", "pattern": "example.net/*", "script": "this-is_my_script-01", "request_limit_fail_open": false}], "success": true, "errors": ["Errors"], "messages": ["Messages"]}`)
				}))
			})
			It(`Invoke ListEdgeFunctionsTriggers successfully`, func() {
				testService, testServiceErr := edgefunctionsapiv1.NewEdgeFunctionsApiV1(&edgefunctionsapiv1.EdgeFunctionsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.ListEdgeFunctionsTriggers(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListEdgeFunctionsTriggersOptions model
				listEdgeFunctionsTriggersOptionsModel := new(edgefunctionsapiv1.ListEdgeFunctionsTriggersOptions)
				listEdgeFunctionsTriggersOptionsModel.XCorrelationID = core.StringPtr("testString")
 				listEdgeFunctionsTriggersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.ListEdgeFunctionsTriggers(listEdgeFunctionsTriggersOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke ListEdgeFunctionsTriggers with error: Operation request error`, func() {
				testService, testServiceErr := edgefunctionsapiv1.NewEdgeFunctionsApiV1(&edgefunctionsapiv1.EdgeFunctionsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ListEdgeFunctionsTriggersOptions model
				listEdgeFunctionsTriggersOptionsModel := new(edgefunctionsapiv1.ListEdgeFunctionsTriggersOptions)
				listEdgeFunctionsTriggersOptionsModel.XCorrelationID = core.StringPtr("testString")
				listEdgeFunctionsTriggersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.ListEdgeFunctionsTriggers(listEdgeFunctionsTriggersOptionsModel)
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
	Describe(`GetEdgeFunctionsTrigger(getEdgeFunctionsTriggerOptions *GetEdgeFunctionsTriggerOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getEdgeFunctionsTriggerPath := "/v1/testString/zones/testString/workers/routes/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getEdgeFunctionsTriggerPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetEdgeFunctionsTrigger with error: Operation response processing error`, func() {
				testService, testServiceErr := edgefunctionsapiv1.NewEdgeFunctionsApiV1(&edgefunctionsapiv1.EdgeFunctionsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetEdgeFunctionsTriggerOptions model
				getEdgeFunctionsTriggerOptionsModel := new(edgefunctionsapiv1.GetEdgeFunctionsTriggerOptions)
				getEdgeFunctionsTriggerOptionsModel.RouteID = core.StringPtr("testString")
				getEdgeFunctionsTriggerOptionsModel.XCorrelationID = core.StringPtr("testString")
				getEdgeFunctionsTriggerOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.GetEdgeFunctionsTrigger(getEdgeFunctionsTriggerOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetEdgeFunctionsTrigger(getEdgeFunctionsTriggerOptions *GetEdgeFunctionsTriggerOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getEdgeFunctionsTriggerPath := "/v1/testString/zones/testString/workers/routes/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getEdgeFunctionsTriggerPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"result": {"id": "9a7806061c88ada191ed06f989cc3dac", "pattern": "example.net/*", "script": "this-is_my_script-01", "request_limit_fail_open": false}, "success": true, "errors": ["Errors"], "messages": ["Messages"]}`)
				}))
			})
			It(`Invoke GetEdgeFunctionsTrigger successfully`, func() {
				testService, testServiceErr := edgefunctionsapiv1.NewEdgeFunctionsApiV1(&edgefunctionsapiv1.EdgeFunctionsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetEdgeFunctionsTrigger(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetEdgeFunctionsTriggerOptions model
				getEdgeFunctionsTriggerOptionsModel := new(edgefunctionsapiv1.GetEdgeFunctionsTriggerOptions)
				getEdgeFunctionsTriggerOptionsModel.RouteID = core.StringPtr("testString")
				getEdgeFunctionsTriggerOptionsModel.XCorrelationID = core.StringPtr("testString")
 				getEdgeFunctionsTriggerOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetEdgeFunctionsTrigger(getEdgeFunctionsTriggerOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke GetEdgeFunctionsTrigger with error: Operation validation and request error`, func() {
				testService, testServiceErr := edgefunctionsapiv1.NewEdgeFunctionsApiV1(&edgefunctionsapiv1.EdgeFunctionsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetEdgeFunctionsTriggerOptions model
				getEdgeFunctionsTriggerOptionsModel := new(edgefunctionsapiv1.GetEdgeFunctionsTriggerOptions)
				getEdgeFunctionsTriggerOptionsModel.RouteID = core.StringPtr("testString")
				getEdgeFunctionsTriggerOptionsModel.XCorrelationID = core.StringPtr("testString")
				getEdgeFunctionsTriggerOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.GetEdgeFunctionsTrigger(getEdgeFunctionsTriggerOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetEdgeFunctionsTriggerOptions model with no property values
				getEdgeFunctionsTriggerOptionsModelNew := new(edgefunctionsapiv1.GetEdgeFunctionsTriggerOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.GetEdgeFunctionsTrigger(getEdgeFunctionsTriggerOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateEdgeFunctionsTrigger(updateEdgeFunctionsTriggerOptions *UpdateEdgeFunctionsTriggerOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		updateEdgeFunctionsTriggerPath := "/v1/testString/zones/testString/workers/routes/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(updateEdgeFunctionsTriggerPath))
					Expect(req.Method).To(Equal("PUT"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateEdgeFunctionsTrigger with error: Operation response processing error`, func() {
				testService, testServiceErr := edgefunctionsapiv1.NewEdgeFunctionsApiV1(&edgefunctionsapiv1.EdgeFunctionsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the UpdateEdgeFunctionsTriggerOptions model
				updateEdgeFunctionsTriggerOptionsModel := new(edgefunctionsapiv1.UpdateEdgeFunctionsTriggerOptions)
				updateEdgeFunctionsTriggerOptionsModel.RouteID = core.StringPtr("testString")
				updateEdgeFunctionsTriggerOptionsModel.Pattern = core.StringPtr("example.net/*")
				updateEdgeFunctionsTriggerOptionsModel.Script = core.StringPtr("this-is_my_script-01")
				updateEdgeFunctionsTriggerOptionsModel.XCorrelationID = core.StringPtr("testString")
				updateEdgeFunctionsTriggerOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.UpdateEdgeFunctionsTrigger(updateEdgeFunctionsTriggerOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UpdateEdgeFunctionsTrigger(updateEdgeFunctionsTriggerOptions *UpdateEdgeFunctionsTriggerOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		updateEdgeFunctionsTriggerPath := "/v1/testString/zones/testString/workers/routes/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(updateEdgeFunctionsTriggerPath))
					Expect(req.Method).To(Equal("PUT"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"result": {"id": "9a7806061c88ada191ed06f989cc3dac", "pattern": "example.net/*", "script": "this-is_my_script-01", "request_limit_fail_open": false}, "success": true, "errors": ["Errors"], "messages": ["Messages"]}`)
				}))
			})
			It(`Invoke UpdateEdgeFunctionsTrigger successfully`, func() {
				testService, testServiceErr := edgefunctionsapiv1.NewEdgeFunctionsApiV1(&edgefunctionsapiv1.EdgeFunctionsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.UpdateEdgeFunctionsTrigger(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateEdgeFunctionsTriggerOptions model
				updateEdgeFunctionsTriggerOptionsModel := new(edgefunctionsapiv1.UpdateEdgeFunctionsTriggerOptions)
				updateEdgeFunctionsTriggerOptionsModel.RouteID = core.StringPtr("testString")
				updateEdgeFunctionsTriggerOptionsModel.Pattern = core.StringPtr("example.net/*")
				updateEdgeFunctionsTriggerOptionsModel.Script = core.StringPtr("this-is_my_script-01")
				updateEdgeFunctionsTriggerOptionsModel.XCorrelationID = core.StringPtr("testString")
 				updateEdgeFunctionsTriggerOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.UpdateEdgeFunctionsTrigger(updateEdgeFunctionsTriggerOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke UpdateEdgeFunctionsTrigger with error: Operation validation and request error`, func() {
				testService, testServiceErr := edgefunctionsapiv1.NewEdgeFunctionsApiV1(&edgefunctionsapiv1.EdgeFunctionsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the UpdateEdgeFunctionsTriggerOptions model
				updateEdgeFunctionsTriggerOptionsModel := new(edgefunctionsapiv1.UpdateEdgeFunctionsTriggerOptions)
				updateEdgeFunctionsTriggerOptionsModel.RouteID = core.StringPtr("testString")
				updateEdgeFunctionsTriggerOptionsModel.Pattern = core.StringPtr("example.net/*")
				updateEdgeFunctionsTriggerOptionsModel.Script = core.StringPtr("this-is_my_script-01")
				updateEdgeFunctionsTriggerOptionsModel.XCorrelationID = core.StringPtr("testString")
				updateEdgeFunctionsTriggerOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.UpdateEdgeFunctionsTrigger(updateEdgeFunctionsTriggerOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateEdgeFunctionsTriggerOptions model with no property values
				updateEdgeFunctionsTriggerOptionsModelNew := new(edgefunctionsapiv1.UpdateEdgeFunctionsTriggerOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.UpdateEdgeFunctionsTrigger(updateEdgeFunctionsTriggerOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteEdgeFunctionsTrigger(deleteEdgeFunctionsTriggerOptions *DeleteEdgeFunctionsTriggerOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		deleteEdgeFunctionsTriggerPath := "/v1/testString/zones/testString/workers/routes/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(deleteEdgeFunctionsTriggerPath))
					Expect(req.Method).To(Equal("DELETE"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke DeleteEdgeFunctionsTrigger with error: Operation response processing error`, func() {
				testService, testServiceErr := edgefunctionsapiv1.NewEdgeFunctionsApiV1(&edgefunctionsapiv1.EdgeFunctionsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the DeleteEdgeFunctionsTriggerOptions model
				deleteEdgeFunctionsTriggerOptionsModel := new(edgefunctionsapiv1.DeleteEdgeFunctionsTriggerOptions)
				deleteEdgeFunctionsTriggerOptionsModel.RouteID = core.StringPtr("testString")
				deleteEdgeFunctionsTriggerOptionsModel.XCorrelationID = core.StringPtr("testString")
				deleteEdgeFunctionsTriggerOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.DeleteEdgeFunctionsTrigger(deleteEdgeFunctionsTriggerOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`DeleteEdgeFunctionsTrigger(deleteEdgeFunctionsTriggerOptions *DeleteEdgeFunctionsTriggerOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		deleteEdgeFunctionsTriggerPath := "/v1/testString/zones/testString/workers/routes/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(deleteEdgeFunctionsTriggerPath))
					Expect(req.Method).To(Equal("DELETE"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"result": {"id": "9a7806061c88ada191ed06f989cc3dac"}, "success": true, "errors": ["Errors"], "messages": ["Messages"]}`)
				}))
			})
			It(`Invoke DeleteEdgeFunctionsTrigger successfully`, func() {
				testService, testServiceErr := edgefunctionsapiv1.NewEdgeFunctionsApiV1(&edgefunctionsapiv1.EdgeFunctionsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.DeleteEdgeFunctionsTrigger(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DeleteEdgeFunctionsTriggerOptions model
				deleteEdgeFunctionsTriggerOptionsModel := new(edgefunctionsapiv1.DeleteEdgeFunctionsTriggerOptions)
				deleteEdgeFunctionsTriggerOptionsModel.RouteID = core.StringPtr("testString")
				deleteEdgeFunctionsTriggerOptionsModel.XCorrelationID = core.StringPtr("testString")
 				deleteEdgeFunctionsTriggerOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.DeleteEdgeFunctionsTrigger(deleteEdgeFunctionsTriggerOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke DeleteEdgeFunctionsTrigger with error: Operation validation and request error`, func() {
				testService, testServiceErr := edgefunctionsapiv1.NewEdgeFunctionsApiV1(&edgefunctionsapiv1.EdgeFunctionsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the DeleteEdgeFunctionsTriggerOptions model
				deleteEdgeFunctionsTriggerOptionsModel := new(edgefunctionsapiv1.DeleteEdgeFunctionsTriggerOptions)
				deleteEdgeFunctionsTriggerOptionsModel.RouteID = core.StringPtr("testString")
				deleteEdgeFunctionsTriggerOptionsModel.XCorrelationID = core.StringPtr("testString")
				deleteEdgeFunctionsTriggerOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.DeleteEdgeFunctionsTrigger(deleteEdgeFunctionsTriggerOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the DeleteEdgeFunctionsTriggerOptions model with no property values
				deleteEdgeFunctionsTriggerOptionsModelNew := new(edgefunctionsapiv1.DeleteEdgeFunctionsTriggerOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.DeleteEdgeFunctionsTrigger(deleteEdgeFunctionsTriggerOptionsModelNew)
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
			testService, _ := edgefunctionsapiv1.NewEdgeFunctionsApiV1(&edgefunctionsapiv1.EdgeFunctionsApiV1Options{
				URL:           "http://edgefunctionsapiv1modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
				Crn: core.StringPtr(crn),
				ZoneIdentifier: core.StringPtr(zoneIdentifier),
			})
			It(`Invoke NewCreateEdgeFunctionsTriggerOptions successfully`, func() {
				// Construct an instance of the CreateEdgeFunctionsTriggerOptions model
				createEdgeFunctionsTriggerOptionsModel := testService.NewCreateEdgeFunctionsTriggerOptions()
				createEdgeFunctionsTriggerOptionsModel.SetPattern("example.net/*")
				createEdgeFunctionsTriggerOptionsModel.SetScript("this-is_my_script-01")
				createEdgeFunctionsTriggerOptionsModel.SetXCorrelationID("testString")
				createEdgeFunctionsTriggerOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createEdgeFunctionsTriggerOptionsModel).ToNot(BeNil())
				Expect(createEdgeFunctionsTriggerOptionsModel.Pattern).To(Equal(core.StringPtr("example.net/*")))
				Expect(createEdgeFunctionsTriggerOptionsModel.Script).To(Equal(core.StringPtr("this-is_my_script-01")))
				Expect(createEdgeFunctionsTriggerOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(createEdgeFunctionsTriggerOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteEdgeFunctionsActionOptions successfully`, func() {
				// Construct an instance of the DeleteEdgeFunctionsActionOptions model
				scriptName := "testString"
				deleteEdgeFunctionsActionOptionsModel := testService.NewDeleteEdgeFunctionsActionOptions(scriptName)
				deleteEdgeFunctionsActionOptionsModel.SetScriptName("testString")
				deleteEdgeFunctionsActionOptionsModel.SetXCorrelationID("testString")
				deleteEdgeFunctionsActionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteEdgeFunctionsActionOptionsModel).ToNot(BeNil())
				Expect(deleteEdgeFunctionsActionOptionsModel.ScriptName).To(Equal(core.StringPtr("testString")))
				Expect(deleteEdgeFunctionsActionOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(deleteEdgeFunctionsActionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteEdgeFunctionsTriggerOptions successfully`, func() {
				// Construct an instance of the DeleteEdgeFunctionsTriggerOptions model
				routeID := "testString"
				deleteEdgeFunctionsTriggerOptionsModel := testService.NewDeleteEdgeFunctionsTriggerOptions(routeID)
				deleteEdgeFunctionsTriggerOptionsModel.SetRouteID("testString")
				deleteEdgeFunctionsTriggerOptionsModel.SetXCorrelationID("testString")
				deleteEdgeFunctionsTriggerOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteEdgeFunctionsTriggerOptionsModel).ToNot(BeNil())
				Expect(deleteEdgeFunctionsTriggerOptionsModel.RouteID).To(Equal(core.StringPtr("testString")))
				Expect(deleteEdgeFunctionsTriggerOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(deleteEdgeFunctionsTriggerOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetEdgeFunctionsActionOptions successfully`, func() {
				// Construct an instance of the GetEdgeFunctionsActionOptions model
				scriptName := "testString"
				getEdgeFunctionsActionOptionsModel := testService.NewGetEdgeFunctionsActionOptions(scriptName)
				getEdgeFunctionsActionOptionsModel.SetScriptName("testString")
				getEdgeFunctionsActionOptionsModel.SetXCorrelationID("testString")
				getEdgeFunctionsActionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getEdgeFunctionsActionOptionsModel).ToNot(BeNil())
				Expect(getEdgeFunctionsActionOptionsModel.ScriptName).To(Equal(core.StringPtr("testString")))
				Expect(getEdgeFunctionsActionOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(getEdgeFunctionsActionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetEdgeFunctionsTriggerOptions successfully`, func() {
				// Construct an instance of the GetEdgeFunctionsTriggerOptions model
				routeID := "testString"
				getEdgeFunctionsTriggerOptionsModel := testService.NewGetEdgeFunctionsTriggerOptions(routeID)
				getEdgeFunctionsTriggerOptionsModel.SetRouteID("testString")
				getEdgeFunctionsTriggerOptionsModel.SetXCorrelationID("testString")
				getEdgeFunctionsTriggerOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getEdgeFunctionsTriggerOptionsModel).ToNot(BeNil())
				Expect(getEdgeFunctionsTriggerOptionsModel.RouteID).To(Equal(core.StringPtr("testString")))
				Expect(getEdgeFunctionsTriggerOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(getEdgeFunctionsTriggerOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListEdgeFunctionsActionsOptions successfully`, func() {
				// Construct an instance of the ListEdgeFunctionsActionsOptions model
				listEdgeFunctionsActionsOptionsModel := testService.NewListEdgeFunctionsActionsOptions()
				listEdgeFunctionsActionsOptionsModel.SetXCorrelationID("testString")
				listEdgeFunctionsActionsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listEdgeFunctionsActionsOptionsModel).ToNot(BeNil())
				Expect(listEdgeFunctionsActionsOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(listEdgeFunctionsActionsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListEdgeFunctionsTriggersOptions successfully`, func() {
				// Construct an instance of the ListEdgeFunctionsTriggersOptions model
				listEdgeFunctionsTriggersOptionsModel := testService.NewListEdgeFunctionsTriggersOptions()
				listEdgeFunctionsTriggersOptionsModel.SetXCorrelationID("testString")
				listEdgeFunctionsTriggersOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listEdgeFunctionsTriggersOptionsModel).ToNot(BeNil())
				Expect(listEdgeFunctionsTriggersOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(listEdgeFunctionsTriggersOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateEdgeFunctionsActionOptions successfully`, func() {
				// Construct an instance of the UpdateEdgeFunctionsActionOptions model
				scriptName := "testString"
				updateEdgeFunctionsActionOptionsModel := testService.NewUpdateEdgeFunctionsActionOptions(scriptName)
				updateEdgeFunctionsActionOptionsModel.SetScriptName("testString")
				updateEdgeFunctionsActionOptionsModel.SetEdgeFunctionsAction(CreateMockReader("This is a mock file."))
				updateEdgeFunctionsActionOptionsModel.SetXCorrelationID("testString")
				updateEdgeFunctionsActionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateEdgeFunctionsActionOptionsModel).ToNot(BeNil())
				Expect(updateEdgeFunctionsActionOptionsModel.ScriptName).To(Equal(core.StringPtr("testString")))
				Expect(updateEdgeFunctionsActionOptionsModel.EdgeFunctionsAction).To(Equal(CreateMockReader("This is a mock file.")))
				Expect(updateEdgeFunctionsActionOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(updateEdgeFunctionsActionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateEdgeFunctionsTriggerOptions successfully`, func() {
				// Construct an instance of the UpdateEdgeFunctionsTriggerOptions model
				routeID := "testString"
				updateEdgeFunctionsTriggerOptionsModel := testService.NewUpdateEdgeFunctionsTriggerOptions(routeID)
				updateEdgeFunctionsTriggerOptionsModel.SetRouteID("testString")
				updateEdgeFunctionsTriggerOptionsModel.SetPattern("example.net/*")
				updateEdgeFunctionsTriggerOptionsModel.SetScript("this-is_my_script-01")
				updateEdgeFunctionsTriggerOptionsModel.SetXCorrelationID("testString")
				updateEdgeFunctionsTriggerOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateEdgeFunctionsTriggerOptionsModel).ToNot(BeNil())
				Expect(updateEdgeFunctionsTriggerOptionsModel.RouteID).To(Equal(core.StringPtr("testString")))
				Expect(updateEdgeFunctionsTriggerOptionsModel.Pattern).To(Equal(core.StringPtr("example.net/*")))
				Expect(updateEdgeFunctionsTriggerOptionsModel.Script).To(Equal(core.StringPtr("this-is_my_script-01")))
				Expect(updateEdgeFunctionsTriggerOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(updateEdgeFunctionsTriggerOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
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
