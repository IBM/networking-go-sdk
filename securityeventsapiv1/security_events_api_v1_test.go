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

package securityeventsapiv1_test

import (
	"bytes"
	"context"
	"fmt"
	"github.com/IBM/go-sdk-core/v4/core"
	"github.com/IBM/networking-go-sdk/securityeventsapiv1"
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

var _ = Describe(`SecurityEventsApiV1`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		crn := "testString"
		zoneID := "testString"
		It(`Instantiate service client`, func() {
			securityEventsApiService, serviceErr := securityeventsapiv1.NewSecurityEventsApiV1(&securityeventsapiv1.SecurityEventsApiV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
				Crn: core.StringPtr(crn),
				ZoneID: core.StringPtr(zoneID),
			})
			Expect(securityEventsApiService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			securityEventsApiService, serviceErr := securityeventsapiv1.NewSecurityEventsApiV1(&securityeventsapiv1.SecurityEventsApiV1Options{
				URL: "{BAD_URL_STRING",
				Crn: core.StringPtr(crn),
				ZoneID: core.StringPtr(zoneID),
			})
			Expect(securityEventsApiService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			securityEventsApiService, serviceErr := securityeventsapiv1.NewSecurityEventsApiV1(&securityeventsapiv1.SecurityEventsApiV1Options{
				URL: "https://securityeventsapiv1/api",
				Crn: core.StringPtr(crn),
				ZoneID: core.StringPtr(zoneID),
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(securityEventsApiService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Validation Error`, func() {
			securityEventsApiService, serviceErr := securityeventsapiv1.NewSecurityEventsApiV1(&securityeventsapiv1.SecurityEventsApiV1Options{})
			Expect(securityEventsApiService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		crn := "testString"
		zoneID := "testString"
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"SECURITY_EVENTS_API_URL": "https://securityeventsapiv1/api",
				"SECURITY_EVENTS_API_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				securityEventsApiService, serviceErr := securityeventsapiv1.NewSecurityEventsApiV1UsingExternalConfig(&securityeventsapiv1.SecurityEventsApiV1Options{
					Crn: core.StringPtr(crn),
					ZoneID: core.StringPtr(zoneID),
				})
				Expect(securityEventsApiService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := securityEventsApiService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != securityEventsApiService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(securityEventsApiService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(securityEventsApiService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				securityEventsApiService, serviceErr := securityeventsapiv1.NewSecurityEventsApiV1UsingExternalConfig(&securityeventsapiv1.SecurityEventsApiV1Options{
					URL: "https://testService/api",
					Crn: core.StringPtr(crn),
					ZoneID: core.StringPtr(zoneID),
				})
				Expect(securityEventsApiService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(securityEventsApiService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := securityEventsApiService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != securityEventsApiService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(securityEventsApiService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(securityEventsApiService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				securityEventsApiService, serviceErr := securityeventsapiv1.NewSecurityEventsApiV1UsingExternalConfig(&securityeventsapiv1.SecurityEventsApiV1Options{
					Crn: core.StringPtr(crn),
					ZoneID: core.StringPtr(zoneID),
				})
				err := securityEventsApiService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(securityEventsApiService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(securityEventsApiService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := securityEventsApiService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != securityEventsApiService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(securityEventsApiService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(securityEventsApiService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"SECURITY_EVENTS_API_URL": "https://securityeventsapiv1/api",
				"SECURITY_EVENTS_API_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			securityEventsApiService, serviceErr := securityeventsapiv1.NewSecurityEventsApiV1UsingExternalConfig(&securityeventsapiv1.SecurityEventsApiV1Options{
				Crn: core.StringPtr(crn),
				ZoneID: core.StringPtr(zoneID),
			})

			It(`Instantiate service client with error`, func() {
				Expect(securityEventsApiService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"SECURITY_EVENTS_API_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			securityEventsApiService, serviceErr := securityeventsapiv1.NewSecurityEventsApiV1UsingExternalConfig(&securityeventsapiv1.SecurityEventsApiV1Options{
				URL: "{BAD_URL_STRING",
				Crn: core.StringPtr(crn),
				ZoneID: core.StringPtr(zoneID),
			})

			It(`Instantiate service client with error`, func() {
				Expect(securityEventsApiService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = securityeventsapiv1.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`SecurityEvents(securityEventsOptions *SecurityEventsOptions) - Operation response error`, func() {
		crn := "testString"
		zoneID := "testString"
		securityEventsPath := "/v1/testString/zones/testString/security/events"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(securityEventsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["ip_class"]).To(Equal([]string{"unknown"}))

					Expect(req.URL.Query()["method"]).To(Equal([]string{"GET"}))

					Expect(req.URL.Query()["scheme"]).To(Equal([]string{"unknown"}))

					Expect(req.URL.Query()["ip"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["host"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["proto"]).To(Equal([]string{"UNK"}))

					Expect(req.URL.Query()["uri"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["ua"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["colo"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["ray_id"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["kind"]).To(Equal([]string{"firewall"}))

					Expect(req.URL.Query()["action"]).To(Equal([]string{"unknown"}))

					Expect(req.URL.Query()["cursor"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["country"]).To(Equal([]string{"testString"}))


					// TODO: Add check for since query parameter

					Expect(req.URL.Query()["source"]).To(Equal([]string{"unknown"}))

					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))

					Expect(req.URL.Query()["rule_id"]).To(Equal([]string{"testString"}))


					// TODO: Add check for until query parameter

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke SecurityEvents with error: Operation response processing error`, func() {
				securityEventsApiService, serviceErr := securityeventsapiv1.NewSecurityEventsApiV1(&securityeventsapiv1.SecurityEventsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneID: core.StringPtr(zoneID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityEventsApiService).ToNot(BeNil())

				// Construct an instance of the SecurityEventsOptions model
				securityEventsOptionsModel := new(securityeventsapiv1.SecurityEventsOptions)
				securityEventsOptionsModel.IpClass = core.StringPtr("unknown")
				securityEventsOptionsModel.Method = core.StringPtr("GET")
				securityEventsOptionsModel.Scheme = core.StringPtr("unknown")
				securityEventsOptionsModel.Ip = core.StringPtr("testString")
				securityEventsOptionsModel.Host = core.StringPtr("testString")
				securityEventsOptionsModel.Proto = core.StringPtr("UNK")
				securityEventsOptionsModel.URI = core.StringPtr("testString")
				securityEventsOptionsModel.Ua = core.StringPtr("testString")
				securityEventsOptionsModel.Colo = core.StringPtr("testString")
				securityEventsOptionsModel.RayID = core.StringPtr("testString")
				securityEventsOptionsModel.Kind = core.StringPtr("firewall")
				securityEventsOptionsModel.Action = core.StringPtr("unknown")
				securityEventsOptionsModel.Cursor = core.StringPtr("testString")
				securityEventsOptionsModel.Country = core.StringPtr("testString")
				securityEventsOptionsModel.Since = CreateMockDateTime()
				securityEventsOptionsModel.Source = core.StringPtr("unknown")
				securityEventsOptionsModel.Limit = core.Int64Ptr(int64(10))
				securityEventsOptionsModel.RuleID = core.StringPtr("testString")
				securityEventsOptionsModel.Until = CreateMockDateTime()
				securityEventsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := securityEventsApiService.SecurityEvents(securityEventsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				securityEventsApiService.EnableRetries(0, 0)
				result, response, operationErr = securityEventsApiService.SecurityEvents(securityEventsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`SecurityEvents(securityEventsOptions *SecurityEventsOptions)`, func() {
		crn := "testString"
		zoneID := "testString"
		securityEventsPath := "/v1/testString/zones/testString/security/events"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(securityEventsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["ip_class"]).To(Equal([]string{"unknown"}))

					Expect(req.URL.Query()["method"]).To(Equal([]string{"GET"}))

					Expect(req.URL.Query()["scheme"]).To(Equal([]string{"unknown"}))

					Expect(req.URL.Query()["ip"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["host"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["proto"]).To(Equal([]string{"UNK"}))

					Expect(req.URL.Query()["uri"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["ua"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["colo"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["ray_id"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["kind"]).To(Equal([]string{"firewall"}))

					Expect(req.URL.Query()["action"]).To(Equal([]string{"unknown"}))

					Expect(req.URL.Query()["cursor"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["country"]).To(Equal([]string{"testString"}))


					// TODO: Add check for since query parameter

					Expect(req.URL.Query()["source"]).To(Equal([]string{"unknown"}))

					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))

					Expect(req.URL.Query()["rule_id"]).To(Equal([]string{"testString"}))


					// TODO: Add check for until query parameter

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"result": [{"ray_id": "4c6392789858b224", "kind": "firewall", "source": "rateLimit", "action": "drop", "rule_id": "fe38bd35ca284de69b5ecbaa6db87dc3", "ip": "192.168.1.1", "ip_class": "noRecord", "country": "CN", "colo": "HKG", "host": "www.example.com", "method": "GET", "proto": "HTTP/2", "scheme": "https", "ua": "curl/7.61.1", "uri": "/", "occurred_at": "2019-01-01T12:00:00", "matches": [{"rule_id": "fe38bd35ca284de69b5ecbaa6db87dc3", "source": "rateLimit", "action": "drop", "metadata": {"anyKey": "anyValue"}}]}], "result_info": {"cursors": {"after": "bnRIiaU-14b2YBxIefX28h7Zqw50XXPA4Vu4Sa-DPa4qaGH-z47uwtOR0Hm2Y3cSh56raQb1POqaBwGXD44", "before": "dmmGxcD665xj3RiQ8eRqclts94GF3M4KpHEJ7AVekLtOUsHLHssfGaV_d8nZgLszk_iElB9LckPhFgmkTXHX"}, "scanned_range": {"since": "2019-04-12 07:44:18", "until": "2019-04-12 07:44:18"}}, "success": true, "errors": [["Errors"]], "messages": [["Messages"]]}`)
				}))
			})
			It(`Invoke SecurityEvents successfully`, func() {
				securityEventsApiService, serviceErr := securityeventsapiv1.NewSecurityEventsApiV1(&securityeventsapiv1.SecurityEventsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneID: core.StringPtr(zoneID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityEventsApiService).ToNot(BeNil())
				securityEventsApiService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := securityEventsApiService.SecurityEvents(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the SecurityEventsOptions model
				securityEventsOptionsModel := new(securityeventsapiv1.SecurityEventsOptions)
				securityEventsOptionsModel.IpClass = core.StringPtr("unknown")
				securityEventsOptionsModel.Method = core.StringPtr("GET")
				securityEventsOptionsModel.Scheme = core.StringPtr("unknown")
				securityEventsOptionsModel.Ip = core.StringPtr("testString")
				securityEventsOptionsModel.Host = core.StringPtr("testString")
				securityEventsOptionsModel.Proto = core.StringPtr("UNK")
				securityEventsOptionsModel.URI = core.StringPtr("testString")
				securityEventsOptionsModel.Ua = core.StringPtr("testString")
				securityEventsOptionsModel.Colo = core.StringPtr("testString")
				securityEventsOptionsModel.RayID = core.StringPtr("testString")
				securityEventsOptionsModel.Kind = core.StringPtr("firewall")
				securityEventsOptionsModel.Action = core.StringPtr("unknown")
				securityEventsOptionsModel.Cursor = core.StringPtr("testString")
				securityEventsOptionsModel.Country = core.StringPtr("testString")
				securityEventsOptionsModel.Since = CreateMockDateTime()
				securityEventsOptionsModel.Source = core.StringPtr("unknown")
				securityEventsOptionsModel.Limit = core.Int64Ptr(int64(10))
				securityEventsOptionsModel.RuleID = core.StringPtr("testString")
				securityEventsOptionsModel.Until = CreateMockDateTime()
				securityEventsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = securityEventsApiService.SecurityEvents(securityEventsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = securityEventsApiService.SecurityEventsWithContext(ctx, securityEventsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				securityEventsApiService.DisableRetries()
				result, response, operationErr = securityEventsApiService.SecurityEvents(securityEventsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = securityEventsApiService.SecurityEventsWithContext(ctx, securityEventsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke SecurityEvents with error: Operation request error`, func() {
				securityEventsApiService, serviceErr := securityeventsapiv1.NewSecurityEventsApiV1(&securityeventsapiv1.SecurityEventsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneID: core.StringPtr(zoneID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityEventsApiService).ToNot(BeNil())

				// Construct an instance of the SecurityEventsOptions model
				securityEventsOptionsModel := new(securityeventsapiv1.SecurityEventsOptions)
				securityEventsOptionsModel.IpClass = core.StringPtr("unknown")
				securityEventsOptionsModel.Method = core.StringPtr("GET")
				securityEventsOptionsModel.Scheme = core.StringPtr("unknown")
				securityEventsOptionsModel.Ip = core.StringPtr("testString")
				securityEventsOptionsModel.Host = core.StringPtr("testString")
				securityEventsOptionsModel.Proto = core.StringPtr("UNK")
				securityEventsOptionsModel.URI = core.StringPtr("testString")
				securityEventsOptionsModel.Ua = core.StringPtr("testString")
				securityEventsOptionsModel.Colo = core.StringPtr("testString")
				securityEventsOptionsModel.RayID = core.StringPtr("testString")
				securityEventsOptionsModel.Kind = core.StringPtr("firewall")
				securityEventsOptionsModel.Action = core.StringPtr("unknown")
				securityEventsOptionsModel.Cursor = core.StringPtr("testString")
				securityEventsOptionsModel.Country = core.StringPtr("testString")
				securityEventsOptionsModel.Since = CreateMockDateTime()
				securityEventsOptionsModel.Source = core.StringPtr("unknown")
				securityEventsOptionsModel.Limit = core.Int64Ptr(int64(10))
				securityEventsOptionsModel.RuleID = core.StringPtr("testString")
				securityEventsOptionsModel.Until = CreateMockDateTime()
				securityEventsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := securityEventsApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := securityEventsApiService.SecurityEvents(securityEventsOptionsModel)
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
			zoneID := "testString"
			securityEventsApiService, _ := securityeventsapiv1.NewSecurityEventsApiV1(&securityeventsapiv1.SecurityEventsApiV1Options{
				URL:           "http://securityeventsapiv1modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
				Crn: core.StringPtr(crn),
				ZoneID: core.StringPtr(zoneID),
			})
			It(`Invoke NewSecurityEventsOptions successfully`, func() {
				// Construct an instance of the SecurityEventsOptions model
				securityEventsOptionsModel := securityEventsApiService.NewSecurityEventsOptions()
				securityEventsOptionsModel.SetIpClass("unknown")
				securityEventsOptionsModel.SetMethod("GET")
				securityEventsOptionsModel.SetScheme("unknown")
				securityEventsOptionsModel.SetIp("testString")
				securityEventsOptionsModel.SetHost("testString")
				securityEventsOptionsModel.SetProto("UNK")
				securityEventsOptionsModel.SetURI("testString")
				securityEventsOptionsModel.SetUa("testString")
				securityEventsOptionsModel.SetColo("testString")
				securityEventsOptionsModel.SetRayID("testString")
				securityEventsOptionsModel.SetKind("firewall")
				securityEventsOptionsModel.SetAction("unknown")
				securityEventsOptionsModel.SetCursor("testString")
				securityEventsOptionsModel.SetCountry("testString")
				securityEventsOptionsModel.SetSince(CreateMockDateTime())
				securityEventsOptionsModel.SetSource("unknown")
				securityEventsOptionsModel.SetLimit(int64(10))
				securityEventsOptionsModel.SetRuleID("testString")
				securityEventsOptionsModel.SetUntil(CreateMockDateTime())
				securityEventsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(securityEventsOptionsModel).ToNot(BeNil())
				Expect(securityEventsOptionsModel.IpClass).To(Equal(core.StringPtr("unknown")))
				Expect(securityEventsOptionsModel.Method).To(Equal(core.StringPtr("GET")))
				Expect(securityEventsOptionsModel.Scheme).To(Equal(core.StringPtr("unknown")))
				Expect(securityEventsOptionsModel.Ip).To(Equal(core.StringPtr("testString")))
				Expect(securityEventsOptionsModel.Host).To(Equal(core.StringPtr("testString")))
				Expect(securityEventsOptionsModel.Proto).To(Equal(core.StringPtr("UNK")))
				Expect(securityEventsOptionsModel.URI).To(Equal(core.StringPtr("testString")))
				Expect(securityEventsOptionsModel.Ua).To(Equal(core.StringPtr("testString")))
				Expect(securityEventsOptionsModel.Colo).To(Equal(core.StringPtr("testString")))
				Expect(securityEventsOptionsModel.RayID).To(Equal(core.StringPtr("testString")))
				Expect(securityEventsOptionsModel.Kind).To(Equal(core.StringPtr("firewall")))
				Expect(securityEventsOptionsModel.Action).To(Equal(core.StringPtr("unknown")))
				Expect(securityEventsOptionsModel.Cursor).To(Equal(core.StringPtr("testString")))
				Expect(securityEventsOptionsModel.Country).To(Equal(core.StringPtr("testString")))
				Expect(securityEventsOptionsModel.Since).To(Equal(CreateMockDateTime()))
				Expect(securityEventsOptionsModel.Source).To(Equal(core.StringPtr("unknown")))
				Expect(securityEventsOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(10))))
				Expect(securityEventsOptionsModel.RuleID).To(Equal(core.StringPtr("testString")))
				Expect(securityEventsOptionsModel.Until).To(Equal(CreateMockDateTime()))
				Expect(securityEventsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
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
