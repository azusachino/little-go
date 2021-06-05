package main

import (
	"context"
	"fmt"
	"log"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

var cli *clientv3.Client

func init() {
	cli, _ = clientv3.New(clientv3.Config{
		Endpoints:   []string{"104.224.188.126:2379", "64.64.224.190:2379"},
		DialTimeout: 5 * time.Second,
	})
}

func main() {
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*5)
	defer cancel()

	cli.Put(ctx, "/mysql/host", "sh-cynosdbmysql-grp-gdj3i6os.sql.tencentcdb.com")
	cli.Put(ctx, "/mysql/port", "20958")
	cli.Put(ctx, "/mysql/username", "root")
	cli.Put(ctx, "/mysql/password", "Azusa1111/")

	res, _ := cli.Get(ctx, "/mysql/host")
	fmt.Println(res)
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
