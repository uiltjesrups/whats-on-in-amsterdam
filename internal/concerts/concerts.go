package concerts

import (
	"log"
	"sort"
	"time"

	"github.com/uiltjesrups/whats-on-in-amsterdam/internal/utils"
)

type Venue struct {
	Name            string
	Url             string
	ConcertProvider ConcertProvider
}

type ConcertProvider interface {
	GetConcerts() []Concert
}

type Concert struct {
	Description string
	Date        time.Time
	Venue       Venue
}

type titleProvider interface {
	getTitle() (string, error)
}

type dateProvider interface {
	getDate() (time.Time, error)
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

type GroupedConcerts map[time.Time][]Concert

func GroupConcertsByDate(concerts []Concert) GroupedConcerts {
	groupedConcerts := make(GroupedConcerts)

	for _, concert := range concerts {
		date := concert.Date.UTC().Truncate(24 * time.Hour)
		today := utils.CurrentDate()
		if date.After(today) || date.Equal(today) {
			concerts := append(groupedConcerts[date], concert)
			groupedConcerts[date] = concerts
		}
	}

	for _, concerts := range groupedConcerts {
		sort.Slice(concerts, func(i, j int) bool {
			return concerts[i].Date.Before(concerts[j].Date)
		})
	}

	return groupedConcerts
}
