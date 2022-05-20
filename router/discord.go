package router

import (
	"errors"
	"net/http"
	"disoauth/utils"

	"github.com/golang-jwt/jwt"
)

func CheckAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		if auth == "" {
			utils.WriteJsonError(&w, errors.New("no authorization header"))
			return
		}

		token, err := utils.DecodeJwt(auth)
		if err != nil {
			utils.WriteJsonError(&w, errors.New("could not decode JWT token:"+err.Error()))
			return
		}
		if !token.Valid {
			utils.WriteJsonError(&w, errors.New("invalid JWT token"))
			return
		}

		next.ServeHTTP(w, r)
	}
}

func Me(w http.ResponseWriter, r *http.Request) {
	auth := r.Header.Get("Authorization")
	token, _ := utils.DecodeJwt(auth)
	claims := token.Claims.(jwt.MapClaims)
	access_token := claims["access_token"].(string)
	user, err := utils.FetchDiscordUser(access_token)
	if err != nil {
		utils.WriteJsonError(&w, errors.New("could not fetch user from discord: "+err.Error()))
		return
	}

	utils.WriteJsonString(&w, user)
}

func Guilds(w http.ResponseWriter, r *http.Request) {
	auth := r.Header.Get("Authorization")
	token, _ := utils.DecodeJwt(auth)

	claims := token.Claims.(jwt.MapClaims)
	access_token := claims["access_token"].(string)
	guilds, err := utils.FetchDiscordGuilds(access_token)
	if err != nil {
		utils.WriteJsonError(&w, errors.New("could not fetch guilds from discord: "+err.Error()))
		return
	}

	utils.WriteJsonString(&w, guilds)
}

func Connections(w http.ResponseWriter, r  *http.Request) {
	auth := r.Header.Get("Authorization")
	token, _ := utils.DecodeJwt(auth)

	claims := token.Claims.(jwt.MapClaims)
	access_token := claims["access_token"].(string)
	connections, err := utils.FetchDiscordConnections(access_token)
	if err != nil {
		utils.WriteJsonError(&w, errors.New("could not fetch connections from discord: "+err.Error()))
		return
	}

	utils.WriteJsonString(&w, connections)
}
