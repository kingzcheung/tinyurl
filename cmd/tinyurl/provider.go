package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/alexedwards/scs/mysqlstore"
	"github.com/alexedwards/scs/sqlite3store"
	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi/v5"
	"github.com/golang/groupcache"
	"github.com/google/wire"
	"github.com/kingzcheung/tinyurl/config"
	"github.com/kingzcheung/tinyurl/core"
	"github.com/kingzcheung/tinyurl/core/migrate"
	"github.com/kingzcheung/tinyurl/handler/api"
	"github.com/kingzcheung/tinyurl/handler/auth"
	"github.com/kingzcheung/tinyurl/handler/health"
	"github.com/kingzcheung/tinyurl/handler/web"
	"github.com/kingzcheung/tinyurl/store/url"
	"github.com/kingzcheung/tinyurl/store/user"
	"github.com/sirupsen/logrus"
	"github.com/speps/go-hashids"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net/http"
	"time"
)

var storeSet = wire.NewSet(
	provideDb,
	user.NewUserStore,
	url.NewUrlStore,
)

var serverSet = wire.NewSet(
	providerHashIds,
	provideUrlCache,
	healthProvide,
	auth.NewServer,
	api.NewServer,
	web.NewServer,
	provideRouter,
	provideSession,
	provideServer,
)

func providerHashIds(config *config.Config) (*hashids.HashID, error) {
	hd := hashids.NewData()
	hd.Salt = config.Server.Hash.HashSalt
	hd.MinLength = config.Server.Hash.HashMinLength
	return hashids.NewWithData(hd)
}

func healthProvide() health.Health {
	return "OK"
}
func provideUrlCache(store core.UrlStore) *groupcache.Group {
	var urlGroup = groupcache.NewGroup("url", 1<<20, groupcache.GetterFunc(func(
		ctx context.Context,
		key string,
		dest groupcache.Sink,
	) error {
		code := store.GetByCode(nil, key)
		if code.ExpiredAt > 0 {
			if code.ExpiredAt < time.Now().Unix() {
				logrus.WithField("url", code).Error("The URL has expired.")
				return fmt.Errorf("source null , key: %v", key)
			}
		}
		if code == nil {
			return fmt.Errorf("source null , key: %v", key)
		}
		marshal, err := json.Marshal(code)
		if err != nil {
			return err
		}
		return dest.SetBytes(marshal)
	}))
	return urlGroup
}

func provideDb(config *config.Config) (*gorm.DB, error) {
	var dial gorm.Dialector
	switch config.Database.Drive {
	case "mysql":
		dial = mysql.Open(config.Database.Datasource)
	case "sqlite3":
		dial = sqlite.Open(config.Database.Datasource)
	default:
		dial = sqlite.Open(config.Database.Datasource)

	}
	if config.Database.Drive == "sqlite3" {
		dial = sqlite.Open(config.Database.Datasource)
	}

	db, err := gorm.Open(dial, &gorm.Config{})
	if err != nil {
		return nil, err
	}
	if config.Database.Debug {
		db = db.Debug()
	}

	err = _initDb(config, db)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func _initDb(c *config.Config, db *gorm.DB) error {
	err := migrate.Migrate(db)
	if err != nil {
		return err
	}

	err = migrate.InitUser(c, db)
	if err != nil {
		return err
	}
	return nil
}

func provideRouter(api *api.Server, health health.Health, web *web.Server, auth *auth.Server) *chi.Mux {
	r := chi.NewRouter()
	r.Mount(api.Pattern(), api.Handle())
	r.Mount(health.Pattern(), health.Handle())
	r.Mount(web.Pattern(), web.Handle())
	r.Mount(auth.Pattern(), auth.Handle())

	//printRoute(r.Routes())
	return r
}

func printRoute(routes []chi.Route) {
	for _, route := range routes {
		for _, subroute := range route.SubRoutes.Routes() {
			fmt.Printf("%s%s\n", route.Pattern, subroute.Pattern)
		}
	}
}

func provideSession(cnf *config.Config, gormDB *gorm.DB) (*scs.SessionManager, error) {
	//db, err := bbolt.Open("/tmp/session.db", 0600, nil)
	//if err != nil {
	//	return nil, err
	//}

	db, err := gormDB.DB()
	if err != nil {
		return nil, err
	}
	sessionManager := scs.New()
	sessionManager.Lifetime = 24 * time.Hour
	sessionManager.Cookie.Name = "TR_SESSION_ID"
	if cnf.Database.Drive == "sqlite3" {
		sessionManager.Store = sqlite3store.New(db)
	} else {
		sessionManager.Store = mysqlstore.New(db)
	}

	return sessionManager, nil
}

func provideServer(handler *chi.Mux, config *config.Config, manager *scs.SessionManager) *http.Server {
	//cors := middleware.Cors(handler)

	return &http.Server{
		Handler: manager.LoadAndSave(handler),
		Addr:    fmt.Sprintf(":%d", config.Server.Http.Port),
	}
}
