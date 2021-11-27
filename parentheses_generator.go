package main

import (
	"log"
	"math/rand"
	"strconv"
)

//RandParentheses  generates a random sequence of parentheses of the length
func RandParentheses(len string) string {
	num, err := strconv.Atoi(len)
	if err != nil {
		log.Fatal(err)
	}
	var s string
	for i := 0; i < num; i++ {
		s += M[strconv.Itoa(rand.Intn(6)+1)]
	}
	return s
}
