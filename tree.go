package tinygo

type node struct {
}

type methodTree struct {
	method string
	root   *node
}

type methodsTree []methodTree

func (n *node) getValue(path string) {

}
