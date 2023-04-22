namespace go eeianux.base

struct Base{
    1: optional string LogId
}

struct BaseResp{
    1: optional string LogId
    2: i64 status
    3: string message
}