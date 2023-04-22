namespace go eeianux.lmq
include "base.thrift"

struct PingReq{
    1: string msg (api.body="msg")
}

struct PingResp{
    1: string msg (api.body="msg")
    255: optional base.BaseResp baseResp
}

struct ProduceReq{
    1: required string topic (api.body="topic")
    2: optional string tag (api.body="tag")
    255: optional base.BaseResp baseResp
}

struct ProduceResp{
    1: optional i64 msgId (api.body="msg_id")
    255: optional base.BaseResp baseResp
}

struct ConsumeReq{
    1: optional string topic (api.body="topic")
    2: optional string tag (api.body="tag")
    255: optional base.BaseResp baseResp
}

struct ConsumeResp{
    255: optional base.BaseResp baseResp
}

service LMQService{
    PingResp Ping (1: PingReq request) (api.post="/ping")
    ProduceResp Produce(1: ProduceReq req) (api.post="/produce")
    ConsumeResp Consume(1: ConsumeReq req) (api.post="/consume")
}