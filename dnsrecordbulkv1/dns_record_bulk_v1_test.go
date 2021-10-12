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
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/networking-go-sdk/dnsrecordbulkv1"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe(`DnsRecordBulkV1`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		It(`Instantiate service client`, func() {
			dnsRecordBulkService, serviceErr := dnsrecordbulkv1.NewDnsRecordBulkV1(&dnsrecordbulkv1.DnsRecordBulkV1Options{
				Authenticator:  &core.NoAuthAuthenticator{},
				Crn:            core.StringPtr(crn),
				ZoneIdentifier: core.StringPtr(zoneIdentifier),
			})
			Expect(dnsRecordBulkService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			dnsRecordBulkService, serviceErr := dnsrecordbulkv1.NewDnsRecordBulkV1(&dnsrecordbulkv1.DnsRecordBulkV1Options{
				URL:            "{BAD_URL_STRING",
				Crn:            core.StringPtr(crn),
				ZoneIdentifier: core.StringPtr(zoneIdentifier),
			})
			Expect(dnsRecordBulkService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			dnsRecordBulkService, serviceErr := dnsrecordbulkv1.NewDnsRecordBulkV1(&dnsrecordbulkv1.DnsRecordBulkV1Options{
				URL:            "https://dnsrecordbulkv1/api",
				Crn:            core.StringPtr(crn),
				ZoneIdentifier: core.StringPtr(zoneIdentifier),
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(dnsRecordBulkService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Validation Error`, func() {
			dnsRecordBulkService, serviceErr := dnsrecordbulkv1.NewDnsRecordBulkV1(&dnsrecordbulkv1.DnsRecordBulkV1Options{})
			Expect(dnsRecordBulkService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"DNS_RECORD_BULK_URL":       "https://dnsrecordbulkv1/api",
				"DNS_RECORD_BULK_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				dnsRecordBulkService, serviceErr := dnsrecordbulkv1.NewDnsRecordBulkV1UsingExternalConfig(&dnsrecordbulkv1.DnsRecordBulkV1Options{
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(dnsRecordBulkService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := dnsRecordBulkService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != dnsRecordBulkService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(dnsRecordBulkService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(dnsRecordBulkService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				dnsRecordBulkService, serviceErr := dnsrecordbulkv1.NewDnsRecordBulkV1UsingExternalConfig(&dnsrecordbulkv1.DnsRecordBulkV1Options{
					URL:            "https://testService/api",
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(dnsRecordBulkService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(dnsRecordBulkService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := dnsRecordBulkService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != dnsRecordBulkService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(dnsRecordBulkService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(dnsRecordBulkService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				dnsRecordBulkService, serviceErr := dnsrecordbulkv1.NewDnsRecordBulkV1UsingExternalConfig(&dnsrecordbulkv1.DnsRecordBulkV1Options{
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				err := dnsRecordBulkService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(dnsRecordBulkService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(dnsRecordBulkService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := dnsRecordBulkService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != dnsRecordBulkService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(dnsRecordBulkService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(dnsRecordBulkService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"DNS_RECORD_BULK_URL":       "https://dnsrecordbulkv1/api",
				"DNS_RECORD_BULK_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			dnsRecordBulkService, serviceErr := dnsrecordbulkv1.NewDnsRecordBulkV1UsingExternalConfig(&dnsrecordbulkv1.DnsRecordBulkV1Options{
				Crn:            core.StringPtr(crn),
				ZoneIdentifier: core.StringPtr(zoneIdentifier),
			})

			It(`Instantiate service client with error`, func() {
				Expect(dnsRecordBulkService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"DNS_RECORD_BULK_AUTH_TYPE": "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			dnsRecordBulkService, serviceErr := dnsrecordbulkv1.NewDnsRecordBulkV1UsingExternalConfig(&dnsrecordbulkv1.DnsRecordBulkV1Options{
				URL:            "{BAD_URL_STRING",
				Crn:            core.StringPtr(crn),
				ZoneIdentifier: core.StringPtr(zoneIdentifier),
			})

			It(`Instantiate service client with error`, func() {
				Expect(dnsRecordBulkService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = dnsrecordbulkv1.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})

	Describe(`GetDnsRecordsBulk(getDnsRecordsBulkOptions *GetDnsRecordsBulkOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getDnsRecordsBulkPath := "/v1/testString/zones/testString/dns_records_bulk"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getDnsRecordsBulkPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "text/plain; charset=utf-8")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `This is a mock binary response.`)
				}))
			})
			It(`Invoke GetDnsRecordsBulk successfully`, func() {
				dnsRecordBulkService, serviceErr := dnsrecordbulkv1.NewDnsRecordBulkV1(&dnsrecordbulkv1.DnsRecordBulkV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsRecordBulkService).ToNot(BeNil())
				dnsRecordBulkService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := dnsRecordBulkService.GetDnsRecordsBulk(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetDnsRecordsBulkOptions model
				getDnsRecordsBulkOptionsModel := new(dnsrecordbulkv1.GetDnsRecordsBulkOptions)
				getDnsRecordsBulkOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = dnsRecordBulkService.GetDnsRecordsBulk(getDnsRecordsBulkOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = dnsRecordBulkService.GetDnsRecordsBulkWithContext(ctx, getDnsRecordsBulkOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				dnsRecordBulkService.DisableRetries()
				result, response, operationErr = dnsRecordBulkService.GetDnsRecordsBulk(getDnsRecordsBulkOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = dnsRecordBulkService.GetDnsRecordsBulkWithContext(ctx, getDnsRecordsBulkOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke GetDnsRecordsBulk with error: Operation request error`, func() {
				dnsRecordBulkService, serviceErr := dnsrecordbulkv1.NewDnsRecordBulkV1(&dnsrecordbulkv1.DnsRecordBulkV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsRecordBulkService).ToNot(BeNil())

				// Construct an instance of the GetDnsRecordsBulkOptions model
				getDnsRecordsBulkOptionsModel := new(dnsrecordbulkv1.GetDnsRecordsBulkOptions)
				getDnsRecordsBulkOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := dnsRecordBulkService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := dnsRecordBulkService.GetDnsRecordsBulk(getDnsRecordsBulkOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(postDnsRecordsBulkPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke PostDnsRecordsBulk with error: Operation response processing error`, func() {
				dnsRecordBulkService, serviceErr := dnsrecordbulkv1.NewDnsRecordBulkV1(&dnsrecordbulkv1.DnsRecordBulkV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsRecordBulkService).ToNot(BeNil())

				// Construct an instance of the PostDnsRecordsBulkOptions model
				postDnsRecordsBulkOptionsModel := new(dnsrecordbulkv1.PostDnsRecordsBulkOptions)
				postDnsRecordsBulkOptionsModel.File = CreateMockReader("This is a mock file.")
				postDnsRecordsBulkOptionsModel.FileContentType = core.StringPtr("testString")
				postDnsRecordsBulkOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := dnsRecordBulkService.PostDnsRecordsBulk(postDnsRecordsBulkOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				dnsRecordBulkService.EnableRetries(0, 0)
				result, response, operationErr = dnsRecordBulkService.PostDnsRecordsBulk(postDnsRecordsBulkOptionsModel)
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
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(postDnsRecordsBulkPath))
					Expect(req.Method).To(Equal("POST"))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [{"code": 4, "message": "Message"}], "result": {"recs_added": 5, "total_records_parsed": 5}, "timing": {"start_time": "2014-03-01T12:20:00Z", "end_time": "2014-03-01T12:20:01Z", "process_time": 1}}`)
				}))
			})
			It(`Invoke PostDnsRecordsBulk successfully`, func() {
				dnsRecordBulkService, serviceErr := dnsrecordbulkv1.NewDnsRecordBulkV1(&dnsrecordbulkv1.DnsRecordBulkV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsRecordBulkService).ToNot(BeNil())
				dnsRecordBulkService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := dnsRecordBulkService.PostDnsRecordsBulk(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the PostDnsRecordsBulkOptions model
				postDnsRecordsBulkOptionsModel := new(dnsrecordbulkv1.PostDnsRecordsBulkOptions)
				postDnsRecordsBulkOptionsModel.File = CreateMockReader("This is a mock file.")
				postDnsRecordsBulkOptionsModel.FileContentType = core.StringPtr("testString")
				postDnsRecordsBulkOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = dnsRecordBulkService.PostDnsRecordsBulk(postDnsRecordsBulkOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = dnsRecordBulkService.PostDnsRecordsBulkWithContext(ctx, postDnsRecordsBulkOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				dnsRecordBulkService.DisableRetries()
				result, response, operationErr = dnsRecordBulkService.PostDnsRecordsBulk(postDnsRecordsBulkOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = dnsRecordBulkService.PostDnsRecordsBulkWithContext(ctx, postDnsRecordsBulkOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke PostDnsRecordsBulk with error: Param validation error`, func() {
				dnsRecordBulkService, serviceErr := dnsrecordbulkv1.NewDnsRecordBulkV1(&dnsrecordbulkv1.DnsRecordBulkV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsRecordBulkService).ToNot(BeNil())

				// Construct an instance of the PostDnsRecordsBulkOptions model
				postDnsRecordsBulkOptionsModel := new(dnsrecordbulkv1.PostDnsRecordsBulkOptions)
				// Invoke operation with invalid options model (negative test)
				result, response, operationErr := dnsRecordBulkService.PostDnsRecordsBulk(postDnsRecordsBulkOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			It(`Invoke PostDnsRecordsBulk with error: Operation request error`, func() {
				dnsRecordBulkService, serviceErr := dnsrecordbulkv1.NewDnsRecordBulkV1(&dnsrecordbulkv1.DnsRecordBulkV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsRecordBulkService).ToNot(BeNil())

				// Construct an instance of the PostDnsRecordsBulkOptions model
				postDnsRecordsBulkOptionsModel := new(dnsrecordbulkv1.PostDnsRecordsBulkOptions)
				postDnsRecordsBulkOptionsModel.File = CreateMockReader("This is a mock file.")
				postDnsRecordsBulkOptionsModel.FileContentType = core.StringPtr("testString")
				postDnsRecordsBulkOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := dnsRecordBulkService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := dnsRecordBulkService.PostDnsRecordsBulk(postDnsRecordsBulkOptionsModel)
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
			dnsRecordBulkService, _ := dnsrecordbulkv1.NewDnsRecordBulkV1(&dnsrecordbulkv1.DnsRecordBulkV1Options{
				URL:            "http://dnsrecordbulkv1modelgenerator.com",
				Authenticator:  &core.NoAuthAuthenticator{},
				Crn:            core.StringPtr(crn),
				ZoneIdentifier: core.StringPtr(zoneIdentifier),
			})
			It(`Invoke NewGetDnsRecordsBulkOptions successfully`, func() {
				// Construct an instance of the GetDnsRecordsBulkOptions model
				getDnsRecordsBulkOptionsModel := dnsRecordBulkService.NewGetDnsRecordsBulkOptions()
				getDnsRecordsBulkOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getDnsRecordsBulkOptionsModel).ToNot(BeNil())
				Expect(getDnsRecordsBulkOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewPostDnsRecordsBulkOptions successfully`, func() {
				// Construct an instance of the PostDnsRecordsBulkOptions model
				postDnsRecordsBulkOptionsModel := dnsRecordBulkService.NewPostDnsRecordsBulkOptions()
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
