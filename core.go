package wxopenapi

import (
	"fmt"
	"path"
	"runtime"
	"strings"

	"github.com/thelark/request"
)

const (
	wxOpenUrl = "open.weixin.qq.com"
)

var wxRequest = request.New(wxOpenUrl)

type response interface {
	Error() error
}

type ErrorReturn struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

func (err *ErrorReturn) Error() error {
	if err != nil {
		if err.ErrCode != 0 {
			return fmt.Errorf(err.ErrMsg)
		}
	}
	return nil
}
func checkError(rsp response) error {
	return rsp.Error()
}

type api interface {
	set(k, v string)
}

type option func(api)

func WithAppID(appID string) option {
	return func(self api) {
		self.set("AppID", appID)
	}
}
func WithComponentAppID(componentAppID string) option {
	return func(self api) {
		self.set("ComponentAppID", componentAppID)
	}
}
func WithComponentAppSecret(componentAppSecret string) option {
	return func(self api) {
		self.set("ComponentAppSecret", componentAppSecret)
	}
}
func WithComponentAccessToken(componentAccessToken string) option {
	return func(self api) {
		self.set("ComponentAccessToken", componentAccessToken)
	}
}

func WithRedirectUri(redirectUri string) option {
	return func(self api) {
		self.set("RedirectUri", redirectUri)
	}
}
func WithScope(scope string) option {
	return func(self api) {
		self.set("Scope", scope)
	}
}
func WithState(state string) option {
	return func(self api) {
		self.set("State", state)
	}
}

// 根据文件名称获取请求路由
func getBasePath() string {
	_, file, _, ok := runtime.Caller(1)
	if ok {
		path := strings.TrimSuffix(path.Base(file), path.Ext(file))
		path = strings.ReplaceAll(path, ".", "/")
		return path
	}
	return ""
}
