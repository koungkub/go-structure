package utils

import (
	"context"
	"io"

	"github.com/labstack/echo/v4"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/sirupsen/logrus"

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

func ExtractSpanAndGetLog(name string, c echo.Context) (opentracing.Span, *logrus.Entry) {

	log := c.Get("log").(*logrus.Entry)

	spanCtx, _ := opentracing.GlobalTracer().Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(c.Request().Header))
	span := opentracing.GlobalTracer().StartSpan(name, ext.RPCServerOption(spanCtx))

	if span.BaggageItem("request.id") == "" {
		span.SetBaggageItem("request.id", c.Response().Header().Get(echo.HeaderXRequestID))
	}

	span.SetTag("request.id", span.BaggageItem("request.id"))
	log.WithFields(logrus.Fields{
		"requestID": span.BaggageItem("request.id"),
	})

	return span, log
}

func GetContextWithSpan(span opentracing.Span) (context.Context, context.CancelFunc) {

	context, cancel := GetContext()
	ctx := opentracing.ContextWithSpan(context, span)

	return ctx, cancel
}
