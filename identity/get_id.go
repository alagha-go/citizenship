package identity

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)


func GetID(name string, ID int, Token, Cookies string) Identity {
	var Identity Identity
	client := &http.Client{}
	var data = strings.NewReader(fmt.Sprintf(`id_number=%d&id_type=citizen&first_name=%s&_csrf_token=%s`, ID, name, Token))
	req, err := http.NewRequest("POST", "https://accounts.ecitizen.go.ke/lookup", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64; rv:100.0) Gecko/20100101 Firefox/100.0")
	req.Header.Set("Accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Set("Accept-Language", "en-US,en;q=0.5")
	req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("Origin", "https://accounts.ecitizen.go.ke")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Referer", "https://accounts.ecitizen.go.ke/register/citizen")
	req.Header.Set("Cookie", Cookies)
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(body, &Identity)
	return Identity
}