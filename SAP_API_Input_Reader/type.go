package sap_api_input_reader

type EC_MC struct {
	ConnectionKey string `json:"connection_key"`
	Result        bool   `json:"result"`
	RedisKey      string `json:"redis_key"`
	Filepath      string `json:"filepath"`
	PurchaseContract struct {
		PurchaseContract          string      `json:"document_no"`
		Plant                     string      `json:"deliver_to"`
		OrderQuantity             string      `json:"quantity"`
		PickedQuantity            string      `json:"picked_quantity"`
		NetPriceAmount            string      `json:"price"`
	    Batch                     string      `json:"batch"`
	} `json:"document"`
	ProductionOrder struct {
		DocumentNo           string      `json:"document_no"`
		Status               string      `json:"status"`
		DeliverTo            string      `json:"deliver_to"`
		Quantity             string      `json:"quantity"`
		CompletedQuantity    string      `json:"completed_quantity"`
	    PlannedStartDate     string      `json:"planned_start_date"`
	    PlannedValidatedDate string      `json:"planned_validated_date"`
	    ActualStartDate      string      `json:"actual_start_date"`
	    ActualValidatedDate  string      `json:"actual_validated_date"`
	    Batch                string      `json:"batch"`
		Work              struct {
			WorkNo                   string      `json:"work_no"`
			Quantity                 string      `json:"quantity"`
			CompletedQuantity        string      `json:"completed_quantity"`
			ErroredQuantity          string      `json:"errored_quantity"`
			Component                string      `json:"component"`
			PlannedComponentQuantity string      `json:"planned_component_quantity"`
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
	Accepter                        []string `json:"accepter"`
	MaterialCode                      string `json:"material_code"`
	Supplier                          string `json:"plant/supplier"`
	Stock                             string `json:"stock"`
	PurchaseContractType              string `json:"document_type"`
	PurchaseContractNo                string `json:"document_no"`
	ValidatedDate                     string `json:"validated_date"`
    Deleted                           bool   `json:"deleted"`
}

type SDC struct {
	ConnectionKey    string `json:"connection_key"`
	Result           bool   `json:"result"`
	RedisKey         string `json:"redis_key"`
	Filepath         string `json:"filepath"`
	PurchaseContract struct {
		PurchaseContract               string      `json:"PurchaseContract"`
		PurchaseContractType           string      `json:"PurchaseContractType"`
		CompanyCode                    string      `json:"CompanyCode"`
		PurchasingDocumentDeletionCode string      `json:"PurchasingDocumentDeletionCode"`
		CreationDate                   string      `json:"CreationDate"`
		Supplier                       string      `json:"Supplier"`
		PurchasingOrganization         string      `json:"PurchasingOrganization"`
		PurchasingGroup                string      `json:"PurchasingGroup"`
		PaymentTerms                   string      `json:"PaymentTerms"`
		NetPaymentDays                 string      `json:"NetPaymentDays"`
		DocumentCurrency               string      `json:"DocumentCurrency"`
		ExchangeRate                   string      `json:"ExchangeRate"`
		ValidityStartDate              string      `json:"ValidityStartDate"`
		ValidityEndDate                string      `json:"ValidityEndDate"`
		SupplierRespSalesPersonName    string      `json:"SupplierRespSalesPersonName"`
		SupplierPhoneNumber            string      `json:"SupplierPhoneNumber"`
		IncotermsClassification        string      `json:"IncotermsClassification"`
		PurchaseContractTargetAmount   string      `json:"PurchaseContractTargetAmount"`
		InvoicingParty                 string      `json:"InvoicingParty"`
		ReleaseCode                    string      `json:"ReleaseCode"`
		LastChangeDateTime             string      `json:"LastChangeDateTime"`
		PurchasingProcessingStatus     string      `json:"PurchasingProcessingStatus"`
		PurchasingProcessingStatusName string      `json:"PurchasingProcessingStatusName"`
		PurgContractIsInPreparation    bool        `json:"PurgContractIsInPreparation"`
		PurchaseContractItem           struct {
			PurchaseContractItem           string      `json:"PurchaseContractItem"`
			PurchasingContractDeletionCode string      `json:"PurchasingContractDeletionCode"`
			PurchaseContractItemText       string      `json:"PurchaseContractItemText"`
			CompanyCode                    string      `json:"CompanyCode"`
			Plant                          string      `json:"Plant"`
			StorageLocation                string      `json:"StorageLocation"`
			MaterialGroup                  string      `json:"MaterialGroup"`
			SupplierMaterialNumber         string      `json:"SupplierMaterialNumber"`
			OrderQuantityUnit              string      `json:"OrderQuantityUnit"`
			TargetQuantity                 string      `json:"TargetQuantity"`
			PurgDocReleaseOrderQuantity    string      `json:"PurgDocReleaseOrderQuantity"`
			OrderPriceUnit                 string      `json:"OrderPriceUnit"`
			ContractNetPriceAmount         string      `json:"ContractNetPriceAmount"`
			DocumentCurrency               string      `json:"DocumentCurrency"`
			NetPriceQuantity               string      `json:"NetPriceQuantity"`
			TaxCode                        string      `json:"TaxCode"`
			TaxCountry                     string      `json:"TaxCountry"`
			StockType                      string      `json:"StockType"`
			IsInfoRecordUpdated            string      `json:"IsInfoRecordUpdated"`
			PriceIsToBePrinted             bool        `json:"PriceIsToBePrinted"`
			PurgDocEstimatedPrice          bool        `json:"PurgDocEstimatedPrice"`
			PlannedDeliveryDurationInDays  string      `json:"PlannedDeliveryDurationInDays"`
			OverdelivTolrtdLmtRatioInPct   string      `json:"OverdelivTolrtdLmtRatioInPct"`
			UnlimitedOverdeliveryIsAllowed bool        `json:"UnlimitedOverdeliveryIsAllowed"`
			UnderdelivTolrtdLmtRatioInPct  string      `json:"UnderdelivTolrtdLmtRatioInPct"`
			PurchasingDocumentItemCategory string      `json:"PurchasingDocumentItemCategory"`
			AccountAssignmentCategory      string      `json:"AccountAssignmentCategory"`
			GoodsReceiptIsExpected         bool        `json:"GoodsReceiptIsExpected"`
			GoodsReceiptIsNonValuated      bool        `json:"GoodsReceiptIsNonValuated"`
			InvoiceIsExpected              bool        `json:"InvoiceIsExpected"`
			InvoiceIsGoodsReceiptBased     bool        `json:"InvoiceIsGoodsReceiptBased"`
			ManualDeliveryAddressID        string      `json:"ManualDeliveryAddressID"`
			VolumeUnit                     string      `json:"VolumeUnit"`
			Subcontractor                  string      `json:"Subcontractor"`
			EvaldRcptSettlmtIsAllowed      bool        `json:"EvaldRcptSettlmtIsAllowed"`
			Material                       string      `json:"Material"`
			ProductType                    string      `json:"ProductType"`
			MaterialType                   string      `json:"MaterialType"`
			ItemAddress                    struct {
				AddressID              string `json:"AddressID"`
				CityName               string `json:"CityName"`
				PostalCode             string `json:"PostalCode"`
				StreetName             string `json:"StreetName"`
				Country                string `json:"Country"`
				CorrespondenceLanguage string `json:"CorrespondenceLanguage"`
				Region                 string `json:"Region"`
				PhoneNumber            string `json:"PhoneNumber"`
				FaxNumber              string `json:"FaxNumber"`
			} `json:"ItemAddress"`
			ItemCondition struct {
				ConditionValidityEndDate     string      `json:"ConditionValidityEndDate"`
				ConditionType                string      `json:"ConditionType"`
				ConditionRecord              string      `json:"ConditionRecord"`
				ConditionSequentialNumber    string      `json:"ConditionSequentialNumber"`
				ConditionValidityStartDate   string      `json:"ConditionValidityStartDate"`
				PricingScaleType             string      `json:"PricingScaleType"`
				PricingScaleBasis            string      `json:"PricingScaleBasis"`
				ConditionScaleQuantity       string      `json:"ConditionScaleQuantity"`
				ConditionScaleQuantityUnit   string      `json:"ConditionScaleQuantityUnit"`
				ConditionScaleAmount         string      `json:"ConditionScaleAmount"`
				ConditionScaleAmountCurrency string      `json:"ConditionScaleAmountCurrency"`
				ConditionCalculationType     string      `json:"ConditionCalculationType"`
				ConditionRateValue           string      `json:"ConditionRateValue"`
				ConditionRateValueUnit       string      `json:"ConditionRateValueUnit"`
				ConditionQuantity            string      `json:"ConditionQuantity"`
				ConditionQuantityUnit        string      `json:"ConditionQuantityUnit"`
				BaseUnit                     string      `json:"BaseUnit"`
				ConditionIsDeleted           bool        `json:"ConditionIsDeleted"`
				PaymentTerms                 string      `json:"PaymentTerms"`
				ConditionReleaseStatus       string      `json:"ConditionReleaseStatus"`
			} `json:"ItemCondition"`
		} `json:"PurchaseContractItem"`
	} `json:"PurchaseContract"`
	APISchema          string   `json:"api_schema"`
	Accepter           []string `json:"accepter"`
	PurchaseContractNo string   `json:"purchase_contract_no"`
	Deleted            bool     `json:"deleted"`
}