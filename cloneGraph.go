package main

type CNode struct {
	Val       int
	Neighbors []*CNode
}

///DFS
// func cloneGraph(node *CNode) *CNode {

// 	visited := make(map[int]*CNode)
// 	return cloneGraphHelper(node, visited)
// }

// func cloneGraphHelper(node *CNode, visited map[int]*CNode) *CNode {

// 	if node == nil {
// 		return nil
// 	}

// 	if v, ok := visited[node.Val]; ok {
// 		return v
// 	}
// 	newNode := &CNode{Val: node.Val}
// 	visited[node.Val] = newNode

// 	for _, n := range node.Neighbors {
// 		newNode.Neighbors = append(newNode.Neighbors, cloneGraphHelper(n, visited))
// 	}

// 	return newNode
// }

// BFS Go/Golang
// func cloneGraph(node *Node) *Node {

// 	if node == nil {
// 		return node
// 	}

// 	// BFS, using queue data structure
// 	q := make([]*Node, 0)

// 	// HashMap to store visited nodes
// 	visited := make(map[int]*Node)

// 	q = append(q, node)
// 	newNode := &Node{Val: node.Val}
// 	visited[node.Val] = newNode

// 	for len(q) != 0 {
// 		temp := q[0] // read the first element from the queue
// 		q = q[1:]    // pop the element out of the queue

// 		for _, v := range temp.Neighbors {

// 			// if the node is not visited
// 			if _, ok := visited[v.Val]; !ok {
// 				visited[v.Val] = &Node{Val: v.Val}
// 				q = append(q, v)
// 			}
// 			visited[temp.Val].Neighbors = append(visited[temp.Val].Neighbors, visited[v.Val])
// 		}

// 	}
// 	return newNode
// }
