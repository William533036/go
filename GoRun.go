package main

import (
	"fmt"
	"time"
)

func main()  {
	for i:= 0; i<188; i++ {
		go func(i int) {
			for{
				fmt.Printf("hello from" +
					"gorutine %d\n", i)
			}
		}(i)
	}
	time.Sleep(time.Second)
}
