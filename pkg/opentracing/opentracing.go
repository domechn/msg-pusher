// Copyright (c) 2018, dmc (814172254@qq.com),
//
// Authors: dmc,
//
// Distribution:.
package opentracing

import (
	"context"
	"io"
	"net"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	"github.com/uber/jaeger-client-go/zipkin"
	"github.com/uber/jaeger-lib/metrics"
)

const (
	jaegerTracing     = "jaeger"
	zipkinPropagation = "zipkin"
)

// Config 配置文件
type Config struct {
	Provider      string
	JaegerTracing JaegerTracing
}

func InitConfig(url string) Config {
	_, _, err := net.SplitHostPort(url)
	if url == "" || err != nil {
		return initNilConfig()
	}
	return initJaegerConfig(url)
}

// initNilConfig 返回一个空的配置
func initNilConfig() Config {
	return Config{}
}

// initJaegerConfig 返回一个jaeger的配置
func initJaegerConfig(serverURL string) Config {
	return Config{
		Provider: jaegerTracing,
		JaegerTracing: JaegerTracing{
			ServiceName:         "gateway",
			SamplingServerURL:   serverURL,
			SamplingParam:       1.0,
			SamplingType:        "const",
			BufferFlushInterval: time.Second,
			LogSpans:            false,
			QueueSize:           1000,
		},
	}
}

// Tracing tracing的基类
type Tracing struct {
	config Config
	tracer opentracing.Tracer
	closer io.Closer
}

type noopCloser struct{}

func (n noopCloser) Close() error { return nil }

// New 按照配置信息初始化tracing
func New(cfg Config) *Tracing {
	return &Tracing{
		config: cfg,
	}
}

// Setup 根据tracing的属性选择合适的opentracing客户端
func (t *Tracing) Setup() (err error) {
	pro := t.config.Provider
	log.Debug("Initializing distributed tracing")

	switch pro {
	case jaegerTracing:
		log.Debug("Select jaeger as tracing system")
		t.tracer, t.closer, err = t.buildJaeger(t.config.JaegerTracing)
	default:
		log.Debug("No tracer selected")
		t.tracer, t.closer, err = opentracing.NoopTracer{}, noopCloser{}, nil
	}
	// 设置全局的tracer
	opentracing.SetGlobalTracer(t.tracer)
	return
}

// Close 关闭tracer
func (t *Tracing) Close() {
	if t.closer != nil {
		t.closer.Close()
	}
}

// buildJaeger 创建jaeger system
func (t *Tracing) buildJaeger(cfg JaegerTracing) (opentracing.Tracer, io.Closer, error) {
	svrName := cfg.ServiceName
	conf := jaegercfg.Configuration{
		ServiceName: svrName,
		Sampler: &jaegercfg.SamplerConfig{
			Param: cfg.SamplingParam,
			Type:  cfg.SamplingType,
		},
		Reporter: &jaegercfg.ReporterConfig{
			QueueSize:           cfg.QueueSize,
			LocalAgentHostPort:  cfg.SamplingServerURL,
			BufferFlushInterval: cfg.BufferFlushInterval,
			LogSpans:            cfg.LogSpans,
		},
	}
	tracerMetrics := jaeger.NewMetrics(metrics.NullFactory, nil)
	sampler, err := conf.Sampler.NewSampler(svrName, tracerMetrics)
	tracerLogger := &jagerLogger{}
	if err != nil {
		return nil, nil, err
	}

	reporter, err := conf.Reporter.NewReporter(svrName, tracerMetrics, tracerLogger)
	if err != nil {
		return nil, nil, err
	}

	var (
		tracer opentracing.Tracer
		closer io.Closer
	)

	switch cfg.PropagationFormat {
	case zipkinPropagation:
		log.Debug("Using zipkin b3 http propagation format")
		zipkinPropagator := zipkin.NewZipkinB3HTTPHeaderPropagator()
		tracer, closer = jaeger.NewTracer(svrName, sampler, reporter,
			jaeger.TracerOptions.Metrics(tracerMetrics),
			jaeger.TracerOptions.Logger(tracerLogger),
			jaeger.TracerOptions.Injector(opentracing.HTTPHeaders, zipkinPropagator),
			jaeger.TracerOptions.Extractor(opentracing.HTTPHeaders, zipkinPropagator),
			jaeger.TracerOptions.ZipkinSharedRPCSpan(true),
		)
	default:
		log.Debug("Using jaeger propagation format")
		tracer, closer = jaeger.NewTracer(svrName, sampler, reporter,
			jaeger.TracerOptions.Metrics(tracerMetrics),
			jaeger.TracerOptions.Logger(tracerLogger),
		)
	}

	return tracer, closer, nil
}

// ToContext 将span写入request的context中
func ToContext(r *http.Request, span opentracing.Span) *http.Request {
	return r.WithContext(opentracing.ContextWithSpan(r.Context(), span))
}

// FromContext 从context中取出span
func FromContext(ctx context.Context, name string) opentracing.Span {
	span, _ := opentracing.StartSpanFromContext(ctx, name)
	return span
}
