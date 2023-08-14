package concerts

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type Sexyland Venue

func (sexyLand Sexyland) GetConcerts() []Concert {
	resp, err := http.Get(sexyLand.Url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal("GetConcerts", err)
	}

	agendaItems := doc.Find("div.event")
	concerts := []Concert{}
	agendaItems.Each(func(i int, s *goquery.Selection) {
		processAgendaItem(sexyLand, s, &concerts)
	})
	return concerts
}

func getTitle(s *goquery.Selection) (string, error) {
	title := s.Find("div.event-wrap > h3").Text()
	if title == "" {
		html, err := s.Html()
		if err != nil {
			return "",
				fmt.Errorf("getTitle: title not found, selection: %s", err)

		}
		return "",
			fmt.Errorf("getTitle: title not found, selection: %s", html)
	}
	return title, nil
}

func getTime(s *goquery.Selection) (time.Time, error) {
	timeStr := s.Find("div.event-wrap > p").First().Text()
	timeParts := strings.Split(timeStr, " - ")

	if len(timeParts) != 2 {
		html, err := s.Html()
		if err != nil {
			return time.Time{},
				fmt.Errorf("getTime: timeParts not found, selection: %s", err)
		}
		return time.Time{},
			fmt.Errorf("getTime: timeParts not found, selection: %s", html)
	}

	// Parse the start time only.
	date, err := time.Parse("15:04", timeParts[0])
	if err != nil {
		return time.Time{}, err
	}

	return date, nil
}

func extractDate(s *goquery.Selection) (time.Time, error) {
	// Check if the input string has the required format
	d := s.AttrOr("data-date", "")
	if len(d) < 20 {
		return time.Time{}, fmt.Errorf("Invalid date format: %s ", d)
	}

	// Extract the date part
	dateString := d[10:20]

	layout := "2006/01/02"

	date, err := time.Parse(layout, dateString)
	if err != nil {
		return time.Time{}, err
	}

	return date, nil
}

func processAgendaItem(sexyLand Sexyland, s *goquery.Selection, concerts *[]Concert) {
	date, err := extractDate(s)

	if err != nil {
		log.Printf("processAgendaItem: extractdate: %s, error: %s", date, err)
		return
	}
	title, err := getTitle(s)
	if err != nil {
		log.Printf("processAgendaItem: title: %s, error: %s", title, err)
		return
	}

	t, err := getTime(s)
	if err != nil {
		log.Printf("processAgendaItem: title: %s, error: %s", title, err)
		return
	}

	*concerts = append(*concerts, Concert{
		Description: title,
		Date:        date.Add(t.Sub(time.Date(0, 1, 1, 0, 0, 0, 0, time.UTC))),
		Venue:       Venue(sexyLand),
	})
}
