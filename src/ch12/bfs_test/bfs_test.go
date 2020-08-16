package bfs_test

import (
	"fmt"
	"reflect"
	"testing"
)

type Node struct {
Value  int		// 节点值
left  *Node 	// 节点的左子叶引用
right *Node 	// 节点的右子叶引用
}

type Tree struct {
root *Node 		// 根节点引用
}

func NewTree() *Tree {
return &Tree{}
}

func (t *Tree) Find(target int) *Node {
curNode := t.root // 从根节点开始查找
if curNode == nil {
return nil
}
for curNode.Value != target {
if curNode.Value < target {
curNode = curNode.right
}else {
curNode = curNode.left
}

if curNode == nil {
return nil
}
}
return curNode
}

func (t *Tree) Insert(target int) {
node := &Node{Value: target}

if t.root == nil {
t.root = node
return
}

curNode := t.root

for {
if node.Value < curNode.Value {
if curNode.left == nil {
curNode.left = node
return
}
curNode = curNode.left
} else {
if curNode.right == nil {
curNode.right = node
return
}
curNode = curNode.right
}
}
}

func (t *Tree) Delete(target int) {
// 太复杂，暂不研究
}

const (
BEGIN_SORT = iota
MIDDLE_SORT
END_SORT
)

func (t *Tree) Traverse(sortType int) {
switch sortType {
case BEGIN_SORT:
fmt.Print("前序遍历:")
t.BeginSort(t.root)
case MIDDLE_SORT:
fmt.Print("中序遍历:")
t.MiddleSort(t.root)
case END_SORT:
fmt.Print("后序遍历:")
t.EndSort(t.root)
}
}

func (t *Tree) BeginSort(node *Node) {
if node != nil {
fmt.Print(node.Value, " ")
t.BeginSort(node.left)
t.BeginSort(node.right)
}
}
func (t *Tree) MiddleSort(node *Node) {
if node != nil {
t.MiddleSort(node.left)
fmt.Print(node.Value, " ")
t.MiddleSort(node.right)
}
}
func (t *Tree) EndSort(node *Node) {
if node != nil {
t.EndSort(node.left)
t.EndSort(node.right)
fmt.Print(node.Value, " ")
}
}


func TestTree(t *testing.T) {
tree := NewTree()

tree.Insert(1)
tree.Insert(3)
tree.Insert(2)
tree.Insert(5)
tree.Insert(9)
//data,_ := json.Marshal(tree)
//v := string(data)
fmt.Println(reflect.ValueOf(tree))
tree.Traverse(BEGIN_SORT)
fmt.Print("\n")
tree.Traverse(MIDDLE_SORT)
fmt.Print("\n")
tree.Traverse(END_SORT)
fmt.Print("\n")
}