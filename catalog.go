package main

import (
//    "strings"
    "encoding/json"
    "log"
//    "regexp"
)

type Catalog struct {
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

type Node struct {
    Node string
    Address string
    TaggedAddresses map[string]string
}

type Service struct {
    Node string
    Address string
    ServiceID string
    ServiceName string
    ServiceTags []string
    ServiceAddress string
    ServicePort int
}

func serviceRoute(method string, path string, body string) (output string) {
    return
}

func nodeRoute(method string, path string, body string) (output string) {
    return
}

func register(path string, body string) (string) {
    var catalog Catalog

    err := json.Unmarshal([]byte(body), &catalog)
    if err != nil {
        log.Println("error: ", err)
    }

    var node Node
    var serv Service

    node.Node = catalog.Node
    node.Address = catalog.Address
    node.TaggedAddresses = catalog.TaggedAddresses

    serv.Node = catalog.Node
    serv.Address = catalog.Address
    serv.ServiceID = catalog.Service.ID
    serv.ServiceName = catalog.Service.Service
    serv.ServiceTags = catalog.Service.Tags
    serv.ServiceAddress = catalog.Service.Address
    serv.ServicePort = catalog.Service.Port

    nodePath := "/catalog/"+node.Node
    PutRedisObj(nodePath, node)
    PutRedisObj(nodePath+"/"+serv.ServiceName, serv)

    return "true"
}

func CatalogRoutes(method string, path string, body string) (output string) {
    if path == "register" {
        output = register(path, body)
    }
    return
}
