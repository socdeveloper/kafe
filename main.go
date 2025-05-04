package main

import (
	"kafe/orderSubmitHandler"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	handleStatic()
	handleRoutes()
	handleCreateCard()
	handleSubmitOrder()

	log.Println("Server is running")
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func handleSubmitOrder() {
	http.HandleFunc("/submit-order", orderSubmitHandler.SubmitOrderHandler)
}

func handleStatic() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
}

func handleCreateCard() {
	http.HandleFunc("/create_card.html", func(w http.ResponseWriter, r *http.Request) {
		amount := r.URL.Query().Get("amount")

		if amount == "" {
			http.Error(w, "Нет значения", http.StatusBadRequest)
			return
		}

		http.ServeFile(w, r, "./static/create_card.html")
	})
}

func handleRoutes() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/index.html")
	})

	http.HandleFunc("/index.html", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/index.html")
	})

	http.HandleFunc("/gift.html", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/gift.html")
	})

	http.HandleFunc("/vacancii.pdf", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/vacancii.pdf")
	})

	http.HandleFunc("/menu.pdf", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/menu.pdf")
	})

	http.HandleFunc("/gift_choose.html", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/gift_choose.html")
	})

	http.HandleFunc("/success.html", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/success.html")
	})
}
