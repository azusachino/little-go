package main

import (
	"fmt"
	"github.com/azusachino/golong/practices/viper/app"
	"github.com/spf13/viper"
	"log"
)

func main() {
	viper.SetConfigName("app")
	viper.AddConfigPath("practices/viper/conf/")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("parse conf failed: %v", err)
	}
	_ = viper.Unmarshal(app.Setting)
	fmt.Println(app.Setting.AppName)
}
