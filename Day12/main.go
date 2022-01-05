package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"unicode"
)

type node struct {
	name    string
	paths   []*node
	isBig   bool
	isStart bool
	isEnd   bool
}

const START = "start"
const END = "end"

func main() {
	part1 := getPaths("input.txt")
	//part2 := part2("input.txt")
	fmt.Printf("Paths:  %v \n", part1)
}

func getPaths(s string) int {
	startNode := readFileIntoNodes(s)
	fmt.Print(startNode.name + ", ")
	output := traverse(startNode)

	return output
}

func traverse(node *node) int {
	if node.isEnd {
		fmt.Print(node.name + "\n")
		return 1
	} else {
		for _, child := range node.paths {
			if child.name == node.name {
				if node.isBig {
					fmt.Print(child.name + ", ")
					return traverse(child) + 1
				}
			} else {
				if child.name != START {
					fmt.Print(child.name + ", ")
					return traverse(child) + 1
				}
			}
		}
	}
	return 0
}

func readFileIntoNodes(name string) *node {
	var startNode *node
	nodeList := []*node{}
	content, err := ioutil.ReadFile(name)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		nodes := strings.Split(line, "-")
		lNodeName := nodes[0]
		rNodeName := nodes[1]

		var lNode *node
		var rNode *node

		for _, createdNode := range nodeList {
			if createdNode.name == lNodeName {
				lNode = createdNode
			}
			if createdNode.name == rNodeName {
				rNode = createdNode
			}
		}

		if lNode == nil {
			lNode = makeNode(lNodeName)
			if lNode.isStart {
				startNode = lNode
			}
			nodeList = append(nodeList, lNode)
		}

		if rNode == nil {
			rNode = makeNode(rNodeName)
			if rNode.isStart {
				startNode = rNode
			}
			nodeList = append(nodeList, rNode)
		}

		lNode.paths = append(lNode.paths, rNode)
		rNode.paths = append(rNode.paths, lNode)

	}
	return startNode
}

func makeNode(name string) *node {
	if name == START {
		return &node{name: name, isBig: false, isStart: true, isEnd: false, paths: []*node{}}
	}
	if name == END {
		return &node{name: name, isBig: false, isStart: false, isEnd: true, paths: []*node{}}
	}

	isBig := unicode.IsUpper(rune(name[1]))

	return &node{name: name, isBig: isBig, isStart: false, isEnd: false}

}
