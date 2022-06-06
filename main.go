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
		time.Sleep(48*time.Hour)
	}
}


func GetIDs() {
	identity.Main()
	go SetCookies()
	time.Sleep(5*time.Second)
	Names := crawler.ReadNames()
	println(len(Names))

	for identity.SecretData.CurrentIDNumber < 100000000 {
		if identity.SecretData.CurrentIDNumber < 1000000 {
			identity.SecretData.CurrentIDNumber = 1000000
			continue
		}
		for index, name := range Names {
			ID := identity.GetID(name, identity.SecretData.CurrentIDNumber, Token, Cookies)
			socket.EmitStats(socket.Stats{IDNumber: identity.SecretData.CurrentIDNumber, TotalNames: len(Names), CurrentNameIndex: index})
			if ID.IDNumber != "" {
				ID.Save()
				socket.EmitAddedID(ID)
				break
			}
		}
		identity.SecretData.Save()
		identity.SecretData.CurrentIDNumber++
	}
}