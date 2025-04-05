package apigo

import (
	"errors"
	"fmt"
	"net/http"
)

type Err struct {
	Code  int
	Cause string
	Err   error
}

func (o *Err) Is(target error) bool {
	if t, ok := target.(*Err); ok {
		return t.Code == o.Code
	}
	return false
}

func (o *Err) Unwrap() error {
	return o.Err
}

func (o *Err) Error() string {
	return fmt.Sprintf("%d; %s", o.Code, o.Cause)
}

func errIs(err error, httpStatus int) bool {
	var webErr *Err
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

func ErrBadRequest(cause string) *Err {
	return &Err{
		Code:  http.StatusBadRequest,
		Cause: cause,
	}
}

// 404: NotFound

func ErrIsNotFound(err error) bool {
	return errIs(err, http.StatusNotFound)
}

func ErrNotFound(cause string) *Err {
	return &Err{
		Code:  http.StatusNotFound,
		Cause: cause,
	}
}

// 409: Conflict

func ErrIsConflict(err error) bool {
	return errIs(err, http.StatusConflict)
}

func ErrConflict(cause string) *Err {
	return &Err{
		Code:  http.StatusConflict,
		Cause: cause,
	}
}

// 422: UnprocessableEntity

func ErrIsUnprocessableEntity(err error) bool {
	return errIs(err, http.StatusUnprocessableEntity)
}

func ErrUnprocessableEntity(cause string) *Err {
	return &Err{
		Code:  http.StatusUnprocessableEntity,
		Cause: cause,
	}
}

// 429: TooManyRequests

func ErrIsTooManyRequests(err error) bool {
	return errIs(err, http.StatusTooManyRequests)
}

func ErrTooManyRequests(cause string) *Err {
	return &Err{
		Code:  http.StatusTooManyRequests,
		Cause: cause,
	}
}

// 401: Unauthorized

func ErrIsUnauthorized(err error) bool {
	return errIs(err, http.StatusUnauthorized)
}

func ErrUnauthorized(cause string) *Err {
	return &Err{
		Code:  http.StatusUnauthorized,
		Cause: cause,
	}
}

// 403; Forbidden

func ErrIsForbidden(err error) bool {
	return errIs(err, http.StatusForbidden)
}

func ErrForbidden(cause string) *Err {
	return &Err{
		Code:  http.StatusForbidden,
		Cause: cause,
	}
}
