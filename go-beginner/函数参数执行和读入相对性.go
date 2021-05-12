package main

import (
	"bytes"

	"sort"

)

type Value1 map[string]interface{}
func (v Value1) Get(key string) interface{} {
	if v == nil {
		return ""
	}
	vs := v[key]

	return vs	//不加bool判断有没有这个key？
}

func (v Value1) Set(key string, value interface{}) {
	v[key] = value
}

func (v Value1) Encode() string {
	if v == nil {
		return ""
	}
	var buf bytes.Buffer
	keys := make([]string, 0, len(v))
	for k := range v {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		prefix :=k + "="
		if buf.Len() > 0 {
			buf.WriteByte('&')
		}
		buf.WriteString(prefix)

	}
	return buf.String()
}

func main() {
	v:=Value1{}
	v.Set("sign", v.Encode()) //
}
func (v Value1) name() string {
	v["sfd"]=101001
	return "sljfd"
}
