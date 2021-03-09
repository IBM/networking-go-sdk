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

package useragentblockingrulesv1_test

import (
	"bytes"
	"context"
	"fmt"
	"github.com/IBM/go-sdk-core/v4/core"
	"github.com/IBM/networking-go-sdk/useragentblockingrulesv1"
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

var _ = Describe(`UserAgentBlockingRulesV1`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		It(`Instantiate service client`, func() {
			userAgentBlockingRulesService, serviceErr := useragentblockingrulesv1.NewUserAgentBlockingRulesV1(&useragentblockingrulesv1.UserAgentBlockingRulesV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
				Crn: core.StringPtr(crn),
				ZoneIdentifier: core.StringPtr(zoneIdentifier),
			})
			Expect(userAgentBlockingRulesService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			userAgentBlockingRulesService, serviceErr := useragentblockingrulesv1.NewUserAgentBlockingRulesV1(&useragentblockingrulesv1.UserAgentBlockingRulesV1Options{
				URL: "{BAD_URL_STRING",
				Crn: core.StringPtr(crn),
				ZoneIdentifier: core.StringPtr(zoneIdentifier),
			})
			Expect(userAgentBlockingRulesService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			userAgentBlockingRulesService, serviceErr := useragentblockingrulesv1.NewUserAgentBlockingRulesV1(&useragentblockingrulesv1.UserAgentBlockingRulesV1Options{
				URL: "https://useragentblockingrulesv1/api",
				Crn: core.StringPtr(crn),
				ZoneIdentifier: core.StringPtr(zoneIdentifier),
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(userAgentBlockingRulesService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Validation Error`, func() {
			userAgentBlockingRulesService, serviceErr := useragentblockingrulesv1.NewUserAgentBlockingRulesV1(&useragentblockingrulesv1.UserAgentBlockingRulesV1Options{})
			Expect(userAgentBlockingRulesService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"USER_AGENT_BLOCKING_RULES_URL": "https://useragentblockingrulesv1/api",
				"USER_AGENT_BLOCKING_RULES_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				userAgentBlockingRulesService, serviceErr := useragentblockingrulesv1.NewUserAgentBlockingRulesV1UsingExternalConfig(&useragentblockingrulesv1.UserAgentBlockingRulesV1Options{
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(userAgentBlockingRulesService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := userAgentBlockingRulesService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != userAgentBlockingRulesService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(userAgentBlockingRulesService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(userAgentBlockingRulesService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				userAgentBlockingRulesService, serviceErr := useragentblockingrulesv1.NewUserAgentBlockingRulesV1UsingExternalConfig(&useragentblockingrulesv1.UserAgentBlockingRulesV1Options{
					URL: "https://testService/api",
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(userAgentBlockingRulesService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(userAgentBlockingRulesService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := userAgentBlockingRulesService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != userAgentBlockingRulesService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(userAgentBlockingRulesService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(userAgentBlockingRulesService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				userAgentBlockingRulesService, serviceErr := useragentblockingrulesv1.NewUserAgentBlockingRulesV1UsingExternalConfig(&useragentblockingrulesv1.UserAgentBlockingRulesV1Options{
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				err := userAgentBlockingRulesService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(userAgentBlockingRulesService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(userAgentBlockingRulesService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := userAgentBlockingRulesService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != userAgentBlockingRulesService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(userAgentBlockingRulesService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(userAgentBlockingRulesService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"USER_AGENT_BLOCKING_RULES_URL": "https://useragentblockingrulesv1/api",
				"USER_AGENT_BLOCKING_RULES_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			userAgentBlockingRulesService, serviceErr := useragentblockingrulesv1.NewUserAgentBlockingRulesV1UsingExternalConfig(&useragentblockingrulesv1.UserAgentBlockingRulesV1Options{
				Crn: core.StringPtr(crn),
				ZoneIdentifier: core.StringPtr(zoneIdentifier),
			})

			It(`Instantiate service client with error`, func() {
				Expect(userAgentBlockingRulesService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"USER_AGENT_BLOCKING_RULES_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			userAgentBlockingRulesService, serviceErr := useragentblockingrulesv1.NewUserAgentBlockingRulesV1UsingExternalConfig(&useragentblockingrulesv1.UserAgentBlockingRulesV1Options{
				URL: "{BAD_URL_STRING",
				Crn: core.StringPtr(crn),
				ZoneIdentifier: core.StringPtr(zoneIdentifier),
			})

			It(`Instantiate service client with error`, func() {
				Expect(userAgentBlockingRulesService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = useragentblockingrulesv1.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`ListAllZoneUserAgentRules(listAllZoneUserAgentRulesOptions *ListAllZoneUserAgentRulesOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		listAllZoneUserAgentRulesPath := "/v1/testString/zones/testString/firewall/ua_rules"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listAllZoneUserAgentRulesPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["page"]).To(Equal([]string{fmt.Sprint(int64(38))}))

					Expect(req.URL.Query()["per_page"]).To(Equal([]string{fmt.Sprint(int64(5))}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListAllZoneUserAgentRules with error: Operation response processing error`, func() {
				userAgentBlockingRulesService, serviceErr := useragentblockingrulesv1.NewUserAgentBlockingRulesV1(&useragentblockingrulesv1.UserAgentBlockingRulesV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(userAgentBlockingRulesService).ToNot(BeNil())

				// Construct an instance of the ListAllZoneUserAgentRulesOptions model
				listAllZoneUserAgentRulesOptionsModel := new(useragentblockingrulesv1.ListAllZoneUserAgentRulesOptions)
				listAllZoneUserAgentRulesOptionsModel.Page = core.Int64Ptr(int64(38))
				listAllZoneUserAgentRulesOptionsModel.PerPage = core.Int64Ptr(int64(5))
				listAllZoneUserAgentRulesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := userAgentBlockingRulesService.ListAllZoneUserAgentRules(listAllZoneUserAgentRulesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				userAgentBlockingRulesService.EnableRetries(0, 0)
				result, response, operationErr = userAgentBlockingRulesService.ListAllZoneUserAgentRules(listAllZoneUserAgentRulesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ListAllZoneUserAgentRules(listAllZoneUserAgentRulesOptions *ListAllZoneUserAgentRulesOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		listAllZoneUserAgentRulesPath := "/v1/testString/zones/testString/firewall/ua_rules"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listAllZoneUserAgentRulesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["page"]).To(Equal([]string{fmt.Sprint(int64(38))}))

					Expect(req.URL.Query()["per_page"]).To(Equal([]string{fmt.Sprint(int64(5))}))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": [{"id": "92f17202ed8bd63d69a66b86a49a8f6b", "paused": true, "description": "Prevent access from abusive clients identified by this UserAgent to mitigate DDoS attack", "mode": "block", "configuration": {"target": "ua", "value": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_5) AppleWebKit/603.2.4 (KHTML, like Gecko) Version/10.1.1 Safari/603.2.4"}}], "result_info": {"page": 1, "per_page": 2, "count": 1, "total_count": 200}}`)
				}))
			})
			It(`Invoke ListAllZoneUserAgentRules successfully`, func() {
				userAgentBlockingRulesService, serviceErr := useragentblockingrulesv1.NewUserAgentBlockingRulesV1(&useragentblockingrulesv1.UserAgentBlockingRulesV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(userAgentBlockingRulesService).ToNot(BeNil())
				userAgentBlockingRulesService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := userAgentBlockingRulesService.ListAllZoneUserAgentRules(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListAllZoneUserAgentRulesOptions model
				listAllZoneUserAgentRulesOptionsModel := new(useragentblockingrulesv1.ListAllZoneUserAgentRulesOptions)
				listAllZoneUserAgentRulesOptionsModel.Page = core.Int64Ptr(int64(38))
				listAllZoneUserAgentRulesOptionsModel.PerPage = core.Int64Ptr(int64(5))
				listAllZoneUserAgentRulesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = userAgentBlockingRulesService.ListAllZoneUserAgentRules(listAllZoneUserAgentRulesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = userAgentBlockingRulesService.ListAllZoneUserAgentRulesWithContext(ctx, listAllZoneUserAgentRulesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				userAgentBlockingRulesService.DisableRetries()
				result, response, operationErr = userAgentBlockingRulesService.ListAllZoneUserAgentRules(listAllZoneUserAgentRulesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = userAgentBlockingRulesService.ListAllZoneUserAgentRulesWithContext(ctx, listAllZoneUserAgentRulesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke ListAllZoneUserAgentRules with error: Operation request error`, func() {
				userAgentBlockingRulesService, serviceErr := useragentblockingrulesv1.NewUserAgentBlockingRulesV1(&useragentblockingrulesv1.UserAgentBlockingRulesV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(userAgentBlockingRulesService).ToNot(BeNil())

				// Construct an instance of the ListAllZoneUserAgentRulesOptions model
				listAllZoneUserAgentRulesOptionsModel := new(useragentblockingrulesv1.ListAllZoneUserAgentRulesOptions)
				listAllZoneUserAgentRulesOptionsModel.Page = core.Int64Ptr(int64(38))
				listAllZoneUserAgentRulesOptionsModel.PerPage = core.Int64Ptr(int64(5))
				listAllZoneUserAgentRulesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := userAgentBlockingRulesService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := userAgentBlockingRulesService.ListAllZoneUserAgentRules(listAllZoneUserAgentRulesOptionsModel)
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
	Describe(`CreateZoneUserAgentRule(createZoneUserAgentRuleOptions *CreateZoneUserAgentRuleOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		createZoneUserAgentRulePath := "/v1/testString/zones/testString/firewall/ua_rules"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createZoneUserAgentRulePath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateZoneUserAgentRule with error: Operation response processing error`, func() {
				userAgentBlockingRulesService, serviceErr := useragentblockingrulesv1.NewUserAgentBlockingRulesV1(&useragentblockingrulesv1.UserAgentBlockingRulesV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(userAgentBlockingRulesService).ToNot(BeNil())

				// Construct an instance of the UseragentRuleInputConfiguration model
				useragentRuleInputConfigurationModel := new(useragentblockingrulesv1.UseragentRuleInputConfiguration)
				useragentRuleInputConfigurationModel.Target = core.StringPtr("ua")
				useragentRuleInputConfigurationModel.Value = core.StringPtr("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_5) AppleWebKit/603.2.4 (KHTML, like Gecko) Version/10.1.1 Safari/603.2.4")

				// Construct an instance of the CreateZoneUserAgentRuleOptions model
				createZoneUserAgentRuleOptionsModel := new(useragentblockingrulesv1.CreateZoneUserAgentRuleOptions)
				createZoneUserAgentRuleOptionsModel.Paused = core.BoolPtr(true)
				createZoneUserAgentRuleOptionsModel.Description = core.StringPtr("Prevent access from abusive clients identified by this UserAgent to mitigate DDoS attack")
				createZoneUserAgentRuleOptionsModel.Mode = core.StringPtr("block")
				createZoneUserAgentRuleOptionsModel.Configuration = useragentRuleInputConfigurationModel
				createZoneUserAgentRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := userAgentBlockingRulesService.CreateZoneUserAgentRule(createZoneUserAgentRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				userAgentBlockingRulesService.EnableRetries(0, 0)
				result, response, operationErr = userAgentBlockingRulesService.CreateZoneUserAgentRule(createZoneUserAgentRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`CreateZoneUserAgentRule(createZoneUserAgentRuleOptions *CreateZoneUserAgentRuleOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		createZoneUserAgentRulePath := "/v1/testString/zones/testString/firewall/ua_rules"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createZoneUserAgentRulePath))
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
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": "92f17202ed8bd63d69a66b86a49a8f6b", "paused": true, "description": "Prevent access from abusive clients identified by this UserAgent to mitigate DDoS attack", "mode": "block", "configuration": {"target": "ua", "value": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_5) AppleWebKit/603.2.4 (KHTML, like Gecko) Version/10.1.1 Safari/603.2.4"}}}`)
				}))
			})
			It(`Invoke CreateZoneUserAgentRule successfully`, func() {
				userAgentBlockingRulesService, serviceErr := useragentblockingrulesv1.NewUserAgentBlockingRulesV1(&useragentblockingrulesv1.UserAgentBlockingRulesV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(userAgentBlockingRulesService).ToNot(BeNil())
				userAgentBlockingRulesService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := userAgentBlockingRulesService.CreateZoneUserAgentRule(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UseragentRuleInputConfiguration model
				useragentRuleInputConfigurationModel := new(useragentblockingrulesv1.UseragentRuleInputConfiguration)
				useragentRuleInputConfigurationModel.Target = core.StringPtr("ua")
				useragentRuleInputConfigurationModel.Value = core.StringPtr("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_5) AppleWebKit/603.2.4 (KHTML, like Gecko) Version/10.1.1 Safari/603.2.4")

				// Construct an instance of the CreateZoneUserAgentRuleOptions model
				createZoneUserAgentRuleOptionsModel := new(useragentblockingrulesv1.CreateZoneUserAgentRuleOptions)
				createZoneUserAgentRuleOptionsModel.Paused = core.BoolPtr(true)
				createZoneUserAgentRuleOptionsModel.Description = core.StringPtr("Prevent access from abusive clients identified by this UserAgent to mitigate DDoS attack")
				createZoneUserAgentRuleOptionsModel.Mode = core.StringPtr("block")
				createZoneUserAgentRuleOptionsModel.Configuration = useragentRuleInputConfigurationModel
				createZoneUserAgentRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = userAgentBlockingRulesService.CreateZoneUserAgentRule(createZoneUserAgentRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = userAgentBlockingRulesService.CreateZoneUserAgentRuleWithContext(ctx, createZoneUserAgentRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				userAgentBlockingRulesService.DisableRetries()
				result, response, operationErr = userAgentBlockingRulesService.CreateZoneUserAgentRule(createZoneUserAgentRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = userAgentBlockingRulesService.CreateZoneUserAgentRuleWithContext(ctx, createZoneUserAgentRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke CreateZoneUserAgentRule with error: Operation request error`, func() {
				userAgentBlockingRulesService, serviceErr := useragentblockingrulesv1.NewUserAgentBlockingRulesV1(&useragentblockingrulesv1.UserAgentBlockingRulesV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(userAgentBlockingRulesService).ToNot(BeNil())

				// Construct an instance of the UseragentRuleInputConfiguration model
				useragentRuleInputConfigurationModel := new(useragentblockingrulesv1.UseragentRuleInputConfiguration)
				useragentRuleInputConfigurationModel.Target = core.StringPtr("ua")
				useragentRuleInputConfigurationModel.Value = core.StringPtr("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_5) AppleWebKit/603.2.4 (KHTML, like Gecko) Version/10.1.1 Safari/603.2.4")

				// Construct an instance of the CreateZoneUserAgentRuleOptions model
				createZoneUserAgentRuleOptionsModel := new(useragentblockingrulesv1.CreateZoneUserAgentRuleOptions)
				createZoneUserAgentRuleOptionsModel.Paused = core.BoolPtr(true)
				createZoneUserAgentRuleOptionsModel.Description = core.StringPtr("Prevent access from abusive clients identified by this UserAgent to mitigate DDoS attack")
				createZoneUserAgentRuleOptionsModel.Mode = core.StringPtr("block")
				createZoneUserAgentRuleOptionsModel.Configuration = useragentRuleInputConfigurationModel
				createZoneUserAgentRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := userAgentBlockingRulesService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := userAgentBlockingRulesService.CreateZoneUserAgentRule(createZoneUserAgentRuleOptionsModel)
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
	Describe(`DeleteZoneUserAgentRule(deleteZoneUserAgentRuleOptions *DeleteZoneUserAgentRuleOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		deleteZoneUserAgentRulePath := "/v1/testString/zones/testString/firewall/ua_rules/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteZoneUserAgentRulePath))
					Expect(req.Method).To(Equal("DELETE"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke DeleteZoneUserAgentRule with error: Operation response processing error`, func() {
				userAgentBlockingRulesService, serviceErr := useragentblockingrulesv1.NewUserAgentBlockingRulesV1(&useragentblockingrulesv1.UserAgentBlockingRulesV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(userAgentBlockingRulesService).ToNot(BeNil())

				// Construct an instance of the DeleteZoneUserAgentRuleOptions model
				deleteZoneUserAgentRuleOptionsModel := new(useragentblockingrulesv1.DeleteZoneUserAgentRuleOptions)
				deleteZoneUserAgentRuleOptionsModel.UseragentRuleIdentifier = core.StringPtr("testString")
				deleteZoneUserAgentRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := userAgentBlockingRulesService.DeleteZoneUserAgentRule(deleteZoneUserAgentRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				userAgentBlockingRulesService.EnableRetries(0, 0)
				result, response, operationErr = userAgentBlockingRulesService.DeleteZoneUserAgentRule(deleteZoneUserAgentRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`DeleteZoneUserAgentRule(deleteZoneUserAgentRuleOptions *DeleteZoneUserAgentRuleOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		deleteZoneUserAgentRulePath := "/v1/testString/zones/testString/firewall/ua_rules/testString"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteZoneUserAgentRulePath))
					Expect(req.Method).To(Equal("DELETE"))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": "f1aba936b94213e5b8dca0c0dbf1f9cc"}}`)
				}))
			})
			It(`Invoke DeleteZoneUserAgentRule successfully`, func() {
				userAgentBlockingRulesService, serviceErr := useragentblockingrulesv1.NewUserAgentBlockingRulesV1(&useragentblockingrulesv1.UserAgentBlockingRulesV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(userAgentBlockingRulesService).ToNot(BeNil())
				userAgentBlockingRulesService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := userAgentBlockingRulesService.DeleteZoneUserAgentRule(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DeleteZoneUserAgentRuleOptions model
				deleteZoneUserAgentRuleOptionsModel := new(useragentblockingrulesv1.DeleteZoneUserAgentRuleOptions)
				deleteZoneUserAgentRuleOptionsModel.UseragentRuleIdentifier = core.StringPtr("testString")
				deleteZoneUserAgentRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = userAgentBlockingRulesService.DeleteZoneUserAgentRule(deleteZoneUserAgentRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = userAgentBlockingRulesService.DeleteZoneUserAgentRuleWithContext(ctx, deleteZoneUserAgentRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				userAgentBlockingRulesService.DisableRetries()
				result, response, operationErr = userAgentBlockingRulesService.DeleteZoneUserAgentRule(deleteZoneUserAgentRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = userAgentBlockingRulesService.DeleteZoneUserAgentRuleWithContext(ctx, deleteZoneUserAgentRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke DeleteZoneUserAgentRule with error: Operation validation and request error`, func() {
				userAgentBlockingRulesService, serviceErr := useragentblockingrulesv1.NewUserAgentBlockingRulesV1(&useragentblockingrulesv1.UserAgentBlockingRulesV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(userAgentBlockingRulesService).ToNot(BeNil())

				// Construct an instance of the DeleteZoneUserAgentRuleOptions model
				deleteZoneUserAgentRuleOptionsModel := new(useragentblockingrulesv1.DeleteZoneUserAgentRuleOptions)
				deleteZoneUserAgentRuleOptionsModel.UseragentRuleIdentifier = core.StringPtr("testString")
				deleteZoneUserAgentRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := userAgentBlockingRulesService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := userAgentBlockingRulesService.DeleteZoneUserAgentRule(deleteZoneUserAgentRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the DeleteZoneUserAgentRuleOptions model with no property values
				deleteZoneUserAgentRuleOptionsModelNew := new(useragentblockingrulesv1.DeleteZoneUserAgentRuleOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = userAgentBlockingRulesService.DeleteZoneUserAgentRule(deleteZoneUserAgentRuleOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetUserAgentRule(getUserAgentRuleOptions *GetUserAgentRuleOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getUserAgentRulePath := "/v1/testString/zones/testString/firewall/ua_rules/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getUserAgentRulePath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetUserAgentRule with error: Operation response processing error`, func() {
				userAgentBlockingRulesService, serviceErr := useragentblockingrulesv1.NewUserAgentBlockingRulesV1(&useragentblockingrulesv1.UserAgentBlockingRulesV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(userAgentBlockingRulesService).ToNot(BeNil())

				// Construct an instance of the GetUserAgentRuleOptions model
				getUserAgentRuleOptionsModel := new(useragentblockingrulesv1.GetUserAgentRuleOptions)
				getUserAgentRuleOptionsModel.UseragentRuleIdentifier = core.StringPtr("testString")
				getUserAgentRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := userAgentBlockingRulesService.GetUserAgentRule(getUserAgentRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				userAgentBlockingRulesService.EnableRetries(0, 0)
				result, response, operationErr = userAgentBlockingRulesService.GetUserAgentRule(getUserAgentRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetUserAgentRule(getUserAgentRuleOptions *GetUserAgentRuleOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getUserAgentRulePath := "/v1/testString/zones/testString/firewall/ua_rules/testString"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getUserAgentRulePath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": "92f17202ed8bd63d69a66b86a49a8f6b", "paused": true, "description": "Prevent access from abusive clients identified by this UserAgent to mitigate DDoS attack", "mode": "block", "configuration": {"target": "ua", "value": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_5) AppleWebKit/603.2.4 (KHTML, like Gecko) Version/10.1.1 Safari/603.2.4"}}}`)
				}))
			})
			It(`Invoke GetUserAgentRule successfully`, func() {
				userAgentBlockingRulesService, serviceErr := useragentblockingrulesv1.NewUserAgentBlockingRulesV1(&useragentblockingrulesv1.UserAgentBlockingRulesV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(userAgentBlockingRulesService).ToNot(BeNil())
				userAgentBlockingRulesService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := userAgentBlockingRulesService.GetUserAgentRule(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetUserAgentRuleOptions model
				getUserAgentRuleOptionsModel := new(useragentblockingrulesv1.GetUserAgentRuleOptions)
				getUserAgentRuleOptionsModel.UseragentRuleIdentifier = core.StringPtr("testString")
				getUserAgentRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = userAgentBlockingRulesService.GetUserAgentRule(getUserAgentRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = userAgentBlockingRulesService.GetUserAgentRuleWithContext(ctx, getUserAgentRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				userAgentBlockingRulesService.DisableRetries()
				result, response, operationErr = userAgentBlockingRulesService.GetUserAgentRule(getUserAgentRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = userAgentBlockingRulesService.GetUserAgentRuleWithContext(ctx, getUserAgentRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke GetUserAgentRule with error: Operation validation and request error`, func() {
				userAgentBlockingRulesService, serviceErr := useragentblockingrulesv1.NewUserAgentBlockingRulesV1(&useragentblockingrulesv1.UserAgentBlockingRulesV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(userAgentBlockingRulesService).ToNot(BeNil())

				// Construct an instance of the GetUserAgentRuleOptions model
				getUserAgentRuleOptionsModel := new(useragentblockingrulesv1.GetUserAgentRuleOptions)
				getUserAgentRuleOptionsModel.UseragentRuleIdentifier = core.StringPtr("testString")
				getUserAgentRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := userAgentBlockingRulesService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := userAgentBlockingRulesService.GetUserAgentRule(getUserAgentRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetUserAgentRuleOptions model with no property values
				getUserAgentRuleOptionsModelNew := new(useragentblockingrulesv1.GetUserAgentRuleOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = userAgentBlockingRulesService.GetUserAgentRule(getUserAgentRuleOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateUserAgentRule(updateUserAgentRuleOptions *UpdateUserAgentRuleOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		updateUserAgentRulePath := "/v1/testString/zones/testString/firewall/ua_rules/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateUserAgentRulePath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateUserAgentRule with error: Operation response processing error`, func() {
				userAgentBlockingRulesService, serviceErr := useragentblockingrulesv1.NewUserAgentBlockingRulesV1(&useragentblockingrulesv1.UserAgentBlockingRulesV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(userAgentBlockingRulesService).ToNot(BeNil())

				// Construct an instance of the UseragentRuleInputConfiguration model
				useragentRuleInputConfigurationModel := new(useragentblockingrulesv1.UseragentRuleInputConfiguration)
				useragentRuleInputConfigurationModel.Target = core.StringPtr("ua")
				useragentRuleInputConfigurationModel.Value = core.StringPtr("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_5) AppleWebKit/603.2.4 (KHTML, like Gecko) Version/10.1.1 Safari/603.2.4")

				// Construct an instance of the UpdateUserAgentRuleOptions model
				updateUserAgentRuleOptionsModel := new(useragentblockingrulesv1.UpdateUserAgentRuleOptions)
				updateUserAgentRuleOptionsModel.UseragentRuleIdentifier = core.StringPtr("testString")
				updateUserAgentRuleOptionsModel.Paused = core.BoolPtr(true)
				updateUserAgentRuleOptionsModel.Description = core.StringPtr("Prevent access from abusive clients identified by this UserAgent to mitigate DDoS attack")
				updateUserAgentRuleOptionsModel.Mode = core.StringPtr("block")
				updateUserAgentRuleOptionsModel.Configuration = useragentRuleInputConfigurationModel
				updateUserAgentRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := userAgentBlockingRulesService.UpdateUserAgentRule(updateUserAgentRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				userAgentBlockingRulesService.EnableRetries(0, 0)
				result, response, operationErr = userAgentBlockingRulesService.UpdateUserAgentRule(updateUserAgentRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UpdateUserAgentRule(updateUserAgentRuleOptions *UpdateUserAgentRuleOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		updateUserAgentRulePath := "/v1/testString/zones/testString/firewall/ua_rules/testString"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateUserAgentRulePath))
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
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": "92f17202ed8bd63d69a66b86a49a8f6b", "paused": true, "description": "Prevent access from abusive clients identified by this UserAgent to mitigate DDoS attack", "mode": "block", "configuration": {"target": "ua", "value": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_5) AppleWebKit/603.2.4 (KHTML, like Gecko) Version/10.1.1 Safari/603.2.4"}}}`)
				}))
			})
			It(`Invoke UpdateUserAgentRule successfully`, func() {
				userAgentBlockingRulesService, serviceErr := useragentblockingrulesv1.NewUserAgentBlockingRulesV1(&useragentblockingrulesv1.UserAgentBlockingRulesV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(userAgentBlockingRulesService).ToNot(BeNil())
				userAgentBlockingRulesService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := userAgentBlockingRulesService.UpdateUserAgentRule(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UseragentRuleInputConfiguration model
				useragentRuleInputConfigurationModel := new(useragentblockingrulesv1.UseragentRuleInputConfiguration)
				useragentRuleInputConfigurationModel.Target = core.StringPtr("ua")
				useragentRuleInputConfigurationModel.Value = core.StringPtr("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_5) AppleWebKit/603.2.4 (KHTML, like Gecko) Version/10.1.1 Safari/603.2.4")

				// Construct an instance of the UpdateUserAgentRuleOptions model
				updateUserAgentRuleOptionsModel := new(useragentblockingrulesv1.UpdateUserAgentRuleOptions)
				updateUserAgentRuleOptionsModel.UseragentRuleIdentifier = core.StringPtr("testString")
				updateUserAgentRuleOptionsModel.Paused = core.BoolPtr(true)
				updateUserAgentRuleOptionsModel.Description = core.StringPtr("Prevent access from abusive clients identified by this UserAgent to mitigate DDoS attack")
				updateUserAgentRuleOptionsModel.Mode = core.StringPtr("block")
				updateUserAgentRuleOptionsModel.Configuration = useragentRuleInputConfigurationModel
				updateUserAgentRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = userAgentBlockingRulesService.UpdateUserAgentRule(updateUserAgentRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = userAgentBlockingRulesService.UpdateUserAgentRuleWithContext(ctx, updateUserAgentRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				userAgentBlockingRulesService.DisableRetries()
				result, response, operationErr = userAgentBlockingRulesService.UpdateUserAgentRule(updateUserAgentRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = userAgentBlockingRulesService.UpdateUserAgentRuleWithContext(ctx, updateUserAgentRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke UpdateUserAgentRule with error: Operation validation and request error`, func() {
				userAgentBlockingRulesService, serviceErr := useragentblockingrulesv1.NewUserAgentBlockingRulesV1(&useragentblockingrulesv1.UserAgentBlockingRulesV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(userAgentBlockingRulesService).ToNot(BeNil())

				// Construct an instance of the UseragentRuleInputConfiguration model
				useragentRuleInputConfigurationModel := new(useragentblockingrulesv1.UseragentRuleInputConfiguration)
				useragentRuleInputConfigurationModel.Target = core.StringPtr("ua")
				useragentRuleInputConfigurationModel.Value = core.StringPtr("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_5) AppleWebKit/603.2.4 (KHTML, like Gecko) Version/10.1.1 Safari/603.2.4")

				// Construct an instance of the UpdateUserAgentRuleOptions model
				updateUserAgentRuleOptionsModel := new(useragentblockingrulesv1.UpdateUserAgentRuleOptions)
				updateUserAgentRuleOptionsModel.UseragentRuleIdentifier = core.StringPtr("testString")
				updateUserAgentRuleOptionsModel.Paused = core.BoolPtr(true)
				updateUserAgentRuleOptionsModel.Description = core.StringPtr("Prevent access from abusive clients identified by this UserAgent to mitigate DDoS attack")
				updateUserAgentRuleOptionsModel.Mode = core.StringPtr("block")
				updateUserAgentRuleOptionsModel.Configuration = useragentRuleInputConfigurationModel
				updateUserAgentRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := userAgentBlockingRulesService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := userAgentBlockingRulesService.UpdateUserAgentRule(updateUserAgentRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateUserAgentRuleOptions model with no property values
				updateUserAgentRuleOptionsModelNew := new(useragentblockingrulesv1.UpdateUserAgentRuleOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = userAgentBlockingRulesService.UpdateUserAgentRule(updateUserAgentRuleOptionsModelNew)
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
			zoneIdentifier := "testString"
			userAgentBlockingRulesService, _ := useragentblockingrulesv1.NewUserAgentBlockingRulesV1(&useragentblockingrulesv1.UserAgentBlockingRulesV1Options{
				URL:           "http://useragentblockingrulesv1modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
				Crn: core.StringPtr(crn),
				ZoneIdentifier: core.StringPtr(zoneIdentifier),
			})
			It(`Invoke NewCreateZoneUserAgentRuleOptions successfully`, func() {
				// Construct an instance of the UseragentRuleInputConfiguration model
				useragentRuleInputConfigurationModel := new(useragentblockingrulesv1.UseragentRuleInputConfiguration)
				Expect(useragentRuleInputConfigurationModel).ToNot(BeNil())
				useragentRuleInputConfigurationModel.Target = core.StringPtr("ua")
				useragentRuleInputConfigurationModel.Value = core.StringPtr("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_5) AppleWebKit/603.2.4 (KHTML, like Gecko) Version/10.1.1 Safari/603.2.4")
				Expect(useragentRuleInputConfigurationModel.Target).To(Equal(core.StringPtr("ua")))
				Expect(useragentRuleInputConfigurationModel.Value).To(Equal(core.StringPtr("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_5) AppleWebKit/603.2.4 (KHTML, like Gecko) Version/10.1.1 Safari/603.2.4")))

				// Construct an instance of the CreateZoneUserAgentRuleOptions model
				createZoneUserAgentRuleOptionsModel := userAgentBlockingRulesService.NewCreateZoneUserAgentRuleOptions()
				createZoneUserAgentRuleOptionsModel.SetPaused(true)
				createZoneUserAgentRuleOptionsModel.SetDescription("Prevent access from abusive clients identified by this UserAgent to mitigate DDoS attack")
				createZoneUserAgentRuleOptionsModel.SetMode("block")
				createZoneUserAgentRuleOptionsModel.SetConfiguration(useragentRuleInputConfigurationModel)
				createZoneUserAgentRuleOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createZoneUserAgentRuleOptionsModel).ToNot(BeNil())
				Expect(createZoneUserAgentRuleOptionsModel.Paused).To(Equal(core.BoolPtr(true)))
				Expect(createZoneUserAgentRuleOptionsModel.Description).To(Equal(core.StringPtr("Prevent access from abusive clients identified by this UserAgent to mitigate DDoS attack")))
				Expect(createZoneUserAgentRuleOptionsModel.Mode).To(Equal(core.StringPtr("block")))
				Expect(createZoneUserAgentRuleOptionsModel.Configuration).To(Equal(useragentRuleInputConfigurationModel))
				Expect(createZoneUserAgentRuleOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteZoneUserAgentRuleOptions successfully`, func() {
				// Construct an instance of the DeleteZoneUserAgentRuleOptions model
				useragentRuleIdentifier := "testString"
				deleteZoneUserAgentRuleOptionsModel := userAgentBlockingRulesService.NewDeleteZoneUserAgentRuleOptions(useragentRuleIdentifier)
				deleteZoneUserAgentRuleOptionsModel.SetUseragentRuleIdentifier("testString")
				deleteZoneUserAgentRuleOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteZoneUserAgentRuleOptionsModel).ToNot(BeNil())
				Expect(deleteZoneUserAgentRuleOptionsModel.UseragentRuleIdentifier).To(Equal(core.StringPtr("testString")))
				Expect(deleteZoneUserAgentRuleOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetUserAgentRuleOptions successfully`, func() {
				// Construct an instance of the GetUserAgentRuleOptions model
				useragentRuleIdentifier := "testString"
				getUserAgentRuleOptionsModel := userAgentBlockingRulesService.NewGetUserAgentRuleOptions(useragentRuleIdentifier)
				getUserAgentRuleOptionsModel.SetUseragentRuleIdentifier("testString")
				getUserAgentRuleOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getUserAgentRuleOptionsModel).ToNot(BeNil())
				Expect(getUserAgentRuleOptionsModel.UseragentRuleIdentifier).To(Equal(core.StringPtr("testString")))
				Expect(getUserAgentRuleOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListAllZoneUserAgentRulesOptions successfully`, func() {
				// Construct an instance of the ListAllZoneUserAgentRulesOptions model
				listAllZoneUserAgentRulesOptionsModel := userAgentBlockingRulesService.NewListAllZoneUserAgentRulesOptions()
				listAllZoneUserAgentRulesOptionsModel.SetPage(int64(38))
				listAllZoneUserAgentRulesOptionsModel.SetPerPage(int64(5))
				listAllZoneUserAgentRulesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listAllZoneUserAgentRulesOptionsModel).ToNot(BeNil())
				Expect(listAllZoneUserAgentRulesOptionsModel.Page).To(Equal(core.Int64Ptr(int64(38))))
				Expect(listAllZoneUserAgentRulesOptionsModel.PerPage).To(Equal(core.Int64Ptr(int64(5))))
				Expect(listAllZoneUserAgentRulesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateUserAgentRuleOptions successfully`, func() {
				// Construct an instance of the UseragentRuleInputConfiguration model
				useragentRuleInputConfigurationModel := new(useragentblockingrulesv1.UseragentRuleInputConfiguration)
				Expect(useragentRuleInputConfigurationModel).ToNot(BeNil())
				useragentRuleInputConfigurationModel.Target = core.StringPtr("ua")
				useragentRuleInputConfigurationModel.Value = core.StringPtr("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_5) AppleWebKit/603.2.4 (KHTML, like Gecko) Version/10.1.1 Safari/603.2.4")
				Expect(useragentRuleInputConfigurationModel.Target).To(Equal(core.StringPtr("ua")))
				Expect(useragentRuleInputConfigurationModel.Value).To(Equal(core.StringPtr("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_5) AppleWebKit/603.2.4 (KHTML, like Gecko) Version/10.1.1 Safari/603.2.4")))

				// Construct an instance of the UpdateUserAgentRuleOptions model
				useragentRuleIdentifier := "testString"
				updateUserAgentRuleOptionsModel := userAgentBlockingRulesService.NewUpdateUserAgentRuleOptions(useragentRuleIdentifier)
				updateUserAgentRuleOptionsModel.SetUseragentRuleIdentifier("testString")
				updateUserAgentRuleOptionsModel.SetPaused(true)
				updateUserAgentRuleOptionsModel.SetDescription("Prevent access from abusive clients identified by this UserAgent to mitigate DDoS attack")
				updateUserAgentRuleOptionsModel.SetMode("block")
				updateUserAgentRuleOptionsModel.SetConfiguration(useragentRuleInputConfigurationModel)
				updateUserAgentRuleOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateUserAgentRuleOptionsModel).ToNot(BeNil())
				Expect(updateUserAgentRuleOptionsModel.UseragentRuleIdentifier).To(Equal(core.StringPtr("testString")))
				Expect(updateUserAgentRuleOptionsModel.Paused).To(Equal(core.BoolPtr(true)))
				Expect(updateUserAgentRuleOptionsModel.Description).To(Equal(core.StringPtr("Prevent access from abusive clients identified by this UserAgent to mitigate DDoS attack")))
				Expect(updateUserAgentRuleOptionsModel.Mode).To(Equal(core.StringPtr("block")))
				Expect(updateUserAgentRuleOptionsModel.Configuration).To(Equal(useragentRuleInputConfigurationModel))
				Expect(updateUserAgentRuleOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUseragentRuleInputConfiguration successfully`, func() {
				target := "ua"
				value := "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_5) AppleWebKit/603.2.4 (KHTML, like Gecko) Version/10.1.1 Safari/603.2.4"
				model, err := userAgentBlockingRulesService.NewUseragentRuleInputConfiguration(target, value)
				Expect(model).ToNot(BeNil())
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
