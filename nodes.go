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
    CreateIndex int
    ModifyIndex int
}

type NodeService struct {
    ID string
    Service string
    Tags []string
    Address string
    Port int
    EnableTagOverride bool
    CreateIndex int
    ModifyIndex int
}

type NodeWithServices struct {
    Node Node
    Services map[string]NodeService
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

func node(nodeName string) (output string) {
    var outputObj NodeWithServices
    outputObj.Node = unmarshalNode( GetRedis("/catalog/" + nodeName) )
    outputObj.Services = make(map[string]NodeService)

    for _,key := range GlobRedis("/catalog/"+ nodeName + "/*") {
        var nodeService NodeService
        serviceObj := unmarshalService( GetRedis(key) )

        nodeService.ID = serviceObj.ServiceID
        nodeService.Service = serviceObj.ServiceName
        nodeService.Tags = serviceObj.ServiceTags
        nodeService.Address = serviceObj.ServiceAddress
        nodeService.Port = serviceObj.ServicePort
        nodeService.EnableTagOverride = serviceObj.ServiceEnableTagOverride
        nodeService.CreateIndex = serviceObj.CreateIndex
        nodeService.ModifyIndex = serviceObj.ModifyIndex

        outputObj.Services[nodeService.ID] = nodeService
    }

    outputBytes,_ := json.Marshal(outputObj)
    output = string(outputBytes)
    return
}
