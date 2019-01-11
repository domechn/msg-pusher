package backoff

import (
	"context"
	"time"
)

// 在取消context后停止retry
type BackOffContext interface {
	BackOff
	Context() context.Context
}

type backOffContext struct {
	BackOff
	ctx context.Context
}

// WithContext 返回一个带context的 BackOffContext
//
// ctx 不能为 nil
func WithContext(b BackOff, ctx context.Context) BackOffContext {
	if ctx == nil {
		panic("nil context")
	}

	if b, ok := b.(*backOffContext); ok {
		return &backOffContext{
			BackOff: b.BackOff,
			ctx:     ctx,
		}
	}

	return &backOffContext{
		BackOff: b,
		ctx:     ctx,
	}
}

func ensureContext(b BackOff) BackOffContext {
	if cb, ok := b.(BackOffContext); ok {
		return cb
	}
	return WithContext(b, context.Background())
}

func (b *backOffContext) Context() context.Context {
	return b.ctx
}

func (b *backOffContext) NextBackOff() time.Duration {
	select {
	case <-b.Context().Done():
		return Stop
	default:
		return b.BackOff.NextBackOff()
	}
}
