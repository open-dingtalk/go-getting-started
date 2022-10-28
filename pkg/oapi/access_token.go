package oapi

import (
	"context"
	"github.com/open-dingtalk/go-getting-started/pkg/config"
	"github.com/open-dingtalk/go-getting-started/pkg/oapi/oapiidl"
	"time"
)

/**
 * @Author linya.jj
 * @Date 2022/10/9 10:41 AM
 */

// AccessTokenModel access tokens
type AccessTokenModel struct {
	AccessToken string
	ExpireAt    int64
}

// Valid 是否过期，保持5分钟刷新
func (m *AccessTokenModel) Valid() bool {
	if m == nil || m.AccessToken == "" {
		return false
	}

	if time.Now().UnixMilli() > m.ExpireAt-5*60*1000 {
		return false
	}

	return true
}

var (
	accessTokenModel *AccessTokenModel
)

func GetAccessToken(ctx context.Context, oapiClient *DTOAPIClient) (*AccessTokenModel, error) {
	if accessTokenModel.Valid() {
		return accessTokenModel, nil
	}

	//TODO accessToken应该异步刷新，防止大并发下的瞬时频繁请求问题

	accessTokenRequestModel := &oapiidl.AccessTokenRequestModel{
		AppKey:    config.GetConfig().GetInnerAppInfo().GetAppKey(),
		AppSecret: config.GetConfig().GetInnerAppInfo().GetAppSecret(),
	}

	accessTokenResponseModel := &oapiidl.AccessTokenResponseModel{}

	err := oapiClient.RequestV2(ctx, "POST", "/v1.0/oauth2/accessToken",
		accessTokenRequestModel, accessTokenResponseModel)

	if err != nil {
		return nil, err
	}

	accessTokenModel = &AccessTokenModel{
		accessTokenResponseModel.AccessToken,
		time.Now().UnixMilli() + accessTokenResponseModel.ExpireIn*1000,
	}

	return accessTokenModel, nil
}
