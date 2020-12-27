package mocks

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

//go:generate minimock -i pb-dropbox-downloader/internal.DataStorage -o ./mocks\data_storagemock.go

import (
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
)

// DataStorageMock implements internal.DataStorage
type DataStorageMock struct {
	t minimock.Tester

	funcFromMap          func(m1 map[string]string) (err error)
	inspectFuncFromMap   func(m1 map[string]string)
	afterFromMapCounter  uint64
	beforeFromMapCounter uint64
	FromMapMock          mDataStorageMockFromMap

	funcGet          func(s1 string) (s2 string, b1 bool)
	inspectFuncGet   func(s1 string)
	afterGetCounter  uint64
	beforeGetCounter uint64
	GetMock          mDataStorageMockGet

	funcToMap          func() (m1 map[string]string, err error)
	inspectFuncToMap   func()
	afterToMapCounter  uint64
	beforeToMapCounter uint64
	ToMapMock          mDataStorageMockToMap
}

// NewDataStorageMock returns a mock for internal.DataStorage
func NewDataStorageMock(t minimock.Tester) *DataStorageMock {
	m := &DataStorageMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.FromMapMock = mDataStorageMockFromMap{mock: m}
	m.FromMapMock.callArgs = []*DataStorageMockFromMapParams{}

	m.GetMock = mDataStorageMockGet{mock: m}
	m.GetMock.callArgs = []*DataStorageMockGetParams{}

	m.ToMapMock = mDataStorageMockToMap{mock: m}

	return m
}

type mDataStorageMockFromMap struct {
	mock               *DataStorageMock
	defaultExpectation *DataStorageMockFromMapExpectation
	expectations       []*DataStorageMockFromMapExpectation

	callArgs []*DataStorageMockFromMapParams
	mutex    sync.RWMutex
}

// DataStorageMockFromMapExpectation specifies expectation struct of the DataStorage.FromMap
type DataStorageMockFromMapExpectation struct {
	mock    *DataStorageMock
	params  *DataStorageMockFromMapParams
	results *DataStorageMockFromMapResults
	Counter uint64
}

// DataStorageMockFromMapParams contains parameters of the DataStorage.FromMap
type DataStorageMockFromMapParams struct {
	m1 map[string]string
}

// DataStorageMockFromMapResults contains results of the DataStorage.FromMap
type DataStorageMockFromMapResults struct {
	err error
}

// Expect sets up expected params for DataStorage.FromMap
func (mmFromMap *mDataStorageMockFromMap) Expect(m1 map[string]string) *mDataStorageMockFromMap {
	if mmFromMap.mock.funcFromMap != nil {
		mmFromMap.mock.t.Fatalf("DataStorageMock.FromMap mock is already set by Set")
	}

	if mmFromMap.defaultExpectation == nil {
		mmFromMap.defaultExpectation = &DataStorageMockFromMapExpectation{}
	}

	mmFromMap.defaultExpectation.params = &DataStorageMockFromMapParams{m1}
	for _, e := range mmFromMap.expectations {
		if minimock.Equal(e.params, mmFromMap.defaultExpectation.params) {
			mmFromMap.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmFromMap.defaultExpectation.params)
		}
	}

	return mmFromMap
}

// Inspect accepts an inspector function that has same arguments as the DataStorage.FromMap
func (mmFromMap *mDataStorageMockFromMap) Inspect(f func(m1 map[string]string)) *mDataStorageMockFromMap {
	if mmFromMap.mock.inspectFuncFromMap != nil {
		mmFromMap.mock.t.Fatalf("Inspect function is already set for DataStorageMock.FromMap")
	}

	mmFromMap.mock.inspectFuncFromMap = f

	return mmFromMap
}

// Return sets up results that will be returned by DataStorage.FromMap
func (mmFromMap *mDataStorageMockFromMap) Return(err error) *DataStorageMock {
	if mmFromMap.mock.funcFromMap != nil {
		mmFromMap.mock.t.Fatalf("DataStorageMock.FromMap mock is already set by Set")
	}

	if mmFromMap.defaultExpectation == nil {
		mmFromMap.defaultExpectation = &DataStorageMockFromMapExpectation{mock: mmFromMap.mock}
	}
	mmFromMap.defaultExpectation.results = &DataStorageMockFromMapResults{err}
	return mmFromMap.mock
}

