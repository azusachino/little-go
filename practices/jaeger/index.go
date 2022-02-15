package main

import (
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	"io"
	"log"
	"time"
)

func main() {

	closer := initTracer()
	defer closer.Close()

	(func() {
		span := opentracing.GlobalTracer().StartSpan("Root Span")
		defer span.Finish()

		(func() {
			span := span.Tracer().StartSpan("Child Span 1", opentracing.ChildOf(span.Context()))
			defer span.Finish()
			time.Sleep(1 * time.Second)
		})()
		(func() {
			span := span.Tracer().StartSpan("Child Span 2", opentracing.ChildOf(span.Context()))
			defer span.Finish()
			time.Sleep(1 * time.Second)
		})()
		(func() {
			span := span.Tracer().StartSpan("Child Span 3", opentracing.ChildOf(span.Context()))
			defer span.Finish()
			time.Sleep(1 * time.Second)
		})()
	})()

}

func initTracer() io.Closer {
	cfg := jaegercfg.Configuration{
		ServiceName: "my demo service",
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{},
	}
	tracer, closer, err := cfg.NewTracer()
	if err != nil {
		log.Fatal(err)
	}

	opentracing.SetGlobalTracer(tracer)

	return closer
}
