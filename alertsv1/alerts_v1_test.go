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

package alertsv1_test

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
	"github.com/IBM/networking-go-sdk/alertsv1"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe(`AlertsV1`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		crn := "testString"
		It(`Instantiate service client`, func() {
			alertsService, serviceErr := alertsv1.NewAlertsV1(&alertsv1.AlertsV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
				Crn: core.StringPtr(crn),
			})
			Expect(alertsService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			alertsService, serviceErr := alertsv1.NewAlertsV1(&alertsv1.AlertsV1Options{
				URL: "{BAD_URL_STRING",
				Crn: core.StringPtr(crn),
			})
			Expect(alertsService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			alertsService, serviceErr := alertsv1.NewAlertsV1(&alertsv1.AlertsV1Options{
				URL: "https://alertsv1/api",
				Crn: core.StringPtr(crn),
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(alertsService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Validation Error`, func() {
			alertsService, serviceErr := alertsv1.NewAlertsV1(&alertsv1.AlertsV1Options{})
			Expect(alertsService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		crn := "testString"
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"ALERTS_URL": "https://alertsv1/api",
				"ALERTS_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				alertsService, serviceErr := alertsv1.NewAlertsV1UsingExternalConfig(&alertsv1.AlertsV1Options{
					Crn: core.StringPtr(crn),
				})
				Expect(alertsService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := alertsService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != alertsService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(alertsService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(alertsService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				alertsService, serviceErr := alertsv1.NewAlertsV1UsingExternalConfig(&alertsv1.AlertsV1Options{
					URL: "https://testService/api",
					Crn: core.StringPtr(crn),
				})
				Expect(alertsService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(alertsService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := alertsService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != alertsService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(alertsService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(alertsService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				alertsService, serviceErr := alertsv1.NewAlertsV1UsingExternalConfig(&alertsv1.AlertsV1Options{
					Crn: core.StringPtr(crn),
				})
				err := alertsService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(alertsService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(alertsService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := alertsService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != alertsService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(alertsService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(alertsService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"ALERTS_URL": "https://alertsv1/api",
				"ALERTS_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			alertsService, serviceErr := alertsv1.NewAlertsV1UsingExternalConfig(&alertsv1.AlertsV1Options{
				Crn: core.StringPtr(crn),
			})

			It(`Instantiate service client with error`, func() {
				Expect(alertsService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"ALERTS_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			alertsService, serviceErr := alertsv1.NewAlertsV1UsingExternalConfig(&alertsv1.AlertsV1Options{
				URL: "{BAD_URL_STRING",
				Crn: core.StringPtr(crn),
			})

			It(`Instantiate service client with error`, func() {
				Expect(alertsService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = alertsv1.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`GetAlertPolicies(getAlertPoliciesOptions *GetAlertPoliciesOptions) - Operation response error`, func() {
		crn := "testString"
		getAlertPoliciesPath := "/v1/testString/alerting/policies"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getAlertPoliciesPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetAlertPolicies with error: Operation response processing error`, func() {
				alertsService, serviceErr := alertsv1.NewAlertsV1(&alertsv1.AlertsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(alertsService).ToNot(BeNil())

				// Construct an instance of the GetAlertPoliciesOptions model
				getAlertPoliciesOptionsModel := new(alertsv1.GetAlertPoliciesOptions)
				getAlertPoliciesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := alertsService.GetAlertPolicies(getAlertPoliciesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				alertsService.EnableRetries(0, 0)
				result, response, operationErr = alertsService.GetAlertPolicies(getAlertPoliciesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetAlertPolicies(getAlertPoliciesOptions *GetAlertPoliciesOptions)`, func() {
		crn := "testString"
		getAlertPoliciesPath := "/v1/testString/alerting/policies"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getAlertPoliciesPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [{"id": "ID"}], "messages": [{"id": "ID"}], "result": [{"id": "f0413b106d2c4aa9b1553d5d0209c522", "name": "My Alert Policy", "description": "Description for my alert policy", "enabled": true, "alert_type": "dos_attack_l7", "mechanisms": {"email": [{"id": "ID"}], "webhooks": [{"id": "ID"}]}, "created": "2021-09-15T16:33:31.834209Z", "modified": "2021-09-15T16:33:31.834209Z", "conditions": {"anyKey": "anyValue"}, "filters": {"anyKey": "anyValue"}}]}`)
				}))
			})
			It(`Invoke GetAlertPolicies successfully with retries`, func() {
				alertsService, serviceErr := alertsv1.NewAlertsV1(&alertsv1.AlertsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(alertsService).ToNot(BeNil())
				alertsService.EnableRetries(0, 0)

				// Construct an instance of the GetAlertPoliciesOptions model
				getAlertPoliciesOptionsModel := new(alertsv1.GetAlertPoliciesOptions)
				getAlertPoliciesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := alertsService.GetAlertPoliciesWithContext(ctx, getAlertPoliciesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				alertsService.DisableRetries()
				result, response, operationErr := alertsService.GetAlertPolicies(getAlertPoliciesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = alertsService.GetAlertPoliciesWithContext(ctx, getAlertPoliciesOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getAlertPoliciesPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [{"id": "ID"}], "messages": [{"id": "ID"}], "result": [{"id": "f0413b106d2c4aa9b1553d5d0209c522", "name": "My Alert Policy", "description": "Description for my alert policy", "enabled": true, "alert_type": "dos_attack_l7", "mechanisms": {"email": [{"id": "ID"}], "webhooks": [{"id": "ID"}]}, "created": "2021-09-15T16:33:31.834209Z", "modified": "2021-09-15T16:33:31.834209Z", "conditions": {"anyKey": "anyValue"}, "filters": {"anyKey": "anyValue"}}]}`)
				}))
			})
			It(`Invoke GetAlertPolicies successfully`, func() {
				alertsService, serviceErr := alertsv1.NewAlertsV1(&alertsv1.AlertsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(alertsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := alertsService.GetAlertPolicies(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetAlertPoliciesOptions model
				getAlertPoliciesOptionsModel := new(alertsv1.GetAlertPoliciesOptions)
				getAlertPoliciesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = alertsService.GetAlertPolicies(getAlertPoliciesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetAlertPolicies with error: Operation request error`, func() {
				alertsService, serviceErr := alertsv1.NewAlertsV1(&alertsv1.AlertsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(alertsService).ToNot(BeNil())

				// Construct an instance of the GetAlertPoliciesOptions model
				getAlertPoliciesOptionsModel := new(alertsv1.GetAlertPoliciesOptions)
				getAlertPoliciesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := alertsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := alertsService.GetAlertPolicies(getAlertPoliciesOptionsModel)
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
			It(`Invoke GetAlertPolicies successfully`, func() {
				alertsService, serviceErr := alertsv1.NewAlertsV1(&alertsv1.AlertsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(alertsService).ToNot(BeNil())

				// Construct an instance of the GetAlertPoliciesOptions model
				getAlertPoliciesOptionsModel := new(alertsv1.GetAlertPoliciesOptions)
				getAlertPoliciesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := alertsService.GetAlertPolicies(getAlertPoliciesOptionsModel)
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
	Describe(`CreateAlertPolicy(createAlertPolicyOptions *CreateAlertPolicyOptions) - Operation response error`, func() {
		crn := "testString"
		createAlertPolicyPath := "/v1/testString/alerting/policies"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createAlertPolicyPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateAlertPolicy with error: Operation response processing error`, func() {
				alertsService, serviceErr := alertsv1.NewAlertsV1(&alertsv1.AlertsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(alertsService).ToNot(BeNil())

				// Construct an instance of the CreateAlertPolicyInputMechanismsEmailItem model
				createAlertPolicyInputMechanismsEmailItemModel := new(alertsv1.CreateAlertPolicyInputMechanismsEmailItem)
				createAlertPolicyInputMechanismsEmailItemModel.ID = core.StringPtr("mynotifications@email.com")

				// Construct an instance of the CreateAlertPolicyInputMechanismsWebhooksItem model
				createAlertPolicyInputMechanismsWebhooksItemModel := new(alertsv1.CreateAlertPolicyInputMechanismsWebhooksItem)
				createAlertPolicyInputMechanismsWebhooksItemModel.ID = core.StringPtr("f0413b106d2c4aa9b1553d5d0209c522")

				// Construct an instance of the CreateAlertPolicyInputMechanisms model
				createAlertPolicyInputMechanismsModel := new(alertsv1.CreateAlertPolicyInputMechanisms)
				createAlertPolicyInputMechanismsModel.Email = []alertsv1.CreateAlertPolicyInputMechanismsEmailItem{*createAlertPolicyInputMechanismsEmailItemModel}
				createAlertPolicyInputMechanismsModel.Webhooks = []alertsv1.CreateAlertPolicyInputMechanismsWebhooksItem{*createAlertPolicyInputMechanismsWebhooksItemModel}

				// Construct an instance of the CreateAlertPolicyOptions model
				createAlertPolicyOptionsModel := new(alertsv1.CreateAlertPolicyOptions)
				createAlertPolicyOptionsModel.Name = core.StringPtr("My Alert Policy")
				createAlertPolicyOptionsModel.Description = core.StringPtr("A description for my alert policy")
				createAlertPolicyOptionsModel.Enabled = core.BoolPtr(true)
				createAlertPolicyOptionsModel.AlertType = core.StringPtr("dos_attack_l7")
				createAlertPolicyOptionsModel.Mechanisms = createAlertPolicyInputMechanismsModel
				createAlertPolicyOptionsModel.Filters = map[string]interface{}{"anyKey": "anyValue"}
				createAlertPolicyOptionsModel.Conditions = map[string]interface{}{"anyKey": "anyValue"}
				createAlertPolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := alertsService.CreateAlertPolicy(createAlertPolicyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				alertsService.EnableRetries(0, 0)
				result, response, operationErr = alertsService.CreateAlertPolicy(createAlertPolicyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateAlertPolicy(createAlertPolicyOptions *CreateAlertPolicyOptions)`, func() {
		crn := "testString"
		createAlertPolicyPath := "/v1/testString/alerting/policies"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createAlertPolicyPath))
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
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [{"id": "ID"}], "messages": [{"id": "ID"}], "result": {"id": "f0413b106d2c4aa9b1553d5d0209c522"}}`)
				}))
			})
			It(`Invoke CreateAlertPolicy successfully with retries`, func() {
				alertsService, serviceErr := alertsv1.NewAlertsV1(&alertsv1.AlertsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(alertsService).ToNot(BeNil())
				alertsService.EnableRetries(0, 0)

				// Construct an instance of the CreateAlertPolicyInputMechanismsEmailItem model
				createAlertPolicyInputMechanismsEmailItemModel := new(alertsv1.CreateAlertPolicyInputMechanismsEmailItem)
				createAlertPolicyInputMechanismsEmailItemModel.ID = core.StringPtr("mynotifications@email.com")

				// Construct an instance of the CreateAlertPolicyInputMechanismsWebhooksItem model
				createAlertPolicyInputMechanismsWebhooksItemModel := new(alertsv1.CreateAlertPolicyInputMechanismsWebhooksItem)
				createAlertPolicyInputMechanismsWebhooksItemModel.ID = core.StringPtr("f0413b106d2c4aa9b1553d5d0209c522")

				// Construct an instance of the CreateAlertPolicyInputMechanisms model
				createAlertPolicyInputMechanismsModel := new(alertsv1.CreateAlertPolicyInputMechanisms)
				createAlertPolicyInputMechanismsModel.Email = []alertsv1.CreateAlertPolicyInputMechanismsEmailItem{*createAlertPolicyInputMechanismsEmailItemModel}
				createAlertPolicyInputMechanismsModel.Webhooks = []alertsv1.CreateAlertPolicyInputMechanismsWebhooksItem{*createAlertPolicyInputMechanismsWebhooksItemModel}

				// Construct an instance of the CreateAlertPolicyOptions model
				createAlertPolicyOptionsModel := new(alertsv1.CreateAlertPolicyOptions)
				createAlertPolicyOptionsModel.Name = core.StringPtr("My Alert Policy")
				createAlertPolicyOptionsModel.Description = core.StringPtr("A description for my alert policy")
				createAlertPolicyOptionsModel.Enabled = core.BoolPtr(true)
				createAlertPolicyOptionsModel.AlertType = core.StringPtr("dos_attack_l7")
				createAlertPolicyOptionsModel.Mechanisms = createAlertPolicyInputMechanismsModel
				createAlertPolicyOptionsModel.Filters = map[string]interface{}{"anyKey": "anyValue"}
				createAlertPolicyOptionsModel.Conditions = map[string]interface{}{"anyKey": "anyValue"}
				createAlertPolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := alertsService.CreateAlertPolicyWithContext(ctx, createAlertPolicyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				alertsService.DisableRetries()
				result, response, operationErr := alertsService.CreateAlertPolicy(createAlertPolicyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = alertsService.CreateAlertPolicyWithContext(ctx, createAlertPolicyOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createAlertPolicyPath))
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
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [{"id": "ID"}], "messages": [{"id": "ID"}], "result": {"id": "f0413b106d2c4aa9b1553d5d0209c522"}}`)
				}))
			})
			It(`Invoke CreateAlertPolicy successfully`, func() {
				alertsService, serviceErr := alertsv1.NewAlertsV1(&alertsv1.AlertsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(alertsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := alertsService.CreateAlertPolicy(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateAlertPolicyInputMechanismsEmailItem model
				createAlertPolicyInputMechanismsEmailItemModel := new(alertsv1.CreateAlertPolicyInputMechanismsEmailItem)
				createAlertPolicyInputMechanismsEmailItemModel.ID = core.StringPtr("mynotifications@email.com")

				// Construct an instance of the CreateAlertPolicyInputMechanismsWebhooksItem model
				createAlertPolicyInputMechanismsWebhooksItemModel := new(alertsv1.CreateAlertPolicyInputMechanismsWebhooksItem)
				createAlertPolicyInputMechanismsWebhooksItemModel.ID = core.StringPtr("f0413b106d2c4aa9b1553d5d0209c522")

				// Construct an instance of the CreateAlertPolicyInputMechanisms model
				createAlertPolicyInputMechanismsModel := new(alertsv1.CreateAlertPolicyInputMechanisms)
				createAlertPolicyInputMechanismsModel.Email = []alertsv1.CreateAlertPolicyInputMechanismsEmailItem{*createAlertPolicyInputMechanismsEmailItemModel}
				createAlertPolicyInputMechanismsModel.Webhooks = []alertsv1.CreateAlertPolicyInputMechanismsWebhooksItem{*createAlertPolicyInputMechanismsWebhooksItemModel}

				// Construct an instance of the CreateAlertPolicyOptions model
				createAlertPolicyOptionsModel := new(alertsv1.CreateAlertPolicyOptions)
				createAlertPolicyOptionsModel.Name = core.StringPtr("My Alert Policy")
				createAlertPolicyOptionsModel.Description = core.StringPtr("A description for my alert policy")
				createAlertPolicyOptionsModel.Enabled = core.BoolPtr(true)
				createAlertPolicyOptionsModel.AlertType = core.StringPtr("dos_attack_l7")
				createAlertPolicyOptionsModel.Mechanisms = createAlertPolicyInputMechanismsModel
				createAlertPolicyOptionsModel.Filters = map[string]interface{}{"anyKey": "anyValue"}
				createAlertPolicyOptionsModel.Conditions = map[string]interface{}{"anyKey": "anyValue"}
				createAlertPolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = alertsService.CreateAlertPolicy(createAlertPolicyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateAlertPolicy with error: Operation request error`, func() {
				alertsService, serviceErr := alertsv1.NewAlertsV1(&alertsv1.AlertsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(alertsService).ToNot(BeNil())

				// Construct an instance of the CreateAlertPolicyInputMechanismsEmailItem model
				createAlertPolicyInputMechanismsEmailItemModel := new(alertsv1.CreateAlertPolicyInputMechanismsEmailItem)
				createAlertPolicyInputMechanismsEmailItemModel.ID = core.StringPtr("mynotifications@email.com")

				// Construct an instance of the CreateAlertPolicyInputMechanismsWebhooksItem model
				createAlertPolicyInputMechanismsWebhooksItemModel := new(alertsv1.CreateAlertPolicyInputMechanismsWebhooksItem)
				createAlertPolicyInputMechanismsWebhooksItemModel.ID = core.StringPtr("f0413b106d2c4aa9b1553d5d0209c522")

				// Construct an instance of the CreateAlertPolicyInputMechanisms model
				createAlertPolicyInputMechanismsModel := new(alertsv1.CreateAlertPolicyInputMechanisms)
				createAlertPolicyInputMechanismsModel.Email = []alertsv1.CreateAlertPolicyInputMechanismsEmailItem{*createAlertPolicyInputMechanismsEmailItemModel}
				createAlertPolicyInputMechanismsModel.Webhooks = []alertsv1.CreateAlertPolicyInputMechanismsWebhooksItem{*createAlertPolicyInputMechanismsWebhooksItemModel}

				// Construct an instance of the CreateAlertPolicyOptions model
				createAlertPolicyOptionsModel := new(alertsv1.CreateAlertPolicyOptions)
				createAlertPolicyOptionsModel.Name = core.StringPtr("My Alert Policy")
				createAlertPolicyOptionsModel.Description = core.StringPtr("A description for my alert policy")
				createAlertPolicyOptionsModel.Enabled = core.BoolPtr(true)
				createAlertPolicyOptionsModel.AlertType = core.StringPtr("dos_attack_l7")
				createAlertPolicyOptionsModel.Mechanisms = createAlertPolicyInputMechanismsModel
				createAlertPolicyOptionsModel.Filters = map[string]interface{}{"anyKey": "anyValue"}
				createAlertPolicyOptionsModel.Conditions = map[string]interface{}{"anyKey": "anyValue"}
				createAlertPolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := alertsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := alertsService.CreateAlertPolicy(createAlertPolicyOptionsModel)
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
			It(`Invoke CreateAlertPolicy successfully`, func() {
				alertsService, serviceErr := alertsv1.NewAlertsV1(&alertsv1.AlertsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(alertsService).ToNot(BeNil())

				// Construct an instance of the CreateAlertPolicyInputMechanismsEmailItem model
				createAlertPolicyInputMechanismsEmailItemModel := new(alertsv1.CreateAlertPolicyInputMechanismsEmailItem)
				createAlertPolicyInputMechanismsEmailItemModel.ID = core.StringPtr("mynotifications@email.com")

				// Construct an instance of the CreateAlertPolicyInputMechanismsWebhooksItem model
				createAlertPolicyInputMechanismsWebhooksItemModel := new(alertsv1.CreateAlertPolicyInputMechanismsWebhooksItem)
				createAlertPolicyInputMechanismsWebhooksItemModel.ID = core.StringPtr("f0413b106d2c4aa9b1553d5d0209c522")

				// Construct an instance of the CreateAlertPolicyInputMechanisms model
				createAlertPolicyInputMechanismsModel := new(alertsv1.CreateAlertPolicyInputMechanisms)
				createAlertPolicyInputMechanismsModel.Email = []alertsv1.CreateAlertPolicyInputMechanismsEmailItem{*createAlertPolicyInputMechanismsEmailItemModel}
				createAlertPolicyInputMechanismsModel.Webhooks = []alertsv1.CreateAlertPolicyInputMechanismsWebhooksItem{*createAlertPolicyInputMechanismsWebhooksItemModel}

				// Construct an instance of the CreateAlertPolicyOptions model
				createAlertPolicyOptionsModel := new(alertsv1.CreateAlertPolicyOptions)
				createAlertPolicyOptionsModel.Name = core.StringPtr("My Alert Policy")
				createAlertPolicyOptionsModel.Description = core.StringPtr("A description for my alert policy")
				createAlertPolicyOptionsModel.Enabled = core.BoolPtr(true)
				createAlertPolicyOptionsModel.AlertType = core.StringPtr("dos_attack_l7")
				createAlertPolicyOptionsModel.Mechanisms = createAlertPolicyInputMechanismsModel
				createAlertPolicyOptionsModel.Filters = map[string]interface{}{"anyKey": "anyValue"}
				createAlertPolicyOptionsModel.Conditions = map[string]interface{}{"anyKey": "anyValue"}
				createAlertPolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := alertsService.CreateAlertPolicy(createAlertPolicyOptionsModel)
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
	Describe(`GetAlertPolicy(getAlertPolicyOptions *GetAlertPolicyOptions) - Operation response error`, func() {
		crn := "testString"
		getAlertPolicyPath := "/v1/testString/alerting/policies/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getAlertPolicyPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetAlertPolicy with error: Operation response processing error`, func() {
				alertsService, serviceErr := alertsv1.NewAlertsV1(&alertsv1.AlertsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(alertsService).ToNot(BeNil())

				// Construct an instance of the GetAlertPolicyOptions model
				getAlertPolicyOptionsModel := new(alertsv1.GetAlertPolicyOptions)
				getAlertPolicyOptionsModel.PolicyID = core.StringPtr("testString")
				getAlertPolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := alertsService.GetAlertPolicy(getAlertPolicyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				alertsService.EnableRetries(0, 0)
				result, response, operationErr = alertsService.GetAlertPolicy(getAlertPolicyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetAlertPolicy(getAlertPolicyOptions *GetAlertPolicyOptions)`, func() {
		crn := "testString"
		getAlertPolicyPath := "/v1/testString/alerting/policies/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getAlertPolicyPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [{"id": "ID"}], "messages": [{"id": "ID"}], "result": {"id": "f0413b106d2c4aa9b1553d5d0209c522", "name": "My Alert Policy", "description": "Description for my alert policy", "enabled": true, "alert_type": "dos_attack_l7", "mechanisms": {"email": [{"id": "ID"}], "webhooks": [{"id": "ID"}]}, "created": "2021-09-15T16:33:31.834209Z", "modified": "2021-09-15T16:33:31.834209Z", "conditions": {"anyKey": "anyValue"}, "filters": {"anyKey": "anyValue"}}}`)
				}))
			})
			It(`Invoke GetAlertPolicy successfully with retries`, func() {
				alertsService, serviceErr := alertsv1.NewAlertsV1(&alertsv1.AlertsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(alertsService).ToNot(BeNil())
				alertsService.EnableRetries(0, 0)

				// Construct an instance of the GetAlertPolicyOptions model
				getAlertPolicyOptionsModel := new(alertsv1.GetAlertPolicyOptions)
				getAlertPolicyOptionsModel.PolicyID = core.StringPtr("testString")
				getAlertPolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := alertsService.GetAlertPolicyWithContext(ctx, getAlertPolicyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				alertsService.DisableRetries()
				result, response, operationErr := alertsService.GetAlertPolicy(getAlertPolicyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = alertsService.GetAlertPolicyWithContext(ctx, getAlertPolicyOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getAlertPolicyPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [{"id": "ID"}], "messages": [{"id": "ID"}], "result": {"id": "f0413b106d2c4aa9b1553d5d0209c522", "name": "My Alert Policy", "description": "Description for my alert policy", "enabled": true, "alert_type": "dos_attack_l7", "mechanisms": {"email": [{"id": "ID"}], "webhooks": [{"id": "ID"}]}, "created": "2021-09-15T16:33:31.834209Z", "modified": "2021-09-15T16:33:31.834209Z", "conditions": {"anyKey": "anyValue"}, "filters": {"anyKey": "anyValue"}}}`)
				}))
			})
			It(`Invoke GetAlertPolicy successfully`, func() {
				alertsService, serviceErr := alertsv1.NewAlertsV1(&alertsv1.AlertsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(alertsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := alertsService.GetAlertPolicy(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetAlertPolicyOptions model
				getAlertPolicyOptionsModel := new(alertsv1.GetAlertPolicyOptions)
				getAlertPolicyOptionsModel.PolicyID = core.StringPtr("testString")
				getAlertPolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = alertsService.GetAlertPolicy(getAlertPolicyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetAlertPolicy with error: Operation validation and request error`, func() {
				alertsService, serviceErr := alertsv1.NewAlertsV1(&alertsv1.AlertsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(alertsService).ToNot(BeNil())

				// Construct an instance of the GetAlertPolicyOptions model
				getAlertPolicyOptionsModel := new(alertsv1.GetAlertPolicyOptions)
				getAlertPolicyOptionsModel.PolicyID = core.StringPtr("testString")
				getAlertPolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := alertsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := alertsService.GetAlertPolicy(getAlertPolicyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetAlertPolicyOptions model with no property values
				getAlertPolicyOptionsModelNew := new(alertsv1.GetAlertPolicyOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = alertsService.GetAlertPolicy(getAlertPolicyOptionsModelNew)
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
			It(`Invoke GetAlertPolicy successfully`, func() {
				alertsService, serviceErr := alertsv1.NewAlertsV1(&alertsv1.AlertsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(alertsService).ToNot(BeNil())

				// Construct an instance of the GetAlertPolicyOptions model
				getAlertPolicyOptionsModel := new(alertsv1.GetAlertPolicyOptions)
				getAlertPolicyOptionsModel.PolicyID = core.StringPtr("testString")
				getAlertPolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := alertsService.GetAlertPolicy(getAlertPolicyOptionsModel)
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
	Describe(`UpdateAlertPolicy(updateAlertPolicyOptions *UpdateAlertPolicyOptions) - Operation response error`, func() {
		crn := "testString"
		updateAlertPolicyPath := "/v1/testString/alerting/policies/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateAlertPolicyPath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateAlertPolicy with error: Operation response processing error`, func() {
				alertsService, serviceErr := alertsv1.NewAlertsV1(&alertsv1.AlertsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(alertsService).ToNot(BeNil())

				// Construct an instance of the UpdateAlertPolicyInputMechanismsEmailItem model
				updateAlertPolicyInputMechanismsEmailItemModel := new(alertsv1.UpdateAlertPolicyInputMechanismsEmailItem)
				updateAlertPolicyInputMechanismsEmailItemModel.ID = core.StringPtr("mynotifications@email.com")

				// Construct an instance of the UpdateAlertPolicyInputMechanismsWebhooksItem model
				updateAlertPolicyInputMechanismsWebhooksItemModel := new(alertsv1.UpdateAlertPolicyInputMechanismsWebhooksItem)
				updateAlertPolicyInputMechanismsWebhooksItemModel.ID = core.StringPtr("f0413b106d2c4aa9b1553d5d0209c522")

				// Construct an instance of the UpdateAlertPolicyInputMechanisms model
				updateAlertPolicyInputMechanismsModel := new(alertsv1.UpdateAlertPolicyInputMechanisms)
				updateAlertPolicyInputMechanismsModel.Email = []alertsv1.UpdateAlertPolicyInputMechanismsEmailItem{*updateAlertPolicyInputMechanismsEmailItemModel}
				updateAlertPolicyInputMechanismsModel.Webhooks = []alertsv1.UpdateAlertPolicyInputMechanismsWebhooksItem{*updateAlertPolicyInputMechanismsWebhooksItemModel}

				// Construct an instance of the UpdateAlertPolicyOptions model
				updateAlertPolicyOptionsModel := new(alertsv1.UpdateAlertPolicyOptions)
				updateAlertPolicyOptionsModel.PolicyID = core.StringPtr("testString")
				updateAlertPolicyOptionsModel.Name = core.StringPtr("My Alert Policy")
				updateAlertPolicyOptionsModel.Description = core.StringPtr("A description for my alert policy")
				updateAlertPolicyOptionsModel.Enabled = core.BoolPtr(true)
				updateAlertPolicyOptionsModel.AlertType = core.StringPtr("dos_attack_l7")
				updateAlertPolicyOptionsModel.Mechanisms = updateAlertPolicyInputMechanismsModel
				updateAlertPolicyOptionsModel.Conditions = map[string]interface{}{"anyKey": "anyValue"}
				updateAlertPolicyOptionsModel.Filters = map[string]interface{}{"anyKey": "anyValue"}
				updateAlertPolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := alertsService.UpdateAlertPolicy(updateAlertPolicyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				alertsService.EnableRetries(0, 0)
				result, response, operationErr = alertsService.UpdateAlertPolicy(updateAlertPolicyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateAlertPolicy(updateAlertPolicyOptions *UpdateAlertPolicyOptions)`, func() {
		crn := "testString"
		updateAlertPolicyPath := "/v1/testString/alerting/policies/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateAlertPolicyPath))
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
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [{"id": "ID"}], "messages": [{"id": "ID"}], "result": {"id": "f0413b106d2c4aa9b1553d5d0209c522"}}`)
				}))
			})
			It(`Invoke UpdateAlertPolicy successfully with retries`, func() {
				alertsService, serviceErr := alertsv1.NewAlertsV1(&alertsv1.AlertsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(alertsService).ToNot(BeNil())
				alertsService.EnableRetries(0, 0)

				// Construct an instance of the UpdateAlertPolicyInputMechanismsEmailItem model
				updateAlertPolicyInputMechanismsEmailItemModel := new(alertsv1.UpdateAlertPolicyInputMechanismsEmailItem)
				updateAlertPolicyInputMechanismsEmailItemModel.ID = core.StringPtr("mynotifications@email.com")

				// Construct an instance of the UpdateAlertPolicyInputMechanismsWebhooksItem model
				updateAlertPolicyInputMechanismsWebhooksItemModel := new(alertsv1.UpdateAlertPolicyInputMechanismsWebhooksItem)
				updateAlertPolicyInputMechanismsWebhooksItemModel.ID = core.StringPtr("f0413b106d2c4aa9b1553d5d0209c522")

				// Construct an instance of the UpdateAlertPolicyInputMechanisms model
				updateAlertPolicyInputMechanismsModel := new(alertsv1.UpdateAlertPolicyInputMechanisms)
				updateAlertPolicyInputMechanismsModel.Email = []alertsv1.UpdateAlertPolicyInputMechanismsEmailItem{*updateAlertPolicyInputMechanismsEmailItemModel}
				updateAlertPolicyInputMechanismsModel.Webhooks = []alertsv1.UpdateAlertPolicyInputMechanismsWebhooksItem{*updateAlertPolicyInputMechanismsWebhooksItemModel}

				// Construct an instance of the UpdateAlertPolicyOptions model
				updateAlertPolicyOptionsModel := new(alertsv1.UpdateAlertPolicyOptions)
				updateAlertPolicyOptionsModel.PolicyID = core.StringPtr("testString")
				updateAlertPolicyOptionsModel.Name = core.StringPtr("My Alert Policy")
				updateAlertPolicyOptionsModel.Description = core.StringPtr("A description for my alert policy")
				updateAlertPolicyOptionsModel.Enabled = core.BoolPtr(true)
				updateAlertPolicyOptionsModel.AlertType = core.StringPtr("dos_attack_l7")
				updateAlertPolicyOptionsModel.Mechanisms = updateAlertPolicyInputMechanismsModel
				updateAlertPolicyOptionsModel.Conditions = map[string]interface{}{"anyKey": "anyValue"}
				updateAlertPolicyOptionsModel.Filters = map[string]interface{}{"anyKey": "anyValue"}
				updateAlertPolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := alertsService.UpdateAlertPolicyWithContext(ctx, updateAlertPolicyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				alertsService.DisableRetries()
				result, response, operationErr := alertsService.UpdateAlertPolicy(updateAlertPolicyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = alertsService.UpdateAlertPolicyWithContext(ctx, updateAlertPolicyOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(updateAlertPolicyPath))
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
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [{"id": "ID"}], "messages": [{"id": "ID"}], "result": {"id": "f0413b106d2c4aa9b1553d5d0209c522"}}`)
				}))
			})
			It(`Invoke UpdateAlertPolicy successfully`, func() {
				alertsService, serviceErr := alertsv1.NewAlertsV1(&alertsv1.AlertsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(alertsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := alertsService.UpdateAlertPolicy(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateAlertPolicyInputMechanismsEmailItem model
				updateAlertPolicyInputMechanismsEmailItemModel := new(alertsv1.UpdateAlertPolicyInputMechanismsEmailItem)
				updateAlertPolicyInputMechanismsEmailItemModel.ID = core.StringPtr("mynotifications@email.com")

				// Construct an instance of the UpdateAlertPolicyInputMechanismsWebhooksItem model
				updateAlertPolicyInputMechanismsWebhooksItemModel := new(alertsv1.UpdateAlertPolicyInputMechanismsWebhooksItem)
				updateAlertPolicyInputMechanismsWebhooksItemModel.ID = core.StringPtr("f0413b106d2c4aa9b1553d5d0209c522")

				// Construct an instance of the UpdateAlertPolicyInputMechanisms model
				updateAlertPolicyInputMechanismsModel := new(alertsv1.UpdateAlertPolicyInputMechanisms)
				updateAlertPolicyInputMechanismsModel.Email = []alertsv1.UpdateAlertPolicyInputMechanismsEmailItem{*updateAlertPolicyInputMechanismsEmailItemModel}
				updateAlertPolicyInputMechanismsModel.Webhooks = []alertsv1.UpdateAlertPolicyInputMechanismsWebhooksItem{*updateAlertPolicyInputMechanismsWebhooksItemModel}

				// Construct an instance of the UpdateAlertPolicyOptions model
				updateAlertPolicyOptionsModel := new(alertsv1.UpdateAlertPolicyOptions)
				updateAlertPolicyOptionsModel.PolicyID = core.StringPtr("testString")
				updateAlertPolicyOptionsModel.Name = core.StringPtr("My Alert Policy")
				updateAlertPolicyOptionsModel.Description = core.StringPtr("A description for my alert policy")
				updateAlertPolicyOptionsModel.Enabled = core.BoolPtr(true)
				updateAlertPolicyOptionsModel.AlertType = core.StringPtr("dos_attack_l7")
				updateAlertPolicyOptionsModel.Mechanisms = updateAlertPolicyInputMechanismsModel
				updateAlertPolicyOptionsModel.Conditions = map[string]interface{}{"anyKey": "anyValue"}
				updateAlertPolicyOptionsModel.Filters = map[string]interface{}{"anyKey": "anyValue"}
				updateAlertPolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = alertsService.UpdateAlertPolicy(updateAlertPolicyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateAlertPolicy with error: Operation validation and request error`, func() {
				alertsService, serviceErr := alertsv1.NewAlertsV1(&alertsv1.AlertsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(alertsService).ToNot(BeNil())

				// Construct an instance of the UpdateAlertPolicyInputMechanismsEmailItem model
				updateAlertPolicyInputMechanismsEmailItemModel := new(alertsv1.UpdateAlertPolicyInputMechanismsEmailItem)
				updateAlertPolicyInputMechanismsEmailItemModel.ID = core.StringPtr("mynotifications@email.com")

				// Construct an instance of the UpdateAlertPolicyInputMechanismsWebhooksItem model
				updateAlertPolicyInputMechanismsWebhooksItemModel := new(alertsv1.UpdateAlertPolicyInputMechanismsWebhooksItem)
				updateAlertPolicyInputMechanismsWebhooksItemModel.ID = core.StringPtr("f0413b106d2c4aa9b1553d5d0209c522")

				// Construct an instance of the UpdateAlertPolicyInputMechanisms model
				updateAlertPolicyInputMechanismsModel := new(alertsv1.UpdateAlertPolicyInputMechanisms)
				updateAlertPolicyInputMechanismsModel.Email = []alertsv1.UpdateAlertPolicyInputMechanismsEmailItem{*updateAlertPolicyInputMechanismsEmailItemModel}
				updateAlertPolicyInputMechanismsModel.Webhooks = []alertsv1.UpdateAlertPolicyInputMechanismsWebhooksItem{*updateAlertPolicyInputMechanismsWebhooksItemModel}

				// Construct an instance of the UpdateAlertPolicyOptions model
				updateAlertPolicyOptionsModel := new(alertsv1.UpdateAlertPolicyOptions)
				updateAlertPolicyOptionsModel.PolicyID = core.StringPtr("testString")
				updateAlertPolicyOptionsModel.Name = core.StringPtr("My Alert Policy")
				updateAlertPolicyOptionsModel.Description = core.StringPtr("A description for my alert policy")
				updateAlertPolicyOptionsModel.Enabled = core.BoolPtr(true)
				updateAlertPolicyOptionsModel.AlertType = core.StringPtr("dos_attack_l7")
				updateAlertPolicyOptionsModel.Mechanisms = updateAlertPolicyInputMechanismsModel
				updateAlertPolicyOptionsModel.Conditions = map[string]interface{}{"anyKey": "anyValue"}
				updateAlertPolicyOptionsModel.Filters = map[string]interface{}{"anyKey": "anyValue"}
				updateAlertPolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := alertsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := alertsService.UpdateAlertPolicy(updateAlertPolicyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateAlertPolicyOptions model with no property values
				updateAlertPolicyOptionsModelNew := new(alertsv1.UpdateAlertPolicyOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = alertsService.UpdateAlertPolicy(updateAlertPolicyOptionsModelNew)
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
			It(`Invoke UpdateAlertPolicy successfully`, func() {
				alertsService, serviceErr := alertsv1.NewAlertsV1(&alertsv1.AlertsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(alertsService).ToNot(BeNil())

				// Construct an instance of the UpdateAlertPolicyInputMechanismsEmailItem model
				updateAlertPolicyInputMechanismsEmailItemModel := new(alertsv1.UpdateAlertPolicyInputMechanismsEmailItem)
				updateAlertPolicyInputMechanismsEmailItemModel.ID = core.StringPtr("mynotifications@email.com")

				// Construct an instance of the UpdateAlertPolicyInputMechanismsWebhooksItem model
				updateAlertPolicyInputMechanismsWebhooksItemModel := new(alertsv1.UpdateAlertPolicyInputMechanismsWebhooksItem)
				updateAlertPolicyInputMechanismsWebhooksItemModel.ID = core.StringPtr("f0413b106d2c4aa9b1553d5d0209c522")

				// Construct an instance of the UpdateAlertPolicyInputMechanisms model
				updateAlertPolicyInputMechanismsModel := new(alertsv1.UpdateAlertPolicyInputMechanisms)
				updateAlertPolicyInputMechanismsModel.Email = []alertsv1.UpdateAlertPolicyInputMechanismsEmailItem{*updateAlertPolicyInputMechanismsEmailItemModel}
				updateAlertPolicyInputMechanismsModel.Webhooks = []alertsv1.UpdateAlertPolicyInputMechanismsWebhooksItem{*updateAlertPolicyInputMechanismsWebhooksItemModel}

				// Construct an instance of the UpdateAlertPolicyOptions model
				updateAlertPolicyOptionsModel := new(alertsv1.UpdateAlertPolicyOptions)
				updateAlertPolicyOptionsModel.PolicyID = core.StringPtr("testString")
				updateAlertPolicyOptionsModel.Name = core.StringPtr("My Alert Policy")
				updateAlertPolicyOptionsModel.Description = core.StringPtr("A description for my alert policy")
				updateAlertPolicyOptionsModel.Enabled = core.BoolPtr(true)
				updateAlertPolicyOptionsModel.AlertType = core.StringPtr("dos_attack_l7")
				updateAlertPolicyOptionsModel.Mechanisms = updateAlertPolicyInputMechanismsModel
				updateAlertPolicyOptionsModel.Conditions = map[string]interface{}{"anyKey": "anyValue"}
				updateAlertPolicyOptionsModel.Filters = map[string]interface{}{"anyKey": "anyValue"}
				updateAlertPolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := alertsService.UpdateAlertPolicy(updateAlertPolicyOptionsModel)
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
	Describe(`DeleteAlertPolicy(deleteAlertPolicyOptions *DeleteAlertPolicyOptions) - Operation response error`, func() {
		crn := "testString"
		deleteAlertPolicyPath := "/v1/testString/alerting/policies/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteAlertPolicyPath))
					Expect(req.Method).To(Equal("DELETE"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke DeleteAlertPolicy with error: Operation response processing error`, func() {
				alertsService, serviceErr := alertsv1.NewAlertsV1(&alertsv1.AlertsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(alertsService).ToNot(BeNil())

				// Construct an instance of the DeleteAlertPolicyOptions model
				deleteAlertPolicyOptionsModel := new(alertsv1.DeleteAlertPolicyOptions)
				deleteAlertPolicyOptionsModel.PolicyID = core.StringPtr("testString")
				deleteAlertPolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := alertsService.DeleteAlertPolicy(deleteAlertPolicyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				alertsService.EnableRetries(0, 0)
				result, response, operationErr = alertsService.DeleteAlertPolicy(deleteAlertPolicyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteAlertPolicy(deleteAlertPolicyOptions *DeleteAlertPolicyOptions)`, func() {
		crn := "testString"
		deleteAlertPolicyPath := "/v1/testString/alerting/policies/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteAlertPolicyPath))
					Expect(req.Method).To(Equal("DELETE"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [{"id": "ID"}], "messages": [{"id": "ID"}], "result": {"id": "f0413b106d2c4aa9b1553d5d0209c522"}}`)
				}))
			})
			It(`Invoke DeleteAlertPolicy successfully with retries`, func() {
				alertsService, serviceErr := alertsv1.NewAlertsV1(&alertsv1.AlertsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(alertsService).ToNot(BeNil())
				alertsService.EnableRetries(0, 0)

				// Construct an instance of the DeleteAlertPolicyOptions model
				deleteAlertPolicyOptionsModel := new(alertsv1.DeleteAlertPolicyOptions)
				deleteAlertPolicyOptionsModel.PolicyID = core.StringPtr("testString")
				deleteAlertPolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := alertsService.DeleteAlertPolicyWithContext(ctx, deleteAlertPolicyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				alertsService.DisableRetries()
				result, response, operationErr := alertsService.DeleteAlertPolicy(deleteAlertPolicyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = alertsService.DeleteAlertPolicyWithContext(ctx, deleteAlertPolicyOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(deleteAlertPolicyPath))
					Expect(req.Method).To(Equal("DELETE"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [{"id": "ID"}], "messages": [{"id": "ID"}], "result": {"id": "f0413b106d2c4aa9b1553d5d0209c522"}}`)
				}))
			})
			It(`Invoke DeleteAlertPolicy successfully`, func() {
				alertsService, serviceErr := alertsv1.NewAlertsV1(&alertsv1.AlertsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(alertsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := alertsService.DeleteAlertPolicy(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DeleteAlertPolicyOptions model
				deleteAlertPolicyOptionsModel := new(alertsv1.DeleteAlertPolicyOptions)
				deleteAlertPolicyOptionsModel.PolicyID = core.StringPtr("testString")
				deleteAlertPolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = alertsService.DeleteAlertPolicy(deleteAlertPolicyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke DeleteAlertPolicy with error: Operation validation and request error`, func() {
				alertsService, serviceErr := alertsv1.NewAlertsV1(&alertsv1.AlertsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(alertsService).ToNot(BeNil())

				// Construct an instance of the DeleteAlertPolicyOptions model
				deleteAlertPolicyOptionsModel := new(alertsv1.DeleteAlertPolicyOptions)
				deleteAlertPolicyOptionsModel.PolicyID = core.StringPtr("testString")
				deleteAlertPolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := alertsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := alertsService.DeleteAlertPolicy(deleteAlertPolicyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the DeleteAlertPolicyOptions model with no property values
				deleteAlertPolicyOptionsModelNew := new(alertsv1.DeleteAlertPolicyOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = alertsService.DeleteAlertPolicy(deleteAlertPolicyOptionsModelNew)
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
			It(`Invoke DeleteAlertPolicy successfully`, func() {
				alertsService, serviceErr := alertsv1.NewAlertsV1(&alertsv1.AlertsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(alertsService).ToNot(BeNil())

				// Construct an instance of the DeleteAlertPolicyOptions model
				deleteAlertPolicyOptionsModel := new(alertsv1.DeleteAlertPolicyOptions)
				deleteAlertPolicyOptionsModel.PolicyID = core.StringPtr("testString")
				deleteAlertPolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := alertsService.DeleteAlertPolicy(deleteAlertPolicyOptionsModel)
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
			alertsService, _ := alertsv1.NewAlertsV1(&alertsv1.AlertsV1Options{
				URL:           "http://alertsv1modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
				Crn: core.StringPtr(crn),
			})
			It(`Invoke NewCreateAlertPolicyOptions successfully`, func() {
				// Construct an instance of the CreateAlertPolicyInputMechanismsEmailItem model
				createAlertPolicyInputMechanismsEmailItemModel := new(alertsv1.CreateAlertPolicyInputMechanismsEmailItem)
				Expect(createAlertPolicyInputMechanismsEmailItemModel).ToNot(BeNil())
				createAlertPolicyInputMechanismsEmailItemModel.ID = core.StringPtr("mynotifications@email.com")
				Expect(createAlertPolicyInputMechanismsEmailItemModel.ID).To(Equal(core.StringPtr("mynotifications@email.com")))

				// Construct an instance of the CreateAlertPolicyInputMechanismsWebhooksItem model
				createAlertPolicyInputMechanismsWebhooksItemModel := new(alertsv1.CreateAlertPolicyInputMechanismsWebhooksItem)
				Expect(createAlertPolicyInputMechanismsWebhooksItemModel).ToNot(BeNil())
				createAlertPolicyInputMechanismsWebhooksItemModel.ID = core.StringPtr("f0413b106d2c4aa9b1553d5d0209c522")
				Expect(createAlertPolicyInputMechanismsWebhooksItemModel.ID).To(Equal(core.StringPtr("f0413b106d2c4aa9b1553d5d0209c522")))

				// Construct an instance of the CreateAlertPolicyInputMechanisms model
				createAlertPolicyInputMechanismsModel := new(alertsv1.CreateAlertPolicyInputMechanisms)
				Expect(createAlertPolicyInputMechanismsModel).ToNot(BeNil())
				createAlertPolicyInputMechanismsModel.Email = []alertsv1.CreateAlertPolicyInputMechanismsEmailItem{*createAlertPolicyInputMechanismsEmailItemModel}
				createAlertPolicyInputMechanismsModel.Webhooks = []alertsv1.CreateAlertPolicyInputMechanismsWebhooksItem{*createAlertPolicyInputMechanismsWebhooksItemModel}
				Expect(createAlertPolicyInputMechanismsModel.Email).To(Equal([]alertsv1.CreateAlertPolicyInputMechanismsEmailItem{*createAlertPolicyInputMechanismsEmailItemModel}))
				Expect(createAlertPolicyInputMechanismsModel.Webhooks).To(Equal([]alertsv1.CreateAlertPolicyInputMechanismsWebhooksItem{*createAlertPolicyInputMechanismsWebhooksItemModel}))

				// Construct an instance of the CreateAlertPolicyOptions model
				createAlertPolicyOptionsModel := alertsService.NewCreateAlertPolicyOptions()
				createAlertPolicyOptionsModel.SetName("My Alert Policy")
				createAlertPolicyOptionsModel.SetDescription("A description for my alert policy")
				createAlertPolicyOptionsModel.SetEnabled(true)
				createAlertPolicyOptionsModel.SetAlertType("dos_attack_l7")
				createAlertPolicyOptionsModel.SetMechanisms(createAlertPolicyInputMechanismsModel)
				createAlertPolicyOptionsModel.SetFilters(map[string]interface{}{"anyKey": "anyValue"})
				createAlertPolicyOptionsModel.SetConditions(map[string]interface{}{"anyKey": "anyValue"})
				createAlertPolicyOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createAlertPolicyOptionsModel).ToNot(BeNil())
				Expect(createAlertPolicyOptionsModel.Name).To(Equal(core.StringPtr("My Alert Policy")))
				Expect(createAlertPolicyOptionsModel.Description).To(Equal(core.StringPtr("A description for my alert policy")))
				Expect(createAlertPolicyOptionsModel.Enabled).To(Equal(core.BoolPtr(true)))
				Expect(createAlertPolicyOptionsModel.AlertType).To(Equal(core.StringPtr("dos_attack_l7")))
				Expect(createAlertPolicyOptionsModel.Mechanisms).To(Equal(createAlertPolicyInputMechanismsModel))
				Expect(createAlertPolicyOptionsModel.Filters).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))
				Expect(createAlertPolicyOptionsModel.Conditions).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))
				Expect(createAlertPolicyOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteAlertPolicyOptions successfully`, func() {
				// Construct an instance of the DeleteAlertPolicyOptions model
				policyID := "testString"
				deleteAlertPolicyOptionsModel := alertsService.NewDeleteAlertPolicyOptions(policyID)
				deleteAlertPolicyOptionsModel.SetPolicyID("testString")
				deleteAlertPolicyOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteAlertPolicyOptionsModel).ToNot(BeNil())
				Expect(deleteAlertPolicyOptionsModel.PolicyID).To(Equal(core.StringPtr("testString")))
				Expect(deleteAlertPolicyOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetAlertPoliciesOptions successfully`, func() {
				// Construct an instance of the GetAlertPoliciesOptions model
				getAlertPoliciesOptionsModel := alertsService.NewGetAlertPoliciesOptions()
				getAlertPoliciesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getAlertPoliciesOptionsModel).ToNot(BeNil())
				Expect(getAlertPoliciesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetAlertPolicyOptions successfully`, func() {
				// Construct an instance of the GetAlertPolicyOptions model
				policyID := "testString"
				getAlertPolicyOptionsModel := alertsService.NewGetAlertPolicyOptions(policyID)
				getAlertPolicyOptionsModel.SetPolicyID("testString")
				getAlertPolicyOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getAlertPolicyOptionsModel).ToNot(BeNil())
				Expect(getAlertPolicyOptionsModel.PolicyID).To(Equal(core.StringPtr("testString")))
				Expect(getAlertPolicyOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateAlertPolicyOptions successfully`, func() {
				// Construct an instance of the UpdateAlertPolicyInputMechanismsEmailItem model
				updateAlertPolicyInputMechanismsEmailItemModel := new(alertsv1.UpdateAlertPolicyInputMechanismsEmailItem)
				Expect(updateAlertPolicyInputMechanismsEmailItemModel).ToNot(BeNil())
				updateAlertPolicyInputMechanismsEmailItemModel.ID = core.StringPtr("mynotifications@email.com")
				Expect(updateAlertPolicyInputMechanismsEmailItemModel.ID).To(Equal(core.StringPtr("mynotifications@email.com")))

				// Construct an instance of the UpdateAlertPolicyInputMechanismsWebhooksItem model
				updateAlertPolicyInputMechanismsWebhooksItemModel := new(alertsv1.UpdateAlertPolicyInputMechanismsWebhooksItem)
				Expect(updateAlertPolicyInputMechanismsWebhooksItemModel).ToNot(BeNil())
				updateAlertPolicyInputMechanismsWebhooksItemModel.ID = core.StringPtr("f0413b106d2c4aa9b1553d5d0209c522")
				Expect(updateAlertPolicyInputMechanismsWebhooksItemModel.ID).To(Equal(core.StringPtr("f0413b106d2c4aa9b1553d5d0209c522")))

				// Construct an instance of the UpdateAlertPolicyInputMechanisms model
				updateAlertPolicyInputMechanismsModel := new(alertsv1.UpdateAlertPolicyInputMechanisms)
				Expect(updateAlertPolicyInputMechanismsModel).ToNot(BeNil())
				updateAlertPolicyInputMechanismsModel.Email = []alertsv1.UpdateAlertPolicyInputMechanismsEmailItem{*updateAlertPolicyInputMechanismsEmailItemModel}
				updateAlertPolicyInputMechanismsModel.Webhooks = []alertsv1.UpdateAlertPolicyInputMechanismsWebhooksItem{*updateAlertPolicyInputMechanismsWebhooksItemModel}
				Expect(updateAlertPolicyInputMechanismsModel.Email).To(Equal([]alertsv1.UpdateAlertPolicyInputMechanismsEmailItem{*updateAlertPolicyInputMechanismsEmailItemModel}))
				Expect(updateAlertPolicyInputMechanismsModel.Webhooks).To(Equal([]alertsv1.UpdateAlertPolicyInputMechanismsWebhooksItem{*updateAlertPolicyInputMechanismsWebhooksItemModel}))

				// Construct an instance of the UpdateAlertPolicyOptions model
				policyID := "testString"
				updateAlertPolicyOptionsModel := alertsService.NewUpdateAlertPolicyOptions(policyID)
				updateAlertPolicyOptionsModel.SetPolicyID("testString")
				updateAlertPolicyOptionsModel.SetName("My Alert Policy")
				updateAlertPolicyOptionsModel.SetDescription("A description for my alert policy")
				updateAlertPolicyOptionsModel.SetEnabled(true)
				updateAlertPolicyOptionsModel.SetAlertType("dos_attack_l7")
				updateAlertPolicyOptionsModel.SetMechanisms(updateAlertPolicyInputMechanismsModel)
				updateAlertPolicyOptionsModel.SetConditions(map[string]interface{}{"anyKey": "anyValue"})
				updateAlertPolicyOptionsModel.SetFilters(map[string]interface{}{"anyKey": "anyValue"})
				updateAlertPolicyOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateAlertPolicyOptionsModel).ToNot(BeNil())
				Expect(updateAlertPolicyOptionsModel.PolicyID).To(Equal(core.StringPtr("testString")))
				Expect(updateAlertPolicyOptionsModel.Name).To(Equal(core.StringPtr("My Alert Policy")))
				Expect(updateAlertPolicyOptionsModel.Description).To(Equal(core.StringPtr("A description for my alert policy")))
				Expect(updateAlertPolicyOptionsModel.Enabled).To(Equal(core.BoolPtr(true)))
				Expect(updateAlertPolicyOptionsModel.AlertType).To(Equal(core.StringPtr("dos_attack_l7")))
				Expect(updateAlertPolicyOptionsModel.Mechanisms).To(Equal(updateAlertPolicyInputMechanismsModel))
				Expect(updateAlertPolicyOptionsModel.Conditions).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))
				Expect(updateAlertPolicyOptionsModel.Filters).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))
				Expect(updateAlertPolicyOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
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
