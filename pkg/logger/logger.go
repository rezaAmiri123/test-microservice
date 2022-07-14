package logger

import "time"

const (
	Topic     = "topic"
	Partition = "partition"
	Message   = "message"
	WorkerID  = "workerID"
	Offset    = "offset"
	Time      = "time"

	GRPC     = "GRPC"
	METHOD   = "METHOD"
	METADATA = "METADATA"
	REQUEST  = "REQUEST"
	REPLY    = "REPLY"
	TIME     = "TIME"
)

// Logger methods interface
type Logger interface {
	InitLogger()
	Sync() error
	Debug(args ...interface{})
	Debugf(template string, args ...interface{})
	Info(args ...interface{})
	Infof(template string, args ...interface{})
	Warn(args ...interface{})
	Warnf(template string, args ...interface{})
	WarnMsg(msg string, err error)
	Error(args ...interface{})
	Errorf(template string, args ...interface{})
	Err(msg string, err error)
	DPanic(args ...interface{})
	DPanicf(template string, args ...interface{})
	Fatal(args ...interface{})
	Fatalf(template string, args ...interface{})
	Printf(template string, args ...interface{})
	WithName(name string)
	KafkaProcessMessage(topic string, partition int, message string, workerID int, offset int64, time time.Time)
	KafkaLogCommittedMessage(topic string, partition int, offset int64)
	GrpcMiddlewareAccessLogger(method string, time time.Duration, metaData map[string][]string, err error)
	GrpcClientInterceptorLogger(method string, req interface{}, reply interface{}, time time.Duration, metaData map[string][]string, err error)
}
