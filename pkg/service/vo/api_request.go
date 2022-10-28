package vo

import (
	"github.com/open-dingtalk/go-getting-started/pkg/utils"
)

/**
 * @Author linya.jj
 * @Date 2022/10/21 10:12 AM
 */

// IApiRequestVO 参数请求interface
type IApiRequestVO interface {
	CheckValid() error
}

// ApiRequestBaseVO 基础模板
type ApiRequestBaseVO struct {
	CorpId             string `json:"corpId"`
	OpenConversationId string `json:"openConversationId"`
}

// CheckValid 实现校验
func (vo *ApiRequestBaseVO) CheckValid() error {
	if vo == nil {
		return utils.NewApiErrorTuple(utils.APIErrorCodeKRequestBodyEmpty,
			utils.APIErrorInfoKRequestBodyEmpty,
			"request body empty",
			"corpId and openConversationId should be specified.",
			"")
	}

	if vo.OpenConversationId == "" {
		return utils.NewApiErrorTuple(utils.APIErrorCodeKParamInvalid,
			utils.APIErrorInfoKParamInvalid,
			"openConversationId empty",
			"openConversationId should be specified.",
			utils.DefaultErrorScope)
	}

	return nil
}

func (vo *ApiRequestBaseVO) GetCorpId() string {
	if vo == nil {
		return ""
	}

	return vo.CorpId
}

func (vo *ApiRequestBaseVO) GetOpenConversationId() string {
	if vo == nil {
		return ""
	}

	return vo.OpenConversationId
}
