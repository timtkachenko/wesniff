package models

import (
	"github.com/google/uuid"
	"time"
)

type UserInfo struct {
	Email     string `json:"email"`
	Nickname  string `json:"nickname"`
	FirstName string `json:"given_name"`
	LastName  string `json:"family_name"`
	Name      string `json:"name"`
}

type User struct {
	ID              uuid.UUID  `gorm:"primary_key;type:uuid;unique;" json:"uuid"`
	Email           string     `json:"email"`
	Secret          string     `json:"secret,omitempty"`
	Token           string     `json:"token,omitempty"`
	FirstName       string     `json:"first_name"`
	LastName        string     `json:"last_name"`
	ActiveAccountID uuid.UUID  `json:"active_account_id"`
	BlockedAt       *time.Time `json:"-"`
	CreatedAt       time.Time  `json:"-"`
	UpdatedAt       time.Time  `json:"-"`
	//
}

func (u *User) BeforeSave() (err error) {
	now := time.Now().UTC()
	u.CreatedAt = now
	u.UpdatedAt = u.CreatedAt
	return
}

func (u *User) BeforeUpdate() (err error) {
	u.UpdatedAt = time.Now().UTC()
	return
}

func (u *User) ToDto(accountId ...uuid.UUID) *UserDto {
	ui := &UserDto{
		ID:    u.ID.String(),
		Email: u.Email,
	}
	return ui
}

func UserFromDto(dto *UserCreateDto) (user *User) {
	user = &User{
		ID:        uuid.New(),
		Secret:    "",
		FirstName: "",
		LastName:  "",
	}

	return
}
