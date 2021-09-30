package main

import "fmt"

type AVLTree struct {
	Root *Node
}

func NewAVLTree(value int) *AVLTree {
	return &AVLTree{
		Root: NewNode(value),
	}
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func Height(node *Node) int {
	if node == nil {
		return 0
	}

	h1 := Height(node.Left)
	h2 := Height(node.Right)
	return (1 + max(h1, h2))
}

func rotateRight(x *Node) *Node {
  y := x.Left
  t := y.Right

  y.Right = x
  x.Left = t

  x.Balancing = Height(x)
	y.Balancing = Height(y)
  fmt.Println(y)

  return y
}

func rotateLeft(x *Node) *Node {
  y := x.Right
  t := y.Left

  y.Left = x
  x.Right = t

  x.Balancing = Height(x)
	y.Balancing = Height(y)

  fmt.Println(y)

  return y
}


func (t *AVLTree) rebalanceTree(n *Node, insertedData int) {
	if n == nil {
		return 
	}

	// check balance factor and rotateLeft if right-heavy and rotateRight if left-heavy
	balance := Height(n.Left) - Height(n.Right)

  if balance > 1 && insertedData < n.Left.Value {
    rotateRight(n)
  }

  // linearly to the right
  if balance < -1 && insertedData > n.Right.Value {
    rotateLeft(n)
  }

  // less than symbol
  if balance > 1 && insertedData > n.Left.Value {
    n.Left = rotateLeft(n.Left)
    rotateRight(n)
  }

  // greater than symbol
  if balance < -1 && insertedData < n.Right.Value {
    n.Right = rotateRight(n.Right)
    rotateLeft(n)
  }

}

func (t *AVLTree) insert(newData int, n *Node) *Node {
	if n == nil {
		return NewNode(newData)
	}

	if newData < n.Value {
		n.Left = t.insert(newData, n.Left)
		n.adjustBalanceFactor(-1)
	} else if newData > n.Value {
		n.Right = t.insert(newData, n.Right)
		n.adjustBalanceFactor(1)
	}

	t.rebalanceTree(n, newData)

	return n
}

// DISPLAY

func (t *AVLTree) DisplayInOrder() {
	t.Root.displayNodesInOrder()
}

func (n *Node) displayNodesInOrder() {
	fmt.Print(n.Value, " balance => ", n.Balancing, " / ")
	if n.Left != nil {
		n.Left.displayNodesInOrder()
	}
	if n.Right != nil {
		n.Right.displayNodesInOrder()
	}
}
