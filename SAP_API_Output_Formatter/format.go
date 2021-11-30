package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-production-routing-reads/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library/logger"
)

func ConvertToProductionRouting(raw []byte, l *logger.Logger) *ProductionRouting {
	pm := &responses.ProductionRouting{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		l.Error(err)
		return nil
	}
	if len(pm.D.Results) == 0 {
		l.Error("Result data is not exist.")
		return nil
	}
	if len(pm.D.Results) > 1 {
		l.Error("raw data has too many Results. %d Results exist. expected only 1 Result. Use the first of Results array", len(pm.D.Results))
	}
	data := pm.D.Results[0]

	return &ProductionRouting{
		Product                      data.Product,
		Plant                        data.Plant,
		ProductionRoutingGroup       data.ProductionRoutingGroup,
		ProductionRouting            data.ProductionRouting,
		ProductionRoutingGroupDesc   data.ProductionRoutingGroup_desc,
		BillOfOperationsStatus       data.BillOfOperationsStatus,
		MinimumLotSizeQuantity       data.MinimumLotSizeQuantity,
		MaximumLotSizeQuantity       data.MaximumLotSizeQuantity,
		ValidityStartDate            data.ValidityStartDate,
		ValidityEndDate              data.ValidityEndDate,
		IsMarkedForDeletion          data.IsMarkedForDeletion,
		ProdnRtgOpBOMItemInternalID  data.ProdnRtgOpBOMItemInternalID,
		ProductionRoutingOpIntID     data.ProductionRoutingOpIntID,
		BillOfMaterial               data.BillOfMaterial,
		BillOfMaterialItemNodeNumber data.BillOfMaterialItemNodeNumber,
		MatlCompIsMarkedForBackflush data.MatlCompIsMarkedForBackflush,
		WorkCenterInternalID         data.WorkCenterInternalID,
		Operation                    data.Operation,
		OperationText                data.OperationText,
		PurchasingInfoRecord         data.PurchasingInfoRecord,
		MaterialGroup                data.MaterialGroup,
		PurchasingGroup              data.PurchasingGroup,
		Supplier                     data.Supplier,
		PlannedDeliveryDuration      data.PlannedDeliveryDuration,
		NumberOfOperationPriceUnits  data.NumberOfOperationPriceUnits,
		OpExternalProcessingPrice    data.OpExternalProcessingPrice,
	}
}
