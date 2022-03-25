package edgefunctionsapiv1_test

import (
	"fmt"
	"os"

	"github.com/IBM/go-sdk-core/v5/core"
	. "github.com/IBM/networking-go-sdk/edgefunctionsapiv1"
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

var _ = Describe(`edgefunctionsapiv1_test`, func() {
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
	url := os.Getenv("URL")

	globalOpt := &EdgeFunctionsApiV1Options{
		ServiceName:    "cis_services",
		URL:            serviceURL,
		Authenticator:  authenticator,
		Crn:            &crn,
		ZoneIdentifier: &zone_id,
	}

	testService, testServiceErr := NewEdgeFunctionsApiV1(globalOpt)
	if testServiceErr != nil {
		fmt.Println(testServiceErr)
	}
	Describe(`edgefunctionsapiv1_test`, func() {
		Context(`edgefunctionsapiv1_test`, func() {
			BeforeEach(func() {
				shouldSkipTest()
				// list all scripts test
				listOpt := testService.NewListEdgeFunctionsActionsOptions()
				listResult, listResp, listErr := testService.ListEdgeFunctionsActions(listOpt)
				Expect(listErr).To(BeNil())
				Expect(listResp).ToNot(BeNil())
				Expect(listResult).ToNot(BeNil())
				Expect(*listResult.Success).To(BeTrue())

				// delete all trigger test
				for _, function := range listResult.Result {
					var scriptName string
					for _, trigger := range function.Routes {
						scriptName = *trigger.Script
						delOpt := testService.NewDeleteEdgeFunctionsTriggerOptions(*trigger.ID)
						delResult, delResp, delErr := testService.DeleteEdgeFunctionsTrigger(delOpt)
						Expect(delErr).To(BeNil())
						Expect(delResp).ToNot(BeNil())
						Expect(delResult).ToNot(BeNil())
						Expect(*delResult.Success).To(BeTrue())
					}
					// delete script test
					if len(scriptName) != 0 {
						delOpt := testService.NewDeleteEdgeFunctionsActionOptions(scriptName)
						delResult, delResp, delErr := testService.DeleteEdgeFunctionsAction(delOpt)
						Expect(delErr).To(BeNil())
						Expect(delResp).ToNot(BeNil())
						Expect(delResult).ToNot(BeNil())
						Expect(*delResult.Success).To(BeTrue())
					}
				}
				delOpt := testService.NewDeleteEdgeFunctionsActionOptions("test-script")
				_, _, _ = testService.DeleteEdgeFunctionsAction(delOpt)
			})
			AfterEach(func() {
				shouldSkipTest()
				// list all scripts test
				listOpt := testService.NewListEdgeFunctionsActionsOptions()
				listResult, listResp, listErr := testService.ListEdgeFunctionsActions(listOpt)
				Expect(listErr).To(BeNil())
				Expect(listResp).ToNot(BeNil())
				Expect(listResult).ToNot(BeNil())
				Expect(*listResult.Success).To(BeTrue())

				// delete all trigger test
				for _, function := range listResult.Result {
					var scriptName string
					for _, trigger := range function.Routes {
						scriptName = *trigger.Script
						delOpt := testService.NewDeleteEdgeFunctionsTriggerOptions(*trigger.ID)
						delResult, delResp, delErr := testService.DeleteEdgeFunctionsTrigger(delOpt)
						Expect(delErr).To(BeNil())
						Expect(delResp).ToNot(BeNil())
						Expect(delResult).ToNot(BeNil())
						Expect(*delResult.Success).To(BeTrue())
					}
					// delete script test
					if len(scriptName) != 0 {
						delOpt := testService.NewDeleteEdgeFunctionsActionOptions(scriptName)
						delResult, delResp, delErr := testService.DeleteEdgeFunctionsAction(delOpt)
						Expect(delErr).To(BeNil())
						Expect(delResp).ToNot(BeNil())
						Expect(delResult).ToNot(BeNil())
						Expect(*delResult.Success).To(BeTrue())
					}
				}
				delOpt := testService.NewDeleteEdgeFunctionsActionOptions("test-script")
				_, _, _ = testService.DeleteEdgeFunctionsAction(delOpt)
			})
			It(`upload/delete/get Edge Functions actions and triggers`, func() {
				shouldSkipTest()

				// Script upload test
				reader, err := os.Open("./script.js")
				Expect(err).To(BeNil())
				uploadOpt := testService.NewUpdateEdgeFunctionsActionOptions("test-script")
				uploadOpt.SetEdgeFunctionsAction(reader)
				uploadResult, uploadResp, uploadErr := testService.UpdateEdgeFunctionsAction(uploadOpt)
				Expect(uploadErr).To(BeNil())
				Expect(uploadResp).ToNot(BeNil())
				Expect(uploadResult).ToNot(BeNil())
				Expect(*uploadResult.Success).To(BeTrue())

				reader, err = os.Open("./script.js")
				Expect(err).To(BeNil())
				uploadOpt = testService.NewUpdateEdgeFunctionsActionOptions("test-script1")
				uploadOpt.SetEdgeFunctionsAction(reader)
				uploadResult, uploadResp, uploadErr = testService.UpdateEdgeFunctionsAction(uploadOpt)
				Expect(uploadErr).To(BeNil())
				Expect(uploadResp).ToNot(BeNil())
				Expect(uploadResult).ToNot(BeNil())
				Expect(*uploadResult.Success).To(BeTrue())

				// Get the script test
				getOpt := testService.NewGetEdgeFunctionsActionOptions("test-script1")
				getResult, getResp, getErr := testService.GetEdgeFunctionsAction(getOpt)
				Expect(getErr).To(BeNil())
				Expect(getResp).ToNot(BeNil())
				Expect(getResult).ToNot(BeNil())

				p := make([]byte, 256)
				f, err := os.Create("/tmp/script.js")
				Expect(err).To(BeNil())
				_, err = getResult.Read(p)
				Expect(err).To(BeNil())
				for len(p) > 0 {
					_, err := f.Write(p)
					if err != nil {
						break
					}
					_, err = getResult.Read(p)
					if err != nil {
						break
					}
				}
				f.Close()
				os.Remove("/tmp/script.js")

				// Trigger creation test
				pattern := fmt.Sprintf("%s.%s/*", "example", url)
				createOpt := testService.NewCreateEdgeFunctionsTriggerOptions()
				createOpt.SetScript("test-script")
				createOpt.SetPattern(pattern)
				createResult, createResp, createErr := testService.CreateEdgeFunctionsTrigger(createOpt)
				Expect(createErr).To(BeNil())
				Expect(createResp).ToNot(BeNil())
				Expect(createResult).ToNot(BeNil())
				Expect(*createResult.Success).To(BeTrue())

				// Trigger update test
				pattern = fmt.Sprintf("%s.%s/*", "test-example", url)
				updateOpt := testService.NewUpdateEdgeFunctionsTriggerOptions(*createResult.Result.ID)
				updateOpt.SetPattern(pattern)
				updateOpt.SetScript("test-script1")
				updateResult, updateResp, updateErr := testService.UpdateEdgeFunctionsTrigger(updateOpt)
				Expect(updateErr).To(BeNil())
				Expect(updateResp).ToNot(BeNil())
				Expect(updateResult).ToNot(BeNil())
				Expect(*updateResult.Success).To(BeTrue())

				// Trigger get test
				triggerGetOpt := testService.NewGetEdgeFunctionsTriggerOptions(*updateResult.Result.ID)
				triggerGetResult, triggerGetResp, triggerGetErr := testService.GetEdgeFunctionsTrigger(triggerGetOpt)
				Expect(triggerGetErr).To(BeNil())
				Expect(triggerGetResp).ToNot(BeNil())
				Expect(triggerGetResult).ToNot(BeNil())
				Expect(*triggerGetResult.Success).To(BeTrue())

				// List all triggers test
				listTriggerOpt := testService.NewListEdgeFunctionsTriggersOptions()
				listTriggerResult, listTriggerResp, listTriggerErr := testService.ListEdgeFunctionsTriggers(listTriggerOpt)
				Expect(listTriggerErr).To(BeNil())
				Expect(listTriggerResp).ToNot(BeNil())
				Expect(listTriggerResult).ToNot(BeNil())
				Expect(*listTriggerResult.Success).To(BeTrue())

				// list all scripts test
				listOpt := testService.NewListEdgeFunctionsActionsOptions()
				listResult, listResp, listErr := testService.ListEdgeFunctionsActions(listOpt)
				Expect(listErr).To(BeNil())
				Expect(listResp).ToNot(BeNil())
				Expect(listResult).ToNot(BeNil())
				Expect(*listResult.Success).To(BeTrue())

				// delete all trigger test
				for _, function := range listResult.Result {
					var scriptName string
					for _, trigger := range function.Routes {
						scriptName = *trigger.Script
						delOpt := testService.NewDeleteEdgeFunctionsTriggerOptions(*trigger.ID)
						delResult, delResp, delErr := testService.DeleteEdgeFunctionsTrigger(delOpt)
						Expect(delErr).To(BeNil())
						Expect(delResp).ToNot(BeNil())
						Expect(delResult).ToNot(BeNil())
						Expect(*delResult.Success).To(BeTrue())
					}
					// delete script test
					if len(scriptName) != 0 {
						// delete script test
						delOpt := testService.NewDeleteEdgeFunctionsActionOptions(scriptName)
						delResult, delResp, delErr := testService.DeleteEdgeFunctionsAction(delOpt)
						Expect(delErr).To(BeNil())
						Expect(delResp).ToNot(BeNil())
						Expect(delResult).ToNot(BeNil())
						Expect(*delResult.Success).To(BeTrue())
					}
				}
			})
		})
	})
})
