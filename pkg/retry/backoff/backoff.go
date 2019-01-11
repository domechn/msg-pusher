// backoff 重试算法
//
// 使用retry，或retryNotify 调用可能失败的方法
//
//
// demo 在demo目录下
package backoff

import "time"

type BackOff interface {
	// NextBackOff 返回retry前等待的时间
	// 或返回 backoff.Stop 判断停止后是否要做一些事情.
	//
	// Example usage:
	//
	// 	duration := backoff.NextBackOff();
	// 	if (duration == backoff.Stop) {
	// 		// Do not retry operation.
	// 	} else {
	// 		// Sleep for duration and retry operation.
	// 	}
	//
	NextBackOff() time.Duration

	// 重置状态.
	Reset()
}

// 停止标识
const Stop time.Duration = -1

// ZeroBackOff backoff时间为0，不等待直接retry
type ZeroBackOff struct{}

func (b *ZeroBackOff) Reset() {}

func (b *ZeroBackOff) NextBackOff() time.Duration { return 0 }

// 停止retry
type StopBackOff struct{}

func (b *StopBackOff) Reset() {}

func (b *StopBackOff) NextBackOff() time.Duration { return Stop }

// ConstantBackOff 重试的时间间隔相同,会不断的重试，重试的时间间隔会随着调用NextBackOff() 而增长
type ConstantBackOff struct {
	Interval time.Duration
}

func (b *ConstantBackOff) Reset()                     {}
func (b *ConstantBackOff) NextBackOff() time.Duration { return b.Interval }

func NewConstantBackOff(d time.Duration) *ConstantBackOff {
	return &ConstantBackOff{Interval: d}
}
