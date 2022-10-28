package oapiidl

/**
 * @Author linya.jj
 * @Date 2022/10/9 2:31 PM
 */

// RobotGroupMessagesSendRequestModel 机器人发送群聊消息
type RobotGroupMessagesSendRequestModel struct {
	MsgParam           *string `json:"msgParam" required:"true"`
	MsgKey             *string `json:"msgKey" required:"true"`
	OpenConversationId *string `json:"openConversationId" required:"true"`
	RobotCode          *string `json:"robotCode" required:"true"`
	CoolAppCode        *string `json:"coolAppCode" required:"true"`
}

type RobotGroupMessagesSendResponseModel struct {
	ProcessQueryKey *string `json:"processQueryKey"`
}
