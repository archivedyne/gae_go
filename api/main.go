package main

import (
	"fmt"
	"net/http"
)

const PROJECT_ID = "onyx-day-154915"

func init() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, html)
	})
}

const html = `
<html>
<a href="/">top</a>
<br>
<strong>datastore</strong>
<br>
<a href="/d/set">set</a>
<a href="/d/get">get</a>
<a href="/d/set2">set2</a>
<a href="/d/get2">get2</a>
<br>
<strong>memocached</strong>
<br>
<a href="/m/set">set</a>
<a href="/m/get">get</a>
<a href="/m/del">del</a>
<br>
<strong>user</strong>
<br>
<a href="/u/login">login</a>
<a href="/u/oauth">oauth</a>
<br>
<strong>storage</strong>
<br>
<a href="/s/set">set</a>
<a href="/s/get">get</a>
</html>
`
