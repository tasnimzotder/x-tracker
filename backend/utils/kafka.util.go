package utils

import (
	"fmt"
	"github.com/IBM/sarama"
)

const (
	ProducerPort       = ":8080"
	KafkaServerAddress = "localhost:9092"
	KafkaTopic         = "notifications"
)

type YourMessageHandler struct{}

func (YourMessageHandler) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (YourMessageHandler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }
func (YourMessageHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		fmt.Printf("Message: %s, Offset: %d\n", msg.Value, msg.Offset)

		// YOUR LOGIC TO HANDLE LOCATION UPDATES HERE

		sess.MarkMessage(msg, "") // Mark message as processed
	}
	return nil
}

func SetupKafka() (sarama.SyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true

	broker := []string{KafkaServerAddress}

	producer, err := sarama.NewSyncProducer(broker, config)
	if err != nil {
		//return nil, fmt.Errorf("new sync producer: %w", err)
	}

	//group, err := sarama.NewConsumerGroup(broker, "group-1", config)
	//if err != nil {
	//
	//}
	//
	//defer func() {
	//	_ = group.Close()
	//}()
	//
	//ctx, cancel := context.WithCancel(context.Background())
	//
	//handler := YourMessageHandler{}
	//
	//go func() {
	//	for {
	//		if err := group.Consume(ctx, []string{KafkaTopic}, handler); err != nil {
	//			fmt.Println("Error from consumer: ", err)
	//		}
	//		if ctx.Err() != nil {
	//			return
	//		}
	//	}
	//}()
	//
	//signals := make(chan os.Signal, 1)
	//signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	//<-signals
	//cancel()
	//
	return producer, nil

}

func SendKafkaMessage(
	producer sarama.SyncProducer,
	topic string,
	key string,
	value string,
) error {

	msg := &sarama.ProducerMessage{
		Topic: topic,
		Key:   sarama.StringEncoder(key),
		Value: sarama.StringEncoder(value),
	}

	_, _, err := producer.SendMessage(msg)
	if err != nil {
		return fmt.Errorf("send message: %w", err)
	}

	return nil
}
