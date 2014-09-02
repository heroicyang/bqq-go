// A Tecent Business QQ client in go.
// https://github.com/heroicyang/bqq-go
//
// @author Herioc Yang <me@heroicyang.com>

package bqq

// Get company information.
// Please refer to API wiki: http://open.b.qq.com/wiki/api:corporation_get
func (session *Session) GetCompanyInfo() (Result, error) {
	return session.Get("/api/corporation/get", nil)
}

// Get departments list with optional `timestamp`.
// Please refer to API wiki: http://open.b.qq.com/wiki/api:dept_list
func (session *Session) GetDepartments(timestamp int) (Result, error) {
	params := Params{"timestamp": timestamp}
	return session.Get("/api/dept/list", params)
}

// Get departments with dept_ids(comma-separated of dept_id).
// Please refer to API wiki: http://open.b.qq.com/wiki/api:dept_info
func (session *Session) GetDepartmentsByIds(ids string) (Result, error) {
	params := Params{"dept_ids": ids}
	return session.Get("/api/dept/info", params)
}

// Get users list with optional `timestamp`.
// Please refer to API wiki: http://open.b.qq.com/wiki/api:user_list
func (session *Session) GetUsers(timestamp int) (Result, error) {
	params := Params{"timestamp": timestamp}
	return session.Get("/api/user/list", params)
}

// Get users by open_ids with open_ids(comma-separated of open_id).
// Please refer to API wiki: http://open.b.qq.com/wiki/api:user_info
func (session *Session) GetUsersByOpenIds(openIds string) (Result, error) {
	params := Params{"open_ids": openIds}
	return session.Get("/api/user/info", params)
}

// Get users face with open_ids(comma-separated of open_id).
// Please refer to API wiki: http://open.b.qq.com/wiki/api:api_user_face
func (session *Session) GetUsersFace(openIds string, typeId int) (Result, error) {
	if typeId == 0 {
		typeId = 4
	}

	params := Params{"open_ids": openIds, "type_id": typeId}
	return session.Get("/api/user/face", params)
}

// Get users email with open_ids(comma-separated of open_id).
// Please refer to API wiki: http://open.b.qq.com/wiki/api:api_user_email
func (session *Session) GetUsersEmail(openIds string) (Result, error) {
	params := Params{"open_ids": openIds}
	return session.Get("/api/user/email", params)
}

// Get users mobile with open_ids(comma-separated of open_id).
// Please refer to API wiki: http://open.b.qq.com/wiki/api:api_user_mobile
func (session *Session) GetUsersMobile(openIds string) (Result, error) {
	params := Params{"open_ids": openIds}
	return session.Get("/api/user/mobile", params)
}

// Get users qq with open_ids(comma-separated of open_id).
// Please refer to API wiki: http://open.b.qq.com/wiki/api:api_user_qq
func (session *Session) GetUsersQQ(openIds string) (Result, error) {
	params := Params{"open_ids": openIds}
	return session.Get("/api/user/qq", params)
}

// Send tips to QQ client.
// Please refer to API wiki: http://open.b.qq.com/wiki/api:tips_send
func (session *Session) SendTip(data map[string]interface{}) (Result, error) {
	if _, ok := data["receivers"]; !ok {
		data["to_all"] = 1
	}

	return session.Post("/api/tips/send", data)
}

// Send broadcast to QQ client.
// Please refer to API wiki: http://open.b.qq.com/wiki/api:broadcast_send
func (session *Session) SendBroadcast(data map[string]interface{}) (Result, error) {
	return session.Post("/api/broadcast/send", data)
}

// Send sms to QQ client.
// Please refer to API wiki: http://open.b.qq.com/wiki/api:sms_send
func (session *Session) SendSms(data map[string]interface{}) (Result, error) {
	return session.Post("/api/sms/send", data)
}
