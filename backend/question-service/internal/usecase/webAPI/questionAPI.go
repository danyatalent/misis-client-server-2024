package webAPI

import (
	"encoding/xml"
	"fmt"
	"github.com/danyatalent/misis-client-server-2024/backend/question-service/internal/models"
	"io"
	"log/slog"
	"net/http"
)

const (
	URL = "https://db.chgk.info/xml/random/types1/complexity2"
)

type QuestionWebAPI struct {
	logger *slog.Logger
}

func New(logger *slog.Logger) *QuestionWebAPI {
	return &QuestionWebAPI{
		logger: logger,
	}
}

func (w *QuestionWebAPI) GetOneQuestion() (*models.Question, error) {
	// Request the XML.
	op := "QuestionWebAPI.GetOneQuestion"
	w.logger = w.logger.With(op)

	res, err := http.Get(URL)
	if err != nil {
		w.logger.Error("error requesting XML")
		return nil, fmt.Errorf("error requesting XML: %w", err)
	}

	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		w.logger.Error("error status code")
		return nil, fmt.Errorf("error status code: %d", res.StatusCode)
	}
	data, err := io.ReadAll(res.Body)
	if err != nil {
		w.logger.Error("error reading body")
		return nil, fmt.Errorf("error reading body: %w", err)
	}
	search := SearchDTO{}
	err = xml.Unmarshal(data, &search)

	if err != nil {
		w.logger.Error("error unmarshalling data")
		return nil, fmt.Errorf("error unmarshalling data: %w", err)
	}
	//

	m := toDomain(&search.Questions[0])
	//return questions, nil
	return m, nil
}

func (w *QuestionWebAPI) GetAllQuestions() ([]*models.Question, error) {
	op := "QuestionWebAPI.GetAllQuestions"
	w.logger = w.logger.With(op)
	res, err := http.Get(URL)
	if err != nil {
		w.logger.Error("error requesting XML")
		return nil, fmt.Errorf("error requesting XML: %w", err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		w.logger.Error("error status code")
		return nil, fmt.Errorf("error status code: %d", res.StatusCode)
	}
	data, err := io.ReadAll(res.Body)
	if err != nil {
		w.logger.Error("error reading body")
		return nil, fmt.Errorf("error reading body: %w", err)
	}
	search := SearchDTO{}
	err = xml.Unmarshal(data, &search)
	if err != nil {
		w.logger.Error("error unmarshalling data")
		return nil, fmt.Errorf("error unmarshalling data: %w", err)
	}
	questions := make([]*models.Question, len(search.Questions))
	for i, q := range search.Questions {
		questions[i] = toDomain(&q)
	}
	return questions, nil
}

type QuestionDTO struct {
	XMLName             xml.Name `xml:"question"`
	TourFileName        string   `xml:"tourFileName"`
	TournamentFileName  string   `xml:"tournamentFileName"`
	QuestionID          string   `xml:"QuestionId"`
	ParentID            string   `xml:"ParentId"`
	Number              int      `xml:"Number"`
	Type                string   `xml:"Type"`
	TypeNum             int      `xml:"TypeNum"`
	TextID              string   `xml:"TextId"`
	Question            string   `xml:"Question"`
	Answer              string   `xml:"Answer"`
	PassCriteria        string   `xml:"PassCriteria"`
	Authors             string   `xml:"Authors"`
	Sources             string   `xml:"Sources"`
	Comments            string   `xml:"Comments"`
	Rating              string   `xml:"Rating"`
	RatingNumber        string   `xml:"RatingNumber"`
	Complexity          string   `xml:"Complexity"`
	Topic               string   `xml:"Topic"`
	ProcessedBySearch   string   `xml:"ProcessedBySearch"`
	ParentTextID        string   `xml:"parent_text_id,ParentTextId"`
	TourID              string   `xml:"tourId"`
	TournamentID        string   `xml:"tournamentId"`
	TourTitle           string   `xml:"tourTitle"`
	TournamentTitle     string   `xml:"tournamentTitle"`
	TourType            string   `xml:"tourType"`
	TournamentType      string   `xml:"tournamentType"`
	TourPlayedAt        string   `xml:"tourPlayedAt"`
	TournamentPlayedAt  string   `xml:"tournamentPlayedAt"`
	TourPlayedAt2       string   `xml:"tourPlayedAt2"`
	TournamentPlayedAt2 string   `xml:"tournamentPlayedAt2"`
	Notices             string   `xml:"Notices"`
}

type SearchDTO struct {
	XMLName   xml.Name      `xml:"search"`
	Questions []QuestionDTO `xml:"question"`
}

func toDomain(q *QuestionDTO) *models.Question {
	return &models.Question{
		ID:      q.QuestionID,
		Text:    q.Question,
		Answer:  q.Answer,
		Comment: q.Comments,
		Author:  q.Authors,
	}
}
