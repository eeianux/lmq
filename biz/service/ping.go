package service

import (
	"context"
	"fmt"
	"github.com/eeianux/lmq/biz/model/eeianux/lmq"
)

type PingService struct{}

func (p PingService) Run(ctx context.Context, req *lmq.PingReq) (*lmq.PingResp, error) {
	return &lmq.PingResp{Msg: fmt.Sprintf("pong! msg:%s", req.Msg)}, nil
}
