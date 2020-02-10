package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/artyom/guesslanguage"
)

func handler(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open("page.html")
	if err != nil {
		fmt.Fprintf(w, "Pizdec, oshibka!!")
		return
	}
	if _, err := io.Copy(w, f); err != nil {
		fmt.Fprintf(w, "Pizdec, oshibka!!")
		return
	}
}

func guessHandler(w http.ResponseWriter, r *http.Request) {
	text := r.FormValue("text")
	lang, err := guesslanguage.Guess(text)
	if err != nil { 
		return 	
	}
	fmt.Fprintf(w,`<!doctype html>
	<style>
		.center {
			text-align: center;
		}
	</style>
	<title>GuessLang</title>
	<h1 class=center>%s</h1>`, lang)
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/guess", guessHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}