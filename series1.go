package main

import "fmt"

func main() {

	var a = 1
	var b = 1
	no_of_digit := 0
	fmt.Println("Enter a digit")
	fmt.Scanln(&no_of_digit)

	fmt.Print(a, ",", b)

	for i := 3; i <= no_of_digit; i++ {
		next_num := a + b

		if i == no_of_digit {
			fmt.Println(",", next_num)
		} else {
			fmt.Print(",", next_num)
		}

		a = b
		b = next_num
	}

}
