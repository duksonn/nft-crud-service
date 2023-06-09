package marketplacehdlr

import "nft-crud-service/internal/domain"

type nftResponse struct {
	Id          string   `json:"id"`
	Image       string   `json:"image"`
	Description string   `json:"description"`
	Owner       string   `json:"owner"`
	CoCreators  []string `json:"co_creators"`
	CreatedAt   string   `json:"created_at"`
	CreatedBy   string   `json:"created_by"`
}

func buildNftResponse(nft domain.Nft) nftResponse {
	return nftResponse{
		Id:          nft.Id(),
		Image:       nft.Image(),
		Description: nft.Description(),
		Owner:       nft.Owner(),
		CoCreators:  nft.CoCreators(),
		CreatedAt:   nft.CreatedAt(),
		CreatedBy:   nft.CreatedBy(),
	}
}

type nftListResponse struct {
	Data []nftResponse `json:"data"`
	Next *int          `json:"next"`
	Took *int          `json:"took"`
}

func buildNftResponseAsList(nfts domain.NftList) nftListResponse {
	var resp nftListResponse
	var nftsResponse []nftResponse
	for _, nft := range nfts.Data() {
		nftsResponse = append(nftsResponse, nftResponse{
			Id:          nft.Id(),
			Image:       nft.Image(),
			Description: nft.Description(),
			Owner:       nft.Owner(),
			CoCreators:  nft.CoCreators(),
			CreatedAt:   nft.CreatedAt(),
			CreatedBy:   nft.CreatedBy(),
		})
	}
	resp.Data = nftsResponse
	resp.Next = nfts.Next()
	resp.Took = nfts.Took()
	return resp
}

type userResponse struct {
	Id      string  `json:"id"`
	Balance float64 `json:"balance"`
}

type nftUsersResponse struct {
	Nft   nftResponse    `json:"nft"`
	Users []userResponse `json:"users"`
}

func buildNftUsersResponse(nftUsers domain.NftUsers) nftUsersResponse {
	var users []userResponse
	for _, u := range nftUsers.Users() {
		users = append(users, userResponse{
			Id:      u.Id(),
			Balance: u.Balance(),
		})
	}
	return nftUsersResponse{
		Nft:   buildNftResponse(nftUsers.Nft()),
		Users: users,
	}
}
