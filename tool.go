package helper

import "time"

// 重试机制
type Fun func() error

// f 操作
// times 尝试次数
// sleepTime 每次尝试之前休眠的事件
func Retry(f Fun, times int, sleepTime time.Duration) (err error) {

	for i := 0; i < times; i++ {
		if err = f(); err != nil {
			time.Sleep(sleepTime)
			continue
		} else {
			break
		}
	}

	return
}
