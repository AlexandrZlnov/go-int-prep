// Задача:
// На очередность выполнения defer
// Что выведет код и в каком порядке?

package main

import "fmt"

func PPrintln(msg string, i *int) {
	fmt.Printf("%s %d\n", msg, *i) // 3: pb=2
}

func tt() (a int) {
	a = 1
	defer func() { // 4: a=5+1 После return будут выполнены defer: 5:a=5 и 4:5+1=6. поскольку используется именованное возвращаемое значнеи defer переопределит "a" который будет равен 6
		a++
	}()
	defer func() { // 5: a = 5
		a = 5
	}()
	return 3
}

func main() {
	a := 1
	b := 2
	pb := &b

	defer fmt.Println("a=", a)   // 1: a=1
	defer fmt.Println("b=", *pb) // 2: b=2
	defer PPrintln("pb=", pb)    // 3: pb передается как указатель, его значение будет вычислено позже при выполнении функции в defer. К моменту выполнения defer *pb = 20

	a = 10
	*pb = 20

	fmt.Println(tt()) // 6
}

// Ответ:
// 6
// pb=20
// b=2
// a=1
