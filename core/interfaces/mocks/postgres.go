// Code generated by MockGen. DO NOT EDIT.
// Source: core/interfaces/postgres.go

// Package mock_interfaces is a generated GoMock package.
package mock_interfaces

import (
	context "context"
	sql "database/sql"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	sqlx "github.com/jmoiron/sqlx"
)

// MockISqlDatabase is a mock of ISqlDatabase interface.
type MockISqlDatabase struct {
	ctrl     *gomock.Controller
	recorder *MockISqlDatabaseMockRecorder
}

// MockISqlDatabaseMockRecorder is the mock recorder for MockISqlDatabase.
type MockISqlDatabaseMockRecorder struct {
	mock *MockISqlDatabase
}

// NewMockISqlDatabase creates a new mock instance.
func NewMockISqlDatabase(ctrl *gomock.Controller) *MockISqlDatabase {
	mock := &MockISqlDatabase{ctrl: ctrl}
	mock.recorder = &MockISqlDatabaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockISqlDatabase) EXPECT() *MockISqlDatabaseMockRecorder {
	return m.recorder
}

// Exec mocks base method.
func (m *MockISqlDatabase) Exec(query string, args ...interface{}) (sql.Result, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{query}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Exec", varargs...)
	ret0, _ := ret[0].(sql.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exec indicates an expected call of Exec.
func (mr *MockISqlDatabaseMockRecorder) Exec(query interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{query}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exec", reflect.TypeOf((*MockISqlDatabase)(nil).Exec), varargs...)
}

// ExecContext mocks base method.
func (m *MockISqlDatabase) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, query}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ExecContext", varargs...)
	ret0, _ := ret[0].(sql.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExecContext indicates an expected call of ExecContext.
func (mr *MockISqlDatabaseMockRecorder) ExecContext(ctx, query interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, query}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExecContext", reflect.TypeOf((*MockISqlDatabase)(nil).ExecContext), varargs...)
}

// GetContext mocks base method.
func (m *MockISqlDatabase) GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, dest, query}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetContext", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetContext indicates an expected call of GetContext.
func (mr *MockISqlDatabaseMockRecorder) GetContext(ctx, dest, query interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, dest, query}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetContext", reflect.TypeOf((*MockISqlDatabase)(nil).GetContext), varargs...)
}

// NamedExecContext mocks base method.
func (m *MockISqlDatabase) NamedExecContext(ctx context.Context, query string, arg interface{}) (sql.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NamedExecContext", ctx, query, arg)
	ret0, _ := ret[0].(sql.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NamedExecContext indicates an expected call of NamedExecContext.
func (mr *MockISqlDatabaseMockRecorder) NamedExecContext(ctx, query, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NamedExecContext", reflect.TypeOf((*MockISqlDatabase)(nil).NamedExecContext), ctx, query, arg)
}

// Prepare mocks base method.
func (m *MockISqlDatabase) Prepare(query string) (*sql.Stmt, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Prepare", query)
	ret0, _ := ret[0].(*sql.Stmt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Prepare indicates an expected call of Prepare.
func (mr *MockISqlDatabaseMockRecorder) Prepare(query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Prepare", reflect.TypeOf((*MockISqlDatabase)(nil).Prepare), query)
}

// PrepareContext mocks base method.
func (m *MockISqlDatabase) PrepareContext(ctx context.Context, query string) (*sql.Stmt, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PrepareContext", ctx, query)
	ret0, _ := ret[0].(*sql.Stmt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PrepareContext indicates an expected call of PrepareContext.
func (mr *MockISqlDatabaseMockRecorder) PrepareContext(ctx, query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PrepareContext", reflect.TypeOf((*MockISqlDatabase)(nil).PrepareContext), ctx, query)
}

// Query mocks base method.
func (m *MockISqlDatabase) Query(query string, args ...interface{}) (*sql.Rows, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{query}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Query", varargs...)
	ret0, _ := ret[0].(*sql.Rows)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Query indicates an expected call of Query.
func (mr *MockISqlDatabaseMockRecorder) Query(query interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{query}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Query", reflect.TypeOf((*MockISqlDatabase)(nil).Query), varargs...)
}

// QueryContext mocks base method.
func (m *MockISqlDatabase) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, query}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "QueryContext", varargs...)
	ret0, _ := ret[0].(*sql.Rows)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryContext indicates an expected call of QueryContext.
func (mr *MockISqlDatabaseMockRecorder) QueryContext(ctx, query interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, query}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryContext", reflect.TypeOf((*MockISqlDatabase)(nil).QueryContext), varargs...)
}

// QueryRow mocks base method.
func (m *MockISqlDatabase) QueryRow(query string, args ...interface{}) *sql.Row {
	m.ctrl.T.Helper()
	varargs := []interface{}{query}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "QueryRow", varargs...)
	ret0, _ := ret[0].(*sql.Row)
	return ret0
}

// QueryRow indicates an expected call of QueryRow.
func (mr *MockISqlDatabaseMockRecorder) QueryRow(query interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{query}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryRow", reflect.TypeOf((*MockISqlDatabase)(nil).QueryRow), varargs...)
}

// QueryRowContext mocks base method.
func (m *MockISqlDatabase) QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, query}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "QueryRowContext", varargs...)
	ret0, _ := ret[0].(*sql.Row)
	return ret0
}

