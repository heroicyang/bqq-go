// A Tencent Business QQ client in go.
// https://github.com/heroicyang/bqq-go
//
// @author Herioc Yang <me@heroicyang.com>

package bqq

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

const (
	APP_ID        = "12345"
	APP_SECRET    = "asdfg"
	ACCESS_TOKEN  = "iKWvqqUUNYAPYDY1qjdLU3KPcg8J2kiz"
	REFRESH_TOKEN = "eyGhsdEJzv3Ayi4IVYKkqmcMUac8BhPe"
	OPEN_ID       = "KwglVwWP99XdMNEhvbHJ1I5Uvu"
	COMPANY_ID    = "zTO8ehphrLtEBX28OKD99gbLRbqDSUxn"
	COMPANY_TOKEN = "dukJ7o9bVMk8zsQ8tYrsBpk5dS5pbkuY"
)

func tokenHandler(w http.ResponseWriter, r *http.Request) {
	var res *Result

	if r.FormValue("code") == "" {
		res = &Result{
			"ret": 1,
			"msg": "code required",
		}
	} else if r.FormValue("state") == "" {
		res = &Result{
			"ret": 1,
			"msg": "state required",
		}
	} else {
		res = &Result{
			"ret": 0,
			"data": map[string]interface{}{
				"access_token":  ACCESS_TOKEN,
				"refresh_token": REFRESH_TOKEN,
				"expires_in":    720000,
				"open_id":       OPEN_ID,
				"state":         r.FormValue("state"),
			},
		}
	}

	result, _ := json.Marshal(res)

	w.Header().Set("Content-Type", "application/json")
	w.Write(result)
}

func companyTokenHandler(w http.ResponseWriter, r *http.Request) {
	var res *Result

	if r.FormValue("code") == "" {
		res = &Result{
			"ret": 1,
			"msg": "code required",
		}
	} else if r.FormValue("state") == "" {
		res = &Result{
			"ret": 1,
			"msg": "state required",
		}
	} else {
		res = &Result{
			"ret": 0,
			"data": map[string]interface{}{
				"company_id":    COMPANY_ID,
				"company_token": COMPANY_TOKEN,
				"refresh_token": REFRESH_TOKEN,
				"expires_in":    720000,
				"state":         r.FormValue("state"),
			},
		}
	}

	result, _ := json.Marshal(res)

	w.Header().Set("Content-Type", "application/json")
	w.Write(result)
}

func checkAccess(r *http.Request) (canAccess bool) {
	if r.FormValue("app_id") != APP_ID || r.FormValue("company_id") != COMPANY_ID ||
		r.FormValue("company_token") != COMPANY_TOKEN {
		canAccess = false
		return
	}

	canAccess = true
	return
}

func companyInfoHandler(w http.ResponseWriter, r *http.Request) {
	var res string

	if !checkAccess(r) {
		res = `{ "ret": 1, "msg": "auth required" }`
	} else {
		res = `{
			"ret": 0,
			"data": {
				"company_name":     "Company Name",
				"company_fullname": "Company FullName"
			}
		}`
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(res))
}

func departmentsListHandler(w http.ResponseWriter, r *http.Request) {
	var res string

	if !checkAccess(r) {
		res = `{ "ret": 1, "msg": "auth required" }`
	} else {
		res = `{
			"ret": 0,
			"data": {
				"timestamp": 1408414724,
				"items": [{
					"dept_id":   74537005,
					"p_dept_id": 1136750599,
					"dept_name": "test"
				}]
			}
		}`
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(res))
}

func departmentsInfoHandler(w http.ResponseWriter, r *http.Request) {
	var res string

	if !checkAccess(r) {
		res = `{ "ret": 1, "msg": "auth required" }`
	} else {
		res = `{
			"ret": 0,
			"data": {
				"74537005": {
					"dept_id": 74537005,
		      "p_dept_id": 1136750599,
		      "dept_name": "test"
				}
			}
		}`
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(res))
}

func usersListHandler(w http.ResponseWriter, r *http.Request) {
	var res string

	if !checkAccess(r) {
		res = `{ "ret": 1, "msg": "auth required" }`
	} else {
		res = `{
			"ret": 0,
			"data": {
				"timestamp": 1408414724,
				"items": [{
					"open_id": "2c0c7cdf67fd7d442db2390dce393bce",
	        "gender": 1,
	        "account": "yx",
	        "realname": "Heroic Yang",
	        "p_dept_id": "74537005",
	        "mobile": 0,
	        "hidden": 0,
	        "role_id": 0
				}]
			}
		}`
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(res))
}

