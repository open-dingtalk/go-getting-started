package manager

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/open-dingtalk/go-getting-started/pkg/oapi"
	"github.com/open-dingtalk/go-getting-started/pkg/oapi/oapiidl"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

/**
 * @Author linya.jj
 * @Date 2022/10/19 11:09 AM
 */

type MockHttpClient struct {
}

func (c *MockHttpClient) Do(req *http.Request) (*http.Response, error) {
	if strings.Contains(req.URL.Path, "/v1.0/oauth2/accessToken") {
		m := oapiidl.AccessTokenResponseModel{
			AccessToken: "accesstoken001",
			ExpireIn:    1000 * 1000,
		}
		mb, _ := json.Marshal(m)
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBuffer(mb)),
		}, nil
	} else if strings.Contains(req.URL.Path, "/topapi/v2/user/getuserinfo") {
		errMsg := ""
		errCode := 0
		userId := "userId001"
		m := oapiidl.UserGetUserInfoResponseModel{
			ErrCode: &errCode,
			ErrMsg:  &errMsg,
			Result: &oapiidl.UserGetUserInfoResultModel{
				Userid: &userId,
			},
		}
		mb, _ := json.Marshal(m)
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBuffer(mb)),
		}, nil
	} else if strings.Contains(req.URL.Path, "/topapi/v2/user/get") {
		errMsg := ""
		errCode := 0
		userId := "userId001"
		avatar := "avatar001"
		name := "name001"
		m := oapiidl.UserGetResponseModel{
			ErrCode: &errCode,
			ErrMsg:  &errMsg,
			Result: &oapiidl.UserGetResultModel{
				Userid: &userId,
				Avatar: &avatar,
				Name:   &name,
			},
		}
		mb, _ := json.Marshal(m)
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBuffer(mb)),
		}, nil
	}

	return &http.Response{
		StatusCode: 200,
		Body:       ioutil.NopCloser(bytes.NewBuffer([]byte("{}"))),
	}, nil
}

func TestCoolAppManager_SendTextMessage(t *testing.T) {
	err := NewCoolAppManager(oapi.NewDTOAPIClient(&MockHttpClient{})).SendTextMessage(context.Background(), "openconversation001", "text")
	assert.Nil(t, err)
}

func TestCoolAppManager_SendMessageCard(t *testing.T) {
	err := NewCoolAppManager(oapi.NewDTOAPIClient(&MockHttpClient{})).SendMessageCard(context.Background(), "openconversation001")
	assert.Nil(t, err)
}

func TestCoolAppManager_SendTopCard(t *testing.T) {
	err := NewCoolAppManager(oapi.NewDTOAPIClient(&MockHttpClient{})).SendTopCard(context.Background(), "openconversation001")
	assert.Nil(t, err)
}

func TestCoolAppManager_GetUserInfo(t *testing.T) {
	_, err := NewCoolAppManager(oapi.NewDTOAPIClient(&MockHttpClient{})).GetUserInfo(context.Background(), "requestAuthCode")
	assert.Nil(t, err)
}
