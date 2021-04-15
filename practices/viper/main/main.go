package main

import (
	"fmt"
	"github.com/spf13/viper"
)

func init() {
	viper.SetDefault("ContentDir", "content")
	viper.SetDefault("language", "go")
}

func main() {
	viper.SetConfigName("app")
	viper.SetConfigType("yml")
	viper.AddConfigPath("E:\\Projects\\project-github\\little-go\\practices\\viper\\main")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	spring := viper.GetStringMap("spring")
	fmt.Println(spring["application"].(map[string]interface{})["name"])
}
