/*
 * (C) Copyright IBM Corp. 2020.
 */

package routingv1_test

import (
	"fmt"
	"os"

	"github.com/IBM/go-sdk-core/v5/core"
	. "github.com/IBM/networking-go-sdk/routingv1"
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

var _ = Describe(`cisipapiv1_test`, func() {
	BeforeEach(func() {
		Skip("Skipping Tests")
	})

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
	zone_id := os.Getenv("ZONE_ID")
	globalOptions := &RoutingV1Options{
		ServiceName:    "cis_services",
		URL:            serviceURL,
		Authenticator:  authenticator,
		Crn:            &crn,
		ZoneIdentifier: &zone_id,
	}

	service, serviceErr := NewRoutingV1(globalOptions)
	if serviceErr != nil {
		fmt.Println(serviceErr)
	}
	Describe(`routingv1_test`, func() {
		Context(`routingv1_test`, func() {
			It(`smart route setting`, func() {
				shouldSkipTest()
				updateOpt := service.NewUpdateSmartRoutingOptions()
				updateOpt.SetValue(UpdateSmartRoutingOptions_Value_On)
				updateResult, updateResp, updateErr := service.UpdateSmartRouting(updateOpt)
				Expect(updateErr).To(BeNil())
				Expect(updateResp).ToNot(BeNil())
				Expect(updateResult).ToNot(BeNil())
				Expect(*updateResult.Success).Should(BeTrue())

				getOpt := service.NewGetSmartRoutingOptions()
				getResult, getResp, getErr := service.GetSmartRouting(getOpt)
				Expect(getErr).To(BeNil())
				Expect(getResp).ToNot(BeNil())
				Expect(getResult).ToNot(BeNil())
				Expect(*getResult.Success).Should(BeTrue())

				updateOpt.SetValue(UpdateSmartRoutingOptions_Value_Off)
				updateResult, updateResp, updateErr = service.UpdateSmartRouting(updateOpt)
				Expect(updateErr).To(BeNil())
				Expect(updateResp).ToNot(BeNil())
				Expect(updateResult).ToNot(BeNil())
				Expect(*updateResult.Success).Should(BeTrue())

			})
		})
	})
})
