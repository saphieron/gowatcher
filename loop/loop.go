package loop

import (
	"context"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/saphieron/gowatcher/executor"
	"github.com/saphieron/gowatcher/liberrors"
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
	looper.wg = new(sync.WaitGroup)
	executorChan := make(chan executor.CommandExecutor, 1)
	errorChan := make(chan error, 10)
	resultErrorChan := make(chan error, 1)

	looper.wg.Add(3)
	go looper.schedule(command, commandArgs, executorChan)
	go looper.process(executorChan, errorChan)
	go looper.errorCollector(errorChan, resultErrorChan)
	looper.wg.Wait()
	select {
	case err := <-resultErrorChan:
		if err != nil {
			logging.ErrorLog.Printf("loop execution encountered errors: %s", err.Error())
		}
		return err
	default:
	}
	return nil
}

func (looper *Looper) schedule(command string, commandArgs []string, outChan chan<- executor.CommandExecutor) {
	defer looper.wg.Done()
	running := true
	for running {
		exec := executor.NewCommandExecutorWithContext(looper.ctx, command, commandArgs...)
		outChan <- exec
		after := time.After(looper.interval)
		select {
		case <-looper.ctx.Done():
			running = false
		case <-after:
		}
	}
}

func (looper *Looper) process(inChan <-chan executor.CommandExecutor, errorChan chan<- error) {
	defer looper.wg.Done()
	running := true
	var exec executor.CommandExecutor
	for running {
		select {
		case exec = <-inChan:
		case <-looper.ctx.Done():
			running = false
			continue
		}
		err := printOutput(exec)
		if err != nil {
			wrapErrorAndReportIt("looped processor encountered error with command execution: %w", err, errorChan)
			running = false
			continue
		}
	}
}

func (looper *Looper) errorCollector(errorChan <-chan error, resultErrChan chan<- error) {
	defer looper.wg.Done()
	running := true
	encounteredError := false
	collection := liberrors.NewCollectedError("loop execution")
	for running {
		select {
		case <-looper.ctx.Done():
			running = false
		case loopErr := <-errorChan:
			collection.AppendError(loopErr)
			encounteredError = true
		}
	}
	if encounteredError {
		resultErrChan <- collection
	}
	close(resultErrChan)
}

// // we assume that the ctx used in the CommandExecutor listens to the same cancel as this one here.
// func (looper *Looper) executeLoop(command string, commandArgs []string, errorChan chan<- error) {
// 	defer looper.wg.Done()
// 	running := true
// 	var err error
// 	for running {
// 		exec := executor.NewCommandExecutorWithContext(looper.ctx, command, commandArgs...)
// 		startingTime := time.Now()
// 		select {
// 		case <-looper.ctx.Done():
// 			running = false
// 		default:
// 			err = printOutput(exec)
// 			if err != nil {
// 				wrapErrorAndReportIt("looped failed to extract command's output: %w", err, errorChan)
// 				running = false
// 				continue
// 			}
// 		}

// 		looper.waitIfNecessary(startingTime)
// 	}
// 	close(errorChan)
// }

func wrapErrorAndReportIt(wrapText string, err error, c chan<- error) {
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
		return fmt.Errorf("rendering to console failed: %w", err)
	}
	return nil
}
