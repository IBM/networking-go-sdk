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

package wafrulegroupsapiv1_test

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
	"github.com/IBM/networking-go-sdk/wafrulegroupsapiv1"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe(`WafRuleGroupsApiV1`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		crn := "testString"
		zoneID := "testString"
		It(`Instantiate service client`, func() {
			wafRuleGroupsApiService, serviceErr := wafrulegroupsapiv1.NewWafRuleGroupsApiV1(&wafrulegroupsapiv1.WafRuleGroupsApiV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
				Crn:           core.StringPtr(crn),
				ZoneID:        core.StringPtr(zoneID),
			})
			Expect(wafRuleGroupsApiService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			wafRuleGroupsApiService, serviceErr := wafrulegroupsapiv1.NewWafRuleGroupsApiV1(&wafrulegroupsapiv1.WafRuleGroupsApiV1Options{
				URL:    "{BAD_URL_STRING",
				Crn:    core.StringPtr(crn),
				ZoneID: core.StringPtr(zoneID),
			})
			Expect(wafRuleGroupsApiService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			wafRuleGroupsApiService, serviceErr := wafrulegroupsapiv1.NewWafRuleGroupsApiV1(&wafrulegroupsapiv1.WafRuleGroupsApiV1Options{
				URL:    "https://wafrulegroupsapiv1/api",
				Crn:    core.StringPtr(crn),
				ZoneID: core.StringPtr(zoneID),
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(wafRuleGroupsApiService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Validation Error`, func() {
			wafRuleGroupsApiService, serviceErr := wafrulegroupsapiv1.NewWafRuleGroupsApiV1(&wafrulegroupsapiv1.WafRuleGroupsApiV1Options{})
			Expect(wafRuleGroupsApiService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		crn := "testString"
		zoneID := "testString"
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"WAF_RULE_GROUPS_API_URL":       "https://wafrulegroupsapiv1/api",
				"WAF_RULE_GROUPS_API_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				wafRuleGroupsApiService, serviceErr := wafrulegroupsapiv1.NewWafRuleGroupsApiV1UsingExternalConfig(&wafrulegroupsapiv1.WafRuleGroupsApiV1Options{
					Crn:    core.StringPtr(crn),
					ZoneID: core.StringPtr(zoneID),
				})
				Expect(wafRuleGroupsApiService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := wafRuleGroupsApiService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != wafRuleGroupsApiService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(wafRuleGroupsApiService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(wafRuleGroupsApiService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				wafRuleGroupsApiService, serviceErr := wafrulegroupsapiv1.NewWafRuleGroupsApiV1UsingExternalConfig(&wafrulegroupsapiv1.WafRuleGroupsApiV1Options{
					URL:    "https://testService/api",
					Crn:    core.StringPtr(crn),
					ZoneID: core.StringPtr(zoneID),
				})
				Expect(wafRuleGroupsApiService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(wafRuleGroupsApiService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := wafRuleGroupsApiService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != wafRuleGroupsApiService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(wafRuleGroupsApiService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(wafRuleGroupsApiService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				wafRuleGroupsApiService, serviceErr := wafrulegroupsapiv1.NewWafRuleGroupsApiV1UsingExternalConfig(&wafrulegroupsapiv1.WafRuleGroupsApiV1Options{
					Crn:    core.StringPtr(crn),
					ZoneID: core.StringPtr(zoneID),
				})
				err := wafRuleGroupsApiService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(wafRuleGroupsApiService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(wafRuleGroupsApiService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := wafRuleGroupsApiService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != wafRuleGroupsApiService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(wafRuleGroupsApiService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(wafRuleGroupsApiService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"WAF_RULE_GROUPS_API_URL":       "https://wafrulegroupsapiv1/api",
				"WAF_RULE_GROUPS_API_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			wafRuleGroupsApiService, serviceErr := wafrulegroupsapiv1.NewWafRuleGroupsApiV1UsingExternalConfig(&wafrulegroupsapiv1.WafRuleGroupsApiV1Options{
				Crn:    core.StringPtr(crn),
				ZoneID: core.StringPtr(zoneID),
			})

			It(`Instantiate service client with error`, func() {
				Expect(wafRuleGroupsApiService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"WAF_RULE_GROUPS_API_AUTH_TYPE": "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			wafRuleGroupsApiService, serviceErr := wafrulegroupsapiv1.NewWafRuleGroupsApiV1UsingExternalConfig(&wafrulegroupsapiv1.WafRuleGroupsApiV1Options{
				URL:    "{BAD_URL_STRING",
				Crn:    core.StringPtr(crn),
				ZoneID: core.StringPtr(zoneID),
			})

			It(`Instantiate service client with error`, func() {
				Expect(wafRuleGroupsApiService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = wafrulegroupsapiv1.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`ListWafRuleGroups(listWafRuleGroupsOptions *ListWafRuleGroupsOptions) - Operation response error`, func() {
		crn := "testString"
		zoneID := "testString"
		listWafRuleGroupsPath := "/v1/testString/zones/testString/firewall/waf/packages/testString/groups"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listWafRuleGroupsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["name"]).To(Equal([]string{"Wordpress-rules"}))

					Expect(req.URL.Query()["mode"]).To(Equal([]string{"true"}))

					Expect(req.URL.Query()["rules_count"]).To(Equal([]string{"10"}))

					Expect(req.URL.Query()["page"]).To(Equal([]string{fmt.Sprint(int64(1))}))

					Expect(req.URL.Query()["per_page"]).To(Equal([]string{fmt.Sprint(int64(50))}))

					Expect(req.URL.Query()["order"]).To(Equal([]string{"status"}))

					Expect(req.URL.Query()["direction"]).To(Equal([]string{"desc"}))

					Expect(req.URL.Query()["match"]).To(Equal([]string{"all"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListWafRuleGroups with error: Operation response processing error`, func() {
				wafRuleGroupsApiService, serviceErr := wafrulegroupsapiv1.NewWafRuleGroupsApiV1(&wafrulegroupsapiv1.WafRuleGroupsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ZoneID:        core.StringPtr(zoneID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(wafRuleGroupsApiService).ToNot(BeNil())

				// Construct an instance of the ListWafRuleGroupsOptions model
				listWafRuleGroupsOptionsModel := new(wafrulegroupsapiv1.ListWafRuleGroupsOptions)
				listWafRuleGroupsOptionsModel.PkgID = core.StringPtr("testString")
				listWafRuleGroupsOptionsModel.Name = core.StringPtr("Wordpress-rules")
				listWafRuleGroupsOptionsModel.Mode = core.StringPtr("true")
				listWafRuleGroupsOptionsModel.RulesCount = core.StringPtr("10")
				listWafRuleGroupsOptionsModel.Page = core.Int64Ptr(int64(1))
				listWafRuleGroupsOptionsModel.PerPage = core.Int64Ptr(int64(50))
				listWafRuleGroupsOptionsModel.Order = core.StringPtr("status")
				listWafRuleGroupsOptionsModel.Direction = core.StringPtr("desc")
				listWafRuleGroupsOptionsModel.Match = core.StringPtr("all")
				listWafRuleGroupsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := wafRuleGroupsApiService.ListWafRuleGroups(listWafRuleGroupsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				wafRuleGroupsApiService.EnableRetries(0, 0)
				result, response, operationErr = wafRuleGroupsApiService.ListWafRuleGroups(listWafRuleGroupsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ListWafRuleGroups(listWafRuleGroupsOptions *ListWafRuleGroupsOptions)`, func() {
		crn := "testString"
		zoneID := "testString"
		listWafRuleGroupsPath := "/v1/testString/zones/testString/firewall/waf/packages/testString/groups"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listWafRuleGroupsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["name"]).To(Equal([]string{"Wordpress-rules"}))

					Expect(req.URL.Query()["mode"]).To(Equal([]string{"true"}))

					Expect(req.URL.Query()["rules_count"]).To(Equal([]string{"10"}))

					Expect(req.URL.Query()["page"]).To(Equal([]string{fmt.Sprint(int64(1))}))

					Expect(req.URL.Query()["per_page"]).To(Equal([]string{fmt.Sprint(int64(50))}))

					Expect(req.URL.Query()["order"]).To(Equal([]string{"status"}))

					Expect(req.URL.Query()["direction"]).To(Equal([]string{"desc"}))

					Expect(req.URL.Query()["match"]).To(Equal([]string{"all"}))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": [{"id": "a25a9a7e9c00afc1fb2e0245519d725b", "name": "Project Honey Pot", "description": "Group designed to protect against IP addresses that are a threat and typically used to launch DDoS attacks", "rules_count": 10, "modified_rules_count": 10, "package_id": "a25a9a7e9c00afc1fb2e0245519d725b", "mode": "on", "allowed_modes": ["AllowedModes"]}], "result_info": {"page": 1, "per_page": 2, "count": 1, "total_count": 200}}`)
				}))
			})
			It(`Invoke ListWafRuleGroups successfully`, func() {
				wafRuleGroupsApiService, serviceErr := wafrulegroupsapiv1.NewWafRuleGroupsApiV1(&wafrulegroupsapiv1.WafRuleGroupsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ZoneID:        core.StringPtr(zoneID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(wafRuleGroupsApiService).ToNot(BeNil())
				wafRuleGroupsApiService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := wafRuleGroupsApiService.ListWafRuleGroups(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListWafRuleGroupsOptions model
				listWafRuleGroupsOptionsModel := new(wafrulegroupsapiv1.ListWafRuleGroupsOptions)
				listWafRuleGroupsOptionsModel.PkgID = core.StringPtr("testString")
				listWafRuleGroupsOptionsModel.Name = core.StringPtr("Wordpress-rules")
				listWafRuleGroupsOptionsModel.Mode = core.StringPtr("true")
				listWafRuleGroupsOptionsModel.RulesCount = core.StringPtr("10")
				listWafRuleGroupsOptionsModel.Page = core.Int64Ptr(int64(1))
				listWafRuleGroupsOptionsModel.PerPage = core.Int64Ptr(int64(50))
				listWafRuleGroupsOptionsModel.Order = core.StringPtr("status")
				listWafRuleGroupsOptionsModel.Direction = core.StringPtr("desc")
				listWafRuleGroupsOptionsModel.Match = core.StringPtr("all")
				listWafRuleGroupsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = wafRuleGroupsApiService.ListWafRuleGroups(listWafRuleGroupsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = wafRuleGroupsApiService.ListWafRuleGroupsWithContext(ctx, listWafRuleGroupsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				wafRuleGroupsApiService.DisableRetries()
				result, response, operationErr = wafRuleGroupsApiService.ListWafRuleGroups(listWafRuleGroupsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = wafRuleGroupsApiService.ListWafRuleGroupsWithContext(ctx, listWafRuleGroupsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke ListWafRuleGroups with error: Operation validation and request error`, func() {
				wafRuleGroupsApiService, serviceErr := wafrulegroupsapiv1.NewWafRuleGroupsApiV1(&wafrulegroupsapiv1.WafRuleGroupsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ZoneID:        core.StringPtr(zoneID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(wafRuleGroupsApiService).ToNot(BeNil())

				// Construct an instance of the ListWafRuleGroupsOptions model
				listWafRuleGroupsOptionsModel := new(wafrulegroupsapiv1.ListWafRuleGroupsOptions)
				listWafRuleGroupsOptionsModel.PkgID = core.StringPtr("testString")
				listWafRuleGroupsOptionsModel.Name = core.StringPtr("Wordpress-rules")
				listWafRuleGroupsOptionsModel.Mode = core.StringPtr("true")
				listWafRuleGroupsOptionsModel.RulesCount = core.StringPtr("10")
				listWafRuleGroupsOptionsModel.Page = core.Int64Ptr(int64(1))
				listWafRuleGroupsOptionsModel.PerPage = core.Int64Ptr(int64(50))
				listWafRuleGroupsOptionsModel.Order = core.StringPtr("status")
				listWafRuleGroupsOptionsModel.Direction = core.StringPtr("desc")
				listWafRuleGroupsOptionsModel.Match = core.StringPtr("all")
				listWafRuleGroupsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := wafRuleGroupsApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := wafRuleGroupsApiService.ListWafRuleGroups(listWafRuleGroupsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListWafRuleGroupsOptions model with no property values
				listWafRuleGroupsOptionsModelNew := new(wafrulegroupsapiv1.ListWafRuleGroupsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = wafRuleGroupsApiService.ListWafRuleGroups(listWafRuleGroupsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetWafRuleGroup(getWafRuleGroupOptions *GetWafRuleGroupOptions) - Operation response error`, func() {
		crn := "testString"
		zoneID := "testString"
		getWafRuleGroupPath := "/v1/testString/zones/testString/firewall/waf/packages/testString/groups/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getWafRuleGroupPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetWafRuleGroup with error: Operation response processing error`, func() {
				wafRuleGroupsApiService, serviceErr := wafrulegroupsapiv1.NewWafRuleGroupsApiV1(&wafrulegroupsapiv1.WafRuleGroupsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ZoneID:        core.StringPtr(zoneID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(wafRuleGroupsApiService).ToNot(BeNil())

				// Construct an instance of the GetWafRuleGroupOptions model
				getWafRuleGroupOptionsModel := new(wafrulegroupsapiv1.GetWafRuleGroupOptions)
				getWafRuleGroupOptionsModel.PkgID = core.StringPtr("testString")
				getWafRuleGroupOptionsModel.GroupID = core.StringPtr("testString")
				getWafRuleGroupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := wafRuleGroupsApiService.GetWafRuleGroup(getWafRuleGroupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				wafRuleGroupsApiService.EnableRetries(0, 0)
				result, response, operationErr = wafRuleGroupsApiService.GetWafRuleGroup(getWafRuleGroupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetWafRuleGroup(getWafRuleGroupOptions *GetWafRuleGroupOptions)`, func() {
		crn := "testString"
		zoneID := "testString"
		getWafRuleGroupPath := "/v1/testString/zones/testString/firewall/waf/packages/testString/groups/testString"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getWafRuleGroupPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": "a25a9a7e9c00afc1fb2e0245519d725b", "name": "Project Honey Pot", "description": "Group designed to protect against IP addresses that are a threat and typically used to launch DDoS attacks", "rules_count": 10, "modified_rules_count": 10, "package_id": "a25a9a7e9c00afc1fb2e0245519d725b", "mode": "on", "allowed_modes": ["AllowedModes"]}, "result_info": {"page": 1, "per_page": 2, "count": 1, "total_count": 200}}`)
				}))
			})
			It(`Invoke GetWafRuleGroup successfully`, func() {
				wafRuleGroupsApiService, serviceErr := wafrulegroupsapiv1.NewWafRuleGroupsApiV1(&wafrulegroupsapiv1.WafRuleGroupsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ZoneID:        core.StringPtr(zoneID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(wafRuleGroupsApiService).ToNot(BeNil())
				wafRuleGroupsApiService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := wafRuleGroupsApiService.GetWafRuleGroup(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetWafRuleGroupOptions model
				getWafRuleGroupOptionsModel := new(wafrulegroupsapiv1.GetWafRuleGroupOptions)
				getWafRuleGroupOptionsModel.PkgID = core.StringPtr("testString")
				getWafRuleGroupOptionsModel.GroupID = core.StringPtr("testString")
				getWafRuleGroupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = wafRuleGroupsApiService.GetWafRuleGroup(getWafRuleGroupOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = wafRuleGroupsApiService.GetWafRuleGroupWithContext(ctx, getWafRuleGroupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				wafRuleGroupsApiService.DisableRetries()
				result, response, operationErr = wafRuleGroupsApiService.GetWafRuleGroup(getWafRuleGroupOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = wafRuleGroupsApiService.GetWafRuleGroupWithContext(ctx, getWafRuleGroupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke GetWafRuleGroup with error: Operation validation and request error`, func() {
				wafRuleGroupsApiService, serviceErr := wafrulegroupsapiv1.NewWafRuleGroupsApiV1(&wafrulegroupsapiv1.WafRuleGroupsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ZoneID:        core.StringPtr(zoneID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(wafRuleGroupsApiService).ToNot(BeNil())

				// Construct an instance of the GetWafRuleGroupOptions model
				getWafRuleGroupOptionsModel := new(wafrulegroupsapiv1.GetWafRuleGroupOptions)
				getWafRuleGroupOptionsModel.PkgID = core.StringPtr("testString")
				getWafRuleGroupOptionsModel.GroupID = core.StringPtr("testString")
				getWafRuleGroupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := wafRuleGroupsApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := wafRuleGroupsApiService.GetWafRuleGroup(getWafRuleGroupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetWafRuleGroupOptions model with no property values
				getWafRuleGroupOptionsModelNew := new(wafrulegroupsapiv1.GetWafRuleGroupOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = wafRuleGroupsApiService.GetWafRuleGroup(getWafRuleGroupOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateWafRuleGroup(updateWafRuleGroupOptions *UpdateWafRuleGroupOptions) - Operation response error`, func() {
		crn := "testString"
		zoneID := "testString"
		updateWafRuleGroupPath := "/v1/testString/zones/testString/firewall/waf/packages/testString/groups/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateWafRuleGroupPath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateWafRuleGroup with error: Operation response processing error`, func() {
				wafRuleGroupsApiService, serviceErr := wafrulegroupsapiv1.NewWafRuleGroupsApiV1(&wafrulegroupsapiv1.WafRuleGroupsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ZoneID:        core.StringPtr(zoneID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(wafRuleGroupsApiService).ToNot(BeNil())

				// Construct an instance of the UpdateWafRuleGroupOptions model
				updateWafRuleGroupOptionsModel := new(wafrulegroupsapiv1.UpdateWafRuleGroupOptions)
				updateWafRuleGroupOptionsModel.PkgID = core.StringPtr("testString")
				updateWafRuleGroupOptionsModel.GroupID = core.StringPtr("testString")
				updateWafRuleGroupOptionsModel.Mode = core.StringPtr("on")
				updateWafRuleGroupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := wafRuleGroupsApiService.UpdateWafRuleGroup(updateWafRuleGroupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				wafRuleGroupsApiService.EnableRetries(0, 0)
				result, response, operationErr = wafRuleGroupsApiService.UpdateWafRuleGroup(updateWafRuleGroupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UpdateWafRuleGroup(updateWafRuleGroupOptions *UpdateWafRuleGroupOptions)`, func() {
		crn := "testString"
		zoneID := "testString"
		updateWafRuleGroupPath := "/v1/testString/zones/testString/firewall/waf/packages/testString/groups/testString"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateWafRuleGroupPath))
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
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": "a25a9a7e9c00afc1fb2e0245519d725b", "name": "Project Honey Pot", "description": "Group designed to protect against IP addresses that are a threat and typically used to launch DDoS attacks", "rules_count": 10, "modified_rules_count": 10, "package_id": "a25a9a7e9c00afc1fb2e0245519d725b", "mode": "on", "allowed_modes": ["AllowedModes"]}, "result_info": {"page": 1, "per_page": 2, "count": 1, "total_count": 200}}`)
				}))
			})
			It(`Invoke UpdateWafRuleGroup successfully`, func() {
				wafRuleGroupsApiService, serviceErr := wafrulegroupsapiv1.NewWafRuleGroupsApiV1(&wafrulegroupsapiv1.WafRuleGroupsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ZoneID:        core.StringPtr(zoneID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(wafRuleGroupsApiService).ToNot(BeNil())
				wafRuleGroupsApiService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := wafRuleGroupsApiService.UpdateWafRuleGroup(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateWafRuleGroupOptions model
				updateWafRuleGroupOptionsModel := new(wafrulegroupsapiv1.UpdateWafRuleGroupOptions)
				updateWafRuleGroupOptionsModel.PkgID = core.StringPtr("testString")
				updateWafRuleGroupOptionsModel.GroupID = core.StringPtr("testString")
				updateWafRuleGroupOptionsModel.Mode = core.StringPtr("on")
				updateWafRuleGroupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = wafRuleGroupsApiService.UpdateWafRuleGroup(updateWafRuleGroupOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = wafRuleGroupsApiService.UpdateWafRuleGroupWithContext(ctx, updateWafRuleGroupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				wafRuleGroupsApiService.DisableRetries()
				result, response, operationErr = wafRuleGroupsApiService.UpdateWafRuleGroup(updateWafRuleGroupOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = wafRuleGroupsApiService.UpdateWafRuleGroupWithContext(ctx, updateWafRuleGroupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke UpdateWafRuleGroup with error: Operation validation and request error`, func() {
				wafRuleGroupsApiService, serviceErr := wafrulegroupsapiv1.NewWafRuleGroupsApiV1(&wafrulegroupsapiv1.WafRuleGroupsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ZoneID:        core.StringPtr(zoneID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(wafRuleGroupsApiService).ToNot(BeNil())

				// Construct an instance of the UpdateWafRuleGroupOptions model
				updateWafRuleGroupOptionsModel := new(wafrulegroupsapiv1.UpdateWafRuleGroupOptions)
				updateWafRuleGroupOptionsModel.PkgID = core.StringPtr("testString")
				updateWafRuleGroupOptionsModel.GroupID = core.StringPtr("testString")
				updateWafRuleGroupOptionsModel.Mode = core.StringPtr("on")
				updateWafRuleGroupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := wafRuleGroupsApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := wafRuleGroupsApiService.UpdateWafRuleGroup(updateWafRuleGroupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateWafRuleGroupOptions model with no property values
				updateWafRuleGroupOptionsModelNew := new(wafrulegroupsapiv1.UpdateWafRuleGroupOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = wafRuleGroupsApiService.UpdateWafRuleGroup(updateWafRuleGroupOptionsModelNew)
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
			zoneID := "testString"
			wafRuleGroupsApiService, _ := wafrulegroupsapiv1.NewWafRuleGroupsApiV1(&wafrulegroupsapiv1.WafRuleGroupsApiV1Options{
				URL:           "http://wafrulegroupsapiv1modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
				Crn:           core.StringPtr(crn),
				ZoneID:        core.StringPtr(zoneID),
			})
			It(`Invoke NewGetWafRuleGroupOptions successfully`, func() {
				// Construct an instance of the GetWafRuleGroupOptions model
				pkgID := "testString"
				groupID := "testString"
				getWafRuleGroupOptionsModel := wafRuleGroupsApiService.NewGetWafRuleGroupOptions(pkgID, groupID)
				getWafRuleGroupOptionsModel.SetPkgID("testString")
				getWafRuleGroupOptionsModel.SetGroupID("testString")
				getWafRuleGroupOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getWafRuleGroupOptionsModel).ToNot(BeNil())
				Expect(getWafRuleGroupOptionsModel.PkgID).To(Equal(core.StringPtr("testString")))
				Expect(getWafRuleGroupOptionsModel.GroupID).To(Equal(core.StringPtr("testString")))
				Expect(getWafRuleGroupOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListWafRuleGroupsOptions successfully`, func() {
				// Construct an instance of the ListWafRuleGroupsOptions model
				pkgID := "testString"
				listWafRuleGroupsOptionsModel := wafRuleGroupsApiService.NewListWafRuleGroupsOptions(pkgID)
				listWafRuleGroupsOptionsModel.SetPkgID("testString")
				listWafRuleGroupsOptionsModel.SetName("Wordpress-rules")
				listWafRuleGroupsOptionsModel.SetMode("true")
				listWafRuleGroupsOptionsModel.SetRulesCount("10")
				listWafRuleGroupsOptionsModel.SetPage(int64(1))
				listWafRuleGroupsOptionsModel.SetPerPage(int64(50))
				listWafRuleGroupsOptionsModel.SetOrder("status")
				listWafRuleGroupsOptionsModel.SetDirection("desc")
				listWafRuleGroupsOptionsModel.SetMatch("all")
				listWafRuleGroupsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listWafRuleGroupsOptionsModel).ToNot(BeNil())
				Expect(listWafRuleGroupsOptionsModel.PkgID).To(Equal(core.StringPtr("testString")))
				Expect(listWafRuleGroupsOptionsModel.Name).To(Equal(core.StringPtr("Wordpress-rules")))
				Expect(listWafRuleGroupsOptionsModel.Mode).To(Equal(core.StringPtr("true")))
				Expect(listWafRuleGroupsOptionsModel.RulesCount).To(Equal(core.StringPtr("10")))
				Expect(listWafRuleGroupsOptionsModel.Page).To(Equal(core.Int64Ptr(int64(1))))
				Expect(listWafRuleGroupsOptionsModel.PerPage).To(Equal(core.Int64Ptr(int64(50))))
				Expect(listWafRuleGroupsOptionsModel.Order).To(Equal(core.StringPtr("status")))
				Expect(listWafRuleGroupsOptionsModel.Direction).To(Equal(core.StringPtr("desc")))
				Expect(listWafRuleGroupsOptionsModel.Match).To(Equal(core.StringPtr("all")))
				Expect(listWafRuleGroupsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateWafRuleGroupOptions successfully`, func() {
				// Construct an instance of the UpdateWafRuleGroupOptions model
				pkgID := "testString"
				groupID := "testString"
				updateWafRuleGroupOptionsModel := wafRuleGroupsApiService.NewUpdateWafRuleGroupOptions(pkgID, groupID)
				updateWafRuleGroupOptionsModel.SetPkgID("testString")
				updateWafRuleGroupOptionsModel.SetGroupID("testString")
				updateWafRuleGroupOptionsModel.SetMode("on")
				updateWafRuleGroupOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateWafRuleGroupOptionsModel).ToNot(BeNil())
				Expect(updateWafRuleGroupOptionsModel.PkgID).To(Equal(core.StringPtr("testString")))
				Expect(updateWafRuleGroupOptionsModel.GroupID).To(Equal(core.StringPtr("testString")))
				Expect(updateWafRuleGroupOptionsModel.Mode).To(Equal(core.StringPtr("on")))
				Expect(updateWafRuleGroupOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
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
