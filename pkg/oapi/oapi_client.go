package oapi

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/open-dingtalk/go-getting-started/pkg/oapi/oapiidl"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

/**
 * @Author linya.jj
 * @Date 2022/10/9 10:31 AM
 */

const (
	// OAPIHost 钉钉OpenAPI网关地址
	OAPIHostV1 = "oapi.dingtalk.com"
	OAPIHost   = "api.dingtalk.com"
)

type IHttpClient interface {
	Do(*http.Request) (*http.Response, error)
}

type DTOAPIClient struct {
	httpClient IHttpClient
}

func NewDTOAPIClient(httpClient IHttpClient) *DTOAPIClient {
	if httpClient == nil {
		httpClient = &http.Client{
			Transport: http.DefaultTransport,
			Timeout:   5 * time.Second, //设置超时，包含connection时间、任意重定向时间、读取response body时间
		}
	}

	return &DTOAPIClient{
		httpClient: httpClient,
	}
}

// RequestV2 发起oapi请求,新版接口（api.dingtalk.com）
func (c *DTOAPIClient) RequestV2(ctx context.Context, method, apiPath string, requestModel, responseModel interface{}) error {
	if responseModel == nil {
		return errors.WithStack(fmt.Errorf("ResponseModelMustNotBeNil"))
	}

	u := &url.URL{}
	u.Scheme = "https"
	u.Host = OAPIHost
	u.Path = apiPath

	needAccessToken := !c.NotNeedAccessToken(apiPath)

	accessTokenStr := ""
	if needAccessToken {
		accessToken, err := GetAccessToken(ctx, c)
		if err != nil {
			return err
		}

		accessTokenStr = accessToken.AccessToken
	}

	var requestJsonReader io.Reader
	if requestModel != nil {
		requestJsonBody, err := json.Marshal(requestModel)
		if err != nil {
			return errors.WithStack(err)
		}

		requestJsonReader = bytes.NewReader(requestJsonBody)
	}

	req, err := http.NewRequestWithContext(ctx, method, u.String(), requestJsonReader)
	if err != nil {
		return errors.WithStack(err)
	}

	if needAccessToken {
		req.Header.Set("x-acs-dingtalk-access-token", accessTokenStr)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return errors.WithStack(err)
	}

	defer resp.Body.Close()
	responseJsonBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return errors.WithStack(err)
	}

	if resp.StatusCode != 200 {
		return errors.Wrap(fmt.Errorf("OAPIRequestError"),
			fmt.Sprintf("statusCode=%d,content=%s", resp.StatusCode, string(responseJsonBody)))
	}

	err = json.Unmarshal(responseJsonBody, responseModel)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

// RequestV1 发起oapi请求,老版接口（oapi.dingtalk.com）
func (c *DTOAPIClient) RequestV1(ctx context.Context, method, apiPath string, requestModel interface{}, responseModel oapiidl.IOAPIResponseModelV1) error {
	if responseModel == nil {
		return errors.WithStack(fmt.Errorf("ResponseModelMustNotBeNil"))
	}

	u := &url.URL{}
	u.Scheme = "https"
	u.Host = OAPIHostV1
	u.Path = apiPath

	needAccessToken := !c.NotNeedAccessToken(apiPath)

	accessTokenStr := ""
	if needAccessToken {
		accessToken, err := GetAccessToken(ctx, c)
		if err != nil {
			return err
		}

		accessTokenStr = accessToken.AccessToken
	}

	var requestJsonReader io.Reader
	if requestModel != nil {
		requestJsonBody, err := json.Marshal(requestModel)
		if err != nil {
			return errors.WithStack(err)
		}

		requestJsonReader = bytes.NewReader(requestJsonBody)
	}

	if needAccessToken {
		v := &url.Values{}
		v.Add("access_token", accessTokenStr)

		u.RawQuery = v.Encode()
	}

	req, err := http.NewRequestWithContext(ctx, method, u.String(), requestJsonReader)
	if err != nil {
		return errors.WithStack(err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return errors.WithStack(err)
	}

	defer resp.Body.Close()
	responseJsonBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return errors.WithStack(err)
	}

	// TODO 错误格式化
	if resp.StatusCode != 200 {
		return errors.Wrap(fmt.Errorf("OAPIRequestError"),
			fmt.Sprintf("statusCode=%d,content=%s", resp.StatusCode, string(responseJsonBody)))
	}

	err = json.Unmarshal(responseJsonBody, responseModel)
	if err != nil {
		return errors.WithStack(err)
	}

	if responseModel.GetErrCode() != 0 {
		return errors.Wrap(fmt.Errorf("OAPIRequestError"),
			fmt.Sprintf("errCode=%d,errmsg=%s", responseModel.GetErrCode(), responseModel.GetErrMsg()))
	}

	return nil
}

var (
	donotNeedAccessTokenApiWhiteList = map[string]bool{
		"/v1.0/oauth2/accessToken":     true,
		"/v1.0/oauth2/corpAccessToken": true,
	}
)

// NotNeedAccessToken 不需要accesstoken的接口
func (c *DTOAPIClient) NotNeedAccessToken(apiPath string) bool {
	if !strings.HasPrefix(apiPath, "/") {
		apiPath = "/" + apiPath
	}

	return donotNeedAccessTokenApiWhiteList[apiPath]
}
