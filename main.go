package main

import (
	"crypto/sha256"
	"fmt"
	"os"
)

func main() {
	arr := os.Args[1:]

	if len(arr) == 0 {
		fmt.Println("No arguments passed, we'll use the default list")
		arr = []string{"1", "2", "3", "4", "5"}
	}

	fmt.Printf("root hash is %s\n", getRootHash(arr))
}

type node struct {
	left  *node
	right *node
	hash  []byte
}

func getRootHash(arr []string) string {
	var nodes []node

	for _, str := range arr {
		n := newNode(nil, nil, []byte(str))
		nodes = append(nodes, *n)
	}

	if len(nodes)%2 != 0 {
		addEmptyNode(&nodes)
	}

	for len(nodes) > 1 {
		var newLevel []node

		for j := 0; j < len(nodes)-1; j += 2 {
			node := newNode(&nodes[j], &nodes[j+1], nil)
			newLevel = append(newLevel, *node)
		}

		nodes = newLevel
		if len(nodes) > 1 && len(nodes)%2 != 0 {
			addEmptyNode(&nodes)
		}
	}

	return fmt.Sprintf("%x", nodes[0].hash)
}

func addEmptyNode(nodes *[]node) {
	*nodes = append(*nodes, *newNode(nil, nil, []byte{}))
}

func newNode(left, right *node, data []byte) *node {
	n := node{}

	if left == nil && right == nil {
		hash := sha256.Sum256(data)
		n.hash = hash[:]
	} else {
		prevHashes := append(left.hash, right.hash...)
		hash := sha256.Sum256(prevHashes)
		n.hash = hash[:]
	}

	n.left = left
	n.right = right

	return &n
}
