// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package app

import (
	"fmt"
	"net/http"

	apiv1 "github.com/google/cloud-android-orchestration/api/v1"
)

type AppError struct {
	Msg        string
	StatusCode int
	Err        error
}

func (e *AppError) Error() string {
	if e.Err != nil {
		return e.Msg + ": " + e.Err.Error()
	}
	return e.Msg
}
func (e *AppError) Unwrap() error {
	return e.Err
}

func (e *AppError) JSONResponse() apiv1.Error {
	// Include only the high level error message in the error response, the
	// lower level errors are just for logging
	return apiv1.Error{
		Code:     e.StatusCode,
		ErrorMsg: e.Msg,
	}
}

func NewNotFoundError(msg string, e error) error {
	return &AppError{Msg: msg, StatusCode: http.StatusNotFound, Err: e}
}

func NewBadRequestError(msg string, e error) error {
	return &AppError{Msg: msg, StatusCode: http.StatusBadRequest, Err: e}
}

func NewInvalidQueryParamError(param, value string, err error) error {
	return NewBadRequestError(fmt.Sprintf("Invalid query parameter %q value: %q", param, value), err)
}

func NewMethodNotAllowedError(msg string, e error) error {
	return &AppError{Msg: msg, StatusCode: http.StatusMethodNotAllowed, Err: e}
}

func NewInternalError(msg string, e error) error {
	return &AppError{Msg: msg, StatusCode: http.StatusInternalServerError, Err: e}
}

func NewForbiddenError(msg string, e error) error {
	return &AppError{Msg: msg, StatusCode: http.StatusForbidden, Err: e}
}

func NewServiceUnavailableError(msg string, e error) error {
	return &AppError{Msg: msg, StatusCode: http.StatusServiceUnavailable, Err: e}
}
