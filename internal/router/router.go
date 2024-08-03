package router

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/the-kwisatz-haderach/optidate/internal/dateservice"
)

type ApiError struct {
	Message string `json:"message"`
}

func New(service *dateservice.Service) http.Handler {
	router := httprouter.New()

	// Serve frontend SPA
	router.NotFound = http.FileServer(http.Dir("frontend/dist/"))

	router.GET("/countries", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		resp, err := service.GetCountries(r.Context())
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	})

	router.GET("/calendar/:country", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		var err error
		w.Header().Set("Content-Type", "application/json")
		countryCode := ps.ByName("country")
		if len(countryCode) != 2 {
			w.WriteHeader(http.StatusBadRequest)
			resp := ApiError{"invalid country code"}
			json.NewEncoder(w).Encode(resp)
			return
		}
		tq := r.URL.Query().Get("threshold")
		var threshold int
		if tq != "" {
			threshold, err = strconv.Atoi(tq)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				resp := ApiError{"invalid value for query 'threshold'"}
				json.NewEncoder(w).Encode(resp)
				return
			}
		}
		resp, err := service.GetCalendar(r.Context(), countryCode, threshold)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			resp := ApiError{"unknown error"}
			json.NewEncoder(w).Encode(resp)
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resp)
	})

	return router
}
