package main

import "strconv"

type node struct{
	name string
	isTrade bool
}

func (n node) String() string{
	return n.name + " " + strconv.FormatBool(n.isTrade)
}

func (n *node) setName(s string){
	n.name = s
}