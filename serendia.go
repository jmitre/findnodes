package main

import "fmt"

//serendia may not be the best name, because all continents behave like this
type serendia struct{
	nodes [] node
	index int
}

func (s *serendia) addNode (n node){
	s.nodes = append(s.nodes, n)
}

func (s serendia) Display (){
	fmt.Println(len(s.nodes),s.nodes)
}