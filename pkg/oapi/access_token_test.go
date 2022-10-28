package oapi

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

/**
 * @Author linya.jj
 * @Date 2022/10/9 10:41 AM
 */

func TestAccessTokenModel_Valid_IfNil(t *testing.T) {
	var m *AccessTokenModel
	assert.False(t, m.Valid())

	m = &AccessTokenModel{
		AccessToken: "",
	}
	assert.False(t, m.Valid())
}

func TestAccessTokenModel_Valid_IfOutDate(t *testing.T) {
	m := &AccessTokenModel{
		AccessToken: "accesstoken001",
		ExpireAt:    time.Now().UnixMilli() - 1000*60*1000,
	}

	assert.False(t, m.Valid())
}

func TestAccessTokenModel_Valid_IfOK(t *testing.T) {
	m := &AccessTokenModel{
		AccessToken: "accesstoken001",
		ExpireAt:    time.Now().UnixMilli() + 1000*60*1000,
	}

	assert.True(t, m.Valid())
}

func TestGetAccessToken_IfRequestError(t *testing.T) {
	_, err := GetAccessToken(context.Background(), NewDTOAPIClient(&Mock400HttpClient{}))
	assert.NotNil(t, err)
}

func TestGetAccessToken_IfOK(t *testing.T) {
	token, err := GetAccessToken(context.Background(), NewDTOAPIClient(&MockHttpClient{}))
	assert.Nil(t, err)
	assert.Equal(t, "accesstoken001", token.AccessToken)

	token, err = GetAccessToken(context.Background(), nil)
	assert.Nil(t, err)
	assert.Equal(t, "accesstoken001", token.AccessToken)
}
