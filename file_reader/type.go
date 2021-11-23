package file_reader

type EC_MC struct {
	ConnectionKey string `json:"connection_key"`
	Result        bool   `json:"result"`
	RedisKey      string `json:"redis_key"`
	Filepath      string `json:"filepath"`
	Document    struct {
		DocumentNo                string      `json:"document_no"`
		DeliverTo                 string      `json:"deliver_to"`
		Quantity                  float64     `json:"quantity"`
		PickedQuantity            float64     `json:"picked_quantity"`
		Price                     float64     `json:"price"`
	    Batch                     string      `json:"batch"`
	} `json:"document"`
	ProductionOrder struct {
		DocumentNo           string      `json:"document_no"`
		Status               string      `json:"status"`
		DeliverTo            string      `json:"deliver_to"`
		Quantity             float64     `json:"quantity"`
		CompletedQuantity    float64     `json:"completed_quantity"`
	    PlannedStartDate     string      `json:"planned_start_date"`
	    PlannedValidatedDate string      `json:"planned_validated_date"`
	    ActualStartDate      string      `json:"actual_start_date"`
	    ActualValidatedDate  string      `json:"actual_validated_date"`
	    Batch                string      `json:"batch"`
		Work              struct {
			WorkNo                   string      `json:"work_no"`
			Quantity                 float64     `json:"quantity"`
			CompletedQuantity        float64     `json:"completed_quantity"`
			ErroredQuantity          float64     `json:"errored_quantity"`
			Component                string      `json:"component"`
			PlannedComponentQuantity float64     `json:"planned_component_quantity"`
			PlannedStartDate         string      `json:"planned_start_date"`
			PlannedStartTime         string      `json:"planned_start_time"`
			PlannedValidatedDate     string      `json:"planned_validated_date"`
			PlannedValidatedTime     string      `json:"planned_validated_time"`
			ActualStartDate          string      `json:"actual_start_date"`
			ActualStartTime          string      `json:"actual_start_time"`
			ActualValidatedDate      string      `json:"actual_validated_date"`
			ActualValidatedTime      string      `json:"actual_validated_time"`
		} `json:"work"`
	} `json:"production_order"`
	APISchema                         string `json:"api_schema"`
	Product                           string `json:"material_code"`
	Plant                             string `json:"plant/supplier"`
	Stock                             float64 `json:"stock"`
	DocumentType                      string `json:"document_type"`
	DocumentNo                        string `json:"document_no"`
	PlannedDate                       string `json:"planned_date"`
	ValidatedDate                     string `json:"validated_date"`
	Deleted                           string `json:"deleted"`
}

type SDC struct {
	ConnectionKey     string `json:"connection_key"`
	Result            bool   `json:"result"`
	RedisKey          string `json:"redis_key"`
	Filepath          string `json:"filepath"`
	ProductionRouting struct {
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
		ProductionRoutingOpIntID    struct {
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
		} `json:"ProductionRoutingOpIntID"`
	} `json:"ProductionRouting"`
	APISchema    string `json:"api_schema"`
	Material     string `json:"material_code"`
	Plant        string `json:"plant"`
	Deleted      string `json:"deleted"`
}