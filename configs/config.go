package configs

import (
	"fmt"
	"os"

	"github.com/caarlos0/env/v6"
	"github.com/go-chi/jwtauth"
	"github.com/joho/godotenv"
)

type conf struct {
	DBDriver           string `mapstructure:"DB_DRIVER" env:"DB_DRIVER,required"`
	DBHost             string `mapstructure:"DB_HOST" env:"DB_HOST,required"`
	DBPort             string `mapstructure:"DB_PORT" env:"DB_PORT,required"`
	DBName             string `mapstructure:"DB_NAME" env:"DB_NAME,required"`
	DBUser             string `mapstructure:"DB_USER" env:"DB_USER,required"`
	DBPassword         string `mapstructure:"DB_PASSWORD" env:"DB_PASSWORD,required"`
	DBConnectionString string
	ServerProtocol     string `mapstructure:"SERVER_PROTOCOL" env:"SERVER_PROTOCOL,required"`
	ServerHost         string `mapstructure:"SERVER_HOST" env:"SERVER_HOST,required"`
	ServerPort         string `mapstructure:"SERVER_PORT" env:"SERVER_PORT,required"`
	JWTSecret          string `mapstructure:"JWT_SECRET" env:"JWT_SECRET,required"`
	JWTExpiresIn       int    `mapstructure:"JWT_EXPIRES_IN" env:"JWT_EXPIRES_IN,required"`
	TokenAuthKey       *jwtauth.JWTAuth
}

func LoadConfig() (*conf, error) {
	configs := conf{}

	/* godotenv will load the .env file variables into the os environment variables.
	 * In docker environment we use environment variables directly, this is why we are not handling the error here.
	 */
	godotenv.Load()

	err := env.Parse(&configs)
	if err != nil {
		panic(err)
	}

	fmt.Println("aobaaaaaaaaaaaaa")
	fmt.Println(os.Getenv("DB_NAME"))

	configs.DBConnectionString = fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		configs.DBHost, configs.DBPort, configs.DBUser, configs.DBPassword, configs.DBName,
	)

	configs.TokenAuthKey = jwtauth.New("HS256", []byte(configs.JWTSecret), nil)

	return &configs, nil
}
