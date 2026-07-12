package main
import "fmt"

func Digitlen(n, base int) int{
	if base < 2 || base > 36{
		return -1

	}
	if n < 0{
		n = -n
	}
	if n == 0{
		return 1
	}
	count := 0
	for n > 0 {
		n = n/base
		count ++
	}
	return count

}
func main(){
result := Digitlen(15, 2)
fmt.Println(result)
result1 := Digitlen(16, 1)
fmt.Println(result1)
result2 := Digitlen(-15, 2)
fmt.Println(result2)
result3 := Digitlen(15, 37)
fmt.Println(result3)
result4 := Digitlen(15, -2)
fmt.Println(result4)
result5 := Digitlen(0, 2)
fmt.Println(result5)
}


