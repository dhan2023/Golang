package main

// import (
// 	"fmt"

// 	"github.com/dhan2023/mascot"
// )

// func main() {
// 	fmt.Println(mascot.BestMascot())
// }

import (
	"fmt"
)

func main() {
	a := [...]int{1, 2, 3}
	b := &a
	b[1] = 5

	fmt.Println(a)
	fmt.Println(b)

}
