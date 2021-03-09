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

package globalloadbalancersv1_test

import (
	"bytes"
	"fmt"
	"github.com/IBM/go-sdk-core/v4/core"
	"github.com/IBM/networking-go-sdk/globalloadbalancersv1"
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

var _ = Describe(`GlobalLoadBalancersV1`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			globalLoadBalancersService, serviceErr := globalloadbalancersv1.NewGlobalLoadBalancersV1(&globalloadbalancersv1.GlobalLoadBalancersV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(globalLoadBalancersService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			globalLoadBalancersService, serviceErr := globalloadbalancersv1.NewGlobalLoadBalancersV1(&globalloadbalancersv1.GlobalLoadBalancersV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(globalLoadBalancersService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			globalLoadBalancersService, serviceErr := globalloadbalancersv1.NewGlobalLoadBalancersV1(&globalloadbalancersv1.GlobalLoadBalancersV1Options{
				URL: "https://globalloadbalancersv1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(globalLoadBalancersService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"GLOBAL_LOAD_BALANCERS_URL": "https://globalloadbalancersv1/api",
				"GLOBAL_LOAD_BALANCERS_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				globalLoadBalancersService, serviceErr := globalloadbalancersv1.NewGlobalLoadBalancersV1UsingExternalConfig(&globalloadbalancersv1.GlobalLoadBalancersV1Options{
				})
				Expect(globalLoadBalancersService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				globalLoadBalancersService, serviceErr := globalloadbalancersv1.NewGlobalLoadBalancersV1UsingExternalConfig(&globalloadbalancersv1.GlobalLoadBalancersV1Options{
					URL: "https://testService/api",
				})
				Expect(globalLoadBalancersService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(globalLoadBalancersService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				globalLoadBalancersService, serviceErr := globalloadbalancersv1.NewGlobalLoadBalancersV1UsingExternalConfig(&globalloadbalancersv1.GlobalLoadBalancersV1Options{
				})
				err := globalLoadBalancersService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(globalLoadBalancersService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(globalLoadBalancersService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"GLOBAL_LOAD_BALANCERS_URL": "https://globalloadbalancersv1/api",
				"GLOBAL_LOAD_BALANCERS_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			globalLoadBalancersService, serviceErr := globalloadbalancersv1.NewGlobalLoadBalancersV1UsingExternalConfig(&globalloadbalancersv1.GlobalLoadBalancersV1Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(globalLoadBalancersService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"GLOBAL_LOAD_BALANCERS_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			globalLoadBalancersService, serviceErr := globalloadbalancersv1.NewGlobalLoadBalancersV1UsingExternalConfig(&globalloadbalancersv1.GlobalLoadBalancersV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(globalLoadBalancersService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`ListLoadBalancers(listLoadBalancersOptions *ListLoadBalancersOptions) - Operation response error`, func() {
		listLoadBalancersPath := "/instances/testString/dnszones/testString/load_balancers"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(listLoadBalancersPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListLoadBalancers with error: Operation response processing error`, func() {
				globalLoadBalancersService, serviceErr := globalloadbalancersv1.NewGlobalLoadBalancersV1(&globalloadbalancersv1.GlobalLoadBalancersV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalLoadBalancersService).ToNot(BeNil())

				// Construct an instance of the ListLoadBalancersOptions model
				listLoadBalancersOptionsModel := new(globalloadbalancersv1.ListLoadBalancersOptions)
				listLoadBalancersOptionsModel.InstanceID = core.StringPtr("testString")
				listLoadBalancersOptionsModel.DnszoneID = core.StringPtr("testString")
				listLoadBalancersOptionsModel.XCorrelationID = core.StringPtr("testString")
				listLoadBalancersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := globalLoadBalancersService.ListLoadBalancers(listLoadBalancersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ListLoadBalancers(listLoadBalancersOptions *ListLoadBalancersOptions)`, func() {
		listLoadBalancersPath := "/instances/testString/dnszones/testString/load_balancers"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(listLoadBalancersPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"load_balancers": [{"id": "5365b73c-ce6f-4d6f-ad9f-d9c131b26370", "name": "glb.example.com", "description": "Load balancer for glb.example.com.", "enabled": true, "ttl": 120, "health": "DEGRADED", "fallback_pool": "24ccf79a-4ae0-4769-b4c8-17f8f230072e", "default_pools": ["DefaultPools"], "az_pools": [{"availability_zone": "us-south-1", "pools": ["0fc0bb7c-2fab-476e-8b9b-40fa14bf8e3d"]}], "created_on": "2019-01-01T05:20:00.12345Z", "modified_on": "2019-01-01T05:20:00.12345Z"}], "offset": 1, "limit": 20, "count": 1, "total_count": 200, "first": {"href": "https://api.dns-svcs.cloud.ibm.com/v1/instances/434f6c3e-6014-4124-a61d-2e910bca19b1/dnszones/example.com:d04d3a7a-7f6d-47d4-b811-08c5478fa1a4/resource_records?limit=20"}, "next": {"href": "https://api.dns-svcs.cloud.ibm.com/v1/instances/434f6c3e-6014-4124-a61d-2e910bca19b1/dnszones/example.com:d04d3a7a-7f6d-47d4-b811-08c5478fa1a4/resource_records?offset=20&limit=20"}}`)
				}))
			})
			It(`Invoke ListLoadBalancers successfully`, func() {
				globalLoadBalancersService, serviceErr := globalloadbalancersv1.NewGlobalLoadBalancersV1(&globalloadbalancersv1.GlobalLoadBalancersV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalLoadBalancersService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := globalLoadBalancersService.ListLoadBalancers(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListLoadBalancersOptions model
				listLoadBalancersOptionsModel := new(globalloadbalancersv1.ListLoadBalancersOptions)
				listLoadBalancersOptionsModel.InstanceID = core.StringPtr("testString")
				listLoadBalancersOptionsModel.DnszoneID = core.StringPtr("testString")
				listLoadBalancersOptionsModel.XCorrelationID = core.StringPtr("testString")
				listLoadBalancersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = globalLoadBalancersService.ListLoadBalancers(listLoadBalancersOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke ListLoadBalancers with error: Operation validation and request error`, func() {
				globalLoadBalancersService, serviceErr := globalloadbalancersv1.NewGlobalLoadBalancersV1(&globalloadbalancersv1.GlobalLoadBalancersV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalLoadBalancersService).ToNot(BeNil())

				// Construct an instance of the ListLoadBalancersOptions model
				listLoadBalancersOptionsModel := new(globalloadbalancersv1.ListLoadBalancersOptions)
				listLoadBalancersOptionsModel.InstanceID = core.StringPtr("testString")
				listLoadBalancersOptionsModel.DnszoneID = core.StringPtr("testString")
				listLoadBalancersOptionsModel.XCorrelationID = core.StringPtr("testString")
				listLoadBalancersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := globalLoadBalancersService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := globalLoadBalancersService.ListLoadBalancers(listLoadBalancersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListLoadBalancersOptions model with no property values
				listLoadBalancersOptionsModelNew := new(globalloadbalancersv1.ListLoadBalancersOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = globalLoadBalancersService.ListLoadBalancers(listLoadBalancersOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateLoadBalancer(createLoadBalancerOptions *CreateLoadBalancerOptions) - Operation response error`, func() {
		createLoadBalancerPath := "/instances/testString/dnszones/testString/load_balancers"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(createLoadBalancerPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateLoadBalancer with error: Operation response processing error`, func() {
				globalLoadBalancersService, serviceErr := globalloadbalancersv1.NewGlobalLoadBalancersV1(&globalloadbalancersv1.GlobalLoadBalancersV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalLoadBalancersService).ToNot(BeNil())

				// Construct an instance of the LoadBalancerAzPoolsItem model
				loadBalancerAzPoolsItemModel := new(globalloadbalancersv1.LoadBalancerAzPoolsItem)
				loadBalancerAzPoolsItemModel.AvailabilityZone = core.StringPtr("us-south-1")
				loadBalancerAzPoolsItemModel.Pools = []string{"0fc0bb7c-2fab-476e-8b9b-40fa14bf8e3d"}

				// Construct an instance of the CreateLoadBalancerOptions model
				createLoadBalancerOptionsModel := new(globalloadbalancersv1.CreateLoadBalancerOptions)
				createLoadBalancerOptionsModel.InstanceID = core.StringPtr("testString")
				createLoadBalancerOptionsModel.DnszoneID = core.StringPtr("testString")
				createLoadBalancerOptionsModel.Name = core.StringPtr("glb.example.com")
				createLoadBalancerOptionsModel.Description = core.StringPtr("Load balancer for glb.example.com.")
				createLoadBalancerOptionsModel.Enabled = core.BoolPtr(true)
				createLoadBalancerOptionsModel.TTL = core.Int64Ptr(int64(120))
				createLoadBalancerOptionsModel.FallbackPool = core.StringPtr("24ccf79a-4ae0-4769-b4c8-17f8f230072e")
				createLoadBalancerOptionsModel.DefaultPools = []string{"testString"}
				createLoadBalancerOptionsModel.AzPools = []globalloadbalancersv1.LoadBalancerAzPoolsItem{*loadBalancerAzPoolsItemModel}
				createLoadBalancerOptionsModel.XCorrelationID = core.StringPtr("testString")
				createLoadBalancerOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := globalLoadBalancersService.CreateLoadBalancer(createLoadBalancerOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`CreateLoadBalancer(createLoadBalancerOptions *CreateLoadBalancerOptions)`, func() {
		createLoadBalancerPath := "/instances/testString/dnszones/testString/load_balancers"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(createLoadBalancerPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "5365b73c-ce6f-4d6f-ad9f-d9c131b26370", "name": "glb.example.com", "description": "Load balancer for glb.example.com.", "enabled": true, "ttl": 120, "health": "DEGRADED", "fallback_pool": "24ccf79a-4ae0-4769-b4c8-17f8f230072e", "default_pools": ["DefaultPools"], "az_pools": [{"availability_zone": "us-south-1", "pools": ["0fc0bb7c-2fab-476e-8b9b-40fa14bf8e3d"]}], "created_on": "2019-01-01T05:20:00.12345Z", "modified_on": "2019-01-01T05:20:00.12345Z"}`)
				}))
			})
			It(`Invoke CreateLoadBalancer successfully`, func() {
				globalLoadBalancersService, serviceErr := globalloadbalancersv1.NewGlobalLoadBalancersV1(&globalloadbalancersv1.GlobalLoadBalancersV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalLoadBalancersService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := globalLoadBalancersService.CreateLoadBalancer(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the LoadBalancerAzPoolsItem model
				loadBalancerAzPoolsItemModel := new(globalloadbalancersv1.LoadBalancerAzPoolsItem)
				loadBalancerAzPoolsItemModel.AvailabilityZone = core.StringPtr("us-south-1")
				loadBalancerAzPoolsItemModel.Pools = []string{"0fc0bb7c-2fab-476e-8b9b-40fa14bf8e3d"}

				// Construct an instance of the CreateLoadBalancerOptions model
				createLoadBalancerOptionsModel := new(globalloadbalancersv1.CreateLoadBalancerOptions)
				createLoadBalancerOptionsModel.InstanceID = core.StringPtr("testString")
				createLoadBalancerOptionsModel.DnszoneID = core.StringPtr("testString")
				createLoadBalancerOptionsModel.Name = core.StringPtr("glb.example.com")
				createLoadBalancerOptionsModel.Description = core.StringPtr("Load balancer for glb.example.com.")
				createLoadBalancerOptionsModel.Enabled = core.BoolPtr(true)
				createLoadBalancerOptionsModel.TTL = core.Int64Ptr(int64(120))
				createLoadBalancerOptionsModel.FallbackPool = core.StringPtr("24ccf79a-4ae0-4769-b4c8-17f8f230072e")
				createLoadBalancerOptionsModel.DefaultPools = []string{"testString"}
				createLoadBalancerOptionsModel.AzPools = []globalloadbalancersv1.LoadBalancerAzPoolsItem{*loadBalancerAzPoolsItemModel}
				createLoadBalancerOptionsModel.XCorrelationID = core.StringPtr("testString")
				createLoadBalancerOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = globalLoadBalancersService.CreateLoadBalancer(createLoadBalancerOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke CreateLoadBalancer with error: Operation validation and request error`, func() {
				globalLoadBalancersService, serviceErr := globalloadbalancersv1.NewGlobalLoadBalancersV1(&globalloadbalancersv1.GlobalLoadBalancersV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalLoadBalancersService).ToNot(BeNil())

				// Construct an instance of the LoadBalancerAzPoolsItem model
				loadBalancerAzPoolsItemModel := new(globalloadbalancersv1.LoadBalancerAzPoolsItem)
				loadBalancerAzPoolsItemModel.AvailabilityZone = core.StringPtr("us-south-1")
				loadBalancerAzPoolsItemModel.Pools = []string{"0fc0bb7c-2fab-476e-8b9b-40fa14bf8e3d"}

				// Construct an instance of the CreateLoadBalancerOptions model
				createLoadBalancerOptionsModel := new(globalloadbalancersv1.CreateLoadBalancerOptions)
				createLoadBalancerOptionsModel.InstanceID = core.StringPtr("testString")
				createLoadBalancerOptionsModel.DnszoneID = core.StringPtr("testString")
				createLoadBalancerOptionsModel.Name = core.StringPtr("glb.example.com")
				createLoadBalancerOptionsModel.Description = core.StringPtr("Load balancer for glb.example.com.")
				createLoadBalancerOptionsModel.Enabled = core.BoolPtr(true)
				createLoadBalancerOptionsModel.TTL = core.Int64Ptr(int64(120))
				createLoadBalancerOptionsModel.FallbackPool = core.StringPtr("24ccf79a-4ae0-4769-b4c8-17f8f230072e")
				createLoadBalancerOptionsModel.DefaultPools = []string{"testString"}
				createLoadBalancerOptionsModel.AzPools = []globalloadbalancersv1.LoadBalancerAzPoolsItem{*loadBalancerAzPoolsItemModel}
				createLoadBalancerOptionsModel.XCorrelationID = core.StringPtr("testString")
				createLoadBalancerOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := globalLoadBalancersService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := globalLoadBalancersService.CreateLoadBalancer(createLoadBalancerOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateLoadBalancerOptions model with no property values
				createLoadBalancerOptionsModelNew := new(globalloadbalancersv1.CreateLoadBalancerOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = globalLoadBalancersService.CreateLoadBalancer(createLoadBalancerOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`DeleteLoadBalancer(deleteLoadBalancerOptions *DeleteLoadBalancerOptions)`, func() {
		deleteLoadBalancerPath := "/instances/testString/dnszones/testString/load_balancers/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(deleteLoadBalancerPath))
					Expect(req.Method).To(Equal("DELETE"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteLoadBalancer successfully`, func() {
				globalLoadBalancersService, serviceErr := globalloadbalancersv1.NewGlobalLoadBalancersV1(&globalloadbalancersv1.GlobalLoadBalancersV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalLoadBalancersService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := globalLoadBalancersService.DeleteLoadBalancer(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteLoadBalancerOptions model
				deleteLoadBalancerOptionsModel := new(globalloadbalancersv1.DeleteLoadBalancerOptions)
				deleteLoadBalancerOptionsModel.InstanceID = core.StringPtr("testString")
				deleteLoadBalancerOptionsModel.DnszoneID = core.StringPtr("testString")
				deleteLoadBalancerOptionsModel.LbID = core.StringPtr("testString")
				deleteLoadBalancerOptionsModel.XCorrelationID = core.StringPtr("testString")
				deleteLoadBalancerOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = globalLoadBalancersService.DeleteLoadBalancer(deleteLoadBalancerOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteLoadBalancer with error: Operation validation and request error`, func() {
				globalLoadBalancersService, serviceErr := globalloadbalancersv1.NewGlobalLoadBalancersV1(&globalloadbalancersv1.GlobalLoadBalancersV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalLoadBalancersService).ToNot(BeNil())

				// Construct an instance of the DeleteLoadBalancerOptions model
				deleteLoadBalancerOptionsModel := new(globalloadbalancersv1.DeleteLoadBalancerOptions)
				deleteLoadBalancerOptionsModel.InstanceID = core.StringPtr("testString")
				deleteLoadBalancerOptionsModel.DnszoneID = core.StringPtr("testString")
				deleteLoadBalancerOptionsModel.LbID = core.StringPtr("testString")
				deleteLoadBalancerOptionsModel.XCorrelationID = core.StringPtr("testString")
				deleteLoadBalancerOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := globalLoadBalancersService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := globalLoadBalancersService.DeleteLoadBalancer(deleteLoadBalancerOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteLoadBalancerOptions model with no property values
				deleteLoadBalancerOptionsModelNew := new(globalloadbalancersv1.DeleteLoadBalancerOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = globalLoadBalancersService.DeleteLoadBalancer(deleteLoadBalancerOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetLoadBalancer(getLoadBalancerOptions *GetLoadBalancerOptions) - Operation response error`, func() {
		getLoadBalancerPath := "/instances/testString/dnszones/testString/load_balancers/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getLoadBalancerPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetLoadBalancer with error: Operation response processing error`, func() {
				globalLoadBalancersService, serviceErr := globalloadbalancersv1.NewGlobalLoadBalancersV1(&globalloadbalancersv1.GlobalLoadBalancersV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalLoadBalancersService).ToNot(BeNil())

				// Construct an instance of the GetLoadBalancerOptions model
				getLoadBalancerOptionsModel := new(globalloadbalancersv1.GetLoadBalancerOptions)
				getLoadBalancerOptionsModel.InstanceID = core.StringPtr("testString")
				getLoadBalancerOptionsModel.DnszoneID = core.StringPtr("testString")
				getLoadBalancerOptionsModel.LbID = core.StringPtr("testString")
				getLoadBalancerOptionsModel.XCorrelationID = core.StringPtr("testString")
				getLoadBalancerOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := globalLoadBalancersService.GetLoadBalancer(getLoadBalancerOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetLoadBalancer(getLoadBalancerOptions *GetLoadBalancerOptions)`, func() {
		getLoadBalancerPath := "/instances/testString/dnszones/testString/load_balancers/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getLoadBalancerPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "5365b73c-ce6f-4d6f-ad9f-d9c131b26370", "name": "glb.example.com", "description": "Load balancer for glb.example.com.", "enabled": true, "ttl": 120, "health": "DEGRADED", "fallback_pool": "24ccf79a-4ae0-4769-b4c8-17f8f230072e", "default_pools": ["DefaultPools"], "az_pools": [{"availability_zone": "us-south-1", "pools": ["0fc0bb7c-2fab-476e-8b9b-40fa14bf8e3d"]}], "created_on": "2019-01-01T05:20:00.12345Z", "modified_on": "2019-01-01T05:20:00.12345Z"}`)
				}))
			})
			It(`Invoke GetLoadBalancer successfully`, func() {
				globalLoadBalancersService, serviceErr := globalloadbalancersv1.NewGlobalLoadBalancersV1(&globalloadbalancersv1.GlobalLoadBalancersV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalLoadBalancersService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := globalLoadBalancersService.GetLoadBalancer(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetLoadBalancerOptions model
				getLoadBalancerOptionsModel := new(globalloadbalancersv1.GetLoadBalancerOptions)
				getLoadBalancerOptionsModel.InstanceID = core.StringPtr("testString")
				getLoadBalancerOptionsModel.DnszoneID = core.StringPtr("testString")
				getLoadBalancerOptionsModel.LbID = core.StringPtr("testString")
				getLoadBalancerOptionsModel.XCorrelationID = core.StringPtr("testString")
				getLoadBalancerOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = globalLoadBalancersService.GetLoadBalancer(getLoadBalancerOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke GetLoadBalancer with error: Operation validation and request error`, func() {
				globalLoadBalancersService, serviceErr := globalloadbalancersv1.NewGlobalLoadBalancersV1(&globalloadbalancersv1.GlobalLoadBalancersV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalLoadBalancersService).ToNot(BeNil())

				// Construct an instance of the GetLoadBalancerOptions model
				getLoadBalancerOptionsModel := new(globalloadbalancersv1.GetLoadBalancerOptions)
				getLoadBalancerOptionsModel.InstanceID = core.StringPtr("testString")
				getLoadBalancerOptionsModel.DnszoneID = core.StringPtr("testString")
				getLoadBalancerOptionsModel.LbID = core.StringPtr("testString")
				getLoadBalancerOptionsModel.XCorrelationID = core.StringPtr("testString")
				getLoadBalancerOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := globalLoadBalancersService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := globalLoadBalancersService.GetLoadBalancer(getLoadBalancerOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetLoadBalancerOptions model with no property values
				getLoadBalancerOptionsModelNew := new(globalloadbalancersv1.GetLoadBalancerOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = globalLoadBalancersService.GetLoadBalancer(getLoadBalancerOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateLoadBalancer(updateLoadBalancerOptions *UpdateLoadBalancerOptions) - Operation response error`, func() {
		updateLoadBalancerPath := "/instances/testString/dnszones/testString/load_balancers/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(updateLoadBalancerPath))
					Expect(req.Method).To(Equal("PUT"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateLoadBalancer with error: Operation response processing error`, func() {
				globalLoadBalancersService, serviceErr := globalloadbalancersv1.NewGlobalLoadBalancersV1(&globalloadbalancersv1.GlobalLoadBalancersV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalLoadBalancersService).ToNot(BeNil())

				// Construct an instance of the LoadBalancerAzPoolsItem model
				loadBalancerAzPoolsItemModel := new(globalloadbalancersv1.LoadBalancerAzPoolsItem)
				loadBalancerAzPoolsItemModel.AvailabilityZone = core.StringPtr("us-south-1")
				loadBalancerAzPoolsItemModel.Pools = []string{"0fc0bb7c-2fab-476e-8b9b-40fa14bf8e3d"}

				// Construct an instance of the UpdateLoadBalancerOptions model
				updateLoadBalancerOptionsModel := new(globalloadbalancersv1.UpdateLoadBalancerOptions)
				updateLoadBalancerOptionsModel.InstanceID = core.StringPtr("testString")
				updateLoadBalancerOptionsModel.DnszoneID = core.StringPtr("testString")
				updateLoadBalancerOptionsModel.LbID = core.StringPtr("testString")
				updateLoadBalancerOptionsModel.Name = core.StringPtr("glb.example.com")
				updateLoadBalancerOptionsModel.Description = core.StringPtr("Load balancer for glb.example.com.")
				updateLoadBalancerOptionsModel.Enabled = core.BoolPtr(true)
				updateLoadBalancerOptionsModel.TTL = core.Int64Ptr(int64(120))
				updateLoadBalancerOptionsModel.FallbackPool = core.StringPtr("24ccf79a-4ae0-4769-b4c8-17f8f230072e")
				updateLoadBalancerOptionsModel.DefaultPools = []string{"testString"}
				updateLoadBalancerOptionsModel.AzPools = []globalloadbalancersv1.LoadBalancerAzPoolsItem{*loadBalancerAzPoolsItemModel}
				updateLoadBalancerOptionsModel.XCorrelationID = core.StringPtr("testString")
				updateLoadBalancerOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := globalLoadBalancersService.UpdateLoadBalancer(updateLoadBalancerOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UpdateLoadBalancer(updateLoadBalancerOptions *UpdateLoadBalancerOptions)`, func() {
		updateLoadBalancerPath := "/instances/testString/dnszones/testString/load_balancers/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(updateLoadBalancerPath))
					Expect(req.Method).To(Equal("PUT"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "5365b73c-ce6f-4d6f-ad9f-d9c131b26370", "name": "glb.example.com", "description": "Load balancer for glb.example.com.", "enabled": true, "ttl": 120, "health": "DEGRADED", "fallback_pool": "24ccf79a-4ae0-4769-b4c8-17f8f230072e", "default_pools": ["DefaultPools"], "az_pools": [{"availability_zone": "us-south-1", "pools": ["0fc0bb7c-2fab-476e-8b9b-40fa14bf8e3d"]}], "created_on": "2019-01-01T05:20:00.12345Z", "modified_on": "2019-01-01T05:20:00.12345Z"}`)
				}))
			})
			It(`Invoke UpdateLoadBalancer successfully`, func() {
				globalLoadBalancersService, serviceErr := globalloadbalancersv1.NewGlobalLoadBalancersV1(&globalloadbalancersv1.GlobalLoadBalancersV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalLoadBalancersService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := globalLoadBalancersService.UpdateLoadBalancer(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the LoadBalancerAzPoolsItem model
				loadBalancerAzPoolsItemModel := new(globalloadbalancersv1.LoadBalancerAzPoolsItem)
				loadBalancerAzPoolsItemModel.AvailabilityZone = core.StringPtr("us-south-1")
				loadBalancerAzPoolsItemModel.Pools = []string{"0fc0bb7c-2fab-476e-8b9b-40fa14bf8e3d"}

				// Construct an instance of the UpdateLoadBalancerOptions model
				updateLoadBalancerOptionsModel := new(globalloadbalancersv1.UpdateLoadBalancerOptions)
				updateLoadBalancerOptionsModel.InstanceID = core.StringPtr("testString")
				updateLoadBalancerOptionsModel.DnszoneID = core.StringPtr("testString")
				updateLoadBalancerOptionsModel.LbID = core.StringPtr("testString")
				updateLoadBalancerOptionsModel.Name = core.StringPtr("glb.example.com")
				updateLoadBalancerOptionsModel.Description = core.StringPtr("Load balancer for glb.example.com.")
				updateLoadBalancerOptionsModel.Enabled = core.BoolPtr(true)
				updateLoadBalancerOptionsModel.TTL = core.Int64Ptr(int64(120))
				updateLoadBalancerOptionsModel.FallbackPool = core.StringPtr("24ccf79a-4ae0-4769-b4c8-17f8f230072e")
				updateLoadBalancerOptionsModel.DefaultPools = []string{"testString"}
				updateLoadBalancerOptionsModel.AzPools = []globalloadbalancersv1.LoadBalancerAzPoolsItem{*loadBalancerAzPoolsItemModel}
				updateLoadBalancerOptionsModel.XCorrelationID = core.StringPtr("testString")
				updateLoadBalancerOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = globalLoadBalancersService.UpdateLoadBalancer(updateLoadBalancerOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke UpdateLoadBalancer with error: Operation validation and request error`, func() {
				globalLoadBalancersService, serviceErr := globalloadbalancersv1.NewGlobalLoadBalancersV1(&globalloadbalancersv1.GlobalLoadBalancersV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalLoadBalancersService).ToNot(BeNil())

				// Construct an instance of the LoadBalancerAzPoolsItem model
				loadBalancerAzPoolsItemModel := new(globalloadbalancersv1.LoadBalancerAzPoolsItem)
				loadBalancerAzPoolsItemModel.AvailabilityZone = core.StringPtr("us-south-1")
				loadBalancerAzPoolsItemModel.Pools = []string{"0fc0bb7c-2fab-476e-8b9b-40fa14bf8e3d"}

				// Construct an instance of the UpdateLoadBalancerOptions model
				updateLoadBalancerOptionsModel := new(globalloadbalancersv1.UpdateLoadBalancerOptions)
				updateLoadBalancerOptionsModel.InstanceID = core.StringPtr("testString")
				updateLoadBalancerOptionsModel.DnszoneID = core.StringPtr("testString")
				updateLoadBalancerOptionsModel.LbID = core.StringPtr("testString")
				updateLoadBalancerOptionsModel.Name = core.StringPtr("glb.example.com")
				updateLoadBalancerOptionsModel.Description = core.StringPtr("Load balancer for glb.example.com.")
				updateLoadBalancerOptionsModel.Enabled = core.BoolPtr(true)
				updateLoadBalancerOptionsModel.TTL = core.Int64Ptr(int64(120))
				updateLoadBalancerOptionsModel.FallbackPool = core.StringPtr("24ccf79a-4ae0-4769-b4c8-17f8f230072e")
				updateLoadBalancerOptionsModel.DefaultPools = []string{"testString"}
				updateLoadBalancerOptionsModel.AzPools = []globalloadbalancersv1.LoadBalancerAzPoolsItem{*loadBalancerAzPoolsItemModel}
				updateLoadBalancerOptionsModel.XCorrelationID = core.StringPtr("testString")
				updateLoadBalancerOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := globalLoadBalancersService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := globalLoadBalancersService.UpdateLoadBalancer(updateLoadBalancerOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateLoadBalancerOptions model with no property values
				updateLoadBalancerOptionsModelNew := new(globalloadbalancersv1.UpdateLoadBalancerOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = globalLoadBalancersService.UpdateLoadBalancer(updateLoadBalancerOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			globalLoadBalancersService, serviceErr := globalloadbalancersv1.NewGlobalLoadBalancersV1(&globalloadbalancersv1.GlobalLoadBalancersV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(globalLoadBalancersService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			globalLoadBalancersService, serviceErr := globalloadbalancersv1.NewGlobalLoadBalancersV1(&globalloadbalancersv1.GlobalLoadBalancersV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(globalLoadBalancersService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			globalLoadBalancersService, serviceErr := globalloadbalancersv1.NewGlobalLoadBalancersV1(&globalloadbalancersv1.GlobalLoadBalancersV1Options{
				URL: "https://globalloadbalancersv1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(globalLoadBalancersService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"GLOBAL_LOAD_BALANCERS_URL": "https://globalloadbalancersv1/api",
				"GLOBAL_LOAD_BALANCERS_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				globalLoadBalancersService, serviceErr := globalloadbalancersv1.NewGlobalLoadBalancersV1UsingExternalConfig(&globalloadbalancersv1.GlobalLoadBalancersV1Options{
				})
				Expect(globalLoadBalancersService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				globalLoadBalancersService, serviceErr := globalloadbalancersv1.NewGlobalLoadBalancersV1UsingExternalConfig(&globalloadbalancersv1.GlobalLoadBalancersV1Options{
					URL: "https://testService/api",
				})
				Expect(globalLoadBalancersService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(globalLoadBalancersService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				globalLoadBalancersService, serviceErr := globalloadbalancersv1.NewGlobalLoadBalancersV1UsingExternalConfig(&globalloadbalancersv1.GlobalLoadBalancersV1Options{
				})
				err := globalLoadBalancersService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(globalLoadBalancersService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(globalLoadBalancersService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"GLOBAL_LOAD_BALANCERS_URL": "https://globalloadbalancersv1/api",
				"GLOBAL_LOAD_BALANCERS_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			globalLoadBalancersService, serviceErr := globalloadbalancersv1.NewGlobalLoadBalancersV1UsingExternalConfig(&globalloadbalancersv1.GlobalLoadBalancersV1Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(globalLoadBalancersService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"GLOBAL_LOAD_BALANCERS_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			globalLoadBalancersService, serviceErr := globalloadbalancersv1.NewGlobalLoadBalancersV1UsingExternalConfig(&globalloadbalancersv1.GlobalLoadBalancersV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(globalLoadBalancersService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`ListPools(listPoolsOptions *ListPoolsOptions) - Operation response error`, func() {
		listPoolsPath := "/instances/testString/pools"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(listPoolsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListPools with error: Operation response processing error`, func() {
				globalLoadBalancersService, serviceErr := globalloadbalancersv1.NewGlobalLoadBalancersV1(&globalloadbalancersv1.GlobalLoadBalancersV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalLoadBalancersService).ToNot(BeNil())

				// Construct an instance of the ListPoolsOptions model
				listPoolsOptionsModel := new(globalloadbalancersv1.ListPoolsOptions)
				listPoolsOptionsModel.InstanceID = core.StringPtr("testString")
				listPoolsOptionsModel.XCorrelationID = core.StringPtr("testString")
				listPoolsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := globalLoadBalancersService.ListPools(listPoolsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ListPools(listPoolsOptions *ListPoolsOptions)`, func() {
		listPoolsPath := "/instances/testString/pools"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(listPoolsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"pools": [{"id": "5365b73c-ce6f-4d6f-ad9f-d9c131b26370", "name": "dal10-az-pool", "description": "Load balancer pool for dal10 availability zone.", "enabled": true, "healthy_origins_threshold": 1, "origins": [{"name": "app-server-1", "description": "description of the origin server", "address": "10.10.16.8", "enabled": true, "health": true, "health_failure_reason": "HealthFailureReason"}], "monitor": "7dd6841c-264e-11ea-88df-062967242a6a", "notification_channel": "https://mywebsite.com/dns/webhook", "health": "HEALTHY", "healthcheck_region": "us-south", "healthcheck_subnets": ["0716-a4c0c123-594c-4ef4-ace3-a08858540b5e"], "created_on": "2019-01-01T05:20:00.12345Z", "modified_on": "2019-01-01T05:20:00.12345Z"}], "offset": 1, "limit": 20, "count": 1, "total_count": 200, "first": {"href": "https://api.dns-svcs.cloud.ibm.com/v1/instances/434f6c3e-6014-4124-a61d-2e910bca19b1/dnszones/example.com:d04d3a7a-7f6d-47d4-b811-08c5478fa1a4/resource_records?limit=20"}, "next": {"href": "https://api.dns-svcs.cloud.ibm.com/v1/instances/434f6c3e-6014-4124-a61d-2e910bca19b1/dnszones/example.com:d04d3a7a-7f6d-47d4-b811-08c5478fa1a4/resource_records?offset=20&limit=20"}}`)
				}))
			})
			It(`Invoke ListPools successfully`, func() {
				globalLoadBalancersService, serviceErr := globalloadbalancersv1.NewGlobalLoadBalancersV1(&globalloadbalancersv1.GlobalLoadBalancersV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalLoadBalancersService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := globalLoadBalancersService.ListPools(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListPoolsOptions model
				listPoolsOptionsModel := new(globalloadbalancersv1.ListPoolsOptions)
				listPoolsOptionsModel.InstanceID = core.StringPtr("testString")
				listPoolsOptionsModel.XCorrelationID = core.StringPtr("testString")
				listPoolsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = globalLoadBalancersService.ListPools(listPoolsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke ListPools with error: Operation validation and request error`, func() {
				globalLoadBalancersService, serviceErr := globalloadbalancersv1.NewGlobalLoadBalancersV1(&globalloadbalancersv1.GlobalLoadBalancersV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalLoadBalancersService).ToNot(BeNil())

				// Construct an instance of the ListPoolsOptions model
				listPoolsOptionsModel := new(globalloadbalancersv1.ListPoolsOptions)
				listPoolsOptionsModel.InstanceID = core.StringPtr("testString")
				listPoolsOptionsModel.XCorrelationID = core.StringPtr("testString")
				listPoolsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := globalLoadBalancersService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := globalLoadBalancersService.ListPools(listPoolsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListPoolsOptions model with no property values
				listPoolsOptionsModelNew := new(globalloadbalancersv1.ListPoolsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = globalLoadBalancersService.ListPools(listPoolsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreatePool(createPoolOptions *CreatePoolOptions) - Operation response error`, func() {
		createPoolPath := "/instances/testString/pools"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(createPoolPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreatePool with error: Operation response processing error`, func() {
				globalLoadBalancersService, serviceErr := globalloadbalancersv1.NewGlobalLoadBalancersV1(&globalloadbalancersv1.GlobalLoadBalancersV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalLoadBalancersService).ToNot(BeNil())

				// Construct an instance of the OriginInput model
				originInputModel := new(globalloadbalancersv1.OriginInput)
				originInputModel.Name = core.StringPtr("app-server-1")
				originInputModel.Description = core.StringPtr("description of the origin server")
				originInputModel.Address = core.StringPtr("10.10.16.8")
				originInputModel.Enabled = core.BoolPtr(true)

				// Construct an instance of the CreatePoolOptions model
				createPoolOptionsModel := new(globalloadbalancersv1.CreatePoolOptions)
				createPoolOptionsModel.InstanceID = core.StringPtr("testString")
				createPoolOptionsModel.Name = core.StringPtr("dal10-az-pool")
				createPoolOptionsModel.Description = core.StringPtr("Load balancer pool for dal10 availability zone.")
				createPoolOptionsModel.Enabled = core.BoolPtr(true)
				createPoolOptionsModel.HealthyOriginsThreshold = core.Int64Ptr(int64(1))
				createPoolOptionsModel.Origins = []globalloadbalancersv1.OriginInput{*originInputModel}
				createPoolOptionsModel.Monitor = core.StringPtr("7dd6841c-264e-11ea-88df-062967242a6a")
				createPoolOptionsModel.NotificationChannel = core.StringPtr("https://mywebsite.com/dns/webhook")
				createPoolOptionsModel.HealthcheckRegion = core.StringPtr("us-south")
				createPoolOptionsModel.HealthcheckSubnets = []string{"0716-a4c0c123-594c-4ef4-ace3-a08858540b5e"}
				createPoolOptionsModel.XCorrelationID = core.StringPtr("testString")
				createPoolOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := globalLoadBalancersService.CreatePool(createPoolOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`CreatePool(createPoolOptions *CreatePoolOptions)`, func() {
		createPoolPath := "/instances/testString/pools"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(createPoolPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "5365b73c-ce6f-4d6f-ad9f-d9c131b26370", "name": "dal10-az-pool", "description": "Load balancer pool for dal10 availability zone.", "enabled": true, "healthy_origins_threshold": 1, "origins": [{"name": "app-server-1", "description": "description of the origin server", "address": "10.10.16.8", "enabled": true, "health": true, "health_failure_reason": "HealthFailureReason"}], "monitor": "7dd6841c-264e-11ea-88df-062967242a6a", "notification_channel": "https://mywebsite.com/dns/webhook", "health": "HEALTHY", "healthcheck_region": "us-south", "healthcheck_subnets": ["0716-a4c0c123-594c-4ef4-ace3-a08858540b5e"], "created_on": "2019-01-01T05:20:00.12345Z", "modified_on": "2019-01-01T05:20:00.12345Z"}`)
				}))
			})
			It(`Invoke CreatePool successfully`, func() {
				globalLoadBalancersService, serviceErr := globalloadbalancersv1.NewGlobalLoadBalancersV1(&globalloadbalancersv1.GlobalLoadBalancersV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalLoadBalancersService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := globalLoadBalancersService.CreatePool(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the OriginInput model
				originInputModel := new(globalloadbalancersv1.OriginInput)
				originInputModel.Name = core.StringPtr("app-server-1")
				originInputModel.Description = core.StringPtr("description of the origin server")
				originInputModel.Address = core.StringPtr("10.10.16.8")
				originInputModel.Enabled = core.BoolPtr(true)

				// Construct an instance of the CreatePoolOptions model
				createPoolOptionsModel := new(globalloadbalancersv1.CreatePoolOptions)
				createPoolOptionsModel.InstanceID = core.StringPtr("testString")
				createPoolOptionsModel.Name = core.StringPtr("dal10-az-pool")
				createPoolOptionsModel.Description = core.StringPtr("Load balancer pool for dal10 availability zone.")
				createPoolOptionsModel.Enabled = core.BoolPtr(true)
				createPoolOptionsModel.HealthyOriginsThreshold = core.Int64Ptr(int64(1))
				createPoolOptionsModel.Origins = []globalloadbalancersv1.OriginInput{*originInputModel}
				createPoolOptionsModel.Monitor = core.StringPtr("7dd6841c-264e-11ea-88df-062967242a6a")
				createPoolOptionsModel.NotificationChannel = core.StringPtr("https://mywebsite.com/dns/webhook")
				createPoolOptionsModel.HealthcheckRegion = core.StringPtr("us-south")
				createPoolOptionsModel.HealthcheckSubnets = []string{"0716-a4c0c123-594c-4ef4-ace3-a08858540b5e"}
				createPoolOptionsModel.XCorrelationID = core.StringPtr("testString")
				createPoolOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = globalLoadBalancersService.CreatePool(createPoolOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke CreatePool with error: Operation validation and request error`, func() {
				globalLoadBalancersService, serviceErr := globalloadbalancersv1.NewGlobalLoadBalancersV1(&globalloadbalancersv1.GlobalLoadBalancersV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalLoadBalancersService).ToNot(BeNil())

				// Construct an instance of the OriginInput model
				originInputModel := new(globalloadbalancersv1.OriginInput)
				originInputModel.Name = core.StringPtr("app-server-1")
				originInputModel.Description = core.StringPtr("description of the origin server")
				originInputModel.Address = core.StringPtr("10.10.16.8")
				originInputModel.Enabled = core.BoolPtr(true)

				// Construct an instance of the CreatePoolOptions model
				createPoolOptionsModel := new(globalloadbalancersv1.CreatePoolOptions)
				createPoolOptionsModel.InstanceID = core.StringPtr("testString")
				createPoolOptionsModel.Name = core.StringPtr("dal10-az-pool")
				createPoolOptionsModel.Description = core.StringPtr("Load balancer pool for dal10 availability zone.")
				createPoolOptionsModel.Enabled = core.BoolPtr(true)
				createPoolOptionsModel.HealthyOriginsThreshold = core.Int64Ptr(int64(1))
				createPoolOptionsModel.Origins = []globalloadbalancersv1.OriginInput{*originInputModel}
				createPoolOptionsModel.Monitor = core.StringPtr("7dd6841c-264e-11ea-88df-062967242a6a")
				createPoolOptionsModel.NotificationChannel = core.StringPtr("https://mywebsite.com/dns/webhook")
				createPoolOptionsModel.HealthcheckRegion = core.StringPtr("us-south")
				createPoolOptionsModel.HealthcheckSubnets = []string{"0716-a4c0c123-594c-4ef4-ace3-a08858540b5e"}
				createPoolOptionsModel.XCorrelationID = core.StringPtr("testString")
				createPoolOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := globalLoadBalancersService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := globalLoadBalancersService.CreatePool(createPoolOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreatePoolOptions model with no property values
				createPoolOptionsModelNew := new(globalloadbalancersv1.CreatePoolOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = globalLoadBalancersService.CreatePool(createPoolOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`DeletePool(deletePoolOptions *DeletePoolOptions)`, func() {
		deletePoolPath := "/instances/testString/pools/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(deletePoolPath))
					Expect(req.Method).To(Equal("DELETE"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeletePool successfully`, func() {
				globalLoadBalancersService, serviceErr := globalloadbalancersv1.NewGlobalLoadBalancersV1(&globalloadbalancersv1.GlobalLoadBalancersV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalLoadBalancersService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := globalLoadBalancersService.DeletePool(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeletePoolOptions model
				deletePoolOptionsModel := new(globalloadbalancersv1.DeletePoolOptions)
				deletePoolOptionsModel.InstanceID = core.StringPtr("testString")
				deletePoolOptionsModel.PoolID = core.StringPtr("testString")
				deletePoolOptionsModel.XCorrelationID = core.StringPtr("testString")
				deletePoolOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = globalLoadBalancersService.DeletePool(deletePoolOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeletePool with error: Operation validation and request error`, func() {
				globalLoadBalancersService, serviceErr := globalloadbalancersv1.NewGlobalLoadBalancersV1(&globalloadbalancersv1.GlobalLoadBalancersV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalLoadBalancersService).ToNot(BeNil())

				// Construct an instance of the DeletePoolOptions model
				deletePoolOptionsModel := new(globalloadbalancersv1.DeletePoolOptions)
				deletePoolOptionsModel.InstanceID = core.StringPtr("testString")
				deletePoolOptionsModel.PoolID = core.StringPtr("testString")
				deletePoolOptionsModel.XCorrelationID = core.StringPtr("testString")
				deletePoolOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := globalLoadBalancersService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := globalLoadBalancersService.DeletePool(deletePoolOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeletePoolOptions model with no property values
				deletePoolOptionsModelNew := new(globalloadbalancersv1.DeletePoolOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = globalLoadBalancersService.DeletePool(deletePoolOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetPool(getPoolOptions *GetPoolOptions) - Operation response error`, func() {
		getPoolPath := "/instances/testString/pools/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getPoolPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetPool with error: Operation response processing error`, func() {
				globalLoadBalancersService, serviceErr := globalloadbalancersv1.NewGlobalLoadBalancersV1(&globalloadbalancersv1.GlobalLoadBalancersV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalLoadBalancersService).ToNot(BeNil())

				// Construct an instance of the GetPoolOptions model
				getPoolOptionsModel := new(globalloadbalancersv1.GetPoolOptions)
				getPoolOptionsModel.InstanceID = core.StringPtr("testString")
				getPoolOptionsModel.PoolID = core.StringPtr("testString")
				getPoolOptionsModel.XCorrelationID = core.StringPtr("testString")
				getPoolOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := globalLoadBalancersService.GetPool(getPoolOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetPool(getPoolOptions *GetPoolOptions)`, func() {
		getPoolPath := "/instances/testString/pools/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getPoolPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "5365b73c-ce6f-4d6f-ad9f-d9c131b26370", "name": "dal10-az-pool", "description": "Load balancer pool for dal10 availability zone.", "enabled": true, "healthy_origins_threshold": 1, "origins": [{"name": "app-server-1", "description": "description of the origin server", "address": "10.10.16.8", "enabled": true, "health": true, "health_failure_reason": "HealthFailureReason"}], "monitor": "7dd6841c-264e-11ea-88df-062967242a6a", "notification_channel": "https://mywebsite.com/dns/webhook", "health": "HEALTHY", "healthcheck_region": "us-south", "healthcheck_subnets": ["0716-a4c0c123-594c-4ef4-ace3-a08858540b5e"], "created_on": "2019-01-01T05:20:00.12345Z", "modified_on": "2019-01-01T05:20:00.12345Z"}`)
				}))
			})
			It(`Invoke GetPool successfully`, func() {
				globalLoadBalancersService, serviceErr := globalloadbalancersv1.NewGlobalLoadBalancersV1(&globalloadbalancersv1.GlobalLoadBalancersV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalLoadBalancersService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := globalLoadBalancersService.GetPool(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetPoolOptions model
				getPoolOptionsModel := new(globalloadbalancersv1.GetPoolOptions)
				getPoolOptionsModel.InstanceID = core.StringPtr("testString")
				getPoolOptionsModel.PoolID = core.StringPtr("testString")
				getPoolOptionsModel.XCorrelationID = core.StringPtr("testString")
				getPoolOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = globalLoadBalancersService.GetPool(getPoolOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke GetPool with error: Operation validation and request error`, func() {
				globalLoadBalancersService, serviceErr := globalloadbalancersv1.NewGlobalLoadBalancersV1(&globalloadbalancersv1.GlobalLoadBalancersV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalLoadBalancersService).ToNot(BeNil())

				// Construct an instance of the GetPoolOptions model
				getPoolOptionsModel := new(globalloadbalancersv1.GetPoolOptions)
				getPoolOptionsModel.InstanceID = core.StringPtr("testString")
				getPoolOptionsModel.PoolID = core.StringPtr("testString")
				getPoolOptionsModel.XCorrelationID = core.StringPtr("testString")
				getPoolOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := globalLoadBalancersService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := globalLoadBalancersService.GetPool(getPoolOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetPoolOptions model with no property values
				getPoolOptionsModelNew := new(globalloadbalancersv1.GetPoolOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = globalLoadBalancersService.GetPool(getPoolOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdatePool(updatePoolOptions *UpdatePoolOptions) - Operation response error`, func() {
		updatePoolPath := "/instances/testString/pools/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(updatePoolPath))
					Expect(req.Method).To(Equal("PUT"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdatePool with error: Operation response processing error`, func() {
				globalLoadBalancersService, serviceErr := globalloadbalancersv1.NewGlobalLoadBalancersV1(&globalloadbalancersv1.GlobalLoadBalancersV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalLoadBalancersService).ToNot(BeNil())

				// Construct an instance of the OriginInput model
				originInputModel := new(globalloadbalancersv1.OriginInput)
				originInputModel.Name = core.StringPtr("app-server-1")
				originInputModel.Description = core.StringPtr("description of the origin server")
				originInputModel.Address = core.StringPtr("10.10.16.8")
				originInputModel.Enabled = core.BoolPtr(true)

				// Construct an instance of the UpdatePoolOptions model
				updatePoolOptionsModel := new(globalloadbalancersv1.UpdatePoolOptions)
				updatePoolOptionsModel.InstanceID = core.StringPtr("testString")
				updatePoolOptionsModel.PoolID = core.StringPtr("testString")
				updatePoolOptionsModel.Name = core.StringPtr("dal10-az-pool")
				updatePoolOptionsModel.Description = core.StringPtr("Load balancer pool for dal10 availability zone.")
				updatePoolOptionsModel.Enabled = core.BoolPtr(true)
				updatePoolOptionsModel.HealthyOriginsThreshold = core.Int64Ptr(int64(1))
				updatePoolOptionsModel.Origins = []globalloadbalancersv1.OriginInput{*originInputModel}
				updatePoolOptionsModel.Monitor = core.StringPtr("7dd6841c-264e-11ea-88df-062967242a6a")
				updatePoolOptionsModel.NotificationChannel = core.StringPtr("https://mywebsite.com/dns/webhook")
				updatePoolOptionsModel.HealthcheckRegion = core.StringPtr("us-south")
				updatePoolOptionsModel.HealthcheckSubnets = []string{"0716-a4c0c123-594c-4ef4-ace3-a08858540b5e"}
				updatePoolOptionsModel.XCorrelationID = core.StringPtr("testString")
				updatePoolOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := globalLoadBalancersService.UpdatePool(updatePoolOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UpdatePool(updatePoolOptions *UpdatePoolOptions)`, func() {
		updatePoolPath := "/instances/testString/pools/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(updatePoolPath))
					Expect(req.Method).To(Equal("PUT"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "5365b73c-ce6f-4d6f-ad9f-d9c131b26370", "name": "dal10-az-pool", "description": "Load balancer pool for dal10 availability zone.", "enabled": true, "healthy_origins_threshold": 1, "origins": [{"name": "app-server-1", "description": "description of the origin server", "address": "10.10.16.8", "enabled": true, "health": true, "health_failure_reason": "HealthFailureReason"}], "monitor": "7dd6841c-264e-11ea-88df-062967242a6a", "notification_channel": "https://mywebsite.com/dns/webhook", "health": "HEALTHY", "healthcheck_region": "us-south", "healthcheck_subnets": ["0716-a4c0c123-594c-4ef4-ace3-a08858540b5e"], "created_on": "2019-01-01T05:20:00.12345Z", "modified_on": "2019-01-01T05:20:00.12345Z"}`)
				}))
			})
			It(`Invoke UpdatePool successfully`, func() {
				globalLoadBalancersService, serviceErr := globalloadbalancersv1.NewGlobalLoadBalancersV1(&globalloadbalancersv1.GlobalLoadBalancersV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalLoadBalancersService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := globalLoadBalancersService.UpdatePool(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the OriginInput model
				originInputModel := new(globalloadbalancersv1.OriginInput)
				originInputModel.Name = core.StringPtr("app-server-1")
				originInputModel.Description = core.StringPtr("description of the origin server")
				originInputModel.Address = core.StringPtr("10.10.16.8")
				originInputModel.Enabled = core.BoolPtr(true)

				// Construct an instance of the UpdatePoolOptions model
				updatePoolOptionsModel := new(globalloadbalancersv1.UpdatePoolOptions)
				updatePoolOptionsModel.InstanceID = core.StringPtr("testString")
				updatePoolOptionsModel.PoolID = core.StringPtr("testString")
				updatePoolOptionsModel.Name = core.StringPtr("dal10-az-pool")
				updatePoolOptionsModel.Description = core.StringPtr("Load balancer pool for dal10 availability zone.")
				updatePoolOptionsModel.Enabled = core.BoolPtr(true)
				updatePoolOptionsModel.HealthyOriginsThreshold = core.Int64Ptr(int64(1))
				updatePoolOptionsModel.Origins = []globalloadbalancersv1.OriginInput{*originInputModel}
				updatePoolOptionsModel.Monitor = core.StringPtr("7dd6841c-264e-11ea-88df-062967242a6a")
				updatePoolOptionsModel.NotificationChannel = core.StringPtr("https://mywebsite.com/dns/webhook")
				updatePoolOptionsModel.HealthcheckRegion = core.StringPtr("us-south")
				updatePoolOptionsModel.HealthcheckSubnets = []string{"0716-a4c0c123-594c-4ef4-ace3-a08858540b5e"}
				updatePoolOptionsModel.XCorrelationID = core.StringPtr("testString")
				updatePoolOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = globalLoadBalancersService.UpdatePool(updatePoolOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke UpdatePool with error: Operation validation and request error`, func() {
				globalLoadBalancersService, serviceErr := globalloadbalancersv1.NewGlobalLoadBalancersV1(&globalloadbalancersv1.GlobalLoadBalancersV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalLoadBalancersService).ToNot(BeNil())

				// Construct an instance of the OriginInput model
				originInputModel := new(globalloadbalancersv1.OriginInput)
				originInputModel.Name = core.StringPtr("app-server-1")
				originInputModel.Description = core.StringPtr("description of the origin server")
				originInputModel.Address = core.StringPtr("10.10.16.8")
				originInputModel.Enabled = core.BoolPtr(true)

				// Construct an instance of the UpdatePoolOptions model
				updatePoolOptionsModel := new(globalloadbalancersv1.UpdatePoolOptions)
				updatePoolOptionsModel.InstanceID = core.StringPtr("testString")
				updatePoolOptionsModel.PoolID = core.StringPtr("testString")
				updatePoolOptionsModel.Name = core.StringPtr("dal10-az-pool")
				updatePoolOptionsModel.Description = core.StringPtr("Load balancer pool for dal10 availability zone.")
				updatePoolOptionsModel.Enabled = core.BoolPtr(true)
				updatePoolOptionsModel.HealthyOriginsThreshold = core.Int64Ptr(int64(1))
				updatePoolOptionsModel.Origins = []globalloadbalancersv1.OriginInput{*originInputModel}
				updatePoolOptionsModel.Monitor = core.StringPtr("7dd6841c-264e-11ea-88df-062967242a6a")
				updatePoolOptionsModel.NotificationChannel = core.StringPtr("https://mywebsite.com/dns/webhook")
				updatePoolOptionsModel.HealthcheckRegion = core.StringPtr("us-south")
				updatePoolOptionsModel.HealthcheckSubnets = []string{"0716-a4c0c123-594c-4ef4-ace3-a08858540b5e"}
				updatePoolOptionsModel.XCorrelationID = core.StringPtr("testString")
				updatePoolOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := globalLoadBalancersService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := globalLoadBalancersService.UpdatePool(updatePoolOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdatePoolOptions model with no property values
				updatePoolOptionsModelNew := new(globalloadbalancersv1.UpdatePoolOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = globalLoadBalancersService.UpdatePool(updatePoolOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			globalLoadBalancersService, serviceErr := globalloadbalancersv1.NewGlobalLoadBalancersV1(&globalloadbalancersv1.GlobalLoadBalancersV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(globalLoadBalancersService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			globalLoadBalancersService, serviceErr := globalloadbalancersv1.NewGlobalLoadBalancersV1(&globalloadbalancersv1.GlobalLoadBalancersV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(globalLoadBalancersService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			globalLoadBalancersService, serviceErr := globalloadbalancersv1.NewGlobalLoadBalancersV1(&globalloadbalancersv1.GlobalLoadBalancersV1Options{
				URL: "https://globalloadbalancersv1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(globalLoadBalancersService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"GLOBAL_LOAD_BALANCERS_URL": "https://globalloadbalancersv1/api",
				"GLOBAL_LOAD_BALANCERS_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				globalLoadBalancersService, serviceErr := globalloadbalancersv1.NewGlobalLoadBalancersV1UsingExternalConfig(&globalloadbalancersv1.GlobalLoadBalancersV1Options{
				})
				Expect(globalLoadBalancersService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				globalLoadBalancersService, serviceErr := globalloadbalancersv1.NewGlobalLoadBalancersV1UsingExternalConfig(&globalloadbalancersv1.GlobalLoadBalancersV1Options{
					URL: "https://testService/api",
				})
				Expect(globalLoadBalancersService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(globalLoadBalancersService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				globalLoadBalancersService, serviceErr := globalloadbalancersv1.NewGlobalLoadBalancersV1UsingExternalConfig(&globalloadbalancersv1.GlobalLoadBalancersV1Options{
				})
				err := globalLoadBalancersService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(globalLoadBalancersService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(globalLoadBalancersService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"GLOBAL_LOAD_BALANCERS_URL": "https://globalloadbalancersv1/api",
				"GLOBAL_LOAD_BALANCERS_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			globalLoadBalancersService, serviceErr := globalloadbalancersv1.NewGlobalLoadBalancersV1UsingExternalConfig(&globalloadbalancersv1.GlobalLoadBalancersV1Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(globalLoadBalancersService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"GLOBAL_LOAD_BALANCERS_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			globalLoadBalancersService, serviceErr := globalloadbalancersv1.NewGlobalLoadBalancersV1UsingExternalConfig(&globalloadbalancersv1.GlobalLoadBalancersV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(globalLoadBalancersService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`ListMonitors(listMonitorsOptions *ListMonitorsOptions) - Operation response error`, func() {
		listMonitorsPath := "/instances/testString/monitors"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(listMonitorsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListMonitors with error: Operation response processing error`, func() {
				globalLoadBalancersService, serviceErr := globalloadbalancersv1.NewGlobalLoadBalancersV1(&globalloadbalancersv1.GlobalLoadBalancersV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalLoadBalancersService).ToNot(BeNil())

				// Construct an instance of the ListMonitorsOptions model
				listMonitorsOptionsModel := new(globalloadbalancersv1.ListMonitorsOptions)
				listMonitorsOptionsModel.InstanceID = core.StringPtr("testString")
				listMonitorsOptionsModel.XCorrelationID = core.StringPtr("testString")
				listMonitorsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := globalLoadBalancersService.ListMonitors(listMonitorsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ListMonitors(listMonitorsOptions *ListMonitorsOptions)`, func() {
		listMonitorsPath := "/instances/testString/monitors"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(listMonitorsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"monitors": [{"id": "5365b73c-ce6f-4d6f-ad9f-d9c131b26370", "name": "healthcheck-monitor", "description": "Load balancer monitor for glb.example.com.", "type": "HTTPS", "port": 8080, "interval": 60, "retries": 2, "timeout": 5, "method": "GET", "path": "/health", "headers": [{"name": "Host", "value": ["origin.example.com"]}], "allow_insecure": false, "expected_codes": "200", "expected_body": "alive", "created_on": "2019-01-01T05:20:00.12345Z", "modified_on": "2019-01-01T05:20:00.12345Z"}], "offset": 1, "limit": 20, "count": 1, "total_count": 200, "first": {"href": "https://api.dns-svcs.cloud.ibm.com/v1/instances/434f6c3e-6014-4124-a61d-2e910bca19b1/dnszones/example.com:d04d3a7a-7f6d-47d4-b811-08c5478fa1a4/resource_records?limit=20"}, "next": {"href": "https://api.dns-svcs.cloud.ibm.com/v1/instances/434f6c3e-6014-4124-a61d-2e910bca19b1/dnszones/example.com:d04d3a7a-7f6d-47d4-b811-08c5478fa1a4/resource_records?offset=20&limit=20"}}`)
				}))
			})
			It(`Invoke ListMonitors successfully`, func() {
				globalLoadBalancersService, serviceErr := globalloadbalancersv1.NewGlobalLoadBalancersV1(&globalloadbalancersv1.GlobalLoadBalancersV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalLoadBalancersService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := globalLoadBalancersService.ListMonitors(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListMonitorsOptions model
				listMonitorsOptionsModel := new(globalloadbalancersv1.ListMonitorsOptions)
				listMonitorsOptionsModel.InstanceID = core.StringPtr("testString")
				listMonitorsOptionsModel.XCorrelationID = core.StringPtr("testString")
				listMonitorsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = globalLoadBalancersService.ListMonitors(listMonitorsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke ListMonitors with error: Operation validation and request error`, func() {
				globalLoadBalancersService, serviceErr := globalloadbalancersv1.NewGlobalLoadBalancersV1(&globalloadbalancersv1.GlobalLoadBalancersV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalLoadBalancersService).ToNot(BeNil())

				// Construct an instance of the ListMonitorsOptions model
				listMonitorsOptionsModel := new(globalloadbalancersv1.ListMonitorsOptions)
				listMonitorsOptionsModel.InstanceID = core.StringPtr("testString")
				listMonitorsOptionsModel.XCorrelationID = core.StringPtr("testString")
				listMonitorsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := globalLoadBalancersService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := globalLoadBalancersService.ListMonitors(listMonitorsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListMonitorsOptions model with no property values
				listMonitorsOptionsModelNew := new(globalloadbalancersv1.ListMonitorsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = globalLoadBalancersService.ListMonitors(listMonitorsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateMonitor(createMonitorOptions *CreateMonitorOptions) - Operation response error`, func() {
		createMonitorPath := "/instances/testString/monitors"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(createMonitorPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateMonitor with error: Operation response processing error`, func() {
				globalLoadBalancersService, serviceErr := globalloadbalancersv1.NewGlobalLoadBalancersV1(&globalloadbalancersv1.GlobalLoadBalancersV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalLoadBalancersService).ToNot(BeNil())

				// Construct an instance of the HealthcheckHeader model
				healthcheckHeaderModel := new(globalloadbalancersv1.HealthcheckHeader)
				healthcheckHeaderModel.Name = core.StringPtr("Host")
				healthcheckHeaderModel.Value = []string{"origin.example.com"}

				// Construct an instance of the CreateMonitorOptions model
				createMonitorOptionsModel := new(globalloadbalancersv1.CreateMonitorOptions)
				createMonitorOptionsModel.InstanceID = core.StringPtr("testString")
				createMonitorOptionsModel.Name = core.StringPtr("healthcheck-monitor")
				createMonitorOptionsModel.Description = core.StringPtr("Load balancer monitor for glb.example.com.")
				createMonitorOptionsModel.Type = core.StringPtr("HTTPS")
				createMonitorOptionsModel.Port = core.Int64Ptr(int64(8080))
				createMonitorOptionsModel.Interval = core.Int64Ptr(int64(60))
				createMonitorOptionsModel.Retries = core.Int64Ptr(int64(2))
				createMonitorOptionsModel.Timeout = core.Int64Ptr(int64(5))
				createMonitorOptionsModel.Method = core.StringPtr("GET")
				createMonitorOptionsModel.Path = core.StringPtr("/health")
				createMonitorOptionsModel.HeadersVar = []globalloadbalancersv1.HealthcheckHeader{*healthcheckHeaderModel}
				createMonitorOptionsModel.AllowInsecure = core.BoolPtr(false)
				createMonitorOptionsModel.ExpectedCodes = core.StringPtr("200")
				createMonitorOptionsModel.ExpectedBody = core.StringPtr("alive")
				createMonitorOptionsModel.XCorrelationID = core.StringPtr("testString")
				createMonitorOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := globalLoadBalancersService.CreateMonitor(createMonitorOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`CreateMonitor(createMonitorOptions *CreateMonitorOptions)`, func() {
		createMonitorPath := "/instances/testString/monitors"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(createMonitorPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "5365b73c-ce6f-4d6f-ad9f-d9c131b26370", "name": "healthcheck-monitor", "description": "Load balancer monitor for glb.example.com.", "type": "HTTPS", "port": 8080, "interval": 60, "retries": 2, "timeout": 5, "method": "GET", "path": "/health", "headers": [{"name": "Host", "value": ["origin.example.com"]}], "allow_insecure": false, "expected_codes": "200", "expected_body": "alive", "created_on": "2019-01-01T05:20:00.12345Z", "modified_on": "2019-01-01T05:20:00.12345Z"}`)
				}))
			})
			It(`Invoke CreateMonitor successfully`, func() {
				globalLoadBalancersService, serviceErr := globalloadbalancersv1.NewGlobalLoadBalancersV1(&globalloadbalancersv1.GlobalLoadBalancersV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalLoadBalancersService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := globalLoadBalancersService.CreateMonitor(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the HealthcheckHeader model
				healthcheckHeaderModel := new(globalloadbalancersv1.HealthcheckHeader)
				healthcheckHeaderModel.Name = core.StringPtr("Host")
				healthcheckHeaderModel.Value = []string{"origin.example.com"}

				// Construct an instance of the CreateMonitorOptions model
				createMonitorOptionsModel := new(globalloadbalancersv1.CreateMonitorOptions)
				createMonitorOptionsModel.InstanceID = core.StringPtr("testString")
				createMonitorOptionsModel.Name = core.StringPtr("healthcheck-monitor")
				createMonitorOptionsModel.Description = core.StringPtr("Load balancer monitor for glb.example.com.")
				createMonitorOptionsModel.Type = core.StringPtr("HTTPS")
				createMonitorOptionsModel.Port = core.Int64Ptr(int64(8080))
				createMonitorOptionsModel.Interval = core.Int64Ptr(int64(60))
				createMonitorOptionsModel.Retries = core.Int64Ptr(int64(2))
				createMonitorOptionsModel.Timeout = core.Int64Ptr(int64(5))
				createMonitorOptionsModel.Method = core.StringPtr("GET")
				createMonitorOptionsModel.Path = core.StringPtr("/health")
				createMonitorOptionsModel.HeadersVar = []globalloadbalancersv1.HealthcheckHeader{*healthcheckHeaderModel}
				createMonitorOptionsModel.AllowInsecure = core.BoolPtr(false)
				createMonitorOptionsModel.ExpectedCodes = core.StringPtr("200")
				createMonitorOptionsModel.ExpectedBody = core.StringPtr("alive")
				createMonitorOptionsModel.XCorrelationID = core.StringPtr("testString")
				createMonitorOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = globalLoadBalancersService.CreateMonitor(createMonitorOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke CreateMonitor with error: Operation validation and request error`, func() {
				globalLoadBalancersService, serviceErr := globalloadbalancersv1.NewGlobalLoadBalancersV1(&globalloadbalancersv1.GlobalLoadBalancersV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalLoadBalancersService).ToNot(BeNil())

				// Construct an instance of the HealthcheckHeader model
				healthcheckHeaderModel := new(globalloadbalancersv1.HealthcheckHeader)
				healthcheckHeaderModel.Name = core.StringPtr("Host")
				healthcheckHeaderModel.Value = []string{"origin.example.com"}

				// Construct an instance of the CreateMonitorOptions model
				createMonitorOptionsModel := new(globalloadbalancersv1.CreateMonitorOptions)
				createMonitorOptionsModel.InstanceID = core.StringPtr("testString")
				createMonitorOptionsModel.Name = core.StringPtr("healthcheck-monitor")
				createMonitorOptionsModel.Description = core.StringPtr("Load balancer monitor for glb.example.com.")
				createMonitorOptionsModel.Type = core.StringPtr("HTTPS")
				createMonitorOptionsModel.Port = core.Int64Ptr(int64(8080))
				createMonitorOptionsModel.Interval = core.Int64Ptr(int64(60))
				createMonitorOptionsModel.Retries = core.Int64Ptr(int64(2))
				createMonitorOptionsModel.Timeout = core.Int64Ptr(int64(5))
				createMonitorOptionsModel.Method = core.StringPtr("GET")
				createMonitorOptionsModel.Path = core.StringPtr("/health")
				createMonitorOptionsModel.HeadersVar = []globalloadbalancersv1.HealthcheckHeader{*healthcheckHeaderModel}
				createMonitorOptionsModel.AllowInsecure = core.BoolPtr(false)
				createMonitorOptionsModel.ExpectedCodes = core.StringPtr("200")
				createMonitorOptionsModel.ExpectedBody = core.StringPtr("alive")
				createMonitorOptionsModel.XCorrelationID = core.StringPtr("testString")
				createMonitorOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := globalLoadBalancersService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := globalLoadBalancersService.CreateMonitor(createMonitorOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateMonitorOptions model with no property values
				createMonitorOptionsModelNew := new(globalloadbalancersv1.CreateMonitorOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = globalLoadBalancersService.CreateMonitor(createMonitorOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`DeleteMonitor(deleteMonitorOptions *DeleteMonitorOptions)`, func() {
		deleteMonitorPath := "/instances/testString/monitors/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(deleteMonitorPath))
					Expect(req.Method).To(Equal("DELETE"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteMonitor successfully`, func() {
				globalLoadBalancersService, serviceErr := globalloadbalancersv1.NewGlobalLoadBalancersV1(&globalloadbalancersv1.GlobalLoadBalancersV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalLoadBalancersService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := globalLoadBalancersService.DeleteMonitor(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteMonitorOptions model
				deleteMonitorOptionsModel := new(globalloadbalancersv1.DeleteMonitorOptions)
				deleteMonitorOptionsModel.InstanceID = core.StringPtr("testString")
				deleteMonitorOptionsModel.MonitorID = core.StringPtr("testString")
				deleteMonitorOptionsModel.XCorrelationID = core.StringPtr("testString")
				deleteMonitorOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = globalLoadBalancersService.DeleteMonitor(deleteMonitorOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteMonitor with error: Operation validation and request error`, func() {
				globalLoadBalancersService, serviceErr := globalloadbalancersv1.NewGlobalLoadBalancersV1(&globalloadbalancersv1.GlobalLoadBalancersV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalLoadBalancersService).ToNot(BeNil())

				// Construct an instance of the DeleteMonitorOptions model
				deleteMonitorOptionsModel := new(globalloadbalancersv1.DeleteMonitorOptions)
				deleteMonitorOptionsModel.InstanceID = core.StringPtr("testString")
				deleteMonitorOptionsModel.MonitorID = core.StringPtr("testString")
				deleteMonitorOptionsModel.XCorrelationID = core.StringPtr("testString")
				deleteMonitorOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := globalLoadBalancersService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := globalLoadBalancersService.DeleteMonitor(deleteMonitorOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteMonitorOptions model with no property values
				deleteMonitorOptionsModelNew := new(globalloadbalancersv1.DeleteMonitorOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = globalLoadBalancersService.DeleteMonitor(deleteMonitorOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetMonitor(getMonitorOptions *GetMonitorOptions) - Operation response error`, func() {
		getMonitorPath := "/instances/testString/monitors/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getMonitorPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetMonitor with error: Operation response processing error`, func() {
				globalLoadBalancersService, serviceErr := globalloadbalancersv1.NewGlobalLoadBalancersV1(&globalloadbalancersv1.GlobalLoadBalancersV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalLoadBalancersService).ToNot(BeNil())

				// Construct an instance of the GetMonitorOptions model
				getMonitorOptionsModel := new(globalloadbalancersv1.GetMonitorOptions)
				getMonitorOptionsModel.InstanceID = core.StringPtr("testString")
				getMonitorOptionsModel.MonitorID = core.StringPtr("testString")
				getMonitorOptionsModel.XCorrelationID = core.StringPtr("testString")
				getMonitorOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := globalLoadBalancersService.GetMonitor(getMonitorOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetMonitor(getMonitorOptions *GetMonitorOptions)`, func() {
		getMonitorPath := "/instances/testString/monitors/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getMonitorPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "5365b73c-ce6f-4d6f-ad9f-d9c131b26370", "name": "healthcheck-monitor", "description": "Load balancer monitor for glb.example.com.", "type": "HTTPS", "port": 8080, "interval": 60, "retries": 2, "timeout": 5, "method": "GET", "path": "/health", "headers": [{"name": "Host", "value": ["origin.example.com"]}], "allow_insecure": false, "expected_codes": "200", "expected_body": "alive", "created_on": "2019-01-01T05:20:00.12345Z", "modified_on": "2019-01-01T05:20:00.12345Z"}`)
				}))
			})
			It(`Invoke GetMonitor successfully`, func() {
				globalLoadBalancersService, serviceErr := globalloadbalancersv1.NewGlobalLoadBalancersV1(&globalloadbalancersv1.GlobalLoadBalancersV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalLoadBalancersService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := globalLoadBalancersService.GetMonitor(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetMonitorOptions model
				getMonitorOptionsModel := new(globalloadbalancersv1.GetMonitorOptions)
				getMonitorOptionsModel.InstanceID = core.StringPtr("testString")
				getMonitorOptionsModel.MonitorID = core.StringPtr("testString")
				getMonitorOptionsModel.XCorrelationID = core.StringPtr("testString")
				getMonitorOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = globalLoadBalancersService.GetMonitor(getMonitorOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke GetMonitor with error: Operation validation and request error`, func() {
				globalLoadBalancersService, serviceErr := globalloadbalancersv1.NewGlobalLoadBalancersV1(&globalloadbalancersv1.GlobalLoadBalancersV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalLoadBalancersService).ToNot(BeNil())

				// Construct an instance of the GetMonitorOptions model
				getMonitorOptionsModel := new(globalloadbalancersv1.GetMonitorOptions)
				getMonitorOptionsModel.InstanceID = core.StringPtr("testString")
				getMonitorOptionsModel.MonitorID = core.StringPtr("testString")
				getMonitorOptionsModel.XCorrelationID = core.StringPtr("testString")
				getMonitorOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := globalLoadBalancersService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := globalLoadBalancersService.GetMonitor(getMonitorOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetMonitorOptions model with no property values
				getMonitorOptionsModelNew := new(globalloadbalancersv1.GetMonitorOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = globalLoadBalancersService.GetMonitor(getMonitorOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateMonitor(updateMonitorOptions *UpdateMonitorOptions) - Operation response error`, func() {
		updateMonitorPath := "/instances/testString/monitors/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(updateMonitorPath))
					Expect(req.Method).To(Equal("PUT"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateMonitor with error: Operation response processing error`, func() {
				globalLoadBalancersService, serviceErr := globalloadbalancersv1.NewGlobalLoadBalancersV1(&globalloadbalancersv1.GlobalLoadBalancersV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalLoadBalancersService).ToNot(BeNil())

				// Construct an instance of the HealthcheckHeader model
				healthcheckHeaderModel := new(globalloadbalancersv1.HealthcheckHeader)
				healthcheckHeaderModel.Name = core.StringPtr("Host")
				healthcheckHeaderModel.Value = []string{"origin.example.com"}

				// Construct an instance of the UpdateMonitorOptions model
				updateMonitorOptionsModel := new(globalloadbalancersv1.UpdateMonitorOptions)
				updateMonitorOptionsModel.InstanceID = core.StringPtr("testString")
				updateMonitorOptionsModel.MonitorID = core.StringPtr("testString")
				updateMonitorOptionsModel.Name = core.StringPtr("healthcheck-monitor")
				updateMonitorOptionsModel.Description = core.StringPtr("Load balancer monitor for glb.example.com.")
				updateMonitorOptionsModel.Type = core.StringPtr("HTTPS")
				updateMonitorOptionsModel.Port = core.Int64Ptr(int64(8080))
				updateMonitorOptionsModel.Interval = core.Int64Ptr(int64(60))
				updateMonitorOptionsModel.Retries = core.Int64Ptr(int64(2))
				updateMonitorOptionsModel.Timeout = core.Int64Ptr(int64(5))
				updateMonitorOptionsModel.Method = core.StringPtr("GET")
				updateMonitorOptionsModel.Path = core.StringPtr("/health")
				updateMonitorOptionsModel.HeadersVar = []globalloadbalancersv1.HealthcheckHeader{*healthcheckHeaderModel}
				updateMonitorOptionsModel.AllowInsecure = core.BoolPtr(false)
				updateMonitorOptionsModel.ExpectedCodes = core.StringPtr("200")
				updateMonitorOptionsModel.ExpectedBody = core.StringPtr("alive")
				updateMonitorOptionsModel.XCorrelationID = core.StringPtr("testString")
				updateMonitorOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := globalLoadBalancersService.UpdateMonitor(updateMonitorOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UpdateMonitor(updateMonitorOptions *UpdateMonitorOptions)`, func() {
		updateMonitorPath := "/instances/testString/monitors/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(updateMonitorPath))
					Expect(req.Method).To(Equal("PUT"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "5365b73c-ce6f-4d6f-ad9f-d9c131b26370", "name": "healthcheck-monitor", "description": "Load balancer monitor for glb.example.com.", "type": "HTTPS", "port": 8080, "interval": 60, "retries": 2, "timeout": 5, "method": "GET", "path": "/health", "headers": [{"name": "Host", "value": ["origin.example.com"]}], "allow_insecure": false, "expected_codes": "200", "expected_body": "alive", "created_on": "2019-01-01T05:20:00.12345Z", "modified_on": "2019-01-01T05:20:00.12345Z"}`)
				}))
			})
			It(`Invoke UpdateMonitor successfully`, func() {
				globalLoadBalancersService, serviceErr := globalloadbalancersv1.NewGlobalLoadBalancersV1(&globalloadbalancersv1.GlobalLoadBalancersV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalLoadBalancersService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := globalLoadBalancersService.UpdateMonitor(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the HealthcheckHeader model
				healthcheckHeaderModel := new(globalloadbalancersv1.HealthcheckHeader)
				healthcheckHeaderModel.Name = core.StringPtr("Host")
				healthcheckHeaderModel.Value = []string{"origin.example.com"}

				// Construct an instance of the UpdateMonitorOptions model
				updateMonitorOptionsModel := new(globalloadbalancersv1.UpdateMonitorOptions)
				updateMonitorOptionsModel.InstanceID = core.StringPtr("testString")
				updateMonitorOptionsModel.MonitorID = core.StringPtr("testString")
				updateMonitorOptionsModel.Name = core.StringPtr("healthcheck-monitor")
				updateMonitorOptionsModel.Description = core.StringPtr("Load balancer monitor for glb.example.com.")
				updateMonitorOptionsModel.Type = core.StringPtr("HTTPS")
				updateMonitorOptionsModel.Port = core.Int64Ptr(int64(8080))
				updateMonitorOptionsModel.Interval = core.Int64Ptr(int64(60))
				updateMonitorOptionsModel.Retries = core.Int64Ptr(int64(2))
				updateMonitorOptionsModel.Timeout = core.Int64Ptr(int64(5))
				updateMonitorOptionsModel.Method = core.StringPtr("GET")
				updateMonitorOptionsModel.Path = core.StringPtr("/health")
				updateMonitorOptionsModel.HeadersVar = []globalloadbalancersv1.HealthcheckHeader{*healthcheckHeaderModel}
				updateMonitorOptionsModel.AllowInsecure = core.BoolPtr(false)
				updateMonitorOptionsModel.ExpectedCodes = core.StringPtr("200")
				updateMonitorOptionsModel.ExpectedBody = core.StringPtr("alive")
				updateMonitorOptionsModel.XCorrelationID = core.StringPtr("testString")
				updateMonitorOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = globalLoadBalancersService.UpdateMonitor(updateMonitorOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke UpdateMonitor with error: Operation validation and request error`, func() {
				globalLoadBalancersService, serviceErr := globalloadbalancersv1.NewGlobalLoadBalancersV1(&globalloadbalancersv1.GlobalLoadBalancersV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalLoadBalancersService).ToNot(BeNil())

				// Construct an instance of the HealthcheckHeader model
				healthcheckHeaderModel := new(globalloadbalancersv1.HealthcheckHeader)
				healthcheckHeaderModel.Name = core.StringPtr("Host")
				healthcheckHeaderModel.Value = []string{"origin.example.com"}

				// Construct an instance of the UpdateMonitorOptions model
				updateMonitorOptionsModel := new(globalloadbalancersv1.UpdateMonitorOptions)
				updateMonitorOptionsModel.InstanceID = core.StringPtr("testString")
				updateMonitorOptionsModel.MonitorID = core.StringPtr("testString")
				updateMonitorOptionsModel.Name = core.StringPtr("healthcheck-monitor")
				updateMonitorOptionsModel.Description = core.StringPtr("Load balancer monitor for glb.example.com.")
				updateMonitorOptionsModel.Type = core.StringPtr("HTTPS")
				updateMonitorOptionsModel.Port = core.Int64Ptr(int64(8080))
				updateMonitorOptionsModel.Interval = core.Int64Ptr(int64(60))
				updateMonitorOptionsModel.Retries = core.Int64Ptr(int64(2))
				updateMonitorOptionsModel.Timeout = core.Int64Ptr(int64(5))
				updateMonitorOptionsModel.Method = core.StringPtr("GET")
				updateMonitorOptionsModel.Path = core.StringPtr("/health")
				updateMonitorOptionsModel.HeadersVar = []globalloadbalancersv1.HealthcheckHeader{*healthcheckHeaderModel}
				updateMonitorOptionsModel.AllowInsecure = core.BoolPtr(false)
				updateMonitorOptionsModel.ExpectedCodes = core.StringPtr("200")
				updateMonitorOptionsModel.ExpectedBody = core.StringPtr("alive")
				updateMonitorOptionsModel.XCorrelationID = core.StringPtr("testString")
				updateMonitorOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := globalLoadBalancersService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := globalLoadBalancersService.UpdateMonitor(updateMonitorOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateMonitorOptions model with no property values
				updateMonitorOptionsModelNew := new(globalloadbalancersv1.UpdateMonitorOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = globalLoadBalancersService.UpdateMonitor(updateMonitorOptionsModelNew)
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
			globalLoadBalancersService, _ := globalloadbalancersv1.NewGlobalLoadBalancersV1(&globalloadbalancersv1.GlobalLoadBalancersV1Options{
				URL:           "http://globalloadbalancersv1modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
			})
			It(`Invoke NewCreateLoadBalancerOptions successfully`, func() {
				// Construct an instance of the LoadBalancerAzPoolsItem model
				loadBalancerAzPoolsItemModel := new(globalloadbalancersv1.LoadBalancerAzPoolsItem)
				Expect(loadBalancerAzPoolsItemModel).ToNot(BeNil())
				loadBalancerAzPoolsItemModel.AvailabilityZone = core.StringPtr("us-south-1")
				loadBalancerAzPoolsItemModel.Pools = []string{"0fc0bb7c-2fab-476e-8b9b-40fa14bf8e3d"}
				Expect(loadBalancerAzPoolsItemModel.AvailabilityZone).To(Equal(core.StringPtr("us-south-1")))
				Expect(loadBalancerAzPoolsItemModel.Pools).To(Equal([]string{"0fc0bb7c-2fab-476e-8b9b-40fa14bf8e3d"}))

				// Construct an instance of the CreateLoadBalancerOptions model
				instanceID := "testString"
				dnszoneID := "testString"
				createLoadBalancerOptionsModel := globalLoadBalancersService.NewCreateLoadBalancerOptions(instanceID, dnszoneID)
				createLoadBalancerOptionsModel.SetInstanceID("testString")
				createLoadBalancerOptionsModel.SetDnszoneID("testString")
				createLoadBalancerOptionsModel.SetName("glb.example.com")
				createLoadBalancerOptionsModel.SetDescription("Load balancer for glb.example.com.")
				createLoadBalancerOptionsModel.SetEnabled(true)
				createLoadBalancerOptionsModel.SetTTL(int64(120))
				createLoadBalancerOptionsModel.SetFallbackPool("24ccf79a-4ae0-4769-b4c8-17f8f230072e")
				createLoadBalancerOptionsModel.SetDefaultPools([]string{"testString"})
				createLoadBalancerOptionsModel.SetAzPools([]globalloadbalancersv1.LoadBalancerAzPoolsItem{*loadBalancerAzPoolsItemModel})
				createLoadBalancerOptionsModel.SetXCorrelationID("testString")
				createLoadBalancerOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createLoadBalancerOptionsModel).ToNot(BeNil())
				Expect(createLoadBalancerOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(createLoadBalancerOptionsModel.DnszoneID).To(Equal(core.StringPtr("testString")))
				Expect(createLoadBalancerOptionsModel.Name).To(Equal(core.StringPtr("glb.example.com")))
				Expect(createLoadBalancerOptionsModel.Description).To(Equal(core.StringPtr("Load balancer for glb.example.com.")))
				Expect(createLoadBalancerOptionsModel.Enabled).To(Equal(core.BoolPtr(true)))
				Expect(createLoadBalancerOptionsModel.TTL).To(Equal(core.Int64Ptr(int64(120))))
				Expect(createLoadBalancerOptionsModel.FallbackPool).To(Equal(core.StringPtr("24ccf79a-4ae0-4769-b4c8-17f8f230072e")))
				Expect(createLoadBalancerOptionsModel.DefaultPools).To(Equal([]string{"testString"}))
				Expect(createLoadBalancerOptionsModel.AzPools).To(Equal([]globalloadbalancersv1.LoadBalancerAzPoolsItem{*loadBalancerAzPoolsItemModel}))
				Expect(createLoadBalancerOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(createLoadBalancerOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateMonitorOptions successfully`, func() {
				// Construct an instance of the HealthcheckHeader model
				healthcheckHeaderModel := new(globalloadbalancersv1.HealthcheckHeader)
				Expect(healthcheckHeaderModel).ToNot(BeNil())
				healthcheckHeaderModel.Name = core.StringPtr("Host")
				healthcheckHeaderModel.Value = []string{"origin.example.com"}
				Expect(healthcheckHeaderModel.Name).To(Equal(core.StringPtr("Host")))
				Expect(healthcheckHeaderModel.Value).To(Equal([]string{"origin.example.com"}))

				// Construct an instance of the CreateMonitorOptions model
				instanceID := "testString"
				createMonitorOptionsModel := globalLoadBalancersService.NewCreateMonitorOptions(instanceID)
				createMonitorOptionsModel.SetInstanceID("testString")
				createMonitorOptionsModel.SetName("healthcheck-monitor")
				createMonitorOptionsModel.SetDescription("Load balancer monitor for glb.example.com.")
				createMonitorOptionsModel.SetType("HTTPS")
				createMonitorOptionsModel.SetPort(int64(8080))
				createMonitorOptionsModel.SetInterval(int64(60))
				createMonitorOptionsModel.SetRetries(int64(2))
				createMonitorOptionsModel.SetTimeout(int64(5))
				createMonitorOptionsModel.SetMethod("GET")
				createMonitorOptionsModel.SetPath("/health")
				createMonitorOptionsModel.SetHeadersVar([]globalloadbalancersv1.HealthcheckHeader{*healthcheckHeaderModel})
				createMonitorOptionsModel.SetAllowInsecure(false)
				createMonitorOptionsModel.SetExpectedCodes("200")
				createMonitorOptionsModel.SetExpectedBody("alive")
				createMonitorOptionsModel.SetXCorrelationID("testString")
				createMonitorOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createMonitorOptionsModel).ToNot(BeNil())
				Expect(createMonitorOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(createMonitorOptionsModel.Name).To(Equal(core.StringPtr("healthcheck-monitor")))
				Expect(createMonitorOptionsModel.Description).To(Equal(core.StringPtr("Load balancer monitor for glb.example.com.")))
				Expect(createMonitorOptionsModel.Type).To(Equal(core.StringPtr("HTTPS")))
				Expect(createMonitorOptionsModel.Port).To(Equal(core.Int64Ptr(int64(8080))))
				Expect(createMonitorOptionsModel.Interval).To(Equal(core.Int64Ptr(int64(60))))
				Expect(createMonitorOptionsModel.Retries).To(Equal(core.Int64Ptr(int64(2))))
				Expect(createMonitorOptionsModel.Timeout).To(Equal(core.Int64Ptr(int64(5))))
				Expect(createMonitorOptionsModel.Method).To(Equal(core.StringPtr("GET")))
				Expect(createMonitorOptionsModel.Path).To(Equal(core.StringPtr("/health")))
				Expect(createMonitorOptionsModel.HeadersVar).To(Equal([]globalloadbalancersv1.HealthcheckHeader{*healthcheckHeaderModel}))
				Expect(createMonitorOptionsModel.AllowInsecure).To(Equal(core.BoolPtr(false)))
				Expect(createMonitorOptionsModel.ExpectedCodes).To(Equal(core.StringPtr("200")))
				Expect(createMonitorOptionsModel.ExpectedBody).To(Equal(core.StringPtr("alive")))
				Expect(createMonitorOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(createMonitorOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreatePoolOptions successfully`, func() {
				// Construct an instance of the OriginInput model
				originInputModel := new(globalloadbalancersv1.OriginInput)
				Expect(originInputModel).ToNot(BeNil())
				originInputModel.Name = core.StringPtr("app-server-1")
				originInputModel.Description = core.StringPtr("description of the origin server")
				originInputModel.Address = core.StringPtr("10.10.16.8")
				originInputModel.Enabled = core.BoolPtr(true)
				Expect(originInputModel.Name).To(Equal(core.StringPtr("app-server-1")))
				Expect(originInputModel.Description).To(Equal(core.StringPtr("description of the origin server")))
				Expect(originInputModel.Address).To(Equal(core.StringPtr("10.10.16.8")))
				Expect(originInputModel.Enabled).To(Equal(core.BoolPtr(true)))

				// Construct an instance of the CreatePoolOptions model
				instanceID := "testString"
				createPoolOptionsModel := globalLoadBalancersService.NewCreatePoolOptions(instanceID)
				createPoolOptionsModel.SetInstanceID("testString")
				createPoolOptionsModel.SetName("dal10-az-pool")
				createPoolOptionsModel.SetDescription("Load balancer pool for dal10 availability zone.")
				createPoolOptionsModel.SetEnabled(true)
				createPoolOptionsModel.SetHealthyOriginsThreshold(int64(1))
				createPoolOptionsModel.SetOrigins([]globalloadbalancersv1.OriginInput{*originInputModel})
				createPoolOptionsModel.SetMonitor("7dd6841c-264e-11ea-88df-062967242a6a")
				createPoolOptionsModel.SetNotificationChannel("https://mywebsite.com/dns/webhook")
				createPoolOptionsModel.SetHealthcheckRegion("us-south")
				createPoolOptionsModel.SetHealthcheckSubnets([]string{"0716-a4c0c123-594c-4ef4-ace3-a08858540b5e"})
				createPoolOptionsModel.SetXCorrelationID("testString")
				createPoolOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createPoolOptionsModel).ToNot(BeNil())
				Expect(createPoolOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(createPoolOptionsModel.Name).To(Equal(core.StringPtr("dal10-az-pool")))
				Expect(createPoolOptionsModel.Description).To(Equal(core.StringPtr("Load balancer pool for dal10 availability zone.")))
				Expect(createPoolOptionsModel.Enabled).To(Equal(core.BoolPtr(true)))
				Expect(createPoolOptionsModel.HealthyOriginsThreshold).To(Equal(core.Int64Ptr(int64(1))))
				Expect(createPoolOptionsModel.Origins).To(Equal([]globalloadbalancersv1.OriginInput{*originInputModel}))
				Expect(createPoolOptionsModel.Monitor).To(Equal(core.StringPtr("7dd6841c-264e-11ea-88df-062967242a6a")))
				Expect(createPoolOptionsModel.NotificationChannel).To(Equal(core.StringPtr("https://mywebsite.com/dns/webhook")))
				Expect(createPoolOptionsModel.HealthcheckRegion).To(Equal(core.StringPtr("us-south")))
				Expect(createPoolOptionsModel.HealthcheckSubnets).To(Equal([]string{"0716-a4c0c123-594c-4ef4-ace3-a08858540b5e"}))
				Expect(createPoolOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(createPoolOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteLoadBalancerOptions successfully`, func() {
				// Construct an instance of the DeleteLoadBalancerOptions model
				instanceID := "testString"
				dnszoneID := "testString"
				lbID := "testString"
				deleteLoadBalancerOptionsModel := globalLoadBalancersService.NewDeleteLoadBalancerOptions(instanceID, dnszoneID, lbID)
				deleteLoadBalancerOptionsModel.SetInstanceID("testString")
				deleteLoadBalancerOptionsModel.SetDnszoneID("testString")
				deleteLoadBalancerOptionsModel.SetLbID("testString")
				deleteLoadBalancerOptionsModel.SetXCorrelationID("testString")
				deleteLoadBalancerOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteLoadBalancerOptionsModel).ToNot(BeNil())
				Expect(deleteLoadBalancerOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(deleteLoadBalancerOptionsModel.DnszoneID).To(Equal(core.StringPtr("testString")))
				Expect(deleteLoadBalancerOptionsModel.LbID).To(Equal(core.StringPtr("testString")))
				Expect(deleteLoadBalancerOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(deleteLoadBalancerOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteMonitorOptions successfully`, func() {
				// Construct an instance of the DeleteMonitorOptions model
				instanceID := "testString"
				monitorID := "testString"
				deleteMonitorOptionsModel := globalLoadBalancersService.NewDeleteMonitorOptions(instanceID, monitorID)
				deleteMonitorOptionsModel.SetInstanceID("testString")
				deleteMonitorOptionsModel.SetMonitorID("testString")
				deleteMonitorOptionsModel.SetXCorrelationID("testString")
				deleteMonitorOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteMonitorOptionsModel).ToNot(BeNil())
				Expect(deleteMonitorOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(deleteMonitorOptionsModel.MonitorID).To(Equal(core.StringPtr("testString")))
				Expect(deleteMonitorOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(deleteMonitorOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeletePoolOptions successfully`, func() {
				// Construct an instance of the DeletePoolOptions model
				instanceID := "testString"
				poolID := "testString"
				deletePoolOptionsModel := globalLoadBalancersService.NewDeletePoolOptions(instanceID, poolID)
				deletePoolOptionsModel.SetInstanceID("testString")
				deletePoolOptionsModel.SetPoolID("testString")
				deletePoolOptionsModel.SetXCorrelationID("testString")
				deletePoolOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deletePoolOptionsModel).ToNot(BeNil())
				Expect(deletePoolOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(deletePoolOptionsModel.PoolID).To(Equal(core.StringPtr("testString")))
				Expect(deletePoolOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(deletePoolOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetLoadBalancerOptions successfully`, func() {
				// Construct an instance of the GetLoadBalancerOptions model
				instanceID := "testString"
				dnszoneID := "testString"
				lbID := "testString"
				getLoadBalancerOptionsModel := globalLoadBalancersService.NewGetLoadBalancerOptions(instanceID, dnszoneID, lbID)
				getLoadBalancerOptionsModel.SetInstanceID("testString")
				getLoadBalancerOptionsModel.SetDnszoneID("testString")
				getLoadBalancerOptionsModel.SetLbID("testString")
				getLoadBalancerOptionsModel.SetXCorrelationID("testString")
				getLoadBalancerOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getLoadBalancerOptionsModel).ToNot(BeNil())
				Expect(getLoadBalancerOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(getLoadBalancerOptionsModel.DnszoneID).To(Equal(core.StringPtr("testString")))
				Expect(getLoadBalancerOptionsModel.LbID).To(Equal(core.StringPtr("testString")))
				Expect(getLoadBalancerOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(getLoadBalancerOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetMonitorOptions successfully`, func() {
				// Construct an instance of the GetMonitorOptions model
				instanceID := "testString"
				monitorID := "testString"
				getMonitorOptionsModel := globalLoadBalancersService.NewGetMonitorOptions(instanceID, monitorID)
				getMonitorOptionsModel.SetInstanceID("testString")
				getMonitorOptionsModel.SetMonitorID("testString")
				getMonitorOptionsModel.SetXCorrelationID("testString")
				getMonitorOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getMonitorOptionsModel).ToNot(BeNil())
				Expect(getMonitorOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(getMonitorOptionsModel.MonitorID).To(Equal(core.StringPtr("testString")))
				Expect(getMonitorOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(getMonitorOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetPoolOptions successfully`, func() {
				// Construct an instance of the GetPoolOptions model
				instanceID := "testString"
				poolID := "testString"
				getPoolOptionsModel := globalLoadBalancersService.NewGetPoolOptions(instanceID, poolID)
				getPoolOptionsModel.SetInstanceID("testString")
				getPoolOptionsModel.SetPoolID("testString")
				getPoolOptionsModel.SetXCorrelationID("testString")
				getPoolOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getPoolOptionsModel).ToNot(BeNil())
				Expect(getPoolOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(getPoolOptionsModel.PoolID).To(Equal(core.StringPtr("testString")))
				Expect(getPoolOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(getPoolOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListLoadBalancersOptions successfully`, func() {
				// Construct an instance of the ListLoadBalancersOptions model
				instanceID := "testString"
				dnszoneID := "testString"
				listLoadBalancersOptionsModel := globalLoadBalancersService.NewListLoadBalancersOptions(instanceID, dnszoneID)
				listLoadBalancersOptionsModel.SetInstanceID("testString")
				listLoadBalancersOptionsModel.SetDnszoneID("testString")
				listLoadBalancersOptionsModel.SetXCorrelationID("testString")
				listLoadBalancersOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listLoadBalancersOptionsModel).ToNot(BeNil())
				Expect(listLoadBalancersOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(listLoadBalancersOptionsModel.DnszoneID).To(Equal(core.StringPtr("testString")))
				Expect(listLoadBalancersOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(listLoadBalancersOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListMonitorsOptions successfully`, func() {
				// Construct an instance of the ListMonitorsOptions model
				instanceID := "testString"
				listMonitorsOptionsModel := globalLoadBalancersService.NewListMonitorsOptions(instanceID)
				listMonitorsOptionsModel.SetInstanceID("testString")
				listMonitorsOptionsModel.SetXCorrelationID("testString")
				listMonitorsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listMonitorsOptionsModel).ToNot(BeNil())
				Expect(listMonitorsOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(listMonitorsOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(listMonitorsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListPoolsOptions successfully`, func() {
				// Construct an instance of the ListPoolsOptions model
				instanceID := "testString"
				listPoolsOptionsModel := globalLoadBalancersService.NewListPoolsOptions(instanceID)
				listPoolsOptionsModel.SetInstanceID("testString")
				listPoolsOptionsModel.SetXCorrelationID("testString")
				listPoolsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listPoolsOptionsModel).ToNot(BeNil())
				Expect(listPoolsOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(listPoolsOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(listPoolsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateLoadBalancerOptions successfully`, func() {
				// Construct an instance of the LoadBalancerAzPoolsItem model
				loadBalancerAzPoolsItemModel := new(globalloadbalancersv1.LoadBalancerAzPoolsItem)
				Expect(loadBalancerAzPoolsItemModel).ToNot(BeNil())
				loadBalancerAzPoolsItemModel.AvailabilityZone = core.StringPtr("us-south-1")
				loadBalancerAzPoolsItemModel.Pools = []string{"0fc0bb7c-2fab-476e-8b9b-40fa14bf8e3d"}
				Expect(loadBalancerAzPoolsItemModel.AvailabilityZone).To(Equal(core.StringPtr("us-south-1")))
				Expect(loadBalancerAzPoolsItemModel.Pools).To(Equal([]string{"0fc0bb7c-2fab-476e-8b9b-40fa14bf8e3d"}))

				// Construct an instance of the UpdateLoadBalancerOptions model
				instanceID := "testString"
				dnszoneID := "testString"
				lbID := "testString"
				updateLoadBalancerOptionsModel := globalLoadBalancersService.NewUpdateLoadBalancerOptions(instanceID, dnszoneID, lbID)
				updateLoadBalancerOptionsModel.SetInstanceID("testString")
				updateLoadBalancerOptionsModel.SetDnszoneID("testString")
				updateLoadBalancerOptionsModel.SetLbID("testString")
				updateLoadBalancerOptionsModel.SetName("glb.example.com")
				updateLoadBalancerOptionsModel.SetDescription("Load balancer for glb.example.com.")
				updateLoadBalancerOptionsModel.SetEnabled(true)
				updateLoadBalancerOptionsModel.SetTTL(int64(120))
				updateLoadBalancerOptionsModel.SetFallbackPool("24ccf79a-4ae0-4769-b4c8-17f8f230072e")
				updateLoadBalancerOptionsModel.SetDefaultPools([]string{"testString"})
				updateLoadBalancerOptionsModel.SetAzPools([]globalloadbalancersv1.LoadBalancerAzPoolsItem{*loadBalancerAzPoolsItemModel})
				updateLoadBalancerOptionsModel.SetXCorrelationID("testString")
				updateLoadBalancerOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateLoadBalancerOptionsModel).ToNot(BeNil())
				Expect(updateLoadBalancerOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(updateLoadBalancerOptionsModel.DnszoneID).To(Equal(core.StringPtr("testString")))
				Expect(updateLoadBalancerOptionsModel.LbID).To(Equal(core.StringPtr("testString")))
				Expect(updateLoadBalancerOptionsModel.Name).To(Equal(core.StringPtr("glb.example.com")))
				Expect(updateLoadBalancerOptionsModel.Description).To(Equal(core.StringPtr("Load balancer for glb.example.com.")))
				Expect(updateLoadBalancerOptionsModel.Enabled).To(Equal(core.BoolPtr(true)))
				Expect(updateLoadBalancerOptionsModel.TTL).To(Equal(core.Int64Ptr(int64(120))))
				Expect(updateLoadBalancerOptionsModel.FallbackPool).To(Equal(core.StringPtr("24ccf79a-4ae0-4769-b4c8-17f8f230072e")))
				Expect(updateLoadBalancerOptionsModel.DefaultPools).To(Equal([]string{"testString"}))
				Expect(updateLoadBalancerOptionsModel.AzPools).To(Equal([]globalloadbalancersv1.LoadBalancerAzPoolsItem{*loadBalancerAzPoolsItemModel}))
				Expect(updateLoadBalancerOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(updateLoadBalancerOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateMonitorOptions successfully`, func() {
				// Construct an instance of the HealthcheckHeader model
				healthcheckHeaderModel := new(globalloadbalancersv1.HealthcheckHeader)
				Expect(healthcheckHeaderModel).ToNot(BeNil())
				healthcheckHeaderModel.Name = core.StringPtr("Host")
				healthcheckHeaderModel.Value = []string{"origin.example.com"}
				Expect(healthcheckHeaderModel.Name).To(Equal(core.StringPtr("Host")))
				Expect(healthcheckHeaderModel.Value).To(Equal([]string{"origin.example.com"}))

				// Construct an instance of the UpdateMonitorOptions model
				instanceID := "testString"
				monitorID := "testString"
				updateMonitorOptionsModel := globalLoadBalancersService.NewUpdateMonitorOptions(instanceID, monitorID)
				updateMonitorOptionsModel.SetInstanceID("testString")
				updateMonitorOptionsModel.SetMonitorID("testString")
				updateMonitorOptionsModel.SetName("healthcheck-monitor")
				updateMonitorOptionsModel.SetDescription("Load balancer monitor for glb.example.com.")
				updateMonitorOptionsModel.SetType("HTTPS")
				updateMonitorOptionsModel.SetPort(int64(8080))
				updateMonitorOptionsModel.SetInterval(int64(60))
				updateMonitorOptionsModel.SetRetries(int64(2))
				updateMonitorOptionsModel.SetTimeout(int64(5))
				updateMonitorOptionsModel.SetMethod("GET")
				updateMonitorOptionsModel.SetPath("/health")
				updateMonitorOptionsModel.SetHeadersVar([]globalloadbalancersv1.HealthcheckHeader{*healthcheckHeaderModel})
				updateMonitorOptionsModel.SetAllowInsecure(false)
				updateMonitorOptionsModel.SetExpectedCodes("200")
				updateMonitorOptionsModel.SetExpectedBody("alive")
				updateMonitorOptionsModel.SetXCorrelationID("testString")
				updateMonitorOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateMonitorOptionsModel).ToNot(BeNil())
				Expect(updateMonitorOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(updateMonitorOptionsModel.MonitorID).To(Equal(core.StringPtr("testString")))
				Expect(updateMonitorOptionsModel.Name).To(Equal(core.StringPtr("healthcheck-monitor")))
				Expect(updateMonitorOptionsModel.Description).To(Equal(core.StringPtr("Load balancer monitor for glb.example.com.")))
				Expect(updateMonitorOptionsModel.Type).To(Equal(core.StringPtr("HTTPS")))
				Expect(updateMonitorOptionsModel.Port).To(Equal(core.Int64Ptr(int64(8080))))
				Expect(updateMonitorOptionsModel.Interval).To(Equal(core.Int64Ptr(int64(60))))
				Expect(updateMonitorOptionsModel.Retries).To(Equal(core.Int64Ptr(int64(2))))
				Expect(updateMonitorOptionsModel.Timeout).To(Equal(core.Int64Ptr(int64(5))))
				Expect(updateMonitorOptionsModel.Method).To(Equal(core.StringPtr("GET")))
				Expect(updateMonitorOptionsModel.Path).To(Equal(core.StringPtr("/health")))
				Expect(updateMonitorOptionsModel.HeadersVar).To(Equal([]globalloadbalancersv1.HealthcheckHeader{*healthcheckHeaderModel}))
				Expect(updateMonitorOptionsModel.AllowInsecure).To(Equal(core.BoolPtr(false)))
				Expect(updateMonitorOptionsModel.ExpectedCodes).To(Equal(core.StringPtr("200")))
				Expect(updateMonitorOptionsModel.ExpectedBody).To(Equal(core.StringPtr("alive")))
				Expect(updateMonitorOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(updateMonitorOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdatePoolOptions successfully`, func() {
				// Construct an instance of the OriginInput model
				originInputModel := new(globalloadbalancersv1.OriginInput)
				Expect(originInputModel).ToNot(BeNil())
				originInputModel.Name = core.StringPtr("app-server-1")
				originInputModel.Description = core.StringPtr("description of the origin server")
				originInputModel.Address = core.StringPtr("10.10.16.8")
				originInputModel.Enabled = core.BoolPtr(true)
				Expect(originInputModel.Name).To(Equal(core.StringPtr("app-server-1")))
				Expect(originInputModel.Description).To(Equal(core.StringPtr("description of the origin server")))
				Expect(originInputModel.Address).To(Equal(core.StringPtr("10.10.16.8")))
				Expect(originInputModel.Enabled).To(Equal(core.BoolPtr(true)))

				// Construct an instance of the UpdatePoolOptions model
				instanceID := "testString"
				poolID := "testString"
				updatePoolOptionsModel := globalLoadBalancersService.NewUpdatePoolOptions(instanceID, poolID)
				updatePoolOptionsModel.SetInstanceID("testString")
				updatePoolOptionsModel.SetPoolID("testString")
				updatePoolOptionsModel.SetName("dal10-az-pool")
				updatePoolOptionsModel.SetDescription("Load balancer pool for dal10 availability zone.")
				updatePoolOptionsModel.SetEnabled(true)
				updatePoolOptionsModel.SetHealthyOriginsThreshold(int64(1))
				updatePoolOptionsModel.SetOrigins([]globalloadbalancersv1.OriginInput{*originInputModel})
				updatePoolOptionsModel.SetMonitor("7dd6841c-264e-11ea-88df-062967242a6a")
				updatePoolOptionsModel.SetNotificationChannel("https://mywebsite.com/dns/webhook")
				updatePoolOptionsModel.SetHealthcheckRegion("us-south")
				updatePoolOptionsModel.SetHealthcheckSubnets([]string{"0716-a4c0c123-594c-4ef4-ace3-a08858540b5e"})
				updatePoolOptionsModel.SetXCorrelationID("testString")
				updatePoolOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updatePoolOptionsModel).ToNot(BeNil())
				Expect(updatePoolOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(updatePoolOptionsModel.PoolID).To(Equal(core.StringPtr("testString")))
				Expect(updatePoolOptionsModel.Name).To(Equal(core.StringPtr("dal10-az-pool")))
				Expect(updatePoolOptionsModel.Description).To(Equal(core.StringPtr("Load balancer pool for dal10 availability zone.")))
				Expect(updatePoolOptionsModel.Enabled).To(Equal(core.BoolPtr(true)))
				Expect(updatePoolOptionsModel.HealthyOriginsThreshold).To(Equal(core.Int64Ptr(int64(1))))
				Expect(updatePoolOptionsModel.Origins).To(Equal([]globalloadbalancersv1.OriginInput{*originInputModel}))
				Expect(updatePoolOptionsModel.Monitor).To(Equal(core.StringPtr("7dd6841c-264e-11ea-88df-062967242a6a")))
				Expect(updatePoolOptionsModel.NotificationChannel).To(Equal(core.StringPtr("https://mywebsite.com/dns/webhook")))
				Expect(updatePoolOptionsModel.HealthcheckRegion).To(Equal(core.StringPtr("us-south")))
				Expect(updatePoolOptionsModel.HealthcheckSubnets).To(Equal([]string{"0716-a4c0c123-594c-4ef4-ace3-a08858540b5e"}))
				Expect(updatePoolOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(updatePoolOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewHealthcheckHeader successfully`, func() {
				name := "Host"
				value := []string{"origin.example.com"}
				model, err := globalLoadBalancersService.NewHealthcheckHeader(name, value)
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
