package youtube

import (
	"github.com/radio-t/radio-t-site/publisher/add-to-youtube/client"
	log "github.com/sirupsen/logrus"
	"google.golang.org/api/youtube/v3"
)

func authorize(c *Client) error {
	log.Info("Start authorizing an user")

	client, err := client.New(c.Config.OAuth2, c.Config.TokenPath, false, c.Config.Scopes...)
	if err != nil {
		return errOAuth2HTTPClientCreate(err)
	}

	if _, err = youtube.New(client); err != nil {
		return errYoutubeClientCreate(err)
	}

	log.Info("An user authorized")

	return nil
}