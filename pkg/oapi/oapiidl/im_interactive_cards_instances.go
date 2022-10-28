package oapiidl

/**
 * @Author linya.jj
 * @Date 2022/10/21 3:32 PM
 */

const (
	// ConversationType 会话类型
	ConversationTypeKSingle = 0
	ConversationTypeKGroup  = 1
)

// ImInteractiveCardsInstancesRequestModel 机器人发送吊顶卡片
type ImInteractiveCardsInstancesRequestModel struct {
	CardTemplateId     *string                                                      `json:"cardTemplateId"`
	OpenConversationId *string                                                      `json:"openConversationId"`
	ReceiverUserIdList *[]*string                                                   `json:"receiverUserIdList"`
	OutTrackId         *string                                                      `json:"outTrackId"`
	RobotCode          *string                                                      `json:"robotCode"`
	ConversationType   *int                                                         `json:"conversationType"`
	CallbackRouteKey   *string                                                      `json:"callbackRouteKey"`
	CardData           *ImInteractiveCardsInstancesRequestCardDataModel             `json:"cardData"`
	PrivateData        *map[string]*ImInteractiveCardsInstancesRequestCardDataModel `json:"privateData"`
	ChatBotId          *string                                                      `json:"chatBotId"`
	UserIdType         *int                                                         `json:"userIdType"`
	PullStrategy       *bool                                                        `json:"pullStrategy"`
}

type ImInteractiveCardsInstancesRequestCardDataModel struct {
	CardParamMap        *map[string]interface{} `json:"cardParamMap"`
	CardMediaIdParamMap *map[string]interface{} `json:"cardMediaIdParamMap"`
}

type ImInteractiveCardsInstancesResponseModel struct {
	ProcessQueryKey *string `json:"processQueryKey"`
}
