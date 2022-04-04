package main

import "fmt"

func main()  {
	//管道可声明只取或只存

	//声明只取 //用于 形参的类型
	//var intchan chan<- int
	//intchan = make(chan int,3)
	//haha:= <-intchan 报错只能取不能存 

	//声明只存//用于 形参的类型
	//var intchan2 <-chan int
	// intchan2<- 10 报错只能存不能取 

  
	//在传统方法遍历时，如果不关闭管道会报错
  
	intchan := make(chan int, 10)
	for i := 1; i<=10;i++{

		intchan<- i

	} 

	 stchan := make(chan string , 5 )
	 
	 for i := 1 ; i <=5 ; i++{
              // 拼接字符串
		stchan <- "hello" + fmt.Sprintf("%d",i)

	 } 

	 var a bool

	 for {
      //一个case 里取不到跳到另一个case取不用关闭到全都取不到
		select{
				case v := <-intchan :
					fmt.Println("从intchan取到的",v)

				case v := <-stchan :
					fmt.Println("从stchan取到的",v)	

				default :
				fmt.Println("都取不到了")
				a = true

		}
  
		if a {
			break
		}

	 }


}