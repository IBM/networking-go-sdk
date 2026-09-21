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

				updateOptions := service.NewReplaceZoneAiSecuritySettingsOptions(newEnabled)
				result, response, err := service.ReplaceZoneAiSecuritySettings(updateOptions)
				Expect(err).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
				Expect(*result.Success).To(BeTrue())

				// Restore original value
				restoreOptions := service.NewReplaceZoneAiSecuritySettingsOptions(currentEnabled)
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
		Context(`Create and Delete API Gateway Operation`, func() {
			It(`Create a single API Gateway Operation and delete it`, func() {
				shouldSkipTest()

				method := "POST"
				host := "api.example.com"
				endpoint := "/v1/messages"

				createOptions := service.NewCreateApiGatewayOperationItemOptions(
					method,
					host,
					endpoint,
				)
				createResult, createResponse, createErr := service.CreateApiGatewayOperationItem(createOptions)
				Expect(createErr).To(BeNil())
				Expect(createResponse).ToNot(BeNil())
				Expect(createResult).ToNot(BeNil())
				Expect(*createResult.Success).To(BeTrue())

				operationID := *createResult.Result.OperationID

				// Delete the created operation
				deleteOptions := service.NewDeleteZoneApiGatewayOperationOptions(operationID)
				deleteResponse, deleteErr := service.DeleteZoneApiGatewayOperation(deleteOptions)
				Expect(deleteErr).To(BeNil())
				Expect(deleteResponse).ToNot(BeNil())
			})
		})
	})
})
