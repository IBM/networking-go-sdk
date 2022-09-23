package mtlsv1_test

import (
	"fmt"
	"os"

	"github.com/IBM/go-sdk-core/v5/core"
	. "github.com/IBM/networking-go-sdk/mtlsv1"
	"github.com/joho/godotenv"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const configFile = "../cis.env"

var configLoaded bool = true

func shouldSkipTest() {
	Skip("skipping failing test")
	if !configLoaded {
		Skip("External configuration is not available, skipping...")
	}
}

var _ = Describe(`mtlsv1`, func() {
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
	URL := os.Getenv("URL")

	globalOptions := &MtlsV1Options{
		ServiceName:   DefaultServiceName,
		URL:           serviceURL,
		Authenticator: authenticator,
		Crn:           &crn,
	}

	service, serviceErr := NewMtlsV1(globalOptions)
	if serviceErr != nil {
		fmt.Println(serviceErr)
	}
	Describe(`mtlsv1_accessCertificates_test`, func() {
		Context(`mtlsv1_accessCertificates_test`, func() {
			defer GinkgoRecover()
			BeforeEach(func() {
				shouldSkipTest()

				// List access applications
				listAccOpt := service.NewListAccessApplicationsOptions(zoneId)
				listAccResult, listAccResp, listAccErr := service.ListAccessApplications(listAccOpt)
				Expect(listAccErr).To(BeNil())
				Expect(listAccResp).ToNot(BeNil())
				Expect(listAccResult).ToNot(BeNil())
				Expect(*listAccResult.Success).Should(BeTrue())
				// Delete an access applications
				for _, appId := range listAccResult.Result {
					// List access policy
					listOptPolicy := service.NewListAccessPoliciesOptions(zoneId, *appId.ID)
					listResultPolicy, listRespPolicy, listErrPolicy := service.ListAccessPolicies(listOptPolicy)
					Expect(listErrPolicy).To(BeNil())
					Expect(listRespPolicy).ToNot(BeNil())
					Expect(listResultPolicy).ToNot(BeNil())
					Expect(*listResultPolicy.Success).Should(BeTrue())
					// Delete access policy
					for _, policyId := range listResultPolicy.Result {
						delOptPolicy := service.NewDeleteAccessPolicyOptions(zoneId, *appId.ID, *policyId.ID)
						delResultPolicy, delRespPolicy, delErrPolicy := service.DeleteAccessPolicy(delOptPolicy)
						Expect(delErrPolicy).To(BeNil())
						Expect(delRespPolicy).ToNot(BeNil())
						Expect(delResultPolicy).ToNot(BeNil())
						Expect(*delResultPolicy.Success).Should(BeTrue())
					}
					delAccOpt := service.NewDeleteAccessApplicationOptions(zoneId, *appId.ID)
					delAccResult, delAccResp, delAccErr := service.DeleteAccessApplication(delAccOpt)
					Expect(delAccErr).To(BeNil())
					Expect(delAccResp).ToNot(BeNil())
					Expect(delAccResult).ToNot(BeNil())
					Expect(*delAccResult.Success).Should(BeTrue())
				}

				// List access certificate
				listOpt := service.NewListAccessCertificatesOptions(zoneId)
				listResult, listResp, listErr := service.ListAccessCertificates(listOpt)
				Expect(listErr).To(BeNil())
				Expect(listResp).ToNot(BeNil())
				Expect(listResult).ToNot(BeNil())
				Expect(*listResult.Success).Should(BeTrue())
				// Delete access certificate
				for _, certId := range listResult.Result {
					delOpt := service.NewDeleteAccessCertificateOptions(zoneId, *certId.ID)
					delResult, delResp, delErr := service.DeleteAccessCertificate(delOpt)
					Expect(delErr).To(BeNil())
					Expect(delResp).ToNot(BeNil())
					Expect(delResult).ToNot(BeNil())
					Expect(*delResult.Success).Should(BeTrue())
				}
			})

			AfterEach(func() {
				shouldSkipTest()
				// List access applications
				listAccOpt := service.NewListAccessApplicationsOptions(zoneId)
				listAccResult, listAccResp, listAccErr := service.ListAccessApplications(listAccOpt)
				Expect(listAccErr).To(BeNil())
				Expect(listAccResp).ToNot(BeNil())
				Expect(listAccResult).ToNot(BeNil())
				Expect(*listAccResult.Success).Should(BeTrue())
				// Delete an access applications
				for _, appId := range listAccResult.Result {
					// List access policy
					listOptPolicy := service.NewListAccessPoliciesOptions(zoneId, *appId.ID)
					listResultPolicy, listRespPolicy, listErrPolicy := service.ListAccessPolicies(listOptPolicy)
					Expect(listErrPolicy).To(BeNil())
					Expect(listRespPolicy).ToNot(BeNil())
					Expect(listResultPolicy).ToNot(BeNil())
					Expect(*listResultPolicy.Success).Should(BeTrue())
					// Delete access policy
					for _, policyId := range listResultPolicy.Result {
						delOptPolicy := service.NewDeleteAccessPolicyOptions(zoneId, *appId.ID, *policyId.ID)
						delResultPolicy, delRespPolicy, delErrPolicy := service.DeleteAccessPolicy(delOptPolicy)
						Expect(delErrPolicy).To(BeNil())
						Expect(delRespPolicy).ToNot(BeNil())
						Expect(delResultPolicy).ToNot(BeNil())
						Expect(*delResultPolicy.Success).Should(BeTrue())
					}
					delAccOpt := service.NewDeleteAccessApplicationOptions(zoneId, *appId.ID)
					delAccResult, delAccResp, delAccErr := service.DeleteAccessApplication(delAccOpt)
					Expect(delAccErr).To(BeNil())
					Expect(delAccResp).ToNot(BeNil())
					Expect(delAccResult).ToNot(BeNil())
					Expect(*delAccResult.Success).Should(BeTrue())
				}

				// List  access certificate
				listOpt := service.NewListAccessCertificatesOptions(zoneId)
				listResult, listResp, listErr := service.ListAccessCertificates(listOpt)
				Expect(listErr).To(BeNil())
				Expect(listResp).ToNot(BeNil())
				Expect(listResult).ToNot(BeNil())
				Expect(*listResult.Success).Should(BeTrue())
				// Delete an access certificate
				for _, certId := range listResult.Result {
					delOpt := service.NewDeleteAccessCertificateOptions(zoneId, *certId.ID)
					delResult, delResp, delErr := service.DeleteAccessCertificate(delOpt)
					Expect(delErr).To(BeNil())
					Expect(delResp).ToNot(BeNil())
					Expect(delResult).ToNot(BeNil())
					Expect(*delResult.Success).Should(BeTrue())
				}

			})
			It(`Create| Get | Update | List | Delete Access certificate, Access applications, Access policy, Access certificates settings `, func() {
				// Create an access certificate
				name := "test-cert"
				certificate := "-----BEGIN CERTIFICATE-----\nMIIEFzCCAv+gAwIBAgIJAMhhsP5Ubtu2MA0GCSqGSIb3DQEBCwUAMIGhMQswCQYD\nVQQGEwJpbjESMBAGA1UECAwJa2FybmF0YWthMRIwEAYDVQQHDAliYW5nYWxvcmUx\nDDAKBgNVBAoMA2libTEMMAoGA1UECwwDY2lzMSowKAYDVQQDDCFtdGxzNy5hdXN0\nZXN0LTEwLmNpc3Rlc3QtbG9hZC5jb20xIjAgBgkqhkiG9w0BCQEWE2RhcnVueWEu\nZC5jQGlibS5jb20wHhcNMjIwNDIyMTEwMzU3WhcNMzIwNDE5MTEwMzU3WjCBoTEL\nMAkGA1UEBhMCaW4xEjAQBgNVBAgMCWthcm5hdGFrYTESMBAGA1UEBwwJYmFuZ2Fs\nb3JlMQwwCgYDVQQKDANpYm0xDDAKBgNVBAsMA2NpczEqMCgGA1UEAwwhbXRsczcu\nYXVzdGVzdC0xMC5jaXN0ZXN0LWxvYWQuY29tMSIwIAYJKoZIhvcNAQkBFhNkYXJ1\nbnlhLmQuY0BpYm0uY29tMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA\n3tjgNpucsvwNFPNWl1DXkWGFLzvdMKDdk3PTAJ3AAYFG4jLVDtZurf3qCLZ8fcz+\nnukYdDKhRZYSP9QvGwDTS4mHOTV/6FAYsb7qfke+V8+v0okmCca07KgTUKFR5F9e\nw1NPYW9yRjoVpy/Kgs983WigDBRQeo50wcLYG7APml0ceqsBKZaXOiTVrf2xDSvd\nNn6Qchgd5dmxiP+drypt7BGIf9j8QlN5HvEETfUQQybwJfq9G6KhNKIKcw+IKGIy\nbI03RmItC+eVhwja/t1UldlXt/L3JduwEkq9QNQe080toAZyaQ/9Vymk80DTrffN\njb1YG224XLlflSSdzbUC0QIDAQABo1AwTjAdBgNVHQ4EFgQUs5QUMLmjPfNutr8U\n2zcjT/yH1pYwHwYDVR0jBBgwFoAUs5QUMLmjPfNutr8U2zcjT/yH1pYwDAYDVR0T\nBAUwAwEB/zANBgkqhkiG9w0BAQsFAAOCAQEAPCqm4rXm0ptf0iSp+u4X60A3U3ON\ntSpKq5BU1KGF0i5/ZB1ia1we2ORdOzeoNIhoffmRCg/a//Ba5fLRhktzXMcT/zwC\nDVxH9OAtFoj6/rfEko6s+NP/WtWMd7YF1w4wVvK189YWSUDKbE4MijeDLvEfBi3T\nStNu14p4gN8hkSLX/3Rn9ZmI2wDIpqsYRF5KPfvNZ0iIpvJoBWjS6bbVYGd3yNs+\nrXez+Q36oEFfMcM35EEt3qo2EGu4mljqZxhIae5Hy4sKe4c6s0AfpYA4wTQ97cAg\nQ0Sdw3p+PIqPMOcY1sjRLbvPDHGbzc60LvKhHgt/7Cc5ntvxIjJ9ZUt5Ng==\n-----END CERTIFICATE-----\n"
				options := service.NewCreateAccessCertificateOptions(zoneId)
				options.SetName(name)
				options.SetCertificate(certificate)
				options.SetAssociatedHostnames([]string{URL})
				result, response, operationErr := service.CreateAccessCertificate(options)
				certId := *result.Result.ID
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Get access certificate
				getOptions := service.NewGetAccessCertificateOptions(zoneId, certId)
				getResult, getResp, getErr := service.GetAccessCertificate(getOptions)
				Expect(getErr).To(BeNil())
				Expect(getResp).ToNot(BeNil())
				Expect(getResult).ToNot(BeNil())

				// update access certificate
				updateOption := service.NewUpdateAccessCertificateOptions(zoneId, certId)
				updateOption.SetName("test-newcert")
				updateOption.SetAssociatedHostnames([]string{})
				updateResult, updateResp, updateErr := service.UpdateAccessCertificate(updateOption)
				Expect(updateErr).To(BeNil())
				Expect(updateResp).ToNot(BeNil())
				Expect(updateResult).ToNot(BeNil())
				Expect(*updateResult.Success).Should(BeTrue())

				// Create an access applications
				OptionsApp := service.NewCreateAccessApplicationOptions(zoneId)
				OptionsApp.SetName("mtls-test-app")
				OptionsApp.SetDomain(URL)
				OptionsApp.SetSessionDuration("24h")
				resultApp, responseApp, operationErrApp := service.CreateAccessApplication(OptionsApp)
				appId := *resultApp.Result.ID
				Expect(operationErrApp).To(BeNil())
				Expect(responseApp).ToNot(BeNil())
				Expect(resultApp).ToNot(BeNil())

				// Get an access applications
				getAppOptions := service.NewGetAccessApplicationOptions(zoneId, appId)
				getAppResult, getAppResp, getAppErr := service.GetAccessApplication(getAppOptions)
				Expect(getAppErr).To(BeNil())
				Expect(getAppResp).ToNot(BeNil())
				Expect(getAppResult).ToNot(BeNil())

				// update an access applications
				updateOptionApp := service.NewUpdateAccessApplicationOptions(zoneId, appId)
				updateOptionApp.SetName("mtls-test-app1")
				updateOptionApp.SetDomain(URL)
				updateOptionApp.SetSessionDuration("24h")
				updateResultApp, updateRespApp, updateErrApp := service.UpdateAccessApplication(updateOptionApp)
				Expect(updateErrApp).To(BeNil())
				Expect(updateRespApp).ToNot(BeNil())
				Expect(updateResultApp).ToNot(BeNil())
				Expect(*updateResultApp.Success).Should(BeTrue())

				// Create an  access policy
				policyRuleModel := &PolicyRulePolicyCertRule{
					Certificate: map[string]interface{}{"certifcate": "CA root certificate"},
				}
				policyCnModel := &PolicyCnRuleCommonName{
					CommonName: core.StringPtr("Access Testing CA"),
				}
				policyModel := &PolicyRulePolicyCnRule{
					CommonName: policyCnModel,
				}
				optionsPolicy := service.NewCreateAccessPolicyOptions(zoneId, appId)
				optionsPolicy.SetName("mtls-policy")
				optionsPolicy.SetDecision("non_identity")
				optionsPolicy.SetInclude([]PolicyRuleIntf{policyModel, policyRuleModel})
				resultPolicy, responsePolicy, operationErrPolicy := service.CreateAccessPolicy(optionsPolicy)
				policyId := *resultPolicy.Result.ID
				Expect(operationErrPolicy).To(BeNil())
				Expect(responsePolicy).ToNot(BeNil())
				Expect(resultPolicy).ToNot(BeNil())

				// update an  access policy
				policyUpadteCnModel := &PolicyCnRuleCommonName{
					CommonName: core.StringPtr("Access Testing CA"),
				}
				policyUpadteModel := &PolicyRulePolicyCnRule{
					CommonName: policyUpadteCnModel,
				}
				policyUpadteRuleModel := &PolicyRulePolicyCertRule{
					Certificate: map[string]interface{}{"certifcate": "CA root certificate"},
				}
				optionPolicyUpdate := service.NewUpdateAccessPolicyOptions(zoneId, appId, policyId)
				optionPolicyUpdate.SetName("mtls-test-policy")
				optionPolicyUpdate.SetDecision("non_identity")
				optionPolicyUpdate.SetInclude([]PolicyRuleIntf{policyUpadteModel, policyUpadteRuleModel})
				resultUpadtePolicy, responseUpadtePolicy, operationUpadteErrPolicy := service.UpdateAccessPolicy(optionPolicyUpdate)
				Expect(operationUpadteErrPolicy).To(BeNil())
				Expect(responseUpadtePolicy).ToNot(BeNil())
				Expect(resultUpadtePolicy).ToNot(BeNil())

				// Get an access policy
				getPolicyOptions := service.NewGetAccessPolicyOptions(zoneId, appId, policyId)
				getPolicyResult, getPolicyResp, getPolicyErr := service.GetAccessPolicy(getPolicyOptions)
				Expect(getPolicyErr).To(BeNil())
				Expect(getPolicyResp).ToNot(BeNil())
				Expect(getPolicyResult).ToNot(BeNil())

				// Update an access certificates settings
				accessCertModel := &AccessCertSettingsInputArray{
					Hostname:                    core.StringPtr(URL),
					ClientCertificateForwarding: core.BoolPtr(false),
				}
				updateSettingOption := service.NewUpdateAccessCertSettingsOptions(zoneId)
				updateSettingOption.SetSettings([]AccessCertSettingsInputArray{*accessCertModel})
				updateSettingResult, updateSettingResp, updateSettingErr := service.UpdateAccessCertSettings(updateSettingOption)
				Expect(updateSettingErr).To(BeNil())
				Expect(updateSettingResp).ToNot(BeNil())
				Expect(updateSettingResult).ToNot(BeNil())
				Expect(*updateSettingResult.Success).Should(BeTrue())

				// // Get an access certificates settings
				getSettingOption := service.NewGetAccessCertSettingsOptions(zoneId)
				getSettingResult, getSettingResp, getSettingErr := service.GetAccessCertSettings(getSettingOption)
				Expect(getSettingErr).To(BeNil())
				Expect(getSettingResp).ToNot(BeNil())
				Expect(getSettingResult).ToNot(BeNil())
				Expect(*getSettingResult.Success).Should(BeTrue())

			})
		})
	})
})
