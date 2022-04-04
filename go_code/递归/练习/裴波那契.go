package main
import "fmt"

func bei(i int) int{

	if i == 1 || i == 2{  //终止条件
          
           return 1 
	}else {

           return bei(i-2)+bei(i-1)

	}


}
//猴子吃桃
func tao (n int) int{

      if n == 10 {
		 return 1
	  }else{

           return (tao(n+1)+1)*2

	  }


}




func main(){
  

	fmt.Println(bei(15))
   // fmt.Println(tao(9))

}
