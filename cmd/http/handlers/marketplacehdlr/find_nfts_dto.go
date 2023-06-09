package marketplacehdlr

import (
	"net/http"
	"ssr/internal/application/marketplacesrvs"
	"ssr/pkg/server"
)

type findNftsRequest struct {
	Next *int `json:"next"`
	Took *int `json:"took"`
}

func newFindNftsRequestDTO(r *http.Request) (*findNftsRequest, error) {
	nextParam, err := server.GetIntParam(r, "next", 0)
	if err != nil {
		return nil, err
	}
	tookParam, err := server.GetIntParam(r, "took", 0)
	if err != nil {
		return nil, err
	}
	resp := &findNftsRequest{
		Next: &nextParam,
		Took: &tookParam,
	}
	return resp, nil
}

func (r *findNftsRequest) toFindNftsInput() *marketplacesrvs.FindNftsInput {
	return &marketplacesrvs.FindNftsInput{
		Next: r.Next,
		Took: r.Took,
	}
}
