package main

type Node struct {
  Value int
  Parent *Node
  Left *Node
  Right *Node
  Balancing int
}

func (n *Node) adjustBalanceFactor(amount int) {
	n.Balancing += amount
}

func NewNode(value int) *Node {
  return &Node{
    Value: value,
    Parent: nil,
    Left: nil,
    Right: nil,
    Balancing: 0,
  }
}