package concerts

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type OCCII Venue

var Occii = OCCII{
	Name: "OCCI",
	Url:  "https://occii.org/"}

func (occii OCCII) GetConcerts() ([]Concert, error) {
	resp, err := http.Get(occii.Url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Parse the HTML document
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	getTitle := func(s *goquery.Selection) (string, error) {
		title := s.Find(".occii-event-link").Text()
		if title == "" {
			return "", errors.New("title not found")
		}
		return title, nil
	}

	getDate := func(s *goquery.Selection) (time.Time, error) {
		dateTimeStr := s.Find(".occii-event-times").Text()
		fmt.Println("dateTimeStr: " + dateTimeStr)
		//dateStr := strings.Split(dateTimeStr, newLine)[0]

		layout := "Monday, January 2\x0aDoors open: 15:04"
		date, err := time.Parse(layout, dateTimeStr)

		date = time.Date(2023, date.Month(), date.Day(), date.Hour(), date.Minute(), date.Second(), date.Nanosecond(), date.Location())

		if err != nil {
			log.Fatal(err)
		}
		return date, nil
	}

	// Use the CSS selector to select the element
	agendaItems := doc.Find(".occii-event-display")

	concerts := []Concert{}
	// Iterate over the selected elements and print their text
	agendaItems.Each(func(i int, s *goquery.Selection) {
		title, err := getTitle(s)
		if err != nil {
			log.Printf("error parsing title: %v", err)
			return
		}
		date, err := getDate(s)
		if err != nil {
			log.Printf("error parsing date: %v", err)
			return
		}
		concerts = append(concerts, Concert{Description: title, Date: date,
			Venue: Venue(occii)})
	})
	return concerts, nil
}
