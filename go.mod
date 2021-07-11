module github.com/kingzcheung/tinyurl

go 1.16

require (
	github.com/alexedwards/scs/mysqlstore v0.0.0-20210606090158-85ec2fab6bdf
	github.com/alexedwards/scs/sqlite3store v0.0.0-20210606090158-85ec2fab6bdf
	github.com/alexedwards/scs/v2 v2.4.0
	github.com/go-chi/chi/v5 v5.0.3
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da
	github.com/google/wire v0.5.0
	github.com/matthewhartstonge/argon2 v0.1.4
	github.com/sirupsen/logrus v1.8.1
	github.com/speps/go-hashids v2.0.0+incompatible
	github.com/spf13/viper v1.7.1
	golang.org/x/crypto v0.0.0-20210711020723-a769d52b0f97 // indirect
	golang.org/x/sys v0.0.0-20210630005230-0f9fa26af87c // indirect
	gorm.io/driver/mysql v1.1.1
	gorm.io/driver/sqlite v1.1.4
	gorm.io/gorm v1.21.9
)
