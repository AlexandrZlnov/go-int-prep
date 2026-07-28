package main

/*

func findEquivPair(root *TreeNode) (*TreeNode, *TreeNode) {
    m := map[uint]*TreeNode
    var result1, result2

	found := false

    var traverse func(node *TreeNode) uint

	traverse = func(node *TreeNode) uint {
        if node == nil {
            return 0
        }

        left := traverse(node.Left)
        right := traverse(node.Right)

		current := left | right | 1 << (node.Val - 'A')

        if prev, exists := m[current]; exists && !found {
            result1 = prev
            result2 = node
            found = true

			m[current] = node
		}



        return current
    }

    traverse(root)

    return result1, result2
}
*/
