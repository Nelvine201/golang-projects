package main

import "fmt"
// checknumber returns true if the string contains a number.
func CheckNumber(s string) bool{
	for _,ch:= range s{
		if ch >= '0' && ch <= '9'{
			return true
		}
	}
	return false
}
func main(){
	result := CheckNumber("5 eggs")
	fmt.Println(result)
	result1 := CheckNumber("hello world")
	fmt.Println(result1)
	result2 := CheckNumber("golang-2026")
	fmt.Println(result2)



}





