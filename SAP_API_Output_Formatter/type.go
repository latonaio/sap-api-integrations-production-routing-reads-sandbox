package sap_api_output_formatter

type ProductionRoutingReads struct {
	ConnectionKey string `json:"connection_key"`
	Result        bool   `json:"result"`
	RedisKey      string `json:"redis_key"`
	Filepath      string `json:"filepath"`
	APISchema     string `json:"api_schema"`
	MaterialCode  string `json:"material_code"`
	Plant         string `json:"plant"`
	Deleted       bool   `json:"deleted"`
}

type ProductionRouting struct {
	Product                     string `json:"Product"`
	Plant                       string `json:"Plant"`
	ProductionRoutingGroup      string `json:"ProductionRoutingGroup"`
	ProductionRouting           string `json:"ProductionRouting"`
	ProductionRoutingGroupDesc  string `json:"ProductionRoutingGroup_desc"`
	BillOfOperationsStatus      string `json:"BillOfOperationsStatus"`
	MinimumLotSizeQuantity      string `json:"MinimumLotSizeQuantity"`
	MaximumLotSizeQuantity      string `json:"MaximumLotSizeQuantity"`
	ValidityStartDate           string `json:"ValidityStartDate"`
	ValidityEndDate             string `json:"ValidityEndDate"`
	IsMarkedForDeletion         bool   `json:"IsMarkedForDeletion"`
	ProdnRtgOpBOMItemInternalID string `json:"ProdnRtgOpBOMItemInternalID"`
}

type ProductionRoutingOp struct {
	Product                      string `json:"Product"`
	Plant                        string `json:"Plant"`
	ProductionRoutingGroup       string `json:"ProductionRoutingGroup"`
	ProductionRouting            string `json:"ProductionRouting"`
	ProductionRoutingOpIntID     string `json:"ProductionRoutingOpIntID"`
	BillOfMaterial               string `json:"BillOfMaterial"`
	BillOfMaterialItemNodeNumber string `json:"BillOfMaterialItemNodeNumber"`
	MatlCompIsMarkedForBackflush bool   `json:"MatlCompIsMarkedForBackflush"`
	WorkCenterInternalID         string `json:"WorkCenterInternalID"`
	Operation                    string `json:"Operation"`
	OperationText                string `json:"OperationText"`
	PurchasingInfoRecord         string `json:"PurchasingInfoRecord"`
	MaterialGroup                string `json:"MaterialGroup"`
	PurchasingGroup              string `json:"PurchasingGroup"`
	Supplier                     string `json:"Supplier"`
	PlannedDeliveryDuration      string `json:"PlannedDeliveryDuration"`
	NumberOfOperationPriceUnits  string `json:"NumberOfOperationPriceUnits"`
	OpExternalProcessingPrice    string `json:"OpExternalProcessingPrice"`
}
