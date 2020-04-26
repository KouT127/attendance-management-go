package auth

import (
	"context"
	"golang.org/x/oauth2/google"
	"golang.org/x/xerrors"
	"google.golang.org/api/option"
	"os"
)

const AuthorizedUserIdKey = "authorized_user_id"

func loadCredFromJson() (*option.ClientOption, error) {
	json := os.Getenv("FIREBASE_SERVICE_JSON")
	cred, err := google.CredentialsFromJSON(context.Background(), []byte(json))
	if err != nil {
		return nil, err
	}
	opt := option.WithCredentials(cred)
	return &opt, err
}

func loadCredFromCtx() (*option.ClientOption, error) {
	cred, err := google.FindDefaultCredentials(context.Background())
	if err != nil {
		return nil, err
	}
	opt := option.WithCredentials(cred)
	return &opt, err
}

func NewCredential() *option.ClientOption {
	opt, err := loadCredFromCtx()
	if opt == nil {
		opt, err = loadCredFromJson()
	}
	if err != nil || opt == nil {
		xerrors.New("Load failed")
	}
	return opt
}
