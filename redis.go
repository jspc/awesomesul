package main

import (
    "encoding/base64"
)

func GetRedis(path string) (r string) {
    r,_ = myRedis.Get(path).Result()
    return
}

func PutRedis(path string, value string) (string) {
    encValue := base64.StdEncoding.EncodeToString([]byte(value))
    myRedis.Set(path, encValue, 0).Result()
    return "true"
}
