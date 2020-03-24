package auth

import (
	"context"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"os"
)

const AuthorizedUserIdKey = "authorized_user_id"

func loadCredFromJson() *option.ClientOption {
	json := os.Getenv("FIREBASE_SERVICE_JSON")
	cred, err := google.CredentialsFromJSON(context.Background(), []byte(json))
	if err != nil {
		return nil
	}
	opt := option.WithCredentials(cred)
	return &opt
}

func loadCredFromCtx() *option.ClientOption {
	cred, err := google.FindDefaultCredentials(context.Background())
	if err != nil {
		return nil
	}
	opt := option.WithCredentials(cred)
	return &opt
}

func NewCredential() *option.ClientOption {
	opt := loadCredFromCtx()
	if opt == nil {
		opt = loadCredFromJson()
	}
	return opt
}
