package rest

import (
	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client/context"
)

// RegisterRoutes registers supplychain-related REST handlers to a router
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router) {
  // this line is used by starport scaffolding # 1
		r.HandleFunc("/supplychain/shipment", createShipmentHandler(cliCtx)).Methods("POST")
		r.HandleFunc("/supplychain/shipment", listShipmentHandler(cliCtx, "supplychain")).Methods("GET")
		r.HandleFunc("/supplychain/shipment/{key}", getShipmentHandler(cliCtx, "supplychain")).Methods("GET")
		r.HandleFunc("/supplychain/shipment", setShipmentHandler(cliCtx)).Methods("PUT")
		r.HandleFunc("/supplychain/shipment", deleteShipmentHandler(cliCtx)).Methods("DELETE")

		
		r.HandleFunc("/supplychain/receipt", createReceiptHandler(cliCtx)).Methods("POST")
		r.HandleFunc("/supplychain/receipt", listReceiptHandler(cliCtx, "supplychain")).Methods("GET")
		r.HandleFunc("/supplychain/receipt/{key}", getReceiptHandler(cliCtx, "supplychain")).Methods("GET")
		r.HandleFunc("/supplychain/receipt", setReceiptHandler(cliCtx)).Methods("PUT")
		r.HandleFunc("/supplychain/receipt", deleteReceiptHandler(cliCtx)).Methods("DELETE")

		
}
