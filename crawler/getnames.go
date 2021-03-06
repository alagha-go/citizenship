package crawler

import "github.com/gocolly/colly"

var (
	Names []string
)


func GetAllNames() []string {
	urls := GetUrls()
	for _, url := range urls {
		GetNames(url)
	}
	return Names
}


func GetNames(url string) {
	collector := colly.NewCollector()

	collector.OnHTML(".search-results", func(element *colly.HTMLElement) {
		element.ForEach(".row", func (_ int, element *colly.HTMLElement) {
			name := element.ChildText(".name")
			Names = append(Names, name)
		})
	})

	collector.Visit(url)
}


func GetUrls() []string {
	var Urls []string
	collector := colly.NewCollector()
	collector.OnHTML(".pagination", func (element *colly.HTMLElement) {
		element.ForEach("a", func(_ int, element *colly.HTMLElement) {
			url := "https://forebears.io/" + element.Attr("href")
			Urls = append(Urls, url)
		})
	})
	collector.Visit("https://forebears.io/forenames")
	return Urls
}