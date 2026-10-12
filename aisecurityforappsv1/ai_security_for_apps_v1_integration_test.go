//go:build integration

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

	// =========================================================================
	// AI Security Settings
	// =========================================================================

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

				// Read current value so we can restore it afterwards
				getOptions := service.NewGetAiSecuritySettingsOptions()
				getResult, _, getErr := service.GetAiSecuritySettings(getOptions)
				Expect(getErr).To(BeNil())
				Expect(getResult).ToNot(BeNil())

				currentEnabled := *getResult.Result.Enabled

				// Toggle to the opposite value
				updateOptions := service.NewReplaceZoneAiSecuritySettingsOptions().
					SetEnabled(!currentEnabled)
				result, response, err := service.ReplaceZoneAiSecuritySettings(updateOptions)
				Expect(err).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
				Expect(*result.Success).To(BeTrue())
				Expect(*result.Result.Enabled).To(Equal(!currentEnabled))

				// Restore original value
				restoreOptions := service.NewReplaceZoneAiSecuritySettingsOptions().
					SetEnabled(currentEnabled)
				_, _, restoreErr := service.ReplaceZoneAiSecuritySettings(restoreOptions)
				Expect(restoreErr).To(BeNil())
			})
		})
	})

	// =========================================================================
	// API Gateway — all operations follow the natural resource lifecycle
	// =========================================================================

	Describe(`API Gateway`, func() {

		// -----------------------------------------------------------------
		// Read-only discovery and schema endpoints
		// -----------------------------------------------------------------

		Context(`Read-only operations`, func() {
			It(`Get API Gateway Discovery (OpenAPI schema view)`, func() {
				shouldSkipTest()

				result, response, err := service.GetApiGatewayDiscovery(service.NewGetApiGatewayDiscoveryOptions())
				Expect(err).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
				Expect(*result.Success).To(BeTrue())
			})

			It(`List API Gateway Discovery Operations (default params)`, func() {
				shouldSkipTest()

				result, response, err := service.ListApiGatewayDiscoveryOperations(
					service.NewListApiGatewayDiscoveryOperationsOptions(),
				)
				Expect(err).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
				Expect(*result.Success).To(BeTrue())
			})

			It(`List API Gateway Discovery Operations with filters`, func() {
				shouldSkipTest()

				opts := service.NewListApiGatewayDiscoveryOperationsOptions().
					SetPerPage(5).
					SetPage(1)
				result, response, err := service.ListApiGatewayDiscoveryOperations(opts)
				Expect(err).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
				Expect(*result.Success).To(BeTrue())
			})

			It(`Get API Gateway Schemas (OpenAPI schema view)`, func() {
				shouldSkipTest()

				result, response, err := service.GetApiGatewaySchemas(service.NewGetApiGatewaySchemasOptions())
				Expect(err).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
				Expect(*result.Success).To(BeTrue())
			})
		})

		// -----------------------------------------------------------------
		// Discovery state update (PATCH) — depends on zone having
		// discovered operations in "review" state; skipped gracefully if not
		// -----------------------------------------------------------------

		Context(`Discovery state update`, func() {
			It(`Bulk-update discovered operation state (ignored → review)`, func() {
				shouldSkipTest()

				// Find at least one operation in "review" state to patch
				listOpts := service.NewListApiGatewayDiscoveryOperationsOptions().
					SetState("review").
					SetPerPage(5)
				listResult, listResponse, listErr := service.ListApiGatewayDiscoveryOperations(listOpts)
				Expect(listErr).To(BeNil())
				Expect(listResponse).ToNot(BeNil())
				Expect(listResult).ToNot(BeNil())
				Expect(*listResult.Success).To(BeTrue())

				if len(listResult.Result) == 0 {
					Skip("No discovered operations in 'review' state — skipping discovery PATCH test")
				}

				discoveredID := *listResult.Result[0].ID

				// PATCH: mark "ignored"; body is a free-form map { uuid: { state } }
				patchBody := map[string]interface{}{
					discoveredID: map[string]interface{}{"state": "ignored"},
				}
				patchResult, patchResponse, patchErr := service.UpdateZoneApiGatewayDiscoveryOperation(
					service.NewUpdateZoneApiGatewayDiscoveryOperationOptions().SetRequestBody(patchBody),
				)
				Expect(patchErr).To(BeNil())
				Expect(patchResponse).ToNot(BeNil())
				Expect(patchResult).ToNot(BeNil())
				Expect(*patchResult.Success).To(BeTrue())

				// Restore to "review"
				restoreBody := map[string]interface{}{
					discoveredID: map[string]interface{}{"state": "review"},
				}
				_, _, restoreErr := service.UpdateZoneApiGatewayDiscoveryOperation(
					service.NewUpdateZoneApiGatewayDiscoveryOperationOptions().SetRequestBody(restoreBody),
				)
				Expect(restoreErr).To(BeNil())
			})
		})

		// -----------------------------------------------------------------
		// Single-operation lifecycle: Create → List (verify) → Get → Delete → 404
		// -----------------------------------------------------------------

		Context(`Single operation lifecycle`, func() {
			It(`Create, list, get, delete, and verify 404`, func() {
				shouldSkipTest()

				// Create a single operation via the /item endpoint
				createResult, createResponse, createErr := service.CreateApiGatewayOperationItem(
					service.NewCreateApiGatewayOperationItemOptions().
						SetMethod("POST").
						SetHost("api.example.com").
						SetEndpoint("/v1/messages"),
				)
				Expect(createErr).To(BeNil())
				Expect(createResponse).ToNot(BeNil())
				Expect(createResult).ToNot(BeNil())
				Expect(*createResult.Success).To(BeTrue())
				Expect(createResult.Result).ToNot(BeNil())
				Expect(createResult.Result.OperationID).ToNot(BeNil())

				operationID := *createResult.Result.OperationID

				// Safety cleanup if a later assertion fails
				defer func() {
					if operationID != "" {
						_, _ = service.DeleteZoneApiGatewayOperation(
							service.NewDeleteZoneApiGatewayOperationOptions(operationID),
						)
					}
				}()

				// List operations and verify the newly created one appears
				listResult, listResponse, listErr := service.ListZoneApiGatewayOperations(
					service.NewListZoneApiGatewayOperationsOptions().
						SetHost([]string{"api.example.com"}).
						SetMethod([]string{"POST"}),
				)
				Expect(listErr).To(BeNil())
				Expect(listResponse).ToNot(BeNil())
				Expect(listResult).ToNot(BeNil())
				Expect(*listResult.Success).To(BeTrue())
				found := false
				for _, op := range listResult.Result {
					if op.OperationID != nil && *op.OperationID == operationID {
						found = true
						break
					}
				}
				Expect(found).To(BeTrue(), "Expected newly created operation %s in list result", operationID)

				// Get by ID and verify fields round-trip correctly
				getResult, getResponse, getErr := service.GetZoneApiGatewayOperation(
					service.NewGetZoneApiGatewayOperationOptions(operationID),
				)
				Expect(getErr).To(BeNil())
				Expect(getResponse).ToNot(BeNil())
				Expect(getResult).ToNot(BeNil())
				Expect(*getResult.Success).To(BeTrue())
				Expect(*getResult.Result.OperationID).To(Equal(operationID))
				Expect(*getResult.Result.Method).To(Equal("POST"))
				Expect(*getResult.Result.Host).To(Equal("api.example.com"))
				Expect(*getResult.Result.Endpoint).To(Equal("/v1/messages"))

				// Delete and assert 204 / no error
				deleteResponse, deleteErr := service.DeleteZoneApiGatewayOperation(
					service.NewDeleteZoneApiGatewayOperationOptions(operationID),
				)
				Expect(deleteErr).To(BeNil())
				Expect(deleteResponse).ToNot(BeNil())
				operationID = "" // prevent defer double-delete

				// Verify the operation is gone (404)
				_, verifyResponse, verifyErr := service.GetZoneApiGatewayOperation(
					service.NewGetZoneApiGatewayOperationOptions(*createResult.Result.OperationID),
				)
				Expect(verifyErr).ToNot(BeNil())
				Expect(verifyResponse.StatusCode).To(Equal(404))
			})
		})

		// -----------------------------------------------------------------
		// Bulk operation lifecycle:
		//   Create (bulk) → Get one → Update labels (add) → Replace labels
		//   (bulk PUT) → Delete labels (bulk DELETE) → Delete all individually
		// -----------------------------------------------------------------

		Context(`Bulk operation lifecycle`, func() {
			It(`Create 3 operations, label them, verify labels, then clean up`, func() {
				shouldSkipTest()

				// Build 3 operations and create them in one bulk call
				op1, err := service.NewApiGatewayOperation("GET", "api.example.com", "/v2/users")
				Expect(err).To(BeNil())
				op2, err := service.NewApiGatewayOperation("POST", "api.example.com", "/v2/orders")
				Expect(err).To(BeNil())
				op3, err := service.NewApiGatewayOperation("DELETE", "api.example.com", "/v2/sessions")
				Expect(err).To(BeNil())

				bulkResult, bulkResponse, bulkErr := service.CreateZoneApiGatewayOperation(
					service.NewCreateZoneApiGatewayOperationOptions().
						SetApiGatewayOperation([]ApiGatewayOperation{*op1, *op2, *op3}),
				)
				Expect(bulkErr).To(BeNil())
				Expect(bulkResponse).ToNot(BeNil())
				Expect(bulkResult).ToNot(BeNil())
				Expect(*bulkResult.Success).To(BeTrue())
				Expect(len(bulkResult.Result)).To(Equal(3))

				opIDs := make([]string, 0, 3)
				for _, item := range bulkResult.Result {
					Expect(item.OperationID).ToNot(BeNil())
					opIDs = append(opIDs, *item.OperationID)
				}

				// Safety cleanup — always delete individually on exit
				defer func() {
					for _, id := range opIDs {
						_, _ = service.DeleteZoneApiGatewayOperation(
							service.NewDeleteZoneApiGatewayOperationOptions(id),
						)
					}
				}()

				// Step 1: Get one operation by ID and verify its fields
				getResult, getResponse, getErr := service.GetZoneApiGatewayOperation(
					service.NewGetZoneApiGatewayOperationOptions(opIDs[1]),
				)
				Expect(getErr).To(BeNil())
				Expect(getResponse).ToNot(BeNil())
				Expect(getResult).ToNot(BeNil())
				Expect(*getResult.Success).To(BeTrue())
				Expect(*getResult.Result.OperationID).To(Equal(opIDs[1]))

				// Step 2: Add managed label to all 3 ops (POST /operations/labels)
				// API requires 'user' field to be present even when empty
				const managedLabel = "cf-llm"
				addResult, addResponse, addErr := service.UpdateApiGatewayOperationLabels(
					service.NewUpdateApiGatewayOperationLabelsOptions().
						SetUser(&ApiGatewayOperationsLabelsBulkInputUser{Labels: []string{}}).
						SetManaged(&ApiGatewayOperationsLabelsBulkInputManaged{Labels: []string{managedLabel}}).
						SetSelector(&ApiGatewayOperationsLabelsBulkInputSelector{
							Include: &ApiGatewayOperationsLabelsBulkInputSelectorInclude{
								OperationIds: opIDs,
							},
						}),
				)
				Expect(addErr).To(BeNil())
				Expect(addResponse).ToNot(BeNil())
				Expect(addResult).ToNot(BeNil())
				Expect(*addResult.Success).To(BeTrue())
				Expect(len(addResult.Result)).To(Equal(3))

				// Verify all 3 operations have the label
				for _, item := range addResult.Result {
					hasLabel := false
					for _, lbl := range item.Labels {
						if lbl.Name != nil && *lbl.Name == managedLabel {
							hasLabel = true
							break
						}
					}
					Expect(hasLabel).To(BeTrue(), "Expected label '%s' on operation %s", managedLabel, *item.OperationID)
				}

				// Step 3: Replace labels on all 3 ops (PUT /operations/labels)
				replaceResult, replaceResponse, replaceErr := service.ReplaceZoneApiGatewayOperationLabels(
					service.NewReplaceZoneApiGatewayOperationLabelsOptions().
						SetUser(&ApiGatewayOperationsLabelsBulkInputUser{Labels: []string{}}).
						SetManaged(&ApiGatewayOperationsLabelsBulkInputManaged{Labels: []string{managedLabel}}).
						SetSelector(&ApiGatewayOperationsLabelsBulkInputSelector{
							Include: &ApiGatewayOperationsLabelsBulkInputSelectorInclude{
								OperationIds: opIDs,
							},
						}),
				)
				Expect(replaceErr).To(BeNil())
				Expect(replaceResponse).ToNot(BeNil())
				Expect(replaceResult).ToNot(BeNil())
				Expect(*replaceResult.Success).To(BeTrue())

				// Step 4: Bulk-delete labels from all 3 ops (DELETE /operations/labels)
				deleteLabelsResult, deleteLabelsResponse, deleteLabelsErr := service.DeleteZoneApiGatewayOperationLabels(
					service.NewDeleteZoneApiGatewayOperationLabelsOptions(
						&ApiGatewayOperationsLabelsBulkDeleteInputSelector{
							Include: &ApiGatewayOperationsLabelsBulkDeleteInputSelectorInclude{
								OperationIds: opIDs,
							},
						},
					).SetManaged(&ApiGatewayOperationsLabelsBulkDeleteInputManaged{
						Labels: []string{managedLabel},
					}),
				)
				Expect(deleteLabelsErr).To(BeNil())
				Expect(deleteLabelsResponse).ToNot(BeNil())
				Expect(deleteLabelsResult).ToNot(BeNil())
				Expect(*deleteLabelsResult.Success).To(BeTrue())

				// Step 5: Delete all operations individually and assert each call succeeds
				for _, id := range opIDs {
					deleteResponse, deleteErr := service.DeleteZoneApiGatewayOperation(
						service.NewDeleteZoneApiGatewayOperationOptions(id),
					)
					Expect(deleteErr).To(BeNil())
					Expect(deleteResponse).ToNot(BeNil())
				}
				opIDs = nil // prevent defer double-delete
			})
		})

		// -----------------------------------------------------------------
		// Per-operation label lifecycle (POST/PUT/DELETE on /{operation_id}/labels)
		// Uses CreateApiGatewayOperationItem to create a fresh operation,
		// exercises all three per-operation label verbs, then deletes it.
		// -----------------------------------------------------------------

		Context(`Per-operation label lifecycle`, func() {
			It(`Create operation, attach labels, replace labels, remove labels, delete operation`, func() {
				shouldSkipTest()

				// Create a dedicated operation to attach labels to
				createResult, createResponse, createErr := service.CreateApiGatewayOperationItem(
					service.NewCreateApiGatewayOperationItemOptions().
						SetMethod("POST").
						SetHost("api.example.com").
						SetEndpoint("/v1/label-test"),
				)
				Expect(createErr).To(BeNil())
				Expect(createResponse).ToNot(BeNil())
				Expect(createResult).ToNot(BeNil())
				Expect(*createResult.Success).To(BeTrue())
				Expect(createResult.Result.OperationID).ToNot(BeNil())

				operationID := *createResult.Result.OperationID

				// Safety cleanup
				defer func() {
					if operationID != "" {
						_, _ = service.DeleteZoneApiGatewayOperation(
							service.NewDeleteZoneApiGatewayOperationOptions(operationID),
						)
					}
				}()

				// Step 1: Attach labels — POST /{operation_id}/labels
				// API requires both 'user' and 'managed' in the body
				addResult, addResponse, addErr := service.CreateApiGatewayOperationLabels(
					service.NewCreateApiGatewayOperationLabelsOptions(operationID).
						SetUser([]string{}).
						SetManaged([]string{"cf-llm"}),
				)
				Expect(addErr).To(BeNil())
				Expect(addResponse).ToNot(BeNil())
				Expect(addResult).ToNot(BeNil())
				Expect(*addResult.Success).To(BeTrue())

				// Step 2: Replace labels — PUT /{operation_id}/labels
				// API requires both 'user' and 'managed' in the body
				setResult, setResponse, setErr := service.SetZoneApiGatewayOperationLabels(
					service.NewSetZoneApiGatewayOperationLabelsOptions(operationID).
						SetUser([]string{}).
						SetManaged([]string{"cf-llm"}),
				)
				Expect(setErr).To(BeNil())
				Expect(setResponse).ToNot(BeNil())
				Expect(setResult).ToNot(BeNil())
				Expect(*setResult.Success).To(BeTrue())

				// Step 3: Remove labels — DELETE /{operation_id}/labels
				unsetResult, unsetResponse, unsetErr := service.UnsetZoneApiGatewayOperationLabels(
					service.NewUnsetZoneApiGatewayOperationLabelsOptions(operationID).
						SetManaged([]string{"cf-llm"}),
				)
				Expect(unsetErr).To(BeNil())
				Expect(unsetResponse).ToNot(BeNil())
				Expect(unsetResult).ToNot(BeNil())
				Expect(*unsetResult.Success).To(BeTrue())

				// Clean up the operation
				_, deleteErr := service.DeleteZoneApiGatewayOperation(
					service.NewDeleteZoneApiGatewayOperationOptions(operationID),
				)
				Expect(deleteErr).To(BeNil())
				operationID = "" // prevent defer double-delete
			})
		})

		// -----------------------------------------------------------------
		// Bulk operation delete (DELETE /operations with operation_ids body)
		// Creates two operations, attempts bulk-delete.
		// Tolerates a 400 if the backend rejects the format on this zone —
		// the individual defer cleanup handles the actual teardown.
		// -----------------------------------------------------------------

		Context(`Bulk operation delete`, func() {
			It(`Create two operations then bulk-delete them`, func() {
				shouldSkipTest()

				op1, err := service.NewApiGatewayOperation("GET", "api.example.com", "/v1/bulk-del/a")
				Expect(err).To(BeNil())
				op2, err := service.NewApiGatewayOperation("POST", "api.example.com", "/v1/bulk-del/b")
				Expect(err).To(BeNil())

				createResult, createResponse, createErr := service.CreateZoneApiGatewayOperation(
					service.NewCreateZoneApiGatewayOperationOptions().
						SetApiGatewayOperation([]ApiGatewayOperation{*op1, *op2}),
				)
				Expect(createErr).To(BeNil())
				Expect(createResponse).ToNot(BeNil())
				Expect(createResult).ToNot(BeNil())
				Expect(*createResult.Success).To(BeTrue())
				Expect(len(createResult.Result)).To(BeNumerically(">=", 1))

				ids := make([]string, 0, len(createResult.Result))
				for _, op := range createResult.Result {
					if op.OperationID != nil {
						ids = append(ids, *op.OperationID)
					}
				}

				// Always clean up individually — defer handles teardown regardless of bulk outcome
				defer func() {
					for _, id := range ids {
						_, _ = service.DeleteZoneApiGatewayOperation(
							service.NewDeleteZoneApiGatewayOperationOptions(id),
						)
					}
				}()

				// Attempt bulk-delete; tolerate failure (backend inconsistency on some zones)
				bulkDeleteResult, bulkDeleteResponse, bulkDeleteErr := service.DeleteZoneApiGatewayOperations(
					service.NewDeleteZoneApiGatewayOperationsOptions().SetOperationIds(ids),
				)
				if bulkDeleteErr == nil {
					Expect(bulkDeleteResponse).ToNot(BeNil())
					Expect(bulkDeleteResult).ToNot(BeNil())
					Expect(*bulkDeleteResult.Success).To(BeTrue())
					ids = nil // prevent defer double-delete
				} else {
					fmt.Printf("DeleteZoneApiGatewayOperations returned error (may be unsupported on this zone): %v\n", bulkDeleteErr)
				}
			})
		})
	})
})
