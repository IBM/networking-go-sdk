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

package resourcerecordsv1_test

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/networking-go-sdk/resourcerecordsv1"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe(`ResourceRecordsV1`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			testService, testServiceErr := resourcerecordsv1.NewResourceRecordsV1(&resourcerecordsv1.ResourceRecordsV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(testService).ToNot(BeNil())
			Expect(testServiceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			testService, testServiceErr := resourcerecordsv1.NewResourceRecordsV1(&resourcerecordsv1.ResourceRecordsV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(testService).To(BeNil())
			Expect(testServiceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			testService, testServiceErr := resourcerecordsv1.NewResourceRecordsV1(&resourcerecordsv1.ResourceRecordsV1Options{
				URL: "https://resourcerecordsv1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(testService).To(BeNil())
			Expect(testServiceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"RESOURCE_RECORDS_URL":       "https://resourcerecordsv1/api",
				"RESOURCE_RECORDS_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := resourcerecordsv1.NewResourceRecordsV1UsingExternalConfig(&resourcerecordsv1.ResourceRecordsV1Options{})
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := resourcerecordsv1.NewResourceRecordsV1UsingExternalConfig(&resourcerecordsv1.ResourceRecordsV1Options{
					URL: "https://testService/api",
				})
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				Expect(testService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := resourcerecordsv1.NewResourceRecordsV1UsingExternalConfig(&resourcerecordsv1.ResourceRecordsV1Options{})
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
				"RESOURCE_RECORDS_URL":       "https://resourcerecordsv1/api",
				"RESOURCE_RECORDS_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			testService, testServiceErr := resourcerecordsv1.NewResourceRecordsV1UsingExternalConfig(&resourcerecordsv1.ResourceRecordsV1Options{})

			It(`Instantiate service client with error`, func() {
				Expect(testService).To(BeNil())
				Expect(testServiceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"RESOURCE_RECORDS_AUTH_TYPE": "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			testService, testServiceErr := resourcerecordsv1.NewResourceRecordsV1UsingExternalConfig(&resourcerecordsv1.ResourceRecordsV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(testService).To(BeNil())
				Expect(testServiceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`ListResourceRecords(listResourceRecordsOptions *ListResourceRecordsOptions) - Operation response error`, func() {
		listResourceRecordsPath := "/instances/testString/dnszones/testString/resource_records"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(listResourceRecordsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(38))}))

					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(38))}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListResourceRecords with error: Operation response processing error`, func() {
				testService, testServiceErr := resourcerecordsv1.NewResourceRecordsV1(&resourcerecordsv1.ResourceRecordsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ListResourceRecordsOptions model
				listResourceRecordsOptionsModel := new(resourcerecordsv1.ListResourceRecordsOptions)
				listResourceRecordsOptionsModel.InstanceID = core.StringPtr("testString")
				listResourceRecordsOptionsModel.DnszoneID = core.StringPtr("testString")
				listResourceRecordsOptionsModel.XCorrelationID = core.StringPtr("testString")
				listResourceRecordsOptionsModel.Offset = core.Int64Ptr(int64(38))
				listResourceRecordsOptionsModel.Limit = core.Int64Ptr(int64(38))
				listResourceRecordsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.ListResourceRecords(listResourceRecordsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ListResourceRecords(listResourceRecordsOptions *ListResourceRecordsOptions)`, func() {
		listResourceRecordsPath := "/instances/testString/dnszones/testString/resource_records"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(listResourceRecordsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(38))}))

					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(38))}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"resource_records": [{"id": "SRV:5365b73c-ce6f-4d6f-ad9f-d9c131b26370", "created_on": "2019-01-01T05:20:00.12345Z", "modified_on": "2019-01-01T05:20:00.12345Z", "name": "_sip._udp.test.example.com", "type": "SRV", "ttl": 120, "rdata": {"anyKey": "anyValue"}, "service": "_sip", "protocol": "udp"}], "offset": 0, "limit": 20, "total_count": 200, "first": {"href": "https://api.dns-svcs.cloud.ibm.com/v1/instances/434f6c3e-6014-4124-a61d-2e910bca19b1/dnszones/example.com:d04d3a7a-7f6d-47d4-b811-08c5478fa1a4/resource_records?limit=20"}, "next": {"href": "https://api.dns-svcs.cloud.ibm.com/v1/instances/434f6c3e-6014-4124-a61d-2e910bca19b1/dnszones/example.com:d04d3a7a-7f6d-47d4-b811-08c5478fa1a4/resource_records?offset=20&limit=20"}}`)
				}))
			})
			It(`Invoke ListResourceRecords successfully`, func() {
				testService, testServiceErr := resourcerecordsv1.NewResourceRecordsV1(&resourcerecordsv1.ResourceRecordsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.ListResourceRecords(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListResourceRecordsOptions model
				listResourceRecordsOptionsModel := new(resourcerecordsv1.ListResourceRecordsOptions)
				listResourceRecordsOptionsModel.InstanceID = core.StringPtr("testString")
				listResourceRecordsOptionsModel.DnszoneID = core.StringPtr("testString")
				listResourceRecordsOptionsModel.XCorrelationID = core.StringPtr("testString")
				listResourceRecordsOptionsModel.Offset = core.Int64Ptr(int64(38))
				listResourceRecordsOptionsModel.Limit = core.Int64Ptr(int64(38))
				listResourceRecordsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.ListResourceRecords(listResourceRecordsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke ListResourceRecords with error: Operation validation and request error`, func() {
				testService, testServiceErr := resourcerecordsv1.NewResourceRecordsV1(&resourcerecordsv1.ResourceRecordsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ListResourceRecordsOptions model
				listResourceRecordsOptionsModel := new(resourcerecordsv1.ListResourceRecordsOptions)
				listResourceRecordsOptionsModel.InstanceID = core.StringPtr("testString")
				listResourceRecordsOptionsModel.DnszoneID = core.StringPtr("testString")
				listResourceRecordsOptionsModel.XCorrelationID = core.StringPtr("testString")
				listResourceRecordsOptionsModel.Offset = core.Int64Ptr(int64(38))
				listResourceRecordsOptionsModel.Limit = core.Int64Ptr(int64(38))
				listResourceRecordsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.ListResourceRecords(listResourceRecordsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListResourceRecordsOptions model with no property values
				listResourceRecordsOptionsModelNew := new(resourcerecordsv1.ListResourceRecordsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.ListResourceRecords(listResourceRecordsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateResourceRecord(createResourceRecordOptions *CreateResourceRecordOptions) - Operation response error`, func() {
		createResourceRecordPath := "/instances/testString/dnszones/testString/resource_records"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(createResourceRecordPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateResourceRecord with error: Operation response processing error`, func() {
				testService, testServiceErr := resourcerecordsv1.NewResourceRecordsV1(&resourcerecordsv1.ResourceRecordsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ResourceRecordInputRdataRdataARecord model
				resourceRecordInputRdataModel := new(resourcerecordsv1.ResourceRecordInputRdataRdataARecord)
				resourceRecordInputRdataModel.Ip = core.StringPtr("10.110.201.214")

				// Construct an instance of the CreateResourceRecordOptions model
				createResourceRecordOptionsModel := new(resourcerecordsv1.CreateResourceRecordOptions)
				createResourceRecordOptionsModel.InstanceID = core.StringPtr("testString")
				createResourceRecordOptionsModel.DnszoneID = core.StringPtr("testString")
				createResourceRecordOptionsModel.Name = core.StringPtr("test.example.com")
				createResourceRecordOptionsModel.Type = core.StringPtr("SRV")
				createResourceRecordOptionsModel.Rdata = resourceRecordInputRdataModel
				createResourceRecordOptionsModel.TTL = core.Int64Ptr(int64(120))
				createResourceRecordOptionsModel.Service = core.StringPtr("_sip")
				createResourceRecordOptionsModel.Protocol = core.StringPtr("udp")
				createResourceRecordOptionsModel.XCorrelationID = core.StringPtr("testString")
				createResourceRecordOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.CreateResourceRecord(createResourceRecordOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`CreateResourceRecord(createResourceRecordOptions *CreateResourceRecordOptions)`, func() {
		createResourceRecordPath := "/instances/testString/dnszones/testString/resource_records"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(createResourceRecordPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"id": "SRV:5365b73c-ce6f-4d6f-ad9f-d9c131b26370", "created_on": "2019-01-01T05:20:00.12345Z", "modified_on": "2019-01-01T05:20:00.12345Z", "name": "_sip._udp.test.example.com", "type": "SRV", "ttl": 120, "rdata": {"anyKey": "anyValue"}, "service": "_sip", "protocol": "udp"}`)
				}))
			})
			It(`Invoke CreateResourceRecord successfully`, func() {
				testService, testServiceErr := resourcerecordsv1.NewResourceRecordsV1(&resourcerecordsv1.ResourceRecordsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.CreateResourceRecord(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ResourceRecordInputRdataRdataARecord model
				resourceRecordInputRdataModel := new(resourcerecordsv1.ResourceRecordInputRdataRdataARecord)
				resourceRecordInputRdataModel.Ip = core.StringPtr("10.110.201.214")

				// Construct an instance of the CreateResourceRecordOptions model
				createResourceRecordOptionsModel := new(resourcerecordsv1.CreateResourceRecordOptions)
				createResourceRecordOptionsModel.InstanceID = core.StringPtr("testString")
				createResourceRecordOptionsModel.DnszoneID = core.StringPtr("testString")
				createResourceRecordOptionsModel.Name = core.StringPtr("test.example.com")
				createResourceRecordOptionsModel.Type = core.StringPtr("SRV")
				createResourceRecordOptionsModel.Rdata = resourceRecordInputRdataModel
				createResourceRecordOptionsModel.TTL = core.Int64Ptr(int64(120))
				createResourceRecordOptionsModel.Service = core.StringPtr("_sip")
				createResourceRecordOptionsModel.Protocol = core.StringPtr("udp")
				createResourceRecordOptionsModel.XCorrelationID = core.StringPtr("testString")
				createResourceRecordOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.CreateResourceRecord(createResourceRecordOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke CreateResourceRecord with error: Operation validation and request error`, func() {
				testService, testServiceErr := resourcerecordsv1.NewResourceRecordsV1(&resourcerecordsv1.ResourceRecordsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ResourceRecordInputRdataRdataARecord model
				resourceRecordInputRdataModel := new(resourcerecordsv1.ResourceRecordInputRdataRdataARecord)
				resourceRecordInputRdataModel.Ip = core.StringPtr("10.110.201.214")

				// Construct an instance of the CreateResourceRecordOptions model
				createResourceRecordOptionsModel := new(resourcerecordsv1.CreateResourceRecordOptions)
				createResourceRecordOptionsModel.InstanceID = core.StringPtr("testString")
				createResourceRecordOptionsModel.DnszoneID = core.StringPtr("testString")
				createResourceRecordOptionsModel.Name = core.StringPtr("test.example.com")
				createResourceRecordOptionsModel.Type = core.StringPtr("SRV")
				createResourceRecordOptionsModel.Rdata = resourceRecordInputRdataModel
				createResourceRecordOptionsModel.TTL = core.Int64Ptr(int64(120))
				createResourceRecordOptionsModel.Service = core.StringPtr("_sip")
				createResourceRecordOptionsModel.Protocol = core.StringPtr("udp")
				createResourceRecordOptionsModel.XCorrelationID = core.StringPtr("testString")
				createResourceRecordOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.CreateResourceRecord(createResourceRecordOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateResourceRecordOptions model with no property values
				createResourceRecordOptionsModelNew := new(resourcerecordsv1.CreateResourceRecordOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.CreateResourceRecord(createResourceRecordOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`DeleteResourceRecord(deleteResourceRecordOptions *DeleteResourceRecordOptions)`, func() {
		deleteResourceRecordPath := "/instances/testString/dnszones/testString/resource_records/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(deleteResourceRecordPath))
					Expect(req.Method).To(Equal("DELETE"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteResourceRecord successfully`, func() {
				testService, testServiceErr := resourcerecordsv1.NewResourceRecordsV1(&resourcerecordsv1.ResourceRecordsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := testService.DeleteResourceRecord(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteResourceRecordOptions model
				deleteResourceRecordOptionsModel := new(resourcerecordsv1.DeleteResourceRecordOptions)
				deleteResourceRecordOptionsModel.InstanceID = core.StringPtr("testString")
				deleteResourceRecordOptionsModel.DnszoneID = core.StringPtr("testString")
				deleteResourceRecordOptionsModel.RecordID = core.StringPtr("testString")
				deleteResourceRecordOptionsModel.XCorrelationID = core.StringPtr("testString")
				deleteResourceRecordOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = testService.DeleteResourceRecord(deleteResourceRecordOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteResourceRecord with error: Operation validation and request error`, func() {
				testService, testServiceErr := resourcerecordsv1.NewResourceRecordsV1(&resourcerecordsv1.ResourceRecordsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the DeleteResourceRecordOptions model
				deleteResourceRecordOptionsModel := new(resourcerecordsv1.DeleteResourceRecordOptions)
				deleteResourceRecordOptionsModel.InstanceID = core.StringPtr("testString")
				deleteResourceRecordOptionsModel.DnszoneID = core.StringPtr("testString")
				deleteResourceRecordOptionsModel.RecordID = core.StringPtr("testString")
				deleteResourceRecordOptionsModel.XCorrelationID = core.StringPtr("testString")
				deleteResourceRecordOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := testService.DeleteResourceRecord(deleteResourceRecordOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteResourceRecordOptions model with no property values
				deleteResourceRecordOptionsModelNew := new(resourcerecordsv1.DeleteResourceRecordOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = testService.DeleteResourceRecord(deleteResourceRecordOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetResourceRecord(getResourceRecordOptions *GetResourceRecordOptions) - Operation response error`, func() {
		getResourceRecordPath := "/instances/testString/dnszones/testString/resource_records/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getResourceRecordPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetResourceRecord with error: Operation response processing error`, func() {
				testService, testServiceErr := resourcerecordsv1.NewResourceRecordsV1(&resourcerecordsv1.ResourceRecordsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetResourceRecordOptions model
				getResourceRecordOptionsModel := new(resourcerecordsv1.GetResourceRecordOptions)
				getResourceRecordOptionsModel.InstanceID = core.StringPtr("testString")
				getResourceRecordOptionsModel.DnszoneID = core.StringPtr("testString")
				getResourceRecordOptionsModel.RecordID = core.StringPtr("testString")
				getResourceRecordOptionsModel.XCorrelationID = core.StringPtr("testString")
				getResourceRecordOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.GetResourceRecord(getResourceRecordOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetResourceRecord(getResourceRecordOptions *GetResourceRecordOptions)`, func() {
		getResourceRecordPath := "/instances/testString/dnszones/testString/resource_records/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getResourceRecordPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"id": "SRV:5365b73c-ce6f-4d6f-ad9f-d9c131b26370", "created_on": "2019-01-01T05:20:00.12345Z", "modified_on": "2019-01-01T05:20:00.12345Z", "name": "_sip._udp.test.example.com", "type": "SRV", "ttl": 120, "rdata": {"anyKey": "anyValue"}, "service": "_sip", "protocol": "udp"}`)
				}))
			})
			It(`Invoke GetResourceRecord successfully`, func() {
				testService, testServiceErr := resourcerecordsv1.NewResourceRecordsV1(&resourcerecordsv1.ResourceRecordsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetResourceRecord(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetResourceRecordOptions model
				getResourceRecordOptionsModel := new(resourcerecordsv1.GetResourceRecordOptions)
				getResourceRecordOptionsModel.InstanceID = core.StringPtr("testString")
				getResourceRecordOptionsModel.DnszoneID = core.StringPtr("testString")
				getResourceRecordOptionsModel.RecordID = core.StringPtr("testString")
				getResourceRecordOptionsModel.XCorrelationID = core.StringPtr("testString")
				getResourceRecordOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetResourceRecord(getResourceRecordOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke GetResourceRecord with error: Operation validation and request error`, func() {
				testService, testServiceErr := resourcerecordsv1.NewResourceRecordsV1(&resourcerecordsv1.ResourceRecordsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetResourceRecordOptions model
				getResourceRecordOptionsModel := new(resourcerecordsv1.GetResourceRecordOptions)
				getResourceRecordOptionsModel.InstanceID = core.StringPtr("testString")
				getResourceRecordOptionsModel.DnszoneID = core.StringPtr("testString")
				getResourceRecordOptionsModel.RecordID = core.StringPtr("testString")
				getResourceRecordOptionsModel.XCorrelationID = core.StringPtr("testString")
				getResourceRecordOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.GetResourceRecord(getResourceRecordOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetResourceRecordOptions model with no property values
				getResourceRecordOptionsModelNew := new(resourcerecordsv1.GetResourceRecordOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.GetResourceRecord(getResourceRecordOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateResourceRecord(updateResourceRecordOptions *UpdateResourceRecordOptions) - Operation response error`, func() {
		updateResourceRecordPath := "/instances/testString/dnszones/testString/resource_records/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(updateResourceRecordPath))
					Expect(req.Method).To(Equal("PUT"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateResourceRecord with error: Operation response processing error`, func() {
				testService, testServiceErr := resourcerecordsv1.NewResourceRecordsV1(&resourcerecordsv1.ResourceRecordsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ResourceRecordUpdateInputRdataRdataARecord model
				resourceRecordUpdateInputRdataModel := new(resourcerecordsv1.ResourceRecordUpdateInputRdataRdataARecord)
				resourceRecordUpdateInputRdataModel.Ip = core.StringPtr("10.110.201.214")

				// Construct an instance of the UpdateResourceRecordOptions model
				updateResourceRecordOptionsModel := new(resourcerecordsv1.UpdateResourceRecordOptions)
				updateResourceRecordOptionsModel.InstanceID = core.StringPtr("testString")
				updateResourceRecordOptionsModel.DnszoneID = core.StringPtr("testString")
				updateResourceRecordOptionsModel.RecordID = core.StringPtr("testString")
				updateResourceRecordOptionsModel.Name = core.StringPtr("test.example.com")
				updateResourceRecordOptionsModel.Rdata = resourceRecordUpdateInputRdataModel
				updateResourceRecordOptionsModel.TTL = core.Int64Ptr(int64(120))
				updateResourceRecordOptionsModel.Service = core.StringPtr("_sip")
				updateResourceRecordOptionsModel.Protocol = core.StringPtr("udp")
				updateResourceRecordOptionsModel.XCorrelationID = core.StringPtr("testString")
				updateResourceRecordOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.UpdateResourceRecord(updateResourceRecordOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UpdateResourceRecord(updateResourceRecordOptions *UpdateResourceRecordOptions)`, func() {
		updateResourceRecordPath := "/instances/testString/dnszones/testString/resource_records/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(updateResourceRecordPath))
					Expect(req.Method).To(Equal("PUT"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"id": "SRV:5365b73c-ce6f-4d6f-ad9f-d9c131b26370", "created_on": "2019-01-01T05:20:00.12345Z", "modified_on": "2019-01-01T05:20:00.12345Z", "name": "_sip._udp.test.example.com", "type": "SRV", "ttl": 120, "rdata": {"anyKey": "anyValue"}, "service": "_sip", "protocol": "udp"}`)
				}))
			})
			It(`Invoke UpdateResourceRecord successfully`, func() {
				testService, testServiceErr := resourcerecordsv1.NewResourceRecordsV1(&resourcerecordsv1.ResourceRecordsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.UpdateResourceRecord(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ResourceRecordUpdateInputRdataRdataARecord model
				resourceRecordUpdateInputRdataModel := new(resourcerecordsv1.ResourceRecordUpdateInputRdataRdataARecord)
				resourceRecordUpdateInputRdataModel.Ip = core.StringPtr("10.110.201.214")

				// Construct an instance of the UpdateResourceRecordOptions model
				updateResourceRecordOptionsModel := new(resourcerecordsv1.UpdateResourceRecordOptions)
				updateResourceRecordOptionsModel.InstanceID = core.StringPtr("testString")
				updateResourceRecordOptionsModel.DnszoneID = core.StringPtr("testString")
				updateResourceRecordOptionsModel.RecordID = core.StringPtr("testString")
				updateResourceRecordOptionsModel.Name = core.StringPtr("test.example.com")
				updateResourceRecordOptionsModel.Rdata = resourceRecordUpdateInputRdataModel
				updateResourceRecordOptionsModel.TTL = core.Int64Ptr(int64(120))
				updateResourceRecordOptionsModel.Service = core.StringPtr("_sip")
				updateResourceRecordOptionsModel.Protocol = core.StringPtr("udp")
				updateResourceRecordOptionsModel.XCorrelationID = core.StringPtr("testString")
				updateResourceRecordOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.UpdateResourceRecord(updateResourceRecordOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke UpdateResourceRecord with error: Operation validation and request error`, func() {
				testService, testServiceErr := resourcerecordsv1.NewResourceRecordsV1(&resourcerecordsv1.ResourceRecordsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ResourceRecordUpdateInputRdataRdataARecord model
				resourceRecordUpdateInputRdataModel := new(resourcerecordsv1.ResourceRecordUpdateInputRdataRdataARecord)
				resourceRecordUpdateInputRdataModel.Ip = core.StringPtr("10.110.201.214")

				// Construct an instance of the UpdateResourceRecordOptions model
				updateResourceRecordOptionsModel := new(resourcerecordsv1.UpdateResourceRecordOptions)
				updateResourceRecordOptionsModel.InstanceID = core.StringPtr("testString")
				updateResourceRecordOptionsModel.DnszoneID = core.StringPtr("testString")
				updateResourceRecordOptionsModel.RecordID = core.StringPtr("testString")
				updateResourceRecordOptionsModel.Name = core.StringPtr("test.example.com")
				updateResourceRecordOptionsModel.Rdata = resourceRecordUpdateInputRdataModel
				updateResourceRecordOptionsModel.TTL = core.Int64Ptr(int64(120))
				updateResourceRecordOptionsModel.Service = core.StringPtr("_sip")
				updateResourceRecordOptionsModel.Protocol = core.StringPtr("udp")
				updateResourceRecordOptionsModel.XCorrelationID = core.StringPtr("testString")
				updateResourceRecordOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.UpdateResourceRecord(updateResourceRecordOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateResourceRecordOptions model with no property values
				updateResourceRecordOptionsModelNew := new(resourcerecordsv1.UpdateResourceRecordOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.UpdateResourceRecord(updateResourceRecordOptionsModelNew)
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
			testService, _ := resourcerecordsv1.NewResourceRecordsV1(&resourcerecordsv1.ResourceRecordsV1Options{
				URL:           "http://resourcerecordsv1modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
			})
			It(`Invoke NewCreateResourceRecordOptions successfully`, func() {
				// Construct an instance of the ResourceRecordInputRdataRdataARecord model
				resourceRecordInputRdataModel := new(resourcerecordsv1.ResourceRecordInputRdataRdataARecord)
				Expect(resourceRecordInputRdataModel).ToNot(BeNil())
				resourceRecordInputRdataModel.Ip = core.StringPtr("10.110.201.214")
				Expect(resourceRecordInputRdataModel.Ip).To(Equal(core.StringPtr("10.110.201.214")))

				// Construct an instance of the CreateResourceRecordOptions model
				instanceID := "testString"
				dnszoneID := "testString"
				createResourceRecordOptionsModel := testService.NewCreateResourceRecordOptions(instanceID, dnszoneID)
				createResourceRecordOptionsModel.SetInstanceID("testString")
				createResourceRecordOptionsModel.SetDnszoneID("testString")
				createResourceRecordOptionsModel.SetName("test.example.com")
				createResourceRecordOptionsModel.SetType("SRV")
				createResourceRecordOptionsModel.SetRdata(resourceRecordInputRdataModel)
				createResourceRecordOptionsModel.SetTTL(int64(120))
				createResourceRecordOptionsModel.SetService("_sip")
				createResourceRecordOptionsModel.SetProtocol("udp")
				createResourceRecordOptionsModel.SetXCorrelationID("testString")
				createResourceRecordOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createResourceRecordOptionsModel).ToNot(BeNil())
				Expect(createResourceRecordOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(createResourceRecordOptionsModel.DnszoneID).To(Equal(core.StringPtr("testString")))
				Expect(createResourceRecordOptionsModel.Name).To(Equal(core.StringPtr("test.example.com")))
				Expect(createResourceRecordOptionsModel.Type).To(Equal(core.StringPtr("SRV")))
				Expect(createResourceRecordOptionsModel.Rdata).To(Equal(resourceRecordInputRdataModel))
				Expect(createResourceRecordOptionsModel.TTL).To(Equal(core.Int64Ptr(int64(120))))
				Expect(createResourceRecordOptionsModel.Service).To(Equal(core.StringPtr("_sip")))
				Expect(createResourceRecordOptionsModel.Protocol).To(Equal(core.StringPtr("udp")))
				Expect(createResourceRecordOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(createResourceRecordOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteResourceRecordOptions successfully`, func() {
				// Construct an instance of the DeleteResourceRecordOptions model
				instanceID := "testString"
				dnszoneID := "testString"
				recordID := "testString"
				deleteResourceRecordOptionsModel := testService.NewDeleteResourceRecordOptions(instanceID, dnszoneID, recordID)
				deleteResourceRecordOptionsModel.SetInstanceID("testString")
				deleteResourceRecordOptionsModel.SetDnszoneID("testString")
				deleteResourceRecordOptionsModel.SetRecordID("testString")
				deleteResourceRecordOptionsModel.SetXCorrelationID("testString")
				deleteResourceRecordOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteResourceRecordOptionsModel).ToNot(BeNil())
				Expect(deleteResourceRecordOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(deleteResourceRecordOptionsModel.DnszoneID).To(Equal(core.StringPtr("testString")))
				Expect(deleteResourceRecordOptionsModel.RecordID).To(Equal(core.StringPtr("testString")))
				Expect(deleteResourceRecordOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(deleteResourceRecordOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetResourceRecordOptions successfully`, func() {
				// Construct an instance of the GetResourceRecordOptions model
				instanceID := "testString"
				dnszoneID := "testString"
				recordID := "testString"
				getResourceRecordOptionsModel := testService.NewGetResourceRecordOptions(instanceID, dnszoneID, recordID)
				getResourceRecordOptionsModel.SetInstanceID("testString")
				getResourceRecordOptionsModel.SetDnszoneID("testString")
				getResourceRecordOptionsModel.SetRecordID("testString")
				getResourceRecordOptionsModel.SetXCorrelationID("testString")
				getResourceRecordOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getResourceRecordOptionsModel).ToNot(BeNil())
				Expect(getResourceRecordOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(getResourceRecordOptionsModel.DnszoneID).To(Equal(core.StringPtr("testString")))
				Expect(getResourceRecordOptionsModel.RecordID).To(Equal(core.StringPtr("testString")))
				Expect(getResourceRecordOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(getResourceRecordOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListResourceRecordsOptions successfully`, func() {
				// Construct an instance of the ListResourceRecordsOptions model
				instanceID := "testString"
				dnszoneID := "testString"
				listResourceRecordsOptionsModel := testService.NewListResourceRecordsOptions(instanceID, dnszoneID)
				listResourceRecordsOptionsModel.SetInstanceID("testString")
				listResourceRecordsOptionsModel.SetDnszoneID("testString")
				listResourceRecordsOptionsModel.SetXCorrelationID("testString")
				listResourceRecordsOptionsModel.SetOffset(int64(38))
				listResourceRecordsOptionsModel.SetLimit(int64(38))
				listResourceRecordsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listResourceRecordsOptionsModel).ToNot(BeNil())
				Expect(listResourceRecordsOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(listResourceRecordsOptionsModel.DnszoneID).To(Equal(core.StringPtr("testString")))
				Expect(listResourceRecordsOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(listResourceRecordsOptionsModel.Offset).To(Equal(core.Int64Ptr(int64(38))))
				Expect(listResourceRecordsOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(38))))
				Expect(listResourceRecordsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateResourceRecordOptions successfully`, func() {
				// Construct an instance of the ResourceRecordUpdateInputRdataRdataARecord model
				resourceRecordUpdateInputRdataModel := new(resourcerecordsv1.ResourceRecordUpdateInputRdataRdataARecord)
				Expect(resourceRecordUpdateInputRdataModel).ToNot(BeNil())
				resourceRecordUpdateInputRdataModel.Ip = core.StringPtr("10.110.201.214")
				Expect(resourceRecordUpdateInputRdataModel.Ip).To(Equal(core.StringPtr("10.110.201.214")))

				// Construct an instance of the UpdateResourceRecordOptions model
				instanceID := "testString"
				dnszoneID := "testString"
				recordID := "testString"
				updateResourceRecordOptionsModel := testService.NewUpdateResourceRecordOptions(instanceID, dnszoneID, recordID)
				updateResourceRecordOptionsModel.SetInstanceID("testString")
				updateResourceRecordOptionsModel.SetDnszoneID("testString")
				updateResourceRecordOptionsModel.SetRecordID("testString")
				updateResourceRecordOptionsModel.SetName("test.example.com")
				updateResourceRecordOptionsModel.SetRdata(resourceRecordUpdateInputRdataModel)
				updateResourceRecordOptionsModel.SetTTL(int64(120))
				updateResourceRecordOptionsModel.SetService("_sip")
				updateResourceRecordOptionsModel.SetProtocol("udp")
				updateResourceRecordOptionsModel.SetXCorrelationID("testString")
				updateResourceRecordOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateResourceRecordOptionsModel).ToNot(BeNil())
				Expect(updateResourceRecordOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(updateResourceRecordOptionsModel.DnszoneID).To(Equal(core.StringPtr("testString")))
				Expect(updateResourceRecordOptionsModel.RecordID).To(Equal(core.StringPtr("testString")))
				Expect(updateResourceRecordOptionsModel.Name).To(Equal(core.StringPtr("test.example.com")))
				Expect(updateResourceRecordOptionsModel.Rdata).To(Equal(resourceRecordUpdateInputRdataModel))
				Expect(updateResourceRecordOptionsModel.TTL).To(Equal(core.Int64Ptr(int64(120))))
				Expect(updateResourceRecordOptionsModel.Service).To(Equal(core.StringPtr("_sip")))
				Expect(updateResourceRecordOptionsModel.Protocol).To(Equal(core.StringPtr("udp")))
				Expect(updateResourceRecordOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(updateResourceRecordOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewResourceRecordInputRdataRdataARecord successfully`, func() {
				ip := "10.110.201.214"
				model, err := testService.NewResourceRecordInputRdataRdataARecord(ip)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewResourceRecordInputRdataRdataAaaaRecord successfully`, func() {
				ip := "2019::2019"
				model, err := testService.NewResourceRecordInputRdataRdataAaaaRecord(ip)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewResourceRecordInputRdataRdataCnameRecord successfully`, func() {
				cname := "www.example.com"
				model, err := testService.NewResourceRecordInputRdataRdataCnameRecord(cname)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewResourceRecordInputRdataRdataMxRecord successfully`, func() {
				exchange := "mail.example.com"
				preference := int64(10)
				model, err := testService.NewResourceRecordInputRdataRdataMxRecord(exchange, preference)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewResourceRecordInputRdataRdataPtrRecord successfully`, func() {
				ptrdname := "www.example.com"
				model, err := testService.NewResourceRecordInputRdataRdataPtrRecord(ptrdname)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewResourceRecordInputRdataRdataSrvRecord successfully`, func() {
				port := int64(80)
				priority := int64(10)
				target := "www.example.com"
				weight := int64(10)
				model, err := testService.NewResourceRecordInputRdataRdataSrvRecord(port, priority, target, weight)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewResourceRecordInputRdataRdataTxtRecord successfully`, func() {
				text := "This is a text record"
				model, err := testService.NewResourceRecordInputRdataRdataTxtRecord(text)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewResourceRecordUpdateInputRdataRdataARecord successfully`, func() {
				ip := "10.110.201.214"
				model, err := testService.NewResourceRecordUpdateInputRdataRdataARecord(ip)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewResourceRecordUpdateInputRdataRdataAaaaRecord successfully`, func() {
				ip := "2019::2019"
				model, err := testService.NewResourceRecordUpdateInputRdataRdataAaaaRecord(ip)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewResourceRecordUpdateInputRdataRdataCnameRecord successfully`, func() {
				cname := "www.example.com"
				model, err := testService.NewResourceRecordUpdateInputRdataRdataCnameRecord(cname)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewResourceRecordUpdateInputRdataRdataMxRecord successfully`, func() {
				exchange := "mail.example.com"
				preference := int64(10)
				model, err := testService.NewResourceRecordUpdateInputRdataRdataMxRecord(exchange, preference)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewResourceRecordUpdateInputRdataRdataPtrRecord successfully`, func() {
				ptrdname := "www.example.com"
				model, err := testService.NewResourceRecordUpdateInputRdataRdataPtrRecord(ptrdname)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewResourceRecordUpdateInputRdataRdataSrvRecord successfully`, func() {
				port := int64(80)
				priority := int64(10)
				target := "www.example.com"
				weight := int64(10)
				model, err := testService.NewResourceRecordUpdateInputRdataRdataSrvRecord(port, priority, target, weight)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewResourceRecordUpdateInputRdataRdataTxtRecord successfully`, func() {
				text := "This is a text record"
				model, err := testService.NewResourceRecordUpdateInputRdataRdataTxtRecord(text)
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
