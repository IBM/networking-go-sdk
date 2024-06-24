/*
 * (C) Copyright IBM Corp. 2024.
 */

package sslcertificateapiv1_test

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	. "github.com/IBM/networking-go-sdk/sslcertificateapiv1"
	"github.com/joho/godotenv"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const configFile = "../cis.env"

var configLoaded bool = true
var authenticationSucceeded bool = true

func shouldSkipTest() {
	if !configLoaded {
		Skip("External configuration is not available, skipping...")
	}

	if !authenticationSucceeded {
		Skip("Authentication failed. Check external configuration...")
	}
}

var _ = Describe(`sslcertificateapiv1`, func() {

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
	authErr := authenticator.Authenticate(&http.Request{
		Header: http.Header{},
	})
	if authErr != nil {
		authenticationSucceeded = false
		fmt.Println("Authentication error during setup: ", authErr)
	}
	serviceURL := os.Getenv("API_ENDPOINT")
	crn := os.Getenv("CRN")
	zone_id := os.Getenv("ZONE_ID")
	url := os.Getenv("URL")
	certificate := os.Getenv("CERTIFICATE")
	updateCertificate := os.Getenv("UPDATE_CERTIFICATE")
	privateKey := os.Getenv("PRIVATE_KEY")
	updatePrivateKey := os.Getenv("UPDATE_PRIVATE_KEY")
	globalOptions := &SslCertificateApiV1Options{
		ServiceName:    "cis_services",
		URL:            serviceURL,
		Authenticator:  authenticator,
		Crn:            &crn,
		ZoneIdentifier: &zone_id,
	}

	service, serviceErr := NewSslCertificateApiV1(globalOptions)
	if serviceErr != nil {
		fmt.Println(serviceErr)
	}
	Describe(`order/view/delete ssl certificate packs`, func() {
		Context(`order/view/delete ssl certificate packs`, func() {
			BeforeEach(func() {
				shouldSkipTest()
				// list all certificates
				listOpt := service.NewListCertificatesOptions()
				listOpt.SetXCorrelationID("12345")
				listResult, listResp, listErr := service.ListCertificates(listOpt)
				Expect(listErr).To(BeNil())
				Expect(listResp).ToNot(BeNil())
				Expect(listResult).ToNot(BeNil())
				result := listResult.Result

				for _, cert := range result {
					if *cert.Type == OrderCertificateOptions_Type_Dedicated {
						delOpt := service.NewDeleteCertificateOptions(*cert.ID)
						_, _ = service.DeleteCertificate(delOpt)

					} else {
						delOpt := service.NewDeleteCustomCertificateOptions(*cert.ID)
						_, _ = service.DeleteCustomCertificate(delOpt)
					}

					log.Println("[BeforeEach] Sleeping for 10 second to check the certificate is deleted.")
					time.Sleep(10 * time.Second)

					for i := 0; i < 5; i++ {
						slept := false

						listOpt := service.NewListCertificatesOptions()
						listResult, _, listErr := service.ListCertificates(listOpt)
						if listErr == nil {
							for _, certCheck := range listResult.Result {
								if *certCheck.ID == *cert.ID {
									log.Println("sleeping for 1 minutes.")
									time.Sleep(1 * time.Minute)
									slept = true
									break
								}
							}
						}
						if slept == false {
							break
						}
					}
				}
			})
			AfterEach(func() {
				shouldSkipTest()
				// list all certificates
				getOpt := service.NewListCertificatesOptions()
				getOpt.SetXCorrelationID("12345")
				listResult, listResp, listErr := service.ListCertificates(getOpt)
				Expect(listErr).To(BeNil())
				Expect(listResp).ToNot(BeNil())
				Expect(listResult).ToNot(BeNil())
				result := listResult.Result

				for _, cert := range result {
					if *cert.Type == OrderCertificateOptions_Type_Dedicated {
						delOpt := service.NewDeleteCertificateOptions(*cert.ID)
						_, _ = service.DeleteCertificate(delOpt)
					} else {
						delOpt := service.NewDeleteCustomCertificateOptions(*cert.ID)
						_, _ = service.DeleteCustomCertificate(delOpt)
					}
					log.Println("[AfterEach] Sleeping for 10 second to check the certificate is deleted.")
					time.Sleep(10 * time.Second)
					for i := 0; i < 5; i++ {
						slept := false

						listOpt := service.NewListCertificatesOptions()
						listResult, _, listErr := service.ListCertificates(listOpt)
						if listErr == nil {
							for _, certCheck := range listResult.Result {
								if *certCheck.ID == *cert.ID {
									log.Println("sleeping for 1 minutes.")
									time.Sleep(1 * time.Minute)
									slept = true
									break
								}
							}
						}
						if slept == false {
							break
						}
					}
				}
			})
			It(`order/view/delete ssl certificate packs`, func() {
				Skip("SKipping this test as this method is deprecated")
				shouldSkipTest()
				// order certificate packs
				orderOpt := service.NewOrderCertificateOptions()
				orderOpt.SetHosts([]string{url})
				orderOpt.SetXCorrelationID("12345")
				orderOpt.SetType(OrderCertificateOptions_Type_Dedicated)

				orderResult, orderResp, orderErr := service.OrderCertificate(orderOpt)
				Expect(orderErr).To(BeNil())
				Expect(orderResp).ToNot(BeNil())
				Expect(orderResult).ToNot(BeNil())
				Expect(orderResult.Result.ID).ToNot(BeNil())
				Expect(orderResult.Result.Status).ToNot(BeNil())
				Expect(*orderResult.Result.Type).Should(BeEquivalentTo("advanced"))

				// list all certificates
				listOpt := service.NewListCertificatesOptions()
				listOpt.SetXCorrelationID("12345")
				listResult, listResp, listErr := service.ListCertificates(listOpt)
				Expect(listErr).To(BeNil())
				Expect(listResp).ToNot(BeNil())
				Expect(listResult).ToNot(BeNil())
				for _, cert := range listResult.Result {
					Expect(cert.PrimaryCertificate).ToNot(BeNil())
					Expect(cert.Hosts[0]).Should(BeEquivalentTo(url))
					Expect(cert.Certificates[0].Hosts[0]).Should(BeEquivalentTo(url))
					Expect(cert.Certificates[0].ID).ToNot(BeNil())
				}

				delOpt := service.NewDeleteCertificateOptions(*orderResult.Result.ID)
				delResp, delErr := service.DeleteCertificate(delOpt)
				Expect(delErr).To(BeNil())
				Expect(delResp).ToNot(BeNil())
			})
			It(`upload/view/delete ssl custom certificates`, func() {
				Skip("No need to run this test case...")
				shouldSkipTest()
				// upload certificate packs
				geoOpt, geoErr := service.NewCustomCertReqGeoRestrictions("us")
				Expect(geoErr).To(BeNil())
				uploadOpt := service.NewUploadCustomCertificateOptions()
				uploadOpt.SetCertificate(certificate)
				uploadOpt.SetPrivateKey(privateKey)
				uploadOpt.SetGeoRestrictions(geoOpt)
				uploadOpt.SetBundleMethod(UpdateCustomCertificateOptions_BundleMethod_Optimal)

				uploadResult, uploadResp, uploadErr := service.UploadCustomCertificate(uploadOpt)
				Expect(uploadErr).To(BeNil())
				Expect(uploadResp).ToNot(BeNil())
				Expect(uploadResult).ToNot(BeNil())
				Expect(*uploadResult.Success).Should(BeTrue())
				Expect(uploadResult.Result.ID).ToNot(BeNil())
				Expect(uploadResult.Result.Hosts[0]).Should(BeEquivalentTo(fmt.Sprintf("beta.%s", url)))
				Expect(uploadResult.Result.Issuer).ToNot(BeNil())
				Expect(uploadResult.Result.Priority).ToNot(BeNil())
				Expect(uploadResult.Result.Signature).ToNot(BeNil())
				Expect(uploadResult.Result.Status).ToNot(BeNil())
				Expect(*uploadResult.Result.BundleMethod).Should(BeEquivalentTo(UpdateCustomCertificateOptions_BundleMethod_Optimal))
				Expect(*uploadResult.Result.ZoneID).Should(BeEquivalentTo(zone_id))

				// get custom certificate
				getOpt := service.NewGetCustomCertificateOptions(*uploadResult.Result.ID)
				getResult, getResp, getErr := service.GetCustomCertificate(getOpt)
				Expect(getErr).To(BeNil())
				Expect(getResp).ToNot(BeNil())
				Expect(getResult).ToNot(BeNil())
				Expect(*getResult.Success).Should(BeTrue())
				Expect(getResult.Result.Hosts[0]).Should(BeEquivalentTo(fmt.Sprintf("beta.%s", url)))
				Expect(getResult.Result.Issuer).ToNot(BeNil())
				Expect(getResult.Result.Priority).ToNot(BeNil())
				Expect(getResult.Result.Signature).ToNot(BeNil())
				Expect(getResult.Result.Status).ToNot(BeNil())
				Expect(*getResult.Result.BundleMethod).Should(BeEquivalentTo(UpdateCustomCertificateOptions_BundleMethod_Optimal))
				Expect(*getResult.Result.ZoneID).Should(BeEquivalentTo(zone_id))

				// update custom certificate
				updateGeoOpt, updateErr := service.NewCustomCertReqGeoRestrictions("eu")
				Expect(updateErr).To(BeNil())
				updateOpt := service.NewUpdateCustomCertificateOptions(*uploadResult.Result.ID)
				updateOpt.SetBundleMethod(UpdateCustomCertificateOptions_BundleMethod_Ubiquitous)
				updateOpt.SetCertificate(updateCertificate)
				updateOpt.SetGeoRestrictions(updateGeoOpt)
				updateOpt.SetPrivateKey(updatePrivateKey)

				updateResult, updateResp, updateErr := service.UpdateCustomCertificate(updateOpt)
				Expect(updateErr).To(BeNil())
				Expect(updateResp).ToNot(BeNil())
				Expect(updateResult).ToNot(BeNil())
				Expect(*updateResult.Success).Should(BeTrue())
				Expect(updateResult.Result.Hosts[0]).Should(BeEquivalentTo(fmt.Sprintf("ibm.%s", url)))
				Expect(updateResult.Result.Issuer).ToNot(BeNil())
				Expect(updateResult.Result.Priority).ToNot(BeNil())
				Expect(updateResult.Result.Signature).ToNot(BeNil())
				Expect(updateResult.Result.Status).ToNot(BeNil())
				Expect(*updateResult.Result.BundleMethod).Should(BeEquivalentTo(UpdateCustomCertificateOptions_BundleMethod_Ubiquitous))
				Expect(*updateResult.Result.ZoneID).Should(BeEquivalentTo(zone_id))

				// list all custom certificates
				listOpt := service.NewListCustomCertificatesOptions()
				listResult, listResp, listErr := service.ListCustomCertificates(listOpt)
				Expect(listErr).To(BeNil())
				Expect(listResp).ToNot(BeNil())
				Expect(listResult).ToNot(BeNil())
				Expect(*listResult.Success).Should(BeTrue())

				Expect(listResult.Result[0].Hosts[0]).Should(BeEquivalentTo(fmt.Sprintf("ibm.%s", url)))
				Expect(listResult.Result[0].Issuer).ToNot(BeNil())
				Expect(listResult.Result[0].Priority).ToNot(BeNil())
				Expect(listResult.Result[0].Signature).ToNot(BeNil())
				Expect(listResult.Result[0].Status).ToNot(BeNil())
				Expect(*listResult.Result[0].BundleMethod).Should(BeEquivalentTo(UpdateCustomCertificateOptions_BundleMethod_Ubiquitous))
				Expect(*listResult.Result[0].ZoneID).Should(BeEquivalentTo(zone_id))

				// delete all custom certificates
				for _, cert := range listResult.Result {
					delOpt := service.NewDeleteCustomCertificateOptions(*cert.ID)
					delResp, delErr := service.DeleteCustomCertificate(delOpt)
					Expect(delErr).To(BeNil())
					Expect(delResp).ToNot(BeNil())
				}
			})
			It(`change/get/delete ssl universal certificate setting`, func() {
				Skip("No need to run this test case...")
				shouldSkipTest()
				// upload custom certificate
				geoOpt, geoErr := service.NewCustomCertReqGeoRestrictions("us")
				Expect(geoErr).To(BeNil())
				uploadOpt := service.NewUploadCustomCertificateOptions()
				uploadOpt.SetCertificate(certificate)
				uploadOpt.SetPrivateKey(privateKey)
				uploadOpt.SetGeoRestrictions(geoOpt)
				uploadOpt.SetBundleMethod(UpdateCustomCertificateOptions_BundleMethod_Optimal)

				uploadResult, uploadResp, uploadErr := service.UploadCustomCertificate(uploadOpt)
				Expect(uploadErr).To(BeNil())
				Expect(uploadResp).ToNot(BeNil())
				Expect(uploadResult).ToNot(BeNil())
				Expect(*uploadResult.Success).Should(BeTrue())

				priorityItem, _ := service.NewCertPriorityReqCertificatesItem(*uploadResult.Result.ID, 40)
				priorityOpt := service.NewChangeCertificatePriorityOptions()
				priorityOpt.SetCertificates([]CertPriorityReqCertificatesItem{*priorityItem})

				changePriorityResp, changePriorityErr := service.ChangeCertificatePriority(priorityOpt)
				Expect(changePriorityErr).To(BeNil())
				Expect(changePriorityResp).ToNot(BeNil())

				// get universal certificate setting
				getOpt := service.NewGetUniversalCertificateSettingOptions()
				getResult, getResp, getErr := service.GetUniversalCertificateSetting(getOpt)
				Expect(getErr).To(BeNil())
				Expect(getResp).ToNot(BeNil())
				Expect(getResult).ToNot(BeNil())
				Expect(*getResult.Success).Should(BeTrue())

				// update universal certificate setting
				updateOpt := service.NewChangeUniversalCertificateSettingOptions()
				updateOpt.SetEnabled(false)

				updateResp, updateErr := service.ChangeUniversalCertificateSetting(updateOpt)
				Expect(updateErr).To(BeNil())
				Expect(updateResp).ToNot(BeNil())

				// list all custom certificates
				listOpt := service.NewListCustomCertificatesOptions()
				listResult, listResp, listErr := service.ListCustomCertificates(listOpt)
				Expect(listErr).To(BeNil())
				Expect(listResp).ToNot(BeNil())
				Expect(listResult).ToNot(BeNil())
				Expect(*listResult.Success).Should(BeTrue())

				// delete all custom certificates
				for _, cert := range listResult.Result {
					delOpt := service.NewDeleteCustomCertificateOptions(*cert.ID)
					delResp, delErr := service.DeleteCustomCertificate(delOpt)
					Expect(delErr).To(BeNil())
					Expect(delResp).ToNot(BeNil())
				}
			})
		})
	})
	Describe(`SSL Certificate Settings`, func() {
		It(`change/get ssl certificate setting`, func() {
			shouldSkipTest()
			values := []string{
				ChangeSslSettingOptions_Value_Full,
				ChangeSslSettingOptions_Value_Strict,
				ChangeSslSettingOptions_Value_Off,
				ChangeSslSettingOptions_Value_Flexible,
			}
			changeOpt := service.NewChangeSslSettingOptions()
			for _, val := range values {

				// change ssl certificate setting
				changeOpt.SetValue(val)
				changeResult, changeResp, changeErr := service.ChangeSslSetting(changeOpt)
				Expect(changeErr).To(BeNil())
				Expect(changeResp).ToNot(BeNil())
				Expect(changeResult).ToNot(BeNil())
				Expect(*changeResult.Success).Should(BeTrue())

				getOpt := service.NewGetSslSettingOptions()
				getResult, getResp, getErr := service.GetSslSetting(getOpt)
				Expect(getErr).To(BeNil())
				Expect(getResp).ToNot(BeNil())
				Expect(getResult).ToNot(BeNil())
				Expect(*getResult.Success).Should(BeTrue())
			}
		})
		It(`change/get TLS version 1.2 setting`, func() {
			shouldSkipTest()
			// get TLS version 1.2 setting
			getOpt := service.NewGetTls12SettingOptions()
			getResult, getResp, getErr := service.GetTls12Setting(getOpt)
			Expect(getErr).To(BeNil())
			Expect(getResp).ToNot(BeNil())
			Expect(getResult).ToNot(BeNil())
			Expect(*getResult.Success).Should(BeTrue())

			// change TLS version 1.2 setting
			changeOpt := service.NewChangeTls12SettingOptions()
			changeOpt.SetValue(ChangeTls13SettingOptions_Value_Off)
			changeResult, changeResp, changeErr := service.ChangeTls12Setting(changeOpt)
			Expect(changeErr).To(BeNil())
			Expect(changeResp).ToNot(BeNil())
			Expect(changeResult).ToNot(BeNil())
			Expect(*changeResult.Success).Should(BeTrue())
		})
		It(`change/get TLS version 1.3 setting`, func() {
			shouldSkipTest()
			// get TLS version 1.3 setting
			getOpt := service.NewGetTls13SettingOptions()
			getResult, getResp, getErr := service.GetTls13Setting(getOpt)
			Expect(getErr).To(BeNil())
			Expect(getResp).ToNot(BeNil())
			Expect(getResult).ToNot(BeNil())
			Expect(*getResult.Success).Should(BeTrue())

			// change TLS version 1.3 setting
			changeOpt := service.NewChangeTls13SettingOptions()
			changeOpt.SetValue(ChangeTls13SettingOptions_Value_Off)
			if *getResult.Result.Value == ChangeTls13SettingOptions_Value_Off {
				changeOpt.SetValue(ChangeTls13SettingOptions_Value_On)
			}
			changeResult, changeResp, changeErr := service.ChangeTls13Setting(changeOpt)
			Expect(changeErr).To(BeNil())
			Expect(changeResp).ToNot(BeNil())
			Expect(changeResult).ToNot(BeNil())
			Expect(*changeResult.Success).Should(BeTrue())

			changeOpt.SetValue(ChangeTls13SettingOptions_Value_Off)
			changeResult, changeResp, changeErr = service.ChangeTls13Setting(changeOpt)
			Expect(changeErr).To(BeNil())
			Expect(changeResp).ToNot(BeNil())
			Expect(changeResult).ToNot(BeNil())
			Expect(*changeResult.Success).Should(BeTrue())
		})
	})

	Describe("Order Advanced Certificate", func() {
		It("create/delete advance certificate pack", func() {
			shouldSkipTest()
			// order advanced certificate pack
			orderOpt := service.NewOrderAdvancedCertificateOptions()
			orderOpt.SetCertificateAuthority(OrderAdvancedCertificateOptions_CertificateAuthority_LetsEncrypt)
			orderOpt.SetCloudflareBranding(false)
			orderOpt.SetHosts([]string{url})
			orderOpt.SetType(OrderAdvancedCertificateOptions_Type_Advanced)
			orderOpt.SetValidationMethod(OrderAdvancedCertificateOptions_ValidationMethod_Txt)
			orderOpt.SetValidityDays(90)

			certResult, certResp, certErr := service.OrderAdvancedCertificate(orderOpt)

			Expect(certErr).To(BeNil())
			Expect(certResp).ToNot(BeNil())
			Expect(certResult).ToNot(BeNil())
			Expect(certResult.Result.ID).ToNot(BeNil())
			Expect(certResult.Result.Status).ToNot(BeNil())
			Expect(*certResult.Result.Type).Should(BeEquivalentTo(OrderAdvancedCertificateOptions_Type_Advanced))
			advanceCertId := *certResult.Result.ID

			// restart validation for advaced certificate under validation_timed_out state
			// skipping as there is no way to test now
			// validateOpt := service.NewPatchCertificateOptions(advanceCertId)
			// certResult, certResp, certErr = service.PatchCertificate(validateOpt)
			// Expect(certErr).To(BeNil())
			// Expect(certResp).ToNot(BeNil())
			// Expect(certResult).ToNot(BeNil())

			// delete advance certificate pack
			delOpt := service.NewDeleteCertificateV2Options(advanceCertId)
			delResp, delErr := service.DeleteCertificateV2(delOpt)
			Expect(delErr).To(BeNil())
			Expect(delResp).ToNot(BeNil())

		})

		It("Get SSL verification info of a zone", func() {
			shouldSkipTest()
			Skip("Can be skipped..")
			verificationOpt := service.NewGetSslVerificationOptions()
			result, resp, err := service.GetSslVerification(verificationOpt)

			Expect(err).To(BeNil())
			Expect(resp).ToNot(BeNil())
			Expect(result).ToNot(BeNil())

		})
	})

	Describe("Origin Certificate", func() {
		It("List origin certificate", func() {
			shouldSkipTest()
			certOpt := service.NewListOriginCertificatesOptions(crn, zone_id)
			certResult, certResp, certErr := service.ListOriginCertificates(certOpt)

			Expect(certErr).To(BeNil())
			Expect(certResp).ToNot(BeNil())
			Expect(certResult).ToNot(BeNil())
		})

		It("Create/Get/Delete origin certificate", func() {
			shouldSkipTest()

			// create origin certificate
			certOpt := service.NewCreateOriginCertificateOptions(crn, zone_id)
			certOpt.SetCsr("-----BEGIN CERTIFICATE REQUEST-----\nMIICxzCCAa8CAQAwSDELMAkGA1UEBhMCVVMxFjAUBgNVBAgTDVNhbiBGcmFuY2lz\nY28xCzAJBgNVBAcTAkNBMRQwEgYDVQQDEwtleGFtcGxlLm5ldDCCASIwDQYJKoZI\nhvcNAQEBBQADggEPADCCAQoCggEBALxejtu4b+jPdFeFi6OUsye8TYJQBm3WfCvL\nHu5EvijMO/4Z2TImwASbwUF7Ir8OLgH+mGlQZeqyNvGoSOMEaZVXcYfpR1hlVak8\n4GGVr+04IGfOCqaBokaBFIwzclGZbzKmLGwIQioNxGfqFm6RGYGA3be2Je2iseBc\nN8GV1wYmvYE0RR+yWweJCTJ157exyRzu7sVxaEW9F87zBQLyOnwXc64rflXslRqi\ng7F7w5IaQYOl8yvmk/jEPCAha7fkiUfEpj4N12+oPRiMvleJF98chxjD4MH39c5I\nuOslULhrWunfh7GB1jwWNA9y44H0snrf+xvoy2TcHmxvma9Eln8CAwEAAaA6MDgG\nCSqGSIb3DQEJDjErMCkwJwYDVR0RBCAwHoILZXhhbXBsZS5uZXSCD3d3dy5leGFt\ncGxlLm5ldDANBgkqhkiG9w0BAQsFAAOCAQEAcBaX6dOnI8ncARrI9ZSF2AJX+8mx\npTHY2+Y2C0VvrVDGMtbBRH8R9yMbqWtlxeeNGf//LeMkSKSFa4kbpdx226lfui8/\nauRDBTJGx2R1ccUxmLZXx4my0W5iIMxunu+kez+BDlu7bTT2io0uXMRHue4i6quH\nyc5ibxvbJMjR7dqbcanVE10/34oprzXQsJ/VmSuZNXtjbtSKDlmcpw6To/eeAJ+J\nhXykcUihvHyG4A1m2R6qpANBjnA0pHexfwM/SgfzvpbvUg0T1ubmer8BgTwCKIWs\ndcWYTthM51JIqRBfNqy4QcBnX+GY05yltEEswQI55wdiS3CjTTA67sdbcQ==\n-----END CERTIFICATE REQUEST-----")
			certOpt.SetHostnames([]string{url})
			certOpt.SetRequestType(OriginCertificate_RequestType_OriginRsa)
			certOpt.SetRequestedValidity(5475)

			certResult, certResp, certErr := service.CreateOriginCertificate(certOpt)

			Expect(certErr).To(BeNil())
			Expect(certResp).ToNot(BeNil())
			Expect(certResult).ToNot(BeNil())
			Expect(certResult.Result.ID).ToNot(BeNil())
			Expect(*certResult.Result.RequestType).Should(BeEquivalentTo(OriginCertificate_RequestType_OriginRsa))
			originCertId := *certResult.Result.ID

			// get origin certificate

			getOpt := service.NewGetOriginCertificateOptions(crn, zone_id, originCertId)
			certResult, certResp, certErr = service.GetOriginCertificate(getOpt)

			Expect(certErr).To(BeNil())
			Expect(certResp).ToNot(BeNil())
			Expect(certResult).ToNot(BeNil())
			Expect(certResult.Result.ID).ToNot(BeNil())
			Expect(*certResult.Result.ID).Should(BeEquivalentTo(originCertId))
			Expect(*certResult.Result.RequestType).Should(BeEquivalentTo(OriginCertificate_RequestType_OriginRsa))

			// delete origin certificates

			delOpt := service.NewRevokeOriginCertificateOptions(crn, zone_id, originCertId)
			result, resp, err := service.RevokeOriginCertificate(delOpt)
			Expect(err).To(BeNil())
			Expect(resp).ToNot(BeNil())
			Expect(result).ToNot(BeNil())

		})
	})
})
