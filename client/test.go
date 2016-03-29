package main

import (
	"github.com/hashicorp/consul/api"
	"fmt"
)

func main(){
	// Get a new client
	client, err := api.NewClient(&api.Config{Address:"192.168.1.245:8500"})
	if err != nil {
		panic(err)
	}

	// Get a handle to the KV API
	kv := client.KV()

	// PUT a new KV pair
	p := &api.KVPair{Key: "foo", Value: []byte("test")}
	_, err = kv.Put(p, nil)
	if err != nil {
		panic(err)
	}

	// Lookup the pair
	pair, _, err := kv.Get("foo", nil)
	if err != nil {
		panic(err)
	}
	fmt.Printf("KV: %v", pair)
}