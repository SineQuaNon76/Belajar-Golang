package main

import (
	"GOLANGDASAR/arrayP/calcu"
	"fmt"
)

func main() {
	arr := []int32 {
		12, 
		2500, 
		90, 
		78, 
		56,
	}


	arr = append(arr, 78)

	perulangan(1,2,3,44,5,5,6,78,8)


	arr1 := []interface{}{ // ini interface any
		"Joko",
		15,
		"Surti",
		"Narto",
		167,
	}

	//Slice
	newArr1 := arr1[1:4]
	fmt.Println(newArr1)


	nilai :=  8.5

	if nilai == 10 {
		fmt.Println("Nilai Sempurna")
	} else if nilai >= 9 || nilai >= 8{
		fmt.Println("GG")
	}else if nilai >= 7 {
		fmt.Println("lulus")
	}else {
		fmt.Println("kaga lulus")
	}

	for i := 1; i < 5; i++ {
		for j := i; j < 5; j++ { 
			fmt.Print(j, " ")
		}
	
		fmt.Println()
	}

	// result := calculation.Multiple(1, 12, 3)
	// println("Hasil perkalian = ", result)


	// honda := map[string] string {
	// 	"Beat" : "125 CC",
	// 	"Vario" : "160 CC",
	// 	"PCX" : "160 CC",
	// }

	// for _, value := range honda {
	// 	fmt.Printf("Motor anda: %s \n", value)
	// }



	// Example slice
	myArray := []string{
		"BRI",
		"BNI",
		"BCA",
		"Mandiri",
		"UOB",
	}

	myArray = append(myArray, "TMRW")


	// for i := 0; i < len(myArray); i++ {
	// 	fmt.Printf("Bank : %s\n", myArray[i])
	// }

	for _, bank := range myArray{
		fmt.Printf("Bank : %s\n", bank)
	}
	// end example slice 

	lorem := "Lorem ipsum dolor sit amet consectetur"

	for index, value := range lorem{
		if index % 2 == 1 {
			fmt.Printf("Index : %d, Value %c: \n", index, value)
		}
	}

	inputString := "Golang the best programing language"
	hasil := calculation.MyArray(inputString)
	println(hasil)

	
	ai := calculation.MyAiueo(inputString)
	println(ai)


	// Map
	myMap := map[int]string{
	1 : "Satu",
	2 : "Dua",
	3 : "Tiga",
	}

	for kunci, hasil := range myMap{
		fmt.Printf("Key ke - %d, Value : %s\n", kunci, hasil)
		// fmt.Println(kunci, hasil )
	} 

	// End Map





	// for i := 0; i < len(arr); i++ {
	// 	fmt.Println("Angka ke-", arr[i])
	// }
}

func perulangan(angka ...int){
	// for i := 0; i < len(angka); i++ {
	// 	fmt.Println("Angka Ke-", angka[i])
	// }

	for number := range angka{
		fmt.Println("Angke Ke-", number)
	}
}

