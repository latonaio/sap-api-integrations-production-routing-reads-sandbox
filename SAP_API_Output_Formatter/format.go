package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-production-routing-reads/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library/logger"
	"golang.org/x/xerrors"
)

func ConvertToHeader(raw []byte, l *logger.Logger) ([]Header, error) {
	pm := &responses.Header{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Header. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	header := make([]Header, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		header = append(header, Header{
			ProductionRoutingGroup:        data.ProductionRoutingGroup,
			ProductionRouting:             data.ProductionRouting,
			ProductionRoutingInternalVers: data.ProductionRoutingInternalVers,
			IsMarkedForDeletion:           data.IsMarkedForDeletion,
			BillOfOperationsDesc:          data.BillOfOperationsDesc,
			Plant:                         data.Plant,
			BillOfOperationsUsage:         data.BillOfOperationsUsage,
			BillOfOperationsStatus:        data.BillOfOperationsStatus,
			ResponsiblePlannerGroup:       data.ResponsiblePlannerGroup,
			MinimumLotSizeQuantity:        data.MinimumLotSizeQuantity,
			MaximumLotSizeQuantity:        data.MaximumLotSizeQuantity,
			BillOfOperationsUnit:          data.BillOfOperationsUnit,
			CreationDate:                  data.CreationDate,
			CreatedByUser:                 data.CreatedByUser,
			LastChangeDate:                data.LastChangeDate,
			ValidityStartDate:             data.ValidityStartDate,
			ValidityEndDate:               data.ValidityEndDate,
			ChangeNumber:                  data.ChangeNumber,
			PlainLongText:                 data.PlainLongText,
			ToMatlAssgmt:                  data.ToMatlAssgmt.Deferred.URI,
			ToSequence:                    data.ToSequence.Deferred.URI,
		})
	}

	return header, nil
}

func ConvertToMaterialAssignment(raw []byte, l *logger.Logger) ([]MaterialAssignment, error) {
	pm := &responses.MaterialAssignment{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to MaterialAssignment. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	materialAssignment := make([]MaterialAssignment, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		materialAssignment = append(materialAssignment, MaterialAssignment{
			Product:                        data.Product,
			Plant:                          data.Plant,
			ProductionRoutingGroup:         data.ProductionRoutingGroup,
			ProductionRouting:              data.ProductionRouting,
			ProductionRoutingMatlAssgmt:    data.ProductionRoutingMatlAssgmt,
			ProductionRtgMatlAssgmtIntVers: data.ProductionRtgMatlAssgmtIntVers,
			CreationDate:                   data.CreationDate,
			LastChangeDate:                 data.LastChangeDate,
			ValidityStartDate:              data.ValidityStartDate,
			ValidityEndDate:                data.ValidityEndDate,
			ChangeNumber:                   data.ChangeNumber,
		})
	}

	return materialAssignment, nil
}

func ConvertToToMatlAssgmt(raw []byte, l *logger.Logger) (*ToMatlAssgmt, error) {
	pm := &responses.ToMatlAssgmt{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ToMatlAssgmt. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. expected only 1 Result. Use the first of Results array", len(pm.D.Results))
	}
	data := pm.D.Results[0]

	return &ToMatlAssgmt{
			Product:                        data.Product,
			Plant:                          data.Plant,
			ProductionRoutingGroup:         data.ProductionRoutingGroup,
			ProductionRouting:              data.ProductionRouting,
			ProductionRoutingMatlAssgmt:    data.ProductionRoutingMatlAssgmt,
			ProductionRtgMatlAssgmtIntVers: data.ProductionRtgMatlAssgmtIntVers,
			CreationDate:                   data.CreationDate,
			LastChangeDate:                 data.LastChangeDate,
			ValidityStartDate:              data.ValidityStartDate,
			ValidityEndDate:                data.ValidityEndDate,
			ChangeNumber:                   data.ChangeNumber,
	}, nil
}

func ConvertToToSequence(raw []byte, l *logger.Logger) (*ToSequence, error) {
	pm := &responses.ToSequence{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ToSequence. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. expected only 1 Result. Use the first of Results array", len(pm.D.Results))
	}
	data := pm.D.Results[0]

	return &ToSequence{
			ProductionRoutingGroup:       data.ProductionRoutingGroup,
			ProductionRouting:            data.ProductionRouting,
			ProductionRoutingSequence:    data.ProductionRoutingSequence,
			ProductionRoutingSqncIntVers: data.ProductionRoutingSqncIntVers,
			ChangeNumber:                 data.ChangeNumber,
			ValidityStartDate:            data.ValidityStartDate,
			ValidityEndDate:              data.ValidityEndDate,
			SequenceCategory:             data.SequenceCategory,
			BillOfOperationsRefSequence:  data.BillOfOperationsRefSequence,
			MinimumLotSizeQuantity:       data.MinimumLotSizeQuantity,
			MaximumLotSizeQuantity:       data.MaximumLotSizeQuantity,
			BillOfOperationsUnit:         data.BillOfOperationsUnit,
			SequenceText:                 data.SequenceText,
			CreationDate:                 data.CreationDate,
			LastChangeDate:               data.LastChangeDate,
			ToOperation:                  data.ToOperation.Deferred.URI,
	}, nil
}

func ConvertToToOperation(raw []byte, l *logger.Logger) (*ToOperation, error) {
	pm := &responses.ToOperation{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ToOperation. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. expected only 1 Result. Use the first of Results array", len(pm.D.Results))
	}
	data := pm.D.Results[0]

	return &ToOperation{
			ProductionRoutingGroup:         data.ProductionRoutingGroup,
			ProductionRouting:              data.ProductionRouting,
			ProductionRoutingSequence:      data.ProductionRoutingSequence,
			ProductionRoutingOpIntID:       data.ProductionRoutingOpIntID,
			ProductionRoutingOpIntVersion:  data.ProductionRoutingOpIntVersion,
			Operation:                      data.Operation,
			CreationDate:                   data.CreationDate,
			LastChangeDate:                 data.LastChangeDate,
			ChangeNumber:                   data.ChangeNumber,
			ValidityStartDate:              data.ValidityStartDate,
			ValidityEndDate:                data.ValidityEndDate,
			OperationText:                  data.OperationText,
			LongTextLanguageCode:           data.LongTextLanguageCode,
			Plant:                          data.Plant,
			OperationControlProfile:        data.OperationControlProfile,
			OperationStandardTextCode:      data.OperationStandardTextCode,
			WorkCenterTypeCode:             data.WorkCenterTypeCode,
			WorkCenterInternalID:           data.WorkCenterInternalID,
			CapacityCategoryCode:           data.CapacityCategoryCode,
			OperationCostingRelevancyType:  data.OperationCostingRelevancyType,
			NumberOfTimeTickets:            data.NumberOfTimeTickets,
			NumberOfConfirmationSlips:      data.NumberOfConfirmationSlips,
			OperationSetupType:             data.OperationSetupType,
			OperationSetupGroupCategory:    data.OperationSetupGroupCategory,
			OperationSetupGroup:            data.OperationSetupGroup,
			OperationReferenceQuantity:     data.OperationReferenceQuantity,
			OperationUnit:                  data.OperationUnit,
			OpQtyToBaseQtyNmrtr:            data.OpQtyToBaseQtyNmrtr,
			OpQtyToBaseQtyDnmntr:           data.OpQtyToBaseQtyDnmntr,
			MaximumWaitDuration:            data.MaximumWaitDuration,
			MaximumWaitDurationUnit:        data.MaximumWaitDurationUnit,
			MinimumWaitDuration:            data.MinimumWaitDuration,
			MinimumWaitDurationUnit:        data.MinimumWaitDurationUnit,
			StandardQueueDuration:          data.StandardQueueDuration,
			StandardQueueDurationUnit:      data.StandardQueueDurationUnit,
			MinimumQueueDuration:           data.MinimumQueueDuration,
			MinimumQueueDurationUnit:       data.MinimumQueueDurationUnit,
			StandardMoveDuration:           data.StandardMoveDuration,
			StandardMoveDurationUnit:       data.StandardMoveDurationUnit,
			MinimumMoveDuration:            data.MinimumMoveDuration,
			MinimumMoveDurationUnit:        data.MinimumMoveDurationUnit,
			OpIsExtlyProcdWithSubcontrg:    data.OpIsExtlyProcdWithSubcontrg,
			PurchasingInfoRecord:           data.PurchasingInfoRecord,
			PurchasingOrganization:         data.PurchasingOrganization,
			PlannedDeliveryDuration:        data.PlannedDeliveryDuration,
			MaterialGroup:                  data.MaterialGroup,
			PurchasingGroup:                data.PurchasingGroup,
			Supplier:                       data.Supplier,
			NumberOfOperationPriceUnits:    data.NumberOfOperationPriceUnits,
			CostElement:                    data.CostElement,
			OpExternalProcessingPrice:      data.OpExternalProcessingPrice,
			OpExternalProcessingCurrency:   data.OpExternalProcessingCurrency,
			OperationScrapPercent:          data.OperationScrapPercent,
			ChangedDateTime:                data.ChangedDateTime,
			PlainLongText:                  data.PlainLongText,
	}, nil
}
