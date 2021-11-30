package responses

type ProductionRoutingOp struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
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
		} `json:"results"`
	} `json:"d"`
}
