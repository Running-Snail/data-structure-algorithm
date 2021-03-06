package bst

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
)

func treeFromList(l []int) *Node {
	root := &Node{
		Val: l[0],
	}
	for i := 1; i < len(l); i++ {
		root.Insert(l[i])
	}
	return root
}

func TestInorderString(t *testing.T) {
	fmt.Println("testing inorder")
	root := &Node{
		Val: 3,
	}
	left := &Node{
		Val: 2,
	}
	right := &Node{
		Val: 4,
	}
	root.Left = left
	root.Right = right
	inorder := root.InorderString(" ")
	if inorder != "2 3 4" {
		t.Error("expect inorder 2 3 4 but get")
		t.Error(inorder)
	} else {
		fmt.Println("PASS")
	}
}

func TestInsert(t *testing.T) {
	fmt.Println("testing insert")
	root := &Node{
		Val: 20,
	}
	root.Insert(10).Insert(3).Insert(20).Insert(2).Insert(30).Insert(40)
	inorder := root.InorderString(" ")
	if inorder != "2 3 10 20 20 30 40" {
		t.Error("expect inorder 2 3 10 20 30 40 but get")
		t.Error(inorder)
	} else {
		fmt.Println("PASS")
	}
}

func TestMin(t *testing.T) {
	fmt.Println("testing min")
	root := treeFromList([]int{30, 20, 6, 7, 9, 2, 40})
	min := root.Min()
	if min != 2 {
		t.Error("expect min 2 but get")
		t.Error(min)
	} else {
		fmt.Println("PASS")
	}
}

func TestMax(t *testing.T) {
	fmt.Println("testing max")
	root := treeFromList([]int{30, 20, 6, 7, 9, 2, 40})
	max := root.Max()
	if max != 40 {
		t.Error("expect max 40 but get")
		t.Error(max)
	} else {
		fmt.Println("PASS")
	}
}

func TestSortedSlice(t *testing.T) {
	fmt.Println("testing sorted")
	root := treeFromList([]int{30, 20, 6, 7, 9, 2, 40})
	sorted := root.SortedSlice()
	if fmt.Sprint(sorted) != "[2 6 7 9 20 30 40]" {
		t.Error("expect sorted result is 2 6 7 9 20 30 40 but get")
		t.Error(sorted)
	} else {
		fmt.Println("PASS")
	}
}

func TestDelete(t *testing.T) {
	fmt.Println("testing delete")
	root := NewFromInts([]int{30, 20, 6, 7, 9, 2, 40})
	root.Delete(20).Delete(70).Delete(6)
	inorder := root.InorderString(" ")
	if inorder != "2 7 9 30 40" {
		t.Error("expect delete tree inorder is 2 7 9 30 40 but get")
		t.Error(inorder)
	}

	root = NewFromInts([]int{6, 3, 9, 4, 5})
	root = root.Delete(6)
	inorder = root.InorderString(" ")
	if inorder != "3 4 5 9" {
		t.Error("expect result is 3 4 5 9 but get")
		t.Error(inorder)
	}

	root = NewFromInts([]int{1})
	root = root.Delete(1)
	if root != nil {
		t.Error("expect to be a empty tree")
	}
}

func insertAndPrint(root *Node, v int) {
	root.Insert(v)
	fmt.Println(root.InorderString(","))
}

func deleteAndPrint(root *Node, v int) {
	root.Delete(v)
	fmt.Println(root.InorderString(","))
}

func TestThreadSafe(t *testing.T) {
	fmt.Println("testing thread safe")
	root := &Node{Val: 0}
	n := 10
	wg := sync.WaitGroup{}
	wg.Add(n + n)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			v := rand.Intn(200)
			insertAndPrint(root, v)
		}()
		go func() {
			defer wg.Done()
			v := rand.Intn(200)
			deleteAndPrint(root, v)
		}()
	}
	wg.Wait()
}

func TestNum(t *testing.T) {
	fmt.Println("testing num")
	root := &Node{}
	n := 400
	for i := 1; i < n; i++ {
		root.Insert(rand.Intn(200))
	}
	m := root.Num()
	if m != n {
		t.Errorf("expect get num is %d but get %d\n", n, m)
	} else {
		fmt.Println("PASS")
	}
}

func TestHas(t *testing.T) {
	fmt.Println("testing has")
	root := NewFromInts([]int{2, 5, 1, 0, 5, 4, 7})
	if root.Has(2) == false {
		t.Error("expect to have 2")
	}

	if root.Has(8) == true {
		t.Error("expect not to have 8")
	}
}

func TestDeleteAll(t *testing.T) {
	fmt.Println("testing delete all")
	root := NewFromInts([]int{10, 10, 10, 20, 30, 5, 4})
	root = root.DeleteAll(10)
	inorder := root.InorderString(" ")
	if inorder != "4 5 20 30" {
		t.Error("delete all error")
	}
}

func benchmarkInsert(n int, b *testing.B) {
	l := make([]int, n)
	for i := 0; i < n; i++ {
		l[i] = rand.Intn(200)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		NewFromInts(l)
	}
}

func BenchmarkInsert100(b *testing.B) {
	benchmarkInsert(100, b)
}

func BenchmarkInsert10000(b *testing.B) {
	benchmarkInsert(10000, b)
}

func BenchmarkInsert100000(b *testing.B) {
	benchmarkInsert(100000, b)
}
