package controllers

import (
	"html/template"
	"net/http"
)

func StaticHandler(tpl Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, r, nil)
	}
}

func FAQ(tpl Template) http.HandlerFunc {
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
		tpl.Execute(w, r, questions)
	}
}
