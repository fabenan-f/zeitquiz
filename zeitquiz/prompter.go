package zeitquiz

import (
	"fmt"
	"log"

	"github.com/manifoldco/promptui"
)

var (
	templateReady = &promptui.SelectTemplates{
		Label:    "{{ . | white }}",
		Active:   `{{ if eq . "Ja, bitte"}}` + "\U0001F64F {{ . | yellow }}{{ else }}\U0001F44B {{ . | yellow }}{{ end }}",
		Inactive: "{{ .Text | white }}",
		Selected: `{{ if eq . "Ja, bitte"}}` + "\U0001F4AA {{ . | green }}{{ else }}\U0001F44B {{ . | red }}{{ end }}",
	}
	templateQuestion = &promptui.SelectTemplates{
		Label:    "{{ .Text | white }}",
		Active:   "\U0001F914 {{ .Text | yellow }}",
		Inactive: "{{ .Text | white }}",
		Selected: `Deine Antwort: {{ if .Correct }}{{ .Text | green }}{{ else }}{{ .Text | red }}{{ end }}`,
	}
)

func PromptReady() bool {
	fmt.Println()
	fmt.Println("Bereit fürs (nächste) Quiz?")
	prompt := promptui.Select{
		Label:     "",
		Items:     []string{"Ja, bitte", "Nein, danke"},
		Templates: templateReady,
		HideHelp:  true,
	}
	_, answer, err := prompt.Run()
	if err != nil {
		log.Printf("Error %s when getting answer", err)
		return false
	}
	if answer != "Ja, bitte" {
		return false
	}
	return true
}

func PromptQuiz(quiz Quiz) (PlayerResult, error) {
	playerReturns := []Details{}
	var pointsScored int
	for _, question := range quiz.Questions {
		fmt.Println()
		shuffledAnswers := shuffleAnswers(question.Answers)
		fmt.Println(includeLineBreaks(question.Question))
		prompt := promptui.Select{
			Label:     "",
			Items:     shuffledAnswers,
			Templates: templateQuestion,
			HideHelp:  true,
		}

		answerIndex, _, err := prompt.Run()
		if err != nil {
			log.Printf("Error %s when getting answer", err)
			return PlayerResult{}, err
		}

		var pointScored int
		if shuffledAnswers[answerIndex].Correct {
			fmt.Println("Gratuliere, richtige Antwort!")
			pointScored++
		} else {
			fmt.Printf("Schade, richtige Antwort: %s \n", getCorrectAnswer(shuffledAnswers))
		}
		pointsScored += pointScored

		playerReturn := Details{
			AnswerID:    shuffledAnswers[answerIndex].ID,
			PointScored: pointScored,
			QuestionID:  question.ID,
			TimeTaken:   1000,
		}
		playerReturns = append(playerReturns, playerReturn)
	}

	return PlayerResult{
		Details:      playerReturns,
		PointsScored: pointsScored,
		TotalTime:    1000,
	}, nil
}
