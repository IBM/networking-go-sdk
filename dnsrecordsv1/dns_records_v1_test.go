/**
 * (C) Copyright IBM Corp. 2026.
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

package dnsrecordsv1_test

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/networking-go-sdk/dnsrecordsv1"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe(`DnsRecordsV1`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		It(`Instantiate service client`, func() {
			dnsRecordsService, serviceErr := dnsrecordsv1.NewDnsRecordsV1(&dnsrecordsv1.DnsRecordsV1Options{
				Authenticator:  &core.NoAuthAuthenticator{},
				Crn:            core.StringPtr(crn),
				ZoneIdentifier: core.StringPtr(zoneIdentifier),
			})
			Expect(dnsRecordsService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			dnsRecordsService, serviceErr := dnsrecordsv1.NewDnsRecordsV1(&dnsrecordsv1.DnsRecordsV1Options{
				URL:            "{BAD_URL_STRING",
				Crn:            core.StringPtr(crn),
				ZoneIdentifier: core.StringPtr(zoneIdentifier),
			})
			Expect(dnsRecordsService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			dnsRecordsService, serviceErr := dnsrecordsv1.NewDnsRecordsV1(&dnsrecordsv1.DnsRecordsV1Options{
				URL:            "https://dnsrecordsv1/api",
				Crn:            core.StringPtr(crn),
				ZoneIdentifier: core.StringPtr(zoneIdentifier),
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(dnsRecordsService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Validation Error`, func() {
			dnsRecordsService, serviceErr := dnsrecordsv1.NewDnsRecordsV1(&dnsrecordsv1.DnsRecordsV1Options{})
			Expect(dnsRecordsService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"DNS_RECORDS_URL":       "https://dnsrecordsv1/api",
				"DNS_RECORDS_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				dnsRecordsService, serviceErr := dnsrecordsv1.NewDnsRecordsV1UsingExternalConfig(&dnsrecordsv1.DnsRecordsV1Options{
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(dnsRecordsService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := dnsRecordsService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != dnsRecordsService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(dnsRecordsService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(dnsRecordsService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				dnsRecordsService, serviceErr := dnsrecordsv1.NewDnsRecordsV1UsingExternalConfig(&dnsrecordsv1.DnsRecordsV1Options{
					URL:            "https://testService/api",
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(dnsRecordsService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(dnsRecordsService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := dnsRecordsService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != dnsRecordsService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(dnsRecordsService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(dnsRecordsService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				dnsRecordsService, serviceErr := dnsrecordsv1.NewDnsRecordsV1UsingExternalConfig(&dnsrecordsv1.DnsRecordsV1Options{
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				err := dnsRecordsService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(dnsRecordsService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(dnsRecordsService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := dnsRecordsService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != dnsRecordsService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(dnsRecordsService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(dnsRecordsService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"DNS_RECORDS_URL":       "https://dnsrecordsv1/api",
				"DNS_RECORDS_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			dnsRecordsService, serviceErr := dnsrecordsv1.NewDnsRecordsV1UsingExternalConfig(&dnsrecordsv1.DnsRecordsV1Options{
				Crn:            core.StringPtr(crn),
				ZoneIdentifier: core.StringPtr(zoneIdentifier),
			})

			It(`Instantiate service client with error`, func() {
				Expect(dnsRecordsService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"DNS_RECORDS_AUTH_TYPE": "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			dnsRecordsService, serviceErr := dnsrecordsv1.NewDnsRecordsV1UsingExternalConfig(&dnsrecordsv1.DnsRecordsV1Options{
				URL:            "{BAD_URL_STRING",
				Crn:            core.StringPtr(crn),
				ZoneIdentifier: core.StringPtr(zoneIdentifier),
			})

			It(`Instantiate service client with error`, func() {
				Expect(dnsRecordsService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = dnsrecordsv1.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`ListAllDnsRecords(listAllDnsRecordsOptions *ListAllDnsRecordsOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		listAllDnsRecordsPath := "/v1/testString/zones/testString/dns_records"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listAllDnsRecordsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["type"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["name"]).To(Equal([]string{"host1.test-example.com"}))
					Expect(req.URL.Query()["content"]).To(Equal([]string{"1.2.3.4"}))
					Expect(req.URL.Query()["page"]).To(Equal([]string{fmt.Sprint(int64(1))}))
					Expect(req.URL.Query()["per_page"]).To(Equal([]string{fmt.Sprint(int64(20))}))
					Expect(req.URL.Query()["order"]).To(Equal([]string{"type"}))
					Expect(req.URL.Query()["direction"]).To(Equal([]string{"asc"}))
					Expect(req.URL.Query()["match"]).To(Equal([]string{"any"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListAllDnsRecords with error: Operation response processing error`, func() {
				dnsRecordsService, serviceErr := dnsrecordsv1.NewDnsRecordsV1(&dnsrecordsv1.DnsRecordsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsRecordsService).ToNot(BeNil())

				// Construct an instance of the ListAllDnsRecordsOptions model
				listAllDnsRecordsOptionsModel := new(dnsrecordsv1.ListAllDnsRecordsOptions)
				listAllDnsRecordsOptionsModel.Type = core.StringPtr("testString")
				listAllDnsRecordsOptionsModel.Name = core.StringPtr("host1.test-example.com")
				listAllDnsRecordsOptionsModel.Content = core.StringPtr("1.2.3.4")
				listAllDnsRecordsOptionsModel.Page = core.Int64Ptr(int64(1))
				listAllDnsRecordsOptionsModel.PerPage = core.Int64Ptr(int64(20))
				listAllDnsRecordsOptionsModel.Order = core.StringPtr("type")
				listAllDnsRecordsOptionsModel.Direction = core.StringPtr("asc")
				listAllDnsRecordsOptionsModel.Match = core.StringPtr("any")
				listAllDnsRecordsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := dnsRecordsService.ListAllDnsRecords(listAllDnsRecordsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				dnsRecordsService.EnableRetries(0, 0)
				result, response, operationErr = dnsRecordsService.ListAllDnsRecords(listAllDnsRecordsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListAllDnsRecords(listAllDnsRecordsOptions *ListAllDnsRecordsOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		listAllDnsRecordsPath := "/v1/testString/zones/testString/dns_records"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listAllDnsRecordsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["type"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["name"]).To(Equal([]string{"host1.test-example.com"}))
					Expect(req.URL.Query()["content"]).To(Equal([]string{"1.2.3.4"}))
					Expect(req.URL.Query()["page"]).To(Equal([]string{fmt.Sprint(int64(1))}))
					Expect(req.URL.Query()["per_page"]).To(Equal([]string{fmt.Sprint(int64(20))}))
					Expect(req.URL.Query()["order"]).To(Equal([]string{"type"}))
					Expect(req.URL.Query()["direction"]).To(Equal([]string{"asc"}))
					Expect(req.URL.Query()["match"]).To(Equal([]string{"any"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": [{"id": "f1aba936b94213e5b8dca0c0dbf1f9cc", "created_on": "2014-01-01T05:20:00.12345Z", "modified_on": "2014-01-01T05:20:00.12345Z", "name": "host-1.test-example.com", "type": "A", "content": "169.154.10.10", "zone_id": "023e105f4ecef8ad9ca31a8372d0c353", "zone_name": "test-example.com", "proxiable": true, "proxied": false, "ttl": 120, "priority": 5, "data": {"anyKey": "anyValue"}}], "result_info": {"page": 1, "per_page": 20, "count": 1, "total_count": 2000}}`)
				}))
			})
			It(`Invoke ListAllDnsRecords successfully with retries`, func() {
				dnsRecordsService, serviceErr := dnsrecordsv1.NewDnsRecordsV1(&dnsrecordsv1.DnsRecordsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsRecordsService).ToNot(BeNil())
				dnsRecordsService.EnableRetries(0, 0)

				// Construct an instance of the ListAllDnsRecordsOptions model
				listAllDnsRecordsOptionsModel := new(dnsrecordsv1.ListAllDnsRecordsOptions)
				listAllDnsRecordsOptionsModel.Type = core.StringPtr("testString")
				listAllDnsRecordsOptionsModel.Name = core.StringPtr("host1.test-example.com")
				listAllDnsRecordsOptionsModel.Content = core.StringPtr("1.2.3.4")
				listAllDnsRecordsOptionsModel.Page = core.Int64Ptr(int64(1))
				listAllDnsRecordsOptionsModel.PerPage = core.Int64Ptr(int64(20))
				listAllDnsRecordsOptionsModel.Order = core.StringPtr("type")
				listAllDnsRecordsOptionsModel.Direction = core.StringPtr("asc")
				listAllDnsRecordsOptionsModel.Match = core.StringPtr("any")
				listAllDnsRecordsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := dnsRecordsService.ListAllDnsRecordsWithContext(ctx, listAllDnsRecordsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				dnsRecordsService.DisableRetries()
				result, response, operationErr := dnsRecordsService.ListAllDnsRecords(listAllDnsRecordsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = dnsRecordsService.ListAllDnsRecordsWithContext(ctx, listAllDnsRecordsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listAllDnsRecordsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["type"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["name"]).To(Equal([]string{"host1.test-example.com"}))
					Expect(req.URL.Query()["content"]).To(Equal([]string{"1.2.3.4"}))
					Expect(req.URL.Query()["page"]).To(Equal([]string{fmt.Sprint(int64(1))}))
					Expect(req.URL.Query()["per_page"]).To(Equal([]string{fmt.Sprint(int64(20))}))
					Expect(req.URL.Query()["order"]).To(Equal([]string{"type"}))
					Expect(req.URL.Query()["direction"]).To(Equal([]string{"asc"}))
					Expect(req.URL.Query()["match"]).To(Equal([]string{"any"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": [{"id": "f1aba936b94213e5b8dca0c0dbf1f9cc", "created_on": "2014-01-01T05:20:00.12345Z", "modified_on": "2014-01-01T05:20:00.12345Z", "name": "host-1.test-example.com", "type": "A", "content": "169.154.10.10", "zone_id": "023e105f4ecef8ad9ca31a8372d0c353", "zone_name": "test-example.com", "proxiable": true, "proxied": false, "ttl": 120, "priority": 5, "data": {"anyKey": "anyValue"}}], "result_info": {"page": 1, "per_page": 20, "count": 1, "total_count": 2000}}`)
				}))
			})
			It(`Invoke ListAllDnsRecords successfully`, func() {
				dnsRecordsService, serviceErr := dnsrecordsv1.NewDnsRecordsV1(&dnsrecordsv1.DnsRecordsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsRecordsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := dnsRecordsService.ListAllDnsRecords(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListAllDnsRecordsOptions model
				listAllDnsRecordsOptionsModel := new(dnsrecordsv1.ListAllDnsRecordsOptions)
				listAllDnsRecordsOptionsModel.Type = core.StringPtr("testString")
				listAllDnsRecordsOptionsModel.Name = core.StringPtr("host1.test-example.com")
				listAllDnsRecordsOptionsModel.Content = core.StringPtr("1.2.3.4")
				listAllDnsRecordsOptionsModel.Page = core.Int64Ptr(int64(1))
				listAllDnsRecordsOptionsModel.PerPage = core.Int64Ptr(int64(20))
				listAllDnsRecordsOptionsModel.Order = core.StringPtr("type")
				listAllDnsRecordsOptionsModel.Direction = core.StringPtr("asc")
				listAllDnsRecordsOptionsModel.Match = core.StringPtr("any")
				listAllDnsRecordsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = dnsRecordsService.ListAllDnsRecords(listAllDnsRecordsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListAllDnsRecords with error: Operation request error`, func() {
				dnsRecordsService, serviceErr := dnsrecordsv1.NewDnsRecordsV1(&dnsrecordsv1.DnsRecordsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsRecordsService).ToNot(BeNil())

				// Construct an instance of the ListAllDnsRecordsOptions model
				listAllDnsRecordsOptionsModel := new(dnsrecordsv1.ListAllDnsRecordsOptions)
				listAllDnsRecordsOptionsModel.Type = core.StringPtr("testString")
				listAllDnsRecordsOptionsModel.Name = core.StringPtr("host1.test-example.com")
				listAllDnsRecordsOptionsModel.Content = core.StringPtr("1.2.3.4")
				listAllDnsRecordsOptionsModel.Page = core.Int64Ptr(int64(1))
				listAllDnsRecordsOptionsModel.PerPage = core.Int64Ptr(int64(20))
				listAllDnsRecordsOptionsModel.Order = core.StringPtr("type")
				listAllDnsRecordsOptionsModel.Direction = core.StringPtr("asc")
				listAllDnsRecordsOptionsModel.Match = core.StringPtr("any")
				listAllDnsRecordsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := dnsRecordsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := dnsRecordsService.ListAllDnsRecords(listAllDnsRecordsOptionsModel)
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
			It(`Invoke ListAllDnsRecords successfully`, func() {
				dnsRecordsService, serviceErr := dnsrecordsv1.NewDnsRecordsV1(&dnsrecordsv1.DnsRecordsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsRecordsService).ToNot(BeNil())

				// Construct an instance of the ListAllDnsRecordsOptions model
				listAllDnsRecordsOptionsModel := new(dnsrecordsv1.ListAllDnsRecordsOptions)
				listAllDnsRecordsOptionsModel.Type = core.StringPtr("testString")
				listAllDnsRecordsOptionsModel.Name = core.StringPtr("host1.test-example.com")
				listAllDnsRecordsOptionsModel.Content = core.StringPtr("1.2.3.4")
				listAllDnsRecordsOptionsModel.Page = core.Int64Ptr(int64(1))
				listAllDnsRecordsOptionsModel.PerPage = core.Int64Ptr(int64(20))
				listAllDnsRecordsOptionsModel.Order = core.StringPtr("type")
				listAllDnsRecordsOptionsModel.Direction = core.StringPtr("asc")
				listAllDnsRecordsOptionsModel.Match = core.StringPtr("any")
				listAllDnsRecordsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := dnsRecordsService.ListAllDnsRecords(listAllDnsRecordsOptionsModel)
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
	Describe(`CreateDnsRecord(createDnsRecordOptions *CreateDnsRecordOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		createDnsRecordPath := "/v1/testString/zones/testString/dns_records"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createDnsRecordPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateDnsRecord with error: Operation response processing error`, func() {
				dnsRecordsService, serviceErr := dnsrecordsv1.NewDnsRecordsV1(&dnsrecordsv1.DnsRecordsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsRecordsService).ToNot(BeNil())

				// Construct an instance of the CreateDnsRecordOptions model
				createDnsRecordOptionsModel := new(dnsrecordsv1.CreateDnsRecordOptions)
				createDnsRecordOptionsModel.Name = core.StringPtr("host-1.test-example.com")
				createDnsRecordOptionsModel.Type = core.StringPtr("A")
				createDnsRecordOptionsModel.TTL = core.Int64Ptr(int64(120))
				createDnsRecordOptionsModel.Content = core.StringPtr("1.2.3.4")
				createDnsRecordOptionsModel.Priority = core.Int64Ptr(int64(5))
				createDnsRecordOptionsModel.Proxied = core.BoolPtr(false)
				createDnsRecordOptionsModel.Data = map[string]interface{}{"anyKey": "anyValue"}
				createDnsRecordOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := dnsRecordsService.CreateDnsRecord(createDnsRecordOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				dnsRecordsService.EnableRetries(0, 0)
				result, response, operationErr = dnsRecordsService.CreateDnsRecord(createDnsRecordOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateDnsRecord(createDnsRecordOptions *CreateDnsRecordOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		createDnsRecordPath := "/v1/testString/zones/testString/dns_records"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createDnsRecordPath))
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
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": "f1aba936b94213e5b8dca0c0dbf1f9cc", "created_on": "2014-01-01T05:20:00.12345Z", "modified_on": "2014-01-01T05:20:00.12345Z", "name": "host-1.test-example.com", "type": "A", "content": "169.154.10.10", "zone_id": "023e105f4ecef8ad9ca31a8372d0c353", "zone_name": "test-example.com", "proxiable": true, "proxied": false, "ttl": 120, "priority": 5, "data": {"anyKey": "anyValue"}}}`)
				}))
			})
			It(`Invoke CreateDnsRecord successfully with retries`, func() {
				dnsRecordsService, serviceErr := dnsrecordsv1.NewDnsRecordsV1(&dnsrecordsv1.DnsRecordsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsRecordsService).ToNot(BeNil())
				dnsRecordsService.EnableRetries(0, 0)

				// Construct an instance of the CreateDnsRecordOptions model
				createDnsRecordOptionsModel := new(dnsrecordsv1.CreateDnsRecordOptions)
				createDnsRecordOptionsModel.Name = core.StringPtr("host-1.test-example.com")
				createDnsRecordOptionsModel.Type = core.StringPtr("A")
				createDnsRecordOptionsModel.TTL = core.Int64Ptr(int64(120))
				createDnsRecordOptionsModel.Content = core.StringPtr("1.2.3.4")
				createDnsRecordOptionsModel.Priority = core.Int64Ptr(int64(5))
				createDnsRecordOptionsModel.Proxied = core.BoolPtr(false)
				createDnsRecordOptionsModel.Data = map[string]interface{}{"anyKey": "anyValue"}
				createDnsRecordOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := dnsRecordsService.CreateDnsRecordWithContext(ctx, createDnsRecordOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				dnsRecordsService.DisableRetries()
				result, response, operationErr := dnsRecordsService.CreateDnsRecord(createDnsRecordOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = dnsRecordsService.CreateDnsRecordWithContext(ctx, createDnsRecordOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createDnsRecordPath))
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
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": "f1aba936b94213e5b8dca0c0dbf1f9cc", "created_on": "2014-01-01T05:20:00.12345Z", "modified_on": "2014-01-01T05:20:00.12345Z", "name": "host-1.test-example.com", "type": "A", "content": "169.154.10.10", "zone_id": "023e105f4ecef8ad9ca31a8372d0c353", "zone_name": "test-example.com", "proxiable": true, "proxied": false, "ttl": 120, "priority": 5, "data": {"anyKey": "anyValue"}}}`)
				}))
			})
			It(`Invoke CreateDnsRecord successfully`, func() {
				dnsRecordsService, serviceErr := dnsrecordsv1.NewDnsRecordsV1(&dnsrecordsv1.DnsRecordsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsRecordsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := dnsRecordsService.CreateDnsRecord(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateDnsRecordOptions model
				createDnsRecordOptionsModel := new(dnsrecordsv1.CreateDnsRecordOptions)
				createDnsRecordOptionsModel.Name = core.StringPtr("host-1.test-example.com")
				createDnsRecordOptionsModel.Type = core.StringPtr("A")
				createDnsRecordOptionsModel.TTL = core.Int64Ptr(int64(120))
				createDnsRecordOptionsModel.Content = core.StringPtr("1.2.3.4")
				createDnsRecordOptionsModel.Priority = core.Int64Ptr(int64(5))
				createDnsRecordOptionsModel.Proxied = core.BoolPtr(false)
				createDnsRecordOptionsModel.Data = map[string]interface{}{"anyKey": "anyValue"}
				createDnsRecordOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = dnsRecordsService.CreateDnsRecord(createDnsRecordOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateDnsRecord with error: Operation request error`, func() {
				dnsRecordsService, serviceErr := dnsrecordsv1.NewDnsRecordsV1(&dnsrecordsv1.DnsRecordsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsRecordsService).ToNot(BeNil())

				// Construct an instance of the CreateDnsRecordOptions model
				createDnsRecordOptionsModel := new(dnsrecordsv1.CreateDnsRecordOptions)
				createDnsRecordOptionsModel.Name = core.StringPtr("host-1.test-example.com")
				createDnsRecordOptionsModel.Type = core.StringPtr("A")
				createDnsRecordOptionsModel.TTL = core.Int64Ptr(int64(120))
				createDnsRecordOptionsModel.Content = core.StringPtr("1.2.3.4")
				createDnsRecordOptionsModel.Priority = core.Int64Ptr(int64(5))
				createDnsRecordOptionsModel.Proxied = core.BoolPtr(false)
				createDnsRecordOptionsModel.Data = map[string]interface{}{"anyKey": "anyValue"}
				createDnsRecordOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := dnsRecordsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := dnsRecordsService.CreateDnsRecord(createDnsRecordOptionsModel)
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
			It(`Invoke CreateDnsRecord successfully`, func() {
				dnsRecordsService, serviceErr := dnsrecordsv1.NewDnsRecordsV1(&dnsrecordsv1.DnsRecordsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsRecordsService).ToNot(BeNil())

				// Construct an instance of the CreateDnsRecordOptions model
				createDnsRecordOptionsModel := new(dnsrecordsv1.CreateDnsRecordOptions)
				createDnsRecordOptionsModel.Name = core.StringPtr("host-1.test-example.com")
				createDnsRecordOptionsModel.Type = core.StringPtr("A")
				createDnsRecordOptionsModel.TTL = core.Int64Ptr(int64(120))
				createDnsRecordOptionsModel.Content = core.StringPtr("1.2.3.4")
				createDnsRecordOptionsModel.Priority = core.Int64Ptr(int64(5))
				createDnsRecordOptionsModel.Proxied = core.BoolPtr(false)
				createDnsRecordOptionsModel.Data = map[string]interface{}{"anyKey": "anyValue"}
				createDnsRecordOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := dnsRecordsService.CreateDnsRecord(createDnsRecordOptionsModel)
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
	Describe(`DeleteDnsRecord(deleteDnsRecordOptions *DeleteDnsRecordOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		deleteDnsRecordPath := "/v1/testString/zones/testString/dns_records/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteDnsRecordPath))
					Expect(req.Method).To(Equal("DELETE"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke DeleteDnsRecord with error: Operation response processing error`, func() {
				dnsRecordsService, serviceErr := dnsrecordsv1.NewDnsRecordsV1(&dnsrecordsv1.DnsRecordsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsRecordsService).ToNot(BeNil())

				// Construct an instance of the DeleteDnsRecordOptions model
				deleteDnsRecordOptionsModel := new(dnsrecordsv1.DeleteDnsRecordOptions)
				deleteDnsRecordOptionsModel.DnsrecordIdentifier = core.StringPtr("testString")
				deleteDnsRecordOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := dnsRecordsService.DeleteDnsRecord(deleteDnsRecordOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				dnsRecordsService.EnableRetries(0, 0)
				result, response, operationErr = dnsRecordsService.DeleteDnsRecord(deleteDnsRecordOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteDnsRecord(deleteDnsRecordOptions *DeleteDnsRecordOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		deleteDnsRecordPath := "/v1/testString/zones/testString/dns_records/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteDnsRecordPath))
					Expect(req.Method).To(Equal("DELETE"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": "f1aba936b94213e5b8dca0c0dbf1f9cc"}}`)
				}))
			})
			It(`Invoke DeleteDnsRecord successfully with retries`, func() {
				dnsRecordsService, serviceErr := dnsrecordsv1.NewDnsRecordsV1(&dnsrecordsv1.DnsRecordsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsRecordsService).ToNot(BeNil())
				dnsRecordsService.EnableRetries(0, 0)

				// Construct an instance of the DeleteDnsRecordOptions model
				deleteDnsRecordOptionsModel := new(dnsrecordsv1.DeleteDnsRecordOptions)
				deleteDnsRecordOptionsModel.DnsrecordIdentifier = core.StringPtr("testString")
				deleteDnsRecordOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := dnsRecordsService.DeleteDnsRecordWithContext(ctx, deleteDnsRecordOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				dnsRecordsService.DisableRetries()
				result, response, operationErr := dnsRecordsService.DeleteDnsRecord(deleteDnsRecordOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = dnsRecordsService.DeleteDnsRecordWithContext(ctx, deleteDnsRecordOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(deleteDnsRecordPath))
					Expect(req.Method).To(Equal("DELETE"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": "f1aba936b94213e5b8dca0c0dbf1f9cc"}}`)
				}))
			})
			It(`Invoke DeleteDnsRecord successfully`, func() {
				dnsRecordsService, serviceErr := dnsrecordsv1.NewDnsRecordsV1(&dnsrecordsv1.DnsRecordsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsRecordsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := dnsRecordsService.DeleteDnsRecord(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DeleteDnsRecordOptions model
				deleteDnsRecordOptionsModel := new(dnsrecordsv1.DeleteDnsRecordOptions)
				deleteDnsRecordOptionsModel.DnsrecordIdentifier = core.StringPtr("testString")
				deleteDnsRecordOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = dnsRecordsService.DeleteDnsRecord(deleteDnsRecordOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke DeleteDnsRecord with error: Operation validation and request error`, func() {
				dnsRecordsService, serviceErr := dnsrecordsv1.NewDnsRecordsV1(&dnsrecordsv1.DnsRecordsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsRecordsService).ToNot(BeNil())

				// Construct an instance of the DeleteDnsRecordOptions model
				deleteDnsRecordOptionsModel := new(dnsrecordsv1.DeleteDnsRecordOptions)
				deleteDnsRecordOptionsModel.DnsrecordIdentifier = core.StringPtr("testString")
				deleteDnsRecordOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := dnsRecordsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := dnsRecordsService.DeleteDnsRecord(deleteDnsRecordOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the DeleteDnsRecordOptions model with no property values
				deleteDnsRecordOptionsModelNew := new(dnsrecordsv1.DeleteDnsRecordOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = dnsRecordsService.DeleteDnsRecord(deleteDnsRecordOptionsModelNew)
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
			It(`Invoke DeleteDnsRecord successfully`, func() {
				dnsRecordsService, serviceErr := dnsrecordsv1.NewDnsRecordsV1(&dnsrecordsv1.DnsRecordsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsRecordsService).ToNot(BeNil())

				// Construct an instance of the DeleteDnsRecordOptions model
				deleteDnsRecordOptionsModel := new(dnsrecordsv1.DeleteDnsRecordOptions)
				deleteDnsRecordOptionsModel.DnsrecordIdentifier = core.StringPtr("testString")
				deleteDnsRecordOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := dnsRecordsService.DeleteDnsRecord(deleteDnsRecordOptionsModel)
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
	Describe(`GetDnsRecord(getDnsRecordOptions *GetDnsRecordOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getDnsRecordPath := "/v1/testString/zones/testString/dns_records/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getDnsRecordPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetDnsRecord with error: Operation response processing error`, func() {
				dnsRecordsService, serviceErr := dnsrecordsv1.NewDnsRecordsV1(&dnsrecordsv1.DnsRecordsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsRecordsService).ToNot(BeNil())

				// Construct an instance of the GetDnsRecordOptions model
				getDnsRecordOptionsModel := new(dnsrecordsv1.GetDnsRecordOptions)
				getDnsRecordOptionsModel.DnsrecordIdentifier = core.StringPtr("testString")
				getDnsRecordOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := dnsRecordsService.GetDnsRecord(getDnsRecordOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				dnsRecordsService.EnableRetries(0, 0)
				result, response, operationErr = dnsRecordsService.GetDnsRecord(getDnsRecordOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetDnsRecord(getDnsRecordOptions *GetDnsRecordOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getDnsRecordPath := "/v1/testString/zones/testString/dns_records/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getDnsRecordPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": "f1aba936b94213e5b8dca0c0dbf1f9cc", "created_on": "2014-01-01T05:20:00.12345Z", "modified_on": "2014-01-01T05:20:00.12345Z", "name": "host-1.test-example.com", "type": "A", "content": "169.154.10.10", "zone_id": "023e105f4ecef8ad9ca31a8372d0c353", "zone_name": "test-example.com", "proxiable": true, "proxied": false, "ttl": 120, "priority": 5, "data": {"anyKey": "anyValue"}}}`)
				}))
			})
			It(`Invoke GetDnsRecord successfully with retries`, func() {
				dnsRecordsService, serviceErr := dnsrecordsv1.NewDnsRecordsV1(&dnsrecordsv1.DnsRecordsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsRecordsService).ToNot(BeNil())
				dnsRecordsService.EnableRetries(0, 0)

				// Construct an instance of the GetDnsRecordOptions model
				getDnsRecordOptionsModel := new(dnsrecordsv1.GetDnsRecordOptions)
				getDnsRecordOptionsModel.DnsrecordIdentifier = core.StringPtr("testString")
				getDnsRecordOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := dnsRecordsService.GetDnsRecordWithContext(ctx, getDnsRecordOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				dnsRecordsService.DisableRetries()
				result, response, operationErr := dnsRecordsService.GetDnsRecord(getDnsRecordOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = dnsRecordsService.GetDnsRecordWithContext(ctx, getDnsRecordOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getDnsRecordPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": "f1aba936b94213e5b8dca0c0dbf1f9cc", "created_on": "2014-01-01T05:20:00.12345Z", "modified_on": "2014-01-01T05:20:00.12345Z", "name": "host-1.test-example.com", "type": "A", "content": "169.154.10.10", "zone_id": "023e105f4ecef8ad9ca31a8372d0c353", "zone_name": "test-example.com", "proxiable": true, "proxied": false, "ttl": 120, "priority": 5, "data": {"anyKey": "anyValue"}}}`)
				}))
			})
			It(`Invoke GetDnsRecord successfully`, func() {
				dnsRecordsService, serviceErr := dnsrecordsv1.NewDnsRecordsV1(&dnsrecordsv1.DnsRecordsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsRecordsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := dnsRecordsService.GetDnsRecord(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetDnsRecordOptions model
				getDnsRecordOptionsModel := new(dnsrecordsv1.GetDnsRecordOptions)
				getDnsRecordOptionsModel.DnsrecordIdentifier = core.StringPtr("testString")
				getDnsRecordOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = dnsRecordsService.GetDnsRecord(getDnsRecordOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetDnsRecord with error: Operation validation and request error`, func() {
				dnsRecordsService, serviceErr := dnsrecordsv1.NewDnsRecordsV1(&dnsrecordsv1.DnsRecordsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsRecordsService).ToNot(BeNil())

				// Construct an instance of the GetDnsRecordOptions model
				getDnsRecordOptionsModel := new(dnsrecordsv1.GetDnsRecordOptions)
				getDnsRecordOptionsModel.DnsrecordIdentifier = core.StringPtr("testString")
				getDnsRecordOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := dnsRecordsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := dnsRecordsService.GetDnsRecord(getDnsRecordOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetDnsRecordOptions model with no property values
				getDnsRecordOptionsModelNew := new(dnsrecordsv1.GetDnsRecordOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = dnsRecordsService.GetDnsRecord(getDnsRecordOptionsModelNew)
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
			It(`Invoke GetDnsRecord successfully`, func() {
				dnsRecordsService, serviceErr := dnsrecordsv1.NewDnsRecordsV1(&dnsrecordsv1.DnsRecordsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsRecordsService).ToNot(BeNil())

				// Construct an instance of the GetDnsRecordOptions model
				getDnsRecordOptionsModel := new(dnsrecordsv1.GetDnsRecordOptions)
				getDnsRecordOptionsModel.DnsrecordIdentifier = core.StringPtr("testString")
				getDnsRecordOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := dnsRecordsService.GetDnsRecord(getDnsRecordOptionsModel)
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
	Describe(`UpdateDnsRecord(updateDnsRecordOptions *UpdateDnsRecordOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		updateDnsRecordPath := "/v1/testString/zones/testString/dns_records/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateDnsRecordPath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateDnsRecord with error: Operation response processing error`, func() {
				dnsRecordsService, serviceErr := dnsrecordsv1.NewDnsRecordsV1(&dnsrecordsv1.DnsRecordsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsRecordsService).ToNot(BeNil())

				// Construct an instance of the UpdateDnsRecordOptions model
				updateDnsRecordOptionsModel := new(dnsrecordsv1.UpdateDnsRecordOptions)
				updateDnsRecordOptionsModel.DnsrecordIdentifier = core.StringPtr("testString")
				updateDnsRecordOptionsModel.Name = core.StringPtr("host-1.test-example.com")
				updateDnsRecordOptionsModel.Type = core.StringPtr("A")
				updateDnsRecordOptionsModel.TTL = core.Int64Ptr(int64(120))
				updateDnsRecordOptionsModel.Content = core.StringPtr("1.2.3.4")
				updateDnsRecordOptionsModel.Priority = core.Int64Ptr(int64(5))
				updateDnsRecordOptionsModel.Proxied = core.BoolPtr(false)
				updateDnsRecordOptionsModel.Data = map[string]interface{}{"anyKey": "anyValue"}
				updateDnsRecordOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := dnsRecordsService.UpdateDnsRecord(updateDnsRecordOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				dnsRecordsService.EnableRetries(0, 0)
				result, response, operationErr = dnsRecordsService.UpdateDnsRecord(updateDnsRecordOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateDnsRecord(updateDnsRecordOptions *UpdateDnsRecordOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		updateDnsRecordPath := "/v1/testString/zones/testString/dns_records/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateDnsRecordPath))
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
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": "f1aba936b94213e5b8dca0c0dbf1f9cc", "created_on": "2014-01-01T05:20:00.12345Z", "modified_on": "2014-01-01T05:20:00.12345Z", "name": "host-1.test-example.com", "type": "A", "content": "169.154.10.10", "zone_id": "023e105f4ecef8ad9ca31a8372d0c353", "zone_name": "test-example.com", "proxiable": true, "proxied": false, "ttl": 120, "priority": 5, "data": {"anyKey": "anyValue"}}}`)
				}))
			})
			It(`Invoke UpdateDnsRecord successfully with retries`, func() {
				dnsRecordsService, serviceErr := dnsrecordsv1.NewDnsRecordsV1(&dnsrecordsv1.DnsRecordsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsRecordsService).ToNot(BeNil())
				dnsRecordsService.EnableRetries(0, 0)

				// Construct an instance of the UpdateDnsRecordOptions model
				updateDnsRecordOptionsModel := new(dnsrecordsv1.UpdateDnsRecordOptions)
				updateDnsRecordOptionsModel.DnsrecordIdentifier = core.StringPtr("testString")
				updateDnsRecordOptionsModel.Name = core.StringPtr("host-1.test-example.com")
				updateDnsRecordOptionsModel.Type = core.StringPtr("A")
				updateDnsRecordOptionsModel.TTL = core.Int64Ptr(int64(120))
				updateDnsRecordOptionsModel.Content = core.StringPtr("1.2.3.4")
				updateDnsRecordOptionsModel.Priority = core.Int64Ptr(int64(5))
				updateDnsRecordOptionsModel.Proxied = core.BoolPtr(false)
				updateDnsRecordOptionsModel.Data = map[string]interface{}{"anyKey": "anyValue"}
				updateDnsRecordOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := dnsRecordsService.UpdateDnsRecordWithContext(ctx, updateDnsRecordOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				dnsRecordsService.DisableRetries()
				result, response, operationErr := dnsRecordsService.UpdateDnsRecord(updateDnsRecordOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = dnsRecordsService.UpdateDnsRecordWithContext(ctx, updateDnsRecordOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(updateDnsRecordPath))
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
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": "f1aba936b94213e5b8dca0c0dbf1f9cc", "created_on": "2014-01-01T05:20:00.12345Z", "modified_on": "2014-01-01T05:20:00.12345Z", "name": "host-1.test-example.com", "type": "A", "content": "169.154.10.10", "zone_id": "023e105f4ecef8ad9ca31a8372d0c353", "zone_name": "test-example.com", "proxiable": true, "proxied": false, "ttl": 120, "priority": 5, "data": {"anyKey": "anyValue"}}}`)
				}))
			})
			It(`Invoke UpdateDnsRecord successfully`, func() {
				dnsRecordsService, serviceErr := dnsrecordsv1.NewDnsRecordsV1(&dnsrecordsv1.DnsRecordsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsRecordsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := dnsRecordsService.UpdateDnsRecord(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateDnsRecordOptions model
				updateDnsRecordOptionsModel := new(dnsrecordsv1.UpdateDnsRecordOptions)
				updateDnsRecordOptionsModel.DnsrecordIdentifier = core.StringPtr("testString")
				updateDnsRecordOptionsModel.Name = core.StringPtr("host-1.test-example.com")
				updateDnsRecordOptionsModel.Type = core.StringPtr("A")
				updateDnsRecordOptionsModel.TTL = core.Int64Ptr(int64(120))
				updateDnsRecordOptionsModel.Content = core.StringPtr("1.2.3.4")
				updateDnsRecordOptionsModel.Priority = core.Int64Ptr(int64(5))
				updateDnsRecordOptionsModel.Proxied = core.BoolPtr(false)
				updateDnsRecordOptionsModel.Data = map[string]interface{}{"anyKey": "anyValue"}
				updateDnsRecordOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = dnsRecordsService.UpdateDnsRecord(updateDnsRecordOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateDnsRecord with error: Operation validation and request error`, func() {
				dnsRecordsService, serviceErr := dnsrecordsv1.NewDnsRecordsV1(&dnsrecordsv1.DnsRecordsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsRecordsService).ToNot(BeNil())

				// Construct an instance of the UpdateDnsRecordOptions model
				updateDnsRecordOptionsModel := new(dnsrecordsv1.UpdateDnsRecordOptions)
				updateDnsRecordOptionsModel.DnsrecordIdentifier = core.StringPtr("testString")
				updateDnsRecordOptionsModel.Name = core.StringPtr("host-1.test-example.com")
				updateDnsRecordOptionsModel.Type = core.StringPtr("A")
				updateDnsRecordOptionsModel.TTL = core.Int64Ptr(int64(120))
				updateDnsRecordOptionsModel.Content = core.StringPtr("1.2.3.4")
				updateDnsRecordOptionsModel.Priority = core.Int64Ptr(int64(5))
				updateDnsRecordOptionsModel.Proxied = core.BoolPtr(false)
				updateDnsRecordOptionsModel.Data = map[string]interface{}{"anyKey": "anyValue"}
				updateDnsRecordOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := dnsRecordsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := dnsRecordsService.UpdateDnsRecord(updateDnsRecordOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateDnsRecordOptions model with no property values
				updateDnsRecordOptionsModelNew := new(dnsrecordsv1.UpdateDnsRecordOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = dnsRecordsService.UpdateDnsRecord(updateDnsRecordOptionsModelNew)
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
			It(`Invoke UpdateDnsRecord successfully`, func() {
				dnsRecordsService, serviceErr := dnsrecordsv1.NewDnsRecordsV1(&dnsrecordsv1.DnsRecordsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsRecordsService).ToNot(BeNil())

				// Construct an instance of the UpdateDnsRecordOptions model
				updateDnsRecordOptionsModel := new(dnsrecordsv1.UpdateDnsRecordOptions)
				updateDnsRecordOptionsModel.DnsrecordIdentifier = core.StringPtr("testString")
				updateDnsRecordOptionsModel.Name = core.StringPtr("host-1.test-example.com")
				updateDnsRecordOptionsModel.Type = core.StringPtr("A")
				updateDnsRecordOptionsModel.TTL = core.Int64Ptr(int64(120))
				updateDnsRecordOptionsModel.Content = core.StringPtr("1.2.3.4")
				updateDnsRecordOptionsModel.Priority = core.Int64Ptr(int64(5))
				updateDnsRecordOptionsModel.Proxied = core.BoolPtr(false)
				updateDnsRecordOptionsModel.Data = map[string]interface{}{"anyKey": "anyValue"}
				updateDnsRecordOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := dnsRecordsService.UpdateDnsRecord(updateDnsRecordOptionsModel)
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
	Describe(`BatchDnsRecords(batchDnsRecordsOptions *BatchDnsRecordsOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		batchDnsRecordsPath := "/v1/testString/zones/testString/dns_records/batch"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(batchDnsRecordsPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke BatchDnsRecords with error: Operation response processing error`, func() {
				dnsRecordsService, serviceErr := dnsrecordsv1.NewDnsRecordsV1(&dnsrecordsv1.DnsRecordsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsRecordsService).ToNot(BeNil())

				// Construct an instance of the BatchDnsRecordsRequestDeletesItem model
				batchDnsRecordsRequestDeletesItemModel := new(dnsrecordsv1.BatchDnsRecordsRequestDeletesItem)
				batchDnsRecordsRequestDeletesItemModel.ID = core.StringPtr("testString")

				// Construct an instance of the BatchDnsRecordsRequestPatchesItem model
				batchDnsRecordsRequestPatchesItemModel := new(dnsrecordsv1.BatchDnsRecordsRequestPatchesItem)
				batchDnsRecordsRequestPatchesItemModel.ID = core.StringPtr("testString")
				batchDnsRecordsRequestPatchesItemModel.Name = core.StringPtr("host-1.test-example.com")
				batchDnsRecordsRequestPatchesItemModel.Type = core.StringPtr("A")
				batchDnsRecordsRequestPatchesItemModel.TTL = core.Int64Ptr(int64(120))
				batchDnsRecordsRequestPatchesItemModel.Content = core.StringPtr("1.2.3.4")
				batchDnsRecordsRequestPatchesItemModel.Priority = core.Int64Ptr(int64(5))
				batchDnsRecordsRequestPatchesItemModel.Proxied = core.BoolPtr(false)
				batchDnsRecordsRequestPatchesItemModel.Data = map[string]interface{}{"anyKey": "anyValue"}

				// Construct an instance of the DnsrecordInput model
				dnsrecordInputModel := new(dnsrecordsv1.DnsrecordInput)
				dnsrecordInputModel.Name = core.StringPtr("host-1.test-example.com")
				dnsrecordInputModel.Type = core.StringPtr("A")
				dnsrecordInputModel.TTL = core.Int64Ptr(int64(120))
				dnsrecordInputModel.Content = core.StringPtr("1.2.3.4")
				dnsrecordInputModel.Priority = core.Int64Ptr(int64(5))
				dnsrecordInputModel.Proxied = core.BoolPtr(false)
				dnsrecordInputModel.Data = map[string]interface{}{"anyKey": "anyValue"}

				// Construct an instance of the BatchDnsRecordsRequestPutsItem model
				batchDnsRecordsRequestPutsItemModel := new(dnsrecordsv1.BatchDnsRecordsRequestPutsItem)
				batchDnsRecordsRequestPutsItemModel.ID = core.StringPtr("testString")
				batchDnsRecordsRequestPutsItemModel.Name = core.StringPtr("host-1.test-example.com")
				batchDnsRecordsRequestPutsItemModel.Type = core.StringPtr("A")
				batchDnsRecordsRequestPutsItemModel.TTL = core.Int64Ptr(int64(120))
				batchDnsRecordsRequestPutsItemModel.Content = core.StringPtr("1.2.3.4")
				batchDnsRecordsRequestPutsItemModel.Priority = core.Int64Ptr(int64(5))
				batchDnsRecordsRequestPutsItemModel.Proxied = core.BoolPtr(false)
				batchDnsRecordsRequestPutsItemModel.Data = map[string]interface{}{"anyKey": "anyValue"}

				// Construct an instance of the BatchDnsRecordsOptions model
				batchDnsRecordsOptionsModel := new(dnsrecordsv1.BatchDnsRecordsOptions)
				batchDnsRecordsOptionsModel.Deletes = []dnsrecordsv1.BatchDnsRecordsRequestDeletesItem{*batchDnsRecordsRequestDeletesItemModel}
				batchDnsRecordsOptionsModel.Patches = []dnsrecordsv1.BatchDnsRecordsRequestPatchesItem{*batchDnsRecordsRequestPatchesItemModel}
				batchDnsRecordsOptionsModel.Posts = []dnsrecordsv1.DnsrecordInput{*dnsrecordInputModel}
				batchDnsRecordsOptionsModel.Puts = []dnsrecordsv1.BatchDnsRecordsRequestPutsItem{*batchDnsRecordsRequestPutsItemModel}
				batchDnsRecordsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := dnsRecordsService.BatchDnsRecords(batchDnsRecordsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				dnsRecordsService.EnableRetries(0, 0)
				result, response, operationErr = dnsRecordsService.BatchDnsRecords(batchDnsRecordsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`BatchDnsRecords(batchDnsRecordsOptions *BatchDnsRecordsOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		batchDnsRecordsPath := "/v1/testString/zones/testString/dns_records/batch"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(batchDnsRecordsPath))
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
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"deletes": [{"id": "f1aba936b94213e5b8dca0c0dbf1f9cc", "created_on": "2014-01-01T05:20:00.12345Z", "modified_on": "2014-01-01T05:20:00.12345Z", "name": "host-1.test-example.com", "type": "A", "content": "1.2.3.4", "proxiable": false, "proxied": false, "ttl": 120, "priority": 5, "data": {"anyKey": "anyValue"}, "settings": {"anyKey": "anyValue"}, "meta": {"anyKey": "anyValue"}, "comment": "Comment", "tags": ["Tags"]}], "patches": [{"id": "f1aba936b94213e5b8dca0c0dbf1f9cc", "created_on": "2014-01-01T05:20:00.12345Z", "modified_on": "2014-01-01T05:20:00.12345Z", "name": "host-1.test-example.com", "type": "A", "content": "1.2.3.4", "proxiable": false, "proxied": false, "ttl": 120, "priority": 5, "data": {"anyKey": "anyValue"}, "settings": {"anyKey": "anyValue"}, "meta": {"anyKey": "anyValue"}, "comment": "Comment", "tags": ["Tags"]}], "posts": [{"id": "f1aba936b94213e5b8dca0c0dbf1f9cc", "created_on": "2014-01-01T05:20:00.12345Z", "modified_on": "2014-01-01T05:20:00.12345Z", "name": "host-1.test-example.com", "type": "A", "content": "1.2.3.4", "proxiable": false, "proxied": false, "ttl": 120, "priority": 5, "data": {"anyKey": "anyValue"}, "settings": {"anyKey": "anyValue"}, "meta": {"anyKey": "anyValue"}, "comment": "Comment", "tags": ["Tags"]}], "puts": [{"id": "f1aba936b94213e5b8dca0c0dbf1f9cc", "created_on": "2014-01-01T05:20:00.12345Z", "modified_on": "2014-01-01T05:20:00.12345Z", "name": "host-1.test-example.com", "type": "A", "content": "1.2.3.4", "proxiable": false, "proxied": false, "ttl": 120, "priority": 5, "data": {"anyKey": "anyValue"}, "settings": {"anyKey": "anyValue"}, "meta": {"anyKey": "anyValue"}, "comment": "Comment", "tags": ["Tags"]}]}}`)
				}))
			})
			It(`Invoke BatchDnsRecords successfully with retries`, func() {
				dnsRecordsService, serviceErr := dnsrecordsv1.NewDnsRecordsV1(&dnsrecordsv1.DnsRecordsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsRecordsService).ToNot(BeNil())
				dnsRecordsService.EnableRetries(0, 0)

				// Construct an instance of the BatchDnsRecordsRequestDeletesItem model
				batchDnsRecordsRequestDeletesItemModel := new(dnsrecordsv1.BatchDnsRecordsRequestDeletesItem)
				batchDnsRecordsRequestDeletesItemModel.ID = core.StringPtr("testString")

				// Construct an instance of the BatchDnsRecordsRequestPatchesItem model
				batchDnsRecordsRequestPatchesItemModel := new(dnsrecordsv1.BatchDnsRecordsRequestPatchesItem)
				batchDnsRecordsRequestPatchesItemModel.ID = core.StringPtr("testString")
				batchDnsRecordsRequestPatchesItemModel.Name = core.StringPtr("host-1.test-example.com")
				batchDnsRecordsRequestPatchesItemModel.Type = core.StringPtr("A")
				batchDnsRecordsRequestPatchesItemModel.TTL = core.Int64Ptr(int64(120))
				batchDnsRecordsRequestPatchesItemModel.Content = core.StringPtr("1.2.3.4")
				batchDnsRecordsRequestPatchesItemModel.Priority = core.Int64Ptr(int64(5))
				batchDnsRecordsRequestPatchesItemModel.Proxied = core.BoolPtr(false)
				batchDnsRecordsRequestPatchesItemModel.Data = map[string]interface{}{"anyKey": "anyValue"}

				// Construct an instance of the DnsrecordInput model
				dnsrecordInputModel := new(dnsrecordsv1.DnsrecordInput)
				dnsrecordInputModel.Name = core.StringPtr("host-1.test-example.com")
				dnsrecordInputModel.Type = core.StringPtr("A")
				dnsrecordInputModel.TTL = core.Int64Ptr(int64(120))
				dnsrecordInputModel.Content = core.StringPtr("1.2.3.4")
				dnsrecordInputModel.Priority = core.Int64Ptr(int64(5))
				dnsrecordInputModel.Proxied = core.BoolPtr(false)
				dnsrecordInputModel.Data = map[string]interface{}{"anyKey": "anyValue"}

				// Construct an instance of the BatchDnsRecordsRequestPutsItem model
				batchDnsRecordsRequestPutsItemModel := new(dnsrecordsv1.BatchDnsRecordsRequestPutsItem)
				batchDnsRecordsRequestPutsItemModel.ID = core.StringPtr("testString")
				batchDnsRecordsRequestPutsItemModel.Name = core.StringPtr("host-1.test-example.com")
				batchDnsRecordsRequestPutsItemModel.Type = core.StringPtr("A")
				batchDnsRecordsRequestPutsItemModel.TTL = core.Int64Ptr(int64(120))
				batchDnsRecordsRequestPutsItemModel.Content = core.StringPtr("1.2.3.4")
				batchDnsRecordsRequestPutsItemModel.Priority = core.Int64Ptr(int64(5))
				batchDnsRecordsRequestPutsItemModel.Proxied = core.BoolPtr(false)
				batchDnsRecordsRequestPutsItemModel.Data = map[string]interface{}{"anyKey": "anyValue"}

				// Construct an instance of the BatchDnsRecordsOptions model
				batchDnsRecordsOptionsModel := new(dnsrecordsv1.BatchDnsRecordsOptions)
				batchDnsRecordsOptionsModel.Deletes = []dnsrecordsv1.BatchDnsRecordsRequestDeletesItem{*batchDnsRecordsRequestDeletesItemModel}
				batchDnsRecordsOptionsModel.Patches = []dnsrecordsv1.BatchDnsRecordsRequestPatchesItem{*batchDnsRecordsRequestPatchesItemModel}
				batchDnsRecordsOptionsModel.Posts = []dnsrecordsv1.DnsrecordInput{*dnsrecordInputModel}
				batchDnsRecordsOptionsModel.Puts = []dnsrecordsv1.BatchDnsRecordsRequestPutsItem{*batchDnsRecordsRequestPutsItemModel}
				batchDnsRecordsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := dnsRecordsService.BatchDnsRecordsWithContext(ctx, batchDnsRecordsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				dnsRecordsService.DisableRetries()
				result, response, operationErr := dnsRecordsService.BatchDnsRecords(batchDnsRecordsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = dnsRecordsService.BatchDnsRecordsWithContext(ctx, batchDnsRecordsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(batchDnsRecordsPath))
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
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"deletes": [{"id": "f1aba936b94213e5b8dca0c0dbf1f9cc", "created_on": "2014-01-01T05:20:00.12345Z", "modified_on": "2014-01-01T05:20:00.12345Z", "name": "host-1.test-example.com", "type": "A", "content": "1.2.3.4", "proxiable": false, "proxied": false, "ttl": 120, "priority": 5, "data": {"anyKey": "anyValue"}, "settings": {"anyKey": "anyValue"}, "meta": {"anyKey": "anyValue"}, "comment": "Comment", "tags": ["Tags"]}], "patches": [{"id": "f1aba936b94213e5b8dca0c0dbf1f9cc", "created_on": "2014-01-01T05:20:00.12345Z", "modified_on": "2014-01-01T05:20:00.12345Z", "name": "host-1.test-example.com", "type": "A", "content": "1.2.3.4", "proxiable": false, "proxied": false, "ttl": 120, "priority": 5, "data": {"anyKey": "anyValue"}, "settings": {"anyKey": "anyValue"}, "meta": {"anyKey": "anyValue"}, "comment": "Comment", "tags": ["Tags"]}], "posts": [{"id": "f1aba936b94213e5b8dca0c0dbf1f9cc", "created_on": "2014-01-01T05:20:00.12345Z", "modified_on": "2014-01-01T05:20:00.12345Z", "name": "host-1.test-example.com", "type": "A", "content": "1.2.3.4", "proxiable": false, "proxied": false, "ttl": 120, "priority": 5, "data": {"anyKey": "anyValue"}, "settings": {"anyKey": "anyValue"}, "meta": {"anyKey": "anyValue"}, "comment": "Comment", "tags": ["Tags"]}], "puts": [{"id": "f1aba936b94213e5b8dca0c0dbf1f9cc", "created_on": "2014-01-01T05:20:00.12345Z", "modified_on": "2014-01-01T05:20:00.12345Z", "name": "host-1.test-example.com", "type": "A", "content": "1.2.3.4", "proxiable": false, "proxied": false, "ttl": 120, "priority": 5, "data": {"anyKey": "anyValue"}, "settings": {"anyKey": "anyValue"}, "meta": {"anyKey": "anyValue"}, "comment": "Comment", "tags": ["Tags"]}]}}`)
				}))
			})
			It(`Invoke BatchDnsRecords successfully`, func() {
				dnsRecordsService, serviceErr := dnsrecordsv1.NewDnsRecordsV1(&dnsrecordsv1.DnsRecordsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsRecordsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := dnsRecordsService.BatchDnsRecords(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the BatchDnsRecordsRequestDeletesItem model
				batchDnsRecordsRequestDeletesItemModel := new(dnsrecordsv1.BatchDnsRecordsRequestDeletesItem)
				batchDnsRecordsRequestDeletesItemModel.ID = core.StringPtr("testString")

				// Construct an instance of the BatchDnsRecordsRequestPatchesItem model
				batchDnsRecordsRequestPatchesItemModel := new(dnsrecordsv1.BatchDnsRecordsRequestPatchesItem)
				batchDnsRecordsRequestPatchesItemModel.ID = core.StringPtr("testString")
				batchDnsRecordsRequestPatchesItemModel.Name = core.StringPtr("host-1.test-example.com")
				batchDnsRecordsRequestPatchesItemModel.Type = core.StringPtr("A")
				batchDnsRecordsRequestPatchesItemModel.TTL = core.Int64Ptr(int64(120))
				batchDnsRecordsRequestPatchesItemModel.Content = core.StringPtr("1.2.3.4")
				batchDnsRecordsRequestPatchesItemModel.Priority = core.Int64Ptr(int64(5))
				batchDnsRecordsRequestPatchesItemModel.Proxied = core.BoolPtr(false)
				batchDnsRecordsRequestPatchesItemModel.Data = map[string]interface{}{"anyKey": "anyValue"}

				// Construct an instance of the DnsrecordInput model
				dnsrecordInputModel := new(dnsrecordsv1.DnsrecordInput)
				dnsrecordInputModel.Name = core.StringPtr("host-1.test-example.com")
				dnsrecordInputModel.Type = core.StringPtr("A")
				dnsrecordInputModel.TTL = core.Int64Ptr(int64(120))
				dnsrecordInputModel.Content = core.StringPtr("1.2.3.4")
				dnsrecordInputModel.Priority = core.Int64Ptr(int64(5))
				dnsrecordInputModel.Proxied = core.BoolPtr(false)
				dnsrecordInputModel.Data = map[string]interface{}{"anyKey": "anyValue"}

				// Construct an instance of the BatchDnsRecordsRequestPutsItem model
				batchDnsRecordsRequestPutsItemModel := new(dnsrecordsv1.BatchDnsRecordsRequestPutsItem)
				batchDnsRecordsRequestPutsItemModel.ID = core.StringPtr("testString")
				batchDnsRecordsRequestPutsItemModel.Name = core.StringPtr("host-1.test-example.com")
				batchDnsRecordsRequestPutsItemModel.Type = core.StringPtr("A")
				batchDnsRecordsRequestPutsItemModel.TTL = core.Int64Ptr(int64(120))
				batchDnsRecordsRequestPutsItemModel.Content = core.StringPtr("1.2.3.4")
				batchDnsRecordsRequestPutsItemModel.Priority = core.Int64Ptr(int64(5))
				batchDnsRecordsRequestPutsItemModel.Proxied = core.BoolPtr(false)
				batchDnsRecordsRequestPutsItemModel.Data = map[string]interface{}{"anyKey": "anyValue"}

				// Construct an instance of the BatchDnsRecordsOptions model
				batchDnsRecordsOptionsModel := new(dnsrecordsv1.BatchDnsRecordsOptions)
				batchDnsRecordsOptionsModel.Deletes = []dnsrecordsv1.BatchDnsRecordsRequestDeletesItem{*batchDnsRecordsRequestDeletesItemModel}
				batchDnsRecordsOptionsModel.Patches = []dnsrecordsv1.BatchDnsRecordsRequestPatchesItem{*batchDnsRecordsRequestPatchesItemModel}
				batchDnsRecordsOptionsModel.Posts = []dnsrecordsv1.DnsrecordInput{*dnsrecordInputModel}
				batchDnsRecordsOptionsModel.Puts = []dnsrecordsv1.BatchDnsRecordsRequestPutsItem{*batchDnsRecordsRequestPutsItemModel}
				batchDnsRecordsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = dnsRecordsService.BatchDnsRecords(batchDnsRecordsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke BatchDnsRecords with error: Operation request error`, func() {
				dnsRecordsService, serviceErr := dnsrecordsv1.NewDnsRecordsV1(&dnsrecordsv1.DnsRecordsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsRecordsService).ToNot(BeNil())

				// Construct an instance of the BatchDnsRecordsRequestDeletesItem model
				batchDnsRecordsRequestDeletesItemModel := new(dnsrecordsv1.BatchDnsRecordsRequestDeletesItem)
				batchDnsRecordsRequestDeletesItemModel.ID = core.StringPtr("testString")

				// Construct an instance of the BatchDnsRecordsRequestPatchesItem model
				batchDnsRecordsRequestPatchesItemModel := new(dnsrecordsv1.BatchDnsRecordsRequestPatchesItem)
				batchDnsRecordsRequestPatchesItemModel.ID = core.StringPtr("testString")
				batchDnsRecordsRequestPatchesItemModel.Name = core.StringPtr("host-1.test-example.com")
				batchDnsRecordsRequestPatchesItemModel.Type = core.StringPtr("A")
				batchDnsRecordsRequestPatchesItemModel.TTL = core.Int64Ptr(int64(120))
				batchDnsRecordsRequestPatchesItemModel.Content = core.StringPtr("1.2.3.4")
				batchDnsRecordsRequestPatchesItemModel.Priority = core.Int64Ptr(int64(5))
				batchDnsRecordsRequestPatchesItemModel.Proxied = core.BoolPtr(false)
				batchDnsRecordsRequestPatchesItemModel.Data = map[string]interface{}{"anyKey": "anyValue"}

				// Construct an instance of the DnsrecordInput model
				dnsrecordInputModel := new(dnsrecordsv1.DnsrecordInput)
				dnsrecordInputModel.Name = core.StringPtr("host-1.test-example.com")
				dnsrecordInputModel.Type = core.StringPtr("A")
				dnsrecordInputModel.TTL = core.Int64Ptr(int64(120))
				dnsrecordInputModel.Content = core.StringPtr("1.2.3.4")
				dnsrecordInputModel.Priority = core.Int64Ptr(int64(5))
				dnsrecordInputModel.Proxied = core.BoolPtr(false)
				dnsrecordInputModel.Data = map[string]interface{}{"anyKey": "anyValue"}

				// Construct an instance of the BatchDnsRecordsRequestPutsItem model
				batchDnsRecordsRequestPutsItemModel := new(dnsrecordsv1.BatchDnsRecordsRequestPutsItem)
				batchDnsRecordsRequestPutsItemModel.ID = core.StringPtr("testString")
				batchDnsRecordsRequestPutsItemModel.Name = core.StringPtr("host-1.test-example.com")
				batchDnsRecordsRequestPutsItemModel.Type = core.StringPtr("A")
				batchDnsRecordsRequestPutsItemModel.TTL = core.Int64Ptr(int64(120))
				batchDnsRecordsRequestPutsItemModel.Content = core.StringPtr("1.2.3.4")
				batchDnsRecordsRequestPutsItemModel.Priority = core.Int64Ptr(int64(5))
				batchDnsRecordsRequestPutsItemModel.Proxied = core.BoolPtr(false)
				batchDnsRecordsRequestPutsItemModel.Data = map[string]interface{}{"anyKey": "anyValue"}

				// Construct an instance of the BatchDnsRecordsOptions model
				batchDnsRecordsOptionsModel := new(dnsrecordsv1.BatchDnsRecordsOptions)
				batchDnsRecordsOptionsModel.Deletes = []dnsrecordsv1.BatchDnsRecordsRequestDeletesItem{*batchDnsRecordsRequestDeletesItemModel}
				batchDnsRecordsOptionsModel.Patches = []dnsrecordsv1.BatchDnsRecordsRequestPatchesItem{*batchDnsRecordsRequestPatchesItemModel}
				batchDnsRecordsOptionsModel.Posts = []dnsrecordsv1.DnsrecordInput{*dnsrecordInputModel}
				batchDnsRecordsOptionsModel.Puts = []dnsrecordsv1.BatchDnsRecordsRequestPutsItem{*batchDnsRecordsRequestPutsItemModel}
				batchDnsRecordsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := dnsRecordsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := dnsRecordsService.BatchDnsRecords(batchDnsRecordsOptionsModel)
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
			It(`Invoke BatchDnsRecords successfully`, func() {
				dnsRecordsService, serviceErr := dnsrecordsv1.NewDnsRecordsV1(&dnsrecordsv1.DnsRecordsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(dnsRecordsService).ToNot(BeNil())

				// Construct an instance of the BatchDnsRecordsRequestDeletesItem model
				batchDnsRecordsRequestDeletesItemModel := new(dnsrecordsv1.BatchDnsRecordsRequestDeletesItem)
				batchDnsRecordsRequestDeletesItemModel.ID = core.StringPtr("testString")

				// Construct an instance of the BatchDnsRecordsRequestPatchesItem model
				batchDnsRecordsRequestPatchesItemModel := new(dnsrecordsv1.BatchDnsRecordsRequestPatchesItem)
				batchDnsRecordsRequestPatchesItemModel.ID = core.StringPtr("testString")
				batchDnsRecordsRequestPatchesItemModel.Name = core.StringPtr("host-1.test-example.com")
				batchDnsRecordsRequestPatchesItemModel.Type = core.StringPtr("A")
				batchDnsRecordsRequestPatchesItemModel.TTL = core.Int64Ptr(int64(120))
				batchDnsRecordsRequestPatchesItemModel.Content = core.StringPtr("1.2.3.4")
				batchDnsRecordsRequestPatchesItemModel.Priority = core.Int64Ptr(int64(5))
				batchDnsRecordsRequestPatchesItemModel.Proxied = core.BoolPtr(false)
				batchDnsRecordsRequestPatchesItemModel.Data = map[string]interface{}{"anyKey": "anyValue"}

				// Construct an instance of the DnsrecordInput model
				dnsrecordInputModel := new(dnsrecordsv1.DnsrecordInput)
				dnsrecordInputModel.Name = core.StringPtr("host-1.test-example.com")
				dnsrecordInputModel.Type = core.StringPtr("A")
				dnsrecordInputModel.TTL = core.Int64Ptr(int64(120))
				dnsrecordInputModel.Content = core.StringPtr("1.2.3.4")
				dnsrecordInputModel.Priority = core.Int64Ptr(int64(5))
				dnsrecordInputModel.Proxied = core.BoolPtr(false)
				dnsrecordInputModel.Data = map[string]interface{}{"anyKey": "anyValue"}

				// Construct an instance of the BatchDnsRecordsRequestPutsItem model
				batchDnsRecordsRequestPutsItemModel := new(dnsrecordsv1.BatchDnsRecordsRequestPutsItem)
				batchDnsRecordsRequestPutsItemModel.ID = core.StringPtr("testString")
				batchDnsRecordsRequestPutsItemModel.Name = core.StringPtr("host-1.test-example.com")
				batchDnsRecordsRequestPutsItemModel.Type = core.StringPtr("A")
				batchDnsRecordsRequestPutsItemModel.TTL = core.Int64Ptr(int64(120))
				batchDnsRecordsRequestPutsItemModel.Content = core.StringPtr("1.2.3.4")
				batchDnsRecordsRequestPutsItemModel.Priority = core.Int64Ptr(int64(5))
				batchDnsRecordsRequestPutsItemModel.Proxied = core.BoolPtr(false)
				batchDnsRecordsRequestPutsItemModel.Data = map[string]interface{}{"anyKey": "anyValue"}

				// Construct an instance of the BatchDnsRecordsOptions model
				batchDnsRecordsOptionsModel := new(dnsrecordsv1.BatchDnsRecordsOptions)
				batchDnsRecordsOptionsModel.Deletes = []dnsrecordsv1.BatchDnsRecordsRequestDeletesItem{*batchDnsRecordsRequestDeletesItemModel}
				batchDnsRecordsOptionsModel.Patches = []dnsrecordsv1.BatchDnsRecordsRequestPatchesItem{*batchDnsRecordsRequestPatchesItemModel}
				batchDnsRecordsOptionsModel.Posts = []dnsrecordsv1.DnsrecordInput{*dnsrecordInputModel}
				batchDnsRecordsOptionsModel.Puts = []dnsrecordsv1.BatchDnsRecordsRequestPutsItem{*batchDnsRecordsRequestPutsItemModel}
				batchDnsRecordsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := dnsRecordsService.BatchDnsRecords(batchDnsRecordsOptionsModel)
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
			dnsRecordsService, _ := dnsrecordsv1.NewDnsRecordsV1(&dnsrecordsv1.DnsRecordsV1Options{
				URL:            "http://dnsrecordsv1modelgenerator.com",
				Authenticator:  &core.NoAuthAuthenticator{},
				Crn:            core.StringPtr(crn),
				ZoneIdentifier: core.StringPtr(zoneIdentifier),
			})
			It(`Invoke NewBatchDnsRecordsOptions successfully`, func() {
				// Construct an instance of the BatchDnsRecordsRequestDeletesItem model
				batchDnsRecordsRequestDeletesItemModel := new(dnsrecordsv1.BatchDnsRecordsRequestDeletesItem)
				Expect(batchDnsRecordsRequestDeletesItemModel).ToNot(BeNil())
				batchDnsRecordsRequestDeletesItemModel.ID = core.StringPtr("testString")
				Expect(batchDnsRecordsRequestDeletesItemModel.ID).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the BatchDnsRecordsRequestPatchesItem model
				batchDnsRecordsRequestPatchesItemModel := new(dnsrecordsv1.BatchDnsRecordsRequestPatchesItem)
				Expect(batchDnsRecordsRequestPatchesItemModel).ToNot(BeNil())
				batchDnsRecordsRequestPatchesItemModel.ID = core.StringPtr("testString")
				batchDnsRecordsRequestPatchesItemModel.Name = core.StringPtr("host-1.test-example.com")
				batchDnsRecordsRequestPatchesItemModel.Type = core.StringPtr("A")
				batchDnsRecordsRequestPatchesItemModel.TTL = core.Int64Ptr(int64(120))
				batchDnsRecordsRequestPatchesItemModel.Content = core.StringPtr("1.2.3.4")
				batchDnsRecordsRequestPatchesItemModel.Priority = core.Int64Ptr(int64(5))
				batchDnsRecordsRequestPatchesItemModel.Proxied = core.BoolPtr(false)
				batchDnsRecordsRequestPatchesItemModel.Data = map[string]interface{}{"anyKey": "anyValue"}
				Expect(batchDnsRecordsRequestPatchesItemModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(batchDnsRecordsRequestPatchesItemModel.Name).To(Equal(core.StringPtr("host-1.test-example.com")))
				Expect(batchDnsRecordsRequestPatchesItemModel.Type).To(Equal(core.StringPtr("A")))
				Expect(batchDnsRecordsRequestPatchesItemModel.TTL).To(Equal(core.Int64Ptr(int64(120))))
				Expect(batchDnsRecordsRequestPatchesItemModel.Content).To(Equal(core.StringPtr("1.2.3.4")))
				Expect(batchDnsRecordsRequestPatchesItemModel.Priority).To(Equal(core.Int64Ptr(int64(5))))
				Expect(batchDnsRecordsRequestPatchesItemModel.Proxied).To(Equal(core.BoolPtr(false)))
				Expect(batchDnsRecordsRequestPatchesItemModel.Data).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))

				// Construct an instance of the DnsrecordInput model
				dnsrecordInputModel := new(dnsrecordsv1.DnsrecordInput)
				Expect(dnsrecordInputModel).ToNot(BeNil())
				dnsrecordInputModel.Name = core.StringPtr("host-1.test-example.com")
				dnsrecordInputModel.Type = core.StringPtr("A")
				dnsrecordInputModel.TTL = core.Int64Ptr(int64(120))
				dnsrecordInputModel.Content = core.StringPtr("1.2.3.4")
				dnsrecordInputModel.Priority = core.Int64Ptr(int64(5))
				dnsrecordInputModel.Proxied = core.BoolPtr(false)
				dnsrecordInputModel.Data = map[string]interface{}{"anyKey": "anyValue"}
				Expect(dnsrecordInputModel.Name).To(Equal(core.StringPtr("host-1.test-example.com")))
				Expect(dnsrecordInputModel.Type).To(Equal(core.StringPtr("A")))
				Expect(dnsrecordInputModel.TTL).To(Equal(core.Int64Ptr(int64(120))))
				Expect(dnsrecordInputModel.Content).To(Equal(core.StringPtr("1.2.3.4")))
				Expect(dnsrecordInputModel.Priority).To(Equal(core.Int64Ptr(int64(5))))
				Expect(dnsrecordInputModel.Proxied).To(Equal(core.BoolPtr(false)))
				Expect(dnsrecordInputModel.Data).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))

				// Construct an instance of the BatchDnsRecordsRequestPutsItem model
				batchDnsRecordsRequestPutsItemModel := new(dnsrecordsv1.BatchDnsRecordsRequestPutsItem)
				Expect(batchDnsRecordsRequestPutsItemModel).ToNot(BeNil())
				batchDnsRecordsRequestPutsItemModel.ID = core.StringPtr("testString")
				batchDnsRecordsRequestPutsItemModel.Name = core.StringPtr("host-1.test-example.com")
				batchDnsRecordsRequestPutsItemModel.Type = core.StringPtr("A")
				batchDnsRecordsRequestPutsItemModel.TTL = core.Int64Ptr(int64(120))
				batchDnsRecordsRequestPutsItemModel.Content = core.StringPtr("1.2.3.4")
				batchDnsRecordsRequestPutsItemModel.Priority = core.Int64Ptr(int64(5))
				batchDnsRecordsRequestPutsItemModel.Proxied = core.BoolPtr(false)
				batchDnsRecordsRequestPutsItemModel.Data = map[string]interface{}{"anyKey": "anyValue"}
				Expect(batchDnsRecordsRequestPutsItemModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(batchDnsRecordsRequestPutsItemModel.Name).To(Equal(core.StringPtr("host-1.test-example.com")))
				Expect(batchDnsRecordsRequestPutsItemModel.Type).To(Equal(core.StringPtr("A")))
				Expect(batchDnsRecordsRequestPutsItemModel.TTL).To(Equal(core.Int64Ptr(int64(120))))
				Expect(batchDnsRecordsRequestPutsItemModel.Content).To(Equal(core.StringPtr("1.2.3.4")))
				Expect(batchDnsRecordsRequestPutsItemModel.Priority).To(Equal(core.Int64Ptr(int64(5))))
				Expect(batchDnsRecordsRequestPutsItemModel.Proxied).To(Equal(core.BoolPtr(false)))
				Expect(batchDnsRecordsRequestPutsItemModel.Data).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))

				// Construct an instance of the BatchDnsRecordsOptions model
				batchDnsRecordsOptionsModel := dnsRecordsService.NewBatchDnsRecordsOptions()
				batchDnsRecordsOptionsModel.SetDeletes([]dnsrecordsv1.BatchDnsRecordsRequestDeletesItem{*batchDnsRecordsRequestDeletesItemModel})
				batchDnsRecordsOptionsModel.SetPatches([]dnsrecordsv1.BatchDnsRecordsRequestPatchesItem{*batchDnsRecordsRequestPatchesItemModel})
				batchDnsRecordsOptionsModel.SetPosts([]dnsrecordsv1.DnsrecordInput{*dnsrecordInputModel})
				batchDnsRecordsOptionsModel.SetPuts([]dnsrecordsv1.BatchDnsRecordsRequestPutsItem{*batchDnsRecordsRequestPutsItemModel})
				batchDnsRecordsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(batchDnsRecordsOptionsModel).ToNot(BeNil())
				Expect(batchDnsRecordsOptionsModel.Deletes).To(Equal([]dnsrecordsv1.BatchDnsRecordsRequestDeletesItem{*batchDnsRecordsRequestDeletesItemModel}))
				Expect(batchDnsRecordsOptionsModel.Patches).To(Equal([]dnsrecordsv1.BatchDnsRecordsRequestPatchesItem{*batchDnsRecordsRequestPatchesItemModel}))
				Expect(batchDnsRecordsOptionsModel.Posts).To(Equal([]dnsrecordsv1.DnsrecordInput{*dnsrecordInputModel}))
				Expect(batchDnsRecordsOptionsModel.Puts).To(Equal([]dnsrecordsv1.BatchDnsRecordsRequestPutsItem{*batchDnsRecordsRequestPutsItemModel}))
				Expect(batchDnsRecordsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewBatchDnsRecordsRequestDeletesItem successfully`, func() {
				id := "testString"
				_model, err := dnsRecordsService.NewBatchDnsRecordsRequestDeletesItem(id)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewBatchDnsRecordsRequestPatchesItem successfully`, func() {
				id := "testString"
				_model, err := dnsRecordsService.NewBatchDnsRecordsRequestPatchesItem(id)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewBatchDnsRecordsRequestPutsItem successfully`, func() {
				id := "testString"
				name := "host-1.test-example.com"
				typeVar := "A"
				ttl := int64(120)
				content := "1.2.3.4"
				_model, err := dnsRecordsService.NewBatchDnsRecordsRequestPutsItem(id, name, typeVar, ttl, content)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewCreateDnsRecordOptions successfully`, func() {
				// Construct an instance of the CreateDnsRecordOptions model
				createDnsRecordOptionsModel := dnsRecordsService.NewCreateDnsRecordOptions()
				createDnsRecordOptionsModel.SetName("host-1.test-example.com")
				createDnsRecordOptionsModel.SetType("A")
				createDnsRecordOptionsModel.SetTTL(int64(120))
				createDnsRecordOptionsModel.SetContent("1.2.3.4")
				createDnsRecordOptionsModel.SetPriority(int64(5))
				createDnsRecordOptionsModel.SetProxied(false)
				createDnsRecordOptionsModel.SetData(map[string]interface{}{"anyKey": "anyValue"})
				createDnsRecordOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createDnsRecordOptionsModel).ToNot(BeNil())
				Expect(createDnsRecordOptionsModel.Name).To(Equal(core.StringPtr("host-1.test-example.com")))
				Expect(createDnsRecordOptionsModel.Type).To(Equal(core.StringPtr("A")))
				Expect(createDnsRecordOptionsModel.TTL).To(Equal(core.Int64Ptr(int64(120))))
				Expect(createDnsRecordOptionsModel.Content).To(Equal(core.StringPtr("1.2.3.4")))
				Expect(createDnsRecordOptionsModel.Priority).To(Equal(core.Int64Ptr(int64(5))))
				Expect(createDnsRecordOptionsModel.Proxied).To(Equal(core.BoolPtr(false)))
				Expect(createDnsRecordOptionsModel.Data).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))
				Expect(createDnsRecordOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteDnsRecordOptions successfully`, func() {
				// Construct an instance of the DeleteDnsRecordOptions model
				dnsrecordIdentifier := "testString"
				deleteDnsRecordOptionsModel := dnsRecordsService.NewDeleteDnsRecordOptions(dnsrecordIdentifier)
				deleteDnsRecordOptionsModel.SetDnsrecordIdentifier("testString")
				deleteDnsRecordOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteDnsRecordOptionsModel).ToNot(BeNil())
				Expect(deleteDnsRecordOptionsModel.DnsrecordIdentifier).To(Equal(core.StringPtr("testString")))
				Expect(deleteDnsRecordOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetDnsRecordOptions successfully`, func() {
				// Construct an instance of the GetDnsRecordOptions model
				dnsrecordIdentifier := "testString"
				getDnsRecordOptionsModel := dnsRecordsService.NewGetDnsRecordOptions(dnsrecordIdentifier)
				getDnsRecordOptionsModel.SetDnsrecordIdentifier("testString")
				getDnsRecordOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getDnsRecordOptionsModel).ToNot(BeNil())
				Expect(getDnsRecordOptionsModel.DnsrecordIdentifier).To(Equal(core.StringPtr("testString")))
				Expect(getDnsRecordOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListAllDnsRecordsOptions successfully`, func() {
				// Construct an instance of the ListAllDnsRecordsOptions model
				listAllDnsRecordsOptionsModel := dnsRecordsService.NewListAllDnsRecordsOptions()
				listAllDnsRecordsOptionsModel.SetType("testString")
				listAllDnsRecordsOptionsModel.SetName("host1.test-example.com")
				listAllDnsRecordsOptionsModel.SetContent("1.2.3.4")
				listAllDnsRecordsOptionsModel.SetPage(int64(1))
				listAllDnsRecordsOptionsModel.SetPerPage(int64(20))
				listAllDnsRecordsOptionsModel.SetOrder("type")
				listAllDnsRecordsOptionsModel.SetDirection("asc")
				listAllDnsRecordsOptionsModel.SetMatch("any")
				listAllDnsRecordsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listAllDnsRecordsOptionsModel).ToNot(BeNil())
				Expect(listAllDnsRecordsOptionsModel.Type).To(Equal(core.StringPtr("testString")))
				Expect(listAllDnsRecordsOptionsModel.Name).To(Equal(core.StringPtr("host1.test-example.com")))
				Expect(listAllDnsRecordsOptionsModel.Content).To(Equal(core.StringPtr("1.2.3.4")))
				Expect(listAllDnsRecordsOptionsModel.Page).To(Equal(core.Int64Ptr(int64(1))))
				Expect(listAllDnsRecordsOptionsModel.PerPage).To(Equal(core.Int64Ptr(int64(20))))
				Expect(listAllDnsRecordsOptionsModel.Order).To(Equal(core.StringPtr("type")))
				Expect(listAllDnsRecordsOptionsModel.Direction).To(Equal(core.StringPtr("asc")))
				Expect(listAllDnsRecordsOptionsModel.Match).To(Equal(core.StringPtr("any")))
				Expect(listAllDnsRecordsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateDnsRecordOptions successfully`, func() {
				// Construct an instance of the UpdateDnsRecordOptions model
				dnsrecordIdentifier := "testString"
				updateDnsRecordOptionsModel := dnsRecordsService.NewUpdateDnsRecordOptions(dnsrecordIdentifier)
				updateDnsRecordOptionsModel.SetDnsrecordIdentifier("testString")
				updateDnsRecordOptionsModel.SetName("host-1.test-example.com")
				updateDnsRecordOptionsModel.SetType("A")
				updateDnsRecordOptionsModel.SetTTL(int64(120))
				updateDnsRecordOptionsModel.SetContent("1.2.3.4")
				updateDnsRecordOptionsModel.SetPriority(int64(5))
				updateDnsRecordOptionsModel.SetProxied(false)
				updateDnsRecordOptionsModel.SetData(map[string]interface{}{"anyKey": "anyValue"})
				updateDnsRecordOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateDnsRecordOptionsModel).ToNot(BeNil())
				Expect(updateDnsRecordOptionsModel.DnsrecordIdentifier).To(Equal(core.StringPtr("testString")))
				Expect(updateDnsRecordOptionsModel.Name).To(Equal(core.StringPtr("host-1.test-example.com")))
				Expect(updateDnsRecordOptionsModel.Type).To(Equal(core.StringPtr("A")))
				Expect(updateDnsRecordOptionsModel.TTL).To(Equal(core.Int64Ptr(int64(120))))
				Expect(updateDnsRecordOptionsModel.Content).To(Equal(core.StringPtr("1.2.3.4")))
				Expect(updateDnsRecordOptionsModel.Priority).To(Equal(core.Int64Ptr(int64(5))))
				Expect(updateDnsRecordOptionsModel.Proxied).To(Equal(core.BoolPtr(false)))
				Expect(updateDnsRecordOptionsModel.Data).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))
				Expect(updateDnsRecordOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDnsrecordInput successfully`, func() {
				typeVar := "A"
				_model, err := dnsRecordsService.NewDnsrecordInput(typeVar)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
		})
	})
	Describe(`Model unmarshaling tests`, func() {
		It(`Invoke UnmarshalBatchDnsRecordsRequestDeletesItem successfully`, func() {
			// Construct an instance of the model.
			model := new(dnsrecordsv1.BatchDnsRecordsRequestDeletesItem)
			model.ID = core.StringPtr("testString")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *dnsrecordsv1.BatchDnsRecordsRequestDeletesItem
			err = dnsrecordsv1.UnmarshalBatchDnsRecordsRequestDeletesItem(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalBatchDnsRecordsRequestPatchesItem successfully`, func() {
			// Construct an instance of the model.
			model := new(dnsrecordsv1.BatchDnsRecordsRequestPatchesItem)
			model.ID = core.StringPtr("testString")
			model.Name = core.StringPtr("host-1.test-example.com")
			model.Type = core.StringPtr("A")
			model.TTL = core.Int64Ptr(int64(120))
			model.Content = core.StringPtr("1.2.3.4")
			model.Priority = core.Int64Ptr(int64(5))
			model.Proxied = core.BoolPtr(false)
			model.Data = map[string]interface{}{"anyKey": "anyValue"}

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *dnsrecordsv1.BatchDnsRecordsRequestPatchesItem
			err = dnsrecordsv1.UnmarshalBatchDnsRecordsRequestPatchesItem(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalBatchDnsRecordsRequestPutsItem successfully`, func() {
			// Construct an instance of the model.
			model := new(dnsrecordsv1.BatchDnsRecordsRequestPutsItem)
			model.ID = core.StringPtr("testString")
			model.Name = core.StringPtr("host-1.test-example.com")
			model.Type = core.StringPtr("A")
			model.TTL = core.Int64Ptr(int64(120))
			model.Content = core.StringPtr("1.2.3.4")
			model.Priority = core.Int64Ptr(int64(5))
			model.Proxied = core.BoolPtr(false)
			model.Data = map[string]interface{}{"anyKey": "anyValue"}

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *dnsrecordsv1.BatchDnsRecordsRequestPutsItem
			err = dnsrecordsv1.UnmarshalBatchDnsRecordsRequestPutsItem(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalDnsrecordInput successfully`, func() {
			// Construct an instance of the model.
			model := new(dnsrecordsv1.DnsrecordInput)
			model.Name = core.StringPtr("host-1.test-example.com")
			model.Type = core.StringPtr("A")
			model.TTL = core.Int64Ptr(int64(120))
			model.Content = core.StringPtr("1.2.3.4")
			model.Priority = core.Int64Ptr(int64(5))
			model.Proxied = core.BoolPtr(false)
			model.Data = map[string]interface{}{"anyKey": "anyValue"}

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *dnsrecordsv1.DnsrecordInput
			err = dnsrecordsv1.UnmarshalDnsrecordInput(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
	})
	Describe(`Utility function tests`, func() {
		It(`Invoke CreateMockByteArray() successfully`, func() {
			mockByteArray := CreateMockByteArray("VGhpcyBpcyBhIHRlc3Qgb2YgdGhlIGVtZXJnZW5jeSBicm9hZGNhc3Qgc3lzdGVt")
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

func CreateMockByteArray(encodedString string) *[]byte {
	ba, err := base64.StdEncoding.DecodeString(encodedString)
	if err != nil {
		panic(err)
	}
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
