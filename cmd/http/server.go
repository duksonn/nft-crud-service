package http

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"nft-crud-service/cmd/config"
	"nft-crud-service/cmd/dependencies"
)

func StartServer(cfg *config.Config, dep *dependencies.Dependencies) error {
	router := mux.NewRouter()
	router = routes(*router, dep)

	err := http.ListenAndServe(fmt.Sprintf(":%v", cfg.Server.Port), router)
	if err != nil {
		log.Fatalln("There's an error with the server", err)
	}
	return nil
}
