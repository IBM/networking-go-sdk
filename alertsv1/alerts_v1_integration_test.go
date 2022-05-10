/*
 * (C) Copyright IBM Corp. 2022.
 */
package alertsv1_test

import (
	"fmt"
	"os"

	"github.com/IBM/go-sdk-core/v5/core"
	. "github.com/IBM/networking-go-sdk/alertsv1"
	"github.com/IBM/networking-go-sdk/webhooksv1"
	"github.com/joho/godotenv"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const configFile = "../cis.env"

var configLoaded bool = true

func shouldSkipTest() {
	if !configLoaded {
		Skip("External configuration is not available, skipping...")
	}
}

var _ = Describe(`alertsv1`, func() {
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
	serviceURL := os.Getenv("API_ENDPOINT")
	crn := os.Getenv("CRN")
	xAuthUserToken := os.Getenv("CIS_SERVICES_APIKEY")
	globalOptions := &AlertsV1Options{
		ServiceName:   DefaultServiceName,
		URL:           serviceURL,
		Authenticator: authenticator,
		Crn:           &crn,
	}

	webhookOptions := &webhooksv1.WebhooksV1Options{
		ServiceName:   DefaultServiceName,
		URL:           serviceURL,
		Authenticator: authenticator,
		Crn:           &crn,
	}
	testservice, testserviceErr := webhooksv1.NewWebhooksV1(webhookOptions)
	if testserviceErr != nil {
		fmt.Println(testserviceErr)
	}
	service, serviceErr := NewAlertsV1(globalOptions)
	if serviceErr != nil {
		fmt.Println(serviceErr)
	}
	Describe(`alertsv1_test`, func() {
		Context(`alertsv1_test`, func() {
			BeforeEach(func() {
				shouldSkipTest()
				listOpt := service.NewGetAlertPoliciesOptions()
				listResult, listResp, listErr := service.GetAlertPolicies(listOpt)
				print("flag 1")
				print(listResp)
				Expect(listErr).To(BeNil())
				Expect(listResp).ToNot(BeNil())
				Expect(listResult).ToNot(BeNil())
				Expect(*listResult.Success).Should(BeTrue())
				// Delete an alert webhook
				for _, alerts := range listResult.Result {
					delOpt := service.NewDeleteAlertPolicyOptions(*alerts.ID)
					delResult, delResp, delErr := service.DeleteAlertPolicy(delOpt)
					Expect(delErr).To(BeNil())
					Expect(delResp).ToNot(BeNil())
					Expect(delResult).ToNot(BeNil())
					Expect(*delResult.Success).Should(BeTrue())
				}
			})
			AfterEach(func() {
				shouldSkipTest()
				listOpt := service.NewGetAlertPoliciesOptions()
				listResult, listResp, listErr := service.GetAlertPolicies(listOpt)
				Expect(listErr).To(BeNil())
				Expect(listResp).ToNot(BeNil())
				Expect(listResult).ToNot(BeNil())
				Expect(*listResult.Success).Should(BeTrue())
				// Delete an alert webhook
				for _, alerts := range listResult.Result {
					delOpt := service.NewDeleteAlertPolicyOptions(*alerts.ID)
					delResult, delResp, delErr := service.DeleteAlertPolicy(delOpt)
					Expect(delErr).To(BeNil())
					Expect(delResp).ToNot(BeNil())
					Expect(delResult).ToNot(BeNil())
					Expect(*delResult.Success).Should(BeTrue())
				}
			})
			It(`Create| Get | Update | List | Delete AlertsPolicies`, func() {
				// Create a new webhook
				createName := "My Slack Alert Webhook"
				createURL := "https://app.slack.com/client/T02J3DPUE/D02EHU8UPPH"
				options := testservice.NewCreateAlertWebhookOptions()
				options.SetName(createName)
				options.SetURL(createURL)
				options.SetSecret(xAuthUserToken)
				webhookResult, webhookResponse, webhookOperationErr := testservice.CreateAlertWebhook(options)
				webhook_id := *webhookResult.Result.ID
				Expect(webhookOperationErr).To(BeNil())
				Expect(webhookResponse).ToNot(BeNil())
				Expect(webhookResult).ToNot(BeNil())

				//create Alert Policy
				createAlertName := ("My Alert Policy")
				description := ("A description for my alert policy")

				email := &CreateAlertPolicyInputMechanismsEmailItem{
					ID: core.StringPtr("mynotifications@email.com"),
				}

				webhookId := &CreateAlertPolicyInputMechanismsWebhooksItem{
					ID: core.StringPtr(webhook_id),
				}

				mechanism := &CreateAlertPolicyInputMechanisms{
					Email:    []CreateAlertPolicyInputMechanismsEmailItem{*email},
					Webhooks: []CreateAlertPolicyInputMechanismsWebhooksItem{*webhookId},
				}

				filters := map[string]interface{}{
					"enabled": []interface{}{
						"true",
						"false",
					},
					"pool_id": []interface{}{
						"6e67c08e3bae7eb398101d08def8a68a",
						"df2d9d70fcb194ea60d2e58397cb35a6",
					},
				}

				alertOptions := service.NewCreateAlertPolicyOptions()
				alertOptions.SetName(createAlertName)
				alertOptions.SetDescription(description)
				alertOptions.SetEnabled(true)
				alertOptions.SetAlertType(CreateAlertPolicyOptions_AlertType_G6PoolToggleAlert)
				alertOptions.SetMechanisms(mechanism)
				alertOptions.SetFilters(filters)

				result, response, operationErr := service.CreateAlertPolicy(alertOptions)
				alert_id := *result.Result.ID
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				//List Alert Policy
				listOpt := service.NewGetAlertPoliciesOptions()
				listResult, listResp, listErr := service.GetAlertPolicies(listOpt)
				Expect(listErr).To(BeNil())
				Expect(listResp).ToNot(BeNil())
				Expect(listResult).ToNot(BeNil())
				Expect(*listResult.Success).Should(BeTrue())

				//Update Alert Policy
				updateAlertName := ("My Alert Policy")
				upadteDescription := ("A description for my alert policy")

				updateEmail := &UpdateAlertPolicyInputMechanismsEmailItem{
					ID: core.StringPtr("mynotifications2@email.com"),
				}

				updateWebhookId := &UpdateAlertPolicyInputMechanismsWebhooksItem{
					ID: core.StringPtr(webhook_id),
				}

				UpdateMechanism := &UpdateAlertPolicyInputMechanisms{
					Email:    []UpdateAlertPolicyInputMechanismsEmailItem{*updateEmail},
					Webhooks: []UpdateAlertPolicyInputMechanismsWebhooksItem{*updateWebhookId},
				}
				upadteFilters := map[string]interface{}{
					"enabled": []interface{}{
						"true",
						"false",
					},
					"pool_id": []interface{}{
						"6e67c08e3bae7eb398101d08def8a68a",
						"df2d9d70fcb194ea60d2e58397cb35a6",
					},
				}
				updateConditions := map[string]interface{}{
					"and": []interface{}{
						map[string]interface{}{
							"or": []interface{}{
								map[string]interface{}{
									"==": []interface{}{
										map[string]interface{}{
											"var": "pool_id",
										},
										"6e67c08e3bae7eb398101d08def8a68a",
									},
								},
								map[string]interface{}{
									"==": []interface{}{
										map[string]interface{}{
											"var": "pool_id",
										},
										"df2d9d70fcb194ea60d2e58397cb35a6",
									},
								},
							},
						},
						map[string]interface{}{
							"or": []interface{}{
								map[string]interface{}{
									"==": []interface{}{
										map[string]interface{}{
											"var": "enabled",
										},
										"false",
									},
								},
								map[string]interface{}{
									"==": []interface{}{
										map[string]interface{}{
											"var": "enabled",
										},
										"true",
									},
								},
							},
						},
					},
				}
				updateOptions := service.NewUpdateAlertPolicyOptions(alert_id)
				updateOptions.SetPolicyID(alert_id)
				updateOptions.SetName(updateAlertName)
				updateOptions.SetDescription(upadteDescription)
				updateOptions.SetEnabled(true)
				updateOptions.SetAlertType(CreateAlertPolicyOptions_AlertType_G6PoolToggleAlert)
				updateOptions.SetMechanisms(UpdateMechanism)
				updateOptions.SetConditions(updateConditions)
				updateOptions.SetFilters(upadteFilters)

				updateResult, updateResponse, upadteOperationErr := service.UpdateAlertPolicy(updateOptions)
				upadteAlert_id := *updateResult.Result.ID
				Expect(upadteOperationErr).To(BeNil())
				Expect(updateResponse).ToNot(BeNil())
				Expect(updateResult).ToNot(BeNil())

				//Get alert policy by identifiers
				getOpt := service.NewGetAlertPolicyOptions(upadteAlert_id)
				getOpt.SetPolicyID(upadteAlert_id)
				getResult, getResp, getErr := service.GetAlertPolicy(getOpt)
				Expect(getErr).To(BeNil())
				Expect(getResp).ToNot(BeNil())
				Expect(getResult).ToNot(BeNil())
				Expect(*getResult.Success).Should(BeTrue())

				//Delete Alert Policy
				delOpt := service.NewDeleteAlertPolicyOptions(upadteAlert_id)
				delOpt.SetPolicyID(upadteAlert_id)
				delResult, delResp, delErr := service.DeleteAlertPolicy(delOpt)
				Expect(delErr).To(BeNil())
				Expect(delResp).ToNot(BeNil())
				Expect(delResult).ToNot(BeNil())
				Expect(*delResult.Success).Should(BeTrue())

				//Delete Alert webhooks
				webListOpt := testservice.NewListWebhooksOptions()
				webListResult, webListResp, webListErr := testservice.ListWebhooks(webListOpt)
				Expect(webListErr).To(BeNil())
				Expect(webListResp).ToNot(BeNil())
				Expect(webListResult).ToNot(BeNil())
				Expect(*webListResult.Success).Should(BeTrue())
				for _, webhook := range webListResult.Result {
					webDelOpt := testservice.NewDeleteWebhookOptions(*webhook.ID)
					webDelResult, webDelResp, webDelErr := testservice.DeleteWebhook(webDelOpt)
					Expect(webDelErr).To(BeNil())
					Expect(webDelResp).ToNot(BeNil())
					Expect(webDelResult).ToNot(BeNil())
					Expect(*webDelResult.Success).Should(BeTrue())
				}

			})
		})
	})
})
