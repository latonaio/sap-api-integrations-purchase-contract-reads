package sap_api_caller

import (
	"fmt"
	"io/ioutil"
	sap_api_output_formatter "sap-api-integrations-purchase-contract-reads/SAP_API_Output_Formatter"
	"strings"
	"sync"

	sap_api_request_client_header_setup "github.com/latonaio/sap-api-request-client-header-setup"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
)

type SAPAPICaller struct {
	baseURL         string
	sapClientNumber string
	requestClient   *sap_api_request_client_header_setup.SAPRequestClient
	log             *logger.Logger
}

func NewSAPAPICaller(baseUrl, sapClientNumber string, requestClient *sap_api_request_client_header_setup.SAPRequestClient, l *logger.Logger) *SAPAPICaller {
	return &SAPAPICaller{
		baseURL:         baseUrl,
		requestClient:   requestClient,
		sapClientNumber: sapClientNumber,
		log:             l,
	}
}

func (c *SAPAPICaller) AsyncGetPurchaseContract(purchaseContract, purchaseContractItem string, accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "Header":
			func() {
				c.Header(purchaseContract)
				wg.Done()
			}()
		case "Item":
			func() {
				c.Item(purchaseContract, purchaseContractItem)
				wg.Done()
			}()
		default:
			wg.Done()
		}
	}

	wg.Wait()
}

func (c *SAPAPICaller) Header(purchaseContract string) {
	headerData, err := c.callPurchaseContractSrvAPIRequirementHeader("A_PurchaseContract", purchaseContract)
	if err != nil {
		c.log.Error(err)
	} else {
		c.log.Info(headerData)
	}

	itemData, err := c.callToItem(headerData[0].ToItem)
	if err != nil {
		c.log.Error(err)
	} else {
		c.log.Info(itemData)
	}

	itemAddressData, err := c.callToItemAddress(itemData[0].ToItemAddress)
	if err != nil {
		c.log.Error(err)
	} else {
		c.log.Info(itemAddressData)
	}

	itemConditionData, err := c.callToItemCondition(itemData[0].ToItemCondition)
	if err != nil {
		c.log.Error(err)
	} else {
		c.log.Info(itemConditionData)
	}
	return

}

func (c *SAPAPICaller) callPurchaseContractSrvAPIRequirementHeader(api, purchaseContract string) ([]sap_api_output_formatter.Header, error) {
	url := strings.Join([]string{c.baseURL, "API_PURCHASECONTRACT_PROCESS_SRV", api}, "/")
	param := c.getQueryWithHeader(map[string]string{}, purchaseContract)

	resp, err := c.requestClient.Request("GET", url, param, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToHeader(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callToItem(url string) ([]sap_api_output_formatter.ToItem, error) {
	resp, err := c.requestClient.Request("GET", url, map[string]string{}, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToToItem(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callToItemAddress(url string) ([]sap_api_output_formatter.ToItemAddress, error) {
	resp, err := c.requestClient.Request("GET", url, map[string]string{}, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToToItemAddress(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callToItemCondition(url string) ([]sap_api_output_formatter.ToItemCondition, error) {
	resp, err := c.requestClient.Request("GET", url, map[string]string{}, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToToItemCondition(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) Item(purchaseContract, purchaseContractItem string) {
	itemData, err := c.callPurchaseContractSrvAPIRequirementItem("A_PurchaseContractItem", purchaseContract, purchaseContractItem)
	if err != nil {
		c.log.Error(err)
	} else {
		c.log.Info(itemData)
	}

	itemAddressData, err := c.callToItemAddress(itemData[0].ToItemAddress)
	if err != nil {
		c.log.Error(err)
	} else {
		c.log.Info(itemAddressData)
	}

	itemConditionData, err := c.callToItemCondition(itemData[0].ToItemCondition)
	if err != nil {
		c.log.Error(err)
	} else {
		c.log.Info(itemConditionData)
	}
	return
}

func (c *SAPAPICaller) callPurchaseContractSrvAPIRequirementItem(api, purchaseContract, purchaseContractItem string) ([]sap_api_output_formatter.Item, error) {
	url := strings.Join([]string{c.baseURL, "API_PURCHASECONTRACT_PROCESS_SRV", api}, "/")

	param := c.getQueryWithItem(map[string]string{}, purchaseContract, purchaseContractItem)

	resp, err := c.requestClient.Request("GET", url, param, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToItem(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) getQueryWithHeader(params map[string]string, purchaseContract string) map[string]string {
	if len(params) == 0 {
		params = make(map[string]string, 1)
	}
	params["$filter"] = fmt.Sprintf("PurchaseContract eq '%s'", purchaseContract)
	return params
}

func (c *SAPAPICaller) getQueryWithItem(params map[string]string, purchaseContract, purchaseContractItem string) map[string]string {
	if len(params) == 0 {
		params = make(map[string]string, 1)
	}
	params["$filter"] = fmt.Sprintf("PurchaseContract eq '%s' and PurchaseContractItem eq '%s'", purchaseContract, purchaseContractItem)
	return params
}
