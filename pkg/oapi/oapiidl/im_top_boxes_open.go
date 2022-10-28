package oapiidl

/**
 * @Author linya.jj
 * @Date 2022/10/21 3:33 PM
 */

//ImTopBoxesOpenRequestModel 卡片置顶
type ImTopBoxesOpenRequestModel struct {
	ConversationType   *int      `json:"conversationType"`
	OpenConversationId *string   `json:"openConversationId"`
	ReceiverUserIdList *[]string `json:"receiverUserIdList"`
	RobotCode          *string   `json:"robotCode"`
	OutTrackId         *string   `json:"outTrackId"`
	CoolAppCode        *string   `json:"coolAppCode"`
	ExpiredTime        *int64    `json:"expiredTime"`
	Platforms          *string   `json:"platforms"`
}

type ImTopBoxesOpenResponseModel struct {
}
