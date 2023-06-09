package domain

import "context"

//go:generate mockgen -source=./repository.go -package=mocks -destination=../../mocks/mockgen/marketplace_repository.go

type MarketplaceRepository interface {
	SaveNft(ctx context.Context, nft Nft) (*Nft, error)
	GetUserById(ctx context.Context, userId string) (*User, error)
	FindNfts(ctx context.Context, next, took *int) (*NftList, error)
	GetNftById(ctx context.Context, nftId string) (*Nft, error)
	UpdateUsersBalances(ctx context.Context, balances map[string]float64) ([]User, error)
}
