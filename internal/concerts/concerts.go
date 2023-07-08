package concerts

import (
	"fmt"
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
	GetConcerts() ([]Concert, error)
}

func Gather(concertProvider []ConcertProvider) []Concert {
	concertsAllVenues := []Concert{}
	for _, venue := range concertProvider {
		fmt.Println("Venue: ", venue)
		concerts, _ := venue.GetConcerts()
		for _, concert := range concerts {
			fmt.Println("append concert: ", concert.Venue)
			concertsAllVenues = append(concertsAllVenues, concert)
		}
	}

	return concertsAllVenues
}
