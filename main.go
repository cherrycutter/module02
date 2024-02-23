package main

import (
	"fmt"
	"strconv"
)

func main() {
	str := "100"
	num := 50

	strToNum, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("Error", err)
	}
	numToStr := strconv.Itoa(num)

	fmt.Printf("Строка '%s' теперь число %d\n", str, strToNum)
	fmt.Printf("Число %d теперь строка '%s'\n", num, numToStr)
	fmt.Printf("%T\n", strToNum)
	fmt.Printf("%T\n", numToStr)

}
