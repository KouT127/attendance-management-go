package auth

import (
	"context"
	"golang.org/x/oauth2/google"
	"golang.org/x/xerrors"
	"google.golang.org/api/option"
	"os"
)

const AuthorizedUserIDKey = "authorized_user_id"

func loadCredFromJSON() (*option.ClientOption, error) {
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

func NewCredential() (*option.ClientOption, error) {
	opt, err := loadCredFromCtx()
	if opt == nil {
		opt, err = loadCredFromJSON()
	}
	if err != nil || opt == nil {
		return nil, xerrors.New("Load failed")
	}
	return opt, nil
}
