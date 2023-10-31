package lamportclock

import "sync/atomic"

type LamportClock struct {
	latestTS int32
}

func New(ts int32) *LamportClock {
	lc := &LamportClock{}
	atomic.StoreInt32(&lc.latestTS, ts)
	return lc
}

func (lc *LamportClock) Tick(requestTime int32) int32 {
	ts := atomic.LoadInt32(&lc.latestTS)
	atomic.StoreInt32(&lc.latestTS, max(&ts, lc.latestTS, requestTime))
	atomic.StoreInt32(&lc.latestTS, atomic.AddInt32(&lc.latestTS, 1))
	return lc.latestTS
}

func (lc *LamportClock) GetLatestTime() int32 {
	return atomic.LoadInt32(&lc.latestTS)
}

func (lc *LamportClock) UpdateTo(at int32) {
	atomic.StoreInt32(&lc.latestTS, at)
}

func max(ts *int32, x, y int32) int32 {
	if !atomic.CompareAndSwapInt32(ts, x, y) {
		return x
	}
	return y
}
