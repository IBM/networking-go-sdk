package dnsrecordsv1_test

import (
	"fmt"
	"os"
	"strings"

	"github.com/IBM/go-sdk-core/v5/core"
	. "github.com/IBM/networking-go-sdk/dnsrecordsv1"
	"github.com/joho/godotenv"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const configFile = "../cis.env"

var configLoaded bool = true

func shouldSkipTest() {
	Skip("Authentication failing, skipping...")
	if !configLoaded {
		Skip("External configuration is not available, skipping...")
	}
}

var _ = Describe(`DNSRecordsV1`, func() {
	// BeforeEach(func() {
	// 	Skip("Skipping Tests")
	// })

	if _, err := os.Stat(configFile); err != nil {
		configLoaded = false
	}

	err := godotenv.Load(configFile)
	if err != nil {
		configLoaded = false
		fmt.Println("config is not loaded : ", err)
	}

	authenticator := &core.IamAuthenticator{
		ApiKey: os.Getenv("CIS_SERVICES_APIKEY"),
		URL:    os.Getenv("CIS_SERVICES_AUTH_URL"),
	}
	serviceURL := os.Getenv("API_ENDPOINT")
	crn := os.Getenv("CRN")
	zone_id := os.Getenv("ZONE_ID")
	globalOptions := &DnsRecordsV1Options{
		ServiceName:    "cis_services",
		URL:            serviceURL,
		Authenticator:  authenticator,
		Crn:            &crn,
		ZoneIdentifier: &zone_id,
	}
	testService, testServiceErr := NewDnsRecordsV1(globalOptions)
	if testServiceErr != nil {
		fmt.Println(testServiceErr)
	}
	Describe(`CIS_Frontend_API_Spec-DNSRecords.yaml`, func() {
		Context("DnsRecordsV1Options", func() {
			BeforeEach(func() {
				shouldSkipTest()
				result, response, operationErr := testService.ListAllDnsRecords(testService.NewListAllDnsRecordsOptions())
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
				Expect(*result.Success).Should(BeTrue())

				for _, rule := range result.Result {
					if strings.Contains(*rule.Name, "recordtest") {
						option := testService.NewDeleteDnsRecordOptions(*rule.ID)
						delResult, response, err := testService.DeleteDnsRecord(option)
						Expect(err).To(BeNil())
						Expect(response).ToNot(BeNil())
						Expect(delResult).ToNot(BeNil())
						Expect(*result.Success).Should(BeTrue())
					}
				}
			})
			AfterEach(func() {
				shouldSkipTest()
				result, response, operationErr := testService.ListAllDnsRecords(testService.NewListAllDnsRecordsOptions())
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
				Expect(*result.Success).Should(BeTrue())

				for _, rule := range result.Result {
					if strings.Contains(*rule.Name, "recordtest") {
						option := testService.NewDeleteDnsRecordOptions(*rule.ID)
						delResult, response, err := testService.DeleteDnsRecord(option)
						Expect(err).To(BeNil())
						Expect(response).ToNot(BeNil())
						Expect(delResult).ToNot(BeNil())
						Expect(*result.Success).Should(BeTrue())
					}
				}
			})
			It(`create/delete/get dns A type records`, func() {
				shouldSkipTest()
				mode := CreateDnsRecordOptions_Type_A
				options := testService.NewCreateDnsRecordOptions()
				options.SetName("host-1.recordtest.com")
				options.SetType(mode)
				options.SetContent("1.2.3.4")
				options.SetTTL(900)
				result, response, err := testService.CreateDnsRecord(options)
				Expect(err).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// get DNS Record Options
				getOption := testService.NewGetDnsRecordOptions(*result.Result.ID)
				result, response, err = testService.GetDnsRecord(getOption)
				Expect(err).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
				Expect(*result.Success).Should(BeTrue())

				// update DNS Record Options
				updateOpt := testService.NewUpdateDnsRecordOptions(*result.Result.ID)
				newModes := CreateDnsRecordOptions_Type_Txt
				updateOpt.SetType(newModes)
				updateOpt.SetName("host-1.recordtest.com")
				options.SetTTL(100)
				updateOpt.SetContent("Test Text")
				result, response, err = testService.UpdateDnsRecord(updateOpt)
				Expect(err).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
				Expect(*result.Success).Should(BeTrue())

				// delete Dns Record Options
				deleteOpt := testService.NewDeleteDnsRecordOptions(*result.Result.ID)
				delResult, response, err := testService.DeleteDnsRecord(deleteOpt)
				Expect(err).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(delResult).ToNot(BeNil())
				Expect(*result.Success).Should(BeTrue())
			})
			It(`create/delete/get dns Caa type records`, func() {
				shouldSkipTest()
				mode := CreateDnsRecordOptions_Type_Caa
				options := testService.NewCreateDnsRecordOptions()
				options.SetName("host-1.recordtest.com")
				options.SetType(mode)
				options.SetTTL(900)
				Data := map[string]interface{}{"tag": "http",
					"value": "domain.com"}
				options.SetData(Data)
				result, response, err := testService.CreateDnsRecord(options)
				Expect(err).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// get DNS Reord Options
				getOption := testService.NewGetDnsRecordOptions(*result.Result.ID)
				result, response, err = testService.GetDnsRecord(getOption)
				Expect(err).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
				Expect(*result.Success).Should(BeTrue())

				//Update DNS Record Option
				updateOpt := testService.NewUpdateDnsRecordOptions(*result.Result.ID)
				newModes := CreateDnsRecordOptions_Type_Txt
				updateOpt.SetType(newModes)
				updateOpt.SetName("host-1.recordtest.com")
				options.SetTTL(100)
				updateOpt.SetContent("Test Text")
				result, response, err = testService.UpdateDnsRecord(updateOpt)
				Expect(err).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
				Expect(*result.Success).Should(BeTrue())

				// delete Dns Record Options
				deleteOpt := testService.NewDeleteDnsRecordOptions(*result.Result.ID)
				delResult, response, err := testService.DeleteDnsRecord(deleteOpt)
				Expect(err).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(delResult).ToNot(BeNil())
				Expect(*result.Success).Should(BeTrue())
			})
			It(`create/delete/get dns Cname type records`, func() {
				shouldSkipTest()
				mode := CreateDnsRecordOptions_Type_Cname
				options := testService.NewCreateDnsRecordOptions()
				options.SetName("host-1.recordtest.com")
				options.SetType(mode)
				options.SetTTL(900)
				options.SetContent("domain.com")
				result, response, err := testService.CreateDnsRecord(options)
				Expect(err).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// get DNS Reord Options
				getOption := testService.NewGetDnsRecordOptions(*result.Result.ID)
				result, response, err = testService.GetDnsRecord(getOption)
				Expect(err).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
				Expect(*result.Success).Should(BeTrue())

				// delete Dns Record Options
				updateOpt := testService.NewUpdateDnsRecordOptions(*result.Result.ID)
				newModes := CreateDnsRecordOptions_Type_Txt
				updateOpt.SetType(newModes)
				updateOpt.SetName("host-1.recordtest.com")
				options.SetTTL(100)
				updateOpt.SetContent("Test Text")
				result, response, err = testService.UpdateDnsRecord(updateOpt)
				Expect(err).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
				Expect(*result.Success).Should(BeTrue())

				// delete Dns Record Options
				deleteOpt := testService.NewDeleteDnsRecordOptions(*result.Result.ID)
				delResult, response, err := testService.DeleteDnsRecord(deleteOpt)
				Expect(err).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(delResult).ToNot(BeNil())
				Expect(*result.Success).Should(BeTrue())
			})
			It(`create/delete/get dns Aaaa type records`, func() {
				shouldSkipTest()
				mode := CreateDnsRecordOptions_Type_Aaaa
				options := testService.NewCreateDnsRecordOptions()
				options.SetName("host-1.recordtest.com")
				options.SetType(mode)
				options.SetContent("2001:0db8:85a3:0000:0000:8a2e:0370:7334")
				options.SetTTL(900)
				result, response, err := testService.CreateDnsRecord(options)
				Expect(err).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// get DNS Record Options
				getOption := testService.NewGetDnsRecordOptions(*result.Result.ID)
				result, response, err = testService.GetDnsRecord(getOption)
				Expect(err).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
				Expect(*result.Success).Should(BeTrue())

				// delete Dns Record Options
				updateOpt := testService.NewUpdateDnsRecordOptions(*result.Result.ID)
				newModes := CreateDnsRecordOptions_Type_Txt
				updateOpt.SetType(newModes)
				updateOpt.SetName("host-1.recordtest.com")
				updateOpt.SetContent("Test Text")
				options.SetTTL(100)
				result, response, err = testService.UpdateDnsRecord(updateOpt)
				Expect(err).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
				Expect(*result.Success).Should(BeTrue())

				// delete Dns Record Options
				deleteOpt := testService.NewDeleteDnsRecordOptions(*result.Result.ID)
				delResult, response, err := testService.DeleteDnsRecord(deleteOpt)
				Expect(err).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(delResult).ToNot(BeNil())
				Expect(*result.Success).Should(BeTrue())
			})
			It(`create/delete/get dns Mx type records`, func() {
				shouldSkipTest()
				mode := CreateDnsRecordOptions_Type_Mx
				options := testService.NewCreateDnsRecordOptions()
				options.SetName("host-1.recordtest.com")
				options.SetType(mode)
				//options.SetContent("2001:0db8:85a3:0000:0000:8a2e:0370:7334")
				options.SetTTL(900)
				options.SetContent("domain.com")
				options.SetPriority(int64(1))
				result, response, err := testService.CreateDnsRecord(options)
				Expect(err).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// get DNS Reord Options
				getOption := testService.NewGetDnsRecordOptions(*result.Result.ID)
				result, response, err = testService.GetDnsRecord(getOption)
				Expect(err).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
				Expect(*result.Success).Should(BeTrue())

				updateOpt := testService.NewUpdateDnsRecordOptions(*result.Result.ID)
				newModes := CreateDnsRecordOptions_Type_Txt
				updateOpt.SetType(newModes)
				updateOpt.SetName("host-1.recordtest.com")
				options.SetTTL(100)
				updateOpt.SetContent("Test Text")
				result, response, err = testService.UpdateDnsRecord(updateOpt)
				Expect(err).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
				Expect(*result.Success).Should(BeTrue())

				// delete Dns Record Options
				deleteOpt := testService.NewDeleteDnsRecordOptions(*result.Result.ID)
				delResult, response, err := testService.DeleteDnsRecord(deleteOpt)
				Expect(err).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(delResult).ToNot(BeNil())
				Expect(*result.Success).Should(BeTrue())
			})
			It(`create/delete/get dns Ns type records`, func() {
				shouldSkipTest()
				mode := CreateDnsRecordOptions_Type_Ns
				options := testService.NewCreateDnsRecordOptions()
				options.SetName("host-1.recordtest.com")
				options.SetTTL(900)
				options.SetType(mode)
				options.SetContent("domain.com")
				result, response, err := testService.CreateDnsRecord(options)
				Expect(err).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// get DNS Reord Options
				getOption := testService.NewGetDnsRecordOptions(*result.Result.ID)
				result, response, err = testService.GetDnsRecord(getOption)
				Expect(err).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
				Expect(*result.Success).Should(BeTrue())

				updateOpt := testService.NewUpdateDnsRecordOptions(*result.Result.ID)
				newModes := CreateDnsRecordOptions_Type_Txt
				updateOpt.SetType(newModes)
				updateOpt.SetName("host-1.recordtest.com")
				options.SetTTL(100)
				updateOpt.SetContent("Test Text")
				result, response, err = testService.UpdateDnsRecord(updateOpt)
				Expect(err).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
				Expect(*result.Success).Should(BeTrue())

				// delete Dns Record Options
				deleteOpt := testService.NewDeleteDnsRecordOptions(*result.Result.ID)
				delResult, response, err := testService.DeleteDnsRecord(deleteOpt)
				Expect(err).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(delResult).ToNot(BeNil())
				Expect(*result.Success).Should(BeTrue())
			})
			It(`create/delete/get dns Srv type records`, func() {
				shouldSkipTest()
				mode := CreateDnsRecordOptions_Type_Srv
				options := testService.NewCreateDnsRecordOptions()
				options.SetType(mode)
				options.SetTTL(100)
				Data := map[string]interface{}{
					"name":     "recordtest",
					"priority": int64(1),
					"service":  "_sip.recordtest",
					"proto":    "_udp",
					"weight":   1,
					"port":     1,
					"target":   "domain.com"}
				options.SetData(Data)
				result, response, err := testService.CreateDnsRecord(options)
				Expect(err).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// get DNS Reord Options
				getOption := testService.NewGetDnsRecordOptions(*result.Result.ID)
				result, response, err = testService.GetDnsRecord(getOption)
				Expect(err).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
				Expect(*result.Success).Should(BeTrue())

				updateOpt := testService.NewUpdateDnsRecordOptions(*result.Result.ID)
				newModes := CreateDnsRecordOptions_Type_Txt
				updateOpt.SetType(newModes)
				updateOpt.SetName("host-1.recordtest.com")
				updateOpt.SetContent("Test Text")
				result, response, err = testService.UpdateDnsRecord(updateOpt)
				Expect(err).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
				Expect(*result.Success).Should(BeTrue())

				// delete Dns Record Options
				deleteOpt := testService.NewDeleteDnsRecordOptions(*result.Result.ID)
				delResult, response, err := testService.DeleteDnsRecord(deleteOpt)
				Expect(err).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(delResult).ToNot(BeNil())
				Expect(*result.Success).Should(BeTrue())
			})
			It(`create/delete/get dns Txt type records`, func() {
				shouldSkipTest()
				mode := CreateDnsRecordOptions_Type_Txt
				options := testService.NewCreateDnsRecordOptions()
				options.SetName("host-1.recordtest.com")
				options.SetType(mode)
				options.SetTTL(900)
				options.SetContent("Test Text")

				result, response, err := testService.CreateDnsRecord(options)
				Expect(err).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// get DNS Reord Options
				getOption := testService.NewGetDnsRecordOptions(*result.Result.ID)
				result, response, err = testService.GetDnsRecord(getOption)
				Expect(err).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
				Expect(*result.Success).Should(BeTrue())

				updateOpt := testService.NewUpdateDnsRecordOptions(*result.Result.ID)
				newModes := CreateDnsRecordOptions_Type_A
				updateOpt.SetType(newModes)
				updateOpt.SetName("host-1.recordtest.com")
				options.SetTTL(120)
				updateOpt.SetContent("1.2.3.4")
				result, response, err = testService.UpdateDnsRecord(updateOpt)
				Expect(err).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
				Expect(*result.Success).Should(BeTrue())

				// delete Dns Record Options
				deleteOpt := testService.NewDeleteDnsRecordOptions(*result.Result.ID)
				delResult, response, err := testService.DeleteDnsRecord(deleteOpt)
				Expect(err).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(delResult).ToNot(BeNil())
				Expect(*result.Success).Should(BeTrue())
			})
			It(`batch dns records - posts, puts, patches, and deletes`, func() {
				shouldSkipTest()

				// Pre-create two records to use for puts/patches/deletes in the batch
				createOptA := testService.NewCreateDnsRecordOptions()
				createOptA.SetName("batch-put.recordtest.com")
				createOptA.SetType(CreateDnsRecordOptions_Type_A)
				createOptA.SetContent("1.2.3.4")
				createOptA.SetTTL(900)
				putRecord, response, err := testService.CreateDnsRecord(createOptA)
				Expect(err).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(putRecord).ToNot(BeNil())

				createOptB := testService.NewCreateDnsRecordOptions()
				createOptB.SetName("batch-patch.recordtest.com")
				createOptB.SetType(CreateDnsRecordOptions_Type_A)
				createOptB.SetContent("5.6.7.8")
				createOptB.SetTTL(900)
				patchRecord, response, err := testService.CreateDnsRecord(createOptB)
				Expect(err).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(patchRecord).ToNot(BeNil())

				createOptC := testService.NewCreateDnsRecordOptions()
				createOptC.SetName("batch-delete.recordtest.com")
				createOptC.SetType(CreateDnsRecordOptions_Type_A)
				createOptC.SetContent("9.10.11.12")
				createOptC.SetTTL(900)
				deleteRecord, response, err := testService.CreateDnsRecord(createOptC)
				Expect(err).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(deleteRecord).ToNot(BeNil())

				// Build batch request
				postInput, postErr := testService.NewDnsrecordInput(DnsrecordInput_Type_Txt)
				Expect(postErr).To(BeNil())
				postInput.Name = core.StringPtr("batch-post.recordtest.com")
				postInput.Content = core.StringPtr("batch post content")
				postInput.TTL = core.Int64Ptr(int64(900))

				putItem, putErr := testService.NewBatchDnsRecordsRequestPutsItem(
					*putRecord.Result.ID,
					"batch-put.recordtest.com",
					BatchDnsRecordsRequestPutsItem_Type_A,
					int64(300),
					"1.2.3.5",
				)
				Expect(putErr).To(BeNil())

				patchItem, patchErr := testService.NewBatchDnsRecordsRequestPatchesItem(*patchRecord.Result.ID)
				Expect(patchErr).To(BeNil())
				patchItem.Content = core.StringPtr("5.6.7.9")

				deleteItem, deleteErr := testService.NewBatchDnsRecordsRequestDeletesItem(*deleteRecord.Result.ID)
				Expect(deleteErr).To(BeNil())

				batchOptions := testService.NewBatchDnsRecordsOptions()
				batchOptions.SetPosts([]DnsrecordInput{*postInput})
				batchOptions.SetPuts([]BatchDnsRecordsRequestPutsItem{*putItem})
				batchOptions.SetPatches([]BatchDnsRecordsRequestPatchesItem{*patchItem})
				batchOptions.SetDeletes([]BatchDnsRecordsRequestDeletesItem{*deleteItem})

				batchResult, response, err := testService.BatchDnsRecords(batchOptions)
				Expect(err).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(batchResult).ToNot(BeNil())
				Expect(*batchResult.Success).Should(BeTrue())
				Expect(batchResult.Result).ToNot(BeNil())
				Expect(len(batchResult.Result.Posts)).To(Equal(1))
				Expect(len(batchResult.Result.Puts)).To(Equal(1))
				Expect(len(batchResult.Result.Patches)).To(Equal(1))
				Expect(len(batchResult.Result.Deletes)).To(Equal(1))

				// Clean up records created by the batch post and put
				for _, record := range batchResult.Result.Posts {
					if record.ID != nil {
						cleanOpt := testService.NewDeleteDnsRecordOptions(*record.ID)
						_, _, err = testService.DeleteDnsRecord(cleanOpt)
						Expect(err).To(BeNil())
					}
				}
				for _, record := range batchResult.Result.Puts {
					if record.ID != nil {
						cleanOpt := testService.NewDeleteDnsRecordOptions(*record.ID)
						_, _, err = testService.DeleteDnsRecord(cleanOpt)
						Expect(err).To(BeNil())
					}
				}
				for _, record := range batchResult.Result.Patches {
					if record.ID != nil {
						cleanOpt := testService.NewDeleteDnsRecordOptions(*record.ID)
						_, _, err = testService.DeleteDnsRecord(cleanOpt)
						Expect(err).To(BeNil())
					}
				}
			})
		})
	})
})
