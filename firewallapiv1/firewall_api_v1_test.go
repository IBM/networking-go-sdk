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

package firewallapiv1_test

import (
	"bytes"
	"context"
	"fmt"
	"github.com/IBM/go-sdk-core/v4/core"
	"github.com/IBM/networking-go-sdk/firewallapiv1"
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

var _ = Describe(`FirewallApiV1`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		It(`Instantiate service client`, func() {
			firewallApiService, serviceErr := firewallapiv1.NewFirewallApiV1(&firewallapiv1.FirewallApiV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
				Crn: core.StringPtr(crn),
				ZoneIdentifier: core.StringPtr(zoneIdentifier),
			})
			Expect(firewallApiService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			firewallApiService, serviceErr := firewallapiv1.NewFirewallApiV1(&firewallapiv1.FirewallApiV1Options{
				URL: "{BAD_URL_STRING",
				Crn: core.StringPtr(crn),
				ZoneIdentifier: core.StringPtr(zoneIdentifier),
			})
			Expect(firewallApiService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			firewallApiService, serviceErr := firewallapiv1.NewFirewallApiV1(&firewallapiv1.FirewallApiV1Options{
				URL: "https://firewallapiv1/api",
				Crn: core.StringPtr(crn),
				ZoneIdentifier: core.StringPtr(zoneIdentifier),
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(firewallApiService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Validation Error`, func() {
			firewallApiService, serviceErr := firewallapiv1.NewFirewallApiV1(&firewallapiv1.FirewallApiV1Options{})
			Expect(firewallApiService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"FIREWALL_API_URL": "https://firewallapiv1/api",
				"FIREWALL_API_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				firewallApiService, serviceErr := firewallapiv1.NewFirewallApiV1UsingExternalConfig(&firewallapiv1.FirewallApiV1Options{
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(firewallApiService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := firewallApiService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != firewallApiService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(firewallApiService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(firewallApiService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				firewallApiService, serviceErr := firewallapiv1.NewFirewallApiV1UsingExternalConfig(&firewallapiv1.FirewallApiV1Options{
					URL: "https://testService/api",
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(firewallApiService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(firewallApiService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := firewallApiService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != firewallApiService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(firewallApiService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(firewallApiService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				firewallApiService, serviceErr := firewallapiv1.NewFirewallApiV1UsingExternalConfig(&firewallapiv1.FirewallApiV1Options{
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				err := firewallApiService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(firewallApiService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(firewallApiService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := firewallApiService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != firewallApiService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(firewallApiService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(firewallApiService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"FIREWALL_API_URL": "https://firewallapiv1/api",
				"FIREWALL_API_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			firewallApiService, serviceErr := firewallapiv1.NewFirewallApiV1UsingExternalConfig(&firewallapiv1.FirewallApiV1Options{
				Crn: core.StringPtr(crn),
				ZoneIdentifier: core.StringPtr(zoneIdentifier),
			})

			It(`Instantiate service client with error`, func() {
				Expect(firewallApiService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"FIREWALL_API_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			firewallApiService, serviceErr := firewallapiv1.NewFirewallApiV1UsingExternalConfig(&firewallapiv1.FirewallApiV1Options{
				URL: "{BAD_URL_STRING",
				Crn: core.StringPtr(crn),
				ZoneIdentifier: core.StringPtr(zoneIdentifier),
			})

			It(`Instantiate service client with error`, func() {
				Expect(firewallApiService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = firewallapiv1.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`GetSecurityLevelSetting(getSecurityLevelSettingOptions *GetSecurityLevelSettingOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getSecurityLevelSettingPath := "/v1/testString/zones/testString/settings/security_level"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getSecurityLevelSettingPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetSecurityLevelSetting with error: Operation response processing error`, func() {
				firewallApiService, serviceErr := firewallapiv1.NewFirewallApiV1(&firewallapiv1.FirewallApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(firewallApiService).ToNot(BeNil())

				// Construct an instance of the GetSecurityLevelSettingOptions model
				getSecurityLevelSettingOptionsModel := new(firewallapiv1.GetSecurityLevelSettingOptions)
				getSecurityLevelSettingOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := firewallApiService.GetSecurityLevelSetting(getSecurityLevelSettingOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				firewallApiService.EnableRetries(0, 0)
				result, response, operationErr = firewallApiService.GetSecurityLevelSetting(getSecurityLevelSettingOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetSecurityLevelSetting(getSecurityLevelSettingOptions *GetSecurityLevelSettingOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getSecurityLevelSettingPath := "/v1/testString/zones/testString/settings/security_level"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getSecurityLevelSettingPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"result": {"id": "security_level", "value": "medium", "editable": true, "modified_on": "2014-01-01T05:20:00.12345Z"}, "result_info": {"page": 1, "per_page": 2, "count": 1, "total_count": 200}, "success": true, "errors": [["Errors"]], "messages": [{"status": "OK"}]}`)
				}))
			})
			It(`Invoke GetSecurityLevelSetting successfully`, func() {
				firewallApiService, serviceErr := firewallapiv1.NewFirewallApiV1(&firewallapiv1.FirewallApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(firewallApiService).ToNot(BeNil())
				firewallApiService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := firewallApiService.GetSecurityLevelSetting(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetSecurityLevelSettingOptions model
				getSecurityLevelSettingOptionsModel := new(firewallapiv1.GetSecurityLevelSettingOptions)
				getSecurityLevelSettingOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = firewallApiService.GetSecurityLevelSetting(getSecurityLevelSettingOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = firewallApiService.GetSecurityLevelSettingWithContext(ctx, getSecurityLevelSettingOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				firewallApiService.DisableRetries()
				result, response, operationErr = firewallApiService.GetSecurityLevelSetting(getSecurityLevelSettingOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = firewallApiService.GetSecurityLevelSettingWithContext(ctx, getSecurityLevelSettingOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke GetSecurityLevelSetting with error: Operation request error`, func() {
				firewallApiService, serviceErr := firewallapiv1.NewFirewallApiV1(&firewallapiv1.FirewallApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(firewallApiService).ToNot(BeNil())

				// Construct an instance of the GetSecurityLevelSettingOptions model
				getSecurityLevelSettingOptionsModel := new(firewallapiv1.GetSecurityLevelSettingOptions)
				getSecurityLevelSettingOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := firewallApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := firewallApiService.GetSecurityLevelSetting(getSecurityLevelSettingOptionsModel)
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
	Describe(`SetSecurityLevelSetting(setSecurityLevelSettingOptions *SetSecurityLevelSettingOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		setSecurityLevelSettingPath := "/v1/testString/zones/testString/settings/security_level"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(setSecurityLevelSettingPath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke SetSecurityLevelSetting with error: Operation response processing error`, func() {
				firewallApiService, serviceErr := firewallapiv1.NewFirewallApiV1(&firewallapiv1.FirewallApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(firewallApiService).ToNot(BeNil())

				// Construct an instance of the SetSecurityLevelSettingOptions model
				setSecurityLevelSettingOptionsModel := new(firewallapiv1.SetSecurityLevelSettingOptions)
				setSecurityLevelSettingOptionsModel.Value = core.StringPtr("under_attack")
				setSecurityLevelSettingOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := firewallApiService.SetSecurityLevelSetting(setSecurityLevelSettingOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				firewallApiService.EnableRetries(0, 0)
				result, response, operationErr = firewallApiService.SetSecurityLevelSetting(setSecurityLevelSettingOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`SetSecurityLevelSetting(setSecurityLevelSettingOptions *SetSecurityLevelSettingOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		setSecurityLevelSettingPath := "/v1/testString/zones/testString/settings/security_level"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(setSecurityLevelSettingPath))
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
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"result": {"id": "security_level", "value": "medium", "editable": true, "modified_on": "2014-01-01T05:20:00.12345Z"}, "result_info": {"page": 1, "per_page": 2, "count": 1, "total_count": 200}, "success": true, "errors": [["Errors"]], "messages": [{"status": "OK"}]}`)
				}))
			})
			It(`Invoke SetSecurityLevelSetting successfully`, func() {
				firewallApiService, serviceErr := firewallapiv1.NewFirewallApiV1(&firewallapiv1.FirewallApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(firewallApiService).ToNot(BeNil())
				firewallApiService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := firewallApiService.SetSecurityLevelSetting(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the SetSecurityLevelSettingOptions model
				setSecurityLevelSettingOptionsModel := new(firewallapiv1.SetSecurityLevelSettingOptions)
				setSecurityLevelSettingOptionsModel.Value = core.StringPtr("under_attack")
				setSecurityLevelSettingOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = firewallApiService.SetSecurityLevelSetting(setSecurityLevelSettingOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = firewallApiService.SetSecurityLevelSettingWithContext(ctx, setSecurityLevelSettingOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				firewallApiService.DisableRetries()
				result, response, operationErr = firewallApiService.SetSecurityLevelSetting(setSecurityLevelSettingOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = firewallApiService.SetSecurityLevelSettingWithContext(ctx, setSecurityLevelSettingOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke SetSecurityLevelSetting with error: Operation request error`, func() {
				firewallApiService, serviceErr := firewallapiv1.NewFirewallApiV1(&firewallapiv1.FirewallApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(firewallApiService).ToNot(BeNil())

				// Construct an instance of the SetSecurityLevelSettingOptions model
				setSecurityLevelSettingOptionsModel := new(firewallapiv1.SetSecurityLevelSettingOptions)
				setSecurityLevelSettingOptionsModel.Value = core.StringPtr("under_attack")
				setSecurityLevelSettingOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := firewallApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := firewallApiService.SetSecurityLevelSetting(setSecurityLevelSettingOptionsModel)
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
	Describe(`Model constructor tests`, func() {
		Context(`Using a service client instance`, func() {
			crn := "testString"
			zoneIdentifier := "testString"
			firewallApiService, _ := firewallapiv1.NewFirewallApiV1(&firewallapiv1.FirewallApiV1Options{
				URL:           "http://firewallapiv1modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
				Crn: core.StringPtr(crn),
				ZoneIdentifier: core.StringPtr(zoneIdentifier),
			})
			It(`Invoke NewGetSecurityLevelSettingOptions successfully`, func() {
				// Construct an instance of the GetSecurityLevelSettingOptions model
				getSecurityLevelSettingOptionsModel := firewallApiService.NewGetSecurityLevelSettingOptions()
				getSecurityLevelSettingOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getSecurityLevelSettingOptionsModel).ToNot(BeNil())
				Expect(getSecurityLevelSettingOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewSetSecurityLevelSettingOptions successfully`, func() {
				// Construct an instance of the SetSecurityLevelSettingOptions model
				setSecurityLevelSettingOptionsModel := firewallApiService.NewSetSecurityLevelSettingOptions()
				setSecurityLevelSettingOptionsModel.SetValue("under_attack")
				setSecurityLevelSettingOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(setSecurityLevelSettingOptionsModel).ToNot(BeNil())
				Expect(setSecurityLevelSettingOptionsModel.Value).To(Equal(core.StringPtr("under_attack")))
				Expect(setSecurityLevelSettingOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
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
