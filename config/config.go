package config

import (
	"encoding/json"
	"io/ioutil"
	"log"

	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

var (
	conf Config
)

func init() {
	log.SetOutput(&lumberjack.Logger{
		Filename:   "logs/app.log",
		MaxSize:    10, // megabytes
		MaxBackups: 7,
		MaxAge:     40,   //days
		Compress:   true, // disabled by default
	})

	log.Println("-------- * ------- Starting Logging -------- * -------")
}

// Parce config from config file
func Parce(path string) (*Config, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(b, &conf); err != nil {
		return nil, err
	}

	return &conf, nil
}

func Peek() *Config {
	return &conf
}

type Config struct {
	Server    Server    `json:"server"`
	Database  database  `json:"database"`
	Configure configure `json:"configure"`
}

type Server struct {
	Addr string `json:"addr"`
	Name string `json:"name"`
}

type database struct {
	Addr       string `json:"addr"`
	User       string `json:"user"`
	Pass       string `json:"pass"`
	DBName     string `json:"dbname"`
	ConnString string `json:"connString"`
}

type configure struct {
	Addr string `json:"addr"`
}
