package yunxin

import (
	"encoding/json"
	"reflect"

	"github.com/gorilla/schema"
)

type StringSlice []string

func StringSliceEncoder(v reflect.Value) string {
	ss := v.Interface().(StringSlice)
	data, _ := json.Marshal(ss)
	return string(data)
}

func LoadCustomTypes(e *schema.Encoder) {
	e.RegisterEncoder(StringSlice{}, StringSliceEncoder)
}
