package main

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

type BarConfig struct {
	Baz int `mapstructure:"baz"`
	Too int `mapstructure:"too"`
}

type Config struct {
	Foo int
	Bar BarConfig `mapstructure:"baar"`
}

func main() {
	v := viper.New()
	v.SetEnvPrefix("PREFIX")
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	v.AutomaticEnv()
	v.SetConfigFile("./.config.env")

	err := v.ReadInConfig()
	if err != nil {
		fmt.Println("error: ", err)
	}

	var cfg Config
	err = v.Unmarshal(&cfg)
	if err != nil {
		fmt.Println("error: ", err)
	}

	fmt.Printf("%#v\n", cfg)
}
