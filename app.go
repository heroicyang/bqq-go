// A Tecent Business QQ client in go.
// https://github.com/heroicyang/bqq-go
//
// @author Herioc Yang <me@heroicyang.com>

package bqq

import (
	"errors"
)

// Holds app infomation
type App struct {
	AppId        string
	AppSecret    string
	RedirectUri  string
	BaseEndPoint string
	ClientIP     string
}

// Initialize a new App and sets app id and secret.
func Init(appId, appSecret string) *App {
	return &App{
		AppId:     appId,
		AppSecret: appSecret,
	}
}

// Get authorize url
func (this *App) GetAuthorizeUri(state, ui string) (url string, err error) {
	if state == "" {
		err = errors.New("state required")
		return
	}

	if ui == "" {
		ui = "auto"
	}

	params := Params{
		"response_type": "code",
		"app_id":        this.AppId,
		"redirect_uri":  this.RedirectUri,
		"state":         state,
		"ui":            ui,
	}

	session := &Session{}
	url = this.BaseEndPoint + session.getUrl("/oauth2/authorize", params)

	return
}

// Get user access token
func (this *App) GetAccessToken(code, state string) (res Result, err error) {
	if state == "" {
		err = errors.New("code required")
		return
	}

	if state == "" {
		err = errors.New("state required")
		return
	}

	params := Params{
		"grant_type":   "authorization_code",
		"app_id":       this.AppId,
		"app_secret":   this.AppSecret,
		"redirect_uri": this.RedirectUri,
		"code":         code,
		"state":        state,
	}

	session := &Session{}
	urlStr := this.BaseEndPoint + session.getUrl("/oauth2/token", params)

	var response []byte
	response, err = session.SendGetRequest(urlStr)

	if err != nil {
		return
	}

	res, err = MakeResult(response)
	return
}

// Get company access token
func (this *App) GetCompanyToken(code, state string) (res Result, err error) {
	if state == "" {
		err = errors.New("code required")
		return
	}

	if state == "" {
		err = errors.New("state required")
		return
	}

	params := Params{
		"grant_type":   "authorization_code",
		"app_id":       this.AppId,
		"app_secret":   this.AppSecret,
		"redirect_uri": this.RedirectUri,
		"code":         code,
		"state":        state,
	}

	session := &Session{}
	urlStr := this.BaseEndPoint + session.getUrl("/oauth2/companyToken", params)

	var response []byte
	response, err = session.SendGetRequest(urlStr)

	if err != nil {
		return
	}

	res, err = MakeResult(response)
	return
}

// Create a session based on current app setting
func (this *App) CreateSession(companyId, companyToken string) *Session {
	return &Session{
		companyId:    companyId,
		companyToken: companyToken,
		app:          this,
	}
}
