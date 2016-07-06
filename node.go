package main

import "strconv"
import "strings"

type node struct{
	name string
	isTrade bool
	cost int
	links []*node
}

func (n node) String() string{
	var linkString = ""
	for _,l := range n.links {
		linkString +=  l.name + "\n"
	}
	var returnString = n.name + " " + strconv.FormatBool(n.isTrade) + " cost: " + strconv.Itoa(n.cost)

	if(linkString != ""){
		returnString+= " links: " + linkString
	}

	return strings.TrimSpace(returnString)

}

func (n *node) setName(s string){
	n.name = s
}

func (n *node) addLink(link *node){
	n.links = append(n.links, link)
}
