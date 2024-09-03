package main

import (
	"fmt"
	"net/http"

	"github.com/rs/zerolog"

	"github.com/the-kwisatz-haderach/optidate/internal/cache"
	"github.com/the-kwisatz-haderach/optidate/internal/config"
	"github.com/the-kwisatz-haderach/optidate/internal/dateservice"
	"github.com/the-kwisatz-haderach/optidate/internal/router"
)

func init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
}

func main() {
	cfg := config.ParseConfig()
	cache := cache.New()
	ds := dateservice.New(cache)
	router := router.New(ds)

	fmt.Printf("listening on port %d...", cfg.Port)
	http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), router)
}
