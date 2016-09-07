package main

import (
    "encoding/json"
//    "strings"
)

func GetRedis(path string) (r string) {
    r,_ = myRedis.Get(path).Result()
    return
}

func PutRedis(path string, value string) (string) {
    myRedis.Set(path, value, 0).Result()
    return "true"
}

func PutRedisObj(path string, o interface{}) (string) {
    j,_ := json.Marshal(o)
    return PutRedis(path, string(j))
}

func DelRedis(path string) (string) {
    myRedis.Del(path).Result()
    return "true"
}

func GlobRedis(path string) (items []string) {
    items,_ = myRedis.Keys(path).Result()
    return
}
