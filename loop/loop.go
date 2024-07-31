package loop

import (
	"context"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/saphieron/gowatcher/executor"
	"github.com/saphieron/gowatcher/logging"
)

type Looper struct {
	ctx      context.Context
	interval time.Duration
	wg       *sync.WaitGroup
}

func NewLooper(interval time.Duration, ctx context.Context) *Looper {
	looper := &Looper{
		ctx:      ctx,
		interval: interval,
	}
	return looper
}

func (looper *Looper) Do(command string, commandArgs ...string) error {
	defer logging.ErrorLog.Printf("finished loop")
	looper.wg = new(sync.WaitGroup)
	looper.wg.Add(1)
	errorChan := make(chan error)

	go looper.executeLoop(command, commandArgs, errorChan)

	for err := range errorChan {
		logging.ErrorLog.Printf("loop encountered error: %v", err)
	}
	looper.wg.Wait()
	logging.ErrorLog.Printf("finished waiting")
	// select {
	// case err := <-errorChan:
	// 	if err != nil {
	// 		return err
	// 	}
	// default:
	// 	// nothing to do
	// }
	return nil
}

// we assume that the ctx used in the CommandExecutor listens to the same cancel as this one here.
func (looper *Looper) executeLoop(command string, commandArgs []string, errorChan chan error) {
	defer looper.wg.Done()
	running := true
	var err error
	for running {
		exec := executor.NewCommandExecutorWithContext(looper.ctx, command, commandArgs...)
		startingTime := time.Now()
		select {
		case <-looper.ctx.Done():
			running = false
		default:
			err = printOutput(exec)
			if err != nil {
				wrapErrorAndReportIt("looped failed to extract command's output: %w", err, errorChan)
				running = false
				continue
			}
		}

		looper.waitIfNecessary(startingTime)
	}
	close(errorChan)
}

func wrapErrorAndReportIt(wrapText string, err error, c chan error) {
	wrappedErr := fmt.Errorf(wrapText, err)
	c <- wrappedErr
}

func printOutput(exec executor.CommandExecutor) error {
	output, err := exec.CombinedOutput()
	if err != nil {
		return err
	}
	output = append(output, []byte("\n")...)
	_, err = os.Stdout.Write(output)
	if err != nil {
		return fmt.Errorf("writing to console failed: %w", err)
	}
	return nil
}

func (looper *Looper) waitIfNecessary(startingTime time.Time) {
	execTime := time.Since(startingTime)
	if execTime < looper.interval {
		remainingTime := looper.interval - execTime
		after := time.After(remainingTime)
		done := looper.ctx.Done()
		select {
		case <-done:
		case <-after:
		}
	}
}