// QueryRowContext indicates an expected call of QueryRowContext.
func (mr *MockISqlDatabaseMockRecorder) QueryRowContext(ctx, query interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, query}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryRowContext", reflect.TypeOf((*MockISqlDatabase)(nil).QueryRowContext), varargs...)
}

// QueryRowx mocks base method.
func (m *MockISqlDatabase) QueryRowx(query string, args ...interface{}) *sqlx.Row {
	m.ctrl.T.Helper()
	varargs := []interface{}{query}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "QueryRowx", varargs...)
	ret0, _ := ret[0].(*sqlx.Row)
	return ret0
}

// QueryRowx indicates an expected call of QueryRowx.
func (mr *MockISqlDatabaseMockRecorder) QueryRowx(query interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{query}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryRowx", reflect.TypeOf((*MockISqlDatabase)(nil).QueryRowx), varargs...)
}

// QueryRowxContext mocks base method.
func (m *MockISqlDatabase) QueryRowxContext(ctx context.Context, query string, args ...interface{}) *sqlx.Row {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, query}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "QueryRowxContext", varargs...)
	ret0, _ := ret[0].(*sqlx.Row)
	return ret0
}

// QueryRowxContext indicates an expected call of QueryRowxContext.
func (mr *MockISqlDatabaseMockRecorder) QueryRowxContext(ctx, query interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, query}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryRowxContext", reflect.TypeOf((*MockISqlDatabase)(nil).QueryRowxContext), varargs...)
}

// Queryx mocks base method.
func (m *MockISqlDatabase) Queryx(query string, args ...interface{}) (*sqlx.Rows, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{query}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Queryx", varargs...)
	ret0, _ := ret[0].(*sqlx.Rows)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Queryx indicates an expected call of Queryx.
func (mr *MockISqlDatabaseMockRecorder) Queryx(query interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{query}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Queryx", reflect.TypeOf((*MockISqlDatabase)(nil).Queryx), varargs...)
}

// QueryxContext mocks base method.
func (m *MockISqlDatabase) QueryxContext(ctx context.Context, query string, args ...interface{}) (*sqlx.Rows, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, query}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "QueryxContext", varargs...)
	ret0, _ := ret[0].(*sqlx.Rows)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryxContext indicates an expected call of QueryxContext.
func (mr *MockISqlDatabaseMockRecorder) QueryxContext(ctx, query interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, query}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryxContext", reflect.TypeOf((*MockISqlDatabase)(nil).QueryxContext), varargs...)
}

// SelectContext mocks base method.
func (m *MockISqlDatabase) SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, dest, query}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SelectContext", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// SelectContext indicates an expected call of SelectContext.
func (mr *MockISqlDatabaseMockRecorder) SelectContext(ctx, dest, query interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, dest, query}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectContext", reflect.TypeOf((*MockISqlDatabase)(nil).SelectContext), varargs...)
}

// MockIRow is a mock of IRow interface.
type MockIRow struct {
	ctrl     *gomock.Controller
	recorder *MockIRowMockRecorder
}

// MockIRowMockRecorder is the mock recorder for MockIRow.
type MockIRowMockRecorder struct {
	mock *MockIRow
}

// NewMockIRow creates a new mock instance.
func NewMockIRow(ctrl *gomock.Controller) *MockIRow {
	mock := &MockIRow{ctrl: ctrl}
	mock.recorder = &MockIRowMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIRow) EXPECT() *MockIRowMockRecorder {
	return m.recorder
}

// Scan mocks base method.
func (m *MockIRow) Scan(dest ...any) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range dest {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Scan", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Scan indicates an expected call of Scan.
func (mr *MockIRowMockRecorder) Scan(dest ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Scan", reflect.TypeOf((*MockIRow)(nil).Scan), dest...)
}
