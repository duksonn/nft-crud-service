package marketplacehdlr

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"nft-crud-service/internal/application/marketplacesrvs"
)

type buyNftRequest struct {
	Body buyNftRequestBody
}

type buyNftRequestBody struct {
	NftId   string  `json:"nft_id"`
	BuyerId string  `json:"buyer_id"`
	Amount  float64 `json:"amount"`
}

func newBuyNftRequestDTO(r *http.Request) (*buyNftRequest, error) {
	var requestBody buyNftRequestBody
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &requestBody)
	if err != nil {
		return nil, err
	}

	resp := &buyNftRequest{
		Body: requestBody,
	}

	return resp, nil
}

func (r *buyNftRequest) validate() error {
	if r.Body.NftId == "" {
		return errors.New("nft_id is required in body")
	}
	if r.Body.BuyerId == "" {
		return errors.New("buyer_id is required in body")
	}
	if r.Body.Amount == 0 {
		return errors.New("amount is required in body")
	}
	return nil
}

func (r *buyNftRequest) toBuyNftInput() *marketplacesrvs.BuyNftInput {
	return &marketplacesrvs.BuyNftInput{
		NftId:   r.Body.NftId,
		BuyerId: r.Body.BuyerId,
		Amount:  r.Body.Amount,
	}
}
