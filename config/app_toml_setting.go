package config

import (
	"github.com/BurntSushi/toml"
	"github.com/sirupsen/logrus"
)

type Database struct {
	Host string `toml:"Host"`
	Type string `toml:"Type"`
	User string `toml:"User"`
	Pwd  string `toml:"Pwd"`
	Name string `toml:"Name"`
}
type Server struct {
	HttpPort        string `toml:"HttpPort"`
	JwtHeader       string `toml:"JwtHeader"`
	TokenExpireTime int    `toml:"TokenExpireTime"`
	TokenSecretKey  string `toml:"TokenSecretKey"`
}
type Redis struct {
	Host string `toml:"Host"`
	User string `toml:"User"`
	Pwd  string `toml:"Pwd"`
}
type TomlConfig struct {
	Server   Server   `toml:"Server"`
	Database Database `toml:"Database"`
	Redis    Redis    `toml:"Redis"`
	Log      Log      `toml:"Log"`
}
type Log struct {
	LogPath string `toml:"Log"`
	Level   string `toml:"Level"`
}

var App = TomlConfig{}

func ConfigSetup() {
	const conf = "conf/app.toml"
	_, err := toml.DecodeFile(conf, &App)
	if err != nil {
		logrus.Errorf("toml setup error[%v],while decodefile[%v]", err, conf)

	}
	logrus.Println("config 是", App)
}
