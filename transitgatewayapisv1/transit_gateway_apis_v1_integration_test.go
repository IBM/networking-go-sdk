/**
 * (C) Copyright IBM Corp. 2020,2022.
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

package transitgatewayapisv1_test

/*

How to run this test:

go test -v ./transitgatewayapisv1

*/

import (
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/networking-go-sdk/transitgatewayapisv1"
	"github.com/joho/godotenv"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var configLoaded = false

func shouldSkipTest() {
	Skip("test exceeds the 10 mins travis limit")
	if !configLoaded {
		Skip("External configuration is not available, skipping...")
	}
}

var _ = Describe(`TransitGatewayApisV1`, func() {
	err := godotenv.Load("../transit.env")
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
	options := &transitgatewayapisv1.TransitGatewayApisV1Options{
		ServiceName:   "TransitGatewayApisV1_Mocking",
		Authenticator: authenticator,
		URL:           serviceURL,
		Version:       &version,
	}

	service, err := transitgatewayapisv1.NewTransitGatewayApisV1UsingExternalConfig(options)
	It(`Successfully created TransitGatewayApisV1 service instance`, func() {
		shouldSkipTest()
		Expect(err).To(BeNil())
	})

	// Test fixed variables:
	timestamp := time.Now().Unix()
	gatewayName := "SDK-GO-TEST-Gateway_" + strconv.FormatInt(timestamp, 10)
	connectionName := "SDK-GO-TEST-Connection_" + strconv.FormatInt(timestamp, 10)

	///////////////////////////////////////////////////////////////////////////////
	//                              Pre-Test Cleanup                             //
	///////////////////////////////////////////////////////////////////////////////

	Describe(`PreTest cleanup`, func() {
		Context(`Successfully clean test environment`, func() {
			header := map[string]string{
				"Content-type": "application/json",
			}
			listTransitGatewaysOptions := service.NewListTransitGatewaysOptions().
				SetHeaders(header)

			It(`Checking gateways`, func() {
				// shouldSkipTest()

				result, detailedResponse, err := service.ListTransitGateways(listTransitGatewaysOptions)
				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(200))

				for _, gw := range result.TransitGateways {
					if strings.HasPrefix(*gw.Name, "SDK-GO-TEST") {
						gatewayID := *gw.ID
						listTransitGatewayConnectionsOptions := service.NewListTransitGatewayConnectionsOptions(gatewayID).
							SetTransitGatewayID(gatewayID).
							SetHeaders(header)

						result, detailedResponse, err := service.ListTransitGatewayConnections(listTransitGatewayConnectionsOptions)
						Expect(err).To(BeNil())
						Expect(detailedResponse.StatusCode).To(Equal(200))

						if len(result.Connections) > 0 {
							connIDs := []string{}
							for _, conn := range result.Connections {
								connID := *conn.ID
								if !strings.Contains(*conn.Status, "delet") {
									// Delete GRE Connections first.
									if *conn.NetworkType == "gre_tunnel" {
										deleteTransitGatewayConnectionOptions := service.NewDeleteTransitGatewayConnectionOptions(gatewayID, connID)
										detailedResponse, err := service.DeleteTransitGatewayConnection(deleteTransitGatewayConnectionOptions)
										Expect(err).To(BeNil())
										Expect(detailedResponse.StatusCode).To(Equal(204))
										isResourceAvailable(service, gatewayID, connID, "")
										deleteCheckTest(service, gatewayID, connID, "", "")
									} else {
										connIDs = append(connIDs, connID)
									}
								}
							}
							// Delete Connections from other types.
							for _, curConn := range connIDs {
								deleteTransitGatewayConnectionOptions := service.NewDeleteTransitGatewayConnectionOptions(gatewayID, curConn)
								detailedResponse, err := service.DeleteTransitGatewayConnection(deleteTransitGatewayConnectionOptions)
								Expect(err).To(BeNil())
								Expect(detailedResponse.StatusCode).To(Equal(204))
								deleteCheckTest(service, gatewayID, curConn, "", "")
							}
						}
						// Remove empty gateways
						if !strings.Contains(*gw.Status, "delet") {
							deleteTransitGatewayOptions := service.NewDeleteTransitGatewayOptions(gatewayID)
							detailedResponse, err = service.DeleteTransitGateway(deleteTransitGatewayOptions)
							Expect(err).To(BeNil())
							Expect(detailedResponse.StatusCode).To(Equal(204))
						}
					}
				}
			})
		})
	})

	///////////////////////////////////////////////////////////////////////////////
	//                          Transit Locations Tests                          //
	///////////////////////////////////////////////////////////////////////////////

	Describe(`ListGatewayLocations(listGatewayLocationsOptions *ListGatewayLocationsOptions)`, func() {
		Context(`Success: LIST Transit Locations`, func() {
			It(`Successfully list all transit locations`, func() {
				shouldSkipTest()

				header := map[string]string{
					"Content-type": "application/json",
				}
				listGatewayLocationsOptions := service.NewListGatewayLocationsOptions().
					SetHeaders(header)

				result, detailedResponse, err := service.ListGatewayLocations(listGatewayLocationsOptions)
				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(200))
				Expect(len(result.Locations)).Should(BeNumerically(">", 0))

				firstResource := result.Locations[0]
				Expect(*firstResource.Name).ToNot(BeNil())
				Expect(*firstResource.BillingLocation).ToNot(BeNil())
				Expect(*firstResource.Type).ToNot(BeNil())
			})
		})
	})

	Describe(`GetGatewayLocation(getGatewayLocationOptions *GetGatewayLocationOptions)`, func() {
		Context(`Success: GET Transit location by ID`, func() {
			It(`Successfully get location by instanceID`, func() {
				shouldSkipTest()

				instanceID := "us-south"
				getGatewayLocationOptions := service.NewGetGatewayLocationOptions(instanceID)

				result, detailedResponse, err := service.GetGatewayLocation(getGatewayLocationOptions)
				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(200))
				Expect(*result.Name).To(Equal(instanceID))
				Expect(*result.BillingLocation).ToNot(BeNil())
				Expect(*result.Type).ToNot(BeNil())
				Expect(len(result.LocalConnectionLocations)).Should(BeNumerically(">", 0))
			})
		})

		Context(`Failure: GET location by instanceID`, func() {
			badinstanceID := "abc123"
			getGatewayLocationOptions := &transitgatewayapisv1.GetGatewayLocationOptions{}
			getGatewayLocationOptions.SetName(badinstanceID)
			header := map[string]string{
				"Content-type": "application/json",
			}
			getGatewayLocationOptions.SetHeaders(header)

			It(`Failed to get location by instanceID`, func() {
				shouldSkipTest()

				result, detailedResponse, err := service.GetGatewayLocation(getGatewayLocationOptions)
				Expect(result).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(404))
				Expect(err).Should(HaveOccurred())
			})
		})
	})

	///////////////////////////////////////////////////////////////////////////////
	//                           Transit Gateway Tests                           //
	///////////////////////////////////////////////////////////////////////////////

	Describe(`CreateTransitGateway(createTransitGatewayOptions *CreateTransitGatewayOptions)`, func() {
		Context(`Success: POST Transit Gateway`, func() {
			It(`Successfully created new gateway`, func() {
				shouldSkipTest()

				header := map[string]string{
					"Content-type": "application/json",
				}
				location := os.Getenv("LOCATION")
				createTransitGatewayOptions := service.NewCreateTransitGatewayOptions(location, gatewayName).
					SetHeaders(header)

				result, detailedResponse, err := service.CreateTransitGateway(createTransitGatewayOptions)
				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(201))

				Expect(*result.ID).NotTo(Equal(""))
				Expect(*result.Name).To(Equal(gatewayName))
				Expect(*result.Crn).NotTo(Equal(""))
				Expect(*result.Global).NotTo(BeNil())
				Expect(*result.CreatedAt).NotTo(Equal(""))
				Expect(*result.UpdatedAt).NotTo(Equal(""))
				Expect(*result.Status).To(Equal("pending"))
				Expect(*result.ResourceGroup.ID).NotTo(Equal(""))
				Expect(*result.Location).To(Equal(os.Getenv("LOCATION")))

				os.Setenv("GATEWAY_INSTANCE_ID", *result.ID)
			})
			It("Successfully waits for gateway to report as available", func() {
				shouldSkipTest()

				gatewayID := os.Getenv("GATEWAY_INSTANCE_ID")
				isResourceAvailable(service, gatewayID, "", "")
			})
		})
		Context(`Failure: POST Transit gateway`, func() {
			createTransitGatewayOptions := &transitgatewayapisv1.CreateTransitGatewayOptions{}
			createTransitGatewayOptions.SetName("testString")
			createTransitGatewayOptions.SetLocation("testString")
			header := map[string]string{
				"Content-type": "application/json",
			}
			createTransitGatewayOptions.SetHeaders(header)

			It(`Fail to create new gateway`, func() {
				shouldSkipTest()

				result, detailedResponse, err := service.CreateTransitGateway(createTransitGatewayOptions)
				Expect(result).To(BeNil())
				Expect(detailedResponse.StatusCode).ToNot(Equal(200))
				Expect(err).Should(HaveOccurred())
			})
		})
	})

	Describe(`GetTransitGateway(getTransitGatewayOptions *GetTransitGatewayOptions)`, func() {
		Context(`Success: GET Transit Gateway by instanceID`, func() {
			It(`Successfully get gateway by instanceID`, func() {
				shouldSkipTest()

				gateway_id := os.Getenv("GATEWAY_INSTANCE_ID")
				getTransitGatewayOptions := service.NewGetTransitGatewayOptions(gateway_id)

				result, detailedResponse, err := service.GetTransitGateway(getTransitGatewayOptions)
				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(200))

				Expect(*result.Name).To(Equal(gatewayName))
				Expect(*result.Crn).NotTo(Equal(""))
				Expect(*result.Global).NotTo(BeNil())
				Expect(*result.ID).To(Equal(gateway_id))
				Expect(*result.CreatedAt).NotTo(Equal(""))
				Expect(*result.UpdatedAt).NotTo(Equal(""))
				Expect(*result.Status).To(Equal("available"))
				Expect(*result.ResourceGroup.ID).NotTo(Equal(""))
				Expect(*result.Location).To(Equal(os.Getenv("LOCATION")))
			})
		})

		Context(`Failure: GET gateway by instanceID`, func() {
			badinstanceID := "abc123"
			getTransitGatewayOptions := &transitgatewayapisv1.GetTransitGatewayOptions{}
			getTransitGatewayOptions.SetID(badinstanceID)
			header := map[string]string{
				"Content-type": "application/json",
			}
			getTransitGatewayOptions.SetHeaders(header)

			It(`Failed to get gateway by instanceID`, func() {
				shouldSkipTest()

				result, detailedResponse, err := service.GetTransitGateway(getTransitGatewayOptions)
				Expect(result).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(404))
				Expect(err).Should(HaveOccurred())
			})
		})
	})

	Describe(`UpdateTransitGateway(updateTransitGatewayOptions *UpdateTransitGatewayOptions)`, func() {
		Context(`Success: UPDATE Transit Gateway by instanceID`, func() {
			It(`Successfully update gateway by instanceID`, func() {
				shouldSkipTest()

				gateway_id := os.Getenv("GATEWAY_INSTANCE_ID")
				updateName := "UPDATED-" + gatewayName
				updateTransitGatewayOptions := service.NewUpdateTransitGatewayOptions(gateway_id).
					SetName(updateName)

				result, detailedResponse, err := service.UpdateTransitGateway(updateTransitGatewayOptions)
				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(200))

				Expect(*result.Crn).NotTo(Equal(""))
				Expect(*result.Global).NotTo(BeNil())
				Expect(*result.ID).To(Equal(gateway_id))
				Expect(*result.CreatedAt).NotTo(Equal(""))
				Expect(*result.UpdatedAt).NotTo(Equal(""))
				Expect(*result.Name).To(Equal(updateName))
				Expect(*result.Status).To(Equal("available"))
				Expect(*result.ResourceGroup.ID).NotTo(Equal(""))
				Expect(*result.Location).To(Equal(os.Getenv("LOCATION")))
				gatewayName = *result.Name // Update gateway name global variable.
			})
		})

		Context(`Failure: UPDATE gateway by instanceID`, func() {
			badinstanceID := "abc123"
			instanceName := "UPDATED-" + strconv.FormatInt(timestamp, 10)
			updateTransitGatewayOptions := &transitgatewayapisv1.UpdateTransitGatewayOptions{}
			updateTransitGatewayOptions.SetID(badinstanceID)
			updateTransitGatewayOptions.SetName(instanceName)
			header := map[string]string{
				"Content-type": "application/json",
			}
			updateTransitGatewayOptions.SetHeaders(header)

			It(`Failed to update gateway by instanceID`, func() {
				shouldSkipTest()

				result, detailedResponse, err := service.UpdateTransitGateway(updateTransitGatewayOptions)
				Expect(result).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(404))
				Expect(err).Should(HaveOccurred())
			})
		})
	})

	Describe(`ListTransitGateways(listTransitGatewaysOptions *ListTransitGatewaysOptions)`, func() {
		Context(`Success: LIST Transit Gateways`, func() {
			It(`Successfully list all transit gateways`, func() {
				shouldSkipTest()

				header := map[string]string{
					"Content-type": "application/json",
				}
				listTransitGatewaysOptions := service.NewListTransitGatewaysOptions().
					SetHeaders(header)

				result, detailedResponse, err := service.ListTransitGateways(listTransitGatewaysOptions)
				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(200))
				Expect(len(result.TransitGateways)).Should(BeNumerically(">", 0))

				found := false
				for _, gw := range result.TransitGateways {
					if *gw.ID == os.Getenv("GATEWAY_INSTANCE_ID") {
						Expect(*gw.Name).To(Equal(gatewayName))
						Expect(*gw.Crn).NotTo(Equal(""))
						Expect(*gw.Global).NotTo(BeNil())
						Expect(*gw.CreatedAt).NotTo(Equal(""))
						Expect(*gw.UpdatedAt).NotTo(Equal(""))
						Expect(*gw.Status).To(Equal("available"))
						Expect(*gw.ResourceGroup.ID).NotTo(Equal(""))
						Expect(*gw.Location).To(Equal(os.Getenv("LOCATION")))

						found = true
						break
					}
				}
				Expect(found).To(Equal(true))
			})
		})
	})

	///////////////////////////////////////////////////////////////////////////////
	//                    Transit Gateway Connections Tests                      //
	///////////////////////////////////////////////////////////////////////////////

	Describe(`CreateTransitGatewayConnection(createTransitGatewayConnectionOptions *CreateTransitGatewayConnectionOptions)`, func() {
		Context(`Success: POST Transit Gateway CLASSIC Connection`, func() {
			header := map[string]string{
				"Content-type": "application/json",
			}
			It(`Successfully create new CLASSIC Connection`, func() {
				shouldSkipTest()

				network_type := "classic"
				gatewayID := os.Getenv("GATEWAY_INSTANCE_ID")

				createTransitGatewayConnectionOptions := service.NewCreateTransitGatewayConnectionOptions(
					gatewayID,
					network_type).
					SetHeaders(header).
					SetName("CLASSIC-" + connectionName)

				result, detailedResponse, err := service.CreateTransitGatewayConnection(createTransitGatewayConnectionOptions)
				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(201))

				os.Setenv("CLASSIC_CONN_INSTANCE_ID", *result.ID)
				os.Setenv("CLASSIC_CONN_INSTANCE_NAME", *result.Name)

				Expect(*result.ID).NotTo(Equal(""))
				Expect(*result.CreatedAt).NotTo(Equal(""))
				Expect(*result.UpdatedAt).NotTo(Equal(""))
				Expect(*result.Status).To(Equal("pending"))
				Expect(*result.NetworkType).To(Equal(network_type))
				Expect(*result.Name).To(Equal(os.Getenv("CLASSIC_CONN_INSTANCE_NAME")))
			})

			It("Successfully waits for CLASSIC connection to report as attached", func() {
				shouldSkipTest()

				gatewayID := os.Getenv("GATEWAY_INSTANCE_ID")
				instanceID := os.Getenv("CLASSIC_CONN_INSTANCE_ID")
				isResourceAvailable(service, gatewayID, instanceID, "")
			})
		})

		Context(`Success: POST Transit Gateway VPC Connection`, func() {
			header := map[string]string{
				"Content-type": "application/json",
			}
			It(`Successfully create new VPC Connection`, func() {
				shouldSkipTest()

				network_type := "vpc"
				crn := os.Getenv("VPC_CRN")
				gatewayID := os.Getenv("GATEWAY_INSTANCE_ID")

				createTransitGatewayConnectionOptions := service.NewCreateTransitGatewayConnectionOptions(
					gatewayID,
					network_type).
					SetHeaders(header).
					SetName("VPC-" + connectionName).
					SetNetworkID(crn)

				result, detailedResponse, err := service.CreateTransitGatewayConnection(createTransitGatewayConnectionOptions)
				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(201))

				os.Setenv("VPC_CONN_INSTANCE_ID", *result.ID)
				os.Setenv("VPC_CONN_INSTANCE_NAME", *result.Name)

				Expect(*result.ID).NotTo(Equal(""))
				Expect(*result.NetworkID).To(Equal(crn))
				Expect(*result.CreatedAt).NotTo(Equal(""))
				Expect(*result.UpdatedAt).NotTo(Equal(""))
				Expect(*result.Status).To(Equal("pending"))
				Expect(*result.NetworkType).To(Equal(network_type))
				Expect(*result.Name).To(Equal(os.Getenv("VPC_CONN_INSTANCE_NAME")))
			})

			It("Successfully waits for VPC connection to report as attached", func() {
				shouldSkipTest()

				gatewayID := os.Getenv("GATEWAY_INSTANCE_ID")
				instanceID := os.Getenv("VPC_CONN_INSTANCE_ID")
				isResourceAvailable(service, gatewayID, instanceID, "")
			})
		})

		Context(`Success: POST Transit Gateway DL Connection`, func() {
			header := map[string]string{
				"Content-type": "application/json",
			}
			It(`Successfully create new DL Connection`, func() {
				shouldSkipTest()

				crn := os.Getenv("DL_CRN")
				network_type := "directlink"
				gatewayID := os.Getenv("GATEWAY_INSTANCE_ID")

				createTransitGatewayConnectionOptions := service.NewCreateTransitGatewayConnectionOptions(
					gatewayID,
					network_type).
					SetHeaders(header).
					SetName("DL-" + connectionName).
					SetNetworkID(crn)

				result, detailedResponse, err := service.CreateTransitGatewayConnection(createTransitGatewayConnectionOptions)
				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(201))

				os.Setenv("DL_CONN_INSTANCE_ID", *result.ID)
				os.Setenv("DL_CONN_INSTANCE_NAME", *result.Name)

				Expect(*result.ID).NotTo(Equal(""))
				Expect(*result.NetworkID).To(Equal(crn))
				Expect(*result.CreatedAt).NotTo(Equal(""))
				Expect(*result.UpdatedAt).NotTo(Equal(""))
				Expect(*result.Status).To(Equal("pending"))
				Expect(*result.NetworkType).To(Equal(network_type))
				Expect(*result.Name).To(Equal(os.Getenv("DL_CONN_INSTANCE_NAME")))

			})

			It("Successfully waits for DL connection to report as attached", func() {
				shouldSkipTest()

				gatewayID := os.Getenv("GATEWAY_INSTANCE_ID")
				instanceID := os.Getenv("DL_CONN_INSTANCE_ID")
				isResourceAvailable(service, gatewayID, instanceID, "")
			})
		})

		Context(`Success: POST Transit Gateway GRE Connection`, func() {
			header := map[string]string{
				"Content-type": "application/json",
			}
			It(`Successfully create new GRE Connection`, func() {
				shouldSkipTest()

				zoneStr := "us-south-1"
				network_type := "gre_tunnel"
				gatewayID := os.Getenv("GATEWAY_INSTANCE_ID")
				zone := &transitgatewayapisv1.ZoneIdentity{Name: &zoneStr}

				createTransitGatewayConnectionOptions := service.NewCreateTransitGatewayConnectionOptions(
					gatewayID,
					network_type).
					SetZone(zone).
					SetHeaders(header).
					SetName("GRE-" + connectionName).
					SetLocalTunnelIp("192.168.101.1").
					SetLocalGatewayIp("192.168.100.1").
					SetRemoteTunnelIp("192.168.101.2").
					SetRemoteGatewayIp("10.242.63.12").
					SetBaseConnectionID(os.Getenv("CLASSIC_CONN_INSTANCE_ID"))

				result, detailedResponse, err := service.CreateTransitGatewayConnection(createTransitGatewayConnectionOptions)
				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(201))

				os.Setenv("GRE_CONN_INSTANCE_ID", *result.ID)
				os.Setenv("GRE_CONN_INSTANCE_NAME", *result.Name)

				Expect(*result.ID).NotTo(Equal(""))
				Expect(*result.CreatedAt).NotTo(Equal(""))
				Expect(*result.UpdatedAt).NotTo(Equal(""))
				Expect(*result.Status).To(Equal("pending"))
				Expect(*result.NetworkType).To(Equal(network_type))
				Expect(*result.Name).To(Equal(os.Getenv("GRE_CONN_INSTANCE_NAME")))
				Expect(*result.BaseConnectionID).To(Equal(os.Getenv("CLASSIC_CONN_INSTANCE_ID")))
			})

			It("Successfully waits for GRE connection to report as attached", func() {
				shouldSkipTest()

				gatewayID := os.Getenv("GATEWAY_INSTANCE_ID")
				instanceID := os.Getenv("GRE_CONN_INSTANCE_ID")
				isResourceAvailable(service, gatewayID, instanceID, "")
			})
		})

		Context(`Failure: POST gateway resource`, func() {
			createTransitGatewayConnectionOptions := &transitgatewayapisv1.CreateTransitGatewayConnectionOptions{}
			createTransitGatewayConnectionOptions.SetName("testString")
			createTransitGatewayConnectionOptions.SetTransitGatewayID("testString")
			createTransitGatewayConnectionOptions.SetNetworkType("testString")
			header := map[string]string{
				"Content-type": "application/json",
			}
			createTransitGatewayConnectionOptions.SetHeaders(header)

			It(`Fail to create new gateway resource`, func() {
				shouldSkipTest()

				result, detailedResponse, err := service.CreateTransitGatewayConnection(createTransitGatewayConnectionOptions)
				Expect(result).To(BeNil())
				Expect(detailedResponse.StatusCode).ToNot(Equal(200))
				Expect(err).Should(HaveOccurred())
			})
		})
	})

	Describe(`GetTransitGatewayConnection(getTransitGatewayConnectionOptions *GetTransitGatewayConnectionOptions)`, func() {
		Context(`Success: GET Transit Gateway VPC Connection`, func() {
			It(`Successfully get VPC Connection`, func() {
				shouldSkipTest()

				gatewayID := os.Getenv("GATEWAY_INSTANCE_ID")
				instanceID := os.Getenv("VPC_CONN_INSTANCE_ID")
				getTransitGatewayConnectionOptions := service.NewGetTransitGatewayConnectionOptions(gatewayID, instanceID)

				result, detailedResponse, err := service.GetTransitGatewayConnection(getTransitGatewayConnectionOptions)
				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(200))

				Expect(*result.ID).To(Equal(instanceID))
				Expect(*result.NetworkType).To(Equal("vpc"))
				Expect(*result.CreatedAt).NotTo(Equal(""))
				Expect(*result.UpdatedAt).NotTo(Equal(""))
				Expect(*result.Status).To(Equal("attached"))
				Expect(*result.NetworkID).To(Equal(os.Getenv("VPC_CRN")))
				Expect(*result.Name).To(Equal(os.Getenv("VPC_CONN_INSTANCE_NAME")))
			})
		})

		Context(`Success: GET Transit Gateway CLASSIC Connection`, func() {
			It(`Successfully get CLASSIC Connection`, func() {
				shouldSkipTest()

				gatewayID := os.Getenv("GATEWAY_INSTANCE_ID")
				instanceID := os.Getenv("CLASSIC_CONN_INSTANCE_ID")
				getTransitGatewayConnectionOptions := service.NewGetTransitGatewayConnectionOptions(gatewayID, instanceID)

				result, detailedResponse, err := service.GetTransitGatewayConnection(getTransitGatewayConnectionOptions)
				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(200))
				Expect(*result.ID).To(Equal(instanceID))
				Expect(*result.CreatedAt).NotTo(Equal(""))
				Expect(*result.UpdatedAt).NotTo(Equal(""))
				Expect(*result.Status).To(Equal("attached"))
				Expect(*result.NetworkType).To(Equal("classic"))
				Expect(*result.Name).To(Equal(os.Getenv("CLASSIC_CONN_INSTANCE_NAME")))
			})
		})

		Context(`Success: GET Transit Gateway DL Connection`, func() {
			It(`Successfully get DL Connection`, func() {
				shouldSkipTest()

				gatewayID := os.Getenv("GATEWAY_INSTANCE_ID")
				instanceID := os.Getenv("DL_CONN_INSTANCE_ID")
				getTransitGatewayConnectionOptions := service.NewGetTransitGatewayConnectionOptions(gatewayID, instanceID)

				result, detailedResponse, err := service.GetTransitGatewayConnection(getTransitGatewayConnectionOptions)
				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(200))

				Expect(*result.ID).To(Equal(instanceID))
				Expect(*result.CreatedAt).NotTo(Equal(""))
				Expect(*result.UpdatedAt).NotTo(Equal(""))
				Expect(*result.Status).To(Equal("attached"))
				Expect(*result.NetworkType).To(Equal("directlink"))
				Expect(*result.NetworkID).To(Equal(os.Getenv("DL_CRN")))
				Expect(*result.Name).To(Equal(os.Getenv("DL_CONN_INSTANCE_NAME")))
			})
		})

		Context(`Success: GET Transit Gateway GRE Connection`, func() {
			It(`Successfully get GRE Connection`, func() {
				shouldSkipTest()

				gatewayID := os.Getenv("GATEWAY_INSTANCE_ID")
				instanceID := os.Getenv("GRE_CONN_INSTANCE_ID")
				getTransitGatewayConnectionOptions := service.NewGetTransitGatewayConnectionOptions(gatewayID, instanceID)

				result, detailedResponse, err := service.GetTransitGatewayConnection(getTransitGatewayConnectionOptions)
				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(200))

				Expect(*result.ID).To(Equal(instanceID))
				Expect(*result.CreatedAt).NotTo(Equal(""))
				Expect(*result.UpdatedAt).NotTo(Equal(""))
				Expect(*result.Status).To(Equal("attached"))
				Expect(*result.NetworkType).To(Equal("gre_tunnel"))
				Expect(*result.Name).To(Equal(os.Getenv("GRE_CONN_INSTANCE_NAME")))
				Expect(*result.BaseConnectionID).To(Equal(os.Getenv("CLASSIC_CONN_INSTANCE_ID")))
			})
		})

		Context(`Failure: GET connection by instanceID`, func() {
			badinstanceID := "abc123"
			getTransitGatewayConnectionOptions := &transitgatewayapisv1.GetTransitGatewayConnectionOptions{}
			getTransitGatewayConnectionOptions.SetTransitGatewayID(badinstanceID)
			getTransitGatewayConnectionOptions.SetID(badinstanceID)
			header := map[string]string{
				"Content-type": "application/json",
			}
			getTransitGatewayConnectionOptions.SetHeaders(header)

			It(`Failed to get resource by instanceID`, func() {
				shouldSkipTest()

				result, detailedResponse, err := service.GetTransitGatewayConnection(getTransitGatewayConnectionOptions)
				Expect(result).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(404))
				Expect(err).Should(HaveOccurred())
			})
		})
	})

	Describe(`UpdateTransitGatewayConnection(updateTransitGatewayConnectionOptions *UpdateTransitGatewayConnectionOptions)`, func() {
		Context(`Success: UPDATE Transit Gateway CLASSIC Connection`, func() {
			It(`Successfully update CLASSIC Connection`, func() {
				shouldSkipTest()

				gatewayID := os.Getenv("GATEWAY_INSTANCE_ID")
				instanceID := os.Getenv("CLASSIC_CONN_INSTANCE_ID")
				updateConnectionName := "UPDATED-" + os.Getenv("CLASSIC_CONN_INSTANCE_NAME")
				updateTransitGatewayConnectionOptions := service.NewUpdateTransitGatewayConnectionOptions(
					gatewayID, instanceID).
					SetName(updateConnectionName)

				result, detailedResponse, err := service.UpdateTransitGatewayConnection(updateTransitGatewayConnectionOptions)
				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(200))

				Expect(*result.Name).To(Equal(updateConnectionName))
				Expect(*result.CreatedAt).NotTo(Equal(""))
				Expect(*result.UpdatedAt).NotTo(Equal(""))
				Expect(*result.Status).To(Equal("attached"))
				Expect(*result.NetworkType).To(Equal("classic"))
				Expect(*result.ID).To(Equal(os.Getenv("CLASSIC_CONN_INSTANCE_ID")))

				os.Setenv("CLASSIC_CONN_INSTANCE_NAME", *result.Name)
			})
		})

		Context(`Success: UPDATE Transit Gateway VPC Connection`, func() {
			It(`Successfully update VPC Connection`, func() {
				shouldSkipTest()

				gatewayID := os.Getenv("GATEWAY_INSTANCE_ID")
				instanceID := os.Getenv("VPC_CONN_INSTANCE_ID")
				updateConnectionName := "UPDATED-" + os.Getenv("VPC_CONN_INSTANCE_NAME")
				updateTransitGatewayConnectionOptions := service.NewUpdateTransitGatewayConnectionOptions(
					gatewayID, instanceID).
					SetName(updateConnectionName)

				result, detailedResponse, err := service.UpdateTransitGatewayConnection(updateTransitGatewayConnectionOptions)
				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(200))

				Expect(*result.Name).To(Equal(updateConnectionName))
				Expect(*result.CreatedAt).NotTo(Equal(""))
				Expect(*result.UpdatedAt).NotTo(Equal(""))
				Expect(*result.NetworkType).To(Equal("vpc"))
				Expect(*result.Status).To(Equal("attached"))
				Expect(*result.NetworkID).To(Equal(os.Getenv("VPC_CRN")))
				Expect(*result.ID).To(Equal(os.Getenv("VPC_CONN_INSTANCE_ID")))

				os.Setenv("VPC_CONN_INSTANCE_NAME", *result.Name)
			})
		})

		Context(`Success: UPDATE Transit Gateway DL Connection`, func() {
			It(`Successfully update DL Connection`, func() {
				shouldSkipTest()

				gatewayID := os.Getenv("GATEWAY_INSTANCE_ID")
				instanceID := os.Getenv("DL_CONN_INSTANCE_ID")
				updateConnectionName := "UPDATED-" + os.Getenv("DL_CONN_INSTANCE_NAME")
				updateTransitGatewayConnectionOptions := service.NewUpdateTransitGatewayConnectionOptions(
					gatewayID, instanceID).
					SetName(updateConnectionName)

				result, detailedResponse, err := service.UpdateTransitGatewayConnection(updateTransitGatewayConnectionOptions)
				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(200))

				Expect(*result.Name).To(Equal(updateConnectionName))
				Expect(*result.CreatedAt).NotTo(Equal(""))
				Expect(*result.UpdatedAt).NotTo(Equal(""))
				Expect(*result.Status).To(Equal("attached"))
				Expect(*result.NetworkType).To(Equal("directlink"))
				Expect(*result.NetworkID).To(Equal(os.Getenv("DL_CRN")))
				Expect(*result.ID).To(Equal(os.Getenv("DL_CONN_INSTANCE_ID")))

				os.Setenv("DL_CONN_INSTANCE_NAME", *result.Name)
			})
		})

		Context(`Success: UPDATE Transit Gateway GRE Connection`, func() {
			It(`Successfully update GRE Connection`, func() {
				shouldSkipTest()

				gatewayID := os.Getenv("GATEWAY_INSTANCE_ID")
				instanceID := os.Getenv("GRE_CONN_INSTANCE_ID")
				updateConnectionName := "UPDATED-" + os.Getenv("GRE_CONN_INSTANCE_NAME")
				updateTransitGatewayConnectionOptions := service.NewUpdateTransitGatewayConnectionOptions(
					gatewayID, instanceID).
					SetName(updateConnectionName)

				result, detailedResponse, err := service.UpdateTransitGatewayConnection(updateTransitGatewayConnectionOptions)
				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(200))

				Expect(*result.Name).To(Equal(updateConnectionName))
				Expect(*result.CreatedAt).NotTo(Equal(""))
				Expect(*result.UpdatedAt).NotTo(Equal(""))
				Expect(*result.Status).To(Equal("attached"))
				Expect(*result.NetworkType).To(Equal("gre_tunnel"))
				Expect(*result.ID).To(Equal(os.Getenv("GRE_CONN_INSTANCE_ID")))
				Expect(*result.BaseConnectionID).To(Equal(os.Getenv("CLASSIC_CONN_INSTANCE_ID")))

				os.Setenv("GRE_CONN_INSTANCE_NAME", *result.Name)
			})
		})

		Context(`Failure: UPDATE connection by instanceID`, func() {
			badinstanceID := "abc123"
			instanceName := "UPDATE-" + strconv.FormatInt(timestamp, 10)
			updateTransitGatewayConnectionOptions := &transitgatewayapisv1.UpdateTransitGatewayConnectionOptions{}
			updateTransitGatewayConnectionOptions.SetTransitGatewayID(badinstanceID)
			updateTransitGatewayConnectionOptions.SetID(badinstanceID)
			updateTransitGatewayConnectionOptions.SetName(instanceName)
			header := map[string]string{
				"Content-type": "application/json",
			}
			updateTransitGatewayConnectionOptions.SetHeaders(header)

			It(`Failed to update gateway by instanceID`, func() {
				shouldSkipTest()

				result, detailedResponse, err := service.UpdateTransitGatewayConnection(updateTransitGatewayConnectionOptions)
				Expect(result).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(404))
				Expect(err).Should(HaveOccurred())
			})
		})
	})

	Describe(`ListTransitGatewayConnections(listTransitGatewayConnectionsOptions *ListTransitGatewayConnectionsOptions)`, func() {
		Context(`Success: LIST Transit Gateway Connections`, func() {
			header := map[string]string{
				"Content-type": "application/json",
			}
			It(`Successfully list all gateway connections`, func() {
				shouldSkipTest()

				gatewayID := os.Getenv("GATEWAY_INSTANCE_ID")
				listTransitGatewayConnectionsOptions := service.NewListTransitGatewayConnectionsOptions(gatewayID).
					SetTransitGatewayID(gatewayID).
					SetHeaders(header)

				result, detailedResponse, err := service.ListTransitGatewayConnections(listTransitGatewayConnectionsOptions)
				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(200))
				Expect(len(result.Connections)).Should(BeNumerically(">", 0))

				dl_found := false
				vpc_found := false
				gre_found := false
				classic_found := false
				for _, conn := range result.Connections {
					if *conn.ID == os.Getenv("VPC_CONN_INSTANCE_ID") {
						Expect(*conn.CreatedAt).NotTo(Equal(""))
						Expect(*conn.UpdatedAt).NotTo(Equal(""))
						Expect(*conn.NetworkType).To(Equal("vpc"))
						Expect(*conn.Status).To(Equal("attached"))
						Expect(*conn.NetworkID).To(Equal(os.Getenv("VPC_CRN")))
						Expect(*conn.Name).To(Equal(os.Getenv("VPC_CONN_INSTANCE_NAME")))
						vpc_found = true

					} else if *conn.ID == os.Getenv("DL_CONN_INSTANCE_ID") {
						Expect(*conn.CreatedAt).NotTo(Equal(""))
						Expect(*conn.UpdatedAt).NotTo(Equal(""))
						Expect(*conn.Status).To(Equal("attached"))
						Expect(*conn.NetworkType).To(Equal("directlink"))
						Expect(*conn.NetworkID).To(Equal(os.Getenv("DL_CRN")))
						Expect(*conn.Name).To(Equal(os.Getenv("DL_CONN_INSTANCE_NAME")))
						dl_found = true

					} else if *conn.ID == os.Getenv("GRE_CONN_INSTANCE_ID") {
						Expect(*conn.CreatedAt).NotTo(Equal(""))
						Expect(*conn.UpdatedAt).NotTo(Equal(""))
						Expect(*conn.Status).To(Equal("attached"))
						Expect(*conn.NetworkType).To(Equal("gre_tunnel"))
						Expect(*conn.Name).To(Equal(os.Getenv("GRE_CONN_INSTANCE_NAME")))
						Expect(*conn.BaseConnectionID).To(Equal(os.Getenv("CLASSIC_CONN_INSTANCE_ID")))
						gre_found = true

					} else if *conn.ID == os.Getenv("CLASSIC_CONN_INSTANCE_ID") {
						Expect(*conn.CreatedAt).NotTo(Equal(""))
						Expect(*conn.UpdatedAt).NotTo(Equal(""))
						Expect(*conn.Status).To(Equal("attached"))
						Expect(*conn.NetworkType).To(Equal("classic"))
						Expect(*conn.Name).To(Equal(os.Getenv("CLASSIC_CONN_INSTANCE_NAME")))
						classic_found = true
					}
				}
				Expect(dl_found).To(Equal(true))
				Expect(vpc_found).To(Equal(true))
				Expect(gre_found).To(Equal(true))
				Expect(classic_found).To(Equal(true))
			})
		})
	})

	///////////////////////////////////////////////////////////////////////////////
	//                  Transit Gateway Route Reports Tests                      //
	///////////////////////////////////////////////////////////////////////////////

	Describe(`CreateTransitGatewayRouteReport(createTransitGatewayRouteReportOptions *CreateTransitGatewayRouteReportOptions)`, func() {
		Context(`Success: POST Gateway Route Report`, func() {
			It(`Successfully create Gateway Route Report`, func() {
				shouldSkipTest()

				header := map[string]string{
					"Content-type": "application/json",
				}
				gatewayID := os.Getenv("GATEWAY_INSTANCE_ID")
				createTransitGatewayRouteReportOptions := service.NewCreateTransitGatewayRouteReportOptions(gatewayID).
					SetHeaders(header)

				result, detailedResponse, err := service.CreateTransitGatewayRouteReport(createTransitGatewayRouteReportOptions)
				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(202))

				Expect(*result.ID).NotTo(Equal(""))
				Expect(*result.Status).NotTo(Equal(""))
				Expect(*result.CreatedAt).NotTo(Equal(""))
				Expect(*result.UpdatedAt).NotTo(Equal(""))

				os.Setenv("RR_INSTANCE_ID", *result.ID)
			})

			It("Successfully waits for RR to report as complete", func() {
				shouldSkipTest()

				instanceID := os.Getenv("RR_INSTANCE_ID")
				gatewayID := os.Getenv("GATEWAY_INSTANCE_ID")
				isResourceAvailable(service, gatewayID, "", instanceID)
			})
		})

		Context(`Failure: POST new route report`, func() {
			badinstanceID := "testString"
			header := map[string]string{
				"Content-type": "application/json",
			}

			createTransitGatewayRouteReportOptions := &transitgatewayapisv1.CreateTransitGatewayRouteReportOptions{}
			createTransitGatewayRouteReportOptions.SetTransitGatewayID(badinstanceID)
			createTransitGatewayRouteReportOptions.SetHeaders(header)

			It(`Fail to create new route report`, func() {
				shouldSkipTest()

				result, detailedResponse, err := service.CreateTransitGatewayRouteReport(createTransitGatewayRouteReportOptions)
				Expect(result).To(BeNil())
				Expect(detailedResponse.StatusCode).ToNot(Equal(200))
				Expect(err).Should(HaveOccurred())
			})
		})
	})

	Describe(`GetTransitGatewayRouteReport(getTransitGatewayRouteReportOptions *GetTransitGatewayRouteReportOptions)`, func() {
		Context(`Success: GET Gateway Route Report`, func() {
			It(`Successfully get Route Report by instanceID`, func() {
				shouldSkipTest()

				instanceID := os.Getenv("RR_INSTANCE_ID")
				gatewayID := os.Getenv("GATEWAY_INSTANCE_ID")
				getTransitGatewayRouteReportOptions := service.NewGetTransitGatewayRouteReportOptions(gatewayID, instanceID)

				result, detailedResponse, err := service.GetTransitGatewayRouteReport(getTransitGatewayRouteReportOptions)
				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(200))

				Expect(*result.CreatedAt).NotTo(Equal(""))
				Expect(*result.UpdatedAt).NotTo(Equal(""))
				Expect(&result.Connections).NotTo(BeNil())
				Expect(*result.Status).To(Equal("complete"))
				Expect(*result.ID).To(Equal(os.Getenv("RR_INSTANCE_ID")))
				Expect(len(result.Connections)).Should(BeNumerically(">", 0))

				dl_found := false
				vpc_found := false
				gre_found := false
				classic_found := false

				for _, conn := range result.Connections {
					if *conn.ID == os.Getenv("VPC_CONN_INSTANCE_ID") {
						Expect(*conn.Type).To(Equal("vpc"))
						Expect(*conn.Name).To(Equal(os.Getenv("VPC_CONN_INSTANCE_NAME")))
						vpc_found = true

					} else if *conn.ID == os.Getenv("DL_CONN_INSTANCE_ID") {
						Expect(*conn.Type).To(Equal("directlink"))
						Expect(*conn.Name).To(Equal(os.Getenv("DL_CONN_INSTANCE_NAME")))
						dl_found = true

					} else if *conn.ID == os.Getenv("GRE_CONN_INSTANCE_ID") {
						Expect(*conn.Type).To(Equal("gre_tunnel"))
						Expect(*conn.Name).To(Equal(os.Getenv("GRE_CONN_INSTANCE_NAME")))
						gre_found = true

					} else if *conn.ID == os.Getenv("CLASSIC_CONN_INSTANCE_ID") {
						Expect(*conn.Type).To(Equal("classic"))
						Expect(*conn.Name).To(Equal(os.Getenv("CLASSIC_CONN_INSTANCE_NAME")))
						classic_found = true
					}
				}

				Expect(dl_found).To(Equal(true))
				Expect(vpc_found).To(Equal(true))
				Expect(gre_found).To(Equal(true))
				Expect(classic_found).To(Equal(true))
			})
		})

		Context(`Failure: GET route report by instanceID`, func() {
			badinstanceID := "abc123"
			getTransitGatewayRouteReportOptions := &transitgatewayapisv1.GetTransitGatewayRouteReportOptions{}
			getTransitGatewayRouteReportOptions.SetTransitGatewayID(badinstanceID)
			getTransitGatewayRouteReportOptions.SetID(badinstanceID)
			header := map[string]string{
				"Content-type": "application/json",
			}
			getTransitGatewayRouteReportOptions.SetHeaders(header)

			It(`Failed to get route report by instanceID`, func() {
				shouldSkipTest()

				result, detailedResponse, err := service.GetTransitGatewayRouteReport(getTransitGatewayRouteReportOptions)
				Expect(result).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(404))
				Expect(err).Should(HaveOccurred())
			})
		})
	})

	Describe(`ListTransitGatewayRouteReports(listTransitGatewayRouteReportsOptions *ListTransitGatewayRouteReportsOptions)`, func() {
		Context(`Success: LIST Gateway Route Reports`, func() {
			It(`Successfully list all gateway route reports`, func() {
				shouldSkipTest()

				header := map[string]string{
					"Content-type": "application/json",
				}
				gatewayID := os.Getenv("GATEWAY_INSTANCE_ID")
				listTransitGatewayRouteReportsOptions := service.NewListTransitGatewayRouteReportsOptions(gatewayID).
					SetHeaders(header)

				result, detailedResponse, err := service.ListTransitGatewayRouteReports(listTransitGatewayRouteReportsOptions)
				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(200))
				Expect(len(result.RouteReports)).Should(BeNumerically(">", 0))

				found := false
				for _, rr := range result.RouteReports {
					if *rr.ID == os.Getenv("RR_INSTANCE_ID") {
						Expect(*rr.CreatedAt).NotTo(Equal(""))
						Expect(*rr.UpdatedAt).NotTo(Equal(""))
						Expect(&rr.Connections).NotTo(BeNil())
						Expect(*rr.Status).To(Equal("complete"))
						Expect(len(rr.Connections)).Should(BeNumerically(">", 0))

						found = true
						break
					}
				}
				Expect(found).To(Equal(true))
			})
		})
	})

	///////////////////////////////////////////////////////////////////////////////
	//                   DELETE Transit Gateway Route Report                     //
	///////////////////////////////////////////////////////////////////////////////

	Describe(`DeleteTransitGatewayRouteReport(deleteTransitGatewayRouteReportOptions *DeleteTransitGatewayRouteReportOptions)`, func() {
		Context(`Success: DELETE Gateway Route Report`, func() {
			It(`Successfully delete Route Report by instanceID`, func() {
				shouldSkipTest()

				instanceID := os.Getenv("RR_INSTANCE_ID")
				gatewayID := os.Getenv("GATEWAY_INSTANCE_ID")
				deleteTransitGatewayRouteReportOptions := service.NewDeleteTransitGatewayRouteReportOptions(gatewayID, instanceID)

				detailedResponse, err := service.DeleteTransitGatewayRouteReport(deleteTransitGatewayRouteReportOptions)
				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(204))
			})

			It("Successfully waits for Gateway Route Report to report as deleted", func() {
				shouldSkipTest()

				instanceID := os.Getenv("RR_INSTANCE_ID")
				gatewayID := os.Getenv("GATEWAY_INSTANCE_ID")
				deleteCheckTest(service, gatewayID, "", instanceID, "")
			})
		})

		Context(`Failure: DELETE route report by instanceID`, func() {
			badinstanceID := "abc123"
			deleteTransitGatewayRouteReportOptions := &transitgatewayapisv1.DeleteTransitGatewayRouteReportOptions{}
			deleteTransitGatewayRouteReportOptions.SetTransitGatewayID(badinstanceID)
			deleteTransitGatewayRouteReportOptions.SetID(badinstanceID)
			header := map[string]string{
				"Content-type": "application/json",
			}
			deleteTransitGatewayRouteReportOptions.SetHeaders(header)

			It(`Failed to delete route report by instanceID`, func() {
				shouldSkipTest()

				detailedResponse, err := service.DeleteTransitGatewayRouteReport(deleteTransitGatewayRouteReportOptions)
				Expect(detailedResponse.StatusCode).To(Equal(404))
				Expect(err).Should(HaveOccurred())
			})
		})
	})

	///////////////////////////////////////////////////////////////////////////////
	//              Transit Gateway Connection Prefix Filter Tests               //
	///////////////////////////////////////////////////////////////////////////////

	Describe(`CreateTransitGatewayConnectionPrefixFilter(reateTransitGatewayConnectionPrefixFilterOptions *CreateTransitGatewayConnectionPrefixFilterOptions)`, func() {
		Context(`Success: POST Gateway Connection Prefix Filter`, func() {
			It(`Successfully create Gateway Connection Prefix Filter`, func() {
				shouldSkipTest()

				header := map[string]string{
					"Content-type": "application/json",
				}
				gatewayID := os.Getenv("GATEWAY_INSTANCE_ID")
				classicConnID := os.Getenv("CLASSIC_CONN_INSTANCE_ID")
				action := "permit"
				prefix := "192.168.100.0/24"

				createTransitGatewayConnectionPrefixFilterOptions := service.NewCreateTransitGatewayConnectionPrefixFilterOptions(gatewayID, classicConnID, action, prefix).
					SetHeaders(header)

				result, detailedResponse, err := service.CreateTransitGatewayConnectionPrefixFilter(createTransitGatewayConnectionPrefixFilterOptions)
				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(201))

				Expect(*result.ID).NotTo(Equal(""))
				os.Setenv("PF_INSTANCE_ID", *result.ID)

				Expect(*result.Prefix).To(Equal(prefix))
				Expect(*result.Action).To(Equal(action))
				Expect(*result.CreatedAt).NotTo(Equal(""))
				Expect(*result.UpdatedAt).NotTo(Equal(""))
			})
		})

		Context(`Failure: POST Gateway Connection Prefix Filter`, func() {
			It(`Fail to create a new Gateway Connection Prefix Filter`, func() {
				shouldSkipTest()

				createTransitGatewayConnectionPrefixFilterOptions := &transitgatewayapisv1.CreateTransitGatewayConnectionPrefixFilterOptions{}
				gatewayID := os.Getenv("GATEWAY_INSTANCE_ID")
				classicConnID := os.Getenv("CLASSIC_CONN_INSTANCE_ID")

				createTransitGatewayConnectionPrefixFilterOptions.SetTransitGatewayID(gatewayID)
				createTransitGatewayConnectionPrefixFilterOptions.SetID(classicConnID)
				createTransitGatewayConnectionPrefixFilterOptions.SetAction("testString")
				createTransitGatewayConnectionPrefixFilterOptions.SetPrefix("testString")

				header := map[string]string{
					"Content-type": "application/json",
				}
				createTransitGatewayConnectionPrefixFilterOptions.SetHeaders(header)

				result, detailedResponse, err := service.CreateTransitGatewayConnectionPrefixFilter(createTransitGatewayConnectionPrefixFilterOptions)
				Expect(result).To(BeNil())
				Expect(detailedResponse.StatusCode).ToNot(Equal(200))
				Expect(err).Should(HaveOccurred())
			})
		})
	})

	Describe(`ListTransitGatewayConnectionPrefixFilters(listTransitGatewayConnectionPrefixFiltersOptions *ListTransitGatewayConnectionPrefixFiltersOptions)`, func() {
		Context(`Success: LIST Transit Gateway Connection Prefix Filters`, func() {
			It(`Successfully list all transit gateway connection prefix filters`, func() {
				shouldSkipTest()

				gatewayID := os.Getenv("GATEWAY_INSTANCE_ID")
				classicConnID := os.Getenv("CLASSIC_CONN_INSTANCE_ID")
				pfID := os.Getenv("PF_INSTANCE_ID")

				header := map[string]string{
					"Content-type": "application/json",
				}
				listTransitGatewayConnectionPrefixFiltersOptions := service.NewListTransitGatewayConnectionPrefixFiltersOptions(gatewayID, classicConnID).
					SetHeaders(header)

				result, detailedResponse, err := service.ListTransitGatewayConnectionPrefixFilters(listTransitGatewayConnectionPrefixFiltersOptions)
				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(200))
				Expect(len(result.PrefixFilters)).Should(BeNumerically(">", 0))

				found := false
				for _, pf := range result.PrefixFilters {
					if *pf.ID == pfID {
						Expect(*pf.Action).To(Equal("permit"))
						Expect(*pf.Prefix).To(Equal("192.168.100.0/24"))
						Expect(*pf.CreatedAt).NotTo(Equal(""))
						Expect(*pf.UpdatedAt).NotTo(Equal(""))

						found = true
						break
					}
				}
				Expect(found).To(Equal(true))
			})
		})
	})

	Describe(`UpdateTransitGatewayConnectionPrefixFilter(lupdateTransitGatewayConnectionPrefixFilterOptions *UpdateTransitGatewayConnectionPrefixFilterOptions)`, func() {
		Context(`Success: UPDATE (PATCH) Transit Gateway Connection Prefix Filter`, func() {
			It(`Successfully update transit gateway connection prefix filter`, func() {
				shouldSkipTest()

				gatewayID := os.Getenv("GATEWAY_INSTANCE_ID")
				classicConnID := os.Getenv("CLASSIC_CONN_INSTANCE_ID")
				pfID := os.Getenv("PF_INSTANCE_ID")

				header := map[string]string{
					"Content-type": "application/json",
				}

				newAction := "deny"

				updateTransitGatewayConnectionPrefixFiltersOptions := service.NewUpdateTransitGatewayConnectionPrefixFilterOptions(gatewayID, classicConnID, pfID)

				updateTransitGatewayConnectionPrefixFiltersOptions.SetHeaders(header)
				updateTransitGatewayConnectionPrefixFiltersOptions.SetAction(newAction)

				result, detailedResponse, err := service.UpdateTransitGatewayConnectionPrefixFilter(updateTransitGatewayConnectionPrefixFiltersOptions)
				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(200))
				Expect(*result.ID).To(Equal(pfID))
				Expect(*result.Prefix).To(Equal("192.168.100.0/24"))
				Expect(*result.Action).To(Equal(newAction))
				Expect(*result.CreatedAt).NotTo(Equal(""))
				Expect(*result.UpdatedAt).NotTo(Equal(""))
			})
		})
	})

	Describe(`GetTransitGatewayConnectionPrefixFilter(getTransitGatewayConnectionPrefixFilterOptions *GetTransitGatewayConnectionPrefixFilterOptions)`, func() {
		Context(`Success: GET Transit Gateway Connection Prefix Filter`, func() {
			It(`Successfully get transit gateway connection prefix filter`, func() {
				shouldSkipTest()

				gatewayID := os.Getenv("GATEWAY_INSTANCE_ID")
				classicConnID := os.Getenv("CLASSIC_CONN_INSTANCE_ID")
				pfID := os.Getenv("PF_INSTANCE_ID")

				header := map[string]string{
					"Content-type": "application/json",
				}

				getTransitGatewayConnectionPrefixFiltersOptions := service.NewGetTransitGatewayConnectionPrefixFilterOptions(gatewayID, classicConnID, pfID).SetHeaders(header)

				result, detailedResponse, err := service.GetTransitGatewayConnectionPrefixFilter(getTransitGatewayConnectionPrefixFiltersOptions)
				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(200))
				Expect(*result.ID).To(Equal(pfID))
				Expect(*result.Prefix).To(Equal("192.168.100.0/24"))
				Expect(*result.Action).To(Equal("deny"))
				Expect(*result.CreatedAt).NotTo(Equal(""))
				Expect(*result.UpdatedAt).NotTo(Equal(""))
			})
		})
	})

	///////////////////////////////////////////////////////////////////////////////
	//             DELETE Transit Gateway Connection Prefix Filters              //
	///////////////////////////////////////////////////////////////////////////////

	Describe(`DeleteTransitGatewayConnectionPrefixFilter(deleteTransitGatewayConnectionPrefixFilterOptions *DeleteTransitGatewayConnectionPrefixFilterOptions)`, func() {
		Context(`Success: DELETE Gateway Connection Prefix Filter`, func() {
			It(`Successfully delete gateway connection prefix filter by instanceID`, func() {
				shouldSkipTest()

				pfID := os.Getenv("PF_INSTANCE_ID")
				gatewayID := os.Getenv("GATEWAY_INSTANCE_ID")
				classicConnID := os.Getenv("CLASSIC_CONN_INSTANCE_ID")
				deleteTransitGatewayConnectionPrefixFilterOptions := service.NewDeleteTransitGatewayConnectionPrefixFilterOptions(gatewayID, classicConnID, pfID)

				detailedResponse, err := service.DeleteTransitGatewayConnectionPrefixFilter(deleteTransitGatewayConnectionPrefixFilterOptions)
				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(204))
			})

			It("Successfully waits for Gateway Connection Prefix Filter to report as deleted", func() {
				shouldSkipTest()

				pfID := os.Getenv("PF_INSTANCE_ID")
				gatewayID := os.Getenv("GATEWAY_INSTANCE_ID")
				classicConnID := os.Getenv("CLASSIC_CONN_INSTANCE_ID")
				deleteCheckTest(service, gatewayID, classicConnID, "", pfID)
			})
		})

		Context(`Failure: DELETE prefix filter by FilterID`, func() {
			It(`Successfully verify DELETE failure by FilterID`, func() {
				shouldSkipTest()
				badPfInstanceID := "abc123"
				gatewayID := os.Getenv("GATEWAY_INSTANCE_ID")
				classicConnID := os.Getenv("CLASSIC_CONN_INSTANCE_ID")

				deleteTransitGatewayConnectionPrefixFilterOptions := &transitgatewayapisv1.DeleteTransitGatewayConnectionPrefixFilterOptions{}
				deleteTransitGatewayConnectionPrefixFilterOptions.SetTransitGatewayID(gatewayID)
				deleteTransitGatewayConnectionPrefixFilterOptions.SetID(classicConnID)
				deleteTransitGatewayConnectionPrefixFilterOptions.SetFilterID(badPfInstanceID)

				header := map[string]string{
					"Content-type": "application/json",
				}
				deleteTransitGatewayConnectionPrefixFilterOptions.SetHeaders(header)

				detailedResponse, err := service.DeleteTransitGatewayConnectionPrefixFilter(deleteTransitGatewayConnectionPrefixFilterOptions)
				Expect(detailedResponse.StatusCode).To(Equal(404))
				Expect(err).Should(HaveOccurred())
			})
		})
	})

	///////////////////////////////////////////////////////////////////////////////
	//                    DELETE Transit Gateway Connections                     //
	///////////////////////////////////////////////////////////////////////////////

	Describe(`DeleteTransitGatewayConnection(deleteTransitGatewayConnectionOptions *DeleteTransitGatewayConnectionOptions)`, func() {
		Context(`Success: DELETE Transit GRE connection by instanceID`, func() {
			It(`Successfully delete GRE connection by instanceID`, func() {
				shouldSkipTest()

				gatewayID := os.Getenv("GATEWAY_INSTANCE_ID")
				instanceID := os.Getenv("GRE_CONN_INSTANCE_ID")
				deleteTransitGatewayConnectionOptions := service.NewDeleteTransitGatewayConnectionOptions(gatewayID, instanceID)

				detailedResponse, err := service.DeleteTransitGatewayConnection(deleteTransitGatewayConnectionOptions)
				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(204))
			})

			It("Successfully waits for GRE connection to report as deleted", func() {
				shouldSkipTest()

				gatewayID := os.Getenv("GATEWAY_INSTANCE_ID")
				instanceID := os.Getenv("GRE_CONN_INSTANCE_ID")
				deleteCheckTest(service, gatewayID, instanceID, "", "")
			})
		})

		Context(`Success: DELETE Transit VPC connection by instanceID`, func() {
			It(`Successfully delete VPC connection by instanceID`, func() {
				shouldSkipTest()

				gatewayID := os.Getenv("GATEWAY_INSTANCE_ID")
				instanceID := os.Getenv("VPC_CONN_INSTANCE_ID")
				deleteTransitGatewayConnectionOptions := service.NewDeleteTransitGatewayConnectionOptions(gatewayID, instanceID)

				detailedResponse, err := service.DeleteTransitGatewayConnection(deleteTransitGatewayConnectionOptions)
				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(204))
			})

			It("Successfully waits for VPC connection to report as deleted", func() {
				shouldSkipTest()

				gatewayID := os.Getenv("GATEWAY_INSTANCE_ID")
				instanceID := os.Getenv("VPC_CONN_INSTANCE_ID")
				deleteCheckTest(service, gatewayID, instanceID, "", "")
			})
		})

		Context(`Success: DELETE Transit DL connection by instanceID`, func() {
			It(`Successfully delete DL connection by instanceID`, func() {
				shouldSkipTest()

				gatewayID := os.Getenv("GATEWAY_INSTANCE_ID")
				instanceID := os.Getenv("DL_CONN_INSTANCE_ID")
				deleteTransitGatewayConnectionOptions := service.NewDeleteTransitGatewayConnectionOptions(gatewayID, instanceID)

				detailedResponse, err := service.DeleteTransitGatewayConnection(deleteTransitGatewayConnectionOptions)
				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(204))
			})

			It("Successfully waits for DL connection to report as deleted", func() {
				shouldSkipTest()

				gatewayID := os.Getenv("GATEWAY_INSTANCE_ID")
				instanceID := os.Getenv("DL_CONN_INSTANCE_ID")
				deleteCheckTest(service, gatewayID, instanceID, "", "")
			})
		})

		Context(`Success: DELETE Transit CLASSIC connection by instanceID`, func() {
			It(`Successfully delete CLASSIC connection by instanceID`, func() {
				shouldSkipTest()

				gatewayID := os.Getenv("GATEWAY_INSTANCE_ID")
				instanceID := os.Getenv("CLASSIC_CONN_INSTANCE_ID")
				deleteTransitGatewayConnectionOptions := service.NewDeleteTransitGatewayConnectionOptions(gatewayID, instanceID)

				detailedResponse, err := service.DeleteTransitGatewayConnection(deleteTransitGatewayConnectionOptions)
				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(204))
			})

			It("Successfully waits for CLASSIC connection to report as deleted", func() {
				shouldSkipTest()

				gatewayID := os.Getenv("GATEWAY_INSTANCE_ID")
				instanceID := os.Getenv("CLASSIC_CONN_INSTANCE_ID")
				deleteCheckTest(service, gatewayID, instanceID, "", "")
			})
		})

		Context(`Failure: DELETE connection by instanceID`, func() {
			badinstanceID := "abc123"
			deleteTransitGatewayConnectionOptions := &transitgatewayapisv1.DeleteTransitGatewayConnectionOptions{}
			deleteTransitGatewayConnectionOptions.SetTransitGatewayID(badinstanceID)
			deleteTransitGatewayConnectionOptions.SetID(badinstanceID)
			header := map[string]string{
				"Content-type": "application/json",
			}
			deleteTransitGatewayConnectionOptions.SetHeaders(header)

			It(`Failed to delete resource by instanceID`, func() {
				shouldSkipTest()

				detailedResponse, err := service.DeleteTransitGatewayConnection(deleteTransitGatewayConnectionOptions)
				Expect(detailedResponse.StatusCode).To(Equal(404))
				Expect(err).Should(HaveOccurred())
			})
		})
	})

	///////////////////////////////////////////////////////////////////////////////
	//                           DELETE Transit Gateway                          //
	///////////////////////////////////////////////////////////////////////////////

	Describe(`DeleteTransitGateway(deleteTransitGatewayOptions *DeleteTransitGatewayOptions)`, func() {
		Context(`Success: DELETE delete gateway by instanceID`, func() {
			It(`Successfully delete gateway by instanceID`, func() {
				shouldSkipTest()

				instanceID := os.Getenv("GATEWAY_INSTANCE_ID")
				deleteTransitGatewayOptions := service.NewDeleteTransitGatewayOptions(instanceID)

				detailedResponse, err := service.DeleteTransitGateway(deleteTransitGatewayOptions)
				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(204))
			})
		})

		Context(`Failure: DELETE gateway by instanceID`, func() {
			badinstanceID := "abc123"
			deleteTransitGatewayOptions := &transitgatewayapisv1.DeleteTransitGatewayOptions{}
			deleteTransitGatewayOptions.SetID(badinstanceID)
			header := map[string]string{
				"Content-type": "application/json",
			}
			deleteTransitGatewayOptions.SetHeaders(header)

			It(`Failed to delete gateway by instanceID`, func() {
				shouldSkipTest()

				detailedResponse, err := service.DeleteTransitGateway(deleteTransitGatewayOptions)
				Expect(detailedResponse.StatusCode).To(Equal(404))
				Expect(err).Should(HaveOccurred())
			})
		})
	})
})

