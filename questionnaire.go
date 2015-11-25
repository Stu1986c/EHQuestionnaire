package main

import (
	"log"
	"net/http"
	"text/template"
)

type Answer struct {
	AnswerId   int
	AnswerText string
	Selected   bool
}

type Answers []Answer

type Question struct {
	QuestionId   int
	Answers      Answers
	QuestionText string
}

func loadPage() (*Question, error) {
	return &Question{
		QuestionId:   321,
		QuestionText: "Food safety training is required...",
		Answers: Answers{
			Answer{
				AnswerId:   1,
				AnswerText: "By law",
				Selected:   false,
			},
			Answer{
				AnswerId:   2,
				AnswerText: "Because it improves the image of the Winter Shelter",
				Selected:   false,
			},
			Answer{
				AnswerId:   3,
				AnswerText: "Although non-essential, it is a good idea",
				Selected:   false,
			},
		},
	}, nil
}

func ViewHandler(w http.ResponseWriter, r *http.Request) {

	const tpl = `
	<!DOCTYPE html>
	<html>
	<body>
		<form action="demo" method="POST">
			<h3>{{ .QuestionText }}</h3>
			{{range .Answers}}
				<input type="radio" name="answerGroup" {{if .Selected}}checked{{end}}>{{.AnswerText}}<br>
			{{end}}
		</form>
	</body>
	</html>
	`

	checkErr := func(err error) {
		if err != nil {
			log.Fatal(err)
		}
	}

	t, err := template.New("questionnaire").Parse(tpl)
	checkErr(err)

	pagedata, _ := loadPage()

	err = t.Execute(w, pagedata)
	checkErr(err)
}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", ViewHandler)
	log.Fatal(http.ListenAndServe(":8080", mux))
}
