package main

import (
    "encoding/json"
    "fmt"
    "log"
    "strings"
)


type AgentRegister struct {
    ID string
    Name string
    Tags []string
    Address string
    Port int
    Check struct {
        DeregisterCriticalServiceAfter string
        Script string
        HTTP string
        Interval string
        TTL string
    }
}

// The service as returned from /agent/services is different
// to that returned from /catalog/services becat
type AgentService struct {
    ID string
    Name string
    Tags []string
    Address string
    Port int
    EnableTagOverride bool
    CreateIndex int
    ModifyIndex int
}

// FIXME: This is a crap way of doing this
func unmarshalAgentService(j string) (a AgentService) {
    err := json.Unmarshal([]byte(j), &a)
    if err != nil {
        log.Println("error: ", err)
    }
    return
}

func unmarshalAgentRegister(j string) (a AgentRegister) {
    err := json.Unmarshal([]byte(j), &a)
    if err != nil {
        log.Println("error: ", err)
    }
    return
}

func agentKeyName(recordType string, id string) (k string) {
    k = fmt.Sprintf("/agent/%s/%s", recordType, id)
    return
}

func agentRegisterService(j string) (string) {
    var service AgentService
    a := unmarshalAgentService(j)
    k := agentKeyName("service", a.ID)

    service.ID = a.ID
    service.Name = a.Name
    service.Tags = a.Tags
    service.Address = a.Address
    service.Port = a.Port
    service.EnableTagOverride = false
    service.CreateIndex = 12345
    service.ModifyIndex = 12345

    PutRedisObj(k, service)
    return "true"
}

func agentDeRegisterService(path string) (string) {
    key := strings.Split(path, "/")
    k := agentKeyName("service", key[ len(key) - 1 ])

    DelRedis(k)
    return "true"
}

func agentListChecks() (string) {
    // We're not doing any thing  with serf/ checks.
    // c'est la vie
    return "{}"
}

func agentListServices() (string) {
    var services = make(map[string]AgentService)

    for _,key := range GlobRedis(agentKeyName("service", "*")) {
        serviceObj := unmarshalAgentService( GetRedis(key) )
        services[serviceObj.ID] = serviceObj
    }

    o,_ := json.Marshal(services)
    return string(o)
}

func Agent(path string, body string) (r string) {
    switch {
    case path == "service/register":
        r = agentRegisterService(body)
    case strings.HasPrefix(path, "service/deregister"):
        r = agentDeRegisterService(path)
    case path == "checks":
        r = agentListChecks()
    case path == "services":
        r = agentListServices()
    }

    return
}
