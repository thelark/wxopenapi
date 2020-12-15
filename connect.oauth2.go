package wxopenapi

import (
	"fmt"
	"net/http"
	"reflect"
)

type connectOAuth2 struct {
	AppID          string // 公众号的 appid
	RedirectUri    string // 重定向地址，需要 urlencode，这里填写的是第三方平台的【公众号开发域名】，注意这个配置需要勾选公众号相关全权限集才可以看到
	Scope          string // 授权作用域，拥有多个作用域用逗号（,）分隔
	State          string // 重定向后会带上 state 参数，开发者可以填写任意参数值，最多 128 字节
	ComponentAppID string // 服务方的 appid，在申请创建公众号服务成功后，可在公众号服务详情页找到
}

func (t *connectOAuth2) set(k, v string) {
	_value := reflect.ValueOf(t).Elem()
	_type := reflect.TypeOf(t).Elem()
	if _, ok := _type.FieldByName(k); ok {
		_field := _value.FieldByName(k)
		_field.SetString(v)
	}
}

// REDIRECT https://open.weixin.qq.com/connect/oauth2/authorize?appid=APPID&redirect_uri=REDIRECT_URI&response_type=code&scope=SCOPE&state=STATE&component_appid=component_appid#wechat_redirect
func (t *connectOAuth2) Authorize(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, fmt.Sprintf("%s/authorize?appid=%s&redirect_uri=%s&response_type=code&scope=%s&state=%s&component_appid=%s#wechat_redirect", getBasePath(), t.AppID, t.RedirectUri, t.Scope, t.State, t.ComponentAppID), http.StatusTemporaryRedirect)
}
