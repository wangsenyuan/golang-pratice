package p1847

import "sort"

func closestRoom(rooms [][]int, queries [][]int) []int {
	rs := make([]Room, len(rooms))
	for i, r := range rooms {
		rs[i] = Room{r[0], r[1]}
	}
	sort.Slice(rs, func(i, j int) bool {
		return rs[i].size > rs[j].size
	})

	qs := make([]Query, len(queries))
	for i, q := range queries {
		qs[i] = Query{q[0], q[1], i}
	}

	sort.Slice(qs, func(i, j int) bool {
		return qs[i].size > qs[j].size
	})

	var root *Node
	ans := make([]int, len(queries))
	var j int
	for _, q := range qs {
		for j < len(rs) && rs[j].size >= q.size {
			root = Insert(root, rs[j].id)
			j++
		}

		a := LowerBound(root, q.pref)
		ans[q.id] = -1
		if a != nil {
			ans[q.id] = a.key
		}
		b := Before(root, q.pref)
		if b != nil {
			if ans[q.id] < 0 || abs(q.pref-b.key) <= abs(q.pref-ans[q.id]) {
				ans[q.id] = b.key
			}
		}
	}

	return ans
}

type Room struct {
	id   int
	size int
}

type Query struct {
	pref int
	size int
	id   int
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

/**
* this is a AVL tree
 */
type Node struct {
	key         int
	height      int
	cnt         int
	left, right *Node
}

func (node *Node) Height() int {
	if node == nil {
		return 0
	}
	return node.height
}

func NewNode(key int) *Node {
	node := new(Node)
	node.key = key
	node.height = 1
	node.cnt = 1
	return node
}

func rightRotate(y *Node) *Node {
	x := y.left

	t2 := x.right

	x.right = y
	y.left = t2
	y.height = max(y.left.Height(), y.right.Height()) + 1
	x.height = max(x.left.Height(), x.right.Height()) + 1

	return x
}

func leftRotate(x *Node) *Node {
	y := x.right
	t2 := y.left

	y.left = x
	x.right = t2

	x.height = max(x.left.Height(), x.right.Height()) + 1
	y.height = max(y.left.Height(), y.right.Height()) + 1

	return y
}

func (node *Node) GetBalance() int {
	if node == nil {
		return 0
	}
	return node.left.Height() - node.right.Height()
}

func Insert(node *Node, key int) *Node {
	if node == nil {
		return NewNode(key)
	}
	if node.key == key {
		node.cnt++
		return node
	}

	if node.key > key {
		node.left = Insert(node.left, key)
	} else {
		node.right = Insert(node.right, key)
	}

	node.height = max(node.left.Height(), node.right.Height()) + 1
	balance := node.GetBalance()

	if balance > 1 && key < node.left.key {
		return rightRotate(node)
	}

	if balance < -1 && key > node.right.key {
		return leftRotate(node)
	}

	if balance > 1 && key > node.left.key {
		node.left = leftRotate(node.left)
		return rightRotate(node)
	}

	if balance < -1 && key < node.right.key {
		node.right = rightRotate(node.right)
		return leftRotate(node)
	}

	return node
}

func LowerBound(root *Node, key int) *Node {
	if root == nil {
		return nil
	}
	if root.key >= key {
		res := LowerBound(root.left, key)
		if res != nil {
			return res
		}
		return root
	}
	return LowerBound(root.right, key)
}

func Before(root *Node, key int) *Node {
	if root == nil {
		return nil
	}
	if root.key >= key {
		return Before(root.left, key)
	}
	// root.key < key
	res := Before(root.right, key)
	if res != nil {
		return res
	}
	return root
}

func MinValueNode(root *Node) *Node {
	cur := root

	for cur.left != nil {
		cur = cur.left
	}

	return cur
}

func Delete(root *Node, key int) *Node {
	if root == nil {
		return nil
	}

	if key < root.key {
		root.left = Delete(root.left, key)
	} else if key > root.key {
		root.right = Delete(root.right, key)
	} else {
		root.cnt--
		if root.cnt > 0 {
			return root
		}
		if root.left == nil || root.right == nil {
			tmp := root.left
			if root.left == nil {
				tmp = root.right
			}
			root = tmp
		} else {
			tmp := MinValueNode(root.right)

			root.key = tmp.key
			root.cnt = tmp.cnt
			// make sure tmp node deleted after call delete on root.right
			tmp.cnt = 1
			root.right = Delete(root.right, tmp.key)
		}
	}

	if root == nil {
		return root
	}

	root.height = max(root.left.Height(), root.right.Height()) + 1
	balance := root.GetBalance()

	if balance > 1 && root.left.GetBalance() >= 0 {
		return rightRotate(root)
	}

	if balance > 1 && root.left.GetBalance() < 0 {
		root.left = leftRotate(root.left)
		return rightRotate(root)
	}

	if balance < -1 && root.right.GetBalance() <= 0 {
		return leftRotate(root)
	}

	if balance < -1 && root.right.GetBalance() > 0 {
		root.right = rightRotate(root.right)
		return leftRotate(root)
	}

	return root
}
