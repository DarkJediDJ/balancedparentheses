package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
}

func generate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	length, ok := vars["length"]
	if !ok {
		log.Fatal("Parameter not found")
	}
	len, err := strconv.Atoi(length)
	if err != nil {
		log.Fatal(err)
	}
	line := Parentheses{
		RandParentheses(len),
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

func main() {
	a := App{}
	a.New()

	a.Run(":8082")
}
