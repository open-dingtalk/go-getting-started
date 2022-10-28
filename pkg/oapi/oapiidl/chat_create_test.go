package oapiidl

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

/**
 * @Author linya.jj
 * @Date 2022/10/9 2:31 PM
 */

func TestChatCreateResponseModel_GetErrMsg_IfNotNil(t *testing.T) {
	m := ChatCreateResponseModel{
		ErrMsg: "errmsg",
	}

	assert.Equal(t, "errmsg", m.GetErrMsg())
}

func TestChatCreateResponseModel_GetErrMsg_IfNil(t *testing.T) {
	var m *ChatCreateResponseModel

	assert.Equal(t, "", m.GetErrMsg())
}

func TestChatCreateResponseModel_GetErrCode_IfNotNil(t *testing.T) {
	m := ChatCreateResponseModel{
		ErrCode: 400,
	}

	assert.Equal(t, 400, m.GetErrCode())
}

func TestChatCreateResponseModel_GetErrCode_IfNil(t *testing.T) {
	var m *ChatCreateResponseModel

	assert.Equal(t, 0, m.GetErrCode())
}
