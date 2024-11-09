package main

import (
	"container/list"
	"log"
)

func main() {
	queue := list.New()
	queue.PushBack(1)
	queue.PushBack(2)
	queue.PushBack(3)

	for queue.Len() > 0 {
		item := queue.Front()
		log.Println(item.Value.(int))
		queue.Remove(item)
	}
}
