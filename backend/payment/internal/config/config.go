package config

import (
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"os"
)

type Config struct {
	Port                string `mapstructure:"PORT"`
	Host                string `mapstructure:"HOST"`
	PostgreDsn          string `mapstructure:"DB_URL"`
	JwtSecretKey        string `mapstructure:"JWT_SECRET_KEY"`
	JwtExpirationMinute uint64 `mapstructure:"JWT_EXPIRATION_MINUTES"`
}

var envs = []string{
	"DB_HOST", "DB_NAME", "DB_USER", "DB_PORT", "DB_PASSWORD",
	"TWILIO_ACCOUNT_SID", "TWILIO_AUTHTOKEN", "TWILIO_SERVICES_ID",
}

func LoadConfig() (Config, error) {
	var config Config

	viper.AddConfigPath("./")
	if os.Getenv("ENVIRONMENT") == "DEVELOPMENT" {
		viper.SetConfigName("dev")
	} else {
		viper.SetConfigName("prod")
	}
	viper.SetConfigType("env")
	viper.ReadInConfig()

	for _, env := range envs {
		if err := viper.BindEnv(env); err != nil {
			return config, err
		}
	}

	if err := viper.Unmarshal(&config); err != nil {
		return config, err
	}

	if err := validator.New().Struct(&config); err != nil {
		return config, err
	}

	return config, nil
}

func GetEnvVar(name string) string {
	if !viper.IsSet(name) {
		log.Debug().Msgf("Environment variable %s is not set", name)
		return ""
	}
	value := viper.GetString(name)
	return value
}
