package main

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestNode(t *testing.T){
  var n = node{"Vellia", false, 3}

  assert.Equal(t, n.name, "Vellia", "The name should be initiallized to Vellia")
  assert.Equal(t, n.isTrade, false, "The is trade value should be initialized to false")
  assert.Equal(t, n.cost, 3, "The cost should be set to 3")
}

// func TestRegionInit(t *testing.T){
//   var r = new(region)
//   var emptyArr = []*node(nil)
//
//   assert.Equal(t, r.nodes, emptyArr, "A region should be initialized with an empty nodes array")
// }
//
// func TestRegionAddNode(t *testing.T){
//   var r = new(region)
//   r.addnode(node{"Vellia", false})
//
//   assert.Equal(t, r.nodes[0].name, "Vellia", "The addnode function correctly adds nodes")
//   assert.Equal(t, r.nodes[0].isTrade, false, "The addnode function correctly adds nodes")
// }
//
// func TestFindNode(t *testing.T){
//   var r = new(region)
//   r.addnode(node{"Olivia", false})
//   r.addnode(node{"Vellia", false})
//   r.addnode(node{"Heidel", true})
//   var foundNode = r.findNode("Vellia")
//
//   assert.Equal(t, foundNode.String(), "Vellia false", "Given the correct name, findNode returns the correct node")
// }
//
// func TestMakeTradable(t *testing.T){
//   var r = new(region)
//   r.addnode(node{"Olivia", false})
//   r.addnode(node{"Vellia", false})
//   r.addnode(node{"Heidel", false})
//   r.makeTradable("Vellia")
//   var foundNode = r.findNode("Vellia")
//
//   assert.Equal(t, foundNode.String(), "Vellia true", "makeTradable changes the isTrade attribute to true, on the node specified")
// }
//
// func TestOutputString(t *testing.T){
//   var r = new(region)
//   r.addnode(node{"Olivia", false})
//   r.addnode(node{"Vellia", false})
//   r.addnode(node{"Heidel", true})
//
//   var output = r.String()
//
//   assert.Equal(t, output, "Olivia false\nVellia false\nHeidel true", "after a region is filled with nodes, it should output a string representing those nodes")
// }
