package main 

import  "fmt"

 func sum(str string,num ...int) (string, int) {

    sum :=0
	for _, v := range num {
        sum += v
    }
	return str, sum
 }

func main() {

	str1,result := sum("praveen",1,2,3,4,5)
	fmt.Println("Sum:", result)

	_, result1 :=  sum("hello",1,2,3,4)
	fmt.Println("Sum:", result1)

	str1, result = sum("Bangalore", 1,2,3,4,7,8,10,15,56)
	fmt.Println("Sum:", result)
	fmt.Println("Place:", str1)
		
}