package vo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

/**
 * @Author linya.jj
 * @Date 2022/10/21 10:14 AM
 */

func TestGetUserInfoRequestVO_CheckValid(t *testing.T) {

	var vo *GetUserInfoRequestVO
	assert.NotNil(t, vo.CheckValid())

	vo = &GetUserInfoRequestVO{
		RequestAuthCode: "",
	}
	assert.NotNil(t, vo.CheckValid())

	vo.RequestAuthCode = "code"
	assert.Nil(t, vo.CheckValid())
}
