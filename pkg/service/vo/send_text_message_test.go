package vo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

/**
 * @Author linya.jj
 * @Date 2022/10/21 10:12 AM
 */

func TestTextMessageSendRequestVO_CheckValid_IfOK(t *testing.T) {
	req := &ApiRequestBaseVO{
		CorpId:             "corpId001",
		OpenConversationId: "openConversationId001",
	}

	card := &TextMessageSendRequestVO{
		ApiRequestBaseVO: *req,
		Txt:              "txt001",
	}

	assert.Nil(t, card.CheckValid())
}

func TestTextMessageSendRequestVO_CheckValid_IfNil(t *testing.T) {
	var req *TextMessageSendRequestVO

	assert.NotNil(t, req.CheckValid())

	card := &TextMessageSendRequestVO{
		ApiRequestBaseVO: ApiRequestBaseVO{
			OpenConversationId: "",
		},
		Txt: "txt001",
	}

	assert.NotNil(t, card.CheckValid())
}
