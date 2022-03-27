package main
import (
	"fmt"
)


func main(){


     var haha string


	 hello:

     fmt.Println("爱不爱我")
	 fmt.Scanln(&haha)
	 
	 switch haha {
		
	 case "爱" :
		{
			fmt.Println("我也爱你")
		}

	case "不爱" :
		{
			fmt.Println("哈哈哈，我也是")
		}

    default :
	    {
            fmt.Println("说重点")

            goto hello
			
		}

	 }




}
	
	 

