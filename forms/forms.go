package forms

import (
	"fmt"
	"html/template"
	"net/http"
)

type ContactDetails struct {
	Email   string
	Subject string
	Message string
	Animal  string
}

type SuccessDetails struct {
	Success bool
	Animal  string
}

func getAnimalEmoji(animal string) string {
	switch animal {
	case "chicken":
		return "🐔"
	case "fish":
		return "🐟"
	case "kraken":
		return "🐙"
	default:
		return "🙀"
	}
}

func RunServer() {
	// check template.ParseFS
	tmpl := template.Must(template.ParseFiles("public/form.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
		}

		details := ContactDetails{
			Email:   r.FormValue("email"),
			Subject: r.FormValue("subject"),
			Message: r.FormValue("message"),
			Animal:  r.FormValue("animal"),
		}

		fmt.Printf("Email: %s\nSubject: %s\n", details.Email, details.Subject)
		fmt.Printf("Message: %s\nAnimal: %s\n", details.Message, details.Animal)

		emoji := getAnimalEmoji(details.Animal)
		animalMsg := fmt.Sprintf("%s %s", details.Animal, emoji)

		tmpl.Execute(w, SuccessDetails{Success: true, Animal: animalMsg})
	})
	port := ":80"

	fmt.Printf("Serving files at http://localhost%s\n", port)
	http.ListenAndServe(port, nil)
}
