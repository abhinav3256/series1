package main

import "fmt"

func main() {

	var a = 1
	var b = 1
	no_of_digit := 0
	fmt.Println("Enter a digit")
	fmt.Scanln(&no_of_digit)
	if no_of_digit == 1 {
		fmt.Println(a)
	} else {
		fmt.Println(a)

		fmt.Println(b)
	}

	for i := 3; i <= no_of_digit; i++ {
		next_num := a + b

		fmt.Println(next_num)

		a = b
		b = next_num
	}

}
