package main

import (
	"fmt"
	"sync"
)

type WebConfig struct {
	Port int
} 

var demo *WebConfig

var once sync.Once

func GetConfig() *WebConfig {
	//go提供了内置的方法，用来创建单例方法，通过atomic 原子包达到加锁的效果
	once.Do(func() {
		demo = &WebConfig{Port: 8080}
	})
	return demo
}

func main()  {
	c := GetConfig()
	c2 := GetConfig()

	c.Port = 9090
	fmt.Println(c == c2, c2)
}