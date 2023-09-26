package main

import "fmt"

func main(){
	var angka int32 = 2000
	number := 10

	kata := "Yang benar kamu? %t \n"
	benar := true
	var salah bool = false
	
	fmt.Println("Ini Uang Jajan mu nak", angka * int32(number))

	// %t untuk boolean
	fmt.Printf(kata ,benar)
	fmt.Printf(kata, salah)

	
}