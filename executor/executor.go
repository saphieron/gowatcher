package executor

import (
	"io"
)

type CommandExecutor interface {
	CombinedOutput() ([]byte, error)
	Environ() []string
	Output() ([]byte, error)
	Run() error
	Start() error
	StderrPipe() (io.ReadCloser, error)
	StdinPipe() (io.WriteCloser, error)
	StdoutPipe() (io.ReadCloser, error)
	String() string
	Wait() error
}

// type CommandReporter struct {
// 	CommandExecutor
// }

// func (e *CommandReporter) CombinedOutput() ([]byte, error) {
// 	exec.CombinedOutput()
// 	return nil, nil
// }

// func (exec *CommandReporter) Environ() []string {
// 	return nil
// }

// func (exec *CommandReporter) Output() ([]byte, error) {
// 	return nil, nil
// }

// func (exec *CommandReporter) Run() error {
// 	return nil
// }

// func (exec *CommandReporter) Start() error {
// 	return nil
// }

// func (exec *CommandReporter) StderrPipe() (io.ReadCloser, error) {
// 	return nil, nil
// }

// func (exec *CommandReporter) StdinPipe() (io.WriteCloser, error) {
// 	return nil, nil
// }

// func (exec *CommandReporter) StdoutPipe() (io.ReadCloser, error) {
// 	return nil, nil
// }

// func (exec *CommandReporter) String() string {
// 	return ""
// }

// func (exec *CommandReporter) Wait() error {
// 	return nil
// }
