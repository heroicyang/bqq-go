// A Tencent Business QQ client in go.
// https://github.com/heroicyang/bqq-go
//
// @author Herioc Yang <me@heroicyang.com>

package bqq

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

// An interface to send http request
type HttpClient interface {
	Do(req *http.Request) (resp *http.Response, err error)
	Get(url string) (resp *http.Response, err error)
	Post(url string, body io.Reader) (resp *http.Response, err error)
}

// Holds a session with company id and token
type Session struct {
	companyId    string
	companyToken string
	OpenId       string
	ClientIP     string
	app          *App
	HttpClient   HttpClient
}

// Request an API with `path` the GET method, and optional `params`
func (session *Session) Get(path string, params Params) (res Result, err error) {
	urlStr := session.app.BaseEndPoint + session.getRequestUrl(path, params)

	var response []byte
	response, err = session.SendGetRequest(urlStr)

	if err != nil {
		return
	}

	res, err = MakeResult(response)
	return
}

// Request an API with `path` and `data` the POST method
func (session *Session) Post(path string, data Params) (res Result, err error) {
	urlStr := session.app.BaseEndPoint + session.getRequestUrl(path, nil)

	var response []byte
	response, err = session.SendPostRequest(urlStr, data)

	if err != nil {
		return
	}

	res, err = MakeResult(response)
	return
}

// Get generic url with `path` and optional `params`
func (session *Session) getUrl(path string, params Params) string {
	buf := &bytes.Buffer{}
	buf.WriteString(path)

	if params != nil {
		buf.WriteRune('?')
		params.Encode(buf)
	}

	return buf.String()
}

// Get requset url with `path` and optional `params`
// Encode common params for the url query string
func (session *Session) getRequestUrl(path string, params Params) string {
	if params == nil {
		params = Params{}
	}

	params["company_id"] = session.companyId
	params["company_token"] = session.companyToken
	params["app_id"] = session.app.AppId
	params["oauth_version"] = 2

	if _, exist := params["open_id"]; !exist && session.OpenId != "" {
		params["open_id"] = session.OpenId
	}

	if _, exist := params["client_ip"]; !exist {
		if session.ClientIP != "" {
			params["client_ip"] = session.ClientIP
		} else {
			params["client_ip"] = session.app.ClientIP
		}
	}

	return session.getUrl(path, params)
}

func (session *Session) SendGetRequest(url string) ([]byte, error) {
	var request *http.Request
	var err error

	request, err = http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}

	return session.sendRequest(request)
}

func (session *Session) SendPostRequest(url string, params Params) ([]byte, error) {
	buf := &bytes.Buffer{}
	params.Encode(buf)

	var request *http.Request
	var err error

	request, err = http.NewRequest("POST", url, buf)

	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	return session.sendRequest(request)
}

func (session *Session) sendRequest(request *http.Request) ([]byte, error) {
	var response *http.Response
	var err error

	if session.HttpClient == nil {
		response, err = http.DefaultClient.Do(request)
	} else {
		response, err = session.HttpClient.Do(request)
	}

	if err != nil {
		return nil, fmt.Errorf("cannot reach tencent server. %v", err)
	}

	defer response.Body.Close()

	buf := &bytes.Buffer{}
	_, err = io.Copy(buf, response.Body)

	if err != nil {
		return nil, fmt.Errorf("cannot read tencent response. %v", err)
	}

	return buf.Bytes(), nil
}
