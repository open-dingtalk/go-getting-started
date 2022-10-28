package oapiidl

/**
 * @Author linya.jj
 * @Date 2022/10/23 11:13 AM
 */

//https://open.dingtalk.com/document/orgapp-server/obtain-the-userid-of-a-user-by-using-the-log-free

type UserGetUserInfoRequestModel struct {
	Code *string `json:"code"`
}

type UserGetUserInfoResponseModel struct {
	RequestId *string                     `json:"requestId"`
	ErrMsg    *string                     `json:"errmsg"`
	ErrCode   *int                        `json:"errcode"`
	Result    *UserGetUserInfoResultModel `json:"result"`
}

type UserGetUserInfoResultModel struct {
	AssociatedUnionid *string `json:"associated_unionid"`
	UnionId           *string `json:"unionid"`
	DeviceId          *string `json:"device_id"`
	SysLevel          *int    `json:"sys_level"`
	Name              *string `json:"name"`
	Sys               *bool   `json:"sys"`
	Userid            *string `json:"userid"`
}

func (m *UserGetUserInfoResponseModel) GetErrMsg() string {
	if m == nil {
		return ""
	}
	return *m.ErrMsg
}

func (m *UserGetUserInfoResponseModel) GetErrCode() int {
	if m == nil {
		return 0
	}
	return *m.ErrCode
}
