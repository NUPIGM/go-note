package notes

import "fmt"

func Lable() {
outside:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			fmt.Print("+")
			if i == 9 && j == 4 {
				break outside
			}
		}
		fmt.Println()
	}
}
func Goto() {
	fmt.Println(1)
	fmt.Println(2)
	if true {
		goto lable
	}
	fmt.Println(3)
	fmt.Println(4)
lable:
	fmt.Println(5)
	fmt.Println("<3,4被跳过了>")
}
