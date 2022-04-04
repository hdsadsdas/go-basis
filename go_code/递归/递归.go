package main 
import "fmt"


func digui (i int){

    if i >2{
          i--
		digui(i)
		 
	}
   fmt.Println("i=",i)

}

func digui2(j int){

   if j >2{

        j--
        digui2(j)    

   }else{
      
	 fmt.Println("j=",j)  

   }


}


func main (){

      digui(4)    
      
	  digui2(4)

}