//Set uses given function f to mock the DataStorage.FromMap method
func (mmFromMap *mDataStorageMockFromMap) Set(f func(m1 map[string]string) (err error)) *DataStorageMock {
	if mmFromMap.defaultExpectation != nil {
		mmFromMap.mock.t.Fatalf("Default expectation is already set for the DataStorage.FromMap method")
	}

	if len(mmFromMap.expectations) > 0 {
		mmFromMap.mock.t.Fatalf("Some expectations are already set for the DataStorage.FromMap method")
	}

	mmFromMap.mock.funcFromMap = f
	return mmFromMap.mock
}

// When sets expectation for the DataStorage.FromMap which will trigger the result defined by the following
// Then helper
func (mmFromMap *mDataStorageMockFromMap) When(m1 map[string]string) *DataStorageMockFromMapExpectation {
	if mmFromMap.mock.funcFromMap != nil {
		mmFromMap.mock.t.Fatalf("DataStorageMock.FromMap mock is already set by Set")
	}

	expectation := &DataStorageMockFromMapExpectation{
		mock:   mmFromMap.mock,
		params: &DataStorageMockFromMapParams{m1},
	}
	mmFromMap.expectations = append(mmFromMap.expectations, expectation)
	return expectation
}

// Then sets up DataStorage.FromMap return parameters for the expectation previously defined by the When method
func (e *DataStorageMockFromMapExpectation) Then(err error) *DataStorageMock {
	e.results = &DataStorageMockFromMapResults{err}
	return e.mock
}

// FromMap implements internal.DataStorage
func (mmFromMap *DataStorageMock) FromMap(m1 map[string]string) (err error) {
	mm_atomic.AddUint64(&mmFromMap.beforeFromMapCounter, 1)
	defer mm_atomic.AddUint64(&mmFromMap.afterFromMapCounter, 1)

	if mmFromMap.inspectFuncFromMap != nil {
		mmFromMap.inspectFuncFromMap(m1)
	}

	mm_params := &DataStorageMockFromMapParams{m1}

	// Record call args
	mmFromMap.FromMapMock.mutex.Lock()
	mmFromMap.FromMapMock.callArgs = append(mmFromMap.FromMapMock.callArgs, mm_params)
	mmFromMap.FromMapMock.mutex.Unlock()

	for _, e := range mmFromMap.FromMapMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmFromMap.FromMapMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmFromMap.FromMapMock.defaultExpectation.Counter, 1)
		mm_want := mmFromMap.FromMapMock.defaultExpectation.params
		mm_got := DataStorageMockFromMapParams{m1}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmFromMap.t.Errorf("DataStorageMock.FromMap got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmFromMap.FromMapMock.defaultExpectation.results
		if mm_results == nil {
			mmFromMap.t.Fatal("No results are set for the DataStorageMock.FromMap")
		}
		return (*mm_results).err
	}
	if mmFromMap.funcFromMap != nil {
		return mmFromMap.funcFromMap(m1)
	}
	mmFromMap.t.Fatalf("Unexpected call to DataStorageMock.FromMap. %v", m1)
	return
}

// FromMapAfterCounter returns a count of finished DataStorageMock.FromMap invocations
func (mmFromMap *DataStorageMock) FromMapAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmFromMap.afterFromMapCounter)
}

// FromMapBeforeCounter returns a count of DataStorageMock.FromMap invocations
func (mmFromMap *DataStorageMock) FromMapBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmFromMap.beforeFromMapCounter)
}

// Calls returns a list of arguments used in each call to DataStorageMock.FromMap.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmFromMap *mDataStorageMockFromMap) Calls() []*DataStorageMockFromMapParams {
	mmFromMap.mutex.RLock()

	argCopy := make([]*DataStorageMockFromMapParams, len(mmFromMap.callArgs))
	copy(argCopy, mmFromMap.callArgs)

	mmFromMap.mutex.RUnlock()

	return argCopy
}

// MinimockFromMapDone returns true if the count of the FromMap invocations corresponds
// the number of defined expectations
func (m *DataStorageMock) MinimockFromMapDone() bool {
	for _, e := range m.FromMapMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.FromMapMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterFromMapCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcFromMap != nil && mm_atomic.LoadUint64(&m.afterFromMapCounter) < 1 {
		return false
	}
	return true
}

