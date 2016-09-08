package main

import (
    "encoding/base64"
    "encoding/json"
    "strings"
)

type KVObj struct {
    CreateIndex int
    ModifyIndex int
    LockIndex int
    Key string
    Flags int
    Value string
}

func KV(method string, path string, body string) (string){
    var response []KVObj
    switch {
    case strings.HasSuffix(path, "/"):
        response = recurseKV(method, path)
    case method == "GET":
        response = append(response, createKVObj(path))
    case method == "PUT":
        encBody := base64.StdEncoding.EncodeToString([]byte(body))
        return PutRedis(path, encBody)
    case method == "DELETE":
        return DelRedis(path)
    }

    j, _ := json.Marshal(response)
    return string(j)
}

func createKVObj(path string) (k KVObj) {
    k.CreateIndex = 1001
    k.ModifyIndex = 1001
    k.LockIndex = 0
    k.Key = path
    k.Flags = 0
    k.Value = GetRedis(path)

    return
}

func recurseKV(method string, path string) (kvs []KVObj) {
    for _,kv := range GlobRedis(path+"*") {
        switch method {
        case "GET":
            kvObj := createKVObj(kv)
            kvs = append(kvs, kvObj)
        case "DELETE":
            DelRedis(path)
        }
    }
    return
}
