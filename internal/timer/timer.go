package timer

import (
	"context"
	"time"
)

type Timer interface {
	Start(ctx context.Context, time time.Duration) <-chan struct{}
}

type timerImpl struct{}

func NewTimer() *timerImpl {
	return &timerImpl{}
}

func (t *timerImpl) Start(ctx context.Context, dur time.Duration) <-chan struct{} {
	signalCh := make(chan struct{})

	go func() {
		defer close(signalCh)

		select {
		case <-time.After(dur):
			signalCh <- struct{}{}
			return
		case <-ctx.Done():
			return
		}
	}()

	return signalCh
}