// MinimockFromMapInspect logs each unmet expectation
func (m *DataStorageMock) MinimockFromMapInspect() {
	for _, e := range m.FromMapMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to DataStorageMock.FromMap with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.FromMapMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterFromMapCounter) < 1 {
		if m.FromMapMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to DataStorageMock.FromMap")
		} else {
			m.t.Errorf("Expected call to DataStorageMock.FromMap with params: %#v", *m.FromMapMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcFromMap != nil && mm_atomic.LoadUint64(&m.afterFromMapCounter) < 1 {
		m.t.Error("Expected call to DataStorageMock.FromMap")
	}
}

type mDataStorageMockGet struct {
	mock               *DataStorageMock
	defaultExpectation *DataStorageMockGetExpectation
	expectations       []*DataStorageMockGetExpectation

	callArgs []*DataStorageMockGetParams
	mutex    sync.RWMutex
}

// DataStorageMockGetExpectation specifies expectation struct of the DataStorage.Get
type DataStorageMockGetExpectation struct {
	mock    *DataStorageMock
	params  *DataStorageMockGetParams
	results *DataStorageMockGetResults
	Counter uint64
}

// DataStorageMockGetParams contains parameters of the DataStorage.Get
type DataStorageMockGetParams struct {
	s1 string
}

// DataStorageMockGetResults contains results of the DataStorage.Get
type DataStorageMockGetResults struct {
	s2 string
	b1 bool
}

// Expect sets up expected params for DataStorage.Get
func (mmGet *mDataStorageMockGet) Expect(s1 string) *mDataStorageMockGet {
	if mmGet.mock.funcGet != nil {
		mmGet.mock.t.Fatalf("DataStorageMock.Get mock is already set by Set")
	}

	if mmGet.defaultExpectation == nil {
		mmGet.defaultExpectation = &DataStorageMockGetExpectation{}
	}

	mmGet.defaultExpectation.params = &DataStorageMockGetParams{s1}
	for _, e := range mmGet.expectations {
		if minimock.Equal(e.params, mmGet.defaultExpectation.params) {
			mmGet.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmGet.defaultExpectation.params)
		}
	}

	return mmGet
}

// Inspect accepts an inspector function that has same arguments as the DataStorage.Get
func (mmGet *mDataStorageMockGet) Inspect(f func(s1 string)) *mDataStorageMockGet {
	if mmGet.mock.inspectFuncGet != nil {
		mmGet.mock.t.Fatalf("Inspect function is already set for DataStorageMock.Get")
	}

	mmGet.mock.inspectFuncGet = f

	return mmGet
}

// Return sets up results that will be returned by DataStorage.Get
func (mmGet *mDataStorageMockGet) Return(s2 string, b1 bool) *DataStorageMock {
	if mmGet.mock.funcGet != nil {
		mmGet.mock.t.Fatalf("DataStorageMock.Get mock is already set by Set")
	}

	if mmGet.defaultExpectation == nil {
		mmGet.defaultExpectation = &DataStorageMockGetExpectation{mock: mmGet.mock}
	}
	mmGet.defaultExpectation.results = &DataStorageMockGetResults{s2, b1}
	return mmGet.mock
}

//Set uses given function f to mock the DataStorage.Get method
func (mmGet *mDataStorageMockGet) Set(f func(s1 string) (s2 string, b1 bool)) *DataStorageMock {
	if mmGet.defaultExpectation != nil {
		mmGet.mock.t.Fatalf("Default expectation is already set for the DataStorage.Get method")
	}

	if len(mmGet.expectations) > 0 {
		mmGet.mock.t.Fatalf("Some expectations are already set for the DataStorage.Get method")
	}

	mmGet.mock.funcGet = f
	return mmGet.mock
}

// When sets expectation for the DataStorage.Get which will trigger the result defined by the following
// Then helper
func (mmGet *mDataStorageMockGet) When(s1 string) *DataStorageMockGetExpectation {
	if mmGet.mock.funcGet != nil {
		mmGet.mock.t.Fatalf("DataStorageMock.Get mock is already set by Set")
	}

	expectation := &DataStorageMockGetExpectation{
		mock:   mmGet.mock,
		params: &DataStorageMockGetParams{s1},
	}
	mmGet.expectations = append(mmGet.expectations, expectation)
	return expectation
}

