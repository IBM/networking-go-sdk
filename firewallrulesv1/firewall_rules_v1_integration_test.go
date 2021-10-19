/*
 * (C) Copyright IBM Corp. 2021.
 */

package firewallrulesv1_test

import (
	"fmt"
	"os"
	"strconv"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/joho/godotenv"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	filterv1 "github.com/IBM/networking-go-sdk/filtersv1"
	. "github.com/IBM/networking-go-sdk/firewallrulesv1"
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
	result, response, operationErr := testService.ListAllFilters(testService.NewListAllFiltersOptions(xAuthUserToken, crn, zoneId))
	Expect(operationErr).To(BeNil())
	Expect(response).ToNot(BeNil())
	Expect(result).ToNot(BeNil())
	// Get Filter IDs
	var Filter_IDs []string
	for i := 0; i < len(result.Result); i++ {
		Filter_IDs = append(Filter_IDs, *result.Result[i].ID)
	}

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
		"(ip.src eq 123.60.125.234)",
		"(http.request.uri eq \"/test?number=1\")",
		"not http.request.uri.path matches \"^/api/.*$\"",
		"(http.host eq \"testexample.com\")",
		"(http.user_agent eq \"Mozilla/5.0\")",
		"(ip.src eq 123.60.125.235)",
		"(http.request.uri eq \"/test-update?number=1\")",
		"not http.request.uri.path matches \"^/api-update\\d{2}/.*$\"",
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
	BeforeEach(func() {
		Skip("Skipping Tests")
	})

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
				for i := 0; i < len(result.Result); i++ {
					delOptions := service.NewDeleteFirewallRulesOptions(xAuthUserToken, crn, zoneId, *result.Result[i].ID)
					result, response, deleteErr := service.DeleteFirewallRules(delOptions)
					Expect(deleteErr).To(BeNil())
					Expect(response).ToNot(BeNil())
					Expect(result).ToNot(BeNil())
				}
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
				for i := 0; i < len(result.Result); i++ {
					delOptions := service.NewDeleteFirewallRulesOptions(xAuthUserToken, crn, zoneId, *result.Result[i].ID)
					result, response, deleteErr := service.DeleteFirewallRules(delOptions)
					Expect(deleteErr).To(BeNil())
					Expect(response).ToNot(BeNil())
					Expect(result).ToNot(BeNil())
				}

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

				//Update Firewall Rules
				for i := 0; i < 5; i++ {
					updateOption := service.NewUpdateFirewllRulesOptions(xAuthUserToken, crn, zoneId)
					Expect(updateOption).ToNot(BeNil())
					filterUpdate, filterErr := service.NewFirewallRulesUpdateInputItemFilter(filter_ids[i+5])
					Expect(filterErr).To(BeNil())
					Expect(filterUpdate).ToNot(BeNil())
					firewallRulesUpdate := &FirewallRulesUpdateInputItem{
						ID:          core.StringPtr(*result.Result[i].ID),
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
				for i := 0; i < len(result.Result); i++ {
					delOptions := service.NewDeleteFirewallRulesOptions(xAuthUserToken, crn, zoneId, *result.Result[i].ID)
					result, response, deleteErr := service.DeleteFirewallRules(delOptions)
					Expect(deleteErr).To(BeNil())
					Expect(response).ToNot(BeNil())
					Expect(result).ToNot(BeNil())
				}
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

				// List a Firewall rule by ID
				getFirewallRuleOptionsModel := service.NewGetFirewallRuleOptions(xAuthUserToken, crn, zoneId, *result.Result[0].ID)
				Expect(getFirewallRuleOptionsModel).ToNot(BeNil())
				resultGet, responseGet, errGet := service.GetFirewallRule(getFirewallRuleOptionsModel)
				Expect(errGet).To(BeNil())
				Expect(resultGet).ToNot(BeNil())
				Expect(responseGet).ToNot(BeNil())

				// Update a Firewall Rule
				updateFirewallRuleOptionsModel := service.NewUpdateFirewallRuleOptions(xAuthUserToken, crn, zoneId, *result.Result[0].ID)
				Expect(updateFirewallRuleOptionsModel).ToNot(BeNil())
				updateFirewallRuleOptionsModel.SetAction(actions[1])
				updateFirewallRuleOptionsModel.SetPaused(false)
				updateFirewallRuleOptionsModel.SetDescription("SDK JS challenge site Test")
				firewallFilterInput, firewallInputErr := service.NewFirewallRuleUpdateInputFilter(filter_ids[5])
				Expect(firewallInputErr).To(BeNil())
				Expect(firewallFilterInput).ToNot(BeNil())
				updateFirewallRuleOptionsModel.SetFilter(firewallFilterInput)
				resultUpdate, responseUpdate, errUpdate := service.UpdateFirewallRule(updateFirewallRuleOptionsModel)
				Expect(errUpdate).To(BeNil())
				Expect(resultUpdate).ToNot(BeNil())
				Expect(responseUpdate).ToNot(BeNil())

				//Delete a Firewall Rule
				delFirewallRuleOptionsModel := service.NewDeleteFirewallRuleOptions(xAuthUserToken, crn, zoneId, *result.Result[0].ID)
				Expect(delFirewallRuleOptionsModel).ToNot(BeNil())
				resultDel, responseDel, errDel := service.DeleteFirewallRule(delFirewallRuleOptionsModel)
				Expect(errDel).To(BeNil())
				Expect(resultDel).ToNot(BeNil())
				Expect(responseDel).ToNot(BeNil())
			})
		})
	})
})
