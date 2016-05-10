package main

import "fmt"

type region struct{
	nodes [] node
	index int
}

func (s *region) addnode (n node){
	s.nodes = append(s.nodes, n)
}

func (r region) Display (){
	for _,n := range r.nodes {
		fmt.Println(n)
	}
}