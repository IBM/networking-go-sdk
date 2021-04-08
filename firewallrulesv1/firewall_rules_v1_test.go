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

package firewallrulesv1_test

import (
	"bytes"
	"context"
	"fmt"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/IBM/networking-go-sdk/firewallrulesv1"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"time"
)

var _ = Describe(`FirewallRulesV1`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			firewallRulesService, serviceErr := firewallrulesv1.NewFirewallRulesV1(&firewallrulesv1.FirewallRulesV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(firewallRulesService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			firewallRulesService, serviceErr := firewallrulesv1.NewFirewallRulesV1(&firewallrulesv1.FirewallRulesV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(firewallRulesService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			firewallRulesService, serviceErr := firewallrulesv1.NewFirewallRulesV1(&firewallrulesv1.FirewallRulesV1Options{
				URL: "https://firewallrulesv1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(firewallRulesService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"FIREWALL_RULES_URL": "https://firewallrulesv1/api",
				"FIREWALL_RULES_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				firewallRulesService, serviceErr := firewallrulesv1.NewFirewallRulesV1UsingExternalConfig(&firewallrulesv1.FirewallRulesV1Options{
				})
				Expect(firewallRulesService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := firewallRulesService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != firewallRulesService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(firewallRulesService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(firewallRulesService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				firewallRulesService, serviceErr := firewallrulesv1.NewFirewallRulesV1UsingExternalConfig(&firewallrulesv1.FirewallRulesV1Options{
					URL: "https://testService/api",
				})
				Expect(firewallRulesService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(firewallRulesService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := firewallRulesService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != firewallRulesService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(firewallRulesService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(firewallRulesService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				firewallRulesService, serviceErr := firewallrulesv1.NewFirewallRulesV1UsingExternalConfig(&firewallrulesv1.FirewallRulesV1Options{
				})
				err := firewallRulesService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(firewallRulesService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(firewallRulesService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := firewallRulesService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != firewallRulesService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(firewallRulesService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(firewallRulesService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"FIREWALL_RULES_URL": "https://firewallrulesv1/api",
				"FIREWALL_RULES_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			firewallRulesService, serviceErr := firewallrulesv1.NewFirewallRulesV1UsingExternalConfig(&firewallrulesv1.FirewallRulesV1Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(firewallRulesService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"FIREWALL_RULES_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			firewallRulesService, serviceErr := firewallrulesv1.NewFirewallRulesV1UsingExternalConfig(&firewallrulesv1.FirewallRulesV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(firewallRulesService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = firewallrulesv1.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})

	Describe(`ListAllFirewallRules(listAllFirewallRulesOptions *ListAllFirewallRulesOptions)`, func() {
		listAllFirewallRulesPath := "/v1/testString/zones/testString/firewall/rules"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listAllFirewallRulesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Auth-User-Token"]).ToNot(BeNil())
					Expect(req.Header["X-Auth-User-Token"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Accept"]).ToNot(BeNil())
					Expect(req.Header["Accept"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "*/*")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `This is a mock binary response.`)
				}))
			})
			It(`Invoke ListAllFirewallRules successfully with retries`, func() {
				firewallRulesService, serviceErr := firewallrulesv1.NewFirewallRulesV1(&firewallrulesv1.FirewallRulesV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(firewallRulesService).ToNot(BeNil())
				firewallRulesService.EnableRetries(0, 0)

				// Construct an instance of the ListAllFirewallRulesOptions model
				listAllFirewallRulesOptionsModel := new(firewallrulesv1.ListAllFirewallRulesOptions)
				listAllFirewallRulesOptionsModel.XAuthUserToken = core.StringPtr("testString")
				listAllFirewallRulesOptionsModel.Crn = core.StringPtr("testString")
				listAllFirewallRulesOptionsModel.ZoneIdentifier = core.StringPtr("testString")
				listAllFirewallRulesOptionsModel.Accept = core.StringPtr("testString")
				listAllFirewallRulesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := firewallRulesService.ListAllFirewallRulesWithContext(ctx, listAllFirewallRulesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				firewallRulesService.DisableRetries()
				result, response, operationErr := firewallRulesService.ListAllFirewallRules(listAllFirewallRulesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = firewallRulesService.ListAllFirewallRulesWithContext(ctx, listAllFirewallRulesOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listAllFirewallRulesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Auth-User-Token"]).ToNot(BeNil())
					Expect(req.Header["X-Auth-User-Token"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Accept"]).ToNot(BeNil())
					Expect(req.Header["Accept"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "*/*")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `This is a mock binary response.`)
				}))
			})
			It(`Invoke ListAllFirewallRules successfully`, func() {
				firewallRulesService, serviceErr := firewallrulesv1.NewFirewallRulesV1(&firewallrulesv1.FirewallRulesV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(firewallRulesService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := firewallRulesService.ListAllFirewallRules(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListAllFirewallRulesOptions model
				listAllFirewallRulesOptionsModel := new(firewallrulesv1.ListAllFirewallRulesOptions)
				listAllFirewallRulesOptionsModel.XAuthUserToken = core.StringPtr("testString")
				listAllFirewallRulesOptionsModel.Crn = core.StringPtr("testString")
				listAllFirewallRulesOptionsModel.ZoneIdentifier = core.StringPtr("testString")
				listAllFirewallRulesOptionsModel.Accept = core.StringPtr("testString")
				listAllFirewallRulesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = firewallRulesService.ListAllFirewallRules(listAllFirewallRulesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListAllFirewallRules with error: Operation validation and request error`, func() {
				firewallRulesService, serviceErr := firewallrulesv1.NewFirewallRulesV1(&firewallrulesv1.FirewallRulesV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(firewallRulesService).ToNot(BeNil())

				// Construct an instance of the ListAllFirewallRulesOptions model
				listAllFirewallRulesOptionsModel := new(firewallrulesv1.ListAllFirewallRulesOptions)
				listAllFirewallRulesOptionsModel.XAuthUserToken = core.StringPtr("testString")
				listAllFirewallRulesOptionsModel.Crn = core.StringPtr("testString")
				listAllFirewallRulesOptionsModel.ZoneIdentifier = core.StringPtr("testString")
				listAllFirewallRulesOptionsModel.Accept = core.StringPtr("testString")
				listAllFirewallRulesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := firewallRulesService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := firewallRulesService.ListAllFirewallRules(listAllFirewallRulesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListAllFirewallRulesOptions model with no property values
				listAllFirewallRulesOptionsModelNew := new(firewallrulesv1.ListAllFirewallRulesOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = firewallRulesService.ListAllFirewallRules(listAllFirewallRulesOptionsModelNew)
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
		It(`Instantiate service client`, func() {
			firewallRulesService, serviceErr := firewallrulesv1.NewFirewallRulesV1(&firewallrulesv1.FirewallRulesV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(firewallRulesService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			firewallRulesService, serviceErr := firewallrulesv1.NewFirewallRulesV1(&firewallrulesv1.FirewallRulesV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(firewallRulesService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			firewallRulesService, serviceErr := firewallrulesv1.NewFirewallRulesV1(&firewallrulesv1.FirewallRulesV1Options{
				URL: "https://firewallrulesv1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(firewallRulesService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"FIREWALL_RULES_URL": "https://firewallrulesv1/api",
				"FIREWALL_RULES_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				firewallRulesService, serviceErr := firewallrulesv1.NewFirewallRulesV1UsingExternalConfig(&firewallrulesv1.FirewallRulesV1Options{
				})
				Expect(firewallRulesService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := firewallRulesService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != firewallRulesService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(firewallRulesService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(firewallRulesService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				firewallRulesService, serviceErr := firewallrulesv1.NewFirewallRulesV1UsingExternalConfig(&firewallrulesv1.FirewallRulesV1Options{
					URL: "https://testService/api",
				})
				Expect(firewallRulesService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(firewallRulesService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := firewallRulesService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != firewallRulesService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(firewallRulesService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(firewallRulesService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				firewallRulesService, serviceErr := firewallrulesv1.NewFirewallRulesV1UsingExternalConfig(&firewallrulesv1.FirewallRulesV1Options{
				})
				err := firewallRulesService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(firewallRulesService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(firewallRulesService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := firewallRulesService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != firewallRulesService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(firewallRulesService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(firewallRulesService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"FIREWALL_RULES_URL": "https://firewallrulesv1/api",
				"FIREWALL_RULES_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			firewallRulesService, serviceErr := firewallrulesv1.NewFirewallRulesV1UsingExternalConfig(&firewallrulesv1.FirewallRulesV1Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(firewallRulesService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"FIREWALL_RULES_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			firewallRulesService, serviceErr := firewallrulesv1.NewFirewallRulesV1UsingExternalConfig(&firewallrulesv1.FirewallRulesV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(firewallRulesService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = firewallrulesv1.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})

	Describe(`CreateFirewallRules(createFirewallRulesOptions *CreateFirewallRulesOptions)`, func() {
		createFirewallRulesPath := "/v1/testString/zones/testString/firewall/rules"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createFirewallRulesPath))
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

					Expect(req.Header["X-Auth-User-Token"]).ToNot(BeNil())
					Expect(req.Header["X-Auth-User-Token"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Accept"]).ToNot(BeNil())
					Expect(req.Header["Accept"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "*/*")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `This is a mock binary response.`)
				}))
			})
			It(`Invoke CreateFirewallRules successfully with retries`, func() {
				firewallRulesService, serviceErr := firewallrulesv1.NewFirewallRulesV1(&firewallrulesv1.FirewallRulesV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(firewallRulesService).ToNot(BeNil())
				firewallRulesService.EnableRetries(0, 0)

				// Construct an instance of the FirewallRuleInputWithFilterIdFilter model
				firewallRuleInputWithFilterIdFilterModel := new(firewallrulesv1.FirewallRuleInputWithFilterIdFilter)
				firewallRuleInputWithFilterIdFilterModel.ID = core.StringPtr("6f58318e7fa2477a23112e8118c66f61")

				// Construct an instance of the FirewallRuleInputWithFilterID model
				firewallRuleInputWithFilterIdModel := new(firewallrulesv1.FirewallRuleInputWithFilterID)
				firewallRuleInputWithFilterIdModel.Filter = firewallRuleInputWithFilterIdFilterModel
				firewallRuleInputWithFilterIdModel.Action = core.StringPtr("js_challenge")
				firewallRuleInputWithFilterIdModel.Description = core.StringPtr("JS challenge site")

				// Construct an instance of the CreateFirewallRulesOptions model
				createFirewallRulesOptionsModel := new(firewallrulesv1.CreateFirewallRulesOptions)
				createFirewallRulesOptionsModel.XAuthUserToken = core.StringPtr("testString")
				createFirewallRulesOptionsModel.Crn = core.StringPtr("testString")
				createFirewallRulesOptionsModel.ZoneIdentifier = core.StringPtr("testString")
				createFirewallRulesOptionsModel.Accept = core.StringPtr("testString")
				createFirewallRulesOptionsModel.FirewallRuleInputWithFilterID = []firewallrulesv1.FirewallRuleInputWithFilterID{*firewallRuleInputWithFilterIdModel}
				createFirewallRulesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := firewallRulesService.CreateFirewallRulesWithContext(ctx, createFirewallRulesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				firewallRulesService.DisableRetries()
				result, response, operationErr := firewallRulesService.CreateFirewallRules(createFirewallRulesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = firewallRulesService.CreateFirewallRulesWithContext(ctx, createFirewallRulesOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createFirewallRulesPath))
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

					Expect(req.Header["X-Auth-User-Token"]).ToNot(BeNil())
					Expect(req.Header["X-Auth-User-Token"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Accept"]).ToNot(BeNil())
					Expect(req.Header["Accept"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "*/*")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `This is a mock binary response.`)
				}))
			})
			It(`Invoke CreateFirewallRules successfully`, func() {
				firewallRulesService, serviceErr := firewallrulesv1.NewFirewallRulesV1(&firewallrulesv1.FirewallRulesV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(firewallRulesService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := firewallRulesService.CreateFirewallRules(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the FirewallRuleInputWithFilterIdFilter model
				firewallRuleInputWithFilterIdFilterModel := new(firewallrulesv1.FirewallRuleInputWithFilterIdFilter)
				firewallRuleInputWithFilterIdFilterModel.ID = core.StringPtr("6f58318e7fa2477a23112e8118c66f61")

				// Construct an instance of the FirewallRuleInputWithFilterID model
				firewallRuleInputWithFilterIdModel := new(firewallrulesv1.FirewallRuleInputWithFilterID)
				firewallRuleInputWithFilterIdModel.Filter = firewallRuleInputWithFilterIdFilterModel
				firewallRuleInputWithFilterIdModel.Action = core.StringPtr("js_challenge")
				firewallRuleInputWithFilterIdModel.Description = core.StringPtr("JS challenge site")

				// Construct an instance of the CreateFirewallRulesOptions model
				createFirewallRulesOptionsModel := new(firewallrulesv1.CreateFirewallRulesOptions)
				createFirewallRulesOptionsModel.XAuthUserToken = core.StringPtr("testString")
				createFirewallRulesOptionsModel.Crn = core.StringPtr("testString")
				createFirewallRulesOptionsModel.ZoneIdentifier = core.StringPtr("testString")
				createFirewallRulesOptionsModel.Accept = core.StringPtr("testString")
				createFirewallRulesOptionsModel.FirewallRuleInputWithFilterID = []firewallrulesv1.FirewallRuleInputWithFilterID{*firewallRuleInputWithFilterIdModel}
				createFirewallRulesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = firewallRulesService.CreateFirewallRules(createFirewallRulesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateFirewallRules with error: Operation validation and request error`, func() {
				firewallRulesService, serviceErr := firewallrulesv1.NewFirewallRulesV1(&firewallrulesv1.FirewallRulesV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(firewallRulesService).ToNot(BeNil())

				// Construct an instance of the FirewallRuleInputWithFilterIdFilter model
				firewallRuleInputWithFilterIdFilterModel := new(firewallrulesv1.FirewallRuleInputWithFilterIdFilter)
				firewallRuleInputWithFilterIdFilterModel.ID = core.StringPtr("6f58318e7fa2477a23112e8118c66f61")

				// Construct an instance of the FirewallRuleInputWithFilterID model
				firewallRuleInputWithFilterIdModel := new(firewallrulesv1.FirewallRuleInputWithFilterID)
				firewallRuleInputWithFilterIdModel.Filter = firewallRuleInputWithFilterIdFilterModel
				firewallRuleInputWithFilterIdModel.Action = core.StringPtr("js_challenge")
				firewallRuleInputWithFilterIdModel.Description = core.StringPtr("JS challenge site")

				// Construct an instance of the CreateFirewallRulesOptions model
				createFirewallRulesOptionsModel := new(firewallrulesv1.CreateFirewallRulesOptions)
				createFirewallRulesOptionsModel.XAuthUserToken = core.StringPtr("testString")
				createFirewallRulesOptionsModel.Crn = core.StringPtr("testString")
				createFirewallRulesOptionsModel.ZoneIdentifier = core.StringPtr("testString")
				createFirewallRulesOptionsModel.Accept = core.StringPtr("testString")
				createFirewallRulesOptionsModel.FirewallRuleInputWithFilterID = []firewallrulesv1.FirewallRuleInputWithFilterID{*firewallRuleInputWithFilterIdModel}
				createFirewallRulesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := firewallRulesService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := firewallRulesService.CreateFirewallRules(createFirewallRulesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateFirewallRulesOptions model with no property values
				createFirewallRulesOptionsModelNew := new(firewallrulesv1.CreateFirewallRulesOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = firewallRulesService.CreateFirewallRules(createFirewallRulesOptionsModelNew)
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
		It(`Instantiate service client`, func() {
			firewallRulesService, serviceErr := firewallrulesv1.NewFirewallRulesV1(&firewallrulesv1.FirewallRulesV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(firewallRulesService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			firewallRulesService, serviceErr := firewallrulesv1.NewFirewallRulesV1(&firewallrulesv1.FirewallRulesV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(firewallRulesService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			firewallRulesService, serviceErr := firewallrulesv1.NewFirewallRulesV1(&firewallrulesv1.FirewallRulesV1Options{
				URL: "https://firewallrulesv1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(firewallRulesService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"FIREWALL_RULES_URL": "https://firewallrulesv1/api",
				"FIREWALL_RULES_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				firewallRulesService, serviceErr := firewallrulesv1.NewFirewallRulesV1UsingExternalConfig(&firewallrulesv1.FirewallRulesV1Options{
				})
				Expect(firewallRulesService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := firewallRulesService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != firewallRulesService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(firewallRulesService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(firewallRulesService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				firewallRulesService, serviceErr := firewallrulesv1.NewFirewallRulesV1UsingExternalConfig(&firewallrulesv1.FirewallRulesV1Options{
					URL: "https://testService/api",
				})
				Expect(firewallRulesService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(firewallRulesService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := firewallRulesService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != firewallRulesService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(firewallRulesService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(firewallRulesService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				firewallRulesService, serviceErr := firewallrulesv1.NewFirewallRulesV1UsingExternalConfig(&firewallrulesv1.FirewallRulesV1Options{
				})
				err := firewallRulesService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(firewallRulesService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(firewallRulesService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := firewallRulesService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != firewallRulesService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(firewallRulesService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(firewallRulesService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"FIREWALL_RULES_URL": "https://firewallrulesv1/api",
				"FIREWALL_RULES_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			firewallRulesService, serviceErr := firewallrulesv1.NewFirewallRulesV1UsingExternalConfig(&firewallrulesv1.FirewallRulesV1Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(firewallRulesService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"FIREWALL_RULES_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			firewallRulesService, serviceErr := firewallrulesv1.NewFirewallRulesV1UsingExternalConfig(&firewallrulesv1.FirewallRulesV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(firewallRulesService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = firewallrulesv1.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})

	Describe(`UpdateFirewllRules(updateFirewllRulesOptions *UpdateFirewllRulesOptions)`, func() {
		updateFirewllRulesPath := "/v1/testString/zones/testString/firewall/rules"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateFirewllRulesPath))
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

					Expect(req.Header["X-Auth-User-Token"]).ToNot(BeNil())
					Expect(req.Header["X-Auth-User-Token"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Accept"]).ToNot(BeNil())
					Expect(req.Header["Accept"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "*/*")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `This is a mock binary response.`)
				}))
			})
			It(`Invoke UpdateFirewllRules successfully with retries`, func() {
				firewallRulesService, serviceErr := firewallrulesv1.NewFirewallRulesV1(&firewallrulesv1.FirewallRulesV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(firewallRulesService).ToNot(BeNil())
				firewallRulesService.EnableRetries(0, 0)

				// Construct an instance of the FirewallRulesUpdateInputItemFilter model
				firewallRulesUpdateInputItemFilterModel := new(firewallrulesv1.FirewallRulesUpdateInputItemFilter)
				firewallRulesUpdateInputItemFilterModel.ID = core.StringPtr("6f58318e7fa2477a23112e8118c66f61")

				// Construct an instance of the FirewallRulesUpdateInputItem model
				firewallRulesUpdateInputItemModel := new(firewallrulesv1.FirewallRulesUpdateInputItem)
				firewallRulesUpdateInputItemModel.ID = core.StringPtr("52161eb6af4241bb9d4b32394be72fdf")
				firewallRulesUpdateInputItemModel.Action = core.StringPtr("js_challenge")
				firewallRulesUpdateInputItemModel.Paused = core.BoolPtr(false)
				firewallRulesUpdateInputItemModel.Description = core.StringPtr("JS challenge site")
				firewallRulesUpdateInputItemModel.Filter = firewallRulesUpdateInputItemFilterModel

				// Construct an instance of the UpdateFirewllRulesOptions model
				updateFirewllRulesOptionsModel := new(firewallrulesv1.UpdateFirewllRulesOptions)
				updateFirewllRulesOptionsModel.XAuthUserToken = core.StringPtr("testString")
				updateFirewllRulesOptionsModel.Crn = core.StringPtr("testString")
				updateFirewllRulesOptionsModel.ZoneIdentifier = core.StringPtr("testString")
				updateFirewllRulesOptionsModel.Accept = core.StringPtr("testString")
				updateFirewllRulesOptionsModel.FirewallRulesUpdateInputItem = []firewallrulesv1.FirewallRulesUpdateInputItem{*firewallRulesUpdateInputItemModel}
				updateFirewllRulesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := firewallRulesService.UpdateFirewllRulesWithContext(ctx, updateFirewllRulesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				firewallRulesService.DisableRetries()
				result, response, operationErr := firewallRulesService.UpdateFirewllRules(updateFirewllRulesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = firewallRulesService.UpdateFirewllRulesWithContext(ctx, updateFirewllRulesOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(updateFirewllRulesPath))
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

					Expect(req.Header["X-Auth-User-Token"]).ToNot(BeNil())
					Expect(req.Header["X-Auth-User-Token"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Accept"]).ToNot(BeNil())
					Expect(req.Header["Accept"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "*/*")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `This is a mock binary response.`)
				}))
			})
			It(`Invoke UpdateFirewllRules successfully`, func() {
				firewallRulesService, serviceErr := firewallrulesv1.NewFirewallRulesV1(&firewallrulesv1.FirewallRulesV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(firewallRulesService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := firewallRulesService.UpdateFirewllRules(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the FirewallRulesUpdateInputItemFilter model
				firewallRulesUpdateInputItemFilterModel := new(firewallrulesv1.FirewallRulesUpdateInputItemFilter)
				firewallRulesUpdateInputItemFilterModel.ID = core.StringPtr("6f58318e7fa2477a23112e8118c66f61")

				// Construct an instance of the FirewallRulesUpdateInputItem model
				firewallRulesUpdateInputItemModel := new(firewallrulesv1.FirewallRulesUpdateInputItem)
				firewallRulesUpdateInputItemModel.ID = core.StringPtr("52161eb6af4241bb9d4b32394be72fdf")
				firewallRulesUpdateInputItemModel.Action = core.StringPtr("js_challenge")
				firewallRulesUpdateInputItemModel.Paused = core.BoolPtr(false)
				firewallRulesUpdateInputItemModel.Description = core.StringPtr("JS challenge site")
				firewallRulesUpdateInputItemModel.Filter = firewallRulesUpdateInputItemFilterModel

				// Construct an instance of the UpdateFirewllRulesOptions model
				updateFirewllRulesOptionsModel := new(firewallrulesv1.UpdateFirewllRulesOptions)
				updateFirewllRulesOptionsModel.XAuthUserToken = core.StringPtr("testString")
				updateFirewllRulesOptionsModel.Crn = core.StringPtr("testString")
				updateFirewllRulesOptionsModel.ZoneIdentifier = core.StringPtr("testString")
				updateFirewllRulesOptionsModel.Accept = core.StringPtr("testString")
				updateFirewllRulesOptionsModel.FirewallRulesUpdateInputItem = []firewallrulesv1.FirewallRulesUpdateInputItem{*firewallRulesUpdateInputItemModel}
				updateFirewllRulesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = firewallRulesService.UpdateFirewllRules(updateFirewllRulesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateFirewllRules with error: Operation validation and request error`, func() {
				firewallRulesService, serviceErr := firewallrulesv1.NewFirewallRulesV1(&firewallrulesv1.FirewallRulesV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(firewallRulesService).ToNot(BeNil())

				// Construct an instance of the FirewallRulesUpdateInputItemFilter model
				firewallRulesUpdateInputItemFilterModel := new(firewallrulesv1.FirewallRulesUpdateInputItemFilter)
				firewallRulesUpdateInputItemFilterModel.ID = core.StringPtr("6f58318e7fa2477a23112e8118c66f61")

				// Construct an instance of the FirewallRulesUpdateInputItem model
				firewallRulesUpdateInputItemModel := new(firewallrulesv1.FirewallRulesUpdateInputItem)
				firewallRulesUpdateInputItemModel.ID = core.StringPtr("52161eb6af4241bb9d4b32394be72fdf")
				firewallRulesUpdateInputItemModel.Action = core.StringPtr("js_challenge")
				firewallRulesUpdateInputItemModel.Paused = core.BoolPtr(false)
				firewallRulesUpdateInputItemModel.Description = core.StringPtr("JS challenge site")
				firewallRulesUpdateInputItemModel.Filter = firewallRulesUpdateInputItemFilterModel

				// Construct an instance of the UpdateFirewllRulesOptions model
				updateFirewllRulesOptionsModel := new(firewallrulesv1.UpdateFirewllRulesOptions)
				updateFirewllRulesOptionsModel.XAuthUserToken = core.StringPtr("testString")
				updateFirewllRulesOptionsModel.Crn = core.StringPtr("testString")
				updateFirewllRulesOptionsModel.ZoneIdentifier = core.StringPtr("testString")
				updateFirewllRulesOptionsModel.Accept = core.StringPtr("testString")
				updateFirewllRulesOptionsModel.FirewallRulesUpdateInputItem = []firewallrulesv1.FirewallRulesUpdateInputItem{*firewallRulesUpdateInputItemModel}
				updateFirewllRulesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := firewallRulesService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := firewallRulesService.UpdateFirewllRules(updateFirewllRulesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateFirewllRulesOptions model with no property values
				updateFirewllRulesOptionsModelNew := new(firewallrulesv1.UpdateFirewllRulesOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = firewallRulesService.UpdateFirewllRules(updateFirewllRulesOptionsModelNew)
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
		It(`Instantiate service client`, func() {
			firewallRulesService, serviceErr := firewallrulesv1.NewFirewallRulesV1(&firewallrulesv1.FirewallRulesV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(firewallRulesService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			firewallRulesService, serviceErr := firewallrulesv1.NewFirewallRulesV1(&firewallrulesv1.FirewallRulesV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(firewallRulesService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			firewallRulesService, serviceErr := firewallrulesv1.NewFirewallRulesV1(&firewallrulesv1.FirewallRulesV1Options{
				URL: "https://firewallrulesv1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(firewallRulesService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"FIREWALL_RULES_URL": "https://firewallrulesv1/api",
				"FIREWALL_RULES_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				firewallRulesService, serviceErr := firewallrulesv1.NewFirewallRulesV1UsingExternalConfig(&firewallrulesv1.FirewallRulesV1Options{
				})
				Expect(firewallRulesService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := firewallRulesService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != firewallRulesService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(firewallRulesService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(firewallRulesService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				firewallRulesService, serviceErr := firewallrulesv1.NewFirewallRulesV1UsingExternalConfig(&firewallrulesv1.FirewallRulesV1Options{
					URL: "https://testService/api",
				})
				Expect(firewallRulesService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(firewallRulesService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := firewallRulesService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != firewallRulesService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(firewallRulesService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(firewallRulesService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				firewallRulesService, serviceErr := firewallrulesv1.NewFirewallRulesV1UsingExternalConfig(&firewallrulesv1.FirewallRulesV1Options{
				})
				err := firewallRulesService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(firewallRulesService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(firewallRulesService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := firewallRulesService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != firewallRulesService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(firewallRulesService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(firewallRulesService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"FIREWALL_RULES_URL": "https://firewallrulesv1/api",
				"FIREWALL_RULES_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			firewallRulesService, serviceErr := firewallrulesv1.NewFirewallRulesV1UsingExternalConfig(&firewallrulesv1.FirewallRulesV1Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(firewallRulesService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"FIREWALL_RULES_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			firewallRulesService, serviceErr := firewallrulesv1.NewFirewallRulesV1UsingExternalConfig(&firewallrulesv1.FirewallRulesV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(firewallRulesService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = firewallrulesv1.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})

	Describe(`DeleteFirewallRules(deleteFirewallRulesOptions *DeleteFirewallRulesOptions)`, func() {
		deleteFirewallRulesPath := "/v1/testString/zones/testString/firewall/rules"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteFirewallRulesPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.Header["X-Auth-User-Token"]).ToNot(BeNil())
					Expect(req.Header["X-Auth-User-Token"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Accept"]).ToNot(BeNil())
					Expect(req.Header["Accept"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["id"]).To(Equal([]string{"f2d427378e7542acb295380d352e2ebd"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "*/*")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `This is a mock binary response.`)
				}))
			})
			It(`Invoke DeleteFirewallRules successfully with retries`, func() {
				firewallRulesService, serviceErr := firewallrulesv1.NewFirewallRulesV1(&firewallrulesv1.FirewallRulesV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(firewallRulesService).ToNot(BeNil())
				firewallRulesService.EnableRetries(0, 0)

				// Construct an instance of the DeleteFirewallRulesOptions model
				deleteFirewallRulesOptionsModel := new(firewallrulesv1.DeleteFirewallRulesOptions)
				deleteFirewallRulesOptionsModel.XAuthUserToken = core.StringPtr("testString")
				deleteFirewallRulesOptionsModel.Crn = core.StringPtr("testString")
				deleteFirewallRulesOptionsModel.ZoneIdentifier = core.StringPtr("testString")
				deleteFirewallRulesOptionsModel.ID = core.StringPtr("f2d427378e7542acb295380d352e2ebd")
				deleteFirewallRulesOptionsModel.Accept = core.StringPtr("testString")
				deleteFirewallRulesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := firewallRulesService.DeleteFirewallRulesWithContext(ctx, deleteFirewallRulesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				firewallRulesService.DisableRetries()
				result, response, operationErr := firewallRulesService.DeleteFirewallRules(deleteFirewallRulesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = firewallRulesService.DeleteFirewallRulesWithContext(ctx, deleteFirewallRulesOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(deleteFirewallRulesPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.Header["X-Auth-User-Token"]).ToNot(BeNil())
					Expect(req.Header["X-Auth-User-Token"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Accept"]).ToNot(BeNil())
					Expect(req.Header["Accept"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["id"]).To(Equal([]string{"f2d427378e7542acb295380d352e2ebd"}))
					// Set mock response
					res.Header().Set("Content-type", "*/*")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `This is a mock binary response.`)
				}))
			})
			It(`Invoke DeleteFirewallRules successfully`, func() {
				firewallRulesService, serviceErr := firewallrulesv1.NewFirewallRulesV1(&firewallrulesv1.FirewallRulesV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(firewallRulesService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := firewallRulesService.DeleteFirewallRules(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DeleteFirewallRulesOptions model
				deleteFirewallRulesOptionsModel := new(firewallrulesv1.DeleteFirewallRulesOptions)
				deleteFirewallRulesOptionsModel.XAuthUserToken = core.StringPtr("testString")
				deleteFirewallRulesOptionsModel.Crn = core.StringPtr("testString")
				deleteFirewallRulesOptionsModel.ZoneIdentifier = core.StringPtr("testString")
				deleteFirewallRulesOptionsModel.ID = core.StringPtr("f2d427378e7542acb295380d352e2ebd")
				deleteFirewallRulesOptionsModel.Accept = core.StringPtr("testString")
				deleteFirewallRulesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = firewallRulesService.DeleteFirewallRules(deleteFirewallRulesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke DeleteFirewallRules with error: Operation validation and request error`, func() {
				firewallRulesService, serviceErr := firewallrulesv1.NewFirewallRulesV1(&firewallrulesv1.FirewallRulesV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(firewallRulesService).ToNot(BeNil())

				// Construct an instance of the DeleteFirewallRulesOptions model
				deleteFirewallRulesOptionsModel := new(firewallrulesv1.DeleteFirewallRulesOptions)
				deleteFirewallRulesOptionsModel.XAuthUserToken = core.StringPtr("testString")
				deleteFirewallRulesOptionsModel.Crn = core.StringPtr("testString")
				deleteFirewallRulesOptionsModel.ZoneIdentifier = core.StringPtr("testString")
				deleteFirewallRulesOptionsModel.ID = core.StringPtr("f2d427378e7542acb295380d352e2ebd")
				deleteFirewallRulesOptionsModel.Accept = core.StringPtr("testString")
				deleteFirewallRulesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := firewallRulesService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := firewallRulesService.DeleteFirewallRules(deleteFirewallRulesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the DeleteFirewallRulesOptions model with no property values
				deleteFirewallRulesOptionsModelNew := new(firewallrulesv1.DeleteFirewallRulesOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = firewallRulesService.DeleteFirewallRules(deleteFirewallRulesOptionsModelNew)
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
		It(`Instantiate service client`, func() {
			firewallRulesService, serviceErr := firewallrulesv1.NewFirewallRulesV1(&firewallrulesv1.FirewallRulesV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(firewallRulesService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			firewallRulesService, serviceErr := firewallrulesv1.NewFirewallRulesV1(&firewallrulesv1.FirewallRulesV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(firewallRulesService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			firewallRulesService, serviceErr := firewallrulesv1.NewFirewallRulesV1(&firewallrulesv1.FirewallRulesV1Options{
				URL: "https://firewallrulesv1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(firewallRulesService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"FIREWALL_RULES_URL": "https://firewallrulesv1/api",
				"FIREWALL_RULES_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				firewallRulesService, serviceErr := firewallrulesv1.NewFirewallRulesV1UsingExternalConfig(&firewallrulesv1.FirewallRulesV1Options{
				})
				Expect(firewallRulesService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := firewallRulesService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != firewallRulesService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(firewallRulesService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(firewallRulesService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				firewallRulesService, serviceErr := firewallrulesv1.NewFirewallRulesV1UsingExternalConfig(&firewallrulesv1.FirewallRulesV1Options{
					URL: "https://testService/api",
				})
				Expect(firewallRulesService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(firewallRulesService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := firewallRulesService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != firewallRulesService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(firewallRulesService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(firewallRulesService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				firewallRulesService, serviceErr := firewallrulesv1.NewFirewallRulesV1UsingExternalConfig(&firewallrulesv1.FirewallRulesV1Options{
				})
				err := firewallRulesService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(firewallRulesService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(firewallRulesService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := firewallRulesService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != firewallRulesService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(firewallRulesService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(firewallRulesService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"FIREWALL_RULES_URL": "https://firewallrulesv1/api",
				"FIREWALL_RULES_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			firewallRulesService, serviceErr := firewallrulesv1.NewFirewallRulesV1UsingExternalConfig(&firewallrulesv1.FirewallRulesV1Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(firewallRulesService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"FIREWALL_RULES_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			firewallRulesService, serviceErr := firewallrulesv1.NewFirewallRulesV1UsingExternalConfig(&firewallrulesv1.FirewallRulesV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(firewallRulesService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = firewallrulesv1.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})

	Describe(`GetFirewallRule(getFirewallRuleOptions *GetFirewallRuleOptions)`, func() {
		getFirewallRulePath := "/v1/testString/zones/testString/firewall/rules/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getFirewallRulePath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Auth-User-Token"]).ToNot(BeNil())
					Expect(req.Header["X-Auth-User-Token"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Accept"]).ToNot(BeNil())
					Expect(req.Header["Accept"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "*/*")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `This is a mock binary response.`)
				}))
			})
			It(`Invoke GetFirewallRule successfully with retries`, func() {
				firewallRulesService, serviceErr := firewallrulesv1.NewFirewallRulesV1(&firewallrulesv1.FirewallRulesV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(firewallRulesService).ToNot(BeNil())
				firewallRulesService.EnableRetries(0, 0)

				// Construct an instance of the GetFirewallRuleOptions model
				getFirewallRuleOptionsModel := new(firewallrulesv1.GetFirewallRuleOptions)
				getFirewallRuleOptionsModel.XAuthUserToken = core.StringPtr("testString")
				getFirewallRuleOptionsModel.Crn = core.StringPtr("testString")
				getFirewallRuleOptionsModel.ZoneIdentifier = core.StringPtr("testString")
				getFirewallRuleOptionsModel.FirewallRuleIdentifier = core.StringPtr("testString")
				getFirewallRuleOptionsModel.Accept = core.StringPtr("testString")
				getFirewallRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := firewallRulesService.GetFirewallRuleWithContext(ctx, getFirewallRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				firewallRulesService.DisableRetries()
				result, response, operationErr := firewallRulesService.GetFirewallRule(getFirewallRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = firewallRulesService.GetFirewallRuleWithContext(ctx, getFirewallRuleOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getFirewallRulePath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Auth-User-Token"]).ToNot(BeNil())
					Expect(req.Header["X-Auth-User-Token"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Accept"]).ToNot(BeNil())
					Expect(req.Header["Accept"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "*/*")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `This is a mock binary response.`)
				}))
			})
			It(`Invoke GetFirewallRule successfully`, func() {
				firewallRulesService, serviceErr := firewallrulesv1.NewFirewallRulesV1(&firewallrulesv1.FirewallRulesV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(firewallRulesService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := firewallRulesService.GetFirewallRule(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetFirewallRuleOptions model
				getFirewallRuleOptionsModel := new(firewallrulesv1.GetFirewallRuleOptions)
				getFirewallRuleOptionsModel.XAuthUserToken = core.StringPtr("testString")
				getFirewallRuleOptionsModel.Crn = core.StringPtr("testString")
				getFirewallRuleOptionsModel.ZoneIdentifier = core.StringPtr("testString")
				getFirewallRuleOptionsModel.FirewallRuleIdentifier = core.StringPtr("testString")
				getFirewallRuleOptionsModel.Accept = core.StringPtr("testString")
				getFirewallRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = firewallRulesService.GetFirewallRule(getFirewallRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetFirewallRule with error: Operation validation and request error`, func() {
				firewallRulesService, serviceErr := firewallrulesv1.NewFirewallRulesV1(&firewallrulesv1.FirewallRulesV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(firewallRulesService).ToNot(BeNil())

				// Construct an instance of the GetFirewallRuleOptions model
				getFirewallRuleOptionsModel := new(firewallrulesv1.GetFirewallRuleOptions)
				getFirewallRuleOptionsModel.XAuthUserToken = core.StringPtr("testString")
				getFirewallRuleOptionsModel.Crn = core.StringPtr("testString")
				getFirewallRuleOptionsModel.ZoneIdentifier = core.StringPtr("testString")
				getFirewallRuleOptionsModel.FirewallRuleIdentifier = core.StringPtr("testString")
				getFirewallRuleOptionsModel.Accept = core.StringPtr("testString")
				getFirewallRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := firewallRulesService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := firewallRulesService.GetFirewallRule(getFirewallRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetFirewallRuleOptions model with no property values
				getFirewallRuleOptionsModelNew := new(firewallrulesv1.GetFirewallRuleOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = firewallRulesService.GetFirewallRule(getFirewallRuleOptionsModelNew)
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
		It(`Instantiate service client`, func() {
			firewallRulesService, serviceErr := firewallrulesv1.NewFirewallRulesV1(&firewallrulesv1.FirewallRulesV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(firewallRulesService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			firewallRulesService, serviceErr := firewallrulesv1.NewFirewallRulesV1(&firewallrulesv1.FirewallRulesV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(firewallRulesService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			firewallRulesService, serviceErr := firewallrulesv1.NewFirewallRulesV1(&firewallrulesv1.FirewallRulesV1Options{
				URL: "https://firewallrulesv1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(firewallRulesService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"FIREWALL_RULES_URL": "https://firewallrulesv1/api",
				"FIREWALL_RULES_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				firewallRulesService, serviceErr := firewallrulesv1.NewFirewallRulesV1UsingExternalConfig(&firewallrulesv1.FirewallRulesV1Options{
				})
				Expect(firewallRulesService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := firewallRulesService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != firewallRulesService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(firewallRulesService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(firewallRulesService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				firewallRulesService, serviceErr := firewallrulesv1.NewFirewallRulesV1UsingExternalConfig(&firewallrulesv1.FirewallRulesV1Options{
					URL: "https://testService/api",
				})
				Expect(firewallRulesService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(firewallRulesService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := firewallRulesService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != firewallRulesService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(firewallRulesService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(firewallRulesService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				firewallRulesService, serviceErr := firewallrulesv1.NewFirewallRulesV1UsingExternalConfig(&firewallrulesv1.FirewallRulesV1Options{
				})
				err := firewallRulesService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(firewallRulesService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(firewallRulesService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := firewallRulesService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != firewallRulesService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(firewallRulesService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(firewallRulesService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"FIREWALL_RULES_URL": "https://firewallrulesv1/api",
				"FIREWALL_RULES_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			firewallRulesService, serviceErr := firewallrulesv1.NewFirewallRulesV1UsingExternalConfig(&firewallrulesv1.FirewallRulesV1Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(firewallRulesService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"FIREWALL_RULES_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			firewallRulesService, serviceErr := firewallrulesv1.NewFirewallRulesV1UsingExternalConfig(&firewallrulesv1.FirewallRulesV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(firewallRulesService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = firewallrulesv1.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})

	Describe(`UpdateFirewallRule(updateFirewallRuleOptions *UpdateFirewallRuleOptions)`, func() {
		updateFirewallRulePath := "/v1/testString/zones/testString/firewall/rules/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateFirewallRulePath))
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

					Expect(req.Header["X-Auth-User-Token"]).ToNot(BeNil())
					Expect(req.Header["X-Auth-User-Token"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Accept"]).ToNot(BeNil())
					Expect(req.Header["Accept"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "*/*")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `This is a mock binary response.`)
				}))
			})
			It(`Invoke UpdateFirewallRule successfully with retries`, func() {
				firewallRulesService, serviceErr := firewallrulesv1.NewFirewallRulesV1(&firewallrulesv1.FirewallRulesV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(firewallRulesService).ToNot(BeNil())
				firewallRulesService.EnableRetries(0, 0)

				// Construct an instance of the FirewallRuleUpdateInputFilter model
				firewallRuleUpdateInputFilterModel := new(firewallrulesv1.FirewallRuleUpdateInputFilter)
				firewallRuleUpdateInputFilterModel.ID = core.StringPtr("6f58318e7fa2477a23112e8118c66f61")

				// Construct an instance of the UpdateFirewallRuleOptions model
				updateFirewallRuleOptionsModel := new(firewallrulesv1.UpdateFirewallRuleOptions)
				updateFirewallRuleOptionsModel.XAuthUserToken = core.StringPtr("testString")
				updateFirewallRuleOptionsModel.Crn = core.StringPtr("testString")
				updateFirewallRuleOptionsModel.ZoneIdentifier = core.StringPtr("testString")
				updateFirewallRuleOptionsModel.FirewallRuleIdentifier = core.StringPtr("testString")
				updateFirewallRuleOptionsModel.Accept = core.StringPtr("testString")
				updateFirewallRuleOptionsModel.Action = core.StringPtr("js_challenge")
				updateFirewallRuleOptionsModel.Paused = core.BoolPtr(false)
				updateFirewallRuleOptionsModel.Description = core.StringPtr("JS challenge site")
				updateFirewallRuleOptionsModel.Filter = firewallRuleUpdateInputFilterModel
				updateFirewallRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := firewallRulesService.UpdateFirewallRuleWithContext(ctx, updateFirewallRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				firewallRulesService.DisableRetries()
				result, response, operationErr := firewallRulesService.UpdateFirewallRule(updateFirewallRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = firewallRulesService.UpdateFirewallRuleWithContext(ctx, updateFirewallRuleOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(updateFirewallRulePath))
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

					Expect(req.Header["X-Auth-User-Token"]).ToNot(BeNil())
					Expect(req.Header["X-Auth-User-Token"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Accept"]).ToNot(BeNil())
					Expect(req.Header["Accept"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "*/*")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `This is a mock binary response.`)
				}))
			})
			It(`Invoke UpdateFirewallRule successfully`, func() {
				firewallRulesService, serviceErr := firewallrulesv1.NewFirewallRulesV1(&firewallrulesv1.FirewallRulesV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(firewallRulesService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := firewallRulesService.UpdateFirewallRule(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the FirewallRuleUpdateInputFilter model
				firewallRuleUpdateInputFilterModel := new(firewallrulesv1.FirewallRuleUpdateInputFilter)
				firewallRuleUpdateInputFilterModel.ID = core.StringPtr("6f58318e7fa2477a23112e8118c66f61")

				// Construct an instance of the UpdateFirewallRuleOptions model
				updateFirewallRuleOptionsModel := new(firewallrulesv1.UpdateFirewallRuleOptions)
				updateFirewallRuleOptionsModel.XAuthUserToken = core.StringPtr("testString")
				updateFirewallRuleOptionsModel.Crn = core.StringPtr("testString")
				updateFirewallRuleOptionsModel.ZoneIdentifier = core.StringPtr("testString")
				updateFirewallRuleOptionsModel.FirewallRuleIdentifier = core.StringPtr("testString")
				updateFirewallRuleOptionsModel.Accept = core.StringPtr("testString")
				updateFirewallRuleOptionsModel.Action = core.StringPtr("js_challenge")
				updateFirewallRuleOptionsModel.Paused = core.BoolPtr(false)
				updateFirewallRuleOptionsModel.Description = core.StringPtr("JS challenge site")
				updateFirewallRuleOptionsModel.Filter = firewallRuleUpdateInputFilterModel
				updateFirewallRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = firewallRulesService.UpdateFirewallRule(updateFirewallRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateFirewallRule with error: Operation validation and request error`, func() {
				firewallRulesService, serviceErr := firewallrulesv1.NewFirewallRulesV1(&firewallrulesv1.FirewallRulesV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(firewallRulesService).ToNot(BeNil())

				// Construct an instance of the FirewallRuleUpdateInputFilter model
				firewallRuleUpdateInputFilterModel := new(firewallrulesv1.FirewallRuleUpdateInputFilter)
				firewallRuleUpdateInputFilterModel.ID = core.StringPtr("6f58318e7fa2477a23112e8118c66f61")

				// Construct an instance of the UpdateFirewallRuleOptions model
				updateFirewallRuleOptionsModel := new(firewallrulesv1.UpdateFirewallRuleOptions)
				updateFirewallRuleOptionsModel.XAuthUserToken = core.StringPtr("testString")
				updateFirewallRuleOptionsModel.Crn = core.StringPtr("testString")
				updateFirewallRuleOptionsModel.ZoneIdentifier = core.StringPtr("testString")
				updateFirewallRuleOptionsModel.FirewallRuleIdentifier = core.StringPtr("testString")
				updateFirewallRuleOptionsModel.Accept = core.StringPtr("testString")
				updateFirewallRuleOptionsModel.Action = core.StringPtr("js_challenge")
				updateFirewallRuleOptionsModel.Paused = core.BoolPtr(false)
				updateFirewallRuleOptionsModel.Description = core.StringPtr("JS challenge site")
				updateFirewallRuleOptionsModel.Filter = firewallRuleUpdateInputFilterModel
				updateFirewallRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := firewallRulesService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := firewallRulesService.UpdateFirewallRule(updateFirewallRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateFirewallRuleOptions model with no property values
				updateFirewallRuleOptionsModelNew := new(firewallrulesv1.UpdateFirewallRuleOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = firewallRulesService.UpdateFirewallRule(updateFirewallRuleOptionsModelNew)
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
		It(`Instantiate service client`, func() {
			firewallRulesService, serviceErr := firewallrulesv1.NewFirewallRulesV1(&firewallrulesv1.FirewallRulesV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(firewallRulesService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			firewallRulesService, serviceErr := firewallrulesv1.NewFirewallRulesV1(&firewallrulesv1.FirewallRulesV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(firewallRulesService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			firewallRulesService, serviceErr := firewallrulesv1.NewFirewallRulesV1(&firewallrulesv1.FirewallRulesV1Options{
				URL: "https://firewallrulesv1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(firewallRulesService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"FIREWALL_RULES_URL": "https://firewallrulesv1/api",
				"FIREWALL_RULES_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				firewallRulesService, serviceErr := firewallrulesv1.NewFirewallRulesV1UsingExternalConfig(&firewallrulesv1.FirewallRulesV1Options{
				})
				Expect(firewallRulesService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := firewallRulesService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != firewallRulesService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(firewallRulesService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(firewallRulesService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				firewallRulesService, serviceErr := firewallrulesv1.NewFirewallRulesV1UsingExternalConfig(&firewallrulesv1.FirewallRulesV1Options{
					URL: "https://testService/api",
				})
				Expect(firewallRulesService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(firewallRulesService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := firewallRulesService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != firewallRulesService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(firewallRulesService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(firewallRulesService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				firewallRulesService, serviceErr := firewallrulesv1.NewFirewallRulesV1UsingExternalConfig(&firewallrulesv1.FirewallRulesV1Options{
				})
				err := firewallRulesService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(firewallRulesService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(firewallRulesService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := firewallRulesService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != firewallRulesService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(firewallRulesService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(firewallRulesService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"FIREWALL_RULES_URL": "https://firewallrulesv1/api",
				"FIREWALL_RULES_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			firewallRulesService, serviceErr := firewallrulesv1.NewFirewallRulesV1UsingExternalConfig(&firewallrulesv1.FirewallRulesV1Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(firewallRulesService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"FIREWALL_RULES_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			firewallRulesService, serviceErr := firewallrulesv1.NewFirewallRulesV1UsingExternalConfig(&firewallrulesv1.FirewallRulesV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(firewallRulesService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = firewallrulesv1.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})

	Describe(`DeleteFirewallRule(deleteFirewallRuleOptions *DeleteFirewallRuleOptions)`, func() {
		deleteFirewallRulePath := "/v1/testString/zones/testString/firewall/rules/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteFirewallRulePath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.Header["X-Auth-User-Token"]).ToNot(BeNil())
					Expect(req.Header["X-Auth-User-Token"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Accept"]).ToNot(BeNil())
					Expect(req.Header["Accept"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "*/*")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `This is a mock binary response.`)
				}))
			})
			It(`Invoke DeleteFirewallRule successfully with retries`, func() {
				firewallRulesService, serviceErr := firewallrulesv1.NewFirewallRulesV1(&firewallrulesv1.FirewallRulesV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(firewallRulesService).ToNot(BeNil())
				firewallRulesService.EnableRetries(0, 0)

				// Construct an instance of the DeleteFirewallRuleOptions model
				deleteFirewallRuleOptionsModel := new(firewallrulesv1.DeleteFirewallRuleOptions)
				deleteFirewallRuleOptionsModel.XAuthUserToken = core.StringPtr("testString")
				deleteFirewallRuleOptionsModel.Crn = core.StringPtr("testString")
				deleteFirewallRuleOptionsModel.ZoneIdentifier = core.StringPtr("testString")
				deleteFirewallRuleOptionsModel.FirewallRuleIdentifier = core.StringPtr("testString")
				deleteFirewallRuleOptionsModel.Accept = core.StringPtr("testString")
				deleteFirewallRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := firewallRulesService.DeleteFirewallRuleWithContext(ctx, deleteFirewallRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				firewallRulesService.DisableRetries()
				result, response, operationErr := firewallRulesService.DeleteFirewallRule(deleteFirewallRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = firewallRulesService.DeleteFirewallRuleWithContext(ctx, deleteFirewallRuleOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(deleteFirewallRulePath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.Header["X-Auth-User-Token"]).ToNot(BeNil())
					Expect(req.Header["X-Auth-User-Token"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Accept"]).ToNot(BeNil())
					Expect(req.Header["Accept"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "*/*")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `This is a mock binary response.`)
				}))
			})
			It(`Invoke DeleteFirewallRule successfully`, func() {
				firewallRulesService, serviceErr := firewallrulesv1.NewFirewallRulesV1(&firewallrulesv1.FirewallRulesV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(firewallRulesService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := firewallRulesService.DeleteFirewallRule(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DeleteFirewallRuleOptions model
				deleteFirewallRuleOptionsModel := new(firewallrulesv1.DeleteFirewallRuleOptions)
				deleteFirewallRuleOptionsModel.XAuthUserToken = core.StringPtr("testString")
				deleteFirewallRuleOptionsModel.Crn = core.StringPtr("testString")
				deleteFirewallRuleOptionsModel.ZoneIdentifier = core.StringPtr("testString")
				deleteFirewallRuleOptionsModel.FirewallRuleIdentifier = core.StringPtr("testString")
				deleteFirewallRuleOptionsModel.Accept = core.StringPtr("testString")
				deleteFirewallRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = firewallRulesService.DeleteFirewallRule(deleteFirewallRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke DeleteFirewallRule with error: Operation validation and request error`, func() {
				firewallRulesService, serviceErr := firewallrulesv1.NewFirewallRulesV1(&firewallrulesv1.FirewallRulesV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(firewallRulesService).ToNot(BeNil())

				// Construct an instance of the DeleteFirewallRuleOptions model
				deleteFirewallRuleOptionsModel := new(firewallrulesv1.DeleteFirewallRuleOptions)
				deleteFirewallRuleOptionsModel.XAuthUserToken = core.StringPtr("testString")
				deleteFirewallRuleOptionsModel.Crn = core.StringPtr("testString")
				deleteFirewallRuleOptionsModel.ZoneIdentifier = core.StringPtr("testString")
				deleteFirewallRuleOptionsModel.FirewallRuleIdentifier = core.StringPtr("testString")
				deleteFirewallRuleOptionsModel.Accept = core.StringPtr("testString")
				deleteFirewallRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := firewallRulesService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := firewallRulesService.DeleteFirewallRule(deleteFirewallRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the DeleteFirewallRuleOptions model with no property values
				deleteFirewallRuleOptionsModelNew := new(firewallrulesv1.DeleteFirewallRuleOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = firewallRulesService.DeleteFirewallRule(deleteFirewallRuleOptionsModelNew)
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
			firewallRulesService, _ := firewallrulesv1.NewFirewallRulesV1(&firewallrulesv1.FirewallRulesV1Options{
				URL:           "http://firewallrulesv1modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
			})
			It(`Invoke NewCreateFirewallRulesOptions successfully`, func() {
				// Construct an instance of the FirewallRuleInputWithFilterIdFilter model
				firewallRuleInputWithFilterIdFilterModel := new(firewallrulesv1.FirewallRuleInputWithFilterIdFilter)
				Expect(firewallRuleInputWithFilterIdFilterModel).ToNot(BeNil())
				firewallRuleInputWithFilterIdFilterModel.ID = core.StringPtr("6f58318e7fa2477a23112e8118c66f61")
				Expect(firewallRuleInputWithFilterIdFilterModel.ID).To(Equal(core.StringPtr("6f58318e7fa2477a23112e8118c66f61")))

				// Construct an instance of the FirewallRuleInputWithFilterID model
				firewallRuleInputWithFilterIdModel := new(firewallrulesv1.FirewallRuleInputWithFilterID)
				Expect(firewallRuleInputWithFilterIdModel).ToNot(BeNil())
				firewallRuleInputWithFilterIdModel.Filter = firewallRuleInputWithFilterIdFilterModel
				firewallRuleInputWithFilterIdModel.Action = core.StringPtr("js_challenge")
				firewallRuleInputWithFilterIdModel.Description = core.StringPtr("JS challenge site")
				Expect(firewallRuleInputWithFilterIdModel.Filter).To(Equal(firewallRuleInputWithFilterIdFilterModel))
				Expect(firewallRuleInputWithFilterIdModel.Action).To(Equal(core.StringPtr("js_challenge")))
				Expect(firewallRuleInputWithFilterIdModel.Description).To(Equal(core.StringPtr("JS challenge site")))

				// Construct an instance of the CreateFirewallRulesOptions model
				xAuthUserToken := "testString"
				crn := "testString"
				zoneIdentifier := "testString"
				createFirewallRulesOptionsModel := firewallRulesService.NewCreateFirewallRulesOptions(xAuthUserToken, crn, zoneIdentifier)
				createFirewallRulesOptionsModel.SetXAuthUserToken("testString")
				createFirewallRulesOptionsModel.SetCrn("testString")
				createFirewallRulesOptionsModel.SetZoneIdentifier("testString")
				createFirewallRulesOptionsModel.SetAccept("testString")
				createFirewallRulesOptionsModel.SetFirewallRuleInputWithFilterID([]firewallrulesv1.FirewallRuleInputWithFilterID{*firewallRuleInputWithFilterIdModel})
				createFirewallRulesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createFirewallRulesOptionsModel).ToNot(BeNil())
				Expect(createFirewallRulesOptionsModel.XAuthUserToken).To(Equal(core.StringPtr("testString")))
				Expect(createFirewallRulesOptionsModel.Crn).To(Equal(core.StringPtr("testString")))
				Expect(createFirewallRulesOptionsModel.ZoneIdentifier).To(Equal(core.StringPtr("testString")))
				Expect(createFirewallRulesOptionsModel.Accept).To(Equal(core.StringPtr("testString")))
				Expect(createFirewallRulesOptionsModel.FirewallRuleInputWithFilterID).To(Equal([]firewallrulesv1.FirewallRuleInputWithFilterID{*firewallRuleInputWithFilterIdModel}))
				Expect(createFirewallRulesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteFirewallRuleOptions successfully`, func() {
				// Construct an instance of the DeleteFirewallRuleOptions model
				xAuthUserToken := "testString"
				crn := "testString"
				zoneIdentifier := "testString"
				firewallRuleIdentifier := "testString"
				deleteFirewallRuleOptionsModel := firewallRulesService.NewDeleteFirewallRuleOptions(xAuthUserToken, crn, zoneIdentifier, firewallRuleIdentifier)
				deleteFirewallRuleOptionsModel.SetXAuthUserToken("testString")
				deleteFirewallRuleOptionsModel.SetCrn("testString")
				deleteFirewallRuleOptionsModel.SetZoneIdentifier("testString")
				deleteFirewallRuleOptionsModel.SetFirewallRuleIdentifier("testString")
				deleteFirewallRuleOptionsModel.SetAccept("testString")
				deleteFirewallRuleOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteFirewallRuleOptionsModel).ToNot(BeNil())
				Expect(deleteFirewallRuleOptionsModel.XAuthUserToken).To(Equal(core.StringPtr("testString")))
				Expect(deleteFirewallRuleOptionsModel.Crn).To(Equal(core.StringPtr("testString")))
				Expect(deleteFirewallRuleOptionsModel.ZoneIdentifier).To(Equal(core.StringPtr("testString")))
				Expect(deleteFirewallRuleOptionsModel.FirewallRuleIdentifier).To(Equal(core.StringPtr("testString")))
				Expect(deleteFirewallRuleOptionsModel.Accept).To(Equal(core.StringPtr("testString")))
				Expect(deleteFirewallRuleOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteFirewallRulesOptions successfully`, func() {
				// Construct an instance of the DeleteFirewallRulesOptions model
				xAuthUserToken := "testString"
				crn := "testString"
				zoneIdentifier := "testString"
				id := "f2d427378e7542acb295380d352e2ebd"
				deleteFirewallRulesOptionsModel := firewallRulesService.NewDeleteFirewallRulesOptions(xAuthUserToken, crn, zoneIdentifier, id)
				deleteFirewallRulesOptionsModel.SetXAuthUserToken("testString")
				deleteFirewallRulesOptionsModel.SetCrn("testString")
				deleteFirewallRulesOptionsModel.SetZoneIdentifier("testString")
				deleteFirewallRulesOptionsModel.SetID("f2d427378e7542acb295380d352e2ebd")
				deleteFirewallRulesOptionsModel.SetAccept("testString")
				deleteFirewallRulesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteFirewallRulesOptionsModel).ToNot(BeNil())
				Expect(deleteFirewallRulesOptionsModel.XAuthUserToken).To(Equal(core.StringPtr("testString")))
				Expect(deleteFirewallRulesOptionsModel.Crn).To(Equal(core.StringPtr("testString")))
				Expect(deleteFirewallRulesOptionsModel.ZoneIdentifier).To(Equal(core.StringPtr("testString")))
				Expect(deleteFirewallRulesOptionsModel.ID).To(Equal(core.StringPtr("f2d427378e7542acb295380d352e2ebd")))
				Expect(deleteFirewallRulesOptionsModel.Accept).To(Equal(core.StringPtr("testString")))
				Expect(deleteFirewallRulesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewFirewallRuleInputWithFilterIdFilter successfully`, func() {
				id := "6f58318e7fa2477a23112e8118c66f61"
				model, err := firewallRulesService.NewFirewallRuleInputWithFilterIdFilter(id)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewFirewallRuleUpdateInputFilter successfully`, func() {
				id := "6f58318e7fa2477a23112e8118c66f61"
				model, err := firewallRulesService.NewFirewallRuleUpdateInputFilter(id)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewFirewallRulesUpdateInputItem successfully`, func() {
				id := "52161eb6af4241bb9d4b32394be72fdf"
				action := "js_challenge"
				model, err := firewallRulesService.NewFirewallRulesUpdateInputItem(id, action)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewFirewallRulesUpdateInputItemFilter successfully`, func() {
				id := "6f58318e7fa2477a23112e8118c66f61"
				model, err := firewallRulesService.NewFirewallRulesUpdateInputItemFilter(id)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewGetFirewallRuleOptions successfully`, func() {
				// Construct an instance of the GetFirewallRuleOptions model
				xAuthUserToken := "testString"
				crn := "testString"
				zoneIdentifier := "testString"
				firewallRuleIdentifier := "testString"
				getFirewallRuleOptionsModel := firewallRulesService.NewGetFirewallRuleOptions(xAuthUserToken, crn, zoneIdentifier, firewallRuleIdentifier)
				getFirewallRuleOptionsModel.SetXAuthUserToken("testString")
				getFirewallRuleOptionsModel.SetCrn("testString")
				getFirewallRuleOptionsModel.SetZoneIdentifier("testString")
				getFirewallRuleOptionsModel.SetFirewallRuleIdentifier("testString")
				getFirewallRuleOptionsModel.SetAccept("testString")
				getFirewallRuleOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getFirewallRuleOptionsModel).ToNot(BeNil())
				Expect(getFirewallRuleOptionsModel.XAuthUserToken).To(Equal(core.StringPtr("testString")))
				Expect(getFirewallRuleOptionsModel.Crn).To(Equal(core.StringPtr("testString")))
				Expect(getFirewallRuleOptionsModel.ZoneIdentifier).To(Equal(core.StringPtr("testString")))
				Expect(getFirewallRuleOptionsModel.FirewallRuleIdentifier).To(Equal(core.StringPtr("testString")))
				Expect(getFirewallRuleOptionsModel.Accept).To(Equal(core.StringPtr("testString")))
				Expect(getFirewallRuleOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListAllFirewallRulesOptions successfully`, func() {
				// Construct an instance of the ListAllFirewallRulesOptions model
				xAuthUserToken := "testString"
				crn := "testString"
				zoneIdentifier := "testString"
				listAllFirewallRulesOptionsModel := firewallRulesService.NewListAllFirewallRulesOptions(xAuthUserToken, crn, zoneIdentifier)
				listAllFirewallRulesOptionsModel.SetXAuthUserToken("testString")
				listAllFirewallRulesOptionsModel.SetCrn("testString")
				listAllFirewallRulesOptionsModel.SetZoneIdentifier("testString")
				listAllFirewallRulesOptionsModel.SetAccept("testString")
				listAllFirewallRulesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listAllFirewallRulesOptionsModel).ToNot(BeNil())
				Expect(listAllFirewallRulesOptionsModel.XAuthUserToken).To(Equal(core.StringPtr("testString")))
				Expect(listAllFirewallRulesOptionsModel.Crn).To(Equal(core.StringPtr("testString")))
				Expect(listAllFirewallRulesOptionsModel.ZoneIdentifier).To(Equal(core.StringPtr("testString")))
				Expect(listAllFirewallRulesOptionsModel.Accept).To(Equal(core.StringPtr("testString")))
				Expect(listAllFirewallRulesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateFirewallRuleOptions successfully`, func() {
				// Construct an instance of the FirewallRuleUpdateInputFilter model
				firewallRuleUpdateInputFilterModel := new(firewallrulesv1.FirewallRuleUpdateInputFilter)
				Expect(firewallRuleUpdateInputFilterModel).ToNot(BeNil())
				firewallRuleUpdateInputFilterModel.ID = core.StringPtr("6f58318e7fa2477a23112e8118c66f61")
				Expect(firewallRuleUpdateInputFilterModel.ID).To(Equal(core.StringPtr("6f58318e7fa2477a23112e8118c66f61")))

				// Construct an instance of the UpdateFirewallRuleOptions model
				xAuthUserToken := "testString"
				crn := "testString"
				zoneIdentifier := "testString"
				firewallRuleIdentifier := "testString"
				updateFirewallRuleOptionsModel := firewallRulesService.NewUpdateFirewallRuleOptions(xAuthUserToken, crn, zoneIdentifier, firewallRuleIdentifier)
				updateFirewallRuleOptionsModel.SetXAuthUserToken("testString")
				updateFirewallRuleOptionsModel.SetCrn("testString")
				updateFirewallRuleOptionsModel.SetZoneIdentifier("testString")
				updateFirewallRuleOptionsModel.SetFirewallRuleIdentifier("testString")
				updateFirewallRuleOptionsModel.SetAccept("testString")
				updateFirewallRuleOptionsModel.SetAction("js_challenge")
				updateFirewallRuleOptionsModel.SetPaused(false)
				updateFirewallRuleOptionsModel.SetDescription("JS challenge site")
				updateFirewallRuleOptionsModel.SetFilter(firewallRuleUpdateInputFilterModel)
				updateFirewallRuleOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateFirewallRuleOptionsModel).ToNot(BeNil())
				Expect(updateFirewallRuleOptionsModel.XAuthUserToken).To(Equal(core.StringPtr("testString")))
				Expect(updateFirewallRuleOptionsModel.Crn).To(Equal(core.StringPtr("testString")))
				Expect(updateFirewallRuleOptionsModel.ZoneIdentifier).To(Equal(core.StringPtr("testString")))
				Expect(updateFirewallRuleOptionsModel.FirewallRuleIdentifier).To(Equal(core.StringPtr("testString")))
				Expect(updateFirewallRuleOptionsModel.Accept).To(Equal(core.StringPtr("testString")))
				Expect(updateFirewallRuleOptionsModel.Action).To(Equal(core.StringPtr("js_challenge")))
				Expect(updateFirewallRuleOptionsModel.Paused).To(Equal(core.BoolPtr(false)))
				Expect(updateFirewallRuleOptionsModel.Description).To(Equal(core.StringPtr("JS challenge site")))
				Expect(updateFirewallRuleOptionsModel.Filter).To(Equal(firewallRuleUpdateInputFilterModel))
				Expect(updateFirewallRuleOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateFirewllRulesOptions successfully`, func() {
				// Construct an instance of the FirewallRulesUpdateInputItemFilter model
				firewallRulesUpdateInputItemFilterModel := new(firewallrulesv1.FirewallRulesUpdateInputItemFilter)
				Expect(firewallRulesUpdateInputItemFilterModel).ToNot(BeNil())
				firewallRulesUpdateInputItemFilterModel.ID = core.StringPtr("6f58318e7fa2477a23112e8118c66f61")
				Expect(firewallRulesUpdateInputItemFilterModel.ID).To(Equal(core.StringPtr("6f58318e7fa2477a23112e8118c66f61")))

				// Construct an instance of the FirewallRulesUpdateInputItem model
				firewallRulesUpdateInputItemModel := new(firewallrulesv1.FirewallRulesUpdateInputItem)
				Expect(firewallRulesUpdateInputItemModel).ToNot(BeNil())
				firewallRulesUpdateInputItemModel.ID = core.StringPtr("52161eb6af4241bb9d4b32394be72fdf")
				firewallRulesUpdateInputItemModel.Action = core.StringPtr("js_challenge")
				firewallRulesUpdateInputItemModel.Paused = core.BoolPtr(false)
				firewallRulesUpdateInputItemModel.Description = core.StringPtr("JS challenge site")
				firewallRulesUpdateInputItemModel.Filter = firewallRulesUpdateInputItemFilterModel
				Expect(firewallRulesUpdateInputItemModel.ID).To(Equal(core.StringPtr("52161eb6af4241bb9d4b32394be72fdf")))
				Expect(firewallRulesUpdateInputItemModel.Action).To(Equal(core.StringPtr("js_challenge")))
				Expect(firewallRulesUpdateInputItemModel.Paused).To(Equal(core.BoolPtr(false)))
				Expect(firewallRulesUpdateInputItemModel.Description).To(Equal(core.StringPtr("JS challenge site")))
				Expect(firewallRulesUpdateInputItemModel.Filter).To(Equal(firewallRulesUpdateInputItemFilterModel))

				// Construct an instance of the UpdateFirewllRulesOptions model
				xAuthUserToken := "testString"
				crn := "testString"
				zoneIdentifier := "testString"
				updateFirewllRulesOptionsModel := firewallRulesService.NewUpdateFirewllRulesOptions(xAuthUserToken, crn, zoneIdentifier)
				updateFirewllRulesOptionsModel.SetXAuthUserToken("testString")
				updateFirewllRulesOptionsModel.SetCrn("testString")
				updateFirewllRulesOptionsModel.SetZoneIdentifier("testString")
				updateFirewllRulesOptionsModel.SetAccept("testString")
				updateFirewllRulesOptionsModel.SetFirewallRulesUpdateInputItem([]firewallrulesv1.FirewallRulesUpdateInputItem{*firewallRulesUpdateInputItemModel})
				updateFirewllRulesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateFirewllRulesOptionsModel).ToNot(BeNil())
				Expect(updateFirewllRulesOptionsModel.XAuthUserToken).To(Equal(core.StringPtr("testString")))
				Expect(updateFirewllRulesOptionsModel.Crn).To(Equal(core.StringPtr("testString")))
				Expect(updateFirewllRulesOptionsModel.ZoneIdentifier).To(Equal(core.StringPtr("testString")))
				Expect(updateFirewllRulesOptionsModel.Accept).To(Equal(core.StringPtr("testString")))
				Expect(updateFirewllRulesOptionsModel.FirewallRulesUpdateInputItem).To(Equal([]firewallrulesv1.FirewallRulesUpdateInputItem{*firewallRulesUpdateInputItemModel}))
				Expect(updateFirewllRulesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewFirewallRuleInputWithFilterID successfully`, func() {
				var filter *firewallrulesv1.FirewallRuleInputWithFilterIdFilter = nil
				action := "js_challenge"
				_, err := firewallRulesService.NewFirewallRuleInputWithFilterID(filter, action)
				Expect(err).ToNot(BeNil())
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
