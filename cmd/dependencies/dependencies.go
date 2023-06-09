package dependencies

import (
	"nft-crud-service/cmd/config"
	"nft-crud-service/internal/application/marketplacesrvs"
	"nft-crud-service/internal/infra/markepplacerepo"
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
