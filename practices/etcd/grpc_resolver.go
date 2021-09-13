package etcd

import (
	"context"
	"encoding/json"
	"fmt"
	"go.etcd.io/etcd/api/v3/mvccpb"
	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc/resolver"
	"log"
)

const schema = "etcd"

// Resolver implements grpc.resolve.Builder
type Resolver struct {
	endpoints []string
	service   string
	cli       *clientv3.Client
	cc        resolver.ClientConn
}

func NewResolver(endpoints []string, service string) resolver.Builder {
	return &Resolver{endpoints: endpoints, service: service}
}

// Scheme return etcd schema
func (r *Resolver) Scheme() string {
	// 最好用这种，因为grpc resolver.Register(r)在注册时，会取scheme，如果一个系统有多个grpc发现，就会覆盖之前注册的
	return schema + "_" + r.service
}

func (r *Resolver) ResolveNow(rn resolver.ResolveNowOptions) {
}

func (r *Resolver) Close() {
}

// Build to resolver.Resolver
// 实现grpc.resolve.Builder接口的方法
func (r *Resolver) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	var err error

	r.cli, err = clientv3.New(clientv3.Config{
		Endpoints: r.endpoints,
	})
	if err != nil {
		return nil, fmt.Errorf("grpclb: create clientv3 client failed: %v", err)
	}

	r.cc = cc

	// go r.watch(fmt.Sprintf("/%s/%s/", schema, r.service))
	go r.watch(fmt.Sprintf(r.service))

	return r, nil
}

func (r *Resolver) watch(prefix string) {
	addrDict := make(map[string]resolver.Address)
	update := func() {
		addrList := make([]resolver.Address, 0, len(addrDict))
		for _, v := range addrDict {
			addrList = append(addrList, v)
		}
		r.cc.NewAddress(addrList)
	}

	res, err := r.cli.Get(context.Background(), prefix, clientv3.WithPrefix())
	if err == nil {
		for i, kv := range res.Kvs {
			info := &ServiceInfo{}
			err := json.Unmarshal(kv.Value, info)
			if err != nil {
			}
			addrDict[string(res.Kvs[i].Value)] = resolver.Address{Addr: info.IP}
		}
	}
	update()

	rch := r.cli.Watch(context.Background(), prefix, clientv3.WithPrefix(), clientv3.WithPrevKV())
	for n := range rch {
		for _, ev := range n.Events {
			switch ev.Type {
			case mvccpb.PUT:
				info := &ServiceInfo{}
				err := json.Unmarshal([]byte(ev.Kv.Value), info)
				if err != nil {
					log.Println(err)
				} else {
					addrDict[string(ev.Kv.Key)] = resolver.Address{Addr: info.IP}
				}
			case mvccpb.DELETE:
				delete(addrDict, string(ev.PrevKv.Key))
			}
		}
		update()
	}
}
