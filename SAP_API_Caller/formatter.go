package sap_api_caller

type ProductionRoutingReads struct {
	ConnectionKey     string `json:"connection_key"`
	Result            bool   `json:"result"`
	RedisKey          string `json:"redis_key"`
	Filepath          string `json:"filepath"`
	APISchema         string `json:"api_schema"`
	Material          string `json:"material_code"`
	Plant             string `json:"plant"`
	Deleted           string `json:"deleted"`
}
    
type ProductionRouting struct {
    Product                     string `json:"Product"`
	Plant                       string `json:"Plant"`
	ProductionRoutingGroup      string `json:"ProductionRoutingGroup"`
	ProductionRouting           string `json:"ProductionRouting"`
	ProductionRoutingGroupDesc  string `json:"ProductionRoutingGroup_desc"`
	BillOfOperationsStatus      string `json:"BillOfOperationsStatus"`
	MinimumLotSizeQuantity      float64 `json:"MinimumLotSizeQuantity"`
	MaximumLotSizeQuantity      float64 `json:"MaximumLotSizeQuantity"`
	ValidityStartDate           string `json:"ValidityStartDate"`
	ValidityEndDate             string `json:"ValidityEndDate"`
	IsMarkedForDeletion         string `json:"IsMarkedForDeletion"`
	ProdnRtgOpBOMItemInternalID string `json:"ProdnRtgOpBOMItemInternalID"`
}   

type ProductionRoutingOpIntID    struct {
    Product                      string `json:"Product"`
	Plant                        string `json:"Plant"`
	ProductionRoutingGroup       string `json:"ProductionRoutingGroup"`
	ProductionRouting            string `json:"ProductionRouting"`
	ProductionRoutingOpIntID     int    `json:"ProductionRoutingOpIntID"`
    BillOfMaterial               string `json:"BillOfMaterial"`
    BillOfMaterialItemNodeNumber int    `json:"BillOfMaterialItemNodeNumber"`
    MatlCompIsMarkedForBackflush string `json:"MatlCompIsMarkedForBackflush"`
    WorkCenterInternalID         int    `json:"WorkCenterInternalID"`
    Operation                    string `json:"Operation"`
    OperationText                string `json:"OperationText"`
    PurchasingInfoRecord         string `json:"PurchasingInfoRecord"`
    MaterialGroup                string `json:"MaterialGroup"`
    PurchasingGroup              string `json:"PurchasingGroup"`
    Supplier                     string `json:"Supplier"`
    PlannedDeliveryDuration      int    `json:"PlannedDeliveryDuration"`
    NumberOfOperationPriceUnits  int    `json:"NumberOfOperationPriceUnits"`
    OpExternalProcessingPrice    float64 `json:"OpExternalProcessingPrice"`
}