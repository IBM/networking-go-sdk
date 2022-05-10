/**
 * (C) Copyright IBM Corp. 2022.
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

package mtlsv1_test

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
	"github.com/IBM/networking-go-sdk/mtlsv1"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe(`MtlsV1`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		crn := "testString"
		It(`Instantiate service client`, func() {
			mtlsService, serviceErr := mtlsv1.NewMtlsV1(&mtlsv1.MtlsV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
				Crn:           core.StringPtr(crn),
			})
			Expect(mtlsService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			mtlsService, serviceErr := mtlsv1.NewMtlsV1(&mtlsv1.MtlsV1Options{
				URL: "{BAD_URL_STRING",
				Crn: core.StringPtr(crn),
			})
			Expect(mtlsService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			mtlsService, serviceErr := mtlsv1.NewMtlsV1(&mtlsv1.MtlsV1Options{
				URL: "https://mtlsv1/api",
				Crn: core.StringPtr(crn),
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(mtlsService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Validation Error`, func() {
			mtlsService, serviceErr := mtlsv1.NewMtlsV1(&mtlsv1.MtlsV1Options{})
			Expect(mtlsService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		crn := "testString"
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"MTLS_URL":       "https://mtlsv1/api",
				"MTLS_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				mtlsService, serviceErr := mtlsv1.NewMtlsV1UsingExternalConfig(&mtlsv1.MtlsV1Options{
					Crn: core.StringPtr(crn),
				})
				Expect(mtlsService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := mtlsService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != mtlsService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(mtlsService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(mtlsService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				mtlsService, serviceErr := mtlsv1.NewMtlsV1UsingExternalConfig(&mtlsv1.MtlsV1Options{
					URL: "https://testService/api",
					Crn: core.StringPtr(crn),
				})
				Expect(mtlsService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(mtlsService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := mtlsService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != mtlsService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(mtlsService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(mtlsService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				mtlsService, serviceErr := mtlsv1.NewMtlsV1UsingExternalConfig(&mtlsv1.MtlsV1Options{
					Crn: core.StringPtr(crn),
				})
				err := mtlsService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(mtlsService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(mtlsService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := mtlsService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != mtlsService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(mtlsService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(mtlsService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"MTLS_URL":       "https://mtlsv1/api",
				"MTLS_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			mtlsService, serviceErr := mtlsv1.NewMtlsV1UsingExternalConfig(&mtlsv1.MtlsV1Options{
				Crn: core.StringPtr(crn),
			})

			It(`Instantiate service client with error`, func() {
				Expect(mtlsService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"MTLS_AUTH_TYPE": "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			mtlsService, serviceErr := mtlsv1.NewMtlsV1UsingExternalConfig(&mtlsv1.MtlsV1Options{
				URL: "{BAD_URL_STRING",
				Crn: core.StringPtr(crn),
			})

			It(`Instantiate service client with error`, func() {
				Expect(mtlsService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = mtlsv1.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`ListAccessCertificates(listAccessCertificatesOptions *ListAccessCertificatesOptions) - Operation response error`, func() {
		crn := "testString"
		listAccessCertificatesPath := "/v1/testString/zones/testString/access/certificates"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listAccessCertificatesPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListAccessCertificates with error: Operation response processing error`, func() {
				mtlsService, serviceErr := mtlsv1.NewMtlsV1(&mtlsv1.MtlsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(mtlsService).ToNot(BeNil())

				// Construct an instance of the ListAccessCertificatesOptions model
				listAccessCertificatesOptionsModel := new(mtlsv1.ListAccessCertificatesOptions)
				listAccessCertificatesOptionsModel.ZoneID = core.StringPtr("testString")
				listAccessCertificatesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := mtlsService.ListAccessCertificates(listAccessCertificatesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				mtlsService.EnableRetries(0, 0)
				result, response, operationErr = mtlsService.ListAccessCertificates(listAccessCertificatesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListAccessCertificates(listAccessCertificatesOptions *ListAccessCertificatesOptions)`, func() {
		crn := "testString"
		listAccessCertificatesPath := "/v1/testString/zones/testString/access/certificates"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listAccessCertificatesPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": [{"id": "21a41336-9001-42c4-8440-c79e0cb86e1f", "name": "test-cert", "fingerprint": "MD5 Fingerprint=38:38:B4:FB:3C:33:CE:2C:8E:8E:D1:1B:94:70:C1:5F", "associated_hostnames": ["test.example.com"], "created_at": "2021-04-19T11:09:11Z", "updated_at": "2021-04-19T11:09:11Z", "expires_on": "2026-04-18T06:26:00Z"}]}`)
				}))
			})
			It(`Invoke ListAccessCertificates successfully with retries`, func() {
				mtlsService, serviceErr := mtlsv1.NewMtlsV1(&mtlsv1.MtlsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(mtlsService).ToNot(BeNil())
				mtlsService.EnableRetries(0, 0)

				// Construct an instance of the ListAccessCertificatesOptions model
				listAccessCertificatesOptionsModel := new(mtlsv1.ListAccessCertificatesOptions)
				listAccessCertificatesOptionsModel.ZoneID = core.StringPtr("testString")
				listAccessCertificatesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := mtlsService.ListAccessCertificatesWithContext(ctx, listAccessCertificatesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				mtlsService.DisableRetries()
				result, response, operationErr := mtlsService.ListAccessCertificates(listAccessCertificatesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = mtlsService.ListAccessCertificatesWithContext(ctx, listAccessCertificatesOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listAccessCertificatesPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": [{"id": "21a41336-9001-42c4-8440-c79e0cb86e1f", "name": "test-cert", "fingerprint": "MD5 Fingerprint=38:38:B4:FB:3C:33:CE:2C:8E:8E:D1:1B:94:70:C1:5F", "associated_hostnames": ["test.example.com"], "created_at": "2021-04-19T11:09:11Z", "updated_at": "2021-04-19T11:09:11Z", "expires_on": "2026-04-18T06:26:00Z"}]}`)
				}))
			})
			It(`Invoke ListAccessCertificates successfully`, func() {
				mtlsService, serviceErr := mtlsv1.NewMtlsV1(&mtlsv1.MtlsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(mtlsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := mtlsService.ListAccessCertificates(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListAccessCertificatesOptions model
				listAccessCertificatesOptionsModel := new(mtlsv1.ListAccessCertificatesOptions)
				listAccessCertificatesOptionsModel.ZoneID = core.StringPtr("testString")
				listAccessCertificatesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = mtlsService.ListAccessCertificates(listAccessCertificatesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListAccessCertificates with error: Operation validation and request error`, func() {
				mtlsService, serviceErr := mtlsv1.NewMtlsV1(&mtlsv1.MtlsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(mtlsService).ToNot(BeNil())

				// Construct an instance of the ListAccessCertificatesOptions model
				listAccessCertificatesOptionsModel := new(mtlsv1.ListAccessCertificatesOptions)
				listAccessCertificatesOptionsModel.ZoneID = core.StringPtr("testString")
				listAccessCertificatesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := mtlsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := mtlsService.ListAccessCertificates(listAccessCertificatesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListAccessCertificatesOptions model with no property values
				listAccessCertificatesOptionsModelNew := new(mtlsv1.ListAccessCertificatesOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = mtlsService.ListAccessCertificates(listAccessCertificatesOptionsModelNew)
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
			It(`Invoke ListAccessCertificates successfully`, func() {
				mtlsService, serviceErr := mtlsv1.NewMtlsV1(&mtlsv1.MtlsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(mtlsService).ToNot(BeNil())

				// Construct an instance of the ListAccessCertificatesOptions model
				listAccessCertificatesOptionsModel := new(mtlsv1.ListAccessCertificatesOptions)
				listAccessCertificatesOptionsModel.ZoneID = core.StringPtr("testString")
				listAccessCertificatesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := mtlsService.ListAccessCertificates(listAccessCertificatesOptionsModel)
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
	Describe(`CreateAccessCertificate(createAccessCertificateOptions *CreateAccessCertificateOptions) - Operation response error`, func() {
		crn := "testString"
		createAccessCertificatePath := "/v1/testString/zones/testString/access/certificates"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createAccessCertificatePath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateAccessCertificate with error: Operation response processing error`, func() {
				mtlsService, serviceErr := mtlsv1.NewMtlsV1(&mtlsv1.MtlsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(mtlsService).ToNot(BeNil())

				// Construct an instance of the CreateAccessCertificateOptions model
				createAccessCertificateOptionsModel := new(mtlsv1.CreateAccessCertificateOptions)
				createAccessCertificateOptionsModel.ZoneID = core.StringPtr("testString")
				createAccessCertificateOptionsModel.Name = core.StringPtr("test-cert")
				createAccessCertificateOptionsModel.Certificate = core.StringPtr("-----BEGIN CERTIFICATE-----MIIGAjCCA+qgAwIBAgIJAI7kymlF7CWT...N4RI7KKB7nikiuUf8vhULKy5IX1DrUtmu/B-----END CERTIFICATE-----")
				createAccessCertificateOptionsModel.AssociatedHostnames = []string{"test.example.com"}
				createAccessCertificateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := mtlsService.CreateAccessCertificate(createAccessCertificateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				mtlsService.EnableRetries(0, 0)
				result, response, operationErr = mtlsService.CreateAccessCertificate(createAccessCertificateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateAccessCertificate(createAccessCertificateOptions *CreateAccessCertificateOptions)`, func() {
		crn := "testString"
		createAccessCertificatePath := "/v1/testString/zones/testString/access/certificates"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createAccessCertificatePath))
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
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": "21a41336-9001-42c4-8440-c79e0cb86e1f", "name": "test-cert", "fingerprint": "MD5 Fingerprint=38:38:B4:FB:3C:33:CE:2C:8E:8E:D1:1B:94:70:C1:5F", "associated_hostnames": ["test.example.com"], "created_at": "2021-04-19T11:09:11Z", "updated_at": "2021-04-19T11:09:11Z", "expires_on": "2026-04-18T06:26:00Z"}}`)
				}))
			})
			It(`Invoke CreateAccessCertificate successfully with retries`, func() {
				mtlsService, serviceErr := mtlsv1.NewMtlsV1(&mtlsv1.MtlsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(mtlsService).ToNot(BeNil())
				mtlsService.EnableRetries(0, 0)

				// Construct an instance of the CreateAccessCertificateOptions model
				createAccessCertificateOptionsModel := new(mtlsv1.CreateAccessCertificateOptions)
				createAccessCertificateOptionsModel.ZoneID = core.StringPtr("testString")
				createAccessCertificateOptionsModel.Name = core.StringPtr("test-cert")
				createAccessCertificateOptionsModel.Certificate = core.StringPtr("-----BEGIN CERTIFICATE-----MIIGAjCCA+qgAwIBAgIJAI7kymlF7CWT...N4RI7KKB7nikiuUf8vhULKy5IX10DrUtmu/B-----END CERTIFICATE-----")
				createAccessCertificateOptionsModel.AssociatedHostnames = []string{"test.example.com"}
				createAccessCertificateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := mtlsService.CreateAccessCertificateWithContext(ctx, createAccessCertificateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				mtlsService.DisableRetries()
				result, response, operationErr := mtlsService.CreateAccessCertificate(createAccessCertificateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = mtlsService.CreateAccessCertificateWithContext(ctx, createAccessCertificateOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createAccessCertificatePath))
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
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": "21a41336-9001-42c4-8440-c79e0cb86e1f", "name": "test-cert", "fingerprint": "MD5 Fingerprint=38:38:B4:FB:3C:33:CE:2C:8E:8E:D1:1B:94:70:C1:5F", "associated_hostnames": ["test.example.com"], "created_at": "2021-04-19T11:09:11Z", "updated_at": "2021-04-19T11:09:11Z", "expires_on": "2026-04-18T06:26:00Z"}}`)
				}))
			})
			It(`Invoke CreateAccessCertificate successfully`, func() {
				mtlsService, serviceErr := mtlsv1.NewMtlsV1(&mtlsv1.MtlsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(mtlsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := mtlsService.CreateAccessCertificate(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateAccessCertificateOptions model
				createAccessCertificateOptionsModel := new(mtlsv1.CreateAccessCertificateOptions)
				createAccessCertificateOptionsModel.ZoneID = core.StringPtr("testString")
				createAccessCertificateOptionsModel.Name = core.StringPtr("test-cert")
				createAccessCertificateOptionsModel.Certificate = core.StringPtr("-----BEGIN CERTIFICATE-----MIIGAjCCA+qgAwIBAgIJAI7kymlF7CWT...N4RI7KKB7nikiuUf8vhULKy5IX10DrUtmu/B-----END CERTIFICATE-----")
				createAccessCertificateOptionsModel.AssociatedHostnames = []string{"test.example.com"}
				createAccessCertificateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = mtlsService.CreateAccessCertificate(createAccessCertificateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateAccessCertificate with error: Operation validation and request error`, func() {
				mtlsService, serviceErr := mtlsv1.NewMtlsV1(&mtlsv1.MtlsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(mtlsService).ToNot(BeNil())

				// Construct an instance of the CreateAccessCertificateOptions model
				createAccessCertificateOptionsModel := new(mtlsv1.CreateAccessCertificateOptions)
				createAccessCertificateOptionsModel.ZoneID = core.StringPtr("testString")
				createAccessCertificateOptionsModel.Name = core.StringPtr("test-cert")
				createAccessCertificateOptionsModel.Certificate = core.StringPtr("-----BEGIN CERTIFICATE-----MIIGAjCCA+qgAwIBAgIJAI7kymlF7CWT...N4RI7KKB7nikiuUf8vhULKy5IX10DrUtmu/B-----END CERTIFICATE-----")
				createAccessCertificateOptionsModel.AssociatedHostnames = []string{"test.example.com"}
				createAccessCertificateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := mtlsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := mtlsService.CreateAccessCertificate(createAccessCertificateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateAccessCertificateOptions model with no property values
				createAccessCertificateOptionsModelNew := new(mtlsv1.CreateAccessCertificateOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = mtlsService.CreateAccessCertificate(createAccessCertificateOptionsModelNew)
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
			It(`Invoke CreateAccessCertificate successfully`, func() {
				mtlsService, serviceErr := mtlsv1.NewMtlsV1(&mtlsv1.MtlsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(mtlsService).ToNot(BeNil())

				// Construct an instance of the CreateAccessCertificateOptions model
				createAccessCertificateOptionsModel := new(mtlsv1.CreateAccessCertificateOptions)
				createAccessCertificateOptionsModel.ZoneID = core.StringPtr("testString")
				createAccessCertificateOptionsModel.Name = core.StringPtr("test-cert")
				createAccessCertificateOptionsModel.Certificate = core.StringPtr("-----BEGIN CERTIFICATE-----MIIGAjCCA+qgAwIBAgIJAI7kymlF7CWT...N4RI7KKB7nikiuUf8vhULKy5IX10DrUtmu/B-----END CERTIFICATE-----")
				createAccessCertificateOptionsModel.AssociatedHostnames = []string{"test.example.com"}
				createAccessCertificateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := mtlsService.CreateAccessCertificate(createAccessCertificateOptionsModel)
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
	Describe(`GetAccessCertificate(getAccessCertificateOptions *GetAccessCertificateOptions) - Operation response error`, func() {
		crn := "testString"
		getAccessCertificatePath := "/v1/testString/zones/testString/access/certificates/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getAccessCertificatePath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetAccessCertificate with error: Operation response processing error`, func() {
				mtlsService, serviceErr := mtlsv1.NewMtlsV1(&mtlsv1.MtlsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(mtlsService).ToNot(BeNil())

				// Construct an instance of the GetAccessCertificateOptions model
				getAccessCertificateOptionsModel := new(mtlsv1.GetAccessCertificateOptions)
				getAccessCertificateOptionsModel.ZoneID = core.StringPtr("testString")
				getAccessCertificateOptionsModel.CertID = core.StringPtr("testString")
				getAccessCertificateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := mtlsService.GetAccessCertificate(getAccessCertificateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				mtlsService.EnableRetries(0, 0)
				result, response, operationErr = mtlsService.GetAccessCertificate(getAccessCertificateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetAccessCertificate(getAccessCertificateOptions *GetAccessCertificateOptions)`, func() {
		crn := "testString"
		getAccessCertificatePath := "/v1/testString/zones/testString/access/certificates/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getAccessCertificatePath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": "21a41336-9001-42c4-8440-c79e0cb86e1f", "name": "test-cert", "fingerprint": "MD5 Fingerprint=38:38:B4:FB:3C:33:CE:2C:8E:8E:D1:1B:94:70:C1:5F", "associated_hostnames": ["test.example.com"], "created_at": "2021-04-19T11:09:11Z", "updated_at": "2021-04-19T11:09:11Z", "expires_on": "2026-04-18T06:26:00Z"}}`)
				}))
			})
			It(`Invoke GetAccessCertificate successfully with retries`, func() {
				mtlsService, serviceErr := mtlsv1.NewMtlsV1(&mtlsv1.MtlsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(mtlsService).ToNot(BeNil())
				mtlsService.EnableRetries(0, 0)

				// Construct an instance of the GetAccessCertificateOptions model
				getAccessCertificateOptionsModel := new(mtlsv1.GetAccessCertificateOptions)
				getAccessCertificateOptionsModel.ZoneID = core.StringPtr("testString")
				getAccessCertificateOptionsModel.CertID = core.StringPtr("testString")
				getAccessCertificateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := mtlsService.GetAccessCertificateWithContext(ctx, getAccessCertificateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				mtlsService.DisableRetries()
				result, response, operationErr := mtlsService.GetAccessCertificate(getAccessCertificateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = mtlsService.GetAccessCertificateWithContext(ctx, getAccessCertificateOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getAccessCertificatePath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": "21a41336-9001-42c4-8440-c79e0cb86e1f", "name": "test-cert", "fingerprint": "MD5 Fingerprint=38:38:B4:FB:3C:33:CE:2C:8E:8E:D1:1B:94:70:C1:5F", "associated_hostnames": ["test.example.com"], "created_at": "2021-04-19T11:09:11Z", "updated_at": "2021-04-19T11:09:11Z", "expires_on": "2026-04-18T06:26:00Z"}}`)
				}))
			})
			It(`Invoke GetAccessCertificate successfully`, func() {
				mtlsService, serviceErr := mtlsv1.NewMtlsV1(&mtlsv1.MtlsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(mtlsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := mtlsService.GetAccessCertificate(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetAccessCertificateOptions model
				getAccessCertificateOptionsModel := new(mtlsv1.GetAccessCertificateOptions)
				getAccessCertificateOptionsModel.ZoneID = core.StringPtr("testString")
				getAccessCertificateOptionsModel.CertID = core.StringPtr("testString")
				getAccessCertificateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = mtlsService.GetAccessCertificate(getAccessCertificateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetAccessCertificate with error: Operation validation and request error`, func() {
				mtlsService, serviceErr := mtlsv1.NewMtlsV1(&mtlsv1.MtlsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(mtlsService).ToNot(BeNil())

				// Construct an instance of the GetAccessCertificateOptions model
				getAccessCertificateOptionsModel := new(mtlsv1.GetAccessCertificateOptions)
				getAccessCertificateOptionsModel.ZoneID = core.StringPtr("testString")
				getAccessCertificateOptionsModel.CertID = core.StringPtr("testString")
				getAccessCertificateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := mtlsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := mtlsService.GetAccessCertificate(getAccessCertificateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetAccessCertificateOptions model with no property values
				getAccessCertificateOptionsModelNew := new(mtlsv1.GetAccessCertificateOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = mtlsService.GetAccessCertificate(getAccessCertificateOptionsModelNew)
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
			It(`Invoke GetAccessCertificate successfully`, func() {
				mtlsService, serviceErr := mtlsv1.NewMtlsV1(&mtlsv1.MtlsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(mtlsService).ToNot(BeNil())

				// Construct an instance of the GetAccessCertificateOptions model
				getAccessCertificateOptionsModel := new(mtlsv1.GetAccessCertificateOptions)
				getAccessCertificateOptionsModel.ZoneID = core.StringPtr("testString")
				getAccessCertificateOptionsModel.CertID = core.StringPtr("testString")
				getAccessCertificateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := mtlsService.GetAccessCertificate(getAccessCertificateOptionsModel)
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
	Describe(`UpdateAccessCertificate(updateAccessCertificateOptions *UpdateAccessCertificateOptions) - Operation response error`, func() {
		crn := "testString"
		updateAccessCertificatePath := "/v1/testString/zones/testString/access/certificates/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateAccessCertificatePath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateAccessCertificate with error: Operation response processing error`, func() {
				mtlsService, serviceErr := mtlsv1.NewMtlsV1(&mtlsv1.MtlsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(mtlsService).ToNot(BeNil())

				// Construct an instance of the UpdateAccessCertificateOptions model
				updateAccessCertificateOptionsModel := new(mtlsv1.UpdateAccessCertificateOptions)
				updateAccessCertificateOptionsModel.ZoneID = core.StringPtr("testString")
				updateAccessCertificateOptionsModel.CertID = core.StringPtr("testString")
				updateAccessCertificateOptionsModel.Name = core.StringPtr("test-cert")
				updateAccessCertificateOptionsModel.AssociatedHostnames = []string{"test.example.com"}
				updateAccessCertificateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := mtlsService.UpdateAccessCertificate(updateAccessCertificateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				mtlsService.EnableRetries(0, 0)
				result, response, operationErr = mtlsService.UpdateAccessCertificate(updateAccessCertificateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateAccessCertificate(updateAccessCertificateOptions *UpdateAccessCertificateOptions)`, func() {
		crn := "testString"
		updateAccessCertificatePath := "/v1/testString/zones/testString/access/certificates/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateAccessCertificatePath))
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
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": "21a41336-9001-42c4-8440-c79e0cb86e1f", "name": "test-cert", "fingerprint": "MD5 Fingerprint=38:38:B4:FB:3C:33:CE:2C:8E:8E:D1:1B:94:70:C1:5F", "associated_hostnames": ["test.example.com"], "created_at": "2021-04-19T11:09:11Z", "updated_at": "2021-04-19T11:09:11Z", "expires_on": "2026-04-18T06:26:00Z"}}`)
				}))
			})
			It(`Invoke UpdateAccessCertificate successfully with retries`, func() {
				mtlsService, serviceErr := mtlsv1.NewMtlsV1(&mtlsv1.MtlsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(mtlsService).ToNot(BeNil())
				mtlsService.EnableRetries(0, 0)

				// Construct an instance of the UpdateAccessCertificateOptions model
				updateAccessCertificateOptionsModel := new(mtlsv1.UpdateAccessCertificateOptions)
				updateAccessCertificateOptionsModel.ZoneID = core.StringPtr("testString")
				updateAccessCertificateOptionsModel.CertID = core.StringPtr("testString")
				updateAccessCertificateOptionsModel.Name = core.StringPtr("test-cert")
				updateAccessCertificateOptionsModel.AssociatedHostnames = []string{"test.example.com"}
				updateAccessCertificateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := mtlsService.UpdateAccessCertificateWithContext(ctx, updateAccessCertificateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				mtlsService.DisableRetries()
				result, response, operationErr := mtlsService.UpdateAccessCertificate(updateAccessCertificateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = mtlsService.UpdateAccessCertificateWithContext(ctx, updateAccessCertificateOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(updateAccessCertificatePath))
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
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": "21a41336-9001-42c4-8440-c79e0cb86e1f", "name": "test-cert", "fingerprint": "MD5 Fingerprint=38:38:B4:FB:3C:33:CE:2C:8E:8E:D1:1B:94:70:C1:5F", "associated_hostnames": ["test.example.com"], "created_at": "2021-04-19T11:09:11Z", "updated_at": "2021-04-19T11:09:11Z", "expires_on": "2026-04-18T06:26:00Z"}}`)
				}))
			})
			It(`Invoke UpdateAccessCertificate successfully`, func() {
				mtlsService, serviceErr := mtlsv1.NewMtlsV1(&mtlsv1.MtlsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(mtlsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := mtlsService.UpdateAccessCertificate(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateAccessCertificateOptions model
				updateAccessCertificateOptionsModel := new(mtlsv1.UpdateAccessCertificateOptions)
				updateAccessCertificateOptionsModel.ZoneID = core.StringPtr("testString")
				updateAccessCertificateOptionsModel.CertID = core.StringPtr("testString")
				updateAccessCertificateOptionsModel.Name = core.StringPtr("test-cert")
				updateAccessCertificateOptionsModel.AssociatedHostnames = []string{"test.example.com"}
				updateAccessCertificateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = mtlsService.UpdateAccessCertificate(updateAccessCertificateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateAccessCertificate with error: Operation validation and request error`, func() {
				mtlsService, serviceErr := mtlsv1.NewMtlsV1(&mtlsv1.MtlsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(mtlsService).ToNot(BeNil())

				// Construct an instance of the UpdateAccessCertificateOptions model
				updateAccessCertificateOptionsModel := new(mtlsv1.UpdateAccessCertificateOptions)
				updateAccessCertificateOptionsModel.ZoneID = core.StringPtr("testString")
				updateAccessCertificateOptionsModel.CertID = core.StringPtr("testString")
				updateAccessCertificateOptionsModel.Name = core.StringPtr("test-cert")
				updateAccessCertificateOptionsModel.AssociatedHostnames = []string{"test.example.com"}
				updateAccessCertificateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := mtlsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := mtlsService.UpdateAccessCertificate(updateAccessCertificateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateAccessCertificateOptions model with no property values
				updateAccessCertificateOptionsModelNew := new(mtlsv1.UpdateAccessCertificateOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = mtlsService.UpdateAccessCertificate(updateAccessCertificateOptionsModelNew)
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
			It(`Invoke UpdateAccessCertificate successfully`, func() {
				mtlsService, serviceErr := mtlsv1.NewMtlsV1(&mtlsv1.MtlsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(mtlsService).ToNot(BeNil())

				// Construct an instance of the UpdateAccessCertificateOptions model
				updateAccessCertificateOptionsModel := new(mtlsv1.UpdateAccessCertificateOptions)
				updateAccessCertificateOptionsModel.ZoneID = core.StringPtr("testString")
				updateAccessCertificateOptionsModel.CertID = core.StringPtr("testString")
				updateAccessCertificateOptionsModel.Name = core.StringPtr("test-cert")
				updateAccessCertificateOptionsModel.AssociatedHostnames = []string{"test.example.com"}
				updateAccessCertificateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := mtlsService.UpdateAccessCertificate(updateAccessCertificateOptionsModel)
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
	Describe(`DeleteAccessCertificate(deleteAccessCertificateOptions *DeleteAccessCertificateOptions) - Operation response error`, func() {
		crn := "testString"
		deleteAccessCertificatePath := "/v1/testString/zones/testString/access/certificates/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteAccessCertificatePath))
					Expect(req.Method).To(Equal("DELETE"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke DeleteAccessCertificate with error: Operation response processing error`, func() {
				mtlsService, serviceErr := mtlsv1.NewMtlsV1(&mtlsv1.MtlsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(mtlsService).ToNot(BeNil())

				// Construct an instance of the DeleteAccessCertificateOptions model
				deleteAccessCertificateOptionsModel := new(mtlsv1.DeleteAccessCertificateOptions)
				deleteAccessCertificateOptionsModel.ZoneID = core.StringPtr("testString")
				deleteAccessCertificateOptionsModel.CertID = core.StringPtr("testString")
				deleteAccessCertificateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := mtlsService.DeleteAccessCertificate(deleteAccessCertificateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				mtlsService.EnableRetries(0, 0)
				result, response, operationErr = mtlsService.DeleteAccessCertificate(deleteAccessCertificateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteAccessCertificate(deleteAccessCertificateOptions *DeleteAccessCertificateOptions)`, func() {
		crn := "testString"
		deleteAccessCertificatePath := "/v1/testString/zones/testString/access/certificates/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteAccessCertificatePath))
					Expect(req.Method).To(Equal("DELETE"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": "21a41336-9001-42c4-8440-c79e0cb86e1f"}}`)
				}))
			})
			It(`Invoke DeleteAccessCertificate successfully with retries`, func() {
				mtlsService, serviceErr := mtlsv1.NewMtlsV1(&mtlsv1.MtlsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(mtlsService).ToNot(BeNil())
				mtlsService.EnableRetries(0, 0)

				// Construct an instance of the DeleteAccessCertificateOptions model
				deleteAccessCertificateOptionsModel := new(mtlsv1.DeleteAccessCertificateOptions)
				deleteAccessCertificateOptionsModel.ZoneID = core.StringPtr("testString")
				deleteAccessCertificateOptionsModel.CertID = core.StringPtr("testString")
				deleteAccessCertificateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := mtlsService.DeleteAccessCertificateWithContext(ctx, deleteAccessCertificateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				mtlsService.DisableRetries()
				result, response, operationErr := mtlsService.DeleteAccessCertificate(deleteAccessCertificateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = mtlsService.DeleteAccessCertificateWithContext(ctx, deleteAccessCertificateOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(deleteAccessCertificatePath))
					Expect(req.Method).To(Equal("DELETE"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": "21a41336-9001-42c4-8440-c79e0cb86e1f"}}`)
				}))
			})
			It(`Invoke DeleteAccessCertificate successfully`, func() {
				mtlsService, serviceErr := mtlsv1.NewMtlsV1(&mtlsv1.MtlsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(mtlsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := mtlsService.DeleteAccessCertificate(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DeleteAccessCertificateOptions model
				deleteAccessCertificateOptionsModel := new(mtlsv1.DeleteAccessCertificateOptions)
				deleteAccessCertificateOptionsModel.ZoneID = core.StringPtr("testString")
				deleteAccessCertificateOptionsModel.CertID = core.StringPtr("testString")
				deleteAccessCertificateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = mtlsService.DeleteAccessCertificate(deleteAccessCertificateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke DeleteAccessCertificate with error: Operation validation and request error`, func() {
				mtlsService, serviceErr := mtlsv1.NewMtlsV1(&mtlsv1.MtlsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(mtlsService).ToNot(BeNil())

				// Construct an instance of the DeleteAccessCertificateOptions model
				deleteAccessCertificateOptionsModel := new(mtlsv1.DeleteAccessCertificateOptions)
				deleteAccessCertificateOptionsModel.ZoneID = core.StringPtr("testString")
				deleteAccessCertificateOptionsModel.CertID = core.StringPtr("testString")
				deleteAccessCertificateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := mtlsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := mtlsService.DeleteAccessCertificate(deleteAccessCertificateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the DeleteAccessCertificateOptions model with no property values
				deleteAccessCertificateOptionsModelNew := new(mtlsv1.DeleteAccessCertificateOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = mtlsService.DeleteAccessCertificate(deleteAccessCertificateOptionsModelNew)
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
			It(`Invoke DeleteAccessCertificate successfully`, func() {
				mtlsService, serviceErr := mtlsv1.NewMtlsV1(&mtlsv1.MtlsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(mtlsService).ToNot(BeNil())

				// Construct an instance of the DeleteAccessCertificateOptions model
				deleteAccessCertificateOptionsModel := new(mtlsv1.DeleteAccessCertificateOptions)
				deleteAccessCertificateOptionsModel.ZoneID = core.StringPtr("testString")
				deleteAccessCertificateOptionsModel.CertID = core.StringPtr("testString")
				deleteAccessCertificateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := mtlsService.DeleteAccessCertificate(deleteAccessCertificateOptionsModel)
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
	Describe(`ListAccessApplications(listAccessApplicationsOptions *ListAccessApplicationsOptions) - Operation response error`, func() {
		crn := "testString"
		listAccessApplicationsPath := "/v1/testString/zones/testString/access/apps"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listAccessApplicationsPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListAccessApplications with error: Operation response processing error`, func() {
				mtlsService, serviceErr := mtlsv1.NewMtlsV1(&mtlsv1.MtlsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(mtlsService).ToNot(BeNil())

				// Construct an instance of the ListAccessApplicationsOptions model
				listAccessApplicationsOptionsModel := new(mtlsv1.ListAccessApplicationsOptions)
				listAccessApplicationsOptionsModel.ZoneID = core.StringPtr("testString")
				listAccessApplicationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := mtlsService.ListAccessApplications(listAccessApplicationsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				mtlsService.EnableRetries(0, 0)
				result, response, operationErr = mtlsService.ListAccessApplications(listAccessApplicationsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListAccessApplications(listAccessApplicationsOptions *ListAccessApplicationsOptions)`, func() {
		crn := "testString"
		listAccessApplicationsPath := "/v1/testString/zones/testString/access/apps"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listAccessApplicationsPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": [{"id": "de4526d6-d125-4f95-906f-1757510a9cd8", "name": "mtls-test-app", "domain": "test.example.com", "aud": "f8e1744453ea3679d919fdc6db58cff648f2b14b33a729f780fc02e75a42a008", "policies": [{"id": "acabcdb1-afb3-4f61-9dae-d1a353a93661", "name": "mtls-test-policy", "decision": "non_identity", "include": [{"certificate": {"anyKey": "anyValue"}}], "exclude": [{"certificate": {"anyKey": "anyValue"}}], "precedence": 1, "require": [{"certificate": {"anyKey": "anyValue"}}], "uid": "acabcdb1-afb3-4f61-9dae-d1a353a93661", "created_at": "2021-04-19T08:01:21Z", "updated_at": "2021-04-19T08:01:21Z"}], "allowed_idps": ["699d98642c564d2e855e9661899b7252"], "auto_redirect_to_identity": false, "session_duration": "24h", "type": "self_hosted", "uid": "de4526d6-d125-4f95-906f-1757510a9cd8", "created_at": "2021-04-19T07:59:49Z", "updated_at": "2021-04-19T07:59:49Z"}]}`)
				}))
			})
			It(`Invoke ListAccessApplications successfully with retries`, func() {
				mtlsService, serviceErr := mtlsv1.NewMtlsV1(&mtlsv1.MtlsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(mtlsService).ToNot(BeNil())
				mtlsService.EnableRetries(0, 0)

				// Construct an instance of the ListAccessApplicationsOptions model
				listAccessApplicationsOptionsModel := new(mtlsv1.ListAccessApplicationsOptions)
				listAccessApplicationsOptionsModel.ZoneID = core.StringPtr("testString")
				listAccessApplicationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := mtlsService.ListAccessApplicationsWithContext(ctx, listAccessApplicationsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				mtlsService.DisableRetries()
				result, response, operationErr := mtlsService.ListAccessApplications(listAccessApplicationsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = mtlsService.ListAccessApplicationsWithContext(ctx, listAccessApplicationsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listAccessApplicationsPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": [{"id": "de4526d6-d125-4f95-906f-1757510a9cd8", "name": "mtls-test-app", "domain": "test.example.com", "aud": "f8e1744453ea3679d919fdc6db58cff648f2b14b33a729f780fc02e75a42a008", "policies": [{"id": "acabcdb1-afb3-4f61-9dae-d1a353a93661", "name": "mtls-test-policy", "decision": "non_identity", "include": [{"certificate": {"anyKey": "anyValue"}}], "exclude": [{"certificate": {"anyKey": "anyValue"}}], "precedence": 1, "require": [{"certificate": {"anyKey": "anyValue"}}], "uid": "acabcdb1-afb3-4f61-9dae-d1a353a93661", "created_at": "2021-04-19T08:01:21Z", "updated_at": "2021-04-19T08:01:21Z"}], "allowed_idps": ["699d98642c564d2e855e9661899b7252"], "auto_redirect_to_identity": false, "session_duration": "24h", "type": "self_hosted", "uid": "de4526d6-d125-4f95-906f-1757510a9cd8", "created_at": "2021-04-19T07:59:49Z", "updated_at": "2021-04-19T07:59:49Z"}]}`)
				}))
			})
			It(`Invoke ListAccessApplications successfully`, func() {
				mtlsService, serviceErr := mtlsv1.NewMtlsV1(&mtlsv1.MtlsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(mtlsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := mtlsService.ListAccessApplications(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListAccessApplicationsOptions model
				listAccessApplicationsOptionsModel := new(mtlsv1.ListAccessApplicationsOptions)
				listAccessApplicationsOptionsModel.ZoneID = core.StringPtr("testString")
				listAccessApplicationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = mtlsService.ListAccessApplications(listAccessApplicationsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListAccessApplications with error: Operation validation and request error`, func() {
				mtlsService, serviceErr := mtlsv1.NewMtlsV1(&mtlsv1.MtlsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(mtlsService).ToNot(BeNil())

				// Construct an instance of the ListAccessApplicationsOptions model
				listAccessApplicationsOptionsModel := new(mtlsv1.ListAccessApplicationsOptions)
				listAccessApplicationsOptionsModel.ZoneID = core.StringPtr("testString")
				listAccessApplicationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := mtlsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := mtlsService.ListAccessApplications(listAccessApplicationsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListAccessApplicationsOptions model with no property values
				listAccessApplicationsOptionsModelNew := new(mtlsv1.ListAccessApplicationsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = mtlsService.ListAccessApplications(listAccessApplicationsOptionsModelNew)
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
			It(`Invoke ListAccessApplications successfully`, func() {
				mtlsService, serviceErr := mtlsv1.NewMtlsV1(&mtlsv1.MtlsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(mtlsService).ToNot(BeNil())

				// Construct an instance of the ListAccessApplicationsOptions model
				listAccessApplicationsOptionsModel := new(mtlsv1.ListAccessApplicationsOptions)
				listAccessApplicationsOptionsModel.ZoneID = core.StringPtr("testString")
				listAccessApplicationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := mtlsService.ListAccessApplications(listAccessApplicationsOptionsModel)
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
	Describe(`CreateAccessApplication(createAccessApplicationOptions *CreateAccessApplicationOptions) - Operation response error`, func() {
		crn := "testString"
		createAccessApplicationPath := "/v1/testString/zones/testString/access/apps"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createAccessApplicationPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateAccessApplication with error: Operation response processing error`, func() {
				mtlsService, serviceErr := mtlsv1.NewMtlsV1(&mtlsv1.MtlsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(mtlsService).ToNot(BeNil())

				// Construct an instance of the CreateAccessApplicationOptions model
				createAccessApplicationOptionsModel := new(mtlsv1.CreateAccessApplicationOptions)
				createAccessApplicationOptionsModel.ZoneID = core.StringPtr("testString")
				createAccessApplicationOptionsModel.Name = core.StringPtr("mtls-test-app")
				createAccessApplicationOptionsModel.Domain = core.StringPtr("test.example.com")
				createAccessApplicationOptionsModel.SessionDuration = core.StringPtr("24h")
				createAccessApplicationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := mtlsService.CreateAccessApplication(createAccessApplicationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				mtlsService.EnableRetries(0, 0)
				result, response, operationErr = mtlsService.CreateAccessApplication(createAccessApplicationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateAccessApplication(createAccessApplicationOptions *CreateAccessApplicationOptions)`, func() {
		crn := "testString"
		createAccessApplicationPath := "/v1/testString/zones/testString/access/apps"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createAccessApplicationPath))
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
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": "de4526d6-d125-4f95-906f-1757510a9cd8", "name": "mtls-test-app", "domain": "test.example.com", "aud": "f8e1744453ea3679d919fdc6db58cff648f2b14b33a729f780fc02e75a42a008", "policies": [{"anyKey": "anyValue"}], "allowed_idps": ["AllowedIdps"], "auto_redirect_to_identity": false, "session_duration": "24h", "type": "self_hosted", "uid": "de4526d6-d125-4f95-906f-1757510a9cd8", "created_at": "2021-04-19T07:59:49Z", "updated_at": "2021-04-19T07:59:49Z"}}`)
				}))
			})
			It(`Invoke CreateAccessApplication successfully with retries`, func() {
				mtlsService, serviceErr := mtlsv1.NewMtlsV1(&mtlsv1.MtlsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(mtlsService).ToNot(BeNil())
				mtlsService.EnableRetries(0, 0)

				// Construct an instance of the CreateAccessApplicationOptions model
				createAccessApplicationOptionsModel := new(mtlsv1.CreateAccessApplicationOptions)
				createAccessApplicationOptionsModel.ZoneID = core.StringPtr("testString")
				createAccessApplicationOptionsModel.Name = core.StringPtr("mtls-test-app")
				createAccessApplicationOptionsModel.Domain = core.StringPtr("test.example.com")
				createAccessApplicationOptionsModel.SessionDuration = core.StringPtr("24h")
				createAccessApplicationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := mtlsService.CreateAccessApplicationWithContext(ctx, createAccessApplicationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				mtlsService.DisableRetries()
				result, response, operationErr := mtlsService.CreateAccessApplication(createAccessApplicationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = mtlsService.CreateAccessApplicationWithContext(ctx, createAccessApplicationOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createAccessApplicationPath))
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
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": "de4526d6-d125-4f95-906f-1757510a9cd8", "name": "mtls-test-app", "domain": "test.example.com", "aud": "f8e1744453ea3679d919fdc6db58cff648f2b14b33a729f780fc02e75a42a008", "policies": [{"anyKey": "anyValue"}], "allowed_idps": ["AllowedIdps"], "auto_redirect_to_identity": false, "session_duration": "24h", "type": "self_hosted", "uid": "de4526d6-d125-4f95-906f-1757510a9cd8", "created_at": "2021-04-19T07:59:49Z", "updated_at": "2021-04-19T07:59:49Z"}}`)
				}))
			})
			It(`Invoke CreateAccessApplication successfully`, func() {
				mtlsService, serviceErr := mtlsv1.NewMtlsV1(&mtlsv1.MtlsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(mtlsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := mtlsService.CreateAccessApplication(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateAccessApplicationOptions model
				createAccessApplicationOptionsModel := new(mtlsv1.CreateAccessApplicationOptions)
				createAccessApplicationOptionsModel.ZoneID = core.StringPtr("testString")
				createAccessApplicationOptionsModel.Name = core.StringPtr("mtls-test-app")
				createAccessApplicationOptionsModel.Domain = core.StringPtr("test.example.com")
				createAccessApplicationOptionsModel.SessionDuration = core.StringPtr("24h")
				createAccessApplicationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = mtlsService.CreateAccessApplication(createAccessApplicationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateAccessApplication with error: Operation validation and request error`, func() {
				mtlsService, serviceErr := mtlsv1.NewMtlsV1(&mtlsv1.MtlsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(mtlsService).ToNot(BeNil())

				// Construct an instance of the CreateAccessApplicationOptions model
				createAccessApplicationOptionsModel := new(mtlsv1.CreateAccessApplicationOptions)
				createAccessApplicationOptionsModel.ZoneID = core.StringPtr("testString")
				createAccessApplicationOptionsModel.Name = core.StringPtr("mtls-test-app")
				createAccessApplicationOptionsModel.Domain = core.StringPtr("test.example.com")
				createAccessApplicationOptionsModel.SessionDuration = core.StringPtr("24h")
				createAccessApplicationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := mtlsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := mtlsService.CreateAccessApplication(createAccessApplicationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateAccessApplicationOptions model with no property values
				createAccessApplicationOptionsModelNew := new(mtlsv1.CreateAccessApplicationOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = mtlsService.CreateAccessApplication(createAccessApplicationOptionsModelNew)
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
			It(`Invoke CreateAccessApplication successfully`, func() {
				mtlsService, serviceErr := mtlsv1.NewMtlsV1(&mtlsv1.MtlsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(mtlsService).ToNot(BeNil())

				// Construct an instance of the CreateAccessApplicationOptions model
				createAccessApplicationOptionsModel := new(mtlsv1.CreateAccessApplicationOptions)
				createAccessApplicationOptionsModel.ZoneID = core.StringPtr("testString")
				createAccessApplicationOptionsModel.Name = core.StringPtr("mtls-test-app")
				createAccessApplicationOptionsModel.Domain = core.StringPtr("test.example.com")
				createAccessApplicationOptionsModel.SessionDuration = core.StringPtr("24h")
				createAccessApplicationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := mtlsService.CreateAccessApplication(createAccessApplicationOptionsModel)
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
	Describe(`GetAccessApplication(getAccessApplicationOptions *GetAccessApplicationOptions) - Operation response error`, func() {
		crn := "testString"
		getAccessApplicationPath := "/v1/testString/zones/testString/access/apps/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getAccessApplicationPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetAccessApplication with error: Operation response processing error`, func() {
				mtlsService, serviceErr := mtlsv1.NewMtlsV1(&mtlsv1.MtlsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(mtlsService).ToNot(BeNil())

				// Construct an instance of the GetAccessApplicationOptions model
				getAccessApplicationOptionsModel := new(mtlsv1.GetAccessApplicationOptions)
				getAccessApplicationOptionsModel.ZoneID = core.StringPtr("testString")
				getAccessApplicationOptionsModel.AppID = core.StringPtr("testString")
				getAccessApplicationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := mtlsService.GetAccessApplication(getAccessApplicationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				mtlsService.EnableRetries(0, 0)
				result, response, operationErr = mtlsService.GetAccessApplication(getAccessApplicationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetAccessApplication(getAccessApplicationOptions *GetAccessApplicationOptions)`, func() {
		crn := "testString"
		getAccessApplicationPath := "/v1/testString/zones/testString/access/apps/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getAccessApplicationPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": "de4526d6-d125-4f95-906f-1757510a9cd8", "name": "mtls-test-app", "domain": "test.example.com", "aud": "f8e1744453ea3679d919fdc6db58cff648f2b14b33a729f780fc02e75a42a008", "policies": [{"id": "acabcdb1-afb3-4f61-9dae-d1a353a93661", "name": "mtls-test-policy", "decision": "non_identity", "include": [{"certificate": {"anyKey": "anyValue"}}], "exclude": [{"certificate": {"anyKey": "anyValue"}}], "precedence": 1, "require": [{"certificate": {"anyKey": "anyValue"}}], "uid": "acabcdb1-afb3-4f61-9dae-d1a353a93661", "created_at": "2021-04-19T08:01:21Z", "updated_at": "2021-04-19T08:01:21Z"}], "allowed_idps": ["699d98642c564d2e855e9661899b7252"], "auto_redirect_to_identity": false, "session_duration": "24h", "type": "self_hosted", "uid": "de4526d6-d125-4f95-906f-1757510a9cd8", "created_at": "2021-04-19T07:59:49Z", "updated_at": "2021-04-19T07:59:49Z"}}`)
				}))
			})
			It(`Invoke GetAccessApplication successfully with retries`, func() {
				mtlsService, serviceErr := mtlsv1.NewMtlsV1(&mtlsv1.MtlsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(mtlsService).ToNot(BeNil())
				mtlsService.EnableRetries(0, 0)

				// Construct an instance of the GetAccessApplicationOptions model
				getAccessApplicationOptionsModel := new(mtlsv1.GetAccessApplicationOptions)
				getAccessApplicationOptionsModel.ZoneID = core.StringPtr("testString")
				getAccessApplicationOptionsModel.AppID = core.StringPtr("testString")
				getAccessApplicationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := mtlsService.GetAccessApplicationWithContext(ctx, getAccessApplicationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				mtlsService.DisableRetries()
				result, response, operationErr := mtlsService.GetAccessApplication(getAccessApplicationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = mtlsService.GetAccessApplicationWithContext(ctx, getAccessApplicationOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getAccessApplicationPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": "de4526d6-d125-4f95-906f-1757510a9cd8", "name": "mtls-test-app", "domain": "test.example.com", "aud": "f8e1744453ea3679d919fdc6db58cff648f2b14b33a729f780fc02e75a42a008", "policies": [{"id": "acabcdb1-afb3-4f61-9dae-d1a353a93661", "name": "mtls-test-policy", "decision": "non_identity", "include": [{"certificate": {"anyKey": "anyValue"}}], "exclude": [{"certificate": {"anyKey": "anyValue"}}], "precedence": 1, "require": [{"certificate": {"anyKey": "anyValue"}}], "uid": "acabcdb1-afb3-4f61-9dae-d1a353a93661", "created_at": "2021-04-19T08:01:21Z", "updated_at": "2021-04-19T08:01:21Z"}], "allowed_idps": ["699d98642c564d2e855e9661899b7252"], "auto_redirect_to_identity": false, "session_duration": "24h", "type": "self_hosted", "uid": "de4526d6-d125-4f95-906f-1757510a9cd8", "created_at": "2021-04-19T07:59:49Z", "updated_at": "2021-04-19T07:59:49Z"}}`)
				}))
			})
			It(`Invoke GetAccessApplication successfully`, func() {
				mtlsService, serviceErr := mtlsv1.NewMtlsV1(&mtlsv1.MtlsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(mtlsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := mtlsService.GetAccessApplication(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetAccessApplicationOptions model
				getAccessApplicationOptionsModel := new(mtlsv1.GetAccessApplicationOptions)
				getAccessApplicationOptionsModel.ZoneID = core.StringPtr("testString")
				getAccessApplicationOptionsModel.AppID = core.StringPtr("testString")
				getAccessApplicationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = mtlsService.GetAccessApplication(getAccessApplicationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetAccessApplication with error: Operation validation and request error`, func() {
				mtlsService, serviceErr := mtlsv1.NewMtlsV1(&mtlsv1.MtlsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(mtlsService).ToNot(BeNil())

				// Construct an instance of the GetAccessApplicationOptions model
				getAccessApplicationOptionsModel := new(mtlsv1.GetAccessApplicationOptions)
				getAccessApplicationOptionsModel.ZoneID = core.StringPtr("testString")
				getAccessApplicationOptionsModel.AppID = core.StringPtr("testString")
				getAccessApplicationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := mtlsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := mtlsService.GetAccessApplication(getAccessApplicationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetAccessApplicationOptions model with no property values
				getAccessApplicationOptionsModelNew := new(mtlsv1.GetAccessApplicationOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = mtlsService.GetAccessApplication(getAccessApplicationOptionsModelNew)
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
			It(`Invoke GetAccessApplication successfully`, func() {
				mtlsService, serviceErr := mtlsv1.NewMtlsV1(&mtlsv1.MtlsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(mtlsService).ToNot(BeNil())

				// Construct an instance of the GetAccessApplicationOptions model
				getAccessApplicationOptionsModel := new(mtlsv1.GetAccessApplicationOptions)
				getAccessApplicationOptionsModel.ZoneID = core.StringPtr("testString")
				getAccessApplicationOptionsModel.AppID = core.StringPtr("testString")
				getAccessApplicationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := mtlsService.GetAccessApplication(getAccessApplicationOptionsModel)
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
	Describe(`UpdateAccessApplication(updateAccessApplicationOptions *UpdateAccessApplicationOptions) - Operation response error`, func() {
		crn := "testString"
		updateAccessApplicationPath := "/v1/testString/zones/testString/access/apps/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateAccessApplicationPath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateAccessApplication with error: Operation response processing error`, func() {
				mtlsService, serviceErr := mtlsv1.NewMtlsV1(&mtlsv1.MtlsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(mtlsService).ToNot(BeNil())

				// Construct an instance of the UpdateAccessApplicationOptions model
				updateAccessApplicationOptionsModel := new(mtlsv1.UpdateAccessApplicationOptions)
				updateAccessApplicationOptionsModel.ZoneID = core.StringPtr("testString")
				updateAccessApplicationOptionsModel.AppID = core.StringPtr("testString")
				updateAccessApplicationOptionsModel.Name = core.StringPtr("mtls-test-app")
				updateAccessApplicationOptionsModel.Domain = core.StringPtr("test.example.com")
				updateAccessApplicationOptionsModel.SessionDuration = core.StringPtr("24h")
				updateAccessApplicationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := mtlsService.UpdateAccessApplication(updateAccessApplicationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				mtlsService.EnableRetries(0, 0)
				result, response, operationErr = mtlsService.UpdateAccessApplication(updateAccessApplicationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateAccessApplication(updateAccessApplicationOptions *UpdateAccessApplicationOptions)`, func() {
		crn := "testString"
		updateAccessApplicationPath := "/v1/testString/zones/testString/access/apps/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateAccessApplicationPath))
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
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": "de4526d6-d125-4f95-906f-1757510a9cd8", "name": "mtls-test-app", "domain": "test.example.com", "aud": "f8e1744453ea3679d919fdc6db58cff648f2b14b33a729f780fc02e75a42a008", "policies": [{"id": "acabcdb1-afb3-4f61-9dae-d1a353a93661", "name": "mtls-test-policy", "decision": "non_identity", "include": [{"certificate": {"anyKey": "anyValue"}}], "exclude": [{"certificate": {"anyKey": "anyValue"}}], "precedence": 1, "require": [{"certificate": {"anyKey": "anyValue"}}], "uid": "acabcdb1-afb3-4f61-9dae-d1a353a93661", "created_at": "2021-04-19T08:01:21Z", "updated_at": "2021-04-19T08:01:21Z"}], "allowed_idps": ["699d98642c564d2e855e9661899b7252"], "auto_redirect_to_identity": false, "session_duration": "24h", "type": "self_hosted", "uid": "de4526d6-d125-4f95-906f-1757510a9cd8", "created_at": "2021-04-19T07:59:49Z", "updated_at": "2021-04-19T07:59:49Z"}}`)
				}))
			})
			It(`Invoke UpdateAccessApplication successfully with retries`, func() {
				mtlsService, serviceErr := mtlsv1.NewMtlsV1(&mtlsv1.MtlsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(mtlsService).ToNot(BeNil())
				mtlsService.EnableRetries(0, 0)

				// Construct an instance of the UpdateAccessApplicationOptions model
				updateAccessApplicationOptionsModel := new(mtlsv1.UpdateAccessApplicationOptions)
				updateAccessApplicationOptionsModel.ZoneID = core.StringPtr("testString")
				updateAccessApplicationOptionsModel.AppID = core.StringPtr("testString")
				updateAccessApplicationOptionsModel.Name = core.StringPtr("mtls-test-app")
				updateAccessApplicationOptionsModel.Domain = core.StringPtr("test.example.com")
				updateAccessApplicationOptionsModel.SessionDuration = core.StringPtr("24h")
				updateAccessApplicationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := mtlsService.UpdateAccessApplicationWithContext(ctx, updateAccessApplicationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				mtlsService.DisableRetries()
				result, response, operationErr := mtlsService.UpdateAccessApplication(updateAccessApplicationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = mtlsService.UpdateAccessApplicationWithContext(ctx, updateAccessApplicationOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(updateAccessApplicationPath))
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
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": "de4526d6-d125-4f95-906f-1757510a9cd8", "name": "mtls-test-app", "domain": "test.example.com", "aud": "f8e1744453ea3679d919fdc6db58cff648f2b14b33a729f780fc02e75a42a008", "policies": [{"id": "acabcdb1-afb3-4f61-9dae-d1a353a93661", "name": "mtls-test-policy", "decision": "non_identity", "include": [{"certificate": {"anyKey": "anyValue"}}], "exclude": [{"certificate": {"anyKey": "anyValue"}}], "precedence": 1, "require": [{"certificate": {"anyKey": "anyValue"}}], "uid": "acabcdb1-afb3-4f61-9dae-d1a353a93661", "created_at": "2021-04-19T08:01:21Z", "updated_at": "2021-04-19T08:01:21Z"}], "allowed_idps": ["699d98642c564d2e855e9661899b7252"], "auto_redirect_to_identity": false, "session_duration": "24h", "type": "self_hosted", "uid": "de4526d6-d125-4f95-906f-1757510a9cd8", "created_at": "2021-04-19T07:59:49Z", "updated_at": "2021-04-19T07:59:49Z"}}`)
				}))
			})
			It(`Invoke UpdateAccessApplication successfully`, func() {
				mtlsService, serviceErr := mtlsv1.NewMtlsV1(&mtlsv1.MtlsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(mtlsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := mtlsService.UpdateAccessApplication(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateAccessApplicationOptions model
				updateAccessApplicationOptionsModel := new(mtlsv1.UpdateAccessApplicationOptions)
				updateAccessApplicationOptionsModel.ZoneID = core.StringPtr("testString")
				updateAccessApplicationOptionsModel.AppID = core.StringPtr("testString")
				updateAccessApplicationOptionsModel.Name = core.StringPtr("mtls-test-app")
				updateAccessApplicationOptionsModel.Domain = core.StringPtr("test.example.com")
				updateAccessApplicationOptionsModel.SessionDuration = core.StringPtr("24h")
				updateAccessApplicationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = mtlsService.UpdateAccessApplication(updateAccessApplicationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateAccessApplication with error: Operation validation and request error`, func() {
				mtlsService, serviceErr := mtlsv1.NewMtlsV1(&mtlsv1.MtlsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(mtlsService).ToNot(BeNil())

				// Construct an instance of the UpdateAccessApplicationOptions model
				updateAccessApplicationOptionsModel := new(mtlsv1.UpdateAccessApplicationOptions)
				updateAccessApplicationOptionsModel.ZoneID = core.StringPtr("testString")
				updateAccessApplicationOptionsModel.AppID = core.StringPtr("testString")
				updateAccessApplicationOptionsModel.Name = core.StringPtr("mtls-test-app")
				updateAccessApplicationOptionsModel.Domain = core.StringPtr("test.example.com")
				updateAccessApplicationOptionsModel.SessionDuration = core.StringPtr("24h")
				updateAccessApplicationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := mtlsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := mtlsService.UpdateAccessApplication(updateAccessApplicationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateAccessApplicationOptions model with no property values
				updateAccessApplicationOptionsModelNew := new(mtlsv1.UpdateAccessApplicationOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = mtlsService.UpdateAccessApplication(updateAccessApplicationOptionsModelNew)
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
			It(`Invoke UpdateAccessApplication successfully`, func() {
				mtlsService, serviceErr := mtlsv1.NewMtlsV1(&mtlsv1.MtlsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(mtlsService).ToNot(BeNil())

				// Construct an instance of the UpdateAccessApplicationOptions model
				updateAccessApplicationOptionsModel := new(mtlsv1.UpdateAccessApplicationOptions)
				updateAccessApplicationOptionsModel.ZoneID = core.StringPtr("testString")
				updateAccessApplicationOptionsModel.AppID = core.StringPtr("testString")
				updateAccessApplicationOptionsModel.Name = core.StringPtr("mtls-test-app")
				updateAccessApplicationOptionsModel.Domain = core.StringPtr("test.example.com")
				updateAccessApplicationOptionsModel.SessionDuration = core.StringPtr("24h")
				updateAccessApplicationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := mtlsService.UpdateAccessApplication(updateAccessApplicationOptionsModel)
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
	Describe(`DeleteAccessApplication(deleteAccessApplicationOptions *DeleteAccessApplicationOptions) - Operation response error`, func() {
		crn := "testString"
		deleteAccessApplicationPath := "/v1/testString/zones/testString/access/apps/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteAccessApplicationPath))
					Expect(req.Method).To(Equal("DELETE"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke DeleteAccessApplication with error: Operation response processing error`, func() {
				mtlsService, serviceErr := mtlsv1.NewMtlsV1(&mtlsv1.MtlsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(mtlsService).ToNot(BeNil())

				// Construct an instance of the DeleteAccessApplicationOptions model
				deleteAccessApplicationOptionsModel := new(mtlsv1.DeleteAccessApplicationOptions)
				deleteAccessApplicationOptionsModel.ZoneID = core.StringPtr("testString")
				deleteAccessApplicationOptionsModel.AppID = core.StringPtr("testString")
				deleteAccessApplicationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := mtlsService.DeleteAccessApplication(deleteAccessApplicationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				mtlsService.EnableRetries(0, 0)
				result, response, operationErr = mtlsService.DeleteAccessApplication(deleteAccessApplicationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteAccessApplication(deleteAccessApplicationOptions *DeleteAccessApplicationOptions)`, func() {
		crn := "testString"
		deleteAccessApplicationPath := "/v1/testString/zones/testString/access/apps/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteAccessApplicationPath))
					Expect(req.Method).To(Equal("DELETE"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": "de4526d6-d125-4f95-906f-1757510a9cd8"}}`)
				}))
			})
			It(`Invoke DeleteAccessApplication successfully with retries`, func() {
				mtlsService, serviceErr := mtlsv1.NewMtlsV1(&mtlsv1.MtlsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(mtlsService).ToNot(BeNil())
				mtlsService.EnableRetries(0, 0)

				// Construct an instance of the DeleteAccessApplicationOptions model
				deleteAccessApplicationOptionsModel := new(mtlsv1.DeleteAccessApplicationOptions)
				deleteAccessApplicationOptionsModel.ZoneID = core.StringPtr("testString")
				deleteAccessApplicationOptionsModel.AppID = core.StringPtr("testString")
				deleteAccessApplicationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := mtlsService.DeleteAccessApplicationWithContext(ctx, deleteAccessApplicationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				mtlsService.DisableRetries()
				result, response, operationErr := mtlsService.DeleteAccessApplication(deleteAccessApplicationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = mtlsService.DeleteAccessApplicationWithContext(ctx, deleteAccessApplicationOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(deleteAccessApplicationPath))
					Expect(req.Method).To(Equal("DELETE"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": "de4526d6-d125-4f95-906f-1757510a9cd8"}}`)
				}))
			})
			It(`Invoke DeleteAccessApplication successfully`, func() {
				mtlsService, serviceErr := mtlsv1.NewMtlsV1(&mtlsv1.MtlsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(mtlsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := mtlsService.DeleteAccessApplication(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DeleteAccessApplicationOptions model
				deleteAccessApplicationOptionsModel := new(mtlsv1.DeleteAccessApplicationOptions)
				deleteAccessApplicationOptionsModel.ZoneID = core.StringPtr("testString")
				deleteAccessApplicationOptionsModel.AppID = core.StringPtr("testString")
				deleteAccessApplicationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = mtlsService.DeleteAccessApplication(deleteAccessApplicationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke DeleteAccessApplication with error: Operation validation and request error`, func() {
				mtlsService, serviceErr := mtlsv1.NewMtlsV1(&mtlsv1.MtlsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(mtlsService).ToNot(BeNil())

				// Construct an instance of the DeleteAccessApplicationOptions model
				deleteAccessApplicationOptionsModel := new(mtlsv1.DeleteAccessApplicationOptions)
				deleteAccessApplicationOptionsModel.ZoneID = core.StringPtr("testString")
				deleteAccessApplicationOptionsModel.AppID = core.StringPtr("testString")
				deleteAccessApplicationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := mtlsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := mtlsService.DeleteAccessApplication(deleteAccessApplicationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the DeleteAccessApplicationOptions model with no property values
				deleteAccessApplicationOptionsModelNew := new(mtlsv1.DeleteAccessApplicationOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = mtlsService.DeleteAccessApplication(deleteAccessApplicationOptionsModelNew)
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
			It(`Invoke DeleteAccessApplication successfully`, func() {
				mtlsService, serviceErr := mtlsv1.NewMtlsV1(&mtlsv1.MtlsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(mtlsService).ToNot(BeNil())

				// Construct an instance of the DeleteAccessApplicationOptions model
				deleteAccessApplicationOptionsModel := new(mtlsv1.DeleteAccessApplicationOptions)
				deleteAccessApplicationOptionsModel.ZoneID = core.StringPtr("testString")
				deleteAccessApplicationOptionsModel.AppID = core.StringPtr("testString")
				deleteAccessApplicationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := mtlsService.DeleteAccessApplication(deleteAccessApplicationOptionsModel)
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
	Describe(`ListAccessPolicies(listAccessPoliciesOptions *ListAccessPoliciesOptions) - Operation response error`, func() {
		crn := "testString"
		listAccessPoliciesPath := "/v1/testString/zones/testString/access/apps/testString/policies"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listAccessPoliciesPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListAccessPolicies with error: Operation response processing error`, func() {
				mtlsService, serviceErr := mtlsv1.NewMtlsV1(&mtlsv1.MtlsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(mtlsService).ToNot(BeNil())

				// Construct an instance of the ListAccessPoliciesOptions model
				listAccessPoliciesOptionsModel := new(mtlsv1.ListAccessPoliciesOptions)
				listAccessPoliciesOptionsModel.ZoneID = core.StringPtr("testString")
				listAccessPoliciesOptionsModel.AppID = core.StringPtr("testString")
				listAccessPoliciesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := mtlsService.ListAccessPolicies(listAccessPoliciesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				mtlsService.EnableRetries(0, 0)
				result, response, operationErr = mtlsService.ListAccessPolicies(listAccessPoliciesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListAccessPolicies(listAccessPoliciesOptions *ListAccessPoliciesOptions)`, func() {
		crn := "testString"
		listAccessPoliciesPath := "/v1/testString/zones/testString/access/apps/testString/policies"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listAccessPoliciesPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": [{"id": "acabcdb1-afb3-4f61-9dae-d1a353a93661", "name": "mtls-test-policy", "decision": "non_identity", "include": [{"certificate": {"anyKey": "anyValue"}}], "exclude": [{"certificate": {"anyKey": "anyValue"}}], "precedence": 1, "require": [{"certificate": {"anyKey": "anyValue"}}], "uid": "acabcdb1-afb3-4f61-9dae-d1a353a93661", "created_at": "2021-04-19T08:01:21Z", "updated_at": "2021-04-19T08:01:21Z"}]}`)
				}))
			})
			It(`Invoke ListAccessPolicies successfully with retries`, func() {
				mtlsService, serviceErr := mtlsv1.NewMtlsV1(&mtlsv1.MtlsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(mtlsService).ToNot(BeNil())
				mtlsService.EnableRetries(0, 0)

				// Construct an instance of the ListAccessPoliciesOptions model
				listAccessPoliciesOptionsModel := new(mtlsv1.ListAccessPoliciesOptions)
				listAccessPoliciesOptionsModel.ZoneID = core.StringPtr("testString")
				listAccessPoliciesOptionsModel.AppID = core.StringPtr("testString")
				listAccessPoliciesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := mtlsService.ListAccessPoliciesWithContext(ctx, listAccessPoliciesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				mtlsService.DisableRetries()
				result, response, operationErr := mtlsService.ListAccessPolicies(listAccessPoliciesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = mtlsService.ListAccessPoliciesWithContext(ctx, listAccessPoliciesOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listAccessPoliciesPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": [{"id": "acabcdb1-afb3-4f61-9dae-d1a353a93661", "name": "mtls-test-policy", "decision": "non_identity", "include": [{"certificate": {"anyKey": "anyValue"}}], "exclude": [{"certificate": {"anyKey": "anyValue"}}], "precedence": 1, "require": [{"certificate": {"anyKey": "anyValue"}}], "uid": "acabcdb1-afb3-4f61-9dae-d1a353a93661", "created_at": "2021-04-19T08:01:21Z", "updated_at": "2021-04-19T08:01:21Z"}]}`)
				}))
			})
			It(`Invoke ListAccessPolicies successfully`, func() {
				mtlsService, serviceErr := mtlsv1.NewMtlsV1(&mtlsv1.MtlsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(mtlsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := mtlsService.ListAccessPolicies(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListAccessPoliciesOptions model
				listAccessPoliciesOptionsModel := new(mtlsv1.ListAccessPoliciesOptions)
				listAccessPoliciesOptionsModel.ZoneID = core.StringPtr("testString")
				listAccessPoliciesOptionsModel.AppID = core.StringPtr("testString")
				listAccessPoliciesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = mtlsService.ListAccessPolicies(listAccessPoliciesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListAccessPolicies with error: Operation validation and request error`, func() {
				mtlsService, serviceErr := mtlsv1.NewMtlsV1(&mtlsv1.MtlsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(mtlsService).ToNot(BeNil())

				// Construct an instance of the ListAccessPoliciesOptions model
				listAccessPoliciesOptionsModel := new(mtlsv1.ListAccessPoliciesOptions)
				listAccessPoliciesOptionsModel.ZoneID = core.StringPtr("testString")
				listAccessPoliciesOptionsModel.AppID = core.StringPtr("testString")
				listAccessPoliciesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := mtlsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := mtlsService.ListAccessPolicies(listAccessPoliciesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListAccessPoliciesOptions model with no property values
				listAccessPoliciesOptionsModelNew := new(mtlsv1.ListAccessPoliciesOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = mtlsService.ListAccessPolicies(listAccessPoliciesOptionsModelNew)
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
			It(`Invoke ListAccessPolicies successfully`, func() {
				mtlsService, serviceErr := mtlsv1.NewMtlsV1(&mtlsv1.MtlsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(mtlsService).ToNot(BeNil())

				// Construct an instance of the ListAccessPoliciesOptions model
				listAccessPoliciesOptionsModel := new(mtlsv1.ListAccessPoliciesOptions)
				listAccessPoliciesOptionsModel.ZoneID = core.StringPtr("testString")
				listAccessPoliciesOptionsModel.AppID = core.StringPtr("testString")
				listAccessPoliciesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := mtlsService.ListAccessPolicies(listAccessPoliciesOptionsModel)
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
	Describe(`CreateAccessPolicy(createAccessPolicyOptions *CreateAccessPolicyOptions) - Operation response error`, func() {
		crn := "testString"
		createAccessPolicyPath := "/v1/testString/zones/testString/access/apps/testString/policies"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createAccessPolicyPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateAccessPolicy with error: Operation response processing error`, func() {
				mtlsService, serviceErr := mtlsv1.NewMtlsV1(&mtlsv1.MtlsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(mtlsService).ToNot(BeNil())

				// Construct an instance of the PolicyRulePolicyCertRule model
				policyRuleModel := new(mtlsv1.PolicyRulePolicyCertRule)
				policyRuleModel.Certificate = map[string]interface{}{"anyKey": "anyValue"}

				// Construct an instance of the CreateAccessPolicyOptions model
				createAccessPolicyOptionsModel := new(mtlsv1.CreateAccessPolicyOptions)
				createAccessPolicyOptionsModel.ZoneID = core.StringPtr("testString")
				createAccessPolicyOptionsModel.AppID = core.StringPtr("testString")
				createAccessPolicyOptionsModel.Name = core.StringPtr("mtls-test-policy")
				createAccessPolicyOptionsModel.Decision = core.StringPtr("non_identity")
				createAccessPolicyOptionsModel.Include = []mtlsv1.PolicyRuleIntf{policyRuleModel}
				createAccessPolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := mtlsService.CreateAccessPolicy(createAccessPolicyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				mtlsService.EnableRetries(0, 0)
				result, response, operationErr = mtlsService.CreateAccessPolicy(createAccessPolicyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateAccessPolicy(createAccessPolicyOptions *CreateAccessPolicyOptions)`, func() {
		crn := "testString"
		createAccessPolicyPath := "/v1/testString/zones/testString/access/apps/testString/policies"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createAccessPolicyPath))
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
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": "acabcdb1-afb3-4f61-9dae-d1a353a93661", "name": "mtls-test-policy", "decision": "non_identity", "include": [{"certificate": {"anyKey": "anyValue"}}], "exclude": [{"certificate": {"anyKey": "anyValue"}}], "precedence": 1, "require": [{"certificate": {"anyKey": "anyValue"}}], "uid": "acabcdb1-afb3-4f61-9dae-d1a353a93661", "created_at": "2021-04-19T08:01:21Z", "updated_at": "2021-04-19T08:01:21Z"}}`)
				}))
			})
			It(`Invoke CreateAccessPolicy successfully with retries`, func() {
				mtlsService, serviceErr := mtlsv1.NewMtlsV1(&mtlsv1.MtlsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(mtlsService).ToNot(BeNil())
				mtlsService.EnableRetries(0, 0)

				// Construct an instance of the PolicyRulePolicyCertRule model
				policyRuleModel := new(mtlsv1.PolicyRulePolicyCertRule)
				policyRuleModel.Certificate = map[string]interface{}{"anyKey": "anyValue"}

				// Construct an instance of the CreateAccessPolicyOptions model
				createAccessPolicyOptionsModel := new(mtlsv1.CreateAccessPolicyOptions)
				createAccessPolicyOptionsModel.ZoneID = core.StringPtr("testString")
				createAccessPolicyOptionsModel.AppID = core.StringPtr("testString")
				createAccessPolicyOptionsModel.Name = core.StringPtr("mtls-test-policy")
				createAccessPolicyOptionsModel.Decision = core.StringPtr("non_identity")
				createAccessPolicyOptionsModel.Include = []mtlsv1.PolicyRuleIntf{policyRuleModel}
				createAccessPolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := mtlsService.CreateAccessPolicyWithContext(ctx, createAccessPolicyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				mtlsService.DisableRetries()
				result, response, operationErr := mtlsService.CreateAccessPolicy(createAccessPolicyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = mtlsService.CreateAccessPolicyWithContext(ctx, createAccessPolicyOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createAccessPolicyPath))
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
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": "acabcdb1-afb3-4f61-9dae-d1a353a93661", "name": "mtls-test-policy", "decision": "non_identity", "include": [{"certificate": {"anyKey": "anyValue"}}], "exclude": [{"certificate": {"anyKey": "anyValue"}}], "precedence": 1, "require": [{"certificate": {"anyKey": "anyValue"}}], "uid": "acabcdb1-afb3-4f61-9dae-d1a353a93661", "created_at": "2021-04-19T08:01:21Z", "updated_at": "2021-04-19T08:01:21Z"}}`)
				}))
			})
			It(`Invoke CreateAccessPolicy successfully`, func() {
				mtlsService, serviceErr := mtlsv1.NewMtlsV1(&mtlsv1.MtlsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(mtlsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := mtlsService.CreateAccessPolicy(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the PolicyRulePolicyCertRule model
				policyRuleModel := new(mtlsv1.PolicyRulePolicyCertRule)
				policyRuleModel.Certificate = map[string]interface{}{"anyKey": "anyValue"}

				// Construct an instance of the CreateAccessPolicyOptions model
				createAccessPolicyOptionsModel := new(mtlsv1.CreateAccessPolicyOptions)
				createAccessPolicyOptionsModel.ZoneID = core.StringPtr("testString")
				createAccessPolicyOptionsModel.AppID = core.StringPtr("testString")
				createAccessPolicyOptionsModel.Name = core.StringPtr("mtls-test-policy")
				createAccessPolicyOptionsModel.Decision = core.StringPtr("non_identity")
				createAccessPolicyOptionsModel.Include = []mtlsv1.PolicyRuleIntf{policyRuleModel}
				createAccessPolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = mtlsService.CreateAccessPolicy(createAccessPolicyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateAccessPolicy with error: Operation validation and request error`, func() {
				mtlsService, serviceErr := mtlsv1.NewMtlsV1(&mtlsv1.MtlsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(mtlsService).ToNot(BeNil())

				// Construct an instance of the PolicyRulePolicyCertRule model
				policyRuleModel := new(mtlsv1.PolicyRulePolicyCertRule)
				policyRuleModel.Certificate = map[string]interface{}{"anyKey": "anyValue"}

				// Construct an instance of the CreateAccessPolicyOptions model
				createAccessPolicyOptionsModel := new(mtlsv1.CreateAccessPolicyOptions)
				createAccessPolicyOptionsModel.ZoneID = core.StringPtr("testString")
				createAccessPolicyOptionsModel.AppID = core.StringPtr("testString")
				createAccessPolicyOptionsModel.Name = core.StringPtr("mtls-test-policy")
				createAccessPolicyOptionsModel.Decision = core.StringPtr("non_identity")
				createAccessPolicyOptionsModel.Include = []mtlsv1.PolicyRuleIntf{policyRuleModel}
				createAccessPolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := mtlsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := mtlsService.CreateAccessPolicy(createAccessPolicyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateAccessPolicyOptions model with no property values
				createAccessPolicyOptionsModelNew := new(mtlsv1.CreateAccessPolicyOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = mtlsService.CreateAccessPolicy(createAccessPolicyOptionsModelNew)
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
			It(`Invoke CreateAccessPolicy successfully`, func() {
				mtlsService, serviceErr := mtlsv1.NewMtlsV1(&mtlsv1.MtlsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(mtlsService).ToNot(BeNil())

				// Construct an instance of the PolicyRulePolicyCertRule model
				policyRuleModel := new(mtlsv1.PolicyRulePolicyCertRule)
				policyRuleModel.Certificate = map[string]interface{}{"anyKey": "anyValue"}

				// Construct an instance of the CreateAccessPolicyOptions model
				createAccessPolicyOptionsModel := new(mtlsv1.CreateAccessPolicyOptions)
				createAccessPolicyOptionsModel.ZoneID = core.StringPtr("testString")
				createAccessPolicyOptionsModel.AppID = core.StringPtr("testString")
				createAccessPolicyOptionsModel.Name = core.StringPtr("mtls-test-policy")
				createAccessPolicyOptionsModel.Decision = core.StringPtr("non_identity")
				createAccessPolicyOptionsModel.Include = []mtlsv1.PolicyRuleIntf{policyRuleModel}
				createAccessPolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := mtlsService.CreateAccessPolicy(createAccessPolicyOptionsModel)
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
	Describe(`GetAccessPolicy(getAccessPolicyOptions *GetAccessPolicyOptions) - Operation response error`, func() {
		crn := "testString"
		getAccessPolicyPath := "/v1/testString/zones/testString/access/apps/testString/policies/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getAccessPolicyPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetAccessPolicy with error: Operation response processing error`, func() {
				mtlsService, serviceErr := mtlsv1.NewMtlsV1(&mtlsv1.MtlsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(mtlsService).ToNot(BeNil())

				// Construct an instance of the GetAccessPolicyOptions model
				getAccessPolicyOptionsModel := new(mtlsv1.GetAccessPolicyOptions)
				getAccessPolicyOptionsModel.ZoneID = core.StringPtr("testString")
				getAccessPolicyOptionsModel.AppID = core.StringPtr("testString")
				getAccessPolicyOptionsModel.PolicyID = core.StringPtr("testString")
				getAccessPolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := mtlsService.GetAccessPolicy(getAccessPolicyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				mtlsService.EnableRetries(0, 0)
				result, response, operationErr = mtlsService.GetAccessPolicy(getAccessPolicyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetAccessPolicy(getAccessPolicyOptions *GetAccessPolicyOptions)`, func() {
		crn := "testString"
		getAccessPolicyPath := "/v1/testString/zones/testString/access/apps/testString/policies/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getAccessPolicyPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": "acabcdb1-afb3-4f61-9dae-d1a353a93661", "name": "mtls-test-policy", "decision": "non_identity", "include": [{"certificate": {"anyKey": "anyValue"}}], "exclude": [{"certificate": {"anyKey": "anyValue"}}], "precedence": 1, "require": [{"certificate": {"anyKey": "anyValue"}}], "uid": "acabcdb1-afb3-4f61-9dae-d1a353a93661", "created_at": "2021-04-19T08:01:21Z", "updated_at": "2021-04-19T08:01:21Z"}}`)
				}))
			})
			It(`Invoke GetAccessPolicy successfully with retries`, func() {
				mtlsService, serviceErr := mtlsv1.NewMtlsV1(&mtlsv1.MtlsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(mtlsService).ToNot(BeNil())
				mtlsService.EnableRetries(0, 0)

				// Construct an instance of the GetAccessPolicyOptions model
				getAccessPolicyOptionsModel := new(mtlsv1.GetAccessPolicyOptions)
				getAccessPolicyOptionsModel.ZoneID = core.StringPtr("testString")
				getAccessPolicyOptionsModel.AppID = core.StringPtr("testString")
				getAccessPolicyOptionsModel.PolicyID = core.StringPtr("testString")
				getAccessPolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := mtlsService.GetAccessPolicyWithContext(ctx, getAccessPolicyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				mtlsService.DisableRetries()
				result, response, operationErr := mtlsService.GetAccessPolicy(getAccessPolicyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = mtlsService.GetAccessPolicyWithContext(ctx, getAccessPolicyOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getAccessPolicyPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": "acabcdb1-afb3-4f61-9dae-d1a353a93661", "name": "mtls-test-policy", "decision": "non_identity", "include": [{"certificate": {"anyKey": "anyValue"}}], "exclude": [{"certificate": {"anyKey": "anyValue"}}], "precedence": 1, "require": [{"certificate": {"anyKey": "anyValue"}}], "uid": "acabcdb1-afb3-4f61-9dae-d1a353a93661", "created_at": "2021-04-19T08:01:21Z", "updated_at": "2021-04-19T08:01:21Z"}}`)
				}))
			})
			It(`Invoke GetAccessPolicy successfully`, func() {
				mtlsService, serviceErr := mtlsv1.NewMtlsV1(&mtlsv1.MtlsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(mtlsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := mtlsService.GetAccessPolicy(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetAccessPolicyOptions model
				getAccessPolicyOptionsModel := new(mtlsv1.GetAccessPolicyOptions)
				getAccessPolicyOptionsModel.ZoneID = core.StringPtr("testString")
				getAccessPolicyOptionsModel.AppID = core.StringPtr("testString")
				getAccessPolicyOptionsModel.PolicyID = core.StringPtr("testString")
				getAccessPolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = mtlsService.GetAccessPolicy(getAccessPolicyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetAccessPolicy with error: Operation validation and request error`, func() {
				mtlsService, serviceErr := mtlsv1.NewMtlsV1(&mtlsv1.MtlsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(mtlsService).ToNot(BeNil())

				// Construct an instance of the GetAccessPolicyOptions model
				getAccessPolicyOptionsModel := new(mtlsv1.GetAccessPolicyOptions)
				getAccessPolicyOptionsModel.ZoneID = core.StringPtr("testString")
				getAccessPolicyOptionsModel.AppID = core.StringPtr("testString")
				getAccessPolicyOptionsModel.PolicyID = core.StringPtr("testString")
				getAccessPolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := mtlsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := mtlsService.GetAccessPolicy(getAccessPolicyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetAccessPolicyOptions model with no property values
				getAccessPolicyOptionsModelNew := new(mtlsv1.GetAccessPolicyOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = mtlsService.GetAccessPolicy(getAccessPolicyOptionsModelNew)
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
			It(`Invoke GetAccessPolicy successfully`, func() {
				mtlsService, serviceErr := mtlsv1.NewMtlsV1(&mtlsv1.MtlsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(mtlsService).ToNot(BeNil())

				// Construct an instance of the GetAccessPolicyOptions model
				getAccessPolicyOptionsModel := new(mtlsv1.GetAccessPolicyOptions)
				getAccessPolicyOptionsModel.ZoneID = core.StringPtr("testString")
				getAccessPolicyOptionsModel.AppID = core.StringPtr("testString")
				getAccessPolicyOptionsModel.PolicyID = core.StringPtr("testString")
				getAccessPolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := mtlsService.GetAccessPolicy(getAccessPolicyOptionsModel)
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
	Describe(`UpdateAccessPolicy(updateAccessPolicyOptions *UpdateAccessPolicyOptions) - Operation response error`, func() {
		crn := "testString"
		updateAccessPolicyPath := "/v1/testString/zones/testString/access/apps/testString/policies/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateAccessPolicyPath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateAccessPolicy with error: Operation response processing error`, func() {
				mtlsService, serviceErr := mtlsv1.NewMtlsV1(&mtlsv1.MtlsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(mtlsService).ToNot(BeNil())

				// Construct an instance of the PolicyRulePolicyCertRule model
				policyRuleModel := new(mtlsv1.PolicyRulePolicyCertRule)
				policyRuleModel.Certificate = map[string]interface{}{"anyKey": "anyValue"}

				// Construct an instance of the UpdateAccessPolicyOptions model
				updateAccessPolicyOptionsModel := new(mtlsv1.UpdateAccessPolicyOptions)
				updateAccessPolicyOptionsModel.ZoneID = core.StringPtr("testString")
				updateAccessPolicyOptionsModel.AppID = core.StringPtr("testString")
				updateAccessPolicyOptionsModel.PolicyID = core.StringPtr("testString")
				updateAccessPolicyOptionsModel.Name = core.StringPtr("mtls-test-policy")
				updateAccessPolicyOptionsModel.Decision = core.StringPtr("non_identity")
				updateAccessPolicyOptionsModel.Include = []mtlsv1.PolicyRuleIntf{policyRuleModel}
				updateAccessPolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := mtlsService.UpdateAccessPolicy(updateAccessPolicyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				mtlsService.EnableRetries(0, 0)
				result, response, operationErr = mtlsService.UpdateAccessPolicy(updateAccessPolicyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateAccessPolicy(updateAccessPolicyOptions *UpdateAccessPolicyOptions)`, func() {
		crn := "testString"
		updateAccessPolicyPath := "/v1/testString/zones/testString/access/apps/testString/policies/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateAccessPolicyPath))
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
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": "acabcdb1-afb3-4f61-9dae-d1a353a93661", "name": "mtls-test-policy", "decision": "non_identity", "include": [{"certificate": {"anyKey": "anyValue"}}], "exclude": [{"certificate": {"anyKey": "anyValue"}}], "precedence": 1, "require": [{"certificate": {"anyKey": "anyValue"}}], "uid": "acabcdb1-afb3-4f61-9dae-d1a353a93661", "created_at": "2021-04-19T08:01:21Z", "updated_at": "2021-04-19T08:01:21Z"}}`)
				}))
			})
			It(`Invoke UpdateAccessPolicy successfully with retries`, func() {
				mtlsService, serviceErr := mtlsv1.NewMtlsV1(&mtlsv1.MtlsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(mtlsService).ToNot(BeNil())
				mtlsService.EnableRetries(0, 0)

				// Construct an instance of the PolicyRulePolicyCertRule model
				policyRuleModel := new(mtlsv1.PolicyRulePolicyCertRule)
				policyRuleModel.Certificate = map[string]interface{}{"anyKey": "anyValue"}

				// Construct an instance of the UpdateAccessPolicyOptions model
				updateAccessPolicyOptionsModel := new(mtlsv1.UpdateAccessPolicyOptions)
				updateAccessPolicyOptionsModel.ZoneID = core.StringPtr("testString")
				updateAccessPolicyOptionsModel.AppID = core.StringPtr("testString")
				updateAccessPolicyOptionsModel.PolicyID = core.StringPtr("testString")
				updateAccessPolicyOptionsModel.Name = core.StringPtr("mtls-test-policy")
				updateAccessPolicyOptionsModel.Decision = core.StringPtr("non_identity")
				updateAccessPolicyOptionsModel.Include = []mtlsv1.PolicyRuleIntf{policyRuleModel}
				updateAccessPolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := mtlsService.UpdateAccessPolicyWithContext(ctx, updateAccessPolicyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				mtlsService.DisableRetries()
				result, response, operationErr := mtlsService.UpdateAccessPolicy(updateAccessPolicyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = mtlsService.UpdateAccessPolicyWithContext(ctx, updateAccessPolicyOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(updateAccessPolicyPath))
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
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": "acabcdb1-afb3-4f61-9dae-d1a353a93661", "name": "mtls-test-policy", "decision": "non_identity", "include": [{"certificate": {"anyKey": "anyValue"}}], "exclude": [{"certificate": {"anyKey": "anyValue"}}], "precedence": 1, "require": [{"certificate": {"anyKey": "anyValue"}}], "uid": "acabcdb1-afb3-4f61-9dae-d1a353a93661", "created_at": "2021-04-19T08:01:21Z", "updated_at": "2021-04-19T08:01:21Z"}}`)
				}))
			})
			It(`Invoke UpdateAccessPolicy successfully`, func() {
				mtlsService, serviceErr := mtlsv1.NewMtlsV1(&mtlsv1.MtlsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(mtlsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := mtlsService.UpdateAccessPolicy(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the PolicyRulePolicyCertRule model
				policyRuleModel := new(mtlsv1.PolicyRulePolicyCertRule)
				policyRuleModel.Certificate = map[string]interface{}{"anyKey": "anyValue"}

				// Construct an instance of the UpdateAccessPolicyOptions model
				updateAccessPolicyOptionsModel := new(mtlsv1.UpdateAccessPolicyOptions)
				updateAccessPolicyOptionsModel.ZoneID = core.StringPtr("testString")
				updateAccessPolicyOptionsModel.AppID = core.StringPtr("testString")
				updateAccessPolicyOptionsModel.PolicyID = core.StringPtr("testString")
				updateAccessPolicyOptionsModel.Name = core.StringPtr("mtls-test-policy")
				updateAccessPolicyOptionsModel.Decision = core.StringPtr("non_identity")
				updateAccessPolicyOptionsModel.Include = []mtlsv1.PolicyRuleIntf{policyRuleModel}
				updateAccessPolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = mtlsService.UpdateAccessPolicy(updateAccessPolicyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateAccessPolicy with error: Operation validation and request error`, func() {
				mtlsService, serviceErr := mtlsv1.NewMtlsV1(&mtlsv1.MtlsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(mtlsService).ToNot(BeNil())

				// Construct an instance of the PolicyRulePolicyCertRule model
				policyRuleModel := new(mtlsv1.PolicyRulePolicyCertRule)
				policyRuleModel.Certificate = map[string]interface{}{"anyKey": "anyValue"}

				// Construct an instance of the UpdateAccessPolicyOptions model
				updateAccessPolicyOptionsModel := new(mtlsv1.UpdateAccessPolicyOptions)
				updateAccessPolicyOptionsModel.ZoneID = core.StringPtr("testString")
				updateAccessPolicyOptionsModel.AppID = core.StringPtr("testString")
				updateAccessPolicyOptionsModel.PolicyID = core.StringPtr("testString")
				updateAccessPolicyOptionsModel.Name = core.StringPtr("mtls-test-policy")
				updateAccessPolicyOptionsModel.Decision = core.StringPtr("non_identity")
				updateAccessPolicyOptionsModel.Include = []mtlsv1.PolicyRuleIntf{policyRuleModel}
				updateAccessPolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := mtlsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := mtlsService.UpdateAccessPolicy(updateAccessPolicyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateAccessPolicyOptions model with no property values
				updateAccessPolicyOptionsModelNew := new(mtlsv1.UpdateAccessPolicyOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = mtlsService.UpdateAccessPolicy(updateAccessPolicyOptionsModelNew)
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
			It(`Invoke UpdateAccessPolicy successfully`, func() {
				mtlsService, serviceErr := mtlsv1.NewMtlsV1(&mtlsv1.MtlsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(mtlsService).ToNot(BeNil())

				// Construct an instance of the PolicyRulePolicyCertRule model
				policyRuleModel := new(mtlsv1.PolicyRulePolicyCertRule)
				policyRuleModel.Certificate = map[string]interface{}{"anyKey": "anyValue"}

				// Construct an instance of the UpdateAccessPolicyOptions model
				updateAccessPolicyOptionsModel := new(mtlsv1.UpdateAccessPolicyOptions)
				updateAccessPolicyOptionsModel.ZoneID = core.StringPtr("testString")
				updateAccessPolicyOptionsModel.AppID = core.StringPtr("testString")
				updateAccessPolicyOptionsModel.PolicyID = core.StringPtr("testString")
				updateAccessPolicyOptionsModel.Name = core.StringPtr("mtls-test-policy")
				updateAccessPolicyOptionsModel.Decision = core.StringPtr("non_identity")
				updateAccessPolicyOptionsModel.Include = []mtlsv1.PolicyRuleIntf{policyRuleModel}
				updateAccessPolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := mtlsService.UpdateAccessPolicy(updateAccessPolicyOptionsModel)
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
	Describe(`DeleteAccessPolicy(deleteAccessPolicyOptions *DeleteAccessPolicyOptions) - Operation response error`, func() {
		crn := "testString"
		deleteAccessPolicyPath := "/v1/testString/zones/testString/access/apps/testString/policies/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteAccessPolicyPath))
					Expect(req.Method).To(Equal("DELETE"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke DeleteAccessPolicy with error: Operation response processing error`, func() {
				mtlsService, serviceErr := mtlsv1.NewMtlsV1(&mtlsv1.MtlsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(mtlsService).ToNot(BeNil())

				// Construct an instance of the DeleteAccessPolicyOptions model
				deleteAccessPolicyOptionsModel := new(mtlsv1.DeleteAccessPolicyOptions)
				deleteAccessPolicyOptionsModel.ZoneID = core.StringPtr("testString")
				deleteAccessPolicyOptionsModel.AppID = core.StringPtr("testString")
				deleteAccessPolicyOptionsModel.PolicyID = core.StringPtr("testString")
				deleteAccessPolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := mtlsService.DeleteAccessPolicy(deleteAccessPolicyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				mtlsService.EnableRetries(0, 0)
				result, response, operationErr = mtlsService.DeleteAccessPolicy(deleteAccessPolicyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteAccessPolicy(deleteAccessPolicyOptions *DeleteAccessPolicyOptions)`, func() {
		crn := "testString"
		deleteAccessPolicyPath := "/v1/testString/zones/testString/access/apps/testString/policies/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteAccessPolicyPath))
					Expect(req.Method).To(Equal("DELETE"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": "acabcdb1-afb3-4f61-9dae-d1a353a93661"}}`)
				}))
			})
			It(`Invoke DeleteAccessPolicy successfully with retries`, func() {
				mtlsService, serviceErr := mtlsv1.NewMtlsV1(&mtlsv1.MtlsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(mtlsService).ToNot(BeNil())
				mtlsService.EnableRetries(0, 0)

				// Construct an instance of the DeleteAccessPolicyOptions model
				deleteAccessPolicyOptionsModel := new(mtlsv1.DeleteAccessPolicyOptions)
				deleteAccessPolicyOptionsModel.ZoneID = core.StringPtr("testString")
				deleteAccessPolicyOptionsModel.AppID = core.StringPtr("testString")
				deleteAccessPolicyOptionsModel.PolicyID = core.StringPtr("testString")
				deleteAccessPolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := mtlsService.DeleteAccessPolicyWithContext(ctx, deleteAccessPolicyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				mtlsService.DisableRetries()
				result, response, operationErr := mtlsService.DeleteAccessPolicy(deleteAccessPolicyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = mtlsService.DeleteAccessPolicyWithContext(ctx, deleteAccessPolicyOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(deleteAccessPolicyPath))
					Expect(req.Method).To(Equal("DELETE"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": "acabcdb1-afb3-4f61-9dae-d1a353a93661"}}`)
				}))
			})
			It(`Invoke DeleteAccessPolicy successfully`, func() {
				mtlsService, serviceErr := mtlsv1.NewMtlsV1(&mtlsv1.MtlsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(mtlsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := mtlsService.DeleteAccessPolicy(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DeleteAccessPolicyOptions model
				deleteAccessPolicyOptionsModel := new(mtlsv1.DeleteAccessPolicyOptions)
				deleteAccessPolicyOptionsModel.ZoneID = core.StringPtr("testString")
				deleteAccessPolicyOptionsModel.AppID = core.StringPtr("testString")
				deleteAccessPolicyOptionsModel.PolicyID = core.StringPtr("testString")
				deleteAccessPolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = mtlsService.DeleteAccessPolicy(deleteAccessPolicyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke DeleteAccessPolicy with error: Operation validation and request error`, func() {
				mtlsService, serviceErr := mtlsv1.NewMtlsV1(&mtlsv1.MtlsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(mtlsService).ToNot(BeNil())

				// Construct an instance of the DeleteAccessPolicyOptions model
				deleteAccessPolicyOptionsModel := new(mtlsv1.DeleteAccessPolicyOptions)
				deleteAccessPolicyOptionsModel.ZoneID = core.StringPtr("testString")
				deleteAccessPolicyOptionsModel.AppID = core.StringPtr("testString")
				deleteAccessPolicyOptionsModel.PolicyID = core.StringPtr("testString")
				deleteAccessPolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := mtlsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := mtlsService.DeleteAccessPolicy(deleteAccessPolicyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the DeleteAccessPolicyOptions model with no property values
				deleteAccessPolicyOptionsModelNew := new(mtlsv1.DeleteAccessPolicyOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = mtlsService.DeleteAccessPolicy(deleteAccessPolicyOptionsModelNew)
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
			It(`Invoke DeleteAccessPolicy successfully`, func() {
				mtlsService, serviceErr := mtlsv1.NewMtlsV1(&mtlsv1.MtlsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(mtlsService).ToNot(BeNil())

				// Construct an instance of the DeleteAccessPolicyOptions model
				deleteAccessPolicyOptionsModel := new(mtlsv1.DeleteAccessPolicyOptions)
				deleteAccessPolicyOptionsModel.ZoneID = core.StringPtr("testString")
				deleteAccessPolicyOptionsModel.AppID = core.StringPtr("testString")
				deleteAccessPolicyOptionsModel.PolicyID = core.StringPtr("testString")
				deleteAccessPolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := mtlsService.DeleteAccessPolicy(deleteAccessPolicyOptionsModel)
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
	Describe(`GetAccessCertSettings(getAccessCertSettingsOptions *GetAccessCertSettingsOptions) - Operation response error`, func() {
		crn := "testString"
		getAccessCertSettingsPath := "/v1/testString/zones/testString/access/certificates/settings"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getAccessCertSettingsPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetAccessCertSettings with error: Operation response processing error`, func() {
				mtlsService, serviceErr := mtlsv1.NewMtlsV1(&mtlsv1.MtlsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(mtlsService).ToNot(BeNil())

				// Construct an instance of the GetAccessCertSettingsOptions model
				getAccessCertSettingsOptionsModel := new(mtlsv1.GetAccessCertSettingsOptions)
				getAccessCertSettingsOptionsModel.ZoneID = core.StringPtr("testString")
				getAccessCertSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := mtlsService.GetAccessCertSettings(getAccessCertSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				mtlsService.EnableRetries(0, 0)
				result, response, operationErr = mtlsService.GetAccessCertSettings(getAccessCertSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetAccessCertSettings(getAccessCertSettingsOptions *GetAccessCertSettingsOptions)`, func() {
		crn := "testString"
		getAccessCertSettingsPath := "/v1/testString/zones/testString/access/certificates/settings"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getAccessCertSettingsPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": [{"hostname": "test.example.com", "china_network": false, "client_certificate_forwarding": true}]}`)
				}))
			})
			It(`Invoke GetAccessCertSettings successfully with retries`, func() {
				mtlsService, serviceErr := mtlsv1.NewMtlsV1(&mtlsv1.MtlsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(mtlsService).ToNot(BeNil())
				mtlsService.EnableRetries(0, 0)

				// Construct an instance of the GetAccessCertSettingsOptions model
				getAccessCertSettingsOptionsModel := new(mtlsv1.GetAccessCertSettingsOptions)
				getAccessCertSettingsOptionsModel.ZoneID = core.StringPtr("testString")
				getAccessCertSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := mtlsService.GetAccessCertSettingsWithContext(ctx, getAccessCertSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				mtlsService.DisableRetries()
				result, response, operationErr := mtlsService.GetAccessCertSettings(getAccessCertSettingsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = mtlsService.GetAccessCertSettingsWithContext(ctx, getAccessCertSettingsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getAccessCertSettingsPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": [{"hostname": "test.example.com", "china_network": false, "client_certificate_forwarding": true}]}`)
				}))
			})
			It(`Invoke GetAccessCertSettings successfully`, func() {
				mtlsService, serviceErr := mtlsv1.NewMtlsV1(&mtlsv1.MtlsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(mtlsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := mtlsService.GetAccessCertSettings(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetAccessCertSettingsOptions model
				getAccessCertSettingsOptionsModel := new(mtlsv1.GetAccessCertSettingsOptions)
				getAccessCertSettingsOptionsModel.ZoneID = core.StringPtr("testString")
				getAccessCertSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = mtlsService.GetAccessCertSettings(getAccessCertSettingsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetAccessCertSettings with error: Operation validation and request error`, func() {
				mtlsService, serviceErr := mtlsv1.NewMtlsV1(&mtlsv1.MtlsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(mtlsService).ToNot(BeNil())

				// Construct an instance of the GetAccessCertSettingsOptions model
				getAccessCertSettingsOptionsModel := new(mtlsv1.GetAccessCertSettingsOptions)
				getAccessCertSettingsOptionsModel.ZoneID = core.StringPtr("testString")
				getAccessCertSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := mtlsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := mtlsService.GetAccessCertSettings(getAccessCertSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetAccessCertSettingsOptions model with no property values
				getAccessCertSettingsOptionsModelNew := new(mtlsv1.GetAccessCertSettingsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = mtlsService.GetAccessCertSettings(getAccessCertSettingsOptionsModelNew)
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
			It(`Invoke GetAccessCertSettings successfully`, func() {
				mtlsService, serviceErr := mtlsv1.NewMtlsV1(&mtlsv1.MtlsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(mtlsService).ToNot(BeNil())

				// Construct an instance of the GetAccessCertSettingsOptions model
				getAccessCertSettingsOptionsModel := new(mtlsv1.GetAccessCertSettingsOptions)
				getAccessCertSettingsOptionsModel.ZoneID = core.StringPtr("testString")
				getAccessCertSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := mtlsService.GetAccessCertSettings(getAccessCertSettingsOptionsModel)
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
	Describe(`UpdateAccessCertSettings(updateAccessCertSettingsOptions *UpdateAccessCertSettingsOptions) - Operation response error`, func() {
		crn := "testString"
		updateAccessCertSettingsPath := "/v1/testString/zones/testString/access/certificates/settings"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateAccessCertSettingsPath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateAccessCertSettings with error: Operation response processing error`, func() {
				mtlsService, serviceErr := mtlsv1.NewMtlsV1(&mtlsv1.MtlsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(mtlsService).ToNot(BeNil())

				// Construct an instance of the AccessCertSettingsInputArray model
				accessCertSettingsInputArrayModel := new(mtlsv1.AccessCertSettingsInputArray)
				accessCertSettingsInputArrayModel.Hostname = core.StringPtr("test.example.com")
				accessCertSettingsInputArrayModel.ClientCertificateForwarding = core.BoolPtr(true)

				// Construct an instance of the UpdateAccessCertSettingsOptions model
				updateAccessCertSettingsOptionsModel := new(mtlsv1.UpdateAccessCertSettingsOptions)
				updateAccessCertSettingsOptionsModel.ZoneID = core.StringPtr("testString")
				updateAccessCertSettingsOptionsModel.Settings = []mtlsv1.AccessCertSettingsInputArray{*accessCertSettingsInputArrayModel}
				updateAccessCertSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := mtlsService.UpdateAccessCertSettings(updateAccessCertSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				mtlsService.EnableRetries(0, 0)
				result, response, operationErr = mtlsService.UpdateAccessCertSettings(updateAccessCertSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateAccessCertSettings(updateAccessCertSettingsOptions *UpdateAccessCertSettingsOptions)`, func() {
		crn := "testString"
		updateAccessCertSettingsPath := "/v1/testString/zones/testString/access/certificates/settings"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateAccessCertSettingsPath))
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
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": [{"hostname": "test.example.com", "china_network": false, "client_certificate_forwarding": true}]}`)
				}))
			})
			It(`Invoke UpdateAccessCertSettings successfully with retries`, func() {
				mtlsService, serviceErr := mtlsv1.NewMtlsV1(&mtlsv1.MtlsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(mtlsService).ToNot(BeNil())
				mtlsService.EnableRetries(0, 0)

				// Construct an instance of the AccessCertSettingsInputArray model
				accessCertSettingsInputArrayModel := new(mtlsv1.AccessCertSettingsInputArray)
				accessCertSettingsInputArrayModel.Hostname = core.StringPtr("test.example.com")
				accessCertSettingsInputArrayModel.ClientCertificateForwarding = core.BoolPtr(true)

				// Construct an instance of the UpdateAccessCertSettingsOptions model
				updateAccessCertSettingsOptionsModel := new(mtlsv1.UpdateAccessCertSettingsOptions)
				updateAccessCertSettingsOptionsModel.ZoneID = core.StringPtr("testString")
				updateAccessCertSettingsOptionsModel.Settings = []mtlsv1.AccessCertSettingsInputArray{*accessCertSettingsInputArrayModel}
				updateAccessCertSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := mtlsService.UpdateAccessCertSettingsWithContext(ctx, updateAccessCertSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				mtlsService.DisableRetries()
				result, response, operationErr := mtlsService.UpdateAccessCertSettings(updateAccessCertSettingsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = mtlsService.UpdateAccessCertSettingsWithContext(ctx, updateAccessCertSettingsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(updateAccessCertSettingsPath))
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
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": [{"hostname": "test.example.com", "china_network": false, "client_certificate_forwarding": true}]}`)
				}))
			})
			It(`Invoke UpdateAccessCertSettings successfully`, func() {
				mtlsService, serviceErr := mtlsv1.NewMtlsV1(&mtlsv1.MtlsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(mtlsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := mtlsService.UpdateAccessCertSettings(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the AccessCertSettingsInputArray model
				accessCertSettingsInputArrayModel := new(mtlsv1.AccessCertSettingsInputArray)
				accessCertSettingsInputArrayModel.Hostname = core.StringPtr("test.example.com")
				accessCertSettingsInputArrayModel.ClientCertificateForwarding = core.BoolPtr(true)

				// Construct an instance of the UpdateAccessCertSettingsOptions model
				updateAccessCertSettingsOptionsModel := new(mtlsv1.UpdateAccessCertSettingsOptions)
				updateAccessCertSettingsOptionsModel.ZoneID = core.StringPtr("testString")
				updateAccessCertSettingsOptionsModel.Settings = []mtlsv1.AccessCertSettingsInputArray{*accessCertSettingsInputArrayModel}
				updateAccessCertSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = mtlsService.UpdateAccessCertSettings(updateAccessCertSettingsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateAccessCertSettings with error: Operation validation and request error`, func() {
				mtlsService, serviceErr := mtlsv1.NewMtlsV1(&mtlsv1.MtlsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(mtlsService).ToNot(BeNil())

				// Construct an instance of the AccessCertSettingsInputArray model
				accessCertSettingsInputArrayModel := new(mtlsv1.AccessCertSettingsInputArray)
				accessCertSettingsInputArrayModel.Hostname = core.StringPtr("test.example.com")
				accessCertSettingsInputArrayModel.ClientCertificateForwarding = core.BoolPtr(true)

				// Construct an instance of the UpdateAccessCertSettingsOptions model
				updateAccessCertSettingsOptionsModel := new(mtlsv1.UpdateAccessCertSettingsOptions)
				updateAccessCertSettingsOptionsModel.ZoneID = core.StringPtr("testString")
				updateAccessCertSettingsOptionsModel.Settings = []mtlsv1.AccessCertSettingsInputArray{*accessCertSettingsInputArrayModel}
				updateAccessCertSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := mtlsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := mtlsService.UpdateAccessCertSettings(updateAccessCertSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateAccessCertSettingsOptions model with no property values
				updateAccessCertSettingsOptionsModelNew := new(mtlsv1.UpdateAccessCertSettingsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = mtlsService.UpdateAccessCertSettings(updateAccessCertSettingsOptionsModelNew)
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
			It(`Invoke UpdateAccessCertSettings successfully`, func() {
				mtlsService, serviceErr := mtlsv1.NewMtlsV1(&mtlsv1.MtlsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(mtlsService).ToNot(BeNil())

				// Construct an instance of the AccessCertSettingsInputArray model
				accessCertSettingsInputArrayModel := new(mtlsv1.AccessCertSettingsInputArray)
				accessCertSettingsInputArrayModel.Hostname = core.StringPtr("test.example.com")
				accessCertSettingsInputArrayModel.ClientCertificateForwarding = core.BoolPtr(true)

				// Construct an instance of the UpdateAccessCertSettingsOptions model
				updateAccessCertSettingsOptionsModel := new(mtlsv1.UpdateAccessCertSettingsOptions)
				updateAccessCertSettingsOptionsModel.ZoneID = core.StringPtr("testString")
				updateAccessCertSettingsOptionsModel.Settings = []mtlsv1.AccessCertSettingsInputArray{*accessCertSettingsInputArrayModel}
				updateAccessCertSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := mtlsService.UpdateAccessCertSettings(updateAccessCertSettingsOptionsModel)
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
	Describe(`CreateAccessOrganization(createAccessOrganizationOptions *CreateAccessOrganizationOptions) - Operation response error`, func() {
		crn := "testString"
		createAccessOrganizationPath := "/v1/testString/access/organizations"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createAccessOrganizationPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateAccessOrganization with error: Operation response processing error`, func() {
				mtlsService, serviceErr := mtlsv1.NewMtlsV1(&mtlsv1.MtlsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(mtlsService).ToNot(BeNil())

				// Construct an instance of the CreateAccessOrganizationOptions model
				createAccessOrganizationOptionsModel := new(mtlsv1.CreateAccessOrganizationOptions)
				createAccessOrganizationOptionsModel.Name = core.StringPtr("MTLS enabled")
				createAccessOrganizationOptionsModel.AuthDomain = core.StringPtr("01652b251c3ae2787110a995d8db0135.cloudflareaccess.com")
				createAccessOrganizationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := mtlsService.CreateAccessOrganization(createAccessOrganizationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				mtlsService.EnableRetries(0, 0)
				result, response, operationErr = mtlsService.CreateAccessOrganization(createAccessOrganizationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateAccessOrganization(createAccessOrganizationOptions *CreateAccessOrganizationOptions)`, func() {
		crn := "testString"
		createAccessOrganizationPath := "/v1/testString/access/organizations"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createAccessOrganizationPath))
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
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"auth_domain": "01652b251c3ae2787110a995d8db0135.cloudflareaccess.com", "name": "MTLS enabled", "login_design": {"anyKey": "anyValue"}, "created_at": "2019-08-13T16:31:42Z", "updated_at": "2019-08-13T16:31:42Z"}}`)
				}))
			})
			It(`Invoke CreateAccessOrganization successfully with retries`, func() {
				mtlsService, serviceErr := mtlsv1.NewMtlsV1(&mtlsv1.MtlsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(mtlsService).ToNot(BeNil())
				mtlsService.EnableRetries(0, 0)

				// Construct an instance of the CreateAccessOrganizationOptions model
				createAccessOrganizationOptionsModel := new(mtlsv1.CreateAccessOrganizationOptions)
				createAccessOrganizationOptionsModel.Name = core.StringPtr("MTLS enabled")
				createAccessOrganizationOptionsModel.AuthDomain = core.StringPtr("01652b251c3ae2787110a995d8db0135.cloudflareaccess.com")
				createAccessOrganizationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := mtlsService.CreateAccessOrganizationWithContext(ctx, createAccessOrganizationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				mtlsService.DisableRetries()
				result, response, operationErr := mtlsService.CreateAccessOrganization(createAccessOrganizationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = mtlsService.CreateAccessOrganizationWithContext(ctx, createAccessOrganizationOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createAccessOrganizationPath))
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
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"auth_domain": "01652b251c3ae2787110a995d8db0135.cloudflareaccess.com", "name": "MTLS enabled", "login_design": {"anyKey": "anyValue"}, "created_at": "2019-08-13T16:31:42Z", "updated_at": "2019-08-13T16:31:42Z"}}`)
				}))
			})
			It(`Invoke CreateAccessOrganization successfully`, func() {
				mtlsService, serviceErr := mtlsv1.NewMtlsV1(&mtlsv1.MtlsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(mtlsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := mtlsService.CreateAccessOrganization(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateAccessOrganizationOptions model
				createAccessOrganizationOptionsModel := new(mtlsv1.CreateAccessOrganizationOptions)
				createAccessOrganizationOptionsModel.Name = core.StringPtr("MTLS enabled")
				createAccessOrganizationOptionsModel.AuthDomain = core.StringPtr("01652b251c3ae2787110a995d8db0135.cloudflareaccess.com")
				createAccessOrganizationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = mtlsService.CreateAccessOrganization(createAccessOrganizationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateAccessOrganization with error: Operation request error`, func() {
				mtlsService, serviceErr := mtlsv1.NewMtlsV1(&mtlsv1.MtlsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(mtlsService).ToNot(BeNil())

				// Construct an instance of the CreateAccessOrganizationOptions model
				createAccessOrganizationOptionsModel := new(mtlsv1.CreateAccessOrganizationOptions)
				createAccessOrganizationOptionsModel.Name = core.StringPtr("MTLS enabled")
				createAccessOrganizationOptionsModel.AuthDomain = core.StringPtr("01652b251c3ae2787110a995d8db0135.cloudflareaccess.com")
				createAccessOrganizationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := mtlsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := mtlsService.CreateAccessOrganization(createAccessOrganizationOptionsModel)
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
			It(`Invoke CreateAccessOrganization successfully`, func() {
				mtlsService, serviceErr := mtlsv1.NewMtlsV1(&mtlsv1.MtlsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(serviceErr).To(BeNil())
				Expect(mtlsService).ToNot(BeNil())

				// Construct an instance of the CreateAccessOrganizationOptions model
				createAccessOrganizationOptionsModel := new(mtlsv1.CreateAccessOrganizationOptions)
				createAccessOrganizationOptionsModel.Name = core.StringPtr("MTLS enabled")
				createAccessOrganizationOptionsModel.AuthDomain = core.StringPtr("01652b251c3ae2787110a995d8db0135.cloudflareaccess.com")
				createAccessOrganizationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := mtlsService.CreateAccessOrganization(createAccessOrganizationOptionsModel)
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
			mtlsService, _ := mtlsv1.NewMtlsV1(&mtlsv1.MtlsV1Options{
				URL:           "http://mtlsv1modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
				Crn:           core.StringPtr(crn),
			})
			It(`Invoke NewCreateAccessApplicationOptions successfully`, func() {
				// Construct an instance of the CreateAccessApplicationOptions model
				zoneID := "testString"
				createAccessApplicationOptionsModel := mtlsService.NewCreateAccessApplicationOptions(zoneID)
				createAccessApplicationOptionsModel.SetZoneID("testString")
				createAccessApplicationOptionsModel.SetName("mtls-test-app")
				createAccessApplicationOptionsModel.SetDomain("test.example.com")
				createAccessApplicationOptionsModel.SetSessionDuration("24h")
				createAccessApplicationOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createAccessApplicationOptionsModel).ToNot(BeNil())
				Expect(createAccessApplicationOptionsModel.ZoneID).To(Equal(core.StringPtr("testString")))
				Expect(createAccessApplicationOptionsModel.Name).To(Equal(core.StringPtr("mtls-test-app")))
				Expect(createAccessApplicationOptionsModel.Domain).To(Equal(core.StringPtr("test.example.com")))
				Expect(createAccessApplicationOptionsModel.SessionDuration).To(Equal(core.StringPtr("24h")))
				Expect(createAccessApplicationOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateAccessCertificateOptions successfully`, func() {
				// Construct an instance of the CreateAccessCertificateOptions model
				zoneID := "testString"
				createAccessCertificateOptionsModel := mtlsService.NewCreateAccessCertificateOptions(zoneID)
				createAccessCertificateOptionsModel.SetZoneID("testString")
				createAccessCertificateOptionsModel.SetName("test-cert")
				createAccessCertificateOptionsModel.SetCertificate("-----BEGIN CERTIFICATE-----MIIGAjCCA+qgAwIBAgIJAI7kymlF7CWT...N4RI7KKB7nikiuUf8vhULKy5IX10DrUtmu/B-----END CERTIFICATE-----")
				createAccessCertificateOptionsModel.SetAssociatedHostnames([]string{"test.example.com"})
				createAccessCertificateOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createAccessCertificateOptionsModel).ToNot(BeNil())
				Expect(createAccessCertificateOptionsModel.ZoneID).To(Equal(core.StringPtr("testString")))
				Expect(createAccessCertificateOptionsModel.Name).To(Equal(core.StringPtr("test-cert")))
				Expect(createAccessCertificateOptionsModel.Certificate).To(Equal(core.StringPtr("-----BEGIN CERTIFICATE-----MIIGAjCCA+qgAwIBAgIJAI7kymlF7CWT...N4RI7KKB7nikiuUf8vhULKy5IX10DrUtmu/B-----END CERTIFICATE-----")))
				Expect(createAccessCertificateOptionsModel.AssociatedHostnames).To(Equal([]string{"test.example.com"}))
				Expect(createAccessCertificateOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateAccessOrganizationOptions successfully`, func() {
				// Construct an instance of the CreateAccessOrganizationOptions model
				createAccessOrganizationOptionsModel := mtlsService.NewCreateAccessOrganizationOptions()
				createAccessOrganizationOptionsModel.SetName("MTLS enabled")
				createAccessOrganizationOptionsModel.SetAuthDomain("01652b251c3ae2787110a995d8db0135.cloudflareaccess.com")
				createAccessOrganizationOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createAccessOrganizationOptionsModel).ToNot(BeNil())
				Expect(createAccessOrganizationOptionsModel.Name).To(Equal(core.StringPtr("MTLS enabled")))
				Expect(createAccessOrganizationOptionsModel.AuthDomain).To(Equal(core.StringPtr("01652b251c3ae2787110a995d8db0135.cloudflareaccess.com")))
				Expect(createAccessOrganizationOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateAccessPolicyOptions successfully`, func() {
				// Construct an instance of the PolicyRulePolicyCertRule model
				policyRuleModel := new(mtlsv1.PolicyRulePolicyCertRule)
				Expect(policyRuleModel).ToNot(BeNil())
				policyRuleModel.Certificate = map[string]interface{}{"anyKey": "anyValue"}
				Expect(policyRuleModel.Certificate).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))

				// Construct an instance of the CreateAccessPolicyOptions model
				zoneID := "testString"
				appID := "testString"
				createAccessPolicyOptionsModel := mtlsService.NewCreateAccessPolicyOptions(zoneID, appID)
				createAccessPolicyOptionsModel.SetZoneID("testString")
				createAccessPolicyOptionsModel.SetAppID("testString")
				createAccessPolicyOptionsModel.SetName("mtls-test-policy")
				createAccessPolicyOptionsModel.SetDecision("non_identity")
				createAccessPolicyOptionsModel.SetInclude([]mtlsv1.PolicyRuleIntf{policyRuleModel})
				createAccessPolicyOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createAccessPolicyOptionsModel).ToNot(BeNil())
				Expect(createAccessPolicyOptionsModel.ZoneID).To(Equal(core.StringPtr("testString")))
				Expect(createAccessPolicyOptionsModel.AppID).To(Equal(core.StringPtr("testString")))
				Expect(createAccessPolicyOptionsModel.Name).To(Equal(core.StringPtr("mtls-test-policy")))
				Expect(createAccessPolicyOptionsModel.Decision).To(Equal(core.StringPtr("non_identity")))
				Expect(createAccessPolicyOptionsModel.Include).To(Equal([]mtlsv1.PolicyRuleIntf{policyRuleModel}))
				Expect(createAccessPolicyOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteAccessApplicationOptions successfully`, func() {
				// Construct an instance of the DeleteAccessApplicationOptions model
				zoneID := "testString"
				appID := "testString"
				deleteAccessApplicationOptionsModel := mtlsService.NewDeleteAccessApplicationOptions(zoneID, appID)
				deleteAccessApplicationOptionsModel.SetZoneID("testString")
				deleteAccessApplicationOptionsModel.SetAppID("testString")
				deleteAccessApplicationOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteAccessApplicationOptionsModel).ToNot(BeNil())
				Expect(deleteAccessApplicationOptionsModel.ZoneID).To(Equal(core.StringPtr("testString")))
				Expect(deleteAccessApplicationOptionsModel.AppID).To(Equal(core.StringPtr("testString")))
				Expect(deleteAccessApplicationOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteAccessCertificateOptions successfully`, func() {
				// Construct an instance of the DeleteAccessCertificateOptions model
				zoneID := "testString"
				certID := "testString"
				deleteAccessCertificateOptionsModel := mtlsService.NewDeleteAccessCertificateOptions(zoneID, certID)
				deleteAccessCertificateOptionsModel.SetZoneID("testString")
				deleteAccessCertificateOptionsModel.SetCertID("testString")
				deleteAccessCertificateOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteAccessCertificateOptionsModel).ToNot(BeNil())
				Expect(deleteAccessCertificateOptionsModel.ZoneID).To(Equal(core.StringPtr("testString")))
				Expect(deleteAccessCertificateOptionsModel.CertID).To(Equal(core.StringPtr("testString")))
				Expect(deleteAccessCertificateOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteAccessPolicyOptions successfully`, func() {
				// Construct an instance of the DeleteAccessPolicyOptions model
				zoneID := "testString"
				appID := "testString"
				policyID := "testString"
				deleteAccessPolicyOptionsModel := mtlsService.NewDeleteAccessPolicyOptions(zoneID, appID, policyID)
				deleteAccessPolicyOptionsModel.SetZoneID("testString")
				deleteAccessPolicyOptionsModel.SetAppID("testString")
				deleteAccessPolicyOptionsModel.SetPolicyID("testString")
				deleteAccessPolicyOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteAccessPolicyOptionsModel).ToNot(BeNil())
				Expect(deleteAccessPolicyOptionsModel.ZoneID).To(Equal(core.StringPtr("testString")))
				Expect(deleteAccessPolicyOptionsModel.AppID).To(Equal(core.StringPtr("testString")))
				Expect(deleteAccessPolicyOptionsModel.PolicyID).To(Equal(core.StringPtr("testString")))
				Expect(deleteAccessPolicyOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetAccessApplicationOptions successfully`, func() {
				// Construct an instance of the GetAccessApplicationOptions model
				zoneID := "testString"
				appID := "testString"
				getAccessApplicationOptionsModel := mtlsService.NewGetAccessApplicationOptions(zoneID, appID)
				getAccessApplicationOptionsModel.SetZoneID("testString")
				getAccessApplicationOptionsModel.SetAppID("testString")
				getAccessApplicationOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getAccessApplicationOptionsModel).ToNot(BeNil())
				Expect(getAccessApplicationOptionsModel.ZoneID).To(Equal(core.StringPtr("testString")))
				Expect(getAccessApplicationOptionsModel.AppID).To(Equal(core.StringPtr("testString")))
				Expect(getAccessApplicationOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetAccessCertSettingsOptions successfully`, func() {
				// Construct an instance of the GetAccessCertSettingsOptions model
				zoneID := "testString"
				getAccessCertSettingsOptionsModel := mtlsService.NewGetAccessCertSettingsOptions(zoneID)
				getAccessCertSettingsOptionsModel.SetZoneID("testString")
				getAccessCertSettingsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getAccessCertSettingsOptionsModel).ToNot(BeNil())
				Expect(getAccessCertSettingsOptionsModel.ZoneID).To(Equal(core.StringPtr("testString")))
				Expect(getAccessCertSettingsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetAccessCertificateOptions successfully`, func() {
				// Construct an instance of the GetAccessCertificateOptions model
				zoneID := "testString"
				certID := "testString"
				getAccessCertificateOptionsModel := mtlsService.NewGetAccessCertificateOptions(zoneID, certID)
				getAccessCertificateOptionsModel.SetZoneID("testString")
				getAccessCertificateOptionsModel.SetCertID("testString")
				getAccessCertificateOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getAccessCertificateOptionsModel).ToNot(BeNil())
				Expect(getAccessCertificateOptionsModel.ZoneID).To(Equal(core.StringPtr("testString")))
				Expect(getAccessCertificateOptionsModel.CertID).To(Equal(core.StringPtr("testString")))
				Expect(getAccessCertificateOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetAccessPolicyOptions successfully`, func() {
				// Construct an instance of the GetAccessPolicyOptions model
				zoneID := "testString"
				appID := "testString"
				policyID := "testString"
				getAccessPolicyOptionsModel := mtlsService.NewGetAccessPolicyOptions(zoneID, appID, policyID)
				getAccessPolicyOptionsModel.SetZoneID("testString")
				getAccessPolicyOptionsModel.SetAppID("testString")
				getAccessPolicyOptionsModel.SetPolicyID("testString")
				getAccessPolicyOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getAccessPolicyOptionsModel).ToNot(BeNil())
				Expect(getAccessPolicyOptionsModel.ZoneID).To(Equal(core.StringPtr("testString")))
				Expect(getAccessPolicyOptionsModel.AppID).To(Equal(core.StringPtr("testString")))
				Expect(getAccessPolicyOptionsModel.PolicyID).To(Equal(core.StringPtr("testString")))
				Expect(getAccessPolicyOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListAccessApplicationsOptions successfully`, func() {
				// Construct an instance of the ListAccessApplicationsOptions model
				zoneID := "testString"
				listAccessApplicationsOptionsModel := mtlsService.NewListAccessApplicationsOptions(zoneID)
				listAccessApplicationsOptionsModel.SetZoneID("testString")
				listAccessApplicationsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listAccessApplicationsOptionsModel).ToNot(BeNil())
				Expect(listAccessApplicationsOptionsModel.ZoneID).To(Equal(core.StringPtr("testString")))
				Expect(listAccessApplicationsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListAccessCertificatesOptions successfully`, func() {
				// Construct an instance of the ListAccessCertificatesOptions model
				zoneID := "testString"
				listAccessCertificatesOptionsModel := mtlsService.NewListAccessCertificatesOptions(zoneID)
				listAccessCertificatesOptionsModel.SetZoneID("testString")
				listAccessCertificatesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listAccessCertificatesOptionsModel).ToNot(BeNil())
				Expect(listAccessCertificatesOptionsModel.ZoneID).To(Equal(core.StringPtr("testString")))
				Expect(listAccessCertificatesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListAccessPoliciesOptions successfully`, func() {
				// Construct an instance of the ListAccessPoliciesOptions model
				zoneID := "testString"
				appID := "testString"
				listAccessPoliciesOptionsModel := mtlsService.NewListAccessPoliciesOptions(zoneID, appID)
				listAccessPoliciesOptionsModel.SetZoneID("testString")
				listAccessPoliciesOptionsModel.SetAppID("testString")
				listAccessPoliciesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listAccessPoliciesOptionsModel).ToNot(BeNil())
				Expect(listAccessPoliciesOptionsModel.ZoneID).To(Equal(core.StringPtr("testString")))
				Expect(listAccessPoliciesOptionsModel.AppID).To(Equal(core.StringPtr("testString")))
				Expect(listAccessPoliciesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewPolicyCnRuleCommonName successfully`, func() {
				commonName := "Access Testing CA"
				_model, err := mtlsService.NewPolicyCnRuleCommonName(commonName)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewUpdateAccessApplicationOptions successfully`, func() {
				// Construct an instance of the UpdateAccessApplicationOptions model
				zoneID := "testString"
				appID := "testString"
				updateAccessApplicationOptionsModel := mtlsService.NewUpdateAccessApplicationOptions(zoneID, appID)
				updateAccessApplicationOptionsModel.SetZoneID("testString")
				updateAccessApplicationOptionsModel.SetAppID("testString")
				updateAccessApplicationOptionsModel.SetName("mtls-test-app")
				updateAccessApplicationOptionsModel.SetDomain("test.example.com")
				updateAccessApplicationOptionsModel.SetSessionDuration("24h")
				updateAccessApplicationOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateAccessApplicationOptionsModel).ToNot(BeNil())
				Expect(updateAccessApplicationOptionsModel.ZoneID).To(Equal(core.StringPtr("testString")))
				Expect(updateAccessApplicationOptionsModel.AppID).To(Equal(core.StringPtr("testString")))
				Expect(updateAccessApplicationOptionsModel.Name).To(Equal(core.StringPtr("mtls-test-app")))
				Expect(updateAccessApplicationOptionsModel.Domain).To(Equal(core.StringPtr("test.example.com")))
				Expect(updateAccessApplicationOptionsModel.SessionDuration).To(Equal(core.StringPtr("24h")))
				Expect(updateAccessApplicationOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateAccessCertSettingsOptions successfully`, func() {
				// Construct an instance of the AccessCertSettingsInputArray model
				accessCertSettingsInputArrayModel := new(mtlsv1.AccessCertSettingsInputArray)
				Expect(accessCertSettingsInputArrayModel).ToNot(BeNil())
				accessCertSettingsInputArrayModel.Hostname = core.StringPtr("test.example.com")
				accessCertSettingsInputArrayModel.ClientCertificateForwarding = core.BoolPtr(true)
				Expect(accessCertSettingsInputArrayModel.Hostname).To(Equal(core.StringPtr("test.example.com")))
				Expect(accessCertSettingsInputArrayModel.ClientCertificateForwarding).To(Equal(core.BoolPtr(true)))

				// Construct an instance of the UpdateAccessCertSettingsOptions model
				zoneID := "testString"
				updateAccessCertSettingsOptionsModel := mtlsService.NewUpdateAccessCertSettingsOptions(zoneID)
				updateAccessCertSettingsOptionsModel.SetZoneID("testString")
				updateAccessCertSettingsOptionsModel.SetSettings([]mtlsv1.AccessCertSettingsInputArray{*accessCertSettingsInputArrayModel})
				updateAccessCertSettingsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateAccessCertSettingsOptionsModel).ToNot(BeNil())
				Expect(updateAccessCertSettingsOptionsModel.ZoneID).To(Equal(core.StringPtr("testString")))
				Expect(updateAccessCertSettingsOptionsModel.Settings).To(Equal([]mtlsv1.AccessCertSettingsInputArray{*accessCertSettingsInputArrayModel}))
				Expect(updateAccessCertSettingsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateAccessCertificateOptions successfully`, func() {
				// Construct an instance of the UpdateAccessCertificateOptions model
				zoneID := "testString"
				certID := "testString"
				updateAccessCertificateOptionsModel := mtlsService.NewUpdateAccessCertificateOptions(zoneID, certID)
				updateAccessCertificateOptionsModel.SetZoneID("testString")
				updateAccessCertificateOptionsModel.SetCertID("testString")
				updateAccessCertificateOptionsModel.SetName("test-cert")
				updateAccessCertificateOptionsModel.SetAssociatedHostnames([]string{"test.example.com"})
				updateAccessCertificateOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateAccessCertificateOptionsModel).ToNot(BeNil())
				Expect(updateAccessCertificateOptionsModel.ZoneID).To(Equal(core.StringPtr("testString")))
				Expect(updateAccessCertificateOptionsModel.CertID).To(Equal(core.StringPtr("testString")))
				Expect(updateAccessCertificateOptionsModel.Name).To(Equal(core.StringPtr("test-cert")))
				Expect(updateAccessCertificateOptionsModel.AssociatedHostnames).To(Equal([]string{"test.example.com"}))
				Expect(updateAccessCertificateOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateAccessPolicyOptions successfully`, func() {
				// Construct an instance of the PolicyRulePolicyCertRule model
				policyRuleModel := new(mtlsv1.PolicyRulePolicyCertRule)
				Expect(policyRuleModel).ToNot(BeNil())
				policyRuleModel.Certificate = map[string]interface{}{"anyKey": "anyValue"}
				Expect(policyRuleModel.Certificate).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))

				// Construct an instance of the UpdateAccessPolicyOptions model
				zoneID := "testString"
				appID := "testString"
				policyID := "testString"
				updateAccessPolicyOptionsModel := mtlsService.NewUpdateAccessPolicyOptions(zoneID, appID, policyID)
				updateAccessPolicyOptionsModel.SetZoneID("testString")
				updateAccessPolicyOptionsModel.SetAppID("testString")
				updateAccessPolicyOptionsModel.SetPolicyID("testString")
				updateAccessPolicyOptionsModel.SetName("mtls-test-policy")
				updateAccessPolicyOptionsModel.SetDecision("non_identity")
				updateAccessPolicyOptionsModel.SetInclude([]mtlsv1.PolicyRuleIntf{policyRuleModel})
				updateAccessPolicyOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateAccessPolicyOptionsModel).ToNot(BeNil())
				Expect(updateAccessPolicyOptionsModel.ZoneID).To(Equal(core.StringPtr("testString")))
				Expect(updateAccessPolicyOptionsModel.AppID).To(Equal(core.StringPtr("testString")))
				Expect(updateAccessPolicyOptionsModel.PolicyID).To(Equal(core.StringPtr("testString")))
				Expect(updateAccessPolicyOptionsModel.Name).To(Equal(core.StringPtr("mtls-test-policy")))
				Expect(updateAccessPolicyOptionsModel.Decision).To(Equal(core.StringPtr("non_identity")))
				Expect(updateAccessPolicyOptionsModel.Include).To(Equal([]mtlsv1.PolicyRuleIntf{policyRuleModel}))
				Expect(updateAccessPolicyOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewAccessCertSettingsInputArray successfully`, func() {
				hostname := "test.example.com"
				clientCertificateForwarding := true
				_model, err := mtlsService.NewAccessCertSettingsInputArray(hostname, clientCertificateForwarding)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewPolicyRulePolicyCnRule successfully`, func() {
				var commonName *mtlsv1.PolicyCnRuleCommonName = nil
				_, err := mtlsService.NewPolicyRulePolicyCnRule(commonName)
				Expect(err).ToNot(BeNil())
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
	return ioutil.NopCloser(bytes.NewReader([]byte(mockData)))
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
