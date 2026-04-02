package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Ben Arief", "Ben"))
	fmt.Println(strings.Split("Ben Arief", " "))
	fmt.Println(strings.ToLower("Ben Arief"))
	fmt.Println(strings.ToUpper("Ben Arief"))
	fmt.Println(strings.Trim("   Ben Arief   ", " "))
	fmt.Println(strings.ReplaceAll("ben Arief ben arief", "ben", "windah"))
}