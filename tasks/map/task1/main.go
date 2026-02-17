// Собес: Яндекс
// Задача:
/*
Дан список ненулевой длины, состоящий из направлений. Направление обозначается одним из четырёх символов:
L – Left
R – Right
U – Up
D – Down

Каждый элемент перемещает нас на 1 в заданном направлении.
Известно, что петли (возврат в уже посещённую точку) дают нулевое перемещение и являются пустой тратой времени. Нужно удалить из списка петли:
[R, D, L, U, R] -> [R]
[R, D, L, R, U, U, R] -> [R, U, R]
Важно отметить, что цель не просто попасть в ту же самую конечную точку, но и придерживаться первоначального маршрута (не срезать по прямой):
[D, R, U] -> [D, R, U]

Вернуть нужно массив направлений (а не массив посещённых точек).
Ограничения: O(N) по памяти, O(N) по времени.


func optimize(input []string) []string {

}

*/

package main

import (
	"fmt"
	"time"
)

func optimize(input []string) []string {
	x, y := 0, 0

	currentWay := make(map[[2]int]int, len(input))
	resultWay := make([]string, 0)

	currentWay[[2]int{0, 0}] = 0
	time.Sleep(100 * time.Millisecond)

	for i, step := range input {
		switch step {
		case "R":
			x++
		case "L":
			x--
		case "U":
			y++
		case "D":
			y--
		default:
			continue
		}

		if crossInd, ok := currentWay[[2]int{x, y}]; ok {
			resultWay = resultWay[:crossInd]

		} else {
			currentWay[[2]int{x, y}] = i + 1

			resultWay = append(resultWay, step)

		}

	}
	return resultWay

}

func main() {
	way := []string{"R", "R", "R", "D", "L", "U", "R", "U", "L"}
	fmt.Println(optimize(way))
}
