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

package cisipapiv1_test

import (
	"bytes"
	"context"
	"fmt"
	"github.com/IBM/go-sdk-core/v4/core"
	"github.com/IBM/networking-go-sdk/cisipapiv1"
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

var _ = Describe(`CisIpApiV1`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			cisIpApiService, serviceErr := cisipapiv1.NewCisIpApiV1(&cisipapiv1.CisIpApiV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(cisIpApiService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			cisIpApiService, serviceErr := cisipapiv1.NewCisIpApiV1(&cisipapiv1.CisIpApiV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(cisIpApiService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			cisIpApiService, serviceErr := cisipapiv1.NewCisIpApiV1(&cisipapiv1.CisIpApiV1Options{
				URL: "https://cisipapiv1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(cisIpApiService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"CIS_IP_API_URL": "https://cisipapiv1/api",
				"CIS_IP_API_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				cisIpApiService, serviceErr := cisipapiv1.NewCisIpApiV1UsingExternalConfig(&cisipapiv1.CisIpApiV1Options{
				})
				Expect(cisIpApiService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := cisIpApiService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != cisIpApiService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(cisIpApiService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(cisIpApiService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				cisIpApiService, serviceErr := cisipapiv1.NewCisIpApiV1UsingExternalConfig(&cisipapiv1.CisIpApiV1Options{
					URL: "https://testService/api",
				})
				Expect(cisIpApiService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(cisIpApiService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := cisIpApiService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != cisIpApiService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(cisIpApiService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(cisIpApiService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				cisIpApiService, serviceErr := cisipapiv1.NewCisIpApiV1UsingExternalConfig(&cisipapiv1.CisIpApiV1Options{
				})
				err := cisIpApiService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(cisIpApiService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(cisIpApiService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := cisIpApiService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != cisIpApiService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(cisIpApiService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(cisIpApiService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"CIS_IP_API_URL": "https://cisipapiv1/api",
				"CIS_IP_API_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			cisIpApiService, serviceErr := cisipapiv1.NewCisIpApiV1UsingExternalConfig(&cisipapiv1.CisIpApiV1Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(cisIpApiService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"CIS_IP_API_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			cisIpApiService, serviceErr := cisipapiv1.NewCisIpApiV1UsingExternalConfig(&cisipapiv1.CisIpApiV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(cisIpApiService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = cisipapiv1.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`ListIps(listIpsOptions *ListIpsOptions) - Operation response error`, func() {
		listIpsPath := "/v1/ips"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listIpsPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListIps with error: Operation response processing error`, func() {
				cisIpApiService, serviceErr := cisipapiv1.NewCisIpApiV1(&cisipapiv1.CisIpApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cisIpApiService).ToNot(BeNil())

				// Construct an instance of the ListIpsOptions model
				listIpsOptionsModel := new(cisipapiv1.ListIpsOptions)
				listIpsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cisIpApiService.ListIps(listIpsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				cisIpApiService.EnableRetries(0, 0)
				result, response, operationErr = cisIpApiService.ListIps(listIpsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ListIps(listIpsOptions *ListIpsOptions)`, func() {
		listIpsPath := "/v1/ips"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listIpsPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"ipv4_cidrs": ["180.15.128.0/20"], "ipv6_cidrs": ["2400:cb00::/32"]}}`)
				}))
			})
			It(`Invoke ListIps successfully`, func() {
				cisIpApiService, serviceErr := cisipapiv1.NewCisIpApiV1(&cisipapiv1.CisIpApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cisIpApiService).ToNot(BeNil())
				cisIpApiService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cisIpApiService.ListIps(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListIpsOptions model
				listIpsOptionsModel := new(cisipapiv1.ListIpsOptions)
				listIpsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cisIpApiService.ListIps(listIpsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = cisIpApiService.ListIpsWithContext(ctx, listIpsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				cisIpApiService.DisableRetries()
				result, response, operationErr = cisIpApiService.ListIps(listIpsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = cisIpApiService.ListIpsWithContext(ctx, listIpsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke ListIps with error: Operation request error`, func() {
				cisIpApiService, serviceErr := cisipapiv1.NewCisIpApiV1(&cisipapiv1.CisIpApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cisIpApiService).ToNot(BeNil())

				// Construct an instance of the ListIpsOptions model
				listIpsOptionsModel := new(cisipapiv1.ListIpsOptions)
				listIpsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cisIpApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cisIpApiService.ListIps(listIpsOptionsModel)
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
			cisIpApiService, _ := cisipapiv1.NewCisIpApiV1(&cisipapiv1.CisIpApiV1Options{
				URL:           "http://cisipapiv1modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
			})
			It(`Invoke NewListIpsOptions successfully`, func() {
				// Construct an instance of the ListIpsOptions model
				listIpsOptionsModel := cisIpApiService.NewListIpsOptions()
				listIpsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listIpsOptionsModel).ToNot(BeNil())
				Expect(listIpsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
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
