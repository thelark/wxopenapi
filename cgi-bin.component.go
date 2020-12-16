package wxopenapi

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/thelark/request"
)

type cgiBinComponentOptionName string

const (
	LocationReport  cgiBinComponentOptionName = "location_report"  // 地理位置上报选项 (0:无上报,1:进入会话时上报,2:每 5s 上报)
	VoiceRecognize  cgiBinComponentOptionName = "voice_recognize"  // 语音识别开关选项 (0:关闭语音识别,1:开启语音识别)
	CustomerService cgiBinComponentOptionName = "customer_service" // 多客服开关选项 (0:关闭多客服,1:开启多客服)
)

type cgiBinComponentOptionValue string

const (
	OptionValue0 cgiBinComponentOptionValue = "0"
	OptionValue1 cgiBinComponentOptionValue = "1"
	OptionValue2 cgiBinComponentOptionValue = "2"
)

type cgiBinComponent struct {
	ComponentAppID         string                     `json:"component_appid,omitempty"`          // 第三方平台 appid
	ComponentAppSecret     string                     `json:"component_appsecret,omitempty"`      //
	ComponentVerifyTicket  string                     `json:"component_verify_ticket,omitempty"`  //
	ComponentAccessToken   string                     `json:"component_access_token,omitempty"`   // 第三方平台component_access_token，不是authorizer_access_token
	AuthorizationCode      string                     `json:"authorization_code,omitempty"`       // 授权码, 会在授权成功时返回给第三方平台
	AuthorizerAppID        string                     `json:"authorizer_appid,omitempty"`         // 授权方 appid
	AuthorizerRefreshToken string                     `json:"authorizer_refresh_token,omitempty"` // 刷新令牌，获取授权信息时得到
	OptionName             cgiBinComponentOptionName  `json:"option_name,omitempty"`              // 选项名称
	OptionValue            cgiBinComponentOptionValue `json:"option_value,omitempty"`             // 设置的选项值
	Offset                 int                        `json:"offset,omitempty"`                   // 偏移位置/起始位置
	Count                  int                        `json:"count,omitempty"`                    // 拉取数量，最大为 500
}

func (t *cgiBinComponent) set(k, v string) {
	_value := reflect.ValueOf(t).Elem()
	_type := reflect.TypeOf(t).Elem()
	if _, ok := _type.FieldByName(k); ok {
		_field := _value.FieldByName(k)
		_field.SetString(v)
	}
}

type cgiBinComponentApiComponentToken struct {
	*ErrorReturn
	ComponentAccessToken string `json:"component_access_token"`
	ExpiresIn            int    `json:"expires_in"`
}

// ApiComponentToken 令牌（component_access_token）是第三方平台接口的调用凭据。令牌的获取是有限制的，每个令牌的有效期为 2 小时，请自行做好令牌的管理，在令牌快过期时（比如1小时50分），重新调用接口获取。
// POST https://api.weixin.qq.com/cgi-bin/component/api_component_token
func (t *cgiBinComponent) ApiComponentToken() (*cgiBinComponentApiComponentToken, error) {
	rsp := new(cgiBinComponentApiComponentToken)
	body, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	if err := wxRequest.Get(
		fmt.Sprintf("%s/api_component_token", getBasePath()),
		request.WithBody(body),
		request.WithResponse(&rsp),
	); err != nil {
		return nil, err
	}
	if err := checkError(rsp); err != nil {
		return nil, err
	}
	t.set("ComponentAccessToken", rsp.ComponentAccessToken)
	return rsp, nil
}

type cgiBinComponentApiCreatePreAuthCode struct {
	*ErrorReturn
	PreAuthCode string `json:"pre_auth_code"`
	ExpiresIn   int    `json:"expires_in"`
}

// ApiCreatePreauthcode 预授权码（pre_auth_code）是第三方平台方实现授权托管的必备信息，每个预授权码有效期为 1800秒。需要先获取令牌才能调用。使用过程中如遇到问题，可在开放平台服务商专区发帖交流。
// POST https://api.weixin.qq.com/cgi-bin/component/api_create_preauthcode?component_access_token=COMPONENT_ACCESS_TOKEN
func (t *cgiBinComponent) ApiCreatePreAuthCode() (*cgiBinComponentApiCreatePreAuthCode, error) {
	rsp := new(cgiBinComponentApiCreatePreAuthCode)
	body, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	if err := wxRequest.Get(
		fmt.Sprintf("%s/api_create_preauthcode", getBasePath()),
		request.WithParam("component_access_token", t.ComponentAccessToken),
		request.WithBody(body),
		request.WithResponse(&rsp),
	); err != nil {
		return nil, err
	}
	if err := checkError(rsp); err != nil {
		return nil, err
	}
	return rsp, nil
}

