package concerts

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type Zaal100 Venue

func (zaal100 Zaal100) GetConcerts() []Concert {
	resp, err := http.Get(zaal100.Url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	getTitle := func(s *goquery.Selection) string {
		title := s.Find(".titel").Text()
		if title == "" {
			log.Fatalf("GetConcerts: title not found %v", zaal100)
		}
		return title
	}

	parseDate := func(day, month, year string) (time.Time, error) {
		if year == "" {
			year = strconv.Itoa(time.Now().Year())
		}
		layout := "2006-01-02"
		dateStr := fmt.Sprintf("%s-%s-%s", year, month, day)
		t, err := time.Parse(layout, dateStr)
		if err != nil {
			return time.Time{}, err
		}
		return t, nil
	}

	getDate := func(s *goquery.Selection) (time.Time, error) {
		month := s.Find("span.maand").Text()
		day := s.Find("span.datum").Text()
		year := strconv.Itoa(time.Now().Year())
		t, err := parseDate(day, month, year)
		if err != nil {
			return time.Time{}, err
		}
		return t, nil
	}

	// Use the CSS selector to select the element
	agendaItems := doc.Find(".agenda-item")

	concerts := []Concert{}
	// Iterate over the selected elements and print their text
	agendaItems.Each(func(i int, s *goquery.Selection) {
		title := getTitle(s)
		date, err := getDate(s)
		if err != nil {
			log.Printf("getConcerts: title: %s, error: %s",
				title,
				err)
		} else {
			concerts = append(concerts,
				Concert{Description: title,
					Date:  date,
					Venue: Venue(zaal100)})
		}
	})
	return concerts
}
