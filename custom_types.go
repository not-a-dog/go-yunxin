package yunxin

import (
	"encoding/json"
	"reflect"

	"github.com/gorilla/schema"
)

type StringSlice []string

func JSONFormEncode(v reflect.Value) string {
	value := v.Interface()
	data, _ := json.Marshal(value)
	return string(data)
}

func LoadCustomTypes(e *schema.Encoder) {
	e.RegisterEncoder(StringSlice{}, JSONFormEncode)
	e.RegisterEncoder(AntispamCustom{}, JSONFormEncode)
}
