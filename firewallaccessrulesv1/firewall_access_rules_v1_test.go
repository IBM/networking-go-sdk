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

package firewallaccessrulesv1_test

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
	"github.com/IBM/networking-go-sdk/firewallaccessrulesv1"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe(`FirewallAccessRulesV1`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		crn := "testString"
		It(`Instantiate service client`, func() {
			firewallAccessRulesService, serviceErr := firewallaccessrulesv1.NewFirewallAccessRulesV1(&firewallaccessrulesv1.FirewallAccessRulesV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
				Crn:           core.StringPtr(crn),
			})
			Expect(firewallAccessRulesService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			firewallAccessRulesService, serviceErr := firewallaccessrulesv1.NewFirewallAccessRulesV1(&firewallaccessrulesv1.FirewallAccessRulesV1Options{
				URL: "{BAD_URL_STRING",
				Crn: core.StringPtr(crn),
			})
			Expect(firewallAccessRulesService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			firewallAccessRulesService, serviceErr := firewallaccessrulesv1.NewFirewallAccessRulesV1(&firewallaccessrulesv1.FirewallAccessRulesV1Options{
				URL: "https://firewallaccessrulesv1/api",
				Crn: core.StringPtr(crn),
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(firewallAccessRulesService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Validation Error`, func() {
			firewallAccessRulesService, serviceErr := firewallaccessrulesv1.NewFirewallAccessRulesV1(&firewallaccessrulesv1.FirewallAccessRulesV1Options{})
			Expect(firewallAccessRulesService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		crn := "testString"
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"FIREWALL_ACCESS_RULES_URL":       "https://firewallaccessrulesv1/api",
				"FIREWALL_ACCESS_RULES_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				firewallAccessRulesService, serviceErr := firewallaccessrulesv1.NewFirewallAccessRulesV1UsingExternalConfig(&firewallaccessrulesv1.FirewallAccessRulesV1Options{
					Crn: core.StringPtr(crn),
				})
				Expect(firewallAccessRulesService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := firewallAccessRulesService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != firewallAccessRulesService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(firewallAccessRulesService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(firewallAccessRulesService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				firewallAccessRulesService, serviceErr := firewallaccessrulesv1.NewFirewallAccessRulesV1UsingExternalConfig(&firewallaccessrulesv1.FirewallAccessRulesV1Options{
					URL: "https://testService/api",
					Crn: core.StringPtr(crn),
				})
				Expect(firewallAccessRulesService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(firewallAccessRulesService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := firewallAccessRulesService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != firewallAccessRulesService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(firewallAccessRulesService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(firewallAccessRulesService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				firewallAccessRulesService, serviceErr := firewallaccessrulesv1.NewFirewallAccessRulesV1UsingExternalConfig(&firewallaccessrulesv1.FirewallAccessRulesV1Options{
					Crn: core.StringPtr(crn),
				})
				err := firewallAccessRulesService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(firewallAccessRulesService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(firewallAccessRulesService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := firewallAccessRulesService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != firewallAccessRulesService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(firewallAccessRulesService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(firewallAccessRulesService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"FIREWALL_ACCESS_RULES_URL":       "https://firewallaccessrulesv1/api",
				"FIREWALL_ACCESS_RULES_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			firewallAccessRulesService, serviceErr := firewallaccessrulesv1.NewFirewallAccessRulesV1UsingExternalConfig(&firewallaccessrulesv1.FirewallAccessRulesV1Options{
				Crn: core.StringPtr(crn),
			})

			It(`Instantiate service client with error`, func() {
				Expect(firewallAccessRulesService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"FIREWALL_ACCESS_RULES_AUTH_TYPE": "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			firewallAccessRulesService, serviceErr := firewallaccessrulesv1.NewFirewallAccessRulesV1UsingExternalConfig(&firewallaccessrulesv1.FirewallAccessRulesV1Options{
				URL: "{BAD_URL_STRING",
				Crn: core.StringPtr(crn),
			})

			It(`Instantiate service client with error`, func() {
				Expect(firewallAccessRulesService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = firewallaccessrulesv1.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`ListAllAccountAccessRules(listAllAccountAccessRulesOptions *ListAllAccountAccessRulesOptions) - Operation response error`, func() {
		crn := "testString"
		listAllAccountAccessRulesPath := "/v1/testString/firewall/access_rules/rules"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listAllAccountAccessRulesPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["notes"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["mode"]).To(Equal([]string{"block"}))

					Expect(req.URL.Query()["configuration.target"]).To(Equal([]string{"ip"}))

					Expect(req.URL.Query()["configuration.value"]).To(Equal([]string{"1.2.3.4"}))

					Expect(req.URL.Query()["page"]).To(Equal([]string{fmt.Sprint(int64(38))}))

					Expect(req.URL.Query()["per_page"]).To(Equal([]string{fmt.Sprint(int64(5))}))

					Expect(req.URL.Query()["order"]).To(Equal([]string{"target"}))

					Expect(req.URL.Query()["direction"]).To(Equal([]string{"asc"}))

					Expect(req.URL.Query()["match"]).To(Equal([]string{"any"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListAllAccountAccessRules with error: Operation response processing error`, func() {
				firewallAccessRulesService, serviceErr := firewallaccessrulesv1.NewFirewallAccessRulesV1(&firewallaccessrulesv1.FirewallAccessRulesV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(firewallAccessRulesService).ToNot(BeNil())

				// Construct an instance of the ListAllAccountAccessRulesOptions model
				listAllAccountAccessRulesOptionsModel := new(firewallaccessrulesv1.ListAllAccountAccessRulesOptions)
				listAllAccountAccessRulesOptionsModel.Notes = core.StringPtr("testString")
				listAllAccountAccessRulesOptionsModel.Mode = core.StringPtr("block")
				listAllAccountAccessRulesOptionsModel.ConfigurationTarget = core.StringPtr("ip")
				listAllAccountAccessRulesOptionsModel.ConfigurationValue = core.StringPtr("1.2.3.4")
				listAllAccountAccessRulesOptionsModel.Page = core.Int64Ptr(int64(38))
				listAllAccountAccessRulesOptionsModel.PerPage = core.Int64Ptr(int64(5))
				listAllAccountAccessRulesOptionsModel.Order = core.StringPtr("target")
				listAllAccountAccessRulesOptionsModel.Direction = core.StringPtr("asc")
				listAllAccountAccessRulesOptionsModel.Match = core.StringPtr("any")
				listAllAccountAccessRulesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := firewallAccessRulesService.ListAllAccountAccessRules(listAllAccountAccessRulesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				firewallAccessRulesService.EnableRetries(0, 0)
				result, response, operationErr = firewallAccessRulesService.ListAllAccountAccessRules(listAllAccountAccessRulesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ListAllAccountAccessRules(listAllAccountAccessRulesOptions *ListAllAccountAccessRulesOptions)`, func() {
		crn := "testString"
		listAllAccountAccessRulesPath := "/v1/testString/firewall/access_rules/rules"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listAllAccountAccessRulesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["notes"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["mode"]).To(Equal([]string{"block"}))

					Expect(req.URL.Query()["configuration.target"]).To(Equal([]string{"ip"}))

					Expect(req.URL.Query()["configuration.value"]).To(Equal([]string{"1.2.3.4"}))

					Expect(req.URL.Query()["page"]).To(Equal([]string{fmt.Sprint(int64(38))}))

					Expect(req.URL.Query()["per_page"]).To(Equal([]string{fmt.Sprint(int64(5))}))

					Expect(req.URL.Query()["order"]).To(Equal([]string{"target"}))

					Expect(req.URL.Query()["direction"]).To(Equal([]string{"asc"}))

					Expect(req.URL.Query()["match"]).To(Equal([]string{"any"}))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": [{"id": "92f17202ed8bd63d69a66b86a49a8f6b", "notes": "This rule is set because of an event that occurred and caused X.", "allowed_modes": ["block"], "mode": "block", "scope": {"type": "account"}, "created_on": "2019-01-01T12:00:00", "modified_on": "2019-01-01T12:00:00", "configuration": {"target": "ip", "value": "ip example 198.51.100.4; ip_range example 198.51.100.4/16 ; asn example AS12345; country example AZ"}}], "result_info": {"page": 1, "per_page": 2, "count": 1, "total_count": 200}}`)
				}))
			})
			It(`Invoke ListAllAccountAccessRules successfully`, func() {
				firewallAccessRulesService, serviceErr := firewallaccessrulesv1.NewFirewallAccessRulesV1(&firewallaccessrulesv1.FirewallAccessRulesV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(firewallAccessRulesService).ToNot(BeNil())
				firewallAccessRulesService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := firewallAccessRulesService.ListAllAccountAccessRules(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListAllAccountAccessRulesOptions model
				listAllAccountAccessRulesOptionsModel := new(firewallaccessrulesv1.ListAllAccountAccessRulesOptions)
				listAllAccountAccessRulesOptionsModel.Notes = core.StringPtr("testString")
				listAllAccountAccessRulesOptionsModel.Mode = core.StringPtr("block")
				listAllAccountAccessRulesOptionsModel.ConfigurationTarget = core.StringPtr("ip")
				listAllAccountAccessRulesOptionsModel.ConfigurationValue = core.StringPtr("1.2.3.4")
				listAllAccountAccessRulesOptionsModel.Page = core.Int64Ptr(int64(38))
				listAllAccountAccessRulesOptionsModel.PerPage = core.Int64Ptr(int64(5))
				listAllAccountAccessRulesOptionsModel.Order = core.StringPtr("target")
				listAllAccountAccessRulesOptionsModel.Direction = core.StringPtr("asc")
				listAllAccountAccessRulesOptionsModel.Match = core.StringPtr("any")
				listAllAccountAccessRulesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = firewallAccessRulesService.ListAllAccountAccessRules(listAllAccountAccessRulesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = firewallAccessRulesService.ListAllAccountAccessRulesWithContext(ctx, listAllAccountAccessRulesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				firewallAccessRulesService.DisableRetries()
				result, response, operationErr = firewallAccessRulesService.ListAllAccountAccessRules(listAllAccountAccessRulesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = firewallAccessRulesService.ListAllAccountAccessRulesWithContext(ctx, listAllAccountAccessRulesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke ListAllAccountAccessRules with error: Operation request error`, func() {
				firewallAccessRulesService, serviceErr := firewallaccessrulesv1.NewFirewallAccessRulesV1(&firewallaccessrulesv1.FirewallAccessRulesV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(firewallAccessRulesService).ToNot(BeNil())

				// Construct an instance of the ListAllAccountAccessRulesOptions model
				listAllAccountAccessRulesOptionsModel := new(firewallaccessrulesv1.ListAllAccountAccessRulesOptions)
				listAllAccountAccessRulesOptionsModel.Notes = core.StringPtr("testString")
				listAllAccountAccessRulesOptionsModel.Mode = core.StringPtr("block")
				listAllAccountAccessRulesOptionsModel.ConfigurationTarget = core.StringPtr("ip")
				listAllAccountAccessRulesOptionsModel.ConfigurationValue = core.StringPtr("1.2.3.4")
				listAllAccountAccessRulesOptionsModel.Page = core.Int64Ptr(int64(38))
				listAllAccountAccessRulesOptionsModel.PerPage = core.Int64Ptr(int64(5))
				listAllAccountAccessRulesOptionsModel.Order = core.StringPtr("target")
				listAllAccountAccessRulesOptionsModel.Direction = core.StringPtr("asc")
				listAllAccountAccessRulesOptionsModel.Match = core.StringPtr("any")
				listAllAccountAccessRulesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := firewallAccessRulesService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := firewallAccessRulesService.ListAllAccountAccessRules(listAllAccountAccessRulesOptionsModel)
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
	Describe(`CreateAccountAccessRule(createAccountAccessRuleOptions *CreateAccountAccessRuleOptions) - Operation response error`, func() {
		crn := "testString"
		createAccountAccessRulePath := "/v1/testString/firewall/access_rules/rules"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createAccountAccessRulePath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateAccountAccessRule with error: Operation response processing error`, func() {
				firewallAccessRulesService, serviceErr := firewallaccessrulesv1.NewFirewallAccessRulesV1(&firewallaccessrulesv1.FirewallAccessRulesV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(firewallAccessRulesService).ToNot(BeNil())

				// Construct an instance of the AccountAccessRuleInputConfiguration model
				accountAccessRuleInputConfigurationModel := new(firewallaccessrulesv1.AccountAccessRuleInputConfiguration)
				accountAccessRuleInputConfigurationModel.Target = core.StringPtr("ip")
				accountAccessRuleInputConfigurationModel.Value = core.StringPtr("ip example 198.51.100.4; ip_range example 198.51.100.4/16 ; asn example AS12345; country example AZ")

				// Construct an instance of the CreateAccountAccessRuleOptions model
				createAccountAccessRuleOptionsModel := new(firewallaccessrulesv1.CreateAccountAccessRuleOptions)
				createAccountAccessRuleOptionsModel.Mode = core.StringPtr("block")
				createAccountAccessRuleOptionsModel.Notes = core.StringPtr("This rule is added because of event X that occurred on date xyz")
				createAccountAccessRuleOptionsModel.Configuration = accountAccessRuleInputConfigurationModel
				createAccountAccessRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := firewallAccessRulesService.CreateAccountAccessRule(createAccountAccessRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				firewallAccessRulesService.EnableRetries(0, 0)
				result, response, operationErr = firewallAccessRulesService.CreateAccountAccessRule(createAccountAccessRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`CreateAccountAccessRule(createAccountAccessRuleOptions *CreateAccountAccessRuleOptions)`, func() {
		crn := "testString"
		createAccountAccessRulePath := "/v1/testString/firewall/access_rules/rules"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createAccountAccessRulePath))
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
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": "92f17202ed8bd63d69a66b86a49a8f6b", "notes": "This rule is set because of an event that occurred and caused X.", "allowed_modes": ["block"], "mode": "block", "scope": {"type": "account"}, "created_on": "2019-01-01T12:00:00", "modified_on": "2019-01-01T12:00:00", "configuration": {"target": "ip", "value": "ip example 198.51.100.4; ip_range example 198.51.100.4/16 ; asn example AS12345; country example AZ"}}}`)
				}))
			})
			It(`Invoke CreateAccountAccessRule successfully`, func() {
				firewallAccessRulesService, serviceErr := firewallaccessrulesv1.NewFirewallAccessRulesV1(&firewallaccessrulesv1.FirewallAccessRulesV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(firewallAccessRulesService).ToNot(BeNil())
				firewallAccessRulesService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := firewallAccessRulesService.CreateAccountAccessRule(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the AccountAccessRuleInputConfiguration model
				accountAccessRuleInputConfigurationModel := new(firewallaccessrulesv1.AccountAccessRuleInputConfiguration)
				accountAccessRuleInputConfigurationModel.Target = core.StringPtr("ip")
				accountAccessRuleInputConfigurationModel.Value = core.StringPtr("ip example 198.51.100.4; ip_range example 198.51.100.4/16 ; asn example AS12345; country example AZ")

				// Construct an instance of the CreateAccountAccessRuleOptions model
				createAccountAccessRuleOptionsModel := new(firewallaccessrulesv1.CreateAccountAccessRuleOptions)
				createAccountAccessRuleOptionsModel.Mode = core.StringPtr("block")
				createAccountAccessRuleOptionsModel.Notes = core.StringPtr("This rule is added because of event X that occurred on date xyz")
				createAccountAccessRuleOptionsModel.Configuration = accountAccessRuleInputConfigurationModel
				createAccountAccessRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = firewallAccessRulesService.CreateAccountAccessRule(createAccountAccessRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = firewallAccessRulesService.CreateAccountAccessRuleWithContext(ctx, createAccountAccessRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				firewallAccessRulesService.DisableRetries()
				result, response, operationErr = firewallAccessRulesService.CreateAccountAccessRule(createAccountAccessRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = firewallAccessRulesService.CreateAccountAccessRuleWithContext(ctx, createAccountAccessRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke CreateAccountAccessRule with error: Operation request error`, func() {
				firewallAccessRulesService, serviceErr := firewallaccessrulesv1.NewFirewallAccessRulesV1(&firewallaccessrulesv1.FirewallAccessRulesV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(firewallAccessRulesService).ToNot(BeNil())

				// Construct an instance of the AccountAccessRuleInputConfiguration model
				accountAccessRuleInputConfigurationModel := new(firewallaccessrulesv1.AccountAccessRuleInputConfiguration)
				accountAccessRuleInputConfigurationModel.Target = core.StringPtr("ip")
				accountAccessRuleInputConfigurationModel.Value = core.StringPtr("ip example 198.51.100.4; ip_range example 198.51.100.4/16 ; asn example AS12345; country example AZ")

				// Construct an instance of the CreateAccountAccessRuleOptions model
				createAccountAccessRuleOptionsModel := new(firewallaccessrulesv1.CreateAccountAccessRuleOptions)
				createAccountAccessRuleOptionsModel.Mode = core.StringPtr("block")
				createAccountAccessRuleOptionsModel.Notes = core.StringPtr("This rule is added because of event X that occurred on date xyz")
				createAccountAccessRuleOptionsModel.Configuration = accountAccessRuleInputConfigurationModel
				createAccountAccessRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := firewallAccessRulesService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := firewallAccessRulesService.CreateAccountAccessRule(createAccountAccessRuleOptionsModel)
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
	Describe(`DeleteAccountAccessRule(deleteAccountAccessRuleOptions *DeleteAccountAccessRuleOptions) - Operation response error`, func() {
		crn := "testString"
		deleteAccountAccessRulePath := "/v1/testString/firewall/access_rules/rules/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteAccountAccessRulePath))
					Expect(req.Method).To(Equal("DELETE"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke DeleteAccountAccessRule with error: Operation response processing error`, func() {
				firewallAccessRulesService, serviceErr := firewallaccessrulesv1.NewFirewallAccessRulesV1(&firewallaccessrulesv1.FirewallAccessRulesV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(firewallAccessRulesService).ToNot(BeNil())

				// Construct an instance of the DeleteAccountAccessRuleOptions model
				deleteAccountAccessRuleOptionsModel := new(firewallaccessrulesv1.DeleteAccountAccessRuleOptions)
				deleteAccountAccessRuleOptionsModel.AccessruleIdentifier = core.StringPtr("testString")
				deleteAccountAccessRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := firewallAccessRulesService.DeleteAccountAccessRule(deleteAccountAccessRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				firewallAccessRulesService.EnableRetries(0, 0)
				result, response, operationErr = firewallAccessRulesService.DeleteAccountAccessRule(deleteAccountAccessRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`DeleteAccountAccessRule(deleteAccountAccessRuleOptions *DeleteAccountAccessRuleOptions)`, func() {
		crn := "testString"
		deleteAccountAccessRulePath := "/v1/testString/firewall/access_rules/rules/testString"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteAccountAccessRulePath))
					Expect(req.Method).To(Equal("DELETE"))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": "f1aba936b94213e5b8dca0c0dbf1f9cc"}}`)
				}))
			})
			It(`Invoke DeleteAccountAccessRule successfully`, func() {
				firewallAccessRulesService, serviceErr := firewallaccessrulesv1.NewFirewallAccessRulesV1(&firewallaccessrulesv1.FirewallAccessRulesV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(firewallAccessRulesService).ToNot(BeNil())
				firewallAccessRulesService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := firewallAccessRulesService.DeleteAccountAccessRule(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DeleteAccountAccessRuleOptions model
				deleteAccountAccessRuleOptionsModel := new(firewallaccessrulesv1.DeleteAccountAccessRuleOptions)
				deleteAccountAccessRuleOptionsModel.AccessruleIdentifier = core.StringPtr("testString")
				deleteAccountAccessRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = firewallAccessRulesService.DeleteAccountAccessRule(deleteAccountAccessRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = firewallAccessRulesService.DeleteAccountAccessRuleWithContext(ctx, deleteAccountAccessRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				firewallAccessRulesService.DisableRetries()
				result, response, operationErr = firewallAccessRulesService.DeleteAccountAccessRule(deleteAccountAccessRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = firewallAccessRulesService.DeleteAccountAccessRuleWithContext(ctx, deleteAccountAccessRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke DeleteAccountAccessRule with error: Operation validation and request error`, func() {
				firewallAccessRulesService, serviceErr := firewallaccessrulesv1.NewFirewallAccessRulesV1(&firewallaccessrulesv1.FirewallAccessRulesV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(firewallAccessRulesService).ToNot(BeNil())

				// Construct an instance of the DeleteAccountAccessRuleOptions model
				deleteAccountAccessRuleOptionsModel := new(firewallaccessrulesv1.DeleteAccountAccessRuleOptions)
				deleteAccountAccessRuleOptionsModel.AccessruleIdentifier = core.StringPtr("testString")
				deleteAccountAccessRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := firewallAccessRulesService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := firewallAccessRulesService.DeleteAccountAccessRule(deleteAccountAccessRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the DeleteAccountAccessRuleOptions model with no property values
				deleteAccountAccessRuleOptionsModelNew := new(firewallaccessrulesv1.DeleteAccountAccessRuleOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = firewallAccessRulesService.DeleteAccountAccessRule(deleteAccountAccessRuleOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetAccountAccessRule(getAccountAccessRuleOptions *GetAccountAccessRuleOptions) - Operation response error`, func() {
		crn := "testString"
		getAccountAccessRulePath := "/v1/testString/firewall/access_rules/rules/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getAccountAccessRulePath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetAccountAccessRule with error: Operation response processing error`, func() {
				firewallAccessRulesService, serviceErr := firewallaccessrulesv1.NewFirewallAccessRulesV1(&firewallaccessrulesv1.FirewallAccessRulesV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(firewallAccessRulesService).ToNot(BeNil())

				// Construct an instance of the GetAccountAccessRuleOptions model
				getAccountAccessRuleOptionsModel := new(firewallaccessrulesv1.GetAccountAccessRuleOptions)
				getAccountAccessRuleOptionsModel.AccessruleIdentifier = core.StringPtr("testString")
				getAccountAccessRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := firewallAccessRulesService.GetAccountAccessRule(getAccountAccessRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				firewallAccessRulesService.EnableRetries(0, 0)
				result, response, operationErr = firewallAccessRulesService.GetAccountAccessRule(getAccountAccessRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetAccountAccessRule(getAccountAccessRuleOptions *GetAccountAccessRuleOptions)`, func() {
		crn := "testString"
		getAccountAccessRulePath := "/v1/testString/firewall/access_rules/rules/testString"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getAccountAccessRulePath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": "92f17202ed8bd63d69a66b86a49a8f6b", "notes": "This rule is set because of an event that occurred and caused X.", "allowed_modes": ["block"], "mode": "block", "scope": {"type": "account"}, "created_on": "2019-01-01T12:00:00", "modified_on": "2019-01-01T12:00:00", "configuration": {"target": "ip", "value": "ip example 198.51.100.4; ip_range example 198.51.100.4/16 ; asn example AS12345; country example AZ"}}}`)
				}))
			})
			It(`Invoke GetAccountAccessRule successfully`, func() {
				firewallAccessRulesService, serviceErr := firewallaccessrulesv1.NewFirewallAccessRulesV1(&firewallaccessrulesv1.FirewallAccessRulesV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(firewallAccessRulesService).ToNot(BeNil())
				firewallAccessRulesService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := firewallAccessRulesService.GetAccountAccessRule(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetAccountAccessRuleOptions model
				getAccountAccessRuleOptionsModel := new(firewallaccessrulesv1.GetAccountAccessRuleOptions)
				getAccountAccessRuleOptionsModel.AccessruleIdentifier = core.StringPtr("testString")
				getAccountAccessRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = firewallAccessRulesService.GetAccountAccessRule(getAccountAccessRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = firewallAccessRulesService.GetAccountAccessRuleWithContext(ctx, getAccountAccessRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				firewallAccessRulesService.DisableRetries()
				result, response, operationErr = firewallAccessRulesService.GetAccountAccessRule(getAccountAccessRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = firewallAccessRulesService.GetAccountAccessRuleWithContext(ctx, getAccountAccessRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke GetAccountAccessRule with error: Operation validation and request error`, func() {
				firewallAccessRulesService, serviceErr := firewallaccessrulesv1.NewFirewallAccessRulesV1(&firewallaccessrulesv1.FirewallAccessRulesV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(firewallAccessRulesService).ToNot(BeNil())

				// Construct an instance of the GetAccountAccessRuleOptions model
				getAccountAccessRuleOptionsModel := new(firewallaccessrulesv1.GetAccountAccessRuleOptions)
				getAccountAccessRuleOptionsModel.AccessruleIdentifier = core.StringPtr("testString")
				getAccountAccessRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := firewallAccessRulesService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := firewallAccessRulesService.GetAccountAccessRule(getAccountAccessRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetAccountAccessRuleOptions model with no property values
				getAccountAccessRuleOptionsModelNew := new(firewallaccessrulesv1.GetAccountAccessRuleOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = firewallAccessRulesService.GetAccountAccessRule(getAccountAccessRuleOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateAccountAccessRule(updateAccountAccessRuleOptions *UpdateAccountAccessRuleOptions) - Operation response error`, func() {
		crn := "testString"
		updateAccountAccessRulePath := "/v1/testString/firewall/access_rules/rules/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateAccountAccessRulePath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateAccountAccessRule with error: Operation response processing error`, func() {
				firewallAccessRulesService, serviceErr := firewallaccessrulesv1.NewFirewallAccessRulesV1(&firewallaccessrulesv1.FirewallAccessRulesV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(firewallAccessRulesService).ToNot(BeNil())

				// Construct an instance of the UpdateAccountAccessRuleOptions model
				updateAccountAccessRuleOptionsModel := new(firewallaccessrulesv1.UpdateAccountAccessRuleOptions)
				updateAccountAccessRuleOptionsModel.AccessruleIdentifier = core.StringPtr("testString")
				updateAccountAccessRuleOptionsModel.Mode = core.StringPtr("block")
				updateAccountAccessRuleOptionsModel.Notes = core.StringPtr("This rule is added because of event X that occurred on date xyz")
				updateAccountAccessRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := firewallAccessRulesService.UpdateAccountAccessRule(updateAccountAccessRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				firewallAccessRulesService.EnableRetries(0, 0)
				result, response, operationErr = firewallAccessRulesService.UpdateAccountAccessRule(updateAccountAccessRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UpdateAccountAccessRule(updateAccountAccessRuleOptions *UpdateAccountAccessRuleOptions)`, func() {
		crn := "testString"
		updateAccountAccessRulePath := "/v1/testString/firewall/access_rules/rules/testString"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateAccountAccessRulePath))
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
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": "92f17202ed8bd63d69a66b86a49a8f6b", "notes": "This rule is set because of an event that occurred and caused X.", "allowed_modes": ["block"], "mode": "block", "scope": {"type": "account"}, "created_on": "2019-01-01T12:00:00", "modified_on": "2019-01-01T12:00:00", "configuration": {"target": "ip", "value": "ip example 198.51.100.4; ip_range example 198.51.100.4/16 ; asn example AS12345; country example AZ"}}}`)
				}))
			})
			It(`Invoke UpdateAccountAccessRule successfully`, func() {
				firewallAccessRulesService, serviceErr := firewallaccessrulesv1.NewFirewallAccessRulesV1(&firewallaccessrulesv1.FirewallAccessRulesV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(firewallAccessRulesService).ToNot(BeNil())
				firewallAccessRulesService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := firewallAccessRulesService.UpdateAccountAccessRule(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateAccountAccessRuleOptions model
				updateAccountAccessRuleOptionsModel := new(firewallaccessrulesv1.UpdateAccountAccessRuleOptions)
				updateAccountAccessRuleOptionsModel.AccessruleIdentifier = core.StringPtr("testString")
				updateAccountAccessRuleOptionsModel.Mode = core.StringPtr("block")
				updateAccountAccessRuleOptionsModel.Notes = core.StringPtr("This rule is added because of event X that occurred on date xyz")
				updateAccountAccessRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = firewallAccessRulesService.UpdateAccountAccessRule(updateAccountAccessRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = firewallAccessRulesService.UpdateAccountAccessRuleWithContext(ctx, updateAccountAccessRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				firewallAccessRulesService.DisableRetries()
				result, response, operationErr = firewallAccessRulesService.UpdateAccountAccessRule(updateAccountAccessRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = firewallAccessRulesService.UpdateAccountAccessRuleWithContext(ctx, updateAccountAccessRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke UpdateAccountAccessRule with error: Operation validation and request error`, func() {
				firewallAccessRulesService, serviceErr := firewallaccessrulesv1.NewFirewallAccessRulesV1(&firewallaccessrulesv1.FirewallAccessRulesV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(firewallAccessRulesService).ToNot(BeNil())

				// Construct an instance of the UpdateAccountAccessRuleOptions model
				updateAccountAccessRuleOptionsModel := new(firewallaccessrulesv1.UpdateAccountAccessRuleOptions)
				updateAccountAccessRuleOptionsModel.AccessruleIdentifier = core.StringPtr("testString")
				updateAccountAccessRuleOptionsModel.Mode = core.StringPtr("block")
				updateAccountAccessRuleOptionsModel.Notes = core.StringPtr("This rule is added because of event X that occurred on date xyz")
				updateAccountAccessRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := firewallAccessRulesService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := firewallAccessRulesService.UpdateAccountAccessRule(updateAccountAccessRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateAccountAccessRuleOptions model with no property values
				updateAccountAccessRuleOptionsModelNew := new(firewallaccessrulesv1.UpdateAccountAccessRuleOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = firewallAccessRulesService.UpdateAccountAccessRule(updateAccountAccessRuleOptionsModelNew)
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
			firewallAccessRulesService, _ := firewallaccessrulesv1.NewFirewallAccessRulesV1(&firewallaccessrulesv1.FirewallAccessRulesV1Options{
				URL:           "http://firewallaccessrulesv1modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
				Crn:           core.StringPtr(crn),
			})
			It(`Invoke NewAccountAccessRuleInputConfiguration successfully`, func() {
				target := "ip"
				value := "ip example 198.51.100.4; ip_range example 198.51.100.4/16 ; asn example AS12345; country example AZ"
				model, err := firewallAccessRulesService.NewAccountAccessRuleInputConfiguration(target, value)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewCreateAccountAccessRuleOptions successfully`, func() {
				// Construct an instance of the AccountAccessRuleInputConfiguration model
				accountAccessRuleInputConfigurationModel := new(firewallaccessrulesv1.AccountAccessRuleInputConfiguration)
				Expect(accountAccessRuleInputConfigurationModel).ToNot(BeNil())
				accountAccessRuleInputConfigurationModel.Target = core.StringPtr("ip")
				accountAccessRuleInputConfigurationModel.Value = core.StringPtr("ip example 198.51.100.4; ip_range example 198.51.100.4/16 ; asn example AS12345; country example AZ")
				Expect(accountAccessRuleInputConfigurationModel.Target).To(Equal(core.StringPtr("ip")))
				Expect(accountAccessRuleInputConfigurationModel.Value).To(Equal(core.StringPtr("ip example 198.51.100.4; ip_range example 198.51.100.4/16 ; asn example AS12345; country example AZ")))

				// Construct an instance of the CreateAccountAccessRuleOptions model
				createAccountAccessRuleOptionsModel := firewallAccessRulesService.NewCreateAccountAccessRuleOptions()
				createAccountAccessRuleOptionsModel.SetMode("block")
				createAccountAccessRuleOptionsModel.SetNotes("This rule is added because of event X that occurred on date xyz")
				createAccountAccessRuleOptionsModel.SetConfiguration(accountAccessRuleInputConfigurationModel)
				createAccountAccessRuleOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createAccountAccessRuleOptionsModel).ToNot(BeNil())
				Expect(createAccountAccessRuleOptionsModel.Mode).To(Equal(core.StringPtr("block")))
				Expect(createAccountAccessRuleOptionsModel.Notes).To(Equal(core.StringPtr("This rule is added because of event X that occurred on date xyz")))
				Expect(createAccountAccessRuleOptionsModel.Configuration).To(Equal(accountAccessRuleInputConfigurationModel))
				Expect(createAccountAccessRuleOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteAccountAccessRuleOptions successfully`, func() {
				// Construct an instance of the DeleteAccountAccessRuleOptions model
				accessruleIdentifier := "testString"
				deleteAccountAccessRuleOptionsModel := firewallAccessRulesService.NewDeleteAccountAccessRuleOptions(accessruleIdentifier)
				deleteAccountAccessRuleOptionsModel.SetAccessruleIdentifier("testString")
				deleteAccountAccessRuleOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteAccountAccessRuleOptionsModel).ToNot(BeNil())
				Expect(deleteAccountAccessRuleOptionsModel.AccessruleIdentifier).To(Equal(core.StringPtr("testString")))
				Expect(deleteAccountAccessRuleOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetAccountAccessRuleOptions successfully`, func() {
				// Construct an instance of the GetAccountAccessRuleOptions model
				accessruleIdentifier := "testString"
				getAccountAccessRuleOptionsModel := firewallAccessRulesService.NewGetAccountAccessRuleOptions(accessruleIdentifier)
				getAccountAccessRuleOptionsModel.SetAccessruleIdentifier("testString")
				getAccountAccessRuleOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getAccountAccessRuleOptionsModel).ToNot(BeNil())
				Expect(getAccountAccessRuleOptionsModel.AccessruleIdentifier).To(Equal(core.StringPtr("testString")))
				Expect(getAccountAccessRuleOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListAllAccountAccessRulesOptions successfully`, func() {
				// Construct an instance of the ListAllAccountAccessRulesOptions model
				listAllAccountAccessRulesOptionsModel := firewallAccessRulesService.NewListAllAccountAccessRulesOptions()
				listAllAccountAccessRulesOptionsModel.SetNotes("testString")
				listAllAccountAccessRulesOptionsModel.SetMode("block")
				listAllAccountAccessRulesOptionsModel.SetConfigurationTarget("ip")
				listAllAccountAccessRulesOptionsModel.SetConfigurationValue("1.2.3.4")
				listAllAccountAccessRulesOptionsModel.SetPage(int64(38))
				listAllAccountAccessRulesOptionsModel.SetPerPage(int64(5))
				listAllAccountAccessRulesOptionsModel.SetOrder("target")
				listAllAccountAccessRulesOptionsModel.SetDirection("asc")
				listAllAccountAccessRulesOptionsModel.SetMatch("any")
				listAllAccountAccessRulesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listAllAccountAccessRulesOptionsModel).ToNot(BeNil())
				Expect(listAllAccountAccessRulesOptionsModel.Notes).To(Equal(core.StringPtr("testString")))
				Expect(listAllAccountAccessRulesOptionsModel.Mode).To(Equal(core.StringPtr("block")))
				Expect(listAllAccountAccessRulesOptionsModel.ConfigurationTarget).To(Equal(core.StringPtr("ip")))
				Expect(listAllAccountAccessRulesOptionsModel.ConfigurationValue).To(Equal(core.StringPtr("1.2.3.4")))
				Expect(listAllAccountAccessRulesOptionsModel.Page).To(Equal(core.Int64Ptr(int64(38))))
				Expect(listAllAccountAccessRulesOptionsModel.PerPage).To(Equal(core.Int64Ptr(int64(5))))
				Expect(listAllAccountAccessRulesOptionsModel.Order).To(Equal(core.StringPtr("target")))
				Expect(listAllAccountAccessRulesOptionsModel.Direction).To(Equal(core.StringPtr("asc")))
				Expect(listAllAccountAccessRulesOptionsModel.Match).To(Equal(core.StringPtr("any")))
				Expect(listAllAccountAccessRulesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateAccountAccessRuleOptions successfully`, func() {
				// Construct an instance of the UpdateAccountAccessRuleOptions model
				accessruleIdentifier := "testString"
				updateAccountAccessRuleOptionsModel := firewallAccessRulesService.NewUpdateAccountAccessRuleOptions(accessruleIdentifier)
				updateAccountAccessRuleOptionsModel.SetAccessruleIdentifier("testString")
				updateAccountAccessRuleOptionsModel.SetMode("block")
				updateAccountAccessRuleOptionsModel.SetNotes("This rule is added because of event X that occurred on date xyz")
				updateAccountAccessRuleOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateAccountAccessRuleOptionsModel).ToNot(BeNil())
				Expect(updateAccountAccessRuleOptionsModel.AccessruleIdentifier).To(Equal(core.StringPtr("testString")))
				Expect(updateAccountAccessRuleOptionsModel.Mode).To(Equal(core.StringPtr("block")))
				Expect(updateAccountAccessRuleOptionsModel.Notes).To(Equal(core.StringPtr("This rule is added because of event X that occurred on date xyz")))
				Expect(updateAccountAccessRuleOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
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
	return io.NopCloser(bytes.NewReader([]byte(mockData)))
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
