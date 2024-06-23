package provider

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

var (
	token         = Token{}
	AUTH_URL      = os.Getenv("AUTH_URL")
	CLIENT_ID     = os.Getenv("CLIENT_ID")
	CLIENT_SECRET = os.Getenv("CLIENT_SECRET")
)

type Token struct {
	TokenType   string `json:"token_type"`
	AccessToken string `json:"access_token"`
	ExpiresIn   string `json:"expires_in"`
	Token       string `json:"token"`
}

// Query the auth API using the app's client_id and client_secret (from the environment)
// and returns a Token struct on success
func GetToken() (Token, error) {
	if token.AccessToken != "" {
		return token, nil
	}

	values := url.Values{"client_id": {CLIENT_ID}, "client_secret": {CLIENT_SECRET}}
	resp, err := http.PostForm(AUTH_URL+"/api/login?auto_login=true", values)
	if err != nil {
		fmt.Println("[ERROR] http.PostForm:", err.Error())
		return token, err
	}

	data, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		fmt.Println("[ERROR] io.ReadAll:", err.Error())
		return token, err
	}

	err = json.Unmarshal(data, &token)
	if err != nil {
		fmt.Println("[ERROR] json.Unmarshal:", err.Error())
		return token, err
	}

	return token, nil
}
