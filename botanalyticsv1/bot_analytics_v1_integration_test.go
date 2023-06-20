/**
 * (C) Copyright IBM Corp. 2023.
 */

package botanalyticsv1_test

import (
	"fmt"
	"net/http"
	"os"

	"github.com/IBM/go-sdk-core/v5/core"
	. "github.com/IBM/networking-go-sdk/botanalyticsv1"
	"github.com/joho/godotenv"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const configFile = "../cis.env"

var configLoaded bool = true
var authenticationSucceeded bool = true

var _ = Describe(`BotAnalyticsV1`, func() {
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
	zoneId := os.Getenv("ZONE_ID")
	globalOptions := &BotAnalyticsV1Options{
		ServiceName:    DefaultServiceName,
		URL:            serviceURL,
		Authenticator:  authenticator,
		Crn:            &crn,
		ZoneIdentifier: &zoneId,
	}

	service, serviceErr := NewBotAnalyticsV1(globalOptions)
	if serviceErr != nil {
		fmt.Println(serviceErr)
	}
	Describe(`botanalyticsv1_test`, func() {
		Context(`botanalyticsv1_test`, func() {
			It(`Get Bot Analytics Settings`, func() {
				Skip("Skipping as CF APIs are failing")

				since := CreateMockDateTime("2023-06-13T00:00:00Z")
				until := CreateMockDateTime("2023-06-14T00:00:00Z")

				// Get Bot Analytics Score Source
				getBotScoreOptionsModel := new(GetBotScoreOptions)
				getBotScoreOptionsModel.Since = since
				getBotScoreOptionsModel.Until = until

				getResult, getResp, getErr := service.GetBotScore(getBotScoreOptionsModel)
				Expect(getErr).To(BeNil())
				Expect(getResp).ToNot(BeNil())
				Expect(getResult).ToNot(BeNil())

				// Get Bot Analytics Time Series
				getBotTimeseriesOptionsModel := new(GetBotTimeseriesOptions)
				getBotTimeseriesOptionsModel.Since = since
				getBotTimeseriesOptionsModel.Until = until

				Result, Resp, Err := service.GetBotTimeseries(getBotTimeseriesOptionsModel)
				Expect(Err).To(BeNil())
				Expect(Resp).ToNot(BeNil())
				Expect(Result).ToNot(BeNil())

				// Get Bot Analytics Top Attributes
				getBotTopnsOptionsModel := new(GetBotTopnsOptions)
				getBotTopnsOptionsModel.Since = since
				getBotTopnsOptionsModel.Until = until

				result, resp, err := service.GetBotTopns(getBotTopnsOptionsModel)
				Expect(err).To(BeNil())
				Expect(resp).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
		})
	})
})
