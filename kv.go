package main

import (
    "encoding/json"
)

type KVObj struct {
    CreateIndex string
    ModifyIndex string
    LockIndex string
    Key string
    Flags int
    Value string
}

func KV(method string, path string, body string) (string){
    var response []KVObj
    var o KVObj
    o.CreateIndex = "100"
    o.ModifyIndex = "200"
    o.LockIndex = "200"
    o.Key = path
    o.Flags = 0

    switch method {
    case "GET":
        o.Value = GetRedis(path)
    case "PUT":
        return PutRedis(path, body)
    // case "DELETE":
    //     return DelRedis(path)
    }

    if len(o.Value) == 0 {
        return o.Value
    }

    j, _ := json.Marshal( append(response, o) )
    return string(j)
}
