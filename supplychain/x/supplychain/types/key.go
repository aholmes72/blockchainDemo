package types

const (
	// ModuleName is the name of the module
	ModuleName = "supplychain"

	// StoreKey to be used when creating the KVStore
	StoreKey = ModuleName

	// RouterKey to be used for routing msgs
	RouterKey = ModuleName

	// QuerierRoute to be used for querier msgs
	QuerierRoute = ModuleName
)

const (
	ReceiptPrefix = "receipt-value-"
	ReceiptCountPrefix = "receipt-count-"
)
		
const (
	ShipmentPrefix = "shipment-value-"
	ShipmentCountPrefix = "shipment-count-"
)
		