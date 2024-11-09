# Binary Search Tree (BST) in Go

This is a generic implementation of a Binary Search Tree (BST) in Go, using generics to allow flexibility for different ordered types. This implementation includes basic operations such as insertion and search, as well as an in-order string representation of the BST.

## Features

- **Insert:** Adds a new node to the tree with the specified value.
- **Find:** Searches for a node with a given value and returns it if found.
- **String Representation:** Prints the tree values in sorted order.

# Min-Heap in Go

This is an implementation of a Min-Heap (priority queue) in Go using the `container/heap` package. This Min-Heap allows efficient insertion and retrieval of the smallest element, supporting operations to maintain a sorted order from minimum to maximum.

## Features

- **Push:** Adds a new element to the heap while maintaining the min-heap property.
- **Pop:** Removes and returns the smallest element from the heap.
- **Peek (Min Element):** Accesses the smallest element without removing it.
- **Initialization:** Initializes a slice as a min-heap to ensure proper structure before use.


# Linked List in Go

This project implements a simple singly linked list in Go, providing basic list operations such as adding, deleting, retrieving nodes, and additional functionalities like swapping pairs and reversing nodes in groups.

## Features

- **Add:** Adds a new node to the end of the list.
- **Get:** Retrieves a node by value.
- **GetByIndex:** Retrieves a node by index.
- **Delete:** Removes a node by value.
- **DeleteByIndex:** Removes a node by index.
- **GetList:** Returns all values in the list as a slice.
- **SwapPairs:** Swaps every two adjacent nodes in the list.
- **ReverseKGroup:** Reverses nodes in groups of a specified size `k`.

# Generic Queue in Go

This project implements a generic queue in Go using Go generics, allowing enqueue and dequeue operations for any ordered type. The queue follows the FIFO (First-In-First-Out) principle.

## Features

- **Enqueue:** Adds an item to the end of the queue.
- **Dequeue:** Removes and returns the item at the front of the queue. Returns an error if the queue is empty.
- **IsEmpty:** Checks if the queue is empty.

# Queue Implementation Using `container/list` in Go

This project demonstrates a simple queue implementation in Go using the `container/list` package. The `container/list` package provides a doubly linked list, which allows efficient enqueue and dequeue operations at both ends of the list. This example enqueues items and then dequeues them in FIFO (First-In-First-Out) order.

## Features

- **Enqueue:** Adds an item to the end of the queue.
- **Dequeue:** Removes and retrieves the item from the front of the queue.
# Ring Buffer Implementation Using `container/ring` in Go

This project demonstrates a simple ring buffer (circular buffer) implementation in Go using the `container/ring` package. A ring buffer is a fixed-size, circular data structure where elements wrap around to the beginning when the end of the buffer is reached. This example initializes a ring, assigns values, and iterates over it with an adjustable starting position.

## Features

- **Initialization:** Creates a fixed-size ring buffer.
- **Assign Values:** Fills the ring buffer with values.
- **Iteration (Do):** Iterates over each element in the ring buffer.
- **Move:** Shifts the starting position within the ring, allowing flexible traversal.
# Generic Stack in Go

This project implements a generic stack in Go using Go generics, providing basic stack operations such as push, pop, and checking if the stack is empty. The stack follows the LIFO (Last-In-First-Out) principle.

## Features

- **Push:** Adds an item to the top of the stack.
- **Pop:** Removes and returns the item from the top of the stack. Returns an error if the stack is empty.
- **IsEmpty:** Checks if the stack is empty.