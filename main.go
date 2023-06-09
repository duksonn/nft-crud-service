package main

import (
	"fmt"
	"nft-crud-service/cmd/config"
	"nft-crud-service/cmd/dependencies"
	"nft-crud-service/cmd/http"
	"os"
	"os/signal"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		panic(fmt.Sprintf("main: can't read config: %s", err.Error()))
	}

	deps := dependencies.Init(cfg)
	fmt.Println("dep init successfully")

	err = http.StartServer(cfg, deps)

	// Wait for terminate signal to shut down server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}
