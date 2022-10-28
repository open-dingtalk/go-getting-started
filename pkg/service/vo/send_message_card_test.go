package vo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

/**
 * @Author linya.jj
 * @Date 2022/10/21 10:12 AM
 */

func TestMessageCardSendRequestVO_CheckValid_IfOK(t *testing.T) {
	req := &ApiRequestBaseVO{
		CorpId:             "corpId001",
		OpenConversationId: "openConversationId001",
	}

	card := &MessageCardSendRequestVO{
		*req,
	}

	assert.Nil(t, card.CheckValid())
}

func TestMessageCardSendRequestVO_CheckValid_IfNil(t *testing.T) {
	var req *MessageCardSendRequestVO

	assert.NotNil(t, req.CheckValid())

	card := &MessageCardSendRequestVO{
		ApiRequestBaseVO: ApiRequestBaseVO{
			OpenConversationId: "",
		},
	}

	assert.NotNil(t, card.CheckValid())
}
