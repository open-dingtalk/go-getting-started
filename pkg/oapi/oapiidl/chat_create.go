package oapiidl

/**
 * @Author linya.jj
 * @Date 2022/10/21 3:34 PM
 */

//ChatCreateRequestModel 创建群会话
type ChatCreateRequestModel struct {
	Name                string   `json:"name"`
	Owner               string   `json:"owner"`
	UserIdList          []string `json:"useridlist"`
	ShowHistoryType     int      `json:"showHistoryType"`
	Searchable          int      `json:"searchable"`
	ValidationType      int      `json:"validationType"`
	MentionAllAuthority int      `json:"mentionAllAuthority"`
	ManagementType      int      `json:"managementType"`
	ChatBannedType      int      `json:"chatBannedType"`
}

type ChatCreateResponseModel struct {
	OpenConversationId string `json:"openConversationId"`
	ChatId             string `json:"chatid"`
	ConversationTag    int    `json:"conversationTag"`
	ErrMsg             string `json:"errmsg"`
	ErrCode            int    `json:"errcode"`
}

func (m *ChatCreateResponseModel) GetErrMsg() string {
	if m == nil {
		return ""
	}
	return m.ErrMsg
}

func (m *ChatCreateResponseModel) GetErrCode() int {
	if m == nil {
		return 0
	}
	return m.ErrCode
}

type IOAPIResponseModelV1 interface {
	GetErrMsg() string
	GetErrCode() int
}
