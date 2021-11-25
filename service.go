package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	len := vars["length"]
	line := Parentheses{
		RandParentheses(len),
	}
	res, err := json.Marshal(line)
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}
func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/generate/{length}", returnSingleArticle)
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func main() {
	handleRequests()
}
