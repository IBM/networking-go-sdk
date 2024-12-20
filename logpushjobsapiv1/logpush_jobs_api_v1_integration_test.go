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

package logpushjobsapiv1_test

import (
	"fmt"
	"os"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/networking-go-sdk/logpushjobsapiv1"
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

var _ = Describe(`LogpushJobsApiV1`, func() {
	defer GinkgoRecover()
	if _, err := os.Stat(configFile); err != nil {
		configLoaded = false
	}

	err := godotenv.Load(configFile)
	if err != nil {
		configLoaded = false
	}

	authenticator := &core.IamAuthenticator{
		ApiKey: os.Getenv("CIS_SERVICES_APIKEY"),
		URL:    os.Getenv("CIS_SERVICES_AUTH_URL"),
	}

	crn := os.Getenv("CRN")
	zoneId := os.Getenv("ZONE_ID")
	serviceURL := os.Getenv("API_ENDPOINT")
	IngressKey := os.Getenv("INGRESS_KEY")
	LogdnaRegion := os.Getenv("LOGDNA_REGION")
	LogdnaDomain := os.Getenv("DOMAIN_NAME")
	CosBucket := os.Getenv("COS_BUCKET")
	CosRegion := os.Getenv("COS_REGION")
	CosInstance := os.Getenv("COS_INSTANCE")
	OwnershipToken := os.Getenv("OWNERSHIP_TOKEN")

	dataset := "http_requests"
	globalOptions := &logpushjobsapiv1.LogpushJobsApiV1Options{
		ServiceName:   logpushjobsapiv1.DefaultServiceName,
		Crn:           &crn,
		ZoneID:        &zoneId,
		URL:           serviceURL,
		Dataset:       &dataset,
		Authenticator: authenticator,
	}

	testService, testServiceErr := logpushjobsapiv1.NewLogpushJobsApiV1(globalOptions)
	Expect(testServiceErr).To(BeNil())

	Describe(`LogpushJobsApiV1_test`, func() {
		Context(`LogpushJobsApiV1 All Jobs`, func() {

			BeforeEach(func() {
				shouldSkipTest()
				//List Logpush Jobs
				listOptions := testService.NewGetLogpushJobsV2Options()
				result, response, operationErr := testService.GetLogpushJobsV2(listOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				//Delete Logpush Jobs
				for _, job := range result.Result {
					delOptions := testService.NewDeleteLogpushJobV2Options(*job.ID)
					result, response, deleteErr := testService.DeleteLogpushJobV2(delOptions)
					Expect(deleteErr).To(BeNil())
					Expect(response).ToNot(BeNil())
					Expect(result).ToNot(BeNil())
				}
			})
			AfterEach(func() {
				shouldSkipTest()
				// List all Logpush Jobs
				listOptions := testService.NewGetLogpushJobsV2Options()
				result, response, operationErr := testService.GetLogpushJobsV2(listOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				//Delete all Logpush Jobs
				for _, job := range result.Result {
					delOptions := testService.NewDeleteLogpushJobV2Options(*job.ID)
					result, response, deleteErr := testService.DeleteLogpushJobV2(delOptions)
					Expect(deleteErr).To(BeNil())
					Expect(response).ToNot(BeNil())
					Expect(result).ToNot(BeNil())
				}
			})

			It(`create/update/delete/get logpush jobs`, func() {
				shouldSkipTest()

				options := testService.NewCreateLogpushJobV2Options()
				createLogpushJobV2RequestModel := &logpushjobsapiv1.CreateLogpushJobV2RequestLogpushJobLogdnaReq{

					Name:           core.StringPtr("Test123"),
					Enabled:        core.BoolPtr(false),
					LogpullOptions: core.StringPtr("timestamps=rfc3339&timestamps=rfc3339"),
					Logdna:         map[string]interface{}{"ingress_key": IngressKey, "region": LogdnaRegion, "hostname": LogdnaDomain},
					Dataset:        core.StringPtr("http_requests"),
					Frequency:      core.StringPtr("high"),
				}

				options.SetCreateLogpushJobV2Request(createLogpushJobV2RequestModel)
				result, response, operationErr := testService.CreateLogpushJobV2(options)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				job := result.Result

				// List all Logpush Jobs
				listOptions := testService.NewGetLogpushJobsV2Options()
				getResult, getResponse, getErr := testService.GetLogpushJobsV2(listOptions)
				Expect(getErr).To(BeNil())
				Expect(getResponse).ToNot(BeNil())
				Expect(getResult).ToNot(BeNil())

				allJobs := getResult.Result

				// Get Logpush Job by jobID
				getJob := allJobs[0]
				getOptions := testService.NewGetLogpushJobV2Options(*getJob.ID)
				result, response, operationErr = testService.GetLogpushJobV2(getOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				//Update Logpush Jobs
				updateOptions := testService.NewUpdateLogpushJobV2Options(*job.ID)
				updateLogpushJobV2RequestModel := &logpushjobsapiv1.UpdateLogpushJobV2RequestLogpushJobsUpdateLogdnaReq{
					Enabled:        core.BoolPtr(false),
					LogpullOptions: core.StringPtr("timestamps=rfc3339&timestamps=rfc3339"),
					Logdna:         map[string]interface{}{"ingress_key": IngressKey, "region": LogdnaRegion, "hostname": LogdnaDomain},
					Frequency:      core.StringPtr("high"),
				}

				updateOptions.SetUpdateLogpushJobV2Request(updateLogpushJobV2RequestModel)

				updateResult, updateResponse, updateErr := testService.UpdateLogpushJobV2(updateOptions)
				Expect(updateErr).To(BeNil())
				Expect(updateResponse).ToNot(BeNil())
				Expect(updateResult).ToNot(BeNil())

				//Delete Logpush Jobs
				for _, thisJob := range allJobs {
					delOptions := testService.NewDeleteLogpushJobV2Options(*thisJob.ID)
					delResult, delResponse, delReleteErr := testService.DeleteLogpushJobV2(delOptions)
					Expect(delReleteErr).To(BeNil())
					Expect(delResponse).ToNot(BeNil())
					Expect(delResult).ToNot(BeNil())
				}
			})

			It(`Post\Validate Logpush Ownership challange`, func() {
				//Send ownership to destination
				ownershipOptions := testService.NewGetLogpushOwnershipV2Options()

				ownershipOptions.Cos = map[string]interface{}{"bucket_name": CosBucket, "region": CosRegion, "id": CosInstance}

				ownershipResult, ownershipResponse, ownershipErr := testService.GetLogpushOwnershipV2(ownershipOptions)
				Expect(ownershipErr).To(BeNil())
				Expect(ownershipResponse).ToNot(BeNil())
				Expect(ownershipResult).ToNot(BeNil())

				// Validate Logpush Ownership Challange
				validationOptions := testService.NewValidateLogpushOwnershipChallengeV2Options()

				validationOptions.Cos = map[string]interface{}{"bucket_name": CosBucket, "region": CosRegion, "id": CosInstance}
				validationOptions.OwnershipChallenge = core.StringPtr(OwnershipToken)

				validationResult, validationResponse, validationErr := testService.ValidateLogpushOwnershipChallengeV2(validationOptions)

				Expect(validationErr).To(BeNil())
				Expect(validationResponse).ToNot(BeNil())
				Expect(validationResult).ToNot(BeNil())

			})

			It(`List Fields/Jobs`, func() {
				//List available fields
				FieldOptions := testService.NewListFieldsForDatasetV2Options()
				fieldResult, fieldResponse, fieldErr := testService.ListFieldsForDatasetV2(FieldOptions)

				Expect(fieldErr).To(BeNil())
				Expect(fieldResponse).ToNot(BeNil())
				Expect(fieldResult).ToNot(BeNil())

				//List logpush jobs
				JobsOptions := testService.NewListLogpushJobsForDatasetV2Options()

				jobsResult, _, _ := testService.ListLogpushJobsForDatasetV2(JobsOptions)

				// Expect(jobsErr).To(BeNil())
				// Expect(jobsResponse).ToNot(BeNil())
				// Expect(jobsResult).ToNot(BeNil())

				// fmt.Println(jobsErr)
				// fmt.Println(jobsResponse)
				fmt.Println(jobsResult)
			})
		})
	})
})
