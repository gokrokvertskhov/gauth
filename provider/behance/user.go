package box

import (
	"fmt"

	"github.com/tamnd/gauth"
	"github.com/tamnd/httpclient"
)

var endpointProfile = "https://www.behance.net/v2/users/"

func User(token *gauth.AccessToken) (*gauth.User, error) {
	URL := endpointProfile + "?access_token=" + token.Token
	fmt.Println(URL)
	content, err := httpclient.String(URL)
	if err != nil {
		return nil, err
	}

	user := gauth.User{
		Raw: content,
	}

	return &user, nil
}
