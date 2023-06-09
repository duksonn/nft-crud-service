package marketplacesrvs

import (
	"context"
	"ssr/internal/domain"
)

type FindNftsInput struct {
	Next *int
	Took *int
}

func (s *MarketplaceServices) FindNfts(ctx context.Context, input *FindNftsInput) (*domain.NftList, error) {
	nfts, err := s.marketplaceRepository.FindNfts(ctx, input.Next, input.Took)
	if err != nil {
		return nil, err
	}
	return nfts, nil
}
