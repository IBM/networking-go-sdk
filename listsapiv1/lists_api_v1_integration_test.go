/**
 * (C) Copyright IBM Corp. 2025.
 */

package listsapiv1_test

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/networking-go-sdk/listsapiv1"
	"github.com/joho/godotenv"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const configFile = "../cis.env"

var configLoaded bool = true
var authenticationSucceeded bool = true

func shouldSkipTest() {
	//Skip("Skipping...")

	if !configLoaded {
		Skip("External configuration is not available, skipping...")
	}

	if !authenticationSucceeded {
		Skip("Authentication failed. Check external configuration...")
	}
}

var _ = Describe(`ListsApiV1 Integration Tests`, func() {

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

	var listsService *listsapiv1.ListsApiV1
	serviceURL := os.Getenv("API_ENDPOINT")
	crn := os.Getenv("CRN")
	var listId string
	var itemId string
	var operationId string

	Describe("Client initialization", func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It("Successfully construct the service client instance", func() {
			listsServiceOptions := &listsapiv1.ListsApiV1Options{
				ServiceName:   "listsapi",
				URL:           serviceURL,
				Crn:           &crn,
				Authenticator: authenticator,
				ListID:        &listId,
				ItemID:        &itemId,
				OperationID:   &operationId,
			}

			listsService, err = listsapiv1.NewListsApiV1(listsServiceOptions)
			Expect(err).To(BeNil())
			Expect(listsService).ToNot(BeNil())
			Expect(listsService.Service.Options.URL).To(Equal(serviceURL))

			listsService.EnableRetries(4, 30*time.Second)
		})
	})

	Describe(`Managed Lists`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`List Managed Lists`, func() {
			getManagedListsOptions := &listsapiv1.GetManagedListsOptions{}

			managedListsResp, response, err := listsService.GetManagedLists(getManagedListsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(managedListsResp).ToNot(BeNil())
		})
	})

	Describe(`Custom Lists`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`Create ip List`, func() {
			createCustomListsOptions := &listsapiv1.CreateCustomListsOptions{
				Kind:        core.StringPtr("ip"),
				Name:        core.StringPtr("ip_list"),
				Description: core.StringPtr("ip list"),
			}

			customListResp, response, err := listsService.CreateCustomLists(createCustomListsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(customListResp).ToNot(BeNil())
			listsService.ListID = customListResp.Result.ID

		})

		It(`Get Custom List`, func() {
			getCustomListOptions := &listsapiv1.GetCustomListOptions{}

			customListResp, response, err := listsService.GetCustomList(getCustomListOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(customListResp).ToNot(BeNil())
		})

		It(`Update Custom List`, func() {
			updateCustomListOptions := &listsapiv1.UpdateCustomListOptions{
				Description: core.StringPtr("change description"),
			}

			customListResp, response, err := listsService.UpdateCustomList(updateCustomListOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(customListResp).ToNot(BeNil())
		})

		It(`Get Custom Lists`, func() {
			getCustomListsOptions := &listsapiv1.GetCustomListsOptions{}

			customListResp, response, err := listsService.GetCustomLists(getCustomListsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(customListResp).ToNot(BeNil())
		})

		It(`Create List Item`, func() {
			createListItemsReqItemModel := &listsapiv1.CreateListItemsReqItem{
				Comment: core.StringPtr("list of IPs."),
				Ip:      core.StringPtr("172.64.0.0"),
			}

			createListItemsOptions := &listsapiv1.CreateListItemsOptions{
				CreateListItemsReqItem: []listsapiv1.CreateListItemsReqItem{*createListItemsReqItemModel},
			}

			listOperationResp, response, err := listsService.CreateListItems(createListItemsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(listOperationResp).ToNot(BeNil())
			listsService.OperationID = listOperationResp.Result.OperationID
		})

		It(`Get Operation Status(`, func() {
			getOperationStatusOptions := &listsapiv1.GetOperationStatusOptions{}

			operationStatusResp, response, err := listsService.GetOperationStatus(getOperationStatusOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(operationStatusResp).ToNot(BeNil())
		})

		It(`Get List Items`, func() {
			getListItemsOptions := &listsapiv1.GetListItemsOptions{}

			listItemsResp, response, err := listsService.GetListItems(getListItemsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(listItemsResp).ToNot(BeNil())
			listsService.ItemID = listItemsResp.Result[0].ID
		})

		It(`Get List Item`, func() {
			getListItemOptions := &listsapiv1.GetListItemOptions{}

			listItemResp, response, err := listsService.GetListItem(getListItemOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(listItemResp).ToNot(BeNil())
		})

		It(`Update List Items`, func() {
			createListItemsReqItemModel := &listsapiv1.CreateListItemsReqItem{
				Comment: core.StringPtr("list of IPs."),
				Ip:      core.StringPtr("172.64.0.1"),
			}

			updateListItemsOptions := &listsapiv1.UpdateListItemsOptions{
				CreateListItemsReqItem: []listsapiv1.CreateListItemsReqItem{*createListItemsReqItemModel},
			}

			listOperationResp, response, err := listsService.UpdateListItems(updateListItemsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(listOperationResp).ToNot(BeNil())
		})

		It(`Delete List Items`, func() {
			deleteListItemsReqItemsItemModel := &listsapiv1.DeleteListItemsReqItemsItem{
				ID: listsService.ItemID,
			}

			deleteListItemsOptions := &listsapiv1.DeleteListItemsOptions{
				Items: []listsapiv1.DeleteListItemsReqItemsItem{*deleteListItemsReqItemsItemModel},
			}

			listOperationResp, response, err := listsService.DeleteListItems(deleteListItemsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(listOperationResp).ToNot(BeNil())
		})

		It(`Delete Custom List`, func() {
			deleteCustomListsOptions := &listsapiv1.DeleteCustomListOptions{}

			customListResp, response, err := listsService.DeleteCustomList(deleteCustomListsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(customListResp).ToNot(BeNil())
		})

	})

})
