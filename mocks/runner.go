// Code generated by github.com/efritz/go-mockgen 0.1.0; DO NOT EDIT.
// This file was generated by robots at
// 2019-06-17T20:42:19-05:00
// using the command
// $ go-mockgen -f github.com/go-nacelle/process -i Runner -o runner.go

package mocks

import (
	config "github.com/go-nacelle/config"
	process "github.com/go-nacelle/process"
	"sync"
	"time"
)

// MockRunner is a mock implementation of the Runner interface (from the
// package github.com/go-nacelle/process) used for unit testing.
type MockRunner struct {
	// RunFunc is an instance of a mock function object controlling the
	// behavior of the method Run.
	RunFunc *RunnerRunFunc
	// ShutdownFunc is an instance of a mock function object controlling the
	// behavior of the method Shutdown.
	ShutdownFunc *RunnerShutdownFunc
}

// NewMockRunner creates a new mock of the Runner interface. All methods
// return zero values for all results, unless overwritten.
func NewMockRunner() *MockRunner {
	return &MockRunner{
		RunFunc: &RunnerRunFunc{
			defaultHook: func(config.Config) <-chan error {
				return nil
			},
		},
		ShutdownFunc: &RunnerShutdownFunc{
			defaultHook: func(time.Duration) error {
				return nil
			},
		},
	}
}

// NewMockRunnerFrom creates a new mock of the MockRunner interface. All
// methods delegate to the given implementation, unless overwritten.
func NewMockRunnerFrom(i process.Runner) *MockRunner {
	return &MockRunner{
		RunFunc: &RunnerRunFunc{
			defaultHook: i.Run,
		},
		ShutdownFunc: &RunnerShutdownFunc{
			defaultHook: i.Shutdown,
		},
	}
}

// RunnerRunFunc describes the behavior when the Run method of the parent
// MockRunner instance is invoked.
type RunnerRunFunc struct {
	defaultHook func(config.Config) <-chan error
	hooks       []func(config.Config) <-chan error
	history     []RunnerRunFuncCall
	mutex       sync.Mutex
}

// Run delegates to the next hook function in the queue and stores the
// parameter and result values of this invocation.
func (m *MockRunner) Run(v0 config.Config) <-chan error {
	r0 := m.RunFunc.nextHook()(v0)
	m.RunFunc.appendCall(RunnerRunFuncCall{v0, r0})
	return r0
}

// SetDefaultHook sets function that is called when the Run method of the
// parent MockRunner instance is invoked and the hook queue is empty.
func (f *RunnerRunFunc) SetDefaultHook(hook func(config.Config) <-chan error) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// Run method of the parent MockRunner instance invokes the hook at the
// front of the queue and discards it. After the queue is empty, the default
// hook function is invoked for any future action.
func (f *RunnerRunFunc) PushHook(hook func(config.Config) <-chan error) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultDefaultHook with a function that returns
// the given values.
func (f *RunnerRunFunc) SetDefaultReturn(r0 <-chan error) {
	f.SetDefaultHook(func(config.Config) <-chan error {
		return r0
	})
}

// PushReturn calls PushDefaultHook with a function that returns the given
// values.
func (f *RunnerRunFunc) PushReturn(r0 <-chan error) {
	f.PushHook(func(config.Config) <-chan error {
		return r0
	})
}

func (f *RunnerRunFunc) nextHook() func(config.Config) <-chan error {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *RunnerRunFunc) appendCall(r0 RunnerRunFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of RunnerRunFuncCall objects describing the
// invocations of this function.
func (f *RunnerRunFunc) History() []RunnerRunFuncCall {
	f.mutex.Lock()
	history := make([]RunnerRunFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// RunnerRunFuncCall is an object that describes an invocation of method Run
// on an instance of MockRunner.
type RunnerRunFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 config.Config
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 <-chan error
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c RunnerRunFuncCall) Args() []interface{} {
	return []interface{}{c.Arg0}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c RunnerRunFuncCall) Results() []interface{} {
	return []interface{}{c.Result0}
}

// RunnerShutdownFunc describes the behavior when the Shutdown method of the
// parent MockRunner instance is invoked.
type RunnerShutdownFunc struct {
	defaultHook func(time.Duration) error
	hooks       []func(time.Duration) error
	history     []RunnerShutdownFuncCall
	mutex       sync.Mutex
}

// Shutdown delegates to the next hook function in the queue and stores the
// parameter and result values of this invocation.
func (m *MockRunner) Shutdown(v0 time.Duration) error {
	r0 := m.ShutdownFunc.nextHook()(v0)
	m.ShutdownFunc.appendCall(RunnerShutdownFuncCall{v0, r0})
	return r0
}

// SetDefaultHook sets function that is called when the Shutdown method of
// the parent MockRunner instance is invoked and the hook queue is empty.
func (f *RunnerShutdownFunc) SetDefaultHook(hook func(time.Duration) error) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// Shutdown method of the parent MockRunner instance invokes the hook at the
// front of the queue and discards it. After the queue is empty, the default
// hook function is invoked for any future action.
func (f *RunnerShutdownFunc) PushHook(hook func(time.Duration) error) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultDefaultHook with a function that returns
// the given values.
func (f *RunnerShutdownFunc) SetDefaultReturn(r0 error) {
	f.SetDefaultHook(func(time.Duration) error {
		return r0
	})
}

// PushReturn calls PushDefaultHook with a function that returns the given
// values.
func (f *RunnerShutdownFunc) PushReturn(r0 error) {
	f.PushHook(func(time.Duration) error {
		return r0
	})
}

func (f *RunnerShutdownFunc) nextHook() func(time.Duration) error {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *RunnerShutdownFunc) appendCall(r0 RunnerShutdownFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of RunnerShutdownFuncCall objects describing
// the invocations of this function.
func (f *RunnerShutdownFunc) History() []RunnerShutdownFuncCall {
	f.mutex.Lock()
	history := make([]RunnerShutdownFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// RunnerShutdownFuncCall is an object that describes an invocation of
// method Shutdown on an instance of MockRunner.
type RunnerShutdownFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 time.Duration
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 error
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c RunnerShutdownFuncCall) Args() []interface{} {
	return []interface{}{c.Arg0}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c RunnerShutdownFuncCall) Results() []interface{} {
	return []interface{}{c.Result0}
}
