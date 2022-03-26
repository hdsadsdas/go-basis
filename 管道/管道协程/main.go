package main

import "fmt"

//创建一个写入管道
func xie(this chan int){
  
  	for i := 1 ; i<=50 ; i++{

		this <- i

	  }
	  //全部写完关闭管道
   close(this)
}

func du(this chan int, this2 chan bool){

	for{
        //ok表示读到最后 关闭管道
		v , ok := <- this
		if !ok {
			break
		}
      fmt.Println("读出的数据",v)
	}
	//读完数据后存入一个true
    //传完立马关闭
	this2 <- true
    close(this2)
}


func main(){

	//定义两个管道
	var intchan chan int = make(chan int, 50) 
    var intchan2 chan bool = make(chan bool, 1) 

     go xie(intchan)
     go du(intchan,intchan2)

	for {
		//读到关闭管道break
		_ ,ok := <- intchan2
		if !ok{
                   break
		}
	}


}