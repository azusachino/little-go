package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/spf13/viper"
	clientv3 "go.etcd.io/etcd/client/v3"
)

var cli *clientv3.Client

func init() {
	// 读取配置文件
	viper.AddConfigPath(".")
	viper.SetConfigName("app")
	viper.SetConfigType("toml")
	_ = viper.ReadInConfig()

	etcdCluster := viper.GetStringSlice("etcd.cluster")
	cli, _ = clientv3.New(clientv3.Config{
		Endpoints:   etcdCluster,
		DialTimeout: 5 * time.Second,
	})
}

func main() {
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*5)
	defer cancel()

	res, _ := cli.Get(ctx, "/mysql/host")
	fmt.Println(res)
	sample()
}

func sample() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	// resp, err := cli.Put(ctx, "sample_key", "sample_value")
	resp, err := cli.Get(ctx, "sample_key")
	cancel()
	if err != nil {
		switch err {
		case context.Canceled:
			log.Fatalf("ctx is canceled by another routine: %v", err)
		case context.DeadlineExceeded:
			log.Fatalf("ctx is attached with a deadline is exceeded: %v", err)
		// case rpctypes.ErrEmptyKey:
		// 	log.Fatalf("client-side error: %v", err)
		default:
			log.Fatalf("bad cluster endpoints, which are not etcd servers: %v", err)
		}
	}
	// use the response
	log.Default().Println(resp)
}
