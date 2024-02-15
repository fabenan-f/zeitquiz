package zeitquiz

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type ZeitClient struct {
	http http.Client
}

type Quiz struct {
	ID          int        `json:"id"`
	Description string     `json:"description"`
	Questions   []Question `json:"questions"`
}

type Question struct {
	ID       int      `json:"id"`
	Question string   `json:"question"`
	Answers  []Answer `json:"answers"`
}

type Answer struct {
	ID      int    `json:"id"`
	Text    string `json:"text"`
	Correct bool   `json:"correct"`
}

type Details struct {
	AnswerID    int    `json:"answer_id"`
	PointScored int    `json:"points_scored"`
	QuestionID  int    `json:"question_id"`
	TimeTaken   int    `json:"time_taken_ms"`
	TimingClass string `json:"timing_class"`
}

type PlayerResult struct {
	Details      []Details `json:"details"`
	PointsScored int       `json:"points_scored"`
}

type Stats struct {
	Average    float64 `json:"average"`
	BetterThan int     `json:"better_than_percent"`
	WorseThan  int     `json:"worse_than_percent"`
}

type Result struct {
	Stats    Stats  `json:"stats"`
	NextQuiz string `json:"next_quiz"`
}

func NewZeitClient() ZeitClient {
	return ZeitClient{
		http: http.Client{},
	}
}

func (z *ZeitClient) GetQuiz(url string) (Quiz, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf("Error %s when creating request", err)
		return Quiz{}, err
	}
	req.Header.Set("Accept", "application/json")

	res, err := z.http.Do(req)
	if err != nil {
		log.Printf("Error %s when getting quiz", err)
		return Quiz{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Printf("Error %s when reading body from response", err)
		return Quiz{}, err
	}

	quiz := Quiz{}
	err = json.Unmarshal(body, &quiz)
	if err != nil {
		log.Printf("Error %s when unmarshalling response", err)
		return Quiz{}, err
	}

	return quiz, nil
}

func (z *ZeitClient) PostPlayerResult(playerResult PlayerResult, url string) (Result, error) {
	playerResultJson, err := json.Marshal(playerResult)
	if err != nil {
		log.Printf("Error %s when marshalling result", err)
		return Result{}, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(playerResultJson))
	if err != nil {
		log.Printf("Error %s when creating request", err)
		return Result{}, err
	}
	req.Header.Set("Accept", "application/json")

	res, err := z.http.Do(req)
	if err != nil {
		log.Printf("Error %s when posting results", err)
		return Result{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Printf("Error %s when reading body from response", err)
		return Result{}, err
	}

	result := Result{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Printf("Error %s when unmarshalling response", err)
		return Result{}, err
	}

	return result, nil
}
