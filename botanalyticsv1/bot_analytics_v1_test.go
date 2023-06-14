/**
 * (C) Copyright IBM Corp. 2023.
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

package botanalyticsv1_test

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
	"github.com/IBM/networking-go-sdk/botanalyticsv1"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe(`BotAnalyticsV1`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		It(`Instantiate service client`, func() {
			botAnalyticsService, serviceErr := botanalyticsv1.NewBotAnalyticsV1(&botanalyticsv1.BotAnalyticsV1Options{
				Authenticator:  &core.NoAuthAuthenticator{},
				Crn:            core.StringPtr(crn),
				ZoneIdentifier: core.StringPtr(zoneIdentifier),
			})
			Expect(botAnalyticsService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			botAnalyticsService, serviceErr := botanalyticsv1.NewBotAnalyticsV1(&botanalyticsv1.BotAnalyticsV1Options{
				URL:            "{BAD_URL_STRING",
				Crn:            core.StringPtr(crn),
				ZoneIdentifier: core.StringPtr(zoneIdentifier),
			})
			Expect(botAnalyticsService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			botAnalyticsService, serviceErr := botanalyticsv1.NewBotAnalyticsV1(&botanalyticsv1.BotAnalyticsV1Options{
				URL:            "https://botanalyticsv1/api",
				Crn:            core.StringPtr(crn),
				ZoneIdentifier: core.StringPtr(zoneIdentifier),
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(botAnalyticsService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Validation Error`, func() {
			botAnalyticsService, serviceErr := botanalyticsv1.NewBotAnalyticsV1(&botanalyticsv1.BotAnalyticsV1Options{})
			Expect(botAnalyticsService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"BOT_ANALYTICS_URL":       "https://botanalyticsv1/api",
				"BOT_ANALYTICS_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				botAnalyticsService, serviceErr := botanalyticsv1.NewBotAnalyticsV1UsingExternalConfig(&botanalyticsv1.BotAnalyticsV1Options{
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(botAnalyticsService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := botAnalyticsService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != botAnalyticsService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(botAnalyticsService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(botAnalyticsService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				botAnalyticsService, serviceErr := botanalyticsv1.NewBotAnalyticsV1UsingExternalConfig(&botanalyticsv1.BotAnalyticsV1Options{
					URL:            "https://testService/api",
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(botAnalyticsService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(botAnalyticsService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := botAnalyticsService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != botAnalyticsService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(botAnalyticsService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(botAnalyticsService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				botAnalyticsService, serviceErr := botanalyticsv1.NewBotAnalyticsV1UsingExternalConfig(&botanalyticsv1.BotAnalyticsV1Options{
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				err := botAnalyticsService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(botAnalyticsService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(botAnalyticsService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := botAnalyticsService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != botAnalyticsService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(botAnalyticsService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(botAnalyticsService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"BOT_ANALYTICS_URL":       "https://botanalyticsv1/api",
				"BOT_ANALYTICS_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			botAnalyticsService, serviceErr := botanalyticsv1.NewBotAnalyticsV1UsingExternalConfig(&botanalyticsv1.BotAnalyticsV1Options{
				Crn:            core.StringPtr(crn),
				ZoneIdentifier: core.StringPtr(zoneIdentifier),
			})

			It(`Instantiate service client with error`, func() {
				Expect(botAnalyticsService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"BOT_ANALYTICS_AUTH_TYPE": "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			botAnalyticsService, serviceErr := botanalyticsv1.NewBotAnalyticsV1UsingExternalConfig(&botanalyticsv1.BotAnalyticsV1Options{
				URL:            "{BAD_URL_STRING",
				Crn:            core.StringPtr(crn),
				ZoneIdentifier: core.StringPtr(zoneIdentifier),
			})

			It(`Instantiate service client with error`, func() {
				Expect(botAnalyticsService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = botanalyticsv1.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`GetBotScore(getBotScoreOptions *GetBotScoreOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getBotScorePath := "/v1/testString/zones/testString/bot_analytics/score_source"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getBotScorePath))
					Expect(req.Method).To(Equal("GET"))
					// TODO: Add check for since query parameter
					// TODO: Add check for until query parameter
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetBotScore with error: Operation response processing error`, func() {
				botAnalyticsService, serviceErr := botanalyticsv1.NewBotAnalyticsV1(&botanalyticsv1.BotAnalyticsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(botAnalyticsService).ToNot(BeNil())

				// Construct an instance of the GetBotScoreOptions model
				getBotScoreOptionsModel := new(botanalyticsv1.GetBotScoreOptions)
				getBotScoreOptionsModel.Since = CreateMockDateTime("2021-06-10T00:00:00Z")
				getBotScoreOptionsModel.Until = CreateMockDateTime("2021-06-11T00:00:00Z")
				getBotScoreOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := botAnalyticsService.GetBotScore(getBotScoreOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				botAnalyticsService.EnableRetries(0, 0)
				result, response, operationErr = botAnalyticsService.GetBotScore(getBotScoreOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetBotScore(getBotScoreOptions *GetBotScoreOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getBotScorePath := "/v1/testString/zones/testString/bot_analytics/score_source"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getBotScorePath))
					Expect(req.Method).To(Equal("GET"))

					// TODO: Add check for since query parameter
					// TODO: Add check for until query parameter
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": [{"botScore": [{"avg": {"sampleInterval": 14}, "count": 5, "dimensions": {"botScoreSrcName": "BotScoreSrcName"}}]}]}`)
				}))
			})
			It(`Invoke GetBotScore successfully with retries`, func() {
				botAnalyticsService, serviceErr := botanalyticsv1.NewBotAnalyticsV1(&botanalyticsv1.BotAnalyticsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(botAnalyticsService).ToNot(BeNil())
				botAnalyticsService.EnableRetries(0, 0)

				// Construct an instance of the GetBotScoreOptions model
				getBotScoreOptionsModel := new(botanalyticsv1.GetBotScoreOptions)
				getBotScoreOptionsModel.Since = CreateMockDateTime("2021-06-10T00:00:00Z")
				getBotScoreOptionsModel.Until = CreateMockDateTime("2021-06-11T00:00:00Z")
				getBotScoreOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := botAnalyticsService.GetBotScoreWithContext(ctx, getBotScoreOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				botAnalyticsService.DisableRetries()
				result, response, operationErr := botAnalyticsService.GetBotScore(getBotScoreOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = botAnalyticsService.GetBotScoreWithContext(ctx, getBotScoreOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getBotScorePath))
					Expect(req.Method).To(Equal("GET"))

					// TODO: Add check for since query parameter
					// TODO: Add check for until query parameter
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": [{"botScore": [{"avg": {"sampleInterval": 14}, "count": 5, "dimensions": {"botScoreSrcName": "BotScoreSrcName"}}]}]}`)
				}))
			})
			It(`Invoke GetBotScore successfully`, func() {
				botAnalyticsService, serviceErr := botanalyticsv1.NewBotAnalyticsV1(&botanalyticsv1.BotAnalyticsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(botAnalyticsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := botAnalyticsService.GetBotScore(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetBotScoreOptions model
				getBotScoreOptionsModel := new(botanalyticsv1.GetBotScoreOptions)
				getBotScoreOptionsModel.Since = CreateMockDateTime("2021-06-10T00:00:00Z")
				getBotScoreOptionsModel.Until = CreateMockDateTime("2021-06-11T00:00:00Z")
				getBotScoreOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = botAnalyticsService.GetBotScore(getBotScoreOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetBotScore with error: Operation validation and request error`, func() {
				botAnalyticsService, serviceErr := botanalyticsv1.NewBotAnalyticsV1(&botanalyticsv1.BotAnalyticsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(botAnalyticsService).ToNot(BeNil())

				// Construct an instance of the GetBotScoreOptions model
				getBotScoreOptionsModel := new(botanalyticsv1.GetBotScoreOptions)
				getBotScoreOptionsModel.Since = CreateMockDateTime("2021-06-10T00:00:00Z")
				getBotScoreOptionsModel.Until = CreateMockDateTime("2021-06-11T00:00:00Z")
				getBotScoreOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := botAnalyticsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := botAnalyticsService.GetBotScore(getBotScoreOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetBotScoreOptions model with no property values
				getBotScoreOptionsModelNew := new(botanalyticsv1.GetBotScoreOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = botAnalyticsService.GetBotScore(getBotScoreOptionsModelNew)
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
			It(`Invoke GetBotScore successfully`, func() {
				botAnalyticsService, serviceErr := botanalyticsv1.NewBotAnalyticsV1(&botanalyticsv1.BotAnalyticsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(botAnalyticsService).ToNot(BeNil())

				// Construct an instance of the GetBotScoreOptions model
				getBotScoreOptionsModel := new(botanalyticsv1.GetBotScoreOptions)
				getBotScoreOptionsModel.Since = CreateMockDateTime("2021-06-10T00:00:00Z")
				getBotScoreOptionsModel.Until = CreateMockDateTime("2021-06-11T00:00:00Z")
				getBotScoreOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := botAnalyticsService.GetBotScore(getBotScoreOptionsModel)
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
	Describe(`GetBotTimeseries(getBotTimeseriesOptions *GetBotTimeseriesOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getBotTimeseriesPath := "/v1/testString/zones/testString/bot_analytics/timeseries"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getBotTimeseriesPath))
					Expect(req.Method).To(Equal("GET"))
					// TODO: Add check for since query parameter
					// TODO: Add check for until query parameter
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetBotTimeseries with error: Operation response processing error`, func() {
				botAnalyticsService, serviceErr := botanalyticsv1.NewBotAnalyticsV1(&botanalyticsv1.BotAnalyticsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(botAnalyticsService).ToNot(BeNil())

				// Construct an instance of the GetBotTimeseriesOptions model
				getBotTimeseriesOptionsModel := new(botanalyticsv1.GetBotTimeseriesOptions)
				getBotTimeseriesOptionsModel.Since = CreateMockDateTime("2021-06-10T00:00:00Z")
				getBotTimeseriesOptionsModel.Until = CreateMockDateTime("2021-06-11T00:00:00Z")
				getBotTimeseriesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := botAnalyticsService.GetBotTimeseries(getBotTimeseriesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				botAnalyticsService.EnableRetries(0, 0)
				result, response, operationErr = botAnalyticsService.GetBotTimeseries(getBotTimeseriesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetBotTimeseries(getBotTimeseriesOptions *GetBotTimeseriesOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getBotTimeseriesPath := "/v1/testString/zones/testString/bot_analytics/timeseries"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getBotTimeseriesPath))
					Expect(req.Method).To(Equal("GET"))

					// TODO: Add check for since query parameter
					// TODO: Add check for until query parameter
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": [{"botScore": [{"anyKey": "anyValue"}]}]}`)
				}))
			})
			It(`Invoke GetBotTimeseries successfully with retries`, func() {
				botAnalyticsService, serviceErr := botanalyticsv1.NewBotAnalyticsV1(&botanalyticsv1.BotAnalyticsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(botAnalyticsService).ToNot(BeNil())
				botAnalyticsService.EnableRetries(0, 0)

				// Construct an instance of the GetBotTimeseriesOptions model
				getBotTimeseriesOptionsModel := new(botanalyticsv1.GetBotTimeseriesOptions)
				getBotTimeseriesOptionsModel.Since = CreateMockDateTime("2021-06-10T00:00:00Z")
				getBotTimeseriesOptionsModel.Until = CreateMockDateTime("2021-06-11T00:00:00Z")
				getBotTimeseriesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := botAnalyticsService.GetBotTimeseriesWithContext(ctx, getBotTimeseriesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				botAnalyticsService.DisableRetries()
				result, response, operationErr := botAnalyticsService.GetBotTimeseries(getBotTimeseriesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = botAnalyticsService.GetBotTimeseriesWithContext(ctx, getBotTimeseriesOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getBotTimeseriesPath))
					Expect(req.Method).To(Equal("GET"))

					// TODO: Add check for since query parameter
					// TODO: Add check for until query parameter
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": [{"botScore": [{"anyKey": "anyValue"}]}]}`)
				}))
			})
			It(`Invoke GetBotTimeseries successfully`, func() {
				botAnalyticsService, serviceErr := botanalyticsv1.NewBotAnalyticsV1(&botanalyticsv1.BotAnalyticsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(botAnalyticsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := botAnalyticsService.GetBotTimeseries(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetBotTimeseriesOptions model
				getBotTimeseriesOptionsModel := new(botanalyticsv1.GetBotTimeseriesOptions)
				getBotTimeseriesOptionsModel.Since = CreateMockDateTime("2021-06-10T00:00:00Z")
				getBotTimeseriesOptionsModel.Until = CreateMockDateTime("2021-06-11T00:00:00Z")
				getBotTimeseriesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = botAnalyticsService.GetBotTimeseries(getBotTimeseriesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetBotTimeseries with error: Operation validation and request error`, func() {
				botAnalyticsService, serviceErr := botanalyticsv1.NewBotAnalyticsV1(&botanalyticsv1.BotAnalyticsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(botAnalyticsService).ToNot(BeNil())

				// Construct an instance of the GetBotTimeseriesOptions model
				getBotTimeseriesOptionsModel := new(botanalyticsv1.GetBotTimeseriesOptions)
				getBotTimeseriesOptionsModel.Since = CreateMockDateTime("2021-06-10T00:00:00Z")
				getBotTimeseriesOptionsModel.Until = CreateMockDateTime("2021-06-11T00:00:00Z")
				getBotTimeseriesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := botAnalyticsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := botAnalyticsService.GetBotTimeseries(getBotTimeseriesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetBotTimeseriesOptions model with no property values
				getBotTimeseriesOptionsModelNew := new(botanalyticsv1.GetBotTimeseriesOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = botAnalyticsService.GetBotTimeseries(getBotTimeseriesOptionsModelNew)
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
			It(`Invoke GetBotTimeseries successfully`, func() {
				botAnalyticsService, serviceErr := botanalyticsv1.NewBotAnalyticsV1(&botanalyticsv1.BotAnalyticsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(botAnalyticsService).ToNot(BeNil())

				// Construct an instance of the GetBotTimeseriesOptions model
				getBotTimeseriesOptionsModel := new(botanalyticsv1.GetBotTimeseriesOptions)
				getBotTimeseriesOptionsModel.Since = CreateMockDateTime("2021-06-10T00:00:00Z")
				getBotTimeseriesOptionsModel.Until = CreateMockDateTime("2021-06-11T00:00:00Z")
				getBotTimeseriesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := botAnalyticsService.GetBotTimeseries(getBotTimeseriesOptionsModel)
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
	Describe(`GetBotTopns(getBotTopnsOptions *GetBotTopnsOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getBotTopnsPath := "/v1/testString/zones/testString/bot_analytics/top_ns"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getBotTopnsPath))
					Expect(req.Method).To(Equal("GET"))
					// TODO: Add check for since query parameter
					// TODO: Add check for until query parameter
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetBotTopns with error: Operation response processing error`, func() {
				botAnalyticsService, serviceErr := botanalyticsv1.NewBotAnalyticsV1(&botanalyticsv1.BotAnalyticsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(botAnalyticsService).ToNot(BeNil())

				// Construct an instance of the GetBotTopnsOptions model
				getBotTopnsOptionsModel := new(botanalyticsv1.GetBotTopnsOptions)
				getBotTopnsOptionsModel.Since = CreateMockDateTime("2021-06-10T00:00:00Z")
				getBotTopnsOptionsModel.Until = CreateMockDateTime("2021-06-11T00:00:00Z")
				getBotTopnsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := botAnalyticsService.GetBotTopns(getBotTopnsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				botAnalyticsService.EnableRetries(0, 0)
				result, response, operationErr = botAnalyticsService.GetBotTopns(getBotTopnsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetBotTopns(getBotTopnsOptions *GetBotTopnsOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getBotTopnsPath := "/v1/testString/zones/testString/bot_analytics/top_ns"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getBotTopnsPath))
					Expect(req.Method).To(Equal("GET"))

					// TODO: Add check for since query parameter
					// TODO: Add check for until query parameter
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": [{"anyKey": "anyValue"}]}`)
				}))
			})
			It(`Invoke GetBotTopns successfully with retries`, func() {
				botAnalyticsService, serviceErr := botanalyticsv1.NewBotAnalyticsV1(&botanalyticsv1.BotAnalyticsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(botAnalyticsService).ToNot(BeNil())
				botAnalyticsService.EnableRetries(0, 0)

				// Construct an instance of the GetBotTopnsOptions model
				getBotTopnsOptionsModel := new(botanalyticsv1.GetBotTopnsOptions)
				getBotTopnsOptionsModel.Since = CreateMockDateTime("2021-06-10T00:00:00Z")
				getBotTopnsOptionsModel.Until = CreateMockDateTime("2021-06-11T00:00:00Z")
				getBotTopnsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := botAnalyticsService.GetBotTopnsWithContext(ctx, getBotTopnsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				botAnalyticsService.DisableRetries()
				result, response, operationErr := botAnalyticsService.GetBotTopns(getBotTopnsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = botAnalyticsService.GetBotTopnsWithContext(ctx, getBotTopnsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getBotTopnsPath))
					Expect(req.Method).To(Equal("GET"))

					// TODO: Add check for since query parameter
					// TODO: Add check for until query parameter
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": [{"anyKey": "anyValue"}]}`)
				}))
			})
			It(`Invoke GetBotTopns successfully`, func() {
				botAnalyticsService, serviceErr := botanalyticsv1.NewBotAnalyticsV1(&botanalyticsv1.BotAnalyticsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(botAnalyticsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := botAnalyticsService.GetBotTopns(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetBotTopnsOptions model
				getBotTopnsOptionsModel := new(botanalyticsv1.GetBotTopnsOptions)
				getBotTopnsOptionsModel.Since = CreateMockDateTime("2021-06-10T00:00:00Z")
				getBotTopnsOptionsModel.Until = CreateMockDateTime("2021-06-11T00:00:00Z")
				getBotTopnsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = botAnalyticsService.GetBotTopns(getBotTopnsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetBotTopns with error: Operation validation and request error`, func() {
				botAnalyticsService, serviceErr := botanalyticsv1.NewBotAnalyticsV1(&botanalyticsv1.BotAnalyticsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(botAnalyticsService).ToNot(BeNil())

				// Construct an instance of the GetBotTopnsOptions model
				getBotTopnsOptionsModel := new(botanalyticsv1.GetBotTopnsOptions)
				getBotTopnsOptionsModel.Since = CreateMockDateTime("2021-06-10T00:00:00Z")
				getBotTopnsOptionsModel.Until = CreateMockDateTime("2021-06-11T00:00:00Z")
				getBotTopnsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := botAnalyticsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := botAnalyticsService.GetBotTopns(getBotTopnsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetBotTopnsOptions model with no property values
				getBotTopnsOptionsModelNew := new(botanalyticsv1.GetBotTopnsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = botAnalyticsService.GetBotTopns(getBotTopnsOptionsModelNew)
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
			It(`Invoke GetBotTopns successfully`, func() {
				botAnalyticsService, serviceErr := botanalyticsv1.NewBotAnalyticsV1(&botanalyticsv1.BotAnalyticsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(botAnalyticsService).ToNot(BeNil())

				// Construct an instance of the GetBotTopnsOptions model
				getBotTopnsOptionsModel := new(botanalyticsv1.GetBotTopnsOptions)
				getBotTopnsOptionsModel.Since = CreateMockDateTime("2021-06-10T00:00:00Z")
				getBotTopnsOptionsModel.Until = CreateMockDateTime("2021-06-11T00:00:00Z")
				getBotTopnsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := botAnalyticsService.GetBotTopns(getBotTopnsOptionsModel)
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
			botAnalyticsService, _ := botanalyticsv1.NewBotAnalyticsV1(&botanalyticsv1.BotAnalyticsV1Options{
				URL:            "http://botanalyticsv1modelgenerator.com",
				Authenticator:  &core.NoAuthAuthenticator{},
				Crn:            core.StringPtr(crn),
				ZoneIdentifier: core.StringPtr(zoneIdentifier),
			})
			It(`Invoke NewGetBotScoreOptions successfully`, func() {
				// Construct an instance of the GetBotScoreOptions model
				since := CreateMockDateTime("2021-06-10T00:00:00Z")
				until := CreateMockDateTime("2021-06-11T00:00:00Z")
				getBotScoreOptionsModel := botAnalyticsService.NewGetBotScoreOptions(since, until)
				getBotScoreOptionsModel.SetSince(CreateMockDateTime("2021-06-10T00:00:00Z"))
				getBotScoreOptionsModel.SetUntil(CreateMockDateTime("2021-06-11T00:00:00Z"))
				getBotScoreOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getBotScoreOptionsModel).ToNot(BeNil())
				Expect(getBotScoreOptionsModel.Since).To(Equal(CreateMockDateTime("2021-06-10T00:00:00Z")))
				Expect(getBotScoreOptionsModel.Until).To(Equal(CreateMockDateTime("2021-06-11T00:00:00Z")))
				Expect(getBotScoreOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetBotTimeseriesOptions successfully`, func() {
				// Construct an instance of the GetBotTimeseriesOptions model
				since := CreateMockDateTime("2021-06-10T00:00:00Z")
				until := CreateMockDateTime("2021-06-11T00:00:00Z")
				getBotTimeseriesOptionsModel := botAnalyticsService.NewGetBotTimeseriesOptions(since, until)
				getBotTimeseriesOptionsModel.SetSince(CreateMockDateTime("2021-06-10T00:00:00Z"))
				getBotTimeseriesOptionsModel.SetUntil(CreateMockDateTime("2021-06-11T00:00:00Z"))
				getBotTimeseriesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getBotTimeseriesOptionsModel).ToNot(BeNil())
				Expect(getBotTimeseriesOptionsModel.Since).To(Equal(CreateMockDateTime("2021-06-10T00:00:00Z")))
				Expect(getBotTimeseriesOptionsModel.Until).To(Equal(CreateMockDateTime("2021-06-11T00:00:00Z")))
				Expect(getBotTimeseriesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetBotTopnsOptions successfully`, func() {
				// Construct an instance of the GetBotTopnsOptions model
				since := CreateMockDateTime("2021-06-10T00:00:00Z")
				until := CreateMockDateTime("2021-06-11T00:00:00Z")
				getBotTopnsOptionsModel := botAnalyticsService.NewGetBotTopnsOptions(since, until)
				getBotTopnsOptionsModel.SetSince(CreateMockDateTime("2021-06-10T00:00:00Z"))
				getBotTopnsOptionsModel.SetUntil(CreateMockDateTime("2021-06-11T00:00:00Z"))
				getBotTopnsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getBotTopnsOptionsModel).ToNot(BeNil())
				Expect(getBotTopnsOptionsModel.Since).To(Equal(CreateMockDateTime("2021-06-10T00:00:00Z")))
				Expect(getBotTopnsOptionsModel.Until).To(Equal(CreateMockDateTime("2021-06-11T00:00:00Z")))
				Expect(getBotTopnsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
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
