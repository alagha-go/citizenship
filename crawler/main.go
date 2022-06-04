package crawler

import (
	"encoding/json"
	"io/ioutil"
)


func ReadNames() []string {
	var Names []string
	data, _ := ioutil.ReadFile("./DB/names.json")
	json.Unmarshal(data, &Names)
	return Names
}