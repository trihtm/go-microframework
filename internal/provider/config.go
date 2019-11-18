package provider

import (
	"flag"
	"fmt"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"go-microframework/internal/config"
)

func NewConfig() (cf config.Configuration) {
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	_ = viper.BindPFlags(pflag.CommandLine)

	configPath := viper.GetString("config-path")

	if configPath == "" {
		configPath = "./configs"
	}

	viper.SetConfigType("yml")
	viper.SetConfigName("config")
	viper.AddConfigPath(configPath)

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	err = viper.Unmarshal(&cf)
	if err != nil {
		panic(fmt.Errorf("Unable to decode into struct, %v", err))
	}

	return cf
}

func InitConfig(cf config.Configuration) {

}
