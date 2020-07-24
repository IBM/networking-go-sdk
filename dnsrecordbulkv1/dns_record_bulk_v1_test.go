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

package dnsrecordbulkv1_test

import (
	"bytes"
	"fmt"
	"github.com/IBM/go-sdk-core/v4/core"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/IBM/networking-go-sdk/dnsrecordbulkv1"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"time"
)

var _ = Describe(`DnsRecordBulkV1`, func() {
	var testServer *httptest.Server
    Describe(`Service constructor tests`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		It(`Instantiate service client`, func() {
			testService, testServiceErr := dnsrecordbulkv1.NewDnsRecordBulkV1(&dnsrecordbulkv1.DnsRecordBulkV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
				Crn: core.StringPtr(crn),
				ZoneIdentifier: core.StringPtr(zoneIdentifier),
			})
			Expect(testService).ToNot(BeNil())
			Expect(testServiceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			testService, testServiceErr := dnsrecordbulkv1.NewDnsRecordBulkV1(&dnsrecordbulkv1.DnsRecordBulkV1Options{
				URL: "{BAD_URL_STRING",
				Crn: core.StringPtr(crn),
				ZoneIdentifier: core.StringPtr(zoneIdentifier),
			})
			Expect(testService).To(BeNil())
			Expect(testServiceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			testService, testServiceErr := dnsrecordbulkv1.NewDnsRecordBulkV1(&dnsrecordbulkv1.DnsRecordBulkV1Options{
				URL: "https://dnsrecordbulkv1/api",
				Crn: core.StringPtr(crn),
				ZoneIdentifier: core.StringPtr(zoneIdentifier),
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(testService).To(BeNil())
			Expect(testServiceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Validation Error`, func() {
			testService, testServiceErr := dnsrecordbulkv1.NewDnsRecordBulkV1(&dnsrecordbulkv1.DnsRecordBulkV1Options{})
			Expect(testService).To(BeNil())
			Expect(testServiceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"DNS_RECORD_BULK_URL": "https://dnsrecordbulkv1/api",
				"DNS_RECORD_BULK_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := dnsrecordbulkv1.NewDnsRecordBulkV1UsingExternalConfig(&dnsrecordbulkv1.DnsRecordBulkV1Options{
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := dnsrecordbulkv1.NewDnsRecordBulkV1UsingExternalConfig(&dnsrecordbulkv1.DnsRecordBulkV1Options{
					URL: "https://testService/api",
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				Expect(testService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := dnsrecordbulkv1.NewDnsRecordBulkV1UsingExternalConfig(&dnsrecordbulkv1.DnsRecordBulkV1Options{
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				err := testService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				Expect(testService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"DNS_RECORD_BULK_URL": "https://dnsrecordbulkv1/api",
				"DNS_RECORD_BULK_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			testService, testServiceErr := dnsrecordbulkv1.NewDnsRecordBulkV1UsingExternalConfig(&dnsrecordbulkv1.DnsRecordBulkV1Options{
				Crn: core.StringPtr(crn),
				ZoneIdentifier: core.StringPtr(zoneIdentifier),
			})

			It(`Instantiate service client with error`, func() {
				Expect(testService).To(BeNil())
				Expect(testServiceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"DNS_RECORD_BULK_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			testService, testServiceErr := dnsrecordbulkv1.NewDnsRecordBulkV1UsingExternalConfig(&dnsrecordbulkv1.DnsRecordBulkV1Options{
				URL: "{BAD_URL_STRING",
				Crn: core.StringPtr(crn),
				ZoneIdentifier: core.StringPtr(zoneIdentifier),
			})

			It(`Instantiate service client with error`, func() {
				Expect(testService).To(BeNil())
				Expect(testServiceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})

	Describe(`GetDnsRecordsBulk(getDnsRecordsBulkOptions *GetDnsRecordsBulkOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getDnsRecordsBulkPath := "/v1/testString/zones/testString/dns_records_bulk"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getDnsRecordsBulkPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "text/plain; charset=utf-8")
					res.WriteHeader(200)
					fmt.Fprintf(res, `"unknown property type: OperationResponse"`)
				}))
			})
			It(`Invoke GetDnsRecordsBulk successfully`, func() {
				testService, testServiceErr := dnsrecordbulkv1.NewDnsRecordBulkV1(&dnsrecordbulkv1.DnsRecordBulkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetDnsRecordsBulk(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetDnsRecordsBulkOptions model
				getDnsRecordsBulkOptionsModel := new(dnsrecordbulkv1.GetDnsRecordsBulkOptions)
 				getDnsRecordsBulkOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetDnsRecordsBulk(getDnsRecordsBulkOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke GetDnsRecordsBulk with error: Operation request error`, func() {
				testService, testServiceErr := dnsrecordbulkv1.NewDnsRecordBulkV1(&dnsrecordbulkv1.DnsRecordBulkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetDnsRecordsBulkOptions model
				getDnsRecordsBulkOptionsModel := new(dnsrecordbulkv1.GetDnsRecordsBulkOptions)
				getDnsRecordsBulkOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.GetDnsRecordsBulk(getDnsRecordsBulkOptionsModel)
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
	Describe(`PostDnsRecordsBulk(postDnsRecordsBulkOptions *PostDnsRecordsBulkOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		postDnsRecordsBulkPath := "/v1/testString/zones/testString/dns_records_bulk"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(postDnsRecordsBulkPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke PostDnsRecordsBulk with error: Operation response processing error`, func() {
				testService, testServiceErr := dnsrecordbulkv1.NewDnsRecordBulkV1(&dnsrecordbulkv1.DnsRecordBulkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the PostDnsRecordsBulkOptions model
				postDnsRecordsBulkOptionsModel := new(dnsrecordbulkv1.PostDnsRecordsBulkOptions)
				postDnsRecordsBulkOptionsModel.File = CreateMockReader("This is a mock file.")
				postDnsRecordsBulkOptionsModel.FileContentType = core.StringPtr("testString")
				postDnsRecordsBulkOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.PostDnsRecordsBulk(postDnsRecordsBulkOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`PostDnsRecordsBulk(postDnsRecordsBulkOptions *PostDnsRecordsBulkOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		postDnsRecordsBulkPath := "/v1/testString/zones/testString/dns_records_bulk"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(postDnsRecordsBulkPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"success": true, "errors": [["Errors"]], "messages": [{"code": 4, "message": "Message"}], "result": {"recs_added": 5, "total_records_parsed": 5}, "timing": {"start_time": "2014-03-01T12:20:00Z", "end_time": "2014-03-01T12:20:01Z", "process_time": 1}}`)
				}))
			})
			It(`Invoke PostDnsRecordsBulk successfully`, func() {
				testService, testServiceErr := dnsrecordbulkv1.NewDnsRecordBulkV1(&dnsrecordbulkv1.DnsRecordBulkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.PostDnsRecordsBulk(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the PostDnsRecordsBulkOptions model
				postDnsRecordsBulkOptionsModel := new(dnsrecordbulkv1.PostDnsRecordsBulkOptions)
				postDnsRecordsBulkOptionsModel.File = CreateMockReader("This is a mock file.")
				postDnsRecordsBulkOptionsModel.FileContentType = core.StringPtr("testString")
 				postDnsRecordsBulkOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.PostDnsRecordsBulk(postDnsRecordsBulkOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke PostDnsRecordsBulk with error: Param validation error`, func() {
				testService, testServiceErr := dnsrecordbulkv1.NewDnsRecordBulkV1(&dnsrecordbulkv1.DnsRecordBulkV1Options{
					URL:  testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the PostDnsRecordsBulkOptions model
				postDnsRecordsBulkOptionsModel := new(dnsrecordbulkv1.PostDnsRecordsBulkOptions)
				// Invoke operation with invalid options model (negative test)
				result, response, operationErr := testService.PostDnsRecordsBulk(postDnsRecordsBulkOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			It(`Invoke PostDnsRecordsBulk with error: Operation request error`, func() {
				testService, testServiceErr := dnsrecordbulkv1.NewDnsRecordBulkV1(&dnsrecordbulkv1.DnsRecordBulkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the PostDnsRecordsBulkOptions model
				postDnsRecordsBulkOptionsModel := new(dnsrecordbulkv1.PostDnsRecordsBulkOptions)
				postDnsRecordsBulkOptionsModel.File = CreateMockReader("This is a mock file.")
				postDnsRecordsBulkOptionsModel.FileContentType = core.StringPtr("testString")
				postDnsRecordsBulkOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.PostDnsRecordsBulk(postDnsRecordsBulkOptionsModel)
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
			testService, _ := dnsrecordbulkv1.NewDnsRecordBulkV1(&dnsrecordbulkv1.DnsRecordBulkV1Options{
				URL:           "http://dnsrecordbulkv1modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
				Crn: core.StringPtr(crn),
				ZoneIdentifier: core.StringPtr(zoneIdentifier),
			})
			It(`Invoke NewGetDnsRecordsBulkOptions successfully`, func() {
				// Construct an instance of the GetDnsRecordsBulkOptions model
				getDnsRecordsBulkOptionsModel := testService.NewGetDnsRecordsBulkOptions()
				getDnsRecordsBulkOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getDnsRecordsBulkOptionsModel).ToNot(BeNil())
				Expect(getDnsRecordsBulkOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewPostDnsRecordsBulkOptions successfully`, func() {
				// Construct an instance of the PostDnsRecordsBulkOptions model
				postDnsRecordsBulkOptionsModel := testService.NewPostDnsRecordsBulkOptions()
				postDnsRecordsBulkOptionsModel.SetFile(CreateMockReader("This is a mock file."))
				postDnsRecordsBulkOptionsModel.SetFileContentType("testString")
				postDnsRecordsBulkOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(postDnsRecordsBulkOptionsModel).ToNot(BeNil())
				Expect(postDnsRecordsBulkOptionsModel.File).To(Equal(CreateMockReader("This is a mock file.")))
				Expect(postDnsRecordsBulkOptionsModel.FileContentType).To(Equal(core.StringPtr("testString")))
				Expect(postDnsRecordsBulkOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
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
