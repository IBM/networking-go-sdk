/**
 * (C) Copyright IBM Corp. 2021.
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

package filtersv1_test

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
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.ibm.com/ibmcloud/networking-go-sdk/filtersv1"
)

var _ = Describe(`FiltersV1`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			filtersService, serviceErr := filtersv1.NewFiltersV1(&filtersv1.FiltersV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(filtersService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			filtersService, serviceErr := filtersv1.NewFiltersV1(&filtersv1.FiltersV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(filtersService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			filtersService, serviceErr := filtersv1.NewFiltersV1(&filtersv1.FiltersV1Options{
				URL: "https://filtersv1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(filtersService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"FILTERS_URL":       "https://filtersv1/api",
				"FILTERS_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				filtersService, serviceErr := filtersv1.NewFiltersV1UsingExternalConfig(&filtersv1.FiltersV1Options{})
				Expect(filtersService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := filtersService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != filtersService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(filtersService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(filtersService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				filtersService, serviceErr := filtersv1.NewFiltersV1UsingExternalConfig(&filtersv1.FiltersV1Options{
					URL: "https://testService/api",
				})
				Expect(filtersService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(filtersService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := filtersService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != filtersService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(filtersService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(filtersService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				filtersService, serviceErr := filtersv1.NewFiltersV1UsingExternalConfig(&filtersv1.FiltersV1Options{})
				err := filtersService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(filtersService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(filtersService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := filtersService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != filtersService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(filtersService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(filtersService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"FILTERS_URL":       "https://filtersv1/api",
				"FILTERS_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			filtersService, serviceErr := filtersv1.NewFiltersV1UsingExternalConfig(&filtersv1.FiltersV1Options{})

			It(`Instantiate service client with error`, func() {
				Expect(filtersService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"FILTERS_AUTH_TYPE": "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			filtersService, serviceErr := filtersv1.NewFiltersV1UsingExternalConfig(&filtersv1.FiltersV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(filtersService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = filtersv1.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})

	Describe(`ListAllFilters(listAllFiltersOptions *ListAllFiltersOptions)`, func() {
		listAllFiltersPath := "/v1/testString/zones/testString/filters"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listAllFiltersPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Auth-User-Token"]).ToNot(BeNil())
					Expect(req.Header["X-Auth-User-Token"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Accept"]).ToNot(BeNil())
					Expect(req.Header["Accept"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "*/*")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `This is a mock binary response.`)
				}))
			})
			It(`Invoke ListAllFilters successfully with retries`, func() {
				filtersService, serviceErr := filtersv1.NewFiltersV1(&filtersv1.FiltersV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(filtersService).ToNot(BeNil())
				filtersService.EnableRetries(0, 0)

				// Construct an instance of the ListAllFiltersOptions model
				listAllFiltersOptionsModel := new(filtersv1.ListAllFiltersOptions)
				listAllFiltersOptionsModel.XAuthUserToken = core.StringPtr("testString")
				listAllFiltersOptionsModel.Crn = core.StringPtr("testString")
				listAllFiltersOptionsModel.ZoneIdentifier = core.StringPtr("testString")
				listAllFiltersOptionsModel.Accept = core.StringPtr("testString")
				listAllFiltersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := filtersService.ListAllFiltersWithContext(ctx, listAllFiltersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				filtersService.DisableRetries()
				result, response, operationErr := filtersService.ListAllFilters(listAllFiltersOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = filtersService.ListAllFiltersWithContext(ctx, listAllFiltersOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listAllFiltersPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Auth-User-Token"]).ToNot(BeNil())
					Expect(req.Header["X-Auth-User-Token"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Accept"]).ToNot(BeNil())
					Expect(req.Header["Accept"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "*/*")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `This is a mock binary response.`)
				}))
			})
			It(`Invoke ListAllFilters successfully`, func() {
				filtersService, serviceErr := filtersv1.NewFiltersV1(&filtersv1.FiltersV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(filtersService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := filtersService.ListAllFilters(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListAllFiltersOptions model
				listAllFiltersOptionsModel := new(filtersv1.ListAllFiltersOptions)
				listAllFiltersOptionsModel.XAuthUserToken = core.StringPtr("testString")
				listAllFiltersOptionsModel.Crn = core.StringPtr("testString")
				listAllFiltersOptionsModel.ZoneIdentifier = core.StringPtr("testString")
				listAllFiltersOptionsModel.Accept = core.StringPtr("testString")
				listAllFiltersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = filtersService.ListAllFilters(listAllFiltersOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListAllFilters with error: Operation validation and request error`, func() {
				filtersService, serviceErr := filtersv1.NewFiltersV1(&filtersv1.FiltersV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(filtersService).ToNot(BeNil())

				// Construct an instance of the ListAllFiltersOptions model
				listAllFiltersOptionsModel := new(filtersv1.ListAllFiltersOptions)
				listAllFiltersOptionsModel.XAuthUserToken = core.StringPtr("testString")
				listAllFiltersOptionsModel.Crn = core.StringPtr("testString")
				listAllFiltersOptionsModel.ZoneIdentifier = core.StringPtr("testString")
				listAllFiltersOptionsModel.Accept = core.StringPtr("testString")
				listAllFiltersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := filtersService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := filtersService.ListAllFilters(listAllFiltersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListAllFiltersOptions model with no property values
				listAllFiltersOptionsModelNew := new(filtersv1.ListAllFiltersOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = filtersService.ListAllFilters(listAllFiltersOptionsModelNew)
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
			filtersService, serviceErr := filtersv1.NewFiltersV1(&filtersv1.FiltersV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(filtersService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			filtersService, serviceErr := filtersv1.NewFiltersV1(&filtersv1.FiltersV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(filtersService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			filtersService, serviceErr := filtersv1.NewFiltersV1(&filtersv1.FiltersV1Options{
				URL: "https://filtersv1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(filtersService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"FILTERS_URL":       "https://filtersv1/api",
				"FILTERS_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				filtersService, serviceErr := filtersv1.NewFiltersV1UsingExternalConfig(&filtersv1.FiltersV1Options{})
				Expect(filtersService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := filtersService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != filtersService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(filtersService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(filtersService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				filtersService, serviceErr := filtersv1.NewFiltersV1UsingExternalConfig(&filtersv1.FiltersV1Options{
					URL: "https://testService/api",
				})
				Expect(filtersService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(filtersService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := filtersService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != filtersService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(filtersService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(filtersService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				filtersService, serviceErr := filtersv1.NewFiltersV1UsingExternalConfig(&filtersv1.FiltersV1Options{})
				err := filtersService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(filtersService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(filtersService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := filtersService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != filtersService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(filtersService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(filtersService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"FILTERS_URL":       "https://filtersv1/api",
				"FILTERS_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			filtersService, serviceErr := filtersv1.NewFiltersV1UsingExternalConfig(&filtersv1.FiltersV1Options{})

			It(`Instantiate service client with error`, func() {
				Expect(filtersService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"FILTERS_AUTH_TYPE": "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			filtersService, serviceErr := filtersv1.NewFiltersV1UsingExternalConfig(&filtersv1.FiltersV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(filtersService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = filtersv1.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})

	Describe(`CreateFilter(createFilterOptions *CreateFilterOptions)`, func() {
		createFilterPath := "/v1/testString/zones/testString/filters"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createFilterPath))
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

					Expect(req.Header["X-Auth-User-Token"]).ToNot(BeNil())
					Expect(req.Header["X-Auth-User-Token"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Accept"]).ToNot(BeNil())
					Expect(req.Header["Accept"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "*/*")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `This is a mock binary response.`)
				}))
			})
			It(`Invoke CreateFilter successfully with retries`, func() {
				filtersService, serviceErr := filtersv1.NewFiltersV1(&filtersv1.FiltersV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(filtersService).ToNot(BeNil())
				filtersService.EnableRetries(0, 0)

				// Construct an instance of the FilterInput model
				filterInputModel := new(filtersv1.FilterInput)
				filterInputModel.Expression = core.StringPtr(`not http.request.uri.path matches "^/api/.*$"`)
				filterInputModel.Paused = core.BoolPtr(false)
				filterInputModel.Description = core.StringPtr("not /api")

				// Construct an instance of the CreateFilterOptions model
				createFilterOptionsModel := new(filtersv1.CreateFilterOptions)
				createFilterOptionsModel.XAuthUserToken = core.StringPtr("testString")
				createFilterOptionsModel.Crn = core.StringPtr("testString")
				createFilterOptionsModel.ZoneIdentifier = core.StringPtr("testString")
				createFilterOptionsModel.Accept = core.StringPtr("testString")
				createFilterOptionsModel.FilterInput = []filtersv1.FilterInput{*filterInputModel}
				createFilterOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := filtersService.CreateFilterWithContext(ctx, createFilterOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				filtersService.DisableRetries()
				result, response, operationErr := filtersService.CreateFilter(createFilterOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = filtersService.CreateFilterWithContext(ctx, createFilterOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createFilterPath))
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

					Expect(req.Header["X-Auth-User-Token"]).ToNot(BeNil())
					Expect(req.Header["X-Auth-User-Token"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Accept"]).ToNot(BeNil())
					Expect(req.Header["Accept"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "*/*")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `This is a mock binary response.`)
				}))
			})
			It(`Invoke CreateFilter successfully`, func() {
				filtersService, serviceErr := filtersv1.NewFiltersV1(&filtersv1.FiltersV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(filtersService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := filtersService.CreateFilter(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the FilterInput model
				filterInputModel := new(filtersv1.FilterInput)
				filterInputModel.Expression = core.StringPtr(`not http.request.uri.path matches "^/api/.*$"`)
				filterInputModel.Paused = core.BoolPtr(false)
				filterInputModel.Description = core.StringPtr("not /api")

				// Construct an instance of the CreateFilterOptions model
				createFilterOptionsModel := new(filtersv1.CreateFilterOptions)
				createFilterOptionsModel.XAuthUserToken = core.StringPtr("testString")
				createFilterOptionsModel.Crn = core.StringPtr("testString")
				createFilterOptionsModel.ZoneIdentifier = core.StringPtr("testString")
				createFilterOptionsModel.Accept = core.StringPtr("testString")
				createFilterOptionsModel.FilterInput = []filtersv1.FilterInput{*filterInputModel}
				createFilterOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = filtersService.CreateFilter(createFilterOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateFilter with error: Operation validation and request error`, func() {
				filtersService, serviceErr := filtersv1.NewFiltersV1(&filtersv1.FiltersV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(filtersService).ToNot(BeNil())

				// Construct an instance of the FilterInput model
				filterInputModel := new(filtersv1.FilterInput)
				filterInputModel.Expression = core.StringPtr(`not http.request.uri.path matches "^/api/.*$"`)
				filterInputModel.Paused = core.BoolPtr(false)
				filterInputModel.Description = core.StringPtr("not /api")

				// Construct an instance of the CreateFilterOptions model
				createFilterOptionsModel := new(filtersv1.CreateFilterOptions)
				createFilterOptionsModel.XAuthUserToken = core.StringPtr("testString")
				createFilterOptionsModel.Crn = core.StringPtr("testString")
				createFilterOptionsModel.ZoneIdentifier = core.StringPtr("testString")
				createFilterOptionsModel.Accept = core.StringPtr("testString")
				createFilterOptionsModel.FilterInput = []filtersv1.FilterInput{*filterInputModel}
				createFilterOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := filtersService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := filtersService.CreateFilter(createFilterOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateFilterOptions model with no property values
				createFilterOptionsModelNew := new(filtersv1.CreateFilterOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = filtersService.CreateFilter(createFilterOptionsModelNew)
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
			filtersService, serviceErr := filtersv1.NewFiltersV1(&filtersv1.FiltersV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(filtersService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			filtersService, serviceErr := filtersv1.NewFiltersV1(&filtersv1.FiltersV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(filtersService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			filtersService, serviceErr := filtersv1.NewFiltersV1(&filtersv1.FiltersV1Options{
				URL: "https://filtersv1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(filtersService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"FILTERS_URL":       "https://filtersv1/api",
				"FILTERS_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				filtersService, serviceErr := filtersv1.NewFiltersV1UsingExternalConfig(&filtersv1.FiltersV1Options{})
				Expect(filtersService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := filtersService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != filtersService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(filtersService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(filtersService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				filtersService, serviceErr := filtersv1.NewFiltersV1UsingExternalConfig(&filtersv1.FiltersV1Options{
					URL: "https://testService/api",
				})
				Expect(filtersService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(filtersService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := filtersService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != filtersService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(filtersService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(filtersService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				filtersService, serviceErr := filtersv1.NewFiltersV1UsingExternalConfig(&filtersv1.FiltersV1Options{})
				err := filtersService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(filtersService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(filtersService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := filtersService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != filtersService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(filtersService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(filtersService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"FILTERS_URL":       "https://filtersv1/api",
				"FILTERS_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			filtersService, serviceErr := filtersv1.NewFiltersV1UsingExternalConfig(&filtersv1.FiltersV1Options{})

			It(`Instantiate service client with error`, func() {
				Expect(filtersService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"FILTERS_AUTH_TYPE": "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			filtersService, serviceErr := filtersv1.NewFiltersV1UsingExternalConfig(&filtersv1.FiltersV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(filtersService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = filtersv1.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})

	Describe(`UpdateFilters(updateFiltersOptions *UpdateFiltersOptions)`, func() {
		updateFiltersPath := "/v1/testString/zones/testString/filters"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateFiltersPath))
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

					Expect(req.Header["X-Auth-User-Token"]).ToNot(BeNil())
					Expect(req.Header["X-Auth-User-Token"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Accept"]).ToNot(BeNil())
					Expect(req.Header["Accept"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "*/*")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `This is a mock binary response.`)
				}))
			})
			It(`Invoke UpdateFilters successfully with retries`, func() {
				filtersService, serviceErr := filtersv1.NewFiltersV1(&filtersv1.FiltersV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(filtersService).ToNot(BeNil())
				filtersService.EnableRetries(0, 0)

				// Construct an instance of the FilterUpdateInput model
				filterUpdateInputModel := new(filtersv1.FilterUpdateInput)
				filterUpdateInputModel.ID = core.StringPtr("f2a64520581a4209aab12187a0081364")
				filterUpdateInputModel.Expression = core.StringPtr(`not http.request.uri.path matches "^/api/.*$"`)
				filterUpdateInputModel.Description = core.StringPtr("not /api")
				filterUpdateInputModel.Paused = core.BoolPtr(false)

				// Construct an instance of the UpdateFiltersOptions model
				updateFiltersOptionsModel := new(filtersv1.UpdateFiltersOptions)
				updateFiltersOptionsModel.XAuthUserToken = core.StringPtr("testString")
				updateFiltersOptionsModel.Crn = core.StringPtr("testString")
				updateFiltersOptionsModel.ZoneIdentifier = core.StringPtr("testString")
				updateFiltersOptionsModel.Accept = core.StringPtr("testString")
				updateFiltersOptionsModel.FilterUpdateInput = []filtersv1.FilterUpdateInput{*filterUpdateInputModel}
				updateFiltersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := filtersService.UpdateFiltersWithContext(ctx, updateFiltersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				filtersService.DisableRetries()
				result, response, operationErr := filtersService.UpdateFilters(updateFiltersOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = filtersService.UpdateFiltersWithContext(ctx, updateFiltersOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(updateFiltersPath))
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

					Expect(req.Header["X-Auth-User-Token"]).ToNot(BeNil())
					Expect(req.Header["X-Auth-User-Token"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Accept"]).ToNot(BeNil())
					Expect(req.Header["Accept"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "*/*")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `This is a mock binary response.`)
				}))
			})
			It(`Invoke UpdateFilters successfully`, func() {
				filtersService, serviceErr := filtersv1.NewFiltersV1(&filtersv1.FiltersV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(filtersService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := filtersService.UpdateFilters(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the FilterUpdateInput model
				filterUpdateInputModel := new(filtersv1.FilterUpdateInput)
				filterUpdateInputModel.ID = core.StringPtr("f2a64520581a4209aab12187a0081364")
				filterUpdateInputModel.Expression = core.StringPtr(`not http.request.uri.path matches "^/api/.*$"`)
				filterUpdateInputModel.Description = core.StringPtr("not /api")
				filterUpdateInputModel.Paused = core.BoolPtr(false)

				// Construct an instance of the UpdateFiltersOptions model
				updateFiltersOptionsModel := new(filtersv1.UpdateFiltersOptions)
				updateFiltersOptionsModel.XAuthUserToken = core.StringPtr("testString")
				updateFiltersOptionsModel.Crn = core.StringPtr("testString")
				updateFiltersOptionsModel.ZoneIdentifier = core.StringPtr("testString")
				updateFiltersOptionsModel.Accept = core.StringPtr("testString")
				updateFiltersOptionsModel.FilterUpdateInput = []filtersv1.FilterUpdateInput{*filterUpdateInputModel}
				updateFiltersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = filtersService.UpdateFilters(updateFiltersOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateFilters with error: Operation validation and request error`, func() {
				filtersService, serviceErr := filtersv1.NewFiltersV1(&filtersv1.FiltersV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(filtersService).ToNot(BeNil())

				// Construct an instance of the FilterUpdateInput model
				filterUpdateInputModel := new(filtersv1.FilterUpdateInput)
				filterUpdateInputModel.ID = core.StringPtr("f2a64520581a4209aab12187a0081364")
				filterUpdateInputModel.Expression = core.StringPtr(`not http.request.uri.path matches "^/api/.*$"`)
				filterUpdateInputModel.Description = core.StringPtr("not /api")
				filterUpdateInputModel.Paused = core.BoolPtr(false)

				// Construct an instance of the UpdateFiltersOptions model
				updateFiltersOptionsModel := new(filtersv1.UpdateFiltersOptions)
				updateFiltersOptionsModel.XAuthUserToken = core.StringPtr("testString")
				updateFiltersOptionsModel.Crn = core.StringPtr("testString")
				updateFiltersOptionsModel.ZoneIdentifier = core.StringPtr("testString")
				updateFiltersOptionsModel.Accept = core.StringPtr("testString")
				updateFiltersOptionsModel.FilterUpdateInput = []filtersv1.FilterUpdateInput{*filterUpdateInputModel}
				updateFiltersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := filtersService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := filtersService.UpdateFilters(updateFiltersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateFiltersOptions model with no property values
				updateFiltersOptionsModelNew := new(filtersv1.UpdateFiltersOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = filtersService.UpdateFilters(updateFiltersOptionsModelNew)
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
			filtersService, serviceErr := filtersv1.NewFiltersV1(&filtersv1.FiltersV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(filtersService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			filtersService, serviceErr := filtersv1.NewFiltersV1(&filtersv1.FiltersV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(filtersService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			filtersService, serviceErr := filtersv1.NewFiltersV1(&filtersv1.FiltersV1Options{
				URL: "https://filtersv1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(filtersService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"FILTERS_URL":       "https://filtersv1/api",
				"FILTERS_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				filtersService, serviceErr := filtersv1.NewFiltersV1UsingExternalConfig(&filtersv1.FiltersV1Options{})
				Expect(filtersService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := filtersService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != filtersService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(filtersService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(filtersService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				filtersService, serviceErr := filtersv1.NewFiltersV1UsingExternalConfig(&filtersv1.FiltersV1Options{
					URL: "https://testService/api",
				})
				Expect(filtersService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(filtersService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := filtersService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != filtersService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(filtersService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(filtersService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				filtersService, serviceErr := filtersv1.NewFiltersV1UsingExternalConfig(&filtersv1.FiltersV1Options{})
				err := filtersService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(filtersService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(filtersService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := filtersService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != filtersService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(filtersService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(filtersService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"FILTERS_URL":       "https://filtersv1/api",
				"FILTERS_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			filtersService, serviceErr := filtersv1.NewFiltersV1UsingExternalConfig(&filtersv1.FiltersV1Options{})

			It(`Instantiate service client with error`, func() {
				Expect(filtersService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"FILTERS_AUTH_TYPE": "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			filtersService, serviceErr := filtersv1.NewFiltersV1UsingExternalConfig(&filtersv1.FiltersV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(filtersService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = filtersv1.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})

	Describe(`DeleteFilters(deleteFiltersOptions *DeleteFiltersOptions)`, func() {
		deleteFiltersPath := "/v1/testString/zones/testString/filters"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteFiltersPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.Header["X-Auth-User-Token"]).ToNot(BeNil())
					Expect(req.Header["X-Auth-User-Token"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Accept"]).ToNot(BeNil())
					Expect(req.Header["Accept"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["id"]).To(Equal([]string{"b7ff25282d394be7b945e23c7106ce8a"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "*/*")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `This is a mock binary response.`)
				}))
			})
			It(`Invoke DeleteFilters successfully with retries`, func() {
				filtersService, serviceErr := filtersv1.NewFiltersV1(&filtersv1.FiltersV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(filtersService).ToNot(BeNil())
				filtersService.EnableRetries(0, 0)

				// Construct an instance of the DeleteFiltersOptions model
				deleteFiltersOptionsModel := new(filtersv1.DeleteFiltersOptions)
				deleteFiltersOptionsModel.XAuthUserToken = core.StringPtr("testString")
				deleteFiltersOptionsModel.Crn = core.StringPtr("testString")
				deleteFiltersOptionsModel.ZoneIdentifier = core.StringPtr("testString")
				deleteFiltersOptionsModel.ID = core.StringPtr("b7ff25282d394be7b945e23c7106ce8a")
				deleteFiltersOptionsModel.Accept = core.StringPtr("testString")
				deleteFiltersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := filtersService.DeleteFiltersWithContext(ctx, deleteFiltersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				filtersService.DisableRetries()
				result, response, operationErr := filtersService.DeleteFilters(deleteFiltersOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = filtersService.DeleteFiltersWithContext(ctx, deleteFiltersOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(deleteFiltersPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.Header["X-Auth-User-Token"]).ToNot(BeNil())
					Expect(req.Header["X-Auth-User-Token"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Accept"]).ToNot(BeNil())
					Expect(req.Header["Accept"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["id"]).To(Equal([]string{"b7ff25282d394be7b945e23c7106ce8a"}))
					// Set mock response
					res.Header().Set("Content-type", "*/*")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `This is a mock binary response.`)
				}))
			})
			It(`Invoke DeleteFilters successfully`, func() {
				filtersService, serviceErr := filtersv1.NewFiltersV1(&filtersv1.FiltersV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(filtersService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := filtersService.DeleteFilters(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DeleteFiltersOptions model
				deleteFiltersOptionsModel := new(filtersv1.DeleteFiltersOptions)
				deleteFiltersOptionsModel.XAuthUserToken = core.StringPtr("testString")
				deleteFiltersOptionsModel.Crn = core.StringPtr("testString")
				deleteFiltersOptionsModel.ZoneIdentifier = core.StringPtr("testString")
				deleteFiltersOptionsModel.ID = core.StringPtr("b7ff25282d394be7b945e23c7106ce8a")
				deleteFiltersOptionsModel.Accept = core.StringPtr("testString")
				deleteFiltersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = filtersService.DeleteFilters(deleteFiltersOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke DeleteFilters with error: Operation validation and request error`, func() {
				filtersService, serviceErr := filtersv1.NewFiltersV1(&filtersv1.FiltersV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(filtersService).ToNot(BeNil())

				// Construct an instance of the DeleteFiltersOptions model
				deleteFiltersOptionsModel := new(filtersv1.DeleteFiltersOptions)
				deleteFiltersOptionsModel.XAuthUserToken = core.StringPtr("testString")
				deleteFiltersOptionsModel.Crn = core.StringPtr("testString")
				deleteFiltersOptionsModel.ZoneIdentifier = core.StringPtr("testString")
				deleteFiltersOptionsModel.ID = core.StringPtr("b7ff25282d394be7b945e23c7106ce8a")
				deleteFiltersOptionsModel.Accept = core.StringPtr("testString")
				deleteFiltersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := filtersService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := filtersService.DeleteFilters(deleteFiltersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the DeleteFiltersOptions model with no property values
				deleteFiltersOptionsModelNew := new(filtersv1.DeleteFiltersOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = filtersService.DeleteFilters(deleteFiltersOptionsModelNew)
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
			filtersService, serviceErr := filtersv1.NewFiltersV1(&filtersv1.FiltersV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(filtersService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			filtersService, serviceErr := filtersv1.NewFiltersV1(&filtersv1.FiltersV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(filtersService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			filtersService, serviceErr := filtersv1.NewFiltersV1(&filtersv1.FiltersV1Options{
				URL: "https://filtersv1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(filtersService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"FILTERS_URL":       "https://filtersv1/api",
				"FILTERS_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				filtersService, serviceErr := filtersv1.NewFiltersV1UsingExternalConfig(&filtersv1.FiltersV1Options{})
				Expect(filtersService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := filtersService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != filtersService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(filtersService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(filtersService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				filtersService, serviceErr := filtersv1.NewFiltersV1UsingExternalConfig(&filtersv1.FiltersV1Options{
					URL: "https://testService/api",
				})
				Expect(filtersService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(filtersService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := filtersService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != filtersService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(filtersService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(filtersService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				filtersService, serviceErr := filtersv1.NewFiltersV1UsingExternalConfig(&filtersv1.FiltersV1Options{})
				err := filtersService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(filtersService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(filtersService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := filtersService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != filtersService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(filtersService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(filtersService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"FILTERS_URL":       "https://filtersv1/api",
				"FILTERS_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			filtersService, serviceErr := filtersv1.NewFiltersV1UsingExternalConfig(&filtersv1.FiltersV1Options{})

			It(`Instantiate service client with error`, func() {
				Expect(filtersService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"FILTERS_AUTH_TYPE": "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			filtersService, serviceErr := filtersv1.NewFiltersV1UsingExternalConfig(&filtersv1.FiltersV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(filtersService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = filtersv1.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})

	Describe(`GetFilter(getFilterOptions *GetFilterOptions)`, func() {
		getFilterPath := "/v1/testString/zones/testString/filters/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getFilterPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Auth-User-Token"]).ToNot(BeNil())
					Expect(req.Header["X-Auth-User-Token"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Accept"]).ToNot(BeNil())
					Expect(req.Header["Accept"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "*/*")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `This is a mock binary response.`)
				}))
			})
			It(`Invoke GetFilter successfully with retries`, func() {
				filtersService, serviceErr := filtersv1.NewFiltersV1(&filtersv1.FiltersV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(filtersService).ToNot(BeNil())
				filtersService.EnableRetries(0, 0)

				// Construct an instance of the GetFilterOptions model
				getFilterOptionsModel := new(filtersv1.GetFilterOptions)
				getFilterOptionsModel.XAuthUserToken = core.StringPtr("testString")
				getFilterOptionsModel.Crn = core.StringPtr("testString")
				getFilterOptionsModel.ZoneIdentifier = core.StringPtr("testString")
				getFilterOptionsModel.FilterIdentifier = core.StringPtr("testString")
				getFilterOptionsModel.Accept = core.StringPtr("testString")
				getFilterOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := filtersService.GetFilterWithContext(ctx, getFilterOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				filtersService.DisableRetries()
				result, response, operationErr := filtersService.GetFilter(getFilterOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = filtersService.GetFilterWithContext(ctx, getFilterOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getFilterPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Auth-User-Token"]).ToNot(BeNil())
					Expect(req.Header["X-Auth-User-Token"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Accept"]).ToNot(BeNil())
					Expect(req.Header["Accept"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "*/*")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `This is a mock binary response.`)
				}))
			})
			It(`Invoke GetFilter successfully`, func() {
				filtersService, serviceErr := filtersv1.NewFiltersV1(&filtersv1.FiltersV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(filtersService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := filtersService.GetFilter(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetFilterOptions model
				getFilterOptionsModel := new(filtersv1.GetFilterOptions)
				getFilterOptionsModel.XAuthUserToken = core.StringPtr("testString")
				getFilterOptionsModel.Crn = core.StringPtr("testString")
				getFilterOptionsModel.ZoneIdentifier = core.StringPtr("testString")
				getFilterOptionsModel.FilterIdentifier = core.StringPtr("testString")
				getFilterOptionsModel.Accept = core.StringPtr("testString")
				getFilterOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = filtersService.GetFilter(getFilterOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetFilter with error: Operation validation and request error`, func() {
				filtersService, serviceErr := filtersv1.NewFiltersV1(&filtersv1.FiltersV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(filtersService).ToNot(BeNil())

				// Construct an instance of the GetFilterOptions model
				getFilterOptionsModel := new(filtersv1.GetFilterOptions)
				getFilterOptionsModel.XAuthUserToken = core.StringPtr("testString")
				getFilterOptionsModel.Crn = core.StringPtr("testString")
				getFilterOptionsModel.ZoneIdentifier = core.StringPtr("testString")
				getFilterOptionsModel.FilterIdentifier = core.StringPtr("testString")
				getFilterOptionsModel.Accept = core.StringPtr("testString")
				getFilterOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := filtersService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := filtersService.GetFilter(getFilterOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetFilterOptions model with no property values
				getFilterOptionsModelNew := new(filtersv1.GetFilterOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = filtersService.GetFilter(getFilterOptionsModelNew)
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
			filtersService, serviceErr := filtersv1.NewFiltersV1(&filtersv1.FiltersV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(filtersService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			filtersService, serviceErr := filtersv1.NewFiltersV1(&filtersv1.FiltersV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(filtersService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			filtersService, serviceErr := filtersv1.NewFiltersV1(&filtersv1.FiltersV1Options{
				URL: "https://filtersv1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(filtersService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"FILTERS_URL":       "https://filtersv1/api",
				"FILTERS_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				filtersService, serviceErr := filtersv1.NewFiltersV1UsingExternalConfig(&filtersv1.FiltersV1Options{})
				Expect(filtersService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := filtersService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != filtersService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(filtersService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(filtersService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				filtersService, serviceErr := filtersv1.NewFiltersV1UsingExternalConfig(&filtersv1.FiltersV1Options{
					URL: "https://testService/api",
				})
				Expect(filtersService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(filtersService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := filtersService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != filtersService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(filtersService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(filtersService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				filtersService, serviceErr := filtersv1.NewFiltersV1UsingExternalConfig(&filtersv1.FiltersV1Options{})
				err := filtersService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(filtersService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(filtersService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := filtersService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != filtersService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(filtersService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(filtersService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"FILTERS_URL":       "https://filtersv1/api",
				"FILTERS_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			filtersService, serviceErr := filtersv1.NewFiltersV1UsingExternalConfig(&filtersv1.FiltersV1Options{})

			It(`Instantiate service client with error`, func() {
				Expect(filtersService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"FILTERS_AUTH_TYPE": "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			filtersService, serviceErr := filtersv1.NewFiltersV1UsingExternalConfig(&filtersv1.FiltersV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(filtersService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = filtersv1.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})

	Describe(`UpdateFilter(updateFilterOptions *UpdateFilterOptions)`, func() {
		updateFilterPath := "/v1/testString/zones/testString/filters/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateFilterPath))
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

					Expect(req.Header["X-Auth-User-Token"]).ToNot(BeNil())
					Expect(req.Header["X-Auth-User-Token"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Accept"]).ToNot(BeNil())
					Expect(req.Header["Accept"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "*/*")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `This is a mock binary response.`)
				}))
			})
			It(`Invoke UpdateFilter successfully with retries`, func() {
				filtersService, serviceErr := filtersv1.NewFiltersV1(&filtersv1.FiltersV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(filtersService).ToNot(BeNil())
				filtersService.EnableRetries(0, 0)

				// Construct an instance of the UpdateFilterOptions model
				updateFilterOptionsModel := new(filtersv1.UpdateFilterOptions)
				updateFilterOptionsModel.XAuthUserToken = core.StringPtr("testString")
				updateFilterOptionsModel.Crn = core.StringPtr("testString")
				updateFilterOptionsModel.ZoneIdentifier = core.StringPtr("testString")
				updateFilterOptionsModel.FilterIdentifier = core.StringPtr("testString")
				updateFilterOptionsModel.Accept = core.StringPtr("testString")
				updateFilterOptionsModel.ID = core.StringPtr("f2a64520581a4209aab12187a0081364")
				updateFilterOptionsModel.Expression = core.StringPtr(`not http.request.uri.path matches "^/api/.*$"`)
				updateFilterOptionsModel.Description = core.StringPtr("not /api")
				updateFilterOptionsModel.Paused = core.BoolPtr(false)
				updateFilterOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := filtersService.UpdateFilterWithContext(ctx, updateFilterOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				filtersService.DisableRetries()
				result, response, operationErr := filtersService.UpdateFilter(updateFilterOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = filtersService.UpdateFilterWithContext(ctx, updateFilterOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(updateFilterPath))
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

					Expect(req.Header["X-Auth-User-Token"]).ToNot(BeNil())
					Expect(req.Header["X-Auth-User-Token"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Accept"]).ToNot(BeNil())
					Expect(req.Header["Accept"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "*/*")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `This is a mock binary response.`)
				}))
			})
			It(`Invoke UpdateFilter successfully`, func() {
				filtersService, serviceErr := filtersv1.NewFiltersV1(&filtersv1.FiltersV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(filtersService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := filtersService.UpdateFilter(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateFilterOptions model
				updateFilterOptionsModel := new(filtersv1.UpdateFilterOptions)
				updateFilterOptionsModel.XAuthUserToken = core.StringPtr("testString")
				updateFilterOptionsModel.Crn = core.StringPtr("testString")
				updateFilterOptionsModel.ZoneIdentifier = core.StringPtr("testString")
				updateFilterOptionsModel.FilterIdentifier = core.StringPtr("testString")
				updateFilterOptionsModel.Accept = core.StringPtr("testString")
				updateFilterOptionsModel.ID = core.StringPtr("f2a64520581a4209aab12187a0081364")
				updateFilterOptionsModel.Expression = core.StringPtr(`not http.request.uri.path matches "^/api/.*$"`)
				updateFilterOptionsModel.Description = core.StringPtr("not /api")
				updateFilterOptionsModel.Paused = core.BoolPtr(false)
				updateFilterOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = filtersService.UpdateFilter(updateFilterOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateFilter with error: Operation validation and request error`, func() {
				filtersService, serviceErr := filtersv1.NewFiltersV1(&filtersv1.FiltersV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(filtersService).ToNot(BeNil())

				// Construct an instance of the UpdateFilterOptions model
				updateFilterOptionsModel := new(filtersv1.UpdateFilterOptions)
				updateFilterOptionsModel.XAuthUserToken = core.StringPtr("testString")
				updateFilterOptionsModel.Crn = core.StringPtr("testString")
				updateFilterOptionsModel.ZoneIdentifier = core.StringPtr("testString")
				updateFilterOptionsModel.FilterIdentifier = core.StringPtr("testString")
				updateFilterOptionsModel.Accept = core.StringPtr("testString")
				updateFilterOptionsModel.ID = core.StringPtr("f2a64520581a4209aab12187a0081364")
				updateFilterOptionsModel.Expression = core.StringPtr(`not http.request.uri.path matches "^/api/.*$"`)
				updateFilterOptionsModel.Description = core.StringPtr("not /api")
				updateFilterOptionsModel.Paused = core.BoolPtr(false)
				updateFilterOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := filtersService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := filtersService.UpdateFilter(updateFilterOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateFilterOptions model with no property values
				updateFilterOptionsModelNew := new(filtersv1.UpdateFilterOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = filtersService.UpdateFilter(updateFilterOptionsModelNew)
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
			filtersService, serviceErr := filtersv1.NewFiltersV1(&filtersv1.FiltersV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(filtersService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			filtersService, serviceErr := filtersv1.NewFiltersV1(&filtersv1.FiltersV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(filtersService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			filtersService, serviceErr := filtersv1.NewFiltersV1(&filtersv1.FiltersV1Options{
				URL: "https://filtersv1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(filtersService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"FILTERS_URL":       "https://filtersv1/api",
				"FILTERS_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				filtersService, serviceErr := filtersv1.NewFiltersV1UsingExternalConfig(&filtersv1.FiltersV1Options{})
				Expect(filtersService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := filtersService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != filtersService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(filtersService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(filtersService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				filtersService, serviceErr := filtersv1.NewFiltersV1UsingExternalConfig(&filtersv1.FiltersV1Options{
					URL: "https://testService/api",
				})
				Expect(filtersService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(filtersService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := filtersService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != filtersService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(filtersService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(filtersService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				filtersService, serviceErr := filtersv1.NewFiltersV1UsingExternalConfig(&filtersv1.FiltersV1Options{})
				err := filtersService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(filtersService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(filtersService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := filtersService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != filtersService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(filtersService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(filtersService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"FILTERS_URL":       "https://filtersv1/api",
				"FILTERS_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			filtersService, serviceErr := filtersv1.NewFiltersV1UsingExternalConfig(&filtersv1.FiltersV1Options{})

			It(`Instantiate service client with error`, func() {
				Expect(filtersService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"FILTERS_AUTH_TYPE": "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			filtersService, serviceErr := filtersv1.NewFiltersV1UsingExternalConfig(&filtersv1.FiltersV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(filtersService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = filtersv1.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})

	Describe(`DeleteFilter(deleteFilterOptions *DeleteFilterOptions)`, func() {
		deleteFilterPath := "/v1/testString/zones/testString/filters/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteFilterPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.Header["X-Auth-User-Token"]).ToNot(BeNil())
					Expect(req.Header["X-Auth-User-Token"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Accept"]).ToNot(BeNil())
					Expect(req.Header["Accept"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "*/*")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `This is a mock binary response.`)
				}))
			})
			It(`Invoke DeleteFilter successfully with retries`, func() {
				filtersService, serviceErr := filtersv1.NewFiltersV1(&filtersv1.FiltersV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(filtersService).ToNot(BeNil())
				filtersService.EnableRetries(0, 0)

				// Construct an instance of the DeleteFilterOptions model
				deleteFilterOptionsModel := new(filtersv1.DeleteFilterOptions)
				deleteFilterOptionsModel.XAuthUserToken = core.StringPtr("testString")
				deleteFilterOptionsModel.Crn = core.StringPtr("testString")
				deleteFilterOptionsModel.ZoneIdentifier = core.StringPtr("testString")
				deleteFilterOptionsModel.FilterIdentifier = core.StringPtr("testString")
				deleteFilterOptionsModel.Accept = core.StringPtr("testString")
				deleteFilterOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := filtersService.DeleteFilterWithContext(ctx, deleteFilterOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				filtersService.DisableRetries()
				result, response, operationErr := filtersService.DeleteFilter(deleteFilterOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = filtersService.DeleteFilterWithContext(ctx, deleteFilterOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(deleteFilterPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.Header["X-Auth-User-Token"]).ToNot(BeNil())
					Expect(req.Header["X-Auth-User-Token"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Accept"]).ToNot(BeNil())
					Expect(req.Header["Accept"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "*/*")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `This is a mock binary response.`)
				}))
			})
			It(`Invoke DeleteFilter successfully`, func() {
				filtersService, serviceErr := filtersv1.NewFiltersV1(&filtersv1.FiltersV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(filtersService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := filtersService.DeleteFilter(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DeleteFilterOptions model
				deleteFilterOptionsModel := new(filtersv1.DeleteFilterOptions)
				deleteFilterOptionsModel.XAuthUserToken = core.StringPtr("testString")
				deleteFilterOptionsModel.Crn = core.StringPtr("testString")
				deleteFilterOptionsModel.ZoneIdentifier = core.StringPtr("testString")
				deleteFilterOptionsModel.FilterIdentifier = core.StringPtr("testString")
				deleteFilterOptionsModel.Accept = core.StringPtr("testString")
				deleteFilterOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = filtersService.DeleteFilter(deleteFilterOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke DeleteFilter with error: Operation validation and request error`, func() {
				filtersService, serviceErr := filtersv1.NewFiltersV1(&filtersv1.FiltersV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(filtersService).ToNot(BeNil())

				// Construct an instance of the DeleteFilterOptions model
				deleteFilterOptionsModel := new(filtersv1.DeleteFilterOptions)
				deleteFilterOptionsModel.XAuthUserToken = core.StringPtr("testString")
				deleteFilterOptionsModel.Crn = core.StringPtr("testString")
				deleteFilterOptionsModel.ZoneIdentifier = core.StringPtr("testString")
				deleteFilterOptionsModel.FilterIdentifier = core.StringPtr("testString")
				deleteFilterOptionsModel.Accept = core.StringPtr("testString")
				deleteFilterOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := filtersService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := filtersService.DeleteFilter(deleteFilterOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the DeleteFilterOptions model with no property values
				deleteFilterOptionsModelNew := new(filtersv1.DeleteFilterOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = filtersService.DeleteFilter(deleteFilterOptionsModelNew)
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
			filtersService, _ := filtersv1.NewFiltersV1(&filtersv1.FiltersV1Options{
				URL:           "http://filtersv1modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
			})
			It(`Invoke NewCreateFilterOptions successfully`, func() {
				// Construct an instance of the FilterInput model
				filterInputModel := new(filtersv1.FilterInput)
				Expect(filterInputModel).ToNot(BeNil())
				filterInputModel.Expression = core.StringPtr(`not http.request.uri.path matches "^/api/.*$"`)
				filterInputModel.Paused = core.BoolPtr(false)
				filterInputModel.Description = core.StringPtr("not /api")
				Expect(filterInputModel.Expression).To(Equal(core.StringPtr(`not http.request.uri.path matches "^/api/.*$"`)))
				Expect(filterInputModel.Paused).To(Equal(core.BoolPtr(false)))
				Expect(filterInputModel.Description).To(Equal(core.StringPtr("not /api")))

				// Construct an instance of the CreateFilterOptions model
				xAuthUserToken := "testString"
				crn := "testString"
				zoneIdentifier := "testString"
				createFilterOptionsModel := filtersService.NewCreateFilterOptions(xAuthUserToken, crn, zoneIdentifier)
				createFilterOptionsModel.SetXAuthUserToken("testString")
				createFilterOptionsModel.SetCrn("testString")
				createFilterOptionsModel.SetZoneIdentifier("testString")
				createFilterOptionsModel.SetAccept("testString")
				createFilterOptionsModel.SetFilterInput([]filtersv1.FilterInput{*filterInputModel})
				createFilterOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createFilterOptionsModel).ToNot(BeNil())
				Expect(createFilterOptionsModel.XAuthUserToken).To(Equal(core.StringPtr("testString")))
				Expect(createFilterOptionsModel.Crn).To(Equal(core.StringPtr("testString")))
				Expect(createFilterOptionsModel.ZoneIdentifier).To(Equal(core.StringPtr("testString")))
				Expect(createFilterOptionsModel.Accept).To(Equal(core.StringPtr("testString")))
				Expect(createFilterOptionsModel.FilterInput).To(Equal([]filtersv1.FilterInput{*filterInputModel}))
				Expect(createFilterOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteFilterOptions successfully`, func() {
				// Construct an instance of the DeleteFilterOptions model
				xAuthUserToken := "testString"
				crn := "testString"
				zoneIdentifier := "testString"
				filterIdentifier := "testString"
				deleteFilterOptionsModel := filtersService.NewDeleteFilterOptions(xAuthUserToken, crn, zoneIdentifier, filterIdentifier)
				deleteFilterOptionsModel.SetXAuthUserToken("testString")
				deleteFilterOptionsModel.SetCrn("testString")
				deleteFilterOptionsModel.SetZoneIdentifier("testString")
				deleteFilterOptionsModel.SetFilterIdentifier("testString")
				deleteFilterOptionsModel.SetAccept("testString")
				deleteFilterOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteFilterOptionsModel).ToNot(BeNil())
				Expect(deleteFilterOptionsModel.XAuthUserToken).To(Equal(core.StringPtr("testString")))
				Expect(deleteFilterOptionsModel.Crn).To(Equal(core.StringPtr("testString")))
				Expect(deleteFilterOptionsModel.ZoneIdentifier).To(Equal(core.StringPtr("testString")))
				Expect(deleteFilterOptionsModel.FilterIdentifier).To(Equal(core.StringPtr("testString")))
				Expect(deleteFilterOptionsModel.Accept).To(Equal(core.StringPtr("testString")))
				Expect(deleteFilterOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteFiltersOptions successfully`, func() {
				// Construct an instance of the DeleteFiltersOptions model
				xAuthUserToken := "testString"
				crn := "testString"
				zoneIdentifier := "testString"
				id := "b7ff25282d394be7b945e23c7106ce8a"
				deleteFiltersOptionsModel := filtersService.NewDeleteFiltersOptions(xAuthUserToken, crn, zoneIdentifier, id)
				deleteFiltersOptionsModel.SetXAuthUserToken("testString")
				deleteFiltersOptionsModel.SetCrn("testString")
				deleteFiltersOptionsModel.SetZoneIdentifier("testString")
				deleteFiltersOptionsModel.SetID("b7ff25282d394be7b945e23c7106ce8a")
				deleteFiltersOptionsModel.SetAccept("testString")
				deleteFiltersOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteFiltersOptionsModel).ToNot(BeNil())
				Expect(deleteFiltersOptionsModel.XAuthUserToken).To(Equal(core.StringPtr("testString")))
				Expect(deleteFiltersOptionsModel.Crn).To(Equal(core.StringPtr("testString")))
				Expect(deleteFiltersOptionsModel.ZoneIdentifier).To(Equal(core.StringPtr("testString")))
				Expect(deleteFiltersOptionsModel.ID).To(Equal(core.StringPtr("b7ff25282d394be7b945e23c7106ce8a")))
				Expect(deleteFiltersOptionsModel.Accept).To(Equal(core.StringPtr("testString")))
				Expect(deleteFiltersOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetFilterOptions successfully`, func() {
				// Construct an instance of the GetFilterOptions model
				xAuthUserToken := "testString"
				crn := "testString"
				zoneIdentifier := "testString"
				filterIdentifier := "testString"
				getFilterOptionsModel := filtersService.NewGetFilterOptions(xAuthUserToken, crn, zoneIdentifier, filterIdentifier)
				getFilterOptionsModel.SetXAuthUserToken("testString")
				getFilterOptionsModel.SetCrn("testString")
				getFilterOptionsModel.SetZoneIdentifier("testString")
				getFilterOptionsModel.SetFilterIdentifier("testString")
				getFilterOptionsModel.SetAccept("testString")
				getFilterOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getFilterOptionsModel).ToNot(BeNil())
				Expect(getFilterOptionsModel.XAuthUserToken).To(Equal(core.StringPtr("testString")))
				Expect(getFilterOptionsModel.Crn).To(Equal(core.StringPtr("testString")))
				Expect(getFilterOptionsModel.ZoneIdentifier).To(Equal(core.StringPtr("testString")))
				Expect(getFilterOptionsModel.FilterIdentifier).To(Equal(core.StringPtr("testString")))
				Expect(getFilterOptionsModel.Accept).To(Equal(core.StringPtr("testString")))
				Expect(getFilterOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListAllFiltersOptions successfully`, func() {
				// Construct an instance of the ListAllFiltersOptions model
				xAuthUserToken := "testString"
				crn := "testString"
				zoneIdentifier := "testString"
				listAllFiltersOptionsModel := filtersService.NewListAllFiltersOptions(xAuthUserToken, crn, zoneIdentifier)
				listAllFiltersOptionsModel.SetXAuthUserToken("testString")
				listAllFiltersOptionsModel.SetCrn("testString")
				listAllFiltersOptionsModel.SetZoneIdentifier("testString")
				listAllFiltersOptionsModel.SetAccept("testString")
				listAllFiltersOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listAllFiltersOptionsModel).ToNot(BeNil())
				Expect(listAllFiltersOptionsModel.XAuthUserToken).To(Equal(core.StringPtr("testString")))
				Expect(listAllFiltersOptionsModel.Crn).To(Equal(core.StringPtr("testString")))
				Expect(listAllFiltersOptionsModel.ZoneIdentifier).To(Equal(core.StringPtr("testString")))
				Expect(listAllFiltersOptionsModel.Accept).To(Equal(core.StringPtr("testString")))
				Expect(listAllFiltersOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateFilterOptions successfully`, func() {
				// Construct an instance of the UpdateFilterOptions model
				xAuthUserToken := "testString"
				crn := "testString"
				zoneIdentifier := "testString"
				filterIdentifier := "testString"
				updateFilterOptionsModel := filtersService.NewUpdateFilterOptions(xAuthUserToken, crn, zoneIdentifier, filterIdentifier)
				updateFilterOptionsModel.SetXAuthUserToken("testString")
				updateFilterOptionsModel.SetCrn("testString")
				updateFilterOptionsModel.SetZoneIdentifier("testString")
				updateFilterOptionsModel.SetFilterIdentifier("testString")
				updateFilterOptionsModel.SetAccept("testString")
				updateFilterOptionsModel.SetID("f2a64520581a4209aab12187a0081364")
				updateFilterOptionsModel.SetExpression(`not http.request.uri.path matches "^/api/.*$"`)
				updateFilterOptionsModel.SetDescription("not /api")
				updateFilterOptionsModel.SetPaused(false)
				updateFilterOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateFilterOptionsModel).ToNot(BeNil())
				Expect(updateFilterOptionsModel.XAuthUserToken).To(Equal(core.StringPtr("testString")))
				Expect(updateFilterOptionsModel.Crn).To(Equal(core.StringPtr("testString")))
				Expect(updateFilterOptionsModel.ZoneIdentifier).To(Equal(core.StringPtr("testString")))
				Expect(updateFilterOptionsModel.FilterIdentifier).To(Equal(core.StringPtr("testString")))
				Expect(updateFilterOptionsModel.Accept).To(Equal(core.StringPtr("testString")))
				Expect(updateFilterOptionsModel.ID).To(Equal(core.StringPtr("f2a64520581a4209aab12187a0081364")))
				Expect(updateFilterOptionsModel.Expression).To(Equal(core.StringPtr(`not http.request.uri.path matches "^/api/.*$"`)))
				Expect(updateFilterOptionsModel.Description).To(Equal(core.StringPtr("not /api")))
				Expect(updateFilterOptionsModel.Paused).To(Equal(core.BoolPtr(false)))
				Expect(updateFilterOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateFiltersOptions successfully`, func() {
				// Construct an instance of the FilterUpdateInput model
				filterUpdateInputModel := new(filtersv1.FilterUpdateInput)
				Expect(filterUpdateInputModel).ToNot(BeNil())
				filterUpdateInputModel.ID = core.StringPtr("f2a64520581a4209aab12187a0081364")
				filterUpdateInputModel.Expression = core.StringPtr(`not http.request.uri.path matches "^/api/.*$"`)
				filterUpdateInputModel.Description = core.StringPtr("not /api")
				filterUpdateInputModel.Paused = core.BoolPtr(false)
				Expect(filterUpdateInputModel.ID).To(Equal(core.StringPtr("f2a64520581a4209aab12187a0081364")))
				Expect(filterUpdateInputModel.Expression).To(Equal(core.StringPtr(`not http.request.uri.path matches "^/api/.*$"`)))
				Expect(filterUpdateInputModel.Description).To(Equal(core.StringPtr("not /api")))
				Expect(filterUpdateInputModel.Paused).To(Equal(core.BoolPtr(false)))

				// Construct an instance of the UpdateFiltersOptions model
				xAuthUserToken := "testString"
				crn := "testString"
				zoneIdentifier := "testString"
				updateFiltersOptionsModel := filtersService.NewUpdateFiltersOptions(xAuthUserToken, crn, zoneIdentifier)
				updateFiltersOptionsModel.SetXAuthUserToken("testString")
				updateFiltersOptionsModel.SetCrn("testString")
				updateFiltersOptionsModel.SetZoneIdentifier("testString")
				updateFiltersOptionsModel.SetAccept("testString")
				updateFiltersOptionsModel.SetFilterUpdateInput([]filtersv1.FilterUpdateInput{*filterUpdateInputModel})
				updateFiltersOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateFiltersOptionsModel).ToNot(BeNil())
				Expect(updateFiltersOptionsModel.XAuthUserToken).To(Equal(core.StringPtr("testString")))
				Expect(updateFiltersOptionsModel.Crn).To(Equal(core.StringPtr("testString")))
				Expect(updateFiltersOptionsModel.ZoneIdentifier).To(Equal(core.StringPtr("testString")))
				Expect(updateFiltersOptionsModel.Accept).To(Equal(core.StringPtr("testString")))
				Expect(updateFiltersOptionsModel.FilterUpdateInput).To(Equal([]filtersv1.FilterUpdateInput{*filterUpdateInputModel}))
				Expect(updateFiltersOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewFilterInput successfully`, func() {
				expression := `not http.request.uri.path matches "^/api/.*$"`
				model, err := filtersService.NewFilterInput(expression)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewFilterUpdateInput successfully`, func() {
				id := "f2a64520581a4209aab12187a0081364"
				expression := `not http.request.uri.path matches "^/api/.*$"`
				model, err := filtersService.NewFilterUpdateInput(id, expression)
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
