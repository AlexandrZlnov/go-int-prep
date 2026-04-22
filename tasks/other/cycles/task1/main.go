// Задача вывести еслку в консоль

package main

import "fmt"

func main() {
	height := 11

	for i := 1; i <= height; i++ {
		for j := height - i; j >= 0; j-- {
			fmt.Print(" ")
		}
		for s := 1; s <= 2*i-1; s++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}

	for i := 0; i < 2; i++ {
		for j := height; j > 0; j-- {
			fmt.Printf(" ")
		}
		fmt.Println("|")
	}
}

// Вывод:
//   	      *
//           ***
//          *****
//         *******
//        *********
//       ***********
//      *************
//     ***************
//    *****************
//   *******************
//  *********************
//            |
//            |
