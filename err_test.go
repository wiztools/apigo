package apigo

import (
	"errors"
	"net/http"
	"testing"
)

func TestIs(t *testing.T) {
	err := &ApiErr{
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
	err := &ApiErr{
		Code:  http.StatusNotFound,
		Cause: "404",
	}
	if errors.Is(err, errors.New("fail")) {
		t.Fail()
	}
	if errors.Is(err, &ApiErr{
		Code: http.StatusConflict,
	}) {
		t.Fail()
	}
	if !errors.Is(err, &ApiErr{
		Code: http.StatusNotFound,
	}) {
		t.Fail()
	}
}
