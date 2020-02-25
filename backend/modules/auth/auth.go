package auth

import (
	"context"
	"fmt"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
)

const AuthorizedUserIdKey = "authorized_user_id"

func loadCredFromFile(name string) *option.ClientOption {
	filename := fmt.Sprintf(name)
	opt := option.WithCredentialsFile(filename)
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
		opt = loadCredFromFile("./backend/configs/firebase-service-dev.json")
	}
	return opt
}
