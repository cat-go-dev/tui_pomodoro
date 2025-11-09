package app

import (
	"context"
	"time"
)

type timer interface {
	Start(ctx context.Context, time time.Duration)
	Done(ctx context.Context) struct{}
}

type commands struct {
	timer timer
}

func newCommnads(timer timer) *commands {
	return &commands{
		timer: timer,
	}
}

const defaultDuration = time.Minute * 25

func (c commands) StartTimer(ctx context.Context, duration time.Duration) {
	if duration == 0 {
		duration = defaultDuration
	}

	c.timer.Start(ctx, duration)
}

func (c commands) ListenToTimer(ctx context.Context) {
	c.timer.Done(ctx)
}
