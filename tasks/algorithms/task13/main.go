// Задача:
// Собес: Яндекс
// Дано бинарное дерево с выделенным корнем, в каждой
// вершине которого записана одна буква от A до Z.
// Две вершины считаются эквивалентными, если поддеревья
// этих вершин содержат одинаковое множество (те без учета частоты) букв.
// Базовая версия: нужно найти любую пару эквивалентных вершин.
// Пример:
//  		 A
// 	       /   \
// 	      C     B
// 	     / \   / \
//     	A   D A   D
//     /		   \
//    B	            C

package main

import "fmt"

type Node struct {
	Value rune
	Left  *Node
	Right *Node
}

func FindEquivalent(root *Node) (*Node, *Node) {
	var first, second *Node

	var traverse func(*Node) uint32

	set := make(map[uint32]*Node)

	traverse = func(node *Node) uint32 {
		if node == nil {
			return 0
		}

		leftMask := traverse(node.Left)
		rightMask := traverse(node.Right)

		selfMask := uint32(1) << (node.Value - 'A')

		mask := leftMask | rightMask | selfMask

		if other, ok := set[mask]; ok && first == nil {
			first = other
			second = node
		} else {
			set[mask] = node
		}
		return mask
	}

	traverse(root)

	return first, second
}

func main() {
	root := &Node{
		Value: 'A',
		Left: &Node{
			Value: 'C',
			Left: &Node{
				Value: 'A',
				Left: &Node{
					Value: 'B',
					Left:  nil,
					Right: nil,
				},
				Right: nil,
			},
			Right: &Node{
				Value: 'D',
				Left:  nil,
				Right: nil,
			},
		},
		Right: &Node{
			Value: 'B',
			Left: &Node{
				Value: 'A',
				Left:  nil,
				Right: nil,
			},
			Right: &Node{
				Value: 'D',
				Left:  nil,
				Right: &Node{
					Value: 'C',
					Left:  nil,
					Right: nil,
				},
			},
		},
	}

	a, b := FindEquivalent(root)

	fmt.Println(string(a.Value))
	fmt.Println(string(b.Value))

}
