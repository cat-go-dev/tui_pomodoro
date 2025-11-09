package timer

import (
	"context"
	"time"
)

type timerImpl struct {
	ch chan struct{}
}

func NewTimer() *timerImpl {
	return &timerImpl{
		ch: make(chan struct{}),
	}
}

func (t *timerImpl) Start(ctx context.Context, dur time.Duration) {
	go func() {
		select {
		case <-time.After(dur):
			t.ch <- struct{}{}
			return
		case <-ctx.Done():
			return
		}
	}()
}

func (t *timerImpl) Done(ctx context.Context) struct{} {
	select {
	case v := <-t.ch:
		return v
	case <-ctx.Done():
		return struct{}{}
	}
}
