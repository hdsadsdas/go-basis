package main
import (
	"fmt"
)


func main(){

	//管道的初始化

	var intChan chan int
	intChan = make(chan int, 3)


    n1 := 10

	//存入管道
	intChan <- n1

	fmt.Println("intChan长度",len(intChan))

	//取出

	n2 := <- intChan

	fmt.Println(n2)
   //关闭管道
   //关闭管道过后  只能取出 不能 存入
   close(intChan)

   intchan2 := make(chan int , 20)
   
   //将20个数存入管道
   for i := 1 ; i <= 20; i++{

	    intchan2 <- i

   } 

       //遍历管道
    //遍历时一定要关闭管道不然最后会报错
	
	close(intchan2)
   for v := range intchan2{
	   fmt.Println(v)
   }


}