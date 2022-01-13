package kafka1

import (
	"context"
	"fmt"
	kafkaClient "github.com/rezaAmiri123/test-microservice/pkg/kafka"
	"github.com/rezaAmiri123/test-microservice/pkg/logger/applogger"
	kafkaMessages "github.com/rezaAmiri123/test-microservice/user_service/proto/kafka"
	"github.com/segmentio/kafka-go"
	"google.golang.org/protobuf/proto"
	kafka2 "github.com/rezaAmiri123/test-microservice/pkg/kafka"
	"time"
)

const (
	topic          = "create_user"
	broker1Address = "localhost:9092"
)

func Produce(ctx context.Context) {
	// initialize a counter
	i := 0
	log := applogger.NewAppLogger(applogger.Config{})
	producer := kafkaClient.NewProducer(log, []string{broker1Address})
	// intialize the writer with the broker addresses, and the topic
	//w := kafka.NewWriter(kafka.WriterConfig{
	//	Brokers: []string{broker1Address},
	//	Topic:   kafka2.CreateUserTopic,
	//})
	user := &kafkaMessages.User{
		Username: "username",
		Password: "password",
		Email:    "email@example.com",
	}
	msg := &kafkaMessages.CreateUser{User: user}

	value, err := proto.Marshal(msg)
	if err != nil {
		panic("could not write message " + err.Error())
	}

	for {
		err := producer.PublishMessage(context.Background(), kafka.Message{
			Topic: kafka2.CreateUserTopic,
			Value: value,
			Time:  time.Now().UTC(),
		})
		// each kafka message has a key and value. The key is used
		// to decide which partition (and consequently, which broker)
		// the message gets published on
		//err := w.WriteMessages(ctx, kafka.Message{
		//	Key: []byte(strconv.Itoa(i)),
		//	// create an arbitrary message payload for the value
		//	Value: []byte("hi user " + strconv.Itoa(i)),
		//})
		if err != nil {
			panic("could not write message " + err.Error())
		}

		// log a confirmation once the message is written
		fmt.Println("writes:", i)
		i++
		// sleep for a second
		time.Sleep(time.Second)
	}
}
