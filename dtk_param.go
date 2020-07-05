package dtk_sdk

import (
	"net/url"
	"sort"
	"strings"
)

type DtkParams map[string]string

func NewParam() *DtkParams {
	return &DtkParams{}
}

func (this *DtkParams) Set(key, value string) {
	(*this)[key] = value
}

func (this *DtkParams) Get(key, value string) string {
	return (*this)[key]
}

func (this *DtkParams) queryEscape(s string, encode bool) string {
	if encode {
		return url.QueryEscape(s)
	}
	return s
}

func (this *DtkParams) String(encode bool) string {
	// 对参数排序，按ASCII升序排序
	var keys []string
	for key, _ := range *this {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	// 组织参数key1=value1&key2=value2&...
	var buffer strings.Builder
	for _, key := range keys {
		if buffer.Len() > 0 {
			buffer.WriteByte('&')
		}
		buffer.WriteString(this.queryEscape(key, encode))
		buffer.WriteByte('=')
		buffer.WriteString(this.queryEscape((*this)[key], encode))
	}
	return buffer.String()
}
