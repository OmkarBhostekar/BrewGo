package util

import "github.com/spf13/viper"

type Config struct {
	Environment            string `mapstructure:"ENVIRONMENT"`
	RabbitMQAddress        string `mapstructure:"RABBITMQ_ADDRESS"`
	ProductServiceEndPoint string `mapstructure:"PRODUCT_SERVICE_ENDPOINT"`
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
