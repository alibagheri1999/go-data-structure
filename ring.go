package main

import (
	"container/ring"
	"log"
)

func main() {
	r := ring.New(5)
	current := r
	for i := 0; i < r.Len(); i++ {
		current.Value = i
		current = current.Next()
	}

	r = r.Move(2)

	r.Do(func(p any) {
		if p != nil {
			log.Println(p.(int))
		}
	})
}
