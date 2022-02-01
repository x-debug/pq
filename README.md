![PriorityQueue](./logo.png)

An implementation of priority queue,ported from the [beanstalkd](https://github.com/beanstalkd/beanstalkd/blob/master/heap.c)

## Features

- Push/Pop item
- Remove item on anywhere

## Installing

To start using `pq`, install Go and run `go get`:

```sh
$ go get -u https://github.com/x-debug/pq
```

This will retrieve the library.

## Example
You have to define the item structure first
```go
type MinIntItem struct {
    value int
}

func (item MinIntItem) Less(other Item) bool {
    return item.value > other.(MinIntItem).value
}
```

Push item to PQ
```go
pq := NewPriorityQueue()

pq.Push(MinIntItem{value: 3})
pq.Push(MinIntItem{value: 9})
pq.Push(MinIntItem{value: 1})
```

Pop item from PQ
```go
item := pq.Pop().(MinIntItem)
```

Or remove item from PQ
```go
item := pq.Remove(2).(MinIntItem)
```

more examples see pq_test.go

## References
[CMUQ 15-121 Priority Queues and Heaps](https://www.cs.cmu.edu/~rdriley/121/notes/heaps.html)

## License

`pq` source code is available under the MIT [License](/LICENSE).