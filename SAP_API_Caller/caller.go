package sap_api_caller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"

	"github.com/latonaio/golang-logging-library/logger"
)

type SAPAPICaller struct {
	baseURL string
	apiKey  string
	log     *logger.Logger
}

func NewSAPAPICaller(baseUrl string, l *logger.Logger) *SAPAPICaller {
	return &SAPAPICaller{
		baseURL: baseUrl,
		apiKey:  GetApiKey(),
		log:     l,
	}
}

		
func (c *SAPAPICaller) AsyncGetProductionRouting(Product, Plant, ValidityEndDate, ProductionRoutingGroup, ProductionRouting, ProductionRoutingOpIntID string) {
	wg := &sync.WaitGroup{}

	wg.Add(3)
	go func() {
		c.ProductPlant(Product, Plant, ValidityEndDate)
		wg.Done()
	}()
	go func() {
		c.ProductionRouting(Product, Plant, ProductionRoutingGroup, ProductionRouting, ValidityEndDate)
		wg.Done()
	}()
	go func() {
		c.ProductionRoutingOpIntID(Product, Plant, ProductionRoutingGroup, ProductionRouting, ProductionRoutingOpIntID, ValidityEndDate)
		wg.Done()
	}()
	wg.Wait()
}

func (c *SAPAPICaller) ProductPlant(Product, Plant, ValidityEndDate string) {
	res, err := c.callProductionRoutingSrvAPIRequirementProductPlant("ProductionRoutingHeader", Product, Plant, ValidityEndDate)
	if err != nil {
		c.log.Error(err)
		return
	}

	c.log.Info(res)

}

func (c *SAPAPICaller) ProductionRouting(Product, Plant, ProductionRoutingGroup, ProductionRouting, ValidityEndDate string) {
	res, err := c.callProductionRoutingSrvAPIRequirementProductionRouting("OpDocumentPRTAssgmt", Product, Plant, ProductionRoutingGroup, ProductionRouting, ValidityEndDate)
	if err != nil {
		c.log.Error(err)
		return
	}

	c.log.Info(res)

}

func (c *SAPAPICaller) ProductionRoutingOpIntID(Product, Plant, ProductionRoutingGroup, ProductionRouting, ProductionRoutingOpIntID, ValidityEndDate string) {
	res, err := c.callProductionRoutingSrvAPIRequirementOpIntID("OpDocumentPRTAssgmt", Product, Plant, ProductionRoutingGroup, ProductionRouting, ProductionRoutingOpIntID, ValidityEndDate)
	if err != nil {
		c.log.Error(err)
		return
	}

	c.log.Info(res)

}

func (c *SAPAPICaller) callProductionRoutingSrvAPIRequirementProductPlant(api, Product, Plant, ValidityEndDate string) ([]byte, error) {
	url := strings.Join([]string{c.baseURL, "API_PRODUCTION_ROUTING", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	params := req.URL.Query()
	// params.Add("$select", "Product, Plant, ValidityEndDate")
	params.Add("$filter", fmt.Sprintf("Product eq '%s'and Plant eq '%s' and ValidityEndDate eq '%s'", Product, Plant, ValidityEndDate))
	req.URL.RawQuery = params.Encode()

	req.Header.Set("APIKey", c.apiKey)
	req.Header.Set("Accept", "application/json")

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	return byteArray, nil
}

func (c *SAPAPICaller) callProductionRoutingSrvAPIRequirementProductionRouting(api, Product, Plant, ProductionRoutingGroup, ProductionRouting, ValidityEndDate string) ([]byte, error) {
	url := strings.Join([]string{c.baseURL, "API_PRODUCTION_ROUTING", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	params := req.URL.Query()
	// params.Add("$select", "Product, Plant, ProductionRoutingGroup, ProductionRouting, ValidityEndDate")
	params.Add("$filter", fmt.Sprintf("Product eq '%s' and Plant eq '%s' and ProductionRoutingGroup eq '%s' and ProductionRouting eq '%s' and ValidityEndDate eq '%s'", Product, Plant, ProductionRoutingGroup, ProductionRouting, ValidityEndDate))
	req.URL.RawQuery = params.Encode()

	req.Header.Set("APIKey", c.apiKey)
	req.Header.Set("Accept", "application/json")

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	return byteArray, nil
}

func (c *SAPAPICaller) callProductionRoutingSrvAPIRequirementOpIntID(api, Product, Plant, ProductionRoutingGroup, ProductionRouting, ProductionRoutingOpIntID, ValidityEndDate string) ([]byte, error) {
	url := strings.Join([]string{c.baseURL, "API_PRODUCTION_ROUTING", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	params := req.URL.Query()
	// params.Add("$select", "Product, Plant, ProductionRoutingGroup, ProductionRouting, ProductionRoutingOpIntID, ValidityEndDate")
	params.Add("$filter", fmt.Sprintf("Product eq '%s' and Plant eq '%s' and ProductionRoutingGroup eq '%s' and ProductionRouting eq '%s' and ProductionRoutingOpIntID eq '%s' and ValidityEndDate eq '%s'", Product, Plant, ProductionRoutingGroup, ProductionRouting, ProductionRoutingOpIntID, ValidityEndDate))
	req.URL.RawQuery = params.Encode()

	req.Header.Set("APIKey", c.apiKey)
	req.Header.Set("Accept", "application/json")

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	return byteArray, nil
}