///////////////////////////////////////////////////////////////////////////////
//                           Test Helper Methods                             //
///////////////////////////////////////////////////////////////////////////////

// deleteResourceTest deletes a Transit Resource: Resource delete might not be
// instantaneous.  Poll the Resource looking for a not found. Fail after 4 min
func deleteCheckTest(service *transitgatewayapisv1.TransitGatewayApisV1, gatewayID, connID, rrID, pfID string) {
	timer := 0
	statusCode := 0

	for {
		if connID != "" && rrID == "" && pfID == "" {
			getTransitResourceOptions := service.NewGetTransitGatewayConnectionOptions(gatewayID, connID)
			_, detailedResponse, _ := service.GetTransitGatewayConnection(getTransitResourceOptions)
			statusCode = detailedResponse.StatusCode
		} else if pfID != "" {
			getTransitResourceOptions := service.NewGetTransitGatewayConnectionPrefixFilterOptions(gatewayID, connID, pfID)
			_, detailedResponse, _ := service.GetTransitGatewayConnectionPrefixFilter(getTransitResourceOptions)
			statusCode = detailedResponse.StatusCode
		} else if connID == "" && rrID != "" {
			getTransitResourceOptions := service.NewGetTransitGatewayRouteReportOptions(gatewayID, rrID)
			_, detailedResponse, _ := service.GetTransitGatewayRouteReport(getTransitResourceOptions)
			statusCode = detailedResponse.StatusCode
		} else {
			getTransitResourceOptions := service.NewGetTransitGatewayOptions(gatewayID)
			_, detailedResponse, _ := service.GetTransitGateway(getTransitResourceOptions)
			statusCode = detailedResponse.StatusCode
		}
		// Break loop if a 404 code is found
		if statusCode == 404 {
			Expect(statusCode).To(Equal(404))
			break
		}

		// Other than 404: See if the timeout value has been reached.
		// If so, exit with failure: 4 min timer (24x10sec).
		if timer > 24 {
			Expect(statusCode).To(Equal(404))
			break
		} else {
			// Still exists, wait 5 sec
			time.Sleep(time.Duration(10) * time.Second)
			timer = timer + 1
		}
	}
}

