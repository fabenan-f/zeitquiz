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

type Daily struct {
	Quizzes []DailyQuizOverview `json:"quizzes"`
}

type DailyQuizOverview struct {
	Id int `json:"id"`
}

type QuizOverview struct {
	Quiz Quiz `json:"quiz"`
}

type Quiz struct {
	ID        int        `json:"id"`
	Questions []Question `json:"questions"`
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
	AnswerID    int `json:"answer_id"`
	PointScored int `json:"points_scored"`
	QuestionID  int `json:"question_id"`
	TimeTaken   int `json:"time_taken_ms"`
}

type PlayerResult struct {
	Details      []Details `json:"details"`
	PointsScored int       `json:"points_scored"`
	TotalTime    int       `json:"total_time"`
}

type Stats struct {
	Average Average `json:"average"`
}

type Average struct {
	MeanPoints  string `json:"mean_points"`
	TotalPlayed int    `json:"total_played"`
	MeanTime    string `json:"mean_time"`
}

type Result struct {
	Stats Stats  `json:"stats"`
	UUID  string `json:"uuid"`
}

type Ranking struct {
	Name       string `json:"name"`
	ResultUUID string `json:"resultUUID"`
}

func NewZeitClient() ZeitClient {
	return ZeitClient{
		http: http.Client{},
	}
}

func (z *ZeitClient) GetQuizIds(url string) ([]int, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf("Error %s when creating request", err)
		return nil, err
	}
	req.Header.Set("Accept", "application/json")

	res, err := z.http.Do(req)
	if err != nil {
		log.Printf("Error %s when getting quiz", err)
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Printf("Error %s when reading body from response", err)
		return nil, err
	}

	daily := Daily{}
	err = json.Unmarshal(body, &daily)
	if err != nil {
		log.Printf("Error %s when unmarshalling response", err)
		return nil, err
	}

	return getIds(daily.Quizzes), nil
}

func getIds(quizzes []DailyQuizOverview) []int {
	ids := []int{}
	for _, quiz := range quizzes {
		ids = append(ids, quiz.Id)
	}
	return ids
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

	quizOverview := QuizOverview{}
	err = json.Unmarshal(body, &quizOverview)
	if err != nil {
		log.Printf("Error %s when unmarshalling response", err)
		return Quiz{}, err
	}

	return quizOverview.Quiz, nil
}

func (z *ZeitClient) PostPlayerResult(playerResult PlayerResult, url string, cookies string) (Result, error) {
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
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cookie", cookies)

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

func (z *ZeitClient) PostRanking(ranking Ranking, url, cookies string) (string, error) {
	rankingJson, err := json.Marshal(ranking)
	if err != nil {
		log.Printf("Error %s when marshaling ranking request", err)
		return "", err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(rankingJson))
	if err != nil {
		log.Printf("Error %s when creating ranking request", err)
		return "", err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cookie", cookies)

	res, err := z.http.Do(req)
	if err != nil {
		log.Printf("Error %s when posting results", err)
		return "", err
	}

	return res.Status, nil
}
