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

package firewallrulesv1_test

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
	"github.com/IBM/networking-go-sdk/firewallrulesv1"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
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
				"FIREWALL_RULES_URL":       "https://firewallrulesv1/api",
				"FIREWALL_RULES_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				firewallRulesService, serviceErr := firewallrulesv1.NewFirewallRulesV1UsingExternalConfig(&firewallrulesv1.FirewallRulesV1Options{})
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
				firewallRulesService, serviceErr := firewallrulesv1.NewFirewallRulesV1UsingExternalConfig(&firewallrulesv1.FirewallRulesV1Options{})
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
				"FIREWALL_RULES_URL":       "https://firewallrulesv1/api",
				"FIREWALL_RULES_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			firewallRulesService, serviceErr := firewallrulesv1.NewFirewallRulesV1UsingExternalConfig(&firewallrulesv1.FirewallRulesV1Options{})

			It(`Instantiate service client with error`, func() {
				Expect(firewallRulesService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"FIREWALL_RULES_AUTH_TYPE": "NOAuth",
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
	Describe(`ListAllFirewallRules(listAllFirewallRulesOptions *ListAllFirewallRulesOptions) - Operation response error`, func() {
		listAllFirewallRulesPath := "/v1/testString/zones/testString/firewall/rules"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listAllFirewallRulesPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["X-Auth-User-Token"]).ToNot(BeNil())
					Expect(req.Header["X-Auth-User-Token"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListAllFirewallRules with error: Operation response processing error`, func() {
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
				listAllFirewallRulesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := firewallRulesService.ListAllFirewallRules(listAllFirewallRulesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				firewallRulesService.EnableRetries(0, 0)
				result, response, operationErr = firewallRulesService.ListAllFirewallRules(listAllFirewallRulesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
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
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": [{"id": "52161eb6af4241bb9d4b32394be72fdf", "paused": false, "description": "JS challenge site", "action": "js_challenge", "filter": {"id": "6f58318e7fa2477a23112e8118c66f61", "paused": true, "description": "Login from office", "expression": "ip.src eq 93.184.216.0 and (http.request.uri.path ~ \"^.*/wp-login.php$\" or http.request.uri.path ~ \"^.*/xmlrpc.php$\")"}, "created_on": "2019-01-01T05:20:00.123Z", "modified_on": "2019-01-01T05:20:00.123Z"}], "result_info": {"page": 1, "per_page": 2, "count": 1, "total_count": 200}}`)
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
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": [{"id": "52161eb6af4241bb9d4b32394be72fdf", "paused": false, "description": "JS challenge site", "action": "js_challenge", "filter": {"id": "6f58318e7fa2477a23112e8118c66f61", "paused": true, "description": "Login from office", "expression": "ip.src eq 93.184.216.0 and (http.request.uri.path ~ \"^.*/wp-login.php$\" or http.request.uri.path ~ \"^.*/xmlrpc.php$\")"}, "created_on": "2019-01-01T05:20:00.123Z", "modified_on": "2019-01-01T05:20:00.123Z"}], "result_info": {"page": 1, "per_page": 2, "count": 1, "total_count": 200}}`)
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
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ListAllFirewallRules successfully`, func() {
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
				listAllFirewallRulesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := firewallRulesService.ListAllFirewallRules(listAllFirewallRulesOptionsModel)
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
	Describe(`CreateFirewallRules(createFirewallRulesOptions *CreateFirewallRulesOptions) - Operation response error`, func() {
		createFirewallRulesPath := "/v1/testString/zones/testString/firewall/rules"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createFirewallRulesPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["X-Auth-User-Token"]).ToNot(BeNil())
					Expect(req.Header["X-Auth-User-Token"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateFirewallRules with error: Operation response processing error`, func() {
				firewallRulesService, serviceErr := firewallrulesv1.NewFirewallRulesV1(&firewallrulesv1.FirewallRulesV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(firewallRulesService).ToNot(BeNil())

				// Construct an instance of the FirewallRuleInputFilterID model
				firewallRuleInputFilterModel := new(firewallrulesv1.FirewallRuleInputFilterID)
				firewallRuleInputFilterModel.ID = core.StringPtr("6f58318e7fa2477a23112e8118c66f61")

				// Construct an instance of the FirewallRuleInput model
				firewallRuleInputModel := new(firewallrulesv1.FirewallRuleInput)
				firewallRuleInputModel.Filter = firewallRuleInputFilterModel
				firewallRuleInputModel.Action = core.StringPtr("js_challenge")
				firewallRuleInputModel.Description = core.StringPtr("JS challenge site")
				firewallRuleInputModel.Paused = core.BoolPtr(false)
				firewallRuleInputModel.Priority = core.Int64Ptr(int64(1))

				// Construct an instance of the CreateFirewallRulesOptions model
				createFirewallRulesOptionsModel := new(firewallrulesv1.CreateFirewallRulesOptions)
				createFirewallRulesOptionsModel.XAuthUserToken = core.StringPtr("testString")
				createFirewallRulesOptionsModel.Crn = core.StringPtr("testString")
				createFirewallRulesOptionsModel.ZoneIdentifier = core.StringPtr("testString")
				createFirewallRulesOptionsModel.FirewallRuleInput = []firewallrulesv1.FirewallRuleInput{*firewallRuleInputModel}
				createFirewallRulesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := firewallRulesService.CreateFirewallRules(createFirewallRulesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				firewallRulesService.EnableRetries(0, 0)
				result, response, operationErr = firewallRulesService.CreateFirewallRules(createFirewallRulesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
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
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": [{"id": "52161eb6af4241bb9d4b32394be72fdf", "paused": false, "description": "JS challenge site", "action": "js_challenge", "filter": {"id": "6f58318e7fa2477a23112e8118c66f61", "paused": true, "description": "Login from office", "expression": "ip.src eq 93.184.216.0 and (http.request.uri.path ~ \"^.*/wp-login.php$\" or http.request.uri.path ~ \"^.*/xmlrpc.php$\")"}, "created_on": "2019-01-01T05:20:00.123Z", "modified_on": "2019-01-01T05:20:00.123Z"}]}`)
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

				// Construct an instance of the FirewallRuleInputFilterID model
				firewallRuleInputFilterModel := new(firewallrulesv1.FirewallRuleInputFilterID)
				firewallRuleInputFilterModel.ID = core.StringPtr("6f58318e7fa2477a23112e8118c66f61")

				// Construct an instance of the FirewallRuleInput model
				firewallRuleInputModel := new(firewallrulesv1.FirewallRuleInput)
				firewallRuleInputModel.Filter = firewallRuleInputFilterModel
				firewallRuleInputModel.Action = core.StringPtr("js_challenge")
				firewallRuleInputModel.Description = core.StringPtr("JS challenge site")
				firewallRuleInputModel.Paused = core.BoolPtr(false)
				firewallRuleInputModel.Priority = core.Int64Ptr(int64(1))

				// Construct an instance of the CreateFirewallRulesOptions model
				createFirewallRulesOptionsModel := new(firewallrulesv1.CreateFirewallRulesOptions)
				createFirewallRulesOptionsModel.XAuthUserToken = core.StringPtr("testString")
				createFirewallRulesOptionsModel.Crn = core.StringPtr("testString")
				createFirewallRulesOptionsModel.ZoneIdentifier = core.StringPtr("testString")
				createFirewallRulesOptionsModel.FirewallRuleInput = []firewallrulesv1.FirewallRuleInput{*firewallRuleInputModel}
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
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": [{"id": "52161eb6af4241bb9d4b32394be72fdf", "paused": false, "description": "JS challenge site", "action": "js_challenge", "filter": {"id": "6f58318e7fa2477a23112e8118c66f61", "paused": true, "description": "Login from office", "expression": "ip.src eq 93.184.216.0 and (http.request.uri.path ~ \"^.*/wp-login.php$\" or http.request.uri.path ~ \"^.*/xmlrpc.php$\")"}, "created_on": "2019-01-01T05:20:00.123Z", "modified_on": "2019-01-01T05:20:00.123Z"}]}`)
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

				// Construct an instance of the FirewallRuleInputFilterID model
				firewallRuleInputFilterModel := new(firewallrulesv1.FirewallRuleInputFilterID)
				firewallRuleInputFilterModel.ID = core.StringPtr("6f58318e7fa2477a23112e8118c66f61")

				// Construct an instance of the FirewallRuleInput model
				firewallRuleInputModel := new(firewallrulesv1.FirewallRuleInput)
				firewallRuleInputModel.Filter = firewallRuleInputFilterModel
				firewallRuleInputModel.Action = core.StringPtr("js_challenge")
				firewallRuleInputModel.Description = core.StringPtr("JS challenge site")
				firewallRuleInputModel.Paused = core.BoolPtr(false)
				firewallRuleInputModel.Priority = core.Int64Ptr(int64(1))

				// Construct an instance of the CreateFirewallRulesOptions model
				createFirewallRulesOptionsModel := new(firewallrulesv1.CreateFirewallRulesOptions)
				createFirewallRulesOptionsModel.XAuthUserToken = core.StringPtr("testString")
				createFirewallRulesOptionsModel.Crn = core.StringPtr("testString")
				createFirewallRulesOptionsModel.ZoneIdentifier = core.StringPtr("testString")
				createFirewallRulesOptionsModel.FirewallRuleInput = []firewallrulesv1.FirewallRuleInput{*firewallRuleInputModel}
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

				// Construct an instance of the FirewallRuleInputFilterID model
				firewallRuleInputFilterModel := new(firewallrulesv1.FirewallRuleInputFilterID)
				firewallRuleInputFilterModel.ID = core.StringPtr("6f58318e7fa2477a23112e8118c66f61")

				// Construct an instance of the FirewallRuleInput model
				firewallRuleInputModel := new(firewallrulesv1.FirewallRuleInput)
				firewallRuleInputModel.Filter = firewallRuleInputFilterModel
				firewallRuleInputModel.Action = core.StringPtr("js_challenge")
				firewallRuleInputModel.Description = core.StringPtr("JS challenge site")
				firewallRuleInputModel.Paused = core.BoolPtr(false)
				firewallRuleInputModel.Priority = core.Int64Ptr(int64(1))

				// Construct an instance of the CreateFirewallRulesOptions model
				createFirewallRulesOptionsModel := new(firewallrulesv1.CreateFirewallRulesOptions)
				createFirewallRulesOptionsModel.XAuthUserToken = core.StringPtr("testString")
				createFirewallRulesOptionsModel.Crn = core.StringPtr("testString")
				createFirewallRulesOptionsModel.ZoneIdentifier = core.StringPtr("testString")
				createFirewallRulesOptionsModel.FirewallRuleInput = []firewallrulesv1.FirewallRuleInput{*firewallRuleInputModel}
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
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke CreateFirewallRules successfully`, func() {
				firewallRulesService, serviceErr := firewallrulesv1.NewFirewallRulesV1(&firewallrulesv1.FirewallRulesV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(firewallRulesService).ToNot(BeNil())

				// Construct an instance of the FirewallRuleInputFilterID model
				firewallRuleInputFilterModel := new(firewallrulesv1.FirewallRuleInputFilterID)
				firewallRuleInputFilterModel.ID = core.StringPtr("6f58318e7fa2477a23112e8118c66f61")

				// Construct an instance of the FirewallRuleInput model
				firewallRuleInputModel := new(firewallrulesv1.FirewallRuleInput)
				firewallRuleInputModel.Filter = firewallRuleInputFilterModel
				firewallRuleInputModel.Action = core.StringPtr("js_challenge")
				firewallRuleInputModel.Description = core.StringPtr("JS challenge site")
				firewallRuleInputModel.Paused = core.BoolPtr(false)
				firewallRuleInputModel.Priority = core.Int64Ptr(int64(1))

				// Construct an instance of the CreateFirewallRulesOptions model
				createFirewallRulesOptionsModel := new(firewallrulesv1.CreateFirewallRulesOptions)
				createFirewallRulesOptionsModel.XAuthUserToken = core.StringPtr("testString")
				createFirewallRulesOptionsModel.Crn = core.StringPtr("testString")
				createFirewallRulesOptionsModel.ZoneIdentifier = core.StringPtr("testString")
				createFirewallRulesOptionsModel.FirewallRuleInput = []firewallrulesv1.FirewallRuleInput{*firewallRuleInputModel}
				createFirewallRulesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := firewallRulesService.CreateFirewallRules(createFirewallRulesOptionsModel)
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
	Describe(`UpdateFirewllRules(updateFirewllRulesOptions *UpdateFirewllRulesOptions) - Operation response error`, func() {
		updateFirewllRulesPath := "/v1/testString/zones/testString/firewall/rules"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateFirewllRulesPath))
					Expect(req.Method).To(Equal("PUT"))
					Expect(req.Header["X-Auth-User-Token"]).ToNot(BeNil())
					Expect(req.Header["X-Auth-User-Token"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateFirewllRules with error: Operation response processing error`, func() {
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
				firewallRulesUpdateInputItemModel.Priority = core.Int64Ptr(int64(1))
				firewallRulesUpdateInputItemModel.Description = core.StringPtr("JS challenge site")
				firewallRulesUpdateInputItemModel.Filter = firewallRulesUpdateInputItemFilterModel

				// Construct an instance of the UpdateFirewllRulesOptions model
				updateFirewllRulesOptionsModel := new(firewallrulesv1.UpdateFirewllRulesOptions)
				updateFirewllRulesOptionsModel.XAuthUserToken = core.StringPtr("testString")
				updateFirewllRulesOptionsModel.Crn = core.StringPtr("testString")
				updateFirewllRulesOptionsModel.ZoneIdentifier = core.StringPtr("testString")
				updateFirewllRulesOptionsModel.FirewallRulesUpdateInputItem = []firewallrulesv1.FirewallRulesUpdateInputItem{*firewallRulesUpdateInputItemModel}
				updateFirewllRulesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := firewallRulesService.UpdateFirewllRules(updateFirewllRulesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				firewallRulesService.EnableRetries(0, 0)
				result, response, operationErr = firewallRulesService.UpdateFirewllRules(updateFirewllRulesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
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
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": [{"id": "52161eb6af4241bb9d4b32394be72fdf", "paused": false, "description": "JS challenge site", "action": "js_challenge", "filter": {"id": "6f58318e7fa2477a23112e8118c66f61", "paused": true, "description": "Login from office", "expression": "ip.src eq 93.184.216.0 and (http.request.uri.path ~ \"^.*/wp-login.php$\" or http.request.uri.path ~ \"^.*/xmlrpc.php$\")"}, "created_on": "2019-01-01T05:20:00.123Z", "modified_on": "2019-01-01T05:20:00.123Z"}]}`)
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
				firewallRulesUpdateInputItemModel.Priority = core.Int64Ptr(int64(1))
				firewallRulesUpdateInputItemModel.Description = core.StringPtr("JS challenge site")
				firewallRulesUpdateInputItemModel.Filter = firewallRulesUpdateInputItemFilterModel

				// Construct an instance of the UpdateFirewllRulesOptions model
				updateFirewllRulesOptionsModel := new(firewallrulesv1.UpdateFirewllRulesOptions)
				updateFirewllRulesOptionsModel.XAuthUserToken = core.StringPtr("testString")
				updateFirewllRulesOptionsModel.Crn = core.StringPtr("testString")
				updateFirewllRulesOptionsModel.ZoneIdentifier = core.StringPtr("testString")
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
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": [{"id": "52161eb6af4241bb9d4b32394be72fdf", "paused": false, "description": "JS challenge site", "action": "js_challenge", "filter": {"id": "6f58318e7fa2477a23112e8118c66f61", "paused": true, "description": "Login from office", "expression": "ip.src eq 93.184.216.0 and (http.request.uri.path ~ \"^.*/wp-login.php$\" or http.request.uri.path ~ \"^.*/xmlrpc.php$\")"}, "created_on": "2019-01-01T05:20:00.123Z", "modified_on": "2019-01-01T05:20:00.123Z"}]}`)
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
				firewallRulesUpdateInputItemModel.Priority = core.Int64Ptr(int64(1))
				firewallRulesUpdateInputItemModel.Description = core.StringPtr("JS challenge site")
				firewallRulesUpdateInputItemModel.Filter = firewallRulesUpdateInputItemFilterModel

				// Construct an instance of the UpdateFirewllRulesOptions model
				updateFirewllRulesOptionsModel := new(firewallrulesv1.UpdateFirewllRulesOptions)
				updateFirewllRulesOptionsModel.XAuthUserToken = core.StringPtr("testString")
				updateFirewllRulesOptionsModel.Crn = core.StringPtr("testString")
				updateFirewllRulesOptionsModel.ZoneIdentifier = core.StringPtr("testString")
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
				firewallRulesUpdateInputItemModel.Priority = core.Int64Ptr(int64(1))
				firewallRulesUpdateInputItemModel.Description = core.StringPtr("JS challenge site")
				firewallRulesUpdateInputItemModel.Filter = firewallRulesUpdateInputItemFilterModel

				// Construct an instance of the UpdateFirewllRulesOptions model
				updateFirewllRulesOptionsModel := new(firewallrulesv1.UpdateFirewllRulesOptions)
				updateFirewllRulesOptionsModel.XAuthUserToken = core.StringPtr("testString")
				updateFirewllRulesOptionsModel.Crn = core.StringPtr("testString")
				updateFirewllRulesOptionsModel.ZoneIdentifier = core.StringPtr("testString")
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
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke UpdateFirewllRules successfully`, func() {
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
				firewallRulesUpdateInputItemModel.Priority = core.Int64Ptr(int64(1))
				firewallRulesUpdateInputItemModel.Description = core.StringPtr("JS challenge site")
				firewallRulesUpdateInputItemModel.Filter = firewallRulesUpdateInputItemFilterModel

				// Construct an instance of the UpdateFirewllRulesOptions model
				updateFirewllRulesOptionsModel := new(firewallrulesv1.UpdateFirewllRulesOptions)
				updateFirewllRulesOptionsModel.XAuthUserToken = core.StringPtr("testString")
				updateFirewllRulesOptionsModel.Crn = core.StringPtr("testString")
				updateFirewllRulesOptionsModel.ZoneIdentifier = core.StringPtr("testString")
				updateFirewllRulesOptionsModel.FirewallRulesUpdateInputItem = []firewallrulesv1.FirewallRulesUpdateInputItem{*firewallRulesUpdateInputItemModel}
				updateFirewllRulesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := firewallRulesService.UpdateFirewllRules(updateFirewllRulesOptionsModel)
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
	Describe(`DeleteFirewallRules(deleteFirewallRulesOptions *DeleteFirewallRulesOptions) - Operation response error`, func() {
		deleteFirewallRulesPath := "/v1/testString/zones/testString/firewall/rules"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteFirewallRulesPath))
					Expect(req.Method).To(Equal("DELETE"))
					Expect(req.Header["X-Auth-User-Token"]).ToNot(BeNil())
					Expect(req.Header["X-Auth-User-Token"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["id"]).To(Equal([]string{"f2d427378e7542acb295380d352e2ebd"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke DeleteFirewallRules with error: Operation response processing error`, func() {
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
				deleteFirewallRulesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := firewallRulesService.DeleteFirewallRules(deleteFirewallRulesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				firewallRulesService.EnableRetries(0, 0)
				result, response, operationErr = firewallRulesService.DeleteFirewallRules(deleteFirewallRulesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
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
					Expect(req.URL.Query()["id"]).To(Equal([]string{"f2d427378e7542acb295380d352e2ebd"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": [{"id": "f2d427378e7542acb295380d352e2ebd"}]}`)
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
					Expect(req.URL.Query()["id"]).To(Equal([]string{"f2d427378e7542acb295380d352e2ebd"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": [{"id": "f2d427378e7542acb295380d352e2ebd"}]}`)
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
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke DeleteFirewallRules successfully`, func() {
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
				deleteFirewallRulesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := firewallRulesService.DeleteFirewallRules(deleteFirewallRulesOptionsModel)
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
	Describe(`DeleteFirewallRule(deleteFirewallRuleOptions *DeleteFirewallRuleOptions) - Operation response error`, func() {
		deleteFirewallRulePath := "/v1/testString/zones/testString/firewall/rules/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteFirewallRulePath))
					Expect(req.Method).To(Equal("DELETE"))
					Expect(req.Header["X-Auth-User-Token"]).ToNot(BeNil())
					Expect(req.Header["X-Auth-User-Token"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke DeleteFirewallRule with error: Operation response processing error`, func() {
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
				deleteFirewallRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := firewallRulesService.DeleteFirewallRule(deleteFirewallRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				firewallRulesService.EnableRetries(0, 0)
				result, response, operationErr = firewallRulesService.DeleteFirewallRule(deleteFirewallRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
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
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": "f2d427378e7542acb295380d352e2ebd"}}`)
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
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": "f2d427378e7542acb295380d352e2ebd"}}`)
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
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke DeleteFirewallRule successfully`, func() {
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
				deleteFirewallRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := firewallRulesService.DeleteFirewallRule(deleteFirewallRuleOptionsModel)
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
	Describe(`GetFirewallRule(getFirewallRuleOptions *GetFirewallRuleOptions) - Operation response error`, func() {
		getFirewallRulePath := "/v1/testString/zones/testString/firewall/rules/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getFirewallRulePath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["X-Auth-User-Token"]).ToNot(BeNil())
					Expect(req.Header["X-Auth-User-Token"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetFirewallRule with error: Operation response processing error`, func() {
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
				getFirewallRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := firewallRulesService.GetFirewallRule(getFirewallRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				firewallRulesService.EnableRetries(0, 0)
				result, response, operationErr = firewallRulesService.GetFirewallRule(getFirewallRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
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
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": "52161eb6af4241bb9d4b32394be72fdf", "paused": false, "description": "JS challenge site", "action": "js_challenge", "filter": {"id": "6f58318e7fa2477a23112e8118c66f61", "paused": true, "description": "Login from office", "expression": "ip.src eq 93.184.216.0 and (http.request.uri.path ~ \"^.*/wp-login.php$\" or http.request.uri.path ~ \"^.*/xmlrpc.php$\")"}, "created_on": "2019-01-01T05:20:00.123Z", "modified_on": "2019-01-01T05:20:00.123Z"}}`)
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
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": "52161eb6af4241bb9d4b32394be72fdf", "paused": false, "description": "JS challenge site", "action": "js_challenge", "filter": {"id": "6f58318e7fa2477a23112e8118c66f61", "paused": true, "description": "Login from office", "expression": "ip.src eq 93.184.216.0 and (http.request.uri.path ~ \"^.*/wp-login.php$\" or http.request.uri.path ~ \"^.*/xmlrpc.php$\")"}, "created_on": "2019-01-01T05:20:00.123Z", "modified_on": "2019-01-01T05:20:00.123Z"}}`)
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
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetFirewallRule successfully`, func() {
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
				getFirewallRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := firewallRulesService.GetFirewallRule(getFirewallRuleOptionsModel)
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
	Describe(`UpdateFirewallRule(updateFirewallRuleOptions *UpdateFirewallRuleOptions) - Operation response error`, func() {
		updateFirewallRulePath := "/v1/testString/zones/testString/firewall/rules/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateFirewallRulePath))
					Expect(req.Method).To(Equal("PUT"))
					Expect(req.Header["X-Auth-User-Token"]).ToNot(BeNil())
					Expect(req.Header["X-Auth-User-Token"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateFirewallRule with error: Operation response processing error`, func() {
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
				updateFirewallRuleOptionsModel.Action = core.StringPtr("js_challenge")
				updateFirewallRuleOptionsModel.Paused = core.BoolPtr(false)
				updateFirewallRuleOptionsModel.Priority = core.Int64Ptr(int64(1))
				updateFirewallRuleOptionsModel.Description = core.StringPtr("JS challenge site")
				updateFirewallRuleOptionsModel.Filter = firewallRuleUpdateInputFilterModel
				updateFirewallRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := firewallRulesService.UpdateFirewallRule(updateFirewallRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				firewallRulesService.EnableRetries(0, 0)
				result, response, operationErr = firewallRulesService.UpdateFirewallRule(updateFirewallRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
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
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": "52161eb6af4241bb9d4b32394be72fdf", "paused": false, "description": "JS challenge site", "action": "js_challenge", "filter": {"id": "6f58318e7fa2477a23112e8118c66f61", "paused": true, "description": "Login from office", "expression": "ip.src eq 93.184.216.0 and (http.request.uri.path ~ \"^.*/wp-login.php$\" or http.request.uri.path ~ \"^.*/xmlrpc.php$\")"}, "created_on": "2019-01-01T05:20:00.123Z", "modified_on": "2019-01-01T05:20:00.123Z"}}`)
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
				updateFirewallRuleOptionsModel.Action = core.StringPtr("js_challenge")
				updateFirewallRuleOptionsModel.Paused = core.BoolPtr(false)
				updateFirewallRuleOptionsModel.Priority = core.Int64Ptr(int64(1))
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
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": "52161eb6af4241bb9d4b32394be72fdf", "paused": false, "description": "JS challenge site", "action": "js_challenge", "filter": {"id": "6f58318e7fa2477a23112e8118c66f61", "paused": true, "description": "Login from office", "expression": "ip.src eq 93.184.216.0 and (http.request.uri.path ~ \"^.*/wp-login.php$\" or http.request.uri.path ~ \"^.*/xmlrpc.php$\")"}, "created_on": "2019-01-01T05:20:00.123Z", "modified_on": "2019-01-01T05:20:00.123Z"}}`)
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
				updateFirewallRuleOptionsModel.Action = core.StringPtr("js_challenge")
				updateFirewallRuleOptionsModel.Paused = core.BoolPtr(false)
				updateFirewallRuleOptionsModel.Priority = core.Int64Ptr(int64(1))
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
				updateFirewallRuleOptionsModel.Action = core.StringPtr("js_challenge")
				updateFirewallRuleOptionsModel.Paused = core.BoolPtr(false)
				updateFirewallRuleOptionsModel.Priority = core.Int64Ptr(int64(1))
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
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke UpdateFirewallRule successfully`, func() {
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
				updateFirewallRuleOptionsModel.Action = core.StringPtr("js_challenge")
				updateFirewallRuleOptionsModel.Paused = core.BoolPtr(false)
				updateFirewallRuleOptionsModel.Priority = core.Int64Ptr(int64(1))
				updateFirewallRuleOptionsModel.Description = core.StringPtr("JS challenge site")
				updateFirewallRuleOptionsModel.Filter = firewallRuleUpdateInputFilterModel
				updateFirewallRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := firewallRulesService.UpdateFirewallRule(updateFirewallRuleOptionsModel)
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
			firewallRulesService, _ := firewallrulesv1.NewFirewallRulesV1(&firewallrulesv1.FirewallRulesV1Options{
				URL:           "http://firewallrulesv1modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
			})
			It(`Invoke NewCreateFirewallRulesOptions successfully`, func() {
				// Construct an instance of the FirewallRuleInputFilterID model
				firewallRuleInputFilterModel := new(firewallrulesv1.FirewallRuleInputFilterID)
				Expect(firewallRuleInputFilterModel).ToNot(BeNil())
				firewallRuleInputFilterModel.ID = core.StringPtr("6f58318e7fa2477a23112e8118c66f61")
				Expect(firewallRuleInputFilterModel.ID).To(Equal(core.StringPtr("6f58318e7fa2477a23112e8118c66f61")))

				// Construct an instance of the FirewallRuleInput model
				firewallRuleInputModel := new(firewallrulesv1.FirewallRuleInput)
				Expect(firewallRuleInputModel).ToNot(BeNil())
				firewallRuleInputModel.Filter = firewallRuleInputFilterModel
				firewallRuleInputModel.Action = core.StringPtr("js_challenge")
				firewallRuleInputModel.Description = core.StringPtr("JS challenge site")
				firewallRuleInputModel.Paused = core.BoolPtr(false)
				firewallRuleInputModel.Priority = core.Int64Ptr(int64(1))
				Expect(firewallRuleInputModel.Filter).To(Equal(firewallRuleInputFilterModel))
				Expect(firewallRuleInputModel.Action).To(Equal(core.StringPtr("js_challenge")))
				Expect(firewallRuleInputModel.Description).To(Equal(core.StringPtr("JS challenge site")))
				Expect(firewallRuleInputModel.Paused).To(Equal(core.BoolPtr(false)))
				Expect(firewallRuleInputModel.Priority).To(Equal(core.Int64Ptr(int64(1))))

				// Construct an instance of the CreateFirewallRulesOptions model
				xAuthUserToken := "testString"
				crn := "testString"
				zoneIdentifier := "testString"
				createFirewallRulesOptionsModel := firewallRulesService.NewCreateFirewallRulesOptions(xAuthUserToken, crn, zoneIdentifier)
				createFirewallRulesOptionsModel.SetXAuthUserToken("testString")
				createFirewallRulesOptionsModel.SetCrn("testString")
				createFirewallRulesOptionsModel.SetZoneIdentifier("testString")
				createFirewallRulesOptionsModel.SetFirewallRuleInput([]firewallrulesv1.FirewallRuleInput{*firewallRuleInputModel})
				createFirewallRulesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createFirewallRulesOptionsModel).ToNot(BeNil())
				Expect(createFirewallRulesOptionsModel.XAuthUserToken).To(Equal(core.StringPtr("testString")))
				Expect(createFirewallRulesOptionsModel.Crn).To(Equal(core.StringPtr("testString")))
				Expect(createFirewallRulesOptionsModel.ZoneIdentifier).To(Equal(core.StringPtr("testString")))
				Expect(createFirewallRulesOptionsModel.FirewallRuleInput).To(Equal([]firewallrulesv1.FirewallRuleInput{*firewallRuleInputModel}))
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
				deleteFirewallRuleOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteFirewallRuleOptionsModel).ToNot(BeNil())
				Expect(deleteFirewallRuleOptionsModel.XAuthUserToken).To(Equal(core.StringPtr("testString")))
				Expect(deleteFirewallRuleOptionsModel.Crn).To(Equal(core.StringPtr("testString")))
				Expect(deleteFirewallRuleOptionsModel.ZoneIdentifier).To(Equal(core.StringPtr("testString")))
				Expect(deleteFirewallRuleOptionsModel.FirewallRuleIdentifier).To(Equal(core.StringPtr("testString")))
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
				deleteFirewallRulesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteFirewallRulesOptionsModel).ToNot(BeNil())
				Expect(deleteFirewallRulesOptionsModel.XAuthUserToken).To(Equal(core.StringPtr("testString")))
				Expect(deleteFirewallRulesOptionsModel.Crn).To(Equal(core.StringPtr("testString")))
				Expect(deleteFirewallRulesOptionsModel.ZoneIdentifier).To(Equal(core.StringPtr("testString")))
				Expect(deleteFirewallRulesOptionsModel.ID).To(Equal(core.StringPtr("f2d427378e7542acb295380d352e2ebd")))
				Expect(deleteFirewallRulesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewFirewallRuleUpdateInputFilter successfully`, func() {
				id := "6f58318e7fa2477a23112e8118c66f61"
				_model, err := firewallRulesService.NewFirewallRuleUpdateInputFilter(id)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewFirewallRulesUpdateInputItem successfully`, func() {
				id := "52161eb6af4241bb9d4b32394be72fdf"
				action := "js_challenge"
				_model, err := firewallRulesService.NewFirewallRulesUpdateInputItem(id, action)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewFirewallRulesUpdateInputItemFilter successfully`, func() {
				id := "6f58318e7fa2477a23112e8118c66f61"
				_model, err := firewallRulesService.NewFirewallRulesUpdateInputItemFilter(id)
				Expect(_model).ToNot(BeNil())
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
				getFirewallRuleOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getFirewallRuleOptionsModel).ToNot(BeNil())
				Expect(getFirewallRuleOptionsModel.XAuthUserToken).To(Equal(core.StringPtr("testString")))
				Expect(getFirewallRuleOptionsModel.Crn).To(Equal(core.StringPtr("testString")))
				Expect(getFirewallRuleOptionsModel.ZoneIdentifier).To(Equal(core.StringPtr("testString")))
				Expect(getFirewallRuleOptionsModel.FirewallRuleIdentifier).To(Equal(core.StringPtr("testString")))
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
				listAllFirewallRulesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listAllFirewallRulesOptionsModel).ToNot(BeNil())
				Expect(listAllFirewallRulesOptionsModel.XAuthUserToken).To(Equal(core.StringPtr("testString")))
				Expect(listAllFirewallRulesOptionsModel.Crn).To(Equal(core.StringPtr("testString")))
				Expect(listAllFirewallRulesOptionsModel.ZoneIdentifier).To(Equal(core.StringPtr("testString")))
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
				updateFirewallRuleOptionsModel.SetAction("js_challenge")
				updateFirewallRuleOptionsModel.SetPaused(false)
				updateFirewallRuleOptionsModel.SetPriority(int64(1))
				updateFirewallRuleOptionsModel.SetDescription("JS challenge site")
				updateFirewallRuleOptionsModel.SetFilter(firewallRuleUpdateInputFilterModel)
				updateFirewallRuleOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateFirewallRuleOptionsModel).ToNot(BeNil())
				Expect(updateFirewallRuleOptionsModel.XAuthUserToken).To(Equal(core.StringPtr("testString")))
				Expect(updateFirewallRuleOptionsModel.Crn).To(Equal(core.StringPtr("testString")))
				Expect(updateFirewallRuleOptionsModel.ZoneIdentifier).To(Equal(core.StringPtr("testString")))
				Expect(updateFirewallRuleOptionsModel.FirewallRuleIdentifier).To(Equal(core.StringPtr("testString")))
				Expect(updateFirewallRuleOptionsModel.Action).To(Equal(core.StringPtr("js_challenge")))
				Expect(updateFirewallRuleOptionsModel.Paused).To(Equal(core.BoolPtr(false)))
				Expect(updateFirewallRuleOptionsModel.Priority).To(Equal(core.Int64Ptr(int64(1))))
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
				firewallRulesUpdateInputItemModel.Priority = core.Int64Ptr(int64(1))
				firewallRulesUpdateInputItemModel.Description = core.StringPtr("JS challenge site")
				firewallRulesUpdateInputItemModel.Filter = firewallRulesUpdateInputItemFilterModel
				Expect(firewallRulesUpdateInputItemModel.ID).To(Equal(core.StringPtr("52161eb6af4241bb9d4b32394be72fdf")))
				Expect(firewallRulesUpdateInputItemModel.Action).To(Equal(core.StringPtr("js_challenge")))
				Expect(firewallRulesUpdateInputItemModel.Paused).To(Equal(core.BoolPtr(false)))
				Expect(firewallRulesUpdateInputItemModel.Priority).To(Equal(core.Int64Ptr(int64(1))))
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
				updateFirewllRulesOptionsModel.SetFirewallRulesUpdateInputItem([]firewallrulesv1.FirewallRulesUpdateInputItem{*firewallRulesUpdateInputItemModel})
				updateFirewllRulesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateFirewllRulesOptionsModel).ToNot(BeNil())
				Expect(updateFirewllRulesOptionsModel.XAuthUserToken).To(Equal(core.StringPtr("testString")))
				Expect(updateFirewllRulesOptionsModel.Crn).To(Equal(core.StringPtr("testString")))
				Expect(updateFirewllRulesOptionsModel.ZoneIdentifier).To(Equal(core.StringPtr("testString")))
				Expect(updateFirewllRulesOptionsModel.FirewallRulesUpdateInputItem).To(Equal([]firewallrulesv1.FirewallRulesUpdateInputItem{*firewallRulesUpdateInputItemModel}))
				Expect(updateFirewllRulesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewFirewallRuleInput successfully`, func() {
				var filter firewallrulesv1.FirewallRuleInputFilterIntf = nil
				action := "js_challenge"
				_, err := firewallRulesService.NewFirewallRuleInput(filter, action)
				Expect(err).ToNot(BeNil())
			})
			It(`Invoke NewFirewallRuleInputFilterExpression successfully`, func() {
				expression := `not http.request.uri.path matches "^/api/.*$"`
				_model, err := firewallRulesService.NewFirewallRuleInputFilterExpression(expression)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewFirewallRuleInputFilterID successfully`, func() {
				id := "6f58318e7fa2477a23112e8118c66f61"
				_model, err := firewallRulesService.NewFirewallRuleInputFilterID(id)
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
