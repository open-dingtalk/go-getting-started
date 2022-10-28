package oapiidl

/**
 * @Author linya.jj
 * @Date 2022/10/21 3:31 PM
 */

// AccessTokenRequestModel 获取accesstoken
type AccessTokenRequestModel struct {
	AppKey    string `json:"appKey"`
	AppSecret string `json:"appSecret"`
}

type AccessTokenResponseModel struct {
	AccessToken string `json:"accessToken"`
	ExpireIn    int64  `json:"expireIn"`
}
