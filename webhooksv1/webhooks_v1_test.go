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

package webhooksv1_test

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
	"github.com/IBM/networking-go-sdk/webhooksv1"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe(`WebhooksV1`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		crn := "testString"
		It(`Instantiate service client`, func() {
			webhooksService, serviceErr := webhooksv1.NewWebhooksV1(&webhooksv1.WebhooksV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
				Crn:           core.StringPtr(crn),
			})
			Expect(webhooksService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			webhooksService, serviceErr := webhooksv1.NewWebhooksV1(&webhooksv1.WebhooksV1Options{
				URL: "{BAD_URL_STRING",
				Crn: core.StringPtr(crn),
			})
			Expect(webhooksService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			webhooksService, serviceErr := webhooksv1.NewWebhooksV1(&webhooksv1.WebhooksV1Options{
				URL: "https://webhooksv1/api",
				Crn: core.StringPtr(crn),
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(webhooksService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Validation Error`, func() {
			webhooksService, serviceErr := webhooksv1.NewWebhooksV1(&webhooksv1.WebhooksV1Options{})
			Expect(webhooksService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		crn := "testString"
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"WEBHOOKS_URL":       "https://webhooksv1/api",
				"WEBHOOKS_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				webhooksService, serviceErr := webhooksv1.NewWebhooksV1UsingExternalConfig(&webhooksv1.WebhooksV1Options{
					Crn: core.StringPtr(crn),
				})
				Expect(webhooksService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := webhooksService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != webhooksService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(webhooksService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(webhooksService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				webhooksService, serviceErr := webhooksv1.NewWebhooksV1UsingExternalConfig(&webhooksv1.WebhooksV1Options{
					URL: "https://testService/api",
					Crn: core.StringPtr(crn),
				})
				Expect(webhooksService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(webhooksService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := webhooksService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != webhooksService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(webhooksService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(webhooksService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				webhooksService, serviceErr := webhooksv1.NewWebhooksV1UsingExternalConfig(&webhooksv1.WebhooksV1Options{
					Crn: core.StringPtr(crn),
				})
				err := webhooksService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(webhooksService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(webhooksService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := webhooksService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != webhooksService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(webhooksService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(webhooksService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"WEBHOOKS_URL":       "https://webhooksv1/api",
				"WEBHOOKS_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			webhooksService, serviceErr := webhooksv1.NewWebhooksV1UsingExternalConfig(&webhooksv1.WebhooksV1Options{
				Crn: core.StringPtr(crn),
			})

			It(`Instantiate service client with error`, func() {
				Expect(webhooksService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"WEBHOOKS_AUTH_TYPE": "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			webhooksService, serviceErr := webhooksv1.NewWebhooksV1UsingExternalConfig(&webhooksv1.WebhooksV1Options{
				URL: "{BAD_URL_STRING",
				Crn: core.StringPtr(crn),
			})

			It(`Instantiate service client with error`, func() {
				Expect(webhooksService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = webhooksv1.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`ListWebhooks(listWebhooksOptions *ListWebhooksOptions) - Operation response error`, func() {
		crn := "testString"
		listWebhooksPath := "/v1/testString/alerting/destinations/webhooks"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listWebhooksPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListWebhooks with error: Operation response processing error`, func() {
				webhooksService, serviceErr := webhooksv1.NewWebhooksV1(&webhooksv1.WebhooksV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(webhooksService).ToNot(BeNil())

				// Construct an instance of the ListWebhooksOptions model
				listWebhooksOptionsModel := new(webhooksv1.ListWebhooksOptions)
				listWebhooksOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := webhooksService.ListWebhooks(listWebhooksOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				webhooksService.EnableRetries(0, 0)
				result, response, operationErr = webhooksService.ListWebhooks(listWebhooksOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListWebhooks(listWebhooksOptions *ListWebhooksOptions)`, func() {
		crn := "testString"
		listWebhooksPath := "/v1/testString/alerting/destinations/webhooks"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listWebhooksPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": [{"id": "6d16fcab-3e80-44b3-b59b-a3716237832e", "name": "My Slack Alert Webhook", "url": "https://hooks.slack.com/services/Ds3fdBFbV/456464Gdd", "type": "generic", "created_at": "2021-09-15T16:33:31.834209Z", "last_success": "2021-09-15T16:33:31.834209Z", "last_failure": "2021-09-15T16:33:31.834209Z"}]}`)
				}))
			})
			It(`Invoke ListWebhooks successfully with retries`, func() {
				webhooksService, serviceErr := webhooksv1.NewWebhooksV1(&webhooksv1.WebhooksV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(webhooksService).ToNot(BeNil())
				webhooksService.EnableRetries(0, 0)

				// Construct an instance of the ListWebhooksOptions model
				listWebhooksOptionsModel := new(webhooksv1.ListWebhooksOptions)
				listWebhooksOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := webhooksService.ListWebhooksWithContext(ctx, listWebhooksOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				webhooksService.DisableRetries()
				result, response, operationErr := webhooksService.ListWebhooks(listWebhooksOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = webhooksService.ListWebhooksWithContext(ctx, listWebhooksOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listWebhooksPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": [{"id": "6d16fcab-3e80-44b3-b59b-a3716237832e", "name": "My Slack Alert Webhook", "url": "https://hooks.slack.com/services/Ds3fdBFbV/456464Gdd", "type": "generic", "created_at": "2021-09-15T16:33:31.834209Z", "last_success": "2021-09-15T16:33:31.834209Z", "last_failure": "2021-09-15T16:33:31.834209Z"}]}`)
				}))
			})
			It(`Invoke ListWebhooks successfully`, func() {
				webhooksService, serviceErr := webhooksv1.NewWebhooksV1(&webhooksv1.WebhooksV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(webhooksService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := webhooksService.ListWebhooks(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListWebhooksOptions model
				listWebhooksOptionsModel := new(webhooksv1.ListWebhooksOptions)
				listWebhooksOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = webhooksService.ListWebhooks(listWebhooksOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListWebhooks with error: Operation request error`, func() {
				webhooksService, serviceErr := webhooksv1.NewWebhooksV1(&webhooksv1.WebhooksV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(webhooksService).ToNot(BeNil())

				// Construct an instance of the ListWebhooksOptions model
				listWebhooksOptionsModel := new(webhooksv1.ListWebhooksOptions)
				listWebhooksOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := webhooksService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := webhooksService.ListWebhooks(listWebhooksOptionsModel)
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
			It(`Invoke ListWebhooks successfully`, func() {
				webhooksService, serviceErr := webhooksv1.NewWebhooksV1(&webhooksv1.WebhooksV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(webhooksService).ToNot(BeNil())

				// Construct an instance of the ListWebhooksOptions model
				listWebhooksOptionsModel := new(webhooksv1.ListWebhooksOptions)
				listWebhooksOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := webhooksService.ListWebhooks(listWebhooksOptionsModel)
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
	Describe(`CreateAlertWebhook(createAlertWebhookOptions *CreateAlertWebhookOptions) - Operation response error`, func() {
		crn := "testString"
		createAlertWebhookPath := "/v1/testString/alerting/destinations/webhooks"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createAlertWebhookPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateAlertWebhook with error: Operation response processing error`, func() {
				webhooksService, serviceErr := webhooksv1.NewWebhooksV1(&webhooksv1.WebhooksV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(webhooksService).ToNot(BeNil())

				// Construct an instance of the CreateAlertWebhookOptions model
				createAlertWebhookOptionsModel := new(webhooksv1.CreateAlertWebhookOptions)
				createAlertWebhookOptionsModel.Name = core.StringPtr("My Slack Alert Webhook")
				createAlertWebhookOptionsModel.URL = core.StringPtr("https://hooks.slack.com/services/Ds3fdBFbV/456464Gdd")
				createAlertWebhookOptionsModel.Secret = core.StringPtr("ff1d9b80-b51d-4a06-bf67-6752fae1eb74")
				createAlertWebhookOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := webhooksService.CreateAlertWebhook(createAlertWebhookOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				webhooksService.EnableRetries(0, 0)
				result, response, operationErr = webhooksService.CreateAlertWebhook(createAlertWebhookOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateAlertWebhook(createAlertWebhookOptions *CreateAlertWebhookOptions)`, func() {
		crn := "testString"
		createAlertWebhookPath := "/v1/testString/alerting/destinations/webhooks"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createAlertWebhookPath))
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
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": "6d16fcab-3e80-44b3-b59b-a3716237832e"}}`)
				}))
			})
			It(`Invoke CreateAlertWebhook successfully with retries`, func() {
				webhooksService, serviceErr := webhooksv1.NewWebhooksV1(&webhooksv1.WebhooksV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(webhooksService).ToNot(BeNil())
				webhooksService.EnableRetries(0, 0)

				// Construct an instance of the CreateAlertWebhookOptions model
				createAlertWebhookOptionsModel := new(webhooksv1.CreateAlertWebhookOptions)
				createAlertWebhookOptionsModel.Name = core.StringPtr("My Slack Alert Webhook")
				createAlertWebhookOptionsModel.URL = core.StringPtr("https://hooks.slack.com/services/Ds3fdBFbV/456464Gdd")
				createAlertWebhookOptionsModel.Secret = core.StringPtr("ff1d9b80-b51d-4a06-bf67-6752fae1eb74")
				createAlertWebhookOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := webhooksService.CreateAlertWebhookWithContext(ctx, createAlertWebhookOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				webhooksService.DisableRetries()
				result, response, operationErr := webhooksService.CreateAlertWebhook(createAlertWebhookOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = webhooksService.CreateAlertWebhookWithContext(ctx, createAlertWebhookOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createAlertWebhookPath))
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
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": "6d16fcab-3e80-44b3-b59b-a3716237832e"}}`)
				}))
			})
			It(`Invoke CreateAlertWebhook successfully`, func() {
				webhooksService, serviceErr := webhooksv1.NewWebhooksV1(&webhooksv1.WebhooksV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(webhooksService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := webhooksService.CreateAlertWebhook(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateAlertWebhookOptions model
				createAlertWebhookOptionsModel := new(webhooksv1.CreateAlertWebhookOptions)
				createAlertWebhookOptionsModel.Name = core.StringPtr("My Slack Alert Webhook")
				createAlertWebhookOptionsModel.URL = core.StringPtr("https://hooks.slack.com/services/Ds3fdBFbV/456464Gdd")
				createAlertWebhookOptionsModel.Secret = core.StringPtr("ff1d9b80-b51d-4a06-bf67-6752fae1eb74")
				createAlertWebhookOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = webhooksService.CreateAlertWebhook(createAlertWebhookOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateAlertWebhook with error: Operation request error`, func() {
				webhooksService, serviceErr := webhooksv1.NewWebhooksV1(&webhooksv1.WebhooksV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(webhooksService).ToNot(BeNil())

				// Construct an instance of the CreateAlertWebhookOptions model
				createAlertWebhookOptionsModel := new(webhooksv1.CreateAlertWebhookOptions)
				createAlertWebhookOptionsModel.Name = core.StringPtr("My Slack Alert Webhook")
				createAlertWebhookOptionsModel.URL = core.StringPtr("https://hooks.slack.com/services/Ds3fdBFbV/456464Gdd")
				createAlertWebhookOptionsModel.Secret = core.StringPtr("ff1d9b80-b51d-4a06-bf67-6752fae1eb74")
				createAlertWebhookOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := webhooksService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := webhooksService.CreateAlertWebhook(createAlertWebhookOptionsModel)
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
			It(`Invoke CreateAlertWebhook successfully`, func() {
				webhooksService, serviceErr := webhooksv1.NewWebhooksV1(&webhooksv1.WebhooksV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(webhooksService).ToNot(BeNil())

				// Construct an instance of the CreateAlertWebhookOptions model
				createAlertWebhookOptionsModel := new(webhooksv1.CreateAlertWebhookOptions)
				createAlertWebhookOptionsModel.Name = core.StringPtr("My Slack Alert Webhook")
				createAlertWebhookOptionsModel.URL = core.StringPtr("https://hooks.slack.com/services/Ds3fdBFbV/456464Gdd")
				createAlertWebhookOptionsModel.Secret = core.StringPtr("ff1d9b80-b51d-4a06-bf67-6752fae1eb74")
				createAlertWebhookOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := webhooksService.CreateAlertWebhook(createAlertWebhookOptionsModel)
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
	Describe(`GetWebhook(getWebhookOptions *GetWebhookOptions) - Operation response error`, func() {
		crn := "testString"
		getWebhookPath := "/v1/testString/alerting/destinations/webhooks/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getWebhookPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetWebhook with error: Operation response processing error`, func() {
				webhooksService, serviceErr := webhooksv1.NewWebhooksV1(&webhooksv1.WebhooksV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(webhooksService).ToNot(BeNil())

				// Construct an instance of the GetWebhookOptions model
				getWebhookOptionsModel := new(webhooksv1.GetWebhookOptions)
				getWebhookOptionsModel.WebhookID = core.StringPtr("testString")
				getWebhookOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := webhooksService.GetWebhook(getWebhookOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				webhooksService.EnableRetries(0, 0)
				result, response, operationErr = webhooksService.GetWebhook(getWebhookOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetWebhook(getWebhookOptions *GetWebhookOptions)`, func() {
		crn := "testString"
		getWebhookPath := "/v1/testString/alerting/destinations/webhooks/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getWebhookPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": "6d16fcab-3e80-44b3-b59b-a3716237832e", "name": "My Slack Alert Webhook", "url": "https://hooks.slack.com/services/Ds3fdBFbV/456464Gdd", "type": "generic", "created_at": "2021-09-15T16:33:31.834209Z", "last_success": "2021-09-15T16:33:31.834209Z", "last_failure": "2021-09-15T16:33:31.834209Z"}}`)
				}))
			})
			It(`Invoke GetWebhook successfully with retries`, func() {
				webhooksService, serviceErr := webhooksv1.NewWebhooksV1(&webhooksv1.WebhooksV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(webhooksService).ToNot(BeNil())
				webhooksService.EnableRetries(0, 0)

				// Construct an instance of the GetWebhookOptions model
				getWebhookOptionsModel := new(webhooksv1.GetWebhookOptions)
				getWebhookOptionsModel.WebhookID = core.StringPtr("testString")
				getWebhookOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := webhooksService.GetWebhookWithContext(ctx, getWebhookOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				webhooksService.DisableRetries()
				result, response, operationErr := webhooksService.GetWebhook(getWebhookOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = webhooksService.GetWebhookWithContext(ctx, getWebhookOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getWebhookPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": "6d16fcab-3e80-44b3-b59b-a3716237832e", "name": "My Slack Alert Webhook", "url": "https://hooks.slack.com/services/Ds3fdBFbV/456464Gdd", "type": "generic", "created_at": "2021-09-15T16:33:31.834209Z", "last_success": "2021-09-15T16:33:31.834209Z", "last_failure": "2021-09-15T16:33:31.834209Z"}}`)
				}))
			})
			It(`Invoke GetWebhook successfully`, func() {
				webhooksService, serviceErr := webhooksv1.NewWebhooksV1(&webhooksv1.WebhooksV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(webhooksService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := webhooksService.GetWebhook(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetWebhookOptions model
				getWebhookOptionsModel := new(webhooksv1.GetWebhookOptions)
				getWebhookOptionsModel.WebhookID = core.StringPtr("testString")
				getWebhookOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = webhooksService.GetWebhook(getWebhookOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetWebhook with error: Operation validation and request error`, func() {
				webhooksService, serviceErr := webhooksv1.NewWebhooksV1(&webhooksv1.WebhooksV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(webhooksService).ToNot(BeNil())

				// Construct an instance of the GetWebhookOptions model
				getWebhookOptionsModel := new(webhooksv1.GetWebhookOptions)
				getWebhookOptionsModel.WebhookID = core.StringPtr("testString")
				getWebhookOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := webhooksService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := webhooksService.GetWebhook(getWebhookOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetWebhookOptions model with no property values
				getWebhookOptionsModelNew := new(webhooksv1.GetWebhookOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = webhooksService.GetWebhook(getWebhookOptionsModelNew)
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
			It(`Invoke GetWebhook successfully`, func() {
				webhooksService, serviceErr := webhooksv1.NewWebhooksV1(&webhooksv1.WebhooksV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(webhooksService).ToNot(BeNil())

				// Construct an instance of the GetWebhookOptions model
				getWebhookOptionsModel := new(webhooksv1.GetWebhookOptions)
				getWebhookOptionsModel.WebhookID = core.StringPtr("testString")
				getWebhookOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := webhooksService.GetWebhook(getWebhookOptionsModel)
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
	Describe(`UpdateAlertWebhook(updateAlertWebhookOptions *UpdateAlertWebhookOptions) - Operation response error`, func() {
		crn := "testString"
		updateAlertWebhookPath := "/v1/testString/alerting/destinations/webhooks/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateAlertWebhookPath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateAlertWebhook with error: Operation response processing error`, func() {
				webhooksService, serviceErr := webhooksv1.NewWebhooksV1(&webhooksv1.WebhooksV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(webhooksService).ToNot(BeNil())

				// Construct an instance of the UpdateAlertWebhookOptions model
				updateAlertWebhookOptionsModel := new(webhooksv1.UpdateAlertWebhookOptions)
				updateAlertWebhookOptionsModel.WebhookID = core.StringPtr("testString")
				updateAlertWebhookOptionsModel.Name = core.StringPtr("My Slack Alert Webhook")
				updateAlertWebhookOptionsModel.URL = core.StringPtr("https://hooks.slack.com/services/Ds3fdBFbV/456464Gdd")
				updateAlertWebhookOptionsModel.Secret = core.StringPtr("ff1d9b80-b51d-4a06-bf67-6752fae1eb74")
				updateAlertWebhookOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := webhooksService.UpdateAlertWebhook(updateAlertWebhookOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				webhooksService.EnableRetries(0, 0)
				result, response, operationErr = webhooksService.UpdateAlertWebhook(updateAlertWebhookOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateAlertWebhook(updateAlertWebhookOptions *UpdateAlertWebhookOptions)`, func() {
		crn := "testString"
		updateAlertWebhookPath := "/v1/testString/alerting/destinations/webhooks/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateAlertWebhookPath))
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
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": "6d16fcab-3e80-44b3-b59b-a3716237832e"}}`)
				}))
			})
			It(`Invoke UpdateAlertWebhook successfully with retries`, func() {
				webhooksService, serviceErr := webhooksv1.NewWebhooksV1(&webhooksv1.WebhooksV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(webhooksService).ToNot(BeNil())
				webhooksService.EnableRetries(0, 0)

				// Construct an instance of the UpdateAlertWebhookOptions model
				updateAlertWebhookOptionsModel := new(webhooksv1.UpdateAlertWebhookOptions)
				updateAlertWebhookOptionsModel.WebhookID = core.StringPtr("testString")
				updateAlertWebhookOptionsModel.Name = core.StringPtr("My Slack Alert Webhook")
				updateAlertWebhookOptionsModel.URL = core.StringPtr("https://hooks.slack.com/services/Ds3fdBFbV/456464Gdd")
				updateAlertWebhookOptionsModel.Secret = core.StringPtr("ff1d9b80-b51d-4a06-bf67-6752fae1eb74")
				updateAlertWebhookOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := webhooksService.UpdateAlertWebhookWithContext(ctx, updateAlertWebhookOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				webhooksService.DisableRetries()
				result, response, operationErr := webhooksService.UpdateAlertWebhook(updateAlertWebhookOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = webhooksService.UpdateAlertWebhookWithContext(ctx, updateAlertWebhookOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(updateAlertWebhookPath))
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
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": "6d16fcab-3e80-44b3-b59b-a3716237832e"}}`)
				}))
			})
			It(`Invoke UpdateAlertWebhook successfully`, func() {
				webhooksService, serviceErr := webhooksv1.NewWebhooksV1(&webhooksv1.WebhooksV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(webhooksService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := webhooksService.UpdateAlertWebhook(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateAlertWebhookOptions model
				updateAlertWebhookOptionsModel := new(webhooksv1.UpdateAlertWebhookOptions)
				updateAlertWebhookOptionsModel.WebhookID = core.StringPtr("testString")
				updateAlertWebhookOptionsModel.Name = core.StringPtr("My Slack Alert Webhook")
				updateAlertWebhookOptionsModel.URL = core.StringPtr("https://hooks.slack.com/services/Ds3fdBFbV/456464Gdd")
				updateAlertWebhookOptionsModel.Secret = core.StringPtr("ff1d9b80-b51d-4a06-bf67-6752fae1eb74")
				updateAlertWebhookOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = webhooksService.UpdateAlertWebhook(updateAlertWebhookOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateAlertWebhook with error: Operation validation and request error`, func() {
				webhooksService, serviceErr := webhooksv1.NewWebhooksV1(&webhooksv1.WebhooksV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(webhooksService).ToNot(BeNil())

				// Construct an instance of the UpdateAlertWebhookOptions model
				updateAlertWebhookOptionsModel := new(webhooksv1.UpdateAlertWebhookOptions)
				updateAlertWebhookOptionsModel.WebhookID = core.StringPtr("testString")
				updateAlertWebhookOptionsModel.Name = core.StringPtr("My Slack Alert Webhook")
				updateAlertWebhookOptionsModel.URL = core.StringPtr("https://hooks.slack.com/services/Ds3fdBFbV/456464Gdd")
				updateAlertWebhookOptionsModel.Secret = core.StringPtr("ff1d9b80-b51d-4a06-bf67-6752fae1eb74")
				updateAlertWebhookOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := webhooksService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := webhooksService.UpdateAlertWebhook(updateAlertWebhookOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateAlertWebhookOptions model with no property values
				updateAlertWebhookOptionsModelNew := new(webhooksv1.UpdateAlertWebhookOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = webhooksService.UpdateAlertWebhook(updateAlertWebhookOptionsModelNew)
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
			It(`Invoke UpdateAlertWebhook successfully`, func() {
				webhooksService, serviceErr := webhooksv1.NewWebhooksV1(&webhooksv1.WebhooksV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(webhooksService).ToNot(BeNil())

				// Construct an instance of the UpdateAlertWebhookOptions model
				updateAlertWebhookOptionsModel := new(webhooksv1.UpdateAlertWebhookOptions)
				updateAlertWebhookOptionsModel.WebhookID = core.StringPtr("testString")
				updateAlertWebhookOptionsModel.Name = core.StringPtr("My Slack Alert Webhook")
				updateAlertWebhookOptionsModel.URL = core.StringPtr("https://hooks.slack.com/services/Ds3fdBFbV/456464Gdd")
				updateAlertWebhookOptionsModel.Secret = core.StringPtr("ff1d9b80-b51d-4a06-bf67-6752fae1eb74")
				updateAlertWebhookOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := webhooksService.UpdateAlertWebhook(updateAlertWebhookOptionsModel)
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
	Describe(`DeleteWebhook(deleteWebhookOptions *DeleteWebhookOptions) - Operation response error`, func() {
		crn := "testString"
		deleteWebhookPath := "/v1/testString/alerting/destinations/webhooks/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteWebhookPath))
					Expect(req.Method).To(Equal("DELETE"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke DeleteWebhook with error: Operation response processing error`, func() {
				webhooksService, serviceErr := webhooksv1.NewWebhooksV1(&webhooksv1.WebhooksV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(webhooksService).ToNot(BeNil())

				// Construct an instance of the DeleteWebhookOptions model
				deleteWebhookOptionsModel := new(webhooksv1.DeleteWebhookOptions)
				deleteWebhookOptionsModel.WebhookID = core.StringPtr("testString")
				deleteWebhookOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := webhooksService.DeleteWebhook(deleteWebhookOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				webhooksService.EnableRetries(0, 0)
				result, response, operationErr = webhooksService.DeleteWebhook(deleteWebhookOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteWebhook(deleteWebhookOptions *DeleteWebhookOptions)`, func() {
		crn := "testString"
		deleteWebhookPath := "/v1/testString/alerting/destinations/webhooks/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteWebhookPath))
					Expect(req.Method).To(Equal("DELETE"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": "6d16fcab-3e80-44b3-b59b-a3716237832e"}}`)
				}))
			})
			It(`Invoke DeleteWebhook successfully with retries`, func() {
				webhooksService, serviceErr := webhooksv1.NewWebhooksV1(&webhooksv1.WebhooksV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(webhooksService).ToNot(BeNil())
				webhooksService.EnableRetries(0, 0)

				// Construct an instance of the DeleteWebhookOptions model
				deleteWebhookOptionsModel := new(webhooksv1.DeleteWebhookOptions)
				deleteWebhookOptionsModel.WebhookID = core.StringPtr("testString")
				deleteWebhookOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := webhooksService.DeleteWebhookWithContext(ctx, deleteWebhookOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				webhooksService.DisableRetries()
				result, response, operationErr := webhooksService.DeleteWebhook(deleteWebhookOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = webhooksService.DeleteWebhookWithContext(ctx, deleteWebhookOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(deleteWebhookPath))
					Expect(req.Method).To(Equal("DELETE"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": "6d16fcab-3e80-44b3-b59b-a3716237832e"}}`)
				}))
			})
			It(`Invoke DeleteWebhook successfully`, func() {
				webhooksService, serviceErr := webhooksv1.NewWebhooksV1(&webhooksv1.WebhooksV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(webhooksService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := webhooksService.DeleteWebhook(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DeleteWebhookOptions model
				deleteWebhookOptionsModel := new(webhooksv1.DeleteWebhookOptions)
				deleteWebhookOptionsModel.WebhookID = core.StringPtr("testString")
				deleteWebhookOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = webhooksService.DeleteWebhook(deleteWebhookOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke DeleteWebhook with error: Operation validation and request error`, func() {
				webhooksService, serviceErr := webhooksv1.NewWebhooksV1(&webhooksv1.WebhooksV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(webhooksService).ToNot(BeNil())

				// Construct an instance of the DeleteWebhookOptions model
				deleteWebhookOptionsModel := new(webhooksv1.DeleteWebhookOptions)
				deleteWebhookOptionsModel.WebhookID = core.StringPtr("testString")
				deleteWebhookOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := webhooksService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := webhooksService.DeleteWebhook(deleteWebhookOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the DeleteWebhookOptions model with no property values
				deleteWebhookOptionsModelNew := new(webhooksv1.DeleteWebhookOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = webhooksService.DeleteWebhook(deleteWebhookOptionsModelNew)
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
			It(`Invoke DeleteWebhook successfully`, func() {
				webhooksService, serviceErr := webhooksv1.NewWebhooksV1(&webhooksv1.WebhooksV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(webhooksService).ToNot(BeNil())

				// Construct an instance of the DeleteWebhookOptions model
				deleteWebhookOptionsModel := new(webhooksv1.DeleteWebhookOptions)
				deleteWebhookOptionsModel.WebhookID = core.StringPtr("testString")
				deleteWebhookOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := webhooksService.DeleteWebhook(deleteWebhookOptionsModel)
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
			webhooksService, _ := webhooksv1.NewWebhooksV1(&webhooksv1.WebhooksV1Options{
				URL:           "http://webhooksv1modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
				Crn:           core.StringPtr(crn),
			})
			It(`Invoke NewCreateAlertWebhookOptions successfully`, func() {
				// Construct an instance of the CreateAlertWebhookOptions model
				createAlertWebhookOptionsModel := webhooksService.NewCreateAlertWebhookOptions()
				createAlertWebhookOptionsModel.SetName("My Slack Alert Webhook")
				createAlertWebhookOptionsModel.SetURL("https://hooks.slack.com/services/Ds3fdBFbV/456464Gdd")
				createAlertWebhookOptionsModel.SetSecret("ff1d9b80-b51d-4a06-bf67-6752fae1eb74")
				createAlertWebhookOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createAlertWebhookOptionsModel).ToNot(BeNil())
				Expect(createAlertWebhookOptionsModel.Name).To(Equal(core.StringPtr("My Slack Alert Webhook")))
				Expect(createAlertWebhookOptionsModel.URL).To(Equal(core.StringPtr("https://hooks.slack.com/services/Ds3fdBFbV/456464Gdd")))
				Expect(createAlertWebhookOptionsModel.Secret).To(Equal(core.StringPtr("ff1d9b80-b51d-4a06-bf67-6752fae1eb74")))
				Expect(createAlertWebhookOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteWebhookOptions successfully`, func() {
				// Construct an instance of the DeleteWebhookOptions model
				webhookID := "testString"
				deleteWebhookOptionsModel := webhooksService.NewDeleteWebhookOptions(webhookID)
				deleteWebhookOptionsModel.SetWebhookID("testString")
				deleteWebhookOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteWebhookOptionsModel).ToNot(BeNil())
				Expect(deleteWebhookOptionsModel.WebhookID).To(Equal(core.StringPtr("testString")))
				Expect(deleteWebhookOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetWebhookOptions successfully`, func() {
				// Construct an instance of the GetWebhookOptions model
				webhookID := "testString"
				getWebhookOptionsModel := webhooksService.NewGetWebhookOptions(webhookID)
				getWebhookOptionsModel.SetWebhookID("testString")
				getWebhookOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getWebhookOptionsModel).ToNot(BeNil())
				Expect(getWebhookOptionsModel.WebhookID).To(Equal(core.StringPtr("testString")))
				Expect(getWebhookOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListWebhooksOptions successfully`, func() {
				// Construct an instance of the ListWebhooksOptions model
				listWebhooksOptionsModel := webhooksService.NewListWebhooksOptions()
				listWebhooksOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listWebhooksOptionsModel).ToNot(BeNil())
				Expect(listWebhooksOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateAlertWebhookOptions successfully`, func() {
				// Construct an instance of the UpdateAlertWebhookOptions model
				webhookID := "testString"
				updateAlertWebhookOptionsModel := webhooksService.NewUpdateAlertWebhookOptions(webhookID)
				updateAlertWebhookOptionsModel.SetWebhookID("testString")
				updateAlertWebhookOptionsModel.SetName("My Slack Alert Webhook")
				updateAlertWebhookOptionsModel.SetURL("https://hooks.slack.com/services/Ds3fdBFbV/456464Gdd")
				updateAlertWebhookOptionsModel.SetSecret("ff1d9b80-b51d-4a06-bf67-6752fae1eb74")
				updateAlertWebhookOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateAlertWebhookOptionsModel).ToNot(BeNil())
				Expect(updateAlertWebhookOptionsModel.WebhookID).To(Equal(core.StringPtr("testString")))
				Expect(updateAlertWebhookOptionsModel.Name).To(Equal(core.StringPtr("My Slack Alert Webhook")))
				Expect(updateAlertWebhookOptionsModel.URL).To(Equal(core.StringPtr("https://hooks.slack.com/services/Ds3fdBFbV/456464Gdd")))
				Expect(updateAlertWebhookOptionsModel.Secret).To(Equal(core.StringPtr("ff1d9b80-b51d-4a06-bf67-6752fae1eb74")))
				Expect(updateAlertWebhookOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
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
