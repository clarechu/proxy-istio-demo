package tracing

import (
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	"io"
	"time"
)

func NewTracing(samplingServerURL, localAgentHostPort string) (opentracing.Tracer, io.Closer) {
	cfg := config.Configuration{
		ServiceName: "grpc-server",
		Sampler: &config.SamplerConfig{
			Type:              "const",
			Param:             1,
			SamplingServerURL: samplingServerURL,
		},
		Reporter: &config.ReporterConfig{
			LogSpans:            true,
			BufferFlushInterval: 1 * time.Second,
			LocalAgentHostPort:  localAgentHostPort, // 替换host
		},
	}
	tracer, closer, _ := cfg.NewTracer(
		config.Logger(jaeger.StdLogger),
	)
	opentracing.SetGlobalTracer(tracer)
	return tracer, closer
}
