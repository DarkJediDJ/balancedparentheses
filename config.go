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

type Percents struct {
	Two   int `json:"Percent for 2 is"`
	Four  int `json:"Percent for 4 is"`
	Eight int `json:"Percent for 8 is"`
}

var A App
