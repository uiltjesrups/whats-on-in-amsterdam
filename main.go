package main

import (
	"github.com/uiltjesrups/whats-on-in-amsterdam/internal/concerts"
	"github.com/uiltjesrups/whats-on-in-amsterdam/internal/html"
)

func main() {
	venues := []concerts.ConcertProvider{concerts.Occii}
	concerts := concerts.Gather(venues)
	html.WriteHTML(concerts)
}
