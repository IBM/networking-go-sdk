/**
 * (C) Copyright IBM Corp. 2021.
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

package directlinkv1_test

/*
	How to run this test:
	go test -v ./directlinkv1
*/

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
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

func getPortIdForConnect(ports []directlinkv1.Port) *directlinkv1.Port {
	providerToUse := "DL2-TEST"
	for _, port := range ports {
		if port.ProviderName != nil && strings.Contains(*port.ProviderName, providerToUse) {
			return &port
		}
	}
	return nil
}

var _ = Describe(`DirectLinkV1`, func() {
	defer GinkgoRecover()
	// Skip("Skipping")
	err := godotenv.Load("../directlink.env")
	It(`Successfully loading .env file`, func() {
		if err == nil {
			serviceURL := os.Getenv("SERVICE_URL")
			if serviceURL != "" {
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
	export_route_filters_id := ""
	import_route_filters_id := ""
	etag := ""

	service, err := directlinkv1.NewDirectLinkV1UsingExternalConfig(options)
	It(`Successfully created DirectLinkV1 service instance`, func() {
		shouldSkipTest()
		Expect(err).To(BeNil())
	})

	Describe("Direct Link Gateways", func() {
		timestamp := time.Now().Unix()
		gatewayName := "GO-INT-SDK-" + strconv.FormatInt(timestamp, 10)
		updatedGatewayName := "GO-INT-SDK-PATCH-" + strconv.FormatInt(timestamp, 10)
		bgpAsn := int64(64999)
		crossConnectRouter := "LAB-xcr01.dal09"
		global := true
		locationName := os.Getenv("LOCATION_NAME")
		speedMbps := int64(1000)
		metered := false
		carrierName := "carrier1"
		customerName := "customer1"
		gatewayType := "dedicated"

		// Construct an instance of the GatewayTemplateRouteFilter model
		gatewayTemplateRouteFilterModel := new(directlinkv1.GatewayTemplateRouteFilter)
		gatewayTemplateRouteFilterModel.Action = core.StringPtr("permit")
		gatewayTemplateRouteFilterModel.Ge = core.Int64Ptr(int64(25))
		gatewayTemplateRouteFilterModel.Le = core.Int64Ptr(int64(30))
		gatewayTemplateRouteFilterModel.Prefix = core.StringPtr("192.168.100.0/24")

		model := []directlinkv1.GatewayTemplateRouteFilter{*gatewayTemplateRouteFilterModel}

		invalidGatewayId := "000000000000000000000000000000000000"

		Context("Get non existing gateway", func() {

			getGatewayOptions := service.NewGetGatewayOptions(invalidGatewayId)

			It(`Returns the http response with error code 404`, func() {
				shouldSkipTest()
				result, detailedResponse, err := service.GetGateway(getGatewayOptions)
				Expect(result).To(BeNil())
				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(Equal("Cannot find Gateway"))
				Expect(detailedResponse.StatusCode).To(Equal(404))

			})
		})

		Context("Create gateway", func() {
			gateway, _ := service.NewGatewayTemplateGatewayTypeDedicatedTemplate(bgpAsn, global, metered, gatewayName, speedMbps, gatewayType, carrierName, crossConnectRouter, customerName, locationName)
			gateway.DefaultExportRouteFilter = core.StringPtr("permit")
			gateway.DefaultImportRouteFilter = core.StringPtr("permit")
			gateway.ExportRouteFilters = model
			gateway.ImportRouteFilters = model

			createGatewayOptions := service.NewCreateGatewayOptions(gateway)

			It(`Fails when Invalid BGP is provided`, func() {
				shouldSkipTest()

				gateway, _ := service.NewGatewayTemplateGatewayTypeDedicatedTemplate(65500, global, metered, gatewayName, speedMbps, gatewayType, carrierName, crossConnectRouter, customerName, locationName)
				createGatewayOptions := service.NewCreateGatewayOptions(gateway)
				result, detailedResponse, err := service.CreateGateway(createGatewayOptions)
				Expect(result).To(BeNil())
				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(Equal("BGP AS Number is invalid."))
				Expect(detailedResponse.StatusCode).To(Equal(400))
			})

			It(`Fails when invalid speed_mbps is provided`, func() {
				shouldSkipTest()

				gateway, _ := service.NewGatewayTemplateGatewayTypeDedicatedTemplate(bgpAsn, global, metered, gatewayName, 10000000000, gatewayType, carrierName, crossConnectRouter, customerName, locationName)

				createGatewayOptions := service.NewCreateGatewayOptions(gateway)

				result, detailedResponse, err := service.CreateGateway(createGatewayOptions)
				Expect(result).To(BeNil())
				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(Equal("Cannot find Location with provided 'linkSpeed' and 'OfferingType'."))
				Expect(detailedResponse.StatusCode).To(Equal(400))
			})

			It(`Fails when invalid locations is provided`, func() {
				shouldSkipTest()

				gateway, _ := service.NewGatewayTemplateGatewayTypeDedicatedTemplate(bgpAsn, global, metered, gatewayName, speedMbps, gatewayType, carrierName, crossConnectRouter, customerName, "InvalidCity")

				createGatewayOptions := service.NewCreateGatewayOptions(gateway)

				result, detailedResponse, err := service.CreateGateway(createGatewayOptions)
				Expect(result).To(BeNil())
				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(Equal("Cannot find Location with provided 'linkSpeed' and 'OfferingType'."))
				Expect(detailedResponse.StatusCode).To(Equal(400))
			})

			It(`Successfully Creates a gateway`, func() {
				shouldSkipTest()

				gateway, _ := service.NewGatewayTemplateGatewayTypeDedicatedTemplate(bgpAsn, global, metered, gatewayName, speedMbps, gatewayType, carrierName, crossConnectRouter, customerName, locationName)
				gateway.DefaultExportRouteFilter = core.StringPtr("permit")
				gateway.DefaultImportRouteFilter = core.StringPtr("permit")
				gateway.ExportRouteFilters = model
				gateway.ImportRouteFilters = model

				createGatewayOptions := service.NewCreateGatewayOptions(gateway)

				result, detailedResponse, err := service.CreateGateway(createGatewayOptions)

				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(201))

				os.Setenv("GATEWAY_ID", *result.ID)

				Expect(*result.Name).To(Equal(gatewayName))
				Expect(*result.BgpAsn).To(Equal(bgpAsn))
				Expect(*result.Global).To(Equal(global))
				Expect(*result.Metered).To(Equal(metered))
				Expect(*result.SpeedMbps).To(Equal(speedMbps))
				Expect(*result.Type).To(Equal(gatewayType))
				Expect(*result.CrossConnectRouter).To(Equal(crossConnectRouter))
				Expect(*result.LocationName).To(Equal(locationName))
				Expect(*result.LocationDisplayName).NotTo(Equal(""))
				Expect(*result.BgpCerCidr).NotTo(BeEmpty())
				Expect(*result.BgpIbmCidr).NotTo(Equal(""))
				Expect(*result.BgpIbmAsn).NotTo(Equal(""))
				Expect(*result.BgpStatus).To(Equal("idle"))
				Expect(*result.CreatedAt).NotTo(Equal(""))
				Expect(*result.Crn).To(HavePrefix("crn:v1"))
				Expect(*result.LinkStatus).To(Equal("down"))
				Expect(*result.OperationalStatus).To(Equal("awaiting_loa"))
				Expect(*result.ResourceGroup.ID).NotTo(Equal(""))

				Expect(result.DefaultExportRouteFilter).To(Equal(core.StringPtr("permit")))
				Expect(result.DefaultImportRouteFilter).To(Equal(core.StringPtr("permit")))
			})

			It(`Successfully fetches the created Gateway`, func() {
				shouldSkipTest()

				gatewayId := os.Getenv("GATEWAY_ID")
				getGatewayOptions := service.NewGetGatewayOptions(gatewayId)

				result, detailedResponse, err := service.GetGateway(getGatewayOptions)

				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(200))

				Expect(*result.ID).To(Equal(gatewayId))
				Expect(*result.Name).To(Equal(gatewayName))
				Expect(*result.BgpAsn).To(Equal(bgpAsn))
				Expect(*result.Global).To(Equal(global))
				Expect(*result.Metered).To(Equal(metered))
				Expect(*result.SpeedMbps).To(Equal(speedMbps))
				Expect(*result.Type).To(Equal(gatewayType))
				Expect(*result.CrossConnectRouter).To(Equal(crossConnectRouter))
				Expect(*result.LocationName).To(Equal(locationName))

				Expect(result.DefaultExportRouteFilter).To(Equal(core.StringPtr("permit")))
				Expect(result.DefaultImportRouteFilter).To(Equal(core.StringPtr("permit")))
			})

			It(`Throws an Error when creating a gateway with same name`, func() {
				shouldSkipTest()

				result, detailedResponse, err := service.CreateGateway(createGatewayOptions)
				Expect(result).To(BeNil())
				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(Equal("A gateway with the same name already exists"))
				Expect(detailedResponse.StatusCode).To(Equal(409))
			})

			It(`Successfully list all gateways`, func() {
				shouldSkipTest()
				listGatewaysOptions := service.NewListGatewaysOptions()

				result, detailedResponse, err := service.ListGateways(listGatewaysOptions)
				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(200))

				gateways := result.Gateways
				Expect(len(gateways)).Should(BeNumerically(">", 0))
				found := false
				// find the created gateway and verify the attributes
				gatewayId := os.Getenv("GATEWAY_ID")
				for _, gw := range gateways {
					if *gw.ID == gatewayId {
						found = true
						Expect(*gw.Name).To(Equal(gatewayName))
						Expect(*gw.BgpAsn).To(Equal(bgpAsn))
						Expect(*gw.Global).To(Equal(global))
						Expect(*gw.Metered).To(Equal(metered))
						Expect(*gw.SpeedMbps).To(Equal(speedMbps))
						Expect(*gw.Type).To(Equal(gatewayType))
						Expect(*gw.CrossConnectRouter).To(Equal(crossConnectRouter))
						Expect(*gw.LocationName).To(Equal(locationName))
						Expect(*gw.LocationDisplayName).NotTo(Equal(""))
						Expect(*gw.BgpCerCidr).NotTo(BeEmpty())
						Expect(*gw.BgpIbmCidr).NotTo(Equal(""))
						Expect(*gw.BgpIbmAsn).NotTo(Equal(""))
						Expect(*gw.BgpStatus).To(Equal("idle"))
						Expect(*gw.CreatedAt).NotTo(Equal(""))
						Expect(*gw.Crn).To(HavePrefix("crn:v1"))
						Expect(*gw.LinkStatus).To(Equal("down"))
						Expect(*gw.OperationalStatus).To(Equal("awaiting_loa"))
						Expect(*gw.ResourceGroup.ID).NotTo(Equal(""))
						Expect(gw.DefaultExportRouteFilter).To(Equal(core.StringPtr("permit")))
						Expect(gw.DefaultImportRouteFilter).To(Equal(core.StringPtr("permit")))
						break
					}
				}
				// expect the created gateway to have been found.  If not found, throw an error
				Expect(found).To(Equal(true))
			})

			It(`Successfully create export route filters for a Gateway`, func() {
				shouldSkipTest()
				gatewayId := os.Getenv("GATEWAY_ID")
				// Construct an instance of the CreateGatewayExportRouteFilterOptions model
				createGatewayExportRouteFilterOptionsModel := new(directlinkv1.CreateGatewayExportRouteFilterOptions)
				createGatewayExportRouteFilterOptionsModel.GatewayID = core.StringPtr(gatewayId)
				createGatewayExportRouteFilterOptionsModel.Action = core.StringPtr("permit")
				createGatewayExportRouteFilterOptionsModel.Prefix = core.StringPtr("192.168.100.0/24")
				//createGatewayExportRouteFilterOptionsModel.Before = core.StringPtr("1a15dcab-7e40-45e1-b7c5-bc690eaa9782")
				createGatewayExportRouteFilterOptionsModel.Ge = core.Int64Ptr(int64(25))
				createGatewayExportRouteFilterOptionsModel.Le = core.Int64Ptr(int64(30))
				createGatewayExportRouteFilterOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, detailedResponse, operationErr := service.CreateGatewayExportRouteFilter(createGatewayExportRouteFilterOptionsModel)
				if operationErr != nil {
					fmt.Println(operationErr)
				}
				fmt.Printf("Gateway Id %v", gatewayId)
				fmt.Printf("Export route filter Id %v", *result.ID)
				Expect(operationErr).To(BeNil())
				Expect(result).ToNot(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(201))
				Expect(result.Action).To(Equal(core.StringPtr("permit")))
				//Expect(result.Before).To(Equal("1a15dcab-7e40-45e1-b7c5-bc690eaa9782"))
				Expect(result.Prefix).To(Equal(core.StringPtr("192.168.100.0/24")))
				Expect(result.Ge).To(Equal(core.Int64Ptr(int64(25))))
				Expect(result.Le).To(Equal(core.Int64Ptr(int64(30))))
				export_route_filters_id = *result.ID
				Expect(result.CreatedAt).ToNot(BeNil())
				Expect(result.UpdatedAt).ToNot(BeNil())
			})

			It(`Successfully fetches the export route filters for a Gateway`, func() {
				shouldSkipTest()

				gatewayId := os.Getenv("GATEWAY_ID")

				// Construct an instance of the ListGatewayExportRouteFiltersOptions model
				listGatewayExportRouteFiltersOptionsModel := new(directlinkv1.ListGatewayExportRouteFiltersOptions)
				listGatewayExportRouteFiltersOptionsModel.GatewayID = core.StringPtr(gatewayId)
				// Expect response parsing to fail since we are receiving a text/plain response
				result, detailedResponse, operationErr := service.ListGatewayExportRouteFilters(listGatewayExportRouteFiltersOptionsModel)

				Expect(operationErr).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(200))
				Expect(result.ExportRouteFilters[0].Action).To(Equal(core.StringPtr("permit")))
				Expect(result.ExportRouteFilters[0].Ge).To(Equal(core.Int64Ptr(int64(25))))
				Expect(result.ExportRouteFilters[0].Le).To(Equal(core.Int64Ptr(int64(30))))
				etag = detailedResponse.GetHeaders().Get("etag")
				fmt.Printf("****** get export etag %v", etag)
			})
			It(`Successfully replace existing export route filters for a Gateway`, func() {

				shouldSkipTest()
				gatewayId := os.Getenv("GATEWAY_ID")

				// Construct an instance of the GatewayTemplateRouteFilter model
				gatewayTemplateRouteFilterModel := new(directlinkv1.GatewayTemplateRouteFilter)
				gatewayTemplateRouteFilterModel.Action = core.StringPtr("permit")
				gatewayTemplateRouteFilterModel.Ge = core.Int64Ptr(int64(25))
				gatewayTemplateRouteFilterModel.Le = core.Int64Ptr(int64(30))
				gatewayTemplateRouteFilterModel.Prefix = core.StringPtr("192.168.100.0/24")

				// Construct an instance of the ReplaceGatewayExportRouteFiltersOptions model
				replaceGatewayExportRouteFiltersOptionsModel := new(directlinkv1.ReplaceGatewayExportRouteFiltersOptions)
				replaceGatewayExportRouteFiltersOptionsModel.GatewayID = core.StringPtr(gatewayId)
				replaceGatewayExportRouteFiltersOptionsModel.ExportRouteFilters = []directlinkv1.GatewayTemplateRouteFilter{*gatewayTemplateRouteFilterModel}
				replaceGatewayExportRouteFiltersOptionsModel.IfMatch = core.StringPtr(etag)
				replaceGatewayExportRouteFiltersOptionsModel.Headers = map[string]string{"If-Match": etag}
				fmt.Printf("****** replace export etag %v", etag)

				// Invoke operation with valid options model (positive test)
				result, response, operationErr := service.ReplaceGatewayExportRouteFilters(replaceGatewayExportRouteFiltersOptionsModel)
				if operationErr != nil {
					fmt.Printf("Printing Error %v ", operationErr)
					fmt.Printf("Gateway Id %v ", gatewayId)
				}
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
				export_route_filters_id = *result.ExportRouteFilters[0].ID
			})

			It(`Successfully get export route filters for a Gateway`, func() {

				shouldSkipTest()
				gatewayId := os.Getenv("GATEWAY_ID")

				// Construct an instance of the GetGatewayExportRouteFilterOptions model
				getGatewayExportRouteFilterOptionsModel := new(directlinkv1.GetGatewayExportRouteFilterOptions)
				getGatewayExportRouteFilterOptionsModel.GatewayID = core.StringPtr(gatewayId)
				getGatewayExportRouteFilterOptionsModel.ID = core.StringPtr(export_route_filters_id)
				getGatewayExportRouteFilterOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				result, detailedResponse, operationErr := service.GetGatewayExportRouteFilter(getGatewayExportRouteFilterOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(result).ToNot(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(200))
				Expect(result.Action).To(Equal(core.StringPtr("permit")))
				Expect(result.Prefix).To(Equal(core.StringPtr("192.168.100.0/24")))
				Expect(result.Ge).To(Equal(core.Int64Ptr(int64(25))))
				Expect(result.Le).To(Equal(core.Int64Ptr(int64(30))))
				Expect(result.CreatedAt).ToNot(BeNil())
				Expect(result.UpdatedAt).ToNot(BeNil())
				etag = detailedResponse.GetHeaders().Get("etag")
			})

			It(`Successfully update export route filters for a Gateway`, func() {

				shouldSkipTest()
				gatewayId := os.Getenv("GATEWAY_ID")

				// Construct an instance of the UpdateRouteFilterTemplate model
				updateRouteFilterTemplateModel := new(directlinkv1.UpdateRouteFilterTemplate)
				updateRouteFilterTemplateModel.Action = core.StringPtr("deny")
				updateRouteFilterTemplateModel.Ge = core.Int64Ptr(int64(24))
				updateRouteFilterTemplateModel.Le = core.Int64Ptr(int64(30))
				updateRouteFilterTemplateModel.Prefix = core.StringPtr("192.168.100.0/24")
				updateRouteFilterTemplateModelAsPatch, asPatchErr := updateRouteFilterTemplateModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the CreateGatewayExportRouteFilterOptions model
				updateGatewayExportRouteFilterOptionsModel := new(directlinkv1.UpdateGatewayExportRouteFilterOptions)
				updateGatewayExportRouteFilterOptionsModel.GatewayID = core.StringPtr(gatewayId)
				updateGatewayExportRouteFilterOptionsModel.ID = core.StringPtr(export_route_filters_id)
				updateGatewayExportRouteFilterOptionsModel.UpdateRouteFilterTemplatePatch = updateRouteFilterTemplateModelAsPatch
				// Expect response parsing to fail since we are receiving a text/plain response
				result, detailedResponse, operationErr := service.UpdateGatewayExportRouteFilter(updateGatewayExportRouteFilterOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(result).ToNot(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(200))
				Expect(result.Action).To(Equal(core.StringPtr("deny")))
				// Expect(result.Before).To(Equal("1a15dcab-7e40-45e1-b7c5-bc690eaa9782"))
				Expect(result.Prefix).To(Equal(core.StringPtr("192.168.100.0/24")))
				Expect(result.Ge).To(Equal(core.Int64Ptr(int64(24))))
				Expect(result.Le).To(Equal(core.Int64Ptr(int64(30))))
				Expect(result.CreatedAt).ToNot(BeNil())
				Expect(result.UpdatedAt).ToNot(BeNil())

			})

			It(`Successfully delete export route filters for a Gateway`, func() {

				shouldSkipTest()
				gatewayId := os.Getenv("GATEWAY_ID")

				// Construct an instance of the DeleteGatewayExportRouteFilterOptions model
				deleteGatewayExportRouteFilterOptionsModel := new(directlinkv1.DeleteGatewayExportRouteFilterOptions)
				deleteGatewayExportRouteFilterOptionsModel.GatewayID = core.StringPtr(gatewayId)
				deleteGatewayExportRouteFilterOptionsModel.ID = core.StringPtr(export_route_filters_id)

				// Invoke operation with valid options model (positive test)
				response, operationErr := service.DeleteGatewayExportRouteFilter(deleteGatewayExportRouteFilterOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})

			It(`Successfully create import route filters for a Gateway`, func() {
				shouldSkipTest()
				gatewayId := os.Getenv("GATEWAY_ID")
				// Construct an instance of the createGatewayImportRouteFilterOptionsModel model
				createGatewayImportRouteFilterOptionsModel := new(directlinkv1.CreateGatewayImportRouteFilterOptions)
				createGatewayImportRouteFilterOptionsModel.GatewayID = core.StringPtr(gatewayId)
				createGatewayImportRouteFilterOptionsModel.Action = core.StringPtr("permit")
				createGatewayImportRouteFilterOptionsModel.Prefix = core.StringPtr("192.168.100.0/24")
				//createGatewayImportRouteFilterOptionsModel.Before = core.StringPtr("1a15dcab-7e40-45e1-b7c5-bc690eaa9782")
				createGatewayImportRouteFilterOptionsModel.Ge = core.Int64Ptr(int64(25))
				createGatewayImportRouteFilterOptionsModel.Le = core.Int64Ptr(int64(30))
				createGatewayImportRouteFilterOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, detailedResponse, operationErr := service.CreateGatewayImportRouteFilter(createGatewayImportRouteFilterOptionsModel)
				fmt.Printf("Gateway Id %v", gatewayId)
				fmt.Printf("Import route filter Id %v", *result.ID)
				import_route_filters_id = *result.ID
				Expect(operationErr).To(BeNil())
				Expect(result).ToNot(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(201))
				Expect(result.Action).To(Equal(core.StringPtr("permit")))
				//Expect(result.Before).To(Equal("1a15dcab-7e40-45e1-b7c5-bc690eaa9782"))
				Expect(result.Prefix).To(Equal(core.StringPtr("192.168.100.0/24")))
				Expect(result.Ge).To(Equal(core.Int64Ptr(int64(25))))
				Expect(result.Le).To(Equal(core.Int64Ptr(int64(30))))
				Expect(result.CreatedAt).ToNot(BeNil())
				Expect(result.UpdatedAt).ToNot(BeNil())

			})

			It(`Successfully fetches the import route filters for a Gateway`, func() {
				shouldSkipTest()

				gatewayId := os.Getenv("GATEWAY_ID")

				// Construct an instance of the ListGatewayImportRouteFiltersOptions model
				listGatewayImportRouteFiltersOptionsModel := new(directlinkv1.ListGatewayImportRouteFiltersOptions)
				listGatewayImportRouteFiltersOptionsModel.GatewayID = core.StringPtr(gatewayId)
				// Expect response parsing to fail since we are receiving a text/plain response
				result, detailedResponse, operationErr := service.ListGatewayImportRouteFilters(listGatewayImportRouteFiltersOptionsModel)

				Expect(operationErr).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(200))
				Expect(result.ImportRouteFilters[0].Action).To(Equal(core.StringPtr("permit")))
				Expect(result.ImportRouteFilters[0].Ge).To(Equal(core.Int64Ptr(int64(25))))
				Expect(result.ImportRouteFilters[0].Le).To(Equal(core.Int64Ptr(int64(30))))
				etag = detailedResponse.GetHeaders().Get("etag")
				fmt.Printf("****** get import etag %v", etag)
			})
			It(`Successfully replace existing import route filters for a Gateway`, func() {

				shouldSkipTest()
				gatewayId := os.Getenv("GATEWAY_ID")

				// Construct an instance of the GatewayTemplateRouteFilter model
				gatewayTemplateRouteFilterModel := new(directlinkv1.GatewayTemplateRouteFilter)
				gatewayTemplateRouteFilterModel.Action = core.StringPtr("permit")
				gatewayTemplateRouteFilterModel.Ge = core.Int64Ptr(int64(25))
				gatewayTemplateRouteFilterModel.Le = core.Int64Ptr(int64(30))
				gatewayTemplateRouteFilterModel.Prefix = core.StringPtr("192.168.100.0/24")

				// Construct an instance of the ReplaceGatewayImportRouteFiltersOptions model
				replaceGatewayImportRouteFiltersOptionsModel := new(directlinkv1.ReplaceGatewayImportRouteFiltersOptions)
				replaceGatewayImportRouteFiltersOptionsModel.GatewayID = core.StringPtr(gatewayId)
				replaceGatewayImportRouteFiltersOptionsModel.ImportRouteFilters = []directlinkv1.GatewayTemplateRouteFilter{*gatewayTemplateRouteFilterModel}
				replaceGatewayImportRouteFiltersOptionsModel.IfMatch = core.StringPtr(etag)
				replaceGatewayImportRouteFiltersOptionsModel.Headers = map[string]string{"If-Match": etag}
				fmt.Printf("****** replace import etag %v", etag)

				// Invoke operation with valid options model (positive test)
				result, response, operationErr := service.ReplaceGatewayImportRouteFilters(replaceGatewayImportRouteFiltersOptionsModel)
				if operationErr != nil {
					fmt.Printf("Printing Error %v ", operationErr)
					fmt.Printf("Gateway Id %v ", gatewayId)
				}
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
				import_route_filters_id = *result.ImportRouteFilters[0].ID
			})

			It(`Successfully get import route filters for a Gateway`, func() {

				shouldSkipTest()
				gatewayId := os.Getenv("GATEWAY_ID")

				// Construct an instance of the GetGatewayImportRouteFilterOptions model
				getGatewayImportRouteFilterOptionsModel := new(directlinkv1.GetGatewayImportRouteFilterOptions)
				getGatewayImportRouteFilterOptionsModel.GatewayID = core.StringPtr(gatewayId)
				getGatewayImportRouteFilterOptionsModel.ID = core.StringPtr(import_route_filters_id)

				result, detailedResponse, operationErr := service.GetGatewayImportRouteFilter(getGatewayImportRouteFilterOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(result).ToNot(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(200))
				Expect(result.Action).To(Equal(core.StringPtr("permit")))
				Expect(result.Prefix).To(Equal(core.StringPtr("192.168.100.0/24")))
				Expect(result.Ge).To(Equal(core.Int64Ptr(int64(25))))
				Expect(result.Le).To(Equal(core.Int64Ptr(int64(30))))
				Expect(result.CreatedAt).ToNot(BeNil())
				Expect(result.UpdatedAt).ToNot(BeNil())
				etag = detailedResponse.GetHeaders().Get("etag")

			})

			It(`Successfully update import route filters for a Gateway`, func() {

				shouldSkipTest()
				gatewayId := os.Getenv("GATEWAY_ID")

				// Construct an instance of the UpdateRouteFilterTemplate model
				updateRouteFilterTemplateModel := new(directlinkv1.UpdateRouteFilterTemplate)
				updateRouteFilterTemplateModel.Action = core.StringPtr("deny")
				updateRouteFilterTemplateModel.Ge = core.Int64Ptr(int64(24))
				updateRouteFilterTemplateModel.Le = core.Int64Ptr(int64(25))
				updateRouteFilterTemplateModel.Prefix = core.StringPtr("192.168.100.0/24")
				updateRouteFilterTemplateModelAsPatch, asPatchErr := updateRouteFilterTemplateModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the updateGatewayImportRouteFilterOptionsModel model
				updateGatewayImportRouteFilterOptionsModel := new(directlinkv1.UpdateGatewayImportRouteFilterOptions)
				updateGatewayImportRouteFilterOptionsModel.GatewayID = core.StringPtr(gatewayId)
				updateGatewayImportRouteFilterOptionsModel.ID = core.StringPtr(import_route_filters_id)
				updateGatewayImportRouteFilterOptionsModel.UpdateRouteFilterTemplatePatch = updateRouteFilterTemplateModelAsPatch
				// Expect response parsing to fail since we are receiving a text/plain response
				result, detailedResponse, operationErr := service.UpdateGatewayImportRouteFilter(updateGatewayImportRouteFilterOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(result).ToNot(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(200))
				Expect(result.Action).To(Equal(core.StringPtr("deny")))
				// Expect(result.Before).To(Equal("1a15dcab-7e40-45e1-b7c5-bc690eaa9782"))
				Expect(result.Prefix).To(Equal(core.StringPtr("192.168.100.0/24")))
				Expect(result.Ge).To(Equal(core.Int64Ptr(int64(24))))
				Expect(result.Le).To(Equal(core.Int64Ptr(int64(25))))
				Expect(result.CreatedAt).ToNot(BeNil())
				Expect(result.UpdatedAt).ToNot(BeNil())

			})

			It(`Successfully get import route filters for a Gateway`, func() {

				shouldSkipTest()
				gatewayId := os.Getenv("GATEWAY_ID")

				// Construct an instance of the GetGatewayImportRouteFilterOptions model
				getGatewayImportRouteFilterOptionsModel := new(directlinkv1.GetGatewayImportRouteFilterOptions)
				getGatewayImportRouteFilterOptionsModel.GatewayID = core.StringPtr(gatewayId)
				getGatewayImportRouteFilterOptionsModel.ID = core.StringPtr(import_route_filters_id)

				result, detailedResponse, operationErr := service.GetGatewayImportRouteFilter(getGatewayImportRouteFilterOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(result).ToNot(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(200))
				Expect(result.Action).To(Equal(core.StringPtr("deny")))
				Expect(result.Prefix).To(Equal(core.StringPtr("192.168.100.0/24")))
				Expect(result.Ge).To(Equal(core.Int64Ptr(int64(24))))
				Expect(result.Le).To(Equal(core.Int64Ptr(int64(25))))
				Expect(result.CreatedAt).ToNot(BeNil())
				Expect(result.UpdatedAt).ToNot(BeNil())
				etag = detailedResponse.GetHeaders().Get("etag")

			})

			It(`Successfully delete import route filters for a Gateway`, func() {

				shouldSkipTest()
				gatewayId := os.Getenv("GATEWAY_ID")

				// Construct an instance of the DeleteGatewayImportRouteFilterOptions model
				deleteGatewayImportRouteFilterOptionsModel := new(directlinkv1.DeleteGatewayImportRouteFilterOptions)
				deleteGatewayImportRouteFilterOptionsModel.GatewayID = core.StringPtr(gatewayId)
				deleteGatewayImportRouteFilterOptionsModel.ID = core.StringPtr(import_route_filters_id)

				// Invoke operation with valid options model (positive test)
				response, operationErr := service.DeleteGatewayImportRouteFilter(deleteGatewayImportRouteFilterOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})

		})

		Context("Fail update Gateway", func() {
			It(`Fails if an invalid GatewayID is provided`, func() {
				shouldSkipTest()

				patchGatewayOptions := service.NewUpdateGatewayOptions(invalidGatewayId).SetOperationalStatus("loa_accepted")

				result, detailedResponse, err := service.UpdateGateway(patchGatewayOptions)
				Expect(result).To(BeNil())
				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(Equal("Invalid Gateway Id."))
				Expect(detailedResponse.StatusCode).To(Equal(400))
			})

			It(`Successfully Updates the Gateway`, func() {
				shouldSkipTest()

				gatewayId := os.Getenv("GATEWAY_ID")
				patchGatewayOptions := service.NewUpdateGatewayOptions(gatewayId)
				patchGatewayOptions.SetDefaultImportRouteFilter("deny")
				patchGatewayOptions.SetDefaultExportRouteFilter("deny")

				result, detailedResponse, err := service.UpdateGateway(patchGatewayOptions.SetGlobal(false).SetSpeedMbps(int64(1000)).SetName(updatedGatewayName))
				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(200))

				Expect(*result.ID).To(Equal(gatewayId))
				Expect(*result.Name).To(Equal(updatedGatewayName))
				Expect(*result.BgpAsn).To(Equal(bgpAsn))
				Expect(*result.Global).To(Equal(false))
				Expect(*result.Metered).To(Equal(metered))
				Expect(*result.SpeedMbps).To(Equal(speedMbps))
				Expect(*result.Type).To(Equal(gatewayType))
				Expect(*result.CrossConnectRouter).To(Equal(crossConnectRouter))
				Expect(*result.LocationName).To(Equal(locationName))
				Expect(*result.LocationDisplayName).NotTo(Equal(""))
				Expect(*result.BgpCerCidr).NotTo(BeEmpty())
				Expect(*result.BgpIbmCidr).NotTo(Equal(""))
				Expect(*result.BgpIbmAsn).NotTo(Equal(""))
				Expect(*result.BgpStatus).To(Equal("idle"))
				Expect(*result.CreatedAt).NotTo(Equal(""))
				Expect(*result.Crn).To(HavePrefix("crn:v1"))
				Expect(*result.LinkStatus).To(Equal("down"))
				Expect(*result.OperationalStatus).To(Equal("awaiting_loa"))
				Expect(*result.ResourceGroup.ID).NotTo(Equal(""))
				Expect(*result.DefaultExportRouteFilter).To(Equal("deny"))
				Expect(*result.DefaultImportRouteFilter).To(Equal("deny"))
			})

			It(`Successfully fetches the updated Gateway`, func() {
				shouldSkipTest()

				gatewayId := os.Getenv("GATEWAY_ID")
				getGatewayOptions := service.NewGetGatewayOptions(gatewayId)

				result, detailedResponse, err := service.GetGateway(getGatewayOptions)
				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(200))

				Expect(*result.ID).To(Equal(gatewayId))
				Expect(*result.Name).To(Equal(updatedGatewayName))
				Expect(*result.BgpAsn).To(Equal(bgpAsn))
				Expect(*result.Global).To(Equal(false))
				Expect(*result.Metered).To(Equal(metered))
				Expect(*result.SpeedMbps).To(Equal(speedMbps))
				Expect(*result.Type).To(Equal(gatewayType))
				Expect(*result.CrossConnectRouter).To(Equal(crossConnectRouter))
				Expect(*result.LocationName).To(Equal(locationName))
				Expect(*result.LocationDisplayName).NotTo(Equal(""))
				Expect(*result.BgpCerCidr).NotTo(BeEmpty())
				Expect(*result.BgpIbmCidr).NotTo(Equal(""))
				Expect(*result.BgpIbmAsn).NotTo(Equal(""))
				Expect(*result.BgpStatus).To(Equal("idle"))
				Expect(*result.CreatedAt).NotTo(Equal(""))
				Expect(*result.Crn).To(HavePrefix("crn:v1"))
				Expect(*result.LinkStatus).To(Equal("down"))
				Expect(*result.OperationalStatus).To(Equal("awaiting_loa"))
				Expect(*result.ResourceGroup.ID).NotTo(Equal(""))
				Expect(*result.DefaultExportRouteFilter).To(Equal("deny"))
				Expect(*result.DefaultImportRouteFilter).To(Equal("deny"))
			})
		})

		Context("Delete a gateway", func() {
			It(`Fails if an invalid GatewayID is provided`, func() {
				shouldSkipTest()

				deteleGatewayOptions := service.NewDeleteGatewayOptions(invalidGatewayId)

				detailedResponse, err := service.DeleteGateway(deteleGatewayOptions)
				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(Equal("Invalid Gateway Id."))
				Expect(detailedResponse.StatusCode).To(Equal(400))
			})

			It(`Successfully deletes a gateway`, func() {
				shouldSkipTest()

				gatewayId := os.Getenv("GATEWAY_ID")
				deteleGatewayOptions := service.NewDeleteGatewayOptions(gatewayId)

				detailedResponse, err := service.DeleteGateway(deteleGatewayOptions)
				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(204))
			})
		})

		Context("DirectLink connect gateway", func() {

			// to create a connect gateway, we need to have a port.  List the ports and save the id of the 1st one found
			portId := ""
			portLocationDisplayName := ""
			portLocationName := ""
			timestamp := time.Now().Unix()

			It(`List ports and save the id of the first port`, func() {
				shouldSkipTest()

				listPortsOptions := service.NewListPortsOptions()
				result, detailedResponse, err := service.ListPorts(listPortsOptions)
				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(200))
				port := getPortIdForConnect(result.Ports)
				portId = *port.ID
				portLocationDisplayName = *port.LocationDisplayName
				portLocationName = *port.LocationName
			})

			It(`create connect gateway`, func() {
				shouldSkipTest()

				gatewayName = "GO-INT-SDK-CONNECT-" + strconv.FormatInt(timestamp, 10)
				portIdentity, _ := service.NewGatewayPortIdentity(portId)
				gateway, _ := service.NewGatewayTemplateGatewayTypeConnectTemplate(bgpAsn, global, metered, gatewayName, speedMbps, "connect", portIdentity)
				// Construct an instance of the GatewayTemplateRouteFilter model
				gatewayTemplateRouteFilterModel := new(directlinkv1.GatewayTemplateRouteFilter)
				gatewayTemplateRouteFilterModel.Action = core.StringPtr("permit")
				gatewayTemplateRouteFilterModel.Ge = core.Int64Ptr(int64(25))
				gatewayTemplateRouteFilterModel.Le = core.Int64Ptr(int64(30))
				gatewayTemplateRouteFilterModel.Prefix = core.StringPtr("192.168.100.0/24")

				connect_model := []directlinkv1.GatewayTemplateRouteFilter{*gatewayTemplateRouteFilterModel}
				gateway.DefaultExportRouteFilter = core.StringPtr("deny")
				gateway.DefaultImportRouteFilter = core.StringPtr("deny")
				gateway.ImportRouteFilters = connect_model
				gateway.ExportRouteFilters = connect_model

				createGatewayOptions := service.NewCreateGatewayOptions(gateway)
				result, detailedResponse, err := service.CreateGateway(createGatewayOptions)

				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(201))

				// Save the gateway id for deletion
				os.Setenv("GATEWAY_ID", *result.ID)

				Expect(*result.Name).To(Equal(gatewayName))
				Expect(*result.BgpAsn).To(Equal(bgpAsn))
				Expect(*result.Global).To(Equal(true))
				Expect(*result.Metered).To(Equal(metered))
				Expect(*result.SpeedMbps).To(Equal(speedMbps))
				Expect(*result.LocationName).To(Equal(portLocationName))
				Expect(*result.LocationDisplayName).To(Equal(portLocationDisplayName))
				Expect(*result.BgpCerCidr).NotTo(BeEmpty())
				Expect(*result.BgpIbmCidr).NotTo(Equal(""))
				Expect(*result.BgpIbmAsn).NotTo(Equal(0))
				Expect(*result.BgpStatus).To(Equal("idle"))
				Expect(*result.CreatedAt).NotTo(Equal(""))
				Expect(*result.Crn).To(HavePrefix("crn:v1"))
				Expect(*result.OperationalStatus).To(Equal("create_pending"))
				Expect(*result.ResourceGroup.ID).NotTo(Equal(""))
				Expect(*result.Type).To(Equal("connect"))
				Expect(*result.Port.ID).To(Equal(portId))
				Expect(*result.ProviderApiManaged).To(Equal(false))
				Expect(*result.DefaultExportRouteFilter).To(Equal("deny"))
				Expect(*result.DefaultImportRouteFilter).To(Equal("deny"))
			})

			It("Successfully waits for connect gateway to be provisioned state", func() {
				shouldSkipTest()

				getGatewayOptions := service.NewGetGatewayOptions(os.Getenv("GATEWAY_ID"))

				// before a connect gateway can be deleted, it needs to have operational_status of provisioned.  We need to wait for
				// the new gateway to go to provisioned so we can delete it.
				timer := 0
				for {
					// Get the current status for the gateway
					result, detailedResponse, err := service.GetGateway(getGatewayOptions)
					Expect(err).To(BeNil())
					Expect(detailedResponse.StatusCode).To(Equal(200))

					Expect(*result.Name).To(Equal(gatewayName))
					Expect(*result.BgpAsn).To(Equal(bgpAsn))
					Expect(*result.Global).To(Equal(true))
					Expect(*result.Metered).To(Equal(metered))
					Expect(*result.SpeedMbps).To(Equal(speedMbps))
					Expect(*result.LocationName).To(Equal(portLocationName))
					Expect(*result.LocationDisplayName).To(Equal(portLocationDisplayName))
					Expect(*result.BgpCerCidr).NotTo(BeEmpty())
					Expect(*result.BgpIbmCidr).NotTo(Equal(""))
					Expect(*result.BgpIbmAsn).NotTo(Equal(0))
					Expect(*result.BgpStatus).To(Equal("idle"))
					Expect(*result.CreatedAt).NotTo(Equal(""))
					Expect(*result.Crn).To(HavePrefix("crn:v1"))
					Expect(*result.ResourceGroup.ID).NotTo(Equal(""))
					Expect(*result.Type).To(Equal("connect"))
					Expect(*result.Port.ID).To(Equal(portId))
					Expect(*result.ProviderApiManaged).To(Equal(false))
					Expect(*result.DefaultExportRouteFilter).To(Equal("deny"))
					Expect(*result.DefaultImportRouteFilter).To(Equal("deny"))

					// if operational status is "provisioned" then we are done
					if *result.OperationalStatus == "provisioned" {
						Expect(*result.OperationalStatus).To(Equal("provisioned"))
						break
					}

					// not provisioned yet, see if we have reached the timeout value.  If so, exit with failure
					if timer > 600 { // 5 min timer (24x5sec)
						Expect(*result.OperationalStatus).To(Equal("provisioned")) // timed out fail if status is not provisioned
						break
					} else {
						// Still exists, wait 5 sec
						time.Sleep(time.Duration(5) * time.Second)
						timer = timer + 1
					}
				}
			})
			It("Successfully deletes connect gateway", func() {
				shouldSkipTest()

				gatewayId := os.Getenv("GATEWAY_ID")
				deteleGatewayOptions := service.NewDeleteGatewayOptions(gatewayId)
				detailedResponse, err := service.DeleteGateway(deteleGatewayOptions)

				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(204))
			})
		})

		// Context("DirectLink MACsec Enabled Gateway", func() {
		// 	timestamp := time.Now().Unix()
		// 	gatewayName := "GO-INT-SDK-MACSEC" + strconv.FormatInt(timestamp, 10)
		// 	updatedGatewayName := "GO-INT-SDK-MACSEC-PATCH-" + strconv.FormatInt(timestamp, 10)
		// 	bgpAsn := int64(64999)
		// 	crossConnectRouter := "LAB-xcr01.dal09"
		// 	global := true
		// 	locationName := os.Getenv("LOCATION_NAME")
		// 	speedMbps := int64(1000)
		// 	metered := false
		// 	carrierName := "carrier1"
		// 	customerName := "customer1"
		// 	gatewayType := "dedicated"
		// 	macsecCak := os.Getenv("MACSEC_CAK")
		// 	macsecSakExpiryTime := int64(86400)
		// 	macsecWindowSize := int64(64)

		// 	It("Create a macsec enabled dedicated gateway", func() {
		// 		shouldSkipTest()

		// 		// Construct an instance of the GatewayMacsecCak model
		// 		gatewayMacsecCak := new(directlinkv1.GatewayMacsecConfigTemplatePrimaryCak)
		// 		gatewayMacsecCak.Crn = core.StringPtr(macsecCak)

		// 		// Construct an instance of the GatewayMacsecConfigTemplate model
		// 		gatewayMacsecConfigTemplate := new(directlinkv1.GatewayMacsecConfigTemplate)
		// 		gatewayMacsecConfigTemplate.Active = core.BoolPtr(true)
		// 		gatewayMacsecConfigTemplate.PrimaryCak = gatewayMacsecCak
		// 		gatewayMacsecConfigTemplate.WindowSize = core.Int64Ptr(macsecWindowSize)

		// 		gatewayTemplate := new(directlinkv1.GatewayTemplateGatewayTypeDedicatedTemplate)
		// 		gatewayTemplate.BgpAsn = core.Int64Ptr(bgpAsn)
		// 		gatewayTemplate.Global = core.BoolPtr(global)
		// 		gatewayTemplate.Metered = core.BoolPtr(metered)
		// 		gatewayTemplate.Name = core.StringPtr(gatewayName)
		// 		gatewayTemplate.SpeedMbps = core.Int64Ptr(int64(1000))
		// 		gatewayTemplate.Type = core.StringPtr(gatewayType)
		// 		gatewayTemplate.CarrierName = core.StringPtr(carrierName)
		// 		gatewayTemplate.CrossConnectRouter = core.StringPtr(crossConnectRouter)
		// 		gatewayTemplate.CustomerName = core.StringPtr(customerName)
		// 		gatewayTemplate.LocationName = core.StringPtr(locationName)
		// 		gatewayTemplate.MacsecConfig = gatewayMacsecConfigTemplate

		// 		createGatewayOptions := service.NewCreateGatewayOptions(gatewayTemplate)
		// 		result, detailedResponse, err := service.CreateGateway(createGatewayOptions)
		// 		Expect(err).To(BeNil())
		// 		Expect(detailedResponse.StatusCode).To(Equal(201))

		// 		os.Setenv("GATEWAY_ID", *result.ID)

		// 		Expect(*result.Name).To(Equal(gatewayName))
		// 		Expect(*result.BgpAsn).To(Equal(bgpAsn))
		// 		Expect(*result.Global).To(Equal(global))
		// 		Expect(*result.Metered).To(Equal(metered))
		// 		Expect(*result.SpeedMbps).To(Equal(speedMbps))
		// 		Expect(*result.Type).To(Equal(gatewayType))
		// 		Expect(*result.CrossConnectRouter).To(Equal(crossConnectRouter))
		// 		Expect(*result.LocationName).To(Equal(locationName))
		// 		Expect(*result.OperationalStatus).To(Equal("awaiting_loa"))
		// 		Expect(*result.MacsecConfig.Active).To(Equal(true))
		// 		Expect(*result.MacsecConfig.PrimaryCak.Crn).To(Equal(macsecCak))
		// 		Expect(*result.MacsecConfig.SakExpiryTime).To(Equal(macsecSakExpiryTime))
		// 		Expect(*result.MacsecConfig.WindowSize).To(Equal(macsecWindowSize))
		// 	})

		// 	It("Should successfully update the macsec enabled gateway", func() {
		// 		shouldSkipTest()

		// 		// Construct an instance of the GatewayMacsecCak model
		// 		gatewayMacsecCak := new(directlinkv1.GatewayMacsecConfigPatchTemplateFallbackCak)
		// 		gatewayMacsecCak.Crn = core.StringPtr(macsecCak)

		// 		// Construct an instance of the GatewayMacsecConfigTemplate model
		// 		gatewayMacsecConfigPatchTemplate := new(directlinkv1.GatewayMacsecConfigPatchTemplate)
		// 		gatewayMacsecConfigPatchTemplate.FallbackCak = gatewayMacsecCak

		// 		gatewayId := os.Getenv("GATEWAY_ID")
		// 		patchGatewayOptions := service.NewUpdateGatewayOptions(gatewayId)

		// 		result, detailedResponse, err := service.UpdateGateway(patchGatewayOptions.SetName(updatedGatewayName).SetMacsecConfig(gatewayMacsecConfigPatchTemplate))
		// 		Expect(err).To(BeNil())
		// 		Expect(detailedResponse.StatusCode).To(Equal(200))

		// 		Expect(*result.ID).To(Equal(gatewayId))
		// 		Expect(*result.Name).To(Equal(updatedGatewayName))
		// 		Expect(*result.MacsecConfig.Active).To(Equal(true))
		// 		Expect(*result.MacsecConfig.PrimaryCak.Crn).To(Equal(macsecCak))
		// 		Expect(*result.MacsecConfig.FallbackCak.Crn).To(Equal(macsecCak))
		// 		Expect(*result.MacsecConfig.SakExpiryTime).To(Equal(macsecSakExpiryTime))
		// 		Expect(*result.MacsecConfig.WindowSize).To(Equal(macsecWindowSize))

		// 	})

		// 	It("Successfully waits for macsec enabled gateway to be provisioned state", func() {
		// 		shouldSkipTest()

		// 		getGatewayOptions := service.NewGetGatewayOptions(os.Getenv("GATEWAY_ID"))

		// 		// before a dedicated gateway can be deleted, it needs to have operational_status of provisioned.  We need to wait for
		// 		// the new gateway to go to provisioned so we can delete it.
		// 		timer := 0
		// 		for {
		// 			// Get the current status for the gateway
		// 			result, detailedResponse, err := service.GetGateway(getGatewayOptions)
		// 			Expect(err).To(BeNil())
		// 			Expect(detailedResponse.StatusCode).To(Equal(200))

		// 			Expect(*result.Name).To(Equal(updatedGatewayName))
		// 			Expect(*result.BgpAsn).To(Equal(bgpAsn))
		// 			Expect(*result.Global).To(Equal(true))
		// 			Expect(*result.Metered).To(Equal(metered))
		// 			Expect(*result.SpeedMbps).To(Equal(speedMbps))
		// 			Expect(*result.BgpCerCidr).NotTo(BeEmpty())
		// 			Expect(*result.BgpIbmCidr).NotTo(Equal(""))
		// 			Expect(*result.BgpIbmAsn).NotTo(Equal(0))
		// 			Expect(*result.BgpStatus).To(Equal("idle"))
		// 			Expect(*result.CreatedAt).NotTo(Equal(""))
		// 			Expect(*result.Crn).To(HavePrefix("crn:v1"))
		// 			Expect(*result.ResourceGroup.ID).NotTo(Equal(""))
		// 			Expect(*result.Type).To(Equal("dedicated"))
		// 			Expect(*result.ProviderApiManaged).To(Equal(false))
		// 			Expect(*result.MacsecConfig.Active).To(Equal(true))
		// 			Expect(*result.MacsecConfig.PrimaryCak.Crn).To(Equal(macsecCak))
		// 			Expect(*result.MacsecConfig.FallbackCak.Crn).To(Equal(macsecCak))
		// 			Expect(*result.MacsecConfig.SakExpiryTime).To(Equal(macsecSakExpiryTime))
		// 			Expect(*result.MacsecConfig.WindowSize).To(Equal(macsecWindowSize))

		// 			// if operational status is "provisioned" then we are done
		// 			if *result.OperationalStatus == "provisioned" {
		// 				Expect(*result.OperationalStatus).To(Equal("provisioned"))
		// 				break
		// 			}

		// 			// not provisioned yet, see if we have reached the timeout value.  If so, exit with failure
		// 			if timer > 24 { // 2 min timer (24x5sec)
		// 				Expect(*result.OperationalStatus).To(Equal("provisioned")) // timed out fail if status is not provisioned
		// 				break
		// 			} else {
		// 				// Still exists, wait 5 sec
		// 				time.Sleep(time.Duration(5) * time.Second)
		// 				timer = timer + 1
		// 			}
		// 		}
		// 	})

		// 	It("Successfully deletes macsec enabled gateway gateway", func() {
		// 		shouldSkipTest()

		// 		gatewayId := os.Getenv("GATEWAY_ID")
		// 		deteleGatewayOptions := service.NewDeleteGatewayOptions(gatewayId)
		// 		detailedResponse, err := service.DeleteGateway(deteleGatewayOptions)

		// 		Expect(err).To(BeNil())
		// 		Expect(detailedResponse.StatusCode).To(Equal(204))
		// 	})
		// })
	})

	Describe("Offering Types", func() {

		Context("Locations", func() {
			It("should fetch the locations for the type dedicated", func() {
				shouldSkipTest()

				listOfferingTypeLocationsOptions := service.NewListOfferingTypeLocationsOptions("dedicated")
				result, detailedResponse, err := service.ListOfferingTypeLocations(listOfferingTypeLocationsOptions)
				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(200))
				Expect(len(result.Locations)).Should(BeNumerically(">", 0))
				os.Setenv("OT_DEDICATED_LOCATION_DISPLAY_NAME", *result.Locations[0].DisplayName)
				os.Setenv("OT_DEDICATED_LOCATION_NAME", *result.Locations[0].Name)

				Expect(*result.Locations[0].BillingLocation).NotTo(Equal(""))
				Expect(*result.Locations[0].BuildingColocationOwner).NotTo(Equal(""))
				Expect(*result.Locations[0].LocationType).NotTo(Equal(""))
				// Expect(*result.Locations[0].Market).NotTo(Equal(""))
				Expect(*result.Locations[0].MarketGeography).NotTo(Equal(""))
				Expect(*result.Locations[0].Mzr).NotTo(Equal(""))
				Expect(*result.Locations[0].OfferingType).To(Equal("dedicated"))
				Expect(*result.Locations[0].ProvisionEnabled).NotTo(BeNil())
				//Expect(*result.Locations[0].VpcRegion).NotTo(Equal(""))

			})

			It("should fetch the locations for the type connect", func() {
				shouldSkipTest()

				listOfferingTypeLocationsOptions := service.NewListOfferingTypeLocationsOptions("connect")

				result, detailedResponse, err := service.ListOfferingTypeLocations(listOfferingTypeLocationsOptions)
				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(200))
				Expect(len(result.Locations)).Should(BeNumerically(">", 0))
				os.Setenv("OT_CONNECT_LOCATION_DISPLAY_NAME", *result.Locations[0].DisplayName)
				os.Setenv("OT_CONNECT_LOCATION_NAME", *result.Locations[0].Name)

				Expect(*result.Locations[0].BillingLocation).NotTo(Equal(""))
				Expect(*result.Locations[0].LocationType).NotTo(Equal(""))
				// Expect(*result.Locations[0].Market).NotTo(Equal(""))
				Expect(*result.Locations[0].MarketGeography).NotTo(Equal(""))
				Expect(*result.Locations[0].Mzr).NotTo(Equal(""))
				Expect(*result.Locations[0].OfferingType).To(Equal("connect"))
				Expect(*result.Locations[0].ProvisionEnabled).NotTo(BeNil())
				// Expect(*result.Locations[0].VpcRegion).NotTo(Equal(""))
			})

			It("should return an error for invalid location type", func() {
				shouldSkipTest()

				listOfferingTypeLocationsOptions := service.NewListOfferingTypeLocationsOptions("RANDOM")

				result, detailedResponse, err := service.ListOfferingTypeLocations(listOfferingTypeLocationsOptions)
				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(Equal("offering_type_location: RANDOM"))
				Expect(detailedResponse.StatusCode).To(Equal(404))
				Expect(result).To(BeNil())
			})
		})

		Context("Cross Connect Routers", func() {
			/*
				 It("should list the location info for type dedicated and location short name", func() {
					 shouldSkipTest()

					 listOfferingTypeLocationCrossConnectRoutersOptions := service.NewListOfferingTypeLocationCrossConnectRoutersOptions("dedicated", os.Getenv("OT_DEDICATED_LOCATION_NAME"))

					 result, detailedResponse, err := service.ListOfferingTypeLocationCrossConnectRouters(listOfferingTypeLocationCrossConnectRoutersOptions)

					 Expect(err).To(BeNil())
					 Expect(detailedResponse.StatusCode).To(Equal(200))
					 Expect(len(result.CrossConnectRouters)).Should(BeNumerically(">", 0))

					 Expect(*result.CrossConnectRouters[0].RouterName).NotTo(Equal(""))
					 Expect(*result.CrossConnectRouters[0].TotalConnections).Should(BeNumerically(">=", 0))
				 })

				 It("should list the location info for type dedicated and location display name", func() {
					 shouldSkipTest()

					 listOfferingTypeLocationCrossConnectRoutersOptions := service.NewListOfferingTypeLocationCrossConnectRoutersOptions("dedicated", os.Getenv("OT_DEDICATED_LOCATION_DISPLAY_NAME"))

					 result, detailedResponse, err := service.ListOfferingTypeLocationCrossConnectRouters(listOfferingTypeLocationCrossConnectRoutersOptions)
					 Expect(err).To(BeNil())
					 Expect(detailedResponse.StatusCode).To(Equal(200))
					 Expect(len(result.CrossConnectRouters)).Should(BeNumerically(">", 0))

					 Expect(*result.CrossConnectRouters[0].RouterName).NotTo(Equal(""))
					 Expect(*result.CrossConnectRouters[0].TotalConnections).Should(BeNumerically(">=", 0))
				 })
			*/
			It("should return proper error when unsupported offering type CONNECT is provided", func() {
				shouldSkipTest()

				listOfferingTypeLocationCrossConnectRoutersOptions := service.NewListOfferingTypeLocationCrossConnectRoutersOptions("connect", os.Getenv("OT_CONNECT_LOCATION_NAME"))

				result, detailedResponse, err := service.ListOfferingTypeLocationCrossConnectRouters(listOfferingTypeLocationCrossConnectRoutersOptions)

				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(Equal("The supplied OfferingType is not supported for this call"))
				Expect(detailedResponse.StatusCode).To(Equal(400))
				Expect(result).To(BeNil())
			})

			It("should return proper error when incorrect offering type is provided", func() {
				shouldSkipTest()

				listOfferingTypeLocationCrossConnectRoutersOptions := service.NewListOfferingTypeLocationCrossConnectRoutersOptions("random", os.Getenv("OT_CONNECT_LOCATION_DISPLAY_NAME"))

				result, detailedResponse, err := service.ListOfferingTypeLocationCrossConnectRouters(listOfferingTypeLocationCrossConnectRoutersOptions)
				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(Equal("Invalid Direct Link Offering Type."))
				Expect(detailedResponse.StatusCode).To(Equal(400))
				Expect(result).To(BeNil())
			})

			It("should return proper error when incorrect location is provided", func() {
				shouldSkipTest()

				listOfferingTypeLocationCrossConnectRoutersOptions := service.NewListOfferingTypeLocationCrossConnectRoutersOptions("dedicated", "florida")

				result, detailedResponse, err := service.ListOfferingTypeLocationCrossConnectRouters(listOfferingTypeLocationCrossConnectRoutersOptions)
				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(Equal("Classic Location not found: florida"))
				Expect(detailedResponse.StatusCode).To(Equal(404))
				Expect(result).To(BeNil())
			})
		})

		Context("Offering Speeds", func() {
			It("should fetch the offering speeds for the type dedicated", func() {
				shouldSkipTest()

				listOfferingTypeSpeedsOptions := service.NewListOfferingTypeSpeedsOptions("dedicated")

				result, detailedResponse, err := service.ListOfferingTypeSpeeds(listOfferingTypeSpeedsOptions)
				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(200))
				Expect(len(result.Speeds)).Should(BeNumerically(">", 0))
			})

			It("should fetch the offering speeds for the type connect", func() {
				shouldSkipTest()

				listOfferingTypeSpeedsOptions := service.NewListOfferingTypeSpeedsOptions("connect")

				result, detailedResponse, err := service.ListOfferingTypeSpeeds(listOfferingTypeSpeedsOptions)
				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(200))
				Expect(len(result.Speeds)).Should(BeNumerically(">", 0))
			})

			It("should proper error for invalid offering type", func() {
				shouldSkipTest()

				listOfferingTypeSpeedsOptions := service.NewListOfferingTypeSpeedsOptions("random")

				result, detailedResponse, err := service.ListOfferingTypeSpeeds(listOfferingTypeSpeedsOptions)
				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(Equal("Cannot find OfferingType"))
				Expect(detailedResponse.StatusCode).To(Equal(404))
				Expect(result).To(BeNil())
			})
		})
	})

	Describe("Ports", func() {
		It("should fetch the ports", func() {
			shouldSkipTest()

			listPortsOptions := service.NewListPortsOptions()

			result, detailedResponse, err := service.ListPorts(listPortsOptions)
			Expect(err).To(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(200))
			Expect(len(result.Ports)).Should(BeNumerically(">", 0))

			Expect(*result.Ports[0].ID).NotTo(Equal(""))
			Expect(*result.Ports[0].DirectLinkCount).Should(BeNumerically(">=", 0))
			Expect(*result.Ports[0].Label).NotTo(Equal(""))
			Expect(*result.Ports[0].LocationDisplayName).NotTo(Equal(""))
			Expect(*result.Ports[0].LocationName).NotTo(Equal(""))
			Expect(*result.Ports[0].ProviderName).NotTo(Equal(""))
			Expect(len(result.Ports[0].SupportedLinkSpeeds)).Should(BeNumerically(">=", 0))

			port := getPortIdForConnect(result.Ports)
			os.Setenv("PORT_ID", *port.ID)
			os.Setenv("PORT_LOCATION_DISPLAY_NAME", *port.LocationDisplayName)
			os.Setenv("PORT_LOCATION_NAME", *port.LocationName)
			os.Setenv("PORT_LABEL", *port.Label)

		})

		It("should fetch the port by ID", func() {
			shouldSkipTest()

			portId := os.Getenv("PORT_ID")
			locationDisplayName := os.Getenv("PORT_LOCATION_DISPLAY_NAME")
			locationName := os.Getenv("PORT_LOCATION_NAME")
			label := os.Getenv("PORT_LABEL")
			getPortOptions := service.NewGetPortOptions(portId)

			result, detailedResponse, err := service.GetPort(getPortOptions)
			Expect(err).To(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(200))

			Expect(*result.ID).To(Equal(portId))
			Expect(*result.LocationDisplayName).To(Equal(locationDisplayName))
			Expect(*result.LocationName).To(Equal(locationName))
			Expect(*result.Label).To(Equal(label))
			Expect(*result.DirectLinkCount).Should(BeNumerically(">=", 0))
			Expect(*result.ProviderName).NotTo(Equal(""))
			Expect(len(result.SupportedLinkSpeeds)).Should(BeNumerically(">=", 0))
		})
	})

	Describe("Direct Link Virtual Connections", func() {
		timestamp := time.Now().Unix()
		gatewayName := "GO-INT-VC-SDK-" + strconv.FormatInt(timestamp, 10)
		bgpAsn := int64(64999)
		crossConnectRouter := "LAB-xcr01.dal09"
		global := true
		locationName := os.Getenv("LOCATION_NAME")
		speedMbps := int64(1000)
		metered := false
		carrierName := "carrier1"
		customerName := "customer1"
		gatewayType := "dedicated"

		Context("Create gateway", func() {

			gateway, _ := service.NewGatewayTemplateGatewayTypeDedicatedTemplate(bgpAsn, global, metered, gatewayName, speedMbps, gatewayType, carrierName, crossConnectRouter, customerName, locationName)

			createGatewayOptions := service.NewCreateGatewayOptions(gateway)

			It("Successfully created a gateway", func() {
				shouldSkipTest()

				result, detailedResponse, err := service.CreateGateway(createGatewayOptions)
				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(201))

				os.Setenv("GATEWAY_ID", *result.ID)

				Expect(*result.Name).To(Equal(gatewayName))
				Expect(*result.BgpAsn).To(Equal(bgpAsn))
				Expect(*result.Global).To(Equal(global))
				Expect(*result.Metered).To(Equal(metered))
				Expect(*result.SpeedMbps).To(Equal(speedMbps))
				Expect(*result.Type).To(Equal(gatewayType))
				Expect(*result.CrossConnectRouter).To(Equal(crossConnectRouter))
				Expect(*result.LocationName).To(Equal(locationName))
			})

			It("Successfully create a CLASSIC virtual connection", func() {
				shouldSkipTest()

				vcName := "GO-INT-CLASSIC-VC-SDK-" + strconv.FormatInt(timestamp, 10)
				createGatewayVCOptions := service.NewCreateGatewayVirtualConnectionOptions(os.Getenv("GATEWAY_ID"), vcName, directlinkv1.CreateGatewayVirtualConnectionOptions_Type_Classic)
				result, detailedResponse, err := service.CreateGatewayVirtualConnection(createGatewayVCOptions)
				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(201))

				os.Setenv("CLASSIC_VC_ID", *result.ID)

				Expect(*result.ID).NotTo(Equal(""))
				Expect(*result.Name).To(Equal(vcName))
				Expect(*result.Type).To(Equal(directlinkv1.CreateGatewayVirtualConnectionOptions_Type_Classic))
				Expect(*result.CreatedAt).NotTo(Equal(""))
				Expect(*result.Status).To(Equal("pending"))
			})

			It("Successfully get a CLASSIC virtual connection", func() {
				shouldSkipTest()

				vcName := "GO-INT-CLASSIC-VC-SDK-" + strconv.FormatInt(timestamp, 10)
				getGatewayVCOptions := service.NewGetGatewayVirtualConnectionOptions(os.Getenv("GATEWAY_ID"), os.Getenv("CLASSIC_VC_ID"))
				result, detailedResponse, err := service.GetGatewayVirtualConnection(getGatewayVCOptions)
				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(200))

				Expect(*result.ID).To(Equal(os.Getenv("CLASSIC_VC_ID")))
				Expect(*result.Name).To(Equal(vcName))
				Expect(*result.Type).To(Equal(directlinkv1.CreateGatewayVirtualConnectionOptions_Type_Classic))
				Expect(*result.CreatedAt).NotTo(Equal(""))
				Expect(*result.Status).To(Equal("pending"))
			})

			It("Successfully create a Gen 2 VPC virtual connection", func() {
				shouldSkipTest()

				vcName := "GO-INT-GEN2-VPC-VC-SDK-" + strconv.FormatInt(timestamp, 10)
				vpcCrn := os.Getenv("GEN2_VPC_CRN")
				createGatewayVCOptions := service.NewCreateGatewayVirtualConnectionOptions(os.Getenv("GATEWAY_ID"), vcName, directlinkv1.CreateGatewayVirtualConnectionOptions_Type_Vpc)
				createGatewayVCOptionsWithNetworkID := createGatewayVCOptions.SetNetworkID(vpcCrn)
				result, detailedResponse, err := service.CreateGatewayVirtualConnection(createGatewayVCOptionsWithNetworkID)
				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(201))

				// save the id so it can be deleted later
				os.Setenv("GEN2_VPC_VC_ID", *result.ID)

				Expect(*result.ID).NotTo(Equal(""))
				Expect(*result.Name).To(Equal(vcName))
				Expect(*result.CreatedAt).NotTo(Equal(""))
				Expect(*result.Status).To(Equal("pending"))
				Expect(*result.Type).To(Equal(directlinkv1.CreateGatewayVirtualConnectionOptions_Type_Vpc))
				Expect(*result.NetworkID).To(Equal(vpcCrn))
			})

			It("Successfully get a Gen 2 VPC virtual connection", func() {
				shouldSkipTest()

				getGatewayVCOptions := service.NewGetGatewayVirtualConnectionOptions(os.Getenv("GATEWAY_ID"), os.Getenv("GEN2_VPC_VC_ID"))
				result, detailedResponse, err := service.GetGatewayVirtualConnection(getGatewayVCOptions)
				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(200))

				Expect(*result.ID).To(Equal(os.Getenv("GEN2_VPC_VC_ID")))
				Expect(*result.Name).To(Equal("GO-INT-GEN2-VPC-VC-SDK-" + strconv.FormatInt(timestamp, 10)))
				Expect(*result.CreatedAt).NotTo(Equal(""))
				Expect(*result.Status).To(Equal("pending"))
				Expect(*result.Type).To(Equal(directlinkv1.CreateGatewayVirtualConnectionOptions_Type_Vpc))
				Expect(*result.NetworkID).To(Equal(os.Getenv("GEN2_VPC_CRN")))
			})

			It("Successfully list the virtual connections for a gateway", func() {
				shouldSkipTest()

				listVcOptions := service.NewListGatewayVirtualConnectionsOptions(os.Getenv("GATEWAY_ID"))
				result, detailedResponse, err := service.ListGatewayVirtualConnections(listVcOptions)
				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(200))

				vcs := result.VirtualConnections
				// two VCs were created for the GW, so we should expect 2
				Expect(len(vcs)).Should(BeNumerically("==", 2))

				for _, vc := range vcs {
					if *vc.ID == os.Getenv("GEN2_VPC_VC_ID") {
						Expect(*vc.Name).To(Equal("GO-INT-GEN2-VPC-VC-SDK-" + strconv.FormatInt(timestamp, 10)))
						Expect(*vc.CreatedAt).NotTo(Equal(""))
						Expect(*vc.Status).To(Equal("pending"))
						Expect(*vc.Type).To(Equal(directlinkv1.CreateGatewayVirtualConnectionOptions_Type_Vpc))
						Expect(*vc.NetworkID).To(Equal(os.Getenv("GEN2_VPC_CRN")))
					} else {
						Expect(*vc.ID).To(Equal(os.Getenv("CLASSIC_VC_ID")))
						Expect(*vc.Name).To(Equal("GO-INT-CLASSIC-VC-SDK-" + strconv.FormatInt(timestamp, 10)))
						Expect(*vc.Type).To(Equal(directlinkv1.CreateGatewayVirtualConnectionOptions_Type_Classic))
						Expect(*vc.CreatedAt).NotTo(Equal(""))
						Expect(*vc.Status).To(Equal("pending"))
					}
				}
			})

			It("Successfully Update a virtual connection name", func() {
				shouldSkipTest()

				gatewayId := os.Getenv("GATEWAY_ID")
				vcId := os.Getenv("GEN2_VPC_VC_ID")
				vcName := "GO-INT-GEN2-VPC-VC-PATCH-SDK-" + strconv.FormatInt(timestamp, 10)
				patchGatewayOptions := service.NewUpdateGatewayVirtualConnectionOptions(gatewayId, vcId)
				patchGatewayOptions = patchGatewayOptions.SetName(vcName)

				result, detailedResponse, err := service.UpdateGatewayVirtualConnection(patchGatewayOptions)
				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(200))

				Expect(*result.ID).To(Equal(vcId))
				Expect(*result.Name).To(Equal(vcName))
				Expect(*result.CreatedAt).NotTo(Equal(""))
				Expect(*result.Status).To(Equal("pending"))
				Expect(*result.Type).To(Equal(directlinkv1.CreateGatewayVirtualConnectionOptions_Type_Vpc))
				Expect(*result.NetworkID).To(Equal(os.Getenv("GEN2_VPC_CRN")))
			})

			It("Fail to Update a virtual connection status", func() {
				shouldSkipTest()

				gatewayId := os.Getenv("GATEWAY_ID")
				vcId := os.Getenv("GEN2_VPC_VC_ID")
				patchGatewayOptions := service.NewUpdateGatewayVirtualConnectionOptions(gatewayId, vcId)
				patchGatewayOptions = patchGatewayOptions.SetStatus(directlinkv1.UpdateGatewayVirtualConnectionOptions_Status_Rejected)

				result, detailedResponse, err := service.UpdateGatewayVirtualConnection(patchGatewayOptions)

				// GW owner is not allowed to change the status, but the test calls the API with the status parameter to valid it is allowed.
				Expect(result).To(BeNil())
				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(Equal("gateway owner can't patch vc status."))
				Expect(detailedResponse.StatusCode).To(Equal(400))
			})

			It("Successfully delete a CLASSIC virtual connection for a gateway", func() {
				shouldSkipTest()

				gatewayId := os.Getenv("GATEWAY_ID")
				vcId := os.Getenv("CLASSIC_VC_ID")
				deleteClassicVCOptions := service.NewDeleteGatewayVirtualConnectionOptions(gatewayId, vcId)

				detailedResponse, err := service.DeleteGatewayVirtualConnection(deleteClassicVCOptions)
				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(204))
			})

			It("Successfully waits for CLASSIC virtual connection to report as deleted", func() {
				shouldSkipTest()

				getGatewayVCOptions := service.NewGetGatewayVirtualConnectionOptions(os.Getenv("GATEWAY_ID"), os.Getenv("CLASSIC_VC_ID"))

				// VC delete might not be instantaneous.  Poll the VC looking for a not found.  Fail after 2 min
				timer := 0
				for {
					// Get the current rc for the VC
					_, detailedResponse, _ := service.GetGatewayVirtualConnection(getGatewayVCOptions)

					// if 404 then we are done
					if detailedResponse.StatusCode == 404 {
						Expect(detailedResponse.StatusCode).To(Equal(404)) // response is 404, exit success
						break
					}

					// other than 404, see if we have reached the timeout value.  If so, exit with failure
					if timer > 600 { // 2 min timer (24x5sec)
						Expect(detailedResponse.StatusCode).To(Equal(404)) // timed out fail if code is not 404
						break
					} else {
						// Still exists, wait 5 sec
						time.Sleep(time.Duration(5) * time.Second)
						timer = timer + 1
					}
				}
			})

			It("Successfully deletes GEN 2 VPC virtual connection for a gateway", func() {
				shouldSkipTest()

				gatewayId := os.Getenv("GATEWAY_ID")
				vcId := os.Getenv("GEN2_VPC_VC_ID")
				deleteVpcVcOptions := service.NewDeleteGatewayVirtualConnectionOptions(gatewayId, vcId)

				detailedResponse, err := service.DeleteGatewayVirtualConnection(deleteVpcVcOptions)
				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(204))
			})

			It("Successfully waits for GEN 2 VPC virtual connection to report as deleted", func() {
				shouldSkipTest()

				getGatewayVCOptions := service.NewGetGatewayVirtualConnectionOptions(os.Getenv("GATEWAY_ID"), os.Getenv("GEN2_VPC_VC_ID"))

				// VC delete might not be instantaneous.  Poll the VC looking for a not found.  Fail after 2 min
				timer := 0
				for {
					// Get the current rc for the VC
					_, detailedResponse, _ := service.GetGatewayVirtualConnection(getGatewayVCOptions)

					// if 404 then we are done
					if detailedResponse.StatusCode == 404 {
						Expect(detailedResponse.StatusCode).To(Equal(404)) // response is 404, exit success
						break
					}

					// other than 404, see if we have reached the timeout value.  If so, exit with failure
					if timer > 600 { // 2 min timer (24x5 sec)
						Expect(detailedResponse.StatusCode).To(Equal(404)) // timed out fail if code is not 404
						break
					} else {
						// Still exists, wait 5 sec
						time.Sleep(time.Duration(5) * time.Second)
						timer = timer + 1
					}
				}
			})

			It("Successfully deletes a gateway", func() {
				shouldSkipTest()

				gatewayId := os.Getenv("GATEWAY_ID")
				deteleGatewayOptions := service.NewDeleteGatewayOptions(gatewayId)

				detailedResponse, err := service.DeleteGateway(deteleGatewayOptions)
				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(204))
			})
		})
	})

	Describe("LOA and Completion Notice", func() {
		timestamp := time.Now().Unix()
		gatewayName := "GO-INT-LOA-SDK-" + strconv.FormatInt(timestamp, 10)
		bgpAsn := int64(64999)
		crossConnectRouter := "LAB-xcr01.dal09"
		global := true
		locationName := os.Getenv("LOCATION_NAME")
		speedMbps := int64(1000)
		metered := false
		carrierName := "carrier1"
		customerName := "customer1"
		gatewayType := "dedicated"

		// notes about LOA and CN testing.  When a GW is created, a github issue is also created by dl-rest.  The issue is used for managing the LOA and CN.  In normal operation,
		// an LOA is added to the issue via manual GH interaction.  After that occurs and the GH label changed, then CN upload is allowed.  Since we do not have the ability to
		// do the manual steps for integration testing, the test will only do the following
		//	- Issue GET LOA for a gateway.  It will expect a 404 error since no one has added the LOA to the GH issue
		//  - PUT a completion notice to the gw.  It will fail with a 412 error because the GH issue and GW status are in the wrong state due to no manual interaction
		//  - GET CN for a gw.  It will expect a 404 since the CN could not be uploaded
		//
		Context("Create gateway", func() {
			It("Successfully created a gateway", func() {
				shouldSkipTest()

				gateway, _ := service.NewGatewayTemplateGatewayTypeDedicatedTemplate(bgpAsn, global, metered, gatewayName, speedMbps, gatewayType, carrierName, crossConnectRouter, customerName, locationName)
				createGatewayOptions := service.NewCreateGatewayOptions(gateway)

				result, detailedResponse, err := service.CreateGateway(createGatewayOptions)
				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(201))

				os.Setenv("GATEWAY_ID", *result.ID)
			})

			It("Successfully call loa", func() {
				shouldSkipTest()

				listLOAOptions := service.NewListGatewayLetterOfAuthorizationOptions(os.Getenv("GATEWAY_ID"))
				result, detailedResponse, err := service.ListGatewayLetterOfAuthorization(listLOAOptions)
				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(Equal("Please check whether the resource you are requesting exists."))
				Expect(detailedResponse.StatusCode).To(Equal(404))
				Expect(result).To(BeNil())
			})

			It("Successfully call PUT completion notice", func() {
				shouldSkipTest()

				buffer, err := os.ReadFile("completion_notice.pdf")
				Expect(err).To(BeNil())
				r := io.NopCloser(bytes.NewReader(buffer))

				createCNOptions := service.NewCreateGatewayCompletionNoticeOptions(os.Getenv("GATEWAY_ID"))
				createCNOptions.SetUpload(r)

				detailedResponse, err := service.CreateGatewayCompletionNotice(createCNOptions)

				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(Equal("Invalid gateway status to upload completion notice."))
				Expect(detailedResponse.StatusCode).To(Equal(412))
			})

			It("Successfully call completion notice", func() {
				shouldSkipTest()

				listCNOptions := service.NewListGatewayCompletionNoticeOptions(os.Getenv("GATEWAY_ID"))
				result, detailedResponse, err := service.ListGatewayCompletionNotice(listCNOptions)

				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(Equal("Please check whether the resource you are requesting exists."))
				Expect(detailedResponse.StatusCode).To(Equal(404))
				Expect(result).To(BeNil())
			})

			It("Successfully deletes a gateway", func() {
				shouldSkipTest()

				gatewayId := os.Getenv("GATEWAY_ID")
				deteleGatewayOptions := service.NewDeleteGatewayOptions(gatewayId)

				detailedResponse, err := service.DeleteGateway(deteleGatewayOptions)
				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(204))
			})
		})
	})
	/*
		 Describe("BGP MD5", func() {
			 timestamp := time.Now().Unix()
			 gatewayName := "GO-INT-MD5-SDK-" + strconv.FormatInt(timestamp, 10)
			 bgpAsn := int64(64999)
			 crossConnectRouter := "LAB-xcr01.dal09"
			 global := true
			 locationName := os.Getenv("LOCATION_NAME")
			 speedMbps := int64(1000)
			 metered := false
			 carrierName := "carrier1"
			 customerName := "customer1"
			 gatewayType := "dedicated"
			 authCrn := os.Getenv("AUTHENTICATION_KEY")

			 Context("Create a Gateway with Authentication Key", func() {
				 It("should successfully create a gateway", func() {
					 shouldSkipTest()

					 // gateway, _ := service.NewGatewayTemplateGatewayTypeDedicatedTemplate(bgpAsn, global, metered, gatewayName, speedMbps, gatewayType, carrierName, crossConnectRouter, customerName, locationName)
					 authenticationKey, _ := service.NewGatewayTemplateAuthenticationKey(authCrn)

					 gatewayTemplateModel := new(directlinkv1.GatewayTemplateGatewayTypeDedicatedTemplate)
					 gatewayTemplateModel.AuthenticationKey = authenticationKey
					 gatewayTemplateModel.BgpAsn = core.Int64Ptr(int64(64999))
					 gatewayTemplateModel.Global = core.BoolPtr(true)
					 gatewayTemplateModel.Metered = core.BoolPtr(false)
					 gatewayTemplateModel.Name = core.StringPtr(gatewayName)
					 gatewayTemplateModel.SpeedMbps = core.Int64Ptr(int64(1000))
					 gatewayTemplateModel.Type = core.StringPtr(gatewayType)
					 gatewayTemplateModel.CarrierName = core.StringPtr(carrierName)
					 gatewayTemplateModel.CrossConnectRouter = core.StringPtr(crossConnectRouter)
					 gatewayTemplateModel.CustomerName = core.StringPtr(customerName)
					 gatewayTemplateModel.LocationName = core.StringPtr(locationName)

					 createGatewayOptions := service.NewCreateGatewayOptions(gatewayTemplateModel)

					 result, resp, err := service.CreateGateway(createGatewayOptions)

					 Expect(err).To(BeNil())
					 Expect(resp.StatusCode).To(Equal(201))

					 os.Setenv("GATEWAY_ID", *result.ID)

					 Expect(*result.Name).To(Equal(gatewayName))
					 Expect(*result.AuthenticationKey.Crn).To(Equal(authCrn))
					 Expect(*result.BgpAsn).To(Equal(bgpAsn))
					 Expect(*result.Global).To(Equal(global))
					 Expect(*result.Metered).To(Equal(metered))
					 Expect(*result.SpeedMbps).To(Equal(speedMbps))
					 Expect(*result.Type).To(Equal(gatewayType))
					 Expect(*result.CrossConnectRouter).To(Equal(crossConnectRouter))
					 Expect(*result.LocationName).To(Equal(locationName))
					 Expect(*result.LocationDisplayName).NotTo(Equal(""))
					 Expect(*result.BgpCerCidr).NotTo(BeEmpty())
					 Expect(*result.BgpIbmCidr).NotTo(Equal(""))
					 Expect(*result.BgpIbmAsn).NotTo(Equal(""))
					 Expect(*result.BgpStatus).To(Equal("idle"))
					 Expect(*result.CreatedAt).NotTo(Equal(""))
					 Expect(*result.Crn).To(HavePrefix("crn:v1"))
					 Expect(*result.LinkStatus).To(Equal("down"))
					 Expect(*result.OperationalStatus).To(Equal("awaiting_loa"))
					 Expect(*result.ResourceGroup.ID).NotTo(Equal(""))

				 })
			 })

			 Context("Update the Authentication key for the gateway", func() {
				 It("should successfully clear the auth key", func() {
					 shouldSkipTest()
					 authKey, _ := service.NewGatewayPatchTemplateAuthenticationKey("")
					 gatewayId := os.Getenv("GATEWAY_ID")

					 updateGatewayOptions := service.NewUpdateGatewayOptions(gatewayId).SetAuthenticationKey(authKey)
					 res, resp, err := service.UpdateGateway(updateGatewayOptions)
					 Expect(err).To(BeNil())
					 Expect(resp.StatusCode).To(Equal(200))

					 Expect(*res.ID).To(Equal(gatewayId))
					 Expect(res.AuthenticationKey).To(BeNil())
					 Expect(*res.Name).To(Equal(gatewayName))
				 })
			 })

			 Context("Delete a gateway", func() {
				 It("Successfully deletes a gateway", func() {
					 shouldSkipTest()

					 gatewayId := os.Getenv("GATEWAY_ID")
					 deteleGatewayOptions := service.NewDeleteGatewayOptions(gatewayId)

					 detailedResponse, err := service.DeleteGateway(deteleGatewayOptions)
					 Expect(err).To(BeNil())
					 Expect(detailedResponse.StatusCode).To(Equal(204))
				 })
			 })
		 })
	*/
	Describe("DLAAS", func() {

		Describe("Create/Verify/update a connect gateway", func() {
			timestamp := time.Now().Unix()
			gatewayName := "GO-INT-SDK-Connect-DLAAS-" + strconv.FormatInt(timestamp, 10)
			bgpAsn := int64(64999)
			global := true
			speedMbps := int64(1000)
			metered := false
			// to create a connect gateway, we need to have a port.  List the ports and save the id of the 1st one found
			portId := ""
			portLocationDisplayName := ""
			portLocationName := ""

			It("List ports and save the id of the first port", func() {
				shouldSkipTest()

				listPortsOptions := service.NewListPortsOptions()
				result, detailedResponse, err := service.ListPorts(listPortsOptions)
				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(200))
				port := getPortIdForConnect(result.Ports)
				portId = *port.ID
				portLocationDisplayName = *port.LocationDisplayName
				portLocationName = *port.LocationName
			})

			It("create connect gateway with connection_mode as transit", func() {
				shouldSkipTest()

				portIdentity, _ := service.NewGatewayPortIdentity(portId)
				gateway, _ := service.NewGatewayTemplateGatewayTypeConnectTemplate(bgpAsn, global, metered, gatewayName, speedMbps, "connect", portIdentity)
				gateway.ConnectionMode = core.StringPtr("transit")
				createGatewayOptions := service.NewCreateGatewayOptions(gateway)
				result, detailedResponse, err := service.CreateGateway(createGatewayOptions)

				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(201))

				// Save the gateway id for deletion
				os.Setenv("GATEWAY_ID", *result.ID)

				Expect(*result.Name).To(Equal(gatewayName))
				Expect(*result.BgpAsn).To(Equal(bgpAsn))
				Expect(*result.Global).To(Equal(true))
				Expect(*result.Metered).To(Equal(metered))
				Expect(*result.SpeedMbps).To(Equal(speedMbps))
				Expect(*result.LocationName).To(Equal(portLocationName))
				Expect(*result.LocationDisplayName).To(Equal(portLocationDisplayName))
				Expect(*result.BgpCerCidr).NotTo(BeEmpty())
				Expect(*result.BgpIbmCidr).NotTo(Equal(""))
				Expect(*result.BgpIbmAsn).NotTo(Equal(0))
				Expect(*result.BgpStatus).To(Equal("idle"))
				Expect(*result.CreatedAt).NotTo(Equal(""))
				Expect(*result.Crn).To(HavePrefix("crn:v1"))
				Expect(*result.OperationalStatus).To(Equal("create_pending"))
				Expect(*result.ResourceGroup.ID).NotTo(Equal(""))
				Expect(*result.Type).To(Equal("connect"))
				Expect(*result.Port.ID).To(Equal(portId))
				Expect(*result.ProviderApiManaged).To(Equal(false))
				Expect(*result.ConnectionMode).To(Equal("transit"))
			})

			It("Successfully waits for gateway to be provisioned state", func() {
				shouldSkipTest()

				getGatewayOptions := service.NewGetGatewayOptions(os.Getenv("GATEWAY_ID"))

				// before connection_mode can be updated on a gateway, it needs to have operational_status of provisioned.  We need to wait for
				// the new gateway to go to provisioned so we can delete it.
				timer := 0
				for {
					// Get the current status for the gateway
					result, detailedResponse, err := service.GetGateway(getGatewayOptions)
					Expect(err).To(BeNil())
					Expect(detailedResponse.StatusCode).To(Equal(200))

					Expect(*result.Name).To(Equal(gatewayName))
					Expect(*result.BgpAsn).To(Equal(bgpAsn))
					Expect(*result.Global).To(Equal(true))
					Expect(*result.Metered).To(Equal(metered))
					Expect(*result.SpeedMbps).To(Equal(speedMbps))
					Expect(*result.LocationName).To(Equal(portLocationName))
					Expect(*result.LocationDisplayName).To(Equal(portLocationDisplayName))
					Expect(*result.BgpCerCidr).NotTo(BeEmpty())
					Expect(*result.BgpIbmCidr).NotTo(Equal(""))
					Expect(*result.BgpIbmAsn).NotTo(Equal(0))
					// Expect(*result.BgpStatus).To(Equal("idle"))
					Expect(*result.CreatedAt).NotTo(Equal(""))
					Expect(*result.Crn).To(HavePrefix("crn:v1"))
					Expect(*result.ResourceGroup.ID).NotTo(Equal(""))
					Expect(*result.Type).To(Equal("connect"))
					Expect(*result.Port.ID).To(Equal(portId))
					Expect(*result.ProviderApiManaged).To(Equal(false))

					// if operational status is "provisioned" then we are done
					if *result.OperationalStatus == "provisioned" {
						Expect(*result.OperationalStatus).To(Equal("provisioned"))
						break
					}

					// not provisioned yet, see if we have reached the timeout value.  If so, exit with failure
					if timer > 600 { // 2 min timer (24x5sec)
						Expect(*result.OperationalStatus).To(Equal("provisioned")) // timed out fail if status is not provisioned
						break
					} else {
						// Still exists, wait 5 sec
						time.Sleep(time.Duration(5) * time.Second)
						timer = timer + 1
					}
				}
			})

			It("should successfully switch the connection mode to direct", func() {
				shouldSkipTest()
				gatewayId := os.Getenv("GATEWAY_ID")

				updateGatewayOptions := service.NewUpdateGatewayOptions(gatewayId).SetConnectionMode("direct")
				res, resp, err := service.UpdateGateway(updateGatewayOptions)
				Expect(err).To(BeNil())
				Expect(resp.StatusCode).To(Equal(200))

				Expect(*res.ID).To(Equal(gatewayId))
				Expect(*res.ConnectionMode).To(Equal("direct"))
				Expect(*res.Name).To(Equal(gatewayName))
			})

			It("Successfully waits for gateway to be provisioned state", func() {
				shouldSkipTest()

				getGatewayOptions := service.NewGetGatewayOptions(os.Getenv("GATEWAY_ID"))

				// before connection_mode can be updated on a gateway, it needs to have operational_status of provisioned.  We need to wait for
				// the new gateway to go to provisioned so we can delete it.
				timer := 0
				for {
					// Get the current status for the gateway
					result, detailedResponse, err := service.GetGateway(getGatewayOptions)
					Expect(err).To(BeNil())
					Expect(detailedResponse.StatusCode).To(Equal(200))

					Expect(*result.Name).To(Equal(gatewayName))
					Expect(*result.BgpAsn).To(Equal(bgpAsn))
					Expect(*result.Global).To(Equal(true))
					Expect(*result.Metered).To(Equal(metered))
					Expect(*result.SpeedMbps).To(Equal(speedMbps))
					Expect(*result.LocationName).To(Equal(portLocationName))
					Expect(*result.LocationDisplayName).To(Equal(portLocationDisplayName))
					Expect(*result.BgpCerCidr).NotTo(BeEmpty())
					Expect(*result.BgpIbmCidr).NotTo(Equal(""))
					Expect(*result.BgpIbmAsn).NotTo(Equal(0))
					// Expect(*result.BgpStatus).To(Equal("idle"))
					Expect(*result.CreatedAt).NotTo(Equal(""))
					Expect(*result.Crn).To(HavePrefix("crn:v1"))
					Expect(*result.ResourceGroup.ID).NotTo(Equal(""))
					Expect(*result.Type).To(Equal("connect"))
					Expect(*result.Port.ID).To(Equal(portId))
					Expect(*result.ProviderApiManaged).To(Equal(false))

					// if operational status is "provisioned" then we are done
					if *result.OperationalStatus == "provisioned" {
						Expect(*result.OperationalStatus).To(Equal("provisioned"))
						break
					}

					// not provisioned yet, see if we have reached the timeout value.  If so, exit with failure
					if timer > 600 { // 2 min timer (24x5sec)
						Expect(*result.OperationalStatus).To(Equal("provisioned")) // timed out fail if status is not provisioned
						break
					} else {
						// Still exists, wait 5 sec
						time.Sleep(time.Duration(5) * time.Second)
						timer = timer + 1
					}
				}
			})

			It("Successfully deletes connect gateway", func() {
				shouldSkipTest()

				gatewayId := os.Getenv("GATEWAY_ID")
				deteleGatewayOptions := service.NewDeleteGatewayOptions(gatewayId)
				detailedResponse, err := service.DeleteGateway(deteleGatewayOptions)

				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(204))
			})
		})

		Describe("Create/verify/update a dedicated gateway", func() {
			timestamp := time.Now().Unix()
			gatewayName := "GO-INT-SDK-Dedicated-DLAAS-" + strconv.FormatInt(timestamp, 10)
			bgpAsn := int64(64999)
			crossConnectRouter := "LAB-xcr01.dal09"
			global := true
			locationName := os.Getenv("LOCATION_NAME")
			speedMbps := int64(1000)
			metered := false
			carrierName := "carrier1"
			customerName := "customer1"
			gatewayType := "dedicated"
			connectionMode := "direct"

			It("should successfully create a dedicated gateway with connection mode as direct", func() {
				shouldSkipTest()

				gatewayTemplateModel := new(directlinkv1.GatewayTemplateGatewayTypeDedicatedTemplate)
				gatewayTemplateModel.BgpAsn = core.Int64Ptr(int64(64999))
				gatewayTemplateModel.Global = core.BoolPtr(true)
				gatewayTemplateModel.Metered = core.BoolPtr(false)
				gatewayTemplateModel.Name = core.StringPtr(gatewayName)
				gatewayTemplateModel.SpeedMbps = core.Int64Ptr(int64(1000))
				gatewayTemplateModel.Type = core.StringPtr(gatewayType)
				gatewayTemplateModel.CarrierName = core.StringPtr(carrierName)
				gatewayTemplateModel.CrossConnectRouter = core.StringPtr(crossConnectRouter)
				gatewayTemplateModel.CustomerName = core.StringPtr(customerName)
				gatewayTemplateModel.LocationName = core.StringPtr(locationName)
				gatewayTemplateModel.ConnectionMode = core.StringPtr(connectionMode)

				createGatewayOptions := service.NewCreateGatewayOptions(gatewayTemplateModel)

				result, resp, err := service.CreateGateway(createGatewayOptions)

				Expect(err).To(BeNil())
				Expect(resp.StatusCode).To(Equal(201))

				os.Setenv("GATEWAY_ID", *result.ID)

				Expect(*result.Name).To(Equal(gatewayName))
				Expect(*result.BgpAsn).To(Equal(bgpAsn))
				Expect(*result.Global).To(Equal(global))
				Expect(*result.Metered).To(Equal(metered))
				Expect(*result.SpeedMbps).To(Equal(speedMbps))
				Expect(*result.Type).To(Equal(gatewayType))
				Expect(*result.CrossConnectRouter).To(Equal(crossConnectRouter))
				Expect(*result.LocationName).To(Equal(locationName))
				Expect(*result.LocationDisplayName).NotTo(Equal(""))
				Expect(*result.BgpCerCidr).NotTo(BeEmpty())
				Expect(*result.BgpIbmCidr).NotTo(Equal(""))
				Expect(*result.BgpIbmAsn).NotTo(Equal(""))
				Expect(*result.BgpStatus).To(Equal("idle"))
				Expect(*result.CreatedAt).NotTo(Equal(""))
				Expect(*result.Crn).To(HavePrefix("crn:v1"))
				Expect(*result.LinkStatus).To(Equal("down"))
				Expect(*result.OperationalStatus).To(Equal("awaiting_loa"))
				Expect(*result.ResourceGroup.ID).NotTo(Equal(""))
				Expect(*result.ConnectionMode).To(Equal("direct"))
			})

			It("should successfully switch the connection mode to transit", func() {
				shouldSkipTest()
				gatewayId := os.Getenv("GATEWAY_ID")

				updateGatewayOptions := service.NewUpdateGatewayOptions(gatewayId).SetConnectionMode("transit")
				res, resp, err := service.UpdateGateway(updateGatewayOptions)
				Expect(err).To(BeNil())
				Expect(resp.StatusCode).To(Equal(200))

				Expect(*res.ID).To(Equal(gatewayId))
				Expect(*res.ConnectionMode).To(Equal("transit"))
				Expect(*res.Name).To(Equal(gatewayName))
			})

			It("Successfully deletes a gateway", func() {
				shouldSkipTest()

				gatewayId := os.Getenv("GATEWAY_ID")
				deteleGatewayOptions := service.NewDeleteGatewayOptions(gatewayId)

				detailedResponse, err := service.DeleteGateway(deteleGatewayOptions)
				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(204))
			})
		})
	})

	Describe("BGP IP Update", func() {
		timestamp := time.Now().Unix()
		gatewayName := "GO-INT-BGP-IP-SDK-" + strconv.FormatInt(timestamp, 10)
		bgpAsn := int64(64999)
		crossConnectRouter := "LAB-xcr01.dal09"
		global := true
		locationName := os.Getenv("LOCATION_NAME")
		speedMbps := int64(1000)
		metered := false
		carrierName := "carrier1"
		customerName := "customer1"
		gatewayType := "dedicated"

		Context("Create a Gateway", func() {
			It("should successfully create a gateway", func() {
				shouldSkipTest()

				gatewayTemplateModel := new(directlinkv1.GatewayTemplateGatewayTypeDedicatedTemplate)
				gatewayTemplateModel.BgpAsn = core.Int64Ptr(int64(64999))
				gatewayTemplateModel.Global = core.BoolPtr(true)
				gatewayTemplateModel.Metered = core.BoolPtr(false)
				gatewayTemplateModel.Name = core.StringPtr(gatewayName)
				gatewayTemplateModel.SpeedMbps = core.Int64Ptr(int64(1000))
				gatewayTemplateModel.Type = core.StringPtr(gatewayType)
				gatewayTemplateModel.CarrierName = core.StringPtr(carrierName)
				gatewayTemplateModel.CrossConnectRouter = core.StringPtr(crossConnectRouter)
				gatewayTemplateModel.CustomerName = core.StringPtr(customerName)
				gatewayTemplateModel.LocationName = core.StringPtr(locationName)

				createGatewayOptions := service.NewCreateGatewayOptions(gatewayTemplateModel)

				result, resp, err := service.CreateGateway(createGatewayOptions)

				Expect(err).To(BeNil())
				Expect(resp.StatusCode).To(Equal(201))

				os.Setenv("GATEWAY_ID", *result.ID)
				os.Setenv("BGP_IP_CER", *result.BgpCerCidr)
				os.Setenv("BGP_IP_IBM", *result.BgpIbmCidr)

				Expect(*result.Name).To(Equal(gatewayName))
				Expect(*result.BgpAsn).To(Equal(bgpAsn))
				Expect(*result.Global).To(Equal(global))
				Expect(*result.Metered).To(Equal(metered))
				Expect(*result.SpeedMbps).To(Equal(speedMbps))
				Expect(*result.Type).To(Equal(gatewayType))
				Expect(*result.CrossConnectRouter).To(Equal(crossConnectRouter))
				Expect(*result.LocationName).To(Equal(locationName))
				Expect(*result.LocationDisplayName).NotTo(Equal(""))
				Expect(*result.BgpCerCidr).NotTo(BeEmpty())
				Expect(*result.BgpIbmCidr).NotTo(Equal(""))
				Expect(*result.BgpIbmAsn).NotTo(Equal(""))
				Expect(*result.BgpStatus).To(Equal("idle"))
				Expect(*result.CreatedAt).NotTo(Equal(""))
				Expect(*result.Crn).To(HavePrefix("crn:v1"))
				Expect(*result.LinkStatus).To(Equal("down"))
				Expect(*result.OperationalStatus).To(Equal("awaiting_loa"))
				Expect(*result.ResourceGroup.ID).NotTo(Equal(""))

			})
		})

		Context("Update the BGP ASN for the gateway", func() {
			It("should successfully update the bgp asn", func() {
				shouldSkipTest()
				gatewayId := os.Getenv("GATEWAY_ID")

				bgpAsn := int64(63999)
				updateGatewayOptions := service.NewUpdateGatewayOptions(gatewayId).SetBgpAsn(bgpAsn)
				res, resp, err := service.UpdateGateway(updateGatewayOptions)
				Expect(err).To(BeNil())
				Expect(resp.StatusCode).To(Equal(200))

				Expect(*res.ID).To(Equal(gatewayId))
				Expect(*res.BgpAsn).To(Equal(bgpAsn))
				Expect(*res.Name).To(Equal(gatewayName))
			})
		})

		Context("Update the BGP IP for the gateway", func() {
			It("should either successfully update the BGP IP CER and IBM CIDR", func() {
				shouldSkipTest()
				gatewayId := os.Getenv("GATEWAY_ID")
				updateGatewayOptions := service.NewUpdateGatewayOptions(gatewayId).SetBgpCerCidr("172.17.252.2/29").SetBgpIbmCidr("172.17.252.1/29")
				res, resp, err := service.UpdateGateway(updateGatewayOptions)
				if err != nil {
					Expect(err.Error()).To(Equal("Please make sure localIP and remoteIP are not in use"))
				} else {
					Expect(err).To(BeNil())
					Expect(resp.StatusCode).To(Equal(200))

					Expect(*res.ID).To(Equal(gatewayId))
					Expect(*res.Name).To(Equal(gatewayName))
				}

			})
		})

		Context("Delete a gateway", func() {
			It("Successfully deletes a gateway", func() {
				shouldSkipTest()

				gatewayId := os.Getenv("GATEWAY_ID")
				deteleGatewayOptions := service.NewDeleteGatewayOptions(gatewayId)

				detailedResponse, err := service.DeleteGateway(deteleGatewayOptions)
				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(204))
			})
		})
	})

	Describe("BFD Config", func() {
		timestamp := time.Now().Unix()
		gatewayName := "GO-INT-BFD-SDK-" + strconv.FormatInt(timestamp, 10)
		bgpAsn := int64(64999)
		crossConnectRouter := "LAB-xcr01.dal09"
		locationName := os.Getenv("LOCATION_NAME")
		speedMbps := int64(1000)
		carrierName := "carrier1"
		customerName := "customer1"
		gatewayType := "dedicated"
		bfdInterval := int64(1000)
		bfdMultiplier := int64(10)

		Context("Create a Gateway", func() {
			It("should successfully create a gateway", func() {
				shouldSkipTest()

				// Create a template for BFD Config
				bfdTemplate := new(directlinkv1.GatewayBfdConfigTemplate)
				bfdTemplate.Interval = &bfdInterval
				bfdTemplate.Multiplier = &bfdMultiplier

				// Create a template for Gateway model
				gatewayTemplateModel := new(directlinkv1.GatewayTemplateGatewayTypeDedicatedTemplate)
				gatewayTemplateModel.BgpAsn = core.Int64Ptr(bgpAsn)
				gatewayTemplateModel.Global = core.BoolPtr(true)
				gatewayTemplateModel.Metered = core.BoolPtr(false)
				gatewayTemplateModel.Name = core.StringPtr(gatewayName)
				gatewayTemplateModel.SpeedMbps = core.Int64Ptr(speedMbps)
				gatewayTemplateModel.Type = core.StringPtr(gatewayType)
				gatewayTemplateModel.CarrierName = core.StringPtr(carrierName)
				gatewayTemplateModel.CrossConnectRouter = core.StringPtr(crossConnectRouter)
				gatewayTemplateModel.CustomerName = core.StringPtr(customerName)
				gatewayTemplateModel.LocationName = core.StringPtr(locationName)
				gatewayTemplateModel.BfdConfig = bfdTemplate

				createGatewayOptions := service.NewCreateGatewayOptions(gatewayTemplateModel)

				result, resp, err := service.CreateGateway(createGatewayOptions)

				Expect(err).To(BeNil())
				Expect(resp.StatusCode).To(Equal(201))

				os.Setenv("GATEWAY_ID", *result.ID)

				Expect(*result.Name).To(Equal(gatewayName))
				Expect(*result.Type).To(Equal(gatewayType))
				Expect(*result.CrossConnectRouter).To(Equal(crossConnectRouter))
				Expect(*result.LocationName).To(Equal(locationName))
				Expect(*result.LocationDisplayName).NotTo(Equal(""))
				Expect(*result.BgpStatus).To(Equal("idle"))
				Expect(*result.OperationalStatus).To(Equal("awaiting_loa"))
				Expect(result.BfdConfig).NotTo(BeNil())
				Expect(result.BfdConfig.BfdStatus).NotTo(BeNil())
				Expect(*result.BfdConfig.Interval).To(Equal(bfdInterval))
				Expect(*result.BfdConfig.Multiplier).To(Equal(bfdMultiplier))
			})
		})

		Context("Update the BFD Config for the gateway", func() {
			It("should successfully update the bfd config", func() {
				shouldSkipTest()
				gatewayId := os.Getenv("GATEWAY_ID")

				updatedBfdInterval := int64(400)
				updatedBfdMultiplier := int64(200)

				// Create a template for BFD Config
				bfdPatchTemplate := new(directlinkv1.GatewayBfdPatchTemplate)
				bfdPatchTemplate.Interval = &updatedBfdInterval
				bfdPatchTemplate.Multiplier = &updatedBfdMultiplier

				updateGatewayOptions := service.NewUpdateGatewayOptions(gatewayId).SetBfdConfig(bfdPatchTemplate)
				res, resp, err := service.UpdateGateway(updateGatewayOptions)
				Expect(err).To(BeNil())
				Expect(resp.StatusCode).To(Equal(200))

				Expect(*res.ID).To(Equal(gatewayId))
				Expect(*res.Name).To(Equal(gatewayName))
				Expect(res.BfdConfig).NotTo(BeNil())
				Expect(res.BfdConfig.BfdStatus).NotTo(BeNil())
				Expect(*res.BfdConfig.Interval).To(Equal(updatedBfdInterval))
				Expect(*res.BfdConfig.Multiplier).To(Equal(updatedBfdMultiplier))
			})
		})

		Context("Delete a gateway", func() {
			It("Successfully deletes a gateway", func() {
				shouldSkipTest()

				gatewayId := os.Getenv("GATEWAY_ID")
				deteleGatewayOptions := service.NewDeleteGatewayOptions(gatewayId)

				detailedResponse, err := service.DeleteGateway(deteleGatewayOptions)
				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(204))
			})
		})
	})
})
