package marketplacesrvs_test

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"ssr/internal/application/marketplacesrvs"
	"ssr/internal/domain"
	"ssr/pkg/utils"
	"testing"
	"time"
)

func TestMarketplaceService_BuyNft(t *testing.T) {
	type args struct {
		NftId   string
		BuyerId string
		Amount  float64
	}
	type response struct {
		NftUsers *domain.NftUsers
	}

	t.Parallel()
	tests := []struct {
		name          string
		args          func() args
		mock          func(dep *dependencies, args args)
		response      func() response
		errorResponse *error
	}{
		{
			name: "Error 1. Should return error because buyer doesnt exist in database",
			args: func() args {
				return args{
					NftId:   "nft_1",
					BuyerId: "buyer_3",
					Amount:  10,
				}
			},
			mock: func(dep *dependencies, args args) {
				dep.marketplaceRepository.EXPECT().GetUserById(gomock.Any(), gomock.Any()).Return(nil, errors.New("user 'buyer_3' doesnt exist in database"))
			},
			response:      nil,
			errorResponse: utils.PError(errors.New("user 'buyer_3' doesnt exist in database")),
		},
		{
			name: "Error 2. Should return error because nft doesnt exist in database",
			args: func() args {
				return args{
					NftId:   "nft_1",
					BuyerId: "buyer_3",
					Amount:  10,
				}
			},
			mock: func(dep *dependencies, args args) {
				gomock.InOrder(
					dep.marketplaceRepository.EXPECT().GetUserById(gomock.Any(), gomock.Any()).Return(getUser("buyer_3"), nil),
					dep.marketplaceRepository.EXPECT().GetNftById(gomock.Any(), gomock.Any()).Return(nil, errors.New("nft 'nft_1' doesnt exist in database")),
				)
			},
			response:      nil,
			errorResponse: utils.PError(errors.New("nft 'nft_1' doesnt exist in database")),
		},
		{
			name: "Error 3. Should return error because update users balances failed",
			args: func() args {
				return args{
					NftId:   "nft_1",
					BuyerId: "buyer_3",
					Amount:  10,
				}
			},
			mock: func(dep *dependencies, args args) {
				gomock.InOrder(
					dep.marketplaceRepository.EXPECT().GetUserById(gomock.Any(), gomock.Any()).Return(getUser("buyer_3"), nil),
					dep.marketplaceRepository.EXPECT().GetNftById(gomock.Any(), gomock.Any()).Return(getNft("buyer_1"), nil),
					dep.marketplaceRepository.EXPECT().GetUserById(gomock.Any(), gomock.Any()).Return(getUser("buyer_1"), nil),
					dep.marketplaceRepository.EXPECT().UpdateUsersBalances(gomock.Any(), gomock.Any()).Return(nil, errors.New("error on UpdateUsersBalances")),
				)
			},
			response:      nil,
			errorResponse: utils.PError(errors.New("error on UpdateUsersBalances")),
		},
		{
			name: "Error 4. Should return error because save nft failed",
			args: func() args {
				return args{
					NftId:   "nft_1",
					BuyerId: "buyer_3",
					Amount:  10,
				}
			},
			mock: func(dep *dependencies, args args) {
				userBalance := map[string]float64{"buyer_1": 110, "buyer_3": 90}
				gomock.InOrder(
					dep.marketplaceRepository.EXPECT().GetUserById(gomock.Any(), gomock.Any()).Return(getUser("buyer_3"), nil),
					dep.marketplaceRepository.EXPECT().GetNftById(gomock.Any(), gomock.Any()).Return(getNft("buyer_1"), nil),
					dep.marketplaceRepository.EXPECT().GetUserById(gomock.Any(), gomock.Any()).Return(getUser("buyer_1"), nil),
					dep.marketplaceRepository.EXPECT().UpdateUsersBalances(gomock.Any(), gomock.Any()).Return(getUsers(userBalance), nil),
					dep.marketplaceRepository.EXPECT().SaveNft(gomock.Any(), gomock.Any()).Return(nil, errors.New("error on SaveNft")),
				)
			},
			response:      nil,
			errorResponse: utils.PError(errors.New("error on SaveNft")),
		},
		{
			name: "OK 1. Should return nft and list of users with balances updated",
			args: func() args {
				return args{
					NftId:   "nft_1",
					BuyerId: "buyer_3",
					Amount:  10,
				}
			},
			mock: func(dep *dependencies, args args) {
				userBalance := map[string]float64{"buyer_3": 90, "buyer_1": 110}
				gomock.InOrder(
					dep.marketplaceRepository.EXPECT().GetUserById(gomock.Any(), gomock.Any()).Return(getUser("buyer_3"), nil),
					dep.marketplaceRepository.EXPECT().GetNftById(gomock.Any(), gomock.Any()).Return(getNft("buyer_1"), nil),
					dep.marketplaceRepository.EXPECT().GetUserById(gomock.Any(), gomock.Any()).Return(getUser("buyer_1"), nil),
					dep.marketplaceRepository.EXPECT().UpdateUsersBalances(gomock.Any(), gomock.Any()).Return(getUsers(userBalance), nil),
					dep.marketplaceRepository.EXPECT().SaveNft(gomock.Any(), gomock.Any()).Return(getNft("buyer_3"), nil),
				)
			},
			response: func() response {
				nft := domain.NewNft(
					"nft_1",
					"some image",
					"some description",
					"buyer_3",
					nil,
					time.Now().Format(time.RFC3339),
					"buyer_1",
				)
				users := []domain.User{*domain.NewUser("buyer_3", 90), *domain.NewUser("buyer_1", 110)}
				return response{
					NftUsers: domain.NewNftUsers(*nft, users),
				}
			},
			errorResponse: nil,
		},
		{
			name: "OK 2. Should return nft and list of users with balances updated with cocreators",
			args: func() args {
				return args{
					NftId:   "nft_1",
					BuyerId: "buyer_3",
					Amount:  10,
				}
			},
			mock: func(dep *dependencies, args args) {
				userBalance := map[string]float64{"buyer_3": 90, "buyer_1": 108, "buyer_2": 102}
				gomock.InOrder(
					dep.marketplaceRepository.EXPECT().GetUserById(gomock.Any(), gomock.Any()).Return(getUser("buyer_3"), nil),
					dep.marketplaceRepository.EXPECT().GetNftById(gomock.Any(), gomock.Any()).Return(getNftWithCoCreators("buyer_1"), nil),
					dep.marketplaceRepository.EXPECT().GetUserById(gomock.Any(), gomock.Any()).Return(getUser("buyer_1"), nil),
					dep.marketplaceRepository.EXPECT().GetUserById(gomock.Any(), gomock.Any()).Return(getUser("buyer_2"), nil),
					dep.marketplaceRepository.EXPECT().UpdateUsersBalances(gomock.Any(), gomock.Any()).Return(getUsers(userBalance), nil),
					dep.marketplaceRepository.EXPECT().SaveNft(gomock.Any(), gomock.Any()).Return(getNft("buyer_3"), nil),
				)

			},
			response: func() response {
				nft := domain.NewNft(
					"nft_1",
					"some image",
					"some description",
					"buyer_3",
					nil,
					time.Now().Format(time.RFC3339),
					"buyer_1",
				)
				users := []domain.User{*domain.NewUser("buyer_3", 90), *domain.NewUser("buyer_1", 108), *domain.NewUser("buyer_2", 102)}
				return response{
					NftUsers: domain.NewNftUsers(*nft, users),
				}
			},
			errorResponse: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dependencies := makeDependencies(t)
			args := tt.args()

			s := initService(dependencies)

			tt.mock(dependencies, args)

			resp, err := s.BuyNft(context.TODO(), &marketplacesrvs.BuyNftInput{
				NftId:   args.NftId,
				BuyerId: args.BuyerId,
				Amount:  args.Amount,
			})

			assert.NotNil(t, s)

			if tt.errorResponse != nil {
				assert.Equal(t, *tt.errorResponse, err)
			} else {
				response := tt.response()
				if &response != nil {
					assert.Equal(t, response.NftUsers.Nft(), resp.Nft())
					assert.Equal(t, response.NftUsers.Users(), resp.Users())
				}
			}
		})
	}
}

func getUser(id string) *domain.User {
	return domain.NewUser(id, 100)
}

func getUsers(userBalance map[string]float64) []domain.User {
	var users []domain.User
	for id, balance := range userBalance {
		users = append(users, *domain.NewUser(id, balance))
	}
	return users
}

func getNft(owner string) *domain.Nft {
	return domain.NewNft(
		"nft_1",
		"some image",
		"some description",
		owner,
		nil,
		time.Now().Format(time.RFC3339),
		"buyer_1",
	)
}

func getNftWithCoCreators(owner string) *domain.Nft {
	return domain.NewNft(
		"nft_1",
		"some image",
		"some description",
		owner,
		[]string{"buyer_2"},
		time.Now().Format(time.RFC3339),
		"buyer_1",
	)
}
