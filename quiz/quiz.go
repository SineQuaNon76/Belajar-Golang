package main

import (
	"errors"
	"fmt"
)

func main(){
	scores := []int{
		100,
		80,
		75,
		92,
		70,
		93,
		88,
		67,
	}


	avg := avarage(scores)
	fmt.Printf("Hasil %.3f \n", float64(avg))

	lebih := lebih(scores)
	fmt.Printf("Lebih dari 90 %d \n", lebih)

	for _, value := range scores{
		if value >= 90 {
			fmt.Printf("lebih dariiiiiii 90 %d\n", value)
		}
	}

	// quiz function 1
	nilai := []int{
		10,
		5,
		8,
		9,
		7,
	}
	sum := sum(nilai)
	fmt.Printf("Hasilnya adalah : %d\n", sum)


	//quiz function 2
	calculate, err := calculate(12, 12, "*")
	if err != nil {
		fmt.Printf("Terjadi Error %v", err.Error())
	}else{
		fmt.Printf("Hasilnya adalah : %d \n", calculate)
	}
	
	

}



func avarage(scores[]int) float64{
	result := 0

	for _, value := range scores {
		result += value
	}
	avarage := float64(result) / float64(len(scores))
	return avarage
}



func lebih(scores []int) []int{
	max := []int {}

	for _, value := range scores {
		if value >= 90 {
			max = append(max, value)
		}
	}
	return max
}

func sum(nilai []int) int{
	result := 0

	for _, value := range nilai{
		result += value
	}

	return result
}

func calculate(number, numberOne int, operator string) (int, error) {
	var hasil int 

	if operator == "+" {
		hasil = number + numberOne
	}else if operator == "-" {
		hasil = number - numberOne
	}else if operator == "*" {
		hasil = number * numberOne
	}else if operator == "/" {
		hasil = number / numberOne
	}else{
		return 0, errors.New("Operator Tidak Valid") 
	}

	return hasil, nil
}