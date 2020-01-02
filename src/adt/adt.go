package adt

import (
  "log"
)

type Tree interface {
  isTree()
}

type Node struct {
  l Tree
  r Tree
}

type Leaf struct {
  v int
}

func (_ Node) isTree() {}
func (_ Leaf) isTree() {}

func fold(t Tree, z func(v int) int, f func(l, r int) int) int {
  switch nt := t.(type) {
  case Leaf:
    return z(nt.v)
  case Node:
    return f(fold(nt.l, z, f), fold(nt.r, z, f))
  default:
    log.Fatalf("unexpected type %T", nt)
    return 0
  }
}
