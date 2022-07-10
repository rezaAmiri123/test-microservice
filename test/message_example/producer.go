package message_example

import (
	"context"
	"fmt"
	kafkaMessages "github.com/rezaAmiri123/test-microservice/library_service/proto/kafka"
	kafka2 "github.com/rezaAmiri123/test-microservice/pkg/kafka"
	kafkaClient "github.com/rezaAmiri123/test-microservice/pkg/kafka"
	"github.com/rezaAmiri123/test-microservice/pkg/logger/applogger"
	"github.com/segmentio/kafka-go"
	"google.golang.org/protobuf/proto"
	"time"
)

const (
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

	for {
		msg := &kafkaMessages.ArticleCreateRequest{
			UUID:        fmt.Sprintf("uuid %v", i),
			UserUUID:    fmt.Sprintf("user uuid %v", i),
			Description: fmt.Sprintf("description %v", i),
			Title:       fmt.Sprintf("title %v", i),

			Body: fmt.Sprintf("body %v", i),
		}

		value, err := proto.Marshal(msg)
		if err != nil {
			panic("could not write message " + err.Error())
		}

		err = producer.PublishMessage(context.Background(), kafka.Message{
			Topic: kafka2.CreateArticleTopic,
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
