package main

import (
    "encoding/json"
    "log"
)

func GetRedis(path string) (r string) {
    r,_ = myRedis.Get(path).Result()
    return
}

func PutRedis(path string, value string) (string) {
    _,err := myRedis.Set(path, value, 0).Result()
    if  err != nil {
        log.Println(err)
    }

    return "true"
}

func PutRedisObj(path string, o interface{}) (string) {
    j,err := json.Marshal(o)
    if  err != nil {
        log.Println(err)
    }

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
