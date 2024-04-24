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
				 Authenticator: &core.NoAuthAuthenticator{},
				 Crn: core.StringPtr(crn),
				 ZoneIdentifier: core.StringPtr(zoneIdentifier),
			 })
			 Expect(rulesetsService).ToNot(BeNil())
			 Expect(serviceErr).To(BeNil())
		 })
		 It(`Instantiate service client with error: Invalid URL`, func() {
			 rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
				 URL: "{BAD_URL_STRING",
				 Crn: core.StringPtr(crn),
				 ZoneIdentifier: core.StringPtr(zoneIdentifier),
			 })
			 Expect(rulesetsService).To(BeNil())
			 Expect(serviceErr).ToNot(BeNil())
		 })
		 It(`Instantiate service client with error: Invalid Auth`, func() {
			 rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
				 URL: "https://rulesetsv1/api",
				 Crn: core.StringPtr(crn),
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
				 "RULESETS_URL": "https://rulesetsv1/api",
				 "RULESETS_AUTH_TYPE": "noauth",
			 }
 
			 It(`Create service client using external config successfully`, func() {
				 SetTestEnvironment(testEnvironment)
				 rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1UsingExternalConfig(&rulesetsv1.RulesetsV1Options{
					 Crn: core.StringPtr(crn),
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
					 URL: "https://testService/api",
					 Crn: core.StringPtr(crn),
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
					 Crn: core.StringPtr(crn),
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
				 "RULESETS_URL": "https://rulesetsv1/api",
				 "RULESETS_AUTH_TYPE": "someOtherAuth",
			 }
 
			 SetTestEnvironment(testEnvironment)
			 rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1UsingExternalConfig(&rulesetsv1.RulesetsV1Options{
				 Crn: core.StringPtr(crn),
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
				 "RULESETS_AUTH_TYPE":   "NOAuth",
			 }
 
			 SetTestEnvironment(testEnvironment)
			 rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1UsingExternalConfig(&rulesetsv1.RulesetsV1Options{
				 URL: "{BAD_URL_STRING",
				 Crn: core.StringPtr(crn),
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
	 Describe(`GetInstanceRulesets(getInstanceRulesetsOptions *GetInstanceRulesetsOptions) - Operation response error`, func() {
		 crn := "testString"
		 zoneIdentifier := "testString"
		 getInstanceRulesetsPath := "/v1/testString/rulesets"
		 Context(`Using mock server endpoint with invalid JSON response`, func() {
			 BeforeEach(func() {
				 testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					 defer GinkgoRecover()
 
					 // Verify the contents of the request
					 Expect(req.URL.EscapedPath()).To(Equal(getInstanceRulesetsPath))
					 Expect(req.Method).To(Equal("GET"))
					 res.Header().Set("Content-type", "application/json")
					 res.WriteHeader(200)
					 fmt.Fprint(res, `} this is not valid json {`)
				 }))
			 })
			 It(`Invoke GetInstanceRulesets with error: Operation response processing error`, func() {
				 rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
					 Crn: core.StringPtr(crn),
					 ZoneIdentifier: core.StringPtr(zoneIdentifier),
				 })
				 Expect(serviceErr).To(BeNil())
				 Expect(rulesetsService).ToNot(BeNil())
 
				 // Construct an instance of the GetInstanceRulesetsOptions model
				 getInstanceRulesetsOptionsModel := new(rulesetsv1.GetInstanceRulesetsOptions)
				 getInstanceRulesetsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				 // Expect response parsing to fail since we are receiving a text/plain response
				 result, response, operationErr := rulesetsService.GetInstanceRulesets(getInstanceRulesetsOptionsModel)
				 Expect(operationErr).ToNot(BeNil())
				 Expect(response).ToNot(BeNil())
				 Expect(result).To(BeNil())
 
				 // Enable retries and test again
				 rulesetsService.EnableRetries(0, 0)
				 result, response, operationErr = rulesetsService.GetInstanceRulesets(getInstanceRulesetsOptionsModel)
				 Expect(operationErr).ToNot(BeNil())
				 Expect(response).ToNot(BeNil())
				 Expect(result).To(BeNil())
			 })
			 AfterEach(func() {
				 testServer.Close()
			 })
		 })
	 })
	 Describe(`GetInstanceRulesets(getInstanceRulesetsOptions *GetInstanceRulesetsOptions)`, func() {
		 crn := "testString"
		 zoneIdentifier := "testString"
		 getInstanceRulesetsPath := "/v1/testString/rulesets"
		 Context(`Using mock server endpoint with timeout`, func() {
			 BeforeEach(func() {
				 testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					 defer GinkgoRecover()
 
					 // Verify the contents of the request
					 Expect(req.URL.EscapedPath()).To(Equal(getInstanceRulesetsPath))
					 Expect(req.Method).To(Equal("GET"))
 
					 // Sleep a short time to support a timeout test
					 time.Sleep(100 * time.Millisecond)
 
					 // Set mock response
					 res.Header().Set("Content-type", "application/json")
					 res.WriteHeader(200)
					 fmt.Fprintf(res, "%s", `{"success": true, "errors": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "messages": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "result": [{"description": "Custom instance ruleset", "id": "ID", "kind": "managed", "last_updated": "2000-01-01T00:00:00.000000Z", "name": "Name", "phase": "ddos_l4", "version": "1"}]}`)
				 }))
			 })
			 It(`Invoke GetInstanceRulesets successfully with retries`, func() {
				 rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
					 Crn: core.StringPtr(crn),
					 ZoneIdentifier: core.StringPtr(zoneIdentifier),
				 })
				 Expect(serviceErr).To(BeNil())
				 Expect(rulesetsService).ToNot(BeNil())
				 rulesetsService.EnableRetries(0, 0)
 
				 // Construct an instance of the GetInstanceRulesetsOptions model
				 getInstanceRulesetsOptionsModel := new(rulesetsv1.GetInstanceRulesetsOptions)
				 getInstanceRulesetsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
 
				 // Invoke operation with a Context to test a timeout error
				 ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				 defer cancelFunc()
				 _, _, operationErr := rulesetsService.GetInstanceRulesetsWithContext(ctx, getInstanceRulesetsOptionsModel)
				 Expect(operationErr).ToNot(BeNil())
				 Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
 
				 // Disable retries and test again
				 rulesetsService.DisableRetries()
				 result, response, operationErr := rulesetsService.GetInstanceRulesets(getInstanceRulesetsOptionsModel)
				 Expect(operationErr).To(BeNil())
				 Expect(response).ToNot(BeNil())
				 Expect(result).ToNot(BeNil())
 
				 // Re-test the timeout error with retries disabled
				 ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				 defer cancelFunc2()
				 _, _, operationErr = rulesetsService.GetInstanceRulesetsWithContext(ctx, getInstanceRulesetsOptionsModel)
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
					 Expect(req.URL.EscapedPath()).To(Equal(getInstanceRulesetsPath))
					 Expect(req.Method).To(Equal("GET"))
 
					 // Set mock response
					 res.Header().Set("Content-type", "application/json")
					 res.WriteHeader(200)
					 fmt.Fprintf(res, "%s", `{"success": true, "errors": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "messages": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "result": [{"description": "Custom instance ruleset", "id": "ID", "kind": "managed", "last_updated": "2000-01-01T00:00:00.000000Z", "name": "Name", "phase": "ddos_l4", "version": "1"}]}`)
				 }))
			 })
			 It(`Invoke GetInstanceRulesets successfully`, func() {
				 rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
					 Crn: core.StringPtr(crn),
					 ZoneIdentifier: core.StringPtr(zoneIdentifier),
				 })
				 Expect(serviceErr).To(BeNil())
				 Expect(rulesetsService).ToNot(BeNil())
 
				 // Invoke operation with nil options model (negative test)
				 result, response, operationErr := rulesetsService.GetInstanceRulesets(nil)
				 Expect(operationErr).NotTo(BeNil())
				 Expect(response).To(BeNil())
				 Expect(result).To(BeNil())
 
				 // Construct an instance of the GetInstanceRulesetsOptions model
				 getInstanceRulesetsOptionsModel := new(rulesetsv1.GetInstanceRulesetsOptions)
				 getInstanceRulesetsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
 
				 // Invoke operation with valid options model (positive test)
				 result, response, operationErr = rulesetsService.GetInstanceRulesets(getInstanceRulesetsOptionsModel)
				 Expect(operationErr).To(BeNil())
				 Expect(response).ToNot(BeNil())
				 Expect(result).ToNot(BeNil())
 
			 })
			 It(`Invoke GetInstanceRulesets with error: Operation request error`, func() {
				 rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
					 Crn: core.StringPtr(crn),
					 ZoneIdentifier: core.StringPtr(zoneIdentifier),
				 })
				 Expect(serviceErr).To(BeNil())
				 Expect(rulesetsService).ToNot(BeNil())
 
				 // Construct an instance of the GetInstanceRulesetsOptions model
				 getInstanceRulesetsOptionsModel := new(rulesetsv1.GetInstanceRulesetsOptions)
				 getInstanceRulesetsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				 // Invoke operation with empty URL (negative test)
				 err := rulesetsService.SetServiceURL("")
				 Expect(err).To(BeNil())
				 result, response, operationErr := rulesetsService.GetInstanceRulesets(getInstanceRulesetsOptionsModel)
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
			 It(`Invoke GetInstanceRulesets successfully`, func() {
				 rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
					 Crn: core.StringPtr(crn),
					 ZoneIdentifier: core.StringPtr(zoneIdentifier),
				 })
				 Expect(serviceErr).To(BeNil())
				 Expect(rulesetsService).ToNot(BeNil())
 
				 // Construct an instance of the GetInstanceRulesetsOptions model
				 getInstanceRulesetsOptionsModel := new(rulesetsv1.GetInstanceRulesetsOptions)
				 getInstanceRulesetsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
 
				 // Invoke operation
				 result, response, operationErr := rulesetsService.GetInstanceRulesets(getInstanceRulesetsOptionsModel)
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
	 Describe(`GetInstanceRuleset(getInstanceRulesetOptions *GetInstanceRulesetOptions) - Operation response error`, func() {
		 crn := "testString"
		 zoneIdentifier := "testString"
		 getInstanceRulesetPath := "/v1/testString/rulesets/testString"
		 Context(`Using mock server endpoint with invalid JSON response`, func() {
			 BeforeEach(func() {
				 testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					 defer GinkgoRecover()
 
					 // Verify the contents of the request
					 Expect(req.URL.EscapedPath()).To(Equal(getInstanceRulesetPath))
					 Expect(req.Method).To(Equal("GET"))
					 res.Header().Set("Content-type", "application/json")
					 res.WriteHeader(200)
					 fmt.Fprint(res, `} this is not valid json {`)
				 }))
			 })
			 It(`Invoke GetInstanceRuleset with error: Operation response processing error`, func() {
				 rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
					 Crn: core.StringPtr(crn),
					 ZoneIdentifier: core.StringPtr(zoneIdentifier),
				 })
				 Expect(serviceErr).To(BeNil())
				 Expect(rulesetsService).ToNot(BeNil())
 
				 // Construct an instance of the GetInstanceRulesetOptions model
				 getInstanceRulesetOptionsModel := new(rulesetsv1.GetInstanceRulesetOptions)
				 getInstanceRulesetOptionsModel.RulesetID = core.StringPtr("testString")
				 getInstanceRulesetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				 // Expect response parsing to fail since we are receiving a text/plain response
				 result, response, operationErr := rulesetsService.GetInstanceRuleset(getInstanceRulesetOptionsModel)
				 Expect(operationErr).ToNot(BeNil())
				 Expect(response).ToNot(BeNil())
				 Expect(result).To(BeNil())
 
				 // Enable retries and test again
				 rulesetsService.EnableRetries(0, 0)
				 result, response, operationErr = rulesetsService.GetInstanceRuleset(getInstanceRulesetOptionsModel)
				 Expect(operationErr).ToNot(BeNil())
				 Expect(response).ToNot(BeNil())
				 Expect(result).To(BeNil())
			 })
			 AfterEach(func() {
				 testServer.Close()
			 })
		 })
	 })
	 Describe(`GetInstanceRuleset(getInstanceRulesetOptions *GetInstanceRulesetOptions)`, func() {
		 crn := "testString"
		 zoneIdentifier := "testString"
		 getInstanceRulesetPath := "/v1/testString/rulesets/testString"
		 Context(`Using mock server endpoint with timeout`, func() {
			 BeforeEach(func() {
				 testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					 defer GinkgoRecover()
 
					 // Verify the contents of the request
					 Expect(req.URL.EscapedPath()).To(Equal(getInstanceRulesetPath))
					 Expect(req.Method).To(Equal("GET"))
 
					 // Sleep a short time to support a timeout test
					 time.Sleep(100 * time.Millisecond)
 
					 // Set mock response
					 res.Header().Set("Content-type", "application/json")
					 res.WriteHeader(200)
					 fmt.Fprintf(res, "%s", `{"success": true, "errors": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "messages": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "result": {"description": "Custom instance ruleset", "id": "ID", "kind": "managed", "last_updated": "2000-01-01T00:00:00.000000Z", "name": "Name", "phase": "ddos_l4", "version": "1", "rules": [{"id": "ID", "version": "Version", "action": "Action", "action_parameters": {"id": "ID", "overrides": {"action": "Action", "enabled": false, "sensitivity_level": "high", "rules": [{"id": "ID", "enabled": false, "action": "Action", "sensitivity_level": "high"}], "categories": [{"category": "Category", "enabled": false, "action": "Action"}]}, "version": "Version", "ruleset": "Ruleset", "rulesets": ["Rulesets"], "response": {"content": "{\"success\": false, \"error\": \"you have been blocked\"}", "content_type": "application/json", "status_code": 400}}, "categories": ["Categories"], "enabled": true, "description": "Description", "expression": "ip.src ne 1.1.1.1", "ref": "my_ref", "logging": {"enabled": true}, "last_updated": "2000-01-01T00:00:00.000000Z"}]}}`)
				 }))
			 })
			 It(`Invoke GetInstanceRuleset successfully with retries`, func() {
				 rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
					 Crn: core.StringPtr(crn),
					 ZoneIdentifier: core.StringPtr(zoneIdentifier),
				 })
				 Expect(serviceErr).To(BeNil())
				 Expect(rulesetsService).ToNot(BeNil())
				 rulesetsService.EnableRetries(0, 0)
 
				 // Construct an instance of the GetInstanceRulesetOptions model
				 getInstanceRulesetOptionsModel := new(rulesetsv1.GetInstanceRulesetOptions)
				 getInstanceRulesetOptionsModel.RulesetID = core.StringPtr("testString")
				 getInstanceRulesetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
 
				 // Invoke operation with a Context to test a timeout error
				 ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				 defer cancelFunc()
				 _, _, operationErr := rulesetsService.GetInstanceRulesetWithContext(ctx, getInstanceRulesetOptionsModel)
				 Expect(operationErr).ToNot(BeNil())
				 Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
 
				 // Disable retries and test again
				 rulesetsService.DisableRetries()
				 result, response, operationErr := rulesetsService.GetInstanceRuleset(getInstanceRulesetOptionsModel)
				 Expect(operationErr).To(BeNil())
				 Expect(response).ToNot(BeNil())
				 Expect(result).ToNot(BeNil())
 
				 // Re-test the timeout error with retries disabled
				 ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				 defer cancelFunc2()
				 _, _, operationErr = rulesetsService.GetInstanceRulesetWithContext(ctx, getInstanceRulesetOptionsModel)
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
					 Expect(req.URL.EscapedPath()).To(Equal(getInstanceRulesetPath))
					 Expect(req.Method).To(Equal("GET"))
 
					 // Set mock response
					 res.Header().Set("Content-type", "application/json")
					 res.WriteHeader(200)
					 fmt.Fprintf(res, "%s", `{"success": true, "errors": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "messages": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "result": {"description": "Custom instance ruleset", "id": "ID", "kind": "managed", "last_updated": "2000-01-01T00:00:00.000000Z", "name": "Name", "phase": "ddos_l4", "version": "1", "rules": [{"id": "ID", "version": "Version", "action": "Action", "action_parameters": {"id": "ID", "overrides": {"action": "Action", "enabled": false, "sensitivity_level": "high", "rules": [{"id": "ID", "enabled": false, "action": "Action", "sensitivity_level": "high"}], "categories": [{"category": "Category", "enabled": false, "action": "Action"}]}, "version": "Version", "ruleset": "Ruleset", "rulesets": ["Rulesets"], "response": {"content": "{\"success\": false, \"error\": \"you have been blocked\"}", "content_type": "application/json", "status_code": 400}}, "categories": ["Categories"], "enabled": true, "description": "Description", "expression": "ip.src ne 1.1.1.1", "ref": "my_ref", "logging": {"enabled": true}, "last_updated": "2000-01-01T00:00:00.000000Z"}]}}`)
				 }))
			 })
			 It(`Invoke GetInstanceRuleset successfully`, func() {
				 rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
					 Crn: core.StringPtr(crn),
					 ZoneIdentifier: core.StringPtr(zoneIdentifier),
				 })
				 Expect(serviceErr).To(BeNil())
				 Expect(rulesetsService).ToNot(BeNil())
 
				 // Invoke operation with nil options model (negative test)
				 result, response, operationErr := rulesetsService.GetInstanceRuleset(nil)
				 Expect(operationErr).NotTo(BeNil())
				 Expect(response).To(BeNil())
				 Expect(result).To(BeNil())
 
				 // Construct an instance of the GetInstanceRulesetOptions model
				 getInstanceRulesetOptionsModel := new(rulesetsv1.GetInstanceRulesetOptions)
				 getInstanceRulesetOptionsModel.RulesetID = core.StringPtr("testString")
				 getInstanceRulesetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
 
				 // Invoke operation with valid options model (positive test)
				 result, response, operationErr = rulesetsService.GetInstanceRuleset(getInstanceRulesetOptionsModel)
				 Expect(operationErr).To(BeNil())
				 Expect(response).ToNot(BeNil())
				 Expect(result).ToNot(BeNil())
 
			 })
			 It(`Invoke GetInstanceRuleset with error: Operation validation and request error`, func() {
				 rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
					 Crn: core.StringPtr(crn),
					 ZoneIdentifier: core.StringPtr(zoneIdentifier),
				 })
				 Expect(serviceErr).To(BeNil())
				 Expect(rulesetsService).ToNot(BeNil())
 
				 // Construct an instance of the GetInstanceRulesetOptions model
				 getInstanceRulesetOptionsModel := new(rulesetsv1.GetInstanceRulesetOptions)
				 getInstanceRulesetOptionsModel.RulesetID = core.StringPtr("testString")
				 getInstanceRulesetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				 // Invoke operation with empty URL (negative test)
				 err := rulesetsService.SetServiceURL("")
				 Expect(err).To(BeNil())
				 result, response, operationErr := rulesetsService.GetInstanceRuleset(getInstanceRulesetOptionsModel)
				 Expect(operationErr).ToNot(BeNil())
				 Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				 Expect(response).To(BeNil())
				 Expect(result).To(BeNil())
				 // Construct a second instance of the GetInstanceRulesetOptions model with no property values
				 getInstanceRulesetOptionsModelNew := new(rulesetsv1.GetInstanceRulesetOptions)
				 // Invoke operation with invalid model (negative test)
				 result, response, operationErr = rulesetsService.GetInstanceRuleset(getInstanceRulesetOptionsModelNew)
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
			 It(`Invoke GetInstanceRuleset successfully`, func() {
				 rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
					 Crn: core.StringPtr(crn),
					 ZoneIdentifier: core.StringPtr(zoneIdentifier),
				 })
				 Expect(serviceErr).To(BeNil())
				 Expect(rulesetsService).ToNot(BeNil())
 
				 // Construct an instance of the GetInstanceRulesetOptions model
				 getInstanceRulesetOptionsModel := new(rulesetsv1.GetInstanceRulesetOptions)
				 getInstanceRulesetOptionsModel.RulesetID = core.StringPtr("testString")
				 getInstanceRulesetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
 
				 // Invoke operation
				 result, response, operationErr := rulesetsService.GetInstanceRuleset(getInstanceRulesetOptionsModel)
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
	 Describe(`UpdateInstanceRuleset(updateInstanceRulesetOptions *UpdateInstanceRulesetOptions) - Operation response error`, func() {
		 crn := "testString"
		 zoneIdentifier := "testString"
		 updateInstanceRulesetPath := "/v1/testString/rulesets/testString"
		 Context(`Using mock server endpoint with invalid JSON response`, func() {
			 BeforeEach(func() {
				 testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					 defer GinkgoRecover()
 
					 // Verify the contents of the request
					 Expect(req.URL.EscapedPath()).To(Equal(updateInstanceRulesetPath))
					 Expect(req.Method).To(Equal("PUT"))
					 res.Header().Set("Content-type", "application/json")
					 res.WriteHeader(200)
					 fmt.Fprint(res, `} this is not valid json {`)
				 }))
			 })
			 It(`Invoke UpdateInstanceRuleset with error: Operation response processing error`, func() {
				 rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
					 Crn: core.StringPtr(crn),
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
				 actionParametersResponseModel.Content = core.StringPtr("{\"success\": false, \"error\": \"you have been blocked\"}")
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
 
				 // Construct an instance of the UpdateInstanceRulesetOptions model
				 updateInstanceRulesetOptionsModel := new(rulesetsv1.UpdateInstanceRulesetOptions)
				 updateInstanceRulesetOptionsModel.RulesetID = core.StringPtr("testString")
				 updateInstanceRulesetOptionsModel.Description = core.StringPtr("Custom instance ruleset")
				 updateInstanceRulesetOptionsModel.Kind = core.StringPtr("managed")
				 updateInstanceRulesetOptionsModel.Name = core.StringPtr("testString")
				 updateInstanceRulesetOptionsModel.Phase = core.StringPtr("ddos_l4")
				 updateInstanceRulesetOptionsModel.Rules = []rulesetsv1.RuleCreate{*ruleCreateModel}
				 updateInstanceRulesetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				 // Expect response parsing to fail since we are receiving a text/plain response
				 result, response, operationErr := rulesetsService.UpdateInstanceRuleset(updateInstanceRulesetOptionsModel)
				 Expect(operationErr).ToNot(BeNil())
				 Expect(response).ToNot(BeNil())
				 Expect(result).To(BeNil())
 
				 // Enable retries and test again
				 rulesetsService.EnableRetries(0, 0)
				 result, response, operationErr = rulesetsService.UpdateInstanceRuleset(updateInstanceRulesetOptionsModel)
				 Expect(operationErr).ToNot(BeNil())
				 Expect(response).ToNot(BeNil())
				 Expect(result).To(BeNil())
			 })
			 AfterEach(func() {
				 testServer.Close()
			 })
		 })
	 })
	 Describe(`UpdateInstanceRuleset(updateInstanceRulesetOptions *UpdateInstanceRulesetOptions)`, func() {
		 crn := "testString"
		 zoneIdentifier := "testString"
		 updateInstanceRulesetPath := "/v1/testString/rulesets/testString"
		 Context(`Using mock server endpoint with timeout`, func() {
			 BeforeEach(func() {
				 testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					 defer GinkgoRecover()
 
					 // Verify the contents of the request
					 Expect(req.URL.EscapedPath()).To(Equal(updateInstanceRulesetPath))
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
					 fmt.Fprintf(res, "%s", `{"success": true, "errors": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "messages": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "result": {"description": "Custom instance ruleset", "id": "ID", "kind": "managed", "last_updated": "2000-01-01T00:00:00.000000Z", "name": "Name", "phase": "ddos_l4", "version": "1", "rules": [{"id": "ID", "version": "Version", "action": "Action", "action_parameters": {"id": "ID", "overrides": {"action": "Action", "enabled": false, "sensitivity_level": "high", "rules": [{"id": "ID", "enabled": false, "action": "Action", "sensitivity_level": "high"}], "categories": [{"category": "Category", "enabled": false, "action": "Action"}]}, "version": "Version", "ruleset": "Ruleset", "rulesets": ["Rulesets"], "response": {"content": "{\"success\": false, \"error\": \"you have been blocked\"}", "content_type": "application/json", "status_code": 400}}, "categories": ["Categories"], "enabled": true, "description": "Description", "expression": "ip.src ne 1.1.1.1", "ref": "my_ref", "logging": {"enabled": true}, "last_updated": "2000-01-01T00:00:00.000000Z"}]}}`)
				 }))
			 })
			 It(`Invoke UpdateInstanceRuleset successfully with retries`, func() {
				 rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
					 Crn: core.StringPtr(crn),
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
				 actionParametersResponseModel.Content = core.StringPtr("{\"success\": false, \"error\": \"you have been blocked\"}")
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
 
				 // Construct an instance of the UpdateInstanceRulesetOptions model
				 updateInstanceRulesetOptionsModel := new(rulesetsv1.UpdateInstanceRulesetOptions)
				 updateInstanceRulesetOptionsModel.RulesetID = core.StringPtr("testString")
				 updateInstanceRulesetOptionsModel.Description = core.StringPtr("Custom instance ruleset")
				 updateInstanceRulesetOptionsModel.Kind = core.StringPtr("managed")
				 updateInstanceRulesetOptionsModel.Name = core.StringPtr("testString")
				 updateInstanceRulesetOptionsModel.Phase = core.StringPtr("ddos_l4")
				 updateInstanceRulesetOptionsModel.Rules = []rulesetsv1.RuleCreate{*ruleCreateModel}
				 updateInstanceRulesetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
 
				 // Invoke operation with a Context to test a timeout error
				 ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				 defer cancelFunc()
				 _, _, operationErr := rulesetsService.UpdateInstanceRulesetWithContext(ctx, updateInstanceRulesetOptionsModel)
				 Expect(operationErr).ToNot(BeNil())
				 Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
 
				 // Disable retries and test again
				 rulesetsService.DisableRetries()
				 result, response, operationErr := rulesetsService.UpdateInstanceRuleset(updateInstanceRulesetOptionsModel)
				 Expect(operationErr).To(BeNil())
				 Expect(response).ToNot(BeNil())
				 Expect(result).ToNot(BeNil())
 
				 // Re-test the timeout error with retries disabled
				 ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				 defer cancelFunc2()
				 _, _, operationErr = rulesetsService.UpdateInstanceRulesetWithContext(ctx, updateInstanceRulesetOptionsModel)
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
					 Expect(req.URL.EscapedPath()).To(Equal(updateInstanceRulesetPath))
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
					 fmt.Fprintf(res, "%s", `{"success": true, "errors": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "messages": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "result": {"description": "Custom instance ruleset", "id": "ID", "kind": "managed", "last_updated": "2000-01-01T00:00:00.000000Z", "name": "Name", "phase": "ddos_l4", "version": "1", "rules": [{"id": "ID", "version": "Version", "action": "Action", "action_parameters": {"id": "ID", "overrides": {"action": "Action", "enabled": false, "sensitivity_level": "high", "rules": [{"id": "ID", "enabled": false, "action": "Action", "sensitivity_level": "high"}], "categories": [{"category": "Category", "enabled": false, "action": "Action"}]}, "version": "Version", "ruleset": "Ruleset", "rulesets": ["Rulesets"], "response": {"content": "{\"success\": false, \"error\": \"you have been blocked\"}", "content_type": "application/json", "status_code": 400}}, "categories": ["Categories"], "enabled": true, "description": "Description", "expression": "ip.src ne 1.1.1.1", "ref": "my_ref", "logging": {"enabled": true}, "last_updated": "2000-01-01T00:00:00.000000Z"}]}}`)
				 }))
			 })
			 It(`Invoke UpdateInstanceRuleset successfully`, func() {
				 rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
					 Crn: core.StringPtr(crn),
					 ZoneIdentifier: core.StringPtr(zoneIdentifier),
				 })
				 Expect(serviceErr).To(BeNil())
				 Expect(rulesetsService).ToNot(BeNil())
 
				 // Invoke operation with nil options model (negative test)
				 result, response, operationErr := rulesetsService.UpdateInstanceRuleset(nil)
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
				 actionParametersResponseModel.Content = core.StringPtr("{\"success\": false, \"error\": \"you have been blocked\"}")
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
 
				 // Construct an instance of the UpdateInstanceRulesetOptions model
				 updateInstanceRulesetOptionsModel := new(rulesetsv1.UpdateInstanceRulesetOptions)
				 updateInstanceRulesetOptionsModel.RulesetID = core.StringPtr("testString")
				 updateInstanceRulesetOptionsModel.Description = core.StringPtr("Custom instance ruleset")
				 updateInstanceRulesetOptionsModel.Kind = core.StringPtr("managed")
				 updateInstanceRulesetOptionsModel.Name = core.StringPtr("testString")
				 updateInstanceRulesetOptionsModel.Phase = core.StringPtr("ddos_l4")
				 updateInstanceRulesetOptionsModel.Rules = []rulesetsv1.RuleCreate{*ruleCreateModel}
				 updateInstanceRulesetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
 
				 // Invoke operation with valid options model (positive test)
				 result, response, operationErr = rulesetsService.UpdateInstanceRuleset(updateInstanceRulesetOptionsModel)
				 Expect(operationErr).To(BeNil())
				 Expect(response).ToNot(BeNil())
				 Expect(result).ToNot(BeNil())
 
			 })
			 It(`Invoke UpdateInstanceRuleset with error: Operation validation and request error`, func() {
				 rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
					 Crn: core.StringPtr(crn),
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
				 actionParametersResponseModel.Content = core.StringPtr("{\"success\": false, \"error\": \"you have been blocked\"}")
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
 
				 // Construct an instance of the UpdateInstanceRulesetOptions model
				 updateInstanceRulesetOptionsModel := new(rulesetsv1.UpdateInstanceRulesetOptions)
				 updateInstanceRulesetOptionsModel.RulesetID = core.StringPtr("testString")
				 updateInstanceRulesetOptionsModel.Description = core.StringPtr("Custom instance ruleset")
				 updateInstanceRulesetOptionsModel.Kind = core.StringPtr("managed")
				 updateInstanceRulesetOptionsModel.Name = core.StringPtr("testString")
				 updateInstanceRulesetOptionsModel.Phase = core.StringPtr("ddos_l4")
				 updateInstanceRulesetOptionsModel.Rules = []rulesetsv1.RuleCreate{*ruleCreateModel}
				 updateInstanceRulesetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				 // Invoke operation with empty URL (negative test)
				 err := rulesetsService.SetServiceURL("")
				 Expect(err).To(BeNil())
				 result, response, operationErr := rulesetsService.UpdateInstanceRuleset(updateInstanceRulesetOptionsModel)
				 Expect(operationErr).ToNot(BeNil())
				 Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				 Expect(response).To(BeNil())
				 Expect(result).To(BeNil())
				 // Construct a second instance of the UpdateInstanceRulesetOptions model with no property values
				 updateInstanceRulesetOptionsModelNew := new(rulesetsv1.UpdateInstanceRulesetOptions)
				 // Invoke operation with invalid model (negative test)
				 result, response, operationErr = rulesetsService.UpdateInstanceRuleset(updateInstanceRulesetOptionsModelNew)
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
			 It(`Invoke UpdateInstanceRuleset successfully`, func() {
				 rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
					 Crn: core.StringPtr(crn),
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
				 actionParametersResponseModel.Content = core.StringPtr("{\"success\": false, \"error\": \"you have been blocked\"}")
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
 
				 // Construct an instance of the UpdateInstanceRulesetOptions model
				 updateInstanceRulesetOptionsModel := new(rulesetsv1.UpdateInstanceRulesetOptions)
				 updateInstanceRulesetOptionsModel.RulesetID = core.StringPtr("testString")
				 updateInstanceRulesetOptionsModel.Description = core.StringPtr("Custom instance ruleset")
				 updateInstanceRulesetOptionsModel.Kind = core.StringPtr("managed")
				 updateInstanceRulesetOptionsModel.Name = core.StringPtr("testString")
				 updateInstanceRulesetOptionsModel.Phase = core.StringPtr("ddos_l4")
				 updateInstanceRulesetOptionsModel.Rules = []rulesetsv1.RuleCreate{*ruleCreateModel}
				 updateInstanceRulesetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
 
				 // Invoke operation
				 result, response, operationErr := rulesetsService.UpdateInstanceRuleset(updateInstanceRulesetOptionsModel)
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
	 Describe(`DeleteInstanceRuleset(deleteInstanceRulesetOptions *DeleteInstanceRulesetOptions)`, func() {
		 crn := "testString"
		 zoneIdentifier := "testString"
		 deleteInstanceRulesetPath := "/v1/testString/rulesets/testString"
		 Context(`Using mock server endpoint`, func() {
			 BeforeEach(func() {
				 testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					 defer GinkgoRecover()
 
					 // Verify the contents of the request
					 Expect(req.URL.EscapedPath()).To(Equal(deleteInstanceRulesetPath))
					 Expect(req.Method).To(Equal("DELETE"))
 
					 res.WriteHeader(204)
				 }))
			 })
			 It(`Invoke DeleteInstanceRuleset successfully`, func() {
				 rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
					 Crn: core.StringPtr(crn),
					 ZoneIdentifier: core.StringPtr(zoneIdentifier),
				 })
				 Expect(serviceErr).To(BeNil())
				 Expect(rulesetsService).ToNot(BeNil())
 
				 // Invoke operation with nil options model (negative test)
				 response, operationErr := rulesetsService.DeleteInstanceRuleset(nil)
				 Expect(operationErr).NotTo(BeNil())
				 Expect(response).To(BeNil())
 
				 // Construct an instance of the DeleteInstanceRulesetOptions model
				 deleteInstanceRulesetOptionsModel := new(rulesetsv1.DeleteInstanceRulesetOptions)
				 deleteInstanceRulesetOptionsModel.RulesetID = core.StringPtr("testString")
				 deleteInstanceRulesetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
 
				 // Invoke operation with valid options model (positive test)
				 response, operationErr = rulesetsService.DeleteInstanceRuleset(deleteInstanceRulesetOptionsModel)
				 Expect(operationErr).To(BeNil())
				 Expect(response).ToNot(BeNil())
			 })
			 It(`Invoke DeleteInstanceRuleset with error: Operation validation and request error`, func() {
				 rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
					 Crn: core.StringPtr(crn),
					 ZoneIdentifier: core.StringPtr(zoneIdentifier),
				 })
				 Expect(serviceErr).To(BeNil())
				 Expect(rulesetsService).ToNot(BeNil())
 
				 // Construct an instance of the DeleteInstanceRulesetOptions model
				 deleteInstanceRulesetOptionsModel := new(rulesetsv1.DeleteInstanceRulesetOptions)
				 deleteInstanceRulesetOptionsModel.RulesetID = core.StringPtr("testString")
				 deleteInstanceRulesetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				 // Invoke operation with empty URL (negative test)
				 err := rulesetsService.SetServiceURL("")
				 Expect(err).To(BeNil())
				 response, operationErr := rulesetsService.DeleteInstanceRuleset(deleteInstanceRulesetOptionsModel)
				 Expect(operationErr).ToNot(BeNil())
				 Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				 Expect(response).To(BeNil())
				 // Construct a second instance of the DeleteInstanceRulesetOptions model with no property values
				 deleteInstanceRulesetOptionsModelNew := new(rulesetsv1.DeleteInstanceRulesetOptions)
				 // Invoke operation with invalid model (negative test)
				 response, operationErr = rulesetsService.DeleteInstanceRuleset(deleteInstanceRulesetOptionsModelNew)
				 Expect(operationErr).ToNot(BeNil())
				 Expect(response).To(BeNil())
			 })
			 AfterEach(func() {
				 testServer.Close()
			 })
		 })
	 })
	 Describe(`GetInstanceRulesetVersions(getInstanceRulesetVersionsOptions *GetInstanceRulesetVersionsOptions) - Operation response error`, func() {
		 crn := "testString"
		 zoneIdentifier := "testString"
		 getInstanceRulesetVersionsPath := "/v1/testString/rulesets/testString/versions"
		 Context(`Using mock server endpoint with invalid JSON response`, func() {
			 BeforeEach(func() {
				 testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					 defer GinkgoRecover()
 
					 // Verify the contents of the request
					 Expect(req.URL.EscapedPath()).To(Equal(getInstanceRulesetVersionsPath))
					 Expect(req.Method).To(Equal("GET"))
					 res.Header().Set("Content-type", "application/json")
					 res.WriteHeader(200)
					 fmt.Fprint(res, `} this is not valid json {`)
				 }))
			 })
			 It(`Invoke GetInstanceRulesetVersions with error: Operation response processing error`, func() {
				 rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
					 Crn: core.StringPtr(crn),
					 ZoneIdentifier: core.StringPtr(zoneIdentifier),
				 })
				 Expect(serviceErr).To(BeNil())
				 Expect(rulesetsService).ToNot(BeNil())
 
				 // Construct an instance of the GetInstanceRulesetVersionsOptions model
				 getInstanceRulesetVersionsOptionsModel := new(rulesetsv1.GetInstanceRulesetVersionsOptions)
				 getInstanceRulesetVersionsOptionsModel.RulesetID = core.StringPtr("testString")
				 getInstanceRulesetVersionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				 // Expect response parsing to fail since we are receiving a text/plain response
				 result, response, operationErr := rulesetsService.GetInstanceRulesetVersions(getInstanceRulesetVersionsOptionsModel)
				 Expect(operationErr).ToNot(BeNil())
				 Expect(response).ToNot(BeNil())
				 Expect(result).To(BeNil())
 
				 // Enable retries and test again
				 rulesetsService.EnableRetries(0, 0)
				 result, response, operationErr = rulesetsService.GetInstanceRulesetVersions(getInstanceRulesetVersionsOptionsModel)
				 Expect(operationErr).ToNot(BeNil())
				 Expect(response).ToNot(BeNil())
				 Expect(result).To(BeNil())
			 })
			 AfterEach(func() {
				 testServer.Close()
			 })
		 })
	 })
	 Describe(`GetInstanceRulesetVersions(getInstanceRulesetVersionsOptions *GetInstanceRulesetVersionsOptions)`, func() {
		 crn := "testString"
		 zoneIdentifier := "testString"
		 getInstanceRulesetVersionsPath := "/v1/testString/rulesets/testString/versions"
		 Context(`Using mock server endpoint with timeout`, func() {
			 BeforeEach(func() {
				 testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					 defer GinkgoRecover()
 
					 // Verify the contents of the request
					 Expect(req.URL.EscapedPath()).To(Equal(getInstanceRulesetVersionsPath))
					 Expect(req.Method).To(Equal("GET"))
 
					 // Sleep a short time to support a timeout test
					 time.Sleep(100 * time.Millisecond)
 
					 // Set mock response
					 res.Header().Set("Content-type", "application/json")
					 res.WriteHeader(200)
					 fmt.Fprintf(res, "%s", `{"success": true, "errors": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "messages": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "result": [{"description": "Custom instance ruleset", "id": "ID", "kind": "managed", "last_updated": "2000-01-01T00:00:00.000000Z", "name": "Name", "phase": "ddos_l4", "version": "1"}]}`)
				 }))
			 })
			 It(`Invoke GetInstanceRulesetVersions successfully with retries`, func() {
				 rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
					 Crn: core.StringPtr(crn),
					 ZoneIdentifier: core.StringPtr(zoneIdentifier),
				 })
				 Expect(serviceErr).To(BeNil())
				 Expect(rulesetsService).ToNot(BeNil())
				 rulesetsService.EnableRetries(0, 0)
 
				 // Construct an instance of the GetInstanceRulesetVersionsOptions model
				 getInstanceRulesetVersionsOptionsModel := new(rulesetsv1.GetInstanceRulesetVersionsOptions)
				 getInstanceRulesetVersionsOptionsModel.RulesetID = core.StringPtr("testString")
				 getInstanceRulesetVersionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
 
				 // Invoke operation with a Context to test a timeout error
				 ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				 defer cancelFunc()
				 _, _, operationErr := rulesetsService.GetInstanceRulesetVersionsWithContext(ctx, getInstanceRulesetVersionsOptionsModel)
				 Expect(operationErr).ToNot(BeNil())
				 Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
 
				 // Disable retries and test again
				 rulesetsService.DisableRetries()
				 result, response, operationErr := rulesetsService.GetInstanceRulesetVersions(getInstanceRulesetVersionsOptionsModel)
				 Expect(operationErr).To(BeNil())
				 Expect(response).ToNot(BeNil())
				 Expect(result).ToNot(BeNil())
 
				 // Re-test the timeout error with retries disabled
				 ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				 defer cancelFunc2()
				 _, _, operationErr = rulesetsService.GetInstanceRulesetVersionsWithContext(ctx, getInstanceRulesetVersionsOptionsModel)
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
					 Expect(req.URL.EscapedPath()).To(Equal(getInstanceRulesetVersionsPath))
					 Expect(req.Method).To(Equal("GET"))
 
					 // Set mock response
					 res.Header().Set("Content-type", "application/json")
					 res.WriteHeader(200)
					 fmt.Fprintf(res, "%s", `{"success": true, "errors": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "messages": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "result": [{"description": "Custom instance ruleset", "id": "ID", "kind": "managed", "last_updated": "2000-01-01T00:00:00.000000Z", "name": "Name", "phase": "ddos_l4", "version": "1"}]}`)
				 }))
			 })
			 It(`Invoke GetInstanceRulesetVersions successfully`, func() {
				 rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
					 Crn: core.StringPtr(crn),
					 ZoneIdentifier: core.StringPtr(zoneIdentifier),
				 })
				 Expect(serviceErr).To(BeNil())
				 Expect(rulesetsService).ToNot(BeNil())
 
				 // Invoke operation with nil options model (negative test)
				 result, response, operationErr := rulesetsService.GetInstanceRulesetVersions(nil)
				 Expect(operationErr).NotTo(BeNil())
				 Expect(response).To(BeNil())
				 Expect(result).To(BeNil())
 
				 // Construct an instance of the GetInstanceRulesetVersionsOptions model
				 getInstanceRulesetVersionsOptionsModel := new(rulesetsv1.GetInstanceRulesetVersionsOptions)
				 getInstanceRulesetVersionsOptionsModel.RulesetID = core.StringPtr("testString")
				 getInstanceRulesetVersionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
 
				 // Invoke operation with valid options model (positive test)
				 result, response, operationErr = rulesetsService.GetInstanceRulesetVersions(getInstanceRulesetVersionsOptionsModel)
				 Expect(operationErr).To(BeNil())
				 Expect(response).ToNot(BeNil())
				 Expect(result).ToNot(BeNil())
 
			 })
			 It(`Invoke GetInstanceRulesetVersions with error: Operation validation and request error`, func() {
				 rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
					 Crn: core.StringPtr(crn),
					 ZoneIdentifier: core.StringPtr(zoneIdentifier),
				 })
				 Expect(serviceErr).To(BeNil())
				 Expect(rulesetsService).ToNot(BeNil())
 
				 // Construct an instance of the GetInstanceRulesetVersionsOptions model
				 getInstanceRulesetVersionsOptionsModel := new(rulesetsv1.GetInstanceRulesetVersionsOptions)
				 getInstanceRulesetVersionsOptionsModel.RulesetID = core.StringPtr("testString")
				 getInstanceRulesetVersionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				 // Invoke operation with empty URL (negative test)
				 err := rulesetsService.SetServiceURL("")
				 Expect(err).To(BeNil())
				 result, response, operationErr := rulesetsService.GetInstanceRulesetVersions(getInstanceRulesetVersionsOptionsModel)
				 Expect(operationErr).ToNot(BeNil())
				 Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				 Expect(response).To(BeNil())
				 Expect(result).To(BeNil())
				 // Construct a second instance of the GetInstanceRulesetVersionsOptions model with no property values
				 getInstanceRulesetVersionsOptionsModelNew := new(rulesetsv1.GetInstanceRulesetVersionsOptions)
				 // Invoke operation with invalid model (negative test)
				 result, response, operationErr = rulesetsService.GetInstanceRulesetVersions(getInstanceRulesetVersionsOptionsModelNew)
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
			 It(`Invoke GetInstanceRulesetVersions successfully`, func() {
				 rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
					 Crn: core.StringPtr(crn),
					 ZoneIdentifier: core.StringPtr(zoneIdentifier),
				 })
				 Expect(serviceErr).To(BeNil())
				 Expect(rulesetsService).ToNot(BeNil())
 
				 // Construct an instance of the GetInstanceRulesetVersionsOptions model
				 getInstanceRulesetVersionsOptionsModel := new(rulesetsv1.GetInstanceRulesetVersionsOptions)
				 getInstanceRulesetVersionsOptionsModel.RulesetID = core.StringPtr("testString")
				 getInstanceRulesetVersionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
 
				 // Invoke operation
				 result, response, operationErr := rulesetsService.GetInstanceRulesetVersions(getInstanceRulesetVersionsOptionsModel)
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
	 Describe(`GetInstanceRulesetVersion(getInstanceRulesetVersionOptions *GetInstanceRulesetVersionOptions) - Operation response error`, func() {
		 crn := "testString"
		 zoneIdentifier := "testString"
		 getInstanceRulesetVersionPath := "/v1/testString/rulesets/testString/versions/1"
		 Context(`Using mock server endpoint with invalid JSON response`, func() {
			 BeforeEach(func() {
				 testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					 defer GinkgoRecover()
 
					 // Verify the contents of the request
					 Expect(req.URL.EscapedPath()).To(Equal(getInstanceRulesetVersionPath))
					 Expect(req.Method).To(Equal("GET"))
					 res.Header().Set("Content-type", "application/json")
					 res.WriteHeader(200)
					 fmt.Fprint(res, `} this is not valid json {`)
				 }))
			 })
			 It(`Invoke GetInstanceRulesetVersion with error: Operation response processing error`, func() {
				 rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
					 Crn: core.StringPtr(crn),
					 ZoneIdentifier: core.StringPtr(zoneIdentifier),
				 })
				 Expect(serviceErr).To(BeNil())
				 Expect(rulesetsService).ToNot(BeNil())
 
				 // Construct an instance of the GetInstanceRulesetVersionOptions model
				 getInstanceRulesetVersionOptionsModel := new(rulesetsv1.GetInstanceRulesetVersionOptions)
				 getInstanceRulesetVersionOptionsModel.RulesetID = core.StringPtr("testString")
				 getInstanceRulesetVersionOptionsModel.RulesetVersion = core.StringPtr("1")
				 getInstanceRulesetVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				 // Expect response parsing to fail since we are receiving a text/plain response
				 result, response, operationErr := rulesetsService.GetInstanceRulesetVersion(getInstanceRulesetVersionOptionsModel)
				 Expect(operationErr).ToNot(BeNil())
				 Expect(response).ToNot(BeNil())
				 Expect(result).To(BeNil())
 
				 // Enable retries and test again
				 rulesetsService.EnableRetries(0, 0)
				 result, response, operationErr = rulesetsService.GetInstanceRulesetVersion(getInstanceRulesetVersionOptionsModel)
				 Expect(operationErr).ToNot(BeNil())
				 Expect(response).ToNot(BeNil())
				 Expect(result).To(BeNil())
			 })
			 AfterEach(func() {
				 testServer.Close()
			 })
		 })
	 })
	 Describe(`GetInstanceRulesetVersion(getInstanceRulesetVersionOptions *GetInstanceRulesetVersionOptions)`, func() {
		 crn := "testString"
		 zoneIdentifier := "testString"
		 getInstanceRulesetVersionPath := "/v1/testString/rulesets/testString/versions/1"
		 Context(`Using mock server endpoint with timeout`, func() {
			 BeforeEach(func() {
				 testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					 defer GinkgoRecover()
 
					 // Verify the contents of the request
					 Expect(req.URL.EscapedPath()).To(Equal(getInstanceRulesetVersionPath))
					 Expect(req.Method).To(Equal("GET"))
 
					 // Sleep a short time to support a timeout test
					 time.Sleep(100 * time.Millisecond)
 
					 // Set mock response
					 res.Header().Set("Content-type", "application/json")
					 res.WriteHeader(200)
					 fmt.Fprintf(res, "%s", `{"success": true, "errors": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "messages": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "result": {"description": "Custom instance ruleset", "id": "ID", "kind": "managed", "last_updated": "2000-01-01T00:00:00.000000Z", "name": "Name", "phase": "ddos_l4", "version": "1", "rules": [{"id": "ID", "version": "Version", "action": "Action", "action_parameters": {"id": "ID", "overrides": {"action": "Action", "enabled": false, "sensitivity_level": "high", "rules": [{"id": "ID", "enabled": false, "action": "Action", "sensitivity_level": "high"}], "categories": [{"category": "Category", "enabled": false, "action": "Action"}]}, "version": "Version", "ruleset": "Ruleset", "rulesets": ["Rulesets"], "response": {"content": "{\"success\": false, \"error\": \"you have been blocked\"}", "content_type": "application/json", "status_code": 400}}, "categories": ["Categories"], "enabled": true, "description": "Description", "expression": "ip.src ne 1.1.1.1", "ref": "my_ref", "logging": {"enabled": true}, "last_updated": "2000-01-01T00:00:00.000000Z"}]}}`)
				 }))
			 })
			 It(`Invoke GetInstanceRulesetVersion successfully with retries`, func() {
				 rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
					 Crn: core.StringPtr(crn),
					 ZoneIdentifier: core.StringPtr(zoneIdentifier),
				 })
				 Expect(serviceErr).To(BeNil())
				 Expect(rulesetsService).ToNot(BeNil())
				 rulesetsService.EnableRetries(0, 0)
 
				 // Construct an instance of the GetInstanceRulesetVersionOptions model
				 getInstanceRulesetVersionOptionsModel := new(rulesetsv1.GetInstanceRulesetVersionOptions)
				 getInstanceRulesetVersionOptionsModel.RulesetID = core.StringPtr("testString")
				 getInstanceRulesetVersionOptionsModel.RulesetVersion = core.StringPtr("1")
				 getInstanceRulesetVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
 
				 // Invoke operation with a Context to test a timeout error
				 ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				 defer cancelFunc()
				 _, _, operationErr := rulesetsService.GetInstanceRulesetVersionWithContext(ctx, getInstanceRulesetVersionOptionsModel)
				 Expect(operationErr).ToNot(BeNil())
				 Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
 
				 // Disable retries and test again
				 rulesetsService.DisableRetries()
				 result, response, operationErr := rulesetsService.GetInstanceRulesetVersion(getInstanceRulesetVersionOptionsModel)
				 Expect(operationErr).To(BeNil())
				 Expect(response).ToNot(BeNil())
				 Expect(result).ToNot(BeNil())
 
				 // Re-test the timeout error with retries disabled
				 ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				 defer cancelFunc2()
				 _, _, operationErr = rulesetsService.GetInstanceRulesetVersionWithContext(ctx, getInstanceRulesetVersionOptionsModel)
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
					 Expect(req.URL.EscapedPath()).To(Equal(getInstanceRulesetVersionPath))
					 Expect(req.Method).To(Equal("GET"))
 
					 // Set mock response
					 res.Header().Set("Content-type", "application/json")
					 res.WriteHeader(200)
					 fmt.Fprintf(res, "%s", `{"success": true, "errors": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "messages": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "result": {"description": "Custom instance ruleset", "id": "ID", "kind": "managed", "last_updated": "2000-01-01T00:00:00.000000Z", "name": "Name", "phase": "ddos_l4", "version": "1", "rules": [{"id": "ID", "version": "Version", "action": "Action", "action_parameters": {"id": "ID", "overrides": {"action": "Action", "enabled": false, "sensitivity_level": "high", "rules": [{"id": "ID", "enabled": false, "action": "Action", "sensitivity_level": "high"}], "categories": [{"category": "Category", "enabled": false, "action": "Action"}]}, "version": "Version", "ruleset": "Ruleset", "rulesets": ["Rulesets"], "response": {"content": "{\"success\": false, \"error\": \"you have been blocked\"}", "content_type": "application/json", "status_code": 400}}, "categories": ["Categories"], "enabled": true, "description": "Description", "expression": "ip.src ne 1.1.1.1", "ref": "my_ref", "logging": {"enabled": true}, "last_updated": "2000-01-01T00:00:00.000000Z"}]}}`)
				 }))
			 })
			 It(`Invoke GetInstanceRulesetVersion successfully`, func() {
				 rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
					 Crn: core.StringPtr(crn),
					 ZoneIdentifier: core.StringPtr(zoneIdentifier),
				 })
				 Expect(serviceErr).To(BeNil())
				 Expect(rulesetsService).ToNot(BeNil())
 
				 // Invoke operation with nil options model (negative test)
				 result, response, operationErr := rulesetsService.GetInstanceRulesetVersion(nil)
				 Expect(operationErr).NotTo(BeNil())
				 Expect(response).To(BeNil())
				 Expect(result).To(BeNil())
 
				 // Construct an instance of the GetInstanceRulesetVersionOptions model
				 getInstanceRulesetVersionOptionsModel := new(rulesetsv1.GetInstanceRulesetVersionOptions)
				 getInstanceRulesetVersionOptionsModel.RulesetID = core.StringPtr("testString")
				 getInstanceRulesetVersionOptionsModel.RulesetVersion = core.StringPtr("1")
				 getInstanceRulesetVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
 
				 // Invoke operation with valid options model (positive test)
				 result, response, operationErr = rulesetsService.GetInstanceRulesetVersion(getInstanceRulesetVersionOptionsModel)
				 Expect(operationErr).To(BeNil())
				 Expect(response).ToNot(BeNil())
				 Expect(result).ToNot(BeNil())
 
			 })
			 It(`Invoke GetInstanceRulesetVersion with error: Operation validation and request error`, func() {
				 rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
					 Crn: core.StringPtr(crn),
					 ZoneIdentifier: core.StringPtr(zoneIdentifier),
				 })
				 Expect(serviceErr).To(BeNil())
				 Expect(rulesetsService).ToNot(BeNil())
 
				 // Construct an instance of the GetInstanceRulesetVersionOptions model
				 getInstanceRulesetVersionOptionsModel := new(rulesetsv1.GetInstanceRulesetVersionOptions)
				 getInstanceRulesetVersionOptionsModel.RulesetID = core.StringPtr("testString")
				 getInstanceRulesetVersionOptionsModel.RulesetVersion = core.StringPtr("1")
				 getInstanceRulesetVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				 // Invoke operation with empty URL (negative test)
				 err := rulesetsService.SetServiceURL("")
				 Expect(err).To(BeNil())
				 result, response, operationErr := rulesetsService.GetInstanceRulesetVersion(getInstanceRulesetVersionOptionsModel)
				 Expect(operationErr).ToNot(BeNil())
				 Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				 Expect(response).To(BeNil())
				 Expect(result).To(BeNil())
				 // Construct a second instance of the GetInstanceRulesetVersionOptions model with no property values
				 getInstanceRulesetVersionOptionsModelNew := new(rulesetsv1.GetInstanceRulesetVersionOptions)
				 // Invoke operation with invalid model (negative test)
				 result, response, operationErr = rulesetsService.GetInstanceRulesetVersion(getInstanceRulesetVersionOptionsModelNew)
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
			 It(`Invoke GetInstanceRulesetVersion successfully`, func() {
				 rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
					 Crn: core.StringPtr(crn),
					 ZoneIdentifier: core.StringPtr(zoneIdentifier),
				 })
				 Expect(serviceErr).To(BeNil())
				 Expect(rulesetsService).ToNot(BeNil())
 
				 // Construct an instance of the GetInstanceRulesetVersionOptions model
				 getInstanceRulesetVersionOptionsModel := new(rulesetsv1.GetInstanceRulesetVersionOptions)
				 getInstanceRulesetVersionOptionsModel.RulesetID = core.StringPtr("testString")
				 getInstanceRulesetVersionOptionsModel.RulesetVersion = core.StringPtr("1")
				 getInstanceRulesetVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
 
				 // Invoke operation
				 result, response, operationErr := rulesetsService.GetInstanceRulesetVersion(getInstanceRulesetVersionOptionsModel)
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
	 Describe(`DeleteInstanceRulesetVersion(deleteInstanceRulesetVersionOptions *DeleteInstanceRulesetVersionOptions)`, func() {
		 crn := "testString"
		 zoneIdentifier := "testString"
		 deleteInstanceRulesetVersionPath := "/v1/testString/rulesets/testString/versions/1"
		 Context(`Using mock server endpoint`, func() {
			 BeforeEach(func() {
				 testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					 defer GinkgoRecover()
 
					 // Verify the contents of the request
					 Expect(req.URL.EscapedPath()).To(Equal(deleteInstanceRulesetVersionPath))
					 Expect(req.Method).To(Equal("DELETE"))
 
					 res.WriteHeader(204)
				 }))
			 })
			 It(`Invoke DeleteInstanceRulesetVersion successfully`, func() {
				 rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
					 Crn: core.StringPtr(crn),
					 ZoneIdentifier: core.StringPtr(zoneIdentifier),
				 })
				 Expect(serviceErr).To(BeNil())
				 Expect(rulesetsService).ToNot(BeNil())
 
				 // Invoke operation with nil options model (negative test)
				 response, operationErr := rulesetsService.DeleteInstanceRulesetVersion(nil)
				 Expect(operationErr).NotTo(BeNil())
				 Expect(response).To(BeNil())
 
				 // Construct an instance of the DeleteInstanceRulesetVersionOptions model
				 deleteInstanceRulesetVersionOptionsModel := new(rulesetsv1.DeleteInstanceRulesetVersionOptions)
				 deleteInstanceRulesetVersionOptionsModel.RulesetID = core.StringPtr("testString")
				 deleteInstanceRulesetVersionOptionsModel.RulesetVersion = core.StringPtr("1")
				 deleteInstanceRulesetVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
 
				 // Invoke operation with valid options model (positive test)
				 response, operationErr = rulesetsService.DeleteInstanceRulesetVersion(deleteInstanceRulesetVersionOptionsModel)
				 Expect(operationErr).To(BeNil())
				 Expect(response).ToNot(BeNil())
			 })
			 It(`Invoke DeleteInstanceRulesetVersion with error: Operation validation and request error`, func() {
				 rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
					 Crn: core.StringPtr(crn),
					 ZoneIdentifier: core.StringPtr(zoneIdentifier),
				 })
				 Expect(serviceErr).To(BeNil())
				 Expect(rulesetsService).ToNot(BeNil())
 
				 // Construct an instance of the DeleteInstanceRulesetVersionOptions model
				 deleteInstanceRulesetVersionOptionsModel := new(rulesetsv1.DeleteInstanceRulesetVersionOptions)
				 deleteInstanceRulesetVersionOptionsModel.RulesetID = core.StringPtr("testString")
				 deleteInstanceRulesetVersionOptionsModel.RulesetVersion = core.StringPtr("1")
				 deleteInstanceRulesetVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				 // Invoke operation with empty URL (negative test)
				 err := rulesetsService.SetServiceURL("")
				 Expect(err).To(BeNil())
				 response, operationErr := rulesetsService.DeleteInstanceRulesetVersion(deleteInstanceRulesetVersionOptionsModel)
				 Expect(operationErr).ToNot(BeNil())
				 Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				 Expect(response).To(BeNil())
				 // Construct a second instance of the DeleteInstanceRulesetVersionOptions model with no property values
				 deleteInstanceRulesetVersionOptionsModelNew := new(rulesetsv1.DeleteInstanceRulesetVersionOptions)
				 // Invoke operation with invalid model (negative test)
				 response, operationErr = rulesetsService.DeleteInstanceRulesetVersion(deleteInstanceRulesetVersionOptionsModelNew)
				 Expect(operationErr).ToNot(BeNil())
				 Expect(response).To(BeNil())
			 })
			 AfterEach(func() {
				 testServer.Close()
			 })
		 })
	 })
	 Describe(`GetInstanceEntrypointRuleset(getInstanceEntrypointRulesetOptions *GetInstanceEntrypointRulesetOptions) - Operation response error`, func() {
		 crn := "testString"
		 zoneIdentifier := "testString"
		 getInstanceEntrypointRulesetPath := "/v1/testString/rulesets/phases/ddos_l4/entrypoint"
		 Context(`Using mock server endpoint with invalid JSON response`, func() {
			 BeforeEach(func() {
				 testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					 defer GinkgoRecover()
 
					 // Verify the contents of the request
					 Expect(req.URL.EscapedPath()).To(Equal(getInstanceEntrypointRulesetPath))
					 Expect(req.Method).To(Equal("GET"))
					 res.Header().Set("Content-type", "application/json")
					 res.WriteHeader(200)
					 fmt.Fprint(res, `} this is not valid json {`)
				 }))
			 })
			 It(`Invoke GetInstanceEntrypointRuleset with error: Operation response processing error`, func() {
				 rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
					 Crn: core.StringPtr(crn),
					 ZoneIdentifier: core.StringPtr(zoneIdentifier),
				 })
				 Expect(serviceErr).To(BeNil())
				 Expect(rulesetsService).ToNot(BeNil())
 
				 // Construct an instance of the GetInstanceEntrypointRulesetOptions model
				 getInstanceEntrypointRulesetOptionsModel := new(rulesetsv1.GetInstanceEntrypointRulesetOptions)
				 getInstanceEntrypointRulesetOptionsModel.RulesetPhase = core.StringPtr("ddos_l4")
				 getInstanceEntrypointRulesetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				 // Expect response parsing to fail since we are receiving a text/plain response
				 result, response, operationErr := rulesetsService.GetInstanceEntrypointRuleset(getInstanceEntrypointRulesetOptionsModel)
				 Expect(operationErr).ToNot(BeNil())
				 Expect(response).ToNot(BeNil())
				 Expect(result).To(BeNil())
 
				 // Enable retries and test again
				 rulesetsService.EnableRetries(0, 0)
				 result, response, operationErr = rulesetsService.GetInstanceEntrypointRuleset(getInstanceEntrypointRulesetOptionsModel)
				 Expect(operationErr).ToNot(BeNil())
				 Expect(response).ToNot(BeNil())
				 Expect(result).To(BeNil())
			 })
			 AfterEach(func() {
				 testServer.Close()
			 })
		 })
	 })
	 Describe(`GetInstanceEntrypointRuleset(getInstanceEntrypointRulesetOptions *GetInstanceEntrypointRulesetOptions)`, func() {
		 crn := "testString"
		 zoneIdentifier := "testString"
		 getInstanceEntrypointRulesetPath := "/v1/testString/rulesets/phases/ddos_l4/entrypoint"
		 Context(`Using mock server endpoint with timeout`, func() {
			 BeforeEach(func() {
				 testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					 defer GinkgoRecover()
 
					 // Verify the contents of the request
					 Expect(req.URL.EscapedPath()).To(Equal(getInstanceEntrypointRulesetPath))
					 Expect(req.Method).To(Equal("GET"))
 
					 // Sleep a short time to support a timeout test
					 time.Sleep(100 * time.Millisecond)
 
					 // Set mock response
					 res.Header().Set("Content-type", "application/json")
					 res.WriteHeader(200)
					 fmt.Fprintf(res, "%s", `{"success": true, "errors": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "messages": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "result": {"description": "Custom instance ruleset", "id": "ID", "kind": "managed", "last_updated": "2000-01-01T00:00:00.000000Z", "name": "Name", "phase": "ddos_l4", "version": "1", "rules": [{"id": "ID", "version": "Version", "action": "Action", "action_parameters": {"id": "ID", "overrides": {"action": "Action", "enabled": false, "sensitivity_level": "high", "rules": [{"id": "ID", "enabled": false, "action": "Action", "sensitivity_level": "high"}], "categories": [{"category": "Category", "enabled": false, "action": "Action"}]}, "version": "Version", "ruleset": "Ruleset", "rulesets": ["Rulesets"], "response": {"content": "{\"success\": false, \"error\": \"you have been blocked\"}", "content_type": "application/json", "status_code": 400}}, "categories": ["Categories"], "enabled": true, "description": "Description", "expression": "ip.src ne 1.1.1.1", "ref": "my_ref", "logging": {"enabled": true}, "last_updated": "2000-01-01T00:00:00.000000Z"}]}}`)
				 }))
			 })
			 It(`Invoke GetInstanceEntrypointRuleset successfully with retries`, func() {
				 rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
					 Crn: core.StringPtr(crn),
					 ZoneIdentifier: core.StringPtr(zoneIdentifier),
				 })
				 Expect(serviceErr).To(BeNil())
				 Expect(rulesetsService).ToNot(BeNil())
				 rulesetsService.EnableRetries(0, 0)
 
				 // Construct an instance of the GetInstanceEntrypointRulesetOptions model
				 getInstanceEntrypointRulesetOptionsModel := new(rulesetsv1.GetInstanceEntrypointRulesetOptions)
				 getInstanceEntrypointRulesetOptionsModel.RulesetPhase = core.StringPtr("ddos_l4")
				 getInstanceEntrypointRulesetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
 
				 // Invoke operation with a Context to test a timeout error
				 ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				 defer cancelFunc()
				 _, _, operationErr := rulesetsService.GetInstanceEntrypointRulesetWithContext(ctx, getInstanceEntrypointRulesetOptionsModel)
				 Expect(operationErr).ToNot(BeNil())
				 Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
 
				 // Disable retries and test again
				 rulesetsService.DisableRetries()
				 result, response, operationErr := rulesetsService.GetInstanceEntrypointRuleset(getInstanceEntrypointRulesetOptionsModel)
				 Expect(operationErr).To(BeNil())
				 Expect(response).ToNot(BeNil())
				 Expect(result).ToNot(BeNil())
 
				 // Re-test the timeout error with retries disabled
				 ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				 defer cancelFunc2()
				 _, _, operationErr = rulesetsService.GetInstanceEntrypointRulesetWithContext(ctx, getInstanceEntrypointRulesetOptionsModel)
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
					 Expect(req.URL.EscapedPath()).To(Equal(getInstanceEntrypointRulesetPath))
					 Expect(req.Method).To(Equal("GET"))
 
					 // Set mock response
					 res.Header().Set("Content-type", "application/json")
					 res.WriteHeader(200)
					 fmt.Fprintf(res, "%s", `{"success": true, "errors": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "messages": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "result": {"description": "Custom instance ruleset", "id": "ID", "kind": "managed", "last_updated": "2000-01-01T00:00:00.000000Z", "name": "Name", "phase": "ddos_l4", "version": "1", "rules": [{"id": "ID", "version": "Version", "action": "Action", "action_parameters": {"id": "ID", "overrides": {"action": "Action", "enabled": false, "sensitivity_level": "high", "rules": [{"id": "ID", "enabled": false, "action": "Action", "sensitivity_level": "high"}], "categories": [{"category": "Category", "enabled": false, "action": "Action"}]}, "version": "Version", "ruleset": "Ruleset", "rulesets": ["Rulesets"], "response": {"content": "{\"success\": false, \"error\": \"you have been blocked\"}", "content_type": "application/json", "status_code": 400}}, "categories": ["Categories"], "enabled": true, "description": "Description", "expression": "ip.src ne 1.1.1.1", "ref": "my_ref", "logging": {"enabled": true}, "last_updated": "2000-01-01T00:00:00.000000Z"}]}}`)
				 }))
			 })
			 It(`Invoke GetInstanceEntrypointRuleset successfully`, func() {
				 rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
					 Crn: core.StringPtr(crn),
					 ZoneIdentifier: core.StringPtr(zoneIdentifier),
				 })
				 Expect(serviceErr).To(BeNil())
				 Expect(rulesetsService).ToNot(BeNil())
 
				 // Invoke operation with nil options model (negative test)
				 result, response, operationErr := rulesetsService.GetInstanceEntrypointRuleset(nil)
				 Expect(operationErr).NotTo(BeNil())
				 Expect(response).To(BeNil())
				 Expect(result).To(BeNil())
 
				 // Construct an instance of the GetInstanceEntrypointRulesetOptions model
				 getInstanceEntrypointRulesetOptionsModel := new(rulesetsv1.GetInstanceEntrypointRulesetOptions)
				 getInstanceEntrypointRulesetOptionsModel.RulesetPhase = core.StringPtr("ddos_l4")
				 getInstanceEntrypointRulesetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
 
				 // Invoke operation with valid options model (positive test)
				 result, response, operationErr = rulesetsService.GetInstanceEntrypointRuleset(getInstanceEntrypointRulesetOptionsModel)
				 Expect(operationErr).To(BeNil())
				 Expect(response).ToNot(BeNil())
				 Expect(result).ToNot(BeNil())
 
			 })
			 It(`Invoke GetInstanceEntrypointRuleset with error: Operation validation and request error`, func() {
				 rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
					 Crn: core.StringPtr(crn),
					 ZoneIdentifier: core.StringPtr(zoneIdentifier),
				 })
				 Expect(serviceErr).To(BeNil())
				 Expect(rulesetsService).ToNot(BeNil())
 
				 // Construct an instance of the GetInstanceEntrypointRulesetOptions model
				 getInstanceEntrypointRulesetOptionsModel := new(rulesetsv1.GetInstanceEntrypointRulesetOptions)
				 getInstanceEntrypointRulesetOptionsModel.RulesetPhase = core.StringPtr("ddos_l4")
				 getInstanceEntrypointRulesetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				 // Invoke operation with empty URL (negative test)
				 err := rulesetsService.SetServiceURL("")
				 Expect(err).To(BeNil())
				 result, response, operationErr := rulesetsService.GetInstanceEntrypointRuleset(getInstanceEntrypointRulesetOptionsModel)
				 Expect(operationErr).ToNot(BeNil())
				 Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				 Expect(response).To(BeNil())
				 Expect(result).To(BeNil())
				 // Construct a second instance of the GetInstanceEntrypointRulesetOptions model with no property values
				 getInstanceEntrypointRulesetOptionsModelNew := new(rulesetsv1.GetInstanceEntrypointRulesetOptions)
				 // Invoke operation with invalid model (negative test)
				 result, response, operationErr = rulesetsService.GetInstanceEntrypointRuleset(getInstanceEntrypointRulesetOptionsModelNew)
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
			 It(`Invoke GetInstanceEntrypointRuleset successfully`, func() {
				 rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
					 Crn: core.StringPtr(crn),
					 ZoneIdentifier: core.StringPtr(zoneIdentifier),
				 })
				 Expect(serviceErr).To(BeNil())
				 Expect(rulesetsService).ToNot(BeNil())
 
				 // Construct an instance of the GetInstanceEntrypointRulesetOptions model
				 getInstanceEntrypointRulesetOptionsModel := new(rulesetsv1.GetInstanceEntrypointRulesetOptions)
				 getInstanceEntrypointRulesetOptionsModel.RulesetPhase = core.StringPtr("ddos_l4")
				 getInstanceEntrypointRulesetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
 
				 // Invoke operation
				 result, response, operationErr := rulesetsService.GetInstanceEntrypointRuleset(getInstanceEntrypointRulesetOptionsModel)
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
	 Describe(`UpdateInstanceEntrypointRuleset(updateInstanceEntrypointRulesetOptions *UpdateInstanceEntrypointRulesetOptions) - Operation response error`, func() {
		 crn := "testString"
		 zoneIdentifier := "testString"
		 updateInstanceEntrypointRulesetPath := "/v1/testString/rulesets/phases/ddos_l4/entrypoint"
		 Context(`Using mock server endpoint with invalid JSON response`, func() {
			 BeforeEach(func() {
				 testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					 defer GinkgoRecover()
 
					 // Verify the contents of the request
					 Expect(req.URL.EscapedPath()).To(Equal(updateInstanceEntrypointRulesetPath))
					 Expect(req.Method).To(Equal("PUT"))
					 res.Header().Set("Content-type", "application/json")
					 res.WriteHeader(200)
					 fmt.Fprint(res, `} this is not valid json {`)
				 }))
			 })
			 It(`Invoke UpdateInstanceEntrypointRuleset with error: Operation response processing error`, func() {
				 rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
					 Crn: core.StringPtr(crn),
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
				 actionParametersResponseModel.Content = core.StringPtr("{\"success\": false, \"error\": \"you have been blocked\"}")
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
 
				 // Construct an instance of the UpdateInstanceEntrypointRulesetOptions model
				 updateInstanceEntrypointRulesetOptionsModel := new(rulesetsv1.UpdateInstanceEntrypointRulesetOptions)
				 updateInstanceEntrypointRulesetOptionsModel.RulesetPhase = core.StringPtr("ddos_l4")
				 updateInstanceEntrypointRulesetOptionsModel.Description = core.StringPtr("Custom instance ruleset")
				 updateInstanceEntrypointRulesetOptionsModel.Kind = core.StringPtr("managed")
				 updateInstanceEntrypointRulesetOptionsModel.Name = core.StringPtr("testString")
				 updateInstanceEntrypointRulesetOptionsModel.Phase = core.StringPtr("ddos_l4")
				 updateInstanceEntrypointRulesetOptionsModel.Rules = []rulesetsv1.RuleCreate{*ruleCreateModel}
				 updateInstanceEntrypointRulesetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				 // Expect response parsing to fail since we are receiving a text/plain response
				 result, response, operationErr := rulesetsService.UpdateInstanceEntrypointRuleset(updateInstanceEntrypointRulesetOptionsModel)
				 Expect(operationErr).ToNot(BeNil())
				 Expect(response).ToNot(BeNil())
				 Expect(result).To(BeNil())
 
				 // Enable retries and test again
				 rulesetsService.EnableRetries(0, 0)
				 result, response, operationErr = rulesetsService.UpdateInstanceEntrypointRuleset(updateInstanceEntrypointRulesetOptionsModel)
				 Expect(operationErr).ToNot(BeNil())
				 Expect(response).ToNot(BeNil())
				 Expect(result).To(BeNil())
			 })
			 AfterEach(func() {
				 testServer.Close()
			 })
		 })
	 })
	 Describe(`UpdateInstanceEntrypointRuleset(updateInstanceEntrypointRulesetOptions *UpdateInstanceEntrypointRulesetOptions)`, func() {
		 crn := "testString"
		 zoneIdentifier := "testString"
		 updateInstanceEntrypointRulesetPath := "/v1/testString/rulesets/phases/ddos_l4/entrypoint"
		 Context(`Using mock server endpoint with timeout`, func() {
			 BeforeEach(func() {
				 testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					 defer GinkgoRecover()
 
					 // Verify the contents of the request
					 Expect(req.URL.EscapedPath()).To(Equal(updateInstanceEntrypointRulesetPath))
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
					 fmt.Fprintf(res, "%s", `{"success": true, "errors": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "messages": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "result": {"description": "Custom instance ruleset", "id": "ID", "kind": "managed", "last_updated": "2000-01-01T00:00:00.000000Z", "name": "Name", "phase": "ddos_l4", "version": "1", "rules": [{"id": "ID", "version": "Version", "action": "Action", "action_parameters": {"id": "ID", "overrides": {"action": "Action", "enabled": false, "sensitivity_level": "high", "rules": [{"id": "ID", "enabled": false, "action": "Action", "sensitivity_level": "high"}], "categories": [{"category": "Category", "enabled": false, "action": "Action"}]}, "version": "Version", "ruleset": "Ruleset", "rulesets": ["Rulesets"], "response": {"content": "{\"success\": false, \"error\": \"you have been blocked\"}", "content_type": "application/json", "status_code": 400}}, "categories": ["Categories"], "enabled": true, "description": "Description", "expression": "ip.src ne 1.1.1.1", "ref": "my_ref", "logging": {"enabled": true}, "last_updated": "2000-01-01T00:00:00.000000Z"}]}}`)
				 }))
			 })
			 It(`Invoke UpdateInstanceEntrypointRuleset successfully with retries`, func() {
				 rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
					 Crn: core.StringPtr(crn),
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
				 actionParametersResponseModel.Content = core.StringPtr("{\"success\": false, \"error\": \"you have been blocked\"}")
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
 
				 // Construct an instance of the UpdateInstanceEntrypointRulesetOptions model
				 updateInstanceEntrypointRulesetOptionsModel := new(rulesetsv1.UpdateInstanceEntrypointRulesetOptions)
				 updateInstanceEntrypointRulesetOptionsModel.RulesetPhase = core.StringPtr("ddos_l4")
				 updateInstanceEntrypointRulesetOptionsModel.Description = core.StringPtr("Custom instance ruleset")
				 updateInstanceEntrypointRulesetOptionsModel.Kind = core.StringPtr("managed")
				 updateInstanceEntrypointRulesetOptionsModel.Name = core.StringPtr("testString")
				 updateInstanceEntrypointRulesetOptionsModel.Phase = core.StringPtr("ddos_l4")
				 updateInstanceEntrypointRulesetOptionsModel.Rules = []rulesetsv1.RuleCreate{*ruleCreateModel}
				 updateInstanceEntrypointRulesetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
 
				 // Invoke operation with a Context to test a timeout error
				 ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				 defer cancelFunc()
				 _, _, operationErr := rulesetsService.UpdateInstanceEntrypointRulesetWithContext(ctx, updateInstanceEntrypointRulesetOptionsModel)
				 Expect(operationErr).ToNot(BeNil())
				 Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
 
				 // Disable retries and test again
				 rulesetsService.DisableRetries()
				 result, response, operationErr := rulesetsService.UpdateInstanceEntrypointRuleset(updateInstanceEntrypointRulesetOptionsModel)
				 Expect(operationErr).To(BeNil())
				 Expect(response).ToNot(BeNil())
				 Expect(result).ToNot(BeNil())
 
				 // Re-test the timeout error with retries disabled
				 ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				 defer cancelFunc2()
				 _, _, operationErr = rulesetsService.UpdateInstanceEntrypointRulesetWithContext(ctx, updateInstanceEntrypointRulesetOptionsModel)
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
					 Expect(req.URL.EscapedPath()).To(Equal(updateInstanceEntrypointRulesetPath))
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
					 fmt.Fprintf(res, "%s", `{"success": true, "errors": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "messages": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "result": {"description": "Custom instance ruleset", "id": "ID", "kind": "managed", "last_updated": "2000-01-01T00:00:00.000000Z", "name": "Name", "phase": "ddos_l4", "version": "1", "rules": [{"id": "ID", "version": "Version", "action": "Action", "action_parameters": {"id": "ID", "overrides": {"action": "Action", "enabled": false, "sensitivity_level": "high", "rules": [{"id": "ID", "enabled": false, "action": "Action", "sensitivity_level": "high"}], "categories": [{"category": "Category", "enabled": false, "action": "Action"}]}, "version": "Version", "ruleset": "Ruleset", "rulesets": ["Rulesets"], "response": {"content": "{\"success\": false, \"error\": \"you have been blocked\"}", "content_type": "application/json", "status_code": 400}}, "categories": ["Categories"], "enabled": true, "description": "Description", "expression": "ip.src ne 1.1.1.1", "ref": "my_ref", "logging": {"enabled": true}, "last_updated": "2000-01-01T00:00:00.000000Z"}]}}`)
				 }))
			 })
			 It(`Invoke UpdateInstanceEntrypointRuleset successfully`, func() {
				 rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
					 Crn: core.StringPtr(crn),
					 ZoneIdentifier: core.StringPtr(zoneIdentifier),
				 })
				 Expect(serviceErr).To(BeNil())
				 Expect(rulesetsService).ToNot(BeNil())
 
				 // Invoke operation with nil options model (negative test)
				 result, response, operationErr := rulesetsService.UpdateInstanceEntrypointRuleset(nil)
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
				 actionParametersResponseModel.Content = core.StringPtr("{\"success\": false, \"error\": \"you have been blocked\"}")
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
 
				 // Construct an instance of the UpdateInstanceEntrypointRulesetOptions model
				 updateInstanceEntrypointRulesetOptionsModel := new(rulesetsv1.UpdateInstanceEntrypointRulesetOptions)
				 updateInstanceEntrypointRulesetOptionsModel.RulesetPhase = core.StringPtr("ddos_l4")
				 updateInstanceEntrypointRulesetOptionsModel.Description = core.StringPtr("Custom instance ruleset")
				 updateInstanceEntrypointRulesetOptionsModel.Kind = core.StringPtr("managed")
				 updateInstanceEntrypointRulesetOptionsModel.Name = core.StringPtr("testString")
				 updateInstanceEntrypointRulesetOptionsModel.Phase = core.StringPtr("ddos_l4")
				 updateInstanceEntrypointRulesetOptionsModel.Rules = []rulesetsv1.RuleCreate{*ruleCreateModel}
				 updateInstanceEntrypointRulesetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
 
				 // Invoke operation with valid options model (positive test)
				 result, response, operationErr = rulesetsService.UpdateInstanceEntrypointRuleset(updateInstanceEntrypointRulesetOptionsModel)
				 Expect(operationErr).To(BeNil())
				 Expect(response).ToNot(BeNil())
				 Expect(result).ToNot(BeNil())
 
			 })
			 It(`Invoke UpdateInstanceEntrypointRuleset with error: Operation validation and request error`, func() {
				 rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
					 Crn: core.StringPtr(crn),
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
				 actionParametersResponseModel.Content = core.StringPtr("{\"success\": false, \"error\": \"you have been blocked\"}")
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
 
				 // Construct an instance of the UpdateInstanceEntrypointRulesetOptions model
				 updateInstanceEntrypointRulesetOptionsModel := new(rulesetsv1.UpdateInstanceEntrypointRulesetOptions)
				 updateInstanceEntrypointRulesetOptionsModel.RulesetPhase = core.StringPtr("ddos_l4")
				 updateInstanceEntrypointRulesetOptionsModel.Description = core.StringPtr("Custom instance ruleset")
				 updateInstanceEntrypointRulesetOptionsModel.Kind = core.StringPtr("managed")
				 updateInstanceEntrypointRulesetOptionsModel.Name = core.StringPtr("testString")
				 updateInstanceEntrypointRulesetOptionsModel.Phase = core.StringPtr("ddos_l4")
				 updateInstanceEntrypointRulesetOptionsModel.Rules = []rulesetsv1.RuleCreate{*ruleCreateModel}
				 updateInstanceEntrypointRulesetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				 // Invoke operation with empty URL (negative test)
				 err := rulesetsService.SetServiceURL("")
				 Expect(err).To(BeNil())
				 result, response, operationErr := rulesetsService.UpdateInstanceEntrypointRuleset(updateInstanceEntrypointRulesetOptionsModel)
				 Expect(operationErr).ToNot(BeNil())
				 Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				 Expect(response).To(BeNil())
				 Expect(result).To(BeNil())
				 // Construct a second instance of the UpdateInstanceEntrypointRulesetOptions model with no property values
				 updateInstanceEntrypointRulesetOptionsModelNew := new(rulesetsv1.UpdateInstanceEntrypointRulesetOptions)
				 // Invoke operation with invalid model (negative test)
				 result, response, operationErr = rulesetsService.UpdateInstanceEntrypointRuleset(updateInstanceEntrypointRulesetOptionsModelNew)
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
			 It(`Invoke UpdateInstanceEntrypointRuleset successfully`, func() {
				 rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
					 Crn: core.StringPtr(crn),
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
				 actionParametersResponseModel.Content = core.StringPtr("{\"success\": false, \"error\": \"you have been blocked\"}")
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
 
				 // Construct an instance of the UpdateInstanceEntrypointRulesetOptions model
				 updateInstanceEntrypointRulesetOptionsModel := new(rulesetsv1.UpdateInstanceEntrypointRulesetOptions)
				 updateInstanceEntrypointRulesetOptionsModel.RulesetPhase = core.StringPtr("ddos_l4")
				 updateInstanceEntrypointRulesetOptionsModel.Description = core.StringPtr("Custom instance ruleset")
				 updateInstanceEntrypointRulesetOptionsModel.Kind = core.StringPtr("managed")
				 updateInstanceEntrypointRulesetOptionsModel.Name = core.StringPtr("testString")
				 updateInstanceEntrypointRulesetOptionsModel.Phase = core.StringPtr("ddos_l4")
				 updateInstanceEntrypointRulesetOptionsModel.Rules = []rulesetsv1.RuleCreate{*ruleCreateModel}
				 updateInstanceEntrypointRulesetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
 
				 // Invoke operation
				 result, response, operationErr := rulesetsService.UpdateInstanceEntrypointRuleset(updateInstanceEntrypointRulesetOptionsModel)
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
	 Describe(`GetInstanceEntryPointRulesetVersions(getInstanceEntryPointRulesetVersionsOptions *GetInstanceEntryPointRulesetVersionsOptions) - Operation response error`, func() {
		 crn := "testString"
		 zoneIdentifier := "testString"
		 getInstanceEntryPointRulesetVersionsPath := "/v1/testString/rulesets/phases/ddos_l4/entrypoint/versions"
		 Context(`Using mock server endpoint with invalid JSON response`, func() {
			 BeforeEach(func() {
				 testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					 defer GinkgoRecover()
 
					 // Verify the contents of the request
					 Expect(req.URL.EscapedPath()).To(Equal(getInstanceEntryPointRulesetVersionsPath))
					 Expect(req.Method).To(Equal("GET"))
					 res.Header().Set("Content-type", "application/json")
					 res.WriteHeader(200)
					 fmt.Fprint(res, `} this is not valid json {`)
				 }))
			 })
			 It(`Invoke GetInstanceEntryPointRulesetVersions with error: Operation response processing error`, func() {
				 rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
					 Crn: core.StringPtr(crn),
					 ZoneIdentifier: core.StringPtr(zoneIdentifier),
				 })
				 Expect(serviceErr).To(BeNil())
				 Expect(rulesetsService).ToNot(BeNil())
 
				 // Construct an instance of the GetInstanceEntryPointRulesetVersionsOptions model
				 getInstanceEntryPointRulesetVersionsOptionsModel := new(rulesetsv1.GetInstanceEntryPointRulesetVersionsOptions)
				 getInstanceEntryPointRulesetVersionsOptionsModel.RulesetPhase = core.StringPtr("ddos_l4")
				 getInstanceEntryPointRulesetVersionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				 // Expect response parsing to fail since we are receiving a text/plain response
				 result, response, operationErr := rulesetsService.GetInstanceEntryPointRulesetVersions(getInstanceEntryPointRulesetVersionsOptionsModel)
				 Expect(operationErr).ToNot(BeNil())
				 Expect(response).ToNot(BeNil())
				 Expect(result).To(BeNil())
 
				 // Enable retries and test again
				 rulesetsService.EnableRetries(0, 0)
				 result, response, operationErr = rulesetsService.GetInstanceEntryPointRulesetVersions(getInstanceEntryPointRulesetVersionsOptionsModel)
				 Expect(operationErr).ToNot(BeNil())
				 Expect(response).ToNot(BeNil())
				 Expect(result).To(BeNil())
			 })
			 AfterEach(func() {
				 testServer.Close()
			 })
		 })
	 })
	 Describe(`GetInstanceEntryPointRulesetVersions(getInstanceEntryPointRulesetVersionsOptions *GetInstanceEntryPointRulesetVersionsOptions)`, func() {
		 crn := "testString"
		 zoneIdentifier := "testString"
		 getInstanceEntryPointRulesetVersionsPath := "/v1/testString/rulesets/phases/ddos_l4/entrypoint/versions"
		 Context(`Using mock server endpoint with timeout`, func() {
			 BeforeEach(func() {
				 testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					 defer GinkgoRecover()
 
					 // Verify the contents of the request
					 Expect(req.URL.EscapedPath()).To(Equal(getInstanceEntryPointRulesetVersionsPath))
					 Expect(req.Method).To(Equal("GET"))
 
					 // Sleep a short time to support a timeout test
					 time.Sleep(100 * time.Millisecond)
 
					 // Set mock response
					 res.Header().Set("Content-type", "application/json")
					 res.WriteHeader(200)
					 fmt.Fprintf(res, "%s", `{"success": true, "errors": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "messages": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "result": [{"description": "Custom instance ruleset", "id": "ID", "kind": "managed", "last_updated": "2000-01-01T00:00:00.000000Z", "name": "Name", "phase": "ddos_l4", "version": "1"}]}`)
				 }))
			 })
			 It(`Invoke GetInstanceEntryPointRulesetVersions successfully with retries`, func() {
				 rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
					 Crn: core.StringPtr(crn),
					 ZoneIdentifier: core.StringPtr(zoneIdentifier),
				 })
				 Expect(serviceErr).To(BeNil())
				 Expect(rulesetsService).ToNot(BeNil())
				 rulesetsService.EnableRetries(0, 0)
 
				 // Construct an instance of the GetInstanceEntryPointRulesetVersionsOptions model
				 getInstanceEntryPointRulesetVersionsOptionsModel := new(rulesetsv1.GetInstanceEntryPointRulesetVersionsOptions)
				 getInstanceEntryPointRulesetVersionsOptionsModel.RulesetPhase = core.StringPtr("ddos_l4")
				 getInstanceEntryPointRulesetVersionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
 
				 // Invoke operation with a Context to test a timeout error
				 ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				 defer cancelFunc()
				 _, _, operationErr := rulesetsService.GetInstanceEntryPointRulesetVersionsWithContext(ctx, getInstanceEntryPointRulesetVersionsOptionsModel)
				 Expect(operationErr).ToNot(BeNil())
				 Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
 
				 // Disable retries and test again
				 rulesetsService.DisableRetries()
				 result, response, operationErr := rulesetsService.GetInstanceEntryPointRulesetVersions(getInstanceEntryPointRulesetVersionsOptionsModel)
				 Expect(operationErr).To(BeNil())
				 Expect(response).ToNot(BeNil())
				 Expect(result).ToNot(BeNil())
 
				 // Re-test the timeout error with retries disabled
				 ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				 defer cancelFunc2()
				 _, _, operationErr = rulesetsService.GetInstanceEntryPointRulesetVersionsWithContext(ctx, getInstanceEntryPointRulesetVersionsOptionsModel)
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
					 Expect(req.URL.EscapedPath()).To(Equal(getInstanceEntryPointRulesetVersionsPath))
					 Expect(req.Method).To(Equal("GET"))
 
					 // Set mock response
					 res.Header().Set("Content-type", "application/json")
					 res.WriteHeader(200)
					 fmt.Fprintf(res, "%s", `{"success": true, "errors": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "messages": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "result": [{"description": "Custom instance ruleset", "id": "ID", "kind": "managed", "last_updated": "2000-01-01T00:00:00.000000Z", "name": "Name", "phase": "ddos_l4", "version": "1"}]}`)
				 }))
			 })
			 It(`Invoke GetInstanceEntryPointRulesetVersions successfully`, func() {
				 rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
					 Crn: core.StringPtr(crn),
					 ZoneIdentifier: core.StringPtr(zoneIdentifier),
				 })
				 Expect(serviceErr).To(BeNil())
				 Expect(rulesetsService).ToNot(BeNil())
 
				 // Invoke operation with nil options model (negative test)
				 result, response, operationErr := rulesetsService.GetInstanceEntryPointRulesetVersions(nil)
				 Expect(operationErr).NotTo(BeNil())
				 Expect(response).To(BeNil())
				 Expect(result).To(BeNil())
 
				 // Construct an instance of the GetInstanceEntryPointRulesetVersionsOptions model
				 getInstanceEntryPointRulesetVersionsOptionsModel := new(rulesetsv1.GetInstanceEntryPointRulesetVersionsOptions)
				 getInstanceEntryPointRulesetVersionsOptionsModel.RulesetPhase = core.StringPtr("ddos_l4")
				 getInstanceEntryPointRulesetVersionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
 
				 // Invoke operation with valid options model (positive test)
				 result, response, operationErr = rulesetsService.GetInstanceEntryPointRulesetVersions(getInstanceEntryPointRulesetVersionsOptionsModel)
				 Expect(operationErr).To(BeNil())
				 Expect(response).ToNot(BeNil())
				 Expect(result).ToNot(BeNil())
 
			 })
			 It(`Invoke GetInstanceEntryPointRulesetVersions with error: Operation validation and request error`, func() {
				 rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
					 Crn: core.StringPtr(crn),
					 ZoneIdentifier: core.StringPtr(zoneIdentifier),
				 })
				 Expect(serviceErr).To(BeNil())
				 Expect(rulesetsService).ToNot(BeNil())
 
				 // Construct an instance of the GetInstanceEntryPointRulesetVersionsOptions model
				 getInstanceEntryPointRulesetVersionsOptionsModel := new(rulesetsv1.GetInstanceEntryPointRulesetVersionsOptions)
				 getInstanceEntryPointRulesetVersionsOptionsModel.RulesetPhase = core.StringPtr("ddos_l4")
				 getInstanceEntryPointRulesetVersionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				 // Invoke operation with empty URL (negative test)
				 err := rulesetsService.SetServiceURL("")
				 Expect(err).To(BeNil())
				 result, response, operationErr := rulesetsService.GetInstanceEntryPointRulesetVersions(getInstanceEntryPointRulesetVersionsOptionsModel)
				 Expect(operationErr).ToNot(BeNil())
				 Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				 Expect(response).To(BeNil())
				 Expect(result).To(BeNil())
				 // Construct a second instance of the GetInstanceEntryPointRulesetVersionsOptions model with no property values
				 getInstanceEntryPointRulesetVersionsOptionsModelNew := new(rulesetsv1.GetInstanceEntryPointRulesetVersionsOptions)
				 // Invoke operation with invalid model (negative test)
				 result, response, operationErr = rulesetsService.GetInstanceEntryPointRulesetVersions(getInstanceEntryPointRulesetVersionsOptionsModelNew)
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
			 It(`Invoke GetInstanceEntryPointRulesetVersions successfully`, func() {
				 rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
					 Crn: core.StringPtr(crn),
					 ZoneIdentifier: core.StringPtr(zoneIdentifier),
				 })
				 Expect(serviceErr).To(BeNil())
				 Expect(rulesetsService).ToNot(BeNil())
 
				 // Construct an instance of the GetInstanceEntryPointRulesetVersionsOptions model
				 getInstanceEntryPointRulesetVersionsOptionsModel := new(rulesetsv1.GetInstanceEntryPointRulesetVersionsOptions)
				 getInstanceEntryPointRulesetVersionsOptionsModel.RulesetPhase = core.StringPtr("ddos_l4")
				 getInstanceEntryPointRulesetVersionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
 
				 // Invoke operation
				 result, response, operationErr := rulesetsService.GetInstanceEntryPointRulesetVersions(getInstanceEntryPointRulesetVersionsOptionsModel)
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
	 Describe(`GetInstanceEntryPointRulesetVersion(getInstanceEntryPointRulesetVersionOptions *GetInstanceEntryPointRulesetVersionOptions) - Operation response error`, func() {
		 crn := "testString"
		 zoneIdentifier := "testString"
		 getInstanceEntryPointRulesetVersionPath := "/v1/testString/rulesets/phases/ddos_l4/entrypoint/versions/1"
		 Context(`Using mock server endpoint with invalid JSON response`, func() {
			 BeforeEach(func() {
				 testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					 defer GinkgoRecover()
 
					 // Verify the contents of the request
					 Expect(req.URL.EscapedPath()).To(Equal(getInstanceEntryPointRulesetVersionPath))
					 Expect(req.Method).To(Equal("GET"))
					 res.Header().Set("Content-type", "application/json")
					 res.WriteHeader(200)
					 fmt.Fprint(res, `} this is not valid json {`)
				 }))
			 })
			 It(`Invoke GetInstanceEntryPointRulesetVersion with error: Operation response processing error`, func() {
				 rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
					 Crn: core.StringPtr(crn),
					 ZoneIdentifier: core.StringPtr(zoneIdentifier),
				 })
				 Expect(serviceErr).To(BeNil())
				 Expect(rulesetsService).ToNot(BeNil())
 
				 // Construct an instance of the GetInstanceEntryPointRulesetVersionOptions model
				 getInstanceEntryPointRulesetVersionOptionsModel := new(rulesetsv1.GetInstanceEntryPointRulesetVersionOptions)
				 getInstanceEntryPointRulesetVersionOptionsModel.RulesetPhase = core.StringPtr("ddos_l4")
				 getInstanceEntryPointRulesetVersionOptionsModel.RulesetVersion = core.StringPtr("1")
				 getInstanceEntryPointRulesetVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				 // Expect response parsing to fail since we are receiving a text/plain response
				 result, response, operationErr := rulesetsService.GetInstanceEntryPointRulesetVersion(getInstanceEntryPointRulesetVersionOptionsModel)
				 Expect(operationErr).ToNot(BeNil())
				 Expect(response).ToNot(BeNil())
				 Expect(result).To(BeNil())
 
				 // Enable retries and test again
				 rulesetsService.EnableRetries(0, 0)
				 result, response, operationErr = rulesetsService.GetInstanceEntryPointRulesetVersion(getInstanceEntryPointRulesetVersionOptionsModel)
				 Expect(operationErr).ToNot(BeNil())
				 Expect(response).ToNot(BeNil())
				 Expect(result).To(BeNil())
			 })
			 AfterEach(func() {
				 testServer.Close()
			 })
		 })
	 })
	 Describe(`GetInstanceEntryPointRulesetVersion(getInstanceEntryPointRulesetVersionOptions *GetInstanceEntryPointRulesetVersionOptions)`, func() {
		 crn := "testString"
		 zoneIdentifier := "testString"
		 getInstanceEntryPointRulesetVersionPath := "/v1/testString/rulesets/phases/ddos_l4/entrypoint/versions/1"
		 Context(`Using mock server endpoint with timeout`, func() {
			 BeforeEach(func() {
				 testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					 defer GinkgoRecover()
 
					 // Verify the contents of the request
					 Expect(req.URL.EscapedPath()).To(Equal(getInstanceEntryPointRulesetVersionPath))
					 Expect(req.Method).To(Equal("GET"))
 
					 // Sleep a short time to support a timeout test
					 time.Sleep(100 * time.Millisecond)
 
					 // Set mock response
					 res.Header().Set("Content-type", "application/json")
					 res.WriteHeader(200)
					 fmt.Fprintf(res, "%s", `{"success": true, "errors": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "messages": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "result": {"description": "Custom instance ruleset", "id": "ID", "kind": "managed", "last_updated": "2000-01-01T00:00:00.000000Z", "name": "Name", "phase": "ddos_l4", "version": "1", "rules": [{"id": "ID", "version": "Version", "action": "Action", "action_parameters": {"id": "ID", "overrides": {"action": "Action", "enabled": false, "sensitivity_level": "high", "rules": [{"id": "ID", "enabled": false, "action": "Action", "sensitivity_level": "high"}], "categories": [{"category": "Category", "enabled": false, "action": "Action"}]}, "version": "Version", "ruleset": "Ruleset", "rulesets": ["Rulesets"], "response": {"content": "{\"success\": false, \"error\": \"you have been blocked\"}", "content_type": "application/json", "status_code": 400}}, "categories": ["Categories"], "enabled": true, "description": "Description", "expression": "ip.src ne 1.1.1.1", "ref": "my_ref", "logging": {"enabled": true}, "last_updated": "2000-01-01T00:00:00.000000Z"}]}}`)
				 }))
			 })
			 It(`Invoke GetInstanceEntryPointRulesetVersion successfully with retries`, func() {
				 rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
					 Crn: core.StringPtr(crn),
					 ZoneIdentifier: core.StringPtr(zoneIdentifier),
				 })
				 Expect(serviceErr).To(BeNil())
				 Expect(rulesetsService).ToNot(BeNil())
				 rulesetsService.EnableRetries(0, 0)
 
				 // Construct an instance of the GetInstanceEntryPointRulesetVersionOptions model
				 getInstanceEntryPointRulesetVersionOptionsModel := new(rulesetsv1.GetInstanceEntryPointRulesetVersionOptions)
				 getInstanceEntryPointRulesetVersionOptionsModel.RulesetPhase = core.StringPtr("ddos_l4")
				 getInstanceEntryPointRulesetVersionOptionsModel.RulesetVersion = core.StringPtr("1")
				 getInstanceEntryPointRulesetVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
 
				 // Invoke operation with a Context to test a timeout error
				 ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				 defer cancelFunc()
				 _, _, operationErr := rulesetsService.GetInstanceEntryPointRulesetVersionWithContext(ctx, getInstanceEntryPointRulesetVersionOptionsModel)
				 Expect(operationErr).ToNot(BeNil())
				 Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
 
				 // Disable retries and test again
				 rulesetsService.DisableRetries()
				 result, response, operationErr := rulesetsService.GetInstanceEntryPointRulesetVersion(getInstanceEntryPointRulesetVersionOptionsModel)
				 Expect(operationErr).To(BeNil())
				 Expect(response).ToNot(BeNil())
				 Expect(result).ToNot(BeNil())
 
				 // Re-test the timeout error with retries disabled
				 ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				 defer cancelFunc2()
				 _, _, operationErr = rulesetsService.GetInstanceEntryPointRulesetVersionWithContext(ctx, getInstanceEntryPointRulesetVersionOptionsModel)
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
					 Expect(req.URL.EscapedPath()).To(Equal(getInstanceEntryPointRulesetVersionPath))
					 Expect(req.Method).To(Equal("GET"))
 
					 // Set mock response
					 res.Header().Set("Content-type", "application/json")
					 res.WriteHeader(200)
					 fmt.Fprintf(res, "%s", `{"success": true, "errors": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "messages": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "result": {"description": "Custom instance ruleset", "id": "ID", "kind": "managed", "last_updated": "2000-01-01T00:00:00.000000Z", "name": "Name", "phase": "ddos_l4", "version": "1", "rules": [{"id": "ID", "version": "Version", "action": "Action", "action_parameters": {"id": "ID", "overrides": {"action": "Action", "enabled": false, "sensitivity_level": "high", "rules": [{"id": "ID", "enabled": false, "action": "Action", "sensitivity_level": "high"}], "categories": [{"category": "Category", "enabled": false, "action": "Action"}]}, "version": "Version", "ruleset": "Ruleset", "rulesets": ["Rulesets"], "response": {"content": "{\"success\": false, \"error\": \"you have been blocked\"}", "content_type": "application/json", "status_code": 400}}, "categories": ["Categories"], "enabled": true, "description": "Description", "expression": "ip.src ne 1.1.1.1", "ref": "my_ref", "logging": {"enabled": true}, "last_updated": "2000-01-01T00:00:00.000000Z"}]}}`)
				 }))
			 })
			 It(`Invoke GetInstanceEntryPointRulesetVersion successfully`, func() {
				 rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
					 Crn: core.StringPtr(crn),
					 ZoneIdentifier: core.StringPtr(zoneIdentifier),
				 })
				 Expect(serviceErr).To(BeNil())
				 Expect(rulesetsService).ToNot(BeNil())
 
				 // Invoke operation with nil options model (negative test)
				 result, response, operationErr := rulesetsService.GetInstanceEntryPointRulesetVersion(nil)
				 Expect(operationErr).NotTo(BeNil())
				 Expect(response).To(BeNil())
				 Expect(result).To(BeNil())
 
				 // Construct an instance of the GetInstanceEntryPointRulesetVersionOptions model
				 getInstanceEntryPointRulesetVersionOptionsModel := new(rulesetsv1.GetInstanceEntryPointRulesetVersionOptions)
				 getInstanceEntryPointRulesetVersionOptionsModel.RulesetPhase = core.StringPtr("ddos_l4")
				 getInstanceEntryPointRulesetVersionOptionsModel.RulesetVersion = core.StringPtr("1")
				 getInstanceEntryPointRulesetVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
 
				 // Invoke operation with valid options model (positive test)
				 result, response, operationErr = rulesetsService.GetInstanceEntryPointRulesetVersion(getInstanceEntryPointRulesetVersionOptionsModel)
				 Expect(operationErr).To(BeNil())
				 Expect(response).ToNot(BeNil())
				 Expect(result).ToNot(BeNil())
 
			 })
			 It(`Invoke GetInstanceEntryPointRulesetVersion with error: Operation validation and request error`, func() {
				 rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
					 Crn: core.StringPtr(crn),
					 ZoneIdentifier: core.StringPtr(zoneIdentifier),
				 })
				 Expect(serviceErr).To(BeNil())
				 Expect(rulesetsService).ToNot(BeNil())
 
				 // Construct an instance of the GetInstanceEntryPointRulesetVersionOptions model
				 getInstanceEntryPointRulesetVersionOptionsModel := new(rulesetsv1.GetInstanceEntryPointRulesetVersionOptions)
				 getInstanceEntryPointRulesetVersionOptionsModel.RulesetPhase = core.StringPtr("ddos_l4")
				 getInstanceEntryPointRulesetVersionOptionsModel.RulesetVersion = core.StringPtr("1")
				 getInstanceEntryPointRulesetVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				 // Invoke operation with empty URL (negative test)
				 err := rulesetsService.SetServiceURL("")
				 Expect(err).To(BeNil())
				 result, response, operationErr := rulesetsService.GetInstanceEntryPointRulesetVersion(getInstanceEntryPointRulesetVersionOptionsModel)
				 Expect(operationErr).ToNot(BeNil())
				 Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				 Expect(response).To(BeNil())
				 Expect(result).To(BeNil())
				 // Construct a second instance of the GetInstanceEntryPointRulesetVersionOptions model with no property values
				 getInstanceEntryPointRulesetVersionOptionsModelNew := new(rulesetsv1.GetInstanceEntryPointRulesetVersionOptions)
				 // Invoke operation with invalid model (negative test)
				 result, response, operationErr = rulesetsService.GetInstanceEntryPointRulesetVersion(getInstanceEntryPointRulesetVersionOptionsModelNew)
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
			 It(`Invoke GetInstanceEntryPointRulesetVersion successfully`, func() {
				 rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
					 Crn: core.StringPtr(crn),
					 ZoneIdentifier: core.StringPtr(zoneIdentifier),
				 })
				 Expect(serviceErr).To(BeNil())
				 Expect(rulesetsService).ToNot(BeNil())
 
				 // Construct an instance of the GetInstanceEntryPointRulesetVersionOptions model
				 getInstanceEntryPointRulesetVersionOptionsModel := new(rulesetsv1.GetInstanceEntryPointRulesetVersionOptions)
				 getInstanceEntryPointRulesetVersionOptionsModel.RulesetPhase = core.StringPtr("ddos_l4")
				 getInstanceEntryPointRulesetVersionOptionsModel.RulesetVersion = core.StringPtr("1")
				 getInstanceEntryPointRulesetVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
 
				 // Invoke operation
				 result, response, operationErr := rulesetsService.GetInstanceEntryPointRulesetVersion(getInstanceEntryPointRulesetVersionOptionsModel)
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
	 Describe(`CreateInstanceRulesetRule(createInstanceRulesetRuleOptions *CreateInstanceRulesetRuleOptions) - Operation response error`, func() {
		 crn := "testString"
		 zoneIdentifier := "testString"
		 createInstanceRulesetRulePath := "/v1/testString/rulesets/testString/rules"
		 Context(`Using mock server endpoint with invalid JSON response`, func() {
			 BeforeEach(func() {
				 testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					 defer GinkgoRecover()
 
					 // Verify the contents of the request
					 Expect(req.URL.EscapedPath()).To(Equal(createInstanceRulesetRulePath))
					 Expect(req.Method).To(Equal("POST"))
					 res.Header().Set("Content-type", "application/json")
					 res.WriteHeader(200)
					 fmt.Fprint(res, `} this is not valid json {`)
				 }))
			 })
			 It(`Invoke CreateInstanceRulesetRule with error: Operation response processing error`, func() {
				 rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
					 Crn: core.StringPtr(crn),
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
				 actionParametersResponseModel.Content = core.StringPtr("{\"success\": false, \"error\": \"you have been blocked\"}")
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
 
				 // Construct an instance of the CreateInstanceRulesetRuleOptions model
				 createInstanceRulesetRuleOptionsModel := new(rulesetsv1.CreateInstanceRulesetRuleOptions)
				 createInstanceRulesetRuleOptionsModel.RulesetID = core.StringPtr("testString")
				 createInstanceRulesetRuleOptionsModel.Action = core.StringPtr("testString")
				 createInstanceRulesetRuleOptionsModel.ActionParameters = actionParametersModel
				 createInstanceRulesetRuleOptionsModel.Description = core.StringPtr("testString")
				 createInstanceRulesetRuleOptionsModel.Enabled = core.BoolPtr(true)
				 createInstanceRulesetRuleOptionsModel.Expression = core.StringPtr("ip.src ne 1.1.1.1")
				 createInstanceRulesetRuleOptionsModel.ID = core.StringPtr("testString")
				 createInstanceRulesetRuleOptionsModel.Logging = loggingModel
				 createInstanceRulesetRuleOptionsModel.Ref = core.StringPtr("my_ref")
				 createInstanceRulesetRuleOptionsModel.Position = positionModel
				 createInstanceRulesetRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				 // Expect response parsing to fail since we are receiving a text/plain response
				 result, response, operationErr := rulesetsService.CreateInstanceRulesetRule(createInstanceRulesetRuleOptionsModel)
				 Expect(operationErr).ToNot(BeNil())
				 Expect(response).ToNot(BeNil())
				 Expect(result).To(BeNil())
 
				 // Enable retries and test again
				 rulesetsService.EnableRetries(0, 0)
				 result, response, operationErr = rulesetsService.CreateInstanceRulesetRule(createInstanceRulesetRuleOptionsModel)
				 Expect(operationErr).ToNot(BeNil())
				 Expect(response).ToNot(BeNil())
				 Expect(result).To(BeNil())
			 })
			 AfterEach(func() {
				 testServer.Close()
			 })
		 })
	 })
	 Describe(`CreateInstanceRulesetRule(createInstanceRulesetRuleOptions *CreateInstanceRulesetRuleOptions)`, func() {
		 crn := "testString"
		 zoneIdentifier := "testString"
		 createInstanceRulesetRulePath := "/v1/testString/rulesets/testString/rules"
		 Context(`Using mock server endpoint with timeout`, func() {
			 BeforeEach(func() {
				 testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					 defer GinkgoRecover()
 
					 // Verify the contents of the request
					 Expect(req.URL.EscapedPath()).To(Equal(createInstanceRulesetRulePath))
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
					 fmt.Fprintf(res, "%s", `{"success": true, "errors": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "messages": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "result": {"id": "ID", "version": "Version", "action": "Action", "action_parameters": {"id": "ID", "overrides": {"action": "Action", "enabled": false, "sensitivity_level": "high", "rules": [{"id": "ID", "enabled": false, "action": "Action", "sensitivity_level": "high"}], "categories": [{"category": "Category", "enabled": false, "action": "Action"}]}, "version": "Version", "ruleset": "Ruleset", "rulesets": ["Rulesets"], "response": {"content": "{\"success\": false, \"error\": \"you have been blocked\"}", "content_type": "application/json", "status_code": 400}}, "categories": ["Categories"], "enabled": true, "description": "Description", "expression": "ip.src ne 1.1.1.1", "ref": "my_ref", "logging": {"enabled": true}, "last_updated": "2000-01-01T00:00:00.000000Z"}}`)
				 }))
			 })
			 It(`Invoke CreateInstanceRulesetRule successfully with retries`, func() {
				 rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
					 Crn: core.StringPtr(crn),
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
				 actionParametersResponseModel.Content = core.StringPtr("{\"success\": false, \"error\": \"you have been blocked\"}")
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
 
				 // Construct an instance of the CreateInstanceRulesetRuleOptions model
				 createInstanceRulesetRuleOptionsModel := new(rulesetsv1.CreateInstanceRulesetRuleOptions)
				 createInstanceRulesetRuleOptionsModel.RulesetID = core.StringPtr("testString")
				 createInstanceRulesetRuleOptionsModel.Action = core.StringPtr("testString")
				 createInstanceRulesetRuleOptionsModel.ActionParameters = actionParametersModel
				 createInstanceRulesetRuleOptionsModel.Description = core.StringPtr("testString")
				 createInstanceRulesetRuleOptionsModel.Enabled = core.BoolPtr(true)
				 createInstanceRulesetRuleOptionsModel.Expression = core.StringPtr("ip.src ne 1.1.1.1")
				 createInstanceRulesetRuleOptionsModel.ID = core.StringPtr("testString")
				 createInstanceRulesetRuleOptionsModel.Logging = loggingModel
				 createInstanceRulesetRuleOptionsModel.Ref = core.StringPtr("my_ref")
				 createInstanceRulesetRuleOptionsModel.Position = positionModel
				 createInstanceRulesetRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
 
				 // Invoke operation with a Context to test a timeout error
				 ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				 defer cancelFunc()
				 _, _, operationErr := rulesetsService.CreateInstanceRulesetRuleWithContext(ctx, createInstanceRulesetRuleOptionsModel)
				 Expect(operationErr).ToNot(BeNil())
				 Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
 
				 // Disable retries and test again
				 rulesetsService.DisableRetries()
				 result, response, operationErr := rulesetsService.CreateInstanceRulesetRule(createInstanceRulesetRuleOptionsModel)
				 Expect(operationErr).To(BeNil())
				 Expect(response).ToNot(BeNil())
				 Expect(result).ToNot(BeNil())
 
				 // Re-test the timeout error with retries disabled
				 ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				 defer cancelFunc2()
				 _, _, operationErr = rulesetsService.CreateInstanceRulesetRuleWithContext(ctx, createInstanceRulesetRuleOptionsModel)
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
					 Expect(req.URL.EscapedPath()).To(Equal(createInstanceRulesetRulePath))
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
					 fmt.Fprintf(res, "%s", `{"success": true, "errors": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "messages": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "result": {"id": "ID", "version": "Version", "action": "Action", "action_parameters": {"id": "ID", "overrides": {"action": "Action", "enabled": false, "sensitivity_level": "high", "rules": [{"id": "ID", "enabled": false, "action": "Action", "sensitivity_level": "high"}], "categories": [{"category": "Category", "enabled": false, "action": "Action"}]}, "version": "Version", "ruleset": "Ruleset", "rulesets": ["Rulesets"], "response": {"content": "{\"success\": false, \"error\": \"you have been blocked\"}", "content_type": "application/json", "status_code": 400}}, "categories": ["Categories"], "enabled": true, "description": "Description", "expression": "ip.src ne 1.1.1.1", "ref": "my_ref", "logging": {"enabled": true}, "last_updated": "2000-01-01T00:00:00.000000Z"}}`)
				 }))
			 })
			 It(`Invoke CreateInstanceRulesetRule successfully`, func() {
				 rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
					 Crn: core.StringPtr(crn),
					 ZoneIdentifier: core.StringPtr(zoneIdentifier),
				 })
				 Expect(serviceErr).To(BeNil())
				 Expect(rulesetsService).ToNot(BeNil())
 
				 // Invoke operation with nil options model (negative test)
				 result, response, operationErr := rulesetsService.CreateInstanceRulesetRule(nil)
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
				 actionParametersResponseModel.Content = core.StringPtr("{\"success\": false, \"error\": \"you have been blocked\"}")
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
 
				 // Construct an instance of the CreateInstanceRulesetRuleOptions model
				 createInstanceRulesetRuleOptionsModel := new(rulesetsv1.CreateInstanceRulesetRuleOptions)
				 createInstanceRulesetRuleOptionsModel.RulesetID = core.StringPtr("testString")
				 createInstanceRulesetRuleOptionsModel.Action = core.StringPtr("testString")
				 createInstanceRulesetRuleOptionsModel.ActionParameters = actionParametersModel
				 createInstanceRulesetRuleOptionsModel.Description = core.StringPtr("testString")
				 createInstanceRulesetRuleOptionsModel.Enabled = core.BoolPtr(true)
				 createInstanceRulesetRuleOptionsModel.Expression = core.StringPtr("ip.src ne 1.1.1.1")
				 createInstanceRulesetRuleOptionsModel.ID = core.StringPtr("testString")
				 createInstanceRulesetRuleOptionsModel.Logging = loggingModel
				 createInstanceRulesetRuleOptionsModel.Ref = core.StringPtr("my_ref")
				 createInstanceRulesetRuleOptionsModel.Position = positionModel
				 createInstanceRulesetRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
 
				 // Invoke operation with valid options model (positive test)
				 result, response, operationErr = rulesetsService.CreateInstanceRulesetRule(createInstanceRulesetRuleOptionsModel)
				 Expect(operationErr).To(BeNil())
				 Expect(response).ToNot(BeNil())
				 Expect(result).ToNot(BeNil())
 
			 })
			 It(`Invoke CreateInstanceRulesetRule with error: Operation validation and request error`, func() {
				 rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
					 Crn: core.StringPtr(crn),
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
				 actionParametersResponseModel.Content = core.StringPtr("{\"success\": false, \"error\": \"you have been blocked\"}")
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
 
				 // Construct an instance of the CreateInstanceRulesetRuleOptions model
				 createInstanceRulesetRuleOptionsModel := new(rulesetsv1.CreateInstanceRulesetRuleOptions)
				 createInstanceRulesetRuleOptionsModel.RulesetID = core.StringPtr("testString")
				 createInstanceRulesetRuleOptionsModel.Action = core.StringPtr("testString")
				 createInstanceRulesetRuleOptionsModel.ActionParameters = actionParametersModel
				 createInstanceRulesetRuleOptionsModel.Description = core.StringPtr("testString")
				 createInstanceRulesetRuleOptionsModel.Enabled = core.BoolPtr(true)
				 createInstanceRulesetRuleOptionsModel.Expression = core.StringPtr("ip.src ne 1.1.1.1")
				 createInstanceRulesetRuleOptionsModel.ID = core.StringPtr("testString")
				 createInstanceRulesetRuleOptionsModel.Logging = loggingModel
				 createInstanceRulesetRuleOptionsModel.Ref = core.StringPtr("my_ref")
				 createInstanceRulesetRuleOptionsModel.Position = positionModel
				 createInstanceRulesetRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				 // Invoke operation with empty URL (negative test)
				 err := rulesetsService.SetServiceURL("")
				 Expect(err).To(BeNil())
				 result, response, operationErr := rulesetsService.CreateInstanceRulesetRule(createInstanceRulesetRuleOptionsModel)
				 Expect(operationErr).ToNot(BeNil())
				 Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				 Expect(response).To(BeNil())
				 Expect(result).To(BeNil())
				 // Construct a second instance of the CreateInstanceRulesetRuleOptions model with no property values
				 createInstanceRulesetRuleOptionsModelNew := new(rulesetsv1.CreateInstanceRulesetRuleOptions)
				 // Invoke operation with invalid model (negative test)
				 result, response, operationErr = rulesetsService.CreateInstanceRulesetRule(createInstanceRulesetRuleOptionsModelNew)
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
			 It(`Invoke CreateInstanceRulesetRule successfully`, func() {
				 rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
					 Crn: core.StringPtr(crn),
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
				 actionParametersResponseModel.Content = core.StringPtr("{\"success\": false, \"error\": \"you have been blocked\"}")
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
 
				 // Construct an instance of the CreateInstanceRulesetRuleOptions model
				 createInstanceRulesetRuleOptionsModel := new(rulesetsv1.CreateInstanceRulesetRuleOptions)
				 createInstanceRulesetRuleOptionsModel.RulesetID = core.StringPtr("testString")
				 createInstanceRulesetRuleOptionsModel.Action = core.StringPtr("testString")
				 createInstanceRulesetRuleOptionsModel.ActionParameters = actionParametersModel
				 createInstanceRulesetRuleOptionsModel.Description = core.StringPtr("testString")
				 createInstanceRulesetRuleOptionsModel.Enabled = core.BoolPtr(true)
				 createInstanceRulesetRuleOptionsModel.Expression = core.StringPtr("ip.src ne 1.1.1.1")
				 createInstanceRulesetRuleOptionsModel.ID = core.StringPtr("testString")
				 createInstanceRulesetRuleOptionsModel.Logging = loggingModel
				 createInstanceRulesetRuleOptionsModel.Ref = core.StringPtr("my_ref")
				 createInstanceRulesetRuleOptionsModel.Position = positionModel
				 createInstanceRulesetRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
 
				 // Invoke operation
				 result, response, operationErr := rulesetsService.CreateInstanceRulesetRule(createInstanceRulesetRuleOptionsModel)
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
	 Describe(`UpdateInstanceRulesetRule(updateInstanceRulesetRuleOptions *UpdateInstanceRulesetRuleOptions) - Operation response error`, func() {
		 crn := "testString"
		 zoneIdentifier := "testString"
		 updateInstanceRulesetRulePath := "/v1/testString/rulesets/testString/rules/testString"
		 Context(`Using mock server endpoint with invalid JSON response`, func() {
			 BeforeEach(func() {
				 testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					 defer GinkgoRecover()
 
					 // Verify the contents of the request
					 Expect(req.URL.EscapedPath()).To(Equal(updateInstanceRulesetRulePath))
					 Expect(req.Method).To(Equal("PATCH"))
					 res.Header().Set("Content-type", "application/json")
					 res.WriteHeader(200)
					 fmt.Fprint(res, `} this is not valid json {`)
				 }))
			 })
			 It(`Invoke UpdateInstanceRulesetRule with error: Operation response processing error`, func() {
				 rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
					 Crn: core.StringPtr(crn),
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
				 actionParametersResponseModel.Content = core.StringPtr("{\"success\": false, \"error\": \"you have been blocked\"}")
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
 
				 // Construct an instance of the UpdateInstanceRulesetRuleOptions model
				 updateInstanceRulesetRuleOptionsModel := new(rulesetsv1.UpdateInstanceRulesetRuleOptions)
				 updateInstanceRulesetRuleOptionsModel.RulesetID = core.StringPtr("testString")
				 updateInstanceRulesetRuleOptionsModel.RuleID = core.StringPtr("testString")
				 updateInstanceRulesetRuleOptionsModel.Action = core.StringPtr("testString")
				 updateInstanceRulesetRuleOptionsModel.ActionParameters = actionParametersModel
				 updateInstanceRulesetRuleOptionsModel.Description = core.StringPtr("testString")
				 updateInstanceRulesetRuleOptionsModel.Enabled = core.BoolPtr(true)
				 updateInstanceRulesetRuleOptionsModel.Expression = core.StringPtr("ip.src ne 1.1.1.1")
				 updateInstanceRulesetRuleOptionsModel.ID = core.StringPtr("testString")
				 updateInstanceRulesetRuleOptionsModel.Logging = loggingModel
				 updateInstanceRulesetRuleOptionsModel.Ref = core.StringPtr("my_ref")
				 updateInstanceRulesetRuleOptionsModel.Position = positionModel
				 updateInstanceRulesetRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				 // Expect response parsing to fail since we are receiving a text/plain response
				 result, response, operationErr := rulesetsService.UpdateInstanceRulesetRule(updateInstanceRulesetRuleOptionsModel)
				 Expect(operationErr).ToNot(BeNil())
				 Expect(response).ToNot(BeNil())
				 Expect(result).To(BeNil())
 
				 // Enable retries and test again
				 rulesetsService.EnableRetries(0, 0)
				 result, response, operationErr = rulesetsService.UpdateInstanceRulesetRule(updateInstanceRulesetRuleOptionsModel)
				 Expect(operationErr).ToNot(BeNil())
				 Expect(response).ToNot(BeNil())
				 Expect(result).To(BeNil())
			 })
			 AfterEach(func() {
				 testServer.Close()
			 })
		 })
	 })
	 Describe(`UpdateInstanceRulesetRule(updateInstanceRulesetRuleOptions *UpdateInstanceRulesetRuleOptions)`, func() {
		 crn := "testString"
		 zoneIdentifier := "testString"
		 updateInstanceRulesetRulePath := "/v1/testString/rulesets/testString/rules/testString"
		 Context(`Using mock server endpoint with timeout`, func() {
			 BeforeEach(func() {
				 testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					 defer GinkgoRecover()
 
					 // Verify the contents of the request
					 Expect(req.URL.EscapedPath()).To(Equal(updateInstanceRulesetRulePath))
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
					 fmt.Fprintf(res, "%s", `{"success": true, "errors": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "messages": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "result": {"id": "ID", "version": "Version", "action": "Action", "action_parameters": {"id": "ID", "overrides": {"action": "Action", "enabled": false, "sensitivity_level": "high", "rules": [{"id": "ID", "enabled": false, "action": "Action", "sensitivity_level": "high"}], "categories": [{"category": "Category", "enabled": false, "action": "Action"}]}, "version": "Version", "ruleset": "Ruleset", "rulesets": ["Rulesets"], "response": {"content": "{\"success\": false, \"error\": \"you have been blocked\"}", "content_type": "application/json", "status_code": 400}}, "categories": ["Categories"], "enabled": true, "description": "Description", "expression": "ip.src ne 1.1.1.1", "ref": "my_ref", "logging": {"enabled": true}, "last_updated": "2000-01-01T00:00:00.000000Z"}}`)
				 }))
			 })
			 It(`Invoke UpdateInstanceRulesetRule successfully with retries`, func() {
				 rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
					 Crn: core.StringPtr(crn),
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
				 actionParametersResponseModel.Content = core.StringPtr("{\"success\": false, \"error\": \"you have been blocked\"}")
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
 
				 // Construct an instance of the UpdateInstanceRulesetRuleOptions model
				 updateInstanceRulesetRuleOptionsModel := new(rulesetsv1.UpdateInstanceRulesetRuleOptions)
				 updateInstanceRulesetRuleOptionsModel.RulesetID = core.StringPtr("testString")
				 updateInstanceRulesetRuleOptionsModel.RuleID = core.StringPtr("testString")
				 updateInstanceRulesetRuleOptionsModel.Action = core.StringPtr("testString")
				 updateInstanceRulesetRuleOptionsModel.ActionParameters = actionParametersModel
				 updateInstanceRulesetRuleOptionsModel.Description = core.StringPtr("testString")
				 updateInstanceRulesetRuleOptionsModel.Enabled = core.BoolPtr(true)
				 updateInstanceRulesetRuleOptionsModel.Expression = core.StringPtr("ip.src ne 1.1.1.1")
				 updateInstanceRulesetRuleOptionsModel.ID = core.StringPtr("testString")
				 updateInstanceRulesetRuleOptionsModel.Logging = loggingModel
				 updateInstanceRulesetRuleOptionsModel.Ref = core.StringPtr("my_ref")
				 updateInstanceRulesetRuleOptionsModel.Position = positionModel
				 updateInstanceRulesetRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
 
				 // Invoke operation with a Context to test a timeout error
				 ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				 defer cancelFunc()
				 _, _, operationErr := rulesetsService.UpdateInstanceRulesetRuleWithContext(ctx, updateInstanceRulesetRuleOptionsModel)
				 Expect(operationErr).ToNot(BeNil())
				 Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
 
				 // Disable retries and test again
				 rulesetsService.DisableRetries()
				 result, response, operationErr := rulesetsService.UpdateInstanceRulesetRule(updateInstanceRulesetRuleOptionsModel)
				 Expect(operationErr).To(BeNil())
				 Expect(response).ToNot(BeNil())
				 Expect(result).ToNot(BeNil())
 
				 // Re-test the timeout error with retries disabled
				 ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				 defer cancelFunc2()
				 _, _, operationErr = rulesetsService.UpdateInstanceRulesetRuleWithContext(ctx, updateInstanceRulesetRuleOptionsModel)
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
					 Expect(req.URL.EscapedPath()).To(Equal(updateInstanceRulesetRulePath))
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
					 fmt.Fprintf(res, "%s", `{"success": true, "errors": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "messages": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "result": {"id": "ID", "version": "Version", "action": "Action", "action_parameters": {"id": "ID", "overrides": {"action": "Action", "enabled": false, "sensitivity_level": "high", "rules": [{"id": "ID", "enabled": false, "action": "Action", "sensitivity_level": "high"}], "categories": [{"category": "Category", "enabled": false, "action": "Action"}]}, "version": "Version", "ruleset": "Ruleset", "rulesets": ["Rulesets"], "response": {"content": "{\"success\": false, \"error\": \"you have been blocked\"}", "content_type": "application/json", "status_code": 400}}, "categories": ["Categories"], "enabled": true, "description": "Description", "expression": "ip.src ne 1.1.1.1", "ref": "my_ref", "logging": {"enabled": true}, "last_updated": "2000-01-01T00:00:00.000000Z"}}`)
				 }))
			 })
			 It(`Invoke UpdateInstanceRulesetRule successfully`, func() {
				 rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
					 Crn: core.StringPtr(crn),
					 ZoneIdentifier: core.StringPtr(zoneIdentifier),
				 })
				 Expect(serviceErr).To(BeNil())
				 Expect(rulesetsService).ToNot(BeNil())
 
				 // Invoke operation with nil options model (negative test)
				 result, response, operationErr := rulesetsService.UpdateInstanceRulesetRule(nil)
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
				 actionParametersResponseModel.Content = core.StringPtr("{\"success\": false, \"error\": \"you have been blocked\"}")
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
 
				 // Construct an instance of the UpdateInstanceRulesetRuleOptions model
				 updateInstanceRulesetRuleOptionsModel := new(rulesetsv1.UpdateInstanceRulesetRuleOptions)
				 updateInstanceRulesetRuleOptionsModel.RulesetID = core.StringPtr("testString")
				 updateInstanceRulesetRuleOptionsModel.RuleID = core.StringPtr("testString")
				 updateInstanceRulesetRuleOptionsModel.Action = core.StringPtr("testString")
				 updateInstanceRulesetRuleOptionsModel.ActionParameters = actionParametersModel
				 updateInstanceRulesetRuleOptionsModel.Description = core.StringPtr("testString")
				 updateInstanceRulesetRuleOptionsModel.Enabled = core.BoolPtr(true)
				 updateInstanceRulesetRuleOptionsModel.Expression = core.StringPtr("ip.src ne 1.1.1.1")
				 updateInstanceRulesetRuleOptionsModel.ID = core.StringPtr("testString")
				 updateInstanceRulesetRuleOptionsModel.Logging = loggingModel
				 updateInstanceRulesetRuleOptionsModel.Ref = core.StringPtr("my_ref")
				 updateInstanceRulesetRuleOptionsModel.Position = positionModel
				 updateInstanceRulesetRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
 
				 // Invoke operation with valid options model (positive test)
				 result, response, operationErr = rulesetsService.UpdateInstanceRulesetRule(updateInstanceRulesetRuleOptionsModel)
				 Expect(operationErr).To(BeNil())
				 Expect(response).ToNot(BeNil())
				 Expect(result).ToNot(BeNil())
 
			 })
			 It(`Invoke UpdateInstanceRulesetRule with error: Operation validation and request error`, func() {
				 rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
					 Crn: core.StringPtr(crn),
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
				 actionParametersResponseModel.Content = core.StringPtr("{\"success\": false, \"error\": \"you have been blocked\"}")
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
 
				 // Construct an instance of the UpdateInstanceRulesetRuleOptions model
				 updateInstanceRulesetRuleOptionsModel := new(rulesetsv1.UpdateInstanceRulesetRuleOptions)
				 updateInstanceRulesetRuleOptionsModel.RulesetID = core.StringPtr("testString")
				 updateInstanceRulesetRuleOptionsModel.RuleID = core.StringPtr("testString")
				 updateInstanceRulesetRuleOptionsModel.Action = core.StringPtr("testString")
				 updateInstanceRulesetRuleOptionsModel.ActionParameters = actionParametersModel
				 updateInstanceRulesetRuleOptionsModel.Description = core.StringPtr("testString")
				 updateInstanceRulesetRuleOptionsModel.Enabled = core.BoolPtr(true)
				 updateInstanceRulesetRuleOptionsModel.Expression = core.StringPtr("ip.src ne 1.1.1.1")
				 updateInstanceRulesetRuleOptionsModel.ID = core.StringPtr("testString")
				 updateInstanceRulesetRuleOptionsModel.Logging = loggingModel
				 updateInstanceRulesetRuleOptionsModel.Ref = core.StringPtr("my_ref")
				 updateInstanceRulesetRuleOptionsModel.Position = positionModel
				 updateInstanceRulesetRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				 // Invoke operation with empty URL (negative test)
				 err := rulesetsService.SetServiceURL("")
				 Expect(err).To(BeNil())
				 result, response, operationErr := rulesetsService.UpdateInstanceRulesetRule(updateInstanceRulesetRuleOptionsModel)
				 Expect(operationErr).ToNot(BeNil())
				 Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				 Expect(response).To(BeNil())
				 Expect(result).To(BeNil())
				 // Construct a second instance of the UpdateInstanceRulesetRuleOptions model with no property values
				 updateInstanceRulesetRuleOptionsModelNew := new(rulesetsv1.UpdateInstanceRulesetRuleOptions)
				 // Invoke operation with invalid model (negative test)
				 result, response, operationErr = rulesetsService.UpdateInstanceRulesetRule(updateInstanceRulesetRuleOptionsModelNew)
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
			 It(`Invoke UpdateInstanceRulesetRule successfully`, func() {
				 rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
					 Crn: core.StringPtr(crn),
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
				 actionParametersResponseModel.Content = core.StringPtr("{\"success\": false, \"error\": \"you have been blocked\"}")
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
 
				 // Construct an instance of the UpdateInstanceRulesetRuleOptions model
				 updateInstanceRulesetRuleOptionsModel := new(rulesetsv1.UpdateInstanceRulesetRuleOptions)
				 updateInstanceRulesetRuleOptionsModel.RulesetID = core.StringPtr("testString")
				 updateInstanceRulesetRuleOptionsModel.RuleID = core.StringPtr("testString")
				 updateInstanceRulesetRuleOptionsModel.Action = core.StringPtr("testString")
				 updateInstanceRulesetRuleOptionsModel.ActionParameters = actionParametersModel
				 updateInstanceRulesetRuleOptionsModel.Description = core.StringPtr("testString")
				 updateInstanceRulesetRuleOptionsModel.Enabled = core.BoolPtr(true)
				 updateInstanceRulesetRuleOptionsModel.Expression = core.StringPtr("ip.src ne 1.1.1.1")
				 updateInstanceRulesetRuleOptionsModel.ID = core.StringPtr("testString")
				 updateInstanceRulesetRuleOptionsModel.Logging = loggingModel
				 updateInstanceRulesetRuleOptionsModel.Ref = core.StringPtr("my_ref")
				 updateInstanceRulesetRuleOptionsModel.Position = positionModel
				 updateInstanceRulesetRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
 
				 // Invoke operation
				 result, response, operationErr := rulesetsService.UpdateInstanceRulesetRule(updateInstanceRulesetRuleOptionsModel)
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
	 Describe(`DeleteInstanceRulesetRule(deleteInstanceRulesetRuleOptions *DeleteInstanceRulesetRuleOptions) - Operation response error`, func() {
		 crn := "testString"
		 zoneIdentifier := "testString"
		 deleteInstanceRulesetRulePath := "/v1/testString/rulesets/testString/rules/testString"
		 Context(`Using mock server endpoint with invalid JSON response`, func() {
			 BeforeEach(func() {
				 testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					 defer GinkgoRecover()
 
					 // Verify the contents of the request
					 Expect(req.URL.EscapedPath()).To(Equal(deleteInstanceRulesetRulePath))
					 Expect(req.Method).To(Equal("DELETE"))
					 res.Header().Set("Content-type", "application/json")
					 res.WriteHeader(200)
					 fmt.Fprint(res, `} this is not valid json {`)
				 }))
			 })
			 It(`Invoke DeleteInstanceRulesetRule with error: Operation response processing error`, func() {
				 rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
					 Crn: core.StringPtr(crn),
					 ZoneIdentifier: core.StringPtr(zoneIdentifier),
				 })
				 Expect(serviceErr).To(BeNil())
				 Expect(rulesetsService).ToNot(BeNil())
 
				 // Construct an instance of the DeleteInstanceRulesetRuleOptions model
				 deleteInstanceRulesetRuleOptionsModel := new(rulesetsv1.DeleteInstanceRulesetRuleOptions)
				 deleteInstanceRulesetRuleOptionsModel.RulesetID = core.StringPtr("testString")
				 deleteInstanceRulesetRuleOptionsModel.RuleID = core.StringPtr("testString")
				 deleteInstanceRulesetRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				 // Expect response parsing to fail since we are receiving a text/plain response
				 result, response, operationErr := rulesetsService.DeleteInstanceRulesetRule(deleteInstanceRulesetRuleOptionsModel)
				 Expect(operationErr).ToNot(BeNil())
				 Expect(response).ToNot(BeNil())
				 Expect(result).To(BeNil())
 
				 // Enable retries and test again
				 rulesetsService.EnableRetries(0, 0)
				 result, response, operationErr = rulesetsService.DeleteInstanceRulesetRule(deleteInstanceRulesetRuleOptionsModel)
				 Expect(operationErr).ToNot(BeNil())
				 Expect(response).ToNot(BeNil())
				 Expect(result).To(BeNil())
			 })
			 AfterEach(func() {
				 testServer.Close()
			 })
		 })
	 })
	 Describe(`DeleteInstanceRulesetRule(deleteInstanceRulesetRuleOptions *DeleteInstanceRulesetRuleOptions)`, func() {
		 crn := "testString"
		 zoneIdentifier := "testString"
		 deleteInstanceRulesetRulePath := "/v1/testString/rulesets/testString/rules/testString"
		 Context(`Using mock server endpoint with timeout`, func() {
			 BeforeEach(func() {
				 testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					 defer GinkgoRecover()
 
					 // Verify the contents of the request
					 Expect(req.URL.EscapedPath()).To(Equal(deleteInstanceRulesetRulePath))
					 Expect(req.Method).To(Equal("DELETE"))
 
					 // Sleep a short time to support a timeout test
					 time.Sleep(100 * time.Millisecond)
 
					 // Set mock response
					 res.Header().Set("Content-type", "application/json")
					 res.WriteHeader(200)
					 fmt.Fprintf(res, "%s", `{"success": true, "errors": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "messages": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "result": {"id": "ID", "version": "Version", "action": "Action", "action_parameters": {"id": "ID", "overrides": {"action": "Action", "enabled": false, "sensitivity_level": "high", "rules": [{"id": "ID", "enabled": false, "action": "Action", "sensitivity_level": "high"}], "categories": [{"category": "Category", "enabled": false, "action": "Action"}]}, "version": "Version", "ruleset": "Ruleset", "rulesets": ["Rulesets"], "response": {"content": "{\"success\": false, \"error\": \"you have been blocked\"}", "content_type": "application/json", "status_code": 400}}, "categories": ["Categories"], "enabled": true, "description": "Description", "expression": "ip.src ne 1.1.1.1", "ref": "my_ref", "logging": {"enabled": true}, "last_updated": "2000-01-01T00:00:00.000000Z"}}`)
				 }))
			 })
			 It(`Invoke DeleteInstanceRulesetRule successfully with retries`, func() {
				 rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
					 Crn: core.StringPtr(crn),
					 ZoneIdentifier: core.StringPtr(zoneIdentifier),
				 })
				 Expect(serviceErr).To(BeNil())
				 Expect(rulesetsService).ToNot(BeNil())
				 rulesetsService.EnableRetries(0, 0)
 
				 // Construct an instance of the DeleteInstanceRulesetRuleOptions model
				 deleteInstanceRulesetRuleOptionsModel := new(rulesetsv1.DeleteInstanceRulesetRuleOptions)
				 deleteInstanceRulesetRuleOptionsModel.RulesetID = core.StringPtr("testString")
				 deleteInstanceRulesetRuleOptionsModel.RuleID = core.StringPtr("testString")
				 deleteInstanceRulesetRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
 
				 // Invoke operation with a Context to test a timeout error
				 ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				 defer cancelFunc()
				 _, _, operationErr := rulesetsService.DeleteInstanceRulesetRuleWithContext(ctx, deleteInstanceRulesetRuleOptionsModel)
				 Expect(operationErr).ToNot(BeNil())
				 Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
 
				 // Disable retries and test again
				 rulesetsService.DisableRetries()
				 result, response, operationErr := rulesetsService.DeleteInstanceRulesetRule(deleteInstanceRulesetRuleOptionsModel)
				 Expect(operationErr).To(BeNil())
				 Expect(response).ToNot(BeNil())
				 Expect(result).ToNot(BeNil())
 
				 // Re-test the timeout error with retries disabled
				 ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				 defer cancelFunc2()
				 _, _, operationErr = rulesetsService.DeleteInstanceRulesetRuleWithContext(ctx, deleteInstanceRulesetRuleOptionsModel)
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
					 Expect(req.URL.EscapedPath()).To(Equal(deleteInstanceRulesetRulePath))
					 Expect(req.Method).To(Equal("DELETE"))
 
					 // Set mock response
					 res.Header().Set("Content-type", "application/json")
					 res.WriteHeader(200)
					 fmt.Fprintf(res, "%s", `{"success": true, "errors": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "messages": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "result": {"id": "ID", "version": "Version", "action": "Action", "action_parameters": {"id": "ID", "overrides": {"action": "Action", "enabled": false, "sensitivity_level": "high", "rules": [{"id": "ID", "enabled": false, "action": "Action", "sensitivity_level": "high"}], "categories": [{"category": "Category", "enabled": false, "action": "Action"}]}, "version": "Version", "ruleset": "Ruleset", "rulesets": ["Rulesets"], "response": {"content": "{\"success\": false, \"error\": \"you have been blocked\"}", "content_type": "application/json", "status_code": 400}}, "categories": ["Categories"], "enabled": true, "description": "Description", "expression": "ip.src ne 1.1.1.1", "ref": "my_ref", "logging": {"enabled": true}, "last_updated": "2000-01-01T00:00:00.000000Z"}}`)
				 }))
			 })
			 It(`Invoke DeleteInstanceRulesetRule successfully`, func() {
				 rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
					 Crn: core.StringPtr(crn),
					 ZoneIdentifier: core.StringPtr(zoneIdentifier),
				 })
				 Expect(serviceErr).To(BeNil())
				 Expect(rulesetsService).ToNot(BeNil())
 
				 // Invoke operation with nil options model (negative test)
				 result, response, operationErr := rulesetsService.DeleteInstanceRulesetRule(nil)
				 Expect(operationErr).NotTo(BeNil())
				 Expect(response).To(BeNil())
				 Expect(result).To(BeNil())
 
				 // Construct an instance of the DeleteInstanceRulesetRuleOptions model
				 deleteInstanceRulesetRuleOptionsModel := new(rulesetsv1.DeleteInstanceRulesetRuleOptions)
				 deleteInstanceRulesetRuleOptionsModel.RulesetID = core.StringPtr("testString")
				 deleteInstanceRulesetRuleOptionsModel.RuleID = core.StringPtr("testString")
				 deleteInstanceRulesetRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
 
				 // Invoke operation with valid options model (positive test)
				 result, response, operationErr = rulesetsService.DeleteInstanceRulesetRule(deleteInstanceRulesetRuleOptionsModel)
				 Expect(operationErr).To(BeNil())
				 Expect(response).ToNot(BeNil())
				 Expect(result).ToNot(BeNil())
 
			 })
			 It(`Invoke DeleteInstanceRulesetRule with error: Operation validation and request error`, func() {
				 rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
					 Crn: core.StringPtr(crn),
					 ZoneIdentifier: core.StringPtr(zoneIdentifier),
				 })
				 Expect(serviceErr).To(BeNil())
				 Expect(rulesetsService).ToNot(BeNil())
 
				 // Construct an instance of the DeleteInstanceRulesetRuleOptions model
				 deleteInstanceRulesetRuleOptionsModel := new(rulesetsv1.DeleteInstanceRulesetRuleOptions)
				 deleteInstanceRulesetRuleOptionsModel.RulesetID = core.StringPtr("testString")
				 deleteInstanceRulesetRuleOptionsModel.RuleID = core.StringPtr("testString")
				 deleteInstanceRulesetRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				 // Invoke operation with empty URL (negative test)
				 err := rulesetsService.SetServiceURL("")
				 Expect(err).To(BeNil())
				 result, response, operationErr := rulesetsService.DeleteInstanceRulesetRule(deleteInstanceRulesetRuleOptionsModel)
				 Expect(operationErr).ToNot(BeNil())
				 Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				 Expect(response).To(BeNil())
				 Expect(result).To(BeNil())
				 // Construct a second instance of the DeleteInstanceRulesetRuleOptions model with no property values
				 deleteInstanceRulesetRuleOptionsModelNew := new(rulesetsv1.DeleteInstanceRulesetRuleOptions)
				 // Invoke operation with invalid model (negative test)
				 result, response, operationErr = rulesetsService.DeleteInstanceRulesetRule(deleteInstanceRulesetRuleOptionsModelNew)
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
			 It(`Invoke DeleteInstanceRulesetRule successfully`, func() {
				 rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
					 Crn: core.StringPtr(crn),
					 ZoneIdentifier: core.StringPtr(zoneIdentifier),
				 })
				 Expect(serviceErr).To(BeNil())
				 Expect(rulesetsService).ToNot(BeNil())
 
				 // Construct an instance of the DeleteInstanceRulesetRuleOptions model
				 deleteInstanceRulesetRuleOptionsModel := new(rulesetsv1.DeleteInstanceRulesetRuleOptions)
				 deleteInstanceRulesetRuleOptionsModel.RulesetID = core.StringPtr("testString")
				 deleteInstanceRulesetRuleOptionsModel.RuleID = core.StringPtr("testString")
				 deleteInstanceRulesetRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
 
				 // Invoke operation
				 result, response, operationErr := rulesetsService.DeleteInstanceRulesetRule(deleteInstanceRulesetRuleOptionsModel)
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
	 Describe(`GetInstanceRulesetVersionByTag(getInstanceRulesetVersionByTagOptions *GetInstanceRulesetVersionByTagOptions) - Operation response error`, func() {
		 crn := "testString"
		 zoneIdentifier := "testString"
		 getInstanceRulesetVersionByTagPath := "/v1/testString/rulesets/testString/versions/1/by_tag/testString"
		 Context(`Using mock server endpoint with invalid JSON response`, func() {
			 BeforeEach(func() {
				 testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					 defer GinkgoRecover()
 
					 // Verify the contents of the request
					 Expect(req.URL.EscapedPath()).To(Equal(getInstanceRulesetVersionByTagPath))
					 Expect(req.Method).To(Equal("GET"))
					 res.Header().Set("Content-type", "application/json")
					 res.WriteHeader(200)
					 fmt.Fprint(res, `} this is not valid json {`)
				 }))
			 })
			 It(`Invoke GetInstanceRulesetVersionByTag with error: Operation response processing error`, func() {
				 rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
					 Crn: core.StringPtr(crn),
					 ZoneIdentifier: core.StringPtr(zoneIdentifier),
				 })
				 Expect(serviceErr).To(BeNil())
				 Expect(rulesetsService).ToNot(BeNil())
 
				 // Construct an instance of the GetInstanceRulesetVersionByTagOptions model
				 getInstanceRulesetVersionByTagOptionsModel := new(rulesetsv1.GetInstanceRulesetVersionByTagOptions)
				 getInstanceRulesetVersionByTagOptionsModel.RulesetID = core.StringPtr("testString")
				 getInstanceRulesetVersionByTagOptionsModel.RulesetVersion = core.StringPtr("1")
				 getInstanceRulesetVersionByTagOptionsModel.RuleTag = core.StringPtr("testString")
				 getInstanceRulesetVersionByTagOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				 // Expect response parsing to fail since we are receiving a text/plain response
				 result, response, operationErr := rulesetsService.GetInstanceRulesetVersionByTag(getInstanceRulesetVersionByTagOptionsModel)
				 Expect(operationErr).ToNot(BeNil())
				 Expect(response).ToNot(BeNil())
				 Expect(result).To(BeNil())
 
				 // Enable retries and test again
				 rulesetsService.EnableRetries(0, 0)
				 result, response, operationErr = rulesetsService.GetInstanceRulesetVersionByTag(getInstanceRulesetVersionByTagOptionsModel)
				 Expect(operationErr).ToNot(BeNil())
				 Expect(response).ToNot(BeNil())
				 Expect(result).To(BeNil())
			 })
			 AfterEach(func() {
				 testServer.Close()
			 })
		 })
	 })
	 Describe(`GetInstanceRulesetVersionByTag(getInstanceRulesetVersionByTagOptions *GetInstanceRulesetVersionByTagOptions)`, func() {
		 crn := "testString"
		 zoneIdentifier := "testString"
		 getInstanceRulesetVersionByTagPath := "/v1/testString/rulesets/testString/versions/1/by_tag/testString"
		 Context(`Using mock server endpoint with timeout`, func() {
			 BeforeEach(func() {
				 testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					 defer GinkgoRecover()
 
					 // Verify the contents of the request
					 Expect(req.URL.EscapedPath()).To(Equal(getInstanceRulesetVersionByTagPath))
					 Expect(req.Method).To(Equal("GET"))
 
					 // Sleep a short time to support a timeout test
					 time.Sleep(100 * time.Millisecond)
 
					 // Set mock response
					 res.Header().Set("Content-type", "application/json")
					 res.WriteHeader(200)
					 fmt.Fprintf(res, "%s", `{"success": true, "errors": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "messages": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "result": {"description": "Custom instance ruleset", "id": "ID", "kind": "managed", "last_updated": "2000-01-01T00:00:00.000000Z", "name": "Name", "phase": "ddos_l4", "version": "1", "rules": [{"id": "ID", "version": "Version", "action": "Action", "action_parameters": {"id": "ID", "overrides": {"action": "Action", "enabled": false, "sensitivity_level": "high", "rules": [{"id": "ID", "enabled": false, "action": "Action", "sensitivity_level": "high"}], "categories": [{"category": "Category", "enabled": false, "action": "Action"}]}, "version": "Version", "ruleset": "Ruleset", "rulesets": ["Rulesets"], "response": {"content": "{\"success\": false, \"error\": \"you have been blocked\"}", "content_type": "application/json", "status_code": 400}}, "categories": ["Categories"], "enabled": true, "description": "Description", "expression": "ip.src ne 1.1.1.1", "ref": "my_ref", "logging": {"enabled": true}, "last_updated": "2000-01-01T00:00:00.000000Z"}]}}`)
				 }))
			 })
			 It(`Invoke GetInstanceRulesetVersionByTag successfully with retries`, func() {
				 rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
					 Crn: core.StringPtr(crn),
					 ZoneIdentifier: core.StringPtr(zoneIdentifier),
				 })
				 Expect(serviceErr).To(BeNil())
				 Expect(rulesetsService).ToNot(BeNil())
				 rulesetsService.EnableRetries(0, 0)
 
				 // Construct an instance of the GetInstanceRulesetVersionByTagOptions model
				 getInstanceRulesetVersionByTagOptionsModel := new(rulesetsv1.GetInstanceRulesetVersionByTagOptions)
				 getInstanceRulesetVersionByTagOptionsModel.RulesetID = core.StringPtr("testString")
				 getInstanceRulesetVersionByTagOptionsModel.RulesetVersion = core.StringPtr("1")
				 getInstanceRulesetVersionByTagOptionsModel.RuleTag = core.StringPtr("testString")
				 getInstanceRulesetVersionByTagOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
 
				 // Invoke operation with a Context to test a timeout error
				 ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				 defer cancelFunc()
				 _, _, operationErr := rulesetsService.GetInstanceRulesetVersionByTagWithContext(ctx, getInstanceRulesetVersionByTagOptionsModel)
				 Expect(operationErr).ToNot(BeNil())
				 Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
 
				 // Disable retries and test again
				 rulesetsService.DisableRetries()
				 result, response, operationErr := rulesetsService.GetInstanceRulesetVersionByTag(getInstanceRulesetVersionByTagOptionsModel)
				 Expect(operationErr).To(BeNil())
				 Expect(response).ToNot(BeNil())
				 Expect(result).ToNot(BeNil())
 
				 // Re-test the timeout error with retries disabled
				 ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				 defer cancelFunc2()
				 _, _, operationErr = rulesetsService.GetInstanceRulesetVersionByTagWithContext(ctx, getInstanceRulesetVersionByTagOptionsModel)
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
					 Expect(req.URL.EscapedPath()).To(Equal(getInstanceRulesetVersionByTagPath))
					 Expect(req.Method).To(Equal("GET"))
 
					 // Set mock response
					 res.Header().Set("Content-type", "application/json")
					 res.WriteHeader(200)
					 fmt.Fprintf(res, "%s", `{"success": true, "errors": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "messages": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "result": {"description": "Custom instance ruleset", "id": "ID", "kind": "managed", "last_updated": "2000-01-01T00:00:00.000000Z", "name": "Name", "phase": "ddos_l4", "version": "1", "rules": [{"id": "ID", "version": "Version", "action": "Action", "action_parameters": {"id": "ID", "overrides": {"action": "Action", "enabled": false, "sensitivity_level": "high", "rules": [{"id": "ID", "enabled": false, "action": "Action", "sensitivity_level": "high"}], "categories": [{"category": "Category", "enabled": false, "action": "Action"}]}, "version": "Version", "ruleset": "Ruleset", "rulesets": ["Rulesets"], "response": {"content": "{\"success\": false, \"error\": \"you have been blocked\"}", "content_type": "application/json", "status_code": 400}}, "categories": ["Categories"], "enabled": true, "description": "Description", "expression": "ip.src ne 1.1.1.1", "ref": "my_ref", "logging": {"enabled": true}, "last_updated": "2000-01-01T00:00:00.000000Z"}]}}`)
				 }))
			 })
			 It(`Invoke GetInstanceRulesetVersionByTag successfully`, func() {
				 rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
					 Crn: core.StringPtr(crn),
					 ZoneIdentifier: core.StringPtr(zoneIdentifier),
				 })
				 Expect(serviceErr).To(BeNil())
				 Expect(rulesetsService).ToNot(BeNil())
 
				 // Invoke operation with nil options model (negative test)
				 result, response, operationErr := rulesetsService.GetInstanceRulesetVersionByTag(nil)
				 Expect(operationErr).NotTo(BeNil())
				 Expect(response).To(BeNil())
				 Expect(result).To(BeNil())
 
				 // Construct an instance of the GetInstanceRulesetVersionByTagOptions model
				 getInstanceRulesetVersionByTagOptionsModel := new(rulesetsv1.GetInstanceRulesetVersionByTagOptions)
				 getInstanceRulesetVersionByTagOptionsModel.RulesetID = core.StringPtr("testString")
				 getInstanceRulesetVersionByTagOptionsModel.RulesetVersion = core.StringPtr("1")
				 getInstanceRulesetVersionByTagOptionsModel.RuleTag = core.StringPtr("testString")
				 getInstanceRulesetVersionByTagOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
 
				 // Invoke operation with valid options model (positive test)
				 result, response, operationErr = rulesetsService.GetInstanceRulesetVersionByTag(getInstanceRulesetVersionByTagOptionsModel)
				 Expect(operationErr).To(BeNil())
				 Expect(response).ToNot(BeNil())
				 Expect(result).ToNot(BeNil())
 
			 })
			 It(`Invoke GetInstanceRulesetVersionByTag with error: Operation validation and request error`, func() {
				 rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
					 Crn: core.StringPtr(crn),
					 ZoneIdentifier: core.StringPtr(zoneIdentifier),
				 })
				 Expect(serviceErr).To(BeNil())
				 Expect(rulesetsService).ToNot(BeNil())
 
				 // Construct an instance of the GetInstanceRulesetVersionByTagOptions model
				 getInstanceRulesetVersionByTagOptionsModel := new(rulesetsv1.GetInstanceRulesetVersionByTagOptions)
				 getInstanceRulesetVersionByTagOptionsModel.RulesetID = core.StringPtr("testString")
				 getInstanceRulesetVersionByTagOptionsModel.RulesetVersion = core.StringPtr("1")
				 getInstanceRulesetVersionByTagOptionsModel.RuleTag = core.StringPtr("testString")
				 getInstanceRulesetVersionByTagOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				 // Invoke operation with empty URL (negative test)
				 err := rulesetsService.SetServiceURL("")
				 Expect(err).To(BeNil())
				 result, response, operationErr := rulesetsService.GetInstanceRulesetVersionByTag(getInstanceRulesetVersionByTagOptionsModel)
				 Expect(operationErr).ToNot(BeNil())
				 Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				 Expect(response).To(BeNil())
				 Expect(result).To(BeNil())
				 // Construct a second instance of the GetInstanceRulesetVersionByTagOptions model with no property values
				 getInstanceRulesetVersionByTagOptionsModelNew := new(rulesetsv1.GetInstanceRulesetVersionByTagOptions)
				 // Invoke operation with invalid model (negative test)
				 result, response, operationErr = rulesetsService.GetInstanceRulesetVersionByTag(getInstanceRulesetVersionByTagOptionsModelNew)
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
			 It(`Invoke GetInstanceRulesetVersionByTag successfully`, func() {
				 rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
					 Crn: core.StringPtr(crn),
					 ZoneIdentifier: core.StringPtr(zoneIdentifier),
				 })
				 Expect(serviceErr).To(BeNil())
				 Expect(rulesetsService).ToNot(BeNil())
 
				 // Construct an instance of the GetInstanceRulesetVersionByTagOptions model
				 getInstanceRulesetVersionByTagOptionsModel := new(rulesetsv1.GetInstanceRulesetVersionByTagOptions)
				 getInstanceRulesetVersionByTagOptionsModel.RulesetID = core.StringPtr("testString")
				 getInstanceRulesetVersionByTagOptionsModel.RulesetVersion = core.StringPtr("1")
				 getInstanceRulesetVersionByTagOptionsModel.RuleTag = core.StringPtr("testString")
				 getInstanceRulesetVersionByTagOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
 
				 // Invoke operation
				 result, response, operationErr := rulesetsService.GetInstanceRulesetVersionByTag(getInstanceRulesetVersionByTagOptionsModel)
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
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
					 Crn: core.StringPtr(crn),
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
					 fmt.Fprintf(res, "%s", `{"success": true, "errors": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "messages": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "result": [{"description": "Custom instance ruleset", "id": "ID", "kind": "managed", "last_updated": "2000-01-01T00:00:00.000000Z", "name": "Name", "phase": "ddos_l4", "version": "1"}]}`)
				 }))
			 })
			 It(`Invoke GetZoneRulesets successfully with retries`, func() {
				 rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
					 Crn: core.StringPtr(crn),
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
					 fmt.Fprintf(res, "%s", `{"success": true, "errors": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "messages": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "result": [{"description": "Custom instance ruleset", "id": "ID", "kind": "managed", "last_updated": "2000-01-01T00:00:00.000000Z", "name": "Name", "phase": "ddos_l4", "version": "1"}]}`)
				 }))
			 })
			 It(`Invoke GetZoneRulesets successfully`, func() {
				 rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
					 Crn: core.StringPtr(crn),
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
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
					 Crn: core.StringPtr(crn),
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
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
					 Crn: core.StringPtr(crn),
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
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
					 Crn: core.StringPtr(crn),
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
					 fmt.Fprintf(res, "%s", `{"success": true, "errors": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "messages": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "result": {"description": "Custom instance ruleset", "id": "ID", "kind": "managed", "last_updated": "2000-01-01T00:00:00.000000Z", "name": "Name", "phase": "ddos_l4", "version": "1", "rules": [{"id": "ID", "version": "Version", "action": "Action", "action_parameters": {"id": "ID", "overrides": {"action": "Action", "enabled": false, "sensitivity_level": "high", "rules": [{"id": "ID", "enabled": false, "action": "Action", "sensitivity_level": "high"}], "categories": [{"category": "Category", "enabled": false, "action": "Action"}]}, "version": "Version", "ruleset": "Ruleset", "rulesets": ["Rulesets"], "response": {"content": "{\"success\": false, \"error\": \"you have been blocked\"}", "content_type": "application/json", "status_code": 400}}, "categories": ["Categories"], "enabled": true, "description": "Description", "expression": "ip.src ne 1.1.1.1", "ref": "my_ref", "logging": {"enabled": true}, "last_updated": "2000-01-01T00:00:00.000000Z"}]}}`)
				 }))
			 })
			 It(`Invoke GetZoneRuleset successfully with retries`, func() {
				 rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
					 Crn: core.StringPtr(crn),
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
					 fmt.Fprintf(res, "%s", `{"success": true, "errors": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "messages": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "result": {"description": "Custom instance ruleset", "id": "ID", "kind": "managed", "last_updated": "2000-01-01T00:00:00.000000Z", "name": "Name", "phase": "ddos_l4", "version": "1", "rules": [{"id": "ID", "version": "Version", "action": "Action", "action_parameters": {"id": "ID", "overrides": {"action": "Action", "enabled": false, "sensitivity_level": "high", "rules": [{"id": "ID", "enabled": false, "action": "Action", "sensitivity_level": "high"}], "categories": [{"category": "Category", "enabled": false, "action": "Action"}]}, "version": "Version", "ruleset": "Ruleset", "rulesets": ["Rulesets"], "response": {"content": "{\"success\": false, \"error\": \"you have been blocked\"}", "content_type": "application/json", "status_code": 400}}, "categories": ["Categories"], "enabled": true, "description": "Description", "expression": "ip.src ne 1.1.1.1", "ref": "my_ref", "logging": {"enabled": true}, "last_updated": "2000-01-01T00:00:00.000000Z"}]}}`)
				 }))
			 })
			 It(`Invoke GetZoneRuleset successfully`, func() {
				 rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
					 Crn: core.StringPtr(crn),
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
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
					 Crn: core.StringPtr(crn),
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
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
					 Crn: core.StringPtr(crn),
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
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
					 Crn: core.StringPtr(crn),
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
				 actionParametersResponseModel.Content = core.StringPtr("{\"success\": false, \"error\": \"you have been blocked\"}")
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
				 updateZoneRulesetOptionsModel.Description = core.StringPtr("Custom instance ruleset")
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
					 fmt.Fprintf(res, "%s", `{"success": true, "errors": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "messages": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "result": {"description": "Custom instance ruleset", "id": "ID", "kind": "managed", "last_updated": "2000-01-01T00:00:00.000000Z", "name": "Name", "phase": "ddos_l4", "version": "1", "rules": [{"id": "ID", "version": "Version", "action": "Action", "action_parameters": {"id": "ID", "overrides": {"action": "Action", "enabled": false, "sensitivity_level": "high", "rules": [{"id": "ID", "enabled": false, "action": "Action", "sensitivity_level": "high"}], "categories": [{"category": "Category", "enabled": false, "action": "Action"}]}, "version": "Version", "ruleset": "Ruleset", "rulesets": ["Rulesets"], "response": {"content": "{\"success\": false, \"error\": \"you have been blocked\"}", "content_type": "application/json", "status_code": 400}}, "categories": ["Categories"], "enabled": true, "description": "Description", "expression": "ip.src ne 1.1.1.1", "ref": "my_ref", "logging": {"enabled": true}, "last_updated": "2000-01-01T00:00:00.000000Z"}]}}`)
				 }))
			 })
			 It(`Invoke UpdateZoneRuleset successfully with retries`, func() {
				 rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
					 Crn: core.StringPtr(crn),
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
				 actionParametersResponseModel.Content = core.StringPtr("{\"success\": false, \"error\": \"you have been blocked\"}")
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
				 updateZoneRulesetOptionsModel.Description = core.StringPtr("Custom instance ruleset")
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
					 fmt.Fprintf(res, "%s", `{"success": true, "errors": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "messages": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "result": {"description": "Custom instance ruleset", "id": "ID", "kind": "managed", "last_updated": "2000-01-01T00:00:00.000000Z", "name": "Name", "phase": "ddos_l4", "version": "1", "rules": [{"id": "ID", "version": "Version", "action": "Action", "action_parameters": {"id": "ID", "overrides": {"action": "Action", "enabled": false, "sensitivity_level": "high", "rules": [{"id": "ID", "enabled": false, "action": "Action", "sensitivity_level": "high"}], "categories": [{"category": "Category", "enabled": false, "action": "Action"}]}, "version": "Version", "ruleset": "Ruleset", "rulesets": ["Rulesets"], "response": {"content": "{\"success\": false, \"error\": \"you have been blocked\"}", "content_type": "application/json", "status_code": 400}}, "categories": ["Categories"], "enabled": true, "description": "Description", "expression": "ip.src ne 1.1.1.1", "ref": "my_ref", "logging": {"enabled": true}, "last_updated": "2000-01-01T00:00:00.000000Z"}]}}`)
				 }))
			 })
			 It(`Invoke UpdateZoneRuleset successfully`, func() {
				 rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
					 Crn: core.StringPtr(crn),
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
				 actionParametersResponseModel.Content = core.StringPtr("{\"success\": false, \"error\": \"you have been blocked\"}")
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
				 updateZoneRulesetOptionsModel.Description = core.StringPtr("Custom instance ruleset")
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
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
					 Crn: core.StringPtr(crn),
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
				 actionParametersResponseModel.Content = core.StringPtr("{\"success\": false, \"error\": \"you have been blocked\"}")
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
				 updateZoneRulesetOptionsModel.Description = core.StringPtr("Custom instance ruleset")
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
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
					 Crn: core.StringPtr(crn),
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
				 actionParametersResponseModel.Content = core.StringPtr("{\"success\": false, \"error\": \"you have been blocked\"}")
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
				 updateZoneRulesetOptionsModel.Description = core.StringPtr("Custom instance ruleset")
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
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
					 Crn: core.StringPtr(crn),
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
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
					 Crn: core.StringPtr(crn),
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
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
					 Crn: core.StringPtr(crn),
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
					 fmt.Fprintf(res, "%s", `{"success": true, "errors": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "messages": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "result": [{"description": "Custom instance ruleset", "id": "ID", "kind": "managed", "last_updated": "2000-01-01T00:00:00.000000Z", "name": "Name", "phase": "ddos_l4", "version": "1"}]}`)
				 }))
			 })
			 It(`Invoke GetZoneRulesetVersions successfully with retries`, func() {
				 rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
					 Crn: core.StringPtr(crn),
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
					 fmt.Fprintf(res, "%s", `{"success": true, "errors": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "messages": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "result": [{"description": "Custom instance ruleset", "id": "ID", "kind": "managed", "last_updated": "2000-01-01T00:00:00.000000Z", "name": "Name", "phase": "ddos_l4", "version": "1"}]}`)
				 }))
			 })
			 It(`Invoke GetZoneRulesetVersions successfully`, func() {
				 rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
					 Crn: core.StringPtr(crn),
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
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
					 Crn: core.StringPtr(crn),
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
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
					 Crn: core.StringPtr(crn),
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
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
					 Crn: core.StringPtr(crn),
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
					 fmt.Fprintf(res, "%s", `{"success": true, "errors": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "messages": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "result": {"description": "Custom instance ruleset", "id": "ID", "kind": "managed", "last_updated": "2000-01-01T00:00:00.000000Z", "name": "Name", "phase": "ddos_l4", "version": "1", "rules": [{"id": "ID", "version": "Version", "action": "Action", "action_parameters": {"id": "ID", "overrides": {"action": "Action", "enabled": false, "sensitivity_level": "high", "rules": [{"id": "ID", "enabled": false, "action": "Action", "sensitivity_level": "high"}], "categories": [{"category": "Category", "enabled": false, "action": "Action"}]}, "version": "Version", "ruleset": "Ruleset", "rulesets": ["Rulesets"], "response": {"content": "{\"success\": false, \"error\": \"you have been blocked\"}", "content_type": "application/json", "status_code": 400}}, "categories": ["Categories"], "enabled": true, "description": "Description", "expression": "ip.src ne 1.1.1.1", "ref": "my_ref", "logging": {"enabled": true}, "last_updated": "2000-01-01T00:00:00.000000Z"}]}}`)
				 }))
			 })
			 It(`Invoke GetZoneRulesetVersion successfully with retries`, func() {
				 rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
					 Crn: core.StringPtr(crn),
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
					 fmt.Fprintf(res, "%s", `{"success": true, "errors": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "messages": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "result": {"description": "Custom instance ruleset", "id": "ID", "kind": "managed", "last_updated": "2000-01-01T00:00:00.000000Z", "name": "Name", "phase": "ddos_l4", "version": "1", "rules": [{"id": "ID", "version": "Version", "action": "Action", "action_parameters": {"id": "ID", "overrides": {"action": "Action", "enabled": false, "sensitivity_level": "high", "rules": [{"id": "ID", "enabled": false, "action": "Action", "sensitivity_level": "high"}], "categories": [{"category": "Category", "enabled": false, "action": "Action"}]}, "version": "Version", "ruleset": "Ruleset", "rulesets": ["Rulesets"], "response": {"content": "{\"success\": false, \"error\": \"you have been blocked\"}", "content_type": "application/json", "status_code": 400}}, "categories": ["Categories"], "enabled": true, "description": "Description", "expression": "ip.src ne 1.1.1.1", "ref": "my_ref", "logging": {"enabled": true}, "last_updated": "2000-01-01T00:00:00.000000Z"}]}}`)
				 }))
			 })
			 It(`Invoke GetZoneRulesetVersion successfully`, func() {
				 rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
					 Crn: core.StringPtr(crn),
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
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
					 Crn: core.StringPtr(crn),
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
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
					 Crn: core.StringPtr(crn),
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
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
					 Crn: core.StringPtr(crn),
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
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
					 Crn: core.StringPtr(crn),
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
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
					 Crn: core.StringPtr(crn),
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
					 fmt.Fprintf(res, "%s", `{"success": true, "errors": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "messages": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "result": {"description": "Custom instance ruleset", "id": "ID", "kind": "managed", "last_updated": "2000-01-01T00:00:00.000000Z", "name": "Name", "phase": "ddos_l4", "version": "1", "rules": [{"id": "ID", "version": "Version", "action": "Action", "action_parameters": {"id": "ID", "overrides": {"action": "Action", "enabled": false, "sensitivity_level": "high", "rules": [{"id": "ID", "enabled": false, "action": "Action", "sensitivity_level": "high"}], "categories": [{"category": "Category", "enabled": false, "action": "Action"}]}, "version": "Version", "ruleset": "Ruleset", "rulesets": ["Rulesets"], "response": {"content": "{\"success\": false, \"error\": \"you have been blocked\"}", "content_type": "application/json", "status_code": 400}}, "categories": ["Categories"], "enabled": true, "description": "Description", "expression": "ip.src ne 1.1.1.1", "ref": "my_ref", "logging": {"enabled": true}, "last_updated": "2000-01-01T00:00:00.000000Z"}]}}`)
				 }))
			 })
			 It(`Invoke GetZoneEntrypointRuleset successfully with retries`, func() {
				 rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
					 Crn: core.StringPtr(crn),
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
					 fmt.Fprintf(res, "%s", `{"success": true, "errors": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "messages": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "result": {"description": "Custom instance ruleset", "id": "ID", "kind": "managed", "last_updated": "2000-01-01T00:00:00.000000Z", "name": "Name", "phase": "ddos_l4", "version": "1", "rules": [{"id": "ID", "version": "Version", "action": "Action", "action_parameters": {"id": "ID", "overrides": {"action": "Action", "enabled": false, "sensitivity_level": "high", "rules": [{"id": "ID", "enabled": false, "action": "Action", "sensitivity_level": "high"}], "categories": [{"category": "Category", "enabled": false, "action": "Action"}]}, "version": "Version", "ruleset": "Ruleset", "rulesets": ["Rulesets"], "response": {"content": "{\"success\": false, \"error\": \"you have been blocked\"}", "content_type": "application/json", "status_code": 400}}, "categories": ["Categories"], "enabled": true, "description": "Description", "expression": "ip.src ne 1.1.1.1", "ref": "my_ref", "logging": {"enabled": true}, "last_updated": "2000-01-01T00:00:00.000000Z"}]}}`)
				 }))
			 })
			 It(`Invoke GetZoneEntrypointRuleset successfully`, func() {
				 rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
					 Crn: core.StringPtr(crn),
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
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
					 Crn: core.StringPtr(crn),
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
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
					 Crn: core.StringPtr(crn),
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
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
					 Crn: core.StringPtr(crn),
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
				 actionParametersResponseModel.Content = core.StringPtr("{\"success\": false, \"error\": \"you have been blocked\"}")
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
				 updateZoneEntrypointRulesetOptionsModel.Description = core.StringPtr("Custom instance ruleset")
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
					 fmt.Fprintf(res, "%s", `{"success": true, "errors": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "messages": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "result": {"description": "Custom instance ruleset", "id": "ID", "kind": "managed", "last_updated": "2000-01-01T00:00:00.000000Z", "name": "Name", "phase": "ddos_l4", "version": "1", "rules": [{"id": "ID", "version": "Version", "action": "Action", "action_parameters": {"id": "ID", "overrides": {"action": "Action", "enabled": false, "sensitivity_level": "high", "rules": [{"id": "ID", "enabled": false, "action": "Action", "sensitivity_level": "high"}], "categories": [{"category": "Category", "enabled": false, "action": "Action"}]}, "version": "Version", "ruleset": "Ruleset", "rulesets": ["Rulesets"], "response": {"content": "{\"success\": false, \"error\": \"you have been blocked\"}", "content_type": "application/json", "status_code": 400}}, "categories": ["Categories"], "enabled": true, "description": "Description", "expression": "ip.src ne 1.1.1.1", "ref": "my_ref", "logging": {"enabled": true}, "last_updated": "2000-01-01T00:00:00.000000Z"}]}}`)
				 }))
			 })
			 It(`Invoke UpdateZoneEntrypointRuleset successfully with retries`, func() {
				 rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
					 Crn: core.StringPtr(crn),
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
				 actionParametersResponseModel.Content = core.StringPtr("{\"success\": false, \"error\": \"you have been blocked\"}")
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
				 updateZoneEntrypointRulesetOptionsModel.Description = core.StringPtr("Custom instance ruleset")
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
					 fmt.Fprintf(res, "%s", `{"success": true, "errors": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "messages": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "result": {"description": "Custom instance ruleset", "id": "ID", "kind": "managed", "last_updated": "2000-01-01T00:00:00.000000Z", "name": "Name", "phase": "ddos_l4", "version": "1", "rules": [{"id": "ID", "version": "Version", "action": "Action", "action_parameters": {"id": "ID", "overrides": {"action": "Action", "enabled": false, "sensitivity_level": "high", "rules": [{"id": "ID", "enabled": false, "action": "Action", "sensitivity_level": "high"}], "categories": [{"category": "Category", "enabled": false, "action": "Action"}]}, "version": "Version", "ruleset": "Ruleset", "rulesets": ["Rulesets"], "response": {"content": "{\"success\": false, \"error\": \"you have been blocked\"}", "content_type": "application/json", "status_code": 400}}, "categories": ["Categories"], "enabled": true, "description": "Description", "expression": "ip.src ne 1.1.1.1", "ref": "my_ref", "logging": {"enabled": true}, "last_updated": "2000-01-01T00:00:00.000000Z"}]}}`)
				 }))
			 })
			 It(`Invoke UpdateZoneEntrypointRuleset successfully`, func() {
				 rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
					 Crn: core.StringPtr(crn),
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
				 actionParametersResponseModel.Content = core.StringPtr("{\"success\": false, \"error\": \"you have been blocked\"}")
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
				 updateZoneEntrypointRulesetOptionsModel.Description = core.StringPtr("Custom instance ruleset")
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
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
					 Crn: core.StringPtr(crn),
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
				 actionParametersResponseModel.Content = core.StringPtr("{\"success\": false, \"error\": \"you have been blocked\"}")
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
				 updateZoneEntrypointRulesetOptionsModel.Description = core.StringPtr("Custom instance ruleset")
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
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
					 Crn: core.StringPtr(crn),
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
				 actionParametersResponseModel.Content = core.StringPtr("{\"success\": false, \"error\": \"you have been blocked\"}")
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
				 updateZoneEntrypointRulesetOptionsModel.Description = core.StringPtr("Custom instance ruleset")
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
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
					 Crn: core.StringPtr(crn),
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
					 fmt.Fprintf(res, "%s", `{"success": true, "errors": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "messages": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "result": [{"description": "Custom instance ruleset", "id": "ID", "kind": "managed", "last_updated": "2000-01-01T00:00:00.000000Z", "name": "Name", "phase": "ddos_l4", "version": "1"}]}`)
				 }))
			 })
			 It(`Invoke GetZoneEntryPointRulesetVersions successfully with retries`, func() {
				 rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
					 Crn: core.StringPtr(crn),
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
					 fmt.Fprintf(res, "%s", `{"success": true, "errors": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "messages": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "result": [{"description": "Custom instance ruleset", "id": "ID", "kind": "managed", "last_updated": "2000-01-01T00:00:00.000000Z", "name": "Name", "phase": "ddos_l4", "version": "1"}]}`)
				 }))
			 })
			 It(`Invoke GetZoneEntryPointRulesetVersions successfully`, func() {
				 rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
					 Crn: core.StringPtr(crn),
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
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
					 Crn: core.StringPtr(crn),
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
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
					 Crn: core.StringPtr(crn),
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
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
					 Crn: core.StringPtr(crn),
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
					 fmt.Fprintf(res, "%s", `{"success": true, "errors": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "messages": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "result": {"description": "Custom instance ruleset", "id": "ID", "kind": "managed", "last_updated": "2000-01-01T00:00:00.000000Z", "name": "Name", "phase": "ddos_l4", "version": "1", "rules": [{"id": "ID", "version": "Version", "action": "Action", "action_parameters": {"id": "ID", "overrides": {"action": "Action", "enabled": false, "sensitivity_level": "high", "rules": [{"id": "ID", "enabled": false, "action": "Action", "sensitivity_level": "high"}], "categories": [{"category": "Category", "enabled": false, "action": "Action"}]}, "version": "Version", "ruleset": "Ruleset", "rulesets": ["Rulesets"], "response": {"content": "{\"success\": false, \"error\": \"you have been blocked\"}", "content_type": "application/json", "status_code": 400}}, "categories": ["Categories"], "enabled": true, "description": "Description", "expression": "ip.src ne 1.1.1.1", "ref": "my_ref", "logging": {"enabled": true}, "last_updated": "2000-01-01T00:00:00.000000Z"}]}}`)
				 }))
			 })
			 It(`Invoke GetZoneEntryPointRulesetVersion successfully with retries`, func() {
				 rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
					 Crn: core.StringPtr(crn),
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
					 fmt.Fprintf(res, "%s", `{"success": true, "errors": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "messages": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "result": {"description": "Custom instance ruleset", "id": "ID", "kind": "managed", "last_updated": "2000-01-01T00:00:00.000000Z", "name": "Name", "phase": "ddos_l4", "version": "1", "rules": [{"id": "ID", "version": "Version", "action": "Action", "action_parameters": {"id": "ID", "overrides": {"action": "Action", "enabled": false, "sensitivity_level": "high", "rules": [{"id": "ID", "enabled": false, "action": "Action", "sensitivity_level": "high"}], "categories": [{"category": "Category", "enabled": false, "action": "Action"}]}, "version": "Version", "ruleset": "Ruleset", "rulesets": ["Rulesets"], "response": {"content": "{\"success\": false, \"error\": \"you have been blocked\"}", "content_type": "application/json", "status_code": 400}}, "categories": ["Categories"], "enabled": true, "description": "Description", "expression": "ip.src ne 1.1.1.1", "ref": "my_ref", "logging": {"enabled": true}, "last_updated": "2000-01-01T00:00:00.000000Z"}]}}`)
				 }))
			 })
			 It(`Invoke GetZoneEntryPointRulesetVersion successfully`, func() {
				 rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
					 Crn: core.StringPtr(crn),
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
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
					 Crn: core.StringPtr(crn),
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
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
					 Crn: core.StringPtr(crn),
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
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
					 Crn: core.StringPtr(crn),
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
				 actionParametersResponseModel.Content = core.StringPtr("{\"success\": false, \"error\": \"you have been blocked\"}")
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
					 fmt.Fprintf(res, "%s", `{"success": true, "errors": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "messages": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "result": {"id": "ID", "version": "Version", "action": "Action", "action_parameters": {"id": "ID", "overrides": {"action": "Action", "enabled": false, "sensitivity_level": "high", "rules": [{"id": "ID", "enabled": false, "action": "Action", "sensitivity_level": "high"}], "categories": [{"category": "Category", "enabled": false, "action": "Action"}]}, "version": "Version", "ruleset": "Ruleset", "rulesets": ["Rulesets"], "response": {"content": "{\"success\": false, \"error\": \"you have been blocked\"}", "content_type": "application/json", "status_code": 400}}, "categories": ["Categories"], "enabled": true, "description": "Description", "expression": "ip.src ne 1.1.1.1", "ref": "my_ref", "logging": {"enabled": true}, "last_updated": "2000-01-01T00:00:00.000000Z"}}`)
				 }))
			 })
			 It(`Invoke CreateZoneRulesetRule successfully with retries`, func() {
				 rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
					 Crn: core.StringPtr(crn),
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
				 actionParametersResponseModel.Content = core.StringPtr("{\"success\": false, \"error\": \"you have been blocked\"}")
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
					 fmt.Fprintf(res, "%s", `{"success": true, "errors": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "messages": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "result": {"id": "ID", "version": "Version", "action": "Action", "action_parameters": {"id": "ID", "overrides": {"action": "Action", "enabled": false, "sensitivity_level": "high", "rules": [{"id": "ID", "enabled": false, "action": "Action", "sensitivity_level": "high"}], "categories": [{"category": "Category", "enabled": false, "action": "Action"}]}, "version": "Version", "ruleset": "Ruleset", "rulesets": ["Rulesets"], "response": {"content": "{\"success\": false, \"error\": \"you have been blocked\"}", "content_type": "application/json", "status_code": 400}}, "categories": ["Categories"], "enabled": true, "description": "Description", "expression": "ip.src ne 1.1.1.1", "ref": "my_ref", "logging": {"enabled": true}, "last_updated": "2000-01-01T00:00:00.000000Z"}}`)
				 }))
			 })
			 It(`Invoke CreateZoneRulesetRule successfully`, func() {
				 rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
					 Crn: core.StringPtr(crn),
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
				 actionParametersResponseModel.Content = core.StringPtr("{\"success\": false, \"error\": \"you have been blocked\"}")
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
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
					 Crn: core.StringPtr(crn),
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
				 actionParametersResponseModel.Content = core.StringPtr("{\"success\": false, \"error\": \"you have been blocked\"}")
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
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
					 Crn: core.StringPtr(crn),
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
				 actionParametersResponseModel.Content = core.StringPtr("{\"success\": false, \"error\": \"you have been blocked\"}")
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
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
					 Crn: core.StringPtr(crn),
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
				 actionParametersResponseModel.Content = core.StringPtr("{\"success\": false, \"error\": \"you have been blocked\"}")
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
					 fmt.Fprintf(res, "%s", `{"success": true, "errors": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "messages": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "result": {"id": "ID", "version": "Version", "action": "Action", "action_parameters": {"id": "ID", "overrides": {"action": "Action", "enabled": false, "sensitivity_level": "high", "rules": [{"id": "ID", "enabled": false, "action": "Action", "sensitivity_level": "high"}], "categories": [{"category": "Category", "enabled": false, "action": "Action"}]}, "version": "Version", "ruleset": "Ruleset", "rulesets": ["Rulesets"], "response": {"content": "{\"success\": false, \"error\": \"you have been blocked\"}", "content_type": "application/json", "status_code": 400}}, "categories": ["Categories"], "enabled": true, "description": "Description", "expression": "ip.src ne 1.1.1.1", "ref": "my_ref", "logging": {"enabled": true}, "last_updated": "2000-01-01T00:00:00.000000Z"}}`)
				 }))
			 })
			 It(`Invoke UpdateZoneRulesetRule successfully with retries`, func() {
				 rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
					 Crn: core.StringPtr(crn),
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
				 actionParametersResponseModel.Content = core.StringPtr("{\"success\": false, \"error\": \"you have been blocked\"}")
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
					 fmt.Fprintf(res, "%s", `{"success": true, "errors": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "messages": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "result": {"id": "ID", "version": "Version", "action": "Action", "action_parameters": {"id": "ID", "overrides": {"action": "Action", "enabled": false, "sensitivity_level": "high", "rules": [{"id": "ID", "enabled": false, "action": "Action", "sensitivity_level": "high"}], "categories": [{"category": "Category", "enabled": false, "action": "Action"}]}, "version": "Version", "ruleset": "Ruleset", "rulesets": ["Rulesets"], "response": {"content": "{\"success\": false, \"error\": \"you have been blocked\"}", "content_type": "application/json", "status_code": 400}}, "categories": ["Categories"], "enabled": true, "description": "Description", "expression": "ip.src ne 1.1.1.1", "ref": "my_ref", "logging": {"enabled": true}, "last_updated": "2000-01-01T00:00:00.000000Z"}}`)
				 }))
			 })
			 It(`Invoke UpdateZoneRulesetRule successfully`, func() {
				 rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
					 Crn: core.StringPtr(crn),
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
				 actionParametersResponseModel.Content = core.StringPtr("{\"success\": false, \"error\": \"you have been blocked\"}")
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
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
					 Crn: core.StringPtr(crn),
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
				 actionParametersResponseModel.Content = core.StringPtr("{\"success\": false, \"error\": \"you have been blocked\"}")
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
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
					 Crn: core.StringPtr(crn),
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
				 actionParametersResponseModel.Content = core.StringPtr("{\"success\": false, \"error\": \"you have been blocked\"}")
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
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
					 Crn: core.StringPtr(crn),
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
					 fmt.Fprintf(res, "%s", `{"success": true, "errors": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "messages": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "result": {"id": "ID", "version": "Version", "action": "Action", "action_parameters": {"id": "ID", "overrides": {"action": "Action", "enabled": false, "sensitivity_level": "high", "rules": [{"id": "ID", "enabled": false, "action": "Action", "sensitivity_level": "high"}], "categories": [{"category": "Category", "enabled": false, "action": "Action"}]}, "version": "Version", "ruleset": "Ruleset", "rulesets": ["Rulesets"], "response": {"content": "{\"success\": false, \"error\": \"you have been blocked\"}", "content_type": "application/json", "status_code": 400}}, "categories": ["Categories"], "enabled": true, "description": "Description", "expression": "ip.src ne 1.1.1.1", "ref": "my_ref", "logging": {"enabled": true}, "last_updated": "2000-01-01T00:00:00.000000Z"}}`)
				 }))
			 })
			 It(`Invoke DeleteZoneRulesetRule successfully with retries`, func() {
				 rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
					 Crn: core.StringPtr(crn),
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
					 fmt.Fprintf(res, "%s", `{"success": true, "errors": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "messages": [{"code": 10000, "message": "something failed in the request", "source": {"pointer": "/rules/0/action"}}], "result": {"id": "ID", "version": "Version", "action": "Action", "action_parameters": {"id": "ID", "overrides": {"action": "Action", "enabled": false, "sensitivity_level": "high", "rules": [{"id": "ID", "enabled": false, "action": "Action", "sensitivity_level": "high"}], "categories": [{"category": "Category", "enabled": false, "action": "Action"}]}, "version": "Version", "ruleset": "Ruleset", "rulesets": ["Rulesets"], "response": {"content": "{\"success\": false, \"error\": \"you have been blocked\"}", "content_type": "application/json", "status_code": 400}}, "categories": ["Categories"], "enabled": true, "description": "Description", "expression": "ip.src ne 1.1.1.1", "ref": "my_ref", "logging": {"enabled": true}, "last_updated": "2000-01-01T00:00:00.000000Z"}}`)
				 }))
			 })
			 It(`Invoke DeleteZoneRulesetRule successfully`, func() {
				 rulesetsService, serviceErr := rulesetsv1.NewRulesetsV1(&rulesetsv1.RulesetsV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
					 Crn: core.StringPtr(crn),
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
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
					 Crn: core.StringPtr(crn),
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
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
					 Crn: core.StringPtr(crn),
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
				 URL:           "http://rulesetsv1modelgenerator.com",
				 Authenticator: &core.NoAuthAuthenticator{},
				 Crn: core.StringPtr(crn),
				 ZoneIdentifier: core.StringPtr(zoneIdentifier),
			 })
			 It(`Invoke NewActionParametersResponse successfully`, func() {
				 content := "{\"success\": false, \"error\": \"you have been blocked\"}"
				 contentType := "application/json"
				 statusCode := int64(400)
				 _model, err := rulesetsService.NewActionParametersResponse(content, contentType, statusCode)
				 Expect(_model).ToNot(BeNil())
				 Expect(err).To(BeNil())
			 })
			 It(`Invoke NewCreateInstanceRulesetRuleOptions successfully`, func() {
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
				 actionParametersResponseModel.Content = core.StringPtr("{\"success\": false, \"error\": \"you have been blocked\"}")
				 actionParametersResponseModel.ContentType = core.StringPtr("application/json")
				 actionParametersResponseModel.StatusCode = core.Int64Ptr(int64(400))
				 Expect(actionParametersResponseModel.Content).To(Equal(core.StringPtr("{\"success\": false, \"error\": \"you have been blocked\"}")))
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
 
				 // Construct an instance of the CreateInstanceRulesetRuleOptions model
				 rulesetID := "testString"
				 createInstanceRulesetRuleOptionsModel := rulesetsService.NewCreateInstanceRulesetRuleOptions(rulesetID)
				 createInstanceRulesetRuleOptionsModel.SetRulesetID("testString")
				 createInstanceRulesetRuleOptionsModel.SetAction("testString")
				 createInstanceRulesetRuleOptionsModel.SetActionParameters(actionParametersModel)
				 createInstanceRulesetRuleOptionsModel.SetDescription("testString")
				 createInstanceRulesetRuleOptionsModel.SetEnabled(true)
				 createInstanceRulesetRuleOptionsModel.SetExpression("ip.src ne 1.1.1.1")
				 createInstanceRulesetRuleOptionsModel.SetID("testString")
				 createInstanceRulesetRuleOptionsModel.SetLogging(loggingModel)
				 createInstanceRulesetRuleOptionsModel.SetRef("my_ref")
				 createInstanceRulesetRuleOptionsModel.SetPosition(positionModel)
				 createInstanceRulesetRuleOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				 Expect(createInstanceRulesetRuleOptionsModel).ToNot(BeNil())
				 Expect(createInstanceRulesetRuleOptionsModel.RulesetID).To(Equal(core.StringPtr("testString")))
				 Expect(createInstanceRulesetRuleOptionsModel.Action).To(Equal(core.StringPtr("testString")))
				 Expect(createInstanceRulesetRuleOptionsModel.ActionParameters).To(Equal(actionParametersModel))
				 Expect(createInstanceRulesetRuleOptionsModel.Description).To(Equal(core.StringPtr("testString")))
				 Expect(createInstanceRulesetRuleOptionsModel.Enabled).To(Equal(core.BoolPtr(true)))
				 Expect(createInstanceRulesetRuleOptionsModel.Expression).To(Equal(core.StringPtr("ip.src ne 1.1.1.1")))
				 Expect(createInstanceRulesetRuleOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				 Expect(createInstanceRulesetRuleOptionsModel.Logging).To(Equal(loggingModel))
				 Expect(createInstanceRulesetRuleOptionsModel.Ref).To(Equal(core.StringPtr("my_ref")))
				 Expect(createInstanceRulesetRuleOptionsModel.Position).To(Equal(positionModel))
				 Expect(createInstanceRulesetRuleOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
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
				 actionParametersResponseModel.Content = core.StringPtr("{\"success\": false, \"error\": \"you have been blocked\"}")
				 actionParametersResponseModel.ContentType = core.StringPtr("application/json")
				 actionParametersResponseModel.StatusCode = core.Int64Ptr(int64(400))
				 Expect(actionParametersResponseModel.Content).To(Equal(core.StringPtr("{\"success\": false, \"error\": \"you have been blocked\"}")))
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
			 It(`Invoke NewDeleteInstanceRulesetOptions successfully`, func() {
				 // Construct an instance of the DeleteInstanceRulesetOptions model
				 rulesetID := "testString"
				 deleteInstanceRulesetOptionsModel := rulesetsService.NewDeleteInstanceRulesetOptions(rulesetID)
				 deleteInstanceRulesetOptionsModel.SetRulesetID("testString")
				 deleteInstanceRulesetOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				 Expect(deleteInstanceRulesetOptionsModel).ToNot(BeNil())
				 Expect(deleteInstanceRulesetOptionsModel.RulesetID).To(Equal(core.StringPtr("testString")))
				 Expect(deleteInstanceRulesetOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			 })
			 It(`Invoke NewDeleteInstanceRulesetRuleOptions successfully`, func() {
				 // Construct an instance of the DeleteInstanceRulesetRuleOptions model
				 rulesetID := "testString"
				 ruleID := "testString"
				 deleteInstanceRulesetRuleOptionsModel := rulesetsService.NewDeleteInstanceRulesetRuleOptions(rulesetID, ruleID)
				 deleteInstanceRulesetRuleOptionsModel.SetRulesetID("testString")
				 deleteInstanceRulesetRuleOptionsModel.SetRuleID("testString")
				 deleteInstanceRulesetRuleOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				 Expect(deleteInstanceRulesetRuleOptionsModel).ToNot(BeNil())
				 Expect(deleteInstanceRulesetRuleOptionsModel.RulesetID).To(Equal(core.StringPtr("testString")))
				 Expect(deleteInstanceRulesetRuleOptionsModel.RuleID).To(Equal(core.StringPtr("testString")))
				 Expect(deleteInstanceRulesetRuleOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			 })
			 It(`Invoke NewDeleteInstanceRulesetVersionOptions successfully`, func() {
				 // Construct an instance of the DeleteInstanceRulesetVersionOptions model
				 rulesetID := "testString"
				 rulesetVersion := "1"
				 deleteInstanceRulesetVersionOptionsModel := rulesetsService.NewDeleteInstanceRulesetVersionOptions(rulesetID, rulesetVersion)
				 deleteInstanceRulesetVersionOptionsModel.SetRulesetID("testString")
				 deleteInstanceRulesetVersionOptionsModel.SetRulesetVersion("1")
				 deleteInstanceRulesetVersionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				 Expect(deleteInstanceRulesetVersionOptionsModel).ToNot(BeNil())
				 Expect(deleteInstanceRulesetVersionOptionsModel.RulesetID).To(Equal(core.StringPtr("testString")))
				 Expect(deleteInstanceRulesetVersionOptionsModel.RulesetVersion).To(Equal(core.StringPtr("1")))
				 Expect(deleteInstanceRulesetVersionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
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
			 It(`Invoke NewGetInstanceEntryPointRulesetVersionOptions successfully`, func() {
				 // Construct an instance of the GetInstanceEntryPointRulesetVersionOptions model
				 rulesetPhase := "ddos_l4"
				 rulesetVersion := "1"
				 getInstanceEntryPointRulesetVersionOptionsModel := rulesetsService.NewGetInstanceEntryPointRulesetVersionOptions(rulesetPhase, rulesetVersion)
				 getInstanceEntryPointRulesetVersionOptionsModel.SetRulesetPhase("ddos_l4")
				 getInstanceEntryPointRulesetVersionOptionsModel.SetRulesetVersion("1")
				 getInstanceEntryPointRulesetVersionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				 Expect(getInstanceEntryPointRulesetVersionOptionsModel).ToNot(BeNil())
				 Expect(getInstanceEntryPointRulesetVersionOptionsModel.RulesetPhase).To(Equal(core.StringPtr("ddos_l4")))
				 Expect(getInstanceEntryPointRulesetVersionOptionsModel.RulesetVersion).To(Equal(core.StringPtr("1")))
				 Expect(getInstanceEntryPointRulesetVersionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			 })
			 It(`Invoke NewGetInstanceEntryPointRulesetVersionsOptions successfully`, func() {
				 // Construct an instance of the GetInstanceEntryPointRulesetVersionsOptions model
				 rulesetPhase := "ddos_l4"
				 getInstanceEntryPointRulesetVersionsOptionsModel := rulesetsService.NewGetInstanceEntryPointRulesetVersionsOptions(rulesetPhase)
				 getInstanceEntryPointRulesetVersionsOptionsModel.SetRulesetPhase("ddos_l4")
				 getInstanceEntryPointRulesetVersionsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				 Expect(getInstanceEntryPointRulesetVersionsOptionsModel).ToNot(BeNil())
				 Expect(getInstanceEntryPointRulesetVersionsOptionsModel.RulesetPhase).To(Equal(core.StringPtr("ddos_l4")))
				 Expect(getInstanceEntryPointRulesetVersionsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			 })
			 It(`Invoke NewGetInstanceEntrypointRulesetOptions successfully`, func() {
				 // Construct an instance of the GetInstanceEntrypointRulesetOptions model
				 rulesetPhase := "ddos_l4"
				 getInstanceEntrypointRulesetOptionsModel := rulesetsService.NewGetInstanceEntrypointRulesetOptions(rulesetPhase)
				 getInstanceEntrypointRulesetOptionsModel.SetRulesetPhase("ddos_l4")
				 getInstanceEntrypointRulesetOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				 Expect(getInstanceEntrypointRulesetOptionsModel).ToNot(BeNil())
				 Expect(getInstanceEntrypointRulesetOptionsModel.RulesetPhase).To(Equal(core.StringPtr("ddos_l4")))
				 Expect(getInstanceEntrypointRulesetOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			 })
			 It(`Invoke NewGetInstanceRulesetOptions successfully`, func() {
				 // Construct an instance of the GetInstanceRulesetOptions model
				 rulesetID := "testString"
				 getInstanceRulesetOptionsModel := rulesetsService.NewGetInstanceRulesetOptions(rulesetID)
				 getInstanceRulesetOptionsModel.SetRulesetID("testString")
				 getInstanceRulesetOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				 Expect(getInstanceRulesetOptionsModel).ToNot(BeNil())
				 Expect(getInstanceRulesetOptionsModel.RulesetID).To(Equal(core.StringPtr("testString")))
				 Expect(getInstanceRulesetOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			 })
			 It(`Invoke NewGetInstanceRulesetVersionByTagOptions successfully`, func() {
				 // Construct an instance of the GetInstanceRulesetVersionByTagOptions model
				 rulesetID := "testString"
				 rulesetVersion := "1"
				 ruleTag := "testString"
				 getInstanceRulesetVersionByTagOptionsModel := rulesetsService.NewGetInstanceRulesetVersionByTagOptions(rulesetID, rulesetVersion, ruleTag)
				 getInstanceRulesetVersionByTagOptionsModel.SetRulesetID("testString")
				 getInstanceRulesetVersionByTagOptionsModel.SetRulesetVersion("1")
				 getInstanceRulesetVersionByTagOptionsModel.SetRuleTag("testString")
				 getInstanceRulesetVersionByTagOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				 Expect(getInstanceRulesetVersionByTagOptionsModel).ToNot(BeNil())
				 Expect(getInstanceRulesetVersionByTagOptionsModel.RulesetID).To(Equal(core.StringPtr("testString")))
				 Expect(getInstanceRulesetVersionByTagOptionsModel.RulesetVersion).To(Equal(core.StringPtr("1")))
				 Expect(getInstanceRulesetVersionByTagOptionsModel.RuleTag).To(Equal(core.StringPtr("testString")))
				 Expect(getInstanceRulesetVersionByTagOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			 })
			 It(`Invoke NewGetInstanceRulesetVersionOptions successfully`, func() {
				 // Construct an instance of the GetInstanceRulesetVersionOptions model
				 rulesetID := "testString"
				 rulesetVersion := "1"
				 getInstanceRulesetVersionOptionsModel := rulesetsService.NewGetInstanceRulesetVersionOptions(rulesetID, rulesetVersion)
				 getInstanceRulesetVersionOptionsModel.SetRulesetID("testString")
				 getInstanceRulesetVersionOptionsModel.SetRulesetVersion("1")
				 getInstanceRulesetVersionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				 Expect(getInstanceRulesetVersionOptionsModel).ToNot(BeNil())
				 Expect(getInstanceRulesetVersionOptionsModel.RulesetID).To(Equal(core.StringPtr("testString")))
				 Expect(getInstanceRulesetVersionOptionsModel.RulesetVersion).To(Equal(core.StringPtr("1")))
				 Expect(getInstanceRulesetVersionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			 })
			 It(`Invoke NewGetInstanceRulesetVersionsOptions successfully`, func() {
				 // Construct an instance of the GetInstanceRulesetVersionsOptions model
				 rulesetID := "testString"
				 getInstanceRulesetVersionsOptionsModel := rulesetsService.NewGetInstanceRulesetVersionsOptions(rulesetID)
				 getInstanceRulesetVersionsOptionsModel.SetRulesetID("testString")
				 getInstanceRulesetVersionsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				 Expect(getInstanceRulesetVersionsOptionsModel).ToNot(BeNil())
				 Expect(getInstanceRulesetVersionsOptionsModel.RulesetID).To(Equal(core.StringPtr("testString")))
				 Expect(getInstanceRulesetVersionsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			 })
			 It(`Invoke NewGetInstanceRulesetsOptions successfully`, func() {
				 // Construct an instance of the GetInstanceRulesetsOptions model
				 getInstanceRulesetsOptionsModel := rulesetsService.NewGetInstanceRulesetsOptions()
				 getInstanceRulesetsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				 Expect(getInstanceRulesetsOptionsModel).ToNot(BeNil())
				 Expect(getInstanceRulesetsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
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
			 It(`Invoke NewUpdateInstanceEntrypointRulesetOptions successfully`, func() {
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
				 actionParametersResponseModel.Content = core.StringPtr("{\"success\": false, \"error\": \"you have been blocked\"}")
				 actionParametersResponseModel.ContentType = core.StringPtr("application/json")
				 actionParametersResponseModel.StatusCode = core.Int64Ptr(int64(400))
				 Expect(actionParametersResponseModel.Content).To(Equal(core.StringPtr("{\"success\": false, \"error\": \"you have been blocked\"}")))
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
 
				 // Construct an instance of the UpdateInstanceEntrypointRulesetOptions model
				 rulesetPhase := "ddos_l4"
				 updateInstanceEntrypointRulesetOptionsModel := rulesetsService.NewUpdateInstanceEntrypointRulesetOptions(rulesetPhase)
				 updateInstanceEntrypointRulesetOptionsModel.SetRulesetPhase("ddos_l4")
				 updateInstanceEntrypointRulesetOptionsModel.SetDescription("Custom instance ruleset")
				 updateInstanceEntrypointRulesetOptionsModel.SetKind("managed")
				 updateInstanceEntrypointRulesetOptionsModel.SetName("testString")
				 updateInstanceEntrypointRulesetOptionsModel.SetPhase("ddos_l4")
				 updateInstanceEntrypointRulesetOptionsModel.SetRules([]rulesetsv1.RuleCreate{*ruleCreateModel})
				 updateInstanceEntrypointRulesetOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				 Expect(updateInstanceEntrypointRulesetOptionsModel).ToNot(BeNil())
				 Expect(updateInstanceEntrypointRulesetOptionsModel.RulesetPhase).To(Equal(core.StringPtr("ddos_l4")))
				 Expect(updateInstanceEntrypointRulesetOptionsModel.Description).To(Equal(core.StringPtr("Custom instance ruleset")))
				 Expect(updateInstanceEntrypointRulesetOptionsModel.Kind).To(Equal(core.StringPtr("managed")))
				 Expect(updateInstanceEntrypointRulesetOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				 Expect(updateInstanceEntrypointRulesetOptionsModel.Phase).To(Equal(core.StringPtr("ddos_l4")))
				 Expect(updateInstanceEntrypointRulesetOptionsModel.Rules).To(Equal([]rulesetsv1.RuleCreate{*ruleCreateModel}))
				 Expect(updateInstanceEntrypointRulesetOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			 })
			 It(`Invoke NewUpdateInstanceRulesetOptions successfully`, func() {
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
				 actionParametersResponseModel.Content = core.StringPtr("{\"success\": false, \"error\": \"you have been blocked\"}")
				 actionParametersResponseModel.ContentType = core.StringPtr("application/json")
				 actionParametersResponseModel.StatusCode = core.Int64Ptr(int64(400))
				 Expect(actionParametersResponseModel.Content).To(Equal(core.StringPtr("{\"success\": false, \"error\": \"you have been blocked\"}")))
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
 
				 // Construct an instance of the UpdateInstanceRulesetOptions model
				 rulesetID := "testString"
				 updateInstanceRulesetOptionsModel := rulesetsService.NewUpdateInstanceRulesetOptions(rulesetID)
				 updateInstanceRulesetOptionsModel.SetRulesetID("testString")
				 updateInstanceRulesetOptionsModel.SetDescription("Custom instance ruleset")
				 updateInstanceRulesetOptionsModel.SetKind("managed")
				 updateInstanceRulesetOptionsModel.SetName("testString")
				 updateInstanceRulesetOptionsModel.SetPhase("ddos_l4")
				 updateInstanceRulesetOptionsModel.SetRules([]rulesetsv1.RuleCreate{*ruleCreateModel})
				 updateInstanceRulesetOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				 Expect(updateInstanceRulesetOptionsModel).ToNot(BeNil())
				 Expect(updateInstanceRulesetOptionsModel.RulesetID).To(Equal(core.StringPtr("testString")))
				 Expect(updateInstanceRulesetOptionsModel.Description).To(Equal(core.StringPtr("Custom instance ruleset")))
				 Expect(updateInstanceRulesetOptionsModel.Kind).To(Equal(core.StringPtr("managed")))
				 Expect(updateInstanceRulesetOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				 Expect(updateInstanceRulesetOptionsModel.Phase).To(Equal(core.StringPtr("ddos_l4")))
				 Expect(updateInstanceRulesetOptionsModel.Rules).To(Equal([]rulesetsv1.RuleCreate{*ruleCreateModel}))
				 Expect(updateInstanceRulesetOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			 })
			 It(`Invoke NewUpdateInstanceRulesetRuleOptions successfully`, func() {
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
				 actionParametersResponseModel.Content = core.StringPtr("{\"success\": false, \"error\": \"you have been blocked\"}")
				 actionParametersResponseModel.ContentType = core.StringPtr("application/json")
				 actionParametersResponseModel.StatusCode = core.Int64Ptr(int64(400))
				 Expect(actionParametersResponseModel.Content).To(Equal(core.StringPtr("{\"success\": false, \"error\": \"you have been blocked\"}")))
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
 
				 // Construct an instance of the UpdateInstanceRulesetRuleOptions model
				 rulesetID := "testString"
				 ruleID := "testString"
				 updateInstanceRulesetRuleOptionsModel := rulesetsService.NewUpdateInstanceRulesetRuleOptions(rulesetID, ruleID)
				 updateInstanceRulesetRuleOptionsModel.SetRulesetID("testString")
				 updateInstanceRulesetRuleOptionsModel.SetRuleID("testString")
				 updateInstanceRulesetRuleOptionsModel.SetAction("testString")
				 updateInstanceRulesetRuleOptionsModel.SetActionParameters(actionParametersModel)
				 updateInstanceRulesetRuleOptionsModel.SetDescription("testString")
				 updateInstanceRulesetRuleOptionsModel.SetEnabled(true)
				 updateInstanceRulesetRuleOptionsModel.SetExpression("ip.src ne 1.1.1.1")
				 updateInstanceRulesetRuleOptionsModel.SetID("testString")
				 updateInstanceRulesetRuleOptionsModel.SetLogging(loggingModel)
				 updateInstanceRulesetRuleOptionsModel.SetRef("my_ref")
				 updateInstanceRulesetRuleOptionsModel.SetPosition(positionModel)
				 updateInstanceRulesetRuleOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				 Expect(updateInstanceRulesetRuleOptionsModel).ToNot(BeNil())
				 Expect(updateInstanceRulesetRuleOptionsModel.RulesetID).To(Equal(core.StringPtr("testString")))
				 Expect(updateInstanceRulesetRuleOptionsModel.RuleID).To(Equal(core.StringPtr("testString")))
				 Expect(updateInstanceRulesetRuleOptionsModel.Action).To(Equal(core.StringPtr("testString")))
				 Expect(updateInstanceRulesetRuleOptionsModel.ActionParameters).To(Equal(actionParametersModel))
				 Expect(updateInstanceRulesetRuleOptionsModel.Description).To(Equal(core.StringPtr("testString")))
				 Expect(updateInstanceRulesetRuleOptionsModel.Enabled).To(Equal(core.BoolPtr(true)))
				 Expect(updateInstanceRulesetRuleOptionsModel.Expression).To(Equal(core.StringPtr("ip.src ne 1.1.1.1")))
				 Expect(updateInstanceRulesetRuleOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				 Expect(updateInstanceRulesetRuleOptionsModel.Logging).To(Equal(loggingModel))
				 Expect(updateInstanceRulesetRuleOptionsModel.Ref).To(Equal(core.StringPtr("my_ref")))
				 Expect(updateInstanceRulesetRuleOptionsModel.Position).To(Equal(positionModel))
				 Expect(updateInstanceRulesetRuleOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
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
				 actionParametersResponseModel.Content = core.StringPtr("{\"success\": false, \"error\": \"you have been blocked\"}")
				 actionParametersResponseModel.ContentType = core.StringPtr("application/json")
				 actionParametersResponseModel.StatusCode = core.Int64Ptr(int64(400))
				 Expect(actionParametersResponseModel.Content).To(Equal(core.StringPtr("{\"success\": false, \"error\": \"you have been blocked\"}")))
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
				 updateZoneEntrypointRulesetOptionsModel.SetDescription("Custom instance ruleset")
				 updateZoneEntrypointRulesetOptionsModel.SetKind("managed")
				 updateZoneEntrypointRulesetOptionsModel.SetName("testString")
				 updateZoneEntrypointRulesetOptionsModel.SetPhase("ddos_l4")
				 updateZoneEntrypointRulesetOptionsModel.SetRules([]rulesetsv1.RuleCreate{*ruleCreateModel})
				 updateZoneEntrypointRulesetOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				 Expect(updateZoneEntrypointRulesetOptionsModel).ToNot(BeNil())
				 Expect(updateZoneEntrypointRulesetOptionsModel.RulesetPhase).To(Equal(core.StringPtr("ddos_l4")))
				 Expect(updateZoneEntrypointRulesetOptionsModel.Description).To(Equal(core.StringPtr("Custom instance ruleset")))
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
				 actionParametersResponseModel.Content = core.StringPtr("{\"success\": false, \"error\": \"you have been blocked\"}")
				 actionParametersResponseModel.ContentType = core.StringPtr("application/json")
				 actionParametersResponseModel.StatusCode = core.Int64Ptr(int64(400))
				 Expect(actionParametersResponseModel.Content).To(Equal(core.StringPtr("{\"success\": false, \"error\": \"you have been blocked\"}")))
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
				 updateZoneRulesetOptionsModel.SetDescription("Custom instance ruleset")
				 updateZoneRulesetOptionsModel.SetKind("managed")
				 updateZoneRulesetOptionsModel.SetName("testString")
				 updateZoneRulesetOptionsModel.SetPhase("ddos_l4")
				 updateZoneRulesetOptionsModel.SetRules([]rulesetsv1.RuleCreate{*ruleCreateModel})
				 updateZoneRulesetOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				 Expect(updateZoneRulesetOptionsModel).ToNot(BeNil())
				 Expect(updateZoneRulesetOptionsModel.RulesetID).To(Equal(core.StringPtr("testString")))
				 Expect(updateZoneRulesetOptionsModel.Description).To(Equal(core.StringPtr("Custom instance ruleset")))
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
				 actionParametersResponseModel.Content = core.StringPtr("{\"success\": false, \"error\": \"you have been blocked\"}")
				 actionParametersResponseModel.ContentType = core.StringPtr("application/json")
				 actionParametersResponseModel.StatusCode = core.Int64Ptr(int64(400))
				 Expect(actionParametersResponseModel.Content).To(Equal(core.StringPtr("{\"success\": false, \"error\": \"you have been blocked\"}")))
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
 