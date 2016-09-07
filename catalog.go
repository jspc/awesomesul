package main

import (
    "encoding/json"
    "log"
    "strings"
)

type CatalogRegister struct {
    Datacenter string
    Node string
    Address string
    TaggedAddresses map[string]string
    Service struct {
        ID string
        Service string
        Tags []string
        Address string
        Port int
    }
    Check struct {
        Node string
        CheckID string
        Name string
        Notes string
        Status string
        ServiceID string
    }
}

type CatalogDeregister struct {
    Datacenter string
    Node string
    ServiceID string
    CheckID string
}

func register(path string, body string) (string) {
    var catalog CatalogRegister

    err := json.Unmarshal([]byte(body), &catalog)
    if err != nil {
        log.Println("error: ", err)
    }

    var node Node
    var serv Service

    node.Node = catalog.Node
    node.Address = catalog.Address
    node.TaggedAddresses = catalog.TaggedAddresses
    node.CreateIndex = 12345
    node.ModifyIndex = 12345

    serv.Node = catalog.Node
    serv.Address = catalog.Address
    serv.ServiceID = catalog.Service.ID
    serv.ServiceName = catalog.Service.Service
    serv.ServiceTags = catalog.Service.Tags
    serv.ServiceAddress = catalog.Service.Address
    serv.ServicePort = catalog.Service.Port
    serv.ServiceEnableTagOverride = false
    serv.CreateIndex = 12345
    serv.ModifyIndex = 12345

    nodePath := "/catalog/"+node.Node
    PutRedisObj(nodePath, node)
    PutRedisObj(nodePath+"/"+serv.ServiceName, serv)

    return "true"
}

func deregister(path string, body string) (string) {
    var catalog CatalogDeregister

    err := json.Unmarshal([]byte(body), &catalog)
    if err != nil {
        log.Println("error: ", err)
    }

    nodePath := "/catalog/"+catalog.Node

    if len(catalog.ServiceID) > 0 {
        DelRedis(nodePath+"/"+catalog.ServiceID)
    } else if len(catalog.CheckID) > 0 {
        // We don't do anything with checks yet
    } else {
        DelRedis(nodePath)
        for _,item := range GlobRedis(nodePath+"/*") {
            DelRedis(item)
        }
    }

    return "true"
}

func CatalogRoutes(method string, path string, body string) (output string) {
    switch path {
    case "register":
        output = register(path, body)
    case "deregister":
        output = deregister(path, body)
    case "services":
        output = services(path, body)
    case "nodes":
        output = nodes(path, body)
    default:
        splitCatalogPath := strings.SplitAfterN(path, "/", 2)
        log.Println(splitCatalogPath[0])
        log.Println(splitCatalogPath[1])

        switch splitCatalogPath[0] {
        case "service/":
            output = service(splitCatalogPath[1])
        case "node/":
            output = node(splitCatalogPath[1])
        }
    }
    return
}
