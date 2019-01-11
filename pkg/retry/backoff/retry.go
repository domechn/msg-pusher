package backoff

import "time"

// retry或retryNorify执行的事情
type Operation func() error

//
// retrynotify  的提醒回掉方法,不论成功失败都会执行
type Notify func(error, time.Duration)

//
// 重试有可能失败的操作
func Retry(o Operation, b BackOff) error { return RetryNotify(o, b, nil) }

// 带提醒操作的retry
func RetryNotify(operation Operation, b BackOff, notify Notify) error {
	var err error
	var next time.Duration

	cb := ensureContext(b)

	b.Reset()
	for {
		if err = operation(); err == nil {
			return nil
		}

		if permanent, ok := err.(*PermanentError); ok {
			return permanent.Err
		}

		if next = b.NextBackOff(); next == Stop {
			return err
		}

		if notify != nil {
			notify(err, next)
		}

		t := time.NewTimer(next)

		select {
		case <-cb.Context().Done():
			t.Stop()
			return err
		case <-t.C:
		}
	}
}

// PermanentError 错误信号
type PermanentError struct {
	Err error
}

func (e *PermanentError) Error() string {
	return e.Err.Error()
}

func Permanent(err error) *PermanentError {
	return &PermanentError{
		Err: err,
	}
}
