// A Tecent Business QQ client in go.
// https://github.com/heroicyang/bqq-go
//
// @author Herioc Yang <me@heroicyang.com>

package bqq

import (
	"encoding/json"
	"fmt"
)

// API call result
type Result struct {
	Ret  int                    `json:"ret"`
	Msg  string                 `json:"msg"`
	Data map[string]interface{} `json:"data"`
}

// Make a result from tencent API response.
func MakeResult(jsonBytes []byte) (res Result, err error) {
	err = json.Unmarshal(jsonBytes, &res)

	if err != nil {
		fmt.Errorf("cannot format tencent response. %v", err)
		return
	}

	return
}
