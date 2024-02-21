package zeitquiz

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
)

func includeLineBreaks(question string) string {
	var lastSpace int
	raw := []rune(question)
	for i, char := range raw {
		if char == 32 {
			lastSpace = i
		}
		if i != 0 && i%93 == 0 {
			raw[lastSpace] = 10
		}
	}
	return string(raw)
}

func shuffleAnswers(answers []Answer) []Answer {
	shuffledAnswers := make([]Answer, len(answers))
	random := rand.Perm(3)
	for i, answer := range answers {
		shuffledAnswers[random[i]] = answer
	}
	return shuffledAnswers
}

func getCorrectAnswer(answers []Answer) string {
	for _, answer := range answers {
		if answer.Correct {
			return answer.Text
		}
	}
	return ""
}

func EvaluateResult(pointsScored int, result Result) (string, error) {
	meanPoints, err := strconv.ParseFloat(result.Stats.Average.MeanPoints, 64)
	if err != nil {
		log.Printf("Error %s when parsing mean points", err)
		return "", err
	}

	return fmt.Sprintf("\nErgebnis \U0001F440\n"+
		"Wow, du hast %d Punkte erreicht.\n"+
		"Der Durchschnitt lag bei %.1f Punkten.\n"+
		"%d Spieler haben teilgenommen.\n",
		pointsScored, meanPoints, result.Stats.Average.TotalPlayed), nil
}
