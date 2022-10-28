package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/open-dingtalk/go-getting-started/pkg/manager"
	"github.com/open-dingtalk/go-getting-started/pkg/oapi"
	"github.com/open-dingtalk/go-getting-started/pkg/oapi/oapiidl"
	"github.com/open-dingtalk/go-getting-started/pkg/utils"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

/**
 * @Author linya.jj
 * @Date 2022/10/19 11:14 AM
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
	}

	return &http.Response{
		StatusCode: 200,
		Body:       ioutil.NopCloser(bytes.NewBuffer([]byte("{}"))),
	}, nil
}

func TestCoolAppService_SendTextMessage(t *testing.T) {
}

func TestCoolAppService_SendMessageCard(t *testing.T) {
}

func TestCoolAppService_SendTopCard(t *testing.T) {
}

func TestCoolAppService_getRequestBodyAndCheckValid(t *testing.T) {
}

func TestCoolAppService_writeResponseError(t *testing.T) {
	srv := NewCoolAppService(manager.NewCoolAppManager(oapi.NewDTOAPIClient(&MockHttpClient{})))
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	srv.writeResponseError(ctx, fmt.Errorf("error001"))
	srv.writeResponseError(ctx, utils.NewApiErrorTuple("", "", "", "", ""))
	srv.writeResponseError(ctx, nil)
}

func TestCoolAppService_writeResponseVO(t *testing.T) {
	srv := NewCoolAppService(manager.NewCoolAppManager(oapi.NewDTOAPIClient(&MockHttpClient{})))
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	srv.writeSuccessResponseVO(ctx, "application/json", map[string]string{})
}
