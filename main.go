package main

import "fmt"

func main() {
	firstName := "Fokion"
	fmt.Println(firstName)

	var age int = 32
	fmt.Println("my age is ", age)

	//doing complex numbers
	complexNum := complex(3, 4)
	realNum, imaginaryNum := real(complexNum), imag(complexNum)
	fmt.Println(realNum, imaginaryNum)

}
