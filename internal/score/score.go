package score

import "sync/atomic"

var s int64

func Inc() {
	atomic.AddInt64(&s, 1)
}

func Get() int64 {
	return atomic.LoadInt64(&s)
}

func Reset() {
	atomic.StoreInt64(&s, 0)
}
