package router

import (
	"context"
	"errors"
	"net/http"
	"net/url"
	"os"
	"server/utils"

	"github.com/golang-jwt/jwt"
)

func Auth(w http.ResponseWriter, r *http.Request) {
	state := r.FormValue("state")
	http.Redirect(w, r, conf.AuthCodeURL(state), http.StatusFound)
}

func Callback(w http.ResponseWriter, r *http.Request) {
	token, err := conf.Exchange(context.Background(), r.FormValue("code"))
	state := r.FormValue("state")

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		utils.WriteJsonError(&w, errors.New("invalid code or missing code"))
		return
	}

	encodedToken, err := utils.EncodeJwt(jwt.MapClaims{
		"access_token":  token.AccessToken,
		"token_type":    token.TokenType,
		"expires_in":    token.Expiry.Unix(),
		"refresh_token": token.RefreshToken,
	})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		utils.WriteJsonError(&w, errors.New("failed to encode token: "+err.Error()))
		return
	}
	final_url, err := url.Parse(string(os.Getenv("FINAL_REDIRECT_URL")))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		utils.WriteJsonError(&w, errors.New("failed to parse final redirect url: "+err.Error()))
		return
	}
	query := final_url.Query()
	query.Add("token", encodedToken)
	if state != "" {
		query.Add("state", state)
	}
	final_url.RawQuery = query.Encode()
	
	// if state != "" {
	// 	redirecturl, err := base64.StdEncoding.DecodeString(state)
	// 	if err == nil {
	// 		rurl, err := url.Parse(string(redirecturl))
	// 		if err != nil {
	// 			w.WriteHeader(http.StatusBadRequest)
	// 			utils.WriteJsonError(&w, errors.New("invalid redirect url"))
	// 			return
	// 		}

	// 		query := rurl.Query()
	// 		query.Add("token", encodedToken)
	// 		rurl.RawQuery = query.Encode()

	// 		http.Redirect(w, r, rurl.String(), http.StatusFound)
	// 		return
	// 	}
	// }
	http.Redirect(w, r, final_url.String(), http.StatusFound)
}
