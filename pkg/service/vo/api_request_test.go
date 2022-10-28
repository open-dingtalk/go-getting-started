package vo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

/**
 * @Author linya.jj
 * @Date 2022/10/21 10:12 AM
 */

func TestApiRequestBaseVO_CheckValid_IfOK(t *testing.T) {
	req := &ApiRequestBaseVO{
		CorpId:             "corpId001",
		OpenConversationId: "openConversationId001",
	}

	assert.Nil(t, req.CheckValid())
}

func TestApiRequestBaseVO_CheckValid_IfNil(t *testing.T) {
	var req *ApiRequestBaseVO

	assert.NotNil(t, req.CheckValid())
}

func TestApiRequestBaseVO_CheckValid_IfOpenConversationInvalid(t *testing.T) {
	req := &ApiRequestBaseVO{
		CorpId:             "corpId001",
		OpenConversationId: "",
	}

	assert.NotNil(t, req.CheckValid())
}

func TestApiRequestBaseVO_GetCorpId_IfNil(t *testing.T) {
	var req *ApiRequestBaseVO

	assert.Equal(t, "", req.GetCorpId())
}

func TestApiRequestBaseVO_GetCorpId_IfNotNil(t *testing.T) {
	req := &ApiRequestBaseVO{
		CorpId:             "corpId001",
		OpenConversationId: "",
	}

	assert.Equal(t, "corpId001", req.GetCorpId())
}

func TestApiRequestBaseVO_GetOpenConversationId_IfNil(t *testing.T) {
	var req *ApiRequestBaseVO

	assert.Equal(t, "", req.GetOpenConversationId())
}

func TestApiRequestBaseVO_GetOpenConversationId_IfNotNil(t *testing.T) {
	req := &ApiRequestBaseVO{
		CorpId:             "",
		OpenConversationId: "openConversationId001",
	}

	assert.Equal(t, "openConversationId001", req.GetOpenConversationId())
}
