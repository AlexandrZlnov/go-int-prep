package main

/*
import "fmt"

type Node struct {
	Value rune
	Left  *Node
	Right *Node
}

func FindEquivalent(root *Node) (*Node, *Node) {
	seen := make(map[uint32]*Node)

	var first, second *Node

	var dfs func(*Node) uint32

	dfs = func(node *Node) uint32 {
		if node == nil {
			return 0
		}

		leftMask := dfs(node.Left)
		rightMask := dfs(node.Right)

		// бит текущей буквы
		selfMask := uint32(1) << (node.Value - 'A')

		// множество букв текущего поддерева
		mask := leftMask | rightMask | selfMask

		if other, ok := seen[mask]; ok && first == nil {
			first = other
			second = node
		} else {
			seen[mask] = node
		}

		return mask
	}

	dfs(root)

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
				},
			},
			Right: &Node{
				Value: 'D',
			},
		},
		Right: &Node{
			Value: 'B',
			Left: &Node{
				Value: 'D',
			},
			Right: &Node{
				Value: 'A',
				Right: &Node{
					Value: 'C',
				},
			},
		},
	}

	a, b := FindEquivalent(root)

	if a != nil {
		fmt.Printf("%c %c\n", a.Value, b.Value)
	}
}


*/
