package main

import "strings"

type region struct{
	nodes []node
	index int
}

func (s *region) addnode (n node){
	s.nodes = append(s.nodes, n)
}

func (r region) String () string{
	var returnString = ""
	for _,n := range r.nodes {
		returnString += n.String() + "\n"
	}
	return strings.TrimSpace(returnString)
}
