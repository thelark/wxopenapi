package wxopenapi

import (
	"reflect"
)

type cgiBin struct {
	ComponentAppID       string
	ComponentAppSecret   string
	ComponentAccessToken string
}

func (t *cgiBin) set(k, v string) {
	_value := reflect.ValueOf(t).Elem()
	_type := reflect.TypeOf(t).Elem()
	if _, ok := _type.FieldByName(k); ok {
		_field := _value.FieldByName(k)
		_field.SetString(v)
	}
}
func (t *cgiBin) Component(opts ...option) *cgiBinComponent {
	self := &cgiBinComponent{}
	for _, opt := range opts {
		opt(self)
	}
	return self
}
