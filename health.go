package main

import (
    "encoding/json"
    "log"
    "strings"
    "os"
)

type Check struct {
    Node string
    CheckID string
    Name string
    Status string
    Notes string
    Output string
    ServiceID string
    ServiceName string
    CreateIndex int
    ModifyIndex int
}

type CheckRequest struct {
    ID string
    Name string
    Notes string
    DeregisterCriticalServiceAfter string
    Script string
    DockerContainerID string
    Shell string
    HTTP string
    TCP string
    Interval string
    TTL string
}

func unmarshalCheckRequest(j string)(o CheckRequest) {
    err := json.Unmarshal([]byte(j), &o)
    if err != nil {
        log.Println("error: ", err)
    }
    return
}

func unmarshalCheck(j string)(o Check) {
    err := json.Unmarshal([]byte(j), &o)
    if err != nil {
        log.Println("error: ", err)
    }
    return
}


func HealthChecks(method string, path string, body string)(resp string) {
    switch {
    case method == "PUT" && path == "check/register":
        resp = addCheck(body)
    case path == "checks":
        resp = getChecks ()
    }

    return
}

func gleanServiceID(checkID string)(string) {
    cSplit := strings.SplitAfterN(checkID, ":", 2)
    return cSplit[len(cSplit) -1 ]
}

func addCheck(body string)(string) {
    log.Println(body)

    var check Check
    c:= unmarshalCheckRequest(body)
    h,_ := os.Hostname()

    check.Node = h
    check.CheckID = c.ID
    check.Name = c.Name
    check.Status = "passing"
    check.Notes = c.Notes
    check.Output = "initialise"
    check.ServiceID = gleanServiceID(c.ID)
    check.ServiceName = "fuck knows"
    check.CreateIndex = 0
    check.ModifyIndex = 0

    o,_ := json.Marshal(check)
    PutRedis( agentKeyName("check", check.CheckID), string(o) )

    return "true"
}

func getChecks()(string) {
    var checks = make(map[string]Check)

    for _,key := range GlobRedis(agentKeyName("check", "*")) {
        checkObj := unmarshalCheck( GetRedis(key) )
        checks[checkObj.CheckID] = checkObj
    }

    if len(checks) == 0 {
        return nil
    }

    o,_ := json.Marshal(checks)
    return string(o)
}
