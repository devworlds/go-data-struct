package main

func main() {
	tree := Insert(nil, 12)

	Insert(tree, 15)
	Insert(tree, 11)
	Insert(tree, 1)
	Insert(tree, 9)

	println("Original: ")
	PostOrder(tree)
	InvertTree(tree)
	println("Invertida: ")
	PostOrder(tree)
}
