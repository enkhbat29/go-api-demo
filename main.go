package main

import (
	"encoding/json"
	"net/http"
)

// Хэрэглэгчийн бүтцийг тодорхойлох (JSON тагуудын хамт)
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Туршилтын өгөгдөл (Мэдээллийн сангийн үүрэг гүйцэтгэнэ)
var users = []User{
	{ID: 1, Name: "Бат", Email: "bat@example.com"},
	{ID: 2, Name: "Сувд", Email: "suvd@example.com"},
}

func getUsersHandler(w http.ResponseWriter, r *http.Request) {
	// Буцаах өгөгдлийг JSON форматтай гэдгийг хөтөчид мэдэгдэнэ
	w.Header().Set("Content-Type", "application/json")
	
	// Амжилттай болсныг илтгэх 200 OK статус код илгээнэ
	w.WriteHeader(http.StatusOK)
	
	// users жагсаалтыг JSON болгон хөрвүүлж, хариу (w) руу бичнэ
	json.NewEncoder(w).Encode(users)
}

func main() {
	// Маршрут (Route) тохируулах: /users хаяг дээр getUsersHandler-ийг ажиллуулна
	http.HandleFunc("/users", getUsersHandler)

	// Серверийг 8080 порт дээр асаах
	println("Сервер 8080 порт дээр ажиллаж эхэллээ...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

