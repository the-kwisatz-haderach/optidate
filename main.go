package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/the-kwisatz-haderach/optidate/internal/dateservice"
)

func main() {
	ctx := context.Background()
	router := http.NewServeMux()

	ds := dateservice.New()

	router.HandleFunc("/countries", func(w http.ResponseWriter, r *http.Request) {
		resp, err := ds.GetCountries(ctx)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	})

	http.ListenAndServe(":8000", router)
}
