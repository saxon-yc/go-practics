package config

import "go-practics/internal/model"

type App struct {
	Port      int    `yaml:"port"`
	HTTPS     string `yaml:"https"`
	Pre       string `yaml:"pre"`
	SessionId int    `yaml:"session_id"`
}

type ConfigModel struct {
	Components []model.ComptItem `yaml:"components"`
	App        App               `yaml:"app"`
}
