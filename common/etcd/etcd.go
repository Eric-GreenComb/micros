package etcd

import (
	"flag"
	"log"
	"time"

	"github.com/coreos/etcd/client"
	"golang.org/x/net/context"
)

var keysAPI client.KeysAPI

func init() {
	fs := flag.NewFlagSet("", flag.ExitOnError)
	var (
		etcdAddr = fs.String("etcd.addr", "http://127.0.0.1:2379", "Address for etcd server")
	)

	cfg := client.Config{
		Endpoints: []string{*etcdAddr},
		Transport: client.DefaultTransport,
		// set timeout per request to fail fast when the target endpoint is unavailable
		HeaderTimeoutPerRequest: time.Second,
	}
	c, err := client.New(cfg)
	if err != nil {
		log.Fatal(err)
	}
	keysAPI = client.NewKeysAPI(c)
}

func Set(key, value string) (*client.Response, error) {
	return keysAPI.Set(context.Background(), key, value, nil)
}

func Get(key string) (*client.Response, error) {
	return keysAPI.Get(context.Background(), key, nil)
}

func GetValue(key string) (string, error) {
	resp, err := keysAPI.Get(context.Background(), key, nil)
	return resp.Node.Value, err
}

func GetString(key string) string {
	resp, err := GetValue(key)
	if err != nil {
		return ""
	}
	return resp
}

// key = /banerwai/mongo return multi node
// /banerwai/mongo/conn       localhost:27017
// banerwai/mongo/database    banerwai
func GetService(key string) (*client.Response, error) {
	return keysAPI.Get(context.Background(), key, nil)
}
