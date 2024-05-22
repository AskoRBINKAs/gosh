package main

import (
	"encoding/json"
	"errors"
	"os"
)

type ColorSettings struct {
	UsernameColor string
	HostnameColor string
	PwdColor      string
	InputColor    string
}

type Config struct {
	UserColors      ColorSettings
	UserAliases     map[string]string
	EnvironmentVars map[string]string
	BuiltInColors   map[string]string `json:"-"`
}

func LoadConfig() (Config, error) {
	conf := Config{}
	conf.UserAliases = make(map[string]string)

	dir, _ := os.UserHomeDir()
	file, err := os.Open(dir + "/.gosh.json")
	if err != nil {
		return conf, errors.New("failed to read or open config.json")
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	e := decoder.Decode(&conf)
	if e != nil {
		return conf, errors.New("failed to load config. check file")
	}
	//init colors
	conf.BuiltInColors = make(map[string]string)
	conf.BuiltInColors["reset"] = "\033[0m"
	conf.BuiltInColors["red"] = "\033[31m"
	conf.BuiltInColors["green"] = "\033[32m"
	conf.BuiltInColors["yellow"] = "\033[33m"
	conf.BuiltInColors["blue"] = "\033[34m"
	conf.BuiltInColors["purple"] = "\033[35m"
	conf.BuiltInColors["cyan"] = "\033[36m"
	conf.BuiltInColors["white"] = "\033[37m"
	userAliases = config.UserAliases
	return conf, nil
}

func SaveConfig(c Config) {
	dir, _ := os.UserHomeDir()
	file, err := os.OpenFile(dir+"/.gosh.json", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return
	}
	defer file.Close()
	b, e := json.Marshal(c)
	if e != nil {
		return
	}
	file.Write(b)
}