type cgiBinComponentApiQueryAuth struct {
	*ErrorReturn
	AuthorizationInfo struct {
		AuthorizerAppID        string `json:"authorizer_appid"`
		AuthorizerAccessToken  string `json:"authorizer_access_token"`
		ExpiresIn              int    `json:"expires_in"`
		AuthorizerRefreshToken string `json:"authorizer_refresh_token"`
		FuncInfo               []struct {
			FuncscopeCategory struct {
				ID int `json:"id"`
			} `json:"funcscope_category"`
		} `json:"func_info"`
	} `json:"authorization_info"`
}

// ApiQueryAuth 当用户在第三方平台授权页中完成授权流程后，第三方平台开发者可以在回调 URI 中通过 URL 参数获取授权码。使用以下接口可以换取公众号/小程序的授权信息。建议保存授权信息中的刷新令牌（authorizer_refresh_token）。使用过程中如遇到问题，可在开放平台服务商专区发帖交流。
// POST https://api.weixin.qq.com/cgi-bin/component/api_query_auth?component_access_token=COMPONENT_ACCESS_TOKEN
func (t *cgiBinComponent) ApiQueryAuth() (*cgiBinComponentApiQueryAuth, error) {
	rsp := new(cgiBinComponentApiQueryAuth)
	body, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	if err := wxRequest.Get(
		fmt.Sprintf("%s/api_query_auth", getBasePath()),
		request.WithParam("component_access_token", t.ComponentAccessToken),
		request.WithBody(body),
		request.WithResponse(&rsp),
	); err != nil {
		return nil, err
	}
	if err := checkError(rsp); err != nil {
		return nil, err
	}
	return rsp, nil
}

type cgiBinComponentApiAuthorizerToken struct {
	*ErrorReturn
	AuthorizerAccessToken  string `json:"authorizer_access_token"`
	ExpiresIn              int    `json:"expires_in"`
	AuthorizerRefreshToken string `json:"authorizer_refresh_token"`
}

// ApiAuthorizerToken 在公众号/小程序接口调用令牌（authorizer_access_token）失效时，可以使用刷新令牌（authorizer_refresh_token）获取新的接口调用令牌。使用过程中如遇到问题，可在开放平台服务商专区发帖交流。
// POST https://api.weixin.qq.com/cgi-bin/component/api_authorizer_token?component_access_token=COMPONENT_ACCESS_TOKEN
func (t *cgiBinComponent) ApiAuthorizerToken() (*cgiBinComponentApiAuthorizerToken, error) {
	rsp := new(cgiBinComponentApiAuthorizerToken)
	body, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	if err := wxRequest.Get(
		fmt.Sprintf("%s/api_authorizer_token", getBasePath()),
		request.WithParam("component_access_token", t.ComponentAccessToken),
		request.WithBody(body),
		request.WithResponse(&rsp),
	); err != nil {
		return nil, err
	}
	if err := checkError(rsp); err != nil {
		return nil, err
	}
	return rsp, nil
}

type cgiBinComponentApiGetAuthorizerInfo struct {
	*ErrorReturn
	AuthorizerInfo struct {
		NickName        string `json:"nick_name"`
		HeadImg         string `json:"head_img"`
		ServiceTypeInfo struct {
			ID int `json:"id"`
		} `json:"service_type_info"`
		VerifyTypeInfo struct {
			ID int `json:"id"`
		} `json:"verify_type_info"`
		UserName      string `json:"user_name"`
		PrincipalName string `json:"principal_name"`
		BusinessInfo  struct {
			OpenStore int `json:"open_store"`
			OpenScan  int `json:"open_scan"`
			OpenPay   int `json:"open_pay"`
			OpenCard  int `json:"open_card"`
			OpenShake int `json:"open_shake"`
		} `json:"business_info"`
		Alias     string `json:"alias"`
		QrcodeUrl string `json:"qrcode_url"`
	} `json:"authorizer_info"`
	AuthorizationInfo struct {
		AuthorizerAppid string `json:"authorizer_appid"`
		FuncInfo        []struct {
			FuncscopeCategory struct {
				ID int `json:"id"`
			} `json:"funcscope_category"`
		} `json:"func_info"`
	} `json:"authorization_info"`
}

// ApiGetAuthorizerInfo 该 API 用于获取授权方的基本信息，包括头像、昵称、帐号类型、认证类型、微信号、原始ID和二维码图片URL。使用过程中如遇到问题，可在开放平台服务商专区发帖交流。
// POST https://api.weixin.qq.com/cgi-bin/component/api_get_authorizer_info?component_access_token=COMPONENT_ACCESS_TOKEN
func (t *cgiBinComponent) ApiGetAuthorizerInfo() (*cgiBinComponentApiGetAuthorizerInfo, error) {
	rsp := new(cgiBinComponentApiGetAuthorizerInfo)
	body, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	if err := wxRequest.Get(
		fmt.Sprintf("%s/api_get_authorizer_info", getBasePath()),
		request.WithParam("component_access_token", t.ComponentAccessToken),
		request.WithBody(body),
		request.WithResponse(&rsp),
	); err != nil {
		return nil, err
	}
	if err := checkError(rsp); err != nil {
		return nil, err
	}
	return rsp, nil
}

