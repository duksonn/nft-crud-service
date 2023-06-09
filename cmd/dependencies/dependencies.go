package dependencies

import (
	"ssr/cmd/config"
	"ssr/internal/application/marketplacesrvs"
	"ssr/internal/infra/markepplacerepo"
)

type Dependencies struct {
	MarketplaceService marketplacesrvs.Service
}

func Init(config *config.Config) *Dependencies {
	/** Repositories and service dependencies */
	marketplaceRepo := markepplacerepo.NewRepository(config.MarketplaceMySql)

	marketplaceSrvs := marketplacesrvs.NewServices(&marketplaceRepo)
	return &Dependencies{marketplaceSrvs}
}
