package linear

import (
	"container/list"
	"fmt"
)

// listExample shows how lists work
func listExample() {
	var intList list.List
	intList.PushBack(11)
	intList.PushBack(12)
	intList.PushBack(13)

	for e := intList.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value.(int))
	}
}
