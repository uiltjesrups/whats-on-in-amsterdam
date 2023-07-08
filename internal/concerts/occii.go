package concerts

import (
	"log"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type OCCII Venue

var Occii = OCCII{
	Name: "OCCI",
	Url:  "https://occii.org/"}

func (occii OCCII) GetConcerts() []Concert {
	resp, err := http.Get(occii.Url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	getTitle := func(s *goquery.Selection) string {
		title := s.Find(".occii-event-link").Text()
		if title == "" {
			log.Fatalf("GetConcerts: title not found %v", occii)
		}
		return title
	}

	getDate := func(s *goquery.Selection) time.Time {
		dateTimeStr := s.Find(".occii-event-times").Text()
		layout := "Monday, January 2\x0aDoors open: 15:04"
		date, err := time.Parse(layout, dateTimeStr)
		if err != nil {
			log.Fatal(err)
		}
		date = time.Date(2023,
			date.Month(),
			date.Day(),
			date.Hour(),
			date.Minute(),
			date.Second(),
			date.Nanosecond(),
			date.Location())
		return date
	}

	agendaItems := doc.Find(".occii-event-display")
	concerts := []Concert{}
	agendaItems.Each(func(i int, s *goquery.Selection) {
		title := getTitle(s)
		date := getDate(s)
		concerts = append(concerts, Concert{Description: title, Date: date,
			Venue: Venue(occii)})
	})
	return concerts
}
