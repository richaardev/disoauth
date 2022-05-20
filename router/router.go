package router

import (
	"net/http"
	"os"
	"strings"

	"golang.org/x/oauth2"
)

var conf *oauth2.Config

func Router() {
	conf = &oauth2.Config{
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		RedirectURL:  os.Getenv("REDIRECT_URI"),
		Scopes:       strings.Split(os.Getenv("SCOPES"), ","),
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://discord.com/api/oauth2/authorize",
			TokenURL: "https://discord.com/api/oauth2/token",

			AuthStyle: oauth2.AuthStyleInParams,
		},
	}

	http.HandleFunc("/me/", CheckAuth(Me))
	http.HandleFunc("/me/guilds/", CheckAuth(Guilds))
	http.HandleFunc("/me/connections/", CheckAuth(Connections))
	http.HandleFunc("/auth/", Auth)
	http.HandleFunc("/auth/callback", Callback)
}
