package oapiidl

/**
 * @Author linya.jj
 * @Date 2022/10/21 3:32 PM
 */

// RobotInteractiveCardRequestModel 机器人发送链接卡片
type RobotInteractiveCardRequestModel struct {
	CardTemplateId        *string                                      `json:"cardTemplateId"`
	OpenConversationId    *string                                      `json:"openConversationId"`
	SingleChatReceiver    *string                                      `json:"singleChatReceiver"`
	CardBizId             *string                                      `json:"cardBizId"`
	RobotCode             *string                                      `json:"robotCode"`
	CallbackUrl           *string                                      `json:"callbackUrl"`
	CardData              *string                                      `json:"cardData"`
	UserIdPrivateDataMap  *string                                      `json:"userIdPrivateDataMap"`
	UnionIdPrivateDataMap *string                                      `json:"unionIdPrivateDataMap"`
	SendOptions           *RobotInteractiveCardRequestSendOptionsModel `json:"sendOptions"`
	PullStrategy          *bool                                        `json:"pullStrategy"`
}

type RobotInteractiveCardRequestSendOptionsModel struct {
	AtUserListJson   *string `json:"atUserListJson"`
	AtAll            *bool   `json:"atAll"`
	ReceiverListJson *string `json:"receiverListJson"`
	CardPropertyJson *string `json:"cardPropertyJson"`
}

type RobotInteractiveCardResponseModel struct {
	ProcessQueryKey *string `json:"processQueryKey"`
}
