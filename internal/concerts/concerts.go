package concerts

import (
	"log"
	"time"
)

type Venue struct {
	Name string
	Url  string
}

type Concert struct {
	Description string
	Date        time.Time
	Venue       Venue
}

type ConcertProvider interface {
	GetConcerts() []Concert
}

func Gather(concertProvider []ConcertProvider) []Concert {
	concertsAllVenues := []Concert{}

	for _, venue := range concertProvider {
		concerts := venue.GetConcerts()
		for _, concert := range concerts {
			log.Println("Gather: ", concert)
			concertsAllVenues = append(concertsAllVenues, concert)
		}
	}

	return concertsAllVenues
}
