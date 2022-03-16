package main
import (
   
    "fmt"
    "math/rand"//math下的包
    "time"
)

func main(){

//我们为了生成一个随机数，还需要个rand设置一个种子
//time.Now().Unix()
// rand.Seed(time.Now().Unix())
//随机生成1到100的数
//n := rand.Intn(100)+1//[0 100)不包括100

var a int = 0



for {
    rand.Seed(time.Now().UnixNano())//UnixNano微秒
    n := rand.Intn(100)+1//[0 100)不包括100
    a++ 
  if n == 99 {

	break

  }
  
}
      fmt.Println("生成99 一共用了",a)



 // yi://设置一个标签
  for i := 1; i < 4 ; i++{
     er://标签

	 for j := 1; j < 4  ; j++{

          if j == 2{

             // break
             //break yi
			 break er
		  }
		  fmt.Println("i=",i)
     fmt.Println("j=",j)

	 }

	 


  }







}




