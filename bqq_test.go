// A Tecent Business QQ client in go.
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
			Ret: 1,
			Msg: "code required",
		}
	} else if r.FormValue("state") == "" {
		res = &Result{
			Ret: 1,
			Msg: "state required",
		}
	} else {
		res = &Result{
			Ret: 0,
			Data: map[string]interface{}{
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
			Ret: 1,
			Msg: "code required",
		}
	} else if r.FormValue("state") == "" {
		res = &Result{
			Ret: 1,
			Msg: "state required",
		}
	} else {
		res = &Result{
			Ret: 0,
			Data: map[string]interface{}{
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
	var res *Result

	if !checkAccess(r) {
		res = &Result{
			Ret: 1,
			Msg: "auth required",
		}
	} else {
		res = &Result{
			Ret: 0,
			Data: map[string]interface{}{
				"company_name":     "Company Name",
				"company_fullname": "Company FullName",
			},
		}
	}

	result, _ := json.Marshal(res)

	w.Header().Set("Content-Type", "application/json")
	w.Write(result)
}

func departmentsListHandler(w http.ResponseWriter, r *http.Request) {
	var res *Result

	if !checkAccess(r) {
		res = &Result{
			Ret: 1,
			Msg: "auth required",
		}
	} else {
		res = &Result{
			Ret: 0,
			Data: map[string]interface{}{
				"timestamp": 1408414724,
				"items": [1]map[string]interface{}{
					map[string]interface{}{
						"dept_id":   74537005,
						"p_dept_id": 1136750599,
						"dept_name": "test",
					},
				},
			},
		}
	}

	result, _ := json.Marshal(res)

	w.Header().Set("Content-Type", "application/json")
	w.Write(result)
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
	ts := httptest.NewServer(http.StripPrefix("/oauth2/token", http.HandlerFunc(tokenHandler)))
	app := Init(APP_ID, APP_SECRET)

	app.BaseEndPoint = ts.URL

	defer ts.Close()

	res, err := app.GetAccessToken("code", "state")
	if err != nil {
		t.Fatalf("cannot get access token. [e:%v]", err)
	}

	if res.Ret == 0 {
		t.Logf("access token is '%v'", res.Data)
	}
}

func TestCompanyToken(t *testing.T) {
	ts := httptest.NewServer(http.StripPrefix("/oauth2/companyToken", http.HandlerFunc(tokenHandler)))
	app := Init(APP_ID, APP_SECRET)

	app.BaseEndPoint = ts.URL

	defer ts.Close()

	res, err := app.GetCompanyToken("code", "state")
	if err != nil {
		t.Fatalf("cannot get company token. [e:%v]", err)
	}

	if res.Ret == 0 {
		t.Logf("company token is '%v'", res.Data["company_token"])
	}
}

func TestCompanyInfoAPI(t *testing.T) {
	ts := httptest.NewServer(http.StripPrefix("/api/corporation/get", http.HandlerFunc(companyInfoHandler)))
	app := Init(APP_ID, APP_SECRET)

	app.BaseEndPoint = ts.URL

	defer ts.Close()

	session := app.CreateSession(COMPANY_ID, COMPANY_TOKEN)

	res, err := session.GetCompanyInfo()
	if err != nil {
		t.Fatalf("cannot get company info. [e:%v]", err)
	}

	if res.Ret == 0 {
		t.Logf("company info '%v'", res.Data)
	}
}

func TestDepartmentsListAPI(t *testing.T) {
	ts := httptest.NewServer(http.StripPrefix("/api/dept/list", http.HandlerFunc(departmentsListHandler)))
	app := Init(APP_ID, APP_SECRET)

	app.BaseEndPoint = ts.URL

	defer ts.Close()

	session := app.CreateSession(COMPANY_ID, COMPANY_TOKEN)

	res, err := session.GetDepartments(0)
	if err != nil {
		t.Fatalf("cannot get department list. [e:%v]", err)
	}

	if res.Ret == 0 {
		t.Logf("department list '%v'", res.Data)
	}
}
