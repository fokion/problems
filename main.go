package main

import "fmt"

func main() {
	firstName := "Fokion"

	var lastName *string

	fmt.Println(firstName)

	if lastName == nil {
		fmt.Println("Last name pointer is nil")
		lastName = new(string)
		*lastName = "Sotiropoulos"
		fmt.Println("The memory location of the pointer is", lastName)
		fmt.Println(*lastName)
	}

	pointer := &firstName
	fmt.Println(pointer, *pointer)

	firstName = "Testing"

	fmt.Println(pointer, *pointer)

	var age int = 32
	fmt.Println("my age is ", age)

	//doing complex numbers
	complexNum := complex(3, 4)
	realNum, imaginaryNum := real(complexNum), imag(complexNum)
	fmt.Println(realNum, imaginaryNum)

}
