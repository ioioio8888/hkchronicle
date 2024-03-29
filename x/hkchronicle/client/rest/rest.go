package rest

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/context"

	"github.com/gorilla/mux"
)

const (
	restName = "hkc"
)

// RegisterRoutes - Central function to define routes that get registered by the main application
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router, storeName string) {

	r.HandleFunc(fmt.Sprintf("/%s/hkc", storeName), eventsHandler(cliCtx, storeName)).Methods("GET")
	r.HandleFunc(fmt.Sprintf("/%s/hkc", storeName), buyEventHandler(cliCtx)).Methods("POST")
	r.HandleFunc(fmt.Sprintf("/%s/hkc", storeName), setEventHandler(cliCtx)).Methods("PUT")

	r.HandleFunc(fmt.Sprintf("/%s/hkc/stake", storeName), stakeEventHandler(cliCtx)).Methods("PUT")
	r.HandleFunc(fmt.Sprintf("/%s/hkc/unstake", storeName), unStakeEventHandler(cliCtx)).Methods("PUT")

	r.HandleFunc(fmt.Sprintf("/%s/hkc/{%s}", storeName, restName), resolveEventHandler(cliCtx, storeName)).Methods("GET")
	r.HandleFunc(fmt.Sprintf("/%s/hkc/{%s}/eventDetails", storeName, restName), eventDetailsHandler(cliCtx, storeName)).Methods("GET")
	r.HandleFunc(fmt.Sprintf("/%s/hkc", storeName), deleteEventHandler(cliCtx)).Methods("DELETE")
}
