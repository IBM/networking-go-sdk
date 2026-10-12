/**
 * (C) Copyright IBM Corp. 2026.
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

package aisecurityforappsv1_test

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
	"github.com/IBM/networking-go-sdk/aisecurityforappsv1"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe(`AiSecurityForAppsV1`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		It(`Instantiate service client`, func() {
			aiSecurityForAppsService, serviceErr := aisecurityforappsv1.NewAiSecurityForAppsV1(&aisecurityforappsv1.AiSecurityForAppsV1Options{
				Authenticator:  &core.NoAuthAuthenticator{},
				Crn:            core.StringPtr(crn),
				ZoneIdentifier: core.StringPtr(zoneIdentifier),
			})
			Expect(aiSecurityForAppsService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			aiSecurityForAppsService, serviceErr := aisecurityforappsv1.NewAiSecurityForAppsV1(&aisecurityforappsv1.AiSecurityForAppsV1Options{
				URL:            "{BAD_URL_STRING",
				Crn:            core.StringPtr(crn),
				ZoneIdentifier: core.StringPtr(zoneIdentifier),
			})
			Expect(aiSecurityForAppsService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			aiSecurityForAppsService, serviceErr := aisecurityforappsv1.NewAiSecurityForAppsV1(&aisecurityforappsv1.AiSecurityForAppsV1Options{
				URL:            "https://aisecurityforappsv1/api",
				Crn:            core.StringPtr(crn),
				ZoneIdentifier: core.StringPtr(zoneIdentifier),
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(aiSecurityForAppsService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Validation Error`, func() {
			aiSecurityForAppsService, serviceErr := aisecurityforappsv1.NewAiSecurityForAppsV1(&aisecurityforappsv1.AiSecurityForAppsV1Options{})
			Expect(aiSecurityForAppsService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"AI_SECURITY_FOR_APPS_URL":       "https://aisecurityforappsv1/api",
				"AI_SECURITY_FOR_APPS_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				aiSecurityForAppsService, serviceErr := aisecurityforappsv1.NewAiSecurityForAppsV1UsingExternalConfig(&aisecurityforappsv1.AiSecurityForAppsV1Options{
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(aiSecurityForAppsService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := aiSecurityForAppsService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != aiSecurityForAppsService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(aiSecurityForAppsService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(aiSecurityForAppsService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				aiSecurityForAppsService, serviceErr := aisecurityforappsv1.NewAiSecurityForAppsV1UsingExternalConfig(&aisecurityforappsv1.AiSecurityForAppsV1Options{
					URL:            "https://testService/api",
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(aiSecurityForAppsService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(aiSecurityForAppsService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := aiSecurityForAppsService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != aiSecurityForAppsService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(aiSecurityForAppsService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(aiSecurityForAppsService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				aiSecurityForAppsService, serviceErr := aisecurityforappsv1.NewAiSecurityForAppsV1UsingExternalConfig(&aisecurityforappsv1.AiSecurityForAppsV1Options{
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				err := aiSecurityForAppsService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(aiSecurityForAppsService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(aiSecurityForAppsService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := aiSecurityForAppsService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != aiSecurityForAppsService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(aiSecurityForAppsService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(aiSecurityForAppsService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"AI_SECURITY_FOR_APPS_URL":       "https://aisecurityforappsv1/api",
				"AI_SECURITY_FOR_APPS_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			aiSecurityForAppsService, serviceErr := aisecurityforappsv1.NewAiSecurityForAppsV1UsingExternalConfig(&aisecurityforappsv1.AiSecurityForAppsV1Options{
				Crn:            core.StringPtr(crn),
				ZoneIdentifier: core.StringPtr(zoneIdentifier),
			})

			It(`Instantiate service client with error`, func() {
				Expect(aiSecurityForAppsService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"AI_SECURITY_FOR_APPS_AUTH_TYPE": "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			aiSecurityForAppsService, serviceErr := aisecurityforappsv1.NewAiSecurityForAppsV1UsingExternalConfig(&aisecurityforappsv1.AiSecurityForAppsV1Options{
				URL:            "{BAD_URL_STRING",
				Crn:            core.StringPtr(crn),
				ZoneIdentifier: core.StringPtr(zoneIdentifier),
			})

			It(`Instantiate service client with error`, func() {
				Expect(aiSecurityForAppsService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = aisecurityforappsv1.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`GetAiSecuritySettings(getAiSecuritySettingsOptions *GetAiSecuritySettingsOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getAiSecuritySettingsPath := "/v1/testString/zones/testString/ai_security/settings"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getAiSecuritySettingsPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetAiSecuritySettings with error: Operation response processing error`, func() {
				aiSecurityForAppsService, serviceErr := aisecurityforappsv1.NewAiSecurityForAppsV1(&aisecurityforappsv1.AiSecurityForAppsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(aiSecurityForAppsService).ToNot(BeNil())

				// Construct an instance of the GetAiSecuritySettingsOptions model
				getAiSecuritySettingsOptionsModel := new(aisecurityforappsv1.GetAiSecuritySettingsOptions)
				getAiSecuritySettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := aiSecurityForAppsService.GetAiSecuritySettings(getAiSecuritySettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				aiSecurityForAppsService.EnableRetries(0, 0)
				result, response, operationErr = aiSecurityForAppsService.GetAiSecuritySettings(getAiSecuritySettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetAiSecuritySettings(getAiSecuritySettingsOptions *GetAiSecuritySettingsOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getAiSecuritySettingsPath := "/v1/testString/zones/testString/ai_security/settings"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getAiSecuritySettingsPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"enabled": false}}`)
				}))
			})
			It(`Invoke GetAiSecuritySettings successfully with retries`, func() {
				aiSecurityForAppsService, serviceErr := aisecurityforappsv1.NewAiSecurityForAppsV1(&aisecurityforappsv1.AiSecurityForAppsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(aiSecurityForAppsService).ToNot(BeNil())
				aiSecurityForAppsService.EnableRetries(0, 0)

				// Construct an instance of the GetAiSecuritySettingsOptions model
				getAiSecuritySettingsOptionsModel := new(aisecurityforappsv1.GetAiSecuritySettingsOptions)
				getAiSecuritySettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := aiSecurityForAppsService.GetAiSecuritySettingsWithContext(ctx, getAiSecuritySettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				aiSecurityForAppsService.DisableRetries()
				result, response, operationErr := aiSecurityForAppsService.GetAiSecuritySettings(getAiSecuritySettingsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = aiSecurityForAppsService.GetAiSecuritySettingsWithContext(ctx, getAiSecuritySettingsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getAiSecuritySettingsPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"enabled": false}}`)
				}))
			})
			It(`Invoke GetAiSecuritySettings successfully`, func() {
				aiSecurityForAppsService, serviceErr := aisecurityforappsv1.NewAiSecurityForAppsV1(&aisecurityforappsv1.AiSecurityForAppsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(aiSecurityForAppsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := aiSecurityForAppsService.GetAiSecuritySettings(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetAiSecuritySettingsOptions model
				getAiSecuritySettingsOptionsModel := new(aisecurityforappsv1.GetAiSecuritySettingsOptions)
				getAiSecuritySettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = aiSecurityForAppsService.GetAiSecuritySettings(getAiSecuritySettingsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetAiSecuritySettings with error: Operation request error`, func() {
				aiSecurityForAppsService, serviceErr := aisecurityforappsv1.NewAiSecurityForAppsV1(&aisecurityforappsv1.AiSecurityForAppsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(aiSecurityForAppsService).ToNot(BeNil())

				// Construct an instance of the GetAiSecuritySettingsOptions model
				getAiSecuritySettingsOptionsModel := new(aisecurityforappsv1.GetAiSecuritySettingsOptions)
				getAiSecuritySettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := aiSecurityForAppsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := aiSecurityForAppsService.GetAiSecuritySettings(getAiSecuritySettingsOptionsModel)
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
			It(`Invoke GetAiSecuritySettings successfully`, func() {
				aiSecurityForAppsService, serviceErr := aisecurityforappsv1.NewAiSecurityForAppsV1(&aisecurityforappsv1.AiSecurityForAppsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(aiSecurityForAppsService).ToNot(BeNil())

				// Construct an instance of the GetAiSecuritySettingsOptions model
				getAiSecuritySettingsOptionsModel := new(aisecurityforappsv1.GetAiSecuritySettingsOptions)
				getAiSecuritySettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := aiSecurityForAppsService.GetAiSecuritySettings(getAiSecuritySettingsOptionsModel)
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
	Describe(`ReplaceZoneAiSecuritySettings(replaceZoneAiSecuritySettingsOptions *ReplaceZoneAiSecuritySettingsOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		replaceZoneAiSecuritySettingsPath := "/v1/testString/zones/testString/ai_security/settings"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(replaceZoneAiSecuritySettingsPath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ReplaceZoneAiSecuritySettings with error: Operation response processing error`, func() {
				aiSecurityForAppsService, serviceErr := aisecurityforappsv1.NewAiSecurityForAppsV1(&aisecurityforappsv1.AiSecurityForAppsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(aiSecurityForAppsService).ToNot(BeNil())

				// Construct an instance of the ReplaceZoneAiSecuritySettingsOptions model
				replaceZoneAiSecuritySettingsOptionsModel := new(aisecurityforappsv1.ReplaceZoneAiSecuritySettingsOptions)
				replaceZoneAiSecuritySettingsOptionsModel.Enabled = core.BoolPtr(true)
				replaceZoneAiSecuritySettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := aiSecurityForAppsService.ReplaceZoneAiSecuritySettings(replaceZoneAiSecuritySettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				aiSecurityForAppsService.EnableRetries(0, 0)
				result, response, operationErr = aiSecurityForAppsService.ReplaceZoneAiSecuritySettings(replaceZoneAiSecuritySettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ReplaceZoneAiSecuritySettings(replaceZoneAiSecuritySettingsOptions *ReplaceZoneAiSecuritySettingsOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		replaceZoneAiSecuritySettingsPath := "/v1/testString/zones/testString/ai_security/settings"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(replaceZoneAiSecuritySettingsPath))
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
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"enabled": false}}`)
				}))
			})
			It(`Invoke ReplaceZoneAiSecuritySettings successfully with retries`, func() {
				aiSecurityForAppsService, serviceErr := aisecurityforappsv1.NewAiSecurityForAppsV1(&aisecurityforappsv1.AiSecurityForAppsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(aiSecurityForAppsService).ToNot(BeNil())
				aiSecurityForAppsService.EnableRetries(0, 0)

				// Construct an instance of the ReplaceZoneAiSecuritySettingsOptions model
				replaceZoneAiSecuritySettingsOptionsModel := new(aisecurityforappsv1.ReplaceZoneAiSecuritySettingsOptions)
				replaceZoneAiSecuritySettingsOptionsModel.Enabled = core.BoolPtr(true)
				replaceZoneAiSecuritySettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := aiSecurityForAppsService.ReplaceZoneAiSecuritySettingsWithContext(ctx, replaceZoneAiSecuritySettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				aiSecurityForAppsService.DisableRetries()
				result, response, operationErr := aiSecurityForAppsService.ReplaceZoneAiSecuritySettings(replaceZoneAiSecuritySettingsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = aiSecurityForAppsService.ReplaceZoneAiSecuritySettingsWithContext(ctx, replaceZoneAiSecuritySettingsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(replaceZoneAiSecuritySettingsPath))
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
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"enabled": false}}`)
				}))
			})
			It(`Invoke ReplaceZoneAiSecuritySettings successfully`, func() {
				aiSecurityForAppsService, serviceErr := aisecurityforappsv1.NewAiSecurityForAppsV1(&aisecurityforappsv1.AiSecurityForAppsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(aiSecurityForAppsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := aiSecurityForAppsService.ReplaceZoneAiSecuritySettings(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ReplaceZoneAiSecuritySettingsOptions model
				replaceZoneAiSecuritySettingsOptionsModel := new(aisecurityforappsv1.ReplaceZoneAiSecuritySettingsOptions)
				replaceZoneAiSecuritySettingsOptionsModel.Enabled = core.BoolPtr(true)
				replaceZoneAiSecuritySettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = aiSecurityForAppsService.ReplaceZoneAiSecuritySettings(replaceZoneAiSecuritySettingsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ReplaceZoneAiSecuritySettings with error: Operation request error`, func() {
				aiSecurityForAppsService, serviceErr := aisecurityforappsv1.NewAiSecurityForAppsV1(&aisecurityforappsv1.AiSecurityForAppsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(aiSecurityForAppsService).ToNot(BeNil())

				// Construct an instance of the ReplaceZoneAiSecuritySettingsOptions model
				replaceZoneAiSecuritySettingsOptionsModel := new(aisecurityforappsv1.ReplaceZoneAiSecuritySettingsOptions)
				replaceZoneAiSecuritySettingsOptionsModel.Enabled = core.BoolPtr(true)
				replaceZoneAiSecuritySettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := aiSecurityForAppsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := aiSecurityForAppsService.ReplaceZoneAiSecuritySettings(replaceZoneAiSecuritySettingsOptionsModel)
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
			It(`Invoke ReplaceZoneAiSecuritySettings successfully`, func() {
				aiSecurityForAppsService, serviceErr := aisecurityforappsv1.NewAiSecurityForAppsV1(&aisecurityforappsv1.AiSecurityForAppsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(aiSecurityForAppsService).ToNot(BeNil())

				// Construct an instance of the ReplaceZoneAiSecuritySettingsOptions model
				replaceZoneAiSecuritySettingsOptionsModel := new(aisecurityforappsv1.ReplaceZoneAiSecuritySettingsOptions)
				replaceZoneAiSecuritySettingsOptionsModel.Enabled = core.BoolPtr(true)
				replaceZoneAiSecuritySettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := aiSecurityForAppsService.ReplaceZoneAiSecuritySettings(replaceZoneAiSecuritySettingsOptionsModel)
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
	Describe(`GetApiGatewayDiscovery(getApiGatewayDiscoveryOptions *GetApiGatewayDiscoveryOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getApiGatewayDiscoveryPath := "/v1/testString/zones/testString/api_gateway/discovery"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getApiGatewayDiscoveryPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetApiGatewayDiscovery with error: Operation response processing error`, func() {
				aiSecurityForAppsService, serviceErr := aisecurityforappsv1.NewAiSecurityForAppsV1(&aisecurityforappsv1.AiSecurityForAppsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(aiSecurityForAppsService).ToNot(BeNil())

				// Construct an instance of the GetApiGatewayDiscoveryOptions model
				getApiGatewayDiscoveryOptionsModel := new(aisecurityforappsv1.GetApiGatewayDiscoveryOptions)
				getApiGatewayDiscoveryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := aiSecurityForAppsService.GetApiGatewayDiscovery(getApiGatewayDiscoveryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				aiSecurityForAppsService.EnableRetries(0, 0)
				result, response, operationErr = aiSecurityForAppsService.GetApiGatewayDiscovery(getApiGatewayDiscoveryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetApiGatewayDiscovery(getApiGatewayDiscoveryOptions *GetApiGatewayDiscoveryOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getApiGatewayDiscoveryPath := "/v1/testString/zones/testString/api_gateway/discovery"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getApiGatewayDiscoveryPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"anyKey": "anyValue"}}`)
				}))
			})
			It(`Invoke GetApiGatewayDiscovery successfully with retries`, func() {
				aiSecurityForAppsService, serviceErr := aisecurityforappsv1.NewAiSecurityForAppsV1(&aisecurityforappsv1.AiSecurityForAppsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(aiSecurityForAppsService).ToNot(BeNil())
				aiSecurityForAppsService.EnableRetries(0, 0)

				// Construct an instance of the GetApiGatewayDiscoveryOptions model
				getApiGatewayDiscoveryOptionsModel := new(aisecurityforappsv1.GetApiGatewayDiscoveryOptions)
				getApiGatewayDiscoveryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := aiSecurityForAppsService.GetApiGatewayDiscoveryWithContext(ctx, getApiGatewayDiscoveryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				aiSecurityForAppsService.DisableRetries()
				result, response, operationErr := aiSecurityForAppsService.GetApiGatewayDiscovery(getApiGatewayDiscoveryOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = aiSecurityForAppsService.GetApiGatewayDiscoveryWithContext(ctx, getApiGatewayDiscoveryOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getApiGatewayDiscoveryPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"anyKey": "anyValue"}}`)
				}))
			})
			It(`Invoke GetApiGatewayDiscovery successfully`, func() {
				aiSecurityForAppsService, serviceErr := aisecurityforappsv1.NewAiSecurityForAppsV1(&aisecurityforappsv1.AiSecurityForAppsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(aiSecurityForAppsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := aiSecurityForAppsService.GetApiGatewayDiscovery(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetApiGatewayDiscoveryOptions model
				getApiGatewayDiscoveryOptionsModel := new(aisecurityforappsv1.GetApiGatewayDiscoveryOptions)
				getApiGatewayDiscoveryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = aiSecurityForAppsService.GetApiGatewayDiscovery(getApiGatewayDiscoveryOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetApiGatewayDiscovery with error: Operation request error`, func() {
				aiSecurityForAppsService, serviceErr := aisecurityforappsv1.NewAiSecurityForAppsV1(&aisecurityforappsv1.AiSecurityForAppsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(aiSecurityForAppsService).ToNot(BeNil())

				// Construct an instance of the GetApiGatewayDiscoveryOptions model
				getApiGatewayDiscoveryOptionsModel := new(aisecurityforappsv1.GetApiGatewayDiscoveryOptions)
				getApiGatewayDiscoveryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := aiSecurityForAppsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := aiSecurityForAppsService.GetApiGatewayDiscovery(getApiGatewayDiscoveryOptionsModel)
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
			It(`Invoke GetApiGatewayDiscovery successfully`, func() {
				aiSecurityForAppsService, serviceErr := aisecurityforappsv1.NewAiSecurityForAppsV1(&aisecurityforappsv1.AiSecurityForAppsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(aiSecurityForAppsService).ToNot(BeNil())

				// Construct an instance of the GetApiGatewayDiscoveryOptions model
				getApiGatewayDiscoveryOptionsModel := new(aisecurityforappsv1.GetApiGatewayDiscoveryOptions)
				getApiGatewayDiscoveryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := aiSecurityForAppsService.GetApiGatewayDiscovery(getApiGatewayDiscoveryOptionsModel)
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
	Describe(`ListApiGatewayDiscoveryOperations(listApiGatewayDiscoveryOperationsOptions *ListApiGatewayDiscoveryOperationsOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		listApiGatewayDiscoveryOperationsPath := "/v1/testString/zones/testString/api_gateway/discovery/operations"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listApiGatewayDiscoveryOperationsPath))
					Expect(req.Method).To(Equal("GET"))
					// TODO: Add check for diff query parameter
					Expect(req.URL.Query()["direction"]).To(Equal([]string{"asc"}))
					Expect(req.URL.Query()["endpoint"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["order"]).To(Equal([]string{"host"}))
					Expect(req.URL.Query()["origin"]).To(Equal([]string{"ML"}))
					Expect(req.URL.Query()["state"]).To(Equal([]string{"review"}))
					Expect(req.URL.Query()["page"]).To(Equal([]string{fmt.Sprint(int64(1))}))
					Expect(req.URL.Query()["per_page"]).To(Equal([]string{fmt.Sprint(int64(5))}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListApiGatewayDiscoveryOperations with error: Operation response processing error`, func() {
				aiSecurityForAppsService, serviceErr := aisecurityforappsv1.NewAiSecurityForAppsV1(&aisecurityforappsv1.AiSecurityForAppsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(aiSecurityForAppsService).ToNot(BeNil())

				// Construct an instance of the ListApiGatewayDiscoveryOperationsOptions model
				listApiGatewayDiscoveryOperationsOptionsModel := new(aisecurityforappsv1.ListApiGatewayDiscoveryOperationsOptions)
				listApiGatewayDiscoveryOperationsOptionsModel.Diff = core.BoolPtr(true)
				listApiGatewayDiscoveryOperationsOptionsModel.Direction = core.StringPtr("asc")
				listApiGatewayDiscoveryOperationsOptionsModel.Endpoint = core.StringPtr("testString")
				listApiGatewayDiscoveryOperationsOptionsModel.Host = []string{"testString"}
				listApiGatewayDiscoveryOperationsOptionsModel.Method = []string{"testString"}
				listApiGatewayDiscoveryOperationsOptionsModel.Order = core.StringPtr("host")
				listApiGatewayDiscoveryOperationsOptionsModel.Origin = core.StringPtr("ML")
				listApiGatewayDiscoveryOperationsOptionsModel.State = core.StringPtr("review")
				listApiGatewayDiscoveryOperationsOptionsModel.Page = core.Int64Ptr(int64(1))
				listApiGatewayDiscoveryOperationsOptionsModel.PerPage = core.Int64Ptr(int64(5))
				listApiGatewayDiscoveryOperationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := aiSecurityForAppsService.ListApiGatewayDiscoveryOperations(listApiGatewayDiscoveryOperationsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				aiSecurityForAppsService.EnableRetries(0, 0)
				result, response, operationErr = aiSecurityForAppsService.ListApiGatewayDiscoveryOperations(listApiGatewayDiscoveryOperationsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListApiGatewayDiscoveryOperations(listApiGatewayDiscoveryOperationsOptions *ListApiGatewayDiscoveryOperationsOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		listApiGatewayDiscoveryOperationsPath := "/v1/testString/zones/testString/api_gateway/discovery/operations"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listApiGatewayDiscoveryOperationsPath))
					Expect(req.Method).To(Equal("GET"))

					// TODO: Add check for diff query parameter
					Expect(req.URL.Query()["direction"]).To(Equal([]string{"asc"}))
					Expect(req.URL.Query()["endpoint"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["order"]).To(Equal([]string{"host"}))
					Expect(req.URL.Query()["origin"]).To(Equal([]string{"ML"}))
					Expect(req.URL.Query()["state"]).To(Equal([]string{"review"}))
					Expect(req.URL.Query()["page"]).To(Equal([]string{fmt.Sprint(int64(1))}))
					Expect(req.URL.Query()["per_page"]).To(Equal([]string{fmt.Sprint(int64(5))}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": [{"id": "f174e90a-fafe-4643-bbbc-4a0ed4fc8415", "endpoint": "/v1/messages", "host": "api.example.com", "method": "POST", "last_updated": "2024-01-01T00:00:00.000Z", "origin": ["ML"], "state": "review", "features": {"traffic_stats": {"last_updated": "2019-01-01T12:00:00.000Z", "period_seconds": 13, "requests": 8}}}], "result_info": {"count": 5, "page": 4, "per_page": 7, "total_count": 10}}`)
				}))
			})
			It(`Invoke ListApiGatewayDiscoveryOperations successfully with retries`, func() {
				aiSecurityForAppsService, serviceErr := aisecurityforappsv1.NewAiSecurityForAppsV1(&aisecurityforappsv1.AiSecurityForAppsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(aiSecurityForAppsService).ToNot(BeNil())
				aiSecurityForAppsService.EnableRetries(0, 0)

				// Construct an instance of the ListApiGatewayDiscoveryOperationsOptions model
				listApiGatewayDiscoveryOperationsOptionsModel := new(aisecurityforappsv1.ListApiGatewayDiscoveryOperationsOptions)
				listApiGatewayDiscoveryOperationsOptionsModel.Diff = core.BoolPtr(true)
				listApiGatewayDiscoveryOperationsOptionsModel.Direction = core.StringPtr("asc")
				listApiGatewayDiscoveryOperationsOptionsModel.Endpoint = core.StringPtr("testString")
				listApiGatewayDiscoveryOperationsOptionsModel.Host = []string{"testString"}
				listApiGatewayDiscoveryOperationsOptionsModel.Method = []string{"testString"}
				listApiGatewayDiscoveryOperationsOptionsModel.Order = core.StringPtr("host")
				listApiGatewayDiscoveryOperationsOptionsModel.Origin = core.StringPtr("ML")
				listApiGatewayDiscoveryOperationsOptionsModel.State = core.StringPtr("review")
				listApiGatewayDiscoveryOperationsOptionsModel.Page = core.Int64Ptr(int64(1))
				listApiGatewayDiscoveryOperationsOptionsModel.PerPage = core.Int64Ptr(int64(5))
				listApiGatewayDiscoveryOperationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := aiSecurityForAppsService.ListApiGatewayDiscoveryOperationsWithContext(ctx, listApiGatewayDiscoveryOperationsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				aiSecurityForAppsService.DisableRetries()
				result, response, operationErr := aiSecurityForAppsService.ListApiGatewayDiscoveryOperations(listApiGatewayDiscoveryOperationsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = aiSecurityForAppsService.ListApiGatewayDiscoveryOperationsWithContext(ctx, listApiGatewayDiscoveryOperationsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listApiGatewayDiscoveryOperationsPath))
					Expect(req.Method).To(Equal("GET"))

					// TODO: Add check for diff query parameter
					Expect(req.URL.Query()["direction"]).To(Equal([]string{"asc"}))
					Expect(req.URL.Query()["endpoint"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["order"]).To(Equal([]string{"host"}))
					Expect(req.URL.Query()["origin"]).To(Equal([]string{"ML"}))
					Expect(req.URL.Query()["state"]).To(Equal([]string{"review"}))
					Expect(req.URL.Query()["page"]).To(Equal([]string{fmt.Sprint(int64(1))}))
					Expect(req.URL.Query()["per_page"]).To(Equal([]string{fmt.Sprint(int64(5))}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": [{"id": "f174e90a-fafe-4643-bbbc-4a0ed4fc8415", "endpoint": "/v1/messages", "host": "api.example.com", "method": "POST", "last_updated": "2024-01-01T00:00:00.000Z", "origin": ["ML"], "state": "review", "features": {"traffic_stats": {"last_updated": "2019-01-01T12:00:00.000Z", "period_seconds": 13, "requests": 8}}}], "result_info": {"count": 5, "page": 4, "per_page": 7, "total_count": 10}}`)
				}))
			})
			It(`Invoke ListApiGatewayDiscoveryOperations successfully`, func() {
				aiSecurityForAppsService, serviceErr := aisecurityforappsv1.NewAiSecurityForAppsV1(&aisecurityforappsv1.AiSecurityForAppsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(aiSecurityForAppsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := aiSecurityForAppsService.ListApiGatewayDiscoveryOperations(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListApiGatewayDiscoveryOperationsOptions model
				listApiGatewayDiscoveryOperationsOptionsModel := new(aisecurityforappsv1.ListApiGatewayDiscoveryOperationsOptions)
				listApiGatewayDiscoveryOperationsOptionsModel.Diff = core.BoolPtr(true)
				listApiGatewayDiscoveryOperationsOptionsModel.Direction = core.StringPtr("asc")
				listApiGatewayDiscoveryOperationsOptionsModel.Endpoint = core.StringPtr("testString")
				listApiGatewayDiscoveryOperationsOptionsModel.Host = []string{"testString"}
				listApiGatewayDiscoveryOperationsOptionsModel.Method = []string{"testString"}
				listApiGatewayDiscoveryOperationsOptionsModel.Order = core.StringPtr("host")
				listApiGatewayDiscoveryOperationsOptionsModel.Origin = core.StringPtr("ML")
				listApiGatewayDiscoveryOperationsOptionsModel.State = core.StringPtr("review")
				listApiGatewayDiscoveryOperationsOptionsModel.Page = core.Int64Ptr(int64(1))
				listApiGatewayDiscoveryOperationsOptionsModel.PerPage = core.Int64Ptr(int64(5))
				listApiGatewayDiscoveryOperationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = aiSecurityForAppsService.ListApiGatewayDiscoveryOperations(listApiGatewayDiscoveryOperationsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListApiGatewayDiscoveryOperations with error: Operation request error`, func() {
				aiSecurityForAppsService, serviceErr := aisecurityforappsv1.NewAiSecurityForAppsV1(&aisecurityforappsv1.AiSecurityForAppsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(aiSecurityForAppsService).ToNot(BeNil())

				// Construct an instance of the ListApiGatewayDiscoveryOperationsOptions model
				listApiGatewayDiscoveryOperationsOptionsModel := new(aisecurityforappsv1.ListApiGatewayDiscoveryOperationsOptions)
				listApiGatewayDiscoveryOperationsOptionsModel.Diff = core.BoolPtr(true)
				listApiGatewayDiscoveryOperationsOptionsModel.Direction = core.StringPtr("asc")
				listApiGatewayDiscoveryOperationsOptionsModel.Endpoint = core.StringPtr("testString")
				listApiGatewayDiscoveryOperationsOptionsModel.Host = []string{"testString"}
				listApiGatewayDiscoveryOperationsOptionsModel.Method = []string{"testString"}
				listApiGatewayDiscoveryOperationsOptionsModel.Order = core.StringPtr("host")
				listApiGatewayDiscoveryOperationsOptionsModel.Origin = core.StringPtr("ML")
				listApiGatewayDiscoveryOperationsOptionsModel.State = core.StringPtr("review")
				listApiGatewayDiscoveryOperationsOptionsModel.Page = core.Int64Ptr(int64(1))
				listApiGatewayDiscoveryOperationsOptionsModel.PerPage = core.Int64Ptr(int64(5))
				listApiGatewayDiscoveryOperationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := aiSecurityForAppsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := aiSecurityForAppsService.ListApiGatewayDiscoveryOperations(listApiGatewayDiscoveryOperationsOptionsModel)
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
			It(`Invoke ListApiGatewayDiscoveryOperations successfully`, func() {
				aiSecurityForAppsService, serviceErr := aisecurityforappsv1.NewAiSecurityForAppsV1(&aisecurityforappsv1.AiSecurityForAppsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(aiSecurityForAppsService).ToNot(BeNil())

				// Construct an instance of the ListApiGatewayDiscoveryOperationsOptions model
				listApiGatewayDiscoveryOperationsOptionsModel := new(aisecurityforappsv1.ListApiGatewayDiscoveryOperationsOptions)
				listApiGatewayDiscoveryOperationsOptionsModel.Diff = core.BoolPtr(true)
				listApiGatewayDiscoveryOperationsOptionsModel.Direction = core.StringPtr("asc")
				listApiGatewayDiscoveryOperationsOptionsModel.Endpoint = core.StringPtr("testString")
				listApiGatewayDiscoveryOperationsOptionsModel.Host = []string{"testString"}
				listApiGatewayDiscoveryOperationsOptionsModel.Method = []string{"testString"}
				listApiGatewayDiscoveryOperationsOptionsModel.Order = core.StringPtr("host")
				listApiGatewayDiscoveryOperationsOptionsModel.Origin = core.StringPtr("ML")
				listApiGatewayDiscoveryOperationsOptionsModel.State = core.StringPtr("review")
				listApiGatewayDiscoveryOperationsOptionsModel.Page = core.Int64Ptr(int64(1))
				listApiGatewayDiscoveryOperationsOptionsModel.PerPage = core.Int64Ptr(int64(5))
				listApiGatewayDiscoveryOperationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := aiSecurityForAppsService.ListApiGatewayDiscoveryOperations(listApiGatewayDiscoveryOperationsOptionsModel)
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
	Describe(`UpdateZoneApiGatewayDiscoveryOperation(updateZoneApiGatewayDiscoveryOperationOptions *UpdateZoneApiGatewayDiscoveryOperationOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		updateZoneApiGatewayDiscoveryOperationPath := "/v1/testString/zones/testString/api_gateway/discovery/operations"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateZoneApiGatewayDiscoveryOperationPath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateZoneApiGatewayDiscoveryOperation with error: Operation response processing error`, func() {
				aiSecurityForAppsService, serviceErr := aisecurityforappsv1.NewAiSecurityForAppsV1(&aisecurityforappsv1.AiSecurityForAppsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(aiSecurityForAppsService).ToNot(BeNil())

				// Construct an instance of the UpdateZoneApiGatewayDiscoveryOperationOptions model
				updateZoneApiGatewayDiscoveryOperationOptionsModel := new(aisecurityforappsv1.UpdateZoneApiGatewayDiscoveryOperationOptions)
				updateZoneApiGatewayDiscoveryOperationOptionsModel.RequestBody = map[string]interface{}{"anyKey": "anyValue"}
				updateZoneApiGatewayDiscoveryOperationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := aiSecurityForAppsService.UpdateZoneApiGatewayDiscoveryOperation(updateZoneApiGatewayDiscoveryOperationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				aiSecurityForAppsService.EnableRetries(0, 0)
				result, response, operationErr = aiSecurityForAppsService.UpdateZoneApiGatewayDiscoveryOperation(updateZoneApiGatewayDiscoveryOperationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateZoneApiGatewayDiscoveryOperation(updateZoneApiGatewayDiscoveryOperationOptions *UpdateZoneApiGatewayDiscoveryOperationOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		updateZoneApiGatewayDiscoveryOperationPath := "/v1/testString/zones/testString/api_gateway/discovery/operations"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateZoneApiGatewayDiscoveryOperationPath))
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
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": [{"id": "f174e90a-fafe-4643-bbbc-4a0ed4fc8415", "endpoint": "/v1/messages", "host": "api.example.com", "method": "POST", "last_updated": "2024-01-01T00:00:00.000Z", "origin": ["ML"], "state": "review", "features": {"traffic_stats": {"last_updated": "2019-01-01T12:00:00.000Z", "period_seconds": 13, "requests": 8}}}]}`)
				}))
			})
			It(`Invoke UpdateZoneApiGatewayDiscoveryOperation successfully with retries`, func() {
				aiSecurityForAppsService, serviceErr := aisecurityforappsv1.NewAiSecurityForAppsV1(&aisecurityforappsv1.AiSecurityForAppsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(aiSecurityForAppsService).ToNot(BeNil())
				aiSecurityForAppsService.EnableRetries(0, 0)

				// Construct an instance of the UpdateZoneApiGatewayDiscoveryOperationOptions model
				updateZoneApiGatewayDiscoveryOperationOptionsModel := new(aisecurityforappsv1.UpdateZoneApiGatewayDiscoveryOperationOptions)
				updateZoneApiGatewayDiscoveryOperationOptionsModel.RequestBody = map[string]interface{}{"anyKey": "anyValue"}
				updateZoneApiGatewayDiscoveryOperationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := aiSecurityForAppsService.UpdateZoneApiGatewayDiscoveryOperationWithContext(ctx, updateZoneApiGatewayDiscoveryOperationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				aiSecurityForAppsService.DisableRetries()
				result, response, operationErr := aiSecurityForAppsService.UpdateZoneApiGatewayDiscoveryOperation(updateZoneApiGatewayDiscoveryOperationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = aiSecurityForAppsService.UpdateZoneApiGatewayDiscoveryOperationWithContext(ctx, updateZoneApiGatewayDiscoveryOperationOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(updateZoneApiGatewayDiscoveryOperationPath))
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
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": [{"id": "f174e90a-fafe-4643-bbbc-4a0ed4fc8415", "endpoint": "/v1/messages", "host": "api.example.com", "method": "POST", "last_updated": "2024-01-01T00:00:00.000Z", "origin": ["ML"], "state": "review", "features": {"traffic_stats": {"last_updated": "2019-01-01T12:00:00.000Z", "period_seconds": 13, "requests": 8}}}]}`)
				}))
			})
			It(`Invoke UpdateZoneApiGatewayDiscoveryOperation successfully`, func() {
				aiSecurityForAppsService, serviceErr := aisecurityforappsv1.NewAiSecurityForAppsV1(&aisecurityforappsv1.AiSecurityForAppsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(aiSecurityForAppsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := aiSecurityForAppsService.UpdateZoneApiGatewayDiscoveryOperation(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateZoneApiGatewayDiscoveryOperationOptions model
				updateZoneApiGatewayDiscoveryOperationOptionsModel := new(aisecurityforappsv1.UpdateZoneApiGatewayDiscoveryOperationOptions)
				updateZoneApiGatewayDiscoveryOperationOptionsModel.RequestBody = map[string]interface{}{"anyKey": "anyValue"}
				updateZoneApiGatewayDiscoveryOperationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = aiSecurityForAppsService.UpdateZoneApiGatewayDiscoveryOperation(updateZoneApiGatewayDiscoveryOperationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateZoneApiGatewayDiscoveryOperation with error: Operation request error`, func() {
				aiSecurityForAppsService, serviceErr := aisecurityforappsv1.NewAiSecurityForAppsV1(&aisecurityforappsv1.AiSecurityForAppsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(aiSecurityForAppsService).ToNot(BeNil())

				// Construct an instance of the UpdateZoneApiGatewayDiscoveryOperationOptions model
				updateZoneApiGatewayDiscoveryOperationOptionsModel := new(aisecurityforappsv1.UpdateZoneApiGatewayDiscoveryOperationOptions)
				updateZoneApiGatewayDiscoveryOperationOptionsModel.RequestBody = map[string]interface{}{"anyKey": "anyValue"}
				updateZoneApiGatewayDiscoveryOperationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := aiSecurityForAppsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := aiSecurityForAppsService.UpdateZoneApiGatewayDiscoveryOperation(updateZoneApiGatewayDiscoveryOperationOptionsModel)
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
			It(`Invoke UpdateZoneApiGatewayDiscoveryOperation successfully`, func() {
				aiSecurityForAppsService, serviceErr := aisecurityforappsv1.NewAiSecurityForAppsV1(&aisecurityforappsv1.AiSecurityForAppsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(aiSecurityForAppsService).ToNot(BeNil())

				// Construct an instance of the UpdateZoneApiGatewayDiscoveryOperationOptions model
				updateZoneApiGatewayDiscoveryOperationOptionsModel := new(aisecurityforappsv1.UpdateZoneApiGatewayDiscoveryOperationOptions)
				updateZoneApiGatewayDiscoveryOperationOptionsModel.RequestBody = map[string]interface{}{"anyKey": "anyValue"}
				updateZoneApiGatewayDiscoveryOperationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := aiSecurityForAppsService.UpdateZoneApiGatewayDiscoveryOperation(updateZoneApiGatewayDiscoveryOperationOptionsModel)
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
	Describe(`DeleteZoneApiGatewayOperations(deleteZoneApiGatewayOperationsOptions *DeleteZoneApiGatewayOperationsOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		deleteZoneApiGatewayOperationsPath := "/v1/testString/zones/testString/api_gateway/operations"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteZoneApiGatewayOperationsPath))
					Expect(req.Method).To(Equal("DELETE"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke DeleteZoneApiGatewayOperations with error: Operation response processing error`, func() {
				aiSecurityForAppsService, serviceErr := aisecurityforappsv1.NewAiSecurityForAppsV1(&aisecurityforappsv1.AiSecurityForAppsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(aiSecurityForAppsService).ToNot(BeNil())

				// Construct an instance of the DeleteZoneApiGatewayOperationsOptions model
				deleteZoneApiGatewayOperationsOptionsModel := new(aisecurityforappsv1.DeleteZoneApiGatewayOperationsOptions)
				deleteZoneApiGatewayOperationsOptionsModel.OperationIds = []string{"f174e90a-fafe-4643-bbbc-4a0ed4fc8415"}
				deleteZoneApiGatewayOperationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := aiSecurityForAppsService.DeleteZoneApiGatewayOperations(deleteZoneApiGatewayOperationsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				aiSecurityForAppsService.EnableRetries(0, 0)
				result, response, operationErr = aiSecurityForAppsService.DeleteZoneApiGatewayOperations(deleteZoneApiGatewayOperationsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteZoneApiGatewayOperations(deleteZoneApiGatewayOperationsOptions *DeleteZoneApiGatewayOperationsOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		deleteZoneApiGatewayOperationsPath := "/v1/testString/zones/testString/api_gateway/operations"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteZoneApiGatewayOperationsPath))
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
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": "anyValue"}`)
				}))
			})
			It(`Invoke DeleteZoneApiGatewayOperations successfully with retries`, func() {
				aiSecurityForAppsService, serviceErr := aisecurityforappsv1.NewAiSecurityForAppsV1(&aisecurityforappsv1.AiSecurityForAppsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(aiSecurityForAppsService).ToNot(BeNil())
				aiSecurityForAppsService.EnableRetries(0, 0)

				// Construct an instance of the DeleteZoneApiGatewayOperationsOptions model
				deleteZoneApiGatewayOperationsOptionsModel := new(aisecurityforappsv1.DeleteZoneApiGatewayOperationsOptions)
				deleteZoneApiGatewayOperationsOptionsModel.OperationIds = []string{"f174e90a-fafe-4643-bbbc-4a0ed4fc8415"}
				deleteZoneApiGatewayOperationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := aiSecurityForAppsService.DeleteZoneApiGatewayOperationsWithContext(ctx, deleteZoneApiGatewayOperationsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				aiSecurityForAppsService.DisableRetries()
				result, response, operationErr := aiSecurityForAppsService.DeleteZoneApiGatewayOperations(deleteZoneApiGatewayOperationsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = aiSecurityForAppsService.DeleteZoneApiGatewayOperationsWithContext(ctx, deleteZoneApiGatewayOperationsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(deleteZoneApiGatewayOperationsPath))
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
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": "anyValue"}`)
				}))
			})
			It(`Invoke DeleteZoneApiGatewayOperations successfully`, func() {
				aiSecurityForAppsService, serviceErr := aisecurityforappsv1.NewAiSecurityForAppsV1(&aisecurityforappsv1.AiSecurityForAppsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(aiSecurityForAppsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := aiSecurityForAppsService.DeleteZoneApiGatewayOperations(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DeleteZoneApiGatewayOperationsOptions model
				deleteZoneApiGatewayOperationsOptionsModel := new(aisecurityforappsv1.DeleteZoneApiGatewayOperationsOptions)
				deleteZoneApiGatewayOperationsOptionsModel.OperationIds = []string{"f174e90a-fafe-4643-bbbc-4a0ed4fc8415"}
				deleteZoneApiGatewayOperationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = aiSecurityForAppsService.DeleteZoneApiGatewayOperations(deleteZoneApiGatewayOperationsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke DeleteZoneApiGatewayOperations with error: Operation request error`, func() {
				aiSecurityForAppsService, serviceErr := aisecurityforappsv1.NewAiSecurityForAppsV1(&aisecurityforappsv1.AiSecurityForAppsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(aiSecurityForAppsService).ToNot(BeNil())

				// Construct an instance of the DeleteZoneApiGatewayOperationsOptions model
				deleteZoneApiGatewayOperationsOptionsModel := new(aisecurityforappsv1.DeleteZoneApiGatewayOperationsOptions)
				deleteZoneApiGatewayOperationsOptionsModel.OperationIds = []string{"f174e90a-fafe-4643-bbbc-4a0ed4fc8415"}
				deleteZoneApiGatewayOperationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := aiSecurityForAppsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := aiSecurityForAppsService.DeleteZoneApiGatewayOperations(deleteZoneApiGatewayOperationsOptionsModel)
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
			It(`Invoke DeleteZoneApiGatewayOperations successfully`, func() {
				aiSecurityForAppsService, serviceErr := aisecurityforappsv1.NewAiSecurityForAppsV1(&aisecurityforappsv1.AiSecurityForAppsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(aiSecurityForAppsService).ToNot(BeNil())

				// Construct an instance of the DeleteZoneApiGatewayOperationsOptions model
				deleteZoneApiGatewayOperationsOptionsModel := new(aisecurityforappsv1.DeleteZoneApiGatewayOperationsOptions)
				deleteZoneApiGatewayOperationsOptionsModel.OperationIds = []string{"f174e90a-fafe-4643-bbbc-4a0ed4fc8415"}
				deleteZoneApiGatewayOperationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := aiSecurityForAppsService.DeleteZoneApiGatewayOperations(deleteZoneApiGatewayOperationsOptionsModel)
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
	Describe(`ListZoneApiGatewayOperations(listZoneApiGatewayOperationsOptions *ListZoneApiGatewayOperationsOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		listZoneApiGatewayOperationsPath := "/v1/testString/zones/testString/api_gateway/operations"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listZoneApiGatewayOperationsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["endpoint"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["direction"]).To(Equal([]string{"asc"}))
					Expect(req.URL.Query()["order"]).To(Equal([]string{"host"}))
					Expect(req.URL.Query()["page"]).To(Equal([]string{fmt.Sprint(int64(1))}))
					Expect(req.URL.Query()["per_page"]).To(Equal([]string{fmt.Sprint(int64(1))}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListZoneApiGatewayOperations with error: Operation response processing error`, func() {
				aiSecurityForAppsService, serviceErr := aisecurityforappsv1.NewAiSecurityForAppsV1(&aisecurityforappsv1.AiSecurityForAppsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(aiSecurityForAppsService).ToNot(BeNil())

				// Construct an instance of the ListZoneApiGatewayOperationsOptions model
				listZoneApiGatewayOperationsOptionsModel := new(aisecurityforappsv1.ListZoneApiGatewayOperationsOptions)
				listZoneApiGatewayOperationsOptionsModel.Host = []string{"testString"}
				listZoneApiGatewayOperationsOptionsModel.Method = []string{"GET"}
				listZoneApiGatewayOperationsOptionsModel.Endpoint = core.StringPtr("testString")
				listZoneApiGatewayOperationsOptionsModel.Labels = []string{"testString"}
				listZoneApiGatewayOperationsOptionsModel.Direction = core.StringPtr("asc")
				listZoneApiGatewayOperationsOptionsModel.Order = core.StringPtr("host")
				listZoneApiGatewayOperationsOptionsModel.Page = core.Int64Ptr(int64(1))
				listZoneApiGatewayOperationsOptionsModel.PerPage = core.Int64Ptr(int64(1))
				listZoneApiGatewayOperationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := aiSecurityForAppsService.ListZoneApiGatewayOperations(listZoneApiGatewayOperationsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				aiSecurityForAppsService.EnableRetries(0, 0)
				result, response, operationErr = aiSecurityForAppsService.ListZoneApiGatewayOperations(listZoneApiGatewayOperationsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListZoneApiGatewayOperations(listZoneApiGatewayOperationsOptions *ListZoneApiGatewayOperationsOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		listZoneApiGatewayOperationsPath := "/v1/testString/zones/testString/api_gateway/operations"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listZoneApiGatewayOperationsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["endpoint"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["direction"]).To(Equal([]string{"asc"}))
					Expect(req.URL.Query()["order"]).To(Equal([]string{"host"}))
					Expect(req.URL.Query()["page"]).To(Equal([]string{fmt.Sprint(int64(1))}))
					Expect(req.URL.Query()["per_page"]).To(Equal([]string{fmt.Sprint(int64(1))}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": [{"operation_id": "test-fafe-4643-bbbc-example", "method": "POST", "host": "api.example.com", "endpoint": "/v1/messages", "last_updated": "2024-01-01T00:00:00.000Z"}], "result_info": {"count": 5, "page": 4, "per_page": 7, "total_count": 10}}`)
				}))
			})
			It(`Invoke ListZoneApiGatewayOperations successfully with retries`, func() {
				aiSecurityForAppsService, serviceErr := aisecurityforappsv1.NewAiSecurityForAppsV1(&aisecurityforappsv1.AiSecurityForAppsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(aiSecurityForAppsService).ToNot(BeNil())
				aiSecurityForAppsService.EnableRetries(0, 0)

				// Construct an instance of the ListZoneApiGatewayOperationsOptions model
				listZoneApiGatewayOperationsOptionsModel := new(aisecurityforappsv1.ListZoneApiGatewayOperationsOptions)
				listZoneApiGatewayOperationsOptionsModel.Host = []string{"testString"}
				listZoneApiGatewayOperationsOptionsModel.Method = []string{"GET"}
				listZoneApiGatewayOperationsOptionsModel.Endpoint = core.StringPtr("testString")
				listZoneApiGatewayOperationsOptionsModel.Labels = []string{"testString"}
				listZoneApiGatewayOperationsOptionsModel.Direction = core.StringPtr("asc")
				listZoneApiGatewayOperationsOptionsModel.Order = core.StringPtr("host")
				listZoneApiGatewayOperationsOptionsModel.Page = core.Int64Ptr(int64(1))
				listZoneApiGatewayOperationsOptionsModel.PerPage = core.Int64Ptr(int64(1))
				listZoneApiGatewayOperationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := aiSecurityForAppsService.ListZoneApiGatewayOperationsWithContext(ctx, listZoneApiGatewayOperationsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				aiSecurityForAppsService.DisableRetries()
				result, response, operationErr := aiSecurityForAppsService.ListZoneApiGatewayOperations(listZoneApiGatewayOperationsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = aiSecurityForAppsService.ListZoneApiGatewayOperationsWithContext(ctx, listZoneApiGatewayOperationsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listZoneApiGatewayOperationsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["endpoint"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["direction"]).To(Equal([]string{"asc"}))
					Expect(req.URL.Query()["order"]).To(Equal([]string{"host"}))
					Expect(req.URL.Query()["page"]).To(Equal([]string{fmt.Sprint(int64(1))}))
					Expect(req.URL.Query()["per_page"]).To(Equal([]string{fmt.Sprint(int64(1))}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": [{"operation_id": "test-fafe-4643-bbbc-example", "method": "POST", "host": "api.example.com", "endpoint": "/v1/messages", "last_updated": "2024-01-01T00:00:00.000Z"}], "result_info": {"count": 5, "page": 4, "per_page": 7, "total_count": 10}}`)
				}))
			})
			It(`Invoke ListZoneApiGatewayOperations successfully`, func() {
				aiSecurityForAppsService, serviceErr := aisecurityforappsv1.NewAiSecurityForAppsV1(&aisecurityforappsv1.AiSecurityForAppsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(aiSecurityForAppsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := aiSecurityForAppsService.ListZoneApiGatewayOperations(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListZoneApiGatewayOperationsOptions model
				listZoneApiGatewayOperationsOptionsModel := new(aisecurityforappsv1.ListZoneApiGatewayOperationsOptions)
				listZoneApiGatewayOperationsOptionsModel.Host = []string{"testString"}
				listZoneApiGatewayOperationsOptionsModel.Method = []string{"GET"}
				listZoneApiGatewayOperationsOptionsModel.Endpoint = core.StringPtr("testString")
				listZoneApiGatewayOperationsOptionsModel.Labels = []string{"testString"}
				listZoneApiGatewayOperationsOptionsModel.Direction = core.StringPtr("asc")
				listZoneApiGatewayOperationsOptionsModel.Order = core.StringPtr("host")
				listZoneApiGatewayOperationsOptionsModel.Page = core.Int64Ptr(int64(1))
				listZoneApiGatewayOperationsOptionsModel.PerPage = core.Int64Ptr(int64(1))
				listZoneApiGatewayOperationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = aiSecurityForAppsService.ListZoneApiGatewayOperations(listZoneApiGatewayOperationsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListZoneApiGatewayOperations with error: Operation request error`, func() {
				aiSecurityForAppsService, serviceErr := aisecurityforappsv1.NewAiSecurityForAppsV1(&aisecurityforappsv1.AiSecurityForAppsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(aiSecurityForAppsService).ToNot(BeNil())

				// Construct an instance of the ListZoneApiGatewayOperationsOptions model
				listZoneApiGatewayOperationsOptionsModel := new(aisecurityforappsv1.ListZoneApiGatewayOperationsOptions)
				listZoneApiGatewayOperationsOptionsModel.Host = []string{"testString"}
				listZoneApiGatewayOperationsOptionsModel.Method = []string{"GET"}
				listZoneApiGatewayOperationsOptionsModel.Endpoint = core.StringPtr("testString")
				listZoneApiGatewayOperationsOptionsModel.Labels = []string{"testString"}
				listZoneApiGatewayOperationsOptionsModel.Direction = core.StringPtr("asc")
				listZoneApiGatewayOperationsOptionsModel.Order = core.StringPtr("host")
				listZoneApiGatewayOperationsOptionsModel.Page = core.Int64Ptr(int64(1))
				listZoneApiGatewayOperationsOptionsModel.PerPage = core.Int64Ptr(int64(1))
				listZoneApiGatewayOperationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := aiSecurityForAppsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := aiSecurityForAppsService.ListZoneApiGatewayOperations(listZoneApiGatewayOperationsOptionsModel)
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
			It(`Invoke ListZoneApiGatewayOperations successfully`, func() {
				aiSecurityForAppsService, serviceErr := aisecurityforappsv1.NewAiSecurityForAppsV1(&aisecurityforappsv1.AiSecurityForAppsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(aiSecurityForAppsService).ToNot(BeNil())

				// Construct an instance of the ListZoneApiGatewayOperationsOptions model
				listZoneApiGatewayOperationsOptionsModel := new(aisecurityforappsv1.ListZoneApiGatewayOperationsOptions)
				listZoneApiGatewayOperationsOptionsModel.Host = []string{"testString"}
				listZoneApiGatewayOperationsOptionsModel.Method = []string{"GET"}
				listZoneApiGatewayOperationsOptionsModel.Endpoint = core.StringPtr("testString")
				listZoneApiGatewayOperationsOptionsModel.Labels = []string{"testString"}
				listZoneApiGatewayOperationsOptionsModel.Direction = core.StringPtr("asc")
				listZoneApiGatewayOperationsOptionsModel.Order = core.StringPtr("host")
				listZoneApiGatewayOperationsOptionsModel.Page = core.Int64Ptr(int64(1))
				listZoneApiGatewayOperationsOptionsModel.PerPage = core.Int64Ptr(int64(1))
				listZoneApiGatewayOperationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := aiSecurityForAppsService.ListZoneApiGatewayOperations(listZoneApiGatewayOperationsOptionsModel)
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
	Describe(`CreateZoneApiGatewayOperation(createZoneApiGatewayOperationOptions *CreateZoneApiGatewayOperationOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		createZoneApiGatewayOperationPath := "/v1/testString/zones/testString/api_gateway/operations"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createZoneApiGatewayOperationPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateZoneApiGatewayOperation with error: Operation response processing error`, func() {
				aiSecurityForAppsService, serviceErr := aisecurityforappsv1.NewAiSecurityForAppsV1(&aisecurityforappsv1.AiSecurityForAppsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(aiSecurityForAppsService).ToNot(BeNil())

				// Construct an instance of the ApiGatewayOperation model
				apiGatewayOperationModel := new(aisecurityforappsv1.ApiGatewayOperation)
				apiGatewayOperationModel.Method = core.StringPtr("POST")
				apiGatewayOperationModel.Host = core.StringPtr("api.example.com")
				apiGatewayOperationModel.Endpoint = core.StringPtr("/v1/messages")

				// Construct an instance of the CreateZoneApiGatewayOperationOptions model
				createZoneApiGatewayOperationOptionsModel := new(aisecurityforappsv1.CreateZoneApiGatewayOperationOptions)
				createZoneApiGatewayOperationOptionsModel.ApiGatewayOperation = []aisecurityforappsv1.ApiGatewayOperation{*apiGatewayOperationModel}
				createZoneApiGatewayOperationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := aiSecurityForAppsService.CreateZoneApiGatewayOperation(createZoneApiGatewayOperationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				aiSecurityForAppsService.EnableRetries(0, 0)
				result, response, operationErr = aiSecurityForAppsService.CreateZoneApiGatewayOperation(createZoneApiGatewayOperationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateZoneApiGatewayOperation(createZoneApiGatewayOperationOptions *CreateZoneApiGatewayOperationOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		createZoneApiGatewayOperationPath := "/v1/testString/zones/testString/api_gateway/operations"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createZoneApiGatewayOperationPath))
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
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": [{"operation_id": "f174e90a-fafe-4643-bbbc-4a0ed4fc8415", "method": "POST", "host": "api.example.com", "endpoint": "/v1/messages"}]}`)
				}))
			})
			It(`Invoke CreateZoneApiGatewayOperation successfully with retries`, func() {
				aiSecurityForAppsService, serviceErr := aisecurityforappsv1.NewAiSecurityForAppsV1(&aisecurityforappsv1.AiSecurityForAppsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(aiSecurityForAppsService).ToNot(BeNil())
				aiSecurityForAppsService.EnableRetries(0, 0)

				// Construct an instance of the ApiGatewayOperation model
				apiGatewayOperationModel := new(aisecurityforappsv1.ApiGatewayOperation)
				apiGatewayOperationModel.Method = core.StringPtr("POST")
				apiGatewayOperationModel.Host = core.StringPtr("api.example.com")
				apiGatewayOperationModel.Endpoint = core.StringPtr("/v1/messages")

				// Construct an instance of the CreateZoneApiGatewayOperationOptions model
				createZoneApiGatewayOperationOptionsModel := new(aisecurityforappsv1.CreateZoneApiGatewayOperationOptions)
				createZoneApiGatewayOperationOptionsModel.ApiGatewayOperation = []aisecurityforappsv1.ApiGatewayOperation{*apiGatewayOperationModel}
				createZoneApiGatewayOperationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := aiSecurityForAppsService.CreateZoneApiGatewayOperationWithContext(ctx, createZoneApiGatewayOperationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				aiSecurityForAppsService.DisableRetries()
				result, response, operationErr := aiSecurityForAppsService.CreateZoneApiGatewayOperation(createZoneApiGatewayOperationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = aiSecurityForAppsService.CreateZoneApiGatewayOperationWithContext(ctx, createZoneApiGatewayOperationOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createZoneApiGatewayOperationPath))
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
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": [{"operation_id": "f174e90a-fafe-4643-bbbc-4a0ed4fc8415", "method": "POST", "host": "api.example.com", "endpoint": "/v1/messages"}]}`)
				}))
			})
			It(`Invoke CreateZoneApiGatewayOperation successfully`, func() {
				aiSecurityForAppsService, serviceErr := aisecurityforappsv1.NewAiSecurityForAppsV1(&aisecurityforappsv1.AiSecurityForAppsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(aiSecurityForAppsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := aiSecurityForAppsService.CreateZoneApiGatewayOperation(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ApiGatewayOperation model
				apiGatewayOperationModel := new(aisecurityforappsv1.ApiGatewayOperation)
				apiGatewayOperationModel.Method = core.StringPtr("POST")
				apiGatewayOperationModel.Host = core.StringPtr("api.example.com")
				apiGatewayOperationModel.Endpoint = core.StringPtr("/v1/messages")

				// Construct an instance of the CreateZoneApiGatewayOperationOptions model
				createZoneApiGatewayOperationOptionsModel := new(aisecurityforappsv1.CreateZoneApiGatewayOperationOptions)
				createZoneApiGatewayOperationOptionsModel.ApiGatewayOperation = []aisecurityforappsv1.ApiGatewayOperation{*apiGatewayOperationModel}
				createZoneApiGatewayOperationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = aiSecurityForAppsService.CreateZoneApiGatewayOperation(createZoneApiGatewayOperationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateZoneApiGatewayOperation with error: Operation request error`, func() {
				aiSecurityForAppsService, serviceErr := aisecurityforappsv1.NewAiSecurityForAppsV1(&aisecurityforappsv1.AiSecurityForAppsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(aiSecurityForAppsService).ToNot(BeNil())

				// Construct an instance of the ApiGatewayOperation model
				apiGatewayOperationModel := new(aisecurityforappsv1.ApiGatewayOperation)
				apiGatewayOperationModel.Method = core.StringPtr("POST")
				apiGatewayOperationModel.Host = core.StringPtr("api.example.com")
				apiGatewayOperationModel.Endpoint = core.StringPtr("/v1/messages")

				// Construct an instance of the CreateZoneApiGatewayOperationOptions model
				createZoneApiGatewayOperationOptionsModel := new(aisecurityforappsv1.CreateZoneApiGatewayOperationOptions)
				createZoneApiGatewayOperationOptionsModel.ApiGatewayOperation = []aisecurityforappsv1.ApiGatewayOperation{*apiGatewayOperationModel}
				createZoneApiGatewayOperationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := aiSecurityForAppsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := aiSecurityForAppsService.CreateZoneApiGatewayOperation(createZoneApiGatewayOperationOptionsModel)
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
			It(`Invoke CreateZoneApiGatewayOperation successfully`, func() {
				aiSecurityForAppsService, serviceErr := aisecurityforappsv1.NewAiSecurityForAppsV1(&aisecurityforappsv1.AiSecurityForAppsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(aiSecurityForAppsService).ToNot(BeNil())

				// Construct an instance of the ApiGatewayOperation model
				apiGatewayOperationModel := new(aisecurityforappsv1.ApiGatewayOperation)
				apiGatewayOperationModel.Method = core.StringPtr("POST")
				apiGatewayOperationModel.Host = core.StringPtr("api.example.com")
				apiGatewayOperationModel.Endpoint = core.StringPtr("/v1/messages")

				// Construct an instance of the CreateZoneApiGatewayOperationOptions model
				createZoneApiGatewayOperationOptionsModel := new(aisecurityforappsv1.CreateZoneApiGatewayOperationOptions)
				createZoneApiGatewayOperationOptionsModel.ApiGatewayOperation = []aisecurityforappsv1.ApiGatewayOperation{*apiGatewayOperationModel}
				createZoneApiGatewayOperationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := aiSecurityForAppsService.CreateZoneApiGatewayOperation(createZoneApiGatewayOperationOptionsModel)
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
	Describe(`CreateApiGatewayOperationItem(createApiGatewayOperationItemOptions *CreateApiGatewayOperationItemOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		createApiGatewayOperationItemPath := "/v1/testString/zones/testString/api_gateway/operations/item"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createApiGatewayOperationItemPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateApiGatewayOperationItem with error: Operation response processing error`, func() {
				aiSecurityForAppsService, serviceErr := aisecurityforappsv1.NewAiSecurityForAppsV1(&aisecurityforappsv1.AiSecurityForAppsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(aiSecurityForAppsService).ToNot(BeNil())

				// Construct an instance of the CreateApiGatewayOperationItemOptions model
				createApiGatewayOperationItemOptionsModel := new(aisecurityforappsv1.CreateApiGatewayOperationItemOptions)
				createApiGatewayOperationItemOptionsModel.Method = core.StringPtr("POST")
				createApiGatewayOperationItemOptionsModel.Host = core.StringPtr("api.example.com")
				createApiGatewayOperationItemOptionsModel.Endpoint = core.StringPtr("/v1/messages")
				createApiGatewayOperationItemOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := aiSecurityForAppsService.CreateApiGatewayOperationItem(createApiGatewayOperationItemOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				aiSecurityForAppsService.EnableRetries(0, 0)
				result, response, operationErr = aiSecurityForAppsService.CreateApiGatewayOperationItem(createApiGatewayOperationItemOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateApiGatewayOperationItem(createApiGatewayOperationItemOptions *CreateApiGatewayOperationItemOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		createApiGatewayOperationItemPath := "/v1/testString/zones/testString/api_gateway/operations/item"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createApiGatewayOperationItemPath))
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
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"operation_id": "f174e90a-fafe-4643-bbbc-4a0ed4fc8415", "method": "POST", "host": "api.example.com", "endpoint": "/v1/messages"}}`)
				}))
			})
			It(`Invoke CreateApiGatewayOperationItem successfully with retries`, func() {
				aiSecurityForAppsService, serviceErr := aisecurityforappsv1.NewAiSecurityForAppsV1(&aisecurityforappsv1.AiSecurityForAppsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(aiSecurityForAppsService).ToNot(BeNil())
				aiSecurityForAppsService.EnableRetries(0, 0)

				// Construct an instance of the CreateApiGatewayOperationItemOptions model
				createApiGatewayOperationItemOptionsModel := new(aisecurityforappsv1.CreateApiGatewayOperationItemOptions)
				createApiGatewayOperationItemOptionsModel.Method = core.StringPtr("POST")
				createApiGatewayOperationItemOptionsModel.Host = core.StringPtr("api.example.com")
				createApiGatewayOperationItemOptionsModel.Endpoint = core.StringPtr("/v1/messages")
				createApiGatewayOperationItemOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := aiSecurityForAppsService.CreateApiGatewayOperationItemWithContext(ctx, createApiGatewayOperationItemOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				aiSecurityForAppsService.DisableRetries()
				result, response, operationErr := aiSecurityForAppsService.CreateApiGatewayOperationItem(createApiGatewayOperationItemOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = aiSecurityForAppsService.CreateApiGatewayOperationItemWithContext(ctx, createApiGatewayOperationItemOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createApiGatewayOperationItemPath))
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
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"operation_id": "f174e90a-fafe-4643-bbbc-4a0ed4fc8415", "method": "POST", "host": "api.example.com", "endpoint": "/v1/messages"}}`)
				}))
			})
			It(`Invoke CreateApiGatewayOperationItem successfully`, func() {
				aiSecurityForAppsService, serviceErr := aisecurityforappsv1.NewAiSecurityForAppsV1(&aisecurityforappsv1.AiSecurityForAppsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(aiSecurityForAppsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := aiSecurityForAppsService.CreateApiGatewayOperationItem(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateApiGatewayOperationItemOptions model
				createApiGatewayOperationItemOptionsModel := new(aisecurityforappsv1.CreateApiGatewayOperationItemOptions)
				createApiGatewayOperationItemOptionsModel.Method = core.StringPtr("POST")
				createApiGatewayOperationItemOptionsModel.Host = core.StringPtr("api.example.com")
				createApiGatewayOperationItemOptionsModel.Endpoint = core.StringPtr("/v1/messages")
				createApiGatewayOperationItemOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = aiSecurityForAppsService.CreateApiGatewayOperationItem(createApiGatewayOperationItemOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateApiGatewayOperationItem with error: Operation request error`, func() {
				aiSecurityForAppsService, serviceErr := aisecurityforappsv1.NewAiSecurityForAppsV1(&aisecurityforappsv1.AiSecurityForAppsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(aiSecurityForAppsService).ToNot(BeNil())

				// Construct an instance of the CreateApiGatewayOperationItemOptions model
				createApiGatewayOperationItemOptionsModel := new(aisecurityforappsv1.CreateApiGatewayOperationItemOptions)
				createApiGatewayOperationItemOptionsModel.Method = core.StringPtr("POST")
				createApiGatewayOperationItemOptionsModel.Host = core.StringPtr("api.example.com")
				createApiGatewayOperationItemOptionsModel.Endpoint = core.StringPtr("/v1/messages")
				createApiGatewayOperationItemOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := aiSecurityForAppsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := aiSecurityForAppsService.CreateApiGatewayOperationItem(createApiGatewayOperationItemOptionsModel)
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
			It(`Invoke CreateApiGatewayOperationItem successfully`, func() {
				aiSecurityForAppsService, serviceErr := aisecurityforappsv1.NewAiSecurityForAppsV1(&aisecurityforappsv1.AiSecurityForAppsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(aiSecurityForAppsService).ToNot(BeNil())

				// Construct an instance of the CreateApiGatewayOperationItemOptions model
				createApiGatewayOperationItemOptionsModel := new(aisecurityforappsv1.CreateApiGatewayOperationItemOptions)
				createApiGatewayOperationItemOptionsModel.Method = core.StringPtr("POST")
				createApiGatewayOperationItemOptionsModel.Host = core.StringPtr("api.example.com")
				createApiGatewayOperationItemOptionsModel.Endpoint = core.StringPtr("/v1/messages")
				createApiGatewayOperationItemOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := aiSecurityForAppsService.CreateApiGatewayOperationItem(createApiGatewayOperationItemOptionsModel)
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
	Describe(`ReplaceZoneApiGatewayOperationLabels(replaceZoneApiGatewayOperationLabelsOptions *ReplaceZoneApiGatewayOperationLabelsOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		replaceZoneApiGatewayOperationLabelsPath := "/v1/testString/zones/testString/api_gateway/operations/labels"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(replaceZoneApiGatewayOperationLabelsPath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ReplaceZoneApiGatewayOperationLabels with error: Operation response processing error`, func() {
				aiSecurityForAppsService, serviceErr := aisecurityforappsv1.NewAiSecurityForAppsV1(&aisecurityforappsv1.AiSecurityForAppsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(aiSecurityForAppsService).ToNot(BeNil())

				// Construct an instance of the ApiGatewayOperationsLabelsBulkInputSelectorInclude model
				apiGatewayOperationsLabelsBulkInputSelectorIncludeModel := new(aisecurityforappsv1.ApiGatewayOperationsLabelsBulkInputSelectorInclude)
				apiGatewayOperationsLabelsBulkInputSelectorIncludeModel.OperationIds = []string{"f174e90a-fafe-4643-bbbc-4a0ed4fc8415"}

				// Construct an instance of the ApiGatewayOperationsLabelsBulkInputSelector model
				apiGatewayOperationsLabelsBulkInputSelectorModel := new(aisecurityforappsv1.ApiGatewayOperationsLabelsBulkInputSelector)
				apiGatewayOperationsLabelsBulkInputSelectorModel.Include = apiGatewayOperationsLabelsBulkInputSelectorIncludeModel

				// Construct an instance of the ApiGatewayOperationsLabelsBulkInputUser model
				apiGatewayOperationsLabelsBulkInputUserModel := new(aisecurityforappsv1.ApiGatewayOperationsLabelsBulkInputUser)
				apiGatewayOperationsLabelsBulkInputUserModel.Labels = []string{"testString"}

				// Construct an instance of the ApiGatewayOperationsLabelsBulkInputManaged model
				apiGatewayOperationsLabelsBulkInputManagedModel := new(aisecurityforappsv1.ApiGatewayOperationsLabelsBulkInputManaged)
				apiGatewayOperationsLabelsBulkInputManagedModel.Labels = []string{"testString"}

				// Construct an instance of the ReplaceZoneApiGatewayOperationLabelsOptions model
				replaceZoneApiGatewayOperationLabelsOptionsModel := new(aisecurityforappsv1.ReplaceZoneApiGatewayOperationLabelsOptions)
				replaceZoneApiGatewayOperationLabelsOptionsModel.Selector = apiGatewayOperationsLabelsBulkInputSelectorModel
				replaceZoneApiGatewayOperationLabelsOptionsModel.User = apiGatewayOperationsLabelsBulkInputUserModel
				replaceZoneApiGatewayOperationLabelsOptionsModel.Managed = apiGatewayOperationsLabelsBulkInputManagedModel
				replaceZoneApiGatewayOperationLabelsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := aiSecurityForAppsService.ReplaceZoneApiGatewayOperationLabels(replaceZoneApiGatewayOperationLabelsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				aiSecurityForAppsService.EnableRetries(0, 0)
				result, response, operationErr = aiSecurityForAppsService.ReplaceZoneApiGatewayOperationLabels(replaceZoneApiGatewayOperationLabelsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ReplaceZoneApiGatewayOperationLabels(replaceZoneApiGatewayOperationLabelsOptions *ReplaceZoneApiGatewayOperationLabelsOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		replaceZoneApiGatewayOperationLabelsPath := "/v1/testString/zones/testString/api_gateway/operations/labels"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(replaceZoneApiGatewayOperationLabelsPath))
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
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": [{"operation_id": "f174e90a-fafe-4643-bbbc-4a0ed4fc8415", "labels": [{"name": "cf-llm", "source": "managed"}]}]}`)
				}))
			})
			It(`Invoke ReplaceZoneApiGatewayOperationLabels successfully with retries`, func() {
				aiSecurityForAppsService, serviceErr := aisecurityforappsv1.NewAiSecurityForAppsV1(&aisecurityforappsv1.AiSecurityForAppsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(aiSecurityForAppsService).ToNot(BeNil())
				aiSecurityForAppsService.EnableRetries(0, 0)

				// Construct an instance of the ApiGatewayOperationsLabelsBulkInputSelectorInclude model
				apiGatewayOperationsLabelsBulkInputSelectorIncludeModel := new(aisecurityforappsv1.ApiGatewayOperationsLabelsBulkInputSelectorInclude)
				apiGatewayOperationsLabelsBulkInputSelectorIncludeModel.OperationIds = []string{"f174e90a-fafe-4643-bbbc-4a0ed4fc8415"}

				// Construct an instance of the ApiGatewayOperationsLabelsBulkInputSelector model
				apiGatewayOperationsLabelsBulkInputSelectorModel := new(aisecurityforappsv1.ApiGatewayOperationsLabelsBulkInputSelector)
				apiGatewayOperationsLabelsBulkInputSelectorModel.Include = apiGatewayOperationsLabelsBulkInputSelectorIncludeModel

				// Construct an instance of the ApiGatewayOperationsLabelsBulkInputUser model
				apiGatewayOperationsLabelsBulkInputUserModel := new(aisecurityforappsv1.ApiGatewayOperationsLabelsBulkInputUser)
				apiGatewayOperationsLabelsBulkInputUserModel.Labels = []string{"testString"}

				// Construct an instance of the ApiGatewayOperationsLabelsBulkInputManaged model
				apiGatewayOperationsLabelsBulkInputManagedModel := new(aisecurityforappsv1.ApiGatewayOperationsLabelsBulkInputManaged)
				apiGatewayOperationsLabelsBulkInputManagedModel.Labels = []string{"testString"}

				// Construct an instance of the ReplaceZoneApiGatewayOperationLabelsOptions model
				replaceZoneApiGatewayOperationLabelsOptionsModel := new(aisecurityforappsv1.ReplaceZoneApiGatewayOperationLabelsOptions)
				replaceZoneApiGatewayOperationLabelsOptionsModel.Selector = apiGatewayOperationsLabelsBulkInputSelectorModel
				replaceZoneApiGatewayOperationLabelsOptionsModel.User = apiGatewayOperationsLabelsBulkInputUserModel
				replaceZoneApiGatewayOperationLabelsOptionsModel.Managed = apiGatewayOperationsLabelsBulkInputManagedModel
				replaceZoneApiGatewayOperationLabelsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := aiSecurityForAppsService.ReplaceZoneApiGatewayOperationLabelsWithContext(ctx, replaceZoneApiGatewayOperationLabelsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				aiSecurityForAppsService.DisableRetries()
				result, response, operationErr := aiSecurityForAppsService.ReplaceZoneApiGatewayOperationLabels(replaceZoneApiGatewayOperationLabelsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = aiSecurityForAppsService.ReplaceZoneApiGatewayOperationLabelsWithContext(ctx, replaceZoneApiGatewayOperationLabelsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(replaceZoneApiGatewayOperationLabelsPath))
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
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": [{"operation_id": "f174e90a-fafe-4643-bbbc-4a0ed4fc8415", "labels": [{"name": "cf-llm", "source": "managed"}]}]}`)
				}))
			})
			It(`Invoke ReplaceZoneApiGatewayOperationLabels successfully`, func() {
				aiSecurityForAppsService, serviceErr := aisecurityforappsv1.NewAiSecurityForAppsV1(&aisecurityforappsv1.AiSecurityForAppsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(aiSecurityForAppsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := aiSecurityForAppsService.ReplaceZoneApiGatewayOperationLabels(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ApiGatewayOperationsLabelsBulkInputSelectorInclude model
				apiGatewayOperationsLabelsBulkInputSelectorIncludeModel := new(aisecurityforappsv1.ApiGatewayOperationsLabelsBulkInputSelectorInclude)
				apiGatewayOperationsLabelsBulkInputSelectorIncludeModel.OperationIds = []string{"f174e90a-fafe-4643-bbbc-4a0ed4fc8415"}

				// Construct an instance of the ApiGatewayOperationsLabelsBulkInputSelector model
				apiGatewayOperationsLabelsBulkInputSelectorModel := new(aisecurityforappsv1.ApiGatewayOperationsLabelsBulkInputSelector)
				apiGatewayOperationsLabelsBulkInputSelectorModel.Include = apiGatewayOperationsLabelsBulkInputSelectorIncludeModel

				// Construct an instance of the ApiGatewayOperationsLabelsBulkInputUser model
				apiGatewayOperationsLabelsBulkInputUserModel := new(aisecurityforappsv1.ApiGatewayOperationsLabelsBulkInputUser)
				apiGatewayOperationsLabelsBulkInputUserModel.Labels = []string{"testString"}

				// Construct an instance of the ApiGatewayOperationsLabelsBulkInputManaged model
				apiGatewayOperationsLabelsBulkInputManagedModel := new(aisecurityforappsv1.ApiGatewayOperationsLabelsBulkInputManaged)
				apiGatewayOperationsLabelsBulkInputManagedModel.Labels = []string{"testString"}

				// Construct an instance of the ReplaceZoneApiGatewayOperationLabelsOptions model
				replaceZoneApiGatewayOperationLabelsOptionsModel := new(aisecurityforappsv1.ReplaceZoneApiGatewayOperationLabelsOptions)
				replaceZoneApiGatewayOperationLabelsOptionsModel.Selector = apiGatewayOperationsLabelsBulkInputSelectorModel
				replaceZoneApiGatewayOperationLabelsOptionsModel.User = apiGatewayOperationsLabelsBulkInputUserModel
				replaceZoneApiGatewayOperationLabelsOptionsModel.Managed = apiGatewayOperationsLabelsBulkInputManagedModel
				replaceZoneApiGatewayOperationLabelsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = aiSecurityForAppsService.ReplaceZoneApiGatewayOperationLabels(replaceZoneApiGatewayOperationLabelsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ReplaceZoneApiGatewayOperationLabels with error: Operation request error`, func() {
				aiSecurityForAppsService, serviceErr := aisecurityforappsv1.NewAiSecurityForAppsV1(&aisecurityforappsv1.AiSecurityForAppsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(aiSecurityForAppsService).ToNot(BeNil())

				// Construct an instance of the ApiGatewayOperationsLabelsBulkInputSelectorInclude model
				apiGatewayOperationsLabelsBulkInputSelectorIncludeModel := new(aisecurityforappsv1.ApiGatewayOperationsLabelsBulkInputSelectorInclude)
				apiGatewayOperationsLabelsBulkInputSelectorIncludeModel.OperationIds = []string{"f174e90a-fafe-4643-bbbc-4a0ed4fc8415"}

				// Construct an instance of the ApiGatewayOperationsLabelsBulkInputSelector model
				apiGatewayOperationsLabelsBulkInputSelectorModel := new(aisecurityforappsv1.ApiGatewayOperationsLabelsBulkInputSelector)
				apiGatewayOperationsLabelsBulkInputSelectorModel.Include = apiGatewayOperationsLabelsBulkInputSelectorIncludeModel

				// Construct an instance of the ApiGatewayOperationsLabelsBulkInputUser model
				apiGatewayOperationsLabelsBulkInputUserModel := new(aisecurityforappsv1.ApiGatewayOperationsLabelsBulkInputUser)
				apiGatewayOperationsLabelsBulkInputUserModel.Labels = []string{"testString"}

				// Construct an instance of the ApiGatewayOperationsLabelsBulkInputManaged model
				apiGatewayOperationsLabelsBulkInputManagedModel := new(aisecurityforappsv1.ApiGatewayOperationsLabelsBulkInputManaged)
				apiGatewayOperationsLabelsBulkInputManagedModel.Labels = []string{"testString"}

				// Construct an instance of the ReplaceZoneApiGatewayOperationLabelsOptions model
				replaceZoneApiGatewayOperationLabelsOptionsModel := new(aisecurityforappsv1.ReplaceZoneApiGatewayOperationLabelsOptions)
				replaceZoneApiGatewayOperationLabelsOptionsModel.Selector = apiGatewayOperationsLabelsBulkInputSelectorModel
				replaceZoneApiGatewayOperationLabelsOptionsModel.User = apiGatewayOperationsLabelsBulkInputUserModel
				replaceZoneApiGatewayOperationLabelsOptionsModel.Managed = apiGatewayOperationsLabelsBulkInputManagedModel
				replaceZoneApiGatewayOperationLabelsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := aiSecurityForAppsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := aiSecurityForAppsService.ReplaceZoneApiGatewayOperationLabels(replaceZoneApiGatewayOperationLabelsOptionsModel)
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
			It(`Invoke ReplaceZoneApiGatewayOperationLabels successfully`, func() {
				aiSecurityForAppsService, serviceErr := aisecurityforappsv1.NewAiSecurityForAppsV1(&aisecurityforappsv1.AiSecurityForAppsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(aiSecurityForAppsService).ToNot(BeNil())

				// Construct an instance of the ApiGatewayOperationsLabelsBulkInputSelectorInclude model
				apiGatewayOperationsLabelsBulkInputSelectorIncludeModel := new(aisecurityforappsv1.ApiGatewayOperationsLabelsBulkInputSelectorInclude)
				apiGatewayOperationsLabelsBulkInputSelectorIncludeModel.OperationIds = []string{"f174e90a-fafe-4643-bbbc-4a0ed4fc8415"}

				// Construct an instance of the ApiGatewayOperationsLabelsBulkInputSelector model
				apiGatewayOperationsLabelsBulkInputSelectorModel := new(aisecurityforappsv1.ApiGatewayOperationsLabelsBulkInputSelector)
				apiGatewayOperationsLabelsBulkInputSelectorModel.Include = apiGatewayOperationsLabelsBulkInputSelectorIncludeModel

				// Construct an instance of the ApiGatewayOperationsLabelsBulkInputUser model
				apiGatewayOperationsLabelsBulkInputUserModel := new(aisecurityforappsv1.ApiGatewayOperationsLabelsBulkInputUser)
				apiGatewayOperationsLabelsBulkInputUserModel.Labels = []string{"testString"}

				// Construct an instance of the ApiGatewayOperationsLabelsBulkInputManaged model
				apiGatewayOperationsLabelsBulkInputManagedModel := new(aisecurityforappsv1.ApiGatewayOperationsLabelsBulkInputManaged)
				apiGatewayOperationsLabelsBulkInputManagedModel.Labels = []string{"testString"}

				// Construct an instance of the ReplaceZoneApiGatewayOperationLabelsOptions model
				replaceZoneApiGatewayOperationLabelsOptionsModel := new(aisecurityforappsv1.ReplaceZoneApiGatewayOperationLabelsOptions)
				replaceZoneApiGatewayOperationLabelsOptionsModel.Selector = apiGatewayOperationsLabelsBulkInputSelectorModel
				replaceZoneApiGatewayOperationLabelsOptionsModel.User = apiGatewayOperationsLabelsBulkInputUserModel
				replaceZoneApiGatewayOperationLabelsOptionsModel.Managed = apiGatewayOperationsLabelsBulkInputManagedModel
				replaceZoneApiGatewayOperationLabelsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := aiSecurityForAppsService.ReplaceZoneApiGatewayOperationLabels(replaceZoneApiGatewayOperationLabelsOptionsModel)
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
	Describe(`DeleteZoneApiGatewayOperationLabels(deleteZoneApiGatewayOperationLabelsOptions *DeleteZoneApiGatewayOperationLabelsOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		deleteZoneApiGatewayOperationLabelsPath := "/v1/testString/zones/testString/api_gateway/operations/labels"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteZoneApiGatewayOperationLabelsPath))
					Expect(req.Method).To(Equal("DELETE"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke DeleteZoneApiGatewayOperationLabels with error: Operation response processing error`, func() {
				aiSecurityForAppsService, serviceErr := aisecurityforappsv1.NewAiSecurityForAppsV1(&aisecurityforappsv1.AiSecurityForAppsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(aiSecurityForAppsService).ToNot(BeNil())

				// Construct an instance of the ApiGatewayOperationsLabelsBulkDeleteInputSelectorInclude model
				apiGatewayOperationsLabelsBulkDeleteInputSelectorIncludeModel := new(aisecurityforappsv1.ApiGatewayOperationsLabelsBulkDeleteInputSelectorInclude)
				apiGatewayOperationsLabelsBulkDeleteInputSelectorIncludeModel.OperationIds = []string{"f174e90a-fafe-4643-bbbc-4a0ed4fc8415"}

				// Construct an instance of the ApiGatewayOperationsLabelsBulkDeleteInputSelector model
				apiGatewayOperationsLabelsBulkDeleteInputSelectorModel := new(aisecurityforappsv1.ApiGatewayOperationsLabelsBulkDeleteInputSelector)
				apiGatewayOperationsLabelsBulkDeleteInputSelectorModel.Include = apiGatewayOperationsLabelsBulkDeleteInputSelectorIncludeModel

				// Construct an instance of the ApiGatewayOperationsLabelsBulkDeleteInputUser model
				apiGatewayOperationsLabelsBulkDeleteInputUserModel := new(aisecurityforappsv1.ApiGatewayOperationsLabelsBulkDeleteInputUser)
				apiGatewayOperationsLabelsBulkDeleteInputUserModel.Labels = []string{"testString"}

				// Construct an instance of the ApiGatewayOperationsLabelsBulkDeleteInputManaged model
				apiGatewayOperationsLabelsBulkDeleteInputManagedModel := new(aisecurityforappsv1.ApiGatewayOperationsLabelsBulkDeleteInputManaged)
				apiGatewayOperationsLabelsBulkDeleteInputManagedModel.Labels = []string{"testString"}

				// Construct an instance of the DeleteZoneApiGatewayOperationLabelsOptions model
				deleteZoneApiGatewayOperationLabelsOptionsModel := new(aisecurityforappsv1.DeleteZoneApiGatewayOperationLabelsOptions)
				deleteZoneApiGatewayOperationLabelsOptionsModel.Selector = apiGatewayOperationsLabelsBulkDeleteInputSelectorModel
				deleteZoneApiGatewayOperationLabelsOptionsModel.User = apiGatewayOperationsLabelsBulkDeleteInputUserModel
				deleteZoneApiGatewayOperationLabelsOptionsModel.Managed = apiGatewayOperationsLabelsBulkDeleteInputManagedModel
				deleteZoneApiGatewayOperationLabelsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := aiSecurityForAppsService.DeleteZoneApiGatewayOperationLabels(deleteZoneApiGatewayOperationLabelsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				aiSecurityForAppsService.EnableRetries(0, 0)
				result, response, operationErr = aiSecurityForAppsService.DeleteZoneApiGatewayOperationLabels(deleteZoneApiGatewayOperationLabelsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteZoneApiGatewayOperationLabels(deleteZoneApiGatewayOperationLabelsOptions *DeleteZoneApiGatewayOperationLabelsOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		deleteZoneApiGatewayOperationLabelsPath := "/v1/testString/zones/testString/api_gateway/operations/labels"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteZoneApiGatewayOperationLabelsPath))
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
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": "anyValue"}`)
				}))
			})
			It(`Invoke DeleteZoneApiGatewayOperationLabels successfully with retries`, func() {
				aiSecurityForAppsService, serviceErr := aisecurityforappsv1.NewAiSecurityForAppsV1(&aisecurityforappsv1.AiSecurityForAppsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(aiSecurityForAppsService).ToNot(BeNil())
				aiSecurityForAppsService.EnableRetries(0, 0)

				// Construct an instance of the ApiGatewayOperationsLabelsBulkDeleteInputSelectorInclude model
				apiGatewayOperationsLabelsBulkDeleteInputSelectorIncludeModel := new(aisecurityforappsv1.ApiGatewayOperationsLabelsBulkDeleteInputSelectorInclude)
				apiGatewayOperationsLabelsBulkDeleteInputSelectorIncludeModel.OperationIds = []string{"f174e90a-fafe-4643-bbbc-4a0ed4fc8415"}

				// Construct an instance of the ApiGatewayOperationsLabelsBulkDeleteInputSelector model
				apiGatewayOperationsLabelsBulkDeleteInputSelectorModel := new(aisecurityforappsv1.ApiGatewayOperationsLabelsBulkDeleteInputSelector)
				apiGatewayOperationsLabelsBulkDeleteInputSelectorModel.Include = apiGatewayOperationsLabelsBulkDeleteInputSelectorIncludeModel

				// Construct an instance of the ApiGatewayOperationsLabelsBulkDeleteInputUser model
				apiGatewayOperationsLabelsBulkDeleteInputUserModel := new(aisecurityforappsv1.ApiGatewayOperationsLabelsBulkDeleteInputUser)
				apiGatewayOperationsLabelsBulkDeleteInputUserModel.Labels = []string{"testString"}

				// Construct an instance of the ApiGatewayOperationsLabelsBulkDeleteInputManaged model
				apiGatewayOperationsLabelsBulkDeleteInputManagedModel := new(aisecurityforappsv1.ApiGatewayOperationsLabelsBulkDeleteInputManaged)
				apiGatewayOperationsLabelsBulkDeleteInputManagedModel.Labels = []string{"testString"}

				// Construct an instance of the DeleteZoneApiGatewayOperationLabelsOptions model
				deleteZoneApiGatewayOperationLabelsOptionsModel := new(aisecurityforappsv1.DeleteZoneApiGatewayOperationLabelsOptions)
				deleteZoneApiGatewayOperationLabelsOptionsModel.Selector = apiGatewayOperationsLabelsBulkDeleteInputSelectorModel
				deleteZoneApiGatewayOperationLabelsOptionsModel.User = apiGatewayOperationsLabelsBulkDeleteInputUserModel
				deleteZoneApiGatewayOperationLabelsOptionsModel.Managed = apiGatewayOperationsLabelsBulkDeleteInputManagedModel
				deleteZoneApiGatewayOperationLabelsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := aiSecurityForAppsService.DeleteZoneApiGatewayOperationLabelsWithContext(ctx, deleteZoneApiGatewayOperationLabelsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				aiSecurityForAppsService.DisableRetries()
				result, response, operationErr := aiSecurityForAppsService.DeleteZoneApiGatewayOperationLabels(deleteZoneApiGatewayOperationLabelsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = aiSecurityForAppsService.DeleteZoneApiGatewayOperationLabelsWithContext(ctx, deleteZoneApiGatewayOperationLabelsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(deleteZoneApiGatewayOperationLabelsPath))
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
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": "anyValue"}`)
				}))
			})
			It(`Invoke DeleteZoneApiGatewayOperationLabels successfully`, func() {
				aiSecurityForAppsService, serviceErr := aisecurityforappsv1.NewAiSecurityForAppsV1(&aisecurityforappsv1.AiSecurityForAppsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(aiSecurityForAppsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := aiSecurityForAppsService.DeleteZoneApiGatewayOperationLabels(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ApiGatewayOperationsLabelsBulkDeleteInputSelectorInclude model
				apiGatewayOperationsLabelsBulkDeleteInputSelectorIncludeModel := new(aisecurityforappsv1.ApiGatewayOperationsLabelsBulkDeleteInputSelectorInclude)
				apiGatewayOperationsLabelsBulkDeleteInputSelectorIncludeModel.OperationIds = []string{"f174e90a-fafe-4643-bbbc-4a0ed4fc8415"}

				// Construct an instance of the ApiGatewayOperationsLabelsBulkDeleteInputSelector model
				apiGatewayOperationsLabelsBulkDeleteInputSelectorModel := new(aisecurityforappsv1.ApiGatewayOperationsLabelsBulkDeleteInputSelector)
				apiGatewayOperationsLabelsBulkDeleteInputSelectorModel.Include = apiGatewayOperationsLabelsBulkDeleteInputSelectorIncludeModel

				// Construct an instance of the ApiGatewayOperationsLabelsBulkDeleteInputUser model
				apiGatewayOperationsLabelsBulkDeleteInputUserModel := new(aisecurityforappsv1.ApiGatewayOperationsLabelsBulkDeleteInputUser)
				apiGatewayOperationsLabelsBulkDeleteInputUserModel.Labels = []string{"testString"}

				// Construct an instance of the ApiGatewayOperationsLabelsBulkDeleteInputManaged model
				apiGatewayOperationsLabelsBulkDeleteInputManagedModel := new(aisecurityforappsv1.ApiGatewayOperationsLabelsBulkDeleteInputManaged)
				apiGatewayOperationsLabelsBulkDeleteInputManagedModel.Labels = []string{"testString"}

				// Construct an instance of the DeleteZoneApiGatewayOperationLabelsOptions model
				deleteZoneApiGatewayOperationLabelsOptionsModel := new(aisecurityforappsv1.DeleteZoneApiGatewayOperationLabelsOptions)
				deleteZoneApiGatewayOperationLabelsOptionsModel.Selector = apiGatewayOperationsLabelsBulkDeleteInputSelectorModel
				deleteZoneApiGatewayOperationLabelsOptionsModel.User = apiGatewayOperationsLabelsBulkDeleteInputUserModel
				deleteZoneApiGatewayOperationLabelsOptionsModel.Managed = apiGatewayOperationsLabelsBulkDeleteInputManagedModel
				deleteZoneApiGatewayOperationLabelsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = aiSecurityForAppsService.DeleteZoneApiGatewayOperationLabels(deleteZoneApiGatewayOperationLabelsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke DeleteZoneApiGatewayOperationLabels with error: Operation validation and request error`, func() {
				aiSecurityForAppsService, serviceErr := aisecurityforappsv1.NewAiSecurityForAppsV1(&aisecurityforappsv1.AiSecurityForAppsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(aiSecurityForAppsService).ToNot(BeNil())

				// Construct an instance of the ApiGatewayOperationsLabelsBulkDeleteInputSelectorInclude model
				apiGatewayOperationsLabelsBulkDeleteInputSelectorIncludeModel := new(aisecurityforappsv1.ApiGatewayOperationsLabelsBulkDeleteInputSelectorInclude)
				apiGatewayOperationsLabelsBulkDeleteInputSelectorIncludeModel.OperationIds = []string{"f174e90a-fafe-4643-bbbc-4a0ed4fc8415"}

				// Construct an instance of the ApiGatewayOperationsLabelsBulkDeleteInputSelector model
				apiGatewayOperationsLabelsBulkDeleteInputSelectorModel := new(aisecurityforappsv1.ApiGatewayOperationsLabelsBulkDeleteInputSelector)
				apiGatewayOperationsLabelsBulkDeleteInputSelectorModel.Include = apiGatewayOperationsLabelsBulkDeleteInputSelectorIncludeModel

				// Construct an instance of the ApiGatewayOperationsLabelsBulkDeleteInputUser model
				apiGatewayOperationsLabelsBulkDeleteInputUserModel := new(aisecurityforappsv1.ApiGatewayOperationsLabelsBulkDeleteInputUser)
				apiGatewayOperationsLabelsBulkDeleteInputUserModel.Labels = []string{"testString"}

				// Construct an instance of the ApiGatewayOperationsLabelsBulkDeleteInputManaged model
				apiGatewayOperationsLabelsBulkDeleteInputManagedModel := new(aisecurityforappsv1.ApiGatewayOperationsLabelsBulkDeleteInputManaged)
				apiGatewayOperationsLabelsBulkDeleteInputManagedModel.Labels = []string{"testString"}

				// Construct an instance of the DeleteZoneApiGatewayOperationLabelsOptions model
				deleteZoneApiGatewayOperationLabelsOptionsModel := new(aisecurityforappsv1.DeleteZoneApiGatewayOperationLabelsOptions)
				deleteZoneApiGatewayOperationLabelsOptionsModel.Selector = apiGatewayOperationsLabelsBulkDeleteInputSelectorModel
				deleteZoneApiGatewayOperationLabelsOptionsModel.User = apiGatewayOperationsLabelsBulkDeleteInputUserModel
				deleteZoneApiGatewayOperationLabelsOptionsModel.Managed = apiGatewayOperationsLabelsBulkDeleteInputManagedModel
				deleteZoneApiGatewayOperationLabelsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := aiSecurityForAppsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := aiSecurityForAppsService.DeleteZoneApiGatewayOperationLabels(deleteZoneApiGatewayOperationLabelsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the DeleteZoneApiGatewayOperationLabelsOptions model with no property values
				deleteZoneApiGatewayOperationLabelsOptionsModelNew := new(aisecurityforappsv1.DeleteZoneApiGatewayOperationLabelsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = aiSecurityForAppsService.DeleteZoneApiGatewayOperationLabels(deleteZoneApiGatewayOperationLabelsOptionsModelNew)
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
			It(`Invoke DeleteZoneApiGatewayOperationLabels successfully`, func() {
				aiSecurityForAppsService, serviceErr := aisecurityforappsv1.NewAiSecurityForAppsV1(&aisecurityforappsv1.AiSecurityForAppsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(aiSecurityForAppsService).ToNot(BeNil())

				// Construct an instance of the ApiGatewayOperationsLabelsBulkDeleteInputSelectorInclude model
				apiGatewayOperationsLabelsBulkDeleteInputSelectorIncludeModel := new(aisecurityforappsv1.ApiGatewayOperationsLabelsBulkDeleteInputSelectorInclude)
				apiGatewayOperationsLabelsBulkDeleteInputSelectorIncludeModel.OperationIds = []string{"f174e90a-fafe-4643-bbbc-4a0ed4fc8415"}

				// Construct an instance of the ApiGatewayOperationsLabelsBulkDeleteInputSelector model
				apiGatewayOperationsLabelsBulkDeleteInputSelectorModel := new(aisecurityforappsv1.ApiGatewayOperationsLabelsBulkDeleteInputSelector)
				apiGatewayOperationsLabelsBulkDeleteInputSelectorModel.Include = apiGatewayOperationsLabelsBulkDeleteInputSelectorIncludeModel

				// Construct an instance of the ApiGatewayOperationsLabelsBulkDeleteInputUser model
				apiGatewayOperationsLabelsBulkDeleteInputUserModel := new(aisecurityforappsv1.ApiGatewayOperationsLabelsBulkDeleteInputUser)
				apiGatewayOperationsLabelsBulkDeleteInputUserModel.Labels = []string{"testString"}

				// Construct an instance of the ApiGatewayOperationsLabelsBulkDeleteInputManaged model
				apiGatewayOperationsLabelsBulkDeleteInputManagedModel := new(aisecurityforappsv1.ApiGatewayOperationsLabelsBulkDeleteInputManaged)
				apiGatewayOperationsLabelsBulkDeleteInputManagedModel.Labels = []string{"testString"}

				// Construct an instance of the DeleteZoneApiGatewayOperationLabelsOptions model
				deleteZoneApiGatewayOperationLabelsOptionsModel := new(aisecurityforappsv1.DeleteZoneApiGatewayOperationLabelsOptions)
				deleteZoneApiGatewayOperationLabelsOptionsModel.Selector = apiGatewayOperationsLabelsBulkDeleteInputSelectorModel
				deleteZoneApiGatewayOperationLabelsOptionsModel.User = apiGatewayOperationsLabelsBulkDeleteInputUserModel
				deleteZoneApiGatewayOperationLabelsOptionsModel.Managed = apiGatewayOperationsLabelsBulkDeleteInputManagedModel
				deleteZoneApiGatewayOperationLabelsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := aiSecurityForAppsService.DeleteZoneApiGatewayOperationLabels(deleteZoneApiGatewayOperationLabelsOptionsModel)
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
	Describe(`UpdateApiGatewayOperationLabels(updateApiGatewayOperationLabelsOptions *UpdateApiGatewayOperationLabelsOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		updateApiGatewayOperationLabelsPath := "/v1/testString/zones/testString/api_gateway/operations/labels"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateApiGatewayOperationLabelsPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateApiGatewayOperationLabels with error: Operation response processing error`, func() {
				aiSecurityForAppsService, serviceErr := aisecurityforappsv1.NewAiSecurityForAppsV1(&aisecurityforappsv1.AiSecurityForAppsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(aiSecurityForAppsService).ToNot(BeNil())

				// Construct an instance of the ApiGatewayOperationsLabelsBulkInputSelectorInclude model
				apiGatewayOperationsLabelsBulkInputSelectorIncludeModel := new(aisecurityforappsv1.ApiGatewayOperationsLabelsBulkInputSelectorInclude)
				apiGatewayOperationsLabelsBulkInputSelectorIncludeModel.OperationIds = []string{"f174e90a-fafe-4643-bbbc-4a0ed4fc8415"}

				// Construct an instance of the ApiGatewayOperationsLabelsBulkInputSelector model
				apiGatewayOperationsLabelsBulkInputSelectorModel := new(aisecurityforappsv1.ApiGatewayOperationsLabelsBulkInputSelector)
				apiGatewayOperationsLabelsBulkInputSelectorModel.Include = apiGatewayOperationsLabelsBulkInputSelectorIncludeModel

				// Construct an instance of the ApiGatewayOperationsLabelsBulkInputUser model
				apiGatewayOperationsLabelsBulkInputUserModel := new(aisecurityforappsv1.ApiGatewayOperationsLabelsBulkInputUser)
				apiGatewayOperationsLabelsBulkInputUserModel.Labels = []string{"testString"}

				// Construct an instance of the ApiGatewayOperationsLabelsBulkInputManaged model
				apiGatewayOperationsLabelsBulkInputManagedModel := new(aisecurityforappsv1.ApiGatewayOperationsLabelsBulkInputManaged)
				apiGatewayOperationsLabelsBulkInputManagedModel.Labels = []string{"testString"}

				// Construct an instance of the UpdateApiGatewayOperationLabelsOptions model
				updateApiGatewayOperationLabelsOptionsModel := new(aisecurityforappsv1.UpdateApiGatewayOperationLabelsOptions)
				updateApiGatewayOperationLabelsOptionsModel.Selector = apiGatewayOperationsLabelsBulkInputSelectorModel
				updateApiGatewayOperationLabelsOptionsModel.User = apiGatewayOperationsLabelsBulkInputUserModel
				updateApiGatewayOperationLabelsOptionsModel.Managed = apiGatewayOperationsLabelsBulkInputManagedModel
				updateApiGatewayOperationLabelsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := aiSecurityForAppsService.UpdateApiGatewayOperationLabels(updateApiGatewayOperationLabelsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				aiSecurityForAppsService.EnableRetries(0, 0)
				result, response, operationErr = aiSecurityForAppsService.UpdateApiGatewayOperationLabels(updateApiGatewayOperationLabelsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateApiGatewayOperationLabels(updateApiGatewayOperationLabelsOptions *UpdateApiGatewayOperationLabelsOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		updateApiGatewayOperationLabelsPath := "/v1/testString/zones/testString/api_gateway/operations/labels"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateApiGatewayOperationLabelsPath))
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
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": [{"operation_id": "f174e90a-fafe-4643-bbbc-4a0ed4fc8415", "labels": [{"name": "cf-llm", "source": "managed"}]}]}`)
				}))
			})
			It(`Invoke UpdateApiGatewayOperationLabels successfully with retries`, func() {
				aiSecurityForAppsService, serviceErr := aisecurityforappsv1.NewAiSecurityForAppsV1(&aisecurityforappsv1.AiSecurityForAppsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(aiSecurityForAppsService).ToNot(BeNil())
				aiSecurityForAppsService.EnableRetries(0, 0)

				// Construct an instance of the ApiGatewayOperationsLabelsBulkInputSelectorInclude model
				apiGatewayOperationsLabelsBulkInputSelectorIncludeModel := new(aisecurityforappsv1.ApiGatewayOperationsLabelsBulkInputSelectorInclude)
				apiGatewayOperationsLabelsBulkInputSelectorIncludeModel.OperationIds = []string{"f174e90a-fafe-4643-bbbc-4a0ed4fc8415"}

				// Construct an instance of the ApiGatewayOperationsLabelsBulkInputSelector model
				apiGatewayOperationsLabelsBulkInputSelectorModel := new(aisecurityforappsv1.ApiGatewayOperationsLabelsBulkInputSelector)
				apiGatewayOperationsLabelsBulkInputSelectorModel.Include = apiGatewayOperationsLabelsBulkInputSelectorIncludeModel

				// Construct an instance of the ApiGatewayOperationsLabelsBulkInputUser model
				apiGatewayOperationsLabelsBulkInputUserModel := new(aisecurityforappsv1.ApiGatewayOperationsLabelsBulkInputUser)
				apiGatewayOperationsLabelsBulkInputUserModel.Labels = []string{"testString"}

				// Construct an instance of the ApiGatewayOperationsLabelsBulkInputManaged model
				apiGatewayOperationsLabelsBulkInputManagedModel := new(aisecurityforappsv1.ApiGatewayOperationsLabelsBulkInputManaged)
				apiGatewayOperationsLabelsBulkInputManagedModel.Labels = []string{"testString"}

				// Construct an instance of the UpdateApiGatewayOperationLabelsOptions model
				updateApiGatewayOperationLabelsOptionsModel := new(aisecurityforappsv1.UpdateApiGatewayOperationLabelsOptions)
				updateApiGatewayOperationLabelsOptionsModel.Selector = apiGatewayOperationsLabelsBulkInputSelectorModel
				updateApiGatewayOperationLabelsOptionsModel.User = apiGatewayOperationsLabelsBulkInputUserModel
				updateApiGatewayOperationLabelsOptionsModel.Managed = apiGatewayOperationsLabelsBulkInputManagedModel
				updateApiGatewayOperationLabelsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := aiSecurityForAppsService.UpdateApiGatewayOperationLabelsWithContext(ctx, updateApiGatewayOperationLabelsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				aiSecurityForAppsService.DisableRetries()
				result, response, operationErr := aiSecurityForAppsService.UpdateApiGatewayOperationLabels(updateApiGatewayOperationLabelsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = aiSecurityForAppsService.UpdateApiGatewayOperationLabelsWithContext(ctx, updateApiGatewayOperationLabelsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(updateApiGatewayOperationLabelsPath))
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
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": [{"operation_id": "f174e90a-fafe-4643-bbbc-4a0ed4fc8415", "labels": [{"name": "cf-llm", "source": "managed"}]}]}`)
				}))
			})
			It(`Invoke UpdateApiGatewayOperationLabels successfully`, func() {
				aiSecurityForAppsService, serviceErr := aisecurityforappsv1.NewAiSecurityForAppsV1(&aisecurityforappsv1.AiSecurityForAppsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(aiSecurityForAppsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := aiSecurityForAppsService.UpdateApiGatewayOperationLabels(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ApiGatewayOperationsLabelsBulkInputSelectorInclude model
				apiGatewayOperationsLabelsBulkInputSelectorIncludeModel := new(aisecurityforappsv1.ApiGatewayOperationsLabelsBulkInputSelectorInclude)
				apiGatewayOperationsLabelsBulkInputSelectorIncludeModel.OperationIds = []string{"f174e90a-fafe-4643-bbbc-4a0ed4fc8415"}

				// Construct an instance of the ApiGatewayOperationsLabelsBulkInputSelector model
				apiGatewayOperationsLabelsBulkInputSelectorModel := new(aisecurityforappsv1.ApiGatewayOperationsLabelsBulkInputSelector)
				apiGatewayOperationsLabelsBulkInputSelectorModel.Include = apiGatewayOperationsLabelsBulkInputSelectorIncludeModel

				// Construct an instance of the ApiGatewayOperationsLabelsBulkInputUser model
				apiGatewayOperationsLabelsBulkInputUserModel := new(aisecurityforappsv1.ApiGatewayOperationsLabelsBulkInputUser)
				apiGatewayOperationsLabelsBulkInputUserModel.Labels = []string{"testString"}

				// Construct an instance of the ApiGatewayOperationsLabelsBulkInputManaged model
				apiGatewayOperationsLabelsBulkInputManagedModel := new(aisecurityforappsv1.ApiGatewayOperationsLabelsBulkInputManaged)
				apiGatewayOperationsLabelsBulkInputManagedModel.Labels = []string{"testString"}

				// Construct an instance of the UpdateApiGatewayOperationLabelsOptions model
				updateApiGatewayOperationLabelsOptionsModel := new(aisecurityforappsv1.UpdateApiGatewayOperationLabelsOptions)
				updateApiGatewayOperationLabelsOptionsModel.Selector = apiGatewayOperationsLabelsBulkInputSelectorModel
				updateApiGatewayOperationLabelsOptionsModel.User = apiGatewayOperationsLabelsBulkInputUserModel
				updateApiGatewayOperationLabelsOptionsModel.Managed = apiGatewayOperationsLabelsBulkInputManagedModel
				updateApiGatewayOperationLabelsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = aiSecurityForAppsService.UpdateApiGatewayOperationLabels(updateApiGatewayOperationLabelsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateApiGatewayOperationLabels with error: Operation request error`, func() {
				aiSecurityForAppsService, serviceErr := aisecurityforappsv1.NewAiSecurityForAppsV1(&aisecurityforappsv1.AiSecurityForAppsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(aiSecurityForAppsService).ToNot(BeNil())

				// Construct an instance of the ApiGatewayOperationsLabelsBulkInputSelectorInclude model
				apiGatewayOperationsLabelsBulkInputSelectorIncludeModel := new(aisecurityforappsv1.ApiGatewayOperationsLabelsBulkInputSelectorInclude)
				apiGatewayOperationsLabelsBulkInputSelectorIncludeModel.OperationIds = []string{"f174e90a-fafe-4643-bbbc-4a0ed4fc8415"}

				// Construct an instance of the ApiGatewayOperationsLabelsBulkInputSelector model
				apiGatewayOperationsLabelsBulkInputSelectorModel := new(aisecurityforappsv1.ApiGatewayOperationsLabelsBulkInputSelector)
				apiGatewayOperationsLabelsBulkInputSelectorModel.Include = apiGatewayOperationsLabelsBulkInputSelectorIncludeModel

				// Construct an instance of the ApiGatewayOperationsLabelsBulkInputUser model
				apiGatewayOperationsLabelsBulkInputUserModel := new(aisecurityforappsv1.ApiGatewayOperationsLabelsBulkInputUser)
				apiGatewayOperationsLabelsBulkInputUserModel.Labels = []string{"testString"}

				// Construct an instance of the ApiGatewayOperationsLabelsBulkInputManaged model
				apiGatewayOperationsLabelsBulkInputManagedModel := new(aisecurityforappsv1.ApiGatewayOperationsLabelsBulkInputManaged)
				apiGatewayOperationsLabelsBulkInputManagedModel.Labels = []string{"testString"}

				// Construct an instance of the UpdateApiGatewayOperationLabelsOptions model
				updateApiGatewayOperationLabelsOptionsModel := new(aisecurityforappsv1.UpdateApiGatewayOperationLabelsOptions)
				updateApiGatewayOperationLabelsOptionsModel.Selector = apiGatewayOperationsLabelsBulkInputSelectorModel
				updateApiGatewayOperationLabelsOptionsModel.User = apiGatewayOperationsLabelsBulkInputUserModel
				updateApiGatewayOperationLabelsOptionsModel.Managed = apiGatewayOperationsLabelsBulkInputManagedModel
				updateApiGatewayOperationLabelsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := aiSecurityForAppsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := aiSecurityForAppsService.UpdateApiGatewayOperationLabels(updateApiGatewayOperationLabelsOptionsModel)
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
			It(`Invoke UpdateApiGatewayOperationLabels successfully`, func() {
				aiSecurityForAppsService, serviceErr := aisecurityforappsv1.NewAiSecurityForAppsV1(&aisecurityforappsv1.AiSecurityForAppsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(aiSecurityForAppsService).ToNot(BeNil())

				// Construct an instance of the ApiGatewayOperationsLabelsBulkInputSelectorInclude model
				apiGatewayOperationsLabelsBulkInputSelectorIncludeModel := new(aisecurityforappsv1.ApiGatewayOperationsLabelsBulkInputSelectorInclude)
				apiGatewayOperationsLabelsBulkInputSelectorIncludeModel.OperationIds = []string{"f174e90a-fafe-4643-bbbc-4a0ed4fc8415"}

				// Construct an instance of the ApiGatewayOperationsLabelsBulkInputSelector model
				apiGatewayOperationsLabelsBulkInputSelectorModel := new(aisecurityforappsv1.ApiGatewayOperationsLabelsBulkInputSelector)
				apiGatewayOperationsLabelsBulkInputSelectorModel.Include = apiGatewayOperationsLabelsBulkInputSelectorIncludeModel

				// Construct an instance of the ApiGatewayOperationsLabelsBulkInputUser model
				apiGatewayOperationsLabelsBulkInputUserModel := new(aisecurityforappsv1.ApiGatewayOperationsLabelsBulkInputUser)
				apiGatewayOperationsLabelsBulkInputUserModel.Labels = []string{"testString"}

				// Construct an instance of the ApiGatewayOperationsLabelsBulkInputManaged model
				apiGatewayOperationsLabelsBulkInputManagedModel := new(aisecurityforappsv1.ApiGatewayOperationsLabelsBulkInputManaged)
				apiGatewayOperationsLabelsBulkInputManagedModel.Labels = []string{"testString"}

				// Construct an instance of the UpdateApiGatewayOperationLabelsOptions model
				updateApiGatewayOperationLabelsOptionsModel := new(aisecurityforappsv1.UpdateApiGatewayOperationLabelsOptions)
				updateApiGatewayOperationLabelsOptionsModel.Selector = apiGatewayOperationsLabelsBulkInputSelectorModel
				updateApiGatewayOperationLabelsOptionsModel.User = apiGatewayOperationsLabelsBulkInputUserModel
				updateApiGatewayOperationLabelsOptionsModel.Managed = apiGatewayOperationsLabelsBulkInputManagedModel
				updateApiGatewayOperationLabelsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := aiSecurityForAppsService.UpdateApiGatewayOperationLabels(updateApiGatewayOperationLabelsOptionsModel)
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
	Describe(`CreateApiGatewayOperationLabels(createApiGatewayOperationLabelsOptions *CreateApiGatewayOperationLabelsOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		createApiGatewayOperationLabelsPath := "/v1/testString/zones/testString/api_gateway/operations/testString/labels"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createApiGatewayOperationLabelsPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateApiGatewayOperationLabels with error: Operation response processing error`, func() {
				aiSecurityForAppsService, serviceErr := aisecurityforappsv1.NewAiSecurityForAppsV1(&aisecurityforappsv1.AiSecurityForAppsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(aiSecurityForAppsService).ToNot(BeNil())

				// Construct an instance of the CreateApiGatewayOperationLabelsOptions model
				createApiGatewayOperationLabelsOptionsModel := new(aisecurityforappsv1.CreateApiGatewayOperationLabelsOptions)
				createApiGatewayOperationLabelsOptionsModel.OperationID = core.StringPtr("testString")
				createApiGatewayOperationLabelsOptionsModel.User = []string{}
				createApiGatewayOperationLabelsOptionsModel.Managed = []string{"cf-llm"}
				createApiGatewayOperationLabelsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := aiSecurityForAppsService.CreateApiGatewayOperationLabels(createApiGatewayOperationLabelsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				aiSecurityForAppsService.EnableRetries(0, 0)
				result, response, operationErr = aiSecurityForAppsService.CreateApiGatewayOperationLabels(createApiGatewayOperationLabelsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateApiGatewayOperationLabels(createApiGatewayOperationLabelsOptions *CreateApiGatewayOperationLabelsOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		createApiGatewayOperationLabelsPath := "/v1/testString/zones/testString/api_gateway/operations/testString/labels"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createApiGatewayOperationLabelsPath))
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
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"operation_id": "f174e90a-fafe-4643-bbbc-4a0ed4fc8415", "labels": [{"name": "cf-llm", "source": "managed"}]}}`)
				}))
			})
			It(`Invoke CreateApiGatewayOperationLabels successfully with retries`, func() {
				aiSecurityForAppsService, serviceErr := aisecurityforappsv1.NewAiSecurityForAppsV1(&aisecurityforappsv1.AiSecurityForAppsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(aiSecurityForAppsService).ToNot(BeNil())
				aiSecurityForAppsService.EnableRetries(0, 0)

				// Construct an instance of the CreateApiGatewayOperationLabelsOptions model
				createApiGatewayOperationLabelsOptionsModel := new(aisecurityforappsv1.CreateApiGatewayOperationLabelsOptions)
				createApiGatewayOperationLabelsOptionsModel.OperationID = core.StringPtr("testString")
				createApiGatewayOperationLabelsOptionsModel.User = []string{}
				createApiGatewayOperationLabelsOptionsModel.Managed = []string{"cf-llm"}
				createApiGatewayOperationLabelsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := aiSecurityForAppsService.CreateApiGatewayOperationLabelsWithContext(ctx, createApiGatewayOperationLabelsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				aiSecurityForAppsService.DisableRetries()
				result, response, operationErr := aiSecurityForAppsService.CreateApiGatewayOperationLabels(createApiGatewayOperationLabelsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = aiSecurityForAppsService.CreateApiGatewayOperationLabelsWithContext(ctx, createApiGatewayOperationLabelsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createApiGatewayOperationLabelsPath))
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
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"operation_id": "f174e90a-fafe-4643-bbbc-4a0ed4fc8415", "labels": [{"name": "cf-llm", "source": "managed"}]}}`)
				}))
			})
			It(`Invoke CreateApiGatewayOperationLabels successfully`, func() {
				aiSecurityForAppsService, serviceErr := aisecurityforappsv1.NewAiSecurityForAppsV1(&aisecurityforappsv1.AiSecurityForAppsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(aiSecurityForAppsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := aiSecurityForAppsService.CreateApiGatewayOperationLabels(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateApiGatewayOperationLabelsOptions model
				createApiGatewayOperationLabelsOptionsModel := new(aisecurityforappsv1.CreateApiGatewayOperationLabelsOptions)
				createApiGatewayOperationLabelsOptionsModel.OperationID = core.StringPtr("testString")
				createApiGatewayOperationLabelsOptionsModel.User = []string{}
				createApiGatewayOperationLabelsOptionsModel.Managed = []string{"cf-llm"}
				createApiGatewayOperationLabelsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = aiSecurityForAppsService.CreateApiGatewayOperationLabels(createApiGatewayOperationLabelsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateApiGatewayOperationLabels with error: Operation validation and request error`, func() {
				aiSecurityForAppsService, serviceErr := aisecurityforappsv1.NewAiSecurityForAppsV1(&aisecurityforappsv1.AiSecurityForAppsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(aiSecurityForAppsService).ToNot(BeNil())

				// Construct an instance of the CreateApiGatewayOperationLabelsOptions model
				createApiGatewayOperationLabelsOptionsModel := new(aisecurityforappsv1.CreateApiGatewayOperationLabelsOptions)
				createApiGatewayOperationLabelsOptionsModel.OperationID = core.StringPtr("testString")
				createApiGatewayOperationLabelsOptionsModel.User = []string{}
				createApiGatewayOperationLabelsOptionsModel.Managed = []string{"cf-llm"}
				createApiGatewayOperationLabelsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := aiSecurityForAppsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := aiSecurityForAppsService.CreateApiGatewayOperationLabels(createApiGatewayOperationLabelsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateApiGatewayOperationLabelsOptions model with no property values
				createApiGatewayOperationLabelsOptionsModelNew := new(aisecurityforappsv1.CreateApiGatewayOperationLabelsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = aiSecurityForAppsService.CreateApiGatewayOperationLabels(createApiGatewayOperationLabelsOptionsModelNew)
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
			It(`Invoke CreateApiGatewayOperationLabels successfully`, func() {
				aiSecurityForAppsService, serviceErr := aisecurityforappsv1.NewAiSecurityForAppsV1(&aisecurityforappsv1.AiSecurityForAppsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(aiSecurityForAppsService).ToNot(BeNil())

				// Construct an instance of the CreateApiGatewayOperationLabelsOptions model
				createApiGatewayOperationLabelsOptionsModel := new(aisecurityforappsv1.CreateApiGatewayOperationLabelsOptions)
				createApiGatewayOperationLabelsOptionsModel.OperationID = core.StringPtr("testString")
				createApiGatewayOperationLabelsOptionsModel.User = []string{}
				createApiGatewayOperationLabelsOptionsModel.Managed = []string{"cf-llm"}
				createApiGatewayOperationLabelsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := aiSecurityForAppsService.CreateApiGatewayOperationLabels(createApiGatewayOperationLabelsOptionsModel)
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
	Describe(`SetZoneApiGatewayOperationLabels(setZoneApiGatewayOperationLabelsOptions *SetZoneApiGatewayOperationLabelsOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		setZoneApiGatewayOperationLabelsPath := "/v1/testString/zones/testString/api_gateway/operations/testString/labels"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(setZoneApiGatewayOperationLabelsPath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke SetZoneApiGatewayOperationLabels with error: Operation response processing error`, func() {
				aiSecurityForAppsService, serviceErr := aisecurityforappsv1.NewAiSecurityForAppsV1(&aisecurityforappsv1.AiSecurityForAppsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(aiSecurityForAppsService).ToNot(BeNil())

				// Construct an instance of the SetZoneApiGatewayOperationLabelsOptions model
				setZoneApiGatewayOperationLabelsOptionsModel := new(aisecurityforappsv1.SetZoneApiGatewayOperationLabelsOptions)
				setZoneApiGatewayOperationLabelsOptionsModel.OperationID = core.StringPtr("testString")
				setZoneApiGatewayOperationLabelsOptionsModel.User = []string{}
				setZoneApiGatewayOperationLabelsOptionsModel.Managed = []string{"cf-llm"}
				setZoneApiGatewayOperationLabelsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := aiSecurityForAppsService.SetZoneApiGatewayOperationLabels(setZoneApiGatewayOperationLabelsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				aiSecurityForAppsService.EnableRetries(0, 0)
				result, response, operationErr = aiSecurityForAppsService.SetZoneApiGatewayOperationLabels(setZoneApiGatewayOperationLabelsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`SetZoneApiGatewayOperationLabels(setZoneApiGatewayOperationLabelsOptions *SetZoneApiGatewayOperationLabelsOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		setZoneApiGatewayOperationLabelsPath := "/v1/testString/zones/testString/api_gateway/operations/testString/labels"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(setZoneApiGatewayOperationLabelsPath))
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
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"operation_id": "f174e90a-fafe-4643-bbbc-4a0ed4fc8415", "labels": [{"name": "cf-llm", "source": "managed"}]}}`)
				}))
			})
			It(`Invoke SetZoneApiGatewayOperationLabels successfully with retries`, func() {
				aiSecurityForAppsService, serviceErr := aisecurityforappsv1.NewAiSecurityForAppsV1(&aisecurityforappsv1.AiSecurityForAppsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(aiSecurityForAppsService).ToNot(BeNil())
				aiSecurityForAppsService.EnableRetries(0, 0)

				// Construct an instance of the SetZoneApiGatewayOperationLabelsOptions model
				setZoneApiGatewayOperationLabelsOptionsModel := new(aisecurityforappsv1.SetZoneApiGatewayOperationLabelsOptions)
				setZoneApiGatewayOperationLabelsOptionsModel.OperationID = core.StringPtr("testString")
				setZoneApiGatewayOperationLabelsOptionsModel.User = []string{}
				setZoneApiGatewayOperationLabelsOptionsModel.Managed = []string{"cf-llm"}
				setZoneApiGatewayOperationLabelsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := aiSecurityForAppsService.SetZoneApiGatewayOperationLabelsWithContext(ctx, setZoneApiGatewayOperationLabelsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				aiSecurityForAppsService.DisableRetries()
				result, response, operationErr := aiSecurityForAppsService.SetZoneApiGatewayOperationLabels(setZoneApiGatewayOperationLabelsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = aiSecurityForAppsService.SetZoneApiGatewayOperationLabelsWithContext(ctx, setZoneApiGatewayOperationLabelsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(setZoneApiGatewayOperationLabelsPath))
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
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"operation_id": "f174e90a-fafe-4643-bbbc-4a0ed4fc8415", "labels": [{"name": "cf-llm", "source": "managed"}]}}`)
				}))
			})
			It(`Invoke SetZoneApiGatewayOperationLabels successfully`, func() {
				aiSecurityForAppsService, serviceErr := aisecurityforappsv1.NewAiSecurityForAppsV1(&aisecurityforappsv1.AiSecurityForAppsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(aiSecurityForAppsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := aiSecurityForAppsService.SetZoneApiGatewayOperationLabels(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the SetZoneApiGatewayOperationLabelsOptions model
				setZoneApiGatewayOperationLabelsOptionsModel := new(aisecurityforappsv1.SetZoneApiGatewayOperationLabelsOptions)
				setZoneApiGatewayOperationLabelsOptionsModel.OperationID = core.StringPtr("testString")
				setZoneApiGatewayOperationLabelsOptionsModel.User = []string{}
				setZoneApiGatewayOperationLabelsOptionsModel.Managed = []string{"cf-llm"}
				setZoneApiGatewayOperationLabelsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = aiSecurityForAppsService.SetZoneApiGatewayOperationLabels(setZoneApiGatewayOperationLabelsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke SetZoneApiGatewayOperationLabels with error: Operation validation and request error`, func() {
				aiSecurityForAppsService, serviceErr := aisecurityforappsv1.NewAiSecurityForAppsV1(&aisecurityforappsv1.AiSecurityForAppsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(aiSecurityForAppsService).ToNot(BeNil())

				// Construct an instance of the SetZoneApiGatewayOperationLabelsOptions model
				setZoneApiGatewayOperationLabelsOptionsModel := new(aisecurityforappsv1.SetZoneApiGatewayOperationLabelsOptions)
				setZoneApiGatewayOperationLabelsOptionsModel.OperationID = core.StringPtr("testString")
				setZoneApiGatewayOperationLabelsOptionsModel.User = []string{}
				setZoneApiGatewayOperationLabelsOptionsModel.Managed = []string{"cf-llm"}
				setZoneApiGatewayOperationLabelsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := aiSecurityForAppsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := aiSecurityForAppsService.SetZoneApiGatewayOperationLabels(setZoneApiGatewayOperationLabelsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the SetZoneApiGatewayOperationLabelsOptions model with no property values
				setZoneApiGatewayOperationLabelsOptionsModelNew := new(aisecurityforappsv1.SetZoneApiGatewayOperationLabelsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = aiSecurityForAppsService.SetZoneApiGatewayOperationLabels(setZoneApiGatewayOperationLabelsOptionsModelNew)
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
			It(`Invoke SetZoneApiGatewayOperationLabels successfully`, func() {
				aiSecurityForAppsService, serviceErr := aisecurityforappsv1.NewAiSecurityForAppsV1(&aisecurityforappsv1.AiSecurityForAppsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(aiSecurityForAppsService).ToNot(BeNil())

				// Construct an instance of the SetZoneApiGatewayOperationLabelsOptions model
				setZoneApiGatewayOperationLabelsOptionsModel := new(aisecurityforappsv1.SetZoneApiGatewayOperationLabelsOptions)
				setZoneApiGatewayOperationLabelsOptionsModel.OperationID = core.StringPtr("testString")
				setZoneApiGatewayOperationLabelsOptionsModel.User = []string{}
				setZoneApiGatewayOperationLabelsOptionsModel.Managed = []string{"cf-llm"}
				setZoneApiGatewayOperationLabelsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := aiSecurityForAppsService.SetZoneApiGatewayOperationLabels(setZoneApiGatewayOperationLabelsOptionsModel)
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
	Describe(`UnsetZoneApiGatewayOperationLabels(unsetZoneApiGatewayOperationLabelsOptions *UnsetZoneApiGatewayOperationLabelsOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		unsetZoneApiGatewayOperationLabelsPath := "/v1/testString/zones/testString/api_gateway/operations/testString/labels"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(unsetZoneApiGatewayOperationLabelsPath))
					Expect(req.Method).To(Equal("DELETE"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UnsetZoneApiGatewayOperationLabels with error: Operation response processing error`, func() {
				aiSecurityForAppsService, serviceErr := aisecurityforappsv1.NewAiSecurityForAppsV1(&aisecurityforappsv1.AiSecurityForAppsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(aiSecurityForAppsService).ToNot(BeNil())

				// Construct an instance of the UnsetZoneApiGatewayOperationLabelsOptions model
				unsetZoneApiGatewayOperationLabelsOptionsModel := new(aisecurityforappsv1.UnsetZoneApiGatewayOperationLabelsOptions)
				unsetZoneApiGatewayOperationLabelsOptionsModel.OperationID = core.StringPtr("testString")
				unsetZoneApiGatewayOperationLabelsOptionsModel.User = []string{"testString"}
				unsetZoneApiGatewayOperationLabelsOptionsModel.Managed = []string{"testString"}
				unsetZoneApiGatewayOperationLabelsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := aiSecurityForAppsService.UnsetZoneApiGatewayOperationLabels(unsetZoneApiGatewayOperationLabelsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				aiSecurityForAppsService.EnableRetries(0, 0)
				result, response, operationErr = aiSecurityForAppsService.UnsetZoneApiGatewayOperationLabels(unsetZoneApiGatewayOperationLabelsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UnsetZoneApiGatewayOperationLabels(unsetZoneApiGatewayOperationLabelsOptions *UnsetZoneApiGatewayOperationLabelsOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		unsetZoneApiGatewayOperationLabelsPath := "/v1/testString/zones/testString/api_gateway/operations/testString/labels"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(unsetZoneApiGatewayOperationLabelsPath))
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
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": "anyValue"}`)
				}))
			})
			It(`Invoke UnsetZoneApiGatewayOperationLabels successfully with retries`, func() {
				aiSecurityForAppsService, serviceErr := aisecurityforappsv1.NewAiSecurityForAppsV1(&aisecurityforappsv1.AiSecurityForAppsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(aiSecurityForAppsService).ToNot(BeNil())
				aiSecurityForAppsService.EnableRetries(0, 0)

				// Construct an instance of the UnsetZoneApiGatewayOperationLabelsOptions model
				unsetZoneApiGatewayOperationLabelsOptionsModel := new(aisecurityforappsv1.UnsetZoneApiGatewayOperationLabelsOptions)
				unsetZoneApiGatewayOperationLabelsOptionsModel.OperationID = core.StringPtr("testString")
				unsetZoneApiGatewayOperationLabelsOptionsModel.User = []string{"testString"}
				unsetZoneApiGatewayOperationLabelsOptionsModel.Managed = []string{"testString"}
				unsetZoneApiGatewayOperationLabelsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := aiSecurityForAppsService.UnsetZoneApiGatewayOperationLabelsWithContext(ctx, unsetZoneApiGatewayOperationLabelsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				aiSecurityForAppsService.DisableRetries()
				result, response, operationErr := aiSecurityForAppsService.UnsetZoneApiGatewayOperationLabels(unsetZoneApiGatewayOperationLabelsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = aiSecurityForAppsService.UnsetZoneApiGatewayOperationLabelsWithContext(ctx, unsetZoneApiGatewayOperationLabelsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(unsetZoneApiGatewayOperationLabelsPath))
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
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": "anyValue"}`)
				}))
			})
			It(`Invoke UnsetZoneApiGatewayOperationLabels successfully`, func() {
				aiSecurityForAppsService, serviceErr := aisecurityforappsv1.NewAiSecurityForAppsV1(&aisecurityforappsv1.AiSecurityForAppsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(aiSecurityForAppsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := aiSecurityForAppsService.UnsetZoneApiGatewayOperationLabels(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UnsetZoneApiGatewayOperationLabelsOptions model
				unsetZoneApiGatewayOperationLabelsOptionsModel := new(aisecurityforappsv1.UnsetZoneApiGatewayOperationLabelsOptions)
				unsetZoneApiGatewayOperationLabelsOptionsModel.OperationID = core.StringPtr("testString")
				unsetZoneApiGatewayOperationLabelsOptionsModel.User = []string{"testString"}
				unsetZoneApiGatewayOperationLabelsOptionsModel.Managed = []string{"testString"}
				unsetZoneApiGatewayOperationLabelsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = aiSecurityForAppsService.UnsetZoneApiGatewayOperationLabels(unsetZoneApiGatewayOperationLabelsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UnsetZoneApiGatewayOperationLabels with error: Operation validation and request error`, func() {
				aiSecurityForAppsService, serviceErr := aisecurityforappsv1.NewAiSecurityForAppsV1(&aisecurityforappsv1.AiSecurityForAppsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(aiSecurityForAppsService).ToNot(BeNil())

				// Construct an instance of the UnsetZoneApiGatewayOperationLabelsOptions model
				unsetZoneApiGatewayOperationLabelsOptionsModel := new(aisecurityforappsv1.UnsetZoneApiGatewayOperationLabelsOptions)
				unsetZoneApiGatewayOperationLabelsOptionsModel.OperationID = core.StringPtr("testString")
				unsetZoneApiGatewayOperationLabelsOptionsModel.User = []string{"testString"}
				unsetZoneApiGatewayOperationLabelsOptionsModel.Managed = []string{"testString"}
				unsetZoneApiGatewayOperationLabelsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := aiSecurityForAppsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := aiSecurityForAppsService.UnsetZoneApiGatewayOperationLabels(unsetZoneApiGatewayOperationLabelsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UnsetZoneApiGatewayOperationLabelsOptions model with no property values
				unsetZoneApiGatewayOperationLabelsOptionsModelNew := new(aisecurityforappsv1.UnsetZoneApiGatewayOperationLabelsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = aiSecurityForAppsService.UnsetZoneApiGatewayOperationLabels(unsetZoneApiGatewayOperationLabelsOptionsModelNew)
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
			It(`Invoke UnsetZoneApiGatewayOperationLabels successfully`, func() {
				aiSecurityForAppsService, serviceErr := aisecurityforappsv1.NewAiSecurityForAppsV1(&aisecurityforappsv1.AiSecurityForAppsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(aiSecurityForAppsService).ToNot(BeNil())

				// Construct an instance of the UnsetZoneApiGatewayOperationLabelsOptions model
				unsetZoneApiGatewayOperationLabelsOptionsModel := new(aisecurityforappsv1.UnsetZoneApiGatewayOperationLabelsOptions)
				unsetZoneApiGatewayOperationLabelsOptionsModel.OperationID = core.StringPtr("testString")
				unsetZoneApiGatewayOperationLabelsOptionsModel.User = []string{"testString"}
				unsetZoneApiGatewayOperationLabelsOptionsModel.Managed = []string{"testString"}
				unsetZoneApiGatewayOperationLabelsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := aiSecurityForAppsService.UnsetZoneApiGatewayOperationLabels(unsetZoneApiGatewayOperationLabelsOptionsModel)
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
	Describe(`GetZoneApiGatewayOperation(getZoneApiGatewayOperationOptions *GetZoneApiGatewayOperationOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getZoneApiGatewayOperationPath := "/v1/testString/zones/testString/api_gateway/operations/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getZoneApiGatewayOperationPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetZoneApiGatewayOperation with error: Operation response processing error`, func() {
				aiSecurityForAppsService, serviceErr := aisecurityforappsv1.NewAiSecurityForAppsV1(&aisecurityforappsv1.AiSecurityForAppsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(aiSecurityForAppsService).ToNot(BeNil())

				// Construct an instance of the GetZoneApiGatewayOperationOptions model
				getZoneApiGatewayOperationOptionsModel := new(aisecurityforappsv1.GetZoneApiGatewayOperationOptions)
				getZoneApiGatewayOperationOptionsModel.OperationID = core.StringPtr("testString")
				getZoneApiGatewayOperationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := aiSecurityForAppsService.GetZoneApiGatewayOperation(getZoneApiGatewayOperationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				aiSecurityForAppsService.EnableRetries(0, 0)
				result, response, operationErr = aiSecurityForAppsService.GetZoneApiGatewayOperation(getZoneApiGatewayOperationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetZoneApiGatewayOperation(getZoneApiGatewayOperationOptions *GetZoneApiGatewayOperationOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getZoneApiGatewayOperationPath := "/v1/testString/zones/testString/api_gateway/operations/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getZoneApiGatewayOperationPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"operation_id": "f174e90a-fafe-4643-bbbc-4a0ed4fc8415", "method": "POST", "host": "api.example.com", "endpoint": "/v1/messages"}}`)
				}))
			})
			It(`Invoke GetZoneApiGatewayOperation successfully with retries`, func() {
				aiSecurityForAppsService, serviceErr := aisecurityforappsv1.NewAiSecurityForAppsV1(&aisecurityforappsv1.AiSecurityForAppsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(aiSecurityForAppsService).ToNot(BeNil())
				aiSecurityForAppsService.EnableRetries(0, 0)

				// Construct an instance of the GetZoneApiGatewayOperationOptions model
				getZoneApiGatewayOperationOptionsModel := new(aisecurityforappsv1.GetZoneApiGatewayOperationOptions)
				getZoneApiGatewayOperationOptionsModel.OperationID = core.StringPtr("testString")
				getZoneApiGatewayOperationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := aiSecurityForAppsService.GetZoneApiGatewayOperationWithContext(ctx, getZoneApiGatewayOperationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				aiSecurityForAppsService.DisableRetries()
				result, response, operationErr := aiSecurityForAppsService.GetZoneApiGatewayOperation(getZoneApiGatewayOperationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = aiSecurityForAppsService.GetZoneApiGatewayOperationWithContext(ctx, getZoneApiGatewayOperationOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getZoneApiGatewayOperationPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"operation_id": "f174e90a-fafe-4643-bbbc-4a0ed4fc8415", "method": "POST", "host": "api.example.com", "endpoint": "/v1/messages"}}`)
				}))
			})
			It(`Invoke GetZoneApiGatewayOperation successfully`, func() {
				aiSecurityForAppsService, serviceErr := aisecurityforappsv1.NewAiSecurityForAppsV1(&aisecurityforappsv1.AiSecurityForAppsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(aiSecurityForAppsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := aiSecurityForAppsService.GetZoneApiGatewayOperation(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetZoneApiGatewayOperationOptions model
				getZoneApiGatewayOperationOptionsModel := new(aisecurityforappsv1.GetZoneApiGatewayOperationOptions)
				getZoneApiGatewayOperationOptionsModel.OperationID = core.StringPtr("testString")
				getZoneApiGatewayOperationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = aiSecurityForAppsService.GetZoneApiGatewayOperation(getZoneApiGatewayOperationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetZoneApiGatewayOperation with error: Operation validation and request error`, func() {
				aiSecurityForAppsService, serviceErr := aisecurityforappsv1.NewAiSecurityForAppsV1(&aisecurityforappsv1.AiSecurityForAppsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(aiSecurityForAppsService).ToNot(BeNil())

				// Construct an instance of the GetZoneApiGatewayOperationOptions model
				getZoneApiGatewayOperationOptionsModel := new(aisecurityforappsv1.GetZoneApiGatewayOperationOptions)
				getZoneApiGatewayOperationOptionsModel.OperationID = core.StringPtr("testString")
				getZoneApiGatewayOperationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := aiSecurityForAppsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := aiSecurityForAppsService.GetZoneApiGatewayOperation(getZoneApiGatewayOperationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetZoneApiGatewayOperationOptions model with no property values
				getZoneApiGatewayOperationOptionsModelNew := new(aisecurityforappsv1.GetZoneApiGatewayOperationOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = aiSecurityForAppsService.GetZoneApiGatewayOperation(getZoneApiGatewayOperationOptionsModelNew)
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
			It(`Invoke GetZoneApiGatewayOperation successfully`, func() {
				aiSecurityForAppsService, serviceErr := aisecurityforappsv1.NewAiSecurityForAppsV1(&aisecurityforappsv1.AiSecurityForAppsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(aiSecurityForAppsService).ToNot(BeNil())

				// Construct an instance of the GetZoneApiGatewayOperationOptions model
				getZoneApiGatewayOperationOptionsModel := new(aisecurityforappsv1.GetZoneApiGatewayOperationOptions)
				getZoneApiGatewayOperationOptionsModel.OperationID = core.StringPtr("testString")
				getZoneApiGatewayOperationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := aiSecurityForAppsService.GetZoneApiGatewayOperation(getZoneApiGatewayOperationOptionsModel)
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
	Describe(`DeleteZoneApiGatewayOperation(deleteZoneApiGatewayOperationOptions *DeleteZoneApiGatewayOperationOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		deleteZoneApiGatewayOperationPath := "/v1/testString/zones/testString/api_gateway/operations/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteZoneApiGatewayOperationPath))
					Expect(req.Method).To(Equal("DELETE"))

					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteZoneApiGatewayOperation successfully`, func() {
				aiSecurityForAppsService, serviceErr := aisecurityforappsv1.NewAiSecurityForAppsV1(&aisecurityforappsv1.AiSecurityForAppsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(aiSecurityForAppsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := aiSecurityForAppsService.DeleteZoneApiGatewayOperation(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteZoneApiGatewayOperationOptions model
				deleteZoneApiGatewayOperationOptionsModel := new(aisecurityforappsv1.DeleteZoneApiGatewayOperationOptions)
				deleteZoneApiGatewayOperationOptionsModel.OperationID = core.StringPtr("testString")
				deleteZoneApiGatewayOperationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = aiSecurityForAppsService.DeleteZoneApiGatewayOperation(deleteZoneApiGatewayOperationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteZoneApiGatewayOperation with error: Operation validation and request error`, func() {
				aiSecurityForAppsService, serviceErr := aisecurityforappsv1.NewAiSecurityForAppsV1(&aisecurityforappsv1.AiSecurityForAppsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(aiSecurityForAppsService).ToNot(BeNil())

				// Construct an instance of the DeleteZoneApiGatewayOperationOptions model
				deleteZoneApiGatewayOperationOptionsModel := new(aisecurityforappsv1.DeleteZoneApiGatewayOperationOptions)
				deleteZoneApiGatewayOperationOptionsModel.OperationID = core.StringPtr("testString")
				deleteZoneApiGatewayOperationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := aiSecurityForAppsService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := aiSecurityForAppsService.DeleteZoneApiGatewayOperation(deleteZoneApiGatewayOperationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteZoneApiGatewayOperationOptions model with no property values
				deleteZoneApiGatewayOperationOptionsModelNew := new(aisecurityforappsv1.DeleteZoneApiGatewayOperationOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = aiSecurityForAppsService.DeleteZoneApiGatewayOperation(deleteZoneApiGatewayOperationOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetApiGatewaySchemas(getApiGatewaySchemasOptions *GetApiGatewaySchemasOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getApiGatewaySchemasPath := "/v1/testString/zones/testString/api_gateway/schemas"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getApiGatewaySchemasPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetApiGatewaySchemas with error: Operation response processing error`, func() {
				aiSecurityForAppsService, serviceErr := aisecurityforappsv1.NewAiSecurityForAppsV1(&aisecurityforappsv1.AiSecurityForAppsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(aiSecurityForAppsService).ToNot(BeNil())

				// Construct an instance of the GetApiGatewaySchemasOptions model
				getApiGatewaySchemasOptionsModel := new(aisecurityforappsv1.GetApiGatewaySchemasOptions)
				getApiGatewaySchemasOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := aiSecurityForAppsService.GetApiGatewaySchemas(getApiGatewaySchemasOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				aiSecurityForAppsService.EnableRetries(0, 0)
				result, response, operationErr = aiSecurityForAppsService.GetApiGatewaySchemas(getApiGatewaySchemasOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetApiGatewaySchemas(getApiGatewaySchemasOptions *GetApiGatewaySchemasOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getApiGatewaySchemasPath := "/v1/testString/zones/testString/api_gateway/schemas"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getApiGatewaySchemasPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"anyKey": "anyValue"}}`)
				}))
			})
			It(`Invoke GetApiGatewaySchemas successfully with retries`, func() {
				aiSecurityForAppsService, serviceErr := aisecurityforappsv1.NewAiSecurityForAppsV1(&aisecurityforappsv1.AiSecurityForAppsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(aiSecurityForAppsService).ToNot(BeNil())
				aiSecurityForAppsService.EnableRetries(0, 0)

				// Construct an instance of the GetApiGatewaySchemasOptions model
				getApiGatewaySchemasOptionsModel := new(aisecurityforappsv1.GetApiGatewaySchemasOptions)
				getApiGatewaySchemasOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := aiSecurityForAppsService.GetApiGatewaySchemasWithContext(ctx, getApiGatewaySchemasOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				aiSecurityForAppsService.DisableRetries()
				result, response, operationErr := aiSecurityForAppsService.GetApiGatewaySchemas(getApiGatewaySchemasOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = aiSecurityForAppsService.GetApiGatewaySchemasWithContext(ctx, getApiGatewaySchemasOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getApiGatewaySchemasPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"anyKey": "anyValue"}}`)
				}))
			})
			It(`Invoke GetApiGatewaySchemas successfully`, func() {
				aiSecurityForAppsService, serviceErr := aisecurityforappsv1.NewAiSecurityForAppsV1(&aisecurityforappsv1.AiSecurityForAppsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(aiSecurityForAppsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := aiSecurityForAppsService.GetApiGatewaySchemas(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetApiGatewaySchemasOptions model
				getApiGatewaySchemasOptionsModel := new(aisecurityforappsv1.GetApiGatewaySchemasOptions)
				getApiGatewaySchemasOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = aiSecurityForAppsService.GetApiGatewaySchemas(getApiGatewaySchemasOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetApiGatewaySchemas with error: Operation request error`, func() {
				aiSecurityForAppsService, serviceErr := aisecurityforappsv1.NewAiSecurityForAppsV1(&aisecurityforappsv1.AiSecurityForAppsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(aiSecurityForAppsService).ToNot(BeNil())

				// Construct an instance of the GetApiGatewaySchemasOptions model
				getApiGatewaySchemasOptionsModel := new(aisecurityforappsv1.GetApiGatewaySchemasOptions)
				getApiGatewaySchemasOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := aiSecurityForAppsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := aiSecurityForAppsService.GetApiGatewaySchemas(getApiGatewaySchemasOptionsModel)
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
			It(`Invoke GetApiGatewaySchemas successfully`, func() {
				aiSecurityForAppsService, serviceErr := aisecurityforappsv1.NewAiSecurityForAppsV1(&aisecurityforappsv1.AiSecurityForAppsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(aiSecurityForAppsService).ToNot(BeNil())

				// Construct an instance of the GetApiGatewaySchemasOptions model
				getApiGatewaySchemasOptionsModel := new(aisecurityforappsv1.GetApiGatewaySchemasOptions)
				getApiGatewaySchemasOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := aiSecurityForAppsService.GetApiGatewaySchemas(getApiGatewaySchemasOptionsModel)
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
			aiSecurityForAppsService, _ := aisecurityforappsv1.NewAiSecurityForAppsV1(&aisecurityforappsv1.AiSecurityForAppsV1Options{
				URL:            "http://aisecurityforappsv1modelgenerator.com",
				Authenticator:  &core.NoAuthAuthenticator{},
				Crn:            core.StringPtr(crn),
				ZoneIdentifier: core.StringPtr(zoneIdentifier),
			})
			It(`Invoke NewApiGatewayOperationsLabelsBulkDeleteInputSelector successfully`, func() {
				var include *aisecurityforappsv1.ApiGatewayOperationsLabelsBulkDeleteInputSelectorInclude = nil
				_, err := aiSecurityForAppsService.NewApiGatewayOperationsLabelsBulkDeleteInputSelector(include)
				Expect(err).ToNot(BeNil())
			})
			It(`Invoke NewApiGatewayOperationsLabelsBulkInputManaged successfully`, func() {
				labels := []string{"testString"}
				_model, err := aiSecurityForAppsService.NewApiGatewayOperationsLabelsBulkInputManaged(labels)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewApiGatewayOperationsLabelsBulkInputSelector successfully`, func() {
				var include *aisecurityforappsv1.ApiGatewayOperationsLabelsBulkInputSelectorInclude = nil
				_, err := aiSecurityForAppsService.NewApiGatewayOperationsLabelsBulkInputSelector(include)
				Expect(err).ToNot(BeNil())
			})
			It(`Invoke NewApiGatewayOperationsLabelsBulkInputUser successfully`, func() {
				labels := []string{"testString"}
				_model, err := aiSecurityForAppsService.NewApiGatewayOperationsLabelsBulkInputUser(labels)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewCreateApiGatewayOperationItemOptions successfully`, func() {
				// Construct an instance of the CreateApiGatewayOperationItemOptions model
				createApiGatewayOperationItemOptionsModel := aiSecurityForAppsService.NewCreateApiGatewayOperationItemOptions()
				createApiGatewayOperationItemOptionsModel.SetMethod("POST")
				createApiGatewayOperationItemOptionsModel.SetHost("api.example.com")
				createApiGatewayOperationItemOptionsModel.SetEndpoint("/v1/messages")
				createApiGatewayOperationItemOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createApiGatewayOperationItemOptionsModel).ToNot(BeNil())
				Expect(createApiGatewayOperationItemOptionsModel.Method).To(Equal(core.StringPtr("POST")))
				Expect(createApiGatewayOperationItemOptionsModel.Host).To(Equal(core.StringPtr("api.example.com")))
				Expect(createApiGatewayOperationItemOptionsModel.Endpoint).To(Equal(core.StringPtr("/v1/messages")))
				Expect(createApiGatewayOperationItemOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateApiGatewayOperationLabelsOptions successfully`, func() {
				// Construct an instance of the CreateApiGatewayOperationLabelsOptions model
				operationID := "testString"
				createApiGatewayOperationLabelsOptionsModel := aiSecurityForAppsService.NewCreateApiGatewayOperationLabelsOptions(operationID)
				createApiGatewayOperationLabelsOptionsModel.SetOperationID("testString")
				createApiGatewayOperationLabelsOptionsModel.SetUser([]string{})
				createApiGatewayOperationLabelsOptionsModel.SetManaged([]string{"cf-llm"})
				createApiGatewayOperationLabelsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createApiGatewayOperationLabelsOptionsModel).ToNot(BeNil())
				Expect(createApiGatewayOperationLabelsOptionsModel.OperationID).To(Equal(core.StringPtr("testString")))
				Expect(createApiGatewayOperationLabelsOptionsModel.User).To(Equal([]string{}))
				Expect(createApiGatewayOperationLabelsOptionsModel.Managed).To(Equal([]string{"cf-llm"}))
				Expect(createApiGatewayOperationLabelsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateZoneApiGatewayOperationOptions successfully`, func() {
				// Construct an instance of the ApiGatewayOperation model
				apiGatewayOperationModel := new(aisecurityforappsv1.ApiGatewayOperation)
				Expect(apiGatewayOperationModel).ToNot(BeNil())
				apiGatewayOperationModel.Method = core.StringPtr("POST")
				apiGatewayOperationModel.Host = core.StringPtr("api.example.com")
				apiGatewayOperationModel.Endpoint = core.StringPtr("/v1/messages")
				Expect(apiGatewayOperationModel.Method).To(Equal(core.StringPtr("POST")))
				Expect(apiGatewayOperationModel.Host).To(Equal(core.StringPtr("api.example.com")))
				Expect(apiGatewayOperationModel.Endpoint).To(Equal(core.StringPtr("/v1/messages")))

				// Construct an instance of the CreateZoneApiGatewayOperationOptions model
				createZoneApiGatewayOperationOptionsModel := aiSecurityForAppsService.NewCreateZoneApiGatewayOperationOptions()
				createZoneApiGatewayOperationOptionsModel.SetApiGatewayOperation([]aisecurityforappsv1.ApiGatewayOperation{*apiGatewayOperationModel})
				createZoneApiGatewayOperationOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createZoneApiGatewayOperationOptionsModel).ToNot(BeNil())
				Expect(createZoneApiGatewayOperationOptionsModel.ApiGatewayOperation).To(Equal([]aisecurityforappsv1.ApiGatewayOperation{*apiGatewayOperationModel}))
				Expect(createZoneApiGatewayOperationOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteZoneApiGatewayOperationLabelsOptions successfully`, func() {
				// Construct an instance of the ApiGatewayOperationsLabelsBulkDeleteInputSelectorInclude model
				apiGatewayOperationsLabelsBulkDeleteInputSelectorIncludeModel := new(aisecurityforappsv1.ApiGatewayOperationsLabelsBulkDeleteInputSelectorInclude)
				Expect(apiGatewayOperationsLabelsBulkDeleteInputSelectorIncludeModel).ToNot(BeNil())
				apiGatewayOperationsLabelsBulkDeleteInputSelectorIncludeModel.OperationIds = []string{"f174e90a-fafe-4643-bbbc-4a0ed4fc8415"}
				Expect(apiGatewayOperationsLabelsBulkDeleteInputSelectorIncludeModel.OperationIds).To(Equal([]string{"f174e90a-fafe-4643-bbbc-4a0ed4fc8415"}))

				// Construct an instance of the ApiGatewayOperationsLabelsBulkDeleteInputSelector model
				apiGatewayOperationsLabelsBulkDeleteInputSelectorModel := new(aisecurityforappsv1.ApiGatewayOperationsLabelsBulkDeleteInputSelector)
				Expect(apiGatewayOperationsLabelsBulkDeleteInputSelectorModel).ToNot(BeNil())
				apiGatewayOperationsLabelsBulkDeleteInputSelectorModel.Include = apiGatewayOperationsLabelsBulkDeleteInputSelectorIncludeModel
				Expect(apiGatewayOperationsLabelsBulkDeleteInputSelectorModel.Include).To(Equal(apiGatewayOperationsLabelsBulkDeleteInputSelectorIncludeModel))

				// Construct an instance of the ApiGatewayOperationsLabelsBulkDeleteInputUser model
				apiGatewayOperationsLabelsBulkDeleteInputUserModel := new(aisecurityforappsv1.ApiGatewayOperationsLabelsBulkDeleteInputUser)
				Expect(apiGatewayOperationsLabelsBulkDeleteInputUserModel).ToNot(BeNil())
				apiGatewayOperationsLabelsBulkDeleteInputUserModel.Labels = []string{"testString"}
				Expect(apiGatewayOperationsLabelsBulkDeleteInputUserModel.Labels).To(Equal([]string{"testString"}))

				// Construct an instance of the ApiGatewayOperationsLabelsBulkDeleteInputManaged model
				apiGatewayOperationsLabelsBulkDeleteInputManagedModel := new(aisecurityforappsv1.ApiGatewayOperationsLabelsBulkDeleteInputManaged)
				Expect(apiGatewayOperationsLabelsBulkDeleteInputManagedModel).ToNot(BeNil())
				apiGatewayOperationsLabelsBulkDeleteInputManagedModel.Labels = []string{"testString"}
				Expect(apiGatewayOperationsLabelsBulkDeleteInputManagedModel.Labels).To(Equal([]string{"testString"}))

				// Construct an instance of the DeleteZoneApiGatewayOperationLabelsOptions model
				var deleteZoneApiGatewayOperationLabelsOptionsSelector *aisecurityforappsv1.ApiGatewayOperationsLabelsBulkDeleteInputSelector = nil
				deleteZoneApiGatewayOperationLabelsOptionsModel := aiSecurityForAppsService.NewDeleteZoneApiGatewayOperationLabelsOptions(deleteZoneApiGatewayOperationLabelsOptionsSelector)
				deleteZoneApiGatewayOperationLabelsOptionsModel.SetSelector(apiGatewayOperationsLabelsBulkDeleteInputSelectorModel)
				deleteZoneApiGatewayOperationLabelsOptionsModel.SetUser(apiGatewayOperationsLabelsBulkDeleteInputUserModel)
				deleteZoneApiGatewayOperationLabelsOptionsModel.SetManaged(apiGatewayOperationsLabelsBulkDeleteInputManagedModel)
				deleteZoneApiGatewayOperationLabelsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteZoneApiGatewayOperationLabelsOptionsModel).ToNot(BeNil())
				Expect(deleteZoneApiGatewayOperationLabelsOptionsModel.Selector).To(Equal(apiGatewayOperationsLabelsBulkDeleteInputSelectorModel))
				Expect(deleteZoneApiGatewayOperationLabelsOptionsModel.User).To(Equal(apiGatewayOperationsLabelsBulkDeleteInputUserModel))
				Expect(deleteZoneApiGatewayOperationLabelsOptionsModel.Managed).To(Equal(apiGatewayOperationsLabelsBulkDeleteInputManagedModel))
				Expect(deleteZoneApiGatewayOperationLabelsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteZoneApiGatewayOperationOptions successfully`, func() {
				// Construct an instance of the DeleteZoneApiGatewayOperationOptions model
				operationID := "testString"
				deleteZoneApiGatewayOperationOptionsModel := aiSecurityForAppsService.NewDeleteZoneApiGatewayOperationOptions(operationID)
				deleteZoneApiGatewayOperationOptionsModel.SetOperationID("testString")
				deleteZoneApiGatewayOperationOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteZoneApiGatewayOperationOptionsModel).ToNot(BeNil())
				Expect(deleteZoneApiGatewayOperationOptionsModel.OperationID).To(Equal(core.StringPtr("testString")))
				Expect(deleteZoneApiGatewayOperationOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteZoneApiGatewayOperationsOptions successfully`, func() {
				// Construct an instance of the DeleteZoneApiGatewayOperationsOptions model
				deleteZoneApiGatewayOperationsOptionsModel := aiSecurityForAppsService.NewDeleteZoneApiGatewayOperationsOptions()
				deleteZoneApiGatewayOperationsOptionsModel.SetOperationIds([]string{"f174e90a-fafe-4643-bbbc-4a0ed4fc8415"})
				deleteZoneApiGatewayOperationsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteZoneApiGatewayOperationsOptionsModel).ToNot(BeNil())
				Expect(deleteZoneApiGatewayOperationsOptionsModel.OperationIds).To(Equal([]string{"f174e90a-fafe-4643-bbbc-4a0ed4fc8415"}))
				Expect(deleteZoneApiGatewayOperationsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetAiSecuritySettingsOptions successfully`, func() {
				// Construct an instance of the GetAiSecuritySettingsOptions model
				getAiSecuritySettingsOptionsModel := aiSecurityForAppsService.NewGetAiSecuritySettingsOptions()
				getAiSecuritySettingsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getAiSecuritySettingsOptionsModel).ToNot(BeNil())
				Expect(getAiSecuritySettingsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetApiGatewayDiscoveryOptions successfully`, func() {
				// Construct an instance of the GetApiGatewayDiscoveryOptions model
				getApiGatewayDiscoveryOptionsModel := aiSecurityForAppsService.NewGetApiGatewayDiscoveryOptions()
				getApiGatewayDiscoveryOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getApiGatewayDiscoveryOptionsModel).ToNot(BeNil())
				Expect(getApiGatewayDiscoveryOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetApiGatewaySchemasOptions successfully`, func() {
				// Construct an instance of the GetApiGatewaySchemasOptions model
				getApiGatewaySchemasOptionsModel := aiSecurityForAppsService.NewGetApiGatewaySchemasOptions()
				getApiGatewaySchemasOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getApiGatewaySchemasOptionsModel).ToNot(BeNil())
				Expect(getApiGatewaySchemasOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetZoneApiGatewayOperationOptions successfully`, func() {
				// Construct an instance of the GetZoneApiGatewayOperationOptions model
				operationID := "testString"
				getZoneApiGatewayOperationOptionsModel := aiSecurityForAppsService.NewGetZoneApiGatewayOperationOptions(operationID)
				getZoneApiGatewayOperationOptionsModel.SetOperationID("testString")
				getZoneApiGatewayOperationOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getZoneApiGatewayOperationOptionsModel).ToNot(BeNil())
				Expect(getZoneApiGatewayOperationOptionsModel.OperationID).To(Equal(core.StringPtr("testString")))
				Expect(getZoneApiGatewayOperationOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListApiGatewayDiscoveryOperationsOptions successfully`, func() {
				// Construct an instance of the ListApiGatewayDiscoveryOperationsOptions model
				listApiGatewayDiscoveryOperationsOptionsModel := aiSecurityForAppsService.NewListApiGatewayDiscoveryOperationsOptions()
				listApiGatewayDiscoveryOperationsOptionsModel.SetDiff(true)
				listApiGatewayDiscoveryOperationsOptionsModel.SetDirection("asc")
				listApiGatewayDiscoveryOperationsOptionsModel.SetEndpoint("testString")
				listApiGatewayDiscoveryOperationsOptionsModel.SetHost([]string{"testString"})
				listApiGatewayDiscoveryOperationsOptionsModel.SetMethod([]string{"testString"})
				listApiGatewayDiscoveryOperationsOptionsModel.SetOrder("host")
				listApiGatewayDiscoveryOperationsOptionsModel.SetOrigin("ML")
				listApiGatewayDiscoveryOperationsOptionsModel.SetState("review")
				listApiGatewayDiscoveryOperationsOptionsModel.SetPage(int64(1))
				listApiGatewayDiscoveryOperationsOptionsModel.SetPerPage(int64(5))
				listApiGatewayDiscoveryOperationsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listApiGatewayDiscoveryOperationsOptionsModel).ToNot(BeNil())
				Expect(listApiGatewayDiscoveryOperationsOptionsModel.Diff).To(Equal(core.BoolPtr(true)))
				Expect(listApiGatewayDiscoveryOperationsOptionsModel.Direction).To(Equal(core.StringPtr("asc")))
				Expect(listApiGatewayDiscoveryOperationsOptionsModel.Endpoint).To(Equal(core.StringPtr("testString")))
				Expect(listApiGatewayDiscoveryOperationsOptionsModel.Host).To(Equal([]string{"testString"}))
				Expect(listApiGatewayDiscoveryOperationsOptionsModel.Method).To(Equal([]string{"testString"}))
				Expect(listApiGatewayDiscoveryOperationsOptionsModel.Order).To(Equal(core.StringPtr("host")))
				Expect(listApiGatewayDiscoveryOperationsOptionsModel.Origin).To(Equal(core.StringPtr("ML")))
				Expect(listApiGatewayDiscoveryOperationsOptionsModel.State).To(Equal(core.StringPtr("review")))
				Expect(listApiGatewayDiscoveryOperationsOptionsModel.Page).To(Equal(core.Int64Ptr(int64(1))))
				Expect(listApiGatewayDiscoveryOperationsOptionsModel.PerPage).To(Equal(core.Int64Ptr(int64(5))))
				Expect(listApiGatewayDiscoveryOperationsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListZoneApiGatewayOperationsOptions successfully`, func() {
				// Construct an instance of the ListZoneApiGatewayOperationsOptions model
				listZoneApiGatewayOperationsOptionsModel := aiSecurityForAppsService.NewListZoneApiGatewayOperationsOptions()
				listZoneApiGatewayOperationsOptionsModel.SetHost([]string{"testString"})
				listZoneApiGatewayOperationsOptionsModel.SetMethod([]string{"GET"})
				listZoneApiGatewayOperationsOptionsModel.SetEndpoint("testString")
				listZoneApiGatewayOperationsOptionsModel.SetLabels([]string{"testString"})
				listZoneApiGatewayOperationsOptionsModel.SetDirection("asc")
				listZoneApiGatewayOperationsOptionsModel.SetOrder("host")
				listZoneApiGatewayOperationsOptionsModel.SetPage(int64(1))
				listZoneApiGatewayOperationsOptionsModel.SetPerPage(int64(1))
				listZoneApiGatewayOperationsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listZoneApiGatewayOperationsOptionsModel).ToNot(BeNil())
				Expect(listZoneApiGatewayOperationsOptionsModel.Host).To(Equal([]string{"testString"}))
				Expect(listZoneApiGatewayOperationsOptionsModel.Method).To(Equal([]string{"GET"}))
				Expect(listZoneApiGatewayOperationsOptionsModel.Endpoint).To(Equal(core.StringPtr("testString")))
				Expect(listZoneApiGatewayOperationsOptionsModel.Labels).To(Equal([]string{"testString"}))
				Expect(listZoneApiGatewayOperationsOptionsModel.Direction).To(Equal(core.StringPtr("asc")))
				Expect(listZoneApiGatewayOperationsOptionsModel.Order).To(Equal(core.StringPtr("host")))
				Expect(listZoneApiGatewayOperationsOptionsModel.Page).To(Equal(core.Int64Ptr(int64(1))))
				Expect(listZoneApiGatewayOperationsOptionsModel.PerPage).To(Equal(core.Int64Ptr(int64(1))))
				Expect(listZoneApiGatewayOperationsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewReplaceZoneAiSecuritySettingsOptions successfully`, func() {
				// Construct an instance of the ReplaceZoneAiSecuritySettingsOptions model
				replaceZoneAiSecuritySettingsOptionsModel := aiSecurityForAppsService.NewReplaceZoneAiSecuritySettingsOptions()
				replaceZoneAiSecuritySettingsOptionsModel.SetEnabled(true)
				replaceZoneAiSecuritySettingsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(replaceZoneAiSecuritySettingsOptionsModel).ToNot(BeNil())
				Expect(replaceZoneAiSecuritySettingsOptionsModel.Enabled).To(Equal(core.BoolPtr(true)))
				Expect(replaceZoneAiSecuritySettingsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewReplaceZoneApiGatewayOperationLabelsOptions successfully`, func() {
				// Construct an instance of the ApiGatewayOperationsLabelsBulkInputSelectorInclude model
				apiGatewayOperationsLabelsBulkInputSelectorIncludeModel := new(aisecurityforappsv1.ApiGatewayOperationsLabelsBulkInputSelectorInclude)
				Expect(apiGatewayOperationsLabelsBulkInputSelectorIncludeModel).ToNot(BeNil())
				apiGatewayOperationsLabelsBulkInputSelectorIncludeModel.OperationIds = []string{"f174e90a-fafe-4643-bbbc-4a0ed4fc8415"}
				Expect(apiGatewayOperationsLabelsBulkInputSelectorIncludeModel.OperationIds).To(Equal([]string{"f174e90a-fafe-4643-bbbc-4a0ed4fc8415"}))

				// Construct an instance of the ApiGatewayOperationsLabelsBulkInputSelector model
				apiGatewayOperationsLabelsBulkInputSelectorModel := new(aisecurityforappsv1.ApiGatewayOperationsLabelsBulkInputSelector)
				Expect(apiGatewayOperationsLabelsBulkInputSelectorModel).ToNot(BeNil())
				apiGatewayOperationsLabelsBulkInputSelectorModel.Include = apiGatewayOperationsLabelsBulkInputSelectorIncludeModel
				Expect(apiGatewayOperationsLabelsBulkInputSelectorModel.Include).To(Equal(apiGatewayOperationsLabelsBulkInputSelectorIncludeModel))

				// Construct an instance of the ApiGatewayOperationsLabelsBulkInputUser model
				apiGatewayOperationsLabelsBulkInputUserModel := new(aisecurityforappsv1.ApiGatewayOperationsLabelsBulkInputUser)
				Expect(apiGatewayOperationsLabelsBulkInputUserModel).ToNot(BeNil())
				apiGatewayOperationsLabelsBulkInputUserModel.Labels = []string{"testString"}
				Expect(apiGatewayOperationsLabelsBulkInputUserModel.Labels).To(Equal([]string{"testString"}))

				// Construct an instance of the ApiGatewayOperationsLabelsBulkInputManaged model
				apiGatewayOperationsLabelsBulkInputManagedModel := new(aisecurityforappsv1.ApiGatewayOperationsLabelsBulkInputManaged)
				Expect(apiGatewayOperationsLabelsBulkInputManagedModel).ToNot(BeNil())
				apiGatewayOperationsLabelsBulkInputManagedModel.Labels = []string{"testString"}
				Expect(apiGatewayOperationsLabelsBulkInputManagedModel.Labels).To(Equal([]string{"testString"}))

				// Construct an instance of the ReplaceZoneApiGatewayOperationLabelsOptions model
				replaceZoneApiGatewayOperationLabelsOptionsModel := aiSecurityForAppsService.NewReplaceZoneApiGatewayOperationLabelsOptions()
				replaceZoneApiGatewayOperationLabelsOptionsModel.SetSelector(apiGatewayOperationsLabelsBulkInputSelectorModel)
				replaceZoneApiGatewayOperationLabelsOptionsModel.SetUser(apiGatewayOperationsLabelsBulkInputUserModel)
				replaceZoneApiGatewayOperationLabelsOptionsModel.SetManaged(apiGatewayOperationsLabelsBulkInputManagedModel)
				replaceZoneApiGatewayOperationLabelsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(replaceZoneApiGatewayOperationLabelsOptionsModel).ToNot(BeNil())
				Expect(replaceZoneApiGatewayOperationLabelsOptionsModel.Selector).To(Equal(apiGatewayOperationsLabelsBulkInputSelectorModel))
				Expect(replaceZoneApiGatewayOperationLabelsOptionsModel.User).To(Equal(apiGatewayOperationsLabelsBulkInputUserModel))
				Expect(replaceZoneApiGatewayOperationLabelsOptionsModel.Managed).To(Equal(apiGatewayOperationsLabelsBulkInputManagedModel))
				Expect(replaceZoneApiGatewayOperationLabelsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewSetZoneApiGatewayOperationLabelsOptions successfully`, func() {
				// Construct an instance of the SetZoneApiGatewayOperationLabelsOptions model
				operationID := "testString"
				setZoneApiGatewayOperationLabelsOptionsModel := aiSecurityForAppsService.NewSetZoneApiGatewayOperationLabelsOptions(operationID)
				setZoneApiGatewayOperationLabelsOptionsModel.SetOperationID("testString")
				setZoneApiGatewayOperationLabelsOptionsModel.SetUser([]string{})
				setZoneApiGatewayOperationLabelsOptionsModel.SetManaged([]string{"cf-llm"})
				setZoneApiGatewayOperationLabelsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(setZoneApiGatewayOperationLabelsOptionsModel).ToNot(BeNil())
				Expect(setZoneApiGatewayOperationLabelsOptionsModel.OperationID).To(Equal(core.StringPtr("testString")))
				Expect(setZoneApiGatewayOperationLabelsOptionsModel.User).To(Equal([]string{}))
				Expect(setZoneApiGatewayOperationLabelsOptionsModel.Managed).To(Equal([]string{"cf-llm"}))
				Expect(setZoneApiGatewayOperationLabelsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUnsetZoneApiGatewayOperationLabelsOptions successfully`, func() {
				// Construct an instance of the UnsetZoneApiGatewayOperationLabelsOptions model
				operationID := "testString"
				unsetZoneApiGatewayOperationLabelsOptionsModel := aiSecurityForAppsService.NewUnsetZoneApiGatewayOperationLabelsOptions(operationID)
				unsetZoneApiGatewayOperationLabelsOptionsModel.SetOperationID("testString")
				unsetZoneApiGatewayOperationLabelsOptionsModel.SetUser([]string{"testString"})
				unsetZoneApiGatewayOperationLabelsOptionsModel.SetManaged([]string{"testString"})
				unsetZoneApiGatewayOperationLabelsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(unsetZoneApiGatewayOperationLabelsOptionsModel).ToNot(BeNil())
				Expect(unsetZoneApiGatewayOperationLabelsOptionsModel.OperationID).To(Equal(core.StringPtr("testString")))
				Expect(unsetZoneApiGatewayOperationLabelsOptionsModel.User).To(Equal([]string{"testString"}))
				Expect(unsetZoneApiGatewayOperationLabelsOptionsModel.Managed).To(Equal([]string{"testString"}))
				Expect(unsetZoneApiGatewayOperationLabelsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateApiGatewayOperationLabelsOptions successfully`, func() {
				// Construct an instance of the ApiGatewayOperationsLabelsBulkInputSelectorInclude model
				apiGatewayOperationsLabelsBulkInputSelectorIncludeModel := new(aisecurityforappsv1.ApiGatewayOperationsLabelsBulkInputSelectorInclude)
				Expect(apiGatewayOperationsLabelsBulkInputSelectorIncludeModel).ToNot(BeNil())
				apiGatewayOperationsLabelsBulkInputSelectorIncludeModel.OperationIds = []string{"f174e90a-fafe-4643-bbbc-4a0ed4fc8415"}
				Expect(apiGatewayOperationsLabelsBulkInputSelectorIncludeModel.OperationIds).To(Equal([]string{"f174e90a-fafe-4643-bbbc-4a0ed4fc8415"}))

				// Construct an instance of the ApiGatewayOperationsLabelsBulkInputSelector model
				apiGatewayOperationsLabelsBulkInputSelectorModel := new(aisecurityforappsv1.ApiGatewayOperationsLabelsBulkInputSelector)
				Expect(apiGatewayOperationsLabelsBulkInputSelectorModel).ToNot(BeNil())
				apiGatewayOperationsLabelsBulkInputSelectorModel.Include = apiGatewayOperationsLabelsBulkInputSelectorIncludeModel
				Expect(apiGatewayOperationsLabelsBulkInputSelectorModel.Include).To(Equal(apiGatewayOperationsLabelsBulkInputSelectorIncludeModel))

				// Construct an instance of the ApiGatewayOperationsLabelsBulkInputUser model
				apiGatewayOperationsLabelsBulkInputUserModel := new(aisecurityforappsv1.ApiGatewayOperationsLabelsBulkInputUser)
				Expect(apiGatewayOperationsLabelsBulkInputUserModel).ToNot(BeNil())
				apiGatewayOperationsLabelsBulkInputUserModel.Labels = []string{"testString"}
				Expect(apiGatewayOperationsLabelsBulkInputUserModel.Labels).To(Equal([]string{"testString"}))

				// Construct an instance of the ApiGatewayOperationsLabelsBulkInputManaged model
				apiGatewayOperationsLabelsBulkInputManagedModel := new(aisecurityforappsv1.ApiGatewayOperationsLabelsBulkInputManaged)
				Expect(apiGatewayOperationsLabelsBulkInputManagedModel).ToNot(BeNil())
				apiGatewayOperationsLabelsBulkInputManagedModel.Labels = []string{"testString"}
				Expect(apiGatewayOperationsLabelsBulkInputManagedModel.Labels).To(Equal([]string{"testString"}))

				// Construct an instance of the UpdateApiGatewayOperationLabelsOptions model
				updateApiGatewayOperationLabelsOptionsModel := aiSecurityForAppsService.NewUpdateApiGatewayOperationLabelsOptions()
				updateApiGatewayOperationLabelsOptionsModel.SetSelector(apiGatewayOperationsLabelsBulkInputSelectorModel)
				updateApiGatewayOperationLabelsOptionsModel.SetUser(apiGatewayOperationsLabelsBulkInputUserModel)
				updateApiGatewayOperationLabelsOptionsModel.SetManaged(apiGatewayOperationsLabelsBulkInputManagedModel)
				updateApiGatewayOperationLabelsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateApiGatewayOperationLabelsOptionsModel).ToNot(BeNil())
				Expect(updateApiGatewayOperationLabelsOptionsModel.Selector).To(Equal(apiGatewayOperationsLabelsBulkInputSelectorModel))
				Expect(updateApiGatewayOperationLabelsOptionsModel.User).To(Equal(apiGatewayOperationsLabelsBulkInputUserModel))
				Expect(updateApiGatewayOperationLabelsOptionsModel.Managed).To(Equal(apiGatewayOperationsLabelsBulkInputManagedModel))
				Expect(updateApiGatewayOperationLabelsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateZoneApiGatewayDiscoveryOperationOptions successfully`, func() {
				// Construct an instance of the UpdateZoneApiGatewayDiscoveryOperationOptions model
				updateZoneApiGatewayDiscoveryOperationOptionsModel := aiSecurityForAppsService.NewUpdateZoneApiGatewayDiscoveryOperationOptions()
				updateZoneApiGatewayDiscoveryOperationOptionsModel.SetRequestBody(map[string]interface{}{"anyKey": "anyValue"})
				updateZoneApiGatewayDiscoveryOperationOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateZoneApiGatewayDiscoveryOperationOptionsModel).ToNot(BeNil())
				Expect(updateZoneApiGatewayDiscoveryOperationOptionsModel.RequestBody).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))
				Expect(updateZoneApiGatewayDiscoveryOperationOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewApiGatewayOperation successfully`, func() {
				method := "POST"
				host := "api.example.com"
				endpoint := "/v1/messages"
				_model, err := aiSecurityForAppsService.NewApiGatewayOperation(method, host, endpoint)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
		})
	})
	Describe(`Model unmarshaling tests`, func() {
		It(`Invoke UnmarshalApiGatewayOperationsLabelsBulkDeleteInputManaged successfully`, func() {
			// Construct an instance of the model.
			model := new(aisecurityforappsv1.ApiGatewayOperationsLabelsBulkDeleteInputManaged)
			model.Labels = []string{"testString"}

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *aisecurityforappsv1.ApiGatewayOperationsLabelsBulkDeleteInputManaged
			err = aisecurityforappsv1.UnmarshalApiGatewayOperationsLabelsBulkDeleteInputManaged(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalApiGatewayOperationsLabelsBulkDeleteInputSelector successfully`, func() {
			// Construct an instance of the model.
			model := new(aisecurityforappsv1.ApiGatewayOperationsLabelsBulkDeleteInputSelector)
			model.Include = nil

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *aisecurityforappsv1.ApiGatewayOperationsLabelsBulkDeleteInputSelector
			err = aisecurityforappsv1.UnmarshalApiGatewayOperationsLabelsBulkDeleteInputSelector(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalApiGatewayOperationsLabelsBulkDeleteInputSelectorInclude successfully`, func() {
			// Construct an instance of the model.
			model := new(aisecurityforappsv1.ApiGatewayOperationsLabelsBulkDeleteInputSelectorInclude)
			model.OperationIds = []string{"f174e90a-fafe-4643-bbbc-4a0ed4fc8415"}

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *aisecurityforappsv1.ApiGatewayOperationsLabelsBulkDeleteInputSelectorInclude
			err = aisecurityforappsv1.UnmarshalApiGatewayOperationsLabelsBulkDeleteInputSelectorInclude(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalApiGatewayOperationsLabelsBulkDeleteInputUser successfully`, func() {
			// Construct an instance of the model.
			model := new(aisecurityforappsv1.ApiGatewayOperationsLabelsBulkDeleteInputUser)
			model.Labels = []string{"testString"}

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *aisecurityforappsv1.ApiGatewayOperationsLabelsBulkDeleteInputUser
			err = aisecurityforappsv1.UnmarshalApiGatewayOperationsLabelsBulkDeleteInputUser(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalApiGatewayOperationsLabelsBulkInputManaged successfully`, func() {
			// Construct an instance of the model.
			model := new(aisecurityforappsv1.ApiGatewayOperationsLabelsBulkInputManaged)
			model.Labels = []string{"testString"}

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *aisecurityforappsv1.ApiGatewayOperationsLabelsBulkInputManaged
			err = aisecurityforappsv1.UnmarshalApiGatewayOperationsLabelsBulkInputManaged(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalApiGatewayOperationsLabelsBulkInputSelector successfully`, func() {
			// Construct an instance of the model.
			model := new(aisecurityforappsv1.ApiGatewayOperationsLabelsBulkInputSelector)
			model.Include = nil

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *aisecurityforappsv1.ApiGatewayOperationsLabelsBulkInputSelector
			err = aisecurityforappsv1.UnmarshalApiGatewayOperationsLabelsBulkInputSelector(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalApiGatewayOperationsLabelsBulkInputSelectorInclude successfully`, func() {
			// Construct an instance of the model.
			model := new(aisecurityforappsv1.ApiGatewayOperationsLabelsBulkInputSelectorInclude)
			model.OperationIds = []string{"f174e90a-fafe-4643-bbbc-4a0ed4fc8415"}

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *aisecurityforappsv1.ApiGatewayOperationsLabelsBulkInputSelectorInclude
			err = aisecurityforappsv1.UnmarshalApiGatewayOperationsLabelsBulkInputSelectorInclude(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalApiGatewayOperationsLabelsBulkInputUser successfully`, func() {
			// Construct an instance of the model.
			model := new(aisecurityforappsv1.ApiGatewayOperationsLabelsBulkInputUser)
			model.Labels = []string{"testString"}

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *aisecurityforappsv1.ApiGatewayOperationsLabelsBulkInputUser
			err = aisecurityforappsv1.UnmarshalApiGatewayOperationsLabelsBulkInputUser(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalApiGatewayOperation successfully`, func() {
			// Construct an instance of the model.
			model := new(aisecurityforappsv1.ApiGatewayOperation)
			model.Method = core.StringPtr("POST")
			model.Host = core.StringPtr("api.example.com")
			model.Endpoint = core.StringPtr("/v1/messages")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *aisecurityforappsv1.ApiGatewayOperation
			err = aisecurityforappsv1.UnmarshalApiGatewayOperation(raw, &result)
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
