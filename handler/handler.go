package handler

import (
	"github.com/golang/groupcache"
	"gorm.io/gorm"
	"github.com/kingzcheung/tinyurl/config"
	"net/http"
)

type Handler interface {
	Handle() http.Handler
	Pattern() string
}

type Handle struct {
	DB     *gorm.DB
	Config *config.Config
	Cache  *groupcache.Group
}
