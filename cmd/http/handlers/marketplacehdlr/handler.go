package marketplacehdlr

import "ssr/internal/application/marketplacesrvs"

type handler struct {
	services marketplacesrvs.Service
}

func NewMarketplaceHandler(service marketplacesrvs.Service) *handler {
	return &handler{
		services: service,
	}
}
