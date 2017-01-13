package main

import (
	"fmt"
	"google.golang.org/appengine"
	"google.golang.org/appengine/memcache"
	"net/http"
	"time"
)

func init() {
	http.HandleFunc("/m/set", func(w http.ResponseWriter, r *http.Request) {
		ctx := appengine.NewContext(r)
		item := &memcache.Item{
			Key:        "foo",
			Value:      []byte("bar"),
			Expiration: time.Second * 6,
		}
		if err := memcache.Set(ctx, item); err != nil {
			fmt.Fprintf(w, "%s", err.Error())
		}
		w.Write([]byte("OK!"))
	})

	http.HandleFunc("/m/get", func(w http.ResponseWriter, r *http.Request) {
		ctx := appengine.NewContext(r)
		key := "foo"
		item, err := memcache.Get(ctx, key)
		if err != nil {
			fmt.Fprintf(w, "%s \n", err.Error())
		} else {
			fmt.Fprintf(w, "KEY: %s, VALUE: %s \n", item.Key, item.Value)
		}
		w.Write([]byte("OK!"))
	})

	http.HandleFunc("/m/del", func(w http.ResponseWriter, r *http.Request) {
		ctx := appengine.NewContext(r)
		key := "foo"
		err := memcache.Delete(ctx, key)
		if err != nil {
			fmt.Fprintf(w, "%s \n", err.Error())
		}
		w.Write([]byte("OK!"))
	})
}
