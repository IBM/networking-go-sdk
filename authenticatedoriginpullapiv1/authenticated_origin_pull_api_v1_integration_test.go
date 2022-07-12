package authenticatedoriginpullapiv1_test

import (
	"fmt"
	"os"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	. "github.com/IBM/networking-go-sdk/authenticatedoriginpullapiv1"
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

var _ = Describe(`authenticatedoriginpullapiv1`, func() {
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
	globalOptions := &AuthenticatedOriginPullApiV1Options{
		ServiceName:    DefaultServiceName,
		URL:            serviceURL,
		Authenticator:  authenticator,
		Crn:            &crn,
		ZoneIdentifier: &zoneId,
	}

	service, serviceErr := NewAuthenticatedOriginPullApiV1(globalOptions)
	if serviceErr != nil {
		fmt.Println(serviceErr)
	}
	Describe(`authenticatedoriginpullapiv1_test`, func() {
		Context(`authenticatedoriginpullapiv1_test`, func() {
			defer GinkgoRecover()
			BeforeEach(func() {
				shouldSkipTest()
				listOpt := service.NewListZoneOriginPullCertificatesOptions()
				listOpt.SetXCorrelationID("45678")
				listResult, listResp, listErr := service.ListZoneOriginPullCertificates(listOpt)
				Expect(listErr).To(BeNil())
				Expect(listResp).ToNot(BeNil())
				Expect(listResult).ToNot(BeNil())
				Expect(*listResult.Success).Should(BeTrue())
				// Delete an Zone-Level Authenticated Origin Pull
				for _, originPull := range listResult.Result {
					delOpt := service.NewDeleteZoneOriginPullCertificateOptions(*originPull.ID)
					delOpt.SetCertIdentifier(*originPull.ID)
					delOpt.SetXCorrelationID("45678")
					delResult, delResp, delErr := service.DeleteZoneOriginPullCertificate(delOpt)
					Expect(delErr).To(BeNil())
					Expect(delResp).ToNot(BeNil())
					Expect(delResult).ToNot(BeNil())
					Expect(*delResult.Success).Should(BeTrue())
				}
			})
			AfterEach(func() {
				shouldSkipTest()
				listOpt := service.NewListZoneOriginPullCertificatesOptions()
				listOpt.SetXCorrelationID("45678")
				listResult, listResp, listErr := service.ListZoneOriginPullCertificates(listOpt)
				Expect(listErr).To(BeNil())
				Expect(listResp).ToNot(BeNil())
				Expect(listResult).ToNot(BeNil())
				Expect(*listResult.Success).Should(BeTrue())
				// Delete an Zone-Level Authenticated Origin Pull
				for _, originPull := range listResult.Result {
					delOpt := service.NewDeleteZoneOriginPullCertificateOptions(*originPull.ID)
					delOpt.SetCertIdentifier(*originPull.ID)
					delOpt.SetXCorrelationID("45678")
					delResult, delResp, delErr := service.DeleteZoneOriginPullCertificate(delOpt)
					Expect(delErr).To(BeNil())
					Expect(delResp).ToNot(BeNil())
					Expect(delResult).ToNot(BeNil())
					Expect(*delResult.Success).Should(BeTrue())
				}
			})
			It(`Create| Get | Update | List | Delete Zone-Level Authenticated Origin Pull`, func() {
				// Create a new Zone-Level Authenticated Origin Pull
				certificate := "-----BEGIN CERTIFICATE-----\nMIICsjCCAZoCCQDTevWcEr/TDjANBgkqhkiG9w0BAQsFADAbMQswCQYDVQQGEwJJ\nTjEMMAoGA1UECgwDSUJNMB4XDTIyMDcxMjExNDIzMloXDTIzMDcxMjExNDIzMlow\nGzELMAkGA1UEBhMCSU4xDDAKBgNVBAoMA0lCTTCCASIwDQYJKoZIhvcNAQEBBQAD\nggEPADCCAQoCggEBALyRkJwi0FU5sE8ILqiuIDX8Gx6lZ5/cv9wtbLuDrCkPkvPq\nZPOHkBcai2cCXUomRrJL6ouZ79FrMG+kZNp6yY9PwMvzsGExUPVd8GMqbnCT/W/D\nf1JodnzkAmdsiiIyy9wM/E4Zv9HeqtA4tBXRjZSB7NobgAPjwb6D08Zd8SWkIhtL\nB3fKYOahuw0M3yubV1nE74CNvvOlMcB0CsZdt486H5lnY4lSp8nkDKcv5M89i3K1\nYwiJRBvzyHOXmr0MrzmzUyKORchCNH8cNYAjT3+E26AoIA6ca20cLnqAlQ9aUSx9\nLboy8MQ1+G9jxS0ienqp802FfGZB3/mdSRGMPfUCAwEAATANBgkqhkiG9w0BAQsF\nAAOCAQEAEPjcpMhJy5fOQ3x2MEw5q4EitoDcpP5eHXsszOM6apjmw0DP6Ayapv/S\nUK9obfbLqlUvnYd8iZ//OWaSBWZmwUCvrWb+n+gdC6oebdMc/jDpxqE7FPCVd10u\nzuZHnbU1qwrlTyANP7QcQRS8vqg2AqsxNBqfjOj8+IFOZ4Ifi9L3XsjT4meI4S2F\ndTaa8JxaD9I2bsgbBRFCW41DzXVyNTZVFIVPc+1CWD8u5Fq15sa0TwSiwbhbk3Eg\nu12WGJbc/VNotc6fCTVseFa2Rkt+MaiVa5HjnUKj4s591eNw81lX7/t2ktr4gbbb\n8o8Efj50/Ems56AF1XxSKjXuy+Ne1A==\n-----END CERTIFICATE-----\n"
				privateKey := "-----BEGIN PRIVATE KEY-----\nMIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQC8kZCcItBVObBP\nCC6oriA1/BsepWef3L/cLWy7g6wpD5Lz6mTzh5AXGotnAl1KJkayS+qLme/RazBv\npGTaesmPT8DL87BhMVD1XfBjKm5wk/1vw39SaHZ85AJnbIoiMsvcDPxOGb/R3qrQ\nOLQV0Y2UgezaG4AD48G+g9PGXfElpCIbSwd3ymDmobsNDN8rm1dZxO+Ajb7zpTHA\ndArGXbePOh+ZZ2OJUqfJ5AynL+TPPYtytWMIiUQb88hzl5q9DK85s1MijkXIQjR/\nHDWAI09/hNugKCAOnGttHC56gJUPWlEsfS26MvDENfhvY8UtInp6qfNNhXxmQd/5\nnUkRjD31AgMBAAECggEAG50BXSvn8CMDg39CPedJxj4FxqYYF5ve6KIqQwdUJasn\nLNrNL7BRVGDJuyroeFxEjAV78jd3m+YjcKXVMv06GOdn5zXnRURQg63e7tae15OE\nUGKDeZDckQteosCNNdbUbYPlSpaQtW3y+4ziKjPGlNG12naed3NZwslRMMI+9vDi\nUWR7ghKASOFLrmYlh0ehP5UK3OsXFE7Ed42nBoihXMFcBNHMxQm4F31vxWAytvi7\nvxCtbbyMJ7IbPWZEM85krKDikhF5AGnPozbSczFbIVPiQvmbi61Z/WjkkJ6FUdsP\nST524yThlDxikYCWDhw05h1WFd8EObIvHSghHBBVYQKBgQD3wsy4V1BsGLhzLe1H\nJG6s+vzcxpQGcu8OT1I6NBgk5fULcNPeXHtFbTNaoq9KKapZfEkO2pHLucjYdEe2\nfnwzg/QEcdnlh1JcW+hdvEQ414D63JUT001kmlg5XEPfAz89zJu2Kp21WiBZGm5J\nQBf88BJs0y7cDtABedEYRg0JKQKBgQDC1tuzb86/H8LoOSEkmArNA/IFY9fX8SaU\nYBVlOegbQOy14l9a2KRZ+wTZGfn/5osX7uO2hO2F+zH7PZyDR2C/tDY+9Af5POMl\nR+v2cOQDx2Jzqr5T0d165IHpLkQ8L5LKEEMzY9eUiE7mGola/Yu7BMinw0PfCRbi\nwhQZbCYL7QKBgQDQCSSG4OHpcjRmmjizVNcNrk3mP2OJqrYqCNadgqKHUQOaEKoF\n+xeS6yeEwjd3iVa9fsuFimeDbcNEZRbWGIzHYNPja4mv3hl87btGAdAy/lkRy2ft\n1q4UfDj6KQvgVUSj6osQweXcogmpZ7UVEplRzG9cK1Mced+UbanxvNgzSQKBgCy1\nHIiZ+TjF0vVyVnaNJL1SUHCILnjwbsfRHFez59yJE0fQ/8xataun+77NRR5BCl2d\nhUbWTaJWt2tNAeLlt/+FHIVpfYLlQ8HENRLBaLCtSZv869tT5pxSXrTg1utwhyAy\nhxj9qfP9Kw2FvUrRrwRk3p4QIjzFWykBG5eRx1EpAoGBAMmcrGol2vsP5mYqRq0r\n2GhfiZrWdpI/Dem3bHwuoyYTkrMLtDaIVOYkZvdN8bSMlTU0msuSKOAsh9NVjbV5\nTlDH4w4ktWIMg87d1fggaFVPXRAe2rR1OQG9nsfpUFfFRncrX0XwqPw6pHwwC1NF\npbE3oZ9bofWnPfePVc7U/IZx\n-----END PRIVATE KEY-----\n"
				xCorrelationID := "45678"
				options := service.NewUploadZoneOriginPullCertificateOptions()
				options.SetCertificate(certificate)
				options.SetPrivateKey(privateKey)
				options.SetXCorrelationID(xCorrelationID)
				result, response, operationErr := service.UploadZoneOriginPullCertificate(options)
				ID := *result.Result.ID
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Get Zone-Level Authenticated Origin Pull
				getOptions := service.NewGetZoneOriginPullCertificateOptions(ID)
				getOptions.SetCertIdentifier(ID)
				getOptions.SetXCorrelationID(xCorrelationID)
				getResult, getResp, getErr := service.GetZoneOriginPullCertificate(getOptions)
				Expect(getErr).To(BeNil())
				Expect(getResp).ToNot(BeNil())
				Expect(getResult).ToNot(BeNil())

				// Set Zone Origin Pull Settings
				updateOption := service.NewSetZoneOriginPullSettingsOptions()
				updateOption.SetEnabled(true)
				updateOption.SetXCorrelationID(xCorrelationID)
				updateResult, updateResp, updateErr := service.SetZoneOriginPullSettings(updateOption)
				Expect(updateErr).To(BeNil())
				Expect(updateResp).ToNot(BeNil())
				Expect(updateResult).ToNot(BeNil())
				Expect(*updateResult.Success).Should(BeTrue())

				// Get Zone Origin Pull Settings
				getZoneOption := service.NewGetZoneOriginPullSettingsOptions()
				getZoneOption.SetXCorrelationID(xCorrelationID)
				getZoneResult, getZoneResp, getZoneErr := service.GetZoneOriginPullSettings(getZoneOption)
				Expect(getZoneErr).To(BeNil())
				Expect(getZoneResp).ToNot(BeNil())
				Expect(getZoneResult).ToNot(BeNil())
				Expect(*getZoneResult.Success).Should(BeTrue())

				// List Zone-Level Authenticated Origin Pull
				listOpt := service.NewListZoneOriginPullCertificatesOptions()
				listOpt.SetXCorrelationID(xCorrelationID)
				listResult, listResp, listErr := service.ListZoneOriginPullCertificates(listOpt)
				Expect(listErr).To(BeNil())
				Expect(listResp).ToNot(BeNil())
				Expect(listResult).ToNot(BeNil())
				Expect(*listResult.Success).Should(BeTrue())

				// delete Zone-Level Authenticated Origin Pull
				delOpt := service.NewDeleteZoneOriginPullCertificateOptions(ID)
				delOpt.SetCertIdentifier(ID)
				delOpt.SetXCorrelationID(xCorrelationID)
				delResult, delResp, delErr := service.DeleteZoneOriginPullCertificate(delOpt)
				Expect(delErr).To(BeNil())
				Expect(delResp).ToNot(BeNil())
				Expect(delResult).ToNot(BeNil())
				Expect(*delResult.Success).Should(BeTrue())
			})
			It(`Create| Get | Update | List | Delete Hostname-Level Authenticated Origin Pull`, func() {
				// Create a new Zone-Level Authenticated Origin Pull
				certificate := "-----BEGIN CERTIFICATE-----\nMIICsjCCAZoCCQDTevWcEr/TDjANBgkqhkiG9w0BAQsFADAbMQswCQYDVQQGEwJJ\nTjEMMAoGA1UECgwDSUJNMB4XDTIyMDcxMjExNDIzMloXDTIzMDcxMjExNDIzMlow\nGzELMAkGA1UEBhMCSU4xDDAKBgNVBAoMA0lCTTCCASIwDQYJKoZIhvcNAQEBBQAD\nggEPADCCAQoCggEBALyRkJwi0FU5sE8ILqiuIDX8Gx6lZ5/cv9wtbLuDrCkPkvPq\nZPOHkBcai2cCXUomRrJL6ouZ79FrMG+kZNp6yY9PwMvzsGExUPVd8GMqbnCT/W/D\nf1JodnzkAmdsiiIyy9wM/E4Zv9HeqtA4tBXRjZSB7NobgAPjwb6D08Zd8SWkIhtL\nB3fKYOahuw0M3yubV1nE74CNvvOlMcB0CsZdt486H5lnY4lSp8nkDKcv5M89i3K1\nYwiJRBvzyHOXmr0MrzmzUyKORchCNH8cNYAjT3+E26AoIA6ca20cLnqAlQ9aUSx9\nLboy8MQ1+G9jxS0ienqp802FfGZB3/mdSRGMPfUCAwEAATANBgkqhkiG9w0BAQsF\nAAOCAQEAEPjcpMhJy5fOQ3x2MEw5q4EitoDcpP5eHXsszOM6apjmw0DP6Ayapv/S\nUK9obfbLqlUvnYd8iZ//OWaSBWZmwUCvrWb+n+gdC6oebdMc/jDpxqE7FPCVd10u\nzuZHnbU1qwrlTyANP7QcQRS8vqg2AqsxNBqfjOj8+IFOZ4Ifi9L3XsjT4meI4S2F\ndTaa8JxaD9I2bsgbBRFCW41DzXVyNTZVFIVPc+1CWD8u5Fq15sa0TwSiwbhbk3Eg\nu12WGJbc/VNotc6fCTVseFa2Rkt+MaiVa5HjnUKj4s591eNw81lX7/t2ktr4gbbb\n8o8Efj50/Ems56AF1XxSKjXuy+Ne1A==\n-----END CERTIFICATE-----\n"
				privateKey := "-----BEGIN PRIVATE KEY-----\nMIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQC8kZCcItBVObBP\nCC6oriA1/BsepWef3L/cLWy7g6wpD5Lz6mTzh5AXGotnAl1KJkayS+qLme/RazBv\npGTaesmPT8DL87BhMVD1XfBjKm5wk/1vw39SaHZ85AJnbIoiMsvcDPxOGb/R3qrQ\nOLQV0Y2UgezaG4AD48G+g9PGXfElpCIbSwd3ymDmobsNDN8rm1dZxO+Ajb7zpTHA\ndArGXbePOh+ZZ2OJUqfJ5AynL+TPPYtytWMIiUQb88hzl5q9DK85s1MijkXIQjR/\nHDWAI09/hNugKCAOnGttHC56gJUPWlEsfS26MvDENfhvY8UtInp6qfNNhXxmQd/5\nnUkRjD31AgMBAAECggEAG50BXSvn8CMDg39CPedJxj4FxqYYF5ve6KIqQwdUJasn\nLNrNL7BRVGDJuyroeFxEjAV78jd3m+YjcKXVMv06GOdn5zXnRURQg63e7tae15OE\nUGKDeZDckQteosCNNdbUbYPlSpaQtW3y+4ziKjPGlNG12naed3NZwslRMMI+9vDi\nUWR7ghKASOFLrmYlh0ehP5UK3OsXFE7Ed42nBoihXMFcBNHMxQm4F31vxWAytvi7\nvxCtbbyMJ7IbPWZEM85krKDikhF5AGnPozbSczFbIVPiQvmbi61Z/WjkkJ6FUdsP\nST524yThlDxikYCWDhw05h1WFd8EObIvHSghHBBVYQKBgQD3wsy4V1BsGLhzLe1H\nJG6s+vzcxpQGcu8OT1I6NBgk5fULcNPeXHtFbTNaoq9KKapZfEkO2pHLucjYdEe2\nfnwzg/QEcdnlh1JcW+hdvEQ414D63JUT001kmlg5XEPfAz89zJu2Kp21WiBZGm5J\nQBf88BJs0y7cDtABedEYRg0JKQKBgQDC1tuzb86/H8LoOSEkmArNA/IFY9fX8SaU\nYBVlOegbQOy14l9a2KRZ+wTZGfn/5osX7uO2hO2F+zH7PZyDR2C/tDY+9Af5POMl\nR+v2cOQDx2Jzqr5T0d165IHpLkQ8L5LKEEMzY9eUiE7mGola/Yu7BMinw0PfCRbi\nwhQZbCYL7QKBgQDQCSSG4OHpcjRmmjizVNcNrk3mP2OJqrYqCNadgqKHUQOaEKoF\n+xeS6yeEwjd3iVa9fsuFimeDbcNEZRbWGIzHYNPja4mv3hl87btGAdAy/lkRy2ft\n1q4UfDj6KQvgVUSj6osQweXcogmpZ7UVEplRzG9cK1Mced+UbanxvNgzSQKBgCy1\nHIiZ+TjF0vVyVnaNJL1SUHCILnjwbsfRHFez59yJE0fQ/8xataun+77NRR5BCl2d\nhUbWTaJWt2tNAeLlt/+FHIVpfYLlQ8HENRLBaLCtSZv869tT5pxSXrTg1utwhyAy\nhxj9qfP9Kw2FvUrRrwRk3p4QIjzFWykBG5eRx1EpAoGBAMmcrGol2vsP5mYqRq0r\n2GhfiZrWdpI/Dem3bHwuoyYTkrMLtDaIVOYkZvdN8bSMlTU0msuSKOAsh9NVjbV5\nTlDH4w4ktWIMg87d1fggaFVPXRAe2rR1OQG9nsfpUFfFRncrX0XwqPw6pHwwC1NF\npbE3oZ9bofWnPfePVc7U/IZx\n-----END PRIVATE KEY-----\n"
				xCorrelationID := "12345"
				options := service.NewUploadHostnameOriginPullCertificateOptions()
				options.SetCertificate(certificate)
				options.SetPrivateKey(privateKey)
				options.SetXCorrelationID(xCorrelationID)
				result, response, operationErr := service.UploadHostnameOriginPullCertificate(options)
				ID := *result.Result.ID
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
				time.Sleep(5 * time.Second)

				// Get Zone-Level Authenticated Origin Pull
				getOptions := service.NewGetHostnameOriginPullCertificateOptions(ID)
				getOptions.SetCertIdentifier(ID)
				getOptions.SetXCorrelationID(xCorrelationID)
				getResult, getResp, getErr := service.GetHostnameOriginPullCertificate(getOptions)
				Expect(getErr).To(BeNil())
				Expect(getResp).ToNot(BeNil())
				Expect(getResult).ToNot(BeNil())

				// Set Zone Origin Pull Settings
				model := &HostnameOriginPullSettings{
					Hostname: core.StringPtr(URL),
					CertID:   core.StringPtr(ID),
					Enabled:  core.BoolPtr(false),
				}
				setOption := service.NewSetHostnameOriginPullSettingsOptions()
				setOption.SetConfig([]HostnameOriginPullSettings{*model})
				setOption.SetXCorrelationID(xCorrelationID)
				setResult, setResp, setErr := service.SetHostnameOriginPullSettings(setOption)
				Expect(setErr).To(BeNil())
				Expect(setResp).ToNot(BeNil())
				Expect(setResult).ToNot(BeNil())
				Expect(*setResult.Success).Should(BeTrue())

				// Get Hostname-Level Origin Pull Settings
				getHostnameOption := service.NewGetHostnameOriginPullSettingsOptions(URL)
				getHostnameOption.SetHostname(URL)
				getHostnameOption.SetXCorrelationID(xCorrelationID)
				getHostnameResult, getHostnameResp, getHostnameErr := service.GetHostnameOriginPullSettings(getHostnameOption)
				Expect(getHostnameErr).To(BeNil())
				Expect(getHostnameResp).ToNot(BeNil())
				Expect(getHostnameResult).ToNot(BeNil())
				Expect(*getHostnameResult.Success).Should(BeTrue())

				// delete Hostname-Level Authenticated Origin Pull
				delOpt := service.NewDeleteHostnameOriginPullCertificateOptions(ID)
				delOpt.SetCertIdentifier(ID)
				delOpt.SetXCorrelationID("12345")
				delResult, delResp, delErr := service.DeleteHostnameOriginPullCertificate(delOpt)
				Expect(delErr).To(BeNil())
				Expect(delResp).ToNot(BeNil())
				Expect(delResult).ToNot(BeNil())
				Expect(*delResult.Success).Should(BeTrue())
			})
		})
	})
})
