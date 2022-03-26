package webhooksv1_test

import (
	"fmt"
	"os"

	"github.com/IBM/go-sdk-core/v5/core"
	. "github.com/IBM/networking-go-sdk/webhooksv1"
	"github.com/joho/godotenv"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const configFile = "../cis.env"

var configLoaded bool = true

func shouldSkipTest() {
	Skip("Authentication failing, skipping...")
	if !configLoaded {
		Skip("External configuration is not available, skipping...")
	}
}

var _ = Describe(`webhooksv1`, func() {
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
	globalOptions := &WebhooksV1Options{
		ServiceName:   DefaultServiceName,
		URL:           serviceURL,
		Authenticator: authenticator,
		Crn:           &crn,
	}

	service, serviceErr := NewWebhooksV1(globalOptions)
	if serviceErr != nil {
		fmt.Println(serviceErr)
	}
	Describe(`webhooksv1_test`, func() {
		Context(`webhooksv1_test`, func() {
			BeforeEach(func() {
				shouldSkipTest()
				listOpt := service.NewListWebhooksOptions()
				listResult, listResp, listErr := service.ListWebhooks(listOpt)
				Expect(listErr).To(BeNil())
				Expect(listResp).ToNot(BeNil())
				Expect(listResult).ToNot(BeNil())
				Expect(*listResult.Success).Should(BeTrue())
				// Delete an alert webhook
				for _, webhook := range listResult.Result {
					delOpt := service.NewDeleteWebhookOptions(*webhook.ID)
					delResult, delResp, delErr := service.DeleteWebhook(delOpt)
					Expect(delErr).To(BeNil())
					Expect(delResp).ToNot(BeNil())
					Expect(delResult).ToNot(BeNil())
					Expect(*delResult.Success).Should(BeTrue())
				}
			})
			AfterEach(func() {
				shouldSkipTest()
				listOpt := service.NewListWebhooksOptions()
				listResult, listResp, listErr := service.ListWebhooks(listOpt)
				Expect(listErr).To(BeNil())
				Expect(listResp).ToNot(BeNil())
				Expect(listResult).ToNot(BeNil())
				Expect(*listResult.Success).Should(BeTrue())
				// Delete an alert webhook
				for _, webhook := range listResult.Result {
					delOpt := service.NewDeleteWebhookOptions(*webhook.ID)
					delResult, delResp, delErr := service.DeleteWebhook(delOpt)
					Expect(delErr).To(BeNil())
					Expect(delResp).ToNot(BeNil())
					Expect(delResult).ToNot(BeNil())
					Expect(*delResult.Success).Should(BeTrue())
				}
			})
			It(`Create| Get | Update | List | Delete Webhooks`, func() {
				// Create a new webhook
				createName := "My Slack Alert Webhook"
				createURL := "https://app.slack.com/client/T02J3DPUE/D02EHU8UPPH"
				options := service.NewCreateAlertWebhookOptions()
				options.SetName(createName)
				options.SetURL(createURL)
				options.SetSecret(xAuthUserToken)
				result, response, operationErr := service.CreateAlertWebhook(options)
				ID := *result.Result.ID
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Get Webhook by ID
				getOptions := service.NewGetWebhookOptions(ID)
				getResult, getResp, getErr := service.GetWebhook(getOptions)
				Expect(getErr).To(BeNil())
				Expect(getResp).ToNot(BeNil())
				Expect(getResult).ToNot(BeNil())

				// Update Webhook by IDs
				updateName := "My new Alert"
				updateURL := "https://hooks.slack.com/services/Ds3fdBFbV/456464Gdd"
				updateOption := service.NewUpdateAlertWebhookOptions(ID)
				updateOption.SetName(updateName)
				updateOption.SetURL(updateURL)
				updateOption.SetSecret(xAuthUserToken)
				updateResult, updateResp, updateErr := service.UpdateAlertWebhook(updateOption)
				Expect(updateErr).To(BeNil())
				Expect(updateResp).ToNot(BeNil())
				Expect(updateResult).ToNot(BeNil())
				Expect(*updateResult.Success).Should(BeTrue())

				// List Alert-Webhooks
				listOption := service.NewListWebhooksOptions()
				listResult, listResp, listErr := service.ListWebhooks(listOption)
				Expect(listErr).To(BeNil())
				Expect(listResp).ToNot(BeNil())
				Expect(listResult).ToNot(BeNil())
				Expect(*listResult.Success).Should(BeTrue())

				// Delete Webhooks by ID
				delOpt := service.NewDeleteWebhookOptions(ID)
				delResult, delResp, delErr := service.DeleteWebhook(delOpt)
				Expect(delErr).To(BeNil())
				Expect(delResp).ToNot(BeNil())
				Expect(delResult).ToNot(BeNil())
				Expect(*delResult.Success).Should(BeTrue())

			})
		})
	})
})
