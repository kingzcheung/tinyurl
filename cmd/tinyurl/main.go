package main

import (
	"flag"
	"github.com/kingzcheung/tinyurl/config"
	"log"
	"net/http"
)

func main() {

	var (
		conf string
	)
	flag.StringVar(&conf, "conf", "", "config file p")
	flag.Parse()

	var c = config.New(conf)

	app, err := InitializeApp(c)

	if err != nil {
		panic(err)
	}
	log.Printf("Server: http://127.0.0.1:%d",c.Server.Http.Port)
	log.Fatalln(app.server.ListenAndServe())
}

type app struct {
	config *config.Config
	server *http.Server
}

func newApp(config *config.Config, server *http.Server) *app {
	return &app{config: config, server: server}
}
