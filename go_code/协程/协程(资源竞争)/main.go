package main

import (
	"fmt"
	"sync" //加锁的包
	"time"
)

//全局变量
var (

	haha  = make(map[int]int)   

    lock  sync.Mutex
)


func text(n int){

	tes := 1
   for i := 1 ; i<=n ; i++{
       tes *= i 
   }
   
   //加锁
   lock.Lock()

  haha[n] = tes
   //解锁
   lock.Unlock()
}


func main (){

  //开启20个协程
	for i := 1 ; i <= 20 ; i++{
 
		go text(i)
      
	}
   //休眠10秒
	time.Sleep(time.Second * 5)

    //遍历haha
	lock.Lock()
     for i,j := range haha{

		fmt.Printf("map[%d]=%d\n",i,j)
	 }
	 lock.Unlock()
}