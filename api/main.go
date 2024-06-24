package main

import (
	"context"
	"encoding/json"
	"html/template"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
)

type Order struct {
	Name           string
	Phone          string
	Address        string
	PreferableTime string
}

var tmpl = template.Must(template.ParseFiles("form.html"))

func main() {
	http.HandleFunc("/", formHandler)
	http.HandleFunc("/submit", submitHandler)
	http.HandleFunc("/map", mapHandler)
	http.HandleFunc("/api/orders", ordersHandler) // For fetching orders data
	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	tmpl.Execute(w, nil)
}

func submitHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	order := Order{
		Name:           r.FormValue("name"),
		Phone:          r.FormValue("phone"),
		Address:        r.FormValue("address"),
		PreferableTime: r.FormValue("preferableTime"),
	}

	// Save order to MongoDB (you will implement this)
	saveOrder(order)

	http.Redirect(w, r, "/map", http.StatusSeeOther)
}

func saveOrder(order Order) {
	client := ConnectDB()
	collection := client.Database("orderlocator").Collection("orders")

	_, err := collection.InsertOne(context.Background(), order)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Order saved to MongoDB")
}

func mapHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "map.html")
}

func ordersHandler(w http.ResponseWriter, r *http.Request) {
	client := ConnectDB()
	collection := client.Database("orderlocator").Collection("orders")

	cursor, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	var orders []Order
	if err := cursor.All(context.Background(), &orders); err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(orders)
}
