package main

import (
	"context"
	"fmt"
	"github.com/azusachino/little-go/learn-project/ants"

	"github.com/caarlos0/env/v6"
	"time"
)

type config struct {
	Home         string        `env:"HOME"`
	Port         int           `env:"PORT" envDefault:"3000"`
	Password     string        `env:"PASSWORD,unset"`
	IsProduction bool          `env:"PRODUCTION"`
	Hosts        []string      `env:"HOSTS" envSeparator:":"`
	Duration     time.Duration `env:"DURATION"`
	TempFolder   string        `env:"TEMP_FOLDER" envDefault:"${HOME}/tmp" envExpand:"true"`
}

func main() {

	cfg := config{}
	if err := env.Parse(&cfg); err != nil {
		fmt.Printf("%+v\n", err)
	}

	fmt.Printf("%+v\n", cfg)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	for i := 0; i < 10; i++ {
		select {
		case <-ctx.Done():
			// try goto for break loop
			goto label
		default:
			_ = ants.Submit(func() {
				fmt.Println(time.Now().Format("2006-01-02 15:04:05.000"))
			})
		}
	}
label:

	time.Sleep(time.Second * 3)
	fmt.Println(ants.Running())
	ants.Release()
	fmt.Println(ants.Running())

	fmt.Println("Hello Little GO!")
}
