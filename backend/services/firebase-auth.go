package services

import (
	"context"
	"log"

	"firebase.google.com/go"

	"firebase.google.com/go/auth"
	"github.com/Quan0308/main-api/config"
	"github.com/Quan0308/main-api/interfaces"
	"github.com/Quan0308/main-api/models"
	"google.golang.org/api/option"
)

type FireBaseAuthServiceImpl struct {
	app    *firebase.App
	client *auth.Client
}

func NewFireBaseAuthServiceImpl() interfaces.FirebaseAuthService {
	opt := option.WithCredentialsFile(config.Envs.SI_SERVICE_ACCOUNT)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatal("Error init firebase app", err)
		return nil
	}
	client, err := app.Auth(context.Background())
	if err != nil {
		log.Fatal("Error init auth client", err)
	}
	return &FireBaseAuthServiceImpl{app, client}
}

func (f *FireBaseAuthServiceImpl) VerifyIdToken(ctx context.Context, idToken string) (*auth.Token, error) {
	token, err := f.client.VerifyIDToken(ctx, idToken)
	if err != nil {
		log.Printf("error verifying ID token: %v\n", err)
		return nil, err
	}
	return token, nil
}

func (f *FireBaseAuthServiceImpl) GenerateCustomToken(ctx context.Context, uid string, claims *models.CurrentUser) (string, error) {

	claimsMap := map[string]interface{}{
		"id":    claims.Id,
		"roles": claims.Roles,
	}

	token, err := f.client.CustomTokenWithClaims(ctx, uid, claimsMap)

	if err != nil {
		log.Printf("Error generating custom token: %v", err)
		return "", err
	}

	return token, nil
}
