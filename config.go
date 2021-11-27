package main

var M = map[string]string{
	"1": "(",
	"2": "[",
	"3": "{",
	"4": ")",
	"5": "]",
	"6": "}",
}

type Parentheses struct {
	Line string `json:"parentheses"`
}
