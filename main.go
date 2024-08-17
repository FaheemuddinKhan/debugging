package main

import (
	"fmt"
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Println("Config file not found")
		} else {
			log.Printf("unexpected error occured %v", err)
		}
	}
	log.Println("Config file found and successfully parsed")

	fmt.Println(viper.Get("env.name"))

	viper.Set("env.name", "E2")
	fmt.Println(viper.Get("env.name"))

	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
		fmt.Println(viper.Get("env"))
		fmt.Println(viper.AllSettings())
	})
	viper.WatchConfig()

	<-make(chan struct{})
}
