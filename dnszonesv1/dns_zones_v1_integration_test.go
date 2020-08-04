/*
 * (C) Copyright IBM Corp. 2020.
 */

package dnszonesv1_test

import (
	"fmt"
	"os"
	"strings"

	"github.com/IBM/go-sdk-core/core"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/IBM/networking-go-sdk/dnszonesv1"
)

const configFile = "../pdns.env"

var configLoaded bool = true

func shouldSkipTest() {
	if !configLoaded {
		Skip("External configuration is not available, skipping...")
	}
}

var _ = Describe(`dnszonesv1`, func() {
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
	zoneOptions := &DnsZonesV1Options{
		ServiceName:   "pdns_services",
		URL:           serviceURL,
		Authenticator: authenticator,
	}
	service, serviceErr := NewDnsZonesV1(zoneOptions)
	if serviceErr != nil {
		fmt.Println(serviceErr)
	}

	Describe(`dnszonesv1`, func() {
		Context(`dnszonesv1`, func() {
			BeforeEach(func() {
				shouldSkipTest()
				listOptions := service.NewListDnszonesOptions(instanceID)
				listResult, listResp, listErr := service.ListDnszones(listOptions)
				Expect(listErr).To(BeNil())
				Expect(listResp).ToNot(BeNil())
				Expect(listResult).ToNot(BeNil())
				Expect(listResp.GetStatusCode()).To(BeEquivalentTo(200))
				for _, zone := range listResult.Dnszones {
					if strings.Contains(*zone.Name, "zone-example") {
						option := service.NewDeleteDnszoneOptions(instanceID, *zone.ID)
						response, err := service.DeleteDnszone(option)
						Expect(err).To(BeNil())
						Expect(response).ToNot(BeNil())
						Expect(response.GetStatusCode()).To(BeEquivalentTo(204))
					}
				}
			})
			AfterEach(func() {
				shouldSkipTest()
				listOptions := service.NewListDnszonesOptions(instanceID)
				listResult, listResp, listErr := service.ListDnszones(listOptions)
				Expect(listErr).To(BeNil())
				Expect(listResp).ToNot(BeNil())
				Expect(listResult).ToNot(BeNil())
				Expect(listResp.GetStatusCode()).To(BeEquivalentTo(200))
				for _, zone := range listResult.Dnszones {
					if strings.Contains(*zone.Name, "zone-example") {
						option := service.NewDeleteDnszoneOptions(instanceID, *zone.ID)
						response, err := service.DeleteDnszone(option)
						Expect(err).To(BeNil())
						Expect(response).ToNot(BeNil())
						Expect(response.GetStatusCode()).To(BeEquivalentTo(204))
					}
				}
			})
			It(`crate/update/delete/get pdns zones`, func() {
				shouldSkipTest()
				// Create DNS Zone
				zoneName := fmt.Sprintf("zone-example%s.com", uuid.New().String())
				createDnszoneOptions := service.NewCreateDnszoneOptions(instanceID)
				createDnszoneOptions.SetName(zoneName)
				createDnszoneOptions.SetDescription("testString")
				createDnszoneOptions.SetLabel("testString")
				createDnszoneOptions.SetXCorrelationID("abc123")
				result, response, reqErr := service.CreateDnszone(createDnszoneOptions)
				Expect(reqErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
				Expect(response.GetStatusCode()).To(BeEquivalentTo(200))

				updateOptions := service.NewUpdateDnszoneOptions(instanceID, *result.ID)
				updateOptions.SetDescription("description")
				updateOptions.SetLabel("test-label")
				updateOptions.SetXCorrelationID("abc123")
				updateResult, updateResponse, updateErr := service.UpdateDnszone(updateOptions)
				Expect(updateErr).To(BeNil())
				Expect(updateResponse).ToNot(BeNil())
				Expect(updateResult).ToNot(BeNil())
				Expect(updateResponse.GetStatusCode()).To(BeEquivalentTo(200))

				getOptions := service.NewGetDnszoneOptions(instanceID, *result.ID)
				getResult, getResponse, getErr := service.GetDnszone(getOptions)
				Expect(getErr).To(BeNil())
				Expect(getResponse).ToNot(BeNil())
				Expect(getResult).ToNot(BeNil())
				Expect(getResponse.GetStatusCode()).To(BeEquivalentTo(200))

				deleteOptions := service.NewDeleteDnszoneOptions(instanceID, *result.ID)
				deleteResponse, deleteErr := service.DeleteDnszone(deleteOptions)
				Expect(deleteErr).To(BeNil())
				Expect(deleteResponse).ToNot(BeNil())
				Expect(deleteResponse.GetStatusCode()).To(BeEquivalentTo(204))
			})
			It(`list pdns zones`, func() {
				shouldSkipTest()
				// Create DNS Zone
				for i := 1; i < 10; i++ {
					zoneName := fmt.Sprintf("zone-example-%s.com", uuid.New().String())
					createDnszoneOptions := service.NewCreateDnszoneOptions(instanceID)
					createDnszoneOptions.SetName(zoneName)
					createDnszoneOptions.SetDescription("testString")
					createDnszoneOptions.SetLabel("testString")
					createDnszoneOptions.SetXCorrelationID("abc123")
					result, response, reqErr := service.CreateDnszone(createDnszoneOptions)
					Expect(reqErr).To(BeNil())
					Expect(response).ToNot(BeNil())
					Expect(result).ToNot(BeNil())
					Expect(response.GetStatusCode()).To(BeEquivalentTo(200))
				}
				// list dns zone with page limit
				listOptions := service.NewListDnszonesOptions(instanceID)
				listOptions.SetLimit(3)
				listOptions.SetOffset(2)
				listResult, listResp, listErr := service.ListDnszones(listOptions)
				Expect(listErr).To(BeNil())
				Expect(listResp).ToNot(BeNil())
				Expect(listResult).ToNot(BeNil())
				Expect(listResp.GetStatusCode()).To(BeEquivalentTo(200))
			})
		})
	})
})
