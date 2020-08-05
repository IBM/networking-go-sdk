package permittednetworksfordnszonesv1_test

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/IBM/go-sdk-core/core"
	guuid "github.com/google/uuid"
	"github.com/joho/godotenv"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/IBM/networking-go-sdk/dnszonesv1"
	. "github.com/IBM/networking-go-sdk/permittednetworksfordnszonesv1"
)

const configFile = "../pdns.env"

var configLoaded bool = true

func shouldSkipTest() {
	if !configLoaded {
		Skip("External configuration is not available, skipping...")
	}
}

var _ = Describe(`permittednetworksfordnszonesv1`, func() {
	if _, err := os.Stat(configFile); err != nil {
		configLoaded = false
	}

	err := godotenv.Load(configFile)
	if err != nil {
		configLoaded = false
	}

	authenticator := &core.IamAuthenticator{
		ApiKey: os.Getenv("PDNS_SERVICES_APIKEY"),
		URL:    os.Getenv("PDNS_SERVICES_AUTH_URL"),
	}
	serviceURL := os.Getenv("API_ENDPOINT")
	instanceID := os.Getenv("INSTANCE_ID")
	vpcCrn := os.Getenv("VPC_CRN")
	globalOptions := &PermittedNetworksForDnsZonesV1Options{
		ServiceName:   "pdns_services",
		URL:           serviceURL,
		Authenticator: authenticator,
	}

	service, serviceErr := NewPermittedNetworksForDnsZonesV1(globalOptions)
	if serviceErr != nil {
		fmt.Println(serviceErr)
	}

	zoneOptions := &dnszonesv1.DnsZonesV1Options{
		ServiceName:   "pdns_services",
		URL:           serviceURL,
		Authenticator: authenticator,
	}
	zoneService, zoneServiceErr := dnszonesv1.NewDnsZonesV1(zoneOptions)
	if zoneServiceErr != nil {
		fmt.Println(zoneServiceErr)
	}
	Describe(`permittednetworksfordnszonesv1`, func() {
		Context(`permittednetworksfordnszonesv1`, func() {
			var zoneInfo *dnszonesv1.Dnszone
			BeforeEach(func() {
				shouldSkipTest()

				// delete all dns zones
				listOptions := zoneService.NewListDnszonesOptions(instanceID)
				listResult, listResp, listErr := zoneService.ListDnszones(listOptions)
				Expect(listErr).To(BeNil())
				Expect(listResp).ToNot(BeNil())
				Expect(listResult).ToNot(BeNil())
				Expect(listResp.GetStatusCode()).To(BeEquivalentTo(200))
				for _, zone := range listResult.Dnszones {
					if strings.Contains(*zone.Name, "network-example") {

						listPermittedNetworksOptions := service.NewListPermittedNetworksOptions(instanceID, *zone.ID)
						results, response, reqErr := service.ListPermittedNetworks(listPermittedNetworksOptions)
						Expect(reqErr).To(BeNil())
						Expect(response).ToNot(BeNil())
						Expect(results).ToNot(BeNil())
						Expect(response.GetStatusCode()).To(BeEquivalentTo(200))

						for _, nw := range results.PermittedNetworks {

							deletePermittedNetworkOptions := service.NewDeletePermittedNetworkOptions(instanceID, *zone.ID, *nw.ID)
							results, response, reqErr := service.DeletePermittedNetwork(deletePermittedNetworkOptions)
							if reqErr != nil {
								fmt.Printf("sleeping for 5 min, since permitted network deletion will take 5 mins, time is %s", time.Now().String())
								time.Sleep(time.Minute * 5)
							} else {
								Expect(reqErr).To(BeNil())
								Expect(response).ToNot(BeNil())
								Expect(results).ToNot(BeNil())
							}
						}

						option := zoneService.NewDeleteDnszoneOptions(instanceID, *zone.ID)
						response, err := zoneService.DeleteDnszone(option)
						Expect(err).To(BeNil())
						Expect(response).ToNot(BeNil())
						Expect(response.GetStatusCode()).To(BeEquivalentTo(204))
					}
				}

				// Create DNS Zone
				zoneName := fmt.Sprintf("network-example-%s.com", guuid.New().String())
				createDnszoneOptions := zoneService.NewCreateDnszoneOptions(instanceID)
				createDnszoneOptions.SetName(zoneName)
				createDnszoneOptions.SetDescription("testString")
				createDnszoneOptions.SetLabel("testString")
				result, response, reqErr := zoneService.CreateDnszone(createDnszoneOptions)
				Expect(reqErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
				Expect(response.GetStatusCode()).To(BeEquivalentTo(200))
				zoneInfo = result
			})
			AfterEach(func() {
				shouldSkipTest()
				// delete all dns zones
				listOptions := zoneService.NewListDnszonesOptions(instanceID)
				listResult, listResp, listErr := zoneService.ListDnszones(listOptions)
				Expect(listErr).To(BeNil())
				Expect(listResp).ToNot(BeNil())
				Expect(listResult).ToNot(BeNil())
				Expect(listResp.GetStatusCode()).To(BeEquivalentTo(200))
				for _, zone := range listResult.Dnszones {
					if strings.Contains(*zone.Name, "network-example") {
						listPermittedNetworksOptions := service.NewListPermittedNetworksOptions(instanceID, *zone.ID)
						results, response, reqErr := service.ListPermittedNetworks(listPermittedNetworksOptions)
						Expect(reqErr).To(BeNil())
						Expect(response).ToNot(BeNil())
						Expect(results).ToNot(BeNil())
						Expect(response.GetStatusCode()).To(BeEquivalentTo(200))

						for _, nw := range results.PermittedNetworks {
							deletePermittedNetworkOptions := service.NewDeletePermittedNetworkOptions(instanceID, *zone.ID, *nw.ID)
							results, response, reqErr := service.DeletePermittedNetwork(deletePermittedNetworkOptions)
							if reqErr != nil {
								fmt.Printf("sleeping for 5 min, since permitted network deletion will take 5 mins, time now %s.", time.Now().String())
								time.Sleep(time.Minute * 5)
								fmt.Printf("exited sleep at %s.", time.Now().String())
							} else {
								Expect(reqErr).To(BeNil())
								Expect(response).ToNot(BeNil())
								Expect(results).ToNot(BeNil())
							}
						}

						option := zoneService.NewDeleteDnszoneOptions(instanceID, *zone.ID)
						response, err := zoneService.DeleteDnszone(option)
						Expect(err).To(BeNil())
						Expect(response).ToNot(BeNil())
						Expect(response.GetStatusCode()).To(BeEquivalentTo(204))
					}
				}
			})
			It(`create/update/delete/get permitted networks`, func() {
				shouldSkipTest()

				header := map[string]string{
					"test": "teststring",
				}
				// Test Add Permitted Network
				createPermittedNetworkOptions := service.NewCreatePermittedNetworkOptions(instanceID, *zoneInfo.ID)
				permittedNetworkCrn := &PermittedNetworkVpc{
					VpcCrn: &vpcCrn,
				}
				createPermittedNetworkOptions.SetPermittedNetwork(permittedNetworkCrn)
				createPermittedNetworkOptions.SetType(CreatePermittedNetworkOptions_Type_Vpc)
				createPermittedNetworkOptions.SetHeaders(header)
				result, response, reqErr := service.CreatePermittedNetwork(createPermittedNetworkOptions)
				Expect(reqErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
				Expect(response.GetStatusCode()).To(BeEquivalentTo(200))
				Expect(*result.State).To(BeEquivalentTo(PermittedNetwork_State_Active))

				permittednetworkID := result.ID

				// Test List Permitted Networks
				listPermittedNetworksOptions := service.NewListPermittedNetworksOptions(instanceID, *zoneInfo.ID)
				listPermittedNetworksOptions.SetHeaders(header)
				results, response, reqErr := service.ListPermittedNetworks(listPermittedNetworksOptions)
				Expect(reqErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
				Expect(response.GetStatusCode()).To(BeEquivalentTo(200))
				firstResource := results.PermittedNetworks[0]
				Expect(*firstResource.ID).ToNot(BeNil())

				// Test Get Permitted Network
				getPermittedNetworkOptions := service.NewGetPermittedNetworkOptions(instanceID, *zoneInfo.ID, *permittednetworkID)
				getPermittedNetworkOptions.SetHeaders(header)
				result, response, reqErr = service.GetPermittedNetwork(getPermittedNetworkOptions)
				Expect(reqErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
				Expect(response.GetStatusCode()).To(BeEquivalentTo(200))
				Expect(*result.State).To(BeEquivalentTo(PermittedNetwork_State_Active))

				// Test Get Permitted Network Fail
				fgetPermittedNetworkOptions := new(GetPermittedNetworkOptions)
				fgetPermittedNetworkOptions.SetInstanceID(instanceID)
				fgetPermittedNetworkOptions.SetDnszoneID(*zoneInfo.ID)
				fgetPermittedNetworkOptions.SetPermittedNetworkID("invalid_id")
				_, _, reqErr = service.GetPermittedNetwork(fgetPermittedNetworkOptions)
				Expect(reqErr).ToNot(BeNil())

				// Test Remove Permitted Network
				deletePermittedNetworkOptions := service.NewDeletePermittedNetworkOptions(instanceID, *zoneInfo.ID, *permittednetworkID)
				deletePermittedNetworkOptions.SetHeaders(header)
				result, response, reqErr = service.DeletePermittedNetwork(deletePermittedNetworkOptions)
				Expect(reqErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
				Expect(response.GetStatusCode()).To(BeEquivalentTo(202))
				Expect(*result.State).To(BeEquivalentTo(PermittedNetwork_State_RemovalInProgress))

				// Test Rmove Permitted Network Fail
				fdeletePermittedNetworkOptions := new(DeletePermittedNetworkOptions)
				fdeletePermittedNetworkOptions.SetInstanceID(instanceID)
				fdeletePermittedNetworkOptions.SetDnszoneID(*zoneInfo.ID)
				fdeletePermittedNetworkOptions.SetPermittedNetworkID("invalid_id")
				_, _, reqErr = service.DeletePermittedNetwork(fdeletePermittedNetworkOptions)
				Expect(reqErr).ToNot(BeNil())
			})
		})
	})
})
