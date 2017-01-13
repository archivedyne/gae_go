package main

import (
	"net/http"
)

func init() {
	http.HandleFunc("/tweet", tweet)
}

func tweet(w http.ResponseWriter, r *http.Request) {
	const (
		ckey = "Consumer Key"
		csec = "Consumer Secret"
		atok = "Access Token"
		asec = "Access Token Secret"
	)

	if r.Header.Get("X-Appengine-Cron")[0] == 0 {
		return
	}

	anaconda.SetConsumerKey(ckey)
	anaconda.SetConsumerSecret(csec)
	api := anaconda.NewTwitterApi(atok, asec)
	text := "Hello, world!"
	api.PostTweet(text, nil)
	return
}
