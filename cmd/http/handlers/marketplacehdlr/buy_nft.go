package marketplacehdlr

import (
	"fmt"
	"log"
	"net/http"
	"nft-crud-service/pkg/server"
)

func (h *handler) BuyNft(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	/** Build requestDTO */
	requestDTO, err := newBuyNftRequestDTO(r)
	if err != nil {
		log.Println(err.Error())
		server.BadRequest(w, r, "BAD_REQUEST", err.Error())
		return
	}

	/** Validate request */
	err = requestDTO.validate()
	if err != nil {
		fmt.Println(err.Error())
		server.BadRequest(w, r, "BAD_REQUEST", err.Error())
		return
	}

	/** Service */
	nftUsersResponse, err := h.services.BuyNft(ctx, requestDTO.toBuyNftInput())
	if err != nil {
		fmt.Println(err.Error())
		server.InternalServerError(w, r, err)
		return
	}

	/** Build response */
	response := buildNftUsersResponse(*nftUsersResponse)

	server.OK(w, r, response)
}