// Then sets up DataStorage.Get return parameters for the expectation previously defined by the When method
func (e *DataStorageMockGetExpectation) Then(s2 string, b1 bool) *DataStorageMock {
	e.results = &DataStorageMockGetResults{s2, b1}
	return e.mock
}

// Get implements internal.DataStorage
func (mmGet *DataStorageMock) Get(s1 string) (s2 string, b1 bool) {
	mm_atomic.AddUint64(&mmGet.beforeGetCounter, 1)
	defer mm_atomic.AddUint64(&mmGet.afterGetCounter, 1)

	if mmGet.inspectFuncGet != nil {
		mmGet.inspectFuncGet(s1)
	}

	mm_params := &DataStorageMockGetParams{s1}

	// Record call args
	mmGet.GetMock.mutex.Lock()
	mmGet.GetMock.callArgs = append(mmGet.GetMock.callArgs, mm_params)
	mmGet.GetMock.mutex.Unlock()

	for _, e := range mmGet.GetMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.s2, e.results.b1
		}
	}

	if mmGet.GetMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmGet.GetMock.defaultExpectation.Counter, 1)
		mm_want := mmGet.GetMock.defaultExpectation.params
		mm_got := DataStorageMockGetParams{s1}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmGet.t.Errorf("DataStorageMock.Get got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmGet.GetMock.defaultExpectation.results
		if mm_results == nil {
			mmGet.t.Fatal("No results are set for the DataStorageMock.Get")
		}
		return (*mm_results).s2, (*mm_results).b1
	}
	if mmGet.funcGet != nil {
		return mmGet.funcGet(s1)
	}
	mmGet.t.Fatalf("Unexpected call to DataStorageMock.Get. %v", s1)
	return
}

// GetAfterCounter returns a count of finished DataStorageMock.Get invocations
func (mmGet *DataStorageMock) GetAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGet.afterGetCounter)
}

// GetBeforeCounter returns a count of DataStorageMock.Get invocations
func (mmGet *DataStorageMock) GetBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGet.beforeGetCounter)
}

// Calls returns a list of arguments used in each call to DataStorageMock.Get.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmGet *mDataStorageMockGet) Calls() []*DataStorageMockGetParams {
	mmGet.mutex.RLock()

	argCopy := make([]*DataStorageMockGetParams, len(mmGet.callArgs))
	copy(argCopy, mmGet.callArgs)

	mmGet.mutex.RUnlock()

	return argCopy
}

// MinimockGetDone returns true if the count of the Get invocations corresponds
// the number of defined expectations
func (m *DataStorageMock) MinimockGetDone() bool {
	for _, e := range m.GetMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGet != nil && mm_atomic.LoadUint64(&m.afterGetCounter) < 1 {
		return false
	}
	return true
}

