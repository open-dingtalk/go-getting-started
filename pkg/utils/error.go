package utils

import "encoding/json"

/**
 * @Author linya.jj
 * @Date 2022/10/21 10:42 AM
 */

// ApiErrorTuple 错误码五元组定义
type ApiErrorTuple struct {
	Code             string `json:"code"`
	Reason           string `json:"reason"`
	DeveloperMessage string `json:"developerMessage"`
	MoreInfo         string `json:"moreInfo"`
	Scope            string `json:"scope"`
}

// Error 实现error interface
func (e *ApiErrorTuple) Error() string {
	jsonB, err := json.Marshal(e)
	if err != nil {
		return ""
	}

	return string(jsonB)
}

// NewApiErrorTuple 五元组错误码创建函数
func NewApiErrorTuple(code, reason, developrtMessage, moreInfo, scope string) *ApiErrorTuple {
	return &ApiErrorTuple{
		Code:             code,
		Reason:           reason,
		DeveloperMessage: developrtMessage,
		MoreInfo:         moreInfo,
		Scope:            scope,
	}
}

// NewInternalServerErrorTuple 服务器内部错误翻译成api错误
func NewInternalServerErrorTuple(err error) *ApiErrorTuple {
	if err == nil {
		return nil
	}

	return &ApiErrorTuple{
		Code:             APIErrorCodeKInternalServerError,
		Reason:           APIErrorInfoKInternalServerError,
		DeveloperMessage: err.Error(),
		MoreInfo:         "",
		Scope:            DefaultErrorScope,
	}
}
