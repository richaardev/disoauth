package utils

import (
	"io/ioutil"
	"net/http"
)

func FetchDiscordUser(token string) (string, error) {
	req, err := http.NewRequest("GET", "https://discordapp.com/api/users/@me", nil)
	if err != nil {
		return "", err
	}
	headers := req.Header
	headers.Add("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func FetchDiscordGuilds(token string) (string, error) {
	req, err := http.NewRequest("GET", "https://discordapp.com/api/users/@me/guilds", nil)
	if err != nil {
		return "", err
	}
	headers := req.Header
	headers.Add("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func FetchDiscordConnections(token string) (string, error) {
	req, err := http.NewRequest("GET", "https://discordapp.com/api/users/@me/connections", nil)
	if err != nil {
		return "", err
	}
	headers := req.Header
	headers.Add("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}