package vo

import "github.com/open-dingtalk/go-getting-started/pkg/utils"

/**
 * @Author linya.jj
 * @Date 2022/10/21 10:14 AM
 */

type GetUserInfoRequestVO struct {
	RequestAuthCode string
}

func (vo *GetUserInfoRequestVO) CheckValid() error {
	if vo == nil || vo.RequestAuthCode == "" {
		return utils.NewApiErrorTuple(utils.APIErrorCodeKRequestBodyEmpty,
			utils.APIErrorInfoKRequestBodyEmpty,
			"RequestAuthCode empty",
			"",
			"")
	}

	return nil
}

type GetUserInfoResponseVO struct {
	Name   string `json:"name"`
	UserId string `json:"userId"`
	Avatar string `json:"avatar"`
}
