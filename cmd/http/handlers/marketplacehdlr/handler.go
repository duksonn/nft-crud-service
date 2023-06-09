package marketplacehdlr

import "nft-crud-service/internal/application/marketplacesrvs"

type handler struct {
	services marketplacesrvs.Service
}

func NewMarketplaceHandler(service marketplacesrvs.Service) *handler {
	return &handler{
		services: service,
	}
}
