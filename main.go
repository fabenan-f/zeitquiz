package main

import (
	"fmt"
	"strconv"

	"github.com/fabenan-f/zeitquiz/zeitquiz"
)

const banner string = "#######                   #######                                   #####                 \n     #  ###### # #####    #     # #    # #      # #    # ######    #     # #    # # ######\n    #   #      #   #      #     # ##   # #      # ##   # #         #     # #    # #     # \n   #    #####  #   #      #     # # #  # #      # # #  # #####     #     # #    # #    #  \n  #     #      #   #      #     # #  # # #      # #  # # #         #   # # #    # #   #   \n #      #      #   #      #     # #   ## #      # #   ## #         #    #  #    # #  #    \n####### ###### #   #      ####### #    # ###### # #    # ######     #### #  ####  # ######\n"

func main() {
	fmt.Printf(banner)

	zeitClient := zeitquiz.NewZeitClient()
	url := "https://quiz.zeit.de/-/quizzes/daily"

	for {
		ok := zeitquiz.PromptReady()
		if !ok {
			return
		}
		quiz, err := zeitClient.GetQuiz(url)
		if err != nil {
			return
		}

		playerResult, err := zeitquiz.PromptQuiz(quiz)
		if err != nil {
			return
		}
		resultUrl := "https://quiz.zeit.de/-/quiz/" + strconv.Itoa(quiz.ID) + "/result"
		result, err := zeitClient.PostPlayerResult(playerResult, resultUrl)
		if err != nil {
			return
		}

		fmt.Println(zeitquiz.EvaluateResult(playerResult.PointsScored, result))

		url = result.NextQuiz
	}
}
