package utils

import (
	"io"

	"github.com/opentracing/opentracing-go"

	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
)

func GetGlobalTrace(name string) (opentracing.Tracer, io.Closer, error) {

	cfg := jaegercfg.Configuration{
		ServiceName: name,
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans: true,
		},
	}

	return cfg.NewTracer()
}
