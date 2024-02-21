package main

import (
	"fmt"
	"strconv"

	"github.com/fabenan-f/zeitquiz/zeitquiz"
)

const banner string = "#######                   #######                                   #####                 \n     #  ###### # #####    #     # #    # #      # #    # ######    #     # #    # # ######\n    #   #      #   #      #     # ##   # #      # ##   # #         #     # #    # #     # \n   #    #####  #   #      #     # # #  # #      # # #  # #####     #     # #    # #    #  \n  #     #      #   #      #     # #  # # #      # #  # # #         #   # # #    # #   #   \n #      #      #   #      #     # #   ## #      # #   ## #         #    #  #    # #  #    \n####### ###### #   #      ####### #    # ###### # #    # ######     #### #  ####  # ######"

func main() {
	fmt.Println(banner)

	zeitClient := zeitquiz.NewZeitClient()
	url := "https://quiz.zeit.de/quizzes/daily"
	ids, err := zeitClient.GetQuizIds(url)
	if err != nil {
		return
	}

	for _, id := range ids {
		ok := zeitquiz.PromptReady()
		if !ok {
			return
		}
		quizUrl := "https://quiz.zeit.de/states?quizId=" + strconv.Itoa(id)
		quiz, err := zeitClient.GetQuiz(quizUrl)
		if err != nil {
			return
		}

		playerResult, err := zeitquiz.PromptQuiz(quiz)
		if err != nil {
			return
		}
		resultUrl := "https://quiz.zeit.de/results?quizId=" + strconv.Itoa(quiz.ID)
		result, err := zeitClient.PostPlayerResult(playerResult, resultUrl, "")
		if err != nil {
			return
		}

		evaluation, err := zeitquiz.EvaluateResult(playerResult.PointsScored, result)
		if err != nil {
			return
		}

		fmt.Println(evaluation)
	}
}
