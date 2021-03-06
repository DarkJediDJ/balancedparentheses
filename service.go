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

func percents(w http.ResponseWriter, r *http.Request) {
	m := CalculateBalancedPercentage()

	result := Percents{
		m[2],
		m[4],
		m[8],
	}

	res, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")

	_, err = w.Write(res)
	if err != nil {
		log.Fatal(err)
	}
}

// New creates router with handler
func (a *App) New() {
	myRouter := mux.NewRouter().StrictSlash(false)
	myRouter.HandleFunc("/generate/{length:[0-9]+}", generate)
	myRouter.HandleFunc("/generate/result", percents)
	a.Router = myRouter
}

// Run starts server
func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func main() {
	a := App{}
	a.New()

	a.Run(":8082")
}
