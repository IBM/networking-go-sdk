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
				certificate := "-----BEGIN CERTIFICATE-----\nMIIDpjCCAo4CCQDiw+/u+5c4eTANBgkqhkiG9w0BAQsFADCBlDELMAkGA1UEBhMC\nSU4xEjAQBgNVBAgMCUtBUk5BVEFLQTESMBAGA1UEBwwJQkFOR0FMT1JFMQwwCgYD\nVQQKDANJQk0xEjAQBgNVBAsMCUlCTSBDTE9VRDEXMBUGA1UEAwwOaWJtdGVzdG1h\nY2hpbmUxIjAgBgkqhkiG9w0BCQEWE2RhcnVueWEuZC5jQGlibS5jb20wHhcNMjIw\nNDA0MTM1ODE2WhcNMjMwNDA0MTM1ODE2WjCBlDELMAkGA1UEBhMCSU4xEjAQBgNV\nBAgMCUtBUk5BVEFLQTESMBAGA1UEBwwJQkFOR0FMT1JFMQwwCgYDVQQKDANJQk0x\nEjAQBgNVBAsMCUlCTSBDTE9VRDEXMBUGA1UEAwwOaWJtdGVzdG1hY2hpbmUxIjAg\nBgkqhkiG9w0BCQEWE2RhcnVueWEuZC5jQGlibS5jb20wggEiMA0GCSqGSIb3DQEB\nAQUAA4IBDwAwggEKAoIBAQCxg0xZgI+JExNCL41AK7FSphsHGP18/RsmrVHiQxGS\nONnh4pBtMJ+/HnnqEoko52L9GGWadu9494JG4vb1Oz3jBJx6vyOBAfJX9EIO0JCz\n/bDdOgyAl9L4MzXF0T5Mc511jHcwMH8jLgczC7hPVm2Acz68z3OFajViLEq7d3+a\n3pC1YV93P3BWn0tNCnHMfUmiXTg40iCVSl1BDpg1hwQnY/L6zAAF+k2jhCJ6W8Ny\nCukSbZ0sEzrhNteYASzWS9k2KMJT8PxoX6bmDWiVVIGHW08YnOC9OZjxHG8fagFs\npEn2FDFc0KNpH7lpLc1qMfWI/i/7cOkHjpqahuD6z9xLAgMBAAEwDQYJKoZIhvcN\nAQELBQADggEBAJIMKN23ChGVU6v+2GT3nnUh5IcZO5qb2bEJrvlyb30uVD8FoBkP\nh7dXlOGsh6tReLB0HLGOz9bnDO1Xzls73So8Ep3M2Xk/42JdzKwXL+Bw3KKTEHP/\n9QUijuwqFTW13KIwX2PWfpYpZTkQwWpi6FS7io+JtEAfO/c5MuwjaWLBEGm7t+HX\nIG21Z2TyIMhFfFoprZG98BSJA4bdqW5gZL2gijoKEtXYpkx65u+4txV566jg2dDr\nKwnFm3A0zHZ3ObRWNt6Vat0SUqOnMOeb0yGNNoxgnoc2NSXlg3+PH9e9FBs5uKE8\npfOqqBCXtdq9QUKjIJnujw/CsYWW4vqLNRI=\n-----END CERTIFICATE-----\n"
				privateKey := "-----BEGIN RSA PRIVATE KEY-----\nMIIEvAIBADANBgkqhkiG9w0BAQEFAASCBKYwggSiAgEAAoIBAQCxg0xZgI+JExNC\nL41AK7FSphsHGP18/RsmrVHiQxGSONnh4pBtMJ+/HnnqEoko52L9GGWadu9494JG\n4vb1Oz3jBJx6vyOBAfJX9EIO0JCz/bDdOgyAl9L4MzXF0T5Mc511jHcwMH8jLgcz\nC7hPVm2Acz68z3OFajViLEq7d3+a3pC1YV93P3BWn0tNCnHMfUmiXTg40iCVSl1B\nDpg1hwQnY/L6zAAF+k2jhCJ6W8NyCukSbZ0sEzrhNteYASzWS9k2KMJT8PxoX6bm\nDWiVVIGHW08YnOC9OZjxHG8fagFspEn2FDFc0KNpH7lpLc1qMfWI/i/7cOkHjpqa\nhuD6z9xLAgMBAAECggEAbO29NE8HxX3HG55Cd1ZYgfccLsbPBpvqxVkmHko5xhjM\n2yhEqDxmSslQ1qp5MHiM7fLCpn7FhN2dPBKaqPGpkF2MCGaySr//Dqn8v0qNAWZz\n2c19TovcEiKapME6EYAA59lCanfYDKZ6FIDkoQrQNzqBDSvgH8aE67FySoeR7l4O\nX2ltn2iEDdfbx7oFRdvA4mq2JfnIEK7faaVF+AybwDdn0WGyzcvsju85uEtHF/SS\n75BtwFZzyTNuEcjhBWevA9dTjBFMcnpWx2HgxyO/oKcNuiG8KFQMOoPLlfKw8Lwt\ntqhYmoV1ATmLsdsp9v3d2alvO3WlOrzWcBCS91jqQQKBgQDrAbLflXNOtFjwChEQ\nh6JzFbFGCeQw+avkh+h4qMPKpuX+kMzOQRZb30PDY+zxSDD5lJXqRxTNa5TsByVj\n8GAGWvy/84pLmB6KR/ujVJ+DH0OlqyEQTPttrTtrEkI0uw6cO/AGP6Fb8Gp0/3QN\n0np1She1iptPHxi9HFUhjFONvwKBgQDBXsp56yDmRAnl7twrCNvQTCB1dq7nPZE6\n/7N+Vpznon1k9At92rHErgYo9Rib21o/hPNeQJTIaey70ODB7q0BRv2e2PXwfyN3\nJrJwGYRRO8vLO/zrhHVxzjDt405bXR5R/j2IpR2pgrLXqpx+PfzujWFayQFkbzHf\ncqQBEwzsdQKBgGvzztBIHbzEuaoiZa5bL/N/vnw25PzeY+jJya9Ljw0TV8l1iK8i\nVPwE9mLWDyzTBbRQXgFNf6/RQIqfybw72lBxEXO3kwqgqT7KTDy+Dbw062U51Clh\nw4mhLw9DRuhkGRUJr3ufVScfrDdsdUo4Koqga324WxmgZkPQtQaBKIyPAoGAWudN\n9jyj7bwEjzRYCl8Svvxasf3GQWz/DiZQ4k6jWn1Xx5K2qEacFWLeAHkgRXy8E2pT\n4nYnu4OYR77tOh4S9KvD5N4H2DRcntHxRqOoQWwD5RnhT3Kop4SQGfUmy+qdq1wC\n328H371Sh/JruSk484hBQSWHYwinAG1rThn/lFUCgYA3okRpQKVulm7xg/g7pKwy\nxvBVlNVa0BxtqFO+vn//kfW7CMNSJQO5m+X0y7wYoX0yYsXHAX/PusuCsJGuvjnd\nnvKzVqYTUQ5HMA0rk80TizM34loHbrDnCwMS0WJTNyyt/QC+KRcrKHXvuHYgu2J6\nSjci4j+SSooywqCcC4Liww==\n-----END RSA PRIVATE KEY-----\n"
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
				certificate := "-----BEGIN CERTIFICATE-----\nMIIDpjCCAo4CCQDiw+/u+5c4eTANBgkqhkiG9w0BAQsFADCBlDELMAkGA1UEBhMC\nSU4xEjAQBgNVBAgMCUtBUk5BVEFLQTESMBAGA1UEBwwJQkFOR0FMT1JFMQwwCgYD\nVQQKDANJQk0xEjAQBgNVBAsMCUlCTSBDTE9VRDEXMBUGA1UEAwwOaWJtdGVzdG1h\nY2hpbmUxIjAgBgkqhkiG9w0BCQEWE2RhcnVueWEuZC5jQGlibS5jb20wHhcNMjIw\nNDA0MTM1ODE2WhcNMjMwNDA0MTM1ODE2WjCBlDELMAkGA1UEBhMCSU4xEjAQBgNV\nBAgMCUtBUk5BVEFLQTESMBAGA1UEBwwJQkFOR0FMT1JFMQwwCgYDVQQKDANJQk0x\nEjAQBgNVBAsMCUlCTSBDTE9VRDEXMBUGA1UEAwwOaWJtdGVzdG1hY2hpbmUxIjAg\nBgkqhkiG9w0BCQEWE2RhcnVueWEuZC5jQGlibS5jb20wggEiMA0GCSqGSIb3DQEB\nAQUAA4IBDwAwggEKAoIBAQCxg0xZgI+JExNCL41AK7FSphsHGP18/RsmrVHiQxGS\nONnh4pBtMJ+/HnnqEoko52L9GGWadu9494JG4vb1Oz3jBJx6vyOBAfJX9EIO0JCz\n/bDdOgyAl9L4MzXF0T5Mc511jHcwMH8jLgczC7hPVm2Acz68z3OFajViLEq7d3+a\n3pC1YV93P3BWn0tNCnHMfUmiXTg40iCVSl1BDpg1hwQnY/L6zAAF+k2jhCJ6W8Ny\nCukSbZ0sEzrhNteYASzWS9k2KMJT8PxoX6bmDWiVVIGHW08YnOC9OZjxHG8fagFs\npEn2FDFc0KNpH7lpLc1qMfWI/i/7cOkHjpqahuD6z9xLAgMBAAEwDQYJKoZIhvcN\nAQELBQADggEBAJIMKN23ChGVU6v+2GT3nnUh5IcZO5qb2bEJrvlyb30uVD8FoBkP\nh7dXlOGsh6tReLB0HLGOz9bnDO1Xzls73So8Ep3M2Xk/42JdzKwXL+Bw3KKTEHP/\n9QUijuwqFTW13KIwX2PWfpYpZTkQwWpi6FS7io+JtEAfO/c5MuwjaWLBEGm7t+HX\nIG21Z2TyIMhFfFoprZG98BSJA4bdqW5gZL2gijoKEtXYpkx65u+4txV566jg2dDr\nKwnFm3A0zHZ3ObRWNt6Vat0SUqOnMOeb0yGNNoxgnoc2NSXlg3+PH9e9FBs5uKE8\npfOqqBCXtdq9QUKjIJnujw/CsYWW4vqLNRI=\n-----END CERTIFICATE-----\n"
				privateKey := "-----BEGIN RSA PRIVATE KEY-----\nMIIEvAIBADANBgkqhkiG9w0BAQEFAASCBKYwggSiAgEAAoIBAQCxg0xZgI+JExNC\nL41AK7FSphsHGP18/RsmrVHiQxGSONnh4pBtMJ+/HnnqEoko52L9GGWadu9494JG\n4vb1Oz3jBJx6vyOBAfJX9EIO0JCz/bDdOgyAl9L4MzXF0T5Mc511jHcwMH8jLgcz\nC7hPVm2Acz68z3OFajViLEq7d3+a3pC1YV93P3BWn0tNCnHMfUmiXTg40iCVSl1B\nDpg1hwQnY/L6zAAF+k2jhCJ6W8NyCukSbZ0sEzrhNteYASzWS9k2KMJT8PxoX6bm\nDWiVVIGHW08YnOC9OZjxHG8fagFspEn2FDFc0KNpH7lpLc1qMfWI/i/7cOkHjpqa\nhuD6z9xLAgMBAAECggEAbO29NE8HxX3HG55Cd1ZYgfccLsbPBpvqxVkmHko5xhjM\n2yhEqDxmSslQ1qp5MHiM7fLCpn7FhN2dPBKaqPGpkF2MCGaySr//Dqn8v0qNAWZz\n2c19TovcEiKapME6EYAA59lCanfYDKZ6FIDkoQrQNzqBDSvgH8aE67FySoeR7l4O\nX2ltn2iEDdfbx7oFRdvA4mq2JfnIEK7faaVF+AybwDdn0WGyzcvsju85uEtHF/SS\n75BtwFZzyTNuEcjhBWevA9dTjBFMcnpWx2HgxyO/oKcNuiG8KFQMOoPLlfKw8Lwt\ntqhYmoV1ATmLsdsp9v3d2alvO3WlOrzWcBCS91jqQQKBgQDrAbLflXNOtFjwChEQ\nh6JzFbFGCeQw+avkh+h4qMPKpuX+kMzOQRZb30PDY+zxSDD5lJXqRxTNa5TsByVj\n8GAGWvy/84pLmB6KR/ujVJ+DH0OlqyEQTPttrTtrEkI0uw6cO/AGP6Fb8Gp0/3QN\n0np1She1iptPHxi9HFUhjFONvwKBgQDBXsp56yDmRAnl7twrCNvQTCB1dq7nPZE6\n/7N+Vpznon1k9At92rHErgYo9Rib21o/hPNeQJTIaey70ODB7q0BRv2e2PXwfyN3\nJrJwGYRRO8vLO/zrhHVxzjDt405bXR5R/j2IpR2pgrLXqpx+PfzujWFayQFkbzHf\ncqQBEwzsdQKBgGvzztBIHbzEuaoiZa5bL/N/vnw25PzeY+jJya9Ljw0TV8l1iK8i\nVPwE9mLWDyzTBbRQXgFNf6/RQIqfybw72lBxEXO3kwqgqT7KTDy+Dbw062U51Clh\nw4mhLw9DRuhkGRUJr3ufVScfrDdsdUo4Koqga324WxmgZkPQtQaBKIyPAoGAWudN\n9jyj7bwEjzRYCl8Svvxasf3GQWz/DiZQ4k6jWn1Xx5K2qEacFWLeAHkgRXy8E2pT\n4nYnu4OYR77tOh4S9KvD5N4H2DRcntHxRqOoQWwD5RnhT3Kop4SQGfUmy+qdq1wC\n328H371Sh/JruSk484hBQSWHYwinAG1rThn/lFUCgYA3okRpQKVulm7xg/g7pKwy\nxvBVlNVa0BxtqFO+vn//kfW7CMNSJQO5m+X0y7wYoX0yYsXHAX/PusuCsJGuvjnd\nnvKzVqYTUQ5HMA0rk80TizM34loHbrDnCwMS0WJTNyyt/QC+KRcrKHXvuHYgu2J6\nSjci4j+SSooywqCcC4Liww==\n-----END RSA PRIVATE KEY-----\n"
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
