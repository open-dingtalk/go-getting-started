package oapiidl

/**
 * @Author linya.jj
 * @Date 2022/10/23 11:13 AM
 */

//https://open.dingtalk.com/document/orgapp-server/query-user-details

type UserGetRequestModel struct {
	UserId   *string `json:"userid"`
	Language *string `json:"language"`
}

type UserGetResponseModel struct {
	RequestId *string             `json:"requestId"`
	ErrMsg    *string             `json:"errmsg"`
	ErrCode   *int                `json:"errcode"`
	Result    *UserGetResultModel `json:"result"`
}

func (m *UserGetResponseModel) GetErrMsg() string {
	if m == nil {
		return ""
	}
	return *m.ErrMsg
}

func (m *UserGetResponseModel) GetErrCode() int {
	if m == nil {
		return 0
	}
	return *m.ErrCode
}

type UserGetResultModel struct {
	Extension        *string                             `json:"extension"`
	Unionid          *string                             `json:"unionid"`
	Boss             *bool                               `json:"boss"`
	RoleList         *[]*UserGetResultRoleListModel      `json:"role_list"`
	ExclusiveAccount *bool                               `json:"exclusive_account"`
	ManagerUserid    *string                             `json:"manager_userid"`
	Admin            *bool                               `json:"admin"`
	Remark           *string                             `json:"remark"`
	Title            *string                             `json:"title"`
	HiredDate        *string                             `json:"hired_date"`
	Userid           *string                             `json:"userid"`
	WorkPlace        *string                             `json:"work_place"`
	DeptOrderList    *[]*UserGetResultDeptOrderListModel `json:"dept_order_list"`
	RealAuthed       *bool                               `json:"real_authed"`
	DeptIdList       *[]int                              `json:"dept_id_list"`
	JobNumber        *string                             `json:"job_number"`
	Email            *string                             `json:"email"`
	LeaderInDept     *[]*UserGetResultLeaderInDeptModel  `json:"leader_in_dept"`
	Mobile           *string                             `json:"mobile"`
	Active           *bool                               `json:"active"`
	OrgEmail         *string                             `json:"org_email"`
	Telephone        *string                             `json:"telephone"`
	Avatar           *string                             `json:"avatar"`
	HideMobile       *bool                               `json:"hide_mobile"`
	Senior           *bool                               `json:"senior"`
	Name             *string                             `json:"name"`
	UnionEmpExt      *UserGetResultUnionEmpExtmodel      `json:"union_emp_ext"`
	StateCode        *string                             `json:"state_code"`
}

type UserGetResultRoleListModel struct {
	GroupName *string `json:"group_name"`
	Name      *string `json:"name"`
	Id        *int    `json:"id"`
}

type UserGetResultDeptOrderListModel struct {
	DeptId *int `json:"dept_id"`
	Order  *int `json:"order"`
}

type UserGetResultLeaderInDeptModel struct {
	Leader *bool `json:"leader"`
	DeptId *int  `json:"dept_id"`
}

type UserGetResultUnionEmpExtmodel struct {
	UnionEmpMapList *UserGetResultUnionEmpMapmodel `json:"union_emp_map_list"`
	Userid          *string                        `json:"userid"`
	CorpId          *string                        `json:"corp_id"`
}

type UserGetResultUnionEmpMapmodel struct {
	Userid *string `json:"userid"`
	CorpId *string `json:"corp_id"`
}
