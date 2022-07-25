package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/gocolly/colly"
)

// Course stores information about a coursera course
type Course struct {
	Title    string
	URL      string
	Rating   string
	Enrolled string
}

var Limit_Page = 100

func main() {
	// Instantiate default collector
	c := colly.NewCollector(
		// Visit only domains: coursera.org, www.coursera.org
		colly.AllowedDomains("coursera.org", "www.coursera.org"),

		// Cache responses to prevent multiple download of pages
		// even if the collector is restarted
		colly.CacheDir("./coursera_cache"),
	)

	// Create another collector to scrape course details
	detailCollector := c.Clone()

	courses := make([]Course, 0, 5)

	// On every a element which has href attribute call callback
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		// If attribute class is this long string return from callback
		// As this a is irrelevant
		if e.Attr("class") == "Button_1qxkboh-o_O-primary_cv02ee-o_O-md_28awn8-o_O-primaryLink_109aggg" {
			return
		}
		link := e.Attr("href")
		// If link start with browse or includes either signup or login return from callback
		if !strings.HasPrefix(link, "/browse") || strings.Contains(link, "=signup") || strings.Contains(link, "=login") {
			return
		}
		// start scaping the page under the link found
		e.Request.Visit(link)
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		log.Println("visiting", r.URL.String())
	})

	limitPages := 0
	// On every a HTML element which has name attribute call callback
	c.OnHTML(`a[data-click-value]`, func(e *colly.HTMLElement) {
		log.Println("counter", limitPages)
		if limitPages > Limit_Page {
			return
		}
		// Activate detailCollector if the link contains "coursera.org/learn"
		courseURL := e.Request.AbsoluteURL(e.Attr("href"))
		if strings.Contains(courseURL, "coursera.org/learn") {
			detailCollector.Visit(courseURL)
		}
		limitPages++
	})
	// Extract details of the course
	detailCollector.OnHTML(`div[id=rendered-content]`, func(e *colly.HTMLElement) {
		log.Println("Course found", e.Request.URL)
		title := e.ChildText(".banner-title")
		if title == "" {
			log.Println("No title found", e.Request.URL)
		}
		course := Course{
			Title:    title,
			URL:      e.Request.URL.String(),
			Enrolled: e.ChildText(".rc-ProductMetrics > div > span > strong > span"),
			Rating:   e.ChildText(".ratings-count-expertise-style > span > span"),
		}
		courses = append(courses, course)
	})

	// Start scraping on http://coursera.com/browse
	c.Visit("https://coursera.org/browse")

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")

	rankingsJson, _ := json.Marshal(courses)
	err := ioutil.WriteFile("output.json", rankingsJson, 0644)
	if err != nil {
		return
	}
	// Dump json to the standard output
	enc.Encode(courses)
}
