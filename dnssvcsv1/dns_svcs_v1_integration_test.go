/**
 * (C) Copyright IBM Corp. 2024.
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

// package go_build_test

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	dnssvcsv1 "github.com/IBM/networking-go-sdk/dnssvcsv1"
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
	defer GinkgoRecover()
	Skip("Skipping test cases..")
	if _, err := os.Stat(configFile); err != nil {
		configLoaded = false
	}

	err := godotenv.Load(configFile)
	if err != nil {
		configLoaded = false
	}

	// first API key
	authenticator, err := core.GetAuthenticatorFromEnvironment("dns_svcs")
	if err != nil {
		panic(err)
	}
	dnsServicesURL := os.Getenv("DNS_SVCS_URL")
	options := &dnssvcsv1.DnsSvcsV1Options{
		ServiceName:   "dns_svcs",
		URL:           dnsServicesURL,
		Authenticator: authenticator,
	}
	service, serviceErr := dnssvcsv1.NewDnsSvcsV1UsingExternalConfig(options)
	if serviceErr != nil {
		panic(serviceErr)
	}
	ownerAPIKey := os.Getenv("DNS_SVCS_OWNER_APIKEY")
	if ownerAPIKey == "" {
		panic("Cross account owner API key is not set")
	}

	setEnvErr := os.Setenv("DNS_SVCS_APIKEY", ownerAPIKey)
	if setEnvErr != nil {
		panic(setEnvErr)
	}

	// second API key
	authenticatorOwnerDnsInstanceAccount, err := core.GetAuthenticatorFromEnvironment("dns_svcs")
	if err != nil {
		panic(err)
	}
	optionsOwnerDnsInstanceAccount := &dnssvcsv1.DnsSvcsV1Options{
		ServiceName:   "dns_svcs",
		URL:           dnsServicesURL,
		Authenticator: authenticatorOwnerDnsInstanceAccount,
	}
	serviceOwnerDnsInstanceAccount, serviceErr := dnssvcsv1.NewDnsSvcsV1UsingExternalConfig(optionsOwnerDnsInstanceAccount)
	if serviceErr != nil {
		panic(err)
	}

	instanceID := os.Getenv("DNS_SVCS_INSTANCE_ID")
	ownerInstanceID := os.Getenv("DNS_SVCS_OWNER_INSTANCE_ID")
	ownerZoneID := os.Getenv("DNS_SVCS_OWNER_ZONE_ID")
	vpcCrnLzPermittedNetwork := os.Getenv("DNS_SVCS_VPC_CRN_LZ_PERMITTED_NETWORK")
	vpcCrn := os.Getenv("DNS_SVCS_VPC_CRN")
	subnetCrn := os.Getenv("DNS_SVCS_SUBNET_CRN")
	customCrn := os.Getenv("DNS_SVCS_CUSTOMER_LOCATION_SUBNET_CRN")

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
				createDnszoneOptions := service.NewCreateDnszoneOptions(instanceID, zoneName)
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

						// First delete PTR record to avoid any record linking issue
						listOptions2 := service.NewListResourceRecordsOptions(instanceID, *zone.ID)
						listOptions2.SetType("PTR")
						listResult2, _, listErr2 := service.ListResourceRecords(listOptions2)
						Expect(listErr2).To(BeNil())
						for _, ptrRecord := range listResult2.ResourceRecords {
							deleteOpt := service.NewDeleteResourceRecordOptions(instanceID, *zone.ID, *ptrRecord.ID)
							deleteResponse, deleteErr := service.DeleteResourceRecord(deleteOpt)
							Expect(deleteErr).To(BeNil())
							Expect(deleteResponse).ToNot(BeNil())
						}
						// Delete remaining records
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
				createDnszoneOptions := service.NewCreateDnszoneOptions(instanceID, zoneName)
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

						// First delete PTR record to avoid any record linking issue
						listOptions2 := service.NewListResourceRecordsOptions(instanceID, *zone.ID)
						listOptions2.SetType("PTR")
						listResult2, _, listErr2 := service.ListResourceRecords(listOptions2)
						Expect(listErr2).To(BeNil())

						for _, ptrRecord := range listResult2.ResourceRecords {
							deleteOpt := service.NewDeleteResourceRecordOptions(instanceID, *zone.ID, *ptrRecord.ID)
							deleteResponse, deleteErr := service.DeleteResourceRecord(deleteOpt)
							Expect(deleteErr).To(BeNil())
							Expect(deleteResponse).ToNot(BeNil())
						}
						// Delete remaining records now
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
				createResourceRecordOptions := service.NewCreateResourceRecordOptions(instanceID, *zoneInfo.ID, dnssvcsv1.CreateResourceRecordOptions_Type_A)
				createResourceRecordOptions.SetName("testa")
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

				updaterdataARecord, err := service.NewResourceRecordUpdateInputRdataRdataARecord("1.1.1.2")
				Expect(err).To(BeNil())
				updateOpt := service.NewUpdateResourceRecordOptions(instanceID, *zoneInfo.ID, *result.ID, "updatea", updaterdataARecord)
				updateOpt.SetName("updatea")
				updateOpt.SetTTL(300)
				updateOpt.SetXCorrelationID("abc123")
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
				createResourceRecordOptions := service.NewCreateResourceRecordOptions(instanceID, *zoneInfo.ID, dnssvcsv1.CreateResourceRecordOptions_Type_A)
				createResourceRecordOptions.SetName("testa")
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

				createResourcePtrRecordOptions := service.NewCreateResourceRecordOptions(instanceID, *zoneInfo.ID, dnssvcsv1.CreateResourceRecordOptions_Type_Ptr)
				createResourcePtrRecordOptions.SetName("1.1.1.1")
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
				createResourceRecordOptions := service.NewCreateResourceRecordOptions(instanceID, *zoneInfo.ID, dnssvcsv1.CreateResourceRecordOptions_Type_Aaaa)
				createResourceRecordOptions.SetName("testaaaa")
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
				updaterdataAaaaRecord, err := service.NewResourceRecordUpdateInputRdataRdataAaaaRecord("2001::8889")
				Expect(err).To(BeNil())
				updateResourceRecordOptions := service.NewUpdateResourceRecordOptions(instanceID, *zoneInfo.ID, aaaaRecordID, "updateaaaa", updaterdataAaaaRecord)
				updateResourceRecordOptions.SetTTL(300)
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
				createResourceRecordOptions := service.NewCreateResourceRecordOptions(instanceID, *zoneInfo.ID, dnssvcsv1.CreateResourceRecordOptions_Type_Cname)
				createResourceRecordOptions.SetName("testcname")
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
				updaterdataCnameRecord, err := service.NewResourceRecordUpdateInputRdataRdataCnameRecord("updatecnamedata.com")
				Expect(err).To(BeNil())
				updateResourceRecordOptions := service.NewUpdateResourceRecordOptions(instanceID, *zoneInfo.ID, *cnameRecordID, "updatecname", updaterdataCnameRecord)
				updateResourceRecordOptions.SetTTL(300)

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
				createResourceRecordOptions := service.NewCreateResourceRecordOptions(instanceID, *zoneInfo.ID, dnssvcsv1.CreateResourceRecordOptions_Type_Mx)
				createResourceRecordOptions.SetName("testmx")
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
				updaterdataMxRecord, err := service.NewResourceRecordUpdateInputRdataRdataMxRecord("mail1.testmx.com", 2)
				Expect(err).To(BeNil())
				updateResourceRecordOptions := service.NewUpdateResourceRecordOptions(instanceID, *zoneInfo.ID, *mxRecordID, "testupdatemx", updaterdataMxRecord)
				updateResourceRecordOptions.SetTTL(300)
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
				createResourceRecordOptions := service.NewCreateResourceRecordOptions(instanceID, *zoneInfo.ID, dnssvcsv1.CreateResourceRecordOptions_Type_Srv)
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
				updaterdataSrvRecord, err := service.NewResourceRecordUpdateInputRdataRdataSrvRecord(2, 2, "updatesiphost.com", 2)
				Expect(err).To(BeNil())
				updateResourceRecordOptions := service.NewUpdateResourceRecordOptions(instanceID, *zoneInfo.ID, *srvRecordID, "updatesrv", updaterdataSrvRecord)
				updateResourceRecordOptions.SetTTL(300)
				updateResourceRecordOptions.SetService("_sip")
				updateResourceRecordOptions.SetProtocol("udp")
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
				createResourceRecordOptions := service.NewCreateResourceRecordOptions(instanceID, *zoneInfo.ID, dnssvcsv1.CreateResourceRecordOptions_Type_Txt)
				createResourceRecordOptions.SetName("testtxt")
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
			It(`create import records`, func() {
				shouldSkipTest()

				// Import DNS records
				importResourceRecordsOptions := service.NewImportResourceRecordsOptions(instanceID, *zoneInfo.ID)
				zoneName := fmt.Sprintf("test-example%s.com", uuid.New().String())
				f := strings.NewReader(zoneName + ` 1 IN AAAA 2001::888`)
				importResourceRecordsOptions.SetFile(io.NopCloser(f))
				importResourceRecordsOptions.SetXCorrelationID("abc123")
				importResourceRecordsOptions.SetFileContentType("application/json")
				result, response, reqErr := service.ImportResourceRecords(importResourceRecordsOptions)
				Expect(reqErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
				Expect(response.GetStatusCode()).To(BeEquivalentTo(200))
				Expect(*result.TotalRecordsParsed).To(BeEquivalentTo(int64(1)))
			})
			It(`get Export records`, func() {
				shouldSkipTest()

				//create a resource record
				createResourceRecordOptions := service.NewCreateResourceRecordOptions(instanceID, *zoneInfo.ID, dnssvcsv1.CreateResourceRecordOptions_Type_A)
				createResourceRecordOptions.SetName("teststring")
				createResourceRecordOptions.SetTTL(120)
				rdataARecord, err := service.NewResourceRecordInputRdataRdataARecord("1.1.1.1")
				Expect(err).To(BeNil())
				createResourceRecordOptions.SetRdata(rdataARecord)
				createResourceRecordOptions.SetXCorrelationID("abc123")
				rresult, rresponse, rreqErr := service.CreateResourceRecord(createResourceRecordOptions)
				Expect(rreqErr).To(BeNil())
				Expect(rresponse).ToNot(BeNil())
				Expect(rresult).ToNot(BeNil())
				Expect(rresponse.GetStatusCode()).To(BeEquivalentTo(200))
				Expect(*rresult.Type).To(BeEquivalentTo(dnssvcsv1.CreateResourceRecordOptions_Type_A))

				// Export DNS Records.
				exportResourceRecordsOptions := service.NewExportResourceRecordsOptions(instanceID, *zoneInfo.ID)
				exportResourceRecordsOptions.SetXCorrelationID("abc123")
				result, response, reqErr := service.ExportResourceRecords(exportResourceRecordsOptions)
				Expect(reqErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
				Expect(response.GetStatusCode()).To(BeEquivalentTo(200))
			})
		})
		Context(`resourcerecordsv1 - list`, func() {
			var zoneInfo *dnssvcsv1.Dnszone
			BeforeEach(func() {
				shouldSkipTest()
				listOptions := service.NewListDnszonesOptions(instanceID)
				listResult, listResp, listErr := service.ListDnszones(listOptions)
				Expect(listErr).To(BeNil())
				Expect(listResp).ToNot(BeNil())
				Expect(listResult).ToNot(BeNil())
				Expect(listResp.GetStatusCode()).To(BeEquivalentTo(200))

				if len(listResult.Dnszones) != 0 {
					zone := listResult.Dnszones[0]
					// First delete PTR record to avoid any record linking issue
					listOptions2 := service.NewListResourceRecordsOptions(instanceID, *zoneInfo.ID)
					listOptions2.SetType("PTR")
					listResult2, _, listErr2 := service.ListResourceRecords(listOptions2)
					Expect(listErr2).To(BeNil())

					for _, ptrRecord := range listResult2.ResourceRecords {
						deleteOpt := service.NewDeleteResourceRecordOptions(instanceID, *zone.ID, *ptrRecord.ID)
						deleteResponse, deleteErr := service.DeleteResourceRecord(deleteOpt)
						Expect(deleteErr).To(BeNil())
						Expect(deleteResponse).ToNot(BeNil())
					}

					// Now iterate again and delete remaining records
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

				// Create DNS Zone
				zoneName := fmt.Sprintf("test-example%s.com", uuid.New().String())
				createDnszoneOptions := service.NewCreateDnszoneOptions(instanceID, zoneName)
				createDnszoneOptions.SetDescription("testString")
				createDnszoneOptions.SetLabel("testString")
				createDnszoneOptions.SetXCorrelationID("abc123")
				result, response, reqErr := service.CreateDnszone(createDnszoneOptions)
				Expect(reqErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
				Expect(response.GetStatusCode()).To(BeEquivalentTo(200))
				zoneInfo = result

				// Create records for list operation validation by name, type and name+type.
				createResourceRecordOptions1 := service.NewCreateResourceRecordOptions(instanceID, *zoneInfo.ID, dnssvcsv1.CreateResourceRecordOptions_Type_A)
				createResourceRecordOptions1.SetName("atest")
				createResourceRecordOptions1.SetTTL(120)
				rdataARecord, err1 := service.NewResourceRecordInputRdataRdataARecord("1.1.1.1")
				Expect(err1).To(BeNil())
				createResourceRecordOptions1.SetRdata(rdataARecord)
				createResourceRecordOptions1.SetXCorrelationID("abc123")
				result1, response1, reqErr1 := service.CreateResourceRecord(createResourceRecordOptions1)

				Expect(reqErr1).To(BeNil())
				Expect(response1).ToNot(BeNil())
				Expect(result1).ToNot(BeNil())
				Expect(response1.GetStatusCode()).To(BeEquivalentTo(200))
				Expect(*result1.Type).To(BeEquivalentTo(dnssvcsv1.CreateResourceRecordOptions_Type_A))

				// create resource record PTR
				createResourceRecordOptions2 := service.NewCreateResourceRecordOptions(instanceID, *zoneInfo.ID, dnssvcsv1.CreateResourceRecordOptions_Type_A)
				createResourceRecordOptions2.SetName("atest")
				createResourceRecordOptions2.SetTTL(120)
				rdataARecord2, err2 := service.NewResourceRecordInputRdataRdataARecord("1.1.1.2")
				Expect(err2).To(BeNil())
				createResourceRecordOptions2.SetRdata(rdataARecord2)
				createResourceRecordOptions2.SetXCorrelationID("abc123")
				result2, response2, reqErr2 := service.CreateResourceRecord(createResourceRecordOptions2)
				Expect(reqErr2).To(BeNil())
				Expect(response2).ToNot(BeNil())
				Expect(result2).ToNot(BeNil())
				Expect(response2.GetStatusCode()).To(BeEquivalentTo(200))
				Expect(*result2.Type).To(BeEquivalentTo(dnssvcsv1.CreateResourceRecordOptions_Type_A))

				createResourcePtrRecordOptions2 := service.NewCreateResourceRecordOptions(instanceID, *zoneInfo.ID, dnssvcsv1.CreateResourceRecordOptions_Type_Ptr)
				createResourcePtrRecordOptions2.SetName("1.1.1.2")
				createResourcePtrRecordOptions2.SetTTL(120)
				rdataPtrRecord, err3 := service.NewResourceRecordInputRdataRdataPtrRecord("atest." + *zoneInfo.Name)
				Expect(err3).To(BeNil())
				createResourcePtrRecordOptions2.SetRdata(rdataPtrRecord)
				createResourcePtrRecordOptions2.SetXCorrelationID("abc123")
				ptrResult, ptrResponse, ptrErr := service.CreateResourceRecord(createResourcePtrRecordOptions2)
				Expect(ptrErr).To(BeNil())
				Expect(ptrResponse).ToNot(BeNil())
				Expect(ptrResult).ToNot(BeNil())
				Expect(ptrResponse.GetStatusCode()).To(BeEquivalentTo(200))
				Expect(*ptrResult.Type).To(BeEquivalentTo(dnssvcsv1.CreateResourceRecordOptions_Type_Ptr))

				// Test Create Resource Record AAAA
				header := map[string]string{
					"test": "teststring",
				}
				createResourceRecordOptions3 := service.NewCreateResourceRecordOptions(instanceID, *zoneInfo.ID, dnssvcsv1.CreateResourceRecordOptions_Type_Aaaa)
				createResourceRecordOptions3.SetName("testaaaa")
				createResourceRecordOptions3.SetTTL(120)
				rdataAaaaRecord, err4 := service.NewResourceRecordInputRdataRdataAaaaRecord("2001::8888")
				Expect(err4).To(BeNil())
				createResourceRecordOptions3.SetRdata(rdataAaaaRecord)
				createResourceRecordOptions3.SetXCorrelationID("abc123")
				createResourceRecordOptions3.SetHeaders(header)
				result3, response3, reqErr3 := service.CreateResourceRecord(createResourceRecordOptions3)
				Expect(reqErr3).To(BeNil())
				Expect(response3).ToNot(BeNil())
				Expect(result3).ToNot(BeNil())
				Expect(response3.GetStatusCode()).To(BeEquivalentTo(200))
				Expect(*result3.Type).To(BeEquivalentTo(dnssvcsv1.CreateResourceRecordOptions_Type_Aaaa))
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
				if len(listResult.Dnszones) != 0 {
					zone := listResult.Dnszones[0]
					// First delete PTR record to avoid any record linking issue
					listOptions2 := service.NewListResourceRecordsOptions(instanceID, *zoneInfo.ID)
					listOptions2.SetType("PTR")
					listResult2, _, listErr2 := service.ListResourceRecords(listOptions2)
					Expect(listErr2).To(BeNil())

					for _, ptrRecord := range listResult2.ResourceRecords {
						deleteOpt := service.NewDeleteResourceRecordOptions(instanceID, *zone.ID, *ptrRecord.ID)
						deleteResponse, deleteErr := service.DeleteResourceRecord(deleteOpt)
						Expect(deleteErr).To(BeNil())
						Expect(deleteResponse).ToNot(BeNil())
					}
					// Delete remaining records now
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
			})
			Context(`resourcerecordsv1 - list`, func() {
				It(`resourcerecordsv1 - list record by name`, func() {
					// Query 1 by name
					aRecordName := fmt.Sprintf("atest.%s", *zoneInfo.Name)
					listOptions1 := service.NewListResourceRecordsOptions(instanceID, *zoneInfo.ID)
					listOptions1.SetName(aRecordName)
					listResult1, _, _ := service.ListResourceRecords(listOptions1)
					Expect(*listResult1.ResourceRecords[0].Name).To(BeEquivalentTo(aRecordName))
					Expect(len(listResult1.ResourceRecords)).To(BeEquivalentTo(2))

					// Query 2 by name
					ptrRecordName := "1.1.1.2"
					listOptions2 := service.NewListResourceRecordsOptions(instanceID, *zoneInfo.ID)
					listOptions2.SetName(ptrRecordName)
					listResult3, _, _ := service.ListResourceRecords(listOptions2)
					Expect(*listResult3.ResourceRecords[0].Name).To(BeEquivalentTo(ptrRecordName))
					Expect(len(listResult3.ResourceRecords)).To(BeEquivalentTo(1))

					// Query 3 by name
					aaaRecordName := fmt.Sprintf("testaaaa.%s", *zoneInfo.Name)
					listOptions3 := service.NewListResourceRecordsOptions(instanceID, *zoneInfo.ID)
					listOptions3.SetName(aaaRecordName)
					listResult5, _, _ := service.ListResourceRecords(listOptions3)
					Expect(*listResult5.ResourceRecords[0].Name).To(BeEquivalentTo(aaaRecordName))
					Expect(len(listResult5.ResourceRecords)).To(BeEquivalentTo(1))
				})

				It(`resourcerecordsv1 - list record by type`, func() {
					// Query 1 by type
					aRecordName := fmt.Sprintf("atest.%s", *zoneInfo.Name)
					listOptions1 := service.NewListResourceRecordsOptions(instanceID, *zoneInfo.ID)
					listOptions1.SetType("A")
					listResult2, _, _ := service.ListResourceRecords(listOptions1)
					Expect(*listResult2.ResourceRecords[0].Name).To(BeEquivalentTo(aRecordName))
					Expect(len(listResult2.ResourceRecords)).To(BeEquivalentTo(2))

					// Query 2 by type
					ptrRecordName := "1.1.1.2"
					listOptions2 := service.NewListResourceRecordsOptions(instanceID, *zoneInfo.ID)
					listOptions2.SetType("PTR")
					listResult4, _, _ := service.ListResourceRecords(listOptions2)
					Expect(*listResult4.ResourceRecords[0].Name).To(BeEquivalentTo(ptrRecordName))
					Expect(len(listResult4.ResourceRecords)).To(BeEquivalentTo(1))

					// Query 3 by type
					aaaRecordName := fmt.Sprintf("testaaaa.%s", *zoneInfo.Name)
					listOptions3 := service.NewListResourceRecordsOptions(instanceID, *zoneInfo.ID)
					listOptions3.SetType("AAAA")
					listResult6, _, _ := service.ListResourceRecords(listOptions3)
					Expect(*listResult6.ResourceRecords[0].Name).To(BeEquivalentTo(aaaRecordName))
					Expect(len(listResult6.ResourceRecords)).To(BeEquivalentTo(1))

				})

				It(`resourcerecordsv1 - list record by name and type`, func() {
					// Query 1 by type and name
					aRecordName := fmt.Sprintf("atest.%s", *zoneInfo.Name)
					listOptions1 := service.NewListResourceRecordsOptions(instanceID, *zoneInfo.ID)
					listOptions1.SetName(aRecordName)
					listOptions1.SetType("A")
					listResult2, _, _ := service.ListResourceRecords(listOptions1)
					Expect(*listResult2.ResourceRecords[0].Name).To(BeEquivalentTo(aRecordName))
					Expect(len(listResult2.ResourceRecords)).To(BeEquivalentTo(2))

					// Query 1 by type and name
					ptrRecordName := "1.1.1.2"
					listOptions2 := service.NewListResourceRecordsOptions(instanceID, *zoneInfo.ID)
					listOptions2.SetName(ptrRecordName)
					listOptions2.SetType("PTR")
					listResult4, _, _ := service.ListResourceRecords(listOptions2)
					Expect(*listResult4.ResourceRecords[0].Name).To(BeEquivalentTo(ptrRecordName))
					Expect(len(listResult4.ResourceRecords)).To(BeEquivalentTo(1))

					// Query 1 by type and name
					aaaRecordName := fmt.Sprintf("testaaaa.%s", *zoneInfo.Name)
					listOptions3 := service.NewListResourceRecordsOptions(instanceID, *zoneInfo.ID)
					listOptions3.SetName(aaaRecordName)
					listOptions3.SetType("AAAA")
					listResult6, _, _ := service.ListResourceRecords(listOptions3)
					Expect(*listResult6.ResourceRecords[0].Name).To(BeEquivalentTo(aaaRecordName))
					Expect(len(listResult6.ResourceRecords)).To(BeEquivalentTo(1))

				})
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
							deleteOpt.SetXCorrelationID("DeleteMonitor-acctest-" + string(*record.ID))
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
				createDnszoneOptions := service.NewCreateDnszoneOptions(instanceID, zoneName)
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
					monitorName := "testaMonitor-" + strconv.Itoa(i)
					createMonitorOptions := service.NewCreateMonitorOptions(instanceID, monitorName, dnssvcsv1.CreateMonitorOptions_Type_Http)
					createMonitorOptions.SetExpectedCodes("200")
					result, response, reqErr := service.CreateMonitor(createMonitorOptions)
					Expect(reqErr).To(BeNil())
					Expect(response).ToNot(BeNil())
					Expect(result).ToNot(BeNil())
					Expect(response.GetStatusCode()).To(BeEquivalentTo(200))
					Expect(*result.Type).To(BeEquivalentTo(dnssvcsv1.CreateMonitorOptions_Type_Http))

					name := fmt.Sprintf("testpool-%d", i)
					createPoolOptionsOrigins := []dnssvcsv1.OriginInput{}
					createPoolOptions := service.NewCreatePoolOptions(instanceID, name, createPoolOptionsOrigins)
					createPoolOptions.SetDescription("creating pool")
					createPoolOptions.SetEnabled(true)
					createPoolOptions.SetHealthyOriginsThreshold(1)
					createPoolOptions.SetMonitor(*result.ID)
					createPoolOptions.SetHealthcheckRegion("us-south")
					createPoolOptions.SetHealthcheckSubnets([]string{subnetCrn})
					originInputModel := new(dnssvcsv1.OriginInput)
					originInputModel.Name = core.StringPtr("app-server-1")
					originInputModel.Description = core.StringPtr("description of the origin server")
					originInputModel.Address = core.StringPtr("10.10.10.8")
					originInputModel.Enabled = core.BoolPtr(true)
					createPoolOptions.Origins = []dnssvcsv1.OriginInput{*originInputModel}
					resultPool, responsePool, reqErrPool := service.CreatePool(createPoolOptions)
					Expect(reqErrPool).To(BeNil())
					Expect(responsePool).ToNot(BeNil())
					Expect(resultPool).ToNot(BeNil())
					Expect(responsePool.GetStatusCode()).To(BeEquivalentTo(200))
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
			It(`create/update/get/delete PDNS GLB monitor,pool and load balancer`, func() {
				shouldSkipTest()

				// create Load Balancer Monitor
				createMonitorOptions := service.NewCreateMonitorOptions(instanceID, "testa", dnssvcsv1.CreateMonitorOptions_Type_Http)
				createMonitorOptions.SetExpectedCodes("200")
				createMonitorOptions.SetDescription("PDNS Load balancer monitor.")
				createMonitorOptions.SetPort(8080)
				createMonitorOptions.SetInterval(60)
				createMonitorOptions.SetRetries(2)
				createMonitorOptions.SetTimeout(5)
				createMonitorOptions.SetMethod(dnssvcsv1.CreateMonitorOptions_Method_Get)
				createMonitorOptions.SetPath("health")
				createMonitorOptions.SetAllowInsecure(false)
				createMonitorOptions.SetExpectedBody("alive")
				result, response, reqErr := service.CreateMonitor(createMonitorOptions)
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
				createPoolOptionsOrigins := []dnssvcsv1.OriginInput{}
				createPoolOptions := service.NewCreatePoolOptions(instanceID, "testPool", createPoolOptionsOrigins)
				createPoolOptions.SetDescription("creating pool")
				createPoolOptions.SetEnabled(true)
				createPoolOptions.SetHealthyOriginsThreshold(1)
				createPoolOptions.SetMonitor(*updateResult.ID)
				createPoolOptions.SetHealthcheckRegion("us-south")
				createPoolOptions.SetHealthcheckSubnets([]string{subnetCrn})
				originInputModel := new(dnssvcsv1.OriginInput)
				originInputModel.Name = core.StringPtr("app-server-1")
				originInputModel.Description = core.StringPtr("description of the origin server")
				originInputModel.Address = core.StringPtr("10.10.10.8")
				originInputModel.Enabled = core.BoolPtr(true)
				createPoolOptions.Origins = []dnssvcsv1.OriginInput{*originInputModel}
				resultPool, responsePool, reqErrPool := service.CreatePool(createPoolOptions)
				Expect(reqErrPool).To(BeNil())
				Expect(responsePool).ToNot(BeNil())
				Expect(resultPool).ToNot(BeNil())
				Expect(responsePool.GetStatusCode()).To(BeEquivalentTo(200))
				Expect(*resultPool.Name).To(BeEquivalentTo("testPool"))
				Expect(*resultPool.Description).To(BeEquivalentTo("creating pool"))
				Expect(*resultPool.Enabled).To(BeEquivalentTo(true))
				Expect(*resultPool.HealthyOriginsThreshold).To(BeEquivalentTo(1))
				Expect(resultPool.HealthcheckSubnets).To(BeEquivalentTo([]string{subnetCrn}))
				Expect(len(resultPool.HealthcheckVsis)).To(BeIdenticalTo(1))
				Expect(*resultPool.HealthcheckVsis[0].Vpc).To(BeEquivalentTo(vpcCrn))
				Expect(*resultPool.HealthcheckVsis[0].Subnet).To(BeEquivalentTo(subnetCrn))

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
				Expect(getPoolResult.HealthcheckSubnets).To(BeEquivalentTo([]string{subnetCrn}))
				Expect(len(getPoolResult.HealthcheckVsis)).To(BeIdenticalTo(1))
				Expect(*getPoolResult.HealthcheckVsis[0].Vpc).To(BeEquivalentTo(vpcCrn))
				Expect(*getPoolResult.HealthcheckVsis[0].Subnet).To(BeEquivalentTo(subnetCrn))

				//Test Update Pool
				updatePoolOpt := service.NewUpdatePoolOptions(instanceID, *resultPool.ID)
				updatePoolOpt.SetName("updatedtestpool")
				updatePoolOpt.SetDescription("updating testPool")
				createPoolOptions.SetMonitor(*updateResult.ID)
				createPoolOptions.SetHealthcheckRegion("us-south")
				createPoolOptions.SetHealthcheckSubnets([]string{subnetCrn})
				updatePoolResult, updatePoolResponse, updatePoolErr := service.UpdatePool(updatePoolOpt)
				Expect(updatePoolErr).To(BeNil())
				Expect(updatePoolResponse).ToNot(BeNil())
				Expect(updatePoolResult).ToNot(BeNil())
				Expect(updatePoolResponse.GetStatusCode()).To(BeEquivalentTo(200))
				Expect(*updatePoolResult.Name).To(BeEquivalentTo("updatedtestpool"))
				Expect(*updatePoolResult.Description).To(BeEquivalentTo("updating testPool"))
				Expect(updatePoolResult.HealthcheckSubnets).To(BeEquivalentTo([]string{subnetCrn}))
				Expect(len(updatePoolResult.HealthcheckVsis)).To(BeIdenticalTo(1))
				Expect(*updatePoolResult.HealthcheckVsis[0].Vpc).To(BeEquivalentTo(vpcCrn))
				Expect(*updatePoolResult.HealthcheckVsis[0].Subnet).To(BeEquivalentTo(subnetCrn))

				//Test Create Load Balancer
				createLoadBalancerOptions := service.NewCreateLoadBalancerOptions(instanceID, *zoneInfo.ID, "testloadbalancer", *resultPool.ID, []string{*resultPool.ID})
				createLoadBalancerOptions.SetDescription("PDNS Load balancer")
				createLoadBalancerOptions.SetEnabled(true)
				createLoadBalancerOptions.SetTTL(120)
				resultLoadbalancer, responseLoadbalancer, reqErrLoadbalancer := service.CreateLoadBalancer(createLoadBalancerOptions)
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
							delResult, response, reqErr := service.DeletePermittedNetwork(deletePermittedNetworkOptions)
							Expect(reqErr).To(BeNil())
							Expect(response).ToNot(BeNil())
							Expect(delResult).ToNot(BeNil())
						}

						for _, nw := range results.PermittedNetworks {

							for i := 30; i > 0; i-- {
								getPermittedNetworkOptions := service.NewGetPermittedNetworkOptions(instanceID, *zone.ID, *nw.ID)
								_, response, reqErr := service.GetPermittedNetwork(getPermittedNetworkOptions)
								if reqErr != nil {
									if response != nil && response.StatusCode == 404 {
										break
									}
									Expect(reqErr).ToNot(BeNil())
								}
								log.Printf("(BeforeEach) Permitted network (%s) delete is pending. will try after 10 sec", *nw.ID)
								time.Sleep(time.Second * 10)
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
				createDnszoneOptions := service.NewCreateDnszoneOptions(instanceID, zoneName)
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
							Expect(reqErr).To(BeNil())
							Expect(response).ToNot(BeNil())
							Expect(results).ToNot(BeNil())
						}
						for _, nw := range results.PermittedNetworks {
							for i := 30; i > 0; i-- {
								getPermittedNetworkOptions := service.NewGetPermittedNetworkOptions(instanceID, *zone.ID, *nw.ID)
								_, response, reqErr := service.GetPermittedNetwork(getPermittedNetworkOptions)
								if reqErr != nil {
									if response != nil && response.StatusCode == 404 {
										break
									}
									Expect(reqErr).ToNot(BeNil())
								}
								log.Printf("(AfterEach) Permitted network (%s) delete is pending. will try after 10 sec", *nw.ID)
								time.Sleep(time.Second * 10)
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
				permittedNwVPCOption, err := service.NewPermittedNetworkVpc(vpcCrn)
				Expect(err).To(BeNil())
				createPermittedNetworkOptions := service.NewCreatePermittedNetworkOptions(instanceID, *zoneInfo.ID, dnssvcsv1.CreatePermittedNetworkOptions_Type_Vpc, permittedNwVPCOption)
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

				for i := 30; i > 0; i-- {
					getPermittedNetworkOptions := service.NewGetPermittedNetworkOptions(instanceID, *zoneInfo.ID, *permittednetworkID)
					_, response, reqErr := service.GetPermittedNetwork(getPermittedNetworkOptions)
					if reqErr != nil {
						if response != nil && response.StatusCode == 404 {
							break
						}
						Expect(reqErr).ToNot(BeNil())
					}
					log.Printf("Permitted network (%s) delete is pending. will try after 10 sec", *permittednetworkID)
					time.Sleep(time.Second * 10)
				}

				// Test Rmove Permitted Network Fail
				fdeletePermittedNetworkOptions := service.NewDeletePermittedNetworkOptions(instanceID, *zoneInfo.ID, "invalid_id")
				_, _, reqErr = service.DeletePermittedNetwork(fdeletePermittedNetworkOptions)
				Expect(reqErr).ToNot(BeNil())
			})
		})
	})

	Describe(`customresolverv1`, func() {
		Context(`customresolverv1`, func() {
			// var zoneInfo *dnssvcsv1.Dnszone
			BeforeEach(func() {
				shouldSkipTest()

				// delete all custom resolvers
				listCustomResolverOptions := service.NewListCustomResolversOptions(instanceID)
				listCustomResolverOptions.SetXCorrelationID("abc1234")
				Expect(listCustomResolverOptions).ToNot(BeNil())
				resultList, responseList, errList := service.ListCustomResolvers(listCustomResolverOptions)
				Expect(errList).To(BeNil())
				Expect(responseList).ToNot(BeNil())
				Expect(resultList).ToNot(BeNil())

				for i := range resultList.CustomResolvers {
					deleteCustomResolverOptionsModel := service.NewDeleteCustomResolverOptions(instanceID, *resultList.CustomResolvers[i].ID)
					deleteCustomResolverOptionsModel.SetXCorrelationID("abc12387")
					Expect(deleteCustomResolverOptionsModel).ToNot(BeNil())
					responseDel, errDel := service.DeleteCustomResolver(deleteCustomResolverOptionsModel)
					Expect(errDel).To(BeNil())
					Expect(responseDel).ToNot(BeNil())
					Expect(responseDel.GetStatusCode()).To(BeEquivalentTo(204))
				}
			})
			AfterEach(func() {
				shouldSkipTest()
				// delete all custom resolvers
				listCustomResolverOptions := service.NewListCustomResolversOptions(instanceID)
				listCustomResolverOptions.SetXCorrelationID("abc12387")
				Expect(listCustomResolverOptions).ToNot(BeNil())
				resultList, responseList, errList := service.ListCustomResolvers(listCustomResolverOptions)
				Expect(errList).To(BeNil())
				Expect(responseList).ToNot(BeNil())
				Expect(resultList).ToNot(BeNil())

				for i := range resultList.CustomResolvers {
					deleteCustomResolverOptionsModel := service.NewDeleteCustomResolverOptions(instanceID, *resultList.CustomResolvers[i].ID)
					deleteCustomResolverOptionsModel.SetXCorrelationID("abc12387")
					Expect(deleteCustomResolverOptionsModel).ToNot(BeNil())
					responseDel, errDel := service.DeleteCustomResolver(deleteCustomResolverOptionsModel)
					Expect(errDel).To(BeNil())
					Expect(responseDel).ToNot(BeNil())
					Expect(responseDel.GetStatusCode()).To(BeEquivalentTo(204))
				}
			})
			It(`create/list/update/delete/get  custom resolver/custom resolver location/forwarding rule`, func() {
				shouldSkipTest()
				locationInputModel := new(dnssvcsv1.LocationInput)
				locationInputModel.SubnetCrn = core.StringPtr(subnetCrn)
				locationInputModel.Enabled = core.BoolPtr(false)

				createCustomResolverOptions := service.NewCreateCustomResolverOptions(instanceID, "test-resolver1")
				createCustomResolverOptions.SetDescription("Integration test resolver")
				createCustomResolverOptions.SetXCorrelationID("abc12387")
				createCustomResolverOptions.SetLocations([]dnssvcsv1.LocationInput{*locationInputModel})
				createCustomResolverOptions.SetProfile("essential")

				result, response, err := service.CreateCustomResolver(createCustomResolverOptions)
				locationId := result.Locations[0].ID
				Expect(err).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
				Expect(response.GetStatusCode()).To(BeEquivalentTo(200))

				// Test ListAll Custom Resolver
				listCustomResolverOptions := service.NewListCustomResolversOptions(instanceID)
				listCustomResolverOptions.SetXCorrelationID("abc12387")
				Expect(listCustomResolverOptions).ToNot(BeNil())
				resultList, responseList, errList := service.ListCustomResolvers(listCustomResolverOptions)
				Expect(errList).To(BeNil())
				Expect(responseList).ToNot(BeNil())
				Expect(resultList).ToNot(BeNil())
				var customResolverIDs []string
				for i := range resultList.CustomResolvers {
					customResolverIDs = append(customResolverIDs, *resultList.CustomResolvers[i].ID)
				}

				// Test Get Custom Resolver
				getCustomResolverOptionsModel := service.NewGetCustomResolverOptions(instanceID, customResolverIDs[0])
				getCustomResolverOptionsModel.SetXCorrelationID("abc12387")
				Expect(getCustomResolverOptionsModel).ToNot(BeNil())
				resultGet, responseGet, errGet := service.GetCustomResolver(getCustomResolverOptionsModel)
				Expect(errGet).To(BeNil())
				Expect(responseGet).ToNot(BeNil())
				Expect(resultGet).ToNot(BeNil())
				Expect(*resultGet.ID).To(Equal(customResolverIDs[0]))

				// Test Update Custom Resolver
				updateCustomResolverOptionsModel := service.NewUpdateCustomResolverOptions(instanceID, customResolverIDs[0])
				updateCustomResolverOptionsModel.SetName("my-resolver")
				updateCustomResolverOptionsModel.SetDescription("custom resolver2")
				updateCustomResolverOptionsModel.SetEnabled(false)
				updateCustomResolverOptionsModel.SetXCorrelationID("abc12387")
				updateCustomResolverOptionsModel.SetProfile("essential")
				updateCustomResolverOptionsModel.SetAllowDisruptiveUpdates(true)
				Expect(updateCustomResolverOptionsModel).ToNot(BeNil())
				resultUpdate, responseUpdate, errUpdate := service.UpdateCustomResolver(updateCustomResolverOptionsModel)
				Expect(errUpdate).To(BeNil())
				Expect(responseUpdate).ToNot(BeNil())
				Expect(resultUpdate).ToNot(BeNil())
				Expect(responseUpdate.GetStatusCode()).To(BeEquivalentTo(200))
				Expect(*resultUpdate.ID).To(Equal(customResolverIDs[0]))
				Expect(*result.Profile).To(Equal("essential"))

				//Test Update the locations order of Custom Resolver
				updateCustomResolverLocationsOrderOptionsModel := service.NewUpdateCrLocationsOrderOptions(instanceID, customResolverIDs[0], []string{*locationId})
				updateCustomResolverLocationsOrderOptionsModel.SetXCorrelationID("abc12387")
				Expect(updateCustomResolverLocationsOrderOptionsModel).ToNot(BeNil())

				resultCrUpdate, responseCrUpdate, errCrUpdate := service.UpdateCrLocationsOrder(updateCustomResolverLocationsOrderOptionsModel)
				Expect(errCrUpdate).To(BeNil())
				Expect(responseCrUpdate).ToNot(BeNil())
				Expect(resultCrUpdate).ToNot(BeNil())
				Expect(responseCrUpdate.GetStatusCode()).To(BeEquivalentTo(200))
				Expect(*resultCrUpdate.ID).To(Equal(customResolverIDs[0]))

				// Test Add Custom Resolver Location
				addCustomResolverLocationOptionsModel := service.NewAddCustomResolverLocationOptions(instanceID, customResolverIDs[0], subnetCrn)
				addCustomResolverLocationOptionsModel.SetEnabled(false)
				addCustomResolverLocationOptionsModel.SetXCorrelationID("abc12387")
				Expect(addCustomResolverLocationOptionsModel).ToNot(BeNil())

				resAdd, responseAdd, errAdd := service.AddCustomResolverLocation(addCustomResolverLocationOptionsModel)

				Expect(errAdd).To(BeNil())
				Expect(responseAdd).ToNot(BeNil())
				Expect(resAdd).ToNot(BeNil())
				Expect(responseAdd.GetStatusCode()).To(BeEquivalentTo(200))
				Expect(resAdd.DnsServerIp).ToNot(BeNil())

				locationID := resAdd.ID
				addCustomResolverLocationOptionsModel = service.NewAddCustomResolverLocationOptions(instanceID, customResolverIDs[0], customCrn)
				addCustomResolverLocationOptionsModel.SetEnabled(false)
				addCustomResolverLocationOptionsModel.SetXCorrelationID("abc12387")
				Expect(addCustomResolverLocationOptionsModel).ToNot(BeNil())

				resAdd, responseAdd, errAdd = service.AddCustomResolverLocation(addCustomResolverLocationOptionsModel)
				Expect(errAdd).To(BeNil())
				Expect(responseAdd).ToNot(BeNil())
				Expect(resAdd).ToNot(BeNil())
				Expect(responseAdd.GetStatusCode()).To(BeEquivalentTo(200))
				Expect(resAdd.DnsServerIp).ToNot(BeNil())

				// Test ListAll Custom Resolver
				listCustomResolverOptions = service.NewListCustomResolversOptions(instanceID)
				listCustomResolverOptions.SetXCorrelationID("abc12387")
				Expect(listCustomResolverOptions).ToNot(BeNil())
				resultList, responseList, errList = service.ListCustomResolvers(listCustomResolverOptions)
				Expect(errList).To(BeNil())
				Expect(responseList).ToNot(BeNil())
				Expect(resultList).ToNot(BeNil())

				//Test Update Custom Resolver Location

				updateCustomResolverLocationOptionsModel := service.NewUpdateCustomResolverLocationOptions(instanceID, customResolverIDs[0], *locationID)
				updateCustomResolverLocationOptionsModel.SetSubnetCrn(subnetCrn)
				updateCustomResolverLocationOptionsModel.SetEnabled(false)
				updateCustomResolverLocationOptionsModel.SetXCorrelationID("abc12387")
				Expect(updateCustomResolverLocationOptionsModel).ToNot(BeNil())
				resultUp, responseUp, errUp := service.UpdateCustomResolverLocation(updateCustomResolverLocationOptionsModel)
				Expect(errUp).To(BeNil())
				Expect(responseUp).ToNot(BeNil())
				Expect(resultUp).ToNot(BeNil())
				Expect(responseUp.GetStatusCode()).To(BeEquivalentTo(200))
				Expect(resultUp.DnsServerIp).ToNot(BeNil())

				// Test - Create Forwarding Rule

				var forwardingRuleInput dnssvcsv1.ForwardingRuleInputIntf = nil
				createForwardingRuleOptionsModel := service.NewCreateForwardingRuleOptions(instanceID, customResolverIDs[0], forwardingRuleInput)
				createForwardingRuleOptionsModel.SetXCorrelationID("abc12387")

				viewConfigModel := new(dnssvcsv1.ViewConfig)
				viewConfigModel.Name = core.StringPtr("view-example")
				viewConfigModel.Description = core.StringPtr("view example")
				viewConfigModel.Expression = core.StringPtr("ipInRange(source.ip, '10.240.0.0/24') || ipInRange(source.ip, '10.240.1.0/24')")
				viewConfigModel.ForwardTo = []string{"10.240.2.6"}

				// testing Forwarding rule where both views and forward_to is required
				forwardingRuleInputModelBoth := new(dnssvcsv1.ForwardingRuleInputForwardingRuleBoth)
				forwardingRuleInputModelBoth.Description = core.StringPtr("forwarding rule")
				forwardingRuleInputModelBoth.Type = core.StringPtr("zone")
				forwardingRuleInputModelBoth.Match = core.StringPtr("example.com")
				forwardingRuleInputModelBoth.ForwardTo = []string{"161.26.0.7"}
				forwardingRuleInputModelBoth.Views = []dnssvcsv1.ViewConfig{*viewConfigModel}

				createForwardingRuleOptionsModel.SetForwardingRuleInput(forwardingRuleInputModelBoth)
				Expect(createForwardingRuleOptionsModel).ToNot(BeNil())
				resultCreate, responseCreate, errCreate := service.CreateForwardingRule(createForwardingRuleOptionsModel)
				Expect(errCreate).To(BeNil())
				Expect(responseCreate).ToNot(BeNil())
				Expect(resultCreate).ToNot(BeNil())
				Expect(responseCreate.StatusCode).To(BeEquivalentTo(200))
				Expect(resultCreate.ID).ToNot(BeNil())
				forwardingRulesID := resultCreate.ID

				// testing Forwarding rule where only forward_to is required
				forwardingRuleInputModelOnlyForward := new(dnssvcsv1.ForwardingRuleInputForwardingRuleOnlyForward)
				forwardingRuleInputModelOnlyForward.Description = core.StringPtr("forwarding rule")
				forwardingRuleInputModelOnlyForward.Type = core.StringPtr("zone")
				forwardingRuleInputModelOnlyForward.Match = core.StringPtr("example1.com")
				forwardingRuleInputModelOnlyForward.ForwardTo = []string{"161.26.0.7"}

				createForwardingRuleOptionsModel.SetForwardingRuleInput(forwardingRuleInputModelOnlyForward)
				Expect(createForwardingRuleOptionsModel).ToNot(BeNil())
				resultCreate, responseCreate, errCreate = service.CreateForwardingRule(createForwardingRuleOptionsModel)
				Expect(errCreate).To(BeNil())
				Expect(responseCreate).ToNot(BeNil())
				Expect(resultCreate).ToNot(BeNil())
				Expect(responseCreate.StatusCode).To(BeEquivalentTo(200))
				Expect(resultCreate.ID).ToNot(BeNil())

				// testing Forwarding rule where only view is required
				forwardingRuleInputModelOnlyView := new(dnssvcsv1.ForwardingRuleInputForwardingRuleOnlyView)
				forwardingRuleInputModelOnlyView.Description = core.StringPtr("forwarding rule")
				forwardingRuleInputModelOnlyView.Type = core.StringPtr("zone")
				forwardingRuleInputModelOnlyView.Match = core.StringPtr("example2.com")
				forwardingRuleInputModelOnlyView.Views = []dnssvcsv1.ViewConfig{*viewConfigModel}

				createForwardingRuleOptionsModel.SetForwardingRuleInput(forwardingRuleInputModelOnlyView)
				Expect(createForwardingRuleOptionsModel).ToNot(BeNil())
				resultCreate, responseCreate, errCreate = service.CreateForwardingRule(createForwardingRuleOptionsModel)
				Expect(errCreate).To(BeNil())
				Expect(responseCreate).ToNot(BeNil())
				Expect(resultCreate).ToNot(BeNil())
				Expect(responseCreate.StatusCode).To(BeEquivalentTo(200))
				Expect(resultCreate.ID).ToNot(BeNil())

				// List Forwarding Rules
				listForwardingRulesOptionsModel := service.NewListForwardingRulesOptions(instanceID, customResolverIDs[0])
				listForwardingRulesOptionsModel.SetXCorrelationID("abc12387")
				Expect(listForwardingRulesOptionsModel).ToNot(BeNil())
				resultFRList, responseFRList, errFRList := service.ListForwardingRules(listForwardingRulesOptionsModel)
				Expect(errFRList).To(BeNil())
				Expect(responseFRList).ToNot(BeNil())
				Expect(resultFRList).ToNot(BeNil())
				// Test Get a Forwarding Rule
				getForwardingRuleOptionsModel := service.NewGetForwardingRuleOptions(instanceID, customResolverIDs[0], *forwardingRulesID)
				getForwardingRuleOptionsModel.SetXCorrelationID("testString")
				Expect(getForwardingRuleOptionsModel).ToNot(BeNil())
				resultFRGet, responseFRGet, errFRGet := service.GetForwardingRule(getForwardingRuleOptionsModel)
				Expect(errFRGet).To(BeNil())
				Expect(responseFRGet).ToNot(BeNil())
				Expect(resultFRGet).ToNot(BeNil())
				Expect(responseFRGet.StatusCode).To(BeEquivalentTo(200))
				Expect(resultFRGet.ID).ToNot(BeNil())

				// Test Update a Forwarding Rule
				updateForwardingRuleOptionsModel := service.NewUpdateForwardingRuleOptions(instanceID, customResolverIDs[0], *forwardingRulesID)
				updateForwardingRuleOptionsModel.SetXCorrelationID("testString")
				updateForwardingRuleOptionsModel.SetDescription("cli test forwarding rule")
				updateForwardingRuleOptionsModel.SetMatch("test.example.com")
				updateForwardingRuleOptionsModel.SetForwardTo([]string{"161.26.8.8"})
				updateForwardingRuleOptionsModel.SetXCorrelationID("testString")
				updateForwardingRuleOptionsModel.SetViews([]dnssvcsv1.ViewConfig{*viewConfigModel})

				Expect(updateForwardingRuleOptionsModel).ToNot(BeNil())
				resultFRUpdate, responseFRUpdate, errFRUpdate := service.UpdateForwardingRule(updateForwardingRuleOptionsModel)
				Expect(errFRUpdate).To(BeNil())
				Expect(responseFRUpdate).ToNot(BeNil())
				Expect(resultFRUpdate).ToNot(BeNil())
				Expect(responseFRUpdate.StatusCode).To(BeEquivalentTo(200))
				Expect(resultFRUpdate.ID).ToNot(BeNil())

				// Test Delete a Forwarding Rule
				deleteForwardingRuleOptionsModel := service.NewDeleteForwardingRuleOptions(instanceID, customResolverIDs[0], *forwardingRulesID)
				deleteForwardingRuleOptionsModel.SetXCorrelationID("testString")
				Expect(deleteForwardingRuleOptionsModel).ToNot(BeNil())
				responseFRDelete, errFRDelete := service.DeleteForwardingRule(deleteForwardingRuleOptionsModel)
				Expect(errFRDelete).To(BeNil())
				Expect(responseFRDelete).ToNot(BeNil())
				Expect(responseFRDelete.StatusCode).To(BeEquivalentTo(204))

				// Test Delete Custom Resolver Location
				deleteCustomResolverLocationOptionsModel := service.NewDeleteCustomResolverLocationOptions(instanceID, customResolverIDs[0], *locationID)
				deleteCustomResolverLocationOptionsModel.SetXCorrelationID("abc12387")

				Expect(deleteCustomResolverLocationOptionsModel).ToNot(BeNil())
				responseDelete, errDelete := service.DeleteCustomResolverLocation(deleteCustomResolverLocationOptionsModel)
				Expect(errDelete).To(BeNil())
				Expect(responseDelete).ToNot(BeNil())
				Expect(responseDelete.GetStatusCode()).To(BeEquivalentTo(204))

				// Test Delete Custom Resolver
				deleteCustomResolverOptionsModel := service.NewDeleteCustomResolverOptions(instanceID, customResolverIDs[0])
				deleteCustomResolverOptionsModel.SetXCorrelationID("abc12387")
				Expect(deleteCustomResolverOptionsModel).ToNot(BeNil())
				responseDel, errDel := service.DeleteCustomResolver(deleteCustomResolverOptionsModel)
				Expect(errDel).To(BeNil())
				Expect(responseDel).ToNot(BeNil())
				Expect(responseDel.GetStatusCode()).To(BeEquivalentTo(204))
			})
		})
		Context(`customresolverv1 pagination params`, func() {
			var customResolverID string
			var forwardingRulesID [3]*string
			BeforeEach(func() {
				shouldSkipTest()
				// delete all custom resolvers
				listCustomResolverOptions := service.NewListCustomResolversOptions(instanceID)
				listCustomResolverOptions.SetXCorrelationID("abc1234")
				Expect(listCustomResolverOptions).ToNot(BeNil())
				resultList, responseList, errList := service.ListCustomResolvers(listCustomResolverOptions)
				Expect(errList).To(BeNil())
				Expect(responseList).ToNot(BeNil())
				Expect(resultList).ToNot(BeNil())
				for i := range resultList.CustomResolvers {
					deleteCustomResolverOptionsModel := service.NewDeleteCustomResolverOptions(instanceID, *resultList.CustomResolvers[i].ID)
					deleteCustomResolverOptionsModel.SetXCorrelationID("abc12387")
					Expect(deleteCustomResolverOptionsModel).ToNot(BeNil())
					responseDel, errDel := service.DeleteCustomResolver(deleteCustomResolverOptionsModel)
					Expect(errDel).To(BeNil())
					Expect(responseDel).ToNot(BeNil())
					Expect(responseDel.GetStatusCode()).To(BeEquivalentTo(204))
				}
				// Create a CR
				locationInputModel := new(dnssvcsv1.LocationInput)
				locationInputModel.SubnetCrn = core.StringPtr(subnetCrn)
				locationInputModel.Enabled = core.BoolPtr(false)
				createCustomResolverOptions := service.NewCreateCustomResolverOptions(instanceID, "test-resolver")
				createCustomResolverOptions.SetDescription("Integration test resolver")
				createCustomResolverOptions.SetXCorrelationID("abc12387")
				createCustomResolverOptions.SetLocations([]dnssvcsv1.LocationInput{*locationInputModel})
				createCustomResolverOptions.SetProfile("essential")

				result, response, err := service.CreateCustomResolver(createCustomResolverOptions)
				Expect(err).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
				Expect(response.GetStatusCode()).To(BeEquivalentTo(200))
				customResolverID = *result.ID
				// Create forwarding rules
				for i := 1; i <= 3; i++ {
					var forwardingRuleInput dnssvcsv1.ForwardingRuleInputIntf = nil
					createForwardingRuleOptionsModel := service.NewCreateForwardingRuleOptions(instanceID, customResolverID, forwardingRuleInput)

					viewConfigModel := new(dnssvcsv1.ViewConfig)
					viewConfigModel.Name = core.StringPtr("view-example")
					viewConfigModel.Description = core.StringPtr("view example")
					viewConfigModel.Expression = core.StringPtr("ipInRange(source.ip, '10.240.0.0/24') || ipInRange(source.ip, '10.240.1.0/24')")
					viewConfigModel.ForwardTo = []string{"10.240.2.6"}

					forwardingRuleInputModel := new(dnssvcsv1.ForwardingRuleInputForwardingRuleBoth)
					forwardingRuleInputModel.Description = core.StringPtr("test forwarding rule " + strconv.Itoa(i))
					forwardingRuleInputModel.Type = core.StringPtr("zone")
					forwardingRuleInputModel.Match = core.StringPtr(strconv.Itoa(i) + "example.com")
					forwardingRuleInputModel.ForwardTo = []string{"161.26.0.7"}
					forwardingRuleInputModel.Views = []dnssvcsv1.ViewConfig{*viewConfigModel}

					createForwardingRuleOptionsModel.SetForwardingRuleInput(forwardingRuleInputModel)
					createForwardingRuleOptionsModel.SetXCorrelationID("abc12387")
					Expect(createForwardingRuleOptionsModel).ToNot(BeNil())
					resultCreate, responseCreate, errCreate := service.CreateForwardingRule(createForwardingRuleOptionsModel)
					Expect(errCreate).To(BeNil())
					Expect(responseCreate).ToNot(BeNil())
					Expect(resultCreate).ToNot(BeNil())
					Expect(responseCreate.StatusCode).To(BeEquivalentTo(200))
					Expect(resultCreate.ID).ToNot(BeNil())
					Expect(resultCreate.ForwardTo[0]).To(Equal("161.26.0.7"))
					forwardingRulesID[i-1] = resultCreate.ID

				}

			})
			AfterEach(func() {
				shouldSkipTest()
				// delete all custom resolvers
				listCustomResolverOptions := service.NewListCustomResolversOptions(instanceID)
				listCustomResolverOptions.SetXCorrelationID("abc12387")
				Expect(listCustomResolverOptions).ToNot(BeNil())
				resultList, responseList, errList := service.ListCustomResolvers(listCustomResolverOptions)
				Expect(errList).To(BeNil())
				Expect(responseList).ToNot(BeNil())
				Expect(resultList).ToNot(BeNil())

				// Delete forwarding rules
				for i := 0; i <= 2; i++ {
					deleteForwardingRuleOptionsModel := service.NewDeleteForwardingRuleOptions(instanceID, customResolverID, *forwardingRulesID[i])
					deleteForwardingRuleOptionsModel.SetXCorrelationID("testString")
					Expect(deleteForwardingRuleOptionsModel).ToNot(BeNil())
					responseFRDelete, errFRDelete := service.DeleteForwardingRule(deleteForwardingRuleOptionsModel)
					Expect(errFRDelete).To(BeNil())
					Expect(responseFRDelete).ToNot(BeNil())
					Expect(responseFRDelete.StatusCode).To(BeEquivalentTo(204))
				}

				for i := range resultList.CustomResolvers {
					deleteCustomResolverOptionsModel := service.NewDeleteCustomResolverOptions(instanceID, *resultList.CustomResolvers[i].ID)
					deleteCustomResolverOptionsModel.SetXCorrelationID("abc12387")
					Expect(deleteCustomResolverOptionsModel).ToNot(BeNil())
					responseDel, errDel := service.DeleteCustomResolver(deleteCustomResolverOptionsModel)
					Expect(errDel).To(BeNil())
					Expect(responseDel).ToNot(BeNil())
					Expect(responseDel.GetStatusCode()).To(BeEquivalentTo(204))
				}
			})
			Context(`customresolverv1 pagination params`, func() {
				It(`customresolverv1 pagination offset`, func() {
					// Test offset(skip) for forwarding rules
					listForwardingRulesOptionsNew := service.NewListForwardingRulesOptions(instanceID, customResolverID)
					Expect(listForwardingRulesOptionsNew).ToNot(BeNil())
					listForwardingRulesOptionsNew.SetXCorrelationID("abc12387")
					listForwardingRulesOptionsNew.SetOffset(2)
					resultFRList1, responseFRList1, errFRList1 := service.ListForwardingRules(listForwardingRulesOptionsNew)
					Expect(errFRList1).To(BeNil())
					Expect(responseFRList1).ToNot(BeNil())
					Expect(resultFRList1).ToNot(BeNil())
					Expect(len(resultFRList1.ForwardingRules)).To(BeEquivalentTo(2))
					Expect(*resultFRList1.Count).To(BeEquivalentTo(2))
					Expect(*resultFRList1.TotalCount).To(BeEquivalentTo(4))
					Expect(*resultFRList1.Limit).To(BeEquivalentTo(200))
					Expect(*resultFRList1.Offset).To(BeEquivalentTo(2))
					href1 := fmt.Sprintf("/%s/forwarding_rules?offset=0&limit=200", customResolverID)
					Expect(strings.Contains(*resultFRList1.First.Href, href1)).To(BeTrue())
					Expect(strings.Contains(*resultFRList1.Last.Href, href1)).To(BeTrue())
					Expect(resultFRList1.Next).To(BeNil())

					// Skip 2 rules
					listForwardingRulesOptionsNew1 := service.NewListForwardingRulesOptions(instanceID, customResolverID)
					listForwardingRulesOptionsNew1.SetXCorrelationID("abc12387")
					listForwardingRulesOptionsNew1.SetOffset(3)
					Expect(listForwardingRulesOptionsNew1).ToNot(BeNil())
					resultFRList2, responseFRList2, errFRList2 := service.ListForwardingRules(listForwardingRulesOptionsNew1)
					Expect(errFRList2).To(BeNil())
					Expect(responseFRList2).ToNot(BeNil())
					Expect(resultFRList2).ToNot(BeNil())
					Expect(len(resultFRList2.ForwardingRules)).To(BeEquivalentTo(1))
					Expect(*resultFRList2.Count).To(BeEquivalentTo(1))
					Expect(*resultFRList2.TotalCount).To(BeEquivalentTo(4))
					Expect(*resultFRList2.Limit).To(BeEquivalentTo(200))
					Expect(*resultFRList2.Offset).To(BeEquivalentTo(3))
					href2 := fmt.Sprintf("/%s/forwarding_rules?offset=0&limit=200", customResolverID)
					Expect(strings.Contains(*resultFRList2.First.Href, href2)).To(BeTrue())
					Expect(strings.Contains(*resultFRList2.Last.Href, href2)).To(BeTrue())
					Expect(resultFRList2.Next).To(BeNil())

				})
				It(`customresolverv1 pagination limit`, func() {
					// Set limit 1
					listForwardingRulesOptionsNew2 := service.NewListForwardingRulesOptions(instanceID, customResolverID)
					Expect(listForwardingRulesOptionsNew2).ToNot(BeNil())
					listForwardingRulesOptionsNew2.SetXCorrelationID("abc12387")
					listForwardingRulesOptionsNew2.SetLimit(1)
					resultFRList3, responseFRList3, errFRList3 := service.ListForwardingRules(listForwardingRulesOptionsNew2)
					Expect(errFRList3).To(BeNil())
					Expect(responseFRList3).ToNot(BeNil())
					Expect(resultFRList3).ToNot(BeNil())
					Expect(len(resultFRList3.ForwardingRules)).To(BeEquivalentTo(1))
					// Set limit 3
					listForwardingRulesOptionsNew3 := service.NewListForwardingRulesOptions(instanceID, customResolverID)
					listForwardingRulesOptionsNew3.SetXCorrelationID("abc12387")
					listForwardingRulesOptionsNew3.SetLimit(3)
					Expect(listForwardingRulesOptionsNew3).ToNot(BeNil())
					resultFRList4, responseFRList4, errFRList4 := service.ListForwardingRules(listForwardingRulesOptionsNew3)
					Expect(errFRList4).To(BeNil())
					Expect(responseFRList4).ToNot(BeNil())
					Expect(resultFRList4).ToNot(BeNil())
					Expect(len(resultFRList4.ForwardingRules)).To(BeEquivalentTo(3))
					Expect(*resultFRList4.Count).To(BeEquivalentTo(3))
					Expect(*resultFRList4.TotalCount).To(BeEquivalentTo(4))
					Expect(*resultFRList4.Limit).To(BeEquivalentTo(3))
					Expect(*resultFRList4.Offset).To(BeEquivalentTo(0))
					// Case where we have last page which is different than first page
					href4 := fmt.Sprintf("/%s/forwarding_rules?offset=0&limit=3", customResolverID)
					href41 := fmt.Sprintf("/%s/forwarding_rules?offset=1&limit=3", customResolverID)
					Expect(strings.Contains(*resultFRList4.First.Href, href4)).To(BeTrue())
					Expect(strings.Contains(*resultFRList4.Last.Href, href41)).To(BeTrue())
					// We should have next page here
					href5 := fmt.Sprintf("/%s/forwarding_rules?offset=3&limit=3", customResolverID)
					Expect(strings.Contains(*resultFRList4.Next.Href, href5)).To(BeTrue())
				})
				It(`customresolverv1 pagination limit & offset`, func() {
					// Set limit 3 and offset 3
					listForwardingRulesOptionsNew2 := service.NewListForwardingRulesOptions(instanceID, customResolverID)
					Expect(listForwardingRulesOptionsNew2).ToNot(BeNil())
					listForwardingRulesOptionsNew2.SetXCorrelationID("abc12387")
					listForwardingRulesOptionsNew2.SetLimit(3)
					listForwardingRulesOptionsNew2.SetOffset(3)
					resultFRList3, responseFRList3, errFRList3 := service.ListForwardingRules(listForwardingRulesOptionsNew2)
					Expect(errFRList3).To(BeNil())
					Expect(responseFRList3).ToNot(BeNil())
					Expect(resultFRList3).ToNot(BeNil())
					Expect(len(resultFRList3.ForwardingRules)).To(BeEquivalentTo(1))

					// Set limit and offset
					listForwardingRulesOptionsNew3 := service.NewListForwardingRulesOptions(instanceID, customResolverID)
					listForwardingRulesOptionsNew3.SetXCorrelationID("abc12387")
					listForwardingRulesOptionsNew3.SetLimit(2)
					listForwardingRulesOptionsNew3.SetOffset(1)
					Expect(listForwardingRulesOptionsNew3).ToNot(BeNil())
					resultFRList4, responseFRList4, errFRList4 := service.ListForwardingRules(listForwardingRulesOptionsNew3)
					Expect(errFRList4).To(BeNil())
					Expect(responseFRList4).ToNot(BeNil())
					Expect(resultFRList4).ToNot(BeNil())
					Expect(len(resultFRList4.ForwardingRules)).To(BeEquivalentTo(2))
					Expect(*resultFRList4.Count).To(BeEquivalentTo(2))
					Expect(*resultFRList4.TotalCount).To(BeEquivalentTo(4))
					Expect(*resultFRList4.Limit).To(BeEquivalentTo(2))
					Expect(*resultFRList4.Offset).To(BeEquivalentTo(1))
					href5 := fmt.Sprintf("/%s/forwarding_rules?offset=0&limit=2", customResolverID)
					href51 := fmt.Sprintf("/%s/forwarding_rules?offset=2&limit=2", customResolverID)
					Expect(strings.Contains(*resultFRList4.First.Href, href5)).To(BeTrue())
					Expect(strings.Contains(*resultFRList4.Last.Href, href51)).To(BeTrue())
					// We should have next page here
					href6 := fmt.Sprintf("/%s/forwarding_rules?offset=3&limit=2", customResolverID)
					Expect(strings.Contains(*resultFRList4.Next.Href, href6)).To(BeTrue())
				})
			})
		})
	})
	Describe(`secondaryzonev1`, func() {
		Context(`secondaryzonev1`, func() {
			var testCustomResolver *dnssvcsv1.CustomResolver
			BeforeEach(func() {
				shouldSkipTest()

				// delete all custom resolvers
				listCustomResolverOptions := service.NewListCustomResolversOptions(instanceID)
				listCustomResolverOptions.SetXCorrelationID("secondaryzone123")
				Expect(listCustomResolverOptions).ToNot(BeNil())
				resultList, responseList, errList := service.ListCustomResolvers(listCustomResolverOptions)
				Expect(errList).To(BeNil())
				Expect(responseList).ToNot(BeNil())
				Expect(resultList).ToNot(BeNil())

				for i := range resultList.CustomResolvers {
					deleteCustomResolverOptionsModel := service.NewDeleteCustomResolverOptions(instanceID, *resultList.CustomResolvers[i].ID)
					deleteCustomResolverOptionsModel.SetXCorrelationID("secondaryzone123")
					Expect(deleteCustomResolverOptionsModel).ToNot(BeNil())
					responseDel, errDel := service.DeleteCustomResolver(deleteCustomResolverOptionsModel)
					Expect(errDel).To(BeNil())
					Expect(responseDel).ToNot(BeNil())
					Expect(responseDel.GetStatusCode()).To(BeEquivalentTo(204))
				}
				locationInputModel := new(dnssvcsv1.LocationInput)
				locationInputModel.SubnetCrn = core.StringPtr(subnetCrn)
				locationInputModel.Enabled = core.BoolPtr(false)

				// create test custom resolver
				createCustomResolverOptions := service.NewCreateCustomResolverOptions(instanceID, "secondaryzone-test-resolver1")
				createCustomResolverOptions.SetDescription("Integration test resolver")
				createCustomResolverOptions.SetXCorrelationID("secondaryzone123")
				createCustomResolverOptions.SetLocations([]dnssvcsv1.LocationInput{*locationInputModel})
				createCustomResolverOptions.SetProfile("essential")

				resultCreateCustomResolver, responseCreateCustomResolver, err := service.CreateCustomResolver(createCustomResolverOptions)
				Expect(err).To(BeNil())
				Expect(responseCreateCustomResolver).ToNot(BeNil())
				Expect(responseCreateCustomResolver.GetStatusCode()).To(BeEquivalentTo(200))
				Expect(resultCreateCustomResolver).ToNot(BeNil())

				// set test custom resolver
				testCustomResolver = resultCreateCustomResolver
			})
			AfterEach(func() {
				shouldSkipTest()

				// check that custom resolver exists
				getCustomResolverOptions := service.NewGetCustomResolverOptions(instanceID, *testCustomResolver.ID)
				getCustomResolverOptions.SetXCorrelationID("secondaryzone123")
				Expect(getCustomResolverOptions).ToNot(BeNil())
				resultGet, responseGet, errList := service.GetCustomResolver(getCustomResolverOptions)
				Expect(errList).To(BeNil())
				Expect(responseGet).ToNot(BeNil())
				Expect(resultGet).ToNot(BeNil())

				// list secondary zones for custom resolver
				listSecondaryZonesOptions := service.NewListSecondaryZonesOptions(instanceID, *testCustomResolver.ID)
				listSecondaryZonesOptions.SetXCorrelationID("secondaryzone123")
				Expect(listSecondaryZonesOptions).ToNot(BeNil())
				resultList, responseList, errList := service.ListSecondaryZones(listSecondaryZonesOptions)
				Expect(errList).To(BeNil())
				Expect(responseList).ToNot(BeNil())
				Expect(resultList).ToNot(BeNil())

				// delete secondary zones
				for i := range resultList.SecondaryZones {
					deleteSecondaryZoneOptions :=
						service.NewDeleteSecondaryZoneOptions(
							instanceID,
							*testCustomResolver.ID,
							*resultList.SecondaryZones[i].ID,
						)
					deleteSecondaryZoneOptions.SetXCorrelationID("secondaryzone123")
					Expect(deleteSecondaryZoneOptions).ToNot(BeNil())
					responseDel, errDel := service.DeleteSecondaryZone(deleteSecondaryZoneOptions)
					Expect(errDel).To(BeNil())
					Expect(responseDel).ToNot(BeNil())
					Expect(responseDel.GetStatusCode()).To(BeEquivalentTo(204))
				}

				// delete test custom resolver
				deleteCustomResolverOptionsModel := service.NewDeleteCustomResolverOptions(instanceID, *testCustomResolver.ID)
				deleteCustomResolverOptionsModel.SetXCorrelationID("abc12387")
				Expect(deleteCustomResolverOptionsModel).ToNot(BeNil())
				responseDel, errDel := service.DeleteCustomResolver(deleteCustomResolverOptionsModel)
				Expect(errDel).To(BeNil())
				Expect(responseDel).ToNot(BeNil())
				Expect(responseDel.GetStatusCode()).To(BeEquivalentTo(204))
			})
			It(`create/list/update/delete/get secondary zone`, func() {
				shouldSkipTest()

				// create custom resolver
				createSecondaryZoneOptions := service.NewCreateSecondaryZoneOptions(
					instanceID,
					*testCustomResolver.ID,
					"example.com",
					[]string{"10.0.0.7"},
				)
				createSecondaryZoneOptions.SetXCorrelationID("create-secondaryzone123")

				resultCreate, responseCreate, errCreate := service.CreateSecondaryZone(createSecondaryZoneOptions)
				Expect(errCreate).To(BeNil())
				Expect(responseCreate).ToNot(BeNil())
				Expect(resultCreate).ToNot(BeNil())
				Expect(responseCreate.GetStatusCode()).To(BeEquivalentTo(200))

				// list secondary zones for custom resolver
				listSecondaryZonesOptions := service.NewListSecondaryZonesOptions(
					instanceID,
					*testCustomResolver.ID,
				)
				listSecondaryZonesOptions.SetXCorrelationID("list-secondaryzone123")
				Expect(listSecondaryZonesOptions).ToNot(BeNil())
				resultList, responseList, errList := service.ListSecondaryZones(listSecondaryZonesOptions)
				Expect(errList).To(BeNil())
				Expect(responseList).ToNot(BeNil())
				Expect(resultList).ToNot(BeNil())
				Expect(responseCreate.GetStatusCode()).To(BeEquivalentTo(200))

				// get created secondary zone
				getSecondaryZoneOptions := service.NewGetSecondaryZoneOptions(
					instanceID,
					*testCustomResolver.ID,
					*resultCreate.ID,
				)
				getSecondaryZoneOptions.SetXCorrelationID("get-secondaryzone123")
				Expect(getSecondaryZoneOptions).ToNot(BeNil())
				resultGet, responseGet, errGet := service.GetSecondaryZone(getSecondaryZoneOptions)
				Expect(errGet).To(BeNil())
				Expect(responseGet).ToNot(BeNil())
				Expect(resultGet).ToNot(BeNil())
				Expect(responseCreate.GetStatusCode()).To(BeEquivalentTo(200))

				// update created secondary zone
				updateSecondaryZoneOptions := service.NewUpdateSecondaryZoneOptions(
					instanceID,
					*testCustomResolver.ID,
					*resultCreate.ID,
				)
				updateSecondaryZoneOptions.SetXCorrelationID("update-secondaryzone123")
				updateSecondaryZoneOptions.SetDescription("description update")
				Expect(updateSecondaryZoneOptions).ToNot(BeNil())
				resultUpdate, responseUpdate, errUpdate := service.UpdateSecondaryZone(updateSecondaryZoneOptions)
				Expect(errUpdate).To(BeNil())
				Expect(responseUpdate).ToNot(BeNil())
				Expect(responseUpdate.GetStatusCode()).To(BeEquivalentTo(200))
				Expect(resultUpdate).ToNot(BeNil())

				// delete created secondary zone
				deleteSecondaryZoneOptions :=
					service.NewDeleteSecondaryZoneOptions(
						instanceID,
						*testCustomResolver.ID,
						*resultCreate.ID,
					)
				deleteSecondaryZoneOptions.SetXCorrelationID("delete-secondaryzone123")
				Expect(deleteSecondaryZoneOptions).ToNot(BeNil())
				responseDel, errDel := service.DeleteSecondaryZone(deleteSecondaryZoneOptions)
				Expect(errDel).To(BeNil())
				Expect(responseDel).ToNot(BeNil())
				Expect(responseDel.GetStatusCode()).To(BeEquivalentTo(204))

			})
		})
	})

	Describe(`crossaccountsv1`, func() {
		Context(`crossaccountsv1`, func() {
			BeforeEach(func() {
				shouldSkipTest()

				// list linked zones
				listLinkedZonesOptions := service.NewListLinkedZonesOptions(instanceID)
				listLinkedZonesOptions.SetXCorrelationID("listLinkedZones123")
				Expect(listLinkedZonesOptions).ToNot(BeNil())
				resultList, responseList, errList := service.ListLinkedZones(listLinkedZonesOptions)
				Expect(errList).To(BeNil())
				Expect(responseList).ToNot(BeNil())
				Expect(resultList).ToNot(BeNil())

				// delete linked zones
				for _, linkedZone := range resultList.LinkedDnszones {
					// list access request
					listDnszoneAccessRequestsOptions := serviceOwnerDnsInstanceAccount.NewListDnszoneAccessRequestsOptions(
						ownerInstanceID,
						*linkedZone.LinkedTo.ZoneID,
					)
					listDnszoneAccessRequestsOptions.SetXCorrelationID("lzpermittednetworks-listDnszoneaccessrequests-beforeeach")
					resultListAccessRequests, responseListAccessRequests, errListAccessRequests := serviceOwnerDnsInstanceAccount.ListDnszoneAccessRequests(listDnszoneAccessRequestsOptions)
					Expect(errListAccessRequests).To(BeNil())
					Expect(responseListAccessRequests).ToNot(BeNil())
					Expect(responseListAccessRequests.GetStatusCode()).To(BeEquivalentTo(200))
					Expect(resultListAccessRequests).ToNot(BeNil())
					Expect(resultListAccessRequests.AccessRequests).ToNot(BeNil())
					Expect(len(resultListAccessRequests.AccessRequests) > 0).To(BeTrue())

					// update access request
					updateDnszoneAccessRequestsOptions :=
						service.NewUpdateDnszoneAccessRequestOptions(
							ownerInstanceID,
							ownerZoneID,
							*resultListAccessRequests.AccessRequests[0].ID,
							"REVOKE",
						)
					updateDnszoneAccessRequestsOptions.SetXCorrelationID("lzpermittednetworks-updatednszoneaccessrequests-beforeeach")
					resultUpdate, responseUpdate, errUpdate := serviceOwnerDnsInstanceAccount.UpdateDnszoneAccessRequest(updateDnszoneAccessRequestsOptions)
					Expect(errUpdate).To(BeNil())
					Expect(responseUpdate).ToNot(BeNil())
					Expect(responseUpdate.GetStatusCode()).To(BeEquivalentTo(200))
					Expect(resultUpdate).ToNot(BeNil())

					deleteLinkedZonesOptions := service.NewDeleteLinkedZoneOptions(instanceID, *linkedZone.ID)
					response, err := service.DeleteLinkedZone(deleteLinkedZonesOptions)
					Expect(err).To(BeNil())
					Expect(response).ToNot(BeNil())
					Expect(response.GetStatusCode()).To(BeEquivalentTo(204))
				}
			})
			AfterEach(func() {
				shouldSkipTest()

				// list linked zones
				listLinkedZonesOptions := service.NewListLinkedZonesOptions(instanceID)
				listLinkedZonesOptions.SetXCorrelationID("listLinkedZones123")
				Expect(listLinkedZonesOptions).ToNot(BeNil())
				resultList, responseList, errList := service.ListLinkedZones(listLinkedZonesOptions)
				Expect(errList).To(BeNil())
				Expect(responseList).ToNot(BeNil())
				Expect(resultList).ToNot(BeNil())

				// delete linked zones
				for _, linkedZone := range resultList.LinkedDnszones {
					// list access request
					listDnszoneAccessRequestsOptions := serviceOwnerDnsInstanceAccount.NewListDnszoneAccessRequestsOptions(
						ownerInstanceID,
						*linkedZone.LinkedTo.ZoneID,
					)
					listDnszoneAccessRequestsOptions.SetXCorrelationID("lzpermittednetworks-listDnszoneaccessrequests-beforeeach")
					resultListAccessRequests, responseListAccessRequests, errListAccessRequests := serviceOwnerDnsInstanceAccount.ListDnszoneAccessRequests(listDnszoneAccessRequestsOptions)
					Expect(errListAccessRequests).To(BeNil())
					Expect(responseListAccessRequests).ToNot(BeNil())
					Expect(responseListAccessRequests.GetStatusCode()).To(BeEquivalentTo(200))
					Expect(resultListAccessRequests).ToNot(BeNil())
					Expect(resultListAccessRequests.AccessRequests).ToNot(BeNil())
					Expect(len(resultListAccessRequests.AccessRequests) > 0).To(BeTrue())

					// update access request
					updateDnszoneAccessRequestsOptions :=
						service.NewUpdateDnszoneAccessRequestOptions(
							ownerInstanceID,
							ownerZoneID,
							*resultListAccessRequests.AccessRequests[0].ID,
							"REVOKE",
						)
					updateDnszoneAccessRequestsOptions.SetXCorrelationID("lzpermittednetworks-updatednszoneaccessrequests-beforeeach")
					updateDnszoneAccessRequestsOptions.SetAction("REVOKE")
					resultUpdate, responseUpdate, errUpdate := serviceOwnerDnsInstanceAccount.UpdateDnszoneAccessRequest(updateDnszoneAccessRequestsOptions)
					Expect(errUpdate).To(BeNil())
					Expect(responseUpdate).ToNot(BeNil())
					Expect(responseUpdate.GetStatusCode()).To(BeEquivalentTo(200))
					Expect(resultUpdate).ToNot(BeNil())

					// delete linked zone
					deleteLinkedZonesOptions := service.NewDeleteLinkedZoneOptions(instanceID, *linkedZone.ID)
					response, err := service.DeleteLinkedZone(deleteLinkedZonesOptions)
					Expect(err).To(BeNil())
					Expect(response).ToNot(BeNil())
					Expect(response.GetStatusCode()).To(BeEquivalentTo(204))
				}
			})
			It(`create/list/update/delete/get cross accounts (linked zones, access requests, and permitted networks)`, func() {
				// Create Linked Zone

				createLinkedZoneOptions := service.NewCreateLinkedZoneOptions(instanceID, ownerInstanceID, ownerZoneID)
				createLinkedZoneOptions.SetXCorrelationID("create-linkedZone123")
				resultCreateLZ, responseCreateLZ, errCreateLZ := service.CreateLinkedZone(createLinkedZoneOptions)
				Expect(errCreateLZ).To(BeNil())
				Expect(responseCreateLZ).ToNot(BeNil())
				Expect(responseCreateLZ.GetStatusCode()).To(BeEquivalentTo(200))
				Expect(resultCreateLZ).ToNot(BeNil())

				// List Linked Zones
				listLinkedZonesOptions := service.NewListLinkedZonesOptions(instanceID)
				resultListLZ, responseListLZ, errListLZ := service.ListLinkedZones(listLinkedZonesOptions)
				Expect(errListLZ).To(BeNil())
				Expect(responseListLZ).ToNot(BeNil())
				Expect(responseListLZ.GetStatusCode()).To(BeEquivalentTo(200))
				Expect(resultListLZ).ToNot(BeNil())

				// Get Linked Zone
				getLinkedZonesOptions := service.NewGetLinkedZoneOptions(instanceID, *resultCreateLZ.ID)
				resultGetLZ, responseGetLZ, errGetLZ := service.GetLinkedZone(getLinkedZonesOptions)
				Expect(errGetLZ).To(BeNil())
				Expect(responseGetLZ).ToNot(BeNil())
				Expect(responseGetLZ.GetStatusCode()).To(BeEquivalentTo(200))
				Expect(resultGetLZ).ToNot(BeNil())

				// Update Linked Zone
				updateLinkedZonesOptions := service.NewUpdateLinkedZoneOptions(instanceID, *resultCreateLZ.ID)
				updateLinkedZonesOptions.SetDescription("new description")
				resultUpdateLZ, responseUpdateLZ, errUpdateLZ := service.UpdateLinkedZone(updateLinkedZonesOptions)
				Expect(errUpdateLZ).To(BeNil())
				Expect(responseUpdateLZ).ToNot(BeNil())
				Expect(responseUpdateLZ.GetStatusCode()).To(BeEquivalentTo(200))
				Expect(resultUpdateLZ).ToNot(BeNil())

				// List access request
				listDnszoneAccessRequestsOptions := serviceOwnerDnsInstanceAccount.NewListDnszoneAccessRequestsOptions(
					ownerInstanceID,
					*resultCreateLZ.LinkedTo.ZoneID,
				)
				listDnszoneAccessRequestsOptions.SetXCorrelationID("dnszoneaccessrequest123-list")
				resultListAccessRequests, responseListAccessRequests, errListAccessRequests := serviceOwnerDnsInstanceAccount.ListDnszoneAccessRequests(listDnszoneAccessRequestsOptions)
				Expect(errListAccessRequests).To(BeNil())
				Expect(responseListAccessRequests).ToNot(BeNil())
				Expect(responseListAccessRequests.GetStatusCode()).To(BeEquivalentTo(200))
				Expect(resultListAccessRequests).ToNot(BeNil())
				Expect(resultListAccessRequests.AccessRequests).ToNot(BeNil())
				Expect(len(resultListAccessRequests.AccessRequests) > 0).To(BeTrue())

				// get access request
				getDnszoneAccessRequestsOptions :=
					service.NewGetDnszoneAccessRequestOptions(
						ownerInstanceID,
						ownerZoneID,
						*resultListAccessRequests.AccessRequests[0].ID,
					)
				getDnszoneAccessRequestsOptions.SetXCorrelationID("dnszoneaccessrequest123-get")
				resultGetAccessRequests, responseGetAccessRequests, errGetAccessRequests := serviceOwnerDnsInstanceAccount.GetDnszoneAccessRequest(getDnszoneAccessRequestsOptions)
				Expect(errGetAccessRequests).To(BeNil())
				Expect(responseGetAccessRequests).ToNot(BeNil())
				Expect(responseGetAccessRequests.GetStatusCode()).To(BeEquivalentTo(200))
				Expect(resultGetAccessRequests).ToNot(BeNil())

				// Update access request
				updateDnszoneAccessRequestsOptions :=
					service.NewUpdateDnszoneAccessRequestOptions(
						ownerInstanceID,
						ownerZoneID,
						*resultListAccessRequests.AccessRequests[0].ID,
						"APPROVE",
					)
				updateDnszoneAccessRequestsOptions.SetXCorrelationID("dnszoneaccessrequest123-update")
				resultUpdate, responseUpdate, errUpdate := serviceOwnerDnsInstanceAccount.UpdateDnszoneAccessRequest(updateDnszoneAccessRequestsOptions)
				Expect(errUpdate).To(BeNil())
				Expect(responseUpdate).ToNot(BeNil())
				Expect(responseUpdate.GetStatusCode()).To(BeEquivalentTo(200))
				Expect(resultUpdate).ToNot(BeNil())

				// Create LZ Permitted Networks
				var createLzPermittedNetworkOptionsPermittedNetwork *dnssvcsv1.PermittedNetworkVpc = nil
				createLzPermittedNetworkOptions := service.NewCreateLzPermittedNetworkOptions(instanceID, *resultCreateLZ.ID, "vpc", createLzPermittedNetworkOptionsPermittedNetwork)
				createLzPermittedNetworkOptions.SetXCorrelationID("lzpermittednetworks-create")
				createLzPermittedNetworkOptions.SetType("vpc")

				permittedNetwork, errNewPermittedNetworkVpc := service.NewPermittedNetworkVpc(vpcCrnLzPermittedNetwork)
				Expect(errNewPermittedNetworkVpc).To(BeNil())

				createLzPermittedNetworkOptions.SetPermittedNetwork(permittedNetwork)
				resultCreate, responseCreate, errCreate := service.CreateLzPermittedNetwork(createLzPermittedNetworkOptions)
				Expect(errCreate).To(BeNil())
				Expect(responseCreate).ToNot(BeNil())
				Expect(responseCreate.GetStatusCode()).To(BeEquivalentTo(200))
				Expect(resultCreate).ToNot(BeNil())

				// list LZ Permitted Networks
				listLinkedPermittedNetworkOptions := service.NewListLinkedPermittedNetworksOptions(instanceID, *resultCreateLZ.ID)
				listLinkedPermittedNetworkOptions.SetXCorrelationID("lzpermittednetworks-list")
				resultList, responseList, errList := service.ListLinkedPermittedNetworks(listLinkedPermittedNetworkOptions)
				Expect(errList).To(BeNil())
				Expect(responseList).ToNot(BeNil())
				Expect(responseList.GetStatusCode()).To(BeEquivalentTo(200))
				Expect(resultList).ToNot(BeNil())

				// get LZ Permitted Networks
				getLinkedPermittedNetworkOptions := service.NewGetLinkedPermittedNetworkOptions(instanceID, *resultCreateLZ.ID, *resultCreate.ID)
				getLinkedPermittedNetworkOptions.SetXCorrelationID("lzpermittednetworks-get")
				resultGet, responseGet, errGet := service.GetLinkedPermittedNetwork(getLinkedPermittedNetworkOptions)
				Expect(errGet).To(BeNil())
				Expect(responseGet).ToNot(BeNil())
				Expect(responseGet.GetStatusCode()).To(BeEquivalentTo(200))
				Expect(resultGet).ToNot(BeNil())

				// delete LZ Permitted Networks
				deleteLzPermittedNetworkOptions := service.NewDeleteLzPermittedNetworkOptions(instanceID, *resultCreateLZ.ID, *resultCreate.ID)
				deleteLzPermittedNetworkOptions.SetXCorrelationID("lzpermittednetworks-delete")
				resultDelete, responseDelete, errGet := service.DeleteLzPermittedNetwork(deleteLzPermittedNetworkOptions)
				Expect(errGet).To(BeNil())
				Expect(responseDelete).ToNot(BeNil())
				Expect(responseDelete.GetStatusCode()).To(BeEquivalentTo(202))
				Expect(resultDelete).ToNot(BeNil())

				// Delete Linked Zone
				deleteLinkedZoneOptions := service.NewDeleteLinkedZoneOptions(instanceID, *resultCreateLZ.ID)
				responseDeleteLZ, errDeleteLZ := service.DeleteLinkedZone(deleteLinkedZoneOptions)
				Expect(errDeleteLZ).To(BeNil())
				Expect(responseDeleteLZ).ToNot(BeNil())
				Expect(responseDeleteLZ.GetStatusCode()).To(BeEquivalentTo(204))
			})
		})
	})
})
