package zeitquiz

import (
	"fmt"
	"math/rand"
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

func EvaluateResult(pointsScored int, result Result) string {
	return fmt.Sprintf("\n### Ergebnis \U0001F440 ###\n"+
		"Wow, du hast %d Punkte erreicht.\n"+
		"Der Durchschnitt lag bei %.1f Punkten.\n"+
		"\033[32mDamit bist du besser als %d%% der Spieler,\n"+
		"\033[31mjedoch schlechter als %d%% der Spieler.\033[37m\n",
		pointsScored, result.Stats.Average, result.Stats.BetterThan, result.Stats.WorseThan)
}
