package config

import "html/template"

type AppConfig struct {
	TemplateCache map[string]*template.Template

	UseCache bool

	ServerConfig ServerConfiguration
	DBConfig DatabaseConfiguration
}
var EnvConfig *AppConfig

func SetEnvConfig(envConfig *AppConfig) {
	EnvConfig = envConfig
}
type ServerConfiguration struct {
	Port int
}
type DatabaseConfiguration struct {
	DBName string
	DBUser string
	DBPassword string

}