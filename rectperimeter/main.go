package main

import "fmt"

func RectPerimeter(h, w int) int{
	if h < 0 || w < 0 {
		return -1
	}
	return 2 * (h+w)
}
func main(){
	result := RectPerimeter(5, 7)
	fmt.Println(result)
	result1 := RectPerimeter(-5, 7)
	fmt.Println(result1)
}





