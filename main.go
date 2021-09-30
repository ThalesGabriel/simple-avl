package main


func main() {
  // n := NewNode(10)
  // n2 := NewNode(12)
  // n.Right = n2

  tree := NewAVLTree(2)
  root := tree.Root
  root = tree.insert(5, root)
  root = tree.insert(1, root)
  root = tree.insert(7, root)
  root = tree.insert(6, root)
  tree.DisplayInOrder()
}