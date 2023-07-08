package html

import (
	"io/ioutil"
	"log"
	"os"
	"text/template"
	"time"

	"github.com/uiltjesrups/whats-on-in-amsterdam/internal/concerts"
)

func readTemplate() string {
	filePath := "resources/index.template.html"

	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	templateContent := string(content)

	return templateContent
}

type TemplateData struct {
	CurrentTime     time.Time
	GroupedConcerts groupedConcerts
}

type groupedConcerts map[time.Time][]concerts.Concert

func groupConcertsByDate(concerts []concerts.Concert) groupedConcerts {
	groupedConcerts := make(groupedConcerts)

	for _, concert := range concerts {
		date := concert.Date.UTC().Truncate(24 * time.Hour)
		groupedConcerts[date] = append(groupedConcerts[date], concert)
	}

	return groupedConcerts
}

func WriteHTML(concerts []concerts.Concert) {
	tmpl := template.Must(template.New("html").Parse(readTemplate()))

	file, err := os.Create("index.html")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	err = tmpl.Execute(file,
		TemplateData{CurrentTime: time.Now(),
			GroupedConcerts: groupConcertsByDate(concerts),
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("HTML file has been written successfully")
}
