package marketplacehdlr

import (
	"fmt"
	"log"
	"net/http"
	"ssr/pkg/server"
)

func (h *handler) CreateNft(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	/** Build requestDTO */
	requestDTO, err := newCreateNftRequestDTO(r)
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
	nftResponse, err := h.services.CreateNft(ctx, requestDTO.toCreateNftInput())
	if err != nil {
		fmt.Println(err.Error())
		server.InternalServerError(w, r, err)
		return
	}

	/** Build response */
	response := buildNftResponse(*nftResponse)

	server.OK(w, r, response)
}
