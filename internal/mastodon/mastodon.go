package main

import (
	"context"
	"fmt"
	"log"

	"github.com/mattn/go-mastodon"
)

func main() {
	c := mastodon.NewClient(&mastodon.Config{
		Server:       "https://mastodon.social",
		ClientID:     "weY3vAdXHQSV69eL77mmeT9e-hJJIFoCnTXzJ9SZngc",
		ClientSecret: "d5rzLBuwLWVRjiUm2dnhRGlupXzAvqtLySDPzYqEhpg",
	})

	err := c.Authenticate(context.Background(),
		"h6jjcefu5l0475dt@sinenomine.email",
		"LOB5DReheFkV5YSS")
	if err != nil {
		log.Fatal(err)
	}

	toot := mastodon.Toot{
		Status: "hi there!",
	}

	s, err := c.PostStatus(context.Background(), &toot)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(s)

	timeline, err := c.GetTimelineHome(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	for i := len(timeline) - 1; i >= 0; i-- {
		fmt.Println(timeline[i])
	}
	fmt.Print("hi")
}
