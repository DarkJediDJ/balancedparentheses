package main

import (
	"math/rand"
	"strconv"
)

//RandParentheses  generates a random sequence of parentheses of the length
func RandParentheses(len int) string {
	var s string
	for i := 0; i < len; i++ {
		s += M[strconv.Itoa(rand.Intn(6)+1)]
	}
	return s
}
