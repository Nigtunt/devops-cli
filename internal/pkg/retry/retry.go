package retry

import (
	"fmt"
	"time"
)

// Config 重试配置
type Config struct {
	MaxRetries   int           // 最大重试次数
	InitialDelay time.Duration // 初始延迟
	Multiplier   float64       // 延迟倍数
	MaxDelay     time.Duration // 最大延迟
}

// DefaultConfig 默认配置
var DefaultConfig = Config{
	MaxRetries:   3,
	InitialDelay: 100 * time.Millisecond,
	Multiplier:   2.0,
	MaxDelay:     2 * time.Second,
}

// Do 执行带重试的操作
func Do[T any](operation func() (T, error), cfg Config, onError func(int, error)) (T, error) {
	var lastErr error
	var result T

	delay := cfg.InitialDelay

	for attempt := 0; attempt <= cfg.MaxRetries; attempt++ {
		result, lastErr = operation()
		if lastErr == nil {
			return result, nil
		}

		// 调用错误回调
		if onError != nil {
			onError(attempt, lastErr)
		}

		// 最后一次不等待
		if attempt < cfg.MaxRetries {
			time.Sleep(delay)
			delay = time.Duration(float64(delay) * cfg.Multiplier)
			if delay > cfg.MaxDelay {
				delay = cfg.MaxDelay
			}
		}
	}

	return result, fmt.Errorf("重试 %d 次后失败：%w", cfg.MaxRetries, lastErr)
}
