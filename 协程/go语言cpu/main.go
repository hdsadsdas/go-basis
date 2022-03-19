package main

import (
	"fmt"
	"runtime"
)

func main (){

     //获取cpu数量

	 num := runtime.NumCPU()
	 fmt.Println("cpu数目为",num)

	 //可以自己设置使用多少个cpu
	 runtime.GOMAXPROCS(num-1)
	 fmt.Println("ok")

}