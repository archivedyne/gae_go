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

	// カーソル
	http.HandleFunc("/d/get3", func(w http.ResponseWriter, r *http.Request) {
		ctx := appengine.NewContext(r)
		query := datastore.NewQuery(Table)
		cursor, err := datastore.DecodeCursor(string(cursor.Value))
		if err == nil {
			query = query.Start(cursor)
		}
		var p Person
		t := query.Run(c)
		for {
			var p Person
			_, err := t.Next(&p) // 取得する結果のKeyを返す
			if err == datastore.Done {
				break
			}
			if err != nil {
				log.Errorf("fetching next Person: %v", err)
				break
			}
			fmt.Fprintf(w, "%v\n", p)
		}
		cursor, err := t.Cursor() // イテレーターの現在位置を表すカーソル
		if err != nil {
			fmt.Fprintf(w, "%s", err.Error())
			return
		}
		w.Write([]byte([]byte(cursor.String())))
	})
}
