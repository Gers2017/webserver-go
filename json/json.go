package json

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Name           string `json:"name"`
	Password       string `json:"password"`
	FavoriteNumber int    `json:"favoriteNumber"`
	FavoriteAnimal string `json:"favoriteAnimal"`
}

func getValue(value string, fallabackValue string) string {
	if len(value) == 0 {
		return fallabackValue
	}
	return value
}

func getNumberValue(value string, fallabackNumber int) int {
	value = getValue(value, strconv.Itoa(fallabackNumber))
	number, err := strconv.Atoi(value)
	if err != nil {
		number = fallabackNumber
	}
	return number
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func RunServer() {
	http.HandleFunc("/decode", func(w http.ResponseWriter, r *http.Request) {
		var user User
		hashed, _ := hashPassword(user.Password)

		if checkPasswordHash(user.Password, hashed) {
			println("Passwords match!")
		}

		json.NewDecoder(r.Body).Decode(&user)
		fmt.Fprintf(w, "All data submitted! name:%s hashed password: %s, fav number:%d, fav animal: %s", user.Name, hashed, user.FavoriteNumber, user.FavoriteAnimal)
	})

	http.HandleFunc("/encode", func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()

		name := getValue(q.Get("name"), "John Doe")
		password := getValue(q.Get("password"), "SecurePassword")
		favoriteAnimal := getValue(q.Get("favoriteAnimal"), "Octopus 🐙🐙🐙")
		favoriteNumber := getNumberValue(q.Get("favoriteNumber"), 69)

		peter := User{
			Name:           name,
			Password:       password,
			FavoriteNumber: favoriteNumber,
			FavoriteAnimal: favoriteAnimal,
		}

		json.NewEncoder(w).Encode(peter)
	})

	port := ":80"
	fmt.Printf("Serving json at http://localhost%s/\n", port)
	http.ListenAndServe(port, nil)
}
