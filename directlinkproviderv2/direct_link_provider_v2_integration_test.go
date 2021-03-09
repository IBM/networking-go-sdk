/**
 * (C) Copyright IBM Corp. 2020.
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

package directlinkproviderv2_test

/*
How to run this test:
go test -v ./directlinkproviderv2
*/

import (
	"encoding/json"
	"os"
	"strconv"
	"time"

	"github.com/IBM/go-sdk-core/v4/core"
	"github.com/IBM/networking-go-sdk/directlinkproviderv2"
	"github.com/IBM/networking-go-sdk/directlinkv1"
	"github.com/joho/godotenv"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var configLoaded = false

func shouldSkipTest() {
	if !configLoaded {
		Skip("External configuration is not available, skipping...")
	}
}

var _ = Describe(`DirectLinkProviderV2`, func() {
	err := godotenv.Load("../directlink.env")
	It(`Successfully loading .env file`, func() {
		if err == nil {
			serviceURLV2 := os.Getenv("SERVICE_URL_V2")
			if serviceURLV2 != "" {
				configLoaded = true
			}
		}
		if !configLoaded {
			Skip("External configuration could not be loaded, skipping...")
		}
	})

	authenticator := &core.IamAuthenticator{
		ApiKey: os.Getenv("IAMAPIKEY"),
		URL:    "https://iam.test.cloud.ibm.com/identity/token",
	}

	version := time.Now().Format("2006-01-02")
	serviceURL := os.Getenv("SERVICE_URL")
	options := &directlinkv1.DirectLinkV1Options{
		ServiceName:   "DirectLinkV1_Mocking",
		Authenticator: authenticator,
		URL:           serviceURL,
		Version:       &version,
	}

	serviceV1, err := directlinkv1.NewDirectLinkV1UsingExternalConfig(options)
	It(`Successfully created DirectLinkV1 service instance`, func() {
		shouldSkipTest()
		Expect(err).To(BeNil())
	})

	// Create instance of DirectLink V2 services
	authenticatorV2 := &core.IamAuthenticator{
		ApiKey: os.Getenv("PROVIDER_IAMAPIKEY"),
		URL:    "https://iam.test.cloud.ibm.com/identity/token",
	}

	serviceURLV2 := os.Getenv("SERVICE_URL_V2")
	optionsV2 := &directlinkproviderv2.DirectLinkProviderV2Options{
		ServiceName:   "DirectLinkProviderV2_Mocking",
		Authenticator: authenticatorV2,
		URL:           serviceURLV2,
		Version:       &version,
	}

	serviceV2, err := directlinkproviderv2.NewDirectLinkProviderV2(optionsV2)
	It(`Successfully created DirectLinkProviderV2 service instance`, func() {
		shouldSkipTest()
		Expect(err).To(BeNil())
	})

	Describe("Direct Link Provider Ports", func() {
		listPortsOptions := serviceV2.NewListProviderPortsOptions()
		var firstPort directlinkproviderv2.ProviderPort

		It(`Successfully list all ports`, func() {
			shouldSkipTest()

			result, detailedResponse, err := serviceV2.ListProviderPorts(listPortsOptions)
			Expect(err).To(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(200))

			ports := result.Ports
			Expect(len(ports)).Should(BeNumerically(">", 0))

			firstPort = ports[0]
			Expect(*firstPort.ID).NotTo(Equal(""))
			Expect(*firstPort.Label).NotTo(Equal(""))
			Expect(*firstPort.LocationName).NotTo(Equal(""))
			Expect(*firstPort.LocationDisplayName).NotTo(Equal(""))
			Expect(*firstPort.ProviderName).NotTo(Equal(""))
			Expect(len(firstPort.SupportedLinkSpeeds)).Should(BeNumerically(">", 0))
		})

		It(`Successfully GET a specific port`, func() {
			shouldSkipTest()

			portsOptions := serviceV2.NewGetProviderPortOptions(*firstPort.ID)
			result, detailedResponse, err := serviceV2.GetProviderPort(portsOptions)
			Expect(err).To(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(200))

			Expect(*result.ID).To(Equal(*firstPort.ID))
			Expect(*result.Label).To(Equal(*firstPort.Label))
			Expect(*result.LocationName).To(Equal(*firstPort.LocationName))
			Expect(*result.LocationDisplayName).To(Equal(*firstPort.LocationDisplayName))
			Expect(*result.ProviderName).To(Equal(*firstPort.ProviderName))
			Expect(result.SupportedLinkSpeeds).To(Equal(firstPort.SupportedLinkSpeeds))
		})
	})

	Describe("Direct Link Provider Gateways", func() {
		timestamp := time.Now().Unix()
		gatewayName := "GO-INT-SDK-PROVIDER-" + strconv.FormatInt(timestamp, 10)
		updatedGatewayName := "GO-INT-SDK-PROVIDER-PATCH-" + strconv.FormatInt(timestamp, 10)
		bgpAsn := int64(64999)
		customerAccount := os.Getenv("CUSTOMER_ACCT_ID")
		speedMbps := int64(1000)

		// Construct an instance of the ProviderGatewayPortIdentity model
		providerGatewayPortIdentityModel := new(directlinkproviderv2.ProviderGatewayPortIdentity)
		var firstPort directlinkproviderv2.ProviderPort

		Context("Successfully test provider gateway CRUD", func() {
			It(`Successfully get a provider port`, func() {
				shouldSkipTest()

				listPortsOptions := serviceV2.NewListProviderPortsOptions()
				result, detailedResponse, err := serviceV2.ListProviderPorts(listPortsOptions)
				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(200))
				ports := result.Ports
				firstPort = ports[0]
				providerGatewayPortIdentityModel.ID = firstPort.ID
			})

			It(`Successfully create gateway`, func() {
				shouldSkipTest()

				gatewayOptions := new(directlinkproviderv2.CreateProviderGatewayOptions)
				gatewayOptions.BgpAsn = core.Int64Ptr(int64(64999))
				gatewayOptions.CustomerAccountID = core.StringPtr(customerAccount)
				gatewayOptions.Name = core.StringPtr(gatewayName)
				gatewayOptions.Port = providerGatewayPortIdentityModel
				gatewayOptions.SpeedMbps = core.Int64Ptr(1000)

				result, detailedResponse, err := serviceV2.CreateProviderGateway(gatewayOptions)

				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(201))

				Expect(*result.ID).NotTo(Equal(""))
				Expect(*result.Name).To(Equal(gatewayName))
				Expect(*result.BgpAsn).To(Equal(bgpAsn))
				Expect(*result.SpeedMbps).To(Equal(speedMbps))
				Expect(*result.BgpCerCidr).NotTo(BeEmpty())
				Expect(*result.BgpIbmCidr).NotTo(Equal(""))
				Expect(*result.BgpIbmAsn).NotTo(Equal(""))
				Expect(*result.BgpStatus).To(Equal("idle"))
				Expect(*result.CreatedAt).NotTo(Equal(""))
				Expect(*result.Crn).To(HavePrefix("crn:v1"))
				Expect(*result.CustomerAccountID).To(Equal(customerAccount))
				Expect(*result.OperationalStatus).To(Equal("create_pending"))
				Expect(*result.Port.ID).To(Equal(*firstPort.ID))
				Expect(*result.ProviderApiManaged).To(Equal(true))
				Expect(*result.Type).To(Equal("connect"))

				os.Setenv("GATEWAY_ID", *result.ID)
			})

			It(`Successfully get the created gateway`, func() {
				shouldSkipTest()

				getProviderGatewayOptions := new(directlinkproviderv2.GetProviderGatewayOptions)
				getProviderGatewayOptions.ID = core.StringPtr(os.Getenv("GATEWAY_ID"))

				result, detailedResponse, err := serviceV2.GetProviderGateway(getProviderGatewayOptions)

				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(200))

				Expect(*result.ID).NotTo(Equal(""))
				Expect(*result.Name).To(Equal(gatewayName))
				Expect(*result.BgpAsn).To(Equal(bgpAsn))
				Expect(*result.SpeedMbps).To(Equal(speedMbps))
				Expect(*result.BgpCerCidr).NotTo(BeEmpty())
				Expect(*result.BgpIbmCidr).NotTo(Equal(""))
				Expect(*result.BgpIbmAsn).NotTo(Equal(""))
				Expect(*result.BgpStatus).To(Equal("idle"))
				Expect(*result.CreatedAt).NotTo(Equal(""))
				Expect(*result.Crn).To(HavePrefix("crn:v1"))
				Expect(*result.CustomerAccountID).To(Equal(customerAccount))
				Expect(*result.OperationalStatus).To(Equal("create_pending"))
				Expect(*result.Port.ID).To(Equal(*firstPort.ID))
				Expect(*result.ProviderApiManaged).To(Equal(true))
				Expect(*result.Type).To(Equal("connect"))
			})

			It(`Successfully list gateways`, func() {
				shouldSkipTest()

				listGatewayOptions := serviceV2.NewListProviderGatewaysOptions()
				result, detailedResponse, err := serviceV2.ListProviderGateways(listGatewayOptions)

				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(200))

				for _, gw := range result.Gateways {
					if *gw.ID == os.Getenv("GATEWAY_ID") {
						Expect(*gw.Name).To(Equal(gatewayName))
						Expect(*gw.BgpAsn).To(Equal(bgpAsn))
						Expect(*gw.SpeedMbps).To(Equal(speedMbps))
						Expect(*gw.BgpCerCidr).NotTo(BeEmpty())
						Expect(*gw.BgpIbmCidr).NotTo(Equal(""))
						Expect(*gw.BgpIbmAsn).NotTo(Equal(""))
						Expect(*gw.BgpStatus).To(Equal("idle"))
						Expect(*gw.CreatedAt).NotTo(Equal(""))
						Expect(*gw.Crn).To(HavePrefix("crn:v1"))
						Expect(*gw.CustomerAccountID).To(Equal(customerAccount))
						Expect(*gw.OperationalStatus).To(Equal("create_pending"))
						Expect(*gw.Port.ID).To(Equal(*firstPort.ID))
						Expect(*gw.ProviderApiManaged).To(Equal(true))
						Expect(*gw.Type).To(Equal("connect"))
					}
				}
			})

			It(`Successfully fail update of the created gateway due to invalid status`, func() {
				shouldSkipTest()

				updateProviderGatewayOptions := new(directlinkproviderv2.UpdateProviderGatewayOptions)
				updateProviderGatewayOptions.ID = core.StringPtr(os.Getenv("GATEWAY_ID"))
				updateProviderGatewayOptions.Name = core.StringPtr(updatedGatewayName)
				updateProviderGatewayOptions.SpeedMbps = core.Int64Ptr(int64(2000))

				_, detailedResponse, err := serviceV2.UpdateProviderGateway(updateProviderGatewayOptions)

				Expect(err).NotTo(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(400))

				// verify we got the correct error message
				detailedResult := detailedResponse.Result
				errInfo, _ := json.Marshal(detailedResult)
				Expect(string(errInfo)).Should(ContainSubstring("Cannot update a gateway with current status"))
			})

			It("Successfully deletes a gateway", func() {
				shouldSkipTest()

				gatewayId := os.Getenv("GATEWAY_ID")
				deteleGatewayOptions := serviceV2.NewDeleteProviderGatewayOptions(gatewayId)

				_, detailedResponse, _ := serviceV2.DeleteProviderGateway(deteleGatewayOptions)

				Expect(detailedResponse.StatusCode).To(Equal(204))
			})
		})
	})

	Describe("Direct Link Provider Gateways with client API", func() {
		timestamp := time.Now().Unix()
		gatewayName := "GO-INT-SDK-PROVIDER-" + strconv.FormatInt(timestamp, 10)
		updatedGatewayName := "GO-INT-SDK-PROVIDER-PATCH-" + strconv.FormatInt(timestamp, 10)
		bgpAsn := int64(64999)
		customerAccount := os.Getenv("CUSTOMER_ACCT_ID")
		speedMbps := int64(1000)
		updatedSpeedMbps := int64(2000)

		// Construct an instance of the ProviderGatewayPortIdentity model
		providerGatewayPortIdentityModel := new(directlinkproviderv2.ProviderGatewayPortIdentity)
		var firstPort directlinkproviderv2.ProviderPort

		Context("Successfully create/approve/patch gateway", func() {
			It(`Successfully get a provider port`, func() {
				shouldSkipTest()

				listPortsOptions := serviceV2.NewListProviderPortsOptions()
				result, detailedResponse, err := serviceV2.ListProviderPorts(listPortsOptions)
				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(200))
				ports := result.Ports
				firstPort = ports[0]
				providerGatewayPortIdentityModel.ID = firstPort.ID
			})

			It(`Successfully create gateway`, func() {
				shouldSkipTest()

				gatewayOptions := new(directlinkproviderv2.CreateProviderGatewayOptions)
				gatewayOptions.BgpAsn = core.Int64Ptr(bgpAsn)
				gatewayOptions.CustomerAccountID = core.StringPtr(customerAccount)
				gatewayOptions.Name = core.StringPtr(gatewayName)
				gatewayOptions.Port = providerGatewayPortIdentityModel
				gatewayOptions.SpeedMbps = core.Int64Ptr(speedMbps)

				result, detailedResponse, err := serviceV2.CreateProviderGateway(gatewayOptions)

				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(201))

				os.Setenv("GATEWAY_ID", *result.ID)
			})

			It(`Successfully verify client account can see the created gateway`, func() {
				shouldSkipTest()

				getGatewayOptions := serviceV1.NewGetGatewayOptions(os.Getenv("GATEWAY_ID"))

				// Get the current status for the gateway
				result, detailedResponse, err := serviceV1.GetGateway(getGatewayOptions)

				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(200))

				Expect(*result.ID).To(Equal(os.Getenv("GATEWAY_ID")))
				Expect(*result.Name).To(Equal(gatewayName))
				Expect(*result.BgpAsn).To(Equal(bgpAsn))
				Expect(*result.SpeedMbps).To(Equal(speedMbps))
				Expect(*result.BgpCerCidr).NotTo(BeEmpty())
				Expect(*result.BgpIbmCidr).NotTo(Equal(""))
				Expect(*result.BgpIbmAsn).NotTo(Equal(""))
				Expect(*result.BgpStatus).To(Equal("idle"))
				Expect(*result.CreatedAt).NotTo(Equal(""))
				Expect(*result.Crn).To(HavePrefix("crn:v1"))
				Expect(*result.LocationDisplayName).To(Equal(*firstPort.LocationDisplayName))
				Expect(*result.LocationName).To(Equal(*firstPort.LocationName))
				Expect(*result.OperationalStatus).To(Equal("create_pending"))
				Expect(*result.Port.ID).To(Equal(*firstPort.ID))
				Expect(result.ChangeRequest).NotTo(BeNil())
				Expect(*result.ProviderApiManaged).To(Equal(true))
				Expect(*result.Type).To(Equal("connect"))

			})

			It(`Successfully approve gateway create using client account`, func() {
				shouldSkipTest()

				createGatewayActionOptions := serviceV1.NewCreateGatewayActionOptions(os.Getenv("GATEWAY_ID"),
					"create_gateway_approve")
				createGatewayActionOptions.SetMetered(false)
				createGatewayActionOptions.SetGlobal(false)

				// Get the current status for the gateway
				result, detailedResponse, err := serviceV1.CreateGatewayAction(createGatewayActionOptions)

				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(200))

				Expect(*result.ID).To(Equal(os.Getenv("GATEWAY_ID")))
				Expect(*result.Name).To(Equal(gatewayName))
				Expect(*result.BgpAsn).To(Equal(bgpAsn))
				Expect(*result.SpeedMbps).To(Equal(speedMbps))
				Expect(*result.BgpCerCidr).NotTo(BeEmpty())
				Expect(*result.BgpIbmCidr).NotTo(Equal(""))
				Expect(*result.BgpIbmAsn).NotTo(Equal(""))
				Expect(*result.BgpStatus).To(Equal("idle"))
				Expect(*result.CreatedAt).NotTo(Equal(""))
				Expect(*result.Crn).To(HavePrefix("crn:v1"))
				Expect(*result.Global).To(Equal(false))
				Expect(*result.Metered).To(Equal(false))
				Expect(*result.OperationalStatus).To(Equal("create_pending"))
				Expect(*result.Port.ID).To(Equal(*firstPort.ID))
				Expect(*result.ProviderApiManaged).To(Equal(true))
				Expect(*result.Type).To(Equal("connect"))
				Expect(*result.ResourceGroup.ID).NotTo(Equal(""))
			})

			It("Successfully waits for connect gateway to move to provisioned state", func() {
				shouldSkipTest()

				getGatewayOptions := serviceV1.NewGetGatewayOptions(os.Getenv("GATEWAY_ID"))

				// before a connect gateway can be deleted, it needs to have operational_status of provisioned.  We need to wait for
				// the new gateway to go to provisioned so we can delete it.
				timer := 0
				for {
					// Get the current status for the gateway
					result, detailedResponse, err := serviceV1.GetGateway(getGatewayOptions)

					Expect(err).To(BeNil())
					Expect(detailedResponse.StatusCode).To(Equal(200))

					// if operational status is "provisioned" then we are done
					if *result.OperationalStatus == "provisioned" {
						Expect(*result.ID).To(Equal(os.Getenv("GATEWAY_ID")))
						Expect(*result.Name).To(Equal(gatewayName))
						Expect(*result.BgpAsn).To(Equal(bgpAsn))
						Expect(*result.SpeedMbps).To(Equal(speedMbps))
						Expect(*result.BgpCerCidr).NotTo(BeEmpty())
						Expect(*result.BgpIbmCidr).NotTo(Equal(""))
						Expect(*result.BgpIbmAsn).NotTo(Equal(""))
						Expect(*result.BgpStatus).To(Equal("idle"))
						Expect(*result.CreatedAt).NotTo(Equal(""))
						Expect(*result.Crn).To(HavePrefix("crn:v1"))
						Expect(*result.Global).To(Equal(false))
						Expect(*result.Metered).To(Equal(false))
						Expect(*result.OperationalStatus).To(Equal("provisioned"))
						Expect(*result.Port.ID).To(Equal(*firstPort.ID))
						Expect(*result.ProviderApiManaged).To(Equal(true))
						Expect(*result.Type).To(Equal("connect"))
						Expect(*result.Vlan).Should(BeNumerically(">", 0))
						Expect(*result.LocationDisplayName).To(Equal(*firstPort.LocationDisplayName))
						Expect(*result.LocationName).To(Equal(*firstPort.LocationName))
						Expect(*result.ResourceGroup.ID).NotTo(Equal(""))
						break
					}

					// not provisioned yet, see if we have reached the timeout value.  If so, exit with failure
					if timer > 24 { // 2 min timer (24x5sec)
						Expect(*result.OperationalStatus).To(Equal("provisioned")) // timed out fail if status is not provisioned
						break
					} else {
						// Still exists, wait 5 sec
						time.Sleep(time.Duration(5) * time.Second)
						timer = timer + 1
					}
				}
			})

			It(`Successfully update name of the created gateway`, func() {
				shouldSkipTest()

				updateProviderGatewayOptions := new(directlinkproviderv2.UpdateProviderGatewayOptions)
				updateProviderGatewayOptions.ID = core.StringPtr(os.Getenv("GATEWAY_ID"))
				updateProviderGatewayOptions.Name = core.StringPtr(updatedGatewayName)

				result, detailedResponse, err := serviceV2.UpdateProviderGateway(updateProviderGatewayOptions)

				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(200))
				Expect(*result.Name).To(Equal(updatedGatewayName))
				Expect(*result.ID).To(Equal(os.Getenv("GATEWAY_ID")))
				Expect(*result.BgpAsn).To(Equal(bgpAsn))
				Expect(*result.SpeedMbps).To(Equal(speedMbps))
				Expect(*result.BgpCerCidr).NotTo(BeEmpty())
				Expect(*result.BgpIbmCidr).NotTo(Equal(""))
				Expect(*result.BgpIbmAsn).NotTo(Equal(""))
				Expect(*result.BgpStatus).To(Equal("idle"))
				Expect(*result.CreatedAt).NotTo(Equal(""))
				Expect(*result.Crn).To(HavePrefix("crn:v1"))
				Expect(*result.CustomerAccountID).To(Equal(customerAccount))
				Expect(*result.OperationalStatus).To(Equal("provisioned"))
				Expect(*result.Port.ID).To(Equal(*firstPort.ID))
				Expect(*result.ProviderApiManaged).To(Equal(true))
				Expect(*result.Type).To(Equal("connect"))
				Expect(*result.Vlan).NotTo(Equal(""))
				Expect(result.ChangeRequest).To(BeNil())

			})

			It(`Successfully request speed update the gateway`, func() {
				shouldSkipTest()

				updateProviderGatewayOptions := new(directlinkproviderv2.UpdateProviderGatewayOptions)
				updateProviderGatewayOptions.ID = core.StringPtr(os.Getenv("GATEWAY_ID"))
				updateProviderGatewayOptions.SpeedMbps = core.Int64Ptr(updatedSpeedMbps)

				result, detailedResponse, err := serviceV2.UpdateProviderGateway(updateProviderGatewayOptions)

				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(200))
				Expect(*result.Name).To(Equal(updatedGatewayName))
				Expect(*result.ID).To(Equal(os.Getenv("GATEWAY_ID")))
				Expect(*result.BgpAsn).To(Equal(bgpAsn))
				Expect(*result.SpeedMbps).To(Equal(speedMbps)) // speed does not change until client approves
				Expect(*result.BgpCerCidr).NotTo(BeEmpty())
				Expect(*result.BgpIbmCidr).NotTo(Equal(""))
				Expect(*result.BgpIbmAsn).NotTo(Equal(""))
				Expect(*result.BgpStatus).To(Equal("idle"))
				Expect(*result.CreatedAt).NotTo(Equal(""))
				Expect(*result.Crn).To(HavePrefix("crn:v1"))
				Expect(*result.CustomerAccountID).To(Equal(customerAccount))
				Expect(*result.OperationalStatus).To(Equal("provisioned"))
				Expect(*result.Port.ID).To(Equal(*firstPort.ID))
				Expect(*result.ProviderApiManaged).To(Equal(true))
				Expect(*result.Type).To(Equal("connect"))
				Expect(*result.Vlan).NotTo(Equal(""))
				Expect(result.ChangeRequest).NotTo(BeNil())
			})

			It(`Successfully approve gateway speed change using client account`, func() {
				shouldSkipTest()

				createGatewayActionOptions := serviceV1.NewCreateGatewayActionOptions(os.Getenv("GATEWAY_ID"),
					"update_attributes_approve")

				gatewayActionTemplateUpdatesItem := new(directlinkv1.GatewayActionTemplateUpdatesItemGatewayClientSpeedUpdate)
				gatewayActionTemplateUpdatesItem.SpeedMbps = core.Int64Ptr(updatedSpeedMbps)
				createGatewayActionOptions.Updates = []directlinkv1.GatewayActionTemplateUpdatesItemIntf{gatewayActionTemplateUpdatesItem}

				// Get the current status for the gateway
				result, detailedResponse, err := serviceV1.CreateGatewayAction(createGatewayActionOptions)
				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(200))

				Expect(*result.ID).To(Equal(os.Getenv("GATEWAY_ID")))
				Expect(*result.Name).To(Equal(updatedGatewayName))
				Expect(*result.BgpAsn).To(Equal(bgpAsn))
				Expect(*result.SpeedMbps).To(Equal(updatedSpeedMbps))
				Expect(*result.BgpCerCidr).NotTo(BeEmpty())
				Expect(*result.BgpIbmCidr).NotTo(Equal(""))
				Expect(*result.BgpIbmAsn).NotTo(Equal(""))
				Expect(*result.BgpStatus).To(Equal("idle"))
				Expect(*result.CreatedAt).NotTo(Equal(""))
				Expect(*result.Crn).To(HavePrefix("crn:v1"))
				Expect(*result.Global).To(Equal(false))
				Expect(*result.Metered).To(Equal(false))
				Expect(*result.OperationalStatus).To(Equal("provisioned"))
				Expect(*result.Port.ID).To(Equal(*firstPort.ID))
				Expect(*result.ProviderApiManaged).To(Equal(true))
				Expect(*result.Type).To(Equal("connect"))
				Expect(*result.ResourceGroup.ID).NotTo(Equal(""))
			})

			It("Successfully waits for updated gateway to go back to provisioned state", func() {
				shouldSkipTest()

				getProviderGatewayOptions := new(directlinkproviderv2.GetProviderGatewayOptions)
				getProviderGatewayOptions.ID = core.StringPtr(os.Getenv("GATEWAY_ID"))

				// before a connect gateway can be deleted, it needs to have operational_status of provisioned.  We need to wait for
				// the new gateway to go to provisioned so we can delete it.
				timer := 0
				for {
					// Get the current status for the gateway

					result, detailedResponse, err := serviceV2.GetProviderGateway(getProviderGatewayOptions)

					Expect(err).To(BeNil())
					Expect(detailedResponse.StatusCode).To(Equal(200))
					// if operational status is "provisioned" then we are done
					if *result.OperationalStatus == "provisioned" {
						Expect(*result.ID).To(Equal(os.Getenv("GATEWAY_ID")))
						Expect(*result.SpeedMbps).To(Equal(updatedSpeedMbps))
						Expect(*result.OperationalStatus).To(Equal("provisioned"))
						break
					}

					// not provisioned yet, see if we have reached the timeout value.  If so, exit with failure
					if timer > 24 { // 2 min timer (24x5sec)
						Expect(*result.OperationalStatus).To(Equal("provisioned")) // timed out fail if status is not provisioned
						break
					} else {
						// Still exists, wait 5 sec
						time.Sleep(time.Duration(5) * time.Second)
						timer = timer + 1
					}
				}
			})

			It("Successfully request gateway delete using provider account", func() {
				shouldSkipTest()

				gatewayId := os.Getenv("GATEWAY_ID")
				deteleGatewayOptions := serviceV2.NewDeleteProviderGatewayOptions(gatewayId)

				_, detailedResponse, _ := serviceV2.DeleteProviderGateway(deteleGatewayOptions)

				Expect(detailedResponse.StatusCode).To(Equal(202))
			})

			It(`Successfully reject gateway delete using client account`, func() {
				shouldSkipTest()

				createGatewayActionOptions := serviceV1.NewCreateGatewayActionOptions(os.Getenv("GATEWAY_ID"),
					"delete_gateway_reject")

				// Get the current status for the gateway
				_, detailedResponse, _ := serviceV1.CreateGatewayAction(createGatewayActionOptions)

				Expect(detailedResponse.StatusCode).To(Equal(200))
			})

			It(`Successfully verify reject gateway delete using client account`, func() {
				shouldSkipTest()

				getGatewayOptions := serviceV1.NewGetGatewayOptions(os.Getenv("GATEWAY_ID"))

				// Get the current status for the gateway
				result, detailedResponse, err := serviceV1.GetGateway(getGatewayOptions)

				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(200))

				// change request has been reset
				Expect(result.ChangeRequest).To(BeNil())
			})

			It("Successfully waits for gateway to go back to provisioned state after reject gateway delete", func() {
				shouldSkipTest()

				getProviderGatewayOptions := new(directlinkproviderv2.GetProviderGatewayOptions)
				getProviderGatewayOptions.ID = core.StringPtr(os.Getenv("GATEWAY_ID"))

				// before a connect gateway can be deleted, it needs to have operational_status of provisioned.  We need to wait for
				// the new gateway to go to provisioned so we can delete it.
				timer := 0
				for {
					// Get the current status for the gateway
					result, detailedResponse, err := serviceV2.GetProviderGateway(getProviderGatewayOptions)

					Expect(err).To(BeNil())
					Expect(detailedResponse.StatusCode).To(Equal(200))
					// if operational status is "provisioned" then we are done
					if *result.OperationalStatus == "provisioned" {
						Expect(*result.ID).To(Equal(os.Getenv("GATEWAY_ID")))
						Expect(*result.SpeedMbps).To(Equal(updatedSpeedMbps))
						Expect(*result.OperationalStatus).To(Equal("provisioned"))
						break
					}

					// not provisioned yet, see if we have reached the timeout value.  If so, exit with failure
					if timer > 24 { // 2 min timer (24x5sec)
						Expect(*result.OperationalStatus).To(Equal("provisioned")) // timed out fail if status is not provisioned
						break
					} else {
						// Still exists, wait 5 sec
						time.Sleep(time.Duration(5) * time.Second)
						timer = timer + 1
					}
				}
			})

			It("Successfully re-request gateway delete using provider account", func() {
				shouldSkipTest()

				gatewayId := os.Getenv("GATEWAY_ID")
				deteleGatewayOptions := serviceV2.NewDeleteProviderGatewayOptions(gatewayId)

				_, detailedResponse, _ := serviceV2.DeleteProviderGateway(deteleGatewayOptions)

				Expect(detailedResponse.StatusCode).To(Equal(202))
			})

			It(`Successfully approve gateway delete using client account`, func() {
				shouldSkipTest()

				createGatewayActionOptions := serviceV1.NewCreateGatewayActionOptions(os.Getenv("GATEWAY_ID"),
					"delete_gateway_approve")

				// Get the current status for the gateway
				_, detailedResponse, _ := serviceV1.CreateGatewayAction(createGatewayActionOptions)

				Expect(detailedResponse.StatusCode).To(Equal(204))
			})
		})
	})
})
