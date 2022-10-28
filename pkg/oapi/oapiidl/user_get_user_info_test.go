package oapiidl

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

/**
 * @Author linya.jj
 * @Date 2022/10/23 11:13 AM
 */

//https://open.dingtalk.com/document/orgapp-server/obtain-the-userid-of-a-user-by-using-the-log-free

func TestUserGetUserInfoResponseModel_GetErrMsg_IfNotNil(t *testing.T) {
	errmsg := "errmsg"
	m := UserGetUserInfoResponseModel{
		ErrMsg: &errmsg,
	}

	assert.Equal(t, "errmsg", m.GetErrMsg())
}

func TestUserGetUserInfoResponseModel_GetErrMsg_IfNil(t *testing.T) {
	var m *UserGetUserInfoResponseModel

	assert.Equal(t, "", m.GetErrMsg())
}

func TestUserGetUserInfoResponseModel_GetErrCode_IfNotNil(t *testing.T) {
	errCode := 400
	m := UserGetUserInfoResponseModel{
		ErrCode: &errCode,
	}

	assert.Equal(t, 400, m.GetErrCode())
}

func TestUserGetUserInfoResponseModel_GetErrCode_IfNil(t *testing.T) {
	var m *UserGetUserInfoResponseModel

	assert.Equal(t, 0, m.GetErrCode())
}
