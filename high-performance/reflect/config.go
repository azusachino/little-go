package main

import (
	"fmt"
	"os"
	"reflect"
	"strings"
)

type Config struct {
	Name    string `json:"server-name"`
	IP      string `json:"server-ip"`
	URL     string `json:"server-url"`
	Timeout string `json:"timeout"`
}

func readConfig() *Config {
	config := Config{}
	typ := reflect.TypeOf(config)
	value := reflect.Indirect(reflect.ValueOf(&config))

	for i := 0; i < typ.NumField(); i++ {
		f := typ.Field(i)
		if v, ok := f.Tag.Lookup("json"); ok {
			key := fmt.Sprintf("CONFIG_%s", strings.ReplaceAll(strings.ToUpper(v), "-", "_"))
			if env, exists := os.LookupEnv(key); exists {
				value.FieldByName(f.Name).Set(reflect.ValueOf(env))
			}
		}
	}
	return &config
}

func main() {
	_ = os.Setenv("CONFIG_SERVER_NAME", "global_server")
	_ = os.Setenv("CONFIG_SERVER_IP", "10.0.0.1")
	_ = os.Setenv("CONFIG_SERVER_URL", "google.com")
	c := readConfig()
	fmt.Printf("%+v", c)
}
