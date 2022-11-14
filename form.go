package yunxin

import (
	"encoding/json"
	"reflect"

	"github.com/gorilla/schema"
)

var defaultEncoder = schema.NewEncoder()

func RegisterEncoder(value interface{}, encoder func(reflect.Value) string) {
	defaultEncoder.RegisterEncoder(value, encoder)
}

func init() {
	RegisterEncoder(StringSlice{}, JSONFormEncode)
	RegisterEncoder(AntispamCustom{}, JSONFormEncode)
}

type StringSlice []string

func JSONFormEncode(v reflect.Value) string {
	value := v.Interface()
	data, _ := json.Marshal(value)
	return string(data)
}
