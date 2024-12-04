// Code generated by http://github.com/gojuno/minimock (v3.4.3). DO NOT EDIT.

package mocks

//go:generate minimock -i platform-pkg/pkg/database.Client -o client_minimock.go -n ClientMock -p mocks

import (
	mm_database "platform-pkg/pkg/database"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
)

// ClientMock implements mm_database.Client
type ClientMock struct {
	t          minimock.Tester
	finishOnce sync.Once

	funcClose          func() (err error)
	funcCloseOrigin    string
	inspectFuncClose   func()
	afterCloseCounter  uint64
	beforeCloseCounter uint64
	CloseMock          mClientMockClose

	funcDB          func() (d1 mm_database.DB)
	funcDBOrigin    string
	inspectFuncDB   func()
	afterDBCounter  uint64
	beforeDBCounter uint64
	DBMock          mClientMockDB
}

// NewClientMock returns a mock for mm_database.Client
func NewClientMock(t minimock.Tester) *ClientMock {
	m := &ClientMock{t: t}

	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.CloseMock = mClientMockClose{mock: m}

	m.DBMock = mClientMockDB{mock: m}

	t.Cleanup(m.MinimockFinish)

	return m
}

type mClientMockClose struct {
	optional           bool
	mock               *ClientMock
	defaultExpectation *ClientMockCloseExpectation
	expectations       []*ClientMockCloseExpectation

	expectedInvocations       uint64
	expectedInvocationsOrigin string
}

// ClientMockCloseExpectation specifies expectation struct of the Client.Close
type ClientMockCloseExpectation struct {
	mock *ClientMock

	results      *ClientMockCloseResults
	returnOrigin string
	Counter      uint64
}

// ClientMockCloseResults contains results of the Client.Close
type ClientMockCloseResults struct {
	err error
}

// Marks this method to be optional. The default behavior of any method with Return() is '1 or more', meaning
// the test will fail minimock's automatic final call check if the mocked method was not called at least once.
// Optional() makes method check to work in '0 or more' mode.
// It is NOT RECOMMENDED to use this option unless you really need it, as default behaviour helps to
// catch the problems when the expected method call is totally skipped during test run.
func (mmClose *mClientMockClose) Optional() *mClientMockClose {
	mmClose.optional = true
	return mmClose
}

// Expect sets up expected params for Client.Close
func (mmClose *mClientMockClose) Expect() *mClientMockClose {
	if mmClose.mock.funcClose != nil {
		mmClose.mock.t.Fatalf("ClientMock.Close mock is already set by Set")
	}

	if mmClose.defaultExpectation == nil {
		mmClose.defaultExpectation = &ClientMockCloseExpectation{}
	}

	return mmClose
}

// Inspect accepts an inspector function that has same arguments as the Client.Close
func (mmClose *mClientMockClose) Inspect(f func()) *mClientMockClose {
	if mmClose.mock.inspectFuncClose != nil {
		mmClose.mock.t.Fatalf("Inspect function is already set for ClientMock.Close")
	}

	mmClose.mock.inspectFuncClose = f

	return mmClose
}

// Return sets up results that will be returned by Client.Close
func (mmClose *mClientMockClose) Return(err error) *ClientMock {
	if mmClose.mock.funcClose != nil {
		mmClose.mock.t.Fatalf("ClientMock.Close mock is already set by Set")
	}

	if mmClose.defaultExpectation == nil {
		mmClose.defaultExpectation = &ClientMockCloseExpectation{mock: mmClose.mock}
	}
	mmClose.defaultExpectation.results = &ClientMockCloseResults{err}
	mmClose.defaultExpectation.returnOrigin = minimock.CallerInfo(1)
	return mmClose.mock
}

// Set uses given function f to mock the Client.Close method
func (mmClose *mClientMockClose) Set(f func() (err error)) *ClientMock {
	if mmClose.defaultExpectation != nil {
		mmClose.mock.t.Fatalf("Default expectation is already set for the Client.Close method")
	}

	if len(mmClose.expectations) > 0 {
		mmClose.mock.t.Fatalf("Some expectations are already set for the Client.Close method")
	}

	mmClose.mock.funcClose = f
	mmClose.mock.funcCloseOrigin = minimock.CallerInfo(1)
	return mmClose.mock
}

// Times sets number of times Client.Close should be invoked
func (mmClose *mClientMockClose) Times(n uint64) *mClientMockClose {
	if n == 0 {
		mmClose.mock.t.Fatalf("Times of ClientMock.Close mock can not be zero")
	}
	mm_atomic.StoreUint64(&mmClose.expectedInvocations, n)
	mmClose.expectedInvocationsOrigin = minimock.CallerInfo(1)
	return mmClose
}

func (mmClose *mClientMockClose) invocationsDone() bool {
	if len(mmClose.expectations) == 0 && mmClose.defaultExpectation == nil && mmClose.mock.funcClose == nil {
		return true
	}

	totalInvocations := mm_atomic.LoadUint64(&mmClose.mock.afterCloseCounter)
	expectedInvocations := mm_atomic.LoadUint64(&mmClose.expectedInvocations)

	return totalInvocations > 0 && (expectedInvocations == 0 || expectedInvocations == totalInvocations)
}

