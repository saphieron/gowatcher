package loop_test

import (
	"context"
	"testing"
	"time"

	"github.com/saphieron/gowatcher/loop"
)

func TestLoopFake(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	looper := loop.NewLooper(time.Duration(float64(time.Second)*0.1), ctx)
	go func() {
		time.Sleep(time.Second)
		cancel()
	}()
	looper.Start("ls", "-la")

}
