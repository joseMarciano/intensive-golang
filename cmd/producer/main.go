package main

import (
	"encoding/json"
	"github.com/google/uuid"
	amqp "github.com/rabbitmq/amqp091-go"
	"math/rand"
)

type Order struct {
	ID    string
	Price float64
}

func GenerateOrders() Order {
	return Order{
		ID:    uuid.NewString(),
		Price: float64(rand.Intn(50)),
	}
}

func Notify(ch *amqp.Channel, order Order) error {
	body, err := json.Marshal(order)
	if err != nil {
		return err
	}

	err = ch.Publish("amq.direct", "", false, false, amqp.Publishing{
		ContentType: "application/json",
		Body:        body,
	})

	if err != nil {
		return nil
	}

	return nil

}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
	if err != nil {
		panic(err)
	}

	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}

	defer conn.Close()
	defer ch.Close()

	for i := 0; i < 5000; i++ {
		order := GenerateOrders()
		err := Notify(ch, order)

		if err != nil {
			panic(err)
		}
		//fmt.Println(order)
	}
}
