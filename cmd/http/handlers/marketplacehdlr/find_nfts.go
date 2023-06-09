package marketplacehdlr

import (
	"fmt"
	"log"
	"net/http"
	"ssr/pkg/server"
)

func (h *handler) FindNfts(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	/** Build requestDTO */
	requestDTO, err := newFindNftsRequestDTO(r)
	if err != nil {
		log.Println(err.Error())
		server.BadRequest(w, r, "BAD_REQUEST", err.Error())
		return
	}

	/** Service */
	nftListResponse, err := h.services.FindNfts(ctx, requestDTO.toFindNftsInput())
	if err != nil {
		fmt.Println(err.Error())
		server.InternalServerError(w, r, err)
		return
	}

	/** Build response */
	response := buildNftResponseAsList(*nftListResponse)

	server.OK(w, r, response)
}
