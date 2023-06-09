package main

import (
	"fmt"
	"os"
	"os/signal"
	"ssr/cmd/config"
	"ssr/cmd/dependencies"
	"ssr/cmd/http"
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
