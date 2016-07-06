package main

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestNode(t *testing.T){
  var n = node{"Vellia", false, 3, []*node{}}

  assert.Equal(t, n.name, "Vellia", "The name should be initiallized to Vellia")
  assert.Equal(t, n.isTrade, false, "The is trade value should be initialized to false")
  assert.Equal(t, n.cost, 3, "The cost should be set to 3")
}

func TestRegionInit(t *testing.T){
  var r = new(region)
  var emptyArr = []*node(nil)

  assert.Equal(t, r.nodes, emptyArr, "A region should be initialized with an empty nodes array")
}

func TestRegionAddNode(t *testing.T){
  var r = new(region)
  r.addnode(node{"Vellia", false, 3, []*node{}})

  assert.Equal(t, r.nodes[0].name, "Vellia", "The addnode function correctly adds nodes")
  assert.Equal(t, r.nodes[0].isTrade, false, "The addnode function correctly adds nodes")
  assert.Equal(t, r.nodes[0].cost, 3, "The addnode function correctly adds nodes")
}

func TestFindNode(t *testing.T){
  var r = new(region)
  r.addnode(node{"Olivia", false, 2, []*node{}})
  r.addnode(node{"Vellia", false, 3, []*node{}})
  r.addnode(node{"Heidel", true, 2, []*node{}})
  var foundNode = r.findNode("Vellia")

  assert.Equal(t, foundNode.String(), "Vellia false cost: 3", "Given the correct name, findNode returns the correct node")
}

func TestMakeTradable(t *testing.T){
  var r = new(region)
  r.addnode(node{"Olivia", false, 3, []*node{}})
  r.addnode(node{"Vellia", false, 3, []*node{}})
  r.addnode(node{"Heidel", false, 2, []*node{}})
  r.makeTradable("Vellia")
  var foundNode = r.findNode("Vellia")

  assert.Equal(t, foundNode.String(), "Vellia true cost: 3", "makeTradable changes the isTrade attribute to true, on the node specified")
}

func TestOutputString(t *testing.T){
  var r = new(region)
  r.addnode(node{"Olivia", false, 3, []*node{}})
  r.addnode(node{"Vellia", false, 2, []*node{}})
  r.addnode(node{"Heidel", true, 3, []*node{}})

  var output = r.String()

  assert.Equal(t, output, "Olivia false cost: 3\nVellia false cost: 2\nHeidel true cost: 3", "after a region is filled with nodes, it should output a string representing those nodes")
}

func TestNodeLinks(t *testing.T){
  var n = new(node)
  var emptyArr = []*node(nil)

  assert.Equal(t, n.links, emptyArr, "A node should have a list of node links")
}

func TestNodeLinksAdd(t *testing.T){
  n := node{"Olivia", false, 3, []*node{}}
  linkNode := node{"Vellia", false, 3, []*node{}}
  n.addLink(&linkNode)

  assert.Equal(t, n.links[0].String(), linkNode.String(), "A node should have a list of node links")
}

func TestRegionConnect(t *testing.T){
  var r = new(region)
  r.addnode(node{"Valencia", false, 3, []*node{}})
  r.addnode(node{"Serendia", false, 2, []*node{}})
  r.addnode(node{"Heidel", true, 3, []*node{}})

  r.connect("Serendia", "Heidel")
  r.connect("Serendia", "Valencia")
  r.connect("Heidel", "Valencia")
  r.connect("Valencia", "Heidel")
  r.connect("Valencia", "Serendia")

  assert.Equal(t, r.nodes[0].links[0].String(), "Heidel true cost: 3 links: Valencia")
}

func TestConnectErrors(t *testing.T){
  var r = new(region)
  r.addnode(node{"Valencia", false, 3, []*node{}})
  r.addnode(node{"Serendia", false, 2, []*node{}})
  r.addnode(node{"Heidel", true, 3, []*node{}})

  var err = r.connect("Vellia", "Heidel")
  assert.Equal(t, err, -1, "Shoudl return -1 for nodes that don't exist")

  err = r.connect("Valencia", "Serendia")
  assert.Equal(t, err, 1, "Should return 1 for nodes that do exist")
}
