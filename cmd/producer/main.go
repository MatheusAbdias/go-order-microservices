package main

import (
	"encoding/json"
	"math/rand"
	"time"

	"github.com/MatheusAbdias/microservices/internal/order/domain"
	"github.com/google/uuid"
	amqp "github.com/streadway/amqp"
)

func Publish(ch *amqp.Channel, order domain.Order) error {
	body, err := json.Marshal(order)
	if err != nil {
		return err
	}

	err = ch.Publish(
		"amq.direct",
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	if err != nil {
		return err
	}
	return nil
}

func OrdersFactory() domain.Order {
	return domain.Order{
		ID:    uuid.New().String(),
		Price: rand.Float64() * 100,
		Tax:   rand.Float64() * 10,
	}
}
func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	for i := 0; i < 10000000; i++ {
		Publish(ch, OrdersFactory())
		time.Sleep(300 * time.Millisecond)
	}
}
