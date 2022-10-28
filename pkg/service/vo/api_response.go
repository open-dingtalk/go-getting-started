package vo

/**
 * @Author linya.jj
 * @Date 2022/10/23 1:53 PM
 */

type ApiResponseVO struct {
	Success bool        `json:"success"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewSuccessApiResponseVO(data interface{}) *ApiResponseVO {
	return &ApiResponseVO{
		Success: true,
		Code:    200,
		Message: "OK",
		Data:    data,
	}
}
