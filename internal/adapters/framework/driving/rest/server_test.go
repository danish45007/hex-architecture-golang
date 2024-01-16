package rest_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/danish45007/hex-architecture-golang/internal/adapters/framework/driving/rest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockAPIPort is a mock implementation of the APIPort interface.
type MockAPIPort struct {
	mock.Mock
}

// GetAddition is the mock implementation for the GetAddition method.
func (m *MockAPIPort) GetAddition(a, b int32) (int32, error) {
	// Record the method call for expectations
    args := m.Called(a, b)
    
    // Retrieve the results from arguments
    res := args.Get(0).(int32)
    return res, args.Error(1)
}

// GetSubtraction is the mock implementation for the GetSubtraction method.
func (m *MockAPIPort) GetSubtraction(a, b int32) (int32, error) {
	// Record the method call for expectations
    args := m.Called(a, b)
    
    // Retrieve the results from arguments
    res := args.Get(0).(int32)
    return res, args.Error(1)
}

// GetMultiplication is the mock implementation for the GetMultiplication method.
func (m *MockAPIPort) GetMultiplication(a, b int32) (int32, error) {
	// Record the method call for expectations
    args := m.Called(a, b)
    
    // Retrieve the results from arguments
    res := args.Get(0).(int32)
    return res, args.Error(1)
}

// GetDivision is the mock implementation for the GetDivision method.
func (m *MockAPIPort) GetDivision(a, b int32) (int32, error) {
	// Record the method call for expectations
    args := m.Called(a, b)
    
    // Retrieve the results from arguments
    res := args.Get(0).(int32)
    return res, args.Error(1)
}

func TestHandleAddition(t *testing.T) {
	mockAPI := &MockAPIPort{}
	adapter := rest.NewAdapter(mockAPI)

	// Mock API behavior
	mockAPI.On("GetAddition", int32(2), int32(3)).Return(int32(5), nil)

	// Create a request
	req := httptest.NewRequest("GET", "/addition?a=2&b=3", nil)
	w := httptest.NewRecorder()

	// Call the handler
	adapter.GetAddition(w, req)

	// Check the response
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "Result of addition: 5", w.Body.String())

	// Assert that the mocked API method was called
	mockAPI.AssertExpectations(t)
}

func TestHandleSubtraction(t *testing.T) {
	mockAPI := &MockAPIPort{}
	adapter := rest.NewAdapter(mockAPI)

	// Mock API behavior
	mockAPI.On("GetSubtraction", int32(3), int32(2)).Return(int32(1), nil)

	// Create a request
	req := httptest.NewRequest("GET", "/subtraction?a=3&b=2", nil)
	w := httptest.NewRecorder()

	// Call the handler
	adapter.GetSubtraction(w, req)

	// Check the response
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "Result of subtraction: 1", w.Body.String())

	// Assert that the mocked API method was called
	mockAPI.AssertExpectations(t)
}

func TestHandleMultiplication(t *testing.T) {
	mockAPI := &MockAPIPort{}
	adapter := rest.NewAdapter(mockAPI)

	// Mock API behavior
	mockAPI.On("GetMultiplication", int32(3), int32(2)).Return(int32(6), nil)

	// Create a request
	req := httptest.NewRequest("GET", "/multiplication?a=3&b=2", nil)
	w := httptest.NewRecorder()

	// Call the handler
	adapter.GetMultiplication(w, req)

	// Check the response
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "Result of multiplication: 6", w.Body.String())

	// Assert that the mocked API method was called
	mockAPI.AssertExpectations(t)
}

func TestHandleDivision(t *testing.T) {
	mockAPI := &MockAPIPort{}
	adapter := rest.NewAdapter(mockAPI)

	// Mock API behavior
	mockAPI.On("GetDivision", int32(4), int32(2)).Return(int32(2), nil)

	// Create a request
	req := httptest.NewRequest("GET", "/subtraction?a=4&b=2", nil)
	w := httptest.NewRecorder()

	// Call the handler
	adapter.GetDivision(w, req)

	// Check the response
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "Result of division: 2", w.Body.String())

	// Assert that the mocked API method was called
	mockAPI.AssertExpectations(t)
}