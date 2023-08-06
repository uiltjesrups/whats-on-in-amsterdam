package main

import (
	"github.com/uiltjesrups/whats-on-in-amsterdam/internal/concerts"
	"github.com/uiltjesrups/whats-on-in-amsterdam/internal/config"
	"github.com/uiltjesrups/whats-on-in-amsterdam/internal/html"
	"github.com/uiltjesrups/whats-on-in-amsterdam/internal/mastodon"
)

func main() {
	config := config.Parse()
	venues := []concerts.ConcertProvider{
		concerts.OCCII{Name: "OCCI", Url: "https://occii.org/"},
		concerts.Zaal100{Name: "Zaal100", Url: "https://zaal100.nl/"},
		concerts.Sexyland{Name: "SexyLand", Url: "https://www.sexyland.world"},
	}
	concerts := concerts.Gather(venues)

	html.WriteHTML(concerts)

	mastodon.Post(config)
}
