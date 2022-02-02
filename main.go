package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

var hangman Hangman

func main() {
	fmt.Println("start server")
	if len(os.Args) == 1 {
		LoadingWord("words.txt")
	} else {
		LoadingWord(os.Args[1])
	}

	HangmanInit()
	templ := template.Must(template.ParseFiles("./index.html"))
	templ2 := template.Must(template.ParseFiles("./hangman.html"))

	filescss := http.FileServer(http.Dir("./css"))
	http.Handle("/css/", http.StripPrefix("/css/", filescss))
	filesimage := http.FileServer(http.Dir("./image"))
	http.Handle("/image/", http.StripPrefix("/image/", filesimage))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		templ.Execute(w, nil)
	})

	http.HandleFunc("/hangman", func(write http.ResponseWriter, respons *http.Request) {
		hangman.InputUser = respons.FormValue("user_input")
		TryInput(hangman.InputUser)
		if respons.Method != http.MethodPost {
			HangmanInit()
		}
		templ2.Execute(write, hangman)
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
	fmt.Print("close server")
}
