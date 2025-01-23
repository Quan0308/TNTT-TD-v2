package interfaces

import "github.com/Quan0308/main-api/models"

type FirebaseAuthService interface {
	CheckAuthentication(email string, password string)
	UpdatePassword(firebaseUid string)
	GeneratePairTokens(firebaseUid string, claims models.CurrentUser) (string, string)
	ValidateToken(token string) bool
}
