package auth

import (
	"context"
	"net/http"
	"os"

	"golang.org/x/oauth2/clientcredentials"
)

func GetOAuth2Token() *http.Client {
	config := &clientcredentials.Config{
		ClientID:     os.Getenv("SHAREPOINT_CLIENT_ID"),
		ClientSecret: os.Getenv("SHAREPOINT_CLIENT_SECRET"),
		TokenURL:     "https://login.microsoftonline.com/{tenant}/oauth2/v2.0/token",
		Scopes:       []string{"https://graph.microsoft.com/.default"},
	}

	return config.Client(context.Background())
}
