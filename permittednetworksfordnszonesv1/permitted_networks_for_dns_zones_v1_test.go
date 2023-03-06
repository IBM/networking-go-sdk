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

package permittednetworksfordnszonesv1_test

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/networking-go-sdk/permittednetworksfordnszonesv1"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe(`PermittedNetworksForDnsZonesV1`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			testService, testServiceErr := permittednetworksfordnszonesv1.NewPermittedNetworksForDnsZonesV1(&permittednetworksfordnszonesv1.PermittedNetworksForDnsZonesV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(testService).ToNot(BeNil())
			Expect(testServiceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			testService, testServiceErr := permittednetworksfordnszonesv1.NewPermittedNetworksForDnsZonesV1(&permittednetworksfordnszonesv1.PermittedNetworksForDnsZonesV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(testService).To(BeNil())
			Expect(testServiceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			testService, testServiceErr := permittednetworksfordnszonesv1.NewPermittedNetworksForDnsZonesV1(&permittednetworksfordnszonesv1.PermittedNetworksForDnsZonesV1Options{
				URL: "https://permittednetworksfordnszonesv1/api",
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
				"PERMITTED_NETWORKS_FOR_DNS_ZONES_URL":       "https://permittednetworksfordnszonesv1/api",
				"PERMITTED_NETWORKS_FOR_DNS_ZONES_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := permittednetworksfordnszonesv1.NewPermittedNetworksForDnsZonesV1UsingExternalConfig(&permittednetworksfordnszonesv1.PermittedNetworksForDnsZonesV1Options{})
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := permittednetworksfordnszonesv1.NewPermittedNetworksForDnsZonesV1UsingExternalConfig(&permittednetworksfordnszonesv1.PermittedNetworksForDnsZonesV1Options{
					URL: "https://testService/api",
				})
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				Expect(testService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := permittednetworksfordnszonesv1.NewPermittedNetworksForDnsZonesV1UsingExternalConfig(&permittednetworksfordnszonesv1.PermittedNetworksForDnsZonesV1Options{})
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
				"PERMITTED_NETWORKS_FOR_DNS_ZONES_URL":       "https://permittednetworksfordnszonesv1/api",
				"PERMITTED_NETWORKS_FOR_DNS_ZONES_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			testService, testServiceErr := permittednetworksfordnszonesv1.NewPermittedNetworksForDnsZonesV1UsingExternalConfig(&permittednetworksfordnszonesv1.PermittedNetworksForDnsZonesV1Options{})

			It(`Instantiate service client with error`, func() {
				Expect(testService).To(BeNil())
				Expect(testServiceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"PERMITTED_NETWORKS_FOR_DNS_ZONES_AUTH_TYPE": "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			testService, testServiceErr := permittednetworksfordnszonesv1.NewPermittedNetworksForDnsZonesV1UsingExternalConfig(&permittednetworksfordnszonesv1.PermittedNetworksForDnsZonesV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(testService).To(BeNil())
				Expect(testServiceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`ListPermittedNetworks(listPermittedNetworksOptions *ListPermittedNetworksOptions) - Operation response error`, func() {
		listPermittedNetworksPath := "/instances/testString/dnszones/testString/permitted_networks"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(listPermittedNetworksPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["offset"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["limit"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListPermittedNetworks with error: Operation response processing error`, func() {
				testService, testServiceErr := permittednetworksfordnszonesv1.NewPermittedNetworksForDnsZonesV1(&permittednetworksfordnszonesv1.PermittedNetworksForDnsZonesV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ListPermittedNetworksOptions model
				listPermittedNetworksOptionsModel := new(permittednetworksfordnszonesv1.ListPermittedNetworksOptions)
				listPermittedNetworksOptionsModel.InstanceID = core.StringPtr("testString")
				listPermittedNetworksOptionsModel.DnszoneID = core.StringPtr("testString")
				listPermittedNetworksOptionsModel.XCorrelationID = core.StringPtr("testString")
				listPermittedNetworksOptionsModel.Offset = core.StringPtr("testString")
				listPermittedNetworksOptionsModel.Limit = core.StringPtr("testString")
				listPermittedNetworksOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.ListPermittedNetworks(listPermittedNetworksOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ListPermittedNetworks(listPermittedNetworksOptions *ListPermittedNetworksOptions)`, func() {
		listPermittedNetworksPath := "/instances/testString/dnszones/testString/permitted_networks"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(listPermittedNetworksPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["offset"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["limit"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"permitted_networks": [{"id": "fecd0173-3919-456b-b202-3029dfa1b0f7", "created_on": "2019-01-01T05:20:00.12345Z", "modified_on": "2019-01-01T05:20:00.12345Z", "permitted_network": {"vpc_crn": "crn:v1:bluemix:public:is:eu-de:a/bcf1865e99742d38d2d5fc3fb80a5496::vpc:6e6cc326-04d1-4c99-a289-efb3ae4193d6"}, "type": "vpc", "state": "ACTIVE"}], "offset": 0, "limit": 10, "total_count": 200, "first": {"href": "https://api.dns-svcs.cloud.ibm.com/v1/instances/434f6c3e-6014-4124-a61d-2e910bca19b1/dnszones/example.com:252926c6-7d0c-4d37-861a-1faca0041785/permitted_networks?limit=10"}, "next": {"href": "https://api.dns-svcs.cloud.ibm.com/v1/instances/434f6c3e-6014-4124-a61d-2e910bca19b1/dnszones/example.com:252926c6-7d0c-4d37-861a-1faca0041785/permitted_networks?offset=1&limit=10"}}`)
				}))
			})
			It(`Invoke ListPermittedNetworks successfully`, func() {
				testService, testServiceErr := permittednetworksfordnszonesv1.NewPermittedNetworksForDnsZonesV1(&permittednetworksfordnszonesv1.PermittedNetworksForDnsZonesV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.ListPermittedNetworks(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListPermittedNetworksOptions model
				listPermittedNetworksOptionsModel := new(permittednetworksfordnszonesv1.ListPermittedNetworksOptions)
				listPermittedNetworksOptionsModel.InstanceID = core.StringPtr("testString")
				listPermittedNetworksOptionsModel.DnszoneID = core.StringPtr("testString")
				listPermittedNetworksOptionsModel.XCorrelationID = core.StringPtr("testString")
				listPermittedNetworksOptionsModel.Offset = core.StringPtr("testString")
				listPermittedNetworksOptionsModel.Limit = core.StringPtr("testString")
				listPermittedNetworksOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.ListPermittedNetworks(listPermittedNetworksOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke ListPermittedNetworks with error: Operation validation and request error`, func() {
				testService, testServiceErr := permittednetworksfordnszonesv1.NewPermittedNetworksForDnsZonesV1(&permittednetworksfordnszonesv1.PermittedNetworksForDnsZonesV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ListPermittedNetworksOptions model
				listPermittedNetworksOptionsModel := new(permittednetworksfordnszonesv1.ListPermittedNetworksOptions)
				listPermittedNetworksOptionsModel.InstanceID = core.StringPtr("testString")
				listPermittedNetworksOptionsModel.DnszoneID = core.StringPtr("testString")
				listPermittedNetworksOptionsModel.XCorrelationID = core.StringPtr("testString")
				listPermittedNetworksOptionsModel.Offset = core.StringPtr("testString")
				listPermittedNetworksOptionsModel.Limit = core.StringPtr("testString")
				listPermittedNetworksOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.ListPermittedNetworks(listPermittedNetworksOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListPermittedNetworksOptions model with no property values
				listPermittedNetworksOptionsModelNew := new(permittednetworksfordnszonesv1.ListPermittedNetworksOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.ListPermittedNetworks(listPermittedNetworksOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreatePermittedNetwork(createPermittedNetworkOptions *CreatePermittedNetworkOptions) - Operation response error`, func() {
		createPermittedNetworkPath := "/instances/testString/dnszones/testString/permitted_networks"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(createPermittedNetworkPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreatePermittedNetwork with error: Operation response processing error`, func() {
				testService, testServiceErr := permittednetworksfordnszonesv1.NewPermittedNetworksForDnsZonesV1(&permittednetworksfordnszonesv1.PermittedNetworksForDnsZonesV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the PermittedNetworkVpc model
				permittedNetworkVpcModel := new(permittednetworksfordnszonesv1.PermittedNetworkVpc)
				permittedNetworkVpcModel.VpcCrn = core.StringPtr("crn:v1:bluemix:public:is:eu-de:a/bcf1865e99742d38d2d5fc3fb80a5496::vpc:6e6cc326-04d1-4c99-a289-efb3ae4193d6")

				// Construct an instance of the CreatePermittedNetworkOptions model
				createPermittedNetworkOptionsModel := new(permittednetworksfordnszonesv1.CreatePermittedNetworkOptions)
				createPermittedNetworkOptionsModel.InstanceID = core.StringPtr("testString")
				createPermittedNetworkOptionsModel.DnszoneID = core.StringPtr("testString")
				createPermittedNetworkOptionsModel.Type = core.StringPtr("vpc")
				createPermittedNetworkOptionsModel.PermittedNetwork = permittedNetworkVpcModel
				createPermittedNetworkOptionsModel.XCorrelationID = core.StringPtr("testString")
				createPermittedNetworkOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.CreatePermittedNetwork(createPermittedNetworkOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`CreatePermittedNetwork(createPermittedNetworkOptions *CreatePermittedNetworkOptions)`, func() {
		createPermittedNetworkPath := "/instances/testString/dnszones/testString/permitted_networks"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(createPermittedNetworkPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"id": "fecd0173-3919-456b-b202-3029dfa1b0f7", "created_on": "2019-01-01T05:20:00.12345Z", "modified_on": "2019-01-01T05:20:00.12345Z", "permitted_network": {"vpc_crn": "crn:v1:bluemix:public:is:eu-de:a/bcf1865e99742d38d2d5fc3fb80a5496::vpc:6e6cc326-04d1-4c99-a289-efb3ae4193d6"}, "type": "vpc", "state": "ACTIVE"}`)
				}))
			})
			It(`Invoke CreatePermittedNetwork successfully`, func() {
				testService, testServiceErr := permittednetworksfordnszonesv1.NewPermittedNetworksForDnsZonesV1(&permittednetworksfordnszonesv1.PermittedNetworksForDnsZonesV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.CreatePermittedNetwork(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the PermittedNetworkVpc model
				permittedNetworkVpcModel := new(permittednetworksfordnszonesv1.PermittedNetworkVpc)
				permittedNetworkVpcModel.VpcCrn = core.StringPtr("crn:v1:bluemix:public:is:eu-de:a/bcf1865e99742d38d2d5fc3fb80a5496::vpc:6e6cc326-04d1-4c99-a289-efb3ae4193d6")

				// Construct an instance of the CreatePermittedNetworkOptions model
				createPermittedNetworkOptionsModel := new(permittednetworksfordnszonesv1.CreatePermittedNetworkOptions)
				createPermittedNetworkOptionsModel.InstanceID = core.StringPtr("testString")
				createPermittedNetworkOptionsModel.DnszoneID = core.StringPtr("testString")
				createPermittedNetworkOptionsModel.Type = core.StringPtr("vpc")
				createPermittedNetworkOptionsModel.PermittedNetwork = permittedNetworkVpcModel
				createPermittedNetworkOptionsModel.XCorrelationID = core.StringPtr("testString")
				createPermittedNetworkOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.CreatePermittedNetwork(createPermittedNetworkOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke CreatePermittedNetwork with error: Operation validation and request error`, func() {
				testService, testServiceErr := permittednetworksfordnszonesv1.NewPermittedNetworksForDnsZonesV1(&permittednetworksfordnszonesv1.PermittedNetworksForDnsZonesV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the PermittedNetworkVpc model
				permittedNetworkVpcModel := new(permittednetworksfordnszonesv1.PermittedNetworkVpc)
				permittedNetworkVpcModel.VpcCrn = core.StringPtr("crn:v1:bluemix:public:is:eu-de:a/bcf1865e99742d38d2d5fc3fb80a5496::vpc:6e6cc326-04d1-4c99-a289-efb3ae4193d6")

				// Construct an instance of the CreatePermittedNetworkOptions model
				createPermittedNetworkOptionsModel := new(permittednetworksfordnszonesv1.CreatePermittedNetworkOptions)
				createPermittedNetworkOptionsModel.InstanceID = core.StringPtr("testString")
				createPermittedNetworkOptionsModel.DnszoneID = core.StringPtr("testString")
				createPermittedNetworkOptionsModel.Type = core.StringPtr("vpc")
				createPermittedNetworkOptionsModel.PermittedNetwork = permittedNetworkVpcModel
				createPermittedNetworkOptionsModel.XCorrelationID = core.StringPtr("testString")
				createPermittedNetworkOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.CreatePermittedNetwork(createPermittedNetworkOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreatePermittedNetworkOptions model with no property values
				createPermittedNetworkOptionsModelNew := new(permittednetworksfordnszonesv1.CreatePermittedNetworkOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.CreatePermittedNetwork(createPermittedNetworkOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeletePermittedNetwork(deletePermittedNetworkOptions *DeletePermittedNetworkOptions) - Operation response error`, func() {
		deletePermittedNetworkPath := "/instances/testString/dnszones/testString/permitted_networks/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(deletePermittedNetworkPath))
					Expect(req.Method).To(Equal("DELETE"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke DeletePermittedNetwork with error: Operation response processing error`, func() {
				testService, testServiceErr := permittednetworksfordnszonesv1.NewPermittedNetworksForDnsZonesV1(&permittednetworksfordnszonesv1.PermittedNetworksForDnsZonesV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the DeletePermittedNetworkOptions model
				deletePermittedNetworkOptionsModel := new(permittednetworksfordnszonesv1.DeletePermittedNetworkOptions)
				deletePermittedNetworkOptionsModel.InstanceID = core.StringPtr("testString")
				deletePermittedNetworkOptionsModel.DnszoneID = core.StringPtr("testString")
				deletePermittedNetworkOptionsModel.PermittedNetworkID = core.StringPtr("testString")
				deletePermittedNetworkOptionsModel.XCorrelationID = core.StringPtr("testString")
				deletePermittedNetworkOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.DeletePermittedNetwork(deletePermittedNetworkOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`DeletePermittedNetwork(deletePermittedNetworkOptions *DeletePermittedNetworkOptions)`, func() {
		deletePermittedNetworkPath := "/instances/testString/dnszones/testString/permitted_networks/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(deletePermittedNetworkPath))
					Expect(req.Method).To(Equal("DELETE"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, `{"id": "fecd0173-3919-456b-b202-3029dfa1b0f7", "created_on": "2019-01-01T05:20:00.12345Z", "modified_on": "2019-01-01T05:20:00.12345Z", "permitted_network": {"vpc_crn": "crn:v1:bluemix:public:is:eu-de:a/bcf1865e99742d38d2d5fc3fb80a5496::vpc:6e6cc326-04d1-4c99-a289-efb3ae4193d6"}, "type": "vpc", "state": "ACTIVE"}`)
				}))
			})
			It(`Invoke DeletePermittedNetwork successfully`, func() {
				testService, testServiceErr := permittednetworksfordnszonesv1.NewPermittedNetworksForDnsZonesV1(&permittednetworksfordnszonesv1.PermittedNetworksForDnsZonesV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.DeletePermittedNetwork(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DeletePermittedNetworkOptions model
				deletePermittedNetworkOptionsModel := new(permittednetworksfordnszonesv1.DeletePermittedNetworkOptions)
				deletePermittedNetworkOptionsModel.InstanceID = core.StringPtr("testString")
				deletePermittedNetworkOptionsModel.DnszoneID = core.StringPtr("testString")
				deletePermittedNetworkOptionsModel.PermittedNetworkID = core.StringPtr("testString")
				deletePermittedNetworkOptionsModel.XCorrelationID = core.StringPtr("testString")
				deletePermittedNetworkOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.DeletePermittedNetwork(deletePermittedNetworkOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke DeletePermittedNetwork with error: Operation validation and request error`, func() {
				testService, testServiceErr := permittednetworksfordnszonesv1.NewPermittedNetworksForDnsZonesV1(&permittednetworksfordnszonesv1.PermittedNetworksForDnsZonesV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the DeletePermittedNetworkOptions model
				deletePermittedNetworkOptionsModel := new(permittednetworksfordnszonesv1.DeletePermittedNetworkOptions)
				deletePermittedNetworkOptionsModel.InstanceID = core.StringPtr("testString")
				deletePermittedNetworkOptionsModel.DnszoneID = core.StringPtr("testString")
				deletePermittedNetworkOptionsModel.PermittedNetworkID = core.StringPtr("testString")
				deletePermittedNetworkOptionsModel.XCorrelationID = core.StringPtr("testString")
				deletePermittedNetworkOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.DeletePermittedNetwork(deletePermittedNetworkOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the DeletePermittedNetworkOptions model with no property values
				deletePermittedNetworkOptionsModelNew := new(permittednetworksfordnszonesv1.DeletePermittedNetworkOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.DeletePermittedNetwork(deletePermittedNetworkOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetPermittedNetwork(getPermittedNetworkOptions *GetPermittedNetworkOptions) - Operation response error`, func() {
		getPermittedNetworkPath := "/instances/testString/dnszones/testString/permitted_networks/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getPermittedNetworkPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetPermittedNetwork with error: Operation response processing error`, func() {
				testService, testServiceErr := permittednetworksfordnszonesv1.NewPermittedNetworksForDnsZonesV1(&permittednetworksfordnszonesv1.PermittedNetworksForDnsZonesV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetPermittedNetworkOptions model
				getPermittedNetworkOptionsModel := new(permittednetworksfordnszonesv1.GetPermittedNetworkOptions)
				getPermittedNetworkOptionsModel.InstanceID = core.StringPtr("testString")
				getPermittedNetworkOptionsModel.DnszoneID = core.StringPtr("testString")
				getPermittedNetworkOptionsModel.PermittedNetworkID = core.StringPtr("testString")
				getPermittedNetworkOptionsModel.XCorrelationID = core.StringPtr("testString")
				getPermittedNetworkOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.GetPermittedNetwork(getPermittedNetworkOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetPermittedNetwork(getPermittedNetworkOptions *GetPermittedNetworkOptions)`, func() {
		getPermittedNetworkPath := "/instances/testString/dnszones/testString/permitted_networks/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getPermittedNetworkPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"id": "fecd0173-3919-456b-b202-3029dfa1b0f7", "created_on": "2019-01-01T05:20:00.12345Z", "modified_on": "2019-01-01T05:20:00.12345Z", "permitted_network": {"vpc_crn": "crn:v1:bluemix:public:is:eu-de:a/bcf1865e99742d38d2d5fc3fb80a5496::vpc:6e6cc326-04d1-4c99-a289-efb3ae4193d6"}, "type": "vpc", "state": "ACTIVE"}`)
				}))
			})
			It(`Invoke GetPermittedNetwork successfully`, func() {
				testService, testServiceErr := permittednetworksfordnszonesv1.NewPermittedNetworksForDnsZonesV1(&permittednetworksfordnszonesv1.PermittedNetworksForDnsZonesV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetPermittedNetwork(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetPermittedNetworkOptions model
				getPermittedNetworkOptionsModel := new(permittednetworksfordnszonesv1.GetPermittedNetworkOptions)
				getPermittedNetworkOptionsModel.InstanceID = core.StringPtr("testString")
				getPermittedNetworkOptionsModel.DnszoneID = core.StringPtr("testString")
				getPermittedNetworkOptionsModel.PermittedNetworkID = core.StringPtr("testString")
				getPermittedNetworkOptionsModel.XCorrelationID = core.StringPtr("testString")
				getPermittedNetworkOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetPermittedNetwork(getPermittedNetworkOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke GetPermittedNetwork with error: Operation validation and request error`, func() {
				testService, testServiceErr := permittednetworksfordnszonesv1.NewPermittedNetworksForDnsZonesV1(&permittednetworksfordnszonesv1.PermittedNetworksForDnsZonesV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetPermittedNetworkOptions model
				getPermittedNetworkOptionsModel := new(permittednetworksfordnszonesv1.GetPermittedNetworkOptions)
				getPermittedNetworkOptionsModel.InstanceID = core.StringPtr("testString")
				getPermittedNetworkOptionsModel.DnszoneID = core.StringPtr("testString")
				getPermittedNetworkOptionsModel.PermittedNetworkID = core.StringPtr("testString")
				getPermittedNetworkOptionsModel.XCorrelationID = core.StringPtr("testString")
				getPermittedNetworkOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.GetPermittedNetwork(getPermittedNetworkOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetPermittedNetworkOptions model with no property values
				getPermittedNetworkOptionsModelNew := new(permittednetworksfordnszonesv1.GetPermittedNetworkOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.GetPermittedNetwork(getPermittedNetworkOptionsModelNew)
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
			testService, _ := permittednetworksfordnszonesv1.NewPermittedNetworksForDnsZonesV1(&permittednetworksfordnszonesv1.PermittedNetworksForDnsZonesV1Options{
				URL:           "http://permittednetworksfordnszonesv1modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
			})
			It(`Invoke NewCreatePermittedNetworkOptions successfully`, func() {
				// Construct an instance of the PermittedNetworkVpc model
				permittedNetworkVpcModel := new(permittednetworksfordnszonesv1.PermittedNetworkVpc)
				Expect(permittedNetworkVpcModel).ToNot(BeNil())
				permittedNetworkVpcModel.VpcCrn = core.StringPtr("crn:v1:bluemix:public:is:eu-de:a/bcf1865e99742d38d2d5fc3fb80a5496::vpc:6e6cc326-04d1-4c99-a289-efb3ae4193d6")
				Expect(permittedNetworkVpcModel.VpcCrn).To(Equal(core.StringPtr("crn:v1:bluemix:public:is:eu-de:a/bcf1865e99742d38d2d5fc3fb80a5496::vpc:6e6cc326-04d1-4c99-a289-efb3ae4193d6")))

				// Construct an instance of the CreatePermittedNetworkOptions model
				instanceID := "testString"
				dnszoneID := "testString"
				createPermittedNetworkOptionsModel := testService.NewCreatePermittedNetworkOptions(instanceID, dnszoneID)
				createPermittedNetworkOptionsModel.SetInstanceID("testString")
				createPermittedNetworkOptionsModel.SetDnszoneID("testString")
				createPermittedNetworkOptionsModel.SetType("vpc")
				createPermittedNetworkOptionsModel.SetPermittedNetwork(permittedNetworkVpcModel)
				createPermittedNetworkOptionsModel.SetXCorrelationID("testString")
				createPermittedNetworkOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createPermittedNetworkOptionsModel).ToNot(BeNil())
				Expect(createPermittedNetworkOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(createPermittedNetworkOptionsModel.DnszoneID).To(Equal(core.StringPtr("testString")))
				Expect(createPermittedNetworkOptionsModel.Type).To(Equal(core.StringPtr("vpc")))
				Expect(createPermittedNetworkOptionsModel.PermittedNetwork).To(Equal(permittedNetworkVpcModel))
				Expect(createPermittedNetworkOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(createPermittedNetworkOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeletePermittedNetworkOptions successfully`, func() {
				// Construct an instance of the DeletePermittedNetworkOptions model
				instanceID := "testString"
				dnszoneID := "testString"
				permittedNetworkID := "testString"
				deletePermittedNetworkOptionsModel := testService.NewDeletePermittedNetworkOptions(instanceID, dnszoneID, permittedNetworkID)
				deletePermittedNetworkOptionsModel.SetInstanceID("testString")
				deletePermittedNetworkOptionsModel.SetDnszoneID("testString")
				deletePermittedNetworkOptionsModel.SetPermittedNetworkID("testString")
				deletePermittedNetworkOptionsModel.SetXCorrelationID("testString")
				deletePermittedNetworkOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deletePermittedNetworkOptionsModel).ToNot(BeNil())
				Expect(deletePermittedNetworkOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(deletePermittedNetworkOptionsModel.DnszoneID).To(Equal(core.StringPtr("testString")))
				Expect(deletePermittedNetworkOptionsModel.PermittedNetworkID).To(Equal(core.StringPtr("testString")))
				Expect(deletePermittedNetworkOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(deletePermittedNetworkOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetPermittedNetworkOptions successfully`, func() {
				// Construct an instance of the GetPermittedNetworkOptions model
				instanceID := "testString"
				dnszoneID := "testString"
				permittedNetworkID := "testString"
				getPermittedNetworkOptionsModel := testService.NewGetPermittedNetworkOptions(instanceID, dnszoneID, permittedNetworkID)
				getPermittedNetworkOptionsModel.SetInstanceID("testString")
				getPermittedNetworkOptionsModel.SetDnszoneID("testString")
				getPermittedNetworkOptionsModel.SetPermittedNetworkID("testString")
				getPermittedNetworkOptionsModel.SetXCorrelationID("testString")
				getPermittedNetworkOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getPermittedNetworkOptionsModel).ToNot(BeNil())
				Expect(getPermittedNetworkOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(getPermittedNetworkOptionsModel.DnszoneID).To(Equal(core.StringPtr("testString")))
				Expect(getPermittedNetworkOptionsModel.PermittedNetworkID).To(Equal(core.StringPtr("testString")))
				Expect(getPermittedNetworkOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(getPermittedNetworkOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListPermittedNetworksOptions successfully`, func() {
				// Construct an instance of the ListPermittedNetworksOptions model
				instanceID := "testString"
				dnszoneID := "testString"
				listPermittedNetworksOptionsModel := testService.NewListPermittedNetworksOptions(instanceID, dnszoneID)
				listPermittedNetworksOptionsModel.SetInstanceID("testString")
				listPermittedNetworksOptionsModel.SetDnszoneID("testString")
				listPermittedNetworksOptionsModel.SetXCorrelationID("testString")
				listPermittedNetworksOptionsModel.SetOffset("testString")
				listPermittedNetworksOptionsModel.SetLimit("testString")
				listPermittedNetworksOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listPermittedNetworksOptionsModel).ToNot(BeNil())
				Expect(listPermittedNetworksOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(listPermittedNetworksOptionsModel.DnszoneID).To(Equal(core.StringPtr("testString")))
				Expect(listPermittedNetworksOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(listPermittedNetworksOptionsModel.Offset).To(Equal(core.StringPtr("testString")))
				Expect(listPermittedNetworksOptionsModel.Limit).To(Equal(core.StringPtr("testString")))
				Expect(listPermittedNetworksOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewPermittedNetworkVpc successfully`, func() {
				vpcCrn := "crn:v1:bluemix:public:is:eu-de:a/bcf1865e99742d38d2d5fc3fb80a5496::vpc:6e6cc326-04d1-4c99-a289-efb3ae4193d6"
				model, err := testService.NewPermittedNetworkVpc(vpcCrn)
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
