package app

type node struct {
}

type methodTree struct {
	method string
	root   *node
}

type methodsTree []methodTree
