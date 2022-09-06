package tracing

import (
	"io"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
)

type Config struct {
	ServiceName string `mapstructure:"TRACER_SERVICE_NAME"`
	HostPort    string `mapstructure:"TRACER_HOST_PORT"`
	Enable      bool   `mapstructure:"TRACER_ENABLE"`
	LogSpans    bool   `mapstructure:"TRACER_LOG_SPANS"`
}

func NewJaegerTracer(jaegerConfig Config) (opentracing.Tracer, io.Closer, error) {
	cfg := config.Configuration{
		ServiceName: jaegerConfig.ServiceName,
		Sampler: &config.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans:           jaegerConfig.LogSpans,
			LocalAgentHostPort: jaegerConfig.HostPort,
		},
	}

	// Example logger and metrics factory. Use github.com/uber/jaeger-client-go/log
	// and github.com/uber/jaeger-lib/metrics respectively to bind to real logging and metrics
	// frameworks.
	//jLogger := jaegerlog.StdLogger
	//jMetricsFactory := metrics.NullFactory

	// Initialize tracer with a logger and a metrics factory
	//tracer, closer, err:= cfg.NewTracer(
	//	jaegercfg.Logger(jLogger),
	//	jaegercfg.Metrics(jMetricsFactory),
	//)
	// Set the singleton opentracing.Tracer with the Jaeger tracer.
	//opentracing.SetGlobalTracer(tracer)
	//defer closer.Close()
	return cfg.NewTracer(config.Logger(jaeger.StdLogger))

}

/*
	cfg := jaegercfg.Configuration{
		ServiceName: "my_service_name",
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans: true,
		},
	}

	// Example logger and metrics factory. Use github.com/uber/jaeger-client-go/log
	// and github.com/uber/jaeger-lib/metrics respectively to bind to real logging and metrics
	// frameworks.
	jLogger := jaegerlog.StdLogger
	jMetricsFactory := metrics.NullFactory

	// Initialize tracer with a logger and a metrics factory
	return  cfg.NewTracer(
		jaegercfg.Logger(jLogger),
		jaegercfg.Metrics(jMetricsFactory),
	)
}
*/
