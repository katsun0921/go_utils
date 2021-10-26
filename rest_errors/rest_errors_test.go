package rest_errors

import (
  "errors"
  "fmt"
  "github.com/stretchr/testify/assert"
  "net/http"
  "testing"
)


func TestNewRestError(t *testing.T) {
  var TestInterface []interface{}
  err := NewRestError("TestNewRestError", 404, "err", TestInterface)
  assert.NotNil(t, err)
}

func TestNewRestErrorFromBytes(t *testing.T) {
  b := []byte("test")
  err, _ := NewRestErrorFromBytes(b)
  assert.Nil(t, err)
}

func TestNewBadRequestError(t *testing.T) {
  TestErrorMessage := "this is the message"
  TestErrorStatus := http.StatusBadRequest
  TestErrorNew := "server error"
  TestError := fmt.Sprintf("message: %s - status: %d - error: bad request - causes: [%s]",TestErrorMessage, TestErrorStatus, TestErrorNew)
  err := NewBadRequestError("this is the message", errors.New("server error"))
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusBadRequest, err.Status())
	assert.EqualValues(t, "this is the message", err.Message())
	assert.EqualValues(t, TestError, err.Error())

	assert.NotNil(t, err.Causes)
	assert.EqualValues(t, 1, len(err.Causes()))
	assert.EqualValues(t, TestErrorNew, err.Causes()[0])
}

func TestNewNotFoundError(t *testing.T) {
  TestErrorMessage := "this is the message"
  TestErrorStatus := http.StatusNotFound
  TestErrorNew := "database error"
  TestError := fmt.Sprintf("message: %s - status: %d - error: not found - causes: [%s]",TestErrorMessage, TestErrorStatus, TestErrorNew)
  err := NewNotFoundError(TestErrorMessage, errors.New("database error"))
	assert.NotNil(t, err)
	assert.EqualValues(t, TestErrorStatus, err.Status())
	assert.EqualValues(t, TestErrorMessage, err.Message())
	assert.EqualValues(t, TestError, err.Error())

	assert.NotNil(t, err.Causes)
	assert.EqualValues(t, 1, len(err.Causes()))
	assert.EqualValues(t, TestErrorNew, err.Causes()[0])
}

func TestNewInternalServerError(t *testing.T) {
  TestErrorMessage := "this is the message"
  TestErrorStatus := http.StatusInternalServerError
  TestErrorNew := "database error"
  TestError := fmt.Sprintf("message: %s - status: %d - error: internal server error - causes: [%s]",TestErrorMessage, TestErrorStatus, TestErrorNew)
	err := NewInternalServerError(TestErrorMessage, errors.New(TestErrorNew))
	assert.NotNil(t, err)
	assert.EqualValues(t, TestErrorStatus, err.Status())
	assert.EqualValues(t, TestErrorMessage, err.Message())
	assert.EqualValues(t, TestError, err.Error())

	assert.NotNil(t, err.Causes)
	assert.EqualValues(t, 1, len(err.Causes()))
	assert.EqualValues(t, "database error", err.Causes()[0])

	fmt.Println(err.Error())
}
