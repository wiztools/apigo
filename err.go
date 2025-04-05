package apigo

import (
	"errors"
	"fmt"
	"net/http"
)

type ApiErr struct {
	Code  int
	Cause string
	Err   error
}

func (o *ApiErr) Is(target error) bool {
	if t, ok := target.(*ApiErr); ok {
		return t.Code == o.Code
	}
	return false
}

func (o *ApiErr) Unwrap() error {
	return o.Err
}

func (o *ApiErr) Error() string {
	return fmt.Sprintf("%d; %s", o.Code, o.Cause)
}

func errIs(err error, httpStatus int) bool {
	var webErr *ApiErr
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

func ErrBadRequest(cause string) *ApiErr {
	return &ApiErr{
		Code:  http.StatusBadRequest,
		Cause: cause,
	}
}

// 404: NotFound

func ErrIsNotFound(err error) bool {
	return errIs(err, http.StatusNotFound)
}

func ErrNotFound(cause string) *ApiErr {
	return &ApiErr{
		Code:  http.StatusNotFound,
		Cause: cause,
	}
}

// 409: Conflict

func ErrIsConflict(err error) bool {
	return errIs(err, http.StatusConflict)
}

func ErrConflict(cause string) *ApiErr {
	return &ApiErr{
		Code:  http.StatusConflict,
		Cause: cause,
	}
}

// 422: UnprocessableEntity

func ErrIsUnprocessableEntity(err error) bool {
	return errIs(err, http.StatusUnprocessableEntity)
}

func ErrUnprocessableEntity(cause string) *ApiErr {
	return &ApiErr{
		Code:  http.StatusUnprocessableEntity,
		Cause: cause,
	}
}

// 429: TooManyRequests

func ErrIsTooManyRequests(err error) bool {
	return errIs(err, http.StatusTooManyRequests)
}

func ErrTooManyRequests(cause string) *ApiErr {
	return &ApiErr{
		Code:  http.StatusTooManyRequests,
		Cause: cause,
	}
}

// 401: Unauthorized

func ErrIsUnauthorized(err error) bool {
	return errIs(err, http.StatusUnauthorized)
}

func ErrUnauthorized(cause string) *ApiErr {
	return &ApiErr{
		Code:  http.StatusUnauthorized,
		Cause: cause,
	}
}

// 403; Forbidden

func ErrIsForbidden(err error) bool {
	return errIs(err, http.StatusForbidden)
}

func ErrForbidden(cause string) *ApiErr {
	return &ApiErr{
		Code:  http.StatusForbidden,
		Cause: cause,
	}
}
