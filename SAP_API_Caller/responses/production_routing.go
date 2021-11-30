package responses

type ProductionRouting struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
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
		} `json:"results"`
	} `json:"d"`
}
