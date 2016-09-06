package main

import (
    "flag"
    "gopkg.in/redis.v4"
    "log"
    "net/http"
)

var myRedis *redis.Client
var redisAddr string
var redisPassword string
var redisDB int

func init() {
    flag.StringVar(&redisAddr, "address", "localhost:6379", "Redis host address")
    flag.StringVar(&redisPassword, "password", "", "Redis password, should one exist")
    flag.IntVar(&redisDB, "db", 0, "Redis database number")
    flag.Parse()

    myRedis = redis.NewClient(&redis.Options{
        Addr:     redisAddr,
        Password: redisPassword,
        DB:       redisDB,
    })

    log.Println("%T", myRedis)

    pong, err := myRedis.Ping().Result()
    if err != nil {
        log.Fatal(err)
    }
    log.Println(pong)

}

func main(){
    http.HandleFunc("/", router)
    http.ListenAndServe(":8000", nil)
}
