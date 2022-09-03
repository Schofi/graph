package BFS

import (
	"container/list"
	"fmt"
)

func main() {
	vals := []int{1, 2, 3, 4}
	fifo := list.New()

	for _, val := range vals {
		fifo.PushBack(val)
	}

	for val := fifo.Front(); val != nil; val = val.Next() {
		fmt.Println(val.Value)
	}
}
