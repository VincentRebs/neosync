// Code generated by mockery. DO NOT EDIT.

package awsmanager

import (
	context "context"

	dynamodb "github.com/aws/aws-sdk-go-v2/service/dynamodb"
	mock "github.com/stretchr/testify/mock"
)

// MockdynamoDBAPIV2 is an autogenerated mock type for the dynamoDBAPIV2 type
type MockdynamoDBAPIV2 struct {
	mock.Mock
}

type MockdynamoDBAPIV2_Expecter struct {
	mock *mock.Mock
}

func (_m *MockdynamoDBAPIV2) EXPECT() *MockdynamoDBAPIV2_Expecter {
	return &MockdynamoDBAPIV2_Expecter{mock: &_m.Mock}
}

// DescribeTable provides a mock function with given fields: ctx, params, optFns
func (_m *MockdynamoDBAPIV2) DescribeTable(ctx context.Context, params *dynamodb.DescribeTableInput, optFns ...func(*dynamodb.Options)) (*dynamodb.DescribeTableOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DescribeTable")
	}

	var r0 *dynamodb.DescribeTableOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *dynamodb.DescribeTableInput, ...func(*dynamodb.Options)) (*dynamodb.DescribeTableOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *dynamodb.DescribeTableInput, ...func(*dynamodb.Options)) *dynamodb.DescribeTableOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dynamodb.DescribeTableOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *dynamodb.DescribeTableInput, ...func(*dynamodb.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockdynamoDBAPIV2_DescribeTable_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DescribeTable'
type MockdynamoDBAPIV2_DescribeTable_Call struct {
	*mock.Call
}

// DescribeTable is a helper method to define mock.On call
//   - ctx context.Context
//   - params *dynamodb.DescribeTableInput
//   - optFns ...func(*dynamodb.Options)
func (_e *MockdynamoDBAPIV2_Expecter) DescribeTable(ctx interface{}, params interface{}, optFns ...interface{}) *MockdynamoDBAPIV2_DescribeTable_Call {
	return &MockdynamoDBAPIV2_DescribeTable_Call{Call: _e.mock.On("DescribeTable",
		append([]interface{}{ctx, params}, optFns...)...)}
}

func (_c *MockdynamoDBAPIV2_DescribeTable_Call) Run(run func(ctx context.Context, params *dynamodb.DescribeTableInput, optFns ...func(*dynamodb.Options))) *MockdynamoDBAPIV2_DescribeTable_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]func(*dynamodb.Options), len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(func(*dynamodb.Options))
			}
		}
		run(args[0].(context.Context), args[1].(*dynamodb.DescribeTableInput), variadicArgs...)
	})
	return _c
}

func (_c *MockdynamoDBAPIV2_DescribeTable_Call) Return(_a0 *dynamodb.DescribeTableOutput, _a1 error) *MockdynamoDBAPIV2_DescribeTable_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockdynamoDBAPIV2_DescribeTable_Call) RunAndReturn(run func(context.Context, *dynamodb.DescribeTableInput, ...func(*dynamodb.Options)) (*dynamodb.DescribeTableOutput, error)) *MockdynamoDBAPIV2_DescribeTable_Call {
	_c.Call.Return(run)
	return _c
}

