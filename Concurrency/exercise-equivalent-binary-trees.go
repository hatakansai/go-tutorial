package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

func Walk(t *tree.Tree, ch chan int) {
	walkHelper(t, ch)
	close(ch)
}

func walkHelper(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	walkHelper(t.Left, ch)
	ch <- t.Value
	walkHelper(t.Right, ch)
}

func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2

		if !ok1 || !ok2 {
			return ok1 == ok2
		}

		if v1 != v2 {
			return false
		}
	}
}

func main() {
	ch := make(chan int)
	go Walk(tree.New(1), ch)

	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(2), tree.New(2)))
	fmt.Println(Same(tree.New(3), tree.New(3)))

	fmt.Println(Same(tree.New(1), tree.New(2)))
	fmt.Println(Same(tree.New(1), tree.New(3)))
	fmt.Println(Same(tree.New(1), tree.New(4)))

	fmt.Println(Same(tree.New(2), tree.New(1)))
	fmt.Println(Same(tree.New(3), tree.New(1)))
	fmt.Println(Same(tree.New(4), tree.New(1)))
}
