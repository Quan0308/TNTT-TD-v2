package models

import (
	"time"

	userEnum "github.com/Quan0308/main-api/enums/user"
	"github.com/google/uuid"
)

type User struct {
	Id               uuid.UUID       `json:"id" db:"id"`
	FirebaseUid      string          `json:"firebaseUid" db:"firebase_uid"`
	Status           userEnum.Status `json:"status" db:"status"`
	Email            string          `json:"email" db:"email"`
	HolyName         string          `json:"holyName" db:"holy_name"`
	FirstName        string          `json:"firstName" db:"firt_name"`
	LastName         string          `json:"lastName" db:"last_name"`
	Gender           userEnum.Gender `json:"gender" db:"gender"`
	Dob              time.Time       `json:"dob" db:"dob"`
	BaptismDate      time.Time       `json:"baptismDate" db:"baptism_date"`
	EucharistDate    time.Time       `json:"eucharistDate" db:"eucharist_date"`
	ConfirmationDate time.Time       `json:"confirmationDate" db:"confirmation_date"`
	PatronDate       time.Time       `json:"patronDate" db:"patron_date"`
	JoinedDate       time.Time       `json:"joinedDate" db:"joined_date"`
	Address          string          `json:"address" db:"address"`
	Note             string          `json:"note" db:"note"`
	Mobile           string          `json:"mobile" db:"mobile"`
	CreatedAt        time.Time       `json:"createdAt" db:"created_at"`
	UpdatedAt        time.Time       `json:"updatedAt" db:"updated_at"`
	CreatedBy        string          `json:"createdBy" db:"created_by"`
	UpdatedBy        string          `json:"updatedBy" db:"updated_by"`
}

type CurrentUser struct {
	Id    uuid.UUID       `json:"id"`
	Roles []userEnum.Role `json:"roles"`
}