func usersInfoHandler(w http.ResponseWriter, r *http.Request) {
	var res string

	if !checkAccess(r) {
		res = `{ "ret": 1, "msg": "auth required" }`
	} else {
		res = `{
			"ret": 0,
			"data": {
				"2c0c7cdf67fd7d442db2390dce393bce": {
		      "open_id": "2c0c7cdf67fd7d442db2390dce393bce",
		      "gender": 1,
		      "dept_id": 1565199740,
		      "dept_ids": [
		        1565199740,
		        1291619508,
		        140325039
		      ],
		      "role_id": 16777216,
		      "account": "Karl",
		      "name": "卡尔",
		      "titles": [
		        "员工",
		        "秘书",
		        "组长"
		      ],
		      "title": "员工",
		      "email": 0,
		      "mobile": 1,
		      "hidden": 0
		    }
			}
		}`
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(res))
}

func usersFaceHandler(w http.ResponseWriter, r *http.Request) {
	var res string

	if !checkAccess(r) {
		res = `{ "ret": 1, "msg": "auth required" }`
	} else {
		res = `{
			"ret": 0,
			"data": {
				"items": {
		      "2c0c7cdf67fd7d442db2390dce393bce": "http://q4.qlogo.cn/g?b=qq&k=9WO8jIfeqm2smbqibCeBIzw&s=100&t=1401893441"
		    }
			}
		}`
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(res))
}

func usersEmailHandler(w http.ResponseWriter, r *http.Request) {
	var res string

	if !checkAccess(r) {
		res = `{ "ret": 1, "msg": "auth required" }`
	} else {
		res = `{
			"ret": 0,
			"data": {
				"2c0c7cdf67fd7d442db2390dce393bce": "a@b.c"
			}
		}`
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(res))
}

func usersMobileHandler(w http.ResponseWriter, r *http.Request) {
	var res string

	if !checkAccess(r) {
		res = `{ "ret": 1, "msg": "auth required" }`
	} else {
		res = `{
			"ret": 0,
			"data": {
				"2c0c7cdf67fd7d442db2390dce393bce": "13800000000"
			}
		}`
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(res))
}

func usersQQHandler(w http.ResponseWriter, r *http.Request) {
	var res string

	if !checkAccess(r) {
		res = `{ "ret": 1, "msg": "auth required" }`
	} else {
		res = `{
			"ret": 0,
			"data": {
				"2c0c7cdf67fd7d442db2390dce393bce": "10000"
			}
		}`
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(res))
}

