/*
 * (C) Copyright IBM Corp. 2021.
 */

package firewallrulesv1_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/IBM/go-sdk-core/core"
	. "github.com/IBM/networking-go-sdk/firewallrulesv1"
	"github.com/joho/godotenv"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	filterv1 "github.com/IBM/networking-go-sdk/filtersv1" 
)

const configFile = "../cis.env"

var configLoaded bool = true

func shouldSkipTest() {
	if !configLoaded {
		Skip("External configuration is not available, skipping...")
	}
}

func getExistingFilterIds(testService *filterv1.FiltersV1, xAuthUserToken string, crn string, zoneId string) []string {
	// List all Filters
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
	result, response, operationErr := testService.ListAllFilters(testService.NewListAllFiltersOptions(xAuthUserToken, crn, zoneId))
	Expect(operationErr).To(BeNil())
	Expect(response).ToNot(BeNil())
	Expect(result).ToNot(BeNil())
	Expect(result.Read).ToNot(BeNil())
	p := make([]byte, 1024)
	f, err := os.Create("/tmp/filtersforfirewallrules.txt")
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
	byteValue, err := ioutil.ReadFile("/tmp/filtersforfirewallrules.txt")
	Expect(err).To(BeNil())
	filter_res := FilterResult{}
	json_err := json.Unmarshal(byteValue, &filter_res)
	Expect(json_err).To(BeNil())
	//Return Filter Ids
	var Filter_IDs []string
	for i := 0; i < len(filter_res.Result); i++ {
		Filter_IDs = append(Filter_IDs, filter_res.Result[i].ID)
	}
	err = os.Remove("/tmp/filtersforfirewallrules.txt")
	Expect(err).To(BeNil())
	return (Filter_IDs)
}

func createFilters(options *filterv1.FiltersV1Options, xAuthUserToken string, crn string, zoneId string) []string {
	//Get all existing Filters
	testService, testServiceErr := filterv1.NewFiltersV1(options)
	Expect(testServiceErr).To(BeNil())
	//Delete Filters
	Exiting_Filter_IDs := getExistingFilterIds(testService, xAuthUserToken, crn, zoneId)
	for i := 0; i < len(Exiting_Filter_IDs); i++ {
		delOptions := testService.NewDeleteFiltersOptions(xAuthUserToken, crn, zoneId, Exiting_Filter_IDs[i])
		result, response, deleteErr := testService.DeleteFilters(delOptions)
		Expect(deleteErr).To(BeNil())
		Expect(response).ToNot(BeNil())
		Expect(result).ToNot(BeNil())
	}
	//Create New Filters
	expressions := [10]string{
		"(ip.src eq 93.60.125.234)",
		"(http.request.uri eq \"/test?number=1\")",
		"not http.request.uri.path matches \"^/api/.*$\"",
		"(http.host eq \"testexample.com\")",
		"(http.user_agent eq \"Mozilla/5.0\")",
		"(ip.src eq 93.60.125.235)",
		"(http.request.uri eq \"/test-update?number=1\")",
		"not http.request.uri.path matches \"^/api-update/.*$\"",
		"(http.host eq \"testexample-update.com\")",
		"(http.user_agent eq \"Mozilla/6.0\")",
	}
	for i := 0; i < 10; i++ {
		options := testService.NewCreateFilterOptions(xAuthUserToken, crn, zoneId)
		filetrInput := &filterv1.FilterInput{
			Expression:  core.StringPtr(expressions[i]),
			Paused:      core.BoolPtr(false),
			Description: core.StringPtr("Login-Office-SDK" + strconv.Itoa(i)),
		}
		options.SetFilterInput([]filterv1.FilterInput{*filetrInput})
		result, response, operationErr := testService.CreateFilter(options)
		Expect(operationErr).To(BeNil())
		Expect(response).ToNot(BeNil())
		Expect(result).ToNot(BeNil())
	}
	// List all Filters
	New_Filter_IDs := getExistingFilterIds(testService, xAuthUserToken, crn, zoneId)

	return (New_Filter_IDs)
}

