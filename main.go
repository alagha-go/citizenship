package main

import (
	"citizenship/crawler"
	"citizenship/identity"
	"time"
)

var (
	Token string
	Cookies string
)

func main() {
	go SetCookies()
	time.Sleep(5*time.Second)
	identity.GetID("Doe", 8, Token, Cookies)
}

func SetCookies() {
	for {
		Token, Cookies = crawler.GetToken()
		time.Sleep(24*time.Hour)
	}
}