package apigo

import (
	"errors"
	"fmt"
	"net/http"
)

type WebErr struct {
	Code  int
	Cause string
	Err   error
}

func (o *WebErr) Is(target error) bool {
	if t, ok := target.(*WebErr); ok {
		return t.Code == o.Code
	}
	return false
}

func (o *WebErr) Unwrap() error {
	return o.Err
}

func (o *WebErr) Error() string {
	return fmt.Sprintf("%d; %s", o.Code, o.Cause)
}

func errIs(err error, httpStatus int) bool {
	var webErr *WebErr
	if errors.As(err, &webErr) {
		if webErr.Code == httpStatus {
			return true
		}
	}
	return false
}

// 400: BadRequest

func ErrIsBadRequest(err error) bool {
	return errIs(err, http.StatusBadRequest)
}

func ErrBadRequest(cause string) *WebErr {
	return &WebErr{
		Code:  http.StatusBadRequest,
		Cause: cause,
	}
}

// 404: NotFound

func ErrIsNotFound(err error) bool {
	return errIs(err, http.StatusNotFound)
}

func ErrNotFound(cause string) *WebErr {
	return &WebErr{
		Code:  http.StatusNotFound,
		Cause: cause,
	}
}

// 409: Conflict

func ErrIsConflict(err error) bool {
	return errIs(err, http.StatusConflict)
}

func ErrConflict(cause string) *WebErr {
	return &WebErr{
		Code:  http.StatusConflict,
		Cause: cause,
	}
}

// 422: UnprocessableEntity

func ErrIsUnprocessableEntity(err error) bool {
	return errIs(err, http.StatusUnprocessableEntity)
}

func ErrUnprocessableEntity(cause string) *WebErr {
	return &WebErr{
		Code:  http.StatusUnprocessableEntity,
		Cause: cause,
	}
}

// 429: TooManyRequests

func ErrIsTooManyRequests(err error) bool {
	return errIs(err, http.StatusTooManyRequests)
}

func ErrTooManyRequests(cause string) *WebErr {
	return &WebErr{
		Code:  http.StatusTooManyRequests,
		Cause: cause,
	}
}

// 401: Unauthorized

func ErrIsUnauthorized(err error) bool {
	return errIs(err, http.StatusUnauthorized)
}

func ErrUnauthorized(cause string) *WebErr {
	return &WebErr{
		Code:  http.StatusUnauthorized,
		Cause: cause,
	}
}

// 403; Forbidden

func ErrIsForbidden(err error) bool {
	return errIs(err, http.StatusForbidden)
}

func ErrForbidden(cause string) *WebErr {
	return &WebErr{
		Code:  http.StatusForbidden,
		Cause: cause,
	}
}
