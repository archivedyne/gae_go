package main

import (
	"fmt"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
	"net/http"
)

const Table = "Person"

const keyName = "KEY_NAME"

type Person struct {
	Name   string
	Gender string
}

func init() {
	http.HandleFunc("/d/set", func(w http.ResponseWriter, r *http.Request) {
		ctx := appengine.NewContext(r)
		p := &Person{
			Name:   "hoge",
			Gender: "male",
		}
		key := datastore.NewIncompleteKey(ctx, Table, nil)
		if _, err := datastore.Put(ctx, key, p); err != nil {
			log.Errorf(ctx, "Message")
		}
		w.Write([]byte("ok!"))
	})

	http.HandleFunc("/d/set2", func(w http.ResponseWriter, r *http.Request) {
		ctx := appengine.NewContext(r)
		p := &Person{
			Name:   "hoge",
			Gender: "male",
		}
		key := datastore.NewKey(ctx, Table, keyName, 0, nil)
		if _, err := datastore.Put(ctx, key, p); err != nil {
			log.Errorf(ctx, err.Error())
		}
		w.Write([]byte("ok!"))
	})

	http.HandleFunc("/d/get", func(w http.ResponseWriter, r *http.Request) {
		ctx := appengine.NewContext(r)
		var ps []Person
		q := datastore.NewQuery(Table).Filter("Gender=", "male")
		if _, err := q.GetAll(ctx, &ps); err != nil {
			log.Errorf(ctx, err.Error())
		}
		for _, p := range ps {
			fmt.Fprintf(w, "%s %s \n", p.Name, p.Gender)
		}
	})

	http.HandleFunc("/d/get2", func(w http.ResponseWriter, r *http.Request) {
		ctx := appengine.NewContext(r)
		var p Person
		key := datastore.NewKey(ctx, Table, keyName, 0, nil)
		if err := datastore.Get(ctx, key, &p); err != nil {
			log.Errorf(ctx, err.Error())
		}
		fmt.Fprintf(w, "%s %s \n", p.Name, p.Gender)
	})

}
