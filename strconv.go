package main

import (
	"fmt"
	"strconv"
)

func main() {
	result, e := strconv.ParseBool("true")
	if e != nil {
		fmt.Println("error", e.Error())
	} else {
		fmt.Println("result", result)
	}


	resultInt, e := strconv.Atoi("1000")
	if e != nil {
		fmt.Println("error", e.Error())
	} else {
		fmt.Println("result", resultInt)
	}

	binary := strconv.FormatInt(999, 2)
	fmt.Println("binary", binary)

	stringInt := strconv.Itoa(999)
	fmt.Println("stringInt", stringInt)

	hasil, err  := strconv.ParseFloat("100", 8)
	if err != nil {
		fmt.Println("error", err.Error())
	} else {
		fmt.Println("hasil", hasil)
	}
}