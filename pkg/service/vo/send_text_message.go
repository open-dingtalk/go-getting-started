package vo

import (
	"github.com/open-dingtalk/go-getting-started/pkg/utils"
)

/**
 * @Author linya.jj
 * @Date 2022/10/21 10:12 AM
 */

type TextMessageSendRequestVO struct {
	ApiRequestBaseVO
	Txt string `json:"txt"`
}

func (vo *TextMessageSendRequestVO) CheckValid() error {
	if vo == nil {
		return utils.NewApiErrorTuple(utils.APIErrorCodeKRequestBodyEmpty,
			utils.APIErrorInfoKRequestBodyEmpty,
			"request body empty",
			"",
			"")
	}

	if err := vo.ApiRequestBaseVO.CheckValid(); err != nil {
		return err
	}

	return nil
}

type TextMessageSendResponseVO struct {
}
