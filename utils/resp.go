package utils

import "github.com/eeianux/lmq/biz/model/eeianux/base"

const errCode = 400

func BuildResp(err error) *base.BaseResp {
	return &base.BaseResp{
		LogId:   nil,
		Status:  errCode,
		Message: err.Error(),
	}
}
