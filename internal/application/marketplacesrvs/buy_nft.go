package marketplacesrvs

import (
	"context"
	"nft-crud-service/internal/domain"
)

type BuyNftInput struct {
	NftId   string
	BuyerId string
	Amount  float64
}

func (s *MarketplaceServices) BuyNft(ctx context.Context, input *BuyNftInput) (*domain.NftUsers, error) {
	/** Validate is buyer exist in db */
	buyer, err := s.marketplaceRepository.GetUserById(ctx, input.BuyerId)
	if err != nil {
		return nil, err
	}

	/** Get NFT to sell */
	nft, err := s.marketplaceRepository.GetNftById(ctx, input.NftId)
	if err != nil {
		return nil, err
	}

	/** Update balances */
	users, err := s.updateBalances(ctx, *nft, input.Amount, *buyer)
	if err != nil {
		return nil, err
	}

	/** Save new NFT */
	var coCreators = nft.CoCreators()
	coCreators = append(coCreators, nft.Owner())
	newNft := domain.NewNft(nft.Id(), nft.Image(), nft.Description(), input.BuyerId, coCreators, nft.CreatedAt(), nft.CreatedBy())
	savedNft, err := s.marketplaceRepository.SaveNft(ctx, *newNft)
	if err != nil {
		return nil, err
	}

	nftUsers := domain.NewNftUsers(*savedNft, users)

	return nftUsers, nil
}

func (s *MarketplaceServices) updateBalances(ctx context.Context, nft domain.Nft, amount float64, buyer domain.User) ([]domain.User, error) {
	/** get owner */
	owner, err := s.marketplaceRepository.GetUserById(ctx, nft.Owner())
	if err != nil {
		return nil, err
	}

	/** create user balance map to update balances */
	userBalanceMap := make(map[string]float64)

	/** if nft has cocreators split and assign amounts */
	if len(nft.CoCreators()) > 0 {
		/** get 80% amount for owner and 20% to creators */
		owner80Fee, creators20Fee := splitAmount(amount)
		creators20Fee = creators20Fee / float64(len(nft.CoCreators()))

		/** add amount to owner and discount amount from buyer user */
		userBalanceMap[owner.Id()] = owner.Balance() + owner80Fee
		userBalanceMap[buyer.Id()] = buyer.Balance() - amount

		/** calculate creators balance */
		for _, coCreator := range nft.CoCreators() {
			user, err := s.marketplaceRepository.GetUserById(ctx, coCreator)
			if err != nil {
				return nil, err
			}
			userBalanceMap[user.Id()] = user.Balance() + creators20Fee
		}
	} else {
		userBalanceMap[owner.Id()] = owner.Balance() + amount
		userBalanceMap[buyer.Id()] = buyer.Balance() - amount
	}

	users, err := s.marketplaceRepository.UpdateUsersBalances(ctx, userBalanceMap)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func splitAmount(amount float64) (float64, float64) {
	amount20 := (20 * amount) / 100
	return amount - amount20, amount20
}
