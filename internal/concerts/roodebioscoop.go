package concerts

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type Roodebioscoop Venue

func (venue Roodebioscoop) getTitle(s *goquery.Selection) (string, error) {
	title := s.Find(".agenda-item-title").Text()
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

func (venue Roodebioscoop) getDate(s *goquery.Selection) (time.Time, error) {
	dutchMonthToEnglish = map[string]string{
		"jan": "January",
		"feb": "February",
		"maa": "March",
		"apr": "April",
		"mei": "May",
		"jun": "June",
		"jul": "July",
		"aug": "August",
		"sep": "September",
		"okt": "October",
		"nov": "November",
		"dec": "December",
	}
	parseDateTime := func(input string) (time.Time, error) {
		layout := "02 January 2006 15:04"

		// Split the string by white space

		// Parse the date and time
		parsedTime, err := time.Parse(layout, input)
		if err != nil {
			return time.Time{}, fmt.Errorf("error parsing input: %w", err)
		}

		return parsedTime, nil
	}
	dateStr := strings.TrimSpace(s.Find(".agenda-item-day").Text())[3:]

	timeStr := strings.TrimSpace(s.Find(".agenda-item-time").Text()[:6])

	// Split the string by white space
	parts := strings.Fields(dateStr)
	parts[len(parts)-1] = dutchMonthToEnglish[parts[len(parts)-1]]

	currentYear := time.Now().Year()
	currentYearStr := strconv.Itoa(currentYear)
	dateTimeStr := strings.Join(parts, " ") + " " + currentYearStr + " " + timeStr

	if dateTimeStr == "" {
		html, err := s.Html()
		if err != nil {
			return time.Time{},
				fmt.Errorf("getDateB: date not found, selection: %s", err)

		}
		return time.Time{},
			fmt.Errorf("getTitle: title not found, selection: %s", html)
	}

	datetime, err := parseDateTime(dateTimeStr)
	if err != nil {
		return time.Time{}, err
	}
	return datetime, nil
}

func (venue Roodebioscoop) processAgendaItem(s *goquery.Selection, concerts *[]Concert) {
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

func (venue Roodebioscoop) GetConcerts() []Concert {
	resp, err := http.Get(venue.Url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal("GetConcerts", err)
	}

	agendaItems := doc.Find("div.agenda-item")
	concerts := []Concert{}
	agendaItems.Each(func(i int, s *goquery.Selection) {
		venue.processAgendaItem(s, &concerts)
	})
	return concerts
}
