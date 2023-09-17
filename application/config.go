package application

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/spf13/viper"
)

type AppConfig struct {
	RedisPasswd string `json:"REDIS_PASSWORD"`
	AppPort     string `json:"APP_PORT"`
	RedisAddr   string `json:"REDIS_ADDR"`
}

var (
	Config *AppConfig
)

func init() {
	loadConfig()
}
func loadConfig() {
	config := AppConfig{}
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("json")
	viper.AutomaticEnv()
	config.AppPort = viper.GetString("APP_PORT")
	config.RedisPasswd = viper.GetString("REDIS_PASSWORD")
	config.RedisAddr = viper.GetString("REDIS_ADDR")
	Config = &config
}

func GetConfig() *AppConfig {
	return Config
}
