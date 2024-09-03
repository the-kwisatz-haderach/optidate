package router

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/the-kwisatz-haderach/optidate/internal/dateservice"
)

func New(service *dateservice.Service) http.Handler {
	router := httprouter.New()

	// Serve frontend SPA
	router.NotFound = http.FileServer(http.Dir("frontend/dist/"))

	router.GET("/countries", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		resp, err := service.GetCountries(r.Context())
		if err != nil {
			JSONError(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resp)
	})

	router.GET("/calendar/:country", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")
		var opts dateservice.CreateCalendarOptions

		countryCode := ps.ByName("country")
		if len(countryCode) != 2 {
			JSONError(w, "invalid country code", http.StatusBadRequest)
			return
		}

		tq := r.URL.Query().Get("threshold")
		if tq != "" {
			threshold, err := strconv.Atoi(tq)
			if err != nil {
				JSONError(w, "invalid value for query 'threshold'", http.StatusBadRequest)
				return
			}
			opts.Threshold = threshold
		}

		yq := r.URL.Query().Get("year")
		if yq != "" {
			year, err := strconv.Atoi(yq)
			if err != nil {
				JSONError(w, "invalid value for query 'days'", http.StatusBadRequest)
				return
			}
			opts.Year = year
		}

		resp, err := service.GetCalendar(r.Context(), countryCode, opts)
		if err != nil {
			JSONError(w, "unknown error", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resp)
	})

	return router
}
