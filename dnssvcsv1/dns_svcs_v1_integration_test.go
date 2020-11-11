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

package dnssvcsv1_test

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/IBM/go-sdk-core/v4/core"
	"github.com/IBM/networking-go-sdk/dnssvcsv1"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const configFile = "../dns.env"

var configLoaded bool = true

func shouldSkipTest() {
	if !configLoaded {
		Skip("External configuration is not available, skipping...")
	}
}

var _ = Describe(`dnssvcsv1`, func() {
	if _, err := os.Stat(configFile); err != nil {
		configLoaded = false
	}

	err := godotenv.Load(configFile)
	if err != nil {
		configLoaded = false
	}

	authenticator, err := core.GetAuthenticatorFromEnvironment("dns_svcs")
	if err != nil {
		panic(err)
	}
	options := &dnssvcsv1.DnsSvcsV1Options{
		Authenticator: authenticator,
	}
	service, serviceErr := dnssvcsv1.NewDnsSvcsV1UsingExternalConfig(options)
	if serviceErr != nil {
		panic(err)
	}

	instanceID := os.Getenv("DNS_SVCS_INSTANCE_ID")
	vpcCrn := os.Getenv("DNS_SVCS_VPC_CRN")

	Describe(`dnssvcsv1`, func() {
		Context(`dnssvcsv1`, func() {
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
			It(`create/update/delete/get pdns zones`, func() {
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

				// list dns zone with page limit
				listOptions := service.NewListDnszonesOptions(instanceID)
				listOptions.SetLimit(2)
				listOptions.SetOffset(2)
				listResult, listResp, listErr := service.ListDnszones(listOptions)
				Expect(listErr).To(BeNil())
				Expect(listResp).ToNot(BeNil())
				Expect(listResult).ToNot(BeNil())
				Expect(listResp.GetStatusCode()).To(BeEquivalentTo(200))

				deleteOptions := service.NewDeleteDnszoneOptions(instanceID, *result.ID)
				deleteResponse, deleteErr := service.DeleteDnszone(deleteOptions)
				Expect(deleteErr).To(BeNil())
				Expect(deleteResponse).ToNot(BeNil())
				Expect(deleteResponse.GetStatusCode()).To(BeEquivalentTo(204))
			})
		})
	})
	Describe(`resourcerecordsv1`, func() {
		Context(`resourcerecordsv1`, func() {
			var zoneInfo *dnssvcsv1.Dnszone
			BeforeEach(func() {
				shouldSkipTest()

				// delete all dns zones
				listOptions := service.NewListDnszonesOptions(instanceID)
				listResult, listResp, listErr := service.ListDnszones(listOptions)
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
						option := service.NewDeleteDnszoneOptions(instanceID, *zone.ID)
						response, err := service.DeleteDnszone(option)
						Expect(err).To(BeNil())
						Expect(response).ToNot(BeNil())
						Expect(response.GetStatusCode()).To(BeEquivalentTo(204))
					}
				}

				// Create DNS Zone
				zoneName := fmt.Sprintf("test-example%s.com", uuid.New().String())
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
				zoneInfo = result
			})
			AfterEach(func() {
				shouldSkipTest()

				// delete all dns zones
				listOptions := service.NewListDnszonesOptions(instanceID)
				listResult, listResp, listErr := service.ListDnszones(listOptions)
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
						option := service.NewDeleteDnszoneOptions(instanceID, *zone.ID)
						response, err := service.DeleteDnszone(option)
						Expect(err).To(BeNil())
						Expect(response).ToNot(BeNil())
						Expect(response.GetStatusCode()).To(BeEquivalentTo(204))
					}
				}
			})
			It(`create/update/delete/get pdns A records`, func() {
				shouldSkipTest()

				// create resource record
				createResourceRecordOptions := service.NewCreateResourceRecordOptions(instanceID, *zoneInfo.ID)
				createResourceRecordOptions.SetName("testa")
				createResourceRecordOptions.SetType(dnssvcsv1.CreateResourceRecordOptions_Type_A)
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
				Expect(*result.Type).To(BeEquivalentTo(dnssvcsv1.CreateResourceRecordOptions_Type_A))

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
				Expect(*updateResult.Type).To(BeEquivalentTo(dnssvcsv1.CreateResourceRecordOptions_Type_A))

				getOpt := service.NewGetResourceRecordOptions(instanceID, *zoneInfo.ID, *result.ID)
				getOpt.SetXCorrelationID("abc123")
				getResult, getResponse, getErr := service.GetResourceRecord(getOpt)
				Expect(getErr).To(BeNil())
				Expect(getResponse).ToNot(BeNil())
				Expect(getResult).ToNot(BeNil())
				Expect(getResponse.GetStatusCode()).To(BeEquivalentTo(200))
				Expect(*getResult.Type).To(BeEquivalentTo(dnssvcsv1.CreateResourceRecordOptions_Type_A))

				deleteOpt := service.NewDeleteResourceRecordOptions(instanceID, *zoneInfo.ID, *result.ID)
				deleteOpt.SetXCorrelationID("abc123")
				deleteResponse, deleteErr := service.DeleteResourceRecord(deleteOpt)
				Expect(deleteErr).To(BeNil())
				Expect(deleteResponse).ToNot(BeNil())
				Expect(deleteResponse.GetStatusCode()).To(BeEquivalentTo(204))
			})

			It(`create/update/delete/get pdns PTR records`, func() {
				shouldSkipTest()

				// create resource record
				createResourceRecordOptions := service.NewCreateResourceRecordOptions(instanceID, *zoneInfo.ID)
				createResourceRecordOptions.SetName("testa")
				createResourceRecordOptions.SetType(dnssvcsv1.CreateResourceRecordOptions_Type_A)
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
				Expect(*result.Type).To(BeEquivalentTo(dnssvcsv1.CreateResourceRecordOptions_Type_A))

				createResourcePtrRecordOptions := service.NewCreateResourceRecordOptions(instanceID, *zoneInfo.ID)
				createResourcePtrRecordOptions.SetName("1.1.1.1")
				createResourcePtrRecordOptions.SetType(dnssvcsv1.CreateResourceRecordOptions_Type_Ptr)
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
				Expect(*ptrResult.Type).To(BeEquivalentTo(dnssvcsv1.CreateResourceRecordOptions_Type_Ptr))

				getOpt := service.NewGetResourceRecordOptions(instanceID, *zoneInfo.ID, *ptrResult.ID)
				getOpt.SetXCorrelationID("abc123")
				getResult, getResponse, getErr := service.GetResourceRecord(getOpt)
				Expect(getErr).To(BeNil())
				Expect(getResponse).ToNot(BeNil())
				Expect(getResult).ToNot(BeNil())
				Expect(getResponse.GetStatusCode()).To(BeEquivalentTo(200))
				Expect(*getResult.Type).To(BeEquivalentTo(dnssvcsv1.CreateResourceRecordOptions_Type_Ptr))

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
				createResourceRecordOptions.SetType(dnssvcsv1.CreateResourceRecordOptions_Type_Aaaa)
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
				Expect(*result.Type).To(BeEquivalentTo(dnssvcsv1.CreateResourceRecordOptions_Type_Aaaa))
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
				Expect(*result.Type).To(BeEquivalentTo(dnssvcsv1.CreateResourceRecordOptions_Type_Aaaa))
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
				Expect(*result.Type).To(BeEquivalentTo(dnssvcsv1.CreateResourceRecordOptions_Type_Aaaa))
			})
			It(`create/update/delete/get pdns CNAME record`, func() {
				shouldSkipTest()

				header := map[string]string{
					"test": "teststring",
				}
				// Test Create Resource Record CNAME
				createResourceRecordOptions := service.NewCreateResourceRecordOptions(instanceID, *zoneInfo.ID)
				createResourceRecordOptions.SetName("testcname")
				createResourceRecordOptions.SetType(dnssvcsv1.CreateResourceRecordOptions_Type_Cname)
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
				Expect(*result.Type).To(BeEquivalentTo(dnssvcsv1.CreateResourceRecordOptions_Type_Cname))
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
				Expect(*result.Type).To(BeEquivalentTo(dnssvcsv1.CreateResourceRecordOptions_Type_Cname))
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
				createResourceRecordOptions.SetType(dnssvcsv1.CreateResourceRecordOptions_Type_Mx)
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
				Expect(*result.Type).To(BeEquivalentTo(dnssvcsv1.CreateResourceRecordOptions_Type_Mx))
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
				Expect(*result.Type).To(BeEquivalentTo(dnssvcsv1.CreateResourceRecordOptions_Type_Mx))
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
				createResourceRecordOptions.SetType(dnssvcsv1.CreateResourceRecordOptions_Type_Srv)
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
				Expect(*result.Type).To(BeEquivalentTo(dnssvcsv1.CreateResourceRecordOptions_Type_Srv))
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
				Expect(*result.Type).To(BeEquivalentTo(dnssvcsv1.CreateResourceRecordOptions_Type_Srv))
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
				createResourceRecordOptions.SetType(dnssvcsv1.CreateResourceRecordOptions_Type_Txt)
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
				Expect(*result.Type).To(BeEquivalentTo(dnssvcsv1.CreateResourceRecordOptions_Type_Txt))
				Expect(*result.Name).To(BeEquivalentTo("testtxt." + *zoneInfo.Name))

				txtRecordID := result.ID
				// Test Update Resource Record TXT
				updateResourceRecordOptions := &dnssvcsv1.UpdateResourceRecordOptions{}
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
				Expect(*result.Type).To(BeEquivalentTo(dnssvcsv1.CreateResourceRecordOptions_Type_Txt))
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
		})
	})
	Describe(`PDNSgloballoadbalancersv1`, func() {
		var zoneInfo *dnssvcsv1.Dnszone

		Context(`PDNSgloballoadbalancersv1`, func() {
			BeforeEach(func() {
				shouldSkipTest()

				// delete all dns zones
				listOptions := service.NewListDnszonesOptions(instanceID)
				listResult, listResp, listErr := service.ListDnszones(listOptions)
				Expect(listErr).To(BeNil())
				Expect(listResp).ToNot(BeNil())
				Expect(listResult).ToNot(BeNil())
				Expect(listResp.GetStatusCode()).To(BeEquivalentTo(200))
				for _, zone := range listResult.Dnszones {
					if strings.Contains(*zone.Name, "glb-example") {
						//delete all PDNS GLB load balancers
						listLoadBalancerOptions := service.NewListLoadBalancersOptions(instanceID, *zone.ID)
						listLoadBalancerResult, listLoadBalancerResp, listLoadBalancerErr := service.ListLoadBalancers(listLoadBalancerOptions)
						Expect(listLoadBalancerErr).To(BeNil())
						Expect(listLoadBalancerResp).ToNot(BeNil())
						Expect(listLoadBalancerResult).ToNot(BeNil())

						for _, record := range listLoadBalancerResult.LoadBalancers {
							deleteOpt := service.NewDeleteLoadBalancerOptions(instanceID, *zone.ID, *record.ID)
							deleteResponse, deleteErr := service.DeleteLoadBalancer(deleteOpt)
							Expect(deleteErr).To(BeNil())
							Expect(deleteResponse).ToNot(BeNil())
						}
						//delete all PDNS GLB Pools
						listPoolOptions := service.NewListPoolsOptions(instanceID)
						listPoolResult, listPoolResp, listPoolErr := service.ListPools(listPoolOptions)
						Expect(listPoolErr).To(BeNil())
						Expect(listPoolResp).ToNot(BeNil())
						Expect(listPoolResult).ToNot(BeNil())

						for _, record := range listPoolResult.Pools {
							deleteOpt := service.NewDeletePoolOptions(instanceID, *record.ID)
							deleteResponse, deleteErr := service.DeletePool(deleteOpt)
							Expect(deleteErr).To(BeNil())
							Expect(deleteResponse).ToNot(BeNil())
						}
						// delete all PDNS GLB Monitors
						listOptions := service.NewListMonitorsOptions(instanceID)
						listResult, listResp, listErr := service.ListMonitors(listOptions)
						Expect(listErr).To(BeNil())
						Expect(listResp).ToNot(BeNil())
						Expect(listResult).ToNot(BeNil())

						for _, record := range listResult.Monitors {
							deleteOpt := service.NewDeleteMonitorOptions(instanceID, *record.ID)
							deleteResponse, deleteErr := service.DeleteMonitor(deleteOpt)
							Expect(deleteErr).To(BeNil())
							Expect(deleteResponse).ToNot(BeNil())
						}
						// delete zone
						option := service.NewDeleteDnszoneOptions(instanceID, *zone.ID)
						response, err := service.DeleteDnszone(option)
						Expect(err).To(BeNil())
						Expect(response).ToNot(BeNil())
						Expect(response.GetStatusCode()).To(BeEquivalentTo(204))
					}
				}

				// Create DNS Zone
				zoneName := fmt.Sprintf("glb-example%s.com", uuid.New().String())
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
				zoneInfo = result
			})
			AfterEach(func() {
				shouldSkipTest()

				// delete all dns zones
				listOptions := service.NewListDnszonesOptions(instanceID)
				listResult, listResp, listErr := service.ListDnszones(listOptions)
				Expect(listErr).To(BeNil())
				Expect(listResp).ToNot(BeNil())
				Expect(listResult).ToNot(BeNil())
				Expect(listResp.GetStatusCode()).To(BeEquivalentTo(200))
				for _, zone := range listResult.Dnszones {
					if strings.Contains(*zone.Name, "glb-example") {
						//delete all PDNS GLB load balancers
						listLoadBalancerOptions := service.NewListLoadBalancersOptions(instanceID, *zone.ID)
						listLoadBalancerResult, listLoadBalancerResp, listLoadBalancerErr := service.ListLoadBalancers(listLoadBalancerOptions)
						Expect(listLoadBalancerErr).To(BeNil())
						Expect(listLoadBalancerResp).ToNot(BeNil())
						Expect(listLoadBalancerResult).ToNot(BeNil())

						for _, record := range listLoadBalancerResult.LoadBalancers {
							deleteOpt := service.NewDeleteLoadBalancerOptions(instanceID, *zone.ID, *record.ID)
							deleteResponse, deleteErr := service.DeleteLoadBalancer(deleteOpt)
							Expect(deleteErr).To(BeNil())
							Expect(deleteResponse).ToNot(BeNil())
						}
						// delete all PDNS GLB pools
						listPoolOptions := service.NewListPoolsOptions(instanceID)
						listPoolResult, listPoolResp, listPoolErr := service.ListPools(listPoolOptions)
						Expect(listPoolErr).To(BeNil())
						Expect(listPoolResp).ToNot(BeNil())
						Expect(listPoolResult).ToNot(BeNil())

						for _, record := range listPoolResult.Pools {
							deleteOpt := service.NewDeletePoolOptions(instanceID, *record.ID)
							deleteResponse, deleteErr := service.DeletePool(deleteOpt)
							Expect(deleteErr).To(BeNil())
							Expect(deleteResponse).ToNot(BeNil())
						}
						// delete all PDNS GLB Monitors
						listOptions := service.NewListMonitorsOptions(instanceID)
						listResult, listResp, listErr := service.ListMonitors(listOptions)
						Expect(listErr).To(BeNil())
						Expect(listResp).ToNot(BeNil())
						Expect(listResult).ToNot(BeNil())

						for _, record := range listResult.Monitors {
							deleteOpt := service.NewDeleteMonitorOptions(instanceID, *record.ID)
							deleteResponse, deleteErr := service.DeleteMonitor(deleteOpt)
							Expect(deleteErr).To(BeNil())
							Expect(deleteResponse).ToNot(BeNil())
						}
						// delete zone
						option := service.NewDeleteDnszoneOptions(instanceID, *zone.ID)
						response, err := service.DeleteDnszone(option)
						Expect(err).To(BeNil())
						Expect(response).ToNot(BeNil())
						Expect(response.GetStatusCode()).To(BeEquivalentTo(204))
					}
				}
			})
			It(`create/list PDNS load balancer monitor and pool`, func() {
				shouldSkipTest()

				//Create and List Monitor
				for i := 1; i < 4; i++ {
					CreateMonitorOptions := service.NewCreateMonitorOptions(instanceID)
					CreateMonitorOptions.SetName("testaMonitor-" + strconv.Itoa(i))
					CreateMonitorOptions.SetType(dnssvcsv1.CreateMonitorOptions_Type_Http)
					CreateMonitorOptions.SetExpectedCodes("200")
					result, response, reqErr := service.CreateMonitor(CreateMonitorOptions)
					Expect(reqErr).To(BeNil())
					Expect(response).ToNot(BeNil())
					Expect(result).ToNot(BeNil())
					Expect(response.GetStatusCode()).To(BeEquivalentTo(200))
					Expect(*result.Type).To(BeEquivalentTo(dnssvcsv1.CreateMonitorOptions_Type_Http))
				}
				listMonitorOpt := service.NewListMonitorsOptions(instanceID)
				listMonitorResult, listMonitorResponse, listMonitorErr := service.ListMonitors(listMonitorOpt)
				Expect(listMonitorErr).To(BeNil())
				Expect(listMonitorResponse).ToNot(BeNil())
				Expect(listMonitorResult).ToNot(BeNil())
				Expect(listMonitorResponse.GetStatusCode()).To(BeEquivalentTo(200))

				listPoolOpt := service.NewListPoolsOptions(instanceID)
				listPoolResult, listPoolResponse, listPoolErr := service.ListPools(listPoolOpt)
				Expect(listPoolErr).To(BeNil())
				Expect(listPoolResponse).ToNot(BeNil())
				Expect(listPoolResult).ToNot(BeNil())
				Expect(listPoolResponse.GetStatusCode()).To(BeEquivalentTo(200))
			})
			It(`crate/update/get/delete PDNS GLB monitor,pool and load balancer`, func() {
				shouldSkipTest()

				// create Load Balancer Monitor
				CreateMonitorOptions := service.NewCreateMonitorOptions(instanceID)
				CreateMonitorOptions.SetName("testa")
				CreateMonitorOptions.SetExpectedCodes("200")
				CreateMonitorOptions.SetType(dnssvcsv1.CreateMonitorOptions_Type_Http)
				CreateMonitorOptions.SetDescription("PDNS Load balancer monitor.")
				CreateMonitorOptions.SetPort(8080)
				CreateMonitorOptions.SetInterval(60)
				CreateMonitorOptions.SetRetries(2)
				CreateMonitorOptions.SetTimeout(5)
				CreateMonitorOptions.SetMethod(dnssvcsv1.CreateMonitorOptions_Method_Get)
				CreateMonitorOptions.SetPath("health")
				CreateMonitorOptions.SetAllowInsecure(false)
				CreateMonitorOptions.SetExpectedBody("alive")
				result, response, reqErr := service.CreateMonitor(CreateMonitorOptions)
				Expect(reqErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
				Expect(response.GetStatusCode()).To(BeEquivalentTo(200))
				Expect(*result.Type).To(BeEquivalentTo(dnssvcsv1.CreateMonitorOptions_Type_Http))
				Expect(*result.Name).To(BeEquivalentTo("testa"))
				Expect(*result.Description).To(BeEquivalentTo("PDNS Load balancer monitor."))
				Expect(*result.Port).To(BeEquivalentTo(8080))
				Expect(*result.Interval).To(BeEquivalentTo(60))
				Expect(*result.Retries).To(BeEquivalentTo(2))
				Expect(*result.Timeout).To(BeEquivalentTo(5))
				Expect(*result.Method).To(BeEquivalentTo(dnssvcsv1.CreateMonitorOptions_Method_Get))
				Expect(*result.Path).To(BeEquivalentTo("health"))
				Expect(*result.AllowInsecure).To(BeEquivalentTo(false))
				Expect(*result.ExpectedCodes).To(BeEquivalentTo("200"))
				Expect(*result.ExpectedBody).To(BeEquivalentTo("alive"))

				//Test GetMonitor
				getOpt := service.NewGetMonitorOptions(instanceID, *result.ID)
				getResult, getResponse, getErr := service.GetMonitor(getOpt)
				Expect(getErr).To(BeNil())
				Expect(getResponse).ToNot(BeNil())
				Expect(getResult).ToNot(BeNil())
				Expect(getResponse.GetStatusCode()).To(BeEquivalentTo(200))
				Expect(*getResult.Type).To(BeEquivalentTo(dnssvcsv1.CreateMonitorOptions_Type_Http))
				Expect(*getResult.Name).To(BeEquivalentTo("testa"))
				Expect(*getResult.Description).To(BeEquivalentTo("PDNS Load balancer monitor."))
				Expect(*getResult.Port).To(BeEquivalentTo(8080))
				Expect(*getResult.Interval).To(BeEquivalentTo(60))
				Expect(*getResult.Retries).To(BeEquivalentTo(2))
				Expect(*getResult.Timeout).To(BeEquivalentTo(5))
				Expect(*getResult.Method).To(BeEquivalentTo(dnssvcsv1.CreateMonitorOptions_Method_Get))
				Expect(*getResult.Path).To(BeEquivalentTo("health"))
				Expect(*getResult.AllowInsecure).To(BeEquivalentTo(false))
				Expect(*getResult.ExpectedCodes).To(BeEquivalentTo("200"))

				//Test UpdateMonitor
				updateOpt := service.NewUpdateMonitorOptions(instanceID, *result.ID)
				updateOpt.SetName("updatea")
				updateOpt.SetType(dnssvcsv1.UpdateMonitorOptions_Type_Https)
				updateResult, updateResponse, updateErr := service.UpdateMonitor(updateOpt)
				Expect(updateErr).To(BeNil())
				Expect(updateResponse).ToNot(BeNil())
				Expect(updateResult).ToNot(BeNil())
				Expect(updateResponse.GetStatusCode()).To(BeEquivalentTo(200))
				Expect(*updateResult.Name).To(BeEquivalentTo("updatea"))
				Expect(*updateResult.Type).To(BeEquivalentTo(dnssvcsv1.UpdateMonitorOptions_Type_Https))

				//Test CreatePool
				CreatePoolOptions := service.NewCreatePoolOptions(instanceID)
				CreatePoolOptions.SetName("testPool")
				CreatePoolOptions.SetDescription("creating pool")
				CreatePoolOptions.SetEnabled(true)
				CreatePoolOptions.SetHealthyOriginsThreshold(1)
				originInputModel := new(dnssvcsv1.OriginInput)
				originInputModel.Name = core.StringPtr("app-server-1")
				originInputModel.Description = core.StringPtr("description of the origin server")
				originInputModel.Address = core.StringPtr("10.10.10.8")
				originInputModel.Enabled = core.BoolPtr(true)
				CreatePoolOptions.Origins = []dnssvcsv1.OriginInput{*originInputModel}
				resultPool, responsePool, reqErrPool := service.CreatePool(CreatePoolOptions)
				Expect(reqErrPool).To(BeNil())
				Expect(responsePool).ToNot(BeNil())
				Expect(resultPool).ToNot(BeNil())
				Expect(responsePool.GetStatusCode()).To(BeEquivalentTo(200))
				Expect(*resultPool.Name).To(BeEquivalentTo("testPool"))
				Expect(*resultPool.Description).To(BeEquivalentTo("creating pool"))
				Expect(*resultPool.Enabled).To(BeEquivalentTo(true))
				Expect(*resultPool.HealthyOriginsThreshold).To(BeEquivalentTo(1))

				//Test Get Pool
				getPoolOpt := service.NewGetPoolOptions(instanceID, *resultPool.ID)
				getPoolResult, gePooltResponse, getPoolErr := service.GetPool(getPoolOpt)
				Expect(getPoolErr).To(BeNil())
				Expect(gePooltResponse).ToNot(BeNil())
				Expect(getPoolResult).ToNot(BeNil())
				Expect(gePooltResponse.GetStatusCode()).To(BeEquivalentTo(200))
				Expect(*getPoolResult.Name).To(BeEquivalentTo("testPool"))
				Expect(*getPoolResult.Description).To(BeEquivalentTo("creating pool"))
				Expect(*getPoolResult.Enabled).To(BeEquivalentTo(true))
				Expect(*getPoolResult.HealthyOriginsThreshold).To(BeEquivalentTo(1))

				//Test Update Pool
				updatePoolOpt := service.NewUpdatePoolOptions(instanceID, *resultPool.ID)
				updatePoolOpt.SetName("updatedtestpool")
				updatePoolOpt.SetDescription("updating testPool")
				updatePoolResult, updatePoolResponse, updatePoolErr := service.UpdatePool(updatePoolOpt)
				Expect(updatePoolErr).To(BeNil())
				Expect(updatePoolResponse).ToNot(BeNil())
				Expect(updatePoolResult).ToNot(BeNil())
				Expect(updatePoolResponse.GetStatusCode()).To(BeEquivalentTo(200))
				Expect(*updatePoolResult.Name).To(BeEquivalentTo("updatedtestpool"))
				Expect(*updatePoolResult.Description).To(BeEquivalentTo("updating testPool"))

				//Test Create Load Balancer
				CreateLoadBalancerOptions := service.NewCreateLoadBalancerOptions(instanceID, *zoneInfo.ID)
				CreateLoadBalancerOptions.SetName("testloadbalancer")
				CreateLoadBalancerOptions.SetDescription("PDNS Load balancer")
				CreateLoadBalancerOptions.SetEnabled(true)
				CreateLoadBalancerOptions.SetTTL(120)
				CreateLoadBalancerOptions.SetFallbackPool(*resultPool.ID)
				CreateLoadBalancerOptions.SetDefaultPools([]string{*resultPool.ID})
				resultLoadbalancer, responseLoadbalancer, reqErrLoadbalancer := service.CreateLoadBalancer(CreateLoadBalancerOptions)
				Expect(reqErrLoadbalancer).To(BeNil())
				Expect(responseLoadbalancer).ToNot(BeNil())
				Expect(resultLoadbalancer).ToNot(BeNil())
				Expect(responseLoadbalancer.GetStatusCode()).To(BeEquivalentTo(200))
				Expect(*resultLoadbalancer.Name).To(ContainSubstring("testloadbalancer"))
				Expect(*resultLoadbalancer.Description).To(BeEquivalentTo("PDNS Load balancer"))
				Expect(*resultLoadbalancer.Enabled).To(BeEquivalentTo(true))
				Expect(*resultLoadbalancer.TTL).To(BeEquivalentTo(120))
				Expect(*resultLoadbalancer.FallbackPool).To(BeEquivalentTo(*resultPool.ID))

				//Test Get Load Balancer
				getLoadBalancerOpt := service.NewGetLoadBalancerOptions(instanceID, *zoneInfo.ID, *resultLoadbalancer.ID)
				getLoadBalancerResult, getLoadBalancerResponse, getLoadBalancerErr := service.GetLoadBalancer(getLoadBalancerOpt)
				Expect(getLoadBalancerErr).To(BeNil())
				Expect(getLoadBalancerResponse).ToNot(BeNil())
				Expect(getLoadBalancerResult).ToNot(BeNil())
				Expect(getLoadBalancerResponse.GetStatusCode()).To(BeEquivalentTo(200))
				Expect(*getLoadBalancerResult.Name).To(ContainSubstring("testloadbalancer"))
				Expect(*getLoadBalancerResult.Description).To(BeEquivalentTo("PDNS Load balancer"))
				Expect(*getLoadBalancerResult.Enabled).To(BeEquivalentTo(true))
				Expect(*getLoadBalancerResult.TTL).To(BeEquivalentTo(120))
				Expect(*getLoadBalancerResult.FallbackPool).To(BeEquivalentTo(*resultPool.ID))

				//Test Update Load Balancer
				updateLoadBalancerOpt := service.NewUpdateLoadBalancerOptions(instanceID, *zoneInfo.ID, *resultLoadbalancer.ID)
				updateLoadBalancerOpt.SetName("updateLoadBalancer")
				updateLoadBalancerOpt.SetDescription("updating Load Balancer")
				updateLoadBalancerResult, updateLoadBalancerResponse, updateLoadBalancerErr := service.UpdateLoadBalancer(updateLoadBalancerOpt)
				Expect(updateLoadBalancerErr).To(BeNil())
				Expect(updateLoadBalancerResponse).ToNot(BeNil())
				Expect(updateLoadBalancerResult).ToNot(BeNil())
				Expect(updateLoadBalancerResponse.GetStatusCode()).To(BeEquivalentTo(200))
				Expect(*updateLoadBalancerResult.Name).To(ContainSubstring("updateLoadBalancer"))
				Expect(*updateLoadBalancerResult.Description).To(BeEquivalentTo("updating Load Balancer"))
				Expect(*updateLoadBalancerResult.Enabled).To(BeEquivalentTo(true))

				//Test List Load Balancer
				listLoadBalancerOpt := service.NewListLoadBalancersOptions(instanceID, *zoneInfo.ID)
				listLoadBalancerResult, listLoadBalancerResponse, listLoadBalancerErr := service.ListLoadBalancers(listLoadBalancerOpt)
				Expect(listLoadBalancerErr).To(BeNil())
				Expect(listLoadBalancerResponse).ToNot(BeNil())
				Expect(listLoadBalancerResult).ToNot(BeNil())
				Expect(listLoadBalancerResponse.GetStatusCode()).To(BeEquivalentTo(200))

				//Test DeleteLoadBalancer
				deleteLoadBalancerOpt := service.NewDeleteLoadBalancerOptions(instanceID, *zoneInfo.ID, *resultLoadbalancer.ID)
				deleteLoadBalancerResponse, deleteLoadBalancerErr := service.DeleteLoadBalancer(deleteLoadBalancerOpt)
				Expect(deleteLoadBalancerErr).To(BeNil())
				Expect(deleteLoadBalancerResponse).ToNot(BeNil())
				Expect(deleteLoadBalancerResponse.GetStatusCode()).To(BeEquivalentTo(204))

				//Test DeletePool
				deletePoolOpt := service.NewDeletePoolOptions(instanceID, *resultPool.ID)
				deletePoolResponse, deletePoolErr := service.DeletePool(deletePoolOpt)
				Expect(deletePoolErr).To(BeNil())
				Expect(deletePoolResponse).ToNot(BeNil())
				Expect(deletePoolResponse.GetStatusCode()).To(BeEquivalentTo(204))

				//Test DeleteMonitor
				deleteOpt := service.NewDeleteMonitorOptions(instanceID, *result.ID)
				deleteResponse, deleteErr := service.DeleteMonitor(deleteOpt)
				Expect(deleteErr).To(BeNil())
				Expect(deleteResponse).ToNot(BeNil())
				Expect(deleteResponse.GetStatusCode()).To(BeEquivalentTo(204))
			})

		})
	})
	Describe(`permittednetworksfordnszonesv1`, func() {
		Context(`permittednetworksfordnszonesv1`, func() {
			var zoneInfo *dnssvcsv1.Dnszone
			BeforeEach(func() {
				shouldSkipTest()

				// delete all dns zones
				listOptions := service.NewListDnszonesOptions(instanceID)
				listResult, listResp, listErr := service.ListDnszones(listOptions)
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

						option := service.NewDeleteDnszoneOptions(instanceID, *zone.ID)
						response, err := service.DeleteDnszone(option)
						Expect(err).To(BeNil())
						Expect(response).ToNot(BeNil())
						Expect(response.GetStatusCode()).To(BeEquivalentTo(204))
					}
				}

				// Create DNS Zone
				zoneName := fmt.Sprintf("network-example-%s.com", uuid.New().String())
				createDnszoneOptions := service.NewCreateDnszoneOptions(instanceID)
				createDnszoneOptions.SetName(zoneName)
				createDnszoneOptions.SetDescription("testString")
				createDnszoneOptions.SetLabel("testString")
				result, response, reqErr := service.CreateDnszone(createDnszoneOptions)
				Expect(reqErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
				Expect(response.GetStatusCode()).To(BeEquivalentTo(200))
				zoneInfo = result
			})
			AfterEach(func() {
				shouldSkipTest()
				// delete all dns zones
				listOptions := service.NewListDnszonesOptions(instanceID)
				listResult, listResp, listErr := service.ListDnszones(listOptions)
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

						option := service.NewDeleteDnszoneOptions(instanceID, *zone.ID)
						response, err := service.DeleteDnszone(option)
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
				permittedNwVPCOption, err := service.NewPermittedNetworkVpc(vpcCrn)
				Expect(err).To(BeNil())
				createPermittedNetworkOptions.SetPermittedNetwork(permittedNwVPCOption)
				createPermittedNetworkOptions.SetType(dnssvcsv1.CreatePermittedNetworkOptions_Type_Vpc)
				createPermittedNetworkOptions.SetHeaders(header)
				result, response, reqErr := service.CreatePermittedNetwork(createPermittedNetworkOptions)
				Expect(reqErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
				Expect(response.GetStatusCode()).To(BeEquivalentTo(200))
				Expect(*result.State).To(BeEquivalentTo(dnssvcsv1.PermittedNetwork_State_Active))

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
				Expect(*result.State).To(BeEquivalentTo(dnssvcsv1.PermittedNetwork_State_Active))

				// Test Get Permitted Network Fail
				fgetPermittedNetworkOptions := service.NewGetPermittedNetworkOptions(instanceID, *zoneInfo.ID, "invalid_id")
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
				Expect(*result.State).To(BeEquivalentTo(dnssvcsv1.PermittedNetwork_State_RemovalInProgress))

				// Test Rmove Permitted Network Fail
				fdeletePermittedNetworkOptions := service.NewDeletePermittedNetworkOptions(instanceID, *zoneInfo.ID, "invalid_id")
				_, _, reqErr = service.DeletePermittedNetwork(fdeletePermittedNetworkOptions)
				Expect(reqErr).ToNot(BeNil())
			})
		})
	})
})
