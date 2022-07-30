package grpc_example

import (
	"context"
	"fmt"
	messageapi "github.com/rezaAmiri123/test-microservice/message_service/proto/grpc"
	"google.golang.org/grpc"
	"time"
)

const (
	grpcAddress = "message_service:8481"
)

func Client(ctx context.Context) {
	// initialize a counter
	i := 0
	//log := applogger.NewAppLogger(applogger.Config{})
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.DialContext(ctx, grpcAddress, opts...)
	if err != nil {
		return
	}
	messageClient := messageapi.NewMessageServiceClient(conn)
	// intialize the writer with the broker addresses, and the topic
	//w := kafka.NewWriter(kafka.WriterConfig{
	//	Brokers: []string{broker1Address},
	//	Topic:   kafka2.CreateUserTopic,
	//})

	for {
		msg := &messageapi.Email{
			From:    fmt.Sprintf("%d@example.com", i),
			To:      []string{fmt.Sprintf("%d@example.com", i), fmt.Sprintf("%d@example.com", i)},
			Body:    fmt.Sprintf("body %d", i),
			Subject: fmt.Sprintf("subject %d", i),
		}

		value, err := messageClient.CreateEmail(ctx, &messageapi.CreateEmailRequest{
			Email: msg,
		})
		if err != nil {
			panic("could not write message " + err.Error())
		}
		fmt.Println(value.GetUUID())

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
