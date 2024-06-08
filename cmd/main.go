package main

import (
	"fmt"
	"log"
	"my-go-project/internal/handlers"
	"my-go-project/pkg/database"
	"my-go-project/pkg/messaging"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq" // Импорт с побочным эффектом для регистрации драйвера PostgreSQL
)

func main() {
	dataSourceName := fmt.Sprintf("host=db port=5432 user=user password=password dbname=autoparts sslmode=disable")
	database.InitDB(dataSourceName)

	rabbitMQURL := "amqp://user:password@rabbitmq:5672/"
	messaging.InitRabbitMQ(rabbitMQURL)
	defer messaging.CloseRabbitMQ()

	r := mux.NewRouter()
	r.HandleFunc("/parts", handlers.GetParts).Methods("GET")
	r.HandleFunc("/suppliers", handlers.GetSuppliers).Methods("GET")
	r.HandleFunc("/customers", handlers.GetCustomers).Methods("GET")

	http.Handle("/", r)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server is running on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), r))
}
