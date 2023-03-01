/**
 * (C) Copyright IBM Corp. 2023.
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

package botmanagementv1_test

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
	"github.com/IBM/networking-go-sdk/botmanagementv1"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe(`BotManagementV1`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		It(`Instantiate service client`, func() {
			botManagementService, serviceErr := botmanagementv1.NewBotManagementV1(&botmanagementv1.BotManagementV1Options{
				Authenticator:  &core.NoAuthAuthenticator{},
				Crn:            core.StringPtr(crn),
				ZoneIdentifier: core.StringPtr(zoneIdentifier),
			})
			Expect(botManagementService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			botManagementService, serviceErr := botmanagementv1.NewBotManagementV1(&botmanagementv1.BotManagementV1Options{
				URL:            "{BAD_URL_STRING",
				Crn:            core.StringPtr(crn),
				ZoneIdentifier: core.StringPtr(zoneIdentifier),
			})
			Expect(botManagementService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			botManagementService, serviceErr := botmanagementv1.NewBotManagementV1(&botmanagementv1.BotManagementV1Options{
				URL:            "https://botmanagementv1/api",
				Crn:            core.StringPtr(crn),
				ZoneIdentifier: core.StringPtr(zoneIdentifier),
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(botManagementService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Validation Error`, func() {
			botManagementService, serviceErr := botmanagementv1.NewBotManagementV1(&botmanagementv1.BotManagementV1Options{})
			Expect(botManagementService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"BOT_MANAGEMENT_URL":       "https://botmanagementv1/api",
				"BOT_MANAGEMENT_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				botManagementService, serviceErr := botmanagementv1.NewBotManagementV1UsingExternalConfig(&botmanagementv1.BotManagementV1Options{
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(botManagementService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := botManagementService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != botManagementService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(botManagementService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(botManagementService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				botManagementService, serviceErr := botmanagementv1.NewBotManagementV1UsingExternalConfig(&botmanagementv1.BotManagementV1Options{
					URL:            "https://testService/api",
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(botManagementService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(botManagementService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := botManagementService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != botManagementService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(botManagementService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(botManagementService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				botManagementService, serviceErr := botmanagementv1.NewBotManagementV1UsingExternalConfig(&botmanagementv1.BotManagementV1Options{
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				err := botManagementService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(botManagementService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(botManagementService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := botManagementService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != botManagementService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(botManagementService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(botManagementService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"BOT_MANAGEMENT_URL":       "https://botmanagementv1/api",
				"BOT_MANAGEMENT_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			botManagementService, serviceErr := botmanagementv1.NewBotManagementV1UsingExternalConfig(&botmanagementv1.BotManagementV1Options{
				Crn:            core.StringPtr(crn),
				ZoneIdentifier: core.StringPtr(zoneIdentifier),
			})

			It(`Instantiate service client with error`, func() {
				Expect(botManagementService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"BOT_MANAGEMENT_AUTH_TYPE": "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			botManagementService, serviceErr := botmanagementv1.NewBotManagementV1UsingExternalConfig(&botmanagementv1.BotManagementV1Options{
				URL:            "{BAD_URL_STRING",
				Crn:            core.StringPtr(crn),
				ZoneIdentifier: core.StringPtr(zoneIdentifier),
			})

			It(`Instantiate service client with error`, func() {
				Expect(botManagementService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = botmanagementv1.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`GetBotManagement(getBotManagementOptions *GetBotManagementOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getBotManagementPath := "/v1/testString/zones/testString/bot_management"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getBotManagementPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetBotManagement with error: Operation response processing error`, func() {
				botManagementService, serviceErr := botmanagementv1.NewBotManagementV1(&botmanagementv1.BotManagementV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(botManagementService).ToNot(BeNil())

				// Construct an instance of the GetBotManagementOptions model
				getBotManagementOptionsModel := new(botmanagementv1.GetBotManagementOptions)
				getBotManagementOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := botManagementService.GetBotManagement(getBotManagementOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				botManagementService.EnableRetries(0, 0)
				result, response, operationErr = botManagementService.GetBotManagement(getBotManagementOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetBotManagement(getBotManagementOptions *GetBotManagementOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getBotManagementPath := "/v1/testString/zones/testString/bot_management"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getBotManagementPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"fight_mode": false, "session_score": false, "enable_js": false, "auth_id_logging": false, "use_latest_model": false}}`)
				}))
			})
			It(`Invoke GetBotManagement successfully with retries`, func() {
				botManagementService, serviceErr := botmanagementv1.NewBotManagementV1(&botmanagementv1.BotManagementV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(botManagementService).ToNot(BeNil())
				botManagementService.EnableRetries(0, 0)

				// Construct an instance of the GetBotManagementOptions model
				getBotManagementOptionsModel := new(botmanagementv1.GetBotManagementOptions)
				getBotManagementOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := botManagementService.GetBotManagementWithContext(ctx, getBotManagementOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				botManagementService.DisableRetries()
				result, response, operationErr := botManagementService.GetBotManagement(getBotManagementOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = botManagementService.GetBotManagementWithContext(ctx, getBotManagementOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getBotManagementPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"fight_mode": false, "session_score": false, "enable_js": false, "auth_id_logging": false, "use_latest_model": false}}`)
				}))
			})
			It(`Invoke GetBotManagement successfully`, func() {
				botManagementService, serviceErr := botmanagementv1.NewBotManagementV1(&botmanagementv1.BotManagementV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(botManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := botManagementService.GetBotManagement(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetBotManagementOptions model
				getBotManagementOptionsModel := new(botmanagementv1.GetBotManagementOptions)
				getBotManagementOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = botManagementService.GetBotManagement(getBotManagementOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetBotManagement with error: Operation request error`, func() {
				botManagementService, serviceErr := botmanagementv1.NewBotManagementV1(&botmanagementv1.BotManagementV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(botManagementService).ToNot(BeNil())

				// Construct an instance of the GetBotManagementOptions model
				getBotManagementOptionsModel := new(botmanagementv1.GetBotManagementOptions)
				getBotManagementOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := botManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := botManagementService.GetBotManagement(getBotManagementOptionsModel)
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
			It(`Invoke GetBotManagement successfully`, func() {
				botManagementService, serviceErr := botmanagementv1.NewBotManagementV1(&botmanagementv1.BotManagementV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(botManagementService).ToNot(BeNil())

				// Construct an instance of the GetBotManagementOptions model
				getBotManagementOptionsModel := new(botmanagementv1.GetBotManagementOptions)
				getBotManagementOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := botManagementService.GetBotManagement(getBotManagementOptionsModel)
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
	Describe(`UpdateBotManagement(updateBotManagementOptions *UpdateBotManagementOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		updateBotManagementPath := "/v1/testString/zones/testString/bot_management"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateBotManagementPath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateBotManagement with error: Operation response processing error`, func() {
				botManagementService, serviceErr := botmanagementv1.NewBotManagementV1(&botmanagementv1.BotManagementV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(botManagementService).ToNot(BeNil())

				// Construct an instance of the UpdateBotManagementOptions model
				updateBotManagementOptionsModel := new(botmanagementv1.UpdateBotManagementOptions)
				updateBotManagementOptionsModel.FightMode = core.BoolPtr(false)
				updateBotManagementOptionsModel.SessionScore = core.BoolPtr(false)
				updateBotManagementOptionsModel.EnableJs = core.BoolPtr(false)
				updateBotManagementOptionsModel.AuthIdLogging = core.BoolPtr(false)
				updateBotManagementOptionsModel.UseLatestModel = core.BoolPtr(false)
				updateBotManagementOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := botManagementService.UpdateBotManagement(updateBotManagementOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				botManagementService.EnableRetries(0, 0)
				result, response, operationErr = botManagementService.UpdateBotManagement(updateBotManagementOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateBotManagement(updateBotManagementOptions *UpdateBotManagementOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		updateBotManagementPath := "/v1/testString/zones/testString/bot_management"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateBotManagementPath))
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
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"fight_mode": false, "session_score": false, "enable_js": false, "auth_id_logging": false, "use_latest_model": false}}`)
				}))
			})
			It(`Invoke UpdateBotManagement successfully with retries`, func() {
				botManagementService, serviceErr := botmanagementv1.NewBotManagementV1(&botmanagementv1.BotManagementV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(botManagementService).ToNot(BeNil())
				botManagementService.EnableRetries(0, 0)

				// Construct an instance of the UpdateBotManagementOptions model
				updateBotManagementOptionsModel := new(botmanagementv1.UpdateBotManagementOptions)
				updateBotManagementOptionsModel.FightMode = core.BoolPtr(false)
				updateBotManagementOptionsModel.SessionScore = core.BoolPtr(false)
				updateBotManagementOptionsModel.EnableJs = core.BoolPtr(false)
				updateBotManagementOptionsModel.AuthIdLogging = core.BoolPtr(false)
				updateBotManagementOptionsModel.UseLatestModel = core.BoolPtr(false)
				updateBotManagementOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := botManagementService.UpdateBotManagementWithContext(ctx, updateBotManagementOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				botManagementService.DisableRetries()
				result, response, operationErr := botManagementService.UpdateBotManagement(updateBotManagementOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = botManagementService.UpdateBotManagementWithContext(ctx, updateBotManagementOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(updateBotManagementPath))
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
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"fight_mode": false, "session_score": false, "enable_js": false, "auth_id_logging": false, "use_latest_model": false}}`)
				}))
			})
			It(`Invoke UpdateBotManagement successfully`, func() {
				botManagementService, serviceErr := botmanagementv1.NewBotManagementV1(&botmanagementv1.BotManagementV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(botManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := botManagementService.UpdateBotManagement(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateBotManagementOptions model
				updateBotManagementOptionsModel := new(botmanagementv1.UpdateBotManagementOptions)
				updateBotManagementOptionsModel.FightMode = core.BoolPtr(false)
				updateBotManagementOptionsModel.SessionScore = core.BoolPtr(false)
				updateBotManagementOptionsModel.EnableJs = core.BoolPtr(false)
				updateBotManagementOptionsModel.AuthIdLogging = core.BoolPtr(false)
				updateBotManagementOptionsModel.UseLatestModel = core.BoolPtr(false)
				updateBotManagementOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = botManagementService.UpdateBotManagement(updateBotManagementOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateBotManagement with error: Operation request error`, func() {
				botManagementService, serviceErr := botmanagementv1.NewBotManagementV1(&botmanagementv1.BotManagementV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(botManagementService).ToNot(BeNil())

				// Construct an instance of the UpdateBotManagementOptions model
				updateBotManagementOptionsModel := new(botmanagementv1.UpdateBotManagementOptions)
				updateBotManagementOptionsModel.FightMode = core.BoolPtr(false)
				updateBotManagementOptionsModel.SessionScore = core.BoolPtr(false)
				updateBotManagementOptionsModel.EnableJs = core.BoolPtr(false)
				updateBotManagementOptionsModel.AuthIdLogging = core.BoolPtr(false)
				updateBotManagementOptionsModel.UseLatestModel = core.BoolPtr(false)
				updateBotManagementOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := botManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := botManagementService.UpdateBotManagement(updateBotManagementOptionsModel)
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
			It(`Invoke UpdateBotManagement successfully`, func() {
				botManagementService, serviceErr := botmanagementv1.NewBotManagementV1(&botmanagementv1.BotManagementV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(botManagementService).ToNot(BeNil())

				// Construct an instance of the UpdateBotManagementOptions model
				updateBotManagementOptionsModel := new(botmanagementv1.UpdateBotManagementOptions)
				updateBotManagementOptionsModel.FightMode = core.BoolPtr(false)
				updateBotManagementOptionsModel.SessionScore = core.BoolPtr(false)
				updateBotManagementOptionsModel.EnableJs = core.BoolPtr(false)
				updateBotManagementOptionsModel.AuthIdLogging = core.BoolPtr(false)
				updateBotManagementOptionsModel.UseLatestModel = core.BoolPtr(false)
				updateBotManagementOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := botManagementService.UpdateBotManagement(updateBotManagementOptionsModel)
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
			zoneIdentifier := "testString"
			botManagementService, _ := botmanagementv1.NewBotManagementV1(&botmanagementv1.BotManagementV1Options{
				URL:            "http://botmanagementv1modelgenerator.com",
				Authenticator:  &core.NoAuthAuthenticator{},
				Crn:            core.StringPtr(crn),
				ZoneIdentifier: core.StringPtr(zoneIdentifier),
			})
			It(`Invoke NewGetBotManagementOptions successfully`, func() {
				// Construct an instance of the GetBotManagementOptions model
				getBotManagementOptionsModel := botManagementService.NewGetBotManagementOptions()
				getBotManagementOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getBotManagementOptionsModel).ToNot(BeNil())
				Expect(getBotManagementOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateBotManagementOptions successfully`, func() {
				// Construct an instance of the UpdateBotManagementOptions model
				updateBotManagementOptionsModel := botManagementService.NewUpdateBotManagementOptions()
				updateBotManagementOptionsModel.SetFightMode(false)
				updateBotManagementOptionsModel.SetSessionScore(false)
				updateBotManagementOptionsModel.SetEnableJs(false)
				updateBotManagementOptionsModel.SetAuthIdLogging(false)
				updateBotManagementOptionsModel.SetUseLatestModel(false)
				updateBotManagementOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateBotManagementOptionsModel).ToNot(BeNil())
				Expect(updateBotManagementOptionsModel.FightMode).To(Equal(core.BoolPtr(false)))
				Expect(updateBotManagementOptionsModel.SessionScore).To(Equal(core.BoolPtr(false)))
				Expect(updateBotManagementOptionsModel.EnableJs).To(Equal(core.BoolPtr(false)))
				Expect(updateBotManagementOptionsModel.AuthIdLogging).To(Equal(core.BoolPtr(false)))
				Expect(updateBotManagementOptionsModel.UseLatestModel).To(Equal(core.BoolPtr(false)))
				Expect(updateBotManagementOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
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
