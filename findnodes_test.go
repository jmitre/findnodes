package main

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestNode(t *testing.T){
  var n = node{"Vellia", false}

  assert.Equal(t, n.name, "Vellia", "The name should be initiallized to Vellia")
  assert.Equal(t, n.isTrade, false, "The is trade value should be initialized to false")
}

func TestRegionInit(t *testing.T){
  var r = new(region)
  var emptyArr = []node(nil)

  assert.Equal(t, r.nodes, emptyArr, "A region should be initialized with an empty nodes array")
}

func TestRegionAddNode(t *testing.T){
  var r = new(region)
  r.addnode(node{"Vellia", false})

  assert.Equal(t, r.nodes[0].name, "Vellia", "The addnode function correctly adds nodes")
  assert.Equal(t, r.nodes[0].isTrade, false, "The addnode function correctly adds nodes")
}

func TestOutputString(t *testing.T){
  var r = new(region)
  r.addnode(node{"Olivia", false})
  r.addnode(node{"Vellia", false})
  r.addnode(node{"Heidel", true})

  var output = r.String()

  assert.Equal(t, output, "Olivia false\nVellia false\nHeidel true", "After a region is filled with nodes, it should output a string representing those nodes")
}
