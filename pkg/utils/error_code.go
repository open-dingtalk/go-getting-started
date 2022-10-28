package utils

/**
 * @Author linya.jj
 * @Date 2022/10/21 10:49 AM
 */

var (
	// APIErrorCodeKInternalServerError 服务内部错误
	APIErrorCodeKInternalServerError = "500000"

	// APIErrorCodeKParamInvalid 参数错误
	APIErrorCodeKParamInvalid = "400001"
	// APIErrorCodeKRequestBodyEmpty 请求参数为空
	APIErrorCodeKRequestBodyEmpty = "400002"
)

var (
	// APIErrorInfoKInternalServerError 服务内部错误
	APIErrorInfoKInternalServerError = "error.api.internal_server_error"
	// APIErrorInfoKParamInvalid 参数错误
	APIErrorInfoKParamInvalid = "error.api.param_invalid"
	// APIErrorInfoKRequestBodyEmpty 请求参数为空
	APIErrorInfoKRequestBodyEmpty = "error.api.request_body_empty"
)

var (
	DefaultErrorScope = "@go-getting-started"
)