// Close implements mm_database.Client
func (mmClose *ClientMock) Close() (err error) {
	mm_atomic.AddUint64(&mmClose.beforeCloseCounter, 1)
	defer mm_atomic.AddUint64(&mmClose.afterCloseCounter, 1)

	mmClose.t.Helper()

	if mmClose.inspectFuncClose != nil {
		mmClose.inspectFuncClose()
	}

	if mmClose.CloseMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmClose.CloseMock.defaultExpectation.Counter, 1)

		mm_results := mmClose.CloseMock.defaultExpectation.results
		if mm_results == nil {
			mmClose.t.Fatal("No results are set for the ClientMock.Close")
		}
		return (*mm_results).err
	}
	if mmClose.funcClose != nil {
		return mmClose.funcClose()
	}
	mmClose.t.Fatalf("Unexpected call to ClientMock.Close.")
	return
}

// CloseAfterCounter returns a count of finished ClientMock.Close invocations
func (mmClose *ClientMock) CloseAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmClose.afterCloseCounter)
}

// CloseBeforeCounter returns a count of ClientMock.Close invocations
func (mmClose *ClientMock) CloseBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmClose.beforeCloseCounter)
}

// MinimockCloseDone returns true if the count of the Close invocations corresponds
// the number of defined expectations
func (m *ClientMock) MinimockCloseDone() bool {
	if m.CloseMock.optional {
		// Optional methods provide '0 or more' call count restriction.
		return true
	}

	for _, e := range m.CloseMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	return m.CloseMock.invocationsDone()
}

// MinimockCloseInspect logs each unmet expectation
func (m *ClientMock) MinimockCloseInspect() {
	for _, e := range m.CloseMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Error("Expected call to ClientMock.Close")
		}
	}

	afterCloseCounter := mm_atomic.LoadUint64(&m.afterCloseCounter)
	// if default expectation was set then invocations count should be greater than zero
	if m.CloseMock.defaultExpectation != nil && afterCloseCounter < 1 {
		m.t.Errorf("Expected call to ClientMock.Close at\n%s", m.CloseMock.defaultExpectation.returnOrigin)
	}
	// if func was set then invocations count should be greater than zero
	if m.funcClose != nil && afterCloseCounter < 1 {
		m.t.Errorf("Expected call to ClientMock.Close at\n%s", m.funcCloseOrigin)
	}

	if !m.CloseMock.invocationsDone() && afterCloseCounter > 0 {
		m.t.Errorf("Expected %d calls to ClientMock.Close at\n%s but found %d calls",
			mm_atomic.LoadUint64(&m.CloseMock.expectedInvocations), m.CloseMock.expectedInvocationsOrigin, afterCloseCounter)
	}
}

type mClientMockDB struct {
	optional           bool
	mock               *ClientMock
	defaultExpectation *ClientMockDBExpectation
	expectations       []*ClientMockDBExpectation

	expectedInvocations       uint64
	expectedInvocationsOrigin string
}

// ClientMockDBExpectation specifies expectation struct of the Client.DB
type ClientMockDBExpectation struct {
	mock *ClientMock

	results      *ClientMockDBResults
	returnOrigin string
	Counter      uint64
}

// ClientMockDBResults contains results of the Client.DB
type ClientMockDBResults struct {
	d1 mm_database.DB
}

// Marks this method to be optional. The default behavior of any method with Return() is '1 or more', meaning
// the test will fail minimock's automatic final call check if the mocked method was not called at least once.
// Optional() makes method check to work in '0 or more' mode.
// It is NOT RECOMMENDED to use this option unless you really need it, as default behaviour helps to
// catch the problems when the expected method call is totally skipped during test run.
func (mmDB *mClientMockDB) Optional() *mClientMockDB {
	mmDB.optional = true
	return mmDB
}

// Expect sets up expected params for Client.DB
func (mmDB *mClientMockDB) Expect() *mClientMockDB {
	if mmDB.mock.funcDB != nil {
		mmDB.mock.t.Fatalf("ClientMock.DB mock is already set by Set")
	}

	if mmDB.defaultExpectation == nil {
		mmDB.defaultExpectation = &ClientMockDBExpectation{}
	}

	return mmDB
}

// Inspect accepts an inspector function that has same arguments as the Client.DB
func (mmDB *mClientMockDB) Inspect(f func()) *mClientMockDB {
	if mmDB.mock.inspectFuncDB != nil {
		mmDB.mock.t.Fatalf("Inspect function is already set for ClientMock.DB")
	}

	mmDB.mock.inspectFuncDB = f

	return mmDB
}