func sendTipHandler(w http.ResponseWriter, r *http.Request) {
	var res string

	if !checkAccess(r) {
		res = `{ "ret": 1, "msg": "auth required" }`
	} else {
		if r.FormValue("window_title") == "" || r.FormValue("tips_title") == "" ||
			r.FormValue("tips_content") == "" {
			res = `{ "ret": 1, "msg": "parameters is not complete" }`
		} else {
			res = `{ "ret": 0 }`
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(res))
}

func sendBroadcastHandler(w http.ResponseWriter, r *http.Request) {
	var res string

	if !checkAccess(r) {
		res = `{ "ret": 1, "msg": "auth required" }`
	} else {
		if r.FormValue("title") == "" || r.FormValue("content") == "" ||
			r.FormValue("recv_open_ids") == "" {
			res = `{ "ret": 1, "msg": "parameters is not complete" }`
		} else {
			res = `{ "ret": 0 }`
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(res))
}

func sendSmsHandler(w http.ResponseWriter, r *http.Request) {
	var res string

	if !checkAccess(r) {
		res = `{ "ret": 1, "msg": "auth required" }`
	} else {
		if r.FormValue("recv_phones") == "" || r.FormValue("recv_open_ids") == "" ||
			r.FormValue("content") == "" {
			res = `{ "ret": 1, "msg": "parameters is not complete" }`
		} else {
			res = `{ "ret": 0 }`
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(res))
}

func TestAuthorizeUrl(t *testing.T) {
	app := Init(APP_ID, APP_SECRET)
	app.BaseEndPoint = "https://demo.com"

	url, err := app.GetAuthorizeUri("state", "")
	if err != nil {
		t.Fatalf("cannot get authorize url. [e:%v]", err)
	}

	t.Logf("authorize url is: %s", url)
}

func TestAccessToken(t *testing.T) {
	ts := httptest.NewServer(
		http.StripPrefix("/oauth2/token", http.HandlerFunc(tokenHandler)))
	app := Init(APP_ID, APP_SECRET)

	app.BaseEndPoint = ts.URL

	defer ts.Close()

	res, err := app.GetAccessToken("code", "state")
	if err != nil {
		t.Fatalf("cannot get access token. [e:%v]", err)
	}

	if res.Get("ret") == float64(0) {
		t.Logf("access token is '%v'", res.Get("data.access_token"))
	}
}

func TestCompanyToken(t *testing.T) {
	ts := httptest.NewServer(
		http.StripPrefix("/oauth2/companyToken", http.HandlerFunc(companyTokenHandler)))
	app := Init(APP_ID, APP_SECRET)

	app.BaseEndPoint = ts.URL

	defer ts.Close()

	res, err := app.GetCompanyToken("code", "state")
	if err != nil {
		t.Fatalf("cannot get company token. [e:%v]", err)
	}

	if res.Get("ret") == float64(0) {
		t.Logf("company token is '%v'", res.Get("data.company_token"))
	}
}

func TestCompanyInfoAPI(t *testing.T) {
	ts := httptest.NewServer(
		http.StripPrefix("/api/corporation/get", http.HandlerFunc(companyInfoHandler)))
	app := Init(APP_ID, APP_SECRET)

	app.BaseEndPoint = ts.URL

	defer ts.Close()

	session := app.CreateSession(COMPANY_ID, COMPANY_TOKEN)

	res, err := session.GetCompanyInfo()
	if err != nil {
		t.Fatalf("cannot get company info. [e:%v]", err)
	}

	if res.Get("ret") == float64(0) {
		t.Logf("company info '%v'", res.Get("data"))
	}
}

func TestDepartmentsListAPI(t *testing.T) {
	ts := httptest.NewServer(
		http.StripPrefix("/api/dept/list", http.HandlerFunc(departmentsListHandler)))
	app := Init(APP_ID, APP_SECRET)

	app.BaseEndPoint = ts.URL

	defer ts.Close()

	session := app.CreateSession(COMPANY_ID, COMPANY_TOKEN)

	res, err := session.GetDepartments(0)
	if err != nil {
		t.Fatalf("cannot get department list. [e:%v]", err)
	}

	if res.Get("ret") == float64(0) {
		t.Logf("department list '%v'", res.Get("data"))
	}
}

func TestDepartmentsInfoAPI(t *testing.T) {
	ts := httptest.NewServer(
		http.StripPrefix("/api/dept/info", http.HandlerFunc(departmentsInfoHandler)))
	app := Init(APP_ID, APP_SECRET)

	app.BaseEndPoint = ts.URL

	defer ts.Close()

	session := app.CreateSession(COMPANY_ID, COMPANY_TOKEN)

	res, err := session.GetDepartmentsByIds("74537005")
	if err != nil {
		t.Fatalf("cannot get departments info. [e:%v]", err)
	}

	if res.Get("ret") == float64(0) {
		t.Logf("departments info '%v'", res.Get("data"))
	}
}

func TestUsersListAPI(t *testing.T) {
	ts := httptest.NewServer(
		http.StripPrefix("/api/user/list", http.HandlerFunc(usersListHandler)))
	app := Init(APP_ID, APP_SECRET)

	app.BaseEndPoint = ts.URL

	defer ts.Close()

	session := app.CreateSession(COMPANY_ID, COMPANY_TOKEN)

	res, err := session.GetUsers(0)
	if err != nil {
		t.Fatalf("cannot get user list. [e:%v]", err)
	}

	if res.Get("ret") == float64(0) {
		t.Logf("user list '%v'", res.Get("data"))
	}
}

func TestUsersInfoAPI(t *testing.T) {
	ts := httptest.NewServer(
		http.StripPrefix("/api/user/info", http.HandlerFunc(usersInfoHandler)))
	app := Init(APP_ID, APP_SECRET)

	app.BaseEndPoint = ts.URL

	defer ts.Close()

	session := app.CreateSession(COMPANY_ID, COMPANY_TOKEN)

	res, err := session.GetUsersByOpenIds("2c0c7cdf67fd7d442db2390dce393bce")
	if err != nil {
		t.Fatalf("cannot get users info. [e:%v]", err)
	}

	if res.Get("ret") == float64(0) {
		t.Logf("users info '%v'", res.Get("data"))
	}
}

func TestUsersFaceAPI(t *testing.T) {
	ts := httptest.NewServer(
		http.StripPrefix("/api/user/face", http.HandlerFunc(usersFaceHandler)))
	app := Init(APP_ID, APP_SECRET)

	app.BaseEndPoint = ts.URL

	defer ts.Close()

	session := app.CreateSession(COMPANY_ID, COMPANY_TOKEN)

	res, err := session.GetUsersFace("2c0c7cdf67fd7d442db2390dce393bce", 3)
	if err != nil {
		t.Fatalf("cannot get users face. [e:%v]", err)
	}

	if res.Get("ret") == float64(0) {
		t.Logf("users face '%v'", res.Get("data"))
	}
}

func TestUsersEmailAPI(t *testing.T) {
	ts := httptest.NewServer(
		http.StripPrefix("/api/user/email", http.HandlerFunc(usersEmailHandler)))
	app := Init(APP_ID, APP_SECRET)

	app.BaseEndPoint = ts.URL

	defer ts.Close()

	session := app.CreateSession(COMPANY_ID, COMPANY_TOKEN)

	res, err := session.GetUsersEmail("2c0c7cdf67fd7d442db2390dce393bce")
	if err != nil {
		t.Fatalf("cannot get users email. [e:%v]", err)
	}

	if res.Get("ret") == float64(0) {
		t.Logf("users email '%v'", res.Get("data"))
	}
}

func TestUsersMobileAPI(t *testing.T) {
	ts := httptest.NewServer(
		http.StripPrefix("/api/user/mobile", http.HandlerFunc(usersMobileHandler)))
	app := Init(APP_ID, APP_SECRET)

	app.BaseEndPoint = ts.URL

	defer ts.Close()

	session := app.CreateSession(COMPANY_ID, COMPANY_TOKEN)

	res, err := session.GetUsersMobile("2c0c7cdf67fd7d442db2390dce393bce")
	if err != nil {
		t.Fatalf("cannot get users mobile. [e:%v]", err)
	}

	if res.Get("ret") == float64(0) {
		t.Logf("users mobile '%v'", res.Get("data"))
	}
}

func TestUsersQQAPI(t *testing.T) {
	ts := httptest.NewServer(
		http.StripPrefix("/api/user/qq", http.HandlerFunc(usersQQHandler)))
	app := Init(APP_ID, APP_SECRET)

	app.BaseEndPoint = ts.URL

	defer ts.Close()

	session := app.CreateSession(COMPANY_ID, COMPANY_TOKEN)

	res, err := session.GetUsersQQ("2c0c7cdf67fd7d442db2390dce393bce")
	if err != nil {
		t.Fatalf("cannot get users qq. [e:%v]", err)
	}

	if res.Get("ret") == float64(0) {
		t.Logf("users qq '%v'", res.Get("data"))
	}
}

func TestSendTipAPI(t *testing.T) {
	ts := httptest.NewServer(
		http.StripPrefix("/api/tips/send", http.HandlerFunc(sendTipHandler)))
	app := Init(APP_ID, APP_SECRET)

	app.BaseEndPoint = ts.URL

	defer ts.Close()

	session := app.CreateSession(COMPANY_ID, COMPANY_TOKEN)

	data := map[string]interface{}{
		"window_title": "window_title",
		"tips_title":   "tips_title",
		"tips_content": "tips_content",
	}

	res, err := session.SendTip(data)
	if err != nil {
		t.Fatalf("cannot send tip. [e:%v]", err)
	}

	if res.Get("ret") == float64(0) {
		t.Log("tip sended")
	}
}

func TestSendBroadcastAPI(t *testing.T) {
	ts := httptest.NewServer(
		http.StripPrefix("/api/broadcast/send", http.HandlerFunc(sendBroadcastHandler)))
	app := Init(APP_ID, APP_SECRET)

	app.BaseEndPoint = ts.URL

	defer ts.Close()

	session := app.CreateSession(COMPANY_ID, COMPANY_TOKEN)

	data := map[string]interface{}{
		"title":         "title",
		"content":       "content",
		"recv_open_ids": "recv_open_ids",
	}

	res, err := session.SendBroadcast(data)
	if err != nil {
		t.Fatalf("cannot send broadcast. [e:%v]", err)
	}

	if res.Get("ret") == float64(0) {
		t.Log("broadcast sended")
	}
}

func TestSendSmsAPI(t *testing.T) {
	ts := httptest.NewServer(
		http.StripPrefix("/api/sms/send", http.HandlerFunc(sendSmsHandler)))
	app := Init(APP_ID, APP_SECRET)

	app.BaseEndPoint = ts.URL

	defer ts.Close()

	session := app.CreateSession(COMPANY_ID, COMPANY_TOKEN)

	data := map[string]interface{}{
		"recv_phones":   "recv_phones",
		"recv_open_ids": "recv_open_ids",
		"content":       "content",
	}

	res, err := session.SendSms(data)
	if err != nil {
		t.Fatalf("cannot send sms. [e:%v]", err)
	}

	if res.Get("ret") == float64(0) {
		t.Log("sms sended")
	}
}
