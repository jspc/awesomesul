package main

import (
    "encoding/json"
    "log"
)

type Service struct {
    Node string
    Address string
    ServiceID string
    ServiceName string
    ServiceTags []string
    ServiceAddress string
    ServicePort int
}


// FIXME: This, and the func below, are a crap way of doing this
func unmarshalService(j string) (s Service) {
    err := json.Unmarshal([]byte(j), &s)
    if err != nil {
        log.Println("error: ", err)
    }
    return
}

func services(path string, body string) (output string) {
    serviceList := make(map[string][]string)

    for _,key := range GlobRedis("/catalog/*/*") {
        servObj := unmarshalService( GetRedis(key) )
        serviceList[servObj.ServiceName] = servObj.ServiceTags
    }

    outputBytes,_ := json.Marshal(serviceList)
    output = string(outputBytes)
    return
}
