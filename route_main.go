package main

import (
	"gowebprog/ch02/chitchat/data"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	threads, err := data.Threads()
	if err != nil {
		error_message(w, r, "Cannot get threads")
	} else {
		_, err := session(w, r)
		if err != nil {
			generateHTML(w, threads, "layout", "public.navbar", "index")
		} else {
			generateHTML(w, threads, "layout", "public.navbar", "index")
		}
	}
}