var _ = Describe(`firewallapiv1_test`, func() {
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
	globalOptions := &FirewallRulesV1Options{
		ServiceName:   "cis_services",
		URL:           serviceURL,
		Authenticator: authenticator,
	}
	filterOptions := &filterv1.FiltersV1Options{
		ServiceName:   "cis_services",
		URL:           serviceURL,
		Authenticator: authenticator,
	}
	service, serviceErr := NewFirewallRulesV1(globalOptions)
	Expect(service).ToNot(BeNil())
	Expect(serviceErr).To(BeNil())
	if serviceErr != nil {
		fmt.Println(serviceErr)
	}
	actions := [5]string{
		FirewallRuleInputWithFilterID_Action_Allow,
		FirewallRuleInputWithFilterID_Action_Block,
		FirewallRuleInputWithFilterID_Action_Challenge,
		FirewallRuleInputWithFilterID_Action_JsChallenge,
		FirewallRuleInputWithFilterID_Action_Log,
	}
	type FirewallRulesResult struct {
		Result []struct {
			ID          string `json:"id"`
			Paused      bool   `json:"paused"`
			Description string `json:"description"`
			Action      string `json:"action"`
			Filter      struct {
				ID          string `json:"id"`
				Paused      bool   `json:"paused"`
				Description string `json:"description"`
				Expression  string `json:"expression"`
			} `json:"filter"`
			CreatedOn  string `json:"created_on"`
			ModifiedOn string `json:"modified_on"`
		} `json:"result"`
		Success bool `json:"success"`
		Errors  []struct {
		} `json:"errors"`
		Messages []struct {
		} `json:"messages"`
		ResultInfo struct {
			Page       int `json:"page"`
			PerPage    int `json:"per_page"`
			Count      int `json:"count"`
			TotalCount int `json:"total_count"`
		} `json:"result_info"`
	}
	Describe(`firewallrulesv1_test`, func() {
		Context(`firewallrulesv1_test`, func() {
			BeforeEach(func() {
				shouldSkipTest()
				listAllFirewallRulesOptionsModel := service.NewListAllFirewallRulesOptions(xAuthUserToken, crn, zoneId)
				Expect(listAllFirewallRulesOptionsModel).ToNot(BeNil())
				result, resp, listErr := service.ListAllFirewallRules(listAllFirewallRulesOptionsModel)
				Expect(listErr).To(BeNil())
				Expect(resp).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
				//Delete all Firewall Rules
				Expect(result.Read).ToNot(BeNil())
				p := make([]byte, 1024)
				f, err := os.Create("/tmp/firewallRules.txt")
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
				byteValue, err := ioutil.ReadFile("/tmp/firewallRules.txt")
				Expect(err).To(BeNil())
				firewall_res := FirewallRulesResult{}
				json_err := json.Unmarshal(byteValue, &firewall_res)
				Expect(json_err).To(BeNil())

				for i := 0; i < len(firewall_res.Result); i++ {
					delOptions := service.NewDeleteFirewallRulesOptions(xAuthUserToken, crn, zoneId, firewall_res.Result[i].ID)
					result, response, deleteErr := service.DeleteFirewallRules(delOptions)
					Expect(deleteErr).To(BeNil())
					Expect(response).ToNot(BeNil())
					Expect(result).ToNot(BeNil())
				}
				//Remove firewallRules file
				err = os.Remove("/tmp/firewallRules.txt")
				Expect(err).To(BeNil())
			})
			AfterEach(func() {
				shouldSkipTest()
				listAllFirewallRulesOptionsModel := service.NewListAllFirewallRulesOptions(xAuthUserToken, crn, zoneId)
				Expect(listAllFirewallRulesOptionsModel).ToNot(BeNil())
				result, resp, listErr := service.ListAllFirewallRules(listAllFirewallRulesOptionsModel)
				Expect(listErr).To(BeNil())
				Expect(resp).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
				//Delete all Firewall Rules
				Expect(result.Read).ToNot(BeNil())
				p := make([]byte, 1024)
				f, err := os.Create("/tmp/firewallRules.txt")
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
				byteValue, err := ioutil.ReadFile("/tmp/firewallRules.txt")
				Expect(err).To(BeNil())
				firewall_res := FirewallRulesResult{}
				json_err := json.Unmarshal(byteValue, &firewall_res)
				Expect(json_err).To(BeNil())

				for i := 0; i < len(firewall_res.Result); i++ {
					delOptions := service.NewDeleteFirewallRulesOptions(xAuthUserToken, crn, zoneId, firewall_res.Result[i].ID)
					result, response, deleteErr := service.DeleteFirewallRules(delOptions)
					Expect(deleteErr).To(BeNil())
					Expect(response).ToNot(BeNil())
					Expect(result).ToNot(BeNil())
				}
				err = os.Remove("/tmp/firewallRules.txt")
				Expect(err).To(BeNil())

			})
			It(`Create Firewall Rules | List Firewall Rules | Update Firewall Rules | Delete Firewall Rules`, func() {
				// Create Filters
				filter_ids := createFilters(filterOptions, xAuthUserToken, crn, zoneId)
				for i := 0; i < 5; i++ {
					createFirewallRulesOptionsModel := service.NewCreateFirewallRulesOptions(xAuthUserToken, crn, zoneId)
					Expect(createFirewallRulesOptionsModel).ToNot(BeNil())
					filterModel, err := service.NewFirewallRuleInputWithFilterIdFilter(filter_ids[i])
					Expect(filterModel).ToNot(BeNil())
					Expect(err).To(BeNil())
					firewallRules := &FirewallRuleInputWithFilterID{
						Filter:      filterModel,
						Action:      core.StringPtr(actions[i]),
						Description: core.StringPtr("Login-Office-SDK-Test" + strconv.Itoa(i)),
					}

					createFirewallRulesOptionsModel.SetFirewallRuleInputWithFilterID([]FirewallRuleInputWithFilterID{*firewallRules})
					result, response, operationErr := service.CreateFirewallRules(createFirewallRulesOptionsModel)
					Expect(operationErr).To(BeNil())
					Expect(response).ToNot(BeNil())
					Expect(result).ToNot(BeNil())
				}

				// List Firewall Rules
				listAllFirewallRulesOptionsModel := service.NewListAllFirewallRulesOptions(xAuthUserToken, crn, zoneId)
				Expect(listAllFirewallRulesOptionsModel).ToNot(BeNil())
				result, resp, listErr := service.ListAllFirewallRules(listAllFirewallRulesOptionsModel)
				Expect(listErr).To(BeNil())
				Expect(resp).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
				//Delete all Firewall Rules
				Expect(result.Read).ToNot(BeNil())
				p := make([]byte, 1024)
				f, err := os.Create("/tmp/firewallRules.txt")
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
				byteValue, err := ioutil.ReadFile("/tmp/firewallRules.txt")
				Expect(err).To(BeNil())
				firewall_res := FirewallRulesResult{}
				json_err := json.Unmarshal(byteValue, &firewall_res)
				Expect(json_err).To(BeNil())

				//Update Firewall Rules
				for i := 0; i < 5; i++ {
					updateOption := service.NewUpdateFirewllRulesOptions(xAuthUserToken, crn, zoneId)
					Expect(updateOption).ToNot(BeNil())
					filterUpdate, filterErr := service.NewFirewallRulesUpdateInputItemFilter(filter_ids[i+5])
					Expect(filterErr).To(BeNil())
					Expect(filterUpdate).ToNot(BeNil())
					firewallRulesUpdate := &FirewallRulesUpdateInputItem{
						ID:          core.StringPtr(firewall_res.Result[i].ID),
						Action:      core.StringPtr(actions[(i+1)%5]),
						Description: core.StringPtr("Firewall-Rules-Update-SDK-Test" + strconv.Itoa(i)),
						Filter:      filterUpdate,
					}
					updateOption.SetFirewallRulesUpdateInputItem([]FirewallRulesUpdateInputItem{*firewallRulesUpdate})
					updateResult, updateResponse, updateErr := service.UpdateFirewllRules(updateOption)
					Expect(updateErr).To(BeNil())
					Expect(updateResult).ToNot(BeNil())
					Expect(updateResponse).ToNot(BeNil())
				}
				// Delete Firewall Rules
				for i := 0; i < len(firewall_res.Result); i++ {
					delOptions := service.NewDeleteFirewallRulesOptions(xAuthUserToken, crn, zoneId, firewall_res.Result[i].ID)
					result, response, deleteErr := service.DeleteFirewallRules(delOptions)
					Expect(deleteErr).To(BeNil())
					Expect(response).ToNot(BeNil())
					Expect(result).ToNot(BeNil())
				}
				//Remove firewallRules file
				err = os.Remove("/tmp/firewallRules.txt")
				Expect(err).To(BeNil())

			})
			It(`List | Update | Delete single Firewall Rule`, func() {
				//List Firewall rule
				filter_ids := createFilters(filterOptions, xAuthUserToken, crn, zoneId)
				for i := 0; i < 5; i++ {
					createFirewallRulesOptionsModel := service.NewCreateFirewallRulesOptions(xAuthUserToken, crn, zoneId)
					Expect(createFirewallRulesOptionsModel).ToNot(BeNil())
					filterModel, err := service.NewFirewallRuleInputWithFilterIdFilter(filter_ids[i])
					Expect(filterModel).ToNot(BeNil())
					Expect(err).To(BeNil())
					firewallRules := &FirewallRuleInputWithFilterID{
						Filter:      filterModel,
						Action:      core.StringPtr(actions[i]),
						Description: core.StringPtr("Login-Office-SDK-Test" + strconv.Itoa(i)),
					}
					createFirewallRulesOptionsModel.SetFirewallRuleInputWithFilterID([]FirewallRuleInputWithFilterID{*firewallRules})
					result, response, operationErr := service.CreateFirewallRules(createFirewallRulesOptionsModel)
					Expect(operationErr).To(BeNil())
					Expect(response).ToNot(BeNil())
					Expect(result).ToNot(BeNil())
				}

				// List Firewall Rules
				listAllFirewallRulesOptionsModel := service.NewListAllFirewallRulesOptions(xAuthUserToken, crn, zoneId)
				Expect(listAllFirewallRulesOptionsModel).ToNot(BeNil())
				result, resp, listErr := service.ListAllFirewallRules(listAllFirewallRulesOptionsModel)
				Expect(listErr).To(BeNil())
				Expect(resp).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
				//Delete all Firewall Rules
				Expect(result.Read).ToNot(BeNil())
				p := make([]byte, 1024)
				f, err := os.Create("/tmp/firewallRule.txt")
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
				byteValue, err := ioutil.ReadFile("/tmp/firewallRule.txt")
				Expect(err).To(BeNil())
				firewall_res := FirewallRulesResult{}
				json_err := json.Unmarshal(byteValue, &firewall_res)
				Expect(json_err).To(BeNil())

				// List a Firewall rule by ID
				getFirewallRuleOptionsModel := service.NewGetFirewallRuleOptions(xAuthUserToken, crn, zoneId, firewall_res.Result[0].ID)
				Expect(getFirewallRuleOptionsModel).ToNot(BeNil())
				result, response, err := service.GetFirewallRule(getFirewallRuleOptionsModel)
				Expect(err).To(BeNil())
				Expect(result).ToNot(BeNil())
				Expect(response).ToNot(BeNil())

				// Update a Firewall Rule
				updateFirewallRuleOptionsModel := service.NewUpdateFirewallRuleOptions(xAuthUserToken, crn, zoneId, firewall_res.Result[0].ID)
				Expect(updateFirewallRuleOptionsModel).ToNot(BeNil())
				updateFirewallRuleOptionsModel.SetAction(actions[1])
				updateFirewallRuleOptionsModel.SetPaused(false)
				updateFirewallRuleOptionsModel.SetDescription("SDK JS challenge site Test")
				firewallFilterInput, firewallInputErr := service.NewFirewallRuleUpdateInputFilter(filter_ids[5])
				Expect(firewallInputErr).To(BeNil())
				Expect(firewallFilterInput).ToNot(BeNil())
				updateFirewallRuleOptionsModel.SetFilter(firewallFilterInput)
				result, response, err = service.UpdateFirewallRule(updateFirewallRuleOptionsModel)
				Expect(err).To(BeNil())
				Expect(result).ToNot(BeNil())
				Expect(response).ToNot(BeNil())

				//Delete a Firewall Rule
				delFirewallRuleOptionsModel := service.NewDeleteFirewallRuleOptions(xAuthUserToken, crn, zoneId, firewall_res.Result[0].ID)
				Expect(delFirewallRuleOptionsModel).ToNot(BeNil())
				result, response, err = service.DeleteFirewallRule(delFirewallRuleOptionsModel)
				Expect(err).To(BeNil())
				Expect(result).ToNot(BeNil())
				Expect(response).ToNot(BeNil())

				//Remove firewallRules file
				err = os.Remove("/tmp/firewallRule.txt")
				Expect(err).To(BeNil())
			})
		})
	})
})
