package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/go-redis/redis"
)

// 声明一个全局的rdb变量
var rdb *redis.Client

// 初始化连接
func initClient() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err = rdb.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}

func main() {
	initClient()
	rand.Seed(time.Now().UnixNano())
	setstring(2)
}

func setInt() {
	for i := 0; i < 10000; i++ {
		lucky := rand.Intn(100)
		str := strconv.Itoa(i)
		setInfo := rdb.Set(str, lucky, 0)
		if setInfo != nil {
			fmt.Printf("setInfo:%v\n", setInfo)

		}
	}
}

func setstring(n int) {
	value := ""
	for n > 0 {
		value += "a"
		n--
	}
	for i := 0; i < 10000; i++ {

		str := strconv.Itoa(i)
		setInfo := rdb.Set(str, value, 0)
		if setInfo != nil {
			fmt.Printf("setInfo:\n")

		}
	}
}
