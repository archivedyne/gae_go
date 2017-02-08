package main

import (
	"bytes"

	"cloud.google.com/go/storage"

	"fmt"
	"google.golang.org/appengine"
	"google.golang.org/appengine/file"
	"io/ioutil"
	"net/http"
)

// Note: storageは本番環境のみ有効
func init() {
	http.HandleFunc("/s/set", func(w http.ResponseWriter, r *http.Request) {
		ctx := appengine.NewContext(r)

		fileName := "hoge.txt"
		contents := []byte("abcde\n")

		bucketName, err := file.DefaultBucketName(ctx)
		if err != nil {
			fmt.Fprintf(w, "%s", err.Error())
			return
		}

		client, err := storage.NewClient(ctx)
		if err != nil {
			fmt.Fprintf(w, "%s", err.Error())
			return
		}
		defer client.Close()

		wc := client.Bucket(bucketName).Object(fileName).NewWriter(ctx)
		wc.ContentType = "text/plain"
		wc.Metadata = map[string]string{
			"x-goog-meta-foo": "foo",
		}

		if _, err := wc.Write(contents); err != nil {
			fmt.Fprintf(w, "%s", err.Error())
			return
		}
		if err := wc.Close(); err != nil {
			fmt.Fprintf(w, "%s", err.Error())
			return
		}

		w.Write([]byte("ok"))
	})

	http.HandleFunc("/s/get", func(w http.ResponseWriter, r *http.Request) {
		ctx := appengine.NewContext(r)
		fileName := "hoge.txt"

		bucketName, err := file.DefaultBucketName(ctx)
		if err != nil {
			fmt.Fprintf(w, "%s", err.Error())
			return
		}

		client, err := storage.NewClient(ctx)
		if err != nil {
			fmt.Fprintf(w, "%s", err.Error())
			return
		}
		defer client.Close()

		rc, err := client.Bucket(bucketName).Object(fileName).NewReader(ctx)
		if err != nil {
			fmt.Fprintf(w, "%s", err.Error())
			return
		}
		defer rc.Close()

		//
		slurp, err := ioutil.ReadAll(rc)
		if err != nil {
			fmt.Fprintf(w, "%s", err.Error())
		}

		fmt.Fprintf(w, "%s\n", bytes.SplitN(slurp, []byte("\n"), 2)[0])
		if len(slurp) > 1024 {
			fmt.Fprintf(w, "...%s\n", slurp[len(slurp)-1024:])
		} else {
			fmt.Fprintf(w, "%s\n", slurp)
		}

		w.Write([]byte("ok"))
	})

}
