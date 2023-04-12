package controllers

import (
	"html/template"
	"net/http"

	"github.com/SHREYAS-KS-CS104/shotDetect/views"
)

func StaticHandler(tpl views.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, nil)
	}
}

func FAQ(tpl views.Template) http.HandlerFunc {
	questions := []struct {
		Question string
		Answer   template.HTML
	}{
		{
			Question: "What is your name?",
			Answer:   "Yes!",
		},
		{
			Question: "Where are you from?",
			Answer:   "Yes!",
		},
		{
			Question: "What do you do?",
			Answer:   "Yes!",
		},
		{
			Question: "What is What?",
			Answer:   "Yes!",
		},
	}
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, questions)
	}
}
