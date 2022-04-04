package main
import (

	   "fmt"
	   "io"
	   "bufio"
	   "os"

)

func filecopy(n1 string,n2 string){

	  n1file , ok  := os.Open(n1)
	  if ok != nil {
		  fmt.Println("打开文件失败")
	  }
  defer n1file.Close()
	    readn1 := bufio.NewReader(n1file)

		n2file, ok2 := os.OpenFile(n2,os.O_CREATE | os.O_WRONLY ,4423)
      if ok2 != nil {
		  fmt.Println("打开文件2错误")
	  }
   
	  defer n2file.Close()
	  writern2 := bufio.NewWriter(n2file)

	  io.Copy(writern2,readn1)

}

func  main()  {

	n1 := "d:/wenjian.txt"
	n2 := "d:/wenjian3.txt"

	filecopy(n1,n2)
	
	
}