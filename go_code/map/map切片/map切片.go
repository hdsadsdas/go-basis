package main
import "fmt"

func main(){

   xixi := map[string]int{
      "年龄" : 20 ,
      "工资" : 10000,
   }

   lala := map[string]int{
      "年龄" : 21 ,
      "工资" : 15000,
   }

   haha := []map[string]int{
       
     xixi,
     lala,
	  
   } 

  fmt.Println(haha)

  //map切片添加

      new := map[string]int{
      "年龄" : 22 ,
      "工资" : 20000,
      }
 
   haha = append(haha,new)

   fmt.Println(haha)


}
