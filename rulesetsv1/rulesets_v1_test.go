/**
 * (C) Copyright IBM Corp. 2024.
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

package rulesetsv1_test

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
	"github.com/IBM/networking-go-sdk/rulesetsv1"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe(`RulesetsV1`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		It(`Instantiate service client`, func() {
			rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
				Authenticator:  &core.NoAuthAuthenticator{},
				Crn:            core.StringPtr(crn),
				ZoneIdentifier: core.StringPtr(zoneIdentifier),
			})
			Expect(rulesetsService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
				URL:            "{BAD_URL_STRING",
				Crn:            core.StringPtr(crn),
				ZoneIdentifier: core.StringPtr(zoneIdentifier),
			})
			Expect(rulesetsService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
				URL:            "https://rulesetsv1/api",
				Crn:            core.StringPtr(crn),
				ZoneIdentifier: core.StringPtr(zoneIdentifier),
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(rulesetsService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Validation Error`, func() {
			rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{})
			Expect(rulesetsService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"RULESETS_URL":       "https://rulesetsv1/api",
				"RULESETS_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1UsingExternalConfig(&rulesetsv1.RulesetsV1Options{
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(rulesetsService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := rulesetsService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != rulesetsService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(rulesetsService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(rulesetsService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1UsingExternalConfig(&rulesetsv1.RulesetsV1Options{
					URL:            "https://testService/api",
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(rulesetsService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(rulesetsService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := rulesetsService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != rulesetsService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(rulesetsService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(rulesetsService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1UsingExternalConfig(&rulesetsv1.RulesetsV1Options{
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				err := rulesetsService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(rulesetsService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(rulesetsService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := rulesetsService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != rulesetsService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(rulesetsService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(rulesetsService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"RULESETS_URL":       "https://rulesetsv1/api",
				"RULESETS_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1UsingExternalConfig(&rulesetsv1.RulesetsV1Options{
				Crn:            core.StringPtr(crn),
				ZoneIdentifier: core.StringPtr(zoneIdentifier),
			})

			It(`Instantiate service client with error`, func() {
				Expect(rulesetsService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"RULESETS_AUTH_TYPE": "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1UsingExternalConfig(&rulesetsv1.RulesetsV1Options{
				URL:            "{BAD_URL_STRING",
				Crn:            core.StringPtr(crn),
				ZoneIdentifier: core.StringPtr(zoneIdentifier),
			})

			It(`Instantiate service client with error`, func() {
				Expect(rulesetsService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = rulesetsv1.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`GetAccountRulesets(getAccountRulesetsOptions *GetAccountRulesetsOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getAccountRulesetsPath := "/v1/testString/rulesets"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getAccountRulesetsPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetAccountRulesets with error: Operation response processing error`, func() {
				rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(rulesetsService).ToNot(BeNil())

				// Construct an instance of the GetAccountRulesetsOptions model
				getAccountRulesetsOptionsModel := new(rulesetsv1.GetAccountRulesetsOptions)
				getAccountRulesetsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := rulesetsService.GetAccountRulesets(getAccountRulesetsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				rulesetsService.EnableRetries(0, 0)
				result, response, operationErr = rulesetsService.GetAccountRulesets(getAccountRulesetsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetAccountRulesets(getAccountRulesetsOptions *GetAccountRulesetsOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getAccountRulesetsPath := "/v1/testString/rulesets"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getAccountRulesetsPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "messages": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "result": [{"description": "Custom account ruleset", "id": "ID", "kind": "managed", "last_updated": "2000-01-01T00:00:00.000000Z", "name": "Name", "phase": "ddos_l4", "version": "1"}]}`)
				}))
			})
			It(`Invoke GetAccountRulesets successfully with retries`, func() {
				rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(rulesetsService).ToNot(BeNil())
				rulesetsService.EnableRetries(0, 0)

				// Construct an instance of the GetAccountRulesetsOptions model
				getAccountRulesetsOptionsModel := new(rulesetsv1.GetAccountRulesetsOptions)
				getAccountRulesetsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := rulesetsService.GetAccountRulesetsWithContext(ctx, getAccountRulesetsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				rulesetsService.DisableRetries()
				result, response, operationErr := rulesetsService.GetAccountRulesets(getAccountRulesetsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = rulesetsService.GetAccountRulesetsWithContext(ctx, getAccountRulesetsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getAccountRulesetsPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "messages": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "result": [{"description": "Custom account ruleset", "id": "ID", "kind": "managed", "last_updated": "2000-01-01T00:00:00.000000Z", "name": "Name", "phase": "ddos_l4", "version": "1"}]}`)
				}))
			})
			It(`Invoke GetAccountRulesets successfully`, func() {
				rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(rulesetsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := rulesetsService.GetAccountRulesets(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetAccountRulesetsOptions model
				getAccountRulesetsOptionsModel := new(rulesetsv1.GetAccountRulesetsOptions)
				getAccountRulesetsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = rulesetsService.GetAccountRulesets(getAccountRulesetsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetAccountRulesets with error: Operation request error`, func() {
				rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(rulesetsService).ToNot(BeNil())

				// Construct an instance of the GetAccountRulesetsOptions model
				getAccountRulesetsOptionsModel := new(rulesetsv1.GetAccountRulesetsOptions)
				getAccountRulesetsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := rulesetsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := rulesetsService.GetAccountRulesets(getAccountRulesetsOptionsModel)
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
			It(`Invoke GetAccountRulesets successfully`, func() {
				rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(rulesetsService).ToNot(BeNil())

				// Construct an instance of the GetAccountRulesetsOptions model
				getAccountRulesetsOptionsModel := new(rulesetsv1.GetAccountRulesetsOptions)
				getAccountRulesetsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := rulesetsService.GetAccountRulesets(getAccountRulesetsOptionsModel)
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
	Describe(`GetAccountRuleset(getAccountRulesetOptions *GetAccountRulesetOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getAccountRulesetPath := "/v1/testString/rulesets/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getAccountRulesetPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetAccountRuleset with error: Operation response processing error`, func() {
				rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(rulesetsService).ToNot(BeNil())

				// Construct an instance of the GetAccountRulesetOptions model
				getAccountRulesetOptionsModel := new(rulesetsv1.GetAccountRulesetOptions)
				getAccountRulesetOptionsModel.RulesetID = core.StringPtr("testString")
				getAccountRulesetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := rulesetsService.GetAccountRuleset(getAccountRulesetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				rulesetsService.EnableRetries(0, 0)
				result, response, operationErr = rulesetsService.GetAccountRuleset(getAccountRulesetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetAccountRuleset(getAccountRulesetOptions *GetAccountRulesetOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getAccountRulesetPath := "/v1/testString/rulesets/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getAccountRulesetPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "messages": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "result": {"description": "Custom account ruleset", "id": "ID", "kind": "managed", "last_updated": "2000-01-01T00:00:00.000000Z", "name": "Name", "phase": "ddos_l4", "version": "1", "rules": [{"id": "ID", "version": "Version", "action": "Action", "action_parameters": {"id": "ID", "overrides": {"action": "Action", "enabled": false, "sensitivity_level": "high", "rules": [{"id": "ID", "enabled": false, "action": "Action", "sensitivity_level": "high"}], "categories": [{"category": "Category", "enabled": false, "action": "Action"}]}, "version": "Version", "ruleset": "Ruleset", "rulesets": ["Rulesets"], "response": {"content": "{\"success\": false, \"error\": \"you have been blocked\"}", "content_type": "application/json", "status_code": 400}}, "expression": "ip.src ne 1.1.1.1", "ref": "my_ref", "logging": {"enabled": true}, "last_updated": "2000-01-01T00:00:00.000000Z"}]}}`)
				}))
			})
			It(`Invoke GetAccountRuleset successfully with retries`, func() {
				rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(rulesetsService).ToNot(BeNil())
				rulesetsService.EnableRetries(0, 0)

				// Construct an instance of the GetAccountRulesetOptions model
				getAccountRulesetOptionsModel := new(rulesetsv1.GetAccountRulesetOptions)
				getAccountRulesetOptionsModel.RulesetID = core.StringPtr("testString")
				getAccountRulesetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := rulesetsService.GetAccountRulesetWithContext(ctx, getAccountRulesetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				rulesetsService.DisableRetries()
				result, response, operationErr := rulesetsService.GetAccountRuleset(getAccountRulesetOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = rulesetsService.GetAccountRulesetWithContext(ctx, getAccountRulesetOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getAccountRulesetPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "messages": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "result": {"description": "Custom account ruleset", "id": "ID", "kind": "managed", "last_updated": "2000-01-01T00:00:00.000000Z", "name": "Name", "phase": "ddos_l4", "version": "1", "rules": [{"id": "ID", "version": "Version", "action": "Action", "action_parameters": {"id": "ID", "overrides": {"action": "Action", "enabled": false, "sensitivity_level": "high", "rules": [{"id": "ID", "enabled": false, "action": "Action", "sensitivity_level": "high"}], "categories": [{"category": "Category", "enabled": false, "action": "Action"}]}, "version": "Version", "ruleset": "Ruleset", "rulesets": ["Rulesets"], "response": {"content": "{\"success\": false, \"error\": \"you have been blocked\"}", "content_type": "application/json", "status_code": 400}}, "expression": "ip.src ne 1.1.1.1", "ref": "my_ref", "logging": {"enabled": true}, "last_updated": "2000-01-01T00:00:00.000000Z"}]}}`)
				}))
			})
			It(`Invoke GetAccountRuleset successfully`, func() {
				rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(rulesetsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := rulesetsService.GetAccountRuleset(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetAccountRulesetOptions model
				getAccountRulesetOptionsModel := new(rulesetsv1.GetAccountRulesetOptions)
				getAccountRulesetOptionsModel.RulesetID = core.StringPtr("testString")
				getAccountRulesetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = rulesetsService.GetAccountRuleset(getAccountRulesetOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetAccountRuleset with error: Operation validation and request error`, func() {
				rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(rulesetsService).ToNot(BeNil())

				// Construct an instance of the GetAccountRulesetOptions model
				getAccountRulesetOptionsModel := new(rulesetsv1.GetAccountRulesetOptions)
				getAccountRulesetOptionsModel.RulesetID = core.StringPtr("testString")
				getAccountRulesetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := rulesetsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := rulesetsService.GetAccountRuleset(getAccountRulesetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetAccountRulesetOptions model with no property values
				getAccountRulesetOptionsModelNew := new(rulesetsv1.GetAccountRulesetOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = rulesetsService.GetAccountRuleset(getAccountRulesetOptionsModelNew)
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
			It(`Invoke GetAccountRuleset successfully`, func() {
				rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(rulesetsService).ToNot(BeNil())

				// Construct an instance of the GetAccountRulesetOptions model
				getAccountRulesetOptionsModel := new(rulesetsv1.GetAccountRulesetOptions)
				getAccountRulesetOptionsModel.RulesetID = core.StringPtr("testString")
				getAccountRulesetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := rulesetsService.GetAccountRuleset(getAccountRulesetOptionsModel)
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
	Describe(`UpdateAccountRuleset(updateAccountRulesetOptions *UpdateAccountRulesetOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		updateAccountRulesetPath := "/v1/testString/rulesets/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateAccountRulesetPath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateAccountRuleset with error: Operation response processing error`, func() {
				rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(rulesetsService).ToNot(BeNil())

				// Construct an instance of the RulesOverride model
				rulesOverrideModel := new(rulesetsv1.RulesOverride)
				rulesOverrideModel.ID = core.StringPtr("testString")
				rulesOverrideModel.Enabled = core.BoolPtr(true)
				rulesOverrideModel.Action = core.StringPtr("testString")
				rulesOverrideModel.SensitivityLevel = core.StringPtr("high")

				// Construct an instance of the CategoriesOverride model
				categoriesOverrideModel := new(rulesetsv1.CategoriesOverride)
				categoriesOverrideModel.Category = core.StringPtr("testString")
				categoriesOverrideModel.Enabled = core.BoolPtr(true)
				categoriesOverrideModel.Action = core.StringPtr("testString")

				// Construct an instance of the Overrides model
				overridesModel := new(rulesetsv1.Overrides)
				overridesModel.Action = core.StringPtr("testString")
				overridesModel.Enabled = core.BoolPtr(true)
				overridesModel.SensitivityLevel = core.StringPtr("high")
				overridesModel.Rules = []rulesetsv1.RulesOverride{*rulesOverrideModel}
				overridesModel.Categories = []rulesetsv1.CategoriesOverride{*categoriesOverrideModel}

				// Construct an instance of the ActionParametersResponse model
				actionParametersResponseModel := new(rulesetsv1.ActionParametersResponse)
				actionParametersResponseModel.Content = core.StringPtr(`{"success": false, "error": "you have been blocked"}`)
				actionParametersResponseModel.ContentType = core.StringPtr("application/json")
				actionParametersResponseModel.StatusCode = core.Int64Ptr(int64(400))

				// Construct an instance of the ActionParameters model
				actionParametersModel := new(rulesetsv1.ActionParameters)
				actionParametersModel.ID = core.StringPtr("testString")
				actionParametersModel.Overrides = overridesModel
				actionParametersModel.Version = core.StringPtr("testString")
				actionParametersModel.Ruleset = core.StringPtr("testString")
				actionParametersModel.Rulesets = []string{"testString"}
				actionParametersModel.Response = actionParametersResponseModel

				// Construct an instance of the Logging model
				loggingModel := new(rulesetsv1.Logging)
				loggingModel.Enabled = core.BoolPtr(true)

				// Construct an instance of the Position model
				positionModel := new(rulesetsv1.Position)
				positionModel.Before = core.StringPtr("testString")
				positionModel.After = core.StringPtr("testString")
				positionModel.Index = core.Int64Ptr(int64(0))

				// Construct an instance of the RuleCreate model
				ruleCreateModel := new(rulesetsv1.RuleCreate)
				ruleCreateModel.Action = core.StringPtr("testString")
				ruleCreateModel.ActionParameters = actionParametersModel
				ruleCreateModel.Description = core.StringPtr("testString")
				ruleCreateModel.Enabled = core.BoolPtr(true)
				ruleCreateModel.Expression = core.StringPtr("ip.src ne 1.1.1.1")
				ruleCreateModel.ID = core.StringPtr("testString")
				ruleCreateModel.Logging = loggingModel
				ruleCreateModel.Ref = core.StringPtr("my_ref")
				ruleCreateModel.Position = positionModel

				// Construct an instance of the UpdateAccountRulesetOptions model
				updateAccountRulesetOptionsModel := new(rulesetsv1.UpdateAccountRulesetOptions)
				updateAccountRulesetOptionsModel.RulesetID = core.StringPtr("testString")
				updateAccountRulesetOptionsModel.Description = core.StringPtr("Custom account ruleset")
				updateAccountRulesetOptionsModel.Kind = core.StringPtr("managed")
				updateAccountRulesetOptionsModel.Name = core.StringPtr("testString")
				updateAccountRulesetOptionsModel.Phase = core.StringPtr("ddos_l4")
				updateAccountRulesetOptionsModel.Rules = []rulesetsv1.RuleCreate{*ruleCreateModel}
				updateAccountRulesetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := rulesetsService.UpdateAccountRuleset(updateAccountRulesetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				rulesetsService.EnableRetries(0, 0)
				result, response, operationErr = rulesetsService.UpdateAccountRuleset(updateAccountRulesetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateAccountRuleset(updateAccountRulesetOptions *UpdateAccountRulesetOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		updateAccountRulesetPath := "/v1/testString/rulesets/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateAccountRulesetPath))
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
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "messages": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "result": {"description": "Custom account ruleset", "id": "ID", "kind": "managed", "last_updated": "2000-01-01T00:00:00.000000Z", "name": "Name", "phase": "ddos_l4", "version": "1", "rules": [{"id": "ID", "version": "Version", "action": "Action", "action_parameters": {"id": "ID", "overrides": {"action": "Action", "enabled": false, "sensitivity_level": "high", "rules": [{"id": "ID", "enabled": false, "action": "Action", "sensitivity_level": "high"}], "categories": [{"category": "Category", "enabled": false, "action": "Action"}]}, "version": "Version", "ruleset": "Ruleset", "rulesets": ["Rulesets"], "response": {"content": "{\"success\": false, \"error\": \"you have been blocked\"}", "content_type": "application/json", "status_code": 400}}, "expression": "ip.src ne 1.1.1.1", "ref": "my_ref", "logging": {"enabled": true}, "last_updated": "2000-01-01T00:00:00.000000Z"}]}}`)
				}))
			})
			It(`Invoke UpdateAccountRuleset successfully with retries`, func() {
				rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(rulesetsService).ToNot(BeNil())
				rulesetsService.EnableRetries(0, 0)

				// Construct an instance of the RulesOverride model
				rulesOverrideModel := new(rulesetsv1.RulesOverride)
				rulesOverrideModel.ID = core.StringPtr("testString")
				rulesOverrideModel.Enabled = core.BoolPtr(true)
				rulesOverrideModel.Action = core.StringPtr("testString")
				rulesOverrideModel.SensitivityLevel = core.StringPtr("high")

				// Construct an instance of the CategoriesOverride model
				categoriesOverrideModel := new(rulesetsv1.CategoriesOverride)
				categoriesOverrideModel.Category = core.StringPtr("testString")
				categoriesOverrideModel.Enabled = core.BoolPtr(true)
				categoriesOverrideModel.Action = core.StringPtr("testString")

				// Construct an instance of the Overrides model
				overridesModel := new(rulesetsv1.Overrides)
				overridesModel.Action = core.StringPtr("testString")
				overridesModel.Enabled = core.BoolPtr(true)
				overridesModel.SensitivityLevel = core.StringPtr("high")
				overridesModel.Rules = []rulesetsv1.RulesOverride{*rulesOverrideModel}
				overridesModel.Categories = []rulesetsv1.CategoriesOverride{*categoriesOverrideModel}

				// Construct an instance of the ActionParametersResponse model
				actionParametersResponseModel := new(rulesetsv1.ActionParametersResponse)
				actionParametersResponseModel.Content = core.StringPtr(`{"success": false, "error": "you have been blocked"}`)
				actionParametersResponseModel.ContentType = core.StringPtr("application/json")
				actionParametersResponseModel.StatusCode = core.Int64Ptr(int64(400))

				// Construct an instance of the ActionParameters model
				actionParametersModel := new(rulesetsv1.ActionParameters)
				actionParametersModel.ID = core.StringPtr("testString")
				actionParametersModel.Overrides = overridesModel
				actionParametersModel.Version = core.StringPtr("testString")
				actionParametersModel.Ruleset = core.StringPtr("testString")
				actionParametersModel.Rulesets = []string{"testString"}
				actionParametersModel.Response = actionParametersResponseModel

				// Construct an instance of the Logging model
				loggingModel := new(rulesetsv1.Logging)
				loggingModel.Enabled = core.BoolPtr(true)

				// Construct an instance of the Position model
				positionModel := new(rulesetsv1.Position)
				positionModel.Before = core.StringPtr("testString")
				positionModel.After = core.StringPtr("testString")
				positionModel.Index = core.Int64Ptr(int64(0))

				// Construct an instance of the RuleCreate model
				ruleCreateModel := new(rulesetsv1.RuleCreate)
				ruleCreateModel.Action = core.StringPtr("testString")
				ruleCreateModel.ActionParameters = actionParametersModel
				ruleCreateModel.Description = core.StringPtr("testString")
				ruleCreateModel.Enabled = core.BoolPtr(true)
				ruleCreateModel.Expression = core.StringPtr("ip.src ne 1.1.1.1")
				ruleCreateModel.ID = core.StringPtr("testString")
				ruleCreateModel.Logging = loggingModel
				ruleCreateModel.Ref = core.StringPtr("my_ref")
				ruleCreateModel.Position = positionModel

				// Construct an instance of the UpdateAccountRulesetOptions model
				updateAccountRulesetOptionsModel := new(rulesetsv1.UpdateAccountRulesetOptions)
				updateAccountRulesetOptionsModel.RulesetID = core.StringPtr("testString")
				updateAccountRulesetOptionsModel.Description = core.StringPtr("Custom account ruleset")
				updateAccountRulesetOptionsModel.Kind = core.StringPtr("managed")
				updateAccountRulesetOptionsModel.Name = core.StringPtr("testString")
				updateAccountRulesetOptionsModel.Phase = core.StringPtr("ddos_l4")
				updateAccountRulesetOptionsModel.Rules = []rulesetsv1.RuleCreate{*ruleCreateModel}
				updateAccountRulesetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := rulesetsService.UpdateAccountRulesetWithContext(ctx, updateAccountRulesetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				rulesetsService.DisableRetries()
				result, response, operationErr := rulesetsService.UpdateAccountRuleset(updateAccountRulesetOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = rulesetsService.UpdateAccountRulesetWithContext(ctx, updateAccountRulesetOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(updateAccountRulesetPath))
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
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "messages": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "result": {"description": "Custom account ruleset", "id": "ID", "kind": "managed", "last_updated": "2000-01-01T00:00:00.000000Z", "name": "Name", "phase": "ddos_l4", "version": "1", "rules": [{"id": "ID", "version": "Version", "action": "Action", "action_parameters": {"id": "ID", "overrides": {"action": "Action", "enabled": false, "sensitivity_level": "high", "rules": [{"id": "ID", "enabled": false, "action": "Action", "sensitivity_level": "high"}], "categories": [{"category": "Category", "enabled": false, "action": "Action"}]}, "version": "Version", "ruleset": "Ruleset", "rulesets": ["Rulesets"], "response": {"content": "{\"success\": false, \"error\": \"you have been blocked\"}", "content_type": "application/json", "status_code": 400}}, "expression": "ip.src ne 1.1.1.1", "ref": "my_ref", "logging": {"enabled": true}, "last_updated": "2000-01-01T00:00:00.000000Z"}]}}`)
				}))
			})
			It(`Invoke UpdateAccountRuleset successfully`, func() {
				rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(rulesetsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := rulesetsService.UpdateAccountRuleset(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the RulesOverride model
				rulesOverrideModel := new(rulesetsv1.RulesOverride)
				rulesOverrideModel.ID = core.StringPtr("testString")
				rulesOverrideModel.Enabled = core.BoolPtr(true)
				rulesOverrideModel.Action = core.StringPtr("testString")
				rulesOverrideModel.SensitivityLevel = core.StringPtr("high")

				// Construct an instance of the CategoriesOverride model
				categoriesOverrideModel := new(rulesetsv1.CategoriesOverride)
				categoriesOverrideModel.Category = core.StringPtr("testString")
				categoriesOverrideModel.Enabled = core.BoolPtr(true)
				categoriesOverrideModel.Action = core.StringPtr("testString")

				// Construct an instance of the Overrides model
				overridesModel := new(rulesetsv1.Overrides)
				overridesModel.Action = core.StringPtr("testString")
				overridesModel.Enabled = core.BoolPtr(true)
				overridesModel.SensitivityLevel = core.StringPtr("high")
				overridesModel.Rules = []rulesetsv1.RulesOverride{*rulesOverrideModel}
				overridesModel.Categories = []rulesetsv1.CategoriesOverride{*categoriesOverrideModel}

				// Construct an instance of the ActionParametersResponse model
				actionParametersResponseModel := new(rulesetsv1.ActionParametersResponse)
				actionParametersResponseModel.Content = core.StringPtr(`{"success": false, "error": "you have been blocked"}`)
				actionParametersResponseModel.ContentType = core.StringPtr("application/json")
				actionParametersResponseModel.StatusCode = core.Int64Ptr(int64(400))

				// Construct an instance of the ActionParameters model
				actionParametersModel := new(rulesetsv1.ActionParameters)
				actionParametersModel.ID = core.StringPtr("testString")
				actionParametersModel.Overrides = overridesModel
				actionParametersModel.Version = core.StringPtr("testString")
				actionParametersModel.Ruleset = core.StringPtr("testString")
				actionParametersModel.Rulesets = []string{"testString"}
				actionParametersModel.Response = actionParametersResponseModel

				// Construct an instance of the Logging model
				loggingModel := new(rulesetsv1.Logging)
				loggingModel.Enabled = core.BoolPtr(true)

				// Construct an instance of the Position model
				positionModel := new(rulesetsv1.Position)
				positionModel.Before = core.StringPtr("testString")
				positionModel.After = core.StringPtr("testString")
				positionModel.Index = core.Int64Ptr(int64(0))

				// Construct an instance of the RuleCreate model
				ruleCreateModel := new(rulesetsv1.RuleCreate)
				ruleCreateModel.Action = core.StringPtr("testString")
				ruleCreateModel.ActionParameters = actionParametersModel
				ruleCreateModel.Description = core.StringPtr("testString")
				ruleCreateModel.Enabled = core.BoolPtr(true)
				ruleCreateModel.Expression = core.StringPtr("ip.src ne 1.1.1.1")
				ruleCreateModel.ID = core.StringPtr("testString")
				ruleCreateModel.Logging = loggingModel
				ruleCreateModel.Ref = core.StringPtr("my_ref")
				ruleCreateModel.Position = positionModel

				// Construct an instance of the UpdateAccountRulesetOptions model
				updateAccountRulesetOptionsModel := new(rulesetsv1.UpdateAccountRulesetOptions)
				updateAccountRulesetOptionsModel.RulesetID = core.StringPtr("testString")
				updateAccountRulesetOptionsModel.Description = core.StringPtr("Custom account ruleset")
				updateAccountRulesetOptionsModel.Kind = core.StringPtr("managed")
				updateAccountRulesetOptionsModel.Name = core.StringPtr("testString")
				updateAccountRulesetOptionsModel.Phase = core.StringPtr("ddos_l4")
				updateAccountRulesetOptionsModel.Rules = []rulesetsv1.RuleCreate{*ruleCreateModel}
				updateAccountRulesetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = rulesetsService.UpdateAccountRuleset(updateAccountRulesetOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateAccountRuleset with error: Operation validation and request error`, func() {
				rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(rulesetsService).ToNot(BeNil())

				// Construct an instance of the RulesOverride model
				rulesOverrideModel := new(rulesetsv1.RulesOverride)
				rulesOverrideModel.ID = core.StringPtr("testString")
				rulesOverrideModel.Enabled = core.BoolPtr(true)
				rulesOverrideModel.Action = core.StringPtr("testString")
				rulesOverrideModel.SensitivityLevel = core.StringPtr("high")

				// Construct an instance of the CategoriesOverride model
				categoriesOverrideModel := new(rulesetsv1.CategoriesOverride)
				categoriesOverrideModel.Category = core.StringPtr("testString")
				categoriesOverrideModel.Enabled = core.BoolPtr(true)
				categoriesOverrideModel.Action = core.StringPtr("testString")

				// Construct an instance of the Overrides model
				overridesModel := new(rulesetsv1.Overrides)
				overridesModel.Action = core.StringPtr("testString")
				overridesModel.Enabled = core.BoolPtr(true)
				overridesModel.SensitivityLevel = core.StringPtr("high")
				overridesModel.Rules = []rulesetsv1.RulesOverride{*rulesOverrideModel}
				overridesModel.Categories = []rulesetsv1.CategoriesOverride{*categoriesOverrideModel}

				// Construct an instance of the ActionParametersResponse model
				actionParametersResponseModel := new(rulesetsv1.ActionParametersResponse)
				actionParametersResponseModel.Content = core.StringPtr(`{"success": false, "error": "you have been blocked"}`)
				actionParametersResponseModel.ContentType = core.StringPtr("application/json")
				actionParametersResponseModel.StatusCode = core.Int64Ptr(int64(400))

				// Construct an instance of the ActionParameters model
				actionParametersModel := new(rulesetsv1.ActionParameters)
				actionParametersModel.ID = core.StringPtr("testString")
				actionParametersModel.Overrides = overridesModel
				actionParametersModel.Version = core.StringPtr("testString")
				actionParametersModel.Ruleset = core.StringPtr("testString")
				actionParametersModel.Rulesets = []string{"testString"}
				actionParametersModel.Response = actionParametersResponseModel

				// Construct an instance of the Logging model
				loggingModel := new(rulesetsv1.Logging)
				loggingModel.Enabled = core.BoolPtr(true)

				// Construct an instance of the Position model
				positionModel := new(rulesetsv1.Position)
				positionModel.Before = core.StringPtr("testString")
				positionModel.After = core.StringPtr("testString")
				positionModel.Index = core.Int64Ptr(int64(0))

				// Construct an instance of the RuleCreate model
				ruleCreateModel := new(rulesetsv1.RuleCreate)
				ruleCreateModel.Action = core.StringPtr("testString")
				ruleCreateModel.ActionParameters = actionParametersModel
				ruleCreateModel.Description = core.StringPtr("testString")
				ruleCreateModel.Enabled = core.BoolPtr(true)
				ruleCreateModel.Expression = core.StringPtr("ip.src ne 1.1.1.1")
				ruleCreateModel.ID = core.StringPtr("testString")
				ruleCreateModel.Logging = loggingModel
				ruleCreateModel.Ref = core.StringPtr("my_ref")
				ruleCreateModel.Position = positionModel

				// Construct an instance of the UpdateAccountRulesetOptions model
				updateAccountRulesetOptionsModel := new(rulesetsv1.UpdateAccountRulesetOptions)
				updateAccountRulesetOptionsModel.RulesetID = core.StringPtr("testString")
				updateAccountRulesetOptionsModel.Description = core.StringPtr("Custom account ruleset")
				updateAccountRulesetOptionsModel.Kind = core.StringPtr("managed")
				updateAccountRulesetOptionsModel.Name = core.StringPtr("testString")
				updateAccountRulesetOptionsModel.Phase = core.StringPtr("ddos_l4")
				updateAccountRulesetOptionsModel.Rules = []rulesetsv1.RuleCreate{*ruleCreateModel}
				updateAccountRulesetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := rulesetsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := rulesetsService.UpdateAccountRuleset(updateAccountRulesetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateAccountRulesetOptions model with no property values
				updateAccountRulesetOptionsModelNew := new(rulesetsv1.UpdateAccountRulesetOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = rulesetsService.UpdateAccountRuleset(updateAccountRulesetOptionsModelNew)
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
			It(`Invoke UpdateAccountRuleset successfully`, func() {
				rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(rulesetsService).ToNot(BeNil())

				// Construct an instance of the RulesOverride model
				rulesOverrideModel := new(rulesetsv1.RulesOverride)
				rulesOverrideModel.ID = core.StringPtr("testString")
				rulesOverrideModel.Enabled = core.BoolPtr(true)
				rulesOverrideModel.Action = core.StringPtr("testString")
				rulesOverrideModel.SensitivityLevel = core.StringPtr("high")

				// Construct an instance of the CategoriesOverride model
				categoriesOverrideModel := new(rulesetsv1.CategoriesOverride)
				categoriesOverrideModel.Category = core.StringPtr("testString")
				categoriesOverrideModel.Enabled = core.BoolPtr(true)
				categoriesOverrideModel.Action = core.StringPtr("testString")

				// Construct an instance of the Overrides model
				overridesModel := new(rulesetsv1.Overrides)
				overridesModel.Action = core.StringPtr("testString")
				overridesModel.Enabled = core.BoolPtr(true)
				overridesModel.SensitivityLevel = core.StringPtr("high")
				overridesModel.Rules = []rulesetsv1.RulesOverride{*rulesOverrideModel}
				overridesModel.Categories = []rulesetsv1.CategoriesOverride{*categoriesOverrideModel}

				// Construct an instance of the ActionParametersResponse model
				actionParametersResponseModel := new(rulesetsv1.ActionParametersResponse)
				actionParametersResponseModel.Content = core.StringPtr(`{"success": false, "error": "you have been blocked"}`)
				actionParametersResponseModel.ContentType = core.StringPtr("application/json")
				actionParametersResponseModel.StatusCode = core.Int64Ptr(int64(400))

				// Construct an instance of the ActionParameters model
				actionParametersModel := new(rulesetsv1.ActionParameters)
				actionParametersModel.ID = core.StringPtr("testString")
				actionParametersModel.Overrides = overridesModel
				actionParametersModel.Version = core.StringPtr("testString")
				actionParametersModel.Ruleset = core.StringPtr("testString")
				actionParametersModel.Rulesets = []string{"testString"}
				actionParametersModel.Response = actionParametersResponseModel

				// Construct an instance of the Logging model
				loggingModel := new(rulesetsv1.Logging)
				loggingModel.Enabled = core.BoolPtr(true)

				// Construct an instance of the Position model
				positionModel := new(rulesetsv1.Position)
				positionModel.Before = core.StringPtr("testString")
				positionModel.After = core.StringPtr("testString")
				positionModel.Index = core.Int64Ptr(int64(0))

				// Construct an instance of the RuleCreate model
				ruleCreateModel := new(rulesetsv1.RuleCreate)
				ruleCreateModel.Action = core.StringPtr("testString")
				ruleCreateModel.ActionParameters = actionParametersModel
				ruleCreateModel.Description = core.StringPtr("testString")
				ruleCreateModel.Enabled = core.BoolPtr(true)
				ruleCreateModel.Expression = core.StringPtr("ip.src ne 1.1.1.1")
				ruleCreateModel.ID = core.StringPtr("testString")
				ruleCreateModel.Logging = loggingModel
				ruleCreateModel.Ref = core.StringPtr("my_ref")
				ruleCreateModel.Position = positionModel

				// Construct an instance of the UpdateAccountRulesetOptions model
				updateAccountRulesetOptionsModel := new(rulesetsv1.UpdateAccountRulesetOptions)
				updateAccountRulesetOptionsModel.RulesetID = core.StringPtr("testString")
				updateAccountRulesetOptionsModel.Description = core.StringPtr("Custom account ruleset")
				updateAccountRulesetOptionsModel.Kind = core.StringPtr("managed")
				updateAccountRulesetOptionsModel.Name = core.StringPtr("testString")
				updateAccountRulesetOptionsModel.Phase = core.StringPtr("ddos_l4")
				updateAccountRulesetOptionsModel.Rules = []rulesetsv1.RuleCreate{*ruleCreateModel}
				updateAccountRulesetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := rulesetsService.UpdateAccountRuleset(updateAccountRulesetOptionsModel)
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
	Describe(`DeleteAccountRuleset(deleteAccountRulesetOptions *DeleteAccountRulesetOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		deleteAccountRulesetPath := "/v1/testString/rulesets/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteAccountRulesetPath))
					Expect(req.Method).To(Equal("DELETE"))

					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteAccountRuleset successfully`, func() {
				rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(rulesetsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := rulesetsService.DeleteAccountRuleset(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteAccountRulesetOptions model
				deleteAccountRulesetOptionsModel := new(rulesetsv1.DeleteAccountRulesetOptions)
				deleteAccountRulesetOptionsModel.RulesetID = core.StringPtr("testString")
				deleteAccountRulesetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = rulesetsService.DeleteAccountRuleset(deleteAccountRulesetOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteAccountRuleset with error: Operation validation and request error`, func() {
				rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(rulesetsService).ToNot(BeNil())

				// Construct an instance of the DeleteAccountRulesetOptions model
				deleteAccountRulesetOptionsModel := new(rulesetsv1.DeleteAccountRulesetOptions)
				deleteAccountRulesetOptionsModel.RulesetID = core.StringPtr("testString")
				deleteAccountRulesetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := rulesetsService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := rulesetsService.DeleteAccountRuleset(deleteAccountRulesetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteAccountRulesetOptions model with no property values
				deleteAccountRulesetOptionsModelNew := new(rulesetsv1.DeleteAccountRulesetOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = rulesetsService.DeleteAccountRuleset(deleteAccountRulesetOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetAccountRulesetVersions(getAccountRulesetVersionsOptions *GetAccountRulesetVersionsOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getAccountRulesetVersionsPath := "/v1/testString/rulesets/testString/versions"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getAccountRulesetVersionsPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetAccountRulesetVersions with error: Operation response processing error`, func() {
				rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(rulesetsService).ToNot(BeNil())

				// Construct an instance of the GetAccountRulesetVersionsOptions model
				getAccountRulesetVersionsOptionsModel := new(rulesetsv1.GetAccountRulesetVersionsOptions)
				getAccountRulesetVersionsOptionsModel.RulesetID = core.StringPtr("testString")
				getAccountRulesetVersionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := rulesetsService.GetAccountRulesetVersions(getAccountRulesetVersionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				rulesetsService.EnableRetries(0, 0)
				result, response, operationErr = rulesetsService.GetAccountRulesetVersions(getAccountRulesetVersionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetAccountRulesetVersions(getAccountRulesetVersionsOptions *GetAccountRulesetVersionsOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getAccountRulesetVersionsPath := "/v1/testString/rulesets/testString/versions"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getAccountRulesetVersionsPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "messages": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "result": [{"description": "Custom account ruleset", "id": "ID", "kind": "managed", "last_updated": "2000-01-01T00:00:00.000000Z", "name": "Name", "phase": "ddos_l4", "version": "1"}]}`)
				}))
			})
			It(`Invoke GetAccountRulesetVersions successfully with retries`, func() {
				rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(rulesetsService).ToNot(BeNil())
				rulesetsService.EnableRetries(0, 0)

				// Construct an instance of the GetAccountRulesetVersionsOptions model
				getAccountRulesetVersionsOptionsModel := new(rulesetsv1.GetAccountRulesetVersionsOptions)
				getAccountRulesetVersionsOptionsModel.RulesetID = core.StringPtr("testString")
				getAccountRulesetVersionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := rulesetsService.GetAccountRulesetVersionsWithContext(ctx, getAccountRulesetVersionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				rulesetsService.DisableRetries()
				result, response, operationErr := rulesetsService.GetAccountRulesetVersions(getAccountRulesetVersionsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = rulesetsService.GetAccountRulesetVersionsWithContext(ctx, getAccountRulesetVersionsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getAccountRulesetVersionsPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "messages": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "result": [{"description": "Custom account ruleset", "id": "ID", "kind": "managed", "last_updated": "2000-01-01T00:00:00.000000Z", "name": "Name", "phase": "ddos_l4", "version": "1"}]}`)
				}))
			})
			It(`Invoke GetAccountRulesetVersions successfully`, func() {
				rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(rulesetsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := rulesetsService.GetAccountRulesetVersions(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetAccountRulesetVersionsOptions model
				getAccountRulesetVersionsOptionsModel := new(rulesetsv1.GetAccountRulesetVersionsOptions)
				getAccountRulesetVersionsOptionsModel.RulesetID = core.StringPtr("testString")
				getAccountRulesetVersionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = rulesetsService.GetAccountRulesetVersions(getAccountRulesetVersionsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetAccountRulesetVersions with error: Operation validation and request error`, func() {
				rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(rulesetsService).ToNot(BeNil())

				// Construct an instance of the GetAccountRulesetVersionsOptions model
				getAccountRulesetVersionsOptionsModel := new(rulesetsv1.GetAccountRulesetVersionsOptions)
				getAccountRulesetVersionsOptionsModel.RulesetID = core.StringPtr("testString")
				getAccountRulesetVersionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := rulesetsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := rulesetsService.GetAccountRulesetVersions(getAccountRulesetVersionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetAccountRulesetVersionsOptions model with no property values
				getAccountRulesetVersionsOptionsModelNew := new(rulesetsv1.GetAccountRulesetVersionsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = rulesetsService.GetAccountRulesetVersions(getAccountRulesetVersionsOptionsModelNew)
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
			It(`Invoke GetAccountRulesetVersions successfully`, func() {
				rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(rulesetsService).ToNot(BeNil())

				// Construct an instance of the GetAccountRulesetVersionsOptions model
				getAccountRulesetVersionsOptionsModel := new(rulesetsv1.GetAccountRulesetVersionsOptions)
				getAccountRulesetVersionsOptionsModel.RulesetID = core.StringPtr("testString")
				getAccountRulesetVersionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := rulesetsService.GetAccountRulesetVersions(getAccountRulesetVersionsOptionsModel)
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
	Describe(`GetAccountRulesetVersion(getAccountRulesetVersionOptions *GetAccountRulesetVersionOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getAccountRulesetVersionPath := "/v1/testString/rulesets/testString/versions/1"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getAccountRulesetVersionPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetAccountRulesetVersion with error: Operation response processing error`, func() {
				rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(rulesetsService).ToNot(BeNil())

				// Construct an instance of the GetAccountRulesetVersionOptions model
				getAccountRulesetVersionOptionsModel := new(rulesetsv1.GetAccountRulesetVersionOptions)
				getAccountRulesetVersionOptionsModel.RulesetID = core.StringPtr("testString")
				getAccountRulesetVersionOptionsModel.RulesetVersion = core.StringPtr("1")
				getAccountRulesetVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := rulesetsService.GetAccountRulesetVersion(getAccountRulesetVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				rulesetsService.EnableRetries(0, 0)
				result, response, operationErr = rulesetsService.GetAccountRulesetVersion(getAccountRulesetVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetAccountRulesetVersion(getAccountRulesetVersionOptions *GetAccountRulesetVersionOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getAccountRulesetVersionPath := "/v1/testString/rulesets/testString/versions/1"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getAccountRulesetVersionPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "messages": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "result": {"description": "Custom account ruleset", "id": "ID", "kind": "managed", "last_updated": "2000-01-01T00:00:00.000000Z", "name": "Name", "phase": "ddos_l4", "version": "1", "rules": [{"id": "ID", "version": "Version", "action": "Action", "action_parameters": {"id": "ID", "overrides": {"action": "Action", "enabled": false, "sensitivity_level": "high", "rules": [{"id": "ID", "enabled": false, "action": "Action", "sensitivity_level": "high"}], "categories": [{"category": "Category", "enabled": false, "action": "Action"}]}, "version": "Version", "ruleset": "Ruleset", "rulesets": ["Rulesets"], "response": {"content": "{\"success\": false, \"error\": \"you have been blocked\"}", "content_type": "application/json", "status_code": 400}}, "expression": "ip.src ne 1.1.1.1", "ref": "my_ref", "logging": {"enabled": true}, "last_updated": "2000-01-01T00:00:00.000000Z"}]}}`)
				}))
			})
			It(`Invoke GetAccountRulesetVersion successfully with retries`, func() {
				rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(rulesetsService).ToNot(BeNil())
				rulesetsService.EnableRetries(0, 0)

				// Construct an instance of the GetAccountRulesetVersionOptions model
				getAccountRulesetVersionOptionsModel := new(rulesetsv1.GetAccountRulesetVersionOptions)
				getAccountRulesetVersionOptionsModel.RulesetID = core.StringPtr("testString")
				getAccountRulesetVersionOptionsModel.RulesetVersion = core.StringPtr("1")
				getAccountRulesetVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := rulesetsService.GetAccountRulesetVersionWithContext(ctx, getAccountRulesetVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				rulesetsService.DisableRetries()
				result, response, operationErr := rulesetsService.GetAccountRulesetVersion(getAccountRulesetVersionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = rulesetsService.GetAccountRulesetVersionWithContext(ctx, getAccountRulesetVersionOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getAccountRulesetVersionPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "messages": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "result": {"description": "Custom account ruleset", "id": "ID", "kind": "managed", "last_updated": "2000-01-01T00:00:00.000000Z", "name": "Name", "phase": "ddos_l4", "version": "1", "rules": [{"id": "ID", "version": "Version", "action": "Action", "action_parameters": {"id": "ID", "overrides": {"action": "Action", "enabled": false, "sensitivity_level": "high", "rules": [{"id": "ID", "enabled": false, "action": "Action", "sensitivity_level": "high"}], "categories": [{"category": "Category", "enabled": false, "action": "Action"}]}, "version": "Version", "ruleset": "Ruleset", "rulesets": ["Rulesets"], "response": {"content": "{\"success\": false, \"error\": \"you have been blocked\"}", "content_type": "application/json", "status_code": 400}}, "expression": "ip.src ne 1.1.1.1", "ref": "my_ref", "logging": {"enabled": true}, "last_updated": "2000-01-01T00:00:00.000000Z"}]}}`)
				}))
			})
			It(`Invoke GetAccountRulesetVersion successfully`, func() {
				rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(rulesetsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := rulesetsService.GetAccountRulesetVersion(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetAccountRulesetVersionOptions model
				getAccountRulesetVersionOptionsModel := new(rulesetsv1.GetAccountRulesetVersionOptions)
				getAccountRulesetVersionOptionsModel.RulesetID = core.StringPtr("testString")
				getAccountRulesetVersionOptionsModel.RulesetVersion = core.StringPtr("1")
				getAccountRulesetVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = rulesetsService.GetAccountRulesetVersion(getAccountRulesetVersionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetAccountRulesetVersion with error: Operation validation and request error`, func() {
				rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(rulesetsService).ToNot(BeNil())

				// Construct an instance of the GetAccountRulesetVersionOptions model
				getAccountRulesetVersionOptionsModel := new(rulesetsv1.GetAccountRulesetVersionOptions)
				getAccountRulesetVersionOptionsModel.RulesetID = core.StringPtr("testString")
				getAccountRulesetVersionOptionsModel.RulesetVersion = core.StringPtr("1")
				getAccountRulesetVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := rulesetsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := rulesetsService.GetAccountRulesetVersion(getAccountRulesetVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetAccountRulesetVersionOptions model with no property values
				getAccountRulesetVersionOptionsModelNew := new(rulesetsv1.GetAccountRulesetVersionOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = rulesetsService.GetAccountRulesetVersion(getAccountRulesetVersionOptionsModelNew)
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
			It(`Invoke GetAccountRulesetVersion successfully`, func() {
				rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(rulesetsService).ToNot(BeNil())

				// Construct an instance of the GetAccountRulesetVersionOptions model
				getAccountRulesetVersionOptionsModel := new(rulesetsv1.GetAccountRulesetVersionOptions)
				getAccountRulesetVersionOptionsModel.RulesetID = core.StringPtr("testString")
				getAccountRulesetVersionOptionsModel.RulesetVersion = core.StringPtr("1")
				getAccountRulesetVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := rulesetsService.GetAccountRulesetVersion(getAccountRulesetVersionOptionsModel)
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
	Describe(`DeleteAccountRulesetVersion(deleteAccountRulesetVersionOptions *DeleteAccountRulesetVersionOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		deleteAccountRulesetVersionPath := "/v1/testString/rulesets/testString/versions/1"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteAccountRulesetVersionPath))
					Expect(req.Method).To(Equal("DELETE"))

					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteAccountRulesetVersion successfully`, func() {
				rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(rulesetsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := rulesetsService.DeleteAccountRulesetVersion(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteAccountRulesetVersionOptions model
				deleteAccountRulesetVersionOptionsModel := new(rulesetsv1.DeleteAccountRulesetVersionOptions)
				deleteAccountRulesetVersionOptionsModel.RulesetID = core.StringPtr("testString")
				deleteAccountRulesetVersionOptionsModel.RulesetVersion = core.StringPtr("1")
				deleteAccountRulesetVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = rulesetsService.DeleteAccountRulesetVersion(deleteAccountRulesetVersionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteAccountRulesetVersion with error: Operation validation and request error`, func() {
				rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(rulesetsService).ToNot(BeNil())

				// Construct an instance of the DeleteAccountRulesetVersionOptions model
				deleteAccountRulesetVersionOptionsModel := new(rulesetsv1.DeleteAccountRulesetVersionOptions)
				deleteAccountRulesetVersionOptionsModel.RulesetID = core.StringPtr("testString")
				deleteAccountRulesetVersionOptionsModel.RulesetVersion = core.StringPtr("1")
				deleteAccountRulesetVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := rulesetsService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := rulesetsService.DeleteAccountRulesetVersion(deleteAccountRulesetVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteAccountRulesetVersionOptions model with no property values
				deleteAccountRulesetVersionOptionsModelNew := new(rulesetsv1.DeleteAccountRulesetVersionOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = rulesetsService.DeleteAccountRulesetVersion(deleteAccountRulesetVersionOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetAccountEntrypointRuleset(getAccountEntrypointRulesetOptions *GetAccountEntrypointRulesetOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getAccountEntrypointRulesetPath := "/v1/testString/rulesets/phases/ddos_l4/entrypoint"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getAccountEntrypointRulesetPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetAccountEntrypointRuleset with error: Operation response processing error`, func() {
				rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(rulesetsService).ToNot(BeNil())

				// Construct an instance of the GetAccountEntrypointRulesetOptions model
				getAccountEntrypointRulesetOptionsModel := new(rulesetsv1.GetAccountEntrypointRulesetOptions)
				getAccountEntrypointRulesetOptionsModel.RulesetPhase = core.StringPtr("ddos_l4")
				getAccountEntrypointRulesetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := rulesetsService.GetAccountEntrypointRuleset(getAccountEntrypointRulesetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				rulesetsService.EnableRetries(0, 0)
				result, response, operationErr = rulesetsService.GetAccountEntrypointRuleset(getAccountEntrypointRulesetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetAccountEntrypointRuleset(getAccountEntrypointRulesetOptions *GetAccountEntrypointRulesetOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getAccountEntrypointRulesetPath := "/v1/testString/rulesets/phases/ddos_l4/entrypoint"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getAccountEntrypointRulesetPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "messages": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "result": {"description": "Custom account ruleset", "id": "ID", "kind": "managed", "last_updated": "2000-01-01T00:00:00.000000Z", "name": "Name", "phase": "ddos_l4", "version": "1", "rules": [{"id": "ID", "version": "Version", "action": "Action", "action_parameters": {"id": "ID", "overrides": {"action": "Action", "enabled": false, "sensitivity_level": "high", "rules": [{"id": "ID", "enabled": false, "action": "Action", "sensitivity_level": "high"}], "categories": [{"category": "Category", "enabled": false, "action": "Action"}]}, "version": "Version", "ruleset": "Ruleset", "rulesets": ["Rulesets"], "response": {"content": "{\"success\": false, \"error\": \"you have been blocked\"}", "content_type": "application/json", "status_code": 400}}, "expression": "ip.src ne 1.1.1.1", "ref": "my_ref", "logging": {"enabled": true}, "last_updated": "2000-01-01T00:00:00.000000Z"}]}}`)
				}))
			})
			It(`Invoke GetAccountEntrypointRuleset successfully with retries`, func() {
				rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(rulesetsService).ToNot(BeNil())
				rulesetsService.EnableRetries(0, 0)

				// Construct an instance of the GetAccountEntrypointRulesetOptions model
				getAccountEntrypointRulesetOptionsModel := new(rulesetsv1.GetAccountEntrypointRulesetOptions)
				getAccountEntrypointRulesetOptionsModel.RulesetPhase = core.StringPtr("ddos_l4")
				getAccountEntrypointRulesetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := rulesetsService.GetAccountEntrypointRulesetWithContext(ctx, getAccountEntrypointRulesetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				rulesetsService.DisableRetries()
				result, response, operationErr := rulesetsService.GetAccountEntrypointRuleset(getAccountEntrypointRulesetOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = rulesetsService.GetAccountEntrypointRulesetWithContext(ctx, getAccountEntrypointRulesetOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getAccountEntrypointRulesetPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "messages": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "result": {"description": "Custom account ruleset", "id": "ID", "kind": "managed", "last_updated": "2000-01-01T00:00:00.000000Z", "name": "Name", "phase": "ddos_l4", "version": "1", "rules": [{"id": "ID", "version": "Version", "action": "Action", "action_parameters": {"id": "ID", "overrides": {"action": "Action", "enabled": false, "sensitivity_level": "high", "rules": [{"id": "ID", "enabled": false, "action": "Action", "sensitivity_level": "high"}], "categories": [{"category": "Category", "enabled": false, "action": "Action"}]}, "version": "Version", "ruleset": "Ruleset", "rulesets": ["Rulesets"], "response": {"content": "{\"success\": false, \"error\": \"you have been blocked\"}", "content_type": "application/json", "status_code": 400}}, "expression": "ip.src ne 1.1.1.1", "ref": "my_ref", "logging": {"enabled": true}, "last_updated": "2000-01-01T00:00:00.000000Z"}]}}`)
				}))
			})
			It(`Invoke GetAccountEntrypointRuleset successfully`, func() {
				rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(rulesetsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := rulesetsService.GetAccountEntrypointRuleset(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetAccountEntrypointRulesetOptions model
				getAccountEntrypointRulesetOptionsModel := new(rulesetsv1.GetAccountEntrypointRulesetOptions)
				getAccountEntrypointRulesetOptionsModel.RulesetPhase = core.StringPtr("ddos_l4")
				getAccountEntrypointRulesetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = rulesetsService.GetAccountEntrypointRuleset(getAccountEntrypointRulesetOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetAccountEntrypointRuleset with error: Operation validation and request error`, func() {
				rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(rulesetsService).ToNot(BeNil())

				// Construct an instance of the GetAccountEntrypointRulesetOptions model
				getAccountEntrypointRulesetOptionsModel := new(rulesetsv1.GetAccountEntrypointRulesetOptions)
				getAccountEntrypointRulesetOptionsModel.RulesetPhase = core.StringPtr("ddos_l4")
				getAccountEntrypointRulesetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := rulesetsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := rulesetsService.GetAccountEntrypointRuleset(getAccountEntrypointRulesetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetAccountEntrypointRulesetOptions model with no property values
				getAccountEntrypointRulesetOptionsModelNew := new(rulesetsv1.GetAccountEntrypointRulesetOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = rulesetsService.GetAccountEntrypointRuleset(getAccountEntrypointRulesetOptionsModelNew)
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
			It(`Invoke GetAccountEntrypointRuleset successfully`, func() {
				rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(rulesetsService).ToNot(BeNil())

				// Construct an instance of the GetAccountEntrypointRulesetOptions model
				getAccountEntrypointRulesetOptionsModel := new(rulesetsv1.GetAccountEntrypointRulesetOptions)
				getAccountEntrypointRulesetOptionsModel.RulesetPhase = core.StringPtr("ddos_l4")
				getAccountEntrypointRulesetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := rulesetsService.GetAccountEntrypointRuleset(getAccountEntrypointRulesetOptionsModel)
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
	Describe(`UpdateAccountEntrypointRuleset(updateAccountEntrypointRulesetOptions *UpdateAccountEntrypointRulesetOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		updateAccountEntrypointRulesetPath := "/v1/testString/rulesets/phases/ddos_l4/entrypoint"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateAccountEntrypointRulesetPath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateAccountEntrypointRuleset with error: Operation response processing error`, func() {
				rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(rulesetsService).ToNot(BeNil())

				// Construct an instance of the RulesOverride model
				rulesOverrideModel := new(rulesetsv1.RulesOverride)
				rulesOverrideModel.ID = core.StringPtr("testString")
				rulesOverrideModel.Enabled = core.BoolPtr(true)
				rulesOverrideModel.Action = core.StringPtr("testString")
				rulesOverrideModel.SensitivityLevel = core.StringPtr("high")

				// Construct an instance of the CategoriesOverride model
				categoriesOverrideModel := new(rulesetsv1.CategoriesOverride)
				categoriesOverrideModel.Category = core.StringPtr("testString")
				categoriesOverrideModel.Enabled = core.BoolPtr(true)
				categoriesOverrideModel.Action = core.StringPtr("testString")

				// Construct an instance of the Overrides model
				overridesModel := new(rulesetsv1.Overrides)
				overridesModel.Action = core.StringPtr("testString")
				overridesModel.Enabled = core.BoolPtr(true)
				overridesModel.SensitivityLevel = core.StringPtr("high")
				overridesModel.Rules = []rulesetsv1.RulesOverride{*rulesOverrideModel}
				overridesModel.Categories = []rulesetsv1.CategoriesOverride{*categoriesOverrideModel}

				// Construct an instance of the ActionParametersResponse model
				actionParametersResponseModel := new(rulesetsv1.ActionParametersResponse)
				actionParametersResponseModel.Content = core.StringPtr(`{"success": false, "error": "you have been blocked"}`)
				actionParametersResponseModel.ContentType = core.StringPtr("application/json")
				actionParametersResponseModel.StatusCode = core.Int64Ptr(int64(400))

				// Construct an instance of the ActionParameters model
				actionParametersModel := new(rulesetsv1.ActionParameters)
				actionParametersModel.ID = core.StringPtr("testString")
				actionParametersModel.Overrides = overridesModel
				actionParametersModel.Version = core.StringPtr("testString")
				actionParametersModel.Ruleset = core.StringPtr("testString")
				actionParametersModel.Rulesets = []string{"testString"}
				actionParametersModel.Response = actionParametersResponseModel

				// Construct an instance of the Logging model
				loggingModel := new(rulesetsv1.Logging)
				loggingModel.Enabled = core.BoolPtr(true)

				// Construct an instance of the Position model
				positionModel := new(rulesetsv1.Position)
				positionModel.Before = core.StringPtr("testString")
				positionModel.After = core.StringPtr("testString")
				positionModel.Index = core.Int64Ptr(int64(0))

				// Construct an instance of the RuleCreate model
				ruleCreateModel := new(rulesetsv1.RuleCreate)
				ruleCreateModel.Action = core.StringPtr("testString")
				ruleCreateModel.ActionParameters = actionParametersModel
				ruleCreateModel.Description = core.StringPtr("testString")
				ruleCreateModel.Enabled = core.BoolPtr(true)
				ruleCreateModel.Expression = core.StringPtr("ip.src ne 1.1.1.1")
				ruleCreateModel.ID = core.StringPtr("testString")
				ruleCreateModel.Logging = loggingModel
				ruleCreateModel.Ref = core.StringPtr("my_ref")
				ruleCreateModel.Position = positionModel

				// Construct an instance of the UpdateAccountEntrypointRulesetOptions model
				updateAccountEntrypointRulesetOptionsModel := new(rulesetsv1.UpdateAccountEntrypointRulesetOptions)
				updateAccountEntrypointRulesetOptionsModel.RulesetPhase = core.StringPtr("ddos_l4")
				updateAccountEntrypointRulesetOptionsModel.Description = core.StringPtr("Custom account ruleset")
				updateAccountEntrypointRulesetOptionsModel.Kind = core.StringPtr("managed")
				updateAccountEntrypointRulesetOptionsModel.Name = core.StringPtr("testString")
				updateAccountEntrypointRulesetOptionsModel.Phase = core.StringPtr("ddos_l4")
				updateAccountEntrypointRulesetOptionsModel.Rules = []rulesetsv1.RuleCreate{*ruleCreateModel}
				updateAccountEntrypointRulesetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := rulesetsService.UpdateAccountEntrypointRuleset(updateAccountEntrypointRulesetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				rulesetsService.EnableRetries(0, 0)
				result, response, operationErr = rulesetsService.UpdateAccountEntrypointRuleset(updateAccountEntrypointRulesetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateAccountEntrypointRuleset(updateAccountEntrypointRulesetOptions *UpdateAccountEntrypointRulesetOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		updateAccountEntrypointRulesetPath := "/v1/testString/rulesets/phases/ddos_l4/entrypoint"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateAccountEntrypointRulesetPath))
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
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "messages": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "result": {"description": "Custom account ruleset", "id": "ID", "kind": "managed", "last_updated": "2000-01-01T00:00:00.000000Z", "name": "Name", "phase": "ddos_l4", "version": "1", "rules": [{"id": "ID", "version": "Version", "action": "Action", "action_parameters": {"id": "ID", "overrides": {"action": "Action", "enabled": false, "sensitivity_level": "high", "rules": [{"id": "ID", "enabled": false, "action": "Action", "sensitivity_level": "high"}], "categories": [{"category": "Category", "enabled": false, "action": "Action"}]}, "version": "Version", "ruleset": "Ruleset", "rulesets": ["Rulesets"], "response": {"content": "{\"success\": false, \"error\": \"you have been blocked\"}", "content_type": "application/json", "status_code": 400}}, "expression": "ip.src ne 1.1.1.1", "ref": "my_ref", "logging": {"enabled": true}, "last_updated": "2000-01-01T00:00:00.000000Z"}]}}`)
				}))
			})
			It(`Invoke UpdateAccountEntrypointRuleset successfully with retries`, func() {
				rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(rulesetsService).ToNot(BeNil())
				rulesetsService.EnableRetries(0, 0)

				// Construct an instance of the RulesOverride model
				rulesOverrideModel := new(rulesetsv1.RulesOverride)
				rulesOverrideModel.ID = core.StringPtr("testString")
				rulesOverrideModel.Enabled = core.BoolPtr(true)
				rulesOverrideModel.Action = core.StringPtr("testString")
				rulesOverrideModel.SensitivityLevel = core.StringPtr("high")

				// Construct an instance of the CategoriesOverride model
				categoriesOverrideModel := new(rulesetsv1.CategoriesOverride)
				categoriesOverrideModel.Category = core.StringPtr("testString")
				categoriesOverrideModel.Enabled = core.BoolPtr(true)
				categoriesOverrideModel.Action = core.StringPtr("testString")

				// Construct an instance of the Overrides model
				overridesModel := new(rulesetsv1.Overrides)
				overridesModel.Action = core.StringPtr("testString")
				overridesModel.Enabled = core.BoolPtr(true)
				overridesModel.SensitivityLevel = core.StringPtr("high")
				overridesModel.Rules = []rulesetsv1.RulesOverride{*rulesOverrideModel}
				overridesModel.Categories = []rulesetsv1.CategoriesOverride{*categoriesOverrideModel}

				// Construct an instance of the ActionParametersResponse model
				actionParametersResponseModel := new(rulesetsv1.ActionParametersResponse)
				actionParametersResponseModel.Content = core.StringPtr(`{"success": false, "error": "you have been blocked"}`)
				actionParametersResponseModel.ContentType = core.StringPtr("application/json")
				actionParametersResponseModel.StatusCode = core.Int64Ptr(int64(400))

				// Construct an instance of the ActionParameters model
				actionParametersModel := new(rulesetsv1.ActionParameters)
				actionParametersModel.ID = core.StringPtr("testString")
				actionParametersModel.Overrides = overridesModel
				actionParametersModel.Version = core.StringPtr("testString")
				actionParametersModel.Ruleset = core.StringPtr("testString")
				actionParametersModel.Rulesets = []string{"testString"}
				actionParametersModel.Response = actionParametersResponseModel

				// Construct an instance of the Logging model
				loggingModel := new(rulesetsv1.Logging)
				loggingModel.Enabled = core.BoolPtr(true)

				// Construct an instance of the Position model
				positionModel := new(rulesetsv1.Position)
				positionModel.Before = core.StringPtr("testString")
				positionModel.After = core.StringPtr("testString")
				positionModel.Index = core.Int64Ptr(int64(0))

				// Construct an instance of the RuleCreate model
				ruleCreateModel := new(rulesetsv1.RuleCreate)
				ruleCreateModel.Action = core.StringPtr("testString")
				ruleCreateModel.ActionParameters = actionParametersModel
				ruleCreateModel.Description = core.StringPtr("testString")
				ruleCreateModel.Enabled = core.BoolPtr(true)
				ruleCreateModel.Expression = core.StringPtr("ip.src ne 1.1.1.1")
				ruleCreateModel.ID = core.StringPtr("testString")
				ruleCreateModel.Logging = loggingModel
				ruleCreateModel.Ref = core.StringPtr("my_ref")
				ruleCreateModel.Position = positionModel

				// Construct an instance of the UpdateAccountEntrypointRulesetOptions model
				updateAccountEntrypointRulesetOptionsModel := new(rulesetsv1.UpdateAccountEntrypointRulesetOptions)
				updateAccountEntrypointRulesetOptionsModel.RulesetPhase = core.StringPtr("ddos_l4")
				updateAccountEntrypointRulesetOptionsModel.Description = core.StringPtr("Custom account ruleset")
				updateAccountEntrypointRulesetOptionsModel.Kind = core.StringPtr("managed")
				updateAccountEntrypointRulesetOptionsModel.Name = core.StringPtr("testString")
				updateAccountEntrypointRulesetOptionsModel.Phase = core.StringPtr("ddos_l4")
				updateAccountEntrypointRulesetOptionsModel.Rules = []rulesetsv1.RuleCreate{*ruleCreateModel}
				updateAccountEntrypointRulesetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := rulesetsService.UpdateAccountEntrypointRulesetWithContext(ctx, updateAccountEntrypointRulesetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				rulesetsService.DisableRetries()
				result, response, operationErr := rulesetsService.UpdateAccountEntrypointRuleset(updateAccountEntrypointRulesetOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = rulesetsService.UpdateAccountEntrypointRulesetWithContext(ctx, updateAccountEntrypointRulesetOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(updateAccountEntrypointRulesetPath))
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
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "messages": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "result": {"description": "Custom account ruleset", "id": "ID", "kind": "managed", "last_updated": "2000-01-01T00:00:00.000000Z", "name": "Name", "phase": "ddos_l4", "version": "1", "rules": [{"id": "ID", "version": "Version", "action": "Action", "action_parameters": {"id": "ID", "overrides": {"action": "Action", "enabled": false, "sensitivity_level": "high", "rules": [{"id": "ID", "enabled": false, "action": "Action", "sensitivity_level": "high"}], "categories": [{"category": "Category", "enabled": false, "action": "Action"}]}, "version": "Version", "ruleset": "Ruleset", "rulesets": ["Rulesets"], "response": {"content": "{\"success\": false, \"error\": \"you have been blocked\"}", "content_type": "application/json", "status_code": 400}}, "expression": "ip.src ne 1.1.1.1", "ref": "my_ref", "logging": {"enabled": true}, "last_updated": "2000-01-01T00:00:00.000000Z"}]}}`)
				}))
			})
			It(`Invoke UpdateAccountEntrypointRuleset successfully`, func() {
				rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(rulesetsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := rulesetsService.UpdateAccountEntrypointRuleset(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the RulesOverride model
				rulesOverrideModel := new(rulesetsv1.RulesOverride)
				rulesOverrideModel.ID = core.StringPtr("testString")
				rulesOverrideModel.Enabled = core.BoolPtr(true)
				rulesOverrideModel.Action = core.StringPtr("testString")
				rulesOverrideModel.SensitivityLevel = core.StringPtr("high")

				// Construct an instance of the CategoriesOverride model
				categoriesOverrideModel := new(rulesetsv1.CategoriesOverride)
				categoriesOverrideModel.Category = core.StringPtr("testString")
				categoriesOverrideModel.Enabled = core.BoolPtr(true)
				categoriesOverrideModel.Action = core.StringPtr("testString")

				// Construct an instance of the Overrides model
				overridesModel := new(rulesetsv1.Overrides)
				overridesModel.Action = core.StringPtr("testString")
				overridesModel.Enabled = core.BoolPtr(true)
				overridesModel.SensitivityLevel = core.StringPtr("high")
				overridesModel.Rules = []rulesetsv1.RulesOverride{*rulesOverrideModel}
				overridesModel.Categories = []rulesetsv1.CategoriesOverride{*categoriesOverrideModel}

				// Construct an instance of the ActionParametersResponse model
				actionParametersResponseModel := new(rulesetsv1.ActionParametersResponse)
				actionParametersResponseModel.Content = core.StringPtr(`{"success": false, "error": "you have been blocked"}`)
				actionParametersResponseModel.ContentType = core.StringPtr("application/json")
				actionParametersResponseModel.StatusCode = core.Int64Ptr(int64(400))

				// Construct an instance of the ActionParameters model
				actionParametersModel := new(rulesetsv1.ActionParameters)
				actionParametersModel.ID = core.StringPtr("testString")
				actionParametersModel.Overrides = overridesModel
				actionParametersModel.Version = core.StringPtr("testString")
				actionParametersModel.Ruleset = core.StringPtr("testString")
				actionParametersModel.Rulesets = []string{"testString"}
				actionParametersModel.Response = actionParametersResponseModel

				// Construct an instance of the Logging model
				loggingModel := new(rulesetsv1.Logging)
				loggingModel.Enabled = core.BoolPtr(true)

				// Construct an instance of the Position model
				positionModel := new(rulesetsv1.Position)
				positionModel.Before = core.StringPtr("testString")
				positionModel.After = core.StringPtr("testString")
				positionModel.Index = core.Int64Ptr(int64(0))

				// Construct an instance of the RuleCreate model
				ruleCreateModel := new(rulesetsv1.RuleCreate)
				ruleCreateModel.Action = core.StringPtr("testString")
				ruleCreateModel.ActionParameters = actionParametersModel
				ruleCreateModel.Description = core.StringPtr("testString")
				ruleCreateModel.Enabled = core.BoolPtr(true)
				ruleCreateModel.Expression = core.StringPtr("ip.src ne 1.1.1.1")
				ruleCreateModel.ID = core.StringPtr("testString")
				ruleCreateModel.Logging = loggingModel
				ruleCreateModel.Ref = core.StringPtr("my_ref")
				ruleCreateModel.Position = positionModel

				// Construct an instance of the UpdateAccountEntrypointRulesetOptions model
				updateAccountEntrypointRulesetOptionsModel := new(rulesetsv1.UpdateAccountEntrypointRulesetOptions)
				updateAccountEntrypointRulesetOptionsModel.RulesetPhase = core.StringPtr("ddos_l4")
				updateAccountEntrypointRulesetOptionsModel.Description = core.StringPtr("Custom account ruleset")
				updateAccountEntrypointRulesetOptionsModel.Kind = core.StringPtr("managed")
				updateAccountEntrypointRulesetOptionsModel.Name = core.StringPtr("testString")
				updateAccountEntrypointRulesetOptionsModel.Phase = core.StringPtr("ddos_l4")
				updateAccountEntrypointRulesetOptionsModel.Rules = []rulesetsv1.RuleCreate{*ruleCreateModel}
				updateAccountEntrypointRulesetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = rulesetsService.UpdateAccountEntrypointRuleset(updateAccountEntrypointRulesetOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateAccountEntrypointRuleset with error: Operation validation and request error`, func() {
				rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(rulesetsService).ToNot(BeNil())

				// Construct an instance of the RulesOverride model
				rulesOverrideModel := new(rulesetsv1.RulesOverride)
				rulesOverrideModel.ID = core.StringPtr("testString")
				rulesOverrideModel.Enabled = core.BoolPtr(true)
				rulesOverrideModel.Action = core.StringPtr("testString")
				rulesOverrideModel.SensitivityLevel = core.StringPtr("high")

				// Construct an instance of the CategoriesOverride model
				categoriesOverrideModel := new(rulesetsv1.CategoriesOverride)
				categoriesOverrideModel.Category = core.StringPtr("testString")
				categoriesOverrideModel.Enabled = core.BoolPtr(true)
				categoriesOverrideModel.Action = core.StringPtr("testString")

				// Construct an instance of the Overrides model
				overridesModel := new(rulesetsv1.Overrides)
				overridesModel.Action = core.StringPtr("testString")
				overridesModel.Enabled = core.BoolPtr(true)
				overridesModel.SensitivityLevel = core.StringPtr("high")
				overridesModel.Rules = []rulesetsv1.RulesOverride{*rulesOverrideModel}
				overridesModel.Categories = []rulesetsv1.CategoriesOverride{*categoriesOverrideModel}

				// Construct an instance of the ActionParametersResponse model
				actionParametersResponseModel := new(rulesetsv1.ActionParametersResponse)
				actionParametersResponseModel.Content = core.StringPtr(`{"success": false, "error": "you have been blocked"}`)
				actionParametersResponseModel.ContentType = core.StringPtr("application/json")
				actionParametersResponseModel.StatusCode = core.Int64Ptr(int64(400))

				// Construct an instance of the ActionParameters model
				actionParametersModel := new(rulesetsv1.ActionParameters)
				actionParametersModel.ID = core.StringPtr("testString")
				actionParametersModel.Overrides = overridesModel
				actionParametersModel.Version = core.StringPtr("testString")
				actionParametersModel.Ruleset = core.StringPtr("testString")
				actionParametersModel.Rulesets = []string{"testString"}
				actionParametersModel.Response = actionParametersResponseModel

				// Construct an instance of the Logging model
				loggingModel := new(rulesetsv1.Logging)
				loggingModel.Enabled = core.BoolPtr(true)

				// Construct an instance of the Position model
				positionModel := new(rulesetsv1.Position)
				positionModel.Before = core.StringPtr("testString")
				positionModel.After = core.StringPtr("testString")
				positionModel.Index = core.Int64Ptr(int64(0))

				// Construct an instance of the RuleCreate model
				ruleCreateModel := new(rulesetsv1.RuleCreate)
				ruleCreateModel.Action = core.StringPtr("testString")
				ruleCreateModel.ActionParameters = actionParametersModel
				ruleCreateModel.Description = core.StringPtr("testString")
				ruleCreateModel.Enabled = core.BoolPtr(true)
				ruleCreateModel.Expression = core.StringPtr("ip.src ne 1.1.1.1")
				ruleCreateModel.ID = core.StringPtr("testString")
				ruleCreateModel.Logging = loggingModel
				ruleCreateModel.Ref = core.StringPtr("my_ref")
				ruleCreateModel.Position = positionModel

				// Construct an instance of the UpdateAccountEntrypointRulesetOptions model
				updateAccountEntrypointRulesetOptionsModel := new(rulesetsv1.UpdateAccountEntrypointRulesetOptions)
				updateAccountEntrypointRulesetOptionsModel.RulesetPhase = core.StringPtr("ddos_l4")
				updateAccountEntrypointRulesetOptionsModel.Description = core.StringPtr("Custom account ruleset")
				updateAccountEntrypointRulesetOptionsModel.Kind = core.StringPtr("managed")
				updateAccountEntrypointRulesetOptionsModel.Name = core.StringPtr("testString")
				updateAccountEntrypointRulesetOptionsModel.Phase = core.StringPtr("ddos_l4")
				updateAccountEntrypointRulesetOptionsModel.Rules = []rulesetsv1.RuleCreate{*ruleCreateModel}
				updateAccountEntrypointRulesetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := rulesetsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := rulesetsService.UpdateAccountEntrypointRuleset(updateAccountEntrypointRulesetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateAccountEntrypointRulesetOptions model with no property values
				updateAccountEntrypointRulesetOptionsModelNew := new(rulesetsv1.UpdateAccountEntrypointRulesetOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = rulesetsService.UpdateAccountEntrypointRuleset(updateAccountEntrypointRulesetOptionsModelNew)
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
			It(`Invoke UpdateAccountEntrypointRuleset successfully`, func() {
				rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(rulesetsService).ToNot(BeNil())

				// Construct an instance of the RulesOverride model
				rulesOverrideModel := new(rulesetsv1.RulesOverride)
				rulesOverrideModel.ID = core.StringPtr("testString")
				rulesOverrideModel.Enabled = core.BoolPtr(true)
				rulesOverrideModel.Action = core.StringPtr("testString")
				rulesOverrideModel.SensitivityLevel = core.StringPtr("high")

				// Construct an instance of the CategoriesOverride model
				categoriesOverrideModel := new(rulesetsv1.CategoriesOverride)
				categoriesOverrideModel.Category = core.StringPtr("testString")
				categoriesOverrideModel.Enabled = core.BoolPtr(true)
				categoriesOverrideModel.Action = core.StringPtr("testString")

				// Construct an instance of the Overrides model
				overridesModel := new(rulesetsv1.Overrides)
				overridesModel.Action = core.StringPtr("testString")
				overridesModel.Enabled = core.BoolPtr(true)
				overridesModel.SensitivityLevel = core.StringPtr("high")
				overridesModel.Rules = []rulesetsv1.RulesOverride{*rulesOverrideModel}
				overridesModel.Categories = []rulesetsv1.CategoriesOverride{*categoriesOverrideModel}

				// Construct an instance of the ActionParametersResponse model
				actionParametersResponseModel := new(rulesetsv1.ActionParametersResponse)
				actionParametersResponseModel.Content = core.StringPtr(`{"success": false, "error": "you have been blocked"}`)
				actionParametersResponseModel.ContentType = core.StringPtr("application/json")
				actionParametersResponseModel.StatusCode = core.Int64Ptr(int64(400))

				// Construct an instance of the ActionParameters model
				actionParametersModel := new(rulesetsv1.ActionParameters)
				actionParametersModel.ID = core.StringPtr("testString")
				actionParametersModel.Overrides = overridesModel
				actionParametersModel.Version = core.StringPtr("testString")
				actionParametersModel.Ruleset = core.StringPtr("testString")
				actionParametersModel.Rulesets = []string{"testString"}
				actionParametersModel.Response = actionParametersResponseModel

				// Construct an instance of the Logging model
				loggingModel := new(rulesetsv1.Logging)
				loggingModel.Enabled = core.BoolPtr(true)

				// Construct an instance of the Position model
				positionModel := new(rulesetsv1.Position)
				positionModel.Before = core.StringPtr("testString")
				positionModel.After = core.StringPtr("testString")
				positionModel.Index = core.Int64Ptr(int64(0))

				// Construct an instance of the RuleCreate model
				ruleCreateModel := new(rulesetsv1.RuleCreate)
				ruleCreateModel.Action = core.StringPtr("testString")
				ruleCreateModel.ActionParameters = actionParametersModel
				ruleCreateModel.Description = core.StringPtr("testString")
				ruleCreateModel.Enabled = core.BoolPtr(true)
				ruleCreateModel.Expression = core.StringPtr("ip.src ne 1.1.1.1")
				ruleCreateModel.ID = core.StringPtr("testString")
				ruleCreateModel.Logging = loggingModel
				ruleCreateModel.Ref = core.StringPtr("my_ref")
				ruleCreateModel.Position = positionModel

				// Construct an instance of the UpdateAccountEntrypointRulesetOptions model
				updateAccountEntrypointRulesetOptionsModel := new(rulesetsv1.UpdateAccountEntrypointRulesetOptions)
				updateAccountEntrypointRulesetOptionsModel.RulesetPhase = core.StringPtr("ddos_l4")
				updateAccountEntrypointRulesetOptionsModel.Description = core.StringPtr("Custom account ruleset")
				updateAccountEntrypointRulesetOptionsModel.Kind = core.StringPtr("managed")
				updateAccountEntrypointRulesetOptionsModel.Name = core.StringPtr("testString")
				updateAccountEntrypointRulesetOptionsModel.Phase = core.StringPtr("ddos_l4")
				updateAccountEntrypointRulesetOptionsModel.Rules = []rulesetsv1.RuleCreate{*ruleCreateModel}
				updateAccountEntrypointRulesetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := rulesetsService.UpdateAccountEntrypointRuleset(updateAccountEntrypointRulesetOptionsModel)
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
	Describe(`GetAccountEntryPointRulesetVersions(getAccountEntryPointRulesetVersionsOptions *GetAccountEntryPointRulesetVersionsOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getAccountEntryPointRulesetVersionsPath := "/v1/testString/rulesets/phases/ddos_l4/entrypoint/versions"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getAccountEntryPointRulesetVersionsPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetAccountEntryPointRulesetVersions with error: Operation response processing error`, func() {
				rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(rulesetsService).ToNot(BeNil())

				// Construct an instance of the GetAccountEntryPointRulesetVersionsOptions model
				getAccountEntryPointRulesetVersionsOptionsModel := new(rulesetsv1.GetAccountEntryPointRulesetVersionsOptions)
				getAccountEntryPointRulesetVersionsOptionsModel.RulesetPhase = core.StringPtr("ddos_l4")
				getAccountEntryPointRulesetVersionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := rulesetsService.GetAccountEntryPointRulesetVersions(getAccountEntryPointRulesetVersionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				rulesetsService.EnableRetries(0, 0)
				result, response, operationErr = rulesetsService.GetAccountEntryPointRulesetVersions(getAccountEntryPointRulesetVersionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetAccountEntryPointRulesetVersions(getAccountEntryPointRulesetVersionsOptions *GetAccountEntryPointRulesetVersionsOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getAccountEntryPointRulesetVersionsPath := "/v1/testString/rulesets/phases/ddos_l4/entrypoint/versions"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getAccountEntryPointRulesetVersionsPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "messages": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "result": [{"description": "Custom account ruleset", "id": "ID", "kind": "managed", "last_updated": "2000-01-01T00:00:00.000000Z", "name": "Name", "phase": "ddos_l4", "version": "1"}]}`)
				}))
			})
			It(`Invoke GetAccountEntryPointRulesetVersions successfully with retries`, func() {
				rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(rulesetsService).ToNot(BeNil())
				rulesetsService.EnableRetries(0, 0)

				// Construct an instance of the GetAccountEntryPointRulesetVersionsOptions model
				getAccountEntryPointRulesetVersionsOptionsModel := new(rulesetsv1.GetAccountEntryPointRulesetVersionsOptions)
				getAccountEntryPointRulesetVersionsOptionsModel.RulesetPhase = core.StringPtr("ddos_l4")
				getAccountEntryPointRulesetVersionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := rulesetsService.GetAccountEntryPointRulesetVersionsWithContext(ctx, getAccountEntryPointRulesetVersionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				rulesetsService.DisableRetries()
				result, response, operationErr := rulesetsService.GetAccountEntryPointRulesetVersions(getAccountEntryPointRulesetVersionsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = rulesetsService.GetAccountEntryPointRulesetVersionsWithContext(ctx, getAccountEntryPointRulesetVersionsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getAccountEntryPointRulesetVersionsPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "messages": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "result": [{"description": "Custom account ruleset", "id": "ID", "kind": "managed", "last_updated": "2000-01-01T00:00:00.000000Z", "name": "Name", "phase": "ddos_l4", "version": "1"}]}`)
				}))
			})
			It(`Invoke GetAccountEntryPointRulesetVersions successfully`, func() {
				rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(rulesetsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := rulesetsService.GetAccountEntryPointRulesetVersions(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetAccountEntryPointRulesetVersionsOptions model
				getAccountEntryPointRulesetVersionsOptionsModel := new(rulesetsv1.GetAccountEntryPointRulesetVersionsOptions)
				getAccountEntryPointRulesetVersionsOptionsModel.RulesetPhase = core.StringPtr("ddos_l4")
				getAccountEntryPointRulesetVersionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = rulesetsService.GetAccountEntryPointRulesetVersions(getAccountEntryPointRulesetVersionsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetAccountEntryPointRulesetVersions with error: Operation validation and request error`, func() {
				rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(rulesetsService).ToNot(BeNil())

				// Construct an instance of the GetAccountEntryPointRulesetVersionsOptions model
				getAccountEntryPointRulesetVersionsOptionsModel := new(rulesetsv1.GetAccountEntryPointRulesetVersionsOptions)
				getAccountEntryPointRulesetVersionsOptionsModel.RulesetPhase = core.StringPtr("ddos_l4")
				getAccountEntryPointRulesetVersionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := rulesetsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := rulesetsService.GetAccountEntryPointRulesetVersions(getAccountEntryPointRulesetVersionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetAccountEntryPointRulesetVersionsOptions model with no property values
				getAccountEntryPointRulesetVersionsOptionsModelNew := new(rulesetsv1.GetAccountEntryPointRulesetVersionsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = rulesetsService.GetAccountEntryPointRulesetVersions(getAccountEntryPointRulesetVersionsOptionsModelNew)
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
			It(`Invoke GetAccountEntryPointRulesetVersions successfully`, func() {
				rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(rulesetsService).ToNot(BeNil())

				// Construct an instance of the GetAccountEntryPointRulesetVersionsOptions model
				getAccountEntryPointRulesetVersionsOptionsModel := new(rulesetsv1.GetAccountEntryPointRulesetVersionsOptions)
				getAccountEntryPointRulesetVersionsOptionsModel.RulesetPhase = core.StringPtr("ddos_l4")
				getAccountEntryPointRulesetVersionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := rulesetsService.GetAccountEntryPointRulesetVersions(getAccountEntryPointRulesetVersionsOptionsModel)
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
	Describe(`GetAccountEntryPointRulesetVersion(getAccountEntryPointRulesetVersionOptions *GetAccountEntryPointRulesetVersionOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getAccountEntryPointRulesetVersionPath := "/v1/testString/rulesets/phases/ddos_l4/entrypoint/versions/1"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getAccountEntryPointRulesetVersionPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetAccountEntryPointRulesetVersion with error: Operation response processing error`, func() {
				rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(rulesetsService).ToNot(BeNil())

				// Construct an instance of the GetAccountEntryPointRulesetVersionOptions model
				getAccountEntryPointRulesetVersionOptionsModel := new(rulesetsv1.GetAccountEntryPointRulesetVersionOptions)
				getAccountEntryPointRulesetVersionOptionsModel.RulesetPhase = core.StringPtr("ddos_l4")
				getAccountEntryPointRulesetVersionOptionsModel.RulesetVersion = core.StringPtr("1")
				getAccountEntryPointRulesetVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := rulesetsService.GetAccountEntryPointRulesetVersion(getAccountEntryPointRulesetVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				rulesetsService.EnableRetries(0, 0)
				result, response, operationErr = rulesetsService.GetAccountEntryPointRulesetVersion(getAccountEntryPointRulesetVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetAccountEntryPointRulesetVersion(getAccountEntryPointRulesetVersionOptions *GetAccountEntryPointRulesetVersionOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getAccountEntryPointRulesetVersionPath := "/v1/testString/rulesets/phases/ddos_l4/entrypoint/versions/1"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getAccountEntryPointRulesetVersionPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "messages": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "result": {"description": "Custom account ruleset", "id": "ID", "kind": "managed", "last_updated": "2000-01-01T00:00:00.000000Z", "name": "Name", "phase": "ddos_l4", "version": "1", "rules": [{"id": "ID", "version": "Version", "action": "Action", "action_parameters": {"id": "ID", "overrides": {"action": "Action", "enabled": false, "sensitivity_level": "high", "rules": [{"id": "ID", "enabled": false, "action": "Action", "sensitivity_level": "high"}], "categories": [{"category": "Category", "enabled": false, "action": "Action"}]}, "version": "Version", "ruleset": "Ruleset", "rulesets": ["Rulesets"], "response": {"content": "{\"success\": false, \"error\": \"you have been blocked\"}", "content_type": "application/json", "status_code": 400}}, "expression": "ip.src ne 1.1.1.1", "ref": "my_ref", "logging": {"enabled": true}, "last_updated": "2000-01-01T00:00:00.000000Z"}]}}`)
				}))
			})
			It(`Invoke GetAccountEntryPointRulesetVersion successfully with retries`, func() {
				rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(rulesetsService).ToNot(BeNil())
				rulesetsService.EnableRetries(0, 0)

				// Construct an instance of the GetAccountEntryPointRulesetVersionOptions model
				getAccountEntryPointRulesetVersionOptionsModel := new(rulesetsv1.GetAccountEntryPointRulesetVersionOptions)
				getAccountEntryPointRulesetVersionOptionsModel.RulesetPhase = core.StringPtr("ddos_l4")
				getAccountEntryPointRulesetVersionOptionsModel.RulesetVersion = core.StringPtr("1")
				getAccountEntryPointRulesetVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := rulesetsService.GetAccountEntryPointRulesetVersionWithContext(ctx, getAccountEntryPointRulesetVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				rulesetsService.DisableRetries()
				result, response, operationErr := rulesetsService.GetAccountEntryPointRulesetVersion(getAccountEntryPointRulesetVersionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = rulesetsService.GetAccountEntryPointRulesetVersionWithContext(ctx, getAccountEntryPointRulesetVersionOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getAccountEntryPointRulesetVersionPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "messages": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "result": {"description": "Custom account ruleset", "id": "ID", "kind": "managed", "last_updated": "2000-01-01T00:00:00.000000Z", "name": "Name", "phase": "ddos_l4", "version": "1", "rules": [{"id": "ID", "version": "Version", "action": "Action", "action_parameters": {"id": "ID", "overrides": {"action": "Action", "enabled": false, "sensitivity_level": "high", "rules": [{"id": "ID", "enabled": false, "action": "Action", "sensitivity_level": "high"}], "categories": [{"category": "Category", "enabled": false, "action": "Action"}]}, "version": "Version", "ruleset": "Ruleset", "rulesets": ["Rulesets"], "response": {"content": "{\"success\": false, \"error\": \"you have been blocked\"}", "content_type": "application/json", "status_code": 400}}, "expression": "ip.src ne 1.1.1.1", "ref": "my_ref", "logging": {"enabled": true}, "last_updated": "2000-01-01T00:00:00.000000Z"}]}}`)
				}))
			})
			It(`Invoke GetAccountEntryPointRulesetVersion successfully`, func() {
				rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(rulesetsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := rulesetsService.GetAccountEntryPointRulesetVersion(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetAccountEntryPointRulesetVersionOptions model
				getAccountEntryPointRulesetVersionOptionsModel := new(rulesetsv1.GetAccountEntryPointRulesetVersionOptions)
				getAccountEntryPointRulesetVersionOptionsModel.RulesetPhase = core.StringPtr("ddos_l4")
				getAccountEntryPointRulesetVersionOptionsModel.RulesetVersion = core.StringPtr("1")
				getAccountEntryPointRulesetVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = rulesetsService.GetAccountEntryPointRulesetVersion(getAccountEntryPointRulesetVersionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetAccountEntryPointRulesetVersion with error: Operation validation and request error`, func() {
				rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(rulesetsService).ToNot(BeNil())

				// Construct an instance of the GetAccountEntryPointRulesetVersionOptions model
				getAccountEntryPointRulesetVersionOptionsModel := new(rulesetsv1.GetAccountEntryPointRulesetVersionOptions)
				getAccountEntryPointRulesetVersionOptionsModel.RulesetPhase = core.StringPtr("ddos_l4")
				getAccountEntryPointRulesetVersionOptionsModel.RulesetVersion = core.StringPtr("1")
				getAccountEntryPointRulesetVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := rulesetsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := rulesetsService.GetAccountEntryPointRulesetVersion(getAccountEntryPointRulesetVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetAccountEntryPointRulesetVersionOptions model with no property values
				getAccountEntryPointRulesetVersionOptionsModelNew := new(rulesetsv1.GetAccountEntryPointRulesetVersionOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = rulesetsService.GetAccountEntryPointRulesetVersion(getAccountEntryPointRulesetVersionOptionsModelNew)
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
			It(`Invoke GetAccountEntryPointRulesetVersion successfully`, func() {
				rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(rulesetsService).ToNot(BeNil())

				// Construct an instance of the GetAccountEntryPointRulesetVersionOptions model
				getAccountEntryPointRulesetVersionOptionsModel := new(rulesetsv1.GetAccountEntryPointRulesetVersionOptions)
				getAccountEntryPointRulesetVersionOptionsModel.RulesetPhase = core.StringPtr("ddos_l4")
				getAccountEntryPointRulesetVersionOptionsModel.RulesetVersion = core.StringPtr("1")
				getAccountEntryPointRulesetVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := rulesetsService.GetAccountEntryPointRulesetVersion(getAccountEntryPointRulesetVersionOptionsModel)
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
	Describe(`CreateAccountRulesetRule(createAccountRulesetRuleOptions *CreateAccountRulesetRuleOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		createAccountRulesetRulePath := "/v1/testString/rulesets/testString/rules"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createAccountRulesetRulePath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateAccountRulesetRule with error: Operation response processing error`, func() {
				rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(rulesetsService).ToNot(BeNil())

				// Construct an instance of the RulesOverride model
				rulesOverrideModel := new(rulesetsv1.RulesOverride)
				rulesOverrideModel.ID = core.StringPtr("testString")
				rulesOverrideModel.Enabled = core.BoolPtr(true)
				rulesOverrideModel.Action = core.StringPtr("testString")
				rulesOverrideModel.SensitivityLevel = core.StringPtr("high")

				// Construct an instance of the CategoriesOverride model
				categoriesOverrideModel := new(rulesetsv1.CategoriesOverride)
				categoriesOverrideModel.Category = core.StringPtr("testString")
				categoriesOverrideModel.Enabled = core.BoolPtr(true)
				categoriesOverrideModel.Action = core.StringPtr("testString")

				// Construct an instance of the Overrides model
				overridesModel := new(rulesetsv1.Overrides)
				overridesModel.Action = core.StringPtr("testString")
				overridesModel.Enabled = core.BoolPtr(true)
				overridesModel.SensitivityLevel = core.StringPtr("high")
				overridesModel.Rules = []rulesetsv1.RulesOverride{*rulesOverrideModel}
				overridesModel.Categories = []rulesetsv1.CategoriesOverride{*categoriesOverrideModel}

				// Construct an instance of the ActionParametersResponse model
				actionParametersResponseModel := new(rulesetsv1.ActionParametersResponse)
				actionParametersResponseModel.Content = core.StringPtr(`{"success": false, "error": "you have been blocked"}`)
				actionParametersResponseModel.ContentType = core.StringPtr("application/json")
				actionParametersResponseModel.StatusCode = core.Int64Ptr(int64(400))

				// Construct an instance of the ActionParameters model
				actionParametersModel := new(rulesetsv1.ActionParameters)
				actionParametersModel.ID = core.StringPtr("testString")
				actionParametersModel.Overrides = overridesModel
				actionParametersModel.Version = core.StringPtr("testString")
				actionParametersModel.Ruleset = core.StringPtr("testString")
				actionParametersModel.Rulesets = []string{"testString"}
				actionParametersModel.Response = actionParametersResponseModel

				// Construct an instance of the Logging model
				loggingModel := new(rulesetsv1.Logging)
				loggingModel.Enabled = core.BoolPtr(true)

				// Construct an instance of the Position model
				positionModel := new(rulesetsv1.Position)
				positionModel.Before = core.StringPtr("testString")
				positionModel.After = core.StringPtr("testString")
				positionModel.Index = core.Int64Ptr(int64(0))

				// Construct an instance of the CreateAccountRulesetRuleOptions model
				createAccountRulesetRuleOptionsModel := new(rulesetsv1.CreateAccountRulesetRuleOptions)
				createAccountRulesetRuleOptionsModel.RulesetID = core.StringPtr("testString")
				createAccountRulesetRuleOptionsModel.Action = core.StringPtr("testString")
				createAccountRulesetRuleOptionsModel.ActionParameters = actionParametersModel
				createAccountRulesetRuleOptionsModel.Description = core.StringPtr("testString")
				createAccountRulesetRuleOptionsModel.Enabled = core.BoolPtr(true)
				createAccountRulesetRuleOptionsModel.Expression = core.StringPtr("ip.src ne 1.1.1.1")
				createAccountRulesetRuleOptionsModel.ID = core.StringPtr("testString")
				createAccountRulesetRuleOptionsModel.Logging = loggingModel
				createAccountRulesetRuleOptionsModel.Ref = core.StringPtr("my_ref")
				createAccountRulesetRuleOptionsModel.Position = positionModel
				createAccountRulesetRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := rulesetsService.CreateAccountRulesetRule(createAccountRulesetRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				rulesetsService.EnableRetries(0, 0)
				result, response, operationErr = rulesetsService.CreateAccountRulesetRule(createAccountRulesetRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateAccountRulesetRule(createAccountRulesetRuleOptions *CreateAccountRulesetRuleOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		createAccountRulesetRulePath := "/v1/testString/rulesets/testString/rules"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createAccountRulesetRulePath))
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
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "messages": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "result": {"id": "ID", "version": "Version", "action": "Action", "action_parameters": {"id": "ID", "overrides": {"action": "Action", "enabled": false, "sensitivity_level": "high", "rules": [{"id": "ID", "enabled": false, "action": "Action", "sensitivity_level": "high"}], "categories": [{"category": "Category", "enabled": false, "action": "Action"}]}, "version": "Version", "ruleset": "Ruleset", "rulesets": ["Rulesets"], "response": {"content": "{\"success\": false, \"error\": \"you have been blocked\"}", "content_type": "application/json", "status_code": 400}}, "expression": "ip.src ne 1.1.1.1", "ref": "my_ref", "logging": {"enabled": true}, "last_updated": "2000-01-01T00:00:00.000000Z"}}`)
				}))
			})
			It(`Invoke CreateAccountRulesetRule successfully with retries`, func() {
				rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(rulesetsService).ToNot(BeNil())
				rulesetsService.EnableRetries(0, 0)

				// Construct an instance of the RulesOverride model
				rulesOverrideModel := new(rulesetsv1.RulesOverride)
				rulesOverrideModel.ID = core.StringPtr("testString")
				rulesOverrideModel.Enabled = core.BoolPtr(true)
				rulesOverrideModel.Action = core.StringPtr("testString")
				rulesOverrideModel.SensitivityLevel = core.StringPtr("high")

				// Construct an instance of the CategoriesOverride model
				categoriesOverrideModel := new(rulesetsv1.CategoriesOverride)
				categoriesOverrideModel.Category = core.StringPtr("testString")
				categoriesOverrideModel.Enabled = core.BoolPtr(true)
				categoriesOverrideModel.Action = core.StringPtr("testString")

				// Construct an instance of the Overrides model
				overridesModel := new(rulesetsv1.Overrides)
				overridesModel.Action = core.StringPtr("testString")
				overridesModel.Enabled = core.BoolPtr(true)
				overridesModel.SensitivityLevel = core.StringPtr("high")
				overridesModel.Rules = []rulesetsv1.RulesOverride{*rulesOverrideModel}
				overridesModel.Categories = []rulesetsv1.CategoriesOverride{*categoriesOverrideModel}

				// Construct an instance of the ActionParametersResponse model
				actionParametersResponseModel := new(rulesetsv1.ActionParametersResponse)
				actionParametersResponseModel.Content = core.StringPtr(`{"success": false, "error": "you have been blocked"}`)
				actionParametersResponseModel.ContentType = core.StringPtr("application/json")
				actionParametersResponseModel.StatusCode = core.Int64Ptr(int64(400))

				// Construct an instance of the ActionParameters model
				actionParametersModel := new(rulesetsv1.ActionParameters)
				actionParametersModel.ID = core.StringPtr("testString")
				actionParametersModel.Overrides = overridesModel
				actionParametersModel.Version = core.StringPtr("testString")
				actionParametersModel.Ruleset = core.StringPtr("testString")
				actionParametersModel.Rulesets = []string{"testString"}
				actionParametersModel.Response = actionParametersResponseModel

				// Construct an instance of the Logging model
				loggingModel := new(rulesetsv1.Logging)
				loggingModel.Enabled = core.BoolPtr(true)

				// Construct an instance of the Position model
				positionModel := new(rulesetsv1.Position)
				positionModel.Before = core.StringPtr("testString")
				positionModel.After = core.StringPtr("testString")
				positionModel.Index = core.Int64Ptr(int64(0))

				// Construct an instance of the CreateAccountRulesetRuleOptions model
				createAccountRulesetRuleOptionsModel := new(rulesetsv1.CreateAccountRulesetRuleOptions)
				createAccountRulesetRuleOptionsModel.RulesetID = core.StringPtr("testString")
				createAccountRulesetRuleOptionsModel.Action = core.StringPtr("testString")
				createAccountRulesetRuleOptionsModel.ActionParameters = actionParametersModel
				createAccountRulesetRuleOptionsModel.Description = core.StringPtr("testString")
				createAccountRulesetRuleOptionsModel.Enabled = core.BoolPtr(true)
				createAccountRulesetRuleOptionsModel.Expression = core.StringPtr("ip.src ne 1.1.1.1")
				createAccountRulesetRuleOptionsModel.ID = core.StringPtr("testString")
				createAccountRulesetRuleOptionsModel.Logging = loggingModel
				createAccountRulesetRuleOptionsModel.Ref = core.StringPtr("my_ref")
				createAccountRulesetRuleOptionsModel.Position = positionModel
				createAccountRulesetRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := rulesetsService.CreateAccountRulesetRuleWithContext(ctx, createAccountRulesetRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				rulesetsService.DisableRetries()
				result, response, operationErr := rulesetsService.CreateAccountRulesetRule(createAccountRulesetRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = rulesetsService.CreateAccountRulesetRuleWithContext(ctx, createAccountRulesetRuleOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createAccountRulesetRulePath))
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
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "messages": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "result": {"id": "ID", "version": "Version", "action": "Action", "action_parameters": {"id": "ID", "overrides": {"action": "Action", "enabled": false, "sensitivity_level": "high", "rules": [{"id": "ID", "enabled": false, "action": "Action", "sensitivity_level": "high"}], "categories": [{"category": "Category", "enabled": false, "action": "Action"}]}, "version": "Version", "ruleset": "Ruleset", "rulesets": ["Rulesets"], "response": {"content": "{\"success\": false, \"error\": \"you have been blocked\"}", "content_type": "application/json", "status_code": 400}}, "expression": "ip.src ne 1.1.1.1", "ref": "my_ref", "logging": {"enabled": true}, "last_updated": "2000-01-01T00:00:00.000000Z"}}`)
				}))
			})
			It(`Invoke CreateAccountRulesetRule successfully`, func() {
				rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(rulesetsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := rulesetsService.CreateAccountRulesetRule(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the RulesOverride model
				rulesOverrideModel := new(rulesetsv1.RulesOverride)
				rulesOverrideModel.ID = core.StringPtr("testString")
				rulesOverrideModel.Enabled = core.BoolPtr(true)
				rulesOverrideModel.Action = core.StringPtr("testString")
				rulesOverrideModel.SensitivityLevel = core.StringPtr("high")

				// Construct an instance of the CategoriesOverride model
				categoriesOverrideModel := new(rulesetsv1.CategoriesOverride)
				categoriesOverrideModel.Category = core.StringPtr("testString")
				categoriesOverrideModel.Enabled = core.BoolPtr(true)
				categoriesOverrideModel.Action = core.StringPtr("testString")

				// Construct an instance of the Overrides model
				overridesModel := new(rulesetsv1.Overrides)
				overridesModel.Action = core.StringPtr("testString")
				overridesModel.Enabled = core.BoolPtr(true)
				overridesModel.SensitivityLevel = core.StringPtr("high")
				overridesModel.Rules = []rulesetsv1.RulesOverride{*rulesOverrideModel}
				overridesModel.Categories = []rulesetsv1.CategoriesOverride{*categoriesOverrideModel}

				// Construct an instance of the ActionParametersResponse model
				actionParametersResponseModel := new(rulesetsv1.ActionParametersResponse)
				actionParametersResponseModel.Content = core.StringPtr(`{"success": false, "error": "you have been blocked"}`)
				actionParametersResponseModel.ContentType = core.StringPtr("application/json")
				actionParametersResponseModel.StatusCode = core.Int64Ptr(int64(400))

				// Construct an instance of the ActionParameters model
				actionParametersModel := new(rulesetsv1.ActionParameters)
				actionParametersModel.ID = core.StringPtr("testString")
				actionParametersModel.Overrides = overridesModel
				actionParametersModel.Version = core.StringPtr("testString")
				actionParametersModel.Ruleset = core.StringPtr("testString")
				actionParametersModel.Rulesets = []string{"testString"}
				actionParametersModel.Response = actionParametersResponseModel

				// Construct an instance of the Logging model
				loggingModel := new(rulesetsv1.Logging)
				loggingModel.Enabled = core.BoolPtr(true)

				// Construct an instance of the Position model
				positionModel := new(rulesetsv1.Position)
				positionModel.Before = core.StringPtr("testString")
				positionModel.After = core.StringPtr("testString")
				positionModel.Index = core.Int64Ptr(int64(0))

				// Construct an instance of the CreateAccountRulesetRuleOptions model
				createAccountRulesetRuleOptionsModel := new(rulesetsv1.CreateAccountRulesetRuleOptions)
				createAccountRulesetRuleOptionsModel.RulesetID = core.StringPtr("testString")
				createAccountRulesetRuleOptionsModel.Action = core.StringPtr("testString")
				createAccountRulesetRuleOptionsModel.ActionParameters = actionParametersModel
				createAccountRulesetRuleOptionsModel.Description = core.StringPtr("testString")
				createAccountRulesetRuleOptionsModel.Enabled = core.BoolPtr(true)
				createAccountRulesetRuleOptionsModel.Expression = core.StringPtr("ip.src ne 1.1.1.1")
				createAccountRulesetRuleOptionsModel.ID = core.StringPtr("testString")
				createAccountRulesetRuleOptionsModel.Logging = loggingModel
				createAccountRulesetRuleOptionsModel.Ref = core.StringPtr("my_ref")
				createAccountRulesetRuleOptionsModel.Position = positionModel
				createAccountRulesetRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = rulesetsService.CreateAccountRulesetRule(createAccountRulesetRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateAccountRulesetRule with error: Operation validation and request error`, func() {
				rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(rulesetsService).ToNot(BeNil())

				// Construct an instance of the RulesOverride model
				rulesOverrideModel := new(rulesetsv1.RulesOverride)
				rulesOverrideModel.ID = core.StringPtr("testString")
				rulesOverrideModel.Enabled = core.BoolPtr(true)
				rulesOverrideModel.Action = core.StringPtr("testString")
				rulesOverrideModel.SensitivityLevel = core.StringPtr("high")

				// Construct an instance of the CategoriesOverride model
				categoriesOverrideModel := new(rulesetsv1.CategoriesOverride)
				categoriesOverrideModel.Category = core.StringPtr("testString")
				categoriesOverrideModel.Enabled = core.BoolPtr(true)
				categoriesOverrideModel.Action = core.StringPtr("testString")

				// Construct an instance of the Overrides model
				overridesModel := new(rulesetsv1.Overrides)
				overridesModel.Action = core.StringPtr("testString")
				overridesModel.Enabled = core.BoolPtr(true)
				overridesModel.SensitivityLevel = core.StringPtr("high")
				overridesModel.Rules = []rulesetsv1.RulesOverride{*rulesOverrideModel}
				overridesModel.Categories = []rulesetsv1.CategoriesOverride{*categoriesOverrideModel}

				// Construct an instance of the ActionParametersResponse model
				actionParametersResponseModel := new(rulesetsv1.ActionParametersResponse)
				actionParametersResponseModel.Content = core.StringPtr(`{"success": false, "error": "you have been blocked"}`)
				actionParametersResponseModel.ContentType = core.StringPtr("application/json")
				actionParametersResponseModel.StatusCode = core.Int64Ptr(int64(400))

				// Construct an instance of the ActionParameters model
				actionParametersModel := new(rulesetsv1.ActionParameters)
				actionParametersModel.ID = core.StringPtr("testString")
				actionParametersModel.Overrides = overridesModel
				actionParametersModel.Version = core.StringPtr("testString")
				actionParametersModel.Ruleset = core.StringPtr("testString")
				actionParametersModel.Rulesets = []string{"testString"}
				actionParametersModel.Response = actionParametersResponseModel

				// Construct an instance of the Logging model
				loggingModel := new(rulesetsv1.Logging)
				loggingModel.Enabled = core.BoolPtr(true)

				// Construct an instance of the Position model
				positionModel := new(rulesetsv1.Position)
				positionModel.Before = core.StringPtr("testString")
				positionModel.After = core.StringPtr("testString")
				positionModel.Index = core.Int64Ptr(int64(0))

				// Construct an instance of the CreateAccountRulesetRuleOptions model
				createAccountRulesetRuleOptionsModel := new(rulesetsv1.CreateAccountRulesetRuleOptions)
				createAccountRulesetRuleOptionsModel.RulesetID = core.StringPtr("testString")
				createAccountRulesetRuleOptionsModel.Action = core.StringPtr("testString")
				createAccountRulesetRuleOptionsModel.ActionParameters = actionParametersModel
				createAccountRulesetRuleOptionsModel.Description = core.StringPtr("testString")
				createAccountRulesetRuleOptionsModel.Enabled = core.BoolPtr(true)
				createAccountRulesetRuleOptionsModel.Expression = core.StringPtr("ip.src ne 1.1.1.1")
				createAccountRulesetRuleOptionsModel.ID = core.StringPtr("testString")
				createAccountRulesetRuleOptionsModel.Logging = loggingModel
				createAccountRulesetRuleOptionsModel.Ref = core.StringPtr("my_ref")
				createAccountRulesetRuleOptionsModel.Position = positionModel
				createAccountRulesetRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := rulesetsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := rulesetsService.CreateAccountRulesetRule(createAccountRulesetRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateAccountRulesetRuleOptions model with no property values
				createAccountRulesetRuleOptionsModelNew := new(rulesetsv1.CreateAccountRulesetRuleOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = rulesetsService.CreateAccountRulesetRule(createAccountRulesetRuleOptionsModelNew)
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
			It(`Invoke CreateAccountRulesetRule successfully`, func() {
				rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(rulesetsService).ToNot(BeNil())

				// Construct an instance of the RulesOverride model
				rulesOverrideModel := new(rulesetsv1.RulesOverride)
				rulesOverrideModel.ID = core.StringPtr("testString")
				rulesOverrideModel.Enabled = core.BoolPtr(true)
				rulesOverrideModel.Action = core.StringPtr("testString")
				rulesOverrideModel.SensitivityLevel = core.StringPtr("high")

				// Construct an instance of the CategoriesOverride model
				categoriesOverrideModel := new(rulesetsv1.CategoriesOverride)
				categoriesOverrideModel.Category = core.StringPtr("testString")
				categoriesOverrideModel.Enabled = core.BoolPtr(true)
				categoriesOverrideModel.Action = core.StringPtr("testString")

				// Construct an instance of the Overrides model
				overridesModel := new(rulesetsv1.Overrides)
				overridesModel.Action = core.StringPtr("testString")
				overridesModel.Enabled = core.BoolPtr(true)
				overridesModel.SensitivityLevel = core.StringPtr("high")
				overridesModel.Rules = []rulesetsv1.RulesOverride{*rulesOverrideModel}
				overridesModel.Categories = []rulesetsv1.CategoriesOverride{*categoriesOverrideModel}

				// Construct an instance of the ActionParametersResponse model
				actionParametersResponseModel := new(rulesetsv1.ActionParametersResponse)
				actionParametersResponseModel.Content = core.StringPtr(`{"success": false, "error": "you have been blocked"}`)
				actionParametersResponseModel.ContentType = core.StringPtr("application/json")
				actionParametersResponseModel.StatusCode = core.Int64Ptr(int64(400))

				// Construct an instance of the ActionParameters model
				actionParametersModel := new(rulesetsv1.ActionParameters)
				actionParametersModel.ID = core.StringPtr("testString")
				actionParametersModel.Overrides = overridesModel
				actionParametersModel.Version = core.StringPtr("testString")
				actionParametersModel.Ruleset = core.StringPtr("testString")
				actionParametersModel.Rulesets = []string{"testString"}
				actionParametersModel.Response = actionParametersResponseModel

				// Construct an instance of the Logging model
				loggingModel := new(rulesetsv1.Logging)
				loggingModel.Enabled = core.BoolPtr(true)

				// Construct an instance of the Position model
				positionModel := new(rulesetsv1.Position)
				positionModel.Before = core.StringPtr("testString")
				positionModel.After = core.StringPtr("testString")
				positionModel.Index = core.Int64Ptr(int64(0))

				// Construct an instance of the CreateAccountRulesetRuleOptions model
				createAccountRulesetRuleOptionsModel := new(rulesetsv1.CreateAccountRulesetRuleOptions)
				createAccountRulesetRuleOptionsModel.RulesetID = core.StringPtr("testString")
				createAccountRulesetRuleOptionsModel.Action = core.StringPtr("testString")
				createAccountRulesetRuleOptionsModel.ActionParameters = actionParametersModel
				createAccountRulesetRuleOptionsModel.Description = core.StringPtr("testString")
				createAccountRulesetRuleOptionsModel.Enabled = core.BoolPtr(true)
				createAccountRulesetRuleOptionsModel.Expression = core.StringPtr("ip.src ne 1.1.1.1")
				createAccountRulesetRuleOptionsModel.ID = core.StringPtr("testString")
				createAccountRulesetRuleOptionsModel.Logging = loggingModel
				createAccountRulesetRuleOptionsModel.Ref = core.StringPtr("my_ref")
				createAccountRulesetRuleOptionsModel.Position = positionModel
				createAccountRulesetRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := rulesetsService.CreateAccountRulesetRule(createAccountRulesetRuleOptionsModel)
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
	Describe(`UpdateAccountRulesetRule(updateAccountRulesetRuleOptions *UpdateAccountRulesetRuleOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		updateAccountRulesetRulePath := "/v1/testString/rulesets/testString/rules/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateAccountRulesetRulePath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateAccountRulesetRule with error: Operation response processing error`, func() {
				rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(rulesetsService).ToNot(BeNil())

				// Construct an instance of the RulesOverride model
				rulesOverrideModel := new(rulesetsv1.RulesOverride)
				rulesOverrideModel.ID = core.StringPtr("testString")
				rulesOverrideModel.Enabled = core.BoolPtr(true)
				rulesOverrideModel.Action = core.StringPtr("testString")
				rulesOverrideModel.SensitivityLevel = core.StringPtr("high")

				// Construct an instance of the CategoriesOverride model
				categoriesOverrideModel := new(rulesetsv1.CategoriesOverride)
				categoriesOverrideModel.Category = core.StringPtr("testString")
				categoriesOverrideModel.Enabled = core.BoolPtr(true)
				categoriesOverrideModel.Action = core.StringPtr("testString")

				// Construct an instance of the Overrides model
				overridesModel := new(rulesetsv1.Overrides)
				overridesModel.Action = core.StringPtr("testString")
				overridesModel.Enabled = core.BoolPtr(true)
				overridesModel.SensitivityLevel = core.StringPtr("high")
				overridesModel.Rules = []rulesetsv1.RulesOverride{*rulesOverrideModel}
				overridesModel.Categories = []rulesetsv1.CategoriesOverride{*categoriesOverrideModel}

				// Construct an instance of the ActionParametersResponse model
				actionParametersResponseModel := new(rulesetsv1.ActionParametersResponse)
				actionParametersResponseModel.Content = core.StringPtr(`{"success": false, "error": "you have been blocked"}`)
				actionParametersResponseModel.ContentType = core.StringPtr("application/json")
				actionParametersResponseModel.StatusCode = core.Int64Ptr(int64(400))

				// Construct an instance of the ActionParameters model
				actionParametersModel := new(rulesetsv1.ActionParameters)
				actionParametersModel.ID = core.StringPtr("testString")
				actionParametersModel.Overrides = overridesModel
				actionParametersModel.Version = core.StringPtr("testString")
				actionParametersModel.Ruleset = core.StringPtr("testString")
				actionParametersModel.Rulesets = []string{"testString"}
				actionParametersModel.Response = actionParametersResponseModel

				// Construct an instance of the Logging model
				loggingModel := new(rulesetsv1.Logging)
				loggingModel.Enabled = core.BoolPtr(true)

				// Construct an instance of the Position model
				positionModel := new(rulesetsv1.Position)
				positionModel.Before = core.StringPtr("testString")
				positionModel.After = core.StringPtr("testString")
				positionModel.Index = core.Int64Ptr(int64(0))

				// Construct an instance of the UpdateAccountRulesetRuleOptions model
				updateAccountRulesetRuleOptionsModel := new(rulesetsv1.UpdateAccountRulesetRuleOptions)
				updateAccountRulesetRuleOptionsModel.RulesetID = core.StringPtr("testString")
				updateAccountRulesetRuleOptionsModel.RuleID = core.StringPtr("testString")
				updateAccountRulesetRuleOptionsModel.Action = core.StringPtr("testString")
				updateAccountRulesetRuleOptionsModel.ActionParameters = actionParametersModel
				updateAccountRulesetRuleOptionsModel.Description = core.StringPtr("testString")
				updateAccountRulesetRuleOptionsModel.Enabled = core.BoolPtr(true)
				updateAccountRulesetRuleOptionsModel.Expression = core.StringPtr("ip.src ne 1.1.1.1")
				updateAccountRulesetRuleOptionsModel.ID = core.StringPtr("testString")
				updateAccountRulesetRuleOptionsModel.Logging = loggingModel
				updateAccountRulesetRuleOptionsModel.Ref = core.StringPtr("my_ref")
				updateAccountRulesetRuleOptionsModel.Position = positionModel
				updateAccountRulesetRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := rulesetsService.UpdateAccountRulesetRule(updateAccountRulesetRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				rulesetsService.EnableRetries(0, 0)
				result, response, operationErr = rulesetsService.UpdateAccountRulesetRule(updateAccountRulesetRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateAccountRulesetRule(updateAccountRulesetRuleOptions *UpdateAccountRulesetRuleOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		updateAccountRulesetRulePath := "/v1/testString/rulesets/testString/rules/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateAccountRulesetRulePath))
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
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "messages": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "result": {"id": "ID", "version": "Version", "action": "Action", "action_parameters": {"id": "ID", "overrides": {"action": "Action", "enabled": false, "sensitivity_level": "high", "rules": [{"id": "ID", "enabled": false, "action": "Action", "sensitivity_level": "high"}], "categories": [{"category": "Category", "enabled": false, "action": "Action"}]}, "version": "Version", "ruleset": "Ruleset", "rulesets": ["Rulesets"], "response": {"content": "{\"success\": false, \"error\": \"you have been blocked\"}", "content_type": "application/json", "status_code": 400}}, "expression": "ip.src ne 1.1.1.1", "ref": "my_ref", "logging": {"enabled": true}, "last_updated": "2000-01-01T00:00:00.000000Z"}}`)
				}))
			})
			It(`Invoke UpdateAccountRulesetRule successfully with retries`, func() {
				rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(rulesetsService).ToNot(BeNil())
				rulesetsService.EnableRetries(0, 0)

				// Construct an instance of the RulesOverride model
				rulesOverrideModel := new(rulesetsv1.RulesOverride)
				rulesOverrideModel.ID = core.StringPtr("testString")
				rulesOverrideModel.Enabled = core.BoolPtr(true)
				rulesOverrideModel.Action = core.StringPtr("testString")
				rulesOverrideModel.SensitivityLevel = core.StringPtr("high")

				// Construct an instance of the CategoriesOverride model
				categoriesOverrideModel := new(rulesetsv1.CategoriesOverride)
				categoriesOverrideModel.Category = core.StringPtr("testString")
				categoriesOverrideModel.Enabled = core.BoolPtr(true)
				categoriesOverrideModel.Action = core.StringPtr("testString")

				// Construct an instance of the Overrides model
				overridesModel := new(rulesetsv1.Overrides)
				overridesModel.Action = core.StringPtr("testString")
				overridesModel.Enabled = core.BoolPtr(true)
				overridesModel.SensitivityLevel = core.StringPtr("high")
				overridesModel.Rules = []rulesetsv1.RulesOverride{*rulesOverrideModel}
				overridesModel.Categories = []rulesetsv1.CategoriesOverride{*categoriesOverrideModel}

				// Construct an instance of the ActionParametersResponse model
				actionParametersResponseModel := new(rulesetsv1.ActionParametersResponse)
				actionParametersResponseModel.Content = core.StringPtr(`{"success": false, "error": "you have been blocked"}`)
				actionParametersResponseModel.ContentType = core.StringPtr("application/json")
				actionParametersResponseModel.StatusCode = core.Int64Ptr(int64(400))

				// Construct an instance of the ActionParameters model
				actionParametersModel := new(rulesetsv1.ActionParameters)
				actionParametersModel.ID = core.StringPtr("testString")
				actionParametersModel.Overrides = overridesModel
				actionParametersModel.Version = core.StringPtr("testString")
				actionParametersModel.Ruleset = core.StringPtr("testString")
				actionParametersModel.Rulesets = []string{"testString"}
				actionParametersModel.Response = actionParametersResponseModel

				// Construct an instance of the Logging model
				loggingModel := new(rulesetsv1.Logging)
				loggingModel.Enabled = core.BoolPtr(true)

				// Construct an instance of the Position model
				positionModel := new(rulesetsv1.Position)
				positionModel.Before = core.StringPtr("testString")
				positionModel.After = core.StringPtr("testString")
				positionModel.Index = core.Int64Ptr(int64(0))

				// Construct an instance of the UpdateAccountRulesetRuleOptions model
				updateAccountRulesetRuleOptionsModel := new(rulesetsv1.UpdateAccountRulesetRuleOptions)
				updateAccountRulesetRuleOptionsModel.RulesetID = core.StringPtr("testString")
				updateAccountRulesetRuleOptionsModel.RuleID = core.StringPtr("testString")
				updateAccountRulesetRuleOptionsModel.Action = core.StringPtr("testString")
				updateAccountRulesetRuleOptionsModel.ActionParameters = actionParametersModel
				updateAccountRulesetRuleOptionsModel.Description = core.StringPtr("testString")
				updateAccountRulesetRuleOptionsModel.Enabled = core.BoolPtr(true)
				updateAccountRulesetRuleOptionsModel.Expression = core.StringPtr("ip.src ne 1.1.1.1")
				updateAccountRulesetRuleOptionsModel.ID = core.StringPtr("testString")
				updateAccountRulesetRuleOptionsModel.Logging = loggingModel
				updateAccountRulesetRuleOptionsModel.Ref = core.StringPtr("my_ref")
				updateAccountRulesetRuleOptionsModel.Position = positionModel
				updateAccountRulesetRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := rulesetsService.UpdateAccountRulesetRuleWithContext(ctx, updateAccountRulesetRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				rulesetsService.DisableRetries()
				result, response, operationErr := rulesetsService.UpdateAccountRulesetRule(updateAccountRulesetRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = rulesetsService.UpdateAccountRulesetRuleWithContext(ctx, updateAccountRulesetRuleOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(updateAccountRulesetRulePath))
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
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "messages": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "result": {"id": "ID", "version": "Version", "action": "Action", "action_parameters": {"id": "ID", "overrides": {"action": "Action", "enabled": false, "sensitivity_level": "high", "rules": [{"id": "ID", "enabled": false, "action": "Action", "sensitivity_level": "high"}], "categories": [{"category": "Category", "enabled": false, "action": "Action"}]}, "version": "Version", "ruleset": "Ruleset", "rulesets": ["Rulesets"], "response": {"content": "{\"success\": false, \"error\": \"you have been blocked\"}", "content_type": "application/json", "status_code": 400}}, "expression": "ip.src ne 1.1.1.1", "ref": "my_ref", "logging": {"enabled": true}, "last_updated": "2000-01-01T00:00:00.000000Z"}}`)
				}))
			})
			It(`Invoke UpdateAccountRulesetRule successfully`, func() {
				rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(rulesetsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := rulesetsService.UpdateAccountRulesetRule(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the RulesOverride model
				rulesOverrideModel := new(rulesetsv1.RulesOverride)
				rulesOverrideModel.ID = core.StringPtr("testString")
				rulesOverrideModel.Enabled = core.BoolPtr(true)
				rulesOverrideModel.Action = core.StringPtr("testString")
				rulesOverrideModel.SensitivityLevel = core.StringPtr("high")

				// Construct an instance of the CategoriesOverride model
				categoriesOverrideModel := new(rulesetsv1.CategoriesOverride)
				categoriesOverrideModel.Category = core.StringPtr("testString")
				categoriesOverrideModel.Enabled = core.BoolPtr(true)
				categoriesOverrideModel.Action = core.StringPtr("testString")

				// Construct an instance of the Overrides model
				overridesModel := new(rulesetsv1.Overrides)
				overridesModel.Action = core.StringPtr("testString")
				overridesModel.Enabled = core.BoolPtr(true)
				overridesModel.SensitivityLevel = core.StringPtr("high")
				overridesModel.Rules = []rulesetsv1.RulesOverride{*rulesOverrideModel}
				overridesModel.Categories = []rulesetsv1.CategoriesOverride{*categoriesOverrideModel}

				// Construct an instance of the ActionParametersResponse model
				actionParametersResponseModel := new(rulesetsv1.ActionParametersResponse)
				actionParametersResponseModel.Content = core.StringPtr(`{"success": false, "error": "you have been blocked"}`)
				actionParametersResponseModel.ContentType = core.StringPtr("application/json")
				actionParametersResponseModel.StatusCode = core.Int64Ptr(int64(400))

				// Construct an instance of the ActionParameters model
				actionParametersModel := new(rulesetsv1.ActionParameters)
				actionParametersModel.ID = core.StringPtr("testString")
				actionParametersModel.Overrides = overridesModel
				actionParametersModel.Version = core.StringPtr("testString")
				actionParametersModel.Ruleset = core.StringPtr("testString")
				actionParametersModel.Rulesets = []string{"testString"}
				actionParametersModel.Response = actionParametersResponseModel

				// Construct an instance of the Logging model
				loggingModel := new(rulesetsv1.Logging)
				loggingModel.Enabled = core.BoolPtr(true)

				// Construct an instance of the Position model
				positionModel := new(rulesetsv1.Position)
				positionModel.Before = core.StringPtr("testString")
				positionModel.After = core.StringPtr("testString")
				positionModel.Index = core.Int64Ptr(int64(0))

				// Construct an instance of the UpdateAccountRulesetRuleOptions model
				updateAccountRulesetRuleOptionsModel := new(rulesetsv1.UpdateAccountRulesetRuleOptions)
				updateAccountRulesetRuleOptionsModel.RulesetID = core.StringPtr("testString")
				updateAccountRulesetRuleOptionsModel.RuleID = core.StringPtr("testString")
				updateAccountRulesetRuleOptionsModel.Action = core.StringPtr("testString")
				updateAccountRulesetRuleOptionsModel.ActionParameters = actionParametersModel
				updateAccountRulesetRuleOptionsModel.Description = core.StringPtr("testString")
				updateAccountRulesetRuleOptionsModel.Enabled = core.BoolPtr(true)
				updateAccountRulesetRuleOptionsModel.Expression = core.StringPtr("ip.src ne 1.1.1.1")
				updateAccountRulesetRuleOptionsModel.ID = core.StringPtr("testString")
				updateAccountRulesetRuleOptionsModel.Logging = loggingModel
				updateAccountRulesetRuleOptionsModel.Ref = core.StringPtr("my_ref")
				updateAccountRulesetRuleOptionsModel.Position = positionModel
				updateAccountRulesetRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = rulesetsService.UpdateAccountRulesetRule(updateAccountRulesetRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateAccountRulesetRule with error: Operation validation and request error`, func() {
				rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(rulesetsService).ToNot(BeNil())

				// Construct an instance of the RulesOverride model
				rulesOverrideModel := new(rulesetsv1.RulesOverride)
				rulesOverrideModel.ID = core.StringPtr("testString")
				rulesOverrideModel.Enabled = core.BoolPtr(true)
				rulesOverrideModel.Action = core.StringPtr("testString")
				rulesOverrideModel.SensitivityLevel = core.StringPtr("high")

				// Construct an instance of the CategoriesOverride model
				categoriesOverrideModel := new(rulesetsv1.CategoriesOverride)
				categoriesOverrideModel.Category = core.StringPtr("testString")
				categoriesOverrideModel.Enabled = core.BoolPtr(true)
				categoriesOverrideModel.Action = core.StringPtr("testString")

				// Construct an instance of the Overrides model
				overridesModel := new(rulesetsv1.Overrides)
				overridesModel.Action = core.StringPtr("testString")
				overridesModel.Enabled = core.BoolPtr(true)
				overridesModel.SensitivityLevel = core.StringPtr("high")
				overridesModel.Rules = []rulesetsv1.RulesOverride{*rulesOverrideModel}
				overridesModel.Categories = []rulesetsv1.CategoriesOverride{*categoriesOverrideModel}

				// Construct an instance of the ActionParametersResponse model
				actionParametersResponseModel := new(rulesetsv1.ActionParametersResponse)
				actionParametersResponseModel.Content = core.StringPtr(`{"success": false, "error": "you have been blocked"}`)
				actionParametersResponseModel.ContentType = core.StringPtr("application/json")
				actionParametersResponseModel.StatusCode = core.Int64Ptr(int64(400))

				// Construct an instance of the ActionParameters model
				actionParametersModel := new(rulesetsv1.ActionParameters)
				actionParametersModel.ID = core.StringPtr("testString")
				actionParametersModel.Overrides = overridesModel
				actionParametersModel.Version = core.StringPtr("testString")
				actionParametersModel.Ruleset = core.StringPtr("testString")
				actionParametersModel.Rulesets = []string{"testString"}
				actionParametersModel.Response = actionParametersResponseModel

				// Construct an instance of the Logging model
				loggingModel := new(rulesetsv1.Logging)
				loggingModel.Enabled = core.BoolPtr(true)

				// Construct an instance of the Position model
				positionModel := new(rulesetsv1.Position)
				positionModel.Before = core.StringPtr("testString")
				positionModel.After = core.StringPtr("testString")
				positionModel.Index = core.Int64Ptr(int64(0))

				// Construct an instance of the UpdateAccountRulesetRuleOptions model
				updateAccountRulesetRuleOptionsModel := new(rulesetsv1.UpdateAccountRulesetRuleOptions)
				updateAccountRulesetRuleOptionsModel.RulesetID = core.StringPtr("testString")
				updateAccountRulesetRuleOptionsModel.RuleID = core.StringPtr("testString")
				updateAccountRulesetRuleOptionsModel.Action = core.StringPtr("testString")
				updateAccountRulesetRuleOptionsModel.ActionParameters = actionParametersModel
				updateAccountRulesetRuleOptionsModel.Description = core.StringPtr("testString")
				updateAccountRulesetRuleOptionsModel.Enabled = core.BoolPtr(true)
				updateAccountRulesetRuleOptionsModel.Expression = core.StringPtr("ip.src ne 1.1.1.1")
				updateAccountRulesetRuleOptionsModel.ID = core.StringPtr("testString")
				updateAccountRulesetRuleOptionsModel.Logging = loggingModel
				updateAccountRulesetRuleOptionsModel.Ref = core.StringPtr("my_ref")
				updateAccountRulesetRuleOptionsModel.Position = positionModel
				updateAccountRulesetRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := rulesetsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := rulesetsService.UpdateAccountRulesetRule(updateAccountRulesetRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateAccountRulesetRuleOptions model with no property values
				updateAccountRulesetRuleOptionsModelNew := new(rulesetsv1.UpdateAccountRulesetRuleOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = rulesetsService.UpdateAccountRulesetRule(updateAccountRulesetRuleOptionsModelNew)
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
			It(`Invoke UpdateAccountRulesetRule successfully`, func() {
				rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(rulesetsService).ToNot(BeNil())

				// Construct an instance of the RulesOverride model
				rulesOverrideModel := new(rulesetsv1.RulesOverride)
				rulesOverrideModel.ID = core.StringPtr("testString")
				rulesOverrideModel.Enabled = core.BoolPtr(true)
				rulesOverrideModel.Action = core.StringPtr("testString")
				rulesOverrideModel.SensitivityLevel = core.StringPtr("high")

				// Construct an instance of the CategoriesOverride model
				categoriesOverrideModel := new(rulesetsv1.CategoriesOverride)
				categoriesOverrideModel.Category = core.StringPtr("testString")
				categoriesOverrideModel.Enabled = core.BoolPtr(true)
				categoriesOverrideModel.Action = core.StringPtr("testString")

				// Construct an instance of the Overrides model
				overridesModel := new(rulesetsv1.Overrides)
				overridesModel.Action = core.StringPtr("testString")
				overridesModel.Enabled = core.BoolPtr(true)
				overridesModel.SensitivityLevel = core.StringPtr("high")
				overridesModel.Rules = []rulesetsv1.RulesOverride{*rulesOverrideModel}
				overridesModel.Categories = []rulesetsv1.CategoriesOverride{*categoriesOverrideModel}

				// Construct an instance of the ActionParametersResponse model
				actionParametersResponseModel := new(rulesetsv1.ActionParametersResponse)
				actionParametersResponseModel.Content = core.StringPtr(`{"success": false, "error": "you have been blocked"}`)
				actionParametersResponseModel.ContentType = core.StringPtr("application/json")
				actionParametersResponseModel.StatusCode = core.Int64Ptr(int64(400))

				// Construct an instance of the ActionParameters model
				actionParametersModel := new(rulesetsv1.ActionParameters)
				actionParametersModel.ID = core.StringPtr("testString")
				actionParametersModel.Overrides = overridesModel
				actionParametersModel.Version = core.StringPtr("testString")
				actionParametersModel.Ruleset = core.StringPtr("testString")
				actionParametersModel.Rulesets = []string{"testString"}
				actionParametersModel.Response = actionParametersResponseModel

				// Construct an instance of the Logging model
				loggingModel := new(rulesetsv1.Logging)
				loggingModel.Enabled = core.BoolPtr(true)

				// Construct an instance of the Position model
				positionModel := new(rulesetsv1.Position)
				positionModel.Before = core.StringPtr("testString")
				positionModel.After = core.StringPtr("testString")
				positionModel.Index = core.Int64Ptr(int64(0))

				// Construct an instance of the UpdateAccountRulesetRuleOptions model
				updateAccountRulesetRuleOptionsModel := new(rulesetsv1.UpdateAccountRulesetRuleOptions)
				updateAccountRulesetRuleOptionsModel.RulesetID = core.StringPtr("testString")
				updateAccountRulesetRuleOptionsModel.RuleID = core.StringPtr("testString")
				updateAccountRulesetRuleOptionsModel.Action = core.StringPtr("testString")
				updateAccountRulesetRuleOptionsModel.ActionParameters = actionParametersModel
				updateAccountRulesetRuleOptionsModel.Description = core.StringPtr("testString")
				updateAccountRulesetRuleOptionsModel.Enabled = core.BoolPtr(true)
				updateAccountRulesetRuleOptionsModel.Expression = core.StringPtr("ip.src ne 1.1.1.1")
				updateAccountRulesetRuleOptionsModel.ID = core.StringPtr("testString")
				updateAccountRulesetRuleOptionsModel.Logging = loggingModel
				updateAccountRulesetRuleOptionsModel.Ref = core.StringPtr("my_ref")
				updateAccountRulesetRuleOptionsModel.Position = positionModel
				updateAccountRulesetRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := rulesetsService.UpdateAccountRulesetRule(updateAccountRulesetRuleOptionsModel)
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
	Describe(`DeleteAccountRulesetRule(deleteAccountRulesetRuleOptions *DeleteAccountRulesetRuleOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		deleteAccountRulesetRulePath := "/v1/testString/rulesets/testString/rules/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteAccountRulesetRulePath))
					Expect(req.Method).To(Equal("DELETE"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke DeleteAccountRulesetRule with error: Operation response processing error`, func() {
				rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(rulesetsService).ToNot(BeNil())

				// Construct an instance of the DeleteAccountRulesetRuleOptions model
				deleteAccountRulesetRuleOptionsModel := new(rulesetsv1.DeleteAccountRulesetRuleOptions)
				deleteAccountRulesetRuleOptionsModel.RulesetID = core.StringPtr("testString")
				deleteAccountRulesetRuleOptionsModel.RuleID = core.StringPtr("testString")
				deleteAccountRulesetRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := rulesetsService.DeleteAccountRulesetRule(deleteAccountRulesetRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				rulesetsService.EnableRetries(0, 0)
				result, response, operationErr = rulesetsService.DeleteAccountRulesetRule(deleteAccountRulesetRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteAccountRulesetRule(deleteAccountRulesetRuleOptions *DeleteAccountRulesetRuleOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		deleteAccountRulesetRulePath := "/v1/testString/rulesets/testString/rules/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteAccountRulesetRulePath))
					Expect(req.Method).To(Equal("DELETE"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "messages": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "result": {"id": "ID", "version": "Version", "action": "Action", "action_parameters": {"id": "ID", "overrides": {"action": "Action", "enabled": false, "sensitivity_level": "high", "rules": [{"id": "ID", "enabled": false, "action": "Action", "sensitivity_level": "high"}], "categories": [{"category": "Category", "enabled": false, "action": "Action"}]}, "version": "Version", "ruleset": "Ruleset", "rulesets": ["Rulesets"], "response": {"content": "{\"success\": false, \"error\": \"you have been blocked\"}", "content_type": "application/json", "status_code": 400}}, "expression": "ip.src ne 1.1.1.1", "ref": "my_ref", "logging": {"enabled": true}, "last_updated": "2000-01-01T00:00:00.000000Z"}}`)
				}))
			})
			It(`Invoke DeleteAccountRulesetRule successfully with retries`, func() {
				rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(rulesetsService).ToNot(BeNil())
				rulesetsService.EnableRetries(0, 0)

				// Construct an instance of the DeleteAccountRulesetRuleOptions model
				deleteAccountRulesetRuleOptionsModel := new(rulesetsv1.DeleteAccountRulesetRuleOptions)
				deleteAccountRulesetRuleOptionsModel.RulesetID = core.StringPtr("testString")
				deleteAccountRulesetRuleOptionsModel.RuleID = core.StringPtr("testString")
				deleteAccountRulesetRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := rulesetsService.DeleteAccountRulesetRuleWithContext(ctx, deleteAccountRulesetRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				rulesetsService.DisableRetries()
				result, response, operationErr := rulesetsService.DeleteAccountRulesetRule(deleteAccountRulesetRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = rulesetsService.DeleteAccountRulesetRuleWithContext(ctx, deleteAccountRulesetRuleOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(deleteAccountRulesetRulePath))
					Expect(req.Method).To(Equal("DELETE"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "messages": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "result": {"id": "ID", "version": "Version", "action": "Action", "action_parameters": {"id": "ID", "overrides": {"action": "Action", "enabled": false, "sensitivity_level": "high", "rules": [{"id": "ID", "enabled": false, "action": "Action", "sensitivity_level": "high"}], "categories": [{"category": "Category", "enabled": false, "action": "Action"}]}, "version": "Version", "ruleset": "Ruleset", "rulesets": ["Rulesets"], "response": {"content": "{\"success\": false, \"error\": \"you have been blocked\"}", "content_type": "application/json", "status_code": 400}}, "expression": "ip.src ne 1.1.1.1", "ref": "my_ref", "logging": {"enabled": true}, "last_updated": "2000-01-01T00:00:00.000000Z"}}`)
				}))
			})
			It(`Invoke DeleteAccountRulesetRule successfully`, func() {
				rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(rulesetsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := rulesetsService.DeleteAccountRulesetRule(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DeleteAccountRulesetRuleOptions model
				deleteAccountRulesetRuleOptionsModel := new(rulesetsv1.DeleteAccountRulesetRuleOptions)
				deleteAccountRulesetRuleOptionsModel.RulesetID = core.StringPtr("testString")
				deleteAccountRulesetRuleOptionsModel.RuleID = core.StringPtr("testString")
				deleteAccountRulesetRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = rulesetsService.DeleteAccountRulesetRule(deleteAccountRulesetRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke DeleteAccountRulesetRule with error: Operation validation and request error`, func() {
				rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(rulesetsService).ToNot(BeNil())

				// Construct an instance of the DeleteAccountRulesetRuleOptions model
				deleteAccountRulesetRuleOptionsModel := new(rulesetsv1.DeleteAccountRulesetRuleOptions)
				deleteAccountRulesetRuleOptionsModel.RulesetID = core.StringPtr("testString")
				deleteAccountRulesetRuleOptionsModel.RuleID = core.StringPtr("testString")
				deleteAccountRulesetRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := rulesetsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := rulesetsService.DeleteAccountRulesetRule(deleteAccountRulesetRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the DeleteAccountRulesetRuleOptions model with no property values
				deleteAccountRulesetRuleOptionsModelNew := new(rulesetsv1.DeleteAccountRulesetRuleOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = rulesetsService.DeleteAccountRulesetRule(deleteAccountRulesetRuleOptionsModelNew)
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
			It(`Invoke DeleteAccountRulesetRule successfully`, func() {
				rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(rulesetsService).ToNot(BeNil())

				// Construct an instance of the DeleteAccountRulesetRuleOptions model
				deleteAccountRulesetRuleOptionsModel := new(rulesetsv1.DeleteAccountRulesetRuleOptions)
				deleteAccountRulesetRuleOptionsModel.RulesetID = core.StringPtr("testString")
				deleteAccountRulesetRuleOptionsModel.RuleID = core.StringPtr("testString")
				deleteAccountRulesetRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := rulesetsService.DeleteAccountRulesetRule(deleteAccountRulesetRuleOptionsModel)
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
	Describe(`GetAccountRulesetVersionByTag(getAccountRulesetVersionByTagOptions *GetAccountRulesetVersionByTagOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getAccountRulesetVersionByTagPath := "/v1/testString/rulesets/testString/versions/1/by_tag/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getAccountRulesetVersionByTagPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetAccountRulesetVersionByTag with error: Operation response processing error`, func() {
				rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(rulesetsService).ToNot(BeNil())

				// Construct an instance of the GetAccountRulesetVersionByTagOptions model
				getAccountRulesetVersionByTagOptionsModel := new(rulesetsv1.GetAccountRulesetVersionByTagOptions)
				getAccountRulesetVersionByTagOptionsModel.RulesetID = core.StringPtr("testString")
				getAccountRulesetVersionByTagOptionsModel.RulesetVersion = core.StringPtr("1")
				getAccountRulesetVersionByTagOptionsModel.RuleTag = core.StringPtr("testString")
				getAccountRulesetVersionByTagOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := rulesetsService.GetAccountRulesetVersionByTag(getAccountRulesetVersionByTagOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				rulesetsService.EnableRetries(0, 0)
				result, response, operationErr = rulesetsService.GetAccountRulesetVersionByTag(getAccountRulesetVersionByTagOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetAccountRulesetVersionByTag(getAccountRulesetVersionByTagOptions *GetAccountRulesetVersionByTagOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getAccountRulesetVersionByTagPath := "/v1/testString/rulesets/testString/versions/1/by_tag/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getAccountRulesetVersionByTagPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "messages": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "result": {"description": "Custom account ruleset", "id": "ID", "kind": "managed", "last_updated": "2000-01-01T00:00:00.000000Z", "name": "Name", "phase": "ddos_l4", "version": "1", "rules": [{"id": "ID", "version": "Version", "action": "Action", "action_parameters": {"id": "ID", "overrides": {"action": "Action", "enabled": false, "sensitivity_level": "high", "rules": [{"id": "ID", "enabled": false, "action": "Action", "sensitivity_level": "high"}], "categories": [{"category": "Category", "enabled": false, "action": "Action"}]}, "version": "Version", "ruleset": "Ruleset", "rulesets": ["Rulesets"], "response": {"content": "{\"success\": false, \"error\": \"you have been blocked\"}", "content_type": "application/json", "status_code": 400}}, "expression": "ip.src ne 1.1.1.1", "ref": "my_ref", "logging": {"enabled": true}, "last_updated": "2000-01-01T00:00:00.000000Z"}]}}`)
				}))
			})
			It(`Invoke GetAccountRulesetVersionByTag successfully with retries`, func() {
				rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(rulesetsService).ToNot(BeNil())
				rulesetsService.EnableRetries(0, 0)

				// Construct an instance of the GetAccountRulesetVersionByTagOptions model
				getAccountRulesetVersionByTagOptionsModel := new(rulesetsv1.GetAccountRulesetVersionByTagOptions)
				getAccountRulesetVersionByTagOptionsModel.RulesetID = core.StringPtr("testString")
				getAccountRulesetVersionByTagOptionsModel.RulesetVersion = core.StringPtr("1")
				getAccountRulesetVersionByTagOptionsModel.RuleTag = core.StringPtr("testString")
				getAccountRulesetVersionByTagOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := rulesetsService.GetAccountRulesetVersionByTagWithContext(ctx, getAccountRulesetVersionByTagOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				rulesetsService.DisableRetries()
				result, response, operationErr := rulesetsService.GetAccountRulesetVersionByTag(getAccountRulesetVersionByTagOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = rulesetsService.GetAccountRulesetVersionByTagWithContext(ctx, getAccountRulesetVersionByTagOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getAccountRulesetVersionByTagPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "messages": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "result": {"description": "Custom account ruleset", "id": "ID", "kind": "managed", "last_updated": "2000-01-01T00:00:00.000000Z", "name": "Name", "phase": "ddos_l4", "version": "1", "rules": [{"id": "ID", "version": "Version", "action": "Action", "action_parameters": {"id": "ID", "overrides": {"action": "Action", "enabled": false, "sensitivity_level": "high", "rules": [{"id": "ID", "enabled": false, "action": "Action", "sensitivity_level": "high"}], "categories": [{"category": "Category", "enabled": false, "action": "Action"}]}, "version": "Version", "ruleset": "Ruleset", "rulesets": ["Rulesets"], "response": {"content": "{\"success\": false, \"error\": \"you have been blocked\"}", "content_type": "application/json", "status_code": 400}}, "expression": "ip.src ne 1.1.1.1", "ref": "my_ref", "logging": {"enabled": true}, "last_updated": "2000-01-01T00:00:00.000000Z"}]}}`)
				}))
			})
			It(`Invoke GetAccountRulesetVersionByTag successfully`, func() {
				rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(rulesetsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := rulesetsService.GetAccountRulesetVersionByTag(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetAccountRulesetVersionByTagOptions model
				getAccountRulesetVersionByTagOptionsModel := new(rulesetsv1.GetAccountRulesetVersionByTagOptions)
				getAccountRulesetVersionByTagOptionsModel.RulesetID = core.StringPtr("testString")
				getAccountRulesetVersionByTagOptionsModel.RulesetVersion = core.StringPtr("1")
				getAccountRulesetVersionByTagOptionsModel.RuleTag = core.StringPtr("testString")
				getAccountRulesetVersionByTagOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = rulesetsService.GetAccountRulesetVersionByTag(getAccountRulesetVersionByTagOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetAccountRulesetVersionByTag with error: Operation validation and request error`, func() {
				rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(rulesetsService).ToNot(BeNil())

				// Construct an instance of the GetAccountRulesetVersionByTagOptions model
				getAccountRulesetVersionByTagOptionsModel := new(rulesetsv1.GetAccountRulesetVersionByTagOptions)
				getAccountRulesetVersionByTagOptionsModel.RulesetID = core.StringPtr("testString")
				getAccountRulesetVersionByTagOptionsModel.RulesetVersion = core.StringPtr("1")
				getAccountRulesetVersionByTagOptionsModel.RuleTag = core.StringPtr("testString")
				getAccountRulesetVersionByTagOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := rulesetsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := rulesetsService.GetAccountRulesetVersionByTag(getAccountRulesetVersionByTagOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetAccountRulesetVersionByTagOptions model with no property values
				getAccountRulesetVersionByTagOptionsModelNew := new(rulesetsv1.GetAccountRulesetVersionByTagOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = rulesetsService.GetAccountRulesetVersionByTag(getAccountRulesetVersionByTagOptionsModelNew)
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
			It(`Invoke GetAccountRulesetVersionByTag successfully`, func() {
				rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(rulesetsService).ToNot(BeNil())

				// Construct an instance of the GetAccountRulesetVersionByTagOptions model
				getAccountRulesetVersionByTagOptionsModel := new(rulesetsv1.GetAccountRulesetVersionByTagOptions)
				getAccountRulesetVersionByTagOptionsModel.RulesetID = core.StringPtr("testString")
				getAccountRulesetVersionByTagOptionsModel.RulesetVersion = core.StringPtr("1")
				getAccountRulesetVersionByTagOptionsModel.RuleTag = core.StringPtr("testString")
				getAccountRulesetVersionByTagOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := rulesetsService.GetAccountRulesetVersionByTag(getAccountRulesetVersionByTagOptionsModel)
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
	Describe(`GetZoneRulesets(getZoneRulesetsOptions *GetZoneRulesetsOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getZoneRulesetsPath := "/v1/testString/zones/testString/rulesets"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getZoneRulesetsPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetZoneRulesets with error: Operation response processing error`, func() {
				rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(rulesetsService).ToNot(BeNil())

				// Construct an instance of the GetZoneRulesetsOptions model
				getZoneRulesetsOptionsModel := new(rulesetsv1.GetZoneRulesetsOptions)
				getZoneRulesetsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := rulesetsService.GetZoneRulesets(getZoneRulesetsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				rulesetsService.EnableRetries(0, 0)
				result, response, operationErr = rulesetsService.GetZoneRulesets(getZoneRulesetsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetZoneRulesets(getZoneRulesetsOptions *GetZoneRulesetsOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getZoneRulesetsPath := "/v1/testString/zones/testString/rulesets"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getZoneRulesetsPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "messages": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "result": [{"description": "Custom account ruleset", "id": "ID", "kind": "managed", "last_updated": "2000-01-01T00:00:00.000000Z", "name": "Name", "phase": "ddos_l4", "version": "1"}]}`)
				}))
			})
			It(`Invoke GetZoneRulesets successfully with retries`, func() {
				rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(rulesetsService).ToNot(BeNil())
				rulesetsService.EnableRetries(0, 0)

				// Construct an instance of the GetZoneRulesetsOptions model
				getZoneRulesetsOptionsModel := new(rulesetsv1.GetZoneRulesetsOptions)
				getZoneRulesetsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := rulesetsService.GetZoneRulesetsWithContext(ctx, getZoneRulesetsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				rulesetsService.DisableRetries()
				result, response, operationErr := rulesetsService.GetZoneRulesets(getZoneRulesetsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = rulesetsService.GetZoneRulesetsWithContext(ctx, getZoneRulesetsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getZoneRulesetsPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "messages": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "result": [{"description": "Custom account ruleset", "id": "ID", "kind": "managed", "last_updated": "2000-01-01T00:00:00.000000Z", "name": "Name", "phase": "ddos_l4", "version": "1"}]}`)
				}))
			})
			It(`Invoke GetZoneRulesets successfully`, func() {
				rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(rulesetsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := rulesetsService.GetZoneRulesets(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetZoneRulesetsOptions model
				getZoneRulesetsOptionsModel := new(rulesetsv1.GetZoneRulesetsOptions)
				getZoneRulesetsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = rulesetsService.GetZoneRulesets(getZoneRulesetsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetZoneRulesets with error: Operation request error`, func() {
				rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(rulesetsService).ToNot(BeNil())

				// Construct an instance of the GetZoneRulesetsOptions model
				getZoneRulesetsOptionsModel := new(rulesetsv1.GetZoneRulesetsOptions)
				getZoneRulesetsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := rulesetsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := rulesetsService.GetZoneRulesets(getZoneRulesetsOptionsModel)
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
			It(`Invoke GetZoneRulesets successfully`, func() {
				rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(rulesetsService).ToNot(BeNil())

				// Construct an instance of the GetZoneRulesetsOptions model
				getZoneRulesetsOptionsModel := new(rulesetsv1.GetZoneRulesetsOptions)
				getZoneRulesetsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := rulesetsService.GetZoneRulesets(getZoneRulesetsOptionsModel)
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
	Describe(`GetZoneRuleset(getZoneRulesetOptions *GetZoneRulesetOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getZoneRulesetPath := "/v1/testString/zones/testString/rulesets/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getZoneRulesetPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetZoneRuleset with error: Operation response processing error`, func() {
				rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(rulesetsService).ToNot(BeNil())

				// Construct an instance of the GetZoneRulesetOptions model
				getZoneRulesetOptionsModel := new(rulesetsv1.GetZoneRulesetOptions)
				getZoneRulesetOptionsModel.RulesetID = core.StringPtr("testString")
				getZoneRulesetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := rulesetsService.GetZoneRuleset(getZoneRulesetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				rulesetsService.EnableRetries(0, 0)
				result, response, operationErr = rulesetsService.GetZoneRuleset(getZoneRulesetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetZoneRuleset(getZoneRulesetOptions *GetZoneRulesetOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getZoneRulesetPath := "/v1/testString/zones/testString/rulesets/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getZoneRulesetPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "messages": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "result": {"description": "Custom account ruleset", "id": "ID", "kind": "managed", "last_updated": "2000-01-01T00:00:00.000000Z", "name": "Name", "phase": "ddos_l4", "version": "1", "rules": [{"id": "ID", "version": "Version", "action": "Action", "action_parameters": {"id": "ID", "overrides": {"action": "Action", "enabled": false, "sensitivity_level": "high", "rules": [{"id": "ID", "enabled": false, "action": "Action", "sensitivity_level": "high"}], "categories": [{"category": "Category", "enabled": false, "action": "Action"}]}, "version": "Version", "ruleset": "Ruleset", "rulesets": ["Rulesets"], "response": {"content": "{\"success\": false, \"error\": \"you have been blocked\"}", "content_type": "application/json", "status_code": 400}}, "expression": "ip.src ne 1.1.1.1", "ref": "my_ref", "logging": {"enabled": true}, "last_updated": "2000-01-01T00:00:00.000000Z"}]}}`)
				}))
			})
			It(`Invoke GetZoneRuleset successfully with retries`, func() {
				rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(rulesetsService).ToNot(BeNil())
				rulesetsService.EnableRetries(0, 0)

				// Construct an instance of the GetZoneRulesetOptions model
				getZoneRulesetOptionsModel := new(rulesetsv1.GetZoneRulesetOptions)
				getZoneRulesetOptionsModel.RulesetID = core.StringPtr("testString")
				getZoneRulesetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := rulesetsService.GetZoneRulesetWithContext(ctx, getZoneRulesetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				rulesetsService.DisableRetries()
				result, response, operationErr := rulesetsService.GetZoneRuleset(getZoneRulesetOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = rulesetsService.GetZoneRulesetWithContext(ctx, getZoneRulesetOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getZoneRulesetPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "messages": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "result": {"description": "Custom account ruleset", "id": "ID", "kind": "managed", "last_updated": "2000-01-01T00:00:00.000000Z", "name": "Name", "phase": "ddos_l4", "version": "1", "rules": [{"id": "ID", "version": "Version", "action": "Action", "action_parameters": {"id": "ID", "overrides": {"action": "Action", "enabled": false, "sensitivity_level": "high", "rules": [{"id": "ID", "enabled": false, "action": "Action", "sensitivity_level": "high"}], "categories": [{"category": "Category", "enabled": false, "action": "Action"}]}, "version": "Version", "ruleset": "Ruleset", "rulesets": ["Rulesets"], "response": {"content": "{\"success\": false, \"error\": \"you have been blocked\"}", "content_type": "application/json", "status_code": 400}}, "expression": "ip.src ne 1.1.1.1", "ref": "my_ref", "logging": {"enabled": true}, "last_updated": "2000-01-01T00:00:00.000000Z"}]}}`)
				}))
			})
			It(`Invoke GetZoneRuleset successfully`, func() {
				rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(rulesetsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := rulesetsService.GetZoneRuleset(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetZoneRulesetOptions model
				getZoneRulesetOptionsModel := new(rulesetsv1.GetZoneRulesetOptions)
				getZoneRulesetOptionsModel.RulesetID = core.StringPtr("testString")
				getZoneRulesetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = rulesetsService.GetZoneRuleset(getZoneRulesetOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetZoneRuleset with error: Operation validation and request error`, func() {
				rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(rulesetsService).ToNot(BeNil())

				// Construct an instance of the GetZoneRulesetOptions model
				getZoneRulesetOptionsModel := new(rulesetsv1.GetZoneRulesetOptions)
				getZoneRulesetOptionsModel.RulesetID = core.StringPtr("testString")
				getZoneRulesetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := rulesetsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := rulesetsService.GetZoneRuleset(getZoneRulesetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetZoneRulesetOptions model with no property values
				getZoneRulesetOptionsModelNew := new(rulesetsv1.GetZoneRulesetOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = rulesetsService.GetZoneRuleset(getZoneRulesetOptionsModelNew)
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
			It(`Invoke GetZoneRuleset successfully`, func() {
				rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(rulesetsService).ToNot(BeNil())

				// Construct an instance of the GetZoneRulesetOptions model
				getZoneRulesetOptionsModel := new(rulesetsv1.GetZoneRulesetOptions)
				getZoneRulesetOptionsModel.RulesetID = core.StringPtr("testString")
				getZoneRulesetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := rulesetsService.GetZoneRuleset(getZoneRulesetOptionsModel)
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
	Describe(`UpdateZoneRuleset(updateZoneRulesetOptions *UpdateZoneRulesetOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		updateZoneRulesetPath := "/v1/testString/zones/testString/rulesets/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateZoneRulesetPath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateZoneRuleset with error: Operation response processing error`, func() {
				rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(rulesetsService).ToNot(BeNil())

				// Construct an instance of the RulesOverride model
				rulesOverrideModel := new(rulesetsv1.RulesOverride)
				rulesOverrideModel.ID = core.StringPtr("testString")
				rulesOverrideModel.Enabled = core.BoolPtr(true)
				rulesOverrideModel.Action = core.StringPtr("testString")
				rulesOverrideModel.SensitivityLevel = core.StringPtr("high")

				// Construct an instance of the CategoriesOverride model
				categoriesOverrideModel := new(rulesetsv1.CategoriesOverride)
				categoriesOverrideModel.Category = core.StringPtr("testString")
				categoriesOverrideModel.Enabled = core.BoolPtr(true)
				categoriesOverrideModel.Action = core.StringPtr("testString")

				// Construct an instance of the Overrides model
				overridesModel := new(rulesetsv1.Overrides)
				overridesModel.Action = core.StringPtr("testString")
				overridesModel.Enabled = core.BoolPtr(true)
				overridesModel.SensitivityLevel = core.StringPtr("high")
				overridesModel.Rules = []rulesetsv1.RulesOverride{*rulesOverrideModel}
				overridesModel.Categories = []rulesetsv1.CategoriesOverride{*categoriesOverrideModel}

				// Construct an instance of the ActionParametersResponse model
				actionParametersResponseModel := new(rulesetsv1.ActionParametersResponse)
				actionParametersResponseModel.Content = core.StringPtr(`{"success": false, "error": "you have been blocked"}`)
				actionParametersResponseModel.ContentType = core.StringPtr("application/json")
				actionParametersResponseModel.StatusCode = core.Int64Ptr(int64(400))

				// Construct an instance of the ActionParameters model
				actionParametersModel := new(rulesetsv1.ActionParameters)
				actionParametersModel.ID = core.StringPtr("testString")
				actionParametersModel.Overrides = overridesModel
				actionParametersModel.Version = core.StringPtr("testString")
				actionParametersModel.Ruleset = core.StringPtr("testString")
				actionParametersModel.Rulesets = []string{"testString"}
				actionParametersModel.Response = actionParametersResponseModel

				// Construct an instance of the Logging model
				loggingModel := new(rulesetsv1.Logging)
				loggingModel.Enabled = core.BoolPtr(true)

				// Construct an instance of the Position model
				positionModel := new(rulesetsv1.Position)
				positionModel.Before = core.StringPtr("testString")
				positionModel.After = core.StringPtr("testString")
				positionModel.Index = core.Int64Ptr(int64(0))

				// Construct an instance of the RuleCreate model
				ruleCreateModel := new(rulesetsv1.RuleCreate)
				ruleCreateModel.Action = core.StringPtr("testString")
				ruleCreateModel.ActionParameters = actionParametersModel
				ruleCreateModel.Description = core.StringPtr("testString")
				ruleCreateModel.Enabled = core.BoolPtr(true)
				ruleCreateModel.Expression = core.StringPtr("ip.src ne 1.1.1.1")
				ruleCreateModel.ID = core.StringPtr("testString")
				ruleCreateModel.Logging = loggingModel
				ruleCreateModel.Ref = core.StringPtr("my_ref")
				ruleCreateModel.Position = positionModel

				// Construct an instance of the UpdateZoneRulesetOptions model
				updateZoneRulesetOptionsModel := new(rulesetsv1.UpdateZoneRulesetOptions)
				updateZoneRulesetOptionsModel.RulesetID = core.StringPtr("testString")
				updateZoneRulesetOptionsModel.Description = core.StringPtr("Custom account ruleset")
				updateZoneRulesetOptionsModel.Kind = core.StringPtr("managed")
				updateZoneRulesetOptionsModel.Name = core.StringPtr("testString")
				updateZoneRulesetOptionsModel.Phase = core.StringPtr("ddos_l4")
				updateZoneRulesetOptionsModel.Rules = []rulesetsv1.RuleCreate{*ruleCreateModel}
				updateZoneRulesetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := rulesetsService.UpdateZoneRuleset(updateZoneRulesetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				rulesetsService.EnableRetries(0, 0)
				result, response, operationErr = rulesetsService.UpdateZoneRuleset(updateZoneRulesetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateZoneRuleset(updateZoneRulesetOptions *UpdateZoneRulesetOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		updateZoneRulesetPath := "/v1/testString/zones/testString/rulesets/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateZoneRulesetPath))
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
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "messages": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "result": {"description": "Custom account ruleset", "id": "ID", "kind": "managed", "last_updated": "2000-01-01T00:00:00.000000Z", "name": "Name", "phase": "ddos_l4", "version": "1", "rules": [{"id": "ID", "version": "Version", "action": "Action", "action_parameters": {"id": "ID", "overrides": {"action": "Action", "enabled": false, "sensitivity_level": "high", "rules": [{"id": "ID", "enabled": false, "action": "Action", "sensitivity_level": "high"}], "categories": [{"category": "Category", "enabled": false, "action": "Action"}]}, "version": "Version", "ruleset": "Ruleset", "rulesets": ["Rulesets"], "response": {"content": "{\"success\": false, \"error\": \"you have been blocked\"}", "content_type": "application/json", "status_code": 400}}, "expression": "ip.src ne 1.1.1.1", "ref": "my_ref", "logging": {"enabled": true}, "last_updated": "2000-01-01T00:00:00.000000Z"}]}}`)
				}))
			})
			It(`Invoke UpdateZoneRuleset successfully with retries`, func() {
				rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(rulesetsService).ToNot(BeNil())
				rulesetsService.EnableRetries(0, 0)

				// Construct an instance of the RulesOverride model
				rulesOverrideModel := new(rulesetsv1.RulesOverride)
				rulesOverrideModel.ID = core.StringPtr("testString")
				rulesOverrideModel.Enabled = core.BoolPtr(true)
				rulesOverrideModel.Action = core.StringPtr("testString")
				rulesOverrideModel.SensitivityLevel = core.StringPtr("high")

				// Construct an instance of the CategoriesOverride model
				categoriesOverrideModel := new(rulesetsv1.CategoriesOverride)
				categoriesOverrideModel.Category = core.StringPtr("testString")
				categoriesOverrideModel.Enabled = core.BoolPtr(true)
				categoriesOverrideModel.Action = core.StringPtr("testString")

				// Construct an instance of the Overrides model
				overridesModel := new(rulesetsv1.Overrides)
				overridesModel.Action = core.StringPtr("testString")
				overridesModel.Enabled = core.BoolPtr(true)
				overridesModel.SensitivityLevel = core.StringPtr("high")
				overridesModel.Rules = []rulesetsv1.RulesOverride{*rulesOverrideModel}
				overridesModel.Categories = []rulesetsv1.CategoriesOverride{*categoriesOverrideModel}

				// Construct an instance of the ActionParametersResponse model
				actionParametersResponseModel := new(rulesetsv1.ActionParametersResponse)
				actionParametersResponseModel.Content = core.StringPtr(`{"success": false, "error": "you have been blocked"}`)
				actionParametersResponseModel.ContentType = core.StringPtr("application/json")
				actionParametersResponseModel.StatusCode = core.Int64Ptr(int64(400))

				// Construct an instance of the ActionParameters model
				actionParametersModel := new(rulesetsv1.ActionParameters)
				actionParametersModel.ID = core.StringPtr("testString")
				actionParametersModel.Overrides = overridesModel
				actionParametersModel.Version = core.StringPtr("testString")
				actionParametersModel.Ruleset = core.StringPtr("testString")
				actionParametersModel.Rulesets = []string{"testString"}
				actionParametersModel.Response = actionParametersResponseModel

				// Construct an instance of the Logging model
				loggingModel := new(rulesetsv1.Logging)
				loggingModel.Enabled = core.BoolPtr(true)

				// Construct an instance of the Position model
				positionModel := new(rulesetsv1.Position)
				positionModel.Before = core.StringPtr("testString")
				positionModel.After = core.StringPtr("testString")
				positionModel.Index = core.Int64Ptr(int64(0))

				// Construct an instance of the RuleCreate model
				ruleCreateModel := new(rulesetsv1.RuleCreate)
				ruleCreateModel.Action = core.StringPtr("testString")
				ruleCreateModel.ActionParameters = actionParametersModel
				ruleCreateModel.Description = core.StringPtr("testString")
				ruleCreateModel.Enabled = core.BoolPtr(true)
				ruleCreateModel.Expression = core.StringPtr("ip.src ne 1.1.1.1")
				ruleCreateModel.ID = core.StringPtr("testString")
				ruleCreateModel.Logging = loggingModel
				ruleCreateModel.Ref = core.StringPtr("my_ref")
				ruleCreateModel.Position = positionModel

				// Construct an instance of the UpdateZoneRulesetOptions model
				updateZoneRulesetOptionsModel := new(rulesetsv1.UpdateZoneRulesetOptions)
				updateZoneRulesetOptionsModel.RulesetID = core.StringPtr("testString")
				updateZoneRulesetOptionsModel.Description = core.StringPtr("Custom account ruleset")
				updateZoneRulesetOptionsModel.Kind = core.StringPtr("managed")
				updateZoneRulesetOptionsModel.Name = core.StringPtr("testString")
				updateZoneRulesetOptionsModel.Phase = core.StringPtr("ddos_l4")
				updateZoneRulesetOptionsModel.Rules = []rulesetsv1.RuleCreate{*ruleCreateModel}
				updateZoneRulesetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := rulesetsService.UpdateZoneRulesetWithContext(ctx, updateZoneRulesetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				rulesetsService.DisableRetries()
				result, response, operationErr := rulesetsService.UpdateZoneRuleset(updateZoneRulesetOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = rulesetsService.UpdateZoneRulesetWithContext(ctx, updateZoneRulesetOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(updateZoneRulesetPath))
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
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "messages": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "result": {"description": "Custom account ruleset", "id": "ID", "kind": "managed", "last_updated": "2000-01-01T00:00:00.000000Z", "name": "Name", "phase": "ddos_l4", "version": "1", "rules": [{"id": "ID", "version": "Version", "action": "Action", "action_parameters": {"id": "ID", "overrides": {"action": "Action", "enabled": false, "sensitivity_level": "high", "rules": [{"id": "ID", "enabled": false, "action": "Action", "sensitivity_level": "high"}], "categories": [{"category": "Category", "enabled": false, "action": "Action"}]}, "version": "Version", "ruleset": "Ruleset", "rulesets": ["Rulesets"], "response": {"content": "{\"success\": false, \"error\": \"you have been blocked\"}", "content_type": "application/json", "status_code": 400}}, "expression": "ip.src ne 1.1.1.1", "ref": "my_ref", "logging": {"enabled": true}, "last_updated": "2000-01-01T00:00:00.000000Z"}]}}`)
				}))
			})
			It(`Invoke UpdateZoneRuleset successfully`, func() {
				rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(rulesetsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := rulesetsService.UpdateZoneRuleset(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the RulesOverride model
				rulesOverrideModel := new(rulesetsv1.RulesOverride)
				rulesOverrideModel.ID = core.StringPtr("testString")
				rulesOverrideModel.Enabled = core.BoolPtr(true)
				rulesOverrideModel.Action = core.StringPtr("testString")
				rulesOverrideModel.SensitivityLevel = core.StringPtr("high")

				// Construct an instance of the CategoriesOverride model
				categoriesOverrideModel := new(rulesetsv1.CategoriesOverride)
				categoriesOverrideModel.Category = core.StringPtr("testString")
				categoriesOverrideModel.Enabled = core.BoolPtr(true)
				categoriesOverrideModel.Action = core.StringPtr("testString")

				// Construct an instance of the Overrides model
				overridesModel := new(rulesetsv1.Overrides)
				overridesModel.Action = core.StringPtr("testString")
				overridesModel.Enabled = core.BoolPtr(true)
				overridesModel.SensitivityLevel = core.StringPtr("high")
				overridesModel.Rules = []rulesetsv1.RulesOverride{*rulesOverrideModel}
				overridesModel.Categories = []rulesetsv1.CategoriesOverride{*categoriesOverrideModel}

				// Construct an instance of the ActionParametersResponse model
				actionParametersResponseModel := new(rulesetsv1.ActionParametersResponse)
				actionParametersResponseModel.Content = core.StringPtr(`{"success": false, "error": "you have been blocked"}`)
				actionParametersResponseModel.ContentType = core.StringPtr("application/json")
				actionParametersResponseModel.StatusCode = core.Int64Ptr(int64(400))

				// Construct an instance of the ActionParameters model
				actionParametersModel := new(rulesetsv1.ActionParameters)
				actionParametersModel.ID = core.StringPtr("testString")
				actionParametersModel.Overrides = overridesModel
				actionParametersModel.Version = core.StringPtr("testString")
				actionParametersModel.Ruleset = core.StringPtr("testString")
				actionParametersModel.Rulesets = []string{"testString"}
				actionParametersModel.Response = actionParametersResponseModel

				// Construct an instance of the Logging model
				loggingModel := new(rulesetsv1.Logging)
				loggingModel.Enabled = core.BoolPtr(true)

				// Construct an instance of the Position model
				positionModel := new(rulesetsv1.Position)
				positionModel.Before = core.StringPtr("testString")
				positionModel.After = core.StringPtr("testString")
				positionModel.Index = core.Int64Ptr(int64(0))

				// Construct an instance of the RuleCreate model
				ruleCreateModel := new(rulesetsv1.RuleCreate)
				ruleCreateModel.Action = core.StringPtr("testString")
				ruleCreateModel.ActionParameters = actionParametersModel
				ruleCreateModel.Description = core.StringPtr("testString")
				ruleCreateModel.Enabled = core.BoolPtr(true)
				ruleCreateModel.Expression = core.StringPtr("ip.src ne 1.1.1.1")
				ruleCreateModel.ID = core.StringPtr("testString")
				ruleCreateModel.Logging = loggingModel
				ruleCreateModel.Ref = core.StringPtr("my_ref")
				ruleCreateModel.Position = positionModel

				// Construct an instance of the UpdateZoneRulesetOptions model
				updateZoneRulesetOptionsModel := new(rulesetsv1.UpdateZoneRulesetOptions)
				updateZoneRulesetOptionsModel.RulesetID = core.StringPtr("testString")
				updateZoneRulesetOptionsModel.Description = core.StringPtr("Custom account ruleset")
				updateZoneRulesetOptionsModel.Kind = core.StringPtr("managed")
				updateZoneRulesetOptionsModel.Name = core.StringPtr("testString")
				updateZoneRulesetOptionsModel.Phase = core.StringPtr("ddos_l4")
				updateZoneRulesetOptionsModel.Rules = []rulesetsv1.RuleCreate{*ruleCreateModel}
				updateZoneRulesetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = rulesetsService.UpdateZoneRuleset(updateZoneRulesetOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateZoneRuleset with error: Operation validation and request error`, func() {
				rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(rulesetsService).ToNot(BeNil())

				// Construct an instance of the RulesOverride model
				rulesOverrideModel := new(rulesetsv1.RulesOverride)
				rulesOverrideModel.ID = core.StringPtr("testString")
				rulesOverrideModel.Enabled = core.BoolPtr(true)
				rulesOverrideModel.Action = core.StringPtr("testString")
				rulesOverrideModel.SensitivityLevel = core.StringPtr("high")

				// Construct an instance of the CategoriesOverride model
				categoriesOverrideModel := new(rulesetsv1.CategoriesOverride)
				categoriesOverrideModel.Category = core.StringPtr("testString")
				categoriesOverrideModel.Enabled = core.BoolPtr(true)
				categoriesOverrideModel.Action = core.StringPtr("testString")

				// Construct an instance of the Overrides model
				overridesModel := new(rulesetsv1.Overrides)
				overridesModel.Action = core.StringPtr("testString")
				overridesModel.Enabled = core.BoolPtr(true)
				overridesModel.SensitivityLevel = core.StringPtr("high")
				overridesModel.Rules = []rulesetsv1.RulesOverride{*rulesOverrideModel}
				overridesModel.Categories = []rulesetsv1.CategoriesOverride{*categoriesOverrideModel}

				// Construct an instance of the ActionParametersResponse model
				actionParametersResponseModel := new(rulesetsv1.ActionParametersResponse)
				actionParametersResponseModel.Content = core.StringPtr(`{"success": false, "error": "you have been blocked"}`)
				actionParametersResponseModel.ContentType = core.StringPtr("application/json")
				actionParametersResponseModel.StatusCode = core.Int64Ptr(int64(400))

				// Construct an instance of the ActionParameters model
				actionParametersModel := new(rulesetsv1.ActionParameters)
				actionParametersModel.ID = core.StringPtr("testString")
				actionParametersModel.Overrides = overridesModel
				actionParametersModel.Version = core.StringPtr("testString")
				actionParametersModel.Ruleset = core.StringPtr("testString")
				actionParametersModel.Rulesets = []string{"testString"}
				actionParametersModel.Response = actionParametersResponseModel

				// Construct an instance of the Logging model
				loggingModel := new(rulesetsv1.Logging)
				loggingModel.Enabled = core.BoolPtr(true)

				// Construct an instance of the Position model
				positionModel := new(rulesetsv1.Position)
				positionModel.Before = core.StringPtr("testString")
				positionModel.After = core.StringPtr("testString")
				positionModel.Index = core.Int64Ptr(int64(0))

				// Construct an instance of the RuleCreate model
				ruleCreateModel := new(rulesetsv1.RuleCreate)
				ruleCreateModel.Action = core.StringPtr("testString")
				ruleCreateModel.ActionParameters = actionParametersModel
				ruleCreateModel.Description = core.StringPtr("testString")
				ruleCreateModel.Enabled = core.BoolPtr(true)
				ruleCreateModel.Expression = core.StringPtr("ip.src ne 1.1.1.1")
				ruleCreateModel.ID = core.StringPtr("testString")
				ruleCreateModel.Logging = loggingModel
				ruleCreateModel.Ref = core.StringPtr("my_ref")
				ruleCreateModel.Position = positionModel

				// Construct an instance of the UpdateZoneRulesetOptions model
				updateZoneRulesetOptionsModel := new(rulesetsv1.UpdateZoneRulesetOptions)
				updateZoneRulesetOptionsModel.RulesetID = core.StringPtr("testString")
				updateZoneRulesetOptionsModel.Description = core.StringPtr("Custom account ruleset")
				updateZoneRulesetOptionsModel.Kind = core.StringPtr("managed")
				updateZoneRulesetOptionsModel.Name = core.StringPtr("testString")
				updateZoneRulesetOptionsModel.Phase = core.StringPtr("ddos_l4")
				updateZoneRulesetOptionsModel.Rules = []rulesetsv1.RuleCreate{*ruleCreateModel}
				updateZoneRulesetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := rulesetsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := rulesetsService.UpdateZoneRuleset(updateZoneRulesetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateZoneRulesetOptions model with no property values
				updateZoneRulesetOptionsModelNew := new(rulesetsv1.UpdateZoneRulesetOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = rulesetsService.UpdateZoneRuleset(updateZoneRulesetOptionsModelNew)
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
			It(`Invoke UpdateZoneRuleset successfully`, func() {
				rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(rulesetsService).ToNot(BeNil())

				// Construct an instance of the RulesOverride model
				rulesOverrideModel := new(rulesetsv1.RulesOverride)
				rulesOverrideModel.ID = core.StringPtr("testString")
				rulesOverrideModel.Enabled = core.BoolPtr(true)
				rulesOverrideModel.Action = core.StringPtr("testString")
				rulesOverrideModel.SensitivityLevel = core.StringPtr("high")

				// Construct an instance of the CategoriesOverride model
				categoriesOverrideModel := new(rulesetsv1.CategoriesOverride)
				categoriesOverrideModel.Category = core.StringPtr("testString")
				categoriesOverrideModel.Enabled = core.BoolPtr(true)
				categoriesOverrideModel.Action = core.StringPtr("testString")

				// Construct an instance of the Overrides model
				overridesModel := new(rulesetsv1.Overrides)
				overridesModel.Action = core.StringPtr("testString")
				overridesModel.Enabled = core.BoolPtr(true)
				overridesModel.SensitivityLevel = core.StringPtr("high")
				overridesModel.Rules = []rulesetsv1.RulesOverride{*rulesOverrideModel}
				overridesModel.Categories = []rulesetsv1.CategoriesOverride{*categoriesOverrideModel}

				// Construct an instance of the ActionParametersResponse model
				actionParametersResponseModel := new(rulesetsv1.ActionParametersResponse)
				actionParametersResponseModel.Content = core.StringPtr(`{"success": false, "error": "you have been blocked"}`)
				actionParametersResponseModel.ContentType = core.StringPtr("application/json")
				actionParametersResponseModel.StatusCode = core.Int64Ptr(int64(400))

				// Construct an instance of the ActionParameters model
				actionParametersModel := new(rulesetsv1.ActionParameters)
				actionParametersModel.ID = core.StringPtr("testString")
				actionParametersModel.Overrides = overridesModel
				actionParametersModel.Version = core.StringPtr("testString")
				actionParametersModel.Ruleset = core.StringPtr("testString")
				actionParametersModel.Rulesets = []string{"testString"}
				actionParametersModel.Response = actionParametersResponseModel

				// Construct an instance of the Logging model
				loggingModel := new(rulesetsv1.Logging)
				loggingModel.Enabled = core.BoolPtr(true)

				// Construct an instance of the Position model
				positionModel := new(rulesetsv1.Position)
				positionModel.Before = core.StringPtr("testString")
				positionModel.After = core.StringPtr("testString")
				positionModel.Index = core.Int64Ptr(int64(0))

				// Construct an instance of the RuleCreate model
				ruleCreateModel := new(rulesetsv1.RuleCreate)
				ruleCreateModel.Action = core.StringPtr("testString")
				ruleCreateModel.ActionParameters = actionParametersModel
				ruleCreateModel.Description = core.StringPtr("testString")
				ruleCreateModel.Enabled = core.BoolPtr(true)
				ruleCreateModel.Expression = core.StringPtr("ip.src ne 1.1.1.1")
				ruleCreateModel.ID = core.StringPtr("testString")
				ruleCreateModel.Logging = loggingModel
				ruleCreateModel.Ref = core.StringPtr("my_ref")
				ruleCreateModel.Position = positionModel

				// Construct an instance of the UpdateZoneRulesetOptions model
				updateZoneRulesetOptionsModel := new(rulesetsv1.UpdateZoneRulesetOptions)
				updateZoneRulesetOptionsModel.RulesetID = core.StringPtr("testString")
				updateZoneRulesetOptionsModel.Description = core.StringPtr("Custom account ruleset")
				updateZoneRulesetOptionsModel.Kind = core.StringPtr("managed")
				updateZoneRulesetOptionsModel.Name = core.StringPtr("testString")
				updateZoneRulesetOptionsModel.Phase = core.StringPtr("ddos_l4")
				updateZoneRulesetOptionsModel.Rules = []rulesetsv1.RuleCreate{*ruleCreateModel}
				updateZoneRulesetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := rulesetsService.UpdateZoneRuleset(updateZoneRulesetOptionsModel)
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
	Describe(`DeleteZoneRuleset(deleteZoneRulesetOptions *DeleteZoneRulesetOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		deleteZoneRulesetPath := "/v1/testString/zones/testString/rulesets/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteZoneRulesetPath))
					Expect(req.Method).To(Equal("DELETE"))

					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteZoneRuleset successfully`, func() {
				rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(rulesetsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := rulesetsService.DeleteZoneRuleset(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteZoneRulesetOptions model
				deleteZoneRulesetOptionsModel := new(rulesetsv1.DeleteZoneRulesetOptions)
				deleteZoneRulesetOptionsModel.RulesetID = core.StringPtr("testString")
				deleteZoneRulesetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = rulesetsService.DeleteZoneRuleset(deleteZoneRulesetOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteZoneRuleset with error: Operation validation and request error`, func() {
				rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(rulesetsService).ToNot(BeNil())

				// Construct an instance of the DeleteZoneRulesetOptions model
				deleteZoneRulesetOptionsModel := new(rulesetsv1.DeleteZoneRulesetOptions)
				deleteZoneRulesetOptionsModel.RulesetID = core.StringPtr("testString")
				deleteZoneRulesetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := rulesetsService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := rulesetsService.DeleteZoneRuleset(deleteZoneRulesetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteZoneRulesetOptions model with no property values
				deleteZoneRulesetOptionsModelNew := new(rulesetsv1.DeleteZoneRulesetOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = rulesetsService.DeleteZoneRuleset(deleteZoneRulesetOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetZoneRulesetVersions(getZoneRulesetVersionsOptions *GetZoneRulesetVersionsOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getZoneRulesetVersionsPath := "/v1/testString/zones/testString/rulesets/testString/versions"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getZoneRulesetVersionsPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetZoneRulesetVersions with error: Operation response processing error`, func() {
				rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(rulesetsService).ToNot(BeNil())

				// Construct an instance of the GetZoneRulesetVersionsOptions model
				getZoneRulesetVersionsOptionsModel := new(rulesetsv1.GetZoneRulesetVersionsOptions)
				getZoneRulesetVersionsOptionsModel.RulesetID = core.StringPtr("testString")
				getZoneRulesetVersionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := rulesetsService.GetZoneRulesetVersions(getZoneRulesetVersionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				rulesetsService.EnableRetries(0, 0)
				result, response, operationErr = rulesetsService.GetZoneRulesetVersions(getZoneRulesetVersionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetZoneRulesetVersions(getZoneRulesetVersionsOptions *GetZoneRulesetVersionsOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getZoneRulesetVersionsPath := "/v1/testString/zones/testString/rulesets/testString/versions"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getZoneRulesetVersionsPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "messages": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "result": [{"description": "Custom account ruleset", "id": "ID", "kind": "managed", "last_updated": "2000-01-01T00:00:00.000000Z", "name": "Name", "phase": "ddos_l4", "version": "1"}]}`)
				}))
			})
			It(`Invoke GetZoneRulesetVersions successfully with retries`, func() {
				rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(rulesetsService).ToNot(BeNil())
				rulesetsService.EnableRetries(0, 0)

				// Construct an instance of the GetZoneRulesetVersionsOptions model
				getZoneRulesetVersionsOptionsModel := new(rulesetsv1.GetZoneRulesetVersionsOptions)
				getZoneRulesetVersionsOptionsModel.RulesetID = core.StringPtr("testString")
				getZoneRulesetVersionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := rulesetsService.GetZoneRulesetVersionsWithContext(ctx, getZoneRulesetVersionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				rulesetsService.DisableRetries()
				result, response, operationErr := rulesetsService.GetZoneRulesetVersions(getZoneRulesetVersionsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = rulesetsService.GetZoneRulesetVersionsWithContext(ctx, getZoneRulesetVersionsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getZoneRulesetVersionsPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "messages": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "result": [{"description": "Custom account ruleset", "id": "ID", "kind": "managed", "last_updated": "2000-01-01T00:00:00.000000Z", "name": "Name", "phase": "ddos_l4", "version": "1"}]}`)
				}))
			})
			It(`Invoke GetZoneRulesetVersions successfully`, func() {
				rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(rulesetsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := rulesetsService.GetZoneRulesetVersions(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetZoneRulesetVersionsOptions model
				getZoneRulesetVersionsOptionsModel := new(rulesetsv1.GetZoneRulesetVersionsOptions)
				getZoneRulesetVersionsOptionsModel.RulesetID = core.StringPtr("testString")
				getZoneRulesetVersionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = rulesetsService.GetZoneRulesetVersions(getZoneRulesetVersionsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetZoneRulesetVersions with error: Operation validation and request error`, func() {
				rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(rulesetsService).ToNot(BeNil())

				// Construct an instance of the GetZoneRulesetVersionsOptions model
				getZoneRulesetVersionsOptionsModel := new(rulesetsv1.GetZoneRulesetVersionsOptions)
				getZoneRulesetVersionsOptionsModel.RulesetID = core.StringPtr("testString")
				getZoneRulesetVersionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := rulesetsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := rulesetsService.GetZoneRulesetVersions(getZoneRulesetVersionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetZoneRulesetVersionsOptions model with no property values
				getZoneRulesetVersionsOptionsModelNew := new(rulesetsv1.GetZoneRulesetVersionsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = rulesetsService.GetZoneRulesetVersions(getZoneRulesetVersionsOptionsModelNew)
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
			It(`Invoke GetZoneRulesetVersions successfully`, func() {
				rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(rulesetsService).ToNot(BeNil())

				// Construct an instance of the GetZoneRulesetVersionsOptions model
				getZoneRulesetVersionsOptionsModel := new(rulesetsv1.GetZoneRulesetVersionsOptions)
				getZoneRulesetVersionsOptionsModel.RulesetID = core.StringPtr("testString")
				getZoneRulesetVersionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := rulesetsService.GetZoneRulesetVersions(getZoneRulesetVersionsOptionsModel)
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
	Describe(`GetZoneRulesetVersion(getZoneRulesetVersionOptions *GetZoneRulesetVersionOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getZoneRulesetVersionPath := "/v1/testString/zones/testString/rulesets/testString/versions/1"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getZoneRulesetVersionPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetZoneRulesetVersion with error: Operation response processing error`, func() {
				rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(rulesetsService).ToNot(BeNil())

				// Construct an instance of the GetZoneRulesetVersionOptions model
				getZoneRulesetVersionOptionsModel := new(rulesetsv1.GetZoneRulesetVersionOptions)
				getZoneRulesetVersionOptionsModel.RulesetID = core.StringPtr("testString")
				getZoneRulesetVersionOptionsModel.RulesetVersion = core.StringPtr("1")
				getZoneRulesetVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := rulesetsService.GetZoneRulesetVersion(getZoneRulesetVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				rulesetsService.EnableRetries(0, 0)
				result, response, operationErr = rulesetsService.GetZoneRulesetVersion(getZoneRulesetVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetZoneRulesetVersion(getZoneRulesetVersionOptions *GetZoneRulesetVersionOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getZoneRulesetVersionPath := "/v1/testString/zones/testString/rulesets/testString/versions/1"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getZoneRulesetVersionPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "messages": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "result": {"description": "Custom account ruleset", "id": "ID", "kind": "managed", "last_updated": "2000-01-01T00:00:00.000000Z", "name": "Name", "phase": "ddos_l4", "version": "1", "rules": [{"id": "ID", "version": "Version", "action": "Action", "action_parameters": {"id": "ID", "overrides": {"action": "Action", "enabled": false, "sensitivity_level": "high", "rules": [{"id": "ID", "enabled": false, "action": "Action", "sensitivity_level": "high"}], "categories": [{"category": "Category", "enabled": false, "action": "Action"}]}, "version": "Version", "ruleset": "Ruleset", "rulesets": ["Rulesets"], "response": {"content": "{\"success\": false, \"error\": \"you have been blocked\"}", "content_type": "application/json", "status_code": 400}}, "expression": "ip.src ne 1.1.1.1", "ref": "my_ref", "logging": {"enabled": true}, "last_updated": "2000-01-01T00:00:00.000000Z"}]}}`)
				}))
			})
			It(`Invoke GetZoneRulesetVersion successfully with retries`, func() {
				rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(rulesetsService).ToNot(BeNil())
				rulesetsService.EnableRetries(0, 0)

				// Construct an instance of the GetZoneRulesetVersionOptions model
				getZoneRulesetVersionOptionsModel := new(rulesetsv1.GetZoneRulesetVersionOptions)
				getZoneRulesetVersionOptionsModel.RulesetID = core.StringPtr("testString")
				getZoneRulesetVersionOptionsModel.RulesetVersion = core.StringPtr("1")
				getZoneRulesetVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := rulesetsService.GetZoneRulesetVersionWithContext(ctx, getZoneRulesetVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				rulesetsService.DisableRetries()
				result, response, operationErr := rulesetsService.GetZoneRulesetVersion(getZoneRulesetVersionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = rulesetsService.GetZoneRulesetVersionWithContext(ctx, getZoneRulesetVersionOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getZoneRulesetVersionPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "messages": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "result": {"description": "Custom account ruleset", "id": "ID", "kind": "managed", "last_updated": "2000-01-01T00:00:00.000000Z", "name": "Name", "phase": "ddos_l4", "version": "1", "rules": [{"id": "ID", "version": "Version", "action": "Action", "action_parameters": {"id": "ID", "overrides": {"action": "Action", "enabled": false, "sensitivity_level": "high", "rules": [{"id": "ID", "enabled": false, "action": "Action", "sensitivity_level": "high"}], "categories": [{"category": "Category", "enabled": false, "action": "Action"}]}, "version": "Version", "ruleset": "Ruleset", "rulesets": ["Rulesets"], "response": {"content": "{\"success\": false, \"error\": \"you have been blocked\"}", "content_type": "application/json", "status_code": 400}}, "expression": "ip.src ne 1.1.1.1", "ref": "my_ref", "logging": {"enabled": true}, "last_updated": "2000-01-01T00:00:00.000000Z"}]}}`)
				}))
			})
			It(`Invoke GetZoneRulesetVersion successfully`, func() {
				rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(rulesetsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := rulesetsService.GetZoneRulesetVersion(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetZoneRulesetVersionOptions model
				getZoneRulesetVersionOptionsModel := new(rulesetsv1.GetZoneRulesetVersionOptions)
				getZoneRulesetVersionOptionsModel.RulesetID = core.StringPtr("testString")
				getZoneRulesetVersionOptionsModel.RulesetVersion = core.StringPtr("1")
				getZoneRulesetVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = rulesetsService.GetZoneRulesetVersion(getZoneRulesetVersionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetZoneRulesetVersion with error: Operation validation and request error`, func() {
				rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(rulesetsService).ToNot(BeNil())

				// Construct an instance of the GetZoneRulesetVersionOptions model
				getZoneRulesetVersionOptionsModel := new(rulesetsv1.GetZoneRulesetVersionOptions)
				getZoneRulesetVersionOptionsModel.RulesetID = core.StringPtr("testString")
				getZoneRulesetVersionOptionsModel.RulesetVersion = core.StringPtr("1")
				getZoneRulesetVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := rulesetsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := rulesetsService.GetZoneRulesetVersion(getZoneRulesetVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetZoneRulesetVersionOptions model with no property values
				getZoneRulesetVersionOptionsModelNew := new(rulesetsv1.GetZoneRulesetVersionOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = rulesetsService.GetZoneRulesetVersion(getZoneRulesetVersionOptionsModelNew)
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
			It(`Invoke GetZoneRulesetVersion successfully`, func() {
				rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(rulesetsService).ToNot(BeNil())

				// Construct an instance of the GetZoneRulesetVersionOptions model
				getZoneRulesetVersionOptionsModel := new(rulesetsv1.GetZoneRulesetVersionOptions)
				getZoneRulesetVersionOptionsModel.RulesetID = core.StringPtr("testString")
				getZoneRulesetVersionOptionsModel.RulesetVersion = core.StringPtr("1")
				getZoneRulesetVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := rulesetsService.GetZoneRulesetVersion(getZoneRulesetVersionOptionsModel)
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
	Describe(`DeleteZoneRulesetVersion(deleteZoneRulesetVersionOptions *DeleteZoneRulesetVersionOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		deleteZoneRulesetVersionPath := "/v1/testString/zones/testString/rulesets/testString/versions/1"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteZoneRulesetVersionPath))
					Expect(req.Method).To(Equal("DELETE"))

					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteZoneRulesetVersion successfully`, func() {
				rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(rulesetsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := rulesetsService.DeleteZoneRulesetVersion(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteZoneRulesetVersionOptions model
				deleteZoneRulesetVersionOptionsModel := new(rulesetsv1.DeleteZoneRulesetVersionOptions)
				deleteZoneRulesetVersionOptionsModel.RulesetID = core.StringPtr("testString")
				deleteZoneRulesetVersionOptionsModel.RulesetVersion = core.StringPtr("1")
				deleteZoneRulesetVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = rulesetsService.DeleteZoneRulesetVersion(deleteZoneRulesetVersionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteZoneRulesetVersion with error: Operation validation and request error`, func() {
				rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(rulesetsService).ToNot(BeNil())

				// Construct an instance of the DeleteZoneRulesetVersionOptions model
				deleteZoneRulesetVersionOptionsModel := new(rulesetsv1.DeleteZoneRulesetVersionOptions)
				deleteZoneRulesetVersionOptionsModel.RulesetID = core.StringPtr("testString")
				deleteZoneRulesetVersionOptionsModel.RulesetVersion = core.StringPtr("1")
				deleteZoneRulesetVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := rulesetsService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := rulesetsService.DeleteZoneRulesetVersion(deleteZoneRulesetVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteZoneRulesetVersionOptions model with no property values
				deleteZoneRulesetVersionOptionsModelNew := new(rulesetsv1.DeleteZoneRulesetVersionOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = rulesetsService.DeleteZoneRulesetVersion(deleteZoneRulesetVersionOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetZoneEntrypointRuleset(getZoneEntrypointRulesetOptions *GetZoneEntrypointRulesetOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getZoneEntrypointRulesetPath := "/v1/testString/zones/testString/rulesets/phases/ddos_l4/entrypoint"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getZoneEntrypointRulesetPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetZoneEntrypointRuleset with error: Operation response processing error`, func() {
				rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(rulesetsService).ToNot(BeNil())

				// Construct an instance of the GetZoneEntrypointRulesetOptions model
				getZoneEntrypointRulesetOptionsModel := new(rulesetsv1.GetZoneEntrypointRulesetOptions)
				getZoneEntrypointRulesetOptionsModel.RulesetPhase = core.StringPtr("ddos_l4")
				getZoneEntrypointRulesetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := rulesetsService.GetZoneEntrypointRuleset(getZoneEntrypointRulesetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				rulesetsService.EnableRetries(0, 0)
				result, response, operationErr = rulesetsService.GetZoneEntrypointRuleset(getZoneEntrypointRulesetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetZoneEntrypointRuleset(getZoneEntrypointRulesetOptions *GetZoneEntrypointRulesetOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getZoneEntrypointRulesetPath := "/v1/testString/zones/testString/rulesets/phases/ddos_l4/entrypoint"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getZoneEntrypointRulesetPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "messages": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "result": {"description": "Custom account ruleset", "id": "ID", "kind": "managed", "last_updated": "2000-01-01T00:00:00.000000Z", "name": "Name", "phase": "ddos_l4", "version": "1", "rules": [{"id": "ID", "version": "Version", "action": "Action", "action_parameters": {"id": "ID", "overrides": {"action": "Action", "enabled": false, "sensitivity_level": "high", "rules": [{"id": "ID", "enabled": false, "action": "Action", "sensitivity_level": "high"}], "categories": [{"category": "Category", "enabled": false, "action": "Action"}]}, "version": "Version", "ruleset": "Ruleset", "rulesets": ["Rulesets"], "response": {"content": "{\"success\": false, \"error\": \"you have been blocked\"}", "content_type": "application/json", "status_code": 400}}, "expression": "ip.src ne 1.1.1.1", "ref": "my_ref", "logging": {"enabled": true}, "last_updated": "2000-01-01T00:00:00.000000Z"}]}}`)
				}))
			})
			It(`Invoke GetZoneEntrypointRuleset successfully with retries`, func() {
				rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(rulesetsService).ToNot(BeNil())
				rulesetsService.EnableRetries(0, 0)

				// Construct an instance of the GetZoneEntrypointRulesetOptions model
				getZoneEntrypointRulesetOptionsModel := new(rulesetsv1.GetZoneEntrypointRulesetOptions)
				getZoneEntrypointRulesetOptionsModel.RulesetPhase = core.StringPtr("ddos_l4")
				getZoneEntrypointRulesetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := rulesetsService.GetZoneEntrypointRulesetWithContext(ctx, getZoneEntrypointRulesetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				rulesetsService.DisableRetries()
				result, response, operationErr := rulesetsService.GetZoneEntrypointRuleset(getZoneEntrypointRulesetOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = rulesetsService.GetZoneEntrypointRulesetWithContext(ctx, getZoneEntrypointRulesetOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getZoneEntrypointRulesetPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "messages": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "result": {"description": "Custom account ruleset", "id": "ID", "kind": "managed", "last_updated": "2000-01-01T00:00:00.000000Z", "name": "Name", "phase": "ddos_l4", "version": "1", "rules": [{"id": "ID", "version": "Version", "action": "Action", "action_parameters": {"id": "ID", "overrides": {"action": "Action", "enabled": false, "sensitivity_level": "high", "rules": [{"id": "ID", "enabled": false, "action": "Action", "sensitivity_level": "high"}], "categories": [{"category": "Category", "enabled": false, "action": "Action"}]}, "version": "Version", "ruleset": "Ruleset", "rulesets": ["Rulesets"], "response": {"content": "{\"success\": false, \"error\": \"you have been blocked\"}", "content_type": "application/json", "status_code": 400}}, "expression": "ip.src ne 1.1.1.1", "ref": "my_ref", "logging": {"enabled": true}, "last_updated": "2000-01-01T00:00:00.000000Z"}]}}`)
				}))
			})
			It(`Invoke GetZoneEntrypointRuleset successfully`, func() {
				rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(rulesetsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := rulesetsService.GetZoneEntrypointRuleset(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetZoneEntrypointRulesetOptions model
				getZoneEntrypointRulesetOptionsModel := new(rulesetsv1.GetZoneEntrypointRulesetOptions)
				getZoneEntrypointRulesetOptionsModel.RulesetPhase = core.StringPtr("ddos_l4")
				getZoneEntrypointRulesetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = rulesetsService.GetZoneEntrypointRuleset(getZoneEntrypointRulesetOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetZoneEntrypointRuleset with error: Operation validation and request error`, func() {
				rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(rulesetsService).ToNot(BeNil())

				// Construct an instance of the GetZoneEntrypointRulesetOptions model
				getZoneEntrypointRulesetOptionsModel := new(rulesetsv1.GetZoneEntrypointRulesetOptions)
				getZoneEntrypointRulesetOptionsModel.RulesetPhase = core.StringPtr("ddos_l4")
				getZoneEntrypointRulesetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := rulesetsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := rulesetsService.GetZoneEntrypointRuleset(getZoneEntrypointRulesetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetZoneEntrypointRulesetOptions model with no property values
				getZoneEntrypointRulesetOptionsModelNew := new(rulesetsv1.GetZoneEntrypointRulesetOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = rulesetsService.GetZoneEntrypointRuleset(getZoneEntrypointRulesetOptionsModelNew)
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
			It(`Invoke GetZoneEntrypointRuleset successfully`, func() {
				rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(rulesetsService).ToNot(BeNil())

				// Construct an instance of the GetZoneEntrypointRulesetOptions model
				getZoneEntrypointRulesetOptionsModel := new(rulesetsv1.GetZoneEntrypointRulesetOptions)
				getZoneEntrypointRulesetOptionsModel.RulesetPhase = core.StringPtr("ddos_l4")
				getZoneEntrypointRulesetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := rulesetsService.GetZoneEntrypointRuleset(getZoneEntrypointRulesetOptionsModel)
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
	Describe(`UpdateZoneEntrypointRuleset(updateZoneEntrypointRulesetOptions *UpdateZoneEntrypointRulesetOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		updateZoneEntrypointRulesetPath := "/v1/testString/zones/testString/rulesets/phases/ddos_l4/entrypoint"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateZoneEntrypointRulesetPath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateZoneEntrypointRuleset with error: Operation response processing error`, func() {
				rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(rulesetsService).ToNot(BeNil())

				// Construct an instance of the RulesOverride model
				rulesOverrideModel := new(rulesetsv1.RulesOverride)
				rulesOverrideModel.ID = core.StringPtr("testString")
				rulesOverrideModel.Enabled = core.BoolPtr(true)
				rulesOverrideModel.Action = core.StringPtr("testString")
				rulesOverrideModel.SensitivityLevel = core.StringPtr("high")

				// Construct an instance of the CategoriesOverride model
				categoriesOverrideModel := new(rulesetsv1.CategoriesOverride)
				categoriesOverrideModel.Category = core.StringPtr("testString")
				categoriesOverrideModel.Enabled = core.BoolPtr(true)
				categoriesOverrideModel.Action = core.StringPtr("testString")

				// Construct an instance of the Overrides model
				overridesModel := new(rulesetsv1.Overrides)
				overridesModel.Action = core.StringPtr("testString")
				overridesModel.Enabled = core.BoolPtr(true)
				overridesModel.SensitivityLevel = core.StringPtr("high")
				overridesModel.Rules = []rulesetsv1.RulesOverride{*rulesOverrideModel}
				overridesModel.Categories = []rulesetsv1.CategoriesOverride{*categoriesOverrideModel}

				// Construct an instance of the ActionParametersResponse model
				actionParametersResponseModel := new(rulesetsv1.ActionParametersResponse)
				actionParametersResponseModel.Content = core.StringPtr(`{"success": false, "error": "you have been blocked"}`)
				actionParametersResponseModel.ContentType = core.StringPtr("application/json")
				actionParametersResponseModel.StatusCode = core.Int64Ptr(int64(400))

				// Construct an instance of the ActionParameters model
				actionParametersModel := new(rulesetsv1.ActionParameters)
				actionParametersModel.ID = core.StringPtr("testString")
				actionParametersModel.Overrides = overridesModel
				actionParametersModel.Version = core.StringPtr("testString")
				actionParametersModel.Ruleset = core.StringPtr("testString")
				actionParametersModel.Rulesets = []string{"testString"}
				actionParametersModel.Response = actionParametersResponseModel

				// Construct an instance of the Logging model
				loggingModel := new(rulesetsv1.Logging)
				loggingModel.Enabled = core.BoolPtr(true)

				// Construct an instance of the Position model
				positionModel := new(rulesetsv1.Position)
				positionModel.Before = core.StringPtr("testString")
				positionModel.After = core.StringPtr("testString")
				positionModel.Index = core.Int64Ptr(int64(0))

				// Construct an instance of the RuleCreate model
				ruleCreateModel := new(rulesetsv1.RuleCreate)
				ruleCreateModel.Action = core.StringPtr("testString")
				ruleCreateModel.ActionParameters = actionParametersModel
				ruleCreateModel.Description = core.StringPtr("testString")
				ruleCreateModel.Enabled = core.BoolPtr(true)
				ruleCreateModel.Expression = core.StringPtr("ip.src ne 1.1.1.1")
				ruleCreateModel.ID = core.StringPtr("testString")
				ruleCreateModel.Logging = loggingModel
				ruleCreateModel.Ref = core.StringPtr("my_ref")
				ruleCreateModel.Position = positionModel

				// Construct an instance of the UpdateZoneEntrypointRulesetOptions model
				updateZoneEntrypointRulesetOptionsModel := new(rulesetsv1.UpdateZoneEntrypointRulesetOptions)
				updateZoneEntrypointRulesetOptionsModel.RulesetPhase = core.StringPtr("ddos_l4")
				updateZoneEntrypointRulesetOptionsModel.Description = core.StringPtr("Custom account ruleset")
				updateZoneEntrypointRulesetOptionsModel.Kind = core.StringPtr("managed")
				updateZoneEntrypointRulesetOptionsModel.Name = core.StringPtr("testString")
				updateZoneEntrypointRulesetOptionsModel.Phase = core.StringPtr("ddos_l4")
				updateZoneEntrypointRulesetOptionsModel.Rules = []rulesetsv1.RuleCreate{*ruleCreateModel}
				updateZoneEntrypointRulesetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := rulesetsService.UpdateZoneEntrypointRuleset(updateZoneEntrypointRulesetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				rulesetsService.EnableRetries(0, 0)
				result, response, operationErr = rulesetsService.UpdateZoneEntrypointRuleset(updateZoneEntrypointRulesetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateZoneEntrypointRuleset(updateZoneEntrypointRulesetOptions *UpdateZoneEntrypointRulesetOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		updateZoneEntrypointRulesetPath := "/v1/testString/zones/testString/rulesets/phases/ddos_l4/entrypoint"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateZoneEntrypointRulesetPath))
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
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "messages": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "result": {"description": "Custom account ruleset", "id": "ID", "kind": "managed", "last_updated": "2000-01-01T00:00:00.000000Z", "name": "Name", "phase": "ddos_l4", "version": "1", "rules": [{"id": "ID", "version": "Version", "action": "Action", "action_parameters": {"id": "ID", "overrides": {"action": "Action", "enabled": false, "sensitivity_level": "high", "rules": [{"id": "ID", "enabled": false, "action": "Action", "sensitivity_level": "high"}], "categories": [{"category": "Category", "enabled": false, "action": "Action"}]}, "version": "Version", "ruleset": "Ruleset", "rulesets": ["Rulesets"], "response": {"content": "{\"success\": false, \"error\": \"you have been blocked\"}", "content_type": "application/json", "status_code": 400}}, "expression": "ip.src ne 1.1.1.1", "ref": "my_ref", "logging": {"enabled": true}, "last_updated": "2000-01-01T00:00:00.000000Z"}]}}`)
				}))
			})
			It(`Invoke UpdateZoneEntrypointRuleset successfully with retries`, func() {
				rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(rulesetsService).ToNot(BeNil())
				rulesetsService.EnableRetries(0, 0)

				// Construct an instance of the RulesOverride model
				rulesOverrideModel := new(rulesetsv1.RulesOverride)
				rulesOverrideModel.ID = core.StringPtr("testString")
				rulesOverrideModel.Enabled = core.BoolPtr(true)
				rulesOverrideModel.Action = core.StringPtr("testString")
				rulesOverrideModel.SensitivityLevel = core.StringPtr("high")

				// Construct an instance of the CategoriesOverride model
				categoriesOverrideModel := new(rulesetsv1.CategoriesOverride)
				categoriesOverrideModel.Category = core.StringPtr("testString")
				categoriesOverrideModel.Enabled = core.BoolPtr(true)
				categoriesOverrideModel.Action = core.StringPtr("testString")

				// Construct an instance of the Overrides model
				overridesModel := new(rulesetsv1.Overrides)
				overridesModel.Action = core.StringPtr("testString")
				overridesModel.Enabled = core.BoolPtr(true)
				overridesModel.SensitivityLevel = core.StringPtr("high")
				overridesModel.Rules = []rulesetsv1.RulesOverride{*rulesOverrideModel}
				overridesModel.Categories = []rulesetsv1.CategoriesOverride{*categoriesOverrideModel}

				// Construct an instance of the ActionParametersResponse model
				actionParametersResponseModel := new(rulesetsv1.ActionParametersResponse)
				actionParametersResponseModel.Content = core.StringPtr(`{"success": false, "error": "you have been blocked"}`)
				actionParametersResponseModel.ContentType = core.StringPtr("application/json")
				actionParametersResponseModel.StatusCode = core.Int64Ptr(int64(400))

				// Construct an instance of the ActionParameters model
				actionParametersModel := new(rulesetsv1.ActionParameters)
				actionParametersModel.ID = core.StringPtr("testString")
				actionParametersModel.Overrides = overridesModel
				actionParametersModel.Version = core.StringPtr("testString")
				actionParametersModel.Ruleset = core.StringPtr("testString")
				actionParametersModel.Rulesets = []string{"testString"}
				actionParametersModel.Response = actionParametersResponseModel

				// Construct an instance of the Logging model
				loggingModel := new(rulesetsv1.Logging)
				loggingModel.Enabled = core.BoolPtr(true)

				// Construct an instance of the Position model
				positionModel := new(rulesetsv1.Position)
				positionModel.Before = core.StringPtr("testString")
				positionModel.After = core.StringPtr("testString")
				positionModel.Index = core.Int64Ptr(int64(0))

				// Construct an instance of the RuleCreate model
				ruleCreateModel := new(rulesetsv1.RuleCreate)
				ruleCreateModel.Action = core.StringPtr("testString")
				ruleCreateModel.ActionParameters = actionParametersModel
				ruleCreateModel.Description = core.StringPtr("testString")
				ruleCreateModel.Enabled = core.BoolPtr(true)
				ruleCreateModel.Expression = core.StringPtr("ip.src ne 1.1.1.1")
				ruleCreateModel.ID = core.StringPtr("testString")
				ruleCreateModel.Logging = loggingModel
				ruleCreateModel.Ref = core.StringPtr("my_ref")
				ruleCreateModel.Position = positionModel

				// Construct an instance of the UpdateZoneEntrypointRulesetOptions model
				updateZoneEntrypointRulesetOptionsModel := new(rulesetsv1.UpdateZoneEntrypointRulesetOptions)
				updateZoneEntrypointRulesetOptionsModel.RulesetPhase = core.StringPtr("ddos_l4")
				updateZoneEntrypointRulesetOptionsModel.Description = core.StringPtr("Custom account ruleset")
				updateZoneEntrypointRulesetOptionsModel.Kind = core.StringPtr("managed")
				updateZoneEntrypointRulesetOptionsModel.Name = core.StringPtr("testString")
				updateZoneEntrypointRulesetOptionsModel.Phase = core.StringPtr("ddos_l4")
				updateZoneEntrypointRulesetOptionsModel.Rules = []rulesetsv1.RuleCreate{*ruleCreateModel}
				updateZoneEntrypointRulesetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := rulesetsService.UpdateZoneEntrypointRulesetWithContext(ctx, updateZoneEntrypointRulesetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				rulesetsService.DisableRetries()
				result, response, operationErr := rulesetsService.UpdateZoneEntrypointRuleset(updateZoneEntrypointRulesetOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = rulesetsService.UpdateZoneEntrypointRulesetWithContext(ctx, updateZoneEntrypointRulesetOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(updateZoneEntrypointRulesetPath))
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
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "messages": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "result": {"description": "Custom account ruleset", "id": "ID", "kind": "managed", "last_updated": "2000-01-01T00:00:00.000000Z", "name": "Name", "phase": "ddos_l4", "version": "1", "rules": [{"id": "ID", "version": "Version", "action": "Action", "action_parameters": {"id": "ID", "overrides": {"action": "Action", "enabled": false, "sensitivity_level": "high", "rules": [{"id": "ID", "enabled": false, "action": "Action", "sensitivity_level": "high"}], "categories": [{"category": "Category", "enabled": false, "action": "Action"}]}, "version": "Version", "ruleset": "Ruleset", "rulesets": ["Rulesets"], "response": {"content": "{\"success\": false, \"error\": \"you have been blocked\"}", "content_type": "application/json", "status_code": 400}}, "expression": "ip.src ne 1.1.1.1", "ref": "my_ref", "logging": {"enabled": true}, "last_updated": "2000-01-01T00:00:00.000000Z"}]}}`)
				}))
			})
			It(`Invoke UpdateZoneEntrypointRuleset successfully`, func() {
				rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(rulesetsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := rulesetsService.UpdateZoneEntrypointRuleset(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the RulesOverride model
				rulesOverrideModel := new(rulesetsv1.RulesOverride)
				rulesOverrideModel.ID = core.StringPtr("testString")
				rulesOverrideModel.Enabled = core.BoolPtr(true)
				rulesOverrideModel.Action = core.StringPtr("testString")
				rulesOverrideModel.SensitivityLevel = core.StringPtr("high")

				// Construct an instance of the CategoriesOverride model
				categoriesOverrideModel := new(rulesetsv1.CategoriesOverride)
				categoriesOverrideModel.Category = core.StringPtr("testString")
				categoriesOverrideModel.Enabled = core.BoolPtr(true)
				categoriesOverrideModel.Action = core.StringPtr("testString")

				// Construct an instance of the Overrides model
				overridesModel := new(rulesetsv1.Overrides)
				overridesModel.Action = core.StringPtr("testString")
				overridesModel.Enabled = core.BoolPtr(true)
				overridesModel.SensitivityLevel = core.StringPtr("high")
				overridesModel.Rules = []rulesetsv1.RulesOverride{*rulesOverrideModel}
				overridesModel.Categories = []rulesetsv1.CategoriesOverride{*categoriesOverrideModel}

				// Construct an instance of the ActionParametersResponse model
				actionParametersResponseModel := new(rulesetsv1.ActionParametersResponse)
				actionParametersResponseModel.Content = core.StringPtr(`{"success": false, "error": "you have been blocked"}`)
				actionParametersResponseModel.ContentType = core.StringPtr("application/json")
				actionParametersResponseModel.StatusCode = core.Int64Ptr(int64(400))

				// Construct an instance of the ActionParameters model
				actionParametersModel := new(rulesetsv1.ActionParameters)
				actionParametersModel.ID = core.StringPtr("testString")
				actionParametersModel.Overrides = overridesModel
				actionParametersModel.Version = core.StringPtr("testString")
				actionParametersModel.Ruleset = core.StringPtr("testString")
				actionParametersModel.Rulesets = []string{"testString"}
				actionParametersModel.Response = actionParametersResponseModel

				// Construct an instance of the Logging model
				loggingModel := new(rulesetsv1.Logging)
				loggingModel.Enabled = core.BoolPtr(true)

				// Construct an instance of the Position model
				positionModel := new(rulesetsv1.Position)
				positionModel.Before = core.StringPtr("testString")
				positionModel.After = core.StringPtr("testString")
				positionModel.Index = core.Int64Ptr(int64(0))

				// Construct an instance of the RuleCreate model
				ruleCreateModel := new(rulesetsv1.RuleCreate)
				ruleCreateModel.Action = core.StringPtr("testString")
				ruleCreateModel.ActionParameters = actionParametersModel
				ruleCreateModel.Description = core.StringPtr("testString")
				ruleCreateModel.Enabled = core.BoolPtr(true)
				ruleCreateModel.Expression = core.StringPtr("ip.src ne 1.1.1.1")
				ruleCreateModel.ID = core.StringPtr("testString")
				ruleCreateModel.Logging = loggingModel
				ruleCreateModel.Ref = core.StringPtr("my_ref")
				ruleCreateModel.Position = positionModel

				// Construct an instance of the UpdateZoneEntrypointRulesetOptions model
				updateZoneEntrypointRulesetOptionsModel := new(rulesetsv1.UpdateZoneEntrypointRulesetOptions)
				updateZoneEntrypointRulesetOptionsModel.RulesetPhase = core.StringPtr("ddos_l4")
				updateZoneEntrypointRulesetOptionsModel.Description = core.StringPtr("Custom account ruleset")
				updateZoneEntrypointRulesetOptionsModel.Kind = core.StringPtr("managed")
				updateZoneEntrypointRulesetOptionsModel.Name = core.StringPtr("testString")
				updateZoneEntrypointRulesetOptionsModel.Phase = core.StringPtr("ddos_l4")
				updateZoneEntrypointRulesetOptionsModel.Rules = []rulesetsv1.RuleCreate{*ruleCreateModel}
				updateZoneEntrypointRulesetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = rulesetsService.UpdateZoneEntrypointRuleset(updateZoneEntrypointRulesetOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateZoneEntrypointRuleset with error: Operation validation and request error`, func() {
				rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(rulesetsService).ToNot(BeNil())

				// Construct an instance of the RulesOverride model
				rulesOverrideModel := new(rulesetsv1.RulesOverride)
				rulesOverrideModel.ID = core.StringPtr("testString")
				rulesOverrideModel.Enabled = core.BoolPtr(true)
				rulesOverrideModel.Action = core.StringPtr("testString")
				rulesOverrideModel.SensitivityLevel = core.StringPtr("high")

				// Construct an instance of the CategoriesOverride model
				categoriesOverrideModel := new(rulesetsv1.CategoriesOverride)
				categoriesOverrideModel.Category = core.StringPtr("testString")
				categoriesOverrideModel.Enabled = core.BoolPtr(true)
				categoriesOverrideModel.Action = core.StringPtr("testString")

				// Construct an instance of the Overrides model
				overridesModel := new(rulesetsv1.Overrides)
				overridesModel.Action = core.StringPtr("testString")
				overridesModel.Enabled = core.BoolPtr(true)
				overridesModel.SensitivityLevel = core.StringPtr("high")
				overridesModel.Rules = []rulesetsv1.RulesOverride{*rulesOverrideModel}
				overridesModel.Categories = []rulesetsv1.CategoriesOverride{*categoriesOverrideModel}

				// Construct an instance of the ActionParametersResponse model
				actionParametersResponseModel := new(rulesetsv1.ActionParametersResponse)
				actionParametersResponseModel.Content = core.StringPtr(`{"success": false, "error": "you have been blocked"}`)
				actionParametersResponseModel.ContentType = core.StringPtr("application/json")
				actionParametersResponseModel.StatusCode = core.Int64Ptr(int64(400))

				// Construct an instance of the ActionParameters model
				actionParametersModel := new(rulesetsv1.ActionParameters)
				actionParametersModel.ID = core.StringPtr("testString")
				actionParametersModel.Overrides = overridesModel
				actionParametersModel.Version = core.StringPtr("testString")
				actionParametersModel.Ruleset = core.StringPtr("testString")
				actionParametersModel.Rulesets = []string{"testString"}
				actionParametersModel.Response = actionParametersResponseModel

				// Construct an instance of the Logging model
				loggingModel := new(rulesetsv1.Logging)
				loggingModel.Enabled = core.BoolPtr(true)

				// Construct an instance of the Position model
				positionModel := new(rulesetsv1.Position)
				positionModel.Before = core.StringPtr("testString")
				positionModel.After = core.StringPtr("testString")
				positionModel.Index = core.Int64Ptr(int64(0))

				// Construct an instance of the RuleCreate model
				ruleCreateModel := new(rulesetsv1.RuleCreate)
				ruleCreateModel.Action = core.StringPtr("testString")
				ruleCreateModel.ActionParameters = actionParametersModel
				ruleCreateModel.Description = core.StringPtr("testString")
				ruleCreateModel.Enabled = core.BoolPtr(true)
				ruleCreateModel.Expression = core.StringPtr("ip.src ne 1.1.1.1")
				ruleCreateModel.ID = core.StringPtr("testString")
				ruleCreateModel.Logging = loggingModel
				ruleCreateModel.Ref = core.StringPtr("my_ref")
				ruleCreateModel.Position = positionModel

				// Construct an instance of the UpdateZoneEntrypointRulesetOptions model
				updateZoneEntrypointRulesetOptionsModel := new(rulesetsv1.UpdateZoneEntrypointRulesetOptions)
				updateZoneEntrypointRulesetOptionsModel.RulesetPhase = core.StringPtr("ddos_l4")
				updateZoneEntrypointRulesetOptionsModel.Description = core.StringPtr("Custom account ruleset")
				updateZoneEntrypointRulesetOptionsModel.Kind = core.StringPtr("managed")
				updateZoneEntrypointRulesetOptionsModel.Name = core.StringPtr("testString")
				updateZoneEntrypointRulesetOptionsModel.Phase = core.StringPtr("ddos_l4")
				updateZoneEntrypointRulesetOptionsModel.Rules = []rulesetsv1.RuleCreate{*ruleCreateModel}
				updateZoneEntrypointRulesetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := rulesetsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := rulesetsService.UpdateZoneEntrypointRuleset(updateZoneEntrypointRulesetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateZoneEntrypointRulesetOptions model with no property values
				updateZoneEntrypointRulesetOptionsModelNew := new(rulesetsv1.UpdateZoneEntrypointRulesetOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = rulesetsService.UpdateZoneEntrypointRuleset(updateZoneEntrypointRulesetOptionsModelNew)
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
			It(`Invoke UpdateZoneEntrypointRuleset successfully`, func() {
				rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(rulesetsService).ToNot(BeNil())

				// Construct an instance of the RulesOverride model
				rulesOverrideModel := new(rulesetsv1.RulesOverride)
				rulesOverrideModel.ID = core.StringPtr("testString")
				rulesOverrideModel.Enabled = core.BoolPtr(true)
				rulesOverrideModel.Action = core.StringPtr("testString")
				rulesOverrideModel.SensitivityLevel = core.StringPtr("high")

				// Construct an instance of the CategoriesOverride model
				categoriesOverrideModel := new(rulesetsv1.CategoriesOverride)
				categoriesOverrideModel.Category = core.StringPtr("testString")
				categoriesOverrideModel.Enabled = core.BoolPtr(true)
				categoriesOverrideModel.Action = core.StringPtr("testString")

				// Construct an instance of the Overrides model
				overridesModel := new(rulesetsv1.Overrides)
				overridesModel.Action = core.StringPtr("testString")
				overridesModel.Enabled = core.BoolPtr(true)
				overridesModel.SensitivityLevel = core.StringPtr("high")
				overridesModel.Rules = []rulesetsv1.RulesOverride{*rulesOverrideModel}
				overridesModel.Categories = []rulesetsv1.CategoriesOverride{*categoriesOverrideModel}

				// Construct an instance of the ActionParametersResponse model
				actionParametersResponseModel := new(rulesetsv1.ActionParametersResponse)
				actionParametersResponseModel.Content = core.StringPtr(`{"success": false, "error": "you have been blocked"}`)
				actionParametersResponseModel.ContentType = core.StringPtr("application/json")
				actionParametersResponseModel.StatusCode = core.Int64Ptr(int64(400))

				// Construct an instance of the ActionParameters model
				actionParametersModel := new(rulesetsv1.ActionParameters)
				actionParametersModel.ID = core.StringPtr("testString")
				actionParametersModel.Overrides = overridesModel
				actionParametersModel.Version = core.StringPtr("testString")
				actionParametersModel.Ruleset = core.StringPtr("testString")
				actionParametersModel.Rulesets = []string{"testString"}
				actionParametersModel.Response = actionParametersResponseModel

				// Construct an instance of the Logging model
				loggingModel := new(rulesetsv1.Logging)
				loggingModel.Enabled = core.BoolPtr(true)

				// Construct an instance of the Position model
				positionModel := new(rulesetsv1.Position)
				positionModel.Before = core.StringPtr("testString")
				positionModel.After = core.StringPtr("testString")
				positionModel.Index = core.Int64Ptr(int64(0))

				// Construct an instance of the RuleCreate model
				ruleCreateModel := new(rulesetsv1.RuleCreate)
				ruleCreateModel.Action = core.StringPtr("testString")
				ruleCreateModel.ActionParameters = actionParametersModel
				ruleCreateModel.Description = core.StringPtr("testString")
				ruleCreateModel.Enabled = core.BoolPtr(true)
				ruleCreateModel.Expression = core.StringPtr("ip.src ne 1.1.1.1")
				ruleCreateModel.ID = core.StringPtr("testString")
				ruleCreateModel.Logging = loggingModel
				ruleCreateModel.Ref = core.StringPtr("my_ref")
				ruleCreateModel.Position = positionModel

				// Construct an instance of the UpdateZoneEntrypointRulesetOptions model
				updateZoneEntrypointRulesetOptionsModel := new(rulesetsv1.UpdateZoneEntrypointRulesetOptions)
				updateZoneEntrypointRulesetOptionsModel.RulesetPhase = core.StringPtr("ddos_l4")
				updateZoneEntrypointRulesetOptionsModel.Description = core.StringPtr("Custom account ruleset")
				updateZoneEntrypointRulesetOptionsModel.Kind = core.StringPtr("managed")
				updateZoneEntrypointRulesetOptionsModel.Name = core.StringPtr("testString")
				updateZoneEntrypointRulesetOptionsModel.Phase = core.StringPtr("ddos_l4")
				updateZoneEntrypointRulesetOptionsModel.Rules = []rulesetsv1.RuleCreate{*ruleCreateModel}
				updateZoneEntrypointRulesetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := rulesetsService.UpdateZoneEntrypointRuleset(updateZoneEntrypointRulesetOptionsModel)
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
	Describe(`GetZoneEntryPointRulesetVersions(getZoneEntryPointRulesetVersionsOptions *GetZoneEntryPointRulesetVersionsOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getZoneEntryPointRulesetVersionsPath := "/v1/testString/zones/testString/rulesets/phases/ddos_l4/entrypoint/versions"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getZoneEntryPointRulesetVersionsPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetZoneEntryPointRulesetVersions with error: Operation response processing error`, func() {
				rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(rulesetsService).ToNot(BeNil())

				// Construct an instance of the GetZoneEntryPointRulesetVersionsOptions model
				getZoneEntryPointRulesetVersionsOptionsModel := new(rulesetsv1.GetZoneEntryPointRulesetVersionsOptions)
				getZoneEntryPointRulesetVersionsOptionsModel.RulesetPhase = core.StringPtr("ddos_l4")
				getZoneEntryPointRulesetVersionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := rulesetsService.GetZoneEntryPointRulesetVersions(getZoneEntryPointRulesetVersionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				rulesetsService.EnableRetries(0, 0)
				result, response, operationErr = rulesetsService.GetZoneEntryPointRulesetVersions(getZoneEntryPointRulesetVersionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetZoneEntryPointRulesetVersions(getZoneEntryPointRulesetVersionsOptions *GetZoneEntryPointRulesetVersionsOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getZoneEntryPointRulesetVersionsPath := "/v1/testString/zones/testString/rulesets/phases/ddos_l4/entrypoint/versions"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getZoneEntryPointRulesetVersionsPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "messages": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "result": [{"description": "Custom account ruleset", "id": "ID", "kind": "managed", "last_updated": "2000-01-01T00:00:00.000000Z", "name": "Name", "phase": "ddos_l4", "version": "1"}]}`)
				}))
			})
			It(`Invoke GetZoneEntryPointRulesetVersions successfully with retries`, func() {
				rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(rulesetsService).ToNot(BeNil())
				rulesetsService.EnableRetries(0, 0)

				// Construct an instance of the GetZoneEntryPointRulesetVersionsOptions model
				getZoneEntryPointRulesetVersionsOptionsModel := new(rulesetsv1.GetZoneEntryPointRulesetVersionsOptions)
				getZoneEntryPointRulesetVersionsOptionsModel.RulesetPhase = core.StringPtr("ddos_l4")
				getZoneEntryPointRulesetVersionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := rulesetsService.GetZoneEntryPointRulesetVersionsWithContext(ctx, getZoneEntryPointRulesetVersionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				rulesetsService.DisableRetries()
				result, response, operationErr := rulesetsService.GetZoneEntryPointRulesetVersions(getZoneEntryPointRulesetVersionsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = rulesetsService.GetZoneEntryPointRulesetVersionsWithContext(ctx, getZoneEntryPointRulesetVersionsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getZoneEntryPointRulesetVersionsPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "messages": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "result": [{"description": "Custom account ruleset", "id": "ID", "kind": "managed", "last_updated": "2000-01-01T00:00:00.000000Z", "name": "Name", "phase": "ddos_l4", "version": "1"}]}`)
				}))
			})
			It(`Invoke GetZoneEntryPointRulesetVersions successfully`, func() {
				rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(rulesetsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := rulesetsService.GetZoneEntryPointRulesetVersions(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetZoneEntryPointRulesetVersionsOptions model
				getZoneEntryPointRulesetVersionsOptionsModel := new(rulesetsv1.GetZoneEntryPointRulesetVersionsOptions)
				getZoneEntryPointRulesetVersionsOptionsModel.RulesetPhase = core.StringPtr("ddos_l4")
				getZoneEntryPointRulesetVersionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = rulesetsService.GetZoneEntryPointRulesetVersions(getZoneEntryPointRulesetVersionsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetZoneEntryPointRulesetVersions with error: Operation validation and request error`, func() {
				rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(rulesetsService).ToNot(BeNil())

				// Construct an instance of the GetZoneEntryPointRulesetVersionsOptions model
				getZoneEntryPointRulesetVersionsOptionsModel := new(rulesetsv1.GetZoneEntryPointRulesetVersionsOptions)
				getZoneEntryPointRulesetVersionsOptionsModel.RulesetPhase = core.StringPtr("ddos_l4")
				getZoneEntryPointRulesetVersionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := rulesetsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := rulesetsService.GetZoneEntryPointRulesetVersions(getZoneEntryPointRulesetVersionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetZoneEntryPointRulesetVersionsOptions model with no property values
				getZoneEntryPointRulesetVersionsOptionsModelNew := new(rulesetsv1.GetZoneEntryPointRulesetVersionsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = rulesetsService.GetZoneEntryPointRulesetVersions(getZoneEntryPointRulesetVersionsOptionsModelNew)
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
			It(`Invoke GetZoneEntryPointRulesetVersions successfully`, func() {
				rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(rulesetsService).ToNot(BeNil())

				// Construct an instance of the GetZoneEntryPointRulesetVersionsOptions model
				getZoneEntryPointRulesetVersionsOptionsModel := new(rulesetsv1.GetZoneEntryPointRulesetVersionsOptions)
				getZoneEntryPointRulesetVersionsOptionsModel.RulesetPhase = core.StringPtr("ddos_l4")
				getZoneEntryPointRulesetVersionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := rulesetsService.GetZoneEntryPointRulesetVersions(getZoneEntryPointRulesetVersionsOptionsModel)
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
	Describe(`GetZoneEntryPointRulesetVersion(getZoneEntryPointRulesetVersionOptions *GetZoneEntryPointRulesetVersionOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getZoneEntryPointRulesetVersionPath := "/v1/testString/zones/testString/rulesets/phases/ddos_l4/entrypoint/versions/1"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getZoneEntryPointRulesetVersionPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetZoneEntryPointRulesetVersion with error: Operation response processing error`, func() {
				rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(rulesetsService).ToNot(BeNil())

				// Construct an instance of the GetZoneEntryPointRulesetVersionOptions model
				getZoneEntryPointRulesetVersionOptionsModel := new(rulesetsv1.GetZoneEntryPointRulesetVersionOptions)
				getZoneEntryPointRulesetVersionOptionsModel.RulesetPhase = core.StringPtr("ddos_l4")
				getZoneEntryPointRulesetVersionOptionsModel.RulesetVersion = core.StringPtr("1")
				getZoneEntryPointRulesetVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := rulesetsService.GetZoneEntryPointRulesetVersion(getZoneEntryPointRulesetVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				rulesetsService.EnableRetries(0, 0)
				result, response, operationErr = rulesetsService.GetZoneEntryPointRulesetVersion(getZoneEntryPointRulesetVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetZoneEntryPointRulesetVersion(getZoneEntryPointRulesetVersionOptions *GetZoneEntryPointRulesetVersionOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getZoneEntryPointRulesetVersionPath := "/v1/testString/zones/testString/rulesets/phases/ddos_l4/entrypoint/versions/1"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getZoneEntryPointRulesetVersionPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "messages": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "result": {"description": "Custom account ruleset", "id": "ID", "kind": "managed", "last_updated": "2000-01-01T00:00:00.000000Z", "name": "Name", "phase": "ddos_l4", "version": "1", "rules": [{"id": "ID", "version": "Version", "action": "Action", "action_parameters": {"id": "ID", "overrides": {"action": "Action", "enabled": false, "sensitivity_level": "high", "rules": [{"id": "ID", "enabled": false, "action": "Action", "sensitivity_level": "high"}], "categories": [{"category": "Category", "enabled": false, "action": "Action"}]}, "version": "Version", "ruleset": "Ruleset", "rulesets": ["Rulesets"], "response": {"content": "{\"success\": false, \"error\": \"you have been blocked\"}", "content_type": "application/json", "status_code": 400}}, "expression": "ip.src ne 1.1.1.1", "ref": "my_ref", "logging": {"enabled": true}, "last_updated": "2000-01-01T00:00:00.000000Z"}]}}`)
				}))
			})
			It(`Invoke GetZoneEntryPointRulesetVersion successfully with retries`, func() {
				rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(rulesetsService).ToNot(BeNil())
				rulesetsService.EnableRetries(0, 0)

				// Construct an instance of the GetZoneEntryPointRulesetVersionOptions model
				getZoneEntryPointRulesetVersionOptionsModel := new(rulesetsv1.GetZoneEntryPointRulesetVersionOptions)
				getZoneEntryPointRulesetVersionOptionsModel.RulesetPhase = core.StringPtr("ddos_l4")
				getZoneEntryPointRulesetVersionOptionsModel.RulesetVersion = core.StringPtr("1")
				getZoneEntryPointRulesetVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := rulesetsService.GetZoneEntryPointRulesetVersionWithContext(ctx, getZoneEntryPointRulesetVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				rulesetsService.DisableRetries()
				result, response, operationErr := rulesetsService.GetZoneEntryPointRulesetVersion(getZoneEntryPointRulesetVersionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = rulesetsService.GetZoneEntryPointRulesetVersionWithContext(ctx, getZoneEntryPointRulesetVersionOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getZoneEntryPointRulesetVersionPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "messages": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "result": {"description": "Custom account ruleset", "id": "ID", "kind": "managed", "last_updated": "2000-01-01T00:00:00.000000Z", "name": "Name", "phase": "ddos_l4", "version": "1", "rules": [{"id": "ID", "version": "Version", "action": "Action", "action_parameters": {"id": "ID", "overrides": {"action": "Action", "enabled": false, "sensitivity_level": "high", "rules": [{"id": "ID", "enabled": false, "action": "Action", "sensitivity_level": "high"}], "categories": [{"category": "Category", "enabled": false, "action": "Action"}]}, "version": "Version", "ruleset": "Ruleset", "rulesets": ["Rulesets"], "response": {"content": "{\"success\": false, \"error\": \"you have been blocked\"}", "content_type": "application/json", "status_code": 400}}, "expression": "ip.src ne 1.1.1.1", "ref": "my_ref", "logging": {"enabled": true}, "last_updated": "2000-01-01T00:00:00.000000Z"}]}}`)
				}))
			})
			It(`Invoke GetZoneEntryPointRulesetVersion successfully`, func() {
				rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(rulesetsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := rulesetsService.GetZoneEntryPointRulesetVersion(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetZoneEntryPointRulesetVersionOptions model
				getZoneEntryPointRulesetVersionOptionsModel := new(rulesetsv1.GetZoneEntryPointRulesetVersionOptions)
				getZoneEntryPointRulesetVersionOptionsModel.RulesetPhase = core.StringPtr("ddos_l4")
				getZoneEntryPointRulesetVersionOptionsModel.RulesetVersion = core.StringPtr("1")
				getZoneEntryPointRulesetVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = rulesetsService.GetZoneEntryPointRulesetVersion(getZoneEntryPointRulesetVersionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetZoneEntryPointRulesetVersion with error: Operation validation and request error`, func() {
				rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(rulesetsService).ToNot(BeNil())

				// Construct an instance of the GetZoneEntryPointRulesetVersionOptions model
				getZoneEntryPointRulesetVersionOptionsModel := new(rulesetsv1.GetZoneEntryPointRulesetVersionOptions)
				getZoneEntryPointRulesetVersionOptionsModel.RulesetPhase = core.StringPtr("ddos_l4")
				getZoneEntryPointRulesetVersionOptionsModel.RulesetVersion = core.StringPtr("1")
				getZoneEntryPointRulesetVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := rulesetsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := rulesetsService.GetZoneEntryPointRulesetVersion(getZoneEntryPointRulesetVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetZoneEntryPointRulesetVersionOptions model with no property values
				getZoneEntryPointRulesetVersionOptionsModelNew := new(rulesetsv1.GetZoneEntryPointRulesetVersionOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = rulesetsService.GetZoneEntryPointRulesetVersion(getZoneEntryPointRulesetVersionOptionsModelNew)
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
			It(`Invoke GetZoneEntryPointRulesetVersion successfully`, func() {
				rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(rulesetsService).ToNot(BeNil())

				// Construct an instance of the GetZoneEntryPointRulesetVersionOptions model
				getZoneEntryPointRulesetVersionOptionsModel := new(rulesetsv1.GetZoneEntryPointRulesetVersionOptions)
				getZoneEntryPointRulesetVersionOptionsModel.RulesetPhase = core.StringPtr("ddos_l4")
				getZoneEntryPointRulesetVersionOptionsModel.RulesetVersion = core.StringPtr("1")
				getZoneEntryPointRulesetVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := rulesetsService.GetZoneEntryPointRulesetVersion(getZoneEntryPointRulesetVersionOptionsModel)
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
	Describe(`CreateZoneRulesetRule(createZoneRulesetRuleOptions *CreateZoneRulesetRuleOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		createZoneRulesetRulePath := "/v1/testString/zones/testString/rulesets/testString/rules"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createZoneRulesetRulePath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateZoneRulesetRule with error: Operation response processing error`, func() {
				rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(rulesetsService).ToNot(BeNil())

				// Construct an instance of the RulesOverride model
				rulesOverrideModel := new(rulesetsv1.RulesOverride)
				rulesOverrideModel.ID = core.StringPtr("testString")
				rulesOverrideModel.Enabled = core.BoolPtr(true)
				rulesOverrideModel.Action = core.StringPtr("testString")
				rulesOverrideModel.SensitivityLevel = core.StringPtr("high")

				// Construct an instance of the CategoriesOverride model
				categoriesOverrideModel := new(rulesetsv1.CategoriesOverride)
				categoriesOverrideModel.Category = core.StringPtr("testString")
				categoriesOverrideModel.Enabled = core.BoolPtr(true)
				categoriesOverrideModel.Action = core.StringPtr("testString")

				// Construct an instance of the Overrides model
				overridesModel := new(rulesetsv1.Overrides)
				overridesModel.Action = core.StringPtr("testString")
				overridesModel.Enabled = core.BoolPtr(true)
				overridesModel.SensitivityLevel = core.StringPtr("high")
				overridesModel.Rules = []rulesetsv1.RulesOverride{*rulesOverrideModel}
				overridesModel.Categories = []rulesetsv1.CategoriesOverride{*categoriesOverrideModel}

				// Construct an instance of the ActionParametersResponse model
				actionParametersResponseModel := new(rulesetsv1.ActionParametersResponse)
				actionParametersResponseModel.Content = core.StringPtr(`{"success": false, "error": "you have been blocked"}`)
				actionParametersResponseModel.ContentType = core.StringPtr("application/json")
				actionParametersResponseModel.StatusCode = core.Int64Ptr(int64(400))

				// Construct an instance of the ActionParameters model
				actionParametersModel := new(rulesetsv1.ActionParameters)
				actionParametersModel.ID = core.StringPtr("testString")
				actionParametersModel.Overrides = overridesModel
				actionParametersModel.Version = core.StringPtr("testString")
				actionParametersModel.Ruleset = core.StringPtr("testString")
				actionParametersModel.Rulesets = []string{"testString"}
				actionParametersModel.Response = actionParametersResponseModel

				// Construct an instance of the Logging model
				loggingModel := new(rulesetsv1.Logging)
				loggingModel.Enabled = core.BoolPtr(true)

				// Construct an instance of the Position model
				positionModel := new(rulesetsv1.Position)
				positionModel.Before = core.StringPtr("testString")
				positionModel.After = core.StringPtr("testString")
				positionModel.Index = core.Int64Ptr(int64(0))

				// Construct an instance of the CreateZoneRulesetRuleOptions model
				createZoneRulesetRuleOptionsModel := new(rulesetsv1.CreateZoneRulesetRuleOptions)
				createZoneRulesetRuleOptionsModel.RulesetID = core.StringPtr("testString")
				createZoneRulesetRuleOptionsModel.Action = core.StringPtr("testString")
				createZoneRulesetRuleOptionsModel.ActionParameters = actionParametersModel
				createZoneRulesetRuleOptionsModel.Description = core.StringPtr("testString")
				createZoneRulesetRuleOptionsModel.Enabled = core.BoolPtr(true)
				createZoneRulesetRuleOptionsModel.Expression = core.StringPtr("ip.src ne 1.1.1.1")
				createZoneRulesetRuleOptionsModel.ID = core.StringPtr("testString")
				createZoneRulesetRuleOptionsModel.Logging = loggingModel
				createZoneRulesetRuleOptionsModel.Ref = core.StringPtr("my_ref")
				createZoneRulesetRuleOptionsModel.Position = positionModel
				createZoneRulesetRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := rulesetsService.CreateZoneRulesetRule(createZoneRulesetRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				rulesetsService.EnableRetries(0, 0)
				result, response, operationErr = rulesetsService.CreateZoneRulesetRule(createZoneRulesetRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateZoneRulesetRule(createZoneRulesetRuleOptions *CreateZoneRulesetRuleOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		createZoneRulesetRulePath := "/v1/testString/zones/testString/rulesets/testString/rules"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createZoneRulesetRulePath))
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
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "messages": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "result": {"id": "ID", "version": "Version", "action": "Action", "action_parameters": {"id": "ID", "overrides": {"action": "Action", "enabled": false, "sensitivity_level": "high", "rules": [{"id": "ID", "enabled": false, "action": "Action", "sensitivity_level": "high"}], "categories": [{"category": "Category", "enabled": false, "action": "Action"}]}, "version": "Version", "ruleset": "Ruleset", "rulesets": ["Rulesets"], "response": {"content": "{\"success\": false, \"error\": \"you have been blocked\"}", "content_type": "application/json", "status_code": 400}}, "expression": "ip.src ne 1.1.1.1", "ref": "my_ref", "logging": {"enabled": true}, "last_updated": "2000-01-01T00:00:00.000000Z"}}`)
				}))
			})
			It(`Invoke CreateZoneRulesetRule successfully with retries`, func() {
				rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(rulesetsService).ToNot(BeNil())
				rulesetsService.EnableRetries(0, 0)

				// Construct an instance of the RulesOverride model
				rulesOverrideModel := new(rulesetsv1.RulesOverride)
				rulesOverrideModel.ID = core.StringPtr("testString")
				rulesOverrideModel.Enabled = core.BoolPtr(true)
				rulesOverrideModel.Action = core.StringPtr("testString")
				rulesOverrideModel.SensitivityLevel = core.StringPtr("high")

				// Construct an instance of the CategoriesOverride model
				categoriesOverrideModel := new(rulesetsv1.CategoriesOverride)
				categoriesOverrideModel.Category = core.StringPtr("testString")
				categoriesOverrideModel.Enabled = core.BoolPtr(true)
				categoriesOverrideModel.Action = core.StringPtr("testString")

				// Construct an instance of the Overrides model
				overridesModel := new(rulesetsv1.Overrides)
				overridesModel.Action = core.StringPtr("testString")
				overridesModel.Enabled = core.BoolPtr(true)
				overridesModel.SensitivityLevel = core.StringPtr("high")
				overridesModel.Rules = []rulesetsv1.RulesOverride{*rulesOverrideModel}
				overridesModel.Categories = []rulesetsv1.CategoriesOverride{*categoriesOverrideModel}

				// Construct an instance of the ActionParametersResponse model
				actionParametersResponseModel := new(rulesetsv1.ActionParametersResponse)
				actionParametersResponseModel.Content = core.StringPtr(`{"success": false, "error": "you have been blocked"}`)
				actionParametersResponseModel.ContentType = core.StringPtr("application/json")
				actionParametersResponseModel.StatusCode = core.Int64Ptr(int64(400))

				// Construct an instance of the ActionParameters model
				actionParametersModel := new(rulesetsv1.ActionParameters)
				actionParametersModel.ID = core.StringPtr("testString")
				actionParametersModel.Overrides = overridesModel
				actionParametersModel.Version = core.StringPtr("testString")
				actionParametersModel.Ruleset = core.StringPtr("testString")
				actionParametersModel.Rulesets = []string{"testString"}
				actionParametersModel.Response = actionParametersResponseModel

				// Construct an instance of the Logging model
				loggingModel := new(rulesetsv1.Logging)
				loggingModel.Enabled = core.BoolPtr(true)

				// Construct an instance of the Position model
				positionModel := new(rulesetsv1.Position)
				positionModel.Before = core.StringPtr("testString")
				positionModel.After = core.StringPtr("testString")
				positionModel.Index = core.Int64Ptr(int64(0))

				// Construct an instance of the CreateZoneRulesetRuleOptions model
				createZoneRulesetRuleOptionsModel := new(rulesetsv1.CreateZoneRulesetRuleOptions)
				createZoneRulesetRuleOptionsModel.RulesetID = core.StringPtr("testString")
				createZoneRulesetRuleOptionsModel.Action = core.StringPtr("testString")
				createZoneRulesetRuleOptionsModel.ActionParameters = actionParametersModel
				createZoneRulesetRuleOptionsModel.Description = core.StringPtr("testString")
				createZoneRulesetRuleOptionsModel.Enabled = core.BoolPtr(true)
				createZoneRulesetRuleOptionsModel.Expression = core.StringPtr("ip.src ne 1.1.1.1")
				createZoneRulesetRuleOptionsModel.ID = core.StringPtr("testString")
				createZoneRulesetRuleOptionsModel.Logging = loggingModel
				createZoneRulesetRuleOptionsModel.Ref = core.StringPtr("my_ref")
				createZoneRulesetRuleOptionsModel.Position = positionModel
				createZoneRulesetRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := rulesetsService.CreateZoneRulesetRuleWithContext(ctx, createZoneRulesetRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				rulesetsService.DisableRetries()
				result, response, operationErr := rulesetsService.CreateZoneRulesetRule(createZoneRulesetRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = rulesetsService.CreateZoneRulesetRuleWithContext(ctx, createZoneRulesetRuleOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createZoneRulesetRulePath))
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
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "messages": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "result": {"id": "ID", "version": "Version", "action": "Action", "action_parameters": {"id": "ID", "overrides": {"action": "Action", "enabled": false, "sensitivity_level": "high", "rules": [{"id": "ID", "enabled": false, "action": "Action", "sensitivity_level": "high"}], "categories": [{"category": "Category", "enabled": false, "action": "Action"}]}, "version": "Version", "ruleset": "Ruleset", "rulesets": ["Rulesets"], "response": {"content": "{\"success\": false, \"error\": \"you have been blocked\"}", "content_type": "application/json", "status_code": 400}}, "expression": "ip.src ne 1.1.1.1", "ref": "my_ref", "logging": {"enabled": true}, "last_updated": "2000-01-01T00:00:00.000000Z"}}`)
				}))
			})
			It(`Invoke CreateZoneRulesetRule successfully`, func() {
				rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(rulesetsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := rulesetsService.CreateZoneRulesetRule(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the RulesOverride model
				rulesOverrideModel := new(rulesetsv1.RulesOverride)
				rulesOverrideModel.ID = core.StringPtr("testString")
				rulesOverrideModel.Enabled = core.BoolPtr(true)
				rulesOverrideModel.Action = core.StringPtr("testString")
				rulesOverrideModel.SensitivityLevel = core.StringPtr("high")

				// Construct an instance of the CategoriesOverride model
				categoriesOverrideModel := new(rulesetsv1.CategoriesOverride)
				categoriesOverrideModel.Category = core.StringPtr("testString")
				categoriesOverrideModel.Enabled = core.BoolPtr(true)
				categoriesOverrideModel.Action = core.StringPtr("testString")

				// Construct an instance of the Overrides model
				overridesModel := new(rulesetsv1.Overrides)
				overridesModel.Action = core.StringPtr("testString")
				overridesModel.Enabled = core.BoolPtr(true)
				overridesModel.SensitivityLevel = core.StringPtr("high")
				overridesModel.Rules = []rulesetsv1.RulesOverride{*rulesOverrideModel}
				overridesModel.Categories = []rulesetsv1.CategoriesOverride{*categoriesOverrideModel}

				// Construct an instance of the ActionParametersResponse model
				actionParametersResponseModel := new(rulesetsv1.ActionParametersResponse)
				actionParametersResponseModel.Content = core.StringPtr(`{"success": false, "error": "you have been blocked"}`)
				actionParametersResponseModel.ContentType = core.StringPtr("application/json")
				actionParametersResponseModel.StatusCode = core.Int64Ptr(int64(400))

				// Construct an instance of the ActionParameters model
				actionParametersModel := new(rulesetsv1.ActionParameters)
				actionParametersModel.ID = core.StringPtr("testString")
				actionParametersModel.Overrides = overridesModel
				actionParametersModel.Version = core.StringPtr("testString")
				actionParametersModel.Ruleset = core.StringPtr("testString")
				actionParametersModel.Rulesets = []string{"testString"}
				actionParametersModel.Response = actionParametersResponseModel

				// Construct an instance of the Logging model
				loggingModel := new(rulesetsv1.Logging)
				loggingModel.Enabled = core.BoolPtr(true)

				// Construct an instance of the Position model
				positionModel := new(rulesetsv1.Position)
				positionModel.Before = core.StringPtr("testString")
				positionModel.After = core.StringPtr("testString")
				positionModel.Index = core.Int64Ptr(int64(0))

				// Construct an instance of the CreateZoneRulesetRuleOptions model
				createZoneRulesetRuleOptionsModel := new(rulesetsv1.CreateZoneRulesetRuleOptions)
				createZoneRulesetRuleOptionsModel.RulesetID = core.StringPtr("testString")
				createZoneRulesetRuleOptionsModel.Action = core.StringPtr("testString")
				createZoneRulesetRuleOptionsModel.ActionParameters = actionParametersModel
				createZoneRulesetRuleOptionsModel.Description = core.StringPtr("testString")
				createZoneRulesetRuleOptionsModel.Enabled = core.BoolPtr(true)
				createZoneRulesetRuleOptionsModel.Expression = core.StringPtr("ip.src ne 1.1.1.1")
				createZoneRulesetRuleOptionsModel.ID = core.StringPtr("testString")
				createZoneRulesetRuleOptionsModel.Logging = loggingModel
				createZoneRulesetRuleOptionsModel.Ref = core.StringPtr("my_ref")
				createZoneRulesetRuleOptionsModel.Position = positionModel
				createZoneRulesetRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = rulesetsService.CreateZoneRulesetRule(createZoneRulesetRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateZoneRulesetRule with error: Operation validation and request error`, func() {
				rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(rulesetsService).ToNot(BeNil())

				// Construct an instance of the RulesOverride model
				rulesOverrideModel := new(rulesetsv1.RulesOverride)
				rulesOverrideModel.ID = core.StringPtr("testString")
				rulesOverrideModel.Enabled = core.BoolPtr(true)
				rulesOverrideModel.Action = core.StringPtr("testString")
				rulesOverrideModel.SensitivityLevel = core.StringPtr("high")

				// Construct an instance of the CategoriesOverride model
				categoriesOverrideModel := new(rulesetsv1.CategoriesOverride)
				categoriesOverrideModel.Category = core.StringPtr("testString")
				categoriesOverrideModel.Enabled = core.BoolPtr(true)
				categoriesOverrideModel.Action = core.StringPtr("testString")

				// Construct an instance of the Overrides model
				overridesModel := new(rulesetsv1.Overrides)
				overridesModel.Action = core.StringPtr("testString")
				overridesModel.Enabled = core.BoolPtr(true)
				overridesModel.SensitivityLevel = core.StringPtr("high")
				overridesModel.Rules = []rulesetsv1.RulesOverride{*rulesOverrideModel}
				overridesModel.Categories = []rulesetsv1.CategoriesOverride{*categoriesOverrideModel}

				// Construct an instance of the ActionParametersResponse model
				actionParametersResponseModel := new(rulesetsv1.ActionParametersResponse)
				actionParametersResponseModel.Content = core.StringPtr(`{"success": false, "error": "you have been blocked"}`)
				actionParametersResponseModel.ContentType = core.StringPtr("application/json")
				actionParametersResponseModel.StatusCode = core.Int64Ptr(int64(400))

				// Construct an instance of the ActionParameters model
				actionParametersModel := new(rulesetsv1.ActionParameters)
				actionParametersModel.ID = core.StringPtr("testString")
				actionParametersModel.Overrides = overridesModel
				actionParametersModel.Version = core.StringPtr("testString")
				actionParametersModel.Ruleset = core.StringPtr("testString")
				actionParametersModel.Rulesets = []string{"testString"}
				actionParametersModel.Response = actionParametersResponseModel

				// Construct an instance of the Logging model
				loggingModel := new(rulesetsv1.Logging)
				loggingModel.Enabled = core.BoolPtr(true)

				// Construct an instance of the Position model
				positionModel := new(rulesetsv1.Position)
				positionModel.Before = core.StringPtr("testString")
				positionModel.After = core.StringPtr("testString")
				positionModel.Index = core.Int64Ptr(int64(0))

				// Construct an instance of the CreateZoneRulesetRuleOptions model
				createZoneRulesetRuleOptionsModel := new(rulesetsv1.CreateZoneRulesetRuleOptions)
				createZoneRulesetRuleOptionsModel.RulesetID = core.StringPtr("testString")
				createZoneRulesetRuleOptionsModel.Action = core.StringPtr("testString")
				createZoneRulesetRuleOptionsModel.ActionParameters = actionParametersModel
				createZoneRulesetRuleOptionsModel.Description = core.StringPtr("testString")
				createZoneRulesetRuleOptionsModel.Enabled = core.BoolPtr(true)
				createZoneRulesetRuleOptionsModel.Expression = core.StringPtr("ip.src ne 1.1.1.1")
				createZoneRulesetRuleOptionsModel.ID = core.StringPtr("testString")
				createZoneRulesetRuleOptionsModel.Logging = loggingModel
				createZoneRulesetRuleOptionsModel.Ref = core.StringPtr("my_ref")
				createZoneRulesetRuleOptionsModel.Position = positionModel
				createZoneRulesetRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := rulesetsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := rulesetsService.CreateZoneRulesetRule(createZoneRulesetRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateZoneRulesetRuleOptions model with no property values
				createZoneRulesetRuleOptionsModelNew := new(rulesetsv1.CreateZoneRulesetRuleOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = rulesetsService.CreateZoneRulesetRule(createZoneRulesetRuleOptionsModelNew)
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
			It(`Invoke CreateZoneRulesetRule successfully`, func() {
				rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(rulesetsService).ToNot(BeNil())

				// Construct an instance of the RulesOverride model
				rulesOverrideModel := new(rulesetsv1.RulesOverride)
				rulesOverrideModel.ID = core.StringPtr("testString")
				rulesOverrideModel.Enabled = core.BoolPtr(true)
				rulesOverrideModel.Action = core.StringPtr("testString")
				rulesOverrideModel.SensitivityLevel = core.StringPtr("high")

				// Construct an instance of the CategoriesOverride model
				categoriesOverrideModel := new(rulesetsv1.CategoriesOverride)
				categoriesOverrideModel.Category = core.StringPtr("testString")
				categoriesOverrideModel.Enabled = core.BoolPtr(true)
				categoriesOverrideModel.Action = core.StringPtr("testString")

				// Construct an instance of the Overrides model
				overridesModel := new(rulesetsv1.Overrides)
				overridesModel.Action = core.StringPtr("testString")
				overridesModel.Enabled = core.BoolPtr(true)
				overridesModel.SensitivityLevel = core.StringPtr("high")
				overridesModel.Rules = []rulesetsv1.RulesOverride{*rulesOverrideModel}
				overridesModel.Categories = []rulesetsv1.CategoriesOverride{*categoriesOverrideModel}

				// Construct an instance of the ActionParametersResponse model
				actionParametersResponseModel := new(rulesetsv1.ActionParametersResponse)
				actionParametersResponseModel.Content = core.StringPtr(`{"success": false, "error": "you have been blocked"}`)
				actionParametersResponseModel.ContentType = core.StringPtr("application/json")
				actionParametersResponseModel.StatusCode = core.Int64Ptr(int64(400))

				// Construct an instance of the ActionParameters model
				actionParametersModel := new(rulesetsv1.ActionParameters)
				actionParametersModel.ID = core.StringPtr("testString")
				actionParametersModel.Overrides = overridesModel
				actionParametersModel.Version = core.StringPtr("testString")
				actionParametersModel.Ruleset = core.StringPtr("testString")
				actionParametersModel.Rulesets = []string{"testString"}
				actionParametersModel.Response = actionParametersResponseModel

				// Construct an instance of the Logging model
				loggingModel := new(rulesetsv1.Logging)
				loggingModel.Enabled = core.BoolPtr(true)

				// Construct an instance of the Position model
				positionModel := new(rulesetsv1.Position)
				positionModel.Before = core.StringPtr("testString")
				positionModel.After = core.StringPtr("testString")
				positionModel.Index = core.Int64Ptr(int64(0))

				// Construct an instance of the CreateZoneRulesetRuleOptions model
				createZoneRulesetRuleOptionsModel := new(rulesetsv1.CreateZoneRulesetRuleOptions)
				createZoneRulesetRuleOptionsModel.RulesetID = core.StringPtr("testString")
				createZoneRulesetRuleOptionsModel.Action = core.StringPtr("testString")
				createZoneRulesetRuleOptionsModel.ActionParameters = actionParametersModel
				createZoneRulesetRuleOptionsModel.Description = core.StringPtr("testString")
				createZoneRulesetRuleOptionsModel.Enabled = core.BoolPtr(true)
				createZoneRulesetRuleOptionsModel.Expression = core.StringPtr("ip.src ne 1.1.1.1")
				createZoneRulesetRuleOptionsModel.ID = core.StringPtr("testString")
				createZoneRulesetRuleOptionsModel.Logging = loggingModel
				createZoneRulesetRuleOptionsModel.Ref = core.StringPtr("my_ref")
				createZoneRulesetRuleOptionsModel.Position = positionModel
				createZoneRulesetRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := rulesetsService.CreateZoneRulesetRule(createZoneRulesetRuleOptionsModel)
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
	Describe(`UpdateZoneRulesetRule(updateZoneRulesetRuleOptions *UpdateZoneRulesetRuleOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		updateZoneRulesetRulePath := "/v1/testString/zones/testString/rulesets/testString/rules/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateZoneRulesetRulePath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateZoneRulesetRule with error: Operation response processing error`, func() {
				rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(rulesetsService).ToNot(BeNil())

				// Construct an instance of the RulesOverride model
				rulesOverrideModel := new(rulesetsv1.RulesOverride)
				rulesOverrideModel.ID = core.StringPtr("testString")
				rulesOverrideModel.Enabled = core.BoolPtr(true)
				rulesOverrideModel.Action = core.StringPtr("testString")
				rulesOverrideModel.SensitivityLevel = core.StringPtr("high")

				// Construct an instance of the CategoriesOverride model
				categoriesOverrideModel := new(rulesetsv1.CategoriesOverride)
				categoriesOverrideModel.Category = core.StringPtr("testString")
				categoriesOverrideModel.Enabled = core.BoolPtr(true)
				categoriesOverrideModel.Action = core.StringPtr("testString")

				// Construct an instance of the Overrides model
				overridesModel := new(rulesetsv1.Overrides)
				overridesModel.Action = core.StringPtr("testString")
				overridesModel.Enabled = core.BoolPtr(true)
				overridesModel.SensitivityLevel = core.StringPtr("high")
				overridesModel.Rules = []rulesetsv1.RulesOverride{*rulesOverrideModel}
				overridesModel.Categories = []rulesetsv1.CategoriesOverride{*categoriesOverrideModel}

				// Construct an instance of the ActionParametersResponse model
				actionParametersResponseModel := new(rulesetsv1.ActionParametersResponse)
				actionParametersResponseModel.Content = core.StringPtr(`{"success": false, "error": "you have been blocked"}`)
				actionParametersResponseModel.ContentType = core.StringPtr("application/json")
				actionParametersResponseModel.StatusCode = core.Int64Ptr(int64(400))

				// Construct an instance of the ActionParameters model
				actionParametersModel := new(rulesetsv1.ActionParameters)
				actionParametersModel.ID = core.StringPtr("testString")
				actionParametersModel.Overrides = overridesModel
				actionParametersModel.Version = core.StringPtr("testString")
				actionParametersModel.Ruleset = core.StringPtr("testString")
				actionParametersModel.Rulesets = []string{"testString"}
				actionParametersModel.Response = actionParametersResponseModel

				// Construct an instance of the Logging model
				loggingModel := new(rulesetsv1.Logging)
				loggingModel.Enabled = core.BoolPtr(true)

				// Construct an instance of the Position model
				positionModel := new(rulesetsv1.Position)
				positionModel.Before = core.StringPtr("testString")
				positionModel.After = core.StringPtr("testString")
				positionModel.Index = core.Int64Ptr(int64(0))

				// Construct an instance of the UpdateZoneRulesetRuleOptions model
				updateZoneRulesetRuleOptionsModel := new(rulesetsv1.UpdateZoneRulesetRuleOptions)
				updateZoneRulesetRuleOptionsModel.RulesetID = core.StringPtr("testString")
				updateZoneRulesetRuleOptionsModel.RuleID = core.StringPtr("testString")
				updateZoneRulesetRuleOptionsModel.Action = core.StringPtr("testString")
				updateZoneRulesetRuleOptionsModel.ActionParameters = actionParametersModel
				updateZoneRulesetRuleOptionsModel.Description = core.StringPtr("testString")
				updateZoneRulesetRuleOptionsModel.Enabled = core.BoolPtr(true)
				updateZoneRulesetRuleOptionsModel.Expression = core.StringPtr("ip.src ne 1.1.1.1")
				updateZoneRulesetRuleOptionsModel.ID = core.StringPtr("testString")
				updateZoneRulesetRuleOptionsModel.Logging = loggingModel
				updateZoneRulesetRuleOptionsModel.Ref = core.StringPtr("my_ref")
				updateZoneRulesetRuleOptionsModel.Position = positionModel
				updateZoneRulesetRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := rulesetsService.UpdateZoneRulesetRule(updateZoneRulesetRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				rulesetsService.EnableRetries(0, 0)
				result, response, operationErr = rulesetsService.UpdateZoneRulesetRule(updateZoneRulesetRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateZoneRulesetRule(updateZoneRulesetRuleOptions *UpdateZoneRulesetRuleOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		updateZoneRulesetRulePath := "/v1/testString/zones/testString/rulesets/testString/rules/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateZoneRulesetRulePath))
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
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "messages": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "result": {"id": "ID", "version": "Version", "action": "Action", "action_parameters": {"id": "ID", "overrides": {"action": "Action", "enabled": false, "sensitivity_level": "high", "rules": [{"id": "ID", "enabled": false, "action": "Action", "sensitivity_level": "high"}], "categories": [{"category": "Category", "enabled": false, "action": "Action"}]}, "version": "Version", "ruleset": "Ruleset", "rulesets": ["Rulesets"], "response": {"content": "{\"success\": false, \"error\": \"you have been blocked\"}", "content_type": "application/json", "status_code": 400}}, "expression": "ip.src ne 1.1.1.1", "ref": "my_ref", "logging": {"enabled": true}, "last_updated": "2000-01-01T00:00:00.000000Z"}}`)
				}))
			})
			It(`Invoke UpdateZoneRulesetRule successfully with retries`, func() {
				rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(rulesetsService).ToNot(BeNil())
				rulesetsService.EnableRetries(0, 0)

				// Construct an instance of the RulesOverride model
				rulesOverrideModel := new(rulesetsv1.RulesOverride)
				rulesOverrideModel.ID = core.StringPtr("testString")
				rulesOverrideModel.Enabled = core.BoolPtr(true)
				rulesOverrideModel.Action = core.StringPtr("testString")
				rulesOverrideModel.SensitivityLevel = core.StringPtr("high")

				// Construct an instance of the CategoriesOverride model
				categoriesOverrideModel := new(rulesetsv1.CategoriesOverride)
				categoriesOverrideModel.Category = core.StringPtr("testString")
				categoriesOverrideModel.Enabled = core.BoolPtr(true)
				categoriesOverrideModel.Action = core.StringPtr("testString")

				// Construct an instance of the Overrides model
				overridesModel := new(rulesetsv1.Overrides)
				overridesModel.Action = core.StringPtr("testString")
				overridesModel.Enabled = core.BoolPtr(true)
				overridesModel.SensitivityLevel = core.StringPtr("high")
				overridesModel.Rules = []rulesetsv1.RulesOverride{*rulesOverrideModel}
				overridesModel.Categories = []rulesetsv1.CategoriesOverride{*categoriesOverrideModel}

				// Construct an instance of the ActionParametersResponse model
				actionParametersResponseModel := new(rulesetsv1.ActionParametersResponse)
				actionParametersResponseModel.Content = core.StringPtr(`{"success": false, "error": "you have been blocked"}`)
				actionParametersResponseModel.ContentType = core.StringPtr("application/json")
				actionParametersResponseModel.StatusCode = core.Int64Ptr(int64(400))

				// Construct an instance of the ActionParameters model
				actionParametersModel := new(rulesetsv1.ActionParameters)
				actionParametersModel.ID = core.StringPtr("testString")
				actionParametersModel.Overrides = overridesModel
				actionParametersModel.Version = core.StringPtr("testString")
				actionParametersModel.Ruleset = core.StringPtr("testString")
				actionParametersModel.Rulesets = []string{"testString"}
				actionParametersModel.Response = actionParametersResponseModel

				// Construct an instance of the Logging model
				loggingModel := new(rulesetsv1.Logging)
				loggingModel.Enabled = core.BoolPtr(true)

				// Construct an instance of the Position model
				positionModel := new(rulesetsv1.Position)
				positionModel.Before = core.StringPtr("testString")
				positionModel.After = core.StringPtr("testString")
				positionModel.Index = core.Int64Ptr(int64(0))

				// Construct an instance of the UpdateZoneRulesetRuleOptions model
				updateZoneRulesetRuleOptionsModel := new(rulesetsv1.UpdateZoneRulesetRuleOptions)
				updateZoneRulesetRuleOptionsModel.RulesetID = core.StringPtr("testString")
				updateZoneRulesetRuleOptionsModel.RuleID = core.StringPtr("testString")
				updateZoneRulesetRuleOptionsModel.Action = core.StringPtr("testString")
				updateZoneRulesetRuleOptionsModel.ActionParameters = actionParametersModel
				updateZoneRulesetRuleOptionsModel.Description = core.StringPtr("testString")
				updateZoneRulesetRuleOptionsModel.Enabled = core.BoolPtr(true)
				updateZoneRulesetRuleOptionsModel.Expression = core.StringPtr("ip.src ne 1.1.1.1")
				updateZoneRulesetRuleOptionsModel.ID = core.StringPtr("testString")
				updateZoneRulesetRuleOptionsModel.Logging = loggingModel
				updateZoneRulesetRuleOptionsModel.Ref = core.StringPtr("my_ref")
				updateZoneRulesetRuleOptionsModel.Position = positionModel
				updateZoneRulesetRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := rulesetsService.UpdateZoneRulesetRuleWithContext(ctx, updateZoneRulesetRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				rulesetsService.DisableRetries()
				result, response, operationErr := rulesetsService.UpdateZoneRulesetRule(updateZoneRulesetRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = rulesetsService.UpdateZoneRulesetRuleWithContext(ctx, updateZoneRulesetRuleOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(updateZoneRulesetRulePath))
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
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "messages": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "result": {"id": "ID", "version": "Version", "action": "Action", "action_parameters": {"id": "ID", "overrides": {"action": "Action", "enabled": false, "sensitivity_level": "high", "rules": [{"id": "ID", "enabled": false, "action": "Action", "sensitivity_level": "high"}], "categories": [{"category": "Category", "enabled": false, "action": "Action"}]}, "version": "Version", "ruleset": "Ruleset", "rulesets": ["Rulesets"], "response": {"content": "{\"success\": false, \"error\": \"you have been blocked\"}", "content_type": "application/json", "status_code": 400}}, "expression": "ip.src ne 1.1.1.1", "ref": "my_ref", "logging": {"enabled": true}, "last_updated": "2000-01-01T00:00:00.000000Z"}}`)
				}))
			})
			It(`Invoke UpdateZoneRulesetRule successfully`, func() {
				rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(rulesetsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := rulesetsService.UpdateZoneRulesetRule(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the RulesOverride model
				rulesOverrideModel := new(rulesetsv1.RulesOverride)
				rulesOverrideModel.ID = core.StringPtr("testString")
				rulesOverrideModel.Enabled = core.BoolPtr(true)
				rulesOverrideModel.Action = core.StringPtr("testString")
				rulesOverrideModel.SensitivityLevel = core.StringPtr("high")

				// Construct an instance of the CategoriesOverride model
				categoriesOverrideModel := new(rulesetsv1.CategoriesOverride)
				categoriesOverrideModel.Category = core.StringPtr("testString")
				categoriesOverrideModel.Enabled = core.BoolPtr(true)
				categoriesOverrideModel.Action = core.StringPtr("testString")

				// Construct an instance of the Overrides model
				overridesModel := new(rulesetsv1.Overrides)
				overridesModel.Action = core.StringPtr("testString")
				overridesModel.Enabled = core.BoolPtr(true)
				overridesModel.SensitivityLevel = core.StringPtr("high")
				overridesModel.Rules = []rulesetsv1.RulesOverride{*rulesOverrideModel}
				overridesModel.Categories = []rulesetsv1.CategoriesOverride{*categoriesOverrideModel}

				// Construct an instance of the ActionParametersResponse model
				actionParametersResponseModel := new(rulesetsv1.ActionParametersResponse)
				actionParametersResponseModel.Content = core.StringPtr(`{"success": false, "error": "you have been blocked"}`)
				actionParametersResponseModel.ContentType = core.StringPtr("application/json")
				actionParametersResponseModel.StatusCode = core.Int64Ptr(int64(400))

				// Construct an instance of the ActionParameters model
				actionParametersModel := new(rulesetsv1.ActionParameters)
				actionParametersModel.ID = core.StringPtr("testString")
				actionParametersModel.Overrides = overridesModel
				actionParametersModel.Version = core.StringPtr("testString")
				actionParametersModel.Ruleset = core.StringPtr("testString")
				actionParametersModel.Rulesets = []string{"testString"}
				actionParametersModel.Response = actionParametersResponseModel

				// Construct an instance of the Logging model
				loggingModel := new(rulesetsv1.Logging)
				loggingModel.Enabled = core.BoolPtr(true)

				// Construct an instance of the Position model
				positionModel := new(rulesetsv1.Position)
				positionModel.Before = core.StringPtr("testString")
				positionModel.After = core.StringPtr("testString")
				positionModel.Index = core.Int64Ptr(int64(0))

				// Construct an instance of the UpdateZoneRulesetRuleOptions model
				updateZoneRulesetRuleOptionsModel := new(rulesetsv1.UpdateZoneRulesetRuleOptions)
				updateZoneRulesetRuleOptionsModel.RulesetID = core.StringPtr("testString")
				updateZoneRulesetRuleOptionsModel.RuleID = core.StringPtr("testString")
				updateZoneRulesetRuleOptionsModel.Action = core.StringPtr("testString")
				updateZoneRulesetRuleOptionsModel.ActionParameters = actionParametersModel
				updateZoneRulesetRuleOptionsModel.Description = core.StringPtr("testString")
				updateZoneRulesetRuleOptionsModel.Enabled = core.BoolPtr(true)
				updateZoneRulesetRuleOptionsModel.Expression = core.StringPtr("ip.src ne 1.1.1.1")
				updateZoneRulesetRuleOptionsModel.ID = core.StringPtr("testString")
				updateZoneRulesetRuleOptionsModel.Logging = loggingModel
				updateZoneRulesetRuleOptionsModel.Ref = core.StringPtr("my_ref")
				updateZoneRulesetRuleOptionsModel.Position = positionModel
				updateZoneRulesetRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = rulesetsService.UpdateZoneRulesetRule(updateZoneRulesetRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateZoneRulesetRule with error: Operation validation and request error`, func() {
				rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(rulesetsService).ToNot(BeNil())

				// Construct an instance of the RulesOverride model
				rulesOverrideModel := new(rulesetsv1.RulesOverride)
				rulesOverrideModel.ID = core.StringPtr("testString")
				rulesOverrideModel.Enabled = core.BoolPtr(true)
				rulesOverrideModel.Action = core.StringPtr("testString")
				rulesOverrideModel.SensitivityLevel = core.StringPtr("high")

				// Construct an instance of the CategoriesOverride model
				categoriesOverrideModel := new(rulesetsv1.CategoriesOverride)
				categoriesOverrideModel.Category = core.StringPtr("testString")
				categoriesOverrideModel.Enabled = core.BoolPtr(true)
				categoriesOverrideModel.Action = core.StringPtr("testString")

				// Construct an instance of the Overrides model
				overridesModel := new(rulesetsv1.Overrides)
				overridesModel.Action = core.StringPtr("testString")
				overridesModel.Enabled = core.BoolPtr(true)
				overridesModel.SensitivityLevel = core.StringPtr("high")
				overridesModel.Rules = []rulesetsv1.RulesOverride{*rulesOverrideModel}
				overridesModel.Categories = []rulesetsv1.CategoriesOverride{*categoriesOverrideModel}

				// Construct an instance of the ActionParametersResponse model
				actionParametersResponseModel := new(rulesetsv1.ActionParametersResponse)
				actionParametersResponseModel.Content = core.StringPtr(`{"success": false, "error": "you have been blocked"}`)
				actionParametersResponseModel.ContentType = core.StringPtr("application/json")
				actionParametersResponseModel.StatusCode = core.Int64Ptr(int64(400))

				// Construct an instance of the ActionParameters model
				actionParametersModel := new(rulesetsv1.ActionParameters)
				actionParametersModel.ID = core.StringPtr("testString")
				actionParametersModel.Overrides = overridesModel
				actionParametersModel.Version = core.StringPtr("testString")
				actionParametersModel.Ruleset = core.StringPtr("testString")
				actionParametersModel.Rulesets = []string{"testString"}
				actionParametersModel.Response = actionParametersResponseModel

				// Construct an instance of the Logging model
				loggingModel := new(rulesetsv1.Logging)
				loggingModel.Enabled = core.BoolPtr(true)

				// Construct an instance of the Position model
				positionModel := new(rulesetsv1.Position)
				positionModel.Before = core.StringPtr("testString")
				positionModel.After = core.StringPtr("testString")
				positionModel.Index = core.Int64Ptr(int64(0))

				// Construct an instance of the UpdateZoneRulesetRuleOptions model
				updateZoneRulesetRuleOptionsModel := new(rulesetsv1.UpdateZoneRulesetRuleOptions)
				updateZoneRulesetRuleOptionsModel.RulesetID = core.StringPtr("testString")
				updateZoneRulesetRuleOptionsModel.RuleID = core.StringPtr("testString")
				updateZoneRulesetRuleOptionsModel.Action = core.StringPtr("testString")
				updateZoneRulesetRuleOptionsModel.ActionParameters = actionParametersModel
				updateZoneRulesetRuleOptionsModel.Description = core.StringPtr("testString")
				updateZoneRulesetRuleOptionsModel.Enabled = core.BoolPtr(true)
				updateZoneRulesetRuleOptionsModel.Expression = core.StringPtr("ip.src ne 1.1.1.1")
				updateZoneRulesetRuleOptionsModel.ID = core.StringPtr("testString")
				updateZoneRulesetRuleOptionsModel.Logging = loggingModel
				updateZoneRulesetRuleOptionsModel.Ref = core.StringPtr("my_ref")
				updateZoneRulesetRuleOptionsModel.Position = positionModel
				updateZoneRulesetRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := rulesetsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := rulesetsService.UpdateZoneRulesetRule(updateZoneRulesetRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateZoneRulesetRuleOptions model with no property values
				updateZoneRulesetRuleOptionsModelNew := new(rulesetsv1.UpdateZoneRulesetRuleOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = rulesetsService.UpdateZoneRulesetRule(updateZoneRulesetRuleOptionsModelNew)
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
			It(`Invoke UpdateZoneRulesetRule successfully`, func() {
				rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(rulesetsService).ToNot(BeNil())

				// Construct an instance of the RulesOverride model
				rulesOverrideModel := new(rulesetsv1.RulesOverride)
				rulesOverrideModel.ID = core.StringPtr("testString")
				rulesOverrideModel.Enabled = core.BoolPtr(true)
				rulesOverrideModel.Action = core.StringPtr("testString")
				rulesOverrideModel.SensitivityLevel = core.StringPtr("high")

				// Construct an instance of the CategoriesOverride model
				categoriesOverrideModel := new(rulesetsv1.CategoriesOverride)
				categoriesOverrideModel.Category = core.StringPtr("testString")
				categoriesOverrideModel.Enabled = core.BoolPtr(true)
				categoriesOverrideModel.Action = core.StringPtr("testString")

				// Construct an instance of the Overrides model
				overridesModel := new(rulesetsv1.Overrides)
				overridesModel.Action = core.StringPtr("testString")
				overridesModel.Enabled = core.BoolPtr(true)
				overridesModel.SensitivityLevel = core.StringPtr("high")
				overridesModel.Rules = []rulesetsv1.RulesOverride{*rulesOverrideModel}
				overridesModel.Categories = []rulesetsv1.CategoriesOverride{*categoriesOverrideModel}

				// Construct an instance of the ActionParametersResponse model
				actionParametersResponseModel := new(rulesetsv1.ActionParametersResponse)
				actionParametersResponseModel.Content = core.StringPtr(`{"success": false, "error": "you have been blocked"}`)
				actionParametersResponseModel.ContentType = core.StringPtr("application/json")
				actionParametersResponseModel.StatusCode = core.Int64Ptr(int64(400))

				// Construct an instance of the ActionParameters model
				actionParametersModel := new(rulesetsv1.ActionParameters)
				actionParametersModel.ID = core.StringPtr("testString")
				actionParametersModel.Overrides = overridesModel
				actionParametersModel.Version = core.StringPtr("testString")
				actionParametersModel.Ruleset = core.StringPtr("testString")
				actionParametersModel.Rulesets = []string{"testString"}
				actionParametersModel.Response = actionParametersResponseModel

				// Construct an instance of the Logging model
				loggingModel := new(rulesetsv1.Logging)
				loggingModel.Enabled = core.BoolPtr(true)

				// Construct an instance of the Position model
				positionModel := new(rulesetsv1.Position)
				positionModel.Before = core.StringPtr("testString")
				positionModel.After = core.StringPtr("testString")
				positionModel.Index = core.Int64Ptr(int64(0))

				// Construct an instance of the UpdateZoneRulesetRuleOptions model
				updateZoneRulesetRuleOptionsModel := new(rulesetsv1.UpdateZoneRulesetRuleOptions)
				updateZoneRulesetRuleOptionsModel.RulesetID = core.StringPtr("testString")
				updateZoneRulesetRuleOptionsModel.RuleID = core.StringPtr("testString")
				updateZoneRulesetRuleOptionsModel.Action = core.StringPtr("testString")
				updateZoneRulesetRuleOptionsModel.ActionParameters = actionParametersModel
				updateZoneRulesetRuleOptionsModel.Description = core.StringPtr("testString")
				updateZoneRulesetRuleOptionsModel.Enabled = core.BoolPtr(true)
				updateZoneRulesetRuleOptionsModel.Expression = core.StringPtr("ip.src ne 1.1.1.1")
				updateZoneRulesetRuleOptionsModel.ID = core.StringPtr("testString")
				updateZoneRulesetRuleOptionsModel.Logging = loggingModel
				updateZoneRulesetRuleOptionsModel.Ref = core.StringPtr("my_ref")
				updateZoneRulesetRuleOptionsModel.Position = positionModel
				updateZoneRulesetRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := rulesetsService.UpdateZoneRulesetRule(updateZoneRulesetRuleOptionsModel)
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
	Describe(`DeleteZoneRulesetRule(deleteZoneRulesetRuleOptions *DeleteZoneRulesetRuleOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		deleteZoneRulesetRulePath := "/v1/testString/zones/testString/rulesets/testString/rules/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteZoneRulesetRulePath))
					Expect(req.Method).To(Equal("DELETE"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke DeleteZoneRulesetRule with error: Operation response processing error`, func() {
				rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(rulesetsService).ToNot(BeNil())

				// Construct an instance of the DeleteZoneRulesetRuleOptions model
				deleteZoneRulesetRuleOptionsModel := new(rulesetsv1.DeleteZoneRulesetRuleOptions)
				deleteZoneRulesetRuleOptionsModel.RulesetID = core.StringPtr("testString")
				deleteZoneRulesetRuleOptionsModel.RuleID = core.StringPtr("testString")
				deleteZoneRulesetRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := rulesetsService.DeleteZoneRulesetRule(deleteZoneRulesetRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				rulesetsService.EnableRetries(0, 0)
				result, response, operationErr = rulesetsService.DeleteZoneRulesetRule(deleteZoneRulesetRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteZoneRulesetRule(deleteZoneRulesetRuleOptions *DeleteZoneRulesetRuleOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		deleteZoneRulesetRulePath := "/v1/testString/zones/testString/rulesets/testString/rules/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteZoneRulesetRulePath))
					Expect(req.Method).To(Equal("DELETE"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "messages": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "result": {"id": "ID", "version": "Version", "action": "Action", "action_parameters": {"id": "ID", "overrides": {"action": "Action", "enabled": false, "sensitivity_level": "high", "rules": [{"id": "ID", "enabled": false, "action": "Action", "sensitivity_level": "high"}], "categories": [{"category": "Category", "enabled": false, "action": "Action"}]}, "version": "Version", "ruleset": "Ruleset", "rulesets": ["Rulesets"], "response": {"content": "{\"success\": false, \"error\": \"you have been blocked\"}", "content_type": "application/json", "status_code": 400}}, "expression": "ip.src ne 1.1.1.1", "ref": "my_ref", "logging": {"enabled": true}, "last_updated": "2000-01-01T00:00:00.000000Z"}}`)
				}))
			})
			It(`Invoke DeleteZoneRulesetRule successfully with retries`, func() {
				rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(rulesetsService).ToNot(BeNil())
				rulesetsService.EnableRetries(0, 0)

				// Construct an instance of the DeleteZoneRulesetRuleOptions model
				deleteZoneRulesetRuleOptionsModel := new(rulesetsv1.DeleteZoneRulesetRuleOptions)
				deleteZoneRulesetRuleOptionsModel.RulesetID = core.StringPtr("testString")
				deleteZoneRulesetRuleOptionsModel.RuleID = core.StringPtr("testString")
				deleteZoneRulesetRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := rulesetsService.DeleteZoneRulesetRuleWithContext(ctx, deleteZoneRulesetRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				rulesetsService.DisableRetries()
				result, response, operationErr := rulesetsService.DeleteZoneRulesetRule(deleteZoneRulesetRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = rulesetsService.DeleteZoneRulesetRuleWithContext(ctx, deleteZoneRulesetRuleOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(deleteZoneRulesetRulePath))
					Expect(req.Method).To(Equal("DELETE"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "messages": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "result": {"id": "ID", "version": "Version", "action": "Action", "action_parameters": {"id": "ID", "overrides": {"action": "Action", "enabled": false, "sensitivity_level": "high", "rules": [{"id": "ID", "enabled": false, "action": "Action", "sensitivity_level": "high"}], "categories": [{"category": "Category", "enabled": false, "action": "Action"}]}, "version": "Version", "ruleset": "Ruleset", "rulesets": ["Rulesets"], "response": {"content": "{\"success\": false, \"error\": \"you have been blocked\"}", "content_type": "application/json", "status_code": 400}}, "expression": "ip.src ne 1.1.1.1", "ref": "my_ref", "logging": {"enabled": true}, "last_updated": "2000-01-01T00:00:00.000000Z"}}`)
				}))
			})
			It(`Invoke DeleteZoneRulesetRule successfully`, func() {
				rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(rulesetsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := rulesetsService.DeleteZoneRulesetRule(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DeleteZoneRulesetRuleOptions model
				deleteZoneRulesetRuleOptionsModel := new(rulesetsv1.DeleteZoneRulesetRuleOptions)
				deleteZoneRulesetRuleOptionsModel.RulesetID = core.StringPtr("testString")
				deleteZoneRulesetRuleOptionsModel.RuleID = core.StringPtr("testString")
				deleteZoneRulesetRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = rulesetsService.DeleteZoneRulesetRule(deleteZoneRulesetRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke DeleteZoneRulesetRule with error: Operation validation and request error`, func() {
				rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(rulesetsService).ToNot(BeNil())

				// Construct an instance of the DeleteZoneRulesetRuleOptions model
				deleteZoneRulesetRuleOptionsModel := new(rulesetsv1.DeleteZoneRulesetRuleOptions)
				deleteZoneRulesetRuleOptionsModel.RulesetID = core.StringPtr("testString")
				deleteZoneRulesetRuleOptionsModel.RuleID = core.StringPtr("testString")
				deleteZoneRulesetRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := rulesetsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := rulesetsService.DeleteZoneRulesetRule(deleteZoneRulesetRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the DeleteZoneRulesetRuleOptions model with no property values
				deleteZoneRulesetRuleOptionsModelNew := new(rulesetsv1.DeleteZoneRulesetRuleOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = rulesetsService.DeleteZoneRulesetRule(deleteZoneRulesetRuleOptionsModelNew)
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
			It(`Invoke DeleteZoneRulesetRule successfully`, func() {
				rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(rulesetsService).ToNot(BeNil())

				// Construct an instance of the DeleteZoneRulesetRuleOptions model
				deleteZoneRulesetRuleOptionsModel := new(rulesetsv1.DeleteZoneRulesetRuleOptions)
				deleteZoneRulesetRuleOptionsModel.RulesetID = core.StringPtr("testString")
				deleteZoneRulesetRuleOptionsModel.RuleID = core.StringPtr("testString")
				deleteZoneRulesetRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := rulesetsService.DeleteZoneRulesetRule(deleteZoneRulesetRuleOptionsModel)
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
			rulesetsService, _ := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
				URL:            "http://rulesetsv1modelgenerator.com",
				Authenticator:  &core.NoAuthAuthenticator{},
				Crn:            core.StringPtr(crn),
				ZoneIdentifier: core.StringPtr(zoneIdentifier),
			})
			It(`Invoke NewActionParametersResponse successfully`, func() {
				content := `{"success": false, "error": "you have been blocked"}`
				contentType := "application/json"
				statusCode := int64(400)
				_model, err := rulesetsService.NewActionParametersResponse(content, contentType, statusCode)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewCreateAccountRulesetRuleOptions successfully`, func() {
				// Construct an instance of the RulesOverride model
				rulesOverrideModel := new(rulesetsv1.RulesOverride)
				Expect(rulesOverrideModel).ToNot(BeNil())
				rulesOverrideModel.ID = core.StringPtr("testString")
				rulesOverrideModel.Enabled = core.BoolPtr(true)
				rulesOverrideModel.Action = core.StringPtr("testString")
				rulesOverrideModel.SensitivityLevel = core.StringPtr("high")
				Expect(rulesOverrideModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(rulesOverrideModel.Enabled).To(Equal(core.BoolPtr(true)))
				Expect(rulesOverrideModel.Action).To(Equal(core.StringPtr("testString")))
				Expect(rulesOverrideModel.SensitivityLevel).To(Equal(core.StringPtr("high")))

				// Construct an instance of the CategoriesOverride model
				categoriesOverrideModel := new(rulesetsv1.CategoriesOverride)
				Expect(categoriesOverrideModel).ToNot(BeNil())
				categoriesOverrideModel.Category = core.StringPtr("testString")
				categoriesOverrideModel.Enabled = core.BoolPtr(true)
				categoriesOverrideModel.Action = core.StringPtr("testString")
				Expect(categoriesOverrideModel.Category).To(Equal(core.StringPtr("testString")))
				Expect(categoriesOverrideModel.Enabled).To(Equal(core.BoolPtr(true)))
				Expect(categoriesOverrideModel.Action).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the Overrides model
				overridesModel := new(rulesetsv1.Overrides)
				Expect(overridesModel).ToNot(BeNil())
				overridesModel.Action = core.StringPtr("testString")
				overridesModel.Enabled = core.BoolPtr(true)
				overridesModel.SensitivityLevel = core.StringPtr("high")
				overridesModel.Rules = []rulesetsv1.RulesOverride{*rulesOverrideModel}
				overridesModel.Categories = []rulesetsv1.CategoriesOverride{*categoriesOverrideModel}
				Expect(overridesModel.Action).To(Equal(core.StringPtr("testString")))
				Expect(overridesModel.Enabled).To(Equal(core.BoolPtr(true)))
				Expect(overridesModel.SensitivityLevel).To(Equal(core.StringPtr("high")))
				Expect(overridesModel.Rules).To(Equal([]rulesetsv1.RulesOverride{*rulesOverrideModel}))
				Expect(overridesModel.Categories).To(Equal([]rulesetsv1.CategoriesOverride{*categoriesOverrideModel}))

				// Construct an instance of the ActionParametersResponse model
				actionParametersResponseModel := new(rulesetsv1.ActionParametersResponse)
				Expect(actionParametersResponseModel).ToNot(BeNil())
				actionParametersResponseModel.Content = core.StringPtr(`{"success": false, "error": "you have been blocked"}`)
				actionParametersResponseModel.ContentType = core.StringPtr("application/json")
				actionParametersResponseModel.StatusCode = core.Int64Ptr(int64(400))
				Expect(actionParametersResponseModel.Content).To(Equal(core.StringPtr(`{"success": false, "error": "you have been blocked"}`)))
				Expect(actionParametersResponseModel.ContentType).To(Equal(core.StringPtr("application/json")))
				Expect(actionParametersResponseModel.StatusCode).To(Equal(core.Int64Ptr(int64(400))))

				// Construct an instance of the ActionParameters model
				actionParametersModel := new(rulesetsv1.ActionParameters)
				Expect(actionParametersModel).ToNot(BeNil())
				actionParametersModel.ID = core.StringPtr("testString")
				actionParametersModel.Overrides = overridesModel
				actionParametersModel.Version = core.StringPtr("testString")
				actionParametersModel.Ruleset = core.StringPtr("testString")
				actionParametersModel.Rulesets = []string{"testString"}
				actionParametersModel.Response = actionParametersResponseModel
				Expect(actionParametersModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(actionParametersModel.Overrides).To(Equal(overridesModel))
				Expect(actionParametersModel.Version).To(Equal(core.StringPtr("testString")))
				Expect(actionParametersModel.Ruleset).To(Equal(core.StringPtr("testString")))
				Expect(actionParametersModel.Rulesets).To(Equal([]string{"testString"}))
				Expect(actionParametersModel.Response).To(Equal(actionParametersResponseModel))

				// Construct an instance of the Logging model
				loggingModel := new(rulesetsv1.Logging)
				Expect(loggingModel).ToNot(BeNil())
				loggingModel.Enabled = core.BoolPtr(true)
				Expect(loggingModel.Enabled).To(Equal(core.BoolPtr(true)))

				// Construct an instance of the Position model
				positionModel := new(rulesetsv1.Position)
				Expect(positionModel).ToNot(BeNil())
				positionModel.Before = core.StringPtr("testString")
				positionModel.After = core.StringPtr("testString")
				positionModel.Index = core.Int64Ptr(int64(0))
				Expect(positionModel.Before).To(Equal(core.StringPtr("testString")))
				Expect(positionModel.After).To(Equal(core.StringPtr("testString")))
				Expect(positionModel.Index).To(Equal(core.Int64Ptr(int64(0))))

				// Construct an instance of the CreateAccountRulesetRuleOptions model
				rulesetID := "testString"
				createAccountRulesetRuleOptionsModel := rulesetsService.NewCreateAccountRulesetRuleOptions(rulesetID)
				createAccountRulesetRuleOptionsModel.SetRulesetID("testString")
				createAccountRulesetRuleOptionsModel.SetAction("testString")
				createAccountRulesetRuleOptionsModel.SetActionParameters(actionParametersModel)
				createAccountRulesetRuleOptionsModel.SetDescription("testString")
				createAccountRulesetRuleOptionsModel.SetEnabled(true)
				createAccountRulesetRuleOptionsModel.SetExpression("ip.src ne 1.1.1.1")
				createAccountRulesetRuleOptionsModel.SetID("testString")
				createAccountRulesetRuleOptionsModel.SetLogging(loggingModel)
				createAccountRulesetRuleOptionsModel.SetRef("my_ref")
				createAccountRulesetRuleOptionsModel.SetPosition(positionModel)
				createAccountRulesetRuleOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createAccountRulesetRuleOptionsModel).ToNot(BeNil())
				Expect(createAccountRulesetRuleOptionsModel.RulesetID).To(Equal(core.StringPtr("testString")))
				Expect(createAccountRulesetRuleOptionsModel.Action).To(Equal(core.StringPtr("testString")))
				Expect(createAccountRulesetRuleOptionsModel.ActionParameters).To(Equal(actionParametersModel))
				Expect(createAccountRulesetRuleOptionsModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(createAccountRulesetRuleOptionsModel.Enabled).To(Equal(core.BoolPtr(true)))
				Expect(createAccountRulesetRuleOptionsModel.Expression).To(Equal(core.StringPtr("ip.src ne 1.1.1.1")))
				Expect(createAccountRulesetRuleOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(createAccountRulesetRuleOptionsModel.Logging).To(Equal(loggingModel))
				Expect(createAccountRulesetRuleOptionsModel.Ref).To(Equal(core.StringPtr("my_ref")))
				Expect(createAccountRulesetRuleOptionsModel.Position).To(Equal(positionModel))
				Expect(createAccountRulesetRuleOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateZoneRulesetRuleOptions successfully`, func() {
				// Construct an instance of the RulesOverride model
				rulesOverrideModel := new(rulesetsv1.RulesOverride)
				Expect(rulesOverrideModel).ToNot(BeNil())
				rulesOverrideModel.ID = core.StringPtr("testString")
				rulesOverrideModel.Enabled = core.BoolPtr(true)
				rulesOverrideModel.Action = core.StringPtr("testString")
				rulesOverrideModel.SensitivityLevel = core.StringPtr("high")
				Expect(rulesOverrideModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(rulesOverrideModel.Enabled).To(Equal(core.BoolPtr(true)))
				Expect(rulesOverrideModel.Action).To(Equal(core.StringPtr("testString")))
				Expect(rulesOverrideModel.SensitivityLevel).To(Equal(core.StringPtr("high")))

				// Construct an instance of the CategoriesOverride model
				categoriesOverrideModel := new(rulesetsv1.CategoriesOverride)
				Expect(categoriesOverrideModel).ToNot(BeNil())
				categoriesOverrideModel.Category = core.StringPtr("testString")
				categoriesOverrideModel.Enabled = core.BoolPtr(true)
				categoriesOverrideModel.Action = core.StringPtr("testString")
				Expect(categoriesOverrideModel.Category).To(Equal(core.StringPtr("testString")))
				Expect(categoriesOverrideModel.Enabled).To(Equal(core.BoolPtr(true)))
				Expect(categoriesOverrideModel.Action).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the Overrides model
				overridesModel := new(rulesetsv1.Overrides)
				Expect(overridesModel).ToNot(BeNil())
				overridesModel.Action = core.StringPtr("testString")
				overridesModel.Enabled = core.BoolPtr(true)
				overridesModel.SensitivityLevel = core.StringPtr("high")
				overridesModel.Rules = []rulesetsv1.RulesOverride{*rulesOverrideModel}
				overridesModel.Categories = []rulesetsv1.CategoriesOverride{*categoriesOverrideModel}
				Expect(overridesModel.Action).To(Equal(core.StringPtr("testString")))
				Expect(overridesModel.Enabled).To(Equal(core.BoolPtr(true)))
				Expect(overridesModel.SensitivityLevel).To(Equal(core.StringPtr("high")))
				Expect(overridesModel.Rules).To(Equal([]rulesetsv1.RulesOverride{*rulesOverrideModel}))
				Expect(overridesModel.Categories).To(Equal([]rulesetsv1.CategoriesOverride{*categoriesOverrideModel}))

				// Construct an instance of the ActionParametersResponse model
				actionParametersResponseModel := new(rulesetsv1.ActionParametersResponse)
				Expect(actionParametersResponseModel).ToNot(BeNil())
				actionParametersResponseModel.Content = core.StringPtr(`{"success": false, "error": "you have been blocked"}`)
				actionParametersResponseModel.ContentType = core.StringPtr("application/json")
				actionParametersResponseModel.StatusCode = core.Int64Ptr(int64(400))
				Expect(actionParametersResponseModel.Content).To(Equal(core.StringPtr(`{"success": false, "error": "you have been blocked"}`)))
				Expect(actionParametersResponseModel.ContentType).To(Equal(core.StringPtr("application/json")))
				Expect(actionParametersResponseModel.StatusCode).To(Equal(core.Int64Ptr(int64(400))))

				// Construct an instance of the ActionParameters model
				actionParametersModel := new(rulesetsv1.ActionParameters)
				Expect(actionParametersModel).ToNot(BeNil())
				actionParametersModel.ID = core.StringPtr("testString")
				actionParametersModel.Overrides = overridesModel
				actionParametersModel.Version = core.StringPtr("testString")
				actionParametersModel.Ruleset = core.StringPtr("testString")
				actionParametersModel.Rulesets = []string{"testString"}
				actionParametersModel.Response = actionParametersResponseModel
				Expect(actionParametersModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(actionParametersModel.Overrides).To(Equal(overridesModel))
				Expect(actionParametersModel.Version).To(Equal(core.StringPtr("testString")))
				Expect(actionParametersModel.Ruleset).To(Equal(core.StringPtr("testString")))
				Expect(actionParametersModel.Rulesets).To(Equal([]string{"testString"}))
				Expect(actionParametersModel.Response).To(Equal(actionParametersResponseModel))

				// Construct an instance of the Logging model
				loggingModel := new(rulesetsv1.Logging)
				Expect(loggingModel).ToNot(BeNil())
				loggingModel.Enabled = core.BoolPtr(true)
				Expect(loggingModel.Enabled).To(Equal(core.BoolPtr(true)))

				// Construct an instance of the Position model
				positionModel := new(rulesetsv1.Position)
				Expect(positionModel).ToNot(BeNil())
				positionModel.Before = core.StringPtr("testString")
				positionModel.After = core.StringPtr("testString")
				positionModel.Index = core.Int64Ptr(int64(0))
				Expect(positionModel.Before).To(Equal(core.StringPtr("testString")))
				Expect(positionModel.After).To(Equal(core.StringPtr("testString")))
				Expect(positionModel.Index).To(Equal(core.Int64Ptr(int64(0))))

				// Construct an instance of the CreateZoneRulesetRuleOptions model
				rulesetID := "testString"
				createZoneRulesetRuleOptionsModel := rulesetsService.NewCreateZoneRulesetRuleOptions(rulesetID)
				createZoneRulesetRuleOptionsModel.SetRulesetID("testString")
				createZoneRulesetRuleOptionsModel.SetAction("testString")
				createZoneRulesetRuleOptionsModel.SetActionParameters(actionParametersModel)
				createZoneRulesetRuleOptionsModel.SetDescription("testString")
				createZoneRulesetRuleOptionsModel.SetEnabled(true)
				createZoneRulesetRuleOptionsModel.SetExpression("ip.src ne 1.1.1.1")
				createZoneRulesetRuleOptionsModel.SetID("testString")
				createZoneRulesetRuleOptionsModel.SetLogging(loggingModel)
				createZoneRulesetRuleOptionsModel.SetRef("my_ref")
				createZoneRulesetRuleOptionsModel.SetPosition(positionModel)
				createZoneRulesetRuleOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createZoneRulesetRuleOptionsModel).ToNot(BeNil())
				Expect(createZoneRulesetRuleOptionsModel.RulesetID).To(Equal(core.StringPtr("testString")))
				Expect(createZoneRulesetRuleOptionsModel.Action).To(Equal(core.StringPtr("testString")))
				Expect(createZoneRulesetRuleOptionsModel.ActionParameters).To(Equal(actionParametersModel))
				Expect(createZoneRulesetRuleOptionsModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(createZoneRulesetRuleOptionsModel.Enabled).To(Equal(core.BoolPtr(true)))
				Expect(createZoneRulesetRuleOptionsModel.Expression).To(Equal(core.StringPtr("ip.src ne 1.1.1.1")))
				Expect(createZoneRulesetRuleOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(createZoneRulesetRuleOptionsModel.Logging).To(Equal(loggingModel))
				Expect(createZoneRulesetRuleOptionsModel.Ref).To(Equal(core.StringPtr("my_ref")))
				Expect(createZoneRulesetRuleOptionsModel.Position).To(Equal(positionModel))
				Expect(createZoneRulesetRuleOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteAccountRulesetOptions successfully`, func() {
				// Construct an instance of the DeleteAccountRulesetOptions model
				rulesetID := "testString"
				deleteAccountRulesetOptionsModel := rulesetsService.NewDeleteAccountRulesetOptions(rulesetID)
				deleteAccountRulesetOptionsModel.SetRulesetID("testString")
				deleteAccountRulesetOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteAccountRulesetOptionsModel).ToNot(BeNil())
				Expect(deleteAccountRulesetOptionsModel.RulesetID).To(Equal(core.StringPtr("testString")))
				Expect(deleteAccountRulesetOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteAccountRulesetRuleOptions successfully`, func() {
				// Construct an instance of the DeleteAccountRulesetRuleOptions model
				rulesetID := "testString"
				ruleID := "testString"
				deleteAccountRulesetRuleOptionsModel := rulesetsService.NewDeleteAccountRulesetRuleOptions(rulesetID, ruleID)
				deleteAccountRulesetRuleOptionsModel.SetRulesetID("testString")
				deleteAccountRulesetRuleOptionsModel.SetRuleID("testString")
				deleteAccountRulesetRuleOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteAccountRulesetRuleOptionsModel).ToNot(BeNil())
				Expect(deleteAccountRulesetRuleOptionsModel.RulesetID).To(Equal(core.StringPtr("testString")))
				Expect(deleteAccountRulesetRuleOptionsModel.RuleID).To(Equal(core.StringPtr("testString")))
				Expect(deleteAccountRulesetRuleOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteAccountRulesetVersionOptions successfully`, func() {
				// Construct an instance of the DeleteAccountRulesetVersionOptions model
				rulesetID := "testString"
				rulesetVersion := "1"
				deleteAccountRulesetVersionOptionsModel := rulesetsService.NewDeleteAccountRulesetVersionOptions(rulesetID, rulesetVersion)
				deleteAccountRulesetVersionOptionsModel.SetRulesetID("testString")
				deleteAccountRulesetVersionOptionsModel.SetRulesetVersion("1")
				deleteAccountRulesetVersionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteAccountRulesetVersionOptionsModel).ToNot(BeNil())
				Expect(deleteAccountRulesetVersionOptionsModel.RulesetID).To(Equal(core.StringPtr("testString")))
				Expect(deleteAccountRulesetVersionOptionsModel.RulesetVersion).To(Equal(core.StringPtr("1")))
				Expect(deleteAccountRulesetVersionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteZoneRulesetOptions successfully`, func() {
				// Construct an instance of the DeleteZoneRulesetOptions model
				rulesetID := "testString"
				deleteZoneRulesetOptionsModel := rulesetsService.NewDeleteZoneRulesetOptions(rulesetID)
				deleteZoneRulesetOptionsModel.SetRulesetID("testString")
				deleteZoneRulesetOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteZoneRulesetOptionsModel).ToNot(BeNil())
				Expect(deleteZoneRulesetOptionsModel.RulesetID).To(Equal(core.StringPtr("testString")))
				Expect(deleteZoneRulesetOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteZoneRulesetRuleOptions successfully`, func() {
				// Construct an instance of the DeleteZoneRulesetRuleOptions model
				rulesetID := "testString"
				ruleID := "testString"
				deleteZoneRulesetRuleOptionsModel := rulesetsService.NewDeleteZoneRulesetRuleOptions(rulesetID, ruleID)
				deleteZoneRulesetRuleOptionsModel.SetRulesetID("testString")
				deleteZoneRulesetRuleOptionsModel.SetRuleID("testString")
				deleteZoneRulesetRuleOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteZoneRulesetRuleOptionsModel).ToNot(BeNil())
				Expect(deleteZoneRulesetRuleOptionsModel.RulesetID).To(Equal(core.StringPtr("testString")))
				Expect(deleteZoneRulesetRuleOptionsModel.RuleID).To(Equal(core.StringPtr("testString")))
				Expect(deleteZoneRulesetRuleOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteZoneRulesetVersionOptions successfully`, func() {
				// Construct an instance of the DeleteZoneRulesetVersionOptions model
				rulesetID := "testString"
				rulesetVersion := "1"
				deleteZoneRulesetVersionOptionsModel := rulesetsService.NewDeleteZoneRulesetVersionOptions(rulesetID, rulesetVersion)
				deleteZoneRulesetVersionOptionsModel.SetRulesetID("testString")
				deleteZoneRulesetVersionOptionsModel.SetRulesetVersion("1")
				deleteZoneRulesetVersionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteZoneRulesetVersionOptionsModel).ToNot(BeNil())
				Expect(deleteZoneRulesetVersionOptionsModel.RulesetID).To(Equal(core.StringPtr("testString")))
				Expect(deleteZoneRulesetVersionOptionsModel.RulesetVersion).To(Equal(core.StringPtr("1")))
				Expect(deleteZoneRulesetVersionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetAccountEntryPointRulesetVersionOptions successfully`, func() {
				// Construct an instance of the GetAccountEntryPointRulesetVersionOptions model
				rulesetPhase := "ddos_l4"
				rulesetVersion := "1"
				getAccountEntryPointRulesetVersionOptionsModel := rulesetsService.NewGetAccountEntryPointRulesetVersionOptions(rulesetPhase, rulesetVersion)
				getAccountEntryPointRulesetVersionOptionsModel.SetRulesetPhase("ddos_l4")
				getAccountEntryPointRulesetVersionOptionsModel.SetRulesetVersion("1")
				getAccountEntryPointRulesetVersionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getAccountEntryPointRulesetVersionOptionsModel).ToNot(BeNil())
				Expect(getAccountEntryPointRulesetVersionOptionsModel.RulesetPhase).To(Equal(core.StringPtr("ddos_l4")))
				Expect(getAccountEntryPointRulesetVersionOptionsModel.RulesetVersion).To(Equal(core.StringPtr("1")))
				Expect(getAccountEntryPointRulesetVersionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetAccountEntryPointRulesetVersionsOptions successfully`, func() {
				// Construct an instance of the GetAccountEntryPointRulesetVersionsOptions model
				rulesetPhase := "ddos_l4"
				getAccountEntryPointRulesetVersionsOptionsModel := rulesetsService.NewGetAccountEntryPointRulesetVersionsOptions(rulesetPhase)
				getAccountEntryPointRulesetVersionsOptionsModel.SetRulesetPhase("ddos_l4")
				getAccountEntryPointRulesetVersionsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getAccountEntryPointRulesetVersionsOptionsModel).ToNot(BeNil())
				Expect(getAccountEntryPointRulesetVersionsOptionsModel.RulesetPhase).To(Equal(core.StringPtr("ddos_l4")))
				Expect(getAccountEntryPointRulesetVersionsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetAccountEntrypointRulesetOptions successfully`, func() {
				// Construct an instance of the GetAccountEntrypointRulesetOptions model
				rulesetPhase := "ddos_l4"
				getAccountEntrypointRulesetOptionsModel := rulesetsService.NewGetAccountEntrypointRulesetOptions(rulesetPhase)
				getAccountEntrypointRulesetOptionsModel.SetRulesetPhase("ddos_l4")
				getAccountEntrypointRulesetOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getAccountEntrypointRulesetOptionsModel).ToNot(BeNil())
				Expect(getAccountEntrypointRulesetOptionsModel.RulesetPhase).To(Equal(core.StringPtr("ddos_l4")))
				Expect(getAccountEntrypointRulesetOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetAccountRulesetOptions successfully`, func() {
				// Construct an instance of the GetAccountRulesetOptions model
				rulesetID := "testString"
				getAccountRulesetOptionsModel := rulesetsService.NewGetAccountRulesetOptions(rulesetID)
				getAccountRulesetOptionsModel.SetRulesetID("testString")
				getAccountRulesetOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getAccountRulesetOptionsModel).ToNot(BeNil())
				Expect(getAccountRulesetOptionsModel.RulesetID).To(Equal(core.StringPtr("testString")))
				Expect(getAccountRulesetOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetAccountRulesetVersionByTagOptions successfully`, func() {
				// Construct an instance of the GetAccountRulesetVersionByTagOptions model
				rulesetID := "testString"
				rulesetVersion := "1"
				ruleTag := "testString"
				getAccountRulesetVersionByTagOptionsModel := rulesetsService.NewGetAccountRulesetVersionByTagOptions(rulesetID, rulesetVersion, ruleTag)
				getAccountRulesetVersionByTagOptionsModel.SetRulesetID("testString")
				getAccountRulesetVersionByTagOptionsModel.SetRulesetVersion("1")
				getAccountRulesetVersionByTagOptionsModel.SetRuleTag("testString")
				getAccountRulesetVersionByTagOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getAccountRulesetVersionByTagOptionsModel).ToNot(BeNil())
				Expect(getAccountRulesetVersionByTagOptionsModel.RulesetID).To(Equal(core.StringPtr("testString")))
				Expect(getAccountRulesetVersionByTagOptionsModel.RulesetVersion).To(Equal(core.StringPtr("1")))
				Expect(getAccountRulesetVersionByTagOptionsModel.RuleTag).To(Equal(core.StringPtr("testString")))
				Expect(getAccountRulesetVersionByTagOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetAccountRulesetVersionOptions successfully`, func() {
				// Construct an instance of the GetAccountRulesetVersionOptions model
				rulesetID := "testString"
				rulesetVersion := "1"
				getAccountRulesetVersionOptionsModel := rulesetsService.NewGetAccountRulesetVersionOptions(rulesetID, rulesetVersion)
				getAccountRulesetVersionOptionsModel.SetRulesetID("testString")
				getAccountRulesetVersionOptionsModel.SetRulesetVersion("1")
				getAccountRulesetVersionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getAccountRulesetVersionOptionsModel).ToNot(BeNil())
				Expect(getAccountRulesetVersionOptionsModel.RulesetID).To(Equal(core.StringPtr("testString")))
				Expect(getAccountRulesetVersionOptionsModel.RulesetVersion).To(Equal(core.StringPtr("1")))
				Expect(getAccountRulesetVersionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetAccountRulesetVersionsOptions successfully`, func() {
				// Construct an instance of the GetAccountRulesetVersionsOptions model
				rulesetID := "testString"
				getAccountRulesetVersionsOptionsModel := rulesetsService.NewGetAccountRulesetVersionsOptions(rulesetID)
				getAccountRulesetVersionsOptionsModel.SetRulesetID("testString")
				getAccountRulesetVersionsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getAccountRulesetVersionsOptionsModel).ToNot(BeNil())
				Expect(getAccountRulesetVersionsOptionsModel.RulesetID).To(Equal(core.StringPtr("testString")))
				Expect(getAccountRulesetVersionsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetAccountRulesetsOptions successfully`, func() {
				// Construct an instance of the GetAccountRulesetsOptions model
				getAccountRulesetsOptionsModel := rulesetsService.NewGetAccountRulesetsOptions()
				getAccountRulesetsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getAccountRulesetsOptionsModel).ToNot(BeNil())
				Expect(getAccountRulesetsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetZoneEntryPointRulesetVersionOptions successfully`, func() {
				// Construct an instance of the GetZoneEntryPointRulesetVersionOptions model
				rulesetPhase := "ddos_l4"
				rulesetVersion := "1"
				getZoneEntryPointRulesetVersionOptionsModel := rulesetsService.NewGetZoneEntryPointRulesetVersionOptions(rulesetPhase, rulesetVersion)
				getZoneEntryPointRulesetVersionOptionsModel.SetRulesetPhase("ddos_l4")
				getZoneEntryPointRulesetVersionOptionsModel.SetRulesetVersion("1")
				getZoneEntryPointRulesetVersionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getZoneEntryPointRulesetVersionOptionsModel).ToNot(BeNil())
				Expect(getZoneEntryPointRulesetVersionOptionsModel.RulesetPhase).To(Equal(core.StringPtr("ddos_l4")))
				Expect(getZoneEntryPointRulesetVersionOptionsModel.RulesetVersion).To(Equal(core.StringPtr("1")))
				Expect(getZoneEntryPointRulesetVersionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetZoneEntryPointRulesetVersionsOptions successfully`, func() {
				// Construct an instance of the GetZoneEntryPointRulesetVersionsOptions model
				rulesetPhase := "ddos_l4"
				getZoneEntryPointRulesetVersionsOptionsModel := rulesetsService.NewGetZoneEntryPointRulesetVersionsOptions(rulesetPhase)
				getZoneEntryPointRulesetVersionsOptionsModel.SetRulesetPhase("ddos_l4")
				getZoneEntryPointRulesetVersionsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getZoneEntryPointRulesetVersionsOptionsModel).ToNot(BeNil())
				Expect(getZoneEntryPointRulesetVersionsOptionsModel.RulesetPhase).To(Equal(core.StringPtr("ddos_l4")))
				Expect(getZoneEntryPointRulesetVersionsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetZoneEntrypointRulesetOptions successfully`, func() {
				// Construct an instance of the GetZoneEntrypointRulesetOptions model
				rulesetPhase := "ddos_l4"
				getZoneEntrypointRulesetOptionsModel := rulesetsService.NewGetZoneEntrypointRulesetOptions(rulesetPhase)
				getZoneEntrypointRulesetOptionsModel.SetRulesetPhase("ddos_l4")
				getZoneEntrypointRulesetOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getZoneEntrypointRulesetOptionsModel).ToNot(BeNil())
				Expect(getZoneEntrypointRulesetOptionsModel.RulesetPhase).To(Equal(core.StringPtr("ddos_l4")))
				Expect(getZoneEntrypointRulesetOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetZoneRulesetOptions successfully`, func() {
				// Construct an instance of the GetZoneRulesetOptions model
				rulesetID := "testString"
				getZoneRulesetOptionsModel := rulesetsService.NewGetZoneRulesetOptions(rulesetID)
				getZoneRulesetOptionsModel.SetRulesetID("testString")
				getZoneRulesetOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getZoneRulesetOptionsModel).ToNot(BeNil())
				Expect(getZoneRulesetOptionsModel.RulesetID).To(Equal(core.StringPtr("testString")))
				Expect(getZoneRulesetOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetZoneRulesetVersionOptions successfully`, func() {
				// Construct an instance of the GetZoneRulesetVersionOptions model
				rulesetID := "testString"
				rulesetVersion := "1"
				getZoneRulesetVersionOptionsModel := rulesetsService.NewGetZoneRulesetVersionOptions(rulesetID, rulesetVersion)
				getZoneRulesetVersionOptionsModel.SetRulesetID("testString")
				getZoneRulesetVersionOptionsModel.SetRulesetVersion("1")
				getZoneRulesetVersionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getZoneRulesetVersionOptionsModel).ToNot(BeNil())
				Expect(getZoneRulesetVersionOptionsModel.RulesetID).To(Equal(core.StringPtr("testString")))
				Expect(getZoneRulesetVersionOptionsModel.RulesetVersion).To(Equal(core.StringPtr("1")))
				Expect(getZoneRulesetVersionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetZoneRulesetVersionsOptions successfully`, func() {
				// Construct an instance of the GetZoneRulesetVersionsOptions model
				rulesetID := "testString"
				getZoneRulesetVersionsOptionsModel := rulesetsService.NewGetZoneRulesetVersionsOptions(rulesetID)
				getZoneRulesetVersionsOptionsModel.SetRulesetID("testString")
				getZoneRulesetVersionsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getZoneRulesetVersionsOptionsModel).ToNot(BeNil())
				Expect(getZoneRulesetVersionsOptionsModel.RulesetID).To(Equal(core.StringPtr("testString")))
				Expect(getZoneRulesetVersionsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetZoneRulesetsOptions successfully`, func() {
				// Construct an instance of the GetZoneRulesetsOptions model
				getZoneRulesetsOptionsModel := rulesetsService.NewGetZoneRulesetsOptions()
				getZoneRulesetsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getZoneRulesetsOptionsModel).ToNot(BeNil())
				Expect(getZoneRulesetsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateAccountEntrypointRulesetOptions successfully`, func() {
				// Construct an instance of the RulesOverride model
				rulesOverrideModel := new(rulesetsv1.RulesOverride)
				Expect(rulesOverrideModel).ToNot(BeNil())
				rulesOverrideModel.ID = core.StringPtr("testString")
				rulesOverrideModel.Enabled = core.BoolPtr(true)
				rulesOverrideModel.Action = core.StringPtr("testString")
				rulesOverrideModel.SensitivityLevel = core.StringPtr("high")
				Expect(rulesOverrideModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(rulesOverrideModel.Enabled).To(Equal(core.BoolPtr(true)))
				Expect(rulesOverrideModel.Action).To(Equal(core.StringPtr("testString")))
				Expect(rulesOverrideModel.SensitivityLevel).To(Equal(core.StringPtr("high")))

				// Construct an instance of the CategoriesOverride model
				categoriesOverrideModel := new(rulesetsv1.CategoriesOverride)
				Expect(categoriesOverrideModel).ToNot(BeNil())
				categoriesOverrideModel.Category = core.StringPtr("testString")
				categoriesOverrideModel.Enabled = core.BoolPtr(true)
				categoriesOverrideModel.Action = core.StringPtr("testString")
				Expect(categoriesOverrideModel.Category).To(Equal(core.StringPtr("testString")))
				Expect(categoriesOverrideModel.Enabled).To(Equal(core.BoolPtr(true)))
				Expect(categoriesOverrideModel.Action).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the Overrides model
				overridesModel := new(rulesetsv1.Overrides)
				Expect(overridesModel).ToNot(BeNil())
				overridesModel.Action = core.StringPtr("testString")
				overridesModel.Enabled = core.BoolPtr(true)
				overridesModel.SensitivityLevel = core.StringPtr("high")
				overridesModel.Rules = []rulesetsv1.RulesOverride{*rulesOverrideModel}
				overridesModel.Categories = []rulesetsv1.CategoriesOverride{*categoriesOverrideModel}
				Expect(overridesModel.Action).To(Equal(core.StringPtr("testString")))
				Expect(overridesModel.Enabled).To(Equal(core.BoolPtr(true)))
				Expect(overridesModel.SensitivityLevel).To(Equal(core.StringPtr("high")))
				Expect(overridesModel.Rules).To(Equal([]rulesetsv1.RulesOverride{*rulesOverrideModel}))
				Expect(overridesModel.Categories).To(Equal([]rulesetsv1.CategoriesOverride{*categoriesOverrideModel}))

				// Construct an instance of the ActionParametersResponse model
				actionParametersResponseModel := new(rulesetsv1.ActionParametersResponse)
				Expect(actionParametersResponseModel).ToNot(BeNil())
				actionParametersResponseModel.Content = core.StringPtr(`{"success": false, "error": "you have been blocked"}`)
				actionParametersResponseModel.ContentType = core.StringPtr("application/json")
				actionParametersResponseModel.StatusCode = core.Int64Ptr(int64(400))
				Expect(actionParametersResponseModel.Content).To(Equal(core.StringPtr(`{"success": false, "error": "you have been blocked"}`)))
				Expect(actionParametersResponseModel.ContentType).To(Equal(core.StringPtr("application/json")))
				Expect(actionParametersResponseModel.StatusCode).To(Equal(core.Int64Ptr(int64(400))))

				// Construct an instance of the ActionParameters model
				actionParametersModel := new(rulesetsv1.ActionParameters)
				Expect(actionParametersModel).ToNot(BeNil())
				actionParametersModel.ID = core.StringPtr("testString")
				actionParametersModel.Overrides = overridesModel
				actionParametersModel.Version = core.StringPtr("testString")
				actionParametersModel.Ruleset = core.StringPtr("testString")
				actionParametersModel.Rulesets = []string{"testString"}
				actionParametersModel.Response = actionParametersResponseModel
				Expect(actionParametersModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(actionParametersModel.Overrides).To(Equal(overridesModel))
				Expect(actionParametersModel.Version).To(Equal(core.StringPtr("testString")))
				Expect(actionParametersModel.Ruleset).To(Equal(core.StringPtr("testString")))
				Expect(actionParametersModel.Rulesets).To(Equal([]string{"testString"}))
				Expect(actionParametersModel.Response).To(Equal(actionParametersResponseModel))

				// Construct an instance of the Logging model
				loggingModel := new(rulesetsv1.Logging)
				Expect(loggingModel).ToNot(BeNil())
				loggingModel.Enabled = core.BoolPtr(true)
				Expect(loggingModel.Enabled).To(Equal(core.BoolPtr(true)))

				// Construct an instance of the Position model
				positionModel := new(rulesetsv1.Position)
				Expect(positionModel).ToNot(BeNil())
				positionModel.Before = core.StringPtr("testString")
				positionModel.After = core.StringPtr("testString")
				positionModel.Index = core.Int64Ptr(int64(0))
				Expect(positionModel.Before).To(Equal(core.StringPtr("testString")))
				Expect(positionModel.After).To(Equal(core.StringPtr("testString")))
				Expect(positionModel.Index).To(Equal(core.Int64Ptr(int64(0))))

				// Construct an instance of the RuleCreate model
				ruleCreateModel := new(rulesetsv1.RuleCreate)
				Expect(ruleCreateModel).ToNot(BeNil())
				ruleCreateModel.Action = core.StringPtr("testString")
				ruleCreateModel.ActionParameters = actionParametersModel
				ruleCreateModel.Description = core.StringPtr("testString")
				ruleCreateModel.Enabled = core.BoolPtr(true)
				ruleCreateModel.Expression = core.StringPtr("ip.src ne 1.1.1.1")
				ruleCreateModel.ID = core.StringPtr("testString")
				ruleCreateModel.Logging = loggingModel
				ruleCreateModel.Ref = core.StringPtr("my_ref")
				ruleCreateModel.Position = positionModel
				Expect(ruleCreateModel.Action).To(Equal(core.StringPtr("testString")))
				Expect(ruleCreateModel.ActionParameters).To(Equal(actionParametersModel))
				Expect(ruleCreateModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(ruleCreateModel.Enabled).To(Equal(core.BoolPtr(true)))
				Expect(ruleCreateModel.Expression).To(Equal(core.StringPtr("ip.src ne 1.1.1.1")))
				Expect(ruleCreateModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(ruleCreateModel.Logging).To(Equal(loggingModel))
				Expect(ruleCreateModel.Ref).To(Equal(core.StringPtr("my_ref")))
				Expect(ruleCreateModel.Position).To(Equal(positionModel))

				// Construct an instance of the UpdateAccountEntrypointRulesetOptions model
				rulesetPhase := "ddos_l4"
				updateAccountEntrypointRulesetOptionsModel := rulesetsService.NewUpdateAccountEntrypointRulesetOptions(rulesetPhase)
				updateAccountEntrypointRulesetOptionsModel.SetRulesetPhase("ddos_l4")
				updateAccountEntrypointRulesetOptionsModel.SetDescription("Custom account ruleset")
				updateAccountEntrypointRulesetOptionsModel.SetKind("managed")
				updateAccountEntrypointRulesetOptionsModel.SetName("testString")
				updateAccountEntrypointRulesetOptionsModel.SetPhase("ddos_l4")
				updateAccountEntrypointRulesetOptionsModel.SetRules([]rulesetsv1.RuleCreate{*ruleCreateModel})
				updateAccountEntrypointRulesetOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateAccountEntrypointRulesetOptionsModel).ToNot(BeNil())
				Expect(updateAccountEntrypointRulesetOptionsModel.RulesetPhase).To(Equal(core.StringPtr("ddos_l4")))
				Expect(updateAccountEntrypointRulesetOptionsModel.Description).To(Equal(core.StringPtr("Custom account ruleset")))
				Expect(updateAccountEntrypointRulesetOptionsModel.Kind).To(Equal(core.StringPtr("managed")))
				Expect(updateAccountEntrypointRulesetOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(updateAccountEntrypointRulesetOptionsModel.Phase).To(Equal(core.StringPtr("ddos_l4")))
				Expect(updateAccountEntrypointRulesetOptionsModel.Rules).To(Equal([]rulesetsv1.RuleCreate{*ruleCreateModel}))
				Expect(updateAccountEntrypointRulesetOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateAccountRulesetOptions successfully`, func() {
				// Construct an instance of the RulesOverride model
				rulesOverrideModel := new(rulesetsv1.RulesOverride)
				Expect(rulesOverrideModel).ToNot(BeNil())
				rulesOverrideModel.ID = core.StringPtr("testString")
				rulesOverrideModel.Enabled = core.BoolPtr(true)
				rulesOverrideModel.Action = core.StringPtr("testString")
				rulesOverrideModel.SensitivityLevel = core.StringPtr("high")
				Expect(rulesOverrideModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(rulesOverrideModel.Enabled).To(Equal(core.BoolPtr(true)))
				Expect(rulesOverrideModel.Action).To(Equal(core.StringPtr("testString")))
				Expect(rulesOverrideModel.SensitivityLevel).To(Equal(core.StringPtr("high")))

				// Construct an instance of the CategoriesOverride model
				categoriesOverrideModel := new(rulesetsv1.CategoriesOverride)
				Expect(categoriesOverrideModel).ToNot(BeNil())
				categoriesOverrideModel.Category = core.StringPtr("testString")
				categoriesOverrideModel.Enabled = core.BoolPtr(true)
				categoriesOverrideModel.Action = core.StringPtr("testString")
				Expect(categoriesOverrideModel.Category).To(Equal(core.StringPtr("testString")))
				Expect(categoriesOverrideModel.Enabled).To(Equal(core.BoolPtr(true)))
				Expect(categoriesOverrideModel.Action).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the Overrides model
				overridesModel := new(rulesetsv1.Overrides)
				Expect(overridesModel).ToNot(BeNil())
				overridesModel.Action = core.StringPtr("testString")
				overridesModel.Enabled = core.BoolPtr(true)
				overridesModel.SensitivityLevel = core.StringPtr("high")
				overridesModel.Rules = []rulesetsv1.RulesOverride{*rulesOverrideModel}
				overridesModel.Categories = []rulesetsv1.CategoriesOverride{*categoriesOverrideModel}
				Expect(overridesModel.Action).To(Equal(core.StringPtr("testString")))
				Expect(overridesModel.Enabled).To(Equal(core.BoolPtr(true)))
				Expect(overridesModel.SensitivityLevel).To(Equal(core.StringPtr("high")))
				Expect(overridesModel.Rules).To(Equal([]rulesetsv1.RulesOverride{*rulesOverrideModel}))
				Expect(overridesModel.Categories).To(Equal([]rulesetsv1.CategoriesOverride{*categoriesOverrideModel}))

				// Construct an instance of the ActionParametersResponse model
				actionParametersResponseModel := new(rulesetsv1.ActionParametersResponse)
				Expect(actionParametersResponseModel).ToNot(BeNil())
				actionParametersResponseModel.Content = core.StringPtr(`{"success": false, "error": "you have been blocked"}`)
				actionParametersResponseModel.ContentType = core.StringPtr("application/json")
				actionParametersResponseModel.StatusCode = core.Int64Ptr(int64(400))
				Expect(actionParametersResponseModel.Content).To(Equal(core.StringPtr(`{"success": false, "error": "you have been blocked"}`)))
				Expect(actionParametersResponseModel.ContentType).To(Equal(core.StringPtr("application/json")))
				Expect(actionParametersResponseModel.StatusCode).To(Equal(core.Int64Ptr(int64(400))))

				// Construct an instance of the ActionParameters model
				actionParametersModel := new(rulesetsv1.ActionParameters)
				Expect(actionParametersModel).ToNot(BeNil())
				actionParametersModel.ID = core.StringPtr("testString")
				actionParametersModel.Overrides = overridesModel
				actionParametersModel.Version = core.StringPtr("testString")
				actionParametersModel.Ruleset = core.StringPtr("testString")
				actionParametersModel.Rulesets = []string{"testString"}
				actionParametersModel.Response = actionParametersResponseModel
				Expect(actionParametersModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(actionParametersModel.Overrides).To(Equal(overridesModel))
				Expect(actionParametersModel.Version).To(Equal(core.StringPtr("testString")))
				Expect(actionParametersModel.Ruleset).To(Equal(core.StringPtr("testString")))
				Expect(actionParametersModel.Rulesets).To(Equal([]string{"testString"}))
				Expect(actionParametersModel.Response).To(Equal(actionParametersResponseModel))

				// Construct an instance of the Logging model
				loggingModel := new(rulesetsv1.Logging)
				Expect(loggingModel).ToNot(BeNil())
				loggingModel.Enabled = core.BoolPtr(true)
				Expect(loggingModel.Enabled).To(Equal(core.BoolPtr(true)))

				// Construct an instance of the Position model
				positionModel := new(rulesetsv1.Position)
				Expect(positionModel).ToNot(BeNil())
				positionModel.Before = core.StringPtr("testString")
				positionModel.After = core.StringPtr("testString")
				positionModel.Index = core.Int64Ptr(int64(0))
				Expect(positionModel.Before).To(Equal(core.StringPtr("testString")))
				Expect(positionModel.After).To(Equal(core.StringPtr("testString")))
				Expect(positionModel.Index).To(Equal(core.Int64Ptr(int64(0))))

				// Construct an instance of the RuleCreate model
				ruleCreateModel := new(rulesetsv1.RuleCreate)
				Expect(ruleCreateModel).ToNot(BeNil())
				ruleCreateModel.Action = core.StringPtr("testString")
				ruleCreateModel.ActionParameters = actionParametersModel
				ruleCreateModel.Description = core.StringPtr("testString")
				ruleCreateModel.Enabled = core.BoolPtr(true)
				ruleCreateModel.Expression = core.StringPtr("ip.src ne 1.1.1.1")
				ruleCreateModel.ID = core.StringPtr("testString")
				ruleCreateModel.Logging = loggingModel
				ruleCreateModel.Ref = core.StringPtr("my_ref")
				ruleCreateModel.Position = positionModel
				Expect(ruleCreateModel.Action).To(Equal(core.StringPtr("testString")))
				Expect(ruleCreateModel.ActionParameters).To(Equal(actionParametersModel))
				Expect(ruleCreateModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(ruleCreateModel.Enabled).To(Equal(core.BoolPtr(true)))
				Expect(ruleCreateModel.Expression).To(Equal(core.StringPtr("ip.src ne 1.1.1.1")))
				Expect(ruleCreateModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(ruleCreateModel.Logging).To(Equal(loggingModel))
				Expect(ruleCreateModel.Ref).To(Equal(core.StringPtr("my_ref")))
				Expect(ruleCreateModel.Position).To(Equal(positionModel))

				// Construct an instance of the UpdateAccountRulesetOptions model
				rulesetID := "testString"
				updateAccountRulesetOptionsModel := rulesetsService.NewUpdateAccountRulesetOptions(rulesetID)
				updateAccountRulesetOptionsModel.SetRulesetID("testString")
				updateAccountRulesetOptionsModel.SetDescription("Custom account ruleset")
				updateAccountRulesetOptionsModel.SetKind("managed")
				updateAccountRulesetOptionsModel.SetName("testString")
				updateAccountRulesetOptionsModel.SetPhase("ddos_l4")
				updateAccountRulesetOptionsModel.SetRules([]rulesetsv1.RuleCreate{*ruleCreateModel})
				updateAccountRulesetOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateAccountRulesetOptionsModel).ToNot(BeNil())
				Expect(updateAccountRulesetOptionsModel.RulesetID).To(Equal(core.StringPtr("testString")))
				Expect(updateAccountRulesetOptionsModel.Description).To(Equal(core.StringPtr("Custom account ruleset")))
				Expect(updateAccountRulesetOptionsModel.Kind).To(Equal(core.StringPtr("managed")))
				Expect(updateAccountRulesetOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(updateAccountRulesetOptionsModel.Phase).To(Equal(core.StringPtr("ddos_l4")))
				Expect(updateAccountRulesetOptionsModel.Rules).To(Equal([]rulesetsv1.RuleCreate{*ruleCreateModel}))
				Expect(updateAccountRulesetOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateAccountRulesetRuleOptions successfully`, func() {
				// Construct an instance of the RulesOverride model
				rulesOverrideModel := new(rulesetsv1.RulesOverride)
				Expect(rulesOverrideModel).ToNot(BeNil())
				rulesOverrideModel.ID = core.StringPtr("testString")
				rulesOverrideModel.Enabled = core.BoolPtr(true)
				rulesOverrideModel.Action = core.StringPtr("testString")
				rulesOverrideModel.SensitivityLevel = core.StringPtr("high")
				Expect(rulesOverrideModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(rulesOverrideModel.Enabled).To(Equal(core.BoolPtr(true)))
				Expect(rulesOverrideModel.Action).To(Equal(core.StringPtr("testString")))
				Expect(rulesOverrideModel.SensitivityLevel).To(Equal(core.StringPtr("high")))

				// Construct an instance of the CategoriesOverride model
				categoriesOverrideModel := new(rulesetsv1.CategoriesOverride)
				Expect(categoriesOverrideModel).ToNot(BeNil())
				categoriesOverrideModel.Category = core.StringPtr("testString")
				categoriesOverrideModel.Enabled = core.BoolPtr(true)
				categoriesOverrideModel.Action = core.StringPtr("testString")
				Expect(categoriesOverrideModel.Category).To(Equal(core.StringPtr("testString")))
				Expect(categoriesOverrideModel.Enabled).To(Equal(core.BoolPtr(true)))
				Expect(categoriesOverrideModel.Action).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the Overrides model
				overridesModel := new(rulesetsv1.Overrides)
				Expect(overridesModel).ToNot(BeNil())
				overridesModel.Action = core.StringPtr("testString")
				overridesModel.Enabled = core.BoolPtr(true)
				overridesModel.SensitivityLevel = core.StringPtr("high")
				overridesModel.Rules = []rulesetsv1.RulesOverride{*rulesOverrideModel}
				overridesModel.Categories = []rulesetsv1.CategoriesOverride{*categoriesOverrideModel}
				Expect(overridesModel.Action).To(Equal(core.StringPtr("testString")))
				Expect(overridesModel.Enabled).To(Equal(core.BoolPtr(true)))
				Expect(overridesModel.SensitivityLevel).To(Equal(core.StringPtr("high")))
				Expect(overridesModel.Rules).To(Equal([]rulesetsv1.RulesOverride{*rulesOverrideModel}))
				Expect(overridesModel.Categories).To(Equal([]rulesetsv1.CategoriesOverride{*categoriesOverrideModel}))

				// Construct an instance of the ActionParametersResponse model
				actionParametersResponseModel := new(rulesetsv1.ActionParametersResponse)
				Expect(actionParametersResponseModel).ToNot(BeNil())
				actionParametersResponseModel.Content = core.StringPtr(`{"success": false, "error": "you have been blocked"}`)
				actionParametersResponseModel.ContentType = core.StringPtr("application/json")
				actionParametersResponseModel.StatusCode = core.Int64Ptr(int64(400))
				Expect(actionParametersResponseModel.Content).To(Equal(core.StringPtr(`{"success": false, "error": "you have been blocked"}`)))
				Expect(actionParametersResponseModel.ContentType).To(Equal(core.StringPtr("application/json")))
				Expect(actionParametersResponseModel.StatusCode).To(Equal(core.Int64Ptr(int64(400))))

				// Construct an instance of the ActionParameters model
				actionParametersModel := new(rulesetsv1.ActionParameters)
				Expect(actionParametersModel).ToNot(BeNil())
				actionParametersModel.ID = core.StringPtr("testString")
				actionParametersModel.Overrides = overridesModel
				actionParametersModel.Version = core.StringPtr("testString")
				actionParametersModel.Ruleset = core.StringPtr("testString")
				actionParametersModel.Rulesets = []string{"testString"}
				actionParametersModel.Response = actionParametersResponseModel
				Expect(actionParametersModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(actionParametersModel.Overrides).To(Equal(overridesModel))
				Expect(actionParametersModel.Version).To(Equal(core.StringPtr("testString")))
				Expect(actionParametersModel.Ruleset).To(Equal(core.StringPtr("testString")))
				Expect(actionParametersModel.Rulesets).To(Equal([]string{"testString"}))
				Expect(actionParametersModel.Response).To(Equal(actionParametersResponseModel))

				// Construct an instance of the Logging model
				loggingModel := new(rulesetsv1.Logging)
				Expect(loggingModel).ToNot(BeNil())
				loggingModel.Enabled = core.BoolPtr(true)
				Expect(loggingModel.Enabled).To(Equal(core.BoolPtr(true)))

				// Construct an instance of the Position model
				positionModel := new(rulesetsv1.Position)
				Expect(positionModel).ToNot(BeNil())
				positionModel.Before = core.StringPtr("testString")
				positionModel.After = core.StringPtr("testString")
				positionModel.Index = core.Int64Ptr(int64(0))
				Expect(positionModel.Before).To(Equal(core.StringPtr("testString")))
				Expect(positionModel.After).To(Equal(core.StringPtr("testString")))
				Expect(positionModel.Index).To(Equal(core.Int64Ptr(int64(0))))

				// Construct an instance of the UpdateAccountRulesetRuleOptions model
				rulesetID := "testString"
				ruleID := "testString"
				updateAccountRulesetRuleOptionsModel := rulesetsService.NewUpdateAccountRulesetRuleOptions(rulesetID, ruleID)
				updateAccountRulesetRuleOptionsModel.SetRulesetID("testString")
				updateAccountRulesetRuleOptionsModel.SetRuleID("testString")
				updateAccountRulesetRuleOptionsModel.SetAction("testString")
				updateAccountRulesetRuleOptionsModel.SetActionParameters(actionParametersModel)
				updateAccountRulesetRuleOptionsModel.SetDescription("testString")
				updateAccountRulesetRuleOptionsModel.SetEnabled(true)
				updateAccountRulesetRuleOptionsModel.SetExpression("ip.src ne 1.1.1.1")
				updateAccountRulesetRuleOptionsModel.SetID("testString")
				updateAccountRulesetRuleOptionsModel.SetLogging(loggingModel)
				updateAccountRulesetRuleOptionsModel.SetRef("my_ref")
				updateAccountRulesetRuleOptionsModel.SetPosition(positionModel)
				updateAccountRulesetRuleOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateAccountRulesetRuleOptionsModel).ToNot(BeNil())
				Expect(updateAccountRulesetRuleOptionsModel.RulesetID).To(Equal(core.StringPtr("testString")))
				Expect(updateAccountRulesetRuleOptionsModel.RuleID).To(Equal(core.StringPtr("testString")))
				Expect(updateAccountRulesetRuleOptionsModel.Action).To(Equal(core.StringPtr("testString")))
				Expect(updateAccountRulesetRuleOptionsModel.ActionParameters).To(Equal(actionParametersModel))
				Expect(updateAccountRulesetRuleOptionsModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(updateAccountRulesetRuleOptionsModel.Enabled).To(Equal(core.BoolPtr(true)))
				Expect(updateAccountRulesetRuleOptionsModel.Expression).To(Equal(core.StringPtr("ip.src ne 1.1.1.1")))
				Expect(updateAccountRulesetRuleOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(updateAccountRulesetRuleOptionsModel.Logging).To(Equal(loggingModel))
				Expect(updateAccountRulesetRuleOptionsModel.Ref).To(Equal(core.StringPtr("my_ref")))
				Expect(updateAccountRulesetRuleOptionsModel.Position).To(Equal(positionModel))
				Expect(updateAccountRulesetRuleOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateZoneEntrypointRulesetOptions successfully`, func() {
				// Construct an instance of the RulesOverride model
				rulesOverrideModel := new(rulesetsv1.RulesOverride)
				Expect(rulesOverrideModel).ToNot(BeNil())
				rulesOverrideModel.ID = core.StringPtr("testString")
				rulesOverrideModel.Enabled = core.BoolPtr(true)
				rulesOverrideModel.Action = core.StringPtr("testString")
				rulesOverrideModel.SensitivityLevel = core.StringPtr("high")
				Expect(rulesOverrideModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(rulesOverrideModel.Enabled).To(Equal(core.BoolPtr(true)))
				Expect(rulesOverrideModel.Action).To(Equal(core.StringPtr("testString")))
				Expect(rulesOverrideModel.SensitivityLevel).To(Equal(core.StringPtr("high")))

				// Construct an instance of the CategoriesOverride model
				categoriesOverrideModel := new(rulesetsv1.CategoriesOverride)
				Expect(categoriesOverrideModel).ToNot(BeNil())
				categoriesOverrideModel.Category = core.StringPtr("testString")
				categoriesOverrideModel.Enabled = core.BoolPtr(true)
				categoriesOverrideModel.Action = core.StringPtr("testString")
				Expect(categoriesOverrideModel.Category).To(Equal(core.StringPtr("testString")))
				Expect(categoriesOverrideModel.Enabled).To(Equal(core.BoolPtr(true)))
				Expect(categoriesOverrideModel.Action).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the Overrides model
				overridesModel := new(rulesetsv1.Overrides)
				Expect(overridesModel).ToNot(BeNil())
				overridesModel.Action = core.StringPtr("testString")
				overridesModel.Enabled = core.BoolPtr(true)
				overridesModel.SensitivityLevel = core.StringPtr("high")
				overridesModel.Rules = []rulesetsv1.RulesOverride{*rulesOverrideModel}
				overridesModel.Categories = []rulesetsv1.CategoriesOverride{*categoriesOverrideModel}
				Expect(overridesModel.Action).To(Equal(core.StringPtr("testString")))
				Expect(overridesModel.Enabled).To(Equal(core.BoolPtr(true)))
				Expect(overridesModel.SensitivityLevel).To(Equal(core.StringPtr("high")))
				Expect(overridesModel.Rules).To(Equal([]rulesetsv1.RulesOverride{*rulesOverrideModel}))
				Expect(overridesModel.Categories).To(Equal([]rulesetsv1.CategoriesOverride{*categoriesOverrideModel}))

				// Construct an instance of the ActionParametersResponse model
				actionParametersResponseModel := new(rulesetsv1.ActionParametersResponse)
				Expect(actionParametersResponseModel).ToNot(BeNil())
				actionParametersResponseModel.Content = core.StringPtr(`{"success": false, "error": "you have been blocked"}`)
				actionParametersResponseModel.ContentType = core.StringPtr("application/json")
				actionParametersResponseModel.StatusCode = core.Int64Ptr(int64(400))
				Expect(actionParametersResponseModel.Content).To(Equal(core.StringPtr(`{"success": false, "error": "you have been blocked"}`)))
				Expect(actionParametersResponseModel.ContentType).To(Equal(core.StringPtr("application/json")))
				Expect(actionParametersResponseModel.StatusCode).To(Equal(core.Int64Ptr(int64(400))))

				// Construct an instance of the ActionParameters model
				actionParametersModel := new(rulesetsv1.ActionParameters)
				Expect(actionParametersModel).ToNot(BeNil())
				actionParametersModel.ID = core.StringPtr("testString")
				actionParametersModel.Overrides = overridesModel
				actionParametersModel.Version = core.StringPtr("testString")
				actionParametersModel.Ruleset = core.StringPtr("testString")
				actionParametersModel.Rulesets = []string{"testString"}
				actionParametersModel.Response = actionParametersResponseModel
				Expect(actionParametersModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(actionParametersModel.Overrides).To(Equal(overridesModel))
				Expect(actionParametersModel.Version).To(Equal(core.StringPtr("testString")))
				Expect(actionParametersModel.Ruleset).To(Equal(core.StringPtr("testString")))
				Expect(actionParametersModel.Rulesets).To(Equal([]string{"testString"}))
				Expect(actionParametersModel.Response).To(Equal(actionParametersResponseModel))

				// Construct an instance of the Logging model
				loggingModel := new(rulesetsv1.Logging)
				Expect(loggingModel).ToNot(BeNil())
				loggingModel.Enabled = core.BoolPtr(true)
				Expect(loggingModel.Enabled).To(Equal(core.BoolPtr(true)))

				// Construct an instance of the Position model
				positionModel := new(rulesetsv1.Position)
				Expect(positionModel).ToNot(BeNil())
				positionModel.Before = core.StringPtr("testString")
				positionModel.After = core.StringPtr("testString")
				positionModel.Index = core.Int64Ptr(int64(0))
				Expect(positionModel.Before).To(Equal(core.StringPtr("testString")))
				Expect(positionModel.After).To(Equal(core.StringPtr("testString")))
				Expect(positionModel.Index).To(Equal(core.Int64Ptr(int64(0))))

				// Construct an instance of the RuleCreate model
				ruleCreateModel := new(rulesetsv1.RuleCreate)
				Expect(ruleCreateModel).ToNot(BeNil())
				ruleCreateModel.Action = core.StringPtr("testString")
				ruleCreateModel.ActionParameters = actionParametersModel
				ruleCreateModel.Description = core.StringPtr("testString")
				ruleCreateModel.Enabled = core.BoolPtr(true)
				ruleCreateModel.Expression = core.StringPtr("ip.src ne 1.1.1.1")
				ruleCreateModel.ID = core.StringPtr("testString")
				ruleCreateModel.Logging = loggingModel
				ruleCreateModel.Ref = core.StringPtr("my_ref")
				ruleCreateModel.Position = positionModel
				Expect(ruleCreateModel.Action).To(Equal(core.StringPtr("testString")))
				Expect(ruleCreateModel.ActionParameters).To(Equal(actionParametersModel))
				Expect(ruleCreateModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(ruleCreateModel.Enabled).To(Equal(core.BoolPtr(true)))
				Expect(ruleCreateModel.Expression).To(Equal(core.StringPtr("ip.src ne 1.1.1.1")))
				Expect(ruleCreateModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(ruleCreateModel.Logging).To(Equal(loggingModel))
				Expect(ruleCreateModel.Ref).To(Equal(core.StringPtr("my_ref")))
				Expect(ruleCreateModel.Position).To(Equal(positionModel))

				// Construct an instance of the UpdateZoneEntrypointRulesetOptions model
				rulesetPhase := "ddos_l4"
				updateZoneEntrypointRulesetOptionsModel := rulesetsService.NewUpdateZoneEntrypointRulesetOptions(rulesetPhase)
				updateZoneEntrypointRulesetOptionsModel.SetRulesetPhase("ddos_l4")
				updateZoneEntrypointRulesetOptionsModel.SetDescription("Custom account ruleset")
				updateZoneEntrypointRulesetOptionsModel.SetKind("managed")
				updateZoneEntrypointRulesetOptionsModel.SetName("testString")
				updateZoneEntrypointRulesetOptionsModel.SetPhase("ddos_l4")
				updateZoneEntrypointRulesetOptionsModel.SetRules([]rulesetsv1.RuleCreate{*ruleCreateModel})
				updateZoneEntrypointRulesetOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateZoneEntrypointRulesetOptionsModel).ToNot(BeNil())
				Expect(updateZoneEntrypointRulesetOptionsModel.RulesetPhase).To(Equal(core.StringPtr("ddos_l4")))
				Expect(updateZoneEntrypointRulesetOptionsModel.Description).To(Equal(core.StringPtr("Custom account ruleset")))
				Expect(updateZoneEntrypointRulesetOptionsModel.Kind).To(Equal(core.StringPtr("managed")))
				Expect(updateZoneEntrypointRulesetOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(updateZoneEntrypointRulesetOptionsModel.Phase).To(Equal(core.StringPtr("ddos_l4")))
				Expect(updateZoneEntrypointRulesetOptionsModel.Rules).To(Equal([]rulesetsv1.RuleCreate{*ruleCreateModel}))
				Expect(updateZoneEntrypointRulesetOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateZoneRulesetOptions successfully`, func() {
				// Construct an instance of the RulesOverride model
				rulesOverrideModel := new(rulesetsv1.RulesOverride)
				Expect(rulesOverrideModel).ToNot(BeNil())
				rulesOverrideModel.ID = core.StringPtr("testString")
				rulesOverrideModel.Enabled = core.BoolPtr(true)
				rulesOverrideModel.Action = core.StringPtr("testString")
				rulesOverrideModel.SensitivityLevel = core.StringPtr("high")
				Expect(rulesOverrideModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(rulesOverrideModel.Enabled).To(Equal(core.BoolPtr(true)))
				Expect(rulesOverrideModel.Action).To(Equal(core.StringPtr("testString")))
				Expect(rulesOverrideModel.SensitivityLevel).To(Equal(core.StringPtr("high")))

				// Construct an instance of the CategoriesOverride model
				categoriesOverrideModel := new(rulesetsv1.CategoriesOverride)
				Expect(categoriesOverrideModel).ToNot(BeNil())
				categoriesOverrideModel.Category = core.StringPtr("testString")
				categoriesOverrideModel.Enabled = core.BoolPtr(true)
				categoriesOverrideModel.Action = core.StringPtr("testString")
				Expect(categoriesOverrideModel.Category).To(Equal(core.StringPtr("testString")))
				Expect(categoriesOverrideModel.Enabled).To(Equal(core.BoolPtr(true)))
				Expect(categoriesOverrideModel.Action).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the Overrides model
				overridesModel := new(rulesetsv1.Overrides)
				Expect(overridesModel).ToNot(BeNil())
				overridesModel.Action = core.StringPtr("testString")
				overridesModel.Enabled = core.BoolPtr(true)
				overridesModel.SensitivityLevel = core.StringPtr("high")
				overridesModel.Rules = []rulesetsv1.RulesOverride{*rulesOverrideModel}
				overridesModel.Categories = []rulesetsv1.CategoriesOverride{*categoriesOverrideModel}
				Expect(overridesModel.Action).To(Equal(core.StringPtr("testString")))
				Expect(overridesModel.Enabled).To(Equal(core.BoolPtr(true)))
				Expect(overridesModel.SensitivityLevel).To(Equal(core.StringPtr("high")))
				Expect(overridesModel.Rules).To(Equal([]rulesetsv1.RulesOverride{*rulesOverrideModel}))
				Expect(overridesModel.Categories).To(Equal([]rulesetsv1.CategoriesOverride{*categoriesOverrideModel}))

				// Construct an instance of the ActionParametersResponse model
				actionParametersResponseModel := new(rulesetsv1.ActionParametersResponse)
				Expect(actionParametersResponseModel).ToNot(BeNil())
				actionParametersResponseModel.Content = core.StringPtr(`{"success": false, "error": "you have been blocked"}`)
				actionParametersResponseModel.ContentType = core.StringPtr("application/json")
				actionParametersResponseModel.StatusCode = core.Int64Ptr(int64(400))
				Expect(actionParametersResponseModel.Content).To(Equal(core.StringPtr(`{"success": false, "error": "you have been blocked"}`)))
				Expect(actionParametersResponseModel.ContentType).To(Equal(core.StringPtr("application/json")))
				Expect(actionParametersResponseModel.StatusCode).To(Equal(core.Int64Ptr(int64(400))))

				// Construct an instance of the ActionParameters model
				actionParametersModel := new(rulesetsv1.ActionParameters)
				Expect(actionParametersModel).ToNot(BeNil())
				actionParametersModel.ID = core.StringPtr("testString")
				actionParametersModel.Overrides = overridesModel
				actionParametersModel.Version = core.StringPtr("testString")
				actionParametersModel.Ruleset = core.StringPtr("testString")
				actionParametersModel.Rulesets = []string{"testString"}
				actionParametersModel.Response = actionParametersResponseModel
				Expect(actionParametersModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(actionParametersModel.Overrides).To(Equal(overridesModel))
				Expect(actionParametersModel.Version).To(Equal(core.StringPtr("testString")))
				Expect(actionParametersModel.Ruleset).To(Equal(core.StringPtr("testString")))
				Expect(actionParametersModel.Rulesets).To(Equal([]string{"testString"}))
				Expect(actionParametersModel.Response).To(Equal(actionParametersResponseModel))

				// Construct an instance of the Logging model
				loggingModel := new(rulesetsv1.Logging)
				Expect(loggingModel).ToNot(BeNil())
				loggingModel.Enabled = core.BoolPtr(true)
				Expect(loggingModel.Enabled).To(Equal(core.BoolPtr(true)))

				// Construct an instance of the Position model
				positionModel := new(rulesetsv1.Position)
				Expect(positionModel).ToNot(BeNil())
				positionModel.Before = core.StringPtr("testString")
				positionModel.After = core.StringPtr("testString")
				positionModel.Index = core.Int64Ptr(int64(0))
				Expect(positionModel.Before).To(Equal(core.StringPtr("testString")))
				Expect(positionModel.After).To(Equal(core.StringPtr("testString")))
				Expect(positionModel.Index).To(Equal(core.Int64Ptr(int64(0))))

				// Construct an instance of the RuleCreate model
				ruleCreateModel := new(rulesetsv1.RuleCreate)
				Expect(ruleCreateModel).ToNot(BeNil())
				ruleCreateModel.Action = core.StringPtr("testString")
				ruleCreateModel.ActionParameters = actionParametersModel
				ruleCreateModel.Description = core.StringPtr("testString")
				ruleCreateModel.Enabled = core.BoolPtr(true)
				ruleCreateModel.Expression = core.StringPtr("ip.src ne 1.1.1.1")
				ruleCreateModel.ID = core.StringPtr("testString")
				ruleCreateModel.Logging = loggingModel
				ruleCreateModel.Ref = core.StringPtr("my_ref")
				ruleCreateModel.Position = positionModel
				Expect(ruleCreateModel.Action).To(Equal(core.StringPtr("testString")))
				Expect(ruleCreateModel.ActionParameters).To(Equal(actionParametersModel))
				Expect(ruleCreateModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(ruleCreateModel.Enabled).To(Equal(core.BoolPtr(true)))
				Expect(ruleCreateModel.Expression).To(Equal(core.StringPtr("ip.src ne 1.1.1.1")))
				Expect(ruleCreateModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(ruleCreateModel.Logging).To(Equal(loggingModel))
				Expect(ruleCreateModel.Ref).To(Equal(core.StringPtr("my_ref")))
				Expect(ruleCreateModel.Position).To(Equal(positionModel))

				// Construct an instance of the UpdateZoneRulesetOptions model
				rulesetID := "testString"
				updateZoneRulesetOptionsModel := rulesetsService.NewUpdateZoneRulesetOptions(rulesetID)
				updateZoneRulesetOptionsModel.SetRulesetID("testString")
				updateZoneRulesetOptionsModel.SetDescription("Custom account ruleset")
				updateZoneRulesetOptionsModel.SetKind("managed")
				updateZoneRulesetOptionsModel.SetName("testString")
				updateZoneRulesetOptionsModel.SetPhase("ddos_l4")
				updateZoneRulesetOptionsModel.SetRules([]rulesetsv1.RuleCreate{*ruleCreateModel})
				updateZoneRulesetOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateZoneRulesetOptionsModel).ToNot(BeNil())
				Expect(updateZoneRulesetOptionsModel.RulesetID).To(Equal(core.StringPtr("testString")))
				Expect(updateZoneRulesetOptionsModel.Description).To(Equal(core.StringPtr("Custom account ruleset")))
				Expect(updateZoneRulesetOptionsModel.Kind).To(Equal(core.StringPtr("managed")))
				Expect(updateZoneRulesetOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(updateZoneRulesetOptionsModel.Phase).To(Equal(core.StringPtr("ddos_l4")))
				Expect(updateZoneRulesetOptionsModel.Rules).To(Equal([]rulesetsv1.RuleCreate{*ruleCreateModel}))
				Expect(updateZoneRulesetOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateZoneRulesetRuleOptions successfully`, func() {
				// Construct an instance of the RulesOverride model
				rulesOverrideModel := new(rulesetsv1.RulesOverride)
				Expect(rulesOverrideModel).ToNot(BeNil())
				rulesOverrideModel.ID = core.StringPtr("testString")
				rulesOverrideModel.Enabled = core.BoolPtr(true)
				rulesOverrideModel.Action = core.StringPtr("testString")
				rulesOverrideModel.SensitivityLevel = core.StringPtr("high")
				Expect(rulesOverrideModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(rulesOverrideModel.Enabled).To(Equal(core.BoolPtr(true)))
				Expect(rulesOverrideModel.Action).To(Equal(core.StringPtr("testString")))
				Expect(rulesOverrideModel.SensitivityLevel).To(Equal(core.StringPtr("high")))

				// Construct an instance of the CategoriesOverride model
				categoriesOverrideModel := new(rulesetsv1.CategoriesOverride)
				Expect(categoriesOverrideModel).ToNot(BeNil())
				categoriesOverrideModel.Category = core.StringPtr("testString")
				categoriesOverrideModel.Enabled = core.BoolPtr(true)
				categoriesOverrideModel.Action = core.StringPtr("testString")
				Expect(categoriesOverrideModel.Category).To(Equal(core.StringPtr("testString")))
				Expect(categoriesOverrideModel.Enabled).To(Equal(core.BoolPtr(true)))
				Expect(categoriesOverrideModel.Action).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the Overrides model
				overridesModel := new(rulesetsv1.Overrides)
				Expect(overridesModel).ToNot(BeNil())
				overridesModel.Action = core.StringPtr("testString")
				overridesModel.Enabled = core.BoolPtr(true)
				overridesModel.SensitivityLevel = core.StringPtr("high")
				overridesModel.Rules = []rulesetsv1.RulesOverride{*rulesOverrideModel}
				overridesModel.Categories = []rulesetsv1.CategoriesOverride{*categoriesOverrideModel}
				Expect(overridesModel.Action).To(Equal(core.StringPtr("testString")))
				Expect(overridesModel.Enabled).To(Equal(core.BoolPtr(true)))
				Expect(overridesModel.SensitivityLevel).To(Equal(core.StringPtr("high")))
				Expect(overridesModel.Rules).To(Equal([]rulesetsv1.RulesOverride{*rulesOverrideModel}))
				Expect(overridesModel.Categories).To(Equal([]rulesetsv1.CategoriesOverride{*categoriesOverrideModel}))

				// Construct an instance of the ActionParametersResponse model
				actionParametersResponseModel := new(rulesetsv1.ActionParametersResponse)
				Expect(actionParametersResponseModel).ToNot(BeNil())
				actionParametersResponseModel.Content = core.StringPtr(`{"success": false, "error": "you have been blocked"}`)
				actionParametersResponseModel.ContentType = core.StringPtr("application/json")
				actionParametersResponseModel.StatusCode = core.Int64Ptr(int64(400))
				Expect(actionParametersResponseModel.Content).To(Equal(core.StringPtr(`{"success": false, "error": "you have been blocked"}`)))
				Expect(actionParametersResponseModel.ContentType).To(Equal(core.StringPtr("application/json")))
				Expect(actionParametersResponseModel.StatusCode).To(Equal(core.Int64Ptr(int64(400))))

				// Construct an instance of the ActionParameters model
				actionParametersModel := new(rulesetsv1.ActionParameters)
				Expect(actionParametersModel).ToNot(BeNil())
				actionParametersModel.ID = core.StringPtr("testString")
				actionParametersModel.Overrides = overridesModel
				actionParametersModel.Version = core.StringPtr("testString")
				actionParametersModel.Ruleset = core.StringPtr("testString")
				actionParametersModel.Rulesets = []string{"testString"}
				actionParametersModel.Response = actionParametersResponseModel
				Expect(actionParametersModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(actionParametersModel.Overrides).To(Equal(overridesModel))
				Expect(actionParametersModel.Version).To(Equal(core.StringPtr("testString")))
				Expect(actionParametersModel.Ruleset).To(Equal(core.StringPtr("testString")))
				Expect(actionParametersModel.Rulesets).To(Equal([]string{"testString"}))
				Expect(actionParametersModel.Response).To(Equal(actionParametersResponseModel))

				// Construct an instance of the Logging model
				loggingModel := new(rulesetsv1.Logging)
				Expect(loggingModel).ToNot(BeNil())
				loggingModel.Enabled = core.BoolPtr(true)
				Expect(loggingModel.Enabled).To(Equal(core.BoolPtr(true)))

				// Construct an instance of the Position model
				positionModel := new(rulesetsv1.Position)
				Expect(positionModel).ToNot(BeNil())
				positionModel.Before = core.StringPtr("testString")
				positionModel.After = core.StringPtr("testString")
				positionModel.Index = core.Int64Ptr(int64(0))
				Expect(positionModel.Before).To(Equal(core.StringPtr("testString")))
				Expect(positionModel.After).To(Equal(core.StringPtr("testString")))
				Expect(positionModel.Index).To(Equal(core.Int64Ptr(int64(0))))

				// Construct an instance of the UpdateZoneRulesetRuleOptions model
				rulesetID := "testString"
				ruleID := "testString"
				updateZoneRulesetRuleOptionsModel := rulesetsService.NewUpdateZoneRulesetRuleOptions(rulesetID, ruleID)
				updateZoneRulesetRuleOptionsModel.SetRulesetID("testString")
				updateZoneRulesetRuleOptionsModel.SetRuleID("testString")
				updateZoneRulesetRuleOptionsModel.SetAction("testString")
				updateZoneRulesetRuleOptionsModel.SetActionParameters(actionParametersModel)
				updateZoneRulesetRuleOptionsModel.SetDescription("testString")
				updateZoneRulesetRuleOptionsModel.SetEnabled(true)
				updateZoneRulesetRuleOptionsModel.SetExpression("ip.src ne 1.1.1.1")
				updateZoneRulesetRuleOptionsModel.SetID("testString")
				updateZoneRulesetRuleOptionsModel.SetLogging(loggingModel)
				updateZoneRulesetRuleOptionsModel.SetRef("my_ref")
				updateZoneRulesetRuleOptionsModel.SetPosition(positionModel)
				updateZoneRulesetRuleOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateZoneRulesetRuleOptionsModel).ToNot(BeNil())
				Expect(updateZoneRulesetRuleOptionsModel.RulesetID).To(Equal(core.StringPtr("testString")))
				Expect(updateZoneRulesetRuleOptionsModel.RuleID).To(Equal(core.StringPtr("testString")))
				Expect(updateZoneRulesetRuleOptionsModel.Action).To(Equal(core.StringPtr("testString")))
				Expect(updateZoneRulesetRuleOptionsModel.ActionParameters).To(Equal(actionParametersModel))
				Expect(updateZoneRulesetRuleOptionsModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(updateZoneRulesetRuleOptionsModel.Enabled).To(Equal(core.BoolPtr(true)))
				Expect(updateZoneRulesetRuleOptionsModel.Expression).To(Equal(core.StringPtr("ip.src ne 1.1.1.1")))
				Expect(updateZoneRulesetRuleOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(updateZoneRulesetRuleOptionsModel.Logging).To(Equal(loggingModel))
				Expect(updateZoneRulesetRuleOptionsModel.Ref).To(Equal(core.StringPtr("my_ref")))
				Expect(updateZoneRulesetRuleOptionsModel.Position).To(Equal(positionModel))
				Expect(updateZoneRulesetRuleOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewLogging successfully`, func() {
				enabled := true
				_model, err := rulesetsService.NewLogging(enabled)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewRuleCreate successfully`, func() {
				action := "testString"
				expression := "ip.src ne 1.1.1.1"
				_model, err := rulesetsService.NewRuleCreate(action, expression)
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
