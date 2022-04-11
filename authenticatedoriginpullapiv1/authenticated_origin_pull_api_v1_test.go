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

package authenticatedoriginpullapiv1_test

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
	"github.com/IBM/networking-go-sdk/authenticatedoriginpullapiv1"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe(`AuthenticatedOriginPullApiV1`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		It(`Instantiate service client`, func() {
			authenticatedOriginPullApiService, serviceErr := authenticatedoriginpullapiv1.NewAuthenticatedOriginPullApiV1(&authenticatedoriginpullapiv1.AuthenticatedOriginPullApiV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
				Crn: core.StringPtr(crn),
				ZoneIdentifier: core.StringPtr(zoneIdentifier),
			})
			Expect(authenticatedOriginPullApiService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			authenticatedOriginPullApiService, serviceErr := authenticatedoriginpullapiv1.NewAuthenticatedOriginPullApiV1(&authenticatedoriginpullapiv1.AuthenticatedOriginPullApiV1Options{
				URL: "{BAD_URL_STRING",
				Crn: core.StringPtr(crn),
				ZoneIdentifier: core.StringPtr(zoneIdentifier),
			})
			Expect(authenticatedOriginPullApiService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			authenticatedOriginPullApiService, serviceErr := authenticatedoriginpullapiv1.NewAuthenticatedOriginPullApiV1(&authenticatedoriginpullapiv1.AuthenticatedOriginPullApiV1Options{
				URL: "https://authenticatedoriginpullapiv1/api",
				Crn: core.StringPtr(crn),
				ZoneIdentifier: core.StringPtr(zoneIdentifier),
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(authenticatedOriginPullApiService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Validation Error`, func() {
			authenticatedOriginPullApiService, serviceErr := authenticatedoriginpullapiv1.NewAuthenticatedOriginPullApiV1(&authenticatedoriginpullapiv1.AuthenticatedOriginPullApiV1Options{})
			Expect(authenticatedOriginPullApiService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"AUTHENTICATED_ORIGIN_PULL_API_URL": "https://authenticatedoriginpullapiv1/api",
				"AUTHENTICATED_ORIGIN_PULL_API_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				authenticatedOriginPullApiService, serviceErr := authenticatedoriginpullapiv1.NewAuthenticatedOriginPullApiV1UsingExternalConfig(&authenticatedoriginpullapiv1.AuthenticatedOriginPullApiV1Options{
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(authenticatedOriginPullApiService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := authenticatedOriginPullApiService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != authenticatedOriginPullApiService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(authenticatedOriginPullApiService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(authenticatedOriginPullApiService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				authenticatedOriginPullApiService, serviceErr := authenticatedoriginpullapiv1.NewAuthenticatedOriginPullApiV1UsingExternalConfig(&authenticatedoriginpullapiv1.AuthenticatedOriginPullApiV1Options{
					URL: "https://testService/api",
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(authenticatedOriginPullApiService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(authenticatedOriginPullApiService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := authenticatedOriginPullApiService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != authenticatedOriginPullApiService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(authenticatedOriginPullApiService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(authenticatedOriginPullApiService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				authenticatedOriginPullApiService, serviceErr := authenticatedoriginpullapiv1.NewAuthenticatedOriginPullApiV1UsingExternalConfig(&authenticatedoriginpullapiv1.AuthenticatedOriginPullApiV1Options{
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				err := authenticatedOriginPullApiService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(authenticatedOriginPullApiService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(authenticatedOriginPullApiService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := authenticatedOriginPullApiService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != authenticatedOriginPullApiService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(authenticatedOriginPullApiService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(authenticatedOriginPullApiService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"AUTHENTICATED_ORIGIN_PULL_API_URL": "https://authenticatedoriginpullapiv1/api",
				"AUTHENTICATED_ORIGIN_PULL_API_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			authenticatedOriginPullApiService, serviceErr := authenticatedoriginpullapiv1.NewAuthenticatedOriginPullApiV1UsingExternalConfig(&authenticatedoriginpullapiv1.AuthenticatedOriginPullApiV1Options{
				Crn: core.StringPtr(crn),
				ZoneIdentifier: core.StringPtr(zoneIdentifier),
			})

			It(`Instantiate service client with error`, func() {
				Expect(authenticatedOriginPullApiService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"AUTHENTICATED_ORIGIN_PULL_API_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			authenticatedOriginPullApiService, serviceErr := authenticatedoriginpullapiv1.NewAuthenticatedOriginPullApiV1UsingExternalConfig(&authenticatedoriginpullapiv1.AuthenticatedOriginPullApiV1Options{
				URL: "{BAD_URL_STRING",
				Crn: core.StringPtr(crn),
				ZoneIdentifier: core.StringPtr(zoneIdentifier),
			})

			It(`Instantiate service client with error`, func() {
				Expect(authenticatedOriginPullApiService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = authenticatedoriginpullapiv1.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`GetZoneOriginPullSettings(getZoneOriginPullSettingsOptions *GetZoneOriginPullSettingsOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getZoneOriginPullSettingsPath := "/v1/testString/zones/testString/origin_tls_client_auth/settings"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getZoneOriginPullSettingsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetZoneOriginPullSettings with error: Operation response processing error`, func() {
				authenticatedOriginPullApiService, serviceErr := authenticatedoriginpullapiv1.NewAuthenticatedOriginPullApiV1(&authenticatedoriginpullapiv1.AuthenticatedOriginPullApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(authenticatedOriginPullApiService).ToNot(BeNil())

				// Construct an instance of the GetZoneOriginPullSettingsOptions model
				getZoneOriginPullSettingsOptionsModel := new(authenticatedoriginpullapiv1.GetZoneOriginPullSettingsOptions)
				getZoneOriginPullSettingsOptionsModel.XCorrelationID = core.StringPtr("testString")
				getZoneOriginPullSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := authenticatedOriginPullApiService.GetZoneOriginPullSettings(getZoneOriginPullSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				authenticatedOriginPullApiService.EnableRetries(0, 0)
				result, response, operationErr = authenticatedOriginPullApiService.GetZoneOriginPullSettings(getZoneOriginPullSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetZoneOriginPullSettings(getZoneOriginPullSettingsOptions *GetZoneOriginPullSettingsOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getZoneOriginPullSettingsPath := "/v1/testString/zones/testString/origin_tls_client_auth/settings"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getZoneOriginPullSettingsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"result": {"enabled": true}, "success": true, "errors": ["Errors"], "messages": ["Messages"]}`)
				}))
			})
			It(`Invoke GetZoneOriginPullSettings successfully with retries`, func() {
				authenticatedOriginPullApiService, serviceErr := authenticatedoriginpullapiv1.NewAuthenticatedOriginPullApiV1(&authenticatedoriginpullapiv1.AuthenticatedOriginPullApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(authenticatedOriginPullApiService).ToNot(BeNil())
				authenticatedOriginPullApiService.EnableRetries(0, 0)

				// Construct an instance of the GetZoneOriginPullSettingsOptions model
				getZoneOriginPullSettingsOptionsModel := new(authenticatedoriginpullapiv1.GetZoneOriginPullSettingsOptions)
				getZoneOriginPullSettingsOptionsModel.XCorrelationID = core.StringPtr("testString")
				getZoneOriginPullSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := authenticatedOriginPullApiService.GetZoneOriginPullSettingsWithContext(ctx, getZoneOriginPullSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				authenticatedOriginPullApiService.DisableRetries()
				result, response, operationErr := authenticatedOriginPullApiService.GetZoneOriginPullSettings(getZoneOriginPullSettingsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = authenticatedOriginPullApiService.GetZoneOriginPullSettingsWithContext(ctx, getZoneOriginPullSettingsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getZoneOriginPullSettingsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"result": {"enabled": true}, "success": true, "errors": ["Errors"], "messages": ["Messages"]}`)
				}))
			})
			It(`Invoke GetZoneOriginPullSettings successfully`, func() {
				authenticatedOriginPullApiService, serviceErr := authenticatedoriginpullapiv1.NewAuthenticatedOriginPullApiV1(&authenticatedoriginpullapiv1.AuthenticatedOriginPullApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(authenticatedOriginPullApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := authenticatedOriginPullApiService.GetZoneOriginPullSettings(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetZoneOriginPullSettingsOptions model
				getZoneOriginPullSettingsOptionsModel := new(authenticatedoriginpullapiv1.GetZoneOriginPullSettingsOptions)
				getZoneOriginPullSettingsOptionsModel.XCorrelationID = core.StringPtr("testString")
				getZoneOriginPullSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = authenticatedOriginPullApiService.GetZoneOriginPullSettings(getZoneOriginPullSettingsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetZoneOriginPullSettings with error: Operation request error`, func() {
				authenticatedOriginPullApiService, serviceErr := authenticatedoriginpullapiv1.NewAuthenticatedOriginPullApiV1(&authenticatedoriginpullapiv1.AuthenticatedOriginPullApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(authenticatedOriginPullApiService).ToNot(BeNil())

				// Construct an instance of the GetZoneOriginPullSettingsOptions model
				getZoneOriginPullSettingsOptionsModel := new(authenticatedoriginpullapiv1.GetZoneOriginPullSettingsOptions)
				getZoneOriginPullSettingsOptionsModel.XCorrelationID = core.StringPtr("testString")
				getZoneOriginPullSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := authenticatedOriginPullApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := authenticatedOriginPullApiService.GetZoneOriginPullSettings(getZoneOriginPullSettingsOptionsModel)
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
			It(`Invoke GetZoneOriginPullSettings successfully`, func() {
				authenticatedOriginPullApiService, serviceErr := authenticatedoriginpullapiv1.NewAuthenticatedOriginPullApiV1(&authenticatedoriginpullapiv1.AuthenticatedOriginPullApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(authenticatedOriginPullApiService).ToNot(BeNil())

				// Construct an instance of the GetZoneOriginPullSettingsOptions model
				getZoneOriginPullSettingsOptionsModel := new(authenticatedoriginpullapiv1.GetZoneOriginPullSettingsOptions)
				getZoneOriginPullSettingsOptionsModel.XCorrelationID = core.StringPtr("testString")
				getZoneOriginPullSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := authenticatedOriginPullApiService.GetZoneOriginPullSettings(getZoneOriginPullSettingsOptionsModel)
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
	Describe(`SetZoneOriginPullSettings(setZoneOriginPullSettingsOptions *SetZoneOriginPullSettingsOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		setZoneOriginPullSettingsPath := "/v1/testString/zones/testString/origin_tls_client_auth/settings"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(setZoneOriginPullSettingsPath))
					Expect(req.Method).To(Equal("PUT"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke SetZoneOriginPullSettings with error: Operation response processing error`, func() {
				authenticatedOriginPullApiService, serviceErr := authenticatedoriginpullapiv1.NewAuthenticatedOriginPullApiV1(&authenticatedoriginpullapiv1.AuthenticatedOriginPullApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(authenticatedOriginPullApiService).ToNot(BeNil())

				// Construct an instance of the SetZoneOriginPullSettingsOptions model
				setZoneOriginPullSettingsOptionsModel := new(authenticatedoriginpullapiv1.SetZoneOriginPullSettingsOptions)
				setZoneOriginPullSettingsOptionsModel.Enabled = core.BoolPtr(true)
				setZoneOriginPullSettingsOptionsModel.XCorrelationID = core.StringPtr("testString")
				setZoneOriginPullSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := authenticatedOriginPullApiService.SetZoneOriginPullSettings(setZoneOriginPullSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				authenticatedOriginPullApiService.EnableRetries(0, 0)
				result, response, operationErr = authenticatedOriginPullApiService.SetZoneOriginPullSettings(setZoneOriginPullSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`SetZoneOriginPullSettings(setZoneOriginPullSettingsOptions *SetZoneOriginPullSettingsOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		setZoneOriginPullSettingsPath := "/v1/testString/zones/testString/origin_tls_client_auth/settings"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(setZoneOriginPullSettingsPath))
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

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"result": {"enabled": true}, "success": true, "errors": ["Errors"], "messages": ["Messages"]}`)
				}))
			})
			It(`Invoke SetZoneOriginPullSettings successfully with retries`, func() {
				authenticatedOriginPullApiService, serviceErr := authenticatedoriginpullapiv1.NewAuthenticatedOriginPullApiV1(&authenticatedoriginpullapiv1.AuthenticatedOriginPullApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(authenticatedOriginPullApiService).ToNot(BeNil())
				authenticatedOriginPullApiService.EnableRetries(0, 0)

				// Construct an instance of the SetZoneOriginPullSettingsOptions model
				setZoneOriginPullSettingsOptionsModel := new(authenticatedoriginpullapiv1.SetZoneOriginPullSettingsOptions)
				setZoneOriginPullSettingsOptionsModel.Enabled = core.BoolPtr(true)
				setZoneOriginPullSettingsOptionsModel.XCorrelationID = core.StringPtr("testString")
				setZoneOriginPullSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := authenticatedOriginPullApiService.SetZoneOriginPullSettingsWithContext(ctx, setZoneOriginPullSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				authenticatedOriginPullApiService.DisableRetries()
				result, response, operationErr := authenticatedOriginPullApiService.SetZoneOriginPullSettings(setZoneOriginPullSettingsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = authenticatedOriginPullApiService.SetZoneOriginPullSettingsWithContext(ctx, setZoneOriginPullSettingsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(setZoneOriginPullSettingsPath))
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

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"result": {"enabled": true}, "success": true, "errors": ["Errors"], "messages": ["Messages"]}`)
				}))
			})
			It(`Invoke SetZoneOriginPullSettings successfully`, func() {
				authenticatedOriginPullApiService, serviceErr := authenticatedoriginpullapiv1.NewAuthenticatedOriginPullApiV1(&authenticatedoriginpullapiv1.AuthenticatedOriginPullApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(authenticatedOriginPullApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := authenticatedOriginPullApiService.SetZoneOriginPullSettings(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the SetZoneOriginPullSettingsOptions model
				setZoneOriginPullSettingsOptionsModel := new(authenticatedoriginpullapiv1.SetZoneOriginPullSettingsOptions)
				setZoneOriginPullSettingsOptionsModel.Enabled = core.BoolPtr(true)
				setZoneOriginPullSettingsOptionsModel.XCorrelationID = core.StringPtr("testString")
				setZoneOriginPullSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = authenticatedOriginPullApiService.SetZoneOriginPullSettings(setZoneOriginPullSettingsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke SetZoneOriginPullSettings with error: Operation request error`, func() {
				authenticatedOriginPullApiService, serviceErr := authenticatedoriginpullapiv1.NewAuthenticatedOriginPullApiV1(&authenticatedoriginpullapiv1.AuthenticatedOriginPullApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(authenticatedOriginPullApiService).ToNot(BeNil())

				// Construct an instance of the SetZoneOriginPullSettingsOptions model
				setZoneOriginPullSettingsOptionsModel := new(authenticatedoriginpullapiv1.SetZoneOriginPullSettingsOptions)
				setZoneOriginPullSettingsOptionsModel.Enabled = core.BoolPtr(true)
				setZoneOriginPullSettingsOptionsModel.XCorrelationID = core.StringPtr("testString")
				setZoneOriginPullSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := authenticatedOriginPullApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := authenticatedOriginPullApiService.SetZoneOriginPullSettings(setZoneOriginPullSettingsOptionsModel)
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
			It(`Invoke SetZoneOriginPullSettings successfully`, func() {
				authenticatedOriginPullApiService, serviceErr := authenticatedoriginpullapiv1.NewAuthenticatedOriginPullApiV1(&authenticatedoriginpullapiv1.AuthenticatedOriginPullApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(authenticatedOriginPullApiService).ToNot(BeNil())

				// Construct an instance of the SetZoneOriginPullSettingsOptions model
				setZoneOriginPullSettingsOptionsModel := new(authenticatedoriginpullapiv1.SetZoneOriginPullSettingsOptions)
				setZoneOriginPullSettingsOptionsModel.Enabled = core.BoolPtr(true)
				setZoneOriginPullSettingsOptionsModel.XCorrelationID = core.StringPtr("testString")
				setZoneOriginPullSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := authenticatedOriginPullApiService.SetZoneOriginPullSettings(setZoneOriginPullSettingsOptionsModel)
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
	Describe(`ListZoneOriginPullCertificates(listZoneOriginPullCertificatesOptions *ListZoneOriginPullCertificatesOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		listZoneOriginPullCertificatesPath := "/v1/testString/zones/testString/origin_tls_client_auth"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listZoneOriginPullCertificatesPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListZoneOriginPullCertificates with error: Operation response processing error`, func() {
				authenticatedOriginPullApiService, serviceErr := authenticatedoriginpullapiv1.NewAuthenticatedOriginPullApiV1(&authenticatedoriginpullapiv1.AuthenticatedOriginPullApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(authenticatedOriginPullApiService).ToNot(BeNil())

				// Construct an instance of the ListZoneOriginPullCertificatesOptions model
				listZoneOriginPullCertificatesOptionsModel := new(authenticatedoriginpullapiv1.ListZoneOriginPullCertificatesOptions)
				listZoneOriginPullCertificatesOptionsModel.XCorrelationID = core.StringPtr("testString")
				listZoneOriginPullCertificatesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := authenticatedOriginPullApiService.ListZoneOriginPullCertificates(listZoneOriginPullCertificatesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				authenticatedOriginPullApiService.EnableRetries(0, 0)
				result, response, operationErr = authenticatedOriginPullApiService.ListZoneOriginPullCertificates(listZoneOriginPullCertificatesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListZoneOriginPullCertificates(listZoneOriginPullCertificatesOptions *ListZoneOriginPullCertificatesOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		listZoneOriginPullCertificatesPath := "/v1/testString/zones/testString/origin_tls_client_auth"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listZoneOriginPullCertificatesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"result": [{"id": "0f405ba2-8c18-49eb-a30b-28b85427780f", "certificate": "-----BEGIN CERTIFICATE-----\n......\n-----END CERTIFICATE-----\n", "issuer": "GlobalSign", "signature": "SHA256WithRSA", "status": "active", "expires_on": "2100-01-01T05:20:00Z", "uploaded_on": "2100-01-01T05:20:00Z"}], "success": true, "errors": ["Errors"], "messages": ["Messages"]}`)
				}))
			})
			It(`Invoke ListZoneOriginPullCertificates successfully with retries`, func() {
				authenticatedOriginPullApiService, serviceErr := authenticatedoriginpullapiv1.NewAuthenticatedOriginPullApiV1(&authenticatedoriginpullapiv1.AuthenticatedOriginPullApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(authenticatedOriginPullApiService).ToNot(BeNil())
				authenticatedOriginPullApiService.EnableRetries(0, 0)

				// Construct an instance of the ListZoneOriginPullCertificatesOptions model
				listZoneOriginPullCertificatesOptionsModel := new(authenticatedoriginpullapiv1.ListZoneOriginPullCertificatesOptions)
				listZoneOriginPullCertificatesOptionsModel.XCorrelationID = core.StringPtr("testString")
				listZoneOriginPullCertificatesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := authenticatedOriginPullApiService.ListZoneOriginPullCertificatesWithContext(ctx, listZoneOriginPullCertificatesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				authenticatedOriginPullApiService.DisableRetries()
				result, response, operationErr := authenticatedOriginPullApiService.ListZoneOriginPullCertificates(listZoneOriginPullCertificatesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = authenticatedOriginPullApiService.ListZoneOriginPullCertificatesWithContext(ctx, listZoneOriginPullCertificatesOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listZoneOriginPullCertificatesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"result": [{"id": "0f405ba2-8c18-49eb-a30b-28b85427780f", "certificate": "-----BEGIN CERTIFICATE-----\n......\n-----END CERTIFICATE-----\n", "issuer": "GlobalSign", "signature": "SHA256WithRSA", "status": "active", "expires_on": "2100-01-01T05:20:00Z", "uploaded_on": "2100-01-01T05:20:00Z"}], "success": true, "errors": ["Errors"], "messages": ["Messages"]}`)
				}))
			})
			It(`Invoke ListZoneOriginPullCertificates successfully`, func() {
				authenticatedOriginPullApiService, serviceErr := authenticatedoriginpullapiv1.NewAuthenticatedOriginPullApiV1(&authenticatedoriginpullapiv1.AuthenticatedOriginPullApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(authenticatedOriginPullApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := authenticatedOriginPullApiService.ListZoneOriginPullCertificates(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListZoneOriginPullCertificatesOptions model
				listZoneOriginPullCertificatesOptionsModel := new(authenticatedoriginpullapiv1.ListZoneOriginPullCertificatesOptions)
				listZoneOriginPullCertificatesOptionsModel.XCorrelationID = core.StringPtr("testString")
				listZoneOriginPullCertificatesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = authenticatedOriginPullApiService.ListZoneOriginPullCertificates(listZoneOriginPullCertificatesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListZoneOriginPullCertificates with error: Operation request error`, func() {
				authenticatedOriginPullApiService, serviceErr := authenticatedoriginpullapiv1.NewAuthenticatedOriginPullApiV1(&authenticatedoriginpullapiv1.AuthenticatedOriginPullApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(authenticatedOriginPullApiService).ToNot(BeNil())

				// Construct an instance of the ListZoneOriginPullCertificatesOptions model
				listZoneOriginPullCertificatesOptionsModel := new(authenticatedoriginpullapiv1.ListZoneOriginPullCertificatesOptions)
				listZoneOriginPullCertificatesOptionsModel.XCorrelationID = core.StringPtr("testString")
				listZoneOriginPullCertificatesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := authenticatedOriginPullApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := authenticatedOriginPullApiService.ListZoneOriginPullCertificates(listZoneOriginPullCertificatesOptionsModel)
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
			It(`Invoke ListZoneOriginPullCertificates successfully`, func() {
				authenticatedOriginPullApiService, serviceErr := authenticatedoriginpullapiv1.NewAuthenticatedOriginPullApiV1(&authenticatedoriginpullapiv1.AuthenticatedOriginPullApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(authenticatedOriginPullApiService).ToNot(BeNil())

				// Construct an instance of the ListZoneOriginPullCertificatesOptions model
				listZoneOriginPullCertificatesOptionsModel := new(authenticatedoriginpullapiv1.ListZoneOriginPullCertificatesOptions)
				listZoneOriginPullCertificatesOptionsModel.XCorrelationID = core.StringPtr("testString")
				listZoneOriginPullCertificatesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := authenticatedOriginPullApiService.ListZoneOriginPullCertificates(listZoneOriginPullCertificatesOptionsModel)
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
	Describe(`UploadZoneOriginPullCertificate(uploadZoneOriginPullCertificateOptions *UploadZoneOriginPullCertificateOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		uploadZoneOriginPullCertificatePath := "/v1/testString/zones/testString/origin_tls_client_auth"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(uploadZoneOriginPullCertificatePath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UploadZoneOriginPullCertificate with error: Operation response processing error`, func() {
				authenticatedOriginPullApiService, serviceErr := authenticatedoriginpullapiv1.NewAuthenticatedOriginPullApiV1(&authenticatedoriginpullapiv1.AuthenticatedOriginPullApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(authenticatedOriginPullApiService).ToNot(BeNil())

				// Construct an instance of the UploadZoneOriginPullCertificateOptions model
				uploadZoneOriginPullCertificateOptionsModel := new(authenticatedoriginpullapiv1.UploadZoneOriginPullCertificateOptions)
				uploadZoneOriginPullCertificateOptionsModel.Certificate = core.StringPtr("-----BEGIN CERTIFICATE-----\n......\n-----END CERTIFICATE-----\n")
				uploadZoneOriginPullCertificateOptionsModel.PrivateKey = core.StringPtr("-----BEGIN RSA PRIVATE KEY-----\n......\n-----END RSA PRIVATE KEY-----\n")
				uploadZoneOriginPullCertificateOptionsModel.XCorrelationID = core.StringPtr("testString")
				uploadZoneOriginPullCertificateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := authenticatedOriginPullApiService.UploadZoneOriginPullCertificate(uploadZoneOriginPullCertificateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				authenticatedOriginPullApiService.EnableRetries(0, 0)
				result, response, operationErr = authenticatedOriginPullApiService.UploadZoneOriginPullCertificate(uploadZoneOriginPullCertificateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UploadZoneOriginPullCertificate(uploadZoneOriginPullCertificateOptions *UploadZoneOriginPullCertificateOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		uploadZoneOriginPullCertificatePath := "/v1/testString/zones/testString/origin_tls_client_auth"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(uploadZoneOriginPullCertificatePath))
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

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"result": {"id": "0f405ba2-8c18-49eb-a30b-28b85427780f", "certificate": "-----BEGIN CERTIFICATE-----\n......\n-----END CERTIFICATE-----\n", "issuer": "GlobalSign", "signature": "SHA256WithRSA", "status": "active", "expires_on": "2100-01-01T05:20:00Z", "uploaded_on": "2100-01-01T05:20:00Z"}, "success": true, "errors": ["Errors"], "messages": ["Messages"]}`)
				}))
			})
			It(`Invoke UploadZoneOriginPullCertificate successfully with retries`, func() {
				authenticatedOriginPullApiService, serviceErr := authenticatedoriginpullapiv1.NewAuthenticatedOriginPullApiV1(&authenticatedoriginpullapiv1.AuthenticatedOriginPullApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(authenticatedOriginPullApiService).ToNot(BeNil())
				authenticatedOriginPullApiService.EnableRetries(0, 0)

				// Construct an instance of the UploadZoneOriginPullCertificateOptions model
				uploadZoneOriginPullCertificateOptionsModel := new(authenticatedoriginpullapiv1.UploadZoneOriginPullCertificateOptions)
				uploadZoneOriginPullCertificateOptionsModel.Certificate = core.StringPtr("-----BEGIN CERTIFICATE-----\n......\n-----END CERTIFICATE-----\n")
				uploadZoneOriginPullCertificateOptionsModel.PrivateKey = core.StringPtr("-----BEGIN RSA PRIVATE KEY-----\n......\n-----END RSA PRIVATE KEY-----\n")
				uploadZoneOriginPullCertificateOptionsModel.XCorrelationID = core.StringPtr("testString")
				uploadZoneOriginPullCertificateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := authenticatedOriginPullApiService.UploadZoneOriginPullCertificateWithContext(ctx, uploadZoneOriginPullCertificateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				authenticatedOriginPullApiService.DisableRetries()
				result, response, operationErr := authenticatedOriginPullApiService.UploadZoneOriginPullCertificate(uploadZoneOriginPullCertificateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = authenticatedOriginPullApiService.UploadZoneOriginPullCertificateWithContext(ctx, uploadZoneOriginPullCertificateOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(uploadZoneOriginPullCertificatePath))
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

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"result": {"id": "0f405ba2-8c18-49eb-a30b-28b85427780f", "certificate": "-----BEGIN CERTIFICATE-----\n......\n-----END CERTIFICATE-----\n", "issuer": "GlobalSign", "signature": "SHA256WithRSA", "status": "active", "expires_on": "2100-01-01T05:20:00Z", "uploaded_on": "2100-01-01T05:20:00Z"}, "success": true, "errors": ["Errors"], "messages": ["Messages"]}`)
				}))
			})
			It(`Invoke UploadZoneOriginPullCertificate successfully`, func() {
				authenticatedOriginPullApiService, serviceErr := authenticatedoriginpullapiv1.NewAuthenticatedOriginPullApiV1(&authenticatedoriginpullapiv1.AuthenticatedOriginPullApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(authenticatedOriginPullApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := authenticatedOriginPullApiService.UploadZoneOriginPullCertificate(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UploadZoneOriginPullCertificateOptions model
				uploadZoneOriginPullCertificateOptionsModel := new(authenticatedoriginpullapiv1.UploadZoneOriginPullCertificateOptions)
				uploadZoneOriginPullCertificateOptionsModel.Certificate = core.StringPtr("-----BEGIN CERTIFICATE-----\n......\n-----END CERTIFICATE-----\n")
				uploadZoneOriginPullCertificateOptionsModel.PrivateKey = core.StringPtr("-----BEGIN RSA PRIVATE KEY-----\n......\n-----END RSA PRIVATE KEY-----\n")
				uploadZoneOriginPullCertificateOptionsModel.XCorrelationID = core.StringPtr("testString")
				uploadZoneOriginPullCertificateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = authenticatedOriginPullApiService.UploadZoneOriginPullCertificate(uploadZoneOriginPullCertificateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UploadZoneOriginPullCertificate with error: Operation request error`, func() {
				authenticatedOriginPullApiService, serviceErr := authenticatedoriginpullapiv1.NewAuthenticatedOriginPullApiV1(&authenticatedoriginpullapiv1.AuthenticatedOriginPullApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(authenticatedOriginPullApiService).ToNot(BeNil())

				// Construct an instance of the UploadZoneOriginPullCertificateOptions model
				uploadZoneOriginPullCertificateOptionsModel := new(authenticatedoriginpullapiv1.UploadZoneOriginPullCertificateOptions)
				uploadZoneOriginPullCertificateOptionsModel.Certificate = core.StringPtr("-----BEGIN CERTIFICATE-----\n......\n-----END CERTIFICATE-----\n")
				uploadZoneOriginPullCertificateOptionsModel.PrivateKey = core.StringPtr("-----BEGIN RSA PRIVATE KEY-----\n......\n-----END RSA PRIVATE KEY-----\n")
				uploadZoneOriginPullCertificateOptionsModel.XCorrelationID = core.StringPtr("testString")
				uploadZoneOriginPullCertificateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := authenticatedOriginPullApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := authenticatedOriginPullApiService.UploadZoneOriginPullCertificate(uploadZoneOriginPullCertificateOptionsModel)
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
			It(`Invoke UploadZoneOriginPullCertificate successfully`, func() {
				authenticatedOriginPullApiService, serviceErr := authenticatedoriginpullapiv1.NewAuthenticatedOriginPullApiV1(&authenticatedoriginpullapiv1.AuthenticatedOriginPullApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(authenticatedOriginPullApiService).ToNot(BeNil())

				// Construct an instance of the UploadZoneOriginPullCertificateOptions model
				uploadZoneOriginPullCertificateOptionsModel := new(authenticatedoriginpullapiv1.UploadZoneOriginPullCertificateOptions)
				uploadZoneOriginPullCertificateOptionsModel.Certificate = core.StringPtr("-----BEGIN CERTIFICATE-----\n......\n-----END CERTIFICATE-----\n")
				uploadZoneOriginPullCertificateOptionsModel.PrivateKey = core.StringPtr("-----BEGIN RSA PRIVATE KEY-----\n......\n-----END RSA PRIVATE KEY-----\n")
				uploadZoneOriginPullCertificateOptionsModel.XCorrelationID = core.StringPtr("testString")
				uploadZoneOriginPullCertificateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := authenticatedOriginPullApiService.UploadZoneOriginPullCertificate(uploadZoneOriginPullCertificateOptionsModel)
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
	Describe(`GetZoneOriginPullCertificate(getZoneOriginPullCertificateOptions *GetZoneOriginPullCertificateOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getZoneOriginPullCertificatePath := "/v1/testString/zones/testString/origin_tls_client_auth/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getZoneOriginPullCertificatePath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetZoneOriginPullCertificate with error: Operation response processing error`, func() {
				authenticatedOriginPullApiService, serviceErr := authenticatedoriginpullapiv1.NewAuthenticatedOriginPullApiV1(&authenticatedoriginpullapiv1.AuthenticatedOriginPullApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(authenticatedOriginPullApiService).ToNot(BeNil())

				// Construct an instance of the GetZoneOriginPullCertificateOptions model
				getZoneOriginPullCertificateOptionsModel := new(authenticatedoriginpullapiv1.GetZoneOriginPullCertificateOptions)
				getZoneOriginPullCertificateOptionsModel.CertIdentifier = core.StringPtr("testString")
				getZoneOriginPullCertificateOptionsModel.XCorrelationID = core.StringPtr("testString")
				getZoneOriginPullCertificateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := authenticatedOriginPullApiService.GetZoneOriginPullCertificate(getZoneOriginPullCertificateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				authenticatedOriginPullApiService.EnableRetries(0, 0)
				result, response, operationErr = authenticatedOriginPullApiService.GetZoneOriginPullCertificate(getZoneOriginPullCertificateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetZoneOriginPullCertificate(getZoneOriginPullCertificateOptions *GetZoneOriginPullCertificateOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getZoneOriginPullCertificatePath := "/v1/testString/zones/testString/origin_tls_client_auth/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getZoneOriginPullCertificatePath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"result": {"id": "0f405ba2-8c18-49eb-a30b-28b85427780f", "certificate": "-----BEGIN CERTIFICATE-----\n......\n-----END CERTIFICATE-----\n", "issuer": "GlobalSign", "signature": "SHA256WithRSA", "status": "active", "expires_on": "2100-01-01T05:20:00Z", "uploaded_on": "2100-01-01T05:20:00Z"}, "success": true, "errors": ["Errors"], "messages": ["Messages"]}`)
				}))
			})
			It(`Invoke GetZoneOriginPullCertificate successfully with retries`, func() {
				authenticatedOriginPullApiService, serviceErr := authenticatedoriginpullapiv1.NewAuthenticatedOriginPullApiV1(&authenticatedoriginpullapiv1.AuthenticatedOriginPullApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(authenticatedOriginPullApiService).ToNot(BeNil())
				authenticatedOriginPullApiService.EnableRetries(0, 0)

				// Construct an instance of the GetZoneOriginPullCertificateOptions model
				getZoneOriginPullCertificateOptionsModel := new(authenticatedoriginpullapiv1.GetZoneOriginPullCertificateOptions)
				getZoneOriginPullCertificateOptionsModel.CertIdentifier = core.StringPtr("testString")
				getZoneOriginPullCertificateOptionsModel.XCorrelationID = core.StringPtr("testString")
				getZoneOriginPullCertificateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := authenticatedOriginPullApiService.GetZoneOriginPullCertificateWithContext(ctx, getZoneOriginPullCertificateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				authenticatedOriginPullApiService.DisableRetries()
				result, response, operationErr := authenticatedOriginPullApiService.GetZoneOriginPullCertificate(getZoneOriginPullCertificateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = authenticatedOriginPullApiService.GetZoneOriginPullCertificateWithContext(ctx, getZoneOriginPullCertificateOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getZoneOriginPullCertificatePath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"result": {"id": "0f405ba2-8c18-49eb-a30b-28b85427780f", "certificate": "-----BEGIN CERTIFICATE-----\n......\n-----END CERTIFICATE-----\n", "issuer": "GlobalSign", "signature": "SHA256WithRSA", "status": "active", "expires_on": "2100-01-01T05:20:00Z", "uploaded_on": "2100-01-01T05:20:00Z"}, "success": true, "errors": ["Errors"], "messages": ["Messages"]}`)
				}))
			})
			It(`Invoke GetZoneOriginPullCertificate successfully`, func() {
				authenticatedOriginPullApiService, serviceErr := authenticatedoriginpullapiv1.NewAuthenticatedOriginPullApiV1(&authenticatedoriginpullapiv1.AuthenticatedOriginPullApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(authenticatedOriginPullApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := authenticatedOriginPullApiService.GetZoneOriginPullCertificate(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetZoneOriginPullCertificateOptions model
				getZoneOriginPullCertificateOptionsModel := new(authenticatedoriginpullapiv1.GetZoneOriginPullCertificateOptions)
				getZoneOriginPullCertificateOptionsModel.CertIdentifier = core.StringPtr("testString")
				getZoneOriginPullCertificateOptionsModel.XCorrelationID = core.StringPtr("testString")
				getZoneOriginPullCertificateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = authenticatedOriginPullApiService.GetZoneOriginPullCertificate(getZoneOriginPullCertificateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetZoneOriginPullCertificate with error: Operation validation and request error`, func() {
				authenticatedOriginPullApiService, serviceErr := authenticatedoriginpullapiv1.NewAuthenticatedOriginPullApiV1(&authenticatedoriginpullapiv1.AuthenticatedOriginPullApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(authenticatedOriginPullApiService).ToNot(BeNil())

				// Construct an instance of the GetZoneOriginPullCertificateOptions model
				getZoneOriginPullCertificateOptionsModel := new(authenticatedoriginpullapiv1.GetZoneOriginPullCertificateOptions)
				getZoneOriginPullCertificateOptionsModel.CertIdentifier = core.StringPtr("testString")
				getZoneOriginPullCertificateOptionsModel.XCorrelationID = core.StringPtr("testString")
				getZoneOriginPullCertificateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := authenticatedOriginPullApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := authenticatedOriginPullApiService.GetZoneOriginPullCertificate(getZoneOriginPullCertificateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetZoneOriginPullCertificateOptions model with no property values
				getZoneOriginPullCertificateOptionsModelNew := new(authenticatedoriginpullapiv1.GetZoneOriginPullCertificateOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = authenticatedOriginPullApiService.GetZoneOriginPullCertificate(getZoneOriginPullCertificateOptionsModelNew)
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
			It(`Invoke GetZoneOriginPullCertificate successfully`, func() {
				authenticatedOriginPullApiService, serviceErr := authenticatedoriginpullapiv1.NewAuthenticatedOriginPullApiV1(&authenticatedoriginpullapiv1.AuthenticatedOriginPullApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(authenticatedOriginPullApiService).ToNot(BeNil())

				// Construct an instance of the GetZoneOriginPullCertificateOptions model
				getZoneOriginPullCertificateOptionsModel := new(authenticatedoriginpullapiv1.GetZoneOriginPullCertificateOptions)
				getZoneOriginPullCertificateOptionsModel.CertIdentifier = core.StringPtr("testString")
				getZoneOriginPullCertificateOptionsModel.XCorrelationID = core.StringPtr("testString")
				getZoneOriginPullCertificateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := authenticatedOriginPullApiService.GetZoneOriginPullCertificate(getZoneOriginPullCertificateOptionsModel)
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
	Describe(`DeleteZoneOriginPullCertificate(deleteZoneOriginPullCertificateOptions *DeleteZoneOriginPullCertificateOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		deleteZoneOriginPullCertificatePath := "/v1/testString/zones/testString/origin_tls_client_auth/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteZoneOriginPullCertificatePath))
					Expect(req.Method).To(Equal("DELETE"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke DeleteZoneOriginPullCertificate with error: Operation response processing error`, func() {
				authenticatedOriginPullApiService, serviceErr := authenticatedoriginpullapiv1.NewAuthenticatedOriginPullApiV1(&authenticatedoriginpullapiv1.AuthenticatedOriginPullApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(authenticatedOriginPullApiService).ToNot(BeNil())

				// Construct an instance of the DeleteZoneOriginPullCertificateOptions model
				deleteZoneOriginPullCertificateOptionsModel := new(authenticatedoriginpullapiv1.DeleteZoneOriginPullCertificateOptions)
				deleteZoneOriginPullCertificateOptionsModel.CertIdentifier = core.StringPtr("testString")
				deleteZoneOriginPullCertificateOptionsModel.XCorrelationID = core.StringPtr("testString")
				deleteZoneOriginPullCertificateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := authenticatedOriginPullApiService.DeleteZoneOriginPullCertificate(deleteZoneOriginPullCertificateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				authenticatedOriginPullApiService.EnableRetries(0, 0)
				result, response, operationErr = authenticatedOriginPullApiService.DeleteZoneOriginPullCertificate(deleteZoneOriginPullCertificateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteZoneOriginPullCertificate(deleteZoneOriginPullCertificateOptions *DeleteZoneOriginPullCertificateOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		deleteZoneOriginPullCertificatePath := "/v1/testString/zones/testString/origin_tls_client_auth/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteZoneOriginPullCertificatePath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"result": {"id": "0f405ba2-8c18-49eb-a30b-28b85427780f", "certificate": "-----BEGIN CERTIFICATE-----\n......\n-----END CERTIFICATE-----\n", "issuer": "GlobalSign", "signature": "SHA256WithRSA", "status": "active", "expires_on": "2100-01-01T05:20:00Z", "uploaded_on": "2100-01-01T05:20:00Z"}, "success": true, "errors": ["Errors"], "messages": ["Messages"]}`)
				}))
			})
			It(`Invoke DeleteZoneOriginPullCertificate successfully with retries`, func() {
				authenticatedOriginPullApiService, serviceErr := authenticatedoriginpullapiv1.NewAuthenticatedOriginPullApiV1(&authenticatedoriginpullapiv1.AuthenticatedOriginPullApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(authenticatedOriginPullApiService).ToNot(BeNil())
				authenticatedOriginPullApiService.EnableRetries(0, 0)

				// Construct an instance of the DeleteZoneOriginPullCertificateOptions model
				deleteZoneOriginPullCertificateOptionsModel := new(authenticatedoriginpullapiv1.DeleteZoneOriginPullCertificateOptions)
				deleteZoneOriginPullCertificateOptionsModel.CertIdentifier = core.StringPtr("testString")
				deleteZoneOriginPullCertificateOptionsModel.XCorrelationID = core.StringPtr("testString")
				deleteZoneOriginPullCertificateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := authenticatedOriginPullApiService.DeleteZoneOriginPullCertificateWithContext(ctx, deleteZoneOriginPullCertificateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				authenticatedOriginPullApiService.DisableRetries()
				result, response, operationErr := authenticatedOriginPullApiService.DeleteZoneOriginPullCertificate(deleteZoneOriginPullCertificateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = authenticatedOriginPullApiService.DeleteZoneOriginPullCertificateWithContext(ctx, deleteZoneOriginPullCertificateOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(deleteZoneOriginPullCertificatePath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"result": {"id": "0f405ba2-8c18-49eb-a30b-28b85427780f", "certificate": "-----BEGIN CERTIFICATE-----\n......\n-----END CERTIFICATE-----\n", "issuer": "GlobalSign", "signature": "SHA256WithRSA", "status": "active", "expires_on": "2100-01-01T05:20:00Z", "uploaded_on": "2100-01-01T05:20:00Z"}, "success": true, "errors": ["Errors"], "messages": ["Messages"]}`)
				}))
			})
			It(`Invoke DeleteZoneOriginPullCertificate successfully`, func() {
				authenticatedOriginPullApiService, serviceErr := authenticatedoriginpullapiv1.NewAuthenticatedOriginPullApiV1(&authenticatedoriginpullapiv1.AuthenticatedOriginPullApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(authenticatedOriginPullApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := authenticatedOriginPullApiService.DeleteZoneOriginPullCertificate(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DeleteZoneOriginPullCertificateOptions model
				deleteZoneOriginPullCertificateOptionsModel := new(authenticatedoriginpullapiv1.DeleteZoneOriginPullCertificateOptions)
				deleteZoneOriginPullCertificateOptionsModel.CertIdentifier = core.StringPtr("testString")
				deleteZoneOriginPullCertificateOptionsModel.XCorrelationID = core.StringPtr("testString")
				deleteZoneOriginPullCertificateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = authenticatedOriginPullApiService.DeleteZoneOriginPullCertificate(deleteZoneOriginPullCertificateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke DeleteZoneOriginPullCertificate with error: Operation validation and request error`, func() {
				authenticatedOriginPullApiService, serviceErr := authenticatedoriginpullapiv1.NewAuthenticatedOriginPullApiV1(&authenticatedoriginpullapiv1.AuthenticatedOriginPullApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(authenticatedOriginPullApiService).ToNot(BeNil())

				// Construct an instance of the DeleteZoneOriginPullCertificateOptions model
				deleteZoneOriginPullCertificateOptionsModel := new(authenticatedoriginpullapiv1.DeleteZoneOriginPullCertificateOptions)
				deleteZoneOriginPullCertificateOptionsModel.CertIdentifier = core.StringPtr("testString")
				deleteZoneOriginPullCertificateOptionsModel.XCorrelationID = core.StringPtr("testString")
				deleteZoneOriginPullCertificateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := authenticatedOriginPullApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := authenticatedOriginPullApiService.DeleteZoneOriginPullCertificate(deleteZoneOriginPullCertificateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the DeleteZoneOriginPullCertificateOptions model with no property values
				deleteZoneOriginPullCertificateOptionsModelNew := new(authenticatedoriginpullapiv1.DeleteZoneOriginPullCertificateOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = authenticatedOriginPullApiService.DeleteZoneOriginPullCertificate(deleteZoneOriginPullCertificateOptionsModelNew)
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
			It(`Invoke DeleteZoneOriginPullCertificate successfully`, func() {
				authenticatedOriginPullApiService, serviceErr := authenticatedoriginpullapiv1.NewAuthenticatedOriginPullApiV1(&authenticatedoriginpullapiv1.AuthenticatedOriginPullApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(authenticatedOriginPullApiService).ToNot(BeNil())

				// Construct an instance of the DeleteZoneOriginPullCertificateOptions model
				deleteZoneOriginPullCertificateOptionsModel := new(authenticatedoriginpullapiv1.DeleteZoneOriginPullCertificateOptions)
				deleteZoneOriginPullCertificateOptionsModel.CertIdentifier = core.StringPtr("testString")
				deleteZoneOriginPullCertificateOptionsModel.XCorrelationID = core.StringPtr("testString")
				deleteZoneOriginPullCertificateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := authenticatedOriginPullApiService.DeleteZoneOriginPullCertificate(deleteZoneOriginPullCertificateOptionsModel)
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
	Describe(`SetHostnameOriginPullSettings(setHostnameOriginPullSettingsOptions *SetHostnameOriginPullSettingsOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		setHostnameOriginPullSettingsPath := "/v1/testString/zones/testString/origin_tls_client_auth/hostnames"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(setHostnameOriginPullSettingsPath))
					Expect(req.Method).To(Equal("PUT"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke SetHostnameOriginPullSettings with error: Operation response processing error`, func() {
				authenticatedOriginPullApiService, serviceErr := authenticatedoriginpullapiv1.NewAuthenticatedOriginPullApiV1(&authenticatedoriginpullapiv1.AuthenticatedOriginPullApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(authenticatedOriginPullApiService).ToNot(BeNil())

				// Construct an instance of the HostnameOriginPullSettings model
				hostnameOriginPullSettingsModel := new(authenticatedoriginpullapiv1.HostnameOriginPullSettings)
				hostnameOriginPullSettingsModel.Hostname = core.StringPtr("app.example.com")
				hostnameOriginPullSettingsModel.CertID = core.StringPtr("2458ce5a-0c35-4c7f-82c7-8e9487d3ff60")
				hostnameOriginPullSettingsModel.Enabled = core.BoolPtr(true)

				// Construct an instance of the SetHostnameOriginPullSettingsOptions model
				setHostnameOriginPullSettingsOptionsModel := new(authenticatedoriginpullapiv1.SetHostnameOriginPullSettingsOptions)
				setHostnameOriginPullSettingsOptionsModel.Config = []authenticatedoriginpullapiv1.HostnameOriginPullSettings{*hostnameOriginPullSettingsModel}
				setHostnameOriginPullSettingsOptionsModel.XCorrelationID = core.StringPtr("testString")
				setHostnameOriginPullSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := authenticatedOriginPullApiService.SetHostnameOriginPullSettings(setHostnameOriginPullSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				authenticatedOriginPullApiService.EnableRetries(0, 0)
				result, response, operationErr = authenticatedOriginPullApiService.SetHostnameOriginPullSettings(setHostnameOriginPullSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`SetHostnameOriginPullSettings(setHostnameOriginPullSettingsOptions *SetHostnameOriginPullSettingsOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		setHostnameOriginPullSettingsPath := "/v1/testString/zones/testString/origin_tls_client_auth/hostnames"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(setHostnameOriginPullSettingsPath))
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

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"result": [{"hostname": "app.example.com", "cert_id": "2458ce5a-0c35-4c7f-82c7-8e9487d3ff60", "enabled": true, "status": "active", "created_at": "2100-01-01T05:20:00.000Z", "updated_at": "2100-01-01T05:20:00.000Z", "cert_status": "active", "issuer": "GlobalSign", "signature": "SHA256WithRSA", "serial_number": "6743787633689793699141714808227354901", "certificate": "-----BEGIN CERTIFICATE-----\n......\n-----END CERTIFICATE-----\n", "cert_uploaded_on": "2019-10-28T18:11:23.374Z", "cert_updated_at": "2100-01-01T05:20:00.000Z", "expires_on": "2100-01-01T05:20:00.000Z"}], "success": true, "errors": ["Errors"], "messages": ["Messages"]}`)
				}))
			})
			It(`Invoke SetHostnameOriginPullSettings successfully with retries`, func() {
				authenticatedOriginPullApiService, serviceErr := authenticatedoriginpullapiv1.NewAuthenticatedOriginPullApiV1(&authenticatedoriginpullapiv1.AuthenticatedOriginPullApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(authenticatedOriginPullApiService).ToNot(BeNil())
				authenticatedOriginPullApiService.EnableRetries(0, 0)

				// Construct an instance of the HostnameOriginPullSettings model
				hostnameOriginPullSettingsModel := new(authenticatedoriginpullapiv1.HostnameOriginPullSettings)
				hostnameOriginPullSettingsModel.Hostname = core.StringPtr("app.example.com")
				hostnameOriginPullSettingsModel.CertID = core.StringPtr("2458ce5a-0c35-4c7f-82c7-8e9487d3ff60")
				hostnameOriginPullSettingsModel.Enabled = core.BoolPtr(true)

				// Construct an instance of the SetHostnameOriginPullSettingsOptions model
				setHostnameOriginPullSettingsOptionsModel := new(authenticatedoriginpullapiv1.SetHostnameOriginPullSettingsOptions)
				setHostnameOriginPullSettingsOptionsModel.Config = []authenticatedoriginpullapiv1.HostnameOriginPullSettings{*hostnameOriginPullSettingsModel}
				setHostnameOriginPullSettingsOptionsModel.XCorrelationID = core.StringPtr("testString")
				setHostnameOriginPullSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := authenticatedOriginPullApiService.SetHostnameOriginPullSettingsWithContext(ctx, setHostnameOriginPullSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				authenticatedOriginPullApiService.DisableRetries()
				result, response, operationErr := authenticatedOriginPullApiService.SetHostnameOriginPullSettings(setHostnameOriginPullSettingsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = authenticatedOriginPullApiService.SetHostnameOriginPullSettingsWithContext(ctx, setHostnameOriginPullSettingsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(setHostnameOriginPullSettingsPath))
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

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"result": [{"hostname": "app.example.com", "cert_id": "2458ce5a-0c35-4c7f-82c7-8e9487d3ff60", "enabled": true, "status": "active", "created_at": "2100-01-01T05:20:00.000Z", "updated_at": "2100-01-01T05:20:00.000Z", "cert_status": "active", "issuer": "GlobalSign", "signature": "SHA256WithRSA", "serial_number": "6743787633689793699141714808227354901", "certificate": "-----BEGIN CERTIFICATE-----\n......\n-----END CERTIFICATE-----\n", "cert_uploaded_on": "2019-10-28T18:11:23.374Z", "cert_updated_at": "2100-01-01T05:20:00.000Z", "expires_on": "2100-01-01T05:20:00.000Z"}], "success": true, "errors": ["Errors"], "messages": ["Messages"]}`)
				}))
			})
			It(`Invoke SetHostnameOriginPullSettings successfully`, func() {
				authenticatedOriginPullApiService, serviceErr := authenticatedoriginpullapiv1.NewAuthenticatedOriginPullApiV1(&authenticatedoriginpullapiv1.AuthenticatedOriginPullApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(authenticatedOriginPullApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := authenticatedOriginPullApiService.SetHostnameOriginPullSettings(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the HostnameOriginPullSettings model
				hostnameOriginPullSettingsModel := new(authenticatedoriginpullapiv1.HostnameOriginPullSettings)
				hostnameOriginPullSettingsModel.Hostname = core.StringPtr("app.example.com")
				hostnameOriginPullSettingsModel.CertID = core.StringPtr("2458ce5a-0c35-4c7f-82c7-8e9487d3ff60")
				hostnameOriginPullSettingsModel.Enabled = core.BoolPtr(true)

				// Construct an instance of the SetHostnameOriginPullSettingsOptions model
				setHostnameOriginPullSettingsOptionsModel := new(authenticatedoriginpullapiv1.SetHostnameOriginPullSettingsOptions)
				setHostnameOriginPullSettingsOptionsModel.Config = []authenticatedoriginpullapiv1.HostnameOriginPullSettings{*hostnameOriginPullSettingsModel}
				setHostnameOriginPullSettingsOptionsModel.XCorrelationID = core.StringPtr("testString")
				setHostnameOriginPullSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = authenticatedOriginPullApiService.SetHostnameOriginPullSettings(setHostnameOriginPullSettingsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke SetHostnameOriginPullSettings with error: Operation request error`, func() {
				authenticatedOriginPullApiService, serviceErr := authenticatedoriginpullapiv1.NewAuthenticatedOriginPullApiV1(&authenticatedoriginpullapiv1.AuthenticatedOriginPullApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(authenticatedOriginPullApiService).ToNot(BeNil())

				// Construct an instance of the HostnameOriginPullSettings model
				hostnameOriginPullSettingsModel := new(authenticatedoriginpullapiv1.HostnameOriginPullSettings)
				hostnameOriginPullSettingsModel.Hostname = core.StringPtr("app.example.com")
				hostnameOriginPullSettingsModel.CertID = core.StringPtr("2458ce5a-0c35-4c7f-82c7-8e9487d3ff60")
				hostnameOriginPullSettingsModel.Enabled = core.BoolPtr(true)

				// Construct an instance of the SetHostnameOriginPullSettingsOptions model
				setHostnameOriginPullSettingsOptionsModel := new(authenticatedoriginpullapiv1.SetHostnameOriginPullSettingsOptions)
				setHostnameOriginPullSettingsOptionsModel.Config = []authenticatedoriginpullapiv1.HostnameOriginPullSettings{*hostnameOriginPullSettingsModel}
				setHostnameOriginPullSettingsOptionsModel.XCorrelationID = core.StringPtr("testString")
				setHostnameOriginPullSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := authenticatedOriginPullApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := authenticatedOriginPullApiService.SetHostnameOriginPullSettings(setHostnameOriginPullSettingsOptionsModel)
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
			It(`Invoke SetHostnameOriginPullSettings successfully`, func() {
				authenticatedOriginPullApiService, serviceErr := authenticatedoriginpullapiv1.NewAuthenticatedOriginPullApiV1(&authenticatedoriginpullapiv1.AuthenticatedOriginPullApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(authenticatedOriginPullApiService).ToNot(BeNil())

				// Construct an instance of the HostnameOriginPullSettings model
				hostnameOriginPullSettingsModel := new(authenticatedoriginpullapiv1.HostnameOriginPullSettings)
				hostnameOriginPullSettingsModel.Hostname = core.StringPtr("app.example.com")
				hostnameOriginPullSettingsModel.CertID = core.StringPtr("2458ce5a-0c35-4c7f-82c7-8e9487d3ff60")
				hostnameOriginPullSettingsModel.Enabled = core.BoolPtr(true)

				// Construct an instance of the SetHostnameOriginPullSettingsOptions model
				setHostnameOriginPullSettingsOptionsModel := new(authenticatedoriginpullapiv1.SetHostnameOriginPullSettingsOptions)
				setHostnameOriginPullSettingsOptionsModel.Config = []authenticatedoriginpullapiv1.HostnameOriginPullSettings{*hostnameOriginPullSettingsModel}
				setHostnameOriginPullSettingsOptionsModel.XCorrelationID = core.StringPtr("testString")
				setHostnameOriginPullSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := authenticatedOriginPullApiService.SetHostnameOriginPullSettings(setHostnameOriginPullSettingsOptionsModel)
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
	Describe(`GetHostnameOriginPullSettings(getHostnameOriginPullSettingsOptions *GetHostnameOriginPullSettingsOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getHostnameOriginPullSettingsPath := "/v1/testString/zones/testString/origin_tls_client_auth/hostnames/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getHostnameOriginPullSettingsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetHostnameOriginPullSettings with error: Operation response processing error`, func() {
				authenticatedOriginPullApiService, serviceErr := authenticatedoriginpullapiv1.NewAuthenticatedOriginPullApiV1(&authenticatedoriginpullapiv1.AuthenticatedOriginPullApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(authenticatedOriginPullApiService).ToNot(BeNil())

				// Construct an instance of the GetHostnameOriginPullSettingsOptions model
				getHostnameOriginPullSettingsOptionsModel := new(authenticatedoriginpullapiv1.GetHostnameOriginPullSettingsOptions)
				getHostnameOriginPullSettingsOptionsModel.Hostname = core.StringPtr("testString")
				getHostnameOriginPullSettingsOptionsModel.XCorrelationID = core.StringPtr("testString")
				getHostnameOriginPullSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := authenticatedOriginPullApiService.GetHostnameOriginPullSettings(getHostnameOriginPullSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				authenticatedOriginPullApiService.EnableRetries(0, 0)
				result, response, operationErr = authenticatedOriginPullApiService.GetHostnameOriginPullSettings(getHostnameOriginPullSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetHostnameOriginPullSettings(getHostnameOriginPullSettingsOptions *GetHostnameOriginPullSettingsOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getHostnameOriginPullSettingsPath := "/v1/testString/zones/testString/origin_tls_client_auth/hostnames/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getHostnameOriginPullSettingsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"result": {"hostname": "app.example.com", "cert_id": "2458ce5a-0c35-4c7f-82c7-8e9487d3ff60", "enabled": true, "status": "active", "created_at": "2100-01-01T05:20:00.000Z", "updated_at": "2100-01-01T05:20:00.000Z", "cert_status": "active", "issuer": "GlobalSign", "signature": "SHA256WithRSA", "serial_number": "6743787633689793699141714808227354901", "certificate": "-----BEGIN CERTIFICATE-----\n......\n-----END CERTIFICATE-----\n", "cert_uploaded_on": "2019-10-28T18:11:23.374Z", "cert_updated_at": "2100-01-01T05:20:00.000Z", "expires_on": "2100-01-01T05:20:00.000Z"}, "success": true, "errors": ["Errors"], "messages": ["Messages"]}`)
				}))
			})
			It(`Invoke GetHostnameOriginPullSettings successfully with retries`, func() {
				authenticatedOriginPullApiService, serviceErr := authenticatedoriginpullapiv1.NewAuthenticatedOriginPullApiV1(&authenticatedoriginpullapiv1.AuthenticatedOriginPullApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(authenticatedOriginPullApiService).ToNot(BeNil())
				authenticatedOriginPullApiService.EnableRetries(0, 0)

				// Construct an instance of the GetHostnameOriginPullSettingsOptions model
				getHostnameOriginPullSettingsOptionsModel := new(authenticatedoriginpullapiv1.GetHostnameOriginPullSettingsOptions)
				getHostnameOriginPullSettingsOptionsModel.Hostname = core.StringPtr("testString")
				getHostnameOriginPullSettingsOptionsModel.XCorrelationID = core.StringPtr("testString")
				getHostnameOriginPullSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := authenticatedOriginPullApiService.GetHostnameOriginPullSettingsWithContext(ctx, getHostnameOriginPullSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				authenticatedOriginPullApiService.DisableRetries()
				result, response, operationErr := authenticatedOriginPullApiService.GetHostnameOriginPullSettings(getHostnameOriginPullSettingsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = authenticatedOriginPullApiService.GetHostnameOriginPullSettingsWithContext(ctx, getHostnameOriginPullSettingsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getHostnameOriginPullSettingsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"result": {"hostname": "app.example.com", "cert_id": "2458ce5a-0c35-4c7f-82c7-8e9487d3ff60", "enabled": true, "status": "active", "created_at": "2100-01-01T05:20:00.000Z", "updated_at": "2100-01-01T05:20:00.000Z", "cert_status": "active", "issuer": "GlobalSign", "signature": "SHA256WithRSA", "serial_number": "6743787633689793699141714808227354901", "certificate": "-----BEGIN CERTIFICATE-----\n......\n-----END CERTIFICATE-----\n", "cert_uploaded_on": "2019-10-28T18:11:23.374Z", "cert_updated_at": "2100-01-01T05:20:00.000Z", "expires_on": "2100-01-01T05:20:00.000Z"}, "success": true, "errors": ["Errors"], "messages": ["Messages"]}`)
				}))
			})
			It(`Invoke GetHostnameOriginPullSettings successfully`, func() {
				authenticatedOriginPullApiService, serviceErr := authenticatedoriginpullapiv1.NewAuthenticatedOriginPullApiV1(&authenticatedoriginpullapiv1.AuthenticatedOriginPullApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(authenticatedOriginPullApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := authenticatedOriginPullApiService.GetHostnameOriginPullSettings(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetHostnameOriginPullSettingsOptions model
				getHostnameOriginPullSettingsOptionsModel := new(authenticatedoriginpullapiv1.GetHostnameOriginPullSettingsOptions)
				getHostnameOriginPullSettingsOptionsModel.Hostname = core.StringPtr("testString")
				getHostnameOriginPullSettingsOptionsModel.XCorrelationID = core.StringPtr("testString")
				getHostnameOriginPullSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = authenticatedOriginPullApiService.GetHostnameOriginPullSettings(getHostnameOriginPullSettingsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetHostnameOriginPullSettings with error: Operation validation and request error`, func() {
				authenticatedOriginPullApiService, serviceErr := authenticatedoriginpullapiv1.NewAuthenticatedOriginPullApiV1(&authenticatedoriginpullapiv1.AuthenticatedOriginPullApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(authenticatedOriginPullApiService).ToNot(BeNil())

				// Construct an instance of the GetHostnameOriginPullSettingsOptions model
				getHostnameOriginPullSettingsOptionsModel := new(authenticatedoriginpullapiv1.GetHostnameOriginPullSettingsOptions)
				getHostnameOriginPullSettingsOptionsModel.Hostname = core.StringPtr("testString")
				getHostnameOriginPullSettingsOptionsModel.XCorrelationID = core.StringPtr("testString")
				getHostnameOriginPullSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := authenticatedOriginPullApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := authenticatedOriginPullApiService.GetHostnameOriginPullSettings(getHostnameOriginPullSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetHostnameOriginPullSettingsOptions model with no property values
				getHostnameOriginPullSettingsOptionsModelNew := new(authenticatedoriginpullapiv1.GetHostnameOriginPullSettingsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = authenticatedOriginPullApiService.GetHostnameOriginPullSettings(getHostnameOriginPullSettingsOptionsModelNew)
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
			It(`Invoke GetHostnameOriginPullSettings successfully`, func() {
				authenticatedOriginPullApiService, serviceErr := authenticatedoriginpullapiv1.NewAuthenticatedOriginPullApiV1(&authenticatedoriginpullapiv1.AuthenticatedOriginPullApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(authenticatedOriginPullApiService).ToNot(BeNil())

				// Construct an instance of the GetHostnameOriginPullSettingsOptions model
				getHostnameOriginPullSettingsOptionsModel := new(authenticatedoriginpullapiv1.GetHostnameOriginPullSettingsOptions)
				getHostnameOriginPullSettingsOptionsModel.Hostname = core.StringPtr("testString")
				getHostnameOriginPullSettingsOptionsModel.XCorrelationID = core.StringPtr("testString")
				getHostnameOriginPullSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := authenticatedOriginPullApiService.GetHostnameOriginPullSettings(getHostnameOriginPullSettingsOptionsModel)
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
	Describe(`UploadHostnameOriginPullCertificate(uploadHostnameOriginPullCertificateOptions *UploadHostnameOriginPullCertificateOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		uploadHostnameOriginPullCertificatePath := "/v1/testString/zones/testString/origin_tls_client_auth/hostnames/certificates"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(uploadHostnameOriginPullCertificatePath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UploadHostnameOriginPullCertificate with error: Operation response processing error`, func() {
				authenticatedOriginPullApiService, serviceErr := authenticatedoriginpullapiv1.NewAuthenticatedOriginPullApiV1(&authenticatedoriginpullapiv1.AuthenticatedOriginPullApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(authenticatedOriginPullApiService).ToNot(BeNil())

				// Construct an instance of the UploadHostnameOriginPullCertificateOptions model
				uploadHostnameOriginPullCertificateOptionsModel := new(authenticatedoriginpullapiv1.UploadHostnameOriginPullCertificateOptions)
				uploadHostnameOriginPullCertificateOptionsModel.Certificate = core.StringPtr("-----BEGIN CERTIFICATE-----\n......\n-----END CERTIFICATE-----\n")
				uploadHostnameOriginPullCertificateOptionsModel.PrivateKey = core.StringPtr("-----BEGIN RSA PRIVATE KEY-----\n......\n-----END RSA PRIVATE KEY-----\n")
				uploadHostnameOriginPullCertificateOptionsModel.XCorrelationID = core.StringPtr("testString")
				uploadHostnameOriginPullCertificateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := authenticatedOriginPullApiService.UploadHostnameOriginPullCertificate(uploadHostnameOriginPullCertificateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				authenticatedOriginPullApiService.EnableRetries(0, 0)
				result, response, operationErr = authenticatedOriginPullApiService.UploadHostnameOriginPullCertificate(uploadHostnameOriginPullCertificateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UploadHostnameOriginPullCertificate(uploadHostnameOriginPullCertificateOptions *UploadHostnameOriginPullCertificateOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		uploadHostnameOriginPullCertificatePath := "/v1/testString/zones/testString/origin_tls_client_auth/hostnames/certificates"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(uploadHostnameOriginPullCertificatePath))
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

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"result": {"id": "0f405ba2-8c18-49eb-a30b-28b85427780f", "certificate": "-----BEGIN CERTIFICATE-----\n......\n-----END CERTIFICATE-----\n", "issuer": "GlobalSign", "signature": "SHA256WithRSA", "serial_number": "6743787633689793699141714808227354901", "status": "active", "expires_on": "2100-01-01T05:20:00Z", "uploaded_on": "2100-01-01T05:20:00Z"}, "success": true, "errors": ["Errors"], "messages": ["Messages"]}`)
				}))
			})
			It(`Invoke UploadHostnameOriginPullCertificate successfully with retries`, func() {
				authenticatedOriginPullApiService, serviceErr := authenticatedoriginpullapiv1.NewAuthenticatedOriginPullApiV1(&authenticatedoriginpullapiv1.AuthenticatedOriginPullApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(authenticatedOriginPullApiService).ToNot(BeNil())
				authenticatedOriginPullApiService.EnableRetries(0, 0)

				// Construct an instance of the UploadHostnameOriginPullCertificateOptions model
				uploadHostnameOriginPullCertificateOptionsModel := new(authenticatedoriginpullapiv1.UploadHostnameOriginPullCertificateOptions)
				uploadHostnameOriginPullCertificateOptionsModel.Certificate = core.StringPtr("-----BEGIN CERTIFICATE-----\n......\n-----END CERTIFICATE-----\n")
				uploadHostnameOriginPullCertificateOptionsModel.PrivateKey = core.StringPtr("-----BEGIN RSA PRIVATE KEY-----\n......\n-----END RSA PRIVATE KEY-----\n")
				uploadHostnameOriginPullCertificateOptionsModel.XCorrelationID = core.StringPtr("testString")
				uploadHostnameOriginPullCertificateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := authenticatedOriginPullApiService.UploadHostnameOriginPullCertificateWithContext(ctx, uploadHostnameOriginPullCertificateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				authenticatedOriginPullApiService.DisableRetries()
				result, response, operationErr := authenticatedOriginPullApiService.UploadHostnameOriginPullCertificate(uploadHostnameOriginPullCertificateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = authenticatedOriginPullApiService.UploadHostnameOriginPullCertificateWithContext(ctx, uploadHostnameOriginPullCertificateOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(uploadHostnameOriginPullCertificatePath))
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

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"result": {"id": "0f405ba2-8c18-49eb-a30b-28b85427780f", "certificate": "-----BEGIN CERTIFICATE-----\n......\n-----END CERTIFICATE-----\n", "issuer": "GlobalSign", "signature": "SHA256WithRSA", "serial_number": "6743787633689793699141714808227354901", "status": "active", "expires_on": "2100-01-01T05:20:00Z", "uploaded_on": "2100-01-01T05:20:00Z"}, "success": true, "errors": ["Errors"], "messages": ["Messages"]}`)
				}))
			})
			It(`Invoke UploadHostnameOriginPullCertificate successfully`, func() {
				authenticatedOriginPullApiService, serviceErr := authenticatedoriginpullapiv1.NewAuthenticatedOriginPullApiV1(&authenticatedoriginpullapiv1.AuthenticatedOriginPullApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(authenticatedOriginPullApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := authenticatedOriginPullApiService.UploadHostnameOriginPullCertificate(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UploadHostnameOriginPullCertificateOptions model
				uploadHostnameOriginPullCertificateOptionsModel := new(authenticatedoriginpullapiv1.UploadHostnameOriginPullCertificateOptions)
				uploadHostnameOriginPullCertificateOptionsModel.Certificate = core.StringPtr("-----BEGIN CERTIFICATE-----\n......\n-----END CERTIFICATE-----\n")
				uploadHostnameOriginPullCertificateOptionsModel.PrivateKey = core.StringPtr("-----BEGIN RSA PRIVATE KEY-----\n......\n-----END RSA PRIVATE KEY-----\n")
				uploadHostnameOriginPullCertificateOptionsModel.XCorrelationID = core.StringPtr("testString")
				uploadHostnameOriginPullCertificateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = authenticatedOriginPullApiService.UploadHostnameOriginPullCertificate(uploadHostnameOriginPullCertificateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UploadHostnameOriginPullCertificate with error: Operation request error`, func() {
				authenticatedOriginPullApiService, serviceErr := authenticatedoriginpullapiv1.NewAuthenticatedOriginPullApiV1(&authenticatedoriginpullapiv1.AuthenticatedOriginPullApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(authenticatedOriginPullApiService).ToNot(BeNil())

				// Construct an instance of the UploadHostnameOriginPullCertificateOptions model
				uploadHostnameOriginPullCertificateOptionsModel := new(authenticatedoriginpullapiv1.UploadHostnameOriginPullCertificateOptions)
				uploadHostnameOriginPullCertificateOptionsModel.Certificate = core.StringPtr("-----BEGIN CERTIFICATE-----\n......\n-----END CERTIFICATE-----\n")
				uploadHostnameOriginPullCertificateOptionsModel.PrivateKey = core.StringPtr("-----BEGIN RSA PRIVATE KEY-----\n......\n-----END RSA PRIVATE KEY-----\n")
				uploadHostnameOriginPullCertificateOptionsModel.XCorrelationID = core.StringPtr("testString")
				uploadHostnameOriginPullCertificateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := authenticatedOriginPullApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := authenticatedOriginPullApiService.UploadHostnameOriginPullCertificate(uploadHostnameOriginPullCertificateOptionsModel)
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
			It(`Invoke UploadHostnameOriginPullCertificate successfully`, func() {
				authenticatedOriginPullApiService, serviceErr := authenticatedoriginpullapiv1.NewAuthenticatedOriginPullApiV1(&authenticatedoriginpullapiv1.AuthenticatedOriginPullApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(authenticatedOriginPullApiService).ToNot(BeNil())

				// Construct an instance of the UploadHostnameOriginPullCertificateOptions model
				uploadHostnameOriginPullCertificateOptionsModel := new(authenticatedoriginpullapiv1.UploadHostnameOriginPullCertificateOptions)
				uploadHostnameOriginPullCertificateOptionsModel.Certificate = core.StringPtr("-----BEGIN CERTIFICATE-----\n......\n-----END CERTIFICATE-----\n")
				uploadHostnameOriginPullCertificateOptionsModel.PrivateKey = core.StringPtr("-----BEGIN RSA PRIVATE KEY-----\n......\n-----END RSA PRIVATE KEY-----\n")
				uploadHostnameOriginPullCertificateOptionsModel.XCorrelationID = core.StringPtr("testString")
				uploadHostnameOriginPullCertificateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := authenticatedOriginPullApiService.UploadHostnameOriginPullCertificate(uploadHostnameOriginPullCertificateOptionsModel)
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
	Describe(`GetHostnameOriginPullCertificate(getHostnameOriginPullCertificateOptions *GetHostnameOriginPullCertificateOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getHostnameOriginPullCertificatePath := "/v1/testString/zones/testString/origin_tls_client_auth/hostnames/certificates/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getHostnameOriginPullCertificatePath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetHostnameOriginPullCertificate with error: Operation response processing error`, func() {
				authenticatedOriginPullApiService, serviceErr := authenticatedoriginpullapiv1.NewAuthenticatedOriginPullApiV1(&authenticatedoriginpullapiv1.AuthenticatedOriginPullApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(authenticatedOriginPullApiService).ToNot(BeNil())

				// Construct an instance of the GetHostnameOriginPullCertificateOptions model
				getHostnameOriginPullCertificateOptionsModel := new(authenticatedoriginpullapiv1.GetHostnameOriginPullCertificateOptions)
				getHostnameOriginPullCertificateOptionsModel.CertIdentifier = core.StringPtr("testString")
				getHostnameOriginPullCertificateOptionsModel.XCorrelationID = core.StringPtr("testString")
				getHostnameOriginPullCertificateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := authenticatedOriginPullApiService.GetHostnameOriginPullCertificate(getHostnameOriginPullCertificateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				authenticatedOriginPullApiService.EnableRetries(0, 0)
				result, response, operationErr = authenticatedOriginPullApiService.GetHostnameOriginPullCertificate(getHostnameOriginPullCertificateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetHostnameOriginPullCertificate(getHostnameOriginPullCertificateOptions *GetHostnameOriginPullCertificateOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getHostnameOriginPullCertificatePath := "/v1/testString/zones/testString/origin_tls_client_auth/hostnames/certificates/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getHostnameOriginPullCertificatePath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"result": {"id": "0f405ba2-8c18-49eb-a30b-28b85427780f", "certificate": "-----BEGIN CERTIFICATE-----\n......\n-----END CERTIFICATE-----\n", "issuer": "GlobalSign", "signature": "SHA256WithRSA", "serial_number": "6743787633689793699141714808227354901", "status": "active", "expires_on": "2100-01-01T05:20:00Z", "uploaded_on": "2100-01-01T05:20:00Z"}, "success": true, "errors": ["Errors"], "messages": ["Messages"]}`)
				}))
			})
			It(`Invoke GetHostnameOriginPullCertificate successfully with retries`, func() {
				authenticatedOriginPullApiService, serviceErr := authenticatedoriginpullapiv1.NewAuthenticatedOriginPullApiV1(&authenticatedoriginpullapiv1.AuthenticatedOriginPullApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(authenticatedOriginPullApiService).ToNot(BeNil())
				authenticatedOriginPullApiService.EnableRetries(0, 0)

				// Construct an instance of the GetHostnameOriginPullCertificateOptions model
				getHostnameOriginPullCertificateOptionsModel := new(authenticatedoriginpullapiv1.GetHostnameOriginPullCertificateOptions)
				getHostnameOriginPullCertificateOptionsModel.CertIdentifier = core.StringPtr("testString")
				getHostnameOriginPullCertificateOptionsModel.XCorrelationID = core.StringPtr("testString")
				getHostnameOriginPullCertificateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := authenticatedOriginPullApiService.GetHostnameOriginPullCertificateWithContext(ctx, getHostnameOriginPullCertificateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				authenticatedOriginPullApiService.DisableRetries()
				result, response, operationErr := authenticatedOriginPullApiService.GetHostnameOriginPullCertificate(getHostnameOriginPullCertificateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = authenticatedOriginPullApiService.GetHostnameOriginPullCertificateWithContext(ctx, getHostnameOriginPullCertificateOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getHostnameOriginPullCertificatePath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"result": {"id": "0f405ba2-8c18-49eb-a30b-28b85427780f", "certificate": "-----BEGIN CERTIFICATE-----\n......\n-----END CERTIFICATE-----\n", "issuer": "GlobalSign", "signature": "SHA256WithRSA", "serial_number": "6743787633689793699141714808227354901", "status": "active", "expires_on": "2100-01-01T05:20:00Z", "uploaded_on": "2100-01-01T05:20:00Z"}, "success": true, "errors": ["Errors"], "messages": ["Messages"]}`)
				}))
			})
			It(`Invoke GetHostnameOriginPullCertificate successfully`, func() {
				authenticatedOriginPullApiService, serviceErr := authenticatedoriginpullapiv1.NewAuthenticatedOriginPullApiV1(&authenticatedoriginpullapiv1.AuthenticatedOriginPullApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(authenticatedOriginPullApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := authenticatedOriginPullApiService.GetHostnameOriginPullCertificate(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetHostnameOriginPullCertificateOptions model
				getHostnameOriginPullCertificateOptionsModel := new(authenticatedoriginpullapiv1.GetHostnameOriginPullCertificateOptions)
				getHostnameOriginPullCertificateOptionsModel.CertIdentifier = core.StringPtr("testString")
				getHostnameOriginPullCertificateOptionsModel.XCorrelationID = core.StringPtr("testString")
				getHostnameOriginPullCertificateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = authenticatedOriginPullApiService.GetHostnameOriginPullCertificate(getHostnameOriginPullCertificateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetHostnameOriginPullCertificate with error: Operation validation and request error`, func() {
				authenticatedOriginPullApiService, serviceErr := authenticatedoriginpullapiv1.NewAuthenticatedOriginPullApiV1(&authenticatedoriginpullapiv1.AuthenticatedOriginPullApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(authenticatedOriginPullApiService).ToNot(BeNil())

				// Construct an instance of the GetHostnameOriginPullCertificateOptions model
				getHostnameOriginPullCertificateOptionsModel := new(authenticatedoriginpullapiv1.GetHostnameOriginPullCertificateOptions)
				getHostnameOriginPullCertificateOptionsModel.CertIdentifier = core.StringPtr("testString")
				getHostnameOriginPullCertificateOptionsModel.XCorrelationID = core.StringPtr("testString")
				getHostnameOriginPullCertificateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := authenticatedOriginPullApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := authenticatedOriginPullApiService.GetHostnameOriginPullCertificate(getHostnameOriginPullCertificateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetHostnameOriginPullCertificateOptions model with no property values
				getHostnameOriginPullCertificateOptionsModelNew := new(authenticatedoriginpullapiv1.GetHostnameOriginPullCertificateOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = authenticatedOriginPullApiService.GetHostnameOriginPullCertificate(getHostnameOriginPullCertificateOptionsModelNew)
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
			It(`Invoke GetHostnameOriginPullCertificate successfully`, func() {
				authenticatedOriginPullApiService, serviceErr := authenticatedoriginpullapiv1.NewAuthenticatedOriginPullApiV1(&authenticatedoriginpullapiv1.AuthenticatedOriginPullApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(authenticatedOriginPullApiService).ToNot(BeNil())

				// Construct an instance of the GetHostnameOriginPullCertificateOptions model
				getHostnameOriginPullCertificateOptionsModel := new(authenticatedoriginpullapiv1.GetHostnameOriginPullCertificateOptions)
				getHostnameOriginPullCertificateOptionsModel.CertIdentifier = core.StringPtr("testString")
				getHostnameOriginPullCertificateOptionsModel.XCorrelationID = core.StringPtr("testString")
				getHostnameOriginPullCertificateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := authenticatedOriginPullApiService.GetHostnameOriginPullCertificate(getHostnameOriginPullCertificateOptionsModel)
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
	Describe(`DeleteHostnameOriginPullCertificate(deleteHostnameOriginPullCertificateOptions *DeleteHostnameOriginPullCertificateOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		deleteHostnameOriginPullCertificatePath := "/v1/testString/zones/testString/origin_tls_client_auth/hostnames/certificates/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteHostnameOriginPullCertificatePath))
					Expect(req.Method).To(Equal("DELETE"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke DeleteHostnameOriginPullCertificate with error: Operation response processing error`, func() {
				authenticatedOriginPullApiService, serviceErr := authenticatedoriginpullapiv1.NewAuthenticatedOriginPullApiV1(&authenticatedoriginpullapiv1.AuthenticatedOriginPullApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(authenticatedOriginPullApiService).ToNot(BeNil())

				// Construct an instance of the DeleteHostnameOriginPullCertificateOptions model
				deleteHostnameOriginPullCertificateOptionsModel := new(authenticatedoriginpullapiv1.DeleteHostnameOriginPullCertificateOptions)
				deleteHostnameOriginPullCertificateOptionsModel.CertIdentifier = core.StringPtr("testString")
				deleteHostnameOriginPullCertificateOptionsModel.XCorrelationID = core.StringPtr("testString")
				deleteHostnameOriginPullCertificateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := authenticatedOriginPullApiService.DeleteHostnameOriginPullCertificate(deleteHostnameOriginPullCertificateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				authenticatedOriginPullApiService.EnableRetries(0, 0)
				result, response, operationErr = authenticatedOriginPullApiService.DeleteHostnameOriginPullCertificate(deleteHostnameOriginPullCertificateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteHostnameOriginPullCertificate(deleteHostnameOriginPullCertificateOptions *DeleteHostnameOriginPullCertificateOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		deleteHostnameOriginPullCertificatePath := "/v1/testString/zones/testString/origin_tls_client_auth/hostnames/certificates/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteHostnameOriginPullCertificatePath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"result": {"id": "0f405ba2-8c18-49eb-a30b-28b85427780f", "certificate": "-----BEGIN CERTIFICATE-----\n......\n-----END CERTIFICATE-----\n", "issuer": "GlobalSign", "signature": "SHA256WithRSA", "serial_number": "6743787633689793699141714808227354901", "status": "active", "expires_on": "2100-01-01T05:20:00Z", "uploaded_on": "2100-01-01T05:20:00Z"}, "success": true, "errors": ["Errors"], "messages": ["Messages"]}`)
				}))
			})
			It(`Invoke DeleteHostnameOriginPullCertificate successfully with retries`, func() {
				authenticatedOriginPullApiService, serviceErr := authenticatedoriginpullapiv1.NewAuthenticatedOriginPullApiV1(&authenticatedoriginpullapiv1.AuthenticatedOriginPullApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(authenticatedOriginPullApiService).ToNot(BeNil())
				authenticatedOriginPullApiService.EnableRetries(0, 0)

				// Construct an instance of the DeleteHostnameOriginPullCertificateOptions model
				deleteHostnameOriginPullCertificateOptionsModel := new(authenticatedoriginpullapiv1.DeleteHostnameOriginPullCertificateOptions)
				deleteHostnameOriginPullCertificateOptionsModel.CertIdentifier = core.StringPtr("testString")
				deleteHostnameOriginPullCertificateOptionsModel.XCorrelationID = core.StringPtr("testString")
				deleteHostnameOriginPullCertificateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := authenticatedOriginPullApiService.DeleteHostnameOriginPullCertificateWithContext(ctx, deleteHostnameOriginPullCertificateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				authenticatedOriginPullApiService.DisableRetries()
				result, response, operationErr := authenticatedOriginPullApiService.DeleteHostnameOriginPullCertificate(deleteHostnameOriginPullCertificateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = authenticatedOriginPullApiService.DeleteHostnameOriginPullCertificateWithContext(ctx, deleteHostnameOriginPullCertificateOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(deleteHostnameOriginPullCertificatePath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"result": {"id": "0f405ba2-8c18-49eb-a30b-28b85427780f", "certificate": "-----BEGIN CERTIFICATE-----\n......\n-----END CERTIFICATE-----\n", "issuer": "GlobalSign", "signature": "SHA256WithRSA", "serial_number": "6743787633689793699141714808227354901", "status": "active", "expires_on": "2100-01-01T05:20:00Z", "uploaded_on": "2100-01-01T05:20:00Z"}, "success": true, "errors": ["Errors"], "messages": ["Messages"]}`)
				}))
			})
			It(`Invoke DeleteHostnameOriginPullCertificate successfully`, func() {
				authenticatedOriginPullApiService, serviceErr := authenticatedoriginpullapiv1.NewAuthenticatedOriginPullApiV1(&authenticatedoriginpullapiv1.AuthenticatedOriginPullApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(authenticatedOriginPullApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := authenticatedOriginPullApiService.DeleteHostnameOriginPullCertificate(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DeleteHostnameOriginPullCertificateOptions model
				deleteHostnameOriginPullCertificateOptionsModel := new(authenticatedoriginpullapiv1.DeleteHostnameOriginPullCertificateOptions)
				deleteHostnameOriginPullCertificateOptionsModel.CertIdentifier = core.StringPtr("testString")
				deleteHostnameOriginPullCertificateOptionsModel.XCorrelationID = core.StringPtr("testString")
				deleteHostnameOriginPullCertificateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = authenticatedOriginPullApiService.DeleteHostnameOriginPullCertificate(deleteHostnameOriginPullCertificateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke DeleteHostnameOriginPullCertificate with error: Operation validation and request error`, func() {
				authenticatedOriginPullApiService, serviceErr := authenticatedoriginpullapiv1.NewAuthenticatedOriginPullApiV1(&authenticatedoriginpullapiv1.AuthenticatedOriginPullApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(authenticatedOriginPullApiService).ToNot(BeNil())

				// Construct an instance of the DeleteHostnameOriginPullCertificateOptions model
				deleteHostnameOriginPullCertificateOptionsModel := new(authenticatedoriginpullapiv1.DeleteHostnameOriginPullCertificateOptions)
				deleteHostnameOriginPullCertificateOptionsModel.CertIdentifier = core.StringPtr("testString")
				deleteHostnameOriginPullCertificateOptionsModel.XCorrelationID = core.StringPtr("testString")
				deleteHostnameOriginPullCertificateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := authenticatedOriginPullApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := authenticatedOriginPullApiService.DeleteHostnameOriginPullCertificate(deleteHostnameOriginPullCertificateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the DeleteHostnameOriginPullCertificateOptions model with no property values
				deleteHostnameOriginPullCertificateOptionsModelNew := new(authenticatedoriginpullapiv1.DeleteHostnameOriginPullCertificateOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = authenticatedOriginPullApiService.DeleteHostnameOriginPullCertificate(deleteHostnameOriginPullCertificateOptionsModelNew)
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
			It(`Invoke DeleteHostnameOriginPullCertificate successfully`, func() {
				authenticatedOriginPullApiService, serviceErr := authenticatedoriginpullapiv1.NewAuthenticatedOriginPullApiV1(&authenticatedoriginpullapiv1.AuthenticatedOriginPullApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(authenticatedOriginPullApiService).ToNot(BeNil())

				// Construct an instance of the DeleteHostnameOriginPullCertificateOptions model
				deleteHostnameOriginPullCertificateOptionsModel := new(authenticatedoriginpullapiv1.DeleteHostnameOriginPullCertificateOptions)
				deleteHostnameOriginPullCertificateOptionsModel.CertIdentifier = core.StringPtr("testString")
				deleteHostnameOriginPullCertificateOptionsModel.XCorrelationID = core.StringPtr("testString")
				deleteHostnameOriginPullCertificateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := authenticatedOriginPullApiService.DeleteHostnameOriginPullCertificate(deleteHostnameOriginPullCertificateOptionsModel)
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
			authenticatedOriginPullApiService, _ := authenticatedoriginpullapiv1.NewAuthenticatedOriginPullApiV1(&authenticatedoriginpullapiv1.AuthenticatedOriginPullApiV1Options{
				URL:           "http://authenticatedoriginpullapiv1modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
				Crn: core.StringPtr(crn),
				ZoneIdentifier: core.StringPtr(zoneIdentifier),
			})
			It(`Invoke NewDeleteHostnameOriginPullCertificateOptions successfully`, func() {
				// Construct an instance of the DeleteHostnameOriginPullCertificateOptions model
				certIdentifier := "testString"
				deleteHostnameOriginPullCertificateOptionsModel := authenticatedOriginPullApiService.NewDeleteHostnameOriginPullCertificateOptions(certIdentifier)
				deleteHostnameOriginPullCertificateOptionsModel.SetCertIdentifier("testString")
				deleteHostnameOriginPullCertificateOptionsModel.SetXCorrelationID("testString")
				deleteHostnameOriginPullCertificateOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteHostnameOriginPullCertificateOptionsModel).ToNot(BeNil())
				Expect(deleteHostnameOriginPullCertificateOptionsModel.CertIdentifier).To(Equal(core.StringPtr("testString")))
				Expect(deleteHostnameOriginPullCertificateOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(deleteHostnameOriginPullCertificateOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteZoneOriginPullCertificateOptions successfully`, func() {
				// Construct an instance of the DeleteZoneOriginPullCertificateOptions model
				certIdentifier := "testString"
				deleteZoneOriginPullCertificateOptionsModel := authenticatedOriginPullApiService.NewDeleteZoneOriginPullCertificateOptions(certIdentifier)
				deleteZoneOriginPullCertificateOptionsModel.SetCertIdentifier("testString")
				deleteZoneOriginPullCertificateOptionsModel.SetXCorrelationID("testString")
				deleteZoneOriginPullCertificateOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteZoneOriginPullCertificateOptionsModel).ToNot(BeNil())
				Expect(deleteZoneOriginPullCertificateOptionsModel.CertIdentifier).To(Equal(core.StringPtr("testString")))
				Expect(deleteZoneOriginPullCertificateOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(deleteZoneOriginPullCertificateOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetHostnameOriginPullCertificateOptions successfully`, func() {
				// Construct an instance of the GetHostnameOriginPullCertificateOptions model
				certIdentifier := "testString"
				getHostnameOriginPullCertificateOptionsModel := authenticatedOriginPullApiService.NewGetHostnameOriginPullCertificateOptions(certIdentifier)
				getHostnameOriginPullCertificateOptionsModel.SetCertIdentifier("testString")
				getHostnameOriginPullCertificateOptionsModel.SetXCorrelationID("testString")
				getHostnameOriginPullCertificateOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getHostnameOriginPullCertificateOptionsModel).ToNot(BeNil())
				Expect(getHostnameOriginPullCertificateOptionsModel.CertIdentifier).To(Equal(core.StringPtr("testString")))
				Expect(getHostnameOriginPullCertificateOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(getHostnameOriginPullCertificateOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetHostnameOriginPullSettingsOptions successfully`, func() {
				// Construct an instance of the GetHostnameOriginPullSettingsOptions model
				hostname := "testString"
				getHostnameOriginPullSettingsOptionsModel := authenticatedOriginPullApiService.NewGetHostnameOriginPullSettingsOptions(hostname)
				getHostnameOriginPullSettingsOptionsModel.SetHostname("testString")
				getHostnameOriginPullSettingsOptionsModel.SetXCorrelationID("testString")
				getHostnameOriginPullSettingsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getHostnameOriginPullSettingsOptionsModel).ToNot(BeNil())
				Expect(getHostnameOriginPullSettingsOptionsModel.Hostname).To(Equal(core.StringPtr("testString")))
				Expect(getHostnameOriginPullSettingsOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(getHostnameOriginPullSettingsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetZoneOriginPullCertificateOptions successfully`, func() {
				// Construct an instance of the GetZoneOriginPullCertificateOptions model
				certIdentifier := "testString"
				getZoneOriginPullCertificateOptionsModel := authenticatedOriginPullApiService.NewGetZoneOriginPullCertificateOptions(certIdentifier)
				getZoneOriginPullCertificateOptionsModel.SetCertIdentifier("testString")
				getZoneOriginPullCertificateOptionsModel.SetXCorrelationID("testString")
				getZoneOriginPullCertificateOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getZoneOriginPullCertificateOptionsModel).ToNot(BeNil())
				Expect(getZoneOriginPullCertificateOptionsModel.CertIdentifier).To(Equal(core.StringPtr("testString")))
				Expect(getZoneOriginPullCertificateOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(getZoneOriginPullCertificateOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetZoneOriginPullSettingsOptions successfully`, func() {
				// Construct an instance of the GetZoneOriginPullSettingsOptions model
				getZoneOriginPullSettingsOptionsModel := authenticatedOriginPullApiService.NewGetZoneOriginPullSettingsOptions()
				getZoneOriginPullSettingsOptionsModel.SetXCorrelationID("testString")
				getZoneOriginPullSettingsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getZoneOriginPullSettingsOptionsModel).ToNot(BeNil())
				Expect(getZoneOriginPullSettingsOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(getZoneOriginPullSettingsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListZoneOriginPullCertificatesOptions successfully`, func() {
				// Construct an instance of the ListZoneOriginPullCertificatesOptions model
				listZoneOriginPullCertificatesOptionsModel := authenticatedOriginPullApiService.NewListZoneOriginPullCertificatesOptions()
				listZoneOriginPullCertificatesOptionsModel.SetXCorrelationID("testString")
				listZoneOriginPullCertificatesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listZoneOriginPullCertificatesOptionsModel).ToNot(BeNil())
				Expect(listZoneOriginPullCertificatesOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(listZoneOriginPullCertificatesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewSetHostnameOriginPullSettingsOptions successfully`, func() {
				// Construct an instance of the HostnameOriginPullSettings model
				hostnameOriginPullSettingsModel := new(authenticatedoriginpullapiv1.HostnameOriginPullSettings)
				Expect(hostnameOriginPullSettingsModel).ToNot(BeNil())
				hostnameOriginPullSettingsModel.Hostname = core.StringPtr("app.example.com")
				hostnameOriginPullSettingsModel.CertID = core.StringPtr("2458ce5a-0c35-4c7f-82c7-8e9487d3ff60")
				hostnameOriginPullSettingsModel.Enabled = core.BoolPtr(true)
				Expect(hostnameOriginPullSettingsModel.Hostname).To(Equal(core.StringPtr("app.example.com")))
				Expect(hostnameOriginPullSettingsModel.CertID).To(Equal(core.StringPtr("2458ce5a-0c35-4c7f-82c7-8e9487d3ff60")))
				Expect(hostnameOriginPullSettingsModel.Enabled).To(Equal(core.BoolPtr(true)))

				// Construct an instance of the SetHostnameOriginPullSettingsOptions model
				setHostnameOriginPullSettingsOptionsModel := authenticatedOriginPullApiService.NewSetHostnameOriginPullSettingsOptions()
				setHostnameOriginPullSettingsOptionsModel.SetConfig([]authenticatedoriginpullapiv1.HostnameOriginPullSettings{*hostnameOriginPullSettingsModel})
				setHostnameOriginPullSettingsOptionsModel.SetXCorrelationID("testString")
				setHostnameOriginPullSettingsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(setHostnameOriginPullSettingsOptionsModel).ToNot(BeNil())
				Expect(setHostnameOriginPullSettingsOptionsModel.Config).To(Equal([]authenticatedoriginpullapiv1.HostnameOriginPullSettings{*hostnameOriginPullSettingsModel}))
				Expect(setHostnameOriginPullSettingsOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(setHostnameOriginPullSettingsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewSetZoneOriginPullSettingsOptions successfully`, func() {
				// Construct an instance of the SetZoneOriginPullSettingsOptions model
				setZoneOriginPullSettingsOptionsModel := authenticatedOriginPullApiService.NewSetZoneOriginPullSettingsOptions()
				setZoneOriginPullSettingsOptionsModel.SetEnabled(true)
				setZoneOriginPullSettingsOptionsModel.SetXCorrelationID("testString")
				setZoneOriginPullSettingsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(setZoneOriginPullSettingsOptionsModel).ToNot(BeNil())
				Expect(setZoneOriginPullSettingsOptionsModel.Enabled).To(Equal(core.BoolPtr(true)))
				Expect(setZoneOriginPullSettingsOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(setZoneOriginPullSettingsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUploadHostnameOriginPullCertificateOptions successfully`, func() {
				// Construct an instance of the UploadHostnameOriginPullCertificateOptions model
				uploadHostnameOriginPullCertificateOptionsModel := authenticatedOriginPullApiService.NewUploadHostnameOriginPullCertificateOptions()
				uploadHostnameOriginPullCertificateOptionsModel.SetCertificate("-----BEGIN CERTIFICATE-----\n......\n-----END CERTIFICATE-----\n")
				uploadHostnameOriginPullCertificateOptionsModel.SetPrivateKey("-----BEGIN RSA PRIVATE KEY-----\n......\n-----END RSA PRIVATE KEY-----\n")
				uploadHostnameOriginPullCertificateOptionsModel.SetXCorrelationID("testString")
				uploadHostnameOriginPullCertificateOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(uploadHostnameOriginPullCertificateOptionsModel).ToNot(BeNil())
				Expect(uploadHostnameOriginPullCertificateOptionsModel.Certificate).To(Equal(core.StringPtr("-----BEGIN CERTIFICATE-----\n......\n-----END CERTIFICATE-----\n")))
				Expect(uploadHostnameOriginPullCertificateOptionsModel.PrivateKey).To(Equal(core.StringPtr("-----BEGIN RSA PRIVATE KEY-----\n......\n-----END RSA PRIVATE KEY-----\n")))
				Expect(uploadHostnameOriginPullCertificateOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(uploadHostnameOriginPullCertificateOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUploadZoneOriginPullCertificateOptions successfully`, func() {
				// Construct an instance of the UploadZoneOriginPullCertificateOptions model
				uploadZoneOriginPullCertificateOptionsModel := authenticatedOriginPullApiService.NewUploadZoneOriginPullCertificateOptions()
				uploadZoneOriginPullCertificateOptionsModel.SetCertificate("-----BEGIN CERTIFICATE-----\n......\n-----END CERTIFICATE-----\n")
				uploadZoneOriginPullCertificateOptionsModel.SetPrivateKey("-----BEGIN RSA PRIVATE KEY-----\n......\n-----END RSA PRIVATE KEY-----\n")
				uploadZoneOriginPullCertificateOptionsModel.SetXCorrelationID("testString")
				uploadZoneOriginPullCertificateOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(uploadZoneOriginPullCertificateOptionsModel).ToNot(BeNil())
				Expect(uploadZoneOriginPullCertificateOptionsModel.Certificate).To(Equal(core.StringPtr("-----BEGIN CERTIFICATE-----\n......\n-----END CERTIFICATE-----\n")))
				Expect(uploadZoneOriginPullCertificateOptionsModel.PrivateKey).To(Equal(core.StringPtr("-----BEGIN RSA PRIVATE KEY-----\n......\n-----END RSA PRIVATE KEY-----\n")))
				Expect(uploadZoneOriginPullCertificateOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(uploadZoneOriginPullCertificateOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewHostnameOriginPullSettings successfully`, func() {
				hostname := "app.example.com"
				certID := "2458ce5a-0c35-4c7f-82c7-8e9487d3ff60"
				enabled := true
				_model, err := authenticatedOriginPullApiService.NewHostnameOriginPullSettings(hostname, certID, enabled)
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
