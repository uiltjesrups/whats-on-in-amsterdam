package main

import (
	"fmt"
	"time"

	"github.com/uiltjesrups/whats-on-in-amsterdam/internal/concerts"
	"github.com/uiltjesrups/whats-on-in-amsterdam/internal/config"
	"github.com/uiltjesrups/whats-on-in-amsterdam/internal/mastodon"
)

func main() {
	config := config.Parse()
	venues := []concerts.ConcertProvider{
		concerts.OCCII{Name: "OCCII", Url: "https://occii.org"},
		concerts.Zaal100{Name: "Zaal100", Url: "https://zaal100.nl"},
		concerts.Sexyland{Name: "SexyLand", Url: "https://www.sexyland.world"},
	}
	concerts := concerts.GroupConcertsByDate(concerts.Gather(venues))

	for _, concert := range concerts[currentDate()] {
		message := fmt.Sprintf("Today at %s: %s - %s\n%s\n%s",
			concert.Venue.Name,
			concert.Description,
			concert.Date.Format("15:00"),
			concert.Venue.Url,
			"https://uiltjesrups.github.io/whats-on-in-amsterdam/")
		mastodon.Post(config, message)
	}

}

func currentDate() time.Time {
	currentTime := time.Now()
	currentDate := currentTime.UTC().Truncate(24 * time.Hour)
	return currentDate
}

func addOneDay(date time.Time) time.Time {
	oneDay := 24 * time.Hour
	newDate := date.Add(oneDay)
	return newDate
}
