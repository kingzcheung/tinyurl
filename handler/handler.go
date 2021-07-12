package handler

import (
	"github.com/golang/groupcache"
	"github.com/kingzcheung/tinyurl/config"
	"gorm.io/gorm"
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
