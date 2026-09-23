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

//go:build integration

package aisecurityforappsv1_test

import (
	"fmt"
	"net/http"
	"os"

	"github.com/IBM/go-sdk-core/v5/core"
	. "github.com/IBM/networking-go-sdk/aisecurityforappsv1"
	"github.com/joho/godotenv"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const configFile = "../cis.env"

var configLoaded bool = true
var authenticationSucceeded bool = true

func shouldSkipTest() {
	if !configLoaded {
		Skip("External configuration is not available, skipping...")
	}

	if !authenticationSucceeded {
		Skip("Authentication failed. Check external configuration...")
	}
}

var _ = Describe(`AiSecurityForAppsV1 Integration Tests`, func() {

	defer GinkgoRecover()
	if _, err := os.Stat(configFile); err != nil {
		configLoaded = false
	}

	err := godotenv.Load(configFile)
	if err != nil {
		configLoaded = false
	}

	authenticator := &core.IamAuthenticator{
		ApiKey: os.Getenv("CIS_SERVICES_APIKEY"),
		URL:    os.Getenv("CIS_SERVICES_AUTH_URL"),
	}
	authErr := authenticator.Authenticate(&http.Request{
		Header: http.Header{},
	})
	if authErr != nil {
		authenticationSucceeded = false
		fmt.Println("Authentication error during setup: ", authErr)
	}
	serviceURL := os.Getenv("API_ENDPOINT")
	crn := os.Getenv("CRN")
	zoneID := os.Getenv("ZONE_ID")

	globalOptions := &AiSecurityForAppsV1Options{
		ServiceName:    DefaultServiceName,
		URL:            serviceURL,
		Authenticator:  authenticator,
		Crn:            &crn,
		ZoneIdentifier: &zoneID,
	}

	service, serviceErr := NewAiSecurityForAppsV1(globalOptions)
	if serviceErr != nil {
		fmt.Println(serviceErr)
	}

	Describe(`AI Security Settings`, func() {
		Context(`Get and Update AI Security Settings`, func() {
			It(`Get AI Security Settings`, func() {
				shouldSkipTest()

				getOptions := service.NewGetAiSecuritySettingsOptions()
				result, response, err := service.GetAiSecuritySettings(getOptions)
				Expect(err).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
				Expect(*result.Success).To(BeTrue())
			})

			It(`Update AI Security Settings`, func() {
				shouldSkipTest()

				// Get current value
				getOptions := service.NewGetAiSecuritySettingsOptions()
				getResult, _, getErr := service.GetAiSecuritySettings(getOptions)
				Expect(getErr).To(BeNil())
				Expect(getResult).ToNot(BeNil())

				currentEnabled := *getResult.Result.Enabled
				newEnabled := !currentEnabled

				updateOptions := service.NewReplaceZoneAiSecuritySettingsOptions().
					SetEnabled(newEnabled)
				result, response, err := service.ReplaceZoneAiSecuritySettings(updateOptions)
				Expect(err).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
				Expect(*result.Success).To(BeTrue())

				// Restore original value
				restoreOptions := service.NewReplaceZoneAiSecuritySettingsOptions().
					SetEnabled(currentEnabled)
				_, _, restoreErr := service.ReplaceZoneAiSecuritySettings(restoreOptions)
				Expect(restoreErr).To(BeNil())
			})
		})
	})

	Describe(`API Gateway Discovery`, func() {
		Context(`Get API Gateway Discovery`, func() {
			It(`Get API Gateway Discovery`, func() {
				shouldSkipTest()

				getOptions := service.NewGetApiGatewayDiscoveryOptions()
				result, response, err := service.GetApiGatewayDiscovery(getOptions)
				Expect(err).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
				Expect(*result.Success).To(BeTrue())
			})
		})

		Context(`List API Gateway Discovery Operations`, func() {
			It(`List API Gateway Discovery Operations`, func() {
				shouldSkipTest()

				listOptions := service.NewListApiGatewayDiscoveryOperationsOptions()
				result, response, err := service.ListApiGatewayDiscoveryOperations(listOptions)
				Expect(err).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
				Expect(*result.Success).To(BeTrue())
			})
		})
	})

	Describe(`API Gateway Schemas`, func() {
		Context(`Get API Gateway Schemas`, func() {
			It(`Get API Gateway Schemas`, func() {
				shouldSkipTest()

				getOptions := service.NewGetApiGatewaySchemasOptions()
				result, response, err := service.GetApiGatewaySchemas(getOptions)
				Expect(err).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
				Expect(*result.Success).To(BeTrue())
			})
		})
	})

	Describe(`API Gateway Operations`, func() {
		Context(`Create single operation, retrieve it, and delete it`, func() {
			It(`Create a single API Gateway Operation, retrieve it by operation ID, then delete and verify deletion`, func() {
				shouldSkipTest()

				// Create a single operation
				createOptions := service.NewCreateApiGatewayOperationItemOptions().
					SetMethod("POST").
					SetHost("api.example.com").
					SetEndpoint("/v1/messages")

				createResult, createResponse, createErr := service.CreateApiGatewayOperationItem(createOptions)
				Expect(createErr).To(BeNil())
				Expect(createResponse).ToNot(BeNil())
				Expect(createResult).ToNot(BeNil())
				Expect(*createResult.Success).To(BeTrue())
				Expect(createResult.Result).ToNot(BeNil())
				Expect(createResult.Result.OperationID).ToNot(BeNil())

				operationID := *createResult.Result.OperationID

				// Safety cleanup: runs if any subsequent expectation fails
				defer func() {
					if operationID != "" {
						deleteOptions := service.NewDeleteZoneApiGatewayOperationOptions(operationID)
						_, _ = service.DeleteZoneApiGatewayOperation(deleteOptions)
					}
				}()

				// Retrieve the operation by ID
				getOptions := service.NewGetZoneApiGatewayOperationOptions(operationID)
				getResult, getResponse, getErr := service.GetZoneApiGatewayOperation(getOptions)
				Expect(getErr).To(BeNil())
				Expect(getResponse).ToNot(BeNil())
				Expect(getResult).ToNot(BeNil())
				Expect(*getResult.Success).To(BeTrue())
				Expect(getResult.Result).ToNot(BeNil())
				Expect(*getResult.Result.OperationID).To(Equal(operationID))

				// Explicit Delete: verifies delete API response on normal flow
				deleteOptions := service.NewDeleteZoneApiGatewayOperationOptions(operationID)
				deleteResponse, deleteErr := service.DeleteZoneApiGatewayOperation(deleteOptions)
				Expect(deleteErr).To(BeNil())
				Expect(deleteResponse).ToNot(BeNil())

				// Reset operationID so defer does not attempt a redundant delete
				operationID = ""

				// Verify deletion — the operation should no longer be retrievable
				verifyGetOptions := service.NewGetZoneApiGatewayOperationOptions(*createResult.Result.OperationID)
				_, verifyResponse, verifyErr := service.GetZoneApiGatewayOperation(verifyGetOptions)
				Expect(verifyErr).ToNot(BeNil())
				Expect(verifyResponse.StatusCode).To(Equal(404))
			})
		})

		Context(`Create bulk operations, retrieve one, update labels, then delete all`, func() {
			It(`Create 3 API Gateway operations in bulk, retrieve one, update labels, and clean up`, func() {
				shouldSkipTest()

				// Build 3 operations for the bulk create call
				op1, err := service.NewApiGatewayOperation("GET", "api.example.com", "/v2/users")
				Expect(err).To(BeNil())
				op2, err := service.NewApiGatewayOperation("POST", "api.example.com", "/v2/orders")
				Expect(err).To(BeNil())
				op3, err := service.NewApiGatewayOperation("DELETE", "api.example.com", "/v2/sessions")
				Expect(err).To(BeNil())

				bulkCreateOptions := service.NewCreateZoneApiGatewayOperationOptions().
					SetApiGatewayOperation([]ApiGatewayOperation{*op1, *op2, *op3})

				bulkResult, bulkResponse, bulkErr := service.CreateZoneApiGatewayOperation(bulkCreateOptions)
				Expect(bulkErr).To(BeNil())
				Expect(bulkResponse).ToNot(BeNil())
				Expect(bulkResult).ToNot(BeNil())
				Expect(*bulkResult.Success).To(BeTrue())
				Expect(len(bulkResult.Result)).To(Equal(3))

				// Collect all operation IDs
				bulkOperationIDs := []string{}
				for _, item := range bulkResult.Result {
					Expect(item.OperationID).ToNot(BeNil())
					bulkOperationIDs = append(bulkOperationIDs, *item.OperationID)
				}
				Expect(len(bulkOperationIDs)).To(Equal(3))

				// Safety cleanup: runs if any subsequent expectation fails
				defer func() {
					for _, opID := range bulkOperationIDs {
						deleteOptions := service.NewDeleteZoneApiGatewayOperationOptions(opID)
						_, _ = service.DeleteZoneApiGatewayOperation(deleteOptions)
					}
				}()

				// --- Step 1: Retrieve the second operation by its operation ID ---
				secondOperationID := bulkOperationIDs[1]
				getOptions := service.NewGetZoneApiGatewayOperationOptions(secondOperationID)
				getResult, getResponse, getErr := service.GetZoneApiGatewayOperation(getOptions)
				Expect(getErr).To(BeNil())
				Expect(getResponse).ToNot(BeNil())
				Expect(getResult).ToNot(BeNil())
				Expect(*getResult.Success).To(BeTrue())
				Expect(*getResult.Result.OperationID).To(Equal(secondOperationID))

				// --- Step 2: Add managed label to all 3 operations ---
				const managedLabel = "cf-llm"
				addSelector := &ApiGatewayOperationsLabelsInputSelector{
					Include: &ApiGatewayOperationsLabelsInputSelectorInclude{
						OperationIds: bulkOperationIDs,
					},
				}
				addLabelsOptions := service.NewUpdateApiGatewayOperationLabelsOptions().
					SetManaged(&ApiGatewayOperationsLabelsInputManaged{
						Labels: []string{managedLabel},
					}).
					SetSelector(addSelector)

				addResult, addResponse, addErr := service.UpdateApiGatewayOperationLabels(addLabelsOptions)
				Expect(addErr).To(BeNil())
				Expect(addResponse).ToNot(BeNil())
				Expect(addResult).ToNot(BeNil())
				Expect(*addResult.Success).To(BeTrue())
				Expect(len(addResult.Result)).To(Equal(3))

				for _, item := range addResult.Result {
					Expect(item.Labels).To(ContainElement(HaveKeyWithValue("name", managedLabel)))
				}

				// --- Step 3: Remove the first operation ID from the label set ---
				remainingIDs := bulkOperationIDs[1:]
				removeSelector := &ApiGatewayOperationsLabelsInputSelector{
					Include: &ApiGatewayOperationsLabelsInputSelectorInclude{
						OperationIds: remainingIDs,
					},
				}
				removeLabelsOptions := service.NewUpdateApiGatewayOperationLabelsOptions().
					SetManaged(&ApiGatewayOperationsLabelsInputManaged{
						Labels: []string{managedLabel},
					}).
					SetSelector(removeSelector)

				removeResult, removeResponse, removeErr := service.UpdateApiGatewayOperationLabels(removeLabelsOptions)
				Expect(removeErr).To(BeNil())
				Expect(removeResponse).ToNot(BeNil())
				Expect(removeResult).ToNot(BeNil())
				Expect(*removeResult.Success).To(BeTrue())

				for _, item := range removeResult.Result {
					Expect(item.Labels).To(ContainElement(HaveKeyWithValue("name", managedLabel)))
				}

				// --- Step 4: Explicit delete of all operations and verify API responses ---
				for _, opID := range bulkOperationIDs {
					deleteOptions := service.NewDeleteZoneApiGatewayOperationOptions(opID)
					deleteResponse, deleteErr := service.DeleteZoneApiGatewayOperation(deleteOptions)
					Expect(deleteErr).To(BeNil())
					Expect(deleteResponse).ToNot(BeNil())
				}

				// Reset slice so defer does not attempt redundant deletes on normal exit
				bulkOperationIDs = nil
			})
		})
	})
})
