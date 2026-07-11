package main
import "fmt"

func CountCharacter(s string, r rune ) int{
	count :=0
	for _,ch:= range s{
		if ch == r {
			count ++
		}
	}
	return count
}
func main(){
	result :=  CountCharacter("banana",'a')
	fmt.Println(result)
	result1 :=  CountCharacter("hello world",'l')
	fmt.Println(result1)
	result2 :=  CountCharacter("Go Lang",'x')
	fmt.Println(result2)
}
