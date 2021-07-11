package user

import (
	"context"
	"github.com/kingzcheung/tinyurl/core"
	"gorm.io/gorm"
)

type userStore struct {
	db *gorm.DB
}

func NewUserStore(db *gorm.DB) core.UserStore {
	return &userStore{db: db}
}

func (u *userStore) FindUser(ctx context.Context, name string) *core.User {
	var user core.User

	u.db.WithContext(ctx).Where("username = ?",name).First(&user)

	return &user
}

