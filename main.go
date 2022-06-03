package main

import (
	"citizenship/crawler"
	"time"
)

var (
	Token string
	Cookies string
)

func main() {
	
}

func SetCookies() {
	for {
		Token, Cookies = crawler.GetToken()
		time.Sleep(24*time.Hour)
	}
}