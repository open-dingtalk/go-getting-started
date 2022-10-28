package utils

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

/**
 * @Author linya.jj
 * @Date 2022/10/21 10:42 AM
 */

func TestApiErrorTuple_Error_IfNotNil(t *testing.T) {
	apiErr := &ApiErrorTuple{
		Code:             "code001",
		Reason:           "eeason001",
		DeveloperMessage: "developerMessage001",
		MoreInfo:         "moreInfo001",
		Scope:            "scope001",
	}

	apiErr1 := &ApiErrorTuple{}
	err := json.Unmarshal([]byte(apiErr.Error()), apiErr1)
	assert.Nil(t, err)
	assert.Equal(t, apiErr.Code, apiErr1.Code)
	assert.Equal(t, apiErr.Reason, apiErr1.Reason)
	assert.Equal(t, apiErr.DeveloperMessage, apiErr1.DeveloperMessage)
	assert.Equal(t, apiErr.MoreInfo, apiErr1.MoreInfo)
	assert.Equal(t, apiErr.Scope, apiErr1.Scope)
}

func TestApiErrorTuple_Error_IfNil(t *testing.T) {
	var apiErr *ApiErrorTuple
	assert.Equal(t, "null", apiErr.Error())
}

func TestNewApiErrorTuple(t *testing.T) {
	apiErr := &ApiErrorTuple{
		Code:             "code001",
		Reason:           "eeason001",
		DeveloperMessage: "developerMessage001",
		MoreInfo:         "moreInfo001",
		Scope:            "scope001",
	}

	apiErr1 := NewApiErrorTuple(apiErr.Code, apiErr.Reason, apiErr.DeveloperMessage, apiErr.MoreInfo, apiErr.Scope)

	assert.Equal(t, apiErr.Code, apiErr1.Code)
	assert.Equal(t, apiErr.Reason, apiErr1.Reason)
	assert.Equal(t, apiErr.DeveloperMessage, apiErr1.DeveloperMessage)
	assert.Equal(t, apiErr.MoreInfo, apiErr1.MoreInfo)
	assert.Equal(t, apiErr.Scope, apiErr1.Scope)
}

func TestNewInternalServerErrorTuple_IfNotNil(t *testing.T) {
	err := fmt.Errorf("error001")
	serverError := NewInternalServerErrorTuple(err)

	assert.Equal(t, APIErrorCodeKInternalServerError, serverError.Code)
	assert.Equal(t, APIErrorInfoKInternalServerError, serverError.Reason)
	assert.Equal(t, "error001", serverError.DeveloperMessage)
	assert.Equal(t, "", serverError.MoreInfo)
	assert.Equal(t, DefaultErrorScope, serverError.Scope)
}

func TestNewInternalServerErrorTuple_IfNil(t *testing.T) {
	assert.Nil(t, NewInternalServerErrorTuple(nil))
}