type cgiBinComponentApiGetAuthorizerOption struct {
	*ErrorReturn
	AuthorizerAppID string `json:"authorizer_appid"` // 授权公众号或小程序的 appid
	OptionName      string `json:"option_name"`      // 选项名称
	OptionValue     string `json:"option_value"`     // 选项值
}

// ApiGetAuthorizerOption 本 API 用于获取授权方的公众号/小程序的选项设置信息，如：地理位置上报，语音识别开关，多客服开关。使用过程中如遇到问题，可在开放平台服务商专区发帖交流。
// POST https://api.weixin.qq.com/cgi-bin/component/api_get_authorizer_option?component_access_token=COMPONENT_ACCESS_TOKEN
func (t *cgiBinComponent) ApiGetAuthorizerOption() (*cgiBinComponentApiGetAuthorizerOption, error) {
	rsp := new(cgiBinComponentApiGetAuthorizerOption)
	body, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	if err := wxRequest.Get(
		fmt.Sprintf("%s/api_get_authorizer_option", getBasePath()),
		request.WithParam("component_access_token", t.ComponentAccessToken),
		request.WithBody(body),
		request.WithResponse(&rsp),
	); err != nil {
		return nil, err
	}
	if err := checkError(rsp); err != nil {
		return nil, err
	}
	return rsp, nil
}

type cgiBinComponentApiSetAuthorizerOption struct {
	*ErrorReturn
}

// ApiSetAuthorizerOption 本 API 用于设置授权方的公众号/小程序的选项信息，如：地理位置上报，语音识别开关，多客服开关。使用过程中如遇到问题，可在开放平台服务商专区发帖交流。
// POST https://api.weixin.qq.com/cgi-bin/component/api_set_authorizer_option?component_access_token=COMPONENT_ACCESS_TOKEN
func (t *cgiBinComponent) ApiSetAuthorizerOption() (*cgiBinComponentApiSetAuthorizerOption, error) {
	rsp := new(cgiBinComponentApiSetAuthorizerOption)
	body, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	if err := wxRequest.Get(
		fmt.Sprintf("%s/api_set_authorizer_option", getBasePath()),
		request.WithParam("component_access_token", t.ComponentAccessToken),
		request.WithBody(body),
		request.WithResponse(&rsp),
	); err != nil {
		return nil, err
	}
	if err := checkError(rsp); err != nil {
		return nil, err
	}
	return rsp, nil
}

type cgiBinComponentApiGetAuthorizerList struct {
	*ErrorReturn
	TotalCount int `json:"total_count"` // 授权的帐号总数
	List       []struct {
		AuthorizerAppID string `json:"authorizer_appid"` // 已授权的 appid
		RefreshToken    string `json:"refresh_token"`    // 刷新令牌authorizer_access_token
		AuthTime        int64  `json:"auth_time"`        // 授权的时间
	} `json:"list"` // 当前查询的帐号基本信息列表
}

// ApiGetAuthorizerList 使用本 API 拉取当前所有已授权的帐号基本信息。使用过程中如遇到问题，可在开放平台服务商专区发帖交流。
// POST https://api.weixin.qq.com/cgi-bin/component/api_get_authorizer_list?component_access_token=COMPONENT_ACCESS_TOKEN
func (t *cgiBinComponent) ApiGetAuthorizerList() (*cgiBinComponentApiGetAuthorizerList, error) {
	rsp := new(cgiBinComponentApiGetAuthorizerList)
	body, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	if err := wxRequest.Get(
		fmt.Sprintf("%s/api_get_authorizer_list", getBasePath()),
		request.WithParam("component_access_token", t.ComponentAccessToken),
		request.WithBody(body),
		request.WithResponse(&rsp),
	); err != nil {
		return nil, err
	}
	if err := checkError(rsp); err != nil {
		return nil, err
	}
	return rsp, nil
}

type cgiBinComponentClearQuota struct {
	*ErrorReturn
}

// ClearQuota 第三方平台对其所有 API 调用次数清零（只与第三方平台相关，与公众号无关，接口如 api_component_token）
// POST https://api.weixin.qq.com/cgi-bin/component/clear_quota?component_access_token=COMPONENT_ACCESS_TOKEN
func (t *cgiBinComponent) ClearQuota() (*cgiBinComponentClearQuota, error) {
	rsp := new(cgiBinComponentClearQuota)
	body, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	if err := wxRequest.Get(
		fmt.Sprintf("%s/clear_quota", getBasePath()),
		request.WithParam("component_access_token", t.ComponentAccessToken),
		request.WithBody(body),
		request.WithResponse(&rsp),
	); err != nil {
		return nil, err
	}
	if err := checkError(rsp); err != nil {
		return nil, err
	}
	return rsp, nil
}
