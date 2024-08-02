package dateservice

import (
	"fmt"
	"net/http"

	"github.com/rs/zerolog/log"
)

type Service struct{}

const basepath = "https://date.nager.at/api/v3/"

func New() *Service {
	return &Service{}
}

func (s *Service) GetCountries() {
	requestURL := fmt.Sprintf(basepath + "AvailableCountries")
	_, err := http.Get(requestURL)
	if err != nil {
		log.Err(err).Msg("error fetching available countries")
		return
	}
	fmt.Println(requestURL)
}
