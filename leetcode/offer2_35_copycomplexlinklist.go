/*package main

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

var nodemap map[*Node]*Node

func copyRandomList(head *Node) *Node { //leetcode函数，对链表扫描，遇到不存在则创建
	nodemap = make(map[*Node]*Node, 0)
	//return deepcopy(head)
	if head == nil {
		return nil
	}
	node := head
	for node != nil {
		newnode := findnode(node)
		newnode.Next = findnode(node.Next)
		newnode.Random = findnode(node.Random)
		node = node.Next
	}
	return nodemap[head]
}

func deepcopy(node *Node) *Node { //leetcode答案
	if node == nil {
		return nil
	}
	if v, ok := nodemap[node]; ok {
		return v
	}
	newnode := &Node{Val: node.Val, Next: nil, Random: nil}
	nodemap[node] = newnode
	newnode.Next = deepcopy(node.Next)
	newnode.Random = deepcopy(node.Random)
	return newnode
}

func findnode(node *Node) *Node { //查找新结点是否存在，不存在则创建并加入哈希表，有就返回新结点地址
	if node == nil {
		return nil
	}
	if v, ok := nodemap[node]; ok {
		return v
	}
	newnode := &Node{Val: node.Val}
	nodemap[node] = newnode
	return newnode
}

func main() {
	node1 := Node{Val: 1}
	node2 := Node{Val: 2}
	node3 := Node{Val: 3}
	node1.Next = &node2
	node1.Random = &node3
	node2.Next = &node3
	node2.Random = &node1
	node3.Next = nil
	node3.Random = &node2

	newnode := copyRandomList(&node1)
	t := newnode
	for t != nil {
		println(t.Val)
		t = t.Next
	}
}*/
