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

package routingv1_test

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
	"github.com/IBM/networking-go-sdk/routingv1"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe(`RoutingV1`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		It(`Instantiate service client`, func() {
			routingService, serviceErr := routingv1.NewRoutingV1(&routingv1.RoutingV1Options{
				Authenticator:  &core.NoAuthAuthenticator{},
				Crn:            core.StringPtr(crn),
				ZoneIdentifier: core.StringPtr(zoneIdentifier),
			})
			Expect(routingService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			routingService, serviceErr := routingv1.NewRoutingV1(&routingv1.RoutingV1Options{
				URL:            "{BAD_URL_STRING",
				Crn:            core.StringPtr(crn),
				ZoneIdentifier: core.StringPtr(zoneIdentifier),
			})
			Expect(routingService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			routingService, serviceErr := routingv1.NewRoutingV1(&routingv1.RoutingV1Options{
				URL:            "https://routingv1/api",
				Crn:            core.StringPtr(crn),
				ZoneIdentifier: core.StringPtr(zoneIdentifier),
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(routingService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Validation Error`, func() {
			routingService, serviceErr := routingv1.NewRoutingV1(&routingv1.RoutingV1Options{})
			Expect(routingService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"ROUTING_URL":       "https://routingv1/api",
				"ROUTING_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				routingService, serviceErr := routingv1.NewRoutingV1UsingExternalConfig(&routingv1.RoutingV1Options{
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(routingService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := routingService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != routingService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(routingService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(routingService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				routingService, serviceErr := routingv1.NewRoutingV1UsingExternalConfig(&routingv1.RoutingV1Options{
					URL:            "https://testService/api",
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(routingService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(routingService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := routingService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != routingService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(routingService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(routingService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				routingService, serviceErr := routingv1.NewRoutingV1UsingExternalConfig(&routingv1.RoutingV1Options{
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				err := routingService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(routingService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(routingService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := routingService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != routingService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(routingService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(routingService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"ROUTING_URL":       "https://routingv1/api",
				"ROUTING_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			routingService, serviceErr := routingv1.NewRoutingV1UsingExternalConfig(&routingv1.RoutingV1Options{
				Crn:            core.StringPtr(crn),
				ZoneIdentifier: core.StringPtr(zoneIdentifier),
			})

			It(`Instantiate service client with error`, func() {
				Expect(routingService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"ROUTING_AUTH_TYPE": "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			routingService, serviceErr := routingv1.NewRoutingV1UsingExternalConfig(&routingv1.RoutingV1Options{
				URL:            "{BAD_URL_STRING",
				Crn:            core.StringPtr(crn),
				ZoneIdentifier: core.StringPtr(zoneIdentifier),
			})

			It(`Instantiate service client with error`, func() {
				Expect(routingService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = routingv1.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`GetSmartRouting(getSmartRoutingOptions *GetSmartRoutingOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getSmartRoutingPath := "/v1/testString/zones/testString/routing/smart_routing"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getSmartRoutingPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetSmartRouting with error: Operation response processing error`, func() {
				routingService, serviceErr := routingv1.NewRoutingV1(&routingv1.RoutingV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(routingService).ToNot(BeNil())

				// Construct an instance of the GetSmartRoutingOptions model
				getSmartRoutingOptionsModel := new(routingv1.GetSmartRoutingOptions)
				getSmartRoutingOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := routingService.GetSmartRouting(getSmartRoutingOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				routingService.EnableRetries(0, 0)
				result, response, operationErr = routingService.GetSmartRouting(getSmartRoutingOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetSmartRouting(getSmartRoutingOptions *GetSmartRoutingOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getSmartRoutingPath := "/v1/testString/zones/testString/routing/smart_routing"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getSmartRoutingPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"result": {"id": "smart_routing", "value": "off", "editable": true, "modified_on": "2019-01-01T12:00:00"}, "success": true, "errors": [["Errors"]], "messages": [["Messages"]]}`)
				}))
			})
			It(`Invoke GetSmartRouting successfully`, func() {
				routingService, serviceErr := routingv1.NewRoutingV1(&routingv1.RoutingV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(routingService).ToNot(BeNil())
				routingService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := routingService.GetSmartRouting(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetSmartRoutingOptions model
				getSmartRoutingOptionsModel := new(routingv1.GetSmartRoutingOptions)
				getSmartRoutingOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = routingService.GetSmartRouting(getSmartRoutingOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = routingService.GetSmartRoutingWithContext(ctx, getSmartRoutingOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				routingService.DisableRetries()
				result, response, operationErr = routingService.GetSmartRouting(getSmartRoutingOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = routingService.GetSmartRoutingWithContext(ctx, getSmartRoutingOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke GetSmartRouting with error: Operation request error`, func() {
				routingService, serviceErr := routingv1.NewRoutingV1(&routingv1.RoutingV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(routingService).ToNot(BeNil())

				// Construct an instance of the GetSmartRoutingOptions model
				getSmartRoutingOptionsModel := new(routingv1.GetSmartRoutingOptions)
				getSmartRoutingOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := routingService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := routingService.GetSmartRouting(getSmartRoutingOptionsModel)
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
	Describe(`UpdateSmartRouting(updateSmartRoutingOptions *UpdateSmartRoutingOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		updateSmartRoutingPath := "/v1/testString/zones/testString/routing/smart_routing"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateSmartRoutingPath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateSmartRouting with error: Operation response processing error`, func() {
				routingService, serviceErr := routingv1.NewRoutingV1(&routingv1.RoutingV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(routingService).ToNot(BeNil())

				// Construct an instance of the UpdateSmartRoutingOptions model
				updateSmartRoutingOptionsModel := new(routingv1.UpdateSmartRoutingOptions)
				updateSmartRoutingOptionsModel.Value = core.StringPtr("off")
				updateSmartRoutingOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := routingService.UpdateSmartRouting(updateSmartRoutingOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				routingService.EnableRetries(0, 0)
				result, response, operationErr = routingService.UpdateSmartRouting(updateSmartRoutingOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UpdateSmartRouting(updateSmartRoutingOptions *UpdateSmartRoutingOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		updateSmartRoutingPath := "/v1/testString/zones/testString/routing/smart_routing"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateSmartRoutingPath))
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
					fmt.Fprintf(res, "%s", `{"result": {"id": "smart_routing", "value": "off", "editable": true, "modified_on": "2019-01-01T12:00:00"}, "success": true, "errors": [["Errors"]], "messages": [["Messages"]]}`)
				}))
			})
			It(`Invoke UpdateSmartRouting successfully`, func() {
				routingService, serviceErr := routingv1.NewRoutingV1(&routingv1.RoutingV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(routingService).ToNot(BeNil())
				routingService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := routingService.UpdateSmartRouting(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateSmartRoutingOptions model
				updateSmartRoutingOptionsModel := new(routingv1.UpdateSmartRoutingOptions)
				updateSmartRoutingOptionsModel.Value = core.StringPtr("off")
				updateSmartRoutingOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = routingService.UpdateSmartRouting(updateSmartRoutingOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = routingService.UpdateSmartRoutingWithContext(ctx, updateSmartRoutingOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				routingService.DisableRetries()
				result, response, operationErr = routingService.UpdateSmartRouting(updateSmartRoutingOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = routingService.UpdateSmartRoutingWithContext(ctx, updateSmartRoutingOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke UpdateSmartRouting with error: Operation request error`, func() {
				routingService, serviceErr := routingv1.NewRoutingV1(&routingv1.RoutingV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(routingService).ToNot(BeNil())

				// Construct an instance of the UpdateSmartRoutingOptions model
				updateSmartRoutingOptionsModel := new(routingv1.UpdateSmartRoutingOptions)
				updateSmartRoutingOptionsModel.Value = core.StringPtr("off")
				updateSmartRoutingOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := routingService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := routingService.UpdateSmartRouting(updateSmartRoutingOptionsModel)
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
			routingService, _ := routingv1.NewRoutingV1(&routingv1.RoutingV1Options{
				URL:            "http://routingv1modelgenerator.com",
				Authenticator:  &core.NoAuthAuthenticator{},
				Crn:            core.StringPtr(crn),
				ZoneIdentifier: core.StringPtr(zoneIdentifier),
			})
			It(`Invoke NewGetSmartRoutingOptions successfully`, func() {
				// Construct an instance of the GetSmartRoutingOptions model
				getSmartRoutingOptionsModel := routingService.NewGetSmartRoutingOptions()
				getSmartRoutingOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getSmartRoutingOptionsModel).ToNot(BeNil())
				Expect(getSmartRoutingOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateSmartRoutingOptions successfully`, func() {
				// Construct an instance of the UpdateSmartRoutingOptions model
				updateSmartRoutingOptionsModel := routingService.NewUpdateSmartRoutingOptions()
				updateSmartRoutingOptionsModel.SetValue("off")
				updateSmartRoutingOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateSmartRoutingOptionsModel).ToNot(BeNil())
				Expect(updateSmartRoutingOptionsModel.Value).To(Equal(core.StringPtr("off")))
				Expect(updateSmartRoutingOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
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
