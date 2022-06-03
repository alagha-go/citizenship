package crawler

import "github.com/gocolly/colly"


func GetNames() []string {
	var Names []string
	collector := colly.NewCollector()

	collector.OnHTML("tbody", func(element *colly.HTMLElement) {
		element.ForEach("td.sur", func (index int, element *colly.HTMLElement) {
			Names = append(Names, element.Text)
		})
	})

	collector.Visit("https://forebears.io/kenya/forenames")

	println(len(Names))
	return Names
}