package main
import (
	"fmt"
  good "go_code/bao"//别名
  
)
  
// func haha(n1 float64, n2 float64,fangfa string ) (float64){
   
//    var jiuguo float64

//     switch fangfa {
// 	case "+" :
// 		jiuguo = n1 + n2
	
// 	case "-" :
// 		jiuguo = n1 - n2	

// 	case "*" :
// 		jiuguo = n1 * n2

// 	case "/" :
// 		jiuguo = n1 / n2
     
// 	}

//   return jiuguo


// }
  



func main(){

   var n1 float64 = 12
   var n2 float64 = 2 
   var fangfa string = "+"

   jieguo :=good.Haha(n1,n2,fangfa)

  fmt.Println(jieguo)

}