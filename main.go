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

var c redis.Conn
var err error
var rdbOptions = redis.DialPassword(redisPass)

// var rdb *redis.Client

type user struct {
	email     string `redis:"email"`
	firstName string `redis:"firstName"`
	lastName  string `redis:"lastName"`
}

func init() {
	c, err = redis.Dial("tcp", redisAddr)
	if err != nil {
		panic(err)
	}
}

func main() {
	// defer c.Close()

	// n, err := c.Do("HMSET", "brandon@byitkc.com", "firstName", "Brandon", "lastName", "Young", "role", "Administrator")
	// if err != nil {
	// 	panic(err)
	// }
	// a := n.(string)
	// if a != "OK" {
	// 	panic(a)
	// }
	// // var value1 string
	// // var value2 string
	// fmt.Println(a)
	// r, err := redis.Values(c.Do("HGETALL", "brandon@byitkc.com"))
	// var u user
	// redis.ScanStruct(r, &u)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(u)
	ExampleArgs()
}

func ExampleArgs() {
	if err != nil {
		fmt.Println(err)
		return
	}
	defer c.Close()

	var p1, p2 struct {
		Title  string `redis:"title"`
		Author string `redis:"author"`
		Body   string `redis:"body"`
	}

	p1.Title = "Example"
	p1.Author = "Gary"
	p1.Body = "Hello"

	if _, err := c.Do("HMSET", redis.Args{}.Add("id1").AddFlat(&p1)...); err != nil {
		fmt.Println(err)
		return
	}

	m := map[string]string{
		"title":  "Example2",
		"author": "Steve",
		"body":   "Map",
	}

	if _, err := c.Do("HMSET", redis.Args{}.Add("id2").AddFlat(m)...); err != nil {
		fmt.Println(err)
		return
	}

	for _, id := range []string{"id1", "id2"} {

		v, err := redis.Values(c.Do("HGETALL", id))
		if err != nil {
			fmt.Println(err)
			return
		}

		if err := redis.ScanStruct(v, &p2); err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("%+v\n", p2)
	}

	// Output:
	// {Title:Example Author:Gary Body:Hello}
	// {Title:Example2 Author:Steve Body:Map}
}
