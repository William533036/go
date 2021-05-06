package main

import (
	"fmt"
	"time"
)

func main() {
	go a()
	for i :=1;i<6;i++{
		fmt.Println(i)
		time.Sleep(time.Millisecond)
	}


}

func a()  {

	defer b()
	panic("A test panic!")

}

func b()  {
	if demo := recover();demo != nil{
		fmt.Println("Recover panic demo:", demo)
	}
}