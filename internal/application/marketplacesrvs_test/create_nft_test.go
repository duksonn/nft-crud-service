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

func TestMarketplaceService_CreateNft(t *testing.T) {
	creation_date := time.Now().Format(time.RFC3339)
	type args struct {
		Image       string
		Description string
		CoCreators  []string
		User        string
	}
	type response struct {
		Nft *domain.Nft
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
			name: "Error 1. Should return error because cocreator doesnt exist in database",
			args: func() args {
				return args{
					Image:       "some image",
					Description: "some description",
					CoCreators:  []string{"some_co_creator"},
					User:        "the_owner_man",
				}
			},
			mock: func(dep *dependencies, args args) {
				dep.marketplaceRepository.EXPECT().GetUserById(gomock.Any(), gomock.Any()).Return(nil, errors.New("user 'some_co_creator' doesnt exist in database"))
			},
			response:      nil,
			errorResponse: utils.PError(errors.New("user 'some_co_creator' doesnt exist in database")),
		},
		{
			name: "OK 1. Should return nft created",
			args: func() args {
				return args{
					Image:       "some image",
					Description: "some description",
					CoCreators:  []string{"some_co_creator"},
					User:        "the_owner_man",
				}
			},
			mock: func(dep *dependencies, args args) {
				user := domain.NewUser("some_co_creator", 100)
				nft := domain.NewNft(
					"some id",
					"some image",
					"some description",
					"the_owner_man",
					[]string{"some_co_creator"},
					creation_date,
					"the_owner_man",
				)
				gomock.InOrder(
					dep.marketplaceRepository.EXPECT().GetUserById(gomock.Any(), gomock.Any()).Return(user, nil),
					dep.marketplaceRepository.EXPECT().SaveNft(gomock.Any(), gomock.Any()).Return(nft, nil),
				)
			},
			response: func() response {
				return response{
					Nft: domain.NewNft(
						"some id",
						"some image",
						"some description",
						"the_owner_man",
						[]string{"some_co_creator"},
						creation_date,
						"the_owner_man"),
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

			resp, err := s.CreateNft(context.TODO(), &marketplacesrvs.CreateNftInput{
				Image:       args.Image,
				Description: args.Description,
				CoCreators:  args.CoCreators,
				User:        args.User,
			})

			assert.NotNil(t, s)

			if tt.errorResponse != nil {
				assert.Equal(t, *tt.errorResponse, err)
			} else {
				response := tt.response()
				if &response != nil {
					assert.Equal(t, response.Nft.Id(), resp.Id())
					assert.Equal(t, response.Nft.Image(), resp.Image())
					assert.Equal(t, response.Nft.Description(), resp.Description())
					assert.Equal(t, response.Nft.CoCreators(), resp.CoCreators())
					assert.Equal(t, response.Nft.CreatedAt(), resp.CreatedAt())
					assert.Equal(t, response.Nft.CreatedBy(), resp.CreatedBy())
				}
			}
		})
	}
}
