package core

import "context"

type User struct {
	UserID   int    `json:"user_id" gorm:"autoIncrement;primaryKey;column:url_id;type:int"`
	Username string `json:"username" gorm:"column:username;unique;type:string"`
	Password string `json:"-" gorm:"column:password;type:string"`
	Email    string `json:"email" gorm:"column:email;type:string"`
}

func (User) TableName() string {
	return "user"
}

type UserStore interface {
	FindUser(ctx context.Context, name string) *User
}