package crawler

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)


func GetToken() (string, string) {
	var Cookies string
	var Token string
	collector := colly.NewCollector()
	collector.UserAgent = "Mozilla/5.0 (X11; Linux x86_64; rv:100.0) Gecko/20100101 Firefox/100.0"

	collector.OnHTML("head", func(element *colly.HTMLElement) {
		element.ForEach("meta", func(index int, element *colly.HTMLElement) {
			if element.Attr("name") == "csrf-token" {
				Token = element.Attr("content")
				cookies := collector.Cookies("https://accounts.ecitizen.go.ke/register/citizen")
				for _, cookie := range cookies {
					Cookies = Cookies + fmt.Sprintf("%s=%s;", cookie.Name, cookie.Value)
				}
				Cookies = strings.TrimSuffix(Cookies, ";")
				return
			}
		})
	})
	collector.Visit("https://accounts.ecitizen.go.ke/register/citizen")
	return Token, Cookies
}