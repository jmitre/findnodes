package main

import "fmt"

//serendia may not be the best name, because all continents behave like this
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