package mocks

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

//go:generate minimock -i pb-dropbox-downloader/infrastructure.Dropbox -o ./mocks\dropboxmock.go

import (
	mm_infrastructure "pb-dropbox-downloader/infrastructure"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
)

// DropboxMock implements infrastructure.Dropbox
type DropboxMock struct {
	t minimock.Tester

	funcGetFiles          func() (ra1 []mm_infrastructure.RemoteFile)
	inspectFuncGetFiles   func()
	afterGetFilesCounter  uint64
	beforeGetFilesCounter uint64
	GetFilesMock          mDropboxMockGetFiles
}

// NewDropboxMock returns a mock for infrastructure.Dropbox
func NewDropboxMock(t minimock.Tester) *DropboxMock {
	m := &DropboxMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.GetFilesMock = mDropboxMockGetFiles{mock: m}

	return m
}

type mDropboxMockGetFiles struct {
	mock               *DropboxMock
	defaultExpectation *DropboxMockGetFilesExpectation
	expectations       []*DropboxMockGetFilesExpectation
}

// DropboxMockGetFilesExpectation specifies expectation struct of the Dropbox.GetFiles
type DropboxMockGetFilesExpectation struct {
	mock *DropboxMock

	results *DropboxMockGetFilesResults
	Counter uint64
}

// DropboxMockGetFilesResults contains results of the Dropbox.GetFiles
type DropboxMockGetFilesResults struct {
	ra1 []mm_infrastructure.RemoteFile
}

// Expect sets up expected params for Dropbox.GetFiles
func (mmGetFiles *mDropboxMockGetFiles) Expect() *mDropboxMockGetFiles {
	if mmGetFiles.mock.funcGetFiles != nil {
		mmGetFiles.mock.t.Fatalf("DropboxMock.GetFiles mock is already set by Set")
	}

	if mmGetFiles.defaultExpectation == nil {
		mmGetFiles.defaultExpectation = &DropboxMockGetFilesExpectation{}
	}

	return mmGetFiles
}

// Inspect accepts an inspector function that has same arguments as the Dropbox.GetFiles
func (mmGetFiles *mDropboxMockGetFiles) Inspect(f func()) *mDropboxMockGetFiles {
	if mmGetFiles.mock.inspectFuncGetFiles != nil {
		mmGetFiles.mock.t.Fatalf("Inspect function is already set for DropboxMock.GetFiles")
	}

	mmGetFiles.mock.inspectFuncGetFiles = f

	return mmGetFiles
}

// Return sets up results that will be returned by Dropbox.GetFiles
func (mmGetFiles *mDropboxMockGetFiles) Return(ra1 []mm_infrastructure.RemoteFile) *DropboxMock {
	if mmGetFiles.mock.funcGetFiles != nil {
		mmGetFiles.mock.t.Fatalf("DropboxMock.GetFiles mock is already set by Set")
	}

	if mmGetFiles.defaultExpectation == nil {
		mmGetFiles.defaultExpectation = &DropboxMockGetFilesExpectation{mock: mmGetFiles.mock}
	}
	mmGetFiles.defaultExpectation.results = &DropboxMockGetFilesResults{ra1}
	return mmGetFiles.mock
}

//Set uses given function f to mock the Dropbox.GetFiles method
func (mmGetFiles *mDropboxMockGetFiles) Set(f func() (ra1 []mm_infrastructure.RemoteFile)) *DropboxMock {
	if mmGetFiles.defaultExpectation != nil {
		mmGetFiles.mock.t.Fatalf("Default expectation is already set for the Dropbox.GetFiles method")
	}

	if len(mmGetFiles.expectations) > 0 {
		mmGetFiles.mock.t.Fatalf("Some expectations are already set for the Dropbox.GetFiles method")
	}

	mmGetFiles.mock.funcGetFiles = f
	return mmGetFiles.mock
}

// GetFiles implements infrastructure.Dropbox
func (mmGetFiles *DropboxMock) GetFiles() (ra1 []mm_infrastructure.RemoteFile) {
	mm_atomic.AddUint64(&mmGetFiles.beforeGetFilesCounter, 1)
	defer mm_atomic.AddUint64(&mmGetFiles.afterGetFilesCounter, 1)

	if mmGetFiles.inspectFuncGetFiles != nil {
		mmGetFiles.inspectFuncGetFiles()
	}

	if mmGetFiles.GetFilesMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmGetFiles.GetFilesMock.defaultExpectation.Counter, 1)

		mm_results := mmGetFiles.GetFilesMock.defaultExpectation.results
		if mm_results == nil {
			mmGetFiles.t.Fatal("No results are set for the DropboxMock.GetFiles")
		}
		return (*mm_results).ra1
	}
	if mmGetFiles.funcGetFiles != nil {
		return mmGetFiles.funcGetFiles()
	}
	mmGetFiles.t.Fatalf("Unexpected call to DropboxMock.GetFiles.")
	return
}

// GetFilesAfterCounter returns a count of finished DropboxMock.GetFiles invocations
func (mmGetFiles *DropboxMock) GetFilesAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetFiles.afterGetFilesCounter)
}

// GetFilesBeforeCounter returns a count of DropboxMock.GetFiles invocations
func (mmGetFiles *DropboxMock) GetFilesBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetFiles.beforeGetFilesCounter)
}

// MinimockGetFilesDone returns true if the count of the GetFiles invocations corresponds
// the number of defined expectations
func (m *DropboxMock) MinimockGetFilesDone() bool {
	for _, e := range m.GetFilesMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetFilesMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetFilesCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetFiles != nil && mm_atomic.LoadUint64(&m.afterGetFilesCounter) < 1 {
		return false
	}
	return true
}

// MinimockGetFilesInspect logs each unmet expectation
func (m *DropboxMock) MinimockGetFilesInspect() {
	for _, e := range m.GetFilesMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Error("Expected call to DropboxMock.GetFiles")
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetFilesMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetFilesCounter) < 1 {
		m.t.Error("Expected call to DropboxMock.GetFiles")
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetFiles != nil && mm_atomic.LoadUint64(&m.afterGetFilesCounter) < 1 {
		m.t.Error("Expected call to DropboxMock.GetFiles")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *DropboxMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockGetFilesInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *DropboxMock) MinimockWait(timeout mm_time.Duration) {
	timeoutCh := mm_time.After(timeout)
	for {
		if m.minimockDone() {
			return
		}
		select {
		case <-timeoutCh:
			m.MinimockFinish()
			return
		case <-mm_time.After(10 * mm_time.Millisecond):
		}
	}
}

func (m *DropboxMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockGetFilesDone()
}
