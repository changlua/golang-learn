package main

import (
	"fmt"
	"go-redisLearn/cache"
	"time"
)

func main()  {
	cache.RCSet("changlu", 123, 30*time.Minute)
	fmt.Println(cache.RCExists("changlu"))
	get, err := cache.RCGet("changlu")
	if err != nil {
		panic(err)
	}
	fmt.Println(get)
}