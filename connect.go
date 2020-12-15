package wxopenapi

import "reflect"

type connect struct {
	AppID          string
	ComponentAppID string
}

func (t *connect) set(k, v string) {
	_value := reflect.ValueOf(t).Elem()
	_type := reflect.TypeOf(t).Elem()
	if _, ok := _type.FieldByName(k); ok {
		_field := _value.FieldByName(k)
		_field.SetString(v)
	}
}
func (t *connect) OAuth2(opts ...option) *connectOAuth2 {
	self := &connectOAuth2{}
	for _, opt := range opts {
		opt(self)
	}
	return self
}
