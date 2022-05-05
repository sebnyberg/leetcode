package p0359loggerratelimiter

import "math"

type Logger struct {
	msgTs map[string]int
}

func Constructor() Logger {
	return Logger{msgTs: make(map[string]int)}
}

func (this *Logger) ShouldPrintMessage(timestamp int, message string) bool {
	if _, exists := this.msgTs[message]; !exists {
		this.msgTs[message] = math.MinInt32
	}
	if timestamp-this.msgTs[message] >= 10 {
		this.msgTs[message] = timestamp
		return true
	}
	return false
}
