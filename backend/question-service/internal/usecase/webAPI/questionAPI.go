package webAPI

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/danyatalent/misis-client-server-2024/backend/question-service/internal/models"
	"log"
	"net/http"
	"strings"
)

func GetQuestion() ([]models.Question, error) {
	// Request the HTML page.
	res, err := http.Get("https://db.chgk.info/random/answers/types1/complexity2")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	questions := make([]models.Question, 0)

	ss := doc.Find(".random_question").First()
	clone := ss.Clone()
	q := clone.Children().Remove().End().Text()
	q = strings.TrimSpace(q)
	a := ss.Find("p").First().Next().Next().Text()
	log.Println("quest: ", q)
	log.Println("ans: ", a)

	// Find the review items
	doc.Find(".random_question").Each(func(i int, s *goquery.Selection) {
		clone := s.Clone()
		question := clone.Children().Remove().End().Text()
		question = strings.TrimSpace(question)
		answer := s.Find("p").First().Next().Next().Text()
		questions = append(questions, models.Question{
			Text:   question,
			Answer: answer,
		})
		//log.Println(question)

	})
	return questions, nil
}
