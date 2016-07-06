package main

import "strings"

type region struct{
	nodes []*node
	index int
}

func (s *region) addnode (n node){
	s.nodes = append(s.nodes, &n)
}

func (r *region) findNode(name string) *node{
	for _,n := range r.nodes {
		if n.name == name{
			return n
		}
	}
	return nil
}

func (r *region) makeTradable(name string) int32{
	var foundNode = r.findNode(name)
	if(foundNode != nil){
		foundNode.isTrade = true
		return 1
	}
	return -1
}

func (r region) String () string{
	var returnString = ""
	for _,n := range r.nodes {
		returnString += n.String() + "\n"
	}
	return strings.TrimSpace(returnString)
}

func (r *region) connect(source string, destination string) int{
	var sourceNode = r.findNode(source)
	var destinationNode = r.findNode(destination)

	if sourceNode != nil && destinationNode != nil{
		sourceNode.addLink(destinationNode)
		return 1
	}
	return -1
}
