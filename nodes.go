package main

import (
    "encoding/json"
    "log"
    "strings"
)

type Node struct {
    Node string
    Address string
    TaggedAddresses map[string]string
}

func unmarshalNode(j string) (n Node) {
    err := json.Unmarshal([]byte(j), &n)
    if err != nil {
        log.Println("error: ", err)
    }
    return
}


func nodes(path string, body string) (output string) {
    var nodeList []Node

    for _,key := range GlobRedis("/catalog/*") {
        if strings.Count(key, "/") > 2 {
            continue
        }

        nodeObj := unmarshalNode( GetRedis(key) )
        nodeList = append(nodeList, nodeObj)
    }

    outputBytes,_ := json.Marshal(nodeList)
    output = string(outputBytes)
    return

}
