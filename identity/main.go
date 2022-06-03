package identity

import (
	"encoding/json"
	"io/ioutil"
	"log"
)


var (
	SecretData Secret
)

func LoadSecret() {
	data, _ := ioutil.ReadFile("./secret.json")
	json.Unmarshal(data, &SecretData)
}

func (Secret *Secret) Save() {
	data, err := json.Marshal(Secret)
	if err != nil {
		log.Panic(err)
	}
	ioutil.WriteFile("./secret.json", data, 0755)
}