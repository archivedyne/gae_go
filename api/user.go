package main

import (
	"fmt"
	"google.golang.org/appengine"
	"google.golang.org/appengine/user"
	"net/http"
)

func init() {
	http.HandleFunc("/u/login", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "text/html; charset=utf-8")
		ctx := appengine.NewContext(r)
		u := user.Current(ctx)
		if u == nil {
			url, _ := user.LoginURL(ctx, "/")
			fmt.Fprintf(w, `<a href="%s">Sign in or register</a>`, url)
			return
		}
		url, _ := user.LogoutURL(ctx, "/")
		fmt.Fprintf(w, `Welcome, %s! (<a href="%s">sign out</a>)`, u, url)
	})

	http.HandleFunc("/u/oauth", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "text/html; charset=utf-8")
		ctx := appengine.NewContext(r)
		u, err := user.CurrentOAuth(ctx)
		if err != nil {
			http.Error(w, "OAuth Authorization header required", http.StatusUnauthorized)
			return
		}
		if !u.Admin {
			http.Error(w, "Admin login only", http.StatusUnauthorized)
			return
		}
		fmt.Fprintf(w, `Welcome, admin user %s!`, u)
	})
}
