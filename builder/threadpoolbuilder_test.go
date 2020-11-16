package builder

import "testing"

func TestThreadPoolBuilder(t *testing.T) {
	b := &ThreadPoolBuilder{}
	tp := b.SetName("threadpool").SetMaxTotal(10).SetMaxIdle(6).SetMinIdle(4).Build()
	tp.Status()
}
