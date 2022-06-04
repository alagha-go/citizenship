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
	GetIDs()
}

func SetCookies() {
	for {
		Token, Cookies = crawler.GetToken()
		time.Sleep(24*time.Hour)
	}
}


func GetIDs() {
	identity.Main()
	crawler.GetAllNames()
	go SetCookies()
	time.Sleep(5*time.Second)
	Names := crawler.GetAllNames()

	for identity.SecretData.CurrentIDNumber < 100000000 {
		for _, name := range Names {
			ID := identity.GetID(name, identity.SecretData.CurrentIDNumber, Token, Cookies)
			if ID.IDNumber != "" {
				ID.Save()
			}
		}
	}
}