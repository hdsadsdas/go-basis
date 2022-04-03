package main
import (
	"fmt"
	"math/rand"
	"time"
)
func main(){


var haha [5]int
//为了每次生成的随机数不一样，我们需要给一个seed值
rand.Seed(time.Now().Unix())
for i := 0 ; i < len(haha) ; i++{

    haha[i] = rand.Intn(100)// 随机生成0到100的数

}
fmt.Println("交换前",haha)

//反转打印
var xixi int
for i := 0 ; i < len(haha)/2 ; i++{
	xixi = haha[len(haha)-1-i]
	haha[len(haha)-1-i]=haha[i]
    haha[i]=xixi 

}
fmt.Println("交换后",haha)

}