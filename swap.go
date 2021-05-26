package main
import(
	"fmt"
	)
	func main(){
		fmt.Println("\nEnter a number to swap:")
		var num1 int
		fmt.Scanln(&num1)
		
		fmt.Println("\nEnter the second number to swap:")
		var num2 int
		fmt.Scanln(&num2)
		
		fmt.Println("the numbers before swap are")
		fmt.Print(num1,num2)

		var temp int
		temp=num1
		num1=num2
		num2=temp
		
		fmt.Println("\nthe numbers after swap are")
		fmt.Print(num1,num2)
		
		
		
		}
