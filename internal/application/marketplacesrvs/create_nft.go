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
	/** Validate coCreators */
	err := s.validateCoCreators(ctx, input.CoCreators)
	if err != nil {
		return nil, err
	}

	/** Create NFT */
	var id = utils.GenerateUuid()
	var now = time.Now().Format(time.RFC3339)
	var nft = domain.NewNft(id, input.Image, input.Description, input.User, input.CoCreators, now, input.User)

	/** Save NFT */
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
