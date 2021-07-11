package config

import (
	"bytes"
	_ "embed"
	"fmt"
	"github.com/spf13/viper"
	"path"
	"strings"
)

//go:embed config.yml
var appIni []byte

type Config struct {
	AppName  string `mapstructure:"app_name"`
	AppMode  string `mapstructure:"app_mode"`
	Database *dbConfig
	Server   *serverConfig
}

func New(filePath string) *Config {
	var err error
	var configName = "config"
	var configType = "yml"
	if filePath == "" {
		dir, filename := path.Split(filePath)
		viper.AddConfigPath(dir)
		fn := strings.Split(filename, ".")
		if len(fn) == 2 {
			configName = fn[0]
			configType = fn[1]
		}
	}
	viper.SetConfigName(configName)
	viper.SetConfigType(configType)
	viper.AddConfigPath("/etc/tinyurl/")
	viper.AddConfigPath("$HOME/.tinyurl")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./etc")

	//default
	_ = viper.ReadConfig(bytes.NewReader(appIni))
	err = viper.MergeInConfig()
	if err != nil {
		err = viper.ReadConfig(bytes.NewReader(appIni))
		if err != nil {
			panic(err)
		}
	}

	var c Config

	err = viper.Unmarshal(&c)
	if err != nil {
		panic(fmt.Sprintf("unable to decode into struct, %v", err))
	}

	return &c
}

type dbConfig struct {
	Drive      string
	Datasource string
	Debug      bool
}

type serverConfig struct {
	Key         string
	AuthType    string `mapstructure:"auth_type"`
	ExternalUrl string `mapstructure:"external_url"`
	Http        *httpServerConfig
	Hash        *hashServerConfig
	Internal    *internalServerConfig
	CsrfToken   string `mapstructure:"csrf_token"`
}

type internalServerConfig struct {
	Admin    string
	Password string
}

type hashServerConfig struct {
	HashSalt      string `mapstructure:"hash_salt"`
	HashMinLength int    `mapstructure:"hash_min_length"`
}

type httpServerConfig struct {
	Port uint16
}
