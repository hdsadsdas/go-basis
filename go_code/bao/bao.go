package bao



func Haha(n1 float64, n2 float64,fangfa string ) (float64){
   
	var jiuguo float64
 
	 switch fangfa {
	 case "+" :
		 jiuguo = n1 + n2
	 
	 case "-" :
		 jiuguo = n1 - n2	
 
	 case "*" :
		 jiuguo = n1 * n2
 
	 case "/" :
		 jiuguo = n1 / n2
	  
	 }
 
   return jiuguo
 
 
 }