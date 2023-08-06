package mastodon

import (
	"context"
	"log"

	"github.com/mattn/go-mastodon"
	"github.com/uiltjesrups/whats-on-in-amsterdam/internal/config"
)

func Post(config config.Config, status string) {
	log.Println("Post: ", status)
	c := mastodon.NewClient(&mastodon.Config{
		Server:       config.Mastodon.Server,
		ClientID:     config.Mastodon.ClientID,
		ClientSecret: config.Mastodon.ClientSecret,
	})

	err := c.Authenticate(context.Background(),
		config.Mastodon.Email,
		config.Mastodon.Password)
	if err != nil {
		log.Fatal(err)
	}

	toot := mastodon.Toot{
		Status: status,
	}

	_, err = c.PostStatus(context.Background(), &toot)
	if err != nil {
		log.Fatal(err)
	}

}
