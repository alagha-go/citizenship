package identity

import (
	"encoding/json"
	"io/ioutil"
)


var (
	SecretData Secret
)

func LoadSecret() {
	data, _ := ioutil.ReadFile("./secret.json")
	json.Unmarshal(data, &SecretData)
}