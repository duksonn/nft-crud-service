package marketplacehdlr

import (
	"bytes"
	"encoding/json"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"nft-crud-service/internal/domain"
	"testing"
	"time"
)

func TestMarketplaceHandler_BuyNft(t *testing.T) {
	mockDate := time.Date(2022, 01, 30, 00, 00, 00, 00, time.UTC).Format(time.RFC3339)
	type args struct {
		body map[string]interface{}
	}
	tests := []struct {
		name          string
		args          args
		mocks         func(dep *dependencies, a args)
		response      map[string]interface{}
		errorResponse *errResponse
	}{
		{
			name: `Error 1. Should return error because nft_id is required in body`,
			args: args{
				body: map[string]interface{}{
					"buyer_id": "buyer_3",
					"amount":   10,
				},
			},
			mocks:    func(dep *dependencies, args args) {},
			response: nil,
			errorResponse: &errResponse{
				Messages: []string{"nft_id is required in body"},
				Code:     "BAD_REQUEST",
			},
		},
		{
			name: `Error 2. Should return error because buyer_id is required in body`,
			args: args{
				body: map[string]interface{}{
					"nft_id": "nft_1",
					"amount": 10,
				},
			},
			mocks:    func(dep *dependencies, args args) {},
			response: nil,
			errorResponse: &errResponse{
				Messages: []string{"buyer_id is required in body"},
				Code:     "BAD_REQUEST",
			},
		},
		{
			name: `Error 3. Should return error because amount is required in body`,
			args: args{
				body: map[string]interface{}{
					"nft_id":   "nft_1",
					"buyer_id": "buyer_3",
				},
			},
			mocks:    func(dep *dependencies, args args) {},
			response: nil,
			errorResponse: &errResponse{
				Messages: []string{"amount is required in body"},
				Code:     "BAD_REQUEST",
			},
		},
		{
			name: `Ok 1. Should return nft and users updated`,
			args: args{
				body: map[string]interface{}{
					"nft_id":   "nft_1",
					"buyer_id": "buyer_3",
					"amount":   10,
				},
			},
			mocks: func(dep *dependencies, args args) {
				userBalance := map[string]float64{"buyer_3": 90, "buyer_1": 110}
				gomock.InOrder(
					dep.marketplaceRepository.EXPECT().GetUserById(gomock.Any(), gomock.Any()).Return(getUser("buyer_3"), nil),
					dep.marketplaceRepository.EXPECT().GetNftById(gomock.Any(), gomock.Any()).Return(getNft("buyer_1"), nil),
					dep.marketplaceRepository.EXPECT().GetUserById(gomock.Any(), gomock.Any()).Return(getUser("buyer_1"), nil),
					dep.marketplaceRepository.EXPECT().UpdateUsersBalances(gomock.Any(), gomock.Any()).Return(getUsers(userBalance), nil),
					dep.marketplaceRepository.EXPECT().SaveNft(gomock.Any(), gomock.Any()).Return(getNft("buyer_3"), nil),
				)
			},
			response: map[string]interface{}{
				"nft": map[string]interface{}{
					"id":          "nft_1",
					"image":       "some image",
					"description": "some description",
					"owner":       "buyer_3",
					"co_creators": nil,
					"created_at":  mockDate,
					"created_by":  "buyer_1",
				},
				"users": []map[string]interface{}{
					{
						"id":      "buyer_3",
						"balance": 90,
					},
					{
						"id":      "buyer_1",
						"balance": 110,
					},
				},
			},
			errorResponse: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dependencies := makeDependencies(t)
			handler := buildHandler(dependencies)

			tt.mocks(dependencies, tt.args)

			/** Reader */
			path := "/v1/nft/buy"
			method := http.MethodPost
			requestBody, _ := json.Marshal(tt.args.body)
			reader := httptest.NewRequest(method, path, bytes.NewReader(requestBody))
			q := reader.URL.Query()
			reader.URL.RawQuery = q.Encode()

			/** Writer */
			writer := httptest.NewRecorder()

			/** Build Server */
			router := mux.NewRouter()
			router.HandleFunc(path, handler.BuyNft).Methods(method)
			router.ServeHTTP(writer, reader)

			/** Asserts */
			assert.NotNil(t, handler.services)
			if tt.response != nil {
				expectedRespByte, _ := json.Marshal(tt.response)
				var expectedResp nftUsersResponse
				_ = json.Unmarshal(expectedRespByte, &expectedResp)

				var resp nftUsersResponse
				_ = json.Unmarshal(writer.Body.Bytes(), &resp)

				assert.Equal(t, expectedResp, resp)
			}
			if tt.errorResponse != nil {
				var resp errResponse
				_ = json.Unmarshal(writer.Body.Bytes(), &resp)

				assert.Equal(t, *tt.errorResponse, resp)
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
	mockDate := time.Date(2022, 01, 30, 00, 00, 00, 00, time.UTC).Format(time.RFC3339)
	return domain.NewNft(
		"nft_1",
		"some image",
		"some description",
		owner,
		nil,
		mockDate,
		"buyer_1",
	)
}