// Return sets up results that will be returned by Client.DB
func (mmDB *mClientMockDB) Return(d1 mm_database.DB) *ClientMock {
	if mmDB.mock.funcDB != nil {
		mmDB.mock.t.Fatalf("ClientMock.DB mock is already set by Set")
	}

	if mmDB.defaultExpectation == nil {
		mmDB.defaultExpectation = &ClientMockDBExpectation{mock: mmDB.mock}
	}
	mmDB.defaultExpectation.results = &ClientMockDBResults{d1}
	mmDB.defaultExpectation.returnOrigin = minimock.CallerInfo(1)
	return mmDB.mock
}

// Set uses given function f to mock the Client.DB method
func (mmDB *mClientMockDB) Set(f func() (d1 mm_database.DB)) *ClientMock {
	if mmDB.defaultExpectation != nil {
		mmDB.mock.t.Fatalf("Default expectation is already set for the Client.DB method")
	}

	if len(mmDB.expectations) > 0 {
		mmDB.mock.t.Fatalf("Some expectations are already set for the Client.DB method")
	}

	mmDB.mock.funcDB = f
	mmDB.mock.funcDBOrigin = minimock.CallerInfo(1)
	return mmDB.mock
}

// Times sets number of times Client.DB should be invoked
func (mmDB *mClientMockDB) Times(n uint64) *mClientMockDB {
	if n == 0 {
		mmDB.mock.t.Fatalf("Times of ClientMock.DB mock can not be zero")
	}
	mm_atomic.StoreUint64(&mmDB.expectedInvocations, n)
	mmDB.expectedInvocationsOrigin = minimock.CallerInfo(1)
	return mmDB
}

func (mmDB *mClientMockDB) invocationsDone() bool {
	if len(mmDB.expectations) == 0 && mmDB.defaultExpectation == nil && mmDB.mock.funcDB == nil {
		return true
	}

	totalInvocations := mm_atomic.LoadUint64(&mmDB.mock.afterDBCounter)
	expectedInvocations := mm_atomic.LoadUint64(&mmDB.expectedInvocations)

	return totalInvocations > 0 && (expectedInvocations == 0 || expectedInvocations == totalInvocations)
}

// DB implements mm_database.Client
func (mmDB *ClientMock) DB() (d1 mm_database.DB) {
	mm_atomic.AddUint64(&mmDB.beforeDBCounter, 1)
	defer mm_atomic.AddUint64(&mmDB.afterDBCounter, 1)

	mmDB.t.Helper()

	if mmDB.inspectFuncDB != nil {
		mmDB.inspectFuncDB()
	}

	if mmDB.DBMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmDB.DBMock.defaultExpectation.Counter, 1)

		mm_results := mmDB.DBMock.defaultExpectation.results
		if mm_results == nil {
			mmDB.t.Fatal("No results are set for the ClientMock.DB")
		}
		return (*mm_results).d1
	}
	if mmDB.funcDB != nil {
		return mmDB.funcDB()
	}
	mmDB.t.Fatalf("Unexpected call to ClientMock.DB.")
	return
}

// DBAfterCounter returns a count of finished ClientMock.DB invocations
func (mmDB *ClientMock) DBAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmDB.afterDBCounter)
}

// DBBeforeCounter returns a count of ClientMock.DB invocations
func (mmDB *ClientMock) DBBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmDB.beforeDBCounter)
}

// MinimockDBDone returns true if the count of the DB invocations corresponds
// the number of defined expectations
func (m *ClientMock) MinimockDBDone() bool {
	if m.DBMock.optional {
		// Optional methods provide '0 or more' call count restriction.
		return true
	}

	for _, e := range m.DBMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	return m.DBMock.invocationsDone()
}

// MinimockDBInspect logs each unmet expectation
func (m *ClientMock) MinimockDBInspect() {
	for _, e := range m.DBMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Error("Expected call to ClientMock.DB")
		}
	}

	afterDBCounter := mm_atomic.LoadUint64(&m.afterDBCounter)
	// if default expectation was set then invocations count should be greater than zero
	if m.DBMock.defaultExpectation != nil && afterDBCounter < 1 {
		m.t.Errorf("Expected call to ClientMock.DB at\n%s", m.DBMock.defaultExpectation.returnOrigin)
	}
	// if func was set then invocations count should be greater than zero
	if m.funcDB != nil && afterDBCounter < 1 {
		m.t.Errorf("Expected call to ClientMock.DB at\n%s", m.funcDBOrigin)
	}

	if !m.DBMock.invocationsDone() && afterDBCounter > 0 {
		m.t.Errorf("Expected %d calls to ClientMock.DB at\n%s but found %d calls",
			mm_atomic.LoadUint64(&m.DBMock.expectedInvocations), m.DBMock.expectedInvocationsOrigin, afterDBCounter)
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *ClientMock) MinimockFinish() {
	m.finishOnce.Do(func() {
		if !m.minimockDone() {
			m.MinimockCloseInspect()

			m.MinimockDBInspect()
		}
	})
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *ClientMock) MinimockWait(timeout mm_time.Duration) {
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

func (m *ClientMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockCloseDone() &&
		m.MinimockDBDone()
}
