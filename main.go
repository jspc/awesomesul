package main

import (
    "fmt"
    "gopkg.in/redis.v4"
    "net/http"
    "log"
    "strings"
    "io"
    "encoding/json"
    "encoding/base64"
    "bytes"
)

type NormalisedPath struct {
    Service string
    Path string
}

type KVObj struct {
    CreateIndex string
    ModifyIndex string
    LockIndex string
    Key string
    Flags int
    Value string
}

var myRedis *redis.Client
var resp string

func GetRedis(path string) (r string) {
    r,_ = myRedis.Get(path).Result()
    return
}

func PutRedis(path string, value string) (string) {
    encValue := base64.StdEncoding.EncodeToString([]byte(value))
    myRedis.Set(path, encValue, 0).Result()
    return "true"
}

func KV(method string, path string, body string) (string){
    var o KVObj
    o.CreateIndex = "100"
    o.ModifyIndex = "200"
    o.LockIndex = "200"
    o.Key = path
    o.Flags = 0

    switch method {
    case "GET":
        o.Value = GetRedis(path)
    case "PUT":
        return PutRedis(path, body)
    // case "DELETE":
    //     return DelRedis(path)
    }

    if len(o.Value) == 0 {
        return o.Value
    }

    j, _ := json.Marshal(o)
    return fmt.Sprintf("[%s]", string(j))
}

func rwToString(rw io.ReadCloser) (s string){
    buf := new(bytes.Buffer)
    buf.ReadFrom(rw)
    s = buf.String()

    return
}

func normalisePath(path string) (n NormalisedPath){
    cleanPath := strings.Replace(path, "/v1/", "", 1)
    splitPath := strings.SplitAfterN(cleanPath, "/", 2)

    n.Service = strings.Replace(splitPath[0], "/", "", 1)
    n.Path = splitPath[1]
    return
}

func setHeaders(w http.ResponseWriter)(wDup http.ResponseWriter){
    wDup = w
    wDup.Header().Set("Access-Control-Allow-Headers", "requested-with, Content-Type, origin, authorization, accept, client-security-token, cache-control, x-api-key")
    wDup.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, DELETE, PUT")
    wDup.Header().Set("Access-Control-Allow-Origin", "*")
    wDup.Header().Set("Access-Control-Max-Age", "10000")
    wDup.Header().Set("Cache-Control", "no-cache")

    wDup.Header().Set("Content-Type", "application/json")

    return
}

func logRequest(r *http.Request) {
    log.Printf( "%s :: %s %s",
        r.RemoteAddr,
        r.Method,
        r.URL.Path)
}

func router(w http.ResponseWriter, r *http.Request){
    logRequest(r)

    normalisedPath := normalisePath(r.URL.Path)
    normalisedBody := rwToString(r.Body)

    switch normalisedPath.Service {
    case "kv":
        resp = KV(r.Method, normalisedPath.Path, normalisedBody)
    }

    w = setHeaders(w)
    if len(resp) == 0 {
        http.NotFound(w,r)
    }
    fmt.Fprintf(w, resp)
}

func main(){
    myRedis = redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "", // no password set
        DB:       0,  // use default DB
    })

    log.Println("%T", myRedis)

    pong, err := myRedis.Ping().Result()
    log.Println(pong, err)

    http.HandleFunc("/", router)
    http.ListenAndServe(":8000", nil)
}
