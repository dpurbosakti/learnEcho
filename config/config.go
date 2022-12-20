package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func (a *app) initConfig() {
	// main viper config
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("/etc/" + a.Code)
	viper.AddConfigPath("$HOME/" + a.Code)
	viper.AutomaticEnv()

	// default values
	viper.SetDefault("FullName", "Learn echo")
	viper.SetDefault("Version", "0.0.1")
	viper.SetDefault("HttpConf.Host", "127.0.0.1")
	viper.SetDefault("HttpConf.Port", "8000")

	// read the file
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
		panic(err)
	}

	// map to app
	if err := viper.Unmarshal(a); err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
		panic(err)

	}

	// done
	fmt.Println("Config is loaded successfully")
}
