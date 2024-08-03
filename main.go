package main

import (
	"fmt"
	"net/http"

	"github.com/the-kwisatz-haderach/optidate/internal/dateservice"
	"github.com/the-kwisatz-haderach/optidate/internal/router"
)

const PORT = 8000

func main() {
	ds := dateservice.New()
	router := router.New(ds)
	fmt.Printf("listening on port %d...", PORT)
	http.ListenAndServe(fmt.Sprintf(":%d", PORT), router)
}
