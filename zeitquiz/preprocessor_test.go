package zeitquiz

import (
	"testing"
)

func TestIncludeLineBreaks(t *testing.T) {
	tests := []struct {
		given string
		want  string
	}{
		{
			"Paare feiern heute hingegen den Tag der Liebenden.",
			"Paare feiern heute hingegen den Tag der Liebenden.",
		},
		{
			"Für viele Menschen beginnt heute die Fastenzeit. Welches Gericht soll laut einer Legende entstanden sein, weil ein Mönch das Fleischverbot hatte umgehen wollen?",
			"Für viele Menschen beginnt heute die Fastenzeit. Welches Gericht soll laut einer Legende\nentstanden sein, weil ein Mönch das Fleischverbot hatte umgehen wollen?",
		},
	}

	for _, test := range tests {
		got := includeLineBreaks(test.given)
		if got != test.want {
			t.Fatalf("Got %s, not want %s", got, test.want)
		}
	}
}

func TestGetCorrectAnswer(t *testing.T) {
	given := []Answer{
		{
			ID:      1,
			Text:    "wrong text",
			Correct: false,
		},
		{
			ID:      2,
			Text:    "correct text",
			Correct: true,
		},
	}

	want := "correct text"
	got := getCorrectAnswer(given)
	if got != want {
		t.Fatalf("Got %s, not want %s", got, want)
	}
}

func TestEvaluateResult(t *testing.T) {
	givenPoints := 6
	givenResult := Result{
		Stats: Stats{
			Average: Average{
				MeanPoints:  "5.5",
				MeanTime:    "3000",
				TotalPlayed: 300,
			},
		},
	}

	want := "\nErgebnis 👀\nWow, du hast 6 Punkte erreicht.\nDer Durchschnitt lag bei 5.5 Punkten.\n300 Spieler haben teilgenommen.\n"
	got, _ := EvaluateResult(givenPoints, givenResult)
	if want != got {
		t.Fatalf("Got %s, not want %s", got, want)
	}
}
