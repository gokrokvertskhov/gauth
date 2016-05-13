package github

import (
	"net/url"

	"github.com/gokrokvertskhov/gauth"
	"github.com/tamnd/httpclient"
)

var endpointProfile = "https://api.github.com/user"

func User(token *gauth.AccessToken) (*gauth.User, error) {
	URL := endpointProfile + "?access_token=" + url.QueryEscape(token.Token)

	u := struct {
		ID      string `json:"login"`
		Email   string `json:"email"`
		Bio     string `json:"bio"`
		Name    string `json:"name"`
		Link    string `json:"html_url"`
		Location string `json:"location"`
	}{}

	err := httpclient.JSON(URL, &u)
	if err != nil {
		return nil, err
	}

	user := gauth.User{
		ID:          u.ID,
		Name:        u.Name,
		Email:       u.Email,
		Username:    u.Name,
		Location:    u.Location,
		Description: u.Bio,
		Raw:         u,
	}

	return &user, nil
}
