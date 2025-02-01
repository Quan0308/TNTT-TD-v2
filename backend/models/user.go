package models

import (
	"database/sql"
	"time"

	userEnum "github.com/Quan0308/main-api/enums/user"
	"github.com/google/uuid"
)

type User struct {
	Id               uuid.UUID        `json:"id" db:"id"`
	FirebaseUid      sql.NullString   `json:"firebaseUid" db:"firebase_uid"`
	Status           *userEnum.Status `json:"status" db:"status"`
	Email            sql.NullString   `json:"email" db:"email"`
	HolyName         sql.NullString   `json:"holyName" db:"holy_name"`
	FirstName        sql.NullString   `json:"firstName" db:"first_name"`
	LastName         sql.NullString   `json:"lastName" db:"last_name"`
	Gender           *userEnum.Gender `json:"gender" db:"gender"`
	Dob              sql.NullTime     `json:"dob" db:"dob"`
	BaptismDate      sql.NullTime     `json:"baptismDate" db:"baptism_date"`
	EucharistDate    sql.NullTime     `json:"eucharistDate" db:"eucharist_date"`
	ConfirmationDate sql.NullTime     `json:"confirmationDate" db:"confirmation_date"`
	PatronDate       sql.NullTime     `json:"patronDate" db:"patron_date"`
	JoinedDate       sql.NullTime     `json:"joinedDate" db:"joined_date"`
	Address          sql.NullString   `json:"address" db:"address"`
	Note             sql.NullString   `json:"note" db:"note"`
	Mobile           sql.NullString   `json:"mobile" db:"mobile"`
	CreatedOn        *time.Time       `json:"createdOn" db:"created_on"`
	UpdatedOn        *time.Time       `json:"updatedOn" db:"updated_on"`
	CreatedBy        sql.NullString   `json:"createdBy" db:"created_by"`
	UpdatedBy        sql.NullString   `json:"updatedBy" db:"updated_by"`
	CroppedAvatarId  sql.NullString   `json:"croppedAvatarId" db:"cropped_avatar_id"`
	AvatarId         sql.NullString   `json:"avatarId" db:"avatar_id"`
	CoverId          sql.NullString   `json:"coverId" db:"cover_id"`
}

type CurrentUser struct {
	Id    uuid.UUID       `json:"id"`
	Roles []userEnum.Role `json:"roles"`
}
