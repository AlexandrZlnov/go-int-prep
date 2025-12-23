// Что выведет
package main

import (
	"fmt"
)

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	t := a[3:4:4]
	fmt.Println(t[0]) // 4 и если сделать append(t, 6) получим panic
}
