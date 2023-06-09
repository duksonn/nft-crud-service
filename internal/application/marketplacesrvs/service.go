package marketplacesrvs

import (
	"context"
	"ssr/internal/domain"
)

//go:generate mockgen -source=./service.go -package=mocks -destination=../../../mocks/mockgen/marketplace_services.go

type Service interface {
	CreateNft(context.Context, *CreateNftInput) (*domain.Nft, error)
	FindNfts(context.Context, *FindNftsInput) (*domain.NftList, error)
	BuyNft(context.Context, *BuyNftInput) (*domain.NftUsers, error)
}

var _ Service = (*MarketplaceServices)(nil)

type MarketplaceServices struct {
	marketplaceRepository domain.MarketplaceRepository
}

func NewServices(
	marketplaceRepository domain.MarketplaceRepository,
) *MarketplaceServices {
	return &MarketplaceServices{
		marketplaceRepository: marketplaceRepository,
	}
}
