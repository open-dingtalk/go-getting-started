package oapiidl

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

/**
 * @Author linya.jj
 * @Date 2022/10/23 11:13 AM
 */

func TestUserGetResponseModel_GetErrMsg_IfNotNil(t *testing.T) {
	errmsg := "errmsg"
	m := UserGetResponseModel{
		ErrMsg: &errmsg,
	}

	assert.Equal(t, "errmsg", m.GetErrMsg())
}

func TestUserGetResponseModel_GetErrMsg_IfNil(t *testing.T) {
	var m *UserGetResponseModel

	assert.Equal(t, "", m.GetErrMsg())
}

func TestUserGetResponseModel_GetErrCode_IfNotNil(t *testing.T) {
	errCode := 400
	m := UserGetResponseModel{
		ErrCode: &errCode,
	}

	assert.Equal(t, 400, m.GetErrCode())
}

func TestUserGetResponseModel_GetErrCode_IfNil(t *testing.T) {
	var m *UserGetResponseModel

	assert.Equal(t, 0, m.GetErrCode())
}
