package concerts

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type Buiksloterkerk Venue

func (venue Buiksloterkerk) getTitle(s *goquery.Selection) (string, error) {
	title := s.Find("h5").Text()
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

func (venue Buiksloterkerk) getDate(s *goquery.Selection) (time.Time, error) {
	parseDateTime := func(input string) (time.Time, error) {
		layout := "2 January 2006 om 15:04"
		trimmedInput := strings.TrimSpace(input)

		parts := strings.Split(trimmedInput, " om ")
		if len(parts) != 2 {
			return time.Time{}, fmt.Errorf("invalid input format: %s", trimmedInput)
		}

		dateStr := strings.TrimSpace(parts[0])
		timeStr := strings.TrimSpace(parts[1])

		// Convert the Dutch month names to English
		for dutchMonth, englishMonth := range dutchMonthToEnglish {
			dateStr = strings.ReplaceAll(dateStr, dutchMonth, englishMonth)
		}

		// Parse the date and time
		parsedTime, err := time.Parse(layout, dateStr+" om "+timeStr)
		if err != nil {
			return time.Time{}, fmt.Errorf("error parsing input: %w", err)
		}

		return parsedTime, nil
	}
	date := s.Find("h6").Text()
	if date == "" {
		html, err := s.Html()
		if err != nil {
			return time.Time{},
				fmt.Errorf("getDateB: date not found, selection: %s", err)

		}
		return time.Time{},
			fmt.Errorf("getTitle: title not found, selection: %s", html)
	}

	datetime, err := parseDateTime(date)
	if err != nil {
		return time.Time{}, err
	}
	return datetime, nil
}

func (venue Buiksloterkerk) processAgendaItem(s *goquery.Selection, concerts *[]Concert) {
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

func (venue Buiksloterkerk) GetConcerts() []Concert {
	resp, err := http.Get(venue.Url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal("GetConcerts", err)
	}

	agendaItems := doc.Find("div.card-body")

	concerts := []Concert{}
	agendaItems.Each(func(i int, s *goquery.Selection) {
		venue.processAgendaItem(s, &concerts)
	})
	return concerts
}

var dutchMonthToEnglish = map[string]string{
	"januari":   "January",
	"februari":  "February",
	"maart":     "March",
	"april":     "April",
	"mei":       "May",
	"juni":      "June",
	"juli":      "July",
	"augustus":  "August",
	"september": "September",
	"oktober":   "October",
	"november":  "November",
	"december":  "December",
}
