package config 

import (
	"fmt"
	"github.com/spf13/viper"
)

var config Config

type Config struct {
	PostgreSQL PostgreSQL
	Redis      Redis
	JwtToken   JwtToken
}

type PostgreSQL struct {
	Username string
	Password string
	Host     string
	Port     string
	Name     string
	Debug    bool
	SSL      string
}

type Redis struct {
	Host string
	Port string
}

type JwtToken struct {
	AccessTokenSecretKey  string
	RefreshTokenSecretKey string
	AccessTokenMaxAge     int
	RefreshTokenMaxAge    int
}

func Get() Config {
	return config
}

func LoadConfig(path string) (Config, error) {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(path)     // path to look for the config file in

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	err = viper.Unmarshal(&config)

	return config, err
}