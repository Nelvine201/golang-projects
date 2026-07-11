package main

import "fmt"

func RetainFirstHalf(s string) string{
	if len(s) == 1{
		return s
	}
	if len(s) == 0{
		return "0"
	}
	half := len(s) / 2
	
	return s[: half]
}
func main(){
	result := RetainFirstHalf("Hello Nelvine")
	fmt.Println(result)
	result1 := RetainFirstHalf("Retain first half of the string")
	fmt.Println(result1)
}



