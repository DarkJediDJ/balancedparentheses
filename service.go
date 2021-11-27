package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
}

// HERE use NEW LINES, plz. The code becomes unreadable! (DONE)
func generate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	length, ok := vars["length"] // HERE - you should check if element exists here (DONE)
	if !ok {
		log.Fatal("Parameter not found")
	}

	line := Parentheses{
		RandParentheses(length), // HERE - make Atoi here, utility function should not do it, it just returns string
	}

	res, err := json.Marshal(line)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")

	_, err = w.Write(res)
	if err != nil {
		log.Fatal(err)
	}
}

// Init the API
func (a *App) New() {
	myRouter := mux.NewRouter().StrictSlash(false)
	myRouter.HandleFunc("/generate/{length:[0-9]+}", generate)

	a.Router = myRouter
}

// And then run it
func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

// HERE - worked, but you were calling generate function, to the router itself, so the URL parameters was not placed
// to request context. TODO - check what mux.Vars(r) does!
//func handleRequests() {
//	myRouter := mux.NewRouter().StrictSlash(true)
//	myRouter.HandleFunc("/generate/{length}", generate)
//	log.Fatal(http.ListenAndServe(":8080", myRouter))
//}

func main() {
	a := App{}
	a.New()

	a.Run(":8082")
}
