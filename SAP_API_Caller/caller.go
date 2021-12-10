package sap_api_caller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	sap_api_output_formatter "sap-api-integrations-production-routing-reads/SAP_API_Output_Formatter"
	"strings"
	"sync"

	"github.com/latonaio/golang-logging-library/logger"
	"golang.org/x/xerrors"
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

func (c *SAPAPICaller) AsyncGetProductionRouting(productionRoutingGroup, productionRouting, product, plant string, accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "Header":
			func() {
				c.Header(productionRoutingGroup, productionRouting)
				wg.Done()
			}()
		case "ProductPlant":
			func() {
				c.ProductPlant(product, plant)
				wg.Done()
			}()
		default:
			wg.Done()
		}
	}

	wg.Wait()
}

func (c *SAPAPICaller) Header(productionRoutingGroup, productionRouting string) {
	headerData, err := c.callProductionRoutingSrvAPIRequirementHeader("ProductionRoutingHeader", productionRoutingGroup, productionRouting)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(headerData)

	matlAssgmtData, err := c.callToMatlAssgmt(headerData[0].ToMatlAssgmt)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(matlAssgmtData)

	sequenceData, err := c.callToSequence(headerData[0].ToSequence)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(sequenceData)

	operationData, err := c.callToOperation(sequenceData.ToOperation)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(operationData)
}

func (c *SAPAPICaller) callProductionRoutingSrvAPIRequirementHeader(api, productionRoutingGroup, productionRouting string) ([]sap_api_output_formatter.Header, error) {
	url := strings.Join([]string{c.baseURL, "API_PRODUCTION_ROUTING", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithHeader(req, productionRoutingGroup, productionRouting)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToHeader(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callToMatlAssgmt(url string) (*sap_api_output_formatter.ToMatlAssgmt, error) {
	req, _ := http.NewRequest("GET", url, nil)
	c.setHeaderAPIKeyAccept(req)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToToMatlAssgmt(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callToSequence(url string) (*sap_api_output_formatter.ToSequence, error) {
	req, _ := http.NewRequest("GET", url, nil)
	c.setHeaderAPIKeyAccept(req)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToToSequence(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callToOperation(url string) (*sap_api_output_formatter.ToOperation, error) {
	req, _ := http.NewRequest("GET", url, nil)
	c.setHeaderAPIKeyAccept(req)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToToOperation(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) ProductPlant(product, plant string) {
	data, err := c.callProductionRoutingSrvAPIRequirementProductPlant("ProductionRoutingMatlAssgmt", product, plant)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)

	sequenceData, err := c.callToSequence(
		fmt.Sprintf("%s/API_PRODUCTION_ROUTING/ProductionRoutingHeader(ProductionRoutingGroup='%s',ProductionRouting='%s',ProductionRoutingInternalVers='%s')/to_Sequence",
			c.baseURL, data[0].ProductionRoutingGroup, data[0].ProductionRouting, data[0].ProductionRtgMatlAssgmtIntVers,
		),
	)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(sequenceData)

	operationData, err := c.callToOperation(sequenceData.ToOperation)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(operationData)

}

func (c *SAPAPICaller) callProductionRoutingSrvAPIRequirementProductPlant(api, product, plant string) ([]sap_api_output_formatter.MaterialAssignment, error) {
	url := strings.Join([]string{c.baseURL, "API_PRODUCTION_ROUTING", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithProductPlant(req, product, plant)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToMaterialAssignment(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) setHeaderAPIKeyAccept(req *http.Request) {
	req.Header.Set("APIKey", c.apiKey)
	req.Header.Set("Accept", "application/json")
}

func (c *SAPAPICaller) getQueryWithHeader(req *http.Request, productionRoutingGroup, productionRouting string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("ProductionRoutingGroup eq '%s' and ProductionRouting eq '%s'", productionRoutingGroup, productionRouting))
	req.URL.RawQuery = params.Encode()
}

func (c *SAPAPICaller) getQueryWithProductPlant(req *http.Request, product, plant string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("Product eq '%s' and Plant eq '%s'", product, plant))
	req.URL.RawQuery = params.Encode()
}
