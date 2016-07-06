package main

import "strconv"

type node struct{
	name string
	isTrade bool
	cost int
}

func (n node) String() string{
	return n.name + " " + strconv.FormatBool(n.isTrade) + " cost: " + strconv.Itoa(n.cost)
}

func (n *node) setName(s string){
	n.name = s
}
