package calculation

import (
	"fmt"
	"strings"
	
)

func Multiple(number int, number1 int, number2 int) int {
	return number * (number1 + number2)
}



// function yang berfungsi untuk mencetak character genap saja dari sebuah input kalimat
func MyArray(myString string) string{
	hasil := ""
	for index, value := range myString{
		if index % 2 == 0 {
			hasil += fmt.Sprintf("Index ke -: %d, Value: %c\n", index, value)
		}
	}
	return hasil
}


//function menctak hanya a,i,u,e,o
func MyAiueo(myString string) string {
	result := ""
	for i := 0; i < len(myString); i++{
		if strings.Contains("aiueo", string(myString[i])){
			result += fmt.Sprintf("Only a,i,u,e,o : %c\n", myString[i])
		}
	}
	return result
}