// isResourceAvailable checks until the resource status is available/attached/complete. Fail after 2 min.
func isResourceAvailable(service *transitgatewayapisv1.TransitGatewayApisV1, gatewayID, connID, rrID string) {
	timer := 0
	delay := 5
	resourceStatus := "available"

	for {
		breaker := 0
		if connID == "" && rrID != "" {
			getTransitResourceOptions := service.NewGetTransitGatewayRouteReportOptions(gatewayID, rrID)
			response, _, _ := service.GetTransitGatewayRouteReport(getTransitResourceOptions)
			delay = 10
			if *response.Status == "complete" {
				Expect(*response.ID).To(Equal(rrID))
				Expect(&response.Connections).NotTo(BeNil())
				Expect(*response.Status).To(Equal("complete"))
				breaker = 1
			}
		} else if connID != "" && rrID == "" {
			getTransitResourceOptions := service.NewGetTransitGatewayConnectionOptions(gatewayID, connID)
			response, _, _ := service.GetTransitGatewayConnection(getTransitResourceOptions)
			if *response.Status == "attached" {
				Expect(*response.ID).To(Equal(connID))
				Expect(*response.NetworkType).NotTo(Equal(""))
				Expect(*response.Status).To(Equal("attached"))
				breaker = 1
			}
		} else {
			getTransitResourceOptions := service.NewGetTransitGatewayOptions(gatewayID)
			response, _, _ := service.GetTransitGateway(getTransitResourceOptions)
			if *response.Status == "available" {
				Expect(*response.Crn).NotTo(Equal(""))
				Expect(*response.ID).To(Equal(gatewayID))
				Expect(*response.Status).To(Equal("available"))
				breaker = 1
			}
		}
		// Break loop if resourse is available!
		if breaker != 0 {
			Expect(breaker).NotTo(Equal(0))
			break
		}

		// Other than available/attached/complete status: See if we the timeout
		// value has been reached. If so, exit with failure after 2 min timer (24x5sec).
		if timer > 24 {
			// timed out fail if resourse is not available.
			Expect(resourceStatus).To(Equal("non-available"))
			break
		} else {
			// Still exists, wait 5 or 10 secs
			time.Sleep(time.Duration(delay) * time.Second)
			timer = timer + 1
		}
	}
}