// MinimockGetInspect logs each unmet expectation
func (m *DataStorageMock) MinimockGetInspect() {
	for _, e := range m.GetMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to DataStorageMock.Get with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetCounter) < 1 {
		if m.GetMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to DataStorageMock.Get")
		} else {
			m.t.Errorf("Expected call to DataStorageMock.Get with params: %#v", *m.GetMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGet != nil && mm_atomic.LoadUint64(&m.afterGetCounter) < 1 {
		m.t.Error("Expected call to DataStorageMock.Get")
	}
}

type mDataStorageMockToMap struct {
	mock               *DataStorageMock
	defaultExpectation *DataStorageMockToMapExpectation
	expectations       []*DataStorageMockToMapExpectation
}

// DataStorageMockToMapExpectation specifies expectation struct of the DataStorage.ToMap
type DataStorageMockToMapExpectation struct {
	mock *DataStorageMock

	results *DataStorageMockToMapResults
	Counter uint64
}

// DataStorageMockToMapResults contains results of the DataStorage.ToMap
type DataStorageMockToMapResults struct {
	m1  map[string]string
	err error
}

// Expect sets up expected params for DataStorage.ToMap
func (mmToMap *mDataStorageMockToMap) Expect() *mDataStorageMockToMap {
	if mmToMap.mock.funcToMap != nil {
		mmToMap.mock.t.Fatalf("DataStorageMock.ToMap mock is already set by Set")
	}

	if mmToMap.defaultExpectation == nil {
		mmToMap.defaultExpectation = &DataStorageMockToMapExpectation{}
	}

	return mmToMap
}

// Inspect accepts an inspector function that has same arguments as the DataStorage.ToMap
func (mmToMap *mDataStorageMockToMap) Inspect(f func()) *mDataStorageMockToMap {
	if mmToMap.mock.inspectFuncToMap != nil {
		mmToMap.mock.t.Fatalf("Inspect function is already set for DataStorageMock.ToMap")
	}

	mmToMap.mock.inspectFuncToMap = f

	return mmToMap
}

// Return sets up results that will be returned by DataStorage.ToMap
func (mmToMap *mDataStorageMockToMap) Return(m1 map[string]string, err error) *DataStorageMock {
	if mmToMap.mock.funcToMap != nil {
		mmToMap.mock.t.Fatalf("DataStorageMock.ToMap mock is already set by Set")
	}

	if mmToMap.defaultExpectation == nil {
		mmToMap.defaultExpectation = &DataStorageMockToMapExpectation{mock: mmToMap.mock}
	}
	mmToMap.defaultExpectation.results = &DataStorageMockToMapResults{m1, err}
	return mmToMap.mock
}

//Set uses given function f to mock the DataStorage.ToMap method
func (mmToMap *mDataStorageMockToMap) Set(f func() (m1 map[string]string, err error)) *DataStorageMock {
	if mmToMap.defaultExpectation != nil {
		mmToMap.mock.t.Fatalf("Default expectation is already set for the DataStorage.ToMap method")
	}

	if len(mmToMap.expectations) > 0 {
		mmToMap.mock.t.Fatalf("Some expectations are already set for the DataStorage.ToMap method")
	}

	mmToMap.mock.funcToMap = f
	return mmToMap.mock
}

// ToMap implements internal.DataStorage
func (mmToMap *DataStorageMock) ToMap() (m1 map[string]string, err error) {
	mm_atomic.AddUint64(&mmToMap.beforeToMapCounter, 1)
	defer mm_atomic.AddUint64(&mmToMap.afterToMapCounter, 1)

	if mmToMap.inspectFuncToMap != nil {
		mmToMap.inspectFuncToMap()
	}

	if mmToMap.ToMapMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmToMap.ToMapMock.defaultExpectation.Counter, 1)

		mm_results := mmToMap.ToMapMock.defaultExpectation.results
		if mm_results == nil {
			mmToMap.t.Fatal("No results are set for the DataStorageMock.ToMap")
		}
		return (*mm_results).m1, (*mm_results).err
	}
	if mmToMap.funcToMap != nil {
		return mmToMap.funcToMap()
	}
	mmToMap.t.Fatalf("Unexpected call to DataStorageMock.ToMap.")
	return
}

// ToMapAfterCounter returns a count of finished DataStorageMock.ToMap invocations
func (mmToMap *DataStorageMock) ToMapAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmToMap.afterToMapCounter)
}

// ToMapBeforeCounter returns a count of DataStorageMock.ToMap invocations
func (mmToMap *DataStorageMock) ToMapBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmToMap.beforeToMapCounter)
}

// MinimockToMapDone returns true if the count of the ToMap invocations corresponds
// the number of defined expectations
func (m *DataStorageMock) MinimockToMapDone() bool {
	for _, e := range m.ToMapMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.ToMapMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterToMapCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcToMap != nil && mm_atomic.LoadUint64(&m.afterToMapCounter) < 1 {
		return false
	}
	return true
}

// MinimockToMapInspect logs each unmet expectation
func (m *DataStorageMock) MinimockToMapInspect() {
	for _, e := range m.ToMapMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Error("Expected call to DataStorageMock.ToMap")
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.ToMapMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterToMapCounter) < 1 {
		m.t.Error("Expected call to DataStorageMock.ToMap")
	}
	// if func was set then invocations count should be greater than zero
	if m.funcToMap != nil && mm_atomic.LoadUint64(&m.afterToMapCounter) < 1 {
		m.t.Error("Expected call to DataStorageMock.ToMap")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *DataStorageMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockFromMapInspect()

		m.MinimockGetInspect()

		m.MinimockToMapInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *DataStorageMock) MinimockWait(timeout mm_time.Duration) {
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

func (m *DataStorageMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockFromMapDone() &&
		m.MinimockGetDone() &&
		m.MinimockToMapDone()
}
