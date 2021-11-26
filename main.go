package main
 
import (
    sap_api_caller "sap-api-integrations-production-routing-reads/SAP_API_Caller"
    "sap-api-integrations-production-routing-reads/file_reader"

    "github.com/latonaio/golang-logging-library/logger"
)

func main() {
    l := logger.NewLogger()
    fr := file_reader.NewFileReader()
    inoutSDC := fr.ReadSDC("./Inputs//SDC_Production_Routing_sample.json")
    caller := sap_api_caller.NewSAPAPICaller(
        "https://sandbox.api.sap.com/s4hanacloud/sap/opu/odata/sap/", l,
    )

    caller.AsyncGetProductionRouting(
        inoutSDC.Product, Plant, ValidityEndDate,
        inoutSDC.Product, Plant, ProductionRoutingGroup, ProductionRouting, ValidityEndDate,
		inoutSDC.Product, Plant, ProductionRoutingGroup, ProductionRouting, ValidityEndDate, ProductionRoutingOpIntID, ValidityEndDate,
    )
}