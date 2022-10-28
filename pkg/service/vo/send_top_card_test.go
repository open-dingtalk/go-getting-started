package vo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

/**
 * @Author linya.jj
 * @Date 2022/10/21 10:12 AM
 */

func TestTopCardSendRequestVO_CheckValid_IfOK(t *testing.T) {
	req := &ApiRequestBaseVO{
		CorpId:             "corpId001",
		OpenConversationId: "openConversationId001",
	}

	card := &TopCardSendRequestVO{
		*req,
	}

	assert.Nil(t, card.CheckValid())
}

func TestTopCardSendRequestVO_CheckValid_IfNil(t *testing.T) {
	var req *TopCardSendRequestVO

	assert.NotNil(t, req.CheckValid())

	card := &TopCardSendRequestVO{
		ApiRequestBaseVO: ApiRequestBaseVO{
			OpenConversationId: "",
		},
	}

	assert.NotNil(t, card.CheckValid())
}
