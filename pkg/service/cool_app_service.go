package service

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/open-dingtalk/go-getting-started/pkg/logger"
	"github.com/open-dingtalk/go-getting-started/pkg/manager"
	"github.com/open-dingtalk/go-getting-started/pkg/service/vo"
	"github.com/open-dingtalk/go-getting-started/pkg/utils"
	"io/ioutil"
)

/**
 * @Author linya.jj
 * @Date 2022/10/19 11:14 AM
 */

type CoolAppService struct {
	coolAppManager *manager.CoolAppManager
}

func NewCoolAppService(coolAppManager *manager.CoolAppManager) *CoolAppService {
	return &CoolAppService{
		coolAppManager: coolAppManager,
	}
}

func (s *CoolAppService) SendTextMessage(c *gin.Context) {
	var err error

	defer func() {
		if err != nil {
			s.writeResponseError(c, err)
		}
	}()

	requestVO := &vo.TextMessageSendRequestVO{}
	err = s.getRequestFromBodyAndCheckValid(c, requestVO)
	if err != nil {
		return
	}

	err = s.coolAppManager.SendTextMessage(c, requestVO.OpenConversationId, requestVO.Txt)
	if err != nil {
		return
	}

	responseVO := &vo.TextMessageSendResponseVO{}
	s.writeSuccessResponseVO(c, "", responseVO)
}

func (s *CoolAppService) SendMessageCard(c *gin.Context) {
	var err error

	defer func() {
		if err != nil {
			s.writeResponseError(c, err)
		}
	}()

	requestVO := &vo.MessageCardSendRequestVO{}
	err = s.getRequestFromBodyAndCheckValid(c, requestVO)
	if err != nil {
		return
	}

	err = s.coolAppManager.SendMessageCard(c, requestVO.OpenConversationId)
	if err != nil {
		return
	}

	responseVO := &vo.MessageCardSendResponseVO{}
	s.writeSuccessResponseVO(c, "", responseVO)
}

func (s *CoolAppService) SendTopCard(c *gin.Context) {
	var err error

	defer func() {
		if err != nil {
			s.writeResponseError(c, err)
		}
	}()

	requestVO := &vo.TopCardSendRequestVO{}
	err = s.getRequestFromBodyAndCheckValid(c, requestVO)
	if err != nil {
		return
	}

	err = s.coolAppManager.SendTopCard(c, requestVO.OpenConversationId)
	if err != nil {
		return
	}

	responseVO := &vo.TopCardSendResponseVO{}
	s.writeSuccessResponseVO(c, "", responseVO)
}

func (s *CoolAppService) GetUserInfo(c *gin.Context) {
	var err error

	defer func() {
		if err != nil {
			s.writeResponseError(c, err)
		}
	}()

	requestVO := &vo.GetUserInfoRequestVO{
		RequestAuthCode: c.Query("requestAuthCode"),
	}

	err = requestVO.CheckValid()
	if err != nil {
		return
	}

	userInfoModel, err := s.coolAppManager.GetUserInfo(c, requestVO.RequestAuthCode)
	if err != nil {
		return
	}

	responseVO := &vo.GetUserInfoResponseVO{
		Name:   userInfoModel.Name,
		UserId: userInfoModel.UserId,
		Avatar: userInfoModel.Avatar,
	}
	s.writeSuccessResponseVO(c, "", responseVO)
}

func (s *CoolAppService) getRequestFromBodyAndCheckValid(c *gin.Context, apiRequestVO vo.IApiRequestVO) error {
	if apiRequestVO == nil {
		return errors.WithStack(fmt.Errorf("getRequestBody.apiRequestVO nil"))
	}

	data, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		return errors.WithStack(err)
	}

	logger.GetDefaultLogger().Info(string(data))

	err = json.Unmarshal(data, apiRequestVO)
	if err != nil {
		return errors.WithStack(err)
	}

	return apiRequestVO.CheckValid()
}

func (s *CoolAppService) writeResponseError(c *gin.Context, err error) {
	if err == nil {
		c.JSON(500, &vo.ApiResponseVO{
			Success: false,
			Code:    500,
			Message: "",
			Data:    "",
		})
		return
	}

	logger.GetDefaultLogger().Error(err.Error())

	if errorTuple, ok := err.(*utils.ApiErrorTuple); ok && err != nil {
		c.JSON(400, &vo.ApiResponseVO{
			Success: false,
			Code:    400,
			Message: "",
			Data:    errorTuple,
		})
		return
	}

	c.JSON(500, &vo.ApiResponseVO{
		Success: false,
		Code:    500,
		Message: "",
		Data:    utils.NewInternalServerErrorTuple(err),
	})
}

func (s *CoolAppService) writeSuccessResponseVO(c *gin.Context, contentType string, responseVO interface{}) {
	data, err := json.Marshal(vo.NewSuccessApiResponseVO(responseVO))
	if err != nil {
		s.writeResponseError(c, err)
		return
	}

	if contentType == "" {
		contentType = "application/json"
	}

	c.Data(200, contentType, data)
}
