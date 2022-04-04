package haha

func App(n1 int) int {
	
	text := 0 

	for i := 1 ; i <= n1; i++{

         text += i

	}
	return text

}

func GetSub(n1 , n2 int) int {

	return n1 - n2
}