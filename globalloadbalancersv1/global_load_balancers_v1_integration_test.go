package globalloadbalancersv1_test

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/IBM/go-sdk-core/core"
	guuid "github.com/google/uuid"
	"github.com/joho/godotenv"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/IBM/networking-go-sdk/dnszonesv1"
	. "github.com/IBM/networking-go-sdk/globalloadbalancersv1"
)

const configFile = "../pdns.env"

var configLoaded bool = true

func shouldSkipTest() {
	if !configLoaded {
		Skip("External configuration is not available, skipping...")
	}
}

var _ = Describe(`PDNSgloballoadbalancersv1`, func() {
	if _, err := os.Stat(configFile); err != nil {
		configLoaded = false
	}

	err := godotenv.Load(configFile)
	if err != nil {
		configLoaded = false
	}

	authenticator := &core.IamAuthenticator{
		ApiKey: os.Getenv("PDNS_SERVICES_APIKEY"),
		URL:    os.Getenv("PDNS_SERVICES_AUTH_URL"),
	}
	serviceURL := os.Getenv("API_ENDPOINT")
	instanceID := os.Getenv("INSTANCE_ID")
	dnsRecordOptions := &GlobalLoadBalancersV1Options{
		ServiceName:   "pdns_glb_services",
		URL:           serviceURL,
		Authenticator: authenticator,
	}

	zoneOptions := &dnszonesv1.DnsZonesV1Options{
		ServiceName:   "pdns_services",
		URL:           serviceURL,
		Authenticator: authenticator,
	}

	service, serviceErr := NewGlobalLoadBalancersV1(dnsRecordOptions)
	if serviceErr != nil {
		fmt.Println(serviceErr)
	}

	zoneService, zoneServiceErr := dnszonesv1.NewDnsZonesV1(zoneOptions)
	if zoneServiceErr != nil {
		fmt.Println(zoneServiceErr)
	}
	Describe(`PDNSgloballoadbalancersv1`, func() {
		Context(`PDNSgloballoadbalancersv1`, func() {
			var zoneInfo *dnszonesv1.Dnszone
			BeforeEach(func() {
				shouldSkipTest()

				// delete all dns zones
				listOptions := zoneService.NewListDnszonesOptions(instanceID)
				listResult, listResp, listErr := zoneService.ListDnszones(listOptions)
				Expect(listErr).To(BeNil())
				Expect(listResp).ToNot(BeNil())
				Expect(listResult).ToNot(BeNil())
				Expect(listResp.GetStatusCode()).To(BeEquivalentTo(200))
				for _, zone := range listResult.Dnszones {
					if strings.Contains(*zone.Name, "test-example") {
						//delete all PDNS GLB load balancers
						listLoadBalancerOptions := service.NewListLoadBalancersOptions(instanceID, *zone.ID)
						listLoadBalancerResult, listLoadBalancerResp, listLoadBalancerErr := service.ListLoadBalancers(listLoadBalancerOptions)
						Expect(listLoadBalancerErr).To(BeNil())
						Expect(listLoadBalancerResp).ToNot(BeNil())
						Expect(listLoadBalancerResult).ToNot(BeNil())

						for _, record := range listLoadBalancerResult.LoadBalancers {
							deleteOpt := service.NewDeleteLoadBalancerOptions(instanceID, *zone.ID, *record.ID)
							deleteResponse, deleteErr := service.DeleteLoadBalancer(deleteOpt)
							Expect(deleteErr).To(BeNil())
							Expect(deleteResponse).ToNot(BeNil())
						}
						//delete all PDNS GLB Pools
						listPoolOptions := service.NewListPoolsOptions(instanceID)
						listPoolResult, listPoolResp, listPoolErr := service.ListPools(listPoolOptions)
						Expect(listPoolErr).To(BeNil())
						Expect(listPoolResp).ToNot(BeNil())
						Expect(listPoolResult).ToNot(BeNil())

						for _, record := range listPoolResult.Pools {
							deleteOpt := service.NewDeletePoolOptions(instanceID, *record.ID)
							deleteResponse, deleteErr := service.DeletePool(deleteOpt)
							Expect(deleteErr).To(BeNil())
							Expect(deleteResponse).ToNot(BeNil())
						}
						// delete all PDNS GLB Monitors
						listOptions := service.NewListMonitorsOptions(instanceID)
						listResult, listResp, listErr := service.ListMonitors(listOptions)
						Expect(listErr).To(BeNil())
						Expect(listResp).ToNot(BeNil())
						Expect(listResult).ToNot(BeNil())

						for _, record := range listResult.Monitors {
							deleteOpt := service.NewDeleteMonitorOptions(instanceID, *record.ID)
							deleteResponse, deleteErr := service.DeleteMonitor(deleteOpt)
							Expect(deleteErr).To(BeNil())
							Expect(deleteResponse).ToNot(BeNil())
						}
						// delete zone
						option := zoneService.NewDeleteDnszoneOptions(instanceID, *zone.ID)
						response, err := zoneService.DeleteDnszone(option)
						Expect(err).To(BeNil())
						Expect(response).ToNot(BeNil())
						Expect(response.GetStatusCode()).To(BeEquivalentTo(204))
					}
				}

				// Create DNS Zone
				zoneName := fmt.Sprintf("test-example%s.com", guuid.New().String())
				createDnszoneOptions := zoneService.NewCreateDnszoneOptions(instanceID)
				createDnszoneOptions.SetName(zoneName)
				createDnszoneOptions.SetDescription("testString")
				createDnszoneOptions.SetLabel("testString")
				createDnszoneOptions.SetXCorrelationID("abc123")
				result, response, reqErr := zoneService.CreateDnszone(createDnszoneOptions)
				Expect(reqErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
				Expect(response.GetStatusCode()).To(BeEquivalentTo(200))
				zoneInfo = result
			})
			AfterEach(func() {
				shouldSkipTest()

				// delete all dns zones
				listOptions := zoneService.NewListDnszonesOptions(instanceID)
				listResult, listResp, listErr := zoneService.ListDnszones(listOptions)
				Expect(listErr).To(BeNil())
				Expect(listResp).ToNot(BeNil())
				Expect(listResult).ToNot(BeNil())
				Expect(listResp.GetStatusCode()).To(BeEquivalentTo(200))
				for _, zone := range listResult.Dnszones {
					if strings.Contains(*zone.Name, "test-example") {
						//delete all PDNS GLB load balancers
						listLoadBalancerOptions := service.NewListLoadBalancersOptions(instanceID, *zone.ID)
						listLoadBalancerResult, listLoadBalancerResp, listLoadBalancerErr := service.ListLoadBalancers(listLoadBalancerOptions)
						Expect(listLoadBalancerErr).To(BeNil())
						Expect(listLoadBalancerResp).ToNot(BeNil())
						Expect(listLoadBalancerResult).ToNot(BeNil())

						for _, record := range listLoadBalancerResult.LoadBalancers {
							deleteOpt := service.NewDeleteLoadBalancerOptions(instanceID, *zone.ID, *record.ID)
							deleteResponse, deleteErr := service.DeleteLoadBalancer(deleteOpt)
							Expect(deleteErr).To(BeNil())
							Expect(deleteResponse).ToNot(BeNil())
						}
						// delete all PDNS GLB pools
						listPoolOptions := service.NewListPoolsOptions(instanceID)
						listPoolResult, listPoolResp, listPoolErr := service.ListPools(listPoolOptions)
						Expect(listPoolErr).To(BeNil())
						Expect(listPoolResp).ToNot(BeNil())
						Expect(listPoolResult).ToNot(BeNil())

						for _, record := range listPoolResult.Pools {
							deleteOpt := service.NewDeletePoolOptions(instanceID, *record.ID)
							deleteResponse, deleteErr := service.DeletePool(deleteOpt)
							Expect(deleteErr).To(BeNil())
							Expect(deleteResponse).ToNot(BeNil())
						}
						// delete all PDNS GLB Monitors
						listOptions := service.NewListMonitorsOptions(instanceID)
						listResult, listResp, listErr := service.ListMonitors(listOptions)
						Expect(listErr).To(BeNil())
						Expect(listResp).ToNot(BeNil())
						Expect(listResult).ToNot(BeNil())

						for _, record := range listResult.Monitors {
							deleteOpt := service.NewDeleteMonitorOptions(instanceID, *record.ID)
							deleteResponse, deleteErr := service.DeleteMonitor(deleteOpt)
							Expect(deleteErr).To(BeNil())
							Expect(deleteResponse).ToNot(BeNil())
						}
						// delete zone
						option := zoneService.NewDeleteDnszoneOptions(instanceID, *zone.ID)
						response, err := zoneService.DeleteDnszone(option)
						Expect(err).To(BeNil())
						Expect(response).ToNot(BeNil())
						Expect(response.GetStatusCode()).To(BeEquivalentTo(204))
					}
				}
			})
			It(`crate/update/get/delete PDNS GLB monitor,pool and load balancer`, func() {
				shouldSkipTest()

				// create Load Balancer Monitor
				CreateMonitorOptions := service.NewCreateMonitorOptions(instanceID)
				CreateMonitorOptions.SetName("testa")
				CreateMonitorOptions.SetExpectedCodes("200")
				CreateMonitorOptions.SetType(CreateMonitorOptions_Type_Http)
				CreateMonitorOptions.SetDescription("PDNS Load balancer monitor.")
				CreateMonitorOptions.SetPort(8080)
				CreateMonitorOptions.SetInterval(60)
				CreateMonitorOptions.SetRetries(2)
				CreateMonitorOptions.SetTimeout(5)
				CreateMonitorOptions.SetMethod(CreateMonitorOptions_Method_Get)
				CreateMonitorOptions.SetPath("health")
				CreateMonitorOptions.SetAllowInsecure(false)
				CreateMonitorOptions.SetExpectedBody("alive")
				result, response, reqErr := service.CreateMonitor(CreateMonitorOptions)
				Expect(reqErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
				Expect(response.GetStatusCode()).To(BeEquivalentTo(200))
				Expect(*result.Type).To(BeEquivalentTo(CreateMonitorOptions_Type_Http))
				Expect(*result.Name).To(BeEquivalentTo("testa"))
				Expect(*result.Description).To(BeEquivalentTo("PDNS Load balancer monitor."))
				Expect(*result.Port).To(BeEquivalentTo(8080))
				Expect(*result.Interval).To(BeEquivalentTo(60))
				Expect(*result.Retries).To(BeEquivalentTo(2))
				Expect(*result.Timeout).To(BeEquivalentTo(5))
				Expect(*result.Method).To(BeEquivalentTo(CreateMonitorOptions_Method_Get))
				Expect(*result.Path).To(BeEquivalentTo("health"))
				Expect(*result.AllowInsecure).To(BeEquivalentTo(false))
				Expect(*result.ExpectedCodes).To(BeEquivalentTo("200"))
				Expect(*result.ExpectedBody).To(BeEquivalentTo("alive"))

				//Test GetMonitor
				getOpt := service.NewGetMonitorOptions(instanceID, *result.ID)
				getResult, getResponse, getErr := service.GetMonitor(getOpt)
				Expect(getErr).To(BeNil())
				Expect(getResponse).ToNot(BeNil())
				Expect(getResult).ToNot(BeNil())
				Expect(getResponse.GetStatusCode()).To(BeEquivalentTo(200))
				Expect(*getResult.Type).To(BeEquivalentTo(CreateMonitorOptions_Type_Http))
				Expect(*getResult.Name).To(BeEquivalentTo("testa"))
				Expect(*getResult.Description).To(BeEquivalentTo("PDNS Load balancer monitor."))
				Expect(*getResult.Port).To(BeEquivalentTo(8080))
				Expect(*getResult.Interval).To(BeEquivalentTo(60))
				Expect(*getResult.Retries).To(BeEquivalentTo(2))
				Expect(*getResult.Timeout).To(BeEquivalentTo(5))
				Expect(*getResult.Method).To(BeEquivalentTo(CreateMonitorOptions_Method_Get))
				Expect(*getResult.Path).To(BeEquivalentTo("health"))
				Expect(*getResult.AllowInsecure).To(BeEquivalentTo(false))
				Expect(*getResult.ExpectedCodes).To(BeEquivalentTo("200"))

				//Test UpdateMonitor
				updateOpt := service.NewUpdateMonitorOptions(instanceID, *result.ID)
				updateOpt.SetName("updatea")
				updateOpt.SetType(UpdateMonitorOptions_Type_Https)
				updateResult, updateResponse, updateErr := service.UpdateMonitor(updateOpt)
				Expect(updateErr).To(BeNil())
				Expect(updateResponse).ToNot(BeNil())
				Expect(updateResult).ToNot(BeNil())
				Expect(updateResponse.GetStatusCode()).To(BeEquivalentTo(200))
				Expect(*updateResult.Name).To(BeEquivalentTo("updatea"))
				Expect(*updateResult.Type).To(BeEquivalentTo(UpdateMonitorOptions_Type_Https))

				//Test CreatePool
				CreatePoolOptions := service.NewCreatePoolOptions(instanceID)
				CreatePoolOptions.SetName("testPool")
				CreatePoolOptions.SetDescription("creating pool")
				CreatePoolOptions.SetEnabled(true)
				CreatePoolOptions.SetHealthyOriginsThreshold(1)
				originInputModel := new(OriginInput)
				originInputModel.Name = core.StringPtr("app-server-1")
				originInputModel.Description = core.StringPtr("description of the origin server")
				originInputModel.Address = core.StringPtr("10.10.10.8")
				originInputModel.Enabled = core.BoolPtr(true)
				CreatePoolOptions.Origins = []OriginInput{*originInputModel}
				resultPool, responsePool, reqErrPool := service.CreatePool(CreatePoolOptions)
				Expect(reqErrPool).To(BeNil())
				Expect(responsePool).ToNot(BeNil())
				Expect(resultPool).ToNot(BeNil())
				Expect(responsePool.GetStatusCode()).To(BeEquivalentTo(200))
				Expect(*resultPool.Name).To(BeEquivalentTo("testPool"))
				Expect(*resultPool.Description).To(BeEquivalentTo("creating pool"))
				Expect(*resultPool.Enabled).To(BeEquivalentTo(true))
				Expect(*resultPool.HealthyOriginsThreshold).To(BeEquivalentTo(1))

				//Test Get Pool
				getPoolOpt := service.NewGetPoolOptions(instanceID, *resultPool.ID)
				getPoolResult, gePooltResponse, getPoolErr := service.GetPool(getPoolOpt)
				Expect(getPoolErr).To(BeNil())
				Expect(gePooltResponse).ToNot(BeNil())
				Expect(getPoolResult).ToNot(BeNil())
				Expect(gePooltResponse.GetStatusCode()).To(BeEquivalentTo(200))
				Expect(*getPoolResult.Name).To(BeEquivalentTo("testPool"))
				Expect(*getPoolResult.Description).To(BeEquivalentTo("creating pool"))
				Expect(*getPoolResult.Enabled).To(BeEquivalentTo(true))
				Expect(*getPoolResult.HealthyOriginsThreshold).To(BeEquivalentTo(1))

				//Test Update Pool
				updatePoolOpt := service.NewUpdatePoolOptions(instanceID, *resultPool.ID)
				updatePoolOpt.SetName("updatedtestpool")
				updatePoolOpt.SetDescription("updating testPool")
				updatePoolResult, updatePoolResponse, updatePoolErr := service.UpdatePool(updatePoolOpt)
				Expect(updatePoolErr).To(BeNil())
				Expect(updatePoolResponse).ToNot(BeNil())
				Expect(updatePoolResult).ToNot(BeNil())
				Expect(updatePoolResponse.GetStatusCode()).To(BeEquivalentTo(200))
				Expect(*updatePoolResult.Name).To(BeEquivalentTo("updatedtestpool"))
				Expect(*updatePoolResult.Description).To(BeEquivalentTo("updating testPool"))

				//Test Create Load Balancer
				CreateLoadBalancerOptions := service.NewCreateLoadBalancerOptions(instanceID, *zoneInfo.ID)
				CreateLoadBalancerOptions.SetName("testloadbalancer")
				CreateLoadBalancerOptions.SetDescription("PDNS Load balancer")
				CreateLoadBalancerOptions.SetEnabled(true)
				CreateLoadBalancerOptions.SetTTL(120)
				CreateLoadBalancerOptions.SetFallbackPool(*resultPool.ID)
				CreateLoadBalancerOptions.SetDefaultPools([]string{*resultPool.ID})
				resultLoadbalancer, responseLoadbalancer, reqErrLoadbalancer := service.CreateLoadBalancer(CreateLoadBalancerOptions)
				Expect(reqErrLoadbalancer).To(BeNil())
				Expect(responseLoadbalancer).ToNot(BeNil())
				Expect(resultLoadbalancer).ToNot(BeNil())
				Expect(responseLoadbalancer.GetStatusCode()).To(BeEquivalentTo(200))
				Expect(*resultLoadbalancer.Name).To(ContainSubstring("testloadbalancer"))
				Expect(*resultLoadbalancer.Description).To(BeEquivalentTo("PDNS Load balancer"))
				Expect(*resultLoadbalancer.Enabled).To(BeEquivalentTo(true))
				Expect(*resultLoadbalancer.TTL).To(BeEquivalentTo(120))
				Expect(*resultLoadbalancer.FallbackPool).To(BeEquivalentTo(*resultPool.ID))

				//Test Get Load Balancer
				getLoadBalancerOpt := service.NewGetLoadBalancerOptions(instanceID, *zoneInfo.ID, *resultLoadbalancer.ID)
				getLoadBalancerResult, getLoadBalancerResponse, getLoadBalancerErr := service.GetLoadBalancer(getLoadBalancerOpt)
				Expect(getLoadBalancerErr).To(BeNil())
				Expect(getLoadBalancerResponse).ToNot(BeNil())
				Expect(getLoadBalancerResult).ToNot(BeNil())
				Expect(getLoadBalancerResponse.GetStatusCode()).To(BeEquivalentTo(200))
				Expect(*getLoadBalancerResult.Name).To(ContainSubstring("testloadbalancer"))
				Expect(*getLoadBalancerResult.Description).To(BeEquivalentTo("PDNS Load balancer"))
				Expect(*getLoadBalancerResult.Enabled).To(BeEquivalentTo(true))
				Expect(*getLoadBalancerResult.TTL).To(BeEquivalentTo(120))
				Expect(*getLoadBalancerResult.FallbackPool).To(BeEquivalentTo(*resultPool.ID))

				//Test Update Load Balancer
				updateLoadBalancerOpt := service.NewUpdateLoadBalancerOptions(instanceID, *zoneInfo.ID, *resultLoadbalancer.ID)
				updateLoadBalancerOpt.SetName("updateLoadBalancer")
				updateLoadBalancerOpt.SetDescription("updating Load Balancer")
				updateLoadBalancerResult, updateLoadBalancerResponse, updateLoadBalancerErr := service.UpdateLoadBalancer(updateLoadBalancerOpt)
				Expect(updateLoadBalancerErr).To(BeNil())
				Expect(updateLoadBalancerResponse).ToNot(BeNil())
				Expect(updateLoadBalancerResult).ToNot(BeNil())
				Expect(updateLoadBalancerResponse.GetStatusCode()).To(BeEquivalentTo(200))
				Expect(*updateLoadBalancerResult.Name).To(ContainSubstring("updateLoadBalancer"))
				Expect(*updateLoadBalancerResult.Description).To(BeEquivalentTo("updating Load Balancer"))
				Expect(*updateLoadBalancerResult.Enabled).To(BeEquivalentTo(true))

				//Test List Load Balancer
				listLoadBalancerOpt := service.NewListLoadBalancersOptions(instanceID, *zoneInfo.ID)
				listLoadBalancerResult, listLoadBalancerResponse, listLoadBalancerErr := service.ListLoadBalancers(listLoadBalancerOpt)
				Expect(listLoadBalancerErr).To(BeNil())
				Expect(listLoadBalancerResponse).ToNot(BeNil())
				Expect(listLoadBalancerResult).ToNot(BeNil())
				Expect(listLoadBalancerResponse.GetStatusCode()).To(BeEquivalentTo(200))

				//Test DeleteLoadBalancer
				deleteLoadBalancerOpt := service.NewDeleteLoadBalancerOptions(instanceID, *zoneInfo.ID, *resultLoadbalancer.ID)
				deleteLoadBalancerResponse, deleteLoadBalancerErr := service.DeleteLoadBalancer(deleteLoadBalancerOpt)
				Expect(deleteLoadBalancerErr).To(BeNil())
				Expect(deleteLoadBalancerResponse).ToNot(BeNil())
				Expect(deleteLoadBalancerResponse.GetStatusCode()).To(BeEquivalentTo(204))

				//Test DeletePool
				deletePoolOpt := service.NewDeletePoolOptions(instanceID, *resultPool.ID)
				deletePoolResponse, deletePoolErr := service.DeletePool(deletePoolOpt)
				Expect(deletePoolErr).To(BeNil())
				Expect(deletePoolResponse).ToNot(BeNil())
				Expect(deletePoolResponse.GetStatusCode()).To(BeEquivalentTo(204))

				//Test DeleteMonitor
				deleteOpt := service.NewDeleteMonitorOptions(instanceID, *result.ID)
				deleteResponse, deleteErr := service.DeleteMonitor(deleteOpt)
				Expect(deleteErr).To(BeNil())
				Expect(deleteResponse).ToNot(BeNil())
				Expect(deleteResponse.GetStatusCode()).To(BeEquivalentTo(204))
			})
			It(`crate/list PDNS load balancer monitor and pool`, func() {
				shouldSkipTest()

				//Create and List Monitor
				for i := 1; i < 4; i++ {
					CreateMonitorOptions := service.NewCreateMonitorOptions(instanceID)
					CreateMonitorOptions.SetName("testaMonitor-" + strconv.Itoa(i))
					CreateMonitorOptions.SetType(CreateMonitorOptions_Type_Http)
					CreateMonitorOptions.SetExpectedCodes("200")
					result, response, reqErr := service.CreateMonitor(CreateMonitorOptions)
					Expect(reqErr).To(BeNil())
					Expect(response).ToNot(BeNil())
					Expect(result).ToNot(BeNil())
					Expect(response.GetStatusCode()).To(BeEquivalentTo(200))
					Expect(*result.Type).To(BeEquivalentTo(CreateMonitorOptions_Type_Http))
				}
				listMonitorOpt := service.NewListMonitorsOptions(instanceID)
				listMonitorResult, listMonitorResponse, listMonitorErr := service.ListMonitors(listMonitorOpt)
				Expect(listMonitorErr).To(BeNil())
				Expect(listMonitorResponse).ToNot(BeNil())
				Expect(listMonitorResult).ToNot(BeNil())
				Expect(listMonitorResponse.GetStatusCode()).To(BeEquivalentTo(200))

				//Create and List Pool
				for i := 1; i < 4; i++ {
					CreatePoolOptions := service.NewCreatePoolOptions(instanceID)
					CreatePoolOptions.SetName("testPool-" + strconv.Itoa(i))
					CreatePoolOptions.SetDescription("creating pool-" + strconv.Itoa(i))
					CreatePoolOptions.SetEnabled(true)
					CreatePoolOptions.SetHealthyOriginsThreshold(1)
					originInputModel := new(OriginInput)
					originInputModel.Name = core.StringPtr("app-server-1-" + strconv.Itoa(i))
					originInputModel.Description = core.StringPtr("description of the origin server-" + strconv.Itoa(i))
					originInputModel.Address = core.StringPtr("10.10.10.8")
					originInputModel.Enabled = core.BoolPtr(true)
					CreatePoolOptions.Origins = []OriginInput{*originInputModel}
					resultPool, responsePool, reqErrPool := service.CreatePool(CreatePoolOptions)
					Expect(reqErrPool).To(BeNil())
					Expect(responsePool).ToNot(BeNil())
					Expect(resultPool).ToNot(BeNil())
					Expect(responsePool.GetStatusCode()).To(BeEquivalentTo(200))
				}
				listPoolOpt := service.NewListPoolsOptions(instanceID)
				listPoolResult, listPoolResponse, listPoolErr := service.ListPools(listPoolOpt)
				Expect(listPoolErr).To(BeNil())
				Expect(listPoolResponse).ToNot(BeNil())
				Expect(listPoolResult).ToNot(BeNil())
				Expect(listPoolResponse.GetStatusCode()).To(BeEquivalentTo(200))
			})
		})
	})
})
