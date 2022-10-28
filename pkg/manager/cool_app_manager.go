package manager

import (
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/open-dingtalk/go-getting-started/pkg/config"
	"github.com/open-dingtalk/go-getting-started/pkg/manager/model"
	"github.com/open-dingtalk/go-getting-started/pkg/oapi"
	"github.com/open-dingtalk/go-getting-started/pkg/oapi/oapiidl"
	"time"
)

/**
 * @Author linya.jj
 * @Date 2022/10/19 11:09 AM
 */

type CoolAppManager struct {
	dtOapiClient *oapi.DTOAPIClient
}

func NewCoolAppManager(dtOapiClient *oapi.DTOAPIClient) *CoolAppManager {
	return &CoolAppManager{
		dtOapiClient: dtOapiClient,
	}
}

// SendTextMessage https://open.dingtalk.com/document/orgapp-server/send-robot-group-chat-messages
func (m *CoolAppManager) SendTextMessage(ctx context.Context, openConversationId, messageContent string) error {
	if messageContent == "" {
		messageContent = "钉钉，让进步发生！"
	}

	msgParamModel := map[string]string{
		"content": messageContent,
	}

	msgParamContentByte, err := json.Marshal(msgParamModel)
	if err != nil {
		return errors.WithStack(err)
	}

	msgParamContentStr := string(msgParamContentByte)
	msgKey := "sampleText"
	robotCode := config.GetConfig().GetRobotInfo().GetCode()
	coolAppCode := config.GetConfig().GetCoolAppInfo().GetCode()

	textMessageRequestModel := &oapiidl.RobotGroupMessagesSendRequestModel{
		MsgParam:           &msgParamContentStr,
		MsgKey:             &msgKey,
		OpenConversationId: &openConversationId,
		RobotCode:          &robotCode,
		CoolAppCode:        &coolAppCode,
	}

	textMessageResponseModel := &oapiidl.RobotGroupMessagesSendResponseModel{}

	err = m.dtOapiClient.RequestV2(ctx, "POST", "/v1.0/robot/groupMessages/send",
		textMessageRequestModel, textMessageResponseModel)

	return err
}

// SendMessageCard https://open.dingtalk.com/document/group/robots-send-interactive-cards
func (m *CoolAppManager) SendMessageCard(ctx context.Context, openConversationId string) error {
	cardParam := map[string]interface{}{
		// 酷应用介绍视频，内嵌
		"videoUrl": "https://cloud.video.taobao.com/play/u/null/p/1/e/6/t/1/d/ud/352793594610.mp4",
	}

	cardDataB, err := json.Marshal(cardParam)
	if err != nil {
		return errors.WithStack(err)
	}

	cardBizId := "msgcardid" + uuid.New().String()
	robotCode := config.GetConfig().GetRobotInfo().GetCode()
	cardData := string(cardDataB)
	cardTemplateId := config.GetConfig().GetCardInfo().GetMessageCardTemplateId()

	cardRequestModel := &oapiidl.RobotInteractiveCardRequestModel{}
	cardRequestModel.CardTemplateId = &cardTemplateId
	cardRequestModel.OpenConversationId = &openConversationId
	cardRequestModel.CardBizId = &cardBizId
	cardRequestModel.RobotCode = &robotCode
	cardRequestModel.CardData = &cardData

	cardResponseModel := &oapiidl.RobotInteractiveCardResponseModel{}

	err = m.dtOapiClient.RequestV2(ctx, "POST", "/v1.0/im/v1.0/robot/interactiveCards/send",
		cardRequestModel, cardResponseModel)

	return err
}

// SendTopCard https://open.dingtalk.com/document/orgapp-server/create-an-interactive-card-instance-1
func (m *CoolAppManager) SendTopCard(ctx context.Context, openConversationId string) error {
	cardParam := map[string]interface{}{}

	cardBizId := "topcardid" + uuid.New().String()
	robotCode := config.GetConfig().GetRobotInfo().GetCode()
	cardTemplateId := config.GetConfig().GetCardInfo().GetTopCardTemplateId()
	conversationType := oapiidl.ConversationTypeKGroup

	cardRequestModel := &oapiidl.ImInteractiveCardsInstancesRequestModel{}
	cardRequestModel.CardTemplateId = &cardTemplateId
	cardRequestModel.OpenConversationId = &openConversationId
	cardRequestModel.OutTrackId = &cardBizId
	cardRequestModel.RobotCode = &robotCode
	cardRequestModel.CardData = &oapiidl.ImInteractiveCardsInstancesRequestCardDataModel{
		CardParamMap: &cardParam,
	}
	cardRequestModel.ConversationType = &conversationType

	cardResponseModel := &oapiidl.ImInteractiveCardsInstancesResponseModel{}

	err := m.dtOapiClient.RequestV2(ctx, "POST", "/v1.0/im/interactiveCards/instances",
		cardRequestModel, cardResponseModel)
	if err != nil {
		return err
	}

	setTopRequestModel := &oapiidl.ImTopBoxesOpenRequestModel{}
	platforms := "ios|mac|android|win"
	expireTime := time.Now().UnixMilli() + 5*60*1000

	setTopRequestModel.ConversationType = &conversationType
	setTopRequestModel.OpenConversationId = &openConversationId
	setTopRequestModel.OutTrackId = &cardBizId
	setTopRequestModel.RobotCode = &robotCode
	setTopRequestModel.Platforms = &platforms
	setTopRequestModel.ExpiredTime = &expireTime

	setTopResponseModel := &oapiidl.ImTopBoxesOpenResponseModel{}

	err = m.dtOapiClient.RequestV2(ctx, "POST", "/v1.0/im/topBoxes/open",
		setTopRequestModel, setTopResponseModel)

	return err
}

// GetUserInfo 获取用户信息
// 免登：https://open.dingtalk.com/document/orgapp-server/obtain-the-userid-of-a-user-by-using-the-log-free
// 用户信息：https://open.dingtalk.com/document/orgapp-server/query-user-details
func (m *CoolAppManager) GetUserInfo(ctx context.Context, requestAuthInfo string) (*model.UserInfoModel, error) {
	getUserInfoRequestModel := &oapiidl.UserGetUserInfoRequestModel{
		Code: &requestAuthInfo,
	}
	getUserInfoResponseModel := &oapiidl.UserGetUserInfoResponseModel{}

	err := m.dtOapiClient.RequestV1(ctx, "POST", "/topapi/v2/user/getuserinfo",
		getUserInfoRequestModel, getUserInfoResponseModel)
	if err != nil {
		return nil, err
	}

	language := "zh_CN"
	userGetRequestModel := &oapiidl.UserGetRequestModel{
		UserId:   getUserInfoResponseModel.Result.Userid,
		Language: &language,
	}
	userGetResponseModel := &oapiidl.UserGetResponseModel{}

	err = m.dtOapiClient.RequestV1(ctx, "POST", "/topapi/v2/user/get",
		userGetRequestModel, userGetResponseModel)
	if err != nil {
		return nil, err
	}

	userInfoModel := &model.UserInfoModel{
		Name:   *userGetResponseModel.Result.Name,
		UserId: *userGetResponseModel.Result.Userid,
		Avatar: *userGetResponseModel.Result.Avatar,
	}

	return userInfoModel, nil
}
