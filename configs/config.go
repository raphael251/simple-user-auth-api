package configs

import (
	"fmt"
	"log"

	"github.com/go-chi/jwtauth"
	"github.com/spf13/viper"
)

type conf struct {
	DBDriver           string `mapstructure:"DB_DRIVER"`
	DBHost             string `mapstructure:"DB_HOST"`
	DBPort             string `mapstructure:"DB_PORT"`
	DBName             string `mapstructure:"DB_NAME"`
	DBUser             string `mapstructure:"DB_USER"`
	DBPassword         string `mapstructure:"DB_PASSWORD"`
	DBConnectionString string
	ServerProtocol     string `mapstructure:"SERVER_PROTOCOL"`
	ServerHost         string `mapstructure:"SERVER_HOST"`
	ServerPort         string `mapstructure:"SERVER_PORT"`
	JWTSecret          string `mapstructure:"JWT_SECRET"`
	JWTExpiresIn       int    `mapstructure:"JWT_EXPIRES_IN"`
	TokenAuthKey       *jwtauth.JWTAuth
}

func LoadConfig() (*conf, error) {
	var configs *conf

	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Error trying to connect to the database")
		panic(err)
	}

	err = viper.Unmarshal(&configs)
	if err != nil {
		panic(err)
	}

	configs.DBConnectionString = fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		configs.DBHost, configs.DBPort, configs.DBUser, configs.DBPassword, configs.DBName,
	)

	configs.TokenAuthKey = jwtauth.New("HS256", []byte(configs.JWTSecret), nil)

	return configs, nil
}
