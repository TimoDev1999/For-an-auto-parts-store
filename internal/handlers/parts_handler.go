package handlers

import (
	"encoding/json"
	"log"
	"my-go-project/internal/services"
	"my-go-project/pkg/messaging"
	"net/http"

	"github.com/rabbitmq/amqp091-go"
)

func GetParts(w http.ResponseWriter, r *http.Request) {
	parts, err := services.GetAllParts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Пример отправки сообщения в RabbitMQ
	body, err := json.Marshal(parts)
	if err != nil {
		log.Printf("Error marshalling parts: %s", err)
	} else {
		err = messaging.GetChannel().Publish(
			"",      // exchange
			"parts", // routing key
			false,   // mandatory
			false,   // immediate
			amqp091.Publishing{
				ContentType: "application/json",
				Body:        body,
			})
		if err != nil {
			log.Printf("Error publishing message: %s", err)
		}
	}

	json.NewEncoder(w).Encode(parts)
}
