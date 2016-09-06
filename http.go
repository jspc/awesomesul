package main

import (
    "bytes"
    "fmt"
    "io"
    "log"
    "net/http"
    "strings"
)

type NormalisedPath struct {
    Service string
    Path string
}

var resp string

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
