package main

import (
	"fmt"
	"time"
)

func text() {

	for i := 1; i <= 10; i++ {
		
		fmt.Println("text",i)
		time.Sleep(time.Second)
	}


}


func main (){  //主线程结束就结束

	go text() //开启了一个协程

	for i := 1; i <= 10; i++ {
		
		fmt.Println("main",i)
		time.Sleep(time.Second)

	}
}