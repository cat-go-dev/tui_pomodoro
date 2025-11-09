package timer

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_Start_Success(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	timer := NewTimer()
	startTime := time.Now()
	expectedDiff := time.Duration(startTime.Unix()) + time.Millisecond*500
	timerDur := 2 * time.Second

	timer.Start(ctx, timerDur)
	timer.Done(ctx)

	finishTime := time.Now()

	assert.WithinDuration(t, startTime, finishTime, expectedDiff)
}

func Test_Start_ContextDone(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	timer := NewTimer()
	dur := 20 * time.Second
	cancelDur := 2 * time.Second
	startTime := time.Now()
	expectedDiff := cancelDur + time.Millisecond*500

	timer.Start(ctx, dur)

	go func() {
		time.Sleep(cancelDur)
		cancel()
	}()

	doneCh := make(chan struct{})
	go func() {
		timer.Done(ctx)
		doneCh <- struct{}{}
	}()

	select {
	case <-doneCh:
		assert.Fail(t, "error with context done")
	case <-ctx.Done():
		<-doneCh
		finishTime := time.Now()
		assert.WithinDuration(t, startTime, finishTime, expectedDiff)
	}
}
