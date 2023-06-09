package marketplacesrvs

import (
	"context"
	"nft-crud-service/internal/domain"
	"nft-crud-service/pkg/utils"
	"time"
)

type CreateNftInput struct {
	Image       string
	Description string
	CoCreators  []string
	User        string
}

func (s *MarketplaceServices) CreateNft(ctx context.Context, input *CreateNftInput) (*domain.Nft, error) {
	/** validate coCreators */
	err := s.validateCoCreators(ctx, input.CoCreators)
	if err != nil {
		return nil, err
	}

	/** create NFT */
	id := utils.GenerateUuid()
	now := time.Now().Format(time.RFC3339)
	nft := domain.NewNft(id, input.Image, input.Description, input.User, input.CoCreators, now, input.User)

	/** insert database */
	nftRes, err := s.marketplaceRepository.SaveNft(ctx, *nft)
	if err != nil {
		return nil, err
	}

	return nftRes, nil
}

func (s *MarketplaceServices) validateCoCreators(ctx context.Context, coCreators []string) error {
	for _, c := range coCreators {
		_, err := s.marketplaceRepository.GetUserById(ctx, c)
		if err != nil {
			return err
		}
	}
	return nil
}
