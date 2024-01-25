package utils

import "github.com/spf13/viper"

type Config struct {
	DBDriver              string `mapstructure:"DB_DRIVER"`
	DB_SOURCE             string `mapstructure:"DB_SOURCE"`
	SERVER_ADDRESS        string `mapstructure:"SERVER_ADDRESS"`
	RDSSource             string `mapstructure:"RDS_SOURCE"`
	MQTTEndpoint          string `mapstructure:"MQTT_ENDPOINT"`
	MQTTClientID          string `mapstructure:"MQTT_CLIENT_ID"`
	MQTTPort              string `mapstructure:"MQTT_PORT"`
	TwilioAccountSID      string `mapstructure:"TWILIO_ACCOUNT_SID"`
	TwilioAuthToken       string `mapstructure:"TWILIO_AUTH_TOKEN"`
	TwilioPhoneNumberFrom string `mapstructure:"TWILIO_PHONE_NUMBER_FROM"`
	TwilioPhoneNumberTo   string `mapstructure:"TWILIO_PHONE_NUMBER_TO"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	//viper.SetConfigType("env")

	viper.AutomaticEnv()

	//err = viper.ReadInConfig()
	//if err != nil {
	//	return
	//}
	//
	//err = viper.Unmarshal(&config)
	return
}
