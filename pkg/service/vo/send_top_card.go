package vo

import (
	"github.com/open-dingtalk/go-getting-started/pkg/utils"
)

/**
 * @Author linya.jj
 * @Date 2022/10/21 10:12 AM
 */

type TopCardSendRequestVO struct {
	ApiRequestBaseVO
}

func (vo *TopCardSendRequestVO) CheckValid() error {
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

type TopCardSendResponseVO struct {
}