// ListTables provides a mock function with given fields: ctx, params, optFns
func (_m *MockdynamoDBAPIV2) ListTables(ctx context.Context, params *dynamodb.ListTablesInput, optFns ...func(*dynamodb.Options)) (*dynamodb.ListTablesOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListTables")
	}

	var r0 *dynamodb.ListTablesOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *dynamodb.ListTablesInput, ...func(*dynamodb.Options)) (*dynamodb.ListTablesOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *dynamodb.ListTablesInput, ...func(*dynamodb.Options)) *dynamodb.ListTablesOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dynamodb.ListTablesOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *dynamodb.ListTablesInput, ...func(*dynamodb.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockdynamoDBAPIV2_ListTables_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListTables'
type MockdynamoDBAPIV2_ListTables_Call struct {
	*mock.Call
}

// ListTables is a helper method to define mock.On call
//   - ctx context.Context
//   - params *dynamodb.ListTablesInput
//   - optFns ...func(*dynamodb.Options)
func (_e *MockdynamoDBAPIV2_Expecter) ListTables(ctx interface{}, params interface{}, optFns ...interface{}) *MockdynamoDBAPIV2_ListTables_Call {
	return &MockdynamoDBAPIV2_ListTables_Call{Call: _e.mock.On("ListTables",
		append([]interface{}{ctx, params}, optFns...)...)}
}

func (_c *MockdynamoDBAPIV2_ListTables_Call) Run(run func(ctx context.Context, params *dynamodb.ListTablesInput, optFns ...func(*dynamodb.Options))) *MockdynamoDBAPIV2_ListTables_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]func(*dynamodb.Options), len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(func(*dynamodb.Options))
			}
		}
		run(args[0].(context.Context), args[1].(*dynamodb.ListTablesInput), variadicArgs...)
	})
	return _c
}

func (_c *MockdynamoDBAPIV2_ListTables_Call) Return(_a0 *dynamodb.ListTablesOutput, _a1 error) *MockdynamoDBAPIV2_ListTables_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockdynamoDBAPIV2_ListTables_Call) RunAndReturn(run func(context.Context, *dynamodb.ListTablesInput, ...func(*dynamodb.Options)) (*dynamodb.ListTablesOutput, error)) *MockdynamoDBAPIV2_ListTables_Call {
	_c.Call.Return(run)
	return _c
}

// Scan provides a mock function with given fields: ctx, params, optFns
func (_m *MockdynamoDBAPIV2) Scan(ctx context.Context, params *dynamodb.ScanInput, optFns ...func(*dynamodb.Options)) (*dynamodb.ScanOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Scan")
	}

	var r0 *dynamodb.ScanOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *dynamodb.ScanInput, ...func(*dynamodb.Options)) (*dynamodb.ScanOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *dynamodb.ScanInput, ...func(*dynamodb.Options)) *dynamodb.ScanOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dynamodb.ScanOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *dynamodb.ScanInput, ...func(*dynamodb.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockdynamoDBAPIV2_Scan_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Scan'
type MockdynamoDBAPIV2_Scan_Call struct {
	*mock.Call
}

// Scan is a helper method to define mock.On call
//   - ctx context.Context
//   - params *dynamodb.ScanInput
//   - optFns ...func(*dynamodb.Options)
func (_e *MockdynamoDBAPIV2_Expecter) Scan(ctx interface{}, params interface{}, optFns ...interface{}) *MockdynamoDBAPIV2_Scan_Call {
	return &MockdynamoDBAPIV2_Scan_Call{Call: _e.mock.On("Scan",
		append([]interface{}{ctx, params}, optFns...)...)}
}

func (_c *MockdynamoDBAPIV2_Scan_Call) Run(run func(ctx context.Context, params *dynamodb.ScanInput, optFns ...func(*dynamodb.Options))) *MockdynamoDBAPIV2_Scan_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]func(*dynamodb.Options), len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(func(*dynamodb.Options))
			}
		}
		run(args[0].(context.Context), args[1].(*dynamodb.ScanInput), variadicArgs...)
	})
	return _c
}

func (_c *MockdynamoDBAPIV2_Scan_Call) Return(_a0 *dynamodb.ScanOutput, _a1 error) *MockdynamoDBAPIV2_Scan_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockdynamoDBAPIV2_Scan_Call) RunAndReturn(run func(context.Context, *dynamodb.ScanInput, ...func(*dynamodb.Options)) (*dynamodb.ScanOutput, error)) *MockdynamoDBAPIV2_Scan_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockdynamoDBAPIV2 creates a new instance of MockdynamoDBAPIV2. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockdynamoDBAPIV2(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockdynamoDBAPIV2 {
	mock := &MockdynamoDBAPIV2{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
