package agent

import (
	"github.com/opentracing/opentracing-go"
	"github.com/rezaAmiri123/test-microservice/pkg/tracing"
)

func (a *Agent) setupTracing() error {
	if a.TracerConfig.Enable{
		tracer,closer,err := tracing.NewJaegerTracer(a.TracerConfig)
		if err != nil{
			return err
		}
		opentracing.SetGlobalTracer(tracer)
		a.closers=append(a.closers,closer)
	}
	return nil
}
