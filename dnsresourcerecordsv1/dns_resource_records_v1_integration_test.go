/*
 * (C) Copyright IBM Corp. 2020.
 */

package dnsresourcerecordsv1_test

import (
	"fmt"
	"os"
	"strings"

	"github.com/IBM/go-sdk-core/core"
	. "github.com/IBM/networking-go-sdk/dnsresourcerecordsv1"
	"github.com/IBM/networking-go-sdk/dnszonesv1"
	guuid "github.com/google/uuid"
	"github.com/joho/godotenv"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const configFile = "../pdns.env"

var configLoaded bool = true

func shouldSkipTest() {
	if !configLoaded {
		Skip("External configuration is not available, skipping...")
	}
}

var _ = Describe(`dnsresourcerecordsv1`, func() {
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
	dnsRecordOptions := &DnsResourceRecordsV1Options{
		ServiceName:   "pdns_services",
		URL:           serviceURL,
		Authenticator: authenticator,
	}

	zoneOptions := &dnszonesv1.DnsZonesV1Options{
		ServiceName:   "pdns_services",
		URL:           serviceURL,
		Authenticator: authenticator,
	}

	service, serviceErr := NewDnsResourceRecordsV1(dnsRecordOptions)
	if serviceErr != nil {
		fmt.Println(serviceErr)
	}

	zoneService, zoneServiceErr := dnszonesv1.NewDnsZonesV1(zoneOptions)
	if zoneServiceErr != nil {
		fmt.Println(zoneServiceErr)
	}
	Describe(`resourcerecordsv1`, func() {
		Context(`resourcerecordsv1`, func() {
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
					if strings.Contains(*zone.Name, "test-example") {
						// delete all dns zones
						listOptions := service.NewListResourceRecordsOptions(instanceID, *zone.ID)
						listResult, listResp, listErr := service.ListResourceRecords(listOptions)
						Expect(listErr).To(BeNil())
						Expect(listResp).ToNot(BeNil())
						Expect(listResult).ToNot(BeNil())

						for _, record := range listResult.ResourceRecords {
							deleteOpt := service.NewDeleteResourceRecordOptions(instanceID, *zone.ID, *record.ID)
							deleteResponse, deleteErr := service.DeleteResourceRecord(deleteOpt)
							Expect(deleteErr).To(BeNil())
							Expect(deleteResponse).ToNot(BeNil())
						}

						// delete zone
						option := zoneService.NewDeleteDnszoneOptions(instanceID, *zone.ID)
						response, err := zoneService.DeleteDnszone(option)
						Expect(err).To(BeNil())
						Expect(response).ToNot(BeNil())
						Expect(response.GetStatusCode()).To(BeEquivalentTo(204))
					}
				}

				// Create DNS Zone
				zoneName := fmt.Sprintf("test-example%s.com", guuid.New().String())
				createDnszoneOptions := zoneService.NewCreateDnszoneOptions(instanceID)
				createDnszoneOptions.SetName(zoneName)
				createDnszoneOptions.SetDescription("testString")
				createDnszoneOptions.SetLabel("testString")
				createDnszoneOptions.SetXCorrelationID("abc123")
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
					if strings.Contains(*zone.Name, "test-example") {
						// delete all dns zones
						listOptions := service.NewListResourceRecordsOptions(instanceID, *zone.ID)
						listResult, listResp, listErr := service.ListResourceRecords(listOptions)
						Expect(listErr).To(BeNil())
						Expect(listResp).ToNot(BeNil())
						Expect(listResult).ToNot(BeNil())

						for _, record := range listResult.ResourceRecords {
							deleteOpt := service.NewDeleteResourceRecordOptions(instanceID, *zone.ID, *record.ID)
							deleteResponse, deleteErr := service.DeleteResourceRecord(deleteOpt)
							Expect(deleteErr).To(BeNil())
							Expect(deleteResponse).ToNot(BeNil())
						}

						// delete zone
						option := zoneService.NewDeleteDnszoneOptions(instanceID, *zone.ID)
						response, err := zoneService.DeleteDnszone(option)
						Expect(err).To(BeNil())
						Expect(response).ToNot(BeNil())
						Expect(response.GetStatusCode()).To(BeEquivalentTo(204))
					}
				}
			})
			It(`crate/update/delete/get pdns A records`, func() {
				shouldSkipTest()

				// create resource record
				createResourceRecordOptions := service.NewCreateResourceRecordOptions(instanceID, *zoneInfo.ID)
				createResourceRecordOptions.SetName("testa")
				createResourceRecordOptions.SetType(CreateResourceRecordOptions_Type_A)
				createResourceRecordOptions.SetTTL(120)
				rdataARecord, err := service.NewResourceRecordInputRdataRdataARecord("1.1.1.1")
				Expect(err).To(BeNil())
				createResourceRecordOptions.SetRdata(rdataARecord)
				createResourceRecordOptions.SetXCorrelationID("abc123")
				result, response, reqErr := service.CreateResourceRecord(createResourceRecordOptions)
				Expect(reqErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
				Expect(response.GetStatusCode()).To(BeEquivalentTo(200))
				Expect(*result.Type).To(BeEquivalentTo(CreateResourceRecordOptions_Type_A))

				updateOpt := service.NewUpdateResourceRecordOptions(instanceID, *zoneInfo.ID, *result.ID)
				updateOpt.SetName("updatea")
				updateOpt.SetTTL(300)
				updateOpt.SetXCorrelationID("abc123")
				updaterdataARecord, err := service.NewResourceRecordUpdateInputRdataRdataARecord("1.1.1.2")
				Expect(err).To(BeNil())
				updateOpt.SetRdata(updaterdataARecord)
				updateResult, updateResponse, updateErr := service.UpdateResourceRecord(updateOpt)
				Expect(updateErr).To(BeNil())
				Expect(updateResponse).ToNot(BeNil())
				Expect(updateResult).ToNot(BeNil())
				Expect(updateResponse.GetStatusCode()).To(BeEquivalentTo(200))
				Expect(*updateResult.Type).To(BeEquivalentTo(CreateResourceRecordOptions_Type_A))

				getOpt := service.NewGetResourceRecordOptions(instanceID, *zoneInfo.ID, *result.ID)
				getOpt.SetXCorrelationID("abc123")
				getResult, getResponse, getErr := service.GetResourceRecord(getOpt)
				Expect(getErr).To(BeNil())
				Expect(getResponse).ToNot(BeNil())
				Expect(getResult).ToNot(BeNil())
				Expect(getResponse.GetStatusCode()).To(BeEquivalentTo(200))
				Expect(*getResult.Type).To(BeEquivalentTo(CreateResourceRecordOptions_Type_A))

				deleteOpt := service.NewDeleteResourceRecordOptions(instanceID, *zoneInfo.ID, *result.ID)
				deleteOpt.SetXCorrelationID("abc123")
				deleteResponse, deleteErr := service.DeleteResourceRecord(deleteOpt)
				Expect(deleteErr).To(BeNil())
				Expect(deleteResponse).ToNot(BeNil())
				Expect(deleteResponse.GetStatusCode()).To(BeEquivalentTo(204))
			})

			It(`crate/update/delete/get pdns PTR records`, func() {
				shouldSkipTest()

				// create resource record
				createResourceRecordOptions := service.NewCreateResourceRecordOptions(instanceID, *zoneInfo.ID)
				createResourceRecordOptions.SetName("testa")
				createResourceRecordOptions.SetType(CreateResourceRecordOptions_Type_A)
				createResourceRecordOptions.SetTTL(120)
				rdataARecord, err := service.NewResourceRecordInputRdataRdataARecord("1.1.1.1")
				Expect(err).To(BeNil())
				createResourceRecordOptions.SetRdata(rdataARecord)
				createResourceRecordOptions.SetXCorrelationID("abc123")
				result, response, reqErr := service.CreateResourceRecord(createResourceRecordOptions)
				Expect(reqErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
				Expect(response.GetStatusCode()).To(BeEquivalentTo(200))
				Expect(*result.Type).To(BeEquivalentTo(CreateResourceRecordOptions_Type_A))

				createResourcePtrRecordOptions := service.NewCreateResourceRecordOptions(instanceID, *zoneInfo.ID)
				createResourcePtrRecordOptions.SetName("1.1.1.1")
				createResourcePtrRecordOptions.SetType(CreateResourceRecordOptions_Type_Ptr)
				createResourcePtrRecordOptions.SetTTL(120)
				rdataPtrRecord, err := service.NewResourceRecordInputRdataRdataPtrRecord("testa." + *zoneInfo.Name)
				Expect(err).To(BeNil())
				createResourcePtrRecordOptions.SetRdata(rdataPtrRecord)
				createResourcePtrRecordOptions.SetXCorrelationID("abc123")
				ptrResult, ptrResponse, ptrErr := service.CreateResourceRecord(createResourcePtrRecordOptions)
				Expect(ptrErr).To(BeNil())
				Expect(ptrResponse).ToNot(BeNil())
				Expect(ptrResult).ToNot(BeNil())
				Expect(ptrResponse.GetStatusCode()).To(BeEquivalentTo(200))
				Expect(*ptrResult.Type).To(BeEquivalentTo(CreateResourceRecordOptions_Type_Ptr))

				getOpt := service.NewGetResourceRecordOptions(instanceID, *zoneInfo.ID, *ptrResult.ID)
				getOpt.SetXCorrelationID("abc123")
				getResult, getResponse, getErr := service.GetResourceRecord(getOpt)
				Expect(getErr).To(BeNil())
				Expect(getResponse).ToNot(BeNil())
				Expect(getResult).ToNot(BeNil())
				Expect(getResponse.GetStatusCode()).To(BeEquivalentTo(200))
				Expect(*getResult.Type).To(BeEquivalentTo(CreateResourceRecordOptions_Type_Ptr))

				deleteOpt := service.NewDeleteResourceRecordOptions(instanceID, *zoneInfo.ID, *ptrResult.ID)
				deleteOpt.SetXCorrelationID("abc123")
				deleteResponse, deleteErr := service.DeleteResourceRecord(deleteOpt)
				Expect(deleteErr).To(BeNil())
				Expect(deleteResponse).ToNot(BeNil())
				Expect(deleteResponse.GetStatusCode()).To(BeEquivalentTo(204))
			})
			It(`create/update/delete/get pdns AAAA record`, func() {
				shouldSkipTest()

				header := map[string]string{
					"test": "teststring",
				}
				// Test Create Resource Record AAAA
				createResourceRecordOptions := service.NewCreateResourceRecordOptions(instanceID, *zoneInfo.ID)
				createResourceRecordOptions.SetName("testaaaa")
				createResourceRecordOptions.SetType(CreateResourceRecordOptions_Type_Aaaa)
				createResourceRecordOptions.SetTTL(120)
				rdataAaaaRecord, err := service.NewResourceRecordInputRdataRdataAaaaRecord("2001::8888")
				Expect(err).To(BeNil())
				createResourceRecordOptions.SetRdata(rdataAaaaRecord)
				createResourceRecordOptions.SetXCorrelationID("abc123")
				createResourceRecordOptions.SetHeaders(header)
				result, response, reqErr := service.CreateResourceRecord(createResourceRecordOptions)
				Expect(reqErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
				Expect(response.GetStatusCode()).To(BeEquivalentTo(200))
				Expect(*result.Type).To(BeEquivalentTo(CreateResourceRecordOptions_Type_Aaaa))
				Expect(*result.Name).To(BeEquivalentTo("testaaaa." + *zoneInfo.Name))

				aaaaRecordID := *result.ID
				// Test Update Resource Record AAAA
				updateResourceRecordOptions := service.NewUpdateResourceRecordOptions(instanceID, *zoneInfo.ID, aaaaRecordID)
				updateResourceRecordOptions.SetName("updateaaaa")
				updateResourceRecordOptions.SetTTL(300)
				updaterdataAaaaRecord, err := service.NewResourceRecordUpdateInputRdataRdataAaaaRecord("2001::8889")
				Expect(err).To(BeNil())
				updateResourceRecordOptions.SetRdata(updaterdataAaaaRecord)
				updateResourceRecordOptions.SetXCorrelationID("abc123")
				updateResourceRecordOptions.SetHeaders(header)
				result, response, reqErr = service.UpdateResourceRecord(updateResourceRecordOptions)
				Expect(reqErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
				Expect(response.GetStatusCode()).To(BeEquivalentTo(200))
				Expect(*result.Type).To(BeEquivalentTo(CreateResourceRecordOptions_Type_Aaaa))
				Expect(*result.Name).To(BeEquivalentTo("updateaaaa." + *zoneInfo.Name))

				getOpt := service.NewGetResourceRecordOptions(instanceID, *zoneInfo.ID, *result.ID)
				getOpt.SetXCorrelationID("abc123")
				getResult, getResponse, getErr := service.GetResourceRecord(getOpt)
				Expect(getErr).To(BeNil())
				Expect(getResponse).ToNot(BeNil())
				Expect(getResult).ToNot(BeNil())
				Expect(getResponse.GetStatusCode()).To(BeEquivalentTo(200))
				Expect(*getResult.ID).To(BeEquivalentTo(*result.ID))

				// Test Delete Resource Record AAAA
				deleteResourceRecordOptions := service.NewDeleteResourceRecordOptions(instanceID, *zoneInfo.ID, aaaaRecordID)
				deleteResourceRecordOptions.SetXCorrelationID("abc123")
				deleteResourceRecordOptions.SetHeaders(header)
				response, reqErr = service.DeleteResourceRecord(deleteResourceRecordOptions)
				Expect(reqErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
				Expect(response.GetStatusCode()).To(BeEquivalentTo(204))
				Expect(*result.Type).To(BeEquivalentTo(CreateResourceRecordOptions_Type_Aaaa))
			})
			It(`create/update/delete/get pdns CNAME record`, func() {
				shouldSkipTest()

				header := map[string]string{
					"test": "teststring",
				}
				// Test Create Resource Record CNAME
				createResourceRecordOptions := service.NewCreateResourceRecordOptions(instanceID, *zoneInfo.ID)
				createResourceRecordOptions.SetName("testcname")
				createResourceRecordOptions.SetType(CreateResourceRecordOptions_Type_Cname)
				createResourceRecordOptions.SetTTL(120)
				rdataCnameRecord, err := service.NewResourceRecordInputRdataRdataCnameRecord("testcnamedata.com")
				Expect(err).To(BeNil())

				createResourceRecordOptions.SetRdata(rdataCnameRecord)
				createResourceRecordOptions.SetXCorrelationID("abc123")
				createResourceRecordOptions.SetHeaders(header)
				result, response, reqErr := service.CreateResourceRecord(createResourceRecordOptions)
				Expect(reqErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
				Expect(response.GetStatusCode()).To(BeEquivalentTo(200))
				Expect(*result.Type).To(BeEquivalentTo(CreateResourceRecordOptions_Type_Cname))
				Expect(*result.Name).To(BeEquivalentTo("testcname." + *zoneInfo.Name))

				cnameRecordID := result.ID
				// Test Update Resource Record CNAME
				updateResourceRecordOptions := service.NewUpdateResourceRecordOptions(instanceID, *zoneInfo.ID, *cnameRecordID)
				updateResourceRecordOptions.SetName("updatecname")
				updateResourceRecordOptions.SetTTL(300)
				updaterdataCnameRecord, err := service.NewResourceRecordUpdateInputRdataRdataCnameRecord("updatecnamedata.com")
				Expect(err).To(BeNil())
				updateResourceRecordOptions.SetRdata(updaterdataCnameRecord)
				updateResourceRecordOptions.SetXCorrelationID("abc123")
				updateResourceRecordOptions.SetHeaders(header)
				result, response, reqErr = service.UpdateResourceRecord(updateResourceRecordOptions)
				Expect(reqErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
				Expect(response.GetStatusCode()).To(BeEquivalentTo(200))
				Expect(*result.Type).To(BeEquivalentTo(CreateResourceRecordOptions_Type_Cname))
				Expect(*result.Name).To(BeEquivalentTo("updatecname." + *zoneInfo.Name))

				getOpt := service.NewGetResourceRecordOptions(instanceID, *zoneInfo.ID, *result.ID)
				getOpt.SetXCorrelationID("abc123")
				getResult, getResponse, getErr := service.GetResourceRecord(getOpt)
				Expect(getErr).To(BeNil())
				Expect(getResponse).ToNot(BeNil())
				Expect(getResult).ToNot(BeNil())
				Expect(getResponse.GetStatusCode()).To(BeEquivalentTo(200))
				Expect(*getResult.ID).To(BeEquivalentTo(*result.ID))

				// Test Delete Resource Record CNAME
				deleteResourceRecordOptions := service.NewDeleteResourceRecordOptions(instanceID, *zoneInfo.ID, *cnameRecordID)
				deleteResourceRecordOptions.SetXCorrelationID("abc123")
				deleteResourceRecordOptions.SetHeaders(header)
				response, reqErr = service.DeleteResourceRecord(deleteResourceRecordOptions)
				Expect(reqErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
				Expect(response.GetStatusCode()).To(BeEquivalentTo(204))
			})
			It(`create/update/delete/get pdns MX record`, func() {
				shouldSkipTest()

				header := map[string]string{
					"test": "teststring",
				}
				// Test Create Resource Record MX
				createResourceRecordOptions := service.NewCreateResourceRecordOptions(instanceID, *zoneInfo.ID)
				createResourceRecordOptions.SetName("testmx")
				createResourceRecordOptions.SetType(CreateResourceRecordOptions_Type_Mx)
				createResourceRecordOptions.SetTTL(120)
				rdataMxRecord, err := service.NewResourceRecordInputRdataRdataMxRecord("mail.testmx.com", 1)
				Expect(err).To(BeNil())
				createResourceRecordOptions.SetRdata(rdataMxRecord)
				createResourceRecordOptions.SetXCorrelationID("abc123")
				createResourceRecordOptions.SetHeaders(header)
				result, response, reqErr := service.CreateResourceRecord(createResourceRecordOptions)
				Expect(reqErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
				Expect(response.GetStatusCode()).To(BeEquivalentTo(200))
				Expect(*result.Type).To(BeEquivalentTo(CreateResourceRecordOptions_Type_Mx))
				Expect(*result.Name).To(BeEquivalentTo("testmx." + *zoneInfo.Name))

				mxRecordID := result.ID
				// Test Update Resource Record MX
				updateResourceRecordOptions := service.NewUpdateResourceRecordOptions(instanceID, *zoneInfo.ID, *mxRecordID)
				updateResourceRecordOptions.SetName("testupdatemx")
				updateResourceRecordOptions.SetTTL(300)
				updaterdataMxRecord, err := service.NewResourceRecordUpdateInputRdataRdataMxRecord("mail1.testmx.com", 2)
				Expect(err).To(BeNil())
				updateResourceRecordOptions.SetRdata(updaterdataMxRecord)
				updateResourceRecordOptions.SetXCorrelationID("abc123")
				updateResourceRecordOptions.SetHeaders(header)
				result, response, reqErr = service.UpdateResourceRecord(updateResourceRecordOptions)
				Expect(reqErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
				Expect(response.GetStatusCode()).To(BeEquivalentTo(200))
				Expect(*result.Type).To(BeEquivalentTo(CreateResourceRecordOptions_Type_Mx))
				Expect(*result.Name).To(BeEquivalentTo("testupdatemx." + *zoneInfo.Name))

				// Test Delete Resource Record MX
				deleteResourceRecordOptions := service.NewDeleteResourceRecordOptions(instanceID, *zoneInfo.ID, *mxRecordID)
				deleteResourceRecordOptions.SetXCorrelationID("abc123")
				deleteResourceRecordOptions.SetHeaders(header)
				response, reqErr = service.DeleteResourceRecord(deleteResourceRecordOptions)
				Expect(reqErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
				Expect(response.GetStatusCode()).To(BeEquivalentTo(204))
			})
			It(`create/update/delete/get pdns SRV record`, func() {
				shouldSkipTest()

				header := map[string]string{
					"test": "teststring",
				}
				// Test Create Resource Record SRV
				createResourceRecordOptions := service.NewCreateResourceRecordOptions(instanceID, *zoneInfo.ID)
				createResourceRecordOptions.SetName("testsrv")
				createResourceRecordOptions.SetType(CreateResourceRecordOptions_Type_Srv)
				createResourceRecordOptions.SetTTL(120)
				createResourceRecordOptions.SetService("_sip")
				createResourceRecordOptions.SetProtocol("udp")
				rdataSrvRecord, err := service.NewResourceRecordInputRdataRdataSrvRecord(1, 1, "siphost.com", 1)
				Expect(err).To(BeNil())
				createResourceRecordOptions.SetRdata(rdataSrvRecord)
				createResourceRecordOptions.SetXCorrelationID("abc123")
				createResourceRecordOptions.SetHeaders(header)
				result, response, reqErr := service.CreateResourceRecord(createResourceRecordOptions)
				Expect(reqErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
				Expect(response.GetStatusCode()).To(BeEquivalentTo(200))
				Expect(*result.Type).To(BeEquivalentTo(CreateResourceRecordOptions_Type_Srv))
				Expect(*result.Protocol).To(BeEquivalentTo("udp"))

				srvRecordID := result.ID
				// Test Update Resource Record SRV
				updateResourceRecordOptions := service.NewUpdateResourceRecordOptions(instanceID, *zoneInfo.ID, *srvRecordID)
				updateResourceRecordOptions.SetName("updatesrv")
				updateResourceRecordOptions.SetTTL(300)
				updateResourceRecordOptions.SetService("_sip")
				updateResourceRecordOptions.SetProtocol("udp")
				updaterdataSrvRecord, err := service.NewResourceRecordUpdateInputRdataRdataSrvRecord(2, 2, "updatesiphost.com", 2)
				Expect(err).To(BeNil())
				updateResourceRecordOptions.SetRdata(updaterdataSrvRecord)
				updateResourceRecordOptions.SetXCorrelationID("abc123")
				updateResourceRecordOptions.SetHeaders(header)
				result, response, reqErr = service.UpdateResourceRecord(updateResourceRecordOptions)
				Expect(reqErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
				Expect(response.GetStatusCode()).To(BeEquivalentTo(200))
				Expect(*result.Type).To(BeEquivalentTo(CreateResourceRecordOptions_Type_Srv))
				Expect(*result.Protocol).To(BeEquivalentTo("udp"))

				// Test Delete Resource Record SRV
				deleteResourceRecordOptions := service.NewDeleteResourceRecordOptions(instanceID, *zoneInfo.ID, *srvRecordID)
				deleteResourceRecordOptions.SetXCorrelationID("abc123")
				deleteResourceRecordOptions.SetHeaders(header)
				response, reqErr = service.DeleteResourceRecord(deleteResourceRecordOptions)
				Expect(reqErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
				Expect(response.GetStatusCode()).To(BeEquivalentTo(204))
			})
			It(`create/update/delete/get pdns TXT record`, func() {
				shouldSkipTest()

				header := map[string]string{
					"test": "teststring",
				}
				// Test Create Resource Record TXT
				createResourceRecordOptions := service.NewCreateResourceRecordOptions(instanceID, *zoneInfo.ID)
				createResourceRecordOptions.SetName("testtxt")
				createResourceRecordOptions.SetType(CreateResourceRecordOptions_Type_Txt)
				createResourceRecordOptions.SetTTL(120)
				rdataTxtRecord, err := service.NewResourceRecordInputRdataRdataTxtRecord("txtdata string")
				Expect(err).To(BeNil())
				createResourceRecordOptions.SetRdata(rdataTxtRecord)
				createResourceRecordOptions.SetXCorrelationID("abc123")
				createResourceRecordOptions.SetHeaders(header)
				result, response, reqErr := service.CreateResourceRecord(createResourceRecordOptions)
				Expect(reqErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
				Expect(response.GetStatusCode()).To(BeEquivalentTo(200))
				Expect(*result.Type).To(BeEquivalentTo(CreateResourceRecordOptions_Type_Txt))
				Expect(*result.Name).To(BeEquivalentTo("testtxt." + *zoneInfo.Name))

				txtRecordID := result.ID
				// Test Update Resource Record TXT
				updateResourceRecordOptions := &UpdateResourceRecordOptions{}
				updateResourceRecordOptions.SetInstanceID(instanceID)
				updateResourceRecordOptions.SetDnszoneID(*zoneInfo.ID)
				updateResourceRecordOptions.SetRecordID(*txtRecordID)
				updateResourceRecordOptions.SetTTL(300)
				updateResourceRecordOptions.SetName("updatetxt")
				updaterdataTxtRecord, err := service.NewResourceRecordUpdateInputRdataRdataTxtRecord("update txtdata string")
				Expect(err).To(BeNil())
				updateResourceRecordOptions.SetRdata(updaterdataTxtRecord)
				updateResourceRecordOptions.SetXCorrelationID("abc123")
				updateResourceRecordOptions.SetHeaders(header)
				result, response, reqErr = service.UpdateResourceRecord(updateResourceRecordOptions)
				Expect(reqErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
				Expect(response.GetStatusCode()).To(BeEquivalentTo(200))
				Expect(*result.Type).To(BeEquivalentTo(CreateResourceRecordOptions_Type_Txt))
				Expect(*result.Name).To(BeEquivalentTo("updatetxt." + *zoneInfo.Name))

				// Test Delete Resource Record TXT
				deleteResourceRecordOptions := service.NewDeleteResourceRecordOptions(instanceID, *zoneInfo.ID, *txtRecordID)
				deleteResourceRecordOptions.SetXCorrelationID("abc123")
				deleteResourceRecordOptions.SetHeaders(header)
				response, reqErr = service.DeleteResourceRecord(deleteResourceRecordOptions)
				Expect(reqErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
				Expect(response.GetStatusCode()).To(BeEquivalentTo(204))
			})
			It(`crate/update/delete/get pdns A records`, func() {
				shouldSkipTest()
				for i := 1; i < 10; i++ {
					name := fmt.Sprintf("testa%d", i)
					ip := fmt.Sprintf("1.1.1.%d", i)
					// create resource record
					createResourceRecordOptions := service.NewCreateResourceRecordOptions(instanceID, *zoneInfo.ID)
					createResourceRecordOptions.SetName(name)
					createResourceRecordOptions.SetType(CreateResourceRecordOptions_Type_A)
					createResourceRecordOptions.SetTTL(120)
					rdataARecord, err := service.NewResourceRecordInputRdataRdataARecord(ip)
					Expect(err).To(BeNil())
					createResourceRecordOptions.SetRdata(rdataARecord)
					createResourceRecordOptions.SetXCorrelationID("abc123")
					result, response, reqErr := service.CreateResourceRecord(createResourceRecordOptions)
					Expect(reqErr).To(BeNil())
					Expect(response).ToNot(BeNil())
					Expect(result).ToNot(BeNil())
					Expect(response.GetStatusCode()).To(BeEquivalentTo(200))
				}
				// delete all dns zones
				listOptions := service.NewListResourceRecordsOptions(instanceID, *zoneInfo.ID)
				listOptions.SetLimit(3)
				listOptions.SetOffset(3)
				listResult, listResp, listErr := service.ListResourceRecords(listOptions)
				Expect(listErr).To(BeNil())
				Expect(listResp).ToNot(BeNil())
				Expect(listResult).ToNot(BeNil())
				Expect(listResp.GetStatusCode()).To(BeEquivalentTo(200))
			})
		})
	})
})
