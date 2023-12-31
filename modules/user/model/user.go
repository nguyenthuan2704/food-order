package model

import (
	"errors"
	"food-client/common"
	"food-client/component/tokenprovider"
	"time"
)

const (
	EntityName = "User"
)

/*var (
	ErrTitleIsBlank = errors.New("title cannot be blank!")
	ErrItemDeleted  = errors.New("item is deleted!")
	RecordNotFound  = errors.New("record not found")
)*/

type User struct {
	common.SQLModel `json:",inline"`
	Email           string      `json:"email" gorm:"column:email;"`
	Password        string      `json:"-" gorm:"column:password;"`
	LastName        string      `json:"last_name" gorm:"column:last_name;"`
	FirstName       string      `json:"first_name" gorm:"column:first_name;"`
	Phone           string      `json:"phone" gorm:"column:phone;"`
	Role            string      `json:"role" gorm:"column:role;"`
	Avatar          *Avatar     `json:"avatar,omitempty" gorm:"column:avatar;type:json"`
	Status          *UserStatus `json:"status" gorm:"-"`
}

/*func (User) TableName() string { return "users" }*/

func (u *User) GetUserId() int {
	return u.Id
}

func (u *User) GetEmail() string {
	return u.Email
}

func (u *User) GetRole() string {
	return u.Role
}

func (User) TableName() string {
	return "users"
}

func (u *User) Mask(isAdmin bool) {
	u.GenUID(common.DbTypeUser)
}

type UserCreation struct {
	common.SQLModel `json:",inline"`
	Email           string      `json:"email" gorm:"column:email;"`
	Password        string      `json:"-" gorm:"column:password;"`
	LastName        string      `json:"last_name" gorm:"column:last_name;"`
	FirstName       string      `json:"first_name" gorm:"column:first_name;"`
	Phone           string      `json:"phone" gorm:"column:phone;"`
	Role            string      `json:"role" gorm:"column:role;"`
	Avatar          *Avatar     `json:"avatar,omitempty" gorm:"column:avatar;type:json"`
	Status          *UserStatus `json:"status" gorm:"-"`
	CreatedAt       *time.Time  `json:"created_at" gorm:"column:created_at;"`
}

/*func (UserCreation) TableName() string { return User{}.TableName() }*/

func (UserCreation) TableName() string {
	return User{}.TableName()
}

func (u *UserCreation) Mask(isAdmin bool) {
	u.GenUID(common.DbTypeUser)
}

type UserLogin struct {
	Email    string `json:"email" form:"email" gorm:"column:email;"`
	Password string `json:"password" form:"password" gorm:"column:password;"`
}

func (UserLogin) TableName() string {
	return User{}.TableName()
}

type Account struct {
	AccessToken  *tokenprovider.Token `json:"access_token"`
	RefreshToken *tokenprovider.Token `json:"refresh_token"`
}

func NewAccount(at, rt *tokenprovider.Token) *Account {
	return &Account{
		AccessToken:  at,
		RefreshToken: rt,
	}
}

var (
	ErrUsernameOrPasswordInvalid = common.NewCustomError(
		errors.New("username or password invalid"),
		"username or password invalid",
		"ErrUsernameOrPasswordInvalid",
	)

	ErrEmailExisted = common.NewCustomError(
		errors.New("email has already existed"),
		"email has already existed",
		"ErrEmailExisted",
	)
)
