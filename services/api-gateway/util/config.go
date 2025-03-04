package util

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Environment            string        `mapstructure:"ENVIRONMENT"`
	UserServiceEndPoint    string        `mapstructure:"USER_SERVICE_ENDPOINT"`
	ProductServiceEndPoint string        `mapstructure:"PRODUCT_SERVICE_ENDPOINT"`
	OrderServiceEndPoint   string        `mapstructure:"ORDER_SERVICE_ENDPOINT"`
	TokenSymmetricKey      string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration    time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	RefreshTokenDuration   time.Duration `mapstructure:"REFRESH_TOKEN_DURATION"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return
}
