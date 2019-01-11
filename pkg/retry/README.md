# 调用失败后重试

## 示例
```go
func main() {
	const successOn = 3
	var i = 0

	// This function is successful on "successOn" calls.
	f := func() error {
		i++
		log.Printf("function is called %d. time\n", i)

		if i == successOn {
			log.Println("OK")
			return nil
		}

		log.Println("error")
		return errors.New("error")
	}
	bkoff := backoff.NewExponentialBackOff()
	err := backoff.Retry(f, bkoff)
	if err != nil {
		log.Println("unexpected error: %s", err.Error())
	}
	if i != successOn {
		log.Println("invalid number of retries: %d", i)
	}
	//type nf backoff.Notify
	//i = 0
	bkoff.MaxElapsedTime = 40 * time.Second
	nf := func(e error, duration time.Duration) {
		if e != nil {
			log.Println("notify%s", e)
		} else {

		}
		log.Println(bkoff.GetElapsedTime())
	}
	backoff.RetryNotify(f, bkoff, nf)
}
```

## 关键参数
```
    InitialInterval     time.Duration    重试时间间隔基数
   	RandomizationFactor float64          随机系数
   	Multiplier          float64
   	MaxInterval         time.Duration     最大重试时间间隔
   	// 最大等待时间  超过后放弃重试  MaxElapsedTime == 0 永不停止
   	MaxElapsedTime time.Duration

   	默认值：
   	    DefaultInitialInterval     = 500 * time.Millisecond
    	DefaultRandomizationFactor = 0.5
    	DefaultMultiplier          = 1.5
    	DefaultMaxInterval         = 60 * time.Second
    	DefaultMaxElapsedTime      = 15 * time.Minute

```

## 说明
```
 实际的等待时间 =
     RetryInterval * (random value in range [1 - RandomizationFactor, 1 + RandomizationFactor])

Request #  RetryInterval (seconds)  Randomized Interval (seconds)

  1          0.5                     [0.25,   0.75]
  2          0.75                    [0.375,  1.125]
  3          1.125                   [0.562,  1.687]
  4          1.687                   [0.8435, 2.53]
  5          2.53                    [1.265,  3.795]
  6          3.795                   [1.897,  5.692]
  7          5.692                   [2.846,  8.538]
  8          8.538                   [4.269, 12.807]
  9         12.807                   [6.403, 19.210]
 10         19.210                   backoff.Stop
 ```

 > 线程不安全