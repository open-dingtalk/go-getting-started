package oapi

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/open-dingtalk/go-getting-started/pkg/oapi/oapiidl"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

/**
 * @Author linya.jj
 * @Date 2022/10/9 10:31 AM
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

type MockJsonUmmarshalErrorHttpClient struct {
}

func (c *MockJsonUmmarshalErrorHttpClient) Do(req *http.Request) (*http.Response, error) {
	return &http.Response{
		StatusCode: 200,
		Body:       ioutil.NopCloser(bytes.NewBuffer([]byte("123456"))),
	}, nil
}

type Mock400HttpClient struct {
}

func (c *Mock400HttpClient) Do(req *http.Request) (*http.Response, error) {
	return &http.Response{
		StatusCode: 400,
		Body:       ioutil.NopCloser(bytes.NewBuffer([]byte(""))),
	}, nil
}

type ErrorMockHttpClient struct {
}

func (c *ErrorMockHttpClient) Do(req *http.Request) (*http.Response, error) {
	return nil, fmt.Errorf("error")
}

func TestNewDTOAPIClient_IfHttpClientNil(t *testing.T) {
	assert.NotNil(t, NewDTOAPIClient(nil))
}

func TestNewDTOAPIClient_IfHttpClientNotNil(t *testing.T) {
	assert.NotNil(t, NewDTOAPIClient(&MockHttpClient{}))
}

func TestDTOAPIClient_RequestV2_IfParamEmpty(t *testing.T) {
	ctx := context.Background()
	method := "POST"
	apiPath := "/v1.0/oauth2/accessToken"
	requestModel := &oapiidl.AccessTokenRequestModel{}
	var responseModel *oapiidl.AccessTokenResponseModel

	err := NewDTOAPIClient(&MockHttpClient{}).RequestV2(ctx, method, apiPath, requestModel, responseModel)
	assert.NotNil(t, err)
}

func TestDTOAPIClient_RequestV2_IfRequestModelInvalid(t *testing.T) {
	ctx := context.Background()
	method := "POST"
	apiPath := "/v1.0/oauth2/accessToken"
	requestModel := &map[string]interface{}{
		"key": func() {},
	}
	responseModel := &oapiidl.AccessTokenResponseModel{}

	err := NewDTOAPIClient(&MockHttpClient{}).RequestV2(ctx, method, apiPath, requestModel, responseModel)
	assert.NotNil(t, err)
}

func TestDTOAPIClient_RequestV2_IfAccessToken(t *testing.T) {
	ctx := context.Background()
	method := "POST"
	apiPath := "/v1.0/oauth2/accessToken"
	requestModel := &oapiidl.AccessTokenRequestModel{}
	responseModel := &oapiidl.AccessTokenResponseModel{}

	err := NewDTOAPIClient(&MockHttpClient{}).RequestV2(ctx, method, apiPath, requestModel, responseModel)
	assert.Nil(t, err)
}

func TestDTOAPIClient_RequestV2_IfNotAccessToken(t *testing.T) {
	ctx := context.Background()
	method := "POST"
	apiPath := "/v1.0/anyelse"
	requestModel := &oapiidl.ImTopBoxesOpenRequestModel{}
	responseModel := &oapiidl.ImTopBoxesOpenResponseModel{}

	err := NewDTOAPIClient(&MockHttpClient{}).RequestV2(ctx, method, apiPath, requestModel, responseModel)
	assert.Nil(t, err)
}

func TestDTOAPIClient_RequestV2_IfNotAccessTokenAndError(t *testing.T) {
	ctx := context.Background()
	method := "POST"
	apiPath := "/v1.0/anyelse"
	requestModel := &oapiidl.ImTopBoxesOpenRequestModel{}
	responseModel := &oapiidl.ImTopBoxesOpenResponseModel{}

	err := NewDTOAPIClient(&Mock400HttpClient{}).RequestV2(ctx, method, apiPath, requestModel, responseModel)
	assert.NotNil(t, err)
}

func TestDTOAPIClient_RequestV2_IfContextNil(t *testing.T) {
	method := "POST"
	apiPath := "/v1.0/oauth2/accessToken"
	requestModel := &oapiidl.AccessTokenRequestModel{}
	responseModel := &oapiidl.AccessTokenResponseModel{}

	err := NewDTOAPIClient(&MockHttpClient{}).RequestV2(nil, method, apiPath, requestModel, responseModel)
	assert.NotNil(t, err)
}

func TestDTOAPIClient_RequestV2_IfRequestError(t *testing.T) {
	ctx := context.Background()
	method := "POST"
	apiPath := "/v1.0/oauth2/accessToken"
	requestModel := &oapiidl.AccessTokenRequestModel{}
	responseModel := &oapiidl.AccessTokenResponseModel{}

	err := NewDTOAPIClient(&ErrorMockHttpClient{}).RequestV2(ctx, method, apiPath, requestModel, responseModel)
	assert.NotNil(t, err)
}

func TestDTOAPIClient_RequestV2_IfResponseUnmarshalError(t *testing.T) {
	ctx := context.Background()
	method := "POST"
	apiPath := "/v1.0/oauth2/accessToken"
	requestModel := &oapiidl.AccessTokenRequestModel{}
	responseModel := &oapiidl.AccessTokenResponseModel{}

	err := NewDTOAPIClient(&MockJsonUmmarshalErrorHttpClient{}).RequestV2(ctx, method, apiPath, requestModel, responseModel)
	assert.NotNil(t, err)
}

func TestDTOAPIClient_RequestV2_IfOK(t *testing.T) {
	ctx := context.Background()
	method := "POST"
	apiPath := "/v1.0/oauth2/accessToken"
	requestModel := &oapiidl.AccessTokenRequestModel{}
	responseModel := &oapiidl.AccessTokenResponseModel{}

	err := NewDTOAPIClient(&MockHttpClient{}).RequestV2(ctx, method, apiPath, requestModel, responseModel)
	assert.Nil(t, err)
}

type MockV1ResponseModel struct {
}

func (m *MockV1ResponseModel) GetErrMsg() string {
	return ""
}

func (m *MockV1ResponseModel) GetErrCode() int {
	return 0
}

func TestDTOAPIClient_RequestV1_IfNotAccessToken(t *testing.T) {
	ctx := context.Background()
	method := "POST"
	apiPath := "/v1.0/anyelse"
	requestModel := &oapiidl.ImTopBoxesOpenRequestModel{}
	responseModel := &MockV1ResponseModel{}

	err := NewDTOAPIClient(&MockHttpClient{}).RequestV1(ctx, method, apiPath, requestModel, responseModel)
	assert.Nil(t, err)
}

func TestDTOAPIClient_RequestV1_IfNotAccessTokenAndError(t *testing.T) {
	ctx := context.Background()
	method := "POST"
	apiPath := "/v1.0/anyelse"
	requestModel := &oapiidl.ImTopBoxesOpenRequestModel{}
	responseModel := &MockV1ResponseModel{}

	err := NewDTOAPIClient(&Mock400HttpClient{}).RequestV1(ctx, method, apiPath, requestModel, responseModel)
	assert.NotNil(t, err)
}

func TestDTOAPIClient_RequestV1_IfParamEmpty(t *testing.T) {
	ctx := context.Background()
	method := "POST"
	apiPath := "/v1.0/anyelse"
	requestModel := &oapiidl.ImTopBoxesOpenRequestModel{}
	var responseModel *MockV1ResponseModel

	err := NewDTOAPIClient(&MockHttpClient{}).RequestV1(ctx, method, apiPath, requestModel, responseModel)
	assert.NotNil(t, err)
}

func TestDTOAPIClient_RequestV1_IfRequestModelInvalid(t *testing.T) {
	ctx := context.Background()
	method := "POST"
	apiPath := "/v1.0/anyelse"
	requestModel := map[string]interface{}{
		"key": func() {},
	}
	responseModel := &MockV1ResponseModel{}

	err := NewDTOAPIClient(&MockHttpClient{}).RequestV1(ctx, method, apiPath, requestModel, responseModel)
	assert.NotNil(t, err)
}

func TestDTOAPIClient_RequestV1_IfContextNil(t *testing.T) {
	method := "POST"
	apiPath := "/v1.0/oauth2/accessToken"
	requestModel := &oapiidl.AccessTokenRequestModel{}
	responseModel := &MockV1ResponseModel{}

	err := NewDTOAPIClient(&MockHttpClient{}).RequestV1(nil, method, apiPath, requestModel, responseModel)
	assert.NotNil(t, err)
}

func TestDTOAPIClient_RequestV1_IfRequestError(t *testing.T) {
	method := "POST"
	apiPath := "/v1.0/oauth2/accessToken"
	requestModel := &oapiidl.AccessTokenRequestModel{}
	responseModel := &MockV1ResponseModel{}

	err := NewDTOAPIClient(&ErrorMockHttpClient{}).RequestV1(context.Background(), method, apiPath, requestModel, responseModel)
	assert.NotNil(t, err)
}

func TestDTOAPIClient_RequestV1_IfResponseUnmarshalError(t *testing.T) {
	method := "POST"
	apiPath := "/v1.0/oauth2/accessToken"
	requestModel := &oapiidl.AccessTokenRequestModel{}
	responseModel := &MockV1ResponseModel{}

	err := NewDTOAPIClient(&MockJsonUmmarshalErrorHttpClient{}).RequestV1(context.Background(), method, apiPath, requestModel, responseModel)
	assert.NotNil(t, err)
}

func TestDTOAPIClient_RequestV1_IfOK(t *testing.T) {
	ctx := context.Background()
	method := "POST"
	apiPath := "/v1.0/oauth2/accessToken"
	requestModel := &oapiidl.AccessTokenRequestModel{}
	responseModel := &MockV1ResponseModel{}

	err := NewDTOAPIClient(&MockHttpClient{}).RequestV1(ctx, method, apiPath, requestModel, responseModel)
	assert.Nil(t, err)
}

func TestDTOAPIClient_NotNeedAccessToken(t *testing.T) {
	assert.True(t, NewDTOAPIClient(nil).NotNeedAccessToken("/v1.0/oauth2/accessToken"))
	assert.True(t, NewDTOAPIClient(nil).NotNeedAccessToken("/v1.0/oauth2/corpAccessToken"))
	assert.False(t, NewDTOAPIClient(nil).NotNeedAccessToken("/v1.0/accessToken"))
}
