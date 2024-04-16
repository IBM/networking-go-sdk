/**
 * (C) Copyright IBM Corp. 2024.
 */

package rulesetsv1_test

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/networking-go-sdk/rulesetsv1"
	"github.com/joho/godotenv"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

/**
 * This file contains an integration test for the rulesetsv1 package.
 *
 * Notes:
 *
 * The integration test will automatically skip tests if the required config file is not available.
 */

const configFile = "../cis.env"

var configLoaded bool = true
var authenticationSucceeded bool = true

func shouldSkipTest() {
	Skip("Skipping..")
	if !configLoaded {
		Skip("External configuration is not available, skipping...")
	}

	if !authenticationSucceeded {
		Skip("Authentication failed. Check external configuration...")
	}
}

var _ = Describe(`RulesetsV1 Integration Tests`, func() {

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
	authErr := authenticator.Authenticate(&http.Request{
		Header: http.Header{},
	})
	if authErr != nil {
		authenticationSucceeded = false
		fmt.Println("Authentication error during setup: ", authErr)
	}
	serviceURL := os.Getenv("API_ENDPOINT")
	crn := os.Getenv("CRN")
	zoneID := os.Getenv("ZONE_ID")

	var rulesetsService *rulesetsv1.RulesetsV1
	var rulesetToDeployId *string
	var rule1Id *string
	var rule2Id *string
	var rulsetForTestingId *string

	Describe("Client initialization", func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It("Successfully construct the service client instance", func() {
			rulesetsServiceOptions := &rulesetsv1.RulesetsV1Options{
				ServiceName:    "rulesets",
				URL:            serviceURL,
				Crn:            &crn,
				ZoneIdentifier: &zoneID,
				Authenticator:  authenticator,
			}

			rulesetsService, err = rulesetsv1.NewRulesetsV1UsingExternalConfig(rulesetsServiceOptions)
			Expect(err).To(BeNil())
			Expect(rulesetsService).ToNot(BeNil())
			Expect(rulesetsService.Service.Options.URL).To(Equal(serviceURL))

			rulesetsService.EnableRetries(4, 30*time.Second)
		})
	})

	//--------------ACCOUNT RULESETS INTEGRATION TEST-----------------

	Describe("account rulesets", func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It("List/Get account ruleset", func() {
			getAccountRulesetsOptions := &rulesetsv1.GetAccountRulesetsOptions{}

			listRulesetsResp, response, err := rulesetsService.GetAccountRulesets(getAccountRulesetsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(listRulesetsResp).ToNot(BeNil())

			rulesetToDeployId = listRulesetsResp.Result[0].ID
			getAccountRulesetOptions := &rulesetsv1.GetAccountRulesetOptions{
				RulesetID: rulesetToDeployId,
			}

			rulesetResp, response, err := rulesetsService.GetAccountRuleset(getAccountRulesetOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(rulesetResp).ToNot(BeNil())

			rule1Id = rulesetResp.Result.Rules[0].ID
			rule2Id = rulesetResp.Result.Rules[1].ID

		})

		It("Update Account Entrypoint Ruleset", func() {

			// if there is no ep this API will create and ep ruleset.
			// here we are creating a rule which will deploy a managed ruleset
			// Also we are overriding a rule already present in the managed ruleset

			rulesOverrideModel := &rulesetsv1.RulesOverride{
				ID:      rule1Id,
				Enabled: core.BoolPtr(true),
				Action:  core.StringPtr("block"),
			}

			categoriesOverrideModel := &rulesetsv1.CategoriesOverride{
				Category: core.StringPtr("wordpress"),
				Enabled:  core.BoolPtr(true),
				Action:   core.StringPtr("log"),
			}

			overridesModel := &rulesetsv1.Overrides{
				Action:     core.StringPtr("log"),
				Enabled:    core.BoolPtr(true),
				Rules:      []rulesetsv1.RulesOverride{*rulesOverrideModel},
				Categories: []rulesetsv1.CategoriesOverride{*categoriesOverrideModel},
			}

			actionParametersModel := &rulesetsv1.ActionParameters{
				ID:        rulesetToDeployId,
				Overrides: overridesModel,
			}
			positionModel := &rulesetsv1.Position{
				Index: core.Int64Ptr(int64(1)),
			}

			ruleCreateModel := &rulesetsv1.RuleCreate{
				Action:           core.StringPtr("execute"),
				ActionParameters: actionParametersModel,
				Description:      core.StringPtr("Overriding rule"),
				Enabled:          core.BoolPtr(true),
				Expression:       core.StringPtr("(ip.src ne 1.1.1.1) and cf.zone.plan eq \"ENT\""),
				Ref:              rulesetToDeployId,
				Position:         positionModel,
			}

			updateAccountEntrypointRulesetOptions := &rulesetsv1.UpdateAccountEntrypointRulesetOptions{
				RulesetPhase: core.StringPtr("http_request_firewall_managed"),
				Description:  core.StringPtr("creating/updating entrypoint ruleset"),
				Rules:        []rulesetsv1.RuleCreate{*ruleCreateModel},
			}

			rulesetResp, response, err := rulesetsService.UpdateAccountEntrypointRuleset(updateAccountEntrypointRulesetOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(rulesetResp).ToNot(BeNil())

			rulsetForTestingId = rulesetResp.Result.ID
		})

		It("Update Account Ruleset", func() {

			// to get the ruleset id

			getAccountEntrypointRulesetOptions := &rulesetsv1.GetAccountEntrypointRulesetOptions{
				RulesetPhase: core.StringPtr("http_request_firewall_managed"),
			}
			rulesetResp, _, _ := rulesetsService.GetAccountEntrypointRuleset(getAccountEntrypointRulesetOptions)
			ruleset2UpdateId := *rulesetResp.Result.ID

			rulesOverrideModel := &rulesetsv1.RulesOverride{
				ID:      rule1Id,
				Enabled: core.BoolPtr(true),
				Action:  core.StringPtr("block"),
			}

			categoriesOverrideModel := &rulesetsv1.CategoriesOverride{
				Category: core.StringPtr("wordpress"),
				Enabled:  core.BoolPtr(true),
				Action:   core.StringPtr("block"),
			}

			overridesModel := &rulesetsv1.Overrides{
				Action:     core.StringPtr("block"),
				Enabled:    core.BoolPtr(true),
				Rules:      []rulesetsv1.RulesOverride{*rulesOverrideModel},
				Categories: []rulesetsv1.CategoriesOverride{*categoriesOverrideModel},
			}

			actionParametersModel := &rulesetsv1.ActionParameters{
				ID:        rulesetToDeployId,
				Overrides: overridesModel,
			}

			positionModel := &rulesetsv1.Position{
				Index: core.Int64Ptr(int64(1)),
			}

			ruleCreateModel := &rulesetsv1.RuleCreate{
				Action:           core.StringPtr("execute"),
				ActionParameters: actionParametersModel,
				Description:      core.StringPtr("update rules"),
				Enabled:          core.BoolPtr(true),
				Expression:       core.StringPtr("(ip.src ne 1.1.1.2) and cf.zone.plan eq \"ENT\""),
				Ref:              rulesetToDeployId,
				Position:         positionModel,
			}

			updateAccountRulesetOptions := &rulesetsv1.UpdateAccountRulesetOptions{
				RulesetID:   core.StringPtr(ruleset2UpdateId),
				Description: core.StringPtr("updating account ruleset"),
				Rules:       []rulesetsv1.RuleCreate{*ruleCreateModel},
			}

			rulesetResp, response, err := rulesetsService.UpdateAccountRuleset(updateAccountRulesetOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(rulesetResp).ToNot(BeNil())
		})

		It("Get Account Entrypoint Ruleset", func() {
			getAccountEntrypointRulesetOptions := &rulesetsv1.GetAccountEntrypointRulesetOptions{
				RulesetPhase: core.StringPtr("http_request_firewall_managed"),
			}

			rulesetResp, response, err := rulesetsService.GetAccountEntrypointRuleset(getAccountEntrypointRulesetOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(rulesetResp).ToNot(BeNil())
		})

		It("List/Get Account Entry Point Ruleset Version", func() {
			getAccountEntryPointRulesetVersionsOptions := &rulesetsv1.GetAccountEntryPointRulesetVersionsOptions{
				RulesetPhase: core.StringPtr("http_request_firewall_managed"),
			}

			listRulesetsResp, response, err := rulesetsService.GetAccountEntryPointRulesetVersions(getAccountEntryPointRulesetVersionsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(listRulesetsResp).ToNot(BeNil())

			version := *listRulesetsResp.Result[1].Version

			getAccountEntryPointRulesetVersionOptions := &rulesetsv1.GetAccountEntryPointRulesetVersionOptions{
				RulesetPhase:   core.StringPtr("http_request_firewall_managed"),
				RulesetVersion: core.StringPtr(version),
			}

			rulesetResp, response, err := rulesetsService.GetAccountEntryPointRulesetVersion(getAccountEntryPointRulesetVersionOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(rulesetResp).ToNot(BeNil())
		})

		It("List/Get/Delete Account Ruleset Version", func() {
			getAccountRulesetVersionsOptions := &rulesetsv1.GetAccountRulesetVersionsOptions{
				RulesetID: rulsetForTestingId,
			}

			listRulesetsResp, response, err := rulesetsService.GetAccountRulesetVersions(getAccountRulesetVersionsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(listRulesetsResp).ToNot(BeNil())

			version := *listRulesetsResp.Result[1].Version // gives the second last version

			getAccountRulesetVersionOptions := &rulesetsv1.GetAccountRulesetVersionOptions{
				RulesetID:      rulsetForTestingId,
				RulesetVersion: core.StringPtr(version),
			}

			rulesetResp, response, err := rulesetsService.GetAccountRulesetVersion(getAccountRulesetVersionOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(rulesetResp).ToNot(BeNil())

			deleteAccountRulesetVersionOptions := &rulesetsv1.DeleteAccountRulesetVersionOptions{
				RulesetID:      rulsetForTestingId,
				RulesetVersion: core.StringPtr(version),
			}

			response, err = rulesetsService.DeleteAccountRulesetVersion(deleteAccountRulesetVersionOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})

		It("Create Account Ruleset Rule", func() {

			getAccountEntrypointRulesetOptions := &rulesetsv1.GetAccountEntrypointRulesetOptions{
				RulesetPhase: core.StringPtr("http_request_firewall_managed"),
			}
			rulesetResp, _, _ := rulesetsService.GetAccountEntrypointRuleset(getAccountEntrypointRulesetOptions)
			ruleset2UpdateId := *rulesetResp.Result.ID

			rulesOverrideModel := &rulesetsv1.RulesOverride{
				ID:      rule2Id,
				Enabled: core.BoolPtr(true),
				Action:  core.StringPtr("log"),
			}

			categoriesOverrideModel := &rulesetsv1.CategoriesOverride{
				Category: core.StringPtr("wordpress"),
				Enabled:  core.BoolPtr(true),
				Action:   core.StringPtr("log"),
			}

			overridesModel := &rulesetsv1.Overrides{
				Enabled:    core.BoolPtr(false),
				Rules:      []rulesetsv1.RulesOverride{*rulesOverrideModel},
				Categories: []rulesetsv1.CategoriesOverride{*categoriesOverrideModel},
			}

			actionParametersModel := &rulesetsv1.ActionParameters{
				ID:        rulesetToDeployId,
				Overrides: overridesModel,
			}

			positionModel := &rulesetsv1.Position{
				Index: core.Int64Ptr(int64(1)),
			}

			createAccountRulesetRuleOptions := &rulesetsv1.CreateAccountRulesetRuleOptions{
				RulesetID:        core.StringPtr(ruleset2UpdateId),
				Action:           core.StringPtr("execute"),
				ActionParameters: actionParametersModel,
				Description:      core.StringPtr("adding a rule to execute managed rules"),
				Enabled:          core.BoolPtr(true),
				Expression:       core.StringPtr("(ip.src ne 1.1.1.1) and cf.zone.plan eq \"ENT\""),
				Ref:              core.StringPtr(ruleset2UpdateId),
				Position:         positionModel,
			}

			ruleResp, response, err := rulesetsService.CreateAccountRulesetRule(createAccountRulesetRuleOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(ruleResp).ToNot(BeNil())
		})

		It("Update/Delete Account Ruleset Rule", func() {

			getAccountEntrypointRulesetOptions := &rulesetsv1.GetAccountEntrypointRulesetOptions{
				RulesetPhase: core.StringPtr("http_request_firewall_managed"),
			}
			rulesetResp, _, _ := rulesetsService.GetAccountEntrypointRuleset(getAccountEntrypointRulesetOptions)
			ruleset2UpdateId := *rulesetResp.Result.ID
			rule2UpdateId := *rulesetResp.Result.Rules[0].ID

			rulesOverrideModel := &rulesetsv1.RulesOverride{
				ID:      rule2Id,
				Enabled: core.BoolPtr(true),
				Action:  core.StringPtr("block"),
			}

			categoriesOverrideModel := &rulesetsv1.CategoriesOverride{
				Category: core.StringPtr("wordpress"),
				Enabled:  core.BoolPtr(true),
				Action:   core.StringPtr("block"),
			}

			overridesModel := &rulesetsv1.Overrides{
				Action:     core.StringPtr("block"),
				Enabled:    core.BoolPtr(true),
				Rules:      []rulesetsv1.RulesOverride{*rulesOverrideModel},
				Categories: []rulesetsv1.CategoriesOverride{*categoriesOverrideModel},
			}

			actionParametersModel := &rulesetsv1.ActionParameters{
				ID:        rulesetToDeployId,
				Overrides: overridesModel,
			}

			updateAccountRulesetRuleOptions := &rulesetsv1.UpdateAccountRulesetRuleOptions{
				RulesetID:        core.StringPtr(ruleset2UpdateId),
				RuleID:           core.StringPtr(rule2UpdateId),
				Action:           core.StringPtr("execute"),
				ActionParameters: actionParametersModel,
				Description:      core.StringPtr("editting rule"),
				Enabled:          core.BoolPtr(true),
				Expression:       core.StringPtr("(ip.src ne 1.1.1.3) and cf.zone.plan eq \"ENT\""),
			}

			ruleResp, response, err := rulesetsService.UpdateAccountRulesetRule(updateAccountRulesetRuleOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(ruleResp).ToNot(BeNil())

			deleteAccountRulesetRuleOptions := &rulesetsv1.DeleteAccountRulesetRuleOptions{
				RulesetID: core.StringPtr(ruleset2UpdateId),
				RuleID:    core.StringPtr(rule2UpdateId),
			}

			ruleResp, response, err = rulesetsService.DeleteAccountRulesetRule(deleteAccountRulesetRuleOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(ruleResp).ToNot(BeNil())
		})

		It("Delete Account Ruleset", func() {
			deleteAccountRulesetOptions := &rulesetsv1.DeleteAccountRulesetOptions{
				RulesetID: rulsetForTestingId,
			}

			response, err := rulesetsService.DeleteAccountRuleset(deleteAccountRulesetOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})

		It("Get Account Ruleset Version By Tag", func() {
			getAccountRulesetVersionByTagOptions := &rulesetsv1.GetAccountRulesetVersionByTagOptions{
				RulesetID:      rulesetToDeployId,
				RulesetVersion: core.StringPtr("56"),
				RuleTag:        core.StringPtr("wordpress"),
			}

			rulesetResp, response, err := rulesetsService.GetAccountRulesetVersionByTag(getAccountRulesetVersionByTagOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(rulesetResp).ToNot(BeNil())
		})

	})

	//--------------ZONE RULESETS INTEGRATION TEST-----------------

	var zoneRulesetTodeploy *string
	var zoneRulesetTodeploy2 *string

	Describe("Zone rulesets", func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It("List/Get Zone Rulesets", func() {
			getZoneRulesetsOptions := &rulesetsv1.GetZoneRulesetsOptions{}

			listRulesetsResp, response, err := rulesetsService.GetZoneRulesets(getZoneRulesetsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(listRulesetsResp).ToNot(BeNil())

			for _, ruleset := range listRulesetsResp.Result {
				if strings.Contains(*ruleset.Name, "CIS Managed Ruleset") {
					zoneRulesetTodeploy = ruleset.ID
				}
				if strings.Contains(*ruleset.Name, "CIS Managed Free Ruleset") {
					zoneRulesetTodeploy2 = ruleset.ID
				}
			}

			getZoneRulesetOptions := &rulesetsv1.GetZoneRulesetOptions{
				RulesetID: zoneRulesetTodeploy,
			}

			rulesetResp, response, err := rulesetsService.GetZoneRuleset(getZoneRulesetOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(rulesetResp).ToNot(BeNil())

			rule1Id = rulesetResp.Result.Rules[0].ID
			rule2Id = rulesetResp.Result.Rules[2].ID

		})

		It("Update Zone Entrypoint Ruleset", func() {
			rulesOverrideModel := &rulesetsv1.RulesOverride{
				ID:      rule1Id,
				Enabled: core.BoolPtr(true),
				Action:  core.StringPtr("log"),
			}

			categoriesOverrideModel := &rulesetsv1.CategoriesOverride{
				Category: core.StringPtr("wordpress"),
				Enabled:  core.BoolPtr(true),
				Action:   core.StringPtr("log"),
			}

			overridesModel := &rulesetsv1.Overrides{
				Action:     core.StringPtr("log"),
				Enabled:    core.BoolPtr(true),
				Rules:      []rulesetsv1.RulesOverride{*rulesOverrideModel},
				Categories: []rulesetsv1.CategoriesOverride{*categoriesOverrideModel},
			}

			actionParametersModel := &rulesetsv1.ActionParameters{
				ID:        zoneRulesetTodeploy,
				Overrides: overridesModel,
			}

			positionModel := &rulesetsv1.Position{
				Index: core.Int64Ptr(int64(1)),
			}

			ruleCreateModel := &rulesetsv1.RuleCreate{
				Action:           core.StringPtr("execute"),
				ActionParameters: actionParametersModel,
				Description:      core.StringPtr("overriding entrypoint ruleset rule"),
				Enabled:          core.BoolPtr(true),
				Expression:       core.StringPtr("ip.src ne 1.1.1.1"),
				Position:         positionModel,
			}

			updateZoneEntrypointRulesetOptions := &rulesetsv1.UpdateZoneEntrypointRulesetOptions{
				RulesetPhase: core.StringPtr("http_request_firewall_managed"),
				Description:  core.StringPtr("updating entrypoint ruleset"),
				Rules:        []rulesetsv1.RuleCreate{*ruleCreateModel},
			}

			rulesetResp, response, err := rulesetsService.UpdateZoneEntrypointRuleset(updateZoneEntrypointRulesetOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(rulesetResp).ToNot(BeNil())

			rulsetForTestingId = rulesetResp.Result.ID
		})

		It(`UpdateZoneRuleset(updateZoneRulesetOptions *UpdateZoneRulesetOptions)`, func() {
			rulesOverrideModel := &rulesetsv1.RulesOverride{
				ID:      rule1Id,
				Enabled: core.BoolPtr(true),
				Action:  core.StringPtr("block"),
			}

			categoriesOverrideModel := &rulesetsv1.CategoriesOverride{
				Category: core.StringPtr("wordpress"),
				Enabled:  core.BoolPtr(true),
				Action:   core.StringPtr("block"),
			}

			overridesModel := &rulesetsv1.Overrides{
				Action:     core.StringPtr("block"),
				Enabled:    core.BoolPtr(true),
				Rules:      []rulesetsv1.RulesOverride{*rulesOverrideModel},
				Categories: []rulesetsv1.CategoriesOverride{*categoriesOverrideModel},
			}

			actionParametersModel := &rulesetsv1.ActionParameters{
				ID:        zoneRulesetTodeploy,
				Overrides: overridesModel,
			}

			positionModel := &rulesetsv1.Position{
				Index: core.Int64Ptr(int64(1)),
			}

			ruleCreateModel := &rulesetsv1.RuleCreate{
				Action:           core.StringPtr("execute"),
				ActionParameters: actionParametersModel,
				Description:      core.StringPtr("deploying a managed rule"),
				Enabled:          core.BoolPtr(true),
				Expression:       core.StringPtr("ip.src ne 1.1.1.2"),
				Ref:              zoneRulesetTodeploy,
				Position:         positionModel,
			}

			updateZoneRulesetOptions := &rulesetsv1.UpdateZoneRulesetOptions{
				RulesetID:   rulsetForTestingId,
				Description: core.StringPtr("Updating a zone ruleset"),
				Rules:       []rulesetsv1.RuleCreate{*ruleCreateModel},
			}

			rulesetResp, response, err := rulesetsService.UpdateZoneRuleset(updateZoneRulesetOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(rulesetResp).ToNot(BeNil())
		})

		It("Get Zone Entrypoint Ruleset", func() {
			getZoneEntrypointRulesetOptions := &rulesetsv1.GetZoneEntrypointRulesetOptions{
				RulesetPhase: core.StringPtr("http_request_firewall_managed"),
			}

			rulesetResp, response, err := rulesetsService.GetZoneEntrypointRuleset(getZoneEntrypointRulesetOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(rulesetResp).ToNot(BeNil())
		})

		It("List/Get Zone EntryPoint Ruleset Version", func() {
			getZoneEntryPointRulesetVersionsOptions := &rulesetsv1.GetZoneEntryPointRulesetVersionsOptions{
				RulesetPhase: core.StringPtr("http_request_firewall_managed"),
			}

			listRulesetsResp, response, err := rulesetsService.GetZoneEntryPointRulesetVersions(getZoneEntryPointRulesetVersionsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(listRulesetsResp).ToNot(BeNil())

			version := *listRulesetsResp.Result[1].Version

			getZoneEntryPointRulesetVersionOptions := &rulesetsv1.GetZoneEntryPointRulesetVersionOptions{
				RulesetPhase:   core.StringPtr("http_request_firewall_managed"),
				RulesetVersion: core.StringPtr(version),
			}

			rulesetResp, response, err := rulesetsService.GetZoneEntryPointRulesetVersion(getZoneEntryPointRulesetVersionOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(rulesetResp).ToNot(BeNil())
		})

		It("List/Get/Delete Zone Ruleset Version", func() {
			getZoneRulesetVersionsOptions := &rulesetsv1.GetZoneRulesetVersionsOptions{
				RulesetID: rulsetForTestingId,
			}

			listRulesetsResp, response, err := rulesetsService.GetZoneRulesetVersions(getZoneRulesetVersionsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(listRulesetsResp).ToNot(BeNil())

			version := *listRulesetsResp.Result[1].Version

			getZoneRulesetVersionOptions := &rulesetsv1.GetZoneRulesetVersionOptions{
				RulesetID:      rulsetForTestingId,
				RulesetVersion: core.StringPtr(version),
			}

			rulesetResp, response, err := rulesetsService.GetZoneRulesetVersion(getZoneRulesetVersionOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(rulesetResp).ToNot(BeNil())

			deleteZoneRulesetVersionOptions := &rulesetsv1.DeleteZoneRulesetVersionOptions{
				RulesetID:      rulsetForTestingId,
				RulesetVersion: core.StringPtr(version),
			}

			response, err = rulesetsService.DeleteZoneRulesetVersion(deleteZoneRulesetVersionOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})

		It("Create Zone Ruleset Rule", func() {
			rulesOverrideModel := &rulesetsv1.RulesOverride{
				ID:      rule2Id,
				Enabled: core.BoolPtr(true),
				Action:  core.StringPtr("log"),
			}

			categoriesOverrideModel := &rulesetsv1.CategoriesOverride{
				Category: core.StringPtr("wordpress"),
				Enabled:  core.BoolPtr(true),
				Action:   core.StringPtr("log"),
			}

			overridesModel := &rulesetsv1.Overrides{
				Action:     core.StringPtr("log"),
				Enabled:    core.BoolPtr(true),
				Rules:      []rulesetsv1.RulesOverride{*rulesOverrideModel},
				Categories: []rulesetsv1.CategoriesOverride{*categoriesOverrideModel},
			}

			actionParametersModel := &rulesetsv1.ActionParameters{
				ID:        zoneRulesetTodeploy2,
				Overrides: overridesModel,
			}

			positionModel := &rulesetsv1.Position{
				Index: core.Int64Ptr(int64(1)),
			}

			createZoneRulesetRuleOptions := &rulesetsv1.CreateZoneRulesetRuleOptions{
				RulesetID:        rulsetForTestingId,
				Action:           core.StringPtr("execute"),
				ActionParameters: actionParametersModel,
				Description:      core.StringPtr("deploying managed rule"),
				Enabled:          core.BoolPtr(true),
				Expression:       core.StringPtr("ip.src ne 1.1.1.3"),
				Ref:              rulsetForTestingId,
				Position:         positionModel,
			}

			ruleResp, response, err := rulesetsService.CreateZoneRulesetRule(createZoneRulesetRuleOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(ruleResp).ToNot(BeNil())
		})

		It("Update/Delete Zone Ruleset Rule", func() {

			getZoneEntrypointRulesetOptions := &rulesetsv1.GetZoneEntrypointRulesetOptions{
				RulesetPhase: core.StringPtr("http_request_firewall_managed"),
			}

			rulesetResp, _, _ := rulesetsService.GetZoneEntrypointRuleset(getZoneEntrypointRulesetOptions)
			rule2UpdateId := *rulesetResp.Result.Rules[0].ID

			rulesOverrideModel := &rulesetsv1.RulesOverride{
				ID:      rule2Id,
				Enabled: core.BoolPtr(true),
				Action:  core.StringPtr("block"),
			}

			categoriesOverrideModel := &rulesetsv1.CategoriesOverride{
				Category: core.StringPtr("wordpress"),
				Enabled:  core.BoolPtr(true),
				Action:   core.StringPtr("block"),
			}

			overridesModel := &rulesetsv1.Overrides{
				Action:     core.StringPtr("block"),
				Enabled:    core.BoolPtr(true),
				Rules:      []rulesetsv1.RulesOverride{*rulesOverrideModel},
				Categories: []rulesetsv1.CategoriesOverride{*categoriesOverrideModel},
			}

			actionParametersModel := &rulesetsv1.ActionParameters{
				ID:        zoneRulesetTodeploy2,
				Overrides: overridesModel,
			}

			updateZoneRulesetRuleOptions := &rulesetsv1.UpdateZoneRulesetRuleOptions{
				RulesetID:        rulsetForTestingId,
				RuleID:           core.StringPtr(rule2UpdateId),
				Action:           core.StringPtr("execute"),
				ActionParameters: actionParametersModel,
				Description:      core.StringPtr("updating the rule"),
				Enabled:          core.BoolPtr(true),
				Expression:       core.StringPtr("ip.src ne 1.1.1.4"),
				Ref:              rulsetForTestingId,
			}

			ruleResp, response, err := rulesetsService.UpdateZoneRulesetRule(updateZoneRulesetRuleOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(ruleResp).ToNot(BeNil())

			deleteZoneRulesetRuleOptions := &rulesetsv1.DeleteZoneRulesetRuleOptions{
				RulesetID: rulsetForTestingId,
				RuleID:    core.StringPtr(rule2UpdateId),
			}

			ruleResp, response, err = rulesetsService.DeleteZoneRulesetRule(deleteZoneRulesetRuleOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(ruleResp).ToNot(BeNil())
		})

		It("Delete Zone Ruleset", func() {
			deleteZoneRulesetOptions := &rulesetsv1.DeleteZoneRulesetOptions{
				RulesetID: rulsetForTestingId,
			}

			response, err := rulesetsService.DeleteZoneRuleset(deleteZoneRulesetOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})

	})

})
