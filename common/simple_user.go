package common

import "food-client/modules/user/model"

type SimpleUser struct {
	SQLModel  `json:",inline"`
	LastName  string        `json:"last_name" gorm:"column:last_name;"`
	FirstName string        `json:"first_name" gorm:"column:first_name;"`
	Role      string        `json:"role" gorm:"column:role;"`
	Avatar    *model.Avatar `json:"avatar,omitempty" gorm:"column:avatar;type:json"`
}

func (SimpleUser) TableName() string {
	return "users"
}

func (u *SimpleUser) Mask(isAdmin bool) {
	u.GenUID(DbTypeUser)
}
