/*
 * (C) Copyright IBM Corp. 2020.
 */

package filtersv1_test

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/IBM/go-sdk-core/core"
	"github.com/joho/godotenv"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.ibm.com/ibmcloud/networking-go-sdk/filtersv1"
)

const configFile = "../cis.env"

var configLoaded bool = true

func shouldSkipTest() {
	if !configLoaded {
		Skip("External configuration is not available, skipping...")
	}
}

var _ = Describe(`FiltersV1`, func() {
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
		"(ip.src eq 93.60.125.234)",
		"(http.request.uri eq \"/test?number=1\")",
		"not http.request.uri.path matches \"^/api/.*$\"",
		"(http.host eq \"testexample.com\")",
	}

	expressions_update := [4]string{
		"(ip.src eq 93.60.125.235)",
		"(http.request.uri eq \"/test-update?number=1\")",
		"not http.request.uri.path matches \"^/api-update/.*$\"",
		"(http.host eq \"testexample-update.com\")",
	}

	type FilterResult struct {
		Result []struct {
			ID          string `json:"id"`
			Paused      bool   `json:"paused"`
			Description string `json:"description"`
			Expression  string `json:"expression"`
		} `json:"result"`
		Success    bool          `json:"success"`
		Errors     []interface{} `json:"errors"`
		Messages   []interface{} `json:"messages"`
		ResultInfo struct {
			Page       int `json:"page"`
			PerPage    int `json:"per_page"`
			Count      int `json:"count"`
			TotalCount int `json:"total_count"`
			TotalPages int `json:"total_pages"`
		} `json:"result_info"`
	}
	var Filter_IDs []string

	Describe(`FiltersApiv1_test`, func() {
		Context(`FiltersApiv1_all_filters`, func() {
			BeforeEach(func() {
				shouldSkipTest()
				result, response, operationErr := testService.ListAllFilters(testService.NewListAllFiltersOptions(xAuthUserToken, crn, zoneId))
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				//Delete all Flters
				Expect(result.Read).ToNot(BeNil())
				p := make([]byte, 1024)
				f, err := os.Create("/tmp/filters.txt")
				Expect(err).To(BeNil())
				n, err := result.Read(p)
				Expect(err).To(BeNil())
				for len(p) > 0 {
					_, err := f.Write(p[:n])
					if err != nil {
						break
					}
					n, err = result.Read(p[:n])
					if err != nil {
						break
					}
				}
				err = f.Close()
				Expect(err).To(BeNil())
				err = result.Close()
				Expect(err).To(BeNil())
				byteValue, err := ioutil.ReadFile("/tmp/filters.txt")
				Expect(err).To(BeNil())
				filter_res := FilterResult{}
				json_err := json.Unmarshal(byteValue, &filter_res)
				Expect(json_err).To(BeNil())

				for i := 0; i < len(filter_res.Result); i++ {
					delOptions := testService.NewDeleteFiltersOptions(xAuthUserToken, crn, zoneId, filter_res.Result[i].ID)
					result, response, deleteErr := testService.DeleteFilters(delOptions)
					Expect(deleteErr).To(BeNil())
					Expect(response).ToNot(BeNil())
					Expect(result).ToNot(BeNil())
				}
				err = os.Remove("/tmp/filters.txt")
				Expect(err).To(BeNil())
			})
			AfterEach(func() {
				shouldSkipTest()
				result, response, operationErr := testService.ListAllFilters(testService.NewListAllFiltersOptions(xAuthUserToken, crn, zoneId))
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
				Expect(result.Read).ToNot(BeNil())

				//Delete all Flters
				Expect(result.Read).ToNot(BeNil())
				p := make([]byte, 1024)
				f, err := os.Create("/tmp/filters.txt")
				Expect(err).To(BeNil())
				n, err := result.Read(p)
				Expect(err).To(BeNil())
				for len(p) > 0 {
					_, err := f.Write(p[:n])
					if err != nil {
						break
					}
					n, err = result.Read(p[:n])
					if err != nil {
						break
					}
				}
				err = f.Close()
				Expect(err).To(BeNil())
				err = result.Close()
				Expect(err).To(BeNil())
				byteValue, err := ioutil.ReadFile("/tmp/filters.txt")
				Expect(err).To(BeNil())
				filter_res := FilterResult{}
				json_err := json.Unmarshal(byteValue, &filter_res)
				Expect(json_err).To(BeNil())

				for i := 0; i < len(filter_res.Result); i++ {
					delOptions := testService.NewDeleteFiltersOptions(xAuthUserToken, crn, zoneId, filter_res.Result[i].ID)
					result, response, deleteErr := testService.DeleteFilters(delOptions)
					Expect(deleteErr).To(BeNil())
					Expect(response).ToNot(BeNil())
					Expect(result).ToNot(BeNil())
				}
				err = os.Remove("/tmp/filters.txt")
				Expect(err).To(BeNil())
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
				Expect(result.Read).ToNot(BeNil())
				p := make([]byte, 1024)
				f, err := os.Create("/tmp/filters.txt")
				Expect(err).To(BeNil())
				n, err := result.Read(p)
				Expect(err).To(BeNil())
				for len(p) > 0 {
					_, err := f.Write(p[:n])
					if err != nil {
						break
					}
					n, err = result.Read(p[:n])
					if err != nil {
						break
					}
				}
				err = f.Close()
				Expect(err).To(BeNil())
				err = result.Close()
				Expect(err).To(BeNil())
				// Get Filter IDs
				byteValue, err := ioutil.ReadFile("/tmp/filters.txt")
				Expect(err).To(BeNil())
				filter_res := FilterResult{}
				json_err := json.Unmarshal(byteValue, &filter_res)
				Expect(json_err).To(BeNil())
				//Update Filters
				for i := 0; i < len(filter_res.Result); i++ {
					options := testService.NewUpdateFiltersOptions(xAuthUserToken, crn, zoneId)
					filterUpdateInput := &FilterUpdateInput{
						ID:          core.StringPtr(filter_res.Result[i].ID),
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

				err = os.Remove("/tmp/filters.txt")
				Expect(err).To(BeNil())
			})
		})
	})
	Describe(`Individual Filter`, func() {
		Context(`Filter`, func() {
			It(`List\Update\Delete a Filter`, func() {
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
				result, response, listErr := testService.ListAllFilters(testService.NewListAllFiltersOptions(xAuthUserToken, crn, zoneId))
				Expect(listErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
				Expect(result.Read).ToNot(BeNil())
				p := make([]byte, 1024)
				f, err := os.Create("/tmp/filters.txt")
				Expect(err).To(BeNil())
				n, err := result.Read(p)
				Expect(err).To(BeNil())
				for len(p) > 0 {
					_, err := f.Write(p[:n])
					if err != nil {
						break
					}
					n, err = result.Read(p[:n])
					if err != nil {
						break
					}
				}
				err = f.Close()
				Expect(err).To(BeNil())
				err = result.Close()
				Expect(err).To(BeNil())
				// Get Filter ID
				byteValue, err := ioutil.ReadFile("/tmp/filters.txt")
				Expect(err).To(BeNil())
				filter_res := FilterResult{}
				json_err := json.Unmarshal(byteValue, &filter_res)
				Expect(json_err).To(BeNil())

				filter_id := filter_res.Result[0].ID

				//Get a Filter Info
				optionsGet := testService.NewGetFilterOptions(xAuthUserToken, crn, zoneId, filter_id)
				result, response, operationErr = testService.GetFilter(optionsGet)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				//Update a Filter
				optionsUpdate := testService.NewUpdateFilterOptions(xAuthUserToken, crn, zoneId, filter_id)
				optionsUpdate.SetExpression(`not http.request.uri.path matches "^/api/.*$"`)
				optionsUpdate.SetDescription("not /api")
				optionsUpdate.SetPaused(false)
				result, response, operationErr = testService.UpdateFilter(optionsUpdate)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				//Delete a Filter
				optionsDelete := testService.NewDeleteFilterOptions(xAuthUserToken, crn, zoneId, filter_id)
				result, response, operationErr = testService.DeleteFilter(optionsDelete)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
})
