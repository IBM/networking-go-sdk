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

	Describe(`API Gateway Discovery Operations Update`, func() {
		Context(`Update API Gateway Discovery Operation`, func() {
			It(`Update API Gateway Discovery Operation`, func() {
				shouldSkipTest()

				updateOptions := service.NewUpdateZoneApiGatewayDiscoveryOperationOptions()
				result, response, err := service.UpdateZoneApiGatewayDiscoveryOperation(updateOptions)
				Expect(err).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
				Expect(*result.Success).To(BeTrue())
			})
		})
	})

	Describe(`API Gateway Operations List and Delete`, func() {
		Context(`List API Gateway Operations`, func() {
			It(`List API Gateway Operations`, func() {
				shouldSkipTest()

				listOptions := service.NewListZoneApiGatewayOperationsOptions()
				result, response, err := service.ListZoneApiGatewayOperations(listOptions)
				Expect(err).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
				Expect(*result.Success).To(BeTrue())
			})
		})

		Context(`Delete Zone API Gateway Operations (bulk)`, func() {
			It(`Delete Zone API Gateway Operations`, func() {
				shouldSkipTest()

				deleteOptions := service.NewDeleteZoneApiGatewayOperationsOptions()
				result, response, err := service.DeleteZoneApiGatewayOperations(deleteOptions)
				Expect(err).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
				Expect(*result.Success).To(BeTrue())
			})
		})
	})

	Describe(`API Gateway Operation Labels — per-operation`, func() {
		Context(`Add, replace, and remove labels on a single operation`, func() {
			It(`Create operation, manage per-operation labels, then clean up`, func() {
				shouldSkipTest()

				// Create a single operation to work with
				createOptions := service.NewCreateApiGatewayOperationItemOptions().
					SetMethod("POST").
					SetHost("api.example.com").
					SetEndpoint("/v1/label-test")

				createResult, createResponse, createErr := service.CreateApiGatewayOperationItem(createOptions)
				Expect(createErr).To(BeNil())
				Expect(createResponse).ToNot(BeNil())
				Expect(createResult).ToNot(BeNil())
				Expect(*createResult.Success).To(BeTrue())
				Expect(createResult.Result.OperationID).ToNot(BeNil())

				operationID := *createResult.Result.OperationID

				defer func() {
					if operationID != "" {
						deleteOptions := service.NewDeleteZoneApiGatewayOperationOptions(operationID)
						_, _ = service.DeleteZoneApiGatewayOperation(deleteOptions)
					}
				}()

				// Step 1: Add labels (create_api_gateway_operation_labels)
				addLabelsOptions := service.NewCreateApiGatewayOperationLabelsOptions(operationID).
					SetManaged([]string{"cf-llm"})
				addResult, addResponse, addErr := service.CreateApiGatewayOperationLabels(addLabelsOptions)
				Expect(addErr).To(BeNil())
				Expect(addResponse).ToNot(BeNil())
				Expect(addResult).ToNot(BeNil())
				Expect(*addResult.Success).To(BeTrue())

				// Step 2: Replace labels (set_zone_api_gateway_operation_labels)
				setLabelsOptions := service.NewSetZoneApiGatewayOperationLabelsOptions(operationID).
					SetManaged([]string{"cf-llm"})
				setResult, setResponse, setErr := service.SetZoneApiGatewayOperationLabels(setLabelsOptions)
				Expect(setErr).To(BeNil())
				Expect(setResponse).ToNot(BeNil())
				Expect(setResult).ToNot(BeNil())
				Expect(*setResult.Success).To(BeTrue())

				// Step 3: Bulk-replace labels on a selector (replace_zone_api_gateway_operation_labels)
				replaceSelector := &ApiGatewayOperationsLabelsBulkInputSelector{
					Include: &ApiGatewayOperationsLabelsBulkInputSelectorInclude{
						OperationIds: []string{operationID},
					},
				}
				replaceLabelsOptions := service.NewReplaceZoneApiGatewayOperationLabelsOptions().
					SetManaged(&ApiGatewayOperationsLabelsBulkInputManaged{
						Labels: []string{"cf-llm"},
					}).
					SetSelector(replaceSelector)
				replaceResult, replaceResponse, replaceErr := service.ReplaceZoneApiGatewayOperationLabels(replaceLabelsOptions)
				Expect(replaceErr).To(BeNil())
				Expect(replaceResponse).ToNot(BeNil())
				Expect(replaceResult).ToNot(BeNil())
				Expect(*replaceResult.Success).To(BeTrue())

				// Step 4: Remove label from single operation (unset_zone_api_gateway_operation_labels)
				unsetLabelsOptions := service.NewUnsetZoneApiGatewayOperationLabelsOptions(operationID).
					SetManaged([]string{"cf-llm"})
				unsetResult, unsetResponse, unsetErr := service.UnsetZoneApiGatewayOperationLabels(unsetLabelsOptions)
				Expect(unsetErr).To(BeNil())
				Expect(unsetResponse).ToNot(BeNil())
				Expect(unsetResult).ToNot(BeNil())
				Expect(*unsetResult.Success).To(BeTrue())

				// Step 5: Bulk-delete labels on a selector (delete_zone_api_gateway_operation_labels)
				deleteSelector := &ApiGatewayOperationsLabelsBulkDeleteInputSelector{
					Include: &ApiGatewayOperationsLabelsBulkDeleteInputSelectorInclude{
						OperationIds: []string{operationID},
					},
				}
				deleteLabelsOptions := service.NewDeleteZoneApiGatewayOperationLabelsOptions().
					SetManaged(&ApiGatewayOperationsLabelsBulkDeleteInputManaged{
						Labels: []string{"cf-llm"},
					}).
					SetSelector(deleteSelector)
				deleteLabelsResult, deleteLabelsResponse, deleteLabelsErr := service.DeleteZoneApiGatewayOperationLabels(deleteLabelsOptions)
				Expect(deleteLabelsErr).To(BeNil())
				Expect(deleteLabelsResponse).ToNot(BeNil())
				Expect(deleteLabelsResult).ToNot(BeNil())
				Expect(*deleteLabelsResult.Success).To(BeTrue())

				// Cleanup
				deleteOptions := service.NewDeleteZoneApiGatewayOperationOptions(operationID)
				_, deleteErr := service.DeleteZoneApiGatewayOperation(deleteOptions)
				Expect(deleteErr).To(BeNil())
				operationID = ""
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
					addSelector := &ApiGatewayOperationsLabelsBulkInputSelector{
						Include: &ApiGatewayOperationsLabelsBulkInputSelectorInclude{
							OperationIds: bulkOperationIDs,
						},
					}
					addLabelsOptions := service.NewUpdateApiGatewayOperationLabelsOptions().
						SetManaged(&ApiGatewayOperationsLabelsBulkInputManaged{
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
						found := false
						for _, lbl := range item.Labels {
							if lbl.Name != nil && *lbl.Name == managedLabel {
								found = true
								break
							}
						}
						Expect(found).To(BeTrue(), "Expected label '%s' on operation %s", managedLabel, *item.OperationID)
					}
	
					// --- Step 3: Remove the first operation ID from the label set ---
					remainingIDs := bulkOperationIDs[1:]
					removeSelector := &ApiGatewayOperationsLabelsBulkInputSelector{
						Include: &ApiGatewayOperationsLabelsBulkInputSelectorInclude{
							OperationIds: remainingIDs,
						},
					}
					removeLabelsOptions := service.NewUpdateApiGatewayOperationLabelsOptions().
						SetManaged(&ApiGatewayOperationsLabelsBulkInputManaged{
							Labels: []string{managedLabel},
						}).
						SetSelector(removeSelector)
	
					removeResult, removeResponse, removeErr := service.UpdateApiGatewayOperationLabels(removeLabelsOptions)
					Expect(removeErr).To(BeNil())
					Expect(removeResponse).ToNot(BeNil())
					Expect(removeResult).ToNot(BeNil())
					Expect(*removeResult.Success).To(BeTrue())
	
					for _, item := range removeResult.Result {
						found := false
						for _, lbl := range item.Labels {
							if lbl.Name != nil && *lbl.Name == managedLabel {
								found = true
								break
							}
						}
						Expect(found).To(BeTrue(), "Expected label '%s' on operation %s", managedLabel, *item.OperationID)
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
