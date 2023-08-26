package main

import (
	"fmt"

	"github.com/uiltjesrups/whats-on-in-amsterdam/internal/concerts"
	"github.com/uiltjesrups/whats-on-in-amsterdam/internal/config"
	"github.com/uiltjesrups/whats-on-in-amsterdam/internal/html"
	"github.com/uiltjesrups/whats-on-in-amsterdam/internal/mastodon"
	"github.com/uiltjesrups/whats-on-in-amsterdam/internal/utils"
)

func main() {
	config := config.Parse()
	venues := []concerts.ConcertProvider{
		concerts.OCCII{Name: "OCCII", Url: "https://occii.org"},
		concerts.Zaal100{Name: "Zaal100", Url: "https://zaal100.nl"},
		concerts.Sexyland{Name: "SexyLand", Url: "https://www.sexyland.world"},
		concerts.Buiksloterkerk{Name: "Buiksloterkerk", Url: "https://buiksloterkerk.nl/agenda"},
		concerts.Orgelpark{Name: "Orgelpark", Url: "https://www.orgelpark.nl/nl/Agenda"},
		concerts.Roodebioscoop{Name: "De Roode Bioscoop", Url: "https://www.roodebioscoop.nl"},
	}
	concerts := concerts.GroupConcertsByDate(concerts.Gather(venues))

	html.WriteHTML(concerts)

	for _, concert := range concerts[utils.CurrentDate()] {
		message := fmt.Sprintf("Today at %s: %s - %s\n%s\n%s",
			concert.Venue.Name,
			concert.Description,
			concert.Date.Format("15:00"),
			concert.Venue.Url,
			"https://uiltjesrups.github.io/whats-on-in-amsterdam/")
		mastodon.Post(config, message)
	}

}
