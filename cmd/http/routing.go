package http

import (
	"github.com/gorilla/mux"
	"net/http"
	"ssr/cmd/dependencies"
	"ssr/cmd/http/handlers/marketplacehdlr"
)

func routes(router mux.Router, dep *dependencies.Dependencies) *mux.Router {
	marketplaceHdlr := marketplacehdlr.NewMarketplaceHandler(dep.MarketplaceService)

	router.HandleFunc("/v1/nft", marketplaceHdlr.CreateNft).Methods(http.MethodPost)
	router.HandleFunc("/v1/nft", marketplaceHdlr.FindNfts).Methods(http.MethodGet)
	router.HandleFunc("/v1/nft/buy", marketplaceHdlr.BuyNft).Methods(http.MethodPost)

	return &router
}
