// Copyright (c) 2018, dmc (814172254@qq.com),
//
// Authors: dmc,
//
// Distribution:.
package opentracing

import (
	"github.com/sirupsen/logrus"
	"time"
)

type JaegerTracing struct {
	ServiceName string

	SamplingServerURL   string
	SamplingParam       float64
	SamplingType        string
	BufferFlushInterval time.Duration
	LogSpans            bool
	QueueSize           int
	PropagationFormat   string
}

type jagerLogger struct {
}

func (s *jagerLogger) Error(msg string) {
	logrus.Error(msg)
}

func (s *jagerLogger) Infof(msg string, args ...interface{}) {
	logrus.Infof(msg, args...)
}
