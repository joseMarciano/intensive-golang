package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/joseMarciano/intensive-golang/internal/order/infra/database"
	"github.com/joseMarciano/intensive-golang/internal/order/usecase"
	"github.com/joseMarciano/intensive-golang/package/rabbitmq"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(mysql:3306)/orders")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	repository := database.NewOrderRepository(db)
	calculateFinalPriceUseCase := usecase.NewCalculateFinalPriceUseCase(repository)

	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}

	defer ch.Close()

	deliveryMessage := make(chan amqp.Delivery)
	forever := make(chan bool)

	go rabbitmq.Consume(ch, deliveryMessage)
	go worker(deliveryMessage, calculateFinalPriceUseCase, 1)
	go worker(deliveryMessage, calculateFinalPriceUseCase, 2)
	go worker(deliveryMessage, calculateFinalPriceUseCase, 3)

	<-forever
}

func worker(deliveryMessage <-chan amqp.Delivery, useCase *usecase.CalculateFinalPriceUseCase, workerId int) {
	for message := range deliveryMessage {
		var input = usecase.OrderInputDTO{}
		err := json.Unmarshal(message.Body, &input)
		if err != nil {
			fmt.Printf("Error on unmarshal %v", err)
		}

		input.Tax = 5.54

		_, err = useCase.Execute(input)

		if err != nil {
			panic(err)
		}

		println("Received message: ", workerId)
		message.Ack(false)
	}
}
