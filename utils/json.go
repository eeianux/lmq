package utils

import "github.com/bytedance/sonic"

func MarshalString(obj interface{}) string {
	content, _ := sonic.MarshalString(obj)
	return content
}
