package migrate

import (
	"github.com/kingzcheung/tinyurl/config"
	"github.com/kingzcheung/tinyurl/core"
	"github.com/matthewhartstonge/argon2"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&core.Url{},
		&core.Tracking{},
		&core.Analytics{},
		&core.User{},
		&core.Session{},
	)
}

func InitUser(c *config.Config, db *gorm.DB) error {
	var user core.User
	name := c.Server.Internal.Admin
	pwd, err := generatePwd(c.Server.Internal.Password)
	if err != nil {
		return err
	}
	return db.Where(core.User{Username: name}).
		Attrs(core.User{Password: pwd,Username: name}).
		FirstOrCreate(&user).Error
}

func generatePwd(password string) (string, error) {
	argon := argon2.DefaultConfig()
	encoded, err := argon.HashEncoded([]byte(password))
	if err != nil {
		return "", err
	}
	return string(encoded), nil
}
