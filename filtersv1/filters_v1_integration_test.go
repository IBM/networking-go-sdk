/*
 * (C) Copyright IBM Corp. 2021.
 */

package filtersv1_test

import (
	"os"
	"strconv"

	"github.com/IBM/go-sdk-core/v5/core"
	. "github.com/IBM/networking-go-sdk/filtersv1"
	"github.com/joho/godotenv"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const configFile = "../cis.env"

var configLoaded bool = true

func shouldSkipTest() {
	if !configLoaded {
		Skip("External configuration is not available, skipping...")
	}
}

var _ = Describe(`FiltersV1`, func() {
	// BeforeEach(func() {
	// 	Skip("Skipping Tests")
	// })

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
	serviceURL := os.Getenv("API_ENDPOINT")
	crn := os.Getenv("CRN")
	zoneId := os.Getenv("ZONE_ID")
	xAuthUserToken := os.Getenv("CIS_SERVICES_APIKEY")
	globalOptions := &FiltersV1Options{
		ServiceName:   "cis_services",
		URL:           serviceURL,
		Authenticator: authenticator,
	}
	testService, testServiceErr := NewFiltersV1(globalOptions)
	Expect(testServiceErr).To(BeNil())

	expressions := [4]string{
		"(ip.src eq 13.60.125.234)",
		"(http.request.uri eq \"/test?number=1\")",
		"not http.request.uri.path matches \"^/api/[\\W].*$\"",
		"(http.request.uri.path ~ \"^.*/wpt[\\d]-login.php$\" or http.request.uri.path ~ \"^.*/xmlrpc.php$\")",
	}

	expressions_update := [4]string{
		"(ip.src eq 13.60.125.235)",
		"(http.request.uri eq \"/test-update?number=1\")",
		"not http.request.uri.path matches \"^/api-update/.*$\"",
		"(http.host eq \"testexample-update.com\")",
	}

	var Filter_IDs []string

	Describe(`FiltersApiv1_test`, func() {
		Context(`FiltersApiv1_all_filters`, func() {
			defer GinkgoRecover()
			BeforeEach(func() {
				shouldSkipTest()
				result, response, operationErr := testService.ListAllFilters(testService.NewListAllFiltersOptions(xAuthUserToken, crn, zoneId))
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				//Delete all Flters
				for i := 0; i < len(result.Result); i++ {
					delOptions := testService.NewDeleteFiltersOptions(xAuthUserToken, crn, zoneId, *result.Result[i].ID)
					result, response, deleteErr := testService.DeleteFilters(delOptions)
					Expect(deleteErr).To(BeNil())
					Expect(response).ToNot(BeNil())
					Expect(result).ToNot(BeNil())
				}

			})
			AfterEach(func() {
				shouldSkipTest()
				result, response, operationErr := testService.ListAllFilters(testService.NewListAllFiltersOptions(xAuthUserToken, crn, zoneId))
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				//Delete all Flters
				for i := 0; i < len(result.Result); i++ {
					delOptions := testService.NewDeleteFiltersOptions(xAuthUserToken, crn, zoneId, *result.Result[i].ID)
					result, response, deleteErr := testService.DeleteFilters(delOptions)
					Expect(deleteErr).To(BeNil())
					Expect(response).ToNot(BeNil())
					Expect(result).ToNot(BeNil())
				}
			})
			It(`Create Filters\ListAll Filters\Update Filters\Delete Filters`, func() {
				shouldSkipTest()
				//Create Filters
				for i := 0; i < 4; i++ {
					options := testService.NewCreateFilterOptions(xAuthUserToken, crn, zoneId)
					filetrInput := &FilterInput{
						Expression:  core.StringPtr(expressions[i]),
						Paused:      core.BoolPtr(false),
						Description: core.StringPtr("Login-Office-SDK" + strconv.Itoa(i)),
					}
					options.SetFilterInput([]FilterInput{*filetrInput})
					result, response, operationErr := testService.CreateFilter(options)
					Expect(operationErr).To(BeNil())
					Expect(response).ToNot(BeNil())
					Expect(result).ToNot(BeNil())
				}

				// List all Filters
				result, response, operationErr := testService.ListAllFilters(testService.NewListAllFiltersOptions(xAuthUserToken, crn, zoneId))
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				//Update Filters
				for i := 0; i < len(result.Result); i++ {
					options := testService.NewUpdateFiltersOptions(xAuthUserToken, crn, zoneId)
					filterUpdateInput := &FilterUpdateInput{
						ID:          core.StringPtr(*result.Result[i].ID),
						Expression:  core.StringPtr(expressions_update[i]),
						Paused:      core.BoolPtr(false),
						Description: core.StringPtr("Login-SDK-Update" + strconv.Itoa(i)),
					}
					options.SetFilterUpdateInput([]FilterUpdateInput{*filterUpdateInput})
					result, response, updateErr := testService.UpdateFilters(options)
					Expect(updateErr).To(BeNil())
					Expect(response).ToNot(BeNil())
					Expect(result).ToNot(BeNil())
				}

				//Delete Filters
				for i := 0; i < len(Filter_IDs); i++ {
					delOptions := testService.NewDeleteFiltersOptions(xAuthUserToken, crn, zoneId, Filter_IDs[i])
					result, response, deleteErr := testService.DeleteFilters(delOptions)
					Expect(deleteErr).To(BeNil())
					Expect(response).ToNot(BeNil())
					Expect(result).ToNot(BeNil())
				}

			})
		})
	})
	Describe(`Individual Filter`, func() {
		Context(`Filter`, func() {
			It(`List\Update\Delete a Filter`, func() {
				shouldSkipTest()
				//Create a Filter
				options := testService.NewCreateFilterOptions(xAuthUserToken, crn, zoneId)
				filetrInput := &FilterInput{
					Expression:  core.StringPtr(expressions[0]),
					Paused:      core.BoolPtr(false),
					Description: core.StringPtr("Login-Office-SDK-Single-Filter"),
				}
				options.SetFilterInput([]FilterInput{*filetrInput})
				result, response, operationErr := testService.CreateFilter(options)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				//Get Filter ID
				filter_id := result.Result[0].ID

				//Get a Filter Info
				optionsGet := testService.NewGetFilterOptions(xAuthUserToken, crn, zoneId, *filter_id)
				resultGet, responseGet, operationGetErr := testService.GetFilter(optionsGet)
				Expect(operationGetErr).To(BeNil())
				Expect(responseGet).ToNot(BeNil())
				Expect(resultGet).ToNot(BeNil())

				//Update a Filter
				optionsUpdate := testService.NewUpdateFilterOptions(xAuthUserToken, crn, zoneId, *filter_id)
				optionsUpdate.SetExpression(`not http.request.uri.path matches "^/api/.*$"`)
				optionsUpdate.SetDescription("not /api")
				optionsUpdate.SetPaused(false)
				resultUpdate, responseUpdate, operationUpdateErr := testService.UpdateFilter(optionsUpdate)
				Expect(operationUpdateErr).To(BeNil())
				Expect(responseUpdate).ToNot(BeNil())
				Expect(resultUpdate).ToNot(BeNil())

				//Delete a Filter
				optionsDelete := testService.NewDeleteFilterOptions(xAuthUserToken, crn, zoneId, *filter_id)
				resultDel, responseDel, operationDelErr := testService.DeleteFilter(optionsDelete)
				Expect(operationDelErr).To(BeNil())
				Expect(responseDel).ToNot(BeNil())
				Expect(resultDel).ToNot(BeNil())
			})
		})
	})
})
