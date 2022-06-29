package main

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

const (
	redisAddr string = "localhost:6379"
	redisPass string = "password"
	redisDB   int    = 0
)

var rdb redis.Conn
var err error
var rdbOptions = redis.DialPassword(redisPass)

// var rdb *redis.Client

func init() {
	rdb, err = redis.Dial("tcp", redisAddr, rdbOptions)
	if err != nil {
		panic(err)
	}
}

func main() {
	defer rdb.Close()

	n, err := rdb.Do("SET", "mykey", 15)
	if err != nil {
		panic(err)
	}
	a := n.(string)
	if a != "OK" {
		panic(a)
	}
	fmt.Println(a)

	r, err := redis.Int(rdb.Do("GET", "mykey"))
	if err != nil {
		panic(err)
	}
	fmt.Println(r)
}
