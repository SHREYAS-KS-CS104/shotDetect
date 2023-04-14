package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/SHREYAS-KS-CS104/shotDetect/controllers"
	"github.com/SHREYAS-KS-CS104/shotDetect/models"
	"github.com/SHREYAS-KS-CS104/shotDetect/templates"
	"github.com/SHREYAS-KS-CS104/shotDetect/views"
	"github.com/go-chi/chi/v5"
)

func executeTemplate(w http.ResponseWriter, filepath string) {
	t, err := views.Parse(filepath)
	if err != nil {
		log.Printf("parsing templates: %v", err)
		http.Error(w, "There was an erro parsing the template", http.StatusInternalServerError)
		return
	}
	t.Execute(w, nil)
}

func main() {
	r := chi.NewRouter()
	tpl := views.Must(views.ParseFS(templates.FS, "home.gohtml", "tailwind.gohtml"))
	r.Get("/", controllers.StaticHandler(tpl))

	tpl = views.Must(views.ParseFS(templates.FS, "contact.gohtml", "tailwind.gohtml"))
	r.Get("/contact", controllers.StaticHandler(tpl))

	tpl = views.Must(views.ParseFS(templates.FS, "faq.gohtml", "tailwind.gohtml"))
	r.Get("/faq", controllers.FAQ(tpl))

	cfg := models.DefaultPostgresConfig()
	db, err := models.Open(cfg)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	userService := models.UserService{
		DB: db,
	}

	usersC := controllers.Users{
		UserService: &userService,
	}
	usersC.Templates.New = views.Must(views.ParseFS(
		templates.FS,
		"signup.gohtml", "tailwind.gohtml",
	))
	r.Get("/signup", usersC.New)
	r.Post("/users", usersC.Create)

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})
	fmt.Println("Starting the server on :3000...")
	err = http.ListenAndServe(":3000", r)
	if err != nil {
		panic(err)
	}
}
