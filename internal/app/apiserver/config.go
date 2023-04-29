package apiserver

import "github.com/Fox1N69/rest-api/store"

//Config ...
type Config struct{
	BindAdder string `toml:"bind_adder"`
	LogLevel string `toml:"log_level"`
	Store *store.Config
}
 
//NewConfig ...
func NewConfig() *Config {
	return &Config {
		BindAdder: "8080",
		LogLevel: "debug",
		Store: store.NewConfig(),
	}
}