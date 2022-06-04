package main

import (
	"citizenship/crawler"
	"citizenship/identity"
	"citizenship/socket"
	"time"
)

var (
	Token string
	Cookies string
)

func main() {
	go GetIDs()
	socket.Main()
}

func SetCookies() {
	for {
		Token, Cookies = crawler.GetToken()
		time.Sleep(24*time.Hour)
	}
}


func GetIDs() {
	identity.Main()
	go SetCookies()
	time.Sleep(5*time.Second)
	Names := crawler.GetAllNames()

	for identity.SecretData.CurrentIDNumber < 100000000 {
		if identity.SecretData.CurrentIDNumber < 9999999 {
			identity.SecretData.CurrentIDNumber = 9999999
			continue
		}
		for _, name := range Names {
			ID := identity.GetID(name, identity.SecretData.CurrentIDNumber, Token, Cookies)
			if ID.IDNumber != "" {
				ID.Save()
				break
			}
		}
		identity.SecretData.Save()
		identity.SecretData.CurrentIDNumber++
	}
}