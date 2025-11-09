package timer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewTimer(t *testing.T) {
	t.Parallel()

	createdTimer := NewTimer()
	excpectedTimer := &timerImpl{}
	assert.Equal(t, excpectedTimer, createdTimer)
}

// todo: переписать тесты
// func Test_Start_Success(t *testing.T) {
// 	t.Parallel()
//
// 	ctx := context.Background()
// 	timer := NewTimer()
// 	startTime := time.Now()
// 	expectedDiff := time.Duration(startTime.Unix()) + time.Millisecond*500
// 	timerDur := 2 * time.Second
//
// 	tCh := timer.Start(ctx, timerDur)
// 	<-tCh
//
// 	finishTime := time.Now()
//
// 	assert.WithinDuration(t, startTime, finishTime, expectedDiff)
// }
//
// func Test_Start_ContextDone(t *testing.T) {
// 	t.Parallel()
//
// 	ctx, cancel := context.WithCancel(context.Background())
// 	timer := NewTimer()
// 	dur := 20 * time.Second
// 	cancelDur := 2 * time.Second
// 	startTime := time.Now()
// 	expectedDiff := cancelDur + time.Millisecond*500
//
// 	tCh := timer.Start(ctx, dur)
//
// 	go func() {
// 		time.Sleep(cancelDur)
// 		cancel()
// 	}()
//
// 	select {
// 	case <-tCh:
// 		assert.Fail(t, "error with context done")
// 	case <-ctx.Done():
// 		<-tCh
// 		finishTime := time.Now()
// 		assert.WithinDuration(t, startTime, finishTime, expectedDiff)
// 	}
// }
