# queue

Library of generic queue data structures for Go.

* [linked-queue](./linkedqueue/)
* [slice-queue](./slicequeue/)
* [circular-queue](./circularqueue/)
  
## Install

```Go
$ go get github.com/golang-ds/queue
```

## Linked Queue

### Import

```Go
import queue "github.com/golang-ds/queue/linkedqueue"
```

### Use

```Go
q := queue.New[int]()
q.Enqueue(1)
```

## Slice Queue

### Import

```Go
import queue "github.com/golang-ds/queue/slicequeue"
```

### Use

```Go
q := queue.New[int]()
q.Enqueue(1)
```

## Circular Queue

### Import

```Go
import queue "github.com/golang-ds/queue/circularqueue"
```

### Use

```Go
q := queue.New[int]()
q.Enqueue(1)
```
