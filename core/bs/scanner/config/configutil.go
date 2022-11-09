/*
Copyright Lemon Corp. All Rights Reserved.

Written by hama
*/

package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func InitConfig(configPath string) Configuration {

	viper.SetConfigType("yaml")
	viper.SetConfigName("config") // name of config file (without extension)

	if configPath != "" {
		viper.AddConfigPath(configPath)
	} else {
		configPath = "../conf/"
		viper.AddConfigPath(configPath)
	}

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	} else {
		fmt.Println(err)
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Error reading config file, %s", err)
	}

	var configuration Configuration
	err2 := viper.Unmarshal(&configuration)

	if err2 != nil {
		fmt.Println("unable to decode into configuration struct, %v", err2)
	}

	return configuration
}
