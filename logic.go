package main

import (
	"io"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
)

func RandParentheses(len string) string {
	num, err := strconv.Atoi(len)
	if err != nil {
		log.Fatal(err)
	}
	s := ""
	for i := 0; i < num; i++ {
		s += M[strconv.Itoa(rand.Intn(6)+1)]
	}
	return s
}

func GetValue(n int) string {
	resp, err := http.Get("http://localhost:8080/generate/" + strconv.Itoa(n))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	line := string(body)
	commas := strings.LastIndex(line, ":")
	return line[commas+2 : len(line)-2]
}

func GetStatus(n int) int {
	resp, err := http.Get("http://localhost:8080/generate/" + strconv.Itoa(n))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	return resp.StatusCode
}
