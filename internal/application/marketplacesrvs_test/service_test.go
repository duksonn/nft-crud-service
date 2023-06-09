package marketplacesrvs_test

import (
	"github.com/golang/mock/gomock"
	"nft-crud-service/internal/application/marketplacesrvs"
	"nft-crud-service/mocks/mockgen"
	"testing"
)

type dependencies struct {
	marketplaceRepository *mocks.MockMarketplaceRepository
}

func makeDependencies(t *testing.T) *dependencies {
	return &dependencies{
		marketplaceRepository: mocks.NewMockMarketplaceRepository(gomock.NewController(t)),
	}
}

func initService(dep *dependencies) *marketplacesrvs.MarketplaceServices {
	return marketplacesrvs.NewServices(dep.marketplaceRepository)
}
