package marketplacehdlr

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"ssr/internal/application/marketplacesrvs"
)

type createNftRequest struct {
	Body createNftRequestBody
}

type createNftRequestBody struct {
	Image       string   `json:"image"`
	Description string   `json:"description"`
	CoCreators  []string `json:"co_creators"`
	User        string   `json:"user"`
}

func newCreateNftRequestDTO(r *http.Request) (*createNftRequest, error) {
	var requestBody createNftRequestBody
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &requestBody)
	if err != nil {
		return nil, err
	}

	resp := &createNftRequest{
		Body: requestBody,
	}

	return resp, nil
}

func (r *createNftRequest) validate() error {
	if r.Body.Image == "" {
		return errors.New("image is required in body")
	}
	if r.Body.Description == "" {
		return errors.New("description is required in body")
	}
	if r.Body.User == "" {
		return errors.New("user is required in body")
	}
	return nil
}

func (r *createNftRequest) toCreateNftInput() *marketplacesrvs.CreateNftInput {
	return &marketplacesrvs.CreateNftInput{
		Image:       r.Body.Image,
		Description: r.Body.Description,
		CoCreators:  r.Body.CoCreators,
		User:        r.Body.User,
	}
}
