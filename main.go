package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

func handler(w http.ResponseWriter, r *http.Request) {
	client := connectEtcd()
	defer client.Close()
	setKeyValue(client, "test-key", "test-value")
	value := getKeyValue(client, "test-key")
	fmt.Println("Value of test-key:", value)
	fmt.Fprintf(w, "Hello, World!")
}

func connectEtcd() *clientv3.Client {
	cfg := clientv3.Config{
		Endpoints:   []string{"etcd1:12379", "etcd2:22379"},
		DialTimeout: 5 * time.Second,
	}
	client, err := clientv3.New(cfg)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func setKeyValue(client *clientv3.Client, key, value string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err := client.Put(ctx, key, value)
	if err != nil {
		log.Fatal(err)
	}
}

func getKeyValue(client *clientv3.Client, key string) string {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	resp, err := client.Get(ctx, key)
	if err != nil {
		log.Fatal(err)
	}
	if resp.Count == 0 {
		return ""
	}
	return string(resp.Kvs[0].Value)
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
