package dnsrecordbulkv1_test

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/IBM/go-sdk-core/core"
	"github.com/joho/godotenv"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/IBM/networking-go-sdk/dnsrecordbulkv1"
	"github.com/IBM/networking-go-sdk/dnsrecordsv1"
)

const configFile = "../cis.env"

var configLoaded bool = true

func shouldSkipTest() {
	if !configLoaded {
		Skip("External configuration is not available, skipping...")
	}
}

var _ = Describe(`DnsRecordBulkV1`, func() {
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

	globalOptions := &dnsrecordbulkv1.DnsRecordBulkV1Options{
		ServiceName:    "cis_services",
		URL:            serviceURL,
		Authenticator:  authenticator,
		Crn:            &crn,
		ZoneIdentifier: &zone_id,
	}
	globalParameters := &dnsrecordsv1.DnsRecordsV1Options{
		ServiceName:    "cis_services",
		URL:            serviceURL,
		Authenticator:  authenticator,
		Crn:            &crn,
		ZoneIdentifier: &zone_id,
	}
	testService, testServiceErr := dnsrecordbulkv1.NewDnsRecordBulkV1(globalOptions)
	dnsTestService, dnsTestServiceErr := dnsrecordsv1.NewDnsRecordsV1(globalParameters)
	if testServiceErr != nil {
		fmt.Println(testServiceErr)
	}
	if dnsTestServiceErr != nil {
		fmt.Println(dnsTestServiceErr)
	}

	Describe(`CIS_Frontend_API_Spec-DNS_Records_Bulk.yaml`, func() {
		Context("DnsRecordsbulkV1Options_get", func() {
			BeforeEach(func() {
				listResult, listResponse, listErr := dnsTestService.ListAllDnsRecords(dnsTestService.NewListAllDnsRecordsOptions())
				Expect(listErr).To(BeNil())
				Expect(listResponse).ToNot(BeNil())
				Expect(listResult).ToNot(BeNil())
				Expect(*listResult.Success).Should(BeTrue())

				for _, rule := range listResult.Result {
					if strings.Contains(*rule.Name, "example") {
						option := dnsTestService.NewDeleteDnsRecordOptions(*rule.ID)
						delResult, response, err := dnsTestService.DeleteDnsRecord(option)
						Expect(err).To(BeNil())
						Expect(response).ToNot(BeNil())
						Expect(delResult).ToNot(BeNil())
						Expect(*delResult.Success).Should(BeTrue())
					}
				}
				dnsOptions := dnsTestService.NewCreateDnsRecordOptions()
				dnsOptions.SetName("host-9.test-example.com")
				dnsOptions.SetType(dnsrecordsv1.CreateDnsRecordOptions_Type_A)
				dnsOptions.SetContent("12.12.12.1")
				result, response, err := dnsTestService.CreateDnsRecord(dnsOptions)
				Expect(err).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
				Expect(*result.Success).Should(BeTrue())
				dnsOptions.SetName("host-15.test-example.com")
				dnsOptions.SetType(dnsrecordsv1.CreateDnsRecordOptions_Type_Aaaa)
				dnsOptions.SetContent("2001:0db8:85a3:0000:0000:8a2e:0370:7334")
				result, response, err = dnsTestService.CreateDnsRecord(dnsOptions)
				Expect(err).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
				dnsOptions = dnsTestService.NewCreateDnsRecordOptions()
				dnsOptions.SetName("host-16.test-example.com")
				dnsOptions.SetType(dnsrecordsv1.CreateDnsRecordOptions_Type_Caa)
				Data := map[string]interface{}{"tag": "http",
					"value": "domain1.com"}
				dnsOptions.SetData(Data)
				result, response, err = dnsTestService.CreateDnsRecord(dnsOptions)
				Expect(err).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
				dnsOptions = dnsTestService.NewCreateDnsRecordOptions()
				dnsOptions.SetName("host-22.test-example.com")
				dnsOptions.SetType(dnsrecordsv1.CreateDnsRecordOptions_Type_Cname)
				dnsOptions.SetContent("domain2.com")
				result, response, err = dnsTestService.CreateDnsRecord(dnsOptions)
				Expect(err).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
				dnsOptions = dnsTestService.NewCreateDnsRecordOptions()
				dnsOptions.SetName("host-33.test-example.com")
				dnsOptions.SetType(dnsrecordsv1.CreateDnsRecordOptions_Type_Mx)
				dnsOptions.SetPriority(5)
				dnsOptions.SetContent("example-domain.com")
				result, response, err = dnsTestService.CreateDnsRecord(dnsOptions)
				Expect(err).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
				dnsOptions = dnsTestService.NewCreateDnsRecordOptions()
				dnsOptions.SetName("host-44.test-example.com")
				dnsOptions.SetType(dnsrecordsv1.CreateDnsRecordOptions_Type_Ns)
				dnsOptions.SetContent("domain6.com")
				result, response, err = dnsTestService.CreateDnsRecord(dnsOptions)
				Expect(err).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
				dnsOptions = dnsTestService.NewCreateDnsRecordOptions()
				dnsOptions.SetName("host-55.test-example.com")
				dnsOptions.SetType(dnsrecordsv1.CreateDnsRecordOptions_Type_Spf)
				dnsOptions.SetContent("domain7.com")
				result, response, err = dnsTestService.CreateDnsRecord(dnsOptions)
				Expect(err).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
				dnsOptions = dnsTestService.NewCreateDnsRecordOptions()
				dnsOptions.SetName("host-10.test-example.com")
				dnsOptions.SetType(dnsrecordsv1.CreateDnsRecordOptions_Type_Txt)
				dnsOptions.SetContent("Test Text")
				result, response, err = dnsTestService.CreateDnsRecord(dnsOptions)
				Expect(err).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			AfterEach(func() {
				result, response, operationErr := dnsTestService.ListAllDnsRecords(dnsTestService.NewListAllDnsRecordsOptions())
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
				Expect(*result.Success).Should(BeTrue())

				for _, rule := range result.Result {
					if strings.Contains(*rule.Name, "example") {
						option := dnsTestService.NewDeleteDnsRecordOptions(*rule.ID)
						delResult, response, err := dnsTestService.DeleteDnsRecord(option)
						Expect(err).To(BeNil())
						Expect(response).ToNot(BeNil())
						Expect(delResult).ToNot(BeNil())
						Expect(*result.Success).Should(BeTrue())
					}
				}
			})
			It(`get DnsBulk Records`, func() {
				shouldSkipTest()
				testService, testServiceErr = dnsrecordbulkv1.NewDnsRecordBulkV1(globalOptions)
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())
				Expect(err).To(BeNil())

				options := testService.NewGetDnsRecordsBulkOptions()
				result, response, err := testService.GetDnsRecordsBulk(options)
				Expect(err).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
				Expect(result.Read).ToNot(BeNil())
				p := make([]byte, 256)
				f, err := os.Create("/tmp/records.txt")
				Expect(err).To(BeNil())
				_, err = result.Read(p)
				Expect(err).To(BeNil())
				for len(p) > 0 {
					_, err := f.Write(p)
					if err != nil {
						break
					}
					_, err = result.Read(p)
					if err != nil {
						break
					}
				}
				postOptions := testService.NewPostDnsRecordsBulkOptions()
				var reader io.ReadCloser
				reader, err = os.Open("./records.txt")
				Expect(err).To(BeNil())
				postOptions.SetFile(reader)
				postResult, postResponse, postErr := testService.PostDnsRecordsBulk(postOptions)
				Expect(postErr).To(BeNil())
				Expect(postResponse).ToNot(BeNil())
				Expect(postResult).ToNot(BeNil())
				err = reader.Close()
				Expect(err).To(BeNil())
				err = f.Close()
				Expect(err).To(BeNil())
				err = os.Remove("/tmp/records.txt")
				Expect(err).To(BeNil())
				err = result.Close()
				Expect(err).To(BeNil())

			})
		})
	})
})
