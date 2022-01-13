package main

import (
	"context"
	"github.com/rezaAmiri123/test-microservice/test/kafka1"
	"time"
)

func main() {
	// create a new context
	ctx := context.Background()
	// produce messages in a new go routine, since
	// both the produce and consume functions are
	// blocking
	kafka1.Produce(ctx)
	time.Sleep(time.Hour)
	//kafka1.Consume(ctx)
}