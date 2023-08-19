package concerts

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type Orgelpark Venue

func (venue Orgelpark) getTitle(s *goquery.Selection) (string, error) {
	title := s.Find("a section h1").Text()
	if title == "" {
		html, err := s.Html()
		if err != nil {
			return "",
				fmt.Errorf("getTitle: title not found, selection: %s", err)

		}
		return "",
			fmt.Errorf("getTitle: title not found, selection: %s", html)
	}

	return strings.TrimSpace(title), nil
}

func (venue Orgelpark) getDate(s *goquery.Selection) (time.Time, error) {

	dateTimeStr, exists := s.Attr("data-date")

	if exists {
		layout := "2006-01-02T15:04"

		dateTime, err := time.Parse(layout, dateTimeStr)
		if err != nil {
			return time.Time{}, err
		}
		return dateTime, nil
	} else {
		html, _ := s.Html()
		return time.Time{}, fmt.Errorf("getDate: date not found, selection: %s",
			html)
	}

}

func (venue Orgelpark) processAgendaItem(s *goquery.Selection, concerts *[]Concert) {
	fmt.Println("process")
	date, err := venue.getDate(s)

	if err != nil {
		log.Printf("processAgendaItem: extractdate: %s, error: %s", date, err)
		return
	}
	title, err := venue.getTitle(s)
	if err != nil {
		log.Printf("processAgendaItem: title: %s, error: %s", title, err)
		return
	}

	*concerts = append(*concerts, Concert{
		Description: title,
		Date:        date,
		Venue:       Venue(venue),
	})
}

func (venue Orgelpark) GetConcerts() []Concert {
	resp, err := http.Get(venue.Url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal("GetConcerts", err)
	}

	agendaItems := doc.Find("section.item-list.agenda article")
	concerts := []Concert{}
	agendaItems.Each(func(i int, s *goquery.Selection) {
		fmt.Println("i: ", i)
		venue.processAgendaItem(s, &concerts)
	})
	return concerts
}
