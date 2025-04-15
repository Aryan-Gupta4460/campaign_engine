package kafka

import (
	"campaign_engine/internal/utils"
	"context"

	"github.com/segmentio/kafka-go"
)

var Writer *kafka.Writer

func InitProducer() {
	Writer = &kafka.Writer{
		Addr:     kafka.TCP(utils.KafkaBroker),
		Topic:    "bids",
		Balancer: &kafka.LeastBytes{},
	}
}

func SendBidMessage(message []byte) error {
	utils.LogInfo("Sending message to kafka : " + string(message))
	return Writer.WriteMessages(context.Background(), kafka.Message{
		Value: message,
	})
}
