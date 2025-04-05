package apigo

import (
	"errors"
	"net/http"
	"testing"
)

func TestIs(t *testing.T) {
	err := &Err{
		Code:  http.StatusNotFound,
		Cause: "404",
	}
	if !ErrIsNotFound(err) {
		t.Fail()
	}
	if ErrIsConflict(err) {
		t.Fail()
	}
}

func TestErrorsIs(t *testing.T) {
	err := &Err{
		Code:  http.StatusNotFound,
		Cause: "404",
	}
	if errors.Is(err, errors.New("fail")) {
		t.Fail()
	}
	if errors.Is(err, &Err{
		Code: http.StatusConflict,
	}) {
		t.Fail()
	}
	if !errors.Is(err, &Err{
		Code: http.StatusNotFound,
	}) {
		t.Fail()
	}
}
