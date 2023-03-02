/**
 * (C) Copyright IBM Corp. 2023.
 */

package botmanagementv1_test

import (
	"fmt"
	"os"

	"github.com/IBM/go-sdk-core/v5/core"
	. "github.com/IBM/networking-go-sdk/botmanagementv1"
	"github.com/joho/godotenv"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const configFile = "../cis.env"

var configLoaded bool = true

var _ = Describe(`BotManagementV1`, func() {
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
	zoneId := os.Getenv("ZONE_ID")
	globalOptions := &BotManagementV1Options{
		ServiceName:    DefaultServiceName,
		URL:            serviceURL,
		Authenticator:  authenticator,
		Crn:            &crn,
		ZoneIdentifier: &zoneId,
	}

	service, serviceErr := NewBotManagementV1(globalOptions)
	if serviceErr != nil {
		fmt.Println(serviceErr)
	}
	Describe(`botmanagementv1_test`, func() {
		Context(`botmanagementv1_test`, func() {
			It(`Get | Update Bot Management Settings`, func() {

				// Get Bot Management Settings
				getOptions := service.NewGetBotManagementOptions()
				getResult, getResp, getErr := service.GetBotManagement(getOptions)
				Expect(getErr).To(BeNil())
				Expect(getResp).ToNot(BeNil())
				Expect(getResult).ToNot(BeNil())

				// Set Bot Management Settings
				updateOption := service.NewUpdateBotManagementOptions()
				updateOption.SetFightMode(false)
				updateOption.SetSessionScore(false)
				updateOption.SetEnableJs(true)
				updateOption.SetAuthIdLogging(false)
				updateOption.SetUseLatestModel(false)

				updateResult, updateResp, updateErr := service.UpdateBotManagement(updateOption)
				Expect(updateErr).To(BeNil())
				Expect(updateResp).ToNot(BeNil())
				Expect(updateResult).ToNot(BeNil())

			})
		})
	})
})
