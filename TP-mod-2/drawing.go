package main
import "fmt"

func main(){
	var inp, i, j int
	var stars string 

	fmt.Scan(&inp)
	if inp > 2 {
		for i = 0; i < inp; i++ {
		stars = ""
		for j = 0; j < inp; j++ {
			if j == i {
				stars += " "
			} else {
				stars += "*"
			}
		}
		fmt.Println(stars)
	}
	